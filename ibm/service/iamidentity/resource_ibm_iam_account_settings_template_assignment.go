// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.113.1-d76630af-20260320-135953
 */

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
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template_assignment", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createAccountSettingsAssignmentOptions := &iamidentityv1.CreateAccountSettingsAssignmentOptions{}

	templateId, _, err := parseResourceId(d.Get("template_id").(string))
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template_assignment", "create", "parse-resource-id").GetDiag()
	}

	createAccountSettingsAssignmentOptions.SetTemplateID(templateId)
	createAccountSettingsAssignmentOptions.SetTemplateVersion(int64(d.Get("template_version").(int)))
	createAccountSettingsAssignmentOptions.SetTargetType(d.Get("target_type").(string))
	createAccountSettingsAssignmentOptions.SetTarget(d.Get("target").(string))

	templateAssignmentResponse, _, err := iamIdentityClient.CreateAccountSettingsAssignmentWithContext(context, createAccountSettingsAssignmentOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateAccountSettingsAssignmentWithContext failed: %s", err.Error()), "ibm_iam_account_settings_template_assignment", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*templateAssignmentResponse.ID)

	_, err = waitForAssignment(d.Timeout(schema.TimeoutCreate), meta, d, isAccountSettingsTemplateAssigned)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template_assignment", "create", "wait-for-assignment").GetDiag()
	}

	// Read persists the final state. When status is "failed"/"superseded", Read
	// writes template_version=0 so the next plan shows a visible update diff.
	diags := resourceIBMAccountSettingsTemplateAssignmentRead(context, d, meta)
	if status, ok := d.GetOk("status"); ok && (status.(string) == "failed" || status.(string) == "superseded") {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Assignment completed with status '%s'.", status.(string)),
			Detail: fmt.Sprintf(
				"The assignment %s is in a '%s' state. Terraform has marked this resource as tainted.\n"+
					"To retry without destroying and recreating the assignment, run:\n\n"+
					"  terraform untaint ibm_iam_account_settings_template_assignment.<RESOURCE_NAME>\n"+
					"  terraform apply\n\n"+
					"Replace <RESOURCE_NAME> with the name of your resource block (e.g. if your config\n"+
					"is 'resource \"ibm_iam_account_settings_template_assignment\" \"assignment\"', use:\n\n"+
					"  terraform untaint ibm_iam_account_settings_template_assignment.assignment\n"+
					"  terraform apply\n\n"+
					"This will perform an in-place update (PUT) to retry only the failed resources.",
				d.Id(), status.(string)),
		})
	}
	return diags
}

func resourceIBMAccountSettingsTemplateAssignmentRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template_assignment", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getAccountSettingsAssignmentOptions := &iamidentityv1.GetAccountSettingsAssignmentOptions{}

	getAccountSettingsAssignmentOptions.SetAssignmentID(d.Id())

	templateAssignmentResponse, response, err := iamIdentityClient.GetAccountSettingsAssignmentWithContext(context, getAccountSettingsAssignmentOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetAccountSettingsAssignmentWithContext failed: %s", err.Error()), "ibm_iam_account_settings_template_assignment", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("template_id", templateAssignmentResponse.TemplateID); err != nil {
		err = fmt.Errorf("Error setting template_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template_assignment", "read", "set-template_id").GetDiag()
	}
	// When the assignment is in a failed/superseded state, write 0 for
	// template_version so Terraform sees a diff against the config value on the
	// next plan and calls Update, which retries the assignment via PUT.
	templateVersion := flex.IntValue(templateAssignmentResponse.TemplateVersion)
	if templateAssignmentResponse.Status != nil &&
		(*templateAssignmentResponse.Status == "failed" || *templateAssignmentResponse.Status == "superseded") {
		templateVersion = 0
	}
	if err = d.Set("template_version", templateVersion); err != nil {
		err = fmt.Errorf("Error setting template_version: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template_assignment", "read", "set-template_version").GetDiag()
	}
	if err = d.Set("target_type", templateAssignmentResponse.TargetType); err != nil {
		err = fmt.Errorf("Error setting target_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template_assignment", "read", "set-target_type").GetDiag()
	}
	if err = d.Set("target", templateAssignmentResponse.Target); err != nil {
		err = fmt.Errorf("Error setting target: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template_assignment", "read", "set-target").GetDiag()
	}
	if err = d.Set("account_id", templateAssignmentResponse.AccountID); err != nil {
		err = fmt.Errorf("Error setting account_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template_assignment", "read", "set-account_id").GetDiag()
	}
	if err = d.Set("status", templateAssignmentResponse.Status); err != nil {
		err = fmt.Errorf("Error setting status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template_assignment", "read", "set-status").GetDiag()
	}
	var resources []map[string]interface{}
	if !core.IsNil(templateAssignmentResponse.Resources) {
		for _, resourcesItem := range templateAssignmentResponse.Resources {
			resourcesItemMap, err := resourceIBMAccountSettingsTemplateAssignmentTemplateAssignmentResponseResourceToMap(&resourcesItem)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template_assignment", "read", "resources-to-map").GetDiag()
			}
			resources = append(resources, resourcesItemMap)
		}
	}
	if err = d.Set("resources", resources); err != nil {
		err = fmt.Errorf("Error setting resources: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template_assignment", "read", "set-resources").GetDiag()
	}
	if !core.IsNil(templateAssignmentResponse.Href) {
		if err = d.Set("href", templateAssignmentResponse.Href); err != nil {
			err = fmt.Errorf("Error setting href: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template_assignment", "read", "set-href").GetDiag()
		}
	}
	if err = d.Set("created_at", templateAssignmentResponse.CreatedAt); err != nil {
		err = fmt.Errorf("Error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template_assignment", "read", "set-created_at").GetDiag()
	}
	if err = d.Set("created_by_id", templateAssignmentResponse.CreatedByID); err != nil {
		err = fmt.Errorf("Error setting created_by_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template_assignment", "read", "set-created_by_id").GetDiag()
	}
	if err = d.Set("last_modified_at", templateAssignmentResponse.LastModifiedAt); err != nil {
		err = fmt.Errorf("Error setting last_modified_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template_assignment", "read", "set-last_modified_at").GetDiag()
	}
	if err = d.Set("last_modified_by_id", templateAssignmentResponse.LastModifiedByID); err != nil {
		err = fmt.Errorf("Error setting last_modified_by_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template_assignment", "read", "set-last_modified_by_id").GetDiag()
	}
	if err = d.Set("entity_tag", templateAssignmentResponse.EntityTag); err != nil {
		err = fmt.Errorf("Error setting entity_tag: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template_assignment", "read", "set-entity_tag").GetDiag()
	}
	return nil
}

func resourceIBMAccountSettingsTemplateAssignmentUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template_assignment", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateAccountSettingsAssignmentOptions := &iamidentityv1.UpdateAccountSettingsAssignmentOptions{}
	updateAccountSettingsAssignmentOptions.SetAssignmentID(d.Id())
	updateAccountSettingsAssignmentOptions.SetIfMatch(d.Get("entity_tag").(string))

	// Always set template_version on the options. When retrying a failed/superseded
	// assignment, Read wrote 0 to state so HasChange is true and d.Get returns the
	// config value (the real version the user specified).
	updateAccountSettingsAssignmentOptions.SetTemplateVersion(int64(d.Get("template_version").(int)))

	if d.HasChange("template_version") {
		_, response, err := iamIdentityClient.UpdateAccountSettingsAssignmentWithContext(context, updateAccountSettingsAssignmentOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateAccountSettingsAssignmentWithContext failed: %s", err.Error()), "ibm_iam_account_settings_template_assignment", "update")
			log.Printf("[DEBUG]\n%s\n%s", tfErr.GetDebugMessage(), response)
			return tfErr.GetDiag()
		}

		_, err = waitForAssignment(d.Timeout(schema.TimeoutUpdate), meta, d, isAccountSettingsTemplateAssigned)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template_assignment", "update", "wait-for-assignment").GetDiag()
		}
	}

	diags := resourceIBMAccountSettingsTemplateAssignmentRead(context, d, meta)
	if status, ok := d.GetOk("status"); ok && (status.(string) == "failed" || status.(string) == "superseded") {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Assignment completed with status '%s'. Run terraform apply again to retry.", status.(string)),
		})
	}
	return diags
}

func resourceIBMAccountSettingsTemplateAssignmentDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamIdentityClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template_assignment", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteAccountSettingsAssignmentOptions := &iamidentityv1.DeleteAccountSettingsAssignmentOptions{}

	deleteAccountSettingsAssignmentOptions.SetAssignmentID(d.Id())

	_, _, err = iamIdentityClient.DeleteAccountSettingsAssignmentWithContext(context, deleteAccountSettingsAssignmentOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteAccountSettingsAssignmentWithContext failed: %s", err.Error()), "ibm_iam_account_settings_template_assignment", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	_, err = waitForAssignment(d.Timeout(schema.TimeoutDelete), meta, d, isAccountSettingsAssignmentRemoved)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_account_settings_template_assignment", "delete", "wait-for-assignment").GetDiag()
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
		}

		if assignment != nil && assignment.Status != nil && *assignment.Status == "failed" {
			return assignment, READY, fmt.Errorf("[ERROR] The deletion of assignment %s completed with a 'failed' status. Please check the assignment resource for detailed errors", id)
		}

		log.Printf("Assignment removal still in progress\n")

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

			if *assignment.Status == "failed" || *assignment.Status == "superseded" {
				log.Printf("[WARN] Assignment %s completed with status '%s'\n", id, *assignment.Status)
				return assignment, READY, nil
			}

			return assignment, READY, nil
		}

		return assignment, READY, fmt.Errorf("[ERROR] Unexpected status reached for assignment %s.: %s", id, response)
	}
}
