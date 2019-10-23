package globaltaggingv3

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
	Describe("PostQuery", func() {
		Context("When PostQuery is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v3/tags"),
						ghttp.RespondWith(http.StatusOK, `
                           {
                            "items": {
                            }
                          }
                        `),
					),
				)
			})

			It("should return query results", func() {
				resourceID := "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/4ea1882a2d3401ed1e459979941966ea:2ede6105-d368-4f20-b2a3-2e27de37f0da::"
				taggingResult, err := newTagging(server.URL()).GetTags(resourceID)
				Expect(err).NotTo(HaveOccurred())
				Expect(taggingResult).ShouldNot(BeNil())
				//Expect(taggingResult.Items).Should(Equal("5abb6a7d11a1a5001479a0ac"))

			})
		})
		Context("When PostQuery is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v3/tags"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to get query results`),
					),
				)
			})

			It("should return error during post query", func() {
				resourceID := "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/4ea1882a2d3401ed1e459979941966ea:2ede6105-d368-4f20-b2a3-2e27de37f0da::"
				taggingResult, err := newTagging(server.URL()).GetTags(resourceID)
				Expect(err).To(HaveOccurred())
				Expect(taggingResult.Items).Should(Equal(""))
			})
		})
	})
})

func newTagging(url string) Tags {

	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}
	conf := sess.Config.Copy()
	conf.HTTPClient = bluemixHttp.NewHTTPClient(conf)
	conf.Endpoint = &url

	client := client.Client{
		Config:      conf,
		ServiceName: bluemix.GlobalTaggingService,
	}
	return newTaggingAPI(&client)
}
