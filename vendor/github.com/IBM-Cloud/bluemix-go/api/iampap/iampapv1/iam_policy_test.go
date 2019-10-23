package iampapv1

import (
	"log"
	"net/http"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/client"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/onsi/gomega/ghttp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Policy", func() {
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
						ghttp.VerifyRequest(http.MethodPost, "/acms/v1/scopes/a/f4755e41794cfa89cb078e865975f8e5/users/IBMid-270000W34J/policies"),
						ghttp.VerifyBody([]byte(`{"roles":[{"id":"crn:v1:bluemix:public:iam::::role:Viewer"}],"resources":[{"serviceName":"metrics-service"}]}`)),
						ghttp.RespondWith(http.StatusCreated, `{

  							"id": "81796686-5766-42ec-bd16-84894cc7f6ce",
  							"roles": [
    							{
      								"id": "crn:v1:bluemix:public:iam::::role:Viewer",
      								"displayName": "Viewer",
    	  							"description": "Viewers can take actions that do not change state (i.e. read only)."
    							}
							],
 						    "resources": [
    							{
      								"serviceName": "metrics-service",
      								"accountId": "f4755e41794cfa89cb078e865975f8e5"
    							}
  							],
  							"links": {
    							"href": "https://iampap.stage1.ng.bluemix.net/acms/v1/scopes/a%252ff4755e41794cfa89cb078e865975f8e5/users/IBMid-270000W34J/policies/81796686-5766-42ec-bd16-84894cc7f6ce",
    							"link": "self"
  							}
						
						}`),
					),
				)
			})

			It("should return Policy created", func() {
				var role = []Roles{
					Roles{
						ID: "crn:v1:bluemix:public:iam::::role:Viewer",
					},
				}
				var resource = []Resources{
					Resources{
						ServiceName: "metrics-service",
					},
				}
				var iamAccessInfo = AccessPolicyRequest{
					Roles:     role,
					Resources: resource,
				}
				myPolicy, _, err := newPolicy(server.URL()).Create("f4755e41794cfa89cb078e865975f8e5", "IBMid-270000W34J", iamAccessInfo)
				Expect(err).NotTo(HaveOccurred())
				Expect(myPolicy).ShouldNot(BeNil())
				Expect(myPolicy.ID).Should(Equal("81796686-5766-42ec-bd16-84894cc7f6ce"))
				Expect(myPolicy.Roles[0].ID).Should(Equal("crn:v1:bluemix:public:iam::::role:Viewer"))
				Expect(myPolicy.Roles[0].DisplayName).Should(Equal("Viewer"))
				Expect(myPolicy.Resources[0].ServiceName).Should(Equal("metrics-service"))
				Expect(myPolicy.Resources[0].AccountId).Should(Equal("f4755e41794cfa89cb078e865975f8e5"))
			})
		})
		Context("When creation is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/acms/v1/scopes/a/f4755e41794cfa89cb078e865975f8e5/users/IBMid-270000W34J/policies"),
						ghttp.VerifyBody([]byte(`{"roles":[{"id":"crn:v1:bluemix:public:iam::::role:Viewer"}],"resources":[{"serviceName":"metrics-service"}]}`)),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to create policy`),
					),
				)
			})

			It("should return error during policy creation", func() {
				var role = []Roles{
					Roles{
						ID: "crn:v1:bluemix:public:iam::::role:Viewer",
					},
				}
				var resource = []Resources{
					Resources{
						ServiceName: "metrics-service",
					},
				}
				var iamAccessInfo = AccessPolicyRequest{
					Roles:     role,
					Resources: resource,
				}
				myPolicy, _, err := newPolicy(server.URL()).Create("f4755e41794cfa89cb078e865975f8e5", "IBMid-270000W34J", iamAccessInfo)
				Expect(err).To(HaveOccurred())
				Expect(myPolicy).ShouldNot(BeNil())
			})
		})
	})

	Describe("Get", func() {
		Context("When get is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/acms/v1/scopes/a/f4755e41794cfa89cb078e865975f8e5/users/IBMid-270000W34J/policies/81796686-5766-42ec-bd16-84894cc7f6ce"),
						ghttp.RespondWith(http.StatusCreated, `{

  							"id": "81796686-5766-42ec-bd16-84894cc7f6ce",
  							"roles": [
    							{
      								"id": "crn:v1:bluemix:public:iam::::role:Viewer",
      								"displayName": "Viewer",
    	  							"description": "Viewers can take actions that do not change state (i.e. read only)."
    							}
							],
 						    "resources": [
    							{
      								"serviceName": "metrics-service",
      								"accountId": "f4755e41794cfa89cb078e865975f8e5"
    							}
  							],
  							"links": {
    							"href": "https://iampap.stage1.ng.bluemix.net/acms/v1/scopes/a%252ff4755e41794cfa89cb078e865975f8e5/users/IBMid-270000W34J/policies/81796686-5766-42ec-bd16-84894cc7f6ce",
    							"link": "self"
  							}
						
						}`),
					),
				)
			})

			It("should return Policy get", func() {
				myPolicy, err := newPolicy(server.URL()).Get("f4755e41794cfa89cb078e865975f8e5", "IBMid-270000W34J", "81796686-5766-42ec-bd16-84894cc7f6ce")
				Expect(err).NotTo(HaveOccurred())
				Expect(myPolicy).ShouldNot(BeNil())
				Expect(myPolicy.ID).Should(Equal("81796686-5766-42ec-bd16-84894cc7f6ce"))
				Expect(myPolicy.Roles[0].ID).Should(Equal("crn:v1:bluemix:public:iam::::role:Viewer"))
				Expect(myPolicy.Roles[0].DisplayName).Should(Equal("Viewer"))
				Expect(myPolicy.Resources[0].ServiceName).Should(Equal("metrics-service"))
				Expect(myPolicy.Resources[0].AccountId).Should(Equal("f4755e41794cfa89cb078e865975f8e5"))
			})
		})
		Context("When get is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/acms/v1/scopes/a/f4755e41794cfa89cb078e865975f8e5/users/IBMid-270000W34J/policies/81796686-5766-42ec-bd16-84894cc7f6ce"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to get policy`),
					),
				)
			})

			It("should return error during policy get", func() {
				myPolicy, err := newPolicy(server.URL()).Get("f4755e41794cfa89cb078e865975f8e5", "IBMid-270000W34J", "81796686-5766-42ec-bd16-84894cc7f6ce")
				Expect(err).To(HaveOccurred())
				Expect(myPolicy).ShouldNot(BeNil())
			})
		})
	})

	Describe("Update", func() {
		Context("When update is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, "/acms/v1/scopes/a/f4755e41794cfa89cb078e865975f8e5/users/IBMid-270000W34J/policies/81796686-5766-42ec-bd16-84894cc7f6ce"),
						ghttp.RespondWith(http.StatusCreated, `{

  							"id": "81796686-5766-42ec-bd16-84894cc7f6ce",
  							"roles": [
    							{
      								"id": "crn:v1:bluemix:public:iam::::role:Editor",
      								"displayName": "Editor",
    	  							"description": "Editor's can take actions that change state."
    							}
							],
 						    "resources": [
    							{
      								"serviceName": "metrics-service",
      								"accountId": "f4755e41794cfa89cb078e865975f8e5"
    							}
  							],
  							"links": {
    							"href": "https://iampap.stage1.ng.bluemix.net/acms/v1/scopes/a%252ff4755e41794cfa89cb078e865975f8e5/users/IBMid-270000W34J/policies/81796686-5766-42ec-bd16-84894cc7f6ce",
    							"link": "self"
  							}
						
						}`),
					),
				)
			})

			It("should return Policy updated", func() {
				var role = []Roles{
					Roles{
						ID: "crn:v1:bluemix:public:iam::::role:Editor",
					},
				}
				var resource = []Resources{
					Resources{
						ServiceName: "metrics-service",
					},
				}
				var iamAccessInfo = AccessPolicyRequest{
					Roles:     role,
					Resources: resource,
				}
				myPolicy, _, err := newPolicy(server.URL()).Update("f4755e41794cfa89cb078e865975f8e5", "IBMid-270000W34J", "81796686-5766-42ec-bd16-84894cc7f6ce", "W/'206-7VpPyt7UYHmZdu7/wv3cBg'", iamAccessInfo)
				Expect(err).NotTo(HaveOccurred())
				Expect(myPolicy).ShouldNot(BeNil())
				Expect(myPolicy.ID).Should(Equal("81796686-5766-42ec-bd16-84894cc7f6ce"))
				Expect(myPolicy.Roles[0].ID).Should(Equal("crn:v1:bluemix:public:iam::::role:Editor"))
				Expect(myPolicy.Roles[0].DisplayName).Should(Equal("Editor"))
				Expect(myPolicy.Resources[0].ServiceName).Should(Equal("metrics-service"))
				Expect(myPolicy.Resources[0].AccountId).Should(Equal("f4755e41794cfa89cb078e865975f8e5"))
			})
		})
		Context("When update is Failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, "/acms/v1/scopes/a/f4755e41794cfa89cb078e865975f8e5/users/IBMid-270000W34J/policies/81796686-5766-42ec-bd16-84894cc7f6ce"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to update policy`),
					),
				)
			})

			It("should return error during policy update", func() {
				var role = []Roles{
					Roles{
						ID: "crn:v1:bluemix:public:iam::::role:Editor",
					},
				}
				var resource = []Resources{
					Resources{
						ServiceName: "metrics-service",
					},
				}
				var iamAccessInfo = AccessPolicyRequest{
					Roles:     role,
					Resources: resource,
				}
				myPolicy, _, err := newPolicy(server.URL()).Update("f4755e41794cfa89cb078e865975f8e5", "IBMid-270000W34J", "81796686-5766-42ec-bd16-84894cc7f6ce", "W/'206-7VpPyt7UYHmZdu7/wv3cBg'", iamAccessInfo)
				Expect(err).To(HaveOccurred())
				Expect(myPolicy).ShouldNot(BeNil())
			})
		})
	})

	Describe("List", func() {
		Context("When List is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/acms/v1/scopes/a/f4755e41794cfa89cb078e865975f8e5/users/IBMid-270000W34J/policies"),
						ghttp.RespondWith(http.StatusCreated, `{
							"policies": [
    							{
      							"id": "a5ccf06f-c883-4806-a7ee-7eb2bf256d8e",
      							"roles": [
        							{
          								"id": "crn:v1:bluemix:public:iam::::role:Operator",
          								"displayName": "Operator",
         							 	"description": "Operators can take actions required to configure and operate resources."
        							}
     						 	],
     							"resources": [
        							{
          								"serviceName": "key-protect",
          								"accountId": "f4755e41794cfa89cb078e865975f8e5",
          								"region": "us-south"
        							}
      							],
      							"links": {
        							"href": "https://iampap.stage1.ng.bluemix.net/acms/v1/scopes/a%252ff4755e41794cfa89cb078e865975f8e5/users/IBMid-270000W34J/policies/a5ccf06f-c883-4806-a7ee-7eb2bf256d8e",
        							"link": "self"
      							}
    						},
    						{
      						"id": "d7344b3e-dcda-487d-b545-5b1b089a7e85",
      						"roles": [
        						{
         							 "id": "crn:v1:bluemix:public:iam::::role:Editor",
          							"displayName": "Editor",
          							"description": "Editors can take actions that can modify the state and create/delete sub-resources."
        						}
      						],
      						"resources": [
        						{
         							 "serviceName": "genesis",
          							"accountId": "f4755e41794cfa89cb078e865975f8e5"
       							 }
     						 ],
      						"links": {
        							"href": "https://iampap.stage1.ng.bluemix.net/acms/v1/scopes/a%252ff4755e41794cfa89cb078e865975f8e5/users/IBMid-270000W34J/policies/d7344b3e-dcda-487d-b545-5b1b089a7e85",
        							"link": "self"
      							}
    						}]	
     						
						}`),
					),
				)
			})

			It("should return Policy list", func() {
				myPolicy, err := newPolicy(server.URL()).List("f4755e41794cfa89cb078e865975f8e5", "IBMid-270000W34J")
				Expect(err).NotTo(HaveOccurred())
				Expect(myPolicy).ShouldNot(BeNil())
				Expect(myPolicy.Policies[0].ID).Should(Equal("a5ccf06f-c883-4806-a7ee-7eb2bf256d8e"))
				Expect(myPolicy.Policies[0].Roles[0].ID).Should(Equal("crn:v1:bluemix:public:iam::::role:Operator"))
				Expect(myPolicy.Policies[0].Roles[0].DisplayName).Should(Equal("Operator"))
				Expect(myPolicy.Policies[0].Resources[0].ServiceName).Should(Equal("key-protect"))
				Expect(myPolicy.Policies[0].Resources[0].AccountId).Should(Equal("f4755e41794cfa89cb078e865975f8e5"))
			})
		})
		Context("When list is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/acms/v1/scopes/a/f4755e41794cfa89cb078e865975f8e5/users/IBMid-270000W34J/policies"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to List`),
					),
				)
			})

			It("should return error during policy list", func() {
				myPolicy, err := newPolicy(server.URL()).List("f4755e41794cfa89cb078e865975f8e5", "IBMid-270000W34J")
				Expect(err).To(HaveOccurred())
				Expect(myPolicy).ShouldNot(BeNil())
			})
		})
	})

	Describe("Delete", func() {
		Context("When delete is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/acms/v1/scopes/a/f4755e41794cfa89cb078e865975f8e5/users/IBMid-270000W34J/policies/81796686-5766-42ec-bd16-84894cc7f6ce"),
						ghttp.RespondWith(http.StatusCreated, `{
						}`),
					),
				)
			})

			It("should return Policy", func() {
				err := newPolicy(server.URL()).Delete("f4755e41794cfa89cb078e865975f8e5", "IBMid-270000W34J", "81796686-5766-42ec-bd16-84894cc7f6ce")
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When delete is Failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/acms/v1/scopes/a/f4755e41794cfa89cb078e865975f8e5/users/IBMid-270000W34J/policies/81796686-5766-42ec-bd16-84894cc7f6ce"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to delete policy`),
					),
				)
			})

			It("should return error during policy delete", func() {
				err := newPolicy(server.URL()).Delete("f4755e41794cfa89cb078e865975f8e5", "IBMid-270000W34J", "81796686-5766-42ec-bd16-84894cc7f6ce")
				Expect(err).To(HaveOccurred())
			})
		})
	})
})

func newPolicy(url string) IAMPolicy {

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
	return newIAMPolicyAPI(&client)
}
