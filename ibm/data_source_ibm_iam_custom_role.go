/**
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ibm

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
)

func dataSourceIbmIamCustomRole() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmIamCustomRoleRead,

		Schema: map[string]*schema.Schema{
			"role_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The role ID.",
			},
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The role ID.",
			},
			"display_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The display name of the role that is shown in the console.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The description of the role.",
			},
			"actions": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The actions of the role.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The role CRN.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the role that is used in the CRN. Can only be alphanumeric and has to be capitalized.",
			},
			"account_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The account GUID.",
			},
			"service_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The service name.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The UTC timestamp when the role was created.",
			},
			"created_by_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The iam ID of the entity that created the role.",
			},
			"last_modified_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The UTC timestamp when the role was last modified.",
			},
			"last_modified_by_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The iam ID of the entity that last modified the policy.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The href link back to the role.",
			},
		},
	}
}

func dataSourceIbmIamCustomRoleRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamPolicyManagementClient, err := meta.(ClientSession).IamPolicyManagementV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getRoleOptions := &iampolicymanagementv1.GetRoleOptions{}

	getRoleOptions.SetRoleID(d.Get("role_id").(string))

	customRole, response, err := iamPolicyManagementClient.GetRoleWithContext(context, getRoleOptions)
	if err != nil {
		log.Printf("[DEBUG] GetRoleWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	d.SetId(*customRole.ID)
	if err = d.Set("id", customRole.ID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting id: %s", err))
	}
	if err = d.Set("display_name", customRole.DisplayName); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting display_name: %s", err))
	}
	if err = d.Set("description", customRole.Description); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting description: %s", err))
	}
	if err = d.Set("actions", customRole.Actions); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting actions: %s", err))
	}
	if err = d.Set("crn", customRole.CRN); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting crn: %s", err))
	}
	if err = d.Set("name", customRole.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	if err = d.Set("account_id", customRole.AccountID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting account_id: %s", err))
	}
	if err = d.Set("service_name", customRole.ServiceName); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting service_name: %s", err))
	}
	if err = d.Set("created_at", customRole.CreatedAt); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
	}
	if err = d.Set("created_by_id", customRole.CreatedByID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_by_id: %s", err))
	}
	if err = d.Set("last_modified_at", customRole.LastModifiedAt); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting last_modified_at: %s", err))
	}
	if err = d.Set("last_modified_by_id", customRole.LastModifiedByID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting last_modified_by_id: %s", err))
	}
	if err = d.Set("href", customRole.Href); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting href: %s", err))
	}

	return nil
}
