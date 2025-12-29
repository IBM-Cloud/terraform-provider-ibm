// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func DataSourceIBMIsVirtualEndpointGatewayResourceBinding() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsVirtualEndpointGatewayResourceBindingRead,

		Schema: map[string]*schema.Schema{
			"endpoint_gateway_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The endpoint gateway identifier.",
			},
			"endpoint_gateway_resource_binding_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The resource binding identifier.",
			},
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
	}
}

func dataSourceIBMIsVirtualEndpointGatewayResourceBindingRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_virtual_endpoint_gateway_resource_binding", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getEndpointGatewayResourceBindingOptions := &vpcv1.GetEndpointGatewayResourceBindingOptions{}

	getEndpointGatewayResourceBindingOptions.SetEndpointGatewayID(d.Get("endpoint_gateway_id").(string))
	getEndpointGatewayResourceBindingOptions.SetID(d.Get("endpoint_gateway_resource_binding_id").(string))

	endpointGatewayResourceBinding, _, err := vpcClient.GetEndpointGatewayResourceBindingWithContext(context, getEndpointGatewayResourceBindingOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetEndpointGatewayResourceBindingWithContext failed: %s", err.Error()), "(Data) ibm_is_virtual_endpoint_gateway_resource_binding", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *getEndpointGatewayResourceBindingOptions.EndpointGatewayID, *getEndpointGatewayResourceBindingOptions.ID))

	if err = d.Set("created_at", flex.DateTimeToString(endpointGatewayResourceBinding.CreatedAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_is_virtual_endpoint_gateway_resource_binding", "read", "set-created_at").GetDiag()
	}

	if err = d.Set("href", endpointGatewayResourceBinding.Href); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_virtual_endpoint_gateway_resource_binding", "read", "set-href").GetDiag()
	}

	lifecycleReasons := []map[string]interface{}{}
	for _, lifecycleReasonsItem := range endpointGatewayResourceBinding.LifecycleReasons {
		lifecycleReasonsItemMap, err := DataSourceIBMIsVirtualEndpointGatewayResourceBindingEndpointGatewayResourceBindingLifecycleReasonToMap(&lifecycleReasonsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_virtual_endpoint_gateway_resource_binding", "read", "lifecycle_reasons-to-map").GetDiag()
		}
		lifecycleReasons = append(lifecycleReasons, lifecycleReasonsItemMap)
	}
	if err = d.Set("lifecycle_reasons", lifecycleReasons); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting lifecycle_reasons: %s", err), "(Data) ibm_is_virtual_endpoint_gateway_resource_binding", "read", "set-lifecycle_reasons").GetDiag()
	}

	if err = d.Set("lifecycle_state", endpointGatewayResourceBinding.LifecycleState); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting lifecycle_state: %s", err), "(Data) ibm_is_virtual_endpoint_gateway_resource_binding", "read", "set-lifecycle_state").GetDiag()
	}

	if err = d.Set("name", endpointGatewayResourceBinding.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_virtual_endpoint_gateway_resource_binding", "read", "set-name").GetDiag()
	}

	if err = d.Set("resource_type", endpointGatewayResourceBinding.ResourceType); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_type: %s", err), "(Data) ibm_is_virtual_endpoint_gateway_resource_binding", "read", "set-resource_type").GetDiag()
	}

	if err = d.Set("service_endpoint", endpointGatewayResourceBinding.ServiceEndpoint); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting service_endpoint: %s", err), "(Data) ibm_is_virtual_endpoint_gateway_resource_binding", "read", "set-service_endpoint").GetDiag()
	}

	target := []map[string]interface{}{}
	targetMap, err := DataSourceIBMIsVirtualEndpointGatewayResourceBindingEndpointGatewayResourceBindingTargetToMap(endpointGatewayResourceBinding.Target)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_virtual_endpoint_gateway_resource_binding", "read", "target-to-map").GetDiag()
	}
	target = append(target, targetMap)
	if err = d.Set("target", target); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting target: %s", err), "(Data) ibm_is_virtual_endpoint_gateway_resource_binding", "read", "set-target").GetDiag()
	}

	if err = d.Set("type", endpointGatewayResourceBinding.Type); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting type: %s", err), "(Data) ibm_is_virtual_endpoint_gateway_resource_binding", "read", "set-type").GetDiag()
	}

	return nil
}

func DataSourceIBMIsVirtualEndpointGatewayResourceBindingEndpointGatewayResourceBindingLifecycleReasonToMap(model *vpcv1.EndpointGatewayResourceBindingLifecycleReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["code"] = *model.Code
	modelMap["message"] = *model.Message
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func DataSourceIBMIsVirtualEndpointGatewayResourceBindingEndpointGatewayResourceBindingTargetToMap(model vpcv1.EndpointGatewayResourceBindingTargetIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.EndpointGatewayResourceBindingTargetCRN); ok {
		return DataSourceIBMIsVirtualEndpointGatewayResourceBindingEndpointGatewayResourceBindingTargetCRNToMap(model.(*vpcv1.EndpointGatewayResourceBindingTargetCRN))
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

func DataSourceIBMIsVirtualEndpointGatewayResourceBindingEndpointGatewayResourceBindingTargetCRNToMap(model *vpcv1.EndpointGatewayResourceBindingTargetCRN) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	return modelMap, nil
}
