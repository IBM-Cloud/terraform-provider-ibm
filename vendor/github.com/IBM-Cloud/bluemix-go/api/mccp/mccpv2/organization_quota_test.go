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

var _ = Describe("OrgQuotas", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})

	Describe("Get", func() {
		Context("When read of organization quota is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/quota_definitions/23d1cc15-c2c4-4641-b6e5-08e5530b8ea8"),
						ghttp.RespondWith(http.StatusOK, `{
							"metadata": {
							  "guid": "23d1cc15-c2c4-4641-b6e5-08e5530b8ea8",
							  "url": "/v2/quota_definitions/23d1cc15-c2c4-4641-b6e5-08e5530b8ea8",
							  "created_at": "2016-04-16T01:23:49Z",
							  "updated_at": null
							},
							"entity": {
							  "name": "gold_quota",
							  "non_basic_services_allowed": true,
							  "total_services": -1,
							  "total_routes": 4,
							  "total_private_domains": -1,
							  "memory_limit": 5120,
							  "trial_db_allowed": false,
							  "instance_memory_limit": 10240,
							  "app_instance_limit": 10,
							  "app_task_limit": 5,
							  "total_service_keys": -1,
							  "total_reserved_route_ports": 3
							}
						  }`),
					),
				)
			})

			It("should return org quota", func() {
				myorgquota, err := newOrgQuotas(server.URL()).Get("23d1cc15-c2c4-4641-b6e5-08e5530b8ea8")
				Expect(err).NotTo(HaveOccurred())
				Expect(myorgquota).ShouldNot(BeNil())
				Expect(myorgquota.Metadata.GUID).Should(Equal("23d1cc15-c2c4-4641-b6e5-08e5530b8ea8"))
				Expect(myorgquota.Entity.Name).Should(Equal("gold_quota"))
				Expect(myorgquota.Entity.NonBasicServicesAllowed).To(BeTrue())
			})
		})
		Context("When org quota retrieve is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/quota_definitions/be829072-3137-418c-9607-c84e7d77e22a"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve org quota`),
					),
				)
			})

			It("should return error when org quota is retrieved", func() {
				myorgquota, err := newOrgQuotas(server.URL()).Get("be829072-3137-418c-9607-c84e7d77e22a")
				Expect(err).To(HaveOccurred())
				Expect(myorgquota).Should(BeNil())
			})
		})
	})

})

var _ = Describe("Org Quota by Name", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})
	Describe("FindByName()", func() {
		Context("Server return org quota by name", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/quota_definitions", "q=name:testorgquotaupdate"),
						ghttp.RespondWith(http.StatusOK, `{
							"total_results": 1,
							"total_pages": 1,
							"prev_url": null,
							"next_url": null,
							"resources": [
							  {
								"metadata": {
								  "guid": "be829072-3137-418c-9607-c84e7d77e22a",
								  "url": "/v2/quota_definitions/be829072-3137-418c-9607-c84e7d77e22a",
								  "created_at": "2014-05-30T19:07:52Z",
								  "updated_at": "2017-11-10T03:32:50Z"
								},
								"entity": {
								  "name": "testorgquotaupdate",
								  "non_basic_services_allowed": false,
								  "total_services": 10,
								  "total_routes": 500,
								  "total_private_domains": -1,
								  "memory_limit": 2048,
								  "trial_db_allowed": false,
								  "instance_memory_limit": 2048,
								  "app_instance_limit": -1,
								  "app_task_limit": -1,
								  "total_service_keys": -1,
								  "total_reserved_route_ports": 0
								}
							  }
							]
						  }`),
					),
				)
			})

			It("should return org quota by name", func() {
				myorgquota, err := newOrgQuotas(server.URL()).FindByName("testorgquotaupdate")
				Expect(err).NotTo(HaveOccurred())
				Expect(myorgquota).ShouldNot(BeNil())
				Expect(myorgquota.GUID).Should(Equal("be829072-3137-418c-9607-c84e7d77e22a"))
				Expect(myorgquota.Name).Should(Equal("testorgquotaupdate"))
				Expect(myorgquota.NonBasicServicesAllowed).To(BeFalse())
			})

		})

		Context("Server return no org quota by name", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/quota_definitions", "q=name:testorgquotaupdate"),
						ghttp.RespondWith(http.StatusOK, `{
															
						}`),
					),
				)
			})

			It("should return no org quota", func() {
				myorgquota, err := newOrgQuotas(server.URL()).FindByName("testorgquotaupdate")
				Expect(err).To(HaveOccurred())
				Expect(myorgquota).To(BeNil())
			})

		})
		Context("Server return error", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/quota_definitions", "q=name:testorgquotaupdate"),
						ghttp.RespondWith(http.StatusInternalServerError, `{
															
						}`),
					),
				)
			})

			It("should return no orgs", func() {
				myorgquota, err := newOrgQuotas(server.URL()).FindByName("testorgquotaupdate")
				Expect(err).To(HaveOccurred())
				Expect(myorgquota).To(BeNil())
			})

		})

	})
})

func newOrgQuotas(url string) OrgQuotas {
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

	return newOrgQuotasAPI(&client)
}
