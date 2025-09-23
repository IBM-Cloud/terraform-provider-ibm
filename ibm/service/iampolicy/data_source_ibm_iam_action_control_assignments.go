// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.103.0-e8b84313-20250402-201816
 */

package iampolicy

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
)

func DataSourceIBMIAMActionControlAssignments() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMActionControlAssignmentsRead,

		Schema: map[string]*schema.Schema{
			"template_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Optional template ID.",
			},
			"template_version": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Optional action control template version.",
			},
			"assignments": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of action control assignments.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Action control assignment ID.",
						},
						"account_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account GUID that the action control assignments belong to.",
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The href URL that links to the action control assignments API by action control assignment ID.",
						},
						"created_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The UTC timestamp when the action control assignment was created.",
						},
						"created_by_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IAM ID of the entity that created the action control assignment.",
						},
						"last_modified_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The UTC timestamp when the action control assignment was last modified.",
						},
						"last_modified_by_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IAM ID of the entity that last modified the action control assignment.",
						},
						"operation": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The current operation of the action control assignment.",
						},
						"resources": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Resources created when action control template is assigned.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "assignment target account and type.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Assignment target type.",
												},
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "ID of the target account.",
												},
											},
										},
									},
									"action_control": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Set of properties of the assigned resource or error message if assignment failed.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"resource_created": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "On success, it includes the action control assigned.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "action control id.",
															},
														},
													},
												},
												"error_message": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The error response from API.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"trace": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The unique transaction ID for the request.",
															},
															"errors": &schema.Schema{
																Type:        schema.TypeList,
																Computed:    true,
																Description: "The errors encountered during the response.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"code": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The API error code for the error.",
																		},
																		"message": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The error message returned by the API.",
																		},
																		"details": &schema.Schema{
																			Type:        schema.TypeList,
																			Computed:    true,
																			Description: "Additional error details.",
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"conflicts_with": &schema.Schema{
																						Type:        schema.TypeList,
																						Computed:    true,
																						Description: "Details of conflicting resource.",
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"etag": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "The revision number of the resource.",
																								},
																								"role": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "The conflicting role of ID.",
																								},
																								"policy": &schema.Schema{
																									Type:        schema.TypeString,
																									Computed:    true,
																									Description: "The conflicting policy ID.",
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																		"more_info": &schema.Schema{
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Additional info for error.",
																		},
																	},
																},
															},
															"status_code": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The HTTP error code of the response.",
															},
															"name": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Name of the error.",
															},
															"error_code": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Internal error code.",
															},
															"message": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Error message detailing the nature of the error.",
															},
															"code": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Internal status code for the error.",
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
						"template": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The action control template id and version that will be assigned.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Action control template ID.",
									},
									"version": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Action control template version.",
									},
								},
							},
						},
						"target": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "assignment target account and type.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Assignment target type.",
									},
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "ID of the target account.",
									},
								},
							},
						},
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The action control assignment status.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMActionControlAssignmentsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamPolicyManagementClient, err := meta.(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_list_action_control_assignments", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listActionControlAssignmentsOptions := &iampolicymanagementv1.ListActionControlAssignmentsOptions{}

	userDetails, err := meta.(conns.ClientSession).BluemixUserDetails()
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to fetch BluemixUserDetails %s", err))
	}
	accountID := userDetails.UserAccount
	listActionControlAssignmentsOptions.SetAccountID(accountID)

	if _, ok := d.GetOk("template_id"); ok {
		listActionControlAssignmentsOptions.SetTemplateID(d.Get("template_id").(string))
	}
	if _, ok := d.GetOk("template_version"); ok {
		listActionControlAssignmentsOptions.SetTemplateVersion(d.Get("template_version").(string))
	}

	var pager *iampolicymanagementv1.ActionControlAssignmentsPager
	pager, err = iamPolicyManagementClient.NewActionControlAssignmentsPager(listActionControlAssignmentsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_list_action_control_assignments", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	allItems, err := pager.GetAll()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ActionControlAssignmentsPager.GetAll() failed %s", err), "(Data) ibm_list_action_control_assignments", "read")
		log.Printf("[DEBUG] %s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIBMListActionControlAssignmentsID(d))

	mapSlice := []map[string]interface{}{}
	for _, modelItem := range allItems {
		modelMap, err := DataSourceIBMListActionControlAssignmentsActionControlAssignmentToMap(&modelItem)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_list_action_control_assignments", "read", "ActionControlAssignments-to-map").GetDiag()
		}
		mapSlice = append(mapSlice, modelMap)
	}

	if err = d.Set("assignments", mapSlice); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting assignments %s", err), "(Data) ibm_list_action_control_assignments", "read", "assignments-set").GetDiag()
	}

	return nil
}

// dataSourceIBMListActionControlAssignmentsID returns a reasonable ID for the list.
func dataSourceIBMListActionControlAssignmentsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIBMListActionControlAssignmentsActionControlAssignmentToMap(model *iampolicymanagementv1.ActionControlAssignment) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
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
	if model.Operation != nil {
		modelMap["operation"] = *model.Operation
	}
	if model.Resources != nil {
		resources := []map[string]interface{}{}
		for _, resourcesItem := range model.Resources {
			resourcesItemMap, err := DataSourceIBMListActionControlAssignmentsActionControlAssignmentResourceToMap(&resourcesItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			resources = append(resources, resourcesItemMap)
		}
		modelMap["resources"] = resources
	}
	templateMap, err := DataSourceIBMListActionControlAssignmentsActionControlAssignmentTemplateToMap(model.Template)
	if err != nil {
		return modelMap, err
	}
	modelMap["template"] = []map[string]interface{}{templateMap}
	targetMap, err := DataSourceIBMListActionControlAssignmentsAssignmentTargetDetailsToMap(model.Target)
	if err != nil {
		return modelMap, err
	}
	modelMap["target"] = []map[string]interface{}{targetMap}
	if model.Status != nil {
		modelMap["status"] = *model.Status
	}
	return modelMap, nil
}

func DataSourceIBMListActionControlAssignmentsActionControlAssignmentResourceToMap(model *iampolicymanagementv1.ActionControlAssignmentResource) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	targetMap, err := DataSourceIBMListActionControlAssignmentsAssignmentTargetDetailsToMap(model.Target)
	if err != nil {
		return modelMap, err
	}
	modelMap["target"] = []map[string]interface{}{targetMap}
	if model.ActionControl != nil {
		actionControlMap, err := DataSourceIBMListActionControlAssignmentsActionControlAssignmentResourceActionControlToMap(model.ActionControl)
		if err != nil {
			return modelMap, err
		}
		modelMap["action_control"] = []map[string]interface{}{actionControlMap}
	}
	return modelMap, nil
}

func DataSourceIBMListActionControlAssignmentsAssignmentTargetDetailsToMap(model *iampolicymanagementv1.AssignmentTargetDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["type"] = *model.Type
	modelMap["id"] = *model.ID
	return modelMap, nil
}

func DataSourceIBMListActionControlAssignmentsActionControlAssignmentResourceActionControlToMap(model *iampolicymanagementv1.ActionControlAssignmentResourceActionControl) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ResourceCreated != nil {
		resourceCreatedMap, err := DataSourceIBMListActionControlAssignmentsActionControlAssignmentResourceCreatedToMap(model.ResourceCreated)
		if err != nil {
			return modelMap, err
		}
		modelMap["resource_created"] = []map[string]interface{}{resourceCreatedMap}
	}
	if model.ErrorMessage != nil {
		errorMessageMap, err := DataSourceIBMListActionControlAssignmentsErrorResponseToMap(model.ErrorMessage)
		if err != nil {
			return modelMap, err
		}
		modelMap["error_message"] = []map[string]interface{}{errorMessageMap}
	}
	return modelMap, nil
}

func DataSourceIBMListActionControlAssignmentsActionControlAssignmentResourceCreatedToMap(model *iampolicymanagementv1.ActionControlAssignmentResourceCreated) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	return modelMap, nil
}

func DataSourceIBMListActionControlAssignmentsErrorResponseToMap(model *iampolicymanagementv1.AssignmentResourceError) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.ErrorCode != nil {
		modelMap["error_code"] = *model.ErrorCode
	}
	if model.Message != nil {
		modelMap["message"] = *model.Message
	}
	if model.Code != nil {
		modelMap["code"] = *model.Code
	}
	if model.Errors != nil {
		errors := []map[string]interface{}{}
		for _, errorsItem := range model.Errors {
			errorsItemMap, err := DataSourceIBMListActionControlAssignmentsErrorObjectToMap(&errorsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			errors = append(errors, errorsItemMap)
		}
		modelMap["errors"] = errors
	}
	return modelMap, nil
}

func DataSourceIBMListActionControlAssignmentsErrorObjectToMap(model *iampolicymanagementv1.ErrorObject) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["code"] = *model.Code
	modelMap["message"] = *model.Message
	if model.Details != nil {
		detailsMap, err := DataSourceIBMListActionControlAssignmentsErrorDetailsToMap(model.Details)
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

func DataSourceIBMListActionControlAssignmentsErrorDetailsToMap(model *iampolicymanagementv1.ErrorDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ConflictsWith != nil {
		conflictsWithMap, err := DataSourceIBMListActionControlAssignmentsConflictsWithToMap(model.ConflictsWith)
		if err != nil {
			return modelMap, err
		}
		modelMap["conflicts_with"] = []map[string]interface{}{conflictsWithMap}
	}
	return modelMap, nil
}

func DataSourceIBMListActionControlAssignmentsConflictsWithToMap(model *iampolicymanagementv1.ConflictsWith) (map[string]interface{}, error) {
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

func DataSourceIBMListActionControlAssignmentsActionControlAssignmentTemplateToMap(model *iampolicymanagementv1.ActionControlAssignmentTemplate) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = *model.ID
	modelMap["version"] = *model.Version
	return modelMap, nil
}
