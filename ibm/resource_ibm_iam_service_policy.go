package ibm

import (
	"fmt"

	"github.com/IBM-Cloud/bluemix-go/api/iampap/iampapv1"
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

func resourceIBMIAMServicePolicyCreate(d *schema.ResourceData, meta interface{}) error {
	iamClient, err := meta.(ClientSession).IAMAPI()
	if err != nil {
		return err
	}
	serviceIDUUID := d.Get("iam_service_id").(string)

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	var policy iampapv1.Policy

	policy, err = generateAccountPolicy(d, meta)
	if err != nil {
		return err
	}

	serviceID, err := iamClient.ServiceIds().Get(serviceIDUUID)
	if err != nil {
		return err
	}

	iampapClient, err := meta.(ClientSession).IAMPAPAPI()
	if err != nil {
		return err
	}

	policy.Resources[0].SetAccountID(userDetails.userAccount)

	policy.Subjects = []iampapv1.Subject{
		{
			Attributes: []iampapv1.Attribute{
				{
					Name:  "iam_id",
					Value: serviceID.IAMID,
				},
			},
		},
	}

	policy.Type = iampapv1.AccessPolicyType

	servicePolicy, err := iampapClient.V1Policy().Create(policy)

	if err != nil {
		return fmt.Errorf("Error creating servicePolicy: %s", err)
	}

	d.SetId(fmt.Sprintf("%s/%s", serviceIDUUID, servicePolicy.ID))

	return resourceIBMIAMServicePolicyRead(d, meta)
}

func resourceIBMIAMServicePolicyRead(d *schema.ResourceData, meta interface{}) error {

	iampapClient, err := meta.(ClientSession).IAMPAPAPI()
	if err != nil {
		return err
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	serviceIDUUID := parts[0]
	servicePolicyID := parts[1]

	servicePolicy, err := iampapClient.V1Policy().Get(servicePolicyID)
	if err != nil {
		return fmt.Errorf("Error retrieving servicePolicy: %s", err)
	}

	d.Set("iam_service_id", serviceIDUUID)
	roles := make([]string, len(servicePolicy.Roles))
	for i, role := range servicePolicy.Roles {
		roles[i] = role.Name
	}
	d.Set("roles", roles)
	d.Set("version", servicePolicy.Version)
	d.Set("resources", flattenPolicyResource(servicePolicy.Resources))
	if len(servicePolicy.Resources) > 0 {
		if servicePolicy.Resources[0].GetAttribute("serviceType") == "service" {
			d.Set("account_management", false)
		}
		if servicePolicy.Resources[0].GetAttribute("serviceType") == "platform_service" {
			d.Set("account_management", true)
		}
	}

	return nil
}

func resourceIBMIAMServicePolicyUpdate(d *schema.ResourceData, meta interface{}) error {

	if d.HasChange("roles") || d.HasChange("resources") {

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

		userDetails, err := meta.(ClientSession).BluemixUserDetails()
		if err != nil {
			return err
		}

		serviceID, err := iamClient.ServiceIds().Get(serviceIDUUID)
		if err != nil {
			return err
		}

		var policy iampapv1.Policy

		policy, err = generateAccountPolicy(d, meta)
		if err != nil {
			return err
		}

		policy.Resources[0].SetAccountID(userDetails.userAccount)

		policy.Subjects = []iampapv1.Subject{
			{
				Attributes: []iampapv1.Attribute{
					{
						Name:  "iam_id",
						Value: serviceID.IAMID,
					},
				},
			},
		}

		policy.Type = iampapv1.AccessPolicyType

		iampapClient, err := meta.(ClientSession).IAMPAPAPI()
		if err != nil {
			return err
		}

		_, err = iampapClient.V1Policy().Update(servicePolicyID, policy, d.Get("version").(string))
		if err != nil {
			return fmt.Errorf("Error updating service policy: %s", err)
		}

	}

	return resourceIBMIAMServicePolicyRead(d, meta)

}

func resourceIBMIAMServicePolicyDelete(d *schema.ResourceData, meta interface{}) error {
	iampapClient, err := meta.(ClientSession).IAMPAPAPI()
	if err != nil {
		return err
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	servicePolicyID := parts[1]

	err = iampapClient.V1Policy().Delete(servicePolicyID)
	if err != nil {
		return fmt.Errorf("Error deleting service policy: %s", err)
	}

	d.SetId("")

	return nil
}

func resourceIBMIAMServicePolicyExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	iampapClient, err := meta.(ClientSession).IAMPAPAPI()
	if err != nil {
		return false, err
	}
	parts, err := idParts(d.Id())
	if err != nil {
		return false, err
	}
	serviceIDUUID := parts[0]
	servicePolicyID := parts[1]

	servicePolicy, err := iampapClient.V1Policy().Get(servicePolicyID)
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
