package defender

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"bytes"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type payLoad struct{

	name string `json:"name"`
	network string `json:"network"`
	minBalance string `json:"minBalance"`
	pendingTxCost string `json:"pendingTxCost"`
}


func resourceCrud() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDefenderCreate,
		ReadContext:   resourceDefenderRead,
		UpdateContext: resourceDefenderUpdate,
		DeleteContext: resourceDefenderDelete,
		Schema: map[string]*schema.Schema{
			"apikey": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				DefaultFunc: schema.EnvDefaultFunc("DEFENDER_APIKEY", nil),
			},
			"token":{
				Type: 	     schema.TypeString,
				Optional:    true,
				Computed:    true,
				DefaultFunc: schema.EnvDefaultFunc("DEFENDER_TOKEN", nil),
			},
			"relay": &schema.Schema{
				Type:     schema.TypeSet,
				MaxItems: 1,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"network": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"minbalance": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						/*"policies": &schema.Schema{
							Type:     Schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"eip1559pricing": &schema.Schema{
										Type:     schema.TypeBool,
										Required: true,
									},
								},
							},
						},*/
						"pendingtxcost": &schema.Schema{
							Type: schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}//End Resource return
}//End Resource func

func resourceDefenderCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	client := &http.Client{Timeout: 10 * time.Second}
        apikey := d.Get("apikey").(string)
	token  := d.Get("token").(string)

	var diags diag.Diagnostics

	relKey := d.Get("relay").(*schema.Set).List()
	dat,_:= json.Marshal(relKey[0])
	arDat := string(dat)
	var result2 map[string]interface{}
	err := json.Unmarshal([]byte(arDat),&result2)
	if err != nil{
		fmt.Println(err)
	}

	pay := payLoad{}
	pay.name = fmt.Sprint(result2["name"])
	pay.network = fmt.Sprint(result2["network"])
	pay.minBalance = fmt.Sprint(result2["minbalance"])
	pay.pendingTxCost = fmt.Sprint(result2["pendingtxcost"])

	pay2 := `{"name":"` + pay.name + `","network":"` + pay.network + `","minBalance":"` + pay.minBalance + `","pendingTxCost":"` + pay.pendingTxCost + `"}`


	jsonBody := []byte(pay2)
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest("POST", "https://defender-api.openzeppelin.com/relayer/relayers", bodyReader)
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
	reqBody, err := ioutil.ReadAll(r.Body)
	var result map[string]interface{}
	err = json.Unmarshal([]byte(reqBody), &result)
	if err != nil {
		return diag.FromErr(err)
	}
	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
func resourceDefenderRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type   
	var diags diag.Diagnostics
	return diags
}
func resourceDefenderUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}
func resourceDefenderDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}
