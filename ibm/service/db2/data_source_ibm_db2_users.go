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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/cloud-db2-go-sdk/db2saasv1"
)

func DataSourceIbmDb2SaasUsers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmDb2SaasUsersRead,

		Schema: map[string]*schema.Schema{
			"x_deployment_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "CRN deployment id.",
			},
			"count": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total number of resources.",
			},
			"resources": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of user resource.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique identifier for the user.",
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
				},
			},
		},
	}
}

func dataSourceIbmDb2SaasUsersRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	db2saasClient, err := meta.(conns.ClientSession).Db2saasV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_db2_saas_users", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getDb2SaasUserOptions := &db2saasv1.GetDb2SaasUserOptions{}

	getDb2SaasUserOptions.SetXDeploymentID(d.Get("x_deployment_id").(string))

	successGetUserInfo, _, err := db2saasClient.GetDb2SaasUserWithContext(context, getDb2SaasUserOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDb2SaasUserWithContext failed: %s", err.Error()), "(Data) ibm_db2_saas_users", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmDb2SaasUsersID(d))

	if err = d.Set("count", flex.IntValue(successGetUserInfo.Count)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting count: %s", err), "(Data) ibm_db2_saas_users", "read", "set-count").GetDiag()
	}

	resources := []map[string]interface{}{}
	for _, resourcesItem := range successGetUserInfo.Resources {
		resourcesItemMap, err := DataSourceIbmDb2SaasUsersSuccessGetUserInfoResourcesItemToMap(&resourcesItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_db2_saas_users", "read", "resources-to-map").GetDiag()
		}
		resources = append(resources, resourcesItemMap)
	}
	if err = d.Set("resources", resources); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resources: %s", err), "(Data) ibm_db2_saas_users", "read", "set-resources").GetDiag()
	}

	return nil
}

// dataSourceIbmDb2SaasUsersID returns a reasonable ID for the list.
func dataSourceIbmDb2SaasUsersID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmDb2SaasUsersSuccessGetUserInfoResourcesItemToMap(model *db2saasv1.SuccessGetUserInfoResourcesItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.DvRole != nil {
		modelMap["dv_role"] = *model.DvRole
	}
	if model.Metadata != nil {
		metadata := make(map[string]interface{})
		for k, v := range model.Metadata {
			metadata[k] = flex.Stringify(v)
		}
		modelMap["metadata"] = metadata
	}
	if model.FormatedIbmid != nil {
		modelMap["formated_ibmid"] = *model.FormatedIbmid
	}
	if model.Role != nil {
		modelMap["role"] = *model.Role
	}
	if model.Iamid != nil {
		modelMap["iamid"] = *model.Iamid
	}
	if model.PermittedActions != nil {
		modelMap["permitted_actions"] = model.PermittedActions
	}
	if model.AllClean != nil {
		modelMap["all_clean"] = *model.AllClean
	}
	if model.Password != nil {
		modelMap["password"] = *model.Password
	}
	if model.Iam != nil {
		modelMap["iam"] = *model.Iam
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Ibmid != nil {
		modelMap["ibmid"] = *model.Ibmid
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Locked != nil {
		modelMap["locked"] = *model.Locked
	}
	if model.InitErrorMsg != nil {
		modelMap["init_error_msg"] = *model.InitErrorMsg
	}
	if model.Email != nil {
		modelMap["email"] = *model.Email
	}
	if model.Authentication != nil {
		authenticationMap, err := DataSourceIbmDb2SaasUsersSuccessGetUserInfoResourcesItemAuthenticationToMap(model.Authentication)
		if err != nil {
			return modelMap, err
		}
		modelMap["authentication"] = []map[string]interface{}{authenticationMap}
	}
	return modelMap, nil
}

func DataSourceIbmDb2SaasUsersSuccessGetUserInfoResourcesItemAuthenticationToMap(model *db2saasv1.SuccessGetUserInfoResourcesItemAuthentication) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["method"] = *model.Method
	modelMap["policy_id"] = *model.PolicyID
	return modelMap, nil
}
