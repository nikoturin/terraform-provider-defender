package defender

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"strconv"
	"time"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDefender() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceDefenderRead,
		Schema: map[string]*schema.Schema{
			"apikey": {
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
				DefaultFunc: schema.EnvDefaultFunc("DEFENDER_APIKEY",nil),
			},
			"token":{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
				DefaultFunc: schema.EnvDefaultFunc("DEFENDER_TOKEN",nil),
			},
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"relayerid": &schema.Schema{
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
						"createdat": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"paused": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"policies": &schema.Schema{
							Type:	schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeBool,
							},
						},
						"minbalance": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"pendingtxcost": &schema.Schema{
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
	apikey := d.Get("apikey").(string)
	token  := d.Get("token").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if(apikey == "") && (token == ""){
		return nil
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/summary", "https://defender-api.openzeppelin.com/relayer/relayers"), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", apikey)
	req.Header.Set("Authorization", token)

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()
	reqBody, err:= ioutil.ReadAll(r.Body)
	var result map[string]interface{}
	err = json.Unmarshal([]byte(reqBody),&result)
	if err !=nil{
		return diag.FromErr(err)
	}
	mapA := result["items"]
	mapB, _ := json.Marshal(mapA)
	summary := make([]map[string]interface{}, 0)
	err = json.Unmarshal([]byte(strings.ToLower(string(mapB))),&summary)

	if err !=nil{
		return diag.FromErr(err)
	}
	if err := d.Set("items",summary); err != nil {
		return diag.FromErr(err)
	}
	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
