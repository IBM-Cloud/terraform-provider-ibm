/*
 * (C) Copyright IBM Corp. 2020.
 */

package zoneratelimitsv1_test

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/core"
	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.ibm.com/ibmcloud/networking-go-sdk/zoneratelimitsv1"
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
	globalOptions := &ZoneRateLimitsV1Options{
		ServiceName:    "cis_services",
		URL:            serviceURL,
		Authenticator:  authenticator,
		Crn:            &crn,
		ZoneIdentifier: &zone_id,
	}

	service, serviceErr := NewZoneRateLimitsV1(globalOptions)
	if serviceErr != nil {
		fmt.Println(serviceErr)
	}
	Describe(`zoneratelimitsv1_test`, func() {
		Context(`zoneratelimitsv1_test`, func() {
			BeforeEach(func() {
				shouldSkipTest()
				listOpt := service.NewListAllZoneRateLimitsOptions()
				listResult, listResp, listErr := service.ListAllZoneRateLimits(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())

				for _, rate := range listResult.Result {
					// delete range app
					delOpt := service.NewDeleteZoneRateLimitOptions(*rate.ID)
					delResult, delResp, delErr := service.DeleteZoneRateLimit(delOpt)
					Expect(delErr).To(BeNil())
					Expect(delResp).ToNot(BeNil())
					Expect(delResult).ToNot(BeNil())
					Expect(*delResult.Success).Should(BeTrue())
				}
			})
			AfterEach(func() {
				shouldSkipTest()
				listOpt := service.NewListAllZoneRateLimitsOptions()
				listResult, listResp, listErr := service.ListAllZoneRateLimits(listOpt)
				Expect(listErr).To(BeNil())
				Expect(listResp).ToNot(BeNil())
				Expect(listResult).ToNot(BeNil())
				Expect(*listResult.Success).Should(BeTrue())

				for _, rate := range listResult.Result {
					// delete range app
					delOpt := service.NewDeleteZoneRateLimitOptions(*rate.ID)
					delResult, delResp, delErr := service.DeleteZoneRateLimit(delOpt)
					Expect(delErr).To(BeNil())
					Expect(delResp).ToNot(BeNil())
					Expect(delResult).ToNot(BeNil())
					Expect(*delResult.Success).Should(BeTrue())
				}
			})
			It(`zone rate limit modes test [smulate & Ban]`, func() {
				shouldSkipTest()
				url := "*.example.org/path*"
				createModes := []string{
					RatelimitInputAction_Mode_Simulate,
					RatelimitInputAction_Mode_Ban,
				}
				updateModes := []string{
					RatelimitInputAction_Mode_Ban,
					RatelimitInputAction_Mode_Simulate,
				}
				for i, mode := range createModes {

					// create rate limit

					actionOpt, err := service.NewRatelimitInputAction(mode)
					Expect(err).To(BeNil())
					actionOpt.Timeout = core.Int64Ptr(60)
					actionOpt.Response = &RatelimitInputActionResponse{
						ContentType: core.StringPtr(RatelimitInputActionResponse_ContentType_TextPlain),
						Body:        core.StringPtr("This request has been rate-limited."),
					}

					bypassOpt, err := service.NewRatelimitInputBypassItem(RatelimitInputBypassItem_Name_URL, "api.example.com/*")
					Expect(err).To(BeNil())

					correlateOpt, err := service.NewRatelimitInputCorrelate(RatelimitInputCorrelate_By_Nat)
					Expect(err).To(BeNil())

					matchReqOpt, err := service.NewRatelimitInputMatchRequest(url)
					Expect(err).To(BeNil())
					matchReqOpt.Methods = []string{RatelimitInputMatchRequest_Methods_All}
					matchReqOpt.Schemes = []string{RatelimitInputMatchRequest_Schemes_All}

					matchOpt := RatelimitInputMatch{
						Request: matchReqOpt,
					}

					createOpt := service.NewCreateZoneRateLimitsOptions()
					createOpt.SetAction(actionOpt)
					createOpt.SetBypass([]RatelimitInputBypassItem{*bypassOpt})
					createOpt.SetCorrelate(correlateOpt)
					createOpt.SetDisabled(false)
					createOpt.SetMatch(&matchOpt)
					createOpt.SetPeriod(2)
					createOpt.SetThreshold(40)

					createResult, createResp, createErr := service.CreateZoneRateLimits(createOpt)
					Expect(createErr).To(BeNil())
					Expect(createResp).ToNot(BeNil())
					Expect(createResult).ToNot(BeNil())
					Expect(*createResult.Success).Should(BeTrue())

					// update rate limit
					actionOpt.Mode = &updateModes[i]

					updateOpt := service.NewUpdateRateLimitOptions(*createResult.Result.ID)
					updateOpt.SetAction(actionOpt)
					updateOpt.SetBypass([]RatelimitInputBypassItem{*bypassOpt})
					updateOpt.SetCorrelate(correlateOpt)
					updateOpt.SetDisabled(false)
					updateOpt.SetMatch(&matchOpt)
					updateOpt.SetPeriod(2)
					updateOpt.SetThreshold(40)

					updateResult, updateResp, updateErr := service.UpdateRateLimit(updateOpt)
					Expect(updateErr).To(BeNil())
					Expect(updateResp).ToNot(BeNil())
					Expect(updateResult).ToNot(BeNil())
					Expect(*updateResult.Success).Should(BeTrue())

					getOpt := service.NewGetRateLimitOptions(*createResult.Result.ID)
					getResult, getResp, getErr := service.GetRateLimit(getOpt)
					Expect(getErr).To(BeNil())
					Expect(getResp).ToNot(BeNil())
					Expect(getResult).ToNot(BeNil())
					Expect(*getResult.Success).Should(BeTrue())

					deleteOpt := service.NewDeleteZoneRateLimitOptions(*createResult.Result.ID)
					deleteResult, deleteResp, deleteErr := service.DeleteZoneRateLimit(deleteOpt)
					Expect(deleteErr).To(BeNil())
					Expect(deleteResp).ToNot(BeNil())
					Expect(deleteResult).ToNot(BeNil())
					Expect(*deleteResult.Success).Should(BeTrue())
				}
			})
			It(`zone rate limit modes test [challenge & jschallenge]`, func() {
				shouldSkipTest()
				url := "*.example.org/path*"
				createModes := []string{
					RatelimitInputAction_Mode_Challenge,
					RatelimitInputAction_Mode_JsChallenge,
				}
				updateModes := []string{
					RatelimitInputAction_Mode_JsChallenge,
					RatelimitInputAction_Mode_Challenge,
				}
				for i, mode := range createModes {

					// create rate limit

					actionOpt, err := service.NewRatelimitInputAction(mode)
					Expect(err).To(BeNil())

					bypassOpt, err := service.NewRatelimitInputBypassItem(RatelimitInputBypassItem_Name_URL, "api.example.com/*")
					Expect(err).To(BeNil())

					correlateOpt, err := service.NewRatelimitInputCorrelate(RatelimitInputCorrelate_By_Nat)
					Expect(err).To(BeNil())

					matchReqOpt, err := service.NewRatelimitInputMatchRequest(url)
					Expect(err).To(BeNil())
					matchReqOpt.Methods = []string{RatelimitInputMatchRequest_Methods_All}
					matchReqOpt.Schemes = []string{RatelimitInputMatchRequest_Schemes_All}

					matchOpt := RatelimitInputMatch{
						Request: matchReqOpt,
					}

					createOpt := service.NewCreateZoneRateLimitsOptions()
					createOpt.SetAction(actionOpt)
					createOpt.SetBypass([]RatelimitInputBypassItem{*bypassOpt})
					createOpt.SetCorrelate(correlateOpt)
					createOpt.SetDisabled(false)
					createOpt.SetMatch(&matchOpt)
					createOpt.SetPeriod(2)
					createOpt.SetThreshold(40)

					createResult, createResp, createErr := service.CreateZoneRateLimits(createOpt)
					Expect(createErr).To(BeNil())
					Expect(createResp).ToNot(BeNil())
					Expect(createResult).ToNot(BeNil())
					Expect(*createResult.Success).Should(BeTrue())

					// update rate limit
					actionOpt.Mode = &updateModes[i]

					updateOpt := service.NewUpdateRateLimitOptions(*createResult.Result.ID)
					updateOpt.SetAction(actionOpt)
					updateOpt.SetBypass([]RatelimitInputBypassItem{*bypassOpt})
					updateOpt.SetCorrelate(correlateOpt)
					updateOpt.SetDisabled(false)
					updateOpt.SetMatch(&matchOpt)
					updateOpt.SetPeriod(2)
					updateOpt.SetThreshold(40)

					updateResult, updateResp, updateErr := service.UpdateRateLimit(updateOpt)
					Expect(updateErr).To(BeNil())
					Expect(updateResp).ToNot(BeNil())
					Expect(updateResult).ToNot(BeNil())
					Expect(*updateResult.Success).Should(BeTrue())

					getOpt := service.NewGetRateLimitOptions(*createResult.Result.ID)
					getResult, getResp, getErr := service.GetRateLimit(getOpt)
					Expect(getErr).To(BeNil())
					Expect(getResp).ToNot(BeNil())
					Expect(getResult).ToNot(BeNil())
					Expect(*getResult.Success).Should(BeTrue())

					deleteOpt := service.NewDeleteZoneRateLimitOptions(*createResult.Result.ID)
					deleteResult, deleteResp, deleteErr := service.DeleteZoneRateLimit(deleteOpt)
					Expect(deleteErr).To(BeNil())
					Expect(deleteResp).ToNot(BeNil())
					Expect(deleteResult).ToNot(BeNil())
					Expect(*deleteResult.Success).Should(BeTrue())
				}
			})
			It(`zone rate limit modes test [smulate & Ban] with action response content`, func() {
				shouldSkipTest()
				url := "*.example.org/path*"
				createModes := []string{
					RatelimitInputAction_Mode_Simulate,
					RatelimitInputAction_Mode_Ban,
				}
				updateModes := []string{
					RatelimitInputAction_Mode_Ban,
					RatelimitInputAction_Mode_Simulate,
				}
				createResponse1 := &RatelimitInputActionResponse{
					ContentType: core.StringPtr(RatelimitInputActionResponse_ContentType_TextPlain),
					Body:        core.StringPtr("This request has been rate-limited."),
				}
				createResponse2 := &RatelimitInputActionResponse{
					ContentType: core.StringPtr(RatelimitInputActionResponse_ContentType_ApplicationJSON),
					Body:        core.StringPtr("{\"name\": \"rate_limit\", \"msg\": \"This request has been rate-limited.\"}"),
				}

				updateResponse1 := &RatelimitInputActionResponse{
					ContentType: core.StringPtr(RatelimitInputActionResponse_ContentType_TextPlain),
					Body:        core.StringPtr("This request has been rate-limited."),
				}
				updateResponse2 := &RatelimitInputActionResponse{
					ContentType: core.StringPtr(RatelimitInputActionResponse_ContentType_ApplicationJSON),
					Body:        core.StringPtr("{\"name\": \"rate_limit\", \"msg\": \"This request has been rate-limited.\"}"),
				}

				responses := []RatelimitInputActionResponse{*createResponse1, *createResponse2}
				updateResponses := []RatelimitInputActionResponse{*updateResponse1, *updateResponse2}

				for i, mode := range createModes {

					for j, response := range responses {

						// create rate limit

						actionOpt, err := service.NewRatelimitInputAction(mode)
						Expect(err).To(BeNil())
						actionOpt.Timeout = core.Int64Ptr(60)
						actionOpt.Response = &response

						bypassOpt, err := service.NewRatelimitInputBypassItem(RatelimitInputBypassItem_Name_URL, "api.example.com/*")
						Expect(err).To(BeNil())

						correlateOpt, err := service.NewRatelimitInputCorrelate(RatelimitInputCorrelate_By_Nat)
						Expect(err).To(BeNil())

						matchReqOpt, err := service.NewRatelimitInputMatchRequest(url)
						Expect(err).To(BeNil())
						matchReqOpt.Methods = []string{RatelimitInputMatchRequest_Methods_All}
						matchReqOpt.Schemes = []string{RatelimitInputMatchRequest_Schemes_All}

						matchOpt := RatelimitInputMatch{
							Request: matchReqOpt,
						}

						createOpt := service.NewCreateZoneRateLimitsOptions()
						createOpt.SetAction(actionOpt)
						createOpt.SetBypass([]RatelimitInputBypassItem{*bypassOpt})
						createOpt.SetCorrelate(correlateOpt)
						createOpt.SetDisabled(false)
						createOpt.SetMatch(&matchOpt)
						createOpt.SetPeriod(2)
						createOpt.SetThreshold(40)

						createResult, createResp, createErr := service.CreateZoneRateLimits(createOpt)
						Expect(createErr).To(BeNil())
						Expect(createResp).ToNot(BeNil())
						Expect(createResult).ToNot(BeNil())
						Expect(*createResult.Success).Should(BeTrue())

						// update rate limit
						actionOpt.Mode = &updateModes[i]
						actionOpt.Timeout = core.Int64Ptr(60)
						actionOpt.Response = &updateResponses[j]

						updateOpt := service.NewUpdateRateLimitOptions(*createResult.Result.ID)
						updateOpt.SetAction(actionOpt)
						updateOpt.SetBypass([]RatelimitInputBypassItem{*bypassOpt})
						updateOpt.SetCorrelate(correlateOpt)
						updateOpt.SetDisabled(false)
						updateOpt.SetMatch(&matchOpt)
						updateOpt.SetPeriod(2)
						updateOpt.SetThreshold(40)

						updateResult, updateResp, updateErr := service.UpdateRateLimit(updateOpt)
						Expect(updateErr).To(BeNil())
						Expect(updateResp).ToNot(BeNil())
						Expect(updateResult).ToNot(BeNil())
						Expect(*updateResult.Success).Should(BeTrue())

						getOpt := service.NewGetRateLimitOptions(*createResult.Result.ID)
						getResult, getResp, getErr := service.GetRateLimit(getOpt)
						Expect(getErr).To(BeNil())
						Expect(getResp).ToNot(BeNil())
						Expect(getResult).ToNot(BeNil())
						Expect(*getResult.Success).Should(BeTrue())

						deleteOpt := service.NewDeleteZoneRateLimitOptions(*createResult.Result.ID)
						deleteResult, deleteResp, deleteErr := service.DeleteZoneRateLimit(deleteOpt)
						Expect(deleteErr).To(BeNil())
						Expect(deleteResp).ToNot(BeNil())
						Expect(deleteResult).ToNot(BeNil())
						Expect(*deleteResult.Success).Should(BeTrue())
					}
				}
			})
			It(`zone rate limit test with request content`, func() {
				shouldSkipTest()
				url := "*.example.org/path*"
				methods := []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "_ALL_"}
				schemas := []string{"HTTP", "HTTPS", "_ALL_"}
				updateMethods := []string{"POST", "PUT", "DELETE", "PATCH", "HEAD", "_ALL_", "GET"}
				updateSchemas := []string{"HTTPS", "_ALL_", "HTTP"}
				for i, method := range methods {
					for j, schema := range schemas {

						// create rate limit

						actionOpt, err := service.NewRatelimitInputAction(RatelimitInputAction_Mode_Simulate)
						Expect(err).To(BeNil())
						actionOpt.Timeout = core.Int64Ptr(60)
						actionOpt.Response = &RatelimitInputActionResponse{
							ContentType: core.StringPtr(RatelimitInputActionResponse_ContentType_TextPlain),
							Body:        core.StringPtr("This request has been rate-limited."),
						}

						bypassOpt, err := service.NewRatelimitInputBypassItem(RatelimitInputBypassItem_Name_URL, "api.example.com/*")
						Expect(err).To(BeNil())

						correlateOpt, err := service.NewRatelimitInputCorrelate(RatelimitInputCorrelate_By_Nat)
						Expect(err).To(BeNil())

						matchReqOpt, err := service.NewRatelimitInputMatchRequest(url)
						Expect(err).To(BeNil())
						matchReqOpt.Methods = []string{method}
						matchReqOpt.Schemes = []string{schema}

						matchOpt := RatelimitInputMatch{
							Request: matchReqOpt,
						}

						createOpt := service.NewCreateZoneRateLimitsOptions()
						createOpt.SetAction(actionOpt)
						createOpt.SetBypass([]RatelimitInputBypassItem{*bypassOpt})
						createOpt.SetCorrelate(correlateOpt)
						createOpt.SetDisabled(false)
						createOpt.SetMatch(&matchOpt)
						createOpt.SetPeriod(2)
						createOpt.SetThreshold(40)

						createResult, createResp, createErr := service.CreateZoneRateLimits(createOpt)
						Expect(createErr).To(BeNil())
						Expect(createResp).ToNot(BeNil())
						Expect(createResult).ToNot(BeNil())
						Expect(*createResult.Success).Should(BeTrue())

						// update rate limit
						matchReqOpt.Methods = []string{updateMethods[i]}
						matchReqOpt.Schemes = []string{updateSchemas[j]}
						updateOpt := service.NewUpdateRateLimitOptions(*createResult.Result.ID)
						updateOpt.SetAction(actionOpt)
						updateOpt.SetBypass([]RatelimitInputBypassItem{*bypassOpt})
						updateOpt.SetCorrelate(correlateOpt)
						updateOpt.SetDisabled(false)
						updateOpt.SetMatch(&matchOpt)
						updateOpt.SetPeriod(2)
						updateOpt.SetThreshold(40)

						updateResult, updateResp, updateErr := service.UpdateRateLimit(updateOpt)
						Expect(updateErr).To(BeNil())
						Expect(updateResp).ToNot(BeNil())
						Expect(updateResult).ToNot(BeNil())
						Expect(*updateResult.Success).Should(BeTrue())

						getOpt := service.NewGetRateLimitOptions(*createResult.Result.ID)
						getResult, getResp, getErr := service.GetRateLimit(getOpt)
						Expect(getErr).To(BeNil())
						Expect(getResp).ToNot(BeNil())
						Expect(getResult).ToNot(BeNil())
						Expect(*getResult.Success).Should(BeTrue())

						deleteOpt := service.NewDeleteZoneRateLimitOptions(*createResult.Result.ID)
						deleteResult, deleteResp, deleteErr := service.DeleteZoneRateLimit(deleteOpt)
						Expect(deleteErr).To(BeNil())
						Expect(deleteResp).ToNot(BeNil())
						Expect(deleteResult).ToNot(BeNil())
						Expect(*deleteResult.Success).Should(BeTrue())
					}
				}
			})
			It(`zone rate limit modes test with response match`, func() {
				shouldSkipTest()
				url := "*.example.org/path*"
				// create rate limit

				actionOpt, err := service.NewRatelimitInputAction(RatelimitInputAction_Mode_Simulate)
				Expect(err).To(BeNil())
				actionOpt.Timeout = core.Int64Ptr(60)
				actionOpt.Response = &RatelimitInputActionResponse{
					ContentType: core.StringPtr(RatelimitInputActionResponse_ContentType_TextPlain),
					Body:        core.StringPtr("This request has been rate-limited."),
				}

				bypassOpt, err := service.NewRatelimitInputBypassItem(RatelimitInputBypassItem_Name_URL, "api.example.com/*")
				Expect(err).To(BeNil())

				correlateOpt, err := service.NewRatelimitInputCorrelate(RatelimitInputCorrelate_By_Nat)
				Expect(err).To(BeNil())

				matchReqOpt, err := service.NewRatelimitInputMatchRequest(url)
				Expect(err).To(BeNil())
				matchReqOpt.Methods = []string{RatelimitInputMatchRequest_Methods_All}
				matchReqOpt.Schemes = []string{RatelimitInputMatchRequest_Schemes_All}

				respOpt, err := service.NewRatelimitInputMatchResponseHeadersItem("Cf-Cache-Status", RatelimitInputMatchResponseHeadersItem_Op_Eq, RatelimitInputMatchResponseHeadersItem_Value_Hit)
				Expect(err).To(BeNil())
				respItems := []RatelimitInputMatchResponseHeadersItem{*respOpt}

				matchRespOpt := &RatelimitInputMatchResponse{
					HeadersVar: respItems,
				}

				matchOpt := RatelimitInputMatch{
					Request:  matchReqOpt,
					Response: matchRespOpt,
				}

				createOpt := service.NewCreateZoneRateLimitsOptions()
				createOpt.SetAction(actionOpt)
				createOpt.SetBypass([]RatelimitInputBypassItem{*bypassOpt})
				createOpt.SetCorrelate(correlateOpt)
				createOpt.SetDisabled(false)
				createOpt.SetMatch(&matchOpt)
				createOpt.SetPeriod(2)
				createOpt.SetThreshold(40)

				createResult, createResp, createErr := service.CreateZoneRateLimits(createOpt)
				Expect(createErr).To(BeNil())
				Expect(createResp).ToNot(BeNil())
				Expect(createResult).ToNot(BeNil())
				Expect(*createResult.Success).Should(BeTrue())

				// update rate limit
				respOpt, err = service.NewRatelimitInputMatchResponseHeadersItem("Cf-Cache-Status", RatelimitInputMatchResponseHeadersItem_Op_Ne, RatelimitInputMatchResponseHeadersItem_Value_Hit)
				Expect(err).To(BeNil())
				respItems = []RatelimitInputMatchResponseHeadersItem{*respOpt}

				matchRespOpt = &RatelimitInputMatchResponse{
					HeadersVar: respItems,
				}

				matchOpt = RatelimitInputMatch{
					Request:  matchReqOpt,
					Response: matchRespOpt,
				}
				updateOpt := service.NewUpdateRateLimitOptions(*createResult.Result.ID)
				updateOpt.SetAction(actionOpt)
				updateOpt.SetBypass([]RatelimitInputBypassItem{*bypassOpt})
				updateOpt.SetCorrelate(correlateOpt)
				updateOpt.SetDisabled(false)
				updateOpt.SetMatch(&matchOpt)
				updateOpt.SetPeriod(2)
				updateOpt.SetThreshold(40)

				updateResult, updateResp, updateErr := service.UpdateRateLimit(updateOpt)
				Expect(updateErr).To(BeNil())
				Expect(updateResp).ToNot(BeNil())
				Expect(updateResult).ToNot(BeNil())
				Expect(*updateResult.Success).Should(BeTrue())

				getOpt := service.NewGetRateLimitOptions(*createResult.Result.ID)
				getResult, getResp, getErr := service.GetRateLimit(getOpt)
				Expect(getErr).To(BeNil())
				Expect(getResp).ToNot(BeNil())
				Expect(getResult).ToNot(BeNil())
				Expect(*getResult.Success).Should(BeTrue())

				deleteOpt := service.NewDeleteZoneRateLimitOptions(*createResult.Result.ID)
				deleteResult, deleteResp, deleteErr := service.DeleteZoneRateLimit(deleteOpt)
				Expect(deleteErr).To(BeNil())
				Expect(deleteResp).ToNot(BeNil())
				Expect(deleteResult).ToNot(BeNil())
				Expect(*deleteResult.Success).Should(BeTrue())
			})
			It(`zone rate limit modes test [smulate & Ban]`, func() {
				shouldSkipTest()

				createModes := []string{
					RatelimitInputAction_Mode_Simulate,
					RatelimitInputAction_Mode_Ban,
				}
				for j := 1; j < 5; j++ {
					for _, mode := range createModes {
						url := fmt.Sprintf("*.example%d.org/path*", j)
						// create rate limit

						actionOpt, err := service.NewRatelimitInputAction(mode)
						Expect(err).To(BeNil())
						actionOpt.Timeout = core.Int64Ptr(60)
						actionOpt.Response = &RatelimitInputActionResponse{
							ContentType: core.StringPtr(RatelimitInputActionResponse_ContentType_TextPlain),
							Body:        core.StringPtr("This request has been rate-limited."),
						}

						bypassOpt, err := service.NewRatelimitInputBypassItem(RatelimitInputBypassItem_Name_URL, "api.example.com/*")
						Expect(err).To(BeNil())

						correlateOpt, err := service.NewRatelimitInputCorrelate(RatelimitInputCorrelate_By_Nat)
						Expect(err).To(BeNil())

						matchReqOpt, err := service.NewRatelimitInputMatchRequest(url)
						Expect(err).To(BeNil())
						matchReqOpt.Methods = []string{RatelimitInputMatchRequest_Methods_All}
						matchReqOpt.Schemes = []string{RatelimitInputMatchRequest_Schemes_All}

						matchOpt := RatelimitInputMatch{
							Request: matchReqOpt,
						}

						createOpt := service.NewCreateZoneRateLimitsOptions()
						createOpt.SetAction(actionOpt)
						createOpt.SetBypass([]RatelimitInputBypassItem{*bypassOpt})
						createOpt.SetCorrelate(correlateOpt)
						createOpt.SetDisabled(false)
						createOpt.SetMatch(&matchOpt)
						createOpt.SetPeriod(2)
						createOpt.SetThreshold(40)

						createResult, createResp, createErr := service.CreateZoneRateLimits(createOpt)
						Expect(createErr).To(BeNil())
						Expect(createResp).ToNot(BeNil())
						Expect(createResult).ToNot(BeNil())
						Expect(*createResult.Success).Should(BeTrue())

						listOpt := service.NewListAllZoneRateLimitsOptions()
						listResult, listResp, listErr := service.ListAllZoneRateLimits(listOpt)
						Expect(listErr).To(BeNil())
						Expect(listResp).ToNot(BeNil())
						Expect(listResult).ToNot(BeNil())
						Expect(*listResult.Success).Should(BeTrue())
					}
				}
			})
		})
	})
})
