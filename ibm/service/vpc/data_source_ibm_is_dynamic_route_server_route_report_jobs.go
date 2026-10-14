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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func DataSourceIBMIsDynamicRouteServerRouteReportJobs() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMIsDynamicRouteServerRouteReportJobsRead,

		Schema: map[string]*schema.Schema{
			"dynamic_route_server_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The dynamic route server identifier.",
			},
			"sort": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "-created_at",
				Description: "Sorts the returned collection by the specified property name in ascending order. A `-` may be prepended to the name to sort in descending order. For example, the value `-created_at` sorts the collection by the `created_at` property in descending order, and the value `name` sorts it by the `name` property in ascending order.",
			},
			"route_report_jobs": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique identifier for this dynamic route server route report export job.",
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
				},
			},
		},
	}
}

func dataSourceIBMIsDynamicRouteServerRouteReportJobsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := meta.(conns.ClientSession).VpcV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_route_report_jobs", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listDynamicRouteServerRouteReportJobsOptions := &vpcv1.ListDynamicRouteServerRouteReportJobsOptions{}

	listDynamicRouteServerRouteReportJobsOptions.SetDynamicRouteServerID(d.Get("dynamic_route_server_id").(string))
	if _, ok := d.GetOk("sort"); ok {
		listDynamicRouteServerRouteReportJobsOptions.SetSort(d.Get("sort").(string))
	}

	var pager *vpcv1.DynamicRouteServerRouteReportJobsPager
	pager, err = vpcClient.NewDynamicRouteServerRouteReportJobsPager(listDynamicRouteServerRouteReportJobsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_route_report_jobs", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	allItems, err := pager.GetAll()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DynamicRouteServerRouteReportJobsPager.GetAll() failed %s", err), "(Data) ibm_is_dynamic_route_server_route_report_jobs", "read")
		log.Printf("[DEBUG] %s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIBMIsDynamicRouteServerRouteReportJobsID(d))

	mapSlice := []map[string]interface{}{}
	for _, modelItem := range allItems {
		modelMap, err := DataSourceIBMIsDynamicRouteServerRouteReportJobsDynamicRouteServerRouteReportJobToMap(&modelItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_is_dynamic_route_server_route_report_jobs", "read", "DynamicRouteServers-to-map").GetDiag()
		}
		mapSlice = append(mapSlice, modelMap)
	}

	if err = d.Set("route_report_jobs", mapSlice); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting route_report_jobs %s", err), "(Data) ibm_is_dynamic_route_server_route_report_jobs", "read", "route_report_jobs-set").GetDiag()
	}

	return nil
}

// dataSourceIBMIsDynamicRouteServerRouteReportJobsID returns a reasonable ID for the list.
func dataSourceIBMIsDynamicRouteServerRouteReportJobsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIBMIsDynamicRouteServerRouteReportJobsDynamicRouteServerRouteReportJobToMap(model *vpcv1.DynamicRouteServerRouteReportJob) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CompletedAt != nil {
		modelMap["completed_at"] = model.CompletedAt.String()
	}
	modelMap["created_at"] = model.CreatedAt.String()
	modelMap["format"] = *model.Format
	modelMap["href"] = *model.Href
	modelMap["id"] = *model.ID
	if model.JSONSchema != nil {
		jsonSchemaMap, err := DataSourceIBMIsDynamicRouteServerRouteReportJobsDynamicRouteServerRouteReportSchemaToMap(model.JSONSchema)
		if err != nil {
			return modelMap, err
		}
		modelMap["json_schema"] = []map[string]interface{}{jsonSchemaMap}
	}
	modelMap["name"] = *model.Name
	modelMap["resource_type"] = *model.ResourceType
	if model.StartedAt != nil {
		modelMap["started_at"] = model.StartedAt.String()
	}
	modelMap["status"] = *model.Status
	statusReasons := []map[string]interface{}{}
	for _, statusReasonsItem := range model.StatusReasons {
		statusReasonsItemMap, err := DataSourceIBMIsDynamicRouteServerRouteReportJobsDynamicRouteServerRouteReportJobStatusReasonToMap(&statusReasonsItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		statusReasons = append(statusReasons, statusReasonsItemMap)
	}
	modelMap["status_reasons"] = statusReasons
	storageBucketMap, err := DataSourceIBMIsDynamicRouteServerRouteReportJobsDynamicRouteServerCloudObjectStorageBucketReferenceToMap(model.StorageBucket)
	if err != nil {
		return modelMap, err
	}
	modelMap["storage_bucket"] = []map[string]interface{}{storageBucketMap}
	modelMap["storage_href"] = *model.StorageHref
	storageObjectMap, err := DataSourceIBMIsDynamicRouteServerRouteReportJobsCloudObjectStorageObjectReferenceToMap(model.StorageObject)
	if err != nil {
		return modelMap, err
	}
	modelMap["storage_object"] = []map[string]interface{}{storageObjectMap}
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerRouteReportJobsDynamicRouteServerRouteReportSchemaToMap(model *vpcv1.DynamicRouteServerRouteReportSchema) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["href"] = *model.Href
	modelMap["type"] = *model.Type
	modelMap["version"] = *model.Version
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerRouteReportJobsDynamicRouteServerRouteReportJobStatusReasonToMap(model *vpcv1.DynamicRouteServerRouteReportJobStatusReason) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["code"] = *model.Code
	modelMap["message"] = *model.Message
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerRouteReportJobsDynamicRouteServerCloudObjectStorageBucketReferenceToMap(model *vpcv1.DynamicRouteServerCloudObjectStorageBucketReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["crn"] = *model.CRN
	if model.Deleted != nil {
		deletedMap, err := DataSourceIBMIsDynamicRouteServerRouteReportJobsDeletedToMap(model.Deleted)
		if err != nil {
			return modelMap, err
		}
		modelMap["deleted"] = []map[string]interface{}{deletedMap}
	}
	modelMap["name"] = *model.Name
	modelMap["resource_type"] = *model.ResourceType
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerRouteReportJobsDeletedToMap(model *vpcv1.Deleted) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["more_info"] = *model.MoreInfo
	return modelMap, nil
}

func DataSourceIBMIsDynamicRouteServerRouteReportJobsCloudObjectStorageObjectReferenceToMap(model *vpcv1.CloudObjectStorageObjectReference) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = *model.Name
	return modelMap, nil
}
