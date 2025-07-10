//go:build examples

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

package usagemeteringv4_test

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/usagemeteringv4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the Usage Metering service.
//
// The following configuration properties are assumed to be defined:
//
// USAGE_METERING_URL=<service url>
// USAGE_METERING_AUTHTYPE=iam
// USAGE_METERING_APIKEY=<your iam apikey>
// USAGE_METERING_AUTH_URL=<IAM token service URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//

var _ = Describe(`UsageMeteringV4 Examples Tests`, func() {
	const externalConfigFile = "../usage_metering.env"

	var (
		usageMeteringService *usagemeteringv4.UsageMeteringV4
		config               map[string]string
		configLoaded         bool = false
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
			config, err = core.GetServiceProperties(usagemeteringv4.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

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

			usageMeteringServiceOptions := &usagemeteringv4.UsageMeteringV4Options{}

			usageMeteringService, err = usagemeteringv4.NewUsageMeteringV4UsingExternalConfig(usageMeteringServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(usageMeteringService).ToNot(BeNil())
		})
	})

	Describe(`UsageMeteringV4 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReportResourceUsage request example`, func() {
			// Use the same start and end time since we're reporting events.
			endTime := time.Now().Unix() * 1000
			startTime := endTime

			// Use these values in the resource usage being reported below.
			resourceID := "cloudant"
			resourceInstanceID := "crn:v1:staging:public:cloudantnosqldb:us-south:a/f5086e3df886495991303628d21da513:3aafbbee-88e2-4d29-b144-9d267d97064c::"
			planID := "cloudant-standard"
			region := "us-south"

			fmt.Println("\nReportResourceUsage() result:")
			// begin-report_resource_usage

			// Report usage for a mythical resource.
			// Use zero for quantities since this is only an example.
			resourceInstanceUsageModel := usagemeteringv4.ResourceInstanceUsage{
				ResourceInstanceID: &resourceInstanceID,
				PlanID:             &planID,
				Region:             &region,
				Start:              &startTime,
				End:                &endTime,
				MeasuredUsage: []usagemeteringv4.MeasureAndQuantity{
					usagemeteringv4.MeasureAndQuantity{
						Measure:  core.StringPtr("LOOKUP"),
						Quantity: core.Int64Ptr(0),
					},
					usagemeteringv4.MeasureAndQuantity{
						Measure:  core.StringPtr("WRITE"),
						Quantity: core.Int64Ptr(0),
					},
					usagemeteringv4.MeasureAndQuantity{
						Measure:  core.StringPtr("QUERY"),
						Quantity: core.Int64Ptr(0),
					},
					usagemeteringv4.MeasureAndQuantity{
						Measure:  core.StringPtr("GIGABYTE"),
						Quantity: core.Int64Ptr(0),
					},
				},
			}
			reportResourceUsageOptions := usageMeteringService.NewReportResourceUsageOptions(
				resourceID,
				[]usagemeteringv4.ResourceInstanceUsage{resourceInstanceUsageModel},
			)

			responseAccepted, response, err := usageMeteringService.ReportResourceUsage(reportResourceUsageOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(responseAccepted, "", "  ")
			fmt.Println(string(b))

			// end-report_resource_usage

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(responseAccepted).ToNot(BeNil())
		})
	})
})
