/*
 * (C) Copyright IBM Corp. 2020.
 */

package wafrulesapiv1_test

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/ibmcloud/networking-go-sdk/wafrulepackagesapiv1"
	. "github.ibm.com/ibmcloud/networking-go-sdk/wafrulesapiv1"
)

const configFile = "../cis.env"

var configLoaded bool = true

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe(`wafrulesapiv1`, func() {
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
	globalOptions := &WafRulesApiV1Options{
		ServiceName:   "cis_services",
		URL:           serviceURL,
		Authenticator: authenticator,
		Crn:           &crn,
		ZoneID:        &zone_id,
	}
	wafPackOptions := &wafrulepackagesapiv1.WafRulePackagesApiV1Options{
		ServiceName:   "cis_services",
		URL:           serviceURL,
		Authenticator: authenticator,
		Crn:           &crn,
		ZoneID:        &zone_id,
	}
	wafPackService, wafPackServiceErr := wafrulepackagesapiv1.NewWafRulePackagesApiV1(wafPackOptions)
	if wafPackServiceErr != nil {
		fmt.Println(wafPackServiceErr)
	}
	wafRuleService, wafRuleServiceErr := NewWafRulesApiV1(globalOptions)
	if wafRuleServiceErr != nil {
		fmt.Println(wafRuleServiceErr)
	}
	Describe(`list/update/get waf rules api`, func() {
		Context(`list/update/get waf rules api`, func() {
			cisBodyList := []string{
				WafRuleBodyCis_Mode_Block,
				WafRuleBodyCis_Mode_Challenge,
				WafRuleBodyCis_Mode_Disable,
				WafRuleBodyCis_Mode_Simulate,
				WafRuleBodyCis_Mode_Default,
			}
			It(`list/update/get waf rules api`, func() {
				shouldSkipTest()

				// list all WAF Packages
				result, resp, err := wafPackService.ListWafPackages(wafPackService.NewListWafPackagesOptions())
				Expect(err).To(BeNil())
				Expect(resp).ToNot(BeNil())
				Expect(result).ToNot(BeNil())
				Expect(*result.Success).Should(BeTrue())

				for _, pack := range result.Result {
					// list all WAF Rules
					listResult, listResp, listErr := wafRuleService.ListWafRules(wafRuleService.NewListWafRulesOptions(*pack.ID))
					Expect(listErr).To(BeNil())
					Expect(listResp).ToNot(BeNil())
					Expect(listResult).ToNot(BeNil())
					Expect(*listResult.Success).Should(BeTrue())

					for _, rule := range listResult.Result {
						fmt.Println("rule id: ", *rule.ID)
						fmt.Println("modes : ", *rule.Mode)

						if *rule.Mode == WafRuleBodyCis_Mode_Default ||
							*rule.Mode == WafRuleBodyCis_Mode_Block ||
							*rule.Mode == WafRuleBodyCis_Mode_Challenge ||
							*rule.Mode == WafRuleBodyCis_Mode_Disable ||
							*rule.Mode == WafRuleBodyCis_Mode_Simulate {
							for _, cis := range cisBodyList {

								cisOpt, cisErr := wafRuleService.NewWafRuleBodyCis(cis)
								Expect(cisErr).To(BeNil())

								updateOpt := wafRuleService.NewUpdateWafRuleOptions(*pack.ID, *rule.ID)
								updateOpt.SetCis(cisOpt)

								// udpate all WAF Rules
								updateResult, updateResp, updateErr := wafRuleService.UpdateWafRule(updateOpt)
								Expect(updateErr).To(BeNil())
								Expect(updateResp).ToNot(BeNil())
								Expect(updateResult).ToNot(BeNil())
								Expect(*updateResult.Success).Should(BeTrue())

							}
						} else {
							var mode string = WafRuleBodyOwasp_Mode_On
							if *rule.Mode == mode {
								mode = WafRuleBodyOwasp_Mode_Off
							}
							owaspOpt, owaspErr := wafRuleService.NewWafRuleBodyOwasp(mode)
							Expect(owaspErr).To(BeNil())

							updateOpt := wafRuleService.NewUpdateWafRuleOptions(*pack.ID, *rule.ID)
							updateOpt.SetOwasp(owaspOpt)

							// udpate all WAF Rules
							updateResult, updateResp, updateErr := wafRuleService.UpdateWafRule(updateOpt)
							Expect(updateErr).To(BeNil())
							Expect(updateResp).ToNot(BeNil())
							Expect(updateResult).ToNot(BeNil())
							Expect(*updateResult.Success).Should(BeTrue())
						}

						// get WAF rules by id
						getOpt := wafRuleService.NewGetWafRuleOptions(*pack.ID, *rule.ID)
						getResult, getResp, getErr := wafRuleService.GetWafRule(getOpt)
						Expect(getErr).To(BeNil())
						Expect(getResp).ToNot(BeNil())
						Expect(getResult).ToNot(BeNil())
						Expect(*getResult.Success).Should(BeTrue())
					}
				}
			})
		})
	})
})
