package icdv4

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

var _ = Describe("Whitelists", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})
	Describe("Create", func() {
		Context("When creation is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v4/ibm/deployments/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/whitelists/ip_addresses"),
						ghttp.RespondWith(http.StatusCreated, `
                           {
                              "task": {
                                "id": "5abb6a7d11a1a5001479a0ad",
                                "description": "Adding whitelist entry for database",
                                "status": "running",
                                "deployment_id": "59b14b19874a1c0018009483",
                                "progress_percent": 10,
                                "created_at": "2018-03-28T10:21:30Z"
                              }
                            }
                        `),
					),
				)
			})

			It("should return whitelist created", func() {
				whitelistEntry := WhitelistEntry{
					Address:     "172.168.0.1",
					Description: "Address description",
				}
				params := WhitelistReq{
					WhitelistEntry: whitelistEntry,
				}
				target := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				myTask, err := newWhitelist(server.URL()).CreateWhitelist(target, params)
				Expect(err).NotTo(HaveOccurred())
				Expect(myTask).ShouldNot(BeNil())
				Expect(myTask.Id).Should(Equal("5abb6a7d11a1a5001479a0ad"))
				Expect(myTask.Description).Should(Equal("Adding whitelist entry for database"))
				Expect(myTask.Status).Should(Equal("running"))
				Expect(myTask.DeploymentId).Should(Equal("59b14b19874a1c0018009483"))
				Expect(myTask.ProgressPercent).Should(Equal(10))
				Expect(myTask.CreatedAt).Should(Equal("2018-03-28T10:21:30Z"))
			})
		})
		Context("When creation is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v4/ibm/deployments/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/whitelists/ip_addresses"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to create Whitelist`),
					),
				)
			})

			It("should return error during Whitelist creation", func() {
				whitelistEntry := WhitelistEntry{
					Address:     "172.168.0.1",
					Description: "Address description",
				}
				params := WhitelistReq{
					WhitelistEntry: whitelistEntry,
				}
				target := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				myTask, err := newWhitelist(server.URL()).CreateWhitelist(target, params)
				Expect(err).To(HaveOccurred())
				Expect(myTask.Id).Should(Equal(""))
			})
		})
	})
	Describe("Delete", func() {
		Context("When deletion is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v4/ibm/deployments/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/whitelists/ip_addresses/172.168.0.1"),
						ghttp.RespondWith(http.StatusOK, `
                           {
                              "task": {
                                "id": "5abb6a7d11a1a5001479a0af",
                                "description": "Deleting whitelist entry for database",
                                "status": "running",
                                "deployment_id": "59b14b19874a1c0018009483",
                                "progress_percent": 15,
                                "created_at": "2018-03-28T10:25:30Z"
                              }
                            }
                        `),
					),
				)
			})

			It("should return whitelist deleted", func() {
				target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				target2 := "172.168.0.1"
				myTask, err := newWhitelist(server.URL()).DeleteWhitelist(target1, target2)
				Expect(err).NotTo(HaveOccurred())
				Expect(myTask).ShouldNot(BeNil())
				Expect(myTask.Id).Should(Equal("5abb6a7d11a1a5001479a0af"))
				Expect(myTask.Description).Should(Equal("Deleting whitelist entry for database"))
				Expect(myTask.Status).Should(Equal("running"))
				Expect(myTask.DeploymentId).Should(Equal("59b14b19874a1c0018009483"))
				Expect(myTask.ProgressPercent).Should(Equal(15))
				Expect(myTask.CreatedAt).Should(Equal("2018-03-28T10:25:30Z"))
			})
		})
		Context("When deletion is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v4/ibm/deployments/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/whitelists/ip_addresses/172.168.0.1"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to delete Whitelist`),
					),
				)
			})

			It("should return error during Whitelist deletion", func() {
				target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				target2 := "172.168.0.1"
				myTask, err := newWhitelist(server.URL()).DeleteWhitelist(target1, target2)
				Expect(err).To(HaveOccurred())
				Expect(myTask.Id).Should(Equal(""))
			})
		})
	})
	Describe("Get", func() {
		Context("When get is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v4/ibm/deployments/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/whitelists/ip_addresses"),
						ghttp.RespondWith(http.StatusOK, `
                           {
                                "ip_addresses": [
                                    {
                                        "address": "172.168.0.1"
                                    },
                                    {
                                        "address": "172.168.0.2"
                                    }
                                ]
                            }
                        `),
					),
				)
			})

			It("should return whitelist", func() {
				target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				whitelist, err := newWhitelist(server.URL()).GetWhitelist(target1)
				Expect(err).NotTo(HaveOccurred())
				Expect(whitelist).ShouldNot(BeNil())
				Expect(whitelist.WhitelistEntrys[0].Address).Should(Equal("172.168.0.1"))
				Expect(whitelist.WhitelistEntrys[1].Address).Should(Equal("172.168.0.2"))
			})
		})
		Context("When get is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v4/ibm/deployments/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/whitelists/ip_addresses"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to get Whitelist`),
					),
				)
			})

			It("should return error during Whitelist gete", func() {
				target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				_, err := newWhitelist(server.URL()).GetWhitelist(target1)
				Expect(err).To(HaveOccurred())
			})
		})
	})
})

func newWhitelist(url string) Whitelists {

	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}
	conf := sess.Config.Copy()
	conf.HTTPClient = bluemixHttp.NewHTTPClient(conf)
	conf.Endpoint = &url

	client := client.Client{
		Config:      conf,
		ServiceName: bluemix.ICDService,
	}
	return newWhitelistAPI(&client)
}
