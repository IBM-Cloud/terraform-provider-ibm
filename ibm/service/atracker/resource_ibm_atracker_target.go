// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.101.0-62624c1e-20250225-192301
 */

package atracker

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/atrackerv2"
)

const COS_CRN_PARTS = 8

func ResourceIBMAtrackerTarget() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMAtrackerTargetCreate,
		ReadContext:   resourceIBMAtrackerTargetRead,
		UpdateContext: resourceIBMAtrackerTargetUpdate,
		DeleteContext: resourceIBMAtrackerTargetDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_atracker_target", "name"),
				Description:  "The name of the target. The name must be 1000 characters or less, and cannot include any special characters other than `(space) - . _ :`.",
			},
			"target_type": &schema.Schema{
				Type:             schema.TypeString,
				DiffSuppressFunc: flex.ApplyOnce,
				Required:         true,
				ForceNew:         true,
				ValidateFunc:     validate.InvokeValidator("ibm_atracker_target", "target_type"),
				Description:      "The type of the target. It can be cloud_object_storage, event_streams, or cloud_logs. Based on this type you must include cos_endpoint, eventstreams_endpoint or cloudlogs_endpoint.",
			},
			"cos_endpoint": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Property values for a Cloud Object Storage endpoint.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"endpoint": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The host name of the Cloud Object Storage endpoint.",
						},
						"target_crn": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The CRN of the Cloud Object Storage instance.",
						},
						"bucket": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The bucket name under the Cloud Object Storage instance.",
						},
						"api_key": &schema.Schema{
							Type:             schema.TypeString,
							Optional:         true,
							Sensitive:        true,
							Description:      "The IAM API key that has writer access to the Cloud Object Storage instance. This credential is masked in the response. This is required if service_to_service is not enabled.",
							DiffSuppressFunc: flex.ApplyOnce,
						},
						"service_to_service_enabled": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Determines if IBM Cloud Activity Tracker Event Routing has service to service authentication enabled. Set this flag to true if service to service is enabled and do not supply an apikey.",
						},
					},
				},
			},
			"eventstreams_endpoint": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Property values for the Event Streams endpoint.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"target_crn": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The CRN of the Event Streams instance.",
						},
						"brokers": &schema.Schema{
							Type:        schema.TypeList,
							Required:    true,
							Description: "List of broker endpoints.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"topic": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The messsage hub topic defined in the Event Streams instance.",
						},
						"api_key": &schema.Schema{ // pragma: allowlist secret
							Type:             schema.TypeString,
							Optional:         true,
							Sensitive:        true,
							DiffSuppressFunc: flex.ApplyOnce,
							Description:      "The user password (api key) for the message hub topic in the Event Streams instance. This is required if service_to_service is not enabled.",
						},
						"service_to_service_enabled": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Determines if IBM Cloud Activity Tracker Event Routing has service to service authentication enabled. Set this flag to true if service to service is enabled and do not supply an apikey.",
						},
					},
				},
			},
			"cloudlogs_endpoint": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Property values for the IBM Cloud Logs endpoint.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"target_crn": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The CRN of the IBM Cloud Logs instance.",
						},
					},
				},
			},
			"region": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_atracker_target", "region"),
				Description:  "Include this optional field if you want to create a target in a different region other than the one you are connected.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The crn of the target resource.",
			},
			"write_status": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The status of the write attempt to the target with the provided endpoint parameters.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status such as failed or success.",
						},
						"last_failure": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The timestamp of the failure.",
						},
						"reason_for_last_failure": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Detailed description of the cause of the failure.",
						},
					},
				},
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp of the target creation time.",
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The timestamp of the target last updated time.",
			},
			"message": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An optional message containing information about the target.",
			},
			"api_version": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The API version of the target.",
			},
		},
	}
}

func ResourceIBMAtrackerTargetValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[a-zA-Z0-9 -._:]+$`,
			MinValueLength:             1,
			MaxValueLength:             1000,
		},
		validate.ValidateSchema{
			Identifier:                 "target_type",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "cloud_logs, cloud_object_storage, event_streams",
		},
		validate.ValidateSchema{
			Identifier:                 "region",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-zA-Z0-9 -._:]+$`,
			MinValueLength:             3,
			MaxValueLength:             1000,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_atracker_target", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMAtrackerTargetCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	atrackerClient, err := meta.(conns.ClientSession).AtrackerV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_target", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createTargetOptions := &atrackerv2.CreateTargetOptions{}

	createTargetOptions.SetName(d.Get("name").(string))
	createTargetOptions.SetTargetType(d.Get("target_type").(string))
	if _, ok := d.GetOk("cos_endpoint"); ok {
		cosEndpointModel, err := ResourceIBMAtrackerTargetMapToCosEndpointPrototype(d.Get("cos_endpoint.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_target", "create", "parse-cos_endpoint").GetDiag()
		}
		createTargetOptions.SetCosEndpoint(cosEndpointModel)
	}
	if _, ok := d.GetOk("eventstreams_endpoint"); ok {
		eventstreamsEndpointModel, err := ResourceIBMAtrackerTargetMapToEventstreamsEndpointPrototype(d.Get("eventstreams_endpoint.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_target", "create", "parse-eventstreams_endpoint").GetDiag()
		}
		createTargetOptions.SetEventstreamsEndpoint(eventstreamsEndpointModel)
	}
	if _, ok := d.GetOk("cloudlogs_endpoint"); ok {
		cloudlogsEndpointModel, err := ResourceIBMAtrackerTargetMapToCloudLogsEndpointPrototype(d.Get("cloudlogs_endpoint.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_target", "create", "parse-cloudlogs_endpoint").GetDiag()
		}
		createTargetOptions.SetCloudlogsEndpoint(cloudlogsEndpointModel)
	}
	if _, ok := d.GetOk("region"); ok {
		createTargetOptions.SetRegion(d.Get("region").(string))
	}

	target, _, err := atrackerClient.CreateTargetWithContext(context, createTargetOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateTargetWithContext failed: %s", err.Error()), "ibm_atracker_target", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*target.ID)

	return resourceIBMAtrackerTargetRead(context, d, meta)
}

func resourceIBMAtrackerTargetRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	atrackerClient, err := meta.(conns.ClientSession).AtrackerV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_target", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getTargetOptions := &atrackerv2.GetTargetOptions{}

	getTargetOptions.SetID(d.Id())

	target, response, err := atrackerClient.GetTargetWithContext(context, getTargetOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetTargetWithContext failed: %s", err.Error()), "ibm_atracker_target", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("name", target.Name); err != nil {
		err = fmt.Errorf("Error setting name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_target", "read", "set-name").GetDiag()
	}
	if err = d.Set("target_type", target.TargetType); err != nil {
		err = fmt.Errorf("Error setting target_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_target", "read", "set-target_type").GetDiag()
	}
	// Don't report difference if the last parts of CRN are different
	if !core.IsNil(target.CosEndpoint) {
		cosEndpointMap, err := ResourceIBMAtrackerTargetCosEndpointToMap(target.CosEndpoint)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_target", "read", "cos_endpoint-to-map").GetDiag()
		}
		if cosInterface, ok := d.GetOk("cos_endpoint.0"); ok {
			targetCrnExisting := cosInterface.(map[string]interface{})["target_crn"].(string)
			targetCrnIncoming := cosEndpointMap["target_crn"].(*string)
			if len(targetCrnExisting) > 0 && targetCrnIncoming != nil {
				targetCrnExistingParts := strings.Split(targetCrnExisting, ":")
				targetCrnIncomingParts := strings.Split(*targetCrnIncoming, ":")
				isDifferent := false
				for i := 0; i < COS_CRN_PARTS && len(targetCrnExistingParts) > COS_CRN_PARTS-1 && len(targetCrnIncomingParts) > COS_CRN_PARTS-1; i++ {
					if targetCrnExistingParts[i] != targetCrnIncomingParts[i] {
						isDifferent = true
					}
				}
				if !isDifferent {
					cosEndpointMap["target_crn"] = targetCrnExisting
				}
			}
		}
		if err = d.Set("cos_endpoint", []map[string]interface{}{cosEndpointMap}); err != nil {
			err = fmt.Errorf("Error setting cos_endpoint: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_target", "read", "set-cos_endpoint").GetDiag()
		}
	}
	if !core.IsNil(target.EventstreamsEndpoint) {
		eventstreamsEndpointMap, err := ResourceIBMAtrackerTargetEventstreamsEndpointToMap(target.EventstreamsEndpoint)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_target", "read", "eventstreams_endpoint-to-map").GetDiag()
		}
		if err = d.Set("eventstreams_endpoint", []map[string]interface{}{eventstreamsEndpointMap}); err != nil {
			err = fmt.Errorf("Error setting eventstreams_endpoint: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_target", "read", "set-eventstreams_endpoint").GetDiag()
		}
	}
	if !core.IsNil(target.CloudlogsEndpoint) {
		cloudlogsEndpointMap, err := ResourceIBMAtrackerTargetCloudLogsEndpointToMap(target.CloudlogsEndpoint)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_target", "read", "cloudlogs_endpoint-to-map").GetDiag()
		}
		if err = d.Set("cloudlogs_endpoint", []map[string]interface{}{cloudlogsEndpointMap}); err != nil {
			err = fmt.Errorf("Error setting cloudlogs_endpoint: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_target", "read", "set-cloudlogs_endpoint").GetDiag()
		}
	}

	if !core.IsNil(target.CRN) {
		if err = d.Set("crn", target.CRN); err != nil {
			err = fmt.Errorf("Error setting crn: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_target", "read", "set-crn").GetDiag()
		}
	}

	if !core.IsNil(target.Region) {
		if len(*target.Region) > 0 {
			if err = d.Set("region", *target.Region); err != nil {
				err = fmt.Errorf("Error setting region: %s", err)
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_target", "read", "set-region").GetDiag()
			}
		}
	}

	writeStatusMap, err := ResourceIBMAtrackerTargetWriteStatusToMap(target.WriteStatus)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_target", "read", "write_status-to-map").GetDiag()
	}
	if err = d.Set("write_status", []map[string]interface{}{writeStatusMap}); err != nil {
		err = fmt.Errorf("Error setting write_status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_target", "read", "set-write_status").GetDiag()
	}
	if err = d.Set("created_at", flex.DateTimeToString(target.CreatedAt)); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_target", "read", "set-created_at").GetDiag()
	}
	if err = d.Set("updated_at", flex.DateTimeToString(target.UpdatedAt)); err != nil {
		err = fmt.Errorf("Error setting updated_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_target", "read", "set-updated_at").GetDiag()
	}
	if !core.IsNil(target.Message) {
		if err = d.Set("message", target.Message); err != nil {
			err = fmt.Errorf("Error setting message: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_target", "read", "set-message").GetDiag()
		}
	}
	if err = d.Set("api_version", flex.IntValue(target.APIVersion)); err != nil {
		err = fmt.Errorf("Error setting api_version: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_target", "read", "set-api_version").GetDiag()
	}

	return nil
}

func resourceIBMAtrackerTargetUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	atrackerClient, err := meta.(conns.ClientSession).AtrackerV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_target", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	replaceTargetOptions := &atrackerv2.ReplaceTargetOptions{}

	replaceTargetOptions.SetID(d.Id())

	hasChange := false

	if d.HasChange("name") || d.HasChange("cos_endpoint") || d.HasChange("eventstreams_endpoint") || d.HasChange("cloudlogs_endpoint") {
		if _, ok := d.GetOk("name"); ok {
			replaceTargetOptions.SetName(d.Get("name").(string))
		}
		if _, ok := d.GetOk("cos_endpoint.0"); ok {
			cosEndpoint, err := ResourceIBMAtrackerTargetMapToCosEndpointPrototype(d.Get("cos_endpoint.0").(map[string]interface{}))
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_target", "update", "parse-cos_endpoint").GetDiag()
			}
			replaceTargetOptions.SetCosEndpoint(cosEndpoint)
		}

		if _, ok := d.GetOk("eventstreams_endpoint.0"); ok {
			eventstreamsEndpoint, err := ResourceIBMAtrackerTargetMapToEventstreamsEndpointPrototype(d.Get("eventstreams_endpoint.0").(map[string]interface{}))
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_target", "update", "parse-eventstreams_endpoint").GetDiag()
			}
			replaceTargetOptions.SetEventstreamsEndpoint(eventstreamsEndpoint)
		}
		if _, ok := d.GetOk("cloudlogs_endpoint.0"); ok {
			cloudlogsEndpoint, err := ResourceIBMAtrackerTargetMapToCloudLogsEndpointPrototype(d.Get("cloudlogs_endpoint.0").(map[string]interface{}))
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_target", "update", "parse-cloudlogs_endpoint").GetDiag()
			}
			replaceTargetOptions.SetCloudlogsEndpoint(cloudlogsEndpoint)
		}

		hasChange = true
	}

	if hasChange {
		_, _, err = atrackerClient.ReplaceTargetWithContext(context, replaceTargetOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ReplaceTargetWithContext failed: %s", err.Error()), "ibm_atracker_target", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIBMAtrackerTargetRead(context, d, meta)
}

func resourceIBMAtrackerTargetDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	atrackerClient, err := meta.(conns.ClientSession).AtrackerV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_atracker_target", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteTargetOptions := &atrackerv2.DeleteTargetOptions{}

	deleteTargetOptions.SetID(d.Id())

	_, _, err = atrackerClient.DeleteTargetWithContext(context, deleteTargetOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteTargetWithContext failed: %s", err.Error()), "ibm_atracker_target", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIBMAtrackerTargetMapToCosEndpointPrototype(modelMap map[string]interface{}) (*atrackerv2.CosEndpointPrototype, error) {
	model := &atrackerv2.CosEndpointPrototype{}
	model.Endpoint = core.StringPtr(modelMap["endpoint"].(string))
	model.TargetCRN = core.StringPtr(modelMap["target_crn"].(string))
	model.Bucket = core.StringPtr(modelMap["bucket"].(string))
	if modelMap["api_key"] != nil && modelMap["api_key"].(string) != "" {
		model.APIKey = core.StringPtr(modelMap["api_key"].(string))
	}
	if modelMap["service_to_service_enabled"] != nil {
		model.ServiceToServiceEnabled = core.BoolPtr(modelMap["service_to_service_enabled"].(bool))
	}
	return model, nil
}

func ResourceIBMAtrackerTargetMapToEventstreamsEndpointPrototype(modelMap map[string]interface{}) (*atrackerv2.EventstreamsEndpointPrototype, error) {
	model := &atrackerv2.EventstreamsEndpointPrototype{}
	model.TargetCRN = core.StringPtr(modelMap["target_crn"].(string))
	brokers := []string{}
	for _, brokersItem := range modelMap["brokers"].([]interface{}) {
		brokers = append(brokers, brokersItem.(string))
	}
	model.Brokers = brokers
	model.Topic = core.StringPtr(modelMap["topic"].(string))
	if modelMap["api_key"] != nil && modelMap["api_key"].(string) != "" {
		model.APIKey = core.StringPtr(modelMap["api_key"].(string)) // pragma: whitelist secret
	}
	if modelMap["service_to_service_enabled"] != nil {
		model.ServiceToServiceEnabled = core.BoolPtr(modelMap["service_to_service_enabled"].(bool))
	}
	return model, nil
}

func ResourceIBMAtrackerTargetMapToCloudLogsEndpointPrototype(modelMap map[string]interface{}) (*atrackerv2.CloudLogsEndpointPrototype, error) {
	model := &atrackerv2.CloudLogsEndpointPrototype{}
	model.TargetCRN = core.StringPtr(modelMap["target_crn"].(string))
	return model, nil
}

func ResourceIBMAtrackerTargetCosEndpointToMap(model *atrackerv2.CosEndpoint) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["endpoint"] = model.Endpoint
	modelMap["target_crn"] = model.TargetCRN
	modelMap["bucket"] = model.Bucket
	// TODO: remove after deprecation
	modelMap["api_key"] = REDACTED_TEXT // pragma: whitelist secret
	modelMap["service_to_service_enabled"] = model.ServiceToServiceEnabled
	return modelMap, nil
}

func ResourceIBMAtrackerTargetEventstreamsEndpointToMap(model *atrackerv2.EventstreamsEndpoint) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["target_crn"] = model.TargetCRN
	modelMap["brokers"] = model.Brokers
	modelMap["topic"] = model.Topic
	// TODO: remove after deprecation
	modelMap["api_key"] = REDACTED_TEXT // pragma: whitelist secret
	modelMap["service_to_service_enabled"] = model.ServiceToServiceEnabled
	return modelMap, nil
}

func ResourceIBMAtrackerTargetCloudLogsEndpointToMap(model *atrackerv2.CloudLogsEndpoint) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["target_crn"] = model.TargetCRN
	return modelMap, nil
}

func ResourceIBMAtrackerTargetWriteStatusToMap(model *atrackerv2.WriteStatus) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["status"] = model.Status
	if model.LastFailure != nil {
		modelMap["last_failure"] = model.LastFailure.String()
	}
	if model.ReasonForLastFailure != nil {
		modelMap["reason_for_last_failure"] = model.ReasonForLastFailure
	}
	return modelMap, nil
}
