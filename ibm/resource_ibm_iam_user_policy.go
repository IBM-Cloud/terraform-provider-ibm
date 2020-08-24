package ibm

import (
	"fmt"

	"github.com/IBM-Cloud/bluemix-go/api/iampap/iampapv1"
	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceIBMIAMUserPolicy() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMIAMUserPolicyCreate,
		Read:     resourceIBMIAMUserPolicyRead,
		Update:   resourceIBMIAMUserPolicyUpdate,
		Delete:   resourceIBMIAMUserPolicyDelete,
		Exists:   resourceIBMIAMUserPolicyExists,
		Importer: &schema.ResourceImporter{},
		Schema: map[string]*schema.Schema{

			"ibm_id": {
				Description: "The ibm id or email of user",
				Type:        schema.TypeString,
				Required:    true,
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

func resourceIBMIAMUserPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	iampapClient, err := meta.(ClientSession).IAMPAPAPI()
	if err != nil {
		return err
	}

	userEmail := d.Get("ibm_id").(string)

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	accountID := userDetails.userAccount
	var policy iampapv1.Policy
	policy, err = generateAccountPolicyV2(d, meta)
	if err != nil {
		return err
	}

	policy.Resources[0].SetAccountID(accountID)

	policy.Type = iampapv1.AccessPolicyType

	ibmUniqueID, err := getIBMUniqueId(accountID, userEmail, meta)
	if err != nil {
		return err
	}

	policy.Subjects = []iampapv1.Subject{
		{
			Attributes: []iampapv1.Attribute{
				{
					Name:  "iam_id",
					Value: ibmUniqueID,
				},
			},
		},
	}

	userPolicy, err := iampapClient.V1Policy().Create(policy)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s", userEmail, userPolicy.ID))
	return resourceIBMIAMUserPolicyRead(d, meta)
}

func resourceIBMIAMUserPolicyRead(d *schema.ResourceData, meta interface{}) error {
	iampapClient, err := meta.(ClientSession).IAMPAPAPI()
	if err != nil {
		return err
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}

	userEmail := parts[0]
	userPolicyID := parts[1]

	if err != nil {
		return err
	}

	userPolicy, err := iampapClient.V1Policy().Get(userPolicyID)
	if err != nil {
		return err
	}
	d.Set("ibm_id", userEmail)
	roles := make([]string, len(userPolicy.Roles))
	for i, role := range userPolicy.Roles {
		roles[i] = role.Name
	}
	d.Set("roles", roles)
	d.Set("version", userPolicy.Version)
	d.Set("resources", flattenPolicyResource(userPolicy.Resources))
	if len(userPolicy.Resources) > 0 {
		if userPolicy.Resources[0].GetAttribute("serviceType") == "service" {
			d.Set("account_management", false)
		}
		if userPolicy.Resources[0].GetAttribute("serviceType") == "platform_service" {
			d.Set("account_management", true)
		}
	}
	return nil
}

func resourceIBMIAMUserPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	iampapClient, err := meta.(ClientSession).IAMPAPAPI()
	if err != nil {
		return err
	}
	if d.HasChange("roles") || d.HasChange("resources") || d.HasChange("account_management") {
		parts, err := idParts(d.Id())
		if err != nil {
			return err
		}
		userEmail := parts[0]
		userPolicyID := parts[1]

		userDetails, err := meta.(ClientSession).BluemixUserDetails()
		if err != nil {
			return err
		}

		var policy iampapv1.Policy
		accountID := userDetails.userAccount

		policy, err = generateAccountPolicyV2(d, meta)
		if err != nil {
			return err
		}

		ibmUniqueID, err := getIBMUniqueId(accountID, userEmail, meta)
		if err != nil {
			return err
		}

		policy.Resources[0].SetAccountID(accountID)

		policy.Subjects = []iampapv1.Subject{
			{
				Attributes: []iampapv1.Attribute{
					{
						Name:  "iam_id",
						Value: ibmUniqueID,
					},
				},
			},
		}

		policy.Type = iampapv1.AccessPolicyType

		_, err = iampapClient.V1Policy().Update(userPolicyID, policy, d.Get("version").(string))
		if err != nil {
			return fmt.Errorf("Error updating user policy: %s", err)
		}
	}
	return resourceIBMIAMUserPolicyRead(d, meta)
}

func resourceIBMIAMUserPolicyDelete(d *schema.ResourceData, meta interface{}) error {

	iampapClient, err := meta.(ClientSession).IAMPAPAPI()
	if err != nil {
		return err
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return err
	}
	userPolicyID := parts[1]

	err = iampapClient.V1Policy().Delete(userPolicyID)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceIBMIAMUserPolicyExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	iampapClient, err := meta.(ClientSession).IAMPAPAPI()
	if err != nil {
		return false, err
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return false, err
	}
	userEmail := parts[0]
	userPolicyID := parts[1]

	userPolicy, err := iampapClient.V1Policy().Get(userPolicyID)
	if err != nil {
		if apiErr, ok := err.(bmxerror.RequestFailure); ok {
			if apiErr.StatusCode() == 404 {
				return false, nil
			}
		}
		return false, fmt.Errorf("Error communicating with the API: %s", err)
	}

	tempID := fmt.Sprintf("%s/%s", userEmail, userPolicy.ID)

	return tempID == d.Id(), nil

}

func getIBMUniqueId(accountID, userEmail string, meta interface{}) (string, error) {
	userManagement, err := meta.(ClientSession).UserManagementAPI()
	if err != nil {
		return "", err
	}
	client := userManagement.UserInvite()
	res, err := client.GetUsers(accountID)
	if err != nil {
		return "", err
	}
	for _, userInfo := range res.Resources {
		if userInfo.Email == userEmail {
			return userInfo.IamID, nil
		}
	}
	return "", fmt.Errorf("User %s is not found under account %s", userEmail, accountID)
}
