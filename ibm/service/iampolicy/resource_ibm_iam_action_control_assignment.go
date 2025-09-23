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
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
)

func ResourceIBMIAMActionControlAssignment() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMActionControlAssignmentCreate,
		ReadContext:   resourceIBMActionControlAssignmentRead,
		UpdateContext: resourceIBMActionControlAssignmentUpdate,
		DeleteContext: resourceIBMActionControlAssignmentDelete,
		Importer:      &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"target": {
				Type:        schema.TypeMap,
				Required:    true,
				Description: "assignment target details",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"templates": {
				Type:        schema.TypeList,
				Required:    true,
				MinItems:    1,
				MaxItems:    1,
				Description: "action control template details.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "action control template id.",
						},
						"version": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "action control template version.",
						},
					},
				},
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
						"target": {
							Type:        schema.TypeMap,
							Required:    true,
							Description: "assignment target details",
							Elem: &schema.Schema{
								Type: schema.TypeString,
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
																						Optional:    true,
																						Computed:    true,
																						Description: "The conflicting role of ID.",
																					},
																					"policy": &schema.Schema{
																						Type:        schema.TypeString,
																						Optional:    true,
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
																Optional:    true,
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
			"status": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The action control assignment status.",
			},
			"template_version": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "The policy template version.",
			},
		},
	}
}

func resourceIBMActionControlAssignmentCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamPolicyManagementClient, err := meta.(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_assignment", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createActionControlTemplateAssignmentOptions := &iampolicymanagementv1.CreateActionControlTemplateAssignmentOptions{}

	targetModel, err := ResourceIBMActionControlAssignmentMapToAssignmentTargetDetails(d.Get("target").(map[string]interface{}))
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_assignment", "create", "parse-target").GetDiag()
	}
	createActionControlTemplateAssignmentOptions.SetTarget(targetModel)
	var templates []iampolicymanagementv1.ActionControlAssignmentTemplate
	for _, v := range d.Get("templates").([]interface{}) {
		value := v.(map[string]interface{})
		templatesItem, err := ResourceIBMActionControlAssignmentMapToAssignmentTemplateDetails(value)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_assignment", "create", "parse-templates").GetDiag()
		}
		templates = append(templates, *templatesItem)
	}

	createActionControlTemplateAssignmentOptions.SetTemplates(templates)

	actionControlAssignment, _, err := iamPolicyManagementClient.CreateActionControlTemplateAssignmentWithContext(context, createActionControlTemplateAssignmentOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateActionControlTemplateAssignmentWithContext failed: %s", err.Error()), "ibm_iam_action_control_assignment", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*actionControlAssignment.Assignments[0].ID)

	if targetModel.Type != nil && (*targetModel.Type == "Account") {
		log.Printf("[DEBUG] Skipping waitForAssignment for target type: %s", *targetModel.Type)
	} else {
		_, err = waitForAssignment(d.Timeout(schema.TimeoutCreate), meta, d, isActionControlAssignmentAssigned)
		if err != nil {
			return diag.FromErr(fmt.Errorf("error assigning: %s", err))
		}
	}

	return resourceIBMActionControlAssignmentRead(context, d, meta)
}

func resourceIBMActionControlAssignmentRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamPolicyManagementClient, err := meta.(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_assignment", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getActionControlAssignmentOptions := &iampolicymanagementv1.GetActionControlAssignmentOptions{}

	getActionControlAssignmentOptions.SetAssignmentID(d.Id())

	actionControlAssignment, response, err := iamPolicyManagementClient.GetActionControlAssignmentWithContext(context, getActionControlAssignmentOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetActionControlAssignmentWithContext failed: %s", err.Error()), "ibm_iam_action_control_assignment", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	targetMap, err := ResourceIBMActionControlAssignmentAssignmentTargetDetailsToMap(actionControlAssignment.Target)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_assignment", "read", "target-to-map").GetDiag()
	}
	if err = d.Set("target", targetMap); err != nil {
		err = fmt.Errorf("Error setting target: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_assignment", "read", "set-target").GetDiag()
	}
	if !core.IsNil(actionControlAssignment.AccountID) {
		if err = d.Set("account_id", actionControlAssignment.AccountID); err != nil {
			err = fmt.Errorf("Error setting account_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_assignment", "read", "set-account_id").GetDiag()
		}
	}
	if !core.IsNil(actionControlAssignment.Href) {
		if err = d.Set("href", actionControlAssignment.Href); err != nil {
			err = fmt.Errorf("Error setting href: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_assignment", "read", "set-href").GetDiag()
		}
	}
	if !core.IsNil(actionControlAssignment.CreatedAt) {
		if err = d.Set("created_at", flex.DateTimeToString(actionControlAssignment.CreatedAt)); err != nil {
			err = fmt.Errorf("Error setting created_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_assignment", "read", "set-created_at").GetDiag()
		}
	}
	if !core.IsNil(actionControlAssignment.CreatedByID) {
		if err = d.Set("created_by_id", actionControlAssignment.CreatedByID); err != nil {
			err = fmt.Errorf("Error setting created_by_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_assignment", "read", "set-created_by_id").GetDiag()
		}
	}
	if !core.IsNil(actionControlAssignment.LastModifiedAt) {
		if err = d.Set("last_modified_at", flex.DateTimeToString(actionControlAssignment.LastModifiedAt)); err != nil {
			err = fmt.Errorf("Error setting last_modified_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_assignment", "read", "set-last_modified_at").GetDiag()
		}
	}
	if !core.IsNil(actionControlAssignment.LastModifiedByID) {
		if err = d.Set("last_modified_by_id", actionControlAssignment.LastModifiedByID); err != nil {
			err = fmt.Errorf("Error setting last_modified_by_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_assignment", "read", "set-last_modified_by_id").GetDiag()
		}
	}
	if !core.IsNil(actionControlAssignment.Operation) {
		if err = d.Set("operation", actionControlAssignment.Operation); err != nil {
			err = fmt.Errorf("Error setting operation: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_assignment", "read", "set-operation").GetDiag()
		}
	}
	if !core.IsNil(actionControlAssignment.Resources) {
		resources := []map[string]interface{}{}
		for _, resourcesItem := range actionControlAssignment.Resources {
			resourcesItemMap, err := ResourceIBMActionControlAssignmentActionControlAssignmentResourceToMap(&resourcesItem)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_assignment", "read", "resources-to-map").GetDiag()
			}
			resources = append(resources, resourcesItemMap)
		}
		if err = d.Set("resources", resources); err != nil {
			err = fmt.Errorf("Error setting resources: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_assignment", "read", "set-resources").GetDiag()
		}
	}
	if !core.IsNil(actionControlAssignment.Status) {
		if err = d.Set("status", actionControlAssignment.Status); err != nil {
			err = fmt.Errorf("Error setting status: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_assignment", "read", "set-status").GetDiag()
		}
	}
	if !core.IsNil(actionControlAssignment.Template.Version) {
		if err = d.Set("template_version", actionControlAssignment.Template.Version); err != nil {
			err = fmt.Errorf("Error setting template_version: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_assignment", "read", "set-template_version").GetDiag()
		}
	}
	return nil
}

func resourceIBMActionControlAssignmentUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamPolicyManagementClient, err := meta.(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_assignment", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateActionControlAssignmentOptions := &iampolicymanagementv1.UpdateActionControlAssignmentOptions{}

	updateActionControlAssignmentOptions.SetAssignmentID(d.Id())
	targetModel, diags := GetTargetModel(d)
	if diags.HasError() {
		return diags
	}
	hasChange := false

	if d.HasChange("template_version") {
		updateActionControlAssignmentOptions.SetTemplateVersion(d.Get("template_version").(string))
		hasChange = true
	}

	if hasChange {
		getActionControlAssignmentOptions := &iampolicymanagementv1.GetActionControlAssignmentOptions{
			AssignmentID: core.StringPtr(d.Id()),
		}
		_, response, err := iamPolicyManagementClient.GetActionControlAssignmentWithContext(context, getActionControlAssignmentOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				d.SetId("")
				return nil
			}
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetActionControlAssignmentWithContext failed: %s", err.Error()), "ibm_policy_assignment", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		updateActionControlAssignmentOptions.SetIfMatch(response.Headers.Get("ETag"))
		_, _, err = iamPolicyManagementClient.UpdateActionControlAssignmentWithContext(context, updateActionControlAssignmentOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateActionControlAssignmentWithContext failed: %s", err.Error()), "ibm_iam_action_control_assignment", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		if targetModel.Type != nil && (*targetModel.Type == "Account") {
			log.Printf("[DEBUG] Skipping waitForAssignment for target type: %s", *targetModel.Type)
		} else {
			_, err = waitForAssignment(d.Timeout(schema.TimeoutUpdate), meta, d, isActionControlAssignmentAssigned)
			if err != nil {
				return diag.FromErr(fmt.Errorf("error assigning: %s", err))
			}
		}
	}

	return resourceIBMActionControlAssignmentRead(context, d, meta)
}

func resourceIBMActionControlAssignmentDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamPolicyManagementClient, err := meta.(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_assignment", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteActionControlAssignmentOptions := &iampolicymanagementv1.DeleteActionControlAssignmentOptions{}

	deleteActionControlAssignmentOptions.SetAssignmentID(d.Id())

	_, err = iamPolicyManagementClient.DeleteActionControlAssignmentWithContext(context, deleteActionControlAssignmentOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteActionControlAssignmentWithContext failed: %s", err.Error()), "ibm_iam_action_control_assignment", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	targetModel, diags := GetTargetModel(d)
	if diags.HasError() {
		return diags
	}

	if targetModel.Type != nil && (*targetModel.Type == "Account") {
		log.Printf("[DEBUG] Skipping waitForAssignment for target type: %s", *targetModel.Type)
	} else {
		_, err = waitForAssignment(d.Timeout(schema.TimeoutDelete), meta, d, isAccessActionControlAssignedDeleted)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return nil
			} else {
				return diag.FromErr(fmt.Errorf("error assigning: %s", err))
			}
		}
	}

	d.SetId("")

	return nil
}

func ResourceIBMActionControlAssignmentMapToAssignmentTargetDetails(modelMap map[string]interface{}) (*iampolicymanagementv1.AssignmentTargetDetails, error) {
	model := &iampolicymanagementv1.AssignmentTargetDetails{}
	if modelMap["type"] != nil && modelMap["type"].(string) != "" {
		model.Type = core.StringPtr(modelMap["type"].(string))
	}
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	return model, nil
}

func ResourceIBMActionControlAssignmentMapToAssignmentTemplateDetails(modelMap map[string]interface{}) (*iampolicymanagementv1.ActionControlAssignmentTemplate, error) {
	model := &iampolicymanagementv1.ActionControlAssignmentTemplate{}
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["version"] != nil && modelMap["version"].(string) != "" {
		model.Version = core.StringPtr(modelMap["version"].(string))
	}
	return model, nil
}

func ResourceIBMActionControlAssignmentAssignmentTargetDetailsToMap(model *iampolicymanagementv1.AssignmentTargetDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	return modelMap, nil
}

func ResourceIBMActionControlAssignmentActionControlAssignmentResourceToMap(model *iampolicymanagementv1.ActionControlAssignmentResource) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	targetMap, err := ResourceIBMActionControlAssignmentAssignmentTargetDetailsToMap(model.Target)
	if err != nil {
		return modelMap, err
	}
	modelMap["target"] = targetMap
	if model.ActionControl != nil {
		actionControlMap, err := ResourceIBMActionControlAssignmentActionControlAssignmentResourceActionControlToMap(model.ActionControl)
		if err != nil {
			return modelMap, err
		}
		modelMap["action_control"] = []map[string]interface{}{actionControlMap}
	}
	return modelMap, nil
}

func ResourceIBMActionControlAssignmentActionControlAssignmentResourceActionControlToMap(model *iampolicymanagementv1.ActionControlAssignmentResourceActionControl) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ResourceCreated != nil {
		resourceCreatedMap, err := ResourceIBMActionControlAssignmentActionControlAssignmentResourceCreatedToMap(model.ResourceCreated)
		if err != nil {
			return modelMap, err
		}
		modelMap["resource_created"] = []map[string]interface{}{resourceCreatedMap}
	}
	if model.ErrorMessage != nil {
		errorMessageMap, err := ResourceIBMActionControlAssignmentErrorResponseToMap(model.ErrorMessage)
		if err != nil {
			return modelMap, err
		}
		modelMap["error_message"] = []map[string]interface{}{errorMessageMap}
	}
	return modelMap, nil
}

func ResourceIBMActionControlAssignmentActionControlAssignmentResourceCreatedToMap(model *iampolicymanagementv1.ActionControlAssignmentResourceCreated) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	return modelMap, nil
}

func ResourceIBMActionControlAssignmentErrorResponseToMap(model *iampolicymanagementv1.AssignmentResourceError) (map[string]interface{}, error) {
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
			errorsItemMap, err := ResourceIBMActionControlAssignmentErrorObjectToMap(&errorsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			errors = append(errors, errorsItemMap)
		}
		modelMap["errors"] = errors
	}
	return modelMap, nil
}

func ResourceIBMActionControlAssignmentErrorObjectToMap(model *iampolicymanagementv1.ErrorObject) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["code"] = *model.Code
	modelMap["message"] = *model.Message
	if model.Details != nil {
		detailsMap, err := ResourceIBMActionControlAssignmentErrorDetailsToMap(model.Details)
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

func ResourceIBMActionControlAssignmentErrorDetailsToMap(model *iampolicymanagementv1.ErrorDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ConflictsWith != nil {
		conflictsWithMap, err := ResourceIBMActionControlAssignmentConflictsWithToMap(model.ConflictsWith)
		if err != nil {
			return modelMap, err
		}
		modelMap["conflicts_with"] = []map[string]interface{}{conflictsWithMap}
	}
	return modelMap, nil
}

func ResourceIBMActionControlAssignmentConflictsWithToMap(model *iampolicymanagementv1.ConflictsWith) (map[string]interface{}, error) {
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

func ResourceIBMActionControlAssignmentActionControlAssignmentTemplateToMap(model *iampolicymanagementv1.ActionControlAssignmentTemplate) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Version != nil {
		modelMap["version"] = *model.Version
	}
	return modelMap, nil
}

func isActionControlAssignmentAssigned(id string, meta interface{}) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		iamPolicyManagementClient, err := meta.(conns.ClientSession).IAMPolicyManagementV1API()
		if err != nil {
			return nil, READY, err
		}

		getAssignmentActionControlOptions := &iampolicymanagementv1.GetActionControlAssignmentOptions{
			AssignmentID: core.StringPtr(id),
		}

		getAssignmentActionControlOptions.SetAssignmentID(id)

		assignment, response, err := iamPolicyManagementClient.GetActionControlAssignment(getAssignmentActionControlOptions)

		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return nil, READY, err
			}
			return nil, READY, err
		}

		if assignment != nil {
			if *assignment.Status == "accepted" || *assignment.Status == "in_progress" {
				log.Printf("Assignment still in progress\n")
				return assignment, WAITING, nil
			}

			if *assignment.Status == "succeeded" {
				return assignment, READY, nil
			}

			if *assignment.Status == "failed" {
				return assignment, READY, fmt.Errorf("[ERROR] The assignment %s did not complete successfully and has a 'failed' status. Please check assignment resource for detailed errors: %s\n", id, response)
			}
		}

		return assignment, READY, fmt.Errorf("[ERROR] Unexpected status reached for assignment %s.: %s\n", id, response)
	}
}

func isAccessActionControlAssignedDeleted(id string, meta interface{}) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		iamPolicyManagementClient, err := meta.(conns.ClientSession).IAMPolicyManagementV1API()
		if err != nil {
			return nil, READY, err
		}

		getAssignmentActionControlOptions := &iampolicymanagementv1.GetActionControlAssignmentOptions{
			AssignmentID: core.StringPtr(id),
		}

		getAssignmentActionControlOptions.SetAssignmentID(id)

		assignment, response, err := iamPolicyManagementClient.GetActionControlAssignment(getAssignmentActionControlOptions)

		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return nil, READY, err
			}
			return nil, READY, err
		}

		if assignment != nil {
			if *assignment.Status == "accepted" || *assignment.Status == "in_progress" {
				log.Printf("Assignment still in progress\n")
				return assignment, WAITING, nil
			}

			if *assignment.Status == "failed" {
				return assignment, READY, fmt.Errorf("[ERROR] The assignment %s did not complete successfully and has a 'failed' status. Please check assignment resource for detailed errors: %s\n", id, response)
			}
		}

		return assignment, READY, fmt.Errorf("[ERROR] Unexpected status reached for assignment %s.: %s\n", id, response)
	}
}
