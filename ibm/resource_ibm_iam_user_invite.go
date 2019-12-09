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
	users := make([]models.User, 0)
	for _, user := range usersList {
		users = append(users, models.User{Email: user, AccountRole: MEMBER})
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
	inviteUserPayload := models.UserInvite{Users: users, AccessGroup: accessGroups}

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
			users := make([]models.User, 0)
			for _, user := range added {
				users = append(users, models.User{Email: user, AccountRole: MEMBER})
			}
			if len(users) == 0 {
				return fmt.Errorf("Users email not provided")
			}
			InviteUserPayload := models.UserInvite{Users: users}

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
