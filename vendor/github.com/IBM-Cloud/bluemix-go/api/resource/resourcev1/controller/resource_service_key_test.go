package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/IBM-Cloud/bluemix-go/client"
	"github.com/IBM-Cloud/bluemix-go/crn"

	"github.com/IBM-Cloud/bluemix-go"

	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/onsi/gomega/ghttp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ServiceKeys", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})

	Describe("GetKeys()", func() {
		Context("When there is no service key", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/resource_keys"),
						ghttp.RespondWith(http.StatusOK, `{"resources":[]}`),
					),
				)
			})
			It("should return zero service key", func() {
				repo := newTestServiceKeyRepo(server.URL())
				keys, err := repo.GetKeys("")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(keys).Should(BeEmpty())
			})
		})
		Context("When there is one service key", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/resource_keys"),
						ghttp.RespondWith(http.StatusOK, `{"rows_count":1,
																	"resources":[
																		{"id":"fdb5785b-c2c1-4f5e-b4ea-686083af250d",
																		"url":"/v1/resource_keys/fdb5785b-c2c1-4f5e-b4ea-686083af250d",
																		"created_at":"2017-08-04T05:53:55.69509063Z",
																		"updated_at":null,
																		"deleted_at":null,
																		"source_crn":"crn:v1:staging:public:rc-demo-services-cfmysqlv3:us-south:a/560df2058b1e7c402303cc598b3e5540:6d5bd9c5-0412-4f5b-a39a-f322be3d7812::",
																		"parameters":{"role_crn":"crn:v1:bluemix:public:iam::::role:Operator"},
																		"crn":"crn:v1:staging:public:rc-demo-services-cfmysqlv3:us-south:a/560df2058b1e7c402303cc598b3e5540:6d5bd9c5-0412-4f5b-a39a-f322be3d7812:resource-key:fdb5785b-c2c1-4f5e-b4ea-686083af250d",
																		"state":"active",
																		"account_id":"560df2058b1e7c402303cc598b3e5540",
																		"credentials":
																			{"ApiKey-2bdfd49d-889e-4270-8d96-ed4206415b09":
																				{"boundTo":"crn:v1:::iam::a/560df2058b1e7c402303cc598b3e5540::serviceid:ServiceId-7ee2fcde-da3e-4880-93d9-7e2fa26238fc",
																				"name":"fdb5785b-c2c1-4f5e-b4ea-686083af250d-key2",
																				"description":"fdb5785b-c2c1-4f5e-b4ea-686083af250d-key2",
																				"format":"APIKEY",
																				"apiKey":"[PRIVATE DATA HIDDEN]"},
																			"ApiKey-e865a2b1-08e5-4c85-8756-0299286f6a79":
																				{"boundTo":"crn:v1:::iam::a/560df2058b1e7c402303cc598b3e5540::serviceid:ServiceId-7ee2fcde-da3e-4880-93d9-7e2fa26238fc",
																				"name":"fdb5785b-c2c1-4f5e-b4ea-686083af250d-key1","description":"fdb5785b-c2c1-4f5e-b4ea-686083af250d-key1",
																				"format":"APIKEY",
																				"apiKey":"[PRIVATE DATA HIDDEN]"
																			}}}]}`),
					),
				)
			})
			It("should return one service key", func() {
				repo := newTestServiceKeyRepo(server.URL())
				keys, err := repo.GetKeys("")

				Expect(err).ShouldNot(HaveOccurred())

				Expect(keys).Should(HaveLen(1))
				key := keys[0]
				Expect(key.ID).Should(Equal("fdb5785b-c2c1-4f5e-b4ea-686083af250d"))
				Expect(key.SourceCrn.String()).Should(Equal("crn:v1:staging:public:rc-demo-services-cfmysqlv3:us-south:a/560df2058b1e7c402303cc598b3e5540:6d5bd9c5-0412-4f5b-a39a-f322be3d7812::"))
			})
		})
	})

	Describe("CreateKey()", func() {

		var (
			roleCRN   crn.CRN
			sourceCRN crn.CRN

			serverStatus   int
			serverResponse string
		)

		BeforeEach(func() {
			var err error
			roleCRN, err = crn.Parse("crn:v1:bluemix:public:iam::::role:Operator")
			Expect(err).NotTo(HaveOccurred())

			sourceCRN, err = crn.Parse("crn:v1:staging:public:rc-demo-services-cfmysqlv3:us-south:a/560df2058b1e7c402303cc598b3e5540:6d5bd9c5-0412-4f5b-a39a-f322be3d7812::")
			Expect(err).NotTo(HaveOccurred())

			serverStatus = http.StatusOK
		})

		JustBeforeEach(func() {
			server = ghttp.NewServer()
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodPost, "/v1/resource_keys"),
					ghttp.VerifyJSON(fmt.Sprintf(`{
							"name": "KeyName", 
							"source_crn": "%s",
							"parameters": {
								"role_crn": "%s"
							}
						}`, sourceCRN.String(), roleCRN.String())),
					ghttp.RespondWith(serverStatus, serverResponse),
				),
			)
		})

		Context("when creation is successful", func() {
			BeforeEach(func() {
				serverStatus = http.StatusOK
				serverResponse = `{"id":"fdb5785b-c2c1-4f5e-b4ea-686083af250d",
									"url":"/v1/resource_keys/fdb5785b-c2c1-4f5e-b4ea-686083af250d",
									"created_at":"2017-08-04T05:53:55.69509063Z",
									"updated_at":null,
									"deleted_at":null,
									"source_crn":"crn:v1:staging:public:rc-demo-services-cfmysqlv3:us-south:a/560df2058b1e7c402303cc598b3e5540:6d5bd9c5-0412-4f5b-a39a-f322be3d7812::",
									"parameters":{
										"role_crn":"crn:v1:bluemix:public:iam::::role:Operator"},
									"crn":"crn:v1:staging:public:rc-demo-services-cfmysqlv3:us-south:a/560df2058b1e7c402303cc598b3e5540:6d5bd9c5-0412-4f5b-a39a-f322be3d7812:resource-key:fdb5785b-c2c1-4f5e-b4ea-686083af250d",
									"state":"active",
									"account_id":"560df2058b1e7c402303cc598b3e5540",
									"credentials":
										{"ApiKey-2bdfd49d-889e-4270-8d96-ed4206415b09":
											{"boundTo":"crn:v1:::iam::a/560df2058b1e7c402303cc598b3e5540::serviceid:ServiceId-7ee2fcde-da3e-4880-93d9-7e2fa26238fc",
											"name":"fdb5785b-c2c1-4f5e-b4ea-686083af250d-key2","description":"fdb5785b-c2c1-4f5e-b4ea-686083af250d-key2",
											"format":"APIKEY",
											"apiKey":"[PRIVATE DATA HIDDEN]"},
										"ApiKey-e865a2b1-08e5-4c85-8756-0299286f6a79":
											{"boundTo":"crn:v1:::iam::a/560df2058b1e7c402303cc598b3e5540::serviceid:ServiceId-7ee2fcde-da3e-4880-93d9-7e2fa26238fc",
											"name":"fdb5785b-c2c1-4f5e-b4ea-686083af250d-key1",
											"description":"fdb5785b-c2c1-4f5e-b4ea-686083af250d-key1",
											"format":"APIKEY",
											"apiKey":"[PRIVATE DATA HIDDEN]"
									}}}`
			})

			It("should return the new service keys", func() {
				parameters := make(map[string]interface{})
				parameters["role_crn"] = roleCRN
				key, err := newTestServiceKeyRepo(server.URL()).CreateKey(CreateServiceKeyRequest{
					Name:       "KeyName",
					SourceCRN:  sourceCRN,
					Parameters: parameters,
				})

				Expect(err).ShouldNot(HaveOccurred())
				Expect(key).ShouldNot(BeNil())
				Expect(key.ID).Should(Equal("fdb5785b-c2c1-4f5e-b4ea-686083af250d"))
			})
		})

		Context("when creation failed", func() {
			BeforeEach(func() {
				serverStatus = http.StatusBadRequest
				serverResponse = "invalid role"
			})

			It("should return error", func() {
				parameters := make(map[string]interface{})
				parameters["role_crn"] = roleCRN
				_, err := newTestServiceKeyRepo(server.URL()).CreateKey(CreateServiceKeyRequest{
					Name:       "KeyName",
					SourceCRN:  sourceCRN,
					Parameters: parameters,
				})
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("DeleteKey()", func() {
		Context("When deletion is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/resource_keys/db430fea-d29a-4e11-b952-b1e3e18a4ac4"),
						ghttp.RespondWith(http.StatusNoContent, ``),
					),
				)
			})
			It("should return success", func() {
				err := newTestServiceKeyRepo(server.URL()).DeleteKey("db430fea-d29a-4e11-b952-b1e3e18a4ac4")
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("When deletion failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/resource_keys/abc"),
						ghttp.RespondWith(http.StatusNotFound, `{"message":"Not found"}`),
					),
				)
			})
			It("should return error", func() {
				err := newTestServiceKeyRepo(server.URL()).DeleteKey("abc")
				Expect(err).Should(HaveOccurred())
			})
		})
	})
})

func newTestServiceKeyRepo(url string) ResourceServiceKeyRepository {
	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}
	conf := sess.Config.Copy()
	conf.Endpoint = &url

	client := client.Client{
		Config:      conf,
		ServiceName: bluemix.ResourceManagementService,
	}

	return newResourceServiceKeyAPI(&client)
}
