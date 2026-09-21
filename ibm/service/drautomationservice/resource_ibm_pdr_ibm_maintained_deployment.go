// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.116.0-df613dbc-20260803-154903
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
	"github.com/IBM/dra-go-sdk/drautomationservicev1"
)

func ResourceIBMPdrIBMMaintainedDeployment() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPdrIBMMaintainedDeploymentCreate,
		ReadContext:   resourceIBMPdrIBMMaintainedDeploymentRead,
		DeleteContext: resourceIBMPdrIBMMaintainedDeploymentDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				// ValidateFunc: validate.InvokeValidator("ibm_pdr_ibm_maintained_deployment", "instance_id"),
				Description: "Service Instance ID.",
			},
			"stand_by_redeploy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				// ValidateFunc: validate.InvokeValidator("ibm_pdr_ibm_maintained_deployment", "stand_by_redeploy"),
				Description: "Flag to indicate if standby should be redeployed (must be \"true\" or \"false\").",
			},
			"accept_language": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				// ValidateFunc: validate.InvokeValidator("ibm_pdr_ibm_maintained_deployment", "accept_language"),
				Description: "The language requested for the return document.",
			},
			"accepts_incomplete": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				ForceNew:    true,
				Description: "A value of true indicates that both the IBM Cloud platform and the requesting client support asynchronous deprovisioning.",
			},
			"managed_apikey": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Sensitive:   true,
				Description: "Apikey used to manage the workloads by adding the PowerVS instances to the orchestrator.",
			},

			"orchestrator_password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Sensitive:   true,
				Description: "The password that you can use to access your orchestrator.",
			},

			"orchestrator_ha": &schema.Schema{
				Type:        schema.TypeBool,
				Required:    true,
				ForceNew:    true,
				Description: "Indicates whether the orchestrator High Availability (HA) is enabled for the service instance.",
			},
			"dashboard_url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL to the dashboard for managing the DR service instance in IBM Cloud.",
			},
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN (Cloud Resource Name) of the DR service instance.",
			},
		},
	}
}

func ResourceIBMPdrIBMMaintainedDeploymentValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "instance_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[a-zA-Z0-9\-_:\/.]+$`,
			MinValueLength:             1,
			MaxValueLength:             512,
		},
		validate.ValidateSchema{
			Identifier:                 "stand_by_redeploy",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^(true|false)$`,
			MinValueLength:             1,
			MaxValueLength:             5,
		},
		validate.ValidateSchema{
			Identifier:                 "accept_language",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9\-_,;=.*]+$`,
			MinValueLength:             1,
			MaxValueLength:             50,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_pdr_ibm_maintained_deployment", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMPdrIBMMaintainedDeploymentCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_ibm_maintained_deployment", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createIBMMaintainedDeploymentOptions := &drautomationservicev1.CreateIBMMaintainedDeploymentOptions{}

	createIBMMaintainedDeploymentOptions.SetInstanceID(d.Get("instance_id").(string))
	createIBMMaintainedDeploymentOptions.SetOrchestratorHa(d.Get("orchestrator_ha").(bool))
	if _, ok := d.GetOk("managed_apikey"); ok {
		createIBMMaintainedDeploymentOptions.SetManagedApikey(d.Get("managed_apikey").(string))
	}
	if _, ok := d.GetOk("orchestrator_password"); ok {
		createIBMMaintainedDeploymentOptions.SetOrchestratorPassword(d.Get("orchestrator_password").(string))
	}
	if _, ok := d.GetOk("stand_by_redeploy"); ok {
		createIBMMaintainedDeploymentOptions.SetStandByRedeploy(d.Get("stand_by_redeploy").(string))
	}
	if _, ok := d.GetOk("accept_language"); ok {
		createIBMMaintainedDeploymentOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}
	if _, ok := d.GetOk("accepts_incomplete"); ok {
		createIBMMaintainedDeploymentOptions.SetAcceptsIncomplete(d.Get("accepts_incomplete").(bool))
	}

	_, response, err := drAutomationServiceClient.CreateIBMMaintainedDeploymentWithContext(context, createIBMMaintainedDeploymentOptions)
	if err != nil {
		detailedMsg := fmt.Sprintf("CreateIBMMaintainedDeploymentWithContext failed: %s", err.Error())
		// Include HTTP status & raw body if available
		if response != nil {
			detailedMsg = fmt.Sprintf(
				"CreateIBMMaintainedDeploymentWithContext failed: %s (status: %d, response: %s)",
				err.Error(), response.StatusCode, response.Result,
			)
		}
		tfErr := flex.TerraformErrorf(err, detailedMsg, "ibm_pdr_ibm_maintained_deployment", "create")
		log.Printf("[ERROR] %s", detailedMsg)
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *createIBMMaintainedDeploymentOptions.InstanceID))

	return resourceIBMPdrIBMMaintainedDeploymentRead(context, d, meta)
}

func resourceIBMPdrIBMMaintainedDeploymentRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_ibm_maintained_deployment", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getIBMMaintainedDeploymentOptions := &drautomationservicev1.GetIBMMaintainedDeploymentOptions{}

	parts, err := flex.SepIdParts(d.Id(), ":")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_ibm_maintained_deployment", "read", "sep-id-parts").GetDiag()
	}

	getIBMMaintainedDeploymentOptions.SetInstanceID(parts[0])
	getIBMMaintainedDeploymentOptions.SetInstanceID(parts[1])
	if _, ok := d.GetOk("accept_language"); ok {
		getIBMMaintainedDeploymentOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}

	ibmOrchestratorGetDeploymentResponse, response, err := drAutomationServiceClient.GetIBMMaintainedDeploymentWithContext(context, getIBMMaintainedDeploymentOptions)
	if err != nil {
		detailedMsg := fmt.Sprintf("GetIBMMaintainedDeploymentWithContext failed: %s", err.Error())
		// Include HTTP status & raw body if available
		if response != nil {
			detailedMsg = fmt.Sprintf(
				"GetIBMMaintainedDeploymentWithContext failed: %s (status: %d, response: %s)",
				err.Error(), response.StatusCode, response.Result,
			)
		}
		tfErr := flex.TerraformErrorf(err, detailedMsg, "ibm_pdr_ibm_maintained_deployment", "read")
		log.Printf("[ERROR] %s", detailedMsg)
		return tfErr.GetDiag()
	}

	// if err = d.Set("dashboard_url", ibmOrchestratorGetDeploymentResponse.OrchestratorDetails); err != nil {
	// 	err = fmt.Errorf("Error setting dashboard_url: %s", err)
	// 	return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_ibm_maintained_deployment", "read", "set-dashboard_url").GetDiag()
	// }
	if err = d.Set("instance_id", ibmOrchestratorGetDeploymentResponse.ServiceDetails.CRN); err != nil {
		err = fmt.Errorf("Error setting instance_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_pdr_ibm_maintained_deployment", "read", "set-instance_id").GetDiag()
	}

	return nil
}

func resourceIBMPdrIBMMaintainedDeploymentDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.
	d.SetId("")
	return nil
}
