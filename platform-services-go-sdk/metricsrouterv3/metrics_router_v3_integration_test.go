//go:build integration

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

package metricsrouterv3_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/metricsrouterv3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the metricsrouterv3 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`MetricsRouterV3 Integration Tests`, func() {
	const externalConfigFile = "../metrics_router_v3.env"
	const notFoundTargetID = "ffffffff-1111-1111-1111-111111111111"
	const notFoundRouteID = "ffffffff-2222-2222-2222-222222222222"

	var (
		err                               error
		metricsRouterService              *metricsrouterv3.MetricsRouterV3
		metricsrouterServiceNotAuthorized *metricsrouterv3.MetricsRouterV3
		serviceURL                        string
		config                            map[string]string

		// Variables to hold link values
		routeIDLink  string
		routeIDLink1 string
		targetIDLink string
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
			config, err = core.GetServiceProperties(metricsrouterv3.DefaultServiceName)
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
			metricsRouterServiceOptions := &metricsrouterv3.MetricsRouterV3Options{}

			metricsRouterService, err = metricsrouterv3.NewMetricsRouterV3UsingExternalConfig(metricsRouterServiceOptions)
			Expect(err).To(BeNil())
			Expect(metricsRouterService).ToNot(BeNil())
			Expect(metricsRouterService.Service.Options.URL).To(Equal(serviceURL))

			metricsrouterUnauthorizedServiceOptions := &metricsrouterv3.MetricsRouterV3Options{
				ServiceName: "NOT_AUTHORIZED",
			}
			metricsrouterServiceNotAuthorized, err = metricsrouterv3.NewMetricsRouterV3UsingExternalConfig(metricsrouterUnauthorizedServiceOptions)
			Expect(err).To(BeNil())
			Expect(metricsrouterServiceNotAuthorized).ToNot(BeNil())
			Expect(metricsrouterServiceNotAuthorized.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			metricsRouterService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`CreateTarget - Create a target`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateTarget(createTargetOptions *CreateTargetOptions)`, func() {
			createTargetOptions := &metricsrouterv3.CreateTargetOptions{
				Name:           core.StringPtr("my-mr-target"),
				DestinationCRN: core.StringPtr("crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::"),
				Region:         core.StringPtr("us-south"),
			}

			target, response, err := metricsRouterService.CreateTarget(createTargetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(target).ToNot(BeNil())

			targetIDLink = *target.ID
			fmt.Fprintf(GinkgoWriter, "Saved targetIDLink value: %v\n", targetIDLink)
		})

		It(`Returns 403 when user is not authorized`, func() {
			createTargetOptions := &metricsrouterv3.CreateTargetOptions{
				Name:           core.StringPtr("my-mr-target"),
				DestinationCRN: core.StringPtr("crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::"),
				Region:         core.StringPtr("us-south"),
			}

			_, response, err := metricsrouterServiceNotAuthorized.CreateTarget(createTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})
	})

	Describe(`CreateRoute - Create a route`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateRoute(createRouteOptions *CreateRouteOptions)`, func() {
			targetIdentityModel := &metricsrouterv3.TargetIdentity{
				ID: &targetIDLink,
			}

			inclusionFilterPrototypeModel := &metricsrouterv3.InclusionFilterPrototype{
				Operand:  core.StringPtr("location"),
				Operator: core.StringPtr("is"),
				Values:   []string{"us-south"},
			}

			rulePrototypeModel := &metricsrouterv3.RulePrototype{
				Action:           core.StringPtr("send"),
				Targets:          []metricsrouterv3.TargetIdentity{*targetIdentityModel},
				InclusionFilters: []metricsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel},
			}

			createRouteOptions := &metricsrouterv3.CreateRouteOptions{
				Name:  core.StringPtr("my-route"),
				Rules: []metricsrouterv3.RulePrototype{*rulePrototypeModel},
			}

			route, response, err := metricsRouterService.CreateRoute(createRouteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(route).ToNot(BeNil())

			routeIDLink = *route.ID
			fmt.Fprintf(GinkgoWriter, "Saved routeIDLink value: %v\n", routeIDLink)
		})

		It(`CreateRoute(createRouteOptions *CreateRouteOptions) with in operator`, func() {
			targetIdentityModel := &metricsrouterv3.TargetIdentity{
				ID: &targetIDLink,
			}
			inclusionFilterPrototypeModel := &metricsrouterv3.InclusionFilterPrototype{
				Operand:  core.StringPtr("location"),
				Operator: core.StringPtr("in"),
				Values:   []string{"us-south", "us-east"},
			}
			rulePrototypeModel := &metricsrouterv3.RulePrototype{
				Action:           core.StringPtr("send"),
				Targets:          []metricsrouterv3.TargetIdentity{*targetIdentityModel},
				InclusionFilters: []metricsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel},
			}
			createRouteOptions := &metricsrouterv3.CreateRouteOptions{
				Name:  core.StringPtr("my-route-with-in"),
				Rules: []metricsrouterv3.RulePrototype{*rulePrototypeModel},
			}

			route, response, err := metricsRouterService.CreateRoute(createRouteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(route).ToNot(BeNil())

			routeIDLink1 = *route.ID
			fmt.Fprintf(GinkgoWriter, "Saved routeIDLink value: %v\n", routeIDLink1)
		})

		It(`CreateRoute fails when multiple values used with 'is' filter operator`, func() {
			targetIdentityModel := &metricsrouterv3.TargetIdentity{
				ID: &targetIDLink,
			}
			inclusionFilterPrototypeModel := &metricsrouterv3.InclusionFilterPrototype{
				Operand:  core.StringPtr("location"),
				Operator: core.StringPtr("is"),
				Values:   []string{"us-south", "us-east"},
			}
			rulePrototypeModel := &metricsrouterv3.RulePrototype{
				Action:           core.StringPtr("send"),
				Targets:          []metricsrouterv3.TargetIdentity{*targetIdentityModel},
				InclusionFilters: []metricsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel},
			}
			createRouteOptions := &metricsrouterv3.CreateRouteOptions{
				Name:  core.StringPtr("my-route-with-multiple-value-is-operator"),
				Rules: []metricsrouterv3.RulePrototype{*rulePrototypeModel},
			}

			route, response, err := metricsRouterService.CreateRoute(createRouteOptions)
			Expect(err).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(400))
			Expect(route).To(BeNil())
		})

		It(`CreateRoute fails when filter value length greater than limit`, func() {
			targetIdentityModel := &metricsrouterv3.TargetIdentity{
				ID: &targetIDLink,
			}
			inclusionFilterPrototypeModel := &metricsrouterv3.InclusionFilterPrototype{
				Operand:  core.StringPtr("location"),
				Operator: core.StringPtr("in"),
				Values:   []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v"},
			}
			rulePrototypeModel := &metricsrouterv3.RulePrototype{
				Action:           core.StringPtr("send"),
				Targets:          []metricsrouterv3.TargetIdentity{*targetIdentityModel},
				InclusionFilters: []metricsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel},
			}
			createRouteOptions := &metricsrouterv3.CreateRouteOptions{
				Name:  core.StringPtr("my-route-with-filter-length-more-than-limit"),
				Rules: []metricsrouterv3.RulePrototype{*rulePrototypeModel},
			}

			route, response, err := metricsRouterService.CreateRoute(createRouteOptions)
			Expect(err).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(400))
			Expect(route).To(BeNil())
		})

		It(`CreateRoute fails when filter value length equal to zero`, func() {
			targetIdentityModel := &metricsrouterv3.TargetIdentity{
				ID: &targetIDLink,
			}
			inclusionFilterPrototypeModel := &metricsrouterv3.InclusionFilterPrototype{
				Operand:  core.StringPtr("location"),
				Operator: core.StringPtr("in"),
				Values:   []string{""},
			}
			rulePrototypeModel := &metricsrouterv3.RulePrototype{
				Action:           core.StringPtr("send"),
				Targets:          []metricsrouterv3.TargetIdentity{*targetIdentityModel},
				InclusionFilters: []metricsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel},
			}
			createRouteOptions := &metricsrouterv3.CreateRouteOptions{
				Name:  core.StringPtr("my-route-with-zero-length-value"),
				Rules: []metricsrouterv3.RulePrototype{*rulePrototypeModel},
			}

			route, response, err := metricsRouterService.CreateRoute(createRouteOptions)
			Expect(err).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(400))
			Expect(route).To(BeNil())
		})

		It(`CreateRoute fails when filter length exceeded the limit`, func() {
			targetIdentityModel := &metricsrouterv3.TargetIdentity{
				ID: &targetIDLink,
			}
			inclusionFilterPrototypeModel := []metricsrouterv3.InclusionFilterPrototype{
				metricsrouterv3.InclusionFilterPrototype{
					Operand:  core.StringPtr("location"),
					Operator: core.StringPtr("in"),
					Values:   []string{"us-south"},
				},
				metricsrouterv3.InclusionFilterPrototype{
					Operand:  core.StringPtr("location"),
					Operator: core.StringPtr("in"),
					Values:   []string{"us-east"},
				},
				metricsrouterv3.InclusionFilterPrototype{
					Operand:  core.StringPtr("location"),
					Operator: core.StringPtr("in"),
					Values:   []string{"au-syd"},
				},
				metricsrouterv3.InclusionFilterPrototype{
					Operand:  core.StringPtr("location"),
					Operator: core.StringPtr("in"),
					Values:   []string{"eu-de"},
				},
				metricsrouterv3.InclusionFilterPrototype{
					Operand:  core.StringPtr("location"),
					Operator: core.StringPtr("in"),
					Values:   []string{"eu-gb"},
				},
				metricsrouterv3.InclusionFilterPrototype{
					Operand:  core.StringPtr("location"),
					Operator: core.StringPtr("in"),
					Values:   []string{"us-south"},
				},
				metricsrouterv3.InclusionFilterPrototype{
					Operand:  core.StringPtr("location"),
					Operator: core.StringPtr("in"),
					Values:   []string{"us-south"},
				},
				metricsrouterv3.InclusionFilterPrototype{
					Operand:  core.StringPtr("location"),
					Operator: core.StringPtr("in"),
					Values:   []string{"us-south"},
				},
			}

			rulePrototypeModel := &metricsrouterv3.RulePrototype{
				Action:           core.StringPtr("send"),
				Targets:          []metricsrouterv3.TargetIdentity{*targetIdentityModel},
				InclusionFilters: inclusionFilterPrototypeModel,
			}
			createRouteOptions := &metricsrouterv3.CreateRouteOptions{
				Name:  core.StringPtr("my-route-with-more-filter-length"),
				Rules: []metricsrouterv3.RulePrototype{*rulePrototypeModel},
			}

			route, response, err := metricsRouterService.CreateRoute(createRouteOptions)
			Expect(err).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(400))
			Expect(route).To(BeNil())
		})

		It(`Returns 403 when user is not authorized`, func() {
			targetIdentityModel := &metricsrouterv3.TargetIdentity{
				ID: &targetIDLink,
			}
			inclusionFilterPrototypeModel := &metricsrouterv3.InclusionFilterPrototype{
				Operand:  core.StringPtr("location"),
				Operator: core.StringPtr("is"),
				Values:   []string{"us-south"},
			}
			rulePrototypeModel := &metricsrouterv3.RulePrototype{
				Action:           core.StringPtr("send"),
				Targets:          []metricsrouterv3.TargetIdentity{*targetIdentityModel},
				InclusionFilters: []metricsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel},
			}
			createRouteOptions := &metricsrouterv3.CreateRouteOptions{
				Name:  core.StringPtr("my-unauthorized-route"),
				Rules: []metricsrouterv3.RulePrototype{*rulePrototypeModel},
			}

			_, response, err := metricsrouterServiceNotAuthorized.CreateRoute(createRouteOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 400 when input validation fails`, func() {
			targetIdentityModel := &metricsrouterv3.TargetIdentity{
				ID: core.StringPtr(notFoundTargetID),
			}
			inclusionFilterPrototypeModel := &metricsrouterv3.InclusionFilterPrototype{
				Operand:  core.StringPtr("location"),
				Operator: core.StringPtr("is"),
				Values:   []string{"us-south"},
			}
			rulePrototypeModel := &metricsrouterv3.RulePrototype{
				Action:           core.StringPtr("send"),
				Targets:          []metricsrouterv3.TargetIdentity{*targetIdentityModel},
				InclusionFilters: []metricsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel},
			}
			createRouteOptions := &metricsrouterv3.CreateRouteOptions{
				Name:  core.StringPtr("my-wrong-input-route"),
				Rules: []metricsrouterv3.RulePrototype{*rulePrototypeModel},
			}

			_, response, err := metricsRouterService.CreateRoute(createRouteOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})
	})

	Describe(`ListTargets - List targets`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListTargets(listTargetsOptions *ListTargetsOptions)`, func() {
			listTargetsOptions := &metricsrouterv3.ListTargetsOptions{}

			targetCollection, response, err := metricsRouterService.ListTargets(listTargetsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(targetCollection).ToNot(BeNil())
		})

		It(`Returns 403 when user is not authorized)`, func() {

			listTargetsOptions := &metricsrouterv3.ListTargetsOptions{}

			_, response, err := metricsrouterServiceNotAuthorized.ListTargets(listTargetsOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})
	})

	Describe(`GetTarget - Get details of a target`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetTarget(getTargetOptions *GetTargetOptions)`, func() {
			getTargetOptions := &metricsrouterv3.GetTargetOptions{
				ID: &targetIDLink,
			}

			target, response, err := metricsRouterService.GetTarget(getTargetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())
		})

		It(`Returns 403 when user is not authorized`, func() {
			getTargetOptions := &metricsrouterv3.GetTargetOptions{
				ID: &targetIDLink,
			}

			_, response, err := metricsrouterServiceNotAuthorized.GetTarget(getTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when target id is not found`, func() {

			getTargetOptions := &metricsrouterv3.GetTargetOptions{
				ID: core.StringPtr(notFoundTargetID),
			}

			_, response, err := metricsRouterService.GetTarget(getTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})
	})

	Describe(`UpdateTarget - Update a target`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateTarget(updateTargetOptions *UpdateTargetOptions)`, func() {
			updateTargetOptions := &metricsrouterv3.UpdateTargetOptions{
				ID:             &targetIDLink,
				Name:           core.StringPtr("my-mr-target"),
				DestinationCRN: core.StringPtr("crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::"),
			}

			target, response, err := metricsRouterService.UpdateTarget(updateTargetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())
		})

		It(`Returns 404 when target id is not found`, func() {

			updateTargetOptions := &metricsrouterv3.UpdateTargetOptions{
				ID:             core.StringPtr(notFoundTargetID),
				Name:           core.StringPtr("my-mr-target"),
				DestinationCRN: core.StringPtr("crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::"),
			}

			_, response, err := metricsRouterService.UpdateTarget(updateTargetOptions)
			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})
	})

	Describe(`ListRoutes - List routes`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListRoutes(listRoutesOptions *ListRoutesOptions)`, func() {
			listRoutesOptions := &metricsrouterv3.ListRoutesOptions{}

			routeCollection, response, err := metricsRouterService.ListRoutes(listRoutesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(routeCollection).ToNot(BeNil())
		})

		It(`Returns 403 when user is not authorized`, func() {

			listRoutesOptions := &metricsrouterv3.ListRoutesOptions{}

			_, response, err := metricsrouterServiceNotAuthorized.ListRoutes(listRoutesOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})
	})

	Describe(`GetRoute - Get details of a route`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetRoute(getRouteOptions *GetRouteOptions)`, func() {
			getRouteOptions := &metricsrouterv3.GetRouteOptions{
				ID: &routeIDLink,
			}

			route, response, err := metricsRouterService.GetRoute(getRouteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(route).ToNot(BeNil())
		})

		It(`Returns 403 when user is not authorized`, func() {
			getRouteOptions := &metricsrouterv3.GetRouteOptions{
				ID: &routeIDLink,
			}

			_, response, err := metricsrouterServiceNotAuthorized.GetRoute(getRouteOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when route id is not found`, func() {
			getRouteOptions := &metricsrouterv3.GetRouteOptions{
				ID: core.StringPtr(notFoundRouteID),
			}

			_, response, err := metricsRouterService.GetRoute(getRouteOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})
	})

	Describe(`UpdateRoute - Update a route`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateRoute(updateRouteOptions *UpdateRouteOptions)`, func() {
			targetIdentityModel := &metricsrouterv3.TargetIdentity{
				ID: &targetIDLink,
			}
			inclusionFilterPrototypeModel := &metricsrouterv3.InclusionFilterPrototype{
				Operand:  core.StringPtr("location"),
				Operator: core.StringPtr("is"),
				Values:   []string{"us-south"},
			}
			rulePrototypeModel := &metricsrouterv3.RulePrototype{
				Action:           core.StringPtr("send"),
				Targets:          []metricsrouterv3.TargetIdentity{*targetIdentityModel},
				InclusionFilters: []metricsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel},
			}
			updateRouteOptions := &metricsrouterv3.UpdateRouteOptions{
				ID:    &routeIDLink,
				Name:  core.StringPtr("my-route"),
				Rules: []metricsrouterv3.RulePrototype{*rulePrototypeModel},
			}
			route, response, err := metricsRouterService.UpdateRoute(updateRouteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(route).ToNot(BeNil())
		})

		It(`Returns 403 when user is not authorized`, func() {
			targetIdentityModel := &metricsrouterv3.TargetIdentity{
				ID: &targetIDLink,
			}
			inclusionFilterPrototypeModel := &metricsrouterv3.InclusionFilterPrototype{
				Operand:  core.StringPtr("location"),
				Operator: core.StringPtr("is"),
				Values:   []string{"us-south"},
			}
			rulePrototypeModel := &metricsrouterv3.RulePrototype{
				Action:           core.StringPtr("send"),
				Targets:          []metricsrouterv3.TargetIdentity{*targetIdentityModel},
				InclusionFilters: []metricsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel},
			}
			updateRouteOptions := &metricsrouterv3.UpdateRouteOptions{
				ID:    &routeIDLink,
				Name:  core.StringPtr("my-route"),
				Rules: []metricsrouterv3.RulePrototype{*rulePrototypeModel},
			}

			_, response, err := metricsrouterServiceNotAuthorized.UpdateRoute(updateRouteOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when route id is not found`, func() {
			targetIdentityModel := &metricsrouterv3.TargetIdentity{
				ID: &targetIDLink,
			}
			inclusionFilterPrototypeModel := &metricsrouterv3.InclusionFilterPrototype{
				Operand:  core.StringPtr("location"),
				Operator: core.StringPtr("is"),
				Values:   []string{"us-south"},
			}
			rulePrototypeModel := &metricsrouterv3.RulePrototype{
				Action:           core.StringPtr("send"),
				Targets:          []metricsrouterv3.TargetIdentity{*targetIdentityModel},
				InclusionFilters: []metricsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel},
			}
			updateRouteOptions := &metricsrouterv3.UpdateRouteOptions{
				ID:    core.StringPtr(notFoundRouteID),
				Name:  core.StringPtr("my-route"),
				Rules: []metricsrouterv3.RulePrototype{*rulePrototypeModel},
			}

			_, response, err := metricsRouterService.UpdateRoute(updateRouteOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})
	})

	Describe(`GetSettings - Get settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSettings(getSettingsOptions *GetSettingsOptions)`, func() {
			getSettingsOptions := &metricsrouterv3.GetSettingsOptions{}

			setting, response, err := metricsRouterService.GetSettings(getSettingsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(setting).ToNot(BeNil())
		})

		It(`Returns 403 when user is not authorized`, func() {
			getSettingsOptions := &metricsrouterv3.GetSettingsOptions{}

			_, response, err := metricsrouterServiceNotAuthorized.GetSettings(getSettingsOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})
	})

	Describe(`UpdateSettings - Modify settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateSettings(updateSettingsOptions *UpdateSettingsOptions)`, func() {
			targetIdentityModel := &metricsrouterv3.TargetIdentity{
				ID: &targetIDLink,
			}

			updateSettingsOptions := &metricsrouterv3.UpdateSettingsOptions{
				DefaultTargets:         []metricsrouterv3.TargetIdentity{*targetIdentityModel},
				PermittedTargetRegions: []string{"us-south"},
				PrimaryMetadataRegion:  core.StringPtr("us-south"),
				BackupMetadataRegion:   core.StringPtr("us-east"),
				PrivateAPIEndpointOnly: core.BoolPtr(false),
			}

			setting, response, err := metricsRouterService.UpdateSettings(updateSettingsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(setting).ToNot(BeNil())
		})

		It(`DeleteTarget will fail when target is a default target added in settings`, func() {
			deleteTargetOptions := &metricsrouterv3.DeleteTargetOptions{
				ID: &targetIDLink,
			}

			response, err := metricsRouterService.DeleteTarget(deleteTargetOptions)
			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(400))
		})

		It(`Returns 403 when user is not authorized`, func() {
			targetIdentityModel := &metricsrouterv3.TargetIdentity{
				ID: &targetIDLink,
			}

			updateSettingsOptions := &metricsrouterv3.UpdateSettingsOptions{
				DefaultTargets:         []metricsrouterv3.TargetIdentity{*targetIdentityModel},
				PermittedTargetRegions: []string{"us-south"},
				PrimaryMetadataRegion:  core.StringPtr("us-south"),
				BackupMetadataRegion:   core.StringPtr("us-east"),
				PrivateAPIEndpointOnly: core.BoolPtr(false),
			}

			_, response, err := metricsrouterServiceNotAuthorized.UpdateSettings(updateSettingsOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Removing default targets`, func() {
			updateSettingsOptions := &metricsrouterv3.UpdateSettingsOptions{
				DefaultTargets:         []metricsrouterv3.TargetIdentity{},
				PermittedTargetRegions: []string{"us-south"},
				PrimaryMetadataRegion:  core.StringPtr("us-south"),
				BackupMetadataRegion:   core.StringPtr("us-east"),
				PrivateAPIEndpointOnly: core.BoolPtr(false),
			}

			setting, response, err := metricsRouterService.UpdateSettings(updateSettingsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(setting).ToNot(BeNil())
		})

	})

	Describe(`DeleteRoute - Delete a route`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`Returns 403 when user is not authorized`, func() {
			deleteRouteOptions := &metricsrouterv3.DeleteRouteOptions{
				ID: &routeIDLink,
			}

			response, err := metricsrouterServiceNotAuthorized.DeleteRoute(deleteRouteOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when route id is not found`, func() {
			deleteRouteOptions := &metricsrouterv3.DeleteRouteOptions{
				ID: core.StringPtr(notFoundRouteID),
			}

			response, err := metricsRouterService.DeleteRoute(deleteRouteOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`DeleteRoute(deleteRouteOptions *DeleteRouteOptions)`, func() {
			deleteRouteOptions := &metricsrouterv3.DeleteRouteOptions{
				ID: &routeIDLink,
			}

			response, err := metricsRouterService.DeleteRoute(deleteRouteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})

		It(`DeleteRoute with in operator(deleteRouteOptions *DeleteRouteOptions)`, func() {
			deleteRouteOptions := &metricsrouterv3.DeleteRouteOptions{
				ID: &routeIDLink1,
			}

			response, err := metricsRouterService.DeleteRoute(deleteRouteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteTarget - Delete a target`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})

		It(`Returns 403 when user is not authorized`, func() {
			deleteTargetOptions := &metricsrouterv3.DeleteTargetOptions{
				ID: &targetIDLink,
			}

			response, err := metricsrouterServiceNotAuthorized.DeleteTarget(deleteTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(403))
		})

		It(`Returns 404 when target id is not found`, func() {

			deleteTargetOptions := &metricsrouterv3.DeleteTargetOptions{
				ID: core.StringPtr(notFoundTargetID),
			}

			response, err := metricsRouterService.DeleteTarget(deleteTargetOptions)

			Expect(err).NotTo(BeNil())
			Expect(response.StatusCode).To(Equal(404))
		})

		It(`DeleteTarget(deleteTargetOptions *DeleteTargetOptions)`, func() {
			deleteTargetOptions := &metricsrouterv3.DeleteTargetOptions{
				ID: &targetIDLink,
			}

			response, err := metricsRouterService.DeleteTarget(deleteTargetOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})

//
// Utility functions are declared in the unit test file
//
