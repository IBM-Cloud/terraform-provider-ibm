package ibm

import (
	"fmt"

	"github.com/IBM-Cloud/bluemix-go/api/account/accountv1"
	"github.com/IBM-Cloud/bluemix-go/bmxerror"
	"github.com/IBM-Cloud/bluemix-go/models"
	"github.com/hashicorp/terraform/helper/schema"
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

func resourceIBMIAMUserPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	iamClient, err := meta.(ClientSession).IAMAPI()
	if err != nil {
		return err
	}
	userEmail := d.Get("ibm_id").(string)

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	accountID := userDetails.userAccount
	var policy models.Policy
	policy, err = generatePolicy(d, meta, accountID)
	if err != nil {
		return err
	}

	user, err := getAccountUser(accountID, userEmail, meta)
	if err != nil {
		return err
	}

	userPolicy, err := iamClient.UserPolicies().Create("a/"+accountID, user.IbmUniqueId, policy)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s", userEmail, userPolicy.ID))
	return resourceIBMIAMUserPolicyRead(d, meta)
}

func resourceIBMIAMUserPolicyRead(d *schema.ResourceData, meta interface{}) error {
	iamClient, err := meta.(ClientSession).IAMAPI()
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

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	accountID := userDetails.userAccount

	user, err := getAccountUser(accountID, userEmail, meta)
	if err != nil {
		return err
	}

	userPolicy, err := iamClient.UserPolicies().Get("a/"+accountID, user.IbmUniqueId, userPolicyID)
	if err != nil {
		return err
	}
	d.Set("ibm_id", userEmail)
	roles := make([]string, len(userPolicy.Roles))
	for i, role := range userPolicy.Roles {
		roles[i] = role.DisplayName
	}
	d.Set("roles", roles)
	d.Set("version", userPolicy.Version)
	d.Set("resources", flattenPolicyResource(userPolicy.Resources))
	return nil
}

func resourceIBMIAMUserPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	iamClient, err := meta.(ClientSession).IAMAPI()
	if err != nil {
		return err
	}
	if d.HasChange("roles") || d.HasChange("resources") {
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

		var policy models.Policy
		accountID := userDetails.userAccount

		policy, err = generatePolicy(d, meta, accountID)
		if err != nil {
			return err
		}

		user, err := getAccountUser(accountID, userEmail, meta)

		_, err = iamClient.UserPolicies().Update("a/"+accountID, user.IbmUniqueId, userPolicyID, policy, d.Get("version").(string))
		if err != nil {
			return fmt.Errorf("Error updating user policy: %s", err)
		}
	}
	return resourceIBMIAMUserPolicyRead(d, meta)
}

func resourceIBMIAMUserPolicyDelete(d *schema.ResourceData, meta interface{}) error {

	iamClient, err := meta.(ClientSession).IAMAPI()
	if err != nil {
		return err
	}

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

	accountID := userDetails.userAccount

	user, err := getAccountUser(accountID, userEmail, meta)
	if err != nil {
		return err
	}

	err = iamClient.UserPolicies().Delete("a/"+accountID, user.IbmUniqueId, userPolicyID)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceIBMIAMUserPolicyExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	iamClient, err := meta.(ClientSession).IAMAPI()
	if err != nil {
		return false, err
	}

	parts, err := idParts(d.Id())
	if err != nil {
		return false, err
	}
	userEmail := parts[0]
	userPolicyID := parts[1]

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return false, err
	}

	accountID := userDetails.userAccount

	user, err := getAccountUser(accountID, userEmail, meta)
	if err != nil {
		return false, err
	}

	userPolicy, err := iamClient.UserPolicies().Get("a/"+accountID, user.IbmUniqueId, userPolicyID)
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

func getAccountUser(accountID, userEmail string, meta interface{}) (*accountv1.AccountUser, error) {

	accountv1Client, err := meta.(ClientSession).BluemixAcccountv1API()
	if err != nil {
		return nil, err
	}
	accUser, err := accountv1Client.Accounts().FindAccountUserByUserId(accountID, userEmail)
	if err != nil {
		return nil, err
	} else if accUser == nil {
		return nil, fmt.Errorf("User %s is not found under current account", userEmail)
	}

	return accUser, nil
}
