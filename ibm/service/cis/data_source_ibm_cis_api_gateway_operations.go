// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"context"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/networking-go-sdk/aisecurityforappsv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	ibmCISApiGatewayOperationsDS          = "ibm_cis_api_gateway_operations"
	cisApiGatewayOperationsList           = "operations"
	cisApiGatewayOperationsFilterHost     = "host"
	cisApiGatewayOperationsFilterMethod   = "method"
	cisApiGatewayOperationsFilterEndpoint = "endpoint"
	cisApiGatewayOperationsFilterLabels   = "labels"
	cisApiGatewayOperationsDirection      = "direction"
	cisApiGatewayOperationsOrder          = "order"
	cisApiGatewayOperationsPage           = "page"
	cisApiGatewayOperationsPerPage        = "per_page"
	cisApiGatewayOperationsTotalCount     = "total_count"
	cisApiGatewayOperationsLastUpdated    = "last_updated"
)

// DataSourceIBMCISApiGatewayOperations defines the schema for listing API Gateway operations.
func DataSourceIBMCISApiGatewayOperations() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMCISApiGatewayOperationsRead,
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "CIS instance CRN",
				ValidateFunc: validate.InvokeDataSourceValidator(
					ibmCISApiGatewayOperationsDS,
					"cis_id"),
			},
			cisDomainID: {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "Associated CIS domain ID",
				DiffSuppressFunc: suppressDomainIDDiff,
			},
			cisApiGatewayOperationsFilterHost: {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "Filter results to only include the specified hosts",
			},
			cisApiGatewayOperationsFilterMethod: {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "Filter results to only include the specified HTTP methods",
			},
			cisApiGatewayOperationsFilterEndpoint: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter results to only include endpoints containing this pattern",
			},
			cisApiGatewayOperationsFilterLabels: {
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Description: "Filter results to only include operations with these labels applied",
			},
			cisApiGatewayOperationsDirection: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Direction to order results: asc or desc",
			},
			cisApiGatewayOperationsOrder: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Field to order results by (endpoint, host, last_updated, method)",
			},
			cisApiGatewayOperationsPage: {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Page number of paginated results",
			},
			cisApiGatewayOperationsPerPage: {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Maximum number of results per page",
			},
			cisApiGatewayOperationsList: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of API Gateway operations",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						cisApiGatewayOperationID: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "UUID of the operation",
						},
						cisApiGatewayOperationMethod: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "HTTP method",
						},
						cisApiGatewayOperationHost: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "API host",
						},
						cisApiGatewayOperationEndpt: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Endpoint path",
						},
						cisApiGatewayOperationsLastUpdated: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp of last update",
						},
					},
				},
			},
			cisApiGatewayOperationsTotalCount: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Total count of operations matching the query",
			},
		},
	}
}

// DataSourceIBMCISApiGatewayOperationsValidator returns the ResourceValidator for the data source.
func DataSourceIBMCISApiGatewayOperationsValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "cis_id",
			ValidateFunctionIdentifier: validate.ValidateCloudData,
			Type:                       validate.TypeString,
			CloudDataType:              "resource_instance",
			CloudDataRange:             []string{"service:internet-svcs"},
			Required:                   true,
		})
	return &validate.ResourceValidator{
		ResourceName: ibmCISApiGatewayOperationsDS,
		Schema:       validateSchema,
	}
}

func dataSourceIBMCISApiGatewayOperationsRead(d *schema.ResourceData, meta interface{}) error {
	cisClient, err := meta.(conns.ClientSession).CisAiSecurityForAppsSession()
	if err != nil {
		return err
	}

	crn := d.Get(cisID).(string)
	domainID := d.Get(cisDomainID).(string)
	zoneID, _, _ := flex.ConvertTftoCisTwoVar(domainID)
	if zoneID == "" {
		zoneID = domainID
	}
	cisClient.Crn = core.StringPtr(crn)
	cisClient.ZoneIdentifier = core.StringPtr(zoneID)

	opt := cisClient.NewListZoneApiGatewayOperationsOptions()

	if v, ok := d.GetOk(cisApiGatewayOperationsFilterHost); ok {
		opt.SetHost(flex.ExpandStringList(v.([]interface{})))
	}
	if v, ok := d.GetOk(cisApiGatewayOperationsFilterMethod); ok {
		opt.SetMethod(flex.ExpandStringList(v.([]interface{})))
	}
	if v, ok := d.GetOk(cisApiGatewayOperationsFilterEndpoint); ok {
		opt.SetEndpoint(v.(string))
	}
	if v, ok := d.GetOk(cisApiGatewayOperationsFilterLabels); ok {
		opt.SetLabels(flex.ExpandStringList(v.([]interface{})))
	}
	if v, ok := d.GetOk(cisApiGatewayOperationsDirection); ok {
		opt.SetDirection(v.(string))
	}
	if v, ok := d.GetOk(cisApiGatewayOperationsOrder); ok {
		opt.SetOrder(v.(string))
	}
	if v, ok := d.GetOk(cisApiGatewayOperationsPage); ok {
		opt.SetPage(int64(v.(int)))
	}
	if v, ok := d.GetOk(cisApiGatewayOperationsPerPage); ok {
		opt.SetPerPage(int64(v.(int)))
	}

	result, response, err := cisClient.ListZoneApiGatewayOperationsWithContext(context.Background(), opt)
	if err != nil {
		log.Printf("[ERROR] ListZoneApiGatewayOperations failed: %v", response)
		return err
	}

	d.SetId(flex.ConvertCisToTfTwoVar(zoneID, crn))
	d.Set(cisID, crn)
	d.Set(cisDomainID, zoneID)

	if result != nil {
		if result.ResultInfo != nil && result.ResultInfo.TotalCount != nil {
			d.Set(cisApiGatewayOperationsTotalCount, int(*result.ResultInfo.TotalCount))
		}
		if result.Result != nil {
			d.Set(cisApiGatewayOperationsList, flattenZoneOperationsList(result.Result))
		}
	}
	return nil
}

func flattenZoneOperationsList(items []aisecurityforappsv1.ApiGatewayOperationsListRespResultItem) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(items))
	for _, item := range items {
		m := map[string]interface{}{}
		if item.OperationID != nil {
			m[cisApiGatewayOperationID] = *item.OperationID
		}
		if item.Method != nil {
			m[cisApiGatewayOperationMethod] = *item.Method
		}
		if item.Host != nil {
			m[cisApiGatewayOperationHost] = *item.Host
		}
		if item.Endpoint != nil {
			m[cisApiGatewayOperationEndpt] = *item.Endpoint
		}
		if item.LastUpdated != nil {
			m[cisApiGatewayOperationsLastUpdated] = item.LastUpdated.String()
		}
		result = append(result, m)
	}
	return result
}
