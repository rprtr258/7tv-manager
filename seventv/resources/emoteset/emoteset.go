package emoteset

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
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
			"id": {
				Type: schema.TypeString,
				// Required: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"emotes": {
				Type:     schema.TypeList,
				Required: true, // TODO: is it?
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
						},
					},
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

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	name := d.Get("name").(string)
	// items := d.Get("emotes").([]interface{})
	// emoteSet := []api.EmoteSet{}

	// for _, item := range items {
	// 	i := item.(map[string]interface{})

	// 	emoteID := i["id"].(string)
	// 	emoteName := i["name"].(string)

	// 	oi := api.EmoteSet{
	// 		ID: "",
	// 		Emotes: []api.Emote{
	// 			{
	// 				ID:   emoteID,
	// 				Name: emoteName,
	// 			},
	// 		},
	// 	}

	// 	emoteSet = append(emoteSet, oi)
	// }

	id, err := c.CreateEmoteSet(name)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(id)

	read(ctx, d, m)

	return diags
}

func read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(api.Api)
	id := d.Id()

	var diags diag.Diagnostics

	emoteSet, err := c.GetEmoteSet(id)
	if err != nil {
		return diag.FromErr(err)
	}

	tflog.Error(ctx, "Unrecognized API response body", map[string]any{
		"data": emoteSet,
	})
	emotes := []any{}
	for _, emote := range emoteSet.Emotes {
		emotes = append(emotes, map[string]any{
			"id":   emote.ID,
			"name": emote.Name,
		})
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

	if !d.HasChange("emotes") {
		return read(ctx, d, m)
	}

	emotes := d.Get("emotes").([]any)
	_ = lo.Map(
		emotes,
		func(emote any, _ int) api.Emote {
			return api.Emote{}
		},
	)

	_, err := c.UpdateEmoteSet(id, name)
	if err != nil {
		return diag.FromErr(err)
	}

	return read(ctx, d, m)
}

func delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(api.Api)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	err := c.DeleteEmoteSet(id)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
