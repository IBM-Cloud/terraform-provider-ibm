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

var _ = Describe("Service Plan by Label", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})
	Describe("FindPlanInServiceOffering()", func() {
		Context("Server return service plan", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/service_plans"),
						ghttp.RespondWith(http.StatusOK, `{
							"total_results": 2,
							"total_pages": 1,
							"prev_url": null,
							"next_url": null,
							"resources": [
								{
									"metadata": {
										"guid": "e72c6030-mccpe3-4477-9fb1-ca2b0408cbcb",
										"url": "/v2/service_plans/e72c6030-mccpe3-4477-9fb1-ca2b0408cbcb",
										"created_at": "2016-09-08T12:55:17Z",
										"updated_at": "2017-03-17T20:05:06Z"
									},
									"entity": {
										"name": "Lite",
										"free": true,
										"description": "The Lite plan provides access to the full functionality of Cloudant for development and evaluation. The plan has a set amount of provisioned throughput capacity as shown and includes a max of 1GB of encrypted data storage.",
										"service_guid": "14c83ad2-6fd4-439a-8c3a-d1a20f8a2381",
										"unique_id": "cloudant-lite",
										"public": true,
										"active": true,
										"service_url": "/v2/services/14c83ad2-6fd4-439a-8c3a-d1a20f8a2381",
										"service_instances_url": "/v2/service_plans/e72c6030-mccpe3-4477-9fb1-ca2b0408cbcb/service_instances"
									}
								},
								{
									"metadata": {
										"guid": "f5b75238-51ea-4f52-a520-caec9015a68d",
										"url": "/v2/service_plans/f5b75238-51ea-4f52-a520-caec9015a68d",
										"created_at": "2016-09-08T12:55:17Z",
										"updated_at": "2017-03-17T20:05:06Z"
									},
									"entity": {
										"name": "Standard",
										"free": false,
										"description": "The Standard plan provides access to the full functionality of Cloudant and can scale as needed for all use cases. The provisioned throughput capacity starts at 100 lookups/sec, 50 writes/sec, and 5 queries/sec with three additional tiers of capacity that can be toggled in the Cloudant Dashboard to meet application throughput requirements. The Standard plan includes 20GB of encrypted data storage, with additional storage metered for purchase.",
										"service_guid": "14c83ad2-6fd4-439a-8c3a-d1a20f8a2381",
										"unique_id": "cloudant-standard",
										"public": true,
										"active": true,
										"service_url": "/v2/services/14c83ad2-6fd4-439a-8c3a-d1a20f8a2381",
										"service_instances_url": "/v2/service_plans/f5b75238-51ea-4f52-a520-caec9015a68d/service_instances"
									}
								}
							]

						}`),
					),
				)
			})

			It("should return service plan", func() {
				myserviceplan, err := newServicePlan(server.URL()).FindPlanInServiceOffering("14c83ad2-6fd4-439a-8c3a-d1a20f8a2381", "Lite")
				Expect(err).NotTo(HaveOccurred())
				Expect(myserviceplan).ShouldNot(BeNil())
				Expect(myserviceplan.GUID).Should(Equal("e72c6030-mccpe3-4477-9fb1-ca2b0408cbcb"))
				Expect(myserviceplan.Name).Should(Equal("Lite"))
			})

		})

		Context("Server return no space plan", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/service_plans"),
						ghttp.RespondWith(http.StatusOK, `{
							"total_results": 0,
							"resources": [
							]
															
						}`),
					),
				)
			})

			It("should return no service plan", func() {
				myserviceplan, err := newServicePlan(server.URL()).FindPlanInServiceOffering("14c83ad2-6fd4-439a-8c3a-d1a20f8a2381", "Lite")
				Expect(err).To(HaveOccurred())
				Expect(myserviceplan).To(BeNil())
			})

		})
		Context("Server return error", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/service_plans"),
						ghttp.RespondWith(http.StatusInternalServerError, `{
															
						}`),
					),
				)
			})

			It("should return error service plan", func() {
				myserviceplan, err := newServicePlan(server.URL()).FindPlanInServiceOffering("14c83ad2-6fd4-439a-8c3a-d1a20f8a2381", "Lite")
				Expect(err).To(HaveOccurred())
				Expect(myserviceplan).To(BeNil())
			})

		})

	})
})

func newServicePlan(url string) ServicePlans {

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

	return newServicePlanAPI(&client)
}
