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
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func DataSourceIBMIsDynamicRouteServerRouteReportJob() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsDynamicRouteServerRouteReportJobRead,

		Schema: map[string]*schema.Schema{
			"dynamic_route_server_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The dynamic route server identifier.",
			},
			"is_dynamic_route_server_route_report_job_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The route report job identifier.",
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
			"format": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The format used for the route report:`json` - The route report is generated based on the json schema.",
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
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name for this dynamic route server route report export job. The name must not be used by another export job for the image. Changing the name will not affect the exported image name, `storage_object.name`, or `storage_href` values.",
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
							Computed:    true,
							Description: "A link to documentation about this status reason.",
						},
					},
				},
			},
			"storage_bucket": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The Cloud Object Storage bucket of the exported dynamic route server route report object.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
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
							Computed:    true,
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
		},
	}
}

func dataSourceIBMIsDynamicRouteServerRouteReportJobRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_route_report_job", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getDynamicRouteServerRouteReportJobOptions := &vpcv1.GetDynamicRouteServerRouteReportJobOptions{}

	getDynamicRouteServerRouteReportJobOptions.SetDynamicRouteServerID(d.Get("dynamic_route_server_id").(string))
	getDynamicRouteServerRouteReportJobOptions.SetID(d.Get("is_dynamic_route_server_route_report_job_id").(string))

	dynamicRouteServerRouteReportJob, _, err := vpcClient.GetDynamicRouteServerRouteReportJobWithContext(context, getDynamicRouteServerRouteReportJobOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDynamicRouteServerRouteReportJobWithContext failed: %s", err.Error()), "(Data) ibm_is_dynamic_route_server_route_report_job", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *getDynamicRouteServerRouteReportJobOptions.DynamicRouteServerID, *getDynamicRouteServerRouteReportJobOptions.ID))

	if !core.IsNil(dynamicRouteServerRouteReportJob.CompletedAt) {
		if err = d.Set("completed_at", flex.DateTimeToString(dynamicRouteServerRouteReportJob.CompletedAt)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting completed_at: %s", err), "(Data) ibm_is_dynamic_route_server_route_report_job", "read", "set-completed_at").GetDiag()
		}
	}

	if err = d.Set("created_at", flex.DateTimeToString(dynamicRouteServerRouteReportJob.CreatedAt)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_is_dynamic_route_server_route_report_job", "read", "set-created_at").GetDiag()
	}

	if err = d.Set("format", dynamicRouteServerRouteReportJob.Format); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting format: %s", err), "(Data) ibm_is_dynamic_route_server_route_report_job", "read", "set-format").GetDiag()
	}

	if err = d.Set("href", dynamicRouteServerRouteReportJob.Href); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_is_dynamic_route_server_route_report_job", "read", "set-href").GetDiag()
	}

	if !core.IsNil(dynamicRouteServerRouteReportJob.JSONSchema) {
		jsonSchema := []map[string]interface{}{}
		jsonSchemaMap, err := DataSourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerRouteReportSchemaToMap(dynamicRouteServerRouteReportJob.JSONSchema)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_route_report_job", "read", "json_schema-to-map").GetDiag()
		}
		jsonSchema = append(jsonSchema, jsonSchemaMap)
		if err = d.Set("json_schema", jsonSchema); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting json_schema: %s", err), "(Data) ibm_is_dynamic_route_server_route_report_job", "read", "set-json_schema").GetDiag()
		}
	}

	if err = d.Set("name", dynamicRouteServerRouteReportJob.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_is_dynamic_route_server_route_report_job", "read", "set-name").GetDiag()
	}

	if err = d.Set("resource_type", dynamicRouteServerRouteReportJob.ResourceType); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource_type: %s", err), "(Data) ibm_is_dynamic_route_server_route_report_job", "read", "set-resource_type").GetDiag()
	}

	if !core.IsNil(dynamicRouteServerRouteReportJob.StartedAt) {
		if err = d.Set("started_at", flex.DateTimeToString(dynamicRouteServerRouteReportJob.StartedAt)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting started_at: %s", err), "(Data) ibm_is_dynamic_route_server_route_report_job", "read", "set-started_at").GetDiag()
		}
	}

	if err = d.Set("status", dynamicRouteServerRouteReportJob.Status); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_is_dynamic_route_server_route_report_job", "read", "set-status").GetDiag()
	}

	statusReasons := []map[string]interface{}{}
	for _, statusReasonsItem := range dynamicRouteServerRouteReportJob.StatusReasons {
		statusReasonsItemMap, err := DataSourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerRouteReportJobStatusReasonToMap(&statusReasonsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_route_report_job", "read", "status_reasons-to-map").GetDiag()
		}
		statusReasons = append(statusReasons, statusReasonsItemMap)
	}
	if err = d.Set("status_reasons", statusReasons); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status_reasons: %s", err), "(Data) ibm_is_dynamic_route_server_route_report_job", "read", "set-status_reasons").GetDiag()
	}

	storageBucket := []map[string]interface{}{}
	storageBucketMap, err := DataSourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerCloudObjectStorageBucketReferenceToMap(dynamicRouteServerRouteReportJob.StorageBucket)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_route_report_job", "read", "storage_bucket-to-map").GetDiag()
	}
	storageBucket = append(storageBucket, storageBucketMap)
	if err = d.Set("storage_bucket", storageBucket); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting storage_bucket: %s", err), "(Data) ibm_is_dynamic_route_server_route_report_job", "read", "set-storage_bucket").GetDiag()
	}

	if err = d.Set("storage_href", dynamicRouteServerRouteReportJob.StorageHref); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting storage_href: %s", err), "(Data) ibm_is_dynamic_route_server_route_report_job", "read", "set-storage_href").GetDiag()
	}

	storageObject := []map[string]interface{}{}
	storageObjectMap, err := DataSourceIBMIsDynamicRouteServerRouteReportJobCloudObjectStorageObjectReferenceToMap(dynamicRouteServerRouteReportJob.StorageObject)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_route_report_job", "read", "storage_object-to-map").GetDiag()
	}
	storageObject = append(storageObject, storageObjectMap)
	if err = d.Set("storage_object", storageObject); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting storage_object: %s", err), "(Data) ibm_is_dynamic_route_server_route_report_job", "read", "set-storage_object").GetDiag()
	}

	return nil
}

func DataSourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerRouteReportSchemaToMap(model *vpcv1.DynamicRouteServerRouteReportSchema) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["href"] = *model.Href
	modelMap["type"] = *model.Type
	modelMap["version"] = *model.Version
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerRouteReportJobStatusReasonToMap(model *vpcv1.DynamicRouteServerRouteReportJobStatusReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["code"] = *model.Code
	modelMap["message"] = *model.Message
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerRouteReportJobDynamicRouteServerCloudObjectStorageBucketReferenceToMap(model *vpcv1.DynamicRouteServerCloudObjectStorageBucketReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	if model.Deleted != nil {
		deletedMap, err := DataSourceIBMIsDynamicRouteServerRouteReportJobDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	modelMap["name"] = *model.Name
	modelMap["resource_type"] = *model.ResourceType
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerRouteReportJobDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = *model.MoreInfo
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerRouteReportJobCloudObjectStorageObjectReferenceToMap(model *vpcv1.CloudObjectStorageObjectReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = *model.Name
	return modelMap, nil
}
