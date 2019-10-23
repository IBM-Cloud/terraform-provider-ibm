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

var _ = Describe("connections", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})
	Describe("Get Postgres", func() {
		Context("When get is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v4/ibm/deployments/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/users/anyone/connections"),
						ghttp.RespondWith(http.StatusOK, `
                           {
                              "connection": {
                                "postgres": {
                                  "composed": [
                                    "postgres://admin:$PASSWORD@1b8f53db-fc2d-4e24-8470-f82b15c71717.8f7bfd8f3faa4218aec56e069eb46187.databases.appdomain.cloud:32121/ibmclouddb?sslmode=verify-full"
                                  ],
                                  "type": "uri",
                                  "scheme": "postgres",
                                  "hosts": [
                                    {
                                      "hostname": "1b8f53db-fc2d-4e24-8470-f82b15c71717.8f7bfd8f3faa4218aec56e069eb46187.databases.appdomain.cloud",
                                      "port": 32121
                                    }
                                  ],
                                  "path": "/ibmclouddb",
                                  "query_options": null,
                                  "authentication": {
                                    "method": "direct",
                                    "username": "admin",
                                    "password": "$PASSWORD"
                                  },
                                  "database": "ibmclouddb",
                                  "certificate": {
                                    "name": "0b22f14b-7ba2-11e8-b8e9-568642342d40",
                                    "certificate_base64": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURE..."
                                  }
                                },
                                "cli": {
                                  "composed": [
                                    "PGPASSWORD=$PASSWORD PGSSLROOTCERT=0b22f14b-7ba2-11e8-b8e9-568642342d40 psql 'host=1b8f53db-fc2d-4e24-8470-f82b15c71717.8f7bfd8f3faa4218aec56e069eb46187.databases.appdomain.cloud port=32121 dbname=ibmclouddb user=admin sslmode=verify-full"
                                  ],
                                  "type": "cli",
                                  "environment": {
                                    "PGPASSWORD": "$PASSWORD"
                                  },
                                  "bin": "psql",
                                  "arguments": [
                                    [
                                      "host=1b8f53db-fc2d-4e24-8470-f82b15c71717.8f7bfd8f3faa4218aec56e069eb46187.databases.appdomain.cloud port=32121 dbname=ibmclouddb user=admin sslmode=verify-full"
                                    ]
                                  ],
                                  "certificate": {
                                    "name": "0b22f14b-7ba2-11e8-b8e9-568642342d40",
                                    "certificate_base64": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURE..."
                                  }
                                }
                              }
                            }
                        `),
					),
				)
			})

			It("should return connection", func() {
				target := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				userId := "anyone"
				connection, err := newConnection(server.URL()).GetConnection(target, userId)
				Expect(err).NotTo(HaveOccurred())
				Expect(connection).ShouldNot(BeNil())
				Expect(connection.Postgres).ShouldNot(BeNil())
				// Expect(connection.Postgres.Composed[0]).Should(Equal("postgres://admin:$PASSWORD@1b8f53db-fc2d-4e24-8470-f82b15c71717.8f7bfd8f3faa4218aec56e069eb46187.databases.appdomain.cloud:32121/ibmclouddb?sslmode=verify-full"))
				// Expect(connection.Postgres.Type).Should(Equal("uri"))
				// Expect(connection.Postgres.Scheme).Should(Equal("postgres"))
				// Expect(connection.Postgres.Path).Should(Equal("/ibmclouddb"))
				// Expect(connection.Postgres.Hosts[0].HostName).Should(Equal("1b8f53db-fc2d-4e24-8470-f82b15c71717.8f7bfd8f3faa4218aec56e069eb46187.databases.appdomain.cloud"))
				// Expect(connection.Postgres.Hosts[0].Port).Should(Equal(32121))
				// //Expect(connection.Postgres.QueryOptions).Should(Equal(""))
				// Expect(connection.Postgres.Authentication.Method).Should(Equal("direct"))
				// Expect(connection.Postgres.Authentication.UserName).Should(Equal("admin"))
				// Expect(connection.Postgres.Authentication.Password).Should(Equal("$PASSWORD"))
				// Expect(connection.Postgres.Certificate.Name).Should(Equal("0b22f14b-7ba2-11e8-b8e9-568642342d40"))
				// Expect(connection.Postgres.Certificate.CertificateBase64).Should(Equal("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURE..."))
				// Expect(connection.Postgres.Database).Should(BeEquivalentTo("ibmclouddb"))
				// Expect(connection.Cli.Composed[0]).Should(Equal("PGPASSWORD=$PASSWORD PGSSLROOTCERT=0b22f14b-7ba2-11e8-b8e9-568642342d40 psql 'host=1b8f53db-fc2d-4e24-8470-f82b15c71717.8f7bfd8f3faa4218aec56e069eb46187.databases.appdomain.cloud port=32121 dbname=ibmclouddb user=admin sslmode=verify-full"))
				// Expect(connection.Cli.Type).Should(Equal("cli"))
				// Expect(connection.Cli.Bin).Should(Equal("psql"))
				// Expect(connection.Cli.Arguments[0][0]).Should(Equal("host=1b8f53db-fc2d-4e24-8470-f82b15c71717.8f7bfd8f3faa4218aec56e069eb46187.databases.appdomain.cloud port=32121 dbname=ibmclouddb user=admin sslmode=verify-full"))
				// Expect(connection.Cli.Certificate.Name).Should(Equal("0b22f14b-7ba2-11e8-b8e9-568642342d40"))
				// Expect(connection.Cli.Certificate.CertificateBase64).Should(Equal("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURE..."))
			})
		})
		Context("When get is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v4/ibm/deployments/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/users/anyone/connections"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to get connection`),
					),
				)
			})

			It("should return error during get connection", func() {
				target := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				userId := "anyone"
				_, err := newConnection(server.URL()).GetConnection(target, userId)
				Expect(err).To(HaveOccurred())
				//Expect(connection).Should(Equal(""))
			})
		})
	})

	Describe("Get Rediss", func() {
		Context("When get is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v4/ibm/deployments/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/users/anyone/connections"),
						ghttp.RespondWith(http.StatusOK,
							//  `
							//   {"connection":
							//    {
							//   "rediss":{"composed":["rediss://admin:$PASSWORD@9f7c2c13-5061-46ca-8153-f7d6c0baba7b.974550db55eb4ec0983f023940bf637f.databases.appdomain.cloud:30348/0"],
							//   "type":"uri",
							//   "scheme":"rediss",
							//   "hosts":[{"hostname":"9f7c2c13-5061-46ca-8153-f7d6c0baba7b.974550db55eb4ec0983f023940bf637f.databases.appdomain.cloud",
							//   "port":30348,
							//   "protocol":""}],
							//   "path":"/0",
							//   "query_options":{},
							//   "authentication":{"method":"direct","username":"admin","password":"$PASSWORD"},
							//   "certificate":{"name":"c5f07736-d94c-11e8-a2e9-62ec2ed68f84",
							//   "certificate_base64":"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tL"},
							//   "database":0},
							//   "cli":
							//     {
							//       "composed":["REDIS_CERTFILE=c5f07736-d94c-11e8-a2e9-62ec2ed68f84 redli -u rediss://admin:$PASSWORD@9f7c2c13-5061-46ca-8153-f7d6c0baba7b.974550db55eb4ec0983f023940bf637f.databases.appdomain.cloud:30348/0"],
							//       "type":"cli",
							//       "environment":{"REDIS_CERTFILE":"c5f07736-d94c-11e8-a2e9-62ec2ed68f84"},
							//       "bin":"redli",
							//       "arguments":[["-u","rediss://admin:$PASSWORD@9f7c2c13-5061-46ca-8153-f7d6c0baba7b.974550db55eb4ec0983f023940bf637f.databases.appdomain.cloud:30348/0"]],
							//       "certificate":{"name":"c5f07736-d94c-11e8-a2e9-62ec2ed68f84",
							//       "certificate_base64":"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tL"}
							//     }
							//   }
							// }
							// `
							`
                          {"connection":{"rediss":{"composed":["rediss://admin:$PASSWORD@9f7c2c13-5061-46ca-8153-f7d6c0baba7b.974550db55eb4ec0983f023940bf637f.databases.appdomain.cloud:30348/0"],"type":"uri","scheme":"rediss","hosts":[{"hostname":"9f7c2c13-5061-46ca-8153-f7d6c0baba7b.974550db55eb4ec0983f023940bf637f.databases.appdomain.cloud","port":30348,"protocol":""}],"path":"/0","query_options":{},"authentication":{"method":"direct","username":"admin","password":"[PRIVATE DATA HIDDEN]"},"certificate":{"name":"c5f07736-d94c-11e8-a2e9-62ec2ed68f84","certificate_base64":"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUREekNDQWZlZ0F3SUJBZ0lKQU5FSDU4eTIva3pITUEwR0NTcUdTSWIzRFFFQkN3VUFNQjR4SERBYUJnTlYKQkFNTUUwbENUU0JEYkc5MVpDQkVZWFJoWW1GelpYTXdIaGNOTVRnd05qSTFNVFF5T1RBd1doY05Namd3TmpJeQpNVFF5T1RBd1dqQWVNUnd3R2dZRFZRUUREQk5KUWswZ1EyeHZkV1FnUkdGMFlXSmhjMlZ6TUlJQklqQU5CZ2txCmhraUc5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0NBUUVBOGxwYVFHemNGZEdxZU1sbXFqZmZNUHBJUWhxcGQ4cUoKUHIzYklrclhKYlRjSko5dUlja1NVY0NqdzRaL3JTZzhublQxM1NDY09sKzF0bys3a2RNaVU4cU9XS2ljZVlaNQp5K3laWWZDa0dhaVpWZmF6UUJtNDV6QnRGV3YrQUIvOGhmQ1RkTkY3Vlk0c3BhQTNvQkUyYVM3T0FOTlNSWlNLCnB3eTI0SVVnVWNJTEpXK21jdlc4MFZ4K0dYUmZEOVl0dDZQUkpnQmhZdVVCcGd6dm5nbUNNR0JuK2wyS05pU2YKd2VvdllEQ0Q2Vm5nbDIrNlc5UUZBRnRXWFdnRjNpRFFENW5sL240bXJpcE1TWDZVRy9uNjY1N3U3VERkZ2t2QQoxZUtJMkZMellLcG9LQmU1cmNuck03bkhnTmMvbkNkRXM1SmVjSGIxZEh2MVFmUG02cHpJeHdJREFRQUJvMUF3ClRqQWRCZ05WSFE0RUZnUVVLMytYWm8xd3lLcytERW9ZWGJIcnV3U3BYamd3SHdZRFZSMGpCQmd3Rm9BVUszK1gKWm8xd3lLcytERW9ZWGJIcnV3U3BYamd3REFZRFZSMFRCQVV3QXdFQi96QU5CZ2txaGtpRzl3MEJBUXNGQUFPQwpBUUVBSmY1ZHZselVwcWFpeDI2cUpFdXFGRzBJUDU3UVFJNVRDUko2WHQvc3VwUkhvNjNlRHZLdzh6Ujd0bFdRCmxWNVAwTjJ4d3VTbDlacUFKdDcvay8zWmVCK25Zd1BveU8zS3ZLdkFUdW5SdmxQQm40RldWWGVhUHNHKzdmaFMKcXNlam1reW9uWXc3N0hSekdPekpINFpnOFVONm1mcGJhV1NzeWFFeHZxa25DcDlTb1RRUDNENjdBeldxYjF6WQpkb3FxZ0dJWjJueENrcDUvRlh4Ri9UTWI1NXZ0ZVRRd2ZnQnk2MGpWVmtiRjdlVk9XQ3YwS2FOSFBGNWhycWJOCmkrM1hqSjcvcGVGM3hNdlRNb3kzNURjVDNFMlplU1Zqb3VaczE1Tzkwa0kzazJkYVMyT0hKQUJXMHZTajRuTHoKK1BRenAvQjljUW1PTzhkQ2UwNDlRM29hVUE9PQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCgo="},"database":0},"cli":{"composed":["REDIS_CERTFILE=c5f07736-d94c-11e8-a2e9-62ec2ed68f84 redli -u rediss://admin:$PASSWORD@9f7c2c13-5061-46ca-8153-f7d6c0baba7b.974550db55eb4ec0983f023940bf637f.databases.appdomain.cloud:30348/0"],"type":"cli","environment":{"REDIS_CERTFILE":"c5f07736-d94c-11e8-a2e9-62ec2ed68f84"},"bin":"redli","arguments":[["-u","rediss://admin:$PASSWORD@9f7c2c13-5061-46ca-8153-f7d6c0baba7b.974550db55eb4ec0983f023940bf637f.databases.appdomain.cloud:30348/0"]],"certificate":{"name":"c5f07736-d94c-11e8-a2e9-62ec2ed68f84","certificate_base64":"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUREekNDQWZlZ0F3SUJBZ0lKQU5FSDU4eTIva3pITUEwR0NTcUdTSWIzRFFFQkN3VUFNQjR4SERBYUJnTlYKQkFNTUUwbENUU0JEYkc5MVpDQkVZWFJoWW1GelpYTXdIaGNOTVRnd05qSTFNVFF5T1RBd1doY05Namd3TmpJeQpNVFF5T1RBd1dqQWVNUnd3R2dZRFZRUUREQk5KUWswZ1EyeHZkV1FnUkdGMFlXSmhjMlZ6TUlJQklqQU5CZ2txCmhraUc5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0NBUUVBOGxwYVFHemNGZEdxZU1sbXFqZmZNUHBJUWhxcGQ4cUoKUHIzYklrclhKYlRjSko5dUlja1NVY0NqdzRaL3JTZzhublQxM1NDY09sKzF0bys3a2RNaVU4cU9XS2ljZVlaNQp5K3laWWZDa0dhaVpWZmF6UUJtNDV6QnRGV3YrQUIvOGhmQ1RkTkY3Vlk0c3BhQTNvQkUyYVM3T0FOTlNSWlNLCnB3eTI0SVVnVWNJTEpXK21jdlc4MFZ4K0dYUmZEOVl0dDZQUkpnQmhZdVVCcGd6dm5nbUNNR0JuK2wyS05pU2YKd2VvdllEQ0Q2Vm5nbDIrNlc5UUZBRnRXWFdnRjNpRFFENW5sL240bXJpcE1TWDZVRy9uNjY1N3U3VERkZ2t2QQoxZUtJMkZMellLcG9LQmU1cmNuck03bkhnTmMvbkNkRXM1SmVjSGIxZEh2MVFmUG02cHpJeHdJREFRQUJvMUF3ClRqQWRCZ05WSFE0RUZnUVVLMytYWm8xd3lLcytERW9ZWGJIcnV3U3BYamd3SHdZRFZSMGpCQmd3Rm9BVUszK1gKWm8xd3lLcytERW9ZWGJIcnV3U3BYamd3REFZRFZSMFRCQVV3QXdFQi96QU5CZ2txaGtpRzl3MEJBUXNGQUFPQwpBUUVBSmY1ZHZselVwcWFpeDI2cUpFdXFGRzBJUDU3UVFJNVRDUko2WHQvc3VwUkhvNjNlRHZLdzh6Ujd0bFdRCmxWNVAwTjJ4d3VTbDlacUFKdDcvay8zWmVCK25Zd1BveU8zS3ZLdkFUdW5SdmxQQm40RldWWGVhUHNHKzdmaFMKcXNlam1reW9uWXc3N0hSekdPekpINFpnOFVONm1mcGJhV1NzeWFFeHZxa25DcDlTb1RRUDNENjdBeldxYjF6WQpkb3FxZ0dJWjJueENrcDUvRlh4Ri9UTWI1NXZ0ZVRRd2ZnQnk2MGpWVmtiRjdlVk9XQ3YwS2FOSFBGNWhycWJOCmkrM1hqSjcvcGVGM3hNdlRNb3kzNURjVDNFMlplU1Zqb3VaczE1Tzkwa0kzazJkYVMyT0hKQUJXMHZTajRuTHoKK1BRenAvQjljUW1PTzhkQ2UwNDlRM29hVUE9PQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCgo="}}}}
                        `),
					),
				)
			})

			It("should return connection", func() {
				target := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				userId := "anyone"
				connection, err := newConnection(server.URL()).GetConnection(target, userId)
				Expect(err).NotTo(HaveOccurred())
				Expect(connection).ShouldNot(BeNil())
				Expect(connection.Rediss).ShouldNot(BeNil())
				Expect(connection.Rediss.Composed[0]).Should(Equal("rediss://admin:$PASSWORD@9f7c2c13-5061-46ca-8153-f7d6c0baba7b.974550db55eb4ec0983f023940bf637f.databases.appdomain.cloud:30348/0"))
				// Expect(connection.Rediss.Type).Should(Equal("uri"))
				// Expect(connection.Rediss.Scheme).Should(Equal("rediss"))
				// Expect(connection.Rediss.Path).Should(Equal("/0"))
				// Expect(connection.Rediss.Hosts[0].HostName).Should(Equal("9f7c2c13-5061-46ca-8153-f7d6c0baba7b.974550db55eb4ec0983f023940bf637f.databases.appdomain.cloud"))
				// Expect(connection.Rediss.Hosts[0].Port).Should(Equal(30348))
				// //Expect(connection.Rediss.QueryOptions).Should(Equal(""))
				// Expect(connection.Rediss.Authentication.Method).Should(Equal("direct"))
				// Expect(connection.Rediss.Authentication.UserName).Should(Equal("admin"))
				// //Expect(connection.Rediss.Authentication.Password).Should(Equal("$PASSWORD"))
				// Expect(connection.Rediss.Certificate.Name).Should(Equal("c5f07736-d94c-11e8-a2e9-62ec2ed68f84"))
				// Expect(connection.Rediss.Certificate.CertificateBase64).Should(Equal("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUREekNDQWZlZ0F3SUJBZ0lKQU5FSDU4eTIva3pITUEwR0NTcUdTSWIzRFFFQkN3VUFNQjR4SERBYUJnTlYKQkFNTUUwbENUU0JEYkc5MVpDQkVZWFJoWW1GelpYTXdIaGNOTVRnd05qSTFNVFF5T1RBd1doY05Namd3TmpJeQpNVFF5T1RBd1dqQWVNUnd3R2dZRFZRUUREQk5KUWswZ1EyeHZkV1FnUkdGMFlXSmhjMlZ6TUlJQklqQU5CZ2txCmhraUc5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0NBUUVBOGxwYVFHemNGZEdxZU1sbXFqZmZNUHBJUWhxcGQ4cUoKUHIzYklrclhKYlRjSko5dUlja1NVY0NqdzRaL3JTZzhublQxM1NDY09sKzF0bys3a2RNaVU4cU9XS2ljZVlaNQp5K3laWWZDa0dhaVpWZmF6UUJtNDV6QnRGV3YrQUIvOGhmQ1RkTkY3Vlk0c3BhQTNvQkUyYVM3T0FOTlNSWlNLCnB3eTI0SVVnVWNJTEpXK21jdlc4MFZ4K0dYUmZEOVl0dDZQUkpnQmhZdVVCcGd6dm5nbUNNR0JuK2wyS05pU2YKd2VvdllEQ0Q2Vm5nbDIrNlc5UUZBRnRXWFdnRjNpRFFENW5sL240bXJpcE1TWDZVRy9uNjY1N3U3VERkZ2t2QQoxZUtJMkZMellLcG9LQmU1cmNuck03bkhnTmMvbkNkRXM1SmVjSGIxZEh2MVFmUG02cHpJeHdJREFRQUJvMUF3ClRqQWRCZ05WSFE0RUZnUVVLMytYWm8xd3lLcytERW9ZWGJIcnV3U3BYamd3SHdZRFZSMGpCQmd3Rm9BVUszK1gKWm8xd3lLcytERW9ZWGJIcnV3U3BYamd3REFZRFZSMFRCQVV3QXdFQi96QU5CZ2txaGtpRzl3MEJBUXNGQUFPQwpBUUVBSmY1ZHZselVwcWFpeDI2cUpFdXFGRzBJUDU3UVFJNVRDUko2WHQvc3VwUkhvNjNlRHZLdzh6Ujd0bFdRCmxWNVAwTjJ4d3VTbDlacUFKdDcvay8zWmVCK25Zd1BveU8zS3ZLdkFUdW5SdmxQQm40RldWWGVhUHNHKzdmaFMKcXNlam1reW9uWXc3N0hSekdPekpINFpnOFVONm1mcGJhV1NzeWFFeHZxa25DcDlTb1RRUDNENjdBeldxYjF6WQpkb3FxZ0dJWjJueENrcDUvRlh4Ri9UTWI1NXZ0ZVRRd2ZnQnk2MGpWVmtiRjdlVk9XQ3YwS2FOSFBGNWhycWJOCmkrM1hqSjcvcGVGM3hNdlRNb3kzNURjVDNFMlplU1Zqb3VaczE1Tzkwa0kzazJkYVMyT0hKQUJXMHZTajRuTHoKK1BRenAvQjljUW1PTzhkQ2UwNDlRM29hVUE9PQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCgo="))
				// Expect(connection.Rediss.Database).Should(BeEquivalentTo("0"))
				// Expect(connection.Cli.Composed[0]).Should(Equal("REDIS_CERTFILE=c5f07736-d94c-11e8-a2e9-62ec2ed68f84 redli -u rediss://admin:$PASSWORD@9f7c2c13-5061-46ca-8153-f7d6c0baba7b.974550db55eb4ec0983f023940bf637f.databases.appdomain.cloud:30348/0"))
				// Expect(connection.Cli.Type).Should(Equal("cli"))
				// Expect(connection.Cli.Bin).Should(Equal("redli"))
				// Expect(connection.Cli.Arguments[0][0]).Should(Equal("-u"))
				// Expect(connection.Cli.Arguments[0][1]).Should(Equal("rediss://admin:$PASSWORD@9f7c2c13-5061-46ca-8153-f7d6c0baba7b.974550db55eb4ec0983f023940bf637f.databases.appdomain.cloud:30348/0"))
				// Expect(connection.Cli.Certificate.Name).Should(Equal("c5f07736-d94c-11e8-a2e9-62ec2ed68f84"))
				// Expect(connection.Cli.Certificate.CertificateBase64).Should(Equal("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUREekNDQWZlZ0F3SUJBZ0lKQU5FSDU4eTIva3pITUEwR0NTcUdTSWIzRFFFQkN3VUFNQjR4SERBYUJnTlYKQkFNTUUwbENUU0JEYkc5MVpDQkVZWFJoWW1GelpYTXdIaGNOTVRnd05qSTFNVFF5T1RBd1doY05Namd3TmpJeQpNVFF5T1RBd1dqQWVNUnd3R2dZRFZRUUREQk5KUWswZ1EyeHZkV1FnUkdGMFlXSmhjMlZ6TUlJQklqQU5CZ2txCmhraUc5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0NBUUVBOGxwYVFHemNGZEdxZU1sbXFqZmZNUHBJUWhxcGQ4cUoKUHIzYklrclhKYlRjSko5dUlja1NVY0NqdzRaL3JTZzhublQxM1NDY09sKzF0bys3a2RNaVU4cU9XS2ljZVlaNQp5K3laWWZDa0dhaVpWZmF6UUJtNDV6QnRGV3YrQUIvOGhmQ1RkTkY3Vlk0c3BhQTNvQkUyYVM3T0FOTlNSWlNLCnB3eTI0SVVnVWNJTEpXK21jdlc4MFZ4K0dYUmZEOVl0dDZQUkpnQmhZdVVCcGd6dm5nbUNNR0JuK2wyS05pU2YKd2VvdllEQ0Q2Vm5nbDIrNlc5UUZBRnRXWFdnRjNpRFFENW5sL240bXJpcE1TWDZVRy9uNjY1N3U3VERkZ2t2QQoxZUtJMkZMellLcG9LQmU1cmNuck03bkhnTmMvbkNkRXM1SmVjSGIxZEh2MVFmUG02cHpJeHdJREFRQUJvMUF3ClRqQWRCZ05WSFE0RUZnUVVLMytYWm8xd3lLcytERW9ZWGJIcnV3U3BYamd3SHdZRFZSMGpCQmd3Rm9BVUszK1gKWm8xd3lLcytERW9ZWGJIcnV3U3BYamd3REFZRFZSMFRCQVV3QXdFQi96QU5CZ2txaGtpRzl3MEJBUXNGQUFPQwpBUUVBSmY1ZHZselVwcWFpeDI2cUpFdXFGRzBJUDU3UVFJNVRDUko2WHQvc3VwUkhvNjNlRHZLdzh6Ujd0bFdRCmxWNVAwTjJ4d3VTbDlacUFKdDcvay8zWmVCK25Zd1BveU8zS3ZLdkFUdW5SdmxQQm40RldWWGVhUHNHKzdmaFMKcXNlam1reW9uWXc3N0hSekdPekpINFpnOFVONm1mcGJhV1NzeWFFeHZxa25DcDlTb1RRUDNENjdBeldxYjF6WQpkb3FxZ0dJWjJueENrcDUvRlh4Ri9UTWI1NXZ0ZVRRd2ZnQnk2MGpWVmtiRjdlVk9XQ3YwS2FOSFBGNWhycWJOCmkrM1hqSjcvcGVGM3hNdlRNb3kzNURjVDNFMlplU1Zqb3VaczE1Tzkwa0kzazJkYVMyT0hKQUJXMHZTajRuTHoKK1BRenAvQjljUW1PTzhkQ2UwNDlRM29hVUE9PQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCgo="))
			})
		})
		Context("When get is unsuccessful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v4/ibm/deployments/crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a/users/anyone/connections"),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to get connection`),
					),
				)
			})

			It("should return error during get connection", func() {
				target := "crn:v1:staging:public:iam::::apikey:ApiKey-62fefdd1-4557-4c7d-8a1c-f6da7ee2ff3a"
				userId := "anyone"
				_, err := newConnection(server.URL()).GetConnection(target, userId)
				Expect(err).To(HaveOccurred())
				//Expect(connection).Should(Equal(""))
			})
		})
	})
})

func newConnection(url string) Connections {

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
	return newConnectionAPI(&client)
}
