package seventv

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/rprtr258/seventv-tf-provider/internal/api"
)

func dataSourceOrder() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceEmoteSetRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// "emotes": &schema.Schema{
			// 	Type:     schema.TypeList,
			// 	Computed: true,
			// 	Elem: &schema.Resource{
			// 		Schema: map[string]*schema.Schema{
			// 			"coffee_id": &schema.Schema{
			// 				Type:     schema.TypeInt,
			// 				Computed: true,
			// 			},
			// 			"coffee_name": &schema.Schema{
			// 				Type:     schema.TypeString,
			// 				Computed: true,
			// 			},
			// 			"coffee_teaser": &schema.Schema{
			// 				Type:     schema.TypeString,
			// 				Computed: true,
			// 			},
			// 			"coffee_description": &schema.Schema{
			// 				Type:     schema.TypeString,
			// 				Computed: true,
			// 			},
			// 			"coffee_price": &schema.Schema{
			// 				Type:     schema.TypeInt,
			// 				Computed: true,
			// 			},
			// 			"coffee_image": &schema.Schema{
			// 				Type:     schema.TypeString,
			// 				Computed: true,
			// 			},
			// 			"quantity": &schema.Schema{
			// 				Type:     schema.TypeInt,
			// 				Computed: true,
			// 			},
			// 		},
			// 	},
			// },
		},
	}
}

func dataSourceEmoteSetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(api.Api)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	id := d.Get("id").(string)

	emoteSet, err := c.EmoteSet().Read(id)
	if err != nil {
		return diag.FromErr(err)
	}

	orderItems := flattenOrderItemsData(emoteSet)
	if err := d.Set("items", orderItems); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(id)

	return diags
}

func flattenOrderItemsData(emoteSet api.EmoteSet) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"coffee_id": emoteSet.ID,
		},
	}
}
