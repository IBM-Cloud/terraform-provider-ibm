package icdv4

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

var _ = Describe("Users", func() {
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
						ghttp.VerifyRequest(http.MethodPost, "/v4/ibm/deployments/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/users"),
						ghttp.RespondWith(http.StatusCreated, `
                           {
                            "task": {
                              "id": "5abb6a7d11a1a5001479a0ac",
                              "description": "Creating user for database",
                              "status": "running",
                              "deployment_id": "59b14b19874a1c0018009482",
                              "progress_percent": 5,
                              "created_at": "2018-03-28T10:21:30Z"
                            }
                          }
                        `),
					),
				)
			})

			It("should return user created", func() {
				user := User{
					UserName: "admin1",
					Password: "password",
				}
				params := UserReq{
					User: user,
				}
				target := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				myTask, err := newUser(server.URL()).CreateUser(target, params)
				Expect(err).NotTo(HaveOccurred())
				Expect(myTask).ShouldNot(BeNil())
				Expect(myTask.Id).Should(Equal("5abb6a7d11a1a5001479a0ac"))
				Expect(myTask.Description).Should(Equal("Creating user for database"))
				Expect(myTask.Status).Should(Equal("running"))
				Expect(myTask.DeploymentId).Should(Equal("59b14b19874a1c0018009482"))
				Expect(myTask.ProgressPercent).Should(Equal(5))
				Expect(myTask.CreatedAt).Should(Equal("2018-03-28T10:21:30Z"))
			})
		})
		Context("When creation is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v4/ibm/deployments/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/users"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to create User`),
					),
				)
			})

			It("should return error during User creation", func() {
				user := User{
					UserName: "admin1",
					Password: "password",
				}
				params := UserReq{
					User: user,
				}
				target := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				myTask, err := newUser(server.URL()).CreateUser(target, params)
				Expect(err).To(HaveOccurred())
				Expect(myTask.Id).Should(Equal(""))
			})
		})
	})
	Describe("Delete", func() {
		Context("When deletion is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v4/ibm/deployments/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/users/admin1"),
						ghttp.RespondWith(http.StatusOK, `
                           {
                              "task": {
                                "id": "5abb6a7d11a1a5001479a0ae",
                                "description": "Deleting user from database",
                                "status": "running",
                                "deployment_id": "59b14b19874a1c0018009482",
                                "progress_percent": 10,
                                "created_at": "2018-03-28T10:23:30Z"
                              }
                            }
                        `),
					),
				)
			})

			It("should return user deleted", func() {
				target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				target2 := "admin1"
				myTask, err := newUser(server.URL()).DeleteUser(target1, target2)
				Expect(err).NotTo(HaveOccurred())
				Expect(myTask).ShouldNot(BeNil())
				Expect(myTask.Id).Should(Equal("5abb6a7d11a1a5001479a0ae"))
				Expect(myTask.Description).Should(Equal("Deleting user from database"))
				Expect(myTask.Status).Should(Equal("running"))
				Expect(myTask.DeploymentId).Should(Equal("59b14b19874a1c0018009482"))
				Expect(myTask.ProgressPercent).Should(Equal(10))
				Expect(myTask.CreatedAt).Should(Equal("2018-03-28T10:23:30Z"))
			})
		})
		Context("When deletion is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v4/ibm/deployments/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/users/admin1"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to delete User`),
					),
				)
			})

			It("should return error during User deletion", func() {
				target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				target2 := "admin1"
				myTask, err := newUser(server.URL()).DeleteUser(target1, target2)
				Expect(err).To(HaveOccurred())
				Expect(myTask.Id).Should(Equal(""))
			})
		})
	})
	Describe("Update", func() {
		Context("When update is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPatch, "/v4/ibm/deployments/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/users/admin1"),
						ghttp.RespondWith(http.StatusOK, `
                           {
                                "task": {
                                "id": "5abb6a7d11a1a5001479a0ad",
                                "description": "Setting user password for database",
                                "status": "running",
                                "deployment_id": "59b14b19874a1c0018009482",
                                "progress_percent": 5,
                                "created_at": "2018-03-28T10:22:30Z"
                                }
                            }
                        `),
					),
				)
			})

			It("should return user updated", func() {
				target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				target2 := "admin1"
				user := User{
					Password: "password",
				}
				params := UserReq{
					User: user,
				}
				myTask, err := newUser(server.URL()).UpdateUser(target1, target2, params)
				Expect(err).NotTo(HaveOccurred())
				Expect(myTask).ShouldNot(BeNil())
				Expect(myTask.Id).Should(Equal("5abb6a7d11a1a5001479a0ad"))
				Expect(myTask.Description).Should(Equal("Setting user password for database"))
				Expect(myTask.Status).Should(Equal("running"))
				Expect(myTask.DeploymentId).Should(Equal("59b14b19874a1c0018009482"))
				Expect(myTask.ProgressPercent).Should(Equal(5))
				Expect(myTask.CreatedAt).Should(Equal("2018-03-28T10:22:30Z"))
			})
		})
		Context("When update is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPatch, "/v4/ibm/deployments/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/users/admin1"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to update User`),
					),
				)
			})

			It("should return error during User update", func() {
				target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				target2 := "admin1"
				user := User{
					Password: "password",
				}
				params := UserReq{
					User: user,
				}
				myTask, err := newUser(server.URL()).UpdateUser(target1, target2, params)
				Expect(err).To(HaveOccurred())
				Expect(myTask.Id).Should(Equal(""))
			})
		})
	})
})

func newUser(url string) Users {

	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}
	conf := sess.Config.Copy()
	conf.HTTPClient = bluemixHttp.NewHTTPClient(conf)
	conf.Endpoint = &url

	client := client.Client{
		Config:      conf,
		ServiceName: bluemix.ICDService,
	}
	return newUsersAPI(&client)
}
