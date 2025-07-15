//go:build examples

/**
 * (C) Copyright IBM Corp. 2023.
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

package iamaccessgroupsv2_test

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iamaccessgroupsv2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// This file provides an example of how to use the iam-access-groups service.
//
// The following configuration properties are assumed to be defined:
// IAM_ACCESS_GROUPS_URL=<service base url>
// IAM_ACCESS_GROUPS_AUTH_TYPE=iam
// IAM_ACCESS_GROUPS_APIKEY=<IAM apikey>
// IAM_ACCESS_GROUPS_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
var _ = Describe(`IamAccessGroupsV2 Examples Tests`, func() {

	const externalConfigFile = "../iam_access_groups_v2.env"

	var (
		iamAccessGroupsService *iamaccessgroupsv2.IamAccessGroupsV2
		config                 map[string]string

		// Variables to hold link values
		accessGroupETagLink           string
		accessGroupIDLink             string
		testAccountID                 string
		testProfileID                 string
		testClaimRuleID               string
		testClaimRuleEtag             string
		testPolicyTemplateID          string
		testAccountGroupID            string
		testTemplateId                string
		testTemplateVersionEtag       string
		testTemplateLatestVersionEtag string
		testAssignmentID              string
		testAssignmentEtag            string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping examples...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping examples: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(iamaccessgroupsv2.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			} else if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
			}

			testAccountID = config["TEST_ACCOUNT_ID"]
			testProfileID = config["TEST_PROFILE_ID"]
			testPolicyTemplateID = config["TEST_POLICY_TEMPLATE_ID"]
			testAccountGroupID = config["TEST_ACCOUNT_GROUP_ID"]
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			iamAccessGroupsServiceOptions := &iamaccessgroupsv2.IamAccessGroupsV2Options{}

			iamAccessGroupsService, err = iamaccessgroupsv2.NewIamAccessGroupsV2UsingExternalConfig(iamAccessGroupsServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(iamAccessGroupsService).ToNot(BeNil())
		})
	})

	Describe(`IamAccessGroupsV2 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateAccessGroup request example`, func() {
			fmt.Println("\nCreateAccessGroup() result:")
			// begin-create_access_group

			createAccessGroupOptions := iamAccessGroupsService.NewCreateAccessGroupOptions(
				testAccountID,
				"Managers",
			)
			createAccessGroupOptions.SetDescription("Group for managers")

			group, response, err := iamAccessGroupsService.CreateAccessGroup(createAccessGroupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(group, "", "  ")
			fmt.Println(string(b))

			// end-create_access_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(group).ToNot(BeNil())

			accessGroupIDLink = *group.ID
			fmt.Fprintf(GinkgoWriter, "Saved accessGroupIDLink value: %v\n", accessGroupIDLink)

		})
		It(`GetAccessGroup request example`, func() {
			fmt.Println("\nGetAccessGroup() result:")
			// begin-get_access_group

			getAccessGroupOptions := iamAccessGroupsService.NewGetAccessGroupOptions(
				accessGroupIDLink,
			)

			group, response, err := iamAccessGroupsService.GetAccessGroup(getAccessGroupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(group, "", "  ")
			fmt.Println(string(b))

			// end-get_access_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(group).ToNot(BeNil())

			accessGroupETagLink = response.Headers.Get("ETag")
			fmt.Fprintf(GinkgoWriter, "Saved accessGroupETagLink value: %v\n", accessGroupETagLink)

		})
		It(`UpdateAccessGroup request example`, func() {
			fmt.Println("\nUpdateAccessGroup() result:")
			// begin-update_access_group

			updateAccessGroupOptions := iamAccessGroupsService.NewUpdateAccessGroupOptions(
				accessGroupIDLink,
				accessGroupETagLink,
			)
			updateAccessGroupOptions.SetName("Awesome Managers")
			updateAccessGroupOptions.SetDescription("Group for awesome managers.")

			group, response, err := iamAccessGroupsService.UpdateAccessGroup(updateAccessGroupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(group, "", "  ")
			fmt.Println(string(b))

			// end-update_access_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(group).ToNot(BeNil())

		})
		It(`ListAccessGroups request example`, func() {
			fmt.Println("\nListAccessGroups() result:")
			// begin-list_access_groups
			listAccessGroupsOptions := &iamaccessgroupsv2.ListAccessGroupsOptions{
				AccountID: &testAccountID,
			}

			pager, err := iamAccessGroupsService.NewAccessGroupsPager(listAccessGroupsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []iamaccessgroupsv2.Group
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_access_groups
		})
		It(`AddMembersToAccessGroup request example`, func() {
			fmt.Println("\nAddMembersToAccessGroup() result:")
			// begin-add_members_to_access_group

			groupMembers := []iamaccessgroupsv2.AddGroupMembersRequestMembersItem{
				iamaccessgroupsv2.AddGroupMembersRequestMembersItem{
					IamID: core.StringPtr("IBMid-user1"),
					Type:  core.StringPtr("user"),
				},
				iamaccessgroupsv2.AddGroupMembersRequestMembersItem{
					IamID: core.StringPtr("iam-ServiceId-123"),
					Type:  core.StringPtr("service"),
				},
				iamaccessgroupsv2.AddGroupMembersRequestMembersItem{
					IamID: core.StringPtr(testProfileID),
					Type:  core.StringPtr("profile"),
				},
			}

			addMembersToAccessGroupOptions := iamAccessGroupsService.NewAddMembersToAccessGroupOptions(
				accessGroupIDLink,
			)
			addMembersToAccessGroupOptions.SetMembers(groupMembers)

			addGroupMembersResponse, response, err := iamAccessGroupsService.AddMembersToAccessGroup(addMembersToAccessGroupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(addGroupMembersResponse, "", "  ")
			fmt.Println(string(b))

			// end-add_members_to_access_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(207))
			Expect(addGroupMembersResponse).ToNot(BeNil())

		})
		It(`IsMemberOfAccessGroup request example`, func() {
			// begin-is_member_of_access_group

			isMemberOfAccessGroupOptions := iamAccessGroupsService.NewIsMemberOfAccessGroupOptions(
				accessGroupIDLink,
				"IBMid-user1",
			)

			response, err := iamAccessGroupsService.IsMemberOfAccessGroup(isMemberOfAccessGroupOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from IsMemberOfAccessGroup(): %d\n", response.StatusCode)
			}

			// end-is_member_of_access_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`ListAccessGroupMembers request example`, func() {
			fmt.Println("\nListAccessGroupMembers() result:")
			// begin-list_access_group_members
			listAccessGroupMembersOptions := &iamaccessgroupsv2.ListAccessGroupMembersOptions{
				AccessGroupID: &accessGroupIDLink,
			}

			pager, err := iamAccessGroupsService.NewAccessGroupMembersPager(listAccessGroupMembersOptions)
			if err != nil {
				panic(err)
			}

			var allResults []iamaccessgroupsv2.ListGroupMembersResponseMember
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_access_group_members
		})
		It(`RemoveMemberFromAccessGroup request example`, func() {
			// begin-remove_member_from_access_group

			removeMemberFromAccessGroupOptions := iamAccessGroupsService.NewRemoveMemberFromAccessGroupOptions(
				accessGroupIDLink,
				"IBMid-user1",
			)

			response, err := iamAccessGroupsService.RemoveMemberFromAccessGroup(removeMemberFromAccessGroupOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from RemoveMemberFromAccessGroup(): %d\n", response.StatusCode)
			}

			// end-remove_member_from_access_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`RemoveMemberFromAccessGroup request example`, func() {
			// begin-remove_member_from_access_group

			removeMemberFromAccessGroupOptions := iamAccessGroupsService.NewRemoveMemberFromAccessGroupOptions(
				accessGroupIDLink,
				"iam-ServiceId-123",
			)

			response, err := iamAccessGroupsService.RemoveMemberFromAccessGroup(removeMemberFromAccessGroupOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from RemoveMemberFromAccessGroup(): %d\n", response.StatusCode)
			}

			// end-remove_member_from_access_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`RemoveMemberFromAccessGroup request example`, func() {
			// begin-remove_member_from_access_group

			removeMemberFromAccessGroupOptions := iamAccessGroupsService.NewRemoveMemberFromAccessGroupOptions(
				accessGroupIDLink,
				testProfileID,
			)

			response, err := iamAccessGroupsService.RemoveMemberFromAccessGroup(removeMemberFromAccessGroupOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from RemoveMemberFromAccessGroup(): %d\n", response.StatusCode)
			}

			// end-remove_member_from_access_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`AddMemberToMultipleAccessGroups request example`, func() {
			fmt.Println("\nAddMemberToMultipleAccessGroups() result:")
			// begin-add_member_to_multiple_access_groups

			addMemberToMultipleAccessGroupsOptions := iamAccessGroupsService.NewAddMemberToMultipleAccessGroupsOptions(
				testAccountID,
				"IBMid-user1",
			)

			addMemberToMultipleAccessGroupsOptions.SetType("user")
			addMemberToMultipleAccessGroupsOptions.SetGroups([]string{accessGroupIDLink})

			addMembershipMultipleGroupsResponse, response, err := iamAccessGroupsService.AddMemberToMultipleAccessGroups(addMemberToMultipleAccessGroupsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(addMembershipMultipleGroupsResponse, "", "  ")
			fmt.Println(string(b))

			// end-add_member_to_multiple_access_groups

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(207))
			Expect(addMembershipMultipleGroupsResponse).ToNot(BeNil())

		})
		It(`RemoveMemberFromAllAccessGroups request example`, func() {
			fmt.Println("\nRemoveMemberFromAllAccessGroups() result:")
			// begin-remove_member_from_all_access_groups

			removeMemberFromAllAccessGroupsOptions := iamAccessGroupsService.NewRemoveMemberFromAllAccessGroupsOptions(
				testAccountID,
				"IBMid-user1",
			)

			deleteFromAllGroupsResponse, response, err := iamAccessGroupsService.RemoveMemberFromAllAccessGroups(removeMemberFromAllAccessGroupsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(deleteFromAllGroupsResponse, "", "  ")
			fmt.Println(string(b))

			// end-remove_member_from_all_access_groups

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(207))
			Expect(deleteFromAllGroupsResponse).ToNot(BeNil())

		})
		It(`AddAccessGroupRule request example`, func() {
			fmt.Println("\nAddAccessGroupRule() result:")
			// begin-add_access_group_rule

			ruleConditionsModel := &iamaccessgroupsv2.RuleConditions{
				Claim:    core.StringPtr("isManager"),
				Operator: core.StringPtr("EQUALS"),
				Value:    core.StringPtr("true"),
			}

			addAccessGroupRuleOptions := iamAccessGroupsService.NewAddAccessGroupRuleOptions(
				accessGroupIDLink,
				int64(12),
				"https://idp.example.org/SAML2a",
				[]iamaccessgroupsv2.RuleConditions{*ruleConditionsModel},
			)
			addAccessGroupRuleOptions.SetName("Manager group rule")

			rule, response, err := iamAccessGroupsService.AddAccessGroupRule(addAccessGroupRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(rule, "", "  ")
			fmt.Println(string(b))

			// end-add_access_group_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(rule).ToNot(BeNil())
			testClaimRuleID = *rule.ID
		})
		It(`GetAccessGroupRule request example`, func() {
			fmt.Println("\nGetAccessGroupRule() result:")
			// begin-get_access_group_rule

			getAccessGroupRuleOptions := iamAccessGroupsService.NewGetAccessGroupRuleOptions(
				accessGroupIDLink,
				testClaimRuleID,
			)

			rule, response, err := iamAccessGroupsService.GetAccessGroupRule(getAccessGroupRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(rule, "", "  ")
			fmt.Println(string(b))

			// end-get_access_group_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())

			testClaimRuleEtag = response.GetHeaders().Get("Etag")
		})
		It(`ReplaceAccessGroupRule request example`, func() {
			fmt.Println("\nReplaceAccessGroupRule() result:")
			// begin-replace_access_group_rule

			ruleConditionsModel := &iamaccessgroupsv2.RuleConditions{
				Claim:    core.StringPtr("isManager"),
				Operator: core.StringPtr("EQUALS"),
				Value:    core.StringPtr("true"),
			}

			replaceAccessGroupRuleOptions := iamAccessGroupsService.NewReplaceAccessGroupRuleOptions(
				accessGroupIDLink,
				testClaimRuleID,
				testClaimRuleEtag,
				int64(12),
				"https://idp.example.org/SAML2",
				[]iamaccessgroupsv2.RuleConditions{*ruleConditionsModel},
			)
			replaceAccessGroupRuleOptions.SetName("Manager group rule")

			rule, response, err := iamAccessGroupsService.ReplaceAccessGroupRule(replaceAccessGroupRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(rule, "", "  ")
			fmt.Println(string(b))

			// end-replace_access_group_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())

		})
		It(`ListAccessGroupRules request example`, func() {
			fmt.Println("\nListAccessGroupRules() result:")
			// begin-list_access_group_rules

			listAccessGroupRulesOptions := iamAccessGroupsService.NewListAccessGroupRulesOptions(
				accessGroupIDLink,
			)

			rulesList, response, err := iamAccessGroupsService.ListAccessGroupRules(listAccessGroupRulesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(rulesList, "", "  ")
			fmt.Println(string(b))

			// end-list_access_group_rules

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rulesList).ToNot(BeNil())

		})
		It(`RemoveAccessGroupRule request example`, func() {
			// begin-remove_access_group_rule

			removeAccessGroupRuleOptions := iamAccessGroupsService.NewRemoveAccessGroupRuleOptions(
				accessGroupIDLink,
				testClaimRuleID,
			)

			response, err := iamAccessGroupsService.RemoveAccessGroupRule(removeAccessGroupRuleOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from RemoveAccessGroupRule(): %d\n", response.StatusCode)
			}

			// end-remove_access_group_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`GetAccountSettings request example`, func() {
			fmt.Println("\nGetAccountSettings() result:")
			// begin-get_account_settings

			getAccountSettingsOptions := iamAccessGroupsService.NewGetAccountSettingsOptions(
				testAccountID,
			)

			accountSettings, response, err := iamAccessGroupsService.GetAccountSettings(getAccountSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accountSettings, "", "  ")
			fmt.Println(string(b))

			// end-get_account_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSettings).ToNot(BeNil())

		})
		It(`UpdateAccountSettings request example`, func() {
			fmt.Println("\nUpdateAccountSettings() result:")
			// begin-update_account_settings

			updateAccountSettingsOptions := iamAccessGroupsService.NewUpdateAccountSettingsOptions(
				testAccountID,
			)
			updateAccountSettingsOptions.SetPublicAccessEnabled(true)

			accountSettings, response, err := iamAccessGroupsService.UpdateAccountSettings(updateAccountSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accountSettings, "", "  ")
			fmt.Println(string(b))

			// end-update_account_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSettings).ToNot(BeNil())

		})
		It(`DeleteAccessGroup request example`, func() {
			// begin-delete_access_group

			deleteAccessGroupOptions := iamAccessGroupsService.NewDeleteAccessGroupOptions(
				accessGroupIDLink,
			)

			response, err := iamAccessGroupsService.DeleteAccessGroup(deleteAccessGroupOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteAccessGroup(): %d\n", response.StatusCode)
			}

			// end-delete_access_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`CreateTemplate request example`, func() {
			fmt.Println("\nCreateTemplate() result:")
			// begin-create_template

			membersActionControlsModel := &iamaccessgroupsv2.MembersActionControls{
				Add:    core.BoolPtr(true),
				Remove: core.BoolPtr(false),
			}

			membersInputModel := &iamaccessgroupsv2.Members{
				Users:          []string{"IBMid-50PJGPKYJJ", "IBMid-665000T8WY"},
				ActionControls: membersActionControlsModel,
			}

			conditionInputModel := &iamaccessgroupsv2.Conditions{
				Claim:    core.StringPtr("blueGroup"),
				Operator: core.StringPtr("CONTAINS"),
				Value:    core.StringPtr(`"test-bluegroup-saml"`),
			}

			rulesActionControlsModel := &iamaccessgroupsv2.RuleActionControls{
				Remove: core.BoolPtr(false),
			}

			ruleInputModel := &iamaccessgroupsv2.AssertionsRule{
				Name:           core.StringPtr("Manager group rule"),
				Expiration:     core.Int64Ptr(int64(12)),
				RealmName:      core.StringPtr("https://idp.example.org/SAML2"),
				Conditions:     []iamaccessgroupsv2.Conditions{*conditionInputModel},
				ActionControls: rulesActionControlsModel,
			}

			assertionsActionControlsModel := &iamaccessgroupsv2.AssertionsActionControls{
				Add:    core.BoolPtr(false),
				Remove: core.BoolPtr(true),
			}

			assertionsInputModel := &iamaccessgroupsv2.Assertions{
				Rules:          []iamaccessgroupsv2.AssertionsRule{*ruleInputModel},
				ActionControls: assertionsActionControlsModel,
			}

			accessActionControlsModel := &iamaccessgroupsv2.AccessActionControls{
				Add: core.BoolPtr(false),
			}

			groupActionControlsModel := &iamaccessgroupsv2.GroupActionControls{
				Access: accessActionControlsModel,
			}

			accessGroupInputModel := &iamaccessgroupsv2.AccessGroupRequest{
				Name:           core.StringPtr("IAM Admin Group"),
				Description:    core.StringPtr("This access group template allows admin access to all IAM platform services in the account."),
				Members:        membersInputModel,
				Assertions:     assertionsInputModel,
				ActionControls: groupActionControlsModel,
			}

			policyTemplatesInputModel := &iamaccessgroupsv2.PolicyTemplates{
				ID:      &testPolicyTemplateID,
				Version: core.StringPtr("1"),
			}

			createTemplateOptions := iamAccessGroupsService.NewCreateTemplateOptions(
				"IAM Admin Group template",
				testAccountID,
			)
			createTemplateOptions.SetDescription("This access group template allows admin access to all IAM platform services in the account.")
			createTemplateOptions.SetGroup(accessGroupInputModel)
			createTemplateOptions.SetPolicyTemplateReferences([]iamaccessgroupsv2.PolicyTemplates{*policyTemplatesInputModel})

			createTemplateResponse, response, err := iamAccessGroupsService.CreateTemplate(createTemplateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createTemplateResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_template
			testTemplateId = *createTemplateResponse.ID

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createTemplateResponse).ToNot(BeNil())
		})
		It(`ListTemplates request example`, func() {
			fmt.Println("\nListTemplates() result:")
			// begin-list_templates

			listTemplatesOptions := &iamaccessgroupsv2.ListTemplatesOptions{
				AccountID:     &testAccountID,
				TransactionID: core.StringPtr("testString"),
				Limit:         core.Int64Ptr(int64(50)),
				Verbose:       core.BoolPtr(true),
			}

			pager, err := iamAccessGroupsService.NewTemplatesPager(listTemplatesOptions)
			if err != nil {
				panic(err)
			}

			var allResults []iamaccessgroupsv2.GroupTemplate
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))

			// end-list_templates
		})
		It(`CreateTemplateVersion request example`, func() {
			fmt.Println("\nCreateTemplateVersion() result:")
			// begin-create_template_version

			membersActionControlsModel := &iamaccessgroupsv2.MembersActionControls{
				Add:    core.BoolPtr(true),
				Remove: core.BoolPtr(false),
			}

			membersInputModel := &iamaccessgroupsv2.Members{
				Users:          []string{"IBMid-50PJGPKYJJ", "IBMid-665000T8WY"},
				ActionControls: membersActionControlsModel,
			}

			conditionInputModel := &iamaccessgroupsv2.Conditions{
				Claim:    core.StringPtr("blueGroup"),
				Operator: core.StringPtr("CONTAINS"),
				Value:    core.StringPtr(`"test-bluegroup-saml"`),
			}

			ruleInputModel := &iamaccessgroupsv2.AssertionsRule{
				Name:       core.StringPtr("Manager group rule"),
				Expiration: core.Int64Ptr(int64(12)),
				RealmName:  core.StringPtr("https://idp.example.org/SAML2"),
				Conditions: []iamaccessgroupsv2.Conditions{*conditionInputModel},
			}

			assertionsActionControlsModel := &iamaccessgroupsv2.AssertionsActionControls{
				Add: core.BoolPtr(false),
			}

			assertionsInputModel := &iamaccessgroupsv2.Assertions{
				Rules:          []iamaccessgroupsv2.AssertionsRule{*ruleInputModel},
				ActionControls: assertionsActionControlsModel,
			}

			accessActionControlsModel := &iamaccessgroupsv2.AccessActionControls{
				Add: core.BoolPtr(false),
			}

			groupActionControlsModel := &iamaccessgroupsv2.GroupActionControls{
				Access: accessActionControlsModel,
			}

			accessGroupInputModel := &iamaccessgroupsv2.AccessGroupRequest{
				Name:           core.StringPtr("IAM Admin Group 8"),
				Description:    core.StringPtr("This access group template allows admin access to all IAM platform services in the account."),
				Members:        membersInputModel,
				Assertions:     assertionsInputModel,
				ActionControls: groupActionControlsModel,
			}

			policyTemplatesInputModel := &iamaccessgroupsv2.PolicyTemplates{
				ID:      &testPolicyTemplateID,
				Version: core.StringPtr("1"),
			}

			createTemplateVersionOptions := iamAccessGroupsService.NewCreateTemplateVersionOptions(
				testTemplateId,
			)
			createTemplateVersionOptions.SetName("IAM Admin Group template 2")
			createTemplateVersionOptions.SetDescription("This access group template allows admin access to all IAM platform services in the account.")
			createTemplateVersionOptions.SetGroup(accessGroupInputModel)
			createTemplateVersionOptions.SetPolicyTemplateReferences([]iamaccessgroupsv2.PolicyTemplates{*policyTemplatesInputModel})

			createTemplateResponse, response, err := iamAccessGroupsService.CreateTemplateVersion(createTemplateVersionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createTemplateResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_template_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createTemplateResponse).ToNot(BeNil())
		})
		It(`ListTemplateVersions request example`, func() {
			fmt.Println("\nListTemplateVersions() result:")
			// begin-list_template_versions

			listTemplateVersionsOptions := &iamaccessgroupsv2.ListTemplateVersionsOptions{
				TemplateID: &testTemplateId,
				Limit:      core.Int64Ptr(int64(100)),
			}

			pager, err := iamAccessGroupsService.NewTemplateVersionsPager(listTemplateVersionsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []iamaccessgroupsv2.ListTemplateVersionResponse
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))

			// end-list_template_versions
		})
		It(`GetTemplateVersion request example`, func() {
			fmt.Println("\nGetTemplateVersion() result:")
			// begin-get_template_version

			getTemplateVersionOptions := iamAccessGroupsService.NewGetTemplateVersionOptions(
				testTemplateId,
				"2",
			)

			createTemplateResponse, response, err := iamAccessGroupsService.GetTemplateVersion(getTemplateVersionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createTemplateResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_template_version
			testTemplateVersionEtag = response.Headers.Get("ETag")
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(createTemplateResponse).ToNot(BeNil())
		})
		It(`UpdateTemplateVersion request example`, func() {
			fmt.Println("\nUpdateTemplateVersion() result:")
			// begin-update_template_version

			membersActionControlsModel := &iamaccessgroupsv2.MembersActionControls{
				Add:    core.BoolPtr(true),
				Remove: core.BoolPtr(false),
			}

			membersInputModel := &iamaccessgroupsv2.Members{
				Users:          []string{"IBMid-665000T8WY"},
				ActionControls: membersActionControlsModel,
			}

			conditionInputModel := &iamaccessgroupsv2.Conditions{
				Claim:    core.StringPtr("blueGroup"),
				Operator: core.StringPtr("CONTAINS"),
				Value:    core.StringPtr(`"test-bluegroup-saml"`),
			}

			rulesActionControlsModel := &iamaccessgroupsv2.RuleActionControls{
				Remove: core.BoolPtr(false),
			}

			ruleInputModel := &iamaccessgroupsv2.AssertionsRule{
				Name:           core.StringPtr("Manager group rule"),
				Expiration:     core.Int64Ptr(int64(12)),
				RealmName:      core.StringPtr("https://idp.example.org/SAML2"),
				Conditions:     []iamaccessgroupsv2.Conditions{*conditionInputModel},
				ActionControls: rulesActionControlsModel,
			}

			assertionsActionControlsModel := &iamaccessgroupsv2.AssertionsActionControls{
				Add: core.BoolPtr(false),
			}

			assertionsInputModel := &iamaccessgroupsv2.Assertions{
				Rules:          []iamaccessgroupsv2.AssertionsRule{*ruleInputModel},
				ActionControls: assertionsActionControlsModel,
			}

			accessActionControlsModel := &iamaccessgroupsv2.AccessActionControls{
				Add: core.BoolPtr(false),
			}

			groupActionControlsModel := &iamaccessgroupsv2.GroupActionControls{
				Access: accessActionControlsModel,
			}

			accessGroupInputModel := &iamaccessgroupsv2.AccessGroupRequest{
				Name:           core.StringPtr("IAM Admin Group 8"),
				Description:    core.StringPtr("This access group template allows admin access to all IAM platform services in the account."),
				Members:        membersInputModel,
				Assertions:     assertionsInputModel,
				ActionControls: groupActionControlsModel,
			}

			policyTemplatesInputModel := &iamaccessgroupsv2.PolicyTemplates{
				ID:      &testPolicyTemplateID,
				Version: core.StringPtr("1"),
			}

			updateTemplateVersionOptions := iamAccessGroupsService.NewUpdateTemplateVersionOptions(
				testTemplateId,
				"2",
				testTemplateVersionEtag,
			)
			updateTemplateVersionOptions.SetName("IAM Admin Group template 2")
			updateTemplateVersionOptions.SetDescription("This access group template allows admin access to all IAM platform services in the account.")
			updateTemplateVersionOptions.SetGroup(accessGroupInputModel)
			updateTemplateVersionOptions.SetPolicyTemplateReferences([]iamaccessgroupsv2.PolicyTemplates{*policyTemplatesInputModel})
			updateTemplateVersionOptions.SetTransactionID("83adf5bd-de790caa3")

			createTemplateResponse, response, err := iamAccessGroupsService.UpdateTemplateVersion(updateTemplateVersionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createTemplateResponse, "", "  ")
			fmt.Println(string(b))

			// end-update_template_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createTemplateResponse).ToNot(BeNil())
		})
		It(`GetLatestTemplateVersion request example`, func() {
			fmt.Println("\nGetLatestTemplateVersion() result:")
			// begin-get_latest_template_version

			getLatestTemplateVersionOptions := iamAccessGroupsService.NewGetLatestTemplateVersionOptions(
				testTemplateId,
			)

			createTemplateResponse, response, err := iamAccessGroupsService.GetLatestTemplateVersion(getLatestTemplateVersionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createTemplateResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_latest_template_version
			testTemplateLatestVersionEtag = response.Headers.Get("ETag")

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(createTemplateResponse).ToNot(BeNil())
		})
		It(`CommitTemplate request example`, func() {
			// begin-commit_template

			commitTemplateOptions := iamAccessGroupsService.NewCommitTemplateOptions(
				testTemplateId,
				"2",
				testTemplateLatestVersionEtag,
			)

			response, err := iamAccessGroupsService.CommitTemplate(commitTemplateOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from CommitTemplate(): %d\n", response.StatusCode)
			}

			// end-commit_template
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`CreateAssignment request example`, func() {
			fmt.Println("\nCreateAssignment() result:")
			// begin-create_assignment

			createAssignmentOptions := iamAccessGroupsService.NewCreateAssignmentOptions(
				testTemplateId,
				"2",
				"AccountGroup",
				testAccountGroupID,
			)

			templateCreateAssignmentResponse, response, err := iamAccessGroupsService.CreateAssignment(createAssignmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(templateCreateAssignmentResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_assignment
			time.Sleep(60 * time.Second)
			testAssignmentID = *templateCreateAssignmentResponse.ID
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(templateCreateAssignmentResponse).ToNot(BeNil())
		})
		It(`ListAssignments request example`, func() {
			fmt.Println("\nListAssignments() result:")
			// begin-list_assignments

			listAssignmentsOptions := iamAccessGroupsService.NewListAssignmentsOptions(
				testAccountID,
			)

			templatesListAssignmentResponse, response, err := iamAccessGroupsService.ListAssignments(listAssignmentsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(templatesListAssignmentResponse, "", "  ")
			fmt.Println(string(b))

			// end-list_assignments
			time.Sleep(60 * time.Second)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templatesListAssignmentResponse).ToNot(BeNil())
		})
		It(`GetAssignment request example`, func() {
			fmt.Println("\nGetAssignment() result:")
			// begin-get_assignment

			getAssignmentOptions := iamAccessGroupsService.NewGetAssignmentOptions(
				testAssignmentID,
			)

			getTemplateAssignmentResponse, response, err := iamAccessGroupsService.GetAssignment(getAssignmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(getTemplateAssignmentResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_assignment
			testAssignmentEtag = response.Headers.Get("ETag")
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getTemplateAssignmentResponse).ToNot(BeNil())
		})

		It(`UpdateAssignment request example`, func() {
			fmt.Println("\nUpdateAssignment() result:")
			// begin-update_assignment

			updateAssignmentOptions := iamAccessGroupsService.NewUpdateAssignmentOptions(
				testAssignmentID,
				testAssignmentEtag,
				"2",
			)

			getTemplateAssignmentResponse, response, err := iamAccessGroupsService.UpdateAssignment(updateAssignmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(getTemplateAssignmentResponse, "", "  ")
			fmt.Println(string(b))

			// end-update_assignment
			time.Sleep(90 * time.Second)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(getTemplateAssignmentResponse).ToNot(BeNil())
		})
		It(`DeleteAssignment request example`, func() {
			// begin-delete_assignment

			deleteAssignmentOptions := iamAccessGroupsService.NewDeleteAssignmentOptions(
				testAssignmentID,
			)

			response, err := iamAccessGroupsService.DeleteAssignment(deleteAssignmentOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 202 {
				fmt.Printf("\nUnexpected response status code received from DeleteAssignment(): %d\n", response.StatusCode)
			}

			// end-delete_assignment
			time.Sleep(90 * time.Second)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
		It(`DeleteTemplateVersion request example`, func() {
			// begin-delete_template_version

			deleteTemplateVersionOptions := iamAccessGroupsService.NewDeleteTemplateVersionOptions(
				testTemplateId,
				"1",
			)

			response, err := iamAccessGroupsService.DeleteTemplateVersion(deleteTemplateVersionOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteTemplateVersion(): %d\n", response.StatusCode)
			}

			// end-delete_template_version

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteTemplate request example`, func() {
			// begin-delete_template

			deleteTemplateOptions := iamAccessGroupsService.NewDeleteTemplateOptions(
				testTemplateId,
			)

			response, err := iamAccessGroupsService.DeleteTemplate(deleteTemplateOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteTemplate(): %d\n", response.StatusCode)
			}

			// end-delete_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})

	})
})
