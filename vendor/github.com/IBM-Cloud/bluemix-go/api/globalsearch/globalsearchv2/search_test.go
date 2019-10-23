package globalsearchv2

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
						ghttp.VerifyRequest(http.MethodPost, "/v2/resources/search"),
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
				token := ""
				fields := []string{""}
				query := ""
				searchBody := SearchBody{
					Query:  query,
					Fields: fields,
					Token:  token,
				}
				searchResult, err := newSearch(server.URL()).PostQuery(searchBody)
				Expect(err).NotTo(HaveOccurred())
				Expect(searchResult).ShouldNot(BeNil())
				//Expect(searchResult.Items).Should(Equal("5abb6a7d11a1a5001479a0ac"))

			})
		})
		Context("When PostQuery is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v2/resources/search"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to get query results`),
					),
				)
			})

			It("should return error during post query", func() {
				token := ""
				fields := []string{""}
				query := ""
				searchBody := SearchBody{
					Query:  query,
					Fields: fields,
					Token:  token,
				}
				searchResult, err := newSearch(server.URL()).PostQuery(searchBody)
				Expect(err).To(HaveOccurred())
				Expect(searchResult.Items).Should(Equal(""))
			})
		})
	})
})

func newSearch(url string) Searches {

	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}
	conf := sess.Config.Copy()
	conf.HTTPClient = bluemixHttp.NewHTTPClient(conf)
	conf.Endpoint = &url

	client := client.Client{
		Config:      conf,
		ServiceName: bluemix.GlobalSearchService,
	}
	return newSearchAPI(&client)
}
