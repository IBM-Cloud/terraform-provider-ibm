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
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func DataSourceIBMIsVolumeJob() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsVolumeJobRead,

		Schema: map[string]*schema.Schema{
			"volume_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The volume identifier.",
			},
			"is_volume_job_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The volume job identifier.",
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
							Computed:    true,
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
		},
	}
}

func dataSourceIBMIsVolumeJobRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_volume_job", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getVolumeJobOptions := &vpcv1.GetVolumeJobOptions{}

	getVolumeJobOptions.SetVolumeID(d.Get("volume_id").(string))
	getVolumeJobOptions.SetID(d.Get("is_volume_job_id").(string))

	volumeJobIntf, _, err := vpcClient.GetVolumeJobWithContext(context, getVolumeJobOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetVolumeJobWithContext failed: %s", err.Error()), "(Data) ibm_is_volume_job", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	volumeJob := volumeJobIntf.(*vpcv1.VolumeJob)

	d.SetId(fmt.Sprintf("%s/%s", *getVolumeJobOptions.VolumeID, *getVolumeJobOptions.ID))

	if err = d.Set("auto_delete", volumeJob.AutoDelete); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting auto_delete: %s", err), "(Data) ibm_is_volume_job", "read", "set-auto_delete").GetDiag()
	}

	if !core.IsNil(volumeJob.CompletedAt) {
		if err = d.Set("completed_at", flex.DateTimeToString(volumeJob.CompletedAt)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting completed_at: %s", err), "(Data) ibm_is_volume_job", "read", "set-completed_at").GetDiag()
		}
	}

	if err = d.Set("created_at", flex.DateTimeToString(volumeJob.CreatedAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_is_volume_job", "read", "set-created_at").GetDiag()
	}

	if !core.IsNil(volumeJob.EstimatedCompletionAt) {
		if err = d.Set("estimated_completion_at", flex.DateTimeToString(volumeJob.EstimatedCompletionAt)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting estimated_completion_at: %s", err), "(Data) ibm_is_volume_job", "read", "set-estimated_completion_at").GetDiag()
		}
	}

	if err = d.Set("href", volumeJob.Href); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_volume_job", "read", "set-href").GetDiag()
	}

	if err = d.Set("job_type", volumeJob.JobType); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting job_type: %s", err), "(Data) ibm_is_volume_job", "read", "set-job_type").GetDiag()
	}

	if err = d.Set("name", volumeJob.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_volume_job", "read", "set-name").GetDiag()
	}

	if err = d.Set("resource_type", volumeJob.ResourceType); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_type: %s", err), "(Data) ibm_is_volume_job", "read", "set-resource_type").GetDiag()
	}

	if !core.IsNil(volumeJob.StartedAt) {
		if err = d.Set("started_at", flex.DateTimeToString(volumeJob.StartedAt)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting started_at: %s", err), "(Data) ibm_is_volume_job", "read", "set-started_at").GetDiag()
		}
	}

	if err = d.Set("status", volumeJob.Status); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_is_volume_job", "read", "set-status").GetDiag()
	}

	statusReasons := []map[string]interface{}{}
	for _, statusReasonsItem := range volumeJob.StatusReasons {
		statusReasonsItemMap, err := DataSourceIBMIsVolumeJobVolumeJobStatusReasonToMap(&statusReasonsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_volume_job", "read", "status_reasons-to-map").GetDiag()
		}
		statusReasons = append(statusReasons, statusReasonsItemMap)
	}
	if err = d.Set("status_reasons", statusReasons); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status_reasons: %s", err), "(Data) ibm_is_volume_job", "read", "set-status_reasons").GetDiag()
	}

	if !core.IsNil(volumeJob.Parameters) {
		parameters := []map[string]interface{}{}
		parametersMap, err := DataSourceIBMIsVolumeJobVolumeJobTypeMigrateParametersToMap(volumeJob.Parameters)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_volume_job", "read", "parameters-to-map").GetDiag()
		}
		parameters = append(parameters, parametersMap)
		if err = d.Set("parameters", parameters); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting parameters: %s", err), "(Data) ibm_is_volume_job", "read", "set-parameters").GetDiag()
		}
	}

	return nil
}

func DataSourceIBMIsVolumeJobVolumeJobStatusReasonToMap(model *vpcv1.VolumeJobStatusReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["code"] = *model.Code
	modelMap["message"] = *model.Message
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func DataSourceIBMIsVolumeJobVolumeJobTypeMigrateParametersToMap(model *vpcv1.VolumeJobTypeMigrateParameters) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Bandwidth != nil {
		modelMap["bandwidth"] = flex.IntValue(model.Bandwidth)
	}
	if model.Iops != nil {
		modelMap["iops"] = flex.IntValue(model.Iops)
	}
	profileMap, err := DataSourceIBMIsVolumeJobVolumeProfileIdentityToMap(model.Profile)
	if err != nil {
		return modelMap, err
	}
	modelMap["profile"] = []map[string]interface{}{profileMap}
	return modelMap, nil
}

func DataSourceIBMIsVolumeJobVolumeProfileIdentityToMap(model vpcv1.VolumeProfileIdentityIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.VolumeProfileIdentityByName); ok {
		return DataSourceIBMIsVolumeJobVolumeProfileIdentityByNameToMap(model.(*vpcv1.VolumeProfileIdentityByName))
	} else if _, ok := model.(*vpcv1.VolumeProfileIdentityByHref); ok {
		return DataSourceIBMIsVolumeJobVolumeProfileIdentityByHrefToMap(model.(*vpcv1.VolumeProfileIdentityByHref))
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

func DataSourceIBMIsVolumeJobVolumeProfileIdentityByNameToMap(model *vpcv1.VolumeProfileIdentityByName) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = *model.Name
	return modelMap, nil
}

func DataSourceIBMIsVolumeJobVolumeProfileIdentityByHrefToMap(model *vpcv1.VolumeProfileIdentityByHref) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["href"] = *model.Href
	return modelMap, nil
}
