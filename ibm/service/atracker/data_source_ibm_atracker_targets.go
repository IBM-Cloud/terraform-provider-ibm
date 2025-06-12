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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/platform-services-go-sdk/atrackerv2"
)

func DataSourceIBMAtrackerTargets() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMAtrackerTargetsRead,

		Schema: map[string]*schema.Schema{
			"region": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Limit the query to the specified region.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the target resource.",
			},
			"targets": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of target resources.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The uuid of the target resource.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the target resource.",
						},
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The crn of the target resource.",
						},
						"target_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the target.",
						},
						"region": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Included this optional field if you used it to create a target in a different region other than the one you are connected.",
						},
						"cos_endpoint": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Property values for a Cloud Object Storage Endpoint in responses.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"endpoint": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The host name of the Cloud Object Storage endpoint.",
									},
									"target_crn": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The CRN of the Cloud Object Storage instance.",
									},
									"bucket": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The bucket name under the Cloud Object Storage instance.",
									},
									"api_key": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Sensitive:   true,
										Description: "The IAM API key that has writer access to the Cloud Object Storage instance. This credential is masked in the response. This is required if service_to_service is not enabled.",
									},
									"service_to_service_enabled": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Determines if IBM Cloud Activity Tracker Event Routing has service to service authentication enabled. Set this flag to true if service to service is enabled and do not supply an apikey.",
									},
								},
							},
						},
						"eventstreams_endpoint": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Property values for the Event Streams Endpoint in responses.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target_crn": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The CRN of the Event Streams instance.",
									},
									"brokers": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "List of broker endpoints.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"topic": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The messsage hub topic defined in the Event Streams instance.",
									},
									"api_key": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Sensitive:   true,
										Description: "The user password (api key) for the message hub topic in the Event Streams instance. This is required if service_to_service is not enabled.",
									},
									"service_to_service_enabled": &schema.Schema{
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Determines if IBM Cloud Activity Tracker Event Routing has service to service authentication enabled. Set this flag to true if service to service is enabled and do not supply an apikey.",
									},
								},
							},
						},
						"cloudlogs_endpoint": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Property values for the IBM Cloud Logs endpoint in responses.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target_crn": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The CRN of the IBM Cloud Logs instance.",
									},
								},
							},
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
										Computed:    true,
										Description: "The timestamp of the failure.",
									},
									"reason_for_last_failure": &schema.Schema{
										Type:        schema.TypeString,
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
				},
			},
		},
	}
}

func dataSourceIBMAtrackerTargetsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	atrackerClient, err := meta.(conns.ClientSession).AtrackerV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_atracker_targets", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listTargetsOptions := &atrackerv2.ListTargetsOptions{}

	if _, ok := d.GetOk("region"); ok {
		listTargetsOptions.SetRegion(d.Get("region").(string))
	}

	targetList, _, err := atrackerClient.ListTargetsWithContext(context, listTargetsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListTargetsWithContext failed: %s", err.Error()), "(Data) ibm_atracker_targets", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	// Use the provided filter argument and construct a new list with only the requested resource(s)
	var matchTargets []atrackerv2.Target
	var name string
	var suppliedFilter bool

	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
		suppliedFilter = true
		for _, data := range targetList.Targets {
			if *data.Name == name {
				matchTargets = append(matchTargets, data)
			}
		}
	} else {
		matchTargets = targetList.Targets
	}
	targetList.Targets = matchTargets

	if suppliedFilter {
		if len(targetList.Targets) == 0 {
			return flex.DiscriminatedTerraformErrorf(nil, fmt.Sprintf("no Targets found with name %s", name), "(Data) ibm_atracker_targets", "read", "no-collection-found").GetDiag()
		}
		d.SetId(name)
	} else {
		d.SetId(dataSourceIBMAtrackerTargetsID(d))
	}

	targets := []map[string]interface{}{}
	for _, targetsItem := range targetList.Targets {
		targetsItemMap, err := DataSourceIBMAtrackerTargetsTargetToMap(&targetsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_atracker_targets", "read", "targets-to-map").GetDiag()
		}
		targets = append(targets, targetsItemMap)
	}
	if err = d.Set("targets", targets); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting targets: %s", err), "(Data) ibm_atracker_targets", "read", "set-targets").GetDiag()
	}

	return nil
}

// dataSourceIBMAtrackerTargetsID returns a reasonable ID for the list.
func dataSourceIBMAtrackerTargetsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIBMAtrackerTargetsTargetToMap(model *atrackerv2.Target) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.CRN != nil {
		modelMap["crn"] = *model.CRN
	}
	if model.TargetType != nil {
		modelMap["target_type"] = *model.TargetType
	}
	if model.Region != nil {
		modelMap["region"] = *model.Region
	}
	if model.CosEndpoint != nil {
		cosEndpointMap, err := DataSourceIBMAtrackerTargetsCosEndpointToMap(model.CosEndpoint)
		if err != nil {
			return modelMap, err
		}
		modelMap["cos_endpoint"] = []map[string]interface{}{cosEndpointMap}
	}
	if model.EventstreamsEndpoint != nil {
		eventstreamsEndpointMap, err := DataSourceIBMAtrackerTargetsEventstreamsEndpointToMap(model.EventstreamsEndpoint)
		if err != nil {
			return modelMap, err
		}
		modelMap["eventstreams_endpoint"] = []map[string]interface{}{eventstreamsEndpointMap}
	}
	if model.CloudlogsEndpoint != nil {
		cloudlogsEndpointMap, err := DataSourceIBMAtrackerTargetsCloudLogsEndpointToMap(model.CloudlogsEndpoint)
		if err != nil {
			return modelMap, err
		}
		modelMap["cloudlogs_endpoint"] = []map[string]interface{}{cloudlogsEndpointMap}
	}
	if model.WriteStatus != nil {
		writeStatusMap, err := DataSourceIBMAtrackerTargetsWriteStatusToMap(model.WriteStatus)
		if err != nil {
			return modelMap, err
		}
		modelMap["write_status"] = []map[string]interface{}{writeStatusMap}
	}
	if model.CreatedAt != nil {
		modelMap["created_at"] = model.CreatedAt.String()
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	if model.APIVersion != nil {
		modelMap["api_version"] = *model.APIVersion
	}
	return modelMap, nil
}

func DataSourceIBMAtrackerTargetsCosEndpointToMap(model *atrackerv2.CosEndpoint) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Endpoint != nil {
		modelMap["endpoint"] = *model.Endpoint
	}
	if model.TargetCRN != nil {
		modelMap["target_crn"] = *model.TargetCRN
	}
	if model.Bucket != nil {
		modelMap["bucket"] = *model.Bucket
	}
	if model.ServiceToServiceEnabled != nil {
		modelMap["service_to_service_enabled"] = *model.ServiceToServiceEnabled
	}
	return modelMap, nil
}

func DataSourceIBMAtrackerTargetsEventstreamsEndpointToMap(model *atrackerv2.EventstreamsEndpoint) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.TargetCRN != nil {
		modelMap["target_crn"] = *model.TargetCRN
	}
	if model.Brokers != nil {
		modelMap["brokers"] = model.Brokers
	}
	if model.Topic != nil {
		modelMap["topic"] = *model.Topic
	}
	if model.APIKey != nil {
		modelMap["api_key"] = *model.APIKey // pragma: allowlist secret
	}
	if model.ServiceToServiceEnabled != nil {
		modelMap["service_to_service_enabled"] = *model.ServiceToServiceEnabled
	}
	return modelMap, nil
}

func DataSourceIBMAtrackerTargetsCloudLogsEndpointToMap(model *atrackerv2.CloudLogsEndpoint) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.TargetCRN != nil {
		modelMap["target_crn"] = *model.TargetCRN
	}
	return modelMap, nil
}

func DataSourceIBMAtrackerTargetsWriteStatusToMap(model *atrackerv2.WriteStatus) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Status != nil {
		modelMap["status"] = *model.Status
	}
	if model.LastFailure != nil {
		modelMap["last_failure"] = model.LastFailure.String()
	}
	if model.ReasonForLastFailure != nil {
		modelMap["reason_for_last_failure"] = *model.ReasonForLastFailure
	}
	return modelMap, nil
}
