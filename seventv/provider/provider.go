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
			"seventv_emoteset": emoteset.New(),
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
	token := d.Get("token").(string)

	if token == "" {
		return nil, []diag.Diagnostic{{
			Severity: diag.Error,
			Summary:  "Invalid credentials",
			Detail:   "No token provided",
		}}
	}

	c, err := api.NewClient("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1IjoiNjE1MDZlYzM0M2IyZDlkYTBkMzI4ZGMwIiwidiI6MSwiaXNzIjoiN1RWLUFQSS1SRVNUIiwiZXhwIjoxNjgwODE1MjE4fQ.bHajGpk4LT5SuT1vqqpfu1fJznX4Gp3PRaad67NC9P4")
	if err != nil {
		return nil, []diag.Diagnostic{{
			Severity: diag.Error,
			Summary:  "Unable to create HashiCups client",
			Detail:   "Unable to authenticate user for authenticated HashiCups client",
		}}
	}

	return c, nil
}
