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

var _ = Describe("Albs", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})
	//Configure
	Describe("Configure", func() {
		Context("When configuring alb is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/alb/albs"),
						ghttp.VerifyJSON(`{"albID":"123","clusterID":"345","name":"test","albType":"public","enable":true,"state":"active","createdDate":"","numOfInstances":"1","resize":false,"albip":"169.0.0.1","zone": "ams03","disableDeployment":false}`),
						ghttp.RespondWith(http.StatusCreated, `{}`),
					),
				)
			})

			It("should configure alb to a cluster", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}
				params := ALBConfig{
					ALBID: "123", ClusterID: "345", Name: "test", ALBType: "public", Enable: true, State: "active", CreatedDate: "", NumOfInstances: "1", Resize: false, ALBIP: "169.0.0.1", Zone: "ams03", DisableDeployment: false,
				}
				err := newAlbs(server.URL()).ConfigureALB("123", params, false, target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When configuring alb is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/alb/albs"),
						ghttp.VerifyJSON(`{"albID":"123","clusterID":"345","name":"test","albType":"public","enable":true,"state":"active","createdDate":"","numOfInstances":"1","resize":false,"albip":"169.0.0.1","zone": "ams03","disableDeployment":false}
`),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to configure alb`),
					),
				)
			})

			It("should return error during configuring alb", func() {
				params := ALBConfig{
					ALBID: "123", ClusterID: "345", Name: "test", ALBType: "public", Enable: true, State: "active", CreatedDate: "", NumOfInstances: "1", Resize: false, ALBIP: "169.0.0.1", Zone: "ams03", DisableDeployment: false,
				}
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}
				err := newAlbs(server.URL()).ConfigureALB("123", params, false, target)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	//ListClusterALBs
	Describe("List cluster albs", func() {
		Context("When read of cluster albs is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/alb/clusters/test"),
						ghttp.RespondWith(http.StatusOK, `
						  {
						    "alb": [
						      {
						        "albID": "string",
						        "albType": "string",
						        "albip": "string",
						        "clusterID": "string",
						        "createdDate": "string",
						        "enable": true,
						        "name": "string",
						        "numOfInstances": "string",
						        "resize": true,
						        "state": "string",
						        "zone": "string",
						        "disableDeployment":false
						      }
						    ],
						    "dataCenter": "string",
						    "id": "string",
						    "ingressHostname": "string",
						    "ingressSecretName": "string",
						    "isPaid": true,
						    "region": "string"
						  }
						`),
					),
				)
			})

			It("should return cluster albs list", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}
				albs, err := newAlbs(server.URL()).ListClusterALBs("test", target)
				Expect(albs).ShouldNot(BeNil())
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When read of cluster albs is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/alb/clusters/test"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve albs`),
					),
				)
			})

			It("should return error when cluster albs are retrieved", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}
				albs, err := newAlbs(server.URL()).ListClusterALBs("test", target)
				Expect(err).To(HaveOccurred())
				Expect(albs).Should(BeNil())
			})
		})
	})
	//GetAlb
	Describe("Get cluster alb", func() {
		Context("When read of cluster alb is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/alb/albs/testAlb"),
						ghttp.RespondWith(http.StatusOK, `{"albID":"123","clusterID":"345","name":"test","albType":"public","enable":true,"state":"active","createdDate":"","numOfInstances":"1","resize":false,"albip":"169.0.0.1","zone": "ams03","disableDeployment":false}`),
					),
				)
			})

			It("should return albs", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}
				alb, err := newAlbs(server.URL()).GetALB("testAlb", target)
				Expect(alb).ShouldNot(BeNil())
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When read of cluster alb is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/alb/albs/testAlb"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve alb.`),
					),
				)
			})

			It("should return error when alb are retrieved", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}
				_, err := newAlbs(server.URL()).GetALB("testAlb", target)
				Expect(err).To(HaveOccurred())
			})
		})
	})
	//Deploy Alb cert
	Describe("Deploy Alb cert", func() {
		Context("When deploying alb cert is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/alb/albsecrets"),
						ghttp.VerifyJSON(`{"secretName":"test","clusterID":"345","domainName":"testDomain","cloudCertInstanceID":"456","clusterCrn":"crn::cluster","certCrn":"crn::cert","issuerName":"testissue","expiresOn":"","state":"active"}
`),
						ghttp.RespondWith(http.StatusCreated, `{}`),
					),
				)
			})

			It("should deploy alb to a cluster", func() {
				params := ALBSecretConfig{
					SecretName: "test", ClusterID: "345", DomainName: "testDomain", CloudCertInstanceID: "456", ClusterCrn: "crn::cluster", CertCrn: "crn::cert", IssuerName: "testissue", ExpiresOn: "", State: "active",
				}
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}
				err := newAlbs(server.URL()).DeployALBCert(params, target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When deploying alb cert is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v1/alb/albsecrets"),
						ghttp.VerifyJSON(`{"secretName":"test","clusterID":"345","domainName":"testDomain","cloudCertInstanceID":"456","clusterCrn":"crn::cluster","certCrn":"crn::cert","issuerName":"testissue","expiresOn":"","state":"active"}
`),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to deploy alb cert`),
					),
				)
			})

			It("should return error during deploying alb cert", func() {
				params := ALBSecretConfig{
					SecretName: "test", ClusterID: "345", DomainName: "testDomain", CloudCertInstanceID: "456", ClusterCrn: "crn::cluster", CertCrn: "crn::cert", IssuerName: "testissue", ExpiresOn: "", State: "active",
				}
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}
				err := newAlbs(server.URL()).DeployALBCert(params, target)
				Expect(err).To(HaveOccurred())
			})
		})
	})
	//UpdateALBCert
	Describe("Update Alb cert", func() {
		Context("When updating alb cert is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, "/v1/alb/albsecrets"),
						ghttp.VerifyJSON(`{"secretName":"test","clusterID":"345","domainName":"testDomain","cloudCertInstanceID":"456","clusterCrn":"crn::cluster","certCrn":"crn::cert","issuerName":"testissue","expiresOn":"","state":"active"}
`),
						ghttp.RespondWith(http.StatusNoContent, `{}`),
					),
				)
			})

			It("should deploy alb to a cluster", func() {
				params := ALBSecretConfig{
					SecretName: "test", ClusterID: "345", DomainName: "testDomain", CloudCertInstanceID: "456", ClusterCrn: "crn::cluster", CertCrn: "crn::cert", IssuerName: "testissue", ExpiresOn: "", State: "active",
				}
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}
				err := newAlbs(server.URL()).UpdateALBCert(params, target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When deploying alb cert is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, "/v1/alb/albsecrets"),
						ghttp.VerifyJSON(`{"secretName":"test","clusterID":"345","domainName":"testDomain","cloudCertInstanceID":"456","clusterCrn":"crn::cluster","certCrn":"crn::cert","issuerName":"testissue","expiresOn":"","state":"active"}
`),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to update alb cert`),
					),
				)
			})

			It("should return error during deploying alb cert", func() {
				params := ALBSecretConfig{
					SecretName: "test", ClusterID: "345", DomainName: "testDomain", CloudCertInstanceID: "456", ClusterCrn: "crn::cluster", CertCrn: "crn::cert", IssuerName: "testissue", ExpiresOn: "", State: "active",
				}
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}
				err := newAlbs(server.URL()).UpdateALBCert(params, target)
				Expect(err).To(HaveOccurred())
			})
		})
	})
	//ListALBCerts
	Describe("Get cluster alb certs", func() {
		Context("When read of cluster alb certs is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/alb/clusters/test/albsecrets"),
						ghttp.RespondWith(http.StatusOK, `{"id":"123","region":"eu-de","dataCenter":"ams03","isPaid":true,"albSecrets":[{"secretName":"test","clusterID":"test","domainName":"testDomain","cloudCertInstanceID":"456","clusterCrn":"crn::cluster","certCrn":"crn::cert","issuerName":"testissue","expiresOn":"string","state":"active"}]}`),
					),
				)
			})

			It("should return cluster alb certs list", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}
				albCerts, err := newAlbs(server.URL()).ListALBCerts("test", target)
				Expect(albCerts).ShouldNot(BeNil())
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When read of cluster alb certs is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/alb/clusters/test/albsecrets"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve alb certs`),
					),
				)
			})

			It("should return error when cluster alb certss are retrieved", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}
				albs, err := newAlbs(server.URL()).ListALBCerts("test", target)
				Expect(err).To(HaveOccurred())
				Expect(albs).Should(BeNil())
			})
		})
	})
	//GetClusterALBCertBySecretName
	Describe("Get cluster alb cert", func() {
		Context("When read of cluster alb cert is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/alb/clusters/test/albsecrets"),
						ghttp.RespondWith(http.StatusOK, `{"secretName":"test","clusterID":"345","domainName":"testDomain","cloudCertInstanceID":"456","clusterCrn":"crn::cluster","certCrn":"crn::cert","issuerName":"testissue","expiresOn":"","state":"active"}`),
					),
				)
			})

			It("should return albs", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}
				alb, err := newAlbs(server.URL()).GetClusterALBCertBySecretName("test", "testSecret", target)
				Expect(alb).ShouldNot(BeNil())
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When read of cluster alb cert is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/alb/clusters/test/albsecrets"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve alb cert.`),
					),
				)
			})

			It("should return error when alb cert are retrieved", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}
				_, err := newAlbs(server.URL()).GetClusterALBCertBySecretName("test", "testSecret", target)
				Expect(err).To(HaveOccurred())
			})
		})
	})
	//GetClusterALBCertByCertCRN
	Describe("Get cluster alb cert", func() {
		Context("When read of cluster alb cert is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/alb/clusters/test/albsecrets"),
						ghttp.RespondWith(http.StatusOK, `{"secretName":"test","clusterID":"345","domainName":"testDomain","cloudCertInstanceID":"456","clusterCrn":"crn::cluster","certCrn":"crn::cert","issuerName":"testissue","expiresOn":"","state":"active"}`),
					),
				)
			})

			It("should return albs", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}
				alb, err := newAlbs(server.URL()).GetClusterALBCertByCertCRN("test", "testCert", target)
				Expect(alb).ShouldNot(BeNil())
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When read of cluster alb cert is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v1/alb/clusters/test/albsecrets"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve alb cert.`),
					),
				)
			})

			It("should return error when alb cert are retrieved", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}
				_, err := newAlbs(server.URL()).GetClusterALBCertByCertCRN("test", "testCert", target)
				Expect(err).To(HaveOccurred())
			})
		})
	})
	//RemoveALB
	Describe("Delete", func() {
		Context("When delete of alb is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/alb/albs/test"),
						ghttp.RespondWith(http.StatusOK, `{							
						}`),
					),
				)
			})

			It("should delete alb", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}

				err := newAlbs(server.URL()).RemoveALB("test", target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When alb delete is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/alb/albs/test"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to delete alb`),
					),
				)
			})

			It("should return error alb delete", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}
				err := newAlbs(server.URL()).RemoveALB("test", target)
				Expect(err).To(HaveOccurred())
			})
		})
	})
	//RemoveALBCertBySecretName
	Describe("Delete", func() {
		Context("When delete of alb cert is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/alb/clusters/mycluster/albsecrets"),
						ghttp.RespondWith(http.StatusOK, `{							
						}`),
					),
				)
			})

			It("should delete alb cert", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}
				err := newAlbs(server.URL()).RemoveALBCertBySecretName("mycluster", "test", target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When alb cert delete is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/alb/clusters/mycluster/albsecrets"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to delete alb`),
					),
				)
			})

			It("should return error alb cert delete", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}
				err := newAlbs(server.URL()).RemoveALBCertBySecretName("mycluster", "test", target)
				Expect(err).To(HaveOccurred())
			})
		})
	})
	//RemoveALBCertByCertCRN
	Describe("Delete", func() {
		Context("When delete of alb cert is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/alb/clusters/mycluster/albsecrets"),
						ghttp.RespondWith(http.StatusOK, `{							
						}`),
					),
				)
			})

			It("should delete alb cert", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}
				err := newAlbs(server.URL()).RemoveALBCertByCertCRN("mycluster", "test", target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When alb cert delete is failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/v1/alb/clusters/mycluster/albsecrets"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to delete alb`),
					),
				)
			})

			It("should return error alb cert delete", func() {
				target := ClusterTargetHeader{
					OrgID:     "abc",
					SpaceID:   "def",
					AccountID: "ghi",
					Region:    "eu-de",
				}
				err := newAlbs(server.URL()).RemoveALBCertByCertCRN("mycluster", "test", target)
				Expect(err).To(HaveOccurred())
			})
		})
	})

})

func newAlbs(url string) Albs {

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
	return newAlbAPI(&client)
}
