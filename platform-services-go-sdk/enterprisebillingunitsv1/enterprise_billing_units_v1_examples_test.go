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

package enterprisebillingunitsv1_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/enterprisebillingunitsv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the Enterprise Billing Units service.
//
// The following configuration properties are assumed to be defined:
//
// ENTERPRISE_BILLING_UNITS_URL=<service url>
// ENTERPRISE_BILLING_UNITS_AUTHTYPE=iam
// ENTERPRISE_BILLING_UNITS_APIKEY=<your iam apikey>
// ENTERPRISE_BILLING_UNITS_AUTH_URL=<IAM token service URL - omit this if using the production environment>
// ENTERPRISE_BILLING_UNITS_ENTERPRISE_ID=<id of enterprise to use for examples>
// ENTERPRISE_BILLING_UNITS_BILLING_UNIT_ID=<id of billing unit to use for examples>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//

var _ = Describe(`EnterpriseBillingUnitsV1 Examples Tests`, func() {
	const externalConfigFile = "../enterprise_billing_units.env"

	var (
		enterpriseBillingUnitsService *enterprisebillingunitsv1.EnterpriseBillingUnitsV1
		config                        map[string]string
		configLoaded                  bool = false

		enterpriseID  string
		billingUnitID string
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
			config, err = core.GetServiceProperties(enterprisebillingunitsv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			enterpriseID = config["ENTERPRISE_ID"]
			Expect(enterpriseID).ToNot(BeEmpty())

			billingUnitID = config["BILLING_UNIT_ID"]
			Expect(billingUnitID).ToNot(BeEmpty())

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

			enterpriseBillingUnitsServiceOptions := &enterprisebillingunitsv1.EnterpriseBillingUnitsV1Options{}

			enterpriseBillingUnitsService, err = enterprisebillingunitsv1.NewEnterpriseBillingUnitsV1UsingExternalConfig(enterpriseBillingUnitsServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(enterpriseBillingUnitsService).ToNot(BeNil())
		})
	})

	Describe(`EnterpriseBillingUnitsV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetBillingUnit request example`, func() {
			fmt.Println("\nGetBillingUnit() result:")
			// begin-get_billing_unit

			getBillingUnitOptions := enterpriseBillingUnitsService.NewGetBillingUnitOptions(
				billingUnitID,
			)

			billingUnit, response, err := enterpriseBillingUnitsService.GetBillingUnit(getBillingUnitOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(billingUnit, "", "  ")
			fmt.Println(string(b))

			// end-get_billing_unit

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(billingUnit).ToNot(BeNil())

		})
		It(`ListBillingUnits request example`, func() {
			fmt.Println("\nListBillingUnits() result:")
			// begin-list_billing_units

			listBillingUnitsOptions := enterpriseBillingUnitsService.NewListBillingUnitsOptions()
			listBillingUnitsOptions.SetEnterpriseID(enterpriseID)

			pager, err := enterpriseBillingUnitsService.NewBillingUnitsPager(listBillingUnitsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []enterprisebillingunitsv1.BillingUnit
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))

			// end-list_billing_units

			Expect(err).To(BeNil())
			Expect(allResults).ToNot(BeNil())

		})
		It(`ListBillingOptions request example`, func() {
			fmt.Println("\nListBillingOptions() result:")
			// begin-list_billing_options

			listBillingOptionsOptions := enterpriseBillingUnitsService.NewListBillingOptionsOptions(
				billingUnitID,
			)

			pager, err := enterpriseBillingUnitsService.NewBillingOptionsPager(listBillingOptionsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []enterprisebillingunitsv1.BillingOption
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			b, _ := json.MarshalIndent(allResults, "", "  ")
			fmt.Println(string(b))

			// end-list_billing_options

			Expect(err).To(BeNil())
			Expect(allResults).ToNot(BeNil())

		})
		It(`GetCreditPools request example`, func() {
			fmt.Println("\nGetCreditPools() result:")
			// begin-get_credit_pools

			getCreditPoolsOptions := enterpriseBillingUnitsService.NewGetCreditPoolsOptions(
				billingUnitID,
			)

			creditPoolsList, response, err := enterpriseBillingUnitsService.GetCreditPools(getCreditPoolsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(creditPoolsList, "", "  ")
			fmt.Println(string(b))

			// end-get_credit_pools

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(creditPoolsList).ToNot(BeNil())
		})
	})
})
