// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.106.0-09823488-20250707-071701
 */

package db2

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/resourcecontroller"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/cloud-db2-go-sdk/db2saasv1"
	"github.com/IBM/go-sdk-core/v5/core"
)

func ResourceIbmDb2SaasUsers() *schema.Resource {
	riSchema := resourcecontroller.ResourceIBMResourceInstance().Schema

	riSchema["users_config"] = &schema.Schema{
		Description: "The db2 new users gets created (available only for platform users)",
		Optional:    true,
		Type:        schema.TypeList,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"id": {
					Description: "The id of the user",
					Optional:    true,
					Type:        schema.TypeString,
				},
				"iam": {
					Description: "The iam of the user",
					Optional:    true,
					Type:        schema.TypeBool,
				},
				"ibmid": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "The ibmid of the user",
				},
				"name": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "The name of the user",
				},
				"password": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "The password of the user",
				},
				"role": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "The role of the user (say: bluuser)",
				},
				"email": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "The email of the user ",
				},
				"locked": {
					Type:        schema.TypeString,
					Optional:    true,
					Description: "It decribes if user is locked or not",
				},
				"authentication": {
					Type:        schema.TypeList,
					Optional:    true,
					Description: "The authentication for user",
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"method": {
								Type:        schema.TypeString,
								Optional:    true,
								Description: "The method of authentication for user",
							},
							"policy_id": {
								Type:        schema.TypeString,
								Optional:    true,
								Description: "The policy_id for the user",
							},
						},
					},
				},
			},
		},
	}

	return &schema.Resource{
		// CreateContext: resourceIbmDb2SaasUsersCreate,
		ReadContext:   resourceIbmDb2SaasUsersRead,
		UpdateContext: resourceIbmDb2SaasUsersUpdate,
		// DeleteContext: resourceIbmDb2SaasUsersDelete,
		Importer: &schema.ResourceImporter{},

		Schema: riSchema,
		/*map[string]*schema.Schema{
			"x_deployment_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				// ValidateFunc: validate.InvokeValidator("ibm_db2_saas_users", "x_deployment_id"),
				Description: "CRN deployment id.",
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				// ValidateFunc: validate.InvokeValidator("ibm_db2_saas_users", "role"),
				Description: "Role assigned to the user.",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "User's password.",
			},
			"iam": &schema.Schema{
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Indicates if IAM is enabled or not.",
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				// ValidateFunc: validate.InvokeValidator("ibm_db2_saas_users", "name"),
				Description: "The display name of the user.",
			},
			"ibmid": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				// ValidateFunc: validate.InvokeValidator("ibm_db2_saas_users", "ibmid"),
				Description: "IBM ID of the user.",
			},
			"locked": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				// ValidateFunc: validate.InvokeValidator("ibm_db2_saas_users", "locked"),
				Description: "Account lock status for the user.",
			},
			"email": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Email address of the user.",
			},
			"authentication": &schema.Schema{
				Type:        schema.TypeList,
				MinItems:    1,
				MaxItems:    1,
				Required:    true,
				Description: "Authentication details for the user.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"method": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Authentication method.",
						},
						"policy_id": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Policy ID of authentication.",
						},
					},
				},
			},
			"dv_role": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "User's DV role.",
			},
			"metadata": &schema.Schema{
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Metadata associated with the user.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"formated_ibmid": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Formatted IBM ID.",
			},
			"iamid": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IAM ID for the user.",
			},
			"permitted_actions": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of allowed actions of the user.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"all_clean": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates if the user account has no issues.",
			},
			"init_error_msg": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Initial error message.",
			},
		},
		*/
	}
}

func ResourceIbmDb2SaasUsersValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "x_deployment_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^crn(:[A-Za-z0-9\-\.]*){9}$`,
			MinValueLength:             1,
			MaxValueLength:             63,
		},
		validate.ValidateSchema{
			Identifier:                 "role",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "bluadmin, bluuser",
		},
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexp,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^.*$`,
		},
		validate.ValidateSchema{
			Identifier:                 "ibmid",
			ValidateFunctionIdentifier: validate.ValidateRegexp,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^.*$`,
		},
		validate.ValidateSchema{
			Identifier:                 "locked",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "no, yes",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_db2_saas_users", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmDb2SaasUsersCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	db2saasClient, err := meta.(conns.ClientSession).Db2saasV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_db2_saas_users", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	postDb2SaasUserOptions := &db2saasv1.PostDb2SaasUserOptions{}

	postDb2SaasUserOptions.SetXDeploymentID(d.Get("x_deployment_id").(string))
	postDb2SaasUserOptions.SetIam(d.Get("iam").(bool))
	postDb2SaasUserOptions.SetIbmid(d.Get("ibmid").(string))
	postDb2SaasUserOptions.SetName(d.Get("name").(string))
	postDb2SaasUserOptions.SetPassword(d.Get("password").(string))
	postDb2SaasUserOptions.SetRole(d.Get("role").(string))
	postDb2SaasUserOptions.SetEmail(d.Get("email").(string))
	postDb2SaasUserOptions.SetLocked(d.Get("locked").(string))
	authenticationModel, err := ResourceIbmDb2SaasUsersMapToCreateUserAuthentication(d.Get("authentication.0").(map[string]interface{}))
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_db2_saas_users", "create", "parse-authentication").GetDiag()
	}
	postDb2SaasUserOptions.SetAuthentication(authenticationModel)

	successUserResponse, _, err := db2saasClient.PostDb2SaasUserWithContext(context, postDb2SaasUserOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("PostDb2SaasUserWithContext failed: %s", err.Error()), "ibm_db2_saas_users", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*successUserResponse.ID)

	return resourceIbmDb2SaasUsersRead(context, d, meta)
}

func resourceIbmDb2SaasUsersRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	db2saasClient, err := meta.(conns.ClientSession).Db2saasV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_db2_saas_users", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getDb2SaasUserOptions := &db2saasv1.GetDb2SaasUserOptions{}

	// getDb2SaasUserOptions.SetID(d.Id())
	getDb2SaasUserOptions.SetXDeploymentID(d.Get("x_deployment_id").(string))

	successGetUserInfo, response, err := db2saasClient.GetDb2SaasUserWithContext(context, getDb2SaasUserOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDb2SaasUserWithContext failed: %s", err.Error()), "ibm_db2_saas_users", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	resources := []map[string]interface{}{}
	for _, resourcesItem := range successGetUserInfo.Resources {
		resourcesItemMap, err := DataSourceIbmDb2SaasUsersSuccessGetUserInfoResourcesItemToMap(&resourcesItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_db2_users", "read", "resources-to-map").GetDiag()
		}
		resources = append(resources, resourcesItemMap)
	}
	if err = d.Set("resources", resources); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resources: %s", err), "(Data) ibm_db2_users", "read", "set-resources").GetDiag()
	}

	// info := successGetUserInfo.Resources
	// users := make([]map[string]interface{}, 0, len(info))

	// for _, user := range info {
	// 	userMap := make(map[string]interface{})
	// 	userMap["id"] = user.ID
	// 	userMap["name"] = user.Name
	// 	userMap["email"] = user.Email
	// 	userMap["role"] = user.Role
	// 	userMap["locked"] = user.Locked
	// 	userMap["ibmid"] = user.Ibmid
	// 	userMap["iam"] = user.Iam
	// 	userMap["formated_ibmid"] = user.FormatedIbmid
	// 	users = append(users, userMap)
	// }

	// if err = d.Set("role", info.Role); err != nil {
	// 	err = fmt.Errorf("Error setting role: %s", err)
	// 	return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_db2_saas_users", "read", "set-role").GetDiag()
	// }
	// if err = d.Set("password", successGetUserInfo.Password); err != nil {
	// 	err = fmt.Errorf("Error setting password: %s", err)
	// 	return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_db2_saas_users", "read", "set-password").GetDiag()
	// }
	// if err = d.Set("iam", successGetUserInfo.Iam); err != nil {
	// 	err = fmt.Errorf("Error setting iam: %s", err)
	// 	return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_db2_saas_users", "read", "set-iam").GetDiag()
	// }
	// if err = d.Set("name", successGetUserInfo.Name); err != nil {
	// 	err = fmt.Errorf("Error setting name: %s", err)
	// 	return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_db2_saas_users", "read", "set-name").GetDiag()
	// }
	// if err = d.Set("ibmid", successGetUserInfo.Ibmid); err != nil {
	// 	err = fmt.Errorf("Error setting ibmid: %s", err)
	// 	return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_db2_saas_users", "read", "set-ibmid").GetDiag()
	// }
	// if err = d.Set("locked", successGetUserInfo.Locked); err != nil {
	// 	err = fmt.Errorf("Error setting locked: %s", err)
	// 	return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_db2_saas_users", "read", "set-locked").GetDiag()
	// }
	// if err = d.Set("email", successGetUserInfo.Email); err != nil {
	// 	err = fmt.Errorf("Error setting email: %s", err)
	// 	return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_db2_saas_users", "read", "set-email").GetDiag()
	// }
	// authenticationMap, err := ResourceIbmDb2SaasUsersSuccessUserResponseAuthenticationToMap(successGetUserInfo.Authentication)
	// if err != nil {
	// 	return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_db2_saas_users", "read", "authentication-to-map").GetDiag()
	// }
	// if err = d.Set("authentication", []map[string]interface{}{authenticationMap}); err != nil {
	// 	err = fmt.Errorf("Error setting authentication: %s", err)
	// 	return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_db2_saas_users", "read", "set-authentication").GetDiag()
	// }
	// if err = d.Set("dv_role", successGetUserInfo.DvRole); err != nil {
	// 	err = fmt.Errorf("Error setting dv_role: %s", err)
	// 	return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_db2_saas_users", "read", "set-dv_role").GetDiag()
	// }
	// metadata := make(map[string]string)
	// for k, v := range successGetUserInfo.Metadata {
	// 	metadata[k] = flex.Stringify(v)
	// }
	// if err = d.Set("metadata", metadata); err != nil {
	// 	err = fmt.Errorf("Error setting metadata: %s", err)
	// 	return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_db2_saas_users", "read", "set-metadata").GetDiag()
	// }
	// if err = d.Set("formated_ibmid", successGetUserInfo.FormatedIbmid); err != nil {
	// 	err = fmt.Errorf("Error setting formated_ibmid: %s", err)
	// 	return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_db2_saas_users", "read", "set-formated_ibmid").GetDiag()
	// }
	// if err = d.Set("iamid", successGetUserInfo.Iamid); err != nil {
	// 	err = fmt.Errorf("Error setting iamid: %s", err)
	// 	return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_db2_saas_users", "read", "set-iamid").GetDiag()
	// }
	// if err = d.Set("permitted_actions", successGetUserInfo.PermittedActions); err != nil {
	// 	err = fmt.Errorf("Error setting permitted_actions: %s", err)
	// 	return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_db2_saas_users", "read", "set-permitted_actions").GetDiag()
	// }
	// if err = d.Set("all_clean", successGetUserInfo.AllClean); err != nil {
	// 	err = fmt.Errorf("Error setting all_clean: %s", err)
	// 	return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_db2_saas_users", "read", "set-all_clean").GetDiag()
	// }
	// if err = d.Set("init_error_msg", successGetUserInfo.InitErrorMsg); err != nil {
	// 	err = fmt.Errorf("Error setting init_error_msg: %s", err)
	// 	return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_db2_saas_users", "read", "set-init_error_msg").GetDiag()
	// }

	return nil
}

func resourceIbmDb2SaasUsersUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	db2saasClient, err := meta.(conns.ClientSession).Db2saasV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_db2_saas_users", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	putDb2SaasUserOptions := &db2saasv1.PutDb2SaasUserOptions{}

	putDb2SaasUserOptions.SetID(d.Id())
	putDb2SaasUserOptions.SetXDeploymentID(d.Get("x_deployment_id").(string))
	putDb2SaasUserOptions.SetNewID(d.Get("id").(string))
	putDb2SaasUserOptions.SetNewIam(d.Get("iam").(bool))
	putDb2SaasUserOptions.SetNewIbmid(d.Get("ibmid").(string))
	putDb2SaasUserOptions.SetNewName(d.Get("name").(string))
	putDb2SaasUserOptions.SetNewPassword(d.Get("password").(string))
	putDb2SaasUserOptions.SetNewRole(d.Get("role").(string))
	putDb2SaasUserOptions.SetNewEmail(d.Get("email").(string))
	putDb2SaasUserOptions.SetNewLocked(d.Get("locked").(string))
	newAuthentication, err := ResourceIbmDb2SaasUsersMapToUpdateUserAuthentication(d.Get("authentication.0").(map[string]interface{}))
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_db2_saas_users", "update", "parse-authentication").GetDiag()
	}
	putDb2SaasUserOptions.SetNewAuthentication(newAuthentication)

	_, _, err = db2saasClient.PutDb2SaasUserWithContext(context, putDb2SaasUserOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("PutDb2SaasUserWithContext failed: %s", err.Error()), "ibm_db2_saas_users", "update")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	return resourceIbmDb2SaasUsersRead(context, d, meta)
}

func resourceIbmDb2SaasUsersDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	db2saasClient, err := meta.(conns.ClientSession).Db2saasV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_db2_saas_users", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteDb2SaasUserOptions := &db2saasv1.DeleteDb2SaasUserOptions{}

	deleteDb2SaasUserOptions.SetID(d.Id())
	deleteDb2SaasUserOptions.SetXDeploymentID(d.Get("x_deployment_id").(string))

	_, err = db2saasClient.DeleteDb2SaasUserWithContext(context, deleteDb2SaasUserOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteDb2SaasUserWithContext failed: %s", err.Error()), "ibm_db2_saas_users", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIbmDb2SaasUsersMapToCreateUserAuthentication(modelMap map[string]interface{}) (*db2saasv1.CreateUserAuthentication, error) {
	model := &db2saasv1.CreateUserAuthentication{}
	model.Method = core.StringPtr(modelMap["method"].(string))
	model.PolicyID = core.StringPtr(modelMap["policy_id"].(string))
	return model, nil
}

func ResourceIbmDb2SaasUsersMapToUpdateUserAuthentication(modelMap map[string]interface{}) (*db2saasv1.UpdateUserAuthentication, error) {
	model := &db2saasv1.UpdateUserAuthentication{}
	model.Method = core.StringPtr(modelMap["method"].(string))
	model.PolicyID = core.StringPtr(modelMap["policy_id"].(string))
	return model, nil
}

func ResourceIbmDb2SaasUsersSuccessUserResponseAuthenticationToMap(model *db2saasv1.SuccessUserResponseAuthentication) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["method"] = *model.Method
	modelMap["policy_id"] = *model.PolicyID
	return modelMap, nil
}
