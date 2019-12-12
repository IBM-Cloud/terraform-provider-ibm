package ibm

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/IBM-Cloud/bluemix-go/api/iampap/iampapv1"
	v2 "github.com/IBM-Cloud/bluemix-go/api/usermanagement/usermanagementv2"
	"github.com/IBM-Cloud/bluemix-go/models"
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	// MEMBER ...
	MEMBER = "MEMEBER"
	// ACCESS ...
	ACCESS = "access"
)

func resourceIBMUserInvite() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMIAMInviteUsers,
		Read:     resourceIBMIAMGetUsers,
		Update:   resourceIBMIAMUpdateUserProfile,
		Delete:   resourceIBMIAMRemoveUser,
		Exists:   resourceIBMIAMGetUserProfileExists,
		Importer: &schema.ResourceImporter{},
		Schema: map[string]*schema.Schema{

			"users": {
				Description: "List of ibm id or email of user",
				Type:        schema.TypeSet,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"access_groups": {
				Description: "access group ids to associate the inviting user",
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"iam_policy": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

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
					},
				},
			},
		},
	}
}

func resourceIBMIAMInviteUsers(d *schema.ResourceData, meta interface{}) error {
	userManagement, err := meta.(ClientSession).UserManagementAPI()
	if err != nil {
		return err
	}
	client := userManagement.UserInvite()

	usersSet := d.Get("users").(*schema.Set)
	usersList := flattenUsersSet(usersSet)
	users := make([]v2.User, 0)
	for _, user := range usersList {
		users = append(users, v2.User{Email: user, AccountRole: MEMBER})
	}
	if len(users) == 0 {
		return fmt.Errorf("Users email not provided")
	}
	var accessGroups = make([]string, 0)
	if data, ok := d.GetOk("access_groups"); ok {
		for _, accessGroup := range data.([]interface{}) {
			accessGroups = append(accessGroups, fmt.Sprintf("%v", accessGroup))
		}
	}

	var accessPolicies []v2.UserPolicy
	if accessPolicyData, ok := d.GetOk("iam_policy"); ok {
		accessPolicies, err = getPolicies(d, meta, accessPolicyData.([]interface{}))
		if err != nil {
			log.Println("IAM Acess policy: ", err.Error())
			return err
		}
	}

	inviteUserPayload := v2.UserInvite{Users: users, AccessGroup: accessGroups, IAMPolicy: accessPolicies}
	accountID, err := getAccountID(d, meta)
	if err != nil {
		return err
	}

	_, InviteUserError := client.InviteUsers(accountID, inviteUserPayload)
	if InviteUserError != nil {
		return InviteUserError
	}
	d.SetId(time.Now().UTC().String())
	return resourceIBMIAMUpdateUserProfile(d, meta)
}

func resourceIBMIAMGetUsers(d *schema.ResourceData, meta interface{}) error {
	userManagement, err := meta.(ClientSession).UserManagementAPI()
	if err != nil {
		return err
	}
	Client := userManagement.UserInvite()

	accountID, err := getAccountID(d, meta)
	if err != nil {
		return err
	}

	res, err := Client.GetUsers(accountID)
	if err != nil {
		return err
	}

	users := make([]string, 0)
	for _, user := range res.Resources {
		if user.AccountID != accountID {
			users = append(users, user.Email)
		}
	}
	return nil

}

func resourceIBMIAMUpdateUserProfile(d *schema.ResourceData, meta interface{}) error {
	// validate change
	userManagement, err := meta.(ClientSession).UserManagementAPI()
	if err != nil {
		return err
	}
	Client := userManagement.UserInvite()

	if d.HasChange("users") {
		//var removedUsers, addedUsers []string
		accountID, err := getAccountID(d, meta)
		if err != nil {
			return err
		}
		ousrs, nusrs := d.GetChange("users")
		old := ousrs.(*schema.Set)
		new := nusrs.(*schema.Set)

		removed := expandStringList(old.Difference(new).List())
		added := expandStringList(new.Difference(old).List())

		//Update the added users
		if len(added) > 0 {
			users := make([]v2.User, 0)
			for _, user := range added {
				users = append(users, v2.User{Email: user, AccountRole: MEMBER})
			}
			if len(users) == 0 {
				return fmt.Errorf("Users email not provided")
			}

			var accessPolicies []v2.UserPolicy
			if accessPolicyData, ok := d.GetOk("iam_policy"); ok {
				accessPolicies, err = getPolicies(d, meta, accessPolicyData.([]interface{}))
				if err != nil {
					log.Println("IAM Acess policy: ", err.Error())
					return err
				}
			}

			var accessGroups = make([]string, 0)
			if data, ok := d.GetOk("access_groups"); ok {
				for _, accessGroup := range data.([]interface{}) {
					accessGroups = append(accessGroups, fmt.Sprintf("%v", accessGroup))
				}
			}

			InviteUserPayload := v2.UserInvite{Users: users, AccessGroup: accessGroups, IAMPolicy: accessPolicies}

			_, InviteUserError := Client.InviteUsers(accountID, InviteUserPayload)
			if InviteUserError != nil {
				return InviteUserError
			}
		}

		//Update the removed users
		if len(removed) > 0 {
			for _, user := range removed {
				IAMID, err := getUserIAMID(d, meta, user)
				if err != nil {
					return fmt.Errorf("User's IAM ID not found: %s", err.Error())
				}
				Err := Client.RemoveUsers(accountID, IAMID)
				if Err != nil {
					log.Println("Failed to remove user: ", user)
					return Err
				}
			}
		}

	}
	return resourceIBMIAMGetUsers(d, meta)
}

func resourceIBMIAMRemoveUser(d *schema.ResourceData, meta interface{}) error {
	userManagement, err := meta.(ClientSession).UserManagementAPI()
	if err != nil {
		return err
	}
	Client := userManagement.UserInvite()

	accountID, err := getAccountID(d, meta)
	if err != nil {
		return err
	}

	usersSet := d.Get("users").(*schema.Set)
	usersList := flattenUsersSet(usersSet)
	for _, user := range usersList {
		IAMID, err := getUserIAMID(d, meta, user)

		if err != nil {
			return fmt.Errorf("User's IAM ID not found: %s", err.Error())
		}
		Err := Client.RemoveUsers(accountID, IAMID)
		if Err != nil {
			return Err
		}
	}
	return nil
}

func resourceIBMIAMGetUserProfileExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	userManagement, err := meta.(ClientSession).UserManagementAPI()
	if err != nil {
		return false, err
	}
	Client := userManagement.UserInvite()

	accountID, err := getAccountID(d, meta)
	if err != nil {
		return false, err
	}

	usersSet := d.Get("users").(*schema.Set)
	usersList := flattenUsersSet(usersSet)

	res, err := Client.GetUsers(accountID)
	if err != nil {
		return false, err
	}
	var isFound bool
	for _, user := range usersList {

		for _, userInfo := range res.Resources {
			if strings.Compare(userInfo.Email, user) == 0 {
				isFound = true
			}
		}
		if !isFound {
			return false, fmt.Errorf("Didn't find the user : %s", user)
		}
	}
	return true, nil
}

// getAccountID returns accountID
func getAccountID(d *schema.ResourceData, meta interface{}) (string, error) {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return "", err
	}
	return userDetails.userAccount, nil
}

// getUserIAMID ...
func getUserIAMID(d *schema.ResourceData, meta interface{}, user string) (string, error) {
	userManagement, err := meta.(ClientSession).UserManagementAPI()
	if err != nil {
		return "", err
	}
	Client := userManagement.UserInvite()

	accountID, err := getAccountID(d, meta)
	if err != nil {
		return "", err
	}

	res, err := Client.GetUsers(accountID)
	if err != nil {
		return "", err
	}

	for _, userInfo := range res.Resources {
		if strings.Compare(userInfo.Email, user) == 0 {
			return userInfo.IamID, nil
		}
	}
	return "", nil

}

// getPolicies ...
func getPolicies(d *schema.ResourceData, meta interface{}, policies []interface{}) ([]v2.UserPolicy, error) {
	var policyList = make([]v2.UserPolicy, 0)
	for _, policy := range policies {
		p := policy.(map[string]interface{})
		var serviceName string
		policyResource := iampapv1.Resource{}
		if res, ok := p["resources"]; ok {
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

		if len(policyResource.Attributes) == 0 {
			policyResource.SetServiceType("service")
		}

		accountID, err := getAccountID(d, meta)
		if err != nil {
			return policyList, err
		}
		policyResource.SetAccountID(accountID)

		iamClient, err := meta.(ClientSession).IAMAPI()
		if err != nil {
			return policyList, err
		}

		iamRepo := iamClient.ServiceRoles()

		var roles []models.PolicyRole

		if serviceName == "" {
			roles, err = iamRepo.ListSystemDefinedRoles()
		} else {
			roles, err = iamRepo.ListServiceRoles(serviceName)
		}
		if err != nil {
			return policyList, err
		}
		var policyRoles = make([]models.PolicyRole, 0)
		if userRoles, ok := p["roles"]; ok {
			policyRoles, err = getRolesFromRoleNames(expandStringList(userRoles.([]interface{})), roles)
			if err != nil {
				return policyList, err
			}
		}

		policyList = append(policyList, v2.UserPolicy{Roles: iampapv1.ConvertRoleModels(policyRoles), Resources: []iampapv1.Resource{policyResource}, Type: ACCESS})
	}
	return policyList, nil
}
