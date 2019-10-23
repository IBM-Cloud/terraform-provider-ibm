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
	tokenID          = `ae57b394-e39d-595f-ae5d-e46905285df6`
	tokenDescription = "Test Token"
	tokenList        = `{
	"tokens": [
		{
			"_id": "ae57b394-e39d-595f-ae5d-e46905285df6",
			"owner": "abc",
			"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDA0MjMwNTQsImp0aSI6ImFlNTdiMzk0LWUzOWQtNTk1Zi1hZTVkLWU0NjkwNTI4NWRmNiIsImlzcyI6InJlZ2lzdHJ5Lm5nLmJsdWVtaXgubmV0In0.mZBpsw-6RxuvV_Bv7WBk6I6YGJK9f70fEuH5a2cVr3g",
			"secondary_owner": "Test Token"
		}
	]
}`
	token = `{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDA0MjMwNTQsImp0aSI6ImFlNTdiMzk0LWUzOWQtNTk1Zi1hZTVkLWU0NjkwNTI4NWRmNiIsImlzcyI6InJlZ2lzdHJ5Lm5nLmJsdWVtaXgubmV0In0.mZBpsw-6RxuvV_Bv7WBk6I6YGJK9f70fEuH5a2cVr3g"
}`
	tokenError = `{
	"code": "CRG0009E",
	"message": "You are not authorized to access the specified account.",
	"request-id": "72-1540337041.610-99712"
}`
)

var _ = Describe("Tokens", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})
	Describe("GetTokens", func() {
		Context("When get tokens is completed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/api/v1/tokens"),
						ghttp.RespondWith(http.StatusOK, tokenList),
					),
				)
			})

			It("should return get tokens results", func() {
				target := TokenTargetHeader{
					AccountID: "abc",
				}
				respptr, err := newTokens(server.URL()).GetTokens(target)
				Expect(err).NotTo(HaveOccurred())
				Expect(respptr).NotTo(BeNil())
				resp := *respptr
				Expect(resp.Tokens).To(HaveLen(1))
				Expect(resp.Tokens[0].ID).Should(Equal(tokenID))
				Expect(resp.Tokens[0].Owner).Should(Equal("abc"))
				Expect(resp.Tokens[0].Token).Should(Equal("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDA0MjMwNTQsImp0aSI6ImFlNTdiMzk0LWUzOWQtNTk1Zi1hZTVkLWU0NjkwNTI4NWRmNiIsImlzcyI6InJlZ2lzdHJ5Lm5nLmJsdWVtaXgubmV0In0.mZBpsw-6RxuvV_Bv7WBk6I6YGJK9f70fEuH5a2cVr3g"))
				Expect(resp.Tokens[0].Description).Should(Equal("Test Token"))
			})
		})
		Context("When get token fails", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/api/v1/tokens"),
						ghttp.RespondWith(http.StatusInternalServerError, `Internal Error`),
					),
				)
			})

			It("should return error when tokens are retrieved", func() {
				target := TokenTargetHeader{
					AccountID: "abc",
				}
				respptr, err := newTokens(server.URL()).GetTokens(target)
				Expect(err).To(HaveOccurred())
				Expect(respptr).Should(BeNil())
			})
		})
	})

	Describe("GetToken", func() {
		Context("When get token is completed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, fmt.Sprintf("/api/v1/tokens/%s", tokenID)),
						ghttp.RespondWith(http.StatusOK, token),
					),
				)
			})

			It("should return get token results", func() {
				target := TokenTargetHeader{
					AccountID: "abc",
				}
				respptr, err := newTokens(server.URL()).GetToken(tokenID, target)
				Expect(err).NotTo(HaveOccurred())
				Expect(respptr).NotTo(BeNil())
				resp := *respptr
				Expect(resp.ID).Should(Equal(tokenID))
				Expect(resp.Token).Should(Equal("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDA0MjMwNTQsImp0aSI6ImFlNTdiMzk0LWUzOWQtNTk1Zi1hZTVkLWU0NjkwNTI4NWRmNiIsImlzcyI6InJlZ2lzdHJ5Lm5nLmJsdWVtaXgubmV0In0.mZBpsw-6RxuvV_Bv7WBk6I6YGJK9f70fEuH5a2cVr3g"))
			})
		})
		Context("When get token fails", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, fmt.Sprintf("/api/v1/tokens/%s", tokenID)),
						ghttp.RespondWith(http.StatusUnauthorized, tokenError),
					),
				)
			})

			It("should return error when token is goten", func() {
				target := TokenTargetHeader{
					AccountID: "abc",
				}
				resparr, err := newTokens(server.URL()).GetToken(tokenID, target)
				Expect(err).To(HaveOccurred())
				Expect(resparr).Should(BeNil())
			})
		})
	})

	Describe("IssueToken", func() {
		Context("When issue token is completed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/api/v1/tokens"),
						ghttp.RespondWith(http.StatusOK, token),
					),
				)
			})

			It("should return issue token results", func() {
				param := IssueTokenRequest{
					Description: "Test Token",
					Permanent:   false,
					Write:       false,
				}
				target := TokenTargetHeader{
					AccountID: "abc",
				}
				respptr, err := newTokens(server.URL()).IssueToken(param, target)
				Expect(err).NotTo(HaveOccurred())
				Expect(respptr).NotTo(BeNil())
				resp := *respptr
				Expect(resp.ID).Should(Equal(tokenID))
				Expect(resp.Token).Should(Equal("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDA0MjMwNTQsImp0aSI6ImFlNTdiMzk0LWUzOWQtNTk1Zi1hZTVkLWU0NjkwNTI4NWRmNiIsImlzcyI6InJlZ2lzdHJ5Lm5nLmJsdWVtaXgubmV0In0.mZBpsw-6RxuvV_Bv7WBk6I6YGJK9f70fEuH5a2cVr3g"))
			})
		})
		Context("When issue token fails", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/api/v1/tokens"),
						ghttp.RespondWith(http.StatusUnauthorized, tokenError),
					),
				)
			})

			It("should return error when token is issued", func() {
				param := IssueTokenRequest{
					Description: "Test Token",
					Permanent:   false,
					Write:       false,
				}
				target := TokenTargetHeader{
					AccountID: "abc",
				}
				respptr, err := newTokens(server.URL()).IssueToken(param, target)
				Expect(err).To(HaveOccurred())
				Expect(respptr).Should(BeNil())
			})
		})
	})
	Describe("IssueToken", func() {
		Context("When issue token is completed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/api/v1/tokens"),
						ghttp.RespondWith(http.StatusOK, token),
					),
				)
			})

			It("should return issue token results", func() {
				param := IssueTokenRequest{
					Description: "Test Token",
					Permanent:   false,
					Write:       false,
				}
				target := TokenTargetHeader{
					AccountID: "abc",
				}
				respptr, err := newTokens(server.URL()).IssueToken(param, target)
				Expect(err).NotTo(HaveOccurred())
				Expect(respptr).NotTo(BeNil())
				resp := *respptr
				Expect(resp.ID).Should(Equal(tokenID))
				Expect(resp.Token).Should(Equal("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDA0MjMwNTQsImp0aSI6ImFlNTdiMzk0LWUzOWQtNTk1Zi1hZTVkLWU0NjkwNTI4NWRmNiIsImlzcyI6InJlZ2lzdHJ5Lm5nLmJsdWVtaXgubmV0In0.mZBpsw-6RxuvV_Bv7WBk6I6YGJK9f70fEuH5a2cVr3g"))
			})
		})
		Context("When issue token fails", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/api/v1/tokens"),
						ghttp.RespondWith(http.StatusUnauthorized, tokenError),
					),
				)
			})

			It("should return error when token is issued", func() {
				param := IssueTokenRequest{
					Description: "Test Token",
					Permanent:   false,
					Write:       false,
				}
				target := TokenTargetHeader{
					AccountID: "abc",
				}
				resparr, err := newTokens(server.URL()).IssueToken(param, target)
				Expect(err).To(HaveOccurred())
				Expect(resparr).Should(BeNil())
			})
		})
	})
	Describe("DeleteToken", func() {
		Context("When delete token is completed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, fmt.Sprintf("/api/v1/tokens/%s", tokenID)),
						ghttp.RespondWith(http.StatusOK, token),
					),
				)
			})

			It("should delete token results", func() {
				target := TokenTargetHeader{
					AccountID: "abc",
				}
				err := newTokens(server.URL()).DeleteToken(tokenID, target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When dlete token fails", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, fmt.Sprintf("/api/v1/tokens/%s", tokenID)),
						ghttp.RespondWith(http.StatusUnauthorized, tokenError),
					),
				)
			})

			It("should return error when token is deleted", func() {
				target := TokenTargetHeader{
					AccountID: "abc",
				}
				err := newTokens(server.URL()).DeleteToken(tokenID, target)
				Expect(err).To(HaveOccurred())
			})
		})
	})
	Describe("DeleteTokenByDescription", func() {
		Context("When delete token by description is completed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/api/v1/tokens"),
						ghttp.RespondWith(http.StatusOK, tokenDescription),
					),
				)
			})

			It("should delete token by description results", func() {
				target := TokenTargetHeader{
					AccountID: "abc",
				}
				err := newTokens(server.URL()).DeleteTokenByDescription(tokenDescription, target)
				Expect(err).NotTo(HaveOccurred())
			})
		})
		Context("When delete token by description fails", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/api/v1/tokens"),
						ghttp.RespondWith(http.StatusUnauthorized, tokenError),
					),
				)
			})

			It("should return error when token is deleted with description", func() {
				target := TokenTargetHeader{
					AccountID: "abc",
				}
				err := newTokens(server.URL()).DeleteTokenByDescription(tokenID, target)
				Expect(err).To(HaveOccurred())
			})
		})
	})

})

func newTokens(url string) Tokens {

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
	return newTokenAPI(&client)
}
