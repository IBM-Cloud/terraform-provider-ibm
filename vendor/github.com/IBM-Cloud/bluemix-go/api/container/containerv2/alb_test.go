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

var _ = Describe("Albs", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})

	//Create
	Describe("Create", func() {
		Context("When creating alb is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v2/alb/vpc/createAlb"),
						ghttp.VerifyJSON(`{"cluster":"345","type":"public","enableByDefault":true,"zone": "us-south-1"}`),
						ghttp.RespondWith(http.StatusCreated, `{}`),
					),
				)
			})

			It("should create Alb in a cluster", func() {
				target := ClusterTargetHeader{}
				params := AlbCreateReq{
					Cluster: "345", Type: "public", EnableByDefault: true, ZoneAlb: "us-south-1",
				}
				err := newAlbs(server.URL()).CreateAlb(params, target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When creating is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v2/alb/vpc/createAlb"),
						ghttp.VerifyJSON(`{"cluster":"345","type":"public","enableByDefault":true,"zone": "us-south-1"}
`),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to create alb`),
					),
				)
			})

			It("should return error during creating alb", func() {
				params := AlbCreateReq{
					Cluster: "345", Type: "public", EnableByDefault: true, ZoneAlb: "us-south-1",
				}
				target := ClusterTargetHeader{}
				err := newAlbs(server.URL()).CreateAlb(params, target)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	//Enable
	Describe("Enable", func() {
		Context("When Enabling alb is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v2/alb/vpc/enableAlb"),
						ghttp.VerifyJSON(`{"albBuild":"579","albID":"private-crbm64u3ed02o93vv36hb0-alb1","albType":"private","authBuild":"341","cluster":"345","createdDate":"","disableDeployment":true,"loadBalancerHostname":"","name":"","numOfInstances":"","resize":true,"state":"disabled","status":"","enable":true,"zone": "us-south-1"}`),
						ghttp.RespondWith(http.StatusCreated, `{}`),
					),
				)
			})

			It("should enable Alb in a cluster", func() {
				target := ClusterTargetHeader{}
				params := AlbConfig{
					AlbBuild: "579", AlbID: "private-crbm64u3ed02o93vv36hb0-alb1", AuthBuild: "341", Cluster: "345", CreatedDate: "", DisableDeployment: true, LoadBalancerHostname: "", AlbType: "private", Name: "", NumOfInstances: "", Resize: true, State: "disabled", Status: "", Enable: true, ZoneAlb: "us-south-1",
				}
				err := newAlbs(server.URL()).EnableAlb(params, target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When enabling is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v2/alb/vpc/enableAlb"),
						ghttp.VerifyJSON(`{"albBuild":"579","albID":"private-crbm64u3ed02o93vv36hb0-alb1","albType":"private","authBuild":"341","cluster":"345","createdDate":"","disableDeployment":true,"loadBalancerHostname":"","name":"","numOfInstances":"","resize":true,"state":"disabled","status":"","enable":true,"zone": "us-south-1"}
`),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to enable alb`),
					),
				)
			})

			It("should return error during enabling alb", func() {
				params := AlbConfig{
					AlbBuild: "579", AlbID: "private-crbm64u3ed02o93vv36hb0-alb1", AlbType: "private", AuthBuild: "341", CreatedDate: "", DisableDeployment: true, Enable: true, LoadBalancerHostname: "", Name: "", NumOfInstances: "", Resize: true, State: "disabled", Status: "", Cluster: "345", ZoneAlb: "us-south-1",
				}
				target := ClusterTargetHeader{}
				err := newAlbs(server.URL()).EnableAlb(params, target)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	//Disable
	Describe("Disable", func() {
		Context("When Disabling alb is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v2/alb/vpc/disableAlb"),
						ghttp.VerifyJSON(`{"albBuild":"579","albID":"private-crbm64u3ed02o93vv36hb0-alb1","albType":"private","authBuild":"341","cluster":"345","createdDate":"","disableDeployment":true,"loadBalancerHostname":"","name":"","numOfInstances":"","resize":true,"state":"disabled","status":"","enable":true,"zone": "us-south-1"}`),
						ghttp.RespondWith(http.StatusCreated, `{}`),
					),
				)
			})

			It("should disable Alb in a cluster", func() {
				target := ClusterTargetHeader{}
				params := AlbConfig{
					AlbBuild: "579", AlbID: "private-crbm64u3ed02o93vv36hb0-alb1", AlbType: "private", AuthBuild: "341", CreatedDate: "", DisableDeployment: true, Enable: true, LoadBalancerHostname: "", Name: "", NumOfInstances: "", Resize: true, State: "disabled", Status: "", Cluster: "345", ZoneAlb: "us-south-1",
				}
				err := newAlbs(server.URL()).DisableAlb(params, target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When disabling is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v2/alb/vpc/disableAlb"),
						ghttp.VerifyJSON(`{"albBuild":"579","albID":"private-crbm64u3ed02o93vv36hb0-alb1","albType":"private","authBuild":"341","cluster":"345","createdDate":"","disableDeployment":true,"loadBalancerHostname":"","name":"","numOfInstances":"","resize":true,"state":"disabled","status":"","enable":true,"zone": "us-south-1"}
`),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to disable alb`),
					),
				)
			})

			It("should return error during disabling alb", func() {
				params := AlbConfig{
					AlbBuild: "579", AlbID: "private-crbm64u3ed02o93vv36hb0-alb1", AlbType: "private", AuthBuild: "341", CreatedDate: "", DisableDeployment: true, Enable: true, LoadBalancerHostname: "", Name: "", NumOfInstances: "", Resize: true, State: "disabled", Status: "", Cluster: "345", ZoneAlb: "us-south-1",
				}
				target := ClusterTargetHeader{}
				err := newAlbs(server.URL()).DisableAlb(params, target)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	//GetAlbs
	Describe("Get", func() {
		Context("When Get Alb is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/alb/getAlb"),
						ghttp.RespondWith(http.StatusCreated, `{"albBuild": "string","albID": "string","albType": "string","authBuild": "string","cluster": "string","createdDate": "string","disableDeployment": true,"enable": true,"loadBalancerHostname": "string","name": "string","numOfInstances": "string","resize": true,"state": "string","status": "string","zone": "string"}`),
					),
				)
			})

			It("should get Alb in a cluster", func() {
				target := ClusterTargetHeader{}

				_, err := newAlbs(server.URL()).GetAlb("aaa", target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When get alb is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/alb/getAlb"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to get alb`),
					),
				)
			})

			It("should return error during get alb", func() {
				target := ClusterTargetHeader{}
				_, err := newAlbs(server.URL()).GetAlb("aaa", target)
				Expect(err).To(HaveOccurred())
			})
		})
	})

})

func newAlbs(url string) Alb {

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
	return newAlbAPI(&client)
}
