//go:build examples
// +build examples

/**
 * (C) Copyright IBM Corp. 2020, 2021, 2022.
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
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
	"reflect"
	"strconv"
	"time"
)

var (
	vpcService                        *vpcv1.VpcV1
	serviceErr                        error
	configLoaded                      bool = false
	externalConfigFile                     = "../vpc.env"
	backupPolicyID                    string
	backupPolicyPlanID                string
	backupPolicyJobID                 string
	vpcID                             string
	subnetID                          string
	keyID                             string
	imageID                           string
	instanceID                        string
	addressPrefixID                   string
	routingTableID                    string
	routeID                           string
	eth2ID                            string
	floatingIPID                      string
	volumeID                          string
	snapshotID                        string
	volumeAttachmentID                string
	reservedIPID                      string
	reservedIPID2                     string
	ifMatchVolume                     string
	ifMatchBackupPolicy               string
	ifMatchSnapshot                   string
	ifMatchVPNServer                  string
	instanceTemplateID                string
	instanceGroupID                   string
	instanceGroupManagerID            string
	instanceGroupManagerPolicyID      string
	instanceGroupManagerActionID      string
	instanceGroupMembershipID         string
	dedicatedHostGroupID              string
	dedicatedHostID                   string
	publicGatewayID                   string
	diskID                            string
	dhID                              string
	securityGroupID                   string
	ikePolicyID                       string
	ipsecPolicyID                     string
	securityGroupRuleID               string
	networkACLID                      string
	targetID                          string
	networkACLRuleID                  string
	vpnGatewayConnectionID            string
	vpnGatewayID                      string
	endpointGatewayID                 string
	placementGroupID                  string
	loadBalancerID                    string
	listenerID                        string
	policyID                          string
	policyRuleID                      string
	poolID                            string
	poolMemberID                      string
	endpointGatewayTargetID           string
	flowLogID                         string
	dhProfile                         string
	operatingSystemName               string
	instanceProfileName               string
	timestamp                         = strconv.FormatInt(tunix, 10)
	tunix                             = time.Now().Unix()
	zone                              *string
	resourceGroupID                   *string
	bareMetalServerProfileName        string
	bareMetalServerId                 string
	bareMetalServerDiskId             string
	bareMetalServerNetworkInterfaceId string
	vpnClientID                       string
	vpnServerRouteID                  string
	vpnServerID                       string
)

func skipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

func getName(rtype string) string {
	return "gsdk-" + rtype + "-" + timestamp
}

var _ = Describe(`VpcV1 Examples Tests`, func() {
	Describe(`External configuration`, func() {

		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}
			if err = os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile); err == nil {
				configLoaded = true
			}
			Expect(err).To(BeNil())
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			skipTest()
		})
		It("Successfully construct the service client instance", func() {

			// begin-common

			vpcService, serviceErr = vpcv1.NewVpcV1UsingExternalConfig(
				&vpcv1.VpcV1Options{
					ServiceName: "vpcint",
				},
			)
			if serviceErr != nil {
				fmt.Println("Gen2 Service creation failed.", serviceErr)
			}

			// end-common

			Expect(vpcService).ToNot(BeNil())
		})
	})
	Describe(`Variable setting`, func() {
		BeforeEach(func() {
			skipTest()
		})
		It("Setting up required variable", func() {
			listSubnetsOptions := &vpcv1.ListSubnetsOptions{}

			subnetCollection, _, err := vpcService.ListSubnets(listSubnetsOptions)
			zone = subnetCollection.Subnets[0].Zone.Name
			resourceGroupID = subnetCollection.Subnets[0].ResourceGroup.ID
			Expect(subnetCollection).ToNot(BeNil())
			Expect(zone).ToNot(BeNil())
			Expect(resourceGroupID).ToNot(BeNil())
			Expect(err).To(BeNil())

		})
	})

	Describe(`VpcV1 request examples`, func() {
		BeforeEach(func() {
			skipTest()
		})
		It(`ListVpcs request example`, func() {
			fmt.Println("\nListVpcs() result:")

			// begin-list_vpcs
			listVpcsOptions := &vpcv1.ListVpcsOptions{}
			vpcs, response, err := vpcService.ListVpcs(listVpcsOptions)

			// end-list_vpcs
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpcs).ToNot(BeNil())

		})
		It(`CreateVPC request example`, func() {
			fmt.Println("\nCreateVPC() result:")

			classicAccess := true
			manual := "manual"
			// begin-create_vpc

			options := &vpcv1.CreateVPCOptions{
				ResourceGroup: &vpcv1.ResourceGroupIdentity{
					ID: resourceGroupID,
				},
				Name:                    &[]string{"my-vpc"}[0],
				ClassicAccess:           &classicAccess,
				AddressPrefixManagement: &manual,
			}
			vpc, response, err := vpcService.CreateVPC(options)

			// end-create_vpc
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(vpc).ToNot(BeNil())
			vpcID = *vpc.ID
		})
		It(`GetVPC request example`, func() {
			fmt.Println("\nGetVPC() result:")
			// begin-get_vpc

			getVpcOptions := &vpcv1.GetVPCOptions{
				ID: &vpcID,
			}
			vpc, response, err := vpcService.GetVPC(getVpcOptions)
			// end-get_vpc
			if err != nil {
				panic(err)
			}

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpc).ToNot(BeNil())

		})
		It(`UpdateVPC request example`, func() {
			fmt.Println("\nUpdateVPC() result:")
			// begin-update_vpc

			options := &vpcv1.UpdateVPCOptions{
				ID: &vpcID,
			}
			vpcPatchModel := &vpcv1.VPCPatch{
				Name: &[]string{"my-vpc-modified"}[0],
			}
			vpcPatch, asPatchErr := vpcPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.VPCPatch = vpcPatch
			vpc, response, err := vpcService.UpdateVPC(options)

			// end-update_vpc
			if err != nil {
				panic(err)
			}

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpc).ToNot(BeNil())

		})
		It(`GetVPCDefaultNetworkACL request example`, func() {
			fmt.Println("\nGetVPCDefaultNetworkACL() result:")
			// begin-get_vpc_default_network_acl

			options := &vpcv1.GetVPCDefaultNetworkACLOptions{}
			options.SetID(vpcID)
			defaultACL, response, err := vpcService.GetVPCDefaultNetworkACL(options)

			// end-get_vpc_default_network_acl
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(defaultACL).ToNot(BeNil())

		})
		It(`GetVPCDefaultRoutingTable request example`, func() {
			fmt.Println("\nGetVPCDefaultRoutingTable() result:")
			// begin-get_vpc_default_routing_table

			options := vpcService.NewGetVPCDefaultRoutingTableOptions(vpcID)
			defaultRoutingTable, response, err := vpcService.GetVPCDefaultRoutingTable(options)

			// end-get_vpc_default_routing_table

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(defaultRoutingTable).ToNot(BeNil())

		})
		It(`GetVPCDefaultSecurityGroup request example`, func() {
			fmt.Println("\nGetVPCDefaultSecurityGroup() result:")
			// begin-get_vpc_default_security_group

			options := &vpcv1.GetVPCDefaultSecurityGroupOptions{}
			options.SetID(vpcID)
			defaultSG, response, err := vpcService.GetVPCDefaultSecurityGroup(options)
			// end-get_vpc_default_security_group
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(defaultSG).ToNot(BeNil())

		})
		It(`ListVPCAddressPrefixes request example`, func() {
			fmt.Println("\nListVPCAddressPrefixes() result:")
			// begin-list_vpc_address_prefixes
			listVpcAddressPrefixesOptions := &vpcv1.ListVPCAddressPrefixesOptions{}
			listVpcAddressPrefixesOptions.SetVPCID(vpcID)
			addressPrefixes, response, err :=
				vpcService.ListVPCAddressPrefixes(listVpcAddressPrefixesOptions)

			// end-list_vpc_address_prefixes
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(addressPrefixes).ToNot(BeNil())

		})
		It(`CreateVPCAddressPrefix request example`, func() {
			fmt.Println("\nCreateVPCAddressPrefix() result:")
			// begin-create_vpc_address_prefix

			options := &vpcv1.CreateVPCAddressPrefixOptions{}
			options.SetVPCID(vpcID)
			options.SetCIDR("10.0.0.0/24")
			options.SetName("my-address-prefix")
			options.SetZone(&vpcv1.ZoneIdentity{
				Name: zone,
			})
			addressPrefix, response, err := vpcService.CreateVPCAddressPrefix(options)
			// end-create_vpc_address_prefix
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(addressPrefix).ToNot(BeNil())
			addressPrefixID = *addressPrefix.ID

		})
		It(`GetVPCAddressPrefix request example`, func() {
			fmt.Println("\nGetVPCAddressPrefix() result:")
			// begin-get_vpc_address_prefix

			getVpcAddressPrefixOptions := &vpcv1.GetVPCAddressPrefixOptions{}
			getVpcAddressPrefixOptions.SetVPCID(vpcID)
			getVpcAddressPrefixOptions.SetID(addressPrefixID)
			addressPrefix, response, err :=
				vpcService.GetVPCAddressPrefix(getVpcAddressPrefixOptions)

			// end-get_vpc_address_prefix
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(addressPrefix).ToNot(BeNil())

		})
		It(`UpdateVPCAddressPrefix request example`, func() {
			fmt.Println("\nUpdateVPCAddressPrefix() result:")
			isDefault := true
			// begin-update_vpc_address_prefix
			options := &vpcv1.UpdateVPCAddressPrefixOptions{}
			options.SetVPCID(vpcID)
			options.SetID(addressPrefixID)
			addressPrefixPatchModel := &vpcv1.AddressPrefixPatch{}
			addressPrefixPatchModel.Name = &[]string{"my-address-prefix-updated"}[0]
			addressPrefixPatchModel.IsDefault = &isDefault
			addressPrefixPatch, asPatchErr := addressPrefixPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.AddressPrefixPatch = addressPrefixPatch
			addressPrefix, response, err := vpcService.UpdateVPCAddressPrefix(options)

			// end-update_vpc_address_prefix
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(addressPrefix).ToNot(BeNil())

		})
		It(`ListVPCRoutingTables request example`, func() {
			fmt.Println("\nListVPCRoutingTables() result:")
			// begin-list_vpc_routing_tables

			options := vpcService.NewListVPCRoutingTablesOptions(vpcID)
			routingTableCollection, response, err := vpcService.ListVPCRoutingTables(options)

			// end-list_vpc_routing_tables
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(routingTableCollection).ToNot(BeNil())

		})
		It(`CreateVPCRoutingTable request example`, func() {
			fmt.Println("\nCreateVPCRoutingTable() result:")
			routeName := "my-route"
			action := "delegate"
			// begin-create_vpc_routing_table
			routePrototypeModel := &vpcv1.RoutePrototype{
				Action: &action,
				NextHop: &vpcv1.RoutePrototypeNextHopRouteNextHopPrototypeRouteNextHopIP{
					Address: &[]string{"192.168.3.4"}[0],
				},
				Name:        &routeName,
				Destination: &[]string{"192.168.3.0/24"}[0],
				Zone: &vpcv1.ZoneIdentityByName{
					Name: zone,
				},
			}
			name := "my-routing-table"
			options := &vpcv1.CreateVPCRoutingTableOptions{
				VPCID:  &vpcID,
				Name:   &name,
				Routes: []vpcv1.RoutePrototype{*routePrototypeModel},
			}
			routingTable, response, err := vpcService.CreateVPCRoutingTable(options)
			// end-create_vpc_routing_table
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(routingTable).ToNot(BeNil())
			routingTableID = *routingTable.ID
		})
		It(`GetVPCRoutingTable request example`, func() {
			fmt.Println("\nGetVPCRoutingTable() result:")
			// begin-get_vpc_routing_table

			options := &vpcv1.GetVPCRoutingTableOptions{
				VPCID: &vpcID,
				ID:    &routingTableID,
			}
			routingTable, response, err := vpcService.GetVPCRoutingTable(options)
			// end-get_vpc_routing_table
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(routingTable).ToNot(BeNil())
		})
		It(`UpdateVPCRoutingTable request example`, func() {
			fmt.Println("\nUpdateVPCRoutingTable() result:")
			// begin-update_vpc_routing_table

			name := "my-routing-table"
			routingTablePatchModel := &vpcv1.RoutingTablePatch{
				Name: &name,
			}
			routingTablePatchModelAsPatch, asPatchErr := routingTablePatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options := &vpcv1.UpdateVPCRoutingTableOptions{
				VPCID:             &vpcID,
				ID:                &routingTableID,
				RoutingTablePatch: routingTablePatchModelAsPatch,
			}
			routingTable, response, err := vpcService.UpdateVPCRoutingTable(options)

			// end-update_vpc_routing_table
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(routingTable).ToNot(BeNil())

		})
		It(`ListVPCRoutingTableRoutes request example`, func() {
			fmt.Println("\nListVPCRoutingTableRoutes() result:")
			// begin-list_vpc_routing_table_routes

			options := vpcService.NewListVPCRoutingTableRoutesOptions(
				vpcID,
				routingTableID,
			)
			routeCollection, response, err := vpcService.ListVPCRoutingTableRoutes(options)

			// end-list_vpc_routing_table_routes
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(routeCollection).ToNot(BeNil())

		})
		It(`CreateVPCRoutingTableRoute request example`, func() {
			fmt.Println("\nCreateVPCRoutingTableRoute() result:")
			destination := "192.168.77.0/24"
			address := "192.168.3.7"
			// begin-create_vpc_routing_table_route
			zoneIdentityModel := &vpcv1.ZoneIdentityByName{
				Name: zone,
			}
			options := &vpcv1.CreateVPCRoutingTableRouteOptions{
				VPCID:          &vpcID,
				RoutingTableID: &routingTableID,
				Destination:    &destination,
				Zone:           zoneIdentityModel,
				NextHop: &vpcv1.RoutePrototypeNextHopRouteNextHopPrototypeRouteNextHopIP{
					Address: &address,
				},
			}
			route, response, err := vpcService.CreateVPCRoutingTableRoute(options)

			// end-create_vpc_routing_table_route
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(route).ToNot(BeNil())
			routeID = *route.ID
		})
		It(`GetVPCRoutingTableRoute request example`, func() {
			fmt.Println("\nGetVPCRoutingTableRoute() result:")
			// begin-get_vpc_routing_table_route

			options := vpcService.NewGetVPCRoutingTableRouteOptions(
				vpcID,
				routingTableID,
				routeID,
			)
			route, response, err := vpcService.GetVPCRoutingTableRoute(options)

			// end-get_vpc_routing_table_route
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(route).ToNot(BeNil())

		})
		It(`UpdateVPCRoutingTableRoute request example`, func() {
			fmt.Println("\nUpdateVPCRoutingTableRoute() result:")
			// begin-update_vpc_routing_table_route

			name := "my-route-updated"
			routePatchModel := &vpcv1.RoutePatch{
				Name: &name,
			}
			routePatchModelAsPatch, asPatchErr := routePatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options := &vpcv1.UpdateVPCRoutingTableRouteOptions{
				VPCID:          &vpcID,
				RoutingTableID: &routingTableID,
				ID:             &routeID,
				RoutePatch:     routePatchModelAsPatch,
			}
			route, response, err := vpcService.UpdateVPCRoutingTableRoute(options)

			// end-update_vpc_routing_table_route
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(route).ToNot(BeNil())

		})
		It(`ListSubnets request example`, func() {
			fmt.Println("\nListSubnets() result:")
			// begin-list_subnets

			options := &vpcv1.ListSubnetsOptions{}
			subnets, response, err := vpcService.ListSubnets(options)

			// end-list_subnets

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subnets).ToNot(BeNil())

		})
		It(`CreateSubnet request example`, func() {
			fmt.Println("\nCreateSubnet() result:")
			cidrBlock := "10.0.1.0/24"
			// begin-create_subnet

			options := &vpcv1.CreateSubnetOptions{}
			options.SetSubnetPrototype(&vpcv1.SubnetPrototype{
				Ipv4CIDRBlock: &cidrBlock,
				Name:          &[]string{"my-subnet"}[0],
				VPC: &vpcv1.VPCIdentity{
					ID: &vpcID,
				},
				Zone: &vpcv1.ZoneIdentity{
					Name: zone,
				},
			})
			subnet, response, err := vpcService.CreateSubnet(options)

			// end-create_subnet

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subnet).ToNot(BeNil())
			subnetID = *subnet.ID
		})
		It(`GetSubnet request example`, func() {
			fmt.Println("\nGetSubnet() result:")
			// begin-get_subnet

			options := &vpcv1.GetSubnetOptions{}
			options.SetID(subnetID)
			subnet, response, err := vpcService.GetSubnet(options)

			// end-get_subnet

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subnet).ToNot(BeNil())

		})
		It(`UpdateSubnet request example`, func() {
			fmt.Println("\nUpdateSubnet() result:")
			name := getName("subnet")
			networkAclId := &networkACLID
			routingTableId := &[]string{""}[0]
			// begin-update_subnet

			options := &vpcv1.UpdateSubnetOptions{}
			options.SetID(subnetID)
			subnetPatchModel := &vpcv1.SubnetPatch{}
			subnetPatchModel.Name = &name
			subnetPatchModel.NetworkACL = &vpcv1.NetworkACLIdentity{
				ID: networkAclId,
			}
			routingTableIdentityModel := new(vpcv1.RoutingTableIdentityByID)
			routingTableIdentityModel.ID = routingTableId
			subnetPatchModel.RoutingTable = routingTableIdentityModel
			subnetPatch, asPatchErr := subnetPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.SubnetPatch = subnetPatch
			subnet, response, err := vpcService.UpdateSubnet(options)

			// end-update_subnet

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subnet).ToNot(BeNil())

		})
		It(`ReplaceSubnetNetworkACL request example`, func() {
			fmt.Println("\nReplaceSubnetNetworkACL() result:")
			vpcIDentityModel := &vpcv1.VPCIdentityByID{
				ID: &vpcID,
			}
			networkACLPrototypeModel := &vpcv1.NetworkACLPrototypeNetworkACLByRules{
				Name: &[]string{"my-network-acl"}[0],
				VPC:  vpcIDentityModel,
			}
			createNetworkACLOptions := vpcService.NewCreateNetworkACLOptions()
			createNetworkACLOptions.SetNetworkACLPrototype(networkACLPrototypeModel)
			networkACL, _, _ := vpcService.CreateNetworkACL(createNetworkACLOptions)
			Expect(networkACL).ToNot(BeNil())
			networkACLID := networkACL.ID
			// begin-replace_subnet_network_acl

			options := &vpcv1.ReplaceSubnetNetworkACLOptions{}
			options.SetID(subnetID)
			options.SetNetworkACLIdentity(&vpcv1.NetworkACLIdentity{
				ID: networkACLID,
			})
			networkACL, response, err := vpcService.ReplaceSubnetNetworkACL(options)

			// end-replace_subnet_network_acl
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(networkACL).ToNot(BeNil())

		})
		It(`GetSubnetNetworkACL request example`, func() {
			fmt.Println("\nGetSubnetNetworkACL() result:")
			// begin-get_subnet_network_acl

			options := &vpcv1.GetSubnetNetworkACLOptions{}
			options.SetID(subnetID)
			acls, response, err := vpcService.GetSubnetNetworkACL(options)

			// end-get_subnet_network_acl
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(acls).ToNot(BeNil())

		})
		It(`SetSubnetPublicGateway request example`, func() {
			fmt.Println("\nSetSubnetPublicGateway() result:")
			vpcIDentityModel := &vpcv1.VPCIdentityByID{
				ID: &vpcID,
			}

			zoneIdentityModel := &vpcv1.ZoneIdentityByName{
				Name: zone,
			}

			createPublicGatewayOptions := vpcService.NewCreatePublicGatewayOptions(
				vpcIDentityModel,
				zoneIdentityModel,
			)

			publicGateway, _, err := vpcService.CreatePublicGateway(createPublicGatewayOptions)
			if err != nil {
				panic(err)
			}
			Expect(publicGateway).ToNot(BeNil())

			// begin-set_subnet_public_gateway

			options := &vpcv1.SetSubnetPublicGatewayOptions{}
			options.SetID(subnetID)
			options.SetPublicGatewayIdentity(&vpcv1.PublicGatewayIdentity{
				ID: publicGateway.ID,
			})
			publicGateway, response, err := vpcService.SetSubnetPublicGateway(options)
			// end-set_subnet_public_gateway
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(publicGateway).ToNot(BeNil())
		})
		It(`GetSubnetPublicGateway request example`, func() {
			fmt.Println("\nGetSubnetPublicGateway() result:")
			// begin-get_subnet_public_gateway

			options := &vpcv1.GetSubnetPublicGatewayOptions{}
			options.SetID(subnetID)
			publicGateway, response, err := vpcService.GetSubnetPublicGateway(options)

			// end-get_subnet_public_gateway
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(publicGateway).ToNot(BeNil())

		})

		It(`UnsetSubnetPublicGateway request example`, func() {
			// begin-unset_subnet_public_gateway

			options := vpcService.NewUnsetSubnetPublicGatewayOptions(
				subnetID,
			)

			response, err := vpcService.UnsetSubnetPublicGateway(options)

			// end-unset_subnet_public_gateway
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nUnsetSubnetPublicGateway() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`ReplaceSubnetRoutingTable request example`, func() {
			fmt.Println("\nReplaceSubnetRoutingTable() result:")
			// begin-replace_subnet_routing_table

			routingTableIdentityModel := &vpcv1.RoutingTableIdentityByID{
				ID: &routingTableID,
			}
			replaceSubnetRoutingTableOptions := vpcService.NewReplaceSubnetRoutingTableOptions(
				subnetID,
				routingTableIdentityModel,
			)
			routingTable, response, err := vpcService.ReplaceSubnetRoutingTable(
				replaceSubnetRoutingTableOptions,
			)

			// end-replace_subnet_routing_table
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(routingTable).ToNot(BeNil())

		})
		It(`GetSubnetRoutingTable request example`, func() {
			fmt.Println("\nGetSubnetRoutingTable() result:")
			// begin-get_subnet_routing_table
			options := vpcService.NewGetSubnetRoutingTableOptions(subnetID)
			routingTable, response, err := vpcService.GetSubnetRoutingTable(options)

			// end-get_subnet_routing_table
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(routingTable).ToNot(BeNil())

		})
		It(`ListSubnetReservedIps request example`, func() {
			fmt.Println("\nListSubnetReservedIps() result:")
			// begin-list_subnet_reserved_ips

			options := vpcService.NewListSubnetReservedIpsOptions(subnetID)
			reservedIPCollection, response, err := vpcService.ListSubnetReservedIps(options)

			// end-list_subnet_reserved_ips
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reservedIPCollection).ToNot(BeNil())

		})
		It(`CreateSubnetReservedIP request example`, func() {
			fmt.Println("\nCreateSubnetReservedIP() result:")
			name := getName("subnetRip")
			// begin-create_subnet_reserved_ip

			options := vpcService.NewCreateSubnetReservedIPOptions(subnetID)
			options.Name = &name
			reservedIP, response, err := vpcService.CreateSubnetReservedIP(options)

			// end-create_subnet_reserved_ip
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(reservedIP).ToNot(BeNil())
			reservedIPID = *reservedIP.ID

		})
		It(`GetSubnetReservedIP request example`, func() {
			fmt.Println("\nGetSubnetReservedIP() result:")
			// begin-get_subnet_reserved_ip

			options := vpcService.NewGetSubnetReservedIPOptions(subnetID, reservedIPID)
			reservedIP, response, err := vpcService.GetSubnetReservedIP(options)

			// end-get_subnet_reserved_ip
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reservedIP).ToNot(BeNil())

		})
		It(`UpdateSubnetReservedIP request example`, func() {
			fmt.Println("\nUpdateSubnetReservedIP() result:")
			name := getName("subnetRip")
			// begin-update_subnet_reserved_ip

			options := &vpcv1.UpdateSubnetReservedIPOptions{}

			patchBody := new(vpcv1.ReservedIPPatch)
			patchBody.Name = &name
			patchBody.AutoDelete = &[]bool{true}[0]
			reservedIPPatch, asPatchErr := patchBody.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.SetReservedIPPatch(reservedIPPatch)
			options.SetID(reservedIPID)
			options.SetSubnetID(subnetID)
			reservedIP, response, err := vpcService.UpdateSubnetReservedIP(options)

			// end-update_subnet_reserved_ip
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reservedIP).ToNot(BeNil())

		})
		It(`DeleteSubnetReservedIP request example`, func() {
			// begin-delete_subnet_reserved_ip
			deleteSubnetReservedIPOptions := vpcService.NewDeleteSubnetReservedIPOptions(
				subnetID,
				reservedIPID,
			)

			response, err := vpcService.DeleteSubnetReservedIP(deleteSubnetReservedIPOptions)

			// end-delete_subnet_reserved_ip
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteSubnetReservedIP() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
			// create another reserved ip for endpoint gateway
			name := getName("subnetRip")
			options := vpcService.NewCreateSubnetReservedIPOptions(subnetID)
			options.Name = &name
			reservedIP, response, err := vpcService.CreateSubnetReservedIP(options)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(reservedIP).ToNot(BeNil())
			reservedIPID = *reservedIP.ID
		})
		It(`ListImages request example`, func() {
			fmt.Println("\nListImages() result:")
			// begin-list_images
			options := &vpcv1.ListImagesOptions{}
			options.SetVisibility("private")
			images, response, err := vpcService.ListImages(options)

			// end-list_images
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(images).ToNot(BeNil())
		})
		It(`CreateImage request example`, func() {
			fmt.Println("\nCreateImage() result:")
			name := getName("image")
			// begin-create_image

			operatingSystemIdentityModel := &vpcv1.OperatingSystemIdentityByName{
				Name: &[]string{"debian-9-amd64"}[0],
			}

			options := &vpcv1.CreateImageOptions{}
			cosID := "cos://us-south/my-bucket/my-image.qcow2"
			options.SetImagePrototype(&vpcv1.ImagePrototype{
				Name: &name,
				File: &vpcv1.ImageFilePrototype{
					Href: &cosID,
				},
				OperatingSystem: operatingSystemIdentityModel,
			})
			image, response, err := vpcService.CreateImage(options)

			// end-create_image

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(image).ToNot(BeNil())
			imageID = *image.ID
		})
		It(`GetImage request example`, func() {
			fmt.Println("\nGetImage() result:")
			// begin-get_image
			options := &vpcv1.GetImageOptions{}
			options.SetID(imageID)
			image, response, err := vpcService.GetImage(options)
			// end-get_image
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(image).ToNot(BeNil())

		})
		It(`UpdateImage request example`, func() {
			fmt.Println("\nUpdateImage() result:")
			name := getName("image")
			// begin-update_image

			options := &vpcv1.UpdateImageOptions{}
			options.SetID(imageID)
			imagePatchModel := &vpcv1.ImagePatch{
				Name: &name,
			}
			imagePatch, asPatchErr := imagePatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.ImagePatch = imagePatch
			image, response, err := vpcService.UpdateImage(options)

			// end-update_image
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(image).ToNot(BeNil())

		})
		It(`ListOperatingSystems request example`, func() {
			fmt.Println("\nListOperatingSystems() result:")
			// begin-list_operating_systems

			options := &vpcv1.ListOperatingSystemsOptions{}
			operatingSystems, response, err := vpcService.ListOperatingSystems(options)

			// end-list_operating_systems
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(operatingSystems).ToNot(BeNil())
			operatingSystemName = *operatingSystems.OperatingSystems[0].Name
		})
		It(`GetOperatingSystem request example`, func() {
			fmt.Println("\nGetOperatingSystem() result:")
			// begin-get_operating_system

			options := &vpcv1.GetOperatingSystemOptions{}
			options.SetName(operatingSystemName)
			operatingSystem, response, err := vpcService.GetOperatingSystem(options)

			// end-get_operating_system
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(operatingSystem).ToNot(BeNil())

		})
		It(`ListKeys request example`, func() {
			fmt.Println("\nListKeys() result:")
			// begin-list_keys

			listKeysOptions := &vpcv1.ListKeysOptions{}
			keys, response, err := vpcService.ListKeys(listKeysOptions)

			// end-list_keys

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(keys).ToNot(BeNil())

		})
		It(`CreateKey request example`, func() {
			fmt.Println("\nCreateKey() result:")
			name := getName("sshkey")
			publicKey := "AAAAB3NzaC1yc2EAAAADAQABAAABAQDDGe50Bxa5T5NDddrrtbx2Y4/VGbiCgXqnBsYToIUKoFSHTQl5IX3PasGnneKanhcLwWz5M5MoCRvhxTp66NKzIfAz7r+FX9rxgR+ZgcM253YAqOVeIpOU408simDZKriTlN8kYsXL7P34tsWuAJf4MgZtJAQxous/2byetpdCv8ddnT4X3ltOg9w+LqSCPYfNivqH00Eh7S1Ldz7I8aw5WOp5a+sQFP/RbwfpwHp+ny7DfeIOokcuI42tJkoBn7UsLTVpCSmXr2EDRlSWe/1M/iHNRBzaT3CK0+SwZWd2AEjePxSnWKNGIEUJDlUYp7hKhiQcgT5ZAnWU121oc5En"

			// begin-create_key
			options := &vpcv1.CreateKeyOptions{}
			options.SetName(name)
			options.SetPublicKey(publicKey)
			key, response, err := vpcService.CreateKey(options)

			// end-create_key

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(key).ToNot(BeNil())
			keyID = *key.ID
		})
		It(`GetKey request example`, func() {
			fmt.Println("\nGetKey() result:")
			// begin-get_key

			getKeyOptions := &vpcv1.GetKeyOptions{}
			getKeyOptions.SetID(keyID)
			key, response, err := vpcService.GetKey(getKeyOptions)

			// end-get_key

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(key).ToNot(BeNil())

		})
		It(`UpdateKey request example`, func() {
			fmt.Println("\nUpdateKey() result:")
			// begin-update_key

			options := &vpcv1.UpdateKeyOptions{}
			options.SetID(keyID)
			keyPatchModel := &vpcv1.KeyPatch{
				Name: &[]string{"my-key-modified"}[0],
			}
			keyPatch, asPatchErr := keyPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.KeyPatch = keyPatch
			key, response, err := vpcService.UpdateKey(options)

			// end-update_key

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(key).ToNot(BeNil())

		})
		It(`ListFloatingIps request example`, func() {
			fmt.Println("\nListFloatingIps() result:")
			// begin-list_floating_ips
			listFloatingIpsOptions := vpcService.NewListFloatingIpsOptions()
			floatingIPs, response, err :=
				vpcService.ListFloatingIps(listFloatingIpsOptions)

			// end-list_floating_ips
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(floatingIPs).ToNot(BeNil())

		})
		It(`CreateFloatingIP request example`, func() {
			fmt.Println("\nCreateFloatingIP() result:")
			name := getName("floatingIP")
			// begin-create_floating_ip

			options := &vpcv1.CreateFloatingIPOptions{}
			options.SetFloatingIPPrototype(&vpcv1.FloatingIPPrototype{
				Name: &name,
				Zone: &vpcv1.ZoneIdentity{
					Name: zone,
				},
			})
			floatingIP, response, err := vpcService.CreateFloatingIP(options)

			// end-create_floating_ip
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(floatingIP).ToNot(BeNil())
			floatingIPID = *floatingIP.ID
		})
		It(`GetFloatingIP request example`, func() {
			fmt.Println("\nGetFloatingIP() result:")
			// begin-get_floating_ip

			options := vpcService.NewGetFloatingIPOptions(floatingIPID)
			floatingIP, response, err := vpcService.GetFloatingIP(options)

			// end-get_floating_ip
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(floatingIP).ToNot(BeNil())

		})
		It(`UpdateFloatingIP request example`, func() {
			name := getName("fip")
			fmt.Println("\nUpdateFloatingIP() result:")
			// begin-update_floating_ip

			floatingIPPatchModel := &vpcv1.FloatingIPPatch{
				Name: &name,
			}
			floatingIPPatchModelAsPatch, asPatchErr := floatingIPPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}

			updateFloatingIPOptions := vpcService.NewUpdateFloatingIPOptions(
				floatingIPID,
				floatingIPPatchModelAsPatch,
			)

			floatingIP, response, err := vpcService.UpdateFloatingIP(updateFloatingIPOptions)

			// end-update_floating_ip
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(floatingIP).ToNot(BeNil())

		})
		It(`ListVolumes request example`, func() {
			fmt.Println("\nListVolumes() result:")
			// begin-list_volumes

			options := &vpcv1.ListVolumesOptions{}
			volumes, response, err := vpcService.ListVolumes(options)

			// end-list_volumes
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumes).ToNot(BeNil())

		})
		It(`CreateVolume request example`, func() {
			fmt.Println("\nCreateVolume() result:")
			name := getName("vol")
			// begin-create_volume
			options := &vpcv1.CreateVolumeOptions{}
			options.SetVolumePrototype(&vpcv1.VolumePrototype{
				Capacity: &[]int64{100}[0],
				Zone: &vpcv1.ZoneIdentity{
					Name: zone,
				},
				Profile: &vpcv1.VolumeProfileIdentity{
					Name: &[]string{"general-purpose"}[0],
				},
				Name: &name,
			})
			volume, response, err := vpcService.CreateVolume(options)
			// end-create_volume
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(volume).ToNot(BeNil())
			volumeID = *volume.ID
		})
		It(`GetVolume request example`, func() {
			fmt.Println("\nGetVolume() result:")
			// begin-get_volume

			options := &vpcv1.GetVolumeOptions{}
			options.SetID(volumeID)
			volume, response, err := vpcService.GetVolume(options)

			// end-get_volume
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volume).ToNot(BeNil())
			ifMatchVolume = response.GetHeaders()["Etag"][0]
		})
		It(`UpdateVolume request example`, func() {
			fmt.Println("\nUpdateVolume() result:")
			name := getName("vol")
			userTags := []string{"usertag-vol-1"}
			// begin-update_volume

			options := &vpcv1.UpdateVolumeOptions{}
			options.SetID(volumeID)
			options.SetIfMatch(ifMatchVolume)
			volumePatchModel := &vpcv1.VolumePatch{
				Name:     &name,
				UserTags: userTags,
			}
			volumePatch, asPatchErr := volumePatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.VolumePatch = volumePatch
			volume, response, err := vpcService.UpdateVolume(options)
			// end-update_volume
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volume).ToNot(BeNil())

		})
		It(`ListInstanceProfiles request example`, func() {
			fmt.Println("\nListInstanceProfiles() result:")
			// begin-list_instance_profiles

			options := &vpcv1.ListInstanceProfilesOptions{}
			profiles, response, err := vpcService.ListInstanceProfiles(options)

			// end-list_instance_profiles
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profiles).ToNot(BeNil())
			instanceProfileName = *profiles.Profiles[0].Name
		})
		It(`GetInstanceProfile request example`, func() {
			fmt.Println("\nGetInstanceProfile() result:")
			// begin-get_instance_profile

			options := &vpcv1.GetInstanceProfileOptions{}
			options.SetName(instanceProfileName)
			profile, response, err := vpcService.GetInstanceProfile(options)
			// end-get_instance_profile
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profile).ToNot(BeNil())

		})
		It(`ListInstanceTemplates request example`, func() {
			fmt.Println("\nListInstanceTemplates() result:")
			// begin-list_instance_templates

			options := &vpcv1.ListInstanceTemplatesOptions{}
			instanceTemplates, response, err := vpcService.ListInstanceTemplates(options)

			// end-list_instance_templates
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceTemplates).ToNot(BeNil())

		})
		It(`CreateInstanceTemplate request example`, func() {
			fmt.Println("\nCreateInstanceTemplate() result:")
			name := getName("template")
			instanceProfile := []string{"bx2d-2x8"}[0]
			// begin-create_instance_template

			options := &vpcv1.CreateInstanceTemplateOptions{}
			options.SetInstanceTemplatePrototype(&vpcv1.InstanceTemplatePrototype{
				Name: &name,
				Image: &vpcv1.ImageIdentity{
					ID: &imageID,
				},
				Profile: &vpcv1.InstanceProfileIdentity{
					Name: &instanceProfile,
				},
				Zone: &vpcv1.ZoneIdentity{
					Name: zone,
				},
				PrimaryNetworkInterface: &vpcv1.NetworkInterfacePrototype{
					Subnet: &vpcv1.SubnetIdentity{
						ID: &subnetID,
					},
				},
				Keys: []vpcv1.KeyIdentityIntf{
					&vpcv1.KeyIdentity{
						ID: &keyID,
					},
				},
				VPC: &vpcv1.VPCIdentity{
					ID: &vpcID,
				},
			})
			instanceTemplate, response, err := vpcService.CreateInstanceTemplate(options)

			// end-create_instance_template
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(instanceTemplate).ToNot(BeNil())
			instanceTemplateID = *instanceTemplate.(*vpcv1.InstanceTemplate).ID
		})
		It(`GetInstanceTemplate request example`, func() {
			fmt.Println("\nGetInstanceTemplate() result:")
			// begin-get_instance_template

			options := &vpcv1.GetInstanceTemplateOptions{}
			options.SetID(instanceTemplateID)
			instanceTemplate, response, err := vpcService.GetInstanceTemplate(options)

			// end-get_instance_template
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceTemplate).ToNot(BeNil())

		})
		It(`UpdateInstanceTemplate request example`, func() {
			fmt.Println("\nUpdateInstanceTemplate() result:")
			name := getName("template")
			// begin-update_instance_template

			options := &vpcv1.UpdateInstanceTemplateOptions{}
			options.SetID(instanceTemplateID)
			instanceTemplatePatchModel := &vpcv1.InstanceTemplatePatch{
				Name: &name,
			}
			instanceTemplatePatch, asPatchErr := instanceTemplatePatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.InstanceTemplatePatch = instanceTemplatePatch
			instanceTemplate, response, err := vpcService.UpdateInstanceTemplate(options)

			// end-update_instance_template
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceTemplate).ToNot(BeNil())

		})
		It(`ListInstances request example`, func() {
			fmt.Println("\nListInstances() result:")
			// begin-list_instances

			options := &vpcv1.ListInstancesOptions{}
			instances, response, err := vpcService.ListInstances(options)

			// end-list_instances
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instances).ToNot(BeNil())

		})
		It(`CreateInstance request example`, func() {
			fmt.Println("\nCreateInstance() result:")
			crn := "crn:[...]"
			// begin-create_instance
			keyIDentityModel := &vpcv1.KeyIdentityByID{
				ID: &keyID,
			}
			instanceProfileIdentityModel := &vpcv1.InstanceProfileIdentityByName{
				Name: &[]string{"bx2d-2x8"}[0],
			}
			encryptionKeyIdentityModel := &vpcv1.EncryptionKeyIdentityByCRN{
				CRN: &crn,
			}
			volumeProfileIdentityModel := &vpcv1.VolumeProfileIdentityByName{
				Name: &[]string{"5iops-tier"}[0],
			}
			volumeAttachmentPrototypeVolumeModel := &vpcv1.VolumeAttachmentVolumePrototypeInstanceContextVolumePrototypeInstanceContext{
				Name:          &[]string{"my-instance-modified"}[0],
				Capacity:      &[]int64{100}[0],
				EncryptionKey: encryptionKeyIdentityModel,
				Profile:       volumeProfileIdentityModel,
			}
			volumeAttachmentPrototypeModel := &vpcv1.VolumeAttachmentPrototypeInstanceContext{
				Volume: volumeAttachmentPrototypeVolumeModel,
			}
			vpcIDentityModel := &vpcv1.VPCIdentityByID{
				ID: &vpcID,
			}
			imageIDentityModel := &vpcv1.ImageIdentityByID{
				ID: &imageID,
			}
			subnetIDentityModel := &vpcv1.SubnetIdentityByID{
				ID: &subnetID,
			}
			networkInterfacePrototypeModel := &vpcv1.NetworkInterfacePrototype{
				Name:   &[]string{"my-instance-modified"}[0],
				Subnet: subnetIDentityModel,
			}
			zoneIdentityModel := &vpcv1.ZoneIdentityByName{
				Name: zone,
			}
			instancePrototypeModel := &vpcv1.InstancePrototypeInstanceByImage{
				Keys:                    []vpcv1.KeyIdentityIntf{keyIDentityModel},
				Name:                    &[]string{"my-instance-modified"}[0],
				Profile:                 instanceProfileIdentityModel,
				VolumeAttachments:       []vpcv1.VolumeAttachmentPrototypeInstanceContext{*volumeAttachmentPrototypeModel},
				VPC:                     vpcIDentityModel,
				Image:                   imageIDentityModel,
				PrimaryNetworkInterface: networkInterfacePrototypeModel,
				Zone:                    zoneIdentityModel,
			}
			createInstanceOptions := vpcService.NewCreateInstanceOptions(
				instancePrototypeModel,
			)
			instance, response, err := vpcService.CreateInstance(createInstanceOptions)

			// end-create_instance
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(instance).ToNot(BeNil())
			instanceID = *instance.ID
		})
		It(`GetInstance request example`, func() {
			fmt.Println("\nGetInstance() result:")
			// begin-get_instance

			options := &vpcv1.GetInstanceOptions{}
			options.SetID(instanceID)
			instance, response, err := vpcService.GetInstance(options)

			// end-get_instance
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instance).ToNot(BeNil())

		})
		It(`UpdateInstance request example`, func() {
			fmt.Println("\nUpdateInstance() result:")
			// begin-update_instance

			options := &vpcv1.UpdateInstanceOptions{
				ID: &instanceID,
			}
			instancePatchModel := &vpcv1.InstancePatch{
				Name: &[]string{"my-instance-modified"}[0],
			}
			instancePatch, asPatchErr := instancePatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.InstancePatch = instancePatch
			instance, response, err := vpcService.UpdateInstance(options)

			// end-update_instance
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instance).ToNot(BeNil())

		})
		It(`GetInstanceInitialization request example`, func() {
			fmt.Println("\nGetInstanceInitialization() result:")
			// begin-get_instance_initialization
			options := &vpcv1.GetInstanceInitializationOptions{}
			options.SetID(instanceID)
			initData, response, err := vpcService.GetInstanceInitialization(options)

			// end-get_instance_initialization
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(initData).ToNot(BeNil())

		})
		It(`CreateInstanceAction request example`, func() {
			fmt.Println("\nCreateInstanceAction() result:")
			// begin-create_instance_action

			options := &vpcv1.CreateInstanceActionOptions{}
			options.SetInstanceID(instanceID)
			options.SetType("stop")
			action, response, err := vpcService.CreateInstanceAction(options)

			// end-create_instance_action
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(action).ToNot(BeNil())

		})
		It(`CreateInstanceConsoleAccessToken request example`, func() {
			Skip("not runnin with mock")
			fmt.Println("\nCreateInstanceConsoleAccessToken() result:")
			// begin-create_instance_console_access_token
			options := &vpcv1.CreateInstanceConsoleAccessTokenOptions{
				InstanceID:  &instanceID,
				ConsoleType: &[]string{"serial"}[0],
			}

			instanceConsoleAccessToken, response, err :=
				vpcService.CreateInstanceConsoleAccessToken(options)

			// end-create_instance_console_access_token
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceConsoleAccessToken).ToNot(BeNil())

		})
		It(`ListInstanceDisks request example`, func() {
			fmt.Println("\nListInstanceDisks() result:")
			// begin-list_instance_disks

			listInstanceDisksOptions := vpcService.NewListInstanceDisksOptions(
				instanceID,
			)
			instanceDisksCollection, response, err :=
				vpcService.ListInstanceDisks(listInstanceDisksOptions)

			// end-list_instance_disks
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceDisksCollection).ToNot(BeNil())
			diskID = *instanceDisksCollection.Disks[0].ID
		})
		It(`GetInstanceDisk request example`, func() {
			fmt.Println("\nGetInstanceDisk() result:")
			// begin-get_instance_disk

			options := vpcService.NewGetInstanceDiskOptions(
				instanceID,
				diskID,
			)
			instanceDisk, response, err := vpcService.GetInstanceDisk(options)

			// end-get_instance_disk
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceDisk).ToNot(BeNil())

		})
		It(`UpdateInstanceDisk request example`, func() {
			fmt.Println("\nUpdateInstanceDisk() result:")
			name := getName("disk")
			// begin-update_instance_disk

			instanceDiskPatchModel := &vpcv1.InstanceDiskPatch{
				Name: &name,
			}
			instanceDiskPatchModelAsPatch, asPatchErr := instanceDiskPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options := vpcService.NewUpdateInstanceDiskOptions(
				instanceID,
				diskID,
				instanceDiskPatchModelAsPatch,
			)
			instanceDisk, response, err := vpcService.UpdateInstanceDisk(options)

			// end-update_instance_disk
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceDisk).ToNot(BeNil())

		})
		It(`ListInstanceNetworkInterfaces request example`, func() {
			fmt.Println("\nListInstanceNetworkInterfaces() result:")
			// begin-list_instance_network_interfaces

			options := &vpcv1.ListInstanceNetworkInterfacesOptions{}
			options.SetInstanceID(instanceID)
			networkInterfaces, response, err := vpcService.ListInstanceNetworkInterfaces(options)

			// end-list_instance_network_interfaces
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkInterfaces).ToNot(BeNil())

		})
		It(`CreateInstanceNetworkInterface request example`, func() {
			fmt.Println("\nCreateInstanceNetworkInterface() result:")
			// begin-create_instance_network_interface

			options := &vpcv1.CreateInstanceNetworkInterfaceOptions{}
			options.SetInstanceID(instanceID)
			options.SetName("eth1")
			options.SetSubnet(&vpcv1.SubnetIdentityByID{
				ID: &subnetID,
			})
			networkInterface, response, err := vpcService.CreateInstanceNetworkInterface(options)

			// end-create_instance_network_interface
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(networkInterface).ToNot(BeNil())
			eth2ID = *networkInterface.ID
		})
		It(`GetInstanceNetworkInterface request example`, func() {
			fmt.Println("\nGetInstanceNetworkInterface() result:")
			// begin-get_instance_network_interface

			options := &vpcv1.GetInstanceNetworkInterfaceOptions{}
			options.SetID(eth2ID)
			options.SetInstanceID(instanceID)
			networkInterface, response, err := vpcService.GetInstanceNetworkInterface(options)

			// end-get_instance_network_interface
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkInterface).ToNot(BeNil())

		})
		It(`UpdateInstanceNetworkInterface request example`, func() {
			fmt.Println("\nUpdateInstanceNetworkInterface() result:")
			name := getName("nic")
			ipSpoofing := true
			// begin-update_instance_network_interface

			options := &vpcv1.UpdateInstanceNetworkInterfaceOptions{
				InstanceID: &instanceID,
				ID:         &eth2ID,
			}
			instancePatchModel := &vpcv1.NetworkInterfacePatch{
				Name:            &name,
				AllowIPSpoofing: &ipSpoofing,
			}
			networkInterfacePatch, asPatchErr := instancePatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.NetworkInterfacePatch = networkInterfacePatch
			networkInterface, response, err := vpcService.UpdateInstanceNetworkInterface(options)

			// end-update_instance_network_interface
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkInterface).ToNot(BeNil())

		})
		It(`AddInstanceNetworkInterfaceFloatingIP request example`, func() {
			fmt.Println("\nAddInstanceNetworkInterfaceFloatingIP() result:")
			// begin-add_instance_network_interface_floating_ip

			options := &vpcv1.AddInstanceNetworkInterfaceFloatingIPOptions{}
			options.SetID(floatingIPID)
			options.SetInstanceID(instanceID)
			options.SetNetworkInterfaceID(eth2ID)
			floatingIP, response, err :=
				vpcService.AddInstanceNetworkInterfaceFloatingIP(options)

			// end-add_instance_network_interface_floating_ip
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(floatingIP).ToNot(BeNil())

		})
		It(`ListInstanceNetworkInterfaceFloatingIps request example`, func() {
			fmt.Println("\nListInstanceNetworkInterfaceFloatingIps() result:")
			// begin-list_instance_network_interface_floating_ips

			options := &vpcv1.ListInstanceNetworkInterfaceFloatingIpsOptions{}
			options.SetInstanceID(instanceID)
			options.SetNetworkInterfaceID(eth2ID)
			floatingIPs, response, err :=
				vpcService.ListInstanceNetworkInterfaceFloatingIps(options)

			// end-list_instance_network_interface_floating_ips
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(floatingIPs).ToNot(BeNil())

		})

		It(`GetInstanceNetworkInterfaceFloatingIP request example`, func() {
			fmt.Println("\nGetInstanceNetworkInterfaceFloatingIP() result:")
			// begin-get_instance_network_interface_floating_ip

			options := &vpcv1.GetInstanceNetworkInterfaceFloatingIPOptions{}
			options.SetID(floatingIPID)
			options.SetInstanceID(instanceID)
			options.SetNetworkInterfaceID(eth2ID)
			floatingIP, response, err :=
				vpcService.GetInstanceNetworkInterfaceFloatingIP(options)

			// end-get_instance_network_interface_floating_ip
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(floatingIP).ToNot(BeNil())

		})

		It(`ListInstanceVolumeAttachments request example`, func() {
			fmt.Println("\nListInstanceVolumeAttachments() result:")
			// begin-list_instance_volume_attachments

			options := &vpcv1.ListInstanceVolumeAttachmentsOptions{}
			options.SetInstanceID(instanceID)
			volumeAttachments, response, err := vpcService.ListInstanceVolumeAttachments(
				options)

			// end-list_instance_volume_attachments
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeAttachments).ToNot(BeNil())

		})
		It(`CreateInstanceVolumeAttachment request example`, func() {
			fmt.Println("\nCreateInstanceVolumeAttachment() result:")
			// begin-create_instance_volume_attachment

			volumeAttachmentPrototypeVolumeModel := &vpcv1.VolumeAttachmentPrototypeVolumeVolumeIdentityVolumeIdentityByID{
				ID: &volumeID,
			}

			options := vpcService.NewCreateInstanceVolumeAttachmentOptions(
				instanceID,
				volumeAttachmentPrototypeVolumeModel,
			)

			volumeAttachment, response, err := vpcService.CreateInstanceVolumeAttachment(options)

			// end-create_instance_volume_attachment
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(volumeAttachment).ToNot(BeNil())
			volumeAttachmentID = *volumeAttachment.ID
		})
		It(`GetInstanceVolumeAttachment request example`, func() {
			fmt.Println("\nGetInstanceVolumeAttachment() result:")
			// begin-get_instance_volume_attachment

			options := &vpcv1.GetInstanceVolumeAttachmentOptions{}
			options.SetInstanceID(instanceID)
			options.SetID(volumeAttachmentID)
			volumeAttachment, response, err := vpcService.GetInstanceVolumeAttachment(options)

			// end-get_instance_volume_attachment
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeAttachment).ToNot(BeNil())

		})
		It(`UpdateInstanceVolumeAttachment request example`, func() {
			fmt.Println("\nUpdateInstanceVolumeAttachment() result:")
			name := getName("vol-att")
			// begin-update_instance_volume_attachment

			options := &vpcv1.UpdateInstanceVolumeAttachmentOptions{}
			volumeAttachmentPatchModel := &vpcv1.VolumeAttachmentPatch{
				Name: &name,
			}
			volumeAttachmentPatch, asPatchErr := volumeAttachmentPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.SetInstanceID(instanceID)
			options.SetID(volumeAttachmentID)
			options.SetVolumeAttachmentPatch(volumeAttachmentPatch)
			volumeAttachment, response, err := vpcService.UpdateInstanceVolumeAttachment(options)

			// end-update_instance_volume_attachment
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeAttachment).ToNot(BeNil())

		})
		It(`ListInstanceGroups request example`, func() {
			fmt.Println("\nListInstanceGroups() result:")
			// begin-list_instance_groups

			options := &vpcv1.ListInstanceGroupsOptions{}
			instanceGroups, response, err := vpcService.ListInstanceGroups(options)

			// end-list_instance_groups
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroups).ToNot(BeNil())

		})
		It(`CreateInstanceGroup request example`, func() {
			fmt.Println("\nCreateInstanceGroup() result:")
			name := getName("ig")
			// begin-create_instance_group

			options := &vpcv1.CreateInstanceGroupOptions{
				InstanceTemplate: &vpcv1.InstanceTemplateIdentity{
					ID: &instanceTemplateID,
				},
			}
			options.SetName(name)
			var subnetArray = []vpcv1.SubnetIdentityIntf{
				&vpcv1.SubnetIdentity{
					ID: &subnetID,
				},
			}
			options.SetSubnets(subnetArray)
			instanceGroup, response, err := vpcService.CreateInstanceGroup(options)

			// end-create_instance_group
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(instanceGroup).ToNot(BeNil())
			instanceGroupID = *instanceGroup.ID
		})
		It(`GetInstanceGroup request example`, func() {
			fmt.Println("\nGetInstanceGroup() result:")
			// begin-get_instance_group

			options := &vpcv1.GetInstanceGroupOptions{}
			options.SetID(instanceGroupID)
			instanceGroup, response, err := vpcService.GetInstanceGroup(options)

			// end-get_instance_group
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroup).ToNot(BeNil())

		})
		It(`UpdateInstanceGroup request example`, func() {
			fmt.Println("\nUpdateInstanceGroup() result:")
			name := getName("ig")
			// begin-update_instance_group

			options := &vpcv1.UpdateInstanceGroupOptions{}
			options.SetID(instanceGroupID)
			instanceGroupPatchModel := vpcv1.InstanceGroupPatch{}
			instanceGroupPatchModel.Name = &name
			instanceGroupPatchModel.InstanceTemplate = &vpcv1.InstanceTemplateIdentity{
				ID: &instanceTemplateID,
			}
			instanceGroupPatchModel.MembershipCount = &[]int64{5}[0]
			instanceGroupPatch, asPatchErr := instanceGroupPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.InstanceGroupPatch = instanceGroupPatch
			instanceGroup, response, err := vpcService.UpdateInstanceGroup(options)

			// end-update_instance_group
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroup).ToNot(BeNil())

		})
		It(`ListInstanceGroupManagers request example`, func() {
			fmt.Println("\nListInstanceGroupManagers() result:")
			// begin-list_instance_group_managers

			options := &vpcv1.ListInstanceGroupManagersOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			instanceGroupManagers, response, err :=
				vpcService.ListInstanceGroupManagers(options)

			// end-list_instance_group_managers
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManagers).ToNot(BeNil())

		})
		It(`CreateInstanceGroupManager request example`, func() {
			fmt.Println("\nCreateInstanceGroupManager() result:")
			// begin-create_instance_group_manager

			prototype := &vpcv1.InstanceGroupManagerPrototype{
				ManagerType:        &[]string{"autoscale"}[0],
				MaxMembershipCount: &[]int64{5}[0],
			}
			options := &vpcv1.CreateInstanceGroupManagerOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			options.SetInstanceGroupManagerPrototype(prototype)
			instanceGroupManagerIntf, response, err :=
				vpcService.CreateInstanceGroupManager(options)
			instanceGroupManager := instanceGroupManagerIntf.(*vpcv1.InstanceGroupManager)

			// end-create_instance_group_manager
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(instanceGroupManager).ToNot(BeNil())
			instanceGroupManagerID = *instanceGroupManager.ID
		})
		It(`GetInstanceGroupManager request example`, func() {
			fmt.Println("\nGetInstanceGroupManager() result:")
			// begin-get_instance_group_manager

			options := &vpcv1.GetInstanceGroupManagerOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			options.SetID(instanceGroupManagerID)
			instanceGroupManager, response, err := vpcService.GetInstanceGroupManager(options)

			// end-get_instance_group_manager
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManager).ToNot(BeNil())

		})
		It(`UpdateInstanceGroupManager request example`, func() {
			fmt.Println("\nUpdateInstanceGroupManager() result:")
			name := getName("manager")
			// begin-update_instance_group_manager
			instanceGroupManagerPatchModel := &vpcv1.InstanceGroupManagerPatch{}
			instanceGroupManagerPatchModel.Name = &name
			instanceGroupManagerPatch, asPatchErr := instanceGroupManagerPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options := &vpcv1.UpdateInstanceGroupManagerOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			options.SetID(instanceGroupManagerID)
			options.InstanceGroupManagerPatch = instanceGroupManagerPatch
			instanceGroupManager, response, err :=
				vpcService.UpdateInstanceGroupManager(options)

			// end-update_instance_group_manager
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManager).ToNot(BeNil())

		})
		It(`ListInstanceGroupManagerActions request example`, func() {
			fmt.Println("\nListInstanceGroupManagerActions() result:")
			// begin-list_instance_group_manager_actions

			options := &vpcv1.ListInstanceGroupManagerActionsOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			options.SetInstanceGroupManagerID(instanceGroupManagerID)
			instanceGroupManagerActions, response, err :=
				vpcService.ListInstanceGroupManagerActions(options)

			// end-list_instance_group_manager_actions
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManagerActions).ToNot(BeNil())

		})
		It(`CreateInstanceGroupManagerAction request example`, func() {
			fmt.Println("\nCreateInstanceGroupManagerAction() result:")
			name := getName("igAction")
			// begin-create_instance_group_manager_action

			cronSpec := &[]string{"*/5 1,2,3 * * *"}[0]
			instanceGroupManagerScheduledActionGroupPrototypeModel :=
				&vpcv1.InstanceGroupManagerScheduledActionGroupPrototype{
					MembershipCount: &[]int64{5}[0],
				}
			instanceGroupManagerActionPrototypeModel :=
				&vpcv1.InstanceGroupManagerActionPrototypeScheduledActionPrototype{
					Name:     &name,
					CronSpec: cronSpec,
					Group:    instanceGroupManagerScheduledActionGroupPrototypeModel,
				}
			createInstanceGroupManagerActionOptions :=
				&vpcv1.CreateInstanceGroupManagerActionOptions{
					InstanceGroupID:                     &instanceGroupID,
					InstanceGroupManagerID:              &instanceGroupManagerID,
					InstanceGroupManagerActionPrototype: instanceGroupManagerActionPrototypeModel,
				}
			instanceGroupManagerActionIntf, response, err :=
				vpcService.CreateInstanceGroupManagerAction(
					createInstanceGroupManagerActionOptions,
				)
			instanceGroupManagerAction := instanceGroupManagerActionIntf.(*vpcv1.InstanceGroupManagerAction)
			// end-create_instance_group_manager_action
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(instanceGroupManagerAction).ToNot(BeNil())
			instanceGroupManagerActionID = *instanceGroupManagerAction.ID
		})
		It(`GetInstanceGroupManagerAction request example`, func() {
			fmt.Println("\nGetInstanceGroupManagerAction() result:")
			// begin-get_instance_group_manager_action

			options := &vpcv1.GetInstanceGroupManagerActionOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			options.SetInstanceGroupManagerID(instanceGroupManagerID)
			options.SetID(instanceGroupManagerActionID)
			instanceGroupManagerAction, response, err :=
				vpcService.GetInstanceGroupManagerAction(options)

			// end-get_instance_group_manager_action
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManagerAction).ToNot(BeNil())

		})
		It(`UpdateInstanceGroupManagerAction request example`, func() {
			fmt.Println("\nUpdateInstanceGroupManagerAction() result:")
			name := getName("igManager")
			// begin-update_instance_group_manager_action
			cronSpec := &[]string{"*/5 1,2,3 * * *"}[0]
			instanceGroupManagerScheduledActionGroupPatchModel :=
				&vpcv1.InstanceGroupManagerActionGroupPatch{
					MembershipCount: &[]int64{5}[0],
				}
			instanceGroupManagerActionPatchModel :=
				&vpcv1.InstanceGroupManagerActionPatch{
					Name:     &name,
					CronSpec: cronSpec,
					Group:    instanceGroupManagerScheduledActionGroupPatchModel,
				}
			instanceGroupManagerActionPatchModelAsPatch, asPatchErr :=
				instanceGroupManagerActionPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options :=
				&vpcv1.UpdateInstanceGroupManagerActionOptions{
					InstanceGroupID:                 &instanceGroupID,
					InstanceGroupManagerID:          &instanceGroupManagerID,
					ID:                              &instanceGroupManagerActionID,
					InstanceGroupManagerActionPatch: instanceGroupManagerActionPatchModelAsPatch,
				}

			instanceGroupManagerAction, response, err := vpcService.UpdateInstanceGroupManagerAction(options)

			// end-update_instance_group_manager_action
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManagerAction).ToNot(BeNil())

		})
		It(`ListInstanceGroupManagerPolicies request example`, func() {
			fmt.Println("\nListInstanceGroupManagerPolicies() result:")
			// begin-list_instance_group_manager_policies

			options := &vpcv1.ListInstanceGroupManagerPoliciesOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			options.SetInstanceGroupManagerID(instanceGroupManagerID)
			instanceGroupManagerPolicies, response, err :=
				vpcService.ListInstanceGroupManagerPolicies(options)

			// end-list_instance_group_manager_policies
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManagerPolicies).ToNot(BeNil())

		})
		It(`CreateInstanceGroupManagerPolicy request example`, func() {
			fmt.Println("\nCreateInstanceGroupManagerPolicy() result:")
			// begin-create_instance_group_manager_policy

			prototype := &vpcv1.InstanceGroupManagerPolicyPrototype{
				PolicyType:  &[]string{"target"}[0],
				MetricType:  &[]string{"cpu"}[0],
				MetricValue: &[]int64{20}[0],
			}
			options := &vpcv1.CreateInstanceGroupManagerPolicyOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			options.SetInstanceGroupManagerID(instanceGroupManagerID)
			options.SetInstanceGroupManagerPolicyPrototype(prototype)
			instanceGroupManagerPolicyIntf, response, err :=
				vpcService.CreateInstanceGroupManagerPolicy(options)
			instanceGroupManagerPolicy := instanceGroupManagerPolicyIntf.(*vpcv1.InstanceGroupManagerPolicy)
			// end-create_instance_group_manager_policy
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(instanceGroupManagerPolicy).ToNot(BeNil())
			instanceGroupManagerPolicyID = *instanceGroupManagerPolicy.ID
		})
		It(`GetInstanceGroupManagerPolicy request example`, func() {
			fmt.Println("\nGetInstanceGroupManagerPolicy() result:")
			// begin-get_instance_group_manager_policy

			options := &vpcv1.GetInstanceGroupManagerPolicyOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			options.SetInstanceGroupManagerID(instanceGroupManagerID)
			options.SetID(instanceGroupManagerPolicyID)
			instanceGroupManagerPolicy, response, err :=
				vpcService.GetInstanceGroupManagerPolicy(options)

			// end-get_instance_group_manager_policy
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManagerPolicy).ToNot(BeNil())

		})
		It(`UpdateInstanceGroupManagerPolicy request example`, func() {
			fmt.Println("\nUpdateInstanceGroupManagerPolicy() result:")
			name := getName("igPolicy")
			// begin-update_instance_group_manager_policy

			options := &vpcv1.UpdateInstanceGroupManagerPolicyOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			options.SetInstanceGroupManagerID(instanceGroupManagerID)
			options.SetID(instanceGroupManagerPolicyID)
			instanceGroupManagerPolicyPatchModel := &vpcv1.InstanceGroupManagerPolicyPatch{}
			instanceGroupManagerPolicyPatchModel.Name = &name
			instanceGroupManagerPolicyPatchModel.MetricType = &[]string{"cpu"}[0]
			instanceGroupManagerPolicyPatchModel.MetricValue = &[]int64{70}[0]
			instanceGroupManagerPolicyPatchModelAsPatch, asPatchErr :=
				instanceGroupManagerPolicyPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.InstanceGroupManagerPolicyPatch =
				instanceGroupManagerPolicyPatchModelAsPatch
			instanceGroupManagerPolicy, response, err :=
				vpcService.UpdateInstanceGroupManagerPolicy(options)

			// end-update_instance_group_manager_policy
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManagerPolicy).ToNot(BeNil())

		})
		It(`ListInstanceGroupMemberships request example`, func() {
			fmt.Println("\nListInstanceGroupMemberships() result:")
			// begin-list_instance_group_memberships

			options := &vpcv1.ListInstanceGroupMembershipsOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			instanceGroupMemberships, response, err :=
				vpcService.ListInstanceGroupMemberships(options)

			// end-list_instance_group_memberships
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupMemberships).ToNot(BeNil())
			instanceGroupMembershipID = *instanceGroupMemberships.Memberships[0].ID
		})
		It(`GetInstanceGroupMembership request example`, func() {
			fmt.Println("\nGetInstanceGroupMembership() result:")
			// begin-get_instance_group_membership

			options := &vpcv1.GetInstanceGroupMembershipOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			options.SetID(instanceGroupMembershipID)
			instanceGroupMembership, response, err :=
				vpcService.GetInstanceGroupMembership(options)

			// end-get_instance_group_membership
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupMembership).ToNot(BeNil())

		})
		It(`UpdateInstanceGroupMembership request example`, func() {
			fmt.Println("\nUpdateInstanceGroupMembership() result:")
			name := getName("membership")
			// begin-update_instance_group_membership

			options := &vpcv1.UpdateInstanceGroupMembershipOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			options.SetID(instanceGroupMembershipID)
			instanceGroupMembershipPatchModel := &vpcv1.InstanceGroupMembershipPatch{}
			instanceGroupMembershipPatchModel.Name = &name
			instanceGroupMembershipPatch, asPatchErr := instanceGroupMembershipPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.InstanceGroupMembershipPatch = instanceGroupMembershipPatch
			instanceGroupMembership, response, err :=
				vpcService.UpdateInstanceGroupMembership(options)

			// end-update_instance_group_membership
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupMembership).ToNot(BeNil())

		})
		It(`ListDedicatedHostGroups request example`, func() {
			fmt.Println("\nListDedicatedHostGroups() result:")
			// begin-list_dedicated_host_groups

			options := vpcService.NewListDedicatedHostGroupsOptions()
			dedicatedHostGroups, response, err :=
				vpcService.ListDedicatedHostGroups(options)

			// end-list_dedicated_host_groups
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHostGroups).ToNot(BeNil())

		})
		It(`CreateDedicatedHostGroup request example`, func() {
			fmt.Println("\nCreateDedicatedHostGroup() result:")
			name := getName("dhg")
			// begin-create_dedicated_host_group

			options := &vpcv1.CreateDedicatedHostGroupOptions{
				Name:   &name,
				Class:  &[]string{"mx2"}[0],
				Family: &[]string{"balanced"}[0],
				Zone: &vpcv1.ZoneIdentity{
					Name: zone,
				},
			}
			dedicatedHostGroup, response, err := vpcService.CreateDedicatedHostGroup(options)

			// end-create_dedicated_host_group
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(dedicatedHostGroup).ToNot(BeNil())
			dedicatedHostGroupID = *dedicatedHostGroup.ID
		})
		It(`GetDedicatedHostGroup request example`, func() {
			fmt.Println("\nGetDedicatedHostGroup() result:")
			// begin-get_dedicated_host_group

			options := vpcService.NewGetDedicatedHostGroupOptions(dedicatedHostGroupID)
			dedicatedHostGroup, response, err := vpcService.GetDedicatedHostGroup(options)

			// end-get_dedicated_host_group
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHostGroup).ToNot(BeNil())

		})
		It(`UpdateDedicatedHostGroup request example`, func() {
			fmt.Println("\nUpdateDedicatedHostGroup() result:")
			name := getName("dhg")
			// begin-update_dedicated_host_group

			dedicatedHostGroupPatchModel := &vpcv1.DedicatedHostGroupPatch{
				Name: &name,
			}
			dedicatedHostGroupPatchModelAsPatch, asPatchErr := dedicatedHostGroupPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}

			updateDedicatedHostGroupOptions := vpcService.NewUpdateDedicatedHostGroupOptions(
				dedicatedHostGroupID,
				dedicatedHostGroupPatchModelAsPatch,
			)

			dedicatedHostGroup, response, err := vpcService.UpdateDedicatedHostGroup(updateDedicatedHostGroupOptions)

			// end-update_dedicated_host_group
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHostGroup).ToNot(BeNil())

		})
		It(`ListDedicatedHostProfiles request example`, func() {
			fmt.Println("\nListDedicatedHostProfiles() result:")
			// begin-list_dedicated_host_profiles

			options := &vpcv1.ListDedicatedHostProfilesOptions{}
			profiles, response, err := vpcService.ListDedicatedHostProfiles(options)

			// end-list_dedicated_host_profiles
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profiles).ToNot(BeNil())
			dhProfile = *profiles.Profiles[0].Name
		})
		It(`GetDedicatedHostProfile request example`, func() {
			fmt.Println("\nGetDedicatedHostProfile() result:")
			// begin-get_dedicated_host_profile

			options := &vpcv1.GetDedicatedHostProfileOptions{}
			options.SetName(dhProfile)
			profile, response, err := vpcService.GetDedicatedHostProfile(options)

			// end-get_dedicated_host_profile
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profile).ToNot(BeNil())

		})
		It(`ListDedicatedHosts request example`, func() {
			fmt.Println("\nListDedicatedHosts() result:")
			// begin-list_dedicated_hosts

			options := vpcService.NewListDedicatedHostsOptions()
			dedicatedHosts, response, err :=
				vpcService.ListDedicatedHosts(options)

			// end-list_dedicated_hosts
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHosts).ToNot(BeNil())

		})
		It(`CreateDedicatedHost request example`, func() {
			fmt.Println("\nCreateDedicatedHost() result:")
			name := getName("dh")
			// begin-create_dedicated_host

			options := &vpcv1.CreateDedicatedHostOptions{}
			options.SetDedicatedHostPrototype(&vpcv1.DedicatedHostPrototype{
				Name: &name,
				Profile: &vpcv1.DedicatedHostProfileIdentity{
					Name: &dhProfile,
				},
				Group: &vpcv1.DedicatedHostGroupIdentity{
					ID: &dedicatedHostGroupID,
				},
			})
			dedicatedHost, response, err := vpcService.CreateDedicatedHost(options)

			// end-create_dedicated_host
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(dedicatedHost).ToNot(BeNil())
			dedicatedHostID = *dedicatedHost.ID
		})
		It(`GetDedicatedHost request example`, func() {
			fmt.Println("\nGetDedicatedHost() result:")
			// begin-get_dedicated_host

			options := vpcService.NewGetDedicatedHostOptions(dedicatedHostID)
			dedicatedHost, response, err := vpcService.GetDedicatedHost(options)

			// end-get_dedicated_host
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHost).ToNot(BeNil())

		})
		It(`ListDedicatedHostDisks request example`, func() {
			fmt.Println("\nListDedicatedHostDisks() result:")
			options := vpcService.NewListDedicatedHostsOptions()
			dedicatedHosts, response, err :=
				vpcService.ListDedicatedHosts(options)
			for i := range dedicatedHosts.DedicatedHosts {
				if len(dedicatedHosts.DedicatedHosts[i].Disks) > 0 {
					dhID = *dedicatedHosts.DedicatedHosts[i].ID
					break
				}
			}
			// begin-list_dedicated_host_disks

			listDedicatedHostDisksOptions := vpcService.NewListDedicatedHostDisksOptions(
				dhID,
			)
			dedicatedHostDiskCollection, response, err :=
				vpcService.ListDedicatedHostDisks(listDedicatedHostDisksOptions)

			// end-list_dedicated_host_disks
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHostDiskCollection).ToNot(BeNil())
			diskID = *dedicatedHostDiskCollection.Disks[0].ID
		})
		It(`GetDedicatedHostDisk request example`, func() {
			fmt.Println("\nGetDedicatedHostDisk() result:")
			// begin-get_dedicated_host_disk

			getDedicatedHostDiskOptions := vpcService.NewGetDedicatedHostDiskOptions(
				dhID,
				diskID,
			)
			dedicatedHostDisk, response, err :=
				vpcService.GetDedicatedHostDisk(getDedicatedHostDiskOptions)

			// end-get_dedicated_host_disk
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHostDisk).ToNot(BeNil())

		})
		It(`UpdateDedicatedHostDisk request example`, func() {
			fmt.Println("\nUpdateDedicatedHostDisk() result:")
			name := getName("dhdisk")
			// begin-update_dedicated_host_disk

			dedicatedHostDiskPatchModel := &vpcv1.DedicatedHostDiskPatch{
				Name: &name,
			}
			dedicatedHostDiskPatchModelAsPatch, asPatchErr := dedicatedHostDiskPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options := vpcService.NewUpdateDedicatedHostDiskOptions(
				dhID,
				diskID,
				dedicatedHostDiskPatchModelAsPatch,
			)
			dedicatedHostDisk, response, err := vpcService.UpdateDedicatedHostDisk(options)

			// end-update_dedicated_host_disk
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHostDisk).ToNot(BeNil())

		})
		It(`UpdateDedicatedHost request example`, func() {
			fmt.Println("\nUpdateDedicatedHost() result:")
			name := getName("dh")
			// begin-update_dedicated_host
			options := &vpcv1.UpdateDedicatedHostOptions{
				ID: &dedicatedHostID,
			}
			dedicatedHostPatchModel := &vpcv1.DedicatedHostPatch{
				Name: &name,
			}
			dedicatedHostPatchModelAsPatch, asPatchErr := dedicatedHostPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.DedicatedHostPatch = dedicatedHostPatchModelAsPatch
			dedicatedHost, response, err := vpcService.UpdateDedicatedHost(options)
			// end-update_dedicated_host
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHost).ToNot(BeNil())

		})
		It(`ListVolumeProfiles request example`, func() {
			fmt.Println("\nListVolumeProfiles() result:")
			// begin-list_volume_profiles

			options := &vpcv1.ListVolumeProfilesOptions{}
			profiles, response, err := vpcService.ListVolumeProfiles(options)

			// end-list_volume_profiles
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profiles).ToNot(BeNil())

		})
		It(`GetVolumeProfile request example`, func() {
			fmt.Println("\nGetVolumeProfile() result:")
			// begin-get_volume_profile

			options := &vpcv1.GetVolumeProfileOptions{}
			options.SetName("10iops-tier")
			profile, response, err := vpcService.GetVolumeProfile(options)

			// end-get_volume_profile
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profile).ToNot(BeNil())

		})

		It(`ListSnapshots request example`, func() {
			fmt.Println("\nListSnapshots() result:")
			// begin-list_snapshots

			options := &vpcv1.ListSnapshotsOptions{}
			snapshotCollection, response, err := vpcService.ListSnapshots(options)

			// end-list_snapshots
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(snapshotCollection).ToNot(BeNil())

		})
		It(`CreateSnapshot request example`, func() {
			fmt.Println("\nCreateSnapshot() result:")
			name := getName("snapshotone")
			volumeIdentityModel := &vpcv1.VolumeIdentityByID{
				ID: &volumeID,
			}
			secondSnap := &vpcv1.SnapshotPrototypeSnapshotBySourceVolume{
				Name:         &name,
				SourceVolume: volumeIdentityModel,
			}
			secondCreateSnapshotOptions := vpcService.NewCreateSnapshotOptions(
				secondSnap,
			)
			_, _, err := vpcService.CreateSnapshot(secondCreateSnapshotOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			name = getName("snapshottwo")
			// begin-create_snapshot
			options := &vpcv1.SnapshotPrototypeSnapshotBySourceVolume{
				Name:         &name,
				SourceVolume: volumeIdentityModel,
			}
			createSnapshotOptions := vpcService.NewCreateSnapshotOptions(
				options,
			)
			snapshot, response, err := vpcService.CreateSnapshot(createSnapshotOptions)

			// end-create_snapshot
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(snapshot).ToNot(BeNil())
			snapshotID = *snapshot.ID
		})
		It(`GetSnapshot request example`, func() {
			fmt.Println("\nGetSnapshot() result:")
			// begin-get_snapshot

			options := &vpcv1.GetSnapshotOptions{
				ID: &snapshotID,
			}
			snapshot, response, err := vpcService.GetSnapshot(options)

			// end-get_snapshot
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(snapshot).ToNot(BeNil())
			ifMatchSnapshot = response.GetHeaders()["Etag"][0]
		})
		It(`UpdateSnapshot request example`, func() {
			fmt.Println("\nUpdateSnapshot() result:")
			name := getName("snapshot")
			userTags := []string{"usertag-snap-1"}
			// begin-update_snapshot

			snapshotPatchModel := &vpcv1.SnapshotPatch{
				Name:     &name,
				UserTags: userTags,
			}
			snapshotPatchModelAsPatch, asPatchErr := snapshotPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			updateSnapshotOptions := &vpcv1.UpdateSnapshotOptions{
				ID:            &snapshotID,
				SnapshotPatch: snapshotPatchModelAsPatch,
				IfMatch:       &ifMatchSnapshot,
			}
			snapshot, response, err := vpcService.UpdateSnapshot(updateSnapshotOptions)

			// end-update_snapshot
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())

			Expect(response.StatusCode).To(Equal(200))
			Expect(snapshot).ToNot(BeNil())

		})
		It(`ListRegions request example`, func() {
			fmt.Println("\nListRegions() result:")
			// begin-list_regions

			listRegionsOptions := &vpcv1.ListRegionsOptions{}
			regions, response, err := vpcService.ListRegions(listRegionsOptions)

			// end-list_regions
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(regions).ToNot(BeNil())

		})
		It(`GetRegion request example`, func() {
			fmt.Println("\nGetRegion() result:")
			// begin-get_region

			getRegionOptions := &vpcv1.GetRegionOptions{}
			getRegionOptions.SetName("us-east")
			region, response, err := vpcService.GetRegion(getRegionOptions)

			// end-get_region
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(region).ToNot(BeNil())

		})
		It(`ListRegionZones request example`, func() {
			fmt.Println("\nListRegionZones() result:")
			// begin-list_region_zones

			listZonesOptions := &vpcv1.ListRegionZonesOptions{}
			listZonesOptions.SetRegionName("us-east")
			zones, response, err := vpcService.ListRegionZones(listZonesOptions)
			// end-list_region_zones
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(zones).ToNot(BeNil())

		})
		It(`GetRegionZone request example`, func() {
			fmt.Println("\nGetRegionZone() result:")
			// begin-get_region_zone

			getZoneOptions := &vpcv1.GetRegionZoneOptions{}
			getZoneOptions.SetRegionName("us-east")
			getZoneOptions.SetName("us-east-1")
			zone, response, err := vpcService.GetRegionZone(getZoneOptions)

			// end-get_region_zone
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(zone).ToNot(BeNil())

		})
		It(`ListPublicGateways request example`, func() {
			fmt.Println("\nListPublicGateways() result:")
			// begin-list_public_gateways

			options := &vpcv1.ListPublicGatewaysOptions{}
			publicGateways, response, err := vpcService.ListPublicGateways(options)

			// end-list_public_gateways
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(publicGateways).ToNot(BeNil())

		})
		It(`CreatePublicGateway request example`, func() {
			fmt.Println("\nCreatePublicGateway() result:")
			// begin-create_public_gateway

			options := &vpcv1.CreatePublicGatewayOptions{}
			options.SetVPC(&vpcv1.VPCIdentity{
				ID: &vpcID,
			})
			options.SetZone(&vpcv1.ZoneIdentity{
				Name: zone,
			})
			publicGateway, response, err := vpcService.CreatePublicGateway(options)

			// end-create_public_gateway
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(publicGateway).ToNot(BeNil())
			publicGatewayID = *publicGateway.ID
		})
		It(`GetPublicGateway request example`, func() {
			fmt.Println("\nGetPublicGateway() result:")
			// begin-get_public_gateway

			options := &vpcv1.GetPublicGatewayOptions{}
			options.SetID(publicGatewayID)
			publicGateway, response, err := vpcService.GetPublicGateway(options)

			// end-get_public_gateway
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(publicGateway).ToNot(BeNil())

		})
		It(`UpdatePublicGateway request example`, func() {
			fmt.Println("\nUpdatePublicGateway() result:")
			name := getName("pgw")
			// begin-update_public_gateway

			options := &vpcv1.UpdatePublicGatewayOptions{}
			options.SetID(publicGatewayID)
			PublicGatewayPatchModel := &vpcv1.PublicGatewayPatch{
				Name: &name,
			}
			PublicGatewayPatch, asPatchErr := PublicGatewayPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.PublicGatewayPatch = PublicGatewayPatch
			publicGateway, response, err := vpcService.UpdatePublicGateway(options)
			// end-update_public_gateway
			if err != nil {
				panic(err)
			} // 	Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(publicGateway).ToNot(BeNil())

		})
		It(`ListNetworkAcls request example`, func() {
			fmt.Println("\nListNetworkAcls() result:")
			// begin-list_network_acls

			options := &vpcv1.ListNetworkAclsOptions{}
			networkACLCollection, response, err := vpcService.ListNetworkAcls(options)

			// end-list_network_acls
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkACLCollection).ToNot(BeNil())

		})
		It(`CreateNetworkACL request example`, func() {
			fmt.Println("\nCreateNetworkACL() result:")
			name := getName("acl")
			// begin-create_network_acl
			options := &vpcv1.CreateNetworkACLOptions{}
			options.SetNetworkACLPrototype(&vpcv1.NetworkACLPrototype{
				Name: &name,
				VPC: &vpcv1.VPCIdentity{
					ID: &vpcID,
				},
			})
			networkACL, response, err := vpcService.CreateNetworkACL(options)
			// end-create_network_acl
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(networkACL).ToNot(BeNil())
			networkACLID = *networkACL.ID
		})
		It(`GetNetworkACL request example`, func() {
			fmt.Println("\nGetNetworkACL() result:")
			// begin-get_network_acl

			options := &vpcv1.GetNetworkACLOptions{}
			options.SetID(networkACLID)
			networkACL, response, err := vpcService.GetNetworkACL(options)

			// end-get_network_acl
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkACL).ToNot(BeNil())

		})
		It(`UpdateNetworkACL request example`, func() {
			fmt.Println("\nUpdateNetworkACL() result:")
			name := getName("acl")
			// begin-update_network_acl

			options := &vpcv1.UpdateNetworkACLOptions{}
			options.SetID(networkACLID)
			networkACLPatchModel := &vpcv1.NetworkACLPatch{
				Name: &name,
			}
			networkACLPatch, asPatchErr := networkACLPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.NetworkACLPatch = networkACLPatch
			networkACL, response, err := vpcService.UpdateNetworkACL(options)

			// end-update_network_acl
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkACL).ToNot(BeNil())

		})
		It(`ListNetworkACLRules request example`, func() {
			fmt.Println("\nListNetworkACLRules() result:")
			// begin-list_network_acl_rules

			options := &vpcv1.ListNetworkACLRulesOptions{}
			options.SetNetworkACLID(networkACLID)
			networkACLRules, response, err := vpcService.ListNetworkACLRules(options)

			// end-list_network_acl_rules
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkACLRules).ToNot(BeNil())

		})
		It(`CreateNetworkACLRule request example`, func() {
			fmt.Println("\nCreateNetworkACLRule() result:")
			name := getName("aclrule")
			// begin-create_network_acl_rule
			options := &vpcv1.CreateNetworkACLRuleOptions{}
			options.SetNetworkACLID(networkACLID)
			options.SetNetworkACLRulePrototype(&vpcv1.NetworkACLRulePrototype{
				Action:      &[]string{"allow"}[0],
				Destination: &[]string{"192.168.3.2/32"}[0],
				Direction:   &[]string{"inbound"}[0],
				Source:      &[]string{"192.168.3.2/32"}[0],
				Protocol:    &[]string{"all"}[0],
				Name:        &name,
			})
			networkACLRuleIntf, response, err := vpcService.CreateNetworkACLRule(options)
			networkACLRule := networkACLRuleIntf.(*vpcv1.NetworkACLRuleNetworkACLRuleProtocolAll)
			// end-create_network_acl_rule
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(networkACLRule).ToNot(BeNil())
			networkACLRuleID = *networkACLRule.ID
		})
		It(`GetNetworkACLRule request example`, func() {
			fmt.Println("\nGetNetworkACLRule() result:")
			// begin-get_network_acl_rule

			options := &vpcv1.GetNetworkACLRuleOptions{}
			options.SetID(networkACLRuleID)
			options.SetNetworkACLID(networkACLID)
			networkACLRule, response, err := vpcService.GetNetworkACLRule(options)

			// end-get_network_acl_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkACLRule).ToNot(BeNil())

		})
		It(`UpdateNetworkACLRule request example`, func() {
			fmt.Println("\nUpdateNetworkACLRule() result:")
			name := getName("aclrule")
			// begin-update_network_acl_rule
			options := &vpcv1.UpdateNetworkACLRuleOptions{}
			options.SetID(networkACLRuleID)
			options.SetNetworkACLID(networkACLID)
			networkACLRulePatchModel := &vpcv1.NetworkACLRulePatch{
				Name: &name,
			}
			networkACLRulePatch, asPatchErr := networkACLRulePatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.NetworkACLRulePatch = networkACLRulePatch
			networkACLRule, response, err := vpcService.UpdateNetworkACLRule(options)
			// end-update_network_acl_rule
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkACLRule).ToNot(BeNil())

		})
		It(`ListSecurityGroups request example`, func() {
			fmt.Println("\nListSecurityGroups() result:")
			// begin-list_security_groups

			options := &vpcv1.ListSecurityGroupsOptions{}
			securityGroups, response, err := vpcService.ListSecurityGroups(options)

			// end-list_security_groups
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(securityGroups).ToNot(BeNil())

		})
		It(`CreateSecurityGroup request example`, func() {
			fmt.Println("\nCreateSecurityGroup() result:")
			name := getName("sg")
			// begin-create_security_group

			options := &vpcv1.CreateSecurityGroupOptions{}
			options.SetVPC(&vpcv1.VPCIdentity{
				ID: &vpcID,
			})
			options.SetName(name)
			securityGroup, response, err := vpcService.CreateSecurityGroup(options)

			// end-create_security_group
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(securityGroup).ToNot(BeNil())
			securityGroupID = *securityGroup.ID
		})
		It(`GetSecurityGroup request example`, func() {
			fmt.Println("\nGetSecurityGroup() result:")
			// begin-get_security_group

			options := &vpcv1.GetSecurityGroupOptions{}
			options.SetID(securityGroupID)
			securityGroup, response, err := vpcService.GetSecurityGroup(options)

			// end-get_security_group
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(securityGroup).ToNot(BeNil())

		})
		It(`UpdateSecurityGroup request example`, func() {
			fmt.Println("\nUpdateSecurityGroup() result:")
			name := getName("sg")
			// begin-update_security_group
			options := &vpcv1.UpdateSecurityGroupOptions{}
			options.SetID(securityGroupID)
			securityGroupPatchModel := &vpcv1.SecurityGroupPatch{
				Name: &name,
			}
			securityGroupPatch, asPatchErr := securityGroupPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.SecurityGroupPatch = securityGroupPatch
			securityGroup, response, err := vpcService.UpdateSecurityGroup(options)

			// end-update_security_group
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(securityGroup).ToNot(BeNil())

		})
		It(`ListSecurityGroupRules request example`, func() {
			fmt.Println("\nListSecurityGroupRules() result:")
			// begin-list_security_group_rules

			options := &vpcv1.ListSecurityGroupRulesOptions{}
			options.SetSecurityGroupID(securityGroupID)
			rules, response, err := vpcService.ListSecurityGroupRules(options)

			// end-list_security_group_rules
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rules).ToNot(BeNil())

		})
		It(`CreateSecurityGroupRule request example`, func() {
			fmt.Println("\nCreateSecurityGroupRule() result:")
			// begin-create_security_group_rule

			options := &vpcv1.CreateSecurityGroupRuleOptions{}
			options.SetSecurityGroupID(securityGroupID)
			options.SetSecurityGroupRulePrototype(&vpcv1.SecurityGroupRulePrototype{
				Direction: &[]string{"inbound"}[0],
				Protocol:  &[]string{"udp"}[0],
			})
			securityGroupRuleIntf, response, err := vpcService.CreateSecurityGroupRule(options)
			securityGroupRule := securityGroupRuleIntf.(*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolTcpudp)
			// end-create_security_group_rule
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(securityGroupRule).ToNot(BeNil())
			securityGroupRuleID = *securityGroupRule.ID
		})
		It(`GetSecurityGroupRule request example`, func() {
			fmt.Println("\nGetSecurityGroupRule() result:")
			// begin-get_security_group_rule

			options := &vpcv1.GetSecurityGroupRuleOptions{}
			options.SetSecurityGroupID(securityGroupID)
			options.SetID(securityGroupRuleID)
			securityGroupRule, response, err := vpcService.GetSecurityGroupRule(options)

			// end-get_security_group_rule
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(securityGroupRule).ToNot(BeNil())

		})
		It(`UpdateSecurityGroupRule request example`, func() {
			fmt.Println("\nUpdateSecurityGroupRule() result:")
			// begin-update_security_group_rule

			options := &vpcv1.UpdateSecurityGroupRuleOptions{}
			options.SecurityGroupID = &securityGroupID
			options.ID = &securityGroupRuleID
			securityGroupRulePatchModel := &vpcv1.SecurityGroupRulePatch{}
			securityGroupRulePatchModel.Direction = &[]string{"inbound"}[0]

			securityGroupRulePatch, asPatchErr := securityGroupRulePatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.SecurityGroupRulePatch = securityGroupRulePatch
			securityGroupRule, response, err := vpcService.UpdateSecurityGroupRule(options)

			// end-update_security_group_rule
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(securityGroupRule).ToNot(BeNil())

		})
		It(`ListSecurityGroupTargets request example`, func() {
			fmt.Println("\nListSecurityGroupTargets() result:")
			// begin-list_security_group_targets

			options := &vpcv1.ListSecurityGroupTargetsOptions{}
			options.SetSecurityGroupID(securityGroupID)
			targets, response, err := vpcService.ListSecurityGroupTargets(options)

			// end-list_security_group_targets
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(targets).ToNot(BeNil())
		})
		It(`CreateSecurityGroupTargetBinding request example`, func() {
			fmt.Println("\nCreateSecurityGroupTargetBinding() result:")
			// begin-create_security_group_target_binding

			options := vpcService.NewCreateSecurityGroupTargetBindingOptions(
				securityGroupID,
				eth2ID,
			)

			securityGroupTargetReferenceIntf, response, err := vpcService.CreateSecurityGroupTargetBinding(options)
			securityGroupTargetReference := securityGroupTargetReferenceIntf.(*vpcv1.SecurityGroupTargetReference)

			// end-create_security_group_target_binding
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(securityGroupTargetReference).ToNot(BeNil())
			targetID = *securityGroupTargetReference.ID
		})
		It(`GetSecurityGroupTarget request example`, func() {
			fmt.Println("\nGetSecurityGroupTarget() result:")
			// begin-get_security_group_target

			options := &vpcv1.GetSecurityGroupTargetOptions{}
			options.SetSecurityGroupID(securityGroupID)
			options.SetID(targetID)
			target, response, err :=
				vpcService.GetSecurityGroupTarget(options)

			// end-get_security_group_target
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())

		})

		It(`ListIkePolicies request example`, func() {
			fmt.Println("\nListIkePolicies() result:")
			// begin-list_ike_policies

			options := vpcService.NewListIkePoliciesOptions()
			ikePolicies, response, err := vpcService.ListIkePolicies(options)

			// end-list_ike_policies
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ikePolicies).ToNot(BeNil())

		})
		It(`CreateIkePolicy request example`, func() {
			fmt.Println("\nCreateIkePolicy() result:")
			name := getName("ike")
			// begin-create_ike_policy

			options := &vpcv1.CreateIkePolicyOptions{}
			options.SetName(name)
			options.SetAuthenticationAlgorithm("md5")
			options.SetDhGroup(2)
			options.SetEncryptionAlgorithm("aes128")
			options.SetIkeVersion(1)
			ikePolicy, response, err := vpcService.CreateIkePolicy(options)

			// end-create_ike_policy
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(ikePolicy).ToNot(BeNil())
			ikePolicyID = *ikePolicy.ID
		})
		It(`GetIkePolicy request example`, func() {
			fmt.Println("\nGetIkePolicy() result:")
			// begin-get_ike_policy

			options := vpcService.NewGetIkePolicyOptions(ikePolicyID)
			ikePolicy, response, err := vpcService.GetIkePolicy(options)

			// end-get_ike_policy
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ikePolicy).ToNot(BeNil())

		})
		It(`UpdateIkePolicy request example`, func() {
			fmt.Println("\nUpdateIkePolicy() result:")
			name := getName("ike")
			// begin-update_ike_policy

			options := &vpcv1.UpdateIkePolicyOptions{
				ID: &ikePolicyID,
			}
			ikePolicyPatchModel := &vpcv1.IkePolicyPatch{}
			ikePolicyPatchModel.Name = &name
			ikePolicyPatchModel.DhGroup = &[]int64{5}[0]
			ikePolicyPatch, asPatchErr := ikePolicyPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.IkePolicyPatch = ikePolicyPatch
			ikePolicy, response, err := vpcService.UpdateIkePolicy(options)

			// end-update_ike_policy
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ikePolicy).ToNot(BeNil())

		})
		It(`ListIkePolicyConnections request example`, func() {
			fmt.Println("\nListIkePolicyConnections() result:")
			// begin-list_ike_policy_connections

			options := &vpcv1.ListIkePolicyConnectionsOptions{
				ID: &ikePolicyID,
			}
			connections, response, err := vpcService.ListIkePolicyConnections(options)

			// end-list_ike_policy_connections
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(connections).ToNot(BeNil())

		})
		It(`ListIpsecPolicies request example`, func() {
			fmt.Println("\nListIpsecPolicies() result:")
			// begin-list_ipsec_policies

			options := &vpcv1.ListIpsecPoliciesOptions{}
			ipsecPolicies, response, err := vpcService.ListIpsecPolicies(options)

			// end-list_ipsec_policies
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ipsecPolicies).ToNot(BeNil())

		})
		It(`CreateIpsecPolicy request example`, func() {
			fmt.Println("\nCreateIpsecPolicy() result:")
			name := getName("ipsec")
			// begin-create_ipsec_policy

			options := &vpcv1.CreateIpsecPolicyOptions{}
			options.SetName(name)
			options.SetAuthenticationAlgorithm("md5")
			options.SetEncryptionAlgorithm("aes128")
			options.SetPfs("disabled")
			ipsecPolicy, response, err := vpcService.CreateIpsecPolicy(options)
			// end-create_ipsec_policy
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(ipsecPolicy).ToNot(BeNil())
			ipsecPolicyID = *ipsecPolicy.ID
		})
		It(`GetIpsecPolicy request example`, func() {
			fmt.Println("\nGetIpsecPolicy() result:")
			// begin-get_ipsec_policy

			options := vpcService.NewGetIpsecPolicyOptions(ipsecPolicyID)
			ipsecPolicy, response, err := vpcService.GetIpsecPolicy(options)

			// end-get_ipsec_policy
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ipsecPolicy).ToNot(BeNil())

		})
		It(`UpdateIpsecPolicy request example`, func() {
			fmt.Println("\nUpdateIpsecPolicy() result:")
			name := getName("ipsec")
			// begin-update_ipsec_policy

			options := &vpcv1.UpdateIpsecPolicyOptions{
				ID: &ipsecPolicyID,
			}
			ipsecPolicyPatchModel := &vpcv1.IPsecPolicyPatch{
				Name:                    &name,
				AuthenticationAlgorithm: &[]string{"sha256"}[0],
			}
			ipsecPolicyPatch, asPatchErr := ipsecPolicyPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.IPsecPolicyPatch = ipsecPolicyPatch
			ipsecPolicy, response, err := vpcService.UpdateIpsecPolicy(options)
			// end-update_ipsec_policy
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ipsecPolicy).ToNot(BeNil())
		})
		It(`ListIpsecPolicyConnections request example`, func() {
			fmt.Println("\nListIpsecPolicyConnections() result:")
			// begin-list_ipsec_policy_connections

			options := &vpcv1.ListIpsecPolicyConnectionsOptions{
				ID: &ipsecPolicyID,
			}
			connections, response, err :=
				vpcService.ListIpsecPolicyConnections(options)

			// end-list_ipsec_policy_connections
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(connections).ToNot(BeNil())

		})
		It(`ListVPNGateways request example`, func() {
			fmt.Println("\nListVPNGateways() result:")
			// begin-list_vpn_gateways

			listVPNGatewaysOptions := vpcService.NewListVPNGatewaysOptions()
			vpnGatewayCollection, response, err := vpcService.ListVPNGateways(listVPNGatewaysOptions)

			// end-list_vpn_gateways
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnGatewayCollection).ToNot(BeNil())

		})
		It(`CreateVPNGateway request example`, func() {
			fmt.Println("\nCreateVPNGateway() result:")
			name := getName("vpngateway")
			// begin-create_vpn_gateway

			vpnGatewayPrototypeModel := new(vpcv1.VPNGatewayPrototypeVPNGatewayRouteModePrototype)
			vpnGatewayPrototypeModel.Name = &name
			vpnGatewayPrototypeModel.Subnet = &vpcv1.SubnetIdentityByID{
				ID: &subnetID,
			}
			vpnGatewayPrototypeModel.Mode = &[]string{"route"}[0]

			createVPNGatewayOptionsModel := new(vpcv1.CreateVPNGatewayOptions)
			createVPNGatewayOptionsModel.VPNGatewayPrototype = vpnGatewayPrototypeModel
			vpnGatewayIntf, response, err := vpcService.CreateVPNGateway(createVPNGatewayOptionsModel)
			vpnGateway := vpnGatewayIntf.(*vpcv1.VPNGateway)
			// end-create_vpn_gateway
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(vpnGateway).ToNot(BeNil())
			vpnGatewayID = *vpnGateway.ID
		})
		It(`GetVPNGateway request example`, func() {
			fmt.Println("\nGetVPNGateway() result:")
			// begin-get_vpn_gateway

			options := vpcService.NewGetVPNGatewayOptions(vpnGatewayID)
			vpnGateway, response, err := vpcService.GetVPNGateway(options)

			// end-get_vpn_gateway
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnGateway).ToNot(BeNil())

		})
		It(`UpdateVPNGateway request example`, func() {
			fmt.Println("\nUpdateVPNGateway() result:")
			name := getName("vpngateway")
			// begin-update_vpn_gateway

			options := &vpcv1.UpdateVPNGatewayOptions{
				ID: &vpnGatewayID,
			}
			vpnGatewayPatchModel := &vpcv1.VPNGatewayPatch{
				Name: &name,
			}
			vpnGatewayPatch, asPatchErr := vpnGatewayPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.VPNGatewayPatch = vpnGatewayPatch
			vpnGateway, response, err := vpcService.UpdateVPNGateway(options)
			// end-update_vpn_gateway
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnGateway).ToNot(BeNil())

		})
		It(`ListVPNGatewayConnections request example`, func() {
			fmt.Println("\nListVPNGatewayConnections() result:")
			// begin-list_vpn_gateway_connections

			options := &vpcv1.ListVPNGatewayConnectionsOptions{}
			options.SetVPNGatewayID(vpnGatewayID)
			vpnGatewayConnections, response, err := vpcService.ListVPNGatewayConnections(
				options,
			)

			// end-list_vpn_gateway_connections
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnGatewayConnections).ToNot(BeNil())

		})
		It(`CreateVPNGatewayConnection request example`, func() {
			fmt.Println("\nCreateVPNGatewayConnection() result:")
			name := getName("vpnconnection")
			// begin-create_vpn_gateway_connection

			options := &vpcv1.CreateVPNGatewayConnectionOptions{
				VPNGatewayConnectionPrototype: &vpcv1.VPNGatewayConnectionPrototypeVPNGatewayConnectionPolicyModePrototype{
					PeerAddress: &[]string{"192.132.5.0"}[0],
					Psk:         &[]string{"lkj14b1oi0alcniejkso"}[0],
					Name:        &name,
					PeerCIDRs:   []string{"197.155.0.0/28"},
					LocalCIDRs:  []string{"192.132.0.0/28"},
				},
				VPNGatewayID: &vpnGatewayID,
			}
			vpnGatewayConnectionIntf, response, err := vpcService.CreateVPNGatewayConnection(
				options,
			)
			vpnGatewayConnection := vpnGatewayConnectionIntf.(*vpcv1.VPNGatewayConnection)
			// end-create_vpn_gateway_connection
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(vpnGatewayConnection).ToNot(BeNil())
			vpnGatewayConnectionID = *vpnGatewayConnection.ID
		})
		It(`GetVPNGatewayConnection request example`, func() {
			fmt.Println("\nGetVPNGatewayConnection() result:")
			// begin-get_vpn_gateway_connection

			options := &vpcv1.GetVPNGatewayConnectionOptions{}
			options.SetVPNGatewayID(vpnGatewayID)
			options.SetID(vpnGatewayConnectionID)
			vpnGatewayConnection, response, err := vpcService.GetVPNGatewayConnection(options)

			// end-get_vpn_gateway_connection
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnGatewayConnection).ToNot(BeNil())

		})
		It(`UpdateVPNGatewayConnection request example`, func() {
			fmt.Println("\nUpdateVPNGatewayConnection() result:")
			name := getName("vpnConnection")
			// begin-update_vpn_gateway_connection
			options := &vpcv1.UpdateVPNGatewayConnectionOptions{
				ID:           &vpnGatewayConnectionID,
				VPNGatewayID: &vpnGatewayID,
			}
			vpnGatewayConnectionPatchModel := &vpcv1.VPNGatewayConnectionPatch{}
			vpnGatewayConnectionPatchModel.Name = &name
			vpnGatewayConnectionPatchModel.PeerAddress = &[]string{"192.132.5.0"}[0]
			vpnGatewayConnectionPatchModel.Psk = &[]string{"lkj14b1oi0alcniejkso"}[0]
			vpnGatewayConnectionPatch, asPatchErr := vpnGatewayConnectionPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.VPNGatewayConnectionPatch = vpnGatewayConnectionPatch
			vpnGatewayConnection, response, err := vpcService.UpdateVPNGatewayConnection(
				options,
			)

			// end-update_vpn_gateway_connection
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnGatewayConnection).ToNot(BeNil())

		})
		It(`AddVPNGatewayConnectionLocalCIDR request example`, func() {
			// begin-add_vpn_gateway_connection_local_cidr

			options := &vpcv1.AddVPNGatewayConnectionLocalCIDROptions{}
			options.SetVPNGatewayID(vpnGatewayID)
			options.SetID(vpnGatewayConnectionID)
			options.SetCIDRPrefix("192.134.0.0")
			options.SetPrefixLength("28")
			response, err := vpcService.AddVPNGatewayConnectionLocalCIDR(options)

			// end-add_vpn_gateway_connection_local_cidr
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nAddVPNGatewayConnectionLocalCIDR() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`ListVPNGatewayConnectionLocalCidrs request example`, func() {
			fmt.Println("\nListVPNGatewayConnectionLocalCidrs() result:")
			// begin-list_vpn_gateway_connection_local_cidrs

			options := &vpcv1.ListVPNGatewayConnectionLocalCIDRsOptions{}
			options.SetVPNGatewayID(vpnGatewayID)
			options.SetID(vpnGatewayConnectionID)
			localCIDRs, response, err :=
				vpcService.ListVPNGatewayConnectionLocalCIDRs(options)

			// end-list_vpn_gateway_connection_local_cidrs
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(localCIDRs).ToNot(BeNil())

		})
		It(`AddVPNGatewayConnectionPeerCIDR request example`, func() {
			// begin-add_vpn_gateway_connection_peer_cidr

			options := &vpcv1.AddVPNGatewayConnectionPeerCIDROptions{}
			options.SetVPNGatewayID(vpnGatewayID)
			options.SetID(vpnGatewayConnectionID)
			options.SetCIDRPrefix("192.144.0.0")
			options.SetPrefixLength("28")
			response, err := vpcService.AddVPNGatewayConnectionPeerCIDR(options)

			// end-add_vpn_gateway_connection_peer_cidr
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nAddVPNGatewayConnectionPeerCIDR() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`CheckVPNGatewayConnectionLocalCIDR request example`, func() {
			// begin-check_vpn_gateway_connection_local_cidr

			options := &vpcv1.CheckVPNGatewayConnectionLocalCIDROptions{}
			options.SetVPNGatewayID(vpnGatewayID)
			options.SetID(vpnGatewayConnectionID)
			options.SetCIDRPrefix("192.134.0.0")
			options.SetPrefixLength("28")
			response, err := vpcService.CheckVPNGatewayConnectionLocalCIDR(options)

			// end-check_vpn_gateway_connection_local_cidr
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nCheckVPNGatewayConnectionLocalCIDR() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})

		It(`ListVPNGatewayConnectionPeerCidrs request example`, func() {
			fmt.Println("\nListVPNGatewayConnectionPeerCidrs() result:")
			// begin-list_vpn_gateway_connection_peer_cidrs

			options := &vpcv1.ListVPNGatewayConnectionPeerCIDRsOptions{}
			options.SetVPNGatewayID(vpnGatewayID)
			options.SetID(vpnGatewayConnectionID)
			peerCIDRs, response, err :=
				vpcService.ListVPNGatewayConnectionPeerCIDRs(options)

			// end-list_vpn_gateway_connection_peer_cidrs
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(peerCIDRs).ToNot(BeNil())

		})
		It(`CheckVPNGatewayConnectionPeerCIDR request example`, func() {
			// begin-check_vpn_gateway_connection_peer_cidr

			options := &vpcv1.CheckVPNGatewayConnectionPeerCIDROptions{}
			options.SetVPNGatewayID(vpnGatewayID)
			options.SetID(vpnGatewayConnectionID)
			options.SetCIDRPrefix("192.144.0.0")
			options.SetPrefixLength("28")
			response, err := vpcService.CheckVPNGatewayConnectionPeerCIDR(options)
			// end-check_vpn_gateway_connection_peer_cidr
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nCheckVPNGatewayConnectionPeerCIDR() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})

		It(`ListBareMetalServerProfiles request example`, func() {
			fmt.Println("\nListBareMetalServerProfiles() result:")
			// begin-list_bare_metal_server_profiles

			listBareMetalServerProfilesOptions := vpcService.NewListBareMetalServerProfilesOptions()

			bareMetalServerProfileCollection, response, err := vpcService.ListBareMetalServerProfiles(listBareMetalServerProfilesOptions)
			if err != nil {
				panic(err)
			}

			// end-list_bare_metal_server_profiles

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServerProfileCollection).ToNot(BeNil())
			bareMetalServerProfileName = *bareMetalServerProfileCollection.Profiles[0].Name

		})
		It(`GetBareMetalServerProfile request example`, func() {
			fmt.Println("\nGetBareMetalServerProfile() result:")
			// begin-get_bare_metal_server_profile

			getBareMetalServerProfileOptions := &vpcv1.GetBareMetalServerProfileOptions{
				Name: &bareMetalServerProfileName,
			}

			bareMetalServerProfile, response, err := vpcService.GetBareMetalServerProfile(getBareMetalServerProfileOptions)
			if err != nil {
				panic(err)
			}

			// end-get_bare_metal_server_profile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServerProfile).ToNot(BeNil())

		})
		It(`ListBareMetalServers request example`, func() {
			fmt.Println("\nListBareMetalServers() result:")
			// begin-list_bare_metal_servers

			listBareMetalServersOptions := &vpcv1.ListBareMetalServersOptions{}

			bareMetalServerCollection, response, err := vpcService.ListBareMetalServers(listBareMetalServersOptions)
			if err != nil {
				panic(err)
			}

			// end-list_bare_metal_servers

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServerCollection).ToNot(BeNil())

		})
		It(`CreateBareMetalServer request example`, func() {
			fmt.Println("\nCreateBareMetalServer() result:")
			// begin-create_bare_metal_server

			imageIdentityModel := &vpcv1.ImageIdentityByID{
				ID: &imageID,
			}

			keyIdentityModel := &vpcv1.KeyIdentityByID{
				ID: &keyID,
			}

			bareMetalServerInitializationPrototypeModel := &vpcv1.BareMetalServerInitializationPrototype{
				Image: imageIdentityModel,
				Keys:  []vpcv1.KeyIdentityIntf{keyIdentityModel},
			}

			subnetIdentityModel := &vpcv1.SubnetIdentityByID{
				ID: &subnetID,
			}

			bareMetalServerPrimaryNetworkInterfacePrototypeModel := &vpcv1.BareMetalServerPrimaryNetworkInterfacePrototype{
				Subnet: subnetIdentityModel,
			}

			bareMetalServerProfileIdentityModel := &vpcv1.BareMetalServerProfileIdentityByName{
				Name: &bareMetalServerProfileName,
			}

			zoneIdentityModel := &vpcv1.ZoneIdentityByName{
				Name: zone,
			}

			createBareMetalServerOptions := &vpcv1.CreateBareMetalServerOptions{
				Initialization:          bareMetalServerInitializationPrototypeModel,
				PrimaryNetworkInterface: bareMetalServerPrimaryNetworkInterfacePrototypeModel,
				Profile:                 bareMetalServerProfileIdentityModel,
				Zone:                    zoneIdentityModel,
			}
			createBareMetalServerOptions.SetName("my-bare-metal-server")

			bareMetalServer, response, err := vpcService.CreateBareMetalServer(createBareMetalServerOptions)
			if err != nil {
				panic(err)
			}

			// end-create_bare_metal_server

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(bareMetalServer).ToNot(BeNil())
			bareMetalServerId = *bareMetalServer.ID
		})
		It(`CreateBareMetalServerConsoleAccessToken request example`, func() {
			Skip("not runnin with mock")
			fmt.Println("\nCreateBareMetalServerConsoleAccessToken() result:")
			// begin-create_bare_metal_server_console_access_token

			createBareMetalServerConsoleAccessTokenOptions := &vpcv1.CreateBareMetalServerConsoleAccessTokenOptions{
				BareMetalServerID: &bareMetalServerId,
			}
			createBareMetalServerConsoleAccessTokenOptions.SetConsoleType("serial")

			bareMetalServerConsoleAccessToken, response, err := vpcService.CreateBareMetalServerConsoleAccessToken(createBareMetalServerConsoleAccessTokenOptions)
			if err != nil {
				panic(err)
			}

			// end-create_bare_metal_server_console_access_token

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServerConsoleAccessToken).ToNot(BeNil())

		})
		It(`ListBareMetalServerDisks request example`, func() {
			fmt.Println("\nListBareMetalServerDisks() result:")
			// begin-list_bare_metal_server_disks

			listBareMetalServerDisksOptions := &vpcv1.ListBareMetalServerDisksOptions{
				BareMetalServerID: &bareMetalServerId,
			}

			bareMetalServerDiskCollection, response, err := vpcService.ListBareMetalServerDisks(listBareMetalServerDisksOptions)
			if err != nil {
				panic(err)
			}

			// end-list_bare_metal_server_disks

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServerDiskCollection).ToNot(BeNil())
			bareMetalServerDiskId = *bareMetalServerDiskCollection.Disks[0].ID
		})
		It(`GetBareMetalServerDisk request example`, func() {
			fmt.Println("\nGetBareMetalServerDisk() result:")
			// begin-get_bare_metal_server_disk

			getBareMetalServerDiskOptions := &vpcv1.GetBareMetalServerDiskOptions{
				BareMetalServerID: &bareMetalServerId,
				ID:                &bareMetalServerDiskId,
			}

			bareMetalServerDisk, response, err := vpcService.GetBareMetalServerDisk(getBareMetalServerDiskOptions)
			if err != nil {
				panic(err)
			}

			// end-get_bare_metal_server_disk

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServerDisk).ToNot(BeNil())

		})
		It(`UpdateBareMetalServerDisk request example`, func() {
			fmt.Println("\nUpdateBareMetalServerDisk() result:")
			// begin-update_bare_metal_server_disk

			bareMetalServerDiskPatchModel := &vpcv1.BareMetalServerDiskPatch{}
			bareMetalServerDiskPatchModel.Name = &[]string{"my-bare-metal-server-disk-update"}[0]

			bareMetalServerDiskPatchModelAsPatch, asPatchErr := bareMetalServerDiskPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateBareMetalServerDiskOptions := &vpcv1.UpdateBareMetalServerDiskOptions{
				BareMetalServerID:        &bareMetalServerId,
				ID:                       &bareMetalServerDiskId,
				BareMetalServerDiskPatch: bareMetalServerDiskPatchModelAsPatch,
			}

			bareMetalServerDisk, response, err := vpcService.UpdateBareMetalServerDisk(updateBareMetalServerDiskOptions)
			if err != nil {
				panic(err)
			}

			// end-update_bare_metal_server_disk

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServerDisk).ToNot(BeNil())

		})
		It(`ListBareMetalServerNetworkInterfaces request example`, func() {
			fmt.Println("\nListBareMetalServerNetworkInterfaces() result:")
			// begin-list_bare_metal_server_network_interfaces

			listBareMetalServerNetworkInterfacesOptions := &vpcv1.ListBareMetalServerNetworkInterfacesOptions{
				BareMetalServerID: &bareMetalServerId,
			}

			bareMetalServerNetworkInterfaceCollection, response, err := vpcService.ListBareMetalServerNetworkInterfaces(listBareMetalServerNetworkInterfacesOptions)
			if err != nil {
				panic(err)
			}

			// end-list_bare_metal_server_network_interfaces

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServerNetworkInterfaceCollection).ToNot(BeNil())

		})
		It(`CreateBareMetalServerNetworkInterface request example`, func() {
			fmt.Println("\nCreateBareMetalServerNetworkInterface() result:")
			// begin-create_bare_metal_server_network_interface

			subnetIdentityModel := &vpcv1.SubnetIdentityByID{
				ID: &subnetID,
			}

			bareMetalServerNetworkInterfacePrototypeModel := &vpcv1.BareMetalServerNetworkInterfacePrototypeBareMetalServerNetworkInterfaceByPciPrototype{
				InterfaceType: core.StringPtr("pci"),
				Subnet:        subnetIdentityModel,
				Name:          core.StringPtr("my-metal-server-nic"),
			}

			createBareMetalServerNetworkInterfaceOptions := &vpcv1.CreateBareMetalServerNetworkInterfaceOptions{
				BareMetalServerID:                        &bareMetalServerId,
				BareMetalServerNetworkInterfacePrototype: bareMetalServerNetworkInterfacePrototypeModel,
			}

			bareMetalServerNetworkInterface, response, err := vpcService.CreateBareMetalServerNetworkInterface(createBareMetalServerNetworkInterfaceOptions)
			if err != nil {
				panic(err)
			}

			// end-create_bare_metal_server_network_interface

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(bareMetalServerNetworkInterface).ToNot(BeNil())
			switch reflect.TypeOf(bareMetalServerNetworkInterface).String() {
			case "*vpcv1.BareMetalServerNetworkInterfaceByPci":
				{
					nic := bareMetalServerNetworkInterface.(*vpcv1.BareMetalServerNetworkInterfaceByPci)
					bareMetalServerNetworkInterfaceId = *nic.ID
				}
			}
		})
		It(`GetBareMetalServerNetworkInterface request example`, func() {
			fmt.Println("\nGetBareMetalServerNetworkInterface() result:")
			// begin-get_bare_metal_server_network_interface

			getBareMetalServerNetworkInterfaceOptions := &vpcv1.GetBareMetalServerNetworkInterfaceOptions{
				BareMetalServerID: &bareMetalServerId,
				ID:                &bareMetalServerNetworkInterfaceId,
			}

			bareMetalServerNetworkInterface, response, err := vpcService.GetBareMetalServerNetworkInterface(getBareMetalServerNetworkInterfaceOptions)
			if err != nil {
				panic(err)
			}

			// end-get_bare_metal_server_network_interface

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServerNetworkInterface).ToNot(BeNil())

		})
		It(`UpdateBareMetalServerNetworkInterface request example`, func() {
			fmt.Println("\nUpdateBareMetalServerNetworkInterface() result:")
			// begin-update_bare_metal_server_network_interface

			bareMetalServerNetworkInterfacePatchModel := &vpcv1.BareMetalServerNetworkInterfacePatch{
				Name: core.StringPtr("my-metal-server-nic-update"),
			}
			bareMetalServerNetworkInterfacePatchModelAsPatch, asPatchErr := bareMetalServerNetworkInterfacePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateBareMetalServerNetworkInterfaceOptions := &vpcv1.UpdateBareMetalServerNetworkInterfaceOptions{
				BareMetalServerID:                    &bareMetalServerId,
				ID:                                   &bareMetalServerNetworkInterfaceId,
				BareMetalServerNetworkInterfacePatch: bareMetalServerNetworkInterfacePatchModelAsPatch,
			}

			bareMetalServerNetworkInterface, response, err := vpcService.UpdateBareMetalServerNetworkInterface(updateBareMetalServerNetworkInterfaceOptions)
			if err != nil {
				panic(err)
			}

			// end-update_bare_metal_server_network_interface

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServerNetworkInterface).ToNot(BeNil())

		})
		It(`ListBareMetalServerNetworkInterfaceFloatingIps request example`, func() {
			fmt.Println("\nListBareMetalServerNetworkInterfaceFloatingIps() result:")
			// begin-list_bare_metal_server_network_interface_floating_ips

			listBareMetalServerNetworkInterfaceFloatingIpsOptions := &vpcv1.ListBareMetalServerNetworkInterfaceFloatingIpsOptions{
				BareMetalServerID:  &bareMetalServerId,
				NetworkInterfaceID: &bareMetalServerNetworkInterfaceId,
			}

			floatingIPUnpaginatedCollection, response, err := vpcService.ListBareMetalServerNetworkInterfaceFloatingIps(listBareMetalServerNetworkInterfaceFloatingIpsOptions)
			if err != nil {
				panic(err)
			}

			// end-list_bare_metal_server_network_interface_floating_ips

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(floatingIPUnpaginatedCollection).ToNot(BeNil())

		})
		It(`AddBareMetalServerNetworkInterfaceFloatingIP request example`, func() {
			fmt.Println("\nAddBareMetalServerNetworkInterfaceFloatingIP() result:")
			// begin-add_bare_metal_server_network_interface_floating_ip

			addBareMetalServerNetworkInterfaceFloatingIPOptions := &vpcv1.AddBareMetalServerNetworkInterfaceFloatingIPOptions{
				BareMetalServerID:  &bareMetalServerId,
				NetworkInterfaceID: &bareMetalServerNetworkInterfaceId,
				ID:                 &floatingIPID,
			}

			floatingIP, response, err := vpcService.AddBareMetalServerNetworkInterfaceFloatingIP(addBareMetalServerNetworkInterfaceFloatingIPOptions)
			if err != nil {
				panic(err)
			}

			// end-add_bare_metal_server_network_interface_floating_ip

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(floatingIP).ToNot(BeNil())

		})
		It(`GetBareMetalServerNetworkInterfaceFloatingIP request example`, func() {
			fmt.Println("\nGetBareMetalServerNetworkInterfaceFloatingIP() result:")
			// begin-get_bare_metal_server_network_interface_floating_ip

			getBareMetalServerNetworkInterfaceFloatingIPOptions := &vpcv1.GetBareMetalServerNetworkInterfaceFloatingIPOptions{
				BareMetalServerID:  &bareMetalServerId,
				NetworkInterfaceID: &bareMetalServerNetworkInterfaceId,
				ID:                 &floatingIPID,
			}

			floatingIP, response, err := vpcService.GetBareMetalServerNetworkInterfaceFloatingIP(getBareMetalServerNetworkInterfaceFloatingIPOptions)
			if err != nil {
				panic(err)
			}

			// end-get_bare_metal_server_network_interface_floating_ip

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(floatingIP).ToNot(BeNil())

		})
		It(`GetBareMetalServer request example`, func() {
			fmt.Println("\nGetBareMetalServer() result:")
			// begin-get_bare_metal_server

			getBareMetalServerOptions := &vpcv1.GetBareMetalServerOptions{
				ID: &bareMetalServerId,
			}

			bareMetalServer, response, err := vpcService.GetBareMetalServer(getBareMetalServerOptions)
			if err != nil {
				panic(err)
			}

			// end-get_bare_metal_server

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServer).ToNot(BeNil())

		})
		It(`UpdateBareMetalServer request example`, func() {
			fmt.Println("\nUpdateBareMetalServer() result:")
			// begin-update_bare_metal_server

			bareMetalServerPatchModel := &vpcv1.BareMetalServerPatch{
				Name: core.StringPtr("my-metal-server-update"),
			}
			bareMetalServerPatchModelAsPatch, asPatchErr := bareMetalServerPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateBareMetalServerOptions := &vpcv1.UpdateBareMetalServerOptions{
				ID:                   &bareMetalServerId,
				BareMetalServerPatch: bareMetalServerPatchModelAsPatch,
			}

			bareMetalServer, response, err := vpcService.UpdateBareMetalServer(updateBareMetalServerOptions)
			if err != nil {
				panic(err)
			}

			// end-update_bare_metal_server

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServer).ToNot(BeNil())

		})
		It(`GetBareMetalServerInitialization request example`, func() {
			fmt.Println("\nGetBareMetalServerInitialization() result:")
			// begin-get_bare_metal_server_initialization

			getBareMetalServerInitializationOptions := &vpcv1.GetBareMetalServerInitializationOptions{
				ID: &bareMetalServerId,
			}

			bareMetalServerInitialization, response, err := vpcService.GetBareMetalServerInitialization(getBareMetalServerInitializationOptions)
			if err != nil {
				panic(err)
			}

			// end-get_bare_metal_server_initialization

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServerInitialization).ToNot(BeNil())

		})
		It(`RestartBareMetalServer request example`, func() {
			// begin-restart_bare_metal_server

			restartBareMetalServerOptions := &vpcv1.RestartBareMetalServerOptions{
				ID: &bareMetalServerId,
			}

			response, err := vpcService.RestartBareMetalServer(restartBareMetalServerOptions)
			if err != nil {
				panic(err)
			}

			// end-restart_bare_metal_server
			fmt.Printf("\nRestartBareMetalServer() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`StartBareMetalServer request example`, func() {
			// begin-start_bare_metal_server

			startBareMetalServerOptions := &vpcv1.StartBareMetalServerOptions{
				ID: &bareMetalServerId,
			}

			response, err := vpcService.StartBareMetalServer(startBareMetalServerOptions)
			if err != nil {
				panic(err)
			}

			// end-start_bare_metal_server
			fmt.Printf("\nStartBareMetalServer() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`StopBareMetalServer request example`, func() {
			// begin-stop_bare_metal_server

			stopBareMetalServerOptions := &vpcv1.StopBareMetalServerOptions{
				ID:   &bareMetalServerId,
				Type: core.StringPtr("soft"),
			}

			response, err := vpcService.StopBareMetalServer(stopBareMetalServerOptions)
			if err != nil {
				panic(err)
			}

			// end-stop_bare_metal_server
			fmt.Printf("\nStopBareMetalServer() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})

		It(`ListBackupPolicies request example`, func() {
			fmt.Println("\nListBackupPolicies() result:")
			// begin-list_backup_policies

			listBackupPoliciesOptions := vpcService.NewListBackupPoliciesOptions()

			backupPolicyCollection, response, err := vpcService.ListBackupPolicies(listBackupPoliciesOptions)
			if err != nil {
				panic(err)
			}

			// end-list_backup_policies

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(backupPolicyCollection).ToNot(BeNil())

		})
		It(`CreateBackupPolicy request example`, func() {
			fmt.Println("\nCreateBackupPolicy() result:")
			// begin-create_backup_policy

			createBackupPolicyOptions := vpcService.NewCreateBackupPolicyOptions()
			userTags := []string{"tag1", "tag2"}
			createBackupPolicyOptions.SetName("my-backup-policy")
			createBackupPolicyOptions.SetMatchUserTags(userTags)

			backupPolicy, response, err := vpcService.CreateBackupPolicy(createBackupPolicyOptions)
			if err != nil {
				panic(err)
			}

			// end-create_backup_policy
			backupPolicyID = *backupPolicy.ID

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(backupPolicy).ToNot(BeNil())

		})
		It(`CreateBackupPolicyPlan request example`, func() {
			fmt.Println("\nCreateBackupPolicyPlan() result:")
			// begin-create_backup_policy_plan

			createBackupPolicyPlanOptions := vpcService.NewCreateBackupPolicyPlanOptions(
				backupPolicyID,
				"*/5 1,2,3 * * *",
			)
			createBackupPolicyPlanOptions.SetName("my-backup-policy-plan")

			backupPolicyPlan, response, err := vpcService.CreateBackupPolicyPlan(createBackupPolicyPlanOptions)
			if err != nil {
				panic(err)
			}

			// end-create_backup_policy_plan
			backupPolicyPlanID = *backupPolicyPlan.ID
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(backupPolicyPlan).ToNot(BeNil())

		})
		It(`ListBackupPolicyPlans request example`, func() {
			fmt.Println("\nListBackupPolicyPlans() result:")
			// begin-list_backup_policy_plans

			listBackupPolicyPlansOptions := vpcService.NewListBackupPolicyPlansOptions(
				backupPolicyID,
			)

			backupPolicyPlanCollection, response, err := vpcService.ListBackupPolicyPlans(listBackupPolicyPlansOptions)
			if err != nil {
				panic(err)
			}

			// end-list_backup_policy_plans

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(backupPolicyPlanCollection).ToNot(BeNil())

		})

		It(`GetBackupPolicyPlan request example`, func() {
			fmt.Println("\nGetBackupPolicyPlan() result:")
			// begin-get_backup_policy_plan

			getBackupPolicyPlanOptions := vpcService.NewGetBackupPolicyPlanOptions(
				backupPolicyID,
				backupPolicyPlanID,
			)

			backupPolicyPlan, response, err := vpcService.GetBackupPolicyPlan(getBackupPolicyPlanOptions)
			if err != nil {
				panic(err)
			}

			// end-get_backup_policy_plan
			ifMatchBackupPolicy = response.GetHeaders()["Etag"][0]
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(backupPolicyPlan).ToNot(BeNil())

		})
		It(`UpdateBackupPolicyPlan request example`, func() {
			fmt.Println("\nUpdateBackupPolicyPlan() result:")
			// begin-update_backup_policy_plan

			backupPolicyPlanPatchModel := &vpcv1.BackupPolicyPlanPatch{
				Name: core.StringPtr("my-backup-plan-updated"),
			}
			backupPolicyPlanPatchModelAsPatch, asPatchErr := backupPolicyPlanPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateBackupPolicyPlanOptions := vpcService.NewUpdateBackupPolicyPlanOptions(
				backupPolicyID,
				backupPolicyPlanID,
				backupPolicyPlanPatchModelAsPatch,
			)
			updateBackupPolicyPlanOptions.SetIfMatch(ifMatchBackupPolicy)

			backupPolicyPlan, response, err := vpcService.UpdateBackupPolicyPlan(updateBackupPolicyPlanOptions)
			if err != nil {
				panic(err)
			}

			// end-update_backup_policy_plan

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(backupPolicyPlan).ToNot(BeNil())

		})
		It(`GetBackupPolicy request example`, func() {
			fmt.Println("\nGetBackupPolicy() result:")
			// begin-get_backup_policy

			getBackupPolicyOptions := vpcService.NewGetBackupPolicyOptions(
				backupPolicyID,
			)

			backupPolicy, response, err := vpcService.GetBackupPolicy(getBackupPolicyOptions)
			if err != nil {
				panic(err)
			}

			// end-get_backup_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(backupPolicy).ToNot(BeNil())

		})
		It(`UpdateBackupPolicy request example`, func() {
			fmt.Println("\nUpdateBackupPolicy() result:")
			// begin-update_backup_policy

			backupPolicyPatchModel := &vpcv1.BackupPolicyPatch{
				Name: core.StringPtr("my-backup-policy-update"),
			}
			backupPolicyPatchModelAsPatch, asPatchErr := backupPolicyPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateBackupPolicyOptions := vpcService.NewUpdateBackupPolicyOptions(
				backupPolicyID,
				backupPolicyPatchModelAsPatch,
			)
			updateBackupPolicyOptions.SetIfMatch(ifMatchBackupPolicy)

			backupPolicy, response, err := vpcService.UpdateBackupPolicy(updateBackupPolicyOptions)
			if err != nil {
				panic(err)
			}

			// end-update_backup_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(backupPolicy).ToNot(BeNil())

		})

		It(`ListPlacementGroups request example`, func() {
			fmt.Println("\nListPlacementGroups() result:")
			// begin-list_placement_groups

			options := &vpcv1.ListPlacementGroupsOptions{}
			placementGroups, response, err := vpcService.ListPlacementGroups(options)

			// end-list_flow_log_collectors
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(placementGroups).ToNot(BeNil())

		})
		It(`CreatePlacementGroup request example`, func() {
			fmt.Println("\nCreatePlacementGroup() result:")
			name := getName("placement")
			// begin-create_flow_log_collector

			strategy := "host_spread"
			createPlacementGroupOptions := &vpcv1.CreatePlacementGroupOptions{
				Strategy: &strategy,
				Name:     &name,
			}
			placementGroup, response, err := vpcService.CreatePlacementGroup(createPlacementGroupOptions)

			// end-create_flow_log_collector
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(placementGroup).ToNot(BeNil())
			placementGroupID = *placementGroup.ID
		})
		It(`GetPlacementGroup request example`, func() {
			fmt.Println("\nGetPlacementGroup() result:")
			// begin-get_flow_log_collector

			getPlacementGroupOptions := &vpcv1.GetPlacementGroupOptions{
				ID: &placementGroupID,
			}

			placementGroup, response, err := vpcService.GetPlacementGroup(getPlacementGroupOptions)

			// end-get_flow_log_collector
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(placementGroup).ToNot(BeNil())

		})

		It(`UpdatePlacementGroup request example`, func() {
			fmt.Println("\nUpdatePlacementGroup() result:")
			name := getName("fl")
			// begin-update_flow_log_collector

			placementGroupPatchModel := &vpcv1.PlacementGroupPatch{
				Name: &name,
			}
			placementGroupPatchModelAsPatch, asPatchErr := placementGroupPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}

			options := &vpcv1.UpdatePlacementGroupOptions{
				ID:                  &placementGroupID,
				PlacementGroupPatch: placementGroupPatchModelAsPatch,
			}

			placementGroup, response, err := vpcService.UpdatePlacementGroup(options)

			// end-update_flow_log_collector
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(placementGroup).ToNot(BeNil())

		})

		It(`DeletePlacementGroup request example`, func() {
			// begin-delete_flow_log_collector

			deletePlacementGroupOptions := &vpcv1.DeletePlacementGroupOptions{
				ID: &placementGroupID,
			}

			response, err := vpcService.DeletePlacementGroup(deletePlacementGroupOptions)

			// end-delete_flow_log_collector
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeletePlacementGroup() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`ListVPNServers request example`, func() {
			fmt.Println("\nListVPNServers() result:")
			// begin-list_vpn_servers

			listVPNServersOptions := vpcService.NewListVPNServersOptions()
			listVPNServersOptions.SetSort("name")

			vpnServerCollection, response, err := vpcService.ListVPNServers(listVPNServersOptions)
			if err != nil {
				panic(err)
			}

			// end-list_vpn_servers

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnServerCollection).ToNot(BeNil())

		})
		It(`CreateVPNServer request example`, func() {
			fmt.Println("\nCreateVPNServer() result:")
			// begin-create_vpn_server

			certificateInstanceIdentityModel := &vpcv1.CertificateInstanceIdentityByCRN{
				CRN: core.StringPtr("crn:v1:bluemix:public:secrets-manager:us-south:a/123456:36fa422d-080d-4d83-8d2d-86851b4001df:secret:2e786aab-42fa-63ed-14f8-d66d552f4dd5"),
			}

			vpnServerAuthenticationByUsernameIDProviderModel := &vpcv1.VPNServerAuthenticationByUsernameIDProviderByIam{
				ProviderType: core.StringPtr("iam"),
			}

			vpnServerAuthenticationPrototypeModel := &vpcv1.VPNServerAuthenticationPrototypeVPNServerAuthenticationByUsernamePrototype{
				Method:           core.StringPtr("certificate"),
				IdentityProvider: vpnServerAuthenticationByUsernameIDProviderModel,
			}

			subnetIdentityModel := &vpcv1.SubnetIdentityByID{
				ID: core.StringPtr(subnetID),
			}

			createVPNServerOptions := vpcService.NewCreateVPNServerOptions(
				certificateInstanceIdentityModel,
				[]vpcv1.VPNServerAuthenticationPrototypeIntf{vpnServerAuthenticationPrototypeModel},
				"172.16.0.0/16",
				[]vpcv1.SubnetIdentityIntf{subnetIdentityModel},
			)
			createVPNServerOptions.SetName("my-vpn-server")

			vpnServer, response, err := vpcService.CreateVPNServer(createVPNServerOptions)
			if err != nil {
				panic(err)
			}

			// end-create_vpn_server
			vpnServerID = *vpnServer.ID
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(vpnServer).ToNot(BeNil())

		})
		It(`GetVPNServer request example`, func() {
			fmt.Println("\nGetVPNServer() result:")
			// begin-get_vpn_server

			getVPNServerOptions := vpcService.NewGetVPNServerOptions(
				vpnServerID,
			)

			vpnServer, response, err := vpcService.GetVPNServer(getVPNServerOptions)
			if err != nil {
				panic(err)
			}

			// end-get_vpn_server
			ifMatchVPNServer = response.GetHeaders()["Etag"][0]

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnServer).ToNot(BeNil())

		})
		It(`UpdateVPNServer request example`, func() {
			fmt.Println("\nUpdateVPNServer() result:")
			// begin-update_vpn_server

			vpnServerPatchModel := &vpcv1.VPNServerPatch{
				Name: &[]string{"my-vpn-server-modified"}[0],
			}
			vpnServerPatchModelAsPatch, asPatchErr := vpnServerPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateVPNServerOptions := vpcService.NewUpdateVPNServerOptions(
				vpnServerID,
				vpnServerPatchModelAsPatch,
			)
			updateVPNServerOptions.SetIfMatch(ifMatchVPNServer)

			vpnServer, response, err := vpcService.UpdateVPNServer(updateVPNServerOptions)
			if err != nil {
				panic(err)
			}

			// end-update_vpn_server

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnServer).ToNot(BeNil())

		})
		It(`GetVPNServerClientConfiguration request example`, func() {
			fmt.Println("\nGetVPNServerClientConfiguration() result:")
			// begin-get_vpn_server_client_configuration

			getVPNServerClientConfigurationOptions := vpcService.NewGetVPNServerClientConfigurationOptions(
				vpnServerID,
			)

			vpnServerClientConfiguration, response, err := vpcService.GetVPNServerClientConfiguration(getVPNServerClientConfigurationOptions)
			if err != nil {
				panic(err)
			}

			// end-get_vpn_server_client_configuration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnServerClientConfiguration).ToNot(BeNil())

		})
		It(`ListVPNServerClients request example`, func() {
			fmt.Println("\nListVPNServerClients() result:")
			// begin-list_vpn_server_clients

			listVPNServerClientsOptions := vpcService.NewListVPNServerClientsOptions(
				vpnServerID,
			)
			listVPNServerClientsOptions.SetSort("created_at")

			vpnServerClientCollection, response, err := vpcService.ListVPNServerClients(listVPNServerClientsOptions)
			if err != nil {
				panic(err)
			}

			// end-list_vpn_server_clients
			vpnClientID = *vpnServerClientCollection.Clients[0].ID
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnServerClientCollection).ToNot(BeNil())

		})
		It(`GetVPNServerClient request example`, func() {
			fmt.Println("\nGetVPNServerClient() result:")
			// begin-get_vpn_server_client

			getVPNServerClientOptions := vpcService.NewGetVPNServerClientOptions(
				vpnServerID,
				vpnClientID,
			)

			vpnServerClient, response, err := vpcService.GetVPNServerClient(getVPNServerClientOptions)
			if err != nil {
				panic(err)
			}

			// end-get_vpn_server_client

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnServerClient).ToNot(BeNil())

		})
		It(`DisconnectVPNClient request example`, func() {
			// begin-disconnect_vpn_client

			disconnectVPNClientOptions := vpcService.NewDisconnectVPNClientOptions(
				vpnServerID,
				vpnClientID,
			)

			response, err := vpcService.DisconnectVPNClient(disconnectVPNClientOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 202 {
				fmt.Printf("\nUnexpected response status code received from DisconnectVPNClient(): %d\n", response.StatusCode)
			}

			// end-disconnect_vpn_client

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})
		It(`ListVPNServerRoutes request example`, func() {
			fmt.Println("\nListVPNServerRoutes() result:")
			// begin-list_vpn_server_routes

			listVPNServerRoutesOptions := vpcService.NewListVPNServerRoutesOptions(
				vpnServerID,
			)
			listVPNServerRoutesOptions.SetSort("name")

			vpnServerRouteCollection, response, err := vpcService.ListVPNServerRoutes(listVPNServerRoutesOptions)
			if err != nil {
				panic(err)
			}

			// end-list_vpn_server_routes

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnServerRouteCollection).ToNot(BeNil())

		})
		It(`CreateVPNServerRoute request example`, func() {
			fmt.Println("\nCreateVPNServerRoute() result:")
			// begin-create_vpn_server_route

			createVPNServerRouteOptions := vpcService.NewCreateVPNServerRouteOptions(
				vpnServerID,
				"172.16.0.0/16",
			)
			createVPNServerRouteOptions.SetName("my-vpn-server-route")

			vpnServerRoute, response, err := vpcService.CreateVPNServerRoute(createVPNServerRouteOptions)
			if err != nil {
				panic(err)
			}

			// end-create_vpn_server_route
			vpnServerRouteID = *vpnServerRoute.ID
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(vpnServerRoute).ToNot(BeNil())

		})
		It(`GetVPNServerRoute request example`, func() {
			fmt.Println("\nGetVPNServerRoute() result:")
			// begin-get_vpn_server_route

			getVPNServerRouteOptions := vpcService.NewGetVPNServerRouteOptions(
				vpnServerID,
				vpnServerRouteID,
			)

			vpnServerRoute, response, err := vpcService.GetVPNServerRoute(getVPNServerRouteOptions)
			if err != nil {
				panic(err)
			}

			// end-get_vpn_server_route

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnServerRoute).ToNot(BeNil())

		})
		It(`UpdateVPNServerRoute request example`, func() {
			fmt.Println("\nUpdateVPNServerRoute() result:")
			// begin-update_vpn_server_route

			vpnServerRoutePatchModel := &vpcv1.VPNServerRoutePatch{
				Name: &[]string{"my-vpn-server-route-modified"}[0],
			}
			vpnServerRoutePatchModelAsPatch, asPatchErr := vpnServerRoutePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateVPNServerRouteOptions := vpcService.NewUpdateVPNServerRouteOptions(
				vpnServerID,
				vpnServerRouteID,
				vpnServerRoutePatchModelAsPatch,
			)

			vpnServerRoute, response, err := vpcService.UpdateVPNServerRoute(updateVPNServerRouteOptions)
			if err != nil {
				panic(err)
			}

			// end-update_vpn_server_route

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnServerRoute).ToNot(BeNil())

		})
		It(`ListLoadBalancerProfiles request example`, func() {
			fmt.Println("\nListLoadBalancerProfiles() result:")
			// begin-list_load_balancer_profiles

			options := &vpcv1.ListLoadBalancerProfilesOptions{}
			profiles, response, err := vpcService.ListLoadBalancerProfiles(options)

			// end-list_load_balancer_profiles
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profiles).ToNot(BeNil())

		})
		It(`GetLoadBalancerProfile request example`, func() {
			fmt.Println("\nGetLoadBalancerProfile() result:")
			// begin-get_load_balancer_profile
			options := &vpcv1.GetLoadBalancerProfileOptions{}
			options.SetName("network-fixed")
			profile, response, err := vpcService.GetLoadBalancerProfile(options)
			// end-get_load_balancer_profile
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profile).ToNot(BeNil())

		})
		It(`ListLoadBalancers request example`, func() {
			fmt.Println("\nListLoadBalancers() result:")
			// begin-list_load_balancers

			options := &vpcv1.ListLoadBalancersOptions{}
			loadBalancers, response, err := vpcService.ListLoadBalancers(options)

			// end-list_load_balancers
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancers).ToNot(BeNil())

		})
		It(`CreateLoadBalancer request example`, func() {
			fmt.Println("\nCreateLoadBalancer() result:")
			name := getName("lb")
			// begin-create_load_balancer

			options := &vpcv1.CreateLoadBalancerOptions{
				IsPublic: &[]bool{true}[0],
				Name:     &name,
				Subnets: []vpcv1.SubnetIdentityIntf{
					&vpcv1.SubnetIdentity{
						ID: &subnetID,
					},
				},
			}
			loadBalancer, response, err := vpcService.CreateLoadBalancer(options)
			// end-create_load_balancer

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(loadBalancer).ToNot(BeNil())
			loadBalancerID = *loadBalancer.ID
		})
		It(`GetLoadBalancer request example`, func() {
			fmt.Println("\nGetLoadBalancer() result:")
			// begin-get_load_balancer

			options := &vpcv1.GetLoadBalancerOptions{
				ID: &loadBalancerID,
			}
			loadBalancer, response, err := vpcService.GetLoadBalancer(options)

			// end-get_load_balancer
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancer).ToNot(BeNil())

		})
		It(`UpdateLoadBalancer request example`, func() {
			fmt.Println("\nUpdateLoadBalancer() result:")
			name := getName("lb")
			// begin-update_load_balancer

			loadBalancerPatchModel := &vpcv1.LoadBalancerPatch{
				Name: &name,
			}
			loadBalancerPatchModelAsPatch, asPatchErr := loadBalancerPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			updateLoadBalancerOptions := vpcService.NewUpdateLoadBalancerOptions(
				loadBalancerID,
				loadBalancerPatchModelAsPatch,
			)

			loadBalancer, response, err := vpcService.UpdateLoadBalancer(updateLoadBalancerOptions)

			// end-update_load_balancer
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancer).ToNot(BeNil())

		})
		It(`GetLoadBalancerStatistics request example`, func() {
			fmt.Println("\nGetLoadBalancerStatistics() result:")
			// begin-get_load_balancer_statistics

			options := &vpcv1.GetLoadBalancerStatisticsOptions{
				ID: &loadBalancerID,
			}
			statistics, response, err := vpcService.GetLoadBalancerStatistics(options)
			// end-get_load_balancer_statistics
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(statistics).ToNot(BeNil())

		})
		It(`ListLoadBalancerListeners request example`, func() {
			fmt.Println("\nListLoadBalancerListeners() result:")
			// begin-list_load_balancer_listeners

			options := &vpcv1.ListLoadBalancerListenersOptions{
				LoadBalancerID: &loadBalancerID,
			}
			listeners, response, err := vpcService.ListLoadBalancerListeners(options)

			// end-list_load_balancer_listeners
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listeners).ToNot(BeNil())

		})
		It(`CreateLoadBalancerListener request example`, func() {
			fmt.Println("\nCreateLoadBalancerListener() result:")
			// begin-create_load_balancer_listener

			options := &vpcv1.CreateLoadBalancerListenerOptions{
				LoadBalancerID: &loadBalancerID,
			}
			options.SetPort(5656)
			options.SetProtocol("http")
			listener, response, err := vpcService.CreateLoadBalancerListener(options)

			// end-create_load_balancer_listener
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(listener).ToNot(BeNil())
			listenerID = *listener.ID
		})
		It(`GetLoadBalancerListener request example`, func() {
			fmt.Println("\nGetLoadBalancerListener() result:")
			// begin-get_load_balancer_listener

			options := &vpcv1.GetLoadBalancerListenerOptions{
				LoadBalancerID: &loadBalancerID,
				ID:             &listenerID,
			}
			listener, response, err := vpcService.GetLoadBalancerListener(options)

			// end-get_load_balancer_listener
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listener).ToNot(BeNil())

		})
		It(`UpdateLoadBalancerListener request example`, func() {
			fmt.Println("\nUpdateLoadBalancerListener() result:")
			// begin-update_load_balancer_listener

			loadBalancerListenerPatchModel := &vpcv1.LoadBalancerListenerPatch{
				Port: &[]int64{5666}[0],
			}
			loadBalancerListenerPatchModelAsPatch, asPatchErr := loadBalancerListenerPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options := vpcService.NewUpdateLoadBalancerListenerOptions(
				loadBalancerID,
				listenerID,
				loadBalancerListenerPatchModelAsPatch,
			)

			listener, response, err := vpcService.UpdateLoadBalancerListener(options)

			// end-update_load_balancer_listener
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listener).ToNot(BeNil())

		})
		It(`ListLoadBalancerListenerPolicies request example`, func() {
			fmt.Println("\nListLoadBalancerListenerPolicies() result:")
			// begin-list_load_balancer_listener_policies

			options := &vpcv1.ListLoadBalancerListenerPoliciesOptions{
				LoadBalancerID: &loadBalancerID,
				ListenerID:     &listenerID,
			}
			policies, response, err :=
				vpcService.ListLoadBalancerListenerPolicies(options)

			// end-list_load_balancer_listener_policies
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policies).ToNot(BeNil())

		})
		It(`CreateLoadBalancerListenerPolicy request example`, func() {
			fmt.Println("\nCreateLoadBalancerListenerPolicy() result:")
			// begin-create_load_balancer_listener_policy

			options := &vpcv1.CreateLoadBalancerListenerPolicyOptions{
				LoadBalancerID: &loadBalancerID,
				ListenerID:     &listenerID,
			}
			options.SetPriority(2)
			options.SetAction("reject")
			policy, response, err :=
				vpcService.CreateLoadBalancerListenerPolicy(options)

			// end-create_load_balancer_listener_policy
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(policy).ToNot(BeNil())
			policyID = *policy.ID
		})
		It(`GetLoadBalancerListenerPolicy request example`, func() {
			fmt.Println("\nGetLoadBalancerListenerPolicy() result:")
			// begin-get_load_balancer_listener_policy

			options := &vpcv1.GetLoadBalancerListenerPolicyOptions{}
			options.SetLoadBalancerID(loadBalancerID)
			options.SetListenerID(listenerID)
			options.SetID(policyID)
			policy, response, err := vpcService.GetLoadBalancerListenerPolicy(options)

			// end-get_load_balancer_listener_policy
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policy).ToNot(BeNil())

		})
		It(`UpdateLoadBalancerListenerPolicy request example`, func() {
			fmt.Println("\nUpdateLoadBalancerListenerPolicy() result:")
			// begin-update_load_balancer_listener_policy

			options := &vpcv1.UpdateLoadBalancerListenerPolicyOptions{}
			options.SetLoadBalancerID(loadBalancerID)
			options.SetListenerID(listenerID)
			options.SetID(policyID)
			policyPatchModel := &vpcv1.LoadBalancerListenerPolicyPatch{}
			policyPatchModel.Priority = &[]int64{5}[0]
			policyPatch, asPatchErr :=
				policyPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.LoadBalancerListenerPolicyPatch = policyPatch
			policy, response, err :=
				vpcService.UpdateLoadBalancerListenerPolicy(options)

			// end-update_load_balancer_listener_policy
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policy).ToNot(BeNil())
		})
		It(`ListLoadBalancerListenerPolicyRules request example`, func() {
			fmt.Println("\nListLoadBalancerListenerPolicyRules() result:")
			// begin-list_load_balancer_listener_policy_rules

			options := &vpcv1.ListLoadBalancerListenerPolicyRulesOptions{}
			options.SetLoadBalancerID(loadBalancerID)
			options.SetListenerID(listenerID)
			options.SetPolicyID(policyID)
			rules, response, err :=
				vpcService.ListLoadBalancerListenerPolicyRules(options)
			// end-list_load_balancer_listener_policy_rules
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rules).ToNot(BeNil())

		})
		It(`CreateLoadBalancerListenerPolicyRule request example`, func() {
			fmt.Println("\nCreateLoadBalancerListenerPolicyRule() result:")
			// begin-create_load_balancer_listener_policy_rule
			options := &vpcv1.CreateLoadBalancerListenerPolicyRuleOptions{}
			options.SetLoadBalancerID(loadBalancerID)
			options.SetListenerID(listenerID)
			options.SetPolicyID(policyID)
			options.SetCondition("contains")
			options.SetType("hostname")
			options.SetValue("one")
			policyRule, response, err :=
				vpcService.CreateLoadBalancerListenerPolicyRule(options)

			// end-create_load_balancer_listener_policy_rule
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(policyRule).ToNot(BeNil())
			policyRuleID = *policyRule.ID
		})
		It(`GetLoadBalancerListenerPolicyRule request example`, func() {
			fmt.Println("\nGetLoadBalancerListenerPolicyRule() result:")
			// begin-get_load_balancer_listener_policy_rule

			options := &vpcv1.GetLoadBalancerListenerPolicyRuleOptions{}
			options.SetLoadBalancerID(loadBalancerID)
			options.SetListenerID(listenerID)
			options.SetPolicyID(policyID)
			options.SetID(policyRuleID)
			rule, response, err :=
				vpcService.GetLoadBalancerListenerPolicyRule(options)

			// end-get_load_balancer_listener_policy_rule
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())

		})
		It(`UpdateLoadBalancerListenerPolicyRule request example`, func() {
			fmt.Println("\nUpdateLoadBalancerListenerPolicyRule() result:")
			// begin-update_load_balancer_listener_policy_rule

			options := &vpcv1.UpdateLoadBalancerListenerPolicyRuleOptions{}
			options.SetLoadBalancerID(loadBalancerID)
			options.SetListenerID(listenerID)
			options.SetPolicyID(policyID)
			options.SetID(policyRuleID)
			policyRulePatchModel :=
				&vpcv1.LoadBalancerListenerPolicyRulePatch{
					Condition: &[]string{"contains"}[0],
					Type:      &[]string{"header"}[0],
					Value:     &[]string{"app"}[0],
					Field:     &[]string{"MY-APP-HEADER"}[0],
				}
			policyRulePatch, asPatchErr := policyRulePatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.LoadBalancerListenerPolicyRulePatch = policyRulePatch
			rule, response, err :=
				vpcService.UpdateLoadBalancerListenerPolicyRule(options)

			// end-update_load_balancer_listener_policy_rule
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())

		})
		It(`ListLoadBalancerPools request example`, func() {
			fmt.Println("\nListLoadBalancerPools() result:")
			// begin-list_load_balancer_pools
			options := &vpcv1.ListLoadBalancerPoolsOptions{}
			options.SetLoadBalancerID(loadBalancerID)
			pools, response, err := vpcService.ListLoadBalancerPools(options)
			// end-list_load_balancer_pools
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pools).ToNot(BeNil())

		})
		It(`CreateLoadBalancerPool request example`, func() {
			fmt.Println("\nCreateLoadBalancerPool() result:")
			name := getName("pool")
			// begin-create_load_balancer_pool

			options := &vpcv1.CreateLoadBalancerPoolOptions{}
			options.SetLoadBalancerID(loadBalancerID)
			options.SetAlgorithm("round_robin")
			options.SetHealthMonitor(&vpcv1.LoadBalancerPoolHealthMonitorPrototype{
				Delay:      &[]int64{30}[0],
				MaxRetries: &[]int64{3}[0],
				Timeout:    &[]int64{30}[0],
				Type:       &[]string{"http"}[0],
			})
			options.SetName(name)
			options.SetProtocol("http")
			pool, response, err := vpcService.CreateLoadBalancerPool(options)

			// end-create_load_balancer_pool
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(pool).ToNot(BeNil())
			poolID = *pool.ID
		})
		It(`GetLoadBalancerPool request example`, func() {
			fmt.Println("\nGetLoadBalancerPool() result:")
			// begin-get_load_balancer_pool

			options := &vpcv1.GetLoadBalancerPoolOptions{
				LoadBalancerID: &loadBalancerID,
				ID:             &poolID,
			}
			pool, response, err := vpcService.GetLoadBalancerPool(options)

			// end-get_load_balancer_pool
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pool).ToNot(BeNil())

		})
		It(`UpdateLoadBalancerPool request example`, func() {
			fmt.Println("\nUpdateLoadBalancerPool() result:")
			// begin-update_load_balancer_pool

			options := &vpcv1.UpdateLoadBalancerPoolOptions{
				LoadBalancerID: &loadBalancerID,
				ID:             &poolID,
			}
			poolPatchModel := &vpcv1.LoadBalancerPoolPatch{}
			healthMonitorPatchModel := &vpcv1.LoadBalancerPoolHealthMonitorPatch{
				Delay:      &[]int64{30}[0],
				MaxRetries: &[]int64{3}[0],
				Timeout:    &[]int64{30}[0],
				Type:       &[]string{"http"}[0],
			}
			poolPatchModel.HealthMonitor = healthMonitorPatchModel
			sessionPersistence := &vpcv1.LoadBalancerPoolSessionPersistencePatch{
				Type: &[]string{"http_cookie"}[0],
			}
			poolPatchModel.SessionPersistence = sessionPersistence
			LoadBalancerPoolPatch, _ := poolPatchModel.AsPatch()
			options.LoadBalancerPoolPatch = LoadBalancerPoolPatch
			pool, response, err := vpcService.UpdateLoadBalancerPool(options)

			// end-update_load_balancer_pool
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pool).ToNot(BeNil())

		})
		It(`ListLoadBalancerPoolMembers request example`, func() {
			fmt.Println("\nListLoadBalancerPoolMembers() result:")
			// begin-list_load_balancer_pool_members

			options := &vpcv1.ListLoadBalancerPoolMembersOptions{
				LoadBalancerID: &loadBalancerID,
				PoolID:         &poolID,
			}
			members, response, err := vpcService.ListLoadBalancerPoolMembers(options)

			// end-list_load_balancer_pool_members
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(members).ToNot(BeNil())

		})
		It(`CreateLoadBalancerPoolMember request example`, func() {
			fmt.Println("\nCreateLoadBalancerPoolMember() result:")
			// begin-create_load_balancer_pool_member

			options := &vpcv1.CreateLoadBalancerPoolMemberOptions{
				LoadBalancerID: &loadBalancerID,
				PoolID:         &poolID,
				Port:           &[]int64{1234}[0],
				Target: &vpcv1.LoadBalancerPoolMemberTargetPrototypeIP{
					Address: &[]string{"192.168.3.4"}[0],
				},
			}
			member, response, err := vpcService.CreateLoadBalancerPoolMember(options)
			// end-create_load_balancer_pool_member
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(member).ToNot(BeNil())
			poolMemberID = *member.ID
		})

		It(`GetLoadBalancerPoolMember request example`, func() {
			fmt.Println("\nGetLoadBalancerPoolMember() result:")
			// begin-get_load_balancer_pool_member

			options := &vpcv1.GetLoadBalancerPoolMemberOptions{}
			options.SetLoadBalancerID(loadBalancerID)
			options.SetPoolID(poolID)
			options.SetID(poolMemberID)
			member, response, err := vpcService.GetLoadBalancerPoolMember(options)

			// end-get_load_balancer_pool_member
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(member).ToNot(BeNil())

		})
		It(`UpdateLoadBalancerPoolMember request example`, func() {
			fmt.Println("\nUpdateLoadBalancerPoolMember() result:")
			// begin-update_load_balancer_pool_member

			options := &vpcv1.UpdateLoadBalancerPoolMemberOptions{}
			options.SetLoadBalancerID(loadBalancerID)
			options.SetPoolID(poolID)
			options.SetID(poolMemberID)
			loadBalancerPoolMemberPatchModel := &vpcv1.LoadBalancerPoolMemberPatch{
				Port:   &[]int64{1235}[0],
				Weight: &[]int64{50}[0],
			}
			loadBalancerPoolMemberPatch, _ := loadBalancerPoolMemberPatchModel.AsPatch()
			options.LoadBalancerPoolMemberPatch = loadBalancerPoolMemberPatch
			member, response, err := vpcService.UpdateLoadBalancerPoolMember(options)

			// end-update_load_balancer_pool_member
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(member).ToNot(BeNil())

		})
		It(`ReplaceLoadBalancerPoolMembers request example`, func() {
			fmt.Println("\nReplaceLoadBalancerPoolMembers() result:")
			// begin-replace_load_balancer_pool_members

			options := &vpcv1.ReplaceLoadBalancerPoolMembersOptions{}
			options.SetLoadBalancerID(loadBalancerID)
			options.SetPoolID(poolID)
			options.SetMembers([]vpcv1.LoadBalancerPoolMemberPrototype{
				{
					Port: &[]int64{1235}[0],
					Target: &vpcv1.LoadBalancerPoolMemberTargetPrototypeIP{
						Address: &[]string{"192.168.3.5"}[0],
					},
				},
			})
			members, response, err :=
				vpcService.ReplaceLoadBalancerPoolMembers(options)

			// end-replace_load_balancer_pool_members
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(members).ToNot(BeNil())
			poolMemberID = *members.Members[0].ID
		})
		It(`DeleteLoadBalancerPoolMember request example`, func() {
			// begin-delete_load_balancer_pool_member

			options := &vpcv1.DeleteLoadBalancerPoolMemberOptions{}
			options.SetLoadBalancerID(loadBalancerID)
			options.SetPoolID(poolID)
			options.SetID(poolMemberID)
			response, err := vpcService.DeleteLoadBalancerPoolMember(options)

			// end-delete_load_balancer_pool_member
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteLoadBalancerPoolMember() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteLoadBalancerPool request example`, func() {
			// begin-delete_load_balancer_pool

			options := &vpcv1.DeleteLoadBalancerPoolOptions{}
			options.SetLoadBalancerID(loadBalancerID)
			options.SetID(poolID)
			response, err := vpcService.DeleteLoadBalancerPool(options)

			// end-delete_load_balancer_pool
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteLoadBalancerPool() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteLoadBalancerListenerPolicyRule request example`, func() {
			// begin-delete_load_balancer_listener_policy_rule

			options := &vpcv1.DeleteLoadBalancerListenerPolicyRuleOptions{}
			options.SetLoadBalancerID(loadBalancerID)
			options.SetListenerID(listenerID)
			options.SetPolicyID(policyID)
			options.SetID(policyRuleID)
			response, err :=
				vpcService.DeleteLoadBalancerListenerPolicyRule(options)

			// end-delete_load_balancer_listener_policy_rule
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteLoadBalancerListenerPolicyRule() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteLoadBalancerListenerPolicy request example`, func() {
			// begin-delete_load_balancer_listener_policy

			options := &vpcv1.DeleteLoadBalancerListenerPolicyOptions{}
			options.SetLoadBalancerID(loadBalancerID)
			options.SetListenerID(listenerID)
			options.SetID(policyID)
			response, err := vpcService.DeleteLoadBalancerListenerPolicy(options)

			// end-delete_load_balancer_listener_policy
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteLoadBalancerListenerPolicy() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteLoadBalancerListener request example`, func() {
			// begin-delete_load_balancer_listener

			options := &vpcv1.DeleteLoadBalancerListenerOptions{}
			options.SetLoadBalancerID(loadBalancerID)
			options.SetID(listenerID)
			response, err := vpcService.DeleteLoadBalancerListener(options)

			// end-delete_load_balancer_listener
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteLoadBalancerListener() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteLoadBalancer request example`, func() {
			// begin-delete_load_balancer
			deleteVpcOptions := &vpcv1.DeleteLoadBalancerOptions{}
			deleteVpcOptions.SetID(loadBalancerID)
			response, err := vpcService.DeleteLoadBalancer(deleteVpcOptions)

			// end-delete_load_balancer
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteLoadBalancer() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`ListEndpointGateways request example`, func() {
			fmt.Println("\nListEndpointGateways() result:")
			// begin-list_endpoint_gateways

			options := vpcService.NewListEndpointGatewaysOptions()
			endpointGateways, response, err :=
				vpcService.ListEndpointGateways(options)

			// end-list_endpoint_gateways
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(endpointGateways).ToNot(BeNil())

		})
		It(`CreateEndpointGateway request example`, func() {
			fmt.Println("\nCreateEndpointGateway() result:")
			name := getName("egw")
			// begin-create_endpoint_gateway

			options := &vpcv1.CreateEndpointGatewayOptions{}
			options.SetName(name)
			options.SetVPC(&vpcv1.VPCIdentity{
				ID: &vpcID,
			})

			targetName := "ibm-ntp-server"
			providerInfrastructureService := "provider_infrastructure_service"
			options.SetTarget(
				&vpcv1.EndpointGatewayTargetPrototype{
					ResourceType: &providerInfrastructureService,
					Name:         &targetName,
				},
			)
			endpointGateway, response, err := vpcService.CreateEndpointGateway(options)

			// end-create_endpoint_gateway
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(endpointGateway).ToNot(BeNil())
			endpointGatewayID = *endpointGateway.ID
		})
		It(`ListEndpointGatewayIps request example`, func() {
			fmt.Println("\nListEndpointGatewayIps() result:")
			// begin-list_endpoint_gateway_ips

			options := vpcService.NewListEndpointGatewayIpsOptions(endpointGatewayID)
			reservedIPs, response, err := vpcService.ListEndpointGatewayIps(options)

			// end-list_endpoint_gateway_ips
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reservedIPs).ToNot(BeNil())

		})
		It(`AddEndpointGatewayIP request example`, func() {
			fmt.Println("\nAddEndpointGatewayIP() result:")
			// begin-add_endpoint_gateway_ip

			options := vpcService.NewAddEndpointGatewayIPOptions(endpointGatewayID, reservedIPID)
			reservedIP, response, err := vpcService.AddEndpointGatewayIP(options)

			// end-add_endpoint_gateway_ip
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(reservedIP).ToNot(BeNil())
			endpointGatewayTargetID = *reservedIP.ID
		})
		It(`GetEndpointGateway request example`, func() {
			fmt.Println("\nGetEndpointGateway() result:")
			// begin-get_endpoint_gateway_ip

			options := vpcService.NewGetEndpointGatewayOptions(endpointGatewayID)
			endpointGateway, response, err := vpcService.GetEndpointGateway(options)

			// end-get_endpoint_gateway_ip
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(endpointGateway).ToNot(BeNil())

		})

		It(`GetEndpointGatewayIP request example`, func() {
			fmt.Println("\nGetEndpointGatewayIP() result:")
			// begin-get_endpoint_gateway

			options := vpcService.NewGetEndpointGatewayIPOptions(endpointGatewayID, endpointGatewayTargetID)
			reservedIP, response, err := vpcService.GetEndpointGatewayIP(options)

			// end-get_endpoint_gateway
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reservedIP).ToNot(BeNil())

		})
		It(`UpdateEndpointGateway request example`, func() {
			fmt.Println("\nUpdateEndpointGateway() result:")
			name := getName("egw")
			// begin-update_endpoint_gateway

			endpointGatewayPatchModel := new(vpcv1.EndpointGatewayPatch)
			endpointGatewayPatchModel.Name = &name
			endpointGatewayPatchModelAsPatch, asPatchErr := endpointGatewayPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options := &vpcv1.UpdateEndpointGatewayOptions{
				ID:                   &endpointGatewayID,
				EndpointGatewayPatch: endpointGatewayPatchModelAsPatch,
			}
			endpointGateway, response, err := vpcService.UpdateEndpointGateway(options)

			// end-update_endpoint_gateway
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(endpointGateway).ToNot(BeNil())

		})

		It(`RemoveEndpointGatewayIP request example`, func() {
			// begin-remove_endpoint_gateway_ip

			removeEndpointGatewayIPOptions := vpcService.NewRemoveEndpointGatewayIPOptions(
				endpointGatewayID,
				endpointGatewayTargetID,
			)

			response, err := vpcService.RemoveEndpointGatewayIP(removeEndpointGatewayIPOptions)

			// end-remove_endpoint_gateway_ip
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nRemoveEndpointGatewayIP() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteEndpointGateway request example`, func() {
			// begin-delete_endpoint_gateway

			deleteEndpointGatewayOptions := vpcService.NewDeleteEndpointGatewayOptions(
				endpointGatewayID,
			)

			response, err := vpcService.DeleteEndpointGateway(deleteEndpointGatewayOptions)

			// end-delete_endpoint_gateway
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteEndpointGateway() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})

		It(`ListFlowLogCollectors request example`, func() {
			fmt.Println("\nListFlowLogCollectors() result:")
			// begin-list_flow_log_collectors

			options := &vpcv1.ListFlowLogCollectorsOptions{}
			flowLogs, response, err := vpcService.ListFlowLogCollectors(options)

			// end-list_flow_log_collectors
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(flowLogs).ToNot(BeNil())

		})
		It(`CreateFlowLogCollector request example`, func() {
			fmt.Println("\nCreateFlowLogCollector() result:")
			name := getName("flowlog")
			// begin-create_flow_log_collector

			options := &vpcv1.CreateFlowLogCollectorOptions{}
			options.SetName(name)
			options.SetTarget(&vpcv1.FlowLogCollectorTargetPrototypeVPCIdentity{
				ID: &vpcID,
			})
			options.SetStorageBucket(&vpcv1.LegacyCloudObjectStorageBucketIdentity{
				Name: &[]string{"bucket-name"}[0],
			})
			flowLog, response, err := vpcService.CreateFlowLogCollector(options)

			// end-create_flow_log_collector
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(flowLog).ToNot(BeNil())
			flowLogID = *flowLog.ID
		})
		It(`GetFlowLogCollector request example`, func() {
			fmt.Println("\nGetFlowLogCollector() result:")
			// begin-get_flow_log_collector

			options := &vpcv1.GetFlowLogCollectorOptions{}
			options.SetID(flowLogID)
			flowLog, response, err := vpcService.GetFlowLogCollector(options)

			// end-get_flow_log_collector
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(flowLog).ToNot(BeNil())

		})

		It(`UpdateFlowLogCollector request example`, func() {
			fmt.Println("\nUpdateFlowLogCollector() result:")
			name := getName("fl")
			// begin-update_flow_log_collector

			options := &vpcv1.UpdateFlowLogCollectorOptions{}
			options.SetID(flowLogID)
			flowLogCollectorPatchModel := &vpcv1.FlowLogCollectorPatch{
				Active: &[]bool{true}[0],
				Name:   &name,
			}
			flowLogCollectorPatch, asPatchErr := flowLogCollectorPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.FlowLogCollectorPatch = flowLogCollectorPatch
			flowLog, response, err := vpcService.UpdateFlowLogCollector(options)

			// end-update_flow_log_collector
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(flowLog).ToNot(BeNil())

		})

		It(`DeleteFlowLogCollector request example`, func() {
			// begin-delete_flow_log_collector

			deleteFlowLogCollectorOptions := vpcService.NewDeleteFlowLogCollectorOptions(
				flowLogID,
			)

			response, err := vpcService.DeleteFlowLogCollector(deleteFlowLogCollectorOptions)

			// end-delete_flow_log_collector
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteFlowLogCollector() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})

		It(`DeleteBareMetalServerNetworkInterface request example`, func() {
			// begin-delete_bare_metal_server_network_interface

			deleteBareMetalServerNetworkInterfaceOptions := &vpcv1.DeleteBareMetalServerNetworkInterfaceOptions{
				BareMetalServerID: &bareMetalServerId,
				ID:                &bareMetalServerNetworkInterfaceId,
			}

			response, err := vpcService.DeleteBareMetalServerNetworkInterface(deleteBareMetalServerNetworkInterfaceOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteBareMetalServerNetworkInterface(): %d\n", response.StatusCode)
			}
			// end-delete_bare_metal_server_network_interface

			fmt.Printf("\nDeleteBareMetalServerNetworkInterface() response status code: %d\n", response.StatusCode)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})

		It(`DeleteBareMetalServer request example`, func() {
			// begin-delete_bare_metal_server

			deleteBareMetalServerOptions := &vpcv1.DeleteBareMetalServerOptions{
				ID: &bareMetalServerId,
			}

			response, err := vpcService.DeleteBareMetalServer(deleteBareMetalServerOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteBareMetalServer(): %d\n", response.StatusCode)
			}
			// end-delete_bare_metal_server

			fmt.Printf("\nDeleteBareMetalServer() response status code: %d\n", response.StatusCode)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})

		It(`RemoveVPNGatewayConnectionPeerCIDR request example`, func() {
			// begin-remove_vpn_gateway_connection_peer_cidr

			options := &vpcv1.RemoveVPNGatewayConnectionPeerCIDROptions{}
			options.SetVPNGatewayID(vpnGatewayID)
			options.SetID(vpnGatewayConnectionID)
			options.SetCIDRPrefix("192.144.0.0")
			options.SetPrefixLength("28")
			response, err := vpcService.RemoveVPNGatewayConnectionPeerCIDR(options)

			// end-remove_vpn_gateway_connection_peer_cidr
			fmt.Printf("\nRemoveVPNGatewayConnectionPeerCIDR() response status code: %d\n", response.StatusCode)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`RemoveVPNGatewayConnectionLocalCIDR request example`, func() {
			// begin-remove_vpn_gateway_connection_local_cidr

			options := &vpcv1.RemoveVPNGatewayConnectionLocalCIDROptions{}
			options.SetVPNGatewayID(vpnGatewayID)
			options.SetID(vpnGatewayConnectionID)
			options.SetCIDRPrefix("192.134.0.0")
			options.SetPrefixLength("28")
			response, err := vpcService.RemoveVPNGatewayConnectionLocalCIDR(options)

			// end-remove_vpn_gateway_connection_local_cidr
			fmt.Printf("\nRemoveVPNGatewayConnectionLocalCIDR() response status code: %d\n", response.StatusCode)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})

		It(`RemoveInstanceNetworkInterfaceFloatingIP request example`, func() {
			// begin-remove_instance_network_interface_floating_ip

			options := &vpcv1.RemoveInstanceNetworkInterfaceFloatingIPOptions{}
			options.SetID(floatingIPID)
			options.SetInstanceID(instanceID)
			options.SetNetworkInterfaceID(eth2ID)
			response, err :=
				vpcService.RemoveInstanceNetworkInterfaceFloatingIP(options)

			// end-remove_instance_network_interface_floating_ip
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nRemoveInstanceNetworkInterfaceFloatingIP() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteSecurityGroupTargetBinding request example`, func() {
			// begin-delete_security_group_target_binding

			options := &vpcv1.DeleteSecurityGroupTargetBindingOptions{}
			options.SetSecurityGroupID(securityGroupID)
			options.SetID(targetID)
			response, err :=
				vpcService.DeleteSecurityGroupTargetBinding(options)

			// end-delete_security_group_target_binding
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteSecurityGroupTargetBinding() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteInstanceNetworkInterface request example`, func() {
			// begin-delete_instance_network_interface

			options := &vpcv1.DeleteInstanceNetworkInterfaceOptions{}
			options.SetID(eth2ID)
			options.SetInstanceID(instanceID)
			response, err := vpcService.DeleteInstanceNetworkInterface(options)

			// end-delete_instance_network_interface
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteInstanceNetworkInterface() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteInstanceVolumeAttachment request example`, func() {
			// begin-delete_instance_volume_attachment

			options := &vpcv1.DeleteInstanceVolumeAttachmentOptions{}
			options.SetID(volumeAttachmentID)
			options.SetInstanceID(instanceID)
			response, err := vpcService.DeleteInstanceVolumeAttachment(options)

			// end-delete_instance_volume_attachment
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteInstanceVolumeAttachment() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteVolume request example`, func() {
			// begin-delete_volume

			options := &vpcv1.DeleteVolumeOptions{}
			options.SetID(volumeID)
			options.SetIfMatch(ifMatchVolume)
			response, err := vpcService.DeleteVolume(options)

			// end-delete_volume
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteVolume() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteFloatingIP request example`, func() {
			// begin-delete_floating_ip

			options := vpcService.NewDeleteFloatingIPOptions(floatingIPID)
			response, err := vpcService.DeleteFloatingIP(options)

			// end-delete_floating_ip
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteFloatingIP() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})

		It(`DeleteInstanceTemplate request example`, func() {
			// begin-delete_instance_template

			options := &vpcv1.DeleteInstanceTemplateOptions{}
			options.SetID(instanceTemplateID)
			response, err := vpcService.DeleteInstanceTemplate(options)

			// end-delete_instance_template
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteInstanceTemplate() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})

		It(`DeleteInstance request example`, func() {
			// begin-delete_instance

			options := &vpcv1.DeleteInstanceOptions{}
			options.SetID(instanceID)
			response, err := vpcService.DeleteInstance(options)
			// end-delete_instance
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteInstance() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteKey request example`, func() {
			// begin - delete_key

			deleteKeyOptions := &vpcv1.DeleteKeyOptions{}
			deleteKeyOptions.SetID(keyID)
			response, err := vpcService.DeleteKey(deleteKeyOptions)

			// end-delete_key

			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteKey() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteImage request example`, func() {
			// begin-delete_image

			options := &vpcv1.DeleteImageOptions{}
			options.SetID(imageID)
			response, err := vpcService.DeleteImage(options)
			// end-delete_image
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteImage() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})

		It(`DeleteSubnet request example`, func() {
			// begin-delete_subnet

			options := &vpcv1.DeleteSubnetOptions{}
			options.SetID(subnetID)
			response, err := vpcService.DeleteSubnet(options)

			// end-delete_subnet
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteSubnet() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteVPCRoutingTableRoute request example`, func() {
			// begin-delete_vpc_routing_table_route

			options := &vpcv1.DeleteVPCRoutingTableRouteOptions{
				VPCID:          &vpcID,
				RoutingTableID: &routingTableID,
				ID:             &routeID,
			}
			response, err := vpcService.DeleteVPCRoutingTableRoute(options)

			// end-delete_vpc_routing_table_route
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteVPCRoutingTableRoute() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteVPCRoutingTable request example`, func() {
			// begin-delete_vpc_routing_table

			options := vpcService.NewDeleteVPCRoutingTableOptions(
				vpcID,
				routingTableID,
			)
			response, err := vpcService.DeleteVPCRoutingTable(options)

			// end-delete_vpc_routing_table
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteVPCRoutingTable() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteVPCAddressPrefix request example`, func() {
			// begin-delete_vpc_address_prefix

			deleteVpcAddressPrefixOptions := &vpcv1.DeleteVPCAddressPrefixOptions{}
			deleteVpcAddressPrefixOptions.SetVPCID(vpcID)
			deleteVpcAddressPrefixOptions.SetID(addressPrefixID)
			response, err :=
				vpcService.DeleteVPCAddressPrefix(deleteVpcAddressPrefixOptions)

			// end-delete_vpc_address_prefix
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteVPCAddressPrefix() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteVPC request example`, func() {
			// begin-delete_vpc

			deleteVpcOptions := &vpcv1.DeleteVPCOptions{}
			deleteVpcOptions.SetID(vpcID)
			response, err := vpcService.DeleteVPC(deleteVpcOptions)
			// end-delete_vpc
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteVPC() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})

		It(`DeleteSnapshot request example`, func() {
			// begin-delete_snapshot
			options := &vpcv1.DeleteSnapshotOptions{
				ID:      &snapshotID,
				IfMatch: &ifMatchSnapshot,
			}
			response, err := vpcService.DeleteSnapshot(options)

			// end-delete_snapshot
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteSnapshot() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteSnapshots request example`, func() {
			// begin-delete_snapshots

			options := &vpcv1.DeleteSnapshotsOptions{
				SourceVolumeID: &volumeID,
			}
			response, err := vpcService.DeleteSnapshots(options)

			// end-delete_snapshots
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteSnapshots() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})

		It(`DeleteSecurityGroupRule request example`, func() {
			// begin-delete_security_group_rule

			options := &vpcv1.DeleteSecurityGroupRuleOptions{}
			options.SetSecurityGroupID(securityGroupID)
			options.SetID(securityGroupRuleID)
			response, err := vpcService.DeleteSecurityGroupRule(options)
			// end-delete_security_group_rule
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteSecurityGroupRule() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteSecurityGroup request example`, func() {
			// begin-delete_security_group

			options := &vpcv1.DeleteSecurityGroupOptions{}
			options.SetID(securityGroupID)
			response, err := vpcService.DeleteSecurityGroup(options)

			// end-delete_security_group
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteSecurityGroup() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeletePublicGateway request example`, func() {
			// begin-delete_public_gateway

			options := &vpcv1.DeletePublicGatewayOptions{}
			options.SetID(publicGatewayID)
			response, err := vpcService.DeletePublicGateway(options)

			// end-delete_public_gateway
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeletePublicGateway() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteNetworkACLRule request example`, func() {
			// begin-delete_network_acl_rule

			options := &vpcv1.DeleteNetworkACLRuleOptions{}
			options.SetID(networkACLRuleID)
			options.SetNetworkACLID(networkACLID)
			response, err := vpcService.DeleteNetworkACLRule(options)

			// end-delete_network_acl_rule
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteNetworkACLRule() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteNetworkACL request example`, func() {
			// begin-delete_network_acl

			options := &vpcv1.DeleteNetworkACLOptions{}
			options.SetID(networkACLID)
			response, err := vpcService.DeleteNetworkACL(options)

			// end-delete_network_acl
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteNetworkACL() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})

		It(`DeleteInstanceGroupMembership request example`, func() {
			// begin-delete_instance_group_membership

			options := &vpcv1.DeleteInstanceGroupMembershipOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			options.SetID(instanceGroupMembershipID)
			response, err := vpcService.DeleteInstanceGroupMembership(options)

			// end-delete_instance_group_membership
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteInstanceGroupMembership() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteInstanceGroupMemberships request example`, func() {
			// begin-delete_instance_group_memberships

			options := &vpcv1.DeleteInstanceGroupMembershipsOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			response, err := vpcService.DeleteInstanceGroupMemberships(options)

			// end-delete_instance_group_memberships
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteInstanceGroupMemberships() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteInstanceGroupManagerPolicy request example`, func() {
			// begin-delete_instance_group_manager_policy
			options := &vpcv1.DeleteInstanceGroupManagerPolicyOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			options.SetInstanceGroupManagerID(instanceGroupManagerID)
			options.SetID(instanceGroupManagerPolicyID)
			response, err := vpcService.DeleteInstanceGroupManagerPolicy(options)

			// end-delete_instance_group_manager_policy
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteInstanceGroupManagerPolicy() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteInstanceGroupManagerAction request example`, func() {
			// begin-delete_instance_group_manager_action

			options := &vpcv1.DeleteInstanceGroupManagerActionOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			options.SetInstanceGroupManagerID(instanceGroupManagerID)
			options.SetID(instanceGroupManagerActionID)
			response, err := vpcService.DeleteInstanceGroupManagerAction(options)

			// end-delete_instance_group_manager_action
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteInstanceGroupManagerAction() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteInstanceGroupManager request example`, func() {
			// begin-delete_instance_group_manager

			options := &vpcv1.DeleteInstanceGroupManagerOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			options.SetID(instanceGroupManagerID)
			response, err := vpcService.DeleteInstanceGroupManager(options)

			// end-delete_instance_group_manager
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteInstanceGroupManager() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteInstanceGroupLoadBalancer request example`, func() {
			// begin-delete_instance_group_load_balancer

			options := &vpcv1.DeleteInstanceGroupLoadBalancerOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			response, err := vpcService.DeleteInstanceGroupLoadBalancer(options)

			// end-delete_instance_group_load_balancer
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteInstanceGroupLoadBalancer() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteInstanceGroup request example`, func() {
			// begin-delete_instance_group

			options := &vpcv1.DeleteInstanceGroupOptions{}
			options.SetID(instanceGroupID)
			response, err := vpcService.DeleteInstanceGroup(options)

			// end-delete_instance_group
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteInstanceGroup() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})

		It(`DeleteDedicatedHostGroup request example`, func() {
			// begin-delete_dedicated_host_group

			options := vpcService.NewDeleteDedicatedHostGroupOptions(
				dedicatedHostGroupID,
			)
			response, err := vpcService.DeleteDedicatedHostGroup(options)

			// end-delete_dedicated_host_group
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteDedicatedHostGroup() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteDedicatedHost request example`, func() {
			// begin-delete_dedicated_host

			options := vpcService.NewDeleteDedicatedHostOptions(dedicatedHostID)
			response, err := vpcService.DeleteDedicatedHost(options)

			// end-delete_dedicated_host
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteDedicatedHost() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteIkePolicy request example`, func() {
			// begin-delete_ike_policy

			options := vpcService.NewDeleteIkePolicyOptions(ikePolicyID)
			response, err := vpcService.DeleteIkePolicy(options)

			// end-delete_ike_policy
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteIkePolicy() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteIpsecPolicy request example`, func() {
			// begin-delete_ipsec_policy

			options := vpcService.NewDeleteIpsecPolicyOptions(ipsecPolicyID)
			response, err := vpcService.DeleteIpsecPolicy(options)

			// end-delete_ipsec_policy
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteIpsecPolicy() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteVPNServerRoute request example`, func() {
			// begin-delete_vpn_server_route

			deleteVPNServerRouteOptions := vpcService.NewDeleteVPNServerRouteOptions(
				vpnServerID,
				vpnServerRouteID,
			)

			response, err := vpcService.DeleteVPNServerRoute(deleteVPNServerRouteOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 202 {
				fmt.Printf("\nUnexpected response status code received from DeleteVPNServerRoute(): %d\n", response.StatusCode)
			}

			// end-delete_vpn_server_route

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})
		It(`DeleteVPNServerClient request example`, func() {
			// begin-delete_vpn_server_client

			deleteVPNServerClientOptions := vpcService.NewDeleteVPNServerClientOptions(
				vpnServerID,
				vpnClientID,
			)

			response, err := vpcService.DeleteVPNServerClient(deleteVPNServerClientOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 202 {
				fmt.Printf("\nUnexpected response status code received from DeleteVPNServerClient(): %d\n", response.StatusCode)
			}

			// end-delete_vpn_server_client

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})
		It(`DeleteVPNServer request example`, func() {
			// begin-delete_vpn_server

			deleteVPNServerOptions := vpcService.NewDeleteVPNServerOptions(
				vpnServerID,
			)
			deleteVPNServerOptions.SetIfMatch(ifMatchVPNServer)

			response, err := vpcService.DeleteVPNServer(deleteVPNServerOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 202 {
				fmt.Printf("\nUnexpected response status code received from DeleteVPNServer(): %d\n", response.StatusCode)
			}

			// end-delete_vpn_server

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})

		It(`DeleteVPNGatewayConnection request example`, func() {
			// begin-delete_vpn_gateway_connection

			options := &vpcv1.DeleteVPNGatewayConnectionOptions{}
			options.SetVPNGatewayID(vpnGatewayID)
			options.SetID(vpnGatewayConnectionID)
			response, err := vpcService.DeleteVPNGatewayConnection(options)

			// end-delete_vpn_gateway_connection
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteVPNGatewayConnection() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteVPNGateway request example`, func() {
			// begin-delete_vpn_gateway

			options := vpcService.NewDeleteVPNGatewayOptions(vpnGatewayID)
			response, err := vpcService.DeleteVPNGateway(options)

			// end-delete_vpn_gateway
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteVPNGateway() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})

		It(`DeleteBackupPolicyPlan request example`, func() {
			fmt.Println("\nDeleteBackupPolicyPlan() result:")
			// begin-delete_backup_policy_plan

			deleteBackupPolicyPlanOptions := vpcService.NewDeleteBackupPolicyPlanOptions(
				backupPolicyID,
				backupPolicyPlanID,
			)
			deleteBackupPolicyPlanOptions.SetIfMatch(ifMatchBackupPolicy)

			backupPolicyPlan, response, err := vpcService.DeleteBackupPolicyPlan(deleteBackupPolicyPlanOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_backup_policy_plan

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(backupPolicyPlan).ToNot(BeNil())

		})
		It(`DeleteBackupPolicy request example`, func() {
			fmt.Println("\nDeleteBackupPolicy() result:")
			// begin-delete_backup_policy

			deleteBackupPolicyOptions := vpcService.NewDeleteBackupPolicyOptions(
				backupPolicyID,
			)
			deleteBackupPolicyOptions.SetIfMatch(ifMatchBackupPolicy)

			backupPolicy, response, err := vpcService.DeleteBackupPolicy(deleteBackupPolicyOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_backup_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(backupPolicy).ToNot(BeNil())

		})
	})

})
