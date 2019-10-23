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

var _ = Describe("Tasks", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})
	Describe("Get", func() {
		Context("When get is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v4/ibm/tasks/5abb6a7d11a1a5001479a0ac"),
						ghttp.RespondWith(http.StatusOK, `
                           {
                            "task": {
                              "id": "5abb6a7d11a1a5001479a0ac",
                              "description": "Creating task for database",
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

			It("should return task", func() {
				target := "5abb6a7d11a1a5001479a0ac"
				task, err := newTask(server.URL()).GetTask(target)
				Expect(err).NotTo(HaveOccurred())
				Expect(task).ShouldNot(BeNil())
				Expect(task.Id).Should(Equal("5abb6a7d11a1a5001479a0ac"))
				Expect(task.Description).Should(Equal("Creating task for database"))
				Expect(task.Status).Should(Equal("running"))
				Expect(task.DeploymentId).Should(Equal("59b14b19874a1c0018009482"))
				Expect(task.ProgressPercent).Should(Equal(5))
				Expect(task.CreatedAt).Should(Equal("2018-03-28T10:21:30Z"))
			})
		})
		Context("When get is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v4/ibm/tasks/5abb6a7d11a1a5001479a0ac"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to get task`),
					),
				)
			})

			It("should return error during get task", func() {
				target := "5abb6a7d11a1a5001479a0ac"
				task, err := newTask(server.URL()).GetTask(target)
				Expect(err).To(HaveOccurred())
				Expect(task.Id).Should(Equal(""))
			})
		})
	})
})

func newTask(url string) Tasks {

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
	return newTaskAPI(&client)
}
