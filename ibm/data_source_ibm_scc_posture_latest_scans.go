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

func dataSourceIBMSccPostureLatestScans() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMSccPostureLatestScansRead,

		Schema: map[string]*schema.Schema{
			"scan_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the scan.",
			},
			"first": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The URL of the first page of scans.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL of the first page of scans.",
						},
					},
				},
			},
			"last": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The URL of the last page of scans.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL of the last page of scans.",
						},
					},
				},
			},
			"previous": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The URL of the previous page of scans.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL of the previous page of scans.",
						},
					},
				},
			},
			"latest_scans": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The details of a scan.",
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
						"result": &schema.Schema{
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
		},
	}
}

func dataSourceIBMSccPostureLatestScansRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	postureManagementClient, err := meta.(ClientSession).PostureManagementV1()
	if err != nil {
		return diag.FromErr(err)
	}

	listLatestScansOptions := &posturemanagementv1.ListLatestScansOptions{}

	var scansList *posturemanagementv1.ScansList
	var offset int64
	finalList := []posturemanagementv1.ScanItem{}
	var scanID string
	var suppliedFilter bool

	if v, ok := d.GetOk("scan_id"); ok {
		scanID = v.(string)
		suppliedFilter = true
	}

	for {
		listLatestScansOptions.Offset = &offset

		listLatestScansOptions.Limit = core.Int64Ptr(int64(100))
		result, response, err := postureManagementClient.ListLatestScansWithContext(context, listLatestScansOptions)
		scansList = result
		if err != nil {
			log.Printf("[DEBUG] ListLatestScansWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("ListLatestScansWithContext failed %s\n%s", err, response))
		}
		offset = dataSourceScansListGetNext(result.Next)
		if suppliedFilter {
			for _, data := range result.LatestScans {
				if *data.ScanID == scanID {
					finalList = append(finalList, data)
				}
			}
		} else {
			finalList = append(finalList, result.LatestScans...)
		}
		if offset == 0 {
			break
		}
	}

	scansList.LatestScans = finalList

	if suppliedFilter {
		if len(scansList.LatestScans) == 0 {
			return diag.FromErr(fmt.Errorf("no LatestScans found with scanID %s", scanID))
		}
		d.SetId(scanID)
	} else {
		d.SetId(dataSourceIBMSccPostureLatestScansID(d))
	}

	if scansList.First != nil {
		err = d.Set("first", dataSourceScansListFlattenFirst(*scansList.First))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting first %s", err))
		}
	}

	if scansList.Last != nil {
		err = d.Set("last", dataSourceScansListFlattenLast(*scansList.Last))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting last %s", err))
		}
	}

	if scansList.Previous != nil {
		err = d.Set("previous", dataSourceScansListFlattenPrevious(*scansList.Previous))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting previous %s", err))
		}
	}

	if scansList.LatestScans != nil {
		err = d.Set("latest_scans", dataSourceScansListFlattenLatestScans(scansList.LatestScans))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting latest_scans %s", err))
		}
	}

	return nil
}

// dataSourceIBMSccPostureLatestScansID returns a reasonable ID for the list.
func dataSourceIBMSccPostureLatestScansID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func dataSourceScansListFlattenFirst(result posturemanagementv1.ScansListFirst) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceScansListFirstToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceScansListFirstToMap(firstItem posturemanagementv1.ScansListFirst) (firstMap map[string]interface{}) {
	firstMap = map[string]interface{}{}

	if firstItem.Href != nil {
		firstMap["href"] = firstItem.Href
	}

	return firstMap
}

func dataSourceScansListFlattenLast(result posturemanagementv1.ScansListLast) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceScansListLastToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceScansListLastToMap(lastItem posturemanagementv1.ScansListLast) (lastMap map[string]interface{}) {
	lastMap = map[string]interface{}{}

	if lastItem.Href != nil {
		lastMap["href"] = lastItem.Href
	}

	return lastMap
}

func dataSourceScansListFlattenPrevious(result posturemanagementv1.ScansListPrevious) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceScansListPreviousToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceScansListPreviousToMap(previousItem posturemanagementv1.ScansListPrevious) (previousMap map[string]interface{}) {
	previousMap = map[string]interface{}{}

	if previousItem.Href != nil {
		previousMap["href"] = previousItem.Href
	}

	return previousMap
}

func dataSourceScansListFlattenLatestScans(result []posturemanagementv1.ScanItem) (latestScans []map[string]interface{}) {
	for _, latestScansItem := range result {
		latestScans = append(latestScans, dataSourceScansListLatestScansToMap(latestScansItem))
	}

	return latestScans
}

func dataSourceScansListLatestScansToMap(latestScansItem posturemanagementv1.ScanItem) (latestScansMap map[string]interface{}) {
	latestScansMap = map[string]interface{}{}

	if latestScansItem.ScanID != nil {
		latestScansMap["scan_id"] = latestScansItem.ScanID
	}
	if latestScansItem.ScanName != nil {
		latestScansMap["scan_name"] = latestScansItem.ScanName
	}
	if latestScansItem.ScopeID != nil {
		latestScansMap["scope_id"] = latestScansItem.ScopeID
	}
	if latestScansItem.ScopeName != nil {
		latestScansMap["scope_name"] = latestScansItem.ScopeName
	}
	if latestScansItem.ProfileID != nil {
		latestScansMap["profile_id"] = latestScansItem.ProfileID
	}
	if latestScansItem.ProfileName != nil {
		latestScansMap["profile_name"] = latestScansItem.ProfileName
	}
	if latestScansItem.ProfileType != nil {
		latestScansMap["profile_type"] = latestScansItem.ProfileType
	}
	if latestScansItem.GroupProfileID != nil {
		latestScansMap["group_profile_id"] = latestScansItem.GroupProfileID
	}
	if latestScansItem.GroupProfileName != nil {
		latestScansMap["group_profile_name"] = latestScansItem.GroupProfileName
	}
	if latestScansItem.ReportRunBy != nil {
		latestScansMap["report_run_by"] = latestScansItem.ReportRunBy
	}
	if latestScansItem.StartTime != nil {
		latestScansMap["start_time"] = latestScansItem.StartTime.String()
	}
	if latestScansItem.EndTime != nil {
		latestScansMap["end_time"] = latestScansItem.EndTime.String()
	}
	if latestScansItem.Result != nil {
		resultList := []map[string]interface{}{}
		resultMap := dataSourceScansListLatestScansResultToMap(*latestScansItem.Result)
		resultList = append(resultList, resultMap)
		latestScansMap["result"] = resultList
	}

	return latestScansMap
}

func dataSourceScansListLatestScansResultToMap(resultItem posturemanagementv1.ScanResult) (resultMap map[string]interface{}) {
	resultMap = map[string]interface{}{}

	if resultItem.GoalsPassCount != nil {
		resultMap["goals_pass_count"] = resultItem.GoalsPassCount
	}
	if resultItem.GoalsUnableToPerformCount != nil {
		resultMap["goals_unable_to_perform_count"] = resultItem.GoalsUnableToPerformCount
	}
	if resultItem.GoalsNotApplicableCount != nil {
		resultMap["goals_not_applicable_count"] = resultItem.GoalsNotApplicableCount
	}
	if resultItem.GoalsFailCount != nil {
		resultMap["goals_fail_count"] = resultItem.GoalsFailCount
	}
	if resultItem.GoalsTotalCount != nil {
		resultMap["goals_total_count"] = resultItem.GoalsTotalCount
	}
	if resultItem.ControlsPassCount != nil {
		resultMap["controls_pass_count"] = resultItem.ControlsPassCount
	}
	if resultItem.ControlsFailCount != nil {
		resultMap["controls_fail_count"] = resultItem.ControlsFailCount
	}
	if resultItem.ControlsNotApplicableCount != nil {
		resultMap["controls_not_applicable_count"] = resultItem.ControlsNotApplicableCount
	}
	if resultItem.ControlsUnableToPerformCount != nil {
		resultMap["controls_unable_to_perform_count"] = resultItem.ControlsUnableToPerformCount
	}
	if resultItem.ControlsTotalCount != nil {
		resultMap["controls_total_count"] = resultItem.ControlsTotalCount
	}

	return resultMap
}

func dataSourceScansListGetNext(next interface{}) int64 {
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
