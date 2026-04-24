// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.102.0-615ec964-20250307-203034
 */

package codeengine

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/code-engine-go-sdk/codeenginev2"
	"github.com/IBM/go-sdk-core/v5/core"
)

func DataSourceIbmCodeEngineAllowedOutboundDestination() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmCodeEngineAllowedOutboundDestinationRead,

		Schema: map[string]*schema.Schema{
			"project_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the project.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of your allowed outbound destination.",
			},
			"entity_tag": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The version of the allowed outbound destination, which is used to achieve optimistic locking.",
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The current status of the outbound destination.",
			},
			"status_details": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"endpoint_gateway": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Optional information about the endpoint gateway located in the Code Engine VPC that connects to the private path service gateway.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"account_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The account that created the endpoint gateway.",
									},
									"created_at": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The timestamp when the endpoint gateway was created.",
									},
									"ips": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The reserved IPs bound to this endpoint gateway.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name for this endpoint gateway. The name is unique across all endpoint gateways in the VPC.",
									},
								},
							},
						},
						"private_path_service_gateway": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Optional information about the private path service gateway that this allowed outbound destination points to.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The private path service gateway identifier.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of private path service gateway.",
									},
									"service_endpoints": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The fully qualified domain names for this private path service gateway. The domains are used for endpoint gateways to connect to the service and are configured in the VPC for each endpoint gateway.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"reason": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Optional information to provide more context in case of a 'failed' or 'deploying' status.",
						},
					},
				},
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specify the type of the allowed outbound destination. Allowed types are: `cidr_block` and `private_path_service_gateway`.",
			},
			"cidr_block": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IPv4 address range.",
			},
			"private_path_service_gateway_crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN of the Private Path service.",
			},
			"isolation_policy": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Optional property to specify the isolation policy of the private path service gateway. If set to `shared`, other projects within the same account or enterprise account family can connect to Private Path service, too. If set to `dedicated` the gateway can only be used by a single Code Engine project. If not specified the isolation policy will be set to `shared`.",
			},
		},
	}
}

func dataSourceIbmCodeEngineAllowedOutboundDestinationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	codeEngineClient, err := meta.(conns.ClientSession).CodeEngineV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_code_engine_allowed_outbound_destination", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getAllowedOutboundDestinationOptions := &codeenginev2.GetAllowedOutboundDestinationOptions{}

	getAllowedOutboundDestinationOptions.SetProjectID(d.Get("project_id").(string))
	getAllowedOutboundDestinationOptions.SetName(d.Get("name").(string))

	allowedOutboundDestinationIntf, _, err := codeEngineClient.GetAllowedOutboundDestinationWithContext(context, getAllowedOutboundDestinationOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetAllowedOutboundDestinationWithContext failed: %s", err.Error()), "(Data) ibm_code_engine_allowed_outbound_destination", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	allowedOutboundDestination := allowedOutboundDestinationIntf.(*codeenginev2.AllowedOutboundDestination)

	d.SetId(fmt.Sprintf("%s/%s", *getAllowedOutboundDestinationOptions.ProjectID, *getAllowedOutboundDestinationOptions.Name))

	if !core.IsNil(allowedOutboundDestination.EntityTag) {
		if err = d.Set("entity_tag", allowedOutboundDestination.EntityTag); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting entity_tag: %s", err), "(Data) ibm_code_engine_allowed_outbound_destination", "read", "set-entity_tag").GetDiag()
		}
	}

	if !core.IsNil(allowedOutboundDestination.Status) {
		if err = d.Set("status", allowedOutboundDestination.Status); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_code_engine_allowed_outbound_destination", "read", "set-status").GetDiag()
		}
	}

	if !core.IsNil(allowedOutboundDestination.StatusDetails) {
		statusDetails := []map[string]interface{}{}
		statusDetailsMap, err := DataSourceIbmCodeEngineAllowedOutboundDestinationAllowedOutboundStatusDetailsToMap(allowedOutboundDestination.StatusDetails)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_code_engine_allowed_outbound_destination", "read", "status_details-to-map").GetDiag()
		}
		statusDetails = append(statusDetails, statusDetailsMap)
		if err = d.Set("status_details", statusDetails); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status_details: %s", err), "(Data) ibm_code_engine_allowed_outbound_destination", "read", "set-status_details").GetDiag()
		}
	} else {
		// Explicitly set to empty list when StatusDetails is nil
		if err = d.Set("status_details", []map[string]interface{}{}); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status_details: %s", err), "(Data) ibm_code_engine_allowed_outbound_destination", "read", "set-status_details").GetDiag()
		}
	}

	if err = d.Set("type", allowedOutboundDestination.Type); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting type: %s", err), "(Data) ibm_code_engine_allowed_outbound_destination", "read", "set-type").GetDiag()
	}

	if !core.IsNil(allowedOutboundDestination.CidrBlock) {
		if err = d.Set("cidr_block", allowedOutboundDestination.CidrBlock); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting cidr_block: %s", err), "(Data) ibm_code_engine_allowed_outbound_destination", "read", "set-cidr_block").GetDiag()
		}
	}

	if !core.IsNil(allowedOutboundDestination.PrivatePathServiceGatewayCrn) {
		if err = d.Set("private_path_service_gateway_crn", allowedOutboundDestination.PrivatePathServiceGatewayCrn); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting private_path_service_gateway_crn: %s", err), "(Data) ibm_code_engine_allowed_outbound_destination", "read", "set-private_path_service_gateway_crn").GetDiag()
		}
	}

	if !core.IsNil(allowedOutboundDestination.IsolationPolicy) {
		if err = d.Set("isolation_policy", allowedOutboundDestination.IsolationPolicy); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting isolation_policy: %s", err), "(Data) ibm_code_engine_allowed_outbound_destination", "read", "set-isolation_policy").GetDiag()
		}
	}

	return nil
}

func DataSourceIbmCodeEngineAllowedOutboundDestinationAllowedOutboundStatusDetailsToMap(model codeenginev2.AllowedOutboundStatusDetailsIntf) (map[string]interface{}, error) {
	if _, ok := model.(*codeenginev2.AllowedOutboundStatusDetailsPrivatePathServiceGatewayStatusDetails); ok {
		return DataSourceIbmCodeEngineAllowedOutboundDestinationAllowedOutboundStatusDetailsPrivatePathServiceGatewayStatusDetailsToMap(model.(*codeenginev2.AllowedOutboundStatusDetailsPrivatePathServiceGatewayStatusDetails))
	} else if _, ok := model.(*codeenginev2.AllowedOutboundStatusDetails); ok {
		modelMap := make(map[string]interface{})
		model := model.(*codeenginev2.AllowedOutboundStatusDetails)
		if model.EndpointGateway != nil {
			endpointGatewayMap, err := DataSourceIbmCodeEngineAllowedOutboundDestinationEndpointGatewayDetailsToMap(model.EndpointGateway)
			if err != nil {
				return modelMap, err
			}
			modelMap["endpoint_gateway"] = []map[string]interface{}{endpointGatewayMap}
		}
		if model.PrivatePathServiceGateway != nil {
			privatePathServiceGatewayMap, err := DataSourceIbmCodeEngineAllowedOutboundDestinationPrivatePathServiceGatewayDetailsToMap(model.PrivatePathServiceGateway)
			if err != nil {
				return modelMap, err
			}
			modelMap["private_path_service_gateway"] = []map[string]interface{}{privatePathServiceGatewayMap}
		}
		if model.Reason != nil {
			modelMap["reason"] = *model.Reason
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized codeenginev2.AllowedOutboundStatusDetailsIntf subtype encountered")
	}
}

func DataSourceIbmCodeEngineAllowedOutboundDestinationEndpointGatewayDetailsToMap(model *codeenginev2.EndpointGatewayDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.AccountID != nil {
		modelMap["account_id"] = *model.AccountID
	}
	if model.CreatedAt != nil {
		modelMap["created_at"] = *model.CreatedAt
	}
	if model.Ips != nil {
		modelMap["ips"] = model.Ips
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	return modelMap, nil
}

func DataSourceIbmCodeEngineAllowedOutboundDestinationPrivatePathServiceGatewayDetailsToMap(model *codeenginev2.PrivatePathServiceGatewayDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.ServiceEndpoints != nil {
		modelMap["service_endpoints"] = model.ServiceEndpoints
	}
	return modelMap, nil
}

func DataSourceIbmCodeEngineAllowedOutboundDestinationAllowedOutboundStatusDetailsPrivatePathServiceGatewayStatusDetailsToMap(model *codeenginev2.AllowedOutboundStatusDetailsPrivatePathServiceGatewayStatusDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.EndpointGateway != nil {
		endpointGatewayMap, err := DataSourceIbmCodeEngineAllowedOutboundDestinationEndpointGatewayDetailsToMap(model.EndpointGateway)
		if err != nil {
			return modelMap, err
		}
		modelMap["endpoint_gateway"] = []map[string]interface{}{endpointGatewayMap}
	}
	if model.PrivatePathServiceGateway != nil {
		privatePathServiceGatewayMap, err := DataSourceIbmCodeEngineAllowedOutboundDestinationPrivatePathServiceGatewayDetailsToMap(model.PrivatePathServiceGateway)
		if err != nil {
			return modelMap, err
		}
		modelMap["private_path_service_gateway"] = []map[string]interface{}{privatePathServiceGatewayMap}
	}
	if model.Reason != nil {
		modelMap["reason"] = *model.Reason
	}
	return modelMap, nil
}
