/*
 * (C) Copyright IBM Corp. 2020.
 */

package firewallapiv1_test

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.ibm.com/ibmcloud/networking-go-sdk/firewallapiv1"
)

const configFile = "../cis.env"

var configLoaded bool = true

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe(`firewallapiv1_test`, func() {
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
	globalOptions := &FirewallApiV1Options{
		ServiceName:    "cis_services",
		URL:            serviceURL,
		Authenticator:  authenticator,
		Crn:            &crn,
		ZoneIdentifier: &zone_id,
	}

	service, serviceErr := NewFirewallApiV1(globalOptions)
	if serviceErr != nil {
		fmt.Println(serviceErr)
	}
	Describe(`zoneratelimitsv1_test`, func() {
		Context(`zoneratelimitsv1_test`, func() {
			It(`security level setting test`, func() {
				shouldSkipTest()
				values := []string{
					SetSecurityLevelSettingOptions_Value_High,
					SetSecurityLevelSettingOptions_Value_Low,
					SetSecurityLevelSettingOptions_Value_Medium,
					SetSecurityLevelSettingOptions_Value_UnderAttack,
					SetSecurityLevelSettingOptions_Value_EssentiallyOff,
				}
				opt := service.NewSetSecurityLevelSettingOptions()
				getOpt := service.NewGetSecurityLevelSettingOptions()
				for _, value := range values {
					opt.SetValue(value)
					setResult, setResp, setErr := service.SetSecurityLevelSetting(opt)
					Expect(setErr).To(BeNil())
					Expect(setResp).ToNot(BeNil())
					Expect(setResult).ToNot(BeNil())
					Expect(*setResult.Success).Should(BeTrue())
					Expect(*setResult.Result.Value).Should(BeEquivalentTo(value))

					getResult, getResp, getErr := service.GetSecurityLevelSetting(getOpt)
					Expect(getErr).To(BeNil())
					Expect(getResp).ToNot(BeNil())
					Expect(getResult).ToNot(BeNil())
					Expect(*getResult.Success).Should(BeTrue())
					Expect(*setResult.Result.Value).Should(BeEquivalentTo(value))
				}
			})
		})
	})
})
