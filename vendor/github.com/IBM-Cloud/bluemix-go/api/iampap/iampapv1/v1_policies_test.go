package iampapv1

import (
	"log"
	"net/http"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/client"
	"github.com/IBM-Cloud/bluemix-go/session"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("V1Policies", func() {
	var server *ghttp.Server
	var repo V1PolicyRepository

	BeforeEach(func() {
		server = ghttp.NewServer()
		repo = newv1Policy(server.URL())
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("List()", func() {
		Context("When there are policies", func() {
			BeforeEach(func() {
				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodGet, "/v1/policies", "account_id=account1-id&type=access"),
					ghttp.RespondWith(http.StatusOK, `{
						"policies": [
						  {
							"id": "285a9542-c08c-43c6-8f0a-632dc5801b64",
							"type": "access",
							"subjects": [{
								"attributes": [{
									"name": "iam_id",
									"value": "IBMid-060001MNR3"
								}]
							}],
							"roles": [
							  {
								"role_id": "crn:v1:bluemix:public:iam::::role:Editor",
								"display_name": "Editor",
								"description": "As an editor, you can perform all platform actions except for managing the account and assigning access policies."
							  }
							],
							"resources": [{
								"attributes": [{
									"name": "accountId",
									"value": "b8b618cc651496dd7a0634264d071843"
								},{
									"name": "resourceType",
									"value": "resource-group"
								},{
									"name": "resource",
									"value": "247bd6f43b5f7af30b43a13724020b52"
								}]
							}],
							"href": "https://iam.cloud.ibm.com/v1/policies/285a9542-c08c-43c6-8f0a-632dc5801b64",
							"created_at": "2018-04-17T06:53:46.342Z",
							"created_by_id": "IBMid-270004WA4U",
							"last_modified_at": "2018-04-17T06:53:46.342Z",
							"last_modified_by_id": "2018-04-17T06:53:46.342Z"
						  },
						  {
							"id": "1a369e5f-3339-4fbb-b52f-97ff364c9e7a",
							"type": "access",
							"subjects": [{
								"attributes": [{
									"name": "iam_id",
									"value": "IBMid-2700076QDK"
								}]
							}],
							"roles": [
							  {
								"role_id": "crn:v1:bluemix:public:iam::::role:Editor",
								"display_name": "Editor",
								"description": "As an editor, you can perform all platform actions except for managing the account and assigning access policies."
							  }
							],
							"resources": [{
								"attributes": [{
									"name": "accountId",
									"value": "b8b618cc651496dd7a0634264d071843"
								}]
							}],
							"href": "https://iam.cloud.ibm.com/v1/policies/1a369e5f-3339-4fbb-b52f-97ff364c9e7a",
							"created_at": null,
							"created_by_id": null,
							"last_modified_at": null,
							"last_modified_by_id": null
						  },
						  {
							"id": "cecc91ff-b680-4705-ba1c-c7f1fefcf02d",
							"type": "access",
							"subjects": [{
								"attributes": [{
									"name": "access_group_id",
									"value": "AccessGroupId-888620a2-be96-47f9-a214-4ae995ef17ad"
								}]
							}],
							"roles": [
							  {
								"role_id": "crn:v1:bluemix:public:iam::::role:Viewer",
								"display_name": "Viewer",
								"description": "As a viewer, you can view service instances, but you can't modify them."
							  }
							],
							"resources": [{
								"attributes": [{
									"name": "accountId",
									"value": "b8b618cc651496dd7a0634264d071843"
								},{
									"name": "resourceType",
									"value": "resource-group"
								},{
									"name": "resource",
									"value": "7f3f9f3ee8e64bf880ecec527c6f7c39"
								}]
							}],
							"href": "https://iam.cloud.ibm.com/v1/policies/cecc91ff-b680-4705-ba1c-c7f1fefcf02d",
							"created_at": "2018-03-26T09:07:03.853Z",
							"created_by_id": "IBMid-270004WA4U",
							"last_modified_at": "2018-03-26T09:07:03.853Z",
							"last_modified_by_id": "IBMid-270004WA4U"
						  }
						]
					}`),
				))
			})
			It("should return all policies", func() {
				policies, err := repo.List(SearchParams{
					AccountID: "account1-id",
					Type:      AccessPolicyType,
				})
				Expect(err).ShouldNot(HaveOccurred())
				Expect(policies).To(HaveLen(3))
				Expect(policies[0].ID).To(Equal("285a9542-c08c-43c6-8f0a-632dc5801b64"))
				Expect(policies[0].Subjects[0].IAMID()).To(Equal("IBMid-060001MNR3"))
				Expect(policies[0].Resources[0].AccountID()).To(Equal("b8b618cc651496dd7a0634264d071843"))
				Expect(policies[0].Resources[0].ResourceType()).To(Equal("resource-group"))
				Expect(policies[0].Resources[0].Resource()).To(Equal("247bd6f43b5f7af30b43a13724020b52"))
				Expect(policies[0].Roles[0].RoleID).To(Equal("crn:v1:bluemix:public:iam::::role:Editor"))

				Expect(policies[1].ID).To(Equal("1a369e5f-3339-4fbb-b52f-97ff364c9e7a"))
				Expect(policies[1].Subjects[0].IAMID()).To(Equal("IBMid-2700076QDK"))
				Expect(policies[1].Resources[0].AccountID()).To(Equal("b8b618cc651496dd7a0634264d071843"))
				Expect(policies[1].Resources[0].ResourceType()).To(Equal(""))
				Expect(policies[1].Resources[0].Resource()).To(Equal(""))
				Expect(policies[0].Roles[0].RoleID).To(Equal("crn:v1:bluemix:public:iam::::role:Editor"))

				Expect(policies[2].ID).To(Equal("cecc91ff-b680-4705-ba1c-c7f1fefcf02d"))
				Expect(policies[2].Subjects[0].AccessGroupID()).To(Equal("AccessGroupId-888620a2-be96-47f9-a214-4ae995ef17ad"))
				Expect(policies[2].Resources[0].AccountID()).To(Equal("b8b618cc651496dd7a0634264d071843"))
				Expect(policies[2].Resources[0].ResourceType()).To(Equal("resource-group"))
				Expect(policies[2].Resources[0].Resource()).To(Equal("7f3f9f3ee8e64bf880ecec527c6f7c39"))
				Expect(policies[2].Roles[0].RoleID).To(Equal("crn:v1:bluemix:public:iam::::role:Viewer"))
			})
		})

		Context("When there is no policy", func() {
			BeforeEach(func() {
				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodGet, "/v1/policies", "account_id=account1-id&type=access&iam_id=foo"),
					ghttp.RespondWith(http.StatusOK, `{
						"policies": []
					}`),
				))
			})
			It("should return empty", func() {
				policies, err := repo.List(SearchParams{
					AccountID: "account1-id",
					IAMID:     "foo",
					Type:      AccessPolicyType,
				})
				Expect(err).ShouldNot(HaveOccurred())
				Expect(policies).To(HaveLen(0))
			})
		})

		Context("When there is error", func() {
			BeforeEach(func() {
				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodGet, "/v1/policies", "account_id=account1-id&type=access&access_group_id=foo"),
					ghttp.RespondWith(http.StatusUnauthorized, `{
						"trace": "aac0333a26c64caaa5528bfeed3de4b8",
						"errors": [
							{
								"code": "invalid_token",
								"message": "The provided IAM token is invalid."
							}
						],
						"status_code": 401
					}`),
				))
			})
			It("should return empty", func() {
				policies, err := repo.List(SearchParams{
					AccountID:     "account1-id",
					AccessGroupID: "foo",
					Type:          AccessPolicyType,
				})
				Expect(err).Should(HaveOccurred())
				Expect(policies).To(HaveLen(0))
			})
		})
	})

	Describe("Get()", func() {
		Context("When the policy is found", func() {
			BeforeEach(func() {
				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodGet, "/v1/policies/policy-id"),
					ghttp.RespondWith(http.StatusOK, `{
						"id": "policy-id",
						"type": "access",
						"subjects": [{
							"attributes": [{
								"name": "iam_id",
								"value": "IBMid-270006V8HD"
							}]
						}],
						"roles": [
							{
								"role_id": "crn:v1:bluemix:public:iam::::role:Viewer",
								"display_name": "Viewer",
								"description": "As a viewer, you can view service instances, but you can't modify them."
							}
						],
						"resources": [{
							"attributes": [{
								"name": "resourceType",
								"value": "resource-group"
							},{
								"name": "resource",
								"value": "5e0f4f2166a04667aea9440ad4850232"
							}]
						}],
						"href": "https://iam.cloud.ibm.com/v1/policies/0b3cea25-fc07-49b1-986b-90846085cfeb",
						"created_at": "2017-11-10T04:37:44.943Z",
						"created_by_id": "IBMid-310000JVN5",
						"last_modified_at": "2017-11-10T04:37:44.943Z",
						"last_modified_by_id": "IBMid-310000JVN5"
				}`, http.Header{"ETag": []string{"3-a283be20686c2c71c512e73e9ee2ce9c"}}),
				))
			})
			It("should return the policy", func() {
				p, err := repo.Get("policy-id")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(p.ID).To(Equal("policy-id"))
				Expect(p.Type).To(Equal(AccessPolicyType))
				Expect(p.Subjects[0].IAMID()).To(Equal("IBMid-270006V8HD"))
				Expect(p.Roles[0].RoleID).To(Equal("crn:v1:bluemix:public:iam::::role:Viewer"))
				Expect(p.Resources[0].Resource()).To(Equal("5e0f4f2166a04667aea9440ad4850232"))
				Expect(p.Href).To(Equal("https://iam.cloud.ibm.com/v1/policies/0b3cea25-fc07-49b1-986b-90846085cfeb"))
				Expect(p.CreatedAt).To(Equal("2017-11-10T04:37:44.943Z"))
				Expect(p.CreatedByID).To(Equal("IBMid-310000JVN5"))
				Expect(p.LastModifiedAt).To(Equal("2017-11-10T04:37:44.943Z"))
				Expect(p.LastModifiedByID).To(Equal("IBMid-310000JVN5"))
				Expect(p.Version).To(Equal("3-a283be20686c2c71c512e73e9ee2ce9c"))
			})
		})

		Context("When the policy is not found", func() {
			BeforeEach(func() {
				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodGet, "/v1/policies/policy-id"),
					ghttp.RespondWith(http.StatusNotFound, `{
						"trace": "10cdeb89a930420a83f8d295e5dad851",
						"errors": [
							{
								"code": "policy_not_found",
								"message": "Policy with Id 285a9542-c08c-43c6-8f0a-632dc5801b64 not found."
							}
						],
						"status_code": 404
					}`),
				))
			})
			It("should return not found error", func() {
				p, err := repo.Get("policy-id")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("policy_not_found"))
				Expect(p).To(BeZero())
			})
		})

		Context("When there is error", func() {
			BeforeEach(func() {
				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodGet, "/v1/policies/policy-id"),
					ghttp.RespondWith(http.StatusUnauthorized, `{
						"trace": "aac0333a26c64caaa5528bfeed3de4b8",
						"errors": [
							{
								"code": "invalid_token",
								"message": "The provided IAM token is invalid."
							}
						],
						"status_code": 401
					}`),
				))
			})
			It("should return the error", func() {
				p, err := repo.Get("policy-id")
				Expect(err).To(HaveOccurred())
				Expect(p).To(BeZero())
			})
		})
	})

	Describe("Delete()", func() {
		Context("Delete successfully", func() {
			BeforeEach(func() {
				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodDelete, "/v1/policies/policy-id"),
					ghttp.RespondWith(http.StatusNoContent, nil),
				))
			})
			It("should return the error", func() {
				err := repo.Delete("policy-id")
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("Deletion fails", func() {
			BeforeEach(func() {
				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodDelete, "/v1/policies/policy-id"),
					ghttp.RespondWith(http.StatusUnauthorized, `{
						"trace": "aac0333a26c64caaa5528bfeed3de4b8",
						"errors": [
							{
								"code": "invalid_token",
								"message": "The provided IAM token is invalid."
							}
						],
						"status_code": 401
					}`),
				))
			})
			It("should return the error", func() {
				err := repo.Delete("policy-id")
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Create()", func() {
		Context("Create successfully", func() {
			BeforeEach(func() {
				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodPost, "/v1/policies"),
					ghttp.VerifyJSON(`{
						"type": "access",
						"subjects": [{
							"attributes": [{
								"name": "iam_id",
								"value": "IBMid-270006V8HD"
							}]
						}],
						"roles": [{
							"role_id": "crn:v1:bluemix:public:iam::::role:Viewer"
						}],
						"resources": [{
							"attributes": [{
								"name": "resourceType",
								"value": "resource-group"
							},{
								"name": "resource",
								"value": "5e0f4f2166a04667aea9440ad4850232"

							}]
						}]
					}`),
					ghttp.RespondWith(http.StatusOK, `{
						"id": "policy-id",
						"type": "access",
						"subjects": [{
							"attributes": [{
								"name": "iam_id",
								"value": "IBMid-270006V8HD"
							}]
						}],
						"roles": [{
							"role_id": "crn:v1:bluemix:public:iam::::role:Viewer",
							"display_name": "Viewer",
							"description": "As a viewer, you can view service instances, but you can't modify them."
						}],
						"resources": [{
							"attributes": [{
								"name": "resourceType",
								"value": "resource-group"
							},{
								"name": "resource",
								"value": "5e0f4f2166a04667aea9440ad4850232"
							}]
						}],
						"href": "https://iam.cloud.ibm.com/v1/policies/0b3cea25-fc07-49b1-986b-90846085cfeb",
						"created_at": "2017-11-10T04:37:44.943Z",
						"created_by_id": "IBMid-310000JVN5",
						"last_modified_at": "2017-11-10T04:37:44.943Z",
						"last_modified_by_id": "IBMid-310000JVN5"
				}`, http.Header{"ETag": []string{"3-a283be20686c2c71c512e73e9ee2ce9c"}}),
				))
			})
			It("should return the created policy", func() {
				policy, err := repo.Create(Policy{
					Type: AccessPolicyType,
					Subjects: []Subject{
						{
							Attributes: []Attribute{
								{
									Name:  "iam_id",
									Value: "IBMid-270006V8HD",
								},
							},
						},
					},
					Roles: []Role{
						{
							RoleID: "crn:v1:bluemix:public:iam::::role:Viewer",
						},
					},
					Resources: []Resource{
						{
							Attributes: []Attribute{
								{
									Name:  "resourceType",
									Value: "resource-group",
								},
								{
									Name:  "resource",
									Value: "5e0f4f2166a04667aea9440ad4850232",
								},
							},
						},
					},
				})

				Expect(err).ShouldNot(HaveOccurred())
				Expect(policy).ShouldNot(BeZero())
				Expect(policy.Type).To(Equal(AccessPolicyType))
				Expect(policy.Subjects).To(HaveLen(1))
				Expect(policy.Subjects[0].IAMID()).To(Equal("IBMid-270006V8HD"))
				Expect(policy.Roles).To(HaveLen(1))
				Expect(policy.Roles[0].Name).To(Equal("Viewer"))
				Expect(policy.Resources).To(HaveLen(1))
				Expect(policy.Resources[0].Resource()).To(Equal("5e0f4f2166a04667aea9440ad4850232"))
				Expect(policy.Version).To(Equal("3-a283be20686c2c71c512e73e9ee2ce9c"))
			})
		})

		Context("Creation fails", func() {
			BeforeEach(func() {
				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodPost, "/v1/policies"),
					ghttp.VerifyJSON(`{
						"type": "authorization",
						"subjects": [{
							"attributes": [{
								"name": "iam_id",
								"value": "IBMid-270006V8HD"
							}]
						}],
						"roles": [{
							"role_id": "crn:v1:bluemix:public:iam::::role:Viewer"
						}],
						"resources": [{
							"attributes": [{
								"name": "resourceType",
								"value": "resource-group"
							},{
								"name": "resource",
								"value": "5e0f4f2166a04667aea9440ad4850232"
							}]
						}]
					}`),
					ghttp.RespondWith(http.StatusBadRequest, `{
						"trace": "c2567d5d81674a93af790a512e770b8d",
						"errors": [{
							"code": "invalid_body",
							"message": "Invalid body format. Check missing parameters."
						}],
						"status_code": 400
				}`),
				))
			})
			It("should return error", func() {
				policy, err := repo.Create(Policy{
					Type: AuthorizationPolicyType,
					Subjects: []Subject{
						{
							Attributes: []Attribute{
								{
									Name:  "iam_id",
									Value: "IBMid-270006V8HD",
								},
							},
						},
					},
					Roles: []Role{
						{
							RoleID: "crn:v1:bluemix:public:iam::::role:Viewer",
						},
					},
					Resources: []Resource{
						{
							Attributes: []Attribute{
								{
									Name:  "resourceType",
									Value: "resource-group",
								},
								{
									Name:  "resource",
									Value: "5e0f4f2166a04667aea9440ad4850232",
								},
							},
						},
					},
				})

				Expect(err).Should(HaveOccurred())
				Expect(policy).Should(BeZero())
			})
		})
	})

	Describe("Update()", func() {
		Context("Update successfully", func() {
			BeforeEach(func() {
				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodPut, "/v1/policies/policy-id"),
					ghttp.VerifyHeaderKV("If-Match", "abc"),
					ghttp.VerifyJSON(`{
						"type": "access",
						"subjects": [{
							"attributes": [{
								"name": "iam_id",
								"value": "IBMid-270006V8HD"
							}]
						}],
						"roles": [{
							"role_id": "crn:v1:bluemix:public:iam::::role:Viewer"
						}],
						"resources": [{
							"attributes": [{
								"name": "resourceType",
								"value": "resource-group"
							},{
								"name": "resource",
								"value": "5e0f4f2166a04667aea9440ad4850232"
							}]
						}]
					}`),
					ghttp.RespondWith(http.StatusOK, `{
						"id": "policy-id",
						"type": "access",
						"subjects": [{
							"attributes": [{
								"name": "iam_id",
								"value": "IBMid-270006V8HD"
							}]
						}],
						"roles": [{
							"role_id": "crn:v1:bluemix:public:iam::::role:Viewer",
							"display_name": "Viewer",
							"description": "As a viewer, you can view service instances, but you can't modify them."
						}],
						"resources": [{
							"attributes": [{
								"name": "resourceType",
								"value": "resource-group"
							},{
								"name": "resource",
								"value": "5e0f4f2166a04667aea9440ad4850232"
							}]
						}],
						"href": "https://iam.cloud.ibm.com/v1/policies/0b3cea25-fc07-49b1-986b-90846085cfeb",
						"created_at": "2017-11-10T04:37:44.943Z",
						"created_by_id": "IBMid-310000JVN5",
						"last_modified_at": "2017-11-10T04:37:44.943Z",
						"last_modified_by_id": "IBMid-310000JVN5"
				}`, http.Header{"ETag": []string{"3-a283be20686c2c71c512e73e9ee2ce9c"}}),
				))
			})
			It("should return the updated policy", func() {
				policy, err := repo.Update("policy-id", Policy{
					Type: AccessPolicyType,
					Subjects: []Subject{
						{
							Attributes: []Attribute{
								{
									Name:  "iam_id",
									Value: "IBMid-270006V8HD",
								},
							},
						},
					},
					Roles: []Role{
						{
							RoleID: "crn:v1:bluemix:public:iam::::role:Viewer",
						},
					},
					Resources: []Resource{
						{
							Attributes: []Attribute{
								{
									Name:  "resourceType",
									Value: "resource-group",
								},
								{
									Name:  "resource",
									Value: "5e0f4f2166a04667aea9440ad4850232",
								},
							},
						},
					},
				}, "abc")

				Expect(err).ShouldNot(HaveOccurred())
				Expect(policy).ShouldNot(BeZero())
				Expect(policy.Type).To(Equal(AccessPolicyType))
				Expect(policy.Subjects).To(HaveLen(1))
				Expect(policy.Subjects[0].IAMID()).To(Equal("IBMid-270006V8HD"))
				Expect(policy.Roles).To(HaveLen(1))
				Expect(policy.Roles[0].Name).To(Equal("Viewer"))
				Expect(policy.Resources).To(HaveLen(1))
				Expect(policy.Resources[0].Resource()).To(Equal("5e0f4f2166a04667aea9440ad4850232"))
				Expect(policy.Version).To(Equal("3-a283be20686c2c71c512e73e9ee2ce9c"))
			})
		})

		Context("Update fails", func() {
			BeforeEach(func() {
				server.AppendHandlers(ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodPut, "/v1/policies/policy-id"),
					ghttp.VerifyHeaderKV("If-Match", "abc"),
					ghttp.VerifyJSON(`{
						"type": "authorization",
						"subjects": [{
							"attributes": [{
								"name": "iam_id",
								"value": "IBMid-270006V8HD"
							}]
						}],
						"roles": [{
							"role_id": "crn:v1:bluemix:public:iam::::role:Viewer"
						}],
						"resources": [{
							"attributes": [{
								"name": "resourceType",
								"value": "resource-group"
							},{
								"name": "resource",
								"value": "5e0f4f2166a04667aea9440ad4850232"
							}]
						}]
					}`),
					ghttp.RespondWith(http.StatusBadRequest, `{
						"trace": "c2567d5d81674a93af790a512e770b8d",
						"errors": [{
							"code": "invalid_body",
							"message": "Invalid body format. Check missing parameters."
						}],
						"status_code": 400
				}`),
				))
			})
			It("should return error", func() {
				policy, err := repo.Update("policy-id", Policy{
					Type: AuthorizationPolicyType,
					Subjects: []Subject{
						{
							Attributes: []Attribute{
								{
									Name:  "iam_id",
									Value: "IBMid-270006V8HD",
								},
							},
						},
					},
					Roles: []Role{
						{
							RoleID: "crn:v1:bluemix:public:iam::::role:Viewer",
						},
					},
					Resources: []Resource{
						{
							Attributes: []Attribute{
								{
									Name:  "resourceType",
									Value: "resource-group",
								},
								{
									Name:  "resource",
									Value: "5e0f4f2166a04667aea9440ad4850232",
								},
							},
						},
					},
				}, "abc")

				Expect(err).Should(HaveOccurred())
				Expect(policy).Should(BeZero())
			})
		})
	})
})

func newv1Policy(url string) V1PolicyRepository {

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
	return NewV1PolicyRepository(&client)
}
