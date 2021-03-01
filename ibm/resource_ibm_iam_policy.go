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

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
)

func resourceIbmIamPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmIamPolicyCreate,
		ReadContext:   resourceIbmIamPolicyRead,
		UpdateContext: resourceIbmIamPolicyUpdate,
		DeleteContext: resourceIbmIamPolicyDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The policy type; either 'access' or 'authorization'.",
			},
			"subjects": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				Description: "The subjects associated with a policy.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attributes": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of subject attributes.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The name of an attribute.",
									},
									"value": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
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
				Required:    true,
				Description: "A set of role cloud resource names (CRNs) granted by the policy.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"role_id": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The role cloud resource name granted by the policy.",
						},
						"display_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The display name of the role.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The description of the role.",
						},
					},
				},
			},
			"resources": &schema.Schema{
				Type:        schema.TypeList,
				Required:    true,
				Description: "The resources associated with a policy.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attributes": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of resource attributes.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The name of an attribute.",
									},
									"value": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The value of an attribute.",
									},
									"operator": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The operator of an attribute.",
									},
								},
							},
						},
					},
				},
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Customer-defined description.",
			},
			"accept_language": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Translation language code.",
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

func resourceIbmIamPolicyCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamPolicyManagementClient, err := meta.(ClientSession).IamPolicyManagementV1()
	if err != nil {
		return diag.FromErr(err)
	}

	createPolicyOptions := &iampolicymanagementv1.CreatePolicyOptions{}

	createPolicyOptions.SetType(d.Get("type").(string))
	var subjects []iampolicymanagementv1.PolicySubject
	for _, e := range d.Get("subjects").([]interface{}) {
		value := e.(map[string]interface{})
		subjectsItem := resourceIbmIamPolicyMapToPolicySubject(value)
		subjects = append(subjects, subjectsItem)
	}
	createPolicyOptions.SetSubjects(subjects)
	var roles []iampolicymanagementv1.PolicyRole
	for _, e := range d.Get("roles").([]interface{}) {
		value := e.(map[string]interface{})
		rolesItem := resourceIbmIamPolicyMapToPolicyRole(value)
		roles = append(roles, rolesItem)
	}
	createPolicyOptions.SetRoles(roles)
	var resources []iampolicymanagementv1.PolicyResource
	for _, e := range d.Get("resources").([]interface{}) {
		value := e.(map[string]interface{})
		resourcesItem := resourceIbmIamPolicyMapToPolicyResource(value)
		resources = append(resources, resourcesItem)
	}
	createPolicyOptions.SetResources(resources)
	if _, ok := d.GetOk("description"); ok {
		createPolicyOptions.SetDescription(d.Get("description").(string))
	}
	if _, ok := d.GetOk("accept_language"); ok {
		createPolicyOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}

	policy, response, err := iamPolicyManagementClient.CreatePolicyWithContext(context, createPolicyOptions)
	if err != nil {
		log.Printf("[DEBUG] CreatePolicyWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	d.SetId(*policy.ID)

	return resourceIbmIamPolicyRead(context, d, meta)
}

func resourceIbmIamPolicyMapToPolicySubject(policySubjectMap map[string]interface{}) iampolicymanagementv1.PolicySubject {
	policySubject := iampolicymanagementv1.PolicySubject{}

	if policySubjectMap["attributes"] != nil {
		attributes := []iampolicymanagementv1.SubjectAttribute{}
		for _, attributesItem := range policySubjectMap["attributes"].([]interface{}) {
			attributesItemModel := resourceIbmIamPolicyMapToSubjectAttribute(attributesItem.(map[string]interface{}))
			attributes = append(attributes, attributesItemModel)
		}
		policySubject.Attributes = attributes
	}

	return policySubject
}

func resourceIbmIamPolicyMapToSubjectAttribute(subjectAttributeMap map[string]interface{}) iampolicymanagementv1.SubjectAttribute {
	subjectAttribute := iampolicymanagementv1.SubjectAttribute{}

	subjectAttribute.Name = core.StringPtr(subjectAttributeMap["name"].(string))
	subjectAttribute.Value = core.StringPtr(subjectAttributeMap["value"].(string))

	return subjectAttribute
}

func resourceIbmIamPolicyMapToPolicyRole(policyRoleMap map[string]interface{}) iampolicymanagementv1.PolicyRole {
	policyRole := iampolicymanagementv1.PolicyRole{}

	policyRole.RoleID = core.StringPtr(policyRoleMap["role_id"].(string))
	if policyRoleMap["display_name"] != nil {
		policyRole.DisplayName = core.StringPtr(policyRoleMap["display_name"].(string))
	}
	if policyRoleMap["description"] != nil {
		policyRole.Description = core.StringPtr(policyRoleMap["description"].(string))
	}

	return policyRole
}

func resourceIbmIamPolicyMapToPolicyResource(policyResourceMap map[string]interface{}) iampolicymanagementv1.PolicyResource {
	policyResource := iampolicymanagementv1.PolicyResource{}

	if policyResourceMap["attributes"] != nil {
		attributes := []iampolicymanagementv1.ResourceAttribute{}
		for _, attributesItem := range policyResourceMap["attributes"].([]interface{}) {
			attributesItemModel := resourceIbmIamPolicyMapToResourceAttribute(attributesItem.(map[string]interface{}))
			attributes = append(attributes, attributesItemModel)
		}
		policyResource.Attributes = attributes
	}

	return policyResource
}

func resourceIbmIamPolicyMapToResourceAttribute(resourceAttributeMap map[string]interface{}) iampolicymanagementv1.ResourceAttribute {
	resourceAttribute := iampolicymanagementv1.ResourceAttribute{}

	resourceAttribute.Name = core.StringPtr(resourceAttributeMap["name"].(string))
	resourceAttribute.Value = core.StringPtr(resourceAttributeMap["value"].(string))
	if resourceAttributeMap["operator"] != nil {
		resourceAttribute.Operator = core.StringPtr(resourceAttributeMap["operator"].(string))
	}

	return resourceAttribute
}

func resourceIbmIamPolicyRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamPolicyManagementClient, err := meta.(ClientSession).IamPolicyManagementV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getPolicyOptions := &iampolicymanagementv1.GetPolicyOptions{}

	getPolicyOptions.SetPolicyID(d.Id())

	policy, response, err := iamPolicyManagementClient.GetPolicyWithContext(context, getPolicyOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetPolicyWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	if err = d.Set("type", policy.Type); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting type: %s", err))
	}
	subjects := []map[string]interface{}{}
	for _, subjectsItem := range policy.Subjects {
		subjectsItemMap := resourceIbmIamPolicyPolicySubjectToMap(subjectsItem)
		subjects = append(subjects, subjectsItemMap)
	}
	if err = d.Set("subjects", subjects); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting subjects: %s", err))
	}
	roles := []map[string]interface{}{}
	for _, rolesItem := range policy.Roles {
		rolesItemMap := resourceIbmIamPolicyPolicyRoleToMap(rolesItem)
		roles = append(roles, rolesItemMap)
	}
	if err = d.Set("roles", roles); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting roles: %s", err))
	}
	resources := []map[string]interface{}{}
	for _, resourcesItem := range policy.Resources {
		resourcesItemMap := resourceIbmIamPolicyPolicyResourceToMap(resourcesItem)
		resources = append(resources, resourcesItemMap)
	}
	if err = d.Set("resources", resources); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting resources: %s", err))
	}
	if err = d.Set("description", policy.Description); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting description: %s", err))
	}
	// if err = d.Set("accept_language", policy.AcceptLanguage); err != nil {
	// 	return diag.FromErr(fmt.Errorf("Error setting accept_language: %s", err))
	// }
	if err = d.Set("href", policy.Href); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting href: %s", err))
	}
	if err = d.Set("created_at", policy.CreatedAt.String()); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
	}
	if err = d.Set("created_by_id", policy.CreatedByID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_by_id: %s", err))
	}
	if err = d.Set("last_modified_at", policy.LastModifiedAt.String()); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting last_modified_at: %s", err))
	}
	if err = d.Set("last_modified_by_id", policy.LastModifiedByID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting last_modified_by_id: %s", err))
	}

	return nil
}

func resourceIbmIamPolicyPolicySubjectToMap(policySubject iampolicymanagementv1.PolicySubject) map[string]interface{} {
	policySubjectMap := map[string]interface{}{}

	if policySubject.Attributes != nil {
		attributes := []map[string]interface{}{}
		for _, attributesItem := range policySubject.Attributes {
			attributesItemMap := resourceIbmIamPolicySubjectAttributeToMap(attributesItem)
			attributes = append(attributes, attributesItemMap)
			// TODO: handle Attributes of type TypeList -- list of non-primitive, not model items
		}
		policySubjectMap["attributes"] = attributes
	}

	return policySubjectMap
}

func resourceIbmIamPolicySubjectAttributeToMap(subjectAttribute iampolicymanagementv1.SubjectAttribute) map[string]interface{} {
	subjectAttributeMap := map[string]interface{}{}

	subjectAttributeMap["name"] = subjectAttribute.Name
	subjectAttributeMap["value"] = subjectAttribute.Value

	return subjectAttributeMap
}

func resourceIbmIamPolicyPolicyRoleToMap(policyRole iampolicymanagementv1.PolicyRole) map[string]interface{} {
	policyRoleMap := map[string]interface{}{}

	policyRoleMap["role_id"] = policyRole.RoleID
	policyRoleMap["display_name"] = policyRole.DisplayName
	policyRoleMap["description"] = policyRole.Description

	return policyRoleMap
}

func resourceIbmIamPolicyPolicyResourceToMap(policyResource iampolicymanagementv1.PolicyResource) map[string]interface{} {
	policyResourceMap := map[string]interface{}{}

	if policyResource.Attributes != nil {
		attributes := []map[string]interface{}{}
		for _, attributesItem := range policyResource.Attributes {
			attributesItemMap := resourceIbmIamPolicyResourceAttributeToMap(attributesItem)
			attributes = append(attributes, attributesItemMap)
			// TODO: handle Attributes of type TypeList -- list of non-primitive, not model items
		}
		policyResourceMap["attributes"] = attributes
	}

	return policyResourceMap
}

func resourceIbmIamPolicyResourceAttributeToMap(resourceAttribute iampolicymanagementv1.ResourceAttribute) map[string]interface{} {
	resourceAttributeMap := map[string]interface{}{}

	resourceAttributeMap["name"] = resourceAttribute.Name
	resourceAttributeMap["value"] = resourceAttribute.Value
	resourceAttributeMap["operator"] = resourceAttribute.Operator

	return resourceAttributeMap
}

func resourceIbmIamPolicyUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamPolicyManagementClient, err := meta.(ClientSession).IamPolicyManagementV1()
	if err != nil {
		return diag.FromErr(err)
	}

	updatePolicyOptions := &iampolicymanagementv1.UpdatePolicyOptions{}

	updatePolicyOptions.SetPolicyID(d.Id())
	updatePolicyOptions.SetType(d.Get("type").(string))
	var subjects []iampolicymanagementv1.PolicySubject
	for _, e := range d.Get("subjects").([]interface{}) {
		value := e.(map[string]interface{})
		subjectsItem := resourceIbmIamPolicyMapToPolicySubject(value)
		subjects = append(subjects, subjectsItem)
	}
	updatePolicyOptions.SetSubjects(subjects)
	var roles []iampolicymanagementv1.PolicyRole
	for _, e := range d.Get("roles").([]interface{}) {
		value := e.(map[string]interface{})
		rolesItem := resourceIbmIamPolicyMapToPolicyRole(value)
		roles = append(roles, rolesItem)
	}
	updatePolicyOptions.SetRoles(roles)
	var resources []iampolicymanagementv1.PolicyResource
	for _, e := range d.Get("resources").([]interface{}) {
		value := e.(map[string]interface{})
		resourcesItem := resourceIbmIamPolicyMapToPolicyResource(value)
		resources = append(resources, resourcesItem)
	}
	updatePolicyOptions.SetResources(resources)
	if _, ok := d.GetOk("description"); ok {
		updatePolicyOptions.SetDescription(d.Get("description").(string))
	}
	// if _, ok := d.GetOk("accept_language"); ok {
	// 	updatePolicyOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	// }

	_, response, err := iamPolicyManagementClient.UpdatePolicyWithContext(context, updatePolicyOptions)
	if err != nil {
		log.Printf("[DEBUG] UpdatePolicyWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	return resourceIbmIamPolicyRead(context, d, meta)
}

func resourceIbmIamPolicyDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamPolicyManagementClient, err := meta.(ClientSession).IamPolicyManagementV1()
	if err != nil {
		return diag.FromErr(err)
	}

	deletePolicyOptions := &iampolicymanagementv1.DeletePolicyOptions{}

	deletePolicyOptions.SetPolicyID(d.Id())

	response, err := iamPolicyManagementClient.DeletePolicyWithContext(context, deletePolicyOptions)
	if err != nil {
		log.Printf("[DEBUG] DeletePolicyWithContext failed %s\n%s", err, response)
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}
