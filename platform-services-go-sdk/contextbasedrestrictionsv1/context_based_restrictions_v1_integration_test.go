//go:build integration

/**
 * (C) Copyright IBM Corp. 2023.
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

package contextbasedrestrictionsv1_test

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/contextbasedrestrictionsv1"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the contextbasedrestrictionsv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`ContextBasedRestrictionsV1 Integration Tests`, func() {

	const externalConfigFile = "../context_based_restrictions_v1.env"

	const (
		NonExistentID = "1234567890abcdef1234567890abcdef"
		InvalidID     = "this_is_an_invalid_id"
	)

	var (
		err                             error
		contextBasedRestrictionsService *contextbasedrestrictionsv1.ContextBasedRestrictionsV1
		serviceURL                      string
		config                          map[string]string
		testAccountID                   string
		testServiceName                 string
		zoneID                          string
		zoneRev                         string
		ruleID                          string
		ruleRev                         string
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

			err := os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			if err != nil {
				Skip("Error setting IBM_CREDENTIALS_FILE environment variable, skipping tests: " + err.Error())
			}

			config, err = core.GetServiceProperties(contextbasedrestrictionsv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			testAccountID = config["TEST_ACCOUNT_ID"]
			if testAccountID == "" {
				Skip("Unable to load TEST_ACCOUNT_ID configuration property, skipping tests")
			}

			testServiceName = config["TEST_SERVICE_NAME"]
			if testServiceName == "" {
				Skip("Unable to load TEST_SERVICE_NAME configuration property, skipping tests")
			}

			fmt.Printf("\nService URL: %s\n", serviceURL)
			fmt.Printf("Test Account ID: %s\n", testAccountID)
			fmt.Printf("Test Service Name: %s\n", testServiceName)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			contextBasedRestrictionsServiceOptions := &contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{}
			contextBasedRestrictionsService, err = contextbasedrestrictionsv1.NewContextBasedRestrictionsV1UsingExternalConfig(contextBasedRestrictionsServiceOptions)

			Expect(err).To(BeNil())
			Expect(contextBasedRestrictionsService).ToNot(BeNil())
			Expect(contextBasedRestrictionsService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			contextBasedRestrictionsService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`CreateZone - Create a network zone`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateZone(createZoneOptions *CreateZoneOptions)`, func() {
			ipAddressModel := &contextbasedrestrictionsv1.AddressIPAddress{
				Type:  core.StringPtr("ipAddress"),
				Value: core.StringPtr("169.23.56.234"),
			}
			serviceRefAddressModel := &contextbasedrestrictionsv1.AddressServiceRef{
				Type: core.StringPtr(contextbasedrestrictionsv1.AddressServiceRefTypeServicerefConst),
				Ref: &contextbasedrestrictionsv1.ServiceRefValue{
					AccountID:   core.StringPtr(testAccountID),
					ServiceName: core.StringPtr("containers-kubernetes"),
					Location:    core.StringPtr("dal"),
				},
			}

			createZoneOptions := &contextbasedrestrictionsv1.CreateZoneOptions{
				Name:        core.StringPtr("an example of zone"),
				AccountID:   core.StringPtr(testAccountID),
				Description: core.StringPtr("this is an example of zone"),
				Addresses: []contextbasedrestrictionsv1.AddressIntf{
					ipAddressModel,
					serviceRefAddressModel,
				},
				TransactionID: getTransactionID(),
			}

			zone, response, err := contextBasedRestrictionsService.CreateZone(createZoneOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(zone).ToNot(BeNil())
			zoneID = *zone.ID
		})
	})

	Describe(`CreateZone - Create a zone with 'duplicated name' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateZone(createZoneOptions *CreateZoneOptions) with 'duplicated name' error (409)`, func() {
			addressModel := &contextbasedrestrictionsv1.AddressIPAddress{
				Type:  core.StringPtr("ipAddress"),
				Value: core.StringPtr("169.23.56.234"),
			}

			createZoneOptions := &contextbasedrestrictionsv1.CreateZoneOptions{
				Name:          core.StringPtr("an example of zone"),
				AccountID:     core.StringPtr(testAccountID),
				Description:   core.StringPtr("this is an example of zone"),
				Addresses:     []contextbasedrestrictionsv1.AddressIntf{addressModel},
				TransactionID: getTransactionID(),
			}

			zone, response, err := contextBasedRestrictionsService.CreateZone(createZoneOptions)

			Expect(err).To(Not(BeNil()))
			Expect(response.StatusCode).To(Equal(409))
			Expect(zone).To(BeNil())
		})
	})

	Describe(`CreateZone - Create a zone with 'invalid ip address format' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateZone(createZoneOptions *CreateZoneOptions) with 'invalid ip address format' error (400)`, func() {
			addressModel := &contextbasedrestrictionsv1.AddressIPAddress{
				Type:  core.StringPtr("ipAddress"),
				Value: core.StringPtr("169.23.56.234."),
			}

			createZoneOptions := &contextbasedrestrictionsv1.CreateZoneOptions{
				Name:          core.StringPtr("another example of zone"),
				AccountID:     core.StringPtr(testAccountID),
				Description:   core.StringPtr("this is another example of zone"),
				Addresses:     []contextbasedrestrictionsv1.AddressIntf{addressModel},
				TransactionID: getTransactionID(),
			}

			zone, response, err := contextBasedRestrictionsService.CreateZone(createZoneOptions)

			Expect(err).To(Not(BeNil()))
			Expect(response.StatusCode).To(Equal(400))
			Expect(zone).To(BeNil())
		})
	})

	Describe(`ListZones - List zones`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListZones(listZonesOptions *ListZonesOptions)`, func() {
			listZonesOptions := &contextbasedrestrictionsv1.ListZonesOptions{
				AccountID:     core.StringPtr(testAccountID),
				TransactionID: getTransactionID(),
			}

			zoneList, response, err := contextBasedRestrictionsService.ListZones(listZonesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(zoneList).ToNot(BeNil())
		})
	})

	Describe(`ListZones - List zones with 'missing AccountID parameter' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListZones(listZonesOptions *ListZonesOptions) with 'missing AccountID parameter' error`, func() {
			listZonesOptions := &contextbasedrestrictionsv1.ListZonesOptions{
				TransactionID: getTransactionID(),
			}

			zoneList, response, err := contextBasedRestrictionsService.ListZones(listZonesOptions)

			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(ContainSubstring("Field validation for 'AccountID' failed"))
			Expect(response).To(BeNil())
			Expect(zoneList).To(BeNil())
		})
	})

	Describe(`ListZones - List zones with 'invalid AccountID parameter' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListZones(listZonesOptions *ListZonesOptions) with 'invalid AccountID parameter' error (400)`, func() {
			listZonesOptions := &contextbasedrestrictionsv1.ListZonesOptions{
				AccountID:     core.StringPtr(InvalidID),
				TransactionID: getTransactionID(),
			}

			zoneList, response, err := contextBasedRestrictionsService.ListZones(listZonesOptions)

			Expect(err).To(Not(BeNil()))
			Expect(response.StatusCode).To(Equal(400))
			Expect(zoneList).To(BeNil())
		})
	})

	Describe(`GetZone - Get the specified zone`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetZone(getZoneOptions *GetZoneOptions)`, func() {
			getZoneOptions := &contextbasedrestrictionsv1.GetZoneOptions{
				ZoneID:        core.StringPtr(zoneID),
				TransactionID: getTransactionID(),
			}

			zone, response, err := contextBasedRestrictionsService.GetZone(getZoneOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(zone).ToNot(BeNil())
			zoneRev = response.Headers.Get("Etag")
		})
	})

	Describe(`GetZone - Get zone with 'missing required ZoneID parameter' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetZone(getZoneOptions *GetZoneOptions) with 'missing required ZoneID parameter' error`, func() {
			getZoneOptions := &contextbasedrestrictionsv1.GetZoneOptions{
				TransactionID: getTransactionID(),
			}

			zone, response, err := contextBasedRestrictionsService.GetZone(getZoneOptions)

			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(ContainSubstring("Field validation for 'ZoneID' failed"))
			Expect(response).To(BeNil())
			Expect(zone).To(BeNil())
		})
	})

	Describe(`GetZone - Get zone with 'zone not found' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetZone(getZoneOptions *GetZoneOptions) with 'zone not found' error (404)`, func() {
			getZoneOptions := &contextbasedrestrictionsv1.GetZoneOptions{
				ZoneID:        core.StringPtr(NonExistentID),
				TransactionID: getTransactionID(),
			}

			zone, response, err := contextBasedRestrictionsService.GetZone(getZoneOptions)

			Expect(err).To(Not(BeNil()))
			Expect(response.StatusCode).To(Equal(404))
			Expect(zone).To(BeNil())

		})
	})

	Describe(`ReplaceZone - Update the specified zone`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceZone(replaceZoneOptions *ReplaceZoneOptions)`, func() {
			addressModel := &contextbasedrestrictionsv1.AddressIPAddress{
				Type:  core.StringPtr("ipAddress"),
				Value: core.StringPtr("169.23.56.234"),
			}

			replaceZoneOptions := &contextbasedrestrictionsv1.ReplaceZoneOptions{
				ZoneID:        core.StringPtr(zoneID),
				IfMatch:       core.StringPtr(zoneRev),
				Name:          core.StringPtr("an example of updated zone"),
				AccountID:     core.StringPtr(testAccountID),
				Description:   core.StringPtr("this is an example of updated zone"),
				Addresses:     []contextbasedrestrictionsv1.AddressIntf{addressModel},
				TransactionID: getTransactionID(),
			}

			zone, response, err := contextBasedRestrictionsService.ReplaceZone(replaceZoneOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(zone).ToNot(BeNil())
		})
	})

	Describe(`ReplaceZone - Update zone with 'missing required IfMatch parameter' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceZone(replaceZoneOptions *ReplaceZoneOptions) with 'missing required IfMatch parameter' error (400)`, func() {
			addressModel := &contextbasedrestrictionsv1.AddressIPAddress{
				Type:  core.StringPtr("ipAddress"),
				Value: core.StringPtr("169.23.56.234"),
			}

			replaceZoneOptions := &contextbasedrestrictionsv1.ReplaceZoneOptions{
				ZoneID:        core.StringPtr(zoneID),
				Name:          core.StringPtr("an example of zone"),
				AccountID:     core.StringPtr(testAccountID),
				Description:   core.StringPtr("this is an example of zone"),
				Addresses:     []contextbasedrestrictionsv1.AddressIntf{addressModel},
				TransactionID: getTransactionID(),
			}

			zone, response, err := contextBasedRestrictionsService.ReplaceZone(replaceZoneOptions)

			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(ContainSubstring("Field validation for 'IfMatch' failed"))
			Expect(response).To(BeNil())
			Expect(zone).To(BeNil())
		})
	})

	Describe(`ReplaceZone - Update zone with 'zone not found' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceZone(replaceZoneOptions *ReplaceZoneOptions) with 'zone not found' error (404)`, func() {
			addressModel := &contextbasedrestrictionsv1.AddressIPAddress{
				Type:  core.StringPtr("ipAddress"),
				Value: core.StringPtr("169.23.56.234"),
			}

			replaceZoneOptions := &contextbasedrestrictionsv1.ReplaceZoneOptions{
				ZoneID:        core.StringPtr(NonExistentID),
				IfMatch:       core.StringPtr("abc"),
				Name:          core.StringPtr("an example of zone"),
				AccountID:     core.StringPtr(testAccountID),
				Description:   core.StringPtr("this is an example of zone"),
				Addresses:     []contextbasedrestrictionsv1.AddressIntf{addressModel},
				TransactionID: getTransactionID(),
			}

			zone, response, err := contextBasedRestrictionsService.ReplaceZone(replaceZoneOptions)

			Expect(err).To(Not(BeNil()))
			Expect(response.StatusCode).To(Equal(404))
			Expect(zone).To(BeNil())
		})
	})

	Describe(`ReplaceZone - Update zone with 'invalid IfMath parameter' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceZone(replaceZoneOptions *ReplaceZoneOptions) with 'invalid IfMath parameter' error (412)`, func() {
			addressModel := &contextbasedrestrictionsv1.AddressIPAddress{
				Type:  core.StringPtr("ipAddress"),
				Value: core.StringPtr("169.23.56.234"),
			}

			replaceZoneOptions := &contextbasedrestrictionsv1.ReplaceZoneOptions{
				ZoneID:        core.StringPtr(zoneID),
				IfMatch:       core.StringPtr("abc"),
				Name:          core.StringPtr("an example of zone"),
				AccountID:     core.StringPtr(testAccountID),
				Description:   core.StringPtr("this is an example of zone"),
				Addresses:     []contextbasedrestrictionsv1.AddressIntf{addressModel},
				TransactionID: getTransactionID(),
			}

			zone, response, err := contextBasedRestrictionsService.ReplaceZone(replaceZoneOptions)

			Expect(err).To(Not(BeNil()))
			Expect(response.StatusCode).To(Equal(412))
			Expect(zone).To(BeNil())
		})
	})

	Describe(`ListAvailableServiceRefTargets - List available service reference targets`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListAvailableServiceRefTargets(listAvailableServiceRefTargetsOptions *ListAvailableServiceRefTargetsOptions)`, func() {
			listAvailableServiceRefTargetsOptions := &contextbasedrestrictionsv1.ListAvailableServicerefTargetsOptions{
				Type: core.StringPtr(contextbasedrestrictionsv1.ListAvailableServicerefTargetsOptionsTypeAllConst),
			}

			serviceRefTargetList, response, err := contextBasedRestrictionsService.ListAvailableServicerefTargets(listAvailableServiceRefTargetsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceRefTargetList).ToNot(BeNil())
		})
	})

	Describe(`ListAvailableServiceRefTargets - List available service reference targets with 'invalid type parameter' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListAvailableServiceRefTargets(listAvailableServiceRefTargetsOptions *ListAvailableServiceRefTargetsOptions) with 'invalid type parameter' error (400)`, func() {
			listAvailableServiceRefTargetsOptions := &contextbasedrestrictionsv1.ListAvailableServicerefTargetsOptions{
				Type: core.StringPtr("invalid-type"),
			}

			serviceRefTargetList, response, err := contextBasedRestrictionsService.ListAvailableServicerefTargets(listAvailableServiceRefTargetsOptions)

			Expect(err).To(Not(BeNil()))
			Expect(response.StatusCode).To(Equal(400))
			Expect(serviceRefTargetList).To(BeNil())
		})
	})

	Describe(`GetServicerefTarget - Get service reference target for a specified service name`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetServicerefTarget(getServicerefTargetOptions *GetServicerefTargetOptions)`, func() {
			getServicerefTargetOptions := &contextbasedrestrictionsv1.GetServicerefTargetOptions{
				ServiceName: core.StringPtr("containers-kubernetes"),
			}

			serviceRefTarget, response, err := contextBasedRestrictionsService.GetServicerefTarget(getServicerefTargetOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceRefTarget).ToNot(BeNil())
		})
	})

	Describe(`GetServicerefTarget - Get service reference target for a specified service name with 'service_not_found' error `, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetServicerefTarget(getServicerefTargetOptions *GetServicerefTargetOptions)`, func() {
			getServicerefTargetOptions := &contextbasedrestrictionsv1.GetServicerefTargetOptions{
				ServiceName: core.StringPtr("invalid-service"),
			}

			serviceRefTarget, response, err := contextBasedRestrictionsService.GetServicerefTarget(getServicerefTargetOptions)

			Expect(err).To(Not(BeNil()))
			Expect(response.StatusCode).To(Equal(404))
			Expect(serviceRefTarget).To(BeNil())
		})
	})

	Describe(`CreateRule - Create a rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateRule(createRuleOptions *CreateRuleOptions)`, func() {
			ruleContextAttributeModel := &contextbasedrestrictionsv1.RuleContextAttribute{
				Name:  core.StringPtr("networkZoneId"),
				Value: core.StringPtr(zoneID),
			}

			ruleContextModel := &contextbasedrestrictionsv1.RuleContext{
				Attributes: []contextbasedrestrictionsv1.RuleContextAttribute{*ruleContextAttributeModel},
			}

			resourceModel := &contextbasedrestrictionsv1.Resource{
				Attributes: []contextbasedrestrictionsv1.ResourceAttribute{
					{
						Name:  core.StringPtr("accountId"),
						Value: core.StringPtr(testAccountID),
					},
					{
						Name:  core.StringPtr("serviceName"),
						Value: core.StringPtr(testServiceName),
					},
				},
				Tags: []contextbasedrestrictionsv1.ResourceTagAttribute{
					{
						Name:  core.StringPtr("tagName"),
						Value: core.StringPtr("tagValue"),
					},
				},
			}

			createRuleOptions := &contextbasedrestrictionsv1.CreateRuleOptions{
				Description:     core.StringPtr("this is an example of rule"),
				Contexts:        []contextbasedrestrictionsv1.RuleContext{*ruleContextModel},
				Resources:       []contextbasedrestrictionsv1.Resource{*resourceModel},
				EnforcementMode: core.StringPtr(contextbasedrestrictionsv1.CreateRuleOptionsEnforcementModeEnabledConst),
				TransactionID:   getTransactionID(),
			}

			rule, response, err := contextBasedRestrictionsService.CreateRule(createRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(rule).ToNot(BeNil())
			ruleID = *rule.ID
		})
	})

	Describe(`CreateRule - Create a rule with an API type`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateRule(createRuleOptions *CreateRuleOptions)`, func() {
			ruleContextAttributeModel := &contextbasedrestrictionsv1.RuleContextAttribute{
				Name:  core.StringPtr("networkZoneId"),
				Value: core.StringPtr(zoneID),
			}

			ruleContextModel := &contextbasedrestrictionsv1.RuleContext{
				Attributes: []contextbasedrestrictionsv1.RuleContextAttribute{*ruleContextAttributeModel},
			}

			resourceModel := &contextbasedrestrictionsv1.Resource{
				Attributes: []contextbasedrestrictionsv1.ResourceAttribute{
					{
						Name:  core.StringPtr("accountId"),
						Value: core.StringPtr(testAccountID),
					},
					{
						Name:  core.StringPtr("serviceName"),
						Value: core.StringPtr("containers-kubernetes"),
					},
				},
			}

			operationsModel := &contextbasedrestrictionsv1.NewRuleOperations{
				APITypes: []contextbasedrestrictionsv1.NewRuleOperationsAPITypesItem{
					{APITypeID: core.StringPtr("crn:v1:bluemix:public:containers-kubernetes::::api-type:management")},
				},
			}

			createRuleOptions := &contextbasedrestrictionsv1.CreateRuleOptions{
				Description:     core.StringPtr("this is an example of rule"),
				Contexts:        []contextbasedrestrictionsv1.RuleContext{*ruleContextModel},
				Resources:       []contextbasedrestrictionsv1.Resource{*resourceModel},
				Operations:      operationsModel,
				EnforcementMode: core.StringPtr(contextbasedrestrictionsv1.CreateRuleOptionsEnforcementModeEnabledConst),
				TransactionID:   getTransactionID(),
			}

			rule, response, err := contextBasedRestrictionsService.CreateRule(createRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(rule).ToNot(BeNil())

			// cleanup
			deleteRuleOptions := &contextbasedrestrictionsv1.DeleteRuleOptions{
				RuleID:        rule.ID,
				TransactionID: getTransactionID(),
			}

			response, err = contextBasedRestrictionsService.DeleteRule(deleteRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`CreateRule - Create a rule with 'service not cbr enabled' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateRule(createRuleOptions *CreateRuleOptions) with 'service not cbr enabled' error (400)`, func() {
			ruleContextAttributeModel := &contextbasedrestrictionsv1.RuleContextAttribute{
				Name:  core.StringPtr("networkZoneId"),
				Value: core.StringPtr(zoneID),
			}

			ruleContextModel := &contextbasedrestrictionsv1.RuleContext{
				Attributes: []contextbasedrestrictionsv1.RuleContextAttribute{*ruleContextAttributeModel},
			}

			resourceModel := &contextbasedrestrictionsv1.Resource{
				Attributes: []contextbasedrestrictionsv1.ResourceAttribute{
					{
						Name:  core.StringPtr("accountId"),
						Value: core.StringPtr(testAccountID),
					},
					{
						Name:  core.StringPtr("serviceName"),
						Value: core.StringPtr("cbr-not-enabled"),
					},
				},
				Tags: []contextbasedrestrictionsv1.ResourceTagAttribute{
					{
						Name:  core.StringPtr("tagName"),
						Value: core.StringPtr("tagValue"),
					},
				},
			}

			createRuleOptions := &contextbasedrestrictionsv1.CreateRuleOptions{
				Description:     core.StringPtr("this is an example of rule"),
				Contexts:        []contextbasedrestrictionsv1.RuleContext{*ruleContextModel},
				Resources:       []contextbasedrestrictionsv1.Resource{*resourceModel},
				EnforcementMode: core.StringPtr(contextbasedrestrictionsv1.CreateRuleOptionsEnforcementModeReportConst),
				TransactionID:   getTransactionID(),
			}

			rule, response, err := contextBasedRestrictionsService.CreateRule(createRuleOptions)

			Expect(err).To(Not(BeNil()))
			Expect(response.StatusCode).To(Equal(400))
			Expect(rule).To(BeNil())
		})
	})

	Describe(`ListRules - List rules`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListRules(listRulesOptions *ListRulesOptions)`, func() {
			listRulesOptions := &contextbasedrestrictionsv1.ListRulesOptions{
				AccountID:     core.StringPtr(testAccountID),
				TransactionID: getTransactionID(),
			}

			ruleList, response, err := contextBasedRestrictionsService.ListRules(listRulesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ruleList).ToNot(BeNil())
		})
	})

	Describe(`ListRule - List a rule with a valid service_group_id`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		//create rule with service_group_id
		It(`CreateRule(createRuleOptions *CreateRuleOptions) with a valid service_group_id (201)`, func() {
			ruleContextAttributeModel := &contextbasedrestrictionsv1.RuleContextAttribute{
				Name:  core.StringPtr("networkZoneId"),
				Value: core.StringPtr(zoneID),
			}

			ruleContextModel := &contextbasedrestrictionsv1.RuleContext{
				Attributes: []contextbasedrestrictionsv1.RuleContextAttribute{*ruleContextAttributeModel},
			}

			resourceModel := &contextbasedrestrictionsv1.Resource{
				Attributes: []contextbasedrestrictionsv1.ResourceAttribute{
					{
						Name:  core.StringPtr("accountId"),
						Value: core.StringPtr(testAccountID),
					},
					{
						Name:  core.StringPtr("service_group_id"),
						Value: core.StringPtr("IAM"),
					},
				},
			}

			createRuleOptions := &contextbasedrestrictionsv1.CreateRuleOptions{
				Description:     core.StringPtr("this is an example of rule with a service_group_id"),
				Contexts:        []contextbasedrestrictionsv1.RuleContext{*ruleContextModel},
				Resources:       []contextbasedrestrictionsv1.Resource{*resourceModel},
				EnforcementMode: core.StringPtr(contextbasedrestrictionsv1.CreateRuleOptionsEnforcementModeEnabledConst),
				TransactionID:   getTransactionID(),
			}

			rule, response, err := contextBasedRestrictionsService.CreateRule(createRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(rule).ToNot(BeNil())

			//list rule with service_group_id
			listRulesOptions := &contextbasedrestrictionsv1.ListRulesOptions{
				AccountID:      core.StringPtr(testAccountID),
				ServiceGroupID: core.StringPtr("IAM"),
				TransactionID:  getTransactionID(),
			}

			ruleList, response, err := contextBasedRestrictionsService.ListRules(listRulesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(*ruleList.Count).To(Equal(int64(1)))
			Expect(*ruleList.Rules[0].ID).To(Equal(*rule.ID))

			// cleanup
			deleteRuleOptions := &contextbasedrestrictionsv1.DeleteRuleOptions{
				RuleID:        rule.ID,
				TransactionID: getTransactionID(),
			}

			response, err = contextBasedRestrictionsService.DeleteRule(deleteRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`ListRules - List rules with 'missing required AccountID parameter' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListRules(listRulesOptions *ListRulesOptions) with 'missing required AccountID parameter' error (400)`, func() {
			listRulesOptions := &contextbasedrestrictionsv1.ListRulesOptions{
				TransactionID: getTransactionID(),
			}

			ruleList, response, err := contextBasedRestrictionsService.ListRules(listRulesOptions)

			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(ContainSubstring("Field validation for 'AccountID' failed"))
			Expect(response).To(BeNil())
			Expect(ruleList).To(BeNil())
		})
	})

	Describe(`ListRules - List rules with 'invalid AccountID parameter' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListRules(listRulesOptions *ListRulesOptions) with 'invalid AccountID parameter' error (400)`, func() {
			listRulesOptions := &contextbasedrestrictionsv1.ListRulesOptions{
				AccountID:     core.StringPtr(InvalidID),
				TransactionID: getTransactionID(),
			}

			ruleList, response, err := contextBasedRestrictionsService.ListRules(listRulesOptions)

			Expect(err).To(Not(BeNil()))
			Expect(response.StatusCode).To(Equal(400))
			Expect(ruleList).To(BeNil())
		})
	})

	Describe(`GetRule - Get the specified rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetRule(getRuleOptions *GetRuleOptions)`, func() {
			getRuleOptions := &contextbasedrestrictionsv1.GetRuleOptions{
				RuleID:        core.StringPtr(ruleID),
				TransactionID: getTransactionID(),
			}

			rule, response, err := contextBasedRestrictionsService.GetRule(getRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())
			ruleRev = response.Headers.Get("Etag")
		})
	})

	Describe(`GetRule - Get rule with 'missing required RuleID parameter' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetRule(getRuleOptions *GetRuleOptions) with 'missing required RuleID parameter' error`, func() {
			getRuleOptions := &contextbasedrestrictionsv1.GetRuleOptions{
				TransactionID: getTransactionID(),
			}

			rule, response, err := contextBasedRestrictionsService.GetRule(getRuleOptions)

			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(ContainSubstring("Field validation for 'RuleID' failed"))
			Expect(response).To(BeNil())
			Expect(rule).To(BeNil())
		})
	})

	Describe(`GetRule - Get rule with 'rule not found' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetRule(getRuleOptions *GetRuleOptions) with 'rule not found' error (404)`, func() {
			getRuleOptions := &contextbasedrestrictionsv1.GetRuleOptions{
				RuleID:        core.StringPtr(NonExistentID),
				TransactionID: getTransactionID(),
			}

			rule, response, err := contextBasedRestrictionsService.GetRule(getRuleOptions)

			Expect(err).To(Not(BeNil()))
			Expect(response.StatusCode).To(Equal(404))
			Expect(rule).To(BeNil())
		})
	})

	Describe(`ReplaceRule - Update the specified rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceRule(replaceRuleOptions *ReplaceRuleOptions)`, func() {
			ruleContextAttributeModel := &contextbasedrestrictionsv1.RuleContextAttribute{
				Name:  core.StringPtr("networkZoneId"),
				Value: core.StringPtr(zoneID),
			}

			ruleContextModel := &contextbasedrestrictionsv1.RuleContext{
				Attributes: []contextbasedrestrictionsv1.RuleContextAttribute{*ruleContextAttributeModel},
			}

			resourceModel := &contextbasedrestrictionsv1.Resource{
				Attributes: []contextbasedrestrictionsv1.ResourceAttribute{
					{
						Name:  core.StringPtr("accountId"),
						Value: core.StringPtr(testAccountID),
					},
					{
						Name:  core.StringPtr("serviceName"),
						Value: core.StringPtr(testServiceName),
					},
				},
				Tags: []contextbasedrestrictionsv1.ResourceTagAttribute{
					{
						Name:  core.StringPtr("tagName"),
						Value: core.StringPtr("updatedTagValue"),
					},
				},
			}

			replaceRuleOptions := &contextbasedrestrictionsv1.ReplaceRuleOptions{
				RuleID:          core.StringPtr(ruleID),
				IfMatch:         core.StringPtr(ruleRev),
				Description:     core.StringPtr("this is an example of updated rule"),
				Contexts:        []contextbasedrestrictionsv1.RuleContext{*ruleContextModel},
				Resources:       []contextbasedrestrictionsv1.Resource{*resourceModel},
				EnforcementMode: core.StringPtr(contextbasedrestrictionsv1.ReplaceRuleOptionsEnforcementModeDisabledConst),
				TransactionID:   getTransactionID(),
			}

			rule, response, err := contextBasedRestrictionsService.ReplaceRule(replaceRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())
		})
	})

	Describe(`ReplaceRule - Update rule with 'missing required IfMatch parameter' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceRule(replaceRuleOptions *ReplaceRuleOptions) with 'missing required IfMatch parameter' error (400)`, func() {
			ruleContextAttributeModel := &contextbasedrestrictionsv1.RuleContextAttribute{
				Name:  core.StringPtr("networkZoneId"),
				Value: core.StringPtr(zoneID),
			}

			ruleContextModel := &contextbasedrestrictionsv1.RuleContext{
				Attributes: []contextbasedrestrictionsv1.RuleContextAttribute{*ruleContextAttributeModel},
			}

			resourceModel := &contextbasedrestrictionsv1.Resource{
				Attributes: []contextbasedrestrictionsv1.ResourceAttribute{
					{
						Name:  core.StringPtr("accountId"),
						Value: core.StringPtr(testAccountID),
					},
					{
						Name:  core.StringPtr("serviceName"),
						Value: core.StringPtr(testServiceName),
					},
				},
				Tags: []contextbasedrestrictionsv1.ResourceTagAttribute{
					{
						Name:  core.StringPtr("tagName"),
						Value: core.StringPtr("updatedTagValue"),
					},
				},
			}

			replaceRuleOptions := &contextbasedrestrictionsv1.ReplaceRuleOptions{
				RuleID:          core.StringPtr(ruleID),
				Description:     core.StringPtr("this is an example of rule"),
				Contexts:        []contextbasedrestrictionsv1.RuleContext{*ruleContextModel},
				Resources:       []contextbasedrestrictionsv1.Resource{*resourceModel},
				EnforcementMode: core.StringPtr(contextbasedrestrictionsv1.ReplaceRuleOptionsEnforcementModeDisabledConst),
				TransactionID:   getTransactionID(),
			}

			rule, response, err := contextBasedRestrictionsService.ReplaceRule(replaceRuleOptions)

			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(ContainSubstring("Field validation for 'IfMatch' failed"))
			Expect(response).To(BeNil())
			Expect(rule).To(BeNil())
		})
	})

	Describe(`ReplaceRule - Update rule with 'rule not found' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceRule(replaceRuleOptions *ReplaceRuleOptions) with 'rule not found' error (404)`, func() {
			ruleContextAttributeModel := &contextbasedrestrictionsv1.RuleContextAttribute{
				Name:  core.StringPtr("networkZoneId"),
				Value: core.StringPtr(zoneID),
			}

			ruleContextModel := &contextbasedrestrictionsv1.RuleContext{
				Attributes: []contextbasedrestrictionsv1.RuleContextAttribute{*ruleContextAttributeModel},
			}

			resourceModel := &contextbasedrestrictionsv1.Resource{
				Attributes: []contextbasedrestrictionsv1.ResourceAttribute{
					{
						Name:  core.StringPtr("accountId"),
						Value: core.StringPtr(testAccountID),
					},
					{
						Name:  core.StringPtr("serviceName"),
						Value: core.StringPtr(testServiceName),
					},
				},
				Tags: []contextbasedrestrictionsv1.ResourceTagAttribute{
					{
						Name:  core.StringPtr("tagName"),
						Value: core.StringPtr("updatedTagValue"),
					},
				},
			}

			replaceRuleOptions := &contextbasedrestrictionsv1.ReplaceRuleOptions{
				RuleID:          core.StringPtr(NonExistentID),
				IfMatch:         core.StringPtr("abc"),
				Description:     core.StringPtr("this is an example of rule"),
				Contexts:        []contextbasedrestrictionsv1.RuleContext{*ruleContextModel},
				Resources:       []contextbasedrestrictionsv1.Resource{*resourceModel},
				EnforcementMode: core.StringPtr(contextbasedrestrictionsv1.ReplaceRuleOptionsEnforcementModeDisabledConst),
				TransactionID:   getTransactionID(),
			}

			rule, response, err := contextBasedRestrictionsService.ReplaceRule(replaceRuleOptions)

			Expect(err).To(Not(BeNil()))
			Expect(response.StatusCode).To(Equal(404))
			Expect(rule).To(BeNil())
		})
	})

	Describe(`ReplaceRule - Update rule with 'invalid IfMatch parameter' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceRule(replaceRuleOptions *ReplaceRuleOptions) with 'invalid IfMatch parameter' error (412)`, func() {
			ruleContextAttributeModel := &contextbasedrestrictionsv1.RuleContextAttribute{
				Name:  core.StringPtr("networkZoneId"),
				Value: core.StringPtr(zoneID),
			}

			ruleContextModel := &contextbasedrestrictionsv1.RuleContext{
				Attributes: []contextbasedrestrictionsv1.RuleContextAttribute{*ruleContextAttributeModel},
			}

			resourceModel := &contextbasedrestrictionsv1.Resource{
				Attributes: []contextbasedrestrictionsv1.ResourceAttribute{
					{
						Name:  core.StringPtr("accountId"),
						Value: core.StringPtr(testAccountID),
					},
					{
						Name:  core.StringPtr("serviceName"),
						Value: core.StringPtr(testServiceName),
					},
				},
				Tags: []contextbasedrestrictionsv1.ResourceTagAttribute{
					{
						Name:  core.StringPtr("tagName"),
						Value: core.StringPtr("updatedTagValue"),
					},
				},
			}

			replaceRuleOptions := &contextbasedrestrictionsv1.ReplaceRuleOptions{
				RuleID:          core.StringPtr(ruleID),
				IfMatch:         core.StringPtr("abc"),
				Description:     core.StringPtr("this is an example of rule"),
				Contexts:        []contextbasedrestrictionsv1.RuleContext{*ruleContextModel},
				Resources:       []contextbasedrestrictionsv1.Resource{*resourceModel},
				EnforcementMode: core.StringPtr(contextbasedrestrictionsv1.ReplaceRuleOptionsEnforcementModeDisabledConst),
				TransactionID:   getTransactionID(),
			}

			rule, response, err := contextBasedRestrictionsService.ReplaceRule(replaceRuleOptions)

			Expect(err).To(Not(BeNil()))
			Expect(response.StatusCode).To(Equal(412))
			Expect(rule).To(BeNil())
		})
	})

	Describe(`GetAccountSettings - Get the specified account settings`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAccountSettings(getAccountSettingsOptions *GetAccountSettingsOptions)`, func() {
			getAccountSettingsOptions := &contextbasedrestrictionsv1.GetAccountSettingsOptions{
				AccountID:     core.StringPtr(testAccountID),
				TransactionID: getTransactionID(),
			}

			accountSettings, response, err := contextBasedRestrictionsService.GetAccountSettings(getAccountSettingsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSettings).ToNot(BeNil())
		})
	})

	Describe(`GetAccountSettings - Get account settings with 'invalid AccountID parameter' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetAccountSettings(getAccountSettingsOptions *GetAccountSettingsOptions) with 'invalid AccountID parameter' error (400)`, func() {
			getAccountSettingsOptions := &contextbasedrestrictionsv1.GetAccountSettingsOptions{
				AccountID:     core.StringPtr(InvalidID),
				TransactionID: getTransactionID(),
			}

			accountSettings, response, err := contextBasedRestrictionsService.GetAccountSettings(getAccountSettingsOptions)

			Expect(err).To(Not(BeNil()))
			Expect(response.StatusCode).To(Equal(400))
			Expect(accountSettings).To(BeNil())
		})
	})

	Describe(`ListAvailableServiceOperations - List available service operations with Service Name`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListAvailableServiceOperations(listAvailableServiceOperationsOptions *ListAvailableServiceOperationsOptions)`, func() {
			listAvailableServiceOperationsOptions := &contextbasedrestrictionsv1.ListAvailableServiceOperationsOptions{
				ServiceName:   core.StringPtr("containers-kubernetes"),
				TransactionID: getTransactionID(),
			}

			operationsList, response, err := contextBasedRestrictionsService.ListAvailableServiceOperations(listAvailableServiceOperationsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(operationsList).ToNot(BeNil())
			for _, apiType := range operationsList.APITypes {
				Expect(*apiType.Type).ToNot(BeEmpty())
			}
		})
	})
	Describe(`ListAvailableServiceOperations - List available service operations with Service Group ID`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListAvailableServiceOperations(listAvailableServiceOperationsOptions *ListAvailableServiceOperationsOptions)`, func() {
			listAvailableServiceOperationsOptions := &contextbasedrestrictionsv1.ListAvailableServiceOperationsOptions{
				ServiceGroupID: core.StringPtr("IAM"),
				TransactionID:  getTransactionID(),
			}

			operationsList, response, err := contextBasedRestrictionsService.ListAvailableServiceOperations(listAvailableServiceOperationsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(operationsList).ToNot(BeNil())
			for _, apiType := range operationsList.APITypes {
				Expect(*apiType.Type).ToNot(BeEmpty())
			}
		})
	})
	Describe(`ListAvailableServiceOperations - List available service operations for subresource`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListAvailableServiceOperations(listAvailableServiceOperationsOptions *ListAvailableServiceOperationsOptions)`, func() {
			listAvailableServiceOperationsOptions := &contextbasedrestrictionsv1.ListAvailableServiceOperationsOptions{
				ServiceName:   core.StringPtr("iam-access-management"),
				ResourceType:  core.StringPtr("customRole"),
				TransactionID: getTransactionID(),
			}

			operationsList, response, err := contextBasedRestrictionsService.ListAvailableServiceOperations(listAvailableServiceOperationsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(operationsList).ToNot(BeNil())
			for _, apiType := range operationsList.APITypes {
				Expect(*apiType.Type).ToNot(BeEmpty())
			}
		})
	})
	Describe(`ListAvailableServiceOperations - List available service operations with 'mutually exclusive parameters' Error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListAvailableServiceOperations(listAvailableServiceOperationsOptions *ListAvailableServiceOperationsOptions)`, func() {
			listAvailableServiceOperationsOptions := &contextbasedrestrictionsv1.ListAvailableServiceOperationsOptions{
				ServiceName:    core.StringPtr("iam-access-management"),
				ServiceGroupID: core.StringPtr("IAM"),
				TransactionID:  getTransactionID(),
			}

			operationsList, response, err := contextBasedRestrictionsService.ListAvailableServiceOperations(listAvailableServiceOperationsOptions)
			Expect(err).ToNot(BeNil())
			Expect(response.StatusCode).To(Equal(400))
			Expect(operationsList).To(BeNil())
		})
	})

	//
	// Cleanup the created zones and rules
	//

	Describe(`DeleteRule - Delete the specified rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteRule(deleteRuleOptions *DeleteRuleOptions)`, func() {
			deleteRuleOptions := &contextbasedrestrictionsv1.DeleteRuleOptions{
				RuleID:        core.StringPtr(ruleID),
				TransactionID: getTransactionID(),
			}

			response, err := contextBasedRestrictionsService.DeleteRule(deleteRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteRule - Delete rule with 'missing required RuleID parameter' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteRule(deleteRuleOptions *DeleteRuleOptions) with 'missing required RuleID parameter' error`, func() {
			deleteRuleOptions := &contextbasedrestrictionsv1.DeleteRuleOptions{
				TransactionID: getTransactionID(),
			}

			response, err := contextBasedRestrictionsService.DeleteRule(deleteRuleOptions)

			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(ContainSubstring("Field validation for 'RuleID' failed"))
			Expect(response).To(BeNil())
		})
	})

	Describe(`DeleteRule - Delete rule with 'rule not found' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteRule(deleteRuleOptions *DeleteRuleOptions) with 'rule not found' error (404)`, func() {
			deleteRuleOptions := &contextbasedrestrictionsv1.DeleteRuleOptions{
				RuleID:        core.StringPtr(NonExistentID),
				TransactionID: getTransactionID(),
			}

			response, err := contextBasedRestrictionsService.DeleteRule(deleteRuleOptions)

			Expect(err).To(Not(BeNil()))
			Expect(response.StatusCode).To(Equal(404))
		})
	})

	Describe(`DeleteZone - Delete the specified zone`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteZone(deleteZoneOptions *DeleteZoneOptions)`, func() {
			deleteZoneOptions := &contextbasedrestrictionsv1.DeleteZoneOptions{
				ZoneID: core.StringPtr(zoneID),
				// Using the standard X-Correlation-Id header in this case
				XCorrelationID: getTransactionID(),
			}

			response, err := contextBasedRestrictionsService.DeleteZone(deleteZoneOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})

	Describe(`DeleteZone - Delete zone with 'missing required ZoneID parameter' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteZone(deleteZoneOptions *DeleteZoneOptions) with 'missing required ZoneID parameter' error`, func() {
			deleteZoneOptions := &contextbasedrestrictionsv1.DeleteZoneOptions{
				TransactionID: getTransactionID(),
			}

			response, err := contextBasedRestrictionsService.DeleteZone(deleteZoneOptions)

			Expect(err).To(Not(BeNil()))
			Expect(err.Error()).To(ContainSubstring("Field validation for 'ZoneID' failed"))
			Expect(response).To(BeNil())
		})
	})

	Describe(`DeleteZone - Delete zone with 'zone not found' error`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteZone(deleteZoneOptions *DeleteZoneOptions) with 'zone not found' error (404)`, func() {
			deleteZoneOptions := &contextbasedrestrictionsv1.DeleteZoneOptions{
				ZoneID:        core.StringPtr(NonExistentID),
				TransactionID: getTransactionID(),
			}

			response, err := contextBasedRestrictionsService.DeleteZone(deleteZoneOptions)

			Expect(err).To(Not(BeNil()))
			Expect(response.StatusCode).To(Equal(404))
		})
	})
})

//
// Utility functions are declared in the unit test file
//

func getTransactionID() *string {
	return core.StringPtr("sdk-test-" + uuid.New().String())
}
