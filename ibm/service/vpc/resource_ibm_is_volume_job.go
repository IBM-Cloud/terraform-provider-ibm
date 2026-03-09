// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.107.1-41b0fbd0-20250825-080732
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

func ResourceIBMIsVolumeJob() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIsVolumeJobCreate,
		ReadContext:   resourceIBMIsVolumeJobRead,
		UpdateContext: resourceIBMIsVolumeJobUpdate,
		DeleteContext: resourceIBMIsVolumeJobDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"volume_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_volume_job", "volume_id"),
				Description:  "The volume identifier.",
			},
			"start": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_volume_job", "start"),
				Description:  "A server-provided token determining what resource to start the page on.",
			},
			"limit": &schema.Schema{
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      50,
				ValidateFunc: validate.InvokeValidator("ibm_is_volume_job", "limit"),
				Description:  "The number of resources to return on a page.",
			},
			"job_type": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_volume_job", "job_type"),
				Description:  "The type of volume job.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_volume_job", "name"),
				Description:  "The name for this volume job. The name must not be used by another volume job for this volume.",
			},
			"parameters": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "The parameters to use after the volume is migrated.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bandwidth": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The maximum bandwidth (in megabits per second) for the volume.If specified, the volume profile must not have a `bandwidth.type` of `dependent`.",
						},
						"iops": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The maximum I/O operations per second (IOPS) for this volume.If specified, the volume profile must not have a `iops.type` of `dependent`.",
						},
						"profile": &schema.Schema{
							Type:        schema.TypeList,
							MinItems:    1,
							MaxItems:    1,
							Required:    true,
							Description: "Identifies a volume profile by a unique property.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The globally unique name for this volume profile.",
									},
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The URL for this volume profile.",
									},
								},
							},
						},
					},
				},
			},
			"auto_delete": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether this volume job will be automatically deleted after it completes. At present, this is always `false`, but may be modifiable in the future.",
			},
			"completed_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the volume job was completed.If absent, the volume job has not yet completed.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the volume job was created.",
			},
			"estimated_completion_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the volume job is estimated to complete.If absent, the volume job is still queued and has not yet started.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this volume job.",
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			"started_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the volume job was started.If absent, the volume job has not yet started.",
			},
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of this volume job:- `deleting`:   job is being deleted- `failed`:     job could not be completed successfully- `queued`:     job is queued- `running`:    job is in progress- `succeeded`:  job was completed successfully- `canceling`: job is being canceled- `canceled`:  job is canceledThe enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
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
			"is_volume_job_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier for this volume job.",
			},
		},
	}
}

func ResourceIBMIsVolumeJobValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "volume_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[-0-9a-z_]+$`,
			MinValueLength:             1,
			MaxValueLength:             64,
		},
		validate.ValidateSchema{
			Identifier:                 "start",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[ -~]+$`,
			MinValueLength:             1,
			MaxValueLength:             4096,
		},
		validate.ValidateSchema{
			Identifier:                 "limit",
			ValidateFunctionIdentifier: validate.IntBetween,
			Type:                       validate.TypeInt,
			Optional:                   true,
			MinValue:                   "1",
			MaxValue:                   "100",
		},
		validate.ValidateSchema{
			Identifier:                 "job_type",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "migrate",
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

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_volume_job", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMIsVolumeJobCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	bodyModelMap := map[string]interface{}{}
	createVolumeJobOptions := &vpcv1.CreateVolumeJobOptions{}

	bodyModelMap["job_type"] = d.Get("job_type")
	if _, ok := d.GetOk("name"); ok {
		bodyModelMap["name"] = d.Get("name")
	}
	if _, ok := d.GetOk("parameters"); ok {
		bodyModelMap["parameters"] = d.Get("parameters")
	}
	createVolumeJobOptions.SetVolumeID(d.Get("volume_id").(string))
	if _, ok := d.GetOk("start"); ok {
		createVolumeJobOptions.SetStart(d.Get("start").(string))
	}
	if _, ok := d.GetOk("limit"); ok {
		createVolumeJobOptions.SetLimit(int64(d.Get("limit").(int)))
	}
	convertedModel, err := ResourceIBMIsVolumeJobMapToVolumeJobPrototype(bodyModelMap)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job", "create", "parse-request-body").GetDiag()
	}
	createVolumeJobOptions.VolumeJobPrototype = convertedModel

	volumeJobIntf, _, err := vpcClient.CreateVolumeJobWithContext(context, createVolumeJobOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateVolumeJobWithContext failed: %s", err.Error()), "ibm_is_volume_job", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	volumeJob := volumeJobIntf.(*vpcv1.VolumeJob)
	d.SetId(fmt.Sprintf("%s/%s", *createVolumeJobOptions.VolumeID, *volumeJob.ID))

	return resourceIBMIsVolumeJobRead(context, d, meta)
}

func resourceIBMIsVolumeJobRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getVolumeJobOptions := &vpcv1.GetVolumeJobOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job", "read", "sep-id-parts").GetDiag()
	}

	getVolumeJobOptions.SetVolumeID(parts[0])
	getVolumeJobOptions.SetID(parts[1])

	volumeJobIntf, response, err := vpcClient.GetVolumeJobWithContext(context, getVolumeJobOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVolumeJobWithContext failed: %s", err.Error()), "ibm_is_volume_job", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	volumeJob := volumeJobIntf.(*vpcv1.VolumeJob)
	if err = d.Set("job_type", volumeJob.JobType); err != nil {
		err = fmt.Errorf("Error setting job_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job", "read", "set-job_type").GetDiag()
	}
	if !core.IsNil(volumeJob.Name) {
		if err = d.Set("name", volumeJob.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job", "read", "set-name").GetDiag()
		}
	}
	if !core.IsNil(volumeJob.Parameters) {
		parametersMap, err := ResourceIBMIsVolumeJobVolumeJobTypeMigrateParametersToMap(volumeJob.Parameters)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job", "read", "parameters-to-map").GetDiag()
		}
		if err = d.Set("parameters", []map[string]interface{}{parametersMap}); err != nil {
			err = fmt.Errorf("Error setting parameters: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job", "read", "set-parameters").GetDiag()
		}
	}
	if err = d.Set("auto_delete", volumeJob.AutoDelete); err != nil {
		err = fmt.Errorf("Error setting auto_delete: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job", "read", "set-auto_delete").GetDiag()
	}
	if !core.IsNil(volumeJob.CompletedAt) {
		if err = d.Set("completed_at", flex.DateTimeToString(volumeJob.CompletedAt)); err != nil {
			err = fmt.Errorf("Error setting completed_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job", "read", "set-completed_at").GetDiag()
		}
	}
	if err = d.Set("created_at", flex.DateTimeToString(volumeJob.CreatedAt)); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job", "read", "set-created_at").GetDiag()
	}
	if !core.IsNil(volumeJob.EstimatedCompletionAt) {
		if err = d.Set("estimated_completion_at", flex.DateTimeToString(volumeJob.EstimatedCompletionAt)); err != nil {
			err = fmt.Errorf("Error setting estimated_completion_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job", "read", "set-estimated_completion_at").GetDiag()
		}
	}
	if err = d.Set("href", volumeJob.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job", "read", "set-href").GetDiag()
	}
	if err = d.Set("resource_type", volumeJob.ResourceType); err != nil {
		err = fmt.Errorf("Error setting resource_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job", "read", "set-resource_type").GetDiag()
	}
	if !core.IsNil(volumeJob.StartedAt) {
		if err = d.Set("started_at", flex.DateTimeToString(volumeJob.StartedAt)); err != nil {
			err = fmt.Errorf("Error setting started_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job", "read", "set-started_at").GetDiag()
		}
	}
	if err = d.Set("status", volumeJob.Status); err != nil {
		err = fmt.Errorf("Error setting status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job", "read", "set-status").GetDiag()
	}
	statusReasons := []map[string]interface{}{}
	for _, statusReasonsItem := range volumeJob.StatusReasons {
		statusReasonsItemMap, err := ResourceIBMIsVolumeJobVolumeJobStatusReasonToMap(&statusReasonsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job", "read", "status_reasons-to-map").GetDiag()
		}
		statusReasons = append(statusReasons, statusReasonsItemMap)
	}
	if err = d.Set("status_reasons", statusReasons); err != nil {
		err = fmt.Errorf("Error setting status_reasons: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job", "read", "set-status_reasons").GetDiag()
	}
	if err = d.Set("is_volume_job_id", volumeJob.ID); err != nil {
		err = fmt.Errorf("Error setting is_volume_job_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job", "read", "set-is_volume_job_id").GetDiag()
	}

	return nil
}

func resourceIBMIsVolumeJobUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateVolumeJobOptions := &vpcv1.UpdateVolumeJobOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job", "update", "sep-id-parts").GetDiag()
	}

	updateVolumeJobOptions.SetVolumeID(parts[0])
	updateVolumeJobOptions.SetID(parts[1])

	hasChange := false

	patchVals := &vpcv1.VolumeJobPatch{}
	if d.HasChange("volume_id") {
		errMsg := fmt.Sprintf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "volume_id")
		return flex.DiscriminatedTerraformErrorf(nil, errMsg, "ibm_is_volume_job", "update", "volume_id-forces-new").GetDiag()
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
		updateVolumeJobOptions.VolumeJobPatch = ResourceIBMIsVolumeJobVolumeJobPatchAsPatch(patchVals, d)

		_, _, err = vpcClient.UpdateVolumeJobWithContext(context, updateVolumeJobOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateVolumeJobWithContext failed: %s", err.Error()), "ibm_is_volume_job", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIBMIsVolumeJobRead(context, d, meta)
}

func resourceIBMIsVolumeJobDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteVolumeJobOptions := &vpcv1.DeleteVolumeJobOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job", "delete", "sep-id-parts").GetDiag()
	}

	deleteVolumeJobOptions.SetVolumeID(parts[0])
	deleteVolumeJobOptions.SetID(parts[1])

	_, err = vpcClient.DeleteVolumeJobWithContext(context, deleteVolumeJobOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteVolumeJobWithContext failed: %s", err.Error()), "ibm_is_volume_job", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIBMIsVolumeJobMapToVolumeJobTypeMigrateParameters(modelMap map[string]interface{}) (*vpcv1.VolumeJobTypeMigrateParameters, error) {
	model := &vpcv1.VolumeJobTypeMigrateParameters{}
	if modelMap["bandwidth"] != nil {
		model.Bandwidth = core.Int64Ptr(int64(modelMap["bandwidth"].(int)))
	}
	if modelMap["iops"] != nil {
		model.Iops = core.Int64Ptr(int64(modelMap["iops"].(int)))
	}
	ProfileModel, err := ResourceIBMIsVolumeJobMapToVolumeProfileIdentity(modelMap["profile"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Profile = ProfileModel
	return model, nil
}

func ResourceIBMIsVolumeJobMapToVolumeProfileIdentity(modelMap map[string]interface{}) (vpcv1.VolumeProfileIdentityIntf, error) {
	model := &vpcv1.VolumeProfileIdentity{}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["href"] != nil && modelMap["href"].(string) != "" {
		model.Href = core.StringPtr(modelMap["href"].(string))
	}
	return model, nil
}

func ResourceIBMIsVolumeJobMapToVolumeProfileIdentityByName(modelMap map[string]interface{}) (*vpcv1.VolumeProfileIdentityByName, error) {
	model := &vpcv1.VolumeProfileIdentityByName{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	return model, nil
}

func ResourceIBMIsVolumeJobMapToVolumeProfileIdentityByHref(modelMap map[string]interface{}) (*vpcv1.VolumeProfileIdentityByHref, error) {
	model := &vpcv1.VolumeProfileIdentityByHref{}
	model.Href = core.StringPtr(modelMap["href"].(string))
	return model, nil
}

func ResourceIBMIsVolumeJobMapToVolumeJobPrototype(modelMap map[string]interface{}) (vpcv1.VolumeJobPrototypeIntf, error) {
	model := &vpcv1.VolumeJobPrototype{}
	model.JobType = core.StringPtr(modelMap["job_type"].(string))
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["parameters"] != nil && len(modelMap["parameters"].([]interface{})) > 0 {
		ParametersModel, err := ResourceIBMIsVolumeJobMapToVolumeJobTypeMigrateParameters(modelMap["parameters"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Parameters = ParametersModel
	}
	return model, nil
}

func ResourceIBMIsVolumeJobMapToVolumeJobPrototypeVolumeJobTypeMigratePrototype(modelMap map[string]interface{}) (*vpcv1.VolumeJobPrototypeVolumeJobTypeMigratePrototype, error) {
	model := &vpcv1.VolumeJobPrototypeVolumeJobTypeMigratePrototype{}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	model.JobType = core.StringPtr(modelMap["job_type"].(string))
	ParametersModel, err := ResourceIBMIsVolumeJobMapToVolumeJobTypeMigrateParameters(modelMap["parameters"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Parameters = ParametersModel
	return model, nil
}

func ResourceIBMIsVolumeJobVolumeJobTypeMigrateParametersToMap(model *vpcv1.VolumeJobTypeMigrateParameters) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Bandwidth != nil {
		modelMap["bandwidth"] = flex.IntValue(model.Bandwidth)
	}
	if model.Iops != nil {
		modelMap["iops"] = flex.IntValue(model.Iops)
	}
	profileMap, err := ResourceIBMIsVolumeJobVolumeProfileIdentityToMap(model.Profile)
	if err != nil {
		return modelMap, err
	}
	modelMap["profile"] = []map[string]interface{}{profileMap}
	return modelMap, nil
}

func ResourceIBMIsVolumeJobVolumeProfileIdentityToMap(model vpcv1.VolumeProfileIdentityIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.VolumeProfileIdentityByName); ok {
		return ResourceIBMIsVolumeJobVolumeProfileIdentityByNameToMap(model.(*vpcv1.VolumeProfileIdentityByName))
	} else if _, ok := model.(*vpcv1.VolumeProfileIdentityByHref); ok {
		return ResourceIBMIsVolumeJobVolumeProfileIdentityByHrefToMap(model.(*vpcv1.VolumeProfileIdentityByHref))
	} else if _, ok := model.(*vpcv1.VolumeProfileIdentity); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.VolumeProfileIdentity)
		if model.Name != nil {
			modelMap["name"] = *model.Name
		}
		if model.Href != nil {
			modelMap["href"] = *model.Href
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.VolumeProfileIdentityIntf subtype encountered")
	}
}

func ResourceIBMIsVolumeJobVolumeProfileIdentityByNameToMap(model *vpcv1.VolumeProfileIdentityByName) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = *model.Name
	return modelMap, nil
}

func ResourceIBMIsVolumeJobVolumeProfileIdentityByHrefToMap(model *vpcv1.VolumeProfileIdentityByHref) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["href"] = *model.Href
	return modelMap, nil
}

func ResourceIBMIsVolumeJobVolumeJobStatusReasonToMap(model *vpcv1.VolumeJobStatusReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["code"] = *model.Code
	modelMap["message"] = *model.Message
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func ResourceIBMIsVolumeJobVolumeJobPatchAsPatch(patchVals *vpcv1.VolumeJobPatch, d *schema.ResourceData) map[string]interface{} {
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
