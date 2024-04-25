// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iampolicy

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
)

func DataSourceIBMIAMPolicyAssignments() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPolicyAssignmentRead,

		Schema: map[string]*schema.Schema{
			"version": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "specify version of response body format.",
			},
			"accept_language": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "default",
				Description: "Language code for translations* `default` - English* `de` -  German (Standard)* `en` - English* `es` - Spanish (Spain)* `fr` - French (Standard)* `it` - Italian (Standard)* `ja` - Japanese* `ko` - Korean* `pt-br` - Portuguese (Brazil)* `zh-cn` - Chinese (Simplified, PRC)* `zh-tw` - (Chinese, Taiwan).",
			},
			"template_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Optional template id.",
			},
			"template_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Optional policy template version.",
			},
			"assignments": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of policy assignments.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"target": {
							Type:        schema.TypeMap,
							Required:    true,
							Description: "assignment target details",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"options": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The set of properties required for a policy assignment.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"root": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"requester_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"assignment_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Passed in value to correlate with other assignments.",
												},
												"template": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The template id where this policy is being assigned from.",
															},
															"version": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The template version where this policy is being assigned from.",
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
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Policy assignment ID.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account GUID that the policies assignments belong to..",
						},
						"href": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The href URL that links to the policies assignments API by policy assignment ID.",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The UTC timestamp when the policy assignment was created.",
						},
						"created_by_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The iam ID of the entity that created the policy assignment.",
						},
						"last_modified_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The UTC timestamp when the policy assignment was last modified.",
						},
						"last_modified_by_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The iam ID of the entity that last modified the policy assignment.",
						},
						"resources": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Object for each account assigned.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target": {
										Type:        schema.TypeMap,
										Required:    true,
										Description: "assignment target details",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"policy": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Set of properties for the assigned resource.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"resource_created": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "On success, includes the  policy assigned.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "policy id.",
															},
														},
													},
												},
												"status": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "policy status.",
												},
												"error_message": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The error response from API.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"trace": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The unique transaction id for the request.",
															},
															"errors": {
																Type:        schema.TypeList,
																Computed:    true,
																Description: "The errors encountered during the response.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"code": {
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The API error code for the error.",
																		},
																		"message": {
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The error message returned by the API.",
																		},
																		"details": {
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Additional error details.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"conflicts_with": {
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Details of conflicting resource.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"etag": {
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "The revision number of the resource.",
																								},
																								"role": {
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "The conflicting role id.",
																								},
																								"policy": {
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "The conflicting policy id.",
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																		"more_info": {
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Additional info for error.",
																		},
																	},
																},
															},
															"status_code": {
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The http error code of the response.",
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
						"subject": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "subject details of access type assignment.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"template": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "policy template details.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "policy template id.",
									},
									"version": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "policy template version.",
									},
								},
							},
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The policy assignment status.",
						},
						"template_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "policy template id.",
						},
						"template_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "policy template version.",
						},
						"assignment_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Passed in value to correlate with other assignments.",
						},
						"target_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Assignment target type.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMPolicyAssignmentRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamPolicyManagementClient, err := meta.(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_policy_assignment", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	userDetails, err := meta.(conns.ClientSession).BluemixUserDetails()
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to fetch BluemixUserDetails %s", err))
	}
	accountID := userDetails.UserAccount

	listPolicyAssignmentsOptions := &iampolicymanagementv1.ListPolicyAssignmentsOptions{}
	listPolicyAssignmentsOptions.SetAccountID(accountID)

	listPolicyAssignmentsOptions.SetVersion(d.Get("version").(string))
	if _, ok := d.GetOk("accept_language"); ok {
		listPolicyAssignmentsOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}
	if _, ok := d.GetOk("template_id"); ok {
		listPolicyAssignmentsOptions.SetTemplateID(d.Get("template_id").(string))
	}
	if _, ok := d.GetOk("template_version"); ok {
		listPolicyAssignmentsOptions.SetTemplateVersion(d.Get("template_version").(string))
	}

	policyTemplateAssignmentCollection, _, err := iamPolicyManagementClient.ListPolicyAssignmentsWithContext(context, listPolicyAssignmentsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListPolicyAssignmentsWithContext failed: %s", err.Error()), "(Data) ibm_policy_assignment", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIBMPolicyAssignmentID(d))

	assignments := []map[string]interface{}{}
	if policyTemplateAssignmentCollection.Assignments != nil {
		for _, modelItem := range policyTemplateAssignmentCollection.Assignments {
			modelMap, err := DataSourceIBMPolicyAssignmentPolicyTemplateAssignmentItemsToMap(modelItem)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_policy_assignment", "read")
				return tfErr.GetDiag()
			}
			assignments = append(assignments, modelMap)
		}
	}
	if err = d.Set("assignments", assignments); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting assignments: %s", err), "(Data) ibm_policy_assignment", "read")
		return tfErr.GetDiag()
	}

	return nil
}

func ResourceIBMPolicyAssignmentAssignmentTargetDetailsToMap(model *iampolicymanagementv1.AssignmentTargetDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	return modelMap, nil
}

func ResourceIBMPolicyAssignmentResourceTargetDetailsToMap(model *iampolicymanagementv1.AssignmentTemplateDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Version != nil {
		modelMap["version"] = *model.Version
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	return modelMap, nil
}

// dataSourceIBMPolicyAssignmentID returns a reasonable ID for the list.
func dataSourceIBMPolicyAssignmentID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIBMPolicyAssignmentPolicyTemplateAssignmentItemsToMap(model iampolicymanagementv1.PolicyTemplateAssignmentItemsIntf) (map[string]interface{}, error) {
	if _, ok := model.(*iampolicymanagementv1.PolicyTemplateAssignmentItemsPolicyAssignmentV1); ok {
		return DataSourceIBMPolicyAssignmentPolicyTemplateAssignmentItemsPolicyAssignmentV1ToMap(model.(*iampolicymanagementv1.PolicyTemplateAssignmentItemsPolicyAssignmentV1))
	} else if _, ok := model.(*iampolicymanagementv1.PolicyTemplateAssignmentItemsPolicyAssignment); ok {
		return DataSourceIBMPolicyAssignmentPolicyTemplateAssignmentItemsPolicyAssignmentToMap(model.(*iampolicymanagementv1.PolicyTemplateAssignmentItemsPolicyAssignment))
	} else if _, ok := model.(*iampolicymanagementv1.PolicyTemplateAssignmentItems); ok {
		modelMap := make(map[string]interface{})
		model := model.(*iampolicymanagementv1.PolicyTemplateAssignmentItems)
		if model.Target != nil {
			targetMap, err := DataSourceIBMPolicyAssignmentAssignmentTargetDetailsToMap(model.Target)
			if err != nil {
				return modelMap, err
			}
			// modelMap["target"] = []map[string]interface{}{targetMap}
			modelMap["target"] = targetMap
		}
		if model.Options != nil {
			optionsMap, err := DataSourceIBMPolicyAssignmentPolicyAssignmentV1OptionsToMap(model.Options)
			if err != nil {
				return modelMap, err
			}
			modelMap["options"] = []map[string]interface{}{optionsMap}
		}
		if model.ID != nil {
			modelMap["id"] = *model.ID
		}
		if model.AccountID != nil {
			modelMap["account_id"] = *model.AccountID
		}
		if model.Href != nil {
			modelMap["href"] = *model.Href
		}
		if model.CreatedAt != nil {
			modelMap["created_at"] = model.CreatedAt.String()
		}
		if model.CreatedByID != nil {
			modelMap["created_by_id"] = *model.CreatedByID
		}
		if model.LastModifiedAt != nil {
			modelMap["last_modified_at"] = model.LastModifiedAt.String()
		}
		if model.LastModifiedByID != nil {
			modelMap["last_modified_by_id"] = *model.LastModifiedByID
		}
		if model.Resources != nil {
			resources := []map[string]interface{}{}
			for _, resourcesItem := range model.Resources {
				resourcesItemMap, err := DataSourceIBMPolicyAssignmentPolicyAssignmentV1ResourcesToMap(&resourcesItem)
				if err != nil {
					return modelMap, err
				}
				resources = append(resources, resourcesItemMap)
			}
			modelMap["resources"] = resources
		}
		if model.Subject != nil {
			subjectMap, err := DataSourceIBMPolicyAssignmentPolicyAssignmentV1SubjectToMap(model.Subject)
			if err != nil {
				return modelMap, err
			}
			modelMap["subject"] = []map[string]interface{}{subjectMap}
		}
		if model.Template != nil {
			templateMap, err := DataSourceIBMPolicyAssignmentAssignmentTemplateDetailsToMap(model.Template)
			if err != nil {
				return modelMap, err
			}
			modelMap["template"] = []map[string]interface{}{templateMap}
		}
		if model.Status != nil {
			modelMap["status"] = *model.Status
		}
		if model.TemplateID != nil {
			modelMap["template_id"] = *model.TemplateID
		}
		if model.TemplateVersion != nil {
			modelMap["template_version"] = *model.TemplateVersion
		}
		if model.AssignmentID != nil {
			modelMap["assignment_id"] = *model.AssignmentID
		}
		if model.TargetType != nil {
			modelMap["target_type"] = *model.TargetType
		}
		return modelMap, nil
	} else {
		return nil, fmt.Errorf("Unrecognized iampolicymanagementv1.PolicyTemplateAssignmentItemsIntf subtype encountered")
	}
}

func DataSourceIBMPolicyAssignmentAssignmentTargetDetailsToMap(model *iampolicymanagementv1.AssignmentTargetDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	return modelMap, nil
}

func DataSourceIBMPolicyAssignmentPolicyAssignmentV1OptionsToMap(model *iampolicymanagementv1.PolicyAssignmentV1Options) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	rootMap, err := DataSourceIBMPolicyAssignmentPolicyAssignmentV1OptionsRootToMap(model.Root)
	if err != nil {
		return modelMap, err
	}
	modelMap["root"] = []map[string]interface{}{rootMap}
	return modelMap, nil
}

func DataSourceIBMPolicyAssignmentPolicyAssignmentV1OptionsRootToMap(model *iampolicymanagementv1.PolicyAssignmentV1OptionsRoot) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.RequesterID != nil {
		modelMap["requester_id"] = *model.RequesterID
	}
	if model.AssignmentID != nil {
		modelMap["assignment_id"] = *model.AssignmentID
	}
	if model.Template != nil {
		templateMap, err := DataSourceIBMPolicyAssignmentPolicyAssignmentV1OptionsRootTemplateToMap(model.Template)
		if err != nil {
			return modelMap, err
		}
		modelMap["template"] = []map[string]interface{}{templateMap}
	}
	return modelMap, nil
}

func DataSourceIBMPolicyAssignmentPolicyAssignmentV1OptionsRootTemplateToMap(model *iampolicymanagementv1.PolicyAssignmentV1OptionsRootTemplate) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Version != nil {
		modelMap["version"] = *model.Version
	}
	return modelMap, nil
}

func DataSourceIBMPolicyAssignmentPolicyAssignmentV1ResourcesToMap(model *iampolicymanagementv1.PolicyAssignmentV1Resources) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Target != nil {
		targetMap, err := ResourceIBMPolicyAssignmentResourceTargetDetailsToMap(model.Target)
		if err != nil {
			return modelMap, err
		}
		modelMap["target"] = targetMap
	}
	if model.Policy != nil {
		policyMap, err := DataSourceIBMPolicyAssignmentPolicyAssignmentResourcePolicyToMap(model.Policy)
		if err != nil {
			return modelMap, err
		}
		modelMap["policy"] = []map[string]interface{}{policyMap}
	}
	return modelMap, nil
}

func DataSourceIBMPolicyAssignmentPolicyAssignmentResourcePolicyToMap(model *iampolicymanagementv1.PolicyAssignmentResourcePolicy) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ResourceCreated != nil {
		resourceCreatedMap, err := DataSourceIBMPolicyAssignmentAssignmentResourceCreatedToMap(model.ResourceCreated)
		if err != nil {
			return modelMap, err
		}
		modelMap["resource_created"] = []map[string]interface{}{resourceCreatedMap}
	}
	if model.Status != nil {
		modelMap["status"] = *model.Status
	}
	if model.ErrorMessage != nil {
		errorMessageMap, err := DataSourceIBMPolicyAssignmentErrorResponseToMap(model.ErrorMessage)
		if err != nil {
			return modelMap, err
		}
		modelMap["error_message"] = []map[string]interface{}{errorMessageMap}
	}
	return modelMap, nil
}

func DataSourceIBMPolicyAssignmentAssignmentResourceCreatedToMap(model *iampolicymanagementv1.AssignmentResourceCreated) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	return modelMap, nil
}

func DataSourceIBMPolicyAssignmentErrorResponseToMap(model *iampolicymanagementv1.ErrorResponse) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Trace != nil {
		modelMap["trace"] = *model.Trace
	}
	if model.Errors != nil {
		errors := []map[string]interface{}{}
		for _, errorsItem := range model.Errors {
			errorsItemMap, err := DataSourceIBMPolicyAssignmentErrorObjectToMap(&errorsItem)
			if err != nil {
				return modelMap, err
			}
			errors = append(errors, errorsItemMap)
		}
		modelMap["errors"] = errors
	}
	if model.StatusCode != nil {
		modelMap["status_code"] = flex.IntValue(model.StatusCode)
	}
	return modelMap, nil
}

func DataSourceIBMPolicyAssignmentErrorObjectToMap(model *iampolicymanagementv1.ErrorObject) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["code"] = *model.Code
	modelMap["message"] = *model.Message
	if model.Details != nil {
		detailsMap, err := DataSourceIBMPolicyAssignmentErrorDetailsToMap(model.Details)
		if err != nil {
			return modelMap, err
		}
		modelMap["details"] = []map[string]interface{}{detailsMap}
	}
	if model.MoreInfo != nil {
		modelMap["more_info"] = *model.MoreInfo
	}
	return modelMap, nil
}

func DataSourceIBMPolicyAssignmentErrorDetailsToMap(model *iampolicymanagementv1.ErrorDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ConflictsWith != nil {
		conflictsWithMap, err := DataSourceIBMPolicyAssignmentConflictsWithToMap(model.ConflictsWith)
		if err != nil {
			return modelMap, err
		}
		modelMap["conflicts_with"] = []map[string]interface{}{conflictsWithMap}
	}
	return modelMap, nil
}

func DataSourceIBMPolicyAssignmentConflictsWithToMap(model *iampolicymanagementv1.ConflictsWith) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Etag != nil {
		modelMap["etag"] = *model.Etag
	}
	if model.Role != nil {
		modelMap["role"] = *model.Role
	}
	if model.Policy != nil {
		modelMap["policy"] = *model.Policy
	}
	return modelMap, nil
}

func DataSourceIBMPolicyAssignmentPolicyAssignmentV1SubjectToMap(model *iampolicymanagementv1.PolicyAssignmentV1Subject) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	return modelMap, nil
}

func DataSourceIBMPolicyAssignmentAssignmentTemplateDetailsToMap(model *iampolicymanagementv1.AssignmentTemplateDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Version != nil {
		modelMap["version"] = *model.Version
	}
	return modelMap, nil
}

func DataSourceIBMPolicyAssignmentPolicyTemplateAssignmentItemsPolicyAssignmentV1ToMap(model *iampolicymanagementv1.PolicyTemplateAssignmentItemsPolicyAssignmentV1) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	targetMap, err := DataSourceIBMPolicyAssignmentAssignmentTargetDetailsToMap(model.Target)
	if err != nil {
		return modelMap, err
	}
	// modelMap["target"] = []map[string]interface{}{targetMap}
	modelMap["target"] = targetMap
	optionsMap, err := DataSourceIBMPolicyAssignmentPolicyAssignmentV1OptionsToMap(model.Options)
	if err != nil {
		return modelMap, err
	}
	modelMap["options"] = []map[string]interface{}{optionsMap}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.AccountID != nil {
		modelMap["account_id"] = *model.AccountID
	}
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	if model.CreatedAt != nil {
		modelMap["created_at"] = model.CreatedAt.String()
	}
	if model.CreatedByID != nil {
		modelMap["created_by_id"] = *model.CreatedByID
	}
	if model.LastModifiedAt != nil {
		modelMap["last_modified_at"] = model.LastModifiedAt.String()
	}
	if model.LastModifiedByID != nil {
		modelMap["last_modified_by_id"] = *model.LastModifiedByID
	}
	resources := []map[string]interface{}{}
	for _, resourcesItem := range model.Resources {
		resourcesItemMap, err := DataSourceIBMPolicyAssignmentPolicyAssignmentV1ResourcesToMap(&resourcesItem)
		if err != nil {
			return modelMap, err
		}
		resources = append(resources, resourcesItemMap)
	}
	modelMap["resources"] = resources
	if model.Subject != nil {
		subjectMap, err := DataSourceIBMPolicyAssignmentPolicyAssignmentV1SubjectToMap(model.Subject)
		if err != nil {
			return modelMap, err
		}
		modelMap["subject"] = []map[string]interface{}{subjectMap}
	}
	templateMap, err := DataSourceIBMPolicyAssignmentAssignmentTemplateDetailsToMap(model.Template)
	if err != nil {
		return modelMap, err
	}
	modelMap["template"] = []map[string]interface{}{templateMap}
	modelMap["status"] = *model.Status
	return modelMap, nil
}

func DataSourceIBMPolicyAssignmentPolicyTemplateAssignmentItemsPolicyAssignmentToMap(model *iampolicymanagementv1.PolicyTemplateAssignmentItemsPolicyAssignment) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.TemplateID != nil {
		modelMap["template_id"] = *model.TemplateID
	}
	if model.TemplateVersion != nil {
		modelMap["template_version"] = *model.TemplateVersion
	}
	if model.AssignmentID != nil {
		modelMap["assignment_id"] = *model.AssignmentID
	}
	if model.TargetType != nil {
		modelMap["target_type"] = *model.TargetType
	}
	if model.Target != nil {
		modelMap["target"] = *model.Target
	}
	if model.Options != nil {
		options := []map[string]interface{}{}
		for _, optionsItem := range model.Options {
			optionsItemMap, err := DataSourceIBMPolicyAssignmentPolicyAssignmentOptionsToMap(&optionsItem)
			if err != nil {
				return modelMap, err
			}
			options = append(options, optionsItemMap)
		}
		modelMap["options"] = options
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.AccountID != nil {
		modelMap["account_id"] = *model.AccountID
	}
	if model.Href != nil {
		modelMap["href"] = *model.Href
	}
	if model.CreatedAt != nil {
		modelMap["created_at"] = model.CreatedAt.String()
	}
	if model.CreatedByID != nil {
		modelMap["created_by_id"] = *model.CreatedByID
	}
	if model.LastModifiedAt != nil {
		modelMap["last_modified_at"] = model.LastModifiedAt.String()
	}
	if model.LastModifiedByID != nil {
		modelMap["last_modified_by_id"] = *model.LastModifiedByID
	}
	if model.Resources != nil {
		resources := []map[string]interface{}{}
		for _, resourcesItem := range model.Resources {
			resourcesItemMap, err := DataSourceIBMPolicyAssignmentPolicyAssignmentResourcesToMap(&resourcesItem)
			if err != nil {
				return modelMap, err
			}
			resources = append(resources, resourcesItemMap)
		}
		modelMap["resources"] = resources
	}
	if model.Status != nil {
		modelMap["status"] = *model.Status
	}
	return modelMap, nil
}

func DataSourceIBMPolicyAssignmentPolicyAssignmentOptionsToMap(model *iampolicymanagementv1.PolicyAssignmentOptions) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["subject_type"] = *model.SubjectType
	modelMap["subject_id"] = *model.SubjectID
	modelMap["root_requester_id"] = *model.RootRequesterID
	if model.RootTemplateID != nil {
		modelMap["root_template_id"] = *model.RootTemplateID
	}
	if model.RootTemplateVersion != nil {
		modelMap["root_template_version"] = *model.RootTemplateVersion
	}
	return modelMap, nil
}

func DataSourceIBMPolicyAssignmentPolicyAssignmentResourcesToMap(model *iampolicymanagementv1.PolicyAssignmentResources) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Target != nil {
		modelMap["target"] = *model.Target
	}
	if model.Policy != nil {
		policyMap, err := DataSourceIBMPolicyAssignmentPolicyAssignmentResourcePolicyToMap(model.Policy)
		if err != nil {
			return modelMap, err
		}
		modelMap["policy"] = []map[string]interface{}{policyMap}
	}
	return modelMap, nil
}
