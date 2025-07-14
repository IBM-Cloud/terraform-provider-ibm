//go:build examples
// +build examples

/**
 * (C) Copyright IBM Corp. 2024.
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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/atrackerv2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// This file provides an example of how to use the atracker service.
//
// The following configuration properties are assumed to be defined:
// ATRACKER_URL=<service base url>
// ATRACKER_AUTH_TYPE=iam
// ATRACKER_APIKEY=<IAM apikey>
// ATRACKER_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
var _ = Describe(`AtrackerV2 Examples Tests`, func() {

	const externalConfigFile = "../atracker_v2.env"

	var (
		atrackerService *atrackerv2.AtrackerV2
		config          map[string]string

		// Variables to hold link values
		routeIDLink  string
		targetIDLink string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping examples...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping examples: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(atrackerv2.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			} else if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
			}

			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			atrackerServiceOptions := &atrackerv2.AtrackerV2Options{}

			atrackerService, err = atrackerv2.NewAtrackerV2UsingExternalConfig(atrackerServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(atrackerService).ToNot(BeNil())
		})
	})

	Describe(`AtrackerV2 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTarget request example`, func() {
			fmt.Println("\nCreateTarget() result:")
			// begin-create_target

			cosEndpointPrototypeModel := &atrackerv2.CosEndpointPrototype{
				Endpoint:                core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud"),
				TargetCRN:               core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"),
				Bucket:                  core.StringPtr("my-atracker-bucket"),
				APIKey:                  core.StringPtr("xxxxxxxxxxxxxx"),
				ServiceToServiceEnabled: core.BoolPtr(false),
			}
			createTargetOptions := atrackerService.NewCreateTargetOptions(
				"my-cos-target",
				"cloud_object_storage",
			)
			createTargetOptions.SetCosEndpoint(cosEndpointPrototypeModel)
			target, response, err := atrackerService.CreateTarget(createTargetOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(target, "", "  ")
			fmt.Println(string(b))

			// end-create_target

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(target).ToNot(BeNil())

			targetIDLink = *target.ID
			fmt.Fprintf(GinkgoWriter, "Saved targetIDLink value: %v\n", targetIDLink)

		})
		It(`CreateRoute request example`, func() {
			fmt.Println("\nCreateRoute() result:")
			// begin-create_route

			rulePrototypeModel := &atrackerv2.RulePrototype{
				TargetIds: []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"},
			}

			createRouteOptions := atrackerService.NewCreateRouteOptions(
				"my-route",
				[]atrackerv2.RulePrototype{*rulePrototypeModel},
			)

			route, response, err := atrackerService.CreateRoute(createRouteOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(route, "", "  ")
			fmt.Println(string(b))

			// end-create_route

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(route).ToNot(BeNil())

			routeIDLink = *route.ID
			fmt.Fprintf(GinkgoWriter, "Saved routeIDLink value: %v\n", routeIDLink)

		})
		It(`ListTargets request example`, func() {
			fmt.Println("\nListTargets() result:")
			// begin-list_targets

			listTargetsOptions := atrackerService.NewListTargetsOptions()

			targetList, response, err := atrackerService.ListTargets(listTargetsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(targetList, "", "  ")
			fmt.Println(string(b))

			// end-list_targets

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(targetList).ToNot(BeNil())

		})
		It(`GetTarget request example`, func() {
			fmt.Println("\nGetTarget() result:")
			// begin-get_target

			getTargetOptions := atrackerService.NewGetTargetOptions(
				targetIDLink,
			)

			target, response, err := atrackerService.GetTarget(getTargetOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(target, "", "  ")
			fmt.Println(string(b))

			// end-get_target

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())

		})
		It(`ReplaceTarget request example`, func() {
			fmt.Println("\nReplaceTarget() result:")
			// begin-replace_target

			replaceTargetOptions := atrackerService.NewReplaceTargetOptions(
				targetIDLink,
			)

			target, response, err := atrackerService.ReplaceTarget(replaceTargetOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(target, "", "  ")
			fmt.Println(string(b))

			// end-replace_target

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())

		})
		It(`ValidateTarget request example`, func() {
			fmt.Println("\nValidateTarget() result:")
			// begin-validate_target

			validateTargetOptions := atrackerService.NewValidateTargetOptions(
				targetIDLink,
			)

			target, response, err := atrackerService.ValidateTarget(validateTargetOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(target, "", "  ")
			fmt.Println(string(b))

			// end-validate_target

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())

		})
		It(`ListRoutes request example`, func() {
			fmt.Println("\nListRoutes() result:")
			// begin-list_routes

			listRoutesOptions := atrackerService.NewListRoutesOptions()

			routeList, response, err := atrackerService.ListRoutes(listRoutesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(routeList, "", "  ")
			fmt.Println(string(b))

			// end-list_routes

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(routeList).ToNot(BeNil())

		})
		It(`GetRoute request example`, func() {
			fmt.Println("\nGetRoute() result:")
			// begin-get_route

			getRouteOptions := atrackerService.NewGetRouteOptions(
				routeIDLink,
			)

			route, response, err := atrackerService.GetRoute(getRouteOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(route, "", "  ")
			fmt.Println(string(b))

			// end-get_route

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(route).ToNot(BeNil())

		})
		It(`ReplaceRoute request example`, func() {
			fmt.Println("\nReplaceRoute() result:")
			// begin-replace_route

			rulePrototypeModel := &atrackerv2.RulePrototype{
				TargetIds: []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"},
			}

			replaceRouteOptions := atrackerService.NewReplaceRouteOptions(
				routeIDLink,
				"my-route",
				[]atrackerv2.RulePrototype{*rulePrototypeModel},
			)

			route, response, err := atrackerService.ReplaceRoute(replaceRouteOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(route, "", "  ")
			fmt.Println(string(b))

			// end-replace_route

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(route).ToNot(BeNil())

		})
		It(`GetSettings request example`, func() {
			fmt.Println("\nGetSettings() result:")
			// begin-get_settings

			getSettingsOptions := atrackerService.NewGetSettingsOptions()

			settings, response, err := atrackerService.GetSettings(getSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(settings, "", "  ")
			fmt.Println(string(b))

			// end-get_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(settings).ToNot(BeNil())

		})
		It(`PutSettings request example`, func() {
			fmt.Println("\nPutSettings() result:")
			// begin-put_settings

			putSettingsOptions := atrackerService.NewPutSettingsOptions(
				"us-south",
				false,
			)

			settings, response, err := atrackerService.PutSettings(putSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(settings, "", "  ")
			fmt.Println(string(b))

			// end-put_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(settings).ToNot(BeNil())

		})
		It(`DeleteTarget request example`, func() {
			fmt.Println("\nDeleteTarget() result:")
			// begin-delete_target

			deleteTargetOptions := atrackerService.NewDeleteTargetOptions(
				targetIDLink,
			)

			warningReport, response, err := atrackerService.DeleteTarget(deleteTargetOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(warningReport, "", "  ")
			fmt.Println(string(b))

			// end-delete_target

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(warningReport).ToNot(BeNil())

		})
		It(`DeleteRoute request example`, func() {
			// begin-delete_route

			deleteRouteOptions := atrackerService.NewDeleteRouteOptions(
				routeIDLink,
			)

			response, err := atrackerService.DeleteRoute(deleteRouteOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteRoute(): %d\n", response.StatusCode)
			}

			// end-delete_route

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})
