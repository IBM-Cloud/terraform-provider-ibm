// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func DataSourceIBMIsVirtualEndpointGatewayResourceBindings() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsVirtualEndpointGatewayResourceBindingsRead,

		Schema: map[string]*schema.Schema{
			"endpoint_gateway_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The endpoint gateway identifier.",
			},
			"resource_bindings": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A page of resource bindings for the endpoint gateway.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"created_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time that the resource binding was created.",
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL for this endpoint gateway resource binding.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this endpoint gateway resource binding.",
						},
						"lifecycle_reasons": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The reasons for the current `lifecycle_state` (if any).",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"code": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "A reason code for this lifecycle state:- `internal_error`: internal error (contact IBM support)- `resource_suspended_by_provider`: The resource has been suspended (contact IBM  support)The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
									},
									"message": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "An explanation of the reason for this lifecycle state.",
									},
									"more_info": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "A link to documentation about the reason for this lifecycle state.",
									},
								},
							},
						},
						"lifecycle_state": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The lifecycle state of the resource binding.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name for this resource binding. The name is unique across all resource bindings for the endpoint gateway.",
						},
						"resource_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
						"service_endpoint": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The fully qualified domain name of the service endpoint for the resource targeted by this resource binding.",
						},
						"target": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The target for this endpoint gateway resource binding.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"crn": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of resource binding:- `weak`: The binding is not dependent on the existence of the target resource.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMIsVirtualEndpointGatewayResourceBindingsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_virtual_endpoint_gateway_resource_bindings", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listEndpointGatewayResourceBindingsOptions := &vpcv1.ListEndpointGatewayResourceBindingsOptions{}

	listEndpointGatewayResourceBindingsOptions.SetEndpointGatewayID(d.Get("endpoint_gateway_id").(string))

	var pager *vpcv1.EndpointGatewayResourceBindingsPager
	pager, err = vpcClient.NewEndpointGatewayResourceBindingsPager(listEndpointGatewayResourceBindingsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_is_virtual_endpoint_gateway_resource_bindings", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	allItems, err := pager.GetAll()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("EndpointGatewayResourceBindingsPager.GetAll() failed %s", err), "(Data) ibm_is_virtual_endpoint_gateway_resource_bindings", "read")
		log.Printf("[DEBUG] %s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIBMIsVirtualEndpointGatewayResourceBindingsID(d))

	mapSlice := []map[string]interface{}{}
	for _, modelItem := range allItems {
		modelMap, err := DataSourceIBMIsVirtualEndpointGatewayResourceBindingsEndpointGatewayResourceBindingToMap(&modelItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_virtual_endpoint_gateway_resource_bindings", "read", "EndpointGateways-to-map").GetDiag()
		}
		mapSlice = append(mapSlice, modelMap)
	}

	if err = d.Set("resource_bindings", mapSlice); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_bindings %s", err), "(Data) ibm_is_virtual_endpoint_gateway_resource_bindings", "read", "resource_bindings-set").GetDiag()
	}

	return nil
}

// dataSourceIBMIsVirtualEndpointGatewayResourceBindingsID returns a reasonable ID for the list.
func dataSourceIBMIsVirtualEndpointGatewayResourceBindingsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIBMIsVirtualEndpointGatewayResourceBindingsEndpointGatewayResourceBindingToMap(model *vpcv1.EndpointGatewayResourceBinding) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["created_at"] = model.CreatedAt.String()
	modelMap["href"] = *model.Href
	modelMap["id"] = *model.ID
	lifecycleReasons := []map[string]interface{}{}
	for _, lifecycleReasonsItem := range model.LifecycleReasons {
		lifecycleReasonsItemMap, err := DataSourceIBMIsVirtualEndpointGatewayResourceBindingsEndpointGatewayResourceBindingLifecycleReasonToMap(&lifecycleReasonsItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		lifecycleReasons = append(lifecycleReasons, lifecycleReasonsItemMap)
	}
	modelMap["lifecycle_reasons"] = lifecycleReasons
	modelMap["lifecycle_state"] = *model.LifecycleState
	modelMap["name"] = *model.Name
	modelMap["resource_type"] = *model.ResourceType
	modelMap["service_endpoint"] = *model.ServiceEndpoint
	targetMap, err := DataSourceIBMIsVirtualEndpointGatewayResourceBindingsEndpointGatewayResourceBindingTargetToMap(model.Target)
	if err != nil {
		return modelMap, err
	}
	modelMap["target"] = []map[string]interface{}{targetMap}
	modelMap["type"] = *model.Type
	return modelMap, nil
}

func DataSourceIBMIsVirtualEndpointGatewayResourceBindingsEndpointGatewayResourceBindingLifecycleReasonToMap(model *vpcv1.EndpointGatewayResourceBindingLifecycleReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["code"] = *model.Code
	modelMap["message"] = *model.Message
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func DataSourceIBMIsVirtualEndpointGatewayResourceBindingsEndpointGatewayResourceBindingTargetToMap(model vpcv1.EndpointGatewayResourceBindingTargetIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.EndpointGatewayResourceBindingTargetCRN); ok {
		return DataSourceIBMIsVirtualEndpointGatewayResourceBindingsEndpointGatewayResourceBindingTargetCRNToMap(model.(*vpcv1.EndpointGatewayResourceBindingTargetCRN))
	} else if _, ok := model.(*vpcv1.EndpointGatewayResourceBindingTarget); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.EndpointGatewayResourceBindingTarget)
		if model.CRN != nil {
			modelMap["crn"] = *model.CRN
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.EndpointGatewayResourceBindingTargetIntf subtype encountered")
	}
}

func DataSourceIBMIsVirtualEndpointGatewayResourceBindingsEndpointGatewayResourceBindingTargetCRNToMap(model *vpcv1.EndpointGatewayResourceBindingTargetCRN) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	return modelMap, nil
}
