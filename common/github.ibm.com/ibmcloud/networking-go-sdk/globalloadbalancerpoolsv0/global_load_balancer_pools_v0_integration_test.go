/*
 * (C) Copyright IBM Corp. 2020.
 */

package globalloadbalancerpoolsv0_test

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/networking-go-sdk/globalloadbalancermonitorv1"
	. "github.ibm.com/ibmcloud/networking-go-sdk/globalloadbalancerpoolsv0"
)

const configFile = "../cis.env"

var configLoaded bool = true

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe(`GlobalLoadBalancerPoolsV0`, func() {
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
	globalOptions := &GlobalLoadBalancerPoolsV0Options{
		ServiceName:   "cis_services",
		URL:           serviceURL,
		Authenticator: authenticator,
		Crn:           &crn,
	}
	testService, testServiceErr := NewGlobalLoadBalancerPoolsV0(globalOptions)
	if testServiceErr != nil {
		fmt.Println(testServiceErr)
	}

	monitorOptions := &globalloadbalancermonitorv1.GlobalLoadBalancerMonitorV1Options{
		ServiceName:   "cis_services",
		URL:           serviceURL,
		Authenticator: authenticator,
		Crn:           &crn,
	}

	monitorService, testServiceErr := globalloadbalancermonitorv1.NewGlobalLoadBalancerMonitorV1(monitorOptions)
	if testServiceErr != nil {
		fmt.Println(testServiceErr)
	}

	var monitorID string

	Describe(`create/update/get/delete global load balancer pool`, func() {
		Context(`GlobalLoadBalancerPoolsV0`, func() {
			BeforeEach(func() {
				shouldSkipTest()

				/// delete all glb pools
				result, response, operationErr := testService.ListAllLoadBalancerPools(testService.NewListAllLoadBalancerPoolsOptions())
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				for _, pool := range result.Result {
					if strings.Contains(*pool.Name, "glbpooltest") {
						option := testService.NewDeleteLoadBalancerPoolOptions(*pool.ID)
						result, response, operationErr := testService.DeleteLoadBalancerPool(option)
						Expect(operationErr).To(BeNil())
						Expect(response).ToNot(BeNil())
						Expect(result).ToNot(BeNil())
					}
				}

				// deelte all glb monitors
				monitorResult, monitorResponse, monitorErr := monitorService.ListAllLoadBalancerMonitors(monitorService.NewListAllLoadBalancerMonitorsOptions())
				Expect(monitorErr).To(BeNil())
				Expect(monitorResponse).ToNot(BeNil())
				Expect(monitorResult).ToNot(BeNil())
				Expect(*monitorResult.Success).Should(BeTrue())

				for _, glb := range monitorResult.Result {
					if strings.Contains(*glb.Description, "glbpooltest") {
						option := monitorService.NewDeleteLoadBalancerMonitorOptions(*glb.ID)
						result, response, operationErr := monitorService.DeleteLoadBalancerMonitor(option)
						Expect(operationErr).To(BeNil())
						Expect(response).ToNot(BeNil())
						Expect(result).ToNot(BeNil())
						Expect(*result.Success).Should(BeTrue())
					}
				}

				// create glb monitor for pool
				options := monitorService.NewCreateLoadBalancerMonitorOptions()
				options.SetExpectedBody("alive")
				options.SetExpectedCodes("2xx")
				options.SetType("http")
				options.SetDescription("Test glbpooltest 1")
				options.SetMethod("GET")
				options.SetPort(80)
				options.SetPath("/auto/test")
				options.SetTimeout(3)
				options.SetRetries(0)
				options.SetInterval(90)
				options.SetFollowRedirects(true)
				options.SetAllowInsecure(true)

				// create global load balancer monitor with type HTTP
				res, resp, err := monitorService.CreateLoadBalancerMonitor(options)
				Expect(err).To(BeNil())
				Expect(resp).ToNot(BeNil())
				Expect(res).ToNot(BeNil())
				Expect(*res.Success).Should(BeTrue())
				monitorID = *res.Result.ID
			})
			AfterEach(func() {
				/// delete all glb pools
				result, response, operationErr := testService.ListAllLoadBalancerPools(testService.NewListAllLoadBalancerPoolsOptions())
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				for _, pool := range result.Result {
					if strings.Contains(*pool.Name, "glbpooltest") {
						option := testService.NewDeleteLoadBalancerPoolOptions(*pool.ID)
						result, response, operationErr := testService.DeleteLoadBalancerPool(option)
						Expect(operationErr).To(BeNil())
						Expect(response).ToNot(BeNil())
						Expect(result).ToNot(BeNil())
					}
				}

				// deelte all glb monitors
				monitorResult, monitorResponse, monitorErr := monitorService.ListAllLoadBalancerMonitors(monitorService.NewListAllLoadBalancerMonitorsOptions())
				Expect(monitorErr).To(BeNil())
				Expect(monitorResponse).ToNot(BeNil())
				Expect(monitorResult).ToNot(BeNil())
				Expect(*monitorResult.Success).Should(BeTrue())

				for _, glb := range monitorResult.Result {
					if strings.Contains(*glb.Description, "glbpooltest") {
						option := monitorService.NewDeleteLoadBalancerMonitorOptions(*glb.ID)
						result, response, operationErr := monitorService.DeleteLoadBalancerMonitor(option)
						Expect(operationErr).To(BeNil())
						Expect(response).ToNot(BeNil())
						Expect(result).ToNot(BeNil())
						Expect(*result.Success).Should(BeTrue())
					}
				}
			})
			It(`GlobalLoadBalancerPool Test`, func() {
				// create glb pool
				option := testService.NewCreateLoadBalancerPoolOptions()
				option.SetName("glbpooltest-pool1")
				regions := []string{"WEU", "ENAM"}
				option.SetCheckRegions(regions)
				origin := &LoadBalancerPoolReqOriginsItem{
					Name:    core.StringPtr("app-server-1"),
					Address: core.StringPtr("www.test.com"),
					Enabled: core.BoolPtr(true),
				}
				origins := []LoadBalancerPoolReqOriginsItem{*origin}
				option.SetOrigins(origins)
				option.SetDescription("Test GLB Pool 1")
				option.SetMinimumOrigins(1)
				option.SetEnabled(true)
				option.SetNotificationEmail("notify@in.ibm.com")
				option.SetMonitor(monitorID)

				createResult, createResponse, createErr := testService.CreateLoadBalancerPool(option)
				Expect(createErr).To(BeNil())
				Expect(createResponse).ToNot(BeNil())
				Expect(createResult).ToNot(BeNil())
				Expect(*createResult.Success).Should(BeTrue())
				poolID := *createResult.Result.ID

				// update glb pool
				opt := testService.NewEditLoadBalancerPoolOptions(poolID)
				opt.SetName("glbpooltest-pool2")
				regions = []string{"WEU", "WNAM"}
				opt.SetCheckRegions(regions)
				origin = &LoadBalancerPoolReqOriginsItem{
					Name:    core.StringPtr("app-server-2"),
					Address: core.StringPtr("www.test1.com"),
					Enabled: core.BoolPtr(true),
				}
				origins = []LoadBalancerPoolReqOriginsItem{*origin}
				opt.SetOrigins(origins)
				opt.SetDescription("Test GLB Pool 2")
				opt.SetMinimumOrigins(1)
				opt.SetEnabled(true)
				opt.SetNotificationEmail("pool@in.ibm.com")
				opt.SetMonitor(monitorID)

				editResult, editResponse, editErr := testService.EditLoadBalancerPool(opt)
				Expect(editErr).To(BeNil())
				Expect(editResponse).ToNot(BeNil())
				Expect(editResult).ToNot(BeNil())
				Expect(*editResult.Success).Should(BeTrue())

				// get glb pool by id
				getOpt := testService.NewGetLoadBalancerPoolOptions(poolID)
				getResult, getResp, getErr := testService.GetLoadBalancerPool(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				// delete glb pool by id
				delOpt := testService.NewDeleteLoadBalancerPoolOptions(poolID)
				delResult, delResp, delErr := testService.DeleteLoadBalancerPool(delOpt)
				Expect(delErr).To(BeNil())
				Expect(delResp).ToNot(BeNil())
				Expect(delResult).ToNot(BeNil())
				Expect(*delResult.Success).Should(BeTrue())
			})
			It(`List all GLB Pools`, func() {
				// create glb pool
				for i := 1; i < 5; i++ {
					option := testService.NewCreateLoadBalancerPoolOptions()
					option.SetName("glbpooltest-pool" + strconv.Itoa(i))
					regions := []string{"WEU", "ENAM"}
					option.SetCheckRegions(regions)
					print("www.test" + strconv.Itoa(i) + ".com")
					origin := &LoadBalancerPoolReqOriginsItem{
						Name:    core.StringPtr("app-server-" + strconv.Itoa(i)),
						Address: core.StringPtr("www.test" + strconv.Itoa(i) + ".com"),
						Enabled: core.BoolPtr(true),
					}
					origins := []LoadBalancerPoolReqOriginsItem{*origin}
					option.SetOrigins(origins)
					option.SetDescription("Test GLB Pool " + strconv.Itoa(i))
					option.SetMinimumOrigins(1)
					option.SetEnabled(true)
					option.SetNotificationEmail("notify@in.ibm.com")
					option.SetMonitor(monitorID)

					createResult, createResponse, createErr := testService.CreateLoadBalancerPool(option)
					Expect(createErr).To(BeNil())
					Expect(createResponse).ToNot(BeNil())
					Expect(createResult).ToNot(BeNil())
					Expect(*createResult.Success).Should(BeTrue())
				}
				result, response, operationErr := testService.ListAllLoadBalancerPools(testService.NewListAllLoadBalancerPoolsOptions())
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
			})
		})
	})
})
