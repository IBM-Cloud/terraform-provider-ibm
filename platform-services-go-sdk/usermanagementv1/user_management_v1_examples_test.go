//go:build examples

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

package usermanagementv1_test

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/usermanagementv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
	"time"
)

var _ = Describe(`UserManagementV1 Examples Tests`, func() {

	const externalConfigFile = "../user_management.env"

	var (
		userManagementService      *usermanagementv1.UserManagementV1
		userManagementAdminService *usermanagementv1.UserManagementV1
		config                     map[string]string
		configLoaded               bool = false

		accountID     string
		userID        string
		memberEmail   string
		viewerRoleID  string
		accessGroupID string

		deleteUserID string
	)

	var shouldSkipTest = func() {
		if !configLoaded {
			Skip("External configuration is not available, skipping tests...")
		}
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(usermanagementv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			accountID = config["ACCOUNT_ID"]
			Expect(accountID).ToNot(BeEmpty())

			userID = config["USER_ID"]
			Expect(userID).ToNot(BeEmpty())

			memberEmail = config["MEMBER_EMAIL"]
			Expect(memberEmail).ToNot(BeEmpty())

			viewerRoleID = config["VIEWER_ROLE_ID"]
			Expect(viewerRoleID).ToNot(BeEmpty())

			accessGroupID = config["ACCESS_GROUP_ID"]
			Expect(accessGroupID).ToNot(BeEmpty())

			configLoaded = len(config) > 0
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			userManagementServiceOptions := &usermanagementv1.UserManagementV1Options{
				ServiceName: usermanagementv1.DefaultServiceName,
			}

			userManagementService, err = usermanagementv1.NewUserManagementV1UsingExternalConfig(userManagementServiceOptions)

			if err != nil {
				panic(err)
			}

			userManagementAdminServiceOptions := &usermanagementv1.UserManagementV1Options{
				ServiceName: "USER_MANAGEMENT_ADMIN",
			}
			userManagementAdminService, err = usermanagementv1.NewUserManagementV1UsingExternalConfig(userManagementAdminServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(userManagementService).ToNot(BeNil())
			userManagementService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`UserManagementV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`InviteUsers request example`, func() {
			Expect(accountID).ToNot(BeEmpty())

			fmt.Println("\nInviteUsers() result:")
			// begin-invite_users

			inviteUserModel := &usermanagementv1.InviteUser{
				Email:       &memberEmail,
				AccountRole: core.StringPtr("Member"),
			}

			roleModel := &usermanagementv1.Role{
				RoleID: &viewerRoleID,
			}

			attributeModel := &usermanagementv1.Attribute{
				Name:  core.StringPtr("accountId"),
				Value: &accountID,
			}

			attributeModel2 := &usermanagementv1.Attribute{
				Name:  core.StringPtr("resourceGroupId"),
				Value: core.StringPtr("*"),
			}

			resourceModel := &usermanagementv1.Resource{
				Attributes: []usermanagementv1.Attribute{*attributeModel, *attributeModel2},
			}

			inviteUserIamPolicyModel := &usermanagementv1.InviteUserIamPolicy{
				Type:      core.StringPtr("access"),
				Roles:     []usermanagementv1.Role{*roleModel},
				Resources: []usermanagementv1.Resource{*resourceModel},
			}

			inviteUsersOptions := &usermanagementv1.InviteUsersOptions{
				AccountID:    &accountID,
				Users:        []usermanagementv1.InviteUser{*inviteUserModel},
				IamPolicy:    []usermanagementv1.InviteUserIamPolicy{*inviteUserIamPolicyModel},
				AccessGroups: []string{accessGroupID},
			}

			invitedUserList, response, err := userManagementAdminService.InviteUsers(inviteUsersOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(invitedUserList, "", "  ")
			fmt.Println(string(b))

			// end-invite_users

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(invitedUserList).ToNot(BeNil())
			Expect(invitedUserList.Resources).ToNot(BeEmpty())

			for _, res := range invitedUserList.Resources {
				deleteUserID = *res.ID
			}
		})
		It(`ListUsers request example`, func() {
			fmt.Println("\nListUsers() result:")
			// begin-list_users
			listUsersOptions := &usermanagementv1.ListUsersOptions{
				AccountID:       &accountID,
				IncludeSettings: core.BoolPtr(true),
				Search:          core.StringPtr("state:ACTIVE"),
			}

			pager, err := userManagementService.NewUsersPager(listUsersOptions)
			if err != nil {
				panic(err)
			}
			var allResults []usermanagementv1.UserProfile
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_users
		})
		It(`RemoveUser request example`, func() {
			Expect(accountID).ToNot(BeEmpty())
			Expect(deleteUserID).ToNot(BeEmpty())

			// begin-remove_user

			removeUserOptions := userManagementService.NewRemoveUserOptions(
				accountID,
				deleteUserID,
			)

			response, err := userManagementAdminService.RemoveUser(removeUserOptions)
			if err != nil {
				panic(err)
			}

			// end-remove_user
			fmt.Printf("\nRemoveUser() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`GetUserProfile request example`, func() {
			Expect(accountID).ToNot(BeEmpty())
			Expect(userID).ToNot(BeEmpty())

			fmt.Println("\nGetUserProfile() result:")
			// begin-get_user_profile

			getUserProfileOptions := userManagementService.NewGetUserProfileOptions(
				accountID,
				userID,
			)

			userProfile, response, err := userManagementService.GetUserProfile(getUserProfileOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(userProfile, "", "  ")
			fmt.Println(string(b))

			// end-get_user_profile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(userProfile).ToNot(BeNil())

		})
		It(`UpdateUserProfile request example`, func() {
			Expect(accountID).ToNot(BeEmpty())
			Expect(userID).ToNot(BeEmpty())

			// begin-update_user_profile

			updateUserProfileOptions := userManagementService.NewUpdateUserProfileOptions(
				accountID,
				userID,
			)
			updateUserProfileOptions.SetPhonenumber("123456789")

			response, err := userManagementService.UpdateUserProfile(updateUserProfileOptions)
			if err != nil {
				panic(err)
			}

			// end-update_user_profile
			fmt.Printf("\nUpdateUserProfile() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`GetUserSettings request example`, func() {
			Expect(accountID).ToNot(BeEmpty())
			Expect(userID).ToNot(BeEmpty())

			fmt.Println("\nGetUserSettings() result:")
			// begin-get_user_settings

			getUserSettingsOptions := userManagementService.NewGetUserSettingsOptions(
				accountID,
				userID,
			)

			userSettings, response, err := userManagementService.GetUserSettings(getUserSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(userSettings, "", "  ")
			fmt.Println(string(b))

			// end-get_user_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(userSettings).ToNot(BeNil())

		})
		It(`UpdateUserSettings request example`, func() {
			Expect(accountID).ToNot(BeEmpty())
			Expect(userID).ToNot(BeEmpty())

			// begin-update_user_settings

			updateUserSettingsOptions := userManagementService.NewUpdateUserSettingsOptions(
				accountID,
				userID,
			)
			updateUserSettingsOptions.SetSelfManage(true)
			updateUserSettingsOptions.SetAllowedIPAddresses("192.168.0.2,192.168.0.3")

			response, err := userManagementService.UpdateUserSettings(updateUserSettingsOptions)
			if err != nil {
				panic(err)
			}

			// end-update_user_settings
			fmt.Printf("\nUpdateUserSettings() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})
