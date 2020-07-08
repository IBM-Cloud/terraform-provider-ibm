/*
 * (C) Copyright IBM Corp. 2020.
 */

package useragentblockingrulesv1_test

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.ibm.com/ibmcloud/networking-go-sdk/useragentblockingrulesv1"
)

const configFile = "../cis.env"

var configLoaded bool = true

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe(`useragentblockingrulesv1`, func() {
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
	globalOptions := &UserAgentBlockingRulesV1Options{
		ServiceName:    "cis_services",
		URL:            serviceURL,
		Authenticator:  authenticator,
		Crn:            &crn,
		ZoneIdentifier: &zone_id,
	}

	service, serviceErr := NewUserAgentBlockingRulesV1(globalOptions)
	if serviceErr != nil {
		fmt.Println(serviceErr)
	}
	Describe(`list/create/update/delete/get ua rules api`, func() {
		Context(`list/create/update/delete/get ua rules api`, func() {
			BeforeEach(func() {
				shouldSkipTest()
				// list all rules
				listResult, listResp, listErr := service.ListAllZoneUserAgentRules(service.NewListAllZoneUserAgentRulesOptions())
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())

				for _, rule := range listResult.Result {
					// delete rules
					delOpt := service.NewDeleteZoneUserAgentRuleOptions(*rule.ID)
					delResult, delResp, delErr := service.DeleteZoneUserAgentRule(delOpt)
					Expect(delErr).To(BeNil())
					Expect(delResp).ToNot(BeNil())
					Expect(delResult).ToNot(BeNil())
					Expect(*delResult.Success).Should(BeTrue())
				}
			})
			AfterEach(func() {

				shouldSkipTest()
				// list all rules
				listResult, listResp, listErr := service.ListAllZoneUserAgentRules(service.NewListAllZoneUserAgentRulesOptions())
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())

				for _, rule := range listResult.Result {
					// delete rules
					delOpt := service.NewDeleteZoneUserAgentRuleOptions(*rule.ID)
					delResult, delResp, delErr := service.DeleteZoneUserAgentRule(delOpt)
					Expect(delErr).To(BeNil())
					Expect(delResp).ToNot(BeNil())
					Expect(delResult).ToNot(BeNil())
					Expect(*delResult.Success).Should(BeTrue())
				}
			})
			It(`list/create/update/delete/get ua rules api`, func() {
				shouldSkipTest()
				modes := []string{
					CreateZoneUserAgentRuleOptions_Mode_Block,
					CreateZoneUserAgentRuleOptions_Mode_Challenge,
					CreateZoneUserAgentRuleOptions_Mode_JsChallenge,
				}

				for i, mode := range modes {
					s := fmt.Sprintf("Mozilla/%d.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4", i+5)
					uaConfigOpt, uaConfigErr := service.NewUseragentRuleInputConfiguration("ua", s)
					Expect(uaConfigErr).To(BeNil())

					createOpt := service.NewCreateZoneUserAgentRuleOptions()
					createOpt.SetMode(mode)
					createOpt.SetConfiguration(uaConfigOpt)
					createOpt.SetDescription("Test user agent rule for " + mode)
					createOpt.SetPaused(true)

					createResult, createResp, createErr := service.CreateZoneUserAgentRule(createOpt)
					Expect(createErr).To(BeNil())
					Expect(createResp).ToNot(BeNil())
					Expect(createResult).ToNot(BeNil())
					Expect(*createResult.Success).Should(BeTrue())
				}

				// // list ua rules by page
				// listOpt := service.NewListAllZoneUserAgentRulesOptions()
				// listOpt.SetPage(2)
				// listOpt.SetPerPage(2)
				// listResult, listResp, listErr := service.ListAllZoneUserAgentRules(listOpt)
				// Expect(listErr).To(BeNil())
				// Expect(listResp).ToNot(BeNil())
				// Expect(listResult).ToNot(BeNil())
				// Expect(*listResult.Success).Should(BeTrue())

				// list all rules
				listResult, listResp, listErr := service.ListAllZoneUserAgentRules(service.NewListAllZoneUserAgentRulesOptions())
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())

				for i, rule := range listResult.Result {
					// get rules by id
					getOpt := service.NewGetUserAgentRuleOptions(*rule.ID)
					getResult, getResp, getErr := service.GetUserAgentRule(getOpt)
					Expect(getErr).To(BeNil())
					Expect(getResp).ToNot(BeNil())
					Expect(getResult).ToNot(BeNil())
					Expect(*getResult.Success).Should(BeTrue())

					// update rules
					var mode string
					updateOpt := service.NewUpdateUserAgentRuleOptions(*rule.ID)
					if *rule.Mode == UpdateUserAgentRuleOptions_Mode_Block || *rule.Mode == UpdateUserAgentRuleOptions_Mode_Challenge {
						mode = UpdateUserAgentRuleOptions_Mode_JsChallenge
					} else if *rule.Mode == UpdateUserAgentRuleOptions_Mode_Challenge || *rule.Mode == UpdateUserAgentRuleOptions_Mode_JsChallenge {
						mode = UpdateUserAgentRuleOptions_Mode_Block
					} else if *rule.Mode == UpdateUserAgentRuleOptions_Mode_Block || *rule.Mode == UpdateUserAgentRuleOptions_Mode_JsChallenge {
						mode = UpdateUserAgentRuleOptions_Mode_Challenge
					}

					s := fmt.Sprintf("Mozilla/%d.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4", i+10)
					uaConfigOpt, uaConfigErr := service.NewUseragentRuleInputConfiguration("ua", s)
					Expect(uaConfigErr).To(BeNil())

					updateOpt.SetMode(mode)
					updateOpt.SetConfiguration(uaConfigOpt)
					updateOpt.SetDescription("Test user agent rule for " + mode)
					updateOpt.SetPaused(false)

					updateResult, updateResp, updateErr := service.UpdateUserAgentRule(updateOpt)
					Expect(updateErr).To(BeNil())
					Expect(updateResp).ToNot(BeNil())
					Expect(updateResult).ToNot(BeNil())
					Expect(*updateResult.Success).Should(BeTrue())

					// delete rules
					delOpt := service.NewDeleteZoneUserAgentRuleOptions(*rule.ID)
					delResult, delResp, delErr := service.DeleteZoneUserAgentRule(delOpt)
					Expect(delErr).To(BeNil())
					Expect(delResp).ToNot(BeNil())
					Expect(delResult).ToNot(BeNil())
					Expect(*delResult.Success).Should(BeTrue())
				}

			})
		})
	})
})
