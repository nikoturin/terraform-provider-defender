package defender

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDefender() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDefenderRead,
		Schema: map[string]*schema.Schema{
			"summary": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"relayer_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"network": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"create_at": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"paused": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	} //EndReturnResource
} //EndFunction

func dataSourceDefenderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/summary", "https://defender-api.openzeppelin.com/relayer/relayers"), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Api-Key", "")
	req.Header.Add("Authorization", "Bearer")

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	summary := make([]map[string]interface{}, 0)
	err = json.NewDecoder(r.Body).Decode(&summary)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("summary", summary); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags

}
