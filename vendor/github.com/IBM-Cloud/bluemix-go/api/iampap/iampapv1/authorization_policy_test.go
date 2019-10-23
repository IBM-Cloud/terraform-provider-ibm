package iampapv1

import (
	"log"
	"net/http"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/client"
	"github.com/IBM-Cloud/bluemix-go/crn"
	"github.com/IBM-Cloud/bluemix-go/models"
	"github.com/IBM-Cloud/bluemix-go/session"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("AuthorizationPolicies", func() {
	var server *ghttp.Server
	var repo AuthorizationPolicyRepository

	BeforeEach(func() {
		server = ghttp.NewServer()
		repo = newTestAuthorizationPolicyRepo(server.URL())
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("List()", func() {
		Context("when there is no policy", func() {
			BeforeEach(func() {
				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodGet, "/acms/v2/accounts/account1-id/policies"),
					ghttp.RespondWith(http.StatusOK, `{}`),
				))
			})

			It("should return empty result", func() {
				policies, err := repo.List("account1-id", nil)
				Expect(err).Should(Succeed())
				Expect(policies).Should(BeEmpty())
			})
		})
		Context("when there is one authorization policy", func() {
			BeforeEach(func() {
				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodGet, "/acms/v2/accounts/account1-id/policies", "type=authorization"),
					ghttp.RespondWith(http.StatusOK, `{
						"policies": [
						{
							"id": "c723d05f-8f8c-4ef5-a44d-bc400cbde25e",
							"createdById": "IBMid-270006V8HD",
							"updatedById": "IBMid-270006V8HD",
							"roles": [
								{
									"crn": "crn:v1:bluemix:public:iam::::serviceRole:Reader",
									"id": "crn:v1:bluemix:public:iam::::serviceRole:Reader",
									"displayName": "Reader",
									"description": "As a reader, you can perform read-only actions within a service such as viewing service-specific resources."
								}
							],
							"resources": [
								{
									"serviceName": "kms-test",
									"accountId": "2bff70b5d2cc4b400814eca0bb730daa"
								}
							],
							"subjects": [
								{
									"serviceName": "cloud-object-storage",
									"accountId": "2bff70b5d2cc4b400814eca0bb730daa"
								}
							]
						}
					]
					}`),
				))
			})

			It("should return the policy", func() {
				policies, err := repo.List("account1-id", &AuthorizationPolicySearchQuery{
					Type: AuthorizationPolicyType,
				})
				Expect(err).Should(Succeed())
				Expect(policies).Should(HaveLen(1))
				Expect(policies[0].ID).To(Equal("c723d05f-8f8c-4ef5-a44d-bc400cbde25e"))
				Expect(policies[0].Roles[0].ID.String()).To(Equal("crn:v1:bluemix:public:iam::::serviceRole:Reader"))
				Expect(policies[0].Resources[0].AccountID).To(Equal("2bff70b5d2cc4b400814eca0bb730daa"))
				Expect(policies[0].Subjects[0].ServiceName).To(Equal("cloud-object-storage"))
			})
		})

		Context("when there are multiple policies", func() {
			BeforeEach(func() {
				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.RespondWith(http.StatusOK, `{
						"policies": [
						{
							"id": "c723d05f-8f8c-4ef5-a44d-bc400cbde25e",
							"createdById": "IBMid-270006V8HD",
							"updatedById": "IBMid-270006V8HD",
							"roles": [
								{
									"crn": "crn:v1:bluemix:public:iam::::serviceRole:Reader",
									"id": "crn:v1:bluemix:public:iam::::serviceRole:Reader",
									"displayName": "Reader",
									"description": "As a reader, you can perform read-only actions within a service such as viewing service-specific resources."
								}
							],
							"resources": [
								{
									"serviceName": "kms-test",
									"accountId": "2bff70b5d2cc4b400814eca0bb730daa"
								}
							],
							"subjects": [
								{
									"serviceName": "cloud-object-storage",
									"accountId": "2bff70b5d2cc4b400814eca0bb730daa"
								}
							]
						},
						{
							"id": "25afbf1b-d849-4d13-b3d9-a0768b5c29aa",
							"createdById": "IBMid-270006V8HD",
							"updatedById": "IBMid-270006V8HD",
							"roles": [
								{
									"crn": "crn:v1:bluemix:public:iam::::serviceRole:Administrator",
									"id": "crn:v1:bluemix:public:iam::::serviceRole:Administrator",
									"displayName": "Administrator",
									"description": "As an administrator, you can perform all platform actions based on the resource this role is being assigned, including assigning access policies to other users."
								}
							],
							"resources": [
								{
									"serviceName": "kms-test",
									"accountId": "2bff70b5d2cc4b400814eca0bb730daa"
								}
							],
							"subjects": [
								{
									"serviceName": "cloud-object-storage",
									"accountId": "2bff70b5d2cc4b400814eca0bb730daa"
								}
							]
						}
					]
					}`),
				))
			})

			It("should return all results", func() {
				policies, err := repo.List("account1-id", &AuthorizationPolicySearchQuery{
					Type: AccessPolicyType,
				})
				Expect(err).Should(Succeed())
				Expect(policies).Should(HaveLen(2))

				Expect(policies[0].ID).To(Equal("c723d05f-8f8c-4ef5-a44d-bc400cbde25e"))
				Expect(policies[0].Roles[0].ID.String()).To(Equal("crn:v1:bluemix:public:iam::::serviceRole:Reader"))
				Expect(policies[0].Resources[0].AccountID).To(Equal("2bff70b5d2cc4b400814eca0bb730daa"))
				Expect(policies[0].Subjects[0].ServiceName).To(Equal("cloud-object-storage"))

				Expect(policies[1].ID).To(Equal("25afbf1b-d849-4d13-b3d9-a0768b5c29aa"))
				Expect(policies[1].Roles[0].ID.String()).To(Equal("crn:v1:bluemix:public:iam::::serviceRole:Administrator"))
				Expect(policies[1].Resources[0].ServiceName).To(Equal("kms-test"))
				Expect(policies[1].Subjects[0].AccountID).To(Equal("2bff70b5d2cc4b400814eca0bb730daa"))
			})
		})

		Context("when there is error", func() {
			BeforeEach(func() {
				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.RespondWith(http.StatusBadRequest, `[{
						"StatusCode": 400,
						"code": "missing_token_in_request",
						"message": "Token not included in request. Please add token with authorization header."
					}]`),
				))
			})

			It("should return error", func() {
				_, err := repo.List("account1-id", &AuthorizationPolicySearchQuery{
					SubjectID:     "123",
					AccessGroupID: "567",
				})
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe("Get()", func() {
		Context("when policy is found", func() {
			BeforeEach(func() {
				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodGet, "/acms/v2/accounts/account1-id/policies/foo"),
					ghttp.RespondWith(http.StatusOK, `{
						"id": "c723d05f-8f8c-4ef5-a44d-bc400cbde25e",
						"createdById": "IBMid-270006V8HD",
						"updatedById": "IBMid-270006V8HD",
						"roles": [
							{
								"crn": "crn:v1:bluemix:public:iam::::serviceRole:Reader",
								"id": "crn:v1:bluemix:public:iam::::serviceRole:Reader",
								"displayName": "Reader",
								"description": "As a reader, you can perform read-only actions within a service such as viewing service-specific resources."
							}
						],
						"resources": [
							{
								"serviceName": "kms-test",
								"accountId": "2bff70b5d2cc4b400814eca0bb730daa"
							}
						],
						"subjects": [
							{
								"serviceName": "cloud-object-storage",
								"accountId": "2bff70b5d2cc4b400814eca0bb730daa"
							}
						]
					}`),
				))
			})

			It("should return success", func() {
				p, err := repo.Get("account1-id", "foo")
				Expect(err).Should(Succeed())
				Expect(p.ID).To(Equal("c723d05f-8f8c-4ef5-a44d-bc400cbde25e"))
				Expect(p.Roles[0].ID.String()).To(Equal("crn:v1:bluemix:public:iam::::serviceRole:Reader"))
				Expect(p.Resources[0].AccountID).To(Equal("2bff70b5d2cc4b400814eca0bb730daa"))
				Expect(p.Subjects[0].ServiceName).To(Equal("cloud-object-storage"))
			})
		})

		Context("when there is error fails", func() {
			BeforeEach(func() {
				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodGet, "/acms/v2/accounts/account1-id/policies/foo"),
					ghttp.RespondWith(http.StatusNotFound, `{
						"errors": [{
							"StatusCode": 404,
							"code": "not_found",
							"message": "policy not found"
						}],
					}`),
				))
			})

			It("should return error", func() {
				_, err := repo.Get("account1-id", "foo")
				Expect(err).Should(HaveOccurred())
			})
		})

	})

	Describe("Create", func() {
		Context("when creation succeeds", func() {
			BeforeEach(func() {
				header := http.Header{}
				header.Add("Etag", "abc")
				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodPost, "/acms/v2/accounts/account1-id/policies"),
					ghttp.VerifyJSON(`{
						"roles": [
							{
								"id": "crn:v1:bluemix:public:iam::::serviceRole:Reader",
								"displayName": "Reader",
								"description": "As a reader, you can perform read-only actions within a service such as viewing service-specific resources."
							}
						],
						"resources": [
							{
								"serviceName": "kms-test",
								"accountId": "2bff70b5d2cc4b400814eca0bb730daa"
							}
						],
						"subjects": [
							{
								"serviceName": "cloud-object-storage",
								"accountId": "2bff70b5d2cc4b400814eca0bb730daa"
							}
						]
					}`),
					ghttp.RespondWith(http.StatusCreated, `{
						"id": "c723d05f-8f8c-4ef5-a44d-bc400cbde25e",
						"createdById": "IBMid-270006V8HD",
						"updatedById": "IBMid-270006V8HD",
						"roles": [
							{
								"crn": "crn:v1:bluemix:public:iam::::serviceRole:Reader",
								"id": "crn:v1:bluemix:public:iam::::serviceRole:Reader",
								"displayName": "Reader",
								"description": "As a reader, you can perform read-only actions within a service such as viewing service-specific resources."
							}
						],
						"resources": [
							{
								"serviceName": "kms-test",
								"accountId": "2bff70b5d2cc4b400814eca0bb730daa"
							}
						],
						"subjects": [
							{
								"serviceName": "cloud-object-storage",
								"accountId": "2bff70b5d2cc4b400814eca0bb730daa"
							}
						]
					}`, header),
				))
			})
			It("should return created authorization policy", func() {
				p, err := repo.Create("account1-id", AuthorizationPolicy{
					Roles: []models.PolicyRole{
						models.PolicyRole{
							ID: crn.CRN{
								Scheme:          "crn",
								Version:         "v1",
								CName:           "bluemix",
								CType:           "public",
								ServiceName:     "iam",
								Region:          "",
								ScopeType:       "",
								Scope:           "",
								ServiceInstance: "",
								ResourceType:    "serviceRole",
								Resource:        "Reader",
							},
							DisplayName: "Reader",
							Description: "As a reader, you can perform read-only actions within a service such as viewing service-specific resources.",
						},
					},
					Resources: []models.PolicyResource{
						models.PolicyResource{
							ServiceName: "kms-test",
							AccountID:   "2bff70b5d2cc4b400814eca0bb730daa",
						},
					},
					Subjects: []models.PolicyResource{
						models.PolicyResource{
							ServiceName: "cloud-object-storage",
							AccountID:   "2bff70b5d2cc4b400814eca0bb730daa",
						},
					},
				})
				Expect(err).To(Succeed())
				Expect(p.ID).To(Equal("c723d05f-8f8c-4ef5-a44d-bc400cbde25e"))
			})
		})
		Context("when creation fails", func() {
			BeforeEach(func() {
				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodPost, "/acms/v2/accounts/account1-id/policies"),
					ghttp.VerifyJSON(`{
						"roles": [
							{
								"id": "crn:v1:bluemix:public:iam::::serviceRole:Administrator",
								"displayName": "Administrator",
								"description": "As an administrator, you can perform all platform actions based on the resource this role is being assigned, including assigning access policies to other users."
							}
						],
						"resources": [
							{
								"serviceName": "kms-test",
								"accountId": "2bff70b5d2cc4b400814eca0bb730daa"
							}
						],
						"subjects": [
							{
								"serviceName": "cloud-object-storage",
								"accountId": "2bff70b5d2cc4b400814eca0bb730daa"
							}
						]
					}`),
					ghttp.RespondWith(http.StatusUnauthorized, `{
						"errors": [{
							"StatusCode": 401,
							"code": "Invalid token",
							"message": "No groups found for the member test."
						}],
					}`),
				))
			})
			It("should return error", func() {
				_, err := repo.Create("account1-id", AuthorizationPolicy{
					Roles: []models.PolicyRole{
						models.PolicyRole{
							ID: crn.CRN{
								Scheme:          "crn",
								Version:         "v1",
								CName:           "bluemix",
								CType:           "public",
								ServiceName:     "iam",
								Region:          "",
								ScopeType:       "",
								Scope:           "",
								ServiceInstance: "",
								ResourceType:    "serviceRole",
								Resource:        "Administrator",
							},
							DisplayName: "Administrator",
							Description: "As an administrator, you can perform all platform actions based on the resource this role is being assigned, including assigning access policies to other users.",
						},
					},
					Resources: []models.PolicyResource{
						models.PolicyResource{
							ServiceName: "kms-test",
							AccountID:   "2bff70b5d2cc4b400814eca0bb730daa",
						},
					},
					Subjects: []models.PolicyResource{
						models.PolicyResource{
							ServiceName: "cloud-object-storage",
							AccountID:   "2bff70b5d2cc4b400814eca0bb730daa",
						},
					},
				})
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Update", func() {
		Context("when update succeeds", func() {
			BeforeEach(func() {
				header := http.Header{}
				header.Add("Etag", "abc")
				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodPut, "/acms/v2/accounts/account1-id/policies/foo"),
					ghttp.VerifyHeaderKV("If-Match", "abc"),
					ghttp.VerifyJSON(`{
						"roles": [
							{
								"id": "crn:v1:bluemix:public:iam::::serviceRole:Reader",
								"displayName": "Reader",
								"description": "As a reader, you can perform read-only actions within a service such as viewing service-specific resources."
							}
						],
						"resources": [
							{
								"serviceName": "kms-test",
								"accountId": "2bff70b5d2cc4b400814eca0bb730daa"
							}
						],
						"subjects": [
							{
								"serviceName": "cloud-object-storage",
								"accountId": "2bff70b5d2cc4b400814eca0bb730daa"
							}
						]
					}`),
					ghttp.RespondWith(http.StatusOK, `{
						"id": "c723d05f-8f8c-4ef5-a44d-bc400cbde25e",
						"createdById": "IBMid-270006V8HD",
						"updatedById": "IBMid-270006V8HD",
						"roles": [
							{
								"crn": "crn:v1:bluemix:public:iam::::serviceRole:Reader",
								"id": "crn:v1:bluemix:public:iam::::serviceRole:Reader",
								"displayName": "Reader",
								"description": "As a reader, you can perform read-only actions within a service such as viewing service-specific resources."
							}
						],
						"resources": [
							{
								"serviceName": "kms-test",
								"accountId": "2bff70b5d2cc4b400814eca0bb730daa"
							}
						],
						"subjects": [
							{
								"serviceName": "cloud-object-storage",
								"accountId": "2bff70b5d2cc4b400814eca0bb730daa"
							}
						]
					}`, header),
				))
			})
			It("should return updated authorization policy", func() {
				p, err := repo.Update("account1-id", "foo", AuthorizationPolicy{
					Roles: []models.PolicyRole{
						models.PolicyRole{
							ID: crn.CRN{
								Scheme:          "crn",
								Version:         "v1",
								CName:           "bluemix",
								CType:           "public",
								ServiceName:     "iam",
								Region:          "",
								ScopeType:       "",
								Scope:           "",
								ServiceInstance: "",
								ResourceType:    "serviceRole",
								Resource:        "Reader",
							},
							DisplayName: "Reader",
							Description: "As a reader, you can perform read-only actions within a service such as viewing service-specific resources.",
						},
					},
					Resources: []models.PolicyResource{
						models.PolicyResource{
							ServiceName: "kms-test",
							AccountID:   "2bff70b5d2cc4b400814eca0bb730daa",
						},
					},
					Subjects: []models.PolicyResource{
						models.PolicyResource{
							ServiceName: "cloud-object-storage",
							AccountID:   "2bff70b5d2cc4b400814eca0bb730daa",
						},
					},
				}, "abc")
				Expect(err).To(Succeed())
				Expect(p.ID).To(Equal("c723d05f-8f8c-4ef5-a44d-bc400cbde25e"))
				Expect(p.Version).To(Equal("abc"))
			})
		})
		Context("when update fails", func() {
			BeforeEach(func() {
				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodPut, "/acms/v2/accounts/account1-id/policies/foo"),
					ghttp.VerifyHeaderKV("If-Match", "abc"),
					ghttp.VerifyJSON(`{
						"roles": [
							{
								"id": "crn:v1:bluemix:public:iam::::serviceRole:Administrator",
								"displayName": "Administrator",
								"description": "As an administrator, you can perform all platform actions based on the resource this role is being assigned, including assigning access policies to other users."
							}
						],
						"resources": [
							{
								"serviceName": "kms-test",
								"accountId": "2bff70b5d2cc4b400814eca0bb730daa"
							}
						],
						"subjects": [
							{
								"serviceName": "cloud-object-storage",
								"accountId": "2bff70b5d2cc4b400814eca0bb730daa"
							}
						]
					}`),
					ghttp.RespondWith(http.StatusUnauthorized, `{
						"errors": [{
							"StatusCode": 401,
							"code": "Invalid token",
							"message": "No groups found for the member test."
						}],
					}`),
				))
			})
			It("should return error", func() {
				_, err := repo.Update("account1-id", "foo", AuthorizationPolicy{
					Roles: []models.PolicyRole{
						models.PolicyRole{
							ID: crn.CRN{
								Scheme:          "crn",
								Version:         "v1",
								CName:           "bluemix",
								CType:           "public",
								ServiceName:     "iam",
								Region:          "",
								ScopeType:       "",
								Scope:           "",
								ServiceInstance: "",
								ResourceType:    "serviceRole",
								Resource:        "Administrator",
							},
							DisplayName: "Administrator",
							Description: "As an administrator, you can perform all platform actions based on the resource this role is being assigned, including assigning access policies to other users.",
						},
					},
					Resources: []models.PolicyResource{
						models.PolicyResource{
							ServiceName: "kms-test",
							AccountID:   "2bff70b5d2cc4b400814eca0bb730daa",
						},
					},
					Subjects: []models.PolicyResource{
						models.PolicyResource{
							ServiceName: "cloud-object-storage",
							AccountID:   "2bff70b5d2cc4b400814eca0bb730daa",
						},
					},
				}, "abc")
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Delete()", func() {
		Context("when deletion succeeds", func() {
			BeforeEach(func() {
				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodDelete, "/acms/v1/policies/foo"),
					ghttp.RespondWith(http.StatusNoContent, ""),
				))
			})

			It("should return success", func() {
				err := repo.Delete("account1-id", "foo")
				Expect(err).Should(Succeed())
			})
		})

		Context("when deletion fails", func() {
			BeforeEach(func() {
				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodDelete, "/acms/v1/policies/foo"),
					ghttp.RespondWith(http.StatusNotFound, `{
						"errors": [{
							"StatusCode": 404,
							"code": "not_found",
							"message": "policy not found"
						}],
					}`),
				))
			})

			It("should return error", func() {
				err := repo.Delete("account1-id", "foo")
				Expect(err).Should(HaveOccurred())
				Expect(err.Error()).Should(ContainSubstring("not_found"))
			})
		})

	})
})

func newTestAuthorizationPolicyRepo(url string) AuthorizationPolicyRepository {
	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}
	conf := sess.Config.Copy()
	conf.Endpoint = &url
	client := client.Client{
		Config:      conf,
		ServiceName: bluemix.IAMPAPService,
	}
	return NewAuthorizationPolicyRepository(&client)
}
