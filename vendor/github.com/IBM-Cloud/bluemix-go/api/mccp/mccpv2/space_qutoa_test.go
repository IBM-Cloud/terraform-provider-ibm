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

var _ = Describe("SpaceQuotas", func() {
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
						ghttp.VerifyRequest(http.MethodPost, "/v2/space_quota_definitions"),
						ghttp.VerifyBody([]byte(`{"name":"test-quota","organization_guid":"3c1b6f9d-ffe5-43b5-ab91-7be2331dc546","memory_limit":1024,"instance_memory_limit":1024,"total_routes":50,"total_services":150,"non_basic_services_allowed":false}`)),
						ghttp.RespondWith(http.StatusCreated, `{
							 	
								"metadata": {
									"guid": "be829072-3137-418c-9607-c84e7d77e22a",
									"url": "/v2/space_quota_definitions/be829072-3137-418c-9607-c84e7d77e22a",
									"created_at": "2017-05-04T04:31:19Z",
									"updated_at": null
								},
								"entity": {
									"name": "test-quota",
									"organization_guid": "3c1b6f9d-ffe5-43b5-ab91-7be2331dc546",
									"non_basic_services_allowed": false,
									"total_services": 150,
									"total_routes": 50,
									"memory_limit": 1024,
									"instance_memory_limit": 1024,
									"app_instance_limit": -1,
									"app_task_limit": 5,
									"total_service_keys": -1,
									"organization_url": "/v2/organizations/3c1b6f9d-ffe5-43b5-ab91-7be2331dc546",
									"spaces_url": "/v2/space_quota_definitions/be829072-3137-418c-9607-c84e7d77e22a/spaces"
								}
				
						}`),
					),
				)
			})

			It("should return SpaceQuota created", func() {
				payload := SpaceQuotaCreateRequest{
					Name:                    "test-quota",
					OrgGUID:                 "3c1b6f9d-ffe5-43b5-ab91-7be2331dc546",
					MemoryLimitInMB:         1024,
					InstanceMemoryLimitInMB: 1024,
					RoutesLimit:             50,
					ServicesLimit:           150,
					NonBasicServicesAllowed: false,
				}
				myspacequota, err := newSpaceQuotas(server.URL()).Create(payload)
				Expect(err).NotTo(HaveOccurred())
				Expect(myspacequota).ShouldNot(BeNil())
				Expect(myspacequota.Metadata.GUID).Should(Equal("be829072-3137-418c-9607-c84e7d77e22a"))
				Expect(myspacequota.Entity.Name).Should(Equal("test-quota"))
				Expect(myspacequota.Entity.NonBasicServicesAllowed).To(BeFalse())
			})
		})
		Context("When creation is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v2/space_quota_definitions"),
						ghttp.VerifyBody([]byte(`{"name":"test-quota","organization_guid":"3c1b6f9d-ffe5-43b5-ab91-7be2331dc546","memory_limit":1024,"instance_memory_limit":1024,"total_routes":50,"total_services":150,"non_basic_services_allowed":false}`)),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to create space quota`),
					),
				)
			})

			It("should return error during spacequota creation", func() {
				payload := SpaceQuotaCreateRequest{
					Name:                    "test-quota",
					OrgGUID:                 "3c1b6f9d-ffe5-43b5-ab91-7be2331dc546",
					MemoryLimitInMB:         1024,
					InstanceMemoryLimitInMB: 1024,
					RoutesLimit:             50,
					ServicesLimit:           150,
					NonBasicServicesAllowed: false,
				}
				myspacequota, err := newSpaceQuotas(server.URL()).Create(payload)
				Expect(err).To(HaveOccurred())
				Expect(myspacequota).Should(BeNil())
			})
		})

	})

	Describe("Get", func() {
		Context("When read of space quota is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/space_quota_definitions/be829072-3137-418c-9607-c84e7d77e22a"),
						ghttp.RespondWith(http.StatusOK, `{
							 "metadata": {
									"guid": "be829072-3137-418c-9607-c84e7d77e22a",
									"url": "/v2/space_quota_definitions/be829072-3137-418c-9607-c84e7d77e22a",
									"created_at": "2017-05-04T04:31:19Z",
									"updated_at": null
								},
								"entity": {
									"name": "test-quota",
									"organization_guid": "3c1b6f9d-ffe5-43b5-ab91-7be2331dc546",
									"non_basic_services_allowed": false,
									"total_services": 150,
									"total_routes": 50,
									"memory_limit": 1024,
									"instance_memory_limit": 1024,
									"app_instance_limit": -1,
									"app_task_limit": 5,
									"total_service_keys": -1,
									"organization_url": "/v2/organizations/3c1b6f9d-ffe5-43b5-ab91-7be2331dc546",
									"spaces_url": "/v2/space_quota_definitions/be829072-3137-418c-9607-c84e7d77e22a/spaces"
								}		
						}`),
					),
				)
			})

			It("should return space quota", func() {
				myspacequota, err := newSpaceQuotas(server.URL()).Get("be829072-3137-418c-9607-c84e7d77e22a")
				Expect(err).NotTo(HaveOccurred())
				Expect(myspacequota).ShouldNot(BeNil())
				Expect(myspacequota.Metadata.GUID).Should(Equal("be829072-3137-418c-9607-c84e7d77e22a"))
				Expect(myspacequota.Entity.Name).Should(Equal("test-quota"))
				Expect(myspacequota.Entity.NonBasicServicesAllowed).To(BeFalse())
			})
		})
		Context("When space quota retrieve is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/space_quota_definitions/be829072-3137-418c-9607-c84e7d77e22a"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve space quota`),
					),
				)
			})

			It("should return error when space quota is retrieved", func() {
				myspacequota, err := newSpaceQuotas(server.URL()).Get("be829072-3137-418c-9607-c84e7d77e22a")
				Expect(err).To(HaveOccurred())
				Expect(myspacequota).Should(BeNil())
			})
		})
	})
	Describe("Update", func() {
		Context("When update of space quota is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, "/v2/space_quota_definitions/be829072-3137-418c-9607-c84e7d77e22a"),
						ghttp.VerifyBody([]byte(`{"name":"testspacequotaupdate","organization_guid":"3c1b6f9d-ffe5-43b5-ab91-7be2331dc546","memory_limit":1024,"instance_memory_limit":1024,"total_routes":50,"total_services":150,"non_basic_services_allowed":false}`)),
						ghttp.RespondWith(http.StatusCreated, `{
							 "metadata": {
									"guid": "be829072-3137-418c-9607-c84e7d77e22a",
									"url": "/v2/space_quota_definitions/be829072-3137-418c-9607-c84e7d77e22a",
									"created_at": "2017-05-04T04:31:19Z",
									"updated_at": null
								},
								"entity": {
									"name": "testspacequotaupdate",
									"organization_guid": "3c1b6f9d-ffe5-43b5-ab91-7be2331dc546",
									"non_basic_services_allowed": false,
									"total_services": 150,
									"total_routes": 50,
									"memory_limit": 1024,
									"instance_memory_limit": 1024,
									"app_instance_limit": -1,
									"app_task_limit": 5,
									"total_service_keys": -1,
									"organization_url": "/v2/organizations/3c1b6f9d-ffe5-43b5-ab91-7be2331dc546",
									"spaces_url": "/v2/space_quota_definitions/be829072-3137-418c-9607-c84e7d77e22a/spaces"
								}			
						}`),
					),
				)
			})

			It("should return spacequota update", func() {
				payload := SpaceQuotaUpdateRequest{
					Name:                    "testspacequotaupdate",
					OrgGUID:                 "3c1b6f9d-ffe5-43b5-ab91-7be2331dc546",
					MemoryLimitInMB:         1024,
					InstanceMemoryLimitInMB: 1024,
					RoutesLimit:             50,
					ServicesLimit:           150,
					NonBasicServicesAllowed: false,
				}
				myspacequota, err := newSpaceQuotas(server.URL()).Update(payload, "be829072-3137-418c-9607-c84e7d77e22a")
				Expect(err).NotTo(HaveOccurred())
				Expect(myspacequota).ShouldNot(BeNil())
				Expect(myspacequota.Metadata.GUID).Should(Equal("be829072-3137-418c-9607-c84e7d77e22a"))
				Expect(myspacequota.Entity.Name).Should(Equal("testspacequotaupdate"))
				Expect(myspacequota.Entity.NonBasicServicesAllowed).To(BeFalse())
			})
		})
		Context("When spacequota update is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, "/v2/space_quota_definitions/be829072-3137-418c-9607-c84e7d77e22a"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to update space quota`),
					),
				)
			})

			It("should return error spacequota updated", func() {
				payload := SpaceQuotaUpdateRequest{
					Name:                    "test-quota",
					OrgGUID:                 "3c1b6f9d-ffe5-43b5-ab91-7be2331dc546",
					MemoryLimitInMB:         1024,
					InstanceMemoryLimitInMB: 1024,
					RoutesLimit:             50,
					ServicesLimit:           150,
					NonBasicServicesAllowed: false,
				}
				myspacequota, err := newSpaceQuotas(server.URL()).Update(payload, "be829072-3137-418c-9607-c84e7d77e22a")
				Expect(err).To(HaveOccurred())
				Expect(myspacequota).Should(BeNil())
			})
		})
	})

	Describe("Delete", func() {
		Context("When delete of space quota is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v2/space_quota_definitions/be829072-3137-418c-9607-c84e7d77e22a"),
						ghttp.RespondWith(http.StatusOK, `{							
						}`),
					),
				)
			})

			It("should delete Space quota", func() {
				err := newSpaceQuotas(server.URL()).Delete("be829072-3137-418c-9607-c84e7d77e22a")
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When space quota delete is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v2/space_quota_definitions/be829072-3137-418c-9607-c84e7d77e22a"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to delete space quota`),
					),
				)
			})

			It("should return error space quota delete", func() {
				err := newSpaceQuotas(server.URL()).Delete("be829072-3137-418c-9607-c84e7d77e22a")
				Expect(err).To(HaveOccurred())
			})
		})
	})

})

var _ = Describe("Space Quota by Name", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})
	Describe("FindByNameInOrg()", func() {
		Context("Server return space quota by name", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/organizations/3c1b6f9d-ffe5-43b5-ab91-7be2331dc546/space_quota_definitions"),
						ghttp.RespondWith(http.StatusOK, `{
							"total_results": 2,
							"total_pages": 1,
							"prev_url": null,
							"next_url": null,
							"resources": [
							{
								"metadata": {
									"guid": "1d11ba0f-aa87-417c-b390-2e4d035279d6",
									"url": "/v2/space_quota_definitions/1d11ba0f-aa87-417c-b390-2e4d035279d6",
									"created_at": "2017-05-04T06:27:51Z",
									"updated_at": null

								},
								"entity": {
									"name": "new",
									"organization_guid": "3c1b6f9d-ffe5-43b5-ab91-7be2331dc546",
									"non_basic_services_allowed": false,
									"total_services": 0,
									"total_routes": 0,
									"memory_limit": 0,
									"instance_memory_limit": -1,
									"app_instance_limit": -1,
									"app_task_limit": 5,
									"total_service_keys": -1,
									"organization_url": "/v2/organizations/3c1b6f9d-ffe5-43b5-ab91-7be2331dc546",
									"spaces_url": "/v2/space_quota_definitions/1d11ba0f-aa87-417c-b390-2e4d035279d6/spaces"
								}
										
							},
							{
								"metadata": {
									"guid": "be829072-3137-418c-9607-c84e7d77e22a",
									"url": "/v2/space_quota_definitions/be829072-3137-418c-9607-c84e7d77e22a",
									"created_at": "2017-05-04T06:45:46Z",
									"updated_at": "2017-05-04T06:45:46Z"

								},
								"entity": {
									"name": "testspacequotaupdate",
									"organization_guid": "3c1b6f9d-ffe5-43b5-ab91-7be2331dc546",
									"non_basic_services_allowed": false,
									"total_services": 150,
									"total_routes": 50,
									"memory_limit": 1024,
									"instance_memory_limit": 1024,
									"app_instance_limit": -1,
									"app_task_limit": 5,
									"total_service_keys": -1,
									"organization_url": "/v2/organizations/3c1b6f9d-ffe5-43b5-ab91-7be2331dc546",
									"spaces_url": "/v2/space_quota_definitions/be829072-3137-418c-9607-c84e7d77e22a/spaces"
       
								}
							}
							]
															
						}`),
					),
				)
			})

			It("should return space quota by name", func() {
				myspacequota, err := newSpaceQuotas(server.URL()).FindByName("testspacequotaupdate", "3c1b6f9d-ffe5-43b5-ab91-7be2331dc546")
				Expect(err).NotTo(HaveOccurred())
				Expect(myspacequota).ShouldNot(BeNil())
				Expect(myspacequota.GUID).Should(Equal("be829072-3137-418c-9607-c84e7d77e22a"))
				Expect(myspacequota.Name).Should(Equal("testspacequotaupdate"))
				Expect(myspacequota.NonBasicServicesAllowed).To(BeFalse())
			})

		})

		Context("Server return no spacequota by name", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/organizations/3c1b6f9d-ffe5-43b5-ab91-7be2331dc546/space_quota_definitions"),
						ghttp.RespondWith(http.StatusOK, `{
							"total_results": 0,
							"resources": [
							]
															
						}`),
					),
				)
			})

			It("should return no space quota", func() {
				myspacequota, err := newSpaceQuotas(server.URL()).FindByName("testspacequotaupdate", "3c1b6f9d-ffe5-43b5-ab91-7be2331dc546")
				Expect(err).To(HaveOccurred())
				Expect(myspacequota).To(BeNil())
			})

		})
		Context("Server return error", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/organizations/3c1b6f9d-ffe5-43b5-ab91-7be2331dc546/space_quota_definitions"),
						ghttp.RespondWith(http.StatusInternalServerError, `{
															
						}`),
					),
				)
			})

			It("should return no spaces", func() {
				myspacequota, err := newSpaceQuotas(server.URL()).FindByName("testspacequotaupdate", "3c1b6f9d-ffe5-43b5-ab91-7be2331dc546")
				Expect(err).To(HaveOccurred())
				Expect(myspacequota).To(BeNil())
			})

		})

	})
})

func newSpaceQuotas(url string) SpaceQuotas {
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

	return newSpaceQuotasAPI(&client)
}
