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

func dataSourceIbmIamPolicy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmIamPolicyRead,

		Schema: map[string]*schema.Schema{
			"policy_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The policy ID.",
			},
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The policy ID.",
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The policy type; either 'access' or 'authorization'.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Customer-defined description.",
			},
			"subjects": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The subjects associated with a policy.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attributes": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of subject attributes.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of an attribute.",
									},
									"value": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The value of an attribute.",
									},
								},
							},
						},
					},
				},
			},
			"roles": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A set of role cloud resource names (CRNs) granted by the policy.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"role_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The role cloud resource name granted by the policy.",
						},
						"display_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The display name of the role.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the role.",
						},
					},
				},
			},
			"resources": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The resources associated with a policy.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attributes": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of resource attributes.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of an attribute.",
									},
									"value": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The value of an attribute.",
									},
									"operator": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The operator of an attribute.",
									},
								},
							},
						},
					},
				},
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The href link back to the policy.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The UTC timestamp when the policy was created.",
			},
			"created_by_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The iam ID of the entity that created the policy.",
			},
			"last_modified_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The UTC timestamp when the policy was last modified.",
			},
			"last_modified_by_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The iam ID of the entity that last modified the policy.",
			},
		},
	}
}

func dataSourceIbmIamPolicyRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamPolicyManagementClient, err := meta.(ClientSession).IamPolicyManagementV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getPolicyOptions := &iampolicymanagementv1.GetPolicyOptions{}

	getPolicyOptions.SetPolicyID(d.Get("policy_id").(string))

	policy, response, err := iamPolicyManagementClient.GetPolicyWithContext(context, getPolicyOptions)
	if err != nil {
		log.Printf("[DEBUG] GetPolicyWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	d.SetId(*policy.ID)
	if err = d.Set("id", policy.ID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting id: %s", err))
	}
	if err = d.Set("type", policy.Type); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting type: %s", err))
	}
	if err = d.Set("description", policy.Description); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting description: %s", err))
	}

	if policy.Subjects != nil {
		err = d.Set("subjects", dataSourcePolicyFlattenSubjects(policy.Subjects))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting subjects %s", err))
		}
	}

	if policy.Roles != nil {
		err = d.Set("roles", dataSourcePolicyFlattenRoles(policy.Roles))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting roles %s", err))
		}
	}

	if policy.Resources != nil {
		err = d.Set("resources", dataSourcePolicyFlattenResources(policy.Resources))
		if err != nil {
			return diag.FromErr(fmt.Errorf("Error setting resources %s", err))
		}
	}
	if err = d.Set("href", policy.Href); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting href: %s", err))
	}
	if err = d.Set("created_at", policy.CreatedAt); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
	}
	if err = d.Set("created_by_id", policy.CreatedByID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_by_id: %s", err))
	}
	if err = d.Set("last_modified_at", policy.LastModifiedAt); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting last_modified_at: %s", err))
	}
	if err = d.Set("last_modified_by_id", policy.LastModifiedByID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting last_modified_by_id: %s", err))
	}

	return nil
}

func dataSourcePolicyFlattenSubjects(result []iampolicymanagementv1.PolicySubject) (subjects []map[string]interface{}) {
	for _, subjectsItem := range result {
		subjects = append(subjects, dataSourcePolicySubjectsToMap(subjectsItem))
	}

	return subjects
}

func dataSourcePolicySubjectsToMap(subjectsItem iampolicymanagementv1.PolicySubject) (subjectsMap map[string]interface{}) {
	subjectsMap = map[string]interface{}{}

	if subjectsItem.Attributes != nil {
		attributesList := []map[string]interface{}{}
		for _, attributesItem := range subjectsItem.Attributes {
			attributesList = append(attributesList, dataSourcePolicySubjectsAttributesToMap(attributesItem))
		}
		subjectsMap["attributes"] = attributesList
	}

	return subjectsMap
}

func dataSourcePolicySubjectsAttributesToMap(attributesItem iampolicymanagementv1.SubjectAttribute) (attributesMap map[string]interface{}) {
	attributesMap = map[string]interface{}{}

	if attributesItem.Name != nil {
		attributesMap["name"] = attributesItem.Name
	}
	if attributesItem.Value != nil {
		attributesMap["value"] = attributesItem.Value
	}

	return attributesMap
}

func dataSourcePolicyFlattenRoles(result []iampolicymanagementv1.PolicyRole) (roles []map[string]interface{}) {
	for _, rolesItem := range result {
		roles = append(roles, dataSourcePolicyRolesToMap(rolesItem))
	}

	return roles
}

func dataSourcePolicyRolesToMap(rolesItem iampolicymanagementv1.PolicyRole) (rolesMap map[string]interface{}) {
	rolesMap = map[string]interface{}{}

	if rolesItem.RoleID != nil {
		rolesMap["role_id"] = rolesItem.RoleID
	}
	if rolesItem.DisplayName != nil {
		rolesMap["display_name"] = rolesItem.DisplayName
	}
	if rolesItem.Description != nil {
		rolesMap["description"] = rolesItem.Description
	}

	return rolesMap
}

func dataSourcePolicyFlattenResources(result []iampolicymanagementv1.PolicyResource) (resources []map[string]interface{}) {
	for _, resourcesItem := range result {
		resources = append(resources, dataSourcePolicyResourcesToMap(resourcesItem))
	}

	return resources
}

func dataSourcePolicyResourcesToMap(resourcesItem iampolicymanagementv1.PolicyResource) (resourcesMap map[string]interface{}) {
	resourcesMap = map[string]interface{}{}

	if resourcesItem.Attributes != nil {
		attributesList := []map[string]interface{}{}
		for _, attributesItem := range resourcesItem.Attributes {
			attributesList = append(attributesList, dataSourcePolicyResourcesAttributesToMap(attributesItem))
		}
		resourcesMap["attributes"] = attributesList
	}

	return resourcesMap
}

func dataSourcePolicyResourcesAttributesToMap(attributesItem iampolicymanagementv1.ResourceAttribute) (attributesMap map[string]interface{}) {
	attributesMap = map[string]interface{}{}

	if attributesItem.Name != nil {
		attributesMap["name"] = attributesItem.Name
	}
	if attributesItem.Value != nil {
		attributesMap["value"] = attributesItem.Value
	}
	if attributesItem.Operator != nil {
		attributesMap["operator"] = attributesItem.Operator
	}

	return attributesMap
}
