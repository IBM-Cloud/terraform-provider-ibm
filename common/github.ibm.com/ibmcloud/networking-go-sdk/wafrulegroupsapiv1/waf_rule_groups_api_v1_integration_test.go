/*
 * (C) Copyright IBM Corp. 2020.
 */

package wafrulegroupsapiv1_test

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.ibm.com/ibmcloud/networking-go-sdk/wafrulegroupsapiv1"
	"github.ibm.com/ibmcloud/networking-go-sdk/wafrulepackagesapiv1"
)

const configFile = "../cis.env"

var configLoaded bool = true

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe(`wafrulegroupsapiv1_test`, func() {
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
	globalOptions := &WafRuleGroupsApiV1Options{
		ServiceName:   "cis_services",
		URL:           serviceURL,
		Authenticator: authenticator,
		Crn:           &crn,
		ZoneID:        &zone_id,
	}
	service, serviceErr := NewWafRuleGroupsApiV1(globalOptions)
	if serviceErr != nil {
		fmt.Println(serviceErr)
	}

	wafPackageOptions := &wafrulepackagesapiv1.WafRulePackagesApiV1Options{
		ServiceName:   "cis_services",
		URL:           serviceURL,
		Authenticator: authenticator,
		Crn:           &crn,
		ZoneID:        &zone_id,
	}

	packService, packServiceErr := wafrulepackagesapiv1.NewWafRulePackagesApiV1(wafPackageOptions)
	if packServiceErr != nil {
		fmt.Println(packServiceErr)
	}

	Describe(`wafrulegroupsapiv1_test`, func() {
		Context("wafrulegroupsapiv1_test", func() {
			It(`waf rule group test`, func() {
				shouldSkipTest()

				// list waf rule packages
				listPackResult, listPackResp, listPackErr := packService.ListWafPackages(packService.NewListWafPackagesOptions())
				Expect(listPackErr).To(BeNil())
				Expect(listPackResp).ToNot(BeNil())
				Expect(listPackResult).ToNot(BeNil())
				Expect(*listPackResult.Success).Should(BeTrue())

				for _, pack := range listPackResult.Result {
					// list waf rule groups
					listResult, listResp, listErr := service.ListWafRuleGroups(service.NewListWafRuleGroupsOptions(*pack.ID))
					Expect(listErr).To(BeNil())
					Expect(listResp).ToNot(BeNil())
					Expect(listResult).ToNot(BeNil())
					Expect(*listResult.Success).Should(BeTrue())

					for _, grp := range listResult.Result {

						// get WAF rules group
						getOpt := service.NewGetWafRuleGroupOptions(*pack.ID, *grp.ID)
						getResult, getResp, getErr := service.GetWafRuleGroup(getOpt)
						Expect(getErr).To(BeNil())
						Expect(getResp).ToNot(BeNil())
						Expect(getResult).ToNot(BeNil())
						Expect(*getResult.Success).Should(BeTrue())

						// update waf rule group mode
						updateOpt := service.NewUpdateWafRuleGroupOptions(*pack.ID, *grp.ID)
						if *grp.Mode == UpdateWafRuleGroupOptions_Mode_Off {
							updateOpt.SetMode(UpdateWafRuleGroupOptions_Mode_On)
						} else {
							updateOpt.SetMode(UpdateWafRuleGroupOptions_Mode_Off)
						}
						updateResult, updateResp, updateErr := service.UpdateWafRuleGroup(updateOpt)
						Expect(updateErr).To(BeNil())
						Expect(updateResp).ToNot(BeNil())
						Expect(updateResult).ToNot(BeNil())
						Expect(*updateResult.Success).Should(BeTrue())

					}
				}

			})

		})
	})
})
