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

package globalloadbalancermonitorv1_test

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.ibm.com/ibmcloud/networking-go-sdk/globalloadbalancermonitorv1"
)

const configFile = "../cis.env"

var configLoaded bool = true

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe(`GlobalLoadBalancerMonitorV1`, func() {
	if _, err := os.Stat(configFile); err != nil {
		configLoaded = false
	}

	err := godotenv.Load(configFile)
	if err != nil {
		configLoaded = false
	}

	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("CIS_SERVICES_APIKEY"),
		URL:    os.Getenv("CIS_SERVICES_AUTH_URL"),
	}
	serviceURL := os.Getenv("API_ENDPOINT")
	crn := os.Getenv("CRN")
	globalOptions := &GlobalLoadBalancerMonitorV1Options{
		ServiceName:   "cis_services",
		URL:           serviceURL,
		Authenticator: authenticator,
		Crn:           &crn,
	}
	testService, testServiceErr := NewGlobalLoadBalancerMonitorV1(globalOptions)
	if testServiceErr != nil {
		fmt.Println(testServiceErr)
	}
	Describe(`CIS_Frontend_API_Spec-GLB_Monitor.yaml`, func() {
		Context(`Global Load Balancer Monitor`, func() {
			BeforeEach(func() {
				shouldSkipTest()
				result, response, operationErr := testService.ListAllLoadBalancerMonitors(testService.NewListAllLoadBalancerMonitorsOptions())
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())

				for _, glb := range result.Result {
					if strings.Contains(*glb.Description, "GLBMonitor") {
						option := testService.NewDeleteLoadBalancerMonitorOptions(*glb.ID)
						result, response, operationErr := testService.DeleteLoadBalancerMonitor(option)
						Expect(operationErr).To(BeNil())
						Expect(response).ToNot(BeNil())
						Expect(result).ToNot(BeNil())
						Expect(*result.Success).Should(BeTrue())
					}
				}
			})
			AfterEach(func() {
				shouldSkipTest()
				result, response, operationErr := testService.ListAllLoadBalancerMonitors(testService.NewListAllLoadBalancerMonitorsOptions())
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())

				for _, glb := range result.Result {
					if strings.Contains(*glb.Description, "GLBMonitor") {
						option := testService.NewDeleteLoadBalancerMonitorOptions(*glb.ID)
						result, response, operationErr := testService.DeleteLoadBalancerMonitor(option)
						Expect(operationErr).To(BeNil())
						Expect(response).ToNot(BeNil())
						Expect(result).ToNot(BeNil())
						Expect(*result.Success).Should(BeTrue())
					}
				}
			})
			It(`Global Load Balancer Monitor tests`, func() {
				shouldSkipTest()
				options := testService.NewCreateLoadBalancerMonitorOptions()
				options.SetExpectedBody("alive")
				options.SetExpectedCodes("2xx")
				options.SetType("http")
				options.SetDescription("Test GLBMonitor 1")
				options.SetMethod("GET")
				options.SetPort(80)
				options.SetPath("/auto/test")
				options.SetTimeout(3)
				options.SetRetries(0)
				options.SetInterval(90)
				options.SetFollowRedirects(true)
				options.SetAllowInsecure(true)

				// create global load balancer monitor with type HTTP
				result, response, err := testService.CreateLoadBalancerMonitor(options)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())
				monitorID1 := *result.Result.ID

				options = testService.NewCreateLoadBalancerMonitorOptions()
				options.SetType("tcp")
				options.SetDescription("Test GLBMonitor 2")
				options.SetMethod("GET")
				options.SetPort(20)
				options.SetPath("/auto/test")
				options.SetTimeout(3)
				options.SetRetries(0)
				options.SetInterval(90)
				// create global load balancer monitor with type TCP
				result, response, err = testService.CreateLoadBalancerMonitor(options)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())
				monitorID2 := *result.Result.ID

				// get global load balancer monitor with id
				iDs := []string{monitorID1, monitorID2}
				for _, id := range iDs {
					options := testService.NewGetLoadBalancerMonitorOptions(id)
					result, response, err := testService.GetLoadBalancerMonitor(options)
					Expect(err).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
					Expect(*result.Success).Should(BeTrue())
					Expect(*result.Result.ID).To(Equal(id))
				}

				// update global load balancer monitor with id
				editOptions := testService.NewEditLoadBalancerMonitorOptions(monitorID1)
				editOptions.SetExpectedBody("new page")
				editOptions.SetExpectedCodes("2xx")
				editOptions.SetType("https")
				editOptions.SetDescription("Test GLBMonitor 3")
				editOptions.SetMethod("GET")
				editOptions.SetPort(80)
				editOptions.SetPath("/auto/test")
				editOptions.SetTimeout(3)
				editOptions.SetRetries(0)
				editOptions.SetInterval(90)
				editOptions.SetFollowRedirects(true)
				editOptions.SetAllowInsecure(true)

				// create global load balancer monitor with type HTTP
				result, response, err = testService.EditLoadBalancerMonitor(editOptions)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())

				editOptions = testService.NewEditLoadBalancerMonitorOptions(monitorID2)
				editOptions.SetType("tcp")
				editOptions.SetDescription("Test GLBMonitor 4")
				editOptions.SetMethod("GET")
				editOptions.SetPort(12345)
				editOptions.SetPath("/auto/test")
				editOptions.SetTimeout(3)
				editOptions.SetRetries(0)
				editOptions.SetInterval(90)
				// create global load balancer monitor with type TCP
				result, response, err = testService.EditLoadBalancerMonitor(editOptions)
				Expect(err).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())

				// delete global load balancer monitor with id
				for _, id := range iDs {
					options := testService.NewDeleteLoadBalancerMonitorOptions(id)
					result, response, err := testService.DeleteLoadBalancerMonitor(options)
					Expect(err).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
					Expect(*result.Success).Should(BeTrue())
					Expect(*result.Result.ID).To(Equal(id))
				}
			})
			It(`List all global load balancer monitor instances`, func() {
				shouldSkipTest()
				var port int64 = 12345
				for i := 1; i < 10; i++ {
					options := testService.NewCreateLoadBalancerMonitorOptions()
					options.SetType("tcp")
					options.SetDescription("Test GLBMonitor " + strconv.Itoa(i))
					options.SetMethod("GET")
					options.SetPort(port)
					options.SetPath("/auto/test")
					options.SetTimeout(3)
					options.SetRetries(0)
					options.SetInterval(90)
					// create global load balancer monitor with type TCP
					result, response, err := testService.CreateLoadBalancerMonitor(options)
					Expect(err).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
					Expect(*result.Success).Should(BeTrue())
					port++
				}
				option := testService.NewListAllLoadBalancerMonitorsOptions()
				result, response, operationErr := testService.ListAllLoadBalancerMonitors(option)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())
			})
		})
	})
})
