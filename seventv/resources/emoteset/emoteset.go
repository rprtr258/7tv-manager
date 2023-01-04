package emoteset

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

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
			"emotes": {
				Type:     schema.TypeList,
				Required: true,
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

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Id()

	emoteSet, err := c.GetEmoteSet(id)
	if err != nil {
		return diag.FromErr(err)
	}

	orderItems := mapAnyEmotes(emoteSet.Emotes)
	if err := d.Set("emotes", orderItems); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// c := m.(api.Api)

	// orderID := d.Id()

	if !d.HasChange("emotes") {
		return read(ctx, d, m)
	}

	// items := d.Get("items").([]interface{})
	// ois := []hc.OrderItem{}

	// for _, item := range items {
	// 	i := item.(map[string]interface{})

	// 	co := i["coffee"].([]interface{})[0]
	// 	coffee := co.(map[string]interface{})

	// 	oi := hc.OrderItem{
	// 		Coffee: hc.Coffee{
	// 			ID: coffee["id"].(int),
	// 		},
	// 		Quantity: i["quantity"].(int),
	// 	}
	// 	ois = append(ois, oi)
	// }

	// _, err := c.UpdateOrder(orderID, ois, &c.Token)
	// if err != nil {
	// 	return diag.FromErr(err)
	// }

	d.Set("last_updated", time.Now().Format(time.RFC850))

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

func mapAnyEmotes(emotes []api.Emote) []interface{} {
	if emotes == nil {
		return make([]interface{}, 0)
	}

	ois := make([]interface{}, len(emotes))

	for i, emote := range emotes {
		ois[i] = map[string]interface{}{
			"id":   emote.ID,
			"name": emote.Name,
		}
	}

	return ois
}
