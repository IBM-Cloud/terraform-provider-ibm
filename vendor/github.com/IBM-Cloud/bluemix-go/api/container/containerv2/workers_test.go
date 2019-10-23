package containerv2

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

var _ = Describe("Workers", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})

	//ListByWorkerpool
	Describe("Get", func() {
		Context("When Get worker is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/vpc/getWorkers"),
						ghttp.RespondWith(http.StatusCreated, `[
							{
							  "flavor": "string",
							  "health": {
								"message": "string",
								"state": "string"
							  },
							  "id": "string",
							  "kubeVersion": {
								"actual": "string",
								"desired": "string",
								"eos": "string",
								"masterEOS": "string",
								"target": "string"
							  },
							  "lifecycle": {
								"actualState": "string",
								"desiredState": "string",
								"message": "string",
								"messageDate": "string",
								"messageDetails": "string",
								"messageDetailsDate": "string",
								"pendingOperation": "string",
								"reasonForDelete": "string"
							  },
							  "location": "string",
							  "networkInterfaces": [
								{
								  "cidr": "string",
								  "ipAddress": "string",
								  "primary": true,
								  "subnetID": "string"
								}
							  ],
							  "poolID": "string",
							  "poolName": "string"
							}
						  ]`),
					),
				)
			})

			It("should get workers in a cluster", func() {
				target := ClusterTargetHeader{}

				_, err := newWorker(server.URL()).ListByWorkerPool("aaa", "bbb", true, target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When get worker is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/vpc/getWorkers"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to get worker`),
					),
				)
			})

			It("should return error during get worker", func() {
				target := ClusterTargetHeader{}
				_, err := newWorker(server.URL()).ListByWorkerPool("aaa", "bbb", true, target)
				Expect(err).To(HaveOccurred())
			})
		})
	})

})

func newWorker(url string) Workers {

	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}
	conf := sess.Config.Copy()
	conf.HTTPClient = bluemixHttp.NewHTTPClient(conf)
	conf.Endpoint = &url

	client := client.Client{
		Config:      conf,
		ServiceName: bluemix.VpcContainerService,
	}
	return newWorkerAPI(&client)
}
