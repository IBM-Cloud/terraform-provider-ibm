//go:build integration
// +build integration

/**
 * (C) Copyright IBM Corp. 2025.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package atrackerv2_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/atrackerv2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the atrackerv2 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`AtrackerV2 Integration Tests`, func() {

	const externalConfigFile = "../atracker_v2.env"

	const notFoundTargetID = "ffffffff-1111-1111-1111-111111111111"

	const notFoundRouteID = "ffffffff-2222-2222-2222-222222222222"

	const badTargetType = "bad_target_type"

	var (
		err                          error
		atrackerService              *atrackerv2.AtrackerV2
		atrackerServiceNotAuthorized *atrackerv2.AtrackerV2
		serviceURL                   string
		config                       map[string]string
		refreshTokenNotAuthorized    string

		// Variables to hold link values
		routeIDLink   string
		targetIDLink  string
		targetIDLink3 string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(atrackerv2.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			atrackerServiceOptions := &atrackerv2.AtrackerV2Options{}

			atrackerService, err = atrackerv2.NewAtrackerV2UsingExternalConfig(atrackerServiceOptions)

			Expect(err).To(BeNil())
			Expect(atrackerService).ToNot(BeNil())
			Expect(atrackerService.Service.Options.URL).To(Equal(serviceURL))

			atrackerUnauthorizedServiceOptions := &atrackerv2.AtrackerV2Options{
				ServiceName: "NOT_AUTHORIZED",
			}
			atrackerServiceNotAuthorized, err = atrackerv2.NewAtrackerV2UsingExternalConfig(atrackerUnauthorizedServiceOptions)
			Expect(err).To(BeNil())
			Expect(atrackerServiceNotAuthorized).ToNot(BeNil())
			Expect(atrackerServiceNotAuthorized.Service.Options.URL).To(Equal(serviceURL))

			tokenNotAuthorized, err := atrackerServiceNotAuthorized.Service.Options.Authenticator.(*core.IamAuthenticator).RequestToken()
			Expect(err).To(BeNil())
			refreshTokenNotAuthorized = tokenNotAuthorized.RefreshToken
			Expect(refreshTokenNotAuthorized).ToNot(BeNil())

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			atrackerService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`CreateTarget - Create a target`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTarget(createTargetOptions *CreateTargetOptions)`, func() {

			cosEndpointPrototypeModel := &atrackerv2.CosEndpointPrototype{
				Endpoint:                core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud"),
				TargetCRN:               core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
				Bucket:                  core.StringPtr("my-atracker-bucket"),
				APIKey:                  core.StringPtr("xxxxxxxxxxxxxx"),
				ServiceToServiceEnabled: core.BoolPtr(true),
			}

			createTargetOptions := &atrackerv2.CreateTargetOptions{
				Name:        core.StringPtr("my-cos-target"),
				TargetType:  core.StringPtr("cloud_object_storage"),
				CosEndpoint: cosEndpointPrototypeModel,
				Region:      core.StringPtr("us-south"),
			}

			target, response, err := atrackerService.CreateTarget(createTargetOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(target).ToNot(BeNil())

			targetIDLink = *target.ID
			fmt.Fprintf(GinkgoWriter, "Saved cos targetIDLink value: %v\n", targetIDLink)
		})

		It(`CreateTarget(createTargetOptions *CreateTargetOptions)`, func() {

			eventstreamsEndpointPrototypeModel := &atrackerv2.EventstreamsEndpointPrototype{
				TargetCRN:               core.StringPtr("crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
				Topic:                   core.StringPtr("my-test-topic"),
				Brokers:                 []string{"kafka-x:9094"},
				APIKey:                  core.StringPtr("xxxxxxxxxxx"),
				ServiceToServiceEnabled: core.BoolPtr(false),
			}

			createTargetOptions := &atrackerv2.CreateTargetOptions{
				Name:                 core.StringPtr("my-ies-target"),
				TargetType:           core.StringPtr("event_streams"),
				EventstreamsEndpoint: eventstreamsEndpointPrototypeModel,
			}

			target, response, err := atrackerService.CreateTarget(createTargetOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(target).ToNot(BeNil())

			targetIDLink3 = *target.ID
			fmt.Fprintf(GinkgoWriter, "Saved event streams targetIDLink value: %v\n", targetIDLink)
		})

		It(`Returns 400 when backend input validation fails`, func() {
			cosEndpointPrototypeModel := &atrackerv2.CosEndpointPrototype{
				Endpoint:                core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud"),
				TargetCRN:               core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
				Bucket:                  core.StringPtr("my-atracker-bucket"),
				APIKey:                  core.StringPtr("xxxxxxxxxxxxxx"),
				ServiceToServiceEnabled: core.BoolPtr(false),
			}

			createTargetOptions := &atrackerv2.CreateTargetOptions{
				Name:        core.StringPtr("my-cos-target"),
				TargetType:  core.StringPtr(badTargetType),
				CosEndpoint: cosEndpointPrototypeModel,
			}

			_, response, err := atrackerService.CreateTarget(createTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Returns 403 when user is not authorized`, func() {

			cosEndpointPrototypeModel := &atrackerv2.CosEndpointPrototype{
				Endpoint:  core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud"),
				TargetCRN: core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
				Bucket:    core.StringPtr("my-atracker-bucket"),
				APIKey:    core.StringPtr("xxxxxxxxxxxxxx"),
			}

			createTargetOptions := &atrackerv2.CreateTargetOptions{
				Name:        core.StringPtr("my-cos-target"),
				TargetType:  core.StringPtr("cloud_object_storage"),
				CosEndpoint: cosEndpointPrototypeModel,
			}

			_, response, err := atrackerServiceNotAuthorized.CreateTarget(createTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})
	})

	Describe(`CreateRoute - Create a route`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateRoute(createRouteOptions *CreateRouteOptions)`, func() {

			rulePrototypeModel := &atrackerv2.RulePrototype{
				TargetIds: []string{targetIDLink},
				Locations: []string{"us-south"},
			}

			createRouteOptions := &atrackerv2.CreateRouteOptions{
				Name:  core.StringPtr("my-route"),
				Rules: []atrackerv2.RulePrototype{*rulePrototypeModel},
			}

			route, response, err := atrackerService.CreateRoute(createRouteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(route).ToNot(BeNil())
			routeIDLink = *route.ID
			fmt.Fprintf(GinkgoWriter, "Saved routeIDLink value: %v\n", routeIDLink)
		})

		It(`Returns 403 when user is not authorized`, func() {

			rulePrototypeModel := &atrackerv2.RulePrototype{
				TargetIds: []string{targetIDLink},
				Locations: []string{"us-south"},
			}

			createRouteOptions := &atrackerv2.CreateRouteOptions{
				Name:  core.StringPtr("my-route"),
				Rules: []atrackerv2.RulePrototype{*rulePrototypeModel},
			}

			_, response, err := atrackerServiceNotAuthorized.CreateRoute(createRouteOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 400 when input validation fails`, func() {

			rulePrototypeModel := &atrackerv2.RulePrototype{
				TargetIds: []string{notFoundTargetID},
				Locations: []string{"us-south"},
			}

			createRouteOptions := &atrackerv2.CreateRouteOptions{
				Name:  core.StringPtr("my-route"),
				Rules: []atrackerv2.RulePrototype{*rulePrototypeModel},
			}

			_, response, err := atrackerService.CreateRoute(createRouteOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})
	})

	Describe(`ListTargets - List targets`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTargets(listTargetsOptions *ListTargetsOptions)`, func() {

			listTargetsOptions := &atrackerv2.ListTargetsOptions{}

			targetList, response, err := atrackerService.ListTargets(listTargetsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(targetList).ToNot(BeNil())
		})

		It(`Returns 403 when user is not authorized)`, func() {

			listTargetsOptions := &atrackerv2.ListTargetsOptions{}

			_, response, err := atrackerServiceNotAuthorized.ListTargets(listTargetsOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})
	})

	Describe(`GetTarget - Get details of a target`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTarget(getTargetOptions *GetTargetOptions)`, func() {

			getTargetOptions := &atrackerv2.GetTargetOptions{
				ID: &targetIDLink,
			}

			target, response, err := atrackerService.GetTarget(getTargetOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())
		})

		It(`Returns 403 when user is not authorized`, func() {

			getTargetOptions := &atrackerv2.GetTargetOptions{
				ID: core.StringPtr(targetIDLink),
			}

			_, response, err := atrackerServiceNotAuthorized.GetTarget(getTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when target id is not found`, func() {

			getTargetOptions := &atrackerv2.GetTargetOptions{
				ID: core.StringPtr(notFoundTargetID),
			}

			_, response, err := atrackerService.GetTarget(getTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})
	})

	Describe(`ReplaceTarget - Update a target`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceTarget(replaceTargetOptions *ReplaceTargetOptions) for cos type of target`, func() {

			cosEndpointPrototypeModel := &atrackerv2.CosEndpointPrototype{
				Endpoint:                core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud"),
				TargetCRN:               core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
				Bucket:                  core.StringPtr("my-atracker-bucket-modified"),
				APIKey:                  core.StringPtr("xxxxxxxxxxxxxx"),
				ServiceToServiceEnabled: core.BoolPtr(true),
			}
			replaceTargetOptions := &atrackerv2.ReplaceTargetOptions{
				ID:          &targetIDLink,
				Name:        core.StringPtr("my-cos-target"),
				CosEndpoint: cosEndpointPrototypeModel,
			}

			target, response, err := atrackerService.ReplaceTarget(replaceTargetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())
		})

		It(`ReplaceTarget(replaceTargetOptions *ReplaceTargetOptions) for event streams type of target`, func() {

			eventstreamsEndpointPrototypeModel := &atrackerv2.EventstreamsEndpointPrototype{
				TargetCRN:               core.StringPtr("crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
				Topic:                   core.StringPtr("my-test-topic"),
				Brokers:                 []string{"kafka-x:9094"},
				APIKey:                  core.StringPtr("xxxxxxxxxxxxx"),
				ServiceToServiceEnabled: core.BoolPtr(false),
			}

			replaceTargetOptions := &atrackerv2.ReplaceTargetOptions{
				ID:                   &targetIDLink3,
				Name:                 core.StringPtr("my-ies-target-modified"),
				EventstreamsEndpoint: eventstreamsEndpointPrototypeModel,
			}

			target, response, err := atrackerService.ReplaceTarget(replaceTargetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())
		})

		It(`Returns 404 when target id is not found`, func() {

			cosEndpointPrototypeModel := &atrackerv2.CosEndpointPrototype{
				Endpoint:                core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud"),
				TargetCRN:               core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
				Bucket:                  core.StringPtr("my-atracker-bucket"),
				APIKey:                  core.StringPtr("xxxxxxxxxxxxxx"),
				ServiceToServiceEnabled: core.BoolPtr(false),
			}

			replaceTargetOptions := &atrackerv2.ReplaceTargetOptions{
				ID:          core.StringPtr(notFoundTargetID),
				Name:        core.StringPtr("my-cos-target-modified"),
				CosEndpoint: cosEndpointPrototypeModel,
			}
			_, response, err := atrackerService.ReplaceTarget(replaceTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})
	})

	Describe(`ValidateTarget - Validate a target`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ValidateTarget(validateTargetOptions *ValidateTargetOptions)`, func() {

			validateTargetOptions := &atrackerv2.ValidateTargetOptions{
				ID: &targetIDLink,
			}

			target, response, err := atrackerService.ValidateTarget(validateTargetOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())
		})

		It(`Returns 403 when user is not authorized`, func() {

			validateTargetOptions := &atrackerv2.ValidateTargetOptions{
				ID: core.StringPtr(targetIDLink),
			}

			_, response, err := atrackerServiceNotAuthorized.ValidateTarget(validateTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when target id is not found`, func() {

			validateTargetOptions := &atrackerv2.ValidateTargetOptions{
				ID: core.StringPtr(notFoundTargetID),
			}

			_, response, err := atrackerService.ValidateTarget(validateTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})
	})

	Describe(`ListRoutes - List routes`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListRoutes(listRoutesOptions *ListRoutesOptions)`, func() {

			listRoutesOptions := &atrackerv2.ListRoutesOptions{}

			routeList, response, err := atrackerService.ListRoutes(listRoutesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(routeList).ToNot(BeNil())
		})

		It(`Returns 403 when user is not authorized`, func() {

			listRoutesOptions := &atrackerv2.ListRoutesOptions{}

			_, response, err := atrackerServiceNotAuthorized.ListRoutes(listRoutesOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})
	})

	Describe(`GetRoute - Get details of a route`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetRoute(getRouteOptions *GetRouteOptions)`, func() {

			getRouteOptions := &atrackerv2.GetRouteOptions{
				ID: &routeIDLink,
			}

			route, response, err := atrackerService.GetRoute(getRouteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(route).ToNot(BeNil())
		})

		It(`Returns 403 when user is not authorized`, func() {

			getRouteOptions := &atrackerv2.GetRouteOptions{
				ID: core.StringPtr(routeIDLink),
			}

			_, response, err := atrackerServiceNotAuthorized.GetRoute(getRouteOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when route id is not found`, func() {

			getRouteOptions := &atrackerv2.GetRouteOptions{
				ID: core.StringPtr(notFoundRouteID),
			}

			_, response, err := atrackerService.GetRoute(getRouteOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})
	})

	Describe(`ReplaceRoute - Update a route`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceRoute(replaceRouteOptions *ReplaceRouteOptions)`, func() {

			rulePrototypeModel := &atrackerv2.RulePrototype{
				TargetIds: []string{targetIDLink},
				Locations: []string{"us-south"},
			}

			replaceRouteOptions := &atrackerv2.ReplaceRouteOptions{
				ID:    &routeIDLink,
				Name:  core.StringPtr("my-route"),
				Rules: []atrackerv2.RulePrototype{*rulePrototypeModel},
			}

			route, response, err := atrackerService.ReplaceRoute(replaceRouteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(route).ToNot(BeNil())
		})

		It(`Returns 403 when user is not authorized`, func() {

			rulePrototypeModel := &atrackerv2.RulePrototype{
				TargetIds: []string{targetIDLink},
				Locations: []string{"us-south"},
			}

			replaceRouteOptions := &atrackerv2.ReplaceRouteOptions{
				ID:    core.StringPtr(routeIDLink),
				Name:  core.StringPtr("my-route-modified"),
				Rules: []atrackerv2.RulePrototype{*rulePrototypeModel},
			}

			_, response, err := atrackerServiceNotAuthorized.ReplaceRoute(replaceRouteOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when route id is not found`, func() {

			rulePrototypeModel := &atrackerv2.RulePrototype{
				TargetIds: []string{targetIDLink},
				Locations: []string{"us-south"},
			}

			replaceRouteOptions := &atrackerv2.ReplaceRouteOptions{
				ID:    core.StringPtr(notFoundRouteID),
				Name:  core.StringPtr("my-route-modified"),
				Rules: []atrackerv2.RulePrototype{*rulePrototypeModel},
			}

			_, response, err := atrackerService.ReplaceRoute(replaceRouteOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})
	})

	Describe(`GetSettings - Get settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSettings(getSettingsOptions *GetSettingsOptions)`, func() {

			getSettingsOptions := &atrackerv2.GetSettingsOptions{}

			settings, response, err := atrackerService.GetSettings(getSettingsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(settings).ToNot(BeNil())
		})

		It(`Returns 403 when user is not authorized`, func() {

			getSettingsOptions := &atrackerv2.GetSettingsOptions{}

			_, response, err := atrackerServiceNotAuthorized.GetSettings(getSettingsOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})
	})

	Describe(`PutSettings - Modify settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`PutSettings(putSettingsOptions *PutSettingsOptions)`, func() {
			putSettingsOptions := &atrackerv2.PutSettingsOptions{
				MetadataRegionPrimary:  core.StringPtr("us-south"),
				PrivateAPIEndpointOnly: core.BoolPtr(false),
				DefaultTargets:         []string{targetIDLink},
				PermittedTargetRegions: []string{"us-south"},
				MetadataRegionBackup:   core.StringPtr("eu-de"),
			}

			settings, response, err := atrackerService.PutSettings(putSettingsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(settings).ToNot(BeNil())
		})

		It(`Removing default targets`, func() {

			putSettingsOptions := &atrackerv2.PutSettingsOptions{
				DefaultTargets:         []string{},
				PermittedTargetRegions: []string{"us-south"},
				MetadataRegionPrimary:  core.StringPtr("us-south"),
				PrivateAPIEndpointOnly: core.BoolPtr(false),
			}

			settings, response, err := atrackerService.PutSettings(putSettingsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(settings).ToNot(BeNil())
		})
	})

	Describe(`DeleteRoute - Delete a route`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteRoute(deleteRouteOptions *DeleteRouteOptions)`, func() {

			deleteRouteOptions := &atrackerv2.DeleteRouteOptions{
				ID: &routeIDLink,
			}

			response, err := atrackerService.DeleteRoute(deleteRouteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})

		It(`Returns 403 when user is not authorized`, func() {

			deleteRouteOptions := &atrackerv2.DeleteRouteOptions{
				ID: core.StringPtr(routeIDLink),
			}

			response, err := atrackerServiceNotAuthorized.DeleteRoute(deleteRouteOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when route id is not found`, func() {

			deleteRouteOptions := &atrackerv2.DeleteRouteOptions{
				ID: core.StringPtr(notFoundRouteID),
			}

			response, err := atrackerService.DeleteRoute(deleteRouteOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})
	})

	Describe(`DeleteTarget - Delete a target`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`Returns 403 when user is not authorized`, func() {

			deleteTargetOptions := &atrackerv2.DeleteTargetOptions{
				ID: core.StringPtr(targetIDLink),
			}

			_, response, err := atrackerServiceNotAuthorized.DeleteTarget(deleteTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})
		It(`Returns 404 when target id is not found`, func() {

			deleteTargetOptions := &atrackerv2.DeleteTargetOptions{
				ID: core.StringPtr(notFoundTargetID),
			}

			_, response, err := atrackerService.DeleteTarget(deleteTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})
		It(`DeleteTarget(deleteTargetOptions *DeleteTargetOptions)`, func() {

			deleteTargetOptions := &atrackerv2.DeleteTargetOptions{
				ID: &targetIDLink,
			}

			_, response, err := atrackerService.DeleteTarget(deleteTargetOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`DeleteTarget(deleteTargetOptions *DeleteTargetOptions)`, func() {

			deleteTargetOptions := &atrackerv2.DeleteTargetOptions{
				ID: &targetIDLink3,
			}

			_, response, err := atrackerService.DeleteTarget(deleteTargetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
