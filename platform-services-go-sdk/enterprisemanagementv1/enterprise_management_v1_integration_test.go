//go:build integration

/**
 * (C) Copyright IBM Corp. 2021, 2022.
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

package enterprisemanagementv1_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/enterprisemanagementv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the enterprisemanagementv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`EnterpriseManagementV1 Integration Tests`, func() {

	const externalConfigFile = "../enterprise_management.env"

	var (
		err                         error
		enterpriseManagementService *enterprisemanagementv1.EnterpriseManagementV1
		serviceURL                  string
		testConfig                  map[string]string

		enterpriseID string
		accountID    string
		accountIamID string

		accountGroupName          = "Example Account Group"
		accountGroupID            string
		updatedAccountGroupName   = "Updated Example Account Group"
		newParentAccountGroupName = "Second Example Account Group"
		newParentAccountGroupID   string

		exampleAccountName = "Example Account Name"
		exampleAccountID   string

		updatedEnterpriseName = "Updated Enterprise Name"
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			testConfig, err = core.GetServiceProperties(enterprisemanagementv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = testConfig["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			enterpriseID = testConfig["ENTERPRISE_ID"]
			Expect(enterpriseID).ToNot(BeEmpty())

			accountID = testConfig["ACCOUNT_ID"]
			Expect(accountID).ToNot(BeEmpty())

			accountIamID = testConfig["ACCOUNT_IAM_ID"]
			Expect(accountIamID).NotTo(BeEmpty())

			fmt.Fprintf(GinkgoWriter, "Service URL: %s\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			enterpriseManagementServiceOptions := &enterprisemanagementv1.EnterpriseManagementV1Options{}

			enterpriseManagementService, err = enterprisemanagementv1.NewEnterpriseManagementV1UsingExternalConfig(enterpriseManagementServiceOptions)

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			enterpriseManagementService.EnableRetries(4, 30*time.Second)

			Expect(err).To(BeNil())
			Expect(enterpriseManagementService).ToNot(BeNil())
			Expect(enterpriseManagementService.Service.Options.URL).To(Equal(serviceURL))
		})
	})

	Describe(`CreateAccountGroup - Create an account group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateAccountGroup(createAccountGroupOptions *CreateAccountGroupOptions)`, func() {
			var parentCRN = "crn:v1:bluemix:public:enterprise::a/" + accountID + "::enterprise:" + enterpriseID
			createAccountGroupOptions := &enterprisemanagementv1.CreateAccountGroupOptions{
				Parent:              &parentCRN,
				Name:                &accountGroupName,
				PrimaryContactIamID: &accountIamID,
			}

			accountGroupResponse, response, err := enterpriseManagementService.CreateAccountGroup(createAccountGroupOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(accountGroupResponse).ToNot(BeNil())

			accountGroupID = *accountGroupResponse.AccountGroupID
		})

		It(`CreateAccountGroup(createAccountGroupOptions *CreateAccountGroupOptions) - new parent account group`, func() {
			var parentCRN = "crn:v1:bluemix:public:enterprise::a/" + accountID + "::enterprise:" + enterpriseID
			createAccountGroupOptions := &enterprisemanagementv1.CreateAccountGroupOptions{
				Parent:              &parentCRN,
				Name:                &newParentAccountGroupName,
				PrimaryContactIamID: &accountIamID,
			}

			accountGroupResponse, response, err := enterpriseManagementService.CreateAccountGroup(createAccountGroupOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(accountGroupResponse).ToNot(BeNil())

			newParentAccountGroupID = *accountGroupResponse.AccountGroupID
		})
	})

	Describe(`ListAccountGroups`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListAccountGroups(listAccountGroupsOptions *ListAccountGroupsOptions)`, func() {

			accountGroupList := []enterprisemanagementv1.AccountGroup{}
			var moreResults = true
			var nextDocid string
			var resultPerPage = int64(10)

			for moreResults {

				listAccountGroupsOptions := &enterprisemanagementv1.ListAccountGroupsOptions{
					EnterpriseID: &enterpriseID,
					Limit:        &resultPerPage,
					NextDocid:    &nextDocid,
				}

				listAccountGroupsResponse, response, err := enterpriseManagementService.ListAccountGroups(listAccountGroupsOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(listAccountGroupsResponse).ToNot(BeNil())

				accountGroupList = append(accountGroupList, listAccountGroupsResponse.Resources...)

				if listAccountGroupsResponse.NextURL != nil {
					docId, errDocId := core.GetQueryParam(listAccountGroupsResponse.NextURL, "next_docid")
					Expect(errDocId).To(BeNil())
					nextDocid = *docId
				} else {
					moreResults = false
				}
			}

			var foundAccountGroup = false
			for i := range accountGroupList {
				if *accountGroupList[i].ID == accountGroupID {
					foundAccountGroup = true
				}
			}
			Expect(foundAccountGroup).To(BeTrue())

			var foundParentAccountGroup = false
			for i := range accountGroupList {
				if *accountGroupList[i].ID == newParentAccountGroupID {
					foundParentAccountGroup = true
				}
			}
			Expect(foundParentAccountGroup).To(BeTrue())

			fmt.Printf("Received a total of %d account groups.", len(accountGroupList))
		})
		It(`ListAccountGroups(listAccountGroupsOptions *ListAccountGroupsOptions) using AccountGroupsPager`, func() {
			listAccountGroupsOptions := &enterprisemanagementv1.ListAccountGroupsOptions{
				EnterpriseID: &enterpriseID,
			}

			// Test GetNext().
			pager, err := enterpriseManagementService.NewAccountGroupsPager(listAccountGroupsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []enterprisemanagementv1.AccountGroup
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = enterpriseManagementService.NewAccountGroupsPager(listAccountGroupsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListAccountGroups() returned a total of %d item(s) using AccountGroupsPager.\n", len(allResults))
		})
	})

	Describe(`GetAccountGroup - Get account group by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAccountGroup(getAccountGroupOptions *GetAccountGroupOptions)`, func() {

			getAccountGroupOptions := &enterprisemanagementv1.GetAccountGroupOptions{
				AccountGroupID: &accountGroupID,
			}

			accountGroupResponse, response, err := enterpriseManagementService.GetAccountGroup(getAccountGroupOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountGroupResponse).ToNot(BeNil())
			Expect(*accountGroupResponse.ID).To(Equal(accountGroupID))

		})
	})

	Describe(`UpdateAccountGroup - Update an account group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateAccountGroup(updateAccountGroupOptions *UpdateAccountGroupOptions)`, func() {

			updateAccountGroupOptions := &enterprisemanagementv1.UpdateAccountGroupOptions{
				AccountGroupID:      &accountGroupID,
				Name:                &updatedAccountGroupName,
				PrimaryContactIamID: &accountIamID,
			}

			response, err := enterpriseManagementService.UpdateAccountGroup(updateAccountGroupOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteAccountGroup - Delete an account group from the enterprise`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteAccountGroup(deleteAccountGroupOptions *DeleteAccountGroupOptions)`, func() {

			deleteAccountGroupOptions := &enterprisemanagementv1.DeleteAccountGroupOptions{
				AccountGroupID: &accountGroupID,
			}

			response, err := enterpriseManagementService.DeleteAccountGroup(deleteAccountGroupOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`CreateAccount - Create a new account in an enterprise`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateAccount(createAccountOptions *CreateAccountOptions)`, func() {
			var parentCRN = "crn:v1:bluemix:public:enterprise::a/" + accountID + "::account-group:" + accountGroupID
			createAccountRequestTraitsModel := &enterprisemanagementv1.CreateAccountRequestTraits{
				Mfa:                  core.StringPtr(""),
				EnterpriseIamManaged: core.BoolPtr(true),
			}

			createAccountRequestOptionsModel := &enterprisemanagementv1.CreateAccountRequestOptions{
				CreateIamServiceIDWithApikeyAndOwnerPolicies: core.BoolPtr(false),
			}
			createAccountOptions := &enterprisemanagementv1.CreateAccountOptions{
				Parent:     &parentCRN,
				Name:       &exampleAccountName,
				OwnerIamID: &accountIamID,
				Traits:     createAccountRequestTraitsModel,
				Options:    createAccountRequestOptionsModel,
			}

			createAccountResponse, response, err := enterpriseManagementService.CreateAccount(createAccountOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(createAccountResponse).ToNot(BeNil())

			exampleAccountID = *createAccountResponse.AccountID

		})
	})

	Describe(`ListAccounts - List accounts with pagination`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListAccounts(listAccountsOptions *ListAccountsOptions)`, func() {

			var moreResults = true
			var accountsList = []enterprisemanagementv1.Account{}
			var resultPerPage = int64(5)
			var nextDocid string

			for moreResults {
				listAccountsOptions := &enterprisemanagementv1.ListAccountsOptions{
					Limit:          &resultPerPage,
					NextDocid:      &nextDocid,
					AccountGroupID: &accountGroupID,
				}

				listAccountsResponse, response, err := enterpriseManagementService.ListAccounts(listAccountsOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(listAccountsResponse).ToNot(BeNil())

				accountsList = append(accountsList, listAccountsResponse.Resources...)

				if listAccountsResponse.NextURL != nil {
					docID, errDocId := core.GetQueryParam(listAccountsResponse.NextURL, "next_docid")
					Expect(errDocId).To(BeNil())
					nextDocid = *docID
				} else {
					moreResults = false
				}
			}
			var foundExampleAccount = false
			for i := range accountsList {
				if *accountsList[i].ID == exampleAccountID {
					foundExampleAccount = true
				}
			}
			Expect(foundExampleAccount).To(BeTrue())

			fmt.Printf("Received a total of %d accounts.", len(accountsList))
		})
		It(`ListAccounts(listAccountsOptions *ListAccountsOptions) using AccountsPager`, func() {
			listAccountsOptions := &enterprisemanagementv1.ListAccountsOptions{
				AccountGroupID: &accountGroupID,
			}

			// Test GetNext().
			pager, err := enterpriseManagementService.NewAccountsPager(listAccountsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []enterprisemanagementv1.Account
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = enterpriseManagementService.NewAccountsPager(listAccountsOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListAccounts() returned a total of %d item(s) using AccountsPager.\n", len(allResults))
		})
	})

	Describe(`GetAccount - Get account by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAccount(getAccountOptions *GetAccountOptions)`, func() {

			getAccountOptions := &enterprisemanagementv1.GetAccountOptions{
				AccountID: &exampleAccountID,
			}

			account, response, err := enterpriseManagementService.GetAccount(getAccountOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(account).ToNot(BeNil())
			Expect(*account.ID).To(Equal(exampleAccountID))

		})
	})

	Describe(`UpdateAccount - Move an account within the enterprise`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateAccount(updateAccountOptions *UpdateAccountOptions)`, func() {

			var newParentCRN = "crn:v1:bluemix:public:enterprise::a/" + accountID + "::account-group:" + newParentAccountGroupID
			updateAccountOptions := &enterprisemanagementv1.UpdateAccountOptions{
				AccountID: &exampleAccountID,
				Parent:    &newParentCRN,
			}

			response, err := enterpriseManagementService.UpdateAccount(updateAccountOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})
	})

	Describe(`DeleteAccount - Remove an account from its enterprise`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteAccount(deleteAccountOptions *DeleteAccountOptions)`, func() {

			deleteAccountOptions := &enterprisemanagementv1.DeleteAccountOptions{
				AccountID: &exampleAccountID,
			}

			response, err := enterpriseManagementService.DeleteAccount(deleteAccountOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`ListEnterprises - List enterprises`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListEnterprises(listEnterprisesOptions *ListEnterprisesOptions)`, func() {

			var moreResults = true
			var enterpriseList = []enterprisemanagementv1.Enterprise{}
			var resultPerPage = int64(10)
			var nextDocid string

			for moreResults {
				listEnterprisesOptions := &enterprisemanagementv1.ListEnterprisesOptions{
					AccountID: &accountID,
					Limit:     &resultPerPage,
					NextDocid: &nextDocid,
				}

				listEnterprisesResponse, response, err := enterpriseManagementService.ListEnterprises(listEnterprisesOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(listEnterprisesResponse).ToNot(BeNil())

				enterpriseList = append(enterpriseList, listEnterprisesResponse.Resources...)

				if listEnterprisesResponse.NextURL != nil {
					docID, docErr := core.GetQueryParam(listEnterprisesResponse.NextURL, "next_docid")
					Expect(docErr).To(BeNil())
					nextDocid = *docID
				} else {
					moreResults = false
				}
			}
			var found = false
			for i := range enterpriseList {
				if *enterpriseList[i].ID == enterpriseID {
					found = true
				}
			}
			Expect(found).To(BeTrue())
			fmt.Printf("Received a total of %d enterprises.", len(enterpriseList))
		})
		It(`ListEnterprises(listEnterprisesOptions *ListEnterprisesOptions) using EnterprisesPager`, func() {
			listEnterprisesOptions := &enterprisemanagementv1.ListEnterprisesOptions{
				AccountID: &accountID,
			}

			// Test GetNext().
			pager, err := enterpriseManagementService.NewEnterprisesPager(listEnterprisesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []enterprisemanagementv1.Enterprise
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = enterpriseManagementService.NewEnterprisesPager(listEnterprisesOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "ListEnterprises() returned a total of %d item(s) using EnterprisesPager.\n", len(allResults))
		})
	})

	Describe(`GetEnterprise - Get enterprise by ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetEnterprise(getEnterpriseOptions *GetEnterpriseOptions)`, func() {

			getEnterpriseOptions := &enterprisemanagementv1.GetEnterpriseOptions{
				EnterpriseID: &enterpriseID,
			}

			getEnterpriseResponse, response, err := enterpriseManagementService.GetEnterprise(getEnterpriseOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(getEnterpriseResponse).ToNot(BeNil())
			Expect(*getEnterpriseResponse.ID).To(Equal(enterpriseID))

		})
	})

	Describe(`UpdateEnterprise - Update an enterprise`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateEnterprise(updateEnterpriseOptions *UpdateEnterpriseOptions)`, func() {

			updateEnterpriseOptions := &enterprisemanagementv1.UpdateEnterpriseOptions{
				EnterpriseID:        &enterpriseID,
				Name:                &updatedEnterpriseName,
				PrimaryContactIamID: &accountIamID,
			}

			response, err := enterpriseManagementService.UpdateEnterprise(updateEnterpriseOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})
