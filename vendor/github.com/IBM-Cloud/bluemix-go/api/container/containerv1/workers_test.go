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

var _ = Describe("Workers", func() {
	var server *ghttp.Server
	Describe("Add", func() {
		Context("When adding a worker is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/clusters/test/workers"),
						ghttp.RespondWith(http.StatusCreated, `{}`),
					),
				)
			})

			It("should return worker added to cluster", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				params := WorkerParam{
					Isolation: "public", MachineType: "u2c.2x4", Prefix: "test", PrivateVlan: "1764491", PublicVlan: "1764435", WorkerNum: 2,
				}
				err := newWorker(server.URL()).Add("test", params, target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When adding worker is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/clusters/test/workers"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to add worker to cluster`),
					),
				)
			})

			It("should return error during add webhook to cluster", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				params := WorkerParam{
					Isolation: "public", MachineType: "u2c.2x4", Prefix: "test", PrivateVlan: "1764491", PublicVlan: "1764435", WorkerNum: 2,
				}
				err := newWorker(server.URL()).Add("test", params, target)
				Expect(err).To(HaveOccurred())
			})
		})
	})
	//Get
	Describe("Get", func() {
		Context("When retrieving worker is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/workers/abc-123-def-ghi"),
						ghttp.RespondWith(http.StatusOK, `{"ErrorMessage":"","Isolation":"","MachineType":"free","KubeVersion":"","PrivateIP":"","PublicIP":"","PrivateVlan":"vlan","PublicVlan":"vlan","state":"normal","status":"ready"}`),
					),
				)
			})

			It("should return available workers ", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				worker, err := newWorker(server.URL()).Get("abc-123-def-ghi", target)
				Expect(err).NotTo(HaveOccurred())
				Expect(worker).ShouldNot(BeNil())
				Expect(worker.State).Should(Equal("normal"))
			})
		})
		Context("When retrieving worker is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/workers/abc-123-def-ghi"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve workers`),
					),
				)
			})

			It("should return error during retrieveing workers", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				worker, err := newWorker(server.URL()).Get("abc-123-def-ghi", target)
				Expect(err).To(HaveOccurred())
				Expect(worker.ID).Should(Equal(""))
				Expect(worker.State).Should(Equal(""))
			})
		})
	})
	//List
	Describe("List", func() {
		Context("When retrieving available workers of a cluster is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/clusters/myCluster/workers"),
						ghttp.RespondWith(http.StatusOK, `[{"ErrorMessage":"","Isolation":"","MachineType":"free","KubeVersion":"","PrivateIP":"","PublicIP":"","PrivateVlan":"vlan","PublicVlan":"vlan","state":"normal","status":"ready"}]`),
					),
				)
			})

			It("should return available workers ", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				worker, err := newWorker(server.URL()).List("myCluster", target)
				Expect(err).NotTo(HaveOccurred())
				Expect(worker).ShouldNot(BeNil())
				for _, wObj := range worker {
					Expect(wObj).ShouldNot(BeNil())
					Expect(wObj.State).Should(Equal("normal"))
				}
			})
		})
		Context("When retrieving available workers is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/clusters/myCluster/workers"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve workers`),
					),
				)
			})

			It("should return error during retrieveing workers", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				worker, err := newWorker(server.URL()).List("myCluster", target)
				Expect(err).To(HaveOccurred())
				Expect(worker).Should(BeNil())
				Expect(len(worker)).Should(Equal(0))
			})
		})
	})
	//ListByWorkerPool
	Describe("ListByWorkerPool", func() {
		Context("When retrieving available workers belong to a worker pool of a cluster is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/clusters/myCluster/workers"),
						ghttp.RespondWith(http.StatusOK, `[{"ErrorMessage":"","Isolation":"","MachineType":"free","KubeVersion":"","PrivateIP":"","PublicIP":"","PrivateVlan":"vlan","PublicVlan":"vlan","state":"normal","status":"ready"}]`),
					),
				)
			})

			It("should return available workers ", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}

				worker, err := newWorker(server.URL()).ListByWorkerPool("myCluster", "test", false, target)
				Expect(err).NotTo(HaveOccurred())
				Expect(worker).ShouldNot(BeNil())
				for _, wObj := range worker {
					Expect(wObj).ShouldNot(BeNil())
					Expect(wObj.State).Should(Equal("normal"))
				}
			})
		})
		Context("When retrieving available workers is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/clusters/myCluster/workers"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve workers`),
					),
				)
			})

			It("should return error during retrieveing workers", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}

				worker, err := newWorker(server.URL()).ListByWorkerPool("myCluster", "test", false, target)
				Expect(err).To(HaveOccurred())
				Expect(worker).Should(BeNil())
				Expect(len(worker)).Should(Equal(0))
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
						ghttp.VerifyRequest(http.MethodDelete, "/v1/clusters/test/workers/abc-123-def-ghi"),
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
				err := newWorker(server.URL()).Delete("test", "abc-123-def-ghi", target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When cluster delete is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/clusters/test/workers/abc-123-def-ghi"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to delete worker`),
					),
				)
			})

			It("should return error service key delete", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				err := newWorker(server.URL()).Delete("test", "abc-123-def-ghi", target)
				Expect(err).To(HaveOccurred())
			})
		})
	})
	//Update
	Describe("Update", func() {
		Context("When update worker is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, "/v1/clusters/test/workers/abc-123-def-ghi"),
						ghttp.RespondWith(http.StatusCreated, `{}`),
					),
				)
			})

			It("should return worker updated", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				params := WorkerUpdateParam{
					Action: "reload",
				}
				err := newWorker(server.URL()).Update("test", "abc-123-def-ghi", params, target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When updating worker is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, "/v1/clusters/test/workers/abc-123-def-ghi"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to add worker to cluster`),
					),
				)
			})

			It("should return error during updating worker", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				params := WorkerUpdateParam{
					Action: "reload",
				}
				err := newWorker(server.URL()).Update("test", "abc-123-def-ghi", params, target)
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
		ServiceName: bluemix.MccpService,
	}
	return newWorkerAPI(&client)
}
