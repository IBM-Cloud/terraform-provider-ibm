package iamv1

import (
	"log"
	"net/http"

	"github.com/IBM-Cloud/bluemix-go"

	"github.com/IBM-Cloud/bluemix-go/client"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/onsi/gomega/ghttp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ServiceIdRepository", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})

	Describe("List()", func() {
		Context("When there is no service ID", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/serviceids", "boundTo=abc"),
						ghttp.RespondWith(http.StatusOK, `{
							"currentPage": 0,
							"pageSize": 0,
							"items":[]}`),
					),
				)
			})
			It("should return zero service ID", func() {
				ids, err := newTestServiceIDRepo(server.URL()).List("abc")

				Expect(err).ShouldNot(HaveOccurred())
				Expect(ids).Should(BeEmpty())
			})
		})

		Context("When there is only one service ID", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/serviceids", "boundTo=abc"),
						ghttp.RespondWith(http.StatusOK, `{
							"pageSize": 20,
							"items":[{
								"metadata": {
									"iam_id": "iam-ServiceId-06379974-f0e3-4cda-804a-4b6bedb2505b",
									"uuid": "ServiceId-06379974-f0e3-4cda-804a-4b6bedb2505b",
									"crn": "crn:v1:staging:public:iam-identity::a/2bff70b5d2cc4b400814eca0bb730daa::serviceid:ServiceId-06379974-f0e3-4cda-804a-4b6bedb2505b",
									"version": "4-1fd99b9b1f1bd6d013823fe77f1b81b2",
									"createdAt": "2017-10-25T07:29+0000",
									"modifiedAt": "2017-10-25T07:32+0000"
								},
								"entity": {
									"boundTo": "crn:v1:staging:public:::a/2bff70b5d2cc4b400814eca0bb730daa:::",
									"name": "cli-test",
									"description": "service id to cli-test service",
									"uniqueInstanceCrns": []
								}
							}]}`),
					),
				)
			})
			It("should return one service ID", func() {
				ids, err := newTestServiceIDRepo(server.URL()).List("abc")

				Expect(err).ShouldNot(HaveOccurred())
				Expect(ids).Should(HaveLen(1))
				id := ids[0]
				Expect(id.Name).Should(Equal("cli-test"))
				Expect(id.Description).Should(Equal("service id to cli-test service"))
				Expect(id.BoundTo).Should(Equal("crn:v1:staging:public:::a/2bff70b5d2cc4b400814eca0bb730daa:::"))
			})
		})

		Context("When there is pagination", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/serviceids", "boundTo=abc"),
						ghttp.RespondWith(http.StatusOK, `{
							"pageSize": 2,
							"nextPageToken": "eyJkaXIiOiJGIiwiYWNjIjoiMmJmZjcwYjVkMmNjNGI0MDA4MTRlY2EwYmI3MzBkYWEiLCJwZ1MiOjUsInNvcnQiOiIiLCJvZmYiOjV9",
							"items":[{
								"metadata": {
									"iam_id": "iam-ServiceId-06379974-f0e3-4cda-804a-4b6bedb2505b",
									"uuid": "ServiceId-06379974-f0e3-4cda-804a-4b6bedb2505b",
									"crn": "crn:v1:staging:public:iam-identity::a/2bff70b5d2cc4b400814eca0bb730daa::serviceid:ServiceId-06379974-f0e3-4cda-804a-4b6bedb2505b",
									"version": "4-1fd99b9b1f1bd6d013823fe77f1b81b2",
									"createdAt": "2017-10-25T07:29+0000",
									"modifiedAt": "2017-10-25T07:32+0000"
								},
								"entity": {
									"boundTo": "crn:v1:staging:public:::a/2bff70b5d2cc4b400814eca0bb730daa:::",
									"name": "cli-test",
									"description": "service id to cli-test service",
									"uniqueInstanceCrns": []
								}
							},
							{
								"metadata": {
									"iam_id": "iam-ServiceId-5238472e-62fa-437b-9aa7-727f3f846e00",
									"uuid": "ServiceId-5238472e-62fa-437b-9aa7-727f3f846e00",
									"crn": "crn:v1:staging:public:iam-identity::a/2bff70b5d2cc4b400814eca0bb730daa::serviceid:ServiceId-5238472e-62fa-437b-9aa7-727f3f846e00",
									"version": "1-448d54be8b259443ff6169c7a2f83f96",
									"createdAt": "2017-10-25T07:31+0000",
									"modifiedAt": "2017-10-25T07:31+0000"
								},
								"entity": {
									"boundTo": "crn:v1:staging:public:::a/2bff70b5d2cc4b400814eca0bb730daa:::",
									"name": "cli-test-dup",
									"uniqueInstanceCrns": []
								}
							}]}`),
					),
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/serviceids", "boundTo=abc&pagetoken=eyJkaXIiOiJGIiwiYWNjIjoiMmJmZjcwYjVkMmNjNGI0MDA4MTRlY2EwYmI3MzBkYWEiLCJwZ1MiOjUsInNvcnQiOiIiLCJvZmYiOjV9"),
						ghttp.RespondWith(http.StatusOK, `{
							"pageSize": 2,
							"items":[{
								"metadata": {
									"iam_id": "iam-ServiceId-5238472e-62fa-437b-9aa7-727f3f846e00",
									"uuid": "ServiceId-5238472e-62fa-437b-9aa7-727f3f846e00",
									"crn": "crn:v1:staging:public:iam-identity::a/2bff70b5d2cc4b400814eca0bb730daa::serviceid:ServiceId-5238472e-62fa-437b-9aa7-727f3f846e00",
									"version": "1-448d54be8b259443ff6169c7a2f83f96",
									"createdAt": "2017-10-25T07:31+0000",
									"modifiedAt": "2017-10-25T07:31+0000"
								},
								"entity": {
									"boundTo": "crn:v1:staging:public:::a/2bff70b5d2cc4b400814eca0bb730daa:::",
									"name": "cli-test-dup",
									"uniqueInstanceCrns": []
								}
							}]}`),
					),
				)
			})
			It("should return three service IDs", func() {
				ids, err := newTestServiceIDRepo(server.URL()).List("abc")

				Expect(err).ShouldNot(HaveOccurred())
				Expect(ids).Should(HaveLen(3))
				id := ids[0]
				Expect(id.Name).Should(Equal("cli-test"))
				Expect(id.Description).Should(Equal("service id to cli-test service"))
				Expect(id.BoundTo).Should(Equal("crn:v1:staging:public:::a/2bff70b5d2cc4b400814eca0bb730daa:::"))

				id = ids[1]
				Expect(id.Name).Should(Equal("cli-test-dup"))
				Expect(id.Description).Should(Equal(""))
				Expect(id.BoundTo).Should(Equal("crn:v1:staging:public:::a/2bff70b5d2cc4b400814eca0bb730daa:::"))

				id = ids[2]
				Expect(id.Name).Should(Equal("cli-test-dup"))
				Expect(id.Description).Should(Equal(""))
				Expect(id.BoundTo).Should(Equal("crn:v1:staging:public:::a/2bff70b5d2cc4b400814eca0bb730daa:::"))
			})
		})

		Context("When there is error", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/serviceids", "boundTo=abc"),
						ghttp.RespondWith(http.StatusUnauthorized, `{
							"errorCode": "BXNIM0308E",
							"errorMessage": "No authorization header found",
							"context": {
							  "requestId": "2252754335",
							  "requestType": "incoming.ServiceId_List",
							  "startTime": "30.10.2017 06:20:29:504 UTC",
							  "endTime": "30.10.2017 06:20:29:505 UTC",
							  "elapsedTime": "1",
							  "instanceId": "tokenservice/1",
							  "host": "localhost",
							  "threadId": "8d852",
							  "clientIp": "114.255.160.171",
							  "userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.62 Safari/537.36",
							  "locale": "en_US"
							}
						}`),
					),
				)
			})
			It("should return three service IDs", func() {
				ids, err := newTestServiceIDRepo(server.URL()).List("abc")

				Expect(err).Should(HaveOccurred())
				Expect(ids).Should(HaveLen(0))
			})
		})
	})
})

func newTestServiceIDRepo(url string) ServiceIDRepository {
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
	return NewServiceIDRepository(&client)
}
