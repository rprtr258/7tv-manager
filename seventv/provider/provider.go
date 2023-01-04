package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/rprtr258/seventv-tf-provider/internal/api"
	"github.com/rprtr258/seventv-tf-provider/seventv/resources/emoteset"
)

// New - 7tv emotes terraform provider
func New() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Optional: true,
				// DefaultFunc: schema.EnvDefaultFunc("7TV_USERNAME", nil),
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
				// DefaultFunc: schema.EnvDefaultFunc("7TV_PASSWORD", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"emoteset": emoteset.New(),
		},
		// DataSourcesMap: map[string]*schema.Resource{
		// 	"hashicups_coffees":     dataSourceCoffees(),
		// 	"hashicups_order":       dataSourceOrder(),
		// 	"hashicups_ingredients": dataSourceIngredients(),
		// },
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if (username != "") && (password != "") {
		c, err := api.NewClient(&username, &password)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create HashiCups client",
				Detail:   "Unable to authenticate user for authenticated HashiCups client",
			})

			return nil, diags
		}

		return c, diags
	}

	c, err := api.NewClient(nil, nil)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create HashiCups client",
			Detail:   "Unable to create anonymous HashiCups client",
		})
		return nil, diags
	}

	return c, diags
}
