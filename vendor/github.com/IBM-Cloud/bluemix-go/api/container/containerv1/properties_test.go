package containerv1

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

var _ = Describe("Vlans", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})

	//List
	Describe("List", func() {
		Context("When read of vlans is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/datacenters/dal10/vlans"),
						ghttp.RespondWith(http.StatusOK, `[
            {
              "id": "12345",
              "type": "private",
              "properties": {
                "name": "",
                "note": "",
                "primary_router": "something.dal10",
                "vlan_number": "1111",
                "vlan_type": "standard",
                "location": "11",
                "local_disk_storage_capability": "true",
                "san_storage_capability": "true"
              }
            }]`),
					),
				)
			})

			It("should return cluster list", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}

				vlans, err := newVlan(server.URL()).List("dal10", target)
				Expect(vlans).ShouldNot(BeNil())
				for _, vlan := range vlans {
					Expect(err).NotTo(HaveOccurred())
					Expect(vlan.ID).Should(Equal("12345"))
				}
			})
		})
		Context("When read of vlans is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/datacenters/fakedc/vlans"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve vlans`),
					),
				)
			})

			It("should return error when cluster are retrieved", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}

				vlans, err := newVlan(server.URL()).List("fakedc", target)
				Expect(err).To(HaveOccurred())
				Expect(vlans).Should(BeNil())
			})
		})
	})
	//
})

func newVlan(url string) Vlans {

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
	return newVlanAPI(&client)
}
