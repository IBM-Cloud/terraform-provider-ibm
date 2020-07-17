/**
 * (C) Copyright IBM Corp. 2020.
 */

package zonessettingsv1_test

import (
	"fmt"
	"os"
	"strings"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.ibm.com/ibmcloud/networking-go-sdk/zonessettingsv1"
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
	url := os.Getenv("URL_MATCH")
	globalOptions := &ZonesSettingsV1Options{
		ServiceName:    "cis_services",
		URL:            serviceURL,
		Authenticator:  authenticator,
		Crn:            &crn,
		ZoneIdentifier: &zone_id,
	}

	service, serviceErr := NewZonesSettingsV1(globalOptions)
	if serviceErr != nil {
		fmt.Println(serviceErr)
	}
	Describe(`zoneratelimitsv1_test`, func() {
		Context(`zoneratelimitsv1_test`, func() {
			It(`DNSSEC setting test`, func() {
				shouldSkipTest()
				getOpt := service.NewGetZoneDnssecOptions()
				getResult, getResp, getErr := service.GetZoneDnssec(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				updateOpt := service.NewUpdateZoneDnssecOptions()
				updateOpt.SetStatus(UpdateZoneDnssecOptions_Status_Active)
				if *getResult.Result.Status != UpdateZoneDnssecOptions_Status_Disabled {
					updateOpt.SetStatus(UpdateZoneDnssecOptions_Status_Disabled)
				}
				updateResult, updateResp, updateErr := service.UpdateZoneDnssec(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())
			})
			It(`CNAME Flattening setting test`, func() {
				shouldSkipTest()
				getOpt := service.NewGetZoneCnameFlatteningOptions()
				getResult, getResp, getErr := service.GetZoneCnameFlattening(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				updateOpt := service.NewUpdateZoneCnameFlatteningOptions()
				updateOpt.SetValue(UpdateZoneCnameFlatteningOptions_Value_FlattenAll)
				if *getResult.Result.Value != UpdateZoneCnameFlatteningOptions_Value_FlattenAll {
					updateOpt.SetValue(UpdateZoneCnameFlatteningOptions_Value_FlattenAtRoot)
				}
				updateResult, updateResp, updateErr := service.UpdateZoneCnameFlattening(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())
			})
			It(`opportunistic encryption setting test`, func() {
				shouldSkipTest()
				getOpt := service.NewGetOpportunisticEncryptionOptions()
				getResult, getResp, getErr := service.GetOpportunisticEncryption(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				updateOpt := service.NewUpdateOpportunisticEncryptionOptions()
				updateOpt.SetValue(UpdateOpportunisticEncryptionOptions_Value_Off)
				if *getResult.Result.Value != UpdateOpportunisticEncryptionOptions_Value_Off {
					updateOpt.SetValue(UpdateOpportunisticEncryptionOptions_Value_On)
				}
				updateResult, updateResp, updateErr := service.UpdateOpportunisticEncryption(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())
			})
			It(`challenge ttl setting test`, func() {
				shouldSkipTest()
				getOpt := service.NewGetChallengeTtlOptions()
				getResult, getResp, getErr := service.GetChallengeTTL(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				values := []int64{900, 1800, 2700, 3600, 7200, 10800, 14400, 28800, 57600, 86400, 604800, 2592000, 31536000, 300}
				updateOpt := service.NewUpdateChallengeTtlOptions()
				for _, value := range values {
					updateOpt.SetValue(value)
					updateResult, updateResp, updateErr := service.UpdateChallengeTTL(updateOpt)
					Expect(updateErr).To(BeNil())
					Expect(updateResp).ToNot(BeNil())
					Expect(updateResult).ToNot(BeNil())
					Expect(*updateResult.Success).Should(BeTrue())
				}
			})
			It(`automatic https rewrites setting test`, func() {
				shouldSkipTest()
				getOpt := service.NewGetAutomaticHttpsRewritesOptions()
				getResult, getResp, getErr := service.GetAutomaticHttpsRewrites(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				updateOpt := service.NewUpdateAutomaticHttpsRewritesOptions()
				updateOpt.SetValue(UpdateAutomaticHttpsRewritesOptions_Value_Off)
				if *getResult.Result.Value != UpdateAutomaticHttpsRewritesOptions_Value_Off {
					updateOpt.SetValue(UpdateAutomaticHttpsRewritesOptions_Value_On)
				}
				updateResult, updateResp, updateErr := service.UpdateAutomaticHttpsRewrites(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())
			})
			It(` true client IP setting test`, func() {
				shouldSkipTest()
				getOpt := service.NewGetTrueClientIpOptions()
				getResult, getResp, getErr := service.GetTrueClientIp(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				updateOpt := service.NewUpdateTrueClientIpOptions()
				updateOpt.SetValue(UpdateTrueClientIpOptions_Value_Off)
				if *getResult.Result.Value != UpdateTrueClientIpOptions_Value_Off {
					updateOpt.SetValue(UpdateTrueClientIpOptions_Value_On)
				}
				updateResult, updateResp, updateErr := service.UpdateTrueClientIp(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())
			})
			It(`always use https setting test`, func() {
				shouldSkipTest()
				getOpt := service.NewGetAlwaysUseHttpsOptions()
				getResult, getResp, getErr := service.GetAlwaysUseHttps(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				updateOpt := service.NewUpdateAlwaysUseHttpsOptions()
				updateOpt.SetValue(UpdateAlwaysUseHttpsOptions_Value_Off)
				if *getResult.Result.Value != UpdateAlwaysUseHttpsOptions_Value_Off {
					updateOpt.SetValue(UpdateAlwaysUseHttpsOptions_Value_On)
				}
				updateResult, updateResp, updateErr := service.UpdateAlwaysUseHttps(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())
			})
			It(`image load optimization setting test`, func() {
				shouldSkipTest()
				getOpt := service.NewGetImageLoadOptimizationOptions()
				getResult, getResp, getErr := service.GetImageLoadOptimization(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				updateOpt := service.NewUpdateImageLoadOptimizationOptions()
				updateOpt.SetValue(UpdateImageLoadOptimizationOptions_Value_Off)
				if *getResult.Result.Value != UpdateImageLoadOptimizationOptions_Value_Off {
					updateOpt.SetValue(UpdateImageLoadOptimizationOptions_Value_On)
				}
				updateResult, updateResp, updateErr := service.UpdateImageLoadOptimization(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())
			})
			It(`image size optimization setting test`, func() {
				shouldSkipTest()
				getOpt := service.NewGetImageSizeOptimizationOptions()
				getResult, getResp, getErr := service.GetImageSizeOptimization(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				values := []string{
					UpdateImageSizeOptimizationOptions_Value_Lossless,
					UpdateImageSizeOptimizationOptions_Value_Lossy,
					UpdateImageSizeOptimizationOptions_Value_Off,
				}
				updateOpt := service.NewUpdateImageSizeOptimizationOptions()
				for _, value := range values {
					updateOpt.SetValue(value)
					updateResult, updateResp, updateErr := service.UpdateImageSizeOptimization(updateOpt)
					Expect(updateErr).To(BeNil())
					Expect(updateResp).ToNot(BeNil())
					Expect(updateResult).ToNot(BeNil())
					Expect(*updateResult.Success).Should(BeTrue())
				}
			})
			It(`script load optimization setting test`, func() {
				shouldSkipTest()
				getOpt := service.NewGetScriptLoadOptimizationOptions()
				getResult, getResp, getErr := service.GetScriptLoadOptimization(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				updateOpt := service.NewUpdateScriptLoadOptimizationOptions()
				updateOpt.SetValue(UpdateScriptLoadOptimizationOptions_Value_Off)
				if *getResult.Result.Value != UpdateScriptLoadOptimizationOptions_Value_Off {
					updateOpt.SetValue(UpdateScriptLoadOptimizationOptions_Value_On)
				}
				updateResult, updateResp, updateErr := service.UpdateScriptLoadOptimization(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())
			})
			It(`minify setting test`, func() {
				shouldSkipTest()
				getOpt := service.NewGetMinifyOptions()
				getResult, getResp, getErr := service.GetMinify(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				minifyOpt, err := service.NewMinifySettingValue(MinifySettingValue_Css_Off, MinifySettingValue_HTML_Off, MinifySettingValue_Js_Off)
				Expect(err).To(BeNil())

				updateOpt := service.NewUpdateMinifyOptions()
				updateOpt.SetValue(minifyOpt)

				if *getResult.Result.Value.Css != MinifySettingValue_Css_Off {
					minifyOpt, err := service.NewMinifySettingValue(MinifySettingValue_Css_On, MinifySettingValue_HTML_On, MinifySettingValue_Js_On)
					Expect(err).To(BeNil())
					updateOpt.SetValue(minifyOpt)
				}
				updateResult, updateResp, updateErr := service.UpdateMinify(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())
			})
			It(`IP geolocation setting test`, func() {
				shouldSkipTest()
				getOpt := service.NewGetIpGeolocationOptions()
				getResult, getResp, getErr := service.GetIpGeolocation(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				updateOpt := service.NewUpdateIpGeolocationOptions()
				updateOpt.SetValue(UpdateIpGeolocationOptions_Value_Off)
				if *getResult.Result.Value != UpdateIpGeolocationOptions_Value_Off {
					updateOpt.SetValue(UpdateIpGeolocationOptions_Value_On)
				}
				updateResult, updateResp, updateErr := service.UpdateIpGeolocation(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())
			})
			It(`server side exclude setting test`, func() {
				shouldSkipTest()
				getOpt := service.NewGetServerSideExcludeOptions()
				getResult, getResp, getErr := service.GetServerSideExclude(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				updateOpt := service.NewUpdateServerSideExcludeOptions()
				updateOpt.SetValue(UpdateServerSideExcludeOptions_Value_Off)
				if *getResult.Result.Value != UpdateServerSideExcludeOptions_Value_Off {
					updateOpt.SetValue(UpdateServerSideExcludeOptions_Value_On)
				}
				updateResult, updateResp, updateErr := service.UpdateServerSideExclude(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())
			})

			It(`HTTP strict transport security setting test`, func() {
				shouldSkipTest()
				getOpt := service.NewGetSecurityHeaderOptions()
				getResult, getResp, getErr := service.GetSecurityHeader(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				valueOpt, err := service.NewSecurityHeaderSettingValueStrictTransportSecurity(true, 3600, true, true)
				Expect(err).To(BeNil())
				securityOpt, err := service.NewSecurityHeaderSettingValue(valueOpt)
				Expect(err).To(BeNil())

				updateOpt := service.NewUpdateSecurityHeaderOptions()
				updateOpt.SetValue(securityOpt)
				if *getResult.Result.Value.StrictTransportSecurity.Enabled != true {
					valueOpt.Enabled = core.BoolPtr(false)
					valueOpt.IncludeSubdomains = core.BoolPtr(false)
					valueOpt.MaxAge = core.Int64Ptr(7200)
					valueOpt.Nosniff = core.BoolPtr(false)
					securityOpt.StrictTransportSecurity = valueOpt
					updateOpt.SetValue(securityOpt)
				}
				updateResult, updateResp, updateErr := service.UpdateSecurityHeader(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())
			})
			It(`mobile redirect setting test`, func() {
				shouldSkipTest()
				getOpt := service.NewGetMobileRedirectOptions()
				getResult, getResp, getErr := service.GetMobileRedirect(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				value := strings.Split(url, ".")
				mobileOpt, err := service.NewMobileRedirecSettingValue(MobileRedirecSettingValue_Status_Off, value[0], false)
				Expect(err).To(BeNil())

				updateOpt := service.NewUpdateMobileRedirectOptions()
				updateOpt.SetValue(mobileOpt)
				if *getResult.Result.Value.StripURI != false {
					mobileOpt.Status = core.StringPtr(MobileRedirecSettingValue_Status_On)
					mobileOpt.StripURI = core.BoolPtr(false)
					updateOpt.SetValue(mobileOpt)
				}
				updateResult, updateResp, updateErr := service.UpdateMobileRedirect(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())
			})
			It(`prefetch URLs from header setting test`, func() {
				shouldSkipTest()
				getOpt := service.NewGetPrefetchPreloadOptions()
				getResult, getResp, getErr := service.GetPrefetchPreload(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				updateOpt := service.NewUpdatePrefetchPreloadOptions()
				updateOpt.SetValue(UpdatePrefetchPreloadOptions_Value_Off)
				if *getResult.Result.Value != UpdatePrefetchPreloadOptions_Value_Off {
					updateOpt.SetValue(UpdatePrefetchPreloadOptions_Value_On)
				}
				updateResult, updateResp, updateErr := service.UpdatePrefetchPreload(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())
			})
			It(`http/2 setting test`, func() {
				shouldSkipTest()
				getOpt := service.NewGetHttp2Options()
				getResult, getResp, getErr := service.GetHttp2(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				updateOpt := service.NewUpdateHttp2Options()
				updateOpt.SetValue(UpdateHttp2Options_Value_Off)
				if *getResult.Result.Value != UpdateHttp2Options_Value_Off {
					updateOpt.SetValue(UpdateHttp2Options_Value_On)
				}
				updateResult, updateResp, updateErr := service.UpdateHttp2(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())
			})
			It(`ipv6 compatibility setting test`, func() {
				shouldSkipTest()
				getOpt := service.NewGetIpv6Options()
				getResult, getResp, getErr := service.GetIpv6(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				updateOpt := service.NewUpdateIpv6Options()
				updateOpt.SetValue(UpdateIpv6Options_Value_Off)
				if *getResult.Result.Value != UpdateIpv6Options_Value_Off {
					updateOpt.SetValue(UpdateIpv6Options_Value_On)
				}
				updateResult, updateResp, updateErr := service.UpdateIpv6(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())
			})
			It(`pseudo IPv4 setting test`, func() {
				shouldSkipTest()
				getOpt := service.NewGetPseudoIpv4Options()
				getResult, getResp, getErr := service.GetPseudoIpv4(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				values := []string{
					UpdatePseudoIpv4Options_Value_AddHeader,
					UpdatePseudoIpv4Options_Value_OverwriteHeader,
					UpdatePseudoIpv4Options_Value_Off,
				}
				updateOpt := service.NewUpdatePseudoIpv4Options()
				for _, value := range values {
					updateOpt.SetValue(value)
					updateResult, updateResp, updateErr := service.UpdatePseudoIpv4(updateOpt)
					Expect(updateErr).To(BeNil())
					Expect(updateResp).ToNot(BeNil())
					Expect(updateResult).ToNot(BeNil())
					Expect(*updateResult.Success).Should(BeTrue())
				}
			})
			It(`response buffering setting test`, func() {
				shouldSkipTest()
				getOpt := service.NewGetResponseBufferingOptions()
				getResult, getResp, getErr := service.GetResponseBuffering(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				updateOpt := service.NewUpdateResponseBufferingOptions()
				updateOpt.SetValue(UpdateResponseBufferingOptions_Value_Off)
				if *getResult.Result.Value != UpdateResponseBufferingOptions_Value_Off {
					updateOpt.SetValue(UpdateResponseBufferingOptions_Value_On)
				}
				updateResult, updateResp, updateErr := service.UpdateResponseBuffering(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())
			})
			It(`hotlink protection setting test`, func() {
				shouldSkipTest()
				getOpt := service.NewGetHotlinkProtectionOptions()
				getResult, getResp, getErr := service.GetHotlinkProtection(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				updateOpt := service.NewUpdateHotlinkProtectionOptions()
				updateOpt.SetValue(UpdateHotlinkProtectionOptions_Value_Off)
				if *getResult.Result.Value != UpdateHotlinkProtectionOptions_Value_Off {
					updateOpt.SetValue(UpdateHotlinkProtectionOptions_Value_On)
				}
				updateResult, updateResp, updateErr := service.UpdateHotlinkProtection(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())
			})
			It(`maximum upload size setting test`, func() {
				shouldSkipTest()
				getOpt := service.NewGetMaxUploadOptions()
				getResult, getResp, getErr := service.GetMaxUpload(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				values := []int64{125, 150, 175, 200, 225, 250, 275, 300, 325, 350, 375, 400, 425, 450, 475, 500, 100}
				updateOpt := service.NewUpdateMaxUploadOptions()
				for _, value := range values {
					updateOpt.SetValue(value)
					updateResult, updateResp, updateErr := service.UpdateMaxUpload(updateOpt)
					Expect(updateErr).To(BeNil())
					Expect(updateResp).ToNot(BeNil())
					Expect(updateResult).ToNot(BeNil())
					Expect(*updateResult.Success).Should(BeTrue())
				}
			})
			It(`Min TLS version setting test`, func() {
				shouldSkipTest()
				getOpt := service.NewGetMinTlsVersionOptions()
				getResult, getResp, getErr := service.GetMinTlsVersion(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				values := []string{"1.0", "1.1", "1.2"}
				updateOpt := service.NewUpdateMinTlsVersionOptions()
				for _, value := range values {
					updateOpt.SetValue(value)
					updateResult, updateResp, updateErr := service.UpdateMinTlsVersion(updateOpt)
					Expect(updateErr).To(BeNil())
					Expect(updateResp).ToNot(BeNil())
					Expect(updateResult).ToNot(BeNil())
					Expect(*updateResult.Success).Should(BeTrue())
				}
			})
			It(`TLS Client Auth setting test`, func() {
				shouldSkipTest()
				getOpt := service.NewGetTlsClientAuthOptions()
				getResult, getResp, getErr := service.GetTlsClientAuth(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				updateOpt := service.NewUpdateTlsClientAuthOptions()
				updateOpt.SetValue(UpdateTlsClientAuthOptions_Value_Off)
				if *getResult.Result.Value != UpdateTlsClientAuthOptions_Value_Off {
					updateOpt.SetValue(UpdateTlsClientAuthOptions_Value_On)
				}
				updateResult, updateResp, updateErr := service.UpdateTlsClientAuth(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())
			})
			It(`Browser check setting test`, func() {
				shouldSkipTest()
				getOpt := service.NewGetBrowserCheckOptions()
				getResult, getResp, getErr := service.GetBrowserCheck(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				updateOpt := service.NewUpdateBrowserCheckOptions()
				updateOpt.SetValue(UpdateBrowserCheckOptions_Value_Off)
				if *getResult.Result.Value != UpdateBrowserCheckOptions_Value_Off {
					updateOpt.SetValue(UpdateBrowserCheckOptions_Value_On)
				}
				updateResult, updateResp, updateErr := service.UpdateBrowserCheck(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())
			})
			It(`Enable error pages setting test`, func() {
				shouldSkipTest()
				getOpt := service.NewGetEnableErrorPagesOnOptions()
				getResult, getResp, getErr := service.GetEnableErrorPagesOn(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				updateOpt := service.NewUpdateEnableErrorPagesOnOptions()
				updateOpt.SetValue(UpdateEnableErrorPagesOnOptions_Value_Off)
				if *getResult.Result.Value != UpdateEnableErrorPagesOnOptions_Value_Off {
					updateOpt.SetValue(UpdateEnableErrorPagesOnOptions_Value_On)
				}
				updateResult, updateResp, updateErr := service.UpdateEnableErrorPagesOn(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())
			})
			It(`ciphers setting test`, func() {
				shouldSkipTest()
				getOpt := service.NewGetCiphersOptions()
				getResult, getResp, getErr := service.GetCiphers(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				values := []string{
					UpdateCiphersOptions_Value_Aes128GcmSha256,
					UpdateCiphersOptions_Value_Aes128Sha,
					UpdateCiphersOptions_Value_Aes128Sha256,
					UpdateCiphersOptions_Value_Aes256GcmSha384,
					UpdateCiphersOptions_Value_Aes256Sha,
					UpdateCiphersOptions_Value_Aes256Sha256,
					UpdateCiphersOptions_Value_DesCbc3Sha,
					UpdateCiphersOptions_Value_EcdheEcdsaAes128GcmSha256,
					UpdateCiphersOptions_Value_EcdheEcdsaAes128Sha,
					UpdateCiphersOptions_Value_EcdheEcdsaAes128Sha256,
					UpdateCiphersOptions_Value_EcdheEcdsaAes256GcmSha384,
					UpdateCiphersOptions_Value_EcdheEcdsaAes256Sha384,
					UpdateCiphersOptions_Value_EcdheEcdsaChacha20Poly1305,
					UpdateCiphersOptions_Value_EcdheRsaAes128GcmSha256,
					UpdateCiphersOptions_Value_EcdheRsaAes128Sha,
					UpdateCiphersOptions_Value_EcdheRsaAes128Sha256,
					UpdateCiphersOptions_Value_EcdheRsaAes256GcmSha384,
					UpdateCiphersOptions_Value_EcdheRsaAes256Sha,
					UpdateCiphersOptions_Value_EcdheRsaAes256Sha384,
					UpdateCiphersOptions_Value_EcdheRsaChacha20Poly1305,
				}
				updateOpt := service.NewUpdateCiphersOptions()
				for _, value := range values {
					updateOpt.SetValue([]string{value})
					updateResult, updateResp, updateErr := service.UpdateCiphers(updateOpt)
					Expect(updateErr).To(BeNil())
					Expect(updateResp).ToNot(BeNil())
					Expect(updateResult).ToNot(BeNil())
					Expect(*updateResult.Success).Should(BeTrue())
				}
			})
		})
	})
})
