// +build integration

/**
 * (C) Copyright IBM Corp. 2021.
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

package vpcv1_test

import (
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the vpcv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`VpcV1 Integration Tests`, func() {

	const externalConfigFile = "../vpc_v1.env"

	var (
		err        error
		vpcService *vpcv1.VpcV1
		serviceURL string
		config     map[string]string
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
			config, err = core.GetServiceProperties(vpcv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Printf("Service URL: %s\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {

			vpcServiceOptions := &vpcv1.VpcV1Options{
				Version: core.StringPtr("testString"),
			}

			vpcService, err = vpcv1.NewVpcV1UsingExternalConfig(vpcServiceOptions)

			Expect(err).To(BeNil())
			Expect(vpcService).ToNot(BeNil())
			Expect(vpcService.Service.Options.URL).To(Equal(serviceURL))
		})
	})

	Describe(`ListVpcs - List all VPCs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListVpcs(listVpcsOptions *ListVpcsOptions)`, func() {

			listVpcsOptions := &vpcv1.ListVpcsOptions{
				Start:           core.StringPtr("testString"),
				Limit:           core.Int64Ptr(int64(1)),
				ResourceGroupID: core.StringPtr("testString"),
				ClassicAccess:   core.BoolPtr(true),
			}

			vpcCollection, response, err := vpcService.ListVpcs(listVpcsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpcCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateVPC - Create a VPC`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateVPC(createVPCOptions *CreateVPCOptions)`, func() {

			resourceGroupIdentityModel := &vpcv1.ResourceGroupIdentityByID{
				ID: core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345"),
			}

			createVPCOptions := &vpcv1.CreateVPCOptions{
				AddressPrefixManagement: core.StringPtr("manual"),
				ClassicAccess:           core.BoolPtr(false),
				Name:                    core.StringPtr("my-vpc"),
				ResourceGroup:           resourceGroupIdentityModel,
			}

			vpc, response, err := vpcService.CreateVPC(createVPCOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(vpc).ToNot(BeNil())

		})
	})

	Describe(`GetVPC - Retrieve a VPC`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetVPC(getVPCOptions *GetVPCOptions)`, func() {

			getVPCOptions := &vpcv1.GetVPCOptions{
				ID: core.StringPtr("testString"),
			}

			vpc, response, err := vpcService.GetVPC(getVPCOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpc).ToNot(BeNil())

		})
	})

	Describe(`UpdateVPC - Update a VPC`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateVPC(updateVPCOptions *UpdateVPCOptions)`, func() {

			vpcPatchModel := &vpcv1.VPCPatch{
				Name: core.StringPtr("my-vpc"),
			}
			vpcPatchModelAsPatch, asPatchErr := vpcPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateVPCOptions := &vpcv1.UpdateVPCOptions{
				ID:       core.StringPtr("testString"),
				VPCPatch: vpcPatchModelAsPatch,
			}

			vpc, response, err := vpcService.UpdateVPC(updateVPCOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpc).ToNot(BeNil())

		})
	})

	Describe(`GetVPCDefaultNetworkACL - Retrieve a VPC's default network ACL`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetVPCDefaultNetworkACL(getVPCDefaultNetworkACLOptions *GetVPCDefaultNetworkACLOptions)`, func() {

			getVPCDefaultNetworkACLOptions := &vpcv1.GetVPCDefaultNetworkACLOptions{
				ID: core.StringPtr("testString"),
			}

			defaultNetworkACL, response, err := vpcService.GetVPCDefaultNetworkACL(getVPCDefaultNetworkACLOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(defaultNetworkACL).ToNot(BeNil())

		})
	})

	Describe(`GetVPCDefaultRoutingTable - Retrieve a VPC's default routing table`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetVPCDefaultRoutingTable(getVPCDefaultRoutingTableOptions *GetVPCDefaultRoutingTableOptions)`, func() {

			getVPCDefaultRoutingTableOptions := &vpcv1.GetVPCDefaultRoutingTableOptions{
				ID: core.StringPtr("testString"),
			}

			defaultRoutingTable, response, err := vpcService.GetVPCDefaultRoutingTable(getVPCDefaultRoutingTableOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(defaultRoutingTable).ToNot(BeNil())

		})
	})

	Describe(`GetVPCDefaultSecurityGroup - Retrieve a VPC's default security group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetVPCDefaultSecurityGroup(getVPCDefaultSecurityGroupOptions *GetVPCDefaultSecurityGroupOptions)`, func() {

			getVPCDefaultSecurityGroupOptions := &vpcv1.GetVPCDefaultSecurityGroupOptions{
				ID: core.StringPtr("testString"),
			}

			defaultSecurityGroup, response, err := vpcService.GetVPCDefaultSecurityGroup(getVPCDefaultSecurityGroupOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(defaultSecurityGroup).ToNot(BeNil())

		})
	})

	Describe(`ListVPCAddressPrefixes - List all address prefixes for a VPC`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListVPCAddressPrefixes(listVPCAddressPrefixesOptions *ListVPCAddressPrefixesOptions)`, func() {

			listVPCAddressPrefixesOptions := &vpcv1.ListVPCAddressPrefixesOptions{
				VPCID: core.StringPtr("testString"),
				Start: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(1)),
			}

			addressPrefixCollection, response, err := vpcService.ListVPCAddressPrefixes(listVPCAddressPrefixesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(addressPrefixCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateVPCAddressPrefix - Create an address prefix for a VPC`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateVPCAddressPrefix(createVPCAddressPrefixOptions *CreateVPCAddressPrefixOptions)`, func() {

			zoneIdentityModel := &vpcv1.ZoneIdentityByName{
				Name: core.StringPtr("us-south-1"),
			}

			createVPCAddressPrefixOptions := &vpcv1.CreateVPCAddressPrefixOptions{
				VPCID:     core.StringPtr("testString"),
				CIDR:      core.StringPtr("10.0.0.0/24"),
				Zone:      zoneIdentityModel,
				IsDefault: core.BoolPtr(true),
				Name:      core.StringPtr("my-address-prefix-2"),
			}

			addressPrefix, response, err := vpcService.CreateVPCAddressPrefix(createVPCAddressPrefixOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(addressPrefix).ToNot(BeNil())

		})
	})

	Describe(`GetVPCAddressPrefix - Retrieve an address prefix`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetVPCAddressPrefix(getVPCAddressPrefixOptions *GetVPCAddressPrefixOptions)`, func() {

			getVPCAddressPrefixOptions := &vpcv1.GetVPCAddressPrefixOptions{
				VPCID: core.StringPtr("testString"),
				ID:    core.StringPtr("testString"),
			}

			addressPrefix, response, err := vpcService.GetVPCAddressPrefix(getVPCAddressPrefixOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(addressPrefix).ToNot(BeNil())

		})
	})

	Describe(`UpdateVPCAddressPrefix - Update an address prefix`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateVPCAddressPrefix(updateVPCAddressPrefixOptions *UpdateVPCAddressPrefixOptions)`, func() {

			addressPrefixPatchModel := &vpcv1.AddressPrefixPatch{
				IsDefault: core.BoolPtr(false),
				Name:      core.StringPtr("my-address-prefix-2"),
			}
			addressPrefixPatchModelAsPatch, asPatchErr := addressPrefixPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateVPCAddressPrefixOptions := &vpcv1.UpdateVPCAddressPrefixOptions{
				VPCID:              core.StringPtr("testString"),
				ID:                 core.StringPtr("testString"),
				AddressPrefixPatch: addressPrefixPatchModelAsPatch,
			}

			addressPrefix, response, err := vpcService.UpdateVPCAddressPrefix(updateVPCAddressPrefixOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(addressPrefix).ToNot(BeNil())

		})
	})

	Describe(`ListVPCRoutes - List all routes in a VPC's default routing table`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListVPCRoutes(listVPCRoutesOptions *ListVPCRoutesOptions)`, func() {

			listVPCRoutesOptions := &vpcv1.ListVPCRoutesOptions{
				VPCID:    core.StringPtr("testString"),
				ZoneName: core.StringPtr("testString"),
				Start:    core.StringPtr("testString"),
				Limit:    core.Int64Ptr(int64(1)),
			}

			routeCollection, response, err := vpcService.ListVPCRoutes(listVPCRoutesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(routeCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateVPCRoute - Create a route in a VPC's default routing table`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateVPCRoute(createVPCRouteOptions *CreateVPCRouteOptions)`, func() {

			zoneIdentityModel := &vpcv1.ZoneIdentityByName{
				Name: core.StringPtr("us-south-1"),
			}

			routeNextHopPrototypeModel := &vpcv1.RouteNextHopPrototypeRouteNextHopIP{
				Address: core.StringPtr("192.168.3.4"),
			}

			createVPCRouteOptions := &vpcv1.CreateVPCRouteOptions{
				VPCID:       core.StringPtr("testString"),
				Destination: core.StringPtr("192.168.3.0/24"),
				Zone:        zoneIdentityModel,
				Action:      core.StringPtr("delegate"),
				Name:        core.StringPtr("my-route-2"),
				NextHop:     routeNextHopPrototypeModel,
			}

			route, response, err := vpcService.CreateVPCRoute(createVPCRouteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(route).ToNot(BeNil())

		})
	})

	Describe(`GetVPCRoute - Retrieve a VPC route`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetVPCRoute(getVPCRouteOptions *GetVPCRouteOptions)`, func() {

			getVPCRouteOptions := &vpcv1.GetVPCRouteOptions{
				VPCID: core.StringPtr("testString"),
				ID:    core.StringPtr("testString"),
			}

			route, response, err := vpcService.GetVPCRoute(getVPCRouteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(route).ToNot(BeNil())

		})
	})

	Describe(`UpdateVPCRoute - Update a VPC route`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateVPCRoute(updateVPCRouteOptions *UpdateVPCRouteOptions)`, func() {

			routePatchModel := &vpcv1.RoutePatch{
				Name: core.StringPtr("my-route-2"),
			}
			routePatchModelAsPatch, asPatchErr := routePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateVPCRouteOptions := &vpcv1.UpdateVPCRouteOptions{
				VPCID:      core.StringPtr("testString"),
				ID:         core.StringPtr("testString"),
				RoutePatch: routePatchModelAsPatch,
			}

			route, response, err := vpcService.UpdateVPCRoute(updateVPCRouteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(route).ToNot(BeNil())

		})
	})

	Describe(`ListVPCRoutingTables - List all routing tables for a VPC`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListVPCRoutingTables(listVPCRoutingTablesOptions *ListVPCRoutingTablesOptions)`, func() {

			listVPCRoutingTablesOptions := &vpcv1.ListVPCRoutingTablesOptions{
				VPCID:     core.StringPtr("testString"),
				Start:     core.StringPtr("testString"),
				Limit:     core.Int64Ptr(int64(1)),
				IsDefault: core.BoolPtr(true),
			}

			routingTableCollection, response, err := vpcService.ListVPCRoutingTables(listVPCRoutingTablesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(routingTableCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateVPCRoutingTable - Create a routing table for a VPC`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateVPCRoutingTable(createVPCRoutingTableOptions *CreateVPCRoutingTableOptions)`, func() {

			resourceFilterModel := &vpcv1.ResourceFilter{
				ResourceType: core.StringPtr("vpn_gateway"),
			}

			routeNextHopPrototypeModel := &vpcv1.RouteNextHopPrototypeRouteNextHopIP{
				Address: core.StringPtr("192.168.3.4"),
			}

			zoneIdentityModel := &vpcv1.ZoneIdentityByName{
				Name: core.StringPtr("us-south-1"),
			}

			routePrototypeModel := &vpcv1.RoutePrototype{
				Action:      core.StringPtr("delegate"),
				Destination: core.StringPtr("192.168.3.0/24"),
				Name:        core.StringPtr("my-route-2"),
				NextHop:     routeNextHopPrototypeModel,
				Zone:        zoneIdentityModel,
			}

			createVPCRoutingTableOptions := &vpcv1.CreateVPCRoutingTableOptions{
				VPCID:                      core.StringPtr("testString"),
				AcceptRoutesFrom:           []vpcv1.ResourceFilter{*resourceFilterModel},
				Name:                       core.StringPtr("my-routing-table-2"),
				RouteDirectLinkIngress:     core.BoolPtr(true),
				RouteTransitGatewayIngress: core.BoolPtr(true),
				RouteVPCZoneIngress:        core.BoolPtr(true),
				Routes:                     []vpcv1.RoutePrototype{*routePrototypeModel},
			}

			routingTable, response, err := vpcService.CreateVPCRoutingTable(createVPCRoutingTableOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(routingTable).ToNot(BeNil())

		})
	})

	Describe(`GetVPCRoutingTable - Retrieve a VPC routing table`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetVPCRoutingTable(getVPCRoutingTableOptions *GetVPCRoutingTableOptions)`, func() {

			getVPCRoutingTableOptions := &vpcv1.GetVPCRoutingTableOptions{
				VPCID: core.StringPtr("testString"),
				ID:    core.StringPtr("testString"),
			}

			routingTable, response, err := vpcService.GetVPCRoutingTable(getVPCRoutingTableOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(routingTable).ToNot(BeNil())

		})
	})

	Describe(`UpdateVPCRoutingTable - Update a VPC routing table`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateVPCRoutingTable(updateVPCRoutingTableOptions *UpdateVPCRoutingTableOptions)`, func() {

			resourceFilterModel := &vpcv1.ResourceFilter{
				ResourceType: core.StringPtr("vpn_gateway"),
			}

			routingTablePatchModel := &vpcv1.RoutingTablePatch{
				AcceptRoutesFrom:           []vpcv1.ResourceFilter{*resourceFilterModel},
				Name:                       core.StringPtr("my-routing-table-2"),
				RouteDirectLinkIngress:     core.BoolPtr(true),
				RouteTransitGatewayIngress: core.BoolPtr(true),
				RouteVPCZoneIngress:        core.BoolPtr(true),
			}
			routingTablePatchModelAsPatch, asPatchErr := routingTablePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateVPCRoutingTableOptions := &vpcv1.UpdateVPCRoutingTableOptions{
				VPCID:             core.StringPtr("testString"),
				ID:                core.StringPtr("testString"),
				RoutingTablePatch: routingTablePatchModelAsPatch,
				IfMatch:           core.StringPtr("96d225c4-56bd-43d9-98fc-d7148e5c5028"),
			}

			routingTable, response, err := vpcService.UpdateVPCRoutingTable(updateVPCRoutingTableOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(routingTable).ToNot(BeNil())

		})
	})

	Describe(`ListVPCRoutingTableRoutes - List all routes in a VPC routing table`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListVPCRoutingTableRoutes(listVPCRoutingTableRoutesOptions *ListVPCRoutingTableRoutesOptions)`, func() {

			listVPCRoutingTableRoutesOptions := &vpcv1.ListVPCRoutingTableRoutesOptions{
				VPCID:          core.StringPtr("testString"),
				RoutingTableID: core.StringPtr("testString"),
				Start:          core.StringPtr("testString"),
				Limit:          core.Int64Ptr(int64(1)),
			}

			routeCollection, response, err := vpcService.ListVPCRoutingTableRoutes(listVPCRoutingTableRoutesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(routeCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateVPCRoutingTableRoute - Create a route in a VPC routing table`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateVPCRoutingTableRoute(createVPCRoutingTableRouteOptions *CreateVPCRoutingTableRouteOptions)`, func() {

			zoneIdentityModel := &vpcv1.ZoneIdentityByName{
				Name: core.StringPtr("us-south-1"),
			}

			routeNextHopPrototypeModel := &vpcv1.RouteNextHopPrototypeRouteNextHopIP{
				Address: core.StringPtr("192.168.3.4"),
			}

			createVPCRoutingTableRouteOptions := &vpcv1.CreateVPCRoutingTableRouteOptions{
				VPCID:          core.StringPtr("testString"),
				RoutingTableID: core.StringPtr("testString"),
				Destination:    core.StringPtr("192.168.3.0/24"),
				Zone:           zoneIdentityModel,
				Action:         core.StringPtr("delegate"),
				Name:           core.StringPtr("my-route-2"),
				NextHop:        routeNextHopPrototypeModel,
			}

			route, response, err := vpcService.CreateVPCRoutingTableRoute(createVPCRoutingTableRouteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(route).ToNot(BeNil())

		})
	})

	Describe(`GetVPCRoutingTableRoute - Retrieve a VPC routing table route`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetVPCRoutingTableRoute(getVPCRoutingTableRouteOptions *GetVPCRoutingTableRouteOptions)`, func() {

			getVPCRoutingTableRouteOptions := &vpcv1.GetVPCRoutingTableRouteOptions{
				VPCID:          core.StringPtr("testString"),
				RoutingTableID: core.StringPtr("testString"),
				ID:             core.StringPtr("testString"),
			}

			route, response, err := vpcService.GetVPCRoutingTableRoute(getVPCRoutingTableRouteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(route).ToNot(BeNil())

		})
	})

	Describe(`UpdateVPCRoutingTableRoute - Update a VPC routing table route`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateVPCRoutingTableRoute(updateVPCRoutingTableRouteOptions *UpdateVPCRoutingTableRouteOptions)`, func() {

			routePatchModel := &vpcv1.RoutePatch{
				Name: core.StringPtr("my-route-2"),
			}
			routePatchModelAsPatch, asPatchErr := routePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateVPCRoutingTableRouteOptions := &vpcv1.UpdateVPCRoutingTableRouteOptions{
				VPCID:          core.StringPtr("testString"),
				RoutingTableID: core.StringPtr("testString"),
				ID:             core.StringPtr("testString"),
				RoutePatch:     routePatchModelAsPatch,
			}

			route, response, err := vpcService.UpdateVPCRoutingTableRoute(updateVPCRoutingTableRouteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(route).ToNot(BeNil())

		})
	})

	Describe(`ListSubnets - List all subnets`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSubnets(listSubnetsOptions *ListSubnetsOptions)`, func() {

			listSubnetsOptions := &vpcv1.ListSubnetsOptions{
				Start:            core.StringPtr("testString"),
				Limit:            core.Int64Ptr(int64(1)),
				ResourceGroupID:  core.StringPtr("testString"),
				RoutingTableID:   core.StringPtr("testString"),
				RoutingTableName: core.StringPtr("testString"),
			}

			subnetCollection, response, err := vpcService.ListSubnets(listSubnetsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subnetCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateSubnet - Create a subnet`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSubnet(createSubnetOptions *CreateSubnetOptions)`, func() {

			networkACLIdentityModel := &vpcv1.NetworkACLIdentityByID{
				ID: core.StringPtr("a4e28308-8ee7-46ab-8108-9f881f22bdbf"),
			}

			publicGatewayIdentityModel := &vpcv1.PublicGatewayIdentityByID{
				ID: core.StringPtr("dc5431ef-1fc6-4861-adc9-a59d077d1241"),
			}

			resourceGroupIdentityModel := &vpcv1.ResourceGroupIdentityByID{
				ID: core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345"),
			}

			routingTableIdentityModel := &vpcv1.RoutingTableIdentityByID{
				ID: core.StringPtr("6885e83f-03b2-4603-8a86-db2a0f55c840"),
			}

			vpcIdentityModel := &vpcv1.VPCIdentityByID{
				ID: core.StringPtr("4727d842-f94f-4a2d-824a-9bc9b02c523b"),
			}

			zoneIdentityModel := &vpcv1.ZoneIdentityByName{
				Name: core.StringPtr("us-south-1"),
			}

			subnetPrototypeModel := &vpcv1.SubnetPrototypeSubnetByTotalCount{
				IPVersion:             core.StringPtr("ipv4"),
				Name:                  core.StringPtr("my-subnet"),
				NetworkACL:            networkACLIdentityModel,
				PublicGateway:         publicGatewayIdentityModel,
				ResourceGroup:         resourceGroupIdentityModel,
				RoutingTable:          routingTableIdentityModel,
				VPC:                   vpcIdentityModel,
				TotalIpv4AddressCount: core.Int64Ptr(int64(256)),
				Zone:                  zoneIdentityModel,
			}

			createSubnetOptions := &vpcv1.CreateSubnetOptions{
				SubnetPrototype: subnetPrototypeModel,
			}

			subnet, response, err := vpcService.CreateSubnet(createSubnetOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subnet).ToNot(BeNil())

		})
	})

	Describe(`GetSubnet - Retrieve a subnet`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSubnet(getSubnetOptions *GetSubnetOptions)`, func() {

			getSubnetOptions := &vpcv1.GetSubnetOptions{
				ID: core.StringPtr("testString"),
			}

			subnet, response, err := vpcService.GetSubnet(getSubnetOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subnet).ToNot(BeNil())

		})
	})

	Describe(`UpdateSubnet - Update a subnet`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateSubnet(updateSubnetOptions *UpdateSubnetOptions)`, func() {

			networkACLIdentityModel := &vpcv1.NetworkACLIdentityByID{
				ID: core.StringPtr("a4e28308-8ee7-46ab-8108-9f881f22bdbf"),
			}

			publicGatewayIdentityModel := &vpcv1.PublicGatewayIdentityByID{
				ID: core.StringPtr("dc5431ef-1fc6-4861-adc9-a59d077d1241"),
			}

			routingTableIdentityModel := &vpcv1.RoutingTableIdentityByID{
				ID: core.StringPtr("6885e83f-03b2-4603-8a86-db2a0f55c840"),
			}

			subnetPatchModel := &vpcv1.SubnetPatch{
				Name:          core.StringPtr("my-subnet"),
				NetworkACL:    networkACLIdentityModel,
				PublicGateway: publicGatewayIdentityModel,
				RoutingTable:  routingTableIdentityModel,
			}
			subnetPatchModelAsPatch, asPatchErr := subnetPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateSubnetOptions := &vpcv1.UpdateSubnetOptions{
				ID:          core.StringPtr("testString"),
				SubnetPatch: subnetPatchModelAsPatch,
			}

			subnet, response, err := vpcService.UpdateSubnet(updateSubnetOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subnet).ToNot(BeNil())

		})
	})

	Describe(`GetSubnetNetworkACL - Retrieve a subnet's attached network ACL`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSubnetNetworkACL(getSubnetNetworkACLOptions *GetSubnetNetworkACLOptions)`, func() {

			getSubnetNetworkACLOptions := &vpcv1.GetSubnetNetworkACLOptions{
				ID: core.StringPtr("testString"),
			}

			networkACL, response, err := vpcService.GetSubnetNetworkACL(getSubnetNetworkACLOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkACL).ToNot(BeNil())

		})
	})

	Describe(`ReplaceSubnetNetworkACL - Attach a network ACL to a subnet`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceSubnetNetworkACL(replaceSubnetNetworkACLOptions *ReplaceSubnetNetworkACLOptions)`, func() {

			networkACLIdentityModel := &vpcv1.NetworkACLIdentityByID{
				ID: core.StringPtr("a4e28308-8ee7-46ab-8108-9f881f22bdbf"),
			}

			replaceSubnetNetworkACLOptions := &vpcv1.ReplaceSubnetNetworkACLOptions{
				ID:                 core.StringPtr("testString"),
				NetworkACLIdentity: networkACLIdentityModel,
			}

			networkACL, response, err := vpcService.ReplaceSubnetNetworkACL(replaceSubnetNetworkACLOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(networkACL).ToNot(BeNil())

		})
	})

	Describe(`GetSubnetPublicGateway - Retrieve a subnet's attached public gateway`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSubnetPublicGateway(getSubnetPublicGatewayOptions *GetSubnetPublicGatewayOptions)`, func() {

			getSubnetPublicGatewayOptions := &vpcv1.GetSubnetPublicGatewayOptions{
				ID: core.StringPtr("testString"),
			}

			publicGateway, response, err := vpcService.GetSubnetPublicGateway(getSubnetPublicGatewayOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(publicGateway).ToNot(BeNil())

		})
	})

	Describe(`SetSubnetPublicGateway - Attach a public gateway to a subnet`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`SetSubnetPublicGateway(setSubnetPublicGatewayOptions *SetSubnetPublicGatewayOptions)`, func() {

			publicGatewayIdentityModel := &vpcv1.PublicGatewayIdentityByID{
				ID: core.StringPtr("dc5431ef-1fc6-4861-adc9-a59d077d1241"),
			}

			setSubnetPublicGatewayOptions := &vpcv1.SetSubnetPublicGatewayOptions{
				ID:                    core.StringPtr("testString"),
				PublicGatewayIdentity: publicGatewayIdentityModel,
			}

			publicGateway, response, err := vpcService.SetSubnetPublicGateway(setSubnetPublicGatewayOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(publicGateway).ToNot(BeNil())

		})
	})

	Describe(`GetSubnetRoutingTable - Retrieve a subnet's attached routing table`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSubnetRoutingTable(getSubnetRoutingTableOptions *GetSubnetRoutingTableOptions)`, func() {

			getSubnetRoutingTableOptions := &vpcv1.GetSubnetRoutingTableOptions{
				ID: core.StringPtr("testString"),
			}

			routingTable, response, err := vpcService.GetSubnetRoutingTable(getSubnetRoutingTableOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(routingTable).ToNot(BeNil())

		})
	})

	Describe(`ReplaceSubnetRoutingTable - Attach a routing table to a subnet`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceSubnetRoutingTable(replaceSubnetRoutingTableOptions *ReplaceSubnetRoutingTableOptions)`, func() {

			routingTableIdentityModel := &vpcv1.RoutingTableIdentityByID{
				ID: core.StringPtr("1a15dca5-7e33-45e1-b7c5-bc690e569531"),
			}

			replaceSubnetRoutingTableOptions := &vpcv1.ReplaceSubnetRoutingTableOptions{
				ID:                   core.StringPtr("testString"),
				RoutingTableIdentity: routingTableIdentityModel,
			}

			routingTable, response, err := vpcService.ReplaceSubnetRoutingTable(replaceSubnetRoutingTableOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(routingTable).ToNot(BeNil())

		})
	})

	Describe(`ListSubnetReservedIps - List all reserved IPs in a subnet`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSubnetReservedIps(listSubnetReservedIpsOptions *ListSubnetReservedIpsOptions)`, func() {

			listSubnetReservedIpsOptions := &vpcv1.ListSubnetReservedIpsOptions{
				SubnetID: core.StringPtr("testString"),
				Start:    core.StringPtr("testString"),
				Limit:    core.Int64Ptr(int64(1)),
				Sort:     core.StringPtr("name"),
			}

			reservedIPCollection, response, err := vpcService.ListSubnetReservedIps(listSubnetReservedIpsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reservedIPCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateSubnetReservedIP - Reserve an IP in a subnet`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSubnetReservedIP(createSubnetReservedIPOptions *CreateSubnetReservedIPOptions)`, func() {

			reservedIPTargetPrototypeModel := &vpcv1.ReservedIPTargetPrototypeEndpointGatewayIdentityEndpointGatewayIdentityByID{
				ID: core.StringPtr("d7cc5196-9864-48c4-82d8-3f30da41fcc5"),
			}

			createSubnetReservedIPOptions := &vpcv1.CreateSubnetReservedIPOptions{
				SubnetID:   core.StringPtr("testString"),
				Address:    core.StringPtr("192.168.3.4"),
				AutoDelete: core.BoolPtr(false),
				Name:       core.StringPtr("my-reserved-ip"),
				Target:     reservedIPTargetPrototypeModel,
			}

			reservedIP, response, err := vpcService.CreateSubnetReservedIP(createSubnetReservedIPOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(reservedIP).ToNot(BeNil())

		})
	})

	Describe(`GetSubnetReservedIP - Retrieve a reserved IP`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSubnetReservedIP(getSubnetReservedIPOptions *GetSubnetReservedIPOptions)`, func() {

			getSubnetReservedIPOptions := &vpcv1.GetSubnetReservedIPOptions{
				SubnetID: core.StringPtr("testString"),
				ID:       core.StringPtr("testString"),
			}

			reservedIP, response, err := vpcService.GetSubnetReservedIP(getSubnetReservedIPOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reservedIP).ToNot(BeNil())

		})
	})

	Describe(`UpdateSubnetReservedIP - Update a reserved IP`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateSubnetReservedIP(updateSubnetReservedIPOptions *UpdateSubnetReservedIPOptions)`, func() {

			reservedIPPatchModel := &vpcv1.ReservedIPPatch{
				AutoDelete: core.BoolPtr(false),
				Name:       core.StringPtr("my-reserved-ip"),
			}
			reservedIPPatchModelAsPatch, asPatchErr := reservedIPPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateSubnetReservedIPOptions := &vpcv1.UpdateSubnetReservedIPOptions{
				SubnetID:        core.StringPtr("testString"),
				ID:              core.StringPtr("testString"),
				ReservedIPPatch: reservedIPPatchModelAsPatch,
			}

			reservedIP, response, err := vpcService.UpdateSubnetReservedIP(updateSubnetReservedIPOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reservedIP).ToNot(BeNil())

		})
	})

	Describe(`ListImages - List all images`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListImages(listImagesOptions *ListImagesOptions)`, func() {

			listImagesOptions := &vpcv1.ListImagesOptions{
				Start:           core.StringPtr("testString"),
				Limit:           core.Int64Ptr(int64(1)),
				ResourceGroupID: core.StringPtr("testString"),
				Name:            core.StringPtr("testString"),
				Provisionable:   []bool{true},
				Visibility:      core.StringPtr("private"),
			}

			imageCollection, response, err := vpcService.ListImages(listImagesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(imageCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateImage - Create an image`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateImage(createImageOptions *CreateImageOptions)`, func() {

			imageRequiredImageFlagsModel := &vpcv1.ImageRequiredImageFlags{}

			resourceGroupIdentityModel := &vpcv1.ResourceGroupIdentityByID{
				ID: core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345"),
			}

			encryptionKeyIdentityModel := &vpcv1.EncryptionKeyIdentityByCRN{
				CRN: core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/dffc98a0f1f0f95f6613b3b752286b87:e4a29d1a-2ef0-42a6-8fd2-350deb1c647e:key:5437653b-c4b1-447f-9646-b2a2a4cd6179"),
			}

			imageFilePrototypeModel := &vpcv1.ImageFilePrototype{
				Href: core.StringPtr("cos://us-south/my-bucket/my-image.qcow2"),
			}

			operatingSystemIdentityModel := &vpcv1.OperatingSystemIdentityByName{
				Name: core.StringPtr("debian-9-amd64"),
			}

			imagePrototypeModel := &vpcv1.ImagePrototypeImageByFile{
				Name:               core.StringPtr("my-image"),
				RequiredImageFlags: imageRequiredImageFlagsModel,
				ResourceGroup:      resourceGroupIdentityModel,
				EncryptedDataKey:   core.StringPtr("testString"),
				EncryptionKey:      encryptionKeyIdentityModel,
				File:               imageFilePrototypeModel,
				OperatingSystem:    operatingSystemIdentityModel,
			}

			createImageOptions := &vpcv1.CreateImageOptions{
				ImagePrototype: imagePrototypeModel,
			}

			image, response, err := vpcService.CreateImage(createImageOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(image).ToNot(BeNil())

		})
	})

	Describe(`GetImage - Retrieve an image`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetImage(getImageOptions *GetImageOptions)`, func() {

			getImageOptions := &vpcv1.GetImageOptions{
				ID: core.StringPtr("testString"),
			}

			image, response, err := vpcService.GetImage(getImageOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(image).ToNot(BeNil())

		})
	})

	Describe(`UpdateImage - Update an image`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateImage(updateImageOptions *UpdateImageOptions)`, func() {

			imagePatchModel := &vpcv1.ImagePatch{
				Name: core.StringPtr("my-image"),
			}
			imagePatchModelAsPatch, asPatchErr := imagePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateImageOptions := &vpcv1.UpdateImageOptions{
				ID:         core.StringPtr("testString"),
				ImagePatch: imagePatchModelAsPatch,
			}

			image, response, err := vpcService.UpdateImage(updateImageOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(image).ToNot(BeNil())

		})
	})

	Describe(`ListOperatingSystems - List all operating systems`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListOperatingSystems(listOperatingSystemsOptions *ListOperatingSystemsOptions)`, func() {

			listOperatingSystemsOptions := &vpcv1.ListOperatingSystemsOptions{
				Start: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(1)),
			}

			operatingSystemCollection, response, err := vpcService.ListOperatingSystems(listOperatingSystemsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(operatingSystemCollection).ToNot(BeNil())

		})
	})

	Describe(`GetOperatingSystem - Retrieve an operating system`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetOperatingSystem(getOperatingSystemOptions *GetOperatingSystemOptions)`, func() {

			getOperatingSystemOptions := &vpcv1.GetOperatingSystemOptions{
				Name: core.StringPtr("testString"),
			}

			operatingSystem, response, err := vpcService.GetOperatingSystem(getOperatingSystemOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(operatingSystem).ToNot(BeNil())

		})
	})

	Describe(`ListKeys - List all keys`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListKeys(listKeysOptions *ListKeysOptions)`, func() {

			listKeysOptions := &vpcv1.ListKeysOptions{
				Start:           core.StringPtr("testString"),
				Limit:           core.Int64Ptr(int64(1)),
				ResourceGroupID: core.StringPtr("testString"),
			}

			keyCollection, response, err := vpcService.ListKeys(listKeysOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(keyCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateKey - Create a key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateKey(createKeyOptions *CreateKeyOptions)`, func() {

			resourceGroupIdentityModel := &vpcv1.ResourceGroupIdentityByID{
				ID: core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345"),
			}

			createKeyOptions := &vpcv1.CreateKeyOptions{
				PublicKey:     core.StringPtr("AAAAB3NzaC1yc2EAAAADAQABAAABAQDDGe50Bxa5T5NDddrrtbx2Y4/VGbiCgXqnBsYToIUKoFSHTQl5IX3PasGnneKanhcLwWz5M5MoCRvhxTp66NKzIfAz7r+FX9rxgR+ZgcM253YAqOVeIpOU408simDZKriTlN8kYsXL7P34tsWuAJf4MgZtJAQxous/2byetpdCv8ddnT4X3ltOg9w+LqSCPYfNivqH00Eh7S1Ldz7I8aw5WOp5a+sQFP/RbwfpwHp+ny7DfeIOokcuI42tJkoBn7UsLTVpCSmXr2EDRlSWe/1M/iHNRBzaT3CK0+SwZWd2AEjePxSnWKNGIEUJDlUYp7hKhiQcgT5ZAnWU121oc5En"),
				Name:          core.StringPtr("my-key"),
				ResourceGroup: resourceGroupIdentityModel,
				Type:          core.StringPtr("rsa"),
			}

			key, response, err := vpcService.CreateKey(createKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(key).ToNot(BeNil())

		})
	})

	Describe(`GetKey - Retrieve a key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetKey(getKeyOptions *GetKeyOptions)`, func() {

			getKeyOptions := &vpcv1.GetKeyOptions{
				ID: core.StringPtr("testString"),
			}

			key, response, err := vpcService.GetKey(getKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(key).ToNot(BeNil())

		})
	})

	Describe(`UpdateKey - Update a key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateKey(updateKeyOptions *UpdateKeyOptions)`, func() {

			keyPatchModel := &vpcv1.KeyPatch{
				Name: core.StringPtr("my-key"),
			}
			keyPatchModelAsPatch, asPatchErr := keyPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateKeyOptions := &vpcv1.UpdateKeyOptions{
				ID:       core.StringPtr("testString"),
				KeyPatch: keyPatchModelAsPatch,
			}

			key, response, err := vpcService.UpdateKey(updateKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(key).ToNot(BeNil())

		})
	})

	Describe(`ListInstanceProfiles - List all instance profiles`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListInstanceProfiles(listInstanceProfilesOptions *ListInstanceProfilesOptions)`, func() {

			listInstanceProfilesOptions := &vpcv1.ListInstanceProfilesOptions{
				Start: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(1)),
			}

			instanceProfileCollection, response, err := vpcService.ListInstanceProfiles(listInstanceProfilesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceProfileCollection).ToNot(BeNil())

		})
	})

	Describe(`GetInstanceProfile - Retrieve an instance profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInstanceProfile(getInstanceProfileOptions *GetInstanceProfileOptions)`, func() {

			getInstanceProfileOptions := &vpcv1.GetInstanceProfileOptions{
				Name: core.StringPtr("testString"),
			}

			instanceProfile, response, err := vpcService.GetInstanceProfile(getInstanceProfileOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceProfile).ToNot(BeNil())

		})
	})

	Describe(`ListInstanceTemplates - List all instance templates`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListInstanceTemplates(listInstanceTemplatesOptions *ListInstanceTemplatesOptions)`, func() {

			listInstanceTemplatesOptions := &vpcv1.ListInstanceTemplatesOptions{}

			instanceTemplateCollection, response, err := vpcService.ListInstanceTemplates(listInstanceTemplatesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceTemplateCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateInstanceTemplate - Create an instance template`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateInstanceTemplate(createInstanceTemplateOptions *CreateInstanceTemplateOptions)`, func() {

			keyIdentityModel := &vpcv1.KeyIdentityByID{
				ID: core.StringPtr("363f6d70-0000-0001-0000-00000013b96c"),
			}

			networkInterfaceIPPrototypeModel := &vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext{
				Address:    core.StringPtr("10.0.0.5"),
				AutoDelete: core.BoolPtr(false),
				Name:       core.StringPtr("my-reserved-ip"),
			}

			securityGroupIdentityModel := &vpcv1.SecurityGroupIdentityByID{
				ID: core.StringPtr("be5df5ca-12a0-494b-907e-aa6ec2bfa271"),
			}

			subnetIdentityModel := &vpcv1.SubnetIdentityByID{
				ID: core.StringPtr("7ec86020-1c6e-4889-b3f0-a15f2e50f87e"),
			}

			networkInterfacePrototypeModel := &vpcv1.NetworkInterfacePrototype{
				AllowIPSpoofing:    core.BoolPtr(true),
				Ips:                []vpcv1.NetworkInterfaceIPPrototypeIntf{networkInterfaceIPPrototypeModel},
				Name:               core.StringPtr("my-network-interface"),
				PrimaryIpv4Address: core.StringPtr("10.0.0.5"),
				SecurityGroups:     []vpcv1.SecurityGroupIdentityIntf{securityGroupIdentityModel},
				Subnet:             subnetIdentityModel,
			}

			instancePlacementTargetPatchModel := &vpcv1.InstancePlacementTargetPatchDedicatedHostIdentityDedicatedHostIdentityByID{
				ID: core.StringPtr("1e09281b-f177-46fb-baf1-bc152b2e391a"),
			}

			instanceProfileIdentityModel := &vpcv1.InstanceProfileIdentityByName{
				Name: core.StringPtr("bx2-2x8"),
			}

			resourceGroupIdentityModel := &vpcv1.ResourceGroupIdentityByID{
				ID: core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345"),
			}

			volumeAttachmentPrototypeVolumeModel := &vpcv1.VolumeAttachmentPrototypeVolumeVolumeIdentityVolumeIdentityByID{
				ID: core.StringPtr("1a6b7274-678d-4dfb-8981-c71dd9d4daa5"),
			}

			volumeAttachmentPrototypeModel := &vpcv1.VolumeAttachmentPrototype{
				DeleteVolumeOnInstanceDelete: core.BoolPtr(true),
				Name:                         core.StringPtr("my-volume-attachment"),
				Volume:                       volumeAttachmentPrototypeVolumeModel,
			}

			vpcIdentityModel := &vpcv1.VPCIdentityByID{
				ID: core.StringPtr("dc201ab2-8536-4904-86a8-084d84582133"),
			}

			encryptionKeyIdentityModel := &vpcv1.EncryptionKeyIdentityByCRN{
				CRN: core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/dffc98a0f1f0f95f6613b3b752286b87:e4a29d1a-2ef0-42a6-8fd2-350deb1c647e:key:5437653b-c4b1-447f-9646-b2a2a4cd6179"),
			}

			volumeProfileIdentityModel := &vpcv1.VolumeProfileIdentityByName{
				Name: core.StringPtr("general-purpose"),
			}

			volumePrototypeInstanceByImageContextModel := &vpcv1.VolumePrototypeInstanceByImageContext{
				Capacity:      core.Int64Ptr(int64(100)),
				EncryptionKey: encryptionKeyIdentityModel,
				Iops:          core.Int64Ptr(int64(10000)),
				Name:          core.StringPtr("my-volume"),
				Profile:       volumeProfileIdentityModel,
				ResourceGroup: resourceGroupIdentityModel,
			}

			volumeAttachmentPrototypeInstanceByImageContextModel := &vpcv1.VolumeAttachmentPrototypeInstanceByImageContext{
				DeleteVolumeOnInstanceDelete: core.BoolPtr(true),
				Name:                         core.StringPtr("my-volume-attachment"),
				Volume:                       volumePrototypeInstanceByImageContextModel,
			}

			imageIdentityModel := &vpcv1.ImageIdentityByID{
				ID: core.StringPtr("3f9a2d96-830e-4100-9b4c-663225a3f872"),
			}

			zoneIdentityModel := &vpcv1.ZoneIdentityByName{
				Name: core.StringPtr("us-south-1"),
			}

			instanceTemplatePrototypeModel := &vpcv1.InstanceTemplatePrototypeInstanceByImage{
				Keys:                    []vpcv1.KeyIdentityIntf{keyIdentityModel},
				Name:                    core.StringPtr("my-instance-template"),
				NetworkInterfaces:       []vpcv1.NetworkInterfacePrototype{*networkInterfacePrototypeModel},
				PlacementTarget:         instancePlacementTargetPatchModel,
				Profile:                 instanceProfileIdentityModel,
				ResourceGroup:           resourceGroupIdentityModel,
				UserData:                core.StringPtr("testString"),
				VolumeAttachments:       []vpcv1.VolumeAttachmentPrototype{*volumeAttachmentPrototypeModel},
				VPC:                     vpcIdentityModel,
				BootVolumeAttachment:    volumeAttachmentPrototypeInstanceByImageContextModel,
				Image:                   imageIdentityModel,
				PrimaryNetworkInterface: networkInterfacePrototypeModel,
				Zone:                    zoneIdentityModel,
			}

			createInstanceTemplateOptions := &vpcv1.CreateInstanceTemplateOptions{
				InstanceTemplatePrototype: instanceTemplatePrototypeModel,
			}

			instanceTemplate, response, err := vpcService.CreateInstanceTemplate(createInstanceTemplateOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(instanceTemplate).ToNot(BeNil())

		})
	})

	Describe(`GetInstanceTemplate - Retrieve an instance template`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInstanceTemplate(getInstanceTemplateOptions *GetInstanceTemplateOptions)`, func() {

			getInstanceTemplateOptions := &vpcv1.GetInstanceTemplateOptions{
				ID: core.StringPtr("testString"),
			}

			instanceTemplate, response, err := vpcService.GetInstanceTemplate(getInstanceTemplateOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceTemplate).ToNot(BeNil())

		})
	})

	Describe(`UpdateInstanceTemplate - Update an instance template`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateInstanceTemplate(updateInstanceTemplateOptions *UpdateInstanceTemplateOptions)`, func() {

			instanceTemplatePatchModel := &vpcv1.InstanceTemplatePatch{
				Name: core.StringPtr("my-instance-template"),
			}
			instanceTemplatePatchModelAsPatch, asPatchErr := instanceTemplatePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateInstanceTemplateOptions := &vpcv1.UpdateInstanceTemplateOptions{
				ID:                    core.StringPtr("testString"),
				InstanceTemplatePatch: instanceTemplatePatchModelAsPatch,
			}

			instanceTemplate, response, err := vpcService.UpdateInstanceTemplate(updateInstanceTemplateOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceTemplate).ToNot(BeNil())

		})
	})

	Describe(`ListInstances - List all instances`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListInstances(listInstancesOptions *ListInstancesOptions)`, func() {

			listInstancesOptions := &vpcv1.ListInstancesOptions{
				Start:                       core.StringPtr("testString"),
				Limit:                       core.Int64Ptr(int64(1)),
				ResourceGroupID:             core.StringPtr("testString"),
				Name:                        core.StringPtr("testString"),
				VPCID:                       core.StringPtr("testString"),
				VPCCRN:                      core.StringPtr("testString"),
				VPCName:                     core.StringPtr("testString"),
				NetworkInterfacesSubnetID:   core.StringPtr("testString"),
				NetworkInterfacesSubnetCRN:  core.StringPtr("testString"),
				NetworkInterfacesSubnetName: core.StringPtr("testString"),
				DedicatedHostID:             core.StringPtr("testString"),
				DedicatedHostCRN:            core.StringPtr("testString"),
				DedicatedHostName:           core.StringPtr("testString"),
				PlacementGroupID:            core.StringPtr("testString"),
				PlacementGroupCRN:           core.StringPtr("testString"),
				PlacementGroupName:          core.StringPtr("testString"),
			}

			instanceCollection, response, err := vpcService.ListInstances(listInstancesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateInstance - Create an instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateInstance(createInstanceOptions *CreateInstanceOptions)`, func() {

			keyIdentityModel := &vpcv1.KeyIdentityByID{
				ID: core.StringPtr("363f6d70-0000-0001-0000-00000013b96c"),
			}

			networkInterfaceIPPrototypeModel := &vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext{
				Address:    core.StringPtr("10.0.0.5"),
				AutoDelete: core.BoolPtr(false),
				Name:       core.StringPtr("my-reserved-ip"),
			}

			securityGroupIdentityModel := &vpcv1.SecurityGroupIdentityByID{
				ID: core.StringPtr("be5df5ca-12a0-494b-907e-aa6ec2bfa271"),
			}

			subnetIdentityModel := &vpcv1.SubnetIdentityByID{
				ID: core.StringPtr("7ec86020-1c6e-4889-b3f0-a15f2e50f87e"),
			}

			networkInterfacePrototypeModel := &vpcv1.NetworkInterfacePrototype{
				AllowIPSpoofing:    core.BoolPtr(true),
				Ips:                []vpcv1.NetworkInterfaceIPPrototypeIntf{networkInterfaceIPPrototypeModel},
				Name:               core.StringPtr("my-network-interface"),
				PrimaryIpv4Address: core.StringPtr("10.0.0.5"),
				SecurityGroups:     []vpcv1.SecurityGroupIdentityIntf{securityGroupIdentityModel},
				Subnet:             subnetIdentityModel,
			}

			instancePlacementTargetPatchModel := &vpcv1.InstancePlacementTargetPatchDedicatedHostIdentityDedicatedHostIdentityByID{
				ID: core.StringPtr("0787-8c2a09be-ee18-4af2-8ef4-6a6060732221"),
			}

			instanceProfileIdentityModel := &vpcv1.InstanceProfileIdentityByName{
				Name: core.StringPtr("bx2-2x8"),
			}

			resourceGroupIdentityModel := &vpcv1.ResourceGroupIdentityByID{
				ID: core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345"),
			}

			encryptionKeyIdentityModel := &vpcv1.EncryptionKeyIdentityByCRN{
				CRN: core.StringPtr("crn:[...]"),
			}

			volumeProfileIdentityModel := &vpcv1.VolumeProfileIdentityByName{
				Name: core.StringPtr("5iops-tier"),
			}

			volumeAttachmentPrototypeVolumeModel := &vpcv1.VolumeAttachmentPrototypeVolumeVolumePrototypeInstanceContextVolumePrototypeInstanceContextVolumeByCapacity{
				EncryptionKey: encryptionKeyIdentityModel,
				Iops:          core.Int64Ptr(int64(10000)),
				Name:          core.StringPtr("my-data-volume"),
				Profile:       volumeProfileIdentityModel,
				ResourceGroup: resourceGroupIdentityModel,
				Capacity:      core.Int64Ptr(int64(1000)),
			}

			volumeAttachmentPrototypeModel := &vpcv1.VolumeAttachmentPrototype{
				DeleteVolumeOnInstanceDelete: core.BoolPtr(true),
				Name:                         core.StringPtr("my-volume-attachment"),
				Volume:                       volumeAttachmentPrototypeVolumeModel,
			}

			vpcIdentityModel := &vpcv1.VPCIdentityByID{
				ID: core.StringPtr("f0aae929-7047-46d1-92e1-9102b07a7f6f"),
			}

			volumePrototypeInstanceByImageContextModel := &vpcv1.VolumePrototypeInstanceByImageContext{
				Capacity:      core.Int64Ptr(int64(100)),
				EncryptionKey: encryptionKeyIdentityModel,
				Iops:          core.Int64Ptr(int64(10000)),
				Name:          core.StringPtr("my-boot-volume"),
				Profile:       volumeProfileIdentityModel,
				ResourceGroup: resourceGroupIdentityModel,
			}

			volumeAttachmentPrototypeInstanceByImageContextModel := &vpcv1.VolumeAttachmentPrototypeInstanceByImageContext{
				DeleteVolumeOnInstanceDelete: core.BoolPtr(true),
				Name:                         core.StringPtr("my-volume-attachment"),
				Volume:                       volumePrototypeInstanceByImageContextModel,
			}

			imageIdentityModel := &vpcv1.ImageIdentityByID{
				ID: core.StringPtr("9aaf3bcb-dcd7-4de7-bb60-24e39ff9d366"),
			}

			zoneIdentityModel := &vpcv1.ZoneIdentityByName{
				Name: core.StringPtr("us-south-1"),
			}

			instancePrototypeModel := &vpcv1.InstancePrototypeInstanceByImage{
				Keys:                    []vpcv1.KeyIdentityIntf{keyIdentityModel},
				Name:                    core.StringPtr("my-instance"),
				NetworkInterfaces:       []vpcv1.NetworkInterfacePrototype{*networkInterfacePrototypeModel},
				PlacementTarget:         instancePlacementTargetPatchModel,
				Profile:                 instanceProfileIdentityModel,
				ResourceGroup:           resourceGroupIdentityModel,
				UserData:                core.StringPtr("testString"),
				VolumeAttachments:       []vpcv1.VolumeAttachmentPrototype{*volumeAttachmentPrototypeModel},
				VPC:                     vpcIdentityModel,
				BootVolumeAttachment:    volumeAttachmentPrototypeInstanceByImageContextModel,
				Image:                   imageIdentityModel,
				PrimaryNetworkInterface: networkInterfacePrototypeModel,
				Zone:                    zoneIdentityModel,
			}

			createInstanceOptions := &vpcv1.CreateInstanceOptions{
				InstancePrototype: instancePrototypeModel,
			}

			instance, response, err := vpcService.CreateInstance(createInstanceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(instance).ToNot(BeNil())

		})
	})

	Describe(`GetInstance - Retrieve an instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInstance(getInstanceOptions *GetInstanceOptions)`, func() {

			getInstanceOptions := &vpcv1.GetInstanceOptions{
				ID: core.StringPtr("testString"),
			}

			instance, response, err := vpcService.GetInstance(getInstanceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instance).ToNot(BeNil())

		})
	})

	Describe(`UpdateInstance - Update an instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateInstance(updateInstanceOptions *UpdateInstanceOptions)`, func() {

			instancePlacementTargetPatchModel := &vpcv1.InstancePlacementTargetPatchDedicatedHostIdentityDedicatedHostIdentityByID{
				ID: core.StringPtr("1e09281b-f177-46fb-baf1-bc152b2e391a"),
			}

			instancePatchProfileModel := &vpcv1.InstancePatchProfileInstanceProfileIdentityByName{
				Name: core.StringPtr("bc1-4x16"),
			}

			instancePatchModel := &vpcv1.InstancePatch{
				Name:            core.StringPtr("my-instance"),
				PlacementTarget: instancePlacementTargetPatchModel,
				Profile:         instancePatchProfileModel,
			}
			instancePatchModelAsPatch, asPatchErr := instancePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateInstanceOptions := &vpcv1.UpdateInstanceOptions{
				ID:            core.StringPtr("testString"),
				InstancePatch: instancePatchModelAsPatch,
			}

			instance, response, err := vpcService.UpdateInstance(updateInstanceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instance).ToNot(BeNil())

		})
	})

	Describe(`GetInstanceInitialization - Retrieve initialization configuration for an instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInstanceInitialization(getInstanceInitializationOptions *GetInstanceInitializationOptions)`, func() {

			getInstanceInitializationOptions := &vpcv1.GetInstanceInitializationOptions{
				ID: core.StringPtr("testString"),
			}

			instanceInitialization, response, err := vpcService.GetInstanceInitialization(getInstanceInitializationOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceInitialization).ToNot(BeNil())

		})
	})

	Describe(`CreateInstanceAction - Create an instance action`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateInstanceAction(createInstanceActionOptions *CreateInstanceActionOptions)`, func() {

			createInstanceActionOptions := &vpcv1.CreateInstanceActionOptions{
				InstanceID: core.StringPtr("testString"),
				Type:       core.StringPtr("reboot"),
				Force:      core.BoolPtr(true),
			}

			instanceAction, response, err := vpcService.CreateInstanceAction(createInstanceActionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(instanceAction).ToNot(BeNil())

		})
	})

	Describe(`GetInstanceConsole - Retrieve the console WebSocket for an instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInstanceConsole(getInstanceConsoleOptions *GetInstanceConsoleOptions)`, func() {

			getInstanceConsoleOptions := &vpcv1.GetInstanceConsoleOptions{
				InstanceID:  core.StringPtr("testString"),
				AccessToken: core.StringPtr("VGhpcyBJcyBhIHRva2Vu"),
			}

			response, err := vpcService.GetInstanceConsole(getInstanceConsoleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})

	Describe(`CreateInstanceConsoleAccessToken - Create a console access token for an instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateInstanceConsoleAccessToken(createInstanceConsoleAccessTokenOptions *CreateInstanceConsoleAccessTokenOptions)`, func() {

			createInstanceConsoleAccessTokenOptions := &vpcv1.CreateInstanceConsoleAccessTokenOptions{
				InstanceID:  core.StringPtr("testString"),
				ConsoleType: core.StringPtr("serial"),
				Force:       core.BoolPtr(false),
			}

			instanceConsoleAccessToken, response, err := vpcService.CreateInstanceConsoleAccessToken(createInstanceConsoleAccessTokenOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceConsoleAccessToken).ToNot(BeNil())

		})
	})

	Describe(`ListInstanceDisks - List all disks on an instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListInstanceDisks(listInstanceDisksOptions *ListInstanceDisksOptions)`, func() {

			listInstanceDisksOptions := &vpcv1.ListInstanceDisksOptions{
				InstanceID: core.StringPtr("testString"),
			}

			instanceDiskCollection, response, err := vpcService.ListInstanceDisks(listInstanceDisksOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceDiskCollection).ToNot(BeNil())

		})
	})

	Describe(`GetInstanceDisk - Retrieve an instance disk`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInstanceDisk(getInstanceDiskOptions *GetInstanceDiskOptions)`, func() {

			getInstanceDiskOptions := &vpcv1.GetInstanceDiskOptions{
				InstanceID: core.StringPtr("testString"),
				ID:         core.StringPtr("testString"),
			}

			instanceDisk, response, err := vpcService.GetInstanceDisk(getInstanceDiskOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceDisk).ToNot(BeNil())

		})
	})

	Describe(`UpdateInstanceDisk - Update an instance disk`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateInstanceDisk(updateInstanceDiskOptions *UpdateInstanceDiskOptions)`, func() {

			instanceDiskPatchModel := &vpcv1.InstanceDiskPatch{
				Name: core.StringPtr("my-instance-disk-updated"),
			}
			instanceDiskPatchModelAsPatch, asPatchErr := instanceDiskPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateInstanceDiskOptions := &vpcv1.UpdateInstanceDiskOptions{
				InstanceID:        core.StringPtr("testString"),
				ID:                core.StringPtr("testString"),
				InstanceDiskPatch: instanceDiskPatchModelAsPatch,
			}

			instanceDisk, response, err := vpcService.UpdateInstanceDisk(updateInstanceDiskOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceDisk).ToNot(BeNil())

		})
	})

	Describe(`ListInstanceNetworkInterfaces - List all network interfaces on an instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListInstanceNetworkInterfaces(listInstanceNetworkInterfacesOptions *ListInstanceNetworkInterfacesOptions)`, func() {

			listInstanceNetworkInterfacesOptions := &vpcv1.ListInstanceNetworkInterfacesOptions{
				InstanceID: core.StringPtr("testString"),
			}

			networkInterfaceUnpaginatedCollection, response, err := vpcService.ListInstanceNetworkInterfaces(listInstanceNetworkInterfacesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkInterfaceUnpaginatedCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateInstanceNetworkInterface - Create a network interface on an instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateInstanceNetworkInterface(createInstanceNetworkInterfaceOptions *CreateInstanceNetworkInterfaceOptions)`, func() {

			subnetIdentityModel := &vpcv1.SubnetIdentityByID{
				ID: core.StringPtr("7ec86020-1c6e-4889-b3f0-a15f2e50f87e"),
			}

			networkInterfaceIPPrototypeModel := &vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext{
				Address:    core.StringPtr("10.0.0.5"),
				AutoDelete: core.BoolPtr(false),
				Name:       core.StringPtr("my-reserved-ip"),
			}

			securityGroupIdentityModel := &vpcv1.SecurityGroupIdentityByID{
				ID: core.StringPtr("be5df5ca-12a0-494b-907e-aa6ec2bfa271"),
			}

			createInstanceNetworkInterfaceOptions := &vpcv1.CreateInstanceNetworkInterfaceOptions{
				InstanceID:         core.StringPtr("testString"),
				Subnet:             subnetIdentityModel,
				AllowIPSpoofing:    core.BoolPtr(true),
				Ips:                []vpcv1.NetworkInterfaceIPPrototypeIntf{networkInterfaceIPPrototypeModel},
				Name:               core.StringPtr("my-network-interface"),
				PrimaryIpv4Address: core.StringPtr("10.0.0.5"),
				SecurityGroups:     []vpcv1.SecurityGroupIdentityIntf{securityGroupIdentityModel},
			}

			networkInterface, response, err := vpcService.CreateInstanceNetworkInterface(createInstanceNetworkInterfaceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(networkInterface).ToNot(BeNil())

		})
	})

	Describe(`GetInstanceNetworkInterface - Retrieve a network interface`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInstanceNetworkInterface(getInstanceNetworkInterfaceOptions *GetInstanceNetworkInterfaceOptions)`, func() {

			getInstanceNetworkInterfaceOptions := &vpcv1.GetInstanceNetworkInterfaceOptions{
				InstanceID: core.StringPtr("testString"),
				ID:         core.StringPtr("testString"),
			}

			networkInterface, response, err := vpcService.GetInstanceNetworkInterface(getInstanceNetworkInterfaceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkInterface).ToNot(BeNil())

		})
	})

	Describe(`UpdateInstanceNetworkInterface - Update a network interface`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateInstanceNetworkInterface(updateInstanceNetworkInterfaceOptions *UpdateInstanceNetworkInterfaceOptions)`, func() {

			networkInterfacePatchModel := &vpcv1.NetworkInterfacePatch{
				AllowIPSpoofing: core.BoolPtr(true),
				Name:            core.StringPtr("my-network-interface-1"),
			}
			networkInterfacePatchModelAsPatch, asPatchErr := networkInterfacePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateInstanceNetworkInterfaceOptions := &vpcv1.UpdateInstanceNetworkInterfaceOptions{
				InstanceID:            core.StringPtr("testString"),
				ID:                    core.StringPtr("testString"),
				NetworkInterfacePatch: networkInterfacePatchModelAsPatch,
			}

			networkInterface, response, err := vpcService.UpdateInstanceNetworkInterface(updateInstanceNetworkInterfaceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkInterface).ToNot(BeNil())

		})
	})

	Describe(`ListInstanceNetworkInterfaceFloatingIps - List all floating IPs associated with a network interface`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListInstanceNetworkInterfaceFloatingIps(listInstanceNetworkInterfaceFloatingIpsOptions *ListInstanceNetworkInterfaceFloatingIpsOptions)`, func() {

			listInstanceNetworkInterfaceFloatingIpsOptions := &vpcv1.ListInstanceNetworkInterfaceFloatingIpsOptions{
				InstanceID:         core.StringPtr("testString"),
				NetworkInterfaceID: core.StringPtr("testString"),
			}

			floatingIPUnpaginatedCollection, response, err := vpcService.ListInstanceNetworkInterfaceFloatingIps(listInstanceNetworkInterfaceFloatingIpsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(floatingIPUnpaginatedCollection).ToNot(BeNil())

		})
	})

	Describe(`GetInstanceNetworkInterfaceFloatingIP - Retrieve associated floating IP`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInstanceNetworkInterfaceFloatingIP(getInstanceNetworkInterfaceFloatingIPOptions *GetInstanceNetworkInterfaceFloatingIPOptions)`, func() {

			getInstanceNetworkInterfaceFloatingIPOptions := &vpcv1.GetInstanceNetworkInterfaceFloatingIPOptions{
				InstanceID:         core.StringPtr("testString"),
				NetworkInterfaceID: core.StringPtr("testString"),
				ID:                 core.StringPtr("testString"),
			}

			floatingIP, response, err := vpcService.GetInstanceNetworkInterfaceFloatingIP(getInstanceNetworkInterfaceFloatingIPOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(floatingIP).ToNot(BeNil())

		})
	})

	Describe(`AddInstanceNetworkInterfaceFloatingIP - Associate a floating IP with a network interface`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AddInstanceNetworkInterfaceFloatingIP(addInstanceNetworkInterfaceFloatingIPOptions *AddInstanceNetworkInterfaceFloatingIPOptions)`, func() {

			addInstanceNetworkInterfaceFloatingIPOptions := &vpcv1.AddInstanceNetworkInterfaceFloatingIPOptions{
				InstanceID:         core.StringPtr("testString"),
				NetworkInterfaceID: core.StringPtr("testString"),
				ID:                 core.StringPtr("testString"),
			}

			floatingIP, response, err := vpcService.AddInstanceNetworkInterfaceFloatingIP(addInstanceNetworkInterfaceFloatingIPOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(floatingIP).ToNot(BeNil())

		})
	})

	Describe(`ListInstanceVolumeAttachments - List all volumes attachments on an instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListInstanceVolumeAttachments(listInstanceVolumeAttachmentsOptions *ListInstanceVolumeAttachmentsOptions)`, func() {

			listInstanceVolumeAttachmentsOptions := &vpcv1.ListInstanceVolumeAttachmentsOptions{
				InstanceID: core.StringPtr("testString"),
			}

			volumeAttachmentCollection, response, err := vpcService.ListInstanceVolumeAttachments(listInstanceVolumeAttachmentsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeAttachmentCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateInstanceVolumeAttachment - Create a volume attachment on an instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateInstanceVolumeAttachment(createInstanceVolumeAttachmentOptions *CreateInstanceVolumeAttachmentOptions)`, func() {

			volumeAttachmentPrototypeVolumeModel := &vpcv1.VolumeAttachmentPrototypeVolumeVolumeIdentityVolumeIdentityByID{
				ID: core.StringPtr("1a6b7274-678d-4dfb-8981-c71dd9d4daa5"),
			}

			createInstanceVolumeAttachmentOptions := &vpcv1.CreateInstanceVolumeAttachmentOptions{
				InstanceID:                   core.StringPtr("testString"),
				Volume:                       volumeAttachmentPrototypeVolumeModel,
				DeleteVolumeOnInstanceDelete: core.BoolPtr(true),
				Name:                         core.StringPtr("my-volume-attachment"),
			}

			volumeAttachment, response, err := vpcService.CreateInstanceVolumeAttachment(createInstanceVolumeAttachmentOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(volumeAttachment).ToNot(BeNil())

		})
	})

	Describe(`GetInstanceVolumeAttachment - Retrieve a volume attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInstanceVolumeAttachment(getInstanceVolumeAttachmentOptions *GetInstanceVolumeAttachmentOptions)`, func() {

			getInstanceVolumeAttachmentOptions := &vpcv1.GetInstanceVolumeAttachmentOptions{
				InstanceID: core.StringPtr("testString"),
				ID:         core.StringPtr("testString"),
			}

			volumeAttachment, response, err := vpcService.GetInstanceVolumeAttachment(getInstanceVolumeAttachmentOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeAttachment).ToNot(BeNil())

		})
	})

	Describe(`UpdateInstanceVolumeAttachment - Update a volume attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateInstanceVolumeAttachment(updateInstanceVolumeAttachmentOptions *UpdateInstanceVolumeAttachmentOptions)`, func() {

			volumeAttachmentPatchModel := &vpcv1.VolumeAttachmentPatch{
				DeleteVolumeOnInstanceDelete: core.BoolPtr(true),
				Name:                         core.StringPtr("my-volume-attachment"),
			}
			volumeAttachmentPatchModelAsPatch, asPatchErr := volumeAttachmentPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateInstanceVolumeAttachmentOptions := &vpcv1.UpdateInstanceVolumeAttachmentOptions{
				InstanceID:            core.StringPtr("testString"),
				ID:                    core.StringPtr("testString"),
				VolumeAttachmentPatch: volumeAttachmentPatchModelAsPatch,
			}

			volumeAttachment, response, err := vpcService.UpdateInstanceVolumeAttachment(updateInstanceVolumeAttachmentOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeAttachment).ToNot(BeNil())

		})
	})

	Describe(`ListInstanceGroups - List all instance groups`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListInstanceGroups(listInstanceGroupsOptions *ListInstanceGroupsOptions)`, func() {

			listInstanceGroupsOptions := &vpcv1.ListInstanceGroupsOptions{
				Start: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(1)),
			}

			instanceGroupCollection, response, err := vpcService.ListInstanceGroups(listInstanceGroupsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateInstanceGroup - Create an instance group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateInstanceGroup(createInstanceGroupOptions *CreateInstanceGroupOptions)`, func() {

			instanceTemplateIdentityModel := &vpcv1.InstanceTemplateIdentityByID{
				ID: core.StringPtr("a6b1a881-2ce8-41a3-80fc-36316a73f803"),
			}

			subnetIdentityModel := &vpcv1.SubnetIdentityByID{
				ID: core.StringPtr("7ec86020-1c6e-4889-b3f0-a15f2e50f87e"),
			}

			loadBalancerIdentityModel := &vpcv1.LoadBalancerIdentityByID{
				ID: core.StringPtr("dd754295-e9e0-4c9d-bf6c-58fbc59e5727"),
			}

			loadBalancerPoolIdentityModel := &vpcv1.LoadBalancerPoolIdentityByID{
				ID: core.StringPtr("70294e14-4e61-11e8-bcf4-0242ac110004"),
			}

			resourceGroupIdentityModel := &vpcv1.ResourceGroupIdentityByID{
				ID: core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345"),
			}

			createInstanceGroupOptions := &vpcv1.CreateInstanceGroupOptions{
				InstanceTemplate: instanceTemplateIdentityModel,
				Subnets:          []vpcv1.SubnetIdentityIntf{subnetIdentityModel},
				ApplicationPort:  core.Int64Ptr(int64(22)),
				LoadBalancer:     loadBalancerIdentityModel,
				LoadBalancerPool: loadBalancerPoolIdentityModel,
				MembershipCount:  core.Int64Ptr(int64(10)),
				Name:             core.StringPtr("my-instance-group"),
				ResourceGroup:    resourceGroupIdentityModel,
			}

			instanceGroup, response, err := vpcService.CreateInstanceGroup(createInstanceGroupOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(instanceGroup).ToNot(BeNil())

		})
	})

	Describe(`GetInstanceGroup - Retrieve an instance group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInstanceGroup(getInstanceGroupOptions *GetInstanceGroupOptions)`, func() {

			getInstanceGroupOptions := &vpcv1.GetInstanceGroupOptions{
				ID: core.StringPtr("testString"),
			}

			instanceGroup, response, err := vpcService.GetInstanceGroup(getInstanceGroupOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroup).ToNot(BeNil())

		})
	})

	Describe(`UpdateInstanceGroup - Update an instance group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateInstanceGroup(updateInstanceGroupOptions *UpdateInstanceGroupOptions)`, func() {

			instanceTemplateIdentityModel := &vpcv1.InstanceTemplateIdentityByID{
				ID: core.StringPtr("a6b1a881-2ce8-41a3-80fc-36316a73f803"),
			}

			loadBalancerIdentityModel := &vpcv1.LoadBalancerIdentityByID{
				ID: core.StringPtr("dd754295-e9e0-4c9d-bf6c-58fbc59e5727"),
			}

			loadBalancerPoolIdentityModel := &vpcv1.LoadBalancerPoolIdentityByID{
				ID: core.StringPtr("70294e14-4e61-11e8-bcf4-0242ac110004"),
			}

			subnetIdentityModel := &vpcv1.SubnetIdentityByID{
				ID: core.StringPtr("7ec86020-1c6e-4889-b3f0-a15f2e50f87e"),
			}

			instanceGroupPatchModel := &vpcv1.InstanceGroupPatch{
				ApplicationPort:  core.Int64Ptr(int64(22)),
				InstanceTemplate: instanceTemplateIdentityModel,
				LoadBalancer:     loadBalancerIdentityModel,
				LoadBalancerPool: loadBalancerPoolIdentityModel,
				MembershipCount:  core.Int64Ptr(int64(10)),
				Name:             core.StringPtr("my-instance-group"),
				Subnets:          []vpcv1.SubnetIdentityIntf{subnetIdentityModel},
			}
			instanceGroupPatchModelAsPatch, asPatchErr := instanceGroupPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateInstanceGroupOptions := &vpcv1.UpdateInstanceGroupOptions{
				ID:                 core.StringPtr("testString"),
				InstanceGroupPatch: instanceGroupPatchModelAsPatch,
			}

			instanceGroup, response, err := vpcService.UpdateInstanceGroup(updateInstanceGroupOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroup).ToNot(BeNil())

		})
	})

	Describe(`ListInstanceGroupManagers - List all managers for an instance group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListInstanceGroupManagers(listInstanceGroupManagersOptions *ListInstanceGroupManagersOptions)`, func() {

			listInstanceGroupManagersOptions := &vpcv1.ListInstanceGroupManagersOptions{
				InstanceGroupID: core.StringPtr("testString"),
				Start:           core.StringPtr("testString"),
				Limit:           core.Int64Ptr(int64(1)),
			}

			instanceGroupManagerCollection, response, err := vpcService.ListInstanceGroupManagers(listInstanceGroupManagersOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManagerCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateInstanceGroupManager - Create a manager for an instance group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateInstanceGroupManager(createInstanceGroupManagerOptions *CreateInstanceGroupManagerOptions)`, func() {

			instanceGroupManagerPrototypeModel := &vpcv1.InstanceGroupManagerPrototypeInstanceGroupManagerAutoScalePrototype{
				ManagementEnabled:  core.BoolPtr(true),
				Name:               core.StringPtr("my-instance-group-manager"),
				AggregationWindow:  core.Int64Ptr(int64(120)),
				Cooldown:           core.Int64Ptr(int64(210)),
				ManagerType:        core.StringPtr("autoscale"),
				MaxMembershipCount: core.Int64Ptr(int64(10)),
				MinMembershipCount: core.Int64Ptr(int64(10)),
			}

			createInstanceGroupManagerOptions := &vpcv1.CreateInstanceGroupManagerOptions{
				InstanceGroupID:               core.StringPtr("testString"),
				InstanceGroupManagerPrototype: instanceGroupManagerPrototypeModel,
			}

			instanceGroupManager, response, err := vpcService.CreateInstanceGroupManager(createInstanceGroupManagerOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(instanceGroupManager).ToNot(BeNil())

		})
	})

	Describe(`GetInstanceGroupManager - Retrieve an instance group manager`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInstanceGroupManager(getInstanceGroupManagerOptions *GetInstanceGroupManagerOptions)`, func() {

			getInstanceGroupManagerOptions := &vpcv1.GetInstanceGroupManagerOptions{
				InstanceGroupID: core.StringPtr("testString"),
				ID:              core.StringPtr("testString"),
			}

			instanceGroupManager, response, err := vpcService.GetInstanceGroupManager(getInstanceGroupManagerOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManager).ToNot(BeNil())

		})
	})

	Describe(`UpdateInstanceGroupManager - Update an instance group manager`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateInstanceGroupManager(updateInstanceGroupManagerOptions *UpdateInstanceGroupManagerOptions)`, func() {

			instanceGroupManagerPatchModel := &vpcv1.InstanceGroupManagerPatch{
				AggregationWindow:  core.Int64Ptr(int64(120)),
				Cooldown:           core.Int64Ptr(int64(210)),
				ManagementEnabled:  core.BoolPtr(true),
				MaxMembershipCount: core.Int64Ptr(int64(10)),
				MinMembershipCount: core.Int64Ptr(int64(10)),
				Name:               core.StringPtr("my-instance-group-manager"),
			}
			instanceGroupManagerPatchModelAsPatch, asPatchErr := instanceGroupManagerPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateInstanceGroupManagerOptions := &vpcv1.UpdateInstanceGroupManagerOptions{
				InstanceGroupID:           core.StringPtr("testString"),
				ID:                        core.StringPtr("testString"),
				InstanceGroupManagerPatch: instanceGroupManagerPatchModelAsPatch,
			}

			instanceGroupManager, response, err := vpcService.UpdateInstanceGroupManager(updateInstanceGroupManagerOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManager).ToNot(BeNil())

		})
	})

	Describe(`ListInstanceGroupManagerActions - List all actions for an instance group manager`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListInstanceGroupManagerActions(listInstanceGroupManagerActionsOptions *ListInstanceGroupManagerActionsOptions)`, func() {

			listInstanceGroupManagerActionsOptions := &vpcv1.ListInstanceGroupManagerActionsOptions{
				InstanceGroupID:        core.StringPtr("testString"),
				InstanceGroupManagerID: core.StringPtr("testString"),
				Start:                  core.StringPtr("testString"),
				Limit:                  core.Int64Ptr(int64(1)),
			}

			instanceGroupManagerActionsCollection, response, err := vpcService.ListInstanceGroupManagerActions(listInstanceGroupManagerActionsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManagerActionsCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateInstanceGroupManagerAction - Create an instance group manager action`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateInstanceGroupManagerAction(createInstanceGroupManagerActionOptions *CreateInstanceGroupManagerActionOptions)`, func() {

			instanceGroupManagerScheduledActionByManagerManagerModel := &vpcv1.InstanceGroupManagerScheduledActionByManagerManagerInstanceGroupManagerScheduledActionManagerAutoScalePrototypeInstanceGroupManagerScheduledActionManagerAutoScalePrototypeInstanceGroupManagerIdentityByID{
				MaxMembershipCount: core.Int64Ptr(int64(10)),
				MinMembershipCount: core.Int64Ptr(int64(10)),
				ID:                 core.StringPtr("1e09281b-f177-46fb-baf1-bc152b2e391a"),
			}

			instanceGroupManagerActionPrototypeModel := &vpcv1.InstanceGroupManagerActionPrototypeInstanceGroupManagerScheduledActionPrototypeInstanceGroupManagerScheduledActionPrototypeInstanceGroupManagerScheduledActionByRunAtInstanceGroupManagerScheduledActionPrototypeInstanceGroupManagerScheduledActionByRunAtInstanceGroupManagerScheduledActionByRunAtInstanceGroupManagerScheduledActionByManager{
				Name:    core.StringPtr("my-instance-group-manager-action"),
				RunAt:   CreateMockDateTime(),
				Manager: instanceGroupManagerScheduledActionByManagerManagerModel,
			}

			createInstanceGroupManagerActionOptions := &vpcv1.CreateInstanceGroupManagerActionOptions{
				InstanceGroupID:                     core.StringPtr("testString"),
				InstanceGroupManagerID:              core.StringPtr("testString"),
				InstanceGroupManagerActionPrototype: instanceGroupManagerActionPrototypeModel,
			}

			instanceGroupManagerAction, response, err := vpcService.CreateInstanceGroupManagerAction(createInstanceGroupManagerActionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(instanceGroupManagerAction).ToNot(BeNil())

		})
	})

	Describe(`GetInstanceGroupManagerAction - Retrieve specified instance group manager action`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInstanceGroupManagerAction(getInstanceGroupManagerActionOptions *GetInstanceGroupManagerActionOptions)`, func() {

			getInstanceGroupManagerActionOptions := &vpcv1.GetInstanceGroupManagerActionOptions{
				InstanceGroupID:        core.StringPtr("testString"),
				InstanceGroupManagerID: core.StringPtr("testString"),
				ID:                     core.StringPtr("testString"),
			}

			instanceGroupManagerAction, response, err := vpcService.GetInstanceGroupManagerAction(getInstanceGroupManagerActionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManagerAction).ToNot(BeNil())

		})
	})

	Describe(`UpdateInstanceGroupManagerAction - Update specified instance group manager action`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateInstanceGroupManagerAction(updateInstanceGroupManagerActionOptions *UpdateInstanceGroupManagerActionOptions)`, func() {

			instanceGroupManagerScheduledActionGroupPatchModel := &vpcv1.InstanceGroupManagerScheduledActionGroupPatch{
				MembershipCount: core.Int64Ptr(int64(10)),
			}

			instanceGroupManagerScheduledActionByManagerPatchManagerModel := &vpcv1.InstanceGroupManagerScheduledActionByManagerPatchManagerInstanceGroupManagerScheduledActionManagerAutoScalePatch{
				MaxMembershipCount: core.Int64Ptr(int64(10)),
				MinMembershipCount: core.Int64Ptr(int64(10)),
			}

			instanceGroupManagerActionPatchModel := &vpcv1.InstanceGroupManagerActionPatchInstanceGroupManagerScheduledActionPatch{
				Name:     core.StringPtr("my-instance-group-manager-action"),
				CronSpec: core.StringPtr("*/5 1,2,3 * * *"),
				Group:    instanceGroupManagerScheduledActionGroupPatchModel,
				Manager:  instanceGroupManagerScheduledActionByManagerPatchManagerModel,
				RunAt:    CreateMockDateTime(),
			}
			instanceGroupManagerActionPatchModelAsPatch, asPatchErr := instanceGroupManagerActionPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateInstanceGroupManagerActionOptions := &vpcv1.UpdateInstanceGroupManagerActionOptions{
				InstanceGroupID:                 core.StringPtr("testString"),
				InstanceGroupManagerID:          core.StringPtr("testString"),
				ID:                              core.StringPtr("testString"),
				InstanceGroupManagerActionPatch: instanceGroupManagerActionPatchModelAsPatch,
			}

			instanceGroupManagerAction, response, err := vpcService.UpdateInstanceGroupManagerAction(updateInstanceGroupManagerActionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManagerAction).ToNot(BeNil())

		})
	})

	Describe(`ListInstanceGroupManagerPolicies - List all policies for an instance group manager`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListInstanceGroupManagerPolicies(listInstanceGroupManagerPoliciesOptions *ListInstanceGroupManagerPoliciesOptions)`, func() {

			listInstanceGroupManagerPoliciesOptions := &vpcv1.ListInstanceGroupManagerPoliciesOptions{
				InstanceGroupID:        core.StringPtr("testString"),
				InstanceGroupManagerID: core.StringPtr("testString"),
				Start:                  core.StringPtr("testString"),
				Limit:                  core.Int64Ptr(int64(1)),
			}

			instanceGroupManagerPolicyCollection, response, err := vpcService.ListInstanceGroupManagerPolicies(listInstanceGroupManagerPoliciesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManagerPolicyCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateInstanceGroupManagerPolicy - Create a policy for an instance group manager`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateInstanceGroupManagerPolicy(createInstanceGroupManagerPolicyOptions *CreateInstanceGroupManagerPolicyOptions)`, func() {

			instanceGroupManagerPolicyPrototypeModel := &vpcv1.InstanceGroupManagerPolicyPrototypeInstanceGroupManagerTargetPolicyPrototype{
				Name:        core.StringPtr("my-instance-group-manager-policy"),
				MetricType:  core.StringPtr("cpu"),
				MetricValue: core.Int64Ptr(int64(38)),
				PolicyType:  core.StringPtr("target"),
			}

			createInstanceGroupManagerPolicyOptions := &vpcv1.CreateInstanceGroupManagerPolicyOptions{
				InstanceGroupID:                     core.StringPtr("testString"),
				InstanceGroupManagerID:              core.StringPtr("testString"),
				InstanceGroupManagerPolicyPrototype: instanceGroupManagerPolicyPrototypeModel,
			}

			instanceGroupManagerPolicy, response, err := vpcService.CreateInstanceGroupManagerPolicy(createInstanceGroupManagerPolicyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(instanceGroupManagerPolicy).ToNot(BeNil())

		})
	})

	Describe(`GetInstanceGroupManagerPolicy - Retrieve an instance group manager policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInstanceGroupManagerPolicy(getInstanceGroupManagerPolicyOptions *GetInstanceGroupManagerPolicyOptions)`, func() {

			getInstanceGroupManagerPolicyOptions := &vpcv1.GetInstanceGroupManagerPolicyOptions{
				InstanceGroupID:        core.StringPtr("testString"),
				InstanceGroupManagerID: core.StringPtr("testString"),
				ID:                     core.StringPtr("testString"),
			}

			instanceGroupManagerPolicy, response, err := vpcService.GetInstanceGroupManagerPolicy(getInstanceGroupManagerPolicyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManagerPolicy).ToNot(BeNil())

		})
	})

	Describe(`UpdateInstanceGroupManagerPolicy - Update an instance group manager policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateInstanceGroupManagerPolicy(updateInstanceGroupManagerPolicyOptions *UpdateInstanceGroupManagerPolicyOptions)`, func() {

			instanceGroupManagerPolicyPatchModel := &vpcv1.InstanceGroupManagerPolicyPatch{
				MetricType:  core.StringPtr("cpu"),
				MetricValue: core.Int64Ptr(int64(38)),
				Name:        core.StringPtr("my-instance-group-manager-policy"),
			}
			instanceGroupManagerPolicyPatchModelAsPatch, asPatchErr := instanceGroupManagerPolicyPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateInstanceGroupManagerPolicyOptions := &vpcv1.UpdateInstanceGroupManagerPolicyOptions{
				InstanceGroupID:                 core.StringPtr("testString"),
				InstanceGroupManagerID:          core.StringPtr("testString"),
				ID:                              core.StringPtr("testString"),
				InstanceGroupManagerPolicyPatch: instanceGroupManagerPolicyPatchModelAsPatch,
			}

			instanceGroupManagerPolicy, response, err := vpcService.UpdateInstanceGroupManagerPolicy(updateInstanceGroupManagerPolicyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManagerPolicy).ToNot(BeNil())

		})
	})

	Describe(`ListInstanceGroupMemberships - List all memberships for an instance group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListInstanceGroupMemberships(listInstanceGroupMembershipsOptions *ListInstanceGroupMembershipsOptions)`, func() {

			listInstanceGroupMembershipsOptions := &vpcv1.ListInstanceGroupMembershipsOptions{
				InstanceGroupID: core.StringPtr("testString"),
				Start:           core.StringPtr("testString"),
				Limit:           core.Int64Ptr(int64(1)),
			}

			instanceGroupMembershipCollection, response, err := vpcService.ListInstanceGroupMemberships(listInstanceGroupMembershipsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupMembershipCollection).ToNot(BeNil())

		})
	})

	Describe(`GetInstanceGroupMembership - Retrieve an instance group membership`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetInstanceGroupMembership(getInstanceGroupMembershipOptions *GetInstanceGroupMembershipOptions)`, func() {

			getInstanceGroupMembershipOptions := &vpcv1.GetInstanceGroupMembershipOptions{
				InstanceGroupID: core.StringPtr("testString"),
				ID:              core.StringPtr("testString"),
			}

			instanceGroupMembership, response, err := vpcService.GetInstanceGroupMembership(getInstanceGroupMembershipOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupMembership).ToNot(BeNil())

		})
	})

	Describe(`UpdateInstanceGroupMembership - Update an instance group membership`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateInstanceGroupMembership(updateInstanceGroupMembershipOptions *UpdateInstanceGroupMembershipOptions)`, func() {

			instanceGroupMembershipPatchModel := &vpcv1.InstanceGroupMembershipPatch{
				Name: core.StringPtr("my-instance-group-membership"),
			}
			instanceGroupMembershipPatchModelAsPatch, asPatchErr := instanceGroupMembershipPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateInstanceGroupMembershipOptions := &vpcv1.UpdateInstanceGroupMembershipOptions{
				InstanceGroupID:              core.StringPtr("testString"),
				ID:                           core.StringPtr("testString"),
				InstanceGroupMembershipPatch: instanceGroupMembershipPatchModelAsPatch,
			}

			instanceGroupMembership, response, err := vpcService.UpdateInstanceGroupMembership(updateInstanceGroupMembershipOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupMembership).ToNot(BeNil())

		})
	})

	Describe(`ListDedicatedHostGroups - List all dedicated host groups`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDedicatedHostGroups(listDedicatedHostGroupsOptions *ListDedicatedHostGroupsOptions)`, func() {

			listDedicatedHostGroupsOptions := &vpcv1.ListDedicatedHostGroupsOptions{
				Start:           core.StringPtr("testString"),
				Limit:           core.Int64Ptr(int64(1)),
				ResourceGroupID: core.StringPtr("testString"),
				ZoneName:        core.StringPtr("testString"),
			}

			dedicatedHostGroupCollection, response, err := vpcService.ListDedicatedHostGroups(listDedicatedHostGroupsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHostGroupCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateDedicatedHostGroup - Create a dedicated host group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDedicatedHostGroup(createDedicatedHostGroupOptions *CreateDedicatedHostGroupOptions)`, func() {

			resourceGroupIdentityModel := &vpcv1.ResourceGroupIdentityByID{
				ID: core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345"),
			}

			zoneIdentityModel := &vpcv1.ZoneIdentityByName{
				Name: core.StringPtr("us-south-1"),
			}

			createDedicatedHostGroupOptions := &vpcv1.CreateDedicatedHostGroupOptions{
				Class:         core.StringPtr("mx2"),
				Family:        core.StringPtr("balanced"),
				Name:          core.StringPtr("testString"),
				ResourceGroup: resourceGroupIdentityModel,
				Zone:          zoneIdentityModel,
			}

			dedicatedHostGroup, response, err := vpcService.CreateDedicatedHostGroup(createDedicatedHostGroupOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(dedicatedHostGroup).ToNot(BeNil())

		})
	})

	Describe(`GetDedicatedHostGroup - Retrieve a dedicated host group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDedicatedHostGroup(getDedicatedHostGroupOptions *GetDedicatedHostGroupOptions)`, func() {

			getDedicatedHostGroupOptions := &vpcv1.GetDedicatedHostGroupOptions{
				ID: core.StringPtr("testString"),
			}

			dedicatedHostGroup, response, err := vpcService.GetDedicatedHostGroup(getDedicatedHostGroupOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHostGroup).ToNot(BeNil())

		})
	})

	Describe(`UpdateDedicatedHostGroup - Update a dedicated host group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateDedicatedHostGroup(updateDedicatedHostGroupOptions *UpdateDedicatedHostGroupOptions)`, func() {

			dedicatedHostGroupPatchModel := &vpcv1.DedicatedHostGroupPatch{
				Name: core.StringPtr("my-host-group-modified"),
			}
			dedicatedHostGroupPatchModelAsPatch, asPatchErr := dedicatedHostGroupPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateDedicatedHostGroupOptions := &vpcv1.UpdateDedicatedHostGroupOptions{
				ID:                      core.StringPtr("testString"),
				DedicatedHostGroupPatch: dedicatedHostGroupPatchModelAsPatch,
			}

			dedicatedHostGroup, response, err := vpcService.UpdateDedicatedHostGroup(updateDedicatedHostGroupOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHostGroup).ToNot(BeNil())

		})
	})

	Describe(`ListDedicatedHostProfiles - List all dedicated host profiles`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDedicatedHostProfiles(listDedicatedHostProfilesOptions *ListDedicatedHostProfilesOptions)`, func() {

			listDedicatedHostProfilesOptions := &vpcv1.ListDedicatedHostProfilesOptions{
				Start: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(1)),
			}

			dedicatedHostProfileCollection, response, err := vpcService.ListDedicatedHostProfiles(listDedicatedHostProfilesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHostProfileCollection).ToNot(BeNil())

		})
	})

	Describe(`GetDedicatedHostProfile - Retrieve a dedicated host profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDedicatedHostProfile(getDedicatedHostProfileOptions *GetDedicatedHostProfileOptions)`, func() {

			getDedicatedHostProfileOptions := &vpcv1.GetDedicatedHostProfileOptions{
				Name: core.StringPtr("testString"),
			}

			dedicatedHostProfile, response, err := vpcService.GetDedicatedHostProfile(getDedicatedHostProfileOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHostProfile).ToNot(BeNil())

		})
	})

	Describe(`ListDedicatedHosts - List all dedicated hosts`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDedicatedHosts(listDedicatedHostsOptions *ListDedicatedHostsOptions)`, func() {

			listDedicatedHostsOptions := &vpcv1.ListDedicatedHostsOptions{
				DedicatedHostGroupID: core.StringPtr("testString"),
				Start:                core.StringPtr("testString"),
				Limit:                core.Int64Ptr(int64(1)),
				ResourceGroupID:      core.StringPtr("testString"),
				ZoneName:             core.StringPtr("testString"),
			}

			dedicatedHostCollection, response, err := vpcService.ListDedicatedHosts(listDedicatedHostsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHostCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateDedicatedHost - Create a dedicated host`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDedicatedHost(createDedicatedHostOptions *CreateDedicatedHostOptions)`, func() {

			dedicatedHostProfileIdentityModel := &vpcv1.DedicatedHostProfileIdentityByName{
				Name: core.StringPtr("m-62x496"),
			}

			resourceGroupIdentityModel := &vpcv1.ResourceGroupIdentityByID{
				ID: core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345"),
			}

			dedicatedHostGroupIdentityModel := &vpcv1.DedicatedHostGroupIdentityByID{
				ID: core.StringPtr("0c8eccb4-271c-4518-956c-32bfce5cf83b"),
			}

			dedicatedHostPrototypeModel := &vpcv1.DedicatedHostPrototypeDedicatedHostByGroup{
				InstancePlacementEnabled: core.BoolPtr(true),
				Name:                     core.StringPtr("my-host"),
				Profile:                  dedicatedHostProfileIdentityModel,
				ResourceGroup:            resourceGroupIdentityModel,
				Group:                    dedicatedHostGroupIdentityModel,
			}

			createDedicatedHostOptions := &vpcv1.CreateDedicatedHostOptions{
				DedicatedHostPrototype: dedicatedHostPrototypeModel,
			}

			dedicatedHost, response, err := vpcService.CreateDedicatedHost(createDedicatedHostOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(dedicatedHost).ToNot(BeNil())

		})
	})

	Describe(`ListDedicatedHostDisks - List all disks on a dedicated host`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDedicatedHostDisks(listDedicatedHostDisksOptions *ListDedicatedHostDisksOptions)`, func() {

			listDedicatedHostDisksOptions := &vpcv1.ListDedicatedHostDisksOptions{
				DedicatedHostID: core.StringPtr("testString"),
			}

			dedicatedHostDiskCollection, response, err := vpcService.ListDedicatedHostDisks(listDedicatedHostDisksOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHostDiskCollection).ToNot(BeNil())

		})
	})

	Describe(`GetDedicatedHostDisk - Retrieve a dedicated host disk`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDedicatedHostDisk(getDedicatedHostDiskOptions *GetDedicatedHostDiskOptions)`, func() {

			getDedicatedHostDiskOptions := &vpcv1.GetDedicatedHostDiskOptions{
				DedicatedHostID: core.StringPtr("testString"),
				ID:              core.StringPtr("testString"),
			}

			dedicatedHostDisk, response, err := vpcService.GetDedicatedHostDisk(getDedicatedHostDiskOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHostDisk).ToNot(BeNil())

		})
	})

	Describe(`UpdateDedicatedHostDisk - Update a dedicated host disk`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateDedicatedHostDisk(updateDedicatedHostDiskOptions *UpdateDedicatedHostDiskOptions)`, func() {

			dedicatedHostDiskPatchModel := &vpcv1.DedicatedHostDiskPatch{
				Name: core.StringPtr("my-disk-updated"),
			}
			dedicatedHostDiskPatchModelAsPatch, asPatchErr := dedicatedHostDiskPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateDedicatedHostDiskOptions := &vpcv1.UpdateDedicatedHostDiskOptions{
				DedicatedHostID:        core.StringPtr("testString"),
				ID:                     core.StringPtr("testString"),
				DedicatedHostDiskPatch: dedicatedHostDiskPatchModelAsPatch,
			}

			dedicatedHostDisk, response, err := vpcService.UpdateDedicatedHostDisk(updateDedicatedHostDiskOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHostDisk).ToNot(BeNil())

		})
	})

	Describe(`GetDedicatedHost - Retrieve a dedicated host`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDedicatedHost(getDedicatedHostOptions *GetDedicatedHostOptions)`, func() {

			getDedicatedHostOptions := &vpcv1.GetDedicatedHostOptions{
				ID: core.StringPtr("testString"),
			}

			dedicatedHost, response, err := vpcService.GetDedicatedHost(getDedicatedHostOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHost).ToNot(BeNil())

		})
	})

	Describe(`UpdateDedicatedHost - Update a dedicated host`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateDedicatedHost(updateDedicatedHostOptions *UpdateDedicatedHostOptions)`, func() {

			dedicatedHostPatchModel := &vpcv1.DedicatedHostPatch{
				InstancePlacementEnabled: core.BoolPtr(true),
				Name:                     core.StringPtr("my-host"),
			}
			dedicatedHostPatchModelAsPatch, asPatchErr := dedicatedHostPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateDedicatedHostOptions := &vpcv1.UpdateDedicatedHostOptions{
				ID:                 core.StringPtr("testString"),
				DedicatedHostPatch: dedicatedHostPatchModelAsPatch,
			}

			dedicatedHost, response, err := vpcService.UpdateDedicatedHost(updateDedicatedHostOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHost).ToNot(BeNil())

		})
	})

	Describe(`ListPlacementGroups - List all placement groups`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListPlacementGroups(listPlacementGroupsOptions *ListPlacementGroupsOptions)`, func() {

			listPlacementGroupsOptions := &vpcv1.ListPlacementGroupsOptions{
				Start: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(1)),
			}

			placementGroupCollection, response, err := vpcService.ListPlacementGroups(listPlacementGroupsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(placementGroupCollection).ToNot(BeNil())

		})
	})

	Describe(`CreatePlacementGroup - Create a placement group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreatePlacementGroup(createPlacementGroupOptions *CreatePlacementGroupOptions)`, func() {

			resourceGroupIdentityModel := &vpcv1.ResourceGroupIdentityByID{
				ID: core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345"),
			}

			createPlacementGroupOptions := &vpcv1.CreatePlacementGroupOptions{
				Strategy:      core.StringPtr("host_spread"),
				Name:          core.StringPtr("my-placement-group"),
				ResourceGroup: resourceGroupIdentityModel,
			}

			placementGroup, response, err := vpcService.CreatePlacementGroup(createPlacementGroupOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(placementGroup).ToNot(BeNil())

		})
	})

	Describe(`GetPlacementGroup - Retrieve a placement group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPlacementGroup(getPlacementGroupOptions *GetPlacementGroupOptions)`, func() {

			getPlacementGroupOptions := &vpcv1.GetPlacementGroupOptions{
				ID: core.StringPtr("testString"),
			}

			placementGroup, response, err := vpcService.GetPlacementGroup(getPlacementGroupOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(placementGroup).ToNot(BeNil())

		})
	})

	Describe(`UpdatePlacementGroup - Update a placement group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdatePlacementGroup(updatePlacementGroupOptions *UpdatePlacementGroupOptions)`, func() {

			placementGroupPatchModel := &vpcv1.PlacementGroupPatch{
				Name: core.StringPtr("my-placement-group"),
			}
			placementGroupPatchModelAsPatch, asPatchErr := placementGroupPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updatePlacementGroupOptions := &vpcv1.UpdatePlacementGroupOptions{
				ID:                  core.StringPtr("testString"),
				PlacementGroupPatch: placementGroupPatchModelAsPatch,
			}

			placementGroup, response, err := vpcService.UpdatePlacementGroup(updatePlacementGroupOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(placementGroup).ToNot(BeNil())

		})
	})

	Describe(`ListBareMetalServerProfiles - List all bare metal server profiles`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListBareMetalServerProfiles(listBareMetalServerProfilesOptions *ListBareMetalServerProfilesOptions)`, func() {

			listBareMetalServerProfilesOptions := &vpcv1.ListBareMetalServerProfilesOptions{
				Start: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(1)),
			}

			bareMetalServerProfileCollection, response, err := vpcService.ListBareMetalServerProfiles(listBareMetalServerProfilesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServerProfileCollection).ToNot(BeNil())

		})
	})

	Describe(`GetBareMetalServerProfile - Retrieve a bare metal server profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetBareMetalServerProfile(getBareMetalServerProfileOptions *GetBareMetalServerProfileOptions)`, func() {

			getBareMetalServerProfileOptions := &vpcv1.GetBareMetalServerProfileOptions{
				Name: core.StringPtr("testString"),
			}

			bareMetalServerProfile, response, err := vpcService.GetBareMetalServerProfile(getBareMetalServerProfileOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServerProfile).ToNot(BeNil())

		})
	})

	Describe(`ListBareMetalServers - List all bare metal servers`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListBareMetalServers(listBareMetalServersOptions *ListBareMetalServersOptions)`, func() {

			listBareMetalServersOptions := &vpcv1.ListBareMetalServersOptions{
				Start:                       core.StringPtr("testString"),
				Limit:                       core.Int64Ptr(int64(1)),
				ResourceGroupID:             core.StringPtr("testString"),
				Name:                        core.StringPtr("testString"),
				VPCID:                       core.StringPtr("testString"),
				VPCCRN:                      core.StringPtr("testString"),
				VPCName:                     core.StringPtr("testString"),
				NetworkInterfacesSubnetID:   core.StringPtr("testString"),
				NetworkInterfacesSubnetCRN:  core.StringPtr("testString"),
				NetworkInterfacesSubnetName: core.StringPtr("testString"),
				Sort:                        core.StringPtr("name"),
			}

			bareMetalServerCollection, response, err := vpcService.ListBareMetalServers(listBareMetalServersOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServerCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateBareMetalServer - Create a bare metal server`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateBareMetalServer(createBareMetalServerOptions *CreateBareMetalServerOptions)`, func() {

			imageIdentityModel := &vpcv1.ImageIdentityByID{
				ID: core.StringPtr("72b27b5c-f4b0-48bb-b954-5becc7c1dcb8"),
			}

			keyIdentityModel := &vpcv1.KeyIdentityByID{
				ID: core.StringPtr("a6b1a881-2ce8-41a3-80fc-36316a73f803"),
			}

			bareMetalServerInitializationPrototypeModel := &vpcv1.BareMetalServerInitializationPrototype{
				Image:    imageIdentityModel,
				Keys:     []vpcv1.KeyIdentityIntf{keyIdentityModel},
				UserData: core.StringPtr("testString"),
			}

			networkInterfaceIPPrototypeModel := &vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext{
				Address:    core.StringPtr("10.0.0.5"),
				AutoDelete: core.BoolPtr(false),
				Name:       core.StringPtr("my-reserved-ip"),
			}

			securityGroupIdentityModel := &vpcv1.SecurityGroupIdentityByID{
				ID: core.StringPtr("be5df5ca-12a0-494b-907e-aa6ec2bfa271"),
			}

			subnetIdentityModel := &vpcv1.SubnetIdentityByID{
				ID: core.StringPtr("7ec86020-1c6e-4889-b3f0-a15f2e50f87e"),
			}

			bareMetalServerPrimaryNetworkInterfacePrototypeModel := &vpcv1.BareMetalServerPrimaryNetworkInterfacePrototype{
				AllowIPSpoofing:         core.BoolPtr(true),
				AllowedVlans:            []int64{int64(4)},
				EnableInfrastructureNat: core.BoolPtr(true),
				InterfaceType:           core.StringPtr("pci"),
				Ips:                     []vpcv1.NetworkInterfaceIPPrototypeIntf{networkInterfaceIPPrototypeModel},
				Name:                    core.StringPtr("my-network-interface"),
				PrimaryIP:               networkInterfaceIPPrototypeModel,
				SecurityGroups:          []vpcv1.SecurityGroupIdentityIntf{securityGroupIdentityModel},
				Subnet:                  subnetIdentityModel,
			}

			bareMetalServerProfileIdentityModel := &vpcv1.BareMetalServerProfileIdentityByName{
				Name: core.StringPtr("bm2-80x1356"),
			}

			zoneIdentityModel := &vpcv1.ZoneIdentityByName{
				Name: core.StringPtr("us-south-1"),
			}

			bareMetalServerNetworkInterfacePrototypeModel := &vpcv1.BareMetalServerNetworkInterfacePrototypeBareMetalServerNetworkInterfaceByPciPrototype{
				AllowIPSpoofing:         core.BoolPtr(true),
				EnableInfrastructureNat: core.BoolPtr(true),
				InterfaceType:           core.StringPtr("pci"),
				Ips:                     []vpcv1.NetworkInterfaceIPPrototypeIntf{networkInterfaceIPPrototypeModel},
				Name:                    core.StringPtr("my-network-interface"),
				PrimaryIP:               networkInterfaceIPPrototypeModel,
				SecurityGroups:          []vpcv1.SecurityGroupIdentityIntf{securityGroupIdentityModel},
				Subnet:                  subnetIdentityModel,
				AllowedVlans:            []int64{int64(4)},
			}

			resourceGroupIdentityModel := &vpcv1.ResourceGroupIdentityByID{
				ID: core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345"),
			}

			bareMetalServerTrustedPlatformModulePrototypeModel := &vpcv1.BareMetalServerTrustedPlatformModulePrototype{
				Enabled: core.BoolPtr(true),
				Mode:    core.StringPtr("tpm_2"),
			}

			vpcIdentityModel := &vpcv1.VPCIdentityByID{
				ID: core.StringPtr("4727d842-f94f-4a2d-824a-9bc9b02c523b"),
			}

			createBareMetalServerOptions := &vpcv1.CreateBareMetalServerOptions{
				Initialization:          bareMetalServerInitializationPrototypeModel,
				PrimaryNetworkInterface: bareMetalServerPrimaryNetworkInterfacePrototypeModel,
				Profile:                 bareMetalServerProfileIdentityModel,
				Zone:                    zoneIdentityModel,
				EnableSecureBoot:        core.BoolPtr(false),
				Name:                    core.StringPtr("my-server"),
				NetworkInterfaces:       []vpcv1.BareMetalServerNetworkInterfacePrototypeIntf{bareMetalServerNetworkInterfacePrototypeModel},
				ResourceGroup:           resourceGroupIdentityModel,
				TrustedPlatformModule:   bareMetalServerTrustedPlatformModulePrototypeModel,
				VPC:                     vpcIdentityModel,
			}

			bareMetalServer, response, err := vpcService.CreateBareMetalServer(createBareMetalServerOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(bareMetalServer).ToNot(BeNil())

		})
	})

	Describe(`GetBareMetalServerConsole - Retrieve the console WebSocket for a bare metal server`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetBareMetalServerConsole(getBareMetalServerConsoleOptions *GetBareMetalServerConsoleOptions)`, func() {

			getBareMetalServerConsoleOptions := &vpcv1.GetBareMetalServerConsoleOptions{
				BareMetalServerID: core.StringPtr("testString"),
				AccessToken:       core.StringPtr("VGhpcyBJcyBhIHRva2Vu"),
			}

			response, err := vpcService.GetBareMetalServerConsole(getBareMetalServerConsoleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))

		})
	})

	Describe(`CreateBareMetalServerConsoleAccessToken - Create a console access token for a bare metal server`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateBareMetalServerConsoleAccessToken(createBareMetalServerConsoleAccessTokenOptions *CreateBareMetalServerConsoleAccessTokenOptions)`, func() {

			createBareMetalServerConsoleAccessTokenOptions := &vpcv1.CreateBareMetalServerConsoleAccessTokenOptions{
				BareMetalServerID: core.StringPtr("testString"),
				ConsoleType:       core.StringPtr("serial"),
				Force:             core.BoolPtr(false),
			}

			bareMetalServerConsoleAccessToken, response, err := vpcService.CreateBareMetalServerConsoleAccessToken(createBareMetalServerConsoleAccessTokenOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServerConsoleAccessToken).ToNot(BeNil())

		})
	})

	Describe(`ListBareMetalServerDisks - List all disks on a bare metal server`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListBareMetalServerDisks(listBareMetalServerDisksOptions *ListBareMetalServerDisksOptions)`, func() {

			listBareMetalServerDisksOptions := &vpcv1.ListBareMetalServerDisksOptions{
				BareMetalServerID: core.StringPtr("testString"),
			}

			bareMetalServerDiskCollection, response, err := vpcService.ListBareMetalServerDisks(listBareMetalServerDisksOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServerDiskCollection).ToNot(BeNil())

		})
	})

	Describe(`GetBareMetalServerDisk - Retrieve a bare metal server disk`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetBareMetalServerDisk(getBareMetalServerDiskOptions *GetBareMetalServerDiskOptions)`, func() {

			getBareMetalServerDiskOptions := &vpcv1.GetBareMetalServerDiskOptions{
				BareMetalServerID: core.StringPtr("testString"),
				ID:                core.StringPtr("testString"),
			}

			bareMetalServerDisk, response, err := vpcService.GetBareMetalServerDisk(getBareMetalServerDiskOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServerDisk).ToNot(BeNil())

		})
	})

	Describe(`UpdateBareMetalServerDisk - Update a bare metal server disk`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateBareMetalServerDisk(updateBareMetalServerDiskOptions *UpdateBareMetalServerDiskOptions)`, func() {

			bareMetalServerDiskPatchModel := &vpcv1.BareMetalServerDiskPatch{
				Name: core.StringPtr("my-bare-metal-server-disk-updated"),
			}
			bareMetalServerDiskPatchModelAsPatch, asPatchErr := bareMetalServerDiskPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateBareMetalServerDiskOptions := &vpcv1.UpdateBareMetalServerDiskOptions{
				BareMetalServerID:        core.StringPtr("testString"),
				ID:                       core.StringPtr("testString"),
				BareMetalServerDiskPatch: bareMetalServerDiskPatchModelAsPatch,
			}

			bareMetalServerDisk, response, err := vpcService.UpdateBareMetalServerDisk(updateBareMetalServerDiskOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServerDisk).ToNot(BeNil())

		})
	})

	Describe(`ListBareMetalServerNetworkInterfaces - List all network interfaces on a bare metal server`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListBareMetalServerNetworkInterfaces(listBareMetalServerNetworkInterfacesOptions *ListBareMetalServerNetworkInterfacesOptions)`, func() {

			listBareMetalServerNetworkInterfacesOptions := &vpcv1.ListBareMetalServerNetworkInterfacesOptions{
				BareMetalServerID: core.StringPtr("testString"),
				Start:             core.StringPtr("testString"),
				Limit:             core.Int64Ptr(int64(1)),
			}

			bareMetalServerNetworkInterfaceCollection, response, err := vpcService.ListBareMetalServerNetworkInterfaces(listBareMetalServerNetworkInterfacesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServerNetworkInterfaceCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateBareMetalServerNetworkInterface - Create a network interface on a bare metal server`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateBareMetalServerNetworkInterface(createBareMetalServerNetworkInterfaceOptions *CreateBareMetalServerNetworkInterfaceOptions)`, func() {

			networkInterfaceIPPrototypeModel := &vpcv1.NetworkInterfaceIPPrototypeReservedIPPrototypeNetworkInterfaceContext{
				Address:    core.StringPtr("10.0.0.5"),
				AutoDelete: core.BoolPtr(false),
				Name:       core.StringPtr("my-reserved-ip"),
			}

			securityGroupIdentityModel := &vpcv1.SecurityGroupIdentityByID{
				ID: core.StringPtr("be5df5ca-12a0-494b-907e-aa6ec2bfa271"),
			}

			subnetIdentityModel := &vpcv1.SubnetIdentityByID{
				ID: core.StringPtr("7ec86020-1c6e-4889-b3f0-a15f2e50f87e"),
			}

			bareMetalServerNetworkInterfacePrototypeModel := &vpcv1.BareMetalServerNetworkInterfacePrototypeBareMetalServerNetworkInterfaceByPciPrototype{
				AllowIPSpoofing:         core.BoolPtr(true),
				EnableInfrastructureNat: core.BoolPtr(true),
				InterfaceType:           core.StringPtr("pci"),
				Ips:                     []vpcv1.NetworkInterfaceIPPrototypeIntf{networkInterfaceIPPrototypeModel},
				Name:                    core.StringPtr("my-network-interface"),
				PrimaryIP:               networkInterfaceIPPrototypeModel,
				SecurityGroups:          []vpcv1.SecurityGroupIdentityIntf{securityGroupIdentityModel},
				Subnet:                  subnetIdentityModel,
				AllowedVlans:            []int64{int64(4)},
			}

			createBareMetalServerNetworkInterfaceOptions := &vpcv1.CreateBareMetalServerNetworkInterfaceOptions{
				BareMetalServerID:                        core.StringPtr("testString"),
				BareMetalServerNetworkInterfacePrototype: bareMetalServerNetworkInterfacePrototypeModel,
			}

			bareMetalServerNetworkInterface, response, err := vpcService.CreateBareMetalServerNetworkInterface(createBareMetalServerNetworkInterfaceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(bareMetalServerNetworkInterface).ToNot(BeNil())

		})
	})

	Describe(`GetBareMetalServerNetworkInterface - Retrieve a network interface`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetBareMetalServerNetworkInterface(getBareMetalServerNetworkInterfaceOptions *GetBareMetalServerNetworkInterfaceOptions)`, func() {

			getBareMetalServerNetworkInterfaceOptions := &vpcv1.GetBareMetalServerNetworkInterfaceOptions{
				BareMetalServerID: core.StringPtr("testString"),
				ID:                core.StringPtr("testString"),
			}

			bareMetalServerNetworkInterface, response, err := vpcService.GetBareMetalServerNetworkInterface(getBareMetalServerNetworkInterfaceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServerNetworkInterface).ToNot(BeNil())

		})
	})

	Describe(`UpdateBareMetalServerNetworkInterface - Update a network interface`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateBareMetalServerNetworkInterface(updateBareMetalServerNetworkInterfaceOptions *UpdateBareMetalServerNetworkInterfaceOptions)`, func() {

			bareMetalServerNetworkInterfacePatchModel := &vpcv1.BareMetalServerNetworkInterfacePatch{
				AllowIPSpoofing:         core.BoolPtr(true),
				AllowedVlans:            []int64{int64(4)},
				EnableInfrastructureNat: core.BoolPtr(true),
				Name:                    core.StringPtr("my-network-interface"),
			}
			bareMetalServerNetworkInterfacePatchModelAsPatch, asPatchErr := bareMetalServerNetworkInterfacePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateBareMetalServerNetworkInterfaceOptions := &vpcv1.UpdateBareMetalServerNetworkInterfaceOptions{
				BareMetalServerID:                    core.StringPtr("testString"),
				ID:                                   core.StringPtr("testString"),
				BareMetalServerNetworkInterfacePatch: bareMetalServerNetworkInterfacePatchModelAsPatch,
			}

			bareMetalServerNetworkInterface, response, err := vpcService.UpdateBareMetalServerNetworkInterface(updateBareMetalServerNetworkInterfaceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServerNetworkInterface).ToNot(BeNil())

		})
	})

	Describe(`ListBareMetalServerNetworkInterfaceFloatingIps - List all floating IPs associated with a network interface`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListBareMetalServerNetworkInterfaceFloatingIps(listBareMetalServerNetworkInterfaceFloatingIpsOptions *ListBareMetalServerNetworkInterfaceFloatingIpsOptions)`, func() {

			listBareMetalServerNetworkInterfaceFloatingIpsOptions := &vpcv1.ListBareMetalServerNetworkInterfaceFloatingIpsOptions{
				BareMetalServerID:  core.StringPtr("testString"),
				NetworkInterfaceID: core.StringPtr("testString"),
			}

			floatingIPUnpaginatedCollection, response, err := vpcService.ListBareMetalServerNetworkInterfaceFloatingIps(listBareMetalServerNetworkInterfaceFloatingIpsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(floatingIPUnpaginatedCollection).ToNot(BeNil())

		})
	})

	Describe(`GetBareMetalServerNetworkInterfaceFloatingIP - Retrieve associated floating IP`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetBareMetalServerNetworkInterfaceFloatingIP(getBareMetalServerNetworkInterfaceFloatingIPOptions *GetBareMetalServerNetworkInterfaceFloatingIPOptions)`, func() {

			getBareMetalServerNetworkInterfaceFloatingIPOptions := &vpcv1.GetBareMetalServerNetworkInterfaceFloatingIPOptions{
				BareMetalServerID:  core.StringPtr("testString"),
				NetworkInterfaceID: core.StringPtr("testString"),
				ID:                 core.StringPtr("testString"),
			}

			floatingIP, response, err := vpcService.GetBareMetalServerNetworkInterfaceFloatingIP(getBareMetalServerNetworkInterfaceFloatingIPOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(floatingIP).ToNot(BeNil())

		})
	})

	Describe(`AddBareMetalServerNetworkInterfaceFloatingIP - Associate a floating IP with a network interface`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AddBareMetalServerNetworkInterfaceFloatingIP(addBareMetalServerNetworkInterfaceFloatingIPOptions *AddBareMetalServerNetworkInterfaceFloatingIPOptions)`, func() {

			addBareMetalServerNetworkInterfaceFloatingIPOptions := &vpcv1.AddBareMetalServerNetworkInterfaceFloatingIPOptions{
				BareMetalServerID:  core.StringPtr("testString"),
				NetworkInterfaceID: core.StringPtr("testString"),
				ID:                 core.StringPtr("testString"),
			}

			floatingIP, response, err := vpcService.AddBareMetalServerNetworkInterfaceFloatingIP(addBareMetalServerNetworkInterfaceFloatingIPOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(floatingIP).ToNot(BeNil())

		})
	})

	Describe(`GetBareMetalServer - Retrieve a bare metal server`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetBareMetalServer(getBareMetalServerOptions *GetBareMetalServerOptions)`, func() {

			getBareMetalServerOptions := &vpcv1.GetBareMetalServerOptions{
				ID: core.StringPtr("testString"),
			}

			bareMetalServer, response, err := vpcService.GetBareMetalServer(getBareMetalServerOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServer).ToNot(BeNil())

		})
	})

	Describe(`UpdateBareMetalServer - Update a bare metal server`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateBareMetalServer(updateBareMetalServerOptions *UpdateBareMetalServerOptions)`, func() {

			bareMetalServerTrustedPlatformModulePatchModel := &vpcv1.BareMetalServerTrustedPlatformModulePatch{
				Enabled: core.BoolPtr(true),
				Mode:    core.StringPtr("tpm_2"),
			}

			bareMetalServerPatchModel := &vpcv1.BareMetalServerPatch{
				EnableSecureBoot:      core.BoolPtr(false),
				Name:                  core.StringPtr("my-server"),
				TrustedPlatformModule: bareMetalServerTrustedPlatformModulePatchModel,
			}
			bareMetalServerPatchModelAsPatch, asPatchErr := bareMetalServerPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateBareMetalServerOptions := &vpcv1.UpdateBareMetalServerOptions{
				ID:                   core.StringPtr("testString"),
				BareMetalServerPatch: bareMetalServerPatchModelAsPatch,
			}

			bareMetalServer, response, err := vpcService.UpdateBareMetalServer(updateBareMetalServerOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServer).ToNot(BeNil())

		})
	})

	Describe(`GetBareMetalServerInitialization - Retrieve initialization configuration for a bare metal server`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetBareMetalServerInitialization(getBareMetalServerInitializationOptions *GetBareMetalServerInitializationOptions)`, func() {

			getBareMetalServerInitializationOptions := &vpcv1.GetBareMetalServerInitializationOptions{
				ID: core.StringPtr("testString"),
			}

			bareMetalServerInitialization, response, err := vpcService.GetBareMetalServerInitialization(getBareMetalServerInitializationOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServerInitialization).ToNot(BeNil())

		})
	})

	Describe(`CreateBareMetalServerRestart - Restart a bare metal server`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateBareMetalServerRestart(createBareMetalServerRestartOptions *CreateBareMetalServerRestartOptions)`, func() {

			createBareMetalServerRestartOptions := &vpcv1.CreateBareMetalServerRestartOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := vpcService.CreateBareMetalServerRestart(createBareMetalServerRestartOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`CreateBareMetalServerStart - Start a bare metal server`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateBareMetalServerStart(createBareMetalServerStartOptions *CreateBareMetalServerStartOptions)`, func() {

			createBareMetalServerStartOptions := &vpcv1.CreateBareMetalServerStartOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := vpcService.CreateBareMetalServerStart(createBareMetalServerStartOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`CreateBareMetalServerStop - Stop a bare metal server`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateBareMetalServerStop(createBareMetalServerStopOptions *CreateBareMetalServerStopOptions)`, func() {

			createBareMetalServerStopOptions := &vpcv1.CreateBareMetalServerStopOptions{
				ID:   core.StringPtr("testString"),
				Type: core.StringPtr("soft"),
			}

			response, err := vpcService.CreateBareMetalServerStop(createBareMetalServerStopOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`ListVolumeProfiles - List all volume profiles`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListVolumeProfiles(listVolumeProfilesOptions *ListVolumeProfilesOptions)`, func() {

			listVolumeProfilesOptions := &vpcv1.ListVolumeProfilesOptions{
				Start: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(1)),
			}

			volumeProfileCollection, response, err := vpcService.ListVolumeProfiles(listVolumeProfilesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeProfileCollection).ToNot(BeNil())

		})
	})

	Describe(`GetVolumeProfile - Retrieve a volume profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetVolumeProfile(getVolumeProfileOptions *GetVolumeProfileOptions)`, func() {

			getVolumeProfileOptions := &vpcv1.GetVolumeProfileOptions{
				Name: core.StringPtr("testString"),
			}

			volumeProfile, response, err := vpcService.GetVolumeProfile(getVolumeProfileOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeProfile).ToNot(BeNil())

		})
	})

	Describe(`ListVolumes - List all volumes`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListVolumes(listVolumesOptions *ListVolumesOptions)`, func() {

			listVolumesOptions := &vpcv1.ListVolumesOptions{
				Start:    core.StringPtr("testString"),
				Limit:    core.Int64Ptr(int64(1)),
				Name:     core.StringPtr("testString"),
				ZoneName: core.StringPtr("testString"),
			}

			volumeCollection, response, err := vpcService.ListVolumes(listVolumesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateVolume - Create a volume`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateVolume(createVolumeOptions *CreateVolumeOptions)`, func() {

			encryptionKeyIdentityModel := &vpcv1.EncryptionKeyIdentityByCRN{
				CRN: core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/dffc98a0f1f0f95f6613b3b752286b87:e4a29d1a-2ef0-42a6-8fd2-350deb1c647e:key:5437653b-c4b1-447f-9646-b2a2a4cd6179"),
			}

			volumeProfileIdentityModel := &vpcv1.VolumeProfileIdentityByName{
				Name: core.StringPtr("5iops-tier"),
			}

			volumePrototypeResourceGroupModel := &vpcv1.VolumePrototypeResourceGroupResourceGroupIdentityByID{
				ID: core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345"),
			}

			zoneIdentityModel := &vpcv1.ZoneIdentityByName{
				Name: core.StringPtr("us-south-1"),
			}

			volumePrototypeModel := &vpcv1.VolumePrototypeVolumeByCapacity{
				EncryptionKey: encryptionKeyIdentityModel,
				Iops:          core.Int64Ptr(int64(10000)),
				Name:          core.StringPtr("my-volume"),
				Profile:       volumeProfileIdentityModel,
				ResourceGroup: volumePrototypeResourceGroupModel,
				Zone:          zoneIdentityModel,
				Capacity:      core.Int64Ptr(int64(100)),
			}

			createVolumeOptions := &vpcv1.CreateVolumeOptions{
				VolumePrototype: volumePrototypeModel,
			}

			volume, response, err := vpcService.CreateVolume(createVolumeOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(volume).ToNot(BeNil())

		})
	})

	Describe(`GetVolume - Retrieve a volume`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetVolume(getVolumeOptions *GetVolumeOptions)`, func() {

			getVolumeOptions := &vpcv1.GetVolumeOptions{
				ID: core.StringPtr("testString"),
			}

			volume, response, err := vpcService.GetVolume(getVolumeOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volume).ToNot(BeNil())

		})
	})

	Describe(`UpdateVolume - Update a volume`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateVolume(updateVolumeOptions *UpdateVolumeOptions)`, func() {

			volumePatchModel := &vpcv1.VolumePatch{
				Name: core.StringPtr("my-volume"),
			}
			volumePatchModelAsPatch, asPatchErr := volumePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateVolumeOptions := &vpcv1.UpdateVolumeOptions{
				ID:          core.StringPtr("testString"),
				VolumePatch: volumePatchModelAsPatch,
			}

			volume, response, err := vpcService.UpdateVolume(updateVolumeOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volume).ToNot(BeNil())

		})
	})

	Describe(`ListSnapshots - List all snapshots`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSnapshots(listSnapshotsOptions *ListSnapshotsOptions)`, func() {

			listSnapshotsOptions := &vpcv1.ListSnapshotsOptions{
				Start:           core.StringPtr("testString"),
				Limit:           core.Int64Ptr(int64(1)),
				ResourceGroupID: core.StringPtr("testString"),
				Name:            core.StringPtr("testString"),
				SourceVolumeID:  core.StringPtr("testString"),
				SourceVolumeCRN: core.StringPtr("testString"),
				SourceImageID:   core.StringPtr("testString"),
				SourceImageCRN:  core.StringPtr("testString"),
				Sort:            core.StringPtr("name"),
			}

			snapshotCollection, response, err := vpcService.ListSnapshots(listSnapshotsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(snapshotCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateSnapshot - Create a snapshot`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSnapshot(createSnapshotOptions *CreateSnapshotOptions)`, func() {

			resourceGroupIdentityModel := &vpcv1.ResourceGroupIdentityByID{
				ID: core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345"),
			}

			volumeIdentityModel := &vpcv1.VolumeIdentityByID{
				ID: core.StringPtr("1a6b7274-678d-4dfb-8981-c71dd9d4daa5"),
			}

			createSnapshotOptions := &vpcv1.CreateSnapshotOptions{
				Name:          core.StringPtr("my-snapshot"),
				ResourceGroup: resourceGroupIdentityModel,
				SourceVolume:  volumeIdentityModel,
			}

			snapshot, response, err := vpcService.CreateSnapshot(createSnapshotOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(snapshot).ToNot(BeNil())

		})
	})

	Describe(`GetSnapshot - Retrieve a snapshot`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSnapshot(getSnapshotOptions *GetSnapshotOptions)`, func() {

			getSnapshotOptions := &vpcv1.GetSnapshotOptions{
				ID: core.StringPtr("testString"),
			}

			snapshot, response, err := vpcService.GetSnapshot(getSnapshotOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(snapshot).ToNot(BeNil())

		})
	})

	Describe(`UpdateSnapshot - Update a snapshot`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateSnapshot(updateSnapshotOptions *UpdateSnapshotOptions)`, func() {

			snapshotPatchModel := &vpcv1.SnapshotPatch{
				Name: core.StringPtr("my-snapshot"),
			}
			snapshotPatchModelAsPatch, asPatchErr := snapshotPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateSnapshotOptions := &vpcv1.UpdateSnapshotOptions{
				ID:            core.StringPtr("testString"),
				SnapshotPatch: snapshotPatchModelAsPatch,
			}

			snapshot, response, err := vpcService.UpdateSnapshot(updateSnapshotOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(snapshot).ToNot(BeNil())

		})
	})

	Describe(`ListShareProfiles - List all file share profiles`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListShareProfiles(listShareProfilesOptions *ListShareProfilesOptions)`, func() {

			listShareProfilesOptions := &vpcv1.ListShareProfilesOptions{
				Start: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(1)),
				Sort:  core.StringPtr("name"),
			}

			shareProfileCollection, response, err := vpcService.ListShareProfiles(listShareProfilesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(shareProfileCollection).ToNot(BeNil())

		})
	})

	Describe(`GetShareProfile - Retrieve a file share profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetShareProfile(getShareProfileOptions *GetShareProfileOptions)`, func() {

			getShareProfileOptions := &vpcv1.GetShareProfileOptions{
				Name: core.StringPtr("testString"),
			}

			shareProfile, response, err := vpcService.GetShareProfile(getShareProfileOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(shareProfile).ToNot(BeNil())

		})
	})

	Describe(`ListShares - List all file shares`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListShares(listSharesOptions *ListSharesOptions)`, func() {

			listSharesOptions := &vpcv1.ListSharesOptions{
				Start:           core.StringPtr("testString"),
				Limit:           core.Int64Ptr(int64(1)),
				ResourceGroupID: core.StringPtr("testString"),
				Name:            core.StringPtr("testString"),
				Sort:            core.StringPtr("name"),
			}

			shareCollection, response, err := vpcService.ListShares(listSharesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(shareCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateShare - Create a file share`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateShare(createShareOptions *CreateShareOptions)`, func() {

			encryptionKeyIdentityModel := &vpcv1.EncryptionKeyIdentityByCRN{
				CRN: core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/dffc98a0f1f0f95f6613b3b752286b87:e4a29d1a-2ef0-42a6-8fd2-350deb1c647e:key:5437653b-c4b1-447f-9646-b2a2a4cd6179"),
			}

			sharePrototypeProfileModel := &vpcv1.SharePrototypeProfileShareProfileIdentityByName{
				Name: core.StringPtr("tier-3iops"),
			}

			resourceGroupIdentityModel := &vpcv1.ResourceGroupIdentityByID{
				ID: core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345"),
			}

			subnetIdentityModel := &vpcv1.SubnetIdentityByID{
				ID: core.StringPtr("7ec86020-1c6e-4889-b3f0-a15f2e50f87e"),
			}

			vpcIdentityModel := &vpcv1.VPCIdentityByID{
				ID: core.StringPtr("7ec86020-1c6e-4889-b3f0-a15f2e50f87e"),
			}

			shareTargetPrototypeModel := &vpcv1.ShareTargetPrototype{
				Name:   core.StringPtr("my-share-target"),
				Subnet: subnetIdentityModel,
				VPC:    vpcIdentityModel,
			}

			zoneIdentityModel := &vpcv1.ZoneIdentityByName{
				Name: core.StringPtr("us-south-1"),
			}

			createShareOptions := &vpcv1.CreateShareOptions{
				EncryptionKey: encryptionKeyIdentityModel,
				Name:          core.StringPtr("my-share"),
				Profile:       sharePrototypeProfileModel,
				ResourceGroup: resourceGroupIdentityModel,
				Size:          core.Int64Ptr(int64(200)),
				Targets:       []vpcv1.ShareTargetPrototype{*shareTargetPrototypeModel},
				Zone:          zoneIdentityModel,
			}

			share, response, err := vpcService.CreateShare(createShareOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(share).ToNot(BeNil())

		})
	})

	Describe(`GetShare - Retrieve a file share`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetShare(getShareOptions *GetShareOptions)`, func() {

			getShareOptions := &vpcv1.GetShareOptions{
				ID: core.StringPtr("testString"),
			}

			share, response, err := vpcService.GetShare(getShareOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(share).ToNot(BeNil())

		})
	})

	Describe(`UpdateShare - Update a file share`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateShare(updateShareOptions *UpdateShareOptions)`, func() {

			sharePatchModel := &vpcv1.SharePatch{
				Name: core.StringPtr("my-share"),
				Size: core.Int64Ptr(int64(200)),
			}
			sharePatchModelAsPatch, asPatchErr := sharePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateShareOptions := &vpcv1.UpdateShareOptions{
				ID:         core.StringPtr("testString"),
				SharePatch: sharePatchModelAsPatch,
			}

			share, response, err := vpcService.UpdateShare(updateShareOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(share).ToNot(BeNil())

		})
	})

	Describe(`ListShareTargets - List all targets for a file share`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListShareTargets(listShareTargetsOptions *ListShareTargetsOptions)`, func() {

			listShareTargetsOptions := &vpcv1.ListShareTargetsOptions{
				ShareID: core.StringPtr("testString"),
				Name:    core.StringPtr("testString"),
			}

			shareTargetCollection, response, err := vpcService.ListShareTargets(listShareTargetsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(shareTargetCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateShareTarget - Create a target for a file share`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateShareTarget(createShareTargetOptions *CreateShareTargetOptions)`, func() {

			vpcIdentityModel := &vpcv1.VPCIdentityByID{
				ID: core.StringPtr("7ec86020-1c6e-4889-b3f0-a15f2e50f87e"),
			}

			subnetIdentityModel := &vpcv1.SubnetIdentityByID{
				ID: core.StringPtr("7ec86020-1c6e-4889-b3f0-a15f2e50f87e"),
			}

			createShareTargetOptions := &vpcv1.CreateShareTargetOptions{
				ShareID: core.StringPtr("testString"),
				VPC:     vpcIdentityModel,
				Name:    core.StringPtr("my-share-target"),
				Subnet:  subnetIdentityModel,
			}

			shareTarget, response, err := vpcService.CreateShareTarget(createShareTargetOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(shareTarget).ToNot(BeNil())

		})
	})

	Describe(`GetShareTarget - Retrieve a share target`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetShareTarget(getShareTargetOptions *GetShareTargetOptions)`, func() {

			getShareTargetOptions := &vpcv1.GetShareTargetOptions{
				ShareID: core.StringPtr("testString"),
				ID:      core.StringPtr("testString"),
			}

			shareTarget, response, err := vpcService.GetShareTarget(getShareTargetOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(shareTarget).ToNot(BeNil())

		})
	})

	Describe(`UpdateShareTarget - Update a share target`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateShareTarget(updateShareTargetOptions *UpdateShareTargetOptions)`, func() {

			shareTargetPatchModel := &vpcv1.ShareTargetPatch{
				Name: core.StringPtr("my-share-target"),
			}
			shareTargetPatchModelAsPatch, asPatchErr := shareTargetPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateShareTargetOptions := &vpcv1.UpdateShareTargetOptions{
				ShareID:          core.StringPtr("testString"),
				ID:               core.StringPtr("testString"),
				ShareTargetPatch: shareTargetPatchModelAsPatch,
			}

			shareTarget, response, err := vpcService.UpdateShareTarget(updateShareTargetOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(shareTarget).ToNot(BeNil())

		})
	})

	Describe(`ListRegions - List all regions`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListRegions(listRegionsOptions *ListRegionsOptions)`, func() {

			listRegionsOptions := &vpcv1.ListRegionsOptions{}

			regionCollection, response, err := vpcService.ListRegions(listRegionsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(regionCollection).ToNot(BeNil())

		})
	})

	Describe(`GetRegion - Retrieve a region`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetRegion(getRegionOptions *GetRegionOptions)`, func() {

			getRegionOptions := &vpcv1.GetRegionOptions{
				Name: core.StringPtr("testString"),
			}

			region, response, err := vpcService.GetRegion(getRegionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(region).ToNot(BeNil())

		})
	})

	Describe(`ListRegionZones - List all zones in a region`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListRegionZones(listRegionZonesOptions *ListRegionZonesOptions)`, func() {

			listRegionZonesOptions := &vpcv1.ListRegionZonesOptions{
				RegionName: core.StringPtr("testString"),
			}

			zoneCollection, response, err := vpcService.ListRegionZones(listRegionZonesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(zoneCollection).ToNot(BeNil())

		})
	})

	Describe(`GetRegionZone - Retrieve a zone`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetRegionZone(getRegionZoneOptions *GetRegionZoneOptions)`, func() {

			getRegionZoneOptions := &vpcv1.GetRegionZoneOptions{
				RegionName: core.StringPtr("testString"),
				Name:       core.StringPtr("testString"),
			}

			zone, response, err := vpcService.GetRegionZone(getRegionZoneOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(zone).ToNot(BeNil())

		})
	})

	Describe(`ListPublicGateways - List all public gateways`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListPublicGateways(listPublicGatewaysOptions *ListPublicGatewaysOptions)`, func() {

			listPublicGatewaysOptions := &vpcv1.ListPublicGatewaysOptions{
				Start:           core.StringPtr("testString"),
				Limit:           core.Int64Ptr(int64(1)),
				ResourceGroupID: core.StringPtr("testString"),
			}

			publicGatewayCollection, response, err := vpcService.ListPublicGateways(listPublicGatewaysOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(publicGatewayCollection).ToNot(BeNil())

		})
	})

	Describe(`CreatePublicGateway - Create a public gateway`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreatePublicGateway(createPublicGatewayOptions *CreatePublicGatewayOptions)`, func() {

			vpcIdentityModel := &vpcv1.VPCIdentityByID{
				ID: core.StringPtr("4727d842-f94f-4a2d-824a-9bc9b02c523b"),
			}

			zoneIdentityModel := &vpcv1.ZoneIdentityByName{
				Name: core.StringPtr("us-south-1"),
			}

			publicGatewayFloatingIPPrototypeModel := &vpcv1.PublicGatewayFloatingIPPrototypeFloatingIPIdentityFloatingIPIdentityByID{
				ID: core.StringPtr("39300233-9995-4806-89a5-3c1b6eb88689"),
			}

			resourceGroupIdentityModel := &vpcv1.ResourceGroupIdentityByID{
				ID: core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345"),
			}

			createPublicGatewayOptions := &vpcv1.CreatePublicGatewayOptions{
				VPC:           vpcIdentityModel,
				Zone:          zoneIdentityModel,
				FloatingIP:    publicGatewayFloatingIPPrototypeModel,
				Name:          core.StringPtr("my-public-gateway"),
				ResourceGroup: resourceGroupIdentityModel,
			}

			publicGateway, response, err := vpcService.CreatePublicGateway(createPublicGatewayOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(publicGateway).ToNot(BeNil())

		})
	})

	Describe(`GetPublicGateway - Retrieve a public gateway`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetPublicGateway(getPublicGatewayOptions *GetPublicGatewayOptions)`, func() {

			getPublicGatewayOptions := &vpcv1.GetPublicGatewayOptions{
				ID: core.StringPtr("testString"),
			}

			publicGateway, response, err := vpcService.GetPublicGateway(getPublicGatewayOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(publicGateway).ToNot(BeNil())

		})
	})

	Describe(`UpdatePublicGateway - Update a public gateway`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdatePublicGateway(updatePublicGatewayOptions *UpdatePublicGatewayOptions)`, func() {

			publicGatewayPatchModel := &vpcv1.PublicGatewayPatch{
				Name: core.StringPtr("my-public-gateway"),
			}
			publicGatewayPatchModelAsPatch, asPatchErr := publicGatewayPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updatePublicGatewayOptions := &vpcv1.UpdatePublicGatewayOptions{
				ID:                 core.StringPtr("testString"),
				PublicGatewayPatch: publicGatewayPatchModelAsPatch,
			}

			publicGateway, response, err := vpcService.UpdatePublicGateway(updatePublicGatewayOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(publicGateway).ToNot(BeNil())

		})
	})

	Describe(`ListFloatingIps - List all floating IPs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListFloatingIps(listFloatingIpsOptions *ListFloatingIpsOptions)`, func() {

			listFloatingIpsOptions := &vpcv1.ListFloatingIpsOptions{
				Start:           core.StringPtr("testString"),
				Limit:           core.Int64Ptr(int64(1)),
				ResourceGroupID: core.StringPtr("testString"),
			}

			floatingIPCollection, response, err := vpcService.ListFloatingIps(listFloatingIpsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(floatingIPCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateFloatingIP - Reserve a floating IP`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateFloatingIP(createFloatingIPOptions *CreateFloatingIPOptions)`, func() {

			resourceGroupIdentityModel := &vpcv1.ResourceGroupIdentityByID{
				ID: core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345"),
			}

			zoneIdentityModel := &vpcv1.ZoneIdentityByName{
				Name: core.StringPtr("us-south-1"),
			}

			floatingIPPrototypeModel := &vpcv1.FloatingIPPrototypeFloatingIPByZone{
				Name:          core.StringPtr("my-floating-ip"),
				ResourceGroup: resourceGroupIdentityModel,
				Zone:          zoneIdentityModel,
			}

			createFloatingIPOptions := &vpcv1.CreateFloatingIPOptions{
				FloatingIPPrototype: floatingIPPrototypeModel,
			}

			floatingIP, response, err := vpcService.CreateFloatingIP(createFloatingIPOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(floatingIP).ToNot(BeNil())

		})
	})

	Describe(`GetFloatingIP - Retrieve a floating IP`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetFloatingIP(getFloatingIPOptions *GetFloatingIPOptions)`, func() {

			getFloatingIPOptions := &vpcv1.GetFloatingIPOptions{
				ID: core.StringPtr("testString"),
			}

			floatingIP, response, err := vpcService.GetFloatingIP(getFloatingIPOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(floatingIP).ToNot(BeNil())

		})
	})

	Describe(`UpdateFloatingIP - Update a floating IP`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateFloatingIP(updateFloatingIPOptions *UpdateFloatingIPOptions)`, func() {

			floatingIPPatchTargetNetworkInterfaceIdentityModel := &vpcv1.FloatingIPPatchTargetNetworkInterfaceIdentityNetworkInterfaceIdentityByID{
				ID: core.StringPtr("69e55145-cc7d-4d8e-9e1f-cc3fb60b1793"),
			}

			floatingIPPatchModel := &vpcv1.FloatingIPPatch{
				Name:   core.StringPtr("my-floating-ip"),
				Target: floatingIPPatchTargetNetworkInterfaceIdentityModel,
			}
			floatingIPPatchModelAsPatch, asPatchErr := floatingIPPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateFloatingIPOptions := &vpcv1.UpdateFloatingIPOptions{
				ID:              core.StringPtr("testString"),
				FloatingIPPatch: floatingIPPatchModelAsPatch,
			}

			floatingIP, response, err := vpcService.UpdateFloatingIP(updateFloatingIPOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(floatingIP).ToNot(BeNil())

		})
	})

	Describe(`ListNetworkAcls - List all network ACLs`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListNetworkAcls(listNetworkAclsOptions *ListNetworkAclsOptions)`, func() {

			listNetworkAclsOptions := &vpcv1.ListNetworkAclsOptions{
				Start:           core.StringPtr("testString"),
				Limit:           core.Int64Ptr(int64(1)),
				ResourceGroupID: core.StringPtr("testString"),
			}

			networkACLCollection, response, err := vpcService.ListNetworkAcls(listNetworkAclsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkACLCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateNetworkACL - Create a network ACL`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateNetworkACL(createNetworkACLOptions *CreateNetworkACLOptions)`, func() {

			resourceGroupIdentityModel := &vpcv1.ResourceGroupIdentityByID{
				ID: core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345"),
			}

			vpcIdentityModel := &vpcv1.VPCIdentityByID{
				ID: core.StringPtr("f0aae929-7047-46d1-92e1-9102b07a7f6f"),
			}

			networkACLRulePrototypeNetworkACLContextModel := &vpcv1.NetworkACLRulePrototypeNetworkACLContextNetworkACLRuleProtocolIcmp{
				Action:      core.StringPtr("allow"),
				Destination: core.StringPtr("192.168.3.2/32"),
				Direction:   core.StringPtr("inbound"),
				Name:        core.StringPtr("my-rule-2"),
				Source:      core.StringPtr("192.168.3.2/32"),
				Code:        core.Int64Ptr(int64(0)),
				Protocol:    core.StringPtr("icmp"),
				Type:        core.Int64Ptr(int64(8)),
			}

			networkACLPrototypeModel := &vpcv1.NetworkACLPrototypeNetworkACLByRules{
				Name:          core.StringPtr("my-network-acl"),
				ResourceGroup: resourceGroupIdentityModel,
				VPC:           vpcIdentityModel,
				Rules:         []vpcv1.NetworkACLRulePrototypeNetworkACLContextIntf{networkACLRulePrototypeNetworkACLContextModel},
			}

			createNetworkACLOptions := &vpcv1.CreateNetworkACLOptions{
				NetworkACLPrototype: networkACLPrototypeModel,
			}

			networkACL, response, err := vpcService.CreateNetworkACL(createNetworkACLOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(networkACL).ToNot(BeNil())

		})
	})

	Describe(`GetNetworkACL - Retrieve a network ACL`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetNetworkACL(getNetworkACLOptions *GetNetworkACLOptions)`, func() {

			getNetworkACLOptions := &vpcv1.GetNetworkACLOptions{
				ID: core.StringPtr("testString"),
			}

			networkACL, response, err := vpcService.GetNetworkACL(getNetworkACLOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkACL).ToNot(BeNil())

		})
	})

	Describe(`UpdateNetworkACL - Update a network ACL`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateNetworkACL(updateNetworkACLOptions *UpdateNetworkACLOptions)`, func() {

			networkACLPatchModel := &vpcv1.NetworkACLPatch{
				Name: core.StringPtr("my-network-acl"),
			}
			networkACLPatchModelAsPatch, asPatchErr := networkACLPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateNetworkACLOptions := &vpcv1.UpdateNetworkACLOptions{
				ID:              core.StringPtr("testString"),
				NetworkACLPatch: networkACLPatchModelAsPatch,
			}

			networkACL, response, err := vpcService.UpdateNetworkACL(updateNetworkACLOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkACL).ToNot(BeNil())

		})
	})

	Describe(`ListNetworkACLRules - List all rules for a network ACL`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListNetworkACLRules(listNetworkACLRulesOptions *ListNetworkACLRulesOptions)`, func() {

			listNetworkACLRulesOptions := &vpcv1.ListNetworkACLRulesOptions{
				NetworkACLID: core.StringPtr("testString"),
				Start:        core.StringPtr("testString"),
				Limit:        core.Int64Ptr(int64(1)),
				Direction:    core.StringPtr("inbound"),
			}

			networkACLRuleCollection, response, err := vpcService.ListNetworkACLRules(listNetworkACLRulesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkACLRuleCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateNetworkACLRule - Create a rule for a network ACL`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateNetworkACLRule(createNetworkACLRuleOptions *CreateNetworkACLRuleOptions)`, func() {

			networkACLRuleBeforePrototypeModel := &vpcv1.NetworkACLRuleBeforePrototypeNetworkACLRuleIdentityByID{
				ID: core.StringPtr("8daca77a-4980-4d33-8f3e-7038797be8f9"),
			}

			networkACLRulePrototypeModel := &vpcv1.NetworkACLRulePrototypeNetworkACLRuleProtocolAll{
				Action:      core.StringPtr("allow"),
				Before:      networkACLRuleBeforePrototypeModel,
				Destination: core.StringPtr("192.168.3.2/32"),
				Direction:   core.StringPtr("inbound"),
				Name:        core.StringPtr("my-rule-2"),
				Source:      core.StringPtr("192.168.3.2/32"),
				Protocol:    core.StringPtr("all"),
			}

			createNetworkACLRuleOptions := &vpcv1.CreateNetworkACLRuleOptions{
				NetworkACLID:            core.StringPtr("testString"),
				NetworkACLRulePrototype: networkACLRulePrototypeModel,
			}

			networkACLRule, response, err := vpcService.CreateNetworkACLRule(createNetworkACLRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(networkACLRule).ToNot(BeNil())

		})
	})

	Describe(`GetNetworkACLRule - Retrieve a network ACL rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetNetworkACLRule(getNetworkACLRuleOptions *GetNetworkACLRuleOptions)`, func() {

			getNetworkACLRuleOptions := &vpcv1.GetNetworkACLRuleOptions{
				NetworkACLID: core.StringPtr("testString"),
				ID:           core.StringPtr("testString"),
			}

			networkACLRule, response, err := vpcService.GetNetworkACLRule(getNetworkACLRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkACLRule).ToNot(BeNil())

		})
	})

	Describe(`UpdateNetworkACLRule - Update a network ACL rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateNetworkACLRule(updateNetworkACLRuleOptions *UpdateNetworkACLRuleOptions)`, func() {

			networkACLRuleBeforePatchModel := &vpcv1.NetworkACLRuleBeforePatchNetworkACLRuleIdentityByID{
				ID: core.StringPtr("8daca77a-4980-4d33-8f3e-7038797be8f9"),
			}

			networkACLRulePatchModel := &vpcv1.NetworkACLRulePatch{
				Action:             core.StringPtr("allow"),
				Before:             networkACLRuleBeforePatchModel,
				Code:               core.Int64Ptr(int64(0)),
				Destination:        core.StringPtr("192.168.3.2/32"),
				DestinationPortMax: core.Int64Ptr(int64(22)),
				DestinationPortMin: core.Int64Ptr(int64(22)),
				Direction:          core.StringPtr("inbound"),
				Name:               core.StringPtr("my-rule-2"),
				Source:             core.StringPtr("192.168.3.2/32"),
				SourcePortMax:      core.Int64Ptr(int64(65535)),
				SourcePortMin:      core.Int64Ptr(int64(49152)),
				Type:               core.Int64Ptr(int64(8)),
			}
			networkACLRulePatchModelAsPatch, asPatchErr := networkACLRulePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateNetworkACLRuleOptions := &vpcv1.UpdateNetworkACLRuleOptions{
				NetworkACLID:        core.StringPtr("testString"),
				ID:                  core.StringPtr("testString"),
				NetworkACLRulePatch: networkACLRulePatchModelAsPatch,
			}

			networkACLRule, response, err := vpcService.UpdateNetworkACLRule(updateNetworkACLRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkACLRule).ToNot(BeNil())

		})
	})

	Describe(`ListSecurityGroups - List all security groups`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSecurityGroups(listSecurityGroupsOptions *ListSecurityGroupsOptions)`, func() {

			listSecurityGroupsOptions := &vpcv1.ListSecurityGroupsOptions{
				Start:           core.StringPtr("testString"),
				Limit:           core.Int64Ptr(int64(1)),
				ResourceGroupID: core.StringPtr("testString"),
				VPCID:           core.StringPtr("testString"),
				VPCCRN:          core.StringPtr("testString"),
				VPCName:         core.StringPtr("testString"),
			}

			securityGroupCollection, response, err := vpcService.ListSecurityGroups(listSecurityGroupsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(securityGroupCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateSecurityGroup - Create a security group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSecurityGroup(createSecurityGroupOptions *CreateSecurityGroupOptions)`, func() {

			vpcIdentityModel := &vpcv1.VPCIdentityByID{
				ID: core.StringPtr("4727d842-f94f-4a2d-824a-9bc9b02c523b"),
			}

			resourceGroupIdentityModel := &vpcv1.ResourceGroupIdentityByID{
				ID: core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345"),
			}

			securityGroupRuleRemotePrototypeModel := &vpcv1.SecurityGroupRuleRemotePrototypeIP{
				Address: core.StringPtr("192.168.3.4"),
			}

			securityGroupRulePrototypeModel := &vpcv1.SecurityGroupRulePrototypeSecurityGroupRuleProtocolTcpudp{
				Direction: core.StringPtr("inbound"),
				IPVersion: core.StringPtr("ipv4"),
				Remote:    securityGroupRuleRemotePrototypeModel,
				PortMax:   core.Int64Ptr(int64(22)),
				PortMin:   core.Int64Ptr(int64(22)),
				Protocol:  core.StringPtr("udp"),
			}

			createSecurityGroupOptions := &vpcv1.CreateSecurityGroupOptions{
				VPC:           vpcIdentityModel,
				Name:          core.StringPtr("my-security-group"),
				ResourceGroup: resourceGroupIdentityModel,
				Rules:         []vpcv1.SecurityGroupRulePrototypeIntf{securityGroupRulePrototypeModel},
			}

			securityGroup, response, err := vpcService.CreateSecurityGroup(createSecurityGroupOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(securityGroup).ToNot(BeNil())

		})
	})

	Describe(`GetSecurityGroup - Retrieve a security group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSecurityGroup(getSecurityGroupOptions *GetSecurityGroupOptions)`, func() {

			getSecurityGroupOptions := &vpcv1.GetSecurityGroupOptions{
				ID: core.StringPtr("testString"),
			}

			securityGroup, response, err := vpcService.GetSecurityGroup(getSecurityGroupOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(securityGroup).ToNot(BeNil())

		})
	})

	Describe(`UpdateSecurityGroup - Update a security group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateSecurityGroup(updateSecurityGroupOptions *UpdateSecurityGroupOptions)`, func() {

			securityGroupPatchModel := &vpcv1.SecurityGroupPatch{
				Name: core.StringPtr("my-security-group"),
			}
			securityGroupPatchModelAsPatch, asPatchErr := securityGroupPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateSecurityGroupOptions := &vpcv1.UpdateSecurityGroupOptions{
				ID:                 core.StringPtr("testString"),
				SecurityGroupPatch: securityGroupPatchModelAsPatch,
			}

			securityGroup, response, err := vpcService.UpdateSecurityGroup(updateSecurityGroupOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(securityGroup).ToNot(BeNil())

		})
	})

	Describe(`ListSecurityGroupNetworkInterfaces - List all network interfaces associated with a security group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSecurityGroupNetworkInterfaces(listSecurityGroupNetworkInterfacesOptions *ListSecurityGroupNetworkInterfacesOptions)`, func() {

			listSecurityGroupNetworkInterfacesOptions := &vpcv1.ListSecurityGroupNetworkInterfacesOptions{
				SecurityGroupID: core.StringPtr("testString"),
				Start:           core.StringPtr("testString"),
				Limit:           core.Int64Ptr(int64(1)),
			}

			networkInterfaceCollection, response, err := vpcService.ListSecurityGroupNetworkInterfaces(listSecurityGroupNetworkInterfacesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkInterfaceCollection).ToNot(BeNil())

		})
	})

	Describe(`GetSecurityGroupNetworkInterface - Retrieve a network interface in a security group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSecurityGroupNetworkInterface(getSecurityGroupNetworkInterfaceOptions *GetSecurityGroupNetworkInterfaceOptions)`, func() {

			getSecurityGroupNetworkInterfaceOptions := &vpcv1.GetSecurityGroupNetworkInterfaceOptions{
				SecurityGroupID: core.StringPtr("testString"),
				ID:              core.StringPtr("testString"),
			}

			networkInterface, response, err := vpcService.GetSecurityGroupNetworkInterface(getSecurityGroupNetworkInterfaceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkInterface).ToNot(BeNil())

		})
	})

	Describe(`AddSecurityGroupNetworkInterface - Add a network interface to a security group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AddSecurityGroupNetworkInterface(addSecurityGroupNetworkInterfaceOptions *AddSecurityGroupNetworkInterfaceOptions)`, func() {

			addSecurityGroupNetworkInterfaceOptions := &vpcv1.AddSecurityGroupNetworkInterfaceOptions{
				SecurityGroupID: core.StringPtr("testString"),
				ID:              core.StringPtr("testString"),
			}

			networkInterface, response, err := vpcService.AddSecurityGroupNetworkInterface(addSecurityGroupNetworkInterfaceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(networkInterface).ToNot(BeNil())

		})
	})

	Describe(`ListSecurityGroupRules - List all rules in a security group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSecurityGroupRules(listSecurityGroupRulesOptions *ListSecurityGroupRulesOptions)`, func() {

			listSecurityGroupRulesOptions := &vpcv1.ListSecurityGroupRulesOptions{
				SecurityGroupID: core.StringPtr("testString"),
			}

			securityGroupRuleCollection, response, err := vpcService.ListSecurityGroupRules(listSecurityGroupRulesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(securityGroupRuleCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateSecurityGroupRule - Create a rule for a security group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSecurityGroupRule(createSecurityGroupRuleOptions *CreateSecurityGroupRuleOptions)`, func() {

			securityGroupRuleRemotePrototypeModel := &vpcv1.SecurityGroupRuleRemotePrototypeIP{
				Address: core.StringPtr("192.168.3.4"),
			}

			securityGroupRulePrototypeModel := &vpcv1.SecurityGroupRulePrototypeSecurityGroupRuleProtocolTcpudp{
				Direction: core.StringPtr("inbound"),
				IPVersion: core.StringPtr("ipv4"),
				Remote:    securityGroupRuleRemotePrototypeModel,
				PortMax:   core.Int64Ptr(int64(22)),
				PortMin:   core.Int64Ptr(int64(22)),
				Protocol:  core.StringPtr("udp"),
			}

			createSecurityGroupRuleOptions := &vpcv1.CreateSecurityGroupRuleOptions{
				SecurityGroupID:            core.StringPtr("testString"),
				SecurityGroupRulePrototype: securityGroupRulePrototypeModel,
			}

			securityGroupRule, response, err := vpcService.CreateSecurityGroupRule(createSecurityGroupRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(securityGroupRule).ToNot(BeNil())

		})
	})

	Describe(`GetSecurityGroupRule - Retrieve a security group rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSecurityGroupRule(getSecurityGroupRuleOptions *GetSecurityGroupRuleOptions)`, func() {

			getSecurityGroupRuleOptions := &vpcv1.GetSecurityGroupRuleOptions{
				SecurityGroupID: core.StringPtr("testString"),
				ID:              core.StringPtr("testString"),
			}

			securityGroupRule, response, err := vpcService.GetSecurityGroupRule(getSecurityGroupRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(securityGroupRule).ToNot(BeNil())

		})
	})

	Describe(`UpdateSecurityGroupRule - Update a security group rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateSecurityGroupRule(updateSecurityGroupRuleOptions *UpdateSecurityGroupRuleOptions)`, func() {

			securityGroupRuleRemotePatchModel := &vpcv1.SecurityGroupRuleRemotePatchIP{
				Address: core.StringPtr("192.168.3.4"),
			}

			securityGroupRulePatchModel := &vpcv1.SecurityGroupRulePatch{
				Code:      core.Int64Ptr(int64(0)),
				Direction: core.StringPtr("inbound"),
				IPVersion: core.StringPtr("ipv4"),
				PortMax:   core.Int64Ptr(int64(22)),
				PortMin:   core.Int64Ptr(int64(22)),
				Remote:    securityGroupRuleRemotePatchModel,
				Type:      core.Int64Ptr(int64(8)),
			}
			securityGroupRulePatchModelAsPatch, asPatchErr := securityGroupRulePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateSecurityGroupRuleOptions := &vpcv1.UpdateSecurityGroupRuleOptions{
				SecurityGroupID:        core.StringPtr("testString"),
				ID:                     core.StringPtr("testString"),
				SecurityGroupRulePatch: securityGroupRulePatchModelAsPatch,
			}

			securityGroupRule, response, err := vpcService.UpdateSecurityGroupRule(updateSecurityGroupRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(securityGroupRule).ToNot(BeNil())

		})
	})

	Describe(`ListSecurityGroupTargets - List all targets associated with a security group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListSecurityGroupTargets(listSecurityGroupTargetsOptions *ListSecurityGroupTargetsOptions)`, func() {

			listSecurityGroupTargetsOptions := &vpcv1.ListSecurityGroupTargetsOptions{
				SecurityGroupID: core.StringPtr("testString"),
				Start:           core.StringPtr("testString"),
				Limit:           core.Int64Ptr(int64(1)),
			}

			securityGroupTargetCollection, response, err := vpcService.ListSecurityGroupTargets(listSecurityGroupTargetsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(securityGroupTargetCollection).ToNot(BeNil())

		})
	})

	Describe(`GetSecurityGroupTarget - Retrieve a security group target`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetSecurityGroupTarget(getSecurityGroupTargetOptions *GetSecurityGroupTargetOptions)`, func() {

			getSecurityGroupTargetOptions := &vpcv1.GetSecurityGroupTargetOptions{
				SecurityGroupID: core.StringPtr("testString"),
				ID:              core.StringPtr("testString"),
			}

			securityGroupTargetReference, response, err := vpcService.GetSecurityGroupTarget(getSecurityGroupTargetOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(securityGroupTargetReference).ToNot(BeNil())

		})
	})

	Describe(`CreateSecurityGroupTargetBinding - Add a target to a security group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateSecurityGroupTargetBinding(createSecurityGroupTargetBindingOptions *CreateSecurityGroupTargetBindingOptions)`, func() {

			createSecurityGroupTargetBindingOptions := &vpcv1.CreateSecurityGroupTargetBindingOptions{
				SecurityGroupID: core.StringPtr("testString"),
				ID:              core.StringPtr("testString"),
			}

			securityGroupTargetReference, response, err := vpcService.CreateSecurityGroupTargetBinding(createSecurityGroupTargetBindingOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(securityGroupTargetReference).ToNot(BeNil())

		})
	})

	Describe(`ListIkePolicies - List all IKE policies`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListIkePolicies(listIkePoliciesOptions *ListIkePoliciesOptions)`, func() {

			listIkePoliciesOptions := &vpcv1.ListIkePoliciesOptions{
				Start: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(1)),
			}

			ikePolicyCollection, response, err := vpcService.ListIkePolicies(listIkePoliciesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ikePolicyCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateIkePolicy - Create an IKE policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateIkePolicy(createIkePolicyOptions *CreateIkePolicyOptions)`, func() {

			resourceGroupIdentityModel := &vpcv1.ResourceGroupIdentityByID{
				ID: core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345"),
			}

			createIkePolicyOptions := &vpcv1.CreateIkePolicyOptions{
				AuthenticationAlgorithm: core.StringPtr("md5"),
				DhGroup:                 core.Int64Ptr(int64(2)),
				EncryptionAlgorithm:     core.StringPtr("triple_des"),
				IkeVersion:              core.Int64Ptr(int64(1)),
				KeyLifetime:             core.Int64Ptr(int64(28800)),
				Name:                    core.StringPtr("my-ike-policy"),
				ResourceGroup:           resourceGroupIdentityModel,
			}

			ikePolicy, response, err := vpcService.CreateIkePolicy(createIkePolicyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(ikePolicy).ToNot(BeNil())

		})
	})

	Describe(`GetIkePolicy - Retrieve an IKE policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetIkePolicy(getIkePolicyOptions *GetIkePolicyOptions)`, func() {

			getIkePolicyOptions := &vpcv1.GetIkePolicyOptions{
				ID: core.StringPtr("testString"),
			}

			ikePolicy, response, err := vpcService.GetIkePolicy(getIkePolicyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ikePolicy).ToNot(BeNil())

		})
	})

	Describe(`UpdateIkePolicy - Update an IKE policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateIkePolicy(updateIkePolicyOptions *UpdateIkePolicyOptions)`, func() {

			updateIkePolicyOptions := &vpcv1.UpdateIkePolicyOptions{
				ID:             core.StringPtr("testString"),
				IkePolicyPatch: make(map[string]interface{}),
			}

			ikePolicy, response, err := vpcService.UpdateIkePolicy(updateIkePolicyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ikePolicy).ToNot(BeNil())

		})
	})

	Describe(`ListIkePolicyConnections - List all VPN gateway connections that use a specified IKE policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListIkePolicyConnections(listIkePolicyConnectionsOptions *ListIkePolicyConnectionsOptions)`, func() {

			listIkePolicyConnectionsOptions := &vpcv1.ListIkePolicyConnectionsOptions{
				ID: core.StringPtr("testString"),
			}

			vpnGatewayConnectionCollection, response, err := vpcService.ListIkePolicyConnections(listIkePolicyConnectionsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnGatewayConnectionCollection).ToNot(BeNil())

		})
	})

	Describe(`ListIpsecPolicies - List all IPsec policies`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListIpsecPolicies(listIpsecPoliciesOptions *ListIpsecPoliciesOptions)`, func() {

			listIpsecPoliciesOptions := &vpcv1.ListIpsecPoliciesOptions{
				Start: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(1)),
			}

			iPsecPolicyCollection, response, err := vpcService.ListIpsecPolicies(listIpsecPoliciesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(iPsecPolicyCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateIpsecPolicy - Create an IPsec policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateIpsecPolicy(createIpsecPolicyOptions *CreateIpsecPolicyOptions)`, func() {

			resourceGroupIdentityModel := &vpcv1.ResourceGroupIdentityByID{
				ID: core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345"),
			}

			createIpsecPolicyOptions := &vpcv1.CreateIpsecPolicyOptions{
				AuthenticationAlgorithm: core.StringPtr("md5"),
				EncryptionAlgorithm:     core.StringPtr("triple_des"),
				Pfs:                     core.StringPtr("disabled"),
				KeyLifetime:             core.Int64Ptr(int64(3600)),
				Name:                    core.StringPtr("my-ipsec-policy"),
				ResourceGroup:           resourceGroupIdentityModel,
			}

			iPsecPolicy, response, err := vpcService.CreateIpsecPolicy(createIpsecPolicyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(iPsecPolicy).ToNot(BeNil())

		})
	})

	Describe(`GetIpsecPolicy - Retrieve an IPsec policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetIpsecPolicy(getIpsecPolicyOptions *GetIpsecPolicyOptions)`, func() {

			getIpsecPolicyOptions := &vpcv1.GetIpsecPolicyOptions{
				ID: core.StringPtr("testString"),
			}

			iPsecPolicy, response, err := vpcService.GetIpsecPolicy(getIpsecPolicyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(iPsecPolicy).ToNot(BeNil())

		})
	})

	Describe(`UpdateIpsecPolicy - Update an IPsec policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateIpsecPolicy(updateIpsecPolicyOptions *UpdateIpsecPolicyOptions)`, func() {

			iPsecPolicyPatchModel := &vpcv1.IPsecPolicyPatch{
				AuthenticationAlgorithm: core.StringPtr("md5"),
				EncryptionAlgorithm:     core.StringPtr("triple_des"),
				KeyLifetime:             core.Int64Ptr(int64(3600)),
				Name:                    core.StringPtr("my-ipsec-policy"),
				Pfs:                     core.StringPtr("disabled"),
			}
			iPsecPolicyPatchModelAsPatch, asPatchErr := iPsecPolicyPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateIpsecPolicyOptions := &vpcv1.UpdateIpsecPolicyOptions{
				ID:               core.StringPtr("testString"),
				IPsecPolicyPatch: iPsecPolicyPatchModelAsPatch,
			}

			iPsecPolicy, response, err := vpcService.UpdateIpsecPolicy(updateIpsecPolicyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(iPsecPolicy).ToNot(BeNil())

		})
	})

	Describe(`ListIpsecPolicyConnections - List all VPN gateway connections that use a specified IPsec policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListIpsecPolicyConnections(listIpsecPolicyConnectionsOptions *ListIpsecPolicyConnectionsOptions)`, func() {

			listIpsecPolicyConnectionsOptions := &vpcv1.ListIpsecPolicyConnectionsOptions{
				ID: core.StringPtr("testString"),
			}

			vpnGatewayConnectionCollection, response, err := vpcService.ListIpsecPolicyConnections(listIpsecPolicyConnectionsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnGatewayConnectionCollection).ToNot(BeNil())

		})
	})

	Describe(`ListVPNGateways - List all VPN gateways`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListVPNGateways(listVPNGatewaysOptions *ListVPNGatewaysOptions)`, func() {

			listVPNGatewaysOptions := &vpcv1.ListVPNGatewaysOptions{
				Start:           core.StringPtr("testString"),
				Limit:           core.Int64Ptr(int64(1)),
				ResourceGroupID: core.StringPtr("testString"),
				Sort:            core.StringPtr("name"),
				Mode:            core.StringPtr("route"),
			}

			vpnGatewayCollection, response, err := vpcService.ListVPNGateways(listVPNGatewaysOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnGatewayCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateVPNGateway - Create a VPN gateway`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateVPNGateway(createVPNGatewayOptions *CreateVPNGatewayOptions)`, func() {

			resourceGroupIdentityModel := &vpcv1.ResourceGroupIdentityByID{
				ID: core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345"),
			}

			subnetIdentityModel := &vpcv1.SubnetIdentityByID{
				ID: core.StringPtr("7ec86020-1c6e-4889-b3f0-a15f2e50f87e"),
			}

			vpnGatewayLocalAsnModel := &vpcv1.VPNGatewayLocalAsnTwoOctetPrivateAsn{}

			vpnGatewayPrototypeModel := &vpcv1.VPNGatewayPrototypeVPNGatewayRouteModePrototype{
				Name:            core.StringPtr("my-vpn-gateway"),
				ResourceGroup:   resourceGroupIdentityModel,
				Subnet:          subnetIdentityModel,
				AdvertisedCIDRs: []string{"192.168.3.0/24"},
				LocalAsn:        vpnGatewayLocalAsnModel,
				Mode:            core.StringPtr("route"),
			}

			createVPNGatewayOptions := &vpcv1.CreateVPNGatewayOptions{
				VPNGatewayPrototype: vpnGatewayPrototypeModel,
			}

			vpnGateway, response, err := vpcService.CreateVPNGateway(createVPNGatewayOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(vpnGateway).ToNot(BeNil())

		})
	})

	Describe(`GetVPNGateway - Retrieve a VPN gateway`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetVPNGateway(getVPNGatewayOptions *GetVPNGatewayOptions)`, func() {

			getVPNGatewayOptions := &vpcv1.GetVPNGatewayOptions{
				ID: core.StringPtr("testString"),
			}

			vpnGateway, response, err := vpcService.GetVPNGateway(getVPNGatewayOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnGateway).ToNot(BeNil())

		})
	})

	Describe(`UpdateVPNGateway - Update a VPN gateway`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateVPNGateway(updateVPNGatewayOptions *UpdateVPNGatewayOptions)`, func() {

			vpnGatewayLocalAsnModel := &vpcv1.VPNGatewayLocalAsnTwoOctetPrivateAsn{}

			vpnGatewayPatchModel := &vpcv1.VPNGatewayPatch{
				LocalAsn: vpnGatewayLocalAsnModel,
				Name:     core.StringPtr("my-vpn-gateway"),
			}
			vpnGatewayPatchModelAsPatch, asPatchErr := vpnGatewayPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateVPNGatewayOptions := &vpcv1.UpdateVPNGatewayOptions{
				ID:              core.StringPtr("testString"),
				VPNGatewayPatch: vpnGatewayPatchModelAsPatch,
			}

			vpnGateway, response, err := vpcService.UpdateVPNGateway(updateVPNGatewayOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnGateway).ToNot(BeNil())

		})
	})

	Describe(`ListVPNGatewayAdvertisedCIDRs - List all advertised CIDRs for a VPN gateway`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListVPNGatewayAdvertisedCIDRs(listVPNGatewayAdvertisedCIDRsOptions *ListVPNGatewayAdvertisedCIDRsOptions)`, func() {

			listVPNGatewayAdvertisedCIDRsOptions := &vpcv1.ListVPNGatewayAdvertisedCIDRsOptions{
				VPNGatewayID: core.StringPtr("testString"),
			}

			vpnGatewayAdvertisedCIDRs, response, err := vpcService.ListVPNGatewayAdvertisedCIDRs(listVPNGatewayAdvertisedCIDRsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnGatewayAdvertisedCIDRs).ToNot(BeNil())

		})
	})

	Describe(`CheckVPNGatewayAdvertisedCIDR - Check if the specified advertised CIDR exists on a VPN gateway`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CheckVPNGatewayAdvertisedCIDR(checkVPNGatewayAdvertisedCIDROptions *CheckVPNGatewayAdvertisedCIDROptions)`, func() {

			checkVPNGatewayAdvertisedCIDROptions := &vpcv1.CheckVPNGatewayAdvertisedCIDROptions{
				VPNGatewayID: core.StringPtr("testString"),
				CIDRPrefix:   core.StringPtr("testString"),
				PrefixLength: core.StringPtr("testString"),
			}

			response, err := vpcService.CheckVPNGatewayAdvertisedCIDR(checkVPNGatewayAdvertisedCIDROptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`AddVPNGatewayAdvertisedCIDR - Set an advertised CIDR on a VPN gateway`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AddVPNGatewayAdvertisedCIDR(addVPNGatewayAdvertisedCIDROptions *AddVPNGatewayAdvertisedCIDROptions)`, func() {

			addVPNGatewayAdvertisedCIDROptions := &vpcv1.AddVPNGatewayAdvertisedCIDROptions{
				VPNGatewayID: core.StringPtr("testString"),
				CIDRPrefix:   core.StringPtr("testString"),
				PrefixLength: core.StringPtr("testString"),
			}

			response, err := vpcService.AddVPNGatewayAdvertisedCIDR(addVPNGatewayAdvertisedCIDROptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`ListVPNGatewayConnections - List all connections of a VPN gateway`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListVPNGatewayConnections(listVPNGatewayConnectionsOptions *ListVPNGatewayConnectionsOptions)`, func() {

			listVPNGatewayConnectionsOptions := &vpcv1.ListVPNGatewayConnectionsOptions{
				VPNGatewayID: core.StringPtr("testString"),
				Status:       core.StringPtr("testString"),
			}

			vpnGatewayConnectionCollection, response, err := vpcService.ListVPNGatewayConnections(listVPNGatewayConnectionsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnGatewayConnectionCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateVPNGatewayConnection - Create a connection for a VPN gateway`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateVPNGatewayConnection(createVPNGatewayConnectionOptions *CreateVPNGatewayConnectionOptions)`, func() {

			vpnGatewayConnectionDpdPrototypeModel := &vpcv1.VPNGatewayConnectionDpdPrototype{
				Action:   core.StringPtr("restart"),
				Interval: core.Int64Ptr(int64(30)),
				Timeout:  core.Int64Ptr(int64(120)),
			}

			ikePolicyIdentityModel := &vpcv1.IkePolicyIdentityByID{
				ID: core.StringPtr("ddf51bec-3424-11e8-b467-0ed5f89f718b"),
			}

			iPsecPolicyIdentityModel := &vpcv1.IPsecPolicyIdentityByID{
				ID: core.StringPtr("ddf51bec-3424-11e8-b467-0ed5f89f718b"),
			}

			vpnGatewayConnectionPrototypeModel := &vpcv1.VPNGatewayConnectionPrototypeVPNGatewayConnectionStaticRouteModePrototype{
				AdminStateUp:      core.BoolPtr(true),
				DeadPeerDetection: vpnGatewayConnectionDpdPrototypeModel,
				IkePolicy:         ikePolicyIdentityModel,
				IpsecPolicy:       iPsecPolicyIdentityModel,
				Name:              core.StringPtr("my-vpn-connection"),
				PeerAddress:       core.StringPtr("169.21.50.5"),
				Psk:               core.StringPtr("lkj14b1oi0alcniejkso"),
				RoutingProtocol:   core.StringPtr("none"),
			}

			createVPNGatewayConnectionOptions := &vpcv1.CreateVPNGatewayConnectionOptions{
				VPNGatewayID:                  core.StringPtr("testString"),
				VPNGatewayConnectionPrototype: vpnGatewayConnectionPrototypeModel,
			}

			vpnGatewayConnection, response, err := vpcService.CreateVPNGatewayConnection(createVPNGatewayConnectionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(vpnGatewayConnection).ToNot(BeNil())

		})
	})

	Describe(`GetVPNGatewayConnection - Retrieve a VPN gateway connection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetVPNGatewayConnection(getVPNGatewayConnectionOptions *GetVPNGatewayConnectionOptions)`, func() {

			getVPNGatewayConnectionOptions := &vpcv1.GetVPNGatewayConnectionOptions{
				VPNGatewayID: core.StringPtr("testString"),
				ID:           core.StringPtr("testString"),
			}

			vpnGatewayConnection, response, err := vpcService.GetVPNGatewayConnection(getVPNGatewayConnectionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnGatewayConnection).ToNot(BeNil())

		})
	})

	Describe(`UpdateVPNGatewayConnection - Update a VPN gateway connection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateVPNGatewayConnection(updateVPNGatewayConnectionOptions *UpdateVPNGatewayConnectionOptions)`, func() {

			vpnGatewayConnectionDpdPrototypeModel := &vpcv1.VPNGatewayConnectionDpdPrototype{
				Action:   core.StringPtr("restart"),
				Interval: core.Int64Ptr(int64(30)),
				Timeout:  core.Int64Ptr(int64(120)),
			}

			ikePolicyIdentityModel := &vpcv1.IkePolicyIdentityByID{
				ID: core.StringPtr("ddf51bec-3424-11e8-b467-0ed5f89f718b"),
			}

			iPsecPolicyIdentityModel := &vpcv1.IPsecPolicyIdentityByID{
				ID: core.StringPtr("ddf51bec-3424-11e8-b467-0ed5f89f718b"),
			}

			vpnGatewayConnectionPatchModel := &vpcv1.VPNGatewayConnectionPatchVPNGatewayConnectionStaticRouteModePatch{
				AdminStateUp:      core.BoolPtr(true),
				DeadPeerDetection: vpnGatewayConnectionDpdPrototypeModel,
				IkePolicy:         ikePolicyIdentityModel,
				IpsecPolicy:       iPsecPolicyIdentityModel,
				Name:              core.StringPtr("my-vpn-connection"),
				PeerAddress:       core.StringPtr("169.21.50.5"),
				Psk:               core.StringPtr("lkj14b1oi0alcniejkso"),
				RoutingProtocol:   core.StringPtr("none"),
			}
			vpnGatewayConnectionPatchModelAsPatch, asPatchErr := vpnGatewayConnectionPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateVPNGatewayConnectionOptions := &vpcv1.UpdateVPNGatewayConnectionOptions{
				VPNGatewayID:              core.StringPtr("testString"),
				ID:                        core.StringPtr("testString"),
				VPNGatewayConnectionPatch: vpnGatewayConnectionPatchModelAsPatch,
				IfMatch:                   core.StringPtr("96d225c4-56bd-43d9-98fc-d7148e5c5028"),
			}

			vpnGatewayConnection, response, err := vpcService.UpdateVPNGatewayConnection(updateVPNGatewayConnectionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnGatewayConnection).ToNot(BeNil())

		})
	})

	Describe(`ListVPNGatewayConnectionLocalCIDRs - List all local CIDRs for a VPN gateway connection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListVPNGatewayConnectionLocalCIDRs(listVPNGatewayConnectionLocalCIDRsOptions *ListVPNGatewayConnectionLocalCIDRsOptions)`, func() {

			listVPNGatewayConnectionLocalCIDRsOptions := &vpcv1.ListVPNGatewayConnectionLocalCIDRsOptions{
				VPNGatewayID: core.StringPtr("testString"),
				ID:           core.StringPtr("testString"),
			}

			vpnGatewayConnectionLocalCIDRs, response, err := vpcService.ListVPNGatewayConnectionLocalCIDRs(listVPNGatewayConnectionLocalCIDRsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnGatewayConnectionLocalCIDRs).ToNot(BeNil())

		})
	})

	Describe(`CheckVPNGatewayConnectionLocalCIDR - Check if the specified local CIDR exists on a VPN gateway connection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CheckVPNGatewayConnectionLocalCIDR(checkVPNGatewayConnectionLocalCIDROptions *CheckVPNGatewayConnectionLocalCIDROptions)`, func() {

			checkVPNGatewayConnectionLocalCIDROptions := &vpcv1.CheckVPNGatewayConnectionLocalCIDROptions{
				VPNGatewayID: core.StringPtr("testString"),
				ID:           core.StringPtr("testString"),
				CIDRPrefix:   core.StringPtr("testString"),
				PrefixLength: core.StringPtr("testString"),
			}

			response, err := vpcService.CheckVPNGatewayConnectionLocalCIDR(checkVPNGatewayConnectionLocalCIDROptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`AddVPNGatewayConnectionLocalCIDR - Set a local CIDR on a VPN gateway connection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AddVPNGatewayConnectionLocalCIDR(addVPNGatewayConnectionLocalCIDROptions *AddVPNGatewayConnectionLocalCIDROptions)`, func() {

			addVPNGatewayConnectionLocalCIDROptions := &vpcv1.AddVPNGatewayConnectionLocalCIDROptions{
				VPNGatewayID: core.StringPtr("testString"),
				ID:           core.StringPtr("testString"),
				CIDRPrefix:   core.StringPtr("testString"),
				PrefixLength: core.StringPtr("testString"),
			}

			response, err := vpcService.AddVPNGatewayConnectionLocalCIDR(addVPNGatewayConnectionLocalCIDROptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`ListVPNGatewayConnectionPeerCIDRs - List all peer CIDRs for a VPN gateway connection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListVPNGatewayConnectionPeerCIDRs(listVPNGatewayConnectionPeerCIDRsOptions *ListVPNGatewayConnectionPeerCIDRsOptions)`, func() {

			listVPNGatewayConnectionPeerCIDRsOptions := &vpcv1.ListVPNGatewayConnectionPeerCIDRsOptions{
				VPNGatewayID: core.StringPtr("testString"),
				ID:           core.StringPtr("testString"),
			}

			vpnGatewayConnectionPeerCIDRs, response, err := vpcService.ListVPNGatewayConnectionPeerCIDRs(listVPNGatewayConnectionPeerCIDRsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnGatewayConnectionPeerCIDRs).ToNot(BeNil())

		})
	})

	Describe(`CheckVPNGatewayConnectionPeerCIDR - Check if the specified peer CIDR exists on a VPN gateway connection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CheckVPNGatewayConnectionPeerCIDR(checkVPNGatewayConnectionPeerCIDROptions *CheckVPNGatewayConnectionPeerCIDROptions)`, func() {

			checkVPNGatewayConnectionPeerCIDROptions := &vpcv1.CheckVPNGatewayConnectionPeerCIDROptions{
				VPNGatewayID: core.StringPtr("testString"),
				ID:           core.StringPtr("testString"),
				CIDRPrefix:   core.StringPtr("testString"),
				PrefixLength: core.StringPtr("testString"),
			}

			response, err := vpcService.CheckVPNGatewayConnectionPeerCIDR(checkVPNGatewayConnectionPeerCIDROptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`AddVPNGatewayConnectionPeerCIDR - Set a peer CIDR on a VPN gateway connection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AddVPNGatewayConnectionPeerCIDR(addVPNGatewayConnectionPeerCIDROptions *AddVPNGatewayConnectionPeerCIDROptions)`, func() {

			addVPNGatewayConnectionPeerCIDROptions := &vpcv1.AddVPNGatewayConnectionPeerCIDROptions{
				VPNGatewayID: core.StringPtr("testString"),
				ID:           core.StringPtr("testString"),
				CIDRPrefix:   core.StringPtr("testString"),
				PrefixLength: core.StringPtr("testString"),
			}

			response, err := vpcService.AddVPNGatewayConnectionPeerCIDR(addVPNGatewayConnectionPeerCIDROptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`ListLoadBalancerProfiles - List all load balancer profiles`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListLoadBalancerProfiles(listLoadBalancerProfilesOptions *ListLoadBalancerProfilesOptions)`, func() {

			listLoadBalancerProfilesOptions := &vpcv1.ListLoadBalancerProfilesOptions{
				Start: core.StringPtr("testString"),
				Limit: core.Int64Ptr(int64(1)),
			}

			loadBalancerProfileCollection, response, err := vpcService.ListLoadBalancerProfiles(listLoadBalancerProfilesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerProfileCollection).ToNot(BeNil())

		})
	})

	Describe(`GetLoadBalancerProfile - Retrieve a load balancer profile`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetLoadBalancerProfile(getLoadBalancerProfileOptions *GetLoadBalancerProfileOptions)`, func() {

			getLoadBalancerProfileOptions := &vpcv1.GetLoadBalancerProfileOptions{
				Name: core.StringPtr("testString"),
			}

			loadBalancerProfile, response, err := vpcService.GetLoadBalancerProfile(getLoadBalancerProfileOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerProfile).ToNot(BeNil())

		})
	})

	Describe(`ListLoadBalancers - List all load balancers`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListLoadBalancers(listLoadBalancersOptions *ListLoadBalancersOptions)`, func() {

			listLoadBalancersOptions := &vpcv1.ListLoadBalancersOptions{
				Start:           core.StringPtr("testString"),
				Limit:           core.Int64Ptr(int64(1)),
				ResourceGroupID: core.StringPtr("testString"),
			}

			loadBalancerCollection, response, err := vpcService.ListLoadBalancers(listLoadBalancersOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateLoadBalancer - Create a load balancer`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateLoadBalancer(createLoadBalancerOptions *CreateLoadBalancerOptions)`, func() {

			subnetIdentityModel := &vpcv1.SubnetIdentityByID{
				ID: core.StringPtr("7ec86020-1c6e-4889-b3f0-a15f2e50f87e"),
			}

			loadBalancerPoolIdentityByNameModel := &vpcv1.LoadBalancerPoolIdentityByName{
				Name: core.StringPtr("my-load-balancer-pool"),
			}

			loadBalancerListenerPrototypeLoadBalancerContextModel := &vpcv1.LoadBalancerListenerPrototypeLoadBalancerContext{
				AcceptProxyProtocol: core.BoolPtr(true),
				ConnectionLimit:     core.Int64Ptr(int64(2000)),
				DefaultPool:         loadBalancerPoolIdentityByNameModel,
				Port:                core.Int64Ptr(int64(443)),
				Protocol:            core.StringPtr("http"),
			}

			loadBalancerLoggingDatapathModel := &vpcv1.LoadBalancerLoggingDatapath{
				Active: core.BoolPtr(true),
			}

			loadBalancerLoggingModel := &vpcv1.LoadBalancerLogging{
				Datapath: loadBalancerLoggingDatapathModel,
			}

			loadBalancerPoolHealthMonitorPrototypeModel := &vpcv1.LoadBalancerPoolHealthMonitorPrototype{
				Delay:      core.Int64Ptr(int64(5)),
				MaxRetries: core.Int64Ptr(int64(2)),
				Port:       core.Int64Ptr(int64(22)),
				Timeout:    core.Int64Ptr(int64(2)),
				Type:       core.StringPtr("http"),
				URLPath:    core.StringPtr("/"),
			}

			loadBalancerPoolMemberTargetPrototypeModel := &vpcv1.LoadBalancerPoolMemberTargetPrototypeInstanceIdentityInstanceIdentityByID{
				ID: core.StringPtr("1e09281b-f177-46fb-baf1-bc152b2e391a"),
			}

			loadBalancerPoolMemberPrototypeModel := &vpcv1.LoadBalancerPoolMemberPrototype{
				Port:   core.Int64Ptr(int64(80)),
				Target: loadBalancerPoolMemberTargetPrototypeModel,
				Weight: core.Int64Ptr(int64(50)),
			}

			loadBalancerPoolSessionPersistencePrototypeModel := &vpcv1.LoadBalancerPoolSessionPersistencePrototype{
				Type: core.StringPtr("source_ip"),
			}

			loadBalancerPoolPrototypeModel := &vpcv1.LoadBalancerPoolPrototype{
				Algorithm:          core.StringPtr("least_connections"),
				HealthMonitor:      loadBalancerPoolHealthMonitorPrototypeModel,
				Members:            []vpcv1.LoadBalancerPoolMemberPrototype{*loadBalancerPoolMemberPrototypeModel},
				Name:               core.StringPtr("my-load-balancer-pool"),
				Protocol:           core.StringPtr("http"),
				ProxyProtocol:      core.StringPtr("disabled"),
				SessionPersistence: loadBalancerPoolSessionPersistencePrototypeModel,
			}

			loadBalancerProfileIdentityModel := &vpcv1.LoadBalancerProfileIdentityByName{
				Name: core.StringPtr("network-fixed"),
			}

			resourceGroupIdentityModel := &vpcv1.ResourceGroupIdentityByID{
				ID: core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345"),
			}

			securityGroupIdentityModel := &vpcv1.SecurityGroupIdentityByID{
				ID: core.StringPtr("be5df5ca-12a0-494b-907e-aa6ec2bfa271"),
			}

			createLoadBalancerOptions := &vpcv1.CreateLoadBalancerOptions{
				IsPublic:       core.BoolPtr(true),
				Subnets:        []vpcv1.SubnetIdentityIntf{subnetIdentityModel},
				Listeners:      []vpcv1.LoadBalancerListenerPrototypeLoadBalancerContext{*loadBalancerListenerPrototypeLoadBalancerContextModel},
				Logging:        loadBalancerLoggingModel,
				Name:           core.StringPtr("my-load-balancer"),
				Pools:          []vpcv1.LoadBalancerPoolPrototype{*loadBalancerPoolPrototypeModel},
				Profile:        loadBalancerProfileIdentityModel,
				ResourceGroup:  resourceGroupIdentityModel,
				SecurityGroups: []vpcv1.SecurityGroupIdentityIntf{securityGroupIdentityModel},
			}

			loadBalancer, response, err := vpcService.CreateLoadBalancer(createLoadBalancerOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(loadBalancer).ToNot(BeNil())

		})
	})

	Describe(`GetLoadBalancer - Retrieve a load balancer`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetLoadBalancer(getLoadBalancerOptions *GetLoadBalancerOptions)`, func() {

			getLoadBalancerOptions := &vpcv1.GetLoadBalancerOptions{
				ID: core.StringPtr("testString"),
			}

			loadBalancer, response, err := vpcService.GetLoadBalancer(getLoadBalancerOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancer).ToNot(BeNil())

		})
	})

	Describe(`UpdateLoadBalancer - Update a load balancer`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateLoadBalancer(updateLoadBalancerOptions *UpdateLoadBalancerOptions)`, func() {

			loadBalancerLoggingDatapathModel := &vpcv1.LoadBalancerLoggingDatapath{
				Active: core.BoolPtr(true),
			}

			loadBalancerLoggingModel := &vpcv1.LoadBalancerLogging{
				Datapath: loadBalancerLoggingDatapathModel,
			}

			loadBalancerPatchModel := &vpcv1.LoadBalancerPatch{
				Logging: loadBalancerLoggingModel,
				Name:    core.StringPtr("my-load-balancer"),
			}
			loadBalancerPatchModelAsPatch, asPatchErr := loadBalancerPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateLoadBalancerOptions := &vpcv1.UpdateLoadBalancerOptions{
				ID:                core.StringPtr("testString"),
				LoadBalancerPatch: loadBalancerPatchModelAsPatch,
			}

			loadBalancer, response, err := vpcService.UpdateLoadBalancer(updateLoadBalancerOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancer).ToNot(BeNil())

		})
	})

	Describe(`GetLoadBalancerStatistics - List all statistics of a load balancer`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetLoadBalancerStatistics(getLoadBalancerStatisticsOptions *GetLoadBalancerStatisticsOptions)`, func() {

			getLoadBalancerStatisticsOptions := &vpcv1.GetLoadBalancerStatisticsOptions{
				ID: core.StringPtr("testString"),
			}

			loadBalancerStatistics, response, err := vpcService.GetLoadBalancerStatistics(getLoadBalancerStatisticsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerStatistics).ToNot(BeNil())

		})
	})

	Describe(`ListLoadBalancerListeners - List all listeners for a load balancer`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListLoadBalancerListeners(listLoadBalancerListenersOptions *ListLoadBalancerListenersOptions)`, func() {

			listLoadBalancerListenersOptions := &vpcv1.ListLoadBalancerListenersOptions{
				LoadBalancerID: core.StringPtr("testString"),
			}

			loadBalancerListenerCollection, response, err := vpcService.ListLoadBalancerListeners(listLoadBalancerListenersOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerListenerCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateLoadBalancerListener - Create a listener for a load balancer`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateLoadBalancerListener(createLoadBalancerListenerOptions *CreateLoadBalancerListenerOptions)`, func() {

			certificateInstanceIdentityModel := &vpcv1.CertificateInstanceIdentityByCRN{
				CRN: core.StringPtr("crn:v1:bluemix:public:cloudcerts:us-south:a/123456:b8866ea4-b8df-467e-801a-da1db7e020bf:certificate:78ff9c4c97d013fb2a95b21dddde7758"),
			}

			loadBalancerPoolIdentityModel := &vpcv1.LoadBalancerPoolIdentityByID{
				ID: core.StringPtr("70294e14-4e61-11e8-bcf4-0242ac110004"),
			}

			loadBalancerListenerPolicyRulePrototypeModel := &vpcv1.LoadBalancerListenerPolicyRulePrototype{
				Condition: core.StringPtr("contains"),
				Field:     core.StringPtr("MY-APP-HEADER"),
				Type:      core.StringPtr("header"),
				Value:     core.StringPtr("testString"),
			}

			loadBalancerListenerPolicyTargetPrototypeModel := &vpcv1.LoadBalancerListenerPolicyTargetPrototypeLoadBalancerPoolIdentityLoadBalancerPoolIdentityByID{
				ID: core.StringPtr("70294e14-4e61-11e8-bcf4-0242ac110004"),
			}

			loadBalancerListenerPolicyPrototypeModel := &vpcv1.LoadBalancerListenerPolicyPrototype{
				Action:   core.StringPtr("forward"),
				Name:     core.StringPtr("my-policy"),
				Priority: core.Int64Ptr(int64(5)),
				Rules:    []vpcv1.LoadBalancerListenerPolicyRulePrototype{*loadBalancerListenerPolicyRulePrototypeModel},
				Target:   loadBalancerListenerPolicyTargetPrototypeModel,
			}

			createLoadBalancerListenerOptions := &vpcv1.CreateLoadBalancerListenerOptions{
				LoadBalancerID:      core.StringPtr("testString"),
				Port:                core.Int64Ptr(int64(443)),
				Protocol:            core.StringPtr("http"),
				AcceptProxyProtocol: core.BoolPtr(true),
				CertificateInstance: certificateInstanceIdentityModel,
				ConnectionLimit:     core.Int64Ptr(int64(2000)),
				DefaultPool:         loadBalancerPoolIdentityModel,
				Policies:            []vpcv1.LoadBalancerListenerPolicyPrototype{*loadBalancerListenerPolicyPrototypeModel},
			}

			loadBalancerListener, response, err := vpcService.CreateLoadBalancerListener(createLoadBalancerListenerOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(loadBalancerListener).ToNot(BeNil())

		})
	})

	Describe(`GetLoadBalancerListener - Retrieve a load balancer listener`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetLoadBalancerListener(getLoadBalancerListenerOptions *GetLoadBalancerListenerOptions)`, func() {

			getLoadBalancerListenerOptions := &vpcv1.GetLoadBalancerListenerOptions{
				LoadBalancerID: core.StringPtr("testString"),
				ID:             core.StringPtr("testString"),
			}

			loadBalancerListener, response, err := vpcService.GetLoadBalancerListener(getLoadBalancerListenerOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerListener).ToNot(BeNil())

		})
	})

	Describe(`UpdateLoadBalancerListener - Update a load balancer listener`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateLoadBalancerListener(updateLoadBalancerListenerOptions *UpdateLoadBalancerListenerOptions)`, func() {

			certificateInstanceIdentityModel := &vpcv1.CertificateInstanceIdentityByCRN{
				CRN: core.StringPtr("crn:v1:bluemix:public:cloudcerts:us-south:a/123456:b8866ea4-b8df-467e-801a-da1db7e020bf:certificate:78ff9c4c97d013fb2a95b21dddde7758"),
			}

			loadBalancerPoolIdentityModel := &vpcv1.LoadBalancerPoolIdentityByID{
				ID: core.StringPtr("70294e14-4e61-11e8-bcf4-0242ac110004"),
			}

			loadBalancerListenerPatchModel := &vpcv1.LoadBalancerListenerPatch{
				AcceptProxyProtocol: core.BoolPtr(true),
				CertificateInstance: certificateInstanceIdentityModel,
				ConnectionLimit:     core.Int64Ptr(int64(2000)),
				DefaultPool:         loadBalancerPoolIdentityModel,
				Port:                core.Int64Ptr(int64(443)),
				Protocol:            core.StringPtr("http"),
			}
			loadBalancerListenerPatchModelAsPatch, asPatchErr := loadBalancerListenerPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateLoadBalancerListenerOptions := &vpcv1.UpdateLoadBalancerListenerOptions{
				LoadBalancerID:            core.StringPtr("testString"),
				ID:                        core.StringPtr("testString"),
				LoadBalancerListenerPatch: loadBalancerListenerPatchModelAsPatch,
			}

			loadBalancerListener, response, err := vpcService.UpdateLoadBalancerListener(updateLoadBalancerListenerOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerListener).ToNot(BeNil())

		})
	})

	Describe(`ListLoadBalancerListenerPolicies - List all policies for a load balancer listener`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListLoadBalancerListenerPolicies(listLoadBalancerListenerPoliciesOptions *ListLoadBalancerListenerPoliciesOptions)`, func() {

			listLoadBalancerListenerPoliciesOptions := &vpcv1.ListLoadBalancerListenerPoliciesOptions{
				LoadBalancerID: core.StringPtr("testString"),
				ListenerID:     core.StringPtr("testString"),
			}

			loadBalancerListenerPolicyCollection, response, err := vpcService.ListLoadBalancerListenerPolicies(listLoadBalancerListenerPoliciesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerListenerPolicyCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateLoadBalancerListenerPolicy - Create a policy for a load balancer listener`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateLoadBalancerListenerPolicy(createLoadBalancerListenerPolicyOptions *CreateLoadBalancerListenerPolicyOptions)`, func() {

			loadBalancerListenerPolicyRulePrototypeModel := &vpcv1.LoadBalancerListenerPolicyRulePrototype{
				Condition: core.StringPtr("contains"),
				Field:     core.StringPtr("MY-APP-HEADER"),
				Type:      core.StringPtr("header"),
				Value:     core.StringPtr("testString"),
			}

			loadBalancerListenerPolicyTargetPrototypeModel := &vpcv1.LoadBalancerListenerPolicyTargetPrototypeLoadBalancerPoolIdentityLoadBalancerPoolIdentityByID{
				ID: core.StringPtr("70294e14-4e61-11e8-bcf4-0242ac110004"),
			}

			createLoadBalancerListenerPolicyOptions := &vpcv1.CreateLoadBalancerListenerPolicyOptions{
				LoadBalancerID: core.StringPtr("testString"),
				ListenerID:     core.StringPtr("testString"),
				Action:         core.StringPtr("forward"),
				Priority:       core.Int64Ptr(int64(5)),
				Name:           core.StringPtr("my-policy"),
				Rules:          []vpcv1.LoadBalancerListenerPolicyRulePrototype{*loadBalancerListenerPolicyRulePrototypeModel},
				Target:         loadBalancerListenerPolicyTargetPrototypeModel,
			}

			loadBalancerListenerPolicy, response, err := vpcService.CreateLoadBalancerListenerPolicy(createLoadBalancerListenerPolicyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(loadBalancerListenerPolicy).ToNot(BeNil())

		})
	})

	Describe(`GetLoadBalancerListenerPolicy - Retrieve a load balancer listener policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetLoadBalancerListenerPolicy(getLoadBalancerListenerPolicyOptions *GetLoadBalancerListenerPolicyOptions)`, func() {

			getLoadBalancerListenerPolicyOptions := &vpcv1.GetLoadBalancerListenerPolicyOptions{
				LoadBalancerID: core.StringPtr("testString"),
				ListenerID:     core.StringPtr("testString"),
				ID:             core.StringPtr("testString"),
			}

			loadBalancerListenerPolicy, response, err := vpcService.GetLoadBalancerListenerPolicy(getLoadBalancerListenerPolicyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerListenerPolicy).ToNot(BeNil())

		})
	})

	Describe(`UpdateLoadBalancerListenerPolicy - Update a load balancer listener policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateLoadBalancerListenerPolicy(updateLoadBalancerListenerPolicyOptions *UpdateLoadBalancerListenerPolicyOptions)`, func() {

			loadBalancerListenerPolicyTargetPatchModel := &vpcv1.LoadBalancerListenerPolicyTargetPatchLoadBalancerPoolIdentityLoadBalancerPoolIdentityByID{
				ID: core.StringPtr("70294e14-4e61-11e8-bcf4-0242ac110004"),
			}

			loadBalancerListenerPolicyPatchModel := &vpcv1.LoadBalancerListenerPolicyPatch{
				Name:     core.StringPtr("my-policy"),
				Priority: core.Int64Ptr(int64(5)),
				Target:   loadBalancerListenerPolicyTargetPatchModel,
			}
			loadBalancerListenerPolicyPatchModelAsPatch, asPatchErr := loadBalancerListenerPolicyPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateLoadBalancerListenerPolicyOptions := &vpcv1.UpdateLoadBalancerListenerPolicyOptions{
				LoadBalancerID:                  core.StringPtr("testString"),
				ListenerID:                      core.StringPtr("testString"),
				ID:                              core.StringPtr("testString"),
				LoadBalancerListenerPolicyPatch: loadBalancerListenerPolicyPatchModelAsPatch,
			}

			loadBalancerListenerPolicy, response, err := vpcService.UpdateLoadBalancerListenerPolicy(updateLoadBalancerListenerPolicyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerListenerPolicy).ToNot(BeNil())

		})
	})

	Describe(`ListLoadBalancerListenerPolicyRules - List all rules of a load balancer listener policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListLoadBalancerListenerPolicyRules(listLoadBalancerListenerPolicyRulesOptions *ListLoadBalancerListenerPolicyRulesOptions)`, func() {

			listLoadBalancerListenerPolicyRulesOptions := &vpcv1.ListLoadBalancerListenerPolicyRulesOptions{
				LoadBalancerID: core.StringPtr("testString"),
				ListenerID:     core.StringPtr("testString"),
				PolicyID:       core.StringPtr("testString"),
			}

			loadBalancerListenerPolicyRuleCollection, response, err := vpcService.ListLoadBalancerListenerPolicyRules(listLoadBalancerListenerPolicyRulesOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerListenerPolicyRuleCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateLoadBalancerListenerPolicyRule - Create a rule for a load balancer listener policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateLoadBalancerListenerPolicyRule(createLoadBalancerListenerPolicyRuleOptions *CreateLoadBalancerListenerPolicyRuleOptions)`, func() {

			createLoadBalancerListenerPolicyRuleOptions := &vpcv1.CreateLoadBalancerListenerPolicyRuleOptions{
				LoadBalancerID: core.StringPtr("testString"),
				ListenerID:     core.StringPtr("testString"),
				PolicyID:       core.StringPtr("testString"),
				Condition:      core.StringPtr("contains"),
				Type:           core.StringPtr("header"),
				Value:          core.StringPtr("testString"),
				Field:          core.StringPtr("MY-APP-HEADER"),
			}

			loadBalancerListenerPolicyRule, response, err := vpcService.CreateLoadBalancerListenerPolicyRule(createLoadBalancerListenerPolicyRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(loadBalancerListenerPolicyRule).ToNot(BeNil())

		})
	})

	Describe(`GetLoadBalancerListenerPolicyRule - Retrieve a load balancer listener policy rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetLoadBalancerListenerPolicyRule(getLoadBalancerListenerPolicyRuleOptions *GetLoadBalancerListenerPolicyRuleOptions)`, func() {

			getLoadBalancerListenerPolicyRuleOptions := &vpcv1.GetLoadBalancerListenerPolicyRuleOptions{
				LoadBalancerID: core.StringPtr("testString"),
				ListenerID:     core.StringPtr("testString"),
				PolicyID:       core.StringPtr("testString"),
				ID:             core.StringPtr("testString"),
			}

			loadBalancerListenerPolicyRule, response, err := vpcService.GetLoadBalancerListenerPolicyRule(getLoadBalancerListenerPolicyRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerListenerPolicyRule).ToNot(BeNil())

		})
	})

	Describe(`UpdateLoadBalancerListenerPolicyRule - Update a load balancer listener policy rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateLoadBalancerListenerPolicyRule(updateLoadBalancerListenerPolicyRuleOptions *UpdateLoadBalancerListenerPolicyRuleOptions)`, func() {

			loadBalancerListenerPolicyRulePatchModel := &vpcv1.LoadBalancerListenerPolicyRulePatch{
				Condition: core.StringPtr("contains"),
				Field:     core.StringPtr("MY-APP-HEADER"),
				Type:      core.StringPtr("header"),
				Value:     core.StringPtr("testString"),
			}
			loadBalancerListenerPolicyRulePatchModelAsPatch, asPatchErr := loadBalancerListenerPolicyRulePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateLoadBalancerListenerPolicyRuleOptions := &vpcv1.UpdateLoadBalancerListenerPolicyRuleOptions{
				LoadBalancerID:                      core.StringPtr("testString"),
				ListenerID:                          core.StringPtr("testString"),
				PolicyID:                            core.StringPtr("testString"),
				ID:                                  core.StringPtr("testString"),
				LoadBalancerListenerPolicyRulePatch: loadBalancerListenerPolicyRulePatchModelAsPatch,
			}

			loadBalancerListenerPolicyRule, response, err := vpcService.UpdateLoadBalancerListenerPolicyRule(updateLoadBalancerListenerPolicyRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerListenerPolicyRule).ToNot(BeNil())

		})
	})

	Describe(`ListLoadBalancerPools - List all pools of a load balancer`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListLoadBalancerPools(listLoadBalancerPoolsOptions *ListLoadBalancerPoolsOptions)`, func() {

			listLoadBalancerPoolsOptions := &vpcv1.ListLoadBalancerPoolsOptions{
				LoadBalancerID: core.StringPtr("testString"),
			}

			loadBalancerPoolCollection, response, err := vpcService.ListLoadBalancerPools(listLoadBalancerPoolsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerPoolCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateLoadBalancerPool - Create a load balancer pool`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateLoadBalancerPool(createLoadBalancerPoolOptions *CreateLoadBalancerPoolOptions)`, func() {

			loadBalancerPoolHealthMonitorPrototypeModel := &vpcv1.LoadBalancerPoolHealthMonitorPrototype{
				Delay:      core.Int64Ptr(int64(5)),
				MaxRetries: core.Int64Ptr(int64(2)),
				Port:       core.Int64Ptr(int64(22)),
				Timeout:    core.Int64Ptr(int64(2)),
				Type:       core.StringPtr("http"),
				URLPath:    core.StringPtr("/"),
			}

			loadBalancerPoolMemberTargetPrototypeModel := &vpcv1.LoadBalancerPoolMemberTargetPrototypeInstanceIdentityInstanceIdentityByID{
				ID: core.StringPtr("1e09281b-f177-46fb-baf1-bc152b2e391a"),
			}

			loadBalancerPoolMemberPrototypeModel := &vpcv1.LoadBalancerPoolMemberPrototype{
				Port:   core.Int64Ptr(int64(80)),
				Target: loadBalancerPoolMemberTargetPrototypeModel,
				Weight: core.Int64Ptr(int64(50)),
			}

			loadBalancerPoolSessionPersistencePrototypeModel := &vpcv1.LoadBalancerPoolSessionPersistencePrototype{
				Type: core.StringPtr("source_ip"),
			}

			createLoadBalancerPoolOptions := &vpcv1.CreateLoadBalancerPoolOptions{
				LoadBalancerID:     core.StringPtr("testString"),
				Algorithm:          core.StringPtr("least_connections"),
				HealthMonitor:      loadBalancerPoolHealthMonitorPrototypeModel,
				Protocol:           core.StringPtr("http"),
				Members:            []vpcv1.LoadBalancerPoolMemberPrototype{*loadBalancerPoolMemberPrototypeModel},
				Name:               core.StringPtr("my-load-balancer-pool"),
				ProxyProtocol:      core.StringPtr("disabled"),
				SessionPersistence: loadBalancerPoolSessionPersistencePrototypeModel,
			}

			loadBalancerPool, response, err := vpcService.CreateLoadBalancerPool(createLoadBalancerPoolOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(loadBalancerPool).ToNot(BeNil())

		})
	})

	Describe(`GetLoadBalancerPool - Retrieve a load balancer pool`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetLoadBalancerPool(getLoadBalancerPoolOptions *GetLoadBalancerPoolOptions)`, func() {

			getLoadBalancerPoolOptions := &vpcv1.GetLoadBalancerPoolOptions{
				LoadBalancerID: core.StringPtr("testString"),
				ID:             core.StringPtr("testString"),
			}

			loadBalancerPool, response, err := vpcService.GetLoadBalancerPool(getLoadBalancerPoolOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerPool).ToNot(BeNil())

		})
	})

	Describe(`UpdateLoadBalancerPool - Update a load balancer pool`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateLoadBalancerPool(updateLoadBalancerPoolOptions *UpdateLoadBalancerPoolOptions)`, func() {

			loadBalancerPoolHealthMonitorPatchModel := &vpcv1.LoadBalancerPoolHealthMonitorPatch{
				Delay:      core.Int64Ptr(int64(5)),
				MaxRetries: core.Int64Ptr(int64(2)),
				Port:       core.Int64Ptr(int64(22)),
				Timeout:    core.Int64Ptr(int64(2)),
				Type:       core.StringPtr("http"),
				URLPath:    core.StringPtr("/"),
			}

			loadBalancerPoolSessionPersistencePatchModel := &vpcv1.LoadBalancerPoolSessionPersistencePatch{
				Type: core.StringPtr("source_ip"),
			}

			loadBalancerPoolPatchModel := &vpcv1.LoadBalancerPoolPatch{
				Algorithm:          core.StringPtr("least_connections"),
				HealthMonitor:      loadBalancerPoolHealthMonitorPatchModel,
				Name:               core.StringPtr("my-load-balancer-pool"),
				Protocol:           core.StringPtr("http"),
				ProxyProtocol:      core.StringPtr("disabled"),
				SessionPersistence: loadBalancerPoolSessionPersistencePatchModel,
			}
			loadBalancerPoolPatchModelAsPatch, asPatchErr := loadBalancerPoolPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateLoadBalancerPoolOptions := &vpcv1.UpdateLoadBalancerPoolOptions{
				LoadBalancerID:        core.StringPtr("testString"),
				ID:                    core.StringPtr("testString"),
				LoadBalancerPoolPatch: loadBalancerPoolPatchModelAsPatch,
			}

			loadBalancerPool, response, err := vpcService.UpdateLoadBalancerPool(updateLoadBalancerPoolOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerPool).ToNot(BeNil())

		})
	})

	Describe(`ListLoadBalancerPoolMembers - List all members of a load balancer pool`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListLoadBalancerPoolMembers(listLoadBalancerPoolMembersOptions *ListLoadBalancerPoolMembersOptions)`, func() {

			listLoadBalancerPoolMembersOptions := &vpcv1.ListLoadBalancerPoolMembersOptions{
				LoadBalancerID: core.StringPtr("testString"),
				PoolID:         core.StringPtr("testString"),
			}

			loadBalancerPoolMemberCollection, response, err := vpcService.ListLoadBalancerPoolMembers(listLoadBalancerPoolMembersOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerPoolMemberCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateLoadBalancerPoolMember - Create a member in a load balancer pool`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateLoadBalancerPoolMember(createLoadBalancerPoolMemberOptions *CreateLoadBalancerPoolMemberOptions)`, func() {

			loadBalancerPoolMemberTargetPrototypeModel := &vpcv1.LoadBalancerPoolMemberTargetPrototypeInstanceIdentityInstanceIdentityByID{
				ID: core.StringPtr("1e09281b-f177-46fb-baf1-bc152b2e391a"),
			}

			createLoadBalancerPoolMemberOptions := &vpcv1.CreateLoadBalancerPoolMemberOptions{
				LoadBalancerID: core.StringPtr("testString"),
				PoolID:         core.StringPtr("testString"),
				Port:           core.Int64Ptr(int64(80)),
				Target:         loadBalancerPoolMemberTargetPrototypeModel,
				Weight:         core.Int64Ptr(int64(50)),
			}

			loadBalancerPoolMember, response, err := vpcService.CreateLoadBalancerPoolMember(createLoadBalancerPoolMemberOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(loadBalancerPoolMember).ToNot(BeNil())

		})
	})

	Describe(`ReplaceLoadBalancerPoolMembers - Replace load balancer pool members`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceLoadBalancerPoolMembers(replaceLoadBalancerPoolMembersOptions *ReplaceLoadBalancerPoolMembersOptions)`, func() {

			loadBalancerPoolMemberTargetPrototypeModel := &vpcv1.LoadBalancerPoolMemberTargetPrototypeInstanceIdentityInstanceIdentityByID{
				ID: core.StringPtr("1e09281b-f177-46fb-baf1-bc152b2e391a"),
			}

			loadBalancerPoolMemberPrototypeModel := &vpcv1.LoadBalancerPoolMemberPrototype{
				Port:   core.Int64Ptr(int64(80)),
				Target: loadBalancerPoolMemberTargetPrototypeModel,
				Weight: core.Int64Ptr(int64(50)),
			}

			replaceLoadBalancerPoolMembersOptions := &vpcv1.ReplaceLoadBalancerPoolMembersOptions{
				LoadBalancerID: core.StringPtr("testString"),
				PoolID:         core.StringPtr("testString"),
				Members:        []vpcv1.LoadBalancerPoolMemberPrototype{*loadBalancerPoolMemberPrototypeModel},
			}

			loadBalancerPoolMemberCollection, response, err := vpcService.ReplaceLoadBalancerPoolMembers(replaceLoadBalancerPoolMembersOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(loadBalancerPoolMemberCollection).ToNot(BeNil())

		})
	})

	Describe(`GetLoadBalancerPoolMember - Retrieve a load balancer pool member`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetLoadBalancerPoolMember(getLoadBalancerPoolMemberOptions *GetLoadBalancerPoolMemberOptions)`, func() {

			getLoadBalancerPoolMemberOptions := &vpcv1.GetLoadBalancerPoolMemberOptions{
				LoadBalancerID: core.StringPtr("testString"),
				PoolID:         core.StringPtr("testString"),
				ID:             core.StringPtr("testString"),
			}

			loadBalancerPoolMember, response, err := vpcService.GetLoadBalancerPoolMember(getLoadBalancerPoolMemberOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerPoolMember).ToNot(BeNil())

		})
	})

	Describe(`UpdateLoadBalancerPoolMember - Update a load balancer pool member`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateLoadBalancerPoolMember(updateLoadBalancerPoolMemberOptions *UpdateLoadBalancerPoolMemberOptions)`, func() {

			loadBalancerPoolMemberTargetPrototypeModel := &vpcv1.LoadBalancerPoolMemberTargetPrototypeInstanceIdentityInstanceIdentityByID{
				ID: core.StringPtr("1e09281b-f177-46fb-baf1-bc152b2e391a"),
			}

			loadBalancerPoolMemberPatchModel := &vpcv1.LoadBalancerPoolMemberPatch{
				Port:   core.Int64Ptr(int64(80)),
				Target: loadBalancerPoolMemberTargetPrototypeModel,
				Weight: core.Int64Ptr(int64(50)),
			}
			loadBalancerPoolMemberPatchModelAsPatch, asPatchErr := loadBalancerPoolMemberPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateLoadBalancerPoolMemberOptions := &vpcv1.UpdateLoadBalancerPoolMemberOptions{
				LoadBalancerID:              core.StringPtr("testString"),
				PoolID:                      core.StringPtr("testString"),
				ID:                          core.StringPtr("testString"),
				LoadBalancerPoolMemberPatch: loadBalancerPoolMemberPatchModelAsPatch,
			}

			loadBalancerPoolMember, response, err := vpcService.UpdateLoadBalancerPoolMember(updateLoadBalancerPoolMemberOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerPoolMember).ToNot(BeNil())

		})
	})

	Describe(`ListEndpointGateways - List all endpoint gateways`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListEndpointGateways(listEndpointGatewaysOptions *ListEndpointGatewaysOptions)`, func() {

			listEndpointGatewaysOptions := &vpcv1.ListEndpointGatewaysOptions{
				Name:            core.StringPtr("testString"),
				Start:           core.StringPtr("testString"),
				Limit:           core.Int64Ptr(int64(1)),
				ResourceGroupID: core.StringPtr("testString"),
			}

			endpointGatewayCollection, response, err := vpcService.ListEndpointGateways(listEndpointGatewaysOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(endpointGatewayCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateEndpointGateway - Create an endpoint gateway`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateEndpointGateway(createEndpointGatewayOptions *CreateEndpointGatewayOptions)`, func() {

			endpointGatewayTargetPrototypeModel := &vpcv1.EndpointGatewayTargetPrototypeProviderCloudServiceIdentityProviderCloudServiceIdentityByCRN{
				ResourceType: core.StringPtr("provider_infrastructure_service"),
				CRN:          core.StringPtr("crn:v1:bluemix:public:cloudant:us-south:a/123456:3527280b-9327-4411-8020-591092e60353::"),
			}

			vpcIdentityModel := &vpcv1.VPCIdentityByID{
				ID: core.StringPtr("f025b503-ae66-46de-a011-3bd08fd5f7bf"),
			}

			endpointGatewayReservedIPModel := &vpcv1.EndpointGatewayReservedIPReservedIPIdentityReservedIPIdentityByID{
				ID: core.StringPtr("6d353a0f-aeb1-4ae1-832e-1110d10981bb"),
			}

			resourceGroupIdentityModel := &vpcv1.ResourceGroupIdentityByID{
				ID: core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345"),
			}

			createEndpointGatewayOptions := &vpcv1.CreateEndpointGatewayOptions{
				Target:        endpointGatewayTargetPrototypeModel,
				VPC:           vpcIdentityModel,
				Ips:           []vpcv1.EndpointGatewayReservedIPIntf{endpointGatewayReservedIPModel},
				Name:          core.StringPtr("testString"),
				ResourceGroup: resourceGroupIdentityModel,
			}

			endpointGateway, response, err := vpcService.CreateEndpointGateway(createEndpointGatewayOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(endpointGateway).ToNot(BeNil())

		})
	})

	Describe(`ListEndpointGatewayIps - List all reserved IPs bound to an endpoint gateway`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListEndpointGatewayIps(listEndpointGatewayIpsOptions *ListEndpointGatewayIpsOptions)`, func() {

			listEndpointGatewayIpsOptions := &vpcv1.ListEndpointGatewayIpsOptions{
				EndpointGatewayID: core.StringPtr("testString"),
				Start:             core.StringPtr("testString"),
				Limit:             core.Int64Ptr(int64(1)),
				Sort:              core.StringPtr("name"),
			}

			reservedIPCollectionEndpointGatewayContext, response, err := vpcService.ListEndpointGatewayIps(listEndpointGatewayIpsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reservedIPCollectionEndpointGatewayContext).ToNot(BeNil())

		})
	})

	Describe(`GetEndpointGatewayIP - Retrieve a reserved IP bound to an endpoint gateway`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetEndpointGatewayIP(getEndpointGatewayIPOptions *GetEndpointGatewayIPOptions)`, func() {

			getEndpointGatewayIPOptions := &vpcv1.GetEndpointGatewayIPOptions{
				EndpointGatewayID: core.StringPtr("testString"),
				ID:                core.StringPtr("testString"),
			}

			reservedIP, response, err := vpcService.GetEndpointGatewayIP(getEndpointGatewayIPOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reservedIP).ToNot(BeNil())

		})
	})

	Describe(`AddEndpointGatewayIP - Bind a reserved IP to an endpoint gateway`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AddEndpointGatewayIP(addEndpointGatewayIPOptions *AddEndpointGatewayIPOptions)`, func() {

			addEndpointGatewayIPOptions := &vpcv1.AddEndpointGatewayIPOptions{
				EndpointGatewayID: core.StringPtr("testString"),
				ID:                core.StringPtr("testString"),
			}

			reservedIP, response, err := vpcService.AddEndpointGatewayIP(addEndpointGatewayIPOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(reservedIP).ToNot(BeNil())

		})
	})

	Describe(`GetEndpointGateway - Retrieve an endpoint gateway`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetEndpointGateway(getEndpointGatewayOptions *GetEndpointGatewayOptions)`, func() {

			getEndpointGatewayOptions := &vpcv1.GetEndpointGatewayOptions{
				ID: core.StringPtr("testString"),
			}

			endpointGateway, response, err := vpcService.GetEndpointGateway(getEndpointGatewayOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(endpointGateway).ToNot(BeNil())

		})
	})

	Describe(`UpdateEndpointGateway - Update an endpoint gateway`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateEndpointGateway(updateEndpointGatewayOptions *UpdateEndpointGatewayOptions)`, func() {

			endpointGatewayPatchModel := &vpcv1.EndpointGatewayPatch{
				Name: core.StringPtr("my-endpoint-gateway"),
			}
			endpointGatewayPatchModelAsPatch, asPatchErr := endpointGatewayPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateEndpointGatewayOptions := &vpcv1.UpdateEndpointGatewayOptions{
				ID:                   core.StringPtr("testString"),
				EndpointGatewayPatch: endpointGatewayPatchModelAsPatch,
			}

			endpointGateway, response, err := vpcService.UpdateEndpointGateway(updateEndpointGatewayOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(endpointGateway).ToNot(BeNil())

		})
	})

	Describe(`ListFlowLogCollectors - List all flow log collectors`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListFlowLogCollectors(listFlowLogCollectorsOptions *ListFlowLogCollectorsOptions)`, func() {

			listFlowLogCollectorsOptions := &vpcv1.ListFlowLogCollectorsOptions{
				Start:              core.StringPtr("testString"),
				Limit:              core.Int64Ptr(int64(1)),
				ResourceGroupID:    core.StringPtr("testString"),
				Name:               core.StringPtr("testString"),
				VPCID:              core.StringPtr("testString"),
				VPCCRN:             core.StringPtr("testString"),
				VPCName:            core.StringPtr("testString"),
				TargetID:           core.StringPtr("testString"),
				TargetResourceType: core.StringPtr("vpc"),
			}

			flowLogCollectorCollection, response, err := vpcService.ListFlowLogCollectors(listFlowLogCollectorsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(flowLogCollectorCollection).ToNot(BeNil())

		})
	})

	Describe(`CreateFlowLogCollector - Create a flow log collector`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateFlowLogCollector(createFlowLogCollectorOptions *CreateFlowLogCollectorOptions)`, func() {

			cloudObjectStorageBucketIdentityModel := &vpcv1.CloudObjectStorageBucketIdentityByName{
				Name: core.StringPtr("bucket-27200-lwx4cfvcue"),
			}

			flowLogCollectorTargetPrototypeModel := &vpcv1.FlowLogCollectorTargetPrototypeNetworkInterfaceIdentityNetworkInterfaceIdentityNetworkInterfaceIdentityByID{
				ID: core.StringPtr("10c02d81-0ecb-4dc5-897d-28392913b81e"),
			}

			resourceGroupIdentityModel := &vpcv1.ResourceGroupIdentityByID{
				ID: core.StringPtr("fee82deba12e4c0fb69c3b09d1f12345"),
			}

			createFlowLogCollectorOptions := &vpcv1.CreateFlowLogCollectorOptions{
				StorageBucket: cloudObjectStorageBucketIdentityModel,
				Target:        flowLogCollectorTargetPrototypeModel,
				Active:        core.BoolPtr(false),
				Name:          core.StringPtr("my-flow-log-collector"),
				ResourceGroup: resourceGroupIdentityModel,
			}

			flowLogCollector, response, err := vpcService.CreateFlowLogCollector(createFlowLogCollectorOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(flowLogCollector).ToNot(BeNil())

		})
	})

	Describe(`GetFlowLogCollector - Retrieve a flow log collector`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetFlowLogCollector(getFlowLogCollectorOptions *GetFlowLogCollectorOptions)`, func() {

			getFlowLogCollectorOptions := &vpcv1.GetFlowLogCollectorOptions{
				ID: core.StringPtr("testString"),
			}

			flowLogCollector, response, err := vpcService.GetFlowLogCollector(getFlowLogCollectorOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(flowLogCollector).ToNot(BeNil())

		})
	})

	Describe(`UpdateFlowLogCollector - Update a flow log collector`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateFlowLogCollector(updateFlowLogCollectorOptions *UpdateFlowLogCollectorOptions)`, func() {

			flowLogCollectorPatchModel := &vpcv1.FlowLogCollectorPatch{
				Active: core.BoolPtr(true),
				Name:   core.StringPtr("my-flow-log-collector"),
			}
			flowLogCollectorPatchModelAsPatch, asPatchErr := flowLogCollectorPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateFlowLogCollectorOptions := &vpcv1.UpdateFlowLogCollectorOptions{
				ID:                    core.StringPtr("testString"),
				FlowLogCollectorPatch: flowLogCollectorPatchModelAsPatch,
			}

			flowLogCollector, response, err := vpcService.UpdateFlowLogCollector(updateFlowLogCollectorOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(flowLogCollector).ToNot(BeNil())

		})
	})

	Describe(`UnsetSubnetPublicGateway - Detach a public gateway from a subnet`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UnsetSubnetPublicGateway(unsetSubnetPublicGatewayOptions *UnsetSubnetPublicGatewayOptions)`, func() {

			unsetSubnetPublicGatewayOptions := &vpcv1.UnsetSubnetPublicGatewayOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := vpcService.UnsetSubnetPublicGateway(unsetSubnetPublicGatewayOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`RemoveVPNGatewayConnectionPeerCIDR - Remove a peer CIDR from a VPN gateway connection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RemoveVPNGatewayConnectionPeerCIDR(removeVPNGatewayConnectionPeerCIDROptions *RemoveVPNGatewayConnectionPeerCIDROptions)`, func() {

			removeVPNGatewayConnectionPeerCIDROptions := &vpcv1.RemoveVPNGatewayConnectionPeerCIDROptions{
				VPNGatewayID: core.StringPtr("testString"),
				ID:           core.StringPtr("testString"),
				CIDRPrefix:   core.StringPtr("testString"),
				PrefixLength: core.StringPtr("testString"),
			}

			response, err := vpcService.RemoveVPNGatewayConnectionPeerCIDR(removeVPNGatewayConnectionPeerCIDROptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`RemoveVPNGatewayConnectionLocalCIDR - Remove a local CIDR from a VPN gateway connection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RemoveVPNGatewayConnectionLocalCIDR(removeVPNGatewayConnectionLocalCIDROptions *RemoveVPNGatewayConnectionLocalCIDROptions)`, func() {

			removeVPNGatewayConnectionLocalCIDROptions := &vpcv1.RemoveVPNGatewayConnectionLocalCIDROptions{
				VPNGatewayID: core.StringPtr("testString"),
				ID:           core.StringPtr("testString"),
				CIDRPrefix:   core.StringPtr("testString"),
				PrefixLength: core.StringPtr("testString"),
			}

			response, err := vpcService.RemoveVPNGatewayConnectionLocalCIDR(removeVPNGatewayConnectionLocalCIDROptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`RemoveVPNGatewayAdvertisedCIDR - Remove an advertised CIDR from a VPN gateway`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RemoveVPNGatewayAdvertisedCIDR(removeVPNGatewayAdvertisedCIDROptions *RemoveVPNGatewayAdvertisedCIDROptions)`, func() {

			removeVPNGatewayAdvertisedCIDROptions := &vpcv1.RemoveVPNGatewayAdvertisedCIDROptions{
				VPNGatewayID: core.StringPtr("testString"),
				CIDRPrefix:   core.StringPtr("testString"),
				PrefixLength: core.StringPtr("testString"),
			}

			response, err := vpcService.RemoveVPNGatewayAdvertisedCIDR(removeVPNGatewayAdvertisedCIDROptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`RemoveSecurityGroupNetworkInterface - Remove a network interface from a security group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RemoveSecurityGroupNetworkInterface(removeSecurityGroupNetworkInterfaceOptions *RemoveSecurityGroupNetworkInterfaceOptions)`, func() {

			removeSecurityGroupNetworkInterfaceOptions := &vpcv1.RemoveSecurityGroupNetworkInterfaceOptions{
				SecurityGroupID: core.StringPtr("testString"),
				ID:              core.StringPtr("testString"),
			}

			response, err := vpcService.RemoveSecurityGroupNetworkInterface(removeSecurityGroupNetworkInterfaceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`RemoveInstanceNetworkInterfaceFloatingIP - Disassociate a floating IP from a network interface`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RemoveInstanceNetworkInterfaceFloatingIP(removeInstanceNetworkInterfaceFloatingIPOptions *RemoveInstanceNetworkInterfaceFloatingIPOptions)`, func() {

			removeInstanceNetworkInterfaceFloatingIPOptions := &vpcv1.RemoveInstanceNetworkInterfaceFloatingIPOptions{
				InstanceID:         core.StringPtr("testString"),
				NetworkInterfaceID: core.StringPtr("testString"),
				ID:                 core.StringPtr("testString"),
			}

			response, err := vpcService.RemoveInstanceNetworkInterfaceFloatingIP(removeInstanceNetworkInterfaceFloatingIPOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`RemoveEndpointGatewayIP - Unbind a reserved IP from an endpoint gateway`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RemoveEndpointGatewayIP(removeEndpointGatewayIPOptions *RemoveEndpointGatewayIPOptions)`, func() {

			removeEndpointGatewayIPOptions := &vpcv1.RemoveEndpointGatewayIPOptions{
				EndpointGatewayID: core.StringPtr("testString"),
				ID:                core.StringPtr("testString"),
			}

			response, err := vpcService.RemoveEndpointGatewayIP(removeEndpointGatewayIPOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`RemoveBareMetalServerNetworkInterfaceFloatingIP - Disassociate a floating IP from a network interface`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RemoveBareMetalServerNetworkInterfaceFloatingIP(removeBareMetalServerNetworkInterfaceFloatingIPOptions *RemoveBareMetalServerNetworkInterfaceFloatingIPOptions)`, func() {

			removeBareMetalServerNetworkInterfaceFloatingIPOptions := &vpcv1.RemoveBareMetalServerNetworkInterfaceFloatingIPOptions{
				BareMetalServerID:  core.StringPtr("testString"),
				NetworkInterfaceID: core.StringPtr("testString"),
				ID:                 core.StringPtr("testString"),
			}

			response, err := vpcService.RemoveBareMetalServerNetworkInterfaceFloatingIP(removeBareMetalServerNetworkInterfaceFloatingIPOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteVolume - Delete a volume`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteVolume(deleteVolumeOptions *DeleteVolumeOptions)`, func() {

			deleteVolumeOptions := &vpcv1.DeleteVolumeOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteVolume(deleteVolumeOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteVPNGatewayConnection - Delete a VPN gateway connection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteVPNGatewayConnection(deleteVPNGatewayConnectionOptions *DeleteVPNGatewayConnectionOptions)`, func() {

			deleteVPNGatewayConnectionOptions := &vpcv1.DeleteVPNGatewayConnectionOptions{
				VPNGatewayID: core.StringPtr("testString"),
				ID:           core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteVPNGatewayConnection(deleteVPNGatewayConnectionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})
	})

	Describe(`DeleteVPNGateway - Delete a VPN gateway`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteVPNGateway(deleteVPNGatewayOptions *DeleteVPNGatewayOptions)`, func() {

			deleteVPNGatewayOptions := &vpcv1.DeleteVPNGatewayOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteVPNGateway(deleteVPNGatewayOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})
	})

	Describe(`DeleteVPCRoutingTableRoute - Delete a VPC routing table route`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteVPCRoutingTableRoute(deleteVPCRoutingTableRouteOptions *DeleteVPCRoutingTableRouteOptions)`, func() {

			deleteVPCRoutingTableRouteOptions := &vpcv1.DeleteVPCRoutingTableRouteOptions{
				VPCID:          core.StringPtr("testString"),
				RoutingTableID: core.StringPtr("testString"),
				ID:             core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteVPCRoutingTableRoute(deleteVPCRoutingTableRouteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteVPCRoutingTable - Delete a VPC routing table`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteVPCRoutingTable(deleteVPCRoutingTableOptions *DeleteVPCRoutingTableOptions)`, func() {

			deleteVPCRoutingTableOptions := &vpcv1.DeleteVPCRoutingTableOptions{
				VPCID: core.StringPtr("testString"),
				ID:    core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteVPCRoutingTable(deleteVPCRoutingTableOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteVPCRoute - Delete a VPC route`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteVPCRoute(deleteVPCRouteOptions *DeleteVPCRouteOptions)`, func() {

			deleteVPCRouteOptions := &vpcv1.DeleteVPCRouteOptions{
				VPCID: core.StringPtr("testString"),
				ID:    core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteVPCRoute(deleteVPCRouteOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteVPCAddressPrefix - Delete an address prefix`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteVPCAddressPrefix(deleteVPCAddressPrefixOptions *DeleteVPCAddressPrefixOptions)`, func() {

			deleteVPCAddressPrefixOptions := &vpcv1.DeleteVPCAddressPrefixOptions{
				VPCID: core.StringPtr("testString"),
				ID:    core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteVPCAddressPrefix(deleteVPCAddressPrefixOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteVPC - Delete a VPC`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteVPC(deleteVPCOptions *DeleteVPCOptions)`, func() {

			deleteVPCOptions := &vpcv1.DeleteVPCOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteVPC(deleteVPCOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteSubnetReservedIP - Release a reserved IP`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSubnetReservedIP(deleteSubnetReservedIPOptions *DeleteSubnetReservedIPOptions)`, func() {

			deleteSubnetReservedIPOptions := &vpcv1.DeleteSubnetReservedIPOptions{
				SubnetID: core.StringPtr("testString"),
				ID:       core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteSubnetReservedIP(deleteSubnetReservedIPOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteSubnet - Delete a subnet`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSubnet(deleteSubnetOptions *DeleteSubnetOptions)`, func() {

			deleteSubnetOptions := &vpcv1.DeleteSubnetOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteSubnet(deleteSubnetOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteSnapshots - Delete a filtered collection of snapshots`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSnapshots(deleteSnapshotsOptions *DeleteSnapshotsOptions)`, func() {

			deleteSnapshotsOptions := &vpcv1.DeleteSnapshotsOptions{
				SourceVolumeID: core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteSnapshots(deleteSnapshotsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteSnapshot - Delete a snapshot`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSnapshot(deleteSnapshotOptions *DeleteSnapshotOptions)`, func() {

			deleteSnapshotOptions := &vpcv1.DeleteSnapshotOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteSnapshot(deleteSnapshotOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteShareTarget - Delete a share target`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteShareTarget(deleteShareTargetOptions *DeleteShareTargetOptions)`, func() {

			deleteShareTargetOptions := &vpcv1.DeleteShareTargetOptions{
				ShareID: core.StringPtr("testString"),
				ID:      core.StringPtr("testString"),
			}

			shareTarget, response, err := vpcService.DeleteShareTarget(deleteShareTargetOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(shareTarget).ToNot(BeNil())

		})
	})

	Describe(`DeleteShare - Delete a file share`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteShare(deleteShareOptions *DeleteShareOptions)`, func() {

			deleteShareOptions := &vpcv1.DeleteShareOptions{
				ID: core.StringPtr("testString"),
			}

			share, response, err := vpcService.DeleteShare(deleteShareOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(share).ToNot(BeNil())

		})
	})

	Describe(`DeleteSecurityGroupTargetBinding - Remove a target from a security group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSecurityGroupTargetBinding(deleteSecurityGroupTargetBindingOptions *DeleteSecurityGroupTargetBindingOptions)`, func() {

			deleteSecurityGroupTargetBindingOptions := &vpcv1.DeleteSecurityGroupTargetBindingOptions{
				SecurityGroupID: core.StringPtr("testString"),
				ID:              core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteSecurityGroupTargetBinding(deleteSecurityGroupTargetBindingOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteSecurityGroupRule - Delete a security group rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSecurityGroupRule(deleteSecurityGroupRuleOptions *DeleteSecurityGroupRuleOptions)`, func() {

			deleteSecurityGroupRuleOptions := &vpcv1.DeleteSecurityGroupRuleOptions{
				SecurityGroupID: core.StringPtr("testString"),
				ID:              core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteSecurityGroupRule(deleteSecurityGroupRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteSecurityGroup - Delete a security group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteSecurityGroup(deleteSecurityGroupOptions *DeleteSecurityGroupOptions)`, func() {

			deleteSecurityGroupOptions := &vpcv1.DeleteSecurityGroupOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteSecurityGroup(deleteSecurityGroupOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeletePublicGateway - Delete a public gateway`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeletePublicGateway(deletePublicGatewayOptions *DeletePublicGatewayOptions)`, func() {

			deletePublicGatewayOptions := &vpcv1.DeletePublicGatewayOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := vpcService.DeletePublicGateway(deletePublicGatewayOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeletePlacementGroup - Delete a placement group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeletePlacementGroup(deletePlacementGroupOptions *DeletePlacementGroupOptions)`, func() {

			deletePlacementGroupOptions := &vpcv1.DeletePlacementGroupOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := vpcService.DeletePlacementGroup(deletePlacementGroupOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})
	})

	Describe(`DeleteNetworkACLRule - Delete a network ACL rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteNetworkACLRule(deleteNetworkACLRuleOptions *DeleteNetworkACLRuleOptions)`, func() {

			deleteNetworkACLRuleOptions := &vpcv1.DeleteNetworkACLRuleOptions{
				NetworkACLID: core.StringPtr("testString"),
				ID:           core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteNetworkACLRule(deleteNetworkACLRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteNetworkACL - Delete a network ACL`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteNetworkACL(deleteNetworkACLOptions *DeleteNetworkACLOptions)`, func() {

			deleteNetworkACLOptions := &vpcv1.DeleteNetworkACLOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteNetworkACL(deleteNetworkACLOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteLoadBalancerPoolMember - Delete a load balancer pool member`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteLoadBalancerPoolMember(deleteLoadBalancerPoolMemberOptions *DeleteLoadBalancerPoolMemberOptions)`, func() {

			deleteLoadBalancerPoolMemberOptions := &vpcv1.DeleteLoadBalancerPoolMemberOptions{
				LoadBalancerID: core.StringPtr("testString"),
				PoolID:         core.StringPtr("testString"),
				ID:             core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteLoadBalancerPoolMember(deleteLoadBalancerPoolMemberOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteLoadBalancerPool - Delete a load balancer pool`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteLoadBalancerPool(deleteLoadBalancerPoolOptions *DeleteLoadBalancerPoolOptions)`, func() {

			deleteLoadBalancerPoolOptions := &vpcv1.DeleteLoadBalancerPoolOptions{
				LoadBalancerID: core.StringPtr("testString"),
				ID:             core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteLoadBalancerPool(deleteLoadBalancerPoolOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteLoadBalancerListenerPolicyRule - Delete a load balancer listener policy rule`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteLoadBalancerListenerPolicyRule(deleteLoadBalancerListenerPolicyRuleOptions *DeleteLoadBalancerListenerPolicyRuleOptions)`, func() {

			deleteLoadBalancerListenerPolicyRuleOptions := &vpcv1.DeleteLoadBalancerListenerPolicyRuleOptions{
				LoadBalancerID: core.StringPtr("testString"),
				ListenerID:     core.StringPtr("testString"),
				PolicyID:       core.StringPtr("testString"),
				ID:             core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteLoadBalancerListenerPolicyRule(deleteLoadBalancerListenerPolicyRuleOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteLoadBalancerListenerPolicy - Delete a load balancer listener policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteLoadBalancerListenerPolicy(deleteLoadBalancerListenerPolicyOptions *DeleteLoadBalancerListenerPolicyOptions)`, func() {

			deleteLoadBalancerListenerPolicyOptions := &vpcv1.DeleteLoadBalancerListenerPolicyOptions{
				LoadBalancerID: core.StringPtr("testString"),
				ListenerID:     core.StringPtr("testString"),
				ID:             core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteLoadBalancerListenerPolicy(deleteLoadBalancerListenerPolicyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteLoadBalancerListener - Delete a load balancer listener`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteLoadBalancerListener(deleteLoadBalancerListenerOptions *DeleteLoadBalancerListenerOptions)`, func() {

			deleteLoadBalancerListenerOptions := &vpcv1.DeleteLoadBalancerListenerOptions{
				LoadBalancerID: core.StringPtr("testString"),
				ID:             core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteLoadBalancerListener(deleteLoadBalancerListenerOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteLoadBalancer - Delete a load balancer`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteLoadBalancer(deleteLoadBalancerOptions *DeleteLoadBalancerOptions)`, func() {

			deleteLoadBalancerOptions := &vpcv1.DeleteLoadBalancerOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteLoadBalancer(deleteLoadBalancerOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteKey - Delete a key`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteKey(deleteKeyOptions *DeleteKeyOptions)`, func() {

			deleteKeyOptions := &vpcv1.DeleteKeyOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteKey(deleteKeyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})
	})

	Describe(`DeleteIpsecPolicy - Delete an IPsec policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteIpsecPolicy(deleteIpsecPolicyOptions *DeleteIpsecPolicyOptions)`, func() {

			deleteIpsecPolicyOptions := &vpcv1.DeleteIpsecPolicyOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteIpsecPolicy(deleteIpsecPolicyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteInstanceVolumeAttachment - Delete a volume attachment`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteInstanceVolumeAttachment(deleteInstanceVolumeAttachmentOptions *DeleteInstanceVolumeAttachmentOptions)`, func() {

			deleteInstanceVolumeAttachmentOptions := &vpcv1.DeleteInstanceVolumeAttachmentOptions{
				InstanceID: core.StringPtr("testString"),
				ID:         core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteInstanceVolumeAttachment(deleteInstanceVolumeAttachmentOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteInstanceTemplate - Delete an instance template`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteInstanceTemplate(deleteInstanceTemplateOptions *DeleteInstanceTemplateOptions)`, func() {

			deleteInstanceTemplateOptions := &vpcv1.DeleteInstanceTemplateOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteInstanceTemplate(deleteInstanceTemplateOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteInstanceNetworkInterface - Delete a network interface`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteInstanceNetworkInterface(deleteInstanceNetworkInterfaceOptions *DeleteInstanceNetworkInterfaceOptions)`, func() {

			deleteInstanceNetworkInterfaceOptions := &vpcv1.DeleteInstanceNetworkInterfaceOptions{
				InstanceID: core.StringPtr("testString"),
				ID:         core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteInstanceNetworkInterface(deleteInstanceNetworkInterfaceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteInstanceGroupMemberships - Delete all memberships from an instance group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteInstanceGroupMemberships(deleteInstanceGroupMembershipsOptions *DeleteInstanceGroupMembershipsOptions)`, func() {

			deleteInstanceGroupMembershipsOptions := &vpcv1.DeleteInstanceGroupMembershipsOptions{
				InstanceGroupID: core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteInstanceGroupMemberships(deleteInstanceGroupMembershipsOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteInstanceGroupMembership - Delete an instance group membership`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteInstanceGroupMembership(deleteInstanceGroupMembershipOptions *DeleteInstanceGroupMembershipOptions)`, func() {

			deleteInstanceGroupMembershipOptions := &vpcv1.DeleteInstanceGroupMembershipOptions{
				InstanceGroupID: core.StringPtr("testString"),
				ID:              core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteInstanceGroupMembership(deleteInstanceGroupMembershipOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteInstanceGroupManagerPolicy - Delete an instance group manager policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteInstanceGroupManagerPolicy(deleteInstanceGroupManagerPolicyOptions *DeleteInstanceGroupManagerPolicyOptions)`, func() {

			deleteInstanceGroupManagerPolicyOptions := &vpcv1.DeleteInstanceGroupManagerPolicyOptions{
				InstanceGroupID:        core.StringPtr("testString"),
				InstanceGroupManagerID: core.StringPtr("testString"),
				ID:                     core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteInstanceGroupManagerPolicy(deleteInstanceGroupManagerPolicyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteInstanceGroupManagerAction - Delete specified instance group manager action`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteInstanceGroupManagerAction(deleteInstanceGroupManagerActionOptions *DeleteInstanceGroupManagerActionOptions)`, func() {

			deleteInstanceGroupManagerActionOptions := &vpcv1.DeleteInstanceGroupManagerActionOptions{
				InstanceGroupID:        core.StringPtr("testString"),
				InstanceGroupManagerID: core.StringPtr("testString"),
				ID:                     core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteInstanceGroupManagerAction(deleteInstanceGroupManagerActionOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteInstanceGroupManager - Delete an instance group manager`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteInstanceGroupManager(deleteInstanceGroupManagerOptions *DeleteInstanceGroupManagerOptions)`, func() {

			deleteInstanceGroupManagerOptions := &vpcv1.DeleteInstanceGroupManagerOptions{
				InstanceGroupID: core.StringPtr("testString"),
				ID:              core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteInstanceGroupManager(deleteInstanceGroupManagerOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteInstanceGroupLoadBalancer - Delete an instance group load balancer`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteInstanceGroupLoadBalancer(deleteInstanceGroupLoadBalancerOptions *DeleteInstanceGroupLoadBalancerOptions)`, func() {

			deleteInstanceGroupLoadBalancerOptions := &vpcv1.DeleteInstanceGroupLoadBalancerOptions{
				InstanceGroupID: core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteInstanceGroupLoadBalancer(deleteInstanceGroupLoadBalancerOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteInstanceGroup - Delete an instance group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteInstanceGroup(deleteInstanceGroupOptions *DeleteInstanceGroupOptions)`, func() {

			deleteInstanceGroupOptions := &vpcv1.DeleteInstanceGroupOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteInstanceGroup(deleteInstanceGroupOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteInstance - Delete an instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteInstance(deleteInstanceOptions *DeleteInstanceOptions)`, func() {

			deleteInstanceOptions := &vpcv1.DeleteInstanceOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteInstance(deleteInstanceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteImage - Delete an image`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteImage(deleteImageOptions *DeleteImageOptions)`, func() {

			deleteImageOptions := &vpcv1.DeleteImageOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteImage(deleteImageOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})
	})

	Describe(`DeleteIkePolicy - Delete an IKE policy`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteIkePolicy(deleteIkePolicyOptions *DeleteIkePolicyOptions)`, func() {

			deleteIkePolicyOptions := &vpcv1.DeleteIkePolicyOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteIkePolicy(deleteIkePolicyOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteFlowLogCollector - Delete a flow log collector`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteFlowLogCollector(deleteFlowLogCollectorOptions *DeleteFlowLogCollectorOptions)`, func() {

			deleteFlowLogCollectorOptions := &vpcv1.DeleteFlowLogCollectorOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteFlowLogCollector(deleteFlowLogCollectorOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteFloatingIP - Release a floating IP`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteFloatingIP(deleteFloatingIPOptions *DeleteFloatingIPOptions)`, func() {

			deleteFloatingIPOptions := &vpcv1.DeleteFloatingIPOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteFloatingIP(deleteFloatingIPOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteEndpointGateway - Delete an endpoint gateway`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteEndpointGateway(deleteEndpointGatewayOptions *DeleteEndpointGatewayOptions)`, func() {

			deleteEndpointGatewayOptions := &vpcv1.DeleteEndpointGatewayOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteEndpointGateway(deleteEndpointGatewayOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteDedicatedHostGroup - Delete a dedicated host group`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDedicatedHostGroup(deleteDedicatedHostGroupOptions *DeleteDedicatedHostGroupOptions)`, func() {

			deleteDedicatedHostGroupOptions := &vpcv1.DeleteDedicatedHostGroupOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteDedicatedHostGroup(deleteDedicatedHostGroupOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteDedicatedHost - Delete a dedicated host`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDedicatedHost(deleteDedicatedHostOptions *DeleteDedicatedHostOptions)`, func() {

			deleteDedicatedHostOptions := &vpcv1.DeleteDedicatedHostOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteDedicatedHost(deleteDedicatedHostOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteBareMetalServerNetworkInterface - Delete a network interface`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteBareMetalServerNetworkInterface(deleteBareMetalServerNetworkInterfaceOptions *DeleteBareMetalServerNetworkInterfaceOptions)`, func() {

			deleteBareMetalServerNetworkInterfaceOptions := &vpcv1.DeleteBareMetalServerNetworkInterfaceOptions{
				BareMetalServerID: core.StringPtr("testString"),
				ID:                core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteBareMetalServerNetworkInterface(deleteBareMetalServerNetworkInterfaceOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

	Describe(`DeleteBareMetalServer - Delete a bare metal server`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteBareMetalServer(deleteBareMetalServerOptions *DeleteBareMetalServerOptions)`, func() {

			deleteBareMetalServerOptions := &vpcv1.DeleteBareMetalServerOptions{
				ID: core.StringPtr("testString"),
			}

			response, err := vpcService.DeleteBareMetalServer(deleteBareMetalServerOptions)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})

//
// Utility functions are declared in the unit test file
//
