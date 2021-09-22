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

func dataSourceIBMSccPostureProfiles() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMSccPostureProfilesRead,

		Schema: map[string]*schema.Schema{
			"profile_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "An auto-generated unique identifying number of the profile.",
			},
			"first": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The URL of the first page of profiles.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL of the first page of profiles.",
						},
					},
				},
			},
			"last": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The URL of the last page of profiles.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL of the last page of profiles.",
						},
					},
				},
			},
			"previous": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The URL of the previous page of profiles.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL of the previous page of profiles.",
						},
					},
				},
			},
			"profiles": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Profiles.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the profile.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A description of the profile.",
						},
						"version": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The version of the profile.",
						},
						"created_by": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user who created the profile.",
						},
						"modified_by": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The user who last modified the profile.",
						},
						"reason_for_delete": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A reason that you want to delete a profile.",
						},
						"applicability_criteria": &schema.Schema{
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
										Description: "The operatoring system that the profile applies to.",
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
						"profile_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An auto-generated unique identifying number of the profile.",
						},
						"base_profile": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The base profile that the controls are pulled from.",
						},
						"profile_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of profile.",
						},
						"created_time": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time that the profile was created in UTC.",
						},
						"modified_time": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time that the profile was most recently modified in UTC.",
						},
						"enabled": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The profile status. If the profile is enabled, the value is true. If the profile is disabled, the value is false.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMSccPostureProfilesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	postureManagementClient, err := meta.(ClientSession).PostureManagementV1()
	if err != nil {
		return diag.FromErr(err)
	}

	listProfilesOptions := &posturemanagementv1.ListProfilesOptions{}

	var profilesList *posturemanagementv1.ProfilesList
	var offset int64
	finalList := []posturemanagementv1.ProfileItem{}
	var profileID string
	var suppliedFilter bool

	if v, ok := d.GetOk("profile_id"); ok {
		profileID = v.(string)
		suppliedFilter = true
	}

	for {
		listProfilesOptions.Offset = &offset

		listProfilesOptions.Limit = core.Int64Ptr(int64(100))
		result, response, err := postureManagementClient.ListProfilesWithContext(context, listProfilesOptions)
		profilesList = result
		if err != nil {
			log.Printf("[DEBUG] ListProfilesWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("ListProfilesWithContext failed %s\n%s", err, response))
		}
		offset = dataSourceProfilesListGetNext(result.Next)
		if suppliedFilter {
			for _, data := range result.Profiles {
				if *data.ProfileID == profileID {
					finalList = append(finalList, data)
				}
			}
		} else {
			finalList = append(finalList, result.Profiles...)
		}
		if offset == 0 {
			break
		}
	}

	profilesList.Profiles = finalList

	if suppliedFilter {
		if len(profilesList.Profiles) == 0 {
			return diag.FromErr(fmt.Errorf("no Profiles found with profileID %s", profileID))
		}
		d.SetId(profileID)
	} else {
		d.SetId(dataSourceIBMSccPostureProfilesID(d))
	}

	if profilesList.First != nil {
		err = d.Set("first", dataSourceProfilesListFlattenFirst(*profilesList.First))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting first %s", err))
		}
	}

	if profilesList.Last != nil {
		err = d.Set("last", dataSourceProfilesListFlattenLast(*profilesList.Last))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting last %s", err))
		}
	}

	if profilesList.Previous != nil {
		err = d.Set("previous", dataSourceProfilesListFlattenPrevious(*profilesList.Previous))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting previous %s", err))
		}
	}

	if profilesList.Profiles != nil {
		err = d.Set("profiles", dataSourceProfilesListFlattenProfiles(profilesList.Profiles))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting profiles %s", err))
		}
	}

	return nil
}

// dataSourceIBMSccPostureProfilesID returns a reasonable ID for the list.
func dataSourceIBMSccPostureProfilesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func dataSourceProfilesListFlattenFirst(result posturemanagementv1.ProfilesListFirst) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceProfilesListFirstToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceProfilesListFirstToMap(firstItem posturemanagementv1.ProfilesListFirst) (firstMap map[string]interface{}) {
	firstMap = map[string]interface{}{}

	if firstItem.Href != nil {
		firstMap["href"] = firstItem.Href
	}

	return firstMap
}

func dataSourceProfilesListFlattenLast(result posturemanagementv1.ProfilesListLast) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceProfilesListLastToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceProfilesListLastToMap(lastItem posturemanagementv1.ProfilesListLast) (lastMap map[string]interface{}) {
	lastMap = map[string]interface{}{}

	if lastItem.Href != nil {
		lastMap["href"] = lastItem.Href
	}

	return lastMap
}

func dataSourceProfilesListFlattenPrevious(result posturemanagementv1.ProfilesListPrevious) (finalList []map[string]interface{}) {
	finalList = []map[string]interface{}{}
	finalMap := dataSourceProfilesListPreviousToMap(result)
	finalList = append(finalList, finalMap)

	return finalList
}

func dataSourceProfilesListPreviousToMap(previousItem posturemanagementv1.ProfilesListPrevious) (previousMap map[string]interface{}) {
	previousMap = map[string]interface{}{}

	if previousItem.Href != nil {
		previousMap["href"] = previousItem.Href
	}

	return previousMap
}

func dataSourceProfilesListFlattenProfiles(result []posturemanagementv1.ProfileItem) (profiles []map[string]interface{}) {
	for _, profilesItem := range result {
		profiles = append(profiles, dataSourceProfilesListProfilesToMap(profilesItem))
	}

	return profiles
}

func dataSourceProfilesListProfilesToMap(profilesItem posturemanagementv1.ProfileItem) (profilesMap map[string]interface{}) {
	profilesMap = map[string]interface{}{}

	if profilesItem.Name != nil {
		profilesMap["name"] = profilesItem.Name
	}
	if profilesItem.Description != nil {
		profilesMap["description"] = profilesItem.Description
	}
	if profilesItem.Version != nil {
		profilesMap["version"] = profilesItem.Version
	}
	if profilesItem.CreatedBy != nil {
		profilesMap["created_by"] = profilesItem.CreatedBy
	}
	if profilesItem.ModifiedBy != nil {
		profilesMap["modified_by"] = profilesItem.ModifiedBy
	}
	if profilesItem.ReasonForDelete != nil {
		profilesMap["reason_for_delete"] = profilesItem.ReasonForDelete
	}
	if profilesItem.ApplicabilityCriteria != nil {
		applicabilityCriteriaList := []map[string]interface{}{}
		applicabilityCriteriaMap := dataSourceProfilesListProfilesApplicabilityCriteriaToMap(*profilesItem.ApplicabilityCriteria)
		applicabilityCriteriaList = append(applicabilityCriteriaList, applicabilityCriteriaMap)
		profilesMap["applicability_criteria"] = applicabilityCriteriaList
	}
	if profilesItem.ProfileID != nil {
		profilesMap["profile_id"] = profilesItem.ProfileID
	}
	if profilesItem.BaseProfile != nil {
		profilesMap["base_profile"] = profilesItem.BaseProfile
	}
	if profilesItem.ProfileType != nil {
		profilesMap["profile_type"] = profilesItem.ProfileType
	}
	if profilesItem.CreatedTime != nil {
		profilesMap["created_time"] = profilesItem.CreatedTime.String()
	}
	if profilesItem.ModifiedTime != nil {
		profilesMap["modified_time"] = profilesItem.ModifiedTime.String()
	}
	if profilesItem.Enabled != nil {
		profilesMap["enabled"] = profilesItem.Enabled
	}

	return profilesMap
}

func dataSourceProfilesListProfilesApplicabilityCriteriaToMap(applicabilityCriteriaItem posturemanagementv1.ApplicabilityCriteria) (applicabilityCriteriaMap map[string]interface{}) {
	applicabilityCriteriaMap = map[string]interface{}{}

	if applicabilityCriteriaItem.Environment != nil {
		applicabilityCriteriaMap["environment"] = applicabilityCriteriaItem.Environment
	}
	if applicabilityCriteriaItem.Resource != nil {
		applicabilityCriteriaMap["resource"] = applicabilityCriteriaItem.Resource
	}
	if applicabilityCriteriaItem.EnvironmentCategory != nil {
		applicabilityCriteriaMap["environment_category"] = applicabilityCriteriaItem.EnvironmentCategory
	}
	if applicabilityCriteriaItem.ResourceCategory != nil {
		applicabilityCriteriaMap["resource_category"] = applicabilityCriteriaItem.ResourceCategory
	}
	if applicabilityCriteriaItem.ResourceType != nil {
		applicabilityCriteriaMap["resource_type"] = applicabilityCriteriaItem.ResourceType
	}
	if applicabilityCriteriaItem.SoftwareDetails != nil {
		applicabilityCriteriaMap["software_details"] = applicabilityCriteriaItem.SoftwareDetails
	}
	if applicabilityCriteriaItem.OsDetails != nil {
		applicabilityCriteriaMap["os_details"] = applicabilityCriteriaItem.OsDetails
	}
	if applicabilityCriteriaItem.AdditionalDetails != nil {
		applicabilityCriteriaMap["additional_details"] = applicabilityCriteriaItem.AdditionalDetails
	}
	if applicabilityCriteriaItem.EnvironmentCategoryDescription != nil {
		applicabilityCriteriaMap["environment_category_description"] = applicabilityCriteriaItem.EnvironmentCategoryDescription
	}
	if applicabilityCriteriaItem.EnvironmentDescription != nil {
		applicabilityCriteriaMap["environment_description"] = applicabilityCriteriaItem.EnvironmentDescription
	}
	if applicabilityCriteriaItem.ResourceCategoryDescription != nil {
		applicabilityCriteriaMap["resource_category_description"] = applicabilityCriteriaItem.ResourceCategoryDescription
	}
	if applicabilityCriteriaItem.ResourceTypeDescription != nil {
		applicabilityCriteriaMap["resource_type_description"] = applicabilityCriteriaItem.ResourceTypeDescription
	}
	if applicabilityCriteriaItem.ResourceDescription != nil {
		applicabilityCriteriaMap["resource_description"] = applicabilityCriteriaItem.ResourceDescription
	}

	return applicabilityCriteriaMap
}

func dataSourceProfilesListGetNext(next interface{}) int64 {
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
