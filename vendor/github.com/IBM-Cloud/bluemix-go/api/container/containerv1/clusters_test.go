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

var _ = Describe("Clusters", func() {
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
						ghttp.VerifyRequest(http.MethodPost, "/v1/clusters"),
						ghttp.VerifyJSON(`{"dataCenter":"dal10","isolation":"","machineType":"b2c.4x16","name":"testservice","privateVlan":"vlan","publicVlan":"vlan","workerNum":1,"noSubnet":false,"masterVersion":"1.8.1","prefix":"worker","diskEncryption": true,"enableTrusted":true,"privateSeviceEndpoint": false,"publicServiceEndpoint": false}
`),
						ghttp.RespondWith(http.StatusCreated, `{							 	
							 "id": "f91adfe2-76c9-4649-939e-b01c37a3704c"
						}`),
					),
				)
			})

			It("should return cluster created", func() {
				params := ClusterCreateRequest{
					Name: "testservice", Datacenter: "dal10", MachineType: "b2c.4x16", PublicVlan: "vlan", PrivateVlan: "vlan", MasterVersion: "1.8.1", Prefix: "worker", WorkerNum: 1, DiskEncryption: true, EnableTrusted: true,
				}
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				myCluster, err := newCluster(server.URL()).Create(params, target)
				Expect(err).NotTo(HaveOccurred())
				Expect(myCluster).ShouldNot(BeNil())
				Expect(myCluster.ID).Should(Equal("f91adfe2-76c9-4649-939e-b01c37a3704c"))
			})
		})
		Context("When creation is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/clusters"),
						ghttp.VerifyJSON(`{"dataCenter":"dal10","isolation":"","machineType":"free","name":"testservice","privateVlan":"vlan","publicVlan":"vlan","workerNum":1,"noSubnet":false,"masterVersion":"1.8.1","prefix":"worker","diskEncryption": false,"enableTrusted":false,"privateSeviceEndpoint": false,"publicServiceEndpoint": false}
`),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to create cluster`),
					),
				)
			})

			It("should return error during cluster creation", func() {
				params := ClusterCreateRequest{
					Name: "testservice", Datacenter: "dal10", MachineType: "free", PublicVlan: "vlan", PrivateVlan: "vlan", MasterVersion: "1.8.1", Prefix: "worker", WorkerNum: 1,
				}
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				myCluster, err := newCluster(server.URL()).Create(params, target)
				Expect(err).To(HaveOccurred())
				Expect(myCluster.ID).Should(Equal(""))
			})
		})
	})
	//List
	Describe("List", func() {
		Context("When read of clusters is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/clusters"),
						ghttp.RespondWith(http.StatusOK, `[{							 	
              "CreatedDate": "",
              "DataCenter": "dal10",
              "ID": "f91adfe2-76c9-4649-939e-b01c37a3704",
              "IngressHostname": "",
              "IngressSecretName": "",
              "Location": "",
              "MasterKubeVersion": "1.8.1",
              "Prefix": "worker",
              "ModifiedDate": "",
              "Name": "test",
              "Region": "abc",
              "ServerURL": "",
              "State": "normal",
              "IsPaid": false,
              "IsTrusted": true,
              "WorkerCount": 1
              }]`),
					),
				)
			})

			It("should return cluster list", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				myCluster, err := newCluster(server.URL()).List(target)
				Expect(myCluster).ShouldNot(BeNil())
				for _, cluster := range myCluster {
					Expect(err).NotTo(HaveOccurred())
					Expect(cluster.ID).Should(Equal("f91adfe2-76c9-4649-939e-b01c37a3704"))
					Expect(cluster.WorkerCount).Should(Equal(1))
					Expect(cluster.MasterKubeVersion).Should(Equal("1.8.1"))
				}
			})
		})
		Context("When read of clusters is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/clusters"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve clusters`),
					),
				)
			})

			It("should return error when cluster are retrieved", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				myCluster, err := newCluster(server.URL()).List(target)
				Expect(err).To(HaveOccurred())
				Expect(myCluster).Should(BeNil())
			})
		})
	})
	//RefreshAPIServers
	Describe("RefreshAPIServers", func() {
		Context("When refresh of api servers of cluster is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, "/v1/clusters/test/masters"),
						ghttp.RespondWith(http.StatusOK, `{							
						}`),
					),
				)
			})

			It("should refresh api servers", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				err := newCluster(server.URL()).RefreshAPIServers("test", target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When refresh of api servers of cluster is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, "/v1/clusters/test/masters"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to refresh api servers`),
					),
				)
			})

			It("should return error failed to refresh api servers", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				err := newCluster(server.URL()).RefreshAPIServers("test", target)
				Expect(err).To(HaveOccurred())
			})
		})
	})
	//Delete
	Describe("Delete", func() {
		Context("When delete of cluster is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/clusters/test"),
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
				err := newCluster(server.URL()).Delete("test", target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When cluster delete is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/clusters/test"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to delete service key`),
					),
				)
			})

			It("should return error service key delete", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				err := newCluster(server.URL()).Delete("test", target)
				Expect(err).To(HaveOccurred())
			})
		})
	})
	//Find
	Describe("Find", func() {
		Context("When read of cluster is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/clusters/test"),
						ghttp.RespondWith(http.StatusOK, `{							 	
              "CreatedDate": "",
              "DataCenter": "dal10",
              "ID": "f91adfe2-76c9-4649-939e-b01c37a3704",
              "IngressHostname": "",
              "IngressSecretName": "",
              "Location": "",
              "MasterKubeVersion": "",
              "ModifiedDate": "",
              "Name": "test",
              "Region": "abc",
              "ServerURL": "",
              "State": "normal",
              "IsPaid": false,
              "IsTrusted": true,
              "ResourceGroup": "abcd",
              "WorkerCount": 1,
              "workerZones": [
				    "zone"
			  ],
              "Vlans": [{
			  "ID": "177453",
				"Subnets": [
				{
				"Cidr": "159.8.226.208/29",
				"ID": "1541737",
				"Ips": ["159.8.226.210"],
				"Is_ByOIP": false,
				"Is_Public": true
				}]
				}]}`),
					),
				)
			})

			It("should return cluster", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				myCluster, err := newCluster(server.URL()).Find("test", target)
				Expect(err).NotTo(HaveOccurred())
				Expect(myCluster).ShouldNot(BeNil())
				Expect(myCluster.Vlans[0].ID).Should(Equal("177453"))
				Expect(myCluster.Vlans[0].Subnets[0].ID).Should(Equal("1541737"))
				Expect(myCluster.Vlans[0].Subnets[0].Cidr).Should(Equal("159.8.226.208/29"))
				Expect(myCluster.Vlans[0].Subnets[0].IsPublic).Should(Equal(true))
				Expect(myCluster.ResourceGroupID).Should(Equal("abcd"))
			})
		})
		Context("When cluster retrieve is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/clusters/test"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve cluster`),
					),
				)
			})

			It("should return error when cluster is retrieved", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				myCluster, err := newCluster(server.URL()).Find("test", target)
				Expect(err).To(HaveOccurred())
				Expect(myCluster.ID).Should(Equal(""))
			})
		})
	})
	//set credentials
	Describe("set credentials", func() {
		Context("When credential set is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/credentials"),
						ghttp.RespondWith(http.StatusOK, `{}`),
					),
				)
			})

			It("should set credentials", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				err := newCluster(server.URL()).SetCredentials("test", "abcdef-df-fg", target)
				Expect(err).NotTo(HaveOccurred())

			})
		})
		Context("When credential set is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/credentials"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to set credentials`),
					),
				)
			})

			It("should throw error when setting credentials", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				err := newCluster(server.URL()).SetCredentials("test", "abcdef-df-fg", target)
				Expect(err).To(HaveOccurred())

			})
		})
	})
	//Unset credentials
	Describe("unset credentials", func() {
		Context("When unset credential is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/credentials"),
						ghttp.RespondWith(http.StatusOK, `{}`),
					),
				)
			})

			It("should set credentials", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				err := newCluster(server.URL()).UnsetCredentials(target)
				Expect(err).NotTo(HaveOccurred())

			})
		})
		Context("When unset credential is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/credentials"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to unset credentials`),
					),
				)
			})

			It("should set credentials", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				err := newCluster(server.URL()).UnsetCredentials(target)
				Expect(err).To(HaveOccurred())

			})
		})
	})
	//Bind service
	Describe("Bind service", func() {
		Context("When bind service is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/clusters/test/services"),
						ghttp.RespondWith(http.StatusOK, `{}`),
					),
				)
			})

			It("should bind service to a cluster", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				params := ServiceBindRequest{
					ClusterNameOrID: "test", ServiceInstanceNameOrID: "cloudantDB", NamespaceID: "default"}
				serviceResp, err := newCluster(server.URL()).BindService(params, target)
				Expect(err).NotTo(HaveOccurred())
				Expect(serviceResp).ShouldNot(BeNil())
			})
		})
		Context("When bind service is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/clusters/test/services"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to set credentials`),
					),
				)
			})

			It("should throw error when binding service to a cluster", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				params := ServiceBindRequest{
					ClusterNameOrID: "test", ServiceInstanceNameOrID: "cloudantDB", NamespaceID: "default"}
				serviceResp, err := newCluster(server.URL()).BindService(params, target)
				Expect(err).To(HaveOccurred())
				Expect(serviceResp.ServiceInstanceGUID).Should(Equal(""))
				Expect(serviceResp.SecretName).Should(Equal(""))
			})
		})
	})
	//Unbind service
	Describe("UnBind service", func() {
		Context("When bind service is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/clusters/test/services/default/cloudantDB"),
						ghttp.RespondWith(http.StatusOK, `{}`),
					),
				)
			})

			It("should bind service to a cluster", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				err := newCluster(server.URL()).UnBindService("test", "default", "cloudantDB", target)
				Expect(err).NotTo(HaveOccurred())

			})
		})
		Context("When unbind service is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/clusters/test/services/default/cloudantDB"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to unbind service`),
					),
				)
			})

			It("should set credentials", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				err := newCluster(server.URL()).UnBindService("test", "default", "cloudantDB", target)
				Expect(err).To(HaveOccurred())

			})
		})
	})
	//List bound services
	Describe("ListServicesBoundToCluster", func() {
		Context("When read of cluster services is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/clusters/test/services/default"),
						ghttp.RespondWith(http.StatusOK, `[{							 	
              "ServiceName": "testService",
              "ServiceID": "f91adfe2-76c9-4649-939e-b01c37a3704",
              "ServiceKeyName": "kube-testService",
              "Namespace": "default"
              }]`),
					),
				)
			})

			It("should return cluster service list", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				boundServices, err := newCluster(server.URL()).ListServicesBoundToCluster("test", "default", target)
				Expect(boundServices).ShouldNot(BeNil())
				for _, service := range boundServices {
					Expect(err).NotTo(HaveOccurred())
					Expect(service.ServiceName).Should(Equal("testService"))
					Expect(service.ServiceID).Should(Equal("f91adfe2-76c9-4649-939e-b01c37a3704"))
				}
			})
		})
		Context("When read of cluster services is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/clusters/test/services/default"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve clusters`),
					),
				)
			})

			It("should return error when cluster services are retrieved", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				service, err := newCluster(server.URL()).ListServicesBoundToCluster("test", "default", target)
				Expect(err).To(HaveOccurred())
				Expect(service).Should(BeNil())
			})
		})
	})
	//Find Cluster service
	Describe("FindServiceBoundToClusters", func() {
		Context("When read a service bound to cluster is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/clusters/test/services/default"),
						ghttp.RespondWith(http.StatusOK, `[{							 	
              "ServiceName": "testService",
              "ServiceID": "f91adfe2-76c9-4649-939e-b01c37a3704",
              "ServiceKeyName": "kube-testService",
              "Namespace": "default"
              }]`),
					),
				)
			})

			It("should return cluster service list", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				boundService, err := newCluster(server.URL()).FindServiceBoundToCluster("test", "f91adfe2-76c9-4649-939e-b01c37a3704", "default", target)
				Expect(boundService).ShouldNot(BeNil())
				Expect(err).NotTo(HaveOccurred())
				Expect(boundService.ServiceName).Should(Equal("testService"))
				Expect(boundService.ServiceID).Should(Equal("f91adfe2-76c9-4649-939e-b01c37a3704"))
			})
		})
		Context("When read of cluster services is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/clusters/test/services/default"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve clusters`),
					),
				)
			})

			It("should return error when cluster services are retrieved", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				_, err := newCluster(server.URL()).FindServiceBoundToCluster("test", "f91adfe2-76c9-4649-939e-b01c37a3704", "default", target)
				Expect(err).To(HaveOccurred())
			})
		})
	})
	//UpdateClusterWorker
	Describe("UpdateClusterWorker", func() {
		Context("When updating cluster workers is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, "/v1/clusters/test/workers/w1"),
						ghttp.RespondWith(http.StatusNoContent, `{}`),
					),
				)
			})

			It("should return cluster version updated", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				params := UpdateWorkerCommand{
					Action: "reload",
				}
				err := newCluster(server.URL()).UpdateClusterWorker("test", "w1", params, target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When updating cluster workers is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, "/v1/clusters/test/workers/w1"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to update cluster workers`),
					),
				)
			})

			It("should return error during updating cluster version", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				params := UpdateWorkerCommand{
					Action: "reload",
				}
				err := newCluster(server.URL()).UpdateClusterWorker("test", "w1", params, target)
				Expect(err).To(HaveOccurred())
			})
		})
	})
	//Update
	Describe("Update", func() {
		Context("When updating cluster version is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, "/v1/clusters/test"),
						ghttp.RespondWith(http.StatusNoContent, `{}`),
					),
				)
			})

			It("should return cluster version updated", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				params := ClusterUpdateParam{
					Action:  "update",
					Force:   true,
					Version: "1.8.6",
				}
				err := newCluster(server.URL()).Update("test", params, target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When updating cluster version is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, "/v1/clusters/test"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to update cluster version`),
					),
				)
			})

			It("should return error during updating cluster version", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
				}
				params := ClusterUpdateParam{
					Action:  "update",
					Force:   true,
					Version: "1.8.6",
				}
				err := newCluster(server.URL()).Update("test", params, target)
				Expect(err).To(HaveOccurred())
			})
		})
	})
	//
})

func newCluster(url string) Clusters {

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
	return newClusterAPI(&client)
}
