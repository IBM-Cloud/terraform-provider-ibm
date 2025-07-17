//go:build integration

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
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/platform-services-go-sdk/common"
	"github.com/IBM/platform-services-go-sdk/usagereportsv4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the usagereportsv4 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`UsageReportsV4 Integration Tests`, func() {

	const externalConfigFile = "../usage_reports.env"

	var (
		err                 error
		usageReportsService *usagereportsv4.UsageReportsV4
		serviceURL          string
		config              map[string]string

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
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(usagereportsv4.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)

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

			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			usageReportsServiceOptions := &usagereportsv4.UsageReportsV4Options{}

			usageReportsService, err = usagereportsv4.NewUsageReportsV4UsingExternalConfig(usageReportsServiceOptions)
			Expect(err).To(BeNil())
			Expect(usageReportsService).ToNot(BeNil())
			Expect(usageReportsService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			usageReportsService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`GetAccountSummary - Get account summary`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAccountSummary(getAccountSummaryOptions *GetAccountSummaryOptions)`, func() {
			getAccountSummaryOptions := &usagereportsv4.GetAccountSummaryOptions{
				AccountID:    &accountID,
				Billingmonth: &billingMonth,
			}

			accountSummary, response, err := usageReportsService.GetAccountSummary(getAccountSummaryOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSummary).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "\nGetAccountSummary response:\n%s", common.ToJSON(accountSummary))

			Expect(*accountSummary.AccountID).To(Equal(accountID))
			Expect(accountSummary.Offers).ToNot(BeEmpty())
			Expect(accountSummary.Subscription).ToNot(BeNil())
		})
	})

	Describe(`GetAccountUsage - Get account usage`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAccountUsage(getAccountUsageOptions *GetAccountUsageOptions)`, func() {
			getAccountUsageOptions := &usagereportsv4.GetAccountUsageOptions{
				AccountID:      &accountID,
				Billingmonth:   &billingMonth,
				Names:          core.BoolPtr(true),
				AcceptLanguage: core.StringPtr("English"),
			}

			accountUsage, response, err := usageReportsService.GetAccountUsage(getAccountUsageOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountUsage).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "\nGetAccountUsage response:\n%s", common.ToJSON(accountUsage))

			Expect(*accountUsage.AccountID).To(Equal(accountID))
			Expect(*accountUsage.Month).To(Equal(billingMonth))
			Expect(accountUsage.Resources).ToNot(BeEmpty())
		})
	})

	Describe(`GetResourceGroupUsage - Get resource group usage`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetResourceGroupUsage(getResourceGroupUsageOptions *GetResourceGroupUsageOptions)`, func() {
			getResourceGroupUsageOptions := &usagereportsv4.GetResourceGroupUsageOptions{
				AccountID:       &accountID,
				ResourceGroupID: &resourceGroupID,
				Billingmonth:    &billingMonth,
				Names:           core.BoolPtr(true),
			}

			resourceGroupUsage, response, err := usageReportsService.GetResourceGroupUsage(getResourceGroupUsageOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(resourceGroupUsage).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "\nGetResourceGroupUsage response:\n%s", common.ToJSON(resourceGroupUsage))

			Expect(*resourceGroupUsage.AccountID).To(Equal(accountID))
			Expect(*resourceGroupUsage.Month).To(Equal(billingMonth))
			Expect(resourceGroupUsage.Resources).ToNot(BeEmpty())
		})
	})

	Describe(`GetResourceUsageAccount - Get resource instance usage in an account`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetResourceUsageAccount(getResourceUsageAccountOptions *GetResourceUsageAccountOptions)`, func() {

			// Retrieve results one page at a time.
			getResourceUsageAccountOptions := &usagereportsv4.GetResourceUsageAccountOptions{
				AccountID:    &accountID,
				Billingmonth: &billingMonth,
				Names:        core.BoolPtr(true),
				Limit:        core.Int64Ptr(50),
			}

			var results []usagereportsv4.InstanceUsage = make([]usagereportsv4.InstanceUsage, 0)
			var offset *string = nil
			var moreResults bool = true

			for moreResults {
				// Set "Start" parameter for next page of results.
				getResourceUsageAccountOptions.Start = offset

				instancesUsage, response, err := usageReportsService.GetResourceUsageAccount(getResourceUsageAccountOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(instancesUsage).ToNot(BeNil())

				// Add the just-retrieved page to the results.
				if len(instancesUsage.Resources) > 0 {
					results = append(results, instancesUsage.Resources...)
				}

				// Determine offset for next page of results.
				if instancesUsage.Next != nil {
					offset = instancesUsage.Next.Offset
				} else {
					offset = nil
					moreResults = false
				}
			}

			fmt.Fprintf(GinkgoWriter, "\nGetResourceUsageAccount response contained %d total resources.", len(results))
			Expect(results).ToNot(BeEmpty())
		})
		It(`GetResourceUsageAccount(getResourceUsageAccountOptions *GetResourceUsageAccountOptions) with pagination`, func() {
			getResourceUsageAccountOptions := &usagereportsv4.GetResourceUsageAccountOptions{
				AccountID:    core.StringPtr(accountID),
				Billingmonth: core.StringPtr(billingMonth),
				Names:        core.BoolPtr(true),
				Tags:         core.BoolPtr(true),
				Limit:        core.Int64Ptr(int64(30)),
			}

			getResourceUsageAccountOptions.Start = nil

			var allResults []usagereportsv4.InstanceUsage
			for {
				instancesUsage, response, err := usageReportsService.GetResourceUsageAccount(getResourceUsageAccountOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(instancesUsage).ToNot(BeNil())
				allResults = append(allResults, instancesUsage.Resources...)

				getResourceUsageAccountOptions.Start, err = instancesUsage.GetNextStart()
				Expect(err).To(BeNil())

				if getResourceUsageAccountOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`GetResourceUsageAccount(getResourceUsageAccountOptions *GetResourceUsageAccountOptions) using GetResourceUsageAccountPager`, func() {
			getResourceUsageAccountOptions := &usagereportsv4.GetResourceUsageAccountOptions{
				AccountID:    core.StringPtr(accountID),
				Billingmonth: core.StringPtr(billingMonth),
				Names:        core.BoolPtr(true),
				Tags:         core.BoolPtr(true),
				Limit:        core.Int64Ptr(int64(30)),
			}

			// Test GetNext().
			pager, err := usageReportsService.NewGetResourceUsageAccountPager(getResourceUsageAccountOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []usagereportsv4.InstanceUsage
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = usageReportsService.NewGetResourceUsageAccountPager(getResourceUsageAccountOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "GetResourceUsageAccount() returned a total of %d item(s) using GetResourceUsageAccountPager.\n", len(allResults))
		})
	})

	Describe(`GetResourceUsageResourceGroup - Get resource instance usage in a resource group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetResourceUsageResourceGroup(getResourceUsageResourceGroupOptions *GetResourceUsageResourceGroupOptions)`, func() {

			getResourceUsageResourceGroupOptions := &usagereportsv4.GetResourceUsageResourceGroupOptions{
				AccountID:       &accountID,
				ResourceGroupID: &resourceGroupID,
				Billingmonth:    &billingMonth,
				Names:           core.BoolPtr(true),
				Limit:           core.Int64Ptr(50),
			}

			var results []usagereportsv4.InstanceUsage = make([]usagereportsv4.InstanceUsage, 0)
			var offset *string = nil
			var moreResults bool = true

			for moreResults {
				// Set "Start" parameter for next page of results.
				getResourceUsageResourceGroupOptions.Start = offset

				instancesUsage, response, err := usageReportsService.GetResourceUsageResourceGroup(getResourceUsageResourceGroupOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(instancesUsage).ToNot(BeNil())

				// Add the just-retrieved page to the results.
				if len(instancesUsage.Resources) > 0 {
					results = append(results, instancesUsage.Resources...)
				}

				// Determine offset for next page of results.
				if instancesUsage.Next != nil {
					offset = instancesUsage.Next.Offset
				} else {
					offset = nil
					moreResults = false
				}
			}

			fmt.Fprintf(GinkgoWriter, "\nGetResourceUsageResourceGroup response contained %d total resources.", len(results))
			Expect(results).ToNot(BeEmpty())
		})

		It(`GetResourceUsageResourceGroup(getResourceUsageResourceGroupOptions *GetResourceUsageResourceGroupOptions) with pagination`, func() {
			getResourceUsageResourceGroupOptions := &usagereportsv4.GetResourceUsageResourceGroupOptions{
				AccountID:       &accountID,
				ResourceGroupID: &resourceGroupID,
				Billingmonth:    &billingMonth,
				Names:           core.BoolPtr(true),
				Limit:           core.Int64Ptr(50),
			}

			getResourceUsageResourceGroupOptions.Start = nil

			var allResults []usagereportsv4.InstanceUsage
			for {
				instancesUsage, response, err := usageReportsService.GetResourceUsageResourceGroup(getResourceUsageResourceGroupOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(instancesUsage).ToNot(BeNil())
				allResults = append(allResults, instancesUsage.Resources...)

				getResourceUsageResourceGroupOptions.Start, err = instancesUsage.GetNextStart()
				Expect(err).To(BeNil())

				if getResourceUsageResourceGroupOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`GetResourceUsageResourceGroup(getResourceUsageResourceGroupOptions *GetResourceUsageResourceGroupOptions) using GetResourceUsageResourceGroupPager`, func() {
			getResourceUsageResourceGroupOptions := &usagereportsv4.GetResourceUsageResourceGroupOptions{
				AccountID:       &accountID,
				ResourceGroupID: &resourceGroupID,
				Billingmonth:    &billingMonth,
				Names:           core.BoolPtr(true),
				Limit:           core.Int64Ptr(50),
			}

			// Test GetNext().
			pager, err := usageReportsService.NewGetResourceUsageResourceGroupPager(getResourceUsageResourceGroupOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []usagereportsv4.InstanceUsage
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = usageReportsService.NewGetResourceUsageResourceGroupPager(getResourceUsageResourceGroupOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "GetResourceUsageResourceGroup() returned a total of %d item(s) using GetResourceUsageResourceGroupPager.\n", len(allResults))
		})
	})

	Describe(`GetResourceUsageOrg - Get resource instance usage in an organization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetResourceUsageOrg(getResourceUsageOrgOptions *GetResourceUsageOrgOptions)`, func() {

			getResourceUsageOrgOptions := &usagereportsv4.GetResourceUsageOrgOptions{
				AccountID:      &accountID,
				OrganizationID: &orgID,
				Billingmonth:   &billingMonth,
				Names:          core.BoolPtr(true),
				Limit:          core.Int64Ptr(50),
			}

			var results []usagereportsv4.InstanceUsage = make([]usagereportsv4.InstanceUsage, 0)
			var offset *string = nil
			var moreResults bool = true

			for moreResults {
				// Set "Start" parameter for next page of results.
				getResourceUsageOrgOptions.Start = offset

				instancesUsage, response, err := usageReportsService.GetResourceUsageOrg(getResourceUsageOrgOptions)

				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(instancesUsage).ToNot(BeNil())

				// Add the just-retrieved page to the results.
				if len(instancesUsage.Resources) > 0 {
					results = append(results, instancesUsage.Resources...)
				}

				// Determine offset for next page of results.
				if instancesUsage.Next != nil {
					offset = instancesUsage.Next.Offset
				} else {
					offset = nil
					moreResults = false
				}
			}

			fmt.Fprintf(GinkgoWriter, "\nGetResourceUsageOrg response contained %d total resources.", len(results))
			fmt.Fprintf(GinkgoWriter, "\nGetResourceUsageOrg response: %s\n", common.ToJSON(results))
			Expect(results).ToNot(BeEmpty())
		})
		It(`GetResourceUsageOrg(getResourceUsageOrgOptions *GetResourceUsageOrgOptions) with pagination`, func() {
			getResourceUsageOrgOptions := &usagereportsv4.GetResourceUsageOrgOptions{
				AccountID:      &accountID,
				OrganizationID: &orgID,
				Billingmonth:   &billingMonth,
				Names:          core.BoolPtr(true),
				Limit:          core.Int64Ptr(50),
			}

			getResourceUsageOrgOptions.Start = nil

			var allResults []usagereportsv4.InstanceUsage
			for {
				instancesUsage, response, err := usageReportsService.GetResourceUsageOrg(getResourceUsageOrgOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(instancesUsage).ToNot(BeNil())
				allResults = append(allResults, instancesUsage.Resources...)

				getResourceUsageOrgOptions.Start, err = instancesUsage.GetNextStart()
				Expect(err).To(BeNil())

				if getResourceUsageOrgOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`GetResourceUsageOrg(getResourceUsageOrgOptions *GetResourceUsageOrgOptions) using GetResourceUsageOrgPager`, func() {
			getResourceUsageOrgOptions := &usagereportsv4.GetResourceUsageOrgOptions{
				AccountID:      &accountID,
				OrganizationID: &orgID,
				Billingmonth:   &billingMonth,
				Names:          core.BoolPtr(true),
				Limit:          core.Int64Ptr(50),
			}

			// Test GetNext().
			pager, err := usageReportsService.NewGetResourceUsageOrgPager(getResourceUsageOrgOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []usagereportsv4.InstanceUsage
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = usageReportsService.NewGetResourceUsageOrgPager(getResourceUsageOrgOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "GetResourceUsageOrg() returned a total of %d item(s) using GetResourceUsageOrgPager.\n", len(allResults))
		})
	})

	Describe(`GetOrgUsage - Get organization usage`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetOrgUsage(getOrgUsageOptions *GetOrgUsageOptions)`, func() {
			getOrgUsageOptions := &usagereportsv4.GetOrgUsageOptions{
				AccountID:      &accountID,
				OrganizationID: &orgID,
				Billingmonth:   &billingMonth,
				Names:          core.BoolPtr(true),
			}

			orgUsage, response, err := usageReportsService.GetOrgUsage(getOrgUsageOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(orgUsage).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "\nGetOrgUsage response:\n%s", common.ToJSON(orgUsage))

			Expect(*orgUsage.AccountID).To(Equal(accountID))
			Expect(*orgUsage.Month).To(Equal(billingMonth))
			// Expect(orgUsage.Resources).ToNot(BeEmpty())
		})
	})

	Describe(`CreateReportsSnapshotConfig - Setup the snapshot configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateReportsSnapshotConfig(createReportsSnapshotConfigOptions *CreateReportsSnapshotConfigOptions)`, func() {
			createReportsSnapshotConfigOptions := &usagereportsv4.CreateReportsSnapshotConfigOptions{
				AccountID:        core.StringPtr(accountID),
				Interval:         core.StringPtr("daily"),
				CosBucket:        core.StringPtr(cosBucket),
				CosLocation:      core.StringPtr(cosLocation),
				CosReportsFolder: core.StringPtr("IBMCloud-Billing-Reports"),
				ReportTypes:      []string{"account_summary", "enterprise_summary", "account_resource_instance_usage"},
				Versioning:       core.StringPtr("new"),
			}

			snapshotConfig, response, err := usageReportsService.CreateReportsSnapshotConfig(createReportsSnapshotConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(snapshotConfig).ToNot(BeNil())
		})
	})

	Describe(`GetReportsSnapshotConfig - Fetch the snapshot configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReportsSnapshotConfig(getReportsSnapshotConfigOptions *GetReportsSnapshotConfigOptions)`, func() {
			getReportsSnapshotConfigOptions := &usagereportsv4.GetReportsSnapshotConfigOptions{
				AccountID: core.StringPtr(accountID),
			}

			snapshotConfig, response, err := usageReportsService.GetReportsSnapshotConfig(getReportsSnapshotConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(snapshotConfig).ToNot(BeNil())
		})
	})

	Describe(`UpdateReportsSnapshotConfig - Update the snapshot configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateReportsSnapshotConfig(updateReportsSnapshotConfigOptions *UpdateReportsSnapshotConfigOptions)`, func() {
			updateReportsSnapshotConfigOptions := &usagereportsv4.UpdateReportsSnapshotConfigOptions{
				AccountID:        core.StringPtr(accountID),
				Interval:         core.StringPtr("daily"),
				CosBucket:        core.StringPtr(cosBucket),
				CosLocation:      core.StringPtr(cosLocation),
				CosReportsFolder: core.StringPtr("IBMCloud-Billing-Reports"),
				ReportTypes:      []string{"account_summary", "enterprise_summary", "account_resource_instance_usage"},
				Versioning:       core.StringPtr("new"),
			}

			snapshotConfig, response, err := usageReportsService.UpdateReportsSnapshotConfig(updateReportsSnapshotConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(snapshotConfig).ToNot(BeNil())
		})
	})

	Describe(`ValidateReportsSnapshotConfig - Verify billing to COS authorization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ValidateReportsSnapshotConfig(validateReportsSnapshotConfigOptions *ValidateReportsSnapshotConfigOptions)`, func() {
			validateReportsSnapshotConfigOptions := &usagereportsv4.ValidateReportsSnapshotConfigOptions{
				AccountID:        core.StringPtr(accountID),
				Interval:         core.StringPtr("daily"),
				CosBucket:        core.StringPtr(cosBucket),
				CosLocation:      core.StringPtr(cosLocation),
				CosReportsFolder: core.StringPtr("IBMCloud-Billing-Reports"),
				ReportTypes:      []string{"account_summary", "enterprise_summary", "account_resource_instance_usage"},
				Versioning:       core.StringPtr("new"),
			}

			snapshotConfigValidateResponse, response, err := usageReportsService.ValidateReportsSnapshotConfig(validateReportsSnapshotConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(snapshotConfigValidateResponse).ToNot(BeNil())
		})
	})

	Describe(`GetReportsSnapshot - Fetch the current or past snapshots`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetReportsSnapshot(getReportsSnapshotOptions *GetReportsSnapshotOptions) with pagination`, func() {
			getReportsSnapshotOptions := &usagereportsv4.GetReportsSnapshotOptions{
				AccountID: &accountID,
				Month:     &billingMonth,
				Limit:     core.Int64Ptr(int64(30)),
				Start:     core.StringPtr("testString"),
			}
			from, err := strconv.ParseInt(dateFrom, 10, 64)
			if err != nil {
				panic(err)
			}
			to, err := strconv.ParseInt(dateTo, 10, 64)
			if err != nil {
				panic(err)
			}
			getReportsSnapshotOptions.SetDateFrom(from)
			getReportsSnapshotOptions.SetDateTo(to)

			getReportsSnapshotOptions.Start = nil

			var allResults []usagereportsv4.SnapshotListSnapshotsItem
			for {
				snapshotList, response, err := usageReportsService.GetReportsSnapshot(getReportsSnapshotOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(snapshotList).ToNot(BeNil())
				allResults = append(allResults, snapshotList.Snapshots...)

				getReportsSnapshotOptions.Start, err = snapshotList.GetNextStart()
				Expect(err).To(BeNil())

				if getReportsSnapshotOptions.Start == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`GetReportsSnapshot(getReportsSnapshotOptions *GetReportsSnapshotOptions) using GetReportsSnapshotPager`, func() {
			getReportsSnapshotOptions := &usagereportsv4.GetReportsSnapshotOptions{
				AccountID: &accountID,
				Month:     &billingMonth,
				Limit:     core.Int64Ptr(int64(30)),
			}
			from, err := strconv.ParseInt(dateFrom, 10, 64)
			if err != nil {
				panic(err)
			}
			to, err := strconv.ParseInt(dateTo, 10, 64)
			if err != nil {
				panic(err)
			}
			getReportsSnapshotOptions.SetDateFrom(from)
			getReportsSnapshotOptions.SetDateTo(to)

			// Test GetNext().
			pager, err := usageReportsService.NewGetReportsSnapshotPager(getReportsSnapshotOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []usagereportsv4.SnapshotListSnapshotsItem
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = usageReportsService.NewGetReportsSnapshotPager(getReportsSnapshotOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "GetReportsSnapshot() returned a total of %d item(s) using GetReportsSnapshotPager.\n", len(allResults))
		})
		It(`GetReportsSnapshot(getReportsSnapshotOptions *GetReportsSnapshotOptions)`, func() {
			getReportsSnapshotOptions := &usagereportsv4.GetReportsSnapshotOptions{
				AccountID: &accountID,
				Month:     &billingMonth,
			}
			from, err := strconv.ParseInt(dateFrom, 10, 64)
			if err != nil {
				panic(err)
			}
			to, err := strconv.ParseInt(dateTo, 10, 64)
			if err != nil {
				panic(err)
			}
			getReportsSnapshotOptions.SetDateFrom(from)
			getReportsSnapshotOptions.SetDateTo(to)

			snapshotList, response, err := usageReportsService.GetReportsSnapshot(getReportsSnapshotOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(snapshotList).ToNot(BeNil())
		})
	})

	Describe(`DeleteReportsSnapshotConfig - Delete the snapshot configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteReportsSnapshotConfig(deleteReportsSnapshotConfigOptions *DeleteReportsSnapshotConfigOptions)`, func() {
			deleteReportsSnapshotConfigOptions := &usagereportsv4.DeleteReportsSnapshotConfigOptions{
				AccountID: core.StringPtr(accountID),
			}

			response, err := usageReportsService.DeleteReportsSnapshotConfig(deleteReportsSnapshotConfigOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
