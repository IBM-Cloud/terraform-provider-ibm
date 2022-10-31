// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package directlink

import (
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceIBMDLASPrepends() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMDLASPrependsRead,
		Schema: map[string]*schema.Schema{
			dlGatewayId: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Direct Link gateway identifier",
			},

			dlAsPrepends: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of AS Prepend configuration information",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						dlCreatedAt: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time AS Prepend was created",
						},
						ID: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time AS Prepend was created",
						},
						dlLength: {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of times the ASN to appended to the AS Path",
						},
						dlPolicy: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Route type this AS Prepend applies to",
						},
						dlPrefix: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Comma separated list of prefixes this AS Prepend applies to. Maximum of 10 prefixes. If not specified, this AS Prepend applies to all prefixes.",
						},
						dlSpecificPrefixes: {
							Type:        schema.TypeList,
							Description: "Array of prefixes this AS Prepend applies to",
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						dlUpdatedAt: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time AS Prepend was updated",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMDLASPrependsRead(d *schema.ResourceData, meta interface{}) error {
	directLink, err := mydlClient(meta)
	gateway_id := d.Get(dlGatewayId).(string)

	if err != nil {
		return err
	}

	listGatewayAsPrependsOptions := directLink.NewListGatewayAsPrependsOptions(gateway_id)
	result, response, operationErr := directLink.ListGatewayAsPrepends(listGatewayAsPrependsOptions)
	if operationErr != nil {
		log.Println("[WARN] Error listing dl Gateway AS Prepends ", response, err)
		return err
	}

	asPrependList := make([]map[string]interface{}, 0)
	if len(result.AsPrepends) > 0 {

		for _, asPrepend := range result.AsPrepends {
			asPrependItem := map[string]interface{}{}
			asPrependItem[dlResourceId] = asPrepend.ID
			asPrependItem[dlLength] = asPrepend.Length
			asPrependItem[dlPrefix] = asPrepend.Prefix
			asPrependItem[dlSpecificPrefixes] = asPrepend.SpecifiedPrefixes
			asPrependItem[dlPolicy] = asPrepend.Policy
			asPrependItem[dlCreatedAt] = asPrepend.CreatedAt.String()
			asPrependItem[dlUpdatedAt] = asPrepend.UpdatedAt.String()

			asPrependList = append(asPrependList, asPrependItem)
		}
	}

	d.SetId(dataSourceIBMDLASPrependsID(d))
	d.Set(dlAsPrepends, asPrependList)
	return nil
}

// dataSourceIBMDLGatewaysID returns a reasonable ID for a direct link gateways list.
func dataSourceIBMDLASPrependsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
