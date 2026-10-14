// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.114.0-a902401e-20260427-192904
 */

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

func ResourceIBMIsDynamicRouteServerRouteReportJob() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIsDynamicRouteServerRouteReportJobCreate,
		ReadContext:   resourceIBMIsDynamicRouteServerRouteReportJobRead,
		UpdateContext: resourceIBMIsDynamicRouteServerRouteReportJobUpdate,
		DeleteContext: resourceIBMIsDynamicRouteServerRouteReportJobDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"dynamic_route_server_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_route_report_job", "dynamic_route_server_id"),
				Description:  "The dynamic route server identifier.",
			},
			"format": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "json",
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_route_report_job", "format"),
				Description:  "The format used for the route report:`json` - The route report is generated based on the json schema.",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_dynamic_route_server_route_report_job", "name"),
				Description:  "The name for this dynamic route server route report export job. The name must not be used by another export job for the image. Changing the name will not affect the exported image name, `storage_object.name`, or `storage_href` values.",
			},
			"storage_bucket": &schema.Schema{
				Type:        schema.TypeList,
				MinItems:    1,
				MaxItems:    1,
				Required:    true,
				Description: "The Cloud Object Storage bucket of the exported dynamic route server route report object.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The CRN of this Cloud Object Storage bucket.",
						},
						"deleted": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"more_info": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "A link to documentation about deleted resources.",
									},
								},
							},
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The globally unique name of this Cloud Object Storage bucket.",
						},
						"resource_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The resource type.",
						},
					},
				},
			},
			"completed_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the dynamic route server route report export job was completed.If absent, the dynamic route server route report export job has not yet completed.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the dynamic route server route report export job was created.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this dynamic route server route report export job.",
			},
			"json_schema": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "JSON schema that defines the structure of the route report content.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The canonical URI of the JSON Schema that defines the dynamic route server route report.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of schema document.",
						},
						"version": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The version of the route report schema.",
						},
					},
				},
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			"started_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the dynamic route server route report export job started running.If absent, the export job has not yet started.",
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of this dynamic route server route report export job:- `deleting`:Dynamic route server route report export job is being deleted- `failed`:Dynamic route server route report export job could not be completed  successfully- `queued`:Dynamic route server route report export job is queued- `running`:Dynamic route server route report export job is in progress- `succeeded`:Dynamic route server route report export job was completed successfullyThe exported route report object is automatically deleted for `failed` jobs.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
			},
			"status_reasons": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The reasons for the current status (if any).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A snake case string succinctly identifying the status reason.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
						},
						"message": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An explanation of the status reason.",
						},
						"more_info": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "A link to documentation about this status reason.",
						},
					},
				},
			},
			"storage_href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Cloud Object Storage location of the exported dynamic route server route report object. The object at this location will not exist until the job completes successfully. The exported image object is not managed by the IBM VPC service, and may be removed or replaced with a different object by any user or service with IAM authorization to the storage bucket.",
			},
			"storage_object": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The Cloud Object Storage object for the exported image. This object will not exist untilthe job completes successfully. The exported dynamic route server route report object isnot managed by the IBM VPC service, and may be removed or replaced with a differentobject by any user or service with IAM authorization to the storage bucket.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of this Cloud Object Storage object. Names are unique within a Cloud Object Storage bucket.",
						},
					},
				},
			},
			"is_dynamic_route_server_route_report_job_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier for this dynamic route server route report export job.",
			},
		},
	}
}

func ResourceIBMIsDynamicRouteServerRouteReportJobValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "dynamic_route_server_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[-0-9a-z_]+$`,
			MinValueLength:             1,
			MaxValueLength:             64,
		},
		validate.ValidateSchema{
			Identifier:                 "format",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Optional:                   true,
			AllowedValues:              "json",
			Regexp:                     `^[a-z][a-z0-9]*(_[a-z0-9]+)*$`,
			MinValueLength:             1,
			MaxValueLength:             128,
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

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_dynamic_route_server_route_report_job", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMIsDynamicRouteServerRouteReportJobCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createDynamicRouteServerRouteReportJobOptions := &vpcv1.CreateDynamicRouteServerRouteReportJobOptions{}

	createDynamicRouteServerRouteReportJobOptions.SetDynamicRouteServerID(d.Get("dynamic_route_server_id").(string))
	storageBucketModel, err := ResourceIBMIsDynamicRouteServerRouteReportJobMapToDynamicRouteServerCloudObjectStorageBucketIdentity(d.Get("storage_bucket.0").(map[string]interface{}))
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "create", "parse-storage_bucket").GetDiag()
	}
	createDynamicRouteServerRouteReportJobOptions.SetStorageBucket(storageBucketModel)
	if _, ok := d.GetOk("format"); ok {
		createDynamicRouteServerRouteReportJobOptions.SetFormat(d.Get("format").(string))
	}
	if _, ok := d.GetOk("name"); ok {
		createDynamicRouteServerRouteReportJobOptions.SetName(d.Get("name").(string))
	}

	dynamicRouteServerRouteReportJob, _, err := vpcClient.CreateDynamicRouteServerRouteReportJobWithContext(context, createDynamicRouteServerRouteReportJobOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateDynamicRouteServerRouteReportJobWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_route_report_job", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *createDynamicRouteServerRouteReportJobOptions.DynamicRouteServerID, *dynamicRouteServerRouteReportJob.ID))

	return resourceIBMIsDynamicRouteServerRouteReportJobRead(context, d, meta)
}

func resourceIBMIsDynamicRouteServerRouteReportJobRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getDynamicRouteServerRouteReportJobOptions := &vpcv1.GetDynamicRouteServerRouteReportJobOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "sep-id-parts").GetDiag()
	}

	getDynamicRouteServerRouteReportJobOptions.SetDynamicRouteServerID(parts[0])
	getDynamicRouteServerRouteReportJobOptions.SetID(parts[1])

	dynamicRouteServerRouteReportJob, response, err := vpcClient.GetDynamicRouteServerRouteReportJobWithContext(context, getDynamicRouteServerRouteReportJobOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDynamicRouteServerRouteReportJobWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_route_report_job", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if !core.IsNil(dynamicRouteServerRouteReportJob.Format) {
		if err = d.Set("format", dynamicRouteServerRouteReportJob.Format); err != nil {
			err = fmt.Errorf("Error setting format: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "set-format").GetDiag()
		}
	}
	if !core.IsNil(dynamicRouteServerRouteReportJob.Name) {
		if err = d.Set("name", dynamicRouteServerRouteReportJob.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "set-name").GetDiag()
		}
	}
	storageBucketMap, err := ResourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerCloudObjectStorageBucketReferenceToMap(dynamicRouteServerRouteReportJob.StorageBucket)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "storage_bucket-to-map").GetDiag()
	}
	if err = d.Set("storage_bucket", []map[string]interface{}{storageBucketMap}); err != nil {
		err = fmt.Errorf("Error setting storage_bucket: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "set-storage_bucket").GetDiag()
	}
	if !core.IsNil(dynamicRouteServerRouteReportJob.CompletedAt) {
		if err = d.Set("completed_at", flex.DateTimeToString(dynamicRouteServerRouteReportJob.CompletedAt)); err != nil {
			err = fmt.Errorf("Error setting completed_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "set-completed_at").GetDiag()
		}
	}
	if err = d.Set("created_at", flex.DateTimeToString(dynamicRouteServerRouteReportJob.CreatedAt)); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "set-created_at").GetDiag()
	}
	if err = d.Set("href", dynamicRouteServerRouteReportJob.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "set-href").GetDiag()
	}
	if !core.IsNil(dynamicRouteServerRouteReportJob.JSONSchema) {
		jsonSchemaMap, err := ResourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerRouteReportSchemaToMap(dynamicRouteServerRouteReportJob.JSONSchema)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "json_schema-to-map").GetDiag()
		}
		if err = d.Set("json_schema", []map[string]interface{}{jsonSchemaMap}); err != nil {
			err = fmt.Errorf("Error setting json_schema: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "set-json_schema").GetDiag()
		}
	}
	if err = d.Set("resource_type", dynamicRouteServerRouteReportJob.ResourceType); err != nil {
		err = fmt.Errorf("Error setting resource_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "set-resource_type").GetDiag()
	}
	if !core.IsNil(dynamicRouteServerRouteReportJob.StartedAt) {
		if err = d.Set("started_at", flex.DateTimeToString(dynamicRouteServerRouteReportJob.StartedAt)); err != nil {
			err = fmt.Errorf("Error setting started_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "set-started_at").GetDiag()
		}
	}
	if err = d.Set("status", dynamicRouteServerRouteReportJob.Status); err != nil {
		err = fmt.Errorf("Error setting status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "set-status").GetDiag()
	}
	statusReasons := []map[string]interface{}{}
	for _, statusReasonsItem := range dynamicRouteServerRouteReportJob.StatusReasons {
		statusReasonsItemMap, err := ResourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerRouteReportJobStatusReasonToMap(&statusReasonsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "status_reasons-to-map").GetDiag()
		}
		statusReasons = append(statusReasons, statusReasonsItemMap)
	}
	if err = d.Set("status_reasons", statusReasons); err != nil {
		err = fmt.Errorf("Error setting status_reasons: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "set-status_reasons").GetDiag()
	}
	if err = d.Set("storage_href", dynamicRouteServerRouteReportJob.StorageHref); err != nil {
		err = fmt.Errorf("Error setting storage_href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "set-storage_href").GetDiag()
	}
	storageObjectMap, err := ResourceIBMIsDynamicRouteServerRouteReportJobCloudObjectStorageObjectReferenceToMap(dynamicRouteServerRouteReportJob.StorageObject)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "storage_object-to-map").GetDiag()
	}
	if err = d.Set("storage_object", []map[string]interface{}{storageObjectMap}); err != nil {
		err = fmt.Errorf("Error setting storage_object: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "set-storage_object").GetDiag()
	}
	if err = d.Set("is_dynamic_route_server_route_report_job_id", dynamicRouteServerRouteReportJob.ID); err != nil {
		err = fmt.Errorf("Error setting is_dynamic_route_server_route_report_job_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "read", "set-is_dynamic_route_server_route_report_job_id").GetDiag()
	}

	return nil
}

func resourceIBMIsDynamicRouteServerRouteReportJobUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateDynamicRouteServerRouteReportJobOptions := &vpcv1.UpdateDynamicRouteServerRouteReportJobOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "update", "sep-id-parts").GetDiag()
	}

	updateDynamicRouteServerRouteReportJobOptions.SetDynamicRouteServerID(parts[0])
	updateDynamicRouteServerRouteReportJobOptions.SetID(parts[1])

	hasChange := false

	patchVals := &vpcv1.DynamicRouteServerRouteReportJobPatch{}
	if d.HasChange("dynamic_route_server_id") {
		errMsg := fmt.Sprintf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "dynamic_route_server_id")
		return flex.DiscriminatedTerraformErrorf(nil, errMsg, "ibm_is_dynamic_route_server_route_report_job", "update", "dynamic_route_server_id-forces-new").GetDiag()
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
		updateDynamicRouteServerRouteReportJobOptions.DynamicRouteServerRouteReportJobPatch = ResourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerRouteReportJobPatchAsPatch(patchVals, d)

		_, _, err = vpcClient.UpdateDynamicRouteServerRouteReportJobWithContext(context, updateDynamicRouteServerRouteReportJobOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateDynamicRouteServerRouteReportJobWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_route_report_job", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIBMIsDynamicRouteServerRouteReportJobRead(context, d, meta)
}

func resourceIBMIsDynamicRouteServerRouteReportJobDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteDynamicRouteServerRouteReportJobOptions := &vpcv1.DeleteDynamicRouteServerRouteReportJobOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_dynamic_route_server_route_report_job", "delete", "sep-id-parts").GetDiag()
	}

	deleteDynamicRouteServerRouteReportJobOptions.SetDynamicRouteServerID(parts[0])
	deleteDynamicRouteServerRouteReportJobOptions.SetID(parts[1])

	_, err = vpcClient.DeleteDynamicRouteServerRouteReportJobWithContext(context, deleteDynamicRouteServerRouteReportJobOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteDynamicRouteServerRouteReportJobWithContext failed: %s", err.Error()), "ibm_is_dynamic_route_server_route_report_job", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIBMIsDynamicRouteServerRouteReportJobMapToDynamicRouteServerCloudObjectStorageBucketIdentity(modelMap map[string]interface{}) (vpcv1.DynamicRouteServerCloudObjectStorageBucketIdentityIntf, error) {
	model := &vpcv1.DynamicRouteServerCloudObjectStorageBucketIdentity{}
	if modelMap["crn"] != nil && modelMap["crn"].(string) != "" {
		model.CRN = core.StringPtr(modelMap["crn"].(string))
	}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	return model, nil
}

func ResourceIBMIsDynamicRouteServerRouteReportJobMapToDynamicRouteServerCloudObjectStorageBucketIdentityCloudObjectStorageBucketIdentityByCRN(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerCloudObjectStorageBucketIdentityCloudObjectStorageBucketIdentityByCRN, error) {
	model := &vpcv1.DynamicRouteServerCloudObjectStorageBucketIdentityCloudObjectStorageBucketIdentityByCRN{}
	model.CRN = core.StringPtr(modelMap["crn"].(string))
	return model, nil
}

func ResourceIBMIsDynamicRouteServerRouteReportJobMapToDynamicRouteServerCloudObjectStorageBucketIdentityCloudObjectStorageBucketIdentityByName(modelMap map[string]interface{}) (*vpcv1.DynamicRouteServerCloudObjectStorageBucketIdentityCloudObjectStorageBucketIdentityByName, error) {
	model := &vpcv1.DynamicRouteServerCloudObjectStorageBucketIdentityCloudObjectStorageBucketIdentityByName{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	return model, nil
}

func ResourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerCloudObjectStorageBucketReferenceToMap(model *vpcv1.DynamicRouteServerCloudObjectStorageBucketReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	if model.Deleted != nil {
		deletedMap, err := ResourceIBMIsDynamicRouteServerRouteReportJobDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	modelMap["name"] = *model.Name
	modelMap["resource_type"] = *model.ResourceType
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerRouteReportJobDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = *model.MoreInfo
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerRouteReportSchemaToMap(model *vpcv1.DynamicRouteServerRouteReportSchema) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["href"] = *model.Href
	modelMap["type"] = *model.Type
	modelMap["version"] = *model.Version
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerRouteReportJobStatusReasonToMap(model *vpcv1.DynamicRouteServerRouteReportJobStatusReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["code"] = *model.Code
	modelMap["message"] = *model.Message
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerRouteReportJobCloudObjectStorageObjectReferenceToMap(model *vpcv1.CloudObjectStorageObjectReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = *model.Name
	return modelMap, nil
}

func ResourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerRouteReportJobPatchAsPatch(patchVals *vpcv1.DynamicRouteServerRouteReportJobPatch, d *schema.ResourceData) map[string]interface{} {
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
