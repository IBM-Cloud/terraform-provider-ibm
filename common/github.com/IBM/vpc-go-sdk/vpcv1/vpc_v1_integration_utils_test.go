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
package vpcv1_test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
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

// InstantiateVPCService - Instantiate VPC Gen2 service
func InstantiateVPCService() *vpcv1.VpcV1 {
	service, serviceErr := vpcv1.NewVpcV1UsingExternalConfig(
		&vpcv1.VpcV1Options{
			ServiceName: "vpcint",
		},
	)
	// Check successful instantiation
	if serviceErr != nil {
		fmt.Println("Gen2 Service creation failed.", serviceErr)
		return nil
	}
	// return new vpc gen2 service
	return service
}

/**
 * Regions and Zones
 *
 */

// ListRegions - List all regions
// GET
// /regions
func ListRegions(gen2 *vpcv1.VpcV1) (regions *vpcv1.RegionCollection, response *core.DetailedResponse, err error) {
	listRegionsOptions := &vpcv1.ListRegionsOptions{}
	regions, response, err = gen2.ListRegions(listRegionsOptions)
	return
}

// GetRegion - GETd
// /regions/{name}
// Retrieve a region
func GetRegion(vpcService *vpcv1.VpcV1, name string) (region *vpcv1.Region, response *core.DetailedResponse, err error) {
	getRegionOptions := &vpcv1.GetRegionOptions{}
	getRegionOptions.SetName(name)
	region, response, err = vpcService.GetRegion(getRegionOptions)
	return
}

// ListZones - GET
// /regions/{region_name}/zones
// List all zones in a region
func ListZones(vpcService *vpcv1.VpcV1, regionName string) (zones *vpcv1.ZoneCollection, response *core.DetailedResponse, err error) {
	listZonesOptions := &vpcv1.ListRegionZonesOptions{}
	listZonesOptions.SetRegionName(regionName)
	zones, response, err = vpcService.ListRegionZones(listZonesOptions)
	return
}

// GetZone - GET
// /regions/{region_name}/zones/{zone_name}
// Retrieve a zone
func GetZone(vpcService *vpcv1.VpcV1, regionName, zoneName string) (zone *vpcv1.Zone, response *core.DetailedResponse, err error) {
	getZoneOptions := &vpcv1.GetRegionZoneOptions{}
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
func GetFloatingIPsList(vpcService *vpcv1.VpcV1) (floatingIPs *vpcv1.FloatingIPCollection, response *core.DetailedResponse, err error) {
	listFloatingIpsOptions := vpcService.NewListFloatingIpsOptions()
	floatingIPs, response, err = vpcService.ListFloatingIps(listFloatingIpsOptions)
	return
}

// GetFloatingIP - GET
// /floating_ips/{id}
// Retrieve the specified floating IP
func GetFloatingIP(vpcService *vpcv1.VpcV1, id string) (floatingIP *vpcv1.FloatingIP, response *core.DetailedResponse, err error) {
	options := vpcService.NewGetFloatingIPOptions(id)
	floatingIP, response, err = vpcService.GetFloatingIP(options)
	return
}

// ReleaseFloatingIP - DELETE
// /floating_ips/{id}
// Release the specified floating IP
func ReleaseFloatingIP(vpcService *vpcv1.VpcV1, id string) (response *core.DetailedResponse, err error) {
	options := vpcService.NewDeleteFloatingIPOptions(id)
	response, err = vpcService.DeleteFloatingIP(options)
	return response, err
}

// UpdateFloatingIP - PATCH
// /floating_ips/{id}
// Update the specified floating IP
func UpdateFloatingIP(vpcService *vpcv1.VpcV1, id, name string) (floatingIP *vpcv1.FloatingIP, response *core.DetailedResponse, err error) {
	body := &vpcv1.FloatingIPPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcv1.UpdateFloatingIPOptions{
		ID:              &id,
		FloatingIPPatch: patchBody,
	}

	floatingIP, response, err = vpcService.UpdateFloatingIP(options)
	return
}

// CreateFloatingIP - POST
// /floating_ips
// Reserve a floating IP
func CreateFloatingIP(vpcService *vpcv1.VpcV1, zone, name string) (floatingIP *vpcv1.FloatingIP, response *core.DetailedResponse, err error) {
	options := &vpcv1.CreateFloatingIPOptions{}
	options.SetFloatingIPPrototype(&vpcv1.FloatingIPPrototype{
		Name: &name,
		Zone: &vpcv1.ZoneIdentity{
			Name: &zone,
		},
	})
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
func ListKeys(vpcService *vpcv1.VpcV1) (keys *vpcv1.KeyCollection, response *core.DetailedResponse, err error) {
	listKeysOptions := &vpcv1.ListKeysOptions{}
	keys, response, err = vpcService.ListKeys(listKeysOptions)
	return
}

// GetSSHKey - GET
// /keys/{id}
// Retrieve specified key
func GetSSHKey(vpcService *vpcv1.VpcV1, id string) (key *vpcv1.Key, response *core.DetailedResponse, err error) {
	getKeyOptions := &vpcv1.GetKeyOptions{}
	getKeyOptions.SetID(id)
	key, response, err = vpcService.GetKey(getKeyOptions)
	return
}

// UpdateSSHKey - PATCH
// /keys/{id}
// Update specified key
func UpdateSSHKey(vpcService *vpcv1.VpcV1, id, name string) (key *vpcv1.Key, response *core.DetailedResponse, err error) {
	body := &vpcv1.KeyPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	updateKeyOptions := vpcService.NewUpdateKeyOptions(id, patchBody)
	key, response, err = vpcService.UpdateKey(updateKeyOptions)
	return
}

// DeleteSSHKey - DELETE
// /keys/{id}
// Delete specified key
func DeleteSSHKey(vpcService *vpcv1.VpcV1, id string) (response *core.DetailedResponse, err error) {
	deleteKeyOptions := &vpcv1.DeleteKeyOptions{}
	deleteKeyOptions.SetID(id)
	response, err = vpcService.DeleteKey(deleteKeyOptions)
	return response, err
}

// CreateSSHKey - POST
// /keys
// Create a key
func CreateSSHKey(vpcService *vpcv1.VpcV1, name, publicKey string) (key *vpcv1.Key, response *core.DetailedResponse, err error) {
	options := &vpcv1.CreateKeyOptions{}
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
func ListVpcs(vpcService *vpcv1.VpcV1) (vpcs *vpcv1.VPCCollection, response *core.DetailedResponse, err error) {
	listVpcsOptions := &vpcv1.ListVpcsOptions{}
	vpcs, response, err = vpcService.ListVpcs(listVpcsOptions)
	return
}

// GetVPC - GET
// /vpcs/{id}
// Retrieve specified VPC
func GetVPC(vpcService *vpcv1.VpcV1, id string) (vpc *vpcv1.VPC, response *core.DetailedResponse, err error) {
	getVpcOptions := &vpcv1.GetVPCOptions{}
	getVpcOptions.SetID(id)
	vpc, response, err = vpcService.GetVPC(getVpcOptions)
	return
}

// DeleteVPC - DELETE
// /vpcs/{id}
// Delete specified VPC
func DeleteVPC(vpcService *vpcv1.VpcV1, id string) (response *core.DetailedResponse, err error) {
	deleteVpcOptions := &vpcv1.DeleteVPCOptions{}
	deleteVpcOptions.SetID(id)
	response, err = vpcService.DeleteVPC(deleteVpcOptions)
	return response, err
}

// UpdateVPC - PATCH
// /vpcs/{id}
// Update specified VPC
func UpdateVPC(vpcService *vpcv1.VpcV1, id, name string) (vpc *vpcv1.VPC, response *core.DetailedResponse, err error) {
	body := &vpcv1.VPCPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := vpcService.NewUpdateVPCOptions(id, patchBody)
	vpc, response, err = vpcService.UpdateVPC(options)
	return
}

// CreateVPC - POST
// /vpcs
// Create a VPC
func CreateVPC(vpcService *vpcv1.VpcV1, name, resourceGroup string) (vpc *vpcv1.VPC, response *core.DetailedResponse, err error) {
	options := &vpcv1.CreateVPCOptions{}

	options.SetResourceGroup(&vpcv1.ResourceGroupIdentity{
		ID: &resourceGroup,
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
func GetVPCDefaultSecurityGroup(vpcService *vpcv1.VpcV1, id string) (defaultSg *vpcv1.DefaultSecurityGroup, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetVPCDefaultSecurityGroupOptions{}
	options.SetID(id)
	defaultSg, response, err = vpcService.GetVPCDefaultSecurityGroup(options)
	return
}

/**
 * VPC default ACL
 * Getting default security group for a vpc with id
 */

// GetVPCDefaultACL - GET
// /vpcs/{id}/default_network_acl
// Retrieve a VPC's default network acl
func GetVPCDefaultACL(vpcService *vpcv1.VpcV1, id string) (defaultACL *vpcv1.DefaultNetworkACL, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetVPCDefaultNetworkACLOptions{}
	options.SetID(id)
	defaultACL, response, err = vpcService.GetVPCDefaultNetworkACL(options)
	return
}

/**
 * VPC address prefix
 *
 */

// ListVpcAddressPrefixes - GET
// /vpcs/{vpc_id}/address_prefixes
// List all address pool prefixes for a VPC
func ListVpcAddressPrefixes(vpcService *vpcv1.VpcV1, vpcID string) (addressPrefixes *vpcv1.AddressPrefixCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListVPCAddressPrefixesOptions{}
	options.SetVPCID(vpcID)
	addressPrefixes, response, err = vpcService.ListVPCAddressPrefixes(options)
	return
}

// GetVpcAddressPrefix - GET
// /vpcs/{vpc_id}/address_prefixes/{id}
// Retrieve specified address pool prefix
func GetVpcAddressPrefix(vpcService *vpcv1.VpcV1, vpcID, addressPrefixID string) (addressPrefix *vpcv1.AddressPrefix, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetVPCAddressPrefixOptions{}
	options.SetVPCID(vpcID)
	options.SetID(addressPrefixID)
	addressPrefix, response, err = vpcService.GetVPCAddressPrefix(options)
	return
}

// CreateVpcAddressPrefix - POST
// /vpcs/{vpc_id}/address_prefixes
// Create an address pool prefix
func CreateVpcAddressPrefix(vpcService *vpcv1.VpcV1, vpcID, zone, cidr, name string) (addressPrefix *vpcv1.AddressPrefix, response *core.DetailedResponse, err error) {
	options := &vpcv1.CreateVPCAddressPrefixOptions{}
	options.SetVPCID(vpcID)
	options.SetCIDR(cidr)
	options.SetName(name)
	options.SetZone(&vpcv1.ZoneIdentity{
		Name: &zone,
	})
	addressPrefix, response, err = vpcService.CreateVPCAddressPrefix(options)
	return
}

// DeleteVpcAddressPrefix - DELETE
// /vpcs/{vpc_id}/address_prefixes/{id}
// Delete specified address pool prefix
func DeleteVpcAddressPrefix(vpcService *vpcv1.VpcV1, vpcID, addressPrefixID string) (response *core.DetailedResponse, err error) {
	deleteVpcAddressPrefixOptions := &vpcv1.DeleteVPCAddressPrefixOptions{}
	deleteVpcAddressPrefixOptions.SetVPCID(vpcID)
	deleteVpcAddressPrefixOptions.SetID(addressPrefixID)
	response, err = vpcService.DeleteVPCAddressPrefix(deleteVpcAddressPrefixOptions)
	return response, err
}

// UpdateVpcAddressPrefix - PATCH
// /vpcs/{vpc_id}/address_prefixes/{id}
// Update an address pool prefix
func UpdateVpcAddressPrefix(vpcService *vpcv1.VpcV1, vpcID, addressPrefixID, name string) (addressPrefix *vpcv1.AddressPrefix, response *core.DetailedResponse, err error) {
	body := &vpcv1.AddressPrefixPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := vpcService.NewUpdateVPCAddressPrefixOptions(vpcID, addressPrefixID, patchBody)
	addressPrefix, response, err = vpcService.UpdateVPCAddressPrefix(options)
	return
}

/**
 * VPC routes
 *
 */

// ListVpcRoutes - GET
// /vpcs/{vpc_id}/routes
// List all user-defined routes for a VPC
func ListVpcRoutes(vpcService *vpcv1.VpcV1, vpcID string) (routes *vpcv1.RouteCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListVPCRoutesOptions{}
	options.SetVPCID(vpcID)
	routes, response, err = vpcService.ListVPCRoutes(options)
	return
}

// GetVpcRoute - GET
// /vpcs/{vpc_id}/routes/{id}
// Retrieve the specified route
func GetVpcRoute(vpcService *vpcv1.VpcV1, vpcID, routeID string) (route *vpcv1.Route, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetVPCRouteOptions{}
	options.SetVPCID(vpcID)
	options.SetID(routeID)
	route, response, err = vpcService.GetVPCRoute(options)
	return
}

// CreateVpcRoute - POST
// /vpcs/{vpc_id}/routes
// Create a route on your VPC
func CreateVpcRoute(vpcService *vpcv1.VpcV1, vpcID, zone, destination, nextHopAddress, name string) (route *vpcv1.Route, response *core.DetailedResponse, err error) {
	options := &vpcv1.CreateVPCRouteOptions{}
	options.SetVPCID(vpcID)
	options.SetName(name)
	options.SetZone(&vpcv1.ZoneIdentity{
		Name: &zone,
	})
	options.SetNextHop(&vpcv1.RouteNextHopPrototype{
		Address: &nextHopAddress,
	})
	options.SetDestination(destination)
	route, response, err = vpcService.CreateVPCRoute(options)
	return
}

// DeleteVpcRoute - DELETE
// /vpcs/{vpc_id}/routes/{id}
// Delete the specified route
func DeleteVpcRoute(vpcService *vpcv1.VpcV1, vpcID, routeID string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.DeleteVPCRouteOptions{}
	options.SetVPCID(vpcID)
	options.SetID(routeID)
	response, err = vpcService.DeleteVPCRoute(options)
	return response, err
}

// UpdateVpcRoute - PATCH
// /vpcs/{vpc_id}/routes/{id}
// Update a route
func UpdateVpcRoute(vpcService *vpcv1.VpcV1, vpcID, routeID, name string) (route *vpcv1.Route, response *core.DetailedResponse, err error) {
	body := &vpcv1.RoutePatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcv1.UpdateVPCRouteOptions{
		RoutePatch: patchBody,
		VPCID:      &vpcID,
		ID:         &routeID,
	}
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
func ListVolumeProfiles(vpcService *vpcv1.VpcV1) (profiles *vpcv1.VolumeProfileCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListVolumeProfilesOptions{}
	profiles, response, err = vpcService.ListVolumeProfiles(options)
	return
}

// GetVolumeProfile - GET
// /volume/profiles/{name}
// Retrieve specified volume profile
func GetVolumeProfile(vpcService *vpcv1.VpcV1, profileName string) (profile *vpcv1.VolumeProfile, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetVolumeProfileOptions{}
	options.SetName(profileName)
	profile, response, err = vpcService.GetVolumeProfile(options)
	return
}

// ListVolumes - GET
// /volumes
// List all volumes
func ListVolumes(vpcService *vpcv1.VpcV1) (volumes *vpcv1.VolumeCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListVolumesOptions{}
	volumes, response, err = vpcService.ListVolumes(options)
	return
}

// GetVolume - GET
// /volumes/{id}
// Retrieve specified volume
func GetVolume(vpcService *vpcv1.VpcV1, volumeID string) (volume *vpcv1.Volume, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetVolumeOptions{}
	options.SetID(volumeID)
	volume, response, err = vpcService.GetVolume(options)
	return
}

// DeleteVolume - DELETE
// /volumes/{id}
// Delete specified volume
func DeleteVolume(vpcService *vpcv1.VpcV1, id string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.DeleteVolumeOptions{}
	options.SetID(id)
	response, err = vpcService.DeleteVolume(options)
	return response, err
}

// UpdateVolume - PATCH
// /volumes/{id}
// Update specified volume
func UpdateVolume(vpcService *vpcv1.VpcV1, id, name string) (volume *vpcv1.Volume, response *core.DetailedResponse, err error) {
	body := &vpcv1.AddressPrefixPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcv1.UpdateVolumeOptions{
		VolumePatch: patchBody,
		ID:          &id,
	}
	volume, response, err = vpcService.UpdateVolume(options)
	return
}

// CreateVolume - POST
// /volumes
// Create a volume
func CreateVolume(vpcService *vpcv1.VpcV1, name, profileName, zoneName string, capacity int64) (volume *vpcv1.Volume, response *core.DetailedResponse, err error) {
	options := &vpcv1.CreateVolumeOptions{}
	options.SetVolumePrototype(&vpcv1.VolumePrototype{
		Capacity: core.Int64Ptr(capacity),
		Zone: &vpcv1.ZoneIdentity{
			Name: &zoneName,
		},
		Profile: &vpcv1.VolumeProfileIdentity{
			Name: &profileName,
		},
		Name: &name,
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
func ListSubnets(vpcService *vpcv1.VpcV1) (subnets *vpcv1.SubnetCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListSubnetsOptions{}
	subnets, response, err = vpcService.ListSubnets(options)
	return
}

// GetSubnet - GET
// /subnets/{id}
// Retrieve specified subnet
func GetSubnet(vpcService *vpcv1.VpcV1, subnetID string) (subnet *vpcv1.Subnet, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetSubnetOptions{}
	options.SetID(subnetID)
	subnet, response, err = vpcService.GetSubnet(options)
	return
}

// DeleteSubnet - DELETE
// /subnets/{id}
// Delete specified subnet
func DeleteSubnet(vpcService *vpcv1.VpcV1, id string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.DeleteSubnetOptions{}
	options.SetID(id)
	response, err = vpcService.DeleteSubnet(options)
	return response, err
}

// UpdateSubnet - PATCH
// /subnets/{id}
// Update specified subnet
func UpdateSubnet(vpcService *vpcv1.VpcV1, id, name string) (subnet *vpcv1.Subnet, response *core.DetailedResponse, err error) {
	body := &vpcv1.SubnetPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcv1.UpdateSubnetOptions{
		SubnetPatch: patchBody,
	}
	options.SetID(id)
	subnet, response, err = vpcService.UpdateSubnet(options)
	return
}

// CreateSubnet - POST
// /subnets
// Create a subnet
func CreateSubnet(vpcService *vpcv1.VpcV1, vpcID, name, zone string, mock bool) (subnet *vpcv1.Subnet, response *core.DetailedResponse, err error) {
	options := &vpcv1.CreateSubnetOptions{}
	if mock {
		options.SetSubnetPrototype(&vpcv1.SubnetPrototype{
			Ipv4CIDRBlock: core.StringPtr("10.243.0.0/24"),
			Name:          &name,
			VPC: &vpcv1.VPCIdentity{
				ID: &vpcID,
			},
			Zone: &vpcv1.ZoneIdentity{
				Name: &zone,
			},
		})
	} else {
		options.SetSubnetPrototype(&vpcv1.SubnetPrototype{
			Name: &name,
			VPC: &vpcv1.VPCIdentity{
				ID: &vpcID,
			},
			Zone: &vpcv1.ZoneIdentity{
				Name: &zone,
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
func GetSubnetNetworkAcl(vpcService *vpcv1.VpcV1, subnetID string) (subnetACL *vpcv1.NetworkACL, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetSubnetNetworkACLOptions{}
	options.SetID(subnetID)
	subnetACL, response, err = vpcService.GetSubnetNetworkACL(options)
	return
}

// SetSubnetNetworkAclBinding - PUT
// /subnets/{id}/network_acl
// Attach a network ACL to a subnet
func SetSubnetNetworkAclBinding(vpcService *vpcv1.VpcV1, subnetID, id string) (nacl *vpcv1.NetworkACL, response *core.DetailedResponse, err error) {
	options := &vpcv1.ReplaceSubnetNetworkACLOptions{}
	options.SetID(subnetID)
	options.SetNetworkACLIdentity(&vpcv1.NetworkACLIdentity{ID: &id})
	nacl, response, err = vpcService.ReplaceSubnetNetworkACL(options)
	return
}

// DeleteSubnetPublicGatewayBinding - DELETE
// /subnets/{id}/public_gateway
// Detach a public gateway from a subnet
func DeleteSubnetPublicGatewayBinding(vpcService *vpcv1.VpcV1, id string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.UnsetSubnetPublicGatewayOptions{}
	options.SetID(id)
	response, err = vpcService.UnsetSubnetPublicGateway(options)
	return response, err
}

// GetSubnetPublicGateway - GET
// /subnets/{id}/public_gateway
// Retrieve a subnet's attached public gateway
func GetSubnetPublicGateway(vpcService *vpcv1.VpcV1, subnetID string) (pgw *vpcv1.PublicGateway, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetSubnetPublicGatewayOptions{}
	options.SetID(subnetID)
	pgw, response, err = vpcService.GetSubnetPublicGateway(options)
	return
}

// SetSubnetPublicGatewayBinding - PUT
// /subnets/{id}/public_gateway
// Attach a public gateway to a subnet
func CreateSubnetPublicGatewayBinding(vpcService *vpcv1.VpcV1, subnetID, id string) (pgw *vpcv1.PublicGateway, response *core.DetailedResponse, err error) {
	options := &vpcv1.SetSubnetPublicGatewayOptions{}
	options.SetID(subnetID)
	options.SetPublicGatewayIdentity(&vpcv1.PublicGatewayIdentity{ID: &id})
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
func ListImages(vpcService *vpcv1.VpcV1, visibility string) (images *vpcv1.ImageCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListImagesOptions{}
	options.SetVisibility(visibility)
	images, response, err = vpcService.ListImages(options)
	return
}

// GetImage - GET
// /images/{id}
// Retrieve the specified image
func GetImage(vpcService *vpcv1.VpcV1, imageID string) (image *vpcv1.Image, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetImageOptions{}
	options.SetID(imageID)
	image, response, err = vpcService.GetImage(options)
	return
}

// DeleteImage DELETE
// /images/{id}
// Delete specified image
func DeleteImage(vpcService *vpcv1.VpcV1, id string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.DeleteImageOptions{}
	options.SetID(id)
	response, err = vpcService.DeleteImage(options)
	return response, err
}

// UpdateImage PATCH
// /images/{id}
// Update specified image
func UpdateImage(vpcService *vpcv1.VpcV1, id, name string) (image *vpcv1.Image, response *core.DetailedResponse, err error) {
	body := &vpcv1.ImagePatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcv1.UpdateImageOptions{
		ImagePatch: patchBody,
	}
	options.SetID(id)
	image, response, err = vpcService.UpdateImage(options)
	return
}

func CreateImage(vpcService *vpcv1.VpcV1, name string) (image *vpcv1.Image, response *core.DetailedResponse, err error) {
	options := &vpcv1.CreateImageOptions{}
	cosID := "cos://cos-location-of-image-file"
	options.SetImagePrototype(&vpcv1.ImagePrototype{
		Name: &name,
		File: &vpcv1.ImageFilePrototype{
			Href: &cosID,
		},
	})
	image, response, err = vpcService.CreateImage(options)
	return
}

func ListOperatingSystems(vpcService *vpcv1.VpcV1) (operatingSystems *vpcv1.OperatingSystemCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListOperatingSystemsOptions{}
	operatingSystems, response, err = vpcService.ListOperatingSystems(options)
	return
}

func GetOperatingSystem(vpcService *vpcv1.VpcV1, osName string) (os *vpcv1.OperatingSystem, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetOperatingSystemOptions{}
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
func ListInstanceProfiles(vpcService *vpcv1.VpcV1) (profiles *vpcv1.InstanceProfileCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListInstanceProfilesOptions{}
	profiles, response, err = vpcService.ListInstanceProfiles(options)
	return
}

// GetInstanceProfile - GET
// /instance/profiles/{name}
// Retrieve specified instance profile
func GetInstanceProfile(vpcService *vpcv1.VpcV1, profileName string) (profile *vpcv1.InstanceProfile, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetInstanceProfileOptions{}
	options.SetName(profileName)
	profile, response, err = vpcService.GetInstanceProfile(options)
	return
}

// ListInstances GET
// /instances
// List all instances
func ListInstances(vpcService *vpcv1.VpcV1) (instances *vpcv1.InstanceCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListInstancesOptions{}
	instances, response, err = vpcService.ListInstances(options)
	return
}

// GetInstance GET
// instances/{id}
// Retrieve an instance
func GetInstance(vpcService *vpcv1.VpcV1, instanceID string) (instance *vpcv1.Instance, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetInstanceOptions{}
	options.SetID(instanceID)
	instance, response, err = vpcService.GetInstance(options)
	return
}

// DeleteInstance DELETE
// /instances/{id}
// Delete specified instance
func DeleteInstance(vpcService *vpcv1.VpcV1, id string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.DeleteInstanceOptions{}
	options.SetID(id)
	response, err = vpcService.DeleteInstance(options)
	return response, err
}

// UpdateInstance PATCH
// /instances/{id}
// Update specified instance
func UpdateInstance(vpcService *vpcv1.VpcV1, id, name string) (instance *vpcv1.Instance, response *core.DetailedResponse, err error) {
	body := &vpcv1.InstancePatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcv1.UpdateInstanceOptions{
		InstancePatch: patchBody,
		ID:            &id,
	}
	instance, response, err = vpcService.UpdateInstance(options)
	return
}

// CreateInstance POST
// /instances/{instance_id}
// Create an instance action
func CreateInstance(vpcService *vpcv1.VpcV1, name, profileName, imageID, zoneName, subnetID, sshkeyID, vpcID string) (instance *vpcv1.Instance, response *core.DetailedResponse, err error) {
	options := &vpcv1.CreateInstanceOptions{}
	options.SetInstancePrototype(&vpcv1.InstancePrototype{
		Name: &name,
		Image: &vpcv1.ImageIdentity{
			ID: &imageID,
		},
		Profile: &vpcv1.InstanceProfileIdentity{
			Name: &profileName,
		},
		Zone: &vpcv1.ZoneIdentity{
			Name: &zoneName,
		},
		PrimaryNetworkInterface: &vpcv1.NetworkInterfacePrototype{
			Subnet: &vpcv1.SubnetIdentity{
				ID: &subnetID,
			},
		},
		Keys: []vpcv1.KeyIdentityIntf{
			&vpcv1.KeyIdentity{
				ID: &sshkeyID,
			},
		},
		VPC: &vpcv1.VPCIdentity{
			ID: &vpcID,
		},
	})
	instance, response, err = vpcService.CreateInstance(options)
	return
}

// CreateInstanceAction PATCH
// /instances/{instance_id}/actions
// Update specified instance
func CreateInstanceAction(vpcService *vpcv1.VpcV1, instanceID, typeOfAction string) (action *vpcv1.InstanceAction, response *core.DetailedResponse, err error) {
	options := &vpcv1.CreateInstanceActionOptions{}
	options.SetInstanceID(instanceID)
	options.SetType(typeOfAction)
	action, response, err = vpcService.CreateInstanceAction(options)
	return
}

// GetInstanceInitialization GET
// /instances/{id}/initialization
// Retrieve configuration used to initialize the instance.
func GetInstanceInitialization(vpcService *vpcv1.VpcV1, instanceID string) (initData *vpcv1.InstanceInitialization, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetInstanceInitializationOptions{}
	options.SetID(instanceID)
	initData, response, err = vpcService.GetInstanceInitialization(options)
	return
}

// ListNetworkInterfaces GET
// /instances/{instance_id}/network_interfaces
// List all network interfaces on an instance
func ListNetworkInterfaces(vpcService *vpcv1.VpcV1, id string) (networkInterfaces *vpcv1.NetworkInterfaceUnpaginatedCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListInstanceNetworkInterfacesOptions{}
	options.SetInstanceID(id)
	networkInterfaces, response, err = vpcService.ListInstanceNetworkInterfaces(options)
	return
}

// CreateNetworkInterface POST
// /instances/{instance_id}/network_interfaces
// List all network interfaces on an instance
func CreateNetworkInterface(vpcService *vpcv1.VpcV1, id, subnetID string) (networkInterface *vpcv1.NetworkInterface, response *core.DetailedResponse, err error) {
	options := &vpcv1.CreateInstanceNetworkInterfaceOptions{}
	options.SetInstanceID(id)
	options.SetName("eth1")
	options.SetSubnet(&vpcv1.SubnetIdentityByID{
		ID: &subnetID,
	})
	networkInterface, response, err = vpcService.CreateInstanceNetworkInterface(options)
	return
}

// DeleteNetworkInterface Delete
// /instances/{instance_id}/network_interfaces/{id}
// Retrieve specified network interface
func DeleteNetworkInterface(vpcService *vpcv1.VpcV1, instanceID, vnicID string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.DeleteInstanceNetworkInterfaceOptions{}
	options.SetID(vnicID)
	options.SetInstanceID(instanceID)
	response, err = vpcService.DeleteInstanceNetworkInterface(options)
	return response, err
}

// GetNetworkInterface GET
// /instances/{instance_id}/network_interfaces/{id}
// Retrieve specified network interface
func GetNetworkInterface(vpcService *vpcv1.VpcV1, instanceID, networkID string) (networkInterface *vpcv1.NetworkInterface, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetInstanceNetworkInterfaceOptions{}
	options.SetID(networkID)
	options.SetInstanceID(instanceID)
	networkInterface, response, err = vpcService.GetInstanceNetworkInterface(options)
	return
}

// UpdateNetworkInterface PATCH
// /instances/{instance_id}/network_interfaces/{id}
// Update a network interface
func UpdateNetworkInterface(vpcService *vpcv1.VpcV1, instanceID, networkID, name string) (networkInterface *vpcv1.NetworkInterface, response *core.DetailedResponse, err error) {
	body := &vpcv1.NetworkInterfacePatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcv1.UpdateInstanceNetworkInterfaceOptions{
		NetworkInterfacePatch: patchBody,
		ID:                    &networkID,
		InstanceID:            &instanceID,
	}
	networkInterface, response, err = vpcService.UpdateInstanceNetworkInterface(options)
	return
}

// ListNetworkInterfaceFloatingIps GET
// /instances/{instance_id}/network_interfaces
// List all network interfaces on an instance
func ListNetworkInterfaceFloatingIps(vpcService *vpcv1.VpcV1, instanceID, networkID string) (fips *vpcv1.FloatingIPUnpaginatedCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListInstanceNetworkInterfaceFloatingIpsOptions{}
	options.SetInstanceID(instanceID)
	options.SetNetworkInterfaceID(networkID)
	fips, response, err = vpcService.ListInstanceNetworkInterfaceFloatingIps(options)
	return
}

// GetNetworkInterfaceFloatingIp GET
// /instances/{instance_id}/network_interfaces/{network_interface_id}/floating_ips
// List all floating IPs associated with a network interface
func GetNetworkInterfaceFloatingIp(vpcService *vpcv1.VpcV1, instanceID, networkID, fipID string) (fip *vpcv1.FloatingIP, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetInstanceNetworkInterfaceFloatingIPOptions{}
	options.SetID(fipID)
	options.SetInstanceID(instanceID)
	options.SetNetworkInterfaceID(networkID)
	fip, response, err = vpcService.GetInstanceNetworkInterfaceFloatingIP(options)
	return
}

// DeleteNetworkInterfaceFloatingIpBinding DELETE
// /instances/{instance_id}/network_interfaces/{network_interface_id}/floating_ips/{id}
// Disassociate specified floating IP
func DeleteNetworkInterfaceFloatingIpBinding(vpcService *vpcv1.VpcV1, instanceID, networkID, fipID string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.RemoveInstanceNetworkInterfaceFloatingIPOptions{}
	options.SetID(fipID)
	options.SetInstanceID(instanceID)
	options.SetNetworkInterfaceID(networkID)
	response, err = vpcService.RemoveInstanceNetworkInterfaceFloatingIP(options)
	return response, err
}

// CreateNetworkInterfaceFloatingIpBinding PUT
// /instances/{instance_id}/network_interfaces/{network_interface_id}/floating_ips/{id}
// Associate a floating IP with a network interface
func CreateNetworkInterfaceFloatingIpBinding(vpcService *vpcv1.VpcV1, instanceID, networkID, fipID string) (fip *vpcv1.FloatingIP, response *core.DetailedResponse, err error) {
	options := &vpcv1.AddInstanceNetworkInterfaceFloatingIPOptions{}
	options.SetID(fipID)
	options.SetInstanceID(instanceID)
	options.SetNetworkInterfaceID(networkID)
	fip, response, err = vpcService.AddInstanceNetworkInterfaceFloatingIP(options)
	return
}

// ListVolumeAttachments GET
// /instances/{instance_id}/volume_attachments
// List all volumes attached to an instance
func ListVolumeAttachments(vpcService *vpcv1.VpcV1, id string) (volumeAttachments *vpcv1.VolumeAttachmentCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListInstanceVolumeAttachmentsOptions{}
	options.SetInstanceID(id)
	volumeAttachments, response, err = vpcService.ListInstanceVolumeAttachments(options)
	return
}

// CreateVolumeAttachment POST
// /instances/{instance_id}/volume_attachments
// Create a volume attachment, connecting a volume to an instance
func CreateVolumeAttachment(vpcService *vpcv1.VpcV1, instanceID, volumeID, name string) (volumeAttachment *vpcv1.VolumeAttachment, response *core.DetailedResponse, err error) {
	options := &vpcv1.CreateInstanceVolumeAttachmentOptions{}
	options.SetInstanceID(instanceID)
	options.SetVolume(&vpcv1.VolumeIdentity{
		ID: &volumeID,
	})
	options.SetName(name)
	options.SetDeleteVolumeOnInstanceDelete(false)
	volumeAttachment, response, err = vpcService.CreateInstanceVolumeAttachment(options)
	return
}

// DeleteVolumeAttachment DELETE
// /instances/{instance_id}/volume_attachments/{id}
// Delete a volume attachment, detaching a volume from an instance
func DeleteVolumeAttachment(vpcService *vpcv1.VpcV1, instanceID, volumeID string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.DeleteInstanceVolumeAttachmentOptions{}
	options.SetID(volumeID)
	options.SetInstanceID(instanceID)
	response, err = vpcService.DeleteInstanceVolumeAttachment(options)
	return response, err
}

// GetVolumeAttachment GET
// /instances/{instance_id}/volume_attachments/{id}
// Retrieve specified volume attachment
func GetVolumeAttachment(vpcService *vpcv1.VpcV1, instanceID, volumeID string) (volumeAttachment *vpcv1.VolumeAttachment, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetInstanceVolumeAttachmentOptions{}
	options.SetInstanceID(instanceID)
	options.SetID(volumeID)
	volumeAttachment, response, err = vpcService.GetInstanceVolumeAttachment(options)
	return
}

// UpdateVolumeAttachment PATCH
// /instances/{instance_id}/volume_attachments/{id}
// Update a volume attachment
func UpdateVolumeAttachment(vpcService *vpcv1.VpcV1, instanceID, volumeID, name string) (volumeAttachment *vpcv1.VolumeAttachment, response *core.DetailedResponse, err error) {
	body := &vpcv1.VolumeAttachmentPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcv1.UpdateInstanceVolumeAttachmentOptions{
		VolumeAttachmentPatch: patchBody,
		InstanceID:            &instanceID,
		ID:                    &volumeID,
	}
	volumeAttachment, response, err = vpcService.UpdateInstanceVolumeAttachment(options)
	return
}

/**
 * Public Gateway
 *
 */

// ListPublicGateways GET
// /public_gateways
// List all public gateways
func ListPublicGateways(vpcService *vpcv1.VpcV1) (pgws *vpcv1.PublicGatewayCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListPublicGatewaysOptions{}
	pgws, response, err = vpcService.ListPublicGateways(options)
	return
}

// CreatePublicGateway POST
// /public_gateways
// Create a public gateway
func CreatePublicGateway(vpcService *vpcv1.VpcV1, name, vpcID, zoneName string) (pgw *vpcv1.PublicGateway, response *core.DetailedResponse, err error) {
	options := &vpcv1.CreatePublicGatewayOptions{}
	options.SetVPC(&vpcv1.VPCIdentity{
		ID: &vpcID,
	})
	options.SetZone(&vpcv1.ZoneIdentity{
		Name: &zoneName,
	})
	pgw, response, err = vpcService.CreatePublicGateway(options)
	return
}

// DeletePublicGateway DELETE
// /public_gateways/{id}
// Delete specified public gateway
func DeletePublicGateway(vpcService *vpcv1.VpcV1, id string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.DeletePublicGatewayOptions{}
	options.SetID(id)
	response, err = vpcService.DeletePublicGateway(options)
	return response, err
}

// GetPublicGateway GET
// /public_gateways/{id}
// Retrieve specified public gateway
func GetPublicGateway(vpcService *vpcv1.VpcV1, id string) (pgw *vpcv1.PublicGateway, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetPublicGatewayOptions{}
	options.SetID(id)
	pgw, response, err = vpcService.GetPublicGateway(options)
	return
}

// UpdatePublicGateway PATCH
// /public_gateways/{id}
// Update a public gateway's name
func UpdatePublicGateway(vpcService *vpcv1.VpcV1, id, name string) (pgw *vpcv1.PublicGateway, response *core.DetailedResponse, err error) {
	body := &vpcv1.PublicGatewayPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcv1.UpdatePublicGatewayOptions{
		PublicGatewayPatch: patchBody,
		ID:                 &id,
	}
	pgw, response, err = vpcService.UpdatePublicGateway(options)
	return
}

/**
 * Network ACLs not available in gen2 currently
 *
 */

// ListNetworkAcls - GET
// /network_acls
// List all network ACLs
func ListNetworkAcls(vpcService *vpcv1.VpcV1) (networkACLs *vpcv1.NetworkACLCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListNetworkAclsOptions{}
	networkACLs, response, err = vpcService.ListNetworkAcls(options)
	return
}

// CreateNetworkAcl - POST
// /network_acls
// Create a network ACL
func CreateNetworkAcl(vpcService *vpcv1.VpcV1, name, copyableAclID, vpcID string) (networkACL *vpcv1.NetworkACL, response *core.DetailedResponse, err error) {
	options := &vpcv1.CreateNetworkACLOptions{}
	options.SetNetworkACLPrototype(&vpcv1.NetworkACLPrototype{
		Name: &name,
		SourceNetworkACL: &vpcv1.NetworkACLIdentity{
			ID: &copyableAclID,
		},
		VPC: &vpcv1.VPCIdentity{
			ID: &vpcID,
		},
	})
	networkACL, response, err = vpcService.CreateNetworkACL(options)
	return
}

// DeleteNetworkAcl - DELETE
// /network_acls/{id}
// Delete specified network ACL
func DeleteNetworkAcl(vpcService *vpcv1.VpcV1, ID string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.DeleteNetworkACLOptions{}
	options.SetID(ID)
	response, err = vpcService.DeleteNetworkACL(options)
	return response, err
}

// GetNetworkAcl - GET
// /network_acls/{id}
// Retrieve specified network ACL
func GetNetworkAcl(vpcService *vpcv1.VpcV1, ID string) (networkACL *vpcv1.NetworkACL, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetNetworkACLOptions{}
	options.SetID(ID)
	networkACL, response, err = vpcService.GetNetworkACL(options)
	return
}

// UpdateNetworkAcl PATCH
// /network_acls/{id}
// Update a network ACL
func UpdateNetworkAcl(vpcService *vpcv1.VpcV1, id, name string) (networkACL *vpcv1.NetworkACL, response *core.DetailedResponse, err error) {
	body := &vpcv1.NetworkACLPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcv1.UpdateNetworkACLOptions{
		NetworkACLPatch: patchBody,
		ID:              &id,
	}
	networkACL, response, err = vpcService.UpdateNetworkACL(options)
	return
}

// ListNetworkAclRules - GET
// /network_acls/{network_acl_id}/rules
// List all rules for a network ACL
func ListNetworkAclRules(vpcService *vpcv1.VpcV1, aclID string) (networkACL *vpcv1.NetworkACLRuleCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListNetworkACLRulesOptions{}
	options.SetNetworkACLID(aclID)
	networkACL, response, err = vpcService.ListNetworkACLRules(options)
	return
}

// CreateNetworkAclRule - POST
// /network_acls/{network_acl_id}/rules
// Create a rule
func CreateNetworkAclRule(vpcService *vpcv1.VpcV1, name, aclID string) (rules vpcv1.NetworkACLRuleIntf, response *core.DetailedResponse, err error) {
	options := &vpcv1.CreateNetworkACLRuleOptions{}
	options.SetNetworkACLID(aclID)
	options.SetNetworkACLRulePrototype(&vpcv1.NetworkACLRulePrototype{
		Action:      core.StringPtr("allow"),
		Direction:   core.StringPtr("inbound"),
		Destination: core.StringPtr("0.0.0.0/0"),
		Source:      core.StringPtr("0.0.0.0/0"),
		Protocol:    core.StringPtr("all"),
		Name:        &name,
	})
	rules, response, err = vpcService.CreateNetworkACLRule(options)
	return
}

// DeleteNetworkAclRule DELETE
// /network_acls/{network_acl_id}/rules/{id}
// Delete specified rule
func DeleteNetworkAclRule(vpcService *vpcv1.VpcV1, aclID, ruleID string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.DeleteNetworkACLRuleOptions{}
	options.SetID(ruleID)
	options.SetNetworkACLID(aclID)
	response, err = vpcService.DeleteNetworkACLRule(options)
	return response, err
}

// GetNetworkAclRule GET
// /network_acls/{network_acl_id}/rules/{id}
// Retrieve specified rule
func GetNetworkAclRule(vpcService *vpcv1.VpcV1, aclID, ruleID string) (rule vpcv1.NetworkACLRuleIntf, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetNetworkACLRuleOptions{}
	options.SetID(ruleID)
	options.SetNetworkACLID(aclID)
	rule, response, err = vpcService.GetNetworkACLRule(options)
	return
}

// UpdateNetworkAclRule PATCH
// /network_acls/{network_acl_id}/rules/{id}
// Update a rule
func UpdateNetworkAclRule(vpcService *vpcv1.VpcV1, aclID, ruleID, name string) (rule vpcv1.NetworkACLRuleIntf, response *core.DetailedResponse, err error) {
	body := &vpcv1.NetworkACLRulePatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcv1.UpdateNetworkACLRuleOptions{
		NetworkACLRulePatch: patchBody,
		ID:                  &ruleID,
		NetworkACLID:        &aclID,
	}
	rule, response, err = vpcService.UpdateNetworkACLRule(options)
	return
}

/**
 * Security Groups
 *
 */

// ListSecurityGroups GET
// /security_groups
// List all security groups
func ListSecurityGroups(vpcService *vpcv1.VpcV1) (securityGroups *vpcv1.SecurityGroupCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListSecurityGroupsOptions{}
	securityGroups, response, err = vpcService.ListSecurityGroups(options)
	return
}

// CreateSecurityGroup POST
// /security_groups
// Create a security group
func CreateSecurityGroup(vpcService *vpcv1.VpcV1, name, vpcID string) (securityGroup *vpcv1.SecurityGroup, response *core.DetailedResponse, err error) {
	options := &vpcv1.CreateSecurityGroupOptions{}
	options.SetVPC(&vpcv1.VPCIdentity{
		ID: &vpcID,
	})
	options.SetName(name)
	securityGroup, response, err = vpcService.CreateSecurityGroup(options)
	return
}

// DeleteSecurityGroup DELETE
// /security_groups/{id}
// Delete a security group
func DeleteSecurityGroup(vpcService *vpcv1.VpcV1, id string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.DeleteSecurityGroupOptions{}
	options.SetID(id)
	response, err = vpcService.DeleteSecurityGroup(options)
	return response, err
}

// GetSecurityGroup GET
// /security_groups/{id}
// Retrieve a security group
func GetSecurityGroup(vpcService *vpcv1.VpcV1, id string) (securityGroup *vpcv1.SecurityGroup, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetSecurityGroupOptions{}
	options.SetID(id)
	securityGroup, response, err = vpcService.GetSecurityGroup(options)
	return
}

// UpdateSecurityGroup PATCH
// /security_groups/{id}
// Update a security group
func UpdateSecurityGroup(vpcService *vpcv1.VpcV1, id, name string) (securityGroup *vpcv1.SecurityGroup, response *core.DetailedResponse, err error) {
	body := &vpcv1.SecurityGroupPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcv1.UpdateSecurityGroupOptions{
		SecurityGroupPatch: patchBody,
		ID:                 &id,
	}
	securityGroup, response, err = vpcService.UpdateSecurityGroup(options)
	return
}

// ListSecurityGroupNetworkInterfaces GET
// /security_groups/{security_group_id}/network_interfaces
// List a security group's network interfaces
// ListSecurityGroupNetworkInterfaces
func ListSecurityGroupNetworkInterfaces(vpcService *vpcv1.VpcV1, sgID string) (netInterfaces *vpcv1.NetworkInterfaceCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListSecurityGroupNetworkInterfacesOptions{}
	options.SetSecurityGroupID(sgID)
	netInterfaces, response, err = vpcService.ListSecurityGroupNetworkInterfaces(options)
	return
}

// DeleteSecurityGroupNetworkInterfaceBinding DELETE
// /security_groups/{security_group_id}/network_interfaces/{id}
// Remove a network interface from a security group.
func DeleteSecurityGroupNetworkInterfaceBinding(vpcService *vpcv1.VpcV1, id, vnicID string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.RemoveSecurityGroupNetworkInterfaceOptions{}
	options.SetSecurityGroupID(id)
	options.SetID(vnicID)
	response, err = vpcService.RemoveSecurityGroupNetworkInterface(options)
	return response, err
}

// GetSecurityGroupNetworkInterface GET
// /security_groups/{security_group_id}/network_interfaces/{id}
// Retrieve a network interface in a security group
func GetSecurityGroupNetworkInterface(vpcService *vpcv1.VpcV1, id, vnicID string) (netInterface *vpcv1.NetworkInterface, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetSecurityGroupNetworkInterfaceOptions{}
	options.SetSecurityGroupID(id)
	options.SetID(vnicID)
	netInterface, response, err = vpcService.GetSecurityGroupNetworkInterface(options)
	return
}

// CreateSecurityGroupNetworkInterfaceBinding PUT
// /security_groups/{security_group_id}/network_interfaces/{id}
// Add a network interface to a security group
func CreateSecurityGroupNetworkInterfaceBinding(vpcService *vpcv1.VpcV1, id, vnicID string) (netInterface *vpcv1.NetworkInterface, response *core.DetailedResponse, err error) {
	options := &vpcv1.AddSecurityGroupNetworkInterfaceOptions{}
	options.SetSecurityGroupID(id)
	options.SetID(vnicID)
	netInterface, response, err = vpcService.AddSecurityGroupNetworkInterface(options)
	return
}

// ListSecurityGroupRules GET
// /security_groups/{security_group_id}/rules
// List all the rules of a security group
func ListSecurityGroupRules(vpcService *vpcv1.VpcV1, id string) (rules *vpcv1.SecurityGroupRuleCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListSecurityGroupRulesOptions{}
	options.SetSecurityGroupID(id)
	rules, response, err = vpcService.ListSecurityGroupRules(options)
	return
}

// CreateSecurityGroupRule POST
// /security_groups/{security_group_id}/rules
// Create a security group rule
func CreateSecurityGroupRule(vpcService *vpcv1.VpcV1, sgID string) (rule vpcv1.SecurityGroupRuleIntf, response *core.DetailedResponse, err error) {
	options := &vpcv1.CreateSecurityGroupRuleOptions{}
	options.SetSecurityGroupID(sgID)
	options.SetSecurityGroupRulePrototype(&vpcv1.SecurityGroupRulePrototype{
		Direction: core.StringPtr("inbound"),
		Protocol:  core.StringPtr("all"),
		IPVersion: core.StringPtr("ipv4"),
	})
	rule, response, err = vpcService.CreateSecurityGroupRule(options)
	return
}

// DeleteSecurityGroupRule DELETE
// /security_groups/{security_group_id}/rules/{id}
// Delete a security group rule
func DeleteSecurityGroupRule(vpcService *vpcv1.VpcV1, sgID, sgRuleID string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.DeleteSecurityGroupRuleOptions{}
	options.SetSecurityGroupID(sgID)
	options.SetID(sgRuleID)
	response, err = vpcService.DeleteSecurityGroupRule(options)
	return response, err
}

// GetSecurityGroupRule GET
// /security_groups/{security_group_id}/rules/{id}
// Retrieve a security group rule
func GetSecurityGroupRule(vpcService *vpcv1.VpcV1, sgID, sgRuleID string) (rule vpcv1.SecurityGroupRuleIntf, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetSecurityGroupRuleOptions{}
	options.SetSecurityGroupID(sgID)
	options.SetID(sgRuleID)
	rule, response, err = vpcService.GetSecurityGroupRule(options)
	return
}

// UpdateSecurityGroupRule PATCH
// /security_groups/{security_group_id}/rules/{id}
// Update a security group rule
func UpdateSecurityGroupRule(vpcService *vpcv1.VpcV1, sgID, sgRuleID string) (rule vpcv1.SecurityGroupRuleIntf, response *core.DetailedResponse, err error) {
	body := &vpcv1.SecurityGroupRulePatch{
		Remote: &vpcv1.SecurityGroupRuleRemotePatch{
			Address: core.StringPtr("1.1.1.11"),
		},
	}
	patchBody, _ := body.AsPatch()
	options := &vpcv1.UpdateSecurityGroupRuleOptions{
		SecurityGroupRulePatch: patchBody,
	}
	options.SetSecurityGroupID(sgID)
	options.SetID(sgRuleID)
	rule, response, err = vpcService.UpdateSecurityGroupRule(options)
	return
}

/**
 * Load Balancers
 *
 */
func ListLoadBalancerProfiles(vpcService *vpcv1.VpcV1) (profiles *vpcv1.LoadBalancerProfileCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListLoadBalancerProfilesOptions{}
	profiles, response, err = vpcService.ListLoadBalancerProfiles(options)
	return
}

func GetLoadBalancerProfile(vpcService *vpcv1.VpcV1, name string) (profile *vpcv1.LoadBalancerProfile, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetLoadBalancerProfileOptions{}
	options.SetName(name)
	profile, response, err = vpcService.GetLoadBalancerProfile(options)
	return
}

// ListLoadBalancers GET
// /load_balancers
// List all load balancers
func ListLoadBalancers(vpcService *vpcv1.VpcV1) (lbs *vpcv1.LoadBalancerCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListLoadBalancersOptions{}
	lbs, response, err = vpcService.ListLoadBalancers(options)
	return
}

// CreateLoadBalancer POST
// /load_balancers
// Create and provision a load balancer
func CreateLoadBalancer(vpcService *vpcv1.VpcV1, name, subnetID string) (lb *vpcv1.LoadBalancer, response *core.DetailedResponse, err error) {
	options := &vpcv1.CreateLoadBalancerOptions{}
	options.SetIsPublic(true)
	options.SetName(name)
	var subnetArray = []vpcv1.SubnetIdentityIntf{
		&vpcv1.SubnetIdentity{
			ID: &subnetID,
		},
	}
	options.SetSubnets(subnetArray)
	options.SetProfile(&vpcv1.LoadBalancerProfileIdentityByName{Name: core.StringPtr("network_small")})
	lb, response, err = vpcService.CreateLoadBalancer(options)
	return
}

// DeleteLoadBalancer DELETE
// /load_balancers/{id}
// Delete a load balancer
func DeleteLoadBalancer(vpcService *vpcv1.VpcV1, id string) (response *core.DetailedResponse, err error) {
	deleteVpcOptions := &vpcv1.DeleteLoadBalancerOptions{}
	deleteVpcOptions.SetID(id)
	response, err = vpcService.DeleteLoadBalancer(deleteVpcOptions)
	return response, err
}

// GetLoadBalancer GET
// /load_balancers/{id}
// Retrieve a load balancer
func GetLoadBalancer(vpcService *vpcv1.VpcV1, id string) (lb *vpcv1.LoadBalancer, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetLoadBalancerOptions{}
	options.SetID(id)
	lb, response, err = vpcService.GetLoadBalancer(options)
	return
}

// UpdateLoadBalancer PATCH
// /load_balancers/{id}
// Update a load balancer
func UpdateLoadBalancer(vpcService *vpcv1.VpcV1, id, name string) (lb *vpcv1.LoadBalancer, response *core.DetailedResponse, err error) {
	body := &vpcv1.AddressPrefixPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcv1.UpdateLoadBalancerOptions{
		LoadBalancerPatch: patchBody,
		ID:                &id,
	}
	lb, response, err = vpcService.UpdateLoadBalancer(options)
	return
}

// GetLoadBalancerStatistics GET
// /load_balancers/{id}/statistics
// List statistics of a load balancer
func GetLoadBalancerStatistics(vpcService *vpcv1.VpcV1, id string) (lb *vpcv1.LoadBalancerStatistics, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetLoadBalancerStatisticsOptions{}
	options.SetID(id)
	lb, response, err = vpcService.GetLoadBalancerStatistics(options)
	return
}

// ListLoadBalancerListeners GET
// /load_balancers/{load_balancer_id}/listeners
// List all listeners of the load balancer
func ListLoadBalancerListeners(vpcService *vpcv1.VpcV1, id string) (listeners *vpcv1.LoadBalancerListenerCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListLoadBalancerListenersOptions{}
	options.SetLoadBalancerID(id)
	listeners, response, err = vpcService.ListLoadBalancerListeners(options)
	return
}

// CreateLoadBalancerListener POST
// /load_balancers/{load_balancer_id}/listeners
// Create a listener
func CreateLoadBalancerListener(vpcService *vpcv1.VpcV1, lbID string) (listener *vpcv1.LoadBalancerListener, response *core.DetailedResponse, err error) {
	options := &vpcv1.CreateLoadBalancerListenerOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetPort(rand.Int63n(100))
	options.SetProtocol("http")
	listener, response, err = vpcService.CreateLoadBalancerListener(options)
	return
}

// DeleteLoadBalancerListener DELETE
// /load_balancers/{load_balancer_id}/listeners/{id}
// Delete a listener
func DeleteLoadBalancerListener(vpcService *vpcv1.VpcV1, lbID, listenerID string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.DeleteLoadBalancerListenerOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetID(listenerID)
	response, err = vpcService.DeleteLoadBalancerListener(options)
	return response, err
}

// GetLoadBalancerListener GET
// /load_balancers/{load_balancer_id}/listeners/{id}
// Retrieve a listener
func GetLoadBalancerListener(vpcService *vpcv1.VpcV1, lbID, listenerID string) (listener *vpcv1.LoadBalancerListener, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetLoadBalancerListenerOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetID(listenerID)
	listener, response, err = vpcService.GetLoadBalancerListener(options)
	return
}

// UpdateLoadBalancerListener PATCH
// /load_balancers/{load_balancer_id}/listeners/{id}
// Update a listener
func UpdateLoadBalancerListener(vpcService *vpcv1.VpcV1, lbID, listenerID string) (listener *vpcv1.LoadBalancerListener, response *core.DetailedResponse, err error) {
	body := &vpcv1.LoadBalancerListenerPatch{
		Protocol: core.StringPtr("http"),
	}
	patchBody, _ := body.AsPatch()
	options := &vpcv1.UpdateLoadBalancerListenerOptions{
		LoadBalancerListenerPatch: patchBody,
		LoadBalancerID:            &lbID,
		ID:                        &listenerID,
	}
	listener, response, err = vpcService.UpdateLoadBalancerListener(options)
	return
}

// ListLoadBalancerListenerPolicies GET
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies
// List all policies of the load balancer listener
func ListLoadBalancerListenerPolicies(vpcService *vpcv1.VpcV1, lbID, listenerID string) (policies *vpcv1.LoadBalancerListenerPolicyCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListLoadBalancerListenerPoliciesOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	policies, response, err = vpcService.ListLoadBalancerListenerPolicies(options)
	return
}

// CreateLoadBalancerListenerPolicy POST
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies
func CreateLoadBalancerListenerPolicy(vpcService *vpcv1.VpcV1, lbID, listenerID string) (policy *vpcv1.LoadBalancerListenerPolicy, response *core.DetailedResponse, err error) {
	options := &vpcv1.CreateLoadBalancerListenerPolicyOptions{}
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
func DeleteLoadBalancerListenerPolicy(vpcService *vpcv1.VpcV1, lbID, listenerID, policyID string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.DeleteLoadBalancerListenerPolicyOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	options.SetID(policyID)
	response, err = vpcService.DeleteLoadBalancerListenerPolicy(options)
	return response, err
}

// GetLoadBalancerListenerPolicy GET
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies/{id}
// Retrieve a policy of the load balancer listener
func GetLoadBalancerListenerPolicy(vpcService *vpcv1.VpcV1, lbID, listenerID, policyID string) (policy *vpcv1.LoadBalancerListenerPolicy, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetLoadBalancerListenerPolicyOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	options.SetID(policyID)
	policy, response, err = vpcService.GetLoadBalancerListenerPolicy(options)
	return
}

// UpdateLoadBalancerListenerPolicy PATCH
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies/{id}
// Update a policy of the load balancer listener
func UpdateLoadBalancerListenerPolicy(vpcService *vpcv1.VpcV1, lbID, listenerID, policyID, targetPoolID string) (policy *vpcv1.LoadBalancerListenerPolicy, response *core.DetailedResponse, err error) {
	target := &vpcv1.LoadBalancerListenerPolicyTargetPatch{
		ID: &targetPoolID,
	}
	body := &vpcv1.LoadBalancerListenerPolicyPatch{
		Name:     core.StringPtr("my-loadblanacer-listner-policy"),
		Priority: core.Int64Ptr(4),
		Target:   target,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcv1.UpdateLoadBalancerListenerPolicyOptions{
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
func ListLoadBalancerListenerPolicyRules(vpcService *vpcv1.VpcV1, lbID, listenerID, policyID string) (rules *vpcv1.LoadBalancerListenerPolicyRuleCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListLoadBalancerListenerPolicyRulesOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	options.SetPolicyID(policyID)
	rules, response, err = vpcService.ListLoadBalancerListenerPolicyRules(options)
	return
}

// CreateLoadBalancerListenerPolicyRule POST
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies/{policy_id}/rules
// Create a rule for the load balancer listener policy
func CreateLoadBalancerListenerPolicyRule(vpcService *vpcv1.VpcV1, lbID, listenerID, policyID string) (rule *vpcv1.LoadBalancerListenerPolicyRule, response *core.DetailedResponse, err error) {
	options := &vpcv1.CreateLoadBalancerListenerPolicyRuleOptions{}
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
func DeleteLoadBalancerListenerPolicyRule(vpcService *vpcv1.VpcV1, lbID, listenerID, policyID, ruleID string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.DeleteLoadBalancerListenerPolicyRuleOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	options.SetPolicyID(policyID)
	options.SetID(ruleID)
	response, err = vpcService.DeleteLoadBalancerListenerPolicyRule(options)
	return response, err
}

// GetLoadBalancerListenerPolicyRule GET
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies/{policy_id}/rules/{id}
// Retrieve a rule of the load balancer listener policy
func GetLoadBalancerListenerPolicyRule(vpcService *vpcv1.VpcV1, lbID, listenerID, policyID, ruleID string) (rule *vpcv1.LoadBalancerListenerPolicyRule, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetLoadBalancerListenerPolicyRuleOptions{}
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
func UpdateLoadBalancerListenerPolicyRule(vpcService *vpcv1.VpcV1, lbID, listenerID, policyID, ruleID string) (rule *vpcv1.LoadBalancerListenerPolicyRule, response *core.DetailedResponse, err error) {
	body := &vpcv1.LoadBalancerListenerPolicyRulePatch{
		Condition: core.StringPtr("equals"),
		Type:      core.StringPtr("header"),
		Value:     core.StringPtr("1"),
		Field:     core.StringPtr("field-1"),
	}
	patchBody, _ := body.AsPatch()
	options := &vpcv1.UpdateLoadBalancerListenerPolicyRuleOptions{
		LoadBalancerListenerPolicyRulePatch: patchBody,
	}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	options.SetPolicyID(policyID)
	options.SetID(ruleID)
	rule, response, err = vpcService.UpdateLoadBalancerListenerPolicyRule(options)
	return
}

// ListLoadBalancerPools GET
// /load_balancers/{load_balancer_id}/pools
// List all pools of the load balancer
func ListLoadBalancerPools(vpcService *vpcv1.VpcV1, id string) (pools *vpcv1.LoadBalancerPoolCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListLoadBalancerPoolsOptions{}
	options.SetLoadBalancerID(id)
	pools, response, err = vpcService.ListLoadBalancerPools(options)
	return
}

// CreateLoadBalancerPool POST
// /load_balancers/{load_balancer_id}/pools
// Create a load balancer pool
func CreateLoadBalancerPool(vpcService *vpcv1.VpcV1, lbID, name string) (pool *vpcv1.LoadBalancerPool, response *core.DetailedResponse, err error) {
	options := &vpcv1.CreateLoadBalancerPoolOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetAlgorithm("round_robin")
	options.SetHealthMonitor(&vpcv1.LoadBalancerPoolHealthMonitorPrototype{
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
func DeleteLoadBalancerPool(vpcService *vpcv1.VpcV1, lbID, poolID string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.DeleteLoadBalancerPoolOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetID(poolID)
	response, err = vpcService.DeleteLoadBalancerPool(options)
	return response, err
}

// GetLoadBalancerPool GET
// /load_balancers/{load_balancer_id}/pools/{id}
// Retrieve a load balancer pool
func GetLoadBalancerPool(vpcService *vpcv1.VpcV1, lbID, poolID string) (pool *vpcv1.LoadBalancerPool, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetLoadBalancerPoolOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetID(poolID)
	pool, response, err = vpcService.GetLoadBalancerPool(options)
	return
}

// UpdateLoadBalancerPool PATCH
// /load_balancers/{load_balancer_id}/pools/{id}
// Update a load balancer pool
func UpdateLoadBalancerPool(vpcService *vpcv1.VpcV1, lbID, poolID string) (pool *vpcv1.LoadBalancerPool, response *core.DetailedResponse, err error) {
	body := &vpcv1.LoadBalancerPoolPatch{
		Protocol: core.StringPtr("tcp"),
	}
	patchBody, _ := body.AsPatch()
	options := &vpcv1.UpdateLoadBalancerPoolOptions{
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
func ListLoadBalancerPoolMembers(vpcService *vpcv1.VpcV1, lbID, poolID string) (members *vpcv1.LoadBalancerPoolMemberCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListLoadBalancerPoolMembersOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetPoolID(poolID)
	members, response, err = vpcService.ListLoadBalancerPoolMembers(options)
	return
}

// CreateLoadBalancerPoolMember POST
// /load_balancers/{load_balancer_id}/pools/{pool_id}/members
// Create a member in the load balancer pool
func CreateLoadBalancerPoolMember(vpcService *vpcv1.VpcV1, lbID, poolID string) (member *vpcv1.LoadBalancerPoolMember, response *core.DetailedResponse, err error) {
	options := &vpcv1.CreateLoadBalancerPoolMemberOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetPoolID(poolID)
	options.SetPort(1234)
	options.SetTarget(&vpcv1.LoadBalancerPoolMemberTargetPrototype{
		Address: core.StringPtr("12.12.0.0"),
	})
	member, response, err = vpcService.CreateLoadBalancerPoolMember(options)
	return
}

// UpdateLoadBalancerPoolMembers PUT
// /load_balancers/{load_balancer_id}/pools/{pool_id}/members
// Update members of the load balancer pool
func UpdateLoadBalancerPoolMembers(vpcService *vpcv1.VpcV1, lbID, poolID string) (member *vpcv1.LoadBalancerPoolMemberCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ReplaceLoadBalancerPoolMembersOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetPoolID(poolID)
	options.SetMembers([]vpcv1.LoadBalancerPoolMemberPrototype{
		{
			Port: core.Int64Ptr(2345),
			Target: &vpcv1.LoadBalancerPoolMemberTargetPrototype{
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
func DeleteLoadBalancerPoolMember(vpcService *vpcv1.VpcV1, lbID, poolID, memberID string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.DeleteLoadBalancerPoolMemberOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetPoolID(poolID)
	options.SetID(memberID)
	response, err = vpcService.DeleteLoadBalancerPoolMember(options)
	return response, err
}

// GetLoadBalancerPoolMember GET
// /load_balancers/{load_balancer_id}/pools/{pool_id}/members/{id}
// Retrieve a member in the load balancer pool
func GetLoadBalancerPoolMember(vpcService *vpcv1.VpcV1, lbID, poolID, memberID string) (member *vpcv1.LoadBalancerPoolMember, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetLoadBalancerPoolMemberOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetPoolID(poolID)
	options.SetID(memberID)
	member, response, err = vpcService.GetLoadBalancerPoolMember(options)
	return
}

// UpdateLoadBalancerPoolMember PATCH
// /load_balancers/{load_balancer_id}/pools/{pool_id}/members/{id}
func UpdateLoadBalancerPoolMember(vpcService *vpcv1.VpcV1, lbID, poolID, memberID string) (member *vpcv1.LoadBalancerPoolMember, response *core.DetailedResponse, err error) {
	body := &vpcv1.LoadBalancerPoolMemberPatch{
		Port: core.Int64Ptr(3434),
	}
	patchBody, _ := body.AsPatch()
	options := &vpcv1.UpdateLoadBalancerPoolMemberOptions{
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
// List all IKE policies
func ListIkePolicies(vpcService *vpcv1.VpcV1) (ikePolicies *vpcv1.IkePolicyCollection, response *core.DetailedResponse, err error) {
	options := vpcService.NewListIkePoliciesOptions()
	ikePolicies, response, err = vpcService.ListIkePolicies(options)
	return
}

// CreateIkePolicy POST
// /ike_policies
// Create an IKE policy
func CreateIkePolicy(vpcService *vpcv1.VpcV1, name string) (ikePolicy *vpcv1.IkePolicy, response *core.DetailedResponse, err error) {
	options := &vpcv1.CreateIkePolicyOptions{}
	options.SetName(name)
	options.SetAuthenticationAlgorithm("md5")
	options.SetDhGroup(2)
	options.SetEncryptionAlgorithm("aes128")
	options.SetIkeVersion(1)
	ikePolicy, response, err = vpcService.CreateIkePolicy(options)
	return
}

// DeleteIkePolicy DELETE
// /ike_policies/{id}
// Delete an IKE policy
func DeleteIkePolicy(vpcService *vpcv1.VpcV1, id string) (response *core.DetailedResponse, err error) {
	options := vpcService.NewDeleteIkePolicyOptions(id)
	response, err = vpcService.DeleteIkePolicy(options)
	return response, err
}

// GetIkePolicy GET
// /ike_policies/{id}
// Retrieve the specified IKE policy
func GetIkePolicy(vpcService *vpcv1.VpcV1, id string) (ikePolicy *vpcv1.IkePolicy, response *core.DetailedResponse, err error) {
	options := vpcService.NewGetIkePolicyOptions(id)
	ikePolicy, response, err = vpcService.GetIkePolicy(options)
	return
}

// UpdateIkePolicy PATCH
// /ike_policies/{id}
// Update an IKE policy
func UpdateIkePolicy(vpcService *vpcv1.VpcV1, id string) (ikePolicy *vpcv1.IkePolicy, response *core.DetailedResponse, err error) {
	body := &vpcv1.IkePolicyPatch{
		Name:    core.StringPtr("go-ike-policy-2"),
		DhGroup: core.Int64Ptr(5),
	}
	patchBody, _ := body.AsPatch()
	options := &vpcv1.UpdateIkePolicyOptions{
		ID:             &id,
		IkePolicyPatch: patchBody,
	}
	ikePolicy, response, err = vpcService.UpdateIkePolicy(options)
	return
}

// ListVPNGatewayIkePolicyConnections GET
// /ike_policies/{id}/connections
// Lists all the connections that use the specified policy
func ListVPNGatewayIkePolicyConnections(vpcService *vpcv1.VpcV1, id string) (connections *vpcv1.VPNGatewayConnectionCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListIkePolicyConnectionsOptions{
		ID: &id,
	}
	connections, response, err = vpcService.ListIkePolicyConnections(options)
	return
}

// ListIpsecPolicies GET
// /ipsec_policies
// List all IPsec policies
func ListIpsecPolicies(vpcService *vpcv1.VpcV1) (ipsecPolicies *vpcv1.IPsecPolicyCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListIpsecPoliciesOptions{}
	ipsecPolicies, response, err = vpcService.ListIpsecPolicies(options)
	return
}

// CreateIpsecPolicy POST
// /ipsec_policies
// Create an IPsec policy
func CreateIpsecPolicy(vpcService *vpcv1.VpcV1, name string) (ipsecPolicy *vpcv1.IPsecPolicy, response *core.DetailedResponse, err error) {
	options := &vpcv1.CreateIpsecPolicyOptions{}
	options.SetName(name)
	options.SetAuthenticationAlgorithm("md5")
	options.SetEncryptionAlgorithm("aes128")
	options.SetPfs("disabled")
	ipsecPolicy, response, err = vpcService.CreateIpsecPolicy(options)
	return
}

// DeleteIpsecPolicy DELETE
// /ipsec_policies/{id}
// Delete an IPsec policy
func DeleteIpsecPolicy(vpcService *vpcv1.VpcV1, id string) (response *core.DetailedResponse, err error) {
	options := vpcService.NewDeleteIpsecPolicyOptions(id)
	response, err = vpcService.DeleteIpsecPolicy(options)
	return response, err
}

// GetIpsecPolicy GET
// /ipsec_policies/{id}
// Retrieve the specified IPsec policy
func GetIpsecPolicy(vpcService *vpcv1.VpcV1, id string) (ipsecPolicy *vpcv1.IPsecPolicy, response *core.DetailedResponse, err error) {
	options := vpcService.NewGetIpsecPolicyOptions(id)
	ipsecPolicy, response, err = vpcService.GetIpsecPolicy(options)
	return
}

// UpdateIpsecPolicy PATCH
// /ipsec_policies/{id}
// Update an IPsec policy
func UpdateIpsecPolicy(vpcService *vpcv1.VpcV1, id string) (ipsecPolicy *vpcv1.IPsecPolicy, response *core.DetailedResponse, err error) {
	body := &vpcv1.IPsecPolicyPatch{
		EncryptionAlgorithm: core.StringPtr("3des"),
	}
	patchBody, _ := body.AsPatch()
	options := &vpcv1.UpdateIpsecPolicyOptions{
		ID:               &id,
		IPsecPolicyPatch: patchBody,
	}
	ipsecPolicy, response, err = vpcService.UpdateIpsecPolicy(options)
	return
}

// ListVPNGatewayIpsecPolicyConnections GET
// /ipsec_policies/{id}/connections
// Lists all the connections that use the specified policy
func ListIpsecPolicyConnections(vpcService *vpcv1.VpcV1, id string) (connections *vpcv1.VPNGatewayConnectionCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListIpsecPolicyConnectionsOptions{
		ID: &id,
	}
	connections, response, err = vpcService.ListIpsecPolicyConnections(options)
	return
}

// ListVPNGateways GET
// /VPN_gateways
// List all VPN gateways
func ListVPNGateways(vpcService *vpcv1.VpcV1) (gateways *vpcv1.VPNGatewayCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListVPNGatewaysOptions{}
	gateways, response, err = vpcService.ListVPNGateways(options)
	return
}

// CreateVPNGateway POST
// /VPN_gateways
// Create a VPN gateway
func CreateVPNGateway(vpcService *vpcv1.VpcV1, subnetID, name string) (gateway vpcv1.VPNGatewayIntf, response *core.DetailedResponse, err error) {
	options := &vpcv1.CreateVPNGatewayOptions{
		VPNGatewayPrototype: &vpcv1.VPNGatewayPrototype{
			Name: &name,
			Subnet: &vpcv1.SubnetIdentity{
				ID: &subnetID,
			},
		},
	}
	gateway, response, err = vpcService.CreateVPNGateway(options)
	return
}

// DeleteVPNGateway DELETE
// /VPN_gateways/{id}
// Delete a VPN gateway
func DeleteVPNGateway(vpcService *vpcv1.VpcV1, id string) (response *core.DetailedResponse, err error) {
	options := vpcService.NewDeleteVPNGatewayOptions(id)
	response, err = vpcService.DeleteVPNGateway(options)
	return response, err
}

// GetVPNGateway GET
// /VPN_gateways/{id}
// Retrieve the specified VPN gateway
func GetVPNGateway(vpcService *vpcv1.VpcV1, id string) (gateway vpcv1.VPNGatewayIntf, response *core.DetailedResponse, err error) {
	options := vpcService.NewGetVPNGatewayOptions(id)
	gateway, response, err = vpcService.GetVPNGateway(options)
	return
}

// UpdateVPNGateway PATCH
// /VPN_gateways/{id}
// Update a VPN gateway
func UpdateVPNGateway(vpcService *vpcv1.VpcV1, id, name string) (gateway vpcv1.VPNGatewayIntf, response *core.DetailedResponse, err error) {
	body := &vpcv1.VPNGatewayPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcv1.UpdateVPNGatewayOptions{
		ID:              &id,
		VPNGatewayPatch: patchBody,
	}
	gateway, response, err = vpcService.UpdateVPNGateway(options)
	return
}

// ListVPNGatewayConnections GET
// /VPN_gateways/{VPN_gateway_id}/connections
// List all the connections of a VPN gateway
func ListVPNGatewayConnections(vpcService *vpcv1.VpcV1, gatewayID string) (connections *vpcv1.VPNGatewayConnectionCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListVPNGatewayConnectionsOptions{}
	options.SetVPNGatewayID(gatewayID)
	connections, response, err = vpcService.ListVPNGatewayConnections(options)
	return
}

// CreateVPNGatewayConnection POST
// /VPN_gateways/{VPN_gateway_id}/connections
// Create a VPN connection
func CreateVPNGatewayConnection(vpcService *vpcv1.VpcV1, gatewayID, name string) (connections vpcv1.VPNGatewayConnectionIntf, response *core.DetailedResponse, err error) {
	options := &vpcv1.CreateVPNGatewayConnectionOptions{}
	peerAddress := "192.168.0.1"
	psk := "pre-shared-key"
	local := []string{"192.132.0.0/28"}
	peer := []string{"197.155.0.0/28"}
	options.SetVPNGatewayConnectionPrototype(&vpcv1.VPNGatewayConnectionPrototype{
		Name:        &name,
		PeerAddress: &peerAddress,
		Psk:         &psk,
		LocalCIDRs:  local,
		PeerCIDRs:   peer,
	})
	options.SetVPNGatewayID(gatewayID)
	connections, response, err = vpcService.CreateVPNGatewayConnection(options)
	return
}

// DeleteVPNGatewayConnection DELETE
// /VPN_gateways/{VPN_gateway_id}/connections/{id}
// Delete a VPN connection
func DeleteVPNGatewayConnection(vpcService *vpcv1.VpcV1, gatewayID, connID string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.DeleteVPNGatewayConnectionOptions{}
	options.SetVPNGatewayID(gatewayID)
	options.SetID(connID)
	response, err = vpcService.DeleteVPNGatewayConnection(options)
	return response, err
}

// GetVPNGatewayConnection GET
// /VPN_gateways/{VPN_gateway_id}/connections/{id}
// Retrieve the specified VPN connection
func GetVPNGatewayConnection(vpcService *vpcv1.VpcV1, gatewayID, connID string) (connection vpcv1.VPNGatewayConnectionIntf, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetVPNGatewayConnectionOptions{}
	options.SetVPNGatewayID(gatewayID)
	options.SetID(connID)
	connection, response, err = vpcService.GetVPNGatewayConnection(options)
	return
}

// UpdateVPNGatewayConnection PATCH
// /VPN_gateways/{VPN_gateway_id}/connections/{id}
// Update a VPN connection
func UpdateVPNGatewayConnection(vpcService *vpcv1.VpcV1, gatewayID, connID, name string) (connection vpcv1.VPNGatewayConnectionIntf, response *core.DetailedResponse, err error) {
	body := &vpcv1.VPNGatewayConnectionPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcv1.UpdateVPNGatewayConnectionOptions{
		ID:                        &connID,
		VPNGatewayID:              &gatewayID,
		VPNGatewayConnectionPatch: patchBody,
	}
	connection, response, err = vpcService.UpdateVPNGatewayConnection(options)
	return
}

// ListVPNGatewayConnectionLocalCIDRs GET
// /VPN_gateways/{VPN_gateway_id}/connections/{id}/local_cidrs
// List all local CIDRs for a resource
func ListVPNGatewayConnectionLocalCIDRs(vpcService *vpcv1.VpcV1, gatewayID, connID string) (localCIDRs *vpcv1.VPNGatewayConnectionLocalCIDRs, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListVPNGatewayConnectionLocalCIDRsOptions{}
	options.SetVPNGatewayID(gatewayID)
	options.SetID(connID)
	localCIDRs, response, err = vpcService.ListVPNGatewayConnectionLocalCIDRs(options)
	return
}

// DeleteVPNGatewayConnectionLocalCidr DELETE
// /VPN_gateways/{VPN_gateway_id}/connections/{id}/local_cidrs/{prefix_address}/{prefix_length}
// Remove a CIDR from a resource
func DeleteVPNGatewayConnectionLocalCidr(vpcService *vpcv1.VpcV1, gatewayID, connID, prefixAdd, prefixLen string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.RemoveVPNGatewayConnectionLocalCIDROptions{}
	options.SetVPNGatewayID(gatewayID)
	options.SetID(connID)
	options.SetCIDRPrefix(prefixAdd)
	options.SetPrefixLength(prefixLen)
	response, err = vpcService.RemoveVPNGatewayConnectionLocalCIDR(options)
	return response, err
}

// GetVPNGatewayConnectionLocalCidr GET
// /VPN_gateways/{VPN_gateway_id}/connections/{id}/local_cidrs
// Check if a specific CIDR exists on a specific resource
func CheckVPNGatewayConnectionLocalCidr(vpcService *vpcv1.VpcV1, gatewayID, connID, prefixAdd, prefixLen string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.CheckVPNGatewayConnectionLocalCIDROptions{}
	options.SetVPNGatewayID(gatewayID)
	options.SetID(connID)
	options.SetCIDRPrefix(prefixAdd)
	options.SetPrefixLength(prefixLen)
	response, err = vpcService.CheckVPNGatewayConnectionLocalCIDR(options)
	return response, err
}

// SetVPNGatewayConnectionLocalCidr - PUT
// /VPN_gateways/{VPN_gateway_id}/connections/{id}/local_cidrs/{prefix_address}/{prefix_length}
// Set a CIDR on a resource
func SetVPNGatewayConnectionLocalCidr(vpcService *vpcv1.VpcV1, gatewayID, connID, prefixAdd, prefixLen string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.AddVPNGatewayConnectionLocalCIDROptions{}
	options.SetVPNGatewayID(gatewayID)
	options.SetID(connID)
	options.SetCIDRPrefix(prefixAdd)
	options.SetPrefixLength(prefixLen)
	response, err = vpcService.AddVPNGatewayConnectionLocalCIDR(options)
	return response, err
}

// ListVPNGatewayConnectionPeerCIDRs GET
// /VPN_gateways/{VPN_gateway_id}/connections/{id}/peer_cidrs
// List all peer CIDRs for a resource
func ListVPNGatewayConnectionPeerCIDRs(vpcService *vpcv1.VpcV1, gatewayID, connID string) (peerCIDRs *vpcv1.VPNGatewayConnectionPeerCIDRs, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListVPNGatewayConnectionPeerCIDRsOptions{}
	options.SetVPNGatewayID(gatewayID)
	options.SetID(connID)
	peerCIDRs, response, err = vpcService.ListVPNGatewayConnectionPeerCIDRs(options)
	return
}

// DeleteVPNGatewayConnectionPeerCidr DELETE
// /VPN_gateways/{VPN_gateway_id}/connections/{id}/peer_cidrs/{prefix_address}/{prefix_length}
// Remove a CIDR from a resource
func DeleteVPNGatewayConnectionPeerCidr(vpcService *vpcv1.VpcV1, gatewayID, connID, prefixAdd, prefixLen string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.RemoveVPNGatewayConnectionPeerCIDROptions{}
	options.SetVPNGatewayID(gatewayID)
	options.SetID(connID)
	options.SetCIDRPrefix(prefixAdd)
	options.SetPrefixLength(prefixLen)
	response, err = vpcService.RemoveVPNGatewayConnectionPeerCIDR(options)
	return response, err
}

// GetVPNGatewayConnectionPeerCidr GET
// /VPN_gateways/{VPN_gateway_id}/connections/{id}/peer_cidrs/{prefix_address}/{prefix_length}
// Check if a specific CIDR exists on a specific resource
func CheckVPNGatewayConnectionPeerCidr(vpcService *vpcv1.VpcV1, gatewayID, connID, prefixAdd, prefixLen string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.CheckVPNGatewayConnectionPeerCIDROptions{}
	options.SetVPNGatewayID(gatewayID)
	options.SetID(connID)
	options.SetCIDRPrefix(prefixAdd)
	options.SetPrefixLength(prefixLen)
	response, err = vpcService.CheckVPNGatewayConnectionPeerCIDR(options)
	return response, err
}

// SetVPNGatewayConnectionPeerCidr - PUT
// /VPN_gateways/{VPN_gateway_id}/connections/{id}/peer_cidrs/{prefix_address}/{prefix_length}
// Set a CIDR on a resource
func SetVPNGatewayConnectionPeerCidr(vpcService *vpcv1.VpcV1, gatewayID, connID, prefixAdd, prefixLen string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.AddVPNGatewayConnectionPeerCIDROptions{}
	options.SetVPNGatewayID(gatewayID)
	options.SetID(connID)
	options.SetCIDRPrefix(prefixAdd)
	options.SetPrefixLength(prefixLen)
	response, err = vpcService.AddVPNGatewayConnectionPeerCIDR(options)
	return response, err
}

// Flow Logs
// ListFlowLogCollectors - GET
// /flow_log_collectors
// List all flow log collectors
func ListFlowLogCollectors(vpcService *vpcv1.VpcV1) (flowLogs *vpcv1.FlowLogCollectorCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListFlowLogCollectorsOptions{}
	flowLogs, response, err = vpcService.ListFlowLogCollectors(options)
	return
}

// GetFlowLogCollector - GET
// /flow_log_collectors/{id}
// Retrieve the specified flow log collector
func GetFlowLogCollector(vpcService *vpcv1.VpcV1, id string) (flowLog *vpcv1.FlowLogCollector, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetFlowLogCollectorOptions{}
	options.SetID(id)
	flowLog, response, err = vpcService.GetFlowLogCollector(options)
	return
}

// DeleteFlowLogCollector DELETE
// /flow_log_collectors/{id}
// Delete specified flow_log_collector
func DeleteFlowLogCollector(vpcService *vpcv1.VpcV1, id string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.DeleteFlowLogCollectorOptions{}
	options.SetID(id)
	response, err = vpcService.DeleteFlowLogCollector(options)
	return response, err
}

// UpdateFlowLogCollector PATCH
// /flow_log_collectors/{id}
// Update specified flow log collector
func UpdateFlowLogCollector(vpcService *vpcv1.VpcV1, id, name string) (flowLog *vpcv1.FlowLogCollector, response *core.DetailedResponse, err error) {
	model := &vpcv1.FlowLogCollectorPatch{
		Name: &name,
	}
	patch, _ := model.AsPatch()
	options := &vpcv1.UpdateFlowLogCollectorOptions{
		FlowLogCollectorPatch: patch,
		ID:                    &id,
	}
	flowLog, response, err = vpcService.UpdateFlowLogCollector(options)
	return
}

func CreateFlowLogCollector(vpcService *vpcv1.VpcV1, name, bucketName, vpcId string) (flowLog *vpcv1.FlowLogCollector, response *core.DetailedResponse, err error) {

	options := &vpcv1.CreateFlowLogCollectorOptions{}
	options.SetName(name)
	options.SetTarget(&vpcv1.FlowLogCollectorTargetPrototype{
		ID: &vpcId,
	})
	options.SetStorageBucket(&vpcv1.CloudObjectStorageBucketIdentity{
		Name: &bucketName,
	})
	flowLog, response, err = vpcService.CreateFlowLogCollector(options)
	return
}

/**
 * Autoscale
 *
 */
// GET
// /instance/templates
// Get instance templates.
func ListInstanceTemplates(vpcService *vpcv1.VpcV1) (templates *vpcv1.InstanceTemplateCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListInstanceTemplatesOptions{}
	templates, response, err = vpcService.ListInstanceTemplates(options)
	return
}

// POST
// /instance/templates
// Create an instance template
func CreateInstanceTemplate(vpcService *vpcv1.VpcV1, name, imageID, profileName, zoneName, subnetID, vpcID string) (template vpcv1.InstanceTemplateIntf, response *core.DetailedResponse, err error) {
	options := &vpcv1.CreateInstanceTemplateOptions{}
	options.SetInstanceTemplatePrototype(&vpcv1.InstanceTemplatePrototype{
		Name: &name,
		Image: &vpcv1.ImageIdentity{
			ID: &imageID,
		},
		Profile: &vpcv1.InstanceProfileIdentity{
			Name: &profileName,
		},
		Zone: &vpcv1.ZoneIdentity{
			Name: &zoneName,
		},
		PrimaryNetworkInterface: &vpcv1.NetworkInterfacePrototype{
			Subnet: &vpcv1.SubnetIdentity{
				ID: &subnetID,
			},
		},
		VPC: &vpcv1.VPCIdentity{
			ID: &vpcID,
		},
	})
	template, response, err = vpcService.CreateInstanceTemplate(options)
	return
}

// DELETE
// /instance/templates/{id}
// Delete specified instance template
func DeleteInstanceTemplate(vpcService *vpcv1.VpcV1, id string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.DeleteInstanceTemplateOptions{}
	options.SetID(id)
	response, err = vpcService.DeleteInstanceTemplate(options)
	return response, err
}

// GET
// /instance/templates/{id}
// Retrieve specified instance template
func GetInstanceTemplate(vpcService *vpcv1.VpcV1, id string) (template vpcv1.InstanceTemplateIntf, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetInstanceTemplateOptions{}
	options.SetID(id)
	template, response, err = vpcService.GetInstanceTemplate(options)
	return
}

// PATCH
// /instance/templates/{id}
// Update specified instance template
func UpdateInstanceTemplate(vpcService *vpcv1.VpcV1, id, name string) (template vpcv1.InstanceTemplateIntf, response *core.DetailedResponse, err error) {
	body := &vpcv1.InstanceTemplatePatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcv1.UpdateInstanceTemplateOptions{
		InstanceTemplatePatch: patchBody,
	}
	options.SetID(id)
	template, response, err = vpcService.UpdateInstanceTemplate(options)
	return
}

// GET
// /instance_groups
// List all instance groups
func ListInstanceGroups(vpcService *vpcv1.VpcV1) (templates *vpcv1.InstanceGroupCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListInstanceGroupsOptions{}
	templates, response, err = vpcService.ListInstanceGroups(options)
	return
}

// POST
// /instance_groups
// Create an instance group
func CreateInstanceGroup(vpcService *vpcv1.VpcV1, instanceID, name, subnetID string, membership int64) (template *vpcv1.InstanceGroup, response *core.DetailedResponse, err error) {

	options := &vpcv1.CreateInstanceGroupOptions{
		InstanceTemplate: &vpcv1.InstanceTemplateIdentityByID{
			ID: &instanceID,
		},
		Subnets: []vpcv1.SubnetIdentityIntf{
			&vpcv1.SubnetIdentity{
				ID: &subnetID,
			},
		},
		Name:            &name,
		MembershipCount: &membership,
	}
	template, response, err = vpcService.CreateInstanceGroup(options)
	return
}

// DELETE
// /instance_groups/{id}
// Delete specified instance group

func DeleteInstanceGroup(vpcService *vpcv1.VpcV1, id string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.DeleteInstanceGroupOptions{}
	options.SetID(id)
	response, err = vpcService.DeleteInstanceGroup(options)
	return response, err
}

// GET
// /instance_groups/{id}
// Retrieve specified instance group
func GetInstanceGroup(vpcService *vpcv1.VpcV1, id string) (ig *vpcv1.InstanceGroup, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetInstanceGroupOptions{}
	options.SetID(id)
	ig, response, err = vpcService.GetInstanceGroup(options)
	return
}

// PATCH
// /instance_groups/{id}
// Update specified instance group
func UpdateInstanceGroup(vpcService *vpcv1.VpcV1, id, name string) (template *vpcv1.InstanceGroup, response *core.DetailedResponse, err error) {
	body := &vpcv1.InstanceGroupPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcv1.UpdateInstanceGroupOptions{
		InstanceGroupPatch: patchBody,
	}
	options.SetID(id)
	template, response, err = vpcService.UpdateInstanceGroup(options)
	return
}

// DELETE
// /instance_groups/{instance_group_id}/load_balancer
// Delete specified instance group load balancer
func DeleteInstanceGroupLoadBalancer(vpcService *vpcv1.VpcV1, id string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.DeleteInstanceGroupLoadBalancerOptions{}
	options.SetInstanceGroupID(id)
	response, err = vpcService.DeleteInstanceGroupLoadBalancer(options)
	return response, err
}

// GET
// /instance_groups/{instance_group_id}/managers
// List all managers for an instance group
func ListInstanceGroupManagers(vpcService *vpcv1.VpcV1, id string) (templates *vpcv1.InstanceGroupManagerCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListInstanceGroupManagersOptions{}
	options.SetInstanceGroupID(id)
	templates, response, err = vpcService.ListInstanceGroupManagers(options)
	return
}

// POST
// /instance_groups/{instance_group_id}/managers
// Create an instance group manager
func CreateInstanceGroupManager(vpcService *vpcv1.VpcV1, gID, name string) (manager vpcv1.InstanceGroupManagerIntf, response *core.DetailedResponse, err error) {

	options := &vpcv1.CreateInstanceGroupManagerOptions{
		InstanceGroupManagerPrototype: &vpcv1.InstanceGroupManagerPrototype{
			Name:               &name,
			ManagerType:        core.StringPtr("autoscale"),
			MaxMembershipCount: core.Int64Ptr(2),
		},
	}
	options.SetInstanceGroupID(gID)
	manager, response, err = vpcService.CreateInstanceGroupManager(options)
	return
}

// DELETE
// /instance_groups/{instance_group_id}/managers/{id}
// Delete specified instance group manager

func DeleteInstanceGroupManager(vpcService *vpcv1.VpcV1, gID, id string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.DeleteInstanceGroupManagerOptions{}
	options.SetID(id)
	options.SetInstanceGroupID(gID)
	response, err = vpcService.DeleteInstanceGroupManager(options)
	return response, err
}

// GET
// /instance_groups/{instance_group_id}/managers/{id}
// Retrieve specified instance group

func GetInstanceGroupManager(vpcService *vpcv1.VpcV1, gID, id, name string) (manager vpcv1.InstanceGroupManagerIntf, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetInstanceGroupManagerOptions{}
	options.SetID(id)
	options.SetInstanceGroupID(gID)
	manager, response, err = vpcService.GetInstanceGroupManager(options)
	return
}

// PATCH
// /instance_groups/{instance_group_id}/managers/{id}
// Update specified instance group manager
func UpdateInstanceGroupManager(vpcService *vpcv1.VpcV1, gID, id, name string) (manager vpcv1.InstanceGroupManagerIntf, response *core.DetailedResponse, err error) {
	body := &vpcv1.InstanceGroupManagerPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcv1.UpdateInstanceGroupManagerOptions{
		InstanceGroupManagerPatch: patchBody,
	}
	options.SetInstanceGroupID(gID)
	options.SetID(id)
	manager, response, err = vpcService.UpdateInstanceGroupManager(options)
	return
}

// GET
// /instance_groups/{instance_group_id}/managers/{instance_group_manager_id}/policies
// List all policies for an instance group manager
func ListInstanceGroupManagerPolicies(vpcService *vpcv1.VpcV1, gID, mID string) (policies *vpcv1.InstanceGroupManagerPolicyCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListInstanceGroupManagerPoliciesOptions{}
	options.SetInstanceGroupID(gID)
	options.SetInstanceGroupManagerID(mID)
	policies, response, err = vpcService.ListInstanceGroupManagerPolicies(options)
	return
}

// POST
// /instance_groups/{instance_group_id}/managers/{instance_group_manager_id}/policies
// Create an instance group manager policy
func CreateInstanceGroupManagerPolicy(vpcService *vpcv1.VpcV1, gID, mID, name string) (policy vpcv1.InstanceGroupManagerPolicyIntf, response *core.DetailedResponse, err error) {

	options := &vpcv1.CreateInstanceGroupManagerPolicyOptions{
		InstanceGroupManagerPolicyPrototype: &vpcv1.InstanceGroupManagerPolicyPrototype{
			Name:        &name,
			MetricType:  core.StringPtr("cpu"),
			MetricValue: core.Int64Ptr(50),
			PolicyType:  core.StringPtr("target"),
		},
	}
	options.SetInstanceGroupID(gID)
	options.SetInstanceGroupManagerID(mID)
	policy, response, err = vpcService.CreateInstanceGroupManagerPolicy(options)
	return
}

// DELETE
// /instance_groups/{instance_group_id}/managers/{instance_group_manager_id}/policies/{id}
// Delete specified instance group manager policy

func DeleteInstanceGroupManagerPolicy(vpcService *vpcv1.VpcV1, gID, mID, id string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.DeleteInstanceGroupManagerPolicyOptions{}
	options.SetID(id)
	options.SetInstanceGroupID(gID)
	options.SetInstanceGroupManagerID(mID)
	response, err = vpcService.DeleteInstanceGroupManagerPolicy(options)
	return response, err
}

// GET
// /instance_groups/{instance_group_id}/managers/{instance_group_manager_id}/policies/{id}
// Retrieve specified instance group manager policy

func GetInstanceGroupManagerPolicy(vpcService *vpcv1.VpcV1, gID, mID, id string) (template vpcv1.InstanceGroupManagerPolicyIntf, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetInstanceGroupManagerPolicyOptions{}
	options.SetID(id)
	options.SetInstanceGroupID(gID)
	options.SetInstanceGroupManagerID(mID)
	template, response, err = vpcService.GetInstanceGroupManagerPolicy(options)
	return
}

// PATCH
// /instance_groups/{instance_group_id}/managers/{instance_group_manager_id}/policies/{id}
// Update specified instance group manager policy
func UpdateInstanceGroupManagerPolicy(vpcService *vpcv1.VpcV1, igID, mID, id, name string) (template vpcv1.InstanceGroupManagerPolicyIntf, response *core.DetailedResponse, err error) {
	body := &vpcv1.InstanceGroupManagerPolicyPatch{
		MetricType:  core.StringPtr("cpu"),
		MetricValue: core.Int64Ptr(80),
	}
	patchBody, _ := body.AsPatch()
	options := &vpcv1.UpdateInstanceGroupManagerPolicyOptions{
		InstanceGroupManagerPolicyPatch: patchBody,
	}
	options.SetID(id)
	options.SetInstanceGroupID(igID)
	options.SetInstanceGroupManagerID(mID)
	template, response, err = vpcService.UpdateInstanceGroupManagerPolicy(options)
	return
}

// DELETE
// /instance_groups/{instance_group_id}/memberships
// Delete all memberships from the instance group

func DeleteInstanceGroupMemberships(vpcService *vpcv1.VpcV1, igID string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.DeleteInstanceGroupMembershipsOptions{}
	options.SetInstanceGroupID(igID)
	response, err = vpcService.DeleteInstanceGroupMemberships(options)
	return response, err
}

// GET
// /instance_groups/{instance_group_id}/memberships
// List all memberships for the instance group
func ListInstanceGroupMemberships(vpcService *vpcv1.VpcV1, igID string) (members *vpcv1.InstanceGroupMembershipCollection, response *core.DetailedResponse, err error) {
	options := &vpcv1.ListInstanceGroupMembershipsOptions{}
	options.SetInstanceGroupID(igID)
	members, response, err = vpcService.ListInstanceGroupMemberships(options)
	return
}

// DELETE
// /instance_groups/{instance_group_id}/memberships/{id}
// Delete specified instance group membership
func DeleteInstanceGroupMembership(vpcService *vpcv1.VpcV1, igID, id string) (response *core.DetailedResponse, err error) {
	options := &vpcv1.DeleteInstanceGroupMembershipOptions{}
	options.SetInstanceGroupID(igID)
	options.SetID(id)
	response, err = vpcService.DeleteInstanceGroupMembership(options)
	return response, err
}

// GET
// /instance_groups/{instance_group_id}/memberships/{id}
// Retrieve specified instance group membership
func GetInstanceGroupMembership(vpcService *vpcv1.VpcV1, igID, id string) (member *vpcv1.InstanceGroupMembership, response *core.DetailedResponse, err error) {
	options := &vpcv1.GetInstanceGroupMembershipOptions{}
	options.SetID(id)
	options.SetInstanceGroupID(igID)
	member, response, err = vpcService.GetInstanceGroupMembership(options)
	return
}

// PATCH
// /instance_groups/{instance_group_id}/memberships/{id}
// Update specified instance group membership
func UpdateInstanceGroupMembership(vpcService *vpcv1.VpcV1, igID, id, name string) (membership *vpcv1.InstanceGroupMembership, response *core.DetailedResponse, err error) {
	body := &vpcv1.InstanceGroupMembershipPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcv1.UpdateInstanceGroupMembershipOptions{
		InstanceGroupMembershipPatch: patchBody,
	}
	options.SetID(id)
	options.SetInstanceGroupID(igID)
	membership, response, err = vpcService.UpdateInstanceGroupMembership(options)
	return
}

/**
 * Endpoint Gateways
 */

// ListEndpointGateways - GET
// /endpoint_gateway
// List all Endpoint Gateways
func ListEndpointGateways(vpcService *vpcv1.VpcV1) (endpointGateways *vpcv1.EndpointGatewayCollection, response *core.DetailedResponse, err error) {
	listEndpointGatewaysOptions := vpcService.NewListEndpointGatewaysOptions()
	endpointGateways, response, err = vpcService.ListEndpointGateways(listEndpointGatewaysOptions)
	return
}

// GetEndpointGateway - GET
// /endpoint_gateway/{id}
// Retrieve the specified Endpoint Gateway
func GetEndpointGateway(vpcService *vpcv1.VpcV1, id string) (endpointGateway *vpcv1.EndpointGateway, response *core.DetailedResponse, err error) {
	options := vpcService.NewGetEndpointGatewayOptions(id)
	endpointGateway, response, err = vpcService.GetEndpointGateway(options)
	return
}

// DeleteEndpointGateway - DELETE
// /endpoint_gateway/{id}
// Delete the specified Endpoint Gateway
func DeleteEndpointGateway(vpcService *vpcv1.VpcV1, id string) (response *core.DetailedResponse, err error) {
	options := vpcService.NewDeleteEndpointGatewayOptions(id)
	response, err = vpcService.DeleteEndpointGateway(options)
	return response, err
}

// UpdateEndpointGateway - PATCH
// /endpoint_gateway/{id}
// Update the specified Endpoint Gateway
func UpdateEndpointGateway(vpcService *vpcv1.VpcV1, id, name string) (endpointGateway *vpcv1.EndpointGateway, response *core.DetailedResponse, err error) {
	body := &vpcv1.EndpointGatewayPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcv1.UpdateEndpointGatewayOptions{
		ID:                   &id,
		EndpointGatewayPatch: patchBody,
	}

	endpointGateway, response, err = vpcService.UpdateEndpointGateway(options)
	return
}

// CreateEndpointGateway - POST
// /endpoint_gateway
// Reserve a Endpoint Gateway
func CreateEndpointGateway(vpcService *vpcv1.VpcV1, vpcId string) (endpointGateway *vpcv1.EndpointGateway, response *core.DetailedResponse, err error) {
	endpointGatewayTargetPrototypeModel := &vpcv1.EndpointGatewayTargetPrototypeProviderInfrastructureServiceIdentityProviderInfrastructureServiceIdentityByName{
		ResourceType: core.StringPtr("provider_infrastructure_service"),
		Name:         core.StringPtr("ibm-ntp-server"),
	}

	vpcIdentityModel := &vpcv1.VPCIdentityByID{
		ID: &vpcId,
	}

	options := vpcService.NewCreateEndpointGatewayOptions(
		endpointGatewayTargetPrototypeModel,
		vpcIdentityModel,
	)

	endpointGateway, response, err = vpcService.CreateEndpointGateway(options)
	return
}

// GET
// /endpoint_gateways/{endpoint_gateway_id}/ips
// List all reserved IPs bound to an endpoint gateway
func ListEndpointGatewayIps(vpcService *vpcv1.VpcV1, endpointGatewayId string) (reserveIPList *vpcv1.ReservedIPCollectionEndpointGatewayContext, response *core.DetailedResponse, err error) {
	options := vpcService.NewListEndpointGatewayIpsOptions(
		endpointGatewayId,
	)
	reserveIPList, response, err = vpcService.ListEndpointGatewayIps(options)
	return
}

// DELETE
// /endpoint_gateways/{endpoint_gateway_id}/ips/{id}
// Unbind a reserved IP from an endpoint gateway
func RemoveEndpointGatewayIP(vpcService *vpcv1.VpcV1, endpointGatewayId, id string) (response *core.DetailedResponse, err error) {
	options := vpcService.NewRemoveEndpointGatewayIPOptions(endpointGatewayId, id)
	response, err = vpcService.RemoveEndpointGatewayIP(options)
	return response, err
}

// GET
// /endpoint_gateways/{endpoint_gateway_id}/ips/{id}
// Retrieve a reserved IP bound to an endpoint gateway
func GetEndpointGatewayIP(vpcService *vpcv1.VpcV1, endpointGatewayId, id string) (reserveIP *vpcv1.ReservedIP, response *core.DetailedResponse, err error) {
	options := vpcService.NewGetEndpointGatewayIPOptions(
		endpointGatewayId,
		id,
	)
	reserveIP, response, err = vpcService.GetEndpointGatewayIP(options)
	return
}

// PUT
// /endpoint_gateways/{endpoint_gateway_id}/ips/{id}
// Bind a reserved IP to an endpoint gateway
func AddEndpointGatewayIP(vpcService *vpcv1.VpcV1, endpointGatewayId, id string) (reserveIP *vpcv1.ReservedIP, response *core.DetailedResponse, err error) {
	addEndpointGatewayIPOptions := vpcService.NewAddEndpointGatewayIPOptions(
		endpointGatewayId,
		id,
	)

	reserveIP, response, err = vpcService.AddEndpointGatewayIP(addEndpointGatewayIPOptions)
	return
}

/**
 * Routing Tables
 */

// GET
// /subnets/{id}/routing_table
// Retrieve a subnet's attached routing table
func GetSubnetRoutingTable(vpcService *vpcv1.VpcV1, subnetId string) (routingTable *vpcv1.RoutingTable, response *core.DetailedResponse, err error) {
	options := vpcService.NewGetSubnetRoutingTableOptions(
		subnetId,
	)
	routingTable, response, err = vpcService.GetSubnetRoutingTable(options)
	return
}

// PUT
// /subnets/{id}/routing_table
// Attach a routing table to a subnet
func ReplaceSubnetRoutingTable(vpcService *vpcv1.VpcV1, subnetId, routingTableId string) (routingTable *vpcv1.RoutingTable, response *core.DetailedResponse, err error) {
	routingTableIdentityModel := &vpcv1.RoutingTableIdentityByID{
		ID: &routingTableId,
	}
	options := vpcService.NewReplaceSubnetRoutingTableOptions(
		subnetId,
		routingTableIdentityModel,
	)
	routingTable, response, err = vpcService.ReplaceSubnetRoutingTable(options)
	return
}

// GET
// /vpcs/{id}/default_routing_table
// Retrieve a VPC's default routing table
func GetVPCDefaultRoutingTable(vpcService *vpcv1.VpcV1, vpcId string) (defaultRoutingTable *vpcv1.DefaultRoutingTable, response *core.DetailedResponse, err error) {
	getVPCDefaultRoutingTableOptions := vpcService.NewGetVPCDefaultRoutingTableOptions(
		vpcId,
	)

	defaultRoutingTable, response, err = vpcService.GetVPCDefaultRoutingTable(getVPCDefaultRoutingTableOptions)
	return
}

// GET
// /vpcs/{vpc_id}/routing_tables
// List all routing tables for a VPC
func ListVPCRoutingTables(vpcService *vpcv1.VpcV1, vpcId string) (routingTableCollection *vpcv1.RoutingTableCollection, response *core.DetailedResponse, err error) {
	listVPCRoutingTablesOptions := vpcService.NewListVPCRoutingTablesOptions(
		vpcId,
	)

	routingTableCollection, response, err = vpcService.ListVPCRoutingTables(listVPCRoutingTablesOptions)
	return
}

// POST
// /vpcs/{vpc_id}/routing_tables
// Create a VPC routing table
func CreateVPCRoutingTable(vpcService *vpcv1.VpcV1, vpcId, name, zoneName string) (routingTable *vpcv1.RoutingTable, response *core.DetailedResponse, err error) {
	routeNextHopPrototypeModel := &vpcv1.RouteNextHopPrototypeRouteNextHopIP{
		Address: core.StringPtr("192.168.3.4"),
	}

	zoneIdentityModel := &vpcv1.ZoneIdentityByName{
		Name: &zoneName,
	}

	routePrototypeModel := &vpcv1.RoutePrototype{
		Action:      core.StringPtr("delegate"),
		Destination: core.StringPtr("192.168.3.0/24"),
		NextHop:     routeNextHopPrototypeModel,
		Zone:        zoneIdentityModel,
	}

	createVPCRoutingTableOptions := &vpcv1.CreateVPCRoutingTableOptions{
		VPCID:  &vpcId,
		Name:   &name,
		Routes: []vpcv1.RoutePrototype{*routePrototypeModel},
	}

	routingTable, response, err = vpcService.CreateVPCRoutingTable(createVPCRoutingTableOptions)
	return
}

// DELETE
// /vpcs/{vpc_id}/routing_tables/{id}
// Delete specified VPC routing table
func DeleteVPCRoutingTable(vpcService *vpcv1.VpcV1, vpcId, id string) (response *core.DetailedResponse, err error) {
	deleteVPCRoutingTableOptions := vpcService.NewDeleteVPCRoutingTableOptions(vpcId, id)
	response, err = vpcService.DeleteVPCRoutingTable(deleteVPCRoutingTableOptions)
	return response, err
}

// GET
// /vpcs/{vpc_id}/routing_tables/{id}
// Retrieve specified VPC routing table
func GetVPCRoutingTable(vpcService *vpcv1.VpcV1, vpcId, id string) (routingTable *vpcv1.RoutingTable, response *core.DetailedResponse, err error) {
	getVPCRoutingTableOptions := vpcService.NewGetVPCRoutingTableOptions(vpcId, id)
	routingTable, response, err = vpcService.GetVPCRoutingTable(getVPCRoutingTableOptions)
	return
}

// PATCH
// /vpcs/{vpc_id}/routing_tables/{id}
// Update specified VPC routing table
func UpdateVPCRoutingTable(vpcService *vpcv1.VpcV1, vpcId, id, name string) (routingTable *vpcv1.RoutingTable, response *core.DetailedResponse, err error) {
	routingTablePatchModel := &vpcv1.RoutingTablePatch{
		Name: &name,
	}
	routingTablePatchModelAsPatch, _ := routingTablePatchModel.AsPatch()

	updateVPCRoutingTableOptions := &vpcv1.UpdateVPCRoutingTableOptions{
		VPCID:             &vpcId,
		ID:                &id,
		RoutingTablePatch: routingTablePatchModelAsPatch,
	}

	routingTable, response, err = vpcService.UpdateVPCRoutingTable(updateVPCRoutingTableOptions)
	return
}

// GET
// /vpcs/{vpc_id}/routing_tables/{routing_table_id}/routes
// List all the routes of a VPC routing table
func ListVPCRoutingTableRoutes(vpcService *vpcv1.VpcV1, vpcId, routingTableID string) (routeCollection *vpcv1.RouteCollection, response *core.DetailedResponse, err error) {
	listVPCRoutingTableRoutesOptions := &vpcv1.ListVPCRoutingTableRoutesOptions{
		VPCID:          &vpcId,
		RoutingTableID: &routingTableID,
	}

	routeCollection, response, err = vpcService.ListVPCRoutingTableRoutes(listVPCRoutingTableRoutesOptions)
	return
}

// POST
// /vpcs/{vpc_id}/routing_tables/{routing_table_id}/routes
// Create a VPC route
func CreateVPCRoutingTableRoute(vpcService *vpcv1.VpcV1, vpcId, routingTableId, zone string) (route *vpcv1.Route, response *core.DetailedResponse, err error) {
	routeNextHopPrototypeModel := &vpcv1.RouteNextHopPrototypeRouteNextHopIP{
		Address: core.StringPtr("192.168.3.4"),
	}

	zoneIdentityModel := &vpcv1.ZoneIdentityByName{
		Name: &zone,
	}

	createVPCRoutingTableRouteOptions := &vpcv1.CreateVPCRoutingTableRouteOptions{
		VPCID:          &vpcId,
		RoutingTableID: &routingTableId,
		Destination:    core.StringPtr("192.168.3.0/24"),
		NextHop:        routeNextHopPrototypeModel,
		Zone:           zoneIdentityModel,
		Action:         core.StringPtr("delegate"),
	}

	route, response, err = vpcService.CreateVPCRoutingTableRoute(createVPCRoutingTableRouteOptions)
	return
}

// DELETE
// /vpcs/{vpc_id}/routing_tables/{routing_table_id}/routes/{id}
// Delete the specified VPC route
func DeleteVPCRoutingTableRoute(vpcService *vpcv1.VpcV1, vpcId, routingTableId, id string) (response *core.DetailedResponse, err error) {
	deleteVPCRoutingTableRouteOptions := &vpcv1.DeleteVPCRoutingTableRouteOptions{
		VPCID:          &vpcId,
		RoutingTableID: &routingTableId,
		ID:             &id,
	}

	response, err = vpcService.DeleteVPCRoutingTableRoute(deleteVPCRoutingTableRouteOptions)
	return response, err
}

// GET
// /vpcs/{vpc_id}/routing_tables/{routing_table_id}/routes/{id}
// Retrieve the specified VPC route
func GetVPCRoutingTableRoute(vpcService *vpcv1.VpcV1, vpcId, routingTableId, id string) (route *vpcv1.Route, response *core.DetailedResponse, err error) {
	getVPCRoutingTableRouteOptions := &vpcv1.GetVPCRoutingTableRouteOptions{
		VPCID:          &vpcId,
		RoutingTableID: &routingTableId,
		ID:             &id,
	}
	route, response, err = vpcService.GetVPCRoutingTableRoute(getVPCRoutingTableRouteOptions)
	return
}

// PATCH
// /vpcs/{vpc_id}/routing_tables/{routing_table_id}/routes/{id}
// Update the specified VPC route
func UpdateVPCRoutingTableRoute(vpcService *vpcv1.VpcV1, vpcId, routingTableId, id, name string) (route *vpcv1.Route, response *core.DetailedResponse, err error) {
	routePatchModel := &vpcv1.RoutePatch{
		Name: &name,
	}
	routePatchModelAsPatch, _ := routePatchModel.AsPatch()

	updateVPCRoutingTableRouteOptions := &vpcv1.UpdateVPCRoutingTableRouteOptions{
		VPCID:          &vpcId,
		RoutingTableID: &routingTableId,
		ID:             &id,
		RoutePatch:     routePatchModelAsPatch,
	}

	route, response, err = vpcService.UpdateVPCRoutingTableRoute(updateVPCRoutingTableRouteOptions)
	return
}

/**
 * Subnet Reserved IPs
 *
 */

// GET
// /subnets/{subnet_id}/reserved_ips
// List all reserved IPs in a subnet
func ListSubnetReservedIps(vpcService *vpcv1.VpcV1, subnetId string) (reservedIPCollection *vpcv1.ReservedIPCollection, response *core.DetailedResponse, err error) {
	listSubnetReservedIpsOptions := vpcService.NewListSubnetReservedIpsOptions(
		subnetId,
	)
	reservedIPCollection, response, err = vpcService.ListSubnetReservedIps(listSubnetReservedIpsOptions)
	return
}

// POST
// /subnets/{subnet_id}/reserved_ips
// Reserve an IP in a subnet
func CreateSubnetReservedIP(vpcService *vpcv1.VpcV1, subnetId, name string) (reservedIP *vpcv1.ReservedIP, response *core.DetailedResponse, err error) {
	createSubnetReservedIPOptions := &vpcv1.CreateSubnetReservedIPOptions{
		SubnetID: &subnetId,
		Name:     &name,
	}
	reservedIP, response, err = vpcService.CreateSubnetReservedIP(createSubnetReservedIPOptions)
	return
}

// DELETE
// /subnets/{subnet_id}/reserved_ips/{id}
// Release specified reserved IP
func DeleteSubnetReservedIP(vpcService *vpcv1.VpcV1, subnetId, reservedIPId string) (response *core.DetailedResponse, err error) {
	deleteSubnetReservedIPOptions := vpcService.NewDeleteSubnetReservedIPOptions(
		subnetId,
		reservedIPId,
	)

	response, err = vpcService.DeleteSubnetReservedIP(deleteSubnetReservedIPOptions)
	return response, err
}

// GET
// /subnets/{subnet_id}/reserved_ips/{id}
// Retrieve specified reserved IP
func GetSubnetReservedIP(vpcService *vpcv1.VpcV1, subnetId, reservedIPId string) (reservedIP *vpcv1.ReservedIP, response *core.DetailedResponse, err error) {
	getSubnetReservedIPOptions := vpcService.NewGetSubnetReservedIPOptions(
		subnetId,
		reservedIPId,
	)
	reservedIP, response, err = vpcService.GetSubnetReservedIP(getSubnetReservedIPOptions)
	return
}

// PATCH
// /subnets/{subnet_id}/reserved_ips/{id}
// Update specified reserved IP
func UpdateSubnetReservedIP(vpcService *vpcv1.VpcV1, subnetId, reservedIPId, name string) (reservedIP *vpcv1.ReservedIP, response *core.DetailedResponse, err error) {
	reservedIPPatchModel := &vpcv1.ReservedIPPatch{
		Name: &name,
	}
	reservedIPPatchModelAsPatch, _ := reservedIPPatchModel.AsPatch()

	updateSubnetReservedIPOptions := vpcService.NewUpdateSubnetReservedIPOptions(
		subnetId,
		reservedIPId,
		reservedIPPatchModelAsPatch,
	)
	reservedIP, response, err = vpcService.UpdateSubnetReservedIP(updateSubnetReservedIPOptions)
	return
}

// GET
// /dedicated_host/groups
// List all dedicated host groups
func ListDedicatedHostGroups(vpcService *vpcv1.VpcV1) (dedicatedHostGroupCollection *vpcv1.DedicatedHostGroupCollection, response *core.DetailedResponse, err error) {
	listDedicatedHostGroupsOptions := &vpcv1.ListDedicatedHostGroupsOptions{}

	dedicatedHostGroupCollection, response, err = vpcService.ListDedicatedHostGroups(listDedicatedHostGroupsOptions)
	return
}

// POST
// /dedicated_host/groups
// Create a dedicated host group
func CreateDedicatedHostGroup(vpcService *vpcv1.VpcV1, name, zone string) (dedicatedHostGroup *vpcv1.DedicatedHostGroup, response *core.DetailedResponse, err error) {
	zoneIdentityModel := &vpcv1.ZoneIdentityByName{
		Name: &zone,
	}
	createDedicatedHostGroupOptions := &vpcv1.CreateDedicatedHostGroupOptions{
		Class:  core.StringPtr("mx2"),
		Family: core.StringPtr("balanced"),
		Name:   &name,
		Zone:   zoneIdentityModel,
	}

	dedicatedHostGroup, response, err = vpcService.CreateDedicatedHostGroup(createDedicatedHostGroupOptions)
	return
}

// DELETE
// /dedicated_host/groups/{id}
// Delete specified dedicated host group
func DeleteDedicatedHostGroup(vpcService *vpcv1.VpcV1, id *string) (response *core.DetailedResponse, err error) {
	getDedicatedHostGroupOptions := &vpcv1.DeleteDedicatedHostGroupOptions{
		ID: id,
	}
	response, err = vpcService.DeleteDedicatedHostGroup(getDedicatedHostGroupOptions)
	return
}

// GET
// /dedicated_host/groups/{id}
// Retrieve a dedicated host group
func GetDedicatedHostGroup(vpcService *vpcv1.VpcV1, id *string) (dedicatedHostGroup *vpcv1.DedicatedHostGroup, response *core.DetailedResponse, err error) {
	getDedicatedHostGroupOptions := &vpcv1.GetDedicatedHostGroupOptions{
		ID: id,
	}

	dedicatedHostGroup, response, err = vpcService.GetDedicatedHostGroup(getDedicatedHostGroupOptions)
	return
}

// PATCH
// /dedicated_host/groups/{id}
// Update specified dedicated host group
func UpdateDedicatedHostGroup(vpcService *vpcv1.VpcV1, id, name *string) (dedicatedHostGroup *vpcv1.DedicatedHostGroup, response *core.DetailedResponse, err error) {
	model := &vpcv1.DedicatedHostGroupPatch{
		Name: name,
	}
	patch, _ := model.AsPatch()
	getDedicatedHostGroupOptions := &vpcv1.UpdateDedicatedHostGroupOptions{
		ID:                      id,
		DedicatedHostGroupPatch: patch,
	}

	dedicatedHostGroup, response, err = vpcService.UpdateDedicatedHostGroup(getDedicatedHostGroupOptions)
	return
}

// GET
// /dedicated_host/profiles
// List all dedicated host profiles
func ListDedicatedHostProfiles(vpcService *vpcv1.VpcV1) (result *vpcv1.DedicatedHostProfileCollection, response *core.DetailedResponse, err error) {
	listDedicatedHostProfilesOptions := &vpcv1.ListDedicatedHostProfilesOptions{}

	result, response, err = vpcService.ListDedicatedHostProfiles(listDedicatedHostProfilesOptions)
	return
}

// GET
// /dedicated_host/profiles/{name}
// Retrieve specified dedicated host profile
func GetDedicatedHostProfile(vpcService *vpcv1.VpcV1, name *string) (result *vpcv1.DedicatedHostProfile, response *core.DetailedResponse, err error) {
	getDedicatedHostProfileOptions := &vpcv1.GetDedicatedHostProfileOptions{
		Name: name,
	}
	result, response, err = vpcService.GetDedicatedHostProfile(getDedicatedHostProfileOptions)
	return
}

// GET
// /dedicated_hosts
// List all dedicated hosts
func ListDedicatedHosts(vpcService *vpcv1.VpcV1) (dedicatedHostCollection *vpcv1.DedicatedHostCollection, response *core.DetailedResponse, err error) {
	listDedicatedHostsOptions := &vpcv1.ListDedicatedHostsOptions{}

	dedicatedHostCollection, response, err = vpcService.ListDedicatedHosts(listDedicatedHostsOptions)
	return
}

// POST
// /dedicated_hosts
// Create a dedicated host
func CreateDedicatedHost(vpcService *vpcv1.VpcV1, name, profile, groupID *string) (dedicatedHost *vpcv1.DedicatedHost, response *core.DetailedResponse, err error) {
	fmt.Println(" sd", *name, *groupID)
	dedicatedHostProfileIdentityModel := &vpcv1.DedicatedHostProfileIdentityByName{
		Name: profile,
	}

	dedicatedHostGroupIdentityModel := &vpcv1.DedicatedHostGroupIdentityByID{
		ID: groupID,
	}

	dedicatedHostPrototypeModel := &vpcv1.DedicatedHostPrototypeDedicatedHostByGroup{
		Name:    name,
		Profile: dedicatedHostProfileIdentityModel,
		Group:   dedicatedHostGroupIdentityModel,
	}

	createDedicatedHostOptions := &vpcv1.CreateDedicatedHostOptions{
		DedicatedHostPrototype: dedicatedHostPrototypeModel,
	}

	dedicatedHost, response, err = vpcService.CreateDedicatedHost(createDedicatedHostOptions)
	return
}

// DELETE
// /dedicated_hosts/{id}
// Delete specified dedicated host
func DeleteDedicatedHost(vpcService *vpcv1.VpcV1, id *string) (response *core.DetailedResponse, err error) {
	getDedicatedHostOptions := &vpcv1.DeleteDedicatedHostOptions{
		ID: id,
	}
	response, err = vpcService.DeleteDedicatedHost(getDedicatedHostOptions)
	return
}

// GET
// /dedicated_hosts/{id}
// Retrieve a dedicated host
func GetDedicatedHost(vpcService *vpcv1.VpcV1, id string) (dedicatedHost *vpcv1.DedicatedHost, response *core.DetailedResponse, err error) {
	getDedicatedHostOptions := &vpcv1.GetDedicatedHostOptions{
		ID: &id,
	}

	dedicatedHost, response, err = vpcService.GetDedicatedHost(getDedicatedHostOptions)
	return
}

// PATCH
// /dedicated_hosts/{id}
// Update specified dedicated host
func UpdateDedicatedHost(vpcService *vpcv1.VpcV1, name, id *string) (dedicatedHost *vpcv1.DedicatedHost, response *core.DetailedResponse, err error) {
	model := &vpcv1.DedicatedHostPatch{
		Name: name,
	}
	patch, _ := model.AsPatch()
	getDedicatedHostOptions := &vpcv1.UpdateDedicatedHostOptions{
		ID:                 id,
		DedicatedHostPatch: patch,
	}

	dedicatedHost, response, err = vpcService.UpdateDedicatedHost(getDedicatedHostOptions)
	return
}

// Print - Marshal JSON and print
func Print(printObject interface{}) {
	p, _ := json.MarshalIndent(printObject, "", "\t")
	fmt.Println(string(p))
}

// PollInstance - poll and check the status of VSI before performing any action
// ID - resource ID
// status - expected status/ poll until this status is returned
// pollFrequency - number of times polling happens
func PollInstance(vpcService *vpcv1.VpcV1, ID, status string, pollFrequency int) bool {
	count := 1
	for {
		if count < pollFrequency {
			res, _, err := GetInstance(vpcService, ID)
			fmt.Println("Current status of VSI - ", *res.Status)
			fmt.Println("Expected status of VSI - ", status)
			if err != nil && res == nil {
				fmt.Printf("err error: Retrieving instance ID %s with err error message: %s", ID, err)
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
func PollSubnet(vpcService *vpcv1.VpcV1, ID, status string, pollFrequency int) bool {
	count := 1
	for {
		if count < pollFrequency {
			res, _, err := GetSubnet(vpcService, ID)
			fmt.Println("Current status of Subnet - ", *res.Status)
			fmt.Println("Expected status of Subnet - ", status)
			if err != nil && res == nil {
				fmt.Printf("err error: Retrieving subnet ID %s with err error message: %s", ID, err)
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
func PollVolAttachment(vpcService *vpcv1.VpcV1, vpcID, volAttachmentID, status string, pollFrequency int) bool {
	count := 1
	for {
		if count < pollFrequency {
			res, _, err := GetVolumeAttachment(vpcService, vpcID, volAttachmentID)
			fmt.Println("Current status of attachment - ", *res.Status)
			fmt.Println("Expected status of attachment - ", status)
			if err != nil && res == nil {
				fmt.Printf("err error: Retrieving volume attachment ID %s with err error message: %s", vpcID, err)
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
func PollLB(vpcService *vpcv1.VpcV1, lbID, status string, pollFrequency int) bool {
	count := 1
	for {
		if count < pollFrequency {
			res, _, err := GetLoadBalancer(vpcService, lbID)
			fmt.Println("Current status of load balancer - ", *res.ProvisioningStatus)
			fmt.Println("Expected status of load balancer - ", status)
			if err != nil && res == nil {
				fmt.Printf("err error: Retrieving load balancer ID %s with err error message: %s", lbID, err)
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
func PollVPNGateway(vpcService *vpcv1.VpcV1, gatewayID, status string, pollFrequency int) bool {
	count := 1
	for {
		if count < pollFrequency {
			res, _, err := GetVPNGateway(vpcService, gatewayID)
			vpn := res.(*vpcv1.VPNGateway)
			fmt.Println("Current status of VPNGateway - ", *vpn.Status)
			fmt.Println("Expected status of VPNGateway - ", status)
			if err != nil && vpn == nil {
				fmt.Printf("err error: Retrieving VPNGateway ID %s with err error message: %s", gatewayID, err)
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
