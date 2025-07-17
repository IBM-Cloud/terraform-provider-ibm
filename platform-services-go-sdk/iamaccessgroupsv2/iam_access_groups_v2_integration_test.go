//go:build integration

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
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iamaccessgroupsv2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("IAM Access Groups - Integration Tests", func() {
	const externalConfigFile = "../iam_access_groups_v2.env"

	var (
		iamAccessGroupsService *iamaccessgroupsv2.IamAccessGroupsV2
		err                    error
		config                 map[string]string
		configLoaded           bool = false

		testAccountID         string
		testGroupName         string = "SDK Test Group - Golang"
		testGroupDescription  string = "This group is used for integration test purposes. It can be deleted at any time."
		testGroupEtag         string
		testGroupID           string
		testUserID            string = "IBMid-" + strconv.Itoa(rand.Intn(100000))
		testClaimRuleID       string
		testClaimRuleEtag     string
		testAccountSettings   *iamaccessgroupsv2.AccountSettings
		testPolicyTemplateID  string
		testTemplateID        string
		testTemplateEtag      string
		testLatestVersionETag string
		testAccountGroupID    string
		testAssignmentID      string
		testAssignmentEtag    string

		userType   string = "user"
		etagHeader string = "Etag"
	)

	var shouldSkipTest = func() {
		if !configLoaded {
			Skip("External configuration is not available, skipping...")
		}
	}

	It("Successfully load the configuration", func() {
		err = os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
		if err != nil {
			Skip("Could not set IBM_CREDENTIALS_FILE environment variable: " + err.Error())
		}

		config, err = core.GetServiceProperties(iamaccessgroupsv2.DefaultServiceName)
		if err == nil {
			testAccountID = config["TEST_ACCOUNT_ID"]
			testPolicyTemplateID = config["TEST_POLICY_TEMPLATE_ID"]
			testAccountGroupID = config["TEST_ACCOUNT_GROUP_ID"]
			if testAccountID != "" {
				configLoaded = true
			}
		}
		if !configLoaded {
			Skip("External configuration could not be loaded, skipping...")
		}
	})

	It(`Successfully created IamAccessGroupsV2 service instance`, func() {
		shouldSkipTest()

		iamAccessGroupsService, err = iamaccessgroupsv2.NewIamAccessGroupsV2UsingExternalConfig(
			&iamaccessgroupsv2.IamAccessGroupsV2Options{},
		)

		Expect(err).To(BeNil())
		Expect(iamAccessGroupsService).ToNot(BeNil())

		core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
		iamAccessGroupsService.EnableRetries(4, 30*time.Second)
	})

	Describe("Create an access group", func() {

		It("Successfully created an access group", func() {
			shouldSkipTest()

			options := iamAccessGroupsService.NewCreateAccessGroupOptions(testAccountID, testGroupName)
			result, detailedResponse, err := iamAccessGroupsService.CreateAccessGroup(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(201))
			Expect(*result.AccountID).To(Equal(testAccountID))
			Expect(*result.Name).To(Equal(testGroupName))

			testGroupID = *result.ID
		})
	})

	Describe("Get an access group", func() {

		It("Successfully retrieved an access group", func() {
			shouldSkipTest()

			options := iamAccessGroupsService.NewGetAccessGroupOptions(testGroupID)
			result, detailedResponse, err := iamAccessGroupsService.GetAccessGroup(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(*result.AccountID).To(Equal(testAccountID))
			Expect(*result.ID).To(Equal(testGroupID))
			Expect(*result.Name).To(Equal(testGroupName))
			Expect(*result.Description).To(Equal(""))

			testGroupEtag = detailedResponse.GetHeaders().Get(etagHeader)
		})
	})

	Describe("Update an access group description", func() {

		It("Successfully updated an access group description", func() {
			shouldSkipTest()

			options := iamAccessGroupsService.NewUpdateAccessGroupOptions(testGroupID, testGroupEtag)
			options.SetDescription(testGroupDescription)
			result, detailedResponse, err := iamAccessGroupsService.UpdateAccessGroup(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(*result.AccountID).To(Equal(testAccountID))
			Expect(*result.Name).To(Equal(testGroupName))
			Expect(*result.ID).To(Equal(testGroupID))
			Expect(*result.Description).To(Equal(testGroupDescription))
		})
	})

	Describe("List access groups", func() {

		It("Successfully listed the account's access groups", func() {
			shouldSkipTest()

			options := iamAccessGroupsService.NewListAccessGroupsOptions(testAccountID)
			options.SetHidePublicAccess(true)
			result, detailedResponse, err := iamAccessGroupsService.ListAccessGroups(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))

			// confirm the test group is present
			testGroupPresent := false
			for _, group := range result.Groups {
				if *group.ID == testGroupID {
					testGroupPresent = true
				}
			}
			Expect(testGroupPresent).To(BeTrue())
		})

		It(`ListAccessGroups(listAccessGroupsOptions *ListAccessGroupsOptions) using AccessGroupsPager`, func() {
			listAccessGroupsOptions := &iamaccessgroupsv2.ListAccessGroupsOptions{
				AccountID:        &testAccountID,
				HidePublicAccess: core.BoolPtr(true),
			}

			// Test GetNext().
			pager, err := iamAccessGroupsService.NewAccessGroupsPager(listAccessGroupsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []iamaccessgroupsv2.Group
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = iamAccessGroupsService.NewAccessGroupsPager(listAccessGroupsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListAccessGroups() returned a total of %d item(s) using AccessGroupsPager.\n", len(allResults))

			// confirm the test group is present
			testGroupPresent := false
			for _, group := range allResults {
				if *group.ID == testGroupID {
					testGroupPresent = true
				}
			}
			Expect(testGroupPresent).To(BeTrue())
		})
	})

	Describe("Add members to an access group", func() {

		It("Successfully added members to an access group", func() {
			shouldSkipTest()

			addMemberItem, err := iamAccessGroupsService.NewAddGroupMembersRequestMembersItem(testUserID, userType)
			Expect(err).To(BeNil())

			options := iamAccessGroupsService.NewAddMembersToAccessGroupOptions(testGroupID)
			options.Members = append(options.Members, *addMemberItem)
			result, detailedResponse, err := iamAccessGroupsService.AddMembersToAccessGroup(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(207))

			// confirm the test user is present
			testUserPresent := false
			for _, member := range result.Members {
				if *member.IamID == testUserID {
					testUserPresent = true
					Expect(*member.Type).To(Equal(userType))
					Expect(*member.StatusCode).To(Equal(int64(200)))
				}
			}
			Expect(testUserPresent).To(BeTrue())

		})

		It("Successfully added member to multiple access groups", func() {
			shouldSkipTest()

			options := iamAccessGroupsService.NewAddMemberToMultipleAccessGroupsOptions(testAccountID, testUserID)
			options.SetType(userType)
			options.Groups = append(options.Groups, testGroupID)
			result, detailedResponse, err := iamAccessGroupsService.AddMemberToMultipleAccessGroups(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(207))

			// confirm the test user is present
			testGroupPresent := false
			for _, group := range result.Groups {
				if *group.AccessGroupID == testGroupID {
					testGroupPresent = true
					Expect(*group.StatusCode).To(Equal(int64(200)))
				}
			}
			Expect(testGroupPresent).To(BeTrue())

		})
	})

	Describe("Check access group membership", func() {

		It("Successfully checked the membership", func() {
			shouldSkipTest()

			options := iamAccessGroupsService.NewIsMemberOfAccessGroupOptions(testGroupID, testUserID)
			detailedResponse, err := iamAccessGroupsService.IsMemberOfAccessGroup(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(204))
		})
	})

	Describe("List access group memberships", func() {

		It("Successfully listed the memberships", func() {
			shouldSkipTest()

			options := iamAccessGroupsService.NewListAccessGroupMembersOptions(testGroupID)
			result, detailedResponse, err := iamAccessGroupsService.ListAccessGroupMembers(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))

			// confirm the test user is present
			testUserPresent := false
			for _, member := range result.Members {
				if *member.IamID == testUserID {
					testUserPresent = true
				}
			}
			Expect(testUserPresent).To(BeTrue())
		})
	})

	Describe("Delete access group membership", func() {

		It("Successfully deleted the membership", func() {
			shouldSkipTest()

			options := iamAccessGroupsService.NewRemoveMemberFromAccessGroupOptions(testGroupID, testUserID)
			detailedResponse, err := iamAccessGroupsService.RemoveMemberFromAccessGroup(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(204))
		})
	})

	Describe("Delete member from all groups", func() {

		It("Returned that the membership was not found", func() {
			shouldSkipTest()

			options := iamAccessGroupsService.NewRemoveMemberFromAllAccessGroupsOptions(testAccountID, testUserID)
			result, detailedResponse, err := iamAccessGroupsService.RemoveMemberFromAllAccessGroups(options)
			Expect(err).To(Not(BeNil()))
			Expect(result).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(404))
		})
	})

	Describe("Delete multiple members from access group", func() {

		It("Returned that the membership was not found", func() {
			shouldSkipTest()

			options := iamAccessGroupsService.NewRemoveMembersFromAccessGroupOptions(testGroupID)
			options.Members = append(options.Members, testUserID)
			result, detailedResponse, err := iamAccessGroupsService.RemoveMembersFromAccessGroup(options)
			Expect(err).To(Not(BeNil()))
			Expect(result).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(404))
		})
	})

	Describe("Create an access group rule", func() {

		It("Successfully created an access group rule", func() {
			shouldSkipTest()

			testExpiration := int64(24)
			condition, err := iamAccessGroupsService.NewRuleConditions("test claim", "EQUALS", "1")
			Expect(err).To(BeNil())

			options := iamAccessGroupsService.NewAddAccessGroupRuleOptions(testGroupID, testExpiration, "test realm name", []iamaccessgroupsv2.RuleConditions{*condition})

			result, detailedResponse, err := iamAccessGroupsService.AddAccessGroupRule(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(201))
			Expect(*result.AccessGroupID).To(Equal(testGroupID))
			Expect(*result.AccountID).To(Equal(testAccountID))
			Expect(*result.Expiration).To(Equal(testExpiration))

			testClaimRuleID = *result.ID
		})
	})

	Describe("Get an access group rule", func() {

		It("Successfully retrieved the rule", func() {
			shouldSkipTest()

			options := iamAccessGroupsService.NewGetAccessGroupRuleOptions(testGroupID, testClaimRuleID)
			result, detailedResponse, err := iamAccessGroupsService.GetAccessGroupRule(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(*result.ID).To(Equal(testClaimRuleID))
			Expect(*result.AccessGroupID).To(Equal(testGroupID))
			Expect(*result.AccountID).To(Equal(testAccountID))

			testClaimRuleEtag = detailedResponse.GetHeaders().Get(etagHeader)
		})
	})

	Describe("List access group rules", func() {

		It("Successfully listed the rules", func() {
			shouldSkipTest()

			options := iamAccessGroupsService.NewListAccessGroupRulesOptions(testGroupID)
			result, detailedResponse, err := iamAccessGroupsService.ListAccessGroupRules(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))

			// confirm the test user is present
			testClaimRulePresent := false
			for _, claimRule := range result.Rules {
				if *claimRule.ID == testClaimRuleID {
					testClaimRulePresent = true
				}
			}
			Expect(testClaimRulePresent).To(BeTrue())
		})
	})

	Describe("Update an access group rule", func() {

		It("Successfully updated the rule", func() {
			shouldSkipTest()

			testExpiration := int64(24)
			condition, err := iamAccessGroupsService.NewRuleConditions("test claim", "EQUALS", "1")
			Expect(err).To(BeNil())

			options := iamAccessGroupsService.NewReplaceAccessGroupRuleOptions(testGroupID, testClaimRuleID, testClaimRuleEtag, testExpiration, "updated test realm name", []iamaccessgroupsv2.RuleConditions{*condition})

			result, detailedResponse, err := iamAccessGroupsService.ReplaceAccessGroupRule(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(*result.ID).To(Equal(testClaimRuleID))
			Expect(*result.AccessGroupID).To(Equal(testGroupID))
			Expect(*result.AccountID).To(Equal(testAccountID))
		})
	})

	Describe("Delete access group rule", func() {

		It("Successfully deleted the rule", func() {
			shouldSkipTest()

			options := iamAccessGroupsService.NewRemoveAccessGroupRuleOptions(testGroupID, testClaimRuleID)
			detailedResponse, err := iamAccessGroupsService.RemoveAccessGroupRule(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(204))
		})
	})

	Describe("Get account settings", func() {

		It("Successfully retrieved the settings", func() {
			shouldSkipTest()

			options := iamAccessGroupsService.NewGetAccountSettingsOptions(testAccountID)
			result, detailedResponse, err := iamAccessGroupsService.GetAccountSettings(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(*result.AccountID).To(Equal(testAccountID))

			testAccountSettings = result
		})
	})

	Describe("Update account settings", func() {

		It("Successfully updated the settings", func() {
			shouldSkipTest()

			options := iamAccessGroupsService.NewUpdateAccountSettingsOptions(testAccountID)
			options.SetPublicAccessEnabled(*testAccountSettings.PublicAccessEnabled)
			result, detailedResponse, err := iamAccessGroupsService.UpdateAccountSettings(options)
			Expect(err).To(BeNil())
			Expect(detailedResponse.StatusCode).To(Equal(200))
			Expect(*result.AccountID).To(Equal(testAccountID))
			Expect(*result.PublicAccessEnabled).To(Equal(*options.PublicAccessEnabled))
		})
	})

	// clean up all test groups
	AfterSuite(func() {
		if !configLoaded {
			return
		}

		// list all groups in the account (minus the public access group)
		options := iamAccessGroupsService.NewListAccessGroupsOptions(testAccountID)
		options.SetHidePublicAccess(true)
		result, detailedResponse, err := iamAccessGroupsService.ListAccessGroups(options)
		Expect(err).To(BeNil())
		Expect(detailedResponse.StatusCode).To(Equal(200))

		// iterate across the groups
		for _, group := range result.Groups {

			// force delete the test group (or any test groups older than 5 minutes)
			if *group.Name == testGroupName {

				createdAt := time.Time(*group.CreatedAt)
				fiveMinutesAgo := time.Now().Add(-(time.Duration(5) * time.Minute))

				if *group.ID == testGroupID || createdAt.Before(fiveMinutesAgo) {
					options := iamAccessGroupsService.NewDeleteAccessGroupOptions(*group.ID)
					options.SetForce(true)
					detailedResponse, err := iamAccessGroupsService.DeleteAccessGroup(options)
					Expect(err).To(BeNil())
					Expect(detailedResponse.StatusCode).To(Equal(204))
				}
			}
		}
	})

	Describe(`CreateTemplate - Create Template`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTemplate(createTemplateOptions *CreateTemplateOptions)`, func() {
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
				Value:    core.StringPtr("\"test-bluegroup-saml\""),
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

			createTemplateOptions := &iamaccessgroupsv2.CreateTemplateOptions{
				Name:                     core.StringPtr("IAM Admin Group template"),
				AccountID:                &testAccountID,
				Description:              core.StringPtr("This access group template allows admin access to all IAM platform services in the account."),
				Group:                    accessGroupInputModel,
				PolicyTemplateReferences: []iamaccessgroupsv2.PolicyTemplates{*policyTemplatesInputModel},
				TransactionID:            core.StringPtr("testString"),
			}

			createTemplateResponse, response, err := iamAccessGroupsService.CreateTemplate(createTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createTemplateResponse).ToNot(BeNil())
			testTemplateID = *createTemplateResponse.ID
		})
	})

	Describe(`ListTemplates - List Templates`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTemplates(listTemplatesOptions *ListTemplatesOptions) with pagination`, func() {
			listTemplatesOptions := &iamaccessgroupsv2.ListTemplatesOptions{
				AccountID:     &testAccountID,
				TransactionID: core.StringPtr("testString"),
				Limit:         core.Int64Ptr(int64(50)),
				Offset:        core.Int64Ptr(int64(0)),
				Verbose:       core.BoolPtr(true),
			}

			listTemplatesOptions.Offset = nil
			listTemplatesOptions.Limit = core.Int64Ptr(1)

			var allResults []iamaccessgroupsv2.GroupTemplate
			for {
				listTemplatesResponse, response, err := iamAccessGroupsService.ListTemplates(listTemplatesOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(listTemplatesResponse).ToNot(BeNil())
				allResults = append(allResults, listTemplatesResponse.GroupTemplates...)

				listTemplatesOptions.Offset, err = listTemplatesResponse.GetNextOffset()
				Expect(err).To(BeNil())

				if listTemplatesOptions.Offset == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListTemplates(listTemplatesOptions *ListTemplatesOptions) using TemplatesPager`, func() {
			listTemplatesOptions := &iamaccessgroupsv2.ListTemplatesOptions{
				AccountID:     &testAccountID,
				TransactionID: core.StringPtr("testString"),
				Limit:         core.Int64Ptr(int64(50)),
				Verbose:       core.BoolPtr(true),
			}

			// Test GetNext().
			pager, err := iamAccessGroupsService.NewTemplatesPager(listTemplatesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []iamaccessgroupsv2.GroupTemplate
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = iamAccessGroupsService.NewTemplatesPager(listTemplatesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListTemplates() returned a total of %d item(s) using TemplatesPager.\n", len(allResults))
		})
	})

	Describe(`CreateTemplateVersion - Create template version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTemplateVersion(createTemplateVersionOptions *CreateTemplateVersionOptions)`, func() {
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
				Value:    core.StringPtr("\"test-bluegroup-saml\""),
			}

			rulesActionControlsModel := &iamaccessgroupsv2.RuleActionControls{
				Remove: core.BoolPtr(true),
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

			createTemplateVersionOptions := &iamaccessgroupsv2.CreateTemplateVersionOptions{
				TemplateID:               &testTemplateID,
				Name:                     core.StringPtr("IAM Admin Group template 2"),
				Description:              core.StringPtr("This access group template allows admin access to all IAM platform services in the account."),
				Group:                    accessGroupInputModel,
				PolicyTemplateReferences: []iamaccessgroupsv2.PolicyTemplates{*policyTemplatesInputModel},
				TransactionID:            core.StringPtr("testString"),
			}

			createTemplateResponse, response, err := iamAccessGroupsService.CreateTemplateVersion(createTemplateVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createTemplateResponse).ToNot(BeNil())
		})
	})

	Describe(`ListTemplateVersions - List template versions`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTemplateVersions(listTemplateVersionsOptions *ListTemplateVersionsOptions) with pagination`, func() {
			listTemplateVersionsOptions := &iamaccessgroupsv2.ListTemplateVersionsOptions{
				TemplateID: &testTemplateID,
				Limit:      core.Int64Ptr(int64(100)),
				Offset:     core.Int64Ptr(int64(0)),
			}

			listTemplateVersionsOptions.Offset = nil
			listTemplateVersionsOptions.Limit = core.Int64Ptr(1)

			var allResults []iamaccessgroupsv2.ListTemplateVersionResponse
			for {
				listTemplateVersionsResponse, response, err := iamAccessGroupsService.ListTemplateVersions(listTemplateVersionsOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(listTemplateVersionsResponse).ToNot(BeNil())
				allResults = append(allResults, listTemplateVersionsResponse.GroupTemplateVersions...)

				listTemplateVersionsOptions.Offset, err = listTemplateVersionsResponse.GetNextOffset()
				Expect(err).To(BeNil())

				if listTemplateVersionsOptions.Offset == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`ListTemplateVersions(listTemplateVersionsOptions *ListTemplateVersionsOptions) using TemplateVersionsPager`, func() {
			listTemplateVersionsOptions := &iamaccessgroupsv2.ListTemplateVersionsOptions{
				TemplateID: &testTemplateID,
				Limit:      core.Int64Ptr(int64(100)),
			}

			// Test GetNext().
			pager, err := iamAccessGroupsService.NewTemplateVersionsPager(listTemplateVersionsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []iamaccessgroupsv2.ListTemplateVersionResponse
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = iamAccessGroupsService.NewTemplateVersionsPager(listTemplateVersionsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListTemplateVersions() returned a total of %d item(s) using TemplateVersionsPager.\n", len(allResults))
		})
	})

	Describe(`GetTemplateVersion - Get template version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTemplateVersion(getTemplateVersionOptions *GetTemplateVersionOptions)`, func() {
			getTemplateVersionOptions := &iamaccessgroupsv2.GetTemplateVersionOptions{
				TemplateID:    &testTemplateID,
				VersionNum:    core.StringPtr("1"),
				TransactionID: core.StringPtr("testString"),
			}

			createTemplateResponse, response, err := iamAccessGroupsService.GetTemplateVersion(getTemplateVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(createTemplateResponse).ToNot(BeNil())
			testTemplateEtag = response.GetHeaders().Get(etagHeader)

		})
	})

	Describe(`UpdateTemplateVersion - Update template version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateTemplateVersion(updateTemplateVersionOptions *UpdateTemplateVersionOptions)`, func() {
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
				Value:    core.StringPtr("\"test-bluegroup-saml\""),
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

			updateTemplateVersionOptions := &iamaccessgroupsv2.UpdateTemplateVersionOptions{
				TemplateID:               &testTemplateID,
				IfMatch:                  &testTemplateEtag,
				VersionNum:               core.StringPtr("1"),
				Name:                     core.StringPtr("IAM Admin Group template 2"),
				Description:              core.StringPtr("This access group template allows admin access to all IAM platform services in the account."),
				Group:                    accessGroupInputModel,
				PolicyTemplateReferences: []iamaccessgroupsv2.PolicyTemplates{*policyTemplatesInputModel},
			}

			createTemplateResponse, response, err := iamAccessGroupsService.UpdateTemplateVersion(updateTemplateVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createTemplateResponse).ToNot(BeNil())
		})
	})

	Describe(`GetLatestTemplateVersion - Get latest template version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetLatestTemplateVersion(getLatestTemplateVersionOptions *GetLatestTemplateVersionOptions)`, func() {
			getLatestTemplateVersionOptions := &iamaccessgroupsv2.GetLatestTemplateVersionOptions{
				TemplateID:    &testTemplateID,
				TransactionID: core.StringPtr("testString"),
			}

			createTemplateResponse, response, err := iamAccessGroupsService.GetLatestTemplateVersion(getLatestTemplateVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(createTemplateResponse).ToNot(BeNil())
			testLatestVersionETag = response.GetHeaders().Get(etagHeader)

		})
	})

	Describe(`CommitTemplate - Commit a template`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CommitTemplate(commitTemplateOptions *CommitTemplateOptions)`, func() {
			commitTemplateOptions := &iamaccessgroupsv2.CommitTemplateOptions{
				TemplateID:    &testTemplateID,
				VersionNum:    core.StringPtr("2"),
				IfMatch:       &testLatestVersionETag,
				TransactionID: core.StringPtr("testString"),
			}

			response, err := iamAccessGroupsService.CommitTemplate(commitTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`CreateAssignment - Create assignment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateAssignment(createAssignmentOptions *CreateAssignmentOptions)`, func() {
			createAssignmentOptions := &iamaccessgroupsv2.CreateAssignmentOptions{
				TemplateID:      &testTemplateID,
				TemplateVersion: core.StringPtr("2"),
				TargetType:      core.StringPtr("AccountGroup"),
				Target:          &testAccountGroupID,
				TransactionID:   core.StringPtr("testString"),
			}

			templateCreateAssignmentResponse, response, err := iamAccessGroupsService.CreateAssignment(createAssignmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(templateCreateAssignmentResponse).ToNot(BeNil())
			testAssignmentID = *templateCreateAssignmentResponse.ID
			time.Sleep(90 * time.Second)

		})
	})

	Describe(`ListAssignments - List Assignments`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListAssignments(listAssignmentsOptions *ListAssignmentsOptions)`, func() {
			listAssignmentsOptions := &iamaccessgroupsv2.ListAssignmentsOptions{
				AccountID: &testAccountID,
			}

			templatesListAssignmentResponse, response, err := iamAccessGroupsService.ListAssignments(listAssignmentsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(templatesListAssignmentResponse).ToNot(BeNil())
		})
	})

	Describe(`GetAssignment - Get assignment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAssignment(getAssignmentOptions *GetAssignmentOptions)`, func() {
			getAssignmentOptions := &iamaccessgroupsv2.GetAssignmentOptions{
				AssignmentID:  &testAssignmentID,
				TransactionID: core.StringPtr("testString"),
			}

			getTemplateAssignmentResponse, response, err := iamAccessGroupsService.GetAssignment(getAssignmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getTemplateAssignmentResponse).ToNot(BeNil())
			testAssignmentEtag = response.GetHeaders().Get(etagHeader)
		})
	})

	Describe(`UpdateAssignment - Update Assignment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateAssignment(updateAssignmentOptions *UpdateAssignmentOptions)`, func() {
			updateAssignmentOptions := &iamaccessgroupsv2.UpdateAssignmentOptions{
				AssignmentID:    &testAssignmentID,
				IfMatch:         &testAssignmentEtag,
				TemplateVersion: core.StringPtr("2"),
			}

			getTemplateAssignmentResponse, response, err := iamAccessGroupsService.UpdateAssignment(updateAssignmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(getTemplateAssignmentResponse).ToNot(BeNil())
			time.Sleep(90 * time.Second)
		})
	})

	Describe(`DeleteAssignment - Delete assignment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteAssignment(deleteAssignmentOptions *DeleteAssignmentOptions)`, func() {
			deleteAssignmentOptions := &iamaccessgroupsv2.DeleteAssignmentOptions{
				AssignmentID:  &testAssignmentID,
				TransactionID: core.StringPtr("testString"),
			}

			response, err := iamAccessGroupsService.DeleteAssignment(deleteAssignmentOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			time.Sleep(90 * time.Second)
		})
	})

	Describe(`DeleteTemplateVersion - Delete template version`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteTemplateVersion(deleteTemplateVersionOptions *DeleteTemplateVersionOptions)`, func() {
			deleteTemplateVersionOptions := &iamaccessgroupsv2.DeleteTemplateVersionOptions{
				TemplateID:    &testTemplateID,
				VersionNum:    core.StringPtr("2"),
				TransactionID: core.StringPtr("testString"),
			}

			response, err := iamAccessGroupsService.DeleteTemplateVersion(deleteTemplateVersionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteTemplate - Delete template`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteTemplate(deleteTemplateOptions *DeleteTemplateOptions)`, func() {
			deleteTemplateOptions := &iamaccessgroupsv2.DeleteTemplateOptions{
				TemplateID:    &testTemplateID,
				TransactionID: core.StringPtr("testString"),
			}

			response, err := iamAccessGroupsService.DeleteTemplate(deleteTemplateOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})
