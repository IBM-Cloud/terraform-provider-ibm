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

var _ = Describe("workerpools", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})

	//Create
	Describe("Create", func() {
		Context("When creating workerpool is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v2/vpc/createWorkerPool"),
						ghttp.VerifyJSON(`{"cluster":"bm64u3ed02o93vv36hb0","flavor":"b2.4x16","labels":{},"name":"mywork211","vpcID":"6015365a-9d93-4bb4-8248-79ae0db2dc26","workerCount":1,"zones":[]}`),
						ghttp.RespondWith(http.StatusCreated, `{
							"workerPoolID":"string"
						}`),
					),
				)
			})

			It("should create Workerpool in a cluster", func() {
				target := ClusterTargetHeader{}
				params := WorkerPoolRequest{
					Cluster: "bm64u3ed02o93vv36hb0",
					WorkerPoolConfig: WorkerPoolConfig{
						Flavor:      "b2.4x16",
						Name:        "mywork211",
						VpcID:       "6015365a-9d93-4bb4-8248-79ae0db2dc26",
						WorkerCount: 1,
						Zones:       []Zone{},
					},
				}
				_, err := newWorkerPool(server.URL()).CreateWorkerPool(params, target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When creating workerpool is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v2/vpc/createWorkerPool"),
						ghttp.VerifyJSON(`{"cluster":"bm64u3ed02o93vv36hb0","flavor":"b2.4x16","labels":{},"name":"mywork211","vpcID":"6015365a-9d93-4bb4-8248-79ae0db2dc26","workerCount":1,"zones":[]}
`),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to create workerpool`),
					),
				)
			})

			It("should return error during creating workerpool", func() {
				params := WorkerPoolRequest{
					Cluster: "bm64u3ed02o93vv36hb0",
					WorkerPoolConfig: WorkerPoolConfig{
						Flavor:      "b2.4x16",
						Name:        "mywork211",
						VpcID:       "6015365a-9d93-4bb4-8248-79ae0db2dc26",
						WorkerCount: 1,
						Zones:       []Zone{},
					},
				}
				target := ClusterTargetHeader{}
				_, err := newWorkerPool(server.URL()).CreateWorkerPool(params, target)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	//getworkerpools
	Describe("Get", func() {
		Context("When Get workerpool is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/vpc/getWorkerPool"),
						ghttp.RespondWith(http.StatusCreated, `{
							"flavor": "string",
							"id": "string",
							"isolation": "string",
							"labels": {
							},
							"lifecycle": {
							  "actualState": "string",
							  "desiredState": "string"
							},
							"poolName": "string",
							"provider": "string",
							"vpcID": "string",
							"workerCount": 0,
							"zones": [
							  {
								"id": "string",
								"subnets": [
								  {
									"id": "string",
									"primary": true
								  }
								],
								"workerCount": 0
							  }
							]
						  }`),
					),
				)
			})

			It("should get Workerpool in a cluster", func() {
				target := ClusterTargetHeader{}

				_, err := newWorkerPool(server.URL()).GetWorkerPool("aaa", "bbb", target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When get workerpool is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/vpc/getWorkerPool"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to get workerpool`),
					),
				)
			})

			It("should return error during get workerpool", func() {
				target := ClusterTargetHeader{}
				_, err := newWorkerPool(server.URL()).GetWorkerPool("aaa", "bbb", target)
				Expect(err).To(HaveOccurred())
			})
		})
	})
	//Delete
	Describe("Delete", func() {
		Context("When delete of worker is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/clusters/test/workerpools/abc-123-def-ghi"),
						ghttp.RespondWith(http.StatusOK, `{							
						}`),
					),
				)
			})

			It("should delete workerpool", func() {
				target := ClusterTargetHeader{}
				err := newWorkerPool(server.URL()).DeleteWorkerPool("test", "abc-123-def-ghi", target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When cluster delete is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/clusters/test/workerpools/abc-123-def-ghi"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to delete worker`),
					),
				)
			})

			It("should return error service key delete", func() {
				target := ClusterTargetHeader{}
				err := newWorkerPool(server.URL()).DeleteWorkerPool("test", "abc-123-def-ghi", target)
				Expect(err).To(HaveOccurred())
			})
		})
	})
})

func newWorkerPool(url string) WorkerPool {

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
	return newWorkerPoolAPI(&client)
}
