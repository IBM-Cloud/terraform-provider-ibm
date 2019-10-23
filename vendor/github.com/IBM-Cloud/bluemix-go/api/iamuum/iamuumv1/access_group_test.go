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

var _ = Describe("AccessGroupRepository", func() {
	var (
		server *ghttp.Server
	)

	Describe("List()", func() {
		Context("When API error 403 returns", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/groups"),
						ghttp.RespondWith(http.StatusForbidden, `
						{
							"message": "The provided access token does not have the proper authority to access this operation."
						}`),
					),
				)
			})

			It("should return API 403 error", func() {
				_, err := newTestAccessGroupRepo(server.URL()).List("def")
				Expect(err).Should(HaveOccurred())
				Expect(err.Error()).Should(ContainSubstring("Request failed with status code: 403"))
			})
		})

		Context("When other JSON error returns", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/groups"),
						ghttp.RespondWith(http.StatusBadGateway, `{
							"message": "other json error"
						}`),
					),
				)
			})

			It("should return server error", func() {
				_, err := newTestAccessGroupRepo(server.URL()).List("abc")
				Expect(err).Should(HaveOccurred())
				Expect(err.Error()).Should(ContainSubstring("other json error"))
			})
		})

		Context("When no group returned", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/groups"),
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
				groups, err := newTestAccessGroupRepo(server.URL()).List("abc")

				Expect(err).ShouldNot(HaveOccurred())
				Expect(groups).Should(BeEmpty())
			})
		})

		Context("When there is one page", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/groups"),
						ghttp.RespondWith(http.StatusOK, `
						{
							"limit": 50,
							"offset": 0,
							"total_count": 2,
							"groups": [{
								"description": "Editor group",
								"id": "008facc4-412f-463e-bd1b-99dd7dcfa27b",
								"name": "Editor"
							},{
								"description": "Viewer group",
								"id": "048af74a-8435-4783-8ad9-8e207fa24afd",
								"name": "Viewer"
							}]
						}`),
					),
				)
			})

			It("should return one page", func() {
				groups, err := newTestAccessGroupRepo(server.URL()).List("abc")

				Expect(err).ShouldNot(HaveOccurred())
				Expect(groups).Should(HaveLen(2))

				Expect(groups[0].ID).Should(Equal("008facc4-412f-463e-bd1b-99dd7dcfa27b"))
				Expect(groups[0].Name).Should(Equal("Editor"))

				Expect(groups[1].ID).Should(Equal("048af74a-8435-4783-8ad9-8e207fa24afd"))
				Expect(groups[1].Name).Should(Equal("Viewer"))
			})
		})

		Context("When there are multiple pages", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/groups"),
						ghttp.RespondWith(http.StatusOK, `
						{
							"limit": 1,
							"offset": 0,
							"total_count": 2,
							"next": {
								"href": "`+server.URL()+`/v1/groups?page=1"
							},
							"groups": [{
								"description": "Editor group",
								"id": "008facc4-412f-463e-bd1b-99dd7dcfa27b",
								"name": "Editor"
							}]
						}`),
					),
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/groups", "page=1"),
						ghttp.RespondWith(http.StatusOK, `
						{
							"limit": 1,
							"offset": 1,
							"total_count": 2,
							"groups": [{
								"description": "Viewer group",
								"id": "048af74a-8435-4783-8ad9-8e207fa24afd",
								"name": "Viewer"
							}]
						}`),
					),
				)
			})

			It("should return all pages", func() {
				groups, err := newTestAccessGroupRepo(server.URL()).List("abc")

				Expect(err).ShouldNot(HaveOccurred())
				Expect(groups).Should(HaveLen(2))

				Expect(groups[0].ID).Should(Equal("008facc4-412f-463e-bd1b-99dd7dcfa27b"))
				Expect(groups[0].Name).Should(Equal("Editor"))

				Expect(groups[1].ID).Should(Equal("048af74a-8435-4783-8ad9-8e207fa24afd"))
				Expect(groups[1].Name).Should(Equal("Viewer"))
			})
		})
	})

	Describe("Create()", func() {
		Context("When create one group", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/groups"),
						ghttp.RespondWith(http.StatusOK, `
						{
							"description": "abc group",
							"id": "008facc4-412f-463e-bd1b-99dd7dcfa27b",
							"name": "abc"
							
						}`),
					),
				)
			})

			It("should return success", func() {
				response, err := newTestAccessGroupRepo(server.URL()).Create(models.AccessGroup{
					Name:        "abc",
					Description: "abc group",
				}, "89999998-8880")
				Expect(err).ShouldNot(HaveOccurred())

				Expect(response.Name).Should(Equal("abc"))
				Expect(response.Description).Should(Equal("abc group"))
			})

		})
	})

	Describe("Remove()", func() {
		Context("When group is deleted", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/groups/abc"),
						ghttp.RespondWith(http.StatusNoContent, ""),
					),
				)
			})

			It("should return success", func() {
				err := newTestAccessGroupRepo(server.URL()).Delete("abc", false)

				Expect(err).Should(Succeed())
			})
		})

		Context("When group is not found", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/groups/abc"),
						ghttp.RespondWith(http.StatusNotFound, `{
							"StatusCode": 404,
							"code": "not_found",
							"message": "Group abc is not found"
						}`),
					),
				)
			})

			It("should return not found error", func() {
				err := newTestAccessGroupRepo(server.URL()).Delete("abc", false)

				Expect(err).Should(HaveOccurred())
				Expect(err.Error()).Should(ContainSubstring("not_found"))
			})
		})
	})

})

func newTestAccessGroupRepo(url string) AccessGroupRepository {
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
	return NewAccessGroupRepository(&client)
}
