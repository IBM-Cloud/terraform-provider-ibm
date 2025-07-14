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

package partnermanagementv1_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/partnermanagementv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the partnermanagementv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`PartnerManagementV1 Integration Tests`, func() {
	const externalConfigFile = "../partner_management_v1.env"

	var (
		err                      error
		partnerManagementService *partnermanagementv1.PartnerManagementV1
		serviceURL               string
		config                   map[string]string

		partnerId    string
		resellerId   string
		customerId   string
		billingMonth string
		viewpoint    string
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
			config, err = core.GetServiceProperties(partnermanagementv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)

			partnerId = config["PARTNER_ID"]
			Expect(partnerId).ToNot(BeEmpty())

			resellerId = config["RESELLER_ID"]
			Expect(resellerId).ToNot(BeEmpty())

			customerId = config["CUSTOMER_ID"]
			Expect(customerId).ToNot(BeEmpty())

			billingMonth = config["BILLING_MONTH"]
			Expect(billingMonth).ToNot(BeEmpty())

			viewpoint = config["VIEWPOINT"]
			Expect(viewpoint).ToNot(BeEmpty())

			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			partnerManagementServiceOptions := &partnermanagementv1.PartnerManagementV1Options{}

			partnerManagementService, err = partnermanagementv1.NewPartnerManagementV1UsingExternalConfig(partnerManagementServiceOptions)
			Expect(err).To(BeNil())
			Expect(partnerManagementService).ToNot(BeNil())
			Expect(partnerManagementService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			partnerManagementService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`GetResourceUsageReport -  Get rolled up usage report across all end customers and resellers`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetResourceUsageReport(getResourceUsageReportOptions *GetResourceUsageReportOptions) with pagination`, func() {
			getResourceUsageReportOptions := &partnermanagementv1.GetResourceUsageReportOptions{
				PartnerID: &partnerId,
				Month:     &billingMonth,
				Limit:     core.Int64Ptr(int64(30)),
			}

			getResourceUsageReportOptions.Offset = nil
			getResourceUsageReportOptions.Limit = core.Int64Ptr(1)

			var allResults []partnermanagementv1.PartnerUsageReport
			for {
				partnerUsageReportSummary, response, err := partnerManagementService.GetResourceUsageReport(getResourceUsageReportOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(partnerUsageReportSummary).ToNot(BeNil())
				allResults = append(allResults, partnerUsageReportSummary.Reports...)

				getResourceUsageReportOptions.Offset, err = partnerUsageReportSummary.GetNextOffset()
				Expect(err).To(BeNil())

				if getResourceUsageReportOptions.Offset == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`GetResourceUsageReport(getResourceUsageReportOptions *GetResourceUsageReportOptions) using GetResourceUsageReportPager`, func() {
			getResourceUsageReportOptions := &partnermanagementv1.GetResourceUsageReportOptions{
				PartnerID: &partnerId,
				Month:     &billingMonth,
				Limit:     core.Int64Ptr(int64(30)),
			}

			// Test GetNext().
			pager, err := partnerManagementService.NewGetResourceUsageReportPager(getResourceUsageReportOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []partnermanagementv1.PartnerUsageReport
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = partnerManagementService.NewGetResourceUsageReportPager(getResourceUsageReportOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "GetResourceUsageReport() returned a total of %d item(s) using GetResourceUsageReportPager.\n", len(allResults))
		})
	})

	Describe(`GetResourceUsageReport -  Get rolled up usage reports by reseller for partner`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetResourceUsageReport(getResourceUsageReportOptions *GetResourceUsageReportOptions) with pagination`, func() {
			getResourceUsageReportOptions := &partnermanagementv1.GetResourceUsageReportOptions{
				PartnerID: &partnerId,
				Month:     &billingMonth,
				Children:  core.BoolPtr(true),
				Limit:     core.Int64Ptr(int64(30)),
			}

			getResourceUsageReportOptions.Offset = nil
			getResourceUsageReportOptions.Limit = core.Int64Ptr(1)

			var allResults []partnermanagementv1.PartnerUsageReport
			for {
				partnerUsageReportSummary, response, err := partnerManagementService.GetResourceUsageReport(getResourceUsageReportOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(partnerUsageReportSummary).ToNot(BeNil())
				allResults = append(allResults, partnerUsageReportSummary.Reports...)

				getResourceUsageReportOptions.Offset, err = partnerUsageReportSummary.GetNextOffset()
				Expect(err).To(BeNil())

				if getResourceUsageReportOptions.Offset == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`GetResourceUsageReport(getResourceUsageReportOptions *GetResourceUsageReportOptions) using GetResourceUsageReportPager`, func() {
			getResourceUsageReportOptions := &partnermanagementv1.GetResourceUsageReportOptions{
				PartnerID: &partnerId,
				Month:     &billingMonth,
				Children:  core.BoolPtr(true),
				Limit:     core.Int64Ptr(int64(30)),
			}

			// Test GetNext().
			pager, err := partnerManagementService.NewGetResourceUsageReportPager(getResourceUsageReportOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []partnermanagementv1.PartnerUsageReport
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = partnerManagementService.NewGetResourceUsageReportPager(getResourceUsageReportOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "GetResourceUsageReport() returned a total of %d item(s) using GetResourceUsageReportPager.\n", len(allResults))
		})
	})

	Describe(`GetResourceUsageReport -  Get usage report of a specific reseller for partner`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetResourceUsageReport(getResourceUsageReportOptions *GetResourceUsageReportOptions) with pagination`, func() {
			getResourceUsageReportOptions := &partnermanagementv1.GetResourceUsageReportOptions{
				PartnerID:  &partnerId,
				ResellerID: &resellerId,
				Month:      &billingMonth,
				Limit:      core.Int64Ptr(int64(30)),
			}

			getResourceUsageReportOptions.Offset = nil
			getResourceUsageReportOptions.Limit = core.Int64Ptr(1)

			var allResults []partnermanagementv1.PartnerUsageReport
			for {
				partnerUsageReportSummary, response, err := partnerManagementService.GetResourceUsageReport(getResourceUsageReportOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(partnerUsageReportSummary).ToNot(BeNil())
				allResults = append(allResults, partnerUsageReportSummary.Reports...)

				getResourceUsageReportOptions.Offset, err = partnerUsageReportSummary.GetNextOffset()
				Expect(err).To(BeNil())

				if getResourceUsageReportOptions.Offset == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`GetResourceUsageReport(getResourceUsageReportOptions *GetResourceUsageReportOptions) using GetResourceUsageReportPager`, func() {
			getResourceUsageReportOptions := &partnermanagementv1.GetResourceUsageReportOptions{
				PartnerID:  &partnerId,
				ResellerID: &resellerId,
				Month:      &billingMonth,
				Limit:      core.Int64Ptr(int64(30)),
			}

			// Test GetNext().
			pager, err := partnerManagementService.NewGetResourceUsageReportPager(getResourceUsageReportOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []partnermanagementv1.PartnerUsageReport
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = partnerManagementService.NewGetResourceUsageReportPager(getResourceUsageReportOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "GetResourceUsageReport() returned a total of %d item(s) using GetResourceUsageReportPager.\n", len(allResults))
		})
	})

	Describe(`GetResourceUsageReport -  Get usage reports of a specific end_customer for partner`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetResourceUsageReport(getResourceUsageReportOptions *GetResourceUsageReportOptions) with pagination`, func() {
			getResourceUsageReportOptions := &partnermanagementv1.GetResourceUsageReportOptions{
				PartnerID:  &partnerId,
				CustomerID: &customerId,
				Month:      &billingMonth,
				Limit:      core.Int64Ptr(int64(30)),
			}

			getResourceUsageReportOptions.Offset = nil
			getResourceUsageReportOptions.Limit = core.Int64Ptr(1)

			var allResults []partnermanagementv1.PartnerUsageReport
			for {
				partnerUsageReportSummary, response, err := partnerManagementService.GetResourceUsageReport(getResourceUsageReportOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(partnerUsageReportSummary).ToNot(BeNil())
				allResults = append(allResults, partnerUsageReportSummary.Reports...)

				getResourceUsageReportOptions.Offset, err = partnerUsageReportSummary.GetNextOffset()
				Expect(err).To(BeNil())

				if getResourceUsageReportOptions.Offset == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`GetResourceUsageReport(getResourceUsageReportOptions *GetResourceUsageReportOptions) using GetResourceUsageReportPager`, func() {
			getResourceUsageReportOptions := &partnermanagementv1.GetResourceUsageReportOptions{
				PartnerID:  &partnerId,
				CustomerID: &customerId,
				Month:      &billingMonth,
				Limit:      core.Int64Ptr(int64(30)),
			}

			// Test GetNext().
			pager, err := partnerManagementService.NewGetResourceUsageReportPager(getResourceUsageReportOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []partnermanagementv1.PartnerUsageReport
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = partnerManagementService.NewGetResourceUsageReportPager(getResourceUsageReportOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "GetResourceUsageReport() returned a total of %d item(s) using GetResourceUsageReportPager.\n", len(allResults))
		})
	})

	Describe(`GetResourceUsageReport -  Recursively GET usage reports for all end customers of a partner`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetResourceUsageReport(getResourceUsageReportOptions *GetResourceUsageReportOptions) with pagination`, func() {
			getResourceUsageReportOptions := &partnermanagementv1.GetResourceUsageReportOptions{
				PartnerID: &partnerId,
				Month:     &billingMonth,
				Recurse:   core.BoolPtr(true),
				Limit:     core.Int64Ptr(int64(30)),
			}

			getResourceUsageReportOptions.Offset = nil
			getResourceUsageReportOptions.Limit = core.Int64Ptr(1)

			var allResults []partnermanagementv1.PartnerUsageReport
			for {
				partnerUsageReportSummary, response, err := partnerManagementService.GetResourceUsageReport(getResourceUsageReportOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(partnerUsageReportSummary).ToNot(BeNil())
				allResults = append(allResults, partnerUsageReportSummary.Reports...)

				getResourceUsageReportOptions.Offset, err = partnerUsageReportSummary.GetNextOffset()
				Expect(err).To(BeNil())

				if getResourceUsageReportOptions.Offset == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`GetResourceUsageReport(getResourceUsageReportOptions *GetResourceUsageReportOptions) using GetResourceUsageReportPager`, func() {
			getResourceUsageReportOptions := &partnermanagementv1.GetResourceUsageReportOptions{
				PartnerID: &partnerId,
				Month:     &billingMonth,
				Recurse:   core.BoolPtr(true),
				Limit:     core.Int64Ptr(int64(30)),
			}

			// Test GetNext().
			pager, err := partnerManagementService.NewGetResourceUsageReportPager(getResourceUsageReportOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []partnermanagementv1.PartnerUsageReport
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = partnerManagementService.NewGetResourceUsageReportPager(getResourceUsageReportOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "GetResourceUsageReport() returned a total of %d item(s) using GetResourceUsageReportPager.\n", len(allResults))
		})
	})

	Describe(`GetResourceUsageReport -  Get rolled up usage reports for partner by specified viewpoint`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetResourceUsageReport(getResourceUsageReportOptions *GetResourceUsageReportOptions) with pagination`, func() {
			getResourceUsageReportOptions := &partnermanagementv1.GetResourceUsageReportOptions{
				PartnerID: &partnerId,
				Month:     &billingMonth,
				Children:  core.BoolPtr(true),
				Viewpoint: &viewpoint,
				Limit:     core.Int64Ptr(int64(30)),
			}

			getResourceUsageReportOptions.Offset = nil
			getResourceUsageReportOptions.Limit = core.Int64Ptr(1)

			var allResults []partnermanagementv1.PartnerUsageReport
			for {
				partnerUsageReportSummary, response, err := partnerManagementService.GetResourceUsageReport(getResourceUsageReportOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(partnerUsageReportSummary).ToNot(BeNil())
				allResults = append(allResults, partnerUsageReportSummary.Reports...)

				getResourceUsageReportOptions.Offset, err = partnerUsageReportSummary.GetNextOffset()
				Expect(err).To(BeNil())

				if getResourceUsageReportOptions.Offset == nil {
					break
				}
			}
			fmt.Fprintf(GinkgoWriter, "Retrieved a total of %d item(s) with pagination.\n", len(allResults))
		})
		It(`GetResourceUsageReport(getResourceUsageReportOptions *GetResourceUsageReportOptions) using GetResourceUsageReportPager`, func() {
			getResourceUsageReportOptions := &partnermanagementv1.GetResourceUsageReportOptions{
				PartnerID: &partnerId,
				Month:     &billingMonth,
				Children:  core.BoolPtr(true),
				Viewpoint: &viewpoint,
				Limit:     core.Int64Ptr(int64(30)),
			}

			// Test GetNext().
			pager, err := partnerManagementService.NewGetResourceUsageReportPager(getResourceUsageReportOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []partnermanagementv1.PartnerUsageReport
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = partnerManagementService.NewGetResourceUsageReportPager(getResourceUsageReportOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "GetResourceUsageReport() returned a total of %d item(s) using GetResourceUsageReportPager.\n", len(allResults))
		})
	})

	Describe(`GetBillingOptions - Get customers billing options of a partner`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetBillingOptions(getBillingOptionsOptions *GetBillingOptionsOptions)`, func() {
			getBillingOptionsOptions := &partnermanagementv1.GetBillingOptionsOptions{
				PartnerID: &partnerId,
				Date:      &billingMonth,
				Limit:     core.Int64Ptr(int64(30)),
			}

			billingOptionsSummary, response, err := partnerManagementService.GetBillingOptions(getBillingOptionsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(billingOptionsSummary).ToNot(BeNil())
		})
	})

	Describe(`GetBillingOptions - Get customers billing options of a Reseller for a specific partner`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetBillingOptions(getBillingOptionsOptions *GetBillingOptionsOptions)`, func() {
			getBillingOptionsOptions := &partnermanagementv1.GetBillingOptionsOptions{
				PartnerID:  &partnerId,
				ResellerID: &resellerId,
				Date:       &billingMonth,
				Limit:      core.Int64Ptr(int64(30)),
			}

			billingOptionsSummary, response, err := partnerManagementService.GetBillingOptions(getBillingOptionsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(billingOptionsSummary).ToNot(BeNil())
		})
	})

	Describe(`GetBillingOptions - Get customers billing options of an end customer for a specific partner`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetBillingOptions(getBillingOptionsOptions *GetBillingOptionsOptions)`, func() {
			getBillingOptionsOptions := &partnermanagementv1.GetBillingOptionsOptions{
				PartnerID:  &partnerId,
				CustomerID: &customerId,
				Date:       &billingMonth,
				Limit:      core.Int64Ptr(int64(30)),
			}

			billingOptionsSummary, response, err := partnerManagementService.GetBillingOptions(getBillingOptionsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(billingOptionsSummary).ToNot(BeNil())
		})
	})

	Describe(`GetCreditPoolsReport - Get subscription burn-down report of a partner`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCreditPoolsReport(getCreditPoolsReportOptions *GetCreditPoolsReportOptions)`, func() {
			getCreditPoolsReportOptions := &partnermanagementv1.GetCreditPoolsReportOptions{
				PartnerID: &partnerId,
				Date:      &billingMonth,
				Limit:     core.Int64Ptr(int64(30)),
			}

			creditPoolsReportSummary, response, err := partnerManagementService.GetCreditPoolsReport(getCreditPoolsReportOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(creditPoolsReportSummary).ToNot(BeNil())
		})
	})

	Describe(`GetCreditPoolsReport - Get subscription burn-down report of a Reseller for a specific partner`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCreditPoolsReport(getCreditPoolsReportOptions *GetCreditPoolsReportOptions)`, func() {
			getCreditPoolsReportOptions := &partnermanagementv1.GetCreditPoolsReportOptions{
				PartnerID:  &partnerId,
				ResellerID: &resellerId,
				Date:       &billingMonth,
				Limit:      core.Int64Ptr(int64(30)),
			}

			creditPoolsReportSummary, response, err := partnerManagementService.GetCreditPoolsReport(getCreditPoolsReportOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(creditPoolsReportSummary).ToNot(BeNil())
		})
	})

	Describe(`GetCreditPoolsReport - Get subscription burn-down report of a partner`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetCreditPoolsReport(getCreditPoolsReportOptions *GetCreditPoolsReportOptions)`, func() {
			getCreditPoolsReportOptions := &partnermanagementv1.GetCreditPoolsReportOptions{
				PartnerID:  &partnerId,
				CustomerID: &customerId,
				Date:       &billingMonth,
				Limit:      core.Int64Ptr(int64(30)),
			}

			creditPoolsReportSummary, response, err := partnerManagementService.GetCreditPoolsReport(getCreditPoolsReportOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(creditPoolsReportSummary).ToNot(BeNil())
		})
	})
})

//
// Utility functions are declared in the unit test file
//
