/*
 * (C) Copyright IBM Corp. 2020.
 */

package zonelockdownv1_test

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.ibm.com/ibmcloud/networking-go-sdk/zonelockdownv1"
)

const configFile = "../cis.env"

var configLoaded bool = true

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe(`zonelockdownv1`, func() {
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
	globalOptions := &ZoneLockdownV1Options{
		ServiceName:    "cis_services",
		URL:            serviceURL,
		Authenticator:  authenticator,
		Crn:            &crn,
		ZoneIdentifier: &zone_id,
	}

	service, serviceErr := NewZoneLockdownV1(globalOptions)
	if serviceErr != nil {
		fmt.Println(serviceErr)
	}
	Describe(`zoneratelimitsv1_test`, func() {
		Context(`zoneratelimitsv1_test`, func() {
			BeforeEach(func() {
				shouldSkipTest()
				listOpt := service.NewListAllZoneLockownRulesOptions()
				listResult, listResp, listErr := service.ListAllZoneLockownRules(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())

				for _, rule := range listResult.Result {
					deleteOpt := service.NewDeleteZoneLockdownRuleOptions(*rule.ID)
					deleteResult, deleteResp, deleteErr := service.DeleteZoneLockdownRule(deleteOpt)
					Expect(deleteErr).To(BeNil())
					Expect(deleteResp).ToNot(BeNil())
					Expect(deleteResult).ToNot(BeNil())
					Expect(*deleteResult.Success).Should(BeTrue())
				}
			})
			AfterEach(func() {
				shouldSkipTest()
				listOpt := service.NewListAllZoneLockownRulesOptions()
				listResult, listResp, listErr := service.ListAllZoneLockownRules(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())

				for _, rule := range listResult.Result {
					deleteOpt := service.NewDeleteZoneLockdownRuleOptions(*rule.ID)
					deleteResult, deleteResp, deleteErr := service.DeleteZoneLockdownRule(deleteOpt)
					Expect(deleteErr).To(BeNil())
					Expect(deleteResp).ToNot(BeNil())
					Expect(deleteResult).ToNot(BeNil())
					Expect(*deleteResult.Success).Should(BeTrue())
				}
			})
			It(`zone lockdown by url`, func() {
				shouldSkipTest()
				config := LockdownInputConfigurationsItem{
					Target: core.StringPtr(LockdownInputConfigurationsItem_Target_Ip),
					Value:  core.StringPtr("198.51.100.4"),
				}
				configs := []LockdownInputConfigurationsItem{config}
				createOpt := service.NewCreateZoneLockdownRuleOptions()
				createOpt.SetConfigurations(configs)
				createOpt.SetDescription("Lockdown rule")
				createOpt.SetPaused(false)
				createOpt.SetUrls([]string{"api.mysite.com/some/endpoint*"})

				createResult, createResp, createErr := service.CreateZoneLockdownRule(createOpt)
				Expect(createErr).To(BeNil())
				Expect(createResp).ToNot(BeNil())
				Expect(createResult).ToNot(BeNil())
				Expect(*createResult.Success).Should(BeTrue())

				config = LockdownInputConfigurationsItem{
					Target: core.StringPtr(LockdownInputConfigurationsItem_Target_IpRange),
					Value:  core.StringPtr("192.51.100.4/24"),
				}
				configs = []LockdownInputConfigurationsItem{config}
				updateOpt := service.NewUpdateLockdownRuleOptions(*createResult.Result.ID)
				updateOpt.SetConfigurations(configs)
				updateOpt.SetDescription("Lockdown rule with ip range")
				updateOpt.SetPaused(true)
				updateOpt.SetUrls([]string{"api.mysite.com/some/endpoint*", "api.oursite.com/some/endpoint*"})

				updateResult, updateResp, updateErr := service.UpdateLockdownRule(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())

				getOpt := service.NewGetLockdownOptions(*createResult.Result.ID)
				getResult, getResp, getErr := service.GetLockdown(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				deleteOpt := service.NewDeleteZoneLockdownRuleOptions(*createResult.Result.ID)
				deleteResult, deleteResp, deleteErr := service.DeleteZoneLockdownRule(deleteOpt)
				Expect(deleteErr).To(BeNil())
				Expect(deleteResp).ToNot(BeNil())
				Expect(deleteResult).ToNot(BeNil())
				Expect(*deleteResult.Success).Should(BeTrue())
			})
			It(`list all lockdown rules test`, func() {
				shouldSkipTest()
				for i := 1; i < 5; i++ {
					ip := fmt.Sprintf("192.51.100.%d", i)
					desc := fmt.Sprintf("lockdown rule %d", i)
					url := fmt.Sprintf("api.mysite%d.com/some/endpoint*", i)
					config := LockdownInputConfigurationsItem{
						Target: core.StringPtr(LockdownInputConfigurationsItem_Target_Ip),
						Value:  core.StringPtr(ip),
					}
					configs := []LockdownInputConfigurationsItem{config}
					createOpt := service.NewCreateZoneLockdownRuleOptions()
					createOpt.SetConfigurations(configs)
					createOpt.SetDescription(desc)
					createOpt.SetPaused(false)
					createOpt.SetUrls([]string{url})
				}
				listOpt := service.NewListAllZoneLockownRulesOptions()
				listResult, listResp, listErr := service.ListAllZoneLockownRules(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())
			})
		})
	})
})
