/*
 * (C) Copyright IBM Corp. 2020.
 */

package wafapiv1_test

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.ibm.com/ibmcloud/networking-go-sdk/wafapiv1"
)

const configFile = "../cis.env"

var configLoaded bool = true

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe(`securityeventsapiv1_test`, func() {
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
	globalOptions := &WafApiV1Options{
		ServiceName:   "cis_services",
		URL:           serviceURL,
		Authenticator: authenticator,
		Crn:           &crn,
		ZoneID:        &zone_id,
	}

	service, serviceErr := NewWafApiV1(globalOptions)
	if serviceErr != nil {
		fmt.Println(serviceErr)
	}
	Describe(`wafapiv1_test`, func() {
		Context(`wafapiv1_test`, func() {
			It(`waf setting on/off test`, func() {
				shouldSkipTest()

				updateOpt := service.NewUpdateWafSettingsOptions()
				updateOpt.SetValue(UpdateWafSettingsOptions_Value_On)
				updateResult, updateResp, updateErr := service.UpdateWafSettings(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())

				listOpt := service.NewGetWafSettingsOptions()
				listResult, listResp, listErr := service.GetWafSettings(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())

				updateOpt.SetValue(UpdateWafSettingsOptions_Value_Off)
				updateResult, updateResp, updateErr = service.UpdateWafSettings(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())

			})
		})
	})
})
