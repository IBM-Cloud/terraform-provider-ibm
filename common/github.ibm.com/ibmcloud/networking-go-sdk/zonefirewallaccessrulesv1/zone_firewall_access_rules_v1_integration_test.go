/*
 * (C) Copyright IBM Corp. 2020.
 */

package zonefirewallaccessrulesv1_test

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.ibm.com/ibmcloud/networking-go-sdk/zonefirewallaccessrulesv1"
)

const configFile = "../cis.env"

var configLoaded bool = true

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe(`ZoneFirewallAccessRulesV1`, func() {
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
	globalOptions := &ZoneFirewallAccessRulesV1Options{
		ServiceName:    "cis_services",
		URL:            serviceURL,
		Authenticator:  authenticator,
		Crn:            &crn,
		ZoneIdentifier: &zone_id,
	}
	testService, testServiceErr := NewZoneFirewallAccessRulesV1(globalOptions)
	if testServiceErr != nil {
		fmt.Println(testServiceErr)
	}
	// CIS_Frontend_API_Spec-Zone_Firewall_AccessRules.yaml file integration test starts here

	Describe(`ZoneAccessRuleActions`, func() {
		Context("ZoneAccessRuleActions", func() {
			BeforeEach(func() {
				shouldSkipTest()
				result, response, operationErr := testService.ListAllZoneAccessRules(testService.NewListAllZoneAccessRulesOptions())
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				for _, rule := range result.Result {
					option := testService.NewDeleteZoneAccessRuleOptions(*rule.ID)
					result, response, operationErr := testService.DeleteZoneAccessRule(option)
					Expect(operationErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
					Expect(*result.Success).Should(BeTrue())
				}
			})
			AfterEach(func() {
				shouldSkipTest()
				result, response, operationErr := testService.ListAllZoneAccessRules(testService.NewListAllZoneAccessRulesOptions())
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				for _, rule := range result.Result {
					option := testService.NewDeleteZoneAccessRuleOptions(*rule.ID)
					result, response, operationErr := testService.DeleteZoneAccessRule(option)
					Expect(operationErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
					Expect(*result.Success).Should(BeTrue())
				}
			})
			It(`create/update/delete/get zone access rules`, func() {
				shouldSkipTest()
				modes := [4]string{
					CreateZoneAccessRuleOptions_Mode_Block,
					CreateZoneAccessRuleOptions_Mode_Challenge,
					CreateZoneAccessRuleOptions_Mode_JsChallenge,
					CreateZoneAccessRuleOptions_Mode_Whitelist,
				}
				targets := [4]string{
					ZoneAccessRuleObjectConfiguration_Target_Ip,
					ZoneAccessRuleObjectConfiguration_Target_IpRange,
					ZoneAccessRuleObjectConfiguration_Target_Asn,
					ZoneAccessRuleObjectConfiguration_Target_Country,
				}
				configData := [4]string{
					"172.168.1.1",
					"172.168.1.0/24",
					"AS12345",
					"US",
				}

				newModes := [4]string{
					CreateZoneAccessRuleOptions_Mode_Challenge,
					CreateZoneAccessRuleOptions_Mode_JsChallenge,
					CreateZoneAccessRuleOptions_Mode_Whitelist,
					CreateZoneAccessRuleOptions_Mode_Block,
				}

				for i, mode := range modes {
					options := testService.NewCreateZoneAccessRuleOptions()
					options.SetMode(mode)

					configOpt, err := testService.NewZoneAccessRuleInputConfiguration(targets[i], configData[i])
					Expect(err).To(BeNil())
					Expect(configOpt).ToNot(BeNil())

					options.SetConfiguration(configOpt)
					options.SetNotes("This rule is added because of event X that occurred on date xyz")
					// create zone access rule
					result, response, err := testService.CreateZoneAccessRule(options)
					Expect(err).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())

					// get zone access rule
					getOption := testService.NewGetZoneAccessRuleOptions(*result.Result.ID)
					result, response, err = testService.GetZoneAccessRule(getOption)
					Expect(err).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
					Expect(*result.Success).Should(BeTrue())

					// update zone access rule
					updateOpt := testService.NewUpdateZoneAccessRuleOptions(*result.Result.ID)
					updateOpt.SetMode(newModes[i])
					result, response, err = testService.UpdateZoneAccessRule(updateOpt)
					Expect(err).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
					Expect(*result.Success).Should(BeTrue())

					// delete zone access rule
					deleteOpt := testService.NewDeleteZoneAccessRuleOptions(*result.Result.ID)
					delResult, response, err := testService.DeleteZoneAccessRule(deleteOpt)
					Expect(err).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(delResult).ToNot(BeNil())
					Expect(*result.Success).Should(BeTrue())
				}
			})
		})
	})

	Describe(`ListAllZoneAccessRules`, func() {
		Context("ZoneAccessRuleActions", func() {
			BeforeEach(func() {
				shouldSkipTest()
				result, response, operationErr := testService.ListAllZoneAccessRules(testService.NewListAllZoneAccessRulesOptions())
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				for _, rule := range result.Result {
					option := testService.NewDeleteZoneAccessRuleOptions(*rule.ID)
					result, response, operationErr := testService.DeleteZoneAccessRule(option)
					Expect(operationErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
					Expect(*result.Success).Should(BeTrue())
				}
			})
			AfterEach(func() {
				shouldSkipTest()
				result, response, operationErr := testService.ListAllZoneAccessRules(testService.NewListAllZoneAccessRulesOptions())
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				for _, rule := range result.Result {
					option := testService.NewDeleteZoneAccessRuleOptions(*rule.ID)
					result, response, operationErr := testService.DeleteZoneAccessRule(option)
					Expect(operationErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
					Expect(*result.Success).Should(BeTrue())
				}
			})
			It(`Invoke ListAllZoneAccessRules successfully`, func() {
				shouldSkipTest()
				modes := [4]string{
					CreateZoneAccessRuleOptions_Mode_Block,
					CreateZoneAccessRuleOptions_Mode_Challenge,
					CreateZoneAccessRuleOptions_Mode_JsChallenge,
					CreateZoneAccessRuleOptions_Mode_Whitelist,
				}
				targets := [4]string{
					ZoneAccessRuleObjectConfiguration_Target_Ip,
					ZoneAccessRuleObjectConfiguration_Target_IpRange,
					ZoneAccessRuleObjectConfiguration_Target_Asn,
					ZoneAccessRuleObjectConfiguration_Target_Country,
				}
				configData := [4]string{
					"172.168.1.1",
					"172.168.1.0/24",
					"AS12345",
					"US",
				}

				for i, mode := range modes {
					options := testService.NewCreateZoneAccessRuleOptions()
					options.SetMode(mode)

					configOpt, err := testService.NewZoneAccessRuleInputConfiguration(targets[i], configData[i])
					Expect(err).To(BeNil())
					Expect(configOpt).ToNot(BeNil())

					options.SetConfiguration(configOpt)
					options.SetNotes("This rule is added because of event X that occurred on date xyz")
					// create zone access rule
					result, response, err := testService.CreateZoneAccessRule(options)
					Expect(err).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
				}

				result, response, operationErr := testService.ListAllZoneAccessRules(testService.NewListAllZoneAccessRulesOptions())
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				for _, rule := range result.Result {
					option := testService.NewDeleteZoneAccessRuleOptions(*rule.ID)
					result, response, operationErr := testService.DeleteZoneAccessRule(option)
					Expect(operationErr).To(BeNil())
					Expect(response).ToNot(BeNil())
					Expect(result).ToNot(BeNil())
					Expect(*result.Success).Should(BeTrue())
				}
			})
		})
	})
})
