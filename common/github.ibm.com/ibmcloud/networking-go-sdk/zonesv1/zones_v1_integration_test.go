/*
 * (C) Copyright IBM Corp. 2020.
 */

package zonesv1_test

import (
	"fmt"
	"os"
	"strings"

	"github.com/IBM/go-sdk-core/core"
	guuid "github.com/google/uuid"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.ibm.com/ibmcloud/networking-go-sdk/zonesv1"
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
	globalOptions := &ZonesV1Options{
		ServiceName:   "cis_services",
		URL:           serviceURL,
		Authenticator: authenticator,
		Crn:           &crn,
	}
	url := os.Getenv("URL")
	service, serviceErr := NewZonesV1(globalOptions)
	if serviceErr != nil {
		fmt.Println(serviceErr)
	}
	Describe(`zonesv1_test`, func() {
		Context(`zonesv1_test`, func() {
			BeforeEach(func() {
				shouldSkipTest()
				listOpt := service.NewListZonesOptions()
				listResult, listResp, listErr := service.ListZones(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())

				for _, zone := range listResult.Result {
					if strings.Contains(*zone.Name, "uuid-") {
						deleteOpt := service.NewDeleteZoneOptions(*zone.ID)
						deleteResult, deleteResp, deleteErr := service.DeleteZone(deleteOpt)
						Expect(deleteErr).To(BeNil())
						Expect(deleteResp).ToNot(BeNil())
						Expect(deleteResult).ToNot(BeNil())
						Expect(*deleteResult.Success).Should(BeTrue())
					}
				}
			})
			AfterEach(func() {
				shouldSkipTest()
				listOpt := service.NewListZonesOptions()
				listResult, listResp, listErr := service.ListZones(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())

				for _, zone := range listResult.Result {
					if strings.Contains(*zone.Name, "uuid-") {
						deleteOpt := service.NewDeleteZoneOptions(*zone.ID)
						deleteResult, deleteResp, deleteErr := service.DeleteZone(deleteOpt)
						Expect(deleteErr).To(BeNil())
						Expect(deleteResp).ToNot(BeNil())
						Expect(deleteResult).ToNot(BeNil())
						Expect(*deleteResult.Success).Should(BeTrue())
					}
				}
			})
			It(`zones create/update/delete/activation check test`, func() {
				shouldSkipTest()
				// create zone
				zoneName := fmt.Sprintf("uuid-%s.%s", guuid.New().String()[1:6], url)
				createOpt := service.NewCreateZoneOptions()
				createOpt.SetName(zoneName)

				createResult, createResp, createErr := service.CreateZone(createOpt)
				Expect(createErr).To(BeNil())
				Expect(createResp).ToNot(BeNil())
				Expect(createResult).ToNot(BeNil())
				Expect(*createResult.Success).Should(BeTrue())

				// update zone
				updateOpt := service.NewUpdateZoneOptions(*createResult.Result.ID)
				updateOpt.SetPaused(true)
				updateResult, updateResp, updateErr := service.UpdateZone(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())

				// activation check
				checkOpt := service.NewZoneActivationCheckOptions(*createResult.Result.ID)
				checkResult, checkResp, checkErr := service.ZoneActivationCheck(checkOpt)
				Expect(checkErr).To(BeNil())
				Expect(checkResp).ToNot(BeNil())
				Expect(checkResult).ToNot(BeNil())
				Expect(*checkResult.Success).Should(BeTrue())

				getOpt := service.NewGetZoneOptions(*createResult.Result.ID)
				getResult, getResp, getErr := service.GetZone(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				deleteOpt := service.NewDeleteZoneOptions(*createResult.Result.ID)
				deleteResult, deleteResp, deleteErr := service.DeleteZone(deleteOpt)
				Expect(deleteErr).To(BeNil())
				Expect(deleteResp).ToNot(BeNil())
				Expect(deleteResult).ToNot(BeNil())
				Expect(*deleteResult.Success).Should(BeTrue())
			})
			It(`list all zones test`, func() {
				shouldSkipTest()
				// create zone
				for i := 0; i < 5; i++ {
					zoneName := fmt.Sprintf("uuid-%s.%s", guuid.New().String(), url)
					createOpt := service.NewCreateZoneOptions()
					createOpt.SetName(zoneName)

					createResult, createResp, createErr := service.CreateZone(createOpt)
					Expect(createErr).To(BeNil())
					Expect(createResp).ToNot(BeNil())
					Expect(createResult).ToNot(BeNil())
					Expect(*createResult.Success).Should(BeTrue())
				}

				// list all zones
				listOpt := service.NewListZonesOptions()
				listResult, listResp, listErr := service.ListZones(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())

			})
		})
	})
})
