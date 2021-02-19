/*
* IBM Confidential
* Object Code Only Source Materials
* 5747-SM3
* (c) Copyright IBM Corp. 2017,2021
*
* The source code for this program is not published or otherwise divested
* of its trade secrets, irrespective of what has been deposited with the
* U.S. Copyright Office.
 */

package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/IBM-Cloud/bluemix-go/api/iampap/iampapv1"
	"github.com/IBM-Cloud/bluemix-go/api/iampap/iampapv2"
	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/IBM-Cloud/bluemix-go/models"
	"github.com/IBM-Cloud/bluemix-go/utils"
)

func resourceIBMIAMAccessGroupPolicy() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMIAMAccessGroupPolicyCreate,
		Read:     resourceIBMIAMAccessGroupPolicyRead,
		Update:   resourceIBMIAMAccessGroupPolicyUpdate,
		Delete:   resourceIBMIAMAccessGroupPolicyDelete,
		Exists:   resourceIBMIAMAccessGroupPolicyExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"access_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of access group",
				ForceNew:    true,
			},

			"roles": {
				Type:        schema.TypeList,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Role names of the policy definition",
			},

			"resources": {
				Type:          schema.TypeList,
				Optional:      true,
				Computed:      true,
				MaxItems:      1,
				ConflictsWith: []string{"account_management"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Service name of the policy definition",
						},

						"resource_instance_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "ID of resource instance of the policy definition",
						},

						"region": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Region of the policy definition",
						},

						"resource_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Resource type of the policy definition",
						},

						"resource": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Resource of the policy definition",
						},

						"resource_group_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "ID of the resource group.",
						},

						"attributes": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Set resource attributes in the form of 'name=value,name=value....",
							Elem:        schema.TypeString,
						},
					},
				},
			},

			"account_management": {
				Type:          schema.TypeBool,
				Default:       false,
				Optional:      true,
				Description:   "Give access to all account management services",
				ConflictsWith: []string{"resources"},
			},

			"tags": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},

			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceIBMIAMAccessGroupPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	iampapClient, err := meta.(ClientSession).IAMPAPAPI()
	if err != nil {
		return err
	}

	accessgrpID := d.Get("access_group_id").(string)

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	var policy iampapv1.Policy

	policy, err = generateAccountPolicyV2(d, meta)
	if err != nil {
		return err
	}

	policy.Subjects = []iampapv1.Subject{
		{
			Attributes: []iampapv1.Attribute{
				{
					Name:  "access_group_id",
					Value: accessgrpID,
				},
			},
		},
	}

	policy.Type = iampapv1.AccessPolicyType

	policy.Resources[0].SetAccountID(userDetails.userAccount)

	accgrpPolicy, err := iampapClient.V1Policy().Create(policy)

	if err != nil {
		return fmt.Errorf("Error creating access group policy: %s", err)
	}

	d.SetId(fmt.Sprintf("%s/%s", accessgrpID, accgrpPolicy.ID))

	return resourceIBMIAMAccessGroupPolicyRead(d, meta)
}

func resourceIBMIAMAccessGroupPolicyRead(d *schema.ResourceData, meta interface{}) error {

	iampapClient, err := meta.(ClientSession).IAMPAPAPI()
	if err != nil {
		return err
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	accessgrpID := parts[0]
	accgrpPolicyID := parts[1]

	accgrpPolicy, err := iampapClient.V1Policy().Get(accgrpPolicyID)
	if err != nil {
		return fmt.Errorf("Error retrieving access group policy: %s", err)
	}

	if accessgrpID != accgrpPolicy.Subjects[0].GetAttribute("access_group_id") {
		return fmt.Errorf("Policy %s does not belong to access group %s", accgrpPolicyID, accessgrpID)
	}

	d.Set("access_group_id", accessgrpID)
	roles := make([]string, len(accgrpPolicy.Roles))
	for i, role := range accgrpPolicy.Roles {
		roles[i] = role.Name
	}
	d.Set("roles", roles)
	d.Set("version", accgrpPolicy.Version)
	d.Set("resources", flattenPolicyResource(accgrpPolicy.Resources))
	if len(accgrpPolicy.Resources) > 0 {
		if accgrpPolicy.Resources[0].GetAttribute("serviceType") == "service" {
			d.Set("account_management", false)
		}
		if accgrpPolicy.Resources[0].GetAttribute("serviceType") == "platform_service" {
			d.Set("account_management", true)
		}
	}

	return nil
}

func resourceIBMIAMAccessGroupPolicyUpdate(d *schema.ResourceData, meta interface{}) error {

	iampapClient, err := meta.(ClientSession).IAMPAPAPI()
	if err != nil {
		return err
	}
	if d.HasChange("roles") || d.HasChange("resources") || d.HasChange("account_management") {
		parts, err := idParts(d.Id())
		if err != nil {
			return err
		}
		accessgrpID := parts[0]
		accgrpPolicyID := parts[1]

		userDetails, err := meta.(ClientSession).BluemixUserDetails()
		if err != nil {
			return err
		}

		var policy iampapv1.Policy

		policy, err = generateAccountPolicyV2(d, meta)
		if err != nil {
			return err
		}

		policy.Subjects = []iampapv1.Subject{
			{
				Attributes: []iampapv1.Attribute{
					{
						Name:  "access_group_id",
						Value: accessgrpID,
					},
				},
			},
		}

		policy.Type = iampapv1.AccessPolicyType

		policy.Resources[0].SetAccountID(userDetails.userAccount)

		_, err = iampapClient.V1Policy().Update(accgrpPolicyID, policy, d.Get("version").(string))
		if err != nil {
			return fmt.Errorf("Error updating access group policy: %s", err)
		}

	}

	return resourceIBMIAMAccessGroupPolicyRead(d, meta)

}

func resourceIBMIAMAccessGroupPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	iampapClient, err := meta.(ClientSession).IAMPAPAPI()
	if err != nil {
		return err
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	accgrpPolicyID := parts[1]

	err = iampapClient.V1Policy().Delete(accgrpPolicyID)
	if err != nil {
		return fmt.Errorf("Error deleting access group policy: %s", err)
	}

	d.SetId("")

	return nil
}

func resourceIBMIAMAccessGroupPolicyExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	iampapClient, err := meta.(ClientSession).IAMPAPAPI()
	if err != nil {
		return false, err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return false, err
	}

	accgrpPolicyID := parts[1]

	accgrpPolicy, err := iampapClient.V1Policy().Get(accgrpPolicyID)
	if err != nil {
		if apiErr, ok := err.(bmxerror.RequestFailure); ok {
			if apiErr.StatusCode() == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error communicating with the API: %s", err)
	}

	tempID := fmt.Sprintf("%s/%s", accgrpPolicy.Subjects[0].GetAttribute("access_group_id"), accgrpPolicy.ID)

	return tempID == d.Id(), nil
}

func generateAccountPolicy(d *schema.ResourceData, meta interface{}) (iampapv1.Policy, error) {

	var serviceName string
	policyResource := iampapv1.Resource{}

	if res, ok := d.GetOk("resources"); ok {
		resources := res.([]interface{})
		for _, resource := range resources {
			r, _ := resource.(map[string]interface{})
			serviceName = r["service"].(string)
			if r, ok := r["service"]; ok {
				if r.(string) != "" {
					policyResource.SetServiceName(r.(string))
				}
			}
			if r, ok := r["resource_instance_id"]; ok {
				if r.(string) != "" {
					policyResource.SetServiceInstance(r.(string))
				}

			}
			if r, ok := r["region"]; ok {
				if r.(string) != "" {
					policyResource.SetRegion(r.(string))
				}

			}
			if r, ok := r["resource_type"]; ok {
				if r.(string) != "" {
					policyResource.SetResourceType(r.(string))
				}

			}
			if r, ok := r["resource"]; ok {
				if r.(string) != "" {
					policyResource.SetResource(r.(string))
				}

			}
			if r, ok := r["resource_group_id"]; ok {
				if r.(string) != "" {
					policyResource.SetResourceGroupID(r.(string))
				}

			}
			if r, ok := r["attributes"]; ok {
				for k, v := range r.(map[string]interface{}) {
					policyResource.SetAttribute(k, v.(string))
				}

			}

		}
	}

	if d.Get("account_management").(bool) {
		policyResource.SetServiceType("platform_service")
	}

	if len(policyResource.Attributes) == 0 {
		policyResource.SetServiceType("service")
	}

	iamClient, err := meta.(ClientSession).IAMAPI()
	if err != nil {
		return iampapv1.Policy{}, err
	}

	iamRepo := iamClient.ServiceRoles()

	var roles []models.PolicyRole

	if serviceName == "" {
		roles, err = iamRepo.ListSystemDefinedRoles()
	} else {
		roles, err = iamRepo.ListServiceRoles(serviceName)
	}
	if err != nil {
		return iampapv1.Policy{}, err
	}

	policyRoles, err := getRolesFromRoleNames(expandStringList(d.Get("roles").([]interface{})), roles)
	if err != nil {
		return iampapv1.Policy{}, err
	}

	return iampapv1.Policy{Roles: iampapv1.ConvertRoleModels(policyRoles), Resources: []iampapv1.Resource{policyResource}}, nil
}

func generateAccountPolicyV2(d *schema.ResourceData, meta interface{}) (iampapv1.Policy, error) {

	var serviceName string
	var resourceType string
	policyResource := iampapv1.Resource{}

	if res, ok := d.GetOk("resources"); ok {
		resources := res.([]interface{})
		for _, resource := range resources {
			r, _ := resource.(map[string]interface{})
			if r, ok := r["service"]; ok && r != nil {
				serviceName = r.(string)
				if r.(string) != "" {
					policyResource.SetServiceName(r.(string))
				}
			}
			if r, ok := r["resource_instance_id"]; ok {
				if r.(string) != "" {
					resourceType = r.(string)
					policyResource.SetServiceInstance(r.(string))
				}

			}
			if r, ok := r["region"]; ok {
				if r.(string) != "" {
					policyResource.SetRegion(r.(string))
				}

			}
			if r, ok := r["resource_type"]; ok {
				if r.(string) != "" {
					policyResource.SetResourceType(r.(string))
				}

			}
			if r, ok := r["resource"]; ok {
				if r.(string) != "" {
					policyResource.SetResource(r.(string))
				}

			}
			if r, ok := r["resource_group_id"]; ok {
				if r.(string) != "" {
					policyResource.SetResourceGroupID(r.(string))
				}

			}
			if r, ok := r["attributes"]; ok {
				for k, v := range r.(map[string]interface{}) {
					policyResource.SetAttribute(k, v.(string))
				}

			}

		}
	}

	if d.Get("account_management").(bool) {
		policyResource.SetServiceType("platform_service")
	}

	if len(policyResource.Attributes) == 0 {
		policyResource.SetServiceType("service")
	}

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return iampapv1.Policy{}, err
	}

	iamClient, err := meta.(ClientSession).IAMPAPAPIV2()
	if err != nil {
		return iampapv1.Policy{}, err
	}

	iamRepo := iamClient.IAMRoles()

	var roles []iampapv2.Role

	serviceToQuery := serviceName

	if serviceName == "" && // no specific service specified
		!d.Get("account_management").(bool) && // not all account management services
		resourceType != "resource-group" { // not to a resource group
		serviceToQuery = "alliamserviceroles"
	}

	query := iampapv2.RoleQuery{
		AccountID:   userDetails.userAccount,
		ServiceName: serviceToQuery,
	}

	roles, err = iamRepo.ListAll(query)

	policyRoles, err := utils.GetRolesFromRoleNamesV2(expandStringList(d.Get("roles").([]interface{})), roles)
	if err != nil {
		return iampapv1.Policy{}, err
	}

	return iampapv1.Policy{Roles: iampapv1.ConvertV2RoleModels(policyRoles), Resources: []iampapv1.Resource{policyResource}}, nil
}
