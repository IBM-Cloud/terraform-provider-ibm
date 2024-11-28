package db2

import (
	"context"
	"fmt"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/cloud-db2-go-sdk/db2saasv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func DataSourceIBMDB2WhitelistIPs() *schema.Resource {
	return &schema.Resource{
		ReadContext: DataSourceIBMDB2WhitelistIPsRead,

		Schema: map[string]*schema.Schema{
			"ip_addresses": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IP address, in IPv4 format",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the IP address",
						},
					},
				},
			},
		},
	}
}

func DataSourceIBMDB2WhitelistIPsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	db2SaasV1Client, err := meta.(conns.ClientSession).DB2SaasV1()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_db2_whitelist_ips", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getWhitelistIPsOptions := &db2saasv1.GetDb2SaasWhitelistOptions{}

	getWhitelistIPsOptions.SetXDeploymentID(d.Get("x_deployment_id").(string))

	result, response, err := db2SaasV1Client.GetDb2SaasWhitelistWithContext(context, getWhitelistIPsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDb2SaasWhitelistWithContext failed: %s\n%s", err.Error(), response), "(Data) ibm_db2_whitelist_ips", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	var matchWhitelistIPs []db2saasv1.IpAddress

	result.IpAddresses = matchWhitelistIPs

	result2 := []map[string]interface{}{}

	if result.IpAddresses != nil {
		for _, modelItem := range result.IpAddresses {
			modelMap, err := DataSourceIBMDb2WhitelistIPsToMap(&modelItem)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_db2_whitelist_ips", "read")
				return tfErr.GetDiag()
			}
			result2 = append(result2, modelMap)
		}
	}

	if err = d.Set("ip_addresses", result2); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting whitelist IPs: %s", err), "(Data) ibm_db2_whitelist_ips", "read")
		return tfErr.GetDiag()
	}

	return nil
}

func DataSourceIBMDb2WhitelistIPsToMap(model *db2saasv1.IpAddress) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Address != nil {
		modelMap["address"] = *model.Address
	}

	if model.Description != nil {
		modelMap["description"] = *model.Description
	}

	return modelMap, nil
}
