package controller

import (
	"log"
	"net/http"

	"github.com/IBM-Cloud/bluemix-go"

	"github.com/IBM-Cloud/bluemix-go/client"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/onsi/gomega/ghttp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ServiceInstances", func() {
	targetCRN := "crn:v1:d_att288:dedicated::us-south::::d_att288-us-south"
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})

	Describe("ListInstances()", func() {
		Context("When there is no service instance", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/resource_instances"),
						ghttp.RespondWith(http.StatusOK, `{"resources":[]}`),
					),
				)
			})
			It("should return zero service instance", func() {
				repo := newTestServiceInstanceRepo(server.URL())
				instances, err := repo.ListInstances(ServiceInstanceQuery{
					ResourceGroupID: "resource_group_id",
				})

				Expect(err).ShouldNot(HaveOccurred())
				Expect(instances).Should(BeEmpty())
			})
		})
		Context("When there is one service instance", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/resource_instances"),
						ghttp.RespondWith(http.StatusOK, `{
																	"rows_count":1,
																	"resources":[{
																		"id":"foo",
																		"guid":"a83261db-5cd1-46ae-8cfb-8ebbcc7c0184",
																		"url":"/v1/resource_instances/a83261db-5cd1-46ae-8cfb-8ebbcc7c0184",
																		"created_at":"2017-07-31T06:19:45.16112535Z",
																		"updated_at":null,
																		"deleted_at":null,
																		"name":"test-instance",
																		"target_crn":"crn:v1:d_att288:dedicated::us-south::::d_att288-us-south",
																		"account_id":"560df2058b1e7c402303cc598b3e5540",
																		"resource_plan_id":"rc-pb-28c24fccc-4ca6-4ddd-a3be-7746cdce9912",
																		"resource_group_id":"resource_group_id",
																		"create_time":0,"crn":"",
																		"state":"inactive",
																		"type":"service_instance",
																		"resource_id":"fake-resource-id",
																		"dashboard_url":null,
																		"last_operation":null,
																		"account_url":"/v1/accounts/560df2058b1e7c402303cc598b3e5540",
																		"resource_plan_url":"/v1/catalog/regions/ibm:ys1:us-south/plans/rc-pb-28c24fccc-4ca6-4ddd-a3be-7746cdce9912",
																		"resource_bindings_url":"/v1/resource_instances/a83261db-5cd1-46ae-8cfb-8ebbcc7c0184/resource_bindings",
																		"resource_aliases_url":"/v1/resource_instances/a83261db-5cd1-46ae-8cfb-8ebbcc7c0184/resource_aliases",
																		"siblings_url":"/v1/resource_instances/a83261db-5cd1-46ae-8cfb-8ebbcc7c0184/siblings"}]}`),
					),
				)
			})
			It("should return one service instance", func() {
				repo := newTestServiceInstanceRepo(server.URL())
				instances, err := repo.ListInstances(ServiceInstanceQuery{
					ResourceGroupID: "resource_group_id",
				})

				Expect(err).ShouldNot(HaveOccurred())

				Expect(instances).Should(HaveLen(1))
				instance := instances[0]
				Expect(instance.ID).Should(Equal("foo"))
				Expect(instance.ServiceID).Should(Equal("fake-resource-id"))
				Expect(instance.Name).Should(Equal("test-instance"))
				Expect(instance.State).Should(Equal("inactive"))
				Expect(instance.Type).Should(Equal("service_instance"))
			})
		})

		Context("When there are multiple service instances", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/resource_instances"),
						ghttp.RespondWith(http.StatusOK, `{
						"rows_count":3,
						"resources":[{
							"id":"foo",
							"guid":"a83261db-5cd1-46ae-8cfb-8ebbcc7c0184",
							"url":"/v1/resource_instances/a83261db-5cd1-46ae-8cfb-8ebbcc7c0184",
							"created_at":"2017-07-31T06:19:45.16112535Z",
							"updated_at":null,
							"deleted_at":null,
							"name":"test-instance",
							"target_crn":"crn:v1:d_att288:dedicated::us-south::::d_att288-us-south",
							"account_id":"560df2058b1e7c402303cc598b3e5540",
							"resource_plan_id":"rc-pb-28c24fccc-4ca6-4ddd-a3be-7746cdce9912",
							"resource_group_id":"",
							"create_time":0,"crn":"",
							"state":"active",
							"type":"service_instance",
							"resource_id":"fake-resource-id",
							"dashboard_url":null,
							"last_operation":null,
							"account_url":"/v1/accounts/560df2058b1e7c402303cc598b3e5540",
							"resource_plan_url":"/v1/catalog/regions/ibm:ys1:us-south/plans/rc-pb-28c24fccc-4ca6-4ddd-a3be-7746cdce9912",
							"resource_bindings_url":"/v1/resource_instances/a83261db-5cd1-46ae-8cfb-8ebbcc7c0184/resource_bindings",
							"resource_aliases_url":"/v1/resource_instances/a83261db-5cd1-46ae-8cfb-8ebbcc7c0184/resource_aliases",
							"siblings_url":"/v1/resource_instances/a83261db-5cd1-46ae-8cfb-8ebbcc7c0184/siblings"
						},
						{
							"id":"foo1",
							"guid":"dea23694-1a2c-45e4-bfa8-7be5226d998c",
							"url":"/v1/resource_instances/dea23694-1a2c-45e4-bfa8-7be5226d998c",
							"created_at":"2017-07-31T06:20:14.592704474Z",
							"updated_at":null,
							"deleted_at":null,
							"name":"test-instance1",
							"target_crn":"crn:v1:d_att288:dedicated::us-south::::d_att288-us-south",
							"account_id":"560df2058b1e7c402303cc598b3e5540",
							"resource_plan_id":"rc-pb-28c24fccc-4ca6-4ddd-a3be-7746cdce9912",
							"resource_group_id":"",
							"create_time":0,
							"crn":"",
							"state":"inactive",
							"type":"service_instance",
							"resource_id":"fake-resource-id1",
							"dashboard_url":null,
							"last_operation":null,
							"account_url":"/v1/accounts/560df2058b1e7c402303cc598b3e5540",
							"resource_plan_url":"/v1/catalog/regions/ibm:ys1:us-south/plans/rc-pb-28c24fccc-4ca6-4ddd-a3be-7746cdce9912",
							"resource_bindings_url":"/v1/resource_instances/dea23694-1a2c-45e4-bfa8-7be5226d998c/resource_bindings",
							"resource_aliases_url":"/v1/resource_instances/dea23694-1a2c-45e4-bfa8-7be5226d998c/resource_aliases",
							"siblings_url":"/v1/resource_instances/dea23694-1a2c-45e4-bfa8-7be5226d998c/siblings"
						},
						{
							"id":"foo2",
							"guid":"50312f63-f43b-4a67-aa32-8626da609adb",
							"url":"/v1/resource_instances/50312f63-f43b-4a67-aa32-8626da609adb",
							"created_at":"2017-07-31T06:27:46.215093281Z",
							"updated_at":"2017-07-31T07:34:07.740506169Z",
							"deleted_at":null,
							"name":"test-instance2",
							"target_crn":"crn:v1:d_att288:dedicated::us-south::::d_att288-us-south",
							"account_id":"560df2058b1e7c402303cc598b3e5540",
							"resource_plan_id":"rc-pb-28c24fccc-4ca6-4ddd-a3be-7746cdce9912",
							"resource_group_id":"",
							"create_time":0,
							"crn":"",
							"state":"active",
							"type":"service_instance",
							"resource_id":"fake-resource-id2",
							"dashboard_url":null,
							"last_operation":null,
							"account_url":"/v1/accounts/560df2058b1e7c402303cc598b3e5540",
							"resource_plan_url":"/v1/catalog/regions/ibm:ys1:us-south/plans/rc-pb-28c24fccc-4ca6-4ddd-a3be-7746cdce9912",
							"resource_bindings_url":"/v1/resource_instances/50312f63-f43b-4a67-aa32-8626da609adb/resource_bindings",
							"resource_aliases_url":"/v1/resource_instances/50312f63-f43b-4a67-aa32-8626da609adb/resource_aliases",
							"siblings_url":"/v1/resource_instances/50312f63-f43b-4a67-aa32-8626da609adb/siblings"
						}
						]}`),
					),
				)
			})
			It("should return all of them", func() {
				repo := newTestServiceInstanceRepo(server.URL())
				instances, err := repo.ListInstances(ServiceInstanceQuery{
					ResourceGroupID: "resource_group_id",
				})

				Expect(err).ShouldNot(HaveOccurred())

				Expect(instances).Should(HaveLen(3))
				instance := instances[0]
				Expect(instance.ID).Should(Equal("foo"))
				Expect(instance.ServiceID).Should(Equal("fake-resource-id"))
				Expect(instance.Name).Should(Equal("test-instance"))
				Expect(instance.State).Should(Equal("active"))
				Expect(instance.Type).Should(Equal("service_instance"))

				instance = instances[1]
				Expect(instance.ID).Should(Equal("foo1"))
				Expect(instance.ServiceID).Should(Equal("fake-resource-id1"))
				Expect(instance.Name).Should(Equal("test-instance1"))
				Expect(instance.State).Should(Equal("inactive"))
				Expect(instance.Type).Should(Equal("service_instance"))

				instance = instances[2]
				Expect(instance.ID).Should(Equal("foo2"))
				Expect(instance.ServiceID).Should(Equal("fake-resource-id2"))
				Expect(instance.Name).Should(Equal("test-instance2"))
				Expect(instance.State).Should(Equal("active"))
				Expect(instance.Type).Should(Equal("service_instance"))
			})
		})

	})

	Describe("CreateInstance()", func() {
		Context("when creation is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/resource_instances"),
						ghttp.VerifyJSONRepresenting(CreateServiceInstanceRequest{
							Name:            "test-instance-name",
							TargetCrn:       targetCRN,
							ServicePlanID:   "test-resource-plan-id",
							ResourceGroupID: "test-resource-group-id",
							Tags:            []string{},
						}),
						ghttp.RespondWith(http.StatusOK,
							`{
								"id":"crn:v1:staging:public:automation-test:us-south:a/560df2058b1e7c402303cc598b3e5540:ec8b9112-331a-4883-be22-6a023542f4da::",
								"guid":"",
								"url":"/v1/resource_instances/3f50ed1d-070a-4114-baef-bccdc93b48c0",
								"created_at":"2017-08-01T02:32:16.12857946Z",
								"updated_at":null,
								"deleted_at":null,
								"name":"test-instance-name",
								"target_crn":"crn:v1:d_att288:dedicated::us-south::::d_att288-us-south",
								"account_id":"test-account-id",
								"resource_plan_id":"test-resource-plan-id",
								"resource_group_id":"test-resource-group-id",
								"create_time":0,
								"crn":"crn:v1:staging:public:automation-test:us-south:a/560df2058b1e7c402303cc598b3e5540:ec8b9112-331a-4883-be22-6a023542f4da::",
								"state":"active",
								"type":"service_instance",
								"resource_id":"rcdemo21bff2d3d-0872-4dd7-affd-80aed1bb46a5",
								"dashboard_url":"http://rc-performance-test-blue.stage1.ng.bluemix.net/cfs/dashboard/3f50ed1d-070a-4114-baef-bccdc93b48c0",
								"last_operation":null,"account_url":"/v1/accounts/560df2058b1e7c402303cc598b3e5540",
								"resource_plan_url":"/v1/catalog/regions/ibm:ys1:us-south/plans/rcdemo29c8c9fb5-ea26-400e-8ef0-12cd49fd2240",
								"resource_bindings_url":"/v1/resource_instances/3f50ed1d-070a-4114-baef-bccdc93b48c0/resource_bindings",
								"resource_aliases_url":"/v1/resource_instances/3f50ed1d-070a-4114-baef-bccdc93b48c0/resource_aliases",
								"siblings_url":"/v1/resource_instances/3f50ed1d-070a-4114-baef-bccdc93b48c0/siblings"}`),
					),
				)
			})
			It("should return the new service instance", func() {
				instance, err := newTestServiceInstanceRepo(server.URL()).CreateInstance(
					CreateServiceInstanceRequest{
						Name:            "test-instance-name",
						TargetCrn:       targetCRN,
						ServicePlanID:   "test-resource-plan-id",
						ResourceGroupID: "test-resource-group-id",
						Tags:            []string{},
					})

				Expect(err).ShouldNot(HaveOccurred())
				Expect(instance).ShouldNot(BeNil())
				Expect(instance.ID).Should(Equal("crn:v1:staging:public:automation-test:us-south:a/560df2058b1e7c402303cc598b3e5540:ec8b9112-331a-4883-be22-6a023542f4da::"))
			})
		})

		Context("when creation failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/resource_instances"),
						ghttp.VerifyJSONRepresenting(CreateServiceInstanceRequest{
							Name:            "test-instance-name",
							TargetCrn:       targetCRN,
							ServicePlanID:   "test-resource-plan-id",
							ResourceGroupID: "test-resource-group-id",
							Tags:            []string{},
						}),
						ghttp.RespondWith(http.StatusBadRequest, `400 Failed to get the region information for the given region_guid`),
					),
				)
			})
			It("should return error", func() {
				_, err := newTestServiceInstanceRepo(server.URL()).CreateInstance(CreateServiceInstanceRequest{
					Name:            "test-instance-name",
					TargetCrn:       targetCRN,
					ServicePlanID:   "test-resource-plan-id",
					ResourceGroupID: "test-resource-group-id",
					Tags:            []string{},
				})
				Expect(err).To(HaveOccurred())

			})
		})
	})

	Describe("UpdateInstance()", func() {
		Context("when update is successful", func() {
			BeforeEach(func() {
				isDefault := new(bool)
				*isDefault = false
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPatch, "/v1/resource_instances/3f50ed1d-070a-4114-baef-bccdc93b48c0"),
						ghttp.VerifyJSONRepresenting(UpdateServiceInstanceRequest{
							Name:          "new-test-name",
							ServicePlanID: "test-resource-id2",
						}),
						ghttp.RespondWith(http.StatusOK, `
						{"id":"crn:v1:staging:public:automation-test:us-south:a/560df2058b1e7c402303cc598b3e5540:ec8b9112-331a-4883-be22-6a023542f4da::",
						"guid":"",
						"url":"/v1/resource_instances/281ef5c5-bdb2-4fc4-b05a-7126bf1d2923",
						"created_at":"2017-07-31T07:02:16.308084742Z",
						"updated_at":"2017-08-01T02:54:52.188612634Z",
						"deleted_at":null,
						"name":"new-test-name",
						"target_crn":"crn:v1:d_att288:dedicated::us-south::::d_att288-us-south",
						"account_id":"560df2058b1e7c402303cc598b3e5540",
						"resource_plan_id":"test-resource-id2",
						"resource_group_id":"",
						"create_time":0,
						"crn":"crn:v1:staging:public:automation-test:us-south:a/560df2058b1e7c402303cc598b3e5540:ec8b9112-331a-4883-be22-6a023542f4da::",
						"state":"active",
						"type":"service_instance",
						"resource_id":"rcdemo21bff2d3d-0872-4dd7-affd-80aed1bb46a5",
						"dashboard_url":"http://rc-performance-test-blue.stage1.ng.bluemix.net/cfs/dashboard/281ef5c5-bdb2-4fc4-b05a-7126bf1d2923",
						"last_operation":null,
						"account_url":"/v1/accounts/560df2058b1e7c402303cc598b3e5540",
						"resource_plan_url":"/v1/catalog/regions/ibm:ys1:us-south/plans/rcdemo29c8c9fb5-ea26-400e-8ef0-12cd49fd2240",
						"resource_bindings_url":"/v1/resource_instances/281ef5c5-bdb2-4fc4-b05a-7126bf1d2923/resource_bindings",
						"resource_aliases_url":"/v1/resource_instances/281ef5c5-bdb2-4fc4-b05a-7126bf1d2923/resource_aliases",
						"siblings_url":"/v1/resource_instances/281ef5c5-bdb2-4fc4-b05a-7126bf1d2923/siblings"}`),
					),
				)
			})
			It("should return the updated service instance", func() {
				isDefault := new(bool)
				*isDefault = false
				instance, err := newTestServiceInstanceRepo(server.URL()).UpdateInstance("3f50ed1d-070a-4114-baef-bccdc93b48c0", UpdateServiceInstanceRequest{
					Name:          "new-test-name",
					ServicePlanID: "test-resource-id2",
				})

				Expect(err).ShouldNot(HaveOccurred())
				Expect(instance).ShouldNot(BeNil())
				Expect(instance.Name).Should(Equal("new-test-name"))
				Expect(instance.Crn.String()).Should(Equal("crn:v1:staging:public:automation-test:us-south:a/560df2058b1e7c402303cc598b3e5540:ec8b9112-331a-4883-be22-6a023542f4da::"))
				Expect(instance.ServicePlanID).Should(Equal("test-resource-id2"))
			})
		})
	})
	Describe("Delete()", func() {
		Context("When deletion is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/resource_instances/8cbdc8aa-fc59-4e15-9020-c509e7726346"),
						ghttp.RespondWith(http.StatusNoContent, ``),
					),
				)
			})
			It("should return success", func() {
				err := newTestServiceInstanceRepo(server.URL()).DeleteInstance("8cbdc8aa-fc59-4e15-9020-c509e7726346", false)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("When deletion failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/resource_instances/abc"),
						ghttp.RespondWith(http.StatusNotFound, `{"message":"Not found"}`),
					),
				)
			})
			It("should return error", func() {
				err := newTestServiceInstanceRepo(server.URL()).DeleteInstance("abc", false)
				Expect(err).Should(HaveOccurred())
			})
		})
	})
})

func newTestServiceInstanceRepo(url string) ResourceServiceInstanceRepository {

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

	return newResourceServiceInstanceAPI(&client)
}
