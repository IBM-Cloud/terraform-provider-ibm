package management

import (
	"log"
	"net/http"

	"github.com/IBM-Cloud/bluemix-go"

	"github.com/IBM-Cloud/bluemix-go/client"
	"github.com/IBM-Cloud/bluemix-go/models"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/onsi/gomega/ghttp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ResourceGroups", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})

	Describe("List()", func() {
		Context("When there is no user group", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/resource_groups"),
						ghttp.RespondWith(http.StatusOK, `{"resources":[]}`),
					),
				)
			})
			It("should return zero user group", func() {
				repo := newTestResourceGroupRepo(server.URL())
				groups, err := repo.List(&ResourceGroupQuery{})

				Expect(err).ShouldNot(HaveOccurred())
				Expect(groups).Should(BeEmpty())
			})
		})
		Context("When there is one user group", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/resource_groups"),
						ghttp.RespondWith(http.StatusOK, `{
							"resources": [{
								"id": "foo",
								"account_id": "abcdefg",
								"name": "test-group",
								"default": true,
								"state": "ACTIVE",
								"quota_id": "abcdefg",
								"payment_method_id": "payment1",
								"resource_linkages": []
							}]
						}`),
					),
				)
			})
			It("should return zero user group", func() {
				repo := newTestResourceGroupRepo(server.URL())
				groups, err := repo.List(&ResourceGroupQuery{})

				Expect(err).ShouldNot(HaveOccurred())

				Expect(groups).Should(HaveLen(1))
				group := groups[0]
				Expect(group.ID).Should(Equal("foo"))
				Expect(group.AccountID).Should(Equal("abcdefg"))
				Expect(group.Name).Should(Equal("test-group"))
				Expect(group.Default).Should(Equal(true))
				Expect(group.State).Should(Equal("ACTIVE"))
				Expect(group.QuotaID).Should(Equal("abcdefg"))
				Expect(group.PaymentMethodID).Should(Equal("payment1"))
				Expect(group.Linkages).Should(BeEmpty())
			})
		})

		Context("When there are multiple user groups", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/resource_groups"),
						ghttp.RespondWith(http.StatusOK, `{
							"resources": [{
								"id": "foo",
								"account_id": "abcdefg",
								"name": "test-group",
								"default": true,
								"state": "ACTIVE",
								"quota_id": "abcdefg",
								"payment_method_id": "payment1",
								"resource_linkages": []
							},{
								"id": "bar",
								"account_id": "xyz",
								"name": "test-group2",
								"default": false,
								"state": "SUSPENDED",
								"quota_id": "xyz",
								"payment_method_id": "payment2",
								"resource_linkages": [{
									"resource_id": "abc",
									"resource_origin": "CF_ORG"
								},{
									"resource_id": "def",
									"resource_origin": "IMS"
								}]
							}]
						}`),
					),
				)
			})
			It("should return all of them", func() {
				repo := newTestResourceGroupRepo(server.URL())
				groups, err := repo.List(&ResourceGroupQuery{})

				Expect(err).ShouldNot(HaveOccurred())

				Expect(groups).Should(HaveLen(2))
				group := groups[0]
				Expect(group.ID).Should(Equal("foo"))
				Expect(group.AccountID).Should(Equal("abcdefg"))
				Expect(group.Name).Should(Equal("test-group"))
				Expect(group.Default).Should(Equal(true))
				Expect(group.State).Should(Equal("ACTIVE"))
				Expect(group.QuotaID).Should(Equal("abcdefg"))
				Expect(group.PaymentMethodID).Should(Equal("payment1"))
				Expect(group.Linkages).Should(BeEmpty())

				group = groups[1]
				Expect(group.ID).Should(Equal("bar"))
				Expect(group.AccountID).Should(Equal("xyz"))
				Expect(group.Name).Should(Equal("test-group2"))
				Expect(group.Default).Should(Equal(false))
				Expect(group.State).Should(Equal("SUSPENDED"))
				Expect(group.QuotaID).Should(Equal("xyz"))
				Expect(group.PaymentMethodID).Should(Equal("payment2"))
				Expect(group.Linkages).Should(HaveLen(2))
				Expect(group.Linkages[0].ResourceID).Should(Equal("abc"))
				Expect(group.Linkages[0].ResourceOrigin.String()).Should(Equal("CF_ORG"))
				Expect(group.Linkages[1].ResourceID).Should(Equal("def"))
				Expect(group.Linkages[1].ResourceOrigin.String()).Should(Equal("IMS"))
			})
		})

		Context("Query by account ID", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/resource_groups", "account_id=abc"),
						ghttp.RespondWith(http.StatusOK, `{"resources":[]}`),
					),
				)
			})
			It("should get HTTP query 'accout_id'", func() {
				repo := newTestResourceGroupRepo(server.URL())
				groups, err := repo.List(&ResourceGroupQuery{
					AccountID: "abc",
				})

				Expect(err).ShouldNot(HaveOccurred())
				Expect(groups).Should(BeEmpty())
			})
		})

		Context("Query by default", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/resource_groups", "default=true"),
						ghttp.RespondWith(http.StatusOK, `{"resources":[]}`),
					),
				)
			})
			It("should get HTTP query 'accout_id'", func() {
				repo := newTestResourceGroupRepo(server.URL())
				groups, err := repo.List(&ResourceGroupQuery{
					Default: true,
				})

				Expect(err).ShouldNot(HaveOccurred())
				Expect(groups).Should(BeEmpty())
			})
		})

		Context("Query by resource ID and origin", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/resource_groups", "resource_id=abc&resource_origin=CF_ORG"),
						ghttp.RespondWith(http.StatusOK, `{"resources":[]}`),
					),
				)
			})
			It("should get HTTP query 'accout_id'", func() {
				repo := newTestResourceGroupRepo(server.URL())
				groups, err := repo.List(&ResourceGroupQuery{
					ResourceID:     "abc",
					ResourceOrigin: "CF_ORG",
				})

				Expect(err).ShouldNot(HaveOccurred())
				Expect(groups).Should(BeEmpty())
			})
		})

		Context("Query by multiple filters", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/resource_groups", "default=true&resource_id=abc&resource_origin=CF_ORG"),
						ghttp.RespondWith(http.StatusOK, `{"resources":[]}`),
					),
				)
			})
			It("should get HTTP query 'accout_id'", func() {
				repo := newTestResourceGroupRepo(server.URL())
				groups, err := repo.List(&ResourceGroupQuery{
					Default:        true,
					ResourceID:     "abc",
					ResourceOrigin: "CF_ORG",
				})

				Expect(err).ShouldNot(HaveOccurred())
				Expect(groups).Should(BeEmpty())
			})
		})

		Context("When there is backend error", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/resource_groups"),
						ghttp.RespondWith(http.StatusBadRequest, `{"resources":[]}`),
					),
				)
			})
			It("should return error", func() {
				repo := newTestResourceGroupRepo(server.URL())
				groups, err := repo.List(&ResourceGroupQuery{})

				Expect(err).Should(HaveOccurred())
				Expect(groups).Should(BeEmpty())
			})
		})
	})

	Describe("FindByName()", func() {
		Context("When no resource group returned", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/resource_groups"),
						ghttp.RespondWith(http.StatusOK, `{"resources":[]}`),
					),
				)
			})
			It("should return no resource group", func() {
				groups, err := newTestResourceGroupRepo(server.URL()).FindByName(&ResourceGroupQuery{}, "test")

				Expect(err).Should(HaveOccurred())
				Expect(groups).Should(BeEmpty())
			})
		})
		Context("When there is one user group returned having the same", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/resource_groups"),
						ghttp.RespondWith(http.StatusOK, `{
							"resources": [{
								"id": "foo",
								"account_id": "abcdefg",
								"name": "test-group",
								"default": true,
								"state": "ACTIVE",
								"quota_id": "abcdefg",
								"payment_method_id": "payment1",
								"resource_linkages": []
							}]
						}`),
					),
				)
			})
			It("should return that resource group", func() {
				groups, err := newTestResourceGroupRepo(server.URL()).FindByName(&ResourceGroupQuery{}, "test-group")

				Expect(err).ShouldNot(HaveOccurred())

				Expect(groups).Should(HaveLen(1))
				group := groups[0]
				Expect(group.ID).Should(Equal("foo"))
				Expect(group.AccountID).Should(Equal("abcdefg"))
				Expect(group.Name).Should(Equal("test-group"))
				Expect(group.Default).Should(Equal(true))
				Expect(group.State).Should(Equal("ACTIVE"))
				Expect(group.QuotaID).Should(Equal("abcdefg"))
				Expect(group.PaymentMethodID).Should(Equal("payment1"))
				Expect(group.Linkages).Should(BeEmpty())
			})
		})

		Context("When there are multiple resource groups having same name returned", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/resource_groups"),
						ghttp.RespondWith(http.StatusOK, `{
							"resources": [{
								"id": "foo",
								"account_id": "test-account",
								"name": "test-group",
								"default": true,
								"state": "ACTIVE",
								"quota_id": "abcdefg",
								"payment_method_id": "payment1",
								"resource_linkages": []
							},{
								"id": "bar",
								"account_id": "test-account2",
								"name": "test-group",
								"default": false,
								"state": "SUSPENDED",
								"quota_id": "xyz",
								"payment_method_id": "payment2",
								"resource_linkages": [{
									"resource_id": "abc",
									"resource_origin": "CF_ORG"
								},{
									"resource_id": "def",
									"resource_origin": "IMS"
								}]
							}]
						}`),
					),
				)
			})
			It("should return all of them", func() {
				groups, err := newTestResourceGroupRepo(server.URL()).FindByName(&ResourceGroupQuery{}, "test-group")

				Expect(err).ShouldNot(HaveOccurred())

				Expect(groups).Should(HaveLen(2))
				group := groups[0]
				Expect(group.ID).Should(Equal("foo"))
				Expect(group.AccountID).Should(Equal("test-account"))
				Expect(group.Name).Should(Equal("test-group"))
				Expect(group.Default).Should(Equal(true))
				Expect(group.State).Should(Equal("ACTIVE"))
				Expect(group.QuotaID).Should(Equal("abcdefg"))
				Expect(group.PaymentMethodID).Should(Equal("payment1"))
				Expect(group.Linkages).Should(BeEmpty())

				group = groups[1]
				Expect(group.ID).Should(Equal("bar"))
				Expect(group.AccountID).Should(Equal("test-account2"))
				Expect(group.Name).Should(Equal("test-group"))
				Expect(group.Default).Should(Equal(false))
				Expect(group.State).Should(Equal("SUSPENDED"))
				Expect(group.QuotaID).Should(Equal("xyz"))
				Expect(group.PaymentMethodID).Should(Equal("payment2"))
				Expect(group.Linkages).Should(HaveLen(2))
				Expect(group.Linkages[0].ResourceID).Should(Equal("abc"))
				Expect(group.Linkages[0].ResourceOrigin.String()).Should(Equal("CF_ORG"))
				Expect(group.Linkages[1].ResourceID).Should(Equal("def"))
				Expect(group.Linkages[1].ResourceOrigin.String()).Should(Equal("IMS"))
			})
		})

		Context("When there are multiple resource group returned, but none have that name", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/resource_groups", "account_id=abc"),
						ghttp.RespondWith(http.StatusOK, `{
							"resources": [{
								"id": "foo",
								"account_id": "abcdefg",
								"name": "test-group",
								"default": true,
								"state": "ACTIVE",
								"quota_id": "abcdefg",
								"payment_method_id": "payment1",
								"resource_linkages": []
							},{
								"id": "bar",
								"account_id": "xyz",
								"name": "test-group2",
								"default": false,
								"state": "SUSPENDED",
								"quota_id": "xyz",
								"payment_method_id": "payment2",
								"resource_linkages": [{
									"resource_id": "abc",
									"resource_origin": "CF_ORG"
								},{
									"resource_id": "def",
									"resource_origin": "IMS"
								}]
							}]
						}`),
					),
				)
			})
			It("should no resource group", func() {
				groups, err := newTestResourceGroupRepo(server.URL()).FindByName(&ResourceGroupQuery{
					AccountID: "abc",
				}, "foo")

				Expect(err).Should(HaveOccurred())
				Expect(groups).Should(BeEmpty())
			})
		})

		Context("When there is backend error", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/resource_groups"),
						ghttp.RespondWith(http.StatusBadRequest, `{"resources":[]}`),
					),
				)
			})
			It("should return error", func() {
				groups, err := newTestResourceGroupRepo(server.URL()).FindByName(&ResourceGroupQuery{}, "foo")

				Expect(err).Should(HaveOccurred())
				Expect(groups).Should(BeEmpty())
			})
		})
	})

	Describe("Create()", func() {
		Context("when creation is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/resource_groups"),
						ghttp.VerifyJSONRepresenting(models.ResourceGroup{
							Name:      "test",
							AccountID: "test-account-id",
							QuotaID:   "test-quota-id",
						}),
						ghttp.RespondWith(http.StatusOK, `{"id":"7f3f9f3ee8e64bf880ecec527c6f7c39"}`),
					),
				)
			})
			It("should return the new resource group", func() {
				group, err := newTestResourceGroupRepo(server.URL()).Create(models.ResourceGroup{
					Name:      "test",
					AccountID: "test-account-id",
					QuotaID:   "test-quota-id",
				})

				Expect(err).ShouldNot(HaveOccurred())
				Expect(group).ShouldNot(BeNil())
				Expect(group.ID).Should(Equal("7f3f9f3ee8e64bf880ecec527c6f7c39"))
			})
		})

		Context("when creation failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/resource_groups"),
						ghttp.VerifyJSONRepresenting(models.ResourceGroup{
							Name:      "test",
							AccountID: "test-account-id",
							QuotaID:   "test-quota-id",
						}),
						ghttp.RespondWith(http.StatusUnauthorized, `{"Message":"Invalid Authorization"}`),
					),
				)
			})
			It("should return error", func() {
				group, err := newTestResourceGroupRepo(server.URL()).Create(models.ResourceGroup{
					Name:      "test",
					AccountID: "test-account-id",
					QuotaID:   "test-quota-id",
				})

				Expect(err).To(HaveOccurred())
				Expect(group).To(BeNil())
			})
		})
	})

	Describe("Update()", func() {
		Context("when update is successful", func() {
			BeforeEach(func() {
				isDefault := new(bool)
				*isDefault = false
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPatch, "/v1/resource_groups/7f3f9f3ee8e64bf880ecec527c6f7c39"),
						ghttp.VerifyJSONRepresenting(ResourceGroupUpdateRequest{
							Name:    "test",
							QuotaID: "test-quota-id",
							Default: isDefault,
						}),
						ghttp.RespondWith(http.StatusOK, `{
							"id": "7f3f9f3ee8e64bf880ecec527c6f7c39",
							"account_id": "b8b618cc651496dd7a0634264d071843",
							"name": "test",
							"default": false,
							"state": "SUSPENDED",
							"quota_id": "test-quota-id",
							"quota_url": "/v1/quota_definitions/test-quota-id",
							"payment_methods_url": "/v1/resource_groups/7f3f9f3ee8e64bf880ecec527c6f7c39/payment_methods",
							"resource_linkages": [],
							"teams_url": "/v1/resource_groups/7f3f9f3ee8e64bf880ecec527c6f7c39/teams",
							"created_at": "2017-07-28T02:57:51.679Z",
							"updated_at": "2017-07-28T02:57:51.679Z"
						}`),
					),
				)
			})
			It("should return the updated resource group", func() {
				isDefault := new(bool)
				*isDefault = false
				group, err := newTestResourceGroupRepo(server.URL()).Update("7f3f9f3ee8e64bf880ecec527c6f7c39", &ResourceGroupUpdateRequest{
					Name:    "test",
					QuotaID: "test-quota-id",
					Default: isDefault,
				})

				Expect(err).ShouldNot(HaveOccurred())
				Expect(group).ShouldNot(BeNil())
				// TODO: BSS bug
				// Expect(group.ID).Should(Equal("bar"))
				Expect(group.AccountID).Should(Equal("b8b618cc651496dd7a0634264d071843"))
				Expect(group.Name).Should(Equal("test"))
				Expect(group.Default).Should(Equal(false))
				Expect(group.State).Should(Equal("SUSPENDED"))
				Expect(group.QuotaID).Should(Equal("test-quota-id"))
				Expect(group.Linkages).Should(HaveLen(0))
			})
		})

		Context("when not updating `default`", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPatch, "/v1/resource_groups/7f3f9f3ee8e64bf880ecec527c6f7c39"),
						ghttp.VerifyJSONRepresenting(models.ResourceGroup{
							Name:    "test",
							QuotaID: "test-quota-id",
						}),
						ghttp.RespondWith(http.StatusOK, `{
							"id": "7f3f9f3ee8e64bf880ecec527c6f7c39",
							"account_id": "b8b618cc651496dd7a0634264d071843",
							"name": "test",
							"default": true,
							"state": "SUSPENDED",
							"quota_id": "test-quota-id",
							"quota_url": "/v1/quota_definitions/test-quota-id",
							"payment_methods_url": "/v1/resource_groups/7f3f9f3ee8e64bf880ecec527c6f7c39/payment_methods",
							"resource_linkages": [],
							"teams_url": "/v1/resource_groups/7f3f9f3ee8e64bf880ecec527c6f7c39/teams",
							"created_at": "2017-07-28T02:57:51.679Z",
							"updated_at": "2017-07-28T02:57:51.679Z"
						}`),
					),
				)
			})
			It("should return the updated resource group", func() {
				group, err := newTestResourceGroupRepo(server.URL()).Update("7f3f9f3ee8e64bf880ecec527c6f7c39", &ResourceGroupUpdateRequest{
					Name:    "test",
					QuotaID: "test-quota-id",
				})

				Expect(err).ShouldNot(HaveOccurred())
				Expect(group).ShouldNot(BeNil())
				// TODO: BSS bug
				// Expect(group.ID).Should(Equal("bar"))
				Expect(group.AccountID).Should(Equal("b8b618cc651496dd7a0634264d071843"))
				Expect(group.Name).Should(Equal("test"))
				Expect(group.Default).Should(Equal(true))
				Expect(group.State).Should(Equal("SUSPENDED"))
				Expect(group.QuotaID).Should(Equal("test-quota-id"))
				Expect(group.Linkages).Should(HaveLen(0))
			})
		})

		Context("when update failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPatch, "/v1/resource_groups/7f3f9f3ee8e64bf880ecec527c6f7c39"),
						ghttp.VerifyJSONRepresenting(models.ResourceGroup{
							QuotaID: "test-quota-id",
						}),
						ghttp.RespondWith(http.StatusUnauthorized, `{"Message":"Invalid Authorization"}`),
					),
				)
			})
			It("should return error", func() {
				group, err := newTestResourceGroupRepo(server.URL()).Update("7f3f9f3ee8e64bf880ecec527c6f7c39", &ResourceGroupUpdateRequest{
					QuotaID: "test-quota-id",
				})

				Expect(err).To(HaveOccurred())
				Expect(group).To(BeNil())
			})
		})
	})
	Describe("Delete()", func() {
		Context("When deletion is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/resource_groups/abc"),
						ghttp.RespondWith(http.StatusNoContent, ``),
					),
				)
			})
			It("should return success", func() {
				err := newTestResourceGroupRepo(server.URL()).Delete("abc")
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("When deletion failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/resource_groups/abc"),
						ghttp.RespondWith(http.StatusNotFound, `{"message":"Not found"}`),
					),
				)
			})
			It("should return error", func() {
				err := newTestResourceGroupRepo(server.URL()).Delete("abc")
				Expect(err).Should(HaveOccurred())
			})
		})
	})
})

func newTestResourceGroupRepo(url string) ResourceGroupRepository {
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

	return newResourceGroupAPI(&client)
}
