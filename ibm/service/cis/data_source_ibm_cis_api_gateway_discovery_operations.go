// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis

import (
	"context"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	ibmCISApiGatewayDiscoveryOperationsDS = "ibm_cis_api_gateway_discovery_operations_ds"

	cisApiGatewayDiscoveryOpsDiff      = "diff"
	cisApiGatewayDiscoveryOpsDirection = "direction"
	cisApiGatewayDiscoveryOpsOrder     = "order"
	cisApiGatewayDiscoveryOpsOrigin    = "origin"
	cisApiGatewayDiscoveryOpsState     = "state"
	cisApiGatewayDiscoveryOpsPage      = "page"
	cisApiGatewayDiscoveryOpsPerPage   = "per_page"
	cisApiGatewayDiscoveryOpsOps       = "operations"
	cisApiGatewayDiscoveryOpsTotalCnt  = "total_count"
)

// DataSourceIBMCISApiGatewayDiscoveryOperations returns the schema.Resource for the
// ibm_cis_api_gateway_discovery_operations data source.
func DataSourceIBMCISApiGatewayDiscoveryOperations() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMCISApiGatewayDiscoveryOperationsRead,
		Schema: map[string]*schema.Schema{
			cisID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "CIS instance CRN",
				ValidateFunc: validate.InvokeDataSourceValidator(
					ibmCISApiGatewayDiscoveryOperationsDS,
					"cis_id"),
			},
			cisDomainID: {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "Associated CIS domain ID",
				DiffSuppressFunc: suppressDomainIDDiff,
			},
			// Optional filter arguments
			cisApiGatewayDiscoveryOpsDiff: {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "When true, return only operations not yet saved into API Shield Endpoint Management",
			},
			cisApiGatewayDiscoveryOpsDirection: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Direction to order results (asc or desc)",
			},
			cisApiGatewayDiscoveryOpResultEndpoint: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter results to only include endpoints containing this pattern",
			},
			cisApiGatewayDiscoveryOpResultHost: {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter results to only include the specified hosts",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			cisApiGatewayDiscoveryOpResultMethod: {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter results to only include the specified HTTP methods",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			cisApiGatewayDiscoveryOpsOrder: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Field to order results by (endpoint, host, method, traffic_stats.last_updated, traffic_stats.requests)",
			},
			cisApiGatewayDiscoveryOpsOrigin: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter by discovery engine source (ML, LabelDiscovery, SessionIdentifier)",
			},
			cisApiGatewayDiscoveryOpsState: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter results by discovery state (review, saved, ignored)",
			},
			cisApiGatewayDiscoveryOpsPage: {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Page number of paginated results",
			},
			cisApiGatewayDiscoveryOpsPerPage: {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Maximum number of results per page",
			},
			// Computed attributes
			cisApiGatewayDiscoveryOpsTotalCnt: {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Total number of matching operations",
			},
			cisApiGatewayDiscoveryOpsOps: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of discovered API Gateway operations",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						cisApiGatewayDiscoveryOpResultID: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "UUID of the discovered operation",
						},
						cisApiGatewayDiscoveryOpResultEndpoint: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Endpoint path",
						},
						cisApiGatewayDiscoveryOpResultHost: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Host",
						},
						cisApiGatewayDiscoveryOpResultMethod: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "HTTP method",
						},
						cisApiGatewayDiscoveryOpResultState: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Discovery state (review/saved/ignored)",
						},
						cisApiGatewayDiscoveryOpResultOrigin: {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Discovery engine origins",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						cisApiGatewayDiscoveryOpResultLastUpdated: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp of last update (RFC3339)",
						},
					},
				},
			},
		},
	}
}

// DataSourceIBMCISApiGatewayDiscoveryOperationsValidator returns the ResourceValidator
// for the ibm_cis_api_gateway_discovery_operations data source.
func DataSourceIBMCISApiGatewayDiscoveryOperationsValidator() *validate.ResourceValidator {
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
		ResourceName: ibmCISApiGatewayDiscoveryOperationsDS,
		Schema:       validateSchema,
	}
}

func dataSourceIBMCISApiGatewayDiscoveryOperationsRead(d *schema.ResourceData, meta interface{}) error {
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

	opt := cisClient.NewListApiGatewayDiscoveryOperationsOptions()

	if v, ok := d.GetOkExists(cisApiGatewayDiscoveryOpsDiff); ok { //nolint:staticcheck
		opt.SetDiff(v.(bool))
	}
	if v, ok := d.GetOk(cisApiGatewayDiscoveryOpsDirection); ok {
		opt.SetDirection(v.(string))
	}
	if v, ok := d.GetOk(cisApiGatewayDiscoveryOpResultEndpoint); ok {
		opt.SetEndpoint(v.(string))
	}
	if v, ok := d.GetOk(cisApiGatewayDiscoveryOpResultHost); ok {
		opt.SetHost(flex.ExpandStringList(v.([]interface{})))
	}
	if v, ok := d.GetOk(cisApiGatewayDiscoveryOpResultMethod); ok {
		opt.SetMethod(flex.ExpandStringList(v.([]interface{})))
	}
	if v, ok := d.GetOk(cisApiGatewayDiscoveryOpsOrder); ok {
		opt.SetOrder(v.(string))
	}
	if v, ok := d.GetOk(cisApiGatewayDiscoveryOpsOrigin); ok {
		opt.SetOrigin(v.(string))
	}
	if v, ok := d.GetOk(cisApiGatewayDiscoveryOpsState); ok {
		opt.SetState(v.(string))
	}
	if v, ok := d.GetOk(cisApiGatewayDiscoveryOpsPage); ok {
		opt.SetPage(int64(v.(int)))
	}
	if v, ok := d.GetOk(cisApiGatewayDiscoveryOpsPerPage); ok {
		opt.SetPerPage(int64(v.(int)))
	}

	result, response, err := cisClient.ListApiGatewayDiscoveryOperationsWithContext(context.Background(), opt)
	if err != nil {
		log.Printf("[ERROR] ListApiGatewayDiscoveryOperations failed: %v", response)
		return err
	}

	d.SetId(flex.ConvertCisToTfTwoVar(zoneID, crn))
	d.Set(cisID, crn)
	d.Set(cisDomainID, zoneID)

	if result != nil {
		d.Set(cisApiGatewayDiscoveryOpsOps, flattenDiscoveryOperationsList(result.Result))
		if result.ResultInfo != nil && result.ResultInfo.TotalCount != nil {
			d.Set(cisApiGatewayDiscoveryOpsTotalCnt, int(*result.ResultInfo.TotalCount))
		}
	}
	return nil
}
