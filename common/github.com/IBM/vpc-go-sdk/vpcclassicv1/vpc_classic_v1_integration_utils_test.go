// +build integration

/**
 * (C) Copyright IBM Corp. 2020.
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
package vpcclassicv1_test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/vpc-go-sdk/vpcclassicv1"
)

/**
 * REST methods
 *
 */
const (
	POST   = http.MethodPost
	GET    = http.MethodGet
	DELETE = http.MethodDelete
	PUT    = http.MethodPut
	PATCH  = http.MethodPatch
)

// InstantiateVPCService - Instantiate VPC service
func InstantiateVPCService() *vpcclassicv1.VpcClassicV1 {
	service, serviceErr := vpcclassicv1.NewVpcClassicV1UsingExternalConfig(
		&vpcclassicv1.VpcClassicV1Options{},
	)
	// Check successful instantiation
	if serviceErr != nil {
		fmt.Println("Service creation failed. Error - ", serviceErr)
		return nil
	}
	// return new vpc service
	return service
}

/**
 * Regions and Zones
 *
 */

// ListRegions - List all regions
// GET
// /regions
func ListRegions(vpcService *vpcclassicv1.VpcClassicV1) (regions *vpcclassicv1.RegionCollection, response *core.DetailedResponse, err error) {
	listRegionsOptions := &vpcclassicv1.ListRegionsOptions{}
	regions, response, err = vpcService.ListRegions(listRegionsOptions)
	return
}

// GetRegion - GET
// /regions/{name}
// Retrieve a region
func GetRegion(vpcService *vpcclassicv1.VpcClassicV1, name string) (region *vpcclassicv1.Region, response *core.DetailedResponse, err error) {
	getRegionOptions := &vpcclassicv1.GetRegionOptions{}
	getRegionOptions.SetName(name)
	region, response, err = vpcService.GetRegion(getRegionOptions)
	return
}

// ListZones - GET
// /regions/{region_name}/zones
// List all zones in a region
func ListRegionZones(vpcService *vpcclassicv1.VpcClassicV1, regionName string) (zones *vpcclassicv1.ZoneCollection, response *core.DetailedResponse, err error) {
	listZonesOptions := &vpcclassicv1.ListRegionZonesOptions{}
	listZonesOptions.SetRegionName(regionName)
	zones, response, err = vpcService.ListRegionZones(listZonesOptions)
	return
}

// GetZone - GET
// /regions/{region_name}/zones/{zone_name}
// Retrieve a zone
func GetRegionZone(vpcService *vpcclassicv1.VpcClassicV1, regionName, zoneName string) (zone *vpcclassicv1.Zone, response *core.DetailedResponse, err error) {
	getZoneOptions := &vpcclassicv1.GetRegionZoneOptions{}
	getZoneOptions.SetRegionName(regionName)
	getZoneOptions.SetName(zoneName)
	zone, response, err = vpcService.GetRegionZone(getZoneOptions)
	return
}

/**
 * Floating IPs
 */

// GetFloatingIPsList - GET
// /floating_ips
// List all floating IPs
func ListFloatingIps(vpcService *vpcclassicv1.VpcClassicV1) (floatingIPs *vpcclassicv1.FloatingIPCollection, response *core.DetailedResponse, err error) {
	listFloatingIpsOptions := vpcService.NewListFloatingIpsOptions()
	floatingIPs, response, err = vpcService.ListFloatingIps(listFloatingIpsOptions)
	return
}

// GetFloatingIP - GET
// /floating_ips/{id}
// Retrieve the specified floating IP
func GetFloatingIp(vpcService *vpcclassicv1.VpcClassicV1, id string) (floatingIP *vpcclassicv1.FloatingIP, response *core.DetailedResponse, err error) {
	options := vpcService.NewGetFloatingIPOptions(id)
	floatingIP, response, err = vpcService.GetFloatingIP(options)
	return
}

// ReleaseFloatingIP - DELETE
// /floating_ips/{id}
// Release the specified floating IP
func DeleteFloatingIp(vpcService *vpcclassicv1.VpcClassicV1, id string) (response *core.DetailedResponse, err error) {
	options := vpcService.NewDeleteFloatingIPOptions(id)
	response, err = vpcService.DeleteFloatingIP(options)
	return
}

// UpdateFloatingIP - PATCH
// /floating_ips/{id}
// Update the specified floating IP
func UpdateFloatingIp(vpcService *vpcclassicv1.VpcClassicV1, id, name string) (floatingIP *vpcclassicv1.FloatingIP, response *core.DetailedResponse, err error) {
	body := &vpcclassicv1.FloatingIPPatch{
		Name: core.StringPtr(name),
	}
	patchBody, _ := body.AsPatch()
	options := &vpcclassicv1.UpdateFloatingIPOptions{
		ID:              core.StringPtr(id),
		FloatingIPPatch: patchBody,
	}
	floatingIP, response, err = vpcService.UpdateFloatingIP(options)
	return
}

// CreateFloatingIP - POST
// /floating_ips
// Reserve a floating IP
func CreateFloatingIp(vpcService *vpcclassicv1.VpcClassicV1, zone, name string) (floatingIP *vpcclassicv1.FloatingIP, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.CreateFloatingIPOptions{}
	options.SetFloatingIPPrototype(&vpcclassicv1.FloatingIPPrototype{
		Name: core.StringPtr(name),
		Zone: &vpcclassicv1.ZoneIdentity{
			Name: core.StringPtr(zone),
		}})
	floatingIP, response, err = vpcService.CreateFloatingIP(options)
	return
}

/**
 * SSH Keys
 *
 */

// ListKeys - GET
// /keys
// List all keys
func ListKeys(vpcService *vpcclassicv1.VpcClassicV1) (keys *vpcclassicv1.KeyCollection, response *core.DetailedResponse, err error) {
	listKeysOptions := &vpcclassicv1.ListKeysOptions{}
	keys, response, err = vpcService.ListKeys(listKeysOptions)
	return
}

// GetSSHKey - GET
// /keys/{id}
// Retrieve specified key
func GetKey(vpcService *vpcclassicv1.VpcClassicV1, id string) (key *vpcclassicv1.Key, response *core.DetailedResponse, err error) {
	getKeyOptions := &vpcclassicv1.GetKeyOptions{}
	getKeyOptions.SetID(id)
	key, response, err = vpcService.GetKey(getKeyOptions)
	return
}

// UpdateSSHKey - PATCH
// /keys/{id}
// Update specified key
func UpdateKey(vpcService *vpcclassicv1.VpcClassicV1, id, name string) (key *vpcclassicv1.Key, response *core.DetailedResponse, err error) {
	body := &vpcclassicv1.FloatingIPPatch{
		Name: core.StringPtr(name),
	}
	patchBody, _ := body.AsPatch()
	updateKeyOptions := &vpcclassicv1.UpdateKeyOptions{}
	updateKeyOptions.SetID(id)
	updateKeyOptions.SetKeyPatch(patchBody)
	key, response, err = vpcService.UpdateKey(updateKeyOptions)
	return
}

// DeleteSSHKey - DELETE
// /keys/{id}
// Delete specified key
func DeleteKey(vpcService *vpcclassicv1.VpcClassicV1, id string) (response *core.DetailedResponse, err error) {
	deleteKeyOptions := &vpcclassicv1.DeleteKeyOptions{}
	deleteKeyOptions.SetID(id)
	response, err = vpcService.DeleteKey(deleteKeyOptions)
	return
}

// CreateSSHKey - POST
// /keys
// Create a key
func CreateKey(vpcService *vpcclassicv1.VpcClassicV1, name, publicKey string) (key *vpcclassicv1.Key, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.CreateKeyOptions{}
	options.SetName(name)
	options.SetPublicKey(publicKey)
	key, response, err = vpcService.CreateKey(options)
	return
}

/**
 * VPC
 *
 */

// GetVPCsList - GET
// /vpcs
// List all VPCs
func ListVpcs(vpcService *vpcclassicv1.VpcClassicV1) (vpcs *vpcclassicv1.VPCCollection, response *core.DetailedResponse, err error) {
	listVpcsOptions := &vpcclassicv1.ListVpcsOptions{}
	vpcs, response, err = vpcService.ListVpcs(listVpcsOptions)
	return
}

// GetVPC - GET
// /vpcs/{id}
// Retrieve specified VPC
func GetVPC(vpcService *vpcclassicv1.VpcClassicV1, id string) (vpc *vpcclassicv1.VPC, response *core.DetailedResponse, err error) {
	getVpcOptions := &vpcclassicv1.GetVPCOptions{}
	getVpcOptions.SetID(id)
	vpc, response, err = vpcService.GetVPC(getVpcOptions)
	return
}

// DeleteVPC - DELETE
// /vpcs/{id}
// Delete specified VPC
func DeleteVPC(vpcService *vpcclassicv1.VpcClassicV1, id string) (response *core.DetailedResponse, err error) {
	deleteVpcOptions := &vpcclassicv1.DeleteVPCOptions{}
	deleteVpcOptions.SetID(id)
	response, err = vpcService.DeleteVPC(deleteVpcOptions)
	return
}

// UpdateVPC - PATCH
// /vpcs/{id}
// Update specified VPC
func UpdateVPC(vpcService *vpcclassicv1.VpcClassicV1, id, name string) (vpc *vpcclassicv1.VPC, response *core.DetailedResponse, err error) {
	body := &vpcclassicv1.VPCPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcclassicv1.UpdateVPCOptions{
		VPCPatch: patchBody,
		ID:       &id,
	}
	vpc, response, err = vpcService.UpdateVPC(options)
	return
}

// CreateVPC - POST
// /vpcs
// Create a VPC
func CreateVPC(vpcService *vpcclassicv1.VpcClassicV1, name, resourceGroup string) (vpc *vpcclassicv1.VPC, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.CreateVPCOptions{}
	options.SetResourceGroup(&vpcclassicv1.ResourceGroupIdentity{
		ID: core.StringPtr(resourceGroup),
	})
	options.SetName(name)
	vpc, response, err = vpcService.CreateVPC(options)
	return
}

/**
 * VPC default Security group
 * Getting default security group for a vpc with id
 */

// GetVPCDefaultSecurityGroup - GET
// /vpcs/{id}/default_security_group
// Retrieve a VPC's default security group
func GetVPCDefaultSecurityGroup(vpcService *vpcclassicv1.VpcClassicV1, id string) (sg *vpcclassicv1.DefaultSecurityGroup, response *core.DetailedResponse, err error) {
	getVpcDefaultSecurityGroupOptions := &vpcclassicv1.GetVPCDefaultSecurityGroupOptions{}
	getVpcDefaultSecurityGroupOptions.SetID(id)
	sg, response, err = vpcService.GetVPCDefaultSecurityGroup(getVpcDefaultSecurityGroupOptions)
	return
}

/**
 * VPC address prefix
 *
 */

// ListVpcAddressPrefixes - GET
// /vpcs/{vpc_id}/address_prefixes
// List all address pool prefixes for a VPC
func ListVpcAddressPrefixes(vpcService *vpcclassicv1.VpcClassicV1, vpcID string) (addrPrefixes *vpcclassicv1.AddressPrefixCollection, response *core.DetailedResponse, err error) {
	listVpcAddressPrefixesOptions := &vpcclassicv1.ListVPCAddressPrefixesOptions{}
	listVpcAddressPrefixesOptions.SetVPCID(vpcID)
	addrPrefixes, response, err = vpcService.ListVPCAddressPrefixes(listVpcAddressPrefixesOptions)
	return
}

// GetVpcAddressPrefix - GET
// /vpcs/{vpc_id}/address_prefixes/{id}
// Retrieve specified address pool prefix
func GetVpcAddressPrefix(vpcService *vpcclassicv1.VpcClassicV1, vpcID, addressPrefixID string) (addrPrefix *vpcclassicv1.AddressPrefix, response *core.DetailedResponse, err error) {
	getVpcAddressPrefixOptions := &vpcclassicv1.GetVPCAddressPrefixOptions{}
	getVpcAddressPrefixOptions.SetVPCID(vpcID)
	getVpcAddressPrefixOptions.SetID(addressPrefixID)
	addrPrefix, response, err = vpcService.GetVPCAddressPrefix(getVpcAddressPrefixOptions)
	return
}

// CreateVpcAddressPrefix - POST
// /vpcs/{vpc_id}/address_prefixes
// Create an address pool prefix
func CreateVpcAddressPrefix(vpcService *vpcclassicv1.VpcClassicV1, vpcID, zone, cidr, name string) (addrPrefix *vpcclassicv1.AddressPrefix, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.CreateVPCAddressPrefixOptions{}
	options.SetVPCID(vpcID)
	options.SetCIDR(cidr)
	options.SetName(name)
	options.SetZone(&vpcclassicv1.ZoneIdentity{
		Name: core.StringPtr(zone),
	})
	addrPrefix, response, err = vpcService.CreateVPCAddressPrefix(options)
	return
}

// DeleteVpcAddressPrefix - DELETE
// /vpcs/{vpc_id}/address_prefixes/{id}
// Delete specified address pool prefix
func DeleteVpcAddressPrefix(vpcService *vpcclassicv1.VpcClassicV1, vpcID, addressPrefixID string) (response *core.DetailedResponse, err error) {
	deleteVpcAddressPrefixOptions := &vpcclassicv1.DeleteVPCAddressPrefixOptions{}
	deleteVpcAddressPrefixOptions.SetVPCID(vpcID)
	deleteVpcAddressPrefixOptions.SetID(addressPrefixID)
	response, err = vpcService.DeleteVPCAddressPrefix(deleteVpcAddressPrefixOptions)
	return
}

// UpdateVpcAddressPrefix - PATCH
// /vpcs/{vpc_id}/address_prefixes/{id}
// Update an address pool prefix
func UpdateVpcAddressPrefix(vpcService *vpcclassicv1.VpcClassicV1, vpcID, addressPrefixID, name string) (addrPrefix *vpcclassicv1.AddressPrefix, response *core.DetailedResponse, err error) {
	body := &vpcclassicv1.AddressPrefixPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcclassicv1.UpdateVPCAddressPrefixOptions{
		AddressPrefixPatch: patchBody,
		VPCID:              &vpcID,
		ID:                 &addressPrefixID,
	}
	addrPrefix, response, err = vpcService.UpdateVPCAddressPrefix(options)
	return
}

/**
 * VPC routes
 *
 */

// ListVpcRoutes - GET
// /vpcs/{vpc_id}/routes
// List all user-defined routes for a VPC
func ListVpcRoutes(vpcService *vpcclassicv1.VpcClassicV1, vpcID string) (routes *vpcclassicv1.RouteCollection, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.ListVPCRoutesOptions{}
	options.SetVPCID(vpcID)
	routes, response, err = vpcService.ListVPCRoutes(options)
	return
}

// GetVpcRoute - GET
// /vpcs/{vpc_id}/routes/{id}
// Retrieve the specified route
func GetVpcRoute(vpcService *vpcclassicv1.VpcClassicV1, vpcID, routeID string) (route *vpcclassicv1.Route, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.GetVPCRouteOptions{}
	options.SetVPCID(vpcID)
	options.SetID(routeID)
	route, response, err = vpcService.GetVPCRoute(options)
	return
}

// CreateVpcRoute - POST
// /vpcs/{vpc_id}/routes
// Create a route on your VPC
func CreateVpcRoute(vpcService *vpcclassicv1.VpcClassicV1, vpcID, zone, destination, nextHopAddress, name string) (route *vpcclassicv1.Route, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.CreateVPCRouteOptions{}
	options.SetVPCID(vpcID)
	options.SetName(name)
	options.SetZone(&vpcclassicv1.ZoneIdentity{
		Name: core.StringPtr(zone),
	})
	options.SetNextHop(&vpcclassicv1.RouteNextHopPrototype{
		Address: &nextHopAddress,
	})
	options.SetDestination(destination)
	route, response, err = vpcService.CreateVPCRoute(options)
	return
}

// DeleteVpcRoute - DELETE
// /vpcs/{vpc_id}/routes/{id}
// Delete the specified route
func DeleteVpcRoute(vpcService *vpcclassicv1.VpcClassicV1, vpcID, routeID string) (response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.DeleteVPCRouteOptions{}
	options.SetVPCID(vpcID)
	options.SetID(routeID)
	response, err = vpcService.DeleteVPCRoute(options)
	return
}

// UpdateVpcRoute - PATCH
// /vpcs/{vpc_id}/routes/{id}
// Update a route
func UpdateVpcRoute(vpcService *vpcclassicv1.VpcClassicV1, vpcID, routeID, name string) (route *vpcclassicv1.Route, response *core.DetailedResponse, err error) {
	body := &vpcclassicv1.RoutePatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcclassicv1.UpdateVPCRouteOptions{}
	options.SetVPCID(vpcID)
	options.SetID(routeID)
	options.SetRoutePatch(patchBody)
	route, response, err = vpcService.UpdateVPCRoute(options)
	return
}

/**
 * Volumes
 *
 */

// ListVolumeProfiles - GET
// /volume/profiles
// List all volume profiles
func ListVolumeProfiles(vpcService *vpcclassicv1.VpcClassicV1) (profiles *vpcclassicv1.VolumeProfileCollection, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.ListVolumeProfilesOptions{}
	profiles, response, err = vpcService.ListVolumeProfiles(options)
	return
}

// GetVolumeProfile - GET
// /volume/profiles/{name}
// Retrieve specified volume profile
func GetVolumeProfile(vpcService *vpcclassicv1.VpcClassicV1, profileName string) (profile *vpcclassicv1.VolumeProfile, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.GetVolumeProfileOptions{}
	options.SetName(profileName)
	profile, response, err = vpcService.GetVolumeProfile(options)
	return
}

// ListVolumes - GET
// /volumes
// List all volumes
func ListVolumes(vpcService *vpcclassicv1.VpcClassicV1) (volumes *vpcclassicv1.VolumeCollection, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.ListVolumesOptions{}
	volumes, response, err = vpcService.ListVolumes(options)
	return
}

// GetVolume - GET
// /volumes/{id}
// Retrieve specified volume
func GetVolume(vpcService *vpcclassicv1.VpcClassicV1, volumeID string) (volume *vpcclassicv1.Volume, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.GetVolumeOptions{}
	options.SetID(volumeID)
	volume, response, err = vpcService.GetVolume(options)
	return
}

// DeleteVolume - DELETE
// /volumes/{id}
// Delete specified volume
func DeleteVolume(vpcService *vpcclassicv1.VpcClassicV1, id string) (response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.DeleteVolumeOptions{}
	options.SetID(id)
	response, err = vpcService.DeleteVolume(options)
	return
}

// UpdateVolume - PATCH
// /volumes/{id}
// Update specified volume
func UpdateVolume(vpcService *vpcclassicv1.VpcClassicV1, id, name string) (volume *vpcclassicv1.Volume, response *core.DetailedResponse, err error) {
	body := &vpcclassicv1.VolumePatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcclassicv1.UpdateVolumeOptions{
		VolumePatch: patchBody,
	}
	options.SetID(id)
	volume, response, err = vpcService.UpdateVolume(options)
	return
}

// CreateVolume - POST
// /volumes
// Create a volume
func CreateVolume(vpcService *vpcclassicv1.VpcClassicV1, name, profileName, zoneName string, capacity int64) (volume *vpcclassicv1.Volume, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.CreateVolumeOptions{}
	options.SetVolumePrototype(&vpcclassicv1.VolumePrototype{
		Capacity: core.Int64Ptr(capacity),
		Zone: &vpcclassicv1.ZoneIdentity{
			Name: core.StringPtr(zoneName),
		},
		Profile: &vpcclassicv1.VolumeProfileIdentity{
			Name: core.StringPtr(profileName),
		},
		Name: core.StringPtr(name),
	})
	volume, response, err = vpcService.CreateVolume(options)
	return
}

/**
 * Subnets
 *
 */

// ListSubnets - GET
// /subnets
// List all subnets
func ListSubnets(vpcService *vpcclassicv1.VpcClassicV1) (subnets *vpcclassicv1.SubnetCollection, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.ListSubnetsOptions{}
	subnets, response, err = vpcService.ListSubnets(options)
	return
}

// GetSubnet - GET
// /subnets/{id}
// Retrieve specified subnet
func GetSubnet(vpcService *vpcclassicv1.VpcClassicV1, subnetID string) (subnet *vpcclassicv1.Subnet, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.GetSubnetOptions{}
	options.SetID(subnetID)
	subnet, response, err = vpcService.GetSubnet(options)
	return
}

// DeleteSubnet - DELETE
// /subnets/{id}
// Delete specified subnet
func DeleteSubnet(vpcService *vpcclassicv1.VpcClassicV1, id string) (response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.DeleteSubnetOptions{}
	options.SetID(id)
	response, err = vpcService.DeleteSubnet(options)
	return
}

// UpdateSubnet - PATCH
// /subnets/{id}
// Update specified subnet
func UpdateSubnet(vpcService *vpcclassicv1.VpcClassicV1, id, name string) (subnet *vpcclassicv1.Subnet, response *core.DetailedResponse, err error) {
	body := &vpcclassicv1.SubnetPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcclassicv1.UpdateSubnetOptions{}
	options.SetID(id)
	options.SetSubnetPatch(patchBody)
	subnet, response, err = vpcService.UpdateSubnet(options)
	return
}

// CreateSubnet - POST
// /subnets
// Create a subnet
func CreateSubnet(vpcService *vpcclassicv1.VpcClassicV1, vpcID, name, zone string, mock bool) (subnet *vpcclassicv1.Subnet, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.CreateSubnetOptions{}
	if mock {
		options.SetSubnetPrototype(&vpcclassicv1.SubnetPrototype{
			Ipv4CIDRBlock: core.StringPtr("10.243.0.0/24"),
			Name:          core.StringPtr(name),
			VPC: &vpcclassicv1.VPCIdentity{
				ID: core.StringPtr(vpcID),
			},
			Zone: &vpcclassicv1.ZoneIdentity{
				Name: core.StringPtr(zone),
			},
		})
	} else {
		options.SetSubnetPrototype(&vpcclassicv1.SubnetPrototype{
			Name: core.StringPtr(name),
			VPC: &vpcclassicv1.VPCIdentity{
				ID: core.StringPtr(vpcID),
			},
			Zone: &vpcclassicv1.ZoneIdentity{
				Name: core.StringPtr(zone),
			},
			TotalIpv4AddressCount: core.Int64Ptr(128),
		})
	}
	subnet, response, err = vpcService.CreateSubnet(options)
	return
}

// GetSubnetNetworkAcl -GET
// /subnets/{id}/network_acl
// Retrieve a subnet's attached network ACL
func GetSubnetNetworkAcl(vpcService *vpcclassicv1.VpcClassicV1, subnetID string) (nacl *vpcclassicv1.NetworkACL, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.GetSubnetNetworkACLOptions{}
	options.SetID(subnetID)
	nacl, response, err = vpcService.GetSubnetNetworkACL(options)
	return
}

// SetSubnetNetworkAclBinding - PUT
// /subnets/{id}/network_acl
// Attach a network ACL to a subnet
func SetSubnetNetworkAclBinding(vpcService *vpcclassicv1.VpcClassicV1, subnetID, id string) (nacl *vpcclassicv1.NetworkACL, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.ReplaceSubnetNetworkACLOptions{}
	options.SetID(subnetID)
	options.SetNetworkACLIdentity(&vpcclassicv1.NetworkACLIdentity{ID: &id})
	nacl, response, err = vpcService.ReplaceSubnetNetworkACL(options)
	return
}

// DeleteSubnetPublicGatewayBinding - DELETE
// /subnets/{id}/public_gateway
// Detach a public gateway from a subnet
func DeleteSubnetPublicGatewayBinding(vpcService *vpcclassicv1.VpcClassicV1, id string) (response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.UnsetSubnetPublicGatewayOptions{}
	options.SetID(id)
	response, err = vpcService.UnsetSubnetPublicGateway(options)
	return
}

// GetSubnetPublicGateway - GET
// /subnets/{id}/public_gateway
// Retrieve a subnet's attached public gateway
func GetSubnetPublicGateway(vpcService *vpcclassicv1.VpcClassicV1, subnetID string) (pgw *vpcclassicv1.PublicGateway, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.GetSubnetPublicGatewayOptions{}
	options.SetID(subnetID)
	pgw, response, err = vpcService.GetSubnetPublicGateway(options)
	return
}

// CreateSubnetPublicGatewayBindingOptions - PUT
// /subnets/{id}/public_gateway
// Attach a public gateway to a subnet
func CreateSubnetPublicGatewayBinding(vpcService *vpcclassicv1.VpcClassicV1, subnetID, id string) (pgw *vpcclassicv1.PublicGateway, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.SetSubnetPublicGatewayOptions{}
	options.SetID(subnetID)
	options.SetPublicGatewayIdentity(&vpcclassicv1.PublicGatewayIdentity{ID: &id})
	pgw, response, err = vpcService.SetSubnetPublicGateway(options)
	return
}

/**
 * Images
 *
 */

// ListImages - GET
// /images
// List all images
func ListImages(vpcService *vpcclassicv1.VpcClassicV1, visibility string) (images *vpcclassicv1.ImageCollection, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.ListImagesOptions{}
	options.SetVisibility(visibility)
	images, response, err = vpcService.ListImages(options)
	return
}

// GetImage - GET
// /images/{id}
// Retrieve the specified image
func GetImage(vpcService *vpcclassicv1.VpcClassicV1, imageID string) (image *vpcclassicv1.Image, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.GetImageOptions{}
	options.SetID(imageID)
	image, response, err = vpcService.GetImage(options)
	return
}

// DeleteImage DELETE
// /images/{id}
// Delete specified image
func DeleteImage(vpcService *vpcclassicv1.VpcClassicV1, id string) (response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.DeleteImageOptions{}
	options.SetID(id)
	response, err = vpcService.DeleteImage(options)
	return
}

// UpdateImage PATCH
// /images/{id}
// Update specified image
func UpdateImage(vpcService *vpcclassicv1.VpcClassicV1, id, name string) (image *vpcclassicv1.Image, response *core.DetailedResponse, err error) {
	body := &vpcclassicv1.ImagePatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcclassicv1.UpdateImageOptions{}
	options.SetID(id)
	options.SetImagePatch(patchBody)
	image, response, err = vpcService.UpdateImage(options)
	return
}

func CreateImage(vpcService *vpcclassicv1.VpcClassicV1, name string) (image *vpcclassicv1.Image, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.CreateImageOptions{}
	cosID := "cos://cos-location-of-image-file"
	options.SetImagePrototype(&vpcclassicv1.ImagePrototype{
		File: &vpcclassicv1.ImageFilePrototype{
			Href: &cosID,
		},
		Name: &name,
	})
	image, response, err = vpcService.CreateImage(options)
	return
}

func ListOperatingSystems(vpcService *vpcclassicv1.VpcClassicV1) (operatingSystems *vpcclassicv1.OperatingSystemCollection, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.ListOperatingSystemsOptions{}
	operatingSystems, response, err = vpcService.ListOperatingSystems(options)
	return
}

func GetOperatingSystem(vpcService *vpcclassicv1.VpcClassicV1, osName string) (os *vpcclassicv1.OperatingSystem, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.GetOperatingSystemOptions{}
	options.SetName(osName)
	os, response, err = vpcService.GetOperatingSystem(options)
	return
}

/**
 * Instances
 *
 */

// ListInstanceProfiles - GET
// /instance/profiles
// List all instance profiles
func ListInstanceProfiles(vpcService *vpcclassicv1.VpcClassicV1) (profiles *vpcclassicv1.InstanceProfileCollection, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.ListInstanceProfilesOptions{}
	profiles, response, err = vpcService.ListInstanceProfiles(options)
	return
}

// GetInstanceProfile - GET
// /instance/profiles/{name}
// Retrieve specified instance profile
func GetInstanceProfile(vpcService *vpcclassicv1.VpcClassicV1, profileName string) (profile *vpcclassicv1.InstanceProfile, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.GetInstanceProfileOptions{}
	options.SetName(profileName)
	profile, response, err = vpcService.GetInstanceProfile(options)
	return
}

// ListInstances GET
// /instances
// List all instances
func ListInstances(vpcService *vpcclassicv1.VpcClassicV1) (instances *vpcclassicv1.InstanceCollection, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.ListInstancesOptions{}
	instances, response, err = vpcService.ListInstances(options)
	return
}

// GetInstance GET
// instances/{id}
// Retrieve an instance
func GetInstance(vpcService *vpcclassicv1.VpcClassicV1, instanceID string) (instance *vpcclassicv1.Instance, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.GetInstanceOptions{}
	options.SetID(instanceID)
	instance, response, err = vpcService.GetInstance(options)
	return
}

// DeleteInstance DELETE
// /instances/{id}
// Delete specified instance
func DeleteInstance(vpcService *vpcclassicv1.VpcClassicV1, id string) (response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.DeleteInstanceOptions{}
	options.SetID(id)
	response, err = vpcService.DeleteInstance(options)
	return
}

// UpdateInstance PATCH
// /instances/{id}
// Update specified instance
func UpdateInstance(vpcService *vpcclassicv1.VpcClassicV1, id, name string) (instance *vpcclassicv1.Instance, response *core.DetailedResponse, err error) {
	body := &vpcclassicv1.InstancePatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcclassicv1.UpdateInstanceOptions{}
	options.SetID(id)
	options.SetInstancePatch(patchBody)
	instance, response, err = vpcService.UpdateInstance(options)
	return
}

// CreateInstance POST
// /instances
// Create an instance action
func CreateInstance(vpcService *vpcclassicv1.VpcClassicV1, name, profileName, imageID, zoneName, subnetID, sshkeyID, vpcID string) (instance *vpcclassicv1.Instance, response *core.DetailedResponse, err error) {
	key := &vpcclassicv1.KeyIdentity{
		ID: core.StringPtr(sshkeyID),
	}
	options := &vpcclassicv1.CreateInstanceOptions{}
	options.SetInstancePrototype(&vpcclassicv1.InstancePrototype{
		Zone: &vpcclassicv1.ZoneIdentity{
			Name: core.StringPtr(zoneName),
		},
		Image: &vpcclassicv1.ImageIdentity{
			ID: core.StringPtr(imageID),
		},
		Profile: &vpcclassicv1.InstanceProfileIdentity{
			Name: core.StringPtr(profileName),
		},
		Name: &name,
		PrimaryNetworkInterface: &vpcclassicv1.NetworkInterfacePrototype{
			Subnet: &vpcclassicv1.SubnetIdentity{
				ID: core.StringPtr(subnetID),
			},
		},
		Keys: []vpcclassicv1.KeyIdentityIntf{key},
		VPC: &vpcclassicv1.VPCIdentity{
			ID: core.StringPtr(vpcID),
		},
	})
	instance, response, err = vpcService.CreateInstance(options)
	return
}

// CreateInstanceAction POST
// /instances/{id}
// Create an instance action
func CreateInstanceAction(vpcService *vpcclassicv1.VpcClassicV1, instanceID, typeOfAction string) (instance *vpcclassicv1.InstanceAction, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.CreateInstanceActionOptions{}
	options.SetInstanceID(instanceID)
	options.SetType(typeOfAction)
	instance, response, err = vpcService.CreateInstanceAction(options)
	return
}

// GetInstanceInitialization GET
// /instances/{id}/initialization
// Retrieve configuration used to initialize the instance.
func GetInstanceInitialization(vpcService *vpcclassicv1.VpcClassicV1, instanceID string) (instance *vpcclassicv1.InstanceInitialization, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.GetInstanceInitializationOptions{}
	options.SetID(instanceID)
	instance, response, err = vpcService.GetInstanceInitialization(options)
	return
}

// ListNetworkInterfaces GET
// /instances/{instance_id}/network_interfaces
// List all network interfaces on an instance
func ListNetworkInterfaces(vpcService *vpcclassicv1.VpcClassicV1, id string) (netInterfaces *vpcclassicv1.NetworkInterfaceUnpaginatedCollection, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.ListInstanceNetworkInterfacesOptions{}
	options.SetInstanceID(id)
	netInterfaces, response, err = vpcService.ListInstanceNetworkInterfaces(options)
	return
}

// GetNetworkInterface GET
// /instances/{instance_id}/network_interfaces/{id}
// Retrieve specified network interface
func GetNetworkInterface(vpcService *vpcclassicv1.VpcClassicV1, instanceID, networkID string) (netInterface *vpcclassicv1.NetworkInterface, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.GetInstanceNetworkInterfaceOptions{}
	options.SetID(networkID)
	options.SetInstanceID(instanceID)
	netInterface, response, err = vpcService.GetInstanceNetworkInterface(options)
	return
}

// ListNetworkInterfaceFloatingIps GET
// /instances/{instance_id}/network_interfaces
// List all network interfaces on an instance
func ListNetworkInterfaceFloatingIps(vpcService *vpcclassicv1.VpcClassicV1, instanceID, networkID string) (netInterfaceFIPs *vpcclassicv1.FloatingIPUnpaginatedCollection, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.ListInstanceNetworkInterfaceFloatingIpsOptions{}
	options.SetInstanceID(instanceID)
	options.SetNetworkInterfaceID(networkID)
	netInterfaceFIPs, response, err = vpcService.ListInstanceNetworkInterfaceFloatingIps(options)
	return
}

// GetNetworkInterfaceFloatingIp GET
// /instances/{instance_id}/network_interfaces/{network_interface_id}/floating_ips
// List all floating IPs associated with a network interface
func GetNetworkInterfaceFloatingIp(vpcService *vpcclassicv1.VpcClassicV1, instanceID, networkID, fipID string) (netInterfaceFIP *vpcclassicv1.FloatingIP, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.GetInstanceNetworkInterfaceFloatingIPOptions{}
	options.SetID(fipID)
	options.SetInstanceID(instanceID)
	options.SetNetworkInterfaceID(networkID)
	netInterfaceFIP, response, err = vpcService.GetInstanceNetworkInterfaceFloatingIP(options)
	return
}

// DeleteNetworkInterfaceFloatingIpBinding DELETE
// /instances/{instance_id}/network_interfaces/{network_interface_id}/floating_ips/{id}
// Disassociate specified floating IP
func DeleteNetworkInterfaceFloatingIpBinding(vpcService *vpcclassicv1.VpcClassicV1, instanceID, networkID, fipID string) (response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.RemoveInstanceNetworkInterfaceFloatingIPOptions{}
	options.SetID(fipID)
	options.SetInstanceID(instanceID)
	options.SetNetworkInterfaceID(networkID)
	response, err = vpcService.RemoveInstanceNetworkInterfaceFloatingIP(options)
	return
}

// CreateNetworkInterfaceFloatingIpBinding PUT
// /instances/{instance_id}/network_interfaces/{network_interface_id}/floating_ips/{id}
// Associate a floating IP with a network interface
func CreateNetworkInterfaceFloatingIpBinding(vpcService *vpcclassicv1.VpcClassicV1, instanceID, networkID, fipID string) (netInterfaceFIP *vpcclassicv1.FloatingIP, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.AddInstanceNetworkInterfaceFloatingIPOptions{}
	options.SetID(fipID)
	options.SetInstanceID(instanceID)
	options.SetNetworkInterfaceID(networkID)
	netInterfaceFIP, response, err = vpcService.AddInstanceNetworkInterfaceFloatingIP(options)
	return
}

// ListVolumeAttachments GET
// /instances/{instance_id}/volume_attachments
// List all volumes attached to an instance
func ListVolumeAttachments(vpcService *vpcclassicv1.VpcClassicV1, id string) (volumes *vpcclassicv1.VolumeAttachmentCollection, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.ListInstanceVolumeAttachmentsOptions{}
	options.SetInstanceID(id)
	volumes, response, err = vpcService.ListInstanceVolumeAttachments(options)
	return
}

// CreateVolumeAttachment POST
// /instances/{instance_id}/volume_attachments
// Create a volume attachment, connecting a volume to an instance
func CreateVolumeAttachment(vpcService *vpcclassicv1.VpcClassicV1, instanceID, volumeID, name string) (volume *vpcclassicv1.VolumeAttachment, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.CreateInstanceVolumeAttachmentOptions{}
	options.SetInstanceID(instanceID)
	options.SetVolume(&vpcclassicv1.VolumeIdentity{
		ID: core.StringPtr(volumeID),
	})
	options.SetName(name)
	options.SetDeleteVolumeOnInstanceDelete(false)
	volume, response, err = vpcService.CreateInstanceVolumeAttachment(options)
	return
}

// DeleteVolumeAttachment DELETE
// /instances/{instance_id}/volume_attachments/{id}
// Delete a volume attachment, detaching a volume from an instance
func DeleteVolumeAttachment(vpcService *vpcclassicv1.VpcClassicV1, instanceID, volumeID string) (response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.DeleteInstanceVolumeAttachmentOptions{}
	options.SetID(volumeID)
	options.SetInstanceID(instanceID)
	response, err = vpcService.DeleteInstanceVolumeAttachment(options)
	return
}

// GetVolumeAttachment GET
// /instances/{instance_id}/volume_attachments/{id}
// Retrieve specified volume attachment
func GetVolumeAttachment(vpcService *vpcclassicv1.VpcClassicV1, instanceID, volumeID string) (volume *vpcclassicv1.VolumeAttachment, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.GetInstanceVolumeAttachmentOptions{}
	options.SetInstanceID(instanceID)
	options.SetID(volumeID)
	volume, response, err = vpcService.GetInstanceVolumeAttachment(options)
	return
}

// UpdateVolumeAttachment PATCH
// /instances/{instance_id}/volume_attachments/{id}
// Update a volume attachment
func UpdateVolumeAttachment(vpcService *vpcclassicv1.VpcClassicV1, instanceID, volumeID, name string) (volume *vpcclassicv1.VolumeAttachment, response *core.DetailedResponse, err error) {
	body := &vpcclassicv1.VolumeAttachmentPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcclassicv1.UpdateInstanceVolumeAttachmentOptions{}
	options.SetInstanceID(instanceID)
	options.SetID(volumeID)
	options.SetVolumeAttachmentPatch(patchBody)
	volume, response, err = vpcService.UpdateInstanceVolumeAttachment(options)
	return
}

/**
 * Public Gateway
 *
 */

// ListPublicGateways GET
// /public_gateways
// List all public gateways
func ListPublicGateways(vpcService *vpcclassicv1.VpcClassicV1) (pgws *vpcclassicv1.PublicGatewayCollection, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.ListPublicGatewaysOptions{}
	pgws, response, err = vpcService.ListPublicGateways(options)
	return
}

// CreatePublicGateway POST
// /public_gateways
// Create a public gateway
func CreatePublicGateway(vpcService *vpcclassicv1.VpcClassicV1, name, vpcID, zoneName string) (pgw *vpcclassicv1.PublicGateway, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.CreatePublicGatewayOptions{}
	options.SetVPC(&vpcclassicv1.VPCIdentity{
		ID: core.StringPtr(vpcID),
	})
	options.SetZone(&vpcclassicv1.ZoneIdentity{
		Name: core.StringPtr(zoneName),
	})
	pgw, response, err = vpcService.CreatePublicGateway(options)
	return
}

// DeletePublicGateway DELETE
// /public_gateways/{id}
// Delete specified public gateway
func DeletePublicGateway(vpcService *vpcclassicv1.VpcClassicV1, id string) (response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.DeletePublicGatewayOptions{}
	options.SetID(id)
	response, err = vpcService.DeletePublicGateway(options)
	return
}

// GetPublicGateway GET
// /public_gateways/{id}
// Retrieve specified public gateway
func GetPublicGateway(vpcService *vpcclassicv1.VpcClassicV1, id string) (pgw *vpcclassicv1.PublicGateway, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.GetPublicGatewayOptions{}
	options.SetID(id)
	pgw, response, err = vpcService.GetPublicGateway(options)
	return
}

// UpdatePublicGateway PATCH
// /public_gateways/{id}
// Update a public gateway's name
func UpdatePublicGateway(vpcService *vpcclassicv1.VpcClassicV1, id, name string) (pgw *vpcclassicv1.PublicGateway, response *core.DetailedResponse, err error) {
	body := &vpcclassicv1.PublicGatewayPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcclassicv1.UpdatePublicGatewayOptions{}
	options.SetID(id)
	options.SetPublicGatewayPatch(patchBody)
	pgw, response, err = vpcService.UpdatePublicGateway(options)
	return
}

/**
 * Network ACLs
 *
 */

// ListNetworkAcls - GET
// /network_acls
// List all network ACLs
func ListNetworkAcls(vpcService *vpcclassicv1.VpcClassicV1) (nacls *vpcclassicv1.NetworkACLCollection, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.ListNetworkAclsOptions{}
	nacls, response, err = vpcService.ListNetworkAcls(options)
	return
}

// CreateNetworkAcl - POST
// /network_acls
// Create a network ACL
func CreateNetworkAcl(vpcService *vpcclassicv1.VpcClassicV1, name, copyableAclID string) (nacl *vpcclassicv1.NetworkACL, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.CreateNetworkACLOptions{}
	options.SetNetworkACLPrototype(&vpcclassicv1.NetworkACLPrototype{
		Name: core.StringPtr(name),
		SourceNetworkACL: &vpcclassicv1.NetworkACLIdentity{
			ID: core.StringPtr(copyableAclID),
		},
	})
	nacl, response, err = vpcService.CreateNetworkACL(options)
	return
}

// DeleteNetworkAcl - DELETE
// /network_acls/{id}
// Delete specified network ACL
func DeleteNetworkAcl(vpcService *vpcclassicv1.VpcClassicV1, ID string) (response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.DeleteNetworkACLOptions{}
	options.SetID(ID)
	response, err = vpcService.DeleteNetworkACL(options)
	return
}

// GetNetworkAcl - GET
// /network_acls/{id}
// Retrieve specified network ACL
func GetNetworkAcl(vpcService *vpcclassicv1.VpcClassicV1, ID string) (nacl *vpcclassicv1.NetworkACL, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.GetNetworkACLOptions{}
	options.SetID(ID)
	nacl, response, err = vpcService.GetNetworkACL(options)
	return
}

// UpdateNetworkAcl PATCH
// /network_acls/{id}
// Update a network ACL
func UpdateNetworkAcl(vpcService *vpcclassicv1.VpcClassicV1, id, name string) (nacl *vpcclassicv1.NetworkACL, response *core.DetailedResponse, err error) {
	body := &vpcclassicv1.NetworkACLPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcclassicv1.UpdateNetworkACLOptions{}
	options.SetID(id)
	options.SetNetworkACLPatch(patchBody)
	nacl, response, err = vpcService.UpdateNetworkACL(options)
	return
}

// ListNetworkAclRules - GET
// /network_acls/{network_acl_id}/rules
// List all rules for a network ACL
func ListNetworkAclRules(vpcService *vpcclassicv1.VpcClassicV1, aclID string) (naclRules *vpcclassicv1.NetworkACLRuleCollection, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.ListNetworkACLRulesOptions{}
	options.SetNetworkACLID(aclID)
	naclRules, response, err = vpcService.ListNetworkACLRules(options)
	return
}

// CreateNetworkAclRule - POST
// /network_acls/{network_acl_id}/rules
// Create a rule
func CreateNetworkAclRule(vpcService *vpcclassicv1.VpcClassicV1, name, aclID string) (naclRule vpcclassicv1.NetworkACLRuleIntf, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.CreateNetworkACLRuleOptions{}
	options.SetNetworkACLID(aclID)
	options.SetNetworkACLRulePrototype(&vpcclassicv1.NetworkACLRulePrototype{
		Action:      core.StringPtr("allow"),
		Direction:   core.StringPtr("inbound"),
		Destination: core.StringPtr("0.0.0.0/0"),
		Source:      core.StringPtr("0.0.0.0/0"),
		Protocol:    core.StringPtr("all"),
		Name:        core.StringPtr(name),
	})
	naclRule, response, err = vpcService.CreateNetworkACLRule(options)
	return
}

// DeleteNetworkAclRule DELETE
// /network_acls/{network_acl_id}/rules/{id}
// Delete specified rule
func DeleteNetworkAclRule(vpcService *vpcclassicv1.VpcClassicV1, aclID, ruleID string) (response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.DeleteNetworkACLRuleOptions{}
	options.SetID(ruleID)
	options.SetNetworkACLID(aclID)
	response, err = vpcService.DeleteNetworkACLRule(options)
	return
}

// GetNetworkAclRule GET
// /network_acls/{network_acl_id}/rules/{id}
// Retrieve specified rule
func GetNetworkAclRule(vpcService *vpcclassicv1.VpcClassicV1, aclID, ruleID string) (naclRule vpcclassicv1.NetworkACLRuleIntf, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.GetNetworkACLRuleOptions{}
	options.SetID(ruleID)
	options.SetNetworkACLID(aclID)
	naclRule, response, err = vpcService.GetNetworkACLRule(options)
	return
}

// UpdateNetworkAclRule PATCH
// /network_acls/{network_acl_id}/rules/{id}
// Update a rule
func UpdateNetworkAclRule(vpcService *vpcclassicv1.VpcClassicV1, aclID, ruleID, name string) (naclRule vpcclassicv1.NetworkACLRuleIntf, response *core.DetailedResponse, err error) {
	body := &vpcclassicv1.NetworkACLRulePatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcclassicv1.UpdateNetworkACLRuleOptions{}
	options.SetID(ruleID)
	options.SetNetworkACLID(aclID)
	options.SetNetworkACLRulePatch(patchBody)
	naclRule, response, err = vpcService.UpdateNetworkACLRule(options)
	return
}

/**
 * Security Groups
 *
 */

// ListSecurityGroups GET
// /security_groups
// List all security groups
func ListSecurityGroups(vpcService *vpcclassicv1.VpcClassicV1) (sgs *vpcclassicv1.SecurityGroupCollection, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.ListSecurityGroupsOptions{}
	sgs, response, err = vpcService.ListSecurityGroups(options)
	return
}

// CreateSecurityGroup POST
// /security_groups
// Create a security group
func CreateSecurityGroup(vpcService *vpcclassicv1.VpcClassicV1, name, vpcID string) (sg *vpcclassicv1.SecurityGroup, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.CreateSecurityGroupOptions{}

	options.SetVPC(&vpcclassicv1.VPCIdentity{
		ID: core.StringPtr(vpcID),
	})
	options.SetName(name)
	sg, response, err = vpcService.CreateSecurityGroup(options)
	return
}

// DeleteSecurityGroup DELETE
// /security_groups/{id}
// Delete a security group
func DeleteSecurityGroup(vpcService *vpcclassicv1.VpcClassicV1, id string) (response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.DeleteSecurityGroupOptions{}

	options.SetID(id)
	response, err = vpcService.DeleteSecurityGroup(options)
	return
}

// GetSecurityGroup GET
// /security_groups/{id}
// Retrieve a security group
func GetSecurityGroup(vpcService *vpcclassicv1.VpcClassicV1, id string) (sg *vpcclassicv1.SecurityGroup, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.GetSecurityGroupOptions{}
	options.SetID(id)
	sg, response, err = vpcService.GetSecurityGroup(options)
	return
}

// UpdateSecurityGroup PATCH
// /security_groups/{id}
// Update a security group
func UpdateSecurityGroup(vpcService *vpcclassicv1.VpcClassicV1, id, name string) (sg *vpcclassicv1.SecurityGroup, response *core.DetailedResponse, err error) {
	body := &vpcclassicv1.SecurityGroupPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcclassicv1.UpdateSecurityGroupOptions{}
	options.SetID(id)
	options.SetSecurityGroupPatch(patchBody)
	sg, response, err = vpcService.UpdateSecurityGroup(options)
	return
}

// ListSecurityGroupNetworkInterfaces GET
// /security_groups/{security_group_id}/network_interfaces
// List a security group's network interfaces
// ListSecurityGroupNetworkInterfaces
func ListSecurityGroupNetworkInterfaces(vpcService *vpcclassicv1.VpcClassicV1, sgID string) (sgNwInterfaces *vpcclassicv1.NetworkInterfaceCollection, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.ListSecurityGroupNetworkInterfacesOptions{}
	options.SetSecurityGroupID(sgID)
	sgNwInterfaces, response, err = vpcService.ListSecurityGroupNetworkInterfaces(options)
	return
}

// DeleteSecurityGroupNetworkInterfaceBinding DELETE
// /security_groups/{security_group_id}/network_interfaces/{id}
// Remove a network interface from a security group.
func DeleteSecurityGroupNetworkInterfaceBinding(vpcService *vpcclassicv1.VpcClassicV1, id, vnicID string) (response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.RemoveSecurityGroupNetworkInterfaceOptions{}
	options.SetSecurityGroupID(id)
	options.SetID(vnicID)
	response, err = vpcService.RemoveSecurityGroupNetworkInterface(options)
	return
}

// GetSecurityGroupNetworkInterface GET
// /security_groups/{security_group_id}/network_interfaces/{id}
// Retrieve a network interface in a security group
func GetSecurityGroupNetworkInterface(vpcService *vpcclassicv1.VpcClassicV1, id, vnicID string) (sgNwInterface *vpcclassicv1.NetworkInterface, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.GetSecurityGroupNetworkInterfaceOptions{}
	options.SetSecurityGroupID(id)
	options.SetID(vnicID)
	sgNwInterface, response, err = vpcService.GetSecurityGroupNetworkInterface(options)
	return
}

// CreateSecurityGroupNetworkInterfaceBinding PUT
// /security_groups/{security_group_id}/network_interfaces/{id}
// Add a network interface to a security group
func CreateSecurityGroupNetworkInterfaceBinding(vpcService *vpcclassicv1.VpcClassicV1, id, vnicID string) (nic *vpcclassicv1.NetworkInterface, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.AddSecurityGroupNetworkInterfaceOptions{}
	options.SetSecurityGroupID(id)
	options.SetID(vnicID)
	nic, response, err = vpcService.AddSecurityGroupNetworkInterface(options)
	return
}

// ListSecurityGroupRules GET
// /security_groups/{security_group_id}/rules
// List all the rules of a security group
func ListSecurityGroupRules(vpcService *vpcclassicv1.VpcClassicV1, id string) (rules *vpcclassicv1.SecurityGroupRuleCollection, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.ListSecurityGroupRulesOptions{}
	options.SetSecurityGroupID(id)
	rules, response, err = vpcService.ListSecurityGroupRules(options)
	return
}

// CreateSecurityGroupRule POST
// /security_groups/{security_group_id}/rules
// Create a security group rule
func CreateSecurityGroupRule(vpcService *vpcclassicv1.VpcClassicV1, sgID string) (rule vpcclassicv1.SecurityGroupRuleIntf, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.CreateSecurityGroupRuleOptions{}
	options.SetSecurityGroupID(sgID)
	options.SetSecurityGroupRulePrototype(&vpcclassicv1.SecurityGroupRulePrototype{
		Direction: core.StringPtr("inbound"),
		Protocol:  core.StringPtr("all"),
		IPVersion: core.StringPtr("ipv4"),
		// Remote: &vpcclassicv1.SecurityGroupRuleTemplateRemote{
		// CidrBlock: core.StringPtr("192.169.0.0/28"),
		// 	Address: core.StringPtr(""),
		// },
	})
	rule, response, err = vpcService.CreateSecurityGroupRule(options)
	return
}

// DeleteSecurityGroupRule DELETE
// /security_groups/{security_group_id}/rules/{id}
// Delete a security group rule
func DeleteSecurityGroupRule(vpcService *vpcclassicv1.VpcClassicV1, sgID, sgRuleID string) (response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.DeleteSecurityGroupRuleOptions{}
	options.SetSecurityGroupID(sgID)
	options.SetID(sgRuleID)
	response, err = vpcService.DeleteSecurityGroupRule(options)
	return
}

// GetSecurityGroupRule GET
// /security_groups/{security_group_id}/rules/{id}
// Retrieve a security group rule
func GetSecurityGroupRule(vpcService *vpcclassicv1.VpcClassicV1, sgID, sgRuleID string) (rule vpcclassicv1.SecurityGroupRuleIntf, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.GetSecurityGroupRuleOptions{}
	options.SetSecurityGroupID(sgID)
	options.SetID(sgRuleID)
	rule, response, err = vpcService.GetSecurityGroupRule(options)
	return
}

// UpdateSecurityGroupRule PATCH
// /security_groups/{security_group_id}/rules/{id}
// Update a security group rule
func UpdateSecurityGroupRule(vpcService *vpcclassicv1.VpcClassicV1, sgID, sgRuleID string) (rule vpcclassicv1.SecurityGroupRuleIntf, response *core.DetailedResponse, err error) {
	securityGroupRulePatchRemoteModel := new(vpcclassicv1.SecurityGroupRuleRemotePatch)
	securityGroupRulePatchRemoteModel.Address = core.StringPtr("192.168.3.4")
	securityGroupRulePatchModel := new(vpcclassicv1.SecurityGroupRulePatch)
	securityGroupRulePatchModel.Remote = securityGroupRulePatchRemoteModel
	patchBody, _ := securityGroupRulePatchModel.AsPatch()
	options := &vpcclassicv1.UpdateSecurityGroupRuleOptions{}
	options.SetSecurityGroupID(sgID)
	options.SetID(sgRuleID)
	options.SetSecurityGroupRulePatch(patchBody)
	rule, response, err = vpcService.UpdateSecurityGroupRule(options)
	rule, response, err = vpcService.UpdateSecurityGroupRule(options)
	return
}

/**
 * Load Balancers
 *
 */

// ListLoadBalancers GET
// /load_balancers
// List all load balancers
func ListLoadBalancers(vpcService *vpcclassicv1.VpcClassicV1) (lbs *vpcclassicv1.LoadBalancerCollection, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.ListLoadBalancersOptions{}
	lbs, response, err = vpcService.ListLoadBalancers(options)
	return
}

// CreateLoadBalancer POST
// /load_balancers
// Create and provision a load balancer
func CreateLoadBalancer(vpcService *vpcclassicv1.VpcClassicV1, name, subnetID string) (lb *vpcclassicv1.LoadBalancer, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.CreateLoadBalancerOptions{}
	options.SetIsPublic(true)
	options.SetName(name)
	var subnetArray = []vpcclassicv1.SubnetIdentityIntf{
		&vpcclassicv1.SubnetIdentity{
			ID: core.StringPtr(subnetID),
		},
	}
	options.SetSubnets(subnetArray)
	lb, response, err = vpcService.CreateLoadBalancer(options)
	return
}

// DeleteLoadBalancer DELETE
// /load_balancers/{id}
// Delete a load balancer
func DeleteLoadBalancer(vpcService *vpcclassicv1.VpcClassicV1, id string) (response *core.DetailedResponse, err error) {
	deleteVpcOptions := &vpcclassicv1.DeleteLoadBalancerOptions{}
	deleteVpcOptions.SetID(id)
	response, err = vpcService.DeleteLoadBalancer(deleteVpcOptions)
	return
}

// GetLoadBalancer GET
// /load_balancers/{id}
// Retrieve a load balancer
func GetLoadBalancer(vpcService *vpcclassicv1.VpcClassicV1, id string) (lb *vpcclassicv1.LoadBalancer, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.GetLoadBalancerOptions{}
	options.SetID(id)
	lb, response, err = vpcService.GetLoadBalancer(options)
	return
}

// UpdateLoadBalancer PATCH
// /load_balancers/{id}
// Update a load balancer
func UpdateLoadBalancer(vpcService *vpcclassicv1.VpcClassicV1, id, name string) (lb *vpcclassicv1.LoadBalancer, response *core.DetailedResponse, err error) {
	body := &vpcclassicv1.AddressPrefixPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcclassicv1.UpdateLoadBalancerOptions{
		LoadBalancerPatch: patchBody,
		ID:                &id,
	}
	lb, response, err = vpcService.UpdateLoadBalancer(options)
	return
}

// GetLoadBalancerStatistics GET
// /load_balancers/{id}/statistics
// List statistics of a load balancer
func GetLoadBalancerStatistics(vpcService *vpcclassicv1.VpcClassicV1, id string) (stat *vpcclassicv1.LoadBalancerStatistics, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.GetLoadBalancerStatisticsOptions{}
	options.SetID(id)
	stat, response, err = vpcService.GetLoadBalancerStatistics(options)
	return
}

// ListLoadBalancerListeners GET
// /load_balancers/{load_balancer_id}/listeners
// List all listeners of the load balancer
func ListLoadBalancerListeners(vpcService *vpcclassicv1.VpcClassicV1, id string) (listeners *vpcclassicv1.LoadBalancerListenerCollection, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.ListLoadBalancerListenersOptions{}
	options.SetLoadBalancerID(id)
	listeners, response, err = vpcService.ListLoadBalancerListeners(options)
	return
}

// CreateLoadBalancerListener POST
// /load_balancers/{load_balancer_id}/listeners
// Create a listener
func CreateLoadBalancerListener(vpcService *vpcclassicv1.VpcClassicV1, lbID string) (listener *vpcclassicv1.LoadBalancerListener, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.CreateLoadBalancerListenerOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetPort(rand.Int63n(100))
	options.SetProtocol("http")
	listener, response, err = vpcService.CreateLoadBalancerListener(options)
	return
}

// DeleteLoadBalancerListener DELETE
// /load_balancers/{load_balancer_id}/listeners/{id}
// Delete a listener
func DeleteLoadBalancerListener(vpcService *vpcclassicv1.VpcClassicV1, lbID, listenerID string) (response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.DeleteLoadBalancerListenerOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetID(listenerID)
	response, err = vpcService.DeleteLoadBalancerListener(options)
	return
}

// GetLoadBalancerListener GET
// /load_balancers/{load_balancer_id}/listeners/{id}
// Retrieve a listener
func GetLoadBalancerListener(vpcService *vpcclassicv1.VpcClassicV1, lbID, listenerID string) (listener *vpcclassicv1.LoadBalancerListener, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.GetLoadBalancerListenerOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetID(listenerID)
	listener, response, err = vpcService.GetLoadBalancerListener(options)
	return
}

// UpdateLoadBalancerListener PATCH
// /load_balancers/{load_balancer_id}/listeners/{id}
// Update a listener
func UpdateLoadBalancerListener(vpcService *vpcclassicv1.VpcClassicV1, lbID, listenerID string) (listener *vpcclassicv1.LoadBalancerListener, response *core.DetailedResponse, err error) {
	body := &vpcclassicv1.LoadBalancerListenerPatch{
		Protocol: core.StringPtr("tcp"),
	}
	patchBody, _ := body.AsPatch()
	options := &vpcclassicv1.UpdateLoadBalancerListenerOptions{
		LoadBalancerListenerPatch: patchBody,
	}
	options.SetLoadBalancerID(lbID)
	options.SetID(listenerID)
	listener, response, err = vpcService.UpdateLoadBalancerListener(options)
	return
}

// ListLoadBalancerListenerPolicies GET
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies
// List all policies of the load balancer listener
func ListLoadBalancerListenerPolicies(vpcService *vpcclassicv1.VpcClassicV1, lbID, listenerID string) (policies *vpcclassicv1.LoadBalancerListenerPolicyCollection, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.ListLoadBalancerListenerPoliciesOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	policies, response, err = vpcService.ListLoadBalancerListenerPolicies(options)
	return
}

// CreateLoadBalancerListenerPolicy POST
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies
func CreateLoadBalancerListenerPolicy(vpcService *vpcclassicv1.VpcClassicV1, lbID, listenerID string) (policy *vpcclassicv1.LoadBalancerListenerPolicy, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.CreateLoadBalancerListenerPolicyOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	options.SetPriority(2)
	options.SetAction("reject")
	policy, response, err = vpcService.CreateLoadBalancerListenerPolicy(options)
	return
}

// DeleteLoadBalancerListenerPolicy DELETE
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies/{id}
// Delete a policy of the load balancer listener
func DeleteLoadBalancerListenerPolicy(vpcService *vpcclassicv1.VpcClassicV1, lbID, listenerID, policyID string) (response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.DeleteLoadBalancerListenerPolicyOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	options.SetID(policyID)
	response, err = vpcService.DeleteLoadBalancerListenerPolicy(options)
	return
}

// GetLoadBalancerListenerPolicy GET
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies/{id}
// Retrieve a policy of the load balancer listener
func GetLoadBalancerListenerPolicy(vpcService *vpcclassicv1.VpcClassicV1, lbID, listenerID, policyID string) (policy *vpcclassicv1.LoadBalancerListenerPolicy, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.GetLoadBalancerListenerPolicyOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	options.SetID(policyID)
	policy, response, err = vpcService.GetLoadBalancerListenerPolicy(options)
	return
}

// UpdateLoadBalancerListenerPolicy PATCH
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies/{id}
// Update a policy of the load balancer listener
func UpdateLoadBalancerListenerPolicy(vpcService *vpcclassicv1.VpcClassicV1, lbID, listenerID, policyID, targetPoolID string) (policy *vpcclassicv1.LoadBalancerListenerPolicy, response *core.DetailedResponse, err error) {
	target := &vpcclassicv1.LoadBalancerListenerPolicyTargetPatch{
		ID: core.StringPtr(targetPoolID),
	}
	model := new(vpcclassicv1.LoadBalancerListenerPolicyPatch)
	model.Name = core.StringPtr("my-policy")
	model.Target = target
	model.Priority = core.Int64Ptr(4)
	patchBody, _ := model.AsPatch()
	options := &vpcclassicv1.UpdateLoadBalancerListenerPolicyOptions{
		LoadBalancerListenerPolicyPatch: patchBody,
	}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	options.SetID(policyID)
	policy, response, err = vpcService.UpdateLoadBalancerListenerPolicy(options)
	return
}

// ListLoadBalancerListenerPolicyRules GET
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies/{policy_id}/rules
// List all rules of the load balancer listener policy
func ListLoadBalancerListenerPolicyRules(vpcService *vpcclassicv1.VpcClassicV1, lbID, listenerID, policyID string) (rules *vpcclassicv1.LoadBalancerListenerPolicyRuleCollection, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.ListLoadBalancerListenerPolicyRulesOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	options.SetPolicyID(policyID)
	rules, response, err = vpcService.ListLoadBalancerListenerPolicyRules(options)
	return
}

// CreateLoadBalancerListenerPolicyRule POST
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies/{policy_id}/rules
// Create a rule for the load balancer listener policy
func CreateLoadBalancerListenerPolicyRule(vpcService *vpcclassicv1.VpcClassicV1, lbID, listenerID, policyID string) (rule *vpcclassicv1.LoadBalancerListenerPolicyRule, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.CreateLoadBalancerListenerPolicyRuleOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	options.SetPolicyID(policyID)
	options.SetCondition("contains")
	options.SetType("hostname")
	options.SetValue("one")
	rule, response, err = vpcService.CreateLoadBalancerListenerPolicyRule(options)
	return
}

// DeleteLoadBalancerListenerPolicyRule DELETE
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies/{policy_id}/rules/{id}
// Delete a rule from the load balancer listener policy
func DeleteLoadBalancerListenerPolicyRule(vpcService *vpcclassicv1.VpcClassicV1, lbID, listenerID, policyID, ruleID string) (response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.DeleteLoadBalancerListenerPolicyRuleOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	options.SetPolicyID(policyID)
	options.SetID(ruleID)
	response, err = vpcService.DeleteLoadBalancerListenerPolicyRule(options)
	return
}

// GetLoadBalancerListenerPolicyRule GET
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies/{policy_id}/rules/{id}
// Retrieve a rule of the load balancer listener policy
func GetLoadBalancerListenerPolicyRule(vpcService *vpcclassicv1.VpcClassicV1, lbID, listenerID, policyID, ruleID string) (rule *vpcclassicv1.LoadBalancerListenerPolicyRule, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.GetLoadBalancerListenerPolicyRuleOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	options.SetPolicyID(policyID)
	options.SetID(ruleID)
	rule, response, err = vpcService.GetLoadBalancerListenerPolicyRule(options)
	return
}

// UpdateLoadBalancerListenerPolicyRule PATCH
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies/{policy_id}/rules/{id}
// Update a rule of the load balancer listener policy
func UpdateLoadBalancerListenerPolicyRule(vpcService *vpcclassicv1.VpcClassicV1, lbID, listenerID, policyID, ruleID string) (rule *vpcclassicv1.LoadBalancerListenerPolicyRule, response *core.DetailedResponse, err error) {
	body := &vpcclassicv1.LoadBalancerListenerPolicyRulePatch{
		Condition: core.StringPtr("equals"),
		Type:      core.StringPtr("header"),
		Value:     core.StringPtr("1"),
		Field:     core.StringPtr("some-field"),
	}
	patchBody, _ := body.AsPatch()
	options := &vpcclassicv1.UpdateLoadBalancerListenerPolicyRuleOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	options.SetPolicyID(policyID)
	options.SetID(ruleID)
	options.SetLoadBalancerListenerPolicyRulePatch(patchBody)
	rule, response, err = vpcService.UpdateLoadBalancerListenerPolicyRule(options)
	return
}

// ListLoadBalancerPools GET
// /load_balancers/{load_balancer_id}/pools
// List all pools of the load balancer
func ListLoadBalancerPools(vpcService *vpcclassicv1.VpcClassicV1, id string) (pools *vpcclassicv1.LoadBalancerPoolCollection, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.ListLoadBalancerPoolsOptions{}
	options.SetLoadBalancerID(id)
	pools, response, err = vpcService.ListLoadBalancerPools(options)
	return
}

// CreateLoadBalancerPool POST
// /load_balancers/{load_balancer_id}/pools
// Create a load balancer pool
func CreateLoadBalancerPool(vpcService *vpcclassicv1.VpcClassicV1, lbID, name string) (pool *vpcclassicv1.LoadBalancerPool, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.CreateLoadBalancerPoolOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetAlgorithm("round_robin")
	options.SetHealthMonitor(&vpcclassicv1.LoadBalancerPoolHealthMonitorPrototype{
		Delay:      core.Int64Ptr(5),
		MaxRetries: core.Int64Ptr(2),
		Timeout:    core.Int64Ptr(4),
		Type:       core.StringPtr("http"),
	})
	options.SetName(name)
	options.SetProtocol("http")
	pool, response, err = vpcService.CreateLoadBalancerPool(options)
	return
}

// DeleteLoadBalancerPool DELETE
// /load_balancers/{load_balancer_id}/pools/{id}
// Delete a pool
func DeleteLoadBalancerPool(vpcService *vpcclassicv1.VpcClassicV1, lbID, poolID string) (response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.DeleteLoadBalancerPoolOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetID(poolID)
	response, err = vpcService.DeleteLoadBalancerPool(options)
	return
}

// GetLoadBalancerPool GET
// /load_balancers/{load_balancer_id}/pools/{id}
// Retrieve a load balancer pool
func GetLoadBalancerPool(vpcService *vpcclassicv1.VpcClassicV1, lbID, poolID string) (pool *vpcclassicv1.LoadBalancerPool, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.GetLoadBalancerPoolOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetID(poolID)
	pool, response, err = vpcService.GetLoadBalancerPool(options)
	return
}

// UpdateLoadBalancerPool PATCH
// /load_balancers/{load_balancer_id}/pools/{id}
// Update a load balancer pool
func UpdateLoadBalancerPool(vpcService *vpcclassicv1.VpcClassicV1, lbID, poolID string) (pool *vpcclassicv1.LoadBalancerPool, response *core.DetailedResponse, err error) {
	body := &vpcclassicv1.LoadBalancerPoolPatch{
		Protocol: core.StringPtr("tcp"),
	}
	patchBody, _ := body.AsPatch()
	options := &vpcclassicv1.UpdateLoadBalancerPoolOptions{
		LoadBalancerPoolPatch: patchBody,
	}
	options.SetLoadBalancerID(lbID)
	options.SetID(poolID)
	pool, response, err = vpcService.UpdateLoadBalancerPool(options)
	return
}

// ListLoadBalancerPoolMembers GET
// /load_balancers/{load_balancer_id}/pools/{pool_id}/members
// List all members of the load balancer pool
func ListLoadBalancerPoolMembers(vpcService *vpcclassicv1.VpcClassicV1, lbID, poolID string) (members *vpcclassicv1.LoadBalancerPoolMemberCollection, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.ListLoadBalancerPoolMembersOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetPoolID(poolID)
	members, response, err = vpcService.ListLoadBalancerPoolMembers(options)
	return
}

// CreateLoadBalancerPoolMember POST
// /load_balancers/{load_balancer_id}/pools/{pool_id}/members
// Create a member in the load balancer pool
func CreateLoadBalancerPoolMember(vpcService *vpcclassicv1.VpcClassicV1, lbID, poolID string) (member *vpcclassicv1.LoadBalancerPoolMember, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.CreateLoadBalancerPoolMemberOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetPoolID(poolID)
	options.SetPort(1234)
	options.SetTarget(&vpcclassicv1.LoadBalancerPoolMemberTargetPrototype{
		Address: core.StringPtr("12.12.0.0"),
	})
	member, response, err = vpcService.CreateLoadBalancerPoolMember(options)
	return
}

// UpdateLoadBalancerPoolMembers PUT
// /load_balancers/{load_balancer_id}/pools/{pool_id}/members
// Update members of the load balancer pool
func UpdateLoadBalancerPoolMembers(vpcService *vpcclassicv1.VpcClassicV1, lbID, poolID string) (member *vpcclassicv1.LoadBalancerPoolMemberCollection, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.ReplaceLoadBalancerPoolMembersOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetPoolID(poolID)
	options.SetMembers([]vpcclassicv1.LoadBalancerPoolMemberPrototype{
		{
			Port: core.Int64Ptr(2345),
			Target: &vpcclassicv1.LoadBalancerPoolMemberTargetPrototype{
				Address: core.StringPtr("13.13.0.0"),
			},
		},
	})
	member, response, err = vpcService.ReplaceLoadBalancerPoolMembers(options)
	return
}

// DeleteLoadBalancerPoolMember DELETE
// /load_balancers/{load_balancer_id}/pools/{pool_id}/members/{id}
// Delete a member from the load balancer pool
func DeleteLoadBalancerPoolMember(vpcService *vpcclassicv1.VpcClassicV1, lbID, poolID, memberID string) (response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.DeleteLoadBalancerPoolMemberOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetPoolID(poolID)
	options.SetID(memberID)
	response, err = vpcService.DeleteLoadBalancerPoolMember(options)
	return
}

// GetLoadBalancerPoolMember GET
// /load_balancers/{load_balancer_id}/pools/{pool_id}/members/{id}
// Retrieve a member in the load balancer pool
func GetLoadBalancerPoolMember(vpcService *vpcclassicv1.VpcClassicV1, lbID, poolID, memberID string) (member *vpcclassicv1.LoadBalancerPoolMember, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.GetLoadBalancerPoolMemberOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetPoolID(poolID)
	options.SetID(memberID)
	member, response, err = vpcService.GetLoadBalancerPoolMember(options)
	return
}

// UpdateLoadBalancerPoolMember PATCH
// /load_balancers/{load_balancer_id}/pools/{pool_id}/members/{id}
func UpdateLoadBalancerPoolMember(vpcService *vpcclassicv1.VpcClassicV1, lbID, poolID, memberID string) (member *vpcclassicv1.LoadBalancerPoolMember, response *core.DetailedResponse, err error) {
	body := &vpcclassicv1.LoadBalancerPoolMemberPatch{
		Port: core.Int64Ptr(3434),
	}
	patchBody, _ := body.AsPatch()
	options := &vpcclassicv1.UpdateLoadBalancerPoolMemberOptions{
		LoadBalancerPoolMemberPatch: patchBody,
	}
	options.SetLoadBalancerID(lbID)
	options.SetPoolID(poolID)
	options.SetID(memberID)
	member, response, err = vpcService.UpdateLoadBalancerPoolMember(options)
	return
}

/**
 * VPN
 *
 */

// ListIkePolicies GET
// /ike_policies
// List all Ike policies
func ListIkePolicies(vpcService *vpcclassicv1.VpcClassicV1) (policies *vpcclassicv1.IkePolicyCollection, response *core.DetailedResponse, err error) {
	options := vpcService.NewListIkePoliciesOptions()
	policies, response, err = vpcService.ListIkePolicies(options)
	return
}

// CreateIkePolicy POST
// /ike_policies
// Create an Ike policy
func CreateIkePolicy(vpcService *vpcclassicv1.VpcClassicV1, name string) (policy *vpcclassicv1.IkePolicy, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.CreateIkePolicyOptions{}
	options.SetName(name)
	options.SetAuthenticationAlgorithm("md5")
	options.SetDhGroup(2)
	options.SetEncryptionAlgorithm("aes128")
	options.SetIkeVersion(1)
	policy, response, err = vpcService.CreateIkePolicy(options)
	return
}

// DeleteIkePolicy DELETE
// /ike_policies/{id}
// Delete an Ike policy
func DeleteIkePolicy(vpcService *vpcclassicv1.VpcClassicV1, id string) (response *core.DetailedResponse, err error) {
	options := vpcService.NewDeleteIkePolicyOptions(id)
	response, err = vpcService.DeleteIkePolicy(options)
	return
}

// GetIkePolicy GET
// /ike_policies/{id}
// Retrieve the specified Ike policy
func GetIkePolicy(vpcService *vpcclassicv1.VpcClassicV1, id string) (policy *vpcclassicv1.IkePolicy, response *core.DetailedResponse, err error) {
	options := vpcService.NewGetIkePolicyOptions(id)
	policy, response, err = vpcService.GetIkePolicy(options)
	return
}

// UpdateIkePolicy PATCH
// /ike_policies/{id}
// Update an Ike policy
func UpdateIkePolicy(vpcService *vpcclassicv1.VpcClassicV1, id string) (policy *vpcclassicv1.IkePolicy, response *core.DetailedResponse, err error) {
	body := &vpcclassicv1.IkePolicyPatch{
		Name:    core.StringPtr("go-ike-policy-2"),
		DhGroup: core.Int64Ptr(5),
	}
	patchBody, _ := body.AsPatch()
	options := &vpcclassicv1.UpdateIkePolicyOptions{
		ID:             core.StringPtr(id),
		IkePolicyPatch: patchBody,
	}
	policy, response, err = vpcService.UpdateIkePolicy(options)
	return
}

// ListVPNGatewayIkePolicyConnections GET
// /ike_policies/{id}/connections
// Lists all the connections that use the specified policy
func ListVPNGatewayIkePolicyConnections(vpcService *vpcclassicv1.VpcClassicV1, id string) (connections *vpcclassicv1.VPNGatewayConnectionCollection, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.ListIkePolicyConnectionsOptions{
		ID: core.StringPtr(id),
	}
	connections, response, err = vpcService.ListIkePolicyConnections(options)
	return
}

// ListIpsecPolicies GET
// /ipsec_policies
// List all IPsec policies
func ListIpsecPolicies(vpcService *vpcclassicv1.VpcClassicV1) (policies *vpcclassicv1.IPsecPolicyCollection, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.ListIpsecPoliciesOptions{}
	policies, response, err = vpcService.ListIpsecPolicies(options)
	return
}

// CreateIpsecPolicy POST
// /ipsec_policies
// Create an IPsec policy
func CreateIpsecPolicy(vpcService *vpcclassicv1.VpcClassicV1, name string) (policy *vpcclassicv1.IPsecPolicy, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.CreateIpsecPolicyOptions{}
	options.SetName(name)
	options.SetAuthenticationAlgorithm("md5")
	options.SetEncryptionAlgorithm("aes128")
	options.SetPfs("disabled")
	policy, response, err = vpcService.CreateIpsecPolicy(options)
	return
}

// DeleteIpsecPolicy DELETE
// /ipsec_policies/{id}
// Delete an IPsec policy
func DeleteIpsecPolicy(vpcService *vpcclassicv1.VpcClassicV1, id string) (response *core.DetailedResponse, err error) {
	options := vpcService.NewDeleteIpsecPolicyOptions(id)
	response, err = vpcService.DeleteIpsecPolicy(options)
	return
}

// GetIpsecPolicy GET
// /ipsec_policies/{id}
// Retrieve the specified IPsec policy
func GetIpsecPolicy(vpcService *vpcclassicv1.VpcClassicV1, id string) (policy *vpcclassicv1.IPsecPolicy, response *core.DetailedResponse, err error) {
	options := vpcService.NewGetIpsecPolicyOptions(id)
	policy, response, err = vpcService.GetIpsecPolicy(options)
	return
}

// UpdateIpsecPolicy PATCH
// /ipsec_policies/{id}
// Update an IPsec policy
func UpdateIpsecPolicy(vpcService *vpcclassicv1.VpcClassicV1, id string) (policy *vpcclassicv1.IPsecPolicy, response *core.DetailedResponse, err error) {
	body := &vpcclassicv1.IPsecPolicyPatch{
		EncryptionAlgorithm: core.StringPtr("3des"),
	}
	patchBody, _ := body.AsPatch()
	options := &vpcclassicv1.UpdateIpsecPolicyOptions{
		ID:               core.StringPtr(id),
		IPsecPolicyPatch: patchBody,
	}
	policy, response, err = vpcService.UpdateIpsecPolicy(options)
	return
}

// ListVPNGatewayIpsecPolicyConnections GET
// /ipsec_policies/{id}/connections
// Lists all the connections that use the specified policy
func ListVPNGatewayIpsecPolicyConnections(vpcService *vpcclassicv1.VpcClassicV1, id string) (connections *vpcclassicv1.VPNGatewayConnectionCollection, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.ListIpsecPolicyConnectionsOptions{
		ID: core.StringPtr(id),
	}
	connections, response, err = vpcService.ListIpsecPolicyConnections(options)
	return
}

// ListVPNGateways GET
// /vpn_gateways
// List all VPN gateways
func ListVPNGateways(vpcService *vpcclassicv1.VpcClassicV1) (gateways *vpcclassicv1.VPNGatewayCollection, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.ListVPNGatewaysOptions{}
	gateways, response, err = vpcService.ListVPNGateways(options)
	return
}

// CreateVPNGateway POST
// /vpn_gateways
// Create a VPN gateway
func CreateVPNGateway(vpcService *vpcclassicv1.VpcClassicV1, subnetID, name string) (gateway vpcclassicv1.VPNGatewayIntf, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.CreateVPNGatewayOptions{
		VPNGatewayPrototype: &vpcclassicv1.VPNGatewayPrototype{
			Name: &name,
			Subnet: &vpcclassicv1.SubnetIdentity{
				ID: core.StringPtr(subnetID),
			},
		},
	}
	gateway, response, err = vpcService.CreateVPNGateway(options)
	return
}

// DeleteVPNGateway DELETE
// /vpn_gateways/{id}
// Delete a VPN gateway
func DeleteVPNGateway(vpcService *vpcclassicv1.VpcClassicV1, id string) (response *core.DetailedResponse, err error) {
	options := vpcService.NewDeleteVPNGatewayOptions(id)
	response, err = vpcService.DeleteVPNGateway(options)
	return
}

// GetVPNGateway GET
// /vpn_gateways/{id}
// Retrieve the specified VPN gateway
func GetVPNGateway(vpcService *vpcclassicv1.VpcClassicV1, id string) (gateway vpcclassicv1.VPNGatewayIntf, response *core.DetailedResponse, err error) {
	options := vpcService.NewGetVPNGatewayOptions(id)
	gateway, response, err = vpcService.GetVPNGateway(options)
	return
}

// UpdateVPNGateway PATCH
// /vpn_gateways/{id}
// Update a VPN gateway
func UpdateVPNGateway(vpcService *vpcclassicv1.VpcClassicV1, id, name string) (gateway vpcclassicv1.VPNGatewayIntf, response *core.DetailedResponse, err error) {
	body := &vpcclassicv1.VPNGatewayPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcclassicv1.UpdateVPNGatewayOptions{
		ID:              core.StringPtr(id),
		VPNGatewayPatch: patchBody,
	}
	gateway, response, err = vpcService.UpdateVPNGateway(options)
	return
}

// ListVPNGatewayConnections GET
// /vpn_gateways/{vpn_gateway_id}/connections
// List all the connections of a VPN gateway
func ListVPNGatewayConnections(vpcService *vpcclassicv1.VpcClassicV1, gatewayID string) (connections *vpcclassicv1.VPNGatewayConnectionCollection, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.ListVPNGatewayConnectionsOptions{}
	options.SetVPNGatewayID(gatewayID)
	connections, response, err = vpcService.ListVPNGatewayConnections(options)
	return
}

// CreateVPNGatewayConnection POST
// /vpn_gateways/{vpn_gateway_id}/connections
// Create a VPN connection
func CreateVPNGatewayConnection(vpcService *vpcclassicv1.VpcClassicV1, gatewayID, name string) (connection vpcclassicv1.VPNGatewayConnectionIntf, response *core.DetailedResponse, err error) {
	peerAddress := "192.168.0.1"
	psk := "pre-shared-key"
	local := []string{"192.132.0.0/28"}
	peer := []string{"197.155.0.0/28"}
	options := &vpcclassicv1.CreateVPNGatewayConnectionOptions{
		VPNGatewayConnectionPrototype: &vpcclassicv1.VPNGatewayConnectionPrototype{
			Name:        &name,
			PeerAddress: &peerAddress,
			Psk:         &psk,
			LocalCIDRs:  local,
			PeerCIDRs:   peer,
		},
	}
	options.SetVPNGatewayID(gatewayID)
	connection, response, err = vpcService.CreateVPNGatewayConnection(options)
	return
}

// DeleteVPNGatewayConnection DELETE
// /vpn_gateways/{vpn_gateway_id}/connections/{id}
// Delete a VPN connection
func DeleteVPNGatewayConnection(vpcService *vpcclassicv1.VpcClassicV1, gatewayID, connID string) (response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.DeleteVPNGatewayConnectionOptions{}
	options.SetVPNGatewayID(gatewayID)
	options.SetID(connID)
	response, err = vpcService.DeleteVPNGatewayConnection(options)
	return
}

// GetVPNGatewayConnection GET
// /vpn_gateways/{vpn_gateway_id}/connections/{id}
// Retrieve the specified VPN connection
func GetVPNGatewayConnection(vpcService *vpcclassicv1.VpcClassicV1, gatewayID, connID string) (connection vpcclassicv1.VPNGatewayConnectionIntf, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.GetVPNGatewayConnectionOptions{}
	options.SetVPNGatewayID(gatewayID)
	options.SetID(connID)
	connection, response, err = vpcService.GetVPNGatewayConnection(options)
	return
}

// UpdateVPNGatewayConnection PATCH
// /vpn_gateways/{vpn_gateway_id}/connections/{id}
// Update a VPN connection
func UpdateVPNGatewayConnection(vpcService *vpcclassicv1.VpcClassicV1, gatewayID, connID, name string) (connection vpcclassicv1.VPNGatewayConnectionIntf, response *core.DetailedResponse, err error) {
	body := &vpcclassicv1.VPNGatewayConnectionPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcclassicv1.UpdateVPNGatewayConnectionOptions{
		ID:                        core.StringPtr(connID),
		VPNGatewayID:              core.StringPtr(gatewayID),
		VPNGatewayConnectionPatch: patchBody,
	}
	connection, response, err = vpcService.UpdateVPNGatewayConnection(options)
	return
}

// ListVPNGatewayConnectionLocalCIDRs GET
// /vpn_gateways/{vpn_gateway_id}/connections/{id}/local_cidrs
// List all local CIDRs for a resource
func ListVPNGatewayConnectionLocalCIDRs(vpcService *vpcclassicv1.VpcClassicV1, gatewayID, connID string) (cidr *vpcclassicv1.VPNGatewayConnectionLocalCIDRs, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.ListVPNGatewayConnectionLocalCIDRsOptions{}
	options.SetVPNGatewayID(gatewayID)
	options.SetID(connID)
	cidr, response, err = vpcService.ListVPNGatewayConnectionLocalCIDRs(options)
	return
}

// DeleteVPNGatewayConnectionLocalCidr DELETE
// /vpn_gateways/{vpn_gateway_id}/connections/{id}/local_cidrs/{prefix_address}/{prefix_length}
// Remove a CIDR from a resource
func DeleteVPNGatewayConnectionLocalCidr(vpcService *vpcclassicv1.VpcClassicV1, gatewayID, connID, prefixAdd, prefixLen string) (response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.RemoveVPNGatewayConnectionLocalCIDROptions{}
	options.SetVPNGatewayID(gatewayID)
	options.SetID(connID)
	options.SetCIDRPrefix(prefixAdd)
	options.SetPrefixLength(prefixLen)
	response, err = vpcService.RemoveVPNGatewayConnectionLocalCIDR(options)
	return
}

// GetVPNGatewayConnectionLocalCidr GET
// /vpn_gateways/{vpn_gateway_id}/connections/{id}/local_cidrs/{prefix_address}/{prefix_length}
// Check if a specific CIDR exists on a specific resource
func CheckVPNGatewayConnectionLocalCidr(vpcService *vpcclassicv1.VpcClassicV1, gatewayID, connID, prefixAdd, prefixLen string) (response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.CheckVPNGatewayConnectionLocalCIDROptions{}
	options.SetVPNGatewayID(gatewayID)
	options.SetID(connID)
	options.SetCIDRPrefix(prefixAdd)
	options.SetPrefixLength(prefixLen)
	response, err = vpcService.CheckVPNGatewayConnectionLocalCIDR(options)
	return
}

// SetVPNGatewayConnectionLocalCidr - PUT
// /vpn_gateways/{vpn_gateway_id}/connections/{id}/local_cidrs/{prefix_address}/{prefix_length}
// Set a CIDR on a resource
func SetVPNGatewayConnectionLocalCidr(vpcService *vpcclassicv1.VpcClassicV1, gatewayID, connID, prefixAdd, prefixLen string) (response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.AddVPNGatewayConnectionLocalCIDROptions{}
	options.SetVPNGatewayID(gatewayID)
	options.SetID(connID)
	options.SetCIDRPrefix(prefixAdd)
	options.SetPrefixLength(prefixLen)
	response, err = vpcService.AddVPNGatewayConnectionLocalCIDR(options)
	return
}

// ListVPNGatewayConnectionPeerCIDRs GET
// /vpn_gateways/{vpn_gateway_id}/connections/{id}/peer_cidrs
// List all peer CIDRs for a resource
func ListVPNGatewayConnectionPeerCIDRs(vpcService *vpcclassicv1.VpcClassicV1, gatewayID, connID string) (cidr *vpcclassicv1.VPNGatewayConnectionPeerCIDRs, response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.ListVPNGatewayConnectionPeerCIDRsOptions{}
	options.SetVPNGatewayID(gatewayID)
	options.SetID(connID)
	cidr, response, err = vpcService.ListVPNGatewayConnectionPeerCIDRs(options)
	return
}

// DeleteVPNGatewayConnectionPeerCidr DELETE
// /vpn_gateways/{vpn_gateway_id}/connections/{id}/peer_cidrs/{prefix_address}/{prefix_length}
// Remove a CIDR from a resource
func DeleteVPNGatewayConnectionPeerCidr(vpcService *vpcclassicv1.VpcClassicV1, gatewayID, connID, prefixAdd, prefixLen string) (response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.RemoveVPNGatewayConnectionPeerCIDROptions{}
	options.SetVPNGatewayID(gatewayID)
	options.SetID(connID)
	options.SetCIDRPrefix(prefixAdd)
	options.SetPrefixLength(prefixLen)
	response, err = vpcService.RemoveVPNGatewayConnectionPeerCIDR(options)
	return
}

// CheckVPNGatewayConnectionPeerCidr GET
// /vpn_gateways/{vpn_gateway_id}/connections/{id}/peer_cidrs/{prefix_address}/{prefix_length}
// Check if a specific CIDR exists on a specific resource
func CheckVPNGatewayConnectionPeerCidr(vpcService *vpcclassicv1.VpcClassicV1, gatewayID, connID, prefixAdd, prefixLen string) (response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.CheckVPNGatewayConnectionPeerCIDROptions{}
	options.SetVPNGatewayID(gatewayID)
	options.SetID(connID)
	options.SetCIDRPrefix(prefixAdd)
	options.SetPrefixLength(prefixLen)
	response, err = vpcService.CheckVPNGatewayConnectionPeerCIDR(options)
	return
}

// SetVPNGatewayConnectionPeerCidr - PUT
// /vpn_gateways/{vpn_gateway_id}/connections/{id}/peer_cidrs/{prefix_address}/{prefix_length}
// Set a CIDR on a resource
func SetVPNGatewayConnectionPeerCidr(vpcService *vpcclassicv1.VpcClassicV1, gatewayID, connID, prefixAdd, prefixLen string) (response *core.DetailedResponse, err error) {
	options := &vpcclassicv1.AddVPNGatewayConnectionPeerCIDROptions{}
	options.SetVPNGatewayID(gatewayID)
	options.SetID(connID)
	options.SetCIDRPrefix(prefixAdd)
	options.SetPrefixLength(prefixLen)
	response, err = vpcService.AddVPNGatewayConnectionPeerCIDR(options)
	return
}

// PollInstance - poll and check the status of VSI before performing any action
// ID - resource ID
// status - expected status/ poll until this status is returned
// pollFrequency - number of times polling happens
func PollInstance(vpcService *vpcclassicv1.VpcClassicV1, ID, status string, pollFrequency int) bool {
	count := 1
	for {
		if count < pollFrequency {
			res, _, err := GetInstance(vpcService, ID)
			fmt.Println("Current status of VSI - ", *res.Status)
			fmt.Println("Expected status of VSI - ", status)
			if err != nil && res == nil {
				fmt.Printf("Error: Retrieving instance ID %s with err error message: %s", ID, err)
				return false
			}
			if *res.Status == status {
				fmt.Println("Received expected status - ", *res.Status)
				return true
			}
			fmt.Printf("Waiting (60 sec) for resource to change status. Attempt - %d", count)
			time.Sleep(60 * time.Second)
			count++

		}
	}
}

// PollSubnet - poll and check the status of VSI before performing any action
// ID - resource ID
// status - expected status/ poll until this status is returned
// pollFrequency - number of times polling happens
func PollSubnet(vpcService *vpcclassicv1.VpcClassicV1, ID, status string, pollFrequency int) bool {
	count := 1
	for {
		if count < pollFrequency {
			res, _, err := GetSubnet(vpcService, ID)
			fmt.Println("Current status of Subnet - ", *res.Status)
			fmt.Println("Expected status of Subnet - ", status)
			if err != nil && res == nil {
				fmt.Printf("Error: Retrieving subnet ID %s with err error message: %s", ID, err)
				return false
			}
			if *res.Status == status {
				fmt.Println("Received expected status - ", *res.Status)
				return true
			}
			fmt.Printf("Waiting (60 sec) for resource to change status. Attempt - %d", count)
			time.Sleep(60 * time.Second)
			count++

		}
	}
}

// PollVolAttachment - poll and check the status of Volume attachment before performing any action
// ID - resource ID
// status - expected status/ poll until this status is returned
// pollFrequency - number of times polling happens
func PollVolAttachment(vpcService *vpcclassicv1.VpcClassicV1, vpcID, volAttachmentID, status string, pollFrequency int) bool {
	count := 1
	for {
		if count < pollFrequency {
			res, _, err := GetVolumeAttachment(vpcService, vpcID, volAttachmentID)
			fmt.Println("Current status of attachment - ", *res.Status)
			fmt.Println("Expected status of attachment - ", status)
			if err != nil && res == nil {
				fmt.Printf("Error: Retrieving volume attachment ID %s with err error message: %s", vpcID, err)
				return false
			}
			if *res.Status == status {
				fmt.Println("Received expected status - ", *res.Status)
				return true
			}
			fmt.Printf("Waiting (60 sec) for resource to change status. Attempt - %d", count)
			time.Sleep(60 * time.Second)
			count++
		}
	}
}

// PollLB - poll and check the status of LB Listener before performing any action
// ID - resource ID
// status - expected status/ poll until this status is returned
// pollFrequency - number of times polling happens
func PollLB(vpcService *vpcclassicv1.VpcClassicV1, lbID, status string, pollFrequency int) bool {
	count := 1
	for {
		if count < pollFrequency {
			res, _, err := GetLoadBalancer(vpcService, lbID)
			fmt.Println("Current status of load balancer - ", *res.ProvisioningStatus)
			fmt.Println("Expected status of load balancer - ", status)
			if err != nil && res == nil {
				fmt.Printf("Error: Retrieving load balancer ID %s with err error message: %s", lbID, err)
				return false
			}
			if *res.ProvisioningStatus == status {
				fmt.Println("Received expected status - ", *res.ProvisioningStatus)
				return true
			}
			fmt.Printf("Waiting (60 sec) for resource to change status. Attempt - %d", count)
			time.Sleep(60 * time.Second)
			count++
		}
	}
}

// PollVPNGateway - poll and check the status of VPNGateway before performing any action
// ID - resource ID
// status - expected status/ poll until this status is returned
// pollFrequency - number of times polling happens
func PollVPNGateway(vpcService *vpcclassicv1.VpcClassicV1, gatewayID, status string, pollFrequency int) bool {
	count := 1
	for {
		if count < pollFrequency {
			res, _, err := GetVPNGateway(vpcService, gatewayID)
			res2B, _ := json.Marshal(res)
			vpn := &vpcclassicv1.VPNGateway{}
			_ = json.Unmarshal([]byte(string(res2B)), &vpn)
			fmt.Println("Current status of VPNGateway - ", *vpn.Status)
			fmt.Println("Expected status of VPNGateway - ", status)
			if err != nil && vpn == nil {
				fmt.Printf("Error: Retrieving VPNGateway ID %s with err error message: %s", gatewayID, err)
				return false
			}
			if *vpn.Status == status {
				fmt.Println("Received expected status - ", *vpn.Status)
				return true
			}
			fmt.Printf("Waiting (60 sec) for resource to change status. Attempt - %d", count)
			time.Sleep(60 * time.Second)
			count++
		}
	}
}

// Counter - Count number of test run.
type Counter struct {
	count int
}

func (counter Counter) currentValue() int {
	return counter.count
}
func (counter *Counter) increment() {
	counter.count++
}

// Print - Marshal JSON and print
func Print(printObject interface{}) {
	p, _ := json.MarshalIndent(printObject, "", "\t")
	fmt.Println(string(p))
}
