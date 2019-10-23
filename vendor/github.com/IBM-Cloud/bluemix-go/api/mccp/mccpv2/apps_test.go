package mccpv2

import (
	"log"
	"net/http"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/client"
	"github.com/IBM-Cloud/bluemix-go/helpers"
	"github.com/IBM-Cloud/bluemix-go/session"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Apps", func() {
	var server *ghttp.Server
	AfterEach(func() {
		server.Close()
	})

	Describe("List", func() {
		Context("Retrieving apps is successful", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/apps"),
						ghttp.RespondWith(http.StatusOK, `{
  "total_results": 3,
  "total_pages": 1,
  "prev_url": null,
  "next_url": null,
  "resources": [
    {
      "metadata": {
        "guid": "50a46e93-c9dc-4d85-b76a-4f23c95e7cda",
        "url": "/v2/apps/50a46e93-c9dc-4d85-b76a-4f23c95e7cda",
        "created_at": "2016-03-30T18:32:20Z",
        "updated_at": "2017-02-13T05:36:58Z"
      },
      "entity": {
        "name": "abc-test123",
        "production": false,
        "space_guid": "15161161-8b6d-4643-bffb-71569e07f0e9",
        "stack_guid": "ac91d31a-86a3-453b-babf-8d49c9d763fc",
        "buildpack": null,
        "detected_buildpack": "Liberty for Java(TM) (WAR, liberty-2016.3.0_0, buildpack-v2.7-20160321-1358, ibmjdk-1.8.0_20160221, env, spring-auto-reconfiguration-1.10.0_RELEASE)",
        "environment_json": {
          "redacted_message": "[PRIVATE DATA HIDDEN]"
        },
        "memory": 1024,
        "instances": 1,
        "disk_quota": 1024,
        "state": "STARTED",
        "version": "f722emccp5-fc16-4b8b-a8ea-e0d44c56a465",
        "command": null,
        "console": false,
        "debug": null,
        "staging_task_id": "a7808aec96b34e359c47e4320796f221",
        "package_state": "STAGED",
        "health_check_type": "port",
        "health_check_timeout": null,
        "health_check_http_endpoint": null,
        "staging_failed_reason": null,
        "staging_failed_description": null,
        "diego": true,
        "docker_image": null,
        "package_updated_at": "2016-04-07T14:29:33Z",
        "detected_start_command": ".liberty/initial_startup.rb",
        "enable_ssh": true,
        "docker_credentials_json": {
          "redacted_message": "[PRIVATE DATA HIDDEN]"
        },
        "ports": [
          8080
        ],
        "space_url": "/v2/spaces/15161161-8b6d-4643-bffb-71569e07f0e9",
        "stack_url": "/v2/stacks/ac91d31a-86a3-453b-babf-8d49c9d763fc",
        "routes_url": "/v2/apps/50a46e93-c9dc-4d85-b76a-4f23c95e7cda/routes",
        "events_url": "/v2/apps/50a46e93-c9dc-4d85-b76a-4f23c95e7cda/events",
        "service_bindings_url": "/v2/apps/50a46e93-c9dc-4d85-b76a-4f23c95e7cda/service_bindings",
        "route_mappings_url": "/v2/apps/50a46e93-c9dc-4d85-b76a-4f23c95e7cda/route_mappings"
      }
    },
    {
      "metadata": {
        "guid": "413a703d-f9f8-4ed1-a350-a45171c6fe60",
        "url": "/v2/apps/413a703d-f9f8-4ed1-a350-a45171c6fe60",
        "created_at": "2016-03-30T19:17:13Z",
        "updated_at": "2017-02-13T05:37:24Z"
      },
      "entity": {
        "name": "WebSphere-Application-Server-Liberty-at-localhost",
        "production": false,
        "space_guid": "15161161-8b6d-4643-bffb-71569e07f0e9",
        "stack_guid": "ac91d31a-86a3-453b-babf-8d49c9d763fc",
        "buildpack": null,
        "detected_buildpack": null,
        "environment_json": {
          "redacted_message": "[PRIVATE DATA HIDDEN]"
        },
        "memory": 512,
        "instances": 1,
        "disk_quota": 1024,
        "state": "STOPPED",
        "version": "e3f2189d-c0e7-43fe-9e1b-cbd4ff827eee",
        "command": null,
        "console": false,
        "debug": null,
        "staging_task_id": null,
        "package_state": "PENDING",
        "health_check_type": "port",
        "health_check_timeout": null,
        "health_check_http_endpoint": null,
        "staging_failed_reason": null,
        "staging_failed_description": null,
        "diego": true,
        "docker_image": null,
        "package_updated_at": null,
        "detected_start_command": "",
        "enable_ssh": true,
        "docker_credentials_json": {
          "redacted_message": "[PRIVATE DATA HIDDEN]"
        },
        "ports": [
          8080
        ],
        "space_url": "/v2/spaces/15161161-8b6d-4643-bffb-71569e07f0e9",
        "stack_url": "/v2/stacks/ac91d31a-86a3-453b-babf-8d49c9d763fc",
        "routes_url": "/v2/apps/413a703d-f9f8-4ed1-a350-a45171c6fe60/routes",
        "events_url": "/v2/apps/413a703d-f9f8-4ed1-a350-a45171c6fe60/events",
        "service_bindings_url": "/v2/apps/413a703d-f9f8-4ed1-a350-a45171c6fe60/service_bindings",
        "route_mappings_url": "/v2/apps/413a703d-f9f8-4ed1-a350-a45171c6fe60/route_mappings"
      }
    },
    {
      "metadata": {
        "guid": "c92bc8ea-fc00-4011-ad19-36fbd47a4c30",
        "url": "/v2/apps/c92bc8ea-fc00-4011-ad19-36fbd47a4c30",
        "created_at": "2017-05-19T05:31:56Z",
        "updated_at": "2017-05-19T05:32:41Z"
      },
      "entity": {
        "name": "testsakshi1051",
        "production": false,
        "space_guid": "211b690c-1241-496e-b6ae-e487b7ebe4e8",
        "stack_guid": "ac91d31a-86a3-453b-babf-8d49c9d763fc",
        "buildpack": "nodejs_buildpack",
        "detected_buildpack": "",
        "environment_json": {

        },
        "memory": 128,
        "instances": 2,
        "disk_quota": 512,
        "state": "STARTED",
        "version": "b5658b95-7f69-4f35-91a0-4105a850f134",
        "command": null,
        "console": false,
        "debug": null,
        "staging_task_id": "15f5c4610e7b4f5fb42aa98c06a7172c",
        "package_state": "STAGED",
        "health_check_type": "port",
        "health_check_timeout": null,
        "health_check_http_endpoint": null,
        "staging_failed_reason": null,
        "staging_failed_description": null,
        "diego": true,
        "docker_image": null,
        "package_updated_at": "2017-05-19T05:32:02Z",
        "detected_start_command": "npm start",
        "enable_ssh": true,
        "docker_credentials_json": {
          "redacted_message": "[PRIVATE DATA HIDDEN]"
        },
        "ports": [
          8080
        ],
        "space_url": "/v2/spaces/211b690c-1241-496e-b6ae-e487b7ebe4e8",
        "stack_url": "/v2/stacks/ac91d31a-86a3-453b-babf-8d49c9d763fc",
        "routes_url": "/v2/apps/c92bc8ea-fc00-4011-ad19-36fbd47a4c30/routes",
        "events_url": "/v2/apps/c92bc8ea-fc00-4011-ad19-36fbd47a4c30/events",
        "service_bindings_url": "/v2/apps/c92bc8ea-fc00-4011-ad19-36fbd47a4c30/service_bindings",
        "route_mappings_url": "/v2/apps/c92bc8ea-fc00-4011-ad19-36fbd47a4c30/route_mappings"
      }
    }
  ]
}`),
					),
				)
			})

			It("should list all the apps", func() {
				myapps, err := newApps(server.URL()).List()
				Expect(err).To(Succeed())
				Expect(len(myapps)).To(Equal(3))
				app1 := myapps[0]
				Expect(app1.GUID).To(Equal("50a46e93-c9dc-4d85-b76a-4f23c95e7cda"))
				Expect(app1.Name).To(Equal("abc-test123"))
				Expect(app1.SpaceGUID).To(Equal("15161161-8b6d-4643-bffb-71569e07f0e9"))
				app2 := myapps[1]
				Expect(app2.GUID).To(Equal("413a703d-f9f8-4ed1-a350-a45171c6fe60"))
				Expect(app2.Name).To(Equal("WebSphere-Application-Server-Liberty-at-localhost"))
				Expect(app2.SpaceGUID).To(Equal("15161161-8b6d-4643-bffb-71569e07f0e9"))
				app3 := myapps[2]
				Expect(app3.GUID).To(Equal("c92bc8ea-fc00-4011-ad19-36fbd47a4c30"))
				Expect(app3.Name).To(Equal("testsakshi1051"))
				Expect(app3.SpaceGUID).To(Equal("211b690c-1241-496e-b6ae-e487b7ebe4e8"))

			})
		})
		Context("When Server return no apps", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/apps"),
						ghttp.RespondWith(http.StatusOK, `{
							"total_results": 0,
							"resources": [
							]
															
						}`),
					),
				)
			})
			It("should return no apps", func() {
				myapps, err := newApps(server.URL()).List()
				Expect(err).To(Succeed())
				Expect(len(myapps)).To(Equal(0))
			})

		})
		Context("Server return error", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodGet, "/v2/apps"),
						ghttp.RespondWith(http.StatusInternalServerError, `{
															
						}`),
					),
				)
			})

			It("should return error", func() {
				myapps, err := newApps(server.URL()).List()
				Expect(err).To(HaveOccurred())
				Expect(myapps).To(BeNil())
			})
		})

	})

	Describe("Create", func() {
		Context("Create a app", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v2/apps"),
						ghttp.VerifyBody([]byte(`{"name":"testapp","memory":128,"instances":2,"disk_quota":512,"space_guid":"211b690c-1241-496e-b6ae-e487b7ebe4e8","buildpack":"nodejs_buildpack"}`)),
						ghttp.RespondWith(http.StatusCreated, `{
  "metadata": {
    "guid": "18231a23-6056-495f-9ae2-c6c79a15e86c",
    "url": "/v2/apps/18231a23-6056-495f-9ae2-c6c79a15e86c",
    "created_at": "2017-05-19T08:31:38Z",
    "updated_at": null
  },
  "entity": {
    "name": "testapp",
    "production": false,
    "space_guid": "211b690c-1241-496e-b6ae-e487b7ebe4e8",
    "stack_guid": "ac91d31a-86a3-453b-babf-8d49c9d763fc",
    "buildpack": "nodejs_buildpack",
    "detected_buildpack": null,
    "environment_json": {

    },
    "memory": 128,
    "instances": 2,
    "disk_quota": 512,
    "state": "STOPPED",
    "version": "23ba0967-9c2e-4927-9c4a-f5ba2dab6f49",
    "command": null,
    "console": false,
    "debug": null,
    "staging_task_id": null,
    "package_state": "PENDING",
    "health_check_type": "port",
    "health_check_timeout": null,
    "health_check_http_endpoint": null,
    "staging_failed_reason": null,
    "staging_failed_description": null,
    "diego": true,
    "docker_image": null,
    "package_updated_at": null,
    "detected_start_command": "",
    "enable_ssh": true,
    "docker_credentials_json": {
      "redacted_message": "[PRIVATE DATA HIDDEN]"
    },
    "ports": [
      8080
    ],
    "space_url": "/v2/spaces/211b690c-1241-496e-b6ae-e487b7ebe4e8",
    "stack_url": "/v2/stacks/ac91d31a-86a3-453b-babf-8d49c9d763fc",
    "routes_url": "/v2/apps/18231a23-6056-495f-9ae2-c6c79a15e86c/routes",
    "events_url": "/v2/apps/18231a23-6056-495f-9ae2-c6c79a15e86c/events",
    "service_bindings_url": "/v2/apps/18231a23-6056-495f-9ae2-c6c79a15e86c/service_bindings",
    "route_mappings_url": "/v2/apps/18231a23-6056-495f-9ae2-c6c79a15e86c/route_mappings"
  }
}`),
					),
				)
			})

			It("should create the apps", func() {
				var appPayload = AppRequest{
					Name:      helpers.String("testapp"),
					SpaceGUID: helpers.String("211b690c-1241-496e-b6ae-e487b7ebe4e8"),
					BuildPack: helpers.String("nodejs_buildpack"),
					Instances: 2,
					Memory:    128,
					DiskQuota: 512,
				}
				myapp, err := newApps(server.URL()).Create(appPayload)
				Expect(err).NotTo(HaveOccurred())
				Expect(myapp).ShouldNot(BeNil())
				Expect(myapp.Metadata.GUID).Should(Equal("18231a23-6056-495f-9ae2-c6c79a15e86c"))
				Expect(myapp.Entity.Name).Should(Equal("testapp"))
				Expect(myapp.Entity.SpaceGUID).Should(Equal("211b690c-1241-496e-b6ae-e487b7ebe4e8"))

			})
		})
		Context("When app creation failed", func() {
			BeforeEach(func() {
				server = ghttp.NewServer()
				server.SetAllowUnhandledRequests(true)
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest(http.MethodPost, "/v2/apps"),
						ghttp.VerifyBody([]byte(`{"name":"testapp","memory":128,"instances":2,"disk_quota":512,"space_guid":"211b690c-1241-496e-b6ae-e487b7ebe4e8","buildpack":"nodejs_buildpack"}`)),
						ghttp.RespondWith(http.StatusInternalServerError, `Failed to create app`),
					),
				)
			})
			It("should return error while creating app", func() {
				var appPayload = AppRequest{
					Name:      helpers.String("testapp"),
					SpaceGUID: helpers.String("211b690c-1241-496e-b6ae-e487b7ebe4e8"),
					BuildPack: helpers.String("nodejs_buildpack"),
					Instances: 2,
					Memory:    128,
					DiskQuota: 512,
				}
				myapp, err := newApps(server.URL()).Create(appPayload)
				Expect(err).To(HaveOccurred())
				Expect(myapp).Should(BeNil())
			})

		})

		Describe("Get", func() {
			Context("Get a particular app details", func() {
				BeforeEach(func() {
					server = ghttp.NewServer()
					server.AppendHandlers(
						ghttp.CombineHandlers(
							ghttp.VerifyRequest(http.MethodGet, "/v2/apps/26d673fd-7e64-49b1-9d00-20c0edc9094b"),
							ghttp.RespondWith(http.StatusOK, `{
  "metadata": {
    "guid": "26d673fd-7e64-49b1-9d00-20c0edc9094b",
    "url": "/v2/apps/26d673fd-7e64-49b1-9d00-20c0edc9094b",
    "created_at": "2017-05-21T16:00:05Z",
    "updated_at": "2017-05-21T16:00:49Z"
  },
  "entity": {
    "name": "testsakshi1051",
    "production": false,
    "space_guid": "211b690c-1241-496e-b6ae-e487b7ebe4e8",
    "stack_guid": "ac91d31a-86a3-453b-babf-8d49c9d763fc",
    "buildpack": "nodejs_buildpack",
    "detected_buildpack": "",
    "environment_json": {

    },
    "memory": 128,
    "instances": 2,
    "disk_quota": 512,
    "state": "STARTED",
    "version": "6ba05631-fd78-4f20-9ad8-5cd0b37c000c",
    "command": null,
    "console": false,
    "debug": null,
    "staging_task_id": "cd4a74fd9e4a443397a688f15920d830",
    "package_state": "STAGED",
    "health_check_type": "port",
    "health_check_timeout": null,
    "health_check_http_endpoint": null,
    "staging_failed_reason": null,
    "staging_failed_description": null,
    "diego": true,
    "docker_image": null,
    "package_updated_at": "2017-05-21T16:00:10Z",
    "detected_start_command": "npm start",
    "enable_ssh": true,
    "docker_credentials_json": {
      "redacted_message": "[PRIVATE DATA HIDDEN]"
    },
    "ports": [
      8080
    ],
    "space_url": "/v2/spaces/211b690c-1241-496e-b6ae-e487b7ebe4e8",
    "stack_url": "/v2/stacks/ac91d31a-86a3-453b-babf-8d49c9d763fc",
    "routes_url": "/v2/apps/26d673fd-7e64-49b1-9d00-20c0edc9094b/routes",
    "events_url": "/v2/apps/26d673fd-7e64-49b1-9d00-20c0edc9094b/events",
    "service_bindings_url": "/v2/apps/26d673fd-7e64-49b1-9d00-20c0edc9094b/service_bindings",
    "route_mappings_url": "/v2/apps/26d673fd-7e64-49b1-9d00-20c0edc9094b/route_mappings"
  }
}`),
						),
					)
				})

				It("should return the details of particular apps", func() {

					myapp, err := newApps(server.URL()).Get("26d673fd-7e64-49b1-9d00-20c0edc9094b")
					Expect(err).NotTo(HaveOccurred())
					Expect(myapp).ShouldNot(BeNil())
					Expect(myapp.Metadata.GUID).Should(Equal("26d673fd-7e64-49b1-9d00-20c0edc9094b"))
					Expect(myapp.Entity.Name).Should(Equal("testsakshi1051"))
					Expect(myapp.Entity.SpaceGUID).Should(Equal("211b690c-1241-496e-b6ae-e487b7ebe4e8"))

				})
			})
			Context("When app retrievel failed", func() {
				BeforeEach(func() {
					server = ghttp.NewServer()
					server.SetAllowUnhandledRequests(true)
					server.AppendHandlers(
						ghttp.CombineHandlers(
							ghttp.VerifyRequest(http.MethodGet, "/v2/apps/26d673fd-7e64-49b1-9d00-20c0edc9094b"),
							ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve app`),
						),
					)
				})
				It("should return error while retrieving app", func() {
					myapp, err := newApps(server.URL()).Get("26d673fd-7e64-49b1-9d00-20c0edc9094b")
					Expect(err).To(HaveOccurred())
					Expect(myapp).Should(BeNil())
				})

			})

		})

		Describe("Update", func() {
			Context("When update of app is successful", func() {
				BeforeEach(func() {
					server = ghttp.NewServer()
					server.AppendHandlers(
						ghttp.CombineHandlers(
							ghttp.VerifyRequest(http.MethodPut, "/v2/apps/26d673fd-7e64-49b1-9d00-20c0edc9094b"),
							ghttp.VerifyBody([]byte(`{"name":"testappupdate","space_guid":"211b690c-1241-496e-b6ae-e487b7ebe4e8"}`)),
							ghttp.RespondWith(http.StatusOK, `{
  "metadata": {
    "guid": "26d673fd-7e64-49b1-9d00-20c0edc9094b",
    "url": "/v2/apps/26d673fd-7e64-49b1-9d00-20c0edc9094b",
    "created_at": "2017-05-21T16:00:05Z",
    "updated_at": "2017-05-21T16:01:19Z"
  },
  "entity": {
    "name": "testappupdate",
    "production": false,
    "space_guid": "211b690c-1241-496e-b6ae-e487b7ebe4e8",
    "stack_guid": "ac91d31a-86a3-453b-babf-8d49c9d763fc",
    "buildpack": "nodejs_buildpack",
    "detected_buildpack": "",
    "environment_json": {

    },
    "memory": 128,
    "instances": 2,
    "disk_quota": 512,
    "state": "STARTED",
    "version": "6ba05631-fd78-4f20-9ad8-5cd0b37c000c",
    "command": null,
    "console": false,
    "debug": null,
    "staging_task_id": "cd4a74fd9e4a443397a688f15920d830",
    "package_state": "STAGED",
    "health_check_type": "port",
    "health_check_timeout": null,
    "health_check_http_endpoint": null,
    "staging_failed_reason": null,
    "staging_failed_description": null,
    "diego": true,
    "docker_image": null,
    "package_updated_at": "2017-05-21T16:00:10Z",
    "detected_start_command": "npm start",
    "enable_ssh": true,
    "docker_credentials_json": {
      "redacted_message": "[PRIVATE DATA HIDDEN]"
    },
    "ports": [
      8080
    ],
    "space_url": "/v2/spaces/211b690c-1241-496e-b6ae-e487b7ebe4e8",
    "stack_url": "/v2/stacks/ac91d31a-86a3-453b-babf-8d49c9d763fc",
    "routes_url": "/v2/apps/26d673fd-7e64-49b1-9d00-20c0edc9094b/routes",
    "events_url": "/v2/apps/26d673fd-7e64-49b1-9d00-20c0edc9094b/events",
    "service_bindings_url": "/v2/apps/26d673fd-7e64-49b1-9d00-20c0edc9094b/service_bindings",
    "route_mappings_url": "/v2/apps/26d673fd-7e64-49b1-9d00-20c0edc9094b/route_mappings"
  }
}`),
						),
					)
				})

				It("should return app update", func() {
					var appUpdatePayload = AppRequest{
						Name:      helpers.String("testappupdate"),
						SpaceGUID: helpers.String("211b690c-1241-496e-b6ae-e487b7ebe4e8"),
					}
					myapp, err := newApps(server.URL()).Update("26d673fd-7e64-49b1-9d00-20c0edc9094b", appUpdatePayload)
					Expect(err).NotTo(HaveOccurred())
					Expect(myapp).ShouldNot(BeNil())
					Expect(myapp.Metadata.GUID).Should(Equal("26d673fd-7e64-49b1-9d00-20c0edc9094b"))
					Expect(myapp.Entity.Name).Should(Equal("testappupdate"))
					Expect(myapp.Entity.SpaceGUID).Should(Equal("211b690c-1241-496e-b6ae-e487b7ebe4e8"))
				})
			})
			Context("When app update is failed", func() {
				BeforeEach(func() {
					server = ghttp.NewServer()
					server.SetAllowUnhandledRequests(true)
					server.AppendHandlers(
						ghttp.CombineHandlers(
							ghttp.VerifyRequest(http.MethodPut, "/v2/apps/26d673fd-7e64-49b1-9d00-20c0edc9094b"),
							ghttp.VerifyBody([]byte(`{"name":"testappupdate","space_guid":"211b690c-1241-496e-b6ae-e487b7ebe4e8"}`)),
							ghttp.RespondWith(http.StatusInternalServerError, `Failed to retrieve app`),
						),
					)
				})

				It("should return error when app updated", func() {
					var appUpdatePayload = AppRequest{
						Name:      helpers.String("testappupdate"),
						SpaceGUID: helpers.String("211b690c-1241-496e-b6ae-e487b7ebe4e8"),
					}
					myapp, err := newApps(server.URL()).Update("26d673fd-7e64-49b1-9d00-20c0edc9094b", appUpdatePayload)
					Expect(err).To(HaveOccurred())
					Expect(myapp).Should(BeNil())
				})
			})
		})

		Describe("Delete", func() {
			Context("When delete of app is successful", func() {
				BeforeEach(func() {
					server = ghttp.NewServer()
					server.AppendHandlers(
						ghttp.CombineHandlers(
							ghttp.VerifyRequest(http.MethodDelete, "/v2/apps/26d673fd-7e64-49b1-9d00-20c0edc9094b"),
							ghttp.RespondWith(http.StatusNoContent, `{}`),
						),
					)
				})

				It("should delete the app", func() {

					err := newApps(server.URL()).Delete("26d673fd-7e64-49b1-9d00-20c0edc9094b", false, false)
					Expect(err).NotTo(HaveOccurred())

				})
			})
			Context("When app delete is failed", func() {
				BeforeEach(func() {
					server = ghttp.NewServer()
					server.SetAllowUnhandledRequests(true)
					server.AppendHandlers(
						ghttp.CombineHandlers(
							ghttp.VerifyRequest(http.MethodDelete, "/v2/apps/26d673fd-7e64-49b1-9d00-20c0edc9094b"),
							ghttp.RespondWith(http.StatusInternalServerError, `Failed to delete app`),
						),
					)
				})

				It("should return error when app deleted", func() {

					err := newApps(server.URL()).Delete("26d673fd-7e64-49b1-9d00-20c0edc9094b", false, false)
					Expect(err).To(HaveOccurred())

				})
			})
		})

	})
})

func newApps(url string) Apps {

	sess, err := session.New()
	if err != nil {
		log.Fatal(err)
	}
	conf := sess.Config.Copy()
	conf.Endpoint = &url
	client := client.Client{
		Config:      conf,
		ServiceName: bluemix.MccpService,
	}
	return newAppAPI(&client)
}
