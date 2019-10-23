package cisv1

import (
	"log"
	"net/http"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/client"
	bluemixHttp "github.com/IBM-Cloud/bluemix-go/http"
	"github.com/IBM-Cloud/bluemix-go/session"

	"github.com/onsi/gomega/ghttp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Settings", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})

	//Find
	Describe("Get", func() {
		Context("When read of min_tls_version Setting is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/zones/3fefc35e7decadb111dcf85d723a4f20/settings/min_tls_version"),
						ghttp.RespondWith(http.StatusOK, `
                            {
                              "result": {
                                "id": "min_tls_version",
                                "value": "1.2",
                                "modified_on": null,
                                "editable": true
                              },
                              "success": true,
                              "errors": [],
                              "messages": []
                            }
                        `),
					),
				)
			})

			It("should return Setting", func() {
				target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				target2 := "3fefc35e7decadb111dcf85d723a4f20"
				target3 := "min_tls_version"
				mySettingPtr, err := newSetting(server.URL()).GetSetting(target1, target2, target3)
				mySetting := *mySettingPtr
				Expect(err).NotTo(HaveOccurred())
				Expect(mySetting).ShouldNot(BeNil())
				Expect(mySetting.Value).Should(Equal("1.2"))
			})
		})
		Context("When Setting get has failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/zones/3fefc35e7decadb111dcf85d723a4f20/settings/min_tls_version"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve Setting`),
					),
				)
			})

			It("should return error when Setting is retrieved", func() {
				target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				target2 := "3fefc35e7decadb111dcf85d723a4f20"
				target3 := "min_tls_version"
				mySettingPtr, err := newSetting(server.URL()).GetSetting(target1, target2, target3)
				mySetting := mySettingPtr
				Expect(err).To(HaveOccurred())
				Expect(mySetting).Should(BeNil())
			})
		})
	})

	Describe("Update", func() {
		Context("When updating is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPatch, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/zones/3fefc35e7decadb111dcf85d723a4f20/settings/min_tls_version"),
						ghttp.RespondWith(http.StatusCreated, `
                           {
                              "result": {
                                "id": "min_tls_version",
                                "value": "1.2",
                                "modified_on": "2018-10-08T09:37:29.953507Z",
                                "editable": true
                              },
                              "success": true,
                              "errors": [],
                              "messages": []
                            }
                        `),
					),
				)
			})

			It("should return setting created", func() {
				params := SettingsBody{
					Value: "1.2",
				}
				target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				target2 := "3fefc35e7decadb111dcf85d723a4f20"
				target3 := "min_tls_version"
				mySettingPt, err := newSetting(server.URL()).UpdateSetting(target1, target2, target3, params)
				mySetting := *mySettingPt
				Expect(err).NotTo(HaveOccurred())
				Expect(mySetting).ShouldNot(BeNil())
				Expect(mySetting.Value).Should(Equal("1.2"))
			})
		})
		Context("When creation is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPatch, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/zones/3fefc35e7decadb111dcf85d723a4f20/settings/min_tls_version"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to create Setting`),
					),
				)
			})
			It("should return error during Setting creation", func() {
				params := SettingsBody{
					Value: "1.2",
				}
				target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				target2 := "3fefc35e7decadb111dcf85d723a4f20"
				target3 := "min_tls_version"
				mySettingPtr, err := newSetting(server.URL()).UpdateSetting(target1, target2, target3, params)
				mySetting := mySettingPtr
				Expect(err).To(HaveOccurred())
				Expect(mySetting).Should(BeNil())
			})
		})
	})
})

func newSetting(url string) Settings {

	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}
	conf := sess.Config.Copy()
	conf.HTTPClient = bluemixHttp.NewHTTPClient(conf)
	conf.Endpoint = &url

	client := client.Client{
		Config:      conf,
		ServiceName: bluemix.CisService,
	}
	return newSettingsAPI(&client)
}
