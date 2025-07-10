//go:build examples

/**
 * (C) Copyright IBM Corp. 2024.
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

package usagereportsv4_test

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/usagereportsv4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// This file provides an example of how to use the Usage Reports service.
//
// The following configuration properties are assumed to be defined:
// USAGE_REPORTS_URL=<service url>
// USAGE_REPORTS_AUTHTYPE=iam
// USAGE_REPORTS_APIKEY=<IAM api key of user with authority to create rules>
// USAGE_REPORTS_AUTH_URL=<IAM token service URL - omit this if using the production environment>
// USAGE_REPORTS_ACCOUNT_ID=<the id of the account whose usage info will be retrieved>
// USAGE_REPORTS_RESOURCE_GROUP_ID=<the id of the resource group whose usage info will be retrieved>
// USAGE_REPORTS_ORG_ID=<the id of the organization whose usage info will be retrieved>
// USAGE_REPORTS_BILLING_MONTH=<the billing month (yyyy-mm) for which usage info will be retrieved>
// USAGE_REPORTS_COS_BUCKET=<The name of the COS bucket to store the snapshot of the billing reports.>
// USAGE_REPORTS_COS_LOCATION=<Region of the COS instance.>
// USAGE_REPORTS_DATE_FROM=<Timestamp in milliseconds for which billing report snapshot is requested.>
// USAGE_REPORTS_DATE_TO=<Timestamp in milliseconds for which billing report snapshot is requested.>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
var _ = Describe(`UsageReportsV4 Examples Tests`, func() {

	const externalConfigFile = "../usage_reports.env"

	var (
		usageReportsService *usagereportsv4.UsageReportsV4
		config              map[string]string
		configLoaded        bool = false

		accountID       string
		resourceGroupID string
		orgID           string
		billingMonth    string
		cosBucket       string
		cosLocation     string
		dateFrom        string
		dateTo          string
	)

	var shouldSkipTest = func() {
		if !configLoaded {
			Skip("External configuration is not available, skipping examples...")
		}
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping examples: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(usagereportsv4.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			accountID = config["ACCOUNT_ID"]
			Expect(accountID).ToNot(BeEmpty())

			resourceGroupID = config["RESOURCE_GROUP_ID"]
			Expect(resourceGroupID).ToNot(BeEmpty())

			orgID = config["ORG_ID"]
			Expect(orgID).ToNot(BeEmpty())

			billingMonth = config["BILLING_MONTH"]
			Expect(billingMonth).ToNot(BeEmpty())

			cosBucket = config["COS_BUCKET"]
			Expect(cosBucket).ToNot(BeEmpty())

			cosLocation = config["COS_LOCATION"]
			Expect(cosLocation).ToNot(BeEmpty())

			dateFrom = config["DATE_FROM"]
			Expect(dateFrom).ToNot(BeEmpty())

			dateTo = config["DATE_TO"]
			Expect(dateTo).ToNot(BeEmpty())

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

			usageReportsServiceOptions := &usagereportsv4.UsageReportsV4Options{}

			usageReportsService, err = usagereportsv4.NewUsageReportsV4UsingExternalConfig(usageReportsServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(usageReportsService).ToNot(BeNil())
		})
	})

	Describe(`UsageReportsV4 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAccountSummary request example`, func() {
			fmt.Println("\nGetAccountSummary() result:")
			// begin-get_account_summary

			getAccountSummaryOptions := usageReportsService.NewGetAccountSummaryOptions(
				accountID,
				billingMonth,
			)

			accountSummary, response, err := usageReportsService.GetAccountSummary(getAccountSummaryOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accountSummary, "", "  ")
			fmt.Println(string(b))

			// end-get_account_summary

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSummary).ToNot(BeNil())
		})
		It(`GetAccountUsage request example`, func() {
			fmt.Println("\nGetAccountUsage() result:")
			// begin-get_account_usage

			getAccountUsageOptions := usageReportsService.NewGetAccountUsageOptions(
				accountID,
				billingMonth,
			)

			accountUsage, response, err := usageReportsService.GetAccountUsage(getAccountUsageOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accountUsage, "", "  ")
			fmt.Println(string(b))

			// end-get_account_usage

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountUsage).ToNot(BeNil())
		})
		It(`GetResourceGroupUsage request example`, func() {
			fmt.Println("\nGetResourceGroupUsage() result:")
			// begin-get_resource_group_usage

			getResourceGroupUsageOptions := usageReportsService.NewGetResourceGroupUsageOptions(
				accountID,
				resourceGroupID,
				billingMonth,
			)

			resourceGroupUsage, response, err := usageReportsService.GetResourceGroupUsage(getResourceGroupUsageOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(resourceGroupUsage, "", "  ")
			fmt.Println(string(b))

			// end-get_resource_group_usage

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceGroupUsage).ToNot(BeNil())
		})
		It(`GetResourceUsageAccount request example`, func() {
			fmt.Println("\nGetResourceUsageAccount() result:")
			// begin-get_resource_usage_account
			getResourceUsageAccountOptions := &usagereportsv4.GetResourceUsageAccountOptions{
				AccountID:    core.StringPtr(accountID),
				Billingmonth: core.StringPtr(billingMonth),
				Names:        core.BoolPtr(true),
				Tags:         core.BoolPtr(true),
			}

			pager, err := usageReportsService.NewGetResourceUsageAccountPager(getResourceUsageAccountOptions)
			if err != nil {
				panic(err)
			}

			var allResults []usagereportsv4.InstanceUsage
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-get_resource_usage_account
		})
		It(`GetResourceUsageResourceGroup request example`, func() {
			fmt.Println("\nGetResourceUsageResourceGroup() result:")
			// begin-get_resource_usage_resource_group
			getResourceUsageResourceGroupOptions := &usagereportsv4.GetResourceUsageResourceGroupOptions{
				AccountID:       core.StringPtr(accountID),
				ResourceGroupID: core.StringPtr(resourceGroupID),
				Billingmonth:    core.StringPtr(billingMonth),
				Names:           core.BoolPtr(true),
				Tags:            core.BoolPtr(true),
				Limit:           core.Int64Ptr(int64(30)),
			}

			pager, err := usageReportsService.NewGetResourceUsageResourceGroupPager(getResourceUsageResourceGroupOptions)
			if err != nil {
				panic(err)
			}

			var allResults []usagereportsv4.InstanceUsage
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-get_resource_usage_resource_group
		})
		It(`GetResourceUsageOrg request example`, func() {
			fmt.Println("\nGetResourceUsageOrg() result:")
			// begin-get_resource_usage_org
			getResourceUsageOrgOptions := &usagereportsv4.GetResourceUsageOrgOptions{
				AccountID:      core.StringPtr(accountID),
				OrganizationID: core.StringPtr(orgID),
				Billingmonth:   core.StringPtr(billingMonth),
				Names:          core.BoolPtr(true),
				Tags:           core.BoolPtr(true),
				Limit:          core.Int64Ptr(int64(30)),
			}

			pager, err := usageReportsService.NewGetResourceUsageOrgPager(getResourceUsageOrgOptions)
			if err != nil {
				panic(err)
			}

			var allResults []usagereportsv4.InstanceUsage
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-get_resource_usage_org
		})
		It(`GetOrgUsage request example`, func() {
			fmt.Println("\nGetOrgUsage() result:")
			// begin-get_org_usage

			getOrgUsageOptions := usageReportsService.NewGetOrgUsageOptions(
				accountID,
				orgID,
				billingMonth,
			)

			orgUsage, response, err := usageReportsService.GetOrgUsage(getOrgUsageOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(orgUsage, "", "  ")
			fmt.Println(string(b))

			// end-get_org_usage

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(orgUsage).ToNot(BeNil())
		})
		It(`CreateReportsSnapshotConfig request example`, func() {
			fmt.Println("\nCreateReportsSnapshotConfig() result:")
			// begin-create_reports_snapshot_config

			createReportsSnapshotConfigOptions := usageReportsService.NewCreateReportsSnapshotConfigOptions(
				accountID,
				"daily",
				cosBucket,
				cosLocation,
			)

			snapshotConfig, response, err := usageReportsService.CreateReportsSnapshotConfig(createReportsSnapshotConfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(snapshotConfig, "", "  ")
			fmt.Println(string(b))

			// end-create_reports_snapshot_config

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(snapshotConfig).ToNot(BeNil())
		})
		It(`GetReportsSnapshotConfig request example`, func() {
			fmt.Println("\nGetReportsSnapshotConfig() result:")
			// begin-get_reports_snapshot_config

			getReportsSnapshotConfigOptions := usageReportsService.NewGetReportsSnapshotConfigOptions(
				accountID,
			)

			snapshotConfig, response, err := usageReportsService.GetReportsSnapshotConfig(getReportsSnapshotConfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(snapshotConfig, "", "  ")
			fmt.Println(string(b))

			// end-get_reports_snapshot_config

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(snapshotConfig).ToNot(BeNil())
		})
		It(`UpdateReportsSnapshotConfig request example`, func() {
			fmt.Println("\nUpdateReportsSnapshotConfig() result:")
			// begin-update_reports_snapshot_config

			updateReportsSnapshotConfigOptions := usageReportsService.NewUpdateReportsSnapshotConfigOptions(
				accountID,
			)

			snapshotConfig, response, err := usageReportsService.UpdateReportsSnapshotConfig(updateReportsSnapshotConfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(snapshotConfig, "", "  ")
			fmt.Println(string(b))

			// end-update_reports_snapshot_config

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(snapshotConfig).ToNot(BeNil())
		})
		It(`ValidateReportsSnapshotConfig request example`, func() {
			fmt.Println("\nValidateReportsSnapshotConfig() result:")
			// begin-validate_reports_snapshot_config

			validateReportsSnapshotConfigOptions := usageReportsService.NewValidateReportsSnapshotConfigOptions(
				accountID,
			)

			snapshotConfigValidateResponse, response, err := usageReportsService.ValidateReportsSnapshotConfig(validateReportsSnapshotConfigOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(snapshotConfigValidateResponse, "", "  ")
			fmt.Println(string(b))

			// end-validate_reports_snapshot_config

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(snapshotConfigValidateResponse).ToNot(BeNil())
		})
		It(`GetReportsSnapshot request example`, func() {
			fmt.Println("\nGetReportsSnapshot() result:")
			from, _ := strconv.ParseInt(dateFrom, 10, 64)
			to, _ := strconv.ParseInt(dateTo, 10, 64)
			// begin-get_reports_snapshot
			getReportsSnapshotOptions := &usagereportsv4.GetReportsSnapshotOptions{
				AccountID: core.StringPtr(accountID),
				Month:     core.StringPtr(billingMonth),
				DateFrom:  core.Int64Ptr(from),
				DateTo:    core.Int64Ptr(to),
				Limit:     core.Int64Ptr(int64(30)),
			}

			pager, err := usageReportsService.NewGetReportsSnapshotPager(getReportsSnapshotOptions)
			if err != nil {
				panic(err)
			}

			var allResults []usagereportsv4.SnapshotListSnapshotsItem
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-get_reports_snapshot
		})
		It(`DeleteReportsSnapshotConfig request example`, func() {
			// begin-delete_reports_snapshot_config

			deleteReportsSnapshotConfigOptions := usageReportsService.NewDeleteReportsSnapshotConfigOptions(
				accountID,
			)

			response, err := usageReportsService.DeleteReportsSnapshotConfig(deleteReportsSnapshotConfigOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteReportsSnapshotConfig(): %d\n", response.StatusCode)
			}

			// end-delete_reports_snapshot_config

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})
