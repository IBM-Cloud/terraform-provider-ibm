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

var _ = Describe("Deployments", func() {
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
						ghttp.VerifyRequest(http.MethodGet, "/v4/ibm/deployments/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"),
						ghttp.RespondWith(http.StatusOK, `
                           {
                              "deployment": {
                                "id": "1f30581-54f8-41a4-8193-4a04cc022e9b-h",
                                "name": "1f30581-54f8-41a4-8193-4a04cc022e9b-h",
                                "type": "etcd",
                                "version": "3.2.7",
                                "admin_username": "admin"
                              }
                            }
                        `),
					),
				)
			})

			It("should return cdb", func() {
				target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				cdb, err := newCdb(server.URL()).GetCdb(target1)
				Expect(err).NotTo(HaveOccurred())
				Expect(cdb).ShouldNot(BeNil())
				Expect(cdb.Id).Should(Equal("1f30581-54f8-41a4-8193-4a04cc022e9b-h"))
				Expect(cdb.Name).Should(Equal("1f30581-54f8-41a4-8193-4a04cc022e9b-h"))
				Expect(cdb.Type).Should(Equal("etcd"))
				Expect(cdb.Version).Should(Equal("3.2.7"))
				Expect(cdb.AdminUser).Should(Equal("admin"))
			})
		})
		Context("When get is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v4/ibm/deployments/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to get cdb`),
					),
				)
			})

			It("should return error during get cdb", func() {
				target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				cdb, err := newCdb(server.URL()).GetCdb(target1)
				Expect(err).To(HaveOccurred())
				Expect(cdb.Id).Should(Equal(""))
			})
		})
	})
})

func newCdb(url string) Cdbs {

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
	return newCdbAPI(&client)
}
