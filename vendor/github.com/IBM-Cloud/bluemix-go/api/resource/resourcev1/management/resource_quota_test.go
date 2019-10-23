package management

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

var _ = Describe("QuotaDefinitions", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})

	Describe("List()", func() {
		Context("When there are multiple quota definitions", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/quota_definitions"),
						ghttp.RespondWith(http.StatusOK, `{
							"resources":[{
								"_id": "7ce89f4a-4381-4600-b814-3cd9a4f4bdf4",
								"_rev": "1-58b5e60b5bc8736565866f37bd413c1e",
								"name": "Trial Quota",
								"type": "SDQ",
								"number_of_service_instances": 3,
								"number_of_apps": 4,
								"instances_per_app": 4,
								"instance_memory": "1G",
								"total_app_memory": "1G",
								"vsi_limit": 1,
								"service_quotas": [
									{
										"_id": "65012f5b6fa84ecaaac5eab4abc2d0fd",
										"service_id": "rcf0c17db8-35ad-11e7-a919-92ebcb67fe33",
										"limit": 1
									},
									{
										"_id": "8b35eb93b6b14f50b6c4f3ff0dfd0358",
										"service_id": "rcd1a5a758-03a1-11e7-93ae-92361f002671",
										"limit": 1
									}
								],
								"created_at": "2017-05-16T17:12:52.925Z",
								"updated_at": "2017-05-16T17:12:52.925Z"
							},
							{
								"_id": "a3d7b8d01e261c24677937c29ab33f3c",
								"_rev": "1-cbdbf21c8f06da0ae8961e4abfd279e6",
								"name": "Pay-as-you-go Quota",
								"type": "SDQ",
								"number_of_service_instances": 3,
								"number_of_apps": 100,
								"instances_per_app": 32,
								"instance_memory": "512G",
								"total_app_memory": "512G",
								"vsi_limit": 100,
								"service_quotas": [],
								"created_at": "2017-05-16T17:12:52.925Z",
								"updated_at": "2017-05-16T17:12:52.925Z"
							}
						]}`),
					),
				)
			})
			It("should return all of them", func() {
				definitions, err := newTestQuotaDefinitionRepo(server.URL()).List()

				Expect(err).ShouldNot(HaveOccurred())
				Expect(definitions).Should(HaveLen(2))

				definition := definitions[0]
				Expect(definition.ID).Should(Equal("7ce89f4a-4381-4600-b814-3cd9a4f4bdf4"))
				Expect(definition.Revision).Should(Equal("1-58b5e60b5bc8736565866f37bd413c1e"))
				Expect(definition.Name).Should(Equal("Trial Quota"))
				Expect(definition.Type).Should(Equal("SDQ"))
				Expect(definition.ServiceInstanceCountLimit).Should(Equal(3))
				Expect(definition.AppCountLimit).Should(Equal(4))
				Expect(definition.AppInstanceCountLimit).Should(Equal(4))
				Expect(definition.AppInstanceMemoryLimit).Should(Equal("1G"))
				Expect(definition.TotalAppMemoryLimit).Should(Equal("1G"))
				Expect(definition.VSICountLimit).Should(Equal(1))
				Expect(definition.ServiceQuotas).Should(HaveLen(2))
				Expect(definition.CreatedAt).Should(Equal("2017-05-16T17:12:52.925Z"))
				Expect(definition.UpdatedAt).Should(Equal("2017-05-16T17:12:52.925Z"))
				serviceQuota := definition.ServiceQuotas[0]
				Expect(serviceQuota.ID).Should(Equal("65012f5b6fa84ecaaac5eab4abc2d0fd"))
				Expect(serviceQuota.ServiceID).Should(Equal("rcf0c17db8-35ad-11e7-a919-92ebcb67fe33"))
				Expect(serviceQuota.Limit).Should(Equal(1))
				serviceQuota = definition.ServiceQuotas[1]
				Expect(serviceQuota.ID).Should(Equal("8b35eb93b6b14f50b6c4f3ff0dfd0358"))
				Expect(serviceQuota.ServiceID).Should(Equal("rcd1a5a758-03a1-11e7-93ae-92361f002671"))
				Expect(serviceQuota.Limit).Should(Equal(1))

				definition = definitions[1]
				Expect(definition.ID).Should(Equal("a3d7b8d01e261c24677937c29ab33f3c"))
				Expect(definition.Revision).Should(Equal("1-cbdbf21c8f06da0ae8961e4abfd279e6"))
				Expect(definition.Name).Should(Equal("Pay-as-you-go Quota"))
				Expect(definition.Type).Should(Equal("SDQ"))
				Expect(definition.ServiceInstanceCountLimit).Should(Equal(3))
				Expect(definition.AppCountLimit).Should(Equal(100))
				Expect(definition.AppInstanceCountLimit).Should(Equal(32))
				Expect(definition.AppInstanceMemoryLimit).Should(Equal("512G"))
				Expect(definition.TotalAppMemoryLimit).Should(Equal("512G"))
				Expect(definition.VSICountLimit).Should(Equal(100))
				Expect(definition.ServiceQuotas).Should(BeEmpty())
				Expect(definition.CreatedAt).Should(Equal("2017-05-16T17:12:52.925Z"))
				Expect(definition.UpdatedAt).Should(Equal("2017-05-16T17:12:52.925Z"))
			})
		})

		Context("When there is one quota definition", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/quota_definitions"),
						ghttp.RespondWith(http.StatusOK, `{
							"resources":[{
								"_id": "a3d7b8d01e261c24677937c29ab33f3c",
								"_rev": "1-cbdbf21c8f06da0ae8961e4abfd279e6",
								"name": "Pay-as-you-go Quota",
								"type": "SDQ",
								"number_of_service_instances": 3,
								"number_of_apps": 100,
								"instances_per_app": 32,
								"instance_memory": "512G",
								"total_app_memory": "512G",
								"vsi_limit": 100,
								"service_quotas": [],
								"created_at": "2017-05-16T17:12:52.925Z",
								"updated_at": "2017-05-16T17:12:52.925Z"
							}
						]}`),
					),
				)
			})
			It("should return the only one", func() {
				definitions, err := newTestQuotaDefinitionRepo(server.URL()).List()

				Expect(err).ShouldNot(HaveOccurred())
				Expect(definitions).Should(HaveLen(1))

				definition := definitions[0]
				Expect(definition.ID).Should(Equal("a3d7b8d01e261c24677937c29ab33f3c"))
				Expect(definition.Revision).Should(Equal("1-cbdbf21c8f06da0ae8961e4abfd279e6"))
				Expect(definition.Name).Should(Equal("Pay-as-you-go Quota"))
				Expect(definition.Type).Should(Equal("SDQ"))
				Expect(definition.ServiceInstanceCountLimit).Should(Equal(3))
				Expect(definition.AppCountLimit).Should(Equal(100))
				Expect(definition.AppInstanceCountLimit).Should(Equal(32))
				Expect(definition.AppInstanceMemoryLimit).Should(Equal("512G"))
				Expect(definition.TotalAppMemoryLimit).Should(Equal("512G"))
				Expect(definition.VSICountLimit).Should(Equal(100))
				Expect(definition.ServiceQuotas).Should(BeEmpty())
				Expect(definition.CreatedAt).Should(Equal("2017-05-16T17:12:52.925Z"))
				Expect(definition.UpdatedAt).Should(Equal("2017-05-16T17:12:52.925Z"))
			})
		})

		Context("When there is no quota definition", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/quota_definitions"),
						ghttp.RespondWith(http.StatusOK, `{"resources":[]}`),
					),
				)
			})
			It("should return empty", func() {
				definitions, err := newTestQuotaDefinitionRepo(server.URL()).List()

				Expect(err).ShouldNot(HaveOccurred())
				Expect(definitions).Should(BeEmpty())
			})
		})

		Context("When there is backend error", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/quota_definitions"),
						ghttp.RespondWith(http.StatusUnauthorized, `{"Message":"Invalid Authorization"}`),
					),
				)
			})
			It("should return error", func() {
				definitions, err := newTestQuotaDefinitionRepo(server.URL()).List()

				Expect(err).Should(HaveOccurred())
				Expect(definitions).Should(BeEmpty())
			})
		})
	})

	Describe("FindByName()", func() {
		Context("When there are multiple quota definitions having the same name", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/quota_definitions"),
						ghttp.RespondWith(http.StatusOK, `{
							"resources":[{
								"_id": "7ce89f4a-4381-4600-b814-3cd9a4f4bdf4",
								"_rev": "1-58b5e60b5bc8736565866f37bd413c1e",
								"name": "Trial-Quota",
								"type": "SDQ",
								"number_of_service_instances": 3,
								"number_of_apps": 4,
								"instances_per_app": 4,
								"instance_memory": "1G",
								"total_app_memory": "1G",
								"vsi_limit": 1,
								"service_quotas": [
									{
										"_id": "65012f5b6fa84ecaaac5eab4abc2d0fd",
										"service_id": "rcf0c17db8-35ad-11e7-a919-92ebcb67fe33",
										"limit": 1
									},
									{
										"_id": "8b35eb93b6b14f50b6c4f3ff0dfd0358",
										"service_id": "rcd1a5a758-03a1-11e7-93ae-92361f002671",
										"limit": 1
									}
								],
								"created_at": "2017-05-16T17:12:52.925Z",
								"updated_at": "2017-05-16T17:12:52.925Z"
							},
							{
								"_id": "a3d7b8d01e261c24677937c29ab33f3c",
								"_rev": "1-cbdbf21c8f06da0ae8961e4abfd279e6",
								"name": "Trial-Quota",
								"type": "SDQ",
								"number_of_service_instances": 3,
								"number_of_apps": 100,
								"instances_per_app": 32,
								"instance_memory": "512G",
								"total_app_memory": "512G",
								"vsi_limit": 100,
								"service_quotas": [],
								"created_at": "2017-05-16T17:12:52.925Z",
								"updated_at": "2017-05-16T17:12:52.925Z"
							}
						]}`),
					),
				)
			})
			It("should return all of them", func() {
				definitions, err := newTestQuotaDefinitionRepo(server.URL()).FindByName("Trial-Quota")

				Expect(err).ShouldNot(HaveOccurred())
				Expect(definitions).Should(HaveLen(2))

				definition := definitions[0]
				Expect(definition.ID).Should(Equal("7ce89f4a-4381-4600-b814-3cd9a4f4bdf4"))
				Expect(definition.Revision).Should(Equal("1-58b5e60b5bc8736565866f37bd413c1e"))
				Expect(definition.Name).Should(Equal("Trial-Quota"))
				Expect(definition.Type).Should(Equal("SDQ"))
				Expect(definition.ServiceInstanceCountLimit).Should(Equal(3))
				Expect(definition.AppCountLimit).Should(Equal(4))
				Expect(definition.AppInstanceCountLimit).Should(Equal(4))
				Expect(definition.AppInstanceMemoryLimit).Should(Equal("1G"))
				Expect(definition.TotalAppMemoryLimit).Should(Equal("1G"))
				Expect(definition.VSICountLimit).Should(Equal(1))
				Expect(definition.ServiceQuotas).Should(HaveLen(2))
				Expect(definition.CreatedAt).Should(Equal("2017-05-16T17:12:52.925Z"))
				Expect(definition.UpdatedAt).Should(Equal("2017-05-16T17:12:52.925Z"))
				serviceQuota := definition.ServiceQuotas[0]
				Expect(serviceQuota.ID).Should(Equal("65012f5b6fa84ecaaac5eab4abc2d0fd"))
				Expect(serviceQuota.ServiceID).Should(Equal("rcf0c17db8-35ad-11e7-a919-92ebcb67fe33"))
				Expect(serviceQuota.Limit).Should(Equal(1))
				serviceQuota = definition.ServiceQuotas[1]
				Expect(serviceQuota.ID).Should(Equal("8b35eb93b6b14f50b6c4f3ff0dfd0358"))
				Expect(serviceQuota.ServiceID).Should(Equal("rcd1a5a758-03a1-11e7-93ae-92361f002671"))
				Expect(serviceQuota.Limit).Should(Equal(1))

				definition = definitions[1]
				Expect(definition.ID).Should(Equal("a3d7b8d01e261c24677937c29ab33f3c"))
				Expect(definition.Revision).Should(Equal("1-cbdbf21c8f06da0ae8961e4abfd279e6"))
				Expect(definition.Name).Should(Equal("Trial-Quota"))
				Expect(definition.Type).Should(Equal("SDQ"))
				Expect(definition.ServiceInstanceCountLimit).Should(Equal(3))
				Expect(definition.AppCountLimit).Should(Equal(100))
				Expect(definition.AppInstanceCountLimit).Should(Equal(32))
				Expect(definition.AppInstanceMemoryLimit).Should(Equal("512G"))
				Expect(definition.TotalAppMemoryLimit).Should(Equal("512G"))
				Expect(definition.VSICountLimit).Should(Equal(100))
				Expect(definition.ServiceQuotas).Should(BeEmpty())
				Expect(definition.CreatedAt).Should(Equal("2017-05-16T17:12:52.925Z"))
				Expect(definition.UpdatedAt).Should(Equal("2017-05-16T17:12:52.925Z"))
			})
		})

		Context("When there is one quota definition having the name", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/quota_definitions"),
						ghttp.RespondWith(http.StatusOK, `{
							"resources":[{
								"_id": "7ce89f4a-4381-4600-b814-3cd9a4f4bdf4",
								"_rev": "1-58b5e60b5bc8736565866f37bd413c1e",
								"name": "Trial Quota",
								"type": "SDQ",
								"number_of_service_instances": 3,
								"number_of_apps": 4,
								"instances_per_app": 4,
								"instance_memory": "1G",
								"total_app_memory": "1G",
								"vsi_limit": 1,
								"service_quotas": [
									{
										"_id": "65012f5b6fa84ecaaac5eab4abc2d0fd",
										"service_id": "rcf0c17db8-35ad-11e7-a919-92ebcb67fe33",
										"limit": 1
									},
									{
										"_id": "8b35eb93b6b14f50b6c4f3ff0dfd0358",
										"service_id": "rcd1a5a758-03a1-11e7-93ae-92361f002671",
										"limit": 1
									}
								],
								"created_at": "2017-05-16T17:12:52.925Z",
								"updated_at": "2017-05-16T17:12:52.925Z"
							},
							{
								"_id": "a3d7b8d01e261c24677937c29ab33f3c",
								"_rev": "1-cbdbf21c8f06da0ae8961e4abfd279e6",
								"name": "Pay-as-you-go Quota",
								"type": "SDQ",
								"number_of_service_instances": 3,
								"number_of_apps": 100,
								"instances_per_app": 32,
								"instance_memory": "512G",
								"total_app_memory": "512G",
								"vsi_limit": 100,
								"service_quotas": [],
								"created_at": "2017-05-16T17:12:52.925Z",
								"updated_at": "2017-05-16T17:12:52.925Z"
							}
						]}`),
					),
				)
			})
			It("should return the only one", func() {
				definitions, err := newTestQuotaDefinitionRepo(server.URL()).FindByName("Pay-as-you-go Quota")

				Expect(err).ShouldNot(HaveOccurred())
				Expect(definitions).Should(HaveLen(1))

				definition := definitions[0]
				Expect(definition.ID).Should(Equal("a3d7b8d01e261c24677937c29ab33f3c"))
				Expect(definition.Revision).Should(Equal("1-cbdbf21c8f06da0ae8961e4abfd279e6"))
				Expect(definition.Name).Should(Equal("Pay-as-you-go Quota"))
				Expect(definition.Type).Should(Equal("SDQ"))
				Expect(definition.ServiceInstanceCountLimit).Should(Equal(3))
				Expect(definition.AppCountLimit).Should(Equal(100))
				Expect(definition.AppInstanceCountLimit).Should(Equal(32))
				Expect(definition.AppInstanceMemoryLimit).Should(Equal("512G"))
				Expect(definition.TotalAppMemoryLimit).Should(Equal("512G"))
				Expect(definition.VSICountLimit).Should(Equal(100))
				Expect(definition.ServiceQuotas).Should(BeEmpty())
				Expect(definition.CreatedAt).Should(Equal("2017-05-16T17:12:52.925Z"))
				Expect(definition.UpdatedAt).Should(Equal("2017-05-16T17:12:52.925Z"))
			})
		})

		Context("When there is no quota definition having the same name", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/quota_definitions"),
						ghttp.RespondWith(http.StatusOK, `{"resources":[
							{
								"_id": "a3d7b8d01e261c24677937c29ab33f3c",
								"_rev": "1-cbdbf21c8f06da0ae8961e4abfd279e6",
								"name": "Pay-as-you-go Quota",
								"type": "SDQ",
								"number_of_service_instances": 3,
								"number_of_apps": 100,
								"instances_per_app": 32,
								"instance_memory": "512G",
								"total_app_memory": "512G",
								"vsi_limit": 100,
								"service_quotas": [],
								"created_at": "2017-05-16T17:12:52.925Z",
								"updated_at": "2017-05-16T17:12:52.925Z"
							}
						]}`),
					),
				)
			})
			It("should return empty", func() {
				definitions, err := newTestQuotaDefinitionRepo(server.URL()).FindByName("abc")

				Expect(err).Should(HaveOccurred())
				Expect(definitions).Should(BeEmpty())
			})
		})

		Context("When there is no quota definition returned", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/quota_definitions"),
						ghttp.RespondWith(http.StatusOK, `{"resources":[]}`),
					),
				)
			})
			It("should return empty", func() {
				definitions, err := newTestQuotaDefinitionRepo(server.URL()).FindByName("abc")

				Expect(err).Should(HaveOccurred())
				Expect(definitions).Should(BeEmpty())
			})
		})

		Context("When there is backend error", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/quota_definitions"),
						ghttp.RespondWith(http.StatusUnauthorized, `{"Message":"Invalid Authorization"}`),
					),
				)
			})
			It("should return error", func() {
				definitions, err := newTestQuotaDefinitionRepo(server.URL()).FindByName("test")

				Expect(err).Should(HaveOccurred())
				Expect(definitions).Should(BeEmpty())
			})
		})
	})

	Describe("Get()", func() {
		Context("When the quota definition exists", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/quota_definitions/a3d7b8d01e261c24677937c29ab33f3c"),
						ghttp.RespondWith(http.StatusOK, `{
							"name": "Pay-as-you-go Quota",
							"type": "SDQ",
							"number_of_service_instances": 3,
							"number_of_apps": 100,
							"instances_per_app": 32,
							"instance_memory": "512G",
							"total_app_memory": "512G",
							"vsi_limit": 100,
							"service_quotas": [
								{
									"_id": "65012f5b6fa84ecaaac5eab4abc2d0fd",
									"service_id": "rcf0c17db8-35ad-11e7-a919-92ebcb67fe33",
									"limit": 1
								},
								{
									"_id": "8b35eb93b6b14f50b6c4f3ff0dfd0358",
									"service_id": "rcd1a5a758-03a1-11e7-93ae-92361f002671",
									"limit": 1
								}
							],
							"created_at": "2017-05-16T17:12:52.925Z",
							"updated_at": "2017-05-16T17:12:52.925Z",
							"id": "a3d7b8d01e261c24677937c29ab33f3c"
						}`),
					),
				)
			})
			It("should return the quota definition", func() {
				quota, err := newTestQuotaDefinitionRepo(server.URL()).Get("a3d7b8d01e261c24677937c29ab33f3c")

				Expect(err).ShouldNot(HaveOccurred())
				Expect(quota).ShouldNot(BeNil())
				// Bug in BSS: https://github.ibm.com/BSS/resource-manager/issues/15
				//Expect(quota.ID).Should(Equal("a3d7b8d01e261c24677937c29ab33f3c"))
				//Expect(quota.Revision).Should(Equal("1-cbdbf21c8f06da0ae8961e4abfd279e6"))
				Expect(quota.Name).Should(Equal("Pay-as-you-go Quota"))
				Expect(quota.Type).Should(Equal("SDQ"))
				Expect(quota.ServiceInstanceCountLimit).Should(Equal(3))
				Expect(quota.AppCountLimit).Should(Equal(100))
				Expect(quota.AppInstanceCountLimit).Should(Equal(32))
				Expect(quota.AppInstanceMemoryLimit).Should(Equal("512G"))
				Expect(quota.TotalAppMemoryLimit).Should(Equal("512G"))
				Expect(quota.VSICountLimit).Should(Equal(100))
				Expect(quota.ServiceQuotas).Should(HaveLen(2))
				serviceQuota := quota.ServiceQuotas[0]
				Expect(serviceQuota.ID).Should(Equal("65012f5b6fa84ecaaac5eab4abc2d0fd"))
				Expect(serviceQuota.ServiceID).Should(Equal("rcf0c17db8-35ad-11e7-a919-92ebcb67fe33"))
				Expect(serviceQuota.Limit).Should(Equal(1))
				serviceQuota = quota.ServiceQuotas[1]
				Expect(serviceQuota.ID).Should(Equal("8b35eb93b6b14f50b6c4f3ff0dfd0358"))
				Expect(serviceQuota.ServiceID).Should(Equal("rcd1a5a758-03a1-11e7-93ae-92361f002671"))
				Expect(serviceQuota.Limit).Should(Equal(1))
				Expect(quota.CreatedAt).Should(Equal("2017-05-16T17:12:52.925Z"))
				Expect(quota.UpdatedAt).Should(Equal("2017-05-16T17:12:52.925Z"))
			})
		})

		Context("When the quota definition do not exist", func() {
			It("should return error", func() {

			})
		})

		Context("When there is backend error", func() {
			It("should return error", func() {

			})
		})
	})
})

func newTestQuotaDefinitionRepo(url string) ResourceQuotaRepository {
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

	return newResourceQuotaAPI(&client)
}
