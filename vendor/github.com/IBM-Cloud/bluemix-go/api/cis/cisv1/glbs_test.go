package cisv1

import (
	"log"
	"net/http"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/client"
	bluemixHttp "github.com/IBM-Cloud/bluemix-go/http"
	"github.com/IBM-Cloud/bluemix-go/session"

	"github.com/onsi/gomega/ghttp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Glbs", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})

	//List
	Describe("List", func() {
		Context("When read of Glbs is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/zones/3fefc35e7decadb111dcf85d723a4f20/load_balancers"),
						ghttp.RespondWith(http.StatusOK, `
                           {
                              "result": [
                                {
                                  "description": "",
                                  "created_on": "2018-11-23T11:11:16.984683Z",
                                  "modified_on": "2018-11-23T11:11:24.727621Z",
                                  "id": "678106b2b5143fa9560e320961500f81",
                                  "proxied": true,
                                  "enabled": true,
                                  "name": "www.example.com",
                                  "session_affinity": "none",
                                  "fallback_pool": "4112ba6c2974ec43886f90736968e838",
                                  "default_pools": [
                                    "6563ebae141638f92ebbdc4a821bef8c",
                                    "4112ba6c2974ec43886f90736968e838"
                                  ],
                                  "pop_pools": {},
                                  "region_pools": {
                                    "EEU": [
                                      "4112ba6c2974ec43886f90736968e838"
                                    ],
                                    "ENAM": [
                                      "6563ebae141638f92ebbdc4a821bef8c"
                                    ],
                                    "WEU": [
                                      "4112ba6c2974ec43886f90736968e838"
                                    ],
                                    "WNAM": [
                                      "6563ebae141638f92ebbdc4a821bef8c"
                                    ]
                                  }
                                }
                              ],
                              "success": true,
                              "errors": [],
                              "messages": []
                            }
                        `),
					),
				)
			})

			It("should return Glb list", func() {
				target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				target2 := "3fefc35e7decadb111dcf85d723a4f20"
				myGlbs, err := newGlb(server.URL()).ListGlbs(target1, target2)
				Expect(myGlbs).ShouldNot(BeNil())
				for _, Glb := range myGlbs {
					Expect(err).NotTo(HaveOccurred())
					Expect(Glb.Id).Should(Equal("678106b2b5143fa9560e320961500f81"))
				}
			})
		})
		Context("When read of Glbs is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/zones/3fefc35e7decadb111dcf85d723a4f20/load_balancers"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve Glbs`),
					),
				)
			})

			It("should return error when Glb are retrieved", func() {
				target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				target2 := "3fefc35e7decadb111dcf85d723a4f20"
				myGlbPtr, err := newGlb(server.URL()).ListGlbs(target1, target2)
				myGlb := myGlbPtr
				Expect(err).To(HaveOccurred())
				Expect(myGlb).Should(BeNil())
			})
		})
	})

	//Delete
	Describe("Delete", func() {
		Context("When delete of Glb is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/zones/3fefc35e7decadb111dcf85d723a4f20/load_balancers/92859a0f6b4d3e55b953e0e29bb96338"),
						ghttp.RespondWith(http.StatusOK, `{                         
                        }`),
					),
				)
			})

			It("should delete Glb", func() {
				target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				target2 := "3fefc35e7decadb111dcf85d723a4f20"
				params := "92859a0f6b4d3e55b953e0e29bb96338"
				err := newGlb(server.URL()).DeleteGlb(target1, target2, params)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When Glb delete has failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/zones/3fefc35e7decadb111dcf85d723a4f20/load_balancers/92859a0f6b4d3e55b953e0e29bb96338"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to delete service key`),
					),
				)
			})

			It("should return error zone delete", func() {
				target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				target2 := "3fefc35e7decadb111dcf85d723a4f20"
				params := "92859a0f6b4d3e55b953e0e29bb96338"
				err := newGlb(server.URL()).DeleteGlb(target1, target2, params)
				Expect(err).To(HaveOccurred())
			})
		})
	})
	//Find
	Describe("Get", func() {
		Context("When read of Glb is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/zones/3fefc35e7decadb111dcf85d723a4f20/load_balancers/92859a0f6b4d3e55b953e0e29bb96338"),
						ghttp.RespondWith(http.StatusOK, `
                                                        {
                              "result": {
                                "description": "",
                                "created_on": "2018-11-23T11:11:16.984683Z",
                                "modified_on": "2018-11-23T11:11:24.727621Z",
                                "id": "678106b2b5143fa9560e320961500f81",
                                "proxied": true,
                                "enabled": true,
                                "name": "www.example.com",
                                "session_affinity": "none",
                                "fallback_pool": "4112ba6c2974ec43886f90736968e838",
                                "default_pools": [
                                  "6563ebae141638f92ebbdc4a821bef8c",
                                  "4112ba6c2974ec43886f90736968e838"
                                ],
                                "pop_pools": {},
                                "region_pools": {
                                  "EEU": [
                                    "4112ba6c2974ec43886f90736968e838"
                                  ],
                                  "ENAM": [
                                    "6563ebae141638f92ebbdc4a821bef8c"
                                  ],
                                  "WEU": [
                                    "4112ba6c2974ec43886f90736968e838"
                                  ],
                                  "WNAM": [
                                    "6563ebae141638f92ebbdc4a821bef8c"
                                  ]
                                }
                              },
                              "success": true,
                              "errors": [],
                              "messages": []
                            }
                        `),
					),
				)
			})

			It("should return Glb", func() {
				target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				target2 := "3fefc35e7decadb111dcf85d723a4f20"
				params := "92859a0f6b4d3e55b953e0e29bb96338"
				myGlbPtr, err := newGlb(server.URL()).GetGlb(target1, target2, params)
				myGlb := *myGlbPtr
				Expect(err).NotTo(HaveOccurred())
				Expect(myGlb).ShouldNot(BeNil())
				Expect(myGlb.Id).Should(Equal("678106b2b5143fa9560e320961500f81"))
			})
		})
		Context("When Glb get has failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/zones/3fefc35e7decadb111dcf85d723a4f20/load_balancers/92859a0f6b4d3e55b953e0e29bb96338"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve Glb`),
					),
				)
			})

			It("should return error when Glb is retrieved", func() {
				target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				target2 := "3fefc35e7decadb111dcf85d723a4f20"
				params := "92859a0f6b4d3e55b953e0e29bb96338"
				myGlbPtr, err := newGlb(server.URL()).GetGlb(target1, target2, params)
				myGlb := myGlbPtr
				Expect(err).To(HaveOccurred())
				Expect(myGlb).Should(BeNil())
			})
		})
	})
	Describe("Create", func() {
		Context("When creation is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/zones/3fefc35e7decadb111dcf85d723a4f20/load_balancers"),
						ghttp.VerifyJSON(`{"proxied": true, "name": "www.example.com", "session_affinity": "none", "fallback_pool": "4112ba6c2974ec43886f90736968e838", "default_pools": ["6563ebae141638f92ebbdc4a821bef8c", "4112ba6c2974ec43886f90736968e838"]}`),
						ghttp.RespondWith(http.StatusCreated, `
                           {
                              "result": {
                                "description": "",
                                "created_on": "2018-11-26T06:53:23.749062Z",
                                "modified_on": "2018-11-26T06:53:23.749062Z",
                                "id": "07085e0ea3c40225dcab6aff04cf64d9",
                                "ttl": 30,
                                "proxied": true,
                                "enabled": true,
                                "name": "www.example.com",
                                "session_affinity": "none",
                                "fallback_pool": "4112ba6c2974ec43886f90736968e838",
                                "default_pools": [
                                  "6563ebae141638f92ebbdc4a821bef8c",
                                  "4112ba6c2974ec43886f90736968e838"
                                ]
                              },
                              "success": true,
                              "errors": [],
                              "messages": []
                            }
                        `),
					),
				)
			})

			It("should return glb created", func() {
				params := GlbBody{
					Name:            "www.example.com",
					SessionAffinity: "none",
					DefaultPools: []string{
						"6563ebae141638f92ebbdc4a821bef8c",
						"4112ba6c2974ec43886f90736968e838",
					},
					FallbackPool: "4112ba6c2974ec43886f90736968e838",
					Proxied:      true,
				}
				target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				target2 := "3fefc35e7decadb111dcf85d723a4f20"
				myGlbPt, err := newGlb(server.URL()).CreateGlb(target1, target2, params)
				myGlb := *myGlbPt
				Expect(err).NotTo(HaveOccurred())
				Expect(myGlb).ShouldNot(BeNil())
				Expect(myGlb.Id).Should(Equal("07085e0ea3c40225dcab6aff04cf64d9"))
				Expect(myGlb.Name).Should(Equal("www.example.com"))
			})
		})
		Context("When creation is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/zones/3fefc35e7decadb111dcf85d723a4f20/load_balancers"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to create Glb`),
					),
				)
			})
			It("should return error during Glb creation", func() {
				params := GlbBody{
					Name:            "www.example.com",
					SessionAffinity: "none",
					DefaultPools: []string{
						"6563ebae141638f92ebbdc4a821bef8c",
						"4112ba6c2974ec43886f90736968e838",
					},
					FallbackPool: "4112ba6c2974ec43886f90736968e838",
					Proxied:      true,
				}
				target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				target2 := "3fefc35e7decadb111dcf85d723a4f20"
				myGlbPtr, err := newGlb(server.URL()).CreateGlb(target1, target2, params)
				myGlb := myGlbPtr
				Expect(err).To(HaveOccurred())
				Expect(myGlb).Should(BeNil())
			})
		})
	})

	Describe("Update", func() {
		Context("When update is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/zones/3fefc35e7decadb111dcf85d723a4f20/load_balancers"),
						ghttp.VerifyJSON(`{"proxied": true, "name": "www.example.com", "session_affinity": "none", "fallback_pool": "4112ba6c2974ec43886f90736968e838", "default_pools": ["6563ebae141638f92ebbdc4a821bef8c", "4112ba6c2974ec43886f90736968e838"]}`),
						ghttp.RespondWith(http.StatusCreated, `
                           {
                              "result": {
                                "description": "",
                                "created_on": "2018-11-26T06:53:23.749062Z",
                                "modified_on": "2018-11-26T06:53:23.749062Z",
                                "id": "07085e0ea3c40225dcab6aff04cf64d9",
                                "ttl": 30,
                                "proxied": true,
                                "enabled": true,
                                "name": "www.example.com",
                                "session_affinity": "none",
                                "fallback_pool": "4112ba6c2974ec43886f90736968e838",
                                "default_pools": [
                                  "6563ebae141638f92ebbdc4a821bef8c",
                                  "4112ba6c2974ec43886f90736968e838"
                                ]
                              },
                              "success": true,
                              "errors": [],
                              "messages": []
                            }
                        `),
					),
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/zones/3fefc35e7decadb111dcf85d723a4f20/load_balancers/07085e0ea3c40225dcab6aff04cf64d9"),
						ghttp.VerifyJSON(`{"proxied": true, "name": "www.example.com", "session_affinity": "none", "fallback_pool": "4112ba6c2974ec43886f90736968e888", "default_pools": ["6563ebae141638f92ebbdc4a821bef8c", "4112ba6c2974ec43886f90736968e838"]}`),
						ghttp.RespondWith(http.StatusCreated, `
                           {
                              "result": {
                                "description": "",
                                "created_on": "2018-11-26T06:53:23.749062Z",
                                "modified_on": "2018-11-26T06:53:23.749062Z",
                                "id": "07085e0ea3c40225dcab6aff04cf64d9",
                                "ttl": 30,
                                "proxied": true,
                                "enabled": true,
                                "name": "www.example.com",
                                "session_affinity": "none",
                                "fallback_pool": "4112ba6c2974ec43886f90736968e888",
                                "default_pools": [
                                  "6563ebae141638f92ebbdc4a821bef8c",
                                  "4112ba6c2974ec43886f90736968e838"
                                ]
                              },
                              "success": true,
                              "errors": [],
                              "messages": []
                            }
                        `),
					),
				)
			})

			It("should return glb updated", func() {
				params := GlbBody{
					Name:            "www.example.com",
					SessionAffinity: "none",
					DefaultPools: []string{
						"6563ebae141638f92ebbdc4a821bef8c",
						"4112ba6c2974ec43886f90736968e838",
					},
					FallbackPool: "4112ba6c2974ec43886f90736968e838",
					Proxied:      true,
				}
				target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				target2 := "3fefc35e7decadb111dcf85d723a4f20"
				myGlbPt, err := newGlb(server.URL()).CreateGlb(target1, target2, params)
				myGlb := *myGlbPt
				Expect(err).NotTo(HaveOccurred())
				Expect(myGlb).ShouldNot(BeNil())
				Expect(myGlb.Id).Should(Equal("07085e0ea3c40225dcab6aff04cf64d9"))
				Expect(myGlb.Name).Should(Equal("www.example.com"))
				params = GlbBody{
					Name:            "www.example.com",
					SessionAffinity: "none",
					DefaultPools: []string{
						"6563ebae141638f92ebbdc4a821bef8c",
						"4112ba6c2974ec43886f90736968e838",
					},
					FallbackPool: "4112ba6c2974ec43886f90736968e888",
					Proxied:      true,
				}
				target1 = "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				target2 = "3fefc35e7decadb111dcf85d723a4f20"
				target3 := "07085e0ea3c40225dcab6aff04cf64d9"
				myGlbPt, err = newGlb(server.URL()).UpdateGlb(target1, target2, target3, params)
				myGlb = *myGlbPt
				Expect(err).NotTo(HaveOccurred())
				Expect(myGlb).ShouldNot(BeNil())
				Expect(myGlb.Id).Should(Equal("07085e0ea3c40225dcab6aff04cf64d9"))
				Expect(myGlb.Name).Should(Equal("www.example.com"))
				Expect(myGlb.FallbackPool).Should(Equal("4112ba6c2974ec43886f90736968e888"))
			})
		})
		Context("When creation is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/zones/3fefc35e7decadb111dcf85d723a4f20/load_balancers/07085e0ea3c40225dcab6aff04cf64d9"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to create Glb`),
					),
				)
			})
			It("should return error during Glb creation", func() {
				params := GlbBody{
					Name:            "www.example.com",
					SessionAffinity: "none",
					DefaultPools: []string{
						"6563ebae141638f92ebbdc4a821bef8c",
						"4112ba6c2974ec43886f90736968e838",
					},
					FallbackPool: "4112ba6c2974ec43886f90736968e838",
					Proxied:      true,
				}
				target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				target2 := "3fefc35e7decadb111dcf85d723a4f20"
				target3 := "07085e0ea3c40225dcab6aff04cf64d9"
				myGlbPtr, err := newGlb(server.URL()).UpdateGlb(target1, target2, target3, params)
				myGlb := myGlbPtr
				Expect(err).To(HaveOccurred())
				Expect(myGlb).Should(BeNil())
			})
		})
	})
})

func newGlb(url string) Glbs {

	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}
	conf := sess.Config.Copy()
	conf.HTTPClient = bluemixHttp.NewHTTPClient(conf)
	conf.Endpoint = &url

	client := client.Client{
		Config:      conf,
		ServiceName: bluemix.CisService,
	}
	return newGlbAPI(&client)
}
