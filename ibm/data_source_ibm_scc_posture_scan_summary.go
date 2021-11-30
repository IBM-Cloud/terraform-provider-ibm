// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/scc-go-sdk/posturemanagementv1"
)

func dataSourceIBMSccPostureScansSummary() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMSccPostureScansSummaryRead,

		Schema: map[string]*schema.Schema{
			"scan_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Your Scan ID.",
			},
			"profile_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The profile ID. This can be obtained from the Security and Compliance Center UI by clicking on the profile name. The URL contains the ID.",
			},
			"discover_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The scan discovery ID.",
			},
			"profile_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The scan profile name.",
			},
			"scope_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The scan summary scope ID.",
			},
			"controls": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of controls on the scan summary.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"control_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The scan summary control ID.",
						},
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The control status.",
						},
						"external_control_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The external control ID.",
						},
						"control_desciption": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The scan profile name.",
						},
						"goals": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of goals on the control.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"goal_description": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The description of the goal.",
									},
									"goal_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The goal ID.",
									},
									"status": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The goal status.",
									},
									"severity": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The severity of the goal.",
									},
									"completed_time": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The report completed time.",
									},
									"error": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The error on goal validation.",
									},
									"resource_result": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The list of resource results.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"resource_name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The resource name.",
												},
												"resource_types": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The resource type.",
												},
												"resource_status": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The resource control result status.",
												},
												"display_expected_value": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The expected results of a resource.",
												},
												"actual_value": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The actual results of a resource.",
												},
												"results_info": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The results information.",
												},
												"not_applicable_reason": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The reason for goal not applicable for a resource.",
												},
											},
										},
									},
									"goal_applicability_criteria": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The criteria that defines how a profile applies.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"environment": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "A list of environments that a profile can be applied to.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"resource": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "A list of resources that a profile can be used with.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"environment_category": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The type of environment that a profile is able to be applied to.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"resource_category": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The type of resource that a profile is able to be applied to.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"resource_type": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The resource type that the profile applies to.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"software_details": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The software that the profile applies to.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"version": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"os_details": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The operating system that the profile applies to.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
															"version": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"additional_details": &schema.Schema{
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "Any additional details about the profile.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"environment_category_description": &schema.Schema{
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "The type of environment that your scope is targeted to.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"environment_description": &schema.Schema{
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "The environment that your scope is targeted to.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"resource_category_description": &schema.Schema{
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "The type of resource that your scope is targeted to.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"resource_type_description": &schema.Schema{
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "A further classification of the type of resource that your scope is targeted to.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"resource_description": &schema.Schema{
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "The resource that is scanned as part of your scope.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},
						"resource_statistics": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "A scans summary controls.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"resource_pass_count": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The resource count of pass controls.",
									},
									"resource_fail_count": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The resource count of fail controls.",
									},
									"resource_unable_to_perform_count": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The number of resources that were unable to be scanned against a control.",
									},
									"resource_not_applicable_count": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The resource count of not applicable(na) controls.",
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

func dataSourceIBMSccPostureScansSummaryRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	postureManagementClient, err := meta.(ClientSession).PostureManagementV1()
	if err != nil {
		return diag.FromErr(err)
	}

	scansSummaryOptions := &posturemanagementv1.ScansSummaryOptions{}

	scansSummaryOptions.SetScanID(d.Get("scan_id").(string))
	scansSummaryOptions.SetProfileID(d.Get("profile_id").(string))

	summary, response, err := postureManagementClient.ScansSummaryWithContext(context, scansSummaryOptions)
	if err != nil {
		log.Printf("[DEBUG] ScansSummaryWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("ScansSummaryWithContext failed %s\n%s", err, response))
	}

	d.SetId(dataSourceIBMSccPostureScansSummaryID(d))
	if err = d.Set("discover_id", summary.DiscoverID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting discover_id: %s", err))
	}
	if err = d.Set("profile_name", summary.ProfileName); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting profile_name: %s", err))
	}
	if err = d.Set("scope_id", summary.ScopeID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting scope_id: %s", err))
	}

	if summary.Controls != nil {
		err = d.Set("controls", dataSourceSummaryFlattenControls(summary.Controls))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting controls %s", err))
		}
	}

	return nil
}

// dataSourceIBMSccPostureScansSummaryID returns a reasonable ID for the list.
func dataSourceIBMSccPostureScansSummaryID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func dataSourceSummaryFlattenControls(result []posturemanagementv1.Control) (controls []map[string]interface{}) {
	for _, controlsItem := range result {
		controls = append(controls, dataSourceSummaryControlsToMap(controlsItem))
	}

	return controls
}

func dataSourceSummaryControlsToMap(controlsItem posturemanagementv1.Control) (controlsMap map[string]interface{}) {
	controlsMap = map[string]interface{}{}

	if controlsItem.ControlID != nil {
		controlsMap["control_id"] = controlsItem.ControlID
	}
	if controlsItem.Status != nil {
		controlsMap["status"] = controlsItem.Status
	}
	if controlsItem.ExternalControlID != nil {
		controlsMap["external_control_id"] = controlsItem.ExternalControlID
	}
	if controlsItem.ControlDesciption != nil {
		controlsMap["control_desciption"] = controlsItem.ControlDesciption
	}
	if controlsItem.Goals != nil {
		goalsList := []map[string]interface{}{}
		for _, goalsItem := range controlsItem.Goals {
			goalsList = append(goalsList, dataSourceSummaryControlsGoalsToMap(goalsItem))
		}
		controlsMap["goals"] = goalsList
	}
	if controlsItem.ResourceStatistics != nil {
		resourceStatisticsList := []map[string]interface{}{}
		resourceStatisticsMap := dataSourceSummaryControlsResourceStatisticsToMap(*controlsItem.ResourceStatistics)
		resourceStatisticsList = append(resourceStatisticsList, resourceStatisticsMap)
		controlsMap["resource_statistics"] = resourceStatisticsList
	}

	return controlsMap
}

func dataSourceSummaryControlsGoalsToMap(goalsItem posturemanagementv1.Goal) (goalsMap map[string]interface{}) {
	goalsMap = map[string]interface{}{}

	if goalsItem.GoalDescription != nil {
		goalsMap["goal_description"] = goalsItem.GoalDescription
	}
	if goalsItem.GoalID != nil {
		goalsMap["goal_id"] = goalsItem.GoalID
	}
	if goalsItem.Status != nil {
		goalsMap["status"] = goalsItem.Status
	}
	if goalsItem.Severity != nil {
		goalsMap["severity"] = goalsItem.Severity
	}
	if goalsItem.CompletedTime != nil {
		goalsMap["completed_time"] = goalsItem.CompletedTime.String()
	}
	if goalsItem.Error != nil {
		goalsMap["error"] = goalsItem.Error
	}
	if goalsItem.ResourceResult != nil {
		resourceResultList := []map[string]interface{}{}
		for _, resourceResultItem := range goalsItem.ResourceResult {
			resourceResultList = append(resourceResultList, dataSourceSummaryGoalsResourceResultToMap(resourceResultItem))
		}
		goalsMap["resource_result"] = resourceResultList
	}
	if goalsItem.GoalApplicabilityCriteria != nil {
		goalApplicabilityCriteriaList := []map[string]interface{}{}
		goalApplicabilityCriteriaMap := dataSourceSummaryGoalsGoalApplicabilityCriteriaToMap(*goalsItem.GoalApplicabilityCriteria)
		goalApplicabilityCriteriaList = append(goalApplicabilityCriteriaList, goalApplicabilityCriteriaMap)
		goalsMap["goal_applicability_criteria"] = goalApplicabilityCriteriaList
	}

	return goalsMap
}

func dataSourceSummaryGoalsResourceResultToMap(resourceResultItem posturemanagementv1.ResourceResult) (resourceResultMap map[string]interface{}) {
	resourceResultMap = map[string]interface{}{}

	if resourceResultItem.ResourceName != nil {
		resourceResultMap["resource_name"] = resourceResultItem.ResourceName
	}
	if resourceResultItem.ResourceTypes != nil {
		resourceResultMap["resource_types"] = resourceResultItem.ResourceTypes
	}
	if resourceResultItem.ResourceStatus != nil {
		resourceResultMap["resource_status"] = resourceResultItem.ResourceStatus
	}
	if resourceResultItem.DisplayExpectedValue != nil {
		resourceResultMap["display_expected_value"] = resourceResultItem.DisplayExpectedValue
	}
	if resourceResultItem.ActualValue != nil {
		resourceResultMap["actual_value"] = resourceResultItem.ActualValue
	}
	if resourceResultItem.ResultsInfo != nil {
		resourceResultMap["results_info"] = resourceResultItem.ResultsInfo
	}
	if resourceResultItem.NotApplicableReason != nil {
		resourceResultMap["not_applicable_reason"] = resourceResultItem.NotApplicableReason
	}

	return resourceResultMap
}

func dataSourceSummaryGoalsGoalApplicabilityCriteriaToMap(goalApplicabilityCriteriaItem posturemanagementv1.GoalApplicabilityCriteria) (goalApplicabilityCriteriaMap map[string]interface{}) {
	goalApplicabilityCriteriaMap = map[string]interface{}{}

	if goalApplicabilityCriteriaItem.Environment != nil {
		goalApplicabilityCriteriaMap["environment"] = goalApplicabilityCriteriaItem.Environment
	}
	if goalApplicabilityCriteriaItem.Resource != nil {
		goalApplicabilityCriteriaMap["resource"] = goalApplicabilityCriteriaItem.Resource
	}
	if goalApplicabilityCriteriaItem.EnvironmentCategory != nil {
		goalApplicabilityCriteriaMap["environment_category"] = goalApplicabilityCriteriaItem.EnvironmentCategory
	}
	if goalApplicabilityCriteriaItem.ResourceCategory != nil {
		goalApplicabilityCriteriaMap["resource_category"] = goalApplicabilityCriteriaItem.ResourceCategory
	}
	if goalApplicabilityCriteriaItem.ResourceType != nil {
		goalApplicabilityCriteriaMap["resource_type"] = goalApplicabilityCriteriaItem.ResourceType
	}
	if goalApplicabilityCriteriaItem.SoftwareDetails != nil {
		goalApplicabilityCriteriaMap["software_details"] = goalApplicabilityCriteriaItem.SoftwareDetails
	}
	if goalApplicabilityCriteriaItem.OsDetails != nil {
		goalApplicabilityCriteriaMap["os_details"] = goalApplicabilityCriteriaItem.OsDetails
	}
	if goalApplicabilityCriteriaItem.AdditionalDetails != nil {
		goalApplicabilityCriteriaMap["additional_details"] = goalApplicabilityCriteriaItem.AdditionalDetails
	}
	if goalApplicabilityCriteriaItem.EnvironmentCategoryDescription != nil {
		goalApplicabilityCriteriaMap["environment_category_description"] = goalApplicabilityCriteriaItem.EnvironmentCategoryDescription
	}
	if goalApplicabilityCriteriaItem.EnvironmentDescription != nil {
		goalApplicabilityCriteriaMap["environment_description"] = goalApplicabilityCriteriaItem.EnvironmentDescription
	}
	if goalApplicabilityCriteriaItem.ResourceCategoryDescription != nil {
		goalApplicabilityCriteriaMap["resource_category_description"] = goalApplicabilityCriteriaItem.ResourceCategoryDescription
	}
	if goalApplicabilityCriteriaItem.ResourceTypeDescription != nil {
		goalApplicabilityCriteriaMap["resource_type_description"] = goalApplicabilityCriteriaItem.ResourceTypeDescription
	}
	if goalApplicabilityCriteriaItem.ResourceDescription != nil {
		goalApplicabilityCriteriaMap["resource_description"] = goalApplicabilityCriteriaItem.ResourceDescription
	}

	return goalApplicabilityCriteriaMap
}

func dataSourceSummaryControlsResourceStatisticsToMap(resourceStatisticsItem posturemanagementv1.ResourceStatistics) (resourceStatisticsMap map[string]interface{}) {
	resourceStatisticsMap = map[string]interface{}{}

	if resourceStatisticsItem.ResourcePassCount != nil {
		resourceStatisticsMap["resource_pass_count"] = resourceStatisticsItem.ResourcePassCount
	}
	if resourceStatisticsItem.ResourceFailCount != nil {
		resourceStatisticsMap["resource_fail_count"] = resourceStatisticsItem.ResourceFailCount
	}
	if resourceStatisticsItem.ResourceUnableToPerformCount != nil {
		resourceStatisticsMap["resource_unable_to_perform_count"] = resourceStatisticsItem.ResourceUnableToPerformCount
	}
	if resourceStatisticsItem.ResourceNotApplicableCount != nil {
		resourceStatisticsMap["resource_not_applicable_count"] = resourceStatisticsItem.ResourceNotApplicableCount
	}

	return resourceStatisticsMap
}
