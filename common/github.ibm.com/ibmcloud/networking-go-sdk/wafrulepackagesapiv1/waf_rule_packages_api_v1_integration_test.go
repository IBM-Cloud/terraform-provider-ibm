/*
 * (C) Copyright IBM Corp. 2020.
 */

package wafrulepackagesapiv1_test

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.ibm.com/ibmcloud/networking-go-sdk/wafrulepackagesapiv1"
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
	globalOptions := &WafRulePackagesApiV1Options{
		ServiceName:   "cis_services",
		URL:           serviceURL,
		Authenticator: authenticator,
		Crn:           &crn,
		ZoneID:        &zone_id,
	}
	testService, testServiceErr := NewWafRulePackagesApiV1(globalOptions)
	if testServiceErr != nil {
		fmt.Println(testServiceErr)
	}
	Describe(`list/update/get waf rule packages api`, func() {
		Context(`list/update/get waf rule packages api`, func() {
			It(`list/update/get waf rule packages api`, func() {
				shouldSkipTest()
				listResult, listResp, listErr := testService.ListWafPackages(testService.NewListWafPackagesOptions())
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())

				sensitivities := []string{
					UpdateWafPackageOptions_Sensitivity_High,
					UpdateWafPackageOptions_Sensitivity_Low,
					UpdateWafPackageOptions_Sensitivity_Medium,
					UpdateWafPackageOptions_Sensitivity_Off,
				}
				modes := []string{
					UpdateWafPackageOptions_ActionMode_Block,
					UpdateWafPackageOptions_ActionMode_Challenge,
					UpdateWafPackageOptions_ActionMode_Simulate,
				}

				for _, pack := range listResult.Result {

					// get WAF rules packages
					getOpt := testService.NewGetWafPackageOptions(*pack.ID)
					getResult, getResp, getErr := testService.GetWafPackage(getOpt)
					Expect(getErr).To(BeNil())
					Expect(getResp).ToNot(BeNil())
					Expect(getResult).ToNot(BeNil())
					Expect(*getResult.Success).Should(BeTrue())
					if *pack.DetectionMode == "anomaly" {
						for _, sensitivity := range sensitivities {
							for _, mode := range modes {
								updateOpt := testService.NewUpdateWafPackageOptions(*pack.ID)
								updateOpt.SetActionMode(mode)
								updateOpt.SetSensitivity(sensitivity)
								updateResult, updateResp, updateErr := testService.UpdateWafPackage(updateOpt)
								Expect(updateErr).To(BeNil())
								Expect(updateResp).ToNot(BeNil())
								Expect(updateResult).ToNot(BeNil())
								Expect(*updateResult.Success).Should(BeTrue())
							}
						}
					}
				}
			})
		})
	})
})
