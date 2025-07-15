//go:build examples

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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/contextbasedrestrictionsv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// This file provides an example of how to use the Context Based Restrictions service.
//
// The following configuration properties are assumed to be defined:
// CONTEXT_BASED_RESTRICTIONS_URL=<service base url>
// CONTEXT_BASED_RESTRICTIONS_AUTH_TYPE=iam
// CONTEXT_BASED_RESTRICTIONS_APIKEY=<IAM apikey>
// CONTEXT_BASED_RESTRICTIONS_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
// CONTEXT_BASED_RESTRICTIONS_TEST_ACCOUNT_ID=<the id of the account under which test CBR zones and rules are created>
// CONTEXT_BASED_RESTRICTIONS_TEST_SERVICE_NAME=<the name of the service to be associated with the test CBR rules>
// CONTEXT_BASED_RESTRICTIONS_TEST_VPC_CRN=<the CRN of the vpc instance to be associated with the test CBR rules>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
var _ = Describe(`ContextBasedRestrictionsV1 Examples Tests`, func() {

	const externalConfigFile = "../context_based_restrictions_v1.env"

	var (
		contextBasedRestrictionsService *contextbasedrestrictionsv1.ContextBasedRestrictionsV1
		config                          map[string]string
		configLoaded                    bool = false
		accountID                       string
		serviceName                     string
		vpcCRN                          string
		zoneID                          string
		zoneRev                         string
		ruleID                          string
		ruleRev                         string
	)

	var shouldSkipTest = func() {
		if !configLoaded {
			Skip("External configuration is not available, skipping examples...")
		}

	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping examples: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(contextbasedrestrictionsv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			}

			accountID = config["TEST_ACCOUNT_ID"]
			if accountID == "" {
				Skip("Unable to load TEST_ACCOUNT_ID configuration property, skipping tests")
			}

			serviceName = config["TEST_SERVICE_NAME"]
			if serviceName == "" {
				Skip("Unable to load TEST_SERVICE_NAME configuration property, skipping tests")
			}

			vpcCRN = config["TEST_VPC_CRN"]
			if vpcCRN == "" {
				Skip("Unable to load TEST_VPC_CRN configuration property, skipping tests")
			}

			if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
			}

			configLoaded = len(config) > 0
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			contextBasedRestrictionsServiceOptions := &contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{}

			contextBasedRestrictionsService, err = contextbasedrestrictionsv1.NewContextBasedRestrictionsV1UsingExternalConfig(contextBasedRestrictionsServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(contextBasedRestrictionsService).ToNot(BeNil())
		})
	})

	Describe(`ContextBasedRestrictionsV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateZone request example`, func() {
			fmt.Println("\nCreateZone() result:")
			// begin-create_zone

			ipAddressModel := &contextbasedrestrictionsv1.AddressIPAddress{
				Type:  core.StringPtr("ipAddress"),
				Value: core.StringPtr("169.23.56.234"),
			}
			ipAddressV6Model := &contextbasedrestrictionsv1.AddressIPAddress{
				Type:  core.StringPtr("ipAddress"),
				Value: core.StringPtr("3ffe:1900:fe21:4545::"),
			}
			ipRangeAddressModel := &contextbasedrestrictionsv1.AddressIPAddressRange{
				Type:  core.StringPtr("ipRange"),
				Value: core.StringPtr("169.23.22.0-169.23.22.255"),
			}
			ipRangeAddressV6Model := &contextbasedrestrictionsv1.AddressIPAddressRange{
				Type:  core.StringPtr("ipRange"),
				Value: core.StringPtr("3ffe:1900:fe21:4545::-3ffe:1900:fe21:6767::"),
			}
			subnetAddressModel := &contextbasedrestrictionsv1.AddressSubnet{
				Type:  core.StringPtr("subnet"),
				Value: core.StringPtr("192.0.2.0/24"),
			}
			vpcAddressModel := &contextbasedrestrictionsv1.AddressVPC{
				Type:  core.StringPtr("vpc"),
				Value: core.StringPtr(vpcCRN),
			}
			serviceRefAddressModel := &contextbasedrestrictionsv1.AddressServiceRef{
				Type: core.StringPtr("serviceRef"),
				Ref: &contextbasedrestrictionsv1.ServiceRefValue{
					AccountID:   core.StringPtr(accountID),
					ServiceName: core.StringPtr("cloud-object-storage"),
				},
			}
			excludedIPAddressModel := &contextbasedrestrictionsv1.AddressIPAddress{
				Type:  core.StringPtr("ipAddress"),
				Value: core.StringPtr("169.23.22.127"),
			}

			createZoneOptions := contextBasedRestrictionsService.NewCreateZoneOptions()
			createZoneOptions.SetName("an example of zone")
			createZoneOptions.SetAccountID(accountID)
			createZoneOptions.SetDescription("this is an example of zone")
			createZoneOptions.SetAddresses([]contextbasedrestrictionsv1.AddressIntf{ipAddressModel, ipAddressV6Model, ipRangeAddressModel, ipRangeAddressV6Model, subnetAddressModel, vpcAddressModel, serviceRefAddressModel})
			createZoneOptions.SetExcluded([]contextbasedrestrictionsv1.AddressIntf{excludedIPAddressModel})

			zone, response, err := contextBasedRestrictionsService.CreateZone(createZoneOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(zone, "", "  ")
			fmt.Println(string(b))

			// end-create_zone

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(zone).ToNot(BeNil())
			zoneID = *zone.ID

		})
		It(`ListZones request example`, func() {
			fmt.Println("\nListZones() result:")
			// begin-list_zones

			listZonesOptions := contextBasedRestrictionsService.NewListZonesOptions(
				accountID,
			)

			zoneList, response, err := contextBasedRestrictionsService.ListZones(listZonesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(zoneList, "", "  ")
			fmt.Println(string(b))

			// end-list_zones

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(zoneList).ToNot(BeNil())

		})
		It(`GetZone request example`, func() {
			fmt.Println("\nGetZone() result:")
			// begin-get_zone

			getZoneOptions := contextBasedRestrictionsService.NewGetZoneOptions(
				zoneID,
			)

			zone, response, err := contextBasedRestrictionsService.GetZone(getZoneOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(zone, "", "  ")
			fmt.Println(string(b))

			// end-get_zone

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(zone).ToNot(BeNil())
			zoneRev = response.Headers.Get("Etag")

		})
		It(`ReplaceZone request example`, func() {
			fmt.Println("\nReplaceZone() result:")
			// begin-replace_zone

			addressModel := &contextbasedrestrictionsv1.AddressIPAddress{
				Type:  core.StringPtr("ipAddress"),
				Value: core.StringPtr("169.23.56.234"),
			}

			replaceZoneOptions := contextBasedRestrictionsService.NewReplaceZoneOptions(
				zoneID,
				zoneRev,
			)
			replaceZoneOptions.SetName("an example of updated zone")
			replaceZoneOptions.SetAccountID(accountID)
			replaceZoneOptions.SetDescription("this is an example of updated zone")
			replaceZoneOptions.SetAddresses([]contextbasedrestrictionsv1.AddressIntf{addressModel})

			zone, response, err := contextBasedRestrictionsService.ReplaceZone(replaceZoneOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(zone, "", "  ")
			fmt.Println(string(b))

			// end-replace_zone

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(zone).ToNot(BeNil())

		})
		It(`ListAvailableServiceRefTargets request example`, func() {
			fmt.Println("\nListAvailableServiceRefTargets() result:")
			// begin-list_available_serviceref_targets

			listAvailableServiceRefTargetsOptions := contextBasedRestrictionsService.NewListAvailableServicerefTargetsOptions()

			serviceRefTargetList, response, err := contextBasedRestrictionsService.ListAvailableServicerefTargets(listAvailableServiceRefTargetsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(serviceRefTargetList, "", "  ")
			fmt.Println(string(b))

			// end-list_available_serviceref_targets

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceRefTargetList).ToNot(BeNil())

		})
		It(`GetServicerefTarget request example`, func() {
			fmt.Println("\nGetServicerefTarget() result:")
			tempServiceName := serviceName
			serviceName = "containers-kubernetes"
			// begin-get_serviceref_target

			getServicerefTargetOptions := contextBasedRestrictionsService.NewGetServicerefTargetOptions(
				serviceName,
			)

			serviceRefTarget, response, err := contextBasedRestrictionsService.GetServicerefTarget(getServicerefTargetOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(serviceRefTarget, "", "  ")
			fmt.Println(string(b))

			// end-get_serviceref_target
			serviceName = tempServiceName

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceRefTarget).ToNot(BeNil())
		})
		It(`CreateRule request example`, func() {
			fmt.Println("\nCreateRule() result:")
			// begin-create_rule

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
						Value: core.StringPtr(accountID),
					},
					{
						Name:  core.StringPtr("serviceName"),
						Value: core.StringPtr(serviceName),
					},
				},
				Tags: []contextbasedrestrictionsv1.ResourceTagAttribute{
					{
						Name:  core.StringPtr("tagName"),
						Value: core.StringPtr("tagValue"),
					},
				},
			}

			createRuleOptions := contextBasedRestrictionsService.NewCreateRuleOptions()
			createRuleOptions.SetDescription("this is an example of rule")
			createRuleOptions.SetContexts([]contextbasedrestrictionsv1.RuleContext{*ruleContextModel})
			createRuleOptions.SetResources([]contextbasedrestrictionsv1.Resource{*resourceModel})
			createRuleOptions.SetEnforcementMode(contextbasedrestrictionsv1.CreateRuleOptionsEnforcementModeEnabledConst)
			rule, response, err := contextBasedRestrictionsService.CreateRule(createRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(rule, "", "  ")
			fmt.Println(string(b))

			// end-create_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(rule).ToNot(BeNil())
			ruleID = *rule.ID

		})
		It(`ListRules request example`, func() {
			fmt.Println("\nListRules() result:")
			// begin-list_rules

			listRulesOptions := contextBasedRestrictionsService.NewListRulesOptions(
				accountID,
			)

			ruleList, response, err := contextBasedRestrictionsService.ListRules(listRulesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(ruleList, "", "  ")
			fmt.Println(string(b))

			// end-list_rules

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ruleList).ToNot(BeNil())
		})
		It(`GetRule request example`, func() {
			fmt.Println("\nGetRule() result:")
			// begin-get_rule

			getRuleOptions := contextBasedRestrictionsService.NewGetRuleOptions(
				ruleID,
			)

			rule, response, err := contextBasedRestrictionsService.GetRule(getRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(rule, "", "  ")
			fmt.Println(string(b))

			// end-get_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())
			ruleRev = response.Headers.Get("Etag")
		})
		It(`ReplaceRule request example`, func() {
			fmt.Println("\nReplaceRule() result:")
			// begin-replace_rule

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
						Value: core.StringPtr(accountID),
					},
					{
						Name:  core.StringPtr("serviceName"),
						Value: core.StringPtr(serviceName),
					},
				},
				Tags: []contextbasedrestrictionsv1.ResourceTagAttribute{
					{
						Name:  core.StringPtr("tagName"),
						Value: core.StringPtr("updatedTagValue"),
					},
				},
			}

			replaceRuleOptions := contextBasedRestrictionsService.NewReplaceRuleOptions(
				ruleID,
				ruleRev,
			)
			replaceRuleOptions.SetDescription("this is an example of rule")
			replaceRuleOptions.SetContexts([]contextbasedrestrictionsv1.RuleContext{*ruleContextModel})
			replaceRuleOptions.SetResources([]contextbasedrestrictionsv1.Resource{*resourceModel})
			replaceRuleOptions.SetEnforcementMode(contextbasedrestrictionsv1.ReplaceRuleOptionsEnforcementModeDisabledConst)

			rule, response, err := contextBasedRestrictionsService.ReplaceRule(replaceRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(rule, "", "  ")
			fmt.Println(string(b))

			// end-replace_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())

		})
		It(`GetAccountSettings request example`, func() {
			fmt.Println("\nGetAccountSettings() result:")
			// begin-get_account_settings

			getAccountSettingsOptions := contextBasedRestrictionsService.NewGetAccountSettingsOptions(
				accountID,
			)

			accountSettings, response, err := contextBasedRestrictionsService.GetAccountSettings(getAccountSettingsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accountSettings, "", "  ")
			fmt.Println(string(b))

			// end-get_account_settings

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountSettings).ToNot(BeNil())

		})

		It(`ListAvailableServiceOperations request example`, func() {
			fmt.Println("\nListAvailableServiceOperations() result:")
			// begin-list_available_service_operations

			listAvailableServiceOperationsOptions := contextBasedRestrictionsService.NewListAvailableServiceOperationsOptions()
			listAvailableServiceOperationsOptions.SetServiceName("containers-kubernetes")

			operationsList, response, err := contextBasedRestrictionsService.ListAvailableServiceOperations(listAvailableServiceOperationsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(operationsList, "", "  ")
			fmt.Println(string(b))

			// end-list_available_service_operations

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(operationsList).ToNot(BeNil())
		})

		It(`DeleteRule request example`, func() {
			// begin-delete_rule

			deleteRuleOptions := contextBasedRestrictionsService.NewDeleteRuleOptions(
				ruleID,
			)

			response, err := contextBasedRestrictionsService.DeleteRule(deleteRuleOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteRule(): %d\n", response.StatusCode)
			}

			// end-delete_rule
			fmt.Printf("\nDeleteRule() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteZone request example`, func() {
			// begin-delete_zone

			deleteZoneOptions := contextBasedRestrictionsService.NewDeleteZoneOptions(
				zoneID,
			)

			response, err := contextBasedRestrictionsService.DeleteZone(deleteZoneOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteZone(): %d\n", response.StatusCode)
			}

			// end-delete_zone
			fmt.Printf("\nDeleteZone() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})
