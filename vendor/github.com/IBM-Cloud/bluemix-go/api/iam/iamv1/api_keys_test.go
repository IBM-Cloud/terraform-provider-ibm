package iamv1

import (
	"log"
	"net/http"

	"github.com/IBM-Cloud/bluemix-go"

	"github.com/IBM-Cloud/bluemix-go/client"
	"github.com/IBM-Cloud/bluemix-go/models"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/onsi/gomega/ghttp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ApiKeys Repository", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})

	Describe("List() method", func() {
		Context("When there is one api key", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/apikeys", "boundTo=abc"),
						ghttp.RespondWith(http.StatusOK, `
						{
							"currentPage": 1,
							"pageSize": 1,
							"items": [{
								"metadata": {
								"uuid": "ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a",
								"crn": "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a",
								"version": "1-892688a64fb68a35ed2aea31b62cbfef",
								"createdAt": "2017-02-19T12:55+0000",
								"modifiedAt": "2017-02-19T12:55+0000"
								},
								"entity": {
								"boundTo": "crn:v:staging:public:iam:::IBMid:user:270004WA4U",
								"name": "test1",
								"description": "my first test api key",
								"format": "APIKEY"
								}
							}]
						}`),
					),
				)
			})

			It("should return all api keys", func() {
				keys, err := newTestAPIKeyRepo(server.URL()).List("abc")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(keys).Should(HaveLen(1))

				key := keys[0]
				Expect(key.UUID).Should(Equal("ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"))
				Expect(key.Crn).Should(Equal("crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"))
				Expect(key.Version).Should(Equal("1-892688a64fb68a35ed2aea31b62cbfef"))
				Expect(key.CreatedAt).Should(Equal("2017-02-19T12:55+0000"))
				Expect(key.ModifiedAt).Should(Equal("2017-02-19T12:55+0000"))
				Expect(key.Name).Should(Equal("test1"))
				Expect(key.BoundTo).Should(Equal("crn:v:staging:public:iam:::IBMid:user:270004WA4U"))
				Expect(key.Description).Should(Equal("my first test api key"))
				Expect(key.Format).Should(Equal("APIKEY"))
			})
		})

		Context("When there are api keys across multiple pages", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/apikeys", "boundTo=abc"),
						ghttp.RespondWith(http.StatusOK, `
						{
							"currentPage": 1,
							"pageSize": 1,
							"nextPageToken": "abc",
							"items": [{
								"metadata": {
								"uuid": "ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a",
								"crn": "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a",
								"version": "1-892688a64fb68a35ed2aea31b62cbfef",
								"createdAt": "2017-02-19T12:55+0000",
								"modifiedAt": "2017-02-19T12:55+0000"
								},
								"entity": {
								"boundTo": "crn:v:staging:public:iam:::IBMid:user:270004WA4U",
								"name": "test1",
								"description": "my first test api key",
								"format": "APIKEY"
								}
							}]
						}`),
					),
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/apikeys", "boundTo=abc&pagetoken=abc"),
						ghttp.RespondWith(http.StatusOK, `
						{
							"currentPage": 2,
							"pageSize": 1,
							"prevPageToken": "def",
							"items": [{
								"metadata": {
								"uuid": "ApiKey-92fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a",
								"crn": "crn:v2:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a",
								"version": "2-892688a64fb68a35ed2aea31b62cbfef",
								"createdAt": "2017-02-20T12:55+0000",
								"modifiedAt": "2017-02-21T12:55+0000"
								},
								"entity": {
								"boundTo": "crn:v2:staging:public:iam:::IBMid:user:270004WA4U",
								"name": "test2",
								"description": "my second test api key",
								"format": "APIKEY"
								}
							}]
						}`),
					),
				)
			})

			It("should return all api keys", func() {
				keys, err := newTestAPIKeyRepo(server.URL()).List("abc")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(keys).Should(HaveLen(2))

				key1 := keys[0]
				Expect(key1.UUID).Should(Equal("ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"))
				Expect(key1.Crn).Should(Equal("crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"))
				Expect(key1.Version).Should(Equal("1-892688a64fb68a35ed2aea31b62cbfef"))
				Expect(key1.CreatedAt).Should(Equal("2017-02-19T12:55+0000"))
				Expect(key1.ModifiedAt).Should(Equal("2017-02-19T12:55+0000"))
				Expect(key1.Name).Should(Equal("test1"))
				Expect(key1.BoundTo).Should(Equal("crn:v:staging:public:iam:::IBMid:user:270004WA4U"))
				Expect(key1.Description).Should(Equal("my first test api key"))
				Expect(key1.Format).Should(Equal("APIKEY"))

				key2 := keys[1]
				Expect(key2.UUID).Should(Equal("ApiKey-92fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"))
				Expect(key2.Crn).Should(Equal("crn:v2:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"))
				Expect(key2.Version).Should(Equal("2-892688a64fb68a35ed2aea31b62cbfef"))
				Expect(key2.CreatedAt).Should(Equal("2017-02-20T12:55+0000"))
				Expect(key2.ModifiedAt).Should(Equal("2017-02-21T12:55+0000"))
				Expect(key2.Name).Should(Equal("test2"))
				Expect(key2.BoundTo).Should(Equal("crn:v2:staging:public:iam:::IBMid:user:270004WA4U"))
				Expect(key2.Description).Should(Equal("my second test api key"))
				Expect(key2.Format).Should(Equal("APIKEY"))
			})
		})

		Context("When there is no api key", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/apikeys", "boundTo=abc"),
						ghttp.RespondWith(http.StatusNotFound, `
						{
							"context": {
								"requestId": "1",
								"requestType": "2",
								"userAgent": "bluemix-cli",
								"clientIp": "3",
								"instanceId": "4",
								"threadId": "5",
								"host": "localhost",
								"startTime": "12345",
								"endTime": "67890",
								"elapsedTime": "123",
								"locale": "en-us"
							},
							"errorCode": "404",
							"errorMessage": "cannot find api keys",
							"errorDetails": "API keys bound to abc cannot be found."
						}`),
					),
				)
			})

			It("should return error", func() {
				keys, err := newTestAPIKeyRepo(server.URL()).List("abc")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(keys).Should(BeEmpty())
			})
		})

		Context("When there is error", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/apikeys", "boundTo=abc"),
						ghttp.RespondWith(http.StatusBadRequest, `
						{
							"context": {
								"requestId": "1",
								"requestType": "2",
								"userAgent": "bluemix-cli",
								"clientIp": "3",
								"instanceId": "4",
								"threadId": "5",
								"host": "localhost",
								"startTime": "12345",
								"endTime": "67890",
								"elapsedTime": "123",
								"locale": "en-us"
							},
							"errorCode": "400",
							"errorMessage": "input error",
							"errorDetails": "boundTo is missing"
						}`),
					),
				)
			})

			It("should return error", func() {
				keys, err := newTestAPIKeyRepo(server.URL()).List("abc")
				Expect(err).Should(HaveOccurred())
				Expect(keys).Should(BeEmpty())
			})
		})
	})

	Describe("FindByName", func() {
		Context("When there is match", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/apikeys", "boundTo=abc"),
						ghttp.RespondWith(http.StatusOK, `
						{
							"currentPage": 1,
							"pageSize": 1,
							"nextPageToken": "abc",
							"items": [{
								"metadata": {
								"uuid": "ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a",
								"crn": "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a",
								"version": "1-892688a64fb68a35ed2aea31b62cbfef",
								"createdAt": "2017-02-19T12:55+0000",
								"modifiedAt": "2017-02-19T12:55+0000"
								},
								"entity": {
								"boundTo": "crn:v:staging:public:iam:::IBMid:user:270004WA4U",
								"name": "test0",
								"description": "my first test api key",
								"format": "APIKEY"
								}
							},{
								"metadata": {
								"uuid": "ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a",
								"crn": "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a",
								"version": "1-892688a64fb68a35ed2aea31b62cbfef",
								"createdAt": "2017-02-19T12:55+0000",
								"modifiedAt": "2017-02-19T12:55+0000"
								},
								"entity": {
								"boundTo": "crn:v:staging:public:iam:::IBMid:user:270004WA4U",
								"name": "test1",
								"description": "my first test api key",
								"format": "APIKEY"
								}
							}]
						}`),
					),
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/apikeys", "boundTo=abc&pagetoken=abc"),
						ghttp.RespondWith(http.StatusOK, `
						{
							"currentPage": 2,
							"pageSize": 1,
							"prevPageToken": "def",
							"items": [{
								"metadata": {
								"uuid": "ApiKey-92fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a",
								"crn": "crn:v2:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a",
								"version": "2-892688a64fb68a35ed2aea31b62cbfef",
								"createdAt": "2017-02-20T12:55+0000",
								"modifiedAt": "2017-02-21T12:55+0000"
								},
								"entity": {
								"boundTo": "crn:v2:staging:public:iam:::IBMid:user:270004WA4U",
								"name": "test2",
								"description": "my second test api key",
								"format": "APIKEY"
								}
							},{
								"metadata": {
								"uuid": "ApiKey-92fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a",
								"crn": "crn:v2:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a",
								"version": "2-892688a64fb68a35ed2aea31b62cbfef",
								"createdAt": "2017-02-20T12:55+0000",
								"modifiedAt": "2017-02-21T12:55+0000"
								},
								"entity": {
								"boundTo": "crn:v2:staging:public:iam:::IBMid:user:270004WA4U",
								"name": "test1",
								"description": "my second test api key",
								"format": "APIKEY"
								}
							}]
						},`),
					),
				)
			})

			It("should return all match", func() {
				keys, err := newTestAPIKeyRepo(server.URL()).FindByName("test1", "abc")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(keys).Should(HaveLen(2))

				key1 := keys[0]
				Expect(key1.UUID).Should(Equal("ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"))
				key2 := keys[1]
				Expect(key2.UUID).Should(Equal("ApiKey-92fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"))
			})
		})

		Context("When there is no match", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/apikeys", "boundTo=abc"),
						ghttp.RespondWith(http.StatusOK, `
						{
							"currentPage": 1,
							"pageSize": 1,
							"pagetoken": "abc",
							"items": [{
								"metadata": {
								"uuid": "ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a",
								"crn": "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a",
								"version": "1-892688a64fb68a35ed2aea31b62cbfef",
								"createdAt": "2017-02-19T12:55+0000",
								"modifiedAt": "2017-02-19T12:55+0000"
								},
								"entity": {
								"boundTo": "crn:v:staging:public:iam:::IBMid:user:270004WA4U",
								"name": "test0",
								"description": "my first test api key",
								"format": "APIKEY"
								}
							},{
								"metadata": {
								"uuid": "ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a",
								"crn": "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a",
								"version": "1-892688a64fb68a35ed2aea31b62cbfef",
								"createdAt": "2017-02-19T12:55+0000",
								"modifiedAt": "2017-02-19T12:55+0000"
								},
								"entity": {
								"boundTo": "crn:v:staging:public:iam:::IBMid:user:270004WA4U",
								"name": "test1",
								"description": "my first test api key",
								"format": "APIKEY"
								}
							}]
						}`),
					),
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/apikeys", "boundTo=abc&pagetoken=abc"),
						ghttp.RespondWith(http.StatusOK, `
						{
							"currentPage": 2,
							"pageSize": 1,
							"prevPageToken": "def",
							"items": [{
								"metadata": {
								"uuid": "ApiKey-92fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a",
								"crn": "crn:v2:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a",
								"version": "2-892688a64fb68a35ed2aea31b62cbfef",
								"createdAt": "2017-02-20T12:55+0000",
								"modifiedAt": "2017-02-21T12:55+0000"
								},
								"entity": {
								"boundTo": "crn:v2:staging:public:iam:::IBMid:user:270004WA4U",
								"name": "test2",
								"description": "my second test api key",
								"format": "APIKEY"
								}
							},{
								"metadata": {
								"uuid": "ApiKey-92fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a",
								"crn": "crn:v2:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a",
								"version": "2-892688a64fb68a35ed2aea31b62cbfef",
								"createdAt": "2017-02-20T12:55+0000",
								"modifiedAt": "2017-02-21T12:55+0000"
								},
								"entity": {
								"boundTo": "crn:v2:staging:public:iam:::IBMid:user:270004WA4U",
								"name": "test1",
								"description": "my second test api key",
								"format": "APIKEY"
								}
							}]
						},`),
					),
				)
			})

			It("should return no result", func() {
				keys, err := newTestAPIKeyRepo(server.URL()).FindByName("test", "abc")
				Expect(err).ShouldNot(HaveOccurred())
				Expect(keys).Should(BeEmpty())
			})
		})
	})

	Describe("Create", func() {
		Context("When creation is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/apikeys"),
						ghttp.VerifyBody([]byte(`{"name":"test1","description":"my first test api key","boundTo":"self"}`)),
						ghttp.RespondWith(http.StatusCreated, `
						{
							"context": {
								"requestId": "1",
								"requestType": "2",
								"userAgent": "bluemix-cli",
								"clientIp": "3",
								"instanceId": "4",
								"threadId": "5",
								"host": "localhost",
								"startTime": "12345",
								"endTime": "67890",
								"elapsedTime": "123",
								"locale": "en-us"
							},
							"metadata": {
							"uuid": "ApiKey-92fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a",
							"crn": "crn:v2:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a",
							"version": "2-892688a64fb68a35ed2aea31b62cbfef",
							"createdAt": "2017-02-20T12:55+0000",
							"modifiedAt": "2017-02-21T12:55+0000"
							},
							"entity": {
							"boundTo": "crn:v2:staging:public:iam:::IBMid:user:270004WA4U",
							"name": "test1",
							"description": "my first test api key",
							"format": "APIKEY"
							}
						}`),
					),
				)
			})

			It("should return key created", func() {
				key, err := newTestAPIKeyRepo(server.URL()).Create(models.APIKey{
					Name:        "test1",
					Description: "my first test api key",
					BoundTo:     "self",
				})
				Expect(err).ShouldNot(HaveOccurred())
				Expect(key).ShouldNot(BeNil())
				Expect(key.Name).Should(Equal("test1"))
				Expect(key.Description).Should(Equal("my first test api key"))
				Expect(key.BoundTo).Should(Equal("crn:v2:staging:public:iam:::IBMid:user:270004WA4U"))
				Expect(key.Format).Should(Equal("APIKEY"))
			})
		})

		Context("When creation fails", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/apikeys"),
						ghttp.VerifyBody([]byte(`{"name":"test1","description":"my first test api key"}`)),
						ghttp.RespondWith(http.StatusBadRequest, `
						{
							"context": {
								"requestId": "1",
								"requestType": "2",
								"userAgent": "bluemix-cli",
								"clientIp": "3",
								"instanceId": "4",
								"threadId": "5",
								"host": "localhost",
								"startTime": "12345",
								"endTime": "67890",
								"elapsedTime": "123",
								"locale": "en-us"
							},
							"errorCode": "400",
							"errorMessage": "input error",
							"errorDetails": "boundTo is missing"
							}
						}`),
					),
				)
			})

			It("should return error", func() {
				key, err := newTestAPIKeyRepo(server.URL()).Create(models.APIKey{
					Name:        "test1",
					Description: "my first test api key",
				})
				Expect(err).Should(HaveOccurred())
				Expect(key).Should(BeNil())
			})
		})
	})

	Describe("Delete", func() {
		Context("When deletion is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/apikeys/abc"),
						ghttp.RespondWith(http.StatusOK, `
						{
							"context": {
								"requestId": "1",
								"requestType": "2",
								"userAgent": "bluemix-cli",
								"clientIp": "3",
								"instanceId": "4",
								"threadId": "5",
								"host": "localhost",
								"startTime": "12345",
								"endTime": "67890",
								"elapsedTime": "123",
								"locale": "en-us"
							}
						}`),
					),
				)
			})

			It("should succeed", func() {
				err := newTestAPIKeyRepo(server.URL()).Delete("abc")
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		Context("When deletion fails", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodDelete, "/apikeys/abc"),
						ghttp.RespondWith(http.StatusNotFound, `
						{
							"context": {
								"requestId": "1",
								"requestType": "2",
								"userAgent": "bluemix-cli",
								"clientIp": "3",
								"instanceId": "4",
								"threadId": "5",
								"host": "localhost",
								"startTime": "12345",
								"endTime": "67890",
								"elapsedTime": "123",
								"locale": "en-us"
							},
							"errorCode": "400",
							"errorMessage": "Cannot find API Key",
							"errorDetails": "UUID 'abc' is not found." 
						}`),
					),
				)
			})

			It("should succeed", func() {
				err := newTestAPIKeyRepo(server.URL()).Delete("abc")
				Expect(err).Should(HaveOccurred())
			})
		})
	})

	Describe("Update", func() {
		Context("When update is succeesful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, "/apikeys/abc"),
						ghttp.VerifyHeaderKV("If-Match", "1-892688a64fb68a35ed2aea31b62cbfef"),
						ghttp.VerifyBody([]byte(`{"name":"test2","description":"my second test api key"}`)),
						ghttp.RespondWith(http.StatusOK, `
						{
							"context": {
								"requestId": "1",
								"requestType": "2",
								"userAgent": "bluemix-cli",
								"clientIp": "3",
								"instanceId": "4",
								"threadId": "5",
								"host": "localhost",
								"startTime": "12345",
								"endTime": "67890",
								"elapsedTime": "123",
								"locale": "en-us"
							},
							"metadata": {
							"uuid": "ApiKey-92fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a",
							"crn": "crn:v2:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a",
							"version": "2-892688a64fb68a35ed2aea31b62cbfef",
							"createdAt": "2017-02-20T12:55+0000",
							"modifiedAt": "2017-02-21T12:55+0000"
							},
							"entity": {
							"boundTo": "crn:v2:staging:public:iam:::IBMid:user:270004WA4U",
							"name": "test2",
							"description": "my second test api key",
							"format": "APIKEY"
							}
							}
						}`),
					),
				)
			})

			It("should return success", func() {
				key, err := newTestAPIKeyRepo(server.URL()).Update("abc", "1-892688a64fb68a35ed2aea31b62cbfef", models.APIKey{
					Name:        "test2",
					Description: "my second test api key",
				})
				Expect(err).ShouldNot(HaveOccurred())
				Expect(key).ShouldNot(BeNil())
				Expect(key.Name).Should(Equal("test2"))
				Expect(key.Description).Should(Equal("my second test api key"))
				Expect(key.UUID).Should(Equal("ApiKey-92fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"))
			})
		})

		Context("When update fails", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPut, "/apikeys/abc"),
						ghttp.VerifyHeaderKV("If-Match", "1-892688a64fb68a35ed2aea31b62cbfef"),
						ghttp.VerifyBody([]byte(`{"name":"test2","description":"my second test api key"}`)),
						ghttp.RespondWith(http.StatusNotFound, `
						{
							"context": {
								"requestId": "1",
								"requestType": "2",
								"userAgent": "bluemix-cli",
								"clientIp": "3",
								"instanceId": "4",
								"threadId": "5",
								"host": "localhost",
								"startTime": "12345",
								"endTime": "67890",
								"elapsedTime": "123",
								"locale": "en-us"
							},
							"errorCode": "404",
							"errorMessage": "Cannot find API Key",
							"errorDetails": "UUID 'abc' is not found." 
							}
						}`),
					),
				)
			})

			It("should return error", func() {
				key, err := newTestAPIKeyRepo(server.URL()).Update("abc", "1-892688a64fb68a35ed2aea31b62cbfef", models.APIKey{
					Name:        "test2",
					Description: "my second test api key",
				})
				Expect(err).Should(HaveOccurred())
				Expect(key).Should(BeNil())
			})
		})
	})
})

func newTestAPIKeyRepo(url string) APIKeyRepository {
	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}
	conf := sess.Config.Copy()
	conf.Endpoint = &url
	client := client.Client{
		Config:      conf,
		ServiceName: bluemix.IAMService,
	}
	return NewAPIKeyRepository(&client)
}
