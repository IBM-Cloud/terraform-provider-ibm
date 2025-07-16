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

package usagemeteringv4_test

import (
	"fmt"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/platform-services-go-sdk/common"
	"github.com/IBM/platform-services-go-sdk/usagemeteringv4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the usagemeteringv4 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`UsageMeteringV4 Integration Tests`, func() {

	const externalConfigFile = "../usage_metering.env"

	var (
		err                  error
		usageMeteringService *usagemeteringv4.UsageMeteringV4
		serviceURL           string
		config               map[string]string

		resourceID = "cloudant"
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
			config, err = core.GetServiceProperties(usagemeteringv4.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Printf("Service URL: %s\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			usageMeteringServiceOptions := &usagemeteringv4.UsageMeteringV4Options{}

			usageMeteringService, err = usagemeteringv4.NewUsageMeteringV4UsingExternalConfig(usageMeteringServiceOptions)

			Expect(err).To(BeNil())
			Expect(usageMeteringService).ToNot(BeNil())
			Expect(usageMeteringService.Service.Options.URL).To(Equal(serviceURL))
		})
	})

	Describe(`ReportResourceUsage - Report Resource Controller resource usage`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReportResourceUsage(reportResourceUsageOptions *ReportResourceUsageOptions)`, func() {
			// We'll use the same start and end time since we're reporting events.
			endTime := time.Now().Unix() * 1000
			startTime := endTime

			resourceInstanceID := "crn:v1:staging:public:cloudantnosqldb:us-south:a/f5086e3df886495991303628d21da513:3aafbbee-88e2-4d29-b144-9d267d97064c::"
			planID := "cloudant-standard"
			region := "us-south"
			resourceUsage := usagemeteringv4.ResourceInstanceUsage{
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

			reportResourceUsageOptions := &usagemeteringv4.ReportResourceUsageOptions{
				ResourceID:    &resourceID,
				ResourceUsage: []usagemeteringv4.ResourceInstanceUsage{resourceUsage},
			}

			responseAccepted, response, err := usageMeteringService.ReportResourceUsage(reportResourceUsageOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(responseAccepted).ToNot(BeNil())

			fmt.Fprintf(GinkgoWriter, "ReportResourceUsage() result:\n%s\n", common.ToJSON(responseAccepted))
		})
	})
})
