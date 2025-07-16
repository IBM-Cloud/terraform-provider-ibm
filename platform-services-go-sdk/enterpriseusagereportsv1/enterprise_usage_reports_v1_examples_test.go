//go:build examples

/**
 * (C) Copyright IBM Corp. 2020, 2022.
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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/enterpriseusagereportsv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// This file provides an example of how to use the Enterprise Usage Reports service.
//
// The following configuration properties are assumed to be defined:
// ENTERPRISE_USAGE_REPORTS_URL=<service url>
// ENTERPRISE_USAGE_REPORTS_AUTHTYPE=iam
// ENTERPRISE_USAGE_REPORTS_APIKEY=<IAM api key of user with authority to create rules>
// ENTERPRISE_USAGE_REPORTS_AUTH_URL=<IAM token service URL - omit this if using the production environment>
// ENTERPRISE_USAGE_REPORTS_ENTERPRISE_ID=<the id of the enterprise whose usage info will be retrieved>
// ENTERPRISE_USAGE_REPORTS_ACCOUNT_ID=<the id of the account whose usage info will be retrieved>
// ENTERPRISE_USAGE_REPORTS_ACCOUNT_GROUP_ID=<the id of the account group whose usage info will be retrieved>
// ENTERPRISE_USAGE_REPORTS_BILLING_MONTH=<the billing month (yyyy-mm) for which usage info will be retrieved>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
const externalConfigFile = "../enterprise_usage_reports.env"

var (
	enterpriseUsageReportsService *enterpriseusagereportsv1.EnterpriseUsageReportsV1
	config                        map[string]string
	configLoaded                  bool = false

	accountID      string
	accountGroupID string
	enterpriseID   string
	billingMonth   string
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`EnterpriseUsageReportsV1 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(enterpriseusagereportsv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			accountID = config["ACCOUNT_ID"]
			Expect(accountID).ToNot(BeEmpty())

			accountGroupID = config["ACCOUNT_GROUP_ID"]
			Expect(accountGroupID).ToNot(BeEmpty())

			enterpriseID = config["ENTERPRISE_ID"]
			Expect(enterpriseID).ToNot(BeEmpty())

			billingMonth = config["BILLING_MONTH"]
			Expect(billingMonth).ToNot(BeEmpty())
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

			enterpriseUsageReportsServiceOptions := &enterpriseusagereportsv1.EnterpriseUsageReportsV1Options{}

			enterpriseUsageReportsService, err = enterpriseusagereportsv1.NewEnterpriseUsageReportsV1UsingExternalConfig(enterpriseUsageReportsServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(enterpriseUsageReportsService).ToNot(BeNil())
		})
	})

	Describe(`EnterpriseUsageReportsV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetResourceUsageReport request example`, func() {
			fmt.Println("\nGetResourceUsageReport() result:")
			// begin-get_resource_usage_report
			getResourceUsageReportOptions := &enterpriseusagereportsv1.GetResourceUsageReportOptions{
				EnterpriseID: &enterpriseID,
				Month:        &billingMonth,
			}

			pager, err := enterpriseUsageReportsService.NewGetResourceUsageReportPager(getResourceUsageReportOptions)
			if err != nil {
				panic(err)
			}

			var allResults []enterpriseusagereportsv1.ResourceUsageReport
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))
			// end-get_resource_usage_report
		})
	})
})
