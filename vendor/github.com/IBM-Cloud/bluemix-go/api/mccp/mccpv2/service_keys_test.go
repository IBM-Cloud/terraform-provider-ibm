package mccpv2

import (
	"log"
	"net/http"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/client"
	bluemixHttp "github.com/IBM-Cloud/bluemix-go/http"
	"github.com/IBM-Cloud/bluemix-go/session"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("ServiceKeys", func() {
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
						ghttp.VerifyRequest(http.MethodPost, "/v2/service_keys"),
						ghttp.VerifyBody([]byte(`{"name":"testkey","service_instance_guid":"f91adfe2-76c9-4649-939e-b01c37a3704c"}`)),
						ghttp.RespondWith(http.StatusCreated, `{							 	
								"metadata": {
									"guid": "c4432b8e-cb14-4225-aa10-1f775b3b1c92",
									"url": "/v2/service_keys/c4432b8e-cb14-4225-aa10-1f775b3b1c92",
									"created_at": "2017-05-04T08:02:50Z",
									"updated_at": null
								},
								"entity": {
										"name": "testkey",
										"service_instance_guid": "f91adfe2-76c9-4649-939e-b01c37a3704c",
										"credentials": {
											"username": "c59f2b23-c2ee-4bc2-baad-2d231da22c01-bluemix",
											"password":"[PRIVATE DATA HIDDEN]",
											"host": "c59f2b23-c2ee-4bc2-baad-2d231da22c01-bluemix.cloudant.com",
											"port": 443,
											"url": "https://c59f2b23-c2ee-4bc2-baad-2d231da22c01-bluemix:8b54e955ebd83c07e1f6cb78edeece5ad53f92ffa75962daa0087265768f4ee6@c59f2b23-c2ee-4bc2-baad-2d231da22c01-bluemix.cloudant.com"
										},
										"service_instance_url": "/v2/service_instances/f91adfe2-76c9-4649-939e-b01c37a3704c"
									}

						}`),
					),
				)
			})

			It("should return service key created", func() {
				myservicekey, err := newServiceKeys(server.URL()).Create("f91adfe2-76c9-4649-939e-b01c37a3704c", "testkey", nil)
				Expect(err).NotTo(HaveOccurred())
				Expect(myservicekey).ShouldNot(BeNil())
				Expect(myservicekey.Metadata.GUID).Should(Equal("c4432b8e-cb14-4225-aa10-1f775b3b1c92"))
				Expect(myservicekey.Entity.Name).Should(Equal("testkey"))
				Expect(myservicekey.Entity.ServiceInstanceGUID).Should(Equal("f91adfe2-76c9-4649-939e-b01c37a3704c"))
				Expect(myservicekey.Entity.Credentials).ShouldNot(BeNil())
			})
		})
		Context("When creation is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v2/service_keys"),
						ghttp.VerifyBody([]byte(`{"name":"testkey","service_instance_guid":"f91adfe2-76c9-4649-939e-b01c37a3704c"}`)),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to create service key`),
					),
				)
			})

			It("should return error during service key creation", func() {
				myservicekey, err := newServiceKeys(server.URL()).Create("f91adfe2-76c9-4649-939e-b01c37a3704c", "testkey", nil)
				Expect(err).To(HaveOccurred())
				Expect(myservicekey).Should(BeNil())
			})
		})

	})

	Describe("Get", func() {
		Context("When read of service key is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/service_keys/c4432b8e-cb14-4225-aa10-1f775b3b1c92"),
						ghttp.RespondWith(http.StatusOK, `{
							"metadata": {
									"guid": "c4432b8e-cb14-4225-aa10-1f775b3b1c92",
									"url": "/v2/service_keys/c4432b8e-cb14-4225-aa10-1f775b3b1c92",
									"created_at": "2017-05-04T08:02:50Z",
									"updated_at": null
								},
							"entity": {
									"name": "testkey",
									"service_instance_guid": "f91adfe2-76c9-4649-939e-b01c37a3704c",
									"credentials": {
										"username": "c59f2b23-c2ee-4bc2-baad-2d231da22c01-bluemix",
										"password":"[PRIVATE DATA HIDDEN]",
										"host": "c59f2b23-c2ee-4bc2-baad-2d231da22c01-bluemix.cloudant.com",
										"port": 443,
										"url": "https://c59f2b23-c2ee-4bc2-baad-2d231da22c01-bluemix:8b54e955ebd83c07e1f6cb78edeece5ad53f92ffa75962daa0087265768f4ee6@c59f2b23-c2ee-4bc2-baad-2d231da22c01-bluemix.cloudant.com"
									},
									"service_instance_url": "/v2/service_instances/f91adfe2-76c9-4649-939e-b01c37a3704c"
								}

						}`),
					),
				)
			})

			It("should return service key", func() {
				myservicekey, err := newServiceKeys(server.URL()).Get("c4432b8e-cb14-4225-aa10-1f775b3b1c92")
				Expect(err).NotTo(HaveOccurred())
				Expect(myservicekey).ShouldNot(BeNil())
				Expect(myservicekey.Metadata.GUID).Should(Equal("c4432b8e-cb14-4225-aa10-1f775b3b1c92"))
				Expect(myservicekey.Entity.Name).Should(Equal("testkey"))
				Expect(myservicekey.Entity.ServiceInstanceGUID).Should(Equal("f91adfe2-76c9-4649-939e-b01c37a3704c"))
				Expect(myservicekey.Entity.Credentials).ShouldNot(BeNil())
			})
		})
		Context("When service key retrieve is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/service_keys/c4432b8e-cb14-4225-aa10-1f775b3b1c92"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve service key`),
					),
				)
			})

			It("should return error when service key is retrieved", func() {
				myservicekey, err := newServiceKeys(server.URL()).Get("c4432b8e-cb14-4225-aa10-1f775b3b1c92")
				Expect(err).To(HaveOccurred())
				Expect(myservicekey).Should(BeNil())
			})
		})
	})

	Describe("Delete", func() {
		Context("When delete of service key is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v2/service_keys/c4432b8e-cb14-4225-aa10-1f775b3b1c92"),
						ghttp.RespondWith(http.StatusOK, `{							
						}`),
					),
				)
			})

			It("should delete service key", func() {
				err := newServiceKeys(server.URL()).Delete("c4432b8e-cb14-4225-aa10-1f775b3b1c92")
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When service key delete is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v2/service_keys/c4432b8e-cb14-4225-aa10-1f775b3b1c92"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to delete service key`),
					),
				)
			})

			It("should return error service key delete", func() {
				err := newServiceKeys(server.URL()).Delete("c4432b8e-cb14-4225-aa10-1f775b3b1c92")
				Expect(err).To(HaveOccurred())
			})
		})
	})

})

var _ = Describe("Service Key by Name", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})
	Describe("FindByName()", func() {
		Context("Server return service key by name", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/service_instances/f91adfe2-76c9-4649-939e-b01c37a3704c/service_keys"),
						ghttp.RespondWith(http.StatusOK, `{
							"total_results": 1,
							"total_pages": 1,
							"prev_url": null,
							"next_url": null,
							"resources": [
								{
									"metadata": {
										"guid": "c4432b8e-cb14-4225-aa10-1f775b3b1c92",
										"url": "/v2/service_keys/c4432b8e-cb14-4225-aa10-1f775b3b1c92",
										"created_at": "2017-05-04T08:02:50Z",
										"updated_at": null
									},
									"entity": {
										"name": "testkey",
										"service_instance_guid": "f91adfe2-76c9-4649-939e-b01c37a3704c",
										"credentials": {
											"username": "c59f2b23-c2ee-4bc2-baad-2d231da22c01-bluemix",
											"password":"[PRIVATE DATA HIDDEN]",
											"host": "c59f2b23-c2ee-4bc2-baad-2d231da22c01-bluemix.cloudant.com",
											"port": 443,
											"url": "https://c59f2b23-c2ee-4bc2-baad-2d231da22c01-bluemix:8b54e955ebd83c07e1f6cb78edeece5ad53f92ffa75962daa0087265768f4ee6@c59f2b23-c2ee-4bc2-baad-2d231da22c01-bluemix.cloudant.com"
										},
										"service_instance_url": "/v2/service_instances/f91adfe2-76c9-4649-939e-b01c37a3704c"
									}
								}
							]

						}`),
					),
				)
			})

			It("should return service key", func() {
				myservicekey, err := newServiceKeys(server.URL()).FindByName("f91adfe2-76c9-4649-939e-b01c37a3704c", "testkey")
				Expect(err).NotTo(HaveOccurred())
				Expect(myservicekey).ShouldNot(BeNil())
				Expect(myservicekey.GUID).Should(Equal("c4432b8e-cb14-4225-aa10-1f775b3b1c92"))
				Expect(myservicekey.Name).Should(Equal("testkey"))
				Expect(myservicekey.ServiceInstanceGUID).Should(Equal("f91adfe2-76c9-4649-939e-b01c37a3704c"))
				Expect(myservicekey.Credentials).ShouldNot(BeNil())
			})

		})

		Context("Server return no service key", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/service_instances/f91adfe2-76c9-4649-939e-b01c37a3704c/service_keys"),
						ghttp.RespondWith(http.StatusOK, `{
							"total_results": 0,
							"resources": [
							]
															
						}`),
					),
				)
			})

			It("should return no service key", func() {
				myservicekey, err := newServiceKeys(server.URL()).FindByName("f91adfe2-76c9-4649-939e-b01c37a3704c", "testkey")
				Expect(err).To(HaveOccurred())
				Expect(myservicekey).To(BeNil())
			})

		})
		Context("Server return error", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/service_instances/f91adfe2-76c9-4649-939e-b01c37a3704c/service_keys"),
						ghttp.RespondWith(http.StatusInternalServerError, `{
															
						}`),
					),
				)
			})

			It("should return error service key", func() {
				myservicekey, err := newServiceKeys(server.URL()).FindByName("f91adfe2-76c9-4649-939e-b01c37a3704c", "testkey")
				Expect(err).To(HaveOccurred())
				Expect(myservicekey).To(BeNil())
			})

		})

	})
})

func newServiceKeys(url string) ServiceKeys {

	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}
	conf := sess.Config.Copy()
	conf.HTTPClient = bluemixHttp.NewHTTPClient(conf)
	conf.Endpoint = &url

	client := client.Client{
		Config:      conf,
		ServiceName: bluemix.MccpService,
	}
	return newServiceKeyAPI(&client)
}
