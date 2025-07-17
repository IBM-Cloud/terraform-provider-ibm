//go:build integration

/**
 * (C) Copyright IBM Corp. 2020.
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

package enterpriseusagereportsv1_test

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/enterpriseusagereportsv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the enterpriseusagereportsv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`EnterpriseUsageReportsV1 Integration Tests`, func() {

	const externalConfigFile = "../enterprise_usage_reports.env"

	var (
		err                           error
		enterpriseUsageReportsService *enterpriseusagereportsv1.EnterpriseUsageReportsV1
		serviceURL                    string
		config                        map[string]string

		accountID      string
		accountGroupID string
		enterpriseID   string
		billingMonth   string
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
			config, err = core.GetServiceProperties(enterpriseusagereportsv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %s\n", serviceURL)

			accountID = config["ACCOUNT_ID"]
			Expect(accountID).ToNot(BeEmpty())

			accountGroupID = config["ACCOUNT_GROUP_ID"]
			Expect(accountGroupID).ToNot(BeEmpty())

			enterpriseID = config["ENTERPRISE_ID"]
			Expect(enterpriseID).ToNot(BeEmpty())

			billingMonth = config["BILLING_MONTH"]
			Expect(billingMonth).ToNot(BeEmpty())

			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			enterpriseUsageReportsServiceOptions := &enterpriseusagereportsv1.EnterpriseUsageReportsV1Options{}

			enterpriseUsageReportsService, err = enterpriseusagereportsv1.NewEnterpriseUsageReportsV1UsingExternalConfig(enterpriseUsageReportsServiceOptions)

			Expect(err).To(BeNil())
			Expect(enterpriseUsageReportsService).ToNot(BeNil())
			Expect(enterpriseUsageReportsService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelError, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			enterpriseUsageReportsService.EnableRetries(3, 30*time.Second)
		})
	})

	Describe(`GetResourceUsageReport `, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`Using Enterprise ID`, func() {

			// Retrieve the search results one page at a time to test pagination.
			getResourceUsageReportOptions := &enterpriseusagereportsv1.GetResourceUsageReportOptions{
				EnterpriseID: &enterpriseID,
				Month:        &billingMonth,
				Limit:        core.Int64Ptr(1),
			}

			var results []enterpriseusagereportsv1.ResourceUsageReport = make([]enterpriseusagereportsv1.ResourceUsageReport, 0)
			var offset *string = nil
			var moreResults bool = true

			for moreResults {
				getResourceUsageReportOptions.Offset = offset

				reports, response, err := enterpriseUsageReportsService.GetResourceUsageReport(getResourceUsageReportOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(reports).ToNot(BeNil())

				// Add the just-retrieved page to the results.
				if len(reports.Reports) > 0 {
					results = append(results, reports.Reports...)
				}

				// Get the offset of the next page.
				if reports.Next != nil {
					offset = getOffsetFromURL(reports.Next.Href)
				} else {
					offset = nil
				}

				moreResults = (offset != nil)
			}

			// Make sure we got back a non-empty set of results.
			Expect(results).ToNot(BeEmpty())
		})
		It(`Using Account ID`, func() {

			// Retrieve the search results one page at a time to test pagination.
			getResourceUsageReportOptions := &enterpriseusagereportsv1.GetResourceUsageReportOptions{
				AccountID: &accountID,
				Month:     &billingMonth,
				Limit:     core.Int64Ptr(1),
			}

			var results []enterpriseusagereportsv1.ResourceUsageReport = make([]enterpriseusagereportsv1.ResourceUsageReport, 0)
			var offset *string = nil
			var moreResults bool = true

			for moreResults {
				getResourceUsageReportOptions.Offset = offset

				reports, response, err := enterpriseUsageReportsService.GetResourceUsageReport(getResourceUsageReportOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(reports).ToNot(BeNil())

				// Add the just-retrieved page to the results.
				if len(reports.Reports) > 0 {
					results = append(results, reports.Reports...)
				}

				// Get the offset of the next page.
				if reports.Next != nil {
					offset = getOffsetFromURL(reports.Next.Href)
				} else {
					offset = nil
				}

				moreResults = (offset != nil)
			}

			// Make sure we got back a non-empty set of results.
			Expect(results).ToNot(BeEmpty())
		})
		It(`Using Account Group ID`, func() {

			// Retrieve the search results one page at a time to test pagination.
			getResourceUsageReportOptions := &enterpriseusagereportsv1.GetResourceUsageReportOptions{
				AccountGroupID: &accountGroupID,
				Month:          &billingMonth,
				Limit:          core.Int64Ptr(1),
			}

			var results []enterpriseusagereportsv1.ResourceUsageReport = make([]enterpriseusagereportsv1.ResourceUsageReport, 0)
			var offset *string = nil
			var moreResults bool = true

			for moreResults {
				getResourceUsageReportOptions.Offset = offset

				reports, response, err := enterpriseUsageReportsService.GetResourceUsageReport(getResourceUsageReportOptions)
				Expect(err).To(BeNil())
				Expect(response.StatusCode).To(Equal(200))
				Expect(reports).ToNot(BeNil())

				// Add the just-retrieved page to the results.
				if len(reports.Reports) > 0 {
					results = append(results, reports.Reports...)
				}

				// Get the offset of the next page.
				if reports.Next != nil {
					offset = getOffsetFromURL(reports.Next.Href)
				} else {
					offset = nil
				}

				moreResults = (offset != nil)
			}

			// Make sure we got back a non-empty set of results.
			Expect(results).ToNot(BeEmpty())
		})
		It(`Using Account ID with pager`, func() {
			getResourceUsageReportOptions := &enterpriseusagereportsv1.GetResourceUsageReportOptions{
				AccountGroupID: &accountGroupID,
				Month:          &billingMonth,
			}

			// Test GetNext().
			pager, err := enterpriseUsageReportsService.NewGetResourceUsageReportPager(getResourceUsageReportOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			var allResults []enterpriseusagereportsv1.ResourceUsageReport
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				Expect(err).To(BeNil())
				Expect(nextPage).ToNot(BeNil())
				allResults = append(allResults, nextPage...)
			}

			// Test GetAll().
			pager, err = enterpriseUsageReportsService.NewGetResourceUsageReportPager(getResourceUsageReportOptions)
			Expect(err).To(BeNil())
			Expect(pager).ToNot(BeNil())

			allItems, err := pager.GetAll()
			Expect(err).To(BeNil())
			Expect(allItems).ToNot(BeNil())

			Expect(len(allItems)).To(Equal(len(allResults)))
			fmt.Fprintf(GinkgoWriter, "GetResourceUsageReport() returned a total of %d item(s) using GetResourceUsageReportPager.\n", len(allResults))
		})
	})
})

func getOffsetFromURL(sptr *string) *string {
	if sptr == nil {
		return nil
	}

	s := *sptr
	if s == "" {
		return nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return nil
	}

	if u.RawQuery == "" {
		return nil
	}

	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil
	}

	token := q.Get("offset")
	if token == "" {
		return nil
	}
	return &token
}
