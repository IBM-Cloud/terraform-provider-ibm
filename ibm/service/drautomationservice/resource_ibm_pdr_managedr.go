// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.105.0-3c13b041-20250605-193116
 */

package drautomationservice

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
	"github.ibm.com/DRAutomation/dra-go-sdk/drautomationservicev1"
)

func ResourceIbmPdrManagedr() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmPdrManagedrCreate,
		ReadContext:   resourceIbmPdrManagedrRead,
		DeleteContext: resourceIbmPdrManagedrDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "instance id of instance to provision.",
			},
			"stand_by_redeploy": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_pdr_managedr", "stand_by_redeploy"),
				Description:  "Flag to indicate if standby should be redeployed (must be \"true\" or \"false\").",
			},
			"accept_language": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The language requested for the return document.",
			},
			"if_none_match": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "ETag for conditional requests (optional).",
			},
			"accepts_incomplete": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				ForceNew:    true,
				Description: "A value of true indicates that both the IBM Cloud platform and the requesting client support asynchronous deprovisioning.",
			},
			"dashboard_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			// "instance_id": &schema.Schema{
			// 	Type:        schema.TypeString,
			// 	Computed:    true,
			// },
			"etag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func ResourceIbmPdrManagedrValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "stand_by_redeploy",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "false, true",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_pdr_managedr", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmPdrManagedrCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	serviceInstanceManageDrOptions := &drautomationservicev1.ServiceInstanceManageDrOptions{}

	serviceInstanceManageDrOptions.SetInstanceID(d.Get("instance_id").(string))
	serviceInstanceManageDrOptions.SetStandByRedeploy(d.Get("stand_by_redeploy").(string))
	contextModel, err := ResourceIbmPdrManagedrMapToContext(d.Get("context.0").(map[string]interface{}))
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "create", "parse-context").GetDiag()
	}
	serviceInstanceManageDrOptions.SetContext(contextModel)
	serviceInstanceManageDrOptions.SetPlanID(d.Get("plan_id").(string))
	serviceInstanceManageDrOptions.SetServiceID(d.Get("service_id").(string))
	if _, ok := d.GetOk("action"); ok {
		serviceInstanceManageDrOptions.SetAction(d.Get("action").(string))
	}
	if _, ok := d.GetOk("parameters"); ok {
		parametersModel, err := ResourceIbmPdrManagedrMapToManageDrParameters(d.Get("parameters.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "create", "parse-parameters").GetDiag()
		}
		serviceInstanceManageDrOptions.SetParameters(parametersModel)
	}
	if _, ok := d.GetOk("accept_language"); ok {
		serviceInstanceManageDrOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}
	if _, ok := d.GetOk("if_none_match"); ok {
		serviceInstanceManageDrOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}
	if _, ok := d.GetOk("accepts_incomplete"); ok {
		serviceInstanceManageDrOptions.SetAcceptsIncomplete(d.Get("accepts_incomplete").(bool))
	}

	serviceInstanceManageDr, _, err := drAutomationServiceClient.ServiceInstanceManageDrWithContext(context, serviceInstanceManageDrOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ServiceInstanceManageDrWithContext failed: %s", err.Error()), "ibm_pdr_managedr", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *serviceInstanceManageDrOptions.InstanceID, *serviceInstanceManageDr.ID))

	return resourceIbmPdrManagedrRead(context, d, meta)
}

func resourceIbmPdrManagedrRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	serviceInstanceFetchManageDrOptions := &drautomationservicev1.ServiceInstanceFetchManageDrOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "read", "sep-id-parts").GetDiag()
	}

	serviceInstanceFetchManageDrOptions.SetInstanceID(parts[0])
	serviceInstanceFetchManageDrOptions.SetInstanceID(parts[1])
	if _, ok := d.GetOk("accept_language"); ok {
		serviceInstanceFetchManageDrOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}
	if _, ok := d.GetOk("if_none_match"); ok {
		serviceInstanceFetchManageDrOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}

	serviceInstanceManageDr, response, err := drAutomationServiceClient.ServiceInstanceFetchManageDrWithContext(context, serviceInstanceFetchManageDrOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ServiceInstanceFetchManageDrWithContext failed: %s", err.Error()), "ibm_pdr_managedr", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if !core.IsNil(serviceInstanceManageDr.DashboardURL) {
		if err = d.Set("dashboard_url", serviceInstanceManageDr.DashboardURL); err != nil {
			err = fmt.Errorf("Error setting dashboard_url: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "read", "set-dashboard_url").GetDiag()
		}
	}
	if !core.IsNil(serviceInstanceManageDr.ID) {
		if err = d.Set("instance_id", serviceInstanceManageDr.ID); err != nil {
			err = fmt.Errorf("Error setting instance_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_managedr", "read", "set-instance_id").GetDiag()
		}
	}
	if err = d.Set("etag", response.Headers.Get("Etag")); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting etag: %s", err), "ibm_pdr_managedr", "read", "set-etag").GetDiag()
	}

	return nil
}

func resourceIbmPdrManagedrDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.
	d.SetId("")
	return nil
}

func ResourceIbmPdrManagedrMapToContext(modelMap map[string]interface{}) (*drautomationservicev1.Context, error) {
	model := &drautomationservicev1.Context{}
	return model, nil
}

func ResourceIbmPdrManagedrMapToManageDrParameters(modelMap map[string]interface{}) (*drautomationservicev1.ManageDrParameters, error) {
	model := &drautomationservicev1.ManageDrParameters{}
	if modelMap["location"] != nil && modelMap["location"].(string) != "" {
		model.Location = core.StringPtr(modelMap["location"].(string))
	}
	if modelMap["optional_param"] != nil && modelMap["optional_param"].(string) != "" {
		model.OptionalParam = core.StringPtr(modelMap["optional_param"].(string))
	}
	return model, nil
}
