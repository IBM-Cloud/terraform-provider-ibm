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

var _ = Describe("Monitors", func() {
    var server *ghttp.Server
    AfterEach(func() {
        server.Close()
    })
    Describe("Create", func() {
        Context("When creation is successful", func() {
            BeforeEach(func() {
                server = ghttp.NewServer()
                server.AppendHandlers(
                    ghttp.CombineHandlers(
                        ghttp.VerifyRequest(http.MethodPost, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/load_balancers/monitors"),
                        ghttp.RespondWith(http.StatusCreated, `
                           {
                              "result": {
                                "description": "",
                                "created_on": "2018-11-22T09:53:06.416054Z",
                                "modified_on": "2018-11-22T09:53:06.416054Z",
                                "id": "92859a0f6b4d3e55b953e0e29bb96338",
                                "type": "http",
                                "interval": 60,
                                "retries": 2,
                                "timeout": 5,
                                "expected_body": "",
                                "expected_codes": "200",
                                "follow_redirects": true,
                                "allow_insecure": false,
                                "path": "/status",
                                "header": {
                                  "Host": [
                                    "www.example.com"
                                  ],
                                  "X-App-Id": [
                                    "abc123"
                                  ]
                                },
                                "method": "GET"
                              },
                              "success": true,
                              "errors": [],
                              "messages": []
                            }
                        `),
                    ),
                )
            })

            It("should return monitor created", func() {
                params := MonitorBody{
                        ExpCodes: "200",
                        ExpBody: "",
                        Path:  "/status",   
                        MonType:  "http",
                        Method: "GET",
                        Timeout:  5,
                        Retries: 2,
                        Interval: 60, 
                        FollowRedirects: true,  
                        AllowInsecure: false,
                    }
                target := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
                myMonitorPtr, err := newMonitor(server.URL()).CreateMonitor(target, params)
                myMonitor := *myMonitorPtr
                Expect(err).NotTo(HaveOccurred())
                Expect(myMonitor).ShouldNot(BeNil())
                Expect(myMonitor.Id).Should(Equal("92859a0f6b4d3e55b953e0e29bb96338"))
            })
        })
        Context("When creation is unsuccessful", func() {
            BeforeEach(func() {
                server = ghttp.NewServer()
                server.AppendHandlers(
                    ghttp.CombineHandlers(
                        ghttp.VerifyRequest(http.MethodPost, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/load_balancers/monitors"),
                        ghttp.RespondWith(http.StatusInternalServerError, `Failed to create Monitor`),
                    ),
                )
            })
            It("should return error during Monitor creation", func() {
                params := MonitorBody{
                        ExpCodes: "200",
                        ExpBody: "",
                        Path:  "/status",   
                        MonType:  "http",
                        Method: "GET",
                        Timeout:  5,
                        Retries: 2,
                        Interval: 60, 
                        FollowRedirects: true,  
                        AllowInsecure: false,
                    }
                target := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
                myMonitorPtr, err := newMonitor(server.URL()).CreateMonitor(target, params)
                myMonitor := myMonitorPtr
                Expect(err).To(HaveOccurred())
                Expect(myMonitor).Should(BeNil())
            })
        })
    })
    //List
    Describe("List", func() {
        Context("When read of Monitors is successful", func() {
            BeforeEach(func() {
                server = ghttp.NewServer()
                server.AppendHandlers(
                    ghttp.CombineHandlers(
                        ghttp.VerifyRequest(http.MethodGet, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/load_balancers/monitors"),
                        ghttp.RespondWith(http.StatusOK, `
                            {
                              "result": [
                                {
                                  "description": "",
                                  "created_on": "2018-06-02T03:14:09.818402Z",
                                  "modified_on": "2018-11-22T08:54:25.126766Z",
                                  "id": "192c950172152639e21f549bc4a1cd6f",
                                  "type": "http",
                                  "interval": 60,
                                  "retries": 2,
                                  "timeout": 5,
                                  "expected_body": "",
                                  "expected_codes": "200",
                                  "follow_redirects": true,
                                  "allow_insecure": false,
                                  "path": "/status",
                                  "header": {},
                                  "method": "GET"
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

            It("should return Monitor list", func() {
                target := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
                myMonitorsPtr, err := newMonitor(server.URL()).ListMonitors(target)
                myMonitors := *myMonitorsPtr
                Expect(myMonitors).ShouldNot(BeNil())
                for _, Monitor := range myMonitors {
                    Expect(err).NotTo(HaveOccurred())
                    Expect(Monitor.Id).Should(Equal("192c950172152639e21f549bc4a1cd6f"))
                }
            })
        })
        Context("When read of Monitors is unsuccessful", func() {
            BeforeEach(func() {
                server = ghttp.NewServer()
                server.AppendHandlers(
                    ghttp.CombineHandlers(
                        ghttp.VerifyRequest(http.MethodGet, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/load_balancers/monitors"),
                        ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve Monitors`),
                    ),
                )
            })

            It("should return error when Monitor are retrieved", func() {
                target := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a" 
                myMonitorPtr, err := newMonitor(server.URL()).ListMonitors(target)
                myMonitor := myMonitorPtr
                Expect(err).To(HaveOccurred())
                Expect(myMonitor).Should(BeNil())
            })
        })
    })
    //Delete
    Describe("Delete", func() {
        Context("When delete of Monitor is successful", func() {
            BeforeEach(func() {
                server = ghttp.NewServer()
                server.AppendHandlers(
                    ghttp.CombineHandlers(
                        ghttp.VerifyRequest(http.MethodDelete, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/load_balancers/monitors/92859a0f6b4d3e55b953e0e29bb96338"),
                        ghttp.RespondWith(http.StatusOK, `{                         
                        }`),
                    ),
                )
            })

            It("should delete Monitor", func() {
                target := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
                params := "92859a0f6b4d3e55b953e0e29bb96338"
                err := newMonitor(server.URL()).DeleteMonitor(target, params)
                Expect(err).NotTo(HaveOccurred())
            })
        })
        Context("When Monitor delete has failed", func() {
            BeforeEach(func() {
                server = ghttp.NewServer()
                server.AppendHandlers(
                    ghttp.CombineHandlers(
                        ghttp.VerifyRequest(http.MethodDelete, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/load_balancers/monitors/92859a0f6b4d3e55b953e0e29bb96338"),
                        ghttp.RespondWith(http.StatusInternalServerError, `Failed to delete service key`),
                    ),
                )
            })

            It("should return error zone delete", func() {
                target := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
                params := "92859a0f6b4d3e55b953e0e29bb96338"
                err := newMonitor(server.URL()).DeleteMonitor(target, params)
                Expect(err).To(HaveOccurred())
            })
        })
    })
    //Find
    Describe("Get", func() {
        Context("When read of Monitor is successful", func() {
            BeforeEach(func() {
                server = ghttp.NewServer()
                server.AppendHandlers(
                    ghttp.CombineHandlers(
                        ghttp.VerifyRequest(http.MethodGet, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/load_balancers/monitors/92859a0f6b4d3e55b953e0e29bb96338"),
                        ghttp.RespondWith(http.StatusOK, `
                            {
                              "result": {
                                "description": "",
                                "created_on": "2018-06-02T03:14:09.818402Z",
                                "modified_on": "2018-11-22T08:54:25.126766Z",
                                "id": "192c950172152639e21f549bc4a1cd6f",
                                "type": "http",
                                "interval": 60,
                                "retries": 2,
                                "timeout": 5,
                                "expected_body": "",
                                "expected_codes": "200",
                                "follow_redirects": true,
                                "allow_insecure": false,
                                "path": "/status",
                                "header": {},
                                "method": "GET"
                              },
                              "success": true,
                              "errors": [],
                              "messages": []
                            }
                        `),
                    ),
                )
            })

            It("should return Monitor", func() {
               target := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
                params := "92859a0f6b4d3e55b953e0e29bb96338"
                myMonitorPtr, err := newMonitor(server.URL()).GetMonitor(target, params)
                myMonitor := myMonitorPtr
                Expect(err).NotTo(HaveOccurred())
                Expect(myMonitor).ShouldNot(BeNil())
                Expect(myMonitor.Id).Should(Equal("192c950172152639e21f549bc4a1cd6f"))
            })
        })
        Context("When Monitor get has failed", func() {
            BeforeEach(func() {
                server = ghttp.NewServer()
                server.AppendHandlers(
                    ghttp.CombineHandlers(
                        ghttp.VerifyRequest(http.MethodGet, "/v1/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/load_balancers/monitors/92859a0f6b4d3e55b953e0e29bb96338"),
                        ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve Monitor`),
                    ),
                )
            })

            It("should return error when Monitor is retrieved", func() {
                target := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
                params := "92859a0f6b4d3e55b953e0e29bb96338"
                myMonitorPtr, err := newMonitor(server.URL()).GetMonitor(target, params)
                myMonitor := myMonitorPtr
                Expect(err).To(HaveOccurred())
                Expect(myMonitor).Should(BeNil())
            })
        })
    })
 
})

func newMonitor(url string) Monitors {

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
    return newMonitorAPI(&client)
}
