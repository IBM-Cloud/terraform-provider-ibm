// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

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

func ResourceIBMIAMPolicyAssignment() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMPolicyAssignmentCreate,
		ReadContext:   resourceIBMPolicyAssignmentRead,
		UpdateContext: resourceIBMPolicyAssignmentUpdate,
		DeleteContext: resourceIBMPolicyAssignmentDelete,
		Importer:      &schema.ResourceImporter{},

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
				Description: "policy template details.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "policy template id.",
						},
						"version": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "policy template version.",
						},
					},
				},
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
			"template": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "policy template details",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The policy assignment status.",
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

func resourceIBMPolicyAssignmentCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamPolicyManagementClient, err := meta.(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_policy_assignment", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createPolicyTemplateAssignmentOptions := &iampolicymanagementv1.CreatePolicyTemplateAssignmentOptions{}

	createPolicyTemplateAssignmentOptions.SetVersion(d.Get("version").(string))
	targetModel, err := ResourceIBMPolicyAssignmentMapToAssignmentTargetDetails(d.Get("target").(map[string]interface{}))
	if err != nil {
		return diag.FromErr(err)
	}
	createPolicyTemplateAssignmentOptions.SetTarget(targetModel)
	var templates []iampolicymanagementv1.AssignmentTemplateDetails
	for _, v := range d.Get("templates").([]interface{}) {
		value := v.(map[string]interface{})
		templatesItem, err := ResourceIBMPolicyAssignmentMapToAssignmentTemplateDetails(value)
		if err != nil {
			return diag.FromErr(err)
		}
		templates = append(templates, *templatesItem)
	}
	createPolicyTemplateAssignmentOptions.SetTemplates(templates)
	if _, ok := d.GetOk("accept_language"); ok {
		createPolicyTemplateAssignmentOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}
	policyAssignmentV1Collection, _, err := iamPolicyManagementClient.CreatePolicyTemplateAssignmentWithContext(context, createPolicyTemplateAssignmentOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreatePolicyTemplateAssignmentWithContext failed: %s", err.Error()), "ibm_policy_assignment", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*policyAssignmentV1Collection.Assignments[0].ID)

	return resourceIBMPolicyAssignmentRead(context, d, meta)
}

func resourceIBMPolicyAssignmentRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamPolicyManagementClient, err := meta.(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_policy_assignment", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getPolicyAssignmentOptions := &iampolicymanagementv1.GetPolicyAssignmentOptions{
		AssignmentID: core.StringPtr(d.Id()),
		Version:      core.StringPtr("1.0"),
	}

	assignmentResponse, response, err := iamPolicyManagementClient.GetPolicyAssignmentWithContext(context, getPolicyAssignmentOptions)

	assignmentDetails := assignmentResponse.(*iampolicymanagementv1.GetPolicyAssignmentResponse)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetPolicyAssignmentWithContext failed: %s", err.Error()), "ibm_policy_assignment", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	targetMap, err := ResourceIBMPolicyAssignmentAssignmentTargetDetailsToMap(assignmentDetails.Target)
	if err != nil {
		return diag.FromErr(err)
	}

	if err = d.Set("target", targetMap); err != nil {
		return diag.FromErr(fmt.Errorf("error setting target: %s", err))
	}
	if !core.IsNil(assignmentDetails.AccountID) {
		if err = d.Set("account_id", assignmentDetails.AccountID); err != nil {
			return diag.FromErr(fmt.Errorf("error setting account_id: %s", err))
		}
	}
	if !core.IsNil(assignmentDetails.Href) {
		if err = d.Set("href", assignmentDetails.Href); err != nil {
			return diag.FromErr(fmt.Errorf("error setting href: %s", err))
		}
	}
	if !core.IsNil(assignmentDetails.CreatedAt) {
		if err = d.Set("created_at", flex.DateTimeToString(assignmentDetails.CreatedAt)); err != nil {
			return diag.FromErr(fmt.Errorf("error setting created_at: %s", err))
		}
	}
	if !core.IsNil(assignmentDetails.CreatedByID) {
		if err = d.Set("created_by_id", assignmentDetails.CreatedByID); err != nil {
			return diag.FromErr(fmt.Errorf("error setting created_by_id: %s", err))
		}
	}
	if !core.IsNil(assignmentDetails.LastModifiedAt) {
		if err = d.Set("last_modified_at", flex.DateTimeToString(assignmentDetails.LastModifiedAt)); err != nil {
			return diag.FromErr(fmt.Errorf("error setting last_modified_at: %s", err))
		}
	}
	if !core.IsNil(assignmentDetails.LastModifiedByID) {
		if err = d.Set("last_modified_by_id", assignmentDetails.LastModifiedByID); err != nil {
			return diag.FromErr(fmt.Errorf("error setting last_modified_by_id: %s", err))
		}
	}
	resources := []map[string]interface{}{}
	for _, resourcesItem := range assignmentDetails.Resources {
		resourcesItemMap, err := ResourceIBMPolicyAssignmentPolicyAssignmentV1ResourcesToMap(&resourcesItem)
		if err != nil {
			return diag.FromErr(err)
		}
		resources = append(resources, resourcesItemMap)
	}
	if err = d.Set("resources", resources); err != nil {
		return diag.FromErr(fmt.Errorf("error setting resources: %s", err))
	}
	templateMap, err := ResourceIBMPolicyAssignmentAssignmentTemplateDetailsToMap(assignmentDetails.Template)
	if err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("template", templateMap); err != nil {
		return diag.FromErr(fmt.Errorf("error setting template: %s", err))
	}
	if err = d.Set("template_version", assignmentDetails.Template.Version); err != nil {
		return diag.FromErr(fmt.Errorf("error setting template: %s", err))
	}
	if err = d.Set("status", assignmentDetails.Status); err != nil {
		return diag.FromErr(fmt.Errorf("error setting status: %s", err))
	}

	return nil
}

func resourceIBMPolicyAssignmentUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamPolicyManagementClient, err := meta.(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_policy_assignment", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updatePolicyAssignmentOptions := &iampolicymanagementv1.UpdatePolicyAssignmentOptions{}

	updatePolicyAssignmentOptions.SetAssignmentID(d.Id())

	hasChange := false

	if d.HasChange("template_version") {
		updatePolicyAssignmentOptions.SetVersion(d.Get("version").(string))
		updatePolicyAssignmentOptions.SetTemplateVersion(d.Get("template_version").(string))
		hasChange = true
	}

	if hasChange {
		getPolicyAssignmentOptions := &iampolicymanagementv1.GetPolicyAssignmentOptions{
			AssignmentID: core.StringPtr(d.Id()),
			Version:      core.StringPtr("1.0"),
		}
		_, response, err := iamPolicyManagementClient.GetPolicyAssignmentWithContext(context, getPolicyAssignmentOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				d.SetId("")
				return nil
			}
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetPolicyAssignmentWithContext failed: %s", err.Error()), "ibm_policy_assignment", "read")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		updatePolicyAssignmentOptions.SetIfMatch(response.Headers.Get("ETag"))
		_, _, err = iamPolicyManagementClient.UpdatePolicyAssignmentWithContext(context, updatePolicyAssignmentOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdatePolicyAssignmentWithContext failed: %s", err.Error()), "ibm_policy_assignment", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIBMPolicyAssignmentRead(context, d, meta)
}

func resourceIBMPolicyAssignmentDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamPolicyManagementClient, err := meta.(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "ibm_policy_assignment", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deletePolicyAssignmentOptions := &iampolicymanagementv1.DeletePolicyAssignmentOptions{}

	deletePolicyAssignmentOptions.SetAssignmentID(d.Id())

	_, err = iamPolicyManagementClient.DeletePolicyAssignmentWithContext(context, deletePolicyAssignmentOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeletePolicyAssignmentWithContext failed: %s", err.Error()), "ibm_policy_assignment", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIBMPolicyAssignmentMapToAssignmentTargetDetails(modelMap map[string]interface{}) (*iampolicymanagementv1.AssignmentTargetDetails, error) {
	model := &iampolicymanagementv1.AssignmentTargetDetails{}
	if modelMap["type"] != nil && modelMap["type"].(string) != "" {
		model.Type = core.StringPtr(modelMap["type"].(string))
	}
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	return model, nil
}

func ResourceIBMPolicyAssignmentMapToAssignmentTemplateDetails(modelMap map[string]interface{}) (*iampolicymanagementv1.AssignmentTemplateDetails, error) {
	model := &iampolicymanagementv1.AssignmentTemplateDetails{}
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["version"] != nil && modelMap["version"].(string) != "" {
		model.Version = core.StringPtr(modelMap["version"].(string))
	}
	return model, nil
}

func ResourceIBMPolicyAssignmentAssignmentTemplateDetailsToMap(model *iampolicymanagementv1.AssignmentTemplateDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Version != nil {
		modelMap["version"] = *model.Version
	}
	return modelMap, nil
}
