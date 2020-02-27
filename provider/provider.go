package provider

import (
	"errors"

	"github.com/onelogin/onelogin-go-sdk/pkg/client"
	"github.com/onelogin/onelogin-terraform-provider/resources"

	"github.com/hashicorp/terraform/helper/schema"
)

var (
	errClientCredentials = errors.New("client_id or client_sercret or region missing")
)

// Provider creates a new provider with all the neccessary configurations.
// It returns a pointer to the created provider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"client_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"client_secret": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  client.USRegion,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"onelogin_apps": resources.OneloginApps(),
		},
		ConfigureFunc: configProvider,
	}
}

// configProvider configures the provider, and if successful, it returns
// an interface containing the api client.
func configProvider(d *schema.ResourceData) (interface{}, error) {
	clientID := d.Get("client_id").(string)
	clientSecret := d.Get("client_secret").(string)
	region := d.Get("region").(string)
	timeout := client.DefaultTimeout

	oneloginClientConfig, err := client.NewConfig(
		clientID,
		clientSecret,
		region,
		timeout,
	)
	if err == nil {
		return nil, err
	}

	oneloginClient := client.NewClient(oneloginClientConfig)

	return oneloginClient, nil
}