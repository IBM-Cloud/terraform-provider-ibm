package iamuumv1

import (
	"log"
	"net/http"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/client"
	"github.com/IBM-Cloud/bluemix-go/models"
	"github.com/IBM-Cloud/bluemix-go/session"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("AccessGroupMemberRepository", func() {
	var (
		server *ghttp.Server
	)

	Describe("List()", func() {
		Context("When API error 403 returns", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/groups/def/members"),
						ghttp.RespondWith(http.StatusForbidden, `
						{
							"message": "The provided access token does not have the proper authority to access this operation."
						}`),
					),
				)
			})

			It("should return API 403 error", func() {
				_, err := newTestAccessGroupMemberRepo(server.URL()).List("def")
				Expect(err).Should(HaveOccurred())
				Expect(err.Error()).Should(ContainSubstring("Request failed with status code: 403"))
			})
		})

		Context("When other JSON error returns", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/groups/abc/members"),
						ghttp.RespondWith(http.StatusBadGateway, `{
							"message": "other json error"
						}`),
					),
				)
			})

			It("should return server error", func() {
				_, err := newTestAccessGroupMemberRepo(server.URL()).List("abc")
				Expect(err).Should(HaveOccurred())
				Expect(err.Error()).Should(ContainSubstring("other json error"))
			})
		})

		Context("When no group member returned", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/groups/abc/members"),
						ghttp.RespondWith(http.StatusOK, `
						{
							"limit": 50,
							"offset": 0,
							"total_count": 0
						}`),
					),
				)
			})

			It("should return empty", func() {
				members, err := newTestAccessGroupMemberRepo(server.URL()).List("abc")

				Expect(err).ShouldNot(HaveOccurred())
				Expect(members).Should(BeEmpty())
			})
		})

		Context("When there is one page", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/groups/abc/members"),
						ghttp.RespondWith(http.StatusOK, `
						{
							"limit": 50,
							"offset": 0,
							"total_count": 2,
							"members": [{
								"href": "https://iam.stage1.bluemix.net/008facc4-412f-463e-bd1b-99dd7dcfa27b",
								"id": "008facc4-412f-463e-bd1b-99dd7dcfa27b",
								"type": "Editor"
							},{
								"href": "https://iam.stage1.bluemix.net/048af74a-8435-4783-8ad9-8e207fa24afd",
								"id": "048af74a-8435-4783-8ad9-8e207fa24afd",
								"type": "Viewer"
							}]
						}`),
					),
				)
			})

			It("should return one page", func() {
				members, err := newTestAccessGroupMemberRepo(server.URL()).List("abc")

				Expect(err).ShouldNot(HaveOccurred())
				Expect(members).Should(HaveLen(2))

				Expect(members[0].ID).Should(Equal("008facc4-412f-463e-bd1b-99dd7dcfa27b"))
				Expect(members[0].Type).Should(Equal("Editor"))

				Expect(members[1].ID).Should(Equal("048af74a-8435-4783-8ad9-8e207fa24afd"))
				Expect(members[1].Type).Should(Equal("Viewer"))
			})
		})

		Context("When there are multiple pages", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/groups/abc/members"),
						ghttp.RespondWith(http.StatusOK, `
						{
							"limit": 1,
							"offset": 0,
							"total_count": 2,
							"next": {
								"href": "`+server.URL()+`/v1/groups/abc/members?page=1"
							},
							"members": [{
								"href": "https://iam.stage1.bluemix.net/008facc4-412f-463e-bd1b-99dd7dcfa27b",
								"id": "008facc4-412f-463e-bd1b-99dd7dcfa27b",
								"type": "Editor"
							}]
						}`),
					),
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/groups/abc/members", "page=1"),
						ghttp.RespondWith(http.StatusOK, `
						{
							"limit": 1,
							"offset": 1,
							"total_count": 2,
							"members": [{
								"href": "https://iam.stage1.bluemix.net/048af74a-8435-4783-8ad9-8e207fa24afd",
								"id": "048af74a-8435-4783-8ad9-8e207fa24afd",
								"type": "Viewer"
							}]
						}`),
					),
				)
			})

			It("should return all pages", func() {
				members, err := newTestAccessGroupMemberRepo(server.URL()).List("abc")

				Expect(err).ShouldNot(HaveOccurred())
				Expect(members).Should(HaveLen(2))

				Expect(members[0].ID).Should(Equal("008facc4-412f-463e-bd1b-99dd7dcfa27b"))
				Expect(members[0].Type).Should(Equal("Editor"))

				Expect(members[1].ID).Should(Equal("048af74a-8435-4783-8ad9-8e207fa24afd"))
				Expect(members[1].Type).Should(Equal("Viewer"))
			})
		})
	})

	Describe("Add()", func() {
		Context("When add one group member", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, "/v1/groups/abc/members"),
						ghttp.VerifyJSONRepresenting(AddGroupMemberRequest{
							Members: []models.AccessGroupMember{
								models.AccessGroupMember{
									ID:   "test",
									Type: AccessGroupMemberUser,
								},
							},
						}),
						ghttp.RespondWith(http.StatusMultiStatus, `
						{
							"members": [{
								"msg":"",
								"ok":true,
								"id":"test",
								"type":"user"
							}]
						}`),
					),
				)
			})

			It("should return success", func() {
				response, err := newTestAccessGroupMemberRepo(server.URL()).Add("abc", AddGroupMemberRequest{
					Members: []models.AccessGroupMember{
						models.AccessGroupMember{
							ID:   "test",
							Type: AccessGroupMemberUser,
						},
					},
				})

				Expect(err).ShouldNot(HaveOccurred())
				Expect(response.Members).Should(HaveLen(1))
				Expect(response.Members[0].ID).Should(Equal("test"))
				Expect(response.Members[0].Type).Should(Equal(AccessGroupMemberUser))
				Expect(response.Members[0].OK).Should(BeTrue())
			})
		})

		Context("When add multiple group members with partial success", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, "/v1/groups/abc/members"),
						ghttp.VerifyJSONRepresenting(AddGroupMemberRequest{
							Members: []models.AccessGroupMember{
								models.AccessGroupMember{
									ID:   "test",
									Type: AccessGroupMemberUser,
								},
								models.AccessGroupMember{
									ID:   "test2",
									Type: AccessGroupMemberUser,
								},
							},
						}),
						ghttp.RespondWith(http.StatusMultiStatus, `
						{
							"members": [{
								"msg":"",
								"ok":true,
								"id":"test",
								"type":"user"
							},{
								"msg":"",
								"ok":false,
								"id":"test2",
								"type":"user"
							}]
						}`),
					),
				)
			})

			It("should return partial success", func() {
				response, err := newTestAccessGroupMemberRepo(server.URL()).Add("abc", AddGroupMemberRequest{
					Members: []models.AccessGroupMember{
						models.AccessGroupMember{
							ID:   "test",
							Type: AccessGroupMemberUser,
						},
						models.AccessGroupMember{
							ID:   "test2",
							Type: AccessGroupMemberUser,
						},
					},
				})

				Expect(err).ShouldNot(HaveOccurred())
				Expect(response.Members).Should(HaveLen(2))
				Expect(response.Members[0].ID).Should(Equal("test"))
				Expect(response.Members[0].Type).Should(Equal(AccessGroupMemberUser))
				Expect(response.Members[0].OK).Should(BeTrue())
				Expect(response.Members[1].ID).Should(Equal("test2"))
				Expect(response.Members[1].Type).Should(Equal(AccessGroupMemberUser))
				Expect(response.Members[1].OK).Should(BeFalse())
			})
		})

		Context("When access group is not found", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, "/v1/groups/abc/members"),
						ghttp.VerifyJSONRepresenting(AddGroupMemberRequest{
							Members: []models.AccessGroupMember{
								models.AccessGroupMember{
									ID:   "test",
									Type: AccessGroupMemberUser,
								},
							},
						}),
						ghttp.RespondWith(http.StatusNotFound, `
						{
							"StatusCode": 404,
							"code": "not_found",
							"message": "No groups found for the member test."
						}`),
					),
				)
			})

			It("should return not found error", func() {
				_, err := newTestAccessGroupMemberRepo(server.URL()).Add("abc", AddGroupMemberRequest{
					Members: []models.AccessGroupMember{
						models.AccessGroupMember{
							ID:   "test",
							Type: AccessGroupMemberUser,
						},
					},
				})

				Expect(err).Should(HaveOccurred())
				Expect(err.Error()).Should(ContainSubstring("not_found"))
			})
		})
	})

	Describe("Remove()", func() {
		Context("When member is deleted", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/groups/abc/members/test"),
						ghttp.RespondWith(http.StatusNoContent, ""),
					),
				)
			})

			It("should return success", func() {
				err := newTestAccessGroupMemberRepo(server.URL()).Remove("abc", "test")

				Expect(err).Should(Succeed())
			})
		})

		Context("When member is not found", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/groups/abc/members/test"),
						ghttp.RespondWith(http.StatusNotFound, `{
							"StatusCode": 404,
							"code": "not_found",
							"message": "Group member test is not found"
						}`),
					),
				)
			})

			It("should return not found error", func() {
				err := newTestAccessGroupMemberRepo(server.URL()).Remove("abc", "test")

				Expect(err).Should(HaveOccurred())
				Expect(err.Error()).Should(ContainSubstring("not_found"))
			})
		})
	})

})

func newTestAccessGroupMemberRepo(url string) AccessGroupMemberRepository {
	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}
	conf := sess.Config.Copy()
	conf.Endpoint = &url
	client := client.Client{
		Config:      conf,
		ServiceName: bluemix.IAMUUMService,
	}
	return NewAccessGroupMemberRepository(&client)
}
