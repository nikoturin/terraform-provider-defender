package defender

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
	ResourcesMap: map[string]*schema.Resource{
		"defender_modules": resourceCrud(),
	},
	DataSourcesMap: map[string]*schema.Resource{
		"defender_summary": dataSourceDefender(),
	},
    }
}
