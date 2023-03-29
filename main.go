package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CSC_API_KEY", nil),
				Description: "API Key for CSC Domain Manager API",
			},
			"bearer_token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CSC_BEARER_TOKEN", nil),
				Description: "Bearer Token for CSC Domain Manager API",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"csc_domain": resourceCscDomain(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"csc_domain": dataSourceCscDomain(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	apiKey := d.Get("api_key").(string)
	bearerToken := d.Get("bearer_token").(string)
	return cscclient.NewClient(apiKey, bearerToken), nil
}

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: Provider,
	})
}
