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
	"github.com/IBM/cloud-db2-go-sdk/db2saasv1"
)

func DataSourceIbmDb2SaasUsersByID() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmDb2SaasUsersByIDRead,

		Schema: map[string]*schema.Schema{
			"x_deployment_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "CRN deployment id.",
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
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"formated_ibmid": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Formatted IBM ID.",
			},
			"role": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Role assigned to the user.",
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
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"all_clean": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates if the user account has no issues.",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "User's password.",
			},
			"iam": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates if IAM is enabled or not.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The display name of the user.",
			},
			"ibmid": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IBM ID of the user.",
			},
			"locked": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Account lock status for the user.",
			},
			"init_error_msg": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Initial error message.",
			},
			"email": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Email address of the user.",
			},
			"authentication": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Authentication details for the user.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"method": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Authentication method.",
						},
						"policy_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Policy ID of authentication.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmDb2SaasUsersByIDRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	db2saasClient, err := meta.(conns.ClientSession).Db2saasV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_db2_saas_users", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getbyidDb2SaasUserOptions := &db2saasv1.GetbyidDb2SaasUserOptions{}

	getbyidDb2SaasUserOptions.SetXDeploymentID(d.Get("x_deployment_id").(string))

	successGetUserByID, _, err := db2saasClient.GetbyidDb2SaasUserWithContext(context, getbyidDb2SaasUserOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetbyidDb2SaasUserWithContext failed: %s", err.Error()), "(Data) ibm_db2_saas_users", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*successGetUserByID.ID)

	if err = d.Set("dv_role", successGetUserByID.DvRole); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting dv_role: %s", err), "(Data) ibm_db2_saas_users", "read", "set-dv_role").GetDiag()
	}

	convertedMap := make(map[string]interface{}, len(successGetUserByID.Metadata))
	for k, v := range successGetUserByID.Metadata {
		convertedMap[k] = v
	}
	if err = d.Set("metadata", flex.Flatten(convertedMap)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting metadata: %s", err), "(Data) ibm_db2_saas_users", "read", "set-metadata").GetDiag()
	}

	if err = d.Set("formated_ibmid", successGetUserByID.FormatedIbmid); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting formated_ibmid: %s", err), "(Data) ibm_db2_saas_users", "read", "set-formated_ibmid").GetDiag()
	}

	if err = d.Set("role", successGetUserByID.Role); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting role: %s", err), "(Data) ibm_db2_saas_users", "read", "set-role").GetDiag()
	}

	if err = d.Set("iamid", successGetUserByID.Iamid); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting iamid: %s", err), "(Data) ibm_db2_saas_users", "read", "set-iamid").GetDiag()
	}

	permittedActions := []interface{}{}
	for _, permittedActionsItem := range successGetUserByID.PermittedActions {
		permittedActions = append(permittedActions, permittedActionsItem)
	}
	if err = d.Set("permitted_actions", permittedActions); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting permitted_actions: %s", err), "(Data) ibm_db2_saas_users", "read", "set-permitted_actions").GetDiag()
	}

	if err = d.Set("all_clean", successGetUserByID.AllClean); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting all_clean: %s", err), "(Data) ibm_db2_saas_users", "read", "set-all_clean").GetDiag()
	}

	if err = d.Set("password", successGetUserByID.Password); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting password: %s", err), "(Data) ibm_db2_saas_users", "read", "set-password").GetDiag()
	}

	if err = d.Set("iam", successGetUserByID.Iam); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting iam: %s", err), "(Data) ibm_db2_saas_users", "read", "set-iam").GetDiag()
	}

	if err = d.Set("name", successGetUserByID.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_db2_saas_users", "read", "set-name").GetDiag()
	}

	if err = d.Set("ibmid", successGetUserByID.Ibmid); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting ibmid: %s", err), "(Data) ibm_db2_saas_users", "read", "set-ibmid").GetDiag()
	}

	if err = d.Set("locked", successGetUserByID.Locked); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting locked: %s", err), "(Data) ibm_db2_saas_users", "read", "set-locked").GetDiag()
	}

	if err = d.Set("init_error_msg", successGetUserByID.InitErrorMsg); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting init_error_msg: %s", err), "(Data) ibm_db2_saas_users", "read", "set-init_error_msg").GetDiag()
	}

	if err = d.Set("email", successGetUserByID.Email); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting email: %s", err), "(Data) ibm_db2_saas_users", "read", "set-email").GetDiag()
	}

	authentication := []map[string]interface{}{}
	authenticationMap, err := DataSourceIbmDb2SaasUsersSuccessGetUserByIDAuthenticationToMap(successGetUserByID.Authentication)
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_db2_saas_users", "read", "authentication-to-map").GetDiag()
	}
	authentication = append(authentication, authenticationMap)
	if err = d.Set("authentication", authentication); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting authentication: %s", err), "(Data) ibm_db2_saas_users", "read", "set-authentication").GetDiag()
	}

	return nil
}

func DataSourceIbmDb2SaasUsersSuccessGetUserByIDAuthenticationToMap(model *db2saasv1.SuccessGetUserByIDAuthentication) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["method"] = *model.Method
	modelMap["policy_id"] = *model.PolicyID
	return modelMap, nil
}
