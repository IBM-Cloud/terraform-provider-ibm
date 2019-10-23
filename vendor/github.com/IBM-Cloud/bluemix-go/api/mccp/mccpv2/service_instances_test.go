package mccpv2

import (
	"log"
	"net/http"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/client"
	"github.com/IBM-Cloud/bluemix-go/helpers"
	bluemixHttp "github.com/IBM-Cloud/bluemix-go/http"
	"github.com/IBM-Cloud/bluemix-go/session"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("ServiceInstances", func() {
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
						ghttp.VerifyRequest(http.MethodPost, "/v2/service_instances", "accepts_incomplete=true"),
						ghttp.VerifyBody([]byte(`{"name":"my-service-instance","space_guid":"ba013e75-1da1-4eaa-b30d-0f258211e4c1","service_plan_guid":"817b7e86-551c-416a-bfbc-c96feb4e4a64","parameters":{"the_service_broker":"wants this object"},"tags":["accounting","mongodb"]}`)),
						ghttp.RespondWith(http.StatusAccepted, `{
							 "metadata": {
							    "guid": "519b0d69-19e4-4009-a363-461eb117mccp32",
							    "url": "/v2/service_instances/519b0d69-19e4-4009-a363-461eb117mccp32",
							    "created_at": "2016-04-16T01:23:58Z",
							    "updated_at": null
							  },
							  "entity": {
							    "name": "my-service-instance",
							    "credentials": {							
							    },
							    "service_plan_guid": "817b7e86-551c-416a-bfbc-c96feb4e4a64",
							    "space_guid": "ba013e75-1da1-4eaa-b30d-0f258211e4c1",
							    "gateway_data": null,
							    "dashboard_url": null,
							    "type": "managed_service_instance",
							    "last_operation": {
							      "type": "create",
							      "state": "in progress",
							      "description": "",
							      "updated_at": null,
							      "created_at": "2016-04-16T01:23:58Z"
							    },
							    "tags": [
							      "accounting",
							      "mongodb"
							    ],
							    "space_url": "/v2/spaces/ba013e75-1da1-4eaa-b30d-0f258211e4c1",
							    "service_plan_url": "/v2/service_plans/817b7e86-551c-416a-bfbc-c96feb4e4a64",
							    "service_bindings_url": "/v2/service_instances/519b0d69-19e4-4009-a363-461eb117mccp32/service_bindings",
							    "service_keys_url": "/v2/service_instances/519b0d69-19e4-4009-a363-461eb117mccp32/service_keys",
							    "routes_url": "/v2/service_instances/519b0d69-19e4-4009-a363-461eb117mccp32/routes"
							  }						
						}`),
					),
				)
			})

			It("Should create ServiceInstances", func() {
				param := map[string]interface{}{"the_service_broker": "wants this object"}
				tags := []string{"accounting", "mongodb"}
				si, err := newServiceInstances(server.URL()).Create(ServiceInstanceCreateRequest{
					Name:      "my-service-instance",
					PlanGUID:  "817b7e86-551c-416a-bfbc-c96feb4e4a64",
					SpaceGUID: "ba013e75-1da1-4eaa-b30d-0f258211e4c1",
					Params:    param,
					Tags:      tags,
				})

				Expect(err).NotTo(HaveOccurred())
				Expect(si).NotTo(BeNil())
				Expect(si.Metadata.GUID).To(Equal("519b0d69-19e4-4009-a363-461eb117mccp32"))
				Expect(si.Entity.Name).To(Equal("my-service-instance"))
				Expect(si.Entity.ServicePlanGUID).To(Equal("817b7e86-551c-416a-bfbc-c96feb4e4a64"))
				Expect(si.Entity.SpaceGUID).To(Equal("ba013e75-1da1-4eaa-b30d-0f258211e4c1"))
			})
		})

		Context("When creation is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v2/service_instances"),
						ghttp.VerifyBody([]byte(`{"name":"my-service-instance","space_guid":"ba013e75-1da1-4eaa-b30d-0f258211e4c1","service_plan_guid":"817b7e86-551c-416a-bfbc-c96feb4e4a64"}`)),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to create`),
					),
				)
			})

			It("Should return error when created", func() {
				si, err := newServiceInstances(server.URL()).Create(ServiceInstanceCreateRequest{
					Name:      "my-service-instance",
					PlanGUID:  "817b7e86-551c-416a-bfbc-c96feb4e4a64",
					SpaceGUID: "ba013e75-1da1-4eaa-b30d-0f258211e4c1",
				})
				Expect(err).To(HaveOccurred())
				Expect(si).Should(BeNil())
			})
		})
	})

	Describe("FindByName", func() {
		Context("When there is match", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/service_instances", "return_user_provided_service_instances=true&q=name:foo"),
						ghttp.RespondWith(http.StatusOK, `{
							  "total_results": 1,
							  "total_pages": 1,
							  "prev_url": null,
							  "next_url": null,
							  "resources": [
							    {
							      "metadata": {
							        "guid": "76f3500e-60b2-4d8c-b478-f6797f2beedb",
							        "url": "/v2/service_instances/76f3500e-60b2-4d8c-b478-f6797f2beedb",
							        "created_at": "2016-04-16T01:23:58Z",
							        "updated_at": null
							      },
							      "entity": {
							        "name": "foo",
							        "credentials": {
							          "creds-key-39": "creds-val-39"
							        },
							        "service_plan_guid": "623943a0-96aa-4d96-bc34-1852c74020ac",
							        "space_guid": "259dffac-mccp36-433f-9738-58a3b1af5668",
							        "gateway_data": null,
							        "dashboard_url": null,
							        "type": "managed_service_instance",
							        "last_operation": {
							          "type": "create",
							          "state": "succeeded",
							          "description": "service broker-provided description",
							          "updated_at": "2016-04-16T01:23:58Z",
							          "created_at": "2016-04-16T01:23:58Z"
							        },
							        "tags": [
							          "accounting",
							          "mongodb"
							        ],
							        "space_url": "/v2/spaces/259dffac-mccp36-433f-9738-58a3b1af5668",
							        "service_plan_url": "/v2/service_plans/623943a0-96aa-4d96-bc34-1852c74020ac",
							        "service_bindings_url": "/v2/service_instances/76f3500e-60b2-4d8c-b478-f6797f2beedb/service_bindings",
							        "service_keys_url": "/v2/service_instances/76f3500e-60b2-4d8c-b478-f6797f2beedb/service_keys",
							        "routes_url": "/v2/service_instances/76f3500e-60b2-4d8c-b478-f6797f2beedb/routes"
							      }
							    }
							  ]
						}`),
					),
				)
			})

			It("Should return the match", func() {
				si, err := newServiceInstances(server.URL()).FindByName("foo")
				Expect(err).NotTo(HaveOccurred())
				Expect(si).NotTo(BeNil())
				Expect(si.GUID).To(Equal("76f3500e-60b2-4d8c-b478-f6797f2beedb"))
				Expect(si.Name).To(Equal("foo"))
				Expect(si.ServicePlanGUID).To(Equal("623943a0-96aa-4d96-bc34-1852c74020ac"))
				Expect(si.SpaceGUID).To(Equal("259dffac-mccp36-433f-9738-58a3b1af5668"))
			})
		})

		Context("When ServiceInstance FindByName is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/service_instances", "return_user_provided_service_instances=true&q=name:foo"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to find by name`),
					),
				)
			})

			It("Should return error when ServiceInstance is found by name", func() {
				si, err := newServiceInstances(server.URL()).FindByName("foo")
				Expect(err).To(HaveOccurred())
				Expect(si).Should(BeNil())
			})
		})
	})

	Describe("Get", func() {
		Context("When there is one ServiceInstance", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/service_instances/fbfd3591-48a8-482c-af62-f2d85d75229c"),
						ghttp.RespondWith(http.StatusOK, `{
							 "metadata": {
							    "guid": "fbfd3591-48a8-482c-af62-f2d85d75229c",
							    "url": "/v2/service_instances/fbfd3591-48a8-482c-af62-f2d85d75229c",
							    "created_at": "2016-04-16T01:23:58Z",
							    "updated_at": null
							  },
							  "entity": {
							    "name": "name-1014",
							    "credentials": {
							      "creds-key-42": "creds-val-42"
							    },
							    "service_plan_guid": "44f5e1a0-1a08-4099-96ac-f747cc9604f2",
							    "space_guid": "abdf35c0-5c2c-481b-8954-14602b7ce7c2",
							    "gateway_data": null,
							    "dashboard_url": null,
							    "type": "managed_service_instance",
							    "last_operation": {
							      "type": "create",
							      "state": "succeeded",
							      "description": "service broker-provided description",
							      "updated_at": "2016-04-16T01:23:58Z",
							      "created_at": "2016-04-16T01:23:58Z"
							    },
							    "tags": [
							      "accounting",
							      "mongodb"
							    ],
							    "space_url": "/v2/spaces/abdf35c0-5c2c-481b-8954-14602b7ce7c2",
							    "service_plan_url": "/v2/service_plans/44f5e1a0-1a08-4099-96ac-f747cc9604f2",
							    "service_bindings_url": "/v2/service_instances/fbfd3591-48a8-482c-af62-f2d85d75229c/service_bindings",
							    "service_keys_url": "/v2/service_instances/fbfd3591-48a8-482c-af62-f2d85d75229c/service_keys",
							    "routes_url": "/v2/service_instances/fbfd3591-48a8-482c-af62-f2d85d75229c/routes"
							  }
						}`),
					),
				)
			})

			It("Should return the ServiceInstance", func() {
				si, err := newServiceInstances(server.URL()).Get("fbfd3591-48a8-482c-af62-f2d85d75229c")
				Expect(err).NotTo(HaveOccurred())
				Expect(si).NotTo(BeNil())
				Expect(si.Metadata.GUID).To(Equal("fbfd3591-48a8-482c-af62-f2d85d75229c"))
				Expect(si.Entity.Name).To(Equal("name-1014"))
				Expect(si.Entity.ServicePlanGUID).To(Equal("44f5e1a0-1a08-4099-96ac-f747cc9604f2"))
				Expect(si.Entity.SpaceGUID).To(Equal("abdf35c0-5c2c-481b-8954-14602b7ce7c2"))
			})
		})

		Context("When ServiceInstance Get is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/service_instances/fbfd3591-48a8-482c-af62-f2d85d75229c"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to get`),
					),
				)
			})

			It("Should return error when ServiceInstance is retrieved", func() {
				si, err := newServiceInstances(server.URL()).Get("fbfd3591-48a8-482c-af62-f2d85d75229c")
				Expect(err).To(HaveOccurred())
				Expect(si).Should(BeNil())
			})
		})
	})

	Describe("Delete", func() {
		Context("When deletion is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v2/service_instances/8dc26f28-8b60-43b3-bed4-b9b5b0190d05"),
						ghttp.RespondWith(http.StatusAccepted, `{
								"metadata": {
								    "guid": "8dc26f28-8b60-43b3-bed4-b9b5b0190d05",
								    "url": "/v2/service_instances/8dc26f28-8b60-43b3-bed4-b9b5b0190d05",
								    "created_at": "2016-04-16T01:23:58Z",
								    "updated_at": null
								  },
								  "entity": {
								    "name": "name-1001",
								    "credentials": {
								      "creds-key-40": "creds-val-40"
								    },
								    "service_plan_guid": "03eb1037-5b45-4841-9bd3-4badcbc438e8",
								    "space_guid": "5db9fe9d-28ef-4704-a906-74f00d3dd0d9",
								    "gateway_data": null,
								    "dashboard_url": null,
								    "type": "managed_service_instance",
								    "last_operation": {
								      "type": "delete",
								      "state": "in progress",
								      "description": "",
								      "updated_at": "2016-04-16T01:23:58Z",
								      "created_at": "2016-04-16T01:23:58Z"
								    },
								    "tags": [
								      "accounting",
								      "mongodb"
								    ],
								    "space_url": "/v2/spaces/5db9fe9d-28ef-4704-a906-74f00d3dd0d9",
								    "service_plan_url": "/v2/service_plans/03eb1037-5b45-4841-9bd3-4badcbc438e8",
								    "service_bindings_url": "/v2/service_instances/8dc26f28-8b60-43b3-bed4-b9b5b0190d05/service_bindings",
								    "service_keys_url": "/v2/service_instances/8dc26f28-8b60-43b3-bed4-b9b5b0190d05/service_keys",
								    "routes_url": "/v2/service_instances/8dc26f28-8b60-43b3-bed4-b9b5b0190d05/routes"
								  }
						}`),
					),
				)
			})

			It("Should delete ServiceInstance", func() {
				err := newServiceInstances(server.URL()).Delete("8dc26f28-8b60-43b3-bed4-b9b5b0190d05")
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("When ServiceInstance Delete is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v2/service_instances/8dc26f28-8b60-43b3-bed4-b9b5b0190d05"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to delete`),
					),
				)
			})

			It("Should return error when deleted", func() {
				err := newServiceInstances(server.URL()).Delete("8dc26f28-8b60-43b3-bed4-b9b5b0190d05")
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Update", func() {
		Context("When update is succeesful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, "/v2/service_instances/e764af7b-1603-4ba3-b4bf-0b0da98f7ec2", "accepts_incomplete=true"),
						ghttp.VerifyBody([]byte(`{"name":"new-name","parameters":{"the_service_broker":"new service broker"}}`)),
						ghttp.RespondWith(http.StatusAccepted, `{
						  "metadata": {
						    "guid": "e764af7b-1603-4ba3-b4bf-0b0da98f7ec2",
						    "url": "/v2/service_instances/e764af7b-1603-4ba3-b4bf-0b0da98f7ec2",
						    "created_at": "2016-04-16T01:23:58Z",
						    "updated_at": null
						  },
						  "entity": {
						    "name": "new-name",
						    "credentials": {
						      "creds-key-41": "creds-val-41"
						    },
						    "service_plan_guid": "7dadd367-5603-4211-8986-d4d99dfab31c",
						    "space_guid": "0de51925-b333-4c84-9abc-115c917e68d5",
						    "gateway_data": null,
						    "dashboard_url": null,
						    "type": "managed_service_instance",
						    "last_operation": {
						      "type": "update",
						      "state": "in progress",
						      "description": "",
						      "updated_at": "2016-04-16T01:23:58Z",
						      "created_at": "2016-04-16T01:23:58Z"
						    },
						    "tags": [
						
						    ],
						    "space_url": "/v2/spaces/0de51925-b333-4c84-9abc-115c917e68d5",
						    "service_plan_url": "/v2/service_plans/7dadd367-5603-4211-8986-d4d99dfab31c",
						    "service_bindings_url": "/v2/service_instances/e764af7b-1603-4ba3-b4bf-0b0da98f7ec2/service_bindings",
						    "service_keys_url": "/v2/service_instances/e764af7b-1603-4ba3-b4bf-0b0da98f7ec2/service_keys",
						    "routes_url": "/v2/service_instances/e764af7b-1603-4ba3-b4bf-0b0da98f7ec2/routes"
						  }					
						}`),
					),
				)
			})

			It("Should update ServiceInstance", func() {
				param := map[string]interface{}{"the_service_broker": "new service broker"}
				si, err := newServiceInstances(server.URL()).Update("e764af7b-1603-4ba3-b4bf-0b0da98f7ec2", ServiceInstanceUpdateRequest{
					Name:   helpers.String("new-name"),
					Params: param,
				})

				Expect(err).NotTo(HaveOccurred())
				Expect(si).NotTo(BeNil())
				Expect(si.Metadata.GUID).To(Equal("e764af7b-1603-4ba3-b4bf-0b0da98f7ec2"))
				Expect(si.Entity.Name).To(Equal("new-name"))
				Expect(si.Entity.ServicePlanGUID).To(Equal("7dadd367-5603-4211-8986-d4d99dfab31c"))
				Expect(si.Entity.SpaceGUID).To(Equal("0de51925-b333-4c84-9abc-115c917e68d5"))
			})
		})

		Context("When ServiceInstance Update is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, "/v2/service_instances/e764af7b-1603-4ba3-b4bf-0b0da98f7ec2", "accepts_incomplete=true"),
						ghttp.VerifyBody([]byte(`{"name":"new-name"}`)),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to update`),
					),
				)
			})

			It("Should return error when updated", func() {
				si, err := newServiceInstances(server.URL()).Update(
					"e764af7b-1603-4ba3-b4bf-0b0da98f7ec2", ServiceInstanceUpdateRequest{
						Name: helpers.String("new-name"),
					})

				Expect(err).To(HaveOccurred())
				Expect(si).Should(BeNil())
			})
		})
	})
})

func newServiceInstances(url string) ServiceInstances {

	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}
	conf := sess.Config.Copy()
	conf.HTTPClient = bluemixHttp.NewHTTPClient(conf)
	conf.Endpoint = &url

	mccpClient := client.Client{
		Config:      conf,
		ServiceName: bluemix.MccpService,
	}

	return newServiceInstanceAPI(&mccpClient)
}
