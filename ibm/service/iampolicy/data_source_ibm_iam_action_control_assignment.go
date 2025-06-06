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

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
)

func DataSourceIBMIAMActionControlAssignment() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMActionControlAssignmentRead,

		Schema: map[string]*schema.Schema{
			"assignment_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Action control template assignment ID.",
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
	}
}

func dataSourceIBMActionControlAssignmentRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamPolicyManagementClient, err := meta.(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_get_action_control_assignment", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getActionControlAssignmentOptions := &iampolicymanagementv1.GetActionControlAssignmentOptions{}

	getActionControlAssignmentOptions.SetAssignmentID(d.Get("assignment_id").(string))

	actionControlAssignment, _, err := iamPolicyManagementClient.GetActionControlAssignmentWithContext(context, getActionControlAssignmentOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetActionControlAssignmentWithContext failed: %s", err.Error()), "(Data) ibm_get_action_control_assignment", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*actionControlAssignment.ID)

	if !core.IsNil(actionControlAssignment.AccountID) {
		if err = d.Set("account_id", actionControlAssignment.AccountID); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting account_id: %s", err), "(Data) ibm_get_action_control_assignment", "read", "set-account_id").GetDiag()
		}
	}

	if !core.IsNil(actionControlAssignment.Href) {
		if err = d.Set("href", actionControlAssignment.Href); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting href: %s", err), "(Data) ibm_get_action_control_assignment", "read", "set-href").GetDiag()
		}
	}

	if !core.IsNil(actionControlAssignment.CreatedAt) {
		if err = d.Set("created_at", flex.DateTimeToString(actionControlAssignment.CreatedAt)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_get_action_control_assignment", "read", "set-created_at").GetDiag()
		}
	}

	if !core.IsNil(actionControlAssignment.CreatedByID) {
		if err = d.Set("created_by_id", actionControlAssignment.CreatedByID); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting created_by_id: %s", err), "(Data) ibm_get_action_control_assignment", "read", "set-created_by_id").GetDiag()
		}
	}

	if !core.IsNil(actionControlAssignment.LastModifiedAt) {
		if err = d.Set("last_modified_at", flex.DateTimeToString(actionControlAssignment.LastModifiedAt)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting last_modified_at: %s", err), "(Data) ibm_get_action_control_assignment", "read", "set-last_modified_at").GetDiag()
		}
	}

	if !core.IsNil(actionControlAssignment.LastModifiedByID) {
		if err = d.Set("last_modified_by_id", actionControlAssignment.LastModifiedByID); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting last_modified_by_id: %s", err), "(Data) ibm_get_action_control_assignment", "read", "set-last_modified_by_id").GetDiag()
		}
	}

	if !core.IsNil(actionControlAssignment.Operation) {
		if err = d.Set("operation", actionControlAssignment.Operation); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting operation: %s", err), "(Data) ibm_get_action_control_assignment", "read", "set-operation").GetDiag()
		}
	}

	if !core.IsNil(actionControlAssignment.Resources) {
		resources := []map[string]interface{}{}
		for _, resourcesItem := range actionControlAssignment.Resources {
			resourcesItemMap, err := DataSourceIBMGetActionControlAssignmentActionControlAssignmentResourceToMap(&resourcesItem)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_get_action_control_assignment", "read", "resources-to-map").GetDiag()
			}
			resources = append(resources, resourcesItemMap)
		}
		if err = d.Set("resources", resources); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resources: %s", err), "(Data) ibm_get_action_control_assignment", "read", "set-resources").GetDiag()
		}
	}

	template := []map[string]interface{}{}
	templateMap, err := DataSourceIBMGetActionControlAssignmentActionControlAssignmentTemplateToMap(actionControlAssignment.Template)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_get_action_control_assignment", "read", "template-to-map").GetDiag()
	}
	template = append(template, templateMap)
	if err = d.Set("template", template); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting template: %s", err), "(Data) ibm_get_action_control_assignment", "read", "set-template").GetDiag()
	}

	target := []map[string]interface{}{}
	targetMap, err := DataSourceIBMGetActionControlAssignmentAssignmentTargetDetailsToMap(actionControlAssignment.Target)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_get_action_control_assignment", "read", "target-to-map").GetDiag()
	}
	target = append(target, targetMap)
	if err = d.Set("target", target); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting target: %s", err), "(Data) ibm_get_action_control_assignment", "read", "set-target").GetDiag()
	}

	if !core.IsNil(actionControlAssignment.Status) {
		if err = d.Set("status", actionControlAssignment.Status); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting status: %s", err), "(Data) ibm_get_action_control_assignment", "read", "set-status").GetDiag()
		}
	}

	return nil
}

func DataSourceIBMGetActionControlAssignmentActionControlAssignmentResourceToMap(model *iampolicymanagementv1.ActionControlAssignmentResource) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	targetMap, err := DataSourceIBMGetActionControlAssignmentAssignmentTargetDetailsToMap(model.Target)
	if err != nil {
		return modelMap, err
	}
	modelMap["target"] = []map[string]interface{}{targetMap}
	if model.ActionControl != nil {
		actionControlMap, err := DataSourceIBMGetActionControlAssignmentActionControlAssignmentResourceActionControlToMap(model.ActionControl)
		if err != nil {
			return modelMap, err
		}
		modelMap["action_control"] = []map[string]interface{}{actionControlMap}
	}
	return modelMap, nil
}

func DataSourceIBMGetActionControlAssignmentAssignmentTargetDetailsToMap(model *iampolicymanagementv1.AssignmentTargetDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["type"] = *model.Type
	modelMap["id"] = *model.ID
	return modelMap, nil
}

func DataSourceIBMGetActionControlAssignmentActionControlAssignmentResourceActionControlToMap(model *iampolicymanagementv1.ActionControlAssignmentResourceActionControl) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ResourceCreated != nil {
		resourceCreatedMap, err := DataSourceIBMGetActionControlAssignmentActionControlAssignmentResourceCreatedToMap(model.ResourceCreated)
		if err != nil {
			return modelMap, err
		}
		modelMap["resource_created"] = []map[string]interface{}{resourceCreatedMap}
	}
	if model.ErrorMessage != nil {
		errorMessageMap, err := DataSourceIBMGetActionControlAssignmentErrorResponseToMap(model.ErrorMessage)
		if err != nil {
			return modelMap, err
		}
		modelMap["error_message"] = []map[string]interface{}{errorMessageMap}
	}
	return modelMap, nil
}

func DataSourceIBMGetActionControlAssignmentActionControlAssignmentResourceCreatedToMap(model *iampolicymanagementv1.ActionControlAssignmentResourceCreated) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	return modelMap, nil
}

func DataSourceIBMGetActionControlAssignmentErrorResponseToMap(model *iampolicymanagementv1.ErrorResponse) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["trace"] = *model.Trace
	errors := []map[string]interface{}{}
	for _, errorsItem := range model.Errors {
		errorsItemMap, err := DataSourceIBMGetActionControlAssignmentErrorObjectToMap(&errorsItem)
		if err != nil {
			return modelMap, err
		}
		errors = append(errors, errorsItemMap)
	}
	modelMap["errors"] = errors
	modelMap["status_code"] = flex.IntValue(model.StatusCode)
	return modelMap, nil
}

func DataSourceIBMGetActionControlAssignmentErrorObjectToMap(model *iampolicymanagementv1.ErrorObject) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["code"] = *model.Code
	modelMap["message"] = *model.Message
	if model.Details != nil {
		detailsMap, err := DataSourceIBMGetActionControlAssignmentErrorDetailsToMap(model.Details)
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

func DataSourceIBMGetActionControlAssignmentErrorDetailsToMap(model *iampolicymanagementv1.ErrorDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ConflictsWith != nil {
		conflictsWithMap, err := DataSourceIBMGetActionControlAssignmentConflictsWithToMap(model.ConflictsWith)
		if err != nil {
			return modelMap, err
		}
		modelMap["conflicts_with"] = []map[string]interface{}{conflictsWithMap}
	}
	return modelMap, nil
}

func DataSourceIBMGetActionControlAssignmentConflictsWithToMap(model *iampolicymanagementv1.ConflictsWith) (map[string]interface{}, error) {
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

func DataSourceIBMGetActionControlAssignmentActionControlAssignmentTemplateToMap(model *iampolicymanagementv1.ActionControlAssignmentTemplate) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = *model.ID
	modelMap["version"] = *model.Version
	return modelMap, nil
}
