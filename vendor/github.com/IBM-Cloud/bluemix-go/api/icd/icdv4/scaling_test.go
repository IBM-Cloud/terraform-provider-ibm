package icdv4

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

var _ = Describe("Scaling", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})
	Describe("Update", func() {
		Context("When update group is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPatch, "/v4/ibm/deployments/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/groups/member"),
						ghttp.RespondWith(http.StatusOK, `
                           {
                              "task": {
                                "id": "5abb6a7d11a1a5001479a0ab",
                                "description": "Scaling database deployment",
                                "status": "running",
                                "deployment_id": "59b14b19874a1c0018009482",
                                "progress_percent": 5,
                                "created_at": "2018-03-28T10:20:30Z"
                              }
                            }
                        `),
					),
				)
			})
			It("should return group updated", func() {
				target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				target2 := "member"
				memoryReq := MemoryReq{AllocationMb: 2018}
				cpuReq := CpuReq{AllocationCount: 2}
				diskReq := DiskReq{AllocationMb: 2018}
				groupBdy := GroupBdy{
					Memory: &memoryReq,
					Disk:   &diskReq,
					Cpu:    &cpuReq,
				}
				params := GroupReq{
					GroupBdy: groupBdy,
				}
				myTask, err := newGroup(server.URL()).UpdateGroup(target1, target2, params)
				Expect(err).NotTo(HaveOccurred())
				Expect(myTask).ShouldNot(BeNil())
				Expect(myTask.Id).Should(Equal("5abb6a7d11a1a5001479a0ab"))
				Expect(myTask.Description).Should(Equal("Scaling database deployment"))
				Expect(myTask.Status).Should(Equal("running"))
				Expect(myTask.DeploymentId).Should(Equal("59b14b19874a1c0018009482"))
				Expect(myTask.ProgressPercent).Should(Equal(5))
				Expect(myTask.CreatedAt).Should(Equal("2018-03-28T10:20:30Z"))
			})
		})
		Context("When update is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPatch, "/v4/ibm/deployments/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/groups/member"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to update group`),
					),
				)
			})

			It("should return error during group update", func() {
				target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				target2 := "member"
				memoryReq := MemoryReq{AllocationMb: 2018}
				cpuReq := CpuReq{AllocationCount: 2}
				diskReq := DiskReq{AllocationMb: 2018}
				groupBdy := GroupBdy{
					Memory: &memoryReq,
					Disk:   &diskReq,
					Cpu:    &cpuReq,
				}
				params := GroupReq{
					GroupBdy: groupBdy,
				}
				myTask, err := newGroup(server.URL()).UpdateGroup(target1, target2, params)
				Expect(err).To(HaveOccurred())
				Expect(myTask.Id).Should(Equal(""))
			})
		})
	})
	Describe("GetDefault", func() {
		Context("When get default is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v4/ibm/deployables/etcd/groups"),
						ghttp.RespondWith(http.StatusOK, `
                                {
                                  "groups": [
                                    {
                                      "id": "member",
                                      "count": 2,
                                      "memory": {
                                        "units": "mb",
                                        "allocation_mb": 2048,
                                        "minimum_mb": 2048,
                                        "step_size_mb": 256,
                                        "is_adjustable": true
                                      },
                                      "cpu": {
                                        "units": "2",
                                        "allocation_count": 2,
                                        "minimum_count": 2,
                                        "step_size_count": 2,
                                        "is_adjustable": false
                                      },
                                      "disk": {
                                        "units": "mb",
                                        "allocation_mb": 5120,
                                        "minimum_mb": 5120,
                                        "step_size_mb": 2048,
                                        "is_adjustable": false
                                      }
                                    }
                                  ]
                                }
                        `),
					),
				)
			})

			It("should return default groups", func() {
				target1 := "etcd"
				groupList, err := newGroup(server.URL()).GetDefaultGroups(target1)
				Expect(err).NotTo(HaveOccurred())
				Expect(groupList).ShouldNot(BeNil())
				Expect(groupList.Groups[0].Id).Should(Equal("member"))
				Expect(groupList.Groups[0].Count).Should(Equal(2))
				Expect(groupList.Groups[0].Memory.Units).Should(Equal("mb"))
				Expect(groupList.Groups[0].Memory.AllocationMb).Should(Equal(2048))
				Expect(groupList.Groups[0].Memory.MinimumMb).Should(Equal(2048))
				Expect(groupList.Groups[0].Memory.StepSizeMb).Should(Equal(256))
				Expect(groupList.Groups[0].Memory.IsAdjustable).Should(Equal(true))
				Expect(groupList.Groups[0].Cpu.Units).Should(Equal("2"))
				Expect(groupList.Groups[0].Cpu.AllocationCount).Should(Equal(2))
				Expect(groupList.Groups[0].Cpu.MinimumCount).Should(Equal(2))
				Expect(groupList.Groups[0].Cpu.StepSizeCount).Should(Equal(2))
				Expect(groupList.Groups[0].Cpu.IsAdjustable).Should(Equal(false))
				Expect(groupList.Groups[0].Disk.Units).Should(Equal("mb"))
				Expect(groupList.Groups[0].Disk.AllocationMb).Should(Equal(5120))
				Expect(groupList.Groups[0].Disk.MinimumMb).Should(Equal(5120))
				Expect(groupList.Groups[0].Disk.StepSizeMb).Should(Equal(2048))
				Expect(groupList.Groups[0].Disk.IsAdjustable).Should(Equal(false))
			})
		})
		Context("When get default is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v4/ibm/deployables/etcd/groups"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to get groups`),
					),
				)
			})

			It("should return error during group get", func() {
				target1 := "etcd"
				_, err := newGroup(server.URL()).GetDefaultGroups(target1)
				Expect(err).To(HaveOccurred())
			})
		})
	})
	Describe("Get", func() {
		Context("When get is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v4/ibm/deployments/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/groups"),
						ghttp.RespondWith(http.StatusOK, `
                                {
                                  "groups": [
                                    {
                                      "id": "member",
                                      "count": 2,
                                      "memory": {
                                        "units": "mb",
                                        "allocation_mb": 2048,
                                        "minimum_mb": 2048,
                                        "step_size_mb": 256,
                                        "is_adjustable": true
                                      },
                                      "cpu": {
                                        "units": "2",
                                        "allocation_count": 2,
                                        "minimum_count": 2,
                                        "step_size_count": 2,
                                        "is_adjustable": false
                                      },
                                      "disk": {
                                        "units": "mb",
                                        "allocation_mb": 5120,
                                        "minimum_mb": 5120,
                                        "step_size_mb": 2048,
                                        "is_adjustable": false
                                      }
                                    }
                                  ]
                                }
                        `),
					),
				)
			})

			It("should return groups", func() {
				target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				groupList, err := newGroup(server.URL()).GetGroups(target1)
				Expect(err).NotTo(HaveOccurred())
				Expect(groupList).ShouldNot(BeNil())
				Expect(groupList.Groups[0].Id).Should(Equal("member"))
				Expect(groupList.Groups[0].Count).Should(Equal(2))
				Expect(groupList.Groups[0].Memory.Units).Should(Equal("mb"))
				Expect(groupList.Groups[0].Memory.AllocationMb).Should(Equal(2048))
				Expect(groupList.Groups[0].Memory.MinimumMb).Should(Equal(2048))
				Expect(groupList.Groups[0].Memory.StepSizeMb).Should(Equal(256))
				Expect(groupList.Groups[0].Memory.IsAdjustable).Should(Equal(true))
				Expect(groupList.Groups[0].Cpu.Units).Should(Equal("2"))
				Expect(groupList.Groups[0].Cpu.AllocationCount).Should(Equal(2))
				Expect(groupList.Groups[0].Cpu.MinimumCount).Should(Equal(2))
				Expect(groupList.Groups[0].Cpu.StepSizeCount).Should(Equal(2))
				Expect(groupList.Groups[0].Cpu.IsAdjustable).Should(Equal(false))
				Expect(groupList.Groups[0].Disk.Units).Should(Equal("mb"))
				Expect(groupList.Groups[0].Disk.AllocationMb).Should(Equal(5120))
				Expect(groupList.Groups[0].Disk.MinimumMb).Should(Equal(5120))
				Expect(groupList.Groups[0].Disk.StepSizeMb).Should(Equal(2048))
				Expect(groupList.Groups[0].Disk.IsAdjustable).Should(Equal(false))
			})
		})
		Context("When get default is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v4/ibm/deployments/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/groups"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to get groups`),
					),
				)
			})

			It("should return error during group get", func() {
				target1 := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				_, err := newGroup(server.URL()).GetGroups(target1)
				Expect(err).To(HaveOccurred())
			})
		})
	})
})

func newGroup(url string) Groups {

	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}
	conf := sess.Config.Copy()
	conf.HTTPClient = bluemixHttp.NewHTTPClient(conf)
	conf.Endpoint = &url

	client := client.Client{
		Config:      conf,
		ServiceName: bluemix.ICDService,
	}
	return newGroupAPI(&client)
}
