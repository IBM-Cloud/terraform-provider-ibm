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

var _ = Describe("Subnets", func() {
	var server *ghttp.Server
	Describe("Add", func() {
		Context("When adding a subnet is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, "/v1/clusters/test/subnets/1109876"),
						ghttp.RespondWith(http.StatusCreated, `{}`),
					),
				)
			})

			It("should return subnet added to cluster", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				err := newSubnet(server.URL()).AddSubnet("test", "1109876", target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When adding subnet is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, "/v1/clusters/test/subnets/1109876"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to add subnet to cluster`),
					),
				)
			})

			It("should return error during add subnet to cluster", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				err := newSubnet(server.URL()).AddSubnet("test", "1109876", target)
				Expect(err).To(HaveOccurred())
			})
		})
	})
	//List
	Describe("List", func() {
		Context("When retrieving available subnets is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/subnets"),
						ghttp.RespondWith(http.StatusOK, `[{
						"ID": "535642",
						"Type": "private",     
						"VlanID": "1565297",
						"IPAddresses": ["10.98.25.2","10.98.25.3","10.98.25.4"],
						"Properties": {
						"CIDR": "26 ",
						"NetworkIdentifier":"10.130.229.64",
						"Note": "",
						"SubnetType":"additional_primary",
						"DisplayLabel":"",
						"Gateway":""
						}
						}]`),
					),
				)
			})

			It("should return available subnets ", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				subnets, err := newSubnet(server.URL()).List(target)
				Expect(err).NotTo(HaveOccurred())
				Expect(subnets).ShouldNot(BeNil())
				for _, sObj := range subnets {
					Expect(sObj).ShouldNot(BeNil())
					Expect(sObj.ID).Should(Equal("535642"))
					Expect(sObj.Type).Should(Equal("private"))
				}
			})
		})
		Context("When retrieving available subnets is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/subnets"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve subnets`),
					),
				)
			})

			It("should return error during retrieveing subnets", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				subnets, err := newSubnet(server.URL()).List(target)
				Expect(err).To(HaveOccurred())
				Expect(subnets).Should(BeNil())
			})
		})
	})

	//AddClusterUserSubnet
	Describe("AddClusterUserSubnet", func() {
		Context("When adding a user subnet is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/clusters/test/usersubnets"),
						ghttp.RespondWith(http.StatusCreated, `{}`),
					),
				)
			})

			It("should return user subnet added to cluster", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				userSubnet := UserSubnet{
					CIDR:   "9.0.0.1",
					VLANID: "1156748",
				}
				err := newSubnet(server.URL()).AddClusterUserSubnet("test", userSubnet, target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When adding user subnet is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/clusters/test/usersubnets"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to add user subnet to cluster`),
					),
				)
			})

			It("should return error during adding user subnet to cluster", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				userSubnet := UserSubnet{
					CIDR:   "9.0.0.1",
					VLANID: "1156748",
				}
				err := newSubnet(server.URL()).AddClusterUserSubnet("test", userSubnet, target)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	//ListClusterUserSubnets
	Describe("ListClusterUserSubnets", func() {
		Context("When retrieving available user subnets is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/clusters/test/usersubnets"),
						ghttp.RespondWith(http.StatusOK, `[{
						  "ID": "177453",
							"Subnets": [
							{
							"Cidr": "159.8.226.208/29",
							"ID": "1541737",
							"Ips": ["159.8.226.210"],
							"Is_ByOIP": false,
							"Is_Public": true
							}]
							}]`),
					),
				)
			})

			It("should return available subnets ", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				subnets, err := newSubnet(server.URL()).ListClusterUserSubnets("test", target)
				Expect(err).NotTo(HaveOccurred())
				Expect(subnets).ShouldNot(BeNil())
				for _, sObj := range subnets {
					Expect(sObj).ShouldNot(BeNil())
				}
			})
		})
		Context("When retrieving available user subnets is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/clusters/test/usersubnets"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve user subnets`),
					),
				)
			})

			It("should return error during retrieveing user subnets", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				subnets, err := newSubnet(server.URL()).ListClusterUserSubnets("test", target)
				Expect(err).To(HaveOccurred())
				Expect(subnets).Should(BeNil())
			})
		})
	})

	//DeleteClusterUserSubnet
	Describe("DeleteClusterUserSubnet", func() {
		Context("When delete of cluster user subnet is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/clusters/test/usersubnets/110976/vlans/174991"),
						ghttp.RespondWith(http.StatusOK, `{							
						}`),
					),
				)
			})

			It("should delete cluster", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				err := newSubnet(server.URL()).DeleteClusterUserSubnet("test", "110976", "174991", target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When cluster user subnet delete is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/clusters/test/usersubnets/110976/vlans/174991"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to delete cluster user subnet`),
					),
				)
			})

			It("should return error deleting user subnet", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				err := newSubnet(server.URL()).DeleteClusterUserSubnet("test", "110976", "174991", target)
				Expect(err).To(HaveOccurred())
			})
		})
	})
})

func newSubnet(url string) Subnets {

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
	return newSubnetAPI(&client)
}
