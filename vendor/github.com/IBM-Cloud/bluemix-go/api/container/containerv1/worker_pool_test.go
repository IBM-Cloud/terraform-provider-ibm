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

var _ = Describe("WorkerPool", func() {
	var server *ghttp.Server
	Describe("CreateWorkerPool", func() {
		Context("When creating a worker pool is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/clusters/test/workerpools"),
						ghttp.RespondWith(http.StatusCreated, `{"Name":"testpool","Size":5,"MachineType": "u2c.2x4","Isolation": "public","ID":"rtr4tg5", "Region":"us-south", "State":"normal", "ReasonForDelete":"","IsBalanced":true}`),
					),
				)
			})

			It("should return worker pools added to cluster", func() {
				workerPoolProperties := WorkerPoolRequest{
					WorkerPoolConfig: WorkerPoolConfig{
						Name:        "test-pool",
						Size:        5,
						MachineType: "u2c.2x4",
						Isolation:   "public",
					},
					DiskEncryption: true,
				}
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}

				_, err := newWorkerPool(server.URL()).CreateWorkerPool("test", workerPoolProperties, target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When creating worker pool is unsuccessful", func() {
			BeforeEach(func() {

				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/clusters/test/workerpools"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to create worker pools to cluster`),
					),
				)
			})

			It("should return worker pools added to cluster", func() {
				workerPoolProperties := WorkerPoolRequest{
					WorkerPoolConfig: WorkerPoolConfig{
						Name:        "test-pool",
						Size:        5,
						MachineType: "u2c.2x4",
						Isolation:   "public",
					},
					DiskEncryption: true,
				}
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}
				_, err := newWorkerPool(server.URL()).CreateWorkerPool("test", workerPoolProperties, target)
				Expect(err).To(HaveOccurred())
			})
		})
	})
	//List
	Describe("List", func() {
		Context("When retrieving available worker pools of a cluster is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/clusters/myCluster/workerpools"),
						ghttp.RespondWith(http.StatusOK, `[{"Name":"testpool","Size":5,"MachineType": "u2c.2x4","Isolation": "public","ID":"rtr4tg5", "Region":"us-south", "State":"normal","ReasonForDelete":"","IsBalanced":true}]`),
					),
				)
			})

			It("should return available worker pools ", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}

				worker, err := newWorkerPool(server.URL()).ListWorkerPools("myCluster", target)
				Expect(err).NotTo(HaveOccurred())
				Expect(worker).ShouldNot(BeNil())
				for _, wObj := range worker {
					Expect(wObj).ShouldNot(BeNil())
					Expect(wObj.State).Should(Equal("normal"))
				}
			})
		})
		Context("When retrieving available worker pools is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/clusters/myCluster/workerpools"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve workerpools`),
					),
				)
			})

			It("should return error during retrieveing worker pools", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}

				_, err := newWorkerPool(server.URL()).ListWorkerPools("myCluster", target)
				Expect(err).To(HaveOccurred())
			})
		})
	})
	//Get
	Describe("Get", func() {
		Context("When retrieving worker pool of a cluster is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/clusters/myCluster/workerpools/abc-123-def"),
						ghttp.RespondWith(http.StatusOK, `{"Name":"testpool","Size":5,"MachineType": "u2c.2x4","Isolation": "public","ID":"rtr4tg5", "Region":"us-south", "State":"normal", "ReasonForDelete":"","IsBalanced":true}`),
					),
				)
			})

			It("should return worker pool", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}

				_, err := newWorkerPool(server.URL()).GetWorkerPool("myCluster", "abc-123-def", target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When retrieving worker pool is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/clusters/myCluster/workerpools/abc-123-def"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve workerpool`),
					),
				)
			})

			It("should return error during retrieveing worker pool", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}

				_, err := newWorkerPool(server.URL()).GetWorkerPool("myCluster", "abc-123-def", target)
				Expect(err).To(HaveOccurred())
			})
		})
	})
	//Delete
	Describe("Delete", func() {
		Context("When delete of worker pool is successful", func() {
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

			It("should delete worker pool", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}

				err := newWorkerPool(server.URL()).DeleteWorkerPool("test", "abc-123-def-ghi", target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When worker pool delete is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/clusters/test/workerpools/abc-123-def-ghi"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to delete worker pool`),
					),
				)
			})

			It("should return error deleting worker pool", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}

				err := newWorkerPool(server.URL()).DeleteWorkerPool("test", "abc-123-def-ghi", target)
				Expect(err).To(HaveOccurred())
			})
		})
	})
	//Patch worker pool
	Describe("Patch", func() {
		Context("When resize of worker pool is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPatch, "/v1/clusters/test/workerpools/abc-123-def-ghi"),
						ghttp.RespondWith(http.StatusOK, `{							
						}`),
					),
				)
			})

			It("should increase the size of worker pool", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}

				err := newWorkerPool(server.URL()).ResizeWorkerPool("test", "abc-123-def-ghi", 6, target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When resize of worker pool is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPatch, "/v1/clusters/test/workerpools/abc-123-def-ghi"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to resize worker pool`),
					),
				)
			})

			It("should return error resizing worker pool", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}

				err := newWorkerPool(server.URL()).ResizeWorkerPool("test", "abc-123-def-ghi", 6, target)
				Expect(err).To(HaveOccurred())
			})
		})
	})
	//Add zone
	Describe("Post", func() {
		Context("When adding a zone is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/clusters/test/workerpools/abc-123-def-ghi/zones"),
						ghttp.RespondWith(http.StatusOK, `{							
						}`),
					),
				)
			})

			It("should add zone to the specified worker pool", func() {
				workerPoolZone := WorkerPoolZone{
					ID: "abc-123-def-ghi",
					WorkerPoolZoneNetwork: WorkerPoolZoneNetwork{
						PrivateVLAN: "12345",
						PublicVLAN:  "43215",
					},
				}
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}

				err := newWorkerPool(server.URL()).AddZone("test", "abc-123-def-ghi", workerPoolZone, target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When adding zone to worker pool is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/clusters/test/workerpools/abc-123-def-ghi/zones"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to add zone to worker pool`),
					),
				)
			})

			It("should return error adding zone to worker pool", func() {

				workerPoolZone := WorkerPoolZone{
					ID: "abc-123-def-ghi",
					WorkerPoolZoneNetwork: WorkerPoolZoneNetwork{
						PrivateVLAN: "12345",
						PublicVLAN:  "43215",
					},
				}
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}

				err := newWorkerPool(server.URL()).AddZone("test", "abc-123-def-ghi", workerPoolZone, target)
				Expect(err).To(HaveOccurred())
			})
		})
	})
	//Remove zone
	Describe("Delete", func() {
		Context("When delete of zone of a worker pool is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/clusters/test/workerpools/abc-123-def-ghi/zones/dal10"),
						ghttp.RespondWith(http.StatusOK, `{							
						}`),
					),
				)
			})

			It("should delete worker pool in that zone", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}

				err := newWorkerPool(server.URL()).RemoveZone("test", "dal10", "abc-123-def-ghi", target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When delete of zone of a worker pool delete is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/clusters/test/workerpools/abc-123-def-ghi/zones/dal10"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to delete zone of a worker pool`),
					),
				)
			})

			It("should return error deleting worker pool in the specific zone", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}

				err := newWorkerPool(server.URL()).RemoveZone("test", "dal10", "abc-123-def-ghi", target)
				Expect(err).To(HaveOccurred())
			})
		})
	})
	//UpdateZoneNetwork
	Describe("Patch", func() {
		Context("When update of worker pool zone is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPatch, "/v1/clusters/test/workerpools/abc-123-def-ghi/zones/dal10"),
						ghttp.RespondWith(http.StatusOK, `{							
						}`),
					),
				)
			})

			It("should update worker pool zone", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}

				err := newWorkerPool(server.URL()).UpdateZoneNetwork("test", "dal10", "abc-123-def-ghi", "12345", "43215", target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When update of worker pool zone is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPatch, "/v1/clusters/test/workerpools/abc-123-def-ghi/zones/dal10"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to update worker pool zone`),
					),
				)
			})

			It("should return updating worker pool zone", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}

				err := newWorkerPool(server.URL()).UpdateZoneNetwork("test", "dal10", "abc-123-def-ghi", "12345", "43215", target)
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
		ServiceName: bluemix.MccpService,
	}
	return newWorkerPoolAPI(&client)
}
