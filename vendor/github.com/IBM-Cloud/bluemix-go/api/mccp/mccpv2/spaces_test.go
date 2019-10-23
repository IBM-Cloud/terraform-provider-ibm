package mccpv2

import (
	"log"
	"net/http"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/client"
	"github.com/IBM-Cloud/bluemix-go/helpers"
	"github.com/IBM-Cloud/bluemix-go/session"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Spaces", func() {
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
						ghttp.VerifyRequest(http.MethodPost, "/v2/spaces"),
						ghttp.VerifyBody([]byte(`{"name":"testspace","organization_guid":"3c1b6f9d-ffe5-43b5-ab91-7be2331dc546"}`)),
						ghttp.RespondWith(http.StatusCreated, `{
							 "metadata": {
									"guid": "64ff2b7d-b6d9-48c6-94a2-7f4ba670d67b",
									"url": "/v2/space_quota_definitions/64ff2b7d-b6d9-48c6-94a2-7f4ba670d67b",
									"created_at": "2017-05-03T08:52:07Z",
									"updated_at": null

							  },
							  "entity": {
							    "name": "testspace",
								"organization_guid": "3c1b6f9d-ffe5-43b5-ab91-7be2331dc546",
								"space_quota_definition_guid": null,
								"allow_ssh": true
							  }							
						}`),
					),
				)
			})

			It("should return Spaces created", func() {
				payload := SpaceCreateRequest{
					Name:    "testspace",
					OrgGUID: "3c1b6f9d-ffe5-43b5-ab91-7be2331dc546",
				}
				myspace, err := newSpaces(server.URL()).Create(payload)
				Expect(err).NotTo(HaveOccurred())
				Expect(myspace).ShouldNot(BeNil())
				Expect(myspace.Metadata.GUID).Should(Equal("64ff2b7d-b6d9-48c6-94a2-7f4ba670d67b"))
				Expect(myspace.Entity.Name).Should(Equal("testspace"))
				Expect(myspace.Entity.OrgGUID).Should(Equal("3c1b6f9d-ffe5-43b5-ab91-7be2331dc546"))
				Expect(myspace.Entity.SpaceQuotaGUID).Should(Equal(""))
			})
		})
		Context("When creation is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v2/spaces"),
						ghttp.VerifyBody([]byte(`{"name":"testspace","organization_guid":"3c1b6f9d-ffe5-43b5-ab91-7be2331dc546"}`)),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to create space`),
					),
				)
			})

			It("should return error when space is created", func() {
				payload := SpaceCreateRequest{
					Name:    "testspace",
					OrgGUID: "3c1b6f9d-ffe5-43b5-ab91-7be2331dc546",
				}
				myspace, err := newSpaces(server.URL()).Create(payload)
				Expect(err).To(HaveOccurred())
				Expect(myspace).Should(BeNil())
			})
		})
	})

	Describe("Get", func() {
		Context("When read of space is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/spaces/64ff2b7d-b6d9-48c6-94a2-7f4ba670d67b"),
						ghttp.RespondWith(http.StatusOK, `{
							 "metadata": {
									"guid": "64ff2b7d-b6d9-48c6-94a2-7f4ba670d67b",
									"url": "/v2/space_quota_definitions/64ff2b7d-b6d9-48c6-94a2-7f4ba670d67b",
									"created_at": "2017-05-03T08:52:07Z",
									"updated_at": null

							  },
							  "entity": {
							    "name": "testspace",
								"organization_guid": "3c1b6f9d-ffe5-43b5-ab91-7be2331dc546",
								"space_quota_definition_guid": null,
								"allow_ssh": true
							  }							
						}`),
					),
				)
			})

			It("should return Space", func() {
				myspace, err := newSpaces(server.URL()).Get("64ff2b7d-b6d9-48c6-94a2-7f4ba670d67b")
				Expect(err).NotTo(HaveOccurred())
				Expect(myspace).ShouldNot(BeNil())
				Expect(myspace.Metadata.GUID).Should(Equal("64ff2b7d-b6d9-48c6-94a2-7f4ba670d67b"))
				Expect(myspace.Entity.Name).Should(Equal("testspace"))
				Expect(myspace.Entity.OrgGUID).Should(Equal("3c1b6f9d-ffe5-43b5-ab91-7be2331dc546"))
				Expect(myspace.Entity.SpaceQuotaGUID).Should(Equal(""))
			})
		})
		Context("When space retrieve is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/spaces/64ff2b7d-b6d9-48c6-94a2-7f4ba670d67b"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve space`),
					),
				)
			})

			It("should return error when space is retrieved", func() {
				myspace, err := newSpaces(server.URL()).Get("64ff2b7d-b6d9-48c6-94a2-7f4ba670d67b")
				Expect(err).To(HaveOccurred())
				Expect(myspace).Should(BeNil())
			})
		})
	})

	Describe("Update", func() {
		Context("When update of space is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, "/v2/spaces/64ff2b7d-b6d9-48c6-94a2-7f4ba670d67b"),
						ghttp.VerifyBody([]byte(`{"name":"testspaceupdate"}`)),
						ghttp.RespondWith(http.StatusCreated, `{
							 "metadata": {
									"guid": "64ff2b7d-b6d9-48c6-94a2-7f4ba670d67b",
									"url": "/v2/space_quota_definitions/64ff2b7d-b6d9-48c6-94a2-7f4ba670d67b",
									"created_at": "2017-05-03T08:52:07Z",
									"updated_at": "2017-05-03T11:02:23Z"

							  },
							  "entity": {
							    "name": "testspaceupdate",
								"organization_guid": "3c1b6f9d-ffe5-43b5-ab91-7be2331dc546",
								"space_quota_definition_guid": null,
								"allow_ssh": true
							  }							
						}`),
					),
				)
			})

			It("should return space update", func() {
				payload := SpaceUpdateRequest{
					Name: helpers.String("testspaceupdate"),
				}
				myspace, err := newSpaces(server.URL()).Update("64ff2b7d-b6d9-48c6-94a2-7f4ba670d67b", payload)
				Expect(err).NotTo(HaveOccurred())
				Expect(myspace).ShouldNot(BeNil())
				Expect(myspace.Metadata.GUID).Should(Equal("64ff2b7d-b6d9-48c6-94a2-7f4ba670d67b"))
				Expect(myspace.Entity.Name).Should(Equal("testspaceupdate"))
				Expect(myspace.Entity.OrgGUID).Should(Equal("3c1b6f9d-ffe5-43b5-ab91-7be2331dc546"))
				Expect(myspace.Entity.SpaceQuotaGUID).Should(Equal(""))
			})
		})
		Context("When space update is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, "/v2/spaces/64ff2b7d-b6d9-48c6-94a2-7f4ba670d67b"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve space`),
					),
				)
			})

			It("should return error when space updated", func() {
				payload := SpaceUpdateRequest{
					Name: helpers.String("testspaceupdate"),
				}
				myspace, err := newSpaces(server.URL()).Update("64ff2b7d-b6d9-48c6-94a2-7f4ba670d67b", payload)
				Expect(err).To(HaveOccurred())
				Expect(myspace).Should(BeNil())
			})
		})
	})
	Describe("Delete", func() {
		Context("When delete of space is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v2/spaces/64ff2b7d-b6d9-48c6-94a2-7f4ba670d67b"),
						ghttp.RespondWith(http.StatusOK, `{							
						}`),
					),
				)
			})

			It("should delete Space", func() {
				err := newSpaces(server.URL()).Delete("64ff2b7d-b6d9-48c6-94a2-7f4ba670d67b")
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When space update is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v2/spaces/64ff2b7d-b6d9-48c6-94a2-7f4ba670d67b"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve space`),
					),
				)
			})

			It("should return error when space  deleted", func() {
				err := newSpaces(server.URL()).Delete("64ff2b7d-b6d9-48c6-94a2-7f4ba670d67b")
				Expect(err).To(HaveOccurred())
			})
		})
	})

})

var _ = Describe("Space Repository", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})
	Describe("ListSpacesInOrg()", func() {
		Context("Server return one space", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/organizations/3c1b6f9d-ffe5-43b5-ab91-7be2331dc546/spaces"),
						ghttp.RespondWith(http.StatusOK, `{
							"total_results": 1,
							"total_pages": 1,
							"prev_url": null,
							"next_url": null,
							"resources": [
							{
								"metadata": {
									"guid": "9fd6aed4-b36b-438e-832f-9f29a68ad61c",
									"url": "/v2/spaces/9fd6aed4-b36b-438e-832f-9f29a68ad61c",
									"created_at": "2015-08-18T09:45:19Z",
									"updated_at": "2016-03-07T13:41:50Z"
								},
								"entity": {
									"name": "prod",
									"organization_guid": "3c1b6f9d-ffe5-43b5-ab91-7be2331dc546",
									"space_quota_definition_guid": null,
									"allow_ssh": true
								}
										
							}]
							
								
						}`),
					),
				)
			})

			It("should return one space", func() {
				myspaces, err := newSpaces(server.URL()).ListSpacesInOrg("3c1b6f9d-ffe5-43b5-ab91-7be2331dc546", "region")
				Expect(err).To(Succeed())
				Expect(len(myspaces)).To(Equal(1))

				space := myspaces[0]
				Expect(space.GUID).To(Equal("9fd6aed4-b36b-438e-832f-9f29a68ad61c"))
				Expect(space.Name).To(Equal("prod"))

			})

		})
		Context("Server return multiple space", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/organizations/3c1b6f9d-ffe5-43b5-ab91-7be2331dc546/spaces"),
						ghttp.RespondWith(http.StatusOK, `{
							"total_results": 2,
							"total_pages": 1,
							"prev_url": null,
							"next_url": null,
							"resources": [
							{
								"metadata": {
									"guid": "9fd6aed4-b36b-438e-832f-9f29a68ad61c",
									"url": "/v2/spaces/9fd6aed4-b36b-438e-832f-9f29a68ad61c",
									"created_at": "2015-08-18T09:45:19Z",
									"updated_at": "2016-03-07T13:41:50Z"
								},
								"entity": {
									"name": "prod",
									"organization_guid": "3c1b6f9d-ffe5-43b5-ab91-7be2331dc546",
									"space_quota_definition_guid": null,
									"allow_ssh": true
								}
										
							},
							{
								"metadata": {
									"guid": "af759fe9-613d-44c7-81df-65d06d13d723",
									"url": "/v2/spaces/af759fe9-613d-44c7-81df-65d06d13d723",
									"created_at": "2017-05-03T16:55:41Z",
									"updated_at": null
								},
								"entity": {
									"name": "test",
									"organization_guid": "3c1b6f9d-ffe5-43b5-ab91-7be2331dc546",
									"space_quota_definition_guid": null,
									"allow_ssh": true
       
								}
							}
							]
															
						}`),
					),
				)
			})

			It("should return multiple spaces", func() {
				myspaces, err := newSpaces(server.URL()).ListSpacesInOrg("3c1b6f9d-ffe5-43b5-ab91-7be2331dc546", "region")
				Expect(err).To(Succeed())
				Expect(len(myspaces)).To(Equal(2))

				space := myspaces[0]
				Expect(space.GUID).To(Equal("9fd6aed4-b36b-438e-832f-9f29a68ad61c"))
				Expect(space.Name).To(Equal("prod"))

				space = myspaces[1]
				Expect(space.GUID).To(Equal("af759fe9-613d-44c7-81df-65d06d13d723"))
				Expect(space.Name).To(Equal("test"))

			})

		})
		Context("Server return no spaces", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/organizations/3c1b6f9d-ffe5-43b5-ab91-7be2331dc546/spaces"),
						ghttp.RespondWith(http.StatusOK, `{
							"total_results": 0,
							"resources": [
							]
															
						}`),
					),
				)
			})

			It("should return no spaces", func() {
				myspaces, err := newSpaces(server.URL()).ListSpacesInOrg("3c1b6f9d-ffe5-43b5-ab91-7be2331dc546", "region")
				Expect(err).To(Succeed())
				Expect(len(myspaces)).To(Equal(0))
			})

		})
		Context("Server return error", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/organizations/3c1b6f9d-ffe5-43b5-ab91-7be2331dc546/spaces"),
						ghttp.RespondWith(http.StatusInternalServerError, `{
															
						}`),
					),
				)
			})

			It("should return error", func() {
				myspaces, err := newSpaces(server.URL()).ListSpacesInOrg("3c1b6f9d-ffe5-43b5-ab91-7be2331dc546", "region")
				Expect(err).To(HaveOccurred())
				Expect(myspaces).To(BeNil())
			})

		})

	})
})

var _ = Describe("Space by Name Repository", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})
	Describe("FindByNameInOrg()", func() {
		Context("Server return space by name", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/organizations/3c1b6f9d-ffe5-43b5-ab91-7be2331dc546/spaces"),
						ghttp.RespondWith(http.StatusOK, `{
							"total_results": 1,
							"total_pages": 1,
							"prev_url": null,
							"next_url": null,
							"resources": [
							{
								"metadata": {
									"guid": "9fd6aed4-b36b-438e-832f-9f29a68ad61c",
									"url": "/v2/spaces/9fd6aed4-b36b-438e-832f-9f29a68ad61c",
									"created_at": "2015-08-18T09:45:19Z",
									"updated_at": "2016-03-07T13:41:50Z"
								},
								"entity": {
									"name": "prod",
									"organization_guid": "3c1b6f9d-ffe5-43b5-ab91-7be2331dc546",
									"space_quota_definition_guid": null,
									"allow_ssh": true
								}
										
							}]
							
								
						}`),
					),
				)
			})

			It("should return one space", func() {
				myspaces, err := newSpaces(server.URL()).FindByNameInOrg("3c1b6f9d-ffe5-43b5-ab91-7be2331dc546", "prod", "region")
				Expect(err).To(Succeed())
				Expect(myspaces.GUID).To(Equal("9fd6aed4-b36b-438e-832f-9f29a68ad61c"))
				Expect(myspaces.Name).To(Equal("prod"))

			})

		})

		Context("Server return no space by name", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/organizations/3c1b6f9d-ffe5-43b5-ab91-7be2331dc546/spaces"),
						ghttp.RespondWith(http.StatusOK, `{
							"total_results": 0,
							"resources": [
							]
															
						}`),
					),
				)
			})

			It("should return no spaces", func() {
				myspaces, err := newSpaces(server.URL()).FindByNameInOrg("3c1b6f9d-ffe5-43b5-ab91-7be2331dc546", "xyz", "region")
				Expect(err).To(HaveOccurred())
				Expect(myspaces).To(BeNil())
			})

		})
		Context("Server return error", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/organizations/3c1b6f9d-ffe5-43b5-ab91-7be2331dc546/spaces"),
						ghttp.RespondWith(http.StatusInternalServerError, `{
															
						}`),
					),
				)
			})

			It("should return error", func() {
				myspaces, err := newSpaces(server.URL()).FindByNameInOrg("3c1b6f9d-ffe5-43b5-ab91-7be2331dc546", "prod", "region")
				Expect(err).To(HaveOccurred())
				Expect(myspaces).To(BeNil())
			})

		})

	})
})

func newSpaces(url string) Spaces {

	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}
	conf := sess.Config.Copy()
	conf.Endpoint = &url
	client := client.Client{
		Config:      conf,
		ServiceName: bluemix.MccpService,
	}
	return newSpacesAPI(&client)
}
