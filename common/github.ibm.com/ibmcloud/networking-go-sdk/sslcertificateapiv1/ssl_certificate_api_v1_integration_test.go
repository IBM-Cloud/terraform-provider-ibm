/*
 * (C) Copyright IBM Corp. 2020.
 */

package sslcertificateapiv1_test

import (
	"fmt"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.ibm.com/ibmcloud/networking-go-sdk/sslcertificateapiv1"
)

const configFile = "../cis.env"

var configLoaded bool = true

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe(`sslcertificateapiv1`, func() {
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
	url := os.Getenv("URL")
	certificate := os.Getenv("CERTIFICATE")
	updateCertificate := os.Getenv("UPDATE_CERTIFICATE")
	privateKey := os.Getenv("PRIVATE_KEY")
	updatePrivateKey := os.Getenv("UPDATE_PRIVATE_KEY")
	globalOptions := &SslCertificateApiV1Options{
		ServiceName:    "cis_services",
		URL:            serviceURL,
		Authenticator:  authenticator,
		Crn:            &crn,
		ZoneIdentifier: &zone_id,
	}

	service, serviceErr := NewSslCertificateApiV1(globalOptions)
	if serviceErr != nil {
		fmt.Println(serviceErr)
	}
	Describe(`order/view/delete ssl certificate packs`, func() {
		Context(`order/view/delete ssl certificate packs`, func() {
			BeforeEach(func() {
				shouldSkipTest()
				var result []DedicatedCertificatePack
				// sometime certificate deletion takes time, so keeping this code
				for i := 0; i < 5; i++ {
					deleteStatus := true
					// list all certificates
					getOpt := service.NewListCertificatesOptions()
					getOpt.SetXCorrelationID("12345")
					listResult, listResp, listErr := service.ListCertificates(getOpt)
					Expect(listErr).To(BeNil())
					Expect(listResp).ToNot(BeNil())
					Expect(listResult).ToNot(BeNil())
					result = listResult.Result
					for _, cert := range listResult.Result {
						if *cert.Status != "active" {
							fmt.Println("sleeping for 10 sec")
							time.Sleep(time.Second * 10)
							deleteStatus = false
						}
					}
					if deleteStatus == true {
						break
					}
				}

				for _, cert := range result {
					delOpt := service.NewDeleteCertificateOptions(*cert.ID)
					delResp, delErr := service.DeleteCertificate(delOpt)
					Expect(delErr).To(BeNil())
					Expect(delResp).ToNot(BeNil())
				}
			})
			AfterEach(func() {
				shouldSkipTest()
				var result []DedicatedCertificatePack
				// sometime certificate deletion takes time, so keeping this code
				for i := 0; i < 5; i++ {
					deleteStatus := true
					// list all certificates
					getOpt := service.NewListCertificatesOptions()
					getOpt.SetXCorrelationID("12345")
					listResult, listResp, listErr := service.ListCertificates(getOpt)
					Expect(listErr).To(BeNil())
					Expect(listResp).ToNot(BeNil())
					Expect(listResult).ToNot(BeNil())
					result = listResult.Result
					for _, cert := range listResult.Result {
						if *cert.Status != "active" {
							fmt.Println("sleeping for 10 sec")
							time.Sleep(time.Second * 10)
							deleteStatus = false
						}
					}
					if deleteStatus == true {
						break
					}
				}

				// delete all certificates
				for _, cert := range result {
					delOpt := service.NewDeleteCertificateOptions(*cert.ID)
					delResp, delErr := service.DeleteCertificate(delOpt)
					Expect(delErr).To(BeNil())
					Expect(delResp).ToNot(BeNil())
				}
			})
			It(`order/view/delete ssl certificate packs`, func() {
				shouldSkipTest()
				// order certificate packs
				orderOpt := service.NewOrderCertificateOptions()
				orderOpt.SetHosts([]string{url})
				orderOpt.SetXCorrelationID("12345")
				orderOpt.SetType(OrderCertificateOptions_Type_Dedicated)

				orderResult, orderResp, orderErr := service.OrderCertificate(orderOpt)
				Expect(orderErr).To(BeNil())
				Expect(orderResp).ToNot(BeNil())
				Expect(orderResult).ToNot(BeNil())

				// list all certificates
				getOpt := service.NewListCertificatesOptions()
				getOpt.SetXCorrelationID("12345")
				listResult, listResp, listErr := service.ListCertificates(getOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())

				// delete all certificates
				for _, cert := range listResult.Result {
					delOpt := service.NewDeleteCertificateOptions(*cert.ID)
					delResp, delErr := service.DeleteCertificate(delOpt)
					Expect(delErr).To(BeNil())
					Expect(delResp).ToNot(BeNil())
				}
			})
			It(`upload/view/delete ssl custom certificates`, func() {
				shouldSkipTest()
				// upload certificate packs
				geoOpt, geoErr := service.NewCustomCertReqGeoRestrictions("us")
				Expect(geoErr).To(BeNil())
				uploadOpt := service.NewUploadCustomCertificateOptions()
				uploadOpt.SetCertificate(certificate)
				uploadOpt.SetPrivateKey(privateKey)
				uploadOpt.SetGeoRestrictions(geoOpt)
				uploadOpt.SetBundleMethod(UpdateCustomCertificateOptions_BundleMethod_Optimal)

				uploadResult, uploadResp, uploadErr := service.UploadCustomCertificate(uploadOpt)
				Expect(uploadErr).To(BeNil())
				Expect(uploadResp).ToNot(BeNil())
				Expect(uploadResult).ToNot(BeNil())
				Expect(*uploadResult.Success).Should(BeTrue())

				// get custom certificate
				getOpt := service.NewGetCustomCertificateOptions(*uploadResult.Result.ID)
				getResult, getResp, getErr := service.GetCustomCertificate(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				// update custom certificate
				updateOpt := service.NewUpdateCustomCertificateOptions(*uploadResult.Result.ID)
				updateOpt.SetBundleMethod(UpdateCustomCertificateOptions_BundleMethod_Ubiquitous)
				updateOpt.SetCertificate(updateCertificate)
				updateOpt.SetGeoRestrictions(geoOpt)
				updateOpt.SetPrivateKey(updatePrivateKey)

				updateResult, updateResp, updateErr := service.UpdateCustomCertificate(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())

				// list all custom certificates
				listOpt := service.NewListCustomCertificatesOptions()
				listResult, listResp, listErr := service.ListCustomCertificates(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())

				// delete all custom certificates
				for _, cert := range listResult.Result {
					delOpt := service.NewDeleteCustomCertificateOptions(*cert.ID)
					delResp, delErr := service.DeleteCustomCertificate(delOpt)
					Expect(delErr).To(BeNil())
					Expect(delResp).ToNot(BeNil())
				}
			})
			It(`change/get/delete ssl universal certificate setting`, func() {
				shouldSkipTest()
				// upload custom certificate
				geoOpt, geoErr := service.NewCustomCertReqGeoRestrictions("us")
				Expect(geoErr).To(BeNil())
				uploadOpt := service.NewUploadCustomCertificateOptions()
				uploadOpt.SetCertificate(certificate)
				uploadOpt.SetPrivateKey(privateKey)
				uploadOpt.SetGeoRestrictions(geoOpt)
				uploadOpt.SetBundleMethod(UpdateCustomCertificateOptions_BundleMethod_Optimal)

				uploadResult, uploadResp, uploadErr := service.UploadCustomCertificate(uploadOpt)
				Expect(uploadErr).To(BeNil())
				Expect(uploadResp).ToNot(BeNil())
				Expect(uploadResult).ToNot(BeNil())
				Expect(*uploadResult.Success).Should(BeTrue())

				// get universal certificate setting
				getOpt := service.NewGetUniversalCertificateSettingOptions()
				getResult, getResp, getErr := service.GetUniversalCertificateSetting(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				// update universal certificate setting
				updateOpt := service.NewChangeUniversalCertificateSettingOptions()
				updateOpt.SetEnabled(false)

				updateResp, updateErr := service.ChangeUniversalCertificateSetting(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())

				// list all custom certificates
				listOpt := service.NewListCustomCertificatesOptions()
				listResult, listResp, listErr := service.ListCustomCertificates(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())

				// delete all custom certificates
				for _, cert := range listResult.Result {
					delOpt := service.NewDeleteCustomCertificateOptions(*cert.ID)
					delResp, delErr := service.DeleteCustomCertificate(delOpt)
					Expect(delErr).To(BeNil())
					Expect(delResp).ToNot(BeNil())
				}
			})
		})
	})
	Describe(`SSL Certificate Settings`, func() {
		It(`change/get ssl certificate setting`, func() {
			shouldSkipTest()
			values := []string{
				ChangeSslSettingOptions_Value_Full,
				ChangeSslSettingOptions_Value_Strict,
				ChangeSslSettingOptions_Value_Off,
				ChangeSslSettingOptions_Value_Flexible,
			}
			changeOpt := service.NewChangeSslSettingOptions()
			for _, val := range values {

				// change ssl certificate setting
				changeOpt.SetValue(val)
				changeResult, changeResp, changeErr := service.ChangeSslSetting(changeOpt)
				Expect(changeErr).To(BeNil())
				Expect(changeResp).ToNot(BeNil())
				Expect(changeResult).ToNot(BeNil())
				Expect(*changeResult.Success).Should(BeTrue())

				getOpt := service.NewGetSslSettingOptions()
				getResult, getResp, getErr := service.GetSslSetting(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())
			}
		})
		It(`change/get TLS version 1.2 setting`, func() {
			shouldSkipTest()
			// get TLS version 1.2 setting
			getOpt := service.NewGetTls12SettingOptions()
			getResult, getResp, getErr := service.GetTls12Setting(getOpt)
			Expect(getErr).To(BeNil())
			Expect(getResp).ToNot(BeNil())
			Expect(getResult).ToNot(BeNil())
			Expect(*getResult.Success).Should(BeTrue())

			// change TLS version 1.2 setting
			changeOpt := service.NewChangeTls12SettingOptions()
			changeOpt.SetValue(ChangeTls13SettingOptions_Value_Off)
			changeResult, changeResp, changeErr := service.ChangeTls12Setting(changeOpt)
			Expect(changeErr).To(BeNil())
			Expect(changeResp).ToNot(BeNil())
			Expect(changeResult).ToNot(BeNil())
			Expect(*changeResult.Success).Should(BeTrue())
		})
		It(`change/get TLS version 1.3 setting`, func() {
			shouldSkipTest()
			// get TLS version 1.3 setting
			getOpt := service.NewGetTls13SettingOptions()
			getResult, getResp, getErr := service.GetTls13Setting(getOpt)
			Expect(getErr).To(BeNil())
			Expect(getResp).ToNot(BeNil())
			Expect(getResult).ToNot(BeNil())
			Expect(*getResult.Success).Should(BeTrue())

			// change TLS version 1.3 setting
			changeOpt := service.NewChangeTls13SettingOptions()
			changeOpt.SetValue(ChangeTls13SettingOptions_Value_Off)
			if *getResult.Result.Value == ChangeTls13SettingOptions_Value_Off {
				changeOpt.SetValue(ChangeTls13SettingOptions_Value_On)
			}
			changeResult, changeResp, changeErr := service.ChangeTls13Setting(changeOpt)
			Expect(changeErr).To(BeNil())
			Expect(changeResp).ToNot(BeNil())
			Expect(changeResult).ToNot(BeNil())
			Expect(*changeResult.Success).Should(BeTrue())
		})
	})
})
