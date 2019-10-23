package csev2

import (
	"fmt"
	"log"
	"net/http"

	"github.com/IBM-Cloud/bluemix-go"
	bluemixHttp "github.com/IBM-Cloud/bluemix-go/http"

	"github.com/IBM-Cloud/bluemix-go/client"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/onsi/gomega/ghttp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ServiceEndpoints", func() {
	var server *ghttp.Server
	srvID := "test-srv-id"

	AfterEach(func() {
		server.Close()
	})

	Describe("GetServiceEndpoint", func() {
		Context("When service with srvid being srvID exists", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, fmt.Sprintf("/v2/serviceendpoint/%s", srvID)),
						ghttp.RespondWith(http.StatusOK, `{"service":{"srvid":"test-srv-id","service":"test-terraform","customer":"test-customer","serviceAddresses":["10.102.33.133"],"estadoProto":"http","estadoPort":8080,"estadoPath":"/service","estadoResultCode":0,"tcpports":[8080,8081],"udpports":null,"tcpportrange":"","udpportrange":"","region":"us-south","dataCenters":["dal13"],"maxSpeed":"1g","url":"test-terraform.test.cloud.ibm.com","hostname":"","dedicated":1,"multitenant":1,"acl":null,"creationTime":"2019-06-05T05:55:40Z","owner":"IBM","bss":"IBM"},"endpoints":[{"seid":"test-srv-id-dal13","srvid":"test-srv-id","mbid":"test-mbid","crn":"crn:v1:bluemix:public:serviceendpoint:dal13:a/xxxxxx:test-srv-id-dal13::","staticAddress":"166.9.1.144","netmask":"25","dnsStatus":"Y","region":"us-south","dataCenter":"dal13","vlanid":2288443,"status":"Ready","serverGroup":"groupA","serviceStatus":null,"statusDetails":[{"address":"10.102.33.133","ping":1,"estado":1,"ports":["8080:1","8081:1"]}],"heartbeatTime":"2019-06-14T01:32:39Z"}]}`),
					),
				)
			})
			It("should return the service info", func() {
				srvObj, err := newTestCseAPI(server.URL()).GetServiceEndpoint(srvID)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(srvObj.Service.ServiceName).Should(Equal("test-terraform"))
				Expect(srvObj.Service.CustomerName).Should(Equal("test-customer"))
			})
		})

		Context("When service with srvid being srvID does not exist", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, fmt.Sprintf("/v2/serviceendpoint/%s", srvID)),
						ghttp.RespondWith(http.StatusNotFound,
							fmt.Sprintf(`{{"message":"Not found service endpoint with id: %s"}}`, srvID)),
					),
				)
			})
			It("should return nil and an error", func() {
				srvObj, err := newTestCseAPI(server.URL()).GetServiceEndpoint(srvID)
				Expect(err).Should(HaveOccurred())
				Expect(srvObj).Should(BeNil())
			})
		})

		Context("When srvID is empty", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, fmt.Sprintf("/v2/serviceendpoint/%s", srvID)),
						ghttp.RespondWith(http.StatusNotFound,
							fmt.Sprintf(`{{"message":"Not found service endpoint with id: %s"}}`, srvID)),
					),
				)
			})
			It("should return nil and an error", func() {
				srvObj, err := newTestCseAPI(server.URL()).GetServiceEndpoint("")
				Expect(err).Should(HaveOccurred())
				Expect(srvObj).Should(BeNil())
			})
		})
	})

	Describe("CreateServiceEndpoint", func() {
		Context("When creation is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v2/serviceendpoint"),
						ghttp.RespondWith(http.StatusOK, `{"serviceid":"test-srv-id"}`),
					),
				)
			})
			It("should return the srvid of newly created service endpoint", func() {
				payload := SeCreateData{
					ServiceName:      "test-terrafor-11",
					CustomerName:     "test-customer-11",
					ServiceAddresses: []string{"10.102.33.131", "10.102.33.133"},
					Region:           "us-south",
					DataCenters:      []string{"dal10"},
					TCPPorts:         []int{8080, 80},
				}
				srvID, err := newTestCseAPI(server.URL()).CreateServiceEndpoint(payload)

				Expect(err).ShouldNot(HaveOccurred())
				Expect(srvID).Should(Equal("test-srv-id"))
			})
		})

		Context("When creation is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v2/serviceendpoint"),
						ghttp.RespondWith(http.StatusBadRequest, `{"message":"The port 0 in tcpports is invalid"}`),
					),
				)
			})
			It("should return empty srvid and an error", func() {
				payload := SeCreateData{
					ServiceName:      "test-terrafor-11",
					CustomerName:     "test-customer-11",
					ServiceAddresses: []string{"10.102.33.131", "10.102.33.133"},
					Region:           "us-south",
					DataCenters:      []string{"dal10"},
					TCPPorts:         []int{0, 8080, 80},
				}
				srvID, err := newTestCseAPI(server.URL()).CreateServiceEndpoint(payload)

				Expect(err).Should(HaveOccurred())
				Expect(srvID).Should(Equal(""))
			})
		})
	})

	Describe("UpdateServiceEndpoint", func() {
		Context("When update is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, fmt.Sprintf("/v2/serviceendpointtf/%s", srvID)),
						ghttp.RespondWith(http.StatusOK, `
							{"message": "Update the service endpoint and view its status to show update result."}
						`),
					),
				)
			})
			It("should return nil", func() {
				payload := SeUpdateData{
					DataCenters: []string{"dal10", "dal13"},
					TCPPorts:    []int{8080, 80, 8081},
				}
				err := newTestCseAPI(server.URL()).UpdateServiceEndpoint(srvID, payload)

				Expect(err).ShouldNot(HaveOccurred())
			})
		})
		Context("When update is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, fmt.Sprintf("/v2/serviceendpointtf/%s", srvID)),
						ghttp.RespondWith(http.StatusNotFound,
							fmt.Sprintf(`{"message": "Not found service endpoint with id:  %s"}`, srvID)),
					),
				)
			})
			It("should return an error", func() {
				payload := SeUpdateData{
					DataCenters: []string{"dal10", "dal13"},
					TCPPorts:    []int{8080, 80, 8081},
				}
				err := newTestCseAPI(server.URL()).UpdateServiceEndpoint(srvID, payload)
				Expect(err).Should(HaveOccurred())
			})
		})

		Context("When srvID is empty", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, fmt.Sprintf("/v2/serviceendpointtf/%s", srvID)),
						ghttp.RespondWith(http.StatusOK, `
							{"message": "Update the service endpoint and view its status to show update result."}
						`),
					),
				)
			})
			It("should return an error", func() {
				payload := SeUpdateData{
					DataCenters: []string{"dal10", "dal13"},
					TCPPorts:    []int{8080, 80, 8081},
				}
				err := newTestCseAPI(server.URL()).UpdateServiceEndpoint("", payload)

				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe("DeleteServiceEndpoint", func() {
		Context("When delete is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, fmt.Sprintf("/v2/serviceendpoint/%s", srvID)),
						ghttp.RespondWith(http.StatusOK,
							fmt.Sprintf(`{"message": "Success to delete service endpoint: %s"}`, srvID)),
					),
				)
			})
			It("should return nil", func() {
				err := newTestCseAPI(server.URL()).DeleteServiceEndpoint(srvID)
				Expect(err).ShouldNot(HaveOccurred())
			})
		})
		Context("When delete is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, fmt.Sprintf("/v2/serviceendpoint/%s", srvID)),
						ghttp.RespondWith(http.StatusNotFound,
							fmt.Sprintf(`{"message": "Not found service endpoint with id:  %s"}`, srvID)),
					),
				)
			})
			It("should return an error", func() {
				err := newTestCseAPI(server.URL()).DeleteServiceEndpoint(srvID)
				Expect(err).Should(HaveOccurred())
			})
		})

		Context("When srvID is empty", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, fmt.Sprintf("/v2/serviceendpoint/%s", srvID)),
						ghttp.RespondWith(http.StatusOK,
							fmt.Sprintf(`{"message": "Success to delete service endpoint: %s"}`, srvID)),
					),
				)
			})
			It("should return an error", func() {
				err := newTestCseAPI(server.URL()).DeleteServiceEndpoint("")
				Expect(err).Should(HaveOccurred())
			})
		})
	})

})

func newTestCseAPI(url string) ServiceEndpoints {
	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}
	conf := sess.Config.Copy()
	conf.HTTPClient = bluemixHttp.NewHTTPClient(conf)
	conf.Endpoint = &url

	client := client.Client{
		Config:      conf,
		ServiceName: bluemix.CseService,
	}
	return newServiceEndpointsAPI(&client)
}
