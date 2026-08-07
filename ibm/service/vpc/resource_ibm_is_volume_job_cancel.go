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

func ResourceIBMIsVolumeJobCancel() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIsVolumeJobCancelCreate,
		ReadContext:   resourceIBMIsVolumeJobCancelRead,
		DeleteContext: resourceIBMIsVolumeJobCancelDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"volume_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_volume_job_cancel", "volume_id"),
				Description:  "The volume identifier.",
			},
			"job_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of volume job.The enumerated values for this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name for this volume job. The name must not be used by another volume job for this volume.",
			},
			"parameters": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The parameters to use after the volume is migrated.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bandwidth": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The maximum bandwidth (in megabits per second) for the volume.If specified, the volume profile must not have a `bandwidth.type` of `dependent`.",
						},
						"iops": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
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
										Computed:    true,
										Description: "The globally unique name for this volume profile.",
									},
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
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
							Computed:    true,
							Description: "A link to documentation about this status reason.",
						},
					},
				},
			},
			"volume_job_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The unique identifier for this volume job.",
			},
		},
	}
}

func ResourceIBMIsVolumeJobCancelValidator() *validate.ResourceValidator {
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
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_volume_job_cancel", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMIsVolumeJobCancelCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job_cancel", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	volume_id := d.Get("volume_id").(string)
	volume_job_id := d.Get("volume_job_id").(string)
	cancelVolumeJobOptions := &vpcv1.CancelVolumeJobOptions{
		VolumeID: &volume_id,
		ID:       &volume_job_id,
	}

	volumeJobIntf, _, err := vpcClient.CancelVolumeJobWithContext(context, cancelVolumeJobOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CancelVolumeJobWithContext failed: %s", err.Error()), "ibm_is_volume_job_cancel", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	volumeJob := volumeJobIntf.(*vpcv1.VolumeJob)
	d.SetId(fmt.Sprintf("%s/%s", *cancelVolumeJobOptions.VolumeID, *volumeJob.ID))

	return resourceIBMIsVolumeJobRead(context, d, meta)
}

func resourceIBMIsVolumeJobCancelRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job_cancel", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getVolumeJobOptions := &vpcv1.GetVolumeJobOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job_cancel", "read", "sep-id-parts").GetDiag()
	}

	getVolumeJobOptions.SetVolumeID(parts[0])
	getVolumeJobOptions.SetID(parts[1])

	volumeJobIntf, response, err := vpcClient.GetVolumeJobWithContext(context, getVolumeJobOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVolumeJobWithContext failed: %s", err.Error()), "ibm_is_volume_job_cancel", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	volumeJob := volumeJobIntf.(*vpcv1.VolumeJob)
	if err = d.Set("job_type", volumeJob.JobType); err != nil {
		err = fmt.Errorf("Error setting job_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job_cancel", "read", "set-job_type").GetDiag()
	}
	if !core.IsNil(volumeJob.Name) {
		if err = d.Set("name", volumeJob.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job_cancel", "read", "set-name").GetDiag()
		}
	}
	if !core.IsNil(volumeJob.Parameters) {
		parametersMap, err := ResourceIBMIsVolumeJobVolumeJobTypeMigrateParametersToMap(volumeJob.Parameters)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job_cancel", "read", "parameters-to-map").GetDiag()
		}
		if err = d.Set("parameters", []map[string]interface{}{parametersMap}); err != nil {
			err = fmt.Errorf("Error setting parameters: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job_cancel", "read", "set-parameters").GetDiag()
		}
	}
	if err = d.Set("auto_delete", volumeJob.AutoDelete); err != nil {
		err = fmt.Errorf("Error setting auto_delete: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job_cancel", "read", "set-auto_delete").GetDiag()
	}
	if !core.IsNil(volumeJob.CompletedAt) {
		if err = d.Set("completed_at", flex.DateTimeToString(volumeJob.CompletedAt)); err != nil {
			err = fmt.Errorf("Error setting completed_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job_cancel", "read", "set-completed_at").GetDiag()
		}
	}
	if err = d.Set("created_at", flex.DateTimeToString(volumeJob.CreatedAt)); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job_cancel", "read", "set-created_at").GetDiag()
	}
	if !core.IsNil(volumeJob.EstimatedCompletionAt) {
		if err = d.Set("estimated_completion_at", flex.DateTimeToString(volumeJob.EstimatedCompletionAt)); err != nil {
			err = fmt.Errorf("Error setting estimated_completion_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job_cancel", "read", "set-estimated_completion_at").GetDiag()
		}
	}
	if err = d.Set("href", volumeJob.Href); err != nil {
		err = fmt.Errorf("Error setting href: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job_cancel", "read", "set-href").GetDiag()
	}
	if err = d.Set("resource_type", volumeJob.ResourceType); err != nil {
		err = fmt.Errorf("Error setting resource_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job_cancel", "read", "set-resource_type").GetDiag()
	}
	if !core.IsNil(volumeJob.StartedAt) {
		if err = d.Set("started_at", flex.DateTimeToString(volumeJob.StartedAt)); err != nil {
			err = fmt.Errorf("Error setting started_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job_cancel", "read", "set-started_at").GetDiag()
		}
	}
	if err = d.Set("status", volumeJob.Status); err != nil {
		err = fmt.Errorf("Error setting status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job_cancel", "read", "set-status").GetDiag()
	}
	statusReasons := []map[string]interface{}{}
	for _, statusReasonsItem := range volumeJob.StatusReasons {
		statusReasonsItemMap, err := ResourceIBMIsVolumeJobVolumeJobStatusReasonToMap(&statusReasonsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job_cancel", "read", "status_reasons-to-map").GetDiag()
		}
		statusReasons = append(statusReasons, statusReasonsItemMap)
	}
	if err = d.Set("status_reasons", statusReasons); err != nil {
		err = fmt.Errorf("Error setting status_reasons: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job_cancel", "read", "set-status_reasons").GetDiag()
	}
	if err = d.Set("volume_job_id", volumeJob.ID); err != nil {
		err = fmt.Errorf("Error setting volume_job_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_is_volume_job_cancel", "read", "set-volume_job_id").GetDiag()
	}

	return nil
}

func resourceIBMIsVolumeJobCancelDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	d.SetId("")

	return nil
}
