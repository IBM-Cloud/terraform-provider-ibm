package cisv1

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

var _ = Describe("Dns", func() {
    var server *ghttp.Server
    AfterEach(func() {
        server.Close()
    })
    
    //List
    Describe("List", func() {
        Context("When read of Dns is successful", func() {
            BeforeEach(func() {
                server = ghttp.NewServer()
                server.AppendHandlers(
                    ghttp.CombineHandlers(
                        ghttp.VerifyRequest(http.MethodGet, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/zones/3fefc35e7decadb111dcf85d723a4f20/dns_records"),
                        ghttp.RespondWith(http.StatusOK, `
                           {
                              "result": [
                                {
                                  "id": "0f4740fc36065f8a9343c7ed9445f2a4",
                                  "type": "A",
                                  "name": "example.com",
                                  "content": "192.168.127.127",
                                  "proxiable": true,
                                  "proxied": false,
                                  "ttl": 3600,
                                  "locked": false,
                                  "zone_id": "3fefc35e7decadb111dcf85d723a4f20",
                                  "zone_name": "example.com",
                                  "modified_on": "2018-10-11T13:34:45.189800Z",
                                  "created_on": "2018-10-11T13:34:45.189800Z"
                                },
                                {
                                  "id": "0f4740fc36065f8a9343c7ed9445f2a4",
                                  "type": "A",
                                  "name": "www.example.com",
                                  "content": "192.168.127.127",
                                  "proxiable": false,
                                  "proxied": false,
                                  "ttl": 1,
                                  "locked": false,
                                  "zone_id": "3fefc35e7decadb111dcf85d723a4f20",
                                  "zone_name": "example.com",
                                  "modified_on": "2018-10-12T06:04:36.533540Z",
                                  "created_on": "2018-10-12T06:04:36.533540Z"
                                }
                              ],
                              "success": true,
                              "errors": [],
                              "messages": []
                            }
                        `),
                    ),
                )
            })

            It("should return Dns list", func() {
                target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
                target2 := "3fefc35e7decadb111dcf85d723a4f20"
                myDnsPtr, err := newDns(server.URL()).ListDns(target1, target2)
                myDns := *myDnsPtr
                Expect(myDns).ShouldNot(BeNil())
                for _, Dns := range myDns {
                    Expect(err).NotTo(HaveOccurred())
                    Expect(Dns.Id).Should(Equal("0f4740fc36065f8a9343c7ed9445f2a4"))
                    Expect(Dns.Content).Should(Equal("192.168.127.127"))
                }
            })
        })
        Context("When read of Dns is unsuccessful", func() {
            BeforeEach(func() {
                server = ghttp.NewServer()
                server.AppendHandlers(
                    ghttp.CombineHandlers(
                        ghttp.VerifyRequest(http.MethodGet, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/zones/3fefc35e7decadb111dcf85d723a4f20/dns_records"),
                        ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve Dns`),
                    ),
                )
            })

            It("should return error when Dns are retrieved", func() {
                target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
                target2 := "3fefc35e7decadb111dcf85d723a4f20"
                myDnsPtr, err := newDns(server.URL()).ListDns(target1, target2)
                myDns := myDnsPtr
                Expect(err).To(HaveOccurred())
                Expect(myDns).Should(BeNil())
            })
        })
    })
    
    //Delete
    Describe("Delete", func() {
        Context("When delete of Dns is successful", func() {
            BeforeEach(func() {
                server = ghttp.NewServer()
                server.AppendHandlers(
                    ghttp.CombineHandlers(
                        ghttp.VerifyRequest(http.MethodDelete, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/zones/3fefc35e7decadb111dcf85d723a4f20/dns_records/58e96913cff39f73d8901a6a4ea07e16"),
                        ghttp.RespondWith(http.StatusOK, `{                         
                        }`),
                    ),
                )
            })

            It("should delete Dns", func() {
                target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
                target2 := "3fefc35e7decadb111dcf85d723a4f20"
                params := "58e96913cff39f73d8901a6a4ea07e16"
                err := newDns(server.URL()).DeleteDns(target1, target2, params)
                Expect(err).NotTo(HaveOccurred())
            })
        })
        Context("When Dns delete has failed", func() {
            BeforeEach(func() {
                server = ghttp.NewServer()
                server.AppendHandlers(
                    ghttp.CombineHandlers(
                        ghttp.VerifyRequest(http.MethodDelete, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/zones/3fefc35e7decadb111dcf85d723a4f20/dns_records/58e96913cff39f73d8901a6a4ea07e16"),
                        ghttp.RespondWith(http.StatusInternalServerError, `Failed to delete service key`),
                    ),
                )
            })

            It("should return error zone delete", func() {
                target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
                target2 := "3fefc35e7decadb111dcf85d723a4f20"
                params := "58e96913cff39f73d8901a6a4ea07e16"
                err := newDns(server.URL()).DeleteDns(target1, target2, params)
                Expect(err).To(HaveOccurred())
            })
        })
    })
    //Find
    Describe("Get", func() {
        Context("When read of Dns is successful", func() {
            BeforeEach(func() {
                server = ghttp.NewServer()
                server.AppendHandlers(
                    ghttp.CombineHandlers(
                        ghttp.VerifyRequest(http.MethodGet, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/zones/3fefc35e7decadb111dcf85d723a4f20/dns_records/0f4740fc36065f8a9343c7ed9445f2a4"),
                        ghttp.RespondWith(http.StatusOK, `
                                {
                                  "result": {
                                    "id": "0f4740fc36065f8a9343c7ed9445f2a4",
                                    "type": "A",
                                    "name": "www.example.com",
                                    "content": "192.168.127.127",
                                    "proxiable": false,
                                    "proxied": false,
                                    "ttl": 1,
                                    "locked": false,
                                    "zone_id": "3fefc35e7decadb111dcf85d723a4f20",
                                    "zone_name": "example.com",
                                    "modified_on": "2018-10-12T06:04:36.533540Z",
                                    "created_on": "2018-10-12T06:04:36.533540Z"
                                },
                                  "success": true,
                                  "errors": [],
                                  "messages": []
                                }
                        `),
                    ),
                )
            })

            It("should return Dns", func() {
                target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
                target2 := "3fefc35e7decadb111dcf85d723a4f20"
                params := "0f4740fc36065f8a9343c7ed9445f2a4"
                myDnsPtr, err := newDns(server.URL()).GetDns(target1, target2, params)
                myDns := *myDnsPtr
                Expect(err).NotTo(HaveOccurred())
                Expect(myDns).ShouldNot(BeNil())
                Expect(myDns.Id).Should(Equal("0f4740fc36065f8a9343c7ed9445f2a4"))
            })
        })
        Context("When Dns get has failed", func() {
            BeforeEach(func() {
                server = ghttp.NewServer()
                server.AppendHandlers(
                    ghttp.CombineHandlers(
                        ghttp.VerifyRequest(http.MethodGet, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/zones/3fefc35e7decadb111dcf85d723a4f20/dns_records/0f4740fc36065f8a9343c7ed9445f2a4"),
                        ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve Dns`),
                    ),
                )
            })

            It("should return error when Dns is retrieved", func() {
                target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
                target2 := "3fefc35e7decadb111dcf85d723a4f20"
                params := "0f4740fc36065f8a9343c7ed9445f2a4"
                myDnsPtr, err := newDns(server.URL()).GetDns(target1, target2, params)
                myDns := myDnsPtr
                Expect(err).To(HaveOccurred())
                Expect(myDns).Should(BeNil())
            })
        })
    })
    Describe("Create", func() {
        Context("When creation is successful", func() {
            BeforeEach(func() {
                server = ghttp.NewServer()
                server.AppendHandlers(
                    ghttp.CombineHandlers(
                        ghttp.VerifyRequest(http.MethodPost, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/zones/3fefc35e7decadb111dcf85d723a4f20/dns_records"),
                        ghttp.RespondWith(http.StatusCreated, `
                           {
                              "result": {
                                "id": "0f4740fc36065f8a9343c7ed9445f2a4",
                                "type": "A",
                                "name": "www.example.com",
                                "content": "192.168.127.127",
                                "proxiable": false,
                                "proxied": false,
                                "ttl": 1,
                                "locked": false,
                                "zone_id": "3fefc35e7decadb111dcf85d723a4f20",
                                "zone_name": "example.com",
                                "modified_on": "2018-10-12T06:04:36.533540Z",
                                "created_on": "2018-10-12T06:04:36.533540Z"
                              },
                              "success": true,
                              "errors": [],
                              "messages": []
                            }
                        `),
                    ),
                )
            })

            It("should return dns created", func() {
                params := DnsBody{
                        Name: "www.example.com",
                        DnsType: "A",
                        Content: "192.168.127.127",
                    }
                target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
                target2 := "3fefc35e7decadb111dcf85d723a4f20"
                myDnsPt, err := newDns(server.URL()).CreateDns(target1, target2, params)
                myDns := *myDnsPt    
                Expect(err).NotTo(HaveOccurred())
                Expect(myDns).ShouldNot(BeNil())
                Expect(myDns.Id).Should(Equal("0f4740fc36065f8a9343c7ed9445f2a4"))
                Expect(myDns.Name).Should(Equal("www.example.com"))
            })
        })
        Context("When creation is unsuccessful", func() {
            BeforeEach(func() {
                server = ghttp.NewServer()
                server.AppendHandlers(
                    ghttp.CombineHandlers(
                        ghttp.VerifyRequest(http.MethodPost, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/zones/0f4740fc36065f8a9343c7ed9445f2a4/dns_records"),
                        ghttp.RespondWith(http.StatusInternalServerError, `Failed to create Dns`),
                    ),
                )
            })
            It("should return error during Dns creation", func() {
                params := DnsBody{
                        Name: "www.example.com",
                        DnsType: "A",
                        Content: "192.168.127.127",
                    }
                target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
                target2 := "0f4740fc36065f8a9343c7ed9445f2a4"
                myDnsPtr, err := newDns(server.URL()).CreateDns(target1, target2, params)
                myDns := myDnsPtr
                Expect(err).To(HaveOccurred())
                Expect(myDns).Should(BeNil())
            })
        })
    })
})

func newDns(url string) Dns {

    sess, err := session.New()
    if err != nil {
        log.Fatal(err)
    }
    conf := sess.Config.Copy()
    conf.HTTPClient = bluemixHttp.NewHTTPClient(conf)
    conf.Endpoint = &url

    client := client.Client{
        Config:      conf,
        ServiceName: bluemix.CisService,
    }
    return newDnsAPI(&client)
}
