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

var _ = Describe("Pools", func() {
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
						ghttp.VerifyRequest(http.MethodPost, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/load_balancers/pools"),
						ghttp.RespondWith(http.StatusCreated, `
                           {
                              "result": {
                                "description": "",
                                "created_on": "2018-11-22T10:09:44.288581Z",
                                "modified_on": "2018-11-22T10:09:44.288581Z",
                                "id": "4112ba6c2974ec43886f90736968e838",
                                "enabled": true,
                                "minimum_origins": 1,
                                "monitor": "92859a0f6b4d3e55b953e0e29bb96338",
                                "name": "eu-pool",
                                "notification_email": "",
                                "check_regions": [
                                  "EEU"
                                ],
                                "origins": [
                                  {
                                    "name": "eu-origin1",
                                    "address": "150.0.0.1",
                                    "enabled": true,
                                    "weight": 1
                                  },
                                  {
                                    "name": "eu-origin2",
                                    "address": "150.0.0.2",
                                    "enabled": true,
                                    "weight": 1
                                  }
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

			It("should return zone created", func() {

				origins := []Origin{
					{Name: "eu-origin1", Address: "150.0.0.1", Enabled: true, Weight: 1},
					{Name: "eu-origin2", Address: "150.0.0.2", Enabled: true, Weight: 1},
				}
				checkRegions := []string{"EEU"}
				params := PoolBody{
					Name:         "eu-pool",
					Description:  "",
					Origins:      origins,
					CheckRegions: checkRegions,
					Enabled:      true,
					MinOrigins:   1,
					Monitor:      "92859a0f6b4d3e55b953e0e29bb96338",
					NotEmail:     "",
				}
				target := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				myPoolPtr, err := newPool(server.URL()).CreatePool(target, params)
				myPool := *myPoolPtr
				Expect(err).NotTo(HaveOccurred())
				Expect(myPool).ShouldNot(BeNil())
				Expect(myPool.Id).Should(Equal("4112ba6c2974ec43886f90736968e838"))
				Expect(myPool.Name).Should(Equal("eu-pool"))
			})
		})
		Context("When creation is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/load_balancers/pools"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to create Pool`),
					),
				)
			})

			It("should return error during Pool creation", func() {
				origins := []Origin{
					Origin{Name: "eu-origin1", Address: "150.0.0.1", Enabled: true, Weight: 1},
					Origin{Name: "eu-origin2", Address: "150.0.0.2", Enabled: true, Weight: 1},
				}
				checkRegions := []string{"EEU"}
				params := PoolBody{
					Name:         "eu-pool",
					Description:  "",
					Origins:      origins,
					CheckRegions: checkRegions,
					Enabled:      true,
					MinOrigins:   1,
					Monitor:      "92859a0f6b4d3e55b953e0e29bb96338",
					NotEmail:     "",
				}
				target := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				myPoolPtr, err := newPool(server.URL()).CreatePool(target, params)
				myPool := myPoolPtr
				Expect(err).To(HaveOccurred())
				Expect(myPool).Should(BeNil())
			})
		})
	})
	//List
	Describe("List", func() {
		Context("When read of Pools is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/load_balancers/pools"),
						ghttp.RespondWith(http.StatusOK, `
                            {
                  "result": [
                    {
                      "description": "",
                      "created_on": "2018-11-22T10:00:17.220005Z",
                      "modified_on": "2018-11-22T10:00:17.220005Z",
                      "id": "19901040048d70b15014330a6e252ba9",
                      "enabled": true,
                      "minimum_origins": 1,
                      "monitor": "92859a0f6b4d3e55b953e0e29bb96338",
                      "name": "ap-pool",
                      "notification_email": "",
                      "check_regions": [
                        "SEAS"
                      ],
                      "healthy": false,
                      "origins": [
                        {
                          "name": "ap-origin1",
                          "address": "100.0.0.1",
                          "enabled": true,
                          "weight": 1,
                          "healthy": false,
                          "failure_reason": "HTTP timeout occurred"
                        },
                        {
                          "name": "ap-origin2",
                          "address": "100.0.0.2",
                          "enabled": true,
                          "weight": 1,
                          "healthy": false,
                          "failure_reason": "HTTP timeout occurred"
                        }
                      ]
                    },
                    {
                      "description": "",
                      "created_on": "2018-11-20T07:20:06.057298Z",
                      "modified_on": "2018-11-22T10:00:52.867903Z",
                      "id": "19901040048d70b15014330a6e252ba9",
                      "enabled": true,
                      "minimum_origins": 1,
                      "monitor": "92859a0f6b4d3e55b953e0e29bb96338",
                      "name": "us-pool",
                      "notification_email": "",
                      "check_regions": [
                        "WNAM"
                      ],
                      "healthy": false,
                      "origins": [
                        {
                          "name": "us-origin1",
                          "address": "200.0.0.1",
                          "enabled": true,
                          "weight": 1,
                          "healthy": false,
                          "failure_reason": "Unspecified error"
                        },
                        {
                          "name": "us-origin2",
                          "address": "200.0.0.2",
                          "enabled": true,
                          "weight": 1,
                          "healthy": false,
                          "failure_reason": "Unspecified error"
                        }
                      ]
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

			It("should return Pool list", func() {
				target := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				myPools, err := newPool(server.URL()).ListPools(target)
				Expect(myPools).ShouldNot(BeNil())
				for _, Pool := range myPools {
					Expect(err).NotTo(HaveOccurred())
					Expect(Pool.Id).Should(Equal("19901040048d70b15014330a6e252ba9"))
				}
			})
		})
		Context("When read of Pools is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/load_balancers/pools"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve Pools`),
					),
				)
			})

			It("should return error when Pool are retrieved", func() {
				target := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				myPoolPtr, err := newPool(server.URL()).ListPools(target)
				myPool := myPoolPtr
				Expect(err).To(HaveOccurred())
				Expect(myPool).Should(BeNil())
			})
		})
	})
	//Delete
	Describe("Delete", func() {
		Context("When delete of Pool is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/load_balancers/pools/4112ba6c2974ec43886f90736968e838"),
						ghttp.RespondWith(http.StatusOK, `{                         
                        }`),
					),
				)
			})

			It("should delete Pool", func() {
				target := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				params := "4112ba6c2974ec43886f90736968e838"
				err := newPool(server.URL()).DeletePool(target, params)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When Pool delete has failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/load_balancers/pools/4112ba6c2974ec43886f90736968e838"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to delete service key`),
					),
				)
			})

			It("should return error zone delete", func() {
				target := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				params := "4112ba6c2974ec43886f90736968e838"
				err := newPool(server.URL()).DeletePool(target, params)
				Expect(err).To(HaveOccurred())
			})
		})
	})
	//Find
	Describe("Get", func() {
		Context("When read of Pool is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/load_balancers/pools/968826044d875e0841742564d2a4e994"),
						ghttp.RespondWith(http.StatusOK, `
                            {
                              "result": {
                                "description": "",
                                "created_on": "2018-11-20T07:20:06.057298Z",
                                "modified_on": "2018-11-22T10:09:10.629388Z",
                                "id": "968826044d875e0841742564d2a4e994",
                                "enabled": true,
                                "minimum_origins": 1,
                                "monitor": "92859a0f6b4d3e55b953e0e29bb96338",
                                "name": "us-pool",
                                "notification_email": "",
                                "check_regions": [
                                  "EEU"
                                ],
                                "healthy": false,
                                "origins": [
                                  {
                                    "name": "us-origin1",
                                    "address": "200.0.0.1",
                                    "enabled": true,
                                    "weight": 1,
                                    "healthy": false,
                                    "failure_reason": "HTTP timeout occurred"
                                  },
                                  {
                                    "name": "us-origin2",
                                    "address": "200.0.0.2",
                                    "enabled": true,
                                    "weight": 1,
                                    "healthy": false,
                                    "failure_reason": "HTTP timeout occurred"
                                  }
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

			It("should return Pool", func() {
				target := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				params := "968826044d875e0841742564d2a4e994"
				myPoolPtr, err := newPool(server.URL()).GetPool(target, params)
				myPool := myPoolPtr
				Expect(err).NotTo(HaveOccurred())
				Expect(myPool).ShouldNot(BeNil())
				Expect(myPool.Id).Should(Equal("968826044d875e0841742564d2a4e994"))
			})
		})
		Context("When Pool get has failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/load_balancers/pools/968826044d875e0841742564d2a4e994"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve Pool`),
					),
				)
			})

			It("should return error when Pool is retrieved", func() {
				target := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				params := "968826044d875e0841742564d2a4e994"
				myPoolPtr, err := newPool(server.URL()).GetPool(target, params)
				myPool := myPoolPtr
				Expect(err).To(HaveOccurred())
				Expect(myPool).Should(BeNil())
			})
		})
	})
	//Update
	Describe("Update", func() {
		Context("When Update is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/load_balancers/pools"),
						ghttp.RespondWith(http.StatusCreated, `
                           {
                              "result": {
                                "description": "",
                                "created_on": "2018-11-22T10:09:44.288581Z",
                                "modified_on": "2018-11-22T10:09:44.288581Z",
                                "id": "4112ba6c2974ec43886f90736968e838",
                                "enabled": true,
                                "minimum_origins": 1,
                                "monitor": "92859a0f6b4d3e55b953e0e29bb96338",
                                "name": "eu-pool",
                                "notification_email": "",
                                "check_regions": [
                                  "EEU"
                                ],
                                "origins": [
                                  {
                                    "name": "eu-origin1",
                                    "address": "150.0.0.1",
                                    "enabled": true,
                                    "weight": 1
                                  },
                                  {
                                    "name": "eu-origin2",
                                    "address": "150.0.0.2",
                                    "enabled": true,
                                    "weight": 1
                                  }
                                ]
                              },
                              "success": true,
                              "errors": [],
                              "messages": []
                            }
                        `),
					),
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/load_balancers/pools/4112ba6c2974ec43886f90736968e838"),
						ghttp.RespondWith(http.StatusCreated, `
                           {
                              "result": {
                                "description": "",
                                "created_on": "2018-11-22T10:09:44.288581Z",
                                "modified_on": "2018-11-22T10:09:44.288581Z",
                                "id": "4112ba6c2974ec43886f90736968e838",
                                "enabled": true,
                                "minimum_origins": 1,
                                "monitor": "92859a0f6b4d3e55b953e0e29bb96888",
                                "name": "eu-pool",
                                "notification_email": "",
                                "check_regions": [
                                  "EEU"
                                ],
                                "origins": [
                                  {
                                    "name": "eu-origin1",
                                    "address": "150.0.0.1",
                                    "enabled": true,
                                    "weight": 1
                                  },
                                  {
                                    "name": "eu-origin2",
                                    "address": "150.0.0.2",
                                    "enabled": true,
                                    "weight": 1
                                  }
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

			It("should return zone created", func() {

				origins := []Origin{
					{Name: "eu-origin1", Address: "150.0.0.1", Enabled: true, Weight: 1},
					{Name: "eu-origin2", Address: "150.0.0.2", Enabled: true, Weight: 1},
				}
				checkRegions := []string{"EEU"}
				params := PoolBody{
					Name:         "eu-pool",
					Description:  "",
					Origins:      origins,
					CheckRegions: checkRegions,
					Enabled:      true,
					MinOrigins:   1,
					Monitor:      "92859a0f6b4d3e55b953e0e29bb96338",
					NotEmail:     "",
				}
				target := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				myPoolPtr, err := newPool(server.URL()).CreatePool(target, params)
				myPool := *myPoolPtr
				Expect(err).NotTo(HaveOccurred())
				Expect(myPool).ShouldNot(BeNil())
				Expect(myPool.Id).Should(Equal("4112ba6c2974ec43886f90736968e838"))
				Expect(myPool.Name).Should(Equal("eu-pool"))
				params = PoolBody{
					Name:         "eu-pool",
					Description:  "",
					Origins:      origins,
					CheckRegions: checkRegions,
					Enabled:      true,
					MinOrigins:   1,
					Monitor:      "92859a0f6b4d3e55b953e0e29bb96888",
					NotEmail:     "",
				}
				target = "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				poolId := "4112ba6c2974ec43886f90736968e838"
				myPoolPtr, err = newPool(server.URL()).UpdatePool(target, poolId, params)
				myPool = *myPoolPtr
				Expect(err).NotTo(HaveOccurred())
				Expect(myPool).ShouldNot(BeNil())
				Expect(myPool.Id).Should(Equal("4112ba6c2974ec43886f90736968e838"))
				Expect(myPool.Name).Should(Equal("eu-pool"))
				Expect(myPool.Monitor).Should(Equal("92859a0f6b4d3e55b953e0e29bb96888"))
			})
		})
		Context("When Update is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/load_balancers/pools/4112ba6c2974ec43886f90736968e838"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to create Pool`),
					),
				)
			})

			It("should return error during Pool creation", func() {
				origins := []Origin{
					Origin{Name: "eu-origin1", Address: "150.0.0.1", Enabled: true, Weight: 1},
					Origin{Name: "eu-origin2", Address: "150.0.0.2", Enabled: true, Weight: 1},
				}
				checkRegions := []string{"EEU"}
				params := PoolBody{
					Name:         "eu-pool",
					Description:  "",
					Origins:      origins,
					CheckRegions: checkRegions,
					Enabled:      true,
					MinOrigins:   1,
					Monitor:      "92859a0f6b4d3e55b953e0e29bb96338",
					NotEmail:     "",
				}
				target := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				poolId := "4112ba6c2974ec43886f90736968e838"
				myPoolPtr, err := newPool(server.URL()).UpdatePool(target, poolId, params)
				myPool := myPoolPtr
				Expect(err).To(HaveOccurred())
				Expect(myPool).Should(BeNil())
			})
		})
	})
})

func newPool(url string) Pools {

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
	return newPoolAPI(&client)
}
