package accountv2

import (
	"log"
	"net/http"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/client"
	bluemixHttp "github.com/IBM-Cloud/bluemix-go/http"
	"github.com/IBM-Cloud/bluemix-go/session"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

// List Account by org

var _ = Describe("Accounts", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})

	Describe("fing account by org", func() {
		Context("Server return account by org", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/coe/v2/getaccounts"),
						ghttp.VerifyBody([]byte(`{"organizations_region":[{"guid":"5718c403-58aa-40fd-9854-4001901e83bb","region":"us-south"}]}`)),
						ghttp.RespondWith(http.StatusOK, `{  
   							"total_results":1,
   							"resources":[  
      						{  
         						"metadata":{  
            						"guid":"0e50220c5f6e1a13fdabe1c5cc87be32",
            						"url":"/coe/v2/accounts/0e50220c5f6e1a13fdabe1c5cc87be32",
            						"created_at":"2014-08-13T12:00:42.620Z",
            						"updated_at":"2017-04-28T23:33:44.415Z"
         						},
        						 "entity":{  
            					 	"name":"sakshi agarwal's Account",
            					 	"type":"TRIAL",
            					 	"state":"ACTIVE",
            					 	"owner":"a94cd8d9-5f77-4e2e-b2b4-adb23dd26117",
            					 	"owner_userid":"sakshiag@in.ibm.com",
            					 	"owner_unique_id":"270006H0V6",
            					 	"customer_id":"500198622",
            					 	"country_code":"USA",
            						"currency_code":"USD",
           							"billing_country_code":"USA",
            						"terms_and_conditions":{  

            						},
            						"tags":[  

            						],
            						"team_directory_enabled":true,
            						"organizations_region":[  
               						{  
                  						"guid":"5718c403-58aa-40fd-9854-4001901e83bb",
                  						"region":"us-south"
               						},
               						{  
                  						"guid":"900d4fda-0ed2-4679-a0ad-1ffmccp2b3ddd3",
                  						"region":"eu-gb"
               						}
            						],
           							"linkages":[  
               						{  
                  					"origin":"IMS",
                  					"state":"LINKABLE"
               						}
            						],
            						"bluemix_subscriptions":[  
               						{  
                  						"type":"TRIAL",
                  						"state":"ACTIVE",
                  						"payment_method":{  
                     						"type":"TRIAL_CREDIT",
                     						"started":"08/13/2014 12:00:39",
                     						"ended":"07/15/2018"
                  						},
                  						"subscription_id":"500096081",
                  						"part_number":"COE-Trial",
                  						"subscriptionTags":[  
                  						]
               						}
            						],
            						"subscription_id":"500096081",
            						"configuration_id":"",
            						"onboarded":-1
         							}
      							}
   								]
							}
							`),
					),
				)
			})

			It("should return account by org", func() {
				myaccount, err := newAccounts(server.URL()).FindByOrg("5718c403-58aa-40fd-9854-4001901e83bb", "us-south")
				Expect(err).To(Succeed())
				Expect(myaccount).ShouldNot(BeNil())
				Expect(myaccount.GUID).Should(Equal("0e50220c5f6e1a13fdabe1c5cc87be32"))
				Expect(myaccount.OwnerGUID).Should(Equal("a94cd8d9-5f77-4e2e-b2b4-adb23dd26117"))

			})

		})

		Context("Server return no account by org", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/coe/v2/getaccounts"),
						ghttp.VerifyBody([]byte(`{"organizations_region":[{"guid":"5718c403-58aa-40fd-9854-4001901e83bb","region":"xyz"}]}`)),
						ghttp.RespondWith(http.StatusOK, `{
							"total_results": 0,
							"resources": [
							]
															
						}`),
					),
				)
			})

			It("should return no account", func() {
				myaccount, err := newAccounts(server.URL()).FindByOrg("5718c403-58aa-40fd-9854-4001901e83bb", "xyz")
				Expect(err).To(HaveOccurred())
				Expect(myaccount).To(BeNil())
			})

		})
		Context("Server return error", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/coe/v2/getaccounts"),
						ghttp.VerifyBody([]byte(`{"organizations_region":[{"guid":"5718c403-58aa-40fd-9854-4001901e83bb","region":"us-south"}]}`)),
						ghttp.RespondWith(http.StatusInternalServerError, `{
															
						}`),
					),
				)
			})

			It("should return error", func() {
				myaccount, err := newAccounts(server.URL()).FindByOrg("5718c403-58aa-40fd-9854-4001901e83bb", "us-south")
				Expect(err).To(HaveOccurred())
				Expect(myaccount).To(BeNil())
			})

		})

	})
})

//List Accounts

var _ = Describe("List Accounts", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})

	Describe("List Accounts", func() {
		Context("Server return account", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/coe/v2/accounts"),
						ghttp.RespondWith(http.StatusOK, `{
   						 	"next_url":null,
   							"total_results":2,
   							"resources":[
     						 {
        						 "metadata":{
           						 "guid":"707977742b81692deed713861fd1a320",
            					 "url":"/coe/v2/accounts/707977742b81692deed713861fd1a320",
            					 "created_at":"2014-07-21T18:44:57.115Z",
            					 "updated_at":"2017-04-28T23:37:41.414Z"
         					 	 },
        					 	 "entity":{
            					 "name":"Richard Gebhardt's Account",
         					     "type":"TRIAL",
            					 "state":"ACTIVE",
            					 "owner":"62478eea-b929-4331-af4e-9acdf95e6f59",
            					 "owner_userid":"gebhardt@us.ibm.com",
            					 "owner_unique_id":"060001S23E",
            					 "customer_id":"500000244",
            				     "country_code":"USA",
           						 "currency_code":"USD",
            					 "billing_country_code":"USA",
            					 "terms_and_conditions":{
								 },
            					 "tags":[

            					 ],
            					 "team_directory_enabled":true,
           						 "organizations_region":[
             				     {
                  				 "guid":"3d6dfcd2-9d2a-408f-aa57-42383fcd0e3b",
                  				 "region":"us-south"
               					 },
               					 {
                  				 "guid":"ae27ccdd-7240-4f96-82b9-2b1713391c8e",
                  				 "region":"us-south"
               					 },
               					 {
                  				 "guid":"6d01a41a-c124-41c5-bb05-4c4a7509feab",
                  				 "region":"eu-gb"
               					 },
               					 {
                 				 "guid":"fe49a9a9-95c5-40a3-844a-1c6f77fbd4ef",
                  				 "region":"eu-gb"
               					 },
               					 {
                  				 "guid":"3391b99e-1163-4960-a3bc-be90e701bae4",
                  				 "region":"au-syd"
               					 },
              					 {
                  				 "guid":"443d63ce-9ef5-473c-807a-6d63ffa26f0a",
                  				 "region":"eu-de"
               					 },
               					 {
                  				 "guid":"bef1d514-2f57-44c2-84be-e4b07f7c3e0a",
                  			     "region":"eu-de"
               					 }
            					 ],
            				     "linkages":[
               					 {
                  				 	"origin":"IMS",
                  				 	"state":"LINKABLE"
               					 }
           						 ],
            					 "bluemix_subscriptions":[
                                  {
                  					"type":"TRIAL",
                  					"state":"ACTIVE",
                  					"payment_method":{
                     				"type":"TRIAL_CREDIT",
                     				"started":"05/28/2014 21:08:38",
                     				"ended":"07/15/2018"
                  				  },
                  				  "subscription_id":"500036292",
                  				  "part_number":"COE-Trial",
                  				  "subscriptionTags":[
								  ]
               					  }
            					  ],
            					  "subscription_id":"500036292",
            					  "configuration_id":"",
            					  "onboarded":-1
         						  }
     							  },
      						{
         						"metadata":{
            						"guid":"0e50220c5f6e1a13fdabe1c5cc87be32",
            						"url":"/coe/v2/accounts/0e50220c5f6e1a13fdabe1c5cc87be32",
            						"created_at":"2014-08-13T12:00:42.620Z",
            						"updated_at":"2017-04-28T23:33:44.415Z"
         						},
         						"entity":{
            						"name":"sakshi agarwal's Account",
            						"type":"TRIAL",
            						"state":"ACTIVE",
            						"owner":"a94cd8d9-5f77-4e2e-b2b4-adb23dd26117",
            						"owner_userid":"sakshiag@in.ibm.com",
            						"owner_unique_id":"270006H0V6",
            						"customer_id":"500198622",
            						"country_code":"USA",
            						"currency_code":"USD",
           						    "billing_country_code":"USA",
            						"terms_and_conditions":{
            						},
            					"tags":[
            					],
            					"team_directory_enabled":true,
           						"organizations_region":[
               					{
                 					 "guid":"5718c403-58aa-40fd-9854-4001901e83bb",
                  					 "region":"us-south"
               					},
              				    {
                  					"guid":"900d4fda-0ed2-4679-a0ad-1ffmccp2b3ddd3",
                  					"region":"eu-gb"
               					}
            					],
            					"linkages":[
               					{
                  					"origin":"IMS",
                  					"state":"LINKABLE"
              					}
            					],
            					"bluemix_subscriptions":[
               					{
                  					"type":"TRIAL",
                  					"state":"ACTIVE",
                  					"payment_method":{
                     					"type":"TRIAL_CREDIT",
                     					"started":"08/13/2014 12:00:39",
                     					"ended":"07/15/2018"
                  					},
                  					"subscription_id":"500096081",
                  					"part_number":"COE-Trial",
                  					"subscriptionTags":[
                  					]
               					}
           						],
            					"subscription_id":"500096081",
            					"configuration_id":"",
            					"onboarded":-1
         						}
      							}
   							]
							}
						    `),
					),
				)
			})

			It("should return accounts", func() {
				myaccounts, err := newAccounts(server.URL()).List()

				Expect(err).To(Succeed())
				Expect(myaccounts).ShouldNot(BeNil())
				Expect(len(myaccounts)).To(Equal(2))
				account1 := myaccounts[0]
				account2 := myaccounts[1]
				Expect(account1.GUID).Should(Equal("707977742b81692deed713861fd1a320"))
				Expect(account1.OwnerGUID).Should(Equal("62478eea-b929-4331-af4e-9acdf95e6f59"))
				Expect(account2.GUID).Should(Equal("0e50220c5f6e1a13fdabe1c5cc87be32"))
				Expect(account2.OwnerGUID).Should(Equal("a94cd8d9-5f77-4e2e-b2b4-adb23dd26117"))

			})

		})
		Context("Server return no accounts", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/coe/v2/accounts"),
						ghttp.RespondWith(http.StatusNotFound, `{
							"total_results": 0,
							"resources": [
							]
															
						}`),
					),
				)
			})

			It("should return no accounts", func() {
				myaccounts, err := newAccounts(server.URL()).List()
				Expect(err).To(HaveOccurred())
				Expect(len(myaccounts)).To(Equal(0))
			})

		})
		Context("Server return error", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/coe/v2/accounts"),
						ghttp.RespondWith(http.StatusInternalServerError, `{
															
						}`),
					),
				)
			})

			It("should return error", func() {
				myaccounts, err := newAccounts(server.URL()).List()
				Expect(err).To(HaveOccurred())
				Expect(myaccounts).To(BeNil())
			})

		})

	})
})

//list account by owner id

var _ = Describe("List Accounts by owner id", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})

	Describe("List Accounts by owner id", func() {
		Context("Server return account", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/coe/v2/accounts"),
						ghttp.RespondWith(http.StatusOK, `{
   						 	"next_url":null,
   							"total_results":1,
   							"resources":[
     						 {
         						"metadata":{
            						"guid":"0e50220c5f6e1a13fdabe1c5cc87be32",
            						"url":"/coe/v2/accounts/0e50220c5f6e1a13fdabe1c5cc87be32",
            						"created_at":"2014-08-13T12:00:42.620Z",
            						"updated_at":"2017-04-28T23:33:44.415Z"
         						},
         						"entity":{
            						"name":"sakshi agarwal's Account",
            						"type":"TRIAL",
            						"state":"ACTIVE",
            						"owner":"a94cd8d9-5f77-4e2e-b2b4-adb23dd26117",
            						"owner_userid":"sakshiag@in.ibm.com",
            						"owner_unique_id":"270006H0V6",
            						"customer_id":"500198622",
            						"country_code":"USA",
            						"currency_code":"USD",
           						    "billing_country_code":"USA",
            						"terms_and_conditions":{
            						},
            					"tags":[
            					],
            					"team_directory_enabled":true,
           						"organizations_region":[
               					{
                 					 "guid":"5718c403-58aa-40fd-9854-4001901e83bb",
                  					 "region":"us-south"
               					},
              				    {
                  					"guid":"900d4fda-0ed2-4679-a0ad-1ffmccp2b3ddd3",
                  					"region":"eu-gb"
               					}
            					],
            					"linkages":[
               					{
                  					"origin":"IMS",
                  					"state":"LINKABLE"
              					}
            					],
            					"bluemix_subscriptions":[
               					{
                  					"type":"TRIAL",
                  					"state":"ACTIVE",
                  					"payment_method":{
                     					"type":"TRIAL_CREDIT",
                     					"started":"08/13/2014 12:00:39",
                     					"ended":"07/15/2018"
                  					},
                  					"subscription_id":"500096081",
                  					"part_number":"COE-Trial",
                  					"subscriptionTags":[
                  					]
               					}
           						],
            					"subscription_id":"500096081",
            					"configuration_id":"",
            					"onboarded":-1
         						}
      							}
   							]
							}
						    `),
					),
				)
			})

			It("should return accounts by owner id", func() {
				myaccounts, err := newAccounts(server.URL()).FindByOwner("sakshiag@in.ibm.com")
				Expect(err).To(Succeed())
				Expect(myaccounts).ShouldNot(BeNil())
				Expect(myaccounts.GUID).Should(Equal("0e50220c5f6e1a13fdabe1c5cc87be32"))
				Expect(myaccounts.OwnerGUID).Should(Equal("a94cd8d9-5f77-4e2e-b2b4-adb23dd26117"))

			})

		})
		Context("Server return no accounts", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/coe/v2/accounts"),
						ghttp.RespondWith(http.StatusNotFound, `{
							"total_results": 0,
							"resources": [
							]
															
						}`),
					),
				)
			})

			It("should return no accounts", func() {
				myaccounts, err := newAccounts(server.URL()).FindByOwner("sakshiag@in.ibm.com")
				Expect(err).To(HaveOccurred())
				Expect(myaccounts).To(BeNil())
			})

		})
		Context("Server return error", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/coe/v2/accounts"),
						ghttp.RespondWith(http.StatusInternalServerError, `{
															
						}`),
					),
				)
			})

			It("should return error", func() {
				myaccounts, err := newAccounts(server.URL()).FindByOwner("sakshiag@in.ibm.com")
				Expect(err).To(HaveOccurred())
				Expect(myaccounts).To(BeNil())
			})

		})

	})
})

//get account by account id

var _ = Describe("Get Account by account id", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})

	Describe("Get Account by account id", func() {
		Context("Server return account", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/coe/v2/accounts/e9021a4d06e9b108b4a221a3cec47e3d"),
						ghttp.RespondWith(http.StatusOK, `{
							  "metadata": {
							    "guid": "e9021a4d06e9b108b4a221a3cec47e3d",
							    "url": "/coe/v2/accounts/e9021a4d06e9b108b4a221a3cec47e3d",
							    "created_at": "2015-09-02T10:35:17.288Z",
							    "updated_at": "2017-04-28T23:44:27.510Z"
							  },
							  "entity": {
							    "name": "Praveen G's Account",
							    "type": "TRIAL",
							    "state": "ACTIVE",
							    "owner": "a5d45507-836d-4609-8d6f-2972a05c9420",
							    "owner_userid": "praveek9@in.ibm.com",
							    "owner_unique_id": "2700067HF8",
							    "customer_id": "501280599",
							    "country_code": "IND",
							    "currency_code": "INR",
							    "billing_country_code": "IND",
							    "terms_and_conditions": {},
							    "tags": [],
							    "team_directory_enabled": true,
							    "organizations_region": [
							      {
							        "guid": "278f5173-c6b7-41b8-a2da-517daf27162c",
							        "region": "eu-gb"
							      },
							      {
							        "guid": "fc14269d-8ecd-403a-b3a8-0556372b4537",
							        "region": "us-south"
							      }
							    ],
							    "linkages": [
							      {
							        "origin": "IMS",
							        "state": "LINKABLE"
							      }
							    ],
							    "bluemix_subscriptions": [
							      {
							        "type": "TRIAL",
							        "state": "ACTIVE",
							        "payment_method": {
							          "type": "TRIAL_CREDIT",
							          "started": "09/02/2015 10:35:12",
							          "ended": "07/15/2018"
							        },
							        "subscription_id": "500711198",
							        "part_number": "COE-Trial",
							        "subscriptionTags": [],
							        "history": []
							      }
							    ],
							    "subscription_id": "500711198",
							    "configuration_id": "",
							    "onboarded": -1
							  }
							}

						    `),
					),
				)
			})

			It("should return account by account id", func() {
				myaccounts, err := newAccounts(server.URL()).Get("e9021a4d06e9b108b4a221a3cec47e3d")
				Expect(err).To(Succeed())
				Expect(myaccounts).ShouldNot(BeNil())
				Expect(myaccounts.Name).Should(Equal("Praveen G's Account"))
				Expect(myaccounts.GUID).Should(Equal("e9021a4d06e9b108b4a221a3cec47e3d"))

			})

		})
		Context("Server return error", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/coe/v2/accounts/e9021a4d06e9b108b4a221a3cec47e3d"),
						ghttp.RespondWith(http.StatusInternalServerError, `{
															
						}`),
					),
				)
			})

			It("should return error", func() {
				myaccounts, err := newAccounts(server.URL()).Get("e9021a4d06e9b108b4a221a3cec47e3d")
				Expect(err).To(HaveOccurred())
				Expect(myaccounts).To(BeNil())
			})

		})

	})
})

func newAccounts(url string) Accounts {

	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}
	conf := sess.Config.Copy()
	conf.HTTPClient = bluemixHttp.NewHTTPClient(conf)
	conf.Endpoint = &url

	client := client.Client{
		Config:      conf,
		ServiceName: bluemix.AccountService,
	}

	return newAccountAPI(&client)
}
