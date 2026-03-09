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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func DataSourceIBMIsVolumeJobs() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsVolumeJobsRead,

		Schema: map[string]*schema.Schema{
			"volume_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The volume identifier.",
			},
			"jobs": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The jobs for this volume.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this volume job.",
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
				},
			},
		},
	}
}

func dataSourceIBMIsVolumeJobsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_volume_jobs", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listVolumeJobsOptions := &vpcv1.ListVolumeJobsOptions{}

	listVolumeJobsOptions.SetVolumeID(d.Get("volume_id").(string))

	volumeJobPaginatedCollection, _, err := vpcClient.ListVolumeJobsWithContext(context, listVolumeJobsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListVolumeJobsWithContext failed: %s", err.Error()), "(Data) ibm_is_volume_jobs", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIBMIsVolumeJobsID(d))

	jobs := []map[string]interface{}{}
	for _, jobsItem := range volumeJobPaginatedCollection.Jobs {
		jobsItemMap, err := DataSourceIBMIsVolumeJobsVolumeJobToMap(jobsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_volume_jobs", "read", "jobs-to-map").GetDiag()
		}
		jobs = append(jobs, jobsItemMap)
	}
	if err = d.Set("jobs", jobs); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting jobs: %s", err), "(Data) ibm_is_volume_jobs", "read", "set-jobs").GetDiag()
	}

	return nil
}

// dataSourceIBMIsVolumeJobsID returns a reasonable ID for the list.
func dataSourceIBMIsVolumeJobsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIBMIsVolumeJobsVolumeJobToMap(model vpcv1.VolumeJobIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.VolumeJobTypeMigrate); ok {
		return DataSourceIBMIsVolumeJobsVolumeJobTypeMigrateToMap(model.(*vpcv1.VolumeJobTypeMigrate))
	} else if _, ok := model.(*vpcv1.VolumeJob); ok {
		modelMap := make(map[string]interface{})
		model := model.(*vpcv1.VolumeJob)
		modelMap["auto_delete"] = *model.AutoDelete
		if model.CompletedAt != nil {
			modelMap["completed_at"] = model.CompletedAt.String()
		}
		modelMap["created_at"] = model.CreatedAt.String()
		if model.EstimatedCompletionAt != nil {
			modelMap["estimated_completion_at"] = model.EstimatedCompletionAt.String()
		}
		modelMap["href"] = *model.Href
		modelMap["id"] = *model.ID
		modelMap["job_type"] = *model.JobType
		modelMap["name"] = *model.Name
		modelMap["resource_type"] = *model.ResourceType
		if model.StartedAt != nil {
			modelMap["started_at"] = model.StartedAt.String()
		}
		modelMap["status"] = *model.Status
		statusReasons := []map[string]interface{}{}
		for _, statusReasonsItem := range model.StatusReasons {
			statusReasonsItemMap, err := DataSourceIBMIsVolumeJobsVolumeJobStatusReasonToMap(&statusReasonsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			statusReasons = append(statusReasons, statusReasonsItemMap)
		}
		modelMap["status_reasons"] = statusReasons
		if model.Parameters != nil {
			parametersMap, err := DataSourceIBMIsVolumeJobsVolumeJobTypeMigrateParametersToMap(model.Parameters)
			if err != nil {
				return modelMap, err
			}
			modelMap["parameters"] = []map[string]interface{}{parametersMap}
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized vpcv1.VolumeJobIntf subtype encountered")
	}
}

func DataSourceIBMIsVolumeJobsVolumeJobStatusReasonToMap(model *vpcv1.VolumeJobStatusReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["code"] = *model.Code
	modelMap["message"] = *model.Message
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func DataSourceIBMIsVolumeJobsVolumeJobTypeMigrateParametersToMap(model *vpcv1.VolumeJobTypeMigrateParameters) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Bandwidth != nil {
		modelMap["bandwidth"] = flex.IntValue(model.Bandwidth)
	}
	if model.Iops != nil {
		modelMap["iops"] = flex.IntValue(model.Iops)
	}
	profileMap, err := DataSourceIBMIsVolumeJobsVolumeProfileIdentityToMap(model.Profile)
	if err != nil {
		return modelMap, err
	}
	modelMap["profile"] = []map[string]interface{}{profileMap}
	return modelMap, nil
}

func DataSourceIBMIsVolumeJobsVolumeProfileIdentityToMap(model vpcv1.VolumeProfileIdentityIntf) (map[string]interface{}, error) {
	if _, ok := model.(*vpcv1.VolumeProfileIdentityByName); ok {
		return DataSourceIBMIsVolumeJobsVolumeProfileIdentityByNameToMap(model.(*vpcv1.VolumeProfileIdentityByName))
	} else if _, ok := model.(*vpcv1.VolumeProfileIdentityByHref); ok {
		return DataSourceIBMIsVolumeJobsVolumeProfileIdentityByHrefToMap(model.(*vpcv1.VolumeProfileIdentityByHref))
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

func DataSourceIBMIsVolumeJobsVolumeProfileIdentityByNameToMap(model *vpcv1.VolumeProfileIdentityByName) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = *model.Name
	return modelMap, nil
}

func DataSourceIBMIsVolumeJobsVolumeProfileIdentityByHrefToMap(model *vpcv1.VolumeProfileIdentityByHref) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["href"] = *model.Href
	return modelMap, nil
}

func DataSourceIBMIsVolumeJobsVolumeJobTypeMigrateToMap(model *vpcv1.VolumeJobTypeMigrate) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["auto_delete"] = *model.AutoDelete
	if model.CompletedAt != nil {
		modelMap["completed_at"] = model.CompletedAt.String()
	}
	modelMap["created_at"] = model.CreatedAt.String()
	if model.EstimatedCompletionAt != nil {
		modelMap["estimated_completion_at"] = model.EstimatedCompletionAt.String()
	}
	modelMap["href"] = *model.Href
	modelMap["id"] = *model.ID
	modelMap["job_type"] = *model.JobType
	modelMap["name"] = *model.Name
	modelMap["resource_type"] = *model.ResourceType
	if model.StartedAt != nil {
		modelMap["started_at"] = model.StartedAt.String()
	}
	modelMap["status"] = *model.Status
	statusReasons := []map[string]interface{}{}
	for _, statusReasonsItem := range model.StatusReasons {
		statusReasonsItemMap, err := DataSourceIBMIsVolumeJobsVolumeJobStatusReasonToMap(&statusReasonsItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		statusReasons = append(statusReasons, statusReasonsItemMap)
	}
	modelMap["status_reasons"] = statusReasons
	parametersMap, err := DataSourceIBMIsVolumeJobsVolumeJobTypeMigrateParametersToMap(model.Parameters)
	if err != nil {
		return modelMap, err
	}
	modelMap["parameters"] = []map[string]interface{}{parametersMap}
	return modelMap, nil
}
