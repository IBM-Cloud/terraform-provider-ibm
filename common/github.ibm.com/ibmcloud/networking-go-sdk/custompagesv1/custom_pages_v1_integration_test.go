/*
 * (C) Copyright IBM Corp. 2020.
 */

package custompagesv1_test

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.ibm.com/ibmcloud/networking-go-sdk/custompagesv1"
)

const configFile = "../cis.env"

var configLoaded bool = true

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe(`zoneratelimitsv1`, func() {
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
	url := os.Getenv("CUSTOM_PAGE_URL")
	globalOptions := &CustomPagesV1Options{
		ServiceName:    "cis_services",
		URL:            serviceURL,
		Authenticator:  authenticator,
		Crn:            &crn,
		ZoneIdentifier: &zone_id,
	}

	service, serviceErr := NewCustomPagesV1(globalOptions)
	if serviceErr != nil {
		fmt.Println(serviceErr)
	}
	Describe(`custompagesv1_test`, func() {
		Context(`custompagesv1_test`, func() {
			It(`zone custom pages test`, func() {
				shouldSkipTest()
				page_ids := []string{
					UpdateZoneCustomPageOptions_PageIdentifier_BasicChallenge,
					UpdateZoneCustomPageOptions_PageIdentifier_CountryChallenge,
					UpdateZoneCustomPageOptions_PageIdentifier_IpBlock,
					UpdateZoneCustomPageOptions_PageIdentifier_RatelimitBlock,
					UpdateZoneCustomPageOptions_PageIdentifier_UnderAttack,
					UpdateZoneCustomPageOptions_PageIdentifier_WafBlock,
					UpdateZoneCustomPageOptions_PageIdentifier_WafChallenge,
				}
				for _, id := range page_ids {

					// update zone custom pages
					updateOpt := service.NewUpdateZoneCustomPageOptions(id)
					updateOpt.SetState(CustomPageObject_State_Customized)
					updateOpt.SetURL(url)
					updateResult, updateResp, updateErr := service.UpdateZoneCustomPage(updateOpt)
					Expect(updateErr).To(BeNil())
					Expect(updateResp).ToNot(BeNil())
					Expect(updateResult).ToNot(BeNil())
					Expect(*updateResult.Success).Should(BeTrue())

					// get zone custom pages
					getOpt := service.NewGetZoneCustomPageOptions(id)
					getResult, getResp, getErr := service.GetZoneCustomPage(getOpt)
					Expect(getErr).To(BeNil())
					Expect(getResp).ToNot(BeNil())
					Expect(getResult).ToNot(BeNil())
					Expect(*getResult.Success).Should(BeTrue())

					// reset zone custom pages
					updateOpt = service.NewUpdateZoneCustomPageOptions(id)
					updateOpt.SetState(CustomPageObject_State_Default)
					updateOpt.SetURL(url)
					updateResult, updateResp, updateErr = service.UpdateZoneCustomPage(updateOpt)
					Expect(updateErr).To(BeNil())
					Expect(updateResp).ToNot(BeNil())
					Expect(updateResult).ToNot(BeNil())
					Expect(*updateResult.Success).Should(BeTrue())
				}

				listOpt := service.NewListZoneCustomPagesOptions()
				listResult, listResp, listErr := service.ListZoneCustomPages(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())
			})
			It(`zone custom pages test`, func() {
				shouldSkipTest()
				page_ids := []string{
					UpdateInstanceCustomPageOptions_PageIdentifier_BasicChallenge,
					UpdateInstanceCustomPageOptions_PageIdentifier_CountryChallenge,
					UpdateInstanceCustomPageOptions_PageIdentifier_IpBlock,
					UpdateInstanceCustomPageOptions_PageIdentifier_RatelimitBlock,
					UpdateInstanceCustomPageOptions_PageIdentifier_UnderAttack,
					UpdateInstanceCustomPageOptions_PageIdentifier_WafBlock,
					UpdateInstanceCustomPageOptions_PageIdentifier_WafChallenge,
				}
				for _, id := range page_ids {

					// update instance custom pages
					updateOpt := service.NewUpdateInstanceCustomPageOptions(id)
					updateOpt.SetState(CustomPageObject_State_Customized)
					updateOpt.SetURL(url)
					updateResult, updateResp, updateErr := service.UpdateInstanceCustomPage(updateOpt)
					Expect(updateErr).To(BeNil())
					Expect(updateResp).ToNot(BeNil())
					Expect(updateResult).ToNot(BeNil())
					Expect(*updateResult.Success).Should(BeTrue())

					// get instance custom pages
					getOpt := service.NewGetInstanceCustomPageOptions(id)
					getResult, getResp, getErr := service.GetInstanceCustomPage(getOpt)
					Expect(getErr).To(BeNil())
					Expect(getResp).ToNot(BeNil())
					Expect(getResult).ToNot(BeNil())
					Expect(*getResult.Success).Should(BeTrue())

					// reset instance custom pages
					updateOpt = service.NewUpdateInstanceCustomPageOptions(id)
					updateOpt.SetState(CustomPageObject_State_Default)
					updateOpt.SetURL(url)
					updateResult, updateResp, updateErr = service.UpdateInstanceCustomPage(updateOpt)
					Expect(updateErr).To(BeNil())
					Expect(updateResp).ToNot(BeNil())
					Expect(updateResult).ToNot(BeNil())
					Expect(*updateResult.Success).Should(BeTrue())
				}

				listOpt := service.NewListInstanceCustomPagesOptions()
				listResult, listResp, listErr := service.ListInstanceCustomPages(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())
			})
		})
	})
})
