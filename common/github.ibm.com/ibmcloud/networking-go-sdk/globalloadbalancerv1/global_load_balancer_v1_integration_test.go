/*
 * (C) Copyright IBM Corp. 2020.
 */

package globalloadbalancerv1_test

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
	"github.ibm.com/ibmcloud/networking-go-sdk/globalloadbalancerpoolsv0"
	. "github.ibm.com/ibmcloud/networking-go-sdk/globalloadbalancerv1"
)

const configFile = "../cis.env"

var configLoaded bool = true

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe(`GlobalLoadBalancerV1`, func() {
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
	zone_id := os.Getenv("ZONE_ID")
	URL := os.Getenv("URL")
	globalOptions := &globalloadbalancerpoolsv0.GlobalLoadBalancerPoolsV0Options{
		ServiceName:   "cis_services",
		URL:           serviceURL,
		Authenticator: authenticator,
		Crn:           &crn,
	}
	testService, testServiceErr := globalloadbalancerpoolsv0.NewGlobalLoadBalancerPoolsV0(globalOptions)
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

	glbOptions := &GlobalLoadBalancerV1Options{
		ServiceName:    "cis_services",
		URL:            serviceURL,
		Authenticator:  authenticator,
		Crn:            &crn,
		ZoneIdentifier: &zone_id,
	}
	glbService, glbErr := NewGlobalLoadBalancerV1(glbOptions)
	if glbErr != nil {
		fmt.Println(glbErr)
	}

	var (
		monitorID  = ""
		glbPoolID1 = ""
		glbPoolID2 = ""
	)
	Describe(`GlobalLoadBalancerV1`, func() {
		Context(`GlobalLoadBalancer create/delete/update/get context`, func() {
			BeforeEach(func() {
				shouldSkipTest()
				// deelte all glb
				listResult, listResponse, listErr := glbService.ListAllLoadBalancers(glbService.NewListAllLoadBalancersOptions())
				Expect(listErr).To(BeNil())
				Expect(listResponse).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())

				for _, glb := range listResult.Result {
					if strings.Contains(*glb.Name, "glbtest") {
						option := glbService.NewDeleteLoadBalancerOptions(*glb.ID)
						result, response, operationErr := glbService.DeleteLoadBalancer(option)
						Expect(operationErr).To(BeNil())
						Expect(response).ToNot(BeNil())
						Expect(result).ToNot(BeNil())
						Expect(*result.Success).Should(BeTrue())
					}
				}

				/// delete all glb pools
				result, response, operationErr := testService.ListAllLoadBalancerPools(testService.NewListAllLoadBalancerPoolsOptions())
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				for _, pool := range result.Result {
					if strings.Contains(*pool.Name, "glbtest") {
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
					if strings.Contains(*glb.Description, "glbtest") {
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
				options.SetDescription("Test glbtest 1")
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

				// create glb pool 1
				option := testService.NewCreateLoadBalancerPoolOptions()
				option.SetName("glbtest-pool11")
				regions := []string{"WEU", "ENAM"}
				option.SetCheckRegions(regions)
				origin := &globalloadbalancerpoolsv0.LoadBalancerPoolReqOriginsItem{
					Name:    core.StringPtr("app-server-1"),
					Address: core.StringPtr("www.test.com"),
					Enabled: core.BoolPtr(true),
				}
				origins := []globalloadbalancerpoolsv0.LoadBalancerPoolReqOriginsItem{*origin}
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
				glbPoolID1 = *createResult.Result.ID

				// create glb pool 2
				option = testService.NewCreateLoadBalancerPoolOptions()
				option.SetName("glbtest-pool2")
				regions = []string{"WEU", "ENAM"}
				option.SetCheckRegions(regions)
				origin = &globalloadbalancerpoolsv0.LoadBalancerPoolReqOriginsItem{
					Name:    core.StringPtr("app-server-2"),
					Address: core.StringPtr("www.test2.com"),
					Enabled: core.BoolPtr(true),
				}
				origins = []globalloadbalancerpoolsv0.LoadBalancerPoolReqOriginsItem{*origin}
				option.SetOrigins(origins)
				option.SetDescription("Test GLB Pool 2")
				option.SetMinimumOrigins(1)
				option.SetEnabled(true)
				option.SetNotificationEmail("notify@in.ibm.com")
				option.SetMonitor(monitorID)

				createResult, createResponse, createErr = testService.CreateLoadBalancerPool(option)
				Expect(createErr).To(BeNil())
				Expect(createResponse).ToNot(BeNil())
				Expect(createResult).ToNot(BeNil())
				Expect(*createResult.Success).Should(BeTrue())
				glbPoolID2 = *createResult.Result.ID
			})
			AfterEach(func() {
				shouldSkipTest()
				// deelte all glb
				listResult, listResponse, listErr := glbService.ListAllLoadBalancers(glbService.NewListAllLoadBalancersOptions())
				Expect(listErr).To(BeNil())
				Expect(listResponse).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())

				for _, glb := range listResult.Result {
					if strings.Contains(*glb.Name, "glbtest") {
						option := glbService.NewDeleteLoadBalancerOptions(*glb.ID)
						result, response, operationErr := glbService.DeleteLoadBalancer(option)
						Expect(operationErr).To(BeNil())
						Expect(response).ToNot(BeNil())
						Expect(result).ToNot(BeNil())
						Expect(*result.Success).Should(BeTrue())
					}
				}

				/// delete all glb pools
				result, response, operationErr := testService.ListAllLoadBalancerPools(testService.NewListAllLoadBalancerPoolsOptions())
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				for _, pool := range result.Result {
					if strings.Contains(*pool.Name, "glbtest") {
						option := testService.NewDeleteLoadBalancerPoolOptions(*pool.ID)
						result, response, operationErr := testService.DeleteLoadBalancerPool(option)
						Expect(operationErr).To(BeNil())
						Expect(response).ToNot(BeNil())
						Expect(result).ToNot(BeNil())
					}
				}

				// delete all glb monitors
				monitorResult, monitorResponse, monitorErr := monitorService.ListAllLoadBalancerMonitors(monitorService.NewListAllLoadBalancerMonitorsOptions())
				Expect(monitorErr).To(BeNil())
				Expect(monitorResponse).ToNot(BeNil())
				Expect(monitorResult).ToNot(BeNil())
				Expect(*monitorResult.Success).Should(BeTrue())

				for _, glb := range monitorResult.Result {
					if strings.Contains(*glb.Description, "glbtest") {
						option := monitorService.NewDeleteLoadBalancerMonitorOptions(*glb.ID)
						result, response, operationErr := monitorService.DeleteLoadBalancerMonitor(option)
						Expect(operationErr).To(BeNil())
						Expect(response).ToNot(BeNil())
						Expect(result).ToNot(BeNil())
						Expect(*result.Success).Should(BeTrue())
					}
				}
			})
			It(`create/update/delete/get global load balancer`, func() {
				// create glb
				glbOpt := glbService.NewCreateLoadBalancerOptions()
				glbOpt.SetName("glbtest1." + URL)
				glbOpt.SetDefaultPools([]string{glbPoolID1})
				glbOpt.SetFallbackPool(glbPoolID2)
				glbOpt.SetDescription("test glb")
				glbOpt.SetEnabled(false)
				glbOpt.SetSessionAffinity(CreateLoadBalancerOptions_SessionAffinity_Cookie)
				glbOpt.SetSteeringPolicy(CreateLoadBalancerOptions_SteeringPolicy_DynamicLatency)
				createResult, createResp, createErr := glbService.CreateLoadBalancer(glbOpt)
				Expect(createErr).To(BeNil())
				Expect(createResp).ToNot(BeNil())
				Expect(createResult).ToNot(BeNil())
				Expect(*createResult.Success).Should(BeTrue())
				glbID := *createResult.Result.ID

				// update glb
				editOpt := glbService.NewEditLoadBalancerOptions(glbID)
				editOpt.SetName("glbtest2." + URL)
				editOpt.SetDefaultPools([]string{glbPoolID1})
				editOpt.SetFallbackPool(glbPoolID1)
				editOpt.SetDescription("test glb 2")
				editOpt.SetEnabled(true)
				editOpt.SetSessionAffinity(EditLoadBalancerOptions_SessionAffinity_IpCookie)
				editOpt.SetSteeringPolicy(EditLoadBalancerOptions_SteeringPolicy_Geo)
				updateResult, updateResp, updateErr := glbService.EditLoadBalancer(editOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())

				// get glb
				getOpt := glbService.NewGetLoadBalancerSettingsOptions(glbID)
				getResult, getResp, getErr := glbService.GetLoadBalancerSettings(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				// delete glb
				delOpt := glbService.NewDeleteLoadBalancerOptions(glbID)
				delResult, delResp, delErr := glbService.DeleteLoadBalancer(delOpt)
				Expect(delErr).To(BeNil())
				Expect(delResp).ToNot(BeNil())
				Expect(delResult).ToNot(BeNil())
				Expect(*delResult.Success).Should(BeTrue())
			})
			It(`List all GLB`, func() {
				for i := 1; i < 5; i++ {
					// create glb
					glbOpt := glbService.NewCreateLoadBalancerOptions()
					glbOpt.SetName("glbtest" + strconv.Itoa(i) + "." + URL)
					glbOpt.SetDefaultPools([]string{glbPoolID1})
					glbOpt.SetFallbackPool(glbPoolID2)
					glbOpt.SetDescription("test glb " + strconv.Itoa(i))
					glbOpt.SetEnabled(false)
					glbOpt.SetSessionAffinity(CreateLoadBalancerOptions_SessionAffinity_Cookie)
					glbOpt.SetSteeringPolicy(CreateLoadBalancerOptions_SteeringPolicy_DynamicLatency)
					createResult, createResp, createErr := glbService.CreateLoadBalancer(glbOpt)
					Expect(createErr).To(BeNil())
					Expect(createResp).ToNot(BeNil())
					Expect(createResult).ToNot(BeNil())
					Expect(*createResult.Success).Should(BeTrue())

				}
				listResult, listResp, listErr := glbService.ListAllLoadBalancers(glbService.NewListAllLoadBalancersOptions())
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())

			})
		})
	})
})
