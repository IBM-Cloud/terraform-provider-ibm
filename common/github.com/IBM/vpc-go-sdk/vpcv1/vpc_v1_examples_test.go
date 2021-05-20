// +build examples

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
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the vpc service.
//
// The following configuration properties are assumed to be defined:
// VPC_URL=<service base url>
// VPC_AUTH_TYPE=iam
// VPC_APIKEY=<IAM apikey>
// VPC_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
const externalConfigFile = "../vpc_v1.env"

var (
	vpcService   *vpcv1.VpcV1
	config       map[string]string
	configLoaded bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`VpcV1 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(vpcv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
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

			vpcServiceOptions := &vpcv1.VpcV1Options{
				Version: core.StringPtr("testString"),
			}

			vpcService, err = vpcv1.NewVpcV1UsingExternalConfig(vpcServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(vpcService).ToNot(BeNil())
		})
	})

	Describe(`VpcV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListVpcs request example`, func() {
			fmt.Println("\nListVpcs() result:")
			// begin-list_vpcs

			listVpcsOptions := vpcService.NewListVpcsOptions()
			listVpcsOptions.SetClassicAccess(true)

			vpcCollection, response, err := vpcService.ListVpcs(listVpcsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(vpcCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_vpcs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpcCollection).ToNot(BeNil())

		})
		It(`CreateVPC request example`, func() {
			fmt.Println("\nCreateVPC() result:")
			// begin-create_vpc

			createVPCOptions := vpcService.NewCreateVPCOptions()

			vpc, response, err := vpcService.CreateVPC(createVPCOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(vpc, "", "  ")
			fmt.Println(string(b))

			// end-create_vpc

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(vpc).ToNot(BeNil())

		})
		It(`GetVPC request example`, func() {
			fmt.Println("\nGetVPC() result:")
			// begin-get_vpc

			getVPCOptions := vpcService.NewGetVPCOptions(
				"testString",
			)

			vpc, response, err := vpcService.GetVPC(getVPCOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(vpc, "", "  ")
			fmt.Println(string(b))

			// end-get_vpc

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpc).ToNot(BeNil())

		})
		It(`UpdateVPC request example`, func() {
			fmt.Println("\nUpdateVPC() result:")
			// begin-update_vpc

			vpcPatchModel := &vpcv1.VPCPatch{}
			vpcPatchModelAsPatch, asPatchErr := vpcPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateVPCOptions := vpcService.NewUpdateVPCOptions(
				"testString",
				vpcPatchModelAsPatch,
			)

			vpc, response, err := vpcService.UpdateVPC(updateVPCOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(vpc, "", "  ")
			fmt.Println(string(b))

			// end-update_vpc

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpc).ToNot(BeNil())

		})
		It(`GetVPCDefaultNetworkACL request example`, func() {
			fmt.Println("\nGetVPCDefaultNetworkACL() result:")
			// begin-get_vpc_default_network_acl

			getVPCDefaultNetworkACLOptions := vpcService.NewGetVPCDefaultNetworkACLOptions(
				"testString",
			)

			defaultNetworkACL, response, err := vpcService.GetVPCDefaultNetworkACL(getVPCDefaultNetworkACLOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(defaultNetworkACL, "", "  ")
			fmt.Println(string(b))

			// end-get_vpc_default_network_acl

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(defaultNetworkACL).ToNot(BeNil())

		})
		It(`GetVPCDefaultRoutingTable request example`, func() {
			fmt.Println("\nGetVPCDefaultRoutingTable() result:")
			// begin-get_vpc_default_routing_table

			getVPCDefaultRoutingTableOptions := vpcService.NewGetVPCDefaultRoutingTableOptions(
				"testString",
			)

			defaultRoutingTable, response, err := vpcService.GetVPCDefaultRoutingTable(getVPCDefaultRoutingTableOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(defaultRoutingTable, "", "  ")
			fmt.Println(string(b))

			// end-get_vpc_default_routing_table

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(defaultRoutingTable).ToNot(BeNil())

		})
		It(`GetVPCDefaultSecurityGroup request example`, func() {
			fmt.Println("\nGetVPCDefaultSecurityGroup() result:")
			// begin-get_vpc_default_security_group

			getVPCDefaultSecurityGroupOptions := vpcService.NewGetVPCDefaultSecurityGroupOptions(
				"testString",
			)

			defaultSecurityGroup, response, err := vpcService.GetVPCDefaultSecurityGroup(getVPCDefaultSecurityGroupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(defaultSecurityGroup, "", "  ")
			fmt.Println(string(b))

			// end-get_vpc_default_security_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(defaultSecurityGroup).ToNot(BeNil())

		})
		It(`ListVPCAddressPrefixes request example`, func() {
			fmt.Println("\nListVPCAddressPrefixes() result:")
			// begin-list_vpc_address_prefixes

			listVPCAddressPrefixesOptions := vpcService.NewListVPCAddressPrefixesOptions(
				"testString",
			)

			addressPrefixCollection, response, err := vpcService.ListVPCAddressPrefixes(listVPCAddressPrefixesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(addressPrefixCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_vpc_address_prefixes

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(addressPrefixCollection).ToNot(BeNil())

		})
		It(`CreateVPCAddressPrefix request example`, func() {
			fmt.Println("\nCreateVPCAddressPrefix() result:")
			// begin-create_vpc_address_prefix

			zoneIdentityModel := &vpcv1.ZoneIdentityByName{
				Name: core.StringPtr("us-south-1"),
			}

			createVPCAddressPrefixOptions := vpcService.NewCreateVPCAddressPrefixOptions(
				"testString",
				"10.0.0.0/24",
				zoneIdentityModel,
			)

			addressPrefix, response, err := vpcService.CreateVPCAddressPrefix(createVPCAddressPrefixOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(addressPrefix, "", "  ")
			fmt.Println(string(b))

			// end-create_vpc_address_prefix

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(addressPrefix).ToNot(BeNil())

		})
		It(`GetVPCAddressPrefix request example`, func() {
			fmt.Println("\nGetVPCAddressPrefix() result:")
			// begin-get_vpc_address_prefix

			getVPCAddressPrefixOptions := vpcService.NewGetVPCAddressPrefixOptions(
				"testString",
				"testString",
			)

			addressPrefix, response, err := vpcService.GetVPCAddressPrefix(getVPCAddressPrefixOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(addressPrefix, "", "  ")
			fmt.Println(string(b))

			// end-get_vpc_address_prefix

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(addressPrefix).ToNot(BeNil())

		})
		It(`UpdateVPCAddressPrefix request example`, func() {
			fmt.Println("\nUpdateVPCAddressPrefix() result:")
			// begin-update_vpc_address_prefix

			addressPrefixPatchModel := &vpcv1.AddressPrefixPatch{}
			addressPrefixPatchModelAsPatch, asPatchErr := addressPrefixPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateVPCAddressPrefixOptions := vpcService.NewUpdateVPCAddressPrefixOptions(
				"testString",
				"testString",
				addressPrefixPatchModelAsPatch,
			)

			addressPrefix, response, err := vpcService.UpdateVPCAddressPrefix(updateVPCAddressPrefixOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(addressPrefix, "", "  ")
			fmt.Println(string(b))

			// end-update_vpc_address_prefix

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(addressPrefix).ToNot(BeNil())

		})
		It(`ListVPCRoutes request example`, func() {
			fmt.Println("\nListVPCRoutes() result:")
			// begin-list_vpc_routes

			listVPCRoutesOptions := vpcService.NewListVPCRoutesOptions(
				"testString",
			)

			routeCollection, response, err := vpcService.ListVPCRoutes(listVPCRoutesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(routeCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_vpc_routes

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(routeCollection).ToNot(BeNil())

		})
		It(`CreateVPCRoute request example`, func() {
			fmt.Println("\nCreateVPCRoute() result:")
			// begin-create_vpc_route

			zoneIdentityModel := &vpcv1.ZoneIdentityByName{
				Name: core.StringPtr("us-south-1"),
			}

			createVPCRouteOptions := vpcService.NewCreateVPCRouteOptions(
				"testString",
				"192.168.3.0/24",
				zoneIdentityModel,
			)

			route, response, err := vpcService.CreateVPCRoute(createVPCRouteOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(route, "", "  ")
			fmt.Println(string(b))

			// end-create_vpc_route

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(route).ToNot(BeNil())

		})
		It(`GetVPCRoute request example`, func() {
			fmt.Println("\nGetVPCRoute() result:")
			// begin-get_vpc_route

			getVPCRouteOptions := vpcService.NewGetVPCRouteOptions(
				"testString",
				"testString",
			)

			route, response, err := vpcService.GetVPCRoute(getVPCRouteOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(route, "", "  ")
			fmt.Println(string(b))

			// end-get_vpc_route

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(route).ToNot(BeNil())

		})
		It(`UpdateVPCRoute request example`, func() {
			fmt.Println("\nUpdateVPCRoute() result:")
			// begin-update_vpc_route

			routePatchModel := &vpcv1.RoutePatch{}
			routePatchModelAsPatch, asPatchErr := routePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateVPCRouteOptions := vpcService.NewUpdateVPCRouteOptions(
				"testString",
				"testString",
				routePatchModelAsPatch,
			)

			route, response, err := vpcService.UpdateVPCRoute(updateVPCRouteOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(route, "", "  ")
			fmt.Println(string(b))

			// end-update_vpc_route

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(route).ToNot(BeNil())

		})
		It(`ListVPCRoutingTables request example`, func() {
			fmt.Println("\nListVPCRoutingTables() result:")
			// begin-list_vpc_routing_tables

			listVPCRoutingTablesOptions := vpcService.NewListVPCRoutingTablesOptions(
				"testString",
			)

			routingTableCollection, response, err := vpcService.ListVPCRoutingTables(listVPCRoutingTablesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(routingTableCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_vpc_routing_tables

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(routingTableCollection).ToNot(BeNil())

		})
		It(`CreateVPCRoutingTable request example`, func() {
			fmt.Println("\nCreateVPCRoutingTable() result:")
			// begin-create_vpc_routing_table

			createVPCRoutingTableOptions := vpcService.NewCreateVPCRoutingTableOptions(
				"testString",
			)

			routingTable, response, err := vpcService.CreateVPCRoutingTable(createVPCRoutingTableOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(routingTable, "", "  ")
			fmt.Println(string(b))

			// end-create_vpc_routing_table

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(routingTable).ToNot(BeNil())

		})
		It(`GetVPCRoutingTable request example`, func() {
			fmt.Println("\nGetVPCRoutingTable() result:")
			// begin-get_vpc_routing_table

			getVPCRoutingTableOptions := vpcService.NewGetVPCRoutingTableOptions(
				"testString",
				"testString",
			)

			routingTable, response, err := vpcService.GetVPCRoutingTable(getVPCRoutingTableOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(routingTable, "", "  ")
			fmt.Println(string(b))

			// end-get_vpc_routing_table

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(routingTable).ToNot(BeNil())

		})
		It(`UpdateVPCRoutingTable request example`, func() {
			fmt.Println("\nUpdateVPCRoutingTable() result:")
			// begin-update_vpc_routing_table

			routingTablePatchModel := &vpcv1.RoutingTablePatch{}
			routingTablePatchModelAsPatch, asPatchErr := routingTablePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateVPCRoutingTableOptions := vpcService.NewUpdateVPCRoutingTableOptions(
				"testString",
				"testString",
				routingTablePatchModelAsPatch,
			)

			routingTable, response, err := vpcService.UpdateVPCRoutingTable(updateVPCRoutingTableOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(routingTable, "", "  ")
			fmt.Println(string(b))

			// end-update_vpc_routing_table

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(routingTable).ToNot(BeNil())

		})
		It(`ListVPCRoutingTableRoutes request example`, func() {
			fmt.Println("\nListVPCRoutingTableRoutes() result:")
			// begin-list_vpc_routing_table_routes

			listVPCRoutingTableRoutesOptions := vpcService.NewListVPCRoutingTableRoutesOptions(
				"testString",
				"testString",
			)

			routeCollection, response, err := vpcService.ListVPCRoutingTableRoutes(listVPCRoutingTableRoutesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(routeCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_vpc_routing_table_routes

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(routeCollection).ToNot(BeNil())

		})
		It(`CreateVPCRoutingTableRoute request example`, func() {
			fmt.Println("\nCreateVPCRoutingTableRoute() result:")
			// begin-create_vpc_routing_table_route

			zoneIdentityModel := &vpcv1.ZoneIdentityByName{
				Name: core.StringPtr("us-south-1"),
			}

			createVPCRoutingTableRouteOptions := vpcService.NewCreateVPCRoutingTableRouteOptions(
				"testString",
				"testString",
				"192.168.3.0/24",
				zoneIdentityModel,
			)

			route, response, err := vpcService.CreateVPCRoutingTableRoute(createVPCRoutingTableRouteOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(route, "", "  ")
			fmt.Println(string(b))

			// end-create_vpc_routing_table_route

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(route).ToNot(BeNil())

		})
		It(`GetVPCRoutingTableRoute request example`, func() {
			fmt.Println("\nGetVPCRoutingTableRoute() result:")
			// begin-get_vpc_routing_table_route

			getVPCRoutingTableRouteOptions := vpcService.NewGetVPCRoutingTableRouteOptions(
				"testString",
				"testString",
				"testString",
			)

			route, response, err := vpcService.GetVPCRoutingTableRoute(getVPCRoutingTableRouteOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(route, "", "  ")
			fmt.Println(string(b))

			// end-get_vpc_routing_table_route

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(route).ToNot(BeNil())

		})
		It(`UpdateVPCRoutingTableRoute request example`, func() {
			fmt.Println("\nUpdateVPCRoutingTableRoute() result:")
			// begin-update_vpc_routing_table_route

			routePatchModel := &vpcv1.RoutePatch{}
			routePatchModelAsPatch, asPatchErr := routePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateVPCRoutingTableRouteOptions := vpcService.NewUpdateVPCRoutingTableRouteOptions(
				"testString",
				"testString",
				"testString",
				routePatchModelAsPatch,
			)

			route, response, err := vpcService.UpdateVPCRoutingTableRoute(updateVPCRoutingTableRouteOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(route, "", "  ")
			fmt.Println(string(b))

			// end-update_vpc_routing_table_route

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(route).ToNot(BeNil())

		})
		It(`ListSubnets request example`, func() {
			fmt.Println("\nListSubnets() result:")
			// begin-list_subnets

			listSubnetsOptions := vpcService.NewListSubnetsOptions()

			subnetCollection, response, err := vpcService.ListSubnets(listSubnetsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(subnetCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_subnets

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subnetCollection).ToNot(BeNil())

		})
		It(`CreateSubnet request example`, func() {
			fmt.Println("\nCreateSubnet() result:")
			// begin-create_subnet

			vpcIdentityModel := &vpcv1.VPCIdentityByID{
				ID: core.StringPtr("4727d842-f94f-4a2d-824a-9bc9b02c523b"),
			}

			zoneIdentityModel := &vpcv1.ZoneIdentityByName{
				Name: core.StringPtr("us-south-1"),
			}

			subnetPrototypeModel := &vpcv1.SubnetPrototypeSubnetByTotalCount{
				VPC:                   vpcIdentityModel,
				TotalIpv4AddressCount: core.Int64Ptr(int64(256)),
				Zone:                  zoneIdentityModel,
			}

			createSubnetOptions := vpcService.NewCreateSubnetOptions(
				subnetPrototypeModel,
			)

			subnet, response, err := vpcService.CreateSubnet(createSubnetOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(subnet, "", "  ")
			fmt.Println(string(b))

			// end-create_subnet

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subnet).ToNot(BeNil())

		})
		It(`GetSubnet request example`, func() {
			fmt.Println("\nGetSubnet() result:")
			// begin-get_subnet

			getSubnetOptions := vpcService.NewGetSubnetOptions(
				"testString",
			)

			subnet, response, err := vpcService.GetSubnet(getSubnetOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(subnet, "", "  ")
			fmt.Println(string(b))

			// end-get_subnet

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subnet).ToNot(BeNil())

		})
		It(`UpdateSubnet request example`, func() {
			fmt.Println("\nUpdateSubnet() result:")
			// begin-update_subnet

			subnetPatchModel := &vpcv1.SubnetPatch{}
			subnetPatchModelAsPatch, asPatchErr := subnetPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateSubnetOptions := vpcService.NewUpdateSubnetOptions(
				"testString",
				subnetPatchModelAsPatch,
			)

			subnet, response, err := vpcService.UpdateSubnet(updateSubnetOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(subnet, "", "  ")
			fmt.Println(string(b))

			// end-update_subnet

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subnet).ToNot(BeNil())

		})
		It(`GetSubnetNetworkACL request example`, func() {
			fmt.Println("\nGetSubnetNetworkACL() result:")
			// begin-get_subnet_network_acl

			getSubnetNetworkACLOptions := vpcService.NewGetSubnetNetworkACLOptions(
				"testString",
			)

			networkACL, response, err := vpcService.GetSubnetNetworkACL(getSubnetNetworkACLOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(networkACL, "", "  ")
			fmt.Println(string(b))

			// end-get_subnet_network_acl

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkACL).ToNot(BeNil())

		})
		It(`ReplaceSubnetNetworkACL request example`, func() {
			fmt.Println("\nReplaceSubnetNetworkACL() result:")
			// begin-replace_subnet_network_acl

			networkACLIdentityModel := &vpcv1.NetworkACLIdentityByID{
				ID: core.StringPtr("a4e28308-8ee7-46ab-8108-9f881f22bdbf"),
			}

			replaceSubnetNetworkACLOptions := vpcService.NewReplaceSubnetNetworkACLOptions(
				"testString",
				networkACLIdentityModel,
			)

			networkACL, response, err := vpcService.ReplaceSubnetNetworkACL(replaceSubnetNetworkACLOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(networkACL, "", "  ")
			fmt.Println(string(b))

			// end-replace_subnet_network_acl

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(networkACL).ToNot(BeNil())

		})
		It(`GetSubnetPublicGateway request example`, func() {
			fmt.Println("\nGetSubnetPublicGateway() result:")
			// begin-get_subnet_public_gateway

			getSubnetPublicGatewayOptions := vpcService.NewGetSubnetPublicGatewayOptions(
				"testString",
			)

			publicGateway, response, err := vpcService.GetSubnetPublicGateway(getSubnetPublicGatewayOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(publicGateway, "", "  ")
			fmt.Println(string(b))

			// end-get_subnet_public_gateway

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(publicGateway).ToNot(BeNil())

		})
		It(`SetSubnetPublicGateway request example`, func() {
			fmt.Println("\nSetSubnetPublicGateway() result:")
			// begin-set_subnet_public_gateway

			publicGatewayIdentityModel := &vpcv1.PublicGatewayIdentityByID{
				ID: core.StringPtr("dc5431ef-1fc6-4861-adc9-a59d077d1241"),
			}

			setSubnetPublicGatewayOptions := vpcService.NewSetSubnetPublicGatewayOptions(
				"testString",
				publicGatewayIdentityModel,
			)

			publicGateway, response, err := vpcService.SetSubnetPublicGateway(setSubnetPublicGatewayOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(publicGateway, "", "  ")
			fmt.Println(string(b))

			// end-set_subnet_public_gateway

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(publicGateway).ToNot(BeNil())

		})
		It(`GetSubnetRoutingTable request example`, func() {
			fmt.Println("\nGetSubnetRoutingTable() result:")
			// begin-get_subnet_routing_table

			getSubnetRoutingTableOptions := vpcService.NewGetSubnetRoutingTableOptions(
				"testString",
			)

			routingTable, response, err := vpcService.GetSubnetRoutingTable(getSubnetRoutingTableOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(routingTable, "", "  ")
			fmt.Println(string(b))

			// end-get_subnet_routing_table

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(routingTable).ToNot(BeNil())

		})
		It(`ReplaceSubnetRoutingTable request example`, func() {
			fmt.Println("\nReplaceSubnetRoutingTable() result:")
			// begin-replace_subnet_routing_table

			routingTableIdentityModel := &vpcv1.RoutingTableIdentityByID{
				ID: core.StringPtr("1a15dca5-7e33-45e1-b7c5-bc690e569531"),
			}

			replaceSubnetRoutingTableOptions := vpcService.NewReplaceSubnetRoutingTableOptions(
				"testString",
				routingTableIdentityModel,
			)

			routingTable, response, err := vpcService.ReplaceSubnetRoutingTable(replaceSubnetRoutingTableOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(routingTable, "", "  ")
			fmt.Println(string(b))

			// end-replace_subnet_routing_table

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(routingTable).ToNot(BeNil())

		})
		It(`ListSubnetReservedIps request example`, func() {
			fmt.Println("\nListSubnetReservedIps() result:")
			// begin-list_subnet_reserved_ips

			listSubnetReservedIpsOptions := vpcService.NewListSubnetReservedIpsOptions(
				"testString",
			)
			listSubnetReservedIpsOptions.SetSort("name")

			reservedIPCollection, response, err := vpcService.ListSubnetReservedIps(listSubnetReservedIpsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(reservedIPCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_subnet_reserved_ips

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reservedIPCollection).ToNot(BeNil())

		})
		It(`CreateSubnetReservedIP request example`, func() {
			fmt.Println("\nCreateSubnetReservedIP() result:")
			// begin-create_subnet_reserved_ip

			createSubnetReservedIPOptions := vpcService.NewCreateSubnetReservedIPOptions(
				"testString",
			)

			reservedIP, response, err := vpcService.CreateSubnetReservedIP(createSubnetReservedIPOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(reservedIP, "", "  ")
			fmt.Println(string(b))

			// end-create_subnet_reserved_ip

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(reservedIP).ToNot(BeNil())

		})
		It(`GetSubnetReservedIP request example`, func() {
			fmt.Println("\nGetSubnetReservedIP() result:")
			// begin-get_subnet_reserved_ip

			getSubnetReservedIPOptions := vpcService.NewGetSubnetReservedIPOptions(
				"testString",
				"testString",
			)

			reservedIP, response, err := vpcService.GetSubnetReservedIP(getSubnetReservedIPOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(reservedIP, "", "  ")
			fmt.Println(string(b))

			// end-get_subnet_reserved_ip

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reservedIP).ToNot(BeNil())

		})
		It(`UpdateSubnetReservedIP request example`, func() {
			fmt.Println("\nUpdateSubnetReservedIP() result:")
			// begin-update_subnet_reserved_ip

			reservedIPPatchModel := &vpcv1.ReservedIPPatch{}
			reservedIPPatchModelAsPatch, asPatchErr := reservedIPPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateSubnetReservedIPOptions := vpcService.NewUpdateSubnetReservedIPOptions(
				"testString",
				"testString",
				reservedIPPatchModelAsPatch,
			)

			reservedIP, response, err := vpcService.UpdateSubnetReservedIP(updateSubnetReservedIPOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(reservedIP, "", "  ")
			fmt.Println(string(b))

			// end-update_subnet_reserved_ip

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reservedIP).ToNot(BeNil())

		})
		It(`ListImages request example`, func() {
			fmt.Println("\nListImages() result:")
			// begin-list_images

			listImagesOptions := vpcService.NewListImagesOptions()

			imageCollection, response, err := vpcService.ListImages(listImagesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(imageCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_images

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(imageCollection).ToNot(BeNil())

		})
		It(`CreateImage request example`, func() {
			fmt.Println("\nCreateImage() result:")
			// begin-create_image

			imageFilePrototypeModel := &vpcv1.ImageFilePrototype{
				Href: core.StringPtr("cos://us-south/my-bucket/my-image.qcow2"),
			}

			operatingSystemIdentityModel := &vpcv1.OperatingSystemIdentityByName{
				Name: core.StringPtr("debian-9-amd64"),
			}

			imagePrototypeModel := &vpcv1.ImagePrototypeImageByFile{
				File:            imageFilePrototypeModel,
				OperatingSystem: operatingSystemIdentityModel,
			}

			createImageOptions := vpcService.NewCreateImageOptions(
				imagePrototypeModel,
			)

			image, response, err := vpcService.CreateImage(createImageOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(image, "", "  ")
			fmt.Println(string(b))

			// end-create_image

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(image).ToNot(BeNil())

		})
		It(`GetImage request example`, func() {
			fmt.Println("\nGetImage() result:")
			// begin-get_image

			getImageOptions := vpcService.NewGetImageOptions(
				"testString",
			)

			image, response, err := vpcService.GetImage(getImageOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(image, "", "  ")
			fmt.Println(string(b))

			// end-get_image

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(image).ToNot(BeNil())

		})
		It(`UpdateImage request example`, func() {
			fmt.Println("\nUpdateImage() result:")
			// begin-update_image

			imagePatchModel := &vpcv1.ImagePatch{}
			imagePatchModelAsPatch, asPatchErr := imagePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateImageOptions := vpcService.NewUpdateImageOptions(
				"testString",
				imagePatchModelAsPatch,
			)

			image, response, err := vpcService.UpdateImage(updateImageOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(image, "", "  ")
			fmt.Println(string(b))

			// end-update_image

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(image).ToNot(BeNil())

		})
		It(`ListOperatingSystems request example`, func() {
			fmt.Println("\nListOperatingSystems() result:")
			// begin-list_operating_systems

			listOperatingSystemsOptions := vpcService.NewListOperatingSystemsOptions()

			operatingSystemCollection, response, err := vpcService.ListOperatingSystems(listOperatingSystemsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(operatingSystemCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_operating_systems

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(operatingSystemCollection).ToNot(BeNil())

		})
		It(`GetOperatingSystem request example`, func() {
			fmt.Println("\nGetOperatingSystem() result:")
			// begin-get_operating_system

			getOperatingSystemOptions := vpcService.NewGetOperatingSystemOptions(
				"testString",
			)

			operatingSystem, response, err := vpcService.GetOperatingSystem(getOperatingSystemOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(operatingSystem, "", "  ")
			fmt.Println(string(b))

			// end-get_operating_system

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(operatingSystem).ToNot(BeNil())

		})
		It(`ListKeys request example`, func() {
			fmt.Println("\nListKeys() result:")
			// begin-list_keys

			listKeysOptions := vpcService.NewListKeysOptions()

			keyCollection, response, err := vpcService.ListKeys(listKeysOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(keyCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_keys

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(keyCollection).ToNot(BeNil())

		})
		It(`CreateKey request example`, func() {
			fmt.Println("\nCreateKey() result:")
			// begin-create_key

			createKeyOptions := vpcService.NewCreateKeyOptions(
				"AAAAB3NzaC1yc2EAAAADAQABAAABAQDDGe50Bxa5T5NDddrrtbx2Y4/VGbiCgXqnBsYToIUKoFSHTQl5IX3PasGnneKanhcLwWz5M5MoCRvhxTp66NKzIfAz7r+FX9rxgR+ZgcM253YAqOVeIpOU408simDZKriTlN8kYsXL7P34tsWuAJf4MgZtJAQxous/2byetpdCv8ddnT4X3ltOg9w+LqSCPYfNivqH00Eh7S1Ldz7I8aw5WOp5a+sQFP/RbwfpwHp+ny7DfeIOokcuI42tJkoBn7UsLTVpCSmXr2EDRlSWe/1M/iHNRBzaT3CK0+SwZWd2AEjePxSnWKNGIEUJDlUYp7hKhiQcgT5ZAnWU121oc5En",
			)

			key, response, err := vpcService.CreateKey(createKeyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(key, "", "  ")
			fmt.Println(string(b))

			// end-create_key

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(key).ToNot(BeNil())

		})
		It(`GetKey request example`, func() {
			fmt.Println("\nGetKey() result:")
			// begin-get_key

			getKeyOptions := vpcService.NewGetKeyOptions(
				"testString",
			)

			key, response, err := vpcService.GetKey(getKeyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(key, "", "  ")
			fmt.Println(string(b))

			// end-get_key

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(key).ToNot(BeNil())

		})
		It(`UpdateKey request example`, func() {
			fmt.Println("\nUpdateKey() result:")
			// begin-update_key

			keyPatchModel := &vpcv1.KeyPatch{}
			keyPatchModelAsPatch, asPatchErr := keyPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateKeyOptions := vpcService.NewUpdateKeyOptions(
				"testString",
				keyPatchModelAsPatch,
			)

			key, response, err := vpcService.UpdateKey(updateKeyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(key, "", "  ")
			fmt.Println(string(b))

			// end-update_key

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(key).ToNot(BeNil())

		})
		It(`ListInstanceProfiles request example`, func() {
			fmt.Println("\nListInstanceProfiles() result:")
			// begin-list_instance_profiles

			listInstanceProfilesOptions := vpcService.NewListInstanceProfilesOptions()

			instanceProfileCollection, response, err := vpcService.ListInstanceProfiles(listInstanceProfilesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceProfileCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_instance_profiles

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceProfileCollection).ToNot(BeNil())

		})
		It(`GetInstanceProfile request example`, func() {
			fmt.Println("\nGetInstanceProfile() result:")
			// begin-get_instance_profile

			getInstanceProfileOptions := vpcService.NewGetInstanceProfileOptions(
				"testString",
			)

			instanceProfile, response, err := vpcService.GetInstanceProfile(getInstanceProfileOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceProfile, "", "  ")
			fmt.Println(string(b))

			// end-get_instance_profile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceProfile).ToNot(BeNil())

		})
		It(`ListInstanceTemplates request example`, func() {
			fmt.Println("\nListInstanceTemplates() result:")
			// begin-list_instance_templates

			listInstanceTemplatesOptions := vpcService.NewListInstanceTemplatesOptions()

			instanceTemplateCollection, response, err := vpcService.ListInstanceTemplates(listInstanceTemplatesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceTemplateCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_instance_templates

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceTemplateCollection).ToNot(BeNil())

		})
		It(`CreateInstanceTemplate request example`, func() {
			fmt.Println("\nCreateInstanceTemplate() result:")
			// begin-create_instance_template

			keyIdentityModel := &vpcv1.KeyIdentityByID{
				ID: core.StringPtr("363f6d70-0000-0001-0000-00000013b96c"),
			}

			instanceProfileIdentityModel := &vpcv1.InstanceProfileIdentityByName{
				Name: core.StringPtr("bx2-2x8"),
			}

			vpcIdentityModel := &vpcv1.VPCIdentityByID{
				ID: core.StringPtr("dc201ab2-8536-4904-86a8-084d84582133"),
			}

			imageIdentityModel := &vpcv1.ImageIdentityByID{
				ID: core.StringPtr("3f9a2d96-830e-4100-9b4c-663225a3f872"),
			}

			subnetIdentityModel := &vpcv1.SubnetIdentityByID{
				ID: core.StringPtr("0d933c75-492a-4756-9832-1200585dfa79"),
			}

			networkInterfacePrototypeModel := &vpcv1.NetworkInterfacePrototype{
				Subnet: subnetIdentityModel,
			}

			zoneIdentityModel := &vpcv1.ZoneIdentityByName{
				Name: core.StringPtr("us-south-1"),
			}

			instanceTemplatePrototypeModel := &vpcv1.InstanceTemplatePrototypeInstanceByImage{
				Keys:                    []vpcv1.KeyIdentityIntf{keyIdentityModel},
				Name:                    core.StringPtr("my-instance-template"),
				Profile:                 instanceProfileIdentityModel,
				VPC:                     vpcIdentityModel,
				Image:                   imageIdentityModel,
				PrimaryNetworkInterface: networkInterfacePrototypeModel,
				Zone:                    zoneIdentityModel,
			}

			createInstanceTemplateOptions := vpcService.NewCreateInstanceTemplateOptions(
				instanceTemplatePrototypeModel,
			)

			instanceTemplate, response, err := vpcService.CreateInstanceTemplate(createInstanceTemplateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceTemplate, "", "  ")
			fmt.Println(string(b))

			// end-create_instance_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(instanceTemplate).ToNot(BeNil())

		})
		It(`GetInstanceTemplate request example`, func() {
			fmt.Println("\nGetInstanceTemplate() result:")
			// begin-get_instance_template

			getInstanceTemplateOptions := vpcService.NewGetInstanceTemplateOptions(
				"testString",
			)

			instanceTemplate, response, err := vpcService.GetInstanceTemplate(getInstanceTemplateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceTemplate, "", "  ")
			fmt.Println(string(b))

			// end-get_instance_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceTemplate).ToNot(BeNil())

		})
		It(`UpdateInstanceTemplate request example`, func() {
			fmt.Println("\nUpdateInstanceTemplate() result:")
			// begin-update_instance_template

			instanceTemplatePatchModel := &vpcv1.InstanceTemplatePatch{}
			instanceTemplatePatchModelAsPatch, asPatchErr := instanceTemplatePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateInstanceTemplateOptions := vpcService.NewUpdateInstanceTemplateOptions(
				"testString",
				instanceTemplatePatchModelAsPatch,
			)

			instanceTemplate, response, err := vpcService.UpdateInstanceTemplate(updateInstanceTemplateOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceTemplate, "", "  ")
			fmt.Println(string(b))

			// end-update_instance_template

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceTemplate).ToNot(BeNil())

		})
		It(`ListInstances request example`, func() {
			fmt.Println("\nListInstances() result:")
			// begin-list_instances

			listInstancesOptions := vpcService.NewListInstancesOptions()

			instanceCollection, response, err := vpcService.ListInstances(listInstancesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_instances

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceCollection).ToNot(BeNil())

		})
		It(`CreateInstance request example`, func() {
			fmt.Println("\nCreateInstance() result:")
			// begin-create_instance

			keyIdentityModel := &vpcv1.KeyIdentityByID{
				ID: core.StringPtr("363f6d70-0000-0001-0000-00000013b96c"),
			}

			instancePlacementTargetPrototypeModel := &vpcv1.InstancePlacementTargetPrototypeDedicatedHostIdentityDedicatedHostIdentityByID{
				ID: core.StringPtr("0787-8c2a09be-ee18-4af2-8ef4-6a6060732221"),
			}

			instanceProfileIdentityModel := &vpcv1.InstanceProfileIdentityByName{
				Name: core.StringPtr("bx2-2x8"),
			}

			volumeProfileIdentityModel := &vpcv1.VolumeProfileIdentityByName{
				Name: core.StringPtr("5iops-tier"),
			}

			encryptionKeyIdentityModel := &vpcv1.EncryptionKeyIdentityByCRN{
				CRN: core.StringPtr("crn:[...]"),
			}

			volumeAttachmentVolumePrototypeInstanceContextModel := &vpcv1.VolumeAttachmentVolumePrototypeInstanceContextVolumePrototypeInstanceContextVolumePrototypeInstanceContextVolumeByCapacity{
				Name:          core.StringPtr("my-data-volume"),
				Profile:       volumeProfileIdentityModel,
				Capacity:      core.Int64Ptr(int64(1000)),
				EncryptionKey: encryptionKeyIdentityModel,
			}

			volumeAttachmentPrototypeInstanceContextModel := &vpcv1.VolumeAttachmentPrototypeInstanceContext{
				Volume: volumeAttachmentVolumePrototypeInstanceContextModel,
			}

			vpcIdentityModel := &vpcv1.VPCIdentityByID{
				ID: core.StringPtr("f0aae929-7047-46d1-92e1-9102b07a7f6f"),
			}

			volumePrototypeInstanceByImageContextModel := &vpcv1.VolumePrototypeInstanceByImageContext{
				EncryptionKey: encryptionKeyIdentityModel,
				Name:          core.StringPtr("my-boot-volume"),
				Profile:       volumeProfileIdentityModel,
			}

			volumeAttachmentPrototypeInstanceByImageContextModel := &vpcv1.VolumeAttachmentPrototypeInstanceByImageContext{
				Volume: volumePrototypeInstanceByImageContextModel,
			}

			imageIdentityModel := &vpcv1.ImageIdentityByID{
				ID: core.StringPtr("9aaf3bcb-dcd7-4de7-bb60-24e39ff9d366"),
			}

			subnetIdentityModel := &vpcv1.SubnetIdentityByID{
				ID: core.StringPtr("bea6a632-5e13-42a4-b4b8-31dc877abfe4"),
			}

			networkInterfacePrototypeModel := &vpcv1.NetworkInterfacePrototype{
				Name:   core.StringPtr("my-network-interface"),
				Subnet: subnetIdentityModel,
			}

			zoneIdentityModel := &vpcv1.ZoneIdentityByName{
				Name: core.StringPtr("us-south-1"),
			}

			instancePrototypeModel := &vpcv1.InstancePrototypeInstanceByImage{
				Keys:                    []vpcv1.KeyIdentityIntf{keyIdentityModel},
				Name:                    core.StringPtr("my-instance"),
				PlacementTarget:         instancePlacementTargetPrototypeModel,
				Profile:                 instanceProfileIdentityModel,
				VolumeAttachments:       []vpcv1.VolumeAttachmentPrototypeInstanceContext{*volumeAttachmentPrototypeInstanceContextModel},
				VPC:                     vpcIdentityModel,
				BootVolumeAttachment:    volumeAttachmentPrototypeInstanceByImageContextModel,
				Image:                   imageIdentityModel,
				PrimaryNetworkInterface: networkInterfacePrototypeModel,
				Zone:                    zoneIdentityModel,
			}

			createInstanceOptions := vpcService.NewCreateInstanceOptions(
				instancePrototypeModel,
			)

			instance, response, err := vpcService.CreateInstance(createInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instance, "", "  ")
			fmt.Println(string(b))

			// end-create_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(instance).ToNot(BeNil())

		})
		It(`GetInstance request example`, func() {
			fmt.Println("\nGetInstance() result:")
			// begin-get_instance

			getInstanceOptions := vpcService.NewGetInstanceOptions(
				"testString",
			)

			instance, response, err := vpcService.GetInstance(getInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instance, "", "  ")
			fmt.Println(string(b))

			// end-get_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instance).ToNot(BeNil())

		})
		It(`UpdateInstance request example`, func() {
			fmt.Println("\nUpdateInstance() result:")
			// begin-update_instance

			instancePatchModel := &vpcv1.InstancePatch{}
			instancePatchModelAsPatch, asPatchErr := instancePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateInstanceOptions := vpcService.NewUpdateInstanceOptions(
				"testString",
				instancePatchModelAsPatch,
			)

			instance, response, err := vpcService.UpdateInstance(updateInstanceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instance, "", "  ")
			fmt.Println(string(b))

			// end-update_instance

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instance).ToNot(BeNil())

		})
		It(`GetInstanceInitialization request example`, func() {
			fmt.Println("\nGetInstanceInitialization() result:")
			// begin-get_instance_initialization

			getInstanceInitializationOptions := vpcService.NewGetInstanceInitializationOptions(
				"testString",
			)

			instanceInitialization, response, err := vpcService.GetInstanceInitialization(getInstanceInitializationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceInitialization, "", "  ")
			fmt.Println(string(b))

			// end-get_instance_initialization

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceInitialization).ToNot(BeNil())

		})
		It(`CreateInstanceAction request example`, func() {
			fmt.Println("\nCreateInstanceAction() result:")
			// begin-create_instance_action

			createInstanceActionOptions := vpcService.NewCreateInstanceActionOptions(
				"testString",
				"reboot",
			)

			instanceAction, response, err := vpcService.CreateInstanceAction(createInstanceActionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceAction, "", "  ")
			fmt.Println(string(b))

			// end-create_instance_action

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(instanceAction).ToNot(BeNil())

		})
		It(`CreateInstanceConsoleAccessToken request example`, func() {
			fmt.Println("\nCreateInstanceConsoleAccessToken() result:")
			// begin-create_instance_console_access_token

			createInstanceConsoleAccessTokenOptions := vpcService.NewCreateInstanceConsoleAccessTokenOptions(
				"testString",
				"serial",
			)

			instanceConsoleAccessToken, response, err := vpcService.CreateInstanceConsoleAccessToken(createInstanceConsoleAccessTokenOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceConsoleAccessToken, "", "  ")
			fmt.Println(string(b))

			// end-create_instance_console_access_token

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceConsoleAccessToken).ToNot(BeNil())

		})
		It(`ListInstanceDisks request example`, func() {
			fmt.Println("\nListInstanceDisks() result:")
			// begin-list_instance_disks

			listInstanceDisksOptions := vpcService.NewListInstanceDisksOptions(
				"testString",
			)

			instanceDiskCollection, response, err := vpcService.ListInstanceDisks(listInstanceDisksOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceDiskCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_instance_disks

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceDiskCollection).ToNot(BeNil())

		})
		It(`GetInstanceDisk request example`, func() {
			fmt.Println("\nGetInstanceDisk() result:")
			// begin-get_instance_disk

			getInstanceDiskOptions := vpcService.NewGetInstanceDiskOptions(
				"testString",
				"testString",
			)

			instanceDisk, response, err := vpcService.GetInstanceDisk(getInstanceDiskOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceDisk, "", "  ")
			fmt.Println(string(b))

			// end-get_instance_disk

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceDisk).ToNot(BeNil())

		})
		It(`UpdateInstanceDisk request example`, func() {
			fmt.Println("\nUpdateInstanceDisk() result:")
			// begin-update_instance_disk

			instanceDiskPatchModel := &vpcv1.InstanceDiskPatch{}
			instanceDiskPatchModelAsPatch, asPatchErr := instanceDiskPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateInstanceDiskOptions := vpcService.NewUpdateInstanceDiskOptions(
				"testString",
				"testString",
				instanceDiskPatchModelAsPatch,
			)

			instanceDisk, response, err := vpcService.UpdateInstanceDisk(updateInstanceDiskOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceDisk, "", "  ")
			fmt.Println(string(b))

			// end-update_instance_disk

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceDisk).ToNot(BeNil())

		})
		It(`ListInstanceNetworkInterfaces request example`, func() {
			fmt.Println("\nListInstanceNetworkInterfaces() result:")
			// begin-list_instance_network_interfaces

			listInstanceNetworkInterfacesOptions := vpcService.NewListInstanceNetworkInterfacesOptions(
				"testString",
			)

			networkInterfaceUnpaginatedCollection, response, err := vpcService.ListInstanceNetworkInterfaces(listInstanceNetworkInterfacesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(networkInterfaceUnpaginatedCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_instance_network_interfaces

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkInterfaceUnpaginatedCollection).ToNot(BeNil())

		})
		It(`CreateInstanceNetworkInterface request example`, func() {
			fmt.Println("\nCreateInstanceNetworkInterface() result:")
			// begin-create_instance_network_interface

			subnetIdentityModel := &vpcv1.SubnetIdentityByID{
				ID: core.StringPtr("7ec86020-1c6e-4889-b3f0-a15f2e50f87e"),
			}

			createInstanceNetworkInterfaceOptions := vpcService.NewCreateInstanceNetworkInterfaceOptions(
				"testString",
				subnetIdentityModel,
			)

			networkInterface, response, err := vpcService.CreateInstanceNetworkInterface(createInstanceNetworkInterfaceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(networkInterface, "", "  ")
			fmt.Println(string(b))

			// end-create_instance_network_interface

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(networkInterface).ToNot(BeNil())

		})
		It(`GetInstanceNetworkInterface request example`, func() {
			fmt.Println("\nGetInstanceNetworkInterface() result:")
			// begin-get_instance_network_interface

			getInstanceNetworkInterfaceOptions := vpcService.NewGetInstanceNetworkInterfaceOptions(
				"testString",
				"testString",
			)

			networkInterface, response, err := vpcService.GetInstanceNetworkInterface(getInstanceNetworkInterfaceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(networkInterface, "", "  ")
			fmt.Println(string(b))

			// end-get_instance_network_interface

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkInterface).ToNot(BeNil())

		})
		It(`UpdateInstanceNetworkInterface request example`, func() {
			fmt.Println("\nUpdateInstanceNetworkInterface() result:")
			// begin-update_instance_network_interface

			networkInterfacePatchModel := &vpcv1.NetworkInterfacePatch{
				Name: core.StringPtr("my-network-interface-1"),
			}
			networkInterfacePatchModelAsPatch, asPatchErr := networkInterfacePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateInstanceNetworkInterfaceOptions := vpcService.NewUpdateInstanceNetworkInterfaceOptions(
				"testString",
				"testString",
				networkInterfacePatchModelAsPatch,
			)

			networkInterface, response, err := vpcService.UpdateInstanceNetworkInterface(updateInstanceNetworkInterfaceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(networkInterface, "", "  ")
			fmt.Println(string(b))

			// end-update_instance_network_interface

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkInterface).ToNot(BeNil())

		})
		It(`ListInstanceNetworkInterfaceFloatingIps request example`, func() {
			fmt.Println("\nListInstanceNetworkInterfaceFloatingIps() result:")
			// begin-list_instance_network_interface_floating_ips

			listInstanceNetworkInterfaceFloatingIpsOptions := vpcService.NewListInstanceNetworkInterfaceFloatingIpsOptions(
				"testString",
				"testString",
			)

			floatingIPUnpaginatedCollection, response, err := vpcService.ListInstanceNetworkInterfaceFloatingIps(listInstanceNetworkInterfaceFloatingIpsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(floatingIPUnpaginatedCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_instance_network_interface_floating_ips

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(floatingIPUnpaginatedCollection).ToNot(BeNil())

		})
		It(`GetInstanceNetworkInterfaceFloatingIP request example`, func() {
			fmt.Println("\nGetInstanceNetworkInterfaceFloatingIP() result:")
			// begin-get_instance_network_interface_floating_ip

			getInstanceNetworkInterfaceFloatingIPOptions := vpcService.NewGetInstanceNetworkInterfaceFloatingIPOptions(
				"testString",
				"testString",
				"testString",
			)

			floatingIP, response, err := vpcService.GetInstanceNetworkInterfaceFloatingIP(getInstanceNetworkInterfaceFloatingIPOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(floatingIP, "", "  ")
			fmt.Println(string(b))

			// end-get_instance_network_interface_floating_ip

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(floatingIP).ToNot(BeNil())

		})
		It(`AddInstanceNetworkInterfaceFloatingIP request example`, func() {
			fmt.Println("\nAddInstanceNetworkInterfaceFloatingIP() result:")
			// begin-add_instance_network_interface_floating_ip

			addInstanceNetworkInterfaceFloatingIPOptions := vpcService.NewAddInstanceNetworkInterfaceFloatingIPOptions(
				"testString",
				"testString",
				"testString",
			)

			floatingIP, response, err := vpcService.AddInstanceNetworkInterfaceFloatingIP(addInstanceNetworkInterfaceFloatingIPOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(floatingIP, "", "  ")
			fmt.Println(string(b))

			// end-add_instance_network_interface_floating_ip

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(floatingIP).ToNot(BeNil())

		})
		It(`ListInstanceVolumeAttachments request example`, func() {
			fmt.Println("\nListInstanceVolumeAttachments() result:")
			// begin-list_instance_volume_attachments

			listInstanceVolumeAttachmentsOptions := vpcService.NewListInstanceVolumeAttachmentsOptions(
				"testString",
			)

			volumeAttachmentCollection, response, err := vpcService.ListInstanceVolumeAttachments(listInstanceVolumeAttachmentsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(volumeAttachmentCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_instance_volume_attachments

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeAttachmentCollection).ToNot(BeNil())

		})
		It(`CreateInstanceVolumeAttachment request example`, func() {
			fmt.Println("\nCreateInstanceVolumeAttachment() result:")
			// begin-create_instance_volume_attachment

			volumeAttachmentPrototypeVolumeModel := &vpcv1.VolumeAttachmentPrototypeVolumeVolumeIdentityVolumeIdentityByID{
				ID: core.StringPtr("1a6b7274-678d-4dfb-8981-c71dd9d4daa5"),
			}

			createInstanceVolumeAttachmentOptions := vpcService.NewCreateInstanceVolumeAttachmentOptions(
				"testString",
				volumeAttachmentPrototypeVolumeModel,
			)

			volumeAttachment, response, err := vpcService.CreateInstanceVolumeAttachment(createInstanceVolumeAttachmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(volumeAttachment, "", "  ")
			fmt.Println(string(b))

			// end-create_instance_volume_attachment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(volumeAttachment).ToNot(BeNil())

		})
		It(`GetInstanceVolumeAttachment request example`, func() {
			fmt.Println("\nGetInstanceVolumeAttachment() result:")
			// begin-get_instance_volume_attachment

			getInstanceVolumeAttachmentOptions := vpcService.NewGetInstanceVolumeAttachmentOptions(
				"testString",
				"testString",
			)

			volumeAttachment, response, err := vpcService.GetInstanceVolumeAttachment(getInstanceVolumeAttachmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(volumeAttachment, "", "  ")
			fmt.Println(string(b))

			// end-get_instance_volume_attachment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeAttachment).ToNot(BeNil())

		})
		It(`UpdateInstanceVolumeAttachment request example`, func() {
			fmt.Println("\nUpdateInstanceVolumeAttachment() result:")
			// begin-update_instance_volume_attachment

			volumeAttachmentPatchModel := &vpcv1.VolumeAttachmentPatch{}
			volumeAttachmentPatchModelAsPatch, asPatchErr := volumeAttachmentPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateInstanceVolumeAttachmentOptions := vpcService.NewUpdateInstanceVolumeAttachmentOptions(
				"testString",
				"testString",
				volumeAttachmentPatchModelAsPatch,
			)

			volumeAttachment, response, err := vpcService.UpdateInstanceVolumeAttachment(updateInstanceVolumeAttachmentOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(volumeAttachment, "", "  ")
			fmt.Println(string(b))

			// end-update_instance_volume_attachment

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeAttachment).ToNot(BeNil())

		})
		It(`ListInstanceGroups request example`, func() {
			fmt.Println("\nListInstanceGroups() result:")
			// begin-list_instance_groups

			listInstanceGroupsOptions := vpcService.NewListInstanceGroupsOptions()

			instanceGroupCollection, response, err := vpcService.ListInstanceGroups(listInstanceGroupsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceGroupCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_instance_groups

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupCollection).ToNot(BeNil())

		})
		It(`CreateInstanceGroup request example`, func() {
			fmt.Println("\nCreateInstanceGroup() result:")
			// begin-create_instance_group

			instanceTemplateIdentityModel := &vpcv1.InstanceTemplateIdentityByID{
				ID: core.StringPtr("a6b1a881-2ce8-41a3-80fc-36316a73f803"),
			}

			subnetIdentityModel := &vpcv1.SubnetIdentityByID{
				ID: core.StringPtr("7ec86020-1c6e-4889-b3f0-a15f2e50f87e"),
			}

			createInstanceGroupOptions := vpcService.NewCreateInstanceGroupOptions(
				instanceTemplateIdentityModel,
				[]vpcv1.SubnetIdentityIntf{subnetIdentityModel},
			)

			instanceGroup, response, err := vpcService.CreateInstanceGroup(createInstanceGroupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceGroup, "", "  ")
			fmt.Println(string(b))

			// end-create_instance_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(instanceGroup).ToNot(BeNil())

		})
		It(`GetInstanceGroup request example`, func() {
			fmt.Println("\nGetInstanceGroup() result:")
			// begin-get_instance_group

			getInstanceGroupOptions := vpcService.NewGetInstanceGroupOptions(
				"testString",
			)

			instanceGroup, response, err := vpcService.GetInstanceGroup(getInstanceGroupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceGroup, "", "  ")
			fmt.Println(string(b))

			// end-get_instance_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroup).ToNot(BeNil())

		})
		It(`UpdateInstanceGroup request example`, func() {
			fmt.Println("\nUpdateInstanceGroup() result:")
			// begin-update_instance_group

			instanceGroupPatchModel := &vpcv1.InstanceGroupPatch{}
			instanceGroupPatchModelAsPatch, asPatchErr := instanceGroupPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateInstanceGroupOptions := vpcService.NewUpdateInstanceGroupOptions(
				"testString",
				instanceGroupPatchModelAsPatch,
			)

			instanceGroup, response, err := vpcService.UpdateInstanceGroup(updateInstanceGroupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceGroup, "", "  ")
			fmt.Println(string(b))

			// end-update_instance_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroup).ToNot(BeNil())

		})
		It(`ListInstanceGroupManagers request example`, func() {
			fmt.Println("\nListInstanceGroupManagers() result:")
			// begin-list_instance_group_managers

			listInstanceGroupManagersOptions := vpcService.NewListInstanceGroupManagersOptions(
				"testString",
			)

			instanceGroupManagerCollection, response, err := vpcService.ListInstanceGroupManagers(listInstanceGroupManagersOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceGroupManagerCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_instance_group_managers

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManagerCollection).ToNot(BeNil())

		})
		It(`CreateInstanceGroupManager request example`, func() {
			fmt.Println("\nCreateInstanceGroupManager() result:")
			// begin-create_instance_group_manager

			instanceGroupManagerPrototypeModel := &vpcv1.InstanceGroupManagerPrototypeInstanceGroupManagerAutoScalePrototype{
				ManagerType:        core.StringPtr("autoscale"),
				MaxMembershipCount: core.Int64Ptr(int64(10)),
			}

			createInstanceGroupManagerOptions := vpcService.NewCreateInstanceGroupManagerOptions(
				"testString",
				instanceGroupManagerPrototypeModel,
			)

			instanceGroupManager, response, err := vpcService.CreateInstanceGroupManager(createInstanceGroupManagerOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceGroupManager, "", "  ")
			fmt.Println(string(b))

			// end-create_instance_group_manager

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(instanceGroupManager).ToNot(BeNil())

		})
		It(`GetInstanceGroupManager request example`, func() {
			fmt.Println("\nGetInstanceGroupManager() result:")
			// begin-get_instance_group_manager

			getInstanceGroupManagerOptions := vpcService.NewGetInstanceGroupManagerOptions(
				"testString",
				"testString",
			)

			instanceGroupManager, response, err := vpcService.GetInstanceGroupManager(getInstanceGroupManagerOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceGroupManager, "", "  ")
			fmt.Println(string(b))

			// end-get_instance_group_manager

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManager).ToNot(BeNil())

		})
		It(`UpdateInstanceGroupManager request example`, func() {
			fmt.Println("\nUpdateInstanceGroupManager() result:")
			// begin-update_instance_group_manager

			instanceGroupManagerPatchModel := &vpcv1.InstanceGroupManagerPatch{}
			instanceGroupManagerPatchModelAsPatch, asPatchErr := instanceGroupManagerPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateInstanceGroupManagerOptions := vpcService.NewUpdateInstanceGroupManagerOptions(
				"testString",
				"testString",
				instanceGroupManagerPatchModelAsPatch,
			)

			instanceGroupManager, response, err := vpcService.UpdateInstanceGroupManager(updateInstanceGroupManagerOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceGroupManager, "", "  ")
			fmt.Println(string(b))

			// end-update_instance_group_manager

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManager).ToNot(BeNil())

		})
		It(`ListInstanceGroupManagerActions request example`, func() {
			fmt.Println("\nListInstanceGroupManagerActions() result:")
			// begin-list_instance_group_manager_actions

			listInstanceGroupManagerActionsOptions := vpcService.NewListInstanceGroupManagerActionsOptions(
				"testString",
				"testString",
			)

			instanceGroupManagerActionsCollection, response, err := vpcService.ListInstanceGroupManagerActions(listInstanceGroupManagerActionsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceGroupManagerActionsCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_instance_group_manager_actions

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManagerActionsCollection).ToNot(BeNil())

		})
		It(`CreateInstanceGroupManagerAction request example`, func() {
			fmt.Println("\nCreateInstanceGroupManagerAction() result:")
			// begin-create_instance_group_manager_action

			instanceGroupManagerScheduledActionGroupPrototypeModel := &vpcv1.InstanceGroupManagerScheduledActionGroupPrototype{
				MembershipCount: core.Int64Ptr(int64(10)),
			}

			instanceGroupManagerActionPrototypeModel := &vpcv1.InstanceGroupManagerActionPrototypeScheduledActionPrototypeByRunAtByGroup{
				Group: instanceGroupManagerScheduledActionGroupPrototypeModel,
			}

			createInstanceGroupManagerActionOptions := vpcService.NewCreateInstanceGroupManagerActionOptions(
				"testString",
				"testString",
				instanceGroupManagerActionPrototypeModel,
			)

			instanceGroupManagerAction, response, err := vpcService.CreateInstanceGroupManagerAction(createInstanceGroupManagerActionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceGroupManagerAction, "", "  ")
			fmt.Println(string(b))

			// end-create_instance_group_manager_action

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(instanceGroupManagerAction).ToNot(BeNil())

		})
		It(`GetInstanceGroupManagerAction request example`, func() {
			fmt.Println("\nGetInstanceGroupManagerAction() result:")
			// begin-get_instance_group_manager_action

			getInstanceGroupManagerActionOptions := vpcService.NewGetInstanceGroupManagerActionOptions(
				"testString",
				"testString",
				"testString",
			)

			instanceGroupManagerAction, response, err := vpcService.GetInstanceGroupManagerAction(getInstanceGroupManagerActionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceGroupManagerAction, "", "  ")
			fmt.Println(string(b))

			// end-get_instance_group_manager_action

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManagerAction).ToNot(BeNil())

		})
		It(`UpdateInstanceGroupManagerAction request example`, func() {
			fmt.Println("\nUpdateInstanceGroupManagerAction() result:")
			// begin-update_instance_group_manager_action

			instanceGroupManagerActionPatchModel := &vpcv1.InstanceGroupManagerActionPatchScheduledActionPatch{}
			instanceGroupManagerActionPatchModelAsPatch, asPatchErr := instanceGroupManagerActionPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateInstanceGroupManagerActionOptions := vpcService.NewUpdateInstanceGroupManagerActionOptions(
				"testString",
				"testString",
				"testString",
				instanceGroupManagerActionPatchModelAsPatch,
			)

			instanceGroupManagerAction, response, err := vpcService.UpdateInstanceGroupManagerAction(updateInstanceGroupManagerActionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceGroupManagerAction, "", "  ")
			fmt.Println(string(b))

			// end-update_instance_group_manager_action

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManagerAction).ToNot(BeNil())

		})
		It(`ListInstanceGroupManagerPolicies request example`, func() {
			fmt.Println("\nListInstanceGroupManagerPolicies() result:")
			// begin-list_instance_group_manager_policies

			listInstanceGroupManagerPoliciesOptions := vpcService.NewListInstanceGroupManagerPoliciesOptions(
				"testString",
				"testString",
			)

			instanceGroupManagerPolicyCollection, response, err := vpcService.ListInstanceGroupManagerPolicies(listInstanceGroupManagerPoliciesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceGroupManagerPolicyCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_instance_group_manager_policies

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManagerPolicyCollection).ToNot(BeNil())

		})
		It(`CreateInstanceGroupManagerPolicy request example`, func() {
			fmt.Println("\nCreateInstanceGroupManagerPolicy() result:")
			// begin-create_instance_group_manager_policy

			instanceGroupManagerPolicyPrototypeModel := &vpcv1.InstanceGroupManagerPolicyPrototypeInstanceGroupManagerTargetPolicyPrototype{
				MetricType:  core.StringPtr("cpu"),
				MetricValue: core.Int64Ptr(int64(38)),
				PolicyType:  core.StringPtr("target"),
			}

			createInstanceGroupManagerPolicyOptions := vpcService.NewCreateInstanceGroupManagerPolicyOptions(
				"testString",
				"testString",
				instanceGroupManagerPolicyPrototypeModel,
			)

			instanceGroupManagerPolicy, response, err := vpcService.CreateInstanceGroupManagerPolicy(createInstanceGroupManagerPolicyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceGroupManagerPolicy, "", "  ")
			fmt.Println(string(b))

			// end-create_instance_group_manager_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(instanceGroupManagerPolicy).ToNot(BeNil())

		})
		It(`GetInstanceGroupManagerPolicy request example`, func() {
			fmt.Println("\nGetInstanceGroupManagerPolicy() result:")
			// begin-get_instance_group_manager_policy

			getInstanceGroupManagerPolicyOptions := vpcService.NewGetInstanceGroupManagerPolicyOptions(
				"testString",
				"testString",
				"testString",
			)

			instanceGroupManagerPolicy, response, err := vpcService.GetInstanceGroupManagerPolicy(getInstanceGroupManagerPolicyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceGroupManagerPolicy, "", "  ")
			fmt.Println(string(b))

			// end-get_instance_group_manager_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManagerPolicy).ToNot(BeNil())

		})
		It(`UpdateInstanceGroupManagerPolicy request example`, func() {
			fmt.Println("\nUpdateInstanceGroupManagerPolicy() result:")
			// begin-update_instance_group_manager_policy

			instanceGroupManagerPolicyPatchModel := &vpcv1.InstanceGroupManagerPolicyPatch{}
			instanceGroupManagerPolicyPatchModelAsPatch, asPatchErr := instanceGroupManagerPolicyPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateInstanceGroupManagerPolicyOptions := vpcService.NewUpdateInstanceGroupManagerPolicyOptions(
				"testString",
				"testString",
				"testString",
				instanceGroupManagerPolicyPatchModelAsPatch,
			)

			instanceGroupManagerPolicy, response, err := vpcService.UpdateInstanceGroupManagerPolicy(updateInstanceGroupManagerPolicyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceGroupManagerPolicy, "", "  ")
			fmt.Println(string(b))

			// end-update_instance_group_manager_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManagerPolicy).ToNot(BeNil())

		})
		It(`ListInstanceGroupMemberships request example`, func() {
			fmt.Println("\nListInstanceGroupMemberships() result:")
			// begin-list_instance_group_memberships

			listInstanceGroupMembershipsOptions := vpcService.NewListInstanceGroupMembershipsOptions(
				"testString",
			)

			instanceGroupMembershipCollection, response, err := vpcService.ListInstanceGroupMemberships(listInstanceGroupMembershipsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceGroupMembershipCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_instance_group_memberships

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupMembershipCollection).ToNot(BeNil())

		})
		It(`GetInstanceGroupMembership request example`, func() {
			fmt.Println("\nGetInstanceGroupMembership() result:")
			// begin-get_instance_group_membership

			getInstanceGroupMembershipOptions := vpcService.NewGetInstanceGroupMembershipOptions(
				"testString",
				"testString",
			)

			instanceGroupMembership, response, err := vpcService.GetInstanceGroupMembership(getInstanceGroupMembershipOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceGroupMembership, "", "  ")
			fmt.Println(string(b))

			// end-get_instance_group_membership

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupMembership).ToNot(BeNil())

		})
		It(`UpdateInstanceGroupMembership request example`, func() {
			fmt.Println("\nUpdateInstanceGroupMembership() result:")
			// begin-update_instance_group_membership

			instanceGroupMembershipPatchModel := &vpcv1.InstanceGroupMembershipPatch{}
			instanceGroupMembershipPatchModelAsPatch, asPatchErr := instanceGroupMembershipPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateInstanceGroupMembershipOptions := vpcService.NewUpdateInstanceGroupMembershipOptions(
				"testString",
				"testString",
				instanceGroupMembershipPatchModelAsPatch,
			)

			instanceGroupMembership, response, err := vpcService.UpdateInstanceGroupMembership(updateInstanceGroupMembershipOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(instanceGroupMembership, "", "  ")
			fmt.Println(string(b))

			// end-update_instance_group_membership

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupMembership).ToNot(BeNil())

		})
		It(`ListDedicatedHostGroups request example`, func() {
			fmt.Println("\nListDedicatedHostGroups() result:")
			// begin-list_dedicated_host_groups

			listDedicatedHostGroupsOptions := vpcService.NewListDedicatedHostGroupsOptions()

			dedicatedHostGroupCollection, response, err := vpcService.ListDedicatedHostGroups(listDedicatedHostGroupsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dedicatedHostGroupCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_dedicated_host_groups

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHostGroupCollection).ToNot(BeNil())

		})
		It(`CreateDedicatedHostGroup request example`, func() {
			fmt.Println("\nCreateDedicatedHostGroup() result:")
			// begin-create_dedicated_host_group

			zoneIdentityModel := &vpcv1.ZoneIdentityByName{
				Name: core.StringPtr("us-south-1"),
			}

			createDedicatedHostGroupOptions := vpcService.NewCreateDedicatedHostGroupOptions()
			createDedicatedHostGroupOptions.SetClass("mx2")
			createDedicatedHostGroupOptions.SetFamily("balanced")
			createDedicatedHostGroupOptions.SetZone(zoneIdentityModel)

			dedicatedHostGroup, response, err := vpcService.CreateDedicatedHostGroup(createDedicatedHostGroupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dedicatedHostGroup, "", "  ")
			fmt.Println(string(b))

			// end-create_dedicated_host_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(dedicatedHostGroup).ToNot(BeNil())

		})
		It(`GetDedicatedHostGroup request example`, func() {
			fmt.Println("\nGetDedicatedHostGroup() result:")
			// begin-get_dedicated_host_group

			getDedicatedHostGroupOptions := vpcService.NewGetDedicatedHostGroupOptions(
				"testString",
			)

			dedicatedHostGroup, response, err := vpcService.GetDedicatedHostGroup(getDedicatedHostGroupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dedicatedHostGroup, "", "  ")
			fmt.Println(string(b))

			// end-get_dedicated_host_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHostGroup).ToNot(BeNil())

		})
		It(`UpdateDedicatedHostGroup request example`, func() {
			fmt.Println("\nUpdateDedicatedHostGroup() result:")
			// begin-update_dedicated_host_group

			dedicatedHostGroupPatchModel := &vpcv1.DedicatedHostGroupPatch{
				Name: core.StringPtr("my-host-group-modified"),
			}
			dedicatedHostGroupPatchModelAsPatch, asPatchErr := dedicatedHostGroupPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateDedicatedHostGroupOptions := vpcService.NewUpdateDedicatedHostGroupOptions(
				"testString",
				dedicatedHostGroupPatchModelAsPatch,
			)

			dedicatedHostGroup, response, err := vpcService.UpdateDedicatedHostGroup(updateDedicatedHostGroupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dedicatedHostGroup, "", "  ")
			fmt.Println(string(b))

			// end-update_dedicated_host_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHostGroup).ToNot(BeNil())

		})
		It(`ListDedicatedHostProfiles request example`, func() {
			fmt.Println("\nListDedicatedHostProfiles() result:")
			// begin-list_dedicated_host_profiles

			listDedicatedHostProfilesOptions := vpcService.NewListDedicatedHostProfilesOptions()

			dedicatedHostProfileCollection, response, err := vpcService.ListDedicatedHostProfiles(listDedicatedHostProfilesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dedicatedHostProfileCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_dedicated_host_profiles

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHostProfileCollection).ToNot(BeNil())

		})
		It(`GetDedicatedHostProfile request example`, func() {
			fmt.Println("\nGetDedicatedHostProfile() result:")
			// begin-get_dedicated_host_profile

			getDedicatedHostProfileOptions := vpcService.NewGetDedicatedHostProfileOptions(
				"testString",
			)

			dedicatedHostProfile, response, err := vpcService.GetDedicatedHostProfile(getDedicatedHostProfileOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dedicatedHostProfile, "", "  ")
			fmt.Println(string(b))

			// end-get_dedicated_host_profile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHostProfile).ToNot(BeNil())

		})
		It(`ListDedicatedHosts request example`, func() {
			fmt.Println("\nListDedicatedHosts() result:")
			// begin-list_dedicated_hosts

			listDedicatedHostsOptions := vpcService.NewListDedicatedHostsOptions()

			dedicatedHostCollection, response, err := vpcService.ListDedicatedHosts(listDedicatedHostsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dedicatedHostCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_dedicated_hosts

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHostCollection).ToNot(BeNil())

		})
		It(`CreateDedicatedHost request example`, func() {
			fmt.Println("\nCreateDedicatedHost() result:")
			// begin-create_dedicated_host

			dedicatedHostProfileIdentityModel := &vpcv1.DedicatedHostProfileIdentityByName{
				Name: core.StringPtr("m-62x496"),
			}

			dedicatedHostGroupIdentityModel := &vpcv1.DedicatedHostGroupIdentityByID{
				ID: core.StringPtr("0c8eccb4-271c-4518-956c-32bfce5cf83b"),
			}

			dedicatedHostPrototypeModel := &vpcv1.DedicatedHostPrototypeDedicatedHostByGroup{
				Profile: dedicatedHostProfileIdentityModel,
				Group:   dedicatedHostGroupIdentityModel,
			}

			createDedicatedHostOptions := vpcService.NewCreateDedicatedHostOptions(
				dedicatedHostPrototypeModel,
			)

			dedicatedHost, response, err := vpcService.CreateDedicatedHost(createDedicatedHostOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dedicatedHost, "", "  ")
			fmt.Println(string(b))

			// end-create_dedicated_host

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(dedicatedHost).ToNot(BeNil())

		})
		It(`ListDedicatedHostDisks request example`, func() {
			fmt.Println("\nListDedicatedHostDisks() result:")
			// begin-list_dedicated_host_disks

			listDedicatedHostDisksOptions := vpcService.NewListDedicatedHostDisksOptions(
				"testString",
			)

			dedicatedHostDiskCollection, response, err := vpcService.ListDedicatedHostDisks(listDedicatedHostDisksOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dedicatedHostDiskCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_dedicated_host_disks

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHostDiskCollection).ToNot(BeNil())

		})
		It(`GetDedicatedHostDisk request example`, func() {
			fmt.Println("\nGetDedicatedHostDisk() result:")
			// begin-get_dedicated_host_disk

			getDedicatedHostDiskOptions := vpcService.NewGetDedicatedHostDiskOptions(
				"testString",
				"testString",
			)

			dedicatedHostDisk, response, err := vpcService.GetDedicatedHostDisk(getDedicatedHostDiskOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dedicatedHostDisk, "", "  ")
			fmt.Println(string(b))

			// end-get_dedicated_host_disk

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHostDisk).ToNot(BeNil())

		})
		It(`UpdateDedicatedHostDisk request example`, func() {
			fmt.Println("\nUpdateDedicatedHostDisk() result:")
			// begin-update_dedicated_host_disk

			dedicatedHostDiskPatchModel := &vpcv1.DedicatedHostDiskPatch{}
			dedicatedHostDiskPatchModelAsPatch, asPatchErr := dedicatedHostDiskPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateDedicatedHostDiskOptions := vpcService.NewUpdateDedicatedHostDiskOptions(
				"testString",
				"testString",
				dedicatedHostDiskPatchModelAsPatch,
			)

			dedicatedHostDisk, response, err := vpcService.UpdateDedicatedHostDisk(updateDedicatedHostDiskOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dedicatedHostDisk, "", "  ")
			fmt.Println(string(b))

			// end-update_dedicated_host_disk

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHostDisk).ToNot(BeNil())

		})
		It(`GetDedicatedHost request example`, func() {
			fmt.Println("\nGetDedicatedHost() result:")
			// begin-get_dedicated_host

			getDedicatedHostOptions := vpcService.NewGetDedicatedHostOptions(
				"testString",
			)

			dedicatedHost, response, err := vpcService.GetDedicatedHost(getDedicatedHostOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dedicatedHost, "", "  ")
			fmt.Println(string(b))

			// end-get_dedicated_host

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHost).ToNot(BeNil())

		})
		It(`UpdateDedicatedHost request example`, func() {
			fmt.Println("\nUpdateDedicatedHost() result:")
			// begin-update_dedicated_host

			dedicatedHostPatchModel := &vpcv1.DedicatedHostPatch{}
			dedicatedHostPatchModelAsPatch, asPatchErr := dedicatedHostPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateDedicatedHostOptions := vpcService.NewUpdateDedicatedHostOptions(
				"testString",
				dedicatedHostPatchModelAsPatch,
			)

			dedicatedHost, response, err := vpcService.UpdateDedicatedHost(updateDedicatedHostOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(dedicatedHost, "", "  ")
			fmt.Println(string(b))

			// end-update_dedicated_host

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHost).ToNot(BeNil())

		})
		It(`ListVolumeProfiles request example`, func() {
			fmt.Println("\nListVolumeProfiles() result:")
			// begin-list_volume_profiles

			listVolumeProfilesOptions := vpcService.NewListVolumeProfilesOptions()

			volumeProfileCollection, response, err := vpcService.ListVolumeProfiles(listVolumeProfilesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(volumeProfileCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_volume_profiles

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeProfileCollection).ToNot(BeNil())

		})
		It(`GetVolumeProfile request example`, func() {
			fmt.Println("\nGetVolumeProfile() result:")
			// begin-get_volume_profile

			getVolumeProfileOptions := vpcService.NewGetVolumeProfileOptions(
				"testString",
			)

			volumeProfile, response, err := vpcService.GetVolumeProfile(getVolumeProfileOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(volumeProfile, "", "  ")
			fmt.Println(string(b))

			// end-get_volume_profile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeProfile).ToNot(BeNil())

		})
		It(`ListVolumes request example`, func() {
			fmt.Println("\nListVolumes() result:")
			// begin-list_volumes

			listVolumesOptions := vpcService.NewListVolumesOptions()

			volumeCollection, response, err := vpcService.ListVolumes(listVolumesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(volumeCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_volumes

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeCollection).ToNot(BeNil())

		})
		It(`CreateVolume request example`, func() {
			fmt.Println("\nCreateVolume() result:")
			// begin-create_volume

			volumeProfileIdentityModel := &vpcv1.VolumeProfileIdentityByName{
				Name: core.StringPtr("5iops-tier"),
			}

			zoneIdentityModel := &vpcv1.ZoneIdentityByName{
				Name: core.StringPtr("us-south-1"),
			}

			volumePrototypeModel := &vpcv1.VolumePrototypeVolumeByCapacity{
				Profile:  volumeProfileIdentityModel,
				Zone:     zoneIdentityModel,
				Capacity: core.Int64Ptr(int64(100)),
			}

			createVolumeOptions := vpcService.NewCreateVolumeOptions(
				volumePrototypeModel,
			)

			volume, response, err := vpcService.CreateVolume(createVolumeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(volume, "", "  ")
			fmt.Println(string(b))

			// end-create_volume

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(volume).ToNot(BeNil())

		})
		It(`GetVolume request example`, func() {
			fmt.Println("\nGetVolume() result:")
			// begin-get_volume

			getVolumeOptions := vpcService.NewGetVolumeOptions(
				"testString",
			)

			volume, response, err := vpcService.GetVolume(getVolumeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(volume, "", "  ")
			fmt.Println(string(b))

			// end-get_volume

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volume).ToNot(BeNil())

		})
		It(`UpdateVolume request example`, func() {
			fmt.Println("\nUpdateVolume() result:")
			// begin-update_volume

			volumePatchModel := &vpcv1.VolumePatch{}
			volumePatchModelAsPatch, asPatchErr := volumePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateVolumeOptions := vpcService.NewUpdateVolumeOptions(
				"testString",
				volumePatchModelAsPatch,
			)

			volume, response, err := vpcService.UpdateVolume(updateVolumeOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(volume, "", "  ")
			fmt.Println(string(b))

			// end-update_volume

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volume).ToNot(BeNil())

		})
		It(`ListSnapshots request example`, func() {
			fmt.Println("\nListSnapshots() result:")
			// begin-list_snapshots

			listSnapshotsOptions := vpcService.NewListSnapshotsOptions()
			listSnapshotsOptions.SetSort("name")

			snapshotCollection, response, err := vpcService.ListSnapshots(listSnapshotsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(snapshotCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_snapshots

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(snapshotCollection).ToNot(BeNil())

		})
		It(`CreateSnapshot request example`, func() {
			fmt.Println("\nCreateSnapshot() result:")
			// begin-create_snapshot

			volumeIdentityModel := &vpcv1.VolumeIdentityByID{
				ID: core.StringPtr("1a6b7274-678d-4dfb-8981-c71dd9d4daa5"),
			}

			createSnapshotOptions := vpcService.NewCreateSnapshotOptions(
				volumeIdentityModel,
			)

			snapshot, response, err := vpcService.CreateSnapshot(createSnapshotOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(snapshot, "", "  ")
			fmt.Println(string(b))

			// end-create_snapshot

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(snapshot).ToNot(BeNil())

		})
		It(`GetSnapshot request example`, func() {
			fmt.Println("\nGetSnapshot() result:")
			// begin-get_snapshot

			getSnapshotOptions := vpcService.NewGetSnapshotOptions(
				"testString",
			)

			snapshot, response, err := vpcService.GetSnapshot(getSnapshotOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(snapshot, "", "  ")
			fmt.Println(string(b))

			// end-get_snapshot

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(snapshot).ToNot(BeNil())

		})
		It(`UpdateSnapshot request example`, func() {
			fmt.Println("\nUpdateSnapshot() result:")
			// begin-update_snapshot

			snapshotPatchModel := &vpcv1.SnapshotPatch{}
			snapshotPatchModelAsPatch, asPatchErr := snapshotPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateSnapshotOptions := vpcService.NewUpdateSnapshotOptions(
				"testString",
				snapshotPatchModelAsPatch,
			)

			snapshot, response, err := vpcService.UpdateSnapshot(updateSnapshotOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(snapshot, "", "  ")
			fmt.Println(string(b))

			// end-update_snapshot

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(snapshot).ToNot(BeNil())

		})
		It(`ListRegions request example`, func() {
			fmt.Println("\nListRegions() result:")
			// begin-list_regions

			listRegionsOptions := vpcService.NewListRegionsOptions()

			regionCollection, response, err := vpcService.ListRegions(listRegionsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(regionCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_regions

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(regionCollection).ToNot(BeNil())

		})
		It(`GetRegion request example`, func() {
			fmt.Println("\nGetRegion() result:")
			// begin-get_region

			getRegionOptions := vpcService.NewGetRegionOptions(
				"testString",
			)

			region, response, err := vpcService.GetRegion(getRegionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(region, "", "  ")
			fmt.Println(string(b))

			// end-get_region

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(region).ToNot(BeNil())

		})
		It(`ListRegionZones request example`, func() {
			fmt.Println("\nListRegionZones() result:")
			// begin-list_region_zones

			listRegionZonesOptions := vpcService.NewListRegionZonesOptions(
				"testString",
			)

			zoneCollection, response, err := vpcService.ListRegionZones(listRegionZonesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(zoneCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_region_zones

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(zoneCollection).ToNot(BeNil())

		})
		It(`GetRegionZone request example`, func() {
			fmt.Println("\nGetRegionZone() result:")
			// begin-get_region_zone

			getRegionZoneOptions := vpcService.NewGetRegionZoneOptions(
				"testString",
				"testString",
			)

			zone, response, err := vpcService.GetRegionZone(getRegionZoneOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(zone, "", "  ")
			fmt.Println(string(b))

			// end-get_region_zone

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(zone).ToNot(BeNil())

		})
		It(`ListPublicGateways request example`, func() {
			fmt.Println("\nListPublicGateways() result:")
			// begin-list_public_gateways

			listPublicGatewaysOptions := vpcService.NewListPublicGatewaysOptions()

			publicGatewayCollection, response, err := vpcService.ListPublicGateways(listPublicGatewaysOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(publicGatewayCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_public_gateways

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(publicGatewayCollection).ToNot(BeNil())

		})
		It(`CreatePublicGateway request example`, func() {
			fmt.Println("\nCreatePublicGateway() result:")
			// begin-create_public_gateway

			vpcIdentityModel := &vpcv1.VPCIdentityByID{
				ID: core.StringPtr("4727d842-f94f-4a2d-824a-9bc9b02c523b"),
			}

			zoneIdentityModel := &vpcv1.ZoneIdentityByName{
				Name: core.StringPtr("us-south-1"),
			}

			createPublicGatewayOptions := vpcService.NewCreatePublicGatewayOptions(
				vpcIdentityModel,
				zoneIdentityModel,
			)

			publicGateway, response, err := vpcService.CreatePublicGateway(createPublicGatewayOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(publicGateway, "", "  ")
			fmt.Println(string(b))

			// end-create_public_gateway

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(publicGateway).ToNot(BeNil())

		})
		It(`GetPublicGateway request example`, func() {
			fmt.Println("\nGetPublicGateway() result:")
			// begin-get_public_gateway

			getPublicGatewayOptions := vpcService.NewGetPublicGatewayOptions(
				"testString",
			)

			publicGateway, response, err := vpcService.GetPublicGateway(getPublicGatewayOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(publicGateway, "", "  ")
			fmt.Println(string(b))

			// end-get_public_gateway

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(publicGateway).ToNot(BeNil())

		})
		It(`UpdatePublicGateway request example`, func() {
			fmt.Println("\nUpdatePublicGateway() result:")
			// begin-update_public_gateway

			publicGatewayPatchModel := &vpcv1.PublicGatewayPatch{}
			publicGatewayPatchModelAsPatch, asPatchErr := publicGatewayPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updatePublicGatewayOptions := vpcService.NewUpdatePublicGatewayOptions(
				"testString",
				publicGatewayPatchModelAsPatch,
			)

			publicGateway, response, err := vpcService.UpdatePublicGateway(updatePublicGatewayOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(publicGateway, "", "  ")
			fmt.Println(string(b))

			// end-update_public_gateway

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(publicGateway).ToNot(BeNil())

		})
		It(`ListFloatingIps request example`, func() {
			fmt.Println("\nListFloatingIps() result:")
			// begin-list_floating_ips

			listFloatingIpsOptions := vpcService.NewListFloatingIpsOptions()

			floatingIPCollection, response, err := vpcService.ListFloatingIps(listFloatingIpsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(floatingIPCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_floating_ips

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(floatingIPCollection).ToNot(BeNil())

		})
		It(`CreateFloatingIP request example`, func() {
			fmt.Println("\nCreateFloatingIP() result:")
			// begin-create_floating_ip

			zoneIdentityModel := &vpcv1.ZoneIdentityByName{
				Name: core.StringPtr("us-south-1"),
			}

			floatingIPPrototypeModel := &vpcv1.FloatingIPPrototypeFloatingIPByZone{
				Zone: zoneIdentityModel,
			}

			createFloatingIPOptions := vpcService.NewCreateFloatingIPOptions(
				floatingIPPrototypeModel,
			)

			floatingIP, response, err := vpcService.CreateFloatingIP(createFloatingIPOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(floatingIP, "", "  ")
			fmt.Println(string(b))

			// end-create_floating_ip

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(floatingIP).ToNot(BeNil())

		})
		It(`GetFloatingIP request example`, func() {
			fmt.Println("\nGetFloatingIP() result:")
			// begin-get_floating_ip

			getFloatingIPOptions := vpcService.NewGetFloatingIPOptions(
				"testString",
			)

			floatingIP, response, err := vpcService.GetFloatingIP(getFloatingIPOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(floatingIP, "", "  ")
			fmt.Println(string(b))

			// end-get_floating_ip

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(floatingIP).ToNot(BeNil())

		})
		It(`UpdateFloatingIP request example`, func() {
			fmt.Println("\nUpdateFloatingIP() result:")
			// begin-update_floating_ip

			floatingIPPatchModel := &vpcv1.FloatingIPPatch{}
			floatingIPPatchModelAsPatch, asPatchErr := floatingIPPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateFloatingIPOptions := vpcService.NewUpdateFloatingIPOptions(
				"testString",
				floatingIPPatchModelAsPatch,
			)

			floatingIP, response, err := vpcService.UpdateFloatingIP(updateFloatingIPOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(floatingIP, "", "  ")
			fmt.Println(string(b))

			// end-update_floating_ip

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(floatingIP).ToNot(BeNil())

		})
		It(`ListNetworkAcls request example`, func() {
			fmt.Println("\nListNetworkAcls() result:")
			// begin-list_network_acls

			listNetworkAclsOptions := vpcService.NewListNetworkAclsOptions()

			networkACLCollection, response, err := vpcService.ListNetworkAcls(listNetworkAclsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(networkACLCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_network_acls

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkACLCollection).ToNot(BeNil())

		})
		It(`CreateNetworkACL request example`, func() {
			fmt.Println("\nCreateNetworkACL() result:")
			// begin-create_network_acl

			vpcIdentityModel := &vpcv1.VPCIdentityByID{
				ID: core.StringPtr("f0aae929-7047-46d1-92e1-9102b07a7f6f"),
			}

			networkACLPrototypeModel := &vpcv1.NetworkACLPrototypeNetworkACLByRules{
				Name: core.StringPtr("my-network-acl"),
				VPC:  vpcIdentityModel,
			}

			createNetworkACLOptions := vpcService.NewCreateNetworkACLOptions()
			createNetworkACLOptions.SetNetworkACLPrototype(networkACLPrototypeModel)

			networkACL, response, err := vpcService.CreateNetworkACL(createNetworkACLOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(networkACL, "", "  ")
			fmt.Println(string(b))

			// end-create_network_acl

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(networkACL).ToNot(BeNil())

		})
		It(`GetNetworkACL request example`, func() {
			fmt.Println("\nGetNetworkACL() result:")
			// begin-get_network_acl

			getNetworkACLOptions := vpcService.NewGetNetworkACLOptions(
				"testString",
			)

			networkACL, response, err := vpcService.GetNetworkACL(getNetworkACLOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(networkACL, "", "  ")
			fmt.Println(string(b))

			// end-get_network_acl

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkACL).ToNot(BeNil())

		})
		It(`UpdateNetworkACL request example`, func() {
			fmt.Println("\nUpdateNetworkACL() result:")
			// begin-update_network_acl

			networkACLPatchModel := &vpcv1.NetworkACLPatch{}
			networkACLPatchModelAsPatch, asPatchErr := networkACLPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateNetworkACLOptions := vpcService.NewUpdateNetworkACLOptions(
				"testString",
				networkACLPatchModelAsPatch,
			)

			networkACL, response, err := vpcService.UpdateNetworkACL(updateNetworkACLOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(networkACL, "", "  ")
			fmt.Println(string(b))

			// end-update_network_acl

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkACL).ToNot(BeNil())

		})
		It(`ListNetworkACLRules request example`, func() {
			fmt.Println("\nListNetworkACLRules() result:")
			// begin-list_network_acl_rules

			listNetworkACLRulesOptions := vpcService.NewListNetworkACLRulesOptions(
				"testString",
			)

			networkACLRuleCollection, response, err := vpcService.ListNetworkACLRules(listNetworkACLRulesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(networkACLRuleCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_network_acl_rules

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkACLRuleCollection).ToNot(BeNil())

		})
		It(`CreateNetworkACLRule request example`, func() {
			fmt.Println("\nCreateNetworkACLRule() result:")
			// begin-create_network_acl_rule

			networkACLRulePrototypeModel := &vpcv1.NetworkACLRulePrototypeNetworkACLRuleProtocolAll{
				Action:      core.StringPtr("allow"),
				Destination: core.StringPtr("192.168.3.2/32"),
				Direction:   core.StringPtr("inbound"),
				Source:      core.StringPtr("192.168.3.2/32"),
				Protocol:    core.StringPtr("all"),
			}

			createNetworkACLRuleOptions := vpcService.NewCreateNetworkACLRuleOptions(
				"testString",
				networkACLRulePrototypeModel,
			)

			networkACLRule, response, err := vpcService.CreateNetworkACLRule(createNetworkACLRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(networkACLRule, "", "  ")
			fmt.Println(string(b))

			// end-create_network_acl_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(networkACLRule).ToNot(BeNil())

		})
		It(`GetNetworkACLRule request example`, func() {
			fmt.Println("\nGetNetworkACLRule() result:")
			// begin-get_network_acl_rule

			getNetworkACLRuleOptions := vpcService.NewGetNetworkACLRuleOptions(
				"testString",
				"testString",
			)

			networkACLRule, response, err := vpcService.GetNetworkACLRule(getNetworkACLRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(networkACLRule, "", "  ")
			fmt.Println(string(b))

			// end-get_network_acl_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkACLRule).ToNot(BeNil())

		})
		It(`UpdateNetworkACLRule request example`, func() {
			fmt.Println("\nUpdateNetworkACLRule() result:")
			// begin-update_network_acl_rule

			networkACLRulePatchModel := &vpcv1.NetworkACLRulePatch{}
			networkACLRulePatchModelAsPatch, asPatchErr := networkACLRulePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateNetworkACLRuleOptions := vpcService.NewUpdateNetworkACLRuleOptions(
				"testString",
				"testString",
				networkACLRulePatchModelAsPatch,
			)

			networkACLRule, response, err := vpcService.UpdateNetworkACLRule(updateNetworkACLRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(networkACLRule, "", "  ")
			fmt.Println(string(b))

			// end-update_network_acl_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkACLRule).ToNot(BeNil())

		})
		It(`ListSecurityGroups request example`, func() {
			fmt.Println("\nListSecurityGroups() result:")
			// begin-list_security_groups

			listSecurityGroupsOptions := vpcService.NewListSecurityGroupsOptions()

			securityGroupCollection, response, err := vpcService.ListSecurityGroups(listSecurityGroupsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(securityGroupCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_security_groups

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(securityGroupCollection).ToNot(BeNil())

		})
		It(`CreateSecurityGroup request example`, func() {
			fmt.Println("\nCreateSecurityGroup() result:")
			// begin-create_security_group

			vpcIdentityModel := &vpcv1.VPCIdentityByID{
				ID: core.StringPtr("4727d842-f94f-4a2d-824a-9bc9b02c523b"),
			}

			createSecurityGroupOptions := vpcService.NewCreateSecurityGroupOptions(
				vpcIdentityModel,
			)

			securityGroup, response, err := vpcService.CreateSecurityGroup(createSecurityGroupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(securityGroup, "", "  ")
			fmt.Println(string(b))

			// end-create_security_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(securityGroup).ToNot(BeNil())

		})
		It(`GetSecurityGroup request example`, func() {
			fmt.Println("\nGetSecurityGroup() result:")
			// begin-get_security_group

			getSecurityGroupOptions := vpcService.NewGetSecurityGroupOptions(
				"testString",
			)

			securityGroup, response, err := vpcService.GetSecurityGroup(getSecurityGroupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(securityGroup, "", "  ")
			fmt.Println(string(b))

			// end-get_security_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(securityGroup).ToNot(BeNil())

		})
		It(`UpdateSecurityGroup request example`, func() {
			fmt.Println("\nUpdateSecurityGroup() result:")
			// begin-update_security_group

			securityGroupPatchModel := &vpcv1.SecurityGroupPatch{}
			securityGroupPatchModelAsPatch, asPatchErr := securityGroupPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateSecurityGroupOptions := vpcService.NewUpdateSecurityGroupOptions(
				"testString",
				securityGroupPatchModelAsPatch,
			)

			securityGroup, response, err := vpcService.UpdateSecurityGroup(updateSecurityGroupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(securityGroup, "", "  ")
			fmt.Println(string(b))

			// end-update_security_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(securityGroup).ToNot(BeNil())

		})
		It(`ListSecurityGroupNetworkInterfaces request example`, func() {
			fmt.Println("\nListSecurityGroupNetworkInterfaces() result:")
			// begin-list_security_group_network_interfaces

			listSecurityGroupNetworkInterfacesOptions := vpcService.NewListSecurityGroupNetworkInterfacesOptions(
				"testString",
			)

			networkInterfaceCollection, response, err := vpcService.ListSecurityGroupNetworkInterfaces(listSecurityGroupNetworkInterfacesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(networkInterfaceCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_security_group_network_interfaces

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkInterfaceCollection).ToNot(BeNil())

		})
		It(`GetSecurityGroupNetworkInterface request example`, func() {
			fmt.Println("\nGetSecurityGroupNetworkInterface() result:")
			// begin-get_security_group_network_interface

			getSecurityGroupNetworkInterfaceOptions := vpcService.NewGetSecurityGroupNetworkInterfaceOptions(
				"testString",
				"testString",
			)

			networkInterface, response, err := vpcService.GetSecurityGroupNetworkInterface(getSecurityGroupNetworkInterfaceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(networkInterface, "", "  ")
			fmt.Println(string(b))

			// end-get_security_group_network_interface

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkInterface).ToNot(BeNil())

		})
		It(`AddSecurityGroupNetworkInterface request example`, func() {
			fmt.Println("\nAddSecurityGroupNetworkInterface() result:")
			// begin-add_security_group_network_interface

			addSecurityGroupNetworkInterfaceOptions := vpcService.NewAddSecurityGroupNetworkInterfaceOptions(
				"testString",
				"testString",
			)

			networkInterface, response, err := vpcService.AddSecurityGroupNetworkInterface(addSecurityGroupNetworkInterfaceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(networkInterface, "", "  ")
			fmt.Println(string(b))

			// end-add_security_group_network_interface

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(networkInterface).ToNot(BeNil())

		})
		It(`ListSecurityGroupRules request example`, func() {
			fmt.Println("\nListSecurityGroupRules() result:")
			// begin-list_security_group_rules

			listSecurityGroupRulesOptions := vpcService.NewListSecurityGroupRulesOptions(
				"testString",
			)

			securityGroupRuleCollection, response, err := vpcService.ListSecurityGroupRules(listSecurityGroupRulesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(securityGroupRuleCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_security_group_rules

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(securityGroupRuleCollection).ToNot(BeNil())

		})
		It(`CreateSecurityGroupRule request example`, func() {
			fmt.Println("\nCreateSecurityGroupRule() result:")
			// begin-create_security_group_rule

			securityGroupRulePrototypeModel := &vpcv1.SecurityGroupRulePrototypeSecurityGroupRuleProtocolTcpudp{
				Direction: core.StringPtr("inbound"),
				Protocol:  core.StringPtr("udp"),
			}

			createSecurityGroupRuleOptions := vpcService.NewCreateSecurityGroupRuleOptions(
				"testString",
				securityGroupRulePrototypeModel,
			)

			securityGroupRule, response, err := vpcService.CreateSecurityGroupRule(createSecurityGroupRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(securityGroupRule, "", "  ")
			fmt.Println(string(b))

			// end-create_security_group_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(securityGroupRule).ToNot(BeNil())

		})
		It(`GetSecurityGroupRule request example`, func() {
			fmt.Println("\nGetSecurityGroupRule() result:")
			// begin-get_security_group_rule

			getSecurityGroupRuleOptions := vpcService.NewGetSecurityGroupRuleOptions(
				"testString",
				"testString",
			)

			securityGroupRule, response, err := vpcService.GetSecurityGroupRule(getSecurityGroupRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(securityGroupRule, "", "  ")
			fmt.Println(string(b))

			// end-get_security_group_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(securityGroupRule).ToNot(BeNil())

		})
		It(`UpdateSecurityGroupRule request example`, func() {
			fmt.Println("\nUpdateSecurityGroupRule() result:")
			// begin-update_security_group_rule

			securityGroupRulePatchModel := &vpcv1.SecurityGroupRulePatch{}
			securityGroupRulePatchModelAsPatch, asPatchErr := securityGroupRulePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateSecurityGroupRuleOptions := vpcService.NewUpdateSecurityGroupRuleOptions(
				"testString",
				"testString",
				securityGroupRulePatchModelAsPatch,
			)

			securityGroupRule, response, err := vpcService.UpdateSecurityGroupRule(updateSecurityGroupRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(securityGroupRule, "", "  ")
			fmt.Println(string(b))

			// end-update_security_group_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(securityGroupRule).ToNot(BeNil())

		})
		It(`ListSecurityGroupTargets request example`, func() {
			fmt.Println("\nListSecurityGroupTargets() result:")
			// begin-list_security_group_targets

			listSecurityGroupTargetsOptions := vpcService.NewListSecurityGroupTargetsOptions(
				"testString",
			)

			securityGroupTargetCollection, response, err := vpcService.ListSecurityGroupTargets(listSecurityGroupTargetsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(securityGroupTargetCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_security_group_targets

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(securityGroupTargetCollection).ToNot(BeNil())

		})
		It(`GetSecurityGroupTarget request example`, func() {
			fmt.Println("\nGetSecurityGroupTarget() result:")
			// begin-get_security_group_target

			getSecurityGroupTargetOptions := vpcService.NewGetSecurityGroupTargetOptions(
				"testString",
				"testString",
			)

			securityGroupTargetReference, response, err := vpcService.GetSecurityGroupTarget(getSecurityGroupTargetOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(securityGroupTargetReference, "", "  ")
			fmt.Println(string(b))

			// end-get_security_group_target

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(securityGroupTargetReference).ToNot(BeNil())

		})
		It(`CreateSecurityGroupTargetBinding request example`, func() {
			fmt.Println("\nCreateSecurityGroupTargetBinding() result:")
			// begin-create_security_group_target_binding

			createSecurityGroupTargetBindingOptions := vpcService.NewCreateSecurityGroupTargetBindingOptions(
				"testString",
				"testString",
			)

			securityGroupTargetReference, response, err := vpcService.CreateSecurityGroupTargetBinding(createSecurityGroupTargetBindingOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(securityGroupTargetReference, "", "  ")
			fmt.Println(string(b))

			// end-create_security_group_target_binding

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(securityGroupTargetReference).ToNot(BeNil())

		})
		It(`ListIkePolicies request example`, func() {
			fmt.Println("\nListIkePolicies() result:")
			// begin-list_ike_policies

			listIkePoliciesOptions := vpcService.NewListIkePoliciesOptions()

			ikePolicyCollection, response, err := vpcService.ListIkePolicies(listIkePoliciesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(ikePolicyCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_ike_policies

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ikePolicyCollection).ToNot(BeNil())

		})
		It(`CreateIkePolicy request example`, func() {
			fmt.Println("\nCreateIkePolicy() result:")
			// begin-create_ike_policy

			createIkePolicyOptions := vpcService.NewCreateIkePolicyOptions(
				"md5",
				int64(2),
				"triple_des",
				int64(1),
			)

			ikePolicy, response, err := vpcService.CreateIkePolicy(createIkePolicyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(ikePolicy, "", "  ")
			fmt.Println(string(b))

			// end-create_ike_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(ikePolicy).ToNot(BeNil())

		})
		It(`GetIkePolicy request example`, func() {
			fmt.Println("\nGetIkePolicy() result:")
			// begin-get_ike_policy

			getIkePolicyOptions := vpcService.NewGetIkePolicyOptions(
				"testString",
			)

			ikePolicy, response, err := vpcService.GetIkePolicy(getIkePolicyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(ikePolicy, "", "  ")
			fmt.Println(string(b))

			// end-get_ike_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ikePolicy).ToNot(BeNil())

		})
		It(`UpdateIkePolicy request example`, func() {
			fmt.Println("\nUpdateIkePolicy() result:")
			// begin-update_ike_policy

			updateIkePolicyOptions := vpcService.NewUpdateIkePolicyOptions(
				"testString",
				make(map[string]interface{}),
			)

			ikePolicy, response, err := vpcService.UpdateIkePolicy(updateIkePolicyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(ikePolicy, "", "  ")
			fmt.Println(string(b))

			// end-update_ike_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ikePolicy).ToNot(BeNil())

		})
		It(`ListIkePolicyConnections request example`, func() {
			fmt.Println("\nListIkePolicyConnections() result:")
			// begin-list_ike_policy_connections

			listIkePolicyConnectionsOptions := vpcService.NewListIkePolicyConnectionsOptions(
				"testString",
			)

			vpnGatewayConnectionCollection, response, err := vpcService.ListIkePolicyConnections(listIkePolicyConnectionsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(vpnGatewayConnectionCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_ike_policy_connections

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnGatewayConnectionCollection).ToNot(BeNil())

		})
		It(`ListIpsecPolicies request example`, func() {
			fmt.Println("\nListIpsecPolicies() result:")
			// begin-list_ipsec_policies

			listIpsecPoliciesOptions := vpcService.NewListIpsecPoliciesOptions()

			iPsecPolicyCollection, response, err := vpcService.ListIpsecPolicies(listIpsecPoliciesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(iPsecPolicyCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_ipsec_policies

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(iPsecPolicyCollection).ToNot(BeNil())

		})
		It(`CreateIpsecPolicy request example`, func() {
			fmt.Println("\nCreateIpsecPolicy() result:")
			// begin-create_ipsec_policy

			createIpsecPolicyOptions := vpcService.NewCreateIpsecPolicyOptions(
				"md5",
				"triple_des",
				"disabled",
			)

			iPsecPolicy, response, err := vpcService.CreateIpsecPolicy(createIpsecPolicyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(iPsecPolicy, "", "  ")
			fmt.Println(string(b))

			// end-create_ipsec_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(iPsecPolicy).ToNot(BeNil())

		})
		It(`GetIpsecPolicy request example`, func() {
			fmt.Println("\nGetIpsecPolicy() result:")
			// begin-get_ipsec_policy

			getIpsecPolicyOptions := vpcService.NewGetIpsecPolicyOptions(
				"testString",
			)

			iPsecPolicy, response, err := vpcService.GetIpsecPolicy(getIpsecPolicyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(iPsecPolicy, "", "  ")
			fmt.Println(string(b))

			// end-get_ipsec_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(iPsecPolicy).ToNot(BeNil())

		})
		It(`UpdateIpsecPolicy request example`, func() {
			fmt.Println("\nUpdateIpsecPolicy() result:")
			// begin-update_ipsec_policy

			iPsecPolicyPatchModel := &vpcv1.IPsecPolicyPatch{}
			iPsecPolicyPatchModelAsPatch, asPatchErr := iPsecPolicyPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateIpsecPolicyOptions := vpcService.NewUpdateIpsecPolicyOptions(
				"testString",
				iPsecPolicyPatchModelAsPatch,
			)

			iPsecPolicy, response, err := vpcService.UpdateIpsecPolicy(updateIpsecPolicyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(iPsecPolicy, "", "  ")
			fmt.Println(string(b))

			// end-update_ipsec_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(iPsecPolicy).ToNot(BeNil())

		})
		It(`ListIpsecPolicyConnections request example`, func() {
			fmt.Println("\nListIpsecPolicyConnections() result:")
			// begin-list_ipsec_policy_connections

			listIpsecPolicyConnectionsOptions := vpcService.NewListIpsecPolicyConnectionsOptions(
				"testString",
			)

			vpnGatewayConnectionCollection, response, err := vpcService.ListIpsecPolicyConnections(listIpsecPolicyConnectionsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(vpnGatewayConnectionCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_ipsec_policy_connections

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnGatewayConnectionCollection).ToNot(BeNil())

		})
		It(`ListVPNGateways request example`, func() {
			fmt.Println("\nListVPNGateways() result:")
			// begin-list_vpn_gateways

			listVPNGatewaysOptions := vpcService.NewListVPNGatewaysOptions()
			listVPNGatewaysOptions.SetMode("route")

			vpnGatewayCollection, response, err := vpcService.ListVPNGateways(listVPNGatewaysOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(vpnGatewayCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_vpn_gateways

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnGatewayCollection).ToNot(BeNil())

		})
		It(`CreateVPNGateway request example`, func() {
			fmt.Println("\nCreateVPNGateway() result:")
			// begin-create_vpn_gateway

			subnetIdentityModel := &vpcv1.SubnetIdentityByID{
				ID: core.StringPtr("7ec86020-1c6e-4889-b3f0-a15f2e50f87e"),
			}

			vpnGatewayPrototypeModel := &vpcv1.VPNGatewayPrototypeVPNGatewayRouteModePrototype{
				Subnet: subnetIdentityModel,
			}

			createVPNGatewayOptions := vpcService.NewCreateVPNGatewayOptions(
				vpnGatewayPrototypeModel,
			)

			vpnGateway, response, err := vpcService.CreateVPNGateway(createVPNGatewayOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(vpnGateway, "", "  ")
			fmt.Println(string(b))

			// end-create_vpn_gateway

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(vpnGateway).ToNot(BeNil())

		})
		It(`GetVPNGateway request example`, func() {
			fmt.Println("\nGetVPNGateway() result:")
			// begin-get_vpn_gateway

			getVPNGatewayOptions := vpcService.NewGetVPNGatewayOptions(
				"testString",
			)

			vpnGateway, response, err := vpcService.GetVPNGateway(getVPNGatewayOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(vpnGateway, "", "  ")
			fmt.Println(string(b))

			// end-get_vpn_gateway

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnGateway).ToNot(BeNil())

		})
		It(`UpdateVPNGateway request example`, func() {
			fmt.Println("\nUpdateVPNGateway() result:")
			// begin-update_vpn_gateway

			vpnGatewayPatchModel := &vpcv1.VPNGatewayPatch{}
			vpnGatewayPatchModelAsPatch, asPatchErr := vpnGatewayPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateVPNGatewayOptions := vpcService.NewUpdateVPNGatewayOptions(
				"testString",
				vpnGatewayPatchModelAsPatch,
			)

			vpnGateway, response, err := vpcService.UpdateVPNGateway(updateVPNGatewayOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(vpnGateway, "", "  ")
			fmt.Println(string(b))

			// end-update_vpn_gateway

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnGateway).ToNot(BeNil())

		})
		It(`ListVPNGatewayConnections request example`, func() {
			fmt.Println("\nListVPNGatewayConnections() result:")
			// begin-list_vpn_gateway_connections

			listVPNGatewayConnectionsOptions := vpcService.NewListVPNGatewayConnectionsOptions(
				"testString",
			)

			vpnGatewayConnectionCollection, response, err := vpcService.ListVPNGatewayConnections(listVPNGatewayConnectionsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(vpnGatewayConnectionCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_vpn_gateway_connections

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnGatewayConnectionCollection).ToNot(BeNil())

		})
		It(`CreateVPNGatewayConnection request example`, func() {
			fmt.Println("\nCreateVPNGatewayConnection() result:")
			// begin-create_vpn_gateway_connection

			vpnGatewayConnectionPrototypeModel := &vpcv1.VPNGatewayConnectionPrototypeVPNGatewayConnectionStaticRouteModePrototype{
				PeerAddress: core.StringPtr("169.21.50.5"),
				Psk:         core.StringPtr("lkj14b1oi0alcniejkso"),
			}

			createVPNGatewayConnectionOptions := vpcService.NewCreateVPNGatewayConnectionOptions(
				"testString",
				vpnGatewayConnectionPrototypeModel,
			)

			vpnGatewayConnection, response, err := vpcService.CreateVPNGatewayConnection(createVPNGatewayConnectionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(vpnGatewayConnection, "", "  ")
			fmt.Println(string(b))

			// end-create_vpn_gateway_connection

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(vpnGatewayConnection).ToNot(BeNil())

		})
		It(`GetVPNGatewayConnection request example`, func() {
			fmt.Println("\nGetVPNGatewayConnection() result:")
			// begin-get_vpn_gateway_connection

			getVPNGatewayConnectionOptions := vpcService.NewGetVPNGatewayConnectionOptions(
				"testString",
				"testString",
			)

			vpnGatewayConnection, response, err := vpcService.GetVPNGatewayConnection(getVPNGatewayConnectionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(vpnGatewayConnection, "", "  ")
			fmt.Println(string(b))

			// end-get_vpn_gateway_connection

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnGatewayConnection).ToNot(BeNil())

		})
		It(`UpdateVPNGatewayConnection request example`, func() {
			fmt.Println("\nUpdateVPNGatewayConnection() result:")
			// begin-update_vpn_gateway_connection

			vpnGatewayConnectionPatchModel := &vpcv1.VPNGatewayConnectionPatchVPNGatewayConnectionStaticRouteModePatch{}
			vpnGatewayConnectionPatchModelAsPatch, asPatchErr := vpnGatewayConnectionPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateVPNGatewayConnectionOptions := vpcService.NewUpdateVPNGatewayConnectionOptions(
				"testString",
				"testString",
				vpnGatewayConnectionPatchModelAsPatch,
			)

			vpnGatewayConnection, response, err := vpcService.UpdateVPNGatewayConnection(updateVPNGatewayConnectionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(vpnGatewayConnection, "", "  ")
			fmt.Println(string(b))

			// end-update_vpn_gateway_connection

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnGatewayConnection).ToNot(BeNil())

		})
		It(`ListVPNGatewayConnectionLocalCIDRs request example`, func() {
			fmt.Println("\nListVPNGatewayConnectionLocalCIDRs() result:")
			// begin-list_vpn_gateway_connection_local_cidrs

			listVPNGatewayConnectionLocalCIDRsOptions := vpcService.NewListVPNGatewayConnectionLocalCIDRsOptions(
				"testString",
				"testString",
			)

			vpnGatewayConnectionLocalCIDRs, response, err := vpcService.ListVPNGatewayConnectionLocalCIDRs(listVPNGatewayConnectionLocalCIDRsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(vpnGatewayConnectionLocalCIDRs, "", "  ")
			fmt.Println(string(b))

			// end-list_vpn_gateway_connection_local_cidrs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnGatewayConnectionLocalCIDRs).ToNot(BeNil())

		})
		It(`CheckVPNGatewayConnectionLocalCIDR request example`, func() {
			// begin-check_vpn_gateway_connection_local_cidr

			checkVPNGatewayConnectionLocalCIDROptions := vpcService.NewCheckVPNGatewayConnectionLocalCIDROptions(
				"testString",
				"testString",
				"testString",
				"testString",
			)

			response, err := vpcService.CheckVPNGatewayConnectionLocalCIDR(checkVPNGatewayConnectionLocalCIDROptions)
			if err != nil {
				panic(err)
			}

			// end-check_vpn_gateway_connection_local_cidr
			fmt.Printf("\nCheckVPNGatewayConnectionLocalCIDR() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`AddVPNGatewayConnectionLocalCIDR request example`, func() {
			// begin-add_vpn_gateway_connection_local_cidr

			addVPNGatewayConnectionLocalCIDROptions := vpcService.NewAddVPNGatewayConnectionLocalCIDROptions(
				"testString",
				"testString",
				"testString",
				"testString",
			)

			response, err := vpcService.AddVPNGatewayConnectionLocalCIDR(addVPNGatewayConnectionLocalCIDROptions)
			if err != nil {
				panic(err)
			}

			// end-add_vpn_gateway_connection_local_cidr
			fmt.Printf("\nAddVPNGatewayConnectionLocalCIDR() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`ListVPNGatewayConnectionPeerCIDRs request example`, func() {
			fmt.Println("\nListVPNGatewayConnectionPeerCIDRs() result:")
			// begin-list_vpn_gateway_connection_peer_cidrs

			listVPNGatewayConnectionPeerCIDRsOptions := vpcService.NewListVPNGatewayConnectionPeerCIDRsOptions(
				"testString",
				"testString",
			)

			vpnGatewayConnectionPeerCIDRs, response, err := vpcService.ListVPNGatewayConnectionPeerCIDRs(listVPNGatewayConnectionPeerCIDRsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(vpnGatewayConnectionPeerCIDRs, "", "  ")
			fmt.Println(string(b))

			// end-list_vpn_gateway_connection_peer_cidrs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnGatewayConnectionPeerCIDRs).ToNot(BeNil())

		})
		It(`CheckVPNGatewayConnectionPeerCIDR request example`, func() {
			// begin-check_vpn_gateway_connection_peer_cidr

			checkVPNGatewayConnectionPeerCIDROptions := vpcService.NewCheckVPNGatewayConnectionPeerCIDROptions(
				"testString",
				"testString",
				"testString",
				"testString",
			)

			response, err := vpcService.CheckVPNGatewayConnectionPeerCIDR(checkVPNGatewayConnectionPeerCIDROptions)
			if err != nil {
				panic(err)
			}

			// end-check_vpn_gateway_connection_peer_cidr
			fmt.Printf("\nCheckVPNGatewayConnectionPeerCIDR() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`AddVPNGatewayConnectionPeerCIDR request example`, func() {
			// begin-add_vpn_gateway_connection_peer_cidr

			addVPNGatewayConnectionPeerCIDROptions := vpcService.NewAddVPNGatewayConnectionPeerCIDROptions(
				"testString",
				"testString",
				"testString",
				"testString",
			)

			response, err := vpcService.AddVPNGatewayConnectionPeerCIDR(addVPNGatewayConnectionPeerCIDROptions)
			if err != nil {
				panic(err)
			}

			// end-add_vpn_gateway_connection_peer_cidr
			fmt.Printf("\nAddVPNGatewayConnectionPeerCIDR() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`ListLoadBalancerProfiles request example`, func() {
			fmt.Println("\nListLoadBalancerProfiles() result:")
			// begin-list_load_balancer_profiles

			listLoadBalancerProfilesOptions := vpcService.NewListLoadBalancerProfilesOptions()

			loadBalancerProfileCollection, response, err := vpcService.ListLoadBalancerProfiles(listLoadBalancerProfilesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(loadBalancerProfileCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_load_balancer_profiles

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerProfileCollection).ToNot(BeNil())

		})
		It(`GetLoadBalancerProfile request example`, func() {
			fmt.Println("\nGetLoadBalancerProfile() result:")
			// begin-get_load_balancer_profile

			getLoadBalancerProfileOptions := vpcService.NewGetLoadBalancerProfileOptions(
				"testString",
			)

			loadBalancerProfile, response, err := vpcService.GetLoadBalancerProfile(getLoadBalancerProfileOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(loadBalancerProfile, "", "  ")
			fmt.Println(string(b))

			// end-get_load_balancer_profile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerProfile).ToNot(BeNil())

		})
		It(`ListLoadBalancers request example`, func() {
			fmt.Println("\nListLoadBalancers() result:")
			// begin-list_load_balancers

			listLoadBalancersOptions := vpcService.NewListLoadBalancersOptions()

			loadBalancerCollection, response, err := vpcService.ListLoadBalancers(listLoadBalancersOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(loadBalancerCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_load_balancers

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerCollection).ToNot(BeNil())

		})
		It(`CreateLoadBalancer request example`, func() {
			fmt.Println("\nCreateLoadBalancer() result:")
			// begin-create_load_balancer

			subnetIdentityModel := &vpcv1.SubnetIdentityByID{
				ID: core.StringPtr("7ec86020-1c6e-4889-b3f0-a15f2e50f87e"),
			}

			createLoadBalancerOptions := vpcService.NewCreateLoadBalancerOptions(
				true,
				[]vpcv1.SubnetIdentityIntf{subnetIdentityModel},
			)

			loadBalancer, response, err := vpcService.CreateLoadBalancer(createLoadBalancerOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(loadBalancer, "", "  ")
			fmt.Println(string(b))

			// end-create_load_balancer

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(loadBalancer).ToNot(BeNil())

		})
		It(`GetLoadBalancer request example`, func() {
			fmt.Println("\nGetLoadBalancer() result:")
			// begin-get_load_balancer

			getLoadBalancerOptions := vpcService.NewGetLoadBalancerOptions(
				"testString",
			)

			loadBalancer, response, err := vpcService.GetLoadBalancer(getLoadBalancerOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(loadBalancer, "", "  ")
			fmt.Println(string(b))

			// end-get_load_balancer

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancer).ToNot(BeNil())

		})
		It(`UpdateLoadBalancer request example`, func() {
			fmt.Println("\nUpdateLoadBalancer() result:")
			// begin-update_load_balancer

			loadBalancerPatchModel := &vpcv1.LoadBalancerPatch{}
			loadBalancerPatchModelAsPatch, asPatchErr := loadBalancerPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateLoadBalancerOptions := vpcService.NewUpdateLoadBalancerOptions(
				"testString",
				loadBalancerPatchModelAsPatch,
			)

			loadBalancer, response, err := vpcService.UpdateLoadBalancer(updateLoadBalancerOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(loadBalancer, "", "  ")
			fmt.Println(string(b))

			// end-update_load_balancer

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancer).ToNot(BeNil())

		})
		It(`GetLoadBalancerStatistics request example`, func() {
			fmt.Println("\nGetLoadBalancerStatistics() result:")
			// begin-get_load_balancer_statistics

			getLoadBalancerStatisticsOptions := vpcService.NewGetLoadBalancerStatisticsOptions(
				"testString",
			)

			loadBalancerStatistics, response, err := vpcService.GetLoadBalancerStatistics(getLoadBalancerStatisticsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(loadBalancerStatistics, "", "  ")
			fmt.Println(string(b))

			// end-get_load_balancer_statistics

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerStatistics).ToNot(BeNil())

		})
		It(`ListLoadBalancerListeners request example`, func() {
			fmt.Println("\nListLoadBalancerListeners() result:")
			// begin-list_load_balancer_listeners

			listLoadBalancerListenersOptions := vpcService.NewListLoadBalancerListenersOptions(
				"testString",
			)

			loadBalancerListenerCollection, response, err := vpcService.ListLoadBalancerListeners(listLoadBalancerListenersOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(loadBalancerListenerCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_load_balancer_listeners

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerListenerCollection).ToNot(BeNil())

		})
		It(`CreateLoadBalancerListener request example`, func() {
			fmt.Println("\nCreateLoadBalancerListener() result:")
			// begin-create_load_balancer_listener

			createLoadBalancerListenerOptions := vpcService.NewCreateLoadBalancerListenerOptions(
				"testString",
				int64(443),
				"http",
			)

			loadBalancerListener, response, err := vpcService.CreateLoadBalancerListener(createLoadBalancerListenerOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(loadBalancerListener, "", "  ")
			fmt.Println(string(b))

			// end-create_load_balancer_listener

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(loadBalancerListener).ToNot(BeNil())

		})
		It(`GetLoadBalancerListener request example`, func() {
			fmt.Println("\nGetLoadBalancerListener() result:")
			// begin-get_load_balancer_listener

			getLoadBalancerListenerOptions := vpcService.NewGetLoadBalancerListenerOptions(
				"testString",
				"testString",
			)

			loadBalancerListener, response, err := vpcService.GetLoadBalancerListener(getLoadBalancerListenerOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(loadBalancerListener, "", "  ")
			fmt.Println(string(b))

			// end-get_load_balancer_listener

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerListener).ToNot(BeNil())

		})
		It(`UpdateLoadBalancerListener request example`, func() {
			fmt.Println("\nUpdateLoadBalancerListener() result:")
			// begin-update_load_balancer_listener

			loadBalancerListenerPatchModel := &vpcv1.LoadBalancerListenerPatch{}
			loadBalancerListenerPatchModelAsPatch, asPatchErr := loadBalancerListenerPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateLoadBalancerListenerOptions := vpcService.NewUpdateLoadBalancerListenerOptions(
				"testString",
				"testString",
				loadBalancerListenerPatchModelAsPatch,
			)

			loadBalancerListener, response, err := vpcService.UpdateLoadBalancerListener(updateLoadBalancerListenerOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(loadBalancerListener, "", "  ")
			fmt.Println(string(b))

			// end-update_load_balancer_listener

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerListener).ToNot(BeNil())

		})
		It(`ListLoadBalancerListenerPolicies request example`, func() {
			fmt.Println("\nListLoadBalancerListenerPolicies() result:")
			// begin-list_load_balancer_listener_policies

			listLoadBalancerListenerPoliciesOptions := vpcService.NewListLoadBalancerListenerPoliciesOptions(
				"testString",
				"testString",
			)

			loadBalancerListenerPolicyCollection, response, err := vpcService.ListLoadBalancerListenerPolicies(listLoadBalancerListenerPoliciesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(loadBalancerListenerPolicyCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_load_balancer_listener_policies

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerListenerPolicyCollection).ToNot(BeNil())

		})
		It(`CreateLoadBalancerListenerPolicy request example`, func() {
			fmt.Println("\nCreateLoadBalancerListenerPolicy() result:")
			// begin-create_load_balancer_listener_policy

			createLoadBalancerListenerPolicyOptions := vpcService.NewCreateLoadBalancerListenerPolicyOptions(
				"testString",
				"testString",
				"forward",
				int64(5),
			)

			loadBalancerListenerPolicy, response, err := vpcService.CreateLoadBalancerListenerPolicy(createLoadBalancerListenerPolicyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(loadBalancerListenerPolicy, "", "  ")
			fmt.Println(string(b))

			// end-create_load_balancer_listener_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(loadBalancerListenerPolicy).ToNot(BeNil())

		})
		It(`GetLoadBalancerListenerPolicy request example`, func() {
			fmt.Println("\nGetLoadBalancerListenerPolicy() result:")
			// begin-get_load_balancer_listener_policy

			getLoadBalancerListenerPolicyOptions := vpcService.NewGetLoadBalancerListenerPolicyOptions(
				"testString",
				"testString",
				"testString",
			)

			loadBalancerListenerPolicy, response, err := vpcService.GetLoadBalancerListenerPolicy(getLoadBalancerListenerPolicyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(loadBalancerListenerPolicy, "", "  ")
			fmt.Println(string(b))

			// end-get_load_balancer_listener_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerListenerPolicy).ToNot(BeNil())

		})
		It(`UpdateLoadBalancerListenerPolicy request example`, func() {
			fmt.Println("\nUpdateLoadBalancerListenerPolicy() result:")
			// begin-update_load_balancer_listener_policy

			loadBalancerListenerPolicyPatchModel := &vpcv1.LoadBalancerListenerPolicyPatch{}
			loadBalancerListenerPolicyPatchModelAsPatch, asPatchErr := loadBalancerListenerPolicyPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateLoadBalancerListenerPolicyOptions := vpcService.NewUpdateLoadBalancerListenerPolicyOptions(
				"testString",
				"testString",
				"testString",
				loadBalancerListenerPolicyPatchModelAsPatch,
			)

			loadBalancerListenerPolicy, response, err := vpcService.UpdateLoadBalancerListenerPolicy(updateLoadBalancerListenerPolicyOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(loadBalancerListenerPolicy, "", "  ")
			fmt.Println(string(b))

			// end-update_load_balancer_listener_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerListenerPolicy).ToNot(BeNil())

		})
		It(`ListLoadBalancerListenerPolicyRules request example`, func() {
			fmt.Println("\nListLoadBalancerListenerPolicyRules() result:")
			// begin-list_load_balancer_listener_policy_rules

			listLoadBalancerListenerPolicyRulesOptions := vpcService.NewListLoadBalancerListenerPolicyRulesOptions(
				"testString",
				"testString",
				"testString",
			)

			loadBalancerListenerPolicyRuleCollection, response, err := vpcService.ListLoadBalancerListenerPolicyRules(listLoadBalancerListenerPolicyRulesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(loadBalancerListenerPolicyRuleCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_load_balancer_listener_policy_rules

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerListenerPolicyRuleCollection).ToNot(BeNil())

		})
		It(`CreateLoadBalancerListenerPolicyRule request example`, func() {
			fmt.Println("\nCreateLoadBalancerListenerPolicyRule() result:")
			// begin-create_load_balancer_listener_policy_rule

			createLoadBalancerListenerPolicyRuleOptions := vpcService.NewCreateLoadBalancerListenerPolicyRuleOptions(
				"testString",
				"testString",
				"testString",
				"contains",
				"header",
				"testString",
			)

			loadBalancerListenerPolicyRule, response, err := vpcService.CreateLoadBalancerListenerPolicyRule(createLoadBalancerListenerPolicyRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(loadBalancerListenerPolicyRule, "", "  ")
			fmt.Println(string(b))

			// end-create_load_balancer_listener_policy_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(loadBalancerListenerPolicyRule).ToNot(BeNil())

		})
		It(`GetLoadBalancerListenerPolicyRule request example`, func() {
			fmt.Println("\nGetLoadBalancerListenerPolicyRule() result:")
			// begin-get_load_balancer_listener_policy_rule

			getLoadBalancerListenerPolicyRuleOptions := vpcService.NewGetLoadBalancerListenerPolicyRuleOptions(
				"testString",
				"testString",
				"testString",
				"testString",
			)

			loadBalancerListenerPolicyRule, response, err := vpcService.GetLoadBalancerListenerPolicyRule(getLoadBalancerListenerPolicyRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(loadBalancerListenerPolicyRule, "", "  ")
			fmt.Println(string(b))

			// end-get_load_balancer_listener_policy_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerListenerPolicyRule).ToNot(BeNil())

		})
		It(`UpdateLoadBalancerListenerPolicyRule request example`, func() {
			fmt.Println("\nUpdateLoadBalancerListenerPolicyRule() result:")
			// begin-update_load_balancer_listener_policy_rule

			loadBalancerListenerPolicyRulePatchModel := &vpcv1.LoadBalancerListenerPolicyRulePatch{}
			loadBalancerListenerPolicyRulePatchModelAsPatch, asPatchErr := loadBalancerListenerPolicyRulePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateLoadBalancerListenerPolicyRuleOptions := vpcService.NewUpdateLoadBalancerListenerPolicyRuleOptions(
				"testString",
				"testString",
				"testString",
				"testString",
				loadBalancerListenerPolicyRulePatchModelAsPatch,
			)

			loadBalancerListenerPolicyRule, response, err := vpcService.UpdateLoadBalancerListenerPolicyRule(updateLoadBalancerListenerPolicyRuleOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(loadBalancerListenerPolicyRule, "", "  ")
			fmt.Println(string(b))

			// end-update_load_balancer_listener_policy_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerListenerPolicyRule).ToNot(BeNil())

		})
		It(`ListLoadBalancerPools request example`, func() {
			fmt.Println("\nListLoadBalancerPools() result:")
			// begin-list_load_balancer_pools

			listLoadBalancerPoolsOptions := vpcService.NewListLoadBalancerPoolsOptions(
				"testString",
			)

			loadBalancerPoolCollection, response, err := vpcService.ListLoadBalancerPools(listLoadBalancerPoolsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(loadBalancerPoolCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_load_balancer_pools

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerPoolCollection).ToNot(BeNil())

		})
		It(`CreateLoadBalancerPool request example`, func() {
			fmt.Println("\nCreateLoadBalancerPool() result:")
			// begin-create_load_balancer_pool

			loadBalancerPoolHealthMonitorPrototypeModel := &vpcv1.LoadBalancerPoolHealthMonitorPrototype{
				Delay:      core.Int64Ptr(int64(5)),
				MaxRetries: core.Int64Ptr(int64(2)),
				Timeout:    core.Int64Ptr(int64(2)),
				Type:       core.StringPtr("http"),
			}

			createLoadBalancerPoolOptions := vpcService.NewCreateLoadBalancerPoolOptions(
				"testString",
				"least_connections",
				loadBalancerPoolHealthMonitorPrototypeModel,
				"http",
			)

			loadBalancerPool, response, err := vpcService.CreateLoadBalancerPool(createLoadBalancerPoolOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(loadBalancerPool, "", "  ")
			fmt.Println(string(b))

			// end-create_load_balancer_pool

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(loadBalancerPool).ToNot(BeNil())

		})
		It(`GetLoadBalancerPool request example`, func() {
			fmt.Println("\nGetLoadBalancerPool() result:")
			// begin-get_load_balancer_pool

			getLoadBalancerPoolOptions := vpcService.NewGetLoadBalancerPoolOptions(
				"testString",
				"testString",
			)

			loadBalancerPool, response, err := vpcService.GetLoadBalancerPool(getLoadBalancerPoolOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(loadBalancerPool, "", "  ")
			fmt.Println(string(b))

			// end-get_load_balancer_pool

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerPool).ToNot(BeNil())

		})
		It(`UpdateLoadBalancerPool request example`, func() {
			fmt.Println("\nUpdateLoadBalancerPool() result:")
			// begin-update_load_balancer_pool

			loadBalancerPoolPatchModel := &vpcv1.LoadBalancerPoolPatch{}
			loadBalancerPoolPatchModelAsPatch, asPatchErr := loadBalancerPoolPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateLoadBalancerPoolOptions := vpcService.NewUpdateLoadBalancerPoolOptions(
				"testString",
				"testString",
				loadBalancerPoolPatchModelAsPatch,
			)

			loadBalancerPool, response, err := vpcService.UpdateLoadBalancerPool(updateLoadBalancerPoolOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(loadBalancerPool, "", "  ")
			fmt.Println(string(b))

			// end-update_load_balancer_pool

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerPool).ToNot(BeNil())

		})
		It(`ListLoadBalancerPoolMembers request example`, func() {
			fmt.Println("\nListLoadBalancerPoolMembers() result:")
			// begin-list_load_balancer_pool_members

			listLoadBalancerPoolMembersOptions := vpcService.NewListLoadBalancerPoolMembersOptions(
				"testString",
				"testString",
			)

			loadBalancerPoolMemberCollection, response, err := vpcService.ListLoadBalancerPoolMembers(listLoadBalancerPoolMembersOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(loadBalancerPoolMemberCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_load_balancer_pool_members

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerPoolMemberCollection).ToNot(BeNil())

		})
		It(`CreateLoadBalancerPoolMember request example`, func() {
			fmt.Println("\nCreateLoadBalancerPoolMember() result:")
			// begin-create_load_balancer_pool_member

			loadBalancerPoolMemberTargetPrototypeModel := &vpcv1.LoadBalancerPoolMemberTargetPrototypeInstanceIdentityInstanceIdentityByID{
				ID: core.StringPtr("1e09281b-f177-46fb-baf1-bc152b2e391a"),
			}

			createLoadBalancerPoolMemberOptions := vpcService.NewCreateLoadBalancerPoolMemberOptions(
				"testString",
				"testString",
				int64(80),
				loadBalancerPoolMemberTargetPrototypeModel,
			)

			loadBalancerPoolMember, response, err := vpcService.CreateLoadBalancerPoolMember(createLoadBalancerPoolMemberOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(loadBalancerPoolMember, "", "  ")
			fmt.Println(string(b))

			// end-create_load_balancer_pool_member

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(loadBalancerPoolMember).ToNot(BeNil())

		})
		It(`ReplaceLoadBalancerPoolMembers request example`, func() {
			fmt.Println("\nReplaceLoadBalancerPoolMembers() result:")
			// begin-replace_load_balancer_pool_members

			loadBalancerPoolMemberTargetPrototypeModel := &vpcv1.LoadBalancerPoolMemberTargetPrototypeInstanceIdentityInstanceIdentityByID{
				ID: core.StringPtr("1e09281b-f177-46fb-baf1-bc152b2e391a"),
			}

			loadBalancerPoolMemberPrototypeModel := &vpcv1.LoadBalancerPoolMemberPrototype{
				Port:   core.Int64Ptr(int64(80)),
				Target: loadBalancerPoolMemberTargetPrototypeModel,
			}

			replaceLoadBalancerPoolMembersOptions := vpcService.NewReplaceLoadBalancerPoolMembersOptions(
				"testString",
				"testString",
				[]vpcv1.LoadBalancerPoolMemberPrototype{*loadBalancerPoolMemberPrototypeModel},
			)

			loadBalancerPoolMemberCollection, response, err := vpcService.ReplaceLoadBalancerPoolMembers(replaceLoadBalancerPoolMembersOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(loadBalancerPoolMemberCollection, "", "  ")
			fmt.Println(string(b))

			// end-replace_load_balancer_pool_members

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(loadBalancerPoolMemberCollection).ToNot(BeNil())

		})
		It(`GetLoadBalancerPoolMember request example`, func() {
			fmt.Println("\nGetLoadBalancerPoolMember() result:")
			// begin-get_load_balancer_pool_member

			getLoadBalancerPoolMemberOptions := vpcService.NewGetLoadBalancerPoolMemberOptions(
				"testString",
				"testString",
				"testString",
			)

			loadBalancerPoolMember, response, err := vpcService.GetLoadBalancerPoolMember(getLoadBalancerPoolMemberOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(loadBalancerPoolMember, "", "  ")
			fmt.Println(string(b))

			// end-get_load_balancer_pool_member

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerPoolMember).ToNot(BeNil())

		})
		It(`UpdateLoadBalancerPoolMember request example`, func() {
			fmt.Println("\nUpdateLoadBalancerPoolMember() result:")
			// begin-update_load_balancer_pool_member

			loadBalancerPoolMemberPatchModel := &vpcv1.LoadBalancerPoolMemberPatch{}
			loadBalancerPoolMemberPatchModelAsPatch, asPatchErr := loadBalancerPoolMemberPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateLoadBalancerPoolMemberOptions := vpcService.NewUpdateLoadBalancerPoolMemberOptions(
				"testString",
				"testString",
				"testString",
				loadBalancerPoolMemberPatchModelAsPatch,
			)

			loadBalancerPoolMember, response, err := vpcService.UpdateLoadBalancerPoolMember(updateLoadBalancerPoolMemberOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(loadBalancerPoolMember, "", "  ")
			fmt.Println(string(b))

			// end-update_load_balancer_pool_member

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancerPoolMember).ToNot(BeNil())

		})
		It(`ListEndpointGateways request example`, func() {
			fmt.Println("\nListEndpointGateways() result:")
			// begin-list_endpoint_gateways

			listEndpointGatewaysOptions := vpcService.NewListEndpointGatewaysOptions()

			endpointGatewayCollection, response, err := vpcService.ListEndpointGateways(listEndpointGatewaysOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(endpointGatewayCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_endpoint_gateways

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(endpointGatewayCollection).ToNot(BeNil())

		})
		It(`CreateEndpointGateway request example`, func() {
			fmt.Println("\nCreateEndpointGateway() result:")
			// begin-create_endpoint_gateway

			endpointGatewayTargetPrototypeModel := &vpcv1.EndpointGatewayTargetPrototypeProviderCloudServiceIdentityProviderCloudServiceIdentityByCRN{
				ResourceType: core.StringPtr("provider_infrastructure_service"),
				CRN:          core.StringPtr("crn:v1:bluemix:public:cloudant:us-south:a/123456:3527280b-9327-4411-8020-591092e60353::"),
			}

			vpcIdentityModel := &vpcv1.VPCIdentityByID{
				ID: core.StringPtr("f025b503-ae66-46de-a011-3bd08fd5f7bf"),
			}

			createEndpointGatewayOptions := vpcService.NewCreateEndpointGatewayOptions(
				endpointGatewayTargetPrototypeModel,
				vpcIdentityModel,
			)

			endpointGateway, response, err := vpcService.CreateEndpointGateway(createEndpointGatewayOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(endpointGateway, "", "  ")
			fmt.Println(string(b))

			// end-create_endpoint_gateway

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(endpointGateway).ToNot(BeNil())

		})
		It(`ListEndpointGatewayIps request example`, func() {
			fmt.Println("\nListEndpointGatewayIps() result:")
			// begin-list_endpoint_gateway_ips

			listEndpointGatewayIpsOptions := vpcService.NewListEndpointGatewayIpsOptions(
				"testString",
			)
			listEndpointGatewayIpsOptions.SetSort("name")

			reservedIPCollectionEndpointGatewayContext, response, err := vpcService.ListEndpointGatewayIps(listEndpointGatewayIpsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(reservedIPCollectionEndpointGatewayContext, "", "  ")
			fmt.Println(string(b))

			// end-list_endpoint_gateway_ips

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reservedIPCollectionEndpointGatewayContext).ToNot(BeNil())

		})
		It(`GetEndpointGatewayIP request example`, func() {
			fmt.Println("\nGetEndpointGatewayIP() result:")
			// begin-get_endpoint_gateway_ip

			getEndpointGatewayIPOptions := vpcService.NewGetEndpointGatewayIPOptions(
				"testString",
				"testString",
			)

			reservedIP, response, err := vpcService.GetEndpointGatewayIP(getEndpointGatewayIPOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(reservedIP, "", "  ")
			fmt.Println(string(b))

			// end-get_endpoint_gateway_ip

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reservedIP).ToNot(BeNil())

		})
		It(`AddEndpointGatewayIP request example`, func() {
			fmt.Println("\nAddEndpointGatewayIP() result:")
			// begin-add_endpoint_gateway_ip

			addEndpointGatewayIPOptions := vpcService.NewAddEndpointGatewayIPOptions(
				"testString",
				"testString",
			)

			reservedIP, response, err := vpcService.AddEndpointGatewayIP(addEndpointGatewayIPOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(reservedIP, "", "  ")
			fmt.Println(string(b))

			// end-add_endpoint_gateway_ip

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(reservedIP).ToNot(BeNil())

		})
		It(`GetEndpointGateway request example`, func() {
			fmt.Println("\nGetEndpointGateway() result:")
			// begin-get_endpoint_gateway

			getEndpointGatewayOptions := vpcService.NewGetEndpointGatewayOptions(
				"testString",
			)

			endpointGateway, response, err := vpcService.GetEndpointGateway(getEndpointGatewayOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(endpointGateway, "", "  ")
			fmt.Println(string(b))

			// end-get_endpoint_gateway

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(endpointGateway).ToNot(BeNil())

		})
		It(`UpdateEndpointGateway request example`, func() {
			fmt.Println("\nUpdateEndpointGateway() result:")
			// begin-update_endpoint_gateway

			endpointGatewayPatchModel := &vpcv1.EndpointGatewayPatch{}
			endpointGatewayPatchModelAsPatch, asPatchErr := endpointGatewayPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateEndpointGatewayOptions := vpcService.NewUpdateEndpointGatewayOptions(
				"testString",
				endpointGatewayPatchModelAsPatch,
			)

			endpointGateway, response, err := vpcService.UpdateEndpointGateway(updateEndpointGatewayOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(endpointGateway, "", "  ")
			fmt.Println(string(b))

			// end-update_endpoint_gateway

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(endpointGateway).ToNot(BeNil())

		})
		It(`ListFlowLogCollectors request example`, func() {
			fmt.Println("\nListFlowLogCollectors() result:")
			// begin-list_flow_log_collectors

			listFlowLogCollectorsOptions := vpcService.NewListFlowLogCollectorsOptions()

			flowLogCollectorCollection, response, err := vpcService.ListFlowLogCollectors(listFlowLogCollectorsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(flowLogCollectorCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_flow_log_collectors

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(flowLogCollectorCollection).ToNot(BeNil())

		})
		It(`CreateFlowLogCollector request example`, func() {
			fmt.Println("\nCreateFlowLogCollector() result:")
			// begin-create_flow_log_collector

			cloudObjectStorageBucketIdentityModel := &vpcv1.CloudObjectStorageBucketIdentityByName{
				Name: core.StringPtr("bucket-27200-lwx4cfvcue"),
			}

			flowLogCollectorTargetPrototypeModel := &vpcv1.FlowLogCollectorTargetPrototypeNetworkInterfaceIdentityNetworkInterfaceIdentityNetworkInterfaceIdentityByID{
				ID: core.StringPtr("10c02d81-0ecb-4dc5-897d-28392913b81e"),
			}

			createFlowLogCollectorOptions := vpcService.NewCreateFlowLogCollectorOptions(
				cloudObjectStorageBucketIdentityModel,
				flowLogCollectorTargetPrototypeModel,
			)

			flowLogCollector, response, err := vpcService.CreateFlowLogCollector(createFlowLogCollectorOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(flowLogCollector, "", "  ")
			fmt.Println(string(b))

			// end-create_flow_log_collector

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(flowLogCollector).ToNot(BeNil())

		})
		It(`GetFlowLogCollector request example`, func() {
			fmt.Println("\nGetFlowLogCollector() result:")
			// begin-get_flow_log_collector

			getFlowLogCollectorOptions := vpcService.NewGetFlowLogCollectorOptions(
				"testString",
			)

			flowLogCollector, response, err := vpcService.GetFlowLogCollector(getFlowLogCollectorOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(flowLogCollector, "", "  ")
			fmt.Println(string(b))

			// end-get_flow_log_collector

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(flowLogCollector).ToNot(BeNil())

		})
		It(`UpdateFlowLogCollector request example`, func() {
			fmt.Println("\nUpdateFlowLogCollector() result:")
			// begin-update_flow_log_collector

			flowLogCollectorPatchModel := &vpcv1.FlowLogCollectorPatch{}
			flowLogCollectorPatchModelAsPatch, asPatchErr := flowLogCollectorPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateFlowLogCollectorOptions := vpcService.NewUpdateFlowLogCollectorOptions(
				"testString",
				flowLogCollectorPatchModelAsPatch,
			)

			flowLogCollector, response, err := vpcService.UpdateFlowLogCollector(updateFlowLogCollectorOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(flowLogCollector, "", "  ")
			fmt.Println(string(b))

			// end-update_flow_log_collector

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(flowLogCollector).ToNot(BeNil())

		})
		It(`UnsetSubnetPublicGateway request example`, func() {
			// begin-unset_subnet_public_gateway

			unsetSubnetPublicGatewayOptions := vpcService.NewUnsetSubnetPublicGatewayOptions(
				"testString",
			)

			response, err := vpcService.UnsetSubnetPublicGateway(unsetSubnetPublicGatewayOptions)
			if err != nil {
				panic(err)
			}

			// end-unset_subnet_public_gateway
			fmt.Printf("\nUnsetSubnetPublicGateway() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`RemoveVPNGatewayConnectionPeerCIDR request example`, func() {
			// begin-remove_vpn_gateway_connection_peer_cidr

			removeVPNGatewayConnectionPeerCIDROptions := vpcService.NewRemoveVPNGatewayConnectionPeerCIDROptions(
				"testString",
				"testString",
				"testString",
				"testString",
			)

			response, err := vpcService.RemoveVPNGatewayConnectionPeerCIDR(removeVPNGatewayConnectionPeerCIDROptions)
			if err != nil {
				panic(err)
			}

			// end-remove_vpn_gateway_connection_peer_cidr
			fmt.Printf("\nRemoveVPNGatewayConnectionPeerCIDR() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`RemoveVPNGatewayConnectionLocalCIDR request example`, func() {
			// begin-remove_vpn_gateway_connection_local_cidr

			removeVPNGatewayConnectionLocalCIDROptions := vpcService.NewRemoveVPNGatewayConnectionLocalCIDROptions(
				"testString",
				"testString",
				"testString",
				"testString",
			)

			response, err := vpcService.RemoveVPNGatewayConnectionLocalCIDR(removeVPNGatewayConnectionLocalCIDROptions)
			if err != nil {
				panic(err)
			}

			// end-remove_vpn_gateway_connection_local_cidr
			fmt.Printf("\nRemoveVPNGatewayConnectionLocalCIDR() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`RemoveSecurityGroupNetworkInterface request example`, func() {
			// begin-remove_security_group_network_interface

			removeSecurityGroupNetworkInterfaceOptions := vpcService.NewRemoveSecurityGroupNetworkInterfaceOptions(
				"testString",
				"testString",
			)

			response, err := vpcService.RemoveSecurityGroupNetworkInterface(removeSecurityGroupNetworkInterfaceOptions)
			if err != nil {
				panic(err)
			}

			// end-remove_security_group_network_interface
			fmt.Printf("\nRemoveSecurityGroupNetworkInterface() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`RemoveInstanceNetworkInterfaceFloatingIP request example`, func() {
			// begin-remove_instance_network_interface_floating_ip

			removeInstanceNetworkInterfaceFloatingIPOptions := vpcService.NewRemoveInstanceNetworkInterfaceFloatingIPOptions(
				"testString",
				"testString",
				"testString",
			)

			response, err := vpcService.RemoveInstanceNetworkInterfaceFloatingIP(removeInstanceNetworkInterfaceFloatingIPOptions)
			if err != nil {
				panic(err)
			}

			// end-remove_instance_network_interface_floating_ip
			fmt.Printf("\nRemoveInstanceNetworkInterfaceFloatingIP() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`RemoveEndpointGatewayIP request example`, func() {
			// begin-remove_endpoint_gateway_ip

			removeEndpointGatewayIPOptions := vpcService.NewRemoveEndpointGatewayIPOptions(
				"testString",
				"testString",
			)

			response, err := vpcService.RemoveEndpointGatewayIP(removeEndpointGatewayIPOptions)
			if err != nil {
				panic(err)
			}

			// end-remove_endpoint_gateway_ip
			fmt.Printf("\nRemoveEndpointGatewayIP() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteVolume request example`, func() {
			// begin-delete_volume

			deleteVolumeOptions := vpcService.NewDeleteVolumeOptions(
				"testString",
			)

			response, err := vpcService.DeleteVolume(deleteVolumeOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_volume
			fmt.Printf("\nDeleteVolume() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteVPNGatewayConnection request example`, func() {
			// begin-delete_vpn_gateway_connection

			deleteVPNGatewayConnectionOptions := vpcService.NewDeleteVPNGatewayConnectionOptions(
				"testString",
				"testString",
			)

			response, err := vpcService.DeleteVPNGatewayConnection(deleteVPNGatewayConnectionOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_vpn_gateway_connection
			fmt.Printf("\nDeleteVPNGatewayConnection() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})
		It(`DeleteVPNGateway request example`, func() {
			// begin-delete_vpn_gateway

			deleteVPNGatewayOptions := vpcService.NewDeleteVPNGatewayOptions(
				"testString",
			)

			response, err := vpcService.DeleteVPNGateway(deleteVPNGatewayOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_vpn_gateway
			fmt.Printf("\nDeleteVPNGateway() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})
		It(`DeleteVPCRoutingTableRoute request example`, func() {
			// begin-delete_vpc_routing_table_route

			deleteVPCRoutingTableRouteOptions := vpcService.NewDeleteVPCRoutingTableRouteOptions(
				"testString",
				"testString",
				"testString",
			)

			response, err := vpcService.DeleteVPCRoutingTableRoute(deleteVPCRoutingTableRouteOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_vpc_routing_table_route
			fmt.Printf("\nDeleteVPCRoutingTableRoute() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteVPCRoutingTable request example`, func() {
			// begin-delete_vpc_routing_table

			deleteVPCRoutingTableOptions := vpcService.NewDeleteVPCRoutingTableOptions(
				"testString",
				"testString",
			)

			response, err := vpcService.DeleteVPCRoutingTable(deleteVPCRoutingTableOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_vpc_routing_table
			fmt.Printf("\nDeleteVPCRoutingTable() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteVPCRoute request example`, func() {
			// begin-delete_vpc_route

			deleteVPCRouteOptions := vpcService.NewDeleteVPCRouteOptions(
				"testString",
				"testString",
			)

			response, err := vpcService.DeleteVPCRoute(deleteVPCRouteOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_vpc_route
			fmt.Printf("\nDeleteVPCRoute() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteVPCAddressPrefix request example`, func() {
			// begin-delete_vpc_address_prefix

			deleteVPCAddressPrefixOptions := vpcService.NewDeleteVPCAddressPrefixOptions(
				"testString",
				"testString",
			)

			response, err := vpcService.DeleteVPCAddressPrefix(deleteVPCAddressPrefixOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_vpc_address_prefix
			fmt.Printf("\nDeleteVPCAddressPrefix() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteVPC request example`, func() {
			// begin-delete_vpc

			deleteVPCOptions := vpcService.NewDeleteVPCOptions(
				"testString",
			)

			response, err := vpcService.DeleteVPC(deleteVPCOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_vpc
			fmt.Printf("\nDeleteVPC() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteSubnetReservedIP request example`, func() {
			// begin-delete_subnet_reserved_ip

			deleteSubnetReservedIPOptions := vpcService.NewDeleteSubnetReservedIPOptions(
				"testString",
				"testString",
			)

			response, err := vpcService.DeleteSubnetReservedIP(deleteSubnetReservedIPOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_subnet_reserved_ip
			fmt.Printf("\nDeleteSubnetReservedIP() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteSubnet request example`, func() {
			// begin-delete_subnet

			deleteSubnetOptions := vpcService.NewDeleteSubnetOptions(
				"testString",
			)

			response, err := vpcService.DeleteSubnet(deleteSubnetOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_subnet
			fmt.Printf("\nDeleteSubnet() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteSnapshots request example`, func() {
			// begin-delete_snapshots

			deleteSnapshotsOptions := vpcService.NewDeleteSnapshotsOptions(
				"testString",
			)

			response, err := vpcService.DeleteSnapshots(deleteSnapshotsOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_snapshots
			fmt.Printf("\nDeleteSnapshots() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteSnapshot request example`, func() {
			// begin-delete_snapshot

			deleteSnapshotOptions := vpcService.NewDeleteSnapshotOptions(
				"testString",
			)

			response, err := vpcService.DeleteSnapshot(deleteSnapshotOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_snapshot
			fmt.Printf("\nDeleteSnapshot() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteSecurityGroupTargetBinding request example`, func() {
			// begin-delete_security_group_target_binding

			deleteSecurityGroupTargetBindingOptions := vpcService.NewDeleteSecurityGroupTargetBindingOptions(
				"testString",
				"testString",
			)

			response, err := vpcService.DeleteSecurityGroupTargetBinding(deleteSecurityGroupTargetBindingOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_security_group_target_binding
			fmt.Printf("\nDeleteSecurityGroupTargetBinding() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteSecurityGroupRule request example`, func() {
			// begin-delete_security_group_rule

			deleteSecurityGroupRuleOptions := vpcService.NewDeleteSecurityGroupRuleOptions(
				"testString",
				"testString",
			)

			response, err := vpcService.DeleteSecurityGroupRule(deleteSecurityGroupRuleOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_security_group_rule
			fmt.Printf("\nDeleteSecurityGroupRule() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteSecurityGroup request example`, func() {
			// begin-delete_security_group

			deleteSecurityGroupOptions := vpcService.NewDeleteSecurityGroupOptions(
				"testString",
			)

			response, err := vpcService.DeleteSecurityGroup(deleteSecurityGroupOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_security_group
			fmt.Printf("\nDeleteSecurityGroup() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeletePublicGateway request example`, func() {
			// begin-delete_public_gateway

			deletePublicGatewayOptions := vpcService.NewDeletePublicGatewayOptions(
				"testString",
			)

			response, err := vpcService.DeletePublicGateway(deletePublicGatewayOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_public_gateway
			fmt.Printf("\nDeletePublicGateway() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteNetworkACLRule request example`, func() {
			// begin-delete_network_acl_rule

			deleteNetworkACLRuleOptions := vpcService.NewDeleteNetworkACLRuleOptions(
				"testString",
				"testString",
			)

			response, err := vpcService.DeleteNetworkACLRule(deleteNetworkACLRuleOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_network_acl_rule
			fmt.Printf("\nDeleteNetworkACLRule() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteNetworkACL request example`, func() {
			// begin-delete_network_acl

			deleteNetworkACLOptions := vpcService.NewDeleteNetworkACLOptions(
				"testString",
			)

			response, err := vpcService.DeleteNetworkACL(deleteNetworkACLOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_network_acl
			fmt.Printf("\nDeleteNetworkACL() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteLoadBalancerPoolMember request example`, func() {
			// begin-delete_load_balancer_pool_member

			deleteLoadBalancerPoolMemberOptions := vpcService.NewDeleteLoadBalancerPoolMemberOptions(
				"testString",
				"testString",
				"testString",
			)

			response, err := vpcService.DeleteLoadBalancerPoolMember(deleteLoadBalancerPoolMemberOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_load_balancer_pool_member
			fmt.Printf("\nDeleteLoadBalancerPoolMember() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteLoadBalancerPool request example`, func() {
			// begin-delete_load_balancer_pool

			deleteLoadBalancerPoolOptions := vpcService.NewDeleteLoadBalancerPoolOptions(
				"testString",
				"testString",
			)

			response, err := vpcService.DeleteLoadBalancerPool(deleteLoadBalancerPoolOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_load_balancer_pool
			fmt.Printf("\nDeleteLoadBalancerPool() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteLoadBalancerListenerPolicyRule request example`, func() {
			// begin-delete_load_balancer_listener_policy_rule

			deleteLoadBalancerListenerPolicyRuleOptions := vpcService.NewDeleteLoadBalancerListenerPolicyRuleOptions(
				"testString",
				"testString",
				"testString",
				"testString",
			)

			response, err := vpcService.DeleteLoadBalancerListenerPolicyRule(deleteLoadBalancerListenerPolicyRuleOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_load_balancer_listener_policy_rule
			fmt.Printf("\nDeleteLoadBalancerListenerPolicyRule() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteLoadBalancerListenerPolicy request example`, func() {
			// begin-delete_load_balancer_listener_policy

			deleteLoadBalancerListenerPolicyOptions := vpcService.NewDeleteLoadBalancerListenerPolicyOptions(
				"testString",
				"testString",
				"testString",
			)

			response, err := vpcService.DeleteLoadBalancerListenerPolicy(deleteLoadBalancerListenerPolicyOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_load_balancer_listener_policy
			fmt.Printf("\nDeleteLoadBalancerListenerPolicy() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteLoadBalancerListener request example`, func() {
			// begin-delete_load_balancer_listener

			deleteLoadBalancerListenerOptions := vpcService.NewDeleteLoadBalancerListenerOptions(
				"testString",
				"testString",
			)

			response, err := vpcService.DeleteLoadBalancerListener(deleteLoadBalancerListenerOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_load_balancer_listener
			fmt.Printf("\nDeleteLoadBalancerListener() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteLoadBalancer request example`, func() {
			// begin-delete_load_balancer

			deleteLoadBalancerOptions := vpcService.NewDeleteLoadBalancerOptions(
				"testString",
			)

			response, err := vpcService.DeleteLoadBalancer(deleteLoadBalancerOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_load_balancer
			fmt.Printf("\nDeleteLoadBalancer() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteKey request example`, func() {
			// begin-delete_key

			deleteKeyOptions := vpcService.NewDeleteKeyOptions(
				"testString",
			)

			response, err := vpcService.DeleteKey(deleteKeyOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_key
			fmt.Printf("\nDeleteKey() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteIpsecPolicy request example`, func() {
			// begin-delete_ipsec_policy

			deleteIpsecPolicyOptions := vpcService.NewDeleteIpsecPolicyOptions(
				"testString",
			)

			response, err := vpcService.DeleteIpsecPolicy(deleteIpsecPolicyOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_ipsec_policy
			fmt.Printf("\nDeleteIpsecPolicy() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteInstanceVolumeAttachment request example`, func() {
			// begin-delete_instance_volume_attachment

			deleteInstanceVolumeAttachmentOptions := vpcService.NewDeleteInstanceVolumeAttachmentOptions(
				"testString",
				"testString",
			)

			response, err := vpcService.DeleteInstanceVolumeAttachment(deleteInstanceVolumeAttachmentOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_instance_volume_attachment
			fmt.Printf("\nDeleteInstanceVolumeAttachment() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteInstanceTemplate request example`, func() {
			// begin-delete_instance_template

			deleteInstanceTemplateOptions := vpcService.NewDeleteInstanceTemplateOptions(
				"testString",
			)

			response, err := vpcService.DeleteInstanceTemplate(deleteInstanceTemplateOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_instance_template
			fmt.Printf("\nDeleteInstanceTemplate() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteInstanceNetworkInterface request example`, func() {
			// begin-delete_instance_network_interface

			deleteInstanceNetworkInterfaceOptions := vpcService.NewDeleteInstanceNetworkInterfaceOptions(
				"testString",
				"testString",
			)

			response, err := vpcService.DeleteInstanceNetworkInterface(deleteInstanceNetworkInterfaceOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_instance_network_interface
			fmt.Printf("\nDeleteInstanceNetworkInterface() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteInstanceGroupMemberships request example`, func() {
			// begin-delete_instance_group_memberships

			deleteInstanceGroupMembershipsOptions := vpcService.NewDeleteInstanceGroupMembershipsOptions(
				"testString",
			)

			response, err := vpcService.DeleteInstanceGroupMemberships(deleteInstanceGroupMembershipsOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_instance_group_memberships
			fmt.Printf("\nDeleteInstanceGroupMemberships() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteInstanceGroupMembership request example`, func() {
			// begin-delete_instance_group_membership

			deleteInstanceGroupMembershipOptions := vpcService.NewDeleteInstanceGroupMembershipOptions(
				"testString",
				"testString",
			)

			response, err := vpcService.DeleteInstanceGroupMembership(deleteInstanceGroupMembershipOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_instance_group_membership
			fmt.Printf("\nDeleteInstanceGroupMembership() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteInstanceGroupManagerPolicy request example`, func() {
			// begin-delete_instance_group_manager_policy

			deleteInstanceGroupManagerPolicyOptions := vpcService.NewDeleteInstanceGroupManagerPolicyOptions(
				"testString",
				"testString",
				"testString",
			)

			response, err := vpcService.DeleteInstanceGroupManagerPolicy(deleteInstanceGroupManagerPolicyOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_instance_group_manager_policy
			fmt.Printf("\nDeleteInstanceGroupManagerPolicy() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteInstanceGroupManagerAction request example`, func() {
			// begin-delete_instance_group_manager_action

			deleteInstanceGroupManagerActionOptions := vpcService.NewDeleteInstanceGroupManagerActionOptions(
				"testString",
				"testString",
				"testString",
			)

			response, err := vpcService.DeleteInstanceGroupManagerAction(deleteInstanceGroupManagerActionOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_instance_group_manager_action
			fmt.Printf("\nDeleteInstanceGroupManagerAction() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteInstanceGroupManager request example`, func() {
			// begin-delete_instance_group_manager

			deleteInstanceGroupManagerOptions := vpcService.NewDeleteInstanceGroupManagerOptions(
				"testString",
				"testString",
			)

			response, err := vpcService.DeleteInstanceGroupManager(deleteInstanceGroupManagerOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_instance_group_manager
			fmt.Printf("\nDeleteInstanceGroupManager() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteInstanceGroupLoadBalancer request example`, func() {
			// begin-delete_instance_group_load_balancer

			deleteInstanceGroupLoadBalancerOptions := vpcService.NewDeleteInstanceGroupLoadBalancerOptions(
				"testString",
			)

			response, err := vpcService.DeleteInstanceGroupLoadBalancer(deleteInstanceGroupLoadBalancerOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_instance_group_load_balancer
			fmt.Printf("\nDeleteInstanceGroupLoadBalancer() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteInstanceGroup request example`, func() {
			// begin-delete_instance_group

			deleteInstanceGroupOptions := vpcService.NewDeleteInstanceGroupOptions(
				"testString",
			)

			response, err := vpcService.DeleteInstanceGroup(deleteInstanceGroupOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_instance_group
			fmt.Printf("\nDeleteInstanceGroup() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteInstance request example`, func() {
			// begin-delete_instance

			deleteInstanceOptions := vpcService.NewDeleteInstanceOptions(
				"testString",
			)

			response, err := vpcService.DeleteInstance(deleteInstanceOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_instance
			fmt.Printf("\nDeleteInstance() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteImage request example`, func() {
			// begin-delete_image

			deleteImageOptions := vpcService.NewDeleteImageOptions(
				"testString",
			)

			response, err := vpcService.DeleteImage(deleteImageOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_image
			fmt.Printf("\nDeleteImage() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})
		It(`DeleteIkePolicy request example`, func() {
			// begin-delete_ike_policy

			deleteIkePolicyOptions := vpcService.NewDeleteIkePolicyOptions(
				"testString",
			)

			response, err := vpcService.DeleteIkePolicy(deleteIkePolicyOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_ike_policy
			fmt.Printf("\nDeleteIkePolicy() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteFlowLogCollector request example`, func() {
			// begin-delete_flow_log_collector

			deleteFlowLogCollectorOptions := vpcService.NewDeleteFlowLogCollectorOptions(
				"testString",
			)

			response, err := vpcService.DeleteFlowLogCollector(deleteFlowLogCollectorOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_flow_log_collector
			fmt.Printf("\nDeleteFlowLogCollector() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteFloatingIP request example`, func() {
			// begin-delete_floating_ip

			deleteFloatingIPOptions := vpcService.NewDeleteFloatingIPOptions(
				"testString",
			)

			response, err := vpcService.DeleteFloatingIP(deleteFloatingIPOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_floating_ip
			fmt.Printf("\nDeleteFloatingIP() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteEndpointGateway request example`, func() {
			// begin-delete_endpoint_gateway

			deleteEndpointGatewayOptions := vpcService.NewDeleteEndpointGatewayOptions(
				"testString",
			)

			response, err := vpcService.DeleteEndpointGateway(deleteEndpointGatewayOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_endpoint_gateway
			fmt.Printf("\nDeleteEndpointGateway() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteDedicatedHostGroup request example`, func() {
			// begin-delete_dedicated_host_group

			deleteDedicatedHostGroupOptions := vpcService.NewDeleteDedicatedHostGroupOptions(
				"testString",
			)

			response, err := vpcService.DeleteDedicatedHostGroup(deleteDedicatedHostGroupOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_dedicated_host_group
			fmt.Printf("\nDeleteDedicatedHostGroup() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteDedicatedHost request example`, func() {
			// begin-delete_dedicated_host

			deleteDedicatedHostOptions := vpcService.NewDeleteDedicatedHostOptions(
				"testString",
			)

			response, err := vpcService.DeleteDedicatedHost(deleteDedicatedHostOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_dedicated_host
			fmt.Printf("\nDeleteDedicatedHost() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})
