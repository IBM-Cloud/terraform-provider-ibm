/*
 * (C) Copyright IBM Corp. 2020.
 */

package cachingapiv1_test

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.ibm.com/ibmcloud/networking-go-sdk/cachingapiv1"
)

const configFile = "../cis.env"

var configLoaded bool = true

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping...")
	}
}

var _ = Describe(`cachingapiv1_test`, func() {
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
	globalOptions := &CachingApiV1Options{
		ServiceName:   "cis_services",
		URL:           serviceURL,
		Authenticator: authenticator,
		Crn:           &crn,
		ZoneID:        &zone_id,
	}

	service, serviceErr := NewCachingApiV1(globalOptions)
	if serviceErr != nil {
		fmt.Println(serviceErr)
	}
	Describe(`zoneratelimitsv1_test`, func() {
		Context(`zoneratelimitsv1_test`, func() {
			It(`cache purge by urls`, func() {
				shouldSkipTest()
				cacheOpt := service.NewPurgeByUrlsOptions()
				cacheOpt.SetFiles([]string{"http://www.example.com/cat_picture.jpg"})
				cacheResult, cacheResp, cacheErr := service.PurgeByUrls(cacheOpt)
				Expect(cacheErr).To(BeNil())
				Expect(cacheResp).ToNot(BeNil())
				Expect(cacheResult).ToNot(BeNil())
				Expect(*cacheResult.Success).Should(BeTrue())
			})
			It(`cache purge by cache tags`, func() {
				shouldSkipTest()
				cacheOpt := service.NewPurgeByCacheTagsOptions()
				cacheOpt.SetTags([]string{"some-tags"})
				cacheResult, cacheResp, cacheErr := service.PurgeByCacheTags(cacheOpt)
				Expect(cacheErr).To(BeNil())
				Expect(cacheResp).ToNot(BeNil())
				Expect(cacheResult).ToNot(BeNil())
				Expect(*cacheResult.Success).Should(BeTrue())
			})
			It(`cache purge by cache tags`, func() {
				shouldSkipTest()
				cacheOpt := service.NewPurgeByCacheTagsOptions()
				cacheOpt.SetTags([]string{"some-tags"})
				cacheResult, cacheResp, cacheErr := service.PurgeByCacheTags(cacheOpt)
				Expect(cacheErr).To(BeNil())
				Expect(cacheResp).ToNot(BeNil())
				Expect(cacheResult).ToNot(BeNil())
				Expect(*cacheResult.Success).Should(BeTrue())
			})
			It(`cache purge by hosts`, func() {
				shouldSkipTest()
				cacheOpt := service.NewPurgeByHostsOptions()
				cacheOpt.SetHosts([]string{"www.example-host.com"})
				cacheResult, cacheResp, cacheErr := service.PurgeByHosts(cacheOpt)
				Expect(cacheErr).To(BeNil())
				Expect(cacheResp).ToNot(BeNil())
				Expect(cacheResult).ToNot(BeNil())
				Expect(*cacheResult.Success).Should(BeTrue())
			})
			It(`update/get cache level setting`, func() {
				shouldSkipTest()
				cacheLevel := []string{
					UpdateCacheLevelOptions_Value_Simplified,
					UpdateCacheLevelOptions_Value_Aggressive,
					UpdateCacheLevelOptions_Value_Basic,
				}
				cacheOpt := service.NewUpdateCacheLevelOptions()
				for _, level := range cacheLevel {
					cacheOpt.SetValue(level)
					cacheResult, cacheResp, cacheErr := service.UpdateCacheLevel(cacheOpt)
					Expect(cacheErr).To(BeNil())
					Expect(cacheResp).ToNot(BeNil())
					Expect(cacheResult).ToNot(BeNil())
					Expect(*cacheResult.Success).Should(BeTrue())

					getOpt := service.NewGetCacheLevelOptions()
					getResult, getResp, getErr := service.GetCacheLevel(getOpt)
					Expect(getErr).To(BeNil())
					Expect(getResp).ToNot(BeNil())
					Expect(getResult).ToNot(BeNil())
					Expect(*getResult.Success).Should(BeTrue())
					Expect(*getResult.Result.Value).Should(BeEquivalentTo(level))
				}
			})
			It(`update/get browser cache ttl setting`, func() {
				shouldSkipTest()
				ttl := []int64{0, 31536000, 14400}
				cacheOpt := service.NewUpdateBrowserCacheTtlOptions()
				for _, value := range ttl {
					cacheOpt.SetValue(value)
					cacheResult, cacheResp, cacheErr := service.UpdateBrowserCacheTTL(cacheOpt)
					Expect(cacheErr).To(BeNil())
					Expect(cacheResp).ToNot(BeNil())
					Expect(cacheResult).ToNot(BeNil())
					Expect(*cacheResult.Success).Should(BeTrue())

					getOpt := service.NewGetBrowserCacheTtlOptions()
					getResult, getResp, getErr := service.GetBrowserCacheTTL(getOpt)
					Expect(getErr).To(BeNil())
					Expect(getResp).ToNot(BeNil())
					Expect(getResult).ToNot(BeNil())
					Expect(*getResult.Success).Should(BeTrue())
					Expect(*getResult.Result.Value).Should(BeEquivalentTo(value))
				}
			})
			It(`update/get development mode setting`, func() {
				shouldSkipTest()
				cacheOpt := service.NewUpdateDevelopmentModeOptions()
				cacheOpt.SetValue(UpdateDevelopmentModeOptions_Value_On)
				cacheResult, cacheResp, cacheErr := service.UpdateDevelopmentMode(cacheOpt)
				Expect(cacheErr).To(BeNil())
				Expect(cacheResp).ToNot(BeNil())
				Expect(cacheResult).ToNot(BeNil())
				Expect(*cacheResult.Success).Should(BeTrue())

				getOpt := service.NewGetDevelopmentModeOptions()
				getResult, getResp, getErr := service.GetDevelopmentMode(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())
				Expect(*getResult.Result.Value).Should(BeEquivalentTo(UpdateDevelopmentModeOptions_Value_On))

				cacheOpt.SetValue(UpdateDevelopmentModeOptions_Value_Off)
				cacheResult, cacheResp, cacheErr = service.UpdateDevelopmentMode(cacheOpt)
				Expect(cacheErr).To(BeNil())
				Expect(cacheResp).ToNot(BeNil())
				Expect(cacheResult).ToNot(BeNil())
				Expect(*cacheResult.Success).Should(BeTrue())

				getResult, getResp, getErr = service.GetDevelopmentMode(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())
				Expect(*getResult.Result.Value).Should(BeEquivalentTo(UpdateDevelopmentModeOptions_Value_Off))
			})
			It(`update/get string sort setting`, func() {
				shouldSkipTest()
				cacheOpt := service.NewUpdateQueryStringSortOptions()
				cacheOpt.SetValue(UpdateQueryStringSortOptions_Value_On)
				cacheResult, cacheResp, cacheErr := service.UpdateQueryStringSort(cacheOpt)
				Expect(cacheErr).To(BeNil())
				Expect(cacheResp).ToNot(BeNil())
				Expect(cacheResult).ToNot(BeNil())
				Expect(*cacheResult.Success).Should(BeTrue())

				getOpt := service.NewGetQueryStringSortOptions()
				getResult, getResp, getErr := service.GetQueryStringSort(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())
				Expect(*getResult.Result.Value).Should(BeEquivalentTo(UpdateQueryStringSortOptions_Value_On))

				cacheOpt.SetValue(UpdateDevelopmentModeOptions_Value_Off)
				cacheResult, cacheResp, cacheErr = service.UpdateQueryStringSort(cacheOpt)
				Expect(cacheErr).To(BeNil())
				Expect(cacheResp).ToNot(BeNil())
				Expect(cacheResult).ToNot(BeNil())
				Expect(*cacheResult.Success).Should(BeTrue())

				getResult, getResp, getErr = service.GetQueryStringSort(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())
				Expect(*getResult.Result.Value).Should(BeEquivalentTo(UpdateQueryStringSortOptions_Value_Off))
			})
			It(`purge all test`, func() {
				shouldSkipTest()
				cacheOpt := service.NewPurgeAllOptions()
				cacheResult, cacheResp, cacheErr := service.PurgeAll(cacheOpt)
				Expect(cacheErr).To(BeNil())
				Expect(cacheResp).ToNot(BeNil())
				Expect(cacheResult).ToNot(BeNil())
				Expect(*cacheResult.Success).Should(BeTrue())
			})
		})
	})
})
