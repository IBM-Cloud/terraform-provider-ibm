package registryv1

import (
	"fmt"
	"log"
	"net/http"
	"time"

	ibmcloud "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/client"
	ibmcloudHttp "github.com/IBM-Cloud/bluemix-go/http"
	"github.com/IBM-Cloud/bluemix-go/session"

	"github.com/onsi/gomega/ghttp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	imageName = "registry.ng.bluemix.net/gpfs/sklm:2.7"
	imageList = `[
    {
        "Id": "sha256:1002cd429f26a9122df60be5f541a12f09f960b1bd092998f70594e4d8450be8",
        "ParentId": "sha256:48b0f237fb3623275a4e96fddd0aec60067a6abee1b94725b292529124db8f86",
        "DigestTags": {
            "sha256:93390ca2d98a4da78a4496a77e14f751f30e95a8a7c2c172e89c1ca069ad2ef3": [
                "registry.ng.bluemix.net/gpfs/sklm:2.7"
            ]
        },
        "RepoTags": [
            "registry.ng.bluemix.net/gpfs/sklm:2.7"
        ],
        "RepoDigests": [
            "registry.ng.bluemix.net/gpfs/sklm@sha256:93390ca2d98a4da78a4496a77e14f751f30e95a8a7c2c172e89c1ca069ad2ef3"
        ],
        "Created": 1531320068,
        "Size": 2257318704,
        "VirtualSize": 2257318704,
        "Labels": {
            "architecture": "x86_64",
            "authoritative-source-url": "registry.access.redhat.com"
        },
        "Vulnerable": "true",
        "VulnerabilityCount": 24,
        "ConfigurationIssueCount": 0,
        "IssueCount": 24,
        "ExemptIssueCount": 0
	}
]`
	imageInspect = `{
	"Id": "sha256:1002cd429f26a9122df60be5f541a12f09f960b1bd092998f70594e4d8450be8",
	"Parent": "sha256:48b0f237fb3623275a4e96fddd0aec60067a6abee1b94725b292529124db8f86",
	"Comment": "",
	"Created": "2018-07-11T14:41:08.6719209Z",
	"Container": "56b721a4dcad2778022df5b1b91e33a885579c29c7d00f7ce76065cd7dd35727",
	"ContainerConfig": {
		"Hostname": "36ee326478e1",
		"Domainname": "",
		"User": "",
		"AttachStdin": false,
		"AttachStdout": false,
		"AttachStderr": false,
		"ExposedPorts": {
			"443/tcp": {},
			"5696/tcp": {},
			"80/tcp": {},
			"9083/tcp": {}
		},
		"Tty": true,
		"OpenStdin": false,
		"StdinOnce": false,
		"Env": [
			"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
			"container=oci"
		],
		"Cmd": [
			"/bin/sh",
			"-c",
			"#(nop) ",
			"CMD [\"/home/klmfcusr/run_sklm.sh\"]"
		],
		"ArgsEscaped": true,
		"Image": "sha256:48b0f237fb3623275a4e96fddd0aec60067a6abee1b94725b292529124db8f86",
		"Volumes": {
			"/home/sklmdb": {}
		},
		"WorkingDir": "/home/klmfcusr",
		"Entrypoint": null,
		"OnBuild": null,
		"Labels": {
			"architecture": "x86_64",
			"authoritative-source-url": "registry.access.redhat.com"
		}
	},
	"DockerVersion": "18.03.1-ce",
	"Author": "",
	"Config": {
		"Hostname": "36ee326478e1",
		"Domainname": "",
		"User": "",
		"AttachStdin": false,
		"AttachStdout": false,
		"AttachStderr": false,
		"ExposedPorts": {
			"443/tcp": {},
			"5696/tcp": {},
			"80/tcp": {},
			"9083/tcp": {}
		},
		"Tty": true,
		"OpenStdin": false,
		"StdinOnce": false,
		"Env": [
			"PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
			"container=oci"
		],
		"Cmd": [
			"/home/klmfcusr/run_sklm.sh"
		],
		"ArgsEscaped": true,
		"Image": "sha256:48b0f237fb3623275a4e96fddd0aec60067a6abee1b94725b292529124db8f86",
		"Volumes": {
			"/home/sklmdb": {}
		},
		"WorkingDir": "/home/klmfcusr",
		"Entrypoint": null,
		"OnBuild": null,
		"Labels": {
			"architecture": "x86_64",
			"authoritative-source-url": "registry.access.redhat.com"
		}
	},
	"Architecture": "amd64",
	"Os": "linux",
	"Size": 2257318704,
	"VirtualSize": 2257318704,
	"RootFS": {
		"Type": "layers",
		"Layers": [
			"sha256:f4fa6c253d2ff944ef6975be17cd0bb59896b386f9e2b737539400a37a68a80b",
			"sha256:d6a4dd6ace1f76d1410e389c23e515a09eda880da05850b4343e2b39b6ced363",
			"sha256:f98e45f3ca63148407d3800280a4faa7d6b6a7f98f39367e873d496d68293ffd",
			"sha256:f98e45f3ca63148407d3800280a4faa7d6b6a7f98f39367e873d496d68293ffd",
			"sha256:e1844b32012a6cb9c5974a10e903c1098e359b5b6b13bd3fc0dd331b1d455313",
			"sha256:27a571001b44c6d3498e2e37b7310652c62e0ab5cc1f77d33807b83cac5449dc",
			"sha256:55dc1ec7f17c842a4da60d7b1aaa976eb96eabd9f0e9a7376d0a473d1835472f"
		]
	}
}`
	imageVulnerability = `{
	"metadata": {
		"namespace": "",
		"complete": true,
		"crawled_time": "2018-10-18T02:32:59Z",
		"os_supported": true
	},
	"summary": {
		"malware": {
			"compliant": true,
			"reason": ""
		},
		"compliance": {
			"compliance_violations": 5,
			"reason": "",
			"compliant": false,
			"total_compliance_rules": 0,
			"execution_status": ""
		},
		"secureconfig": {
			"misconfigured": 0,
			"correct_output": 0,
			"total_output_docs": 0
		},
		"vulnerability": {
			"total_packages": 0,
			"total_usns_for_distro": 0,
			"vulnerable_usns": 0,
			"vulnerable_packages": 10
		}
	},
	"detail": {
		"compliance": [
			{
				"reason": "File /etc/pam.d/common-password not found",
				"compliant": false,
				"description": "Minimum password length must be 8.",
				"policy_mandated": false
			}
		],
		"vulnerability": [
			{
				"package_name": "firefox 60.1.0-4.el7_5 has vulnerabilities",
				"vulnerabilities": [
					{
						"url": "https://access.redhat.com/errata/RHSA-2018:2692",
						"cveid": [
							"CVE-2017-16541",
							"CVE-2018-12376",
							"CVE-2018-12377",
							"CVE-2018-12378",
							"CVE-2018-12379"
						],
						"summary": "(RHSA-2018:2692) Critical: firefox security update"
					}
				]
			}
		]
	}
}`
	imageDelete = `{ "Untagged": "sha256:1002cd429f26a9122df60be5f541a12f09f960b1bd092998f70594e4d8450be8" }`
)

var _ = Describe("Images", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})

	Describe("GetImages", func() {
		Context("When get images is completed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/api/v1/images"),
						ghttp.RespondWith(http.StatusOK, imageList),
					),
				)
			})

			It("should return get images results", func() {
				params := GetImageRequest{
					IncludeIBM:      false,
					IncludePrivate:  true,
					Namespace:       "",
					Repository:      "",
					Vulnerabilities: true,
				}
				target := ImageTargetHeader{
					AccountID: "abc",
				}
				respptr, err := newImages(server.URL()).GetImages(params, target)
				Expect(err).NotTo(HaveOccurred())
				Expect(respptr).NotTo(BeNil())
				resp := *respptr
				Expect(resp).To(HaveLen(1))
				Expect(resp[0].ID).Should(Equal("sha256:1002cd429f26a9122df60be5f541a12f09f960b1bd092998f70594e4d8450be8"))
				Expect(resp[0].ParentID).Should(Equal("sha256:48b0f237fb3623275a4e96fddd0aec60067a6abee1b94725b292529124db8f86"))
				Expect(resp[0].DigestTags).Should(HaveKey("sha256:93390ca2d98a4da78a4496a77e14f751f30e95a8a7c2c172e89c1ca069ad2ef3"))
				Expect(resp[0].DigestTags["sha256:93390ca2d98a4da78a4496a77e14f751f30e95a8a7c2c172e89c1ca069ad2ef3"]).To(HaveLen(1))
				Expect(resp[0].DigestTags["sha256:93390ca2d98a4da78a4496a77e14f751f30e95a8a7c2c172e89c1ca069ad2ef3"][0]).Should(Equal("registry.ng.bluemix.net/gpfs/sklm:2.7"))
				Expect(resp[0].RepoTags).To(HaveLen(1))
				Expect(resp[0].RepoTags[0]).Should(Equal("registry.ng.bluemix.net/gpfs/sklm:2.7"))
				Expect(resp[0].RepoDigests).To(HaveLen(1))
				Expect(resp[0].RepoDigests[0]).Should(Equal("registry.ng.bluemix.net/gpfs/sklm@sha256:93390ca2d98a4da78a4496a77e14f751f30e95a8a7c2c172e89c1ca069ad2ef3"))
				Expect(resp[0].Created).Should(Equal(1531320068))
				Expect(resp[0].Size).Should(Equal(int64(2257318704)))
				Expect(resp[0].VirtualSize).Should(Equal(int64(2257318704)))
				Expect(resp[0].Labels).Should(HaveKey("architecture"))
				Expect(resp[0].Labels["architecture"]).Should(Equal("x86_64"))
				Expect(resp[0].Labels).Should(HaveKey("authoritative-source-url"))
				Expect(resp[0].Labels["authoritative-source-url"]).Should(Equal("registry.access.redhat.com"))
				Expect(resp[0].Vulnerable).Should(Equal("true"))
				Expect(resp[0].VulnerabilityCount).Should(Equal(24))
				Expect(resp[0].ConfigurationIssueCount).Should(Equal(0))
				Expect(resp[0].IssueCount).Should(Equal(24))
				Expect(resp[0].ExemptIssueCount).Should(Equal(0))
			})
		})
		Context("When get image fails", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/api/v1/images"),
						ghttp.RespondWith(http.StatusInternalServerError, `Internal Error`),
					),
				)
			})

			It("should return error when images are retrieved", func() {
				params := GetImageRequest{
					IncludeIBM:      false,
					IncludePrivate:  true,
					Namespace:       "",
					Repository:      "",
					Vulnerabilities: true,
				}
				target := ImageTargetHeader{
					AccountID: "abc",
				}
				resp, err := newImages(server.URL()).GetImages(params, target)
				Expect(err).To(HaveOccurred())
				Expect(resp).Should(BeNil())
			})
		})
	})

	Describe("InspectImage", func() {
		Context("When Inspect Image is completed", func() {

			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, fmt.Sprintf("/api/v1/images/%s/json", imageName)),
						ghttp.RespondWith(http.StatusOK, imageInspect),
					),
				)
			})

			It("should return image inspect results", func() {
				target := ImageTargetHeader{
					AccountID: "abc",
				}
				respptr, err := newImages(server.URL()).InspectImage(imageName, target)
				Expect(err).NotTo(HaveOccurred())
				Expect(respptr).NotTo(BeNil())
				resp := *respptr
				Expect(resp.ID).Should(Equal("sha256:1002cd429f26a9122df60be5f541a12f09f960b1bd092998f70594e4d8450be8"))
				Expect(resp.Parent).Should(Equal("sha256:48b0f237fb3623275a4e96fddd0aec60067a6abee1b94725b292529124db8f86"))
				Expect(resp.Comment).Should(Equal(""))
				time1, _ := time.Parse(time.RFC3339, "2018-07-11T14:41:08.6719209Z")
				Expect(resp.Created).Should(Equal(time1))
				Expect(resp.Container).Should(Equal("56b721a4dcad2778022df5b1b91e33a885579c29c7d00f7ce76065cd7dd35727"))
				Expect(resp.ContainerConfig.Hostname).Should(Equal("36ee326478e1"))
				Expect(resp.ContainerConfig.Domainname).Should(Equal(""))
				Expect(resp.ContainerConfig.User).Should(Equal(""))
				Expect(resp.ContainerConfig.AttachStdin).Should(Equal(false))
				Expect(resp.ContainerConfig.AttachStdout).Should(Equal(false))
				Expect(resp.ContainerConfig.AttachStderr).Should(Equal(false))
				Expect(resp.ContainerConfig.ExposedPorts).Should(HaveKey("443/tcp"))
				Expect(resp.ContainerConfig.ExposedPorts["443/tcp"]).To(HaveLen(0))
				Expect(resp.ContainerConfig.ExposedPorts).Should(HaveKey("5696/tcp"))
				Expect(resp.ContainerConfig.ExposedPorts["5696/tcp"]).To(HaveLen(0))
				Expect(resp.ContainerConfig.ExposedPorts).Should(HaveKey("80/tcp"))
				Expect(resp.ContainerConfig.ExposedPorts["80/tcp"]).To(HaveLen(0))
				Expect(resp.ContainerConfig.ExposedPorts).Should(HaveKey("9083/tcp"))
				Expect(resp.ContainerConfig.ExposedPorts["9083/tcp"]).To(HaveLen(0))
				Expect(resp.ContainerConfig.Tty).Should(Equal(true))
				Expect(resp.ContainerConfig.OpenStdin).Should(Equal(false))
				Expect(resp.ContainerConfig.StdinOnce).Should(Equal(false))
				Expect(resp.ContainerConfig.Env).To(HaveLen(2))
				Expect(resp.ContainerConfig.Env[0]).Should(Equal("PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"))
				Expect(resp.ContainerConfig.Env[1]).Should(Equal("container=oci"))
				Expect(resp.ContainerConfig.Cmd).To(HaveLen(4))
				Expect(resp.ContainerConfig.Cmd[0]).Should(Equal("/bin/sh"))
				Expect(resp.ContainerConfig.Cmd[1]).Should(Equal("-c"))
				Expect(resp.ContainerConfig.Cmd[2]).Should(Equal("#(nop) "))
				Expect(resp.ContainerConfig.Cmd[3]).Should(Equal("CMD [\"/home/klmfcusr/run_sklm.sh\"]"))
				Expect(resp.ContainerConfig.ArgsEscaped).Should(Equal(true))
				Expect(resp.ContainerConfig.Image).Should(Equal("sha256:48b0f237fb3623275a4e96fddd0aec60067a6abee1b94725b292529124db8f86"))
				Expect(resp.ContainerConfig.Volumes).Should(HaveKey("/home/sklmdb"))
				Expect(resp.ContainerConfig.Volumes["/home/sklmdb"]).To(HaveLen(0))
				Expect(resp.ContainerConfig.WorkingDir).Should(Equal("/home/klmfcusr"))
				Expect(resp.ContainerConfig.Entrypoint).To(HaveLen(0))
				Expect(resp.ContainerConfig.OnBuild).To(HaveLen(0))
				Expect(resp.ContainerConfig.Labels).Should(HaveKey("architecture"))
				Expect(resp.ContainerConfig.Labels["architecture"]).Should(Equal("x86_64"))
				Expect(resp.ContainerConfig.Labels).Should(HaveKey("authoritative-source-url"))
				Expect(resp.ContainerConfig.Labels["authoritative-source-url"]).Should(Equal("registry.access.redhat.com"))
				Expect(resp.DockerVersion).Should(Equal("18.03.1-ce"))
				Expect(resp.Author).Should(Equal(""))
				Expect(resp.Config.Hostname).Should(Equal("36ee326478e1"))
				Expect(resp.Config.Domainname).Should(Equal(""))
				Expect(resp.Config.User).Should(Equal(""))
				Expect(resp.Config.AttachStdin).Should(Equal(false))
				Expect(resp.Config.AttachStdout).Should(Equal(false))
				Expect(resp.Config.AttachStderr).Should(Equal(false))
				Expect(resp.Config.ExposedPorts).Should(HaveKey("443/tcp"))
				Expect(resp.Config.ExposedPorts["443/tcp"]).To(HaveLen(0))
				Expect(resp.Config.ExposedPorts).Should(HaveKey("5696/tcp"))
				Expect(resp.Config.ExposedPorts["5696/tcp"]).To(HaveLen(0))
				Expect(resp.Config.ExposedPorts).Should(HaveKey("80/tcp"))
				Expect(resp.Config.ExposedPorts["80/tcp"]).To(HaveLen(0))
				Expect(resp.Config.ExposedPorts).Should(HaveKey("9083/tcp"))
				Expect(resp.Config.ExposedPorts["9083/tcp"]).To(HaveLen(0))
				Expect(resp.Config.Tty).Should(Equal(true))
				Expect(resp.Config.OpenStdin).Should(Equal(false))
				Expect(resp.Config.StdinOnce).Should(Equal(false))
				Expect(resp.Config.Env).To(HaveLen(2))
				Expect(resp.Config.Env[0]).Should(Equal("PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"))
				Expect(resp.Config.Env[1]).Should(Equal("container=oci"))
				Expect(resp.Config.Cmd).To(HaveLen(1))
				Expect(resp.Config.Cmd[0]).Should(Equal("/home/klmfcusr/run_sklm.sh"))
				Expect(resp.Config.ArgsEscaped).Should(Equal(true))
				Expect(resp.Config.Image).Should(Equal("sha256:48b0f237fb3623275a4e96fddd0aec60067a6abee1b94725b292529124db8f86"))
				Expect(resp.Config.Volumes).Should(HaveKey("/home/sklmdb"))
				Expect(resp.Config.Volumes["/home/sklmdb"]).To(HaveLen(0))
				Expect(resp.Config.WorkingDir).Should(Equal("/home/klmfcusr"))
				Expect(resp.Config.Entrypoint).To(HaveLen(0))
				Expect(resp.Config.OnBuild).To(HaveLen(0))
				Expect(resp.Config.Labels).Should(HaveKey("architecture"))
				Expect(resp.Config.Labels["architecture"]).Should(Equal("x86_64"))
				Expect(resp.Config.Labels).Should(HaveKey("authoritative-source-url"))
				Expect(resp.Config.Labels["authoritative-source-url"]).Should(Equal("registry.access.redhat.com"))
				Expect(resp.Architecture).Should(Equal("amd64"))
				Expect(resp.Os).Should(Equal("linux"))
				Expect(resp.Size).Should(Equal(int64(2257318704)))
				Expect(resp.VirtualSize).Should(Equal(int64(2257318704)))
				Expect(resp.RootFS.Type).Should(Equal("layers"))
				Expect(resp.RootFS.Layers).To(HaveLen(7))
				Expect(resp.RootFS.Layers[0]).Should(Equal("sha256:f4fa6c253d2ff944ef6975be17cd0bb59896b386f9e2b737539400a37a68a80b"))
				Expect(resp.RootFS.Layers[1]).Should(Equal("sha256:d6a4dd6ace1f76d1410e389c23e515a09eda880da05850b4343e2b39b6ced363"))
				Expect(resp.RootFS.Layers[2]).Should(Equal("sha256:f98e45f3ca63148407d3800280a4faa7d6b6a7f98f39367e873d496d68293ffd"))
				Expect(resp.RootFS.Layers[3]).Should(Equal("sha256:f98e45f3ca63148407d3800280a4faa7d6b6a7f98f39367e873d496d68293ffd"))
				Expect(resp.RootFS.Layers[4]).Should(Equal("sha256:e1844b32012a6cb9c5974a10e903c1098e359b5b6b13bd3fc0dd331b1d455313"))
				Expect(resp.RootFS.Layers[5]).Should(Equal("sha256:27a571001b44c6d3498e2e37b7310652c62e0ab5cc1f77d33807b83cac5449dc"))
				Expect(resp.RootFS.Layers[6]).Should(Equal("sha256:55dc1ec7f17c842a4da60d7b1aaa976eb96eabd9f0e9a7376d0a473d1835472f"))
			})
		})
		Context("When inspect image fails", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, fmt.Sprintf("/api/v1/images/%s/json", imageName)),
						ghttp.RespondWith(http.StatusInternalServerError, `Internal Error`),
					),
				)
			})

			It("should return error when image is inspected", func() {
				target := ImageTargetHeader{
					AccountID: "abc",
				}
				resp, err := newImages(server.URL()).InspectImage(imageName, target)
				Expect(err).To(HaveOccurred())
				Expect(resp).Should(BeNil())
			})
		})
	})

	Describe("ImageVulnerabilities", func() {
		Context("When Scan Image is completed", func() {

			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, fmt.Sprintf("/api/v1/images/%s/vulnerabilities", imageName)),
						ghttp.RespondWith(http.StatusOK, imageVulnerability),
					),
				)
			})

			It("should return scan images results", func() {
				param := ImageVulnerabilitiesRequest{
					Advisory: true,
					All:      true,
				}
				target := ImageTargetHeader{
					AccountID: "abc",
				}
				respptr, err := newImages(server.URL()).ImageVulnerabilities(imageName, param, target)

				Expect(err).NotTo(HaveOccurred())
				Expect(respptr).NotTo(BeNil())
				resp := *respptr
				Expect(resp.Metadata.Namespace).Should(Equal(""))
				Expect(resp.Metadata.Complete).Should(Equal(true))
				time1, _ := time.Parse(time.RFC3339, "2018-10-18T02:32:59Z")
				Expect(resp.Metadata.CrawledTime).Should(Equal(time1))
				Expect(resp.Metadata.OsSupported).Should(Equal(true))
				Expect(resp.Summary.Malware.Compliant).Should(Equal(true))
				Expect(resp.Summary.Malware.Reason).Should(Equal(""))
				Expect(resp.Summary.Compliance.ComplianceViolations).Should(Equal(5))
				Expect(resp.Summary.Compliance.Reason).Should(Equal(""))
				Expect(resp.Summary.Compliance.Compliant).Should(Equal(false))
				Expect(resp.Summary.Compliance.TotalComplianceRules).Should(Equal(0))
				Expect(resp.Summary.Compliance.ExecutionStatus).Should(Equal(""))
				Expect(resp.Summary.Secureconfig.Misconfigured).Should(Equal(0))
				Expect(resp.Summary.Secureconfig.CorrectOutput).Should(Equal(0))
				Expect(resp.Summary.Secureconfig.TotalOutputDocs).Should(Equal(0))
				Expect(resp.Summary.Vulnerability.TotalPackages).Should(Equal(0))
				Expect(resp.Summary.Vulnerability.TotalUsnsForDistro).Should(Equal(0))
				Expect(resp.Summary.Vulnerability.VulnerableUsns).Should(Equal(0))
				Expect(resp.Summary.Vulnerability.VulnerablePackages).Should(Equal(10))
				Expect(resp.Detail.Compliance).To(HaveLen(1))
				Expect(resp.Detail.Compliance[0].Reason).Should(Equal("File /etc/pam.d/common-password not found"))
				Expect(resp.Detail.Compliance[0].Compliant).Should(Equal(false))
				Expect(resp.Detail.Compliance[0].Description).Should(Equal("Minimum password length must be 8."))
				Expect(resp.Detail.Compliance[0].PolicyMandated).Should(Equal(false))
				Expect(resp.Detail.Compliance).To(HaveLen(1))
				Expect(resp.Detail.Vulnerability[0].PackageName).Should(Equal("firefox 60.1.0-4.el7_5 has vulnerabilities"))
				Expect(resp.Detail.Vulnerability[0].Vulnerabilities).To(HaveLen(1))
				Expect(resp.Detail.Vulnerability[0].Vulnerabilities[0].URL).Should(Equal("https://access.redhat.com/errata/RHSA-2018:2692"))
				Expect(resp.Detail.Vulnerability[0].Vulnerabilities[0].Cveid).To(HaveLen(5))
				Expect(resp.Detail.Vulnerability[0].Vulnerabilities[0].Cveid[0]).Should(Equal("CVE-2017-16541"))
				Expect(resp.Detail.Vulnerability[0].Vulnerabilities[0].Cveid[1]).Should(Equal("CVE-2018-12376"))
				Expect(resp.Detail.Vulnerability[0].Vulnerabilities[0].Cveid[2]).Should(Equal("CVE-2018-12377"))
				Expect(resp.Detail.Vulnerability[0].Vulnerabilities[0].Cveid[3]).Should(Equal("CVE-2018-12378"))
				Expect(resp.Detail.Vulnerability[0].Vulnerabilities[0].Cveid[4]).Should(Equal("CVE-2018-12379"))
				Expect(resp.Detail.Vulnerability[0].Vulnerabilities[0].Summary).Should(Equal("(RHSA-2018:2692) Critical: firefox security update"))
			})
		})
		Context("When scan image fails", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, fmt.Sprintf("/api/v1/images/%s/vulnerabilities", imageName)),
						ghttp.RespondWith(http.StatusInternalServerError, `Internal Error`),
					),
				)
			})

			It("should return error when image is scanned", func() {
				param := ImageVulnerabilitiesRequest{
					Advisory: true,
					All:      true,
				}
				target := ImageTargetHeader{
					AccountID: "abc",
				}
				resp, err := newImages(server.URL()).ImageVulnerabilities(imageName, param, target)
				Expect(err).To(HaveOccurred())
				Expect(resp).Should(BeNil())
			})
		})
	})

	Describe("DeleteImage", func() {
		Context("When Delete image is completed", func() {

			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, fmt.Sprintf("/api/v1/images/%s", imageName)),
						ghttp.RespondWith(http.StatusOK, imageDelete),
					),
				)
			})

			It("should return delete images results", func() {
				target := ImageTargetHeader{
					AccountID: "abc",
				}
				respptr, err := newImages(server.URL()).DeleteImage(imageName, target)

				Expect(err).NotTo(HaveOccurred())
				resp := *respptr
				Expect(resp.Untagged).Should(Equal("sha256:1002cd429f26a9122df60be5f541a12f09f960b1bd092998f70594e4d8450be8"))
			})
		})
		Context("When delete image fails", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, fmt.Sprintf("/api/v1/images/%s", imageName)),
						ghttp.RespondWith(http.StatusInternalServerError, `Internal Error`),
					),
				)
			})

			It("should return error when image is deleted", func() {
				target := ImageTargetHeader{
					AccountID: "abc",
				}
				resp, err := newImages(server.URL()).DeleteImage(imageName, target)
				Expect(err).To(HaveOccurred())
				Expect(resp).Should(BeNil())
			})
		})
	})
})

func newImages(url string) Images {

	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}
	conf := sess.Config.Copy()
	conf.HTTPClient = ibmcloudHttp.NewHTTPClient(conf)
	conf.Endpoint = &url

	client := client.Client{
		Config:      conf,
		ServiceName: ibmcloud.ContainerRegistryService,
	}
	return newImageAPI(&client)
}
