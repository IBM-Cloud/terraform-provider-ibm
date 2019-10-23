package iamv1

import (
	"log"
	"net/http"

	"github.com/IBM-Cloud/bluemix-go"

	"github.com/IBM-Cloud/bluemix-go/client"
	"github.com/IBM-Cloud/bluemix-go/crn"
	"github.com/IBM-Cloud/bluemix-go/models"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/onsi/gomega/ghttp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ServiceRoles", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})

	Describe("ListServiceRoles()", func() {
		Context("Service roles are returned", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/acms/v1/roles", "serviceName=cloud-object-storage"),
						ghttp.RespondWith(http.StatusOK, `{
							"supportedRoles": [
								{
									"id": "crn:v1:bluemix:public:iam::::serviceRole:Writer",
									"displayName": "Writer",
									"description": "As a writer, you have permissions beyond the reader role, including creating and editing service-specific resources."
								},
								{
									"id": "crn:v1:bluemix:public:iam::::serviceRole:Reader",
									"displayName": "Reader",
									"description": "As a reader, you can perform read-only actions within a service such as viewing service-specific resources."
								},
								{
									"id": "crn:v1:bluemix:public:iam::::serviceRole:Manager",
									"displayName": "Manager",
									"description": "As a manager, you have permissions beyond the writer role to complete privileged actions as defined by the service. In addition, you can create and edit service-specific resources."
								}
							],
							"platformExtensions": {
								"supportedRoles": [
									{
										"id": "crn:v1:bluemix:public:iam::::role:Administrator",
										"displayName": "Administrator",
										"description": "As an administrator, you can perform all platform actions based on the resource this role is being assigned, including assigning access policies to other users."
									},
									{
										"id": "crn:v1:bluemix:public:iam::::role:Operator",
										"displayName": "Operator",
										"description": "As an operator, you can perform platform actions required to configure and operate service instances, such as viewing a service's dashboard."
									},
									{
										"id": "crn:v1:bluemix:public:iam::::role:Viewer",
										"displayName": "Viewer",
										"description": "As a viewer, you can view service instances, but you can't modify them."
									},
									{
										"id": "crn:v1:bluemix:public:iam::::role:Editor",
										"displayName": "Editor",
										"description": "As an editor, you can perform all platform actions except for managing the account and assigning access policies."
									}
								]
							}
						}`),
					),
				)
			})
			It("should return all roles", func() {
				roles, err := newTestServiceRoleRepo(server.URL()).ListServiceRoles("cloud-object-storage")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(roles).Should(HaveLen(7))
				Expect(roles).Should(Equal([]models.PolicyRole{
					models.PolicyRole{
						ID:          crn.CRN{Scheme: "crn", Version: "v1", CName: "bluemix", CType: "public", ServiceName: "iam", ResourceType: "serviceRole", Resource: "Writer"},
						DisplayName: "Writer",
						Description: "As a writer, you have permissions beyond the reader role, including creating and editing service-specific resources.",
					},
					models.PolicyRole{
						ID:          crn.CRN{Scheme: "crn", Version: "v1", CName: "bluemix", CType: "public", ServiceName: "iam", ResourceType: "serviceRole", Resource: "Reader"},
						DisplayName: "Reader",
						Description: "As a reader, you can perform read-only actions within a service such as viewing service-specific resources.",
					},
					models.PolicyRole{
						ID:          crn.CRN{Scheme: "crn", Version: "v1", CName: "bluemix", CType: "public", ServiceName: "iam", ResourceType: "serviceRole", Resource: "Manager"},
						DisplayName: "Manager",
						Description: "As a manager, you have permissions beyond the writer role to complete privileged actions as defined by the service. In addition, you can create and edit service-specific resources.",
					},
					models.PolicyRole{
						ID:          crn.CRN{Scheme: "crn", Version: "v1", CName: "bluemix", CType: "public", ServiceName: "iam", ResourceType: "role", Resource: "Administrator"},
						DisplayName: "Administrator",
						Description: "As an administrator, you can perform all platform actions based on the resource this role is being assigned, including assigning access policies to other users.",
					},
					models.PolicyRole{
						ID:          crn.CRN{Scheme: "crn", Version: "v1", CName: "bluemix", CType: "public", ServiceName: "iam", ResourceType: "role", Resource: "Operator"},
						DisplayName: "Operator",
						Description: "As an operator, you can perform platform actions required to configure and operate service instances, such as viewing a service's dashboard.",
					},
					models.PolicyRole{
						ID:          crn.CRN{Scheme: "crn", Version: "v1", CName: "bluemix", CType: "public", ServiceName: "iam", ResourceType: "role", Resource: "Viewer"},
						DisplayName: "Viewer",
						Description: "As a viewer, you can view service instances, but you can't modify them.",
					},
					models.PolicyRole{
						ID:          crn.CRN{Scheme: "crn", Version: "v1", CName: "bluemix", CType: "public", ServiceName: "iam", ResourceType: "role", Resource: "Editor"},
						DisplayName: "Editor",
						Description: "As an editor, you can perform all platform actions except for managing the account and assigning access policies.",
					},
				}))
			})
		})

		Context("Service not found", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/acms/v1/roles", "serviceName=cloud-object-storage"),
						ghttp.RespondWith(http.StatusNotFound, `{
							"errorsArray": [
								{
									"code": "BXNAC12104",
									"response": "not_found_error",
									"message": "Not Found Service name given returned empty query. ",
									"level": "error",
									"statusCode": 404,
									"description": "Service name given returned empty query. ",
									"transactionId": "string",
									"instanceId": "a4d1e7b5-3f2f-4242-b555-78b7e695bbb1"
								}
							]
						}`),
					),
				)
			})

			It("should return error", func() {
				roles, err := newTestServiceRoleRepo(server.URL()).ListServiceRoles("cloud-object-storage")
				Expect(err).Should(HaveOccurred())
				Expect(roles).Should(BeEmpty())
			})
		})

	})

	Describe("ListSystemDefinedRoles()", func() {
		Context("System defined roles are returned", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/acms/v1/roles"),
						ghttp.RespondWith(http.StatusOK, `{
							"systemDefinedRoles": [
								{
									"crn": "crn:v1:bluemix:public:iam::::serviceRole:IAMAuthz",
									"displayName": "IAMAuthz",
									"description": "IAMAuthz"
								},
								{
									"crn": "crn:v1:bluemix:public:iam::::role:Administrator",
									"displayName": "Administrator",
									"description": "As an administrator, you can perform all platform actions based on the resource this role is being assigned, including assigning access policies to other users."
								},
								{
									"crn": "crn:v1:bluemix:public:iam::::role:Operator",
									"displayName": "Operator",
									"description": "As an operator, you can perform platform actions required to configure and operate service instances, such as viewing a service's dashboard."
								},
								{
									"crn": "crn:v1:bluemix:public:iam::::role:Viewer",
									"displayName": "Viewer",
									"description": "As a viewer, you can view service instances, but you can't modify them."
								},
								{
									"crn": "crn:v1:bluemix:public:iam::::role:Editor",
									"displayName": "Editor",
									"description": "As an editor, you can perform all platform actions except for managing the account and assigning access policies."
								},
								{
									"crn": "crn:v1:bluemix:public:iam::::serviceRole:Reader",
									"displayName": "Reader",
									"description": "As a reader, you can perform read-only actions within a service such as viewing service-specific resources."
								},
								{
									"crn": "crn:v1:bluemix:public:iam::::serviceRole:Writer",
									"displayName": "Writer",
									"description": "As a writer, you have permissions beyond the reader role, including creating and editing service-specific resources."
								},
								{
									"crn": "crn:v1:bluemix:public:iam::::serviceRole:Manager",
									"displayName": "Manager",
									"description": "As a manager, you have permissions beyond the writer role to complete privileged actions as defined by the service. In addition, you can create and edit service-specific resources."
								}
							]
						}`),
					),
				)
			})
			It("should return all roles", func() {
				roles, err := newTestServiceRoleRepo(server.URL()).ListSystemDefinedRoles()
				Expect(err).ShouldNot(HaveOccurred())
				Expect(roles).Should(HaveLen(8))
				Expect(roles).Should(Equal([]models.PolicyRole{
					{
						ID:          crn.CRN{Scheme: "crn", Version: "v1", CName: "bluemix", CType: "public", ServiceName: "iam", ResourceType: "serviceRole", Resource: "IAMAuthz"},
						DisplayName: "IAMAuthz",
						Description: "IAMAuthz",
					},
					{
						ID:          crn.CRN{Scheme: "crn", Version: "v1", CName: "bluemix", CType: "public", ServiceName: "iam", ResourceType: "role", Resource: "Administrator"},
						DisplayName: "Administrator",
						Description: "As an administrator, you can perform all platform actions based on the resource this role is being assigned, including assigning access policies to other users.",
					},
					{
						ID:          crn.CRN{Scheme: "crn", Version: "v1", CName: "bluemix", CType: "public", ServiceName: "iam", ResourceType: "role", Resource: "Operator"},
						DisplayName: "Operator",
						Description: "As an operator, you can perform platform actions required to configure and operate service instances, such as viewing a service's dashboard.",
					},
					{
						ID:          crn.CRN{Scheme: "crn", Version: "v1", CName: "bluemix", CType: "public", ServiceName: "iam", ResourceType: "role", Resource: "Viewer"},
						DisplayName: "Viewer",
						Description: "As a viewer, you can view service instances, but you can't modify them.",
					},
					{
						ID:          crn.CRN{Scheme: "crn", Version: "v1", CName: "bluemix", CType: "public", ServiceName: "iam", ResourceType: "role", Resource: "Editor"},
						DisplayName: "Editor",
						Description: "As an editor, you can perform all platform actions except for managing the account and assigning access policies.",
					},
					{
						ID:          crn.CRN{Scheme: "crn", Version: "v1", CName: "bluemix", CType: "public", ServiceName: "iam", ResourceType: "serviceRole", Resource: "Reader"},
						DisplayName: "Reader",
						Description: "As a reader, you can perform read-only actions within a service such as viewing service-specific resources.",
					},
					{
						ID:          crn.CRN{Scheme: "crn", Version: "v1", CName: "bluemix", CType: "public", ServiceName: "iam", ResourceType: "serviceRole", Resource: "Writer"},
						DisplayName: "Writer",
						Description: "As a writer, you have permissions beyond the reader role, including creating and editing service-specific resources.",
					},
					{
						ID:          crn.CRN{Scheme: "crn", Version: "v1", CName: "bluemix", CType: "public", ServiceName: "iam", ResourceType: "serviceRole", Resource: "Manager"},
						DisplayName: "Manager",
						Description: "As a manager, you have permissions beyond the writer role to complete privileged actions as defined by the service. In addition, you can create and edit service-specific resources.",
					},
				}))
			})
		})

		Context("User token expires", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/acms/v1/roles"),
						ghttp.RespondWith(http.StatusNotFound, `{
							"errors": [
								{
									"code": "BXNAC12101",
									"response": "unauthorized_error",
									"message": "Unauthorized",
									"description": "Your access is not authorized.",
									"level": "error",
									"statusCode": 401,
									"transactionId": "string",
									"instanceId": "fa48b62a-e1bf-4d25-97e1-542d33ad51bb"
								}
							]
						}`),
					),
				)
			})

			It("should return error", func() {
				roles, err := newTestServiceRoleRepo(server.URL()).ListSystemDefinedRoles()
				Expect(err).Should(HaveOccurred())
				Expect(roles).Should(BeEmpty())
			})
		})
	})
})

func newTestServiceRoleRepo(url string) ServiceRoleRepository {
	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}
	conf := sess.Config.Copy()
	conf.Endpoint = &url

	client := client.Client{
		Config:      conf,
		ServiceName: bluemix.AccountService,
	}

	return NewServiceRoleRepository(&client)
}
