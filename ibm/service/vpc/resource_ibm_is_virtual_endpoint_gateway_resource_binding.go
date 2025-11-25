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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func ResourceIBMIsVirtualEndpointGatewayResourceBinding() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIsVirtualEndpointGatewayResourceBindingCreate,
		ReadContext:   resourceIBMIsVirtualEndpointGatewayResourceBindingRead,
		UpdateContext: resourceIBMIsVirtualEndpointGatewayResourceBindingUpdate,
		DeleteContext: resourceIBMIsVirtualEndpointGatewayResourceBindingDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"endpoint_gateway_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_virtual_endpoint_gateway_resource_binding", "endpoint_gateway_id"),
				Description:  "The endpoint gateway identifier.",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_virtual_endpoint_gateway_resource_binding", "name"),
				Description:  "The name for this resource binding. The name is unique across all resource bindings for the endpoint gateway.",
			},
			"target": &schema.Schema{
				Type:        schema.TypeList,
				MinItems:    1,
				MaxItems:    1,
				Required:    true,
				Description: "The target for this endpoint gateway resource binding.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
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
							Optional:    true,
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
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of resource binding:- `weak`: The binding is not dependent on the existence of the target resource.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
			},
			"endpoint_gateway_resource_binding_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier for this endpoint gateway resource binding.",
			},
		},
	}
}

func ResourceIBMIsVirtualEndpointGatewayResourceBindingValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "endpoint_gateway_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[-0-9a-z_]+$`,
			MinValueLength:             1,
			MaxValueLength:             64,
		},
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$`,
			MinValueLength:             1,
			MaxValueLength:             63,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_virtual_endpoint_gateway_resource_binding", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMIsVirtualEndpointGatewayResourceBindingCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_endpoint_gateway_resource_binding", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createEndpointGatewayResourceBindingOptions := &vpcv1.CreateEndpointGatewayResourceBindingOptions{}

	createEndpointGatewayResourceBindingOptions.SetEndpointGatewayID(d.Get("endpoint_gateway_id").(string))
	targetModel, err := ResourceIBMIsVirtualEndpointGatewayResourceBindingMapToEndpointGatewayResourceBindingTargetPrototype(d.Get("target.0").(map[string]interface{}))
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_endpoint_gateway_resource_binding", "create", "parse-target").GetDiag()
	}
	createEndpointGatewayResourceBindingOptions.SetTarget(targetModel)
	if _, ok := d.GetOk("name"); ok {
		createEndpointGatewayResourceBindingOptions.SetName(d.Get("name").(string))
	}

	endpointGatewayResourceBinding, _, err := vpcClient.CreateEndpointGatewayResourceBindingWithContext(context, createEndpointGatewayResourceBindingOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateEndpointGatewayResourceBindingWithContext failed: %s", err.Error()), "ibm_is_virtual_endpoint_gateway_resource_binding", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *createEndpointGatewayResourceBindingOptions.EndpointGatewayID, *endpointGatewayResourceBinding.ID))

	return resourceIBMIsVirtualEndpointGatewayResourceBindingRead(context, d, meta)
}

func resourceIBMIsVirtualEndpointGatewayResourceBindingRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_endpoint_gateway_resource_binding", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getEndpointGatewayResourceBindingOptions := &vpcv1.GetEndpointGatewayResourceBindingOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_endpoint_gateway_resource_binding", "read", "sep-id-parts").GetDiag()
	}

	getEndpointGatewayResourceBindingOptions.SetEndpointGatewayID(parts[0])
	getEndpointGatewayResourceBindingOptions.SetID(parts[1])

	endpointGatewayResourceBinding, response, err := vpcClient.GetEndpointGatewayResourceBindingWithContext(context, getEndpointGatewayResourceBindingOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetEndpointGatewayResourceBindingWithContext failed: %s", err.Error()), "ibm_is_virtual_endpoint_gateway_resource_binding", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if !core.IsNil(endpointGatewayResourceBinding.Name) {
		if err = d.Set("name", endpointGatewayResourceBinding.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_endpoint_gateway_resource_binding", "read", "set-name").GetDiag()
		}
	}
	targetMap, err := ResourceIBMIsVirtualEndpointGatewayResourceBindingEndpointGatewayResourceBindingTargetToMap(endpointGatewayResourceBinding.Target)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_endpoint_gateway_resource_binding", "read", "target-to-map").GetDiag()
	}
	if err = d.Set("target", []map[string]interface{}{targetMap}); err != nil {
		err = fmt.Errorf("Error setting target: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_endpoint_gateway_resource_binding", "read", "set-target").GetDiag()
	}
	if err = d.Set("created_at", flex.DateTimeToString(endpointGatewayResourceBinding.CreatedAt)); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_endpoint_gateway_resource_binding", "read", "set-created_at").GetDiag()
	}
	if err = d.Set("href", endpointGatewayResourceBinding.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_endpoint_gateway_resource_binding", "read", "set-href").GetDiag()
	}
	lifecycleReasons := []map[string]interface{}{}
	for _, lifecycleReasonsItem := range endpointGatewayResourceBinding.LifecycleReasons {
		lifecycleReasonsItemMap, err := ResourceIBMIsVirtualEndpointGatewayResourceBindingEndpointGatewayResourceBindingLifecycleReasonToMap(&lifecycleReasonsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_endpoint_gateway_resource_binding", "read", "lifecycle_reasons-to-map").GetDiag()
		}
		lifecycleReasons = append(lifecycleReasons, lifecycleReasonsItemMap)
	}
	if err = d.Set("lifecycle_reasons", lifecycleReasons); err != nil {
		err = fmt.Errorf("Error setting lifecycle_reasons: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_endpoint_gateway_resource_binding", "read", "set-lifecycle_reasons").GetDiag()
	}
	if err = d.Set("lifecycle_state", endpointGatewayResourceBinding.LifecycleState); err != nil {
		err = fmt.Errorf("Error setting lifecycle_state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_endpoint_gateway_resource_binding", "read", "set-lifecycle_state").GetDiag()
	}
	if err = d.Set("resource_type", endpointGatewayResourceBinding.ResourceType); err != nil {
		err = fmt.Errorf("Error setting resource_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_endpoint_gateway_resource_binding", "read", "set-resource_type").GetDiag()
	}
	if err = d.Set("service_endpoint", endpointGatewayResourceBinding.ServiceEndpoint); err != nil {
		err = fmt.Errorf("Error setting service_endpoint: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_endpoint_gateway_resource_binding", "read", "set-service_endpoint").GetDiag()
	}
	if err = d.Set("type", endpointGatewayResourceBinding.Type); err != nil {
		err = fmt.Errorf("Error setting type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_endpoint_gateway_resource_binding", "read", "set-type").GetDiag()
	}
	if err = d.Set("endpoint_gateway_resource_binding_id", endpointGatewayResourceBinding.ID); err != nil {
		err = fmt.Errorf("Error setting endpoint_gateway_resource_binding_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_endpoint_gateway_resource_binding", "read", "set-endpoint_gateway_resource_binding_id").GetDiag()
	}

	return nil
}

func resourceIBMIsVirtualEndpointGatewayResourceBindingUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_endpoint_gateway_resource_binding", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateEndpointGatewayResourceBindingOptions := &vpcv1.UpdateEndpointGatewayResourceBindingOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_endpoint_gateway_resource_binding", "update", "sep-id-parts").GetDiag()
	}

	updateEndpointGatewayResourceBindingOptions.SetEndpointGatewayID(parts[0])
	updateEndpointGatewayResourceBindingOptions.SetID(parts[1])

	hasChange := false

	patchVals := &vpcv1.EndpointGatewayResourceBindingPatch{}
	if d.HasChange("endpoint_gateway_id") {
		errMsg := fmt.Sprintf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "endpoint_gateway_id")
		return flex.DiscriminatedTerraformErrorf(nil, errMsg, "ibm_is_virtual_endpoint_gateway_resource_binding", "update", "endpoint_gateway_id-forces-new").GetDiag()
	}
	if d.HasChange("name") {
		newName := d.Get("name").(string)
		patchVals.Name = &newName
		hasChange = true
	}

	if hasChange {
		// Fields with `nil` values are omitted from the generic map,
		// so we need to re-add them to support removing arguments
		// in merge-patch operations sent to the service.
		updateEndpointGatewayResourceBindingOptions.EndpointGatewayResourceBindingPatch = ResourceIBMIsVirtualEndpointGatewayResourceBindingEndpointGatewayResourceBindingPatchAsPatch(patchVals, d)

		_, _, err = vpcClient.UpdateEndpointGatewayResourceBindingWithContext(context, updateEndpointGatewayResourceBindingOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateEndpointGatewayResourceBindingWithContext failed: %s", err.Error()), "ibm_is_virtual_endpoint_gateway_resource_binding", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIBMIsVirtualEndpointGatewayResourceBindingRead(context, d, meta)
}

func resourceIBMIsVirtualEndpointGatewayResourceBindingDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_endpoint_gateway_resource_binding", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteEndpointGatewayResourceBindingOptions := &vpcv1.DeleteEndpointGatewayResourceBindingOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_virtual_endpoint_gateway_resource_binding", "delete", "sep-id-parts").GetDiag()
	}

	deleteEndpointGatewayResourceBindingOptions.SetEndpointGatewayID(parts[0])
	deleteEndpointGatewayResourceBindingOptions.SetID(parts[1])

	_, err = vpcClient.DeleteEndpointGatewayResourceBindingWithContext(context, deleteEndpointGatewayResourceBindingOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteEndpointGatewayResourceBindingWithContext failed: %s", err.Error()), "ibm_is_virtual_endpoint_gateway_resource_binding", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIBMIsVirtualEndpointGatewayResourceBindingMapToEndpointGatewayResourceBindingTargetPrototype(modelMap map[string]interface{}) (vpcv1.EndpointGatewayResourceBindingTargetPrototypeIntf, error) {
	model := &vpcv1.EndpointGatewayResourceBindingTargetPrototype{}
	if modelMap["crn"] != nil && modelMap["crn"].(string) != "" {
		model.CRN = core.StringPtr(modelMap["crn"].(string))
	}
	return model, nil
}

func ResourceIBMIsVirtualEndpointGatewayResourceBindingMapToEndpointGatewayResourceBindingTargetPrototypeEndpointGatewayResourceBindingTargetByCRN(modelMap map[string]interface{}) (*vpcv1.EndpointGatewayResourceBindingTargetPrototypeEndpointGatewayResourceBindingTargetByCRN, error) {
	model := &vpcv1.EndpointGatewayResourceBindingTargetPrototypeEndpointGatewayResourceBindingTargetByCRN{}
	model.CRN = core.StringPtr(modelMap["crn"].(string))
	return model, nil
}

func ResourceIBMIsVirtualEndpointGatewayResourceBindingEndpointGatewayResourceBindingTargetToMap(model vpcv1.EndpointGatewayResourceBindingTargetIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.EndpointGatewayResourceBindingTargetCRN); ok {
		return ResourceIBMIsVirtualEndpointGatewayResourceBindingEndpointGatewayResourceBindingTargetCRNToMap(model.(*vpcv1.EndpointGatewayResourceBindingTargetCRN))
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

func ResourceIBMIsVirtualEndpointGatewayResourceBindingEndpointGatewayResourceBindingTargetCRNToMap(model *vpcv1.EndpointGatewayResourceBindingTargetCRN) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	return modelMap, nil
}

func ResourceIBMIsVirtualEndpointGatewayResourceBindingEndpointGatewayResourceBindingLifecycleReasonToMap(model *vpcv1.EndpointGatewayResourceBindingLifecycleReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["code"] = *model.Code
	modelMap["message"] = *model.Message
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func ResourceIBMIsVirtualEndpointGatewayResourceBindingEndpointGatewayResourceBindingPatchAsPatch(patchVals *vpcv1.EndpointGatewayResourceBindingPatch, d *schema.ResourceData) map[string]interface{} {
	patch, _ := patchVals.AsPatch()
	var path string

	path = "name"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["name"] = nil
	} else if !exists {
		delete(patch, "name")
	}

	return patch
}
