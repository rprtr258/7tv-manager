package emoteset

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/samber/lo"

	"github.com/rprtr258/seventv-tf-provider/internal/api"
)

func New() *schema.Resource {
	return &schema.Resource{
		CreateContext: create,
		ReadContext:   read,
		UpdateContext: update,
		DeleteContext: delete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"emotes": {
				Type:     schema.TypeMap,
				Required: true,
				Elem: &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					Computed: true,
				},
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(api.Api)
	name := d.Get("name").(string)
	emotesAny := d.Get("emotes")

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	// diags = append(diags, read(ctx, d, m)...)

	id, err := c.EmoteSet().Create(name)
	if err != nil {
		return diag.FromErr(err)
	}

	emotes := mapEmotesList(emotesAny)
	for emoteID, emoteName := range emotes {
		if err := c.Emote().AddToSet(id, emoteID, emoteName); err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId(id)

	return diags
}

func read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(api.Api)
	id := d.Id()

	var diags diag.Diagnostics

	emoteSet, err := c.EmoteSet().Read(id)
	if err != nil {
		return diag.FromErr(err)
	}

	emotes := map[string]any{}
	for emoteName, emoteID := range emoteSet.Emotes {
		emotes[emoteID] = emoteName
	}

	d.SetId(emoteSet.ID)

	if err := d.Set("name", emoteSet.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("emotes", emotes); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(api.Api)
	id := d.Id()
	name := d.Get("name").(string)
	oldEmotesAny, newEmotesAny := d.GetChange("emotes")

	// TODO: name changed?
	if !d.HasChange("emotes") {
		return read(ctx, d, m)
	}

	_, err := c.EmoteSet().UpdateName(id, name)
	if err != nil {
		return diag.FromErr(err)
	}

	oldEmotes := mapEmotesList(oldEmotesAny)
	newEmotes := mapEmotesList(newEmotesAny)

	for emoteID, update := range diffSets(oldEmotes, newEmotes) {
		switch {
		case update.oldName != nil && update.newName == nil:
			if err := c.Emote().RemoveFromSet(id, emoteID); err != nil {
				return diag.FromErr(err)
			}
		case update.oldName == nil && update.newName != nil:
			if err := c.Emote().AddToSet(id, emoteID, *update.newName); err != nil {
				return diag.FromErr(err)
			}
		case update.oldName != nil && update.newName != nil:
			if *update.oldName == *update.newName {
				continue
			}

			if err := c.Emote().UpdateName(id, emoteID, *update.newName); err != nil {
				return diag.FromErr(err)
			}
		case update.oldName == nil && update.newName == nil:
			return diag.Errorf("id=%s, no old, nor new names", emoteID)
		}
	}

	return read(ctx, d, m)
}

type updateKind struct {
	oldName *string
	newName *string
}

func diffSets(old, new map[string]string) map[string]*updateKind {
	updates := map[string]*updateKind{}

	for emoteID, emoteName := range old {
		emoteName := emoteName
		updates[emoteID] = &updateKind{
			oldName: &emoteName,
		}
	}

	for emoteID, emoteName := range new {
		emoteName := emoteName
		if _, ok := updates[emoteID]; !ok {
			updates[emoteID] = &updateKind{
				newName: &emoteName,
			}
		} else {
			updates[emoteID].newName = &emoteName
		}
	}

	return updates
}

func mapEmotesList(emotes any) map[string]string {
	return lo.MapEntries(
		emotes.(map[string]any),
		func(emoteID string, emoteName any) (string, string) {
			return emoteID, emoteName.(string)
		},
	)
}

func delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(api.Api)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	err := c.EmoteSet().Delete(id)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
