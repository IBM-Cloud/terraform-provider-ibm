// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
)

func ResourceIBMAccountSettingsTemplateAssignment() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMAccountSettingsTemplateAssignmentCreate,
		ReadContext:   resourceIBMAccountSettingsTemplateAssignmentRead,
		UpdateContext: resourceIBMAccountSettingsTemplateAssignmentUpdate,
		DeleteContext: resourceIBMAccountSettingsTemplateAssignmentDelete,
		Importer:      &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"template_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Template Id.",
			},
			"template_version": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Template version.",
			},
			"target_type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_iam_account_settings_template_assignment", "target_type"),
				Description:  "Assignment target type.",
			},
			"target": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Assignment target.",
			},
			"account_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Enterprise account Id.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Assignment status.",
			},
			"resources": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Status breakdown per target account of IAM resources created or errors encountered in attempting to create those IAM resources. IAM resources are only included in the response providing the assignment is not in progress. IAM resources are also only included when getting a single assignment, and excluded by list APIs.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"target": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Target account where the IAM resource is created.",
						},
						"account_settings": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"resource_created": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Body parameters for created resource.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Id of the created resource.",
												},
											},
										},
									},
									"error_message": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Body parameters for assignment error.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
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
												"status_code": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Internal status code for the error.",
												},
											},
										},
									},
									"status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Status for the target account's assignment.",
									},
								},
							},
						},
					},
				},
			},
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Href.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Assignment created at.",
			},
			"created_by_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IAMid of the identity that created the assignment.",
			},
			"last_modified_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Assignment modified at.",
			},
			"last_modified_by_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IAMid of the identity that last modified the assignment.",
			},
			"entity_tag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Entity tag for this assignment record.",
			},
		},
	}
}

func ResourceIBMAccountSettingsTemplateAssignmentValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "target_type",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "Account, AccountGroup",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_iam_account_settings_template_assignment", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMAccountSettingsTemplateAssignmentCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		return diag.FromErr(err)
	}

	createAccountSettingsAssignmentOptions := &iamidentityv1.CreateAccountSettingsAssignmentOptions{}

	templateId, _, err := parseResourceId(d.Get("template_id").(string))
	if err != nil {
		log.Printf("[DEBUG] resourceIBMAccountSettingsTemplateRead failed %s", err)
		return diag.FromErr(fmt.Errorf("resourceIBMAccountSettingsTemplateRead failed %s", err))
	}

	createAccountSettingsAssignmentOptions.SetTemplateID(templateId)
	createAccountSettingsAssignmentOptions.SetTemplateVersion(int64(d.Get("template_version").(int)))
	createAccountSettingsAssignmentOptions.SetTargetType(d.Get("target_type").(string))
	createAccountSettingsAssignmentOptions.SetTarget(d.Get("target").(string))

	templateAssignmentResponse, response, err := iamIdentityClient.CreateAccountSettingsAssignmentWithContext(context, createAccountSettingsAssignmentOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateAccountSettingsAssignmentWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateAccountSettingsAssignmentWithContext failed %s\n%s", err, response))
	}

	d.SetId(*templateAssignmentResponse.ID)

	_, err = waitForAssignment(d.Timeout(schema.TimeoutCreate), meta, d, isAccountSettingsTemplateAssigned)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error assigning %s", err))
	}

	return resourceIBMAccountSettingsTemplateAssignmentRead(context, d, meta)
}

func resourceIBMAccountSettingsTemplateAssignmentRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		return diag.FromErr(err)
	}

	getAccountSettingsAssignmentOptions := &iamidentityv1.GetAccountSettingsAssignmentOptions{}

	getAccountSettingsAssignmentOptions.SetAssignmentID(d.Id())

	templateAssignmentResponse, response, err := iamIdentityClient.GetAccountSettingsAssignmentWithContext(context, getAccountSettingsAssignmentOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetAccountSettingsAssignmentWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetAccountSettingsAssignmentWithContext failed %s\n%s", err, response))
	}

	if err = d.Set("template_id", templateAssignmentResponse.TemplateID); err != nil {
		return diag.FromErr(fmt.Errorf("error setting template_id: %s", err))
	}
	if err = d.Set("template_version", flex.IntValue(templateAssignmentResponse.TemplateVersion)); err != nil {
		return diag.FromErr(fmt.Errorf("error setting template_version: %s", err))
	}
	if err = d.Set("target_type", templateAssignmentResponse.TargetType); err != nil {
		return diag.FromErr(fmt.Errorf("error setting target_type: %s", err))
	}
	if err = d.Set("target", templateAssignmentResponse.Target); err != nil {
		return diag.FromErr(fmt.Errorf("error setting target: %s", err))
	}
	if err = d.Set("account_id", templateAssignmentResponse.AccountID); err != nil {
		return diag.FromErr(fmt.Errorf("error setting account_id: %s", err))
	}
	if err = d.Set("status", templateAssignmentResponse.Status); err != nil {
		return diag.FromErr(fmt.Errorf("error setting status: %s", err))
	}
	var resources []map[string]interface{}
	if !core.IsNil(templateAssignmentResponse.Resources) {
		for _, resourcesItem := range templateAssignmentResponse.Resources {
			resourcesItemMap, err := resourceIBMAccountSettingsTemplateAssignmentTemplateAssignmentResponseResourceToMap(&resourcesItem)
			if err != nil {
				return diag.FromErr(err)
			}
			resources = append(resources, resourcesItemMap)
		}
	}
	if err = d.Set("resources", resources); err != nil {
		return diag.FromErr(fmt.Errorf("error setting resources: %s", err))
	}
	if !core.IsNil(templateAssignmentResponse.Href) {
		if err = d.Set("href", templateAssignmentResponse.Href); err != nil {
			return diag.FromErr(fmt.Errorf("error setting href: %s", err))
		}
	}
	if err = d.Set("created_at", templateAssignmentResponse.CreatedAt); err != nil {
		return diag.FromErr(fmt.Errorf("error setting created_at: %s", err))
	}
	if err = d.Set("created_by_id", templateAssignmentResponse.CreatedByID); err != nil {
		return diag.FromErr(fmt.Errorf("error setting created_by_id: %s", err))
	}
	if err = d.Set("last_modified_at", templateAssignmentResponse.LastModifiedAt); err != nil {
		return diag.FromErr(fmt.Errorf("error setting last_modified_at: %s", err))
	}
	if err = d.Set("last_modified_by_id", templateAssignmentResponse.LastModifiedByID); err != nil {
		return diag.FromErr(fmt.Errorf("error setting last_modified_by_id: %s", err))
	}
	if err = d.Set("entity_tag", templateAssignmentResponse.EntityTag); err != nil {
		return diag.FromErr(fmt.Errorf("error setting entity_tag: %s", err))
	}

	return nil
}

func resourceIBMAccountSettingsTemplateAssignmentUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		return diag.FromErr(err)
	}

	updateAccountSettingsAssignmentOptions := &iamidentityv1.UpdateAccountSettingsAssignmentOptions{}
	updateAccountSettingsAssignmentOptions.SetAssignmentID(d.Id())
	updateAccountSettingsAssignmentOptions.SetIfMatch(d.Get("entity_tag").(string))

	hasChange := false

	if d.HasChange("template_version") {
		updateAccountSettingsAssignmentOptions.SetTemplateVersion(int64(d.Get("template_version").(int)))
		hasChange = true
	}

	if hasChange || d.Get("status") == "failed" { // allow the same version to retry failed assignments
		_, response, err := iamIdentityClient.UpdateAccountSettingsAssignmentWithContext(context, updateAccountSettingsAssignmentOptions)
		if err != nil {
			log.Printf("[DEBUG] UpdateAccountSettingsAssignmentWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("UpdateAccountSettingsAssignmentWithContext failed %s\n%s", err, response))
		}

		_, err = waitForAssignment(d.Timeout(schema.TimeoutUpdate), meta, d, isAccountSettingsTemplateAssigned)
		if err != nil {
			return diag.FromErr(fmt.Errorf("error assigning %s", err))
		}
	}

	return resourceIBMAccountSettingsTemplateAssignmentRead(context, d, meta)
}

func resourceIBMAccountSettingsTemplateAssignmentDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteAccountSettingsAssignmentOptions := &iamidentityv1.DeleteAccountSettingsAssignmentOptions{}

	deleteAccountSettingsAssignmentOptions.SetAssignmentID(d.Id())

	_, response, err := iamIdentityClient.DeleteAccountSettingsAssignmentWithContext(context, deleteAccountSettingsAssignmentOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteAccountSettingsAssignmentWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteAccountSettingsAssignmentWithContext failed %s\n%s", err, response))
	}

	_, err = waitForAssignment(d.Timeout(schema.TimeoutDelete), meta, d, isAccountSettingsAssignmentRemoved)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error removing assignment %s", err))
	}

	d.SetId("")

	return nil
}

func resourceIBMAccountSettingsTemplateAssignmentTemplateAssignmentResponseResourceToMap(model *iamidentityv1.TemplateAssignmentResponseResource) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["target"] = model.Target
	if model.AccountSettings != nil {
		accountSettingsMap, err := resourceIBMAccountSettingsTemplateAssignmentTemplateAssignmentResponseResourceDetailToMap(model.AccountSettings)
		if err != nil {
			return modelMap, err
		}
		modelMap["account_settings"] = []map[string]interface{}{accountSettingsMap}
	}
	if model.PolicyTemplateReferences != nil {
		var policyTemplateRefs []map[string]interface{}
		for _, policyTemplateRefsItem := range model.PolicyTemplateReferences {
			policyTemplateRefsItemMap, err := resourceIBMAccountSettingsTemplateAssignmentTemplateAssignmentResponseResourceDetailToMap(&policyTemplateRefsItem)
			if err != nil {
				return modelMap, err
			}
			policyTemplateRefs = append(policyTemplateRefs, policyTemplateRefsItemMap)
		}
		modelMap["policy_template_references"] = policyTemplateRefs
	}
	return modelMap, nil
}

func resourceIBMAccountSettingsTemplateAssignmentTemplateAssignmentResponseResourceDetailToMap(model *iamidentityv1.TemplateAssignmentResponseResourceDetail) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Version != nil {
		modelMap["version"] = model.Version
	}
	if model.ResourceCreated != nil {
		resourceCreatedMap, err := resourceIBMAccountSettingsTemplateAssignmentTemplateAssignmentResourceToMap(model.ResourceCreated)
		if err != nil {
			return modelMap, err
		}
		modelMap["resource_created"] = []map[string]interface{}{resourceCreatedMap}
	}
	if model.ErrorMessage != nil {
		errorMessageMap, err := resourceIBMAccountSettingsTemplateAssignmentTemplateAssignmentResourceErrorToMap(model.ErrorMessage)
		if err != nil {
			return modelMap, err
		}
		modelMap["error_message"] = []map[string]interface{}{errorMessageMap}
	}
	modelMap["status"] = model.Status
	return modelMap, nil
}

func resourceIBMAccountSettingsTemplateAssignmentTemplateAssignmentResourceToMap(model *iamidentityv1.TemplateAssignmentResource) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	return modelMap, nil
}

func resourceIBMAccountSettingsTemplateAssignmentTemplateAssignmentResourceErrorToMap(model *iamidentityv1.TemplateAssignmentResourceError) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.ErrorCode != nil {
		modelMap["error_code"] = model.ErrorCode
	}
	if model.Message != nil {
		modelMap["message"] = model.Message
	}
	if model.StatusCode != nil {
		modelMap["status_code"] = model.StatusCode
	}
	return modelMap, nil
}

func isAccountSettingsAssignmentRemoved(id string, meta interface{}) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		iamIdentityClient, _ := meta.(conns.ClientSession).IAMIdentityV1API()

		getOptions := &iamidentityv1.GetAccountSettingsAssignmentOptions{
			AssignmentID: &id,
		}
		assignment, response, err := iamIdentityClient.GetAccountSettingsAssignment(getOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				return assignment, READY, nil
			}

			return nil, READY, fmt.Errorf("[ERROR] The assignment %s failed to delete or deletion was not completed within specific timeout period: %s\n%s", id, err, response)
		} else {
			log.Printf("Assignment removal still in progress\n")
		}

		return assignment, WAITING, nil
	}
}

func isAccountSettingsTemplateAssigned(id string, meta interface{}) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {

		iamIdentityClient, _ := meta.(conns.ClientSession).IAMIdentityV1API()

		getOptions := &iamidentityv1.GetAccountSettingsAssignmentOptions{
			AssignmentID: &id,
		}
		assignment, response, err := iamIdentityClient.GetAccountSettingsAssignment(getOptions)
		if err != nil {
			return nil, READY, fmt.Errorf("[ERROR] The assignment %s failed or did not complete within specific timeout period: %s\n%s", id, err, response)
		}

		if assignment != nil {
			if *assignment.Status == "accepted" || *assignment.Status == "in_progress" {
				log.Printf("Assignment still in progress\n")
				return assignment, WAITING, nil
			}

			if *assignment.Status == "failed" {
				return assignment, READY, fmt.Errorf("[ERROR] The assignment %s did complete but with a 'failed' status. Please check assignment resource for detailed errors: %s", id, response)
			}

			return assignment, READY, nil
		}

		return assignment, READY, fmt.Errorf("[ERROR] Unexpected status reached for assignment %s.: %s", id, response)
	}
}
