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

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
)

func DataSourceIBMIAMRoleAssignments() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMRoleAssignmentsRead,

		Schema: map[string]*schema.Schema{
			"template_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Optional template ID.",
			},
			"template_version": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Optional role control template version.",
			},
			"assignments": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of role control assignments.",
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
							Description: "The account GUID that the role control assignments belong to.",
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The href URL that links to the role control assignments API by role control assignment ID.",
						},
						"created_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The UTC timestamp when the role control assignment was created.",
						},
						"created_by_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IAM ID of the entity that created the role control assignment.",
						},
						"last_modified_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The UTC timestamp when the role control assignment was last modified.",
						},
						"last_modified_by_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IAM ID of the entity that last modified the role control assignment.",
						},
						"operation": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The current operation of the role control assignment.",
						},
						"resources": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Resources created when role control template is assigned.",
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
									"role": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Set of properties of the assigned resource or error message if assignment failed.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"resource_created": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "On success, it includes the role control assigned.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "role control id.",
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
																Description: "The unique transrole ID for the request.",
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
							Description: "The role control template id and version that will be assigned.",
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
							Description: "The role control assignment status.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMRoleAssignmentsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamPolicyManagementClient, err := meta.(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_list_role_assignments", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listRoleAssignmentsOptions := &iampolicymanagementv1.ListRoleAssignmentsOptions{}

	userDetails, err := meta.(conns.ClientSession).BluemixUserDetails()
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to fetch BluemixUserDetails %s", err))
	}
	accountID := userDetails.UserAccount
	listRoleAssignmentsOptions.SetAccountID(accountID)

	if _, ok := d.GetOk("template_id"); ok {
		listRoleAssignmentsOptions.SetTemplateID(d.Get("template_id").(string))
	}
	if _, ok := d.GetOk("template_version"); ok {
		listRoleAssignmentsOptions.SetTemplateVersion(d.Get("template_version").(string))
	}

	var pager *iampolicymanagementv1.RoleAssignmentsPager
	pager, err = iamPolicyManagementClient.NewRoleAssignmentsPager(listRoleAssignmentsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_list_role_assignments", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	allItems, err := pager.GetAll()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("RoleAssignmentsPager.GetAll() failed %s", err), "(Data) ibm_list_role_assignments", "read")
		log.Printf("[DEBUG] %s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIBMListRoleAssignmentsID(d))

	mapSlice := []map[string]interface{}{}
	for _, modelItem := range allItems {
		modelMap, err := DataSourceIBMListRoleAssignmentsRoleAssignmentToMap(&modelItem)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_list_role_assignments", "read", "RoleAssignments-to-map").GetDiag()
		}
		mapSlice = append(mapSlice, modelMap)
	}

	if err = d.Set("assignments", mapSlice); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting assignments %s", err), "(Data) ibm_list_role_assignments", "read", "assignments-set").GetDiag()
	}

	return nil
}

// dataSourceIBMListRoleAssignmentsID returns a reasonable ID for the list.
func dataSourceIBMListRoleAssignmentsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIBMListRoleAssignmentsRoleAssignmentToMap(model *iampolicymanagementv1.RoleAssignment) (map[string]interface{}, error) {
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
			resourcesItemMap, err := DataSourceIBMListRoleAssignmentsRoleAssignmentResourceToMap(&resourcesItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			resources = append(resources, resourcesItemMap)
		}
		modelMap["resources"] = resources
	}
	templateMap, err := DataSourceIBMListRoleAssignmentsRoleAssignmentTemplateToMap(model.Template)
	if err != nil {
		return modelMap, err
	}
	modelMap["template"] = []map[string]interface{}{templateMap}
	targetMap, err := DataSourceIBMListRoleAssignmentsAssignmentTargetDetailsToMap(model.Target)
	if err != nil {
		return modelMap, err
	}
	modelMap["target"] = []map[string]interface{}{targetMap}
	if model.Status != nil {
		modelMap["status"] = *model.Status
	}
	return modelMap, nil
}

func DataSourceIBMListRoleAssignmentsRoleAssignmentResourceToMap(model *iampolicymanagementv1.RoleAssignmentResource) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	targetMap, err := DataSourceIBMListRoleAssignmentsAssignmentTargetDetailsToMap(model.Target)
	if err != nil {
		return modelMap, err
	}
	modelMap["target"] = []map[string]interface{}{targetMap}
	if model.Role != nil {
		roleControlMap, err := DataSourceIBMListRoleAssignmentsRoleAssignmentResourceRoleToMap(model.Role)
		if err != nil {
			return modelMap, err
		}
		modelMap["role"] = []map[string]interface{}{roleControlMap}
	}
	return modelMap, nil
}

func DataSourceIBMListRoleAssignmentsAssignmentTargetDetailsToMap(model *iampolicymanagementv1.AssignmentTargetDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["type"] = *model.Type
	modelMap["id"] = *model.ID
	return modelMap, nil
}

func DataSourceIBMListRoleAssignmentsRoleAssignmentResourceRoleToMap(model *iampolicymanagementv1.RoleAssignmentResourceRole) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ResourceCreated != nil {
		resourceCreatedMap, err := DataSourceIBMListRoleAssignmentsRoleAssignmentResourceCreatedToMap(model.ResourceCreated)
		if err != nil {
			return modelMap, err
		}
		modelMap["resource_created"] = []map[string]interface{}{resourceCreatedMap}
	}
	if model.ErrorMessage != nil {
		errorMessageMap, err := DataSourceIBMListRoleAssignmentsErrorResponseToMap(model.ErrorMessage)
		if err != nil {
			return modelMap, err
		}
		modelMap["error_message"] = []map[string]interface{}{errorMessageMap}
	}
	return modelMap, nil
}

func DataSourceIBMListRoleAssignmentsRoleAssignmentResourceCreatedToMap(model *iampolicymanagementv1.RoleAssignmentResourceCreated) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	return modelMap, nil
}

func DataSourceIBMListRoleAssignmentsErrorResponseToMap(model *iampolicymanagementv1.AssignmentResourceError) (map[string]interface{}, error) {
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
			errorsItemMap, err := DataSourceIBMListRoleAssignmentsErrorObjectToMap(&errorsItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			errors = append(errors, errorsItemMap)
		}
		modelMap["errors"] = errors
	}
	return modelMap, nil
}

func DataSourceIBMListRoleAssignmentsErrorObjectToMap(model *iampolicymanagementv1.ErrorObject) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["code"] = *model.Code
	modelMap["message"] = *model.Message
	if model.Details != nil {
		detailsMap, err := DataSourceIBMListRoleAssignmentsErrorDetailsToMap(model.Details)
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

func DataSourceIBMListRoleAssignmentsErrorDetailsToMap(model *iampolicymanagementv1.ErrorDetails) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ConflictsWith != nil {
		conflictsWithMap, err := DataSourceIBMListRoleAssignmentsConflictsWithToMap(model.ConflictsWith)
		if err != nil {
			return modelMap, err
		}
		modelMap["conflicts_with"] = []map[string]interface{}{conflictsWithMap}
	}
	return modelMap, nil
}

func DataSourceIBMListRoleAssignmentsConflictsWithToMap(model *iampolicymanagementv1.ConflictsWith) (map[string]interface{}, error) {
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

func DataSourceIBMListRoleAssignmentsRoleAssignmentTemplateToMap(model *iampolicymanagementv1.RoleAssignmentTemplate) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = *model.ID
	modelMap["version"] = *model.Version
	return modelMap, nil
}
