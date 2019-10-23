package containerv1

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

var _ = Describe("Webhooks", func() {
	var server *ghttp.Server
	Describe("Add", func() {
		Context("When adding a webhook is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/clusters/test/webhooks"),
						ghttp.RespondWith(http.StatusCreated, `{}`),
					),
				)
			})

			It("should return webhook added to cluster", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				params := WebHook{
					Level: "Warning", Type: "slack", URL: "http://slack.com/frwf-grev",
				}
				err := newWebhook(server.URL()).Add("test", params, target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When adding webhook is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/clusters/test/webhooks"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to add webhook to cluster`),
					),
				)
			})

			It("should return error during add webhook to cluster", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				params := WebHook{
					Level: "Warning", Type: "slack", URL: "http://slack.com/frwf-grev",
				}
				err := newWebhook(server.URL()).Add("test", params, target)
				Expect(err).To(HaveOccurred())
			})
		})
	})
	//List
	Describe("List", func() {
		Context("When retrieving available webhooks is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/clusters/test/webhooks"),
						ghttp.RespondWith(http.StatusOK, `
						[{"Level": "Warning", "Type": "slack", "URL": "http://slack.com/frwf-grev"}]
						`),
					),
				)
			})

			It("should return available webhooks ", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				webhooks, err := newWebhook(server.URL()).List("test", target)
				Expect(err).NotTo(HaveOccurred())
				Expect(webhooks).ShouldNot(BeNil())
				for _, wObj := range webhooks {
					Expect(wObj).ShouldNot(BeNil())
					Expect(wObj.Type).Should(Equal("slack"))
				}
			})
		})
		Context("When retrieving available webhooks is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/clusters/test/webhooks"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve webhooks`),
					),
				)
			})

			It("should return error during retrieveing webhooks", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				webhook, err := newWebhook(server.URL()).List("test", target)
				Expect(err).To(HaveOccurred())
				Expect(webhook).Should(BeNil())
			})
		})
	})

})

func newWebhook(url string) Webhooks {

	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}
	conf := sess.Config.Copy()
	conf.HTTPClient = bluemixHttp.NewHTTPClient(conf)
	conf.Endpoint = &url

	client := client.Client{
		Config:      conf,
		ServiceName: bluemix.MccpService,
	}
	return newWebhookAPI(&client)
}
