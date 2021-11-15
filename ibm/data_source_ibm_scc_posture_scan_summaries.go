// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"reflect"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/scc-go-sdk/posturemanagementv1"
)

func dataSourceIBMSccPostureScanSummaries() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMSccPostureScanSummariesRead,

		Schema: map[string]*schema.Schema{
			"profile_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The profile ID. This can be obtained from the Security and Compliance Center UI by clicking on the profile name. The URL contains the ID.",
			},
			"scope_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The scope ID. This can be obtained from the Security and Compliance Center UI by clicking on the scope name. The URL contains the ID.",
			},
			"scan_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the scan.",
			},
			"first": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "he URL of the first scan summary.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL of the first scan summary.",
						},
					},
				},
			},
			"last": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The URL of the last scan summary.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL of the last scan summary.",
						},
					},
				},
			},
			"previous": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The URL of the previous scan summary.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL of the previous scan summary.",
						},
					},
				},
			},
			"summaries": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Summaries.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"scan_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the scan.",
						},
						"scan_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A system generated name that is the combination of 12 characters in the scope name and 12 characters of a profile name.",
						},
						"scope_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the scan.",
						},
						"scope_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the scope.",
						},
						"report_run_by": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The entity that ran the report.",
						},
						"start_time": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the scan was run.",
						},
						"end_time": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The date and time the scan completed.",
						},
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the collector as it completes a scan.",
						},
						"profile": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The result of a profile.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"profile_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of the profile.",
									},
									"profile_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the profile.",
									},
									"profile_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of profile. To learn more about profile types, check out the [docs] (https://cloud.ibm.com/docs/security-compliance?topic=security-compliance-profiles).",
									},
									"validation_result": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The result of a scan.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"goals_pass_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The number of goals that passed the scan.",
												},
												"goals_unable_to_perform_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The number of goals that could not be validated. A control is listed as 'Unable to perform' when information about its associated resource can't be collected.",
												},
												"goals_not_applicable_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The number of goals that are not relevant to the current scan. A scan is listed as 'Not applicable' when information about its associated resource can't be found.",
												},
												"goals_fail_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The number of goals that failed the scan.",
												},
												"goals_total_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The total number of goals that were included in the scan.",
												},
												"controls_pass_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The number of controls that passed the scan.",
												},
												"controls_fail_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The number of controls that failed the scan.",
												},
												"controls_not_applicable_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The number of controls that are not relevant to the current scan. A scan is listed as 'Not applicable' when information about its associated resource can't be found.",
												},
												"controls_unable_to_perform_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The number of controls that could not be validated. A control is listed as 'Unable to perform' when information about its associated resource can't be collected.",
												},
												"controls_total_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The total number of controls that were included in the scan.",
												},
											},
										},
									},
								},
							},
						},
						"group_profiles": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The result of a group profile.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"group_profile_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The group ID of profile.",
									},
									"group_profile_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The group name of the profile.",
									},
									"profile_type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of profile. To learn more about profile types, check out the [docs] (https://cloud.ibm.com/docs/security-compliance?topic=security-compliance-profiles).",
									},
									"validation_result": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The result of a scan.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"goals_pass_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The number of goals that passed the scan.",
												},
												"goals_unable_to_perform_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The number of goals that could not be validated. A control is listed as 'Unable to perform' when information about its associated resource can't be collected.",
												},
												"goals_not_applicable_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The number of goals that are not relevant to the current scan. A scan is listed as 'Not applicable' when information about its associated resource can't be found.",
												},
												"goals_fail_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The number of goals that failed the scan.",
												},
												"goals_total_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The total number of goals that were included in the scan.",
												},
												"controls_pass_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The number of controls that passed the scan.",
												},
												"controls_fail_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The number of controls that failed the scan.",
												},
												"controls_not_applicable_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The number of controls that are not relevant to the current scan. A scan is listed as 'Not applicable' when information about its associated resource can't be found.",
												},
												"controls_unable_to_perform_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The number of controls that could not be validated. A control is listed as 'Unable to perform' when information about its associated resource can't be collected.",
												},
												"controls_total_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The total number of controls that were included in the scan.",
												},
											},
										},
									},
									"profiles": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The result of a each profile in group profile.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"profile_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The ID of the profile.",
												},
												"profile_name": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name of the profile.",
												},
												"profile_type": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The type of profile. To learn more about profile types, check out the [docs] (https://cloud.ibm.com/docs/security-compliance?topic=security-compliance-profiles).",
												},
												"validation_result": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The result of a scan.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"controls_pass_count": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The number of controls that passed the scan.",
															},
															"controls_fail_count": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The number of controls that failed the scan.",
															},
															"controls_not_applicable_count": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The number of controls that are not relevant to the current scan. A scan is listed as 'Not applicable' when information about its associated resource can't be found.",
															},
															"controls_unable_to_perform_count": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The number of controls that could not be validated. A control is listed as 'Unable to perform' when information about its associated resource can't be collected.",
															},
															"controls_total_count": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The total number of controls that were included in the scan.",
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
				},
			},
		},
	}
}

func dataSourceIBMSccPostureScanSummariesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	postureManagementClient, err := meta.(ClientSession).PostureManagementV1()
	if err != nil {
		return diag.FromErr(err)
	}

	scanSummariesOptions := &posturemanagementv1.ScanSummariesOptions{}

	scanSummariesOptions.SetProfileID(d.Get("profile_id").(string))
	scanSummariesOptions.SetScopeID(d.Get("scope_id").(string))

	var summariesList *posturemanagementv1.SummariesList
	var offset int64
	finalList := []posturemanagementv1.SummaryItem{}
	var scanID string
	var suppliedFilter bool

	if v, ok := d.GetOk("scan_id"); ok {
		scanID = v.(string)
		suppliedFilter = true
	}

	for {
		scanSummariesOptions.Offset = &offset

		scanSummariesOptions.Limit = core.Int64Ptr(int64(100))
		result, response, err := postureManagementClient.ScanSummariesWithContext(context, scanSummariesOptions)
		summariesList = result
		if err != nil {
			log.Printf("[DEBUG] ScanSummariesWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("ScanSummariesWithContext failed %s\n%s", err, response))
		}
		offset = dataSourceSummariesListGetNext(result.Next)
		if suppliedFilter {
			for _, data := range result.Summaries {
				if *data.ScanID == scanID {
					finalList = append(finalList, data)
				}
			}
		} else {
			finalList = append(finalList, result.Summaries...)
		}
		if offset == 0 {
			break
		}
	}

	summariesList.Summaries = finalList

	if suppliedFilter {
		if len(summariesList.Summaries) == 0 {
			return diag.FromErr(fmt.Errorf("no Summaries found with scanID %s", scanID))
		}
		d.SetId(scanID)
	} else {
		d.SetId(dataSourceIBMSccPostureScanSummariesID(d))
	}

	if summariesList.First != nil {
		err = d.Set("first", dataSourceSummariesListFlattenFirst(*summariesList.First))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting first %s", err))
		}
	}

	if summariesList.Last != nil {
		err = d.Set("last", dataSourceSummariesListFlattenLast(*summariesList.Last))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting last %s", err))
		}
	}

	if summariesList.Previous != nil {
		err = d.Set("previous", dataSourceSummariesListFlattenPrevious(*summariesList.Previous))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting previous %s", err))
		}
	}

	if summariesList.Summaries != nil {
		err = d.Set("summaries", dataSourceSummariesListFlattenSummaries(summariesList.Summaries))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting summaries %s", err))
		}
	}

	return nil
}

// dataSourceIBMSccPostureScanSummariesID returns a reasonable ID for the list.
func dataSourceIBMSccPostureScanSummariesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func dataSourceSummariesListFlattenFirst(result posturemanagementv1.SummariesListFirst) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceSummariesListFirstToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceSummariesListFirstToMap(firstItem posturemanagementv1.SummariesListFirst) (firstMap map[string]interface{}) {
	firstMap = map[string]interface{}{}

	if firstItem.Href != nil {
		firstMap["href"] = firstItem.Href
	}

	return firstMap
}

func dataSourceSummariesListFlattenLast(result posturemanagementv1.SummariesListLast) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceSummariesListLastToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceSummariesListLastToMap(lastItem posturemanagementv1.SummariesListLast) (lastMap map[string]interface{}) {
	lastMap = map[string]interface{}{}

	if lastItem.Href != nil {
		lastMap["href"] = lastItem.Href
	}

	return lastMap
}

func dataSourceSummariesListFlattenPrevious(result posturemanagementv1.SummariesListPrevious) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceSummariesListPreviousToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceSummariesListPreviousToMap(previousItem posturemanagementv1.SummariesListPrevious) (previousMap map[string]interface{}) {
	previousMap = map[string]interface{}{}

	if previousItem.Href != nil {
		previousMap["href"] = previousItem.Href
	}

	return previousMap
}

func dataSourceSummariesListFlattenSummaries(result []posturemanagementv1.SummaryItem) (summaries []map[string]interface{}) {
	for _, summariesItem := range result {
		summaries = append(summaries, dataSourceSummariesListSummariesToMap(summariesItem))
	}

	return summaries
}

func dataSourceSummariesListSummariesToMap(summariesItem posturemanagementv1.SummaryItem) (summariesMap map[string]interface{}) {
	summariesMap = map[string]interface{}{}

	if summariesItem.ScanID != nil {
		summariesMap["scan_id"] = summariesItem.ScanID
	}
	if summariesItem.ScanName != nil {
		summariesMap["scan_name"] = summariesItem.ScanName
	}
	if summariesItem.ScopeID != nil {
		summariesMap["scope_id"] = summariesItem.ScopeID
	}
	if summariesItem.ScopeName != nil {
		summariesMap["scope_name"] = summariesItem.ScopeName
	}
	if summariesItem.ReportRunBy != nil {
		summariesMap["report_run_by"] = summariesItem.ReportRunBy
	}
	if summariesItem.StartTime != nil {
		summariesMap["start_time"] = summariesItem.StartTime.String()
	}
	if summariesItem.EndTime != nil {
		summariesMap["end_time"] = summariesItem.EndTime.String()
	}
	if summariesItem.Status != nil {
		summariesMap["status"] = summariesItem.Status
	}
	if summariesItem.Profile != nil {
		profileList := []map[string]interface{}{}
		profileMap := dataSourceSummariesListSummariesProfileToMap(*summariesItem.Profile)
		profileList = append(profileList, profileMap)
		summariesMap["profile"] = profileList
	}
	if summariesItem.GroupProfiles != nil {
		groupProfilesList := []map[string]interface{}{}
		groupProfilesMap := dataSourceSummariesListSummariesGroupProfilesToMap(*summariesItem.GroupProfiles)
		groupProfilesList = append(groupProfilesList, groupProfilesMap)
		summariesMap["group_profiles"] = groupProfilesList
	}

	return summariesMap
}

func dataSourceSummariesListSummariesProfileToMap(profileItem posturemanagementv1.ProfileResult) (profileMap map[string]interface{}) {
	profileMap = map[string]interface{}{}

	if profileItem.ProfileID != nil {
		profileMap["profile_id"] = profileItem.ProfileID
	}
	if profileItem.ProfileName != nil {
		profileMap["profile_name"] = profileItem.ProfileName
	}
	if profileItem.ProfileType != nil {
		profileMap["profile_type"] = profileItem.ProfileType
	}
	if profileItem.ValidationResult != nil {
		validationResultList := []map[string]interface{}{}
		validationResultMap := dataSourceSummariesListProfileValidationResultToMap(*profileItem.ValidationResult)
		validationResultList = append(validationResultList, validationResultMap)
		profileMap["validation_result"] = validationResultList
	}

	return profileMap
}

func dataSourceSummariesListProfileValidationResultToMap(validationResultItem posturemanagementv1.ScanResult) (validationResultMap map[string]interface{}) {
	validationResultMap = map[string]interface{}{}

	if validationResultItem.GoalsPassCount != nil {
		validationResultMap["goals_pass_count"] = validationResultItem.GoalsPassCount
	}
	if validationResultItem.GoalsUnableToPerformCount != nil {
		validationResultMap["goals_unable_to_perform_count"] = validationResultItem.GoalsUnableToPerformCount
	}
	if validationResultItem.GoalsNotApplicableCount != nil {
		validationResultMap["goals_not_applicable_count"] = validationResultItem.GoalsNotApplicableCount
	}
	if validationResultItem.GoalsFailCount != nil {
		validationResultMap["goals_fail_count"] = validationResultItem.GoalsFailCount
	}
	if validationResultItem.GoalsTotalCount != nil {
		validationResultMap["goals_total_count"] = validationResultItem.GoalsTotalCount
	}
	if validationResultItem.ControlsPassCount != nil {
		validationResultMap["controls_pass_count"] = validationResultItem.ControlsPassCount
	}
	if validationResultItem.ControlsFailCount != nil {
		validationResultMap["controls_fail_count"] = validationResultItem.ControlsFailCount
	}
	if validationResultItem.ControlsNotApplicableCount != nil {
		validationResultMap["controls_not_applicable_count"] = validationResultItem.ControlsNotApplicableCount
	}
	if validationResultItem.ControlsUnableToPerformCount != nil {
		validationResultMap["controls_unable_to_perform_count"] = validationResultItem.ControlsUnableToPerformCount
	}
	if validationResultItem.ControlsTotalCount != nil {
		validationResultMap["controls_total_count"] = validationResultItem.ControlsTotalCount
	}

	return validationResultMap
}

func dataSourceSummariesListSummariesGroupProfilesToMap(groupProfilesItem posturemanagementv1.GroupProfileResult) (groupProfilesMap map[string]interface{}) {
	groupProfilesMap = map[string]interface{}{}

	if groupProfilesItem.GroupProfileID != nil {
		groupProfilesMap["group_profile_id"] = groupProfilesItem.GroupProfileID
	}
	if groupProfilesItem.GroupProfileName != nil {
		groupProfilesMap["group_profile_name"] = groupProfilesItem.GroupProfileName
	}
	if groupProfilesItem.ProfileType != nil {
		groupProfilesMap["profile_type"] = groupProfilesItem.ProfileType
	}
	if groupProfilesItem.ValidationResult != nil {
		validationResultList := []map[string]interface{}{}
		validationResultMap := dataSourceSummariesListGroupProfilesValidationResultToMap(*groupProfilesItem.ValidationResult)
		validationResultList = append(validationResultList, validationResultMap)
		groupProfilesMap["validation_result"] = validationResultList
	}
	if groupProfilesItem.Profiles != nil {
		profilesList := []map[string]interface{}{}
		for _, profilesItem := range groupProfilesItem.Profiles {
			profilesList = append(profilesList, dataSourceSummariesListGroupProfilesProfilesToMap(profilesItem))
		}
		groupProfilesMap["profiles"] = profilesList
	}

	return groupProfilesMap
}

func dataSourceSummariesListGroupProfilesValidationResultToMap(validationResultItem posturemanagementv1.ScanResult) (validationResultMap map[string]interface{}) {
	validationResultMap = map[string]interface{}{}

	if validationResultItem.GoalsPassCount != nil {
		validationResultMap["goals_pass_count"] = validationResultItem.GoalsPassCount
	}
	if validationResultItem.GoalsUnableToPerformCount != nil {
		validationResultMap["goals_unable_to_perform_count"] = validationResultItem.GoalsUnableToPerformCount
	}
	if validationResultItem.GoalsNotApplicableCount != nil {
		validationResultMap["goals_not_applicable_count"] = validationResultItem.GoalsNotApplicableCount
	}
	if validationResultItem.GoalsFailCount != nil {
		validationResultMap["goals_fail_count"] = validationResultItem.GoalsFailCount
	}
	if validationResultItem.GoalsTotalCount != nil {
		validationResultMap["goals_total_count"] = validationResultItem.GoalsTotalCount
	}
	if validationResultItem.ControlsPassCount != nil {
		validationResultMap["controls_pass_count"] = validationResultItem.ControlsPassCount
	}
	if validationResultItem.ControlsFailCount != nil {
		validationResultMap["controls_fail_count"] = validationResultItem.ControlsFailCount
	}
	if validationResultItem.ControlsNotApplicableCount != nil {
		validationResultMap["controls_not_applicable_count"] = validationResultItem.ControlsNotApplicableCount
	}
	if validationResultItem.ControlsUnableToPerformCount != nil {
		validationResultMap["controls_unable_to_perform_count"] = validationResultItem.ControlsUnableToPerformCount
	}
	if validationResultItem.ControlsTotalCount != nil {
		validationResultMap["controls_total_count"] = validationResultItem.ControlsTotalCount
	}

	return validationResultMap
}

func dataSourceSummariesListGroupProfilesProfilesToMap(profilesItem posturemanagementv1.ProfilesResult) (profilesMap map[string]interface{}) {
	profilesMap = map[string]interface{}{}

	if profilesItem.ProfileID != nil {
		profilesMap["profile_id"] = profilesItem.ProfileID
	}
	if profilesItem.ProfileName != nil {
		profilesMap["profile_name"] = profilesItem.ProfileName
	}
	if profilesItem.ProfileType != nil {
		profilesMap["profile_type"] = profilesItem.ProfileType
	}
	if profilesItem.ValidationResult != nil {
		validationResultList := []map[string]interface{}{}
		validationResultMap := dataSourceSummariesListProfilesValidationResultToMap(*profilesItem.ValidationResult)
		validationResultList = append(validationResultList, validationResultMap)
		profilesMap["validation_result"] = validationResultList
	}

	return profilesMap
}

func dataSourceSummariesListProfilesValidationResultToMap(validationResultItem posturemanagementv1.Results) (validationResultMap map[string]interface{}) {
	validationResultMap = map[string]interface{}{}

	if validationResultItem.ControlsPassCount != nil {
		validationResultMap["controls_pass_count"] = validationResultItem.ControlsPassCount
	}
	if validationResultItem.ControlsFailCount != nil {
		validationResultMap["controls_fail_count"] = validationResultItem.ControlsFailCount
	}
	if validationResultItem.ControlsNotApplicableCount != nil {
		validationResultMap["controls_not_applicable_count"] = validationResultItem.ControlsNotApplicableCount
	}
	if validationResultItem.ControlsUnableToPerformCount != nil {
		validationResultMap["controls_unable_to_perform_count"] = validationResultItem.ControlsUnableToPerformCount
	}
	if validationResultItem.ControlsTotalCount != nil {
		validationResultMap["controls_total_count"] = validationResultItem.ControlsTotalCount
	}

	return validationResultMap
}

func dataSourceSummariesListGetNext(next interface{}) int64 {
	if reflect.ValueOf(next).IsNil() {
		return 0
	}

	u, err := url.Parse(reflect.ValueOf(next).Elem().FieldByName("Href").Elem().String())
	if err != nil {
		return 0
	}

	q := u.Query()
	var page string

	if q.Get("start") != "" {
		page = q.Get("start")
	} else if q.Get("offset") != "" {
		page = q.Get("offset")
	}

	convertedVal, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		return 0
	}
	return convertedVal
}
