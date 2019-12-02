package ibm

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/IBM-Cloud/bluemix-go/models"
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	// MEMBER ...
	MEMBER = "MEMEBER"
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
				Type:        schema.TypeList,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"users_info": {
				Description: "Additional Access to cloud foundry roles",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"org_id": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "Cloud foundry Organization ID",
						},

						"org_roles": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "user roles at the org level",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},

			"access_groups": {
				Description: "List of access group ids",
				Type:        schema.TypeList,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},

			"organization_roles": {
				Description: "Additional Access to cloud foundry roles",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"org_id": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "Cloud foundry Organization ID",
						},

						"org_roles": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "user roles at the org level",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},

						"spaces": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Description: "Sapces within the org",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"space_id": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "ID of the space.",
									},

									"space_roles": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "user roles at the space level",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
					},
				},
			},

			"infrastructure_roles": {
				Description: "Additional list of permissions to classic infrastructure",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"permissions": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "List of Permissions",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},

			"iam_policies": {
				Description: "Additional iam access policies",
				Type:        schema.TypeList,
				Optional:    true,
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
							ForceNew: true,
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
	Client := userManagement.UserInvite()

	userList := d.Get("users")
	users := make([]models.User, 0)
	for _, user := range userList.([]interface{}) {
		users = append(users, models.User{Email: user.(string), AccountRole: MEMBER})
	}
	if len(users) == 0 {
		return fmt.Errorf("Users email not provided")
	}
	InviteUserPayload := models.UserInvite{Users: users}

	AccountID, err := GetAccountID(d, meta)
	if err != nil {
		return err
	}
	_, InviteUserError := Client.InviteUsers(AccountID, InviteUserPayload)
	if InviteUserError != nil {
		return InviteUserError
	}
	d.SetId(time.Now().UTC().String())
	return resourceIBMIAMUpdateUserProfile(d, meta)
}

func resourceIBMIAMGetUsers(d *schema.ResourceData, meta interface{}) error {
	userManagement, err := meta.(ClientSession).UserManagementAPI()
	Client := userManagement.UserInvite()

	AccountID, err := GetAccountID(d, meta)
	if err != nil {
		return err
	}

	res, err := Client.GetUsers(AccountID)
	if err != nil {
		return err
	}

	users := make([]string, 0)
	for _, user := range res.Resources {
		if user.AccountID != AccountID {
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
		AccountID, err := GetAccountID(d, meta)
		if err != nil {
			return err
		}
		oldUsers, newUsers := d.GetChange("users")

		var present bool
		var removed = make([]string, 0)
		var added = make([]string, 0)

		for _, o := range oldUsers.([]interface{}) {
			present = false
			for _, n := range newUsers.([]interface{}) {
				if strings.Compare(o.(string), n.(string)) == 0 {
					present = true
				}
			}
			if !present {
				removed = append(removed, o.(string))

			}
		}

		for _, n := range newUsers.([]interface{}) {
			present = false
			for _, o := range oldUsers.([]interface{}) {
				if strings.Compare(n.(string), o.(string)) == 0 {
					present = true
				}
			}
			if !present {
				added = append(added, n.(string))
			}

		}
		log.Println("Removed users :", removed)
		log.Println("Added users : ", added)

		//Update the added users
		if len(added) > 0 {
			users := make([]models.User, 0)
			for _, user := range added {
				users = append(users, models.User{Email: user, AccountRole: MEMBER})
			}
			if len(users) == 0 {
				return fmt.Errorf("Users email not provided")
			}
			InviteUserPayload := models.UserInvite{Users: users}

			_, InviteUserError := Client.InviteUsers(AccountID, InviteUserPayload)
			if InviteUserError != nil {
				return InviteUserError
			}
		}

		//Update the removed users
		if len(removed) > 0 {
			for _, user := range removed {
				IAMID := GetUserIAMID(d, meta, user)
				if IAMID == "" {
					return fmt.Errorf("User's IAM ID not found")
				}
				err := Client.RemoveUsers(AccountID, IAMID)
				if err != nil {
					log.Println("Failed to remove user: ", user)
					return err
				}
			}
		}

	}
	return resourceIBMIAMGetUsers(d, meta)
}

func resourceIBMIAMRemoveUser(d *schema.ResourceData, meta interface{}) error {
	userManagement, err := meta.(ClientSession).UserManagementAPI()
	Client := userManagement.UserInvite()

	AccountID, err := GetAccountID(d, meta)
	if err != nil {
		return err
	}

	userList := d.Get("users")
	for _, user := range userList.([]interface{}) {
		IAMID := GetUserIAMID(d, meta, user.(string))
		if IAMID == "" {
			return fmt.Errorf("User's IAM ID not found")
		}
		err := Client.RemoveUsers(AccountID, IAMID)
		if err != nil {
			log.Println("Failed to remove user: ", user.(string))
			return err
		}
	}
	return nil
}

func resourceIBMIAMGetUserProfileExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	userManagement, err := meta.(ClientSession).UserManagementAPI()
	Client := userManagement.UserInvite()

	AccountID, err := GetAccountID(d, meta)
	if err != nil {
		return false, err
	}

	userList := d.Get("users")

	res, err := Client.GetUsers(AccountID)
	if err != nil {
		return false, err
	}
	var isFound bool
	for _, user := range userList.([]interface{}) {

		for _, userInfo := range res.Resources {
			if strings.Compare(userInfo.Email, user.(string)) == 0 {
				isFound = true
			}
		}
		if !isFound {
			return false, fmt.Errorf("Didn't find the user : %s", user)
		}
	}
	return true, nil
}

// GetAccountID returns accountID
func GetAccountID(d *schema.ResourceData, meta interface{}) (string, error) {
	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return "", err
	}
	return userDetails.userAccount, nil
}

// GetUserIAMID ...
func GetUserIAMID(d *schema.ResourceData, meta interface{}, user string) string {
	userManagement, err := meta.(ClientSession).UserManagementAPI()
	Client := userManagement.UserInvite()

	AccountID, err := GetAccountID(d, meta)
	if err != nil {
		return ""
	}

	res, err := Client.GetUsers(AccountID)
	if err != nil {
		return ""
	}

	for _, userInfo := range res.Resources {
		if strings.Compare(userInfo.Email, user) == 0 {
			return userInfo.IamID
		}
	}
	return ""

}
