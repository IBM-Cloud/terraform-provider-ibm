package registryv1

import (
	"fmt"
	"log"
	"net/http"

	ibmcloud "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/client"
	ibmcloudHttp "github.com/IBM-Cloud/bluemix-go/http"
	"github.com/IBM-Cloud/bluemix-go/session"

	"github.com/onsi/gomega/ghttp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	namespaceName = `gpfs`
	namespaceList = `[
	"gpfs",
	"devops_insights_dnt_dev",
	"devops_insights_dnt_staging",
	"bkuschel",
	"othomann",
	"bog-jx"
]`
	addNamespace = `{
	"namespace": "gpfs"
}`
	namespaceForbidden = `{
	"code": "CRG0020E",
	"message": "You are not authorized to access the specified resource.",
	"request-id": "4185-1540334590.738-17873666"
}`
)

var _ = Describe("Namespaces", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})
	Describe("GetNamespaces", func() {
		Context("When get namespaces is completed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/api/v1/namespaces"),
						ghttp.RespondWith(http.StatusOK, namespaceList),
					),
				)
			})

			It("should return get namespaces results", func() {
				target := NamespaceTargetHeader{
					AccountID: "abc",
				}
				resparr, err := newNamespaces(server.URL()).GetNamespaces(target)
				Expect(err).NotTo(HaveOccurred())
				Expect(resparr).NotTo(BeNil())
				Expect(resparr).To(HaveLen(6))
				Expect(resparr[0]).Should(Equal("gpfs"))
				Expect(resparr[1]).Should(Equal("devops_insights_dnt_dev"))
				Expect(resparr[2]).Should(Equal("devops_insights_dnt_staging"))
				Expect(resparr[3]).Should(Equal("bkuschel"))
				Expect(resparr[4]).Should(Equal("othomann"))
				Expect(resparr[5]).Should(Equal("bog-jx"))
			})
		})
		Context("When get namespace fails", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/api/v1/namespaces"),
						ghttp.RespondWith(http.StatusInternalServerError, `Internal Error`),
					),
				)
			})

			It("should return error when namespaces are retrieved", func() {
				target := NamespaceTargetHeader{
					AccountID: "abc",
				}
				resparr, err := newNamespaces(server.URL()).GetNamespaces(target)
				Expect(err).To(HaveOccurred())
				Expect(resparr).Should(BeNil())
			})
		})
	})

	Describe("AddNamespace", func() {
		Context("When add namespace is completed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, fmt.Sprintf("/api/v1/namespaces/%s", namespaceName)),
						ghttp.RespondWith(http.StatusOK, addNamespace),
					),
				)
			})

			It("should return add namespace results", func() {
				target := NamespaceTargetHeader{
					AccountID: "abc",
				}
				respptr, err := newNamespaces(server.URL()).AddNamespace(namespaceName, target)
				Expect(err).NotTo(HaveOccurred())
				Expect(respptr).NotTo(BeNil())
				resp := *respptr
				Expect(resp.Namespace).Should(Equal("gpfs"))
			})
		})
		Context("When add namespace fails", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, fmt.Sprintf("/api/v1/namespaces/%s", namespaceName)),
						ghttp.RespondWith(http.StatusForbidden, namespaceForbidden),
					),
				)
			})

			It("should return error when namepsaces is added", func() {
				target := NamespaceTargetHeader{
					AccountID: "abc",
				}
				resparr, err := newNamespaces(server.URL()).AddNamespace(namespaceName, target)
				Expect(err).To(HaveOccurred())
				Expect(resparr).Should(BeNil())
			})
		})
	})

	Describe("DeleteNamespace", func() {
		Context("When Delete namespace is completed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, fmt.Sprintf("/api/v1/namespaces/%s", namespaceName)),
						ghttp.RespondWith(http.StatusNoContent, nil),
					),
				)
			})

			It("should return delete namespace results", func() {
				target := NamespaceTargetHeader{
					AccountID: "abc",
				}
				err := newNamespaces(server.URL()).DeleteNamespace(namespaceName, target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When delete namespace fails", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, fmt.Sprintf("/api/v1/namespaces/%s", namespaceName)),
						ghttp.RespondWith(http.StatusForbidden, namespaceForbidden),
					),
				)
			})

			It("should return error when namespace is deleted", func() {
				target := NamespaceTargetHeader{
					AccountID: "abc",
				}
				err := newNamespaces(server.URL()).DeleteNamespace(namespaceName, target)
				Expect(err).To(HaveOccurred())
			})
		})
	})
})

func newNamespaces(url string) Namespaces {

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
	return newNamespaceAPI(&client)
}
