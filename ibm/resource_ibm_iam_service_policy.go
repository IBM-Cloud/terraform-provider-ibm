package ibm

import (
	"fmt"

	"github.com/IBM-Cloud/bluemix-go/api/iam/iamv1"
	"github.com/IBM-Cloud/bluemix-go/models"

	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceIBMIAMServicePolicy() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMIAMServicePolicyCreate,
		Read:     resourceIBMIAMServicePolicyRead,
		Update:   resourceIBMIAMServicePolicyUpdate,
		Delete:   resourceIBMIAMServicePolicyDelete,
		Exists:   resourceIBMIAMServicePolicyExists,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"iam_service_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "UUID of ServiceID",
				ForceNew:    true,
			},

			"roles": {
				Type:        schema.TypeList,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Role names of the policy definition",
			},

			"resources": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
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
					},
				},
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

func resourceIBMIAMServicePolicyCreate(d *schema.ResourceData, meta interface{}) error {
	iamClient, err := meta.(ClientSession).IAMAPI()
	if err != nil {
		return err
	}
	serviceIDUUID := d.Get("iam_service_id").(string)

	bmxSess, err := meta.(ClientSession).BluemixSession()
	if err != nil {
		return err
	}

	mccpAPI, err := meta.(ClientSession).MccpAPI()
	if err != nil {
		return err
	}
	region, err := mccpAPI.Regions().FindRegionByName(bmxSess.Config.Region)
	if err != nil {
		return err
	}

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	boundTo := GenerateBoundToCRN(*region, userDetails.userAccount)
	var policy models.Policy

	policy, err = generatePolicy(d, meta, userDetails.userAccount)
	if err != nil {
		return err
	}

	serviceID, err := iamClient.ServiceIds().Get(serviceIDUUID)
	if err != nil {
		return err
	}

	servicePolicy, err := iamClient.ServicePolicies().Create(boundTo.ScopeSegment(), serviceID.IAMID, policy)

	if err != nil {
		return fmt.Errorf("Error creating servicePolicy: %s", err)
	}

	d.SetId(fmt.Sprintf("%s/%s", serviceIDUUID, servicePolicy.ID))

	return resourceIBMIAMServicePolicyRead(d, meta)
}

func resourceIBMIAMServicePolicyRead(d *schema.ResourceData, meta interface{}) error {

	iamClient, err := meta.(ClientSession).IAMAPI()
	if err != nil {
		return err
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	serviceIDUUID := parts[0]
	servicePolicyID := parts[1]

	bmxSess, err := meta.(ClientSession).BluemixSession()
	if err != nil {
		return err
	}

	mccpAPI, err := meta.(ClientSession).MccpAPI()
	if err != nil {
		return err
	}
	region, err := mccpAPI.Regions().FindRegionByName(bmxSess.Config.Region)
	if err != nil {
		return err
	}

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	boundTo := GenerateBoundToCRN(*region, userDetails.userAccount)

	serviceID, err := iamClient.ServiceIds().Get(serviceIDUUID)
	if err != nil {
		return err
	}

	servicePolicy, err := iamClient.ServicePolicies().Get(boundTo.ScopeSegment(), serviceID.IAMID, servicePolicyID)
	if err != nil {
		return fmt.Errorf("Error retrieving servicePolicy: %s", err)
	}

	d.Set("iam_service_id", serviceIDUUID)
	roles := make([]string, len(servicePolicy.Roles))
	for i, role := range servicePolicy.Roles {
		roles[i] = role.DisplayName
	}
	d.Set("roles", roles)
	d.Set("version", servicePolicy.Version)
	d.Set("resources", flattenPolicyResource(servicePolicy.Resources))

	return nil
}

func resourceIBMIAMServicePolicyUpdate(d *schema.ResourceData, meta interface{}) error {

	iamClient, err := meta.(ClientSession).IAMAPI()
	if err != nil {
		return err
	}
	if d.HasChange("roles") || d.HasChange("resources") {
		parts, err := idParts(d.Id())
		if err != nil {
			return err
		}
		serviceIDUUID := parts[0]
		servicePolicyID := parts[1]

		bmxSess, err := meta.(ClientSession).BluemixSession()
		if err != nil {
			return err
		}

		mccpAPI, err := meta.(ClientSession).MccpAPI()
		if err != nil {
			return err
		}
		region, err := mccpAPI.Regions().FindRegionByName(bmxSess.Config.Region)
		if err != nil {
			return err
		}

		userDetails, err := meta.(ClientSession).BluemixUserDetails()
		if err != nil {
			return err
		}

		boundTo := GenerateBoundToCRN(*region, userDetails.userAccount)
		var policy models.Policy

		policy, err = generatePolicy(d, meta, userDetails.userAccount)
		if err != nil {
			return err
		}

		serviceID, err := iamClient.ServiceIds().Get(serviceIDUUID)
		if err != nil {
			return err
		}

		_, err = iamClient.ServicePolicies().Update(iamv1.ServicePolicyIdentifier{
			Scope:    boundTo.ScopeSegment(),
			IAMID:    serviceID.IAMID,
			PolicyID: servicePolicyID,
		}, policy, d.Get("version").(string))
		if err != nil {
			return fmt.Errorf("Error updating service policy: %s", err)
		}

	}

	return resourceIBMIAMServicePolicyRead(d, meta)

}

func resourceIBMIAMServicePolicyDelete(d *schema.ResourceData, meta interface{}) error {
	iamClient, err := meta.(ClientSession).IAMAPI()
	if err != nil {
		return err
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	serviceIDUUID := parts[0]
	servicePolicyID := parts[1]

	bmxSess, err := meta.(ClientSession).BluemixSession()
	if err != nil {
		return err
	}

	mccpAPI, err := meta.(ClientSession).MccpAPI()
	if err != nil {
		return err
	}
	region, err := mccpAPI.Regions().FindRegionByName(bmxSess.Config.Region)
	if err != nil {
		return err
	}

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	boundTo := GenerateBoundToCRN(*region, userDetails.userAccount)

	serviceID, err := iamClient.ServiceIds().Get(serviceIDUUID)
	if err != nil {
		return err
	}

	err = iamClient.ServicePolicies().Delete(iamv1.ServicePolicyIdentifier{
		Scope:    boundTo.ScopeSegment(),
		IAMID:    serviceID.IAMID,
		PolicyID: servicePolicyID,
	})
	if err != nil {
		return fmt.Errorf("Error deleting service policy: %s", err)
	}

	d.SetId("")

	return nil
}

func resourceIBMIAMServicePolicyExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	iamClient, err := meta.(ClientSession).IAMAPI()
	if err != nil {
		return false, err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return false, err
	}
	serviceIDUUID := parts[0]
	servicePolicyID := parts[1]

	bmxSess, err := meta.(ClientSession).BluemixSession()
	if err != nil {
		return false, err
	}

	mccpAPI, err := meta.(ClientSession).MccpAPI()
	if err != nil {
		return false, err
	}
	region, err := mccpAPI.Regions().FindRegionByName(bmxSess.Config.Region)
	if err != nil {
		return false, err
	}

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return false, err
	}

	boundTo := GenerateBoundToCRN(*region, userDetails.userAccount)

	serviceID, err := iamClient.ServiceIds().Get(serviceIDUUID)
	if err != nil {
		return false, err
	}

	servicePolicy, err := iamClient.ServicePolicies().Get(boundTo.ScopeSegment(), serviceID.IAMID, servicePolicyID)
	if err != nil {
		return false, fmt.Errorf("Error retrieving servicePolicy: %s", err)
	}
	if err != nil {
		if apiErr, ok := err.(bmxerror.RequestFailure); ok {
			if apiErr.StatusCode() == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error communicating with the API: %s", err)
	}

	tempID := fmt.Sprintf("%s/%s", serviceIDUUID, servicePolicy.ID)

	return tempID == d.Id(), nil
}

func generatePolicy(d *schema.ResourceData, meta interface{}, accountID string) (models.Policy, error) {

	policyResources := []models.PolicyResource{}
	var resources []interface{}
	var serviceName string

	if res, ok := d.GetOk("resources"); ok {
		resources = res.([]interface{})
		for _, resource := range resources {
			r, _ := resource.(map[string]interface{})
			serviceName = r["service"].(string)
			resourceParam := models.PolicyResource{
				ServiceName:     r["service"].(string),
				ServiceInstance: r["resource_instance_id"].(string),
				Region:          r["region"].(string),
				ResourceType:    r["resource_type"].(string),
				Resource:        r["resource"].(string),
				AccountID:       accountID,
				ResourceGroupID: r["resource_group_id"].(string),
			}
			policyResources = append(policyResources, resourceParam)
		}
	} else {
		policyResources = append(policyResources, models.PolicyResource{AccountID: accountID})
	}

	iamClient, err := meta.(ClientSession).IAMAPI()
	if err != nil {
		return models.Policy{}, err
	}

	iamRepo := iamClient.ServiceRoles()

	var roles []models.PolicyRole

	if serviceName == "" {
		roles, err = iamRepo.ListSystemDefinedRoles()
	} else {
		roles, err = iamRepo.ListServiceRoles(serviceName)
	}
	if err != nil {
		return models.Policy{}, err
	}

	policyRoles, err := getRolesFromRoleNames(expandStringList(d.Get("roles").([]interface{})), roles)
	if err != nil {
		return models.Policy{}, err
	}

	return models.Policy{Roles: policyRoles, Resources: policyResources}, nil
}

func getRolesFromRoleNames(roleNames []string, roles []models.PolicyRole) ([]models.PolicyRole, error) {

	filteredRoles := []models.PolicyRole{}
	for _, roleName := range roleNames {
		role, err := findRoleByName(roles, roleName)
		if err != nil {
			return []models.PolicyRole{}, err
		}
		filteredRoles = append(filteredRoles, role)
	}
	return filteredRoles, nil
}
