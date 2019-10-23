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

var _ = Describe("Service Offering by Label", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})
	Describe("FindByLabel()", func() {
		Context("Server return service offering by label", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/services"),
						ghttp.RespondWith(http.StatusOK, `{
							"total_results": 1,
							"total_pages": 1,
							"prev_url": null,
							"next_url": null,
							"resources": [
								{
									"metadata": {
										"guid": "14c83ad2-6fd4-439a-8c3a-d1a20f8a2381",
										"url": "/v2/services/14c83ad2-6fd4-439a-8c3a-d1a20f8a2381",
										"created_at": "2014-06-03T07:04:12Z",
										"updated_at": "2017-03-17T20:05:06Z"
									},
									"entity": {
										"label": "cloudantNoSQLDB",
										"provider": null,
										"url": null,
										"description": "Cloudant NoSQL DB is a fully managed data layer designed for modern web and mobile applications that leverages a flexible JSON schema. Cloudant is built upon and compatible with Apache CouchDB and accessible through a secure HTTPS API, which scales as your application grows. Cloudant is ISO27001 and SOC2 Type 1 certified, and all data is stored in triplicate across separate physical nodes in a cluster for HA/DR within a data center.",
										"long_description": null,
										"version": null,
										"info_url": null,
										"active": true,
										"bindable": true,
										"unique_id": "cloudant",
										"tags": [
											"data_management",
											"ibm_created",
											"lite",
											"ibm_dedicated_public"
										],
										"requires": [

										],
										"documentation_url": null,
										"service_broker_guid": "b39770d9-5e57-4b3d-a5ff-0b1c7c432597",
										"plan_updateable": true,
										"service_plans_url": "/v2/services/14c83ad2-6fd4-439a-8c3a-d1a20f8a2381/service_plans"
									}
								}
							]							
															
						}`),
					),
				)
			})

			It("should return service offering by label", func() {
				myserviceoffering, err := newServiceOffering(server.URL()).FindByLabel("cloudantNoSQLDB")
				Expect(err).NotTo(HaveOccurred())
				Expect(myserviceoffering).ShouldNot(BeNil())
				Expect(myserviceoffering.GUID).Should(Equal("14c83ad2-6fd4-439a-8c3a-d1a20f8a2381"))
				Expect(myserviceoffering.Label).Should(Equal("cloudantNoSQLDB"))
			})

		})

		Context("Server return no space offering by label", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/services"),
						ghttp.RespondWith(http.StatusOK, `{
							"total_results": 0,
							"resources": [
							]
															
						}`),
					),
				)
			})

			It("should return no service offering", func() {
				myserviceoffering, err := newServiceOffering(server.URL()).FindByLabel("cloudantNoSQLDB")
				Expect(err).To(HaveOccurred())
				Expect(myserviceoffering).To(BeNil())
			})

		})
		Context("Server return error", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/services"),
						ghttp.RespondWith(http.StatusInternalServerError, `{
															
						}`),
					),
				)
			})

			It("should return error service offering", func() {
				myserviceoffering, err := newServiceOffering(server.URL()).FindByLabel("cloudantNoSQLDB")
				Expect(err).To(HaveOccurred())
				Expect(myserviceoffering).To(BeNil())
			})

		})

	})
})

func newServiceOffering(url string) ServiceOfferings {

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
	return newServiceOfferingAPI(&client)
}
