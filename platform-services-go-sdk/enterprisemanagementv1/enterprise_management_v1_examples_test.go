//go:build examples

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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/enterprisemanagementv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the Enterprise Management service.
//
// The following configuration properties are assumed to be defined:
// ENTERPRISE_MANAGEMENT_URL=<service base url>
// ENTERPRISE_MANAGEMENT_AUTH_TYPE=iam
// ENTERPRISE_MANAGEMENT_APIKEY=<IAM apikey>
// ENTERPRISE_MANAGEMENT_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
// ENTERPRISE_MANAGEMENT_ENTERPRISE_ID=<ID of the enterprise>
// ENTERPRISE_MANAGEMENT_ACCOUNT_ID=<enterprise account ID>
// ENTERPRISE_MANAGEMENT_ACCOUNT_IAM_ID=<IAM ID of the enterprise account>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//

var _ = Describe(`EnterpriseManagementV1 Examples Tests`, func() {
	const externalConfigFile = "../enterprise_management.env"

	var (
		enterpriseManagementService *enterprisemanagementv1.EnterpriseManagementV1
		config                      map[string]string
		configLoaded                bool = false
		enterpriseID                string
		enterpriseAccountID         string
		enterpriseAccountIamID      string

		accountGroupID          string
		newParentAccountGroupID string
		accountID               string
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
			config, err = core.GetServiceProperties(enterprisemanagementv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			configLoaded = len(config) > 0

			enterpriseID = config["ENTERPRISE_ID"]
			Expect(enterpriseID).ToNot(BeEmpty())

			enterpriseAccountID = config["ACCOUNT_ID"]
			Expect(enterpriseAccountID).ToNot(BeEmpty())

			enterpriseAccountIamID = config["ACCOUNT_IAM_ID"]
			Expect(enterpriseAccountIamID).ToNot(BeEmpty())
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			enterpriseManagementServiceOptions := &enterprisemanagementv1.EnterpriseManagementV1Options{}

			enterpriseManagementService, err = enterprisemanagementv1.NewEnterpriseManagementV1UsingExternalConfig(enterpriseManagementServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(enterpriseManagementService).ToNot(BeNil())
		})
	})

	Describe(`EnterpriseManagementV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateAccountGroup request example`, func() {
			fmt.Println("\nCreateAccountGroup() result:")
			parentCRN := "crn:v1:bluemix:public:enterprise::a/" + enterpriseAccountID + "::enterprise:" + enterpriseID

			// begin-create_account_group
			createAccountGroupOptions := enterpriseManagementService.NewCreateAccountGroupOptions(
				parentCRN,
				"Example Account Group",
				enterpriseAccountIamID,
			)

			createAccountGroupResponse, response, err := enterpriseManagementService.CreateAccountGroup(createAccountGroupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createAccountGroupResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_account_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createAccountGroupResponse).ToNot(BeNil())

			accountGroupID = *createAccountGroupResponse.AccountGroupID
		})
		It(`CreateAccountGroup request example (new parent account group)`, func() {
			fmt.Println("\nCreateAccountGroup(<new-parent>) result:")
			parentCRN := "crn:v1:bluemix:public:enterprise::a/" + enterpriseAccountID + "::enterprise:" + enterpriseID
			createAccountGroupOptions := enterpriseManagementService.NewCreateAccountGroupOptions(
				parentCRN,
				"New Parent Account Group",
				enterpriseAccountIamID,
			)

			createAccountGroupResponse, response, err := enterpriseManagementService.CreateAccountGroup(createAccountGroupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createAccountGroupResponse, "", "  ")
			fmt.Println(string(b))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createAccountGroupResponse).ToNot(BeNil())

			newParentAccountGroupID = *createAccountGroupResponse.AccountGroupID
		})
		It(`ListAccountGroups request example`, func() {
			fmt.Println("\nListAccountGroups() result:")
			// begin-list_account_groups
			listAccountGroupsOptions := &enterprisemanagementv1.ListAccountGroupsOptions{
				EnterpriseID: &enterpriseID,
			}

			pager, err := enterpriseManagementService.NewAccountGroupsPager(listAccountGroupsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []enterprisemanagementv1.AccountGroup
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_account_groups
		})
		It(`GetAccountGroup request example`, func() {
			fmt.Println("\nGetAccountGroup() result:")
			// begin-get_account_group

			getAccountGroupOptions := enterpriseManagementService.NewGetAccountGroupOptions(
				accountGroupID,
			)

			accountGroup, response, err := enterpriseManagementService.GetAccountGroup(getAccountGroupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accountGroup, "", "  ")
			fmt.Println(string(b))

			// end-get_account_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountGroup).ToNot(BeNil())
		})
		It(`UpdateAccountGroup request example`, func() {
			// begin-update_account_group

			updateAccountGroupOptions := enterpriseManagementService.NewUpdateAccountGroupOptions(
				accountGroupID,
			)
			updateAccountGroupOptions.SetName("Updated Example Account Group")
			updateAccountGroupOptions.SetPrimaryContactIamID(enterpriseAccountIamID)

			response, err := enterpriseManagementService.UpdateAccountGroup(updateAccountGroupOptions)
			if err != nil {
				panic(err)
			}

			// end-update_account_group

			Expect(err).To(BeNil())
			fmt.Printf("\nUpdateAccountGroup() response status code: %d\n", response.StatusCode)
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteAccountGroup request example`, func() {
			// begin-delete_account_group
			deleteAccountGroupOptions := enterpriseManagementService.NewDeleteAccountGroupOptions(
				accountGroupID,
			)

			response, err := enterpriseManagementService.DeleteAccountGroup(deleteAccountGroupOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_account_group

			Expect(err).To(BeNil())
			fmt.Printf("\nDeleteAccountGroup() response status code: %d\n", response.StatusCode)
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`ImportAccountToEnterprise request example`, func() {
			Skip("Skip by design")
			importAccountID := "<accountid_to_be_imported>"
			// begin-import_account_to_enterprise

			importAccountToEnterpriseOptions := enterpriseManagementService.NewImportAccountToEnterpriseOptions(
				enterpriseID,
				importAccountID,
			)

			response, err := enterpriseManagementService.ImportAccountToEnterprise(importAccountToEnterpriseOptions)
			if err != nil {
				panic(err)
			}

			// end-import_account_to_enterprise

			Expect(err).To(BeNil())
			fmt.Printf("\nImportAccountToEnterprise() response status code: %d\n", response.StatusCode)
			Expect(response.StatusCode).To(Equal(202))
		})
		It(`CreateAccount request example`, func() {
			fmt.Println("\nCreateAccount() result:")
			var parentCRN = "crn:v1:bluemix:public:enterprise::a/" + enterpriseAccountID + "::account-group:" + accountGroupID

			// begin-create_account
			createAccountOptions := enterpriseManagementService.NewCreateAccountOptions(
				parentCRN,
				"Example Account",
				enterpriseAccountIamID,
			)

			createAccountResponse, response, err := enterpriseManagementService.CreateAccount(createAccountOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createAccountResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_account

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(createAccountResponse).ToNot(BeNil())

			accountID = *createAccountResponse.AccountID
		})
		It(`ListAccounts request example`, func() {
			fmt.Println("\nListAccounts() result:")
			// begin-list_accounts
			listAccountsOptions := &enterprisemanagementv1.ListAccountsOptions{
				EnterpriseID: &enterpriseID,
			}

			pager, err := enterpriseManagementService.NewAccountsPager(listAccountsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []enterprisemanagementv1.Account
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_accounts
		})
		It(`GetAccount request example`, func() {
			fmt.Println("\nGetAccount() result:")
			// begin-get_account

			getAccountOptions := enterpriseManagementService.NewGetAccountOptions(
				accountID,
			)

			account, response, err := enterpriseManagementService.GetAccount(getAccountOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(account, "", "  ")
			fmt.Println(string(b))

			// end-get_account

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(account).ToNot(BeNil())
		})
		It(`UpdateAccount request example`, func() {
			newParentCRN := "crn:v1:bluemix:public:enterprise::a/" + enterpriseAccountID + "::account-group:" + newParentAccountGroupID

			// begin-update_account

			updateAccountOptions := enterpriseManagementService.NewUpdateAccountOptions(
				accountID,
				newParentCRN,
			)

			response, err := enterpriseManagementService.UpdateAccount(updateAccountOptions)
			if err != nil {
				panic(err)
			}

			// end-update_account

			Expect(err).To(BeNil())
			fmt.Printf("\nUpdateAccount() response status code: %d\n", response.StatusCode)
			Expect(response.StatusCode).To(Equal(202))
		})
		It(`DeleteAccount request example`, func() {
			fmt.Println("\nDeleteAccount() result:")
			// begin-delete_account

			deleteAccountOptions := enterpriseManagementService.NewDeleteAccountOptions(
				accountID,
			)

			response, err := enterpriseManagementService.DeleteAccount(deleteAccountOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_account

			Expect(err).To(BeNil())
			fmt.Printf("\nDeleteAccount() response status code: %d\n", response.StatusCode)
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`CreateEnterprise request example`, func() {
			Skip("Skip by design")

			srcAccountID := "<standalone_account_id>"
			contactIamID := "<standalone_account_iam_id>"

			fmt.Println("\nCreateEnterprise() result:")
			// begin-create_enterprise

			createEnterpriseOptions := enterpriseManagementService.NewCreateEnterpriseOptions(
				srcAccountID,
				"Example Enterprise",
				contactIamID,
			)

			createEnterpriseResponse, response, err := enterpriseManagementService.CreateEnterprise(createEnterpriseOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createEnterpriseResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_enterprise

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(createEnterpriseResponse).ToNot(BeNil())
		})
		It(`ListEnterprises request example`, func() {
			fmt.Println("\nListEnterprises() result:")
			// begin-list_enterprises
			listEnterprisesOptions := &enterprisemanagementv1.ListEnterprisesOptions{
				EnterpriseAccountID: &enterpriseAccountID,
			}

			pager, err := enterpriseManagementService.NewEnterprisesPager(listEnterprisesOptions)
			if err != nil {
				panic(err)
			}

			var allResults []enterprisemanagementv1.Enterprise
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-list_enterprises
		})
		It(`GetEnterprise request example`, func() {
			fmt.Println("\nGetEnterprise() result:")
			// begin-get_enterprise

			getEnterpriseOptions := enterpriseManagementService.NewGetEnterpriseOptions(
				enterpriseID,
			)

			enterprise, response, err := enterpriseManagementService.GetEnterprise(getEnterpriseOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(enterprise, "", "  ")
			fmt.Println(string(b))

			// end-get_enterprise

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(enterprise).ToNot(BeNil())
		})
		It(`UpdateEnterprise request example`, func() {
			// begin-update_enterprise

			updateEnterpriseOptions := enterpriseManagementService.NewUpdateEnterpriseOptions(
				enterpriseID,
			)
			updateEnterpriseOptions.SetName("Updated Example Enterprise")
			updateEnterpriseOptions.SetPrimaryContactIamID(enterpriseAccountIamID)

			response, err := enterpriseManagementService.UpdateEnterprise(updateEnterpriseOptions)
			if err != nil {
				panic(err)
			}

			// end-update_enterprise

			Expect(err).To(BeNil())
			fmt.Printf("\nUpdateEnterprise() response status code: %d\n", response.StatusCode)
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})
