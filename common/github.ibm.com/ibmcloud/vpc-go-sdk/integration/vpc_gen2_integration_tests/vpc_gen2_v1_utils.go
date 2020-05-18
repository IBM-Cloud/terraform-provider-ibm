package vpcgen2integration

import (
	"fmt"
	"math/rand"

	"github.com/IBM/go-sdk-core/v3/core"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
)

// InstantiateVPCGen2Service - Instantiate VPC Gen2 service
func InstantiateVPCGen2Service() *vpcv1.VpcV1 {

	gen2Service, gen2ServiceErr := vpcv1.NewVpcV1(&vpcv1.VpcV1Options{
		URL: URL,
		Authenticator: &core.IamAuthenticator{
			ApiKey: APIKey,
			URL:    IAMURL,
		},
	})
	// Check successful instantiation
	if gen2ServiceErr != nil {
		fmt.Println("service error")
		fmt.Println(gen2ServiceErr)
		return nil
	}
	// return new vpc gen2 service
	return gen2Service
}

/**
 * Regions and Zones
 *
 */

// ListRegions - List all regions
// GET
// /regions
func ListRegions(gen2 *vpcv1.VpcV1) (*vpcv1.RegionCollection, *core.DetailedResponse, error) {
	listRegionsOptions := &vpcv1.ListRegionsOptions{}
	result, returnValue, returnValueErr := gen2.ListRegions(listRegionsOptions)
	return result, returnValue, returnValueErr
}

// GetRegion - GET
// /regions/{name}
// Retrieve a region
func GetRegion(vpcService *vpcv1.VpcV1, name string) (*vpcv1.Region, *core.DetailedResponse, error) {
	getRegionOptions := &vpcv1.GetRegionOptions{}
	getRegionOptions.SetName(name)
	result, returnValue, returnValueErr := vpcService.GetRegion(getRegionOptions)
	return result, returnValue, returnValueErr
}

// ListZones - GET
// /regions/{region_name}/zones
// List all zones in a region
func ListZones(vpcService *vpcv1.VpcV1, regionName string) (*vpcv1.ZoneCollection, *core.DetailedResponse, error) {
	listZonesOptions := &vpcv1.ListZonesOptions{}
	listZonesOptions.SetRegionName(regionName)
	result, returnValue, returnValueErr := vpcService.ListZones(listZonesOptions)
	return result, returnValue, returnValueErr
}

// GetZone - GET
// /regions/{region_name}/zones/{zone_name}
// Retrieve a zone
func GetZone(vpcService *vpcv1.VpcV1, regionName, zoneName string) (*vpcv1.Zone, *core.DetailedResponse, error) {
	getZoneOptions := &vpcv1.GetZoneOptions{}
	getZoneOptions.SetRegionName(regionName)
	getZoneOptions.SetZoneName(zoneName)
	result, returnValue, returnValueErr := vpcService.GetZone(getZoneOptions)
	return result, returnValue, returnValueErr
}

/**
 * Floating IPs
 */

// GetFloatingIPsList - GET
// /floating_ips
// List all floating IPs
func GetFloatingIPsList(vpcService *vpcv1.VpcV1) (*vpcv1.FloatingIPCollection, *core.DetailedResponse, error) {
	listFloatingIpsOptions := vpcService.NewListFloatingIpsOptions()
	result, returnValue, returnValueErr := vpcService.ListFloatingIps(listFloatingIpsOptions)
	// TODO: target is not coming back in the response
	return result, returnValue, returnValueErr
}

// GetFloatingIP - GET
// /floating_ips/{id}
// Retrieve the specified floating IP
func GetFloatingIP(vpcService *vpcv1.VpcV1, id string) (*vpcv1.FloatingIP, *core.DetailedResponse, error) {
	options := vpcService.NewGetFloatingIpOptions(id)
	result, returnValue, returnValueErr := vpcService.GetFloatingIp(options)
	return result, returnValue, returnValueErr
}

// ReleaseFloatingIP - DELETE
// /floating_ips/{id}
// Release the specified floating IP
func ReleaseFloatingIP(vpcService *vpcv1.VpcV1, id string) (*core.DetailedResponse, error) {
	options := vpcService.NewReleaseFloatingIpOptions(id)
	returnValue, returnValueErr := vpcService.ReleaseFloatingIp(options)
	return returnValue, returnValueErr
}

// UpdateFloatingIP - PATCH
// /floating_ips/{id}
// Update the specified floating IP
func UpdateFloatingIP(vpcService *vpcv1.VpcV1, id, name string) (*vpcv1.FloatingIP, *core.DetailedResponse, error) {
	options := &vpcv1.UpdateFloatingIpOptions{
		ID:   core.StringPtr(id),
		Name: core.StringPtr(name),
	}
	// To update target
	// updateFloatingIpOptions.SetTarget(&vpcv1.NetworkInterfaceIdentity{
	// 	ID: core.StringPtr(targetId),
	// })
	result, returnValue, returnValueErr := vpcService.UpdateFloatingIp(options)
	return result, returnValue, returnValueErr
}

// CreateFloatingIP - POST
// /floating_ips
// Reserve a floating IP
func CreateFloatingIP(vpcService *vpcv1.VpcV1, zone, name string) (*vpcv1.FloatingIP, *core.DetailedResponse, error) {
	options := &vpcv1.ReserveFloatingIpOptions{}
	options.SetFloatingIPPrototype(&vpcv1.FloatingIPPrototype{
		Name: core.StringPtr(name),
		Zone: &vpcv1.ZoneIdentity{
			Name: core.StringPtr(zone),
		},
	})
	result, returnValue, returnValueErr := vpcService.ReserveFloatingIp(options)
	return result, returnValue, returnValueErr
}

/**
 * SSH Keys
 *
 */

// ListKeys - GET
// /keys
// List all keys
func ListKeys(vpcService *vpcv1.VpcV1) (*vpcv1.KeyCollection, *core.DetailedResponse, error) {
	listKeysOptions := &vpcv1.ListKeysOptions{}
	result, returnValue, returnValueErr := vpcService.ListKeys(listKeysOptions)
	return result, returnValue, returnValueErr
}

// GetSSHKey - GET
// /keys/{id}
// Retrieve specified key
func GetSSHKey(vpcService *vpcv1.VpcV1, id string) (*vpcv1.Key, *core.DetailedResponse, error) {
	getKeyOptions := &vpcv1.GetKeyOptions{}
	getKeyOptions.SetID(id)
	result, returnValue, returnValueErr := vpcService.GetKey(getKeyOptions)
	return result, returnValue, returnValueErr
}

// UpdateSSHKey - PATCH
// /keys/{id}
// Update specified key
func UpdateSSHKey(vpcService *vpcv1.VpcV1, id, name string) (*vpcv1.Key, *core.DetailedResponse, error) {
	updateKeyOptions := &vpcv1.UpdateKeyOptions{}
	updateKeyOptions.SetID(id)
	updateKeyOptions.SetName(name)
	result, returnValue, returnValueErr := vpcService.UpdateKey(updateKeyOptions)
	return result, returnValue, returnValueErr
}

// DeleteSSHKey - DELETE
// /keys/{id}
// Delete specified key
func DeleteSSHKey(vpcService *vpcv1.VpcV1, id string) (*core.DetailedResponse, error) {
	deleteKeyOptions := &vpcv1.DeleteKeyOptions{}
	deleteKeyOptions.SetID(id)
	returnValue, returnValueErr := vpcService.DeleteKey(deleteKeyOptions)
	return returnValue, returnValueErr
}

// CreateSSHKey - POST
// /keys
// Create a key
func CreateSSHKey(vpcService *vpcv1.VpcV1, name, publicKey string) (*vpcv1.Key, *core.DetailedResponse, error) {
	options := &vpcv1.CreateKeyOptions{}

	options.SetName(name)
	options.SetPublicKey(publicKey)
	result, returnValue, returnValueErr := vpcService.CreateKey(options)
	return result, returnValue, returnValueErr
}

/**
 * VPC
 *
 */

// GetVPCsList - GET
// /vpcs
// List all VPCs
func GetVPCsList(vpcService *vpcv1.VpcV1) (*vpcv1.VPCCollection, *core.DetailedResponse, error) {
	listVpcsOptions := &vpcv1.ListVpcsOptions{}
	result, returnValue, returnValueErr := vpcService.ListVpcs(listVpcsOptions)
	return result, returnValue, returnValueErr
}

// GetVPC - GET
// /vpcs/{id}
// Retrieve specified VPC
func GetVPC(vpcService *vpcv1.VpcV1, id string) (*vpcv1.VPC, *core.DetailedResponse, error) {
	getVpcOptions := &vpcv1.GetVpcOptions{}
	getVpcOptions.SetID(id)
	result, returnValue, returnValueErr := vpcService.GetVpc(getVpcOptions)
	return result, returnValue, returnValueErr
}

// DeleteVPC - DELETE
// /vpcs/{id}
// Delete specified VPC
func DeleteVPC(vpcService *vpcv1.VpcV1, id string) (*core.DetailedResponse, error) {
	deleteVpcOptions := &vpcv1.DeleteVpcOptions{}
	deleteVpcOptions.SetID(id)
	returnValue, returnValueErr := vpcService.DeleteVpc(deleteVpcOptions)
	return returnValue, returnValueErr
}

// UpdateVPC - PATCH
// /vpcs/{id}
// Update specified VPC
func UpdateVPC(vpcService *vpcv1.VpcV1, id, name string) (*vpcv1.VPC, *core.DetailedResponse, error) {
	options := &vpcv1.UpdateVpcOptions{
		Name: core.StringPtr(name),
	}
	options.SetID(id)
	result, returnValue, returnValueErr := vpcService.UpdateVpc(options)
	return result, returnValue, returnValueErr
}

// CreateVPC - POST
// /vpcs
// Create a VPC
func CreateVPC(vpcService *vpcv1.VpcV1, name, resourceGroup string) (*vpcv1.VPC, *core.DetailedResponse, error) {
	options := &vpcv1.CreateVpcOptions{}

	options.SetResourceGroup(&vpcv1.ResourceGroupIdentity{
		ID: core.StringPtr(resourceGroup),
	})
	options.SetName(name)
	result, returnValue, returnValueErr := vpcService.CreateVpc(options)
	return result, returnValue, returnValueErr
}

/**
 * VPC default Security group
 * Getting default security group for a vpc with id
 */

// GetVPCDefaultSecurityGroup - GET
// /vpcs/{id}/default_security_group
// Retrieve a VPC's default security group
func GetVPCDefaultSecurityGroup(vpcService *vpcv1.VpcV1, id string) (*vpcv1.DefaultSecurityGroup, *core.DetailedResponse, error) {
	getVpcDefaultSecurityGroupOptions := &vpcv1.GetVpcDefaultSecurityGroupOptions{}
	getVpcDefaultSecurityGroupOptions.SetID(id)
	result, returnValue, returnValueErr := vpcService.GetVpcDefaultSecurityGroup(getVpcDefaultSecurityGroupOptions)
	return result, returnValue, returnValueErr
}

/**
 * VPC address prefix
 *
 */

// ListVpcAddressPrefixes - GET
// /vpcs/{vpc_id}/address_prefixes
// List all address pool prefixes for a VPC
func ListVpcAddressPrefixes(vpcService *vpcv1.VpcV1, vpcID string) (*vpcv1.AddressPrefixCollection, *core.DetailedResponse, error) {
	listVpcAddressPrefixesOptions := &vpcv1.ListVpcAddressPrefixesOptions{}
	listVpcAddressPrefixesOptions.SetVpcID(vpcID)
	result, returnValue, returnValueErr := vpcService.ListVpcAddressPrefixes(listVpcAddressPrefixesOptions)
	return result, returnValue, returnValueErr
}

// GetVpcAddressPrefix - GET
// /vpcs/{vpc_id}/address_prefixes/{id}
// Retrieve specified address pool prefix
func GetVpcAddressPrefix(vpcService *vpcv1.VpcV1, vpcID, addressPrefixID string) (*vpcv1.AddressPrefix, *core.DetailedResponse, error) {
	getVpcAddressPrefixOptions := &vpcv1.GetVpcAddressPrefixOptions{}
	getVpcAddressPrefixOptions.SetVpcID(vpcID)
	getVpcAddressPrefixOptions.SetID(addressPrefixID)
	result, returnValue, returnValueErr := vpcService.GetVpcAddressPrefix(getVpcAddressPrefixOptions)
	return result, returnValue, returnValueErr
}

// CreateVpcAddressPrefix - POST
// /vpcs/{vpc_id}/address_prefixes
// Create an address pool prefix
func CreateVpcAddressPrefix(vpcService *vpcv1.VpcV1, vpcID, zone, cidr, name string) (*vpcv1.AddressPrefix, *core.DetailedResponse, error) {
	options := &vpcv1.CreateVpcAddressPrefixOptions{}
	options.SetVpcID(vpcID)
	options.SetCidr(cidr)
	options.SetName(name)
	options.SetZone(&vpcv1.ZoneIdentity{
		Name: core.StringPtr(zone),
	})
	result, returnValue, returnValueErr := vpcService.CreateVpcAddressPrefix(options)
	return result, returnValue, returnValueErr
}

// DeleteVpcAddressPrefix - DELETE
// /vpcs/{vpc_id}/address_prefixes/{id}
// Delete specified address pool prefix
func DeleteVpcAddressPrefix(vpcService *vpcv1.VpcV1, vpcID, addressPrefixID string) (*core.DetailedResponse, error) {
	deleteVpcAddressPrefixOptions := &vpcv1.DeleteVpcAddressPrefixOptions{}
	deleteVpcAddressPrefixOptions.SetVpcID(vpcID)
	deleteVpcAddressPrefixOptions.SetID(addressPrefixID)
	returnValue, returnValueErr := vpcService.DeleteVpcAddressPrefix(deleteVpcAddressPrefixOptions)
	return returnValue, returnValueErr
}

// UpdateVpcAddressPrefix - PATCH
// /vpcs/{vpc_id}/address_prefixes/{id}
// Update an address pool prefix
func UpdateVpcAddressPrefix(vpcService *vpcv1.VpcV1, vpcID, addressPrefixID, name string) (*vpcv1.AddressPrefix, *core.DetailedResponse, error) {
	options := &vpcv1.UpdateVpcAddressPrefixOptions{}

	options.SetVpcID(vpcID)
	options.SetID(addressPrefixID)
	options.SetName(name)
	result, returnValue, returnValueErr := vpcService.UpdateVpcAddressPrefix(options)
	return result, returnValue, returnValueErr
}

/**
 * VPC routes
 *
 */

// ListVpcRoutes - GET
// /vpcs/{vpc_id}/routes
// List all user-defined routes for a VPC
func ListVpcRoutes(vpcService *vpcv1.VpcV1, vpcID string) (*vpcv1.RouteCollection, *core.DetailedResponse, error) {
	options := &vpcv1.ListVpcRoutesOptions{}

	options.SetVpcID(vpcID)
	result, returnValue, returnValueErr := vpcService.ListVpcRoutes(options)
	return result, returnValue, returnValueErr
}

// GetVpcRoute - GET
// /vpcs/{vpc_id}/routes/{id}
// Retrieve the specified route
func GetVpcRoute(vpcService *vpcv1.VpcV1, vpcID, routeID string) (*vpcv1.Route, *core.DetailedResponse, error) {
	options := &vpcv1.GetVpcRouteOptions{}

	options.SetVpcID(vpcID)
	options.SetID(routeID)
	result, returnValue, returnValueErr := vpcService.GetVpcRoute(options)
	return result, returnValue, returnValueErr
}

// CreateVpcRoute - POST
// /vpcs/{vpc_id}/routes
// Create a route on your VPC
func CreateVpcRoute(vpcService *vpcv1.VpcV1, vpcID, zone, destination, nextHopAddress, name string) (*vpcv1.Route, *core.DetailedResponse, error) {
	options := &vpcv1.CreateVpcRouteOptions{}

	options.SetVpcID(vpcID)
	options.SetName(name)
	options.SetZone(&vpcv1.ZoneIdentity{
		Name: core.StringPtr(zone),
	})
	options.SetNextHop(&vpcv1.RouteNextHopPrototype{
		Address: &nextHopAddress,
	})
	options.SetDestination(destination)
	result, returnValue, returnValueErr := vpcService.CreateVpcRoute(options)
	return result, returnValue, returnValueErr
}

// DeleteVpcRoute - DELETE
// /vpcs/{vpc_id}/routes/{id}
// Delete the specified route
func DeleteVpcRoute(vpcService *vpcv1.VpcV1, vpcID, routeID string) (*core.DetailedResponse, error) {
	options := &vpcv1.DeleteVpcRouteOptions{}
	options.SetVpcID(vpcID)
	options.SetID(routeID)
	returnValue, returnValueErr := vpcService.DeleteVpcRoute(options)
	return returnValue, returnValueErr
}

// UpdateVpcRoute - PATCH
// /vpcs/{vpc_id}/routes/{id}
// Update a route
func UpdateVpcRoute(vpcService *vpcv1.VpcV1, vpcID, routeID, name string) (*vpcv1.Route, *core.DetailedResponse, error) {
	options := &vpcv1.UpdateVpcRouteOptions{}

	options.SetVpcID(vpcID)
	options.SetID(routeID)
	options.SetName(name)
	result, returnValue, returnValueErr := vpcService.UpdateVpcRoute(options)
	return result, returnValue, returnValueErr
}

/**
 * Volumes
 *
 */

// ListVolumeProfiles - GET
// /volume/profiles
// List all volume profiles
func ListVolumeProfiles(vpcService *vpcv1.VpcV1) (*vpcv1.VolumeProfileCollection, *core.DetailedResponse, error) {
	options := &vpcv1.ListVolumeProfilesOptions{}

	result, returnValue, returnValueErr := vpcService.ListVolumeProfiles(options)
	return result, returnValue, returnValueErr
}

// GetVolumeProfile - GET
// /volume/profiles/{name}
// Retrieve specified volume profile
func GetVolumeProfile(vpcService *vpcv1.VpcV1, profileName string) (*vpcv1.VolumeProfile, *core.DetailedResponse, error) {
	options := &vpcv1.GetVolumeProfileOptions{}
	options.SetName(profileName)
	result, returnValue, returnValueErr := vpcService.GetVolumeProfile(options)
	return result, returnValue, returnValueErr
}

// ListVolumes - GET
// /volumes
// List all volumes
func ListVolumes(vpcService *vpcv1.VpcV1) (*vpcv1.VolumeCollection, *core.DetailedResponse, error) {
	options := &vpcv1.ListVolumesOptions{}

	result, returnValue, returnValueErr := vpcService.ListVolumes(options)
	return result, returnValue, returnValueErr
}

// GetVolume - GET
// /volumes/{id}
// Retrieve specified volume
func GetVolume(vpcService *vpcv1.VpcV1, volumeID string) (*vpcv1.Volume, *core.DetailedResponse, error) {
	options := &vpcv1.GetVolumeOptions{}

	options.SetID(volumeID)
	result, returnValue, returnValueErr := vpcService.GetVolume(options)
	return result, returnValue, returnValueErr
}

// DeleteVolume - DELETE
// /volumes/{id}
// Delete specified volume
func DeleteVolume(vpcService *vpcv1.VpcV1, id string) (*core.DetailedResponse, error) {
	options := &vpcv1.DeleteVolumeOptions{}

	options.SetID(id)
	returnValue, returnValueErr := vpcService.DeleteVolume(options)
	return returnValue, returnValueErr
}

// UpdateVolume - PATCH
// /volumes/{id}
// Update specified volume
func UpdateVolume(vpcService *vpcv1.VpcV1, id, name string) (*vpcv1.Volume, *core.DetailedResponse, error) {
	options := &vpcv1.UpdateVolumeOptions{}

	options.SetID(id)
	options.SetName(name)
	result, returnValue, returnValueErr := vpcService.UpdateVolume(options)
	return result, returnValue, returnValueErr
}

// CreateVolume - POST
// /volumes
// Create a volume
func CreateVolume(vpcService *vpcv1.VpcV1, name, profileName, zoneName string, capacity int64) (*vpcv1.Volume, *core.DetailedResponse, error) {
	options := &vpcv1.CreateVolumeOptions{}
	options.SetVolumePrototype(&vpcv1.VolumePrototype{
		Capacity: core.Int64Ptr(capacity),
		Zone: &vpcv1.ZoneIdentity{
			Name: core.StringPtr(zoneName),
		},
		Profile: &vpcv1.VolumeProfileIdentity{
			Name: core.StringPtr(profileName),
		},
		Name: core.StringPtr(name),
	})
	result, returnValue, returnValueErr := vpcService.CreateVolume(options)
	return result, returnValue, returnValueErr
}

/**
 * Subnets
 *
 */

// ListSubnets - GET
// /subnets
// List all subnets
func ListSubnets(vpcService *vpcv1.VpcV1) (*vpcv1.SubnetCollection, *core.DetailedResponse, error) {
	options := &vpcv1.ListSubnetsOptions{}

	result, returnValue, returnValueErr := vpcService.ListSubnets(options)
	return result, returnValue, returnValueErr
}

// GetSubnet - GET
// /subnets/{id}
// Retrieve specified subnet
func GetSubnet(vpcService *vpcv1.VpcV1, subnetID string) (*vpcv1.Subnet, *core.DetailedResponse, error) {
	options := &vpcv1.GetSubnetOptions{}

	options.SetID(subnetID)
	result, returnValue, returnValueErr := vpcService.GetSubnet(options)
	return result, returnValue, returnValueErr
}

// DeleteSubnet - DELETE
// /subnets/{id}
// Delete specified subnet
func DeleteSubnet(vpcService *vpcv1.VpcV1, id string) (*core.DetailedResponse, error) {
	options := &vpcv1.DeleteSubnetOptions{}

	options.SetID(id)
	returnValue, returnValueErr := vpcService.DeleteSubnet(options)
	return returnValue, returnValueErr
}

// UpdateSubnet - PATCH
// /subnets/{id}
// Update specified subnet
func UpdateSubnet(vpcService *vpcv1.VpcV1, id, name string) (*vpcv1.Subnet, *core.DetailedResponse, error) {
	options := &vpcv1.UpdateSubnetOptions{}

	options.SetID(id)
	options.SetName(name)
	result, returnValue, returnValueErr := vpcService.UpdateSubnet(options)
	return result, returnValue, returnValueErr
}

// CreateSubnet - POST
// /subnets
// Create a subnet
func CreateSubnet(vpcService *vpcv1.VpcV1, vpcID, name, zone string) (*vpcv1.Subnet, *core.DetailedResponse, error) {
	options := &vpcv1.CreateSubnetOptions{}

	options.SetSubnetPrototype(&vpcv1.SubnetPrototype{
		Ipv4CidrBlock: core.StringPtr("10.243.0.0/24"),
		Name:          core.StringPtr(name),
		Vpc: &vpcv1.VPCIdentity{
			ID: core.StringPtr(vpcID),
		},
		Zone: &vpcv1.ZoneIdentity{
			Name: core.StringPtr(zone),
		},
	})
	result, returnValue, returnValueErr := vpcService.CreateSubnet(options)
	return result, returnValue, returnValueErr
}

// GetSubnetNetworkAcl -GET
// /subnets/{id}/network_acl
// Retrieve a subnet's attached network ACL
func GetSubnetNetworkAcl(vpcService *vpcv1.VpcV1, subnetID string) (*vpcv1.NetworkACL, *core.DetailedResponse, error) {
	options := &vpcv1.GetSubnetNetworkAclOptions{}
	options.SetID(subnetID)
	result, returnValue, returnValueErr := vpcService.GetSubnetNetworkAcl(options)
	return result, returnValue, returnValueErr
}

// SetSubnetNetworkAclBinding - PUT
// /subnets/{id}/network_acl
// Attach a network ACL to a subnet
func SetSubnetNetworkAclBinding(vpcService *vpcv1.VpcV1, subnetID, id string) (*vpcv1.NetworkACL, *core.DetailedResponse, error) {
	options := &vpcv1.SetSubnetNetworkAclBindingOptions{}
	options.SetID(subnetID)
	options.SetNetworkACLIdentity(&vpcv1.NetworkACLIdentity{ID: &id})
	result, returnValue, returnValueErr := vpcService.SetSubnetNetworkAclBinding(options)
	return result, returnValue, returnValueErr
}

// DeleteSubnetPublicGatewayBinding - DELETE
// /subnets/{id}/public_gateway
// Detach a public gateway from a subnet
func DeleteSubnetPublicGatewayBinding(vpcService *vpcv1.VpcV1, id string) (*core.DetailedResponse, error) {
	options := &vpcv1.DeleteSubnetPublicGatewayBindingOptions{}
	options.SetID(id)
	returnValue, returnValueErr := vpcService.DeleteSubnetPublicGatewayBinding(options)
	return returnValue, returnValueErr
}

// GetSubnetPublicGateway - GET
// /subnets/{id}/public_gateway
// Retrieve a subnet's attached public gateway
func GetSubnetPublicGateway(vpcService *vpcv1.VpcV1, subnetID string) (*vpcv1.PublicGateway, *core.DetailedResponse, error) {
	options := &vpcv1.GetSubnetPublicGatewayOptions{}
	options.SetID(subnetID)
	result, returnValue, returnValueErr := vpcService.GetSubnetPublicGateway(options)
	return result, returnValue, returnValueErr
}

// SetSubnetPublicGatewayBinding - PUT
// /subnets/{id}/public_gateway
// Attach a public gateway to a subnet
func SetSubnetPublicGatewayBinding(vpcService *vpcv1.VpcV1, subnetID, id string) (*vpcv1.PublicGateway, *core.DetailedResponse, error) {
	options := &vpcv1.SetSubnetPublicGatewayBindingOptions{}
	options.SetID(subnetID)
	options.SetPublicGatewayIdentity(&vpcv1.PublicGatewayIdentity{ID: &id})
	result, returnValue, returnValueErr := vpcService.SetSubnetPublicGatewayBinding(options)
	return result, returnValue, returnValueErr
}

/**
 * Images
 *
 */

// ListImages - GET
// /images
// List all images
func ListImages(vpcService *vpcv1.VpcV1, visibility string) (*vpcv1.ImageCollection, *core.DetailedResponse, error) {
	options := &vpcv1.ListImagesOptions{}
	options.SetVisibility(visibility)
	result, returnValue, returnValueErr := vpcService.ListImages(options)
	return result, returnValue, returnValueErr
}

// GetImage - GET
// /images/{id}
// Retrieve the specified image
func GetImage(vpcService *vpcv1.VpcV1, imageID string) (*vpcv1.Image, *core.DetailedResponse, error) {
	options := &vpcv1.GetImageOptions{}
	options.SetID(imageID)
	result, returnValue, returnValueErr := vpcService.GetImage(options)
	return result, returnValue, returnValueErr
}

// DeleteImage DELETE
// /images/{id}
// Delete specified image
func DeleteImage(vpcService *vpcv1.VpcV1, id string) (*core.DetailedResponse, error) {
	options := &vpcv1.DeleteImageOptions{}
	options.SetID(id)
	returnValue, returnValueErr := vpcService.DeleteImage(options)
	return returnValue, returnValueErr
}

// UpdateImage PATCH
// /images/{id}
// Update specified image
func UpdateImage(vpcService *vpcv1.VpcV1, id, name string) (*vpcv1.Image, *core.DetailedResponse, error) {
	options := &vpcv1.UpdateImageOptions{}
	options.SetID(id)
	options.SetName(name)
	result, returnValue, returnValueErr := vpcService.UpdateImage(options)
	return result, returnValue, returnValueErr
}

// TODO
// func CreateImage(vpcService *vpcv1.VpcV1, vpcId, name, cidr string) (*vpcv1.Image,*core.DetailedResponse, error) {
// 	options := &vpcv1.CreateSubnetOptions{}
//
// 	options.SetCreateSubnetRequest(&vpcv1.CreateSubnetRequest{
// 		Ipv4CidrBlock: core.StringPtr(cidr),
// 		Name:          core.StringPtr(name),
// 		Vpc: &vpcv1.VPCIdentity{
// 			ID: core.StringPtr(vpcId),
// 		},
// 	})
// 	returnValue, returnValueErr := vpcService.CreateImage(options)
// 	return result, returnValueErr
// }

func ListOperatingSystems(vpcService *vpcv1.VpcV1) (*vpcv1.OperatingSystemCollection, *core.DetailedResponse, error) {
	options := &vpcv1.ListOperatingSystemsOptions{}
	result, returnValue, returnValueErr := vpcService.ListOperatingSystems(options)
	return result, returnValue, returnValueErr
}

func GetOperatingSystem(vpcService *vpcv1.VpcV1, osName string) (*vpcv1.OperatingSystem, *core.DetailedResponse, error) {
	options := &vpcv1.GetOperatingSystemOptions{}
	options.SetName(osName)
	result, returnValue, returnValueErr := vpcService.GetOperatingSystem(options)
	return result, returnValue, returnValueErr
}

/**
 * Instances
 *
 */

// ListInstanceProfiles - GET
// /instance/profiles
// List all instance profiles
func ListInstanceProfiles(vpcService *vpcv1.VpcV1) (*vpcv1.InstanceProfileCollection, *core.DetailedResponse, error) {
	options := &vpcv1.ListInstanceProfilesOptions{}
	result, returnValue, returnValueErr := vpcService.ListInstanceProfiles(options)
	return result, returnValue, returnValueErr
}

// GetInstanceProfile - GET
// /instance/profiles/{name}
// Retrieve specified instance profile
func GetInstanceProfile(vpcService *vpcv1.VpcV1, profileName string) (*vpcv1.InstanceProfile, *core.DetailedResponse, error) {
	options := &vpcv1.GetInstanceProfileOptions{}
	options.SetName(profileName)
	result, returnValue, returnValueErr := vpcService.GetInstanceProfile(options)
	return result, returnValue, returnValueErr
}

// ListInstances GET
// /instances
// List all instances
func ListInstances(vpcService *vpcv1.VpcV1) (*vpcv1.InstanceCollection, *core.DetailedResponse, error) {
	options := &vpcv1.ListInstancesOptions{}
	result, returnValue, returnValueErr := vpcService.ListInstances(options)
	return result, returnValue, returnValueErr
}

// GetInstance GET
// instances/{id}
// Retrieve an instance
func GetInstance(vpcService *vpcv1.VpcV1, instanceID string) (*vpcv1.Instance, *core.DetailedResponse, error) {
	options := &vpcv1.GetInstanceOptions{}
	options.SetID(instanceID)
	result, returnValue, returnValueErr := vpcService.GetInstance(options)
	return result, returnValue, returnValueErr
}

// DeleteInstance DELETE
// /instances/{id}
// Delete specified instance
func DeleteInstance(vpcService *vpcv1.VpcV1, id string) (*core.DetailedResponse, error) {
	options := &vpcv1.DeleteInstanceOptions{}
	options.SetID(id)
	returnValue, returnValueErr := vpcService.DeleteInstance(options)
	return returnValue, returnValueErr
}

// UpdateInstance PATCH
// /instances/{id}
// Update specified instance
func UpdateInstance(vpcService *vpcv1.VpcV1, id, name string) (*vpcv1.Instance, *core.DetailedResponse, error) {
	options := &vpcv1.UpdateInstanceOptions{}
	options.SetID(id)
	options.SetName(name)
	result, returnValue, returnValueErr := vpcService.UpdateInstance(options)
	return result, returnValue, returnValueErr
}

// CreateInstance POST
// /instances/{instance_id}/actions
// Create an instance action
func CreateInstance(vpcService *vpcv1.VpcV1, name, profileName, imageID, zoneName, subnetID, sshkeyID, vpcID string) (*vpcv1.Instance, *core.DetailedResponse, error) {
	options := &vpcv1.CreateInstanceOptions{}
	options.SetInstancePrototype(&vpcv1.InstancePrototype{
		Name: core.StringPtr(name),
		Image: &vpcv1.ImageIdentity{
			ID: core.StringPtr(imageID),
		},
		Profile: &vpcv1.InstanceProfileIdentity{
			Name: core.StringPtr(profileName),
		},
		Zone: &vpcv1.ZoneIdentity{
			Name: core.StringPtr(zoneName),
		},
		PrimaryNetworkInterface: &vpcv1.NetworkInterfacePrototype{
			Subnet: &vpcv1.SubnetIdentity{
				ID: core.StringPtr(subnetID),
			},
		},
		Keys: []vpcv1.KeyIdentityIntf{
			&vpcv1.KeyIdentity{
				ID: core.StringPtr(sshkeyID),
			},
		},
		Vpc: &vpcv1.VPCIdentity{
			ID: core.StringPtr(vpcID),
		},
	})
	result, returnValue, returnValueErr := vpcService.CreateInstance(options)
	return result, returnValue, returnValueErr
}

// CreateInstanceAction PATCH
// /instances/{id}
// Update specified instance
func CreateInstanceAction(vpcService *vpcv1.VpcV1, instanceID, typeOfAction string) (*vpcv1.InstanceAction, *core.DetailedResponse, error) {
	options := &vpcv1.CreateInstanceActionOptions{}
	options.SetInstanceID(instanceID)
	options.SetType(typeOfAction)
	result, returnValue, returnValueErr := vpcService.CreateInstanceAction(options)
	return result, returnValue, returnValueErr
}

// GetInstanceInitialization GET
// /instances/{id}/initialization
// Retrieve configuration used to initialize the instance.
func GetInstanceInitialization(vpcService *vpcv1.VpcV1, instanceID string) (*vpcv1.InstanceInitialization, *core.DetailedResponse, error) {
	options := &vpcv1.GetInstanceInitializationOptions{}
	options.SetID(instanceID)
	result, returnValue, returnValueErr := vpcService.GetInstanceInitialization(options)
	return result, returnValue, returnValueErr
}

// ListNetworkInterfaces GET
// /instances/{instance_id}/network_interfaces
// List all network interfaces on an instance
func ListNetworkInterfaces(vpcService *vpcv1.VpcV1, id string) (*vpcv1.NetworkInterfaceCollection, *core.DetailedResponse, error) {
	options := &vpcv1.ListNetworkInterfacesOptions{}
	options.SetInstanceID(id)
	result, returnValue, returnValueErr := vpcService.ListNetworkInterfaces(options)
	return result, returnValue, returnValueErr
}

// GetNetworkInterface GET
// /instances/{instance_id}/network_interfaces/{id}
// Retrieve specified network interface
func GetNetworkInterface(vpcService *vpcv1.VpcV1, instanceID, networkID string) (*vpcv1.NetworkInterface, *core.DetailedResponse, error) {
	options := &vpcv1.GetNetworkInterfaceOptions{}
	options.SetID(networkID)
	options.SetInstanceID(instanceID)
	result, returnValue, returnValueErr := vpcService.GetNetworkInterface(options)
	return result, returnValue, returnValueErr
}

// UpdateNetworkInterface PATCH
// /instances/{instance_id}/network_interfaces/{id}
// Update a network interface
func UpdateNetworkInterface(vpcService *vpcv1.VpcV1, instanceID, networkID, name string) (*vpcv1.NetworkInterface, *core.DetailedResponse, error) {
	options := &vpcv1.UpdateNetworkInterfaceOptions{}
	options.SetID(networkID)
	options.SetInstanceID(instanceID)
	options.SetName(name)
	result, returnValue, returnValueErr := vpcService.UpdateNetworkInterface(options)
	return result, returnValue, returnValueErr
}

// ListNetworkInterfaceFloatingIps GET
// /instances/{instance_id}/network_interfaces
// List all network interfaces on an instance
func ListNetworkInterfaceFloatingIps(vpcService *vpcv1.VpcV1, instanceID, networkID string) (*vpcv1.FloatingIPUnpaginatedCollection, *core.DetailedResponse, error) {
	options := &vpcv1.ListNetworkInterfaceFloatingIpsOptions{}
	options.SetInstanceID(instanceID)
	options.SetNetworkInterfaceID(networkID)
	result, returnValue, returnValueErr := vpcService.ListNetworkInterfaceFloatingIps(options)
	return result, returnValue, returnValueErr
}

// GetNetworkInterfaceFloatingIp GET
// /instances/{instance_id}/network_interfaces/{network_interface_id}/floating_ips
// List all floating IPs associated with a network interface
func GetNetworkInterfaceFloatingIp(vpcService *vpcv1.VpcV1, instanceID, networkID, fipID string) (*vpcv1.FloatingIP, *core.DetailedResponse, error) {
	options := &vpcv1.GetNetworkInterfaceFloatingIpOptions{}
	options.SetID(fipID)
	options.SetInstanceID(instanceID)
	options.SetNetworkInterfaceID(networkID)
	result, returnValue, returnValueErr := vpcService.GetNetworkInterfaceFloatingIp(options)
	return result, returnValue, returnValueErr
}

// DeleteNetworkInterfaceFloatingIpBinding DELETE
// /instances/{instance_id}/network_interfaces/{network_interface_id}/floating_ips/{id}
// Disassociate specified floating IP
func DeleteNetworkInterfaceFloatingIpBinding(vpcService *vpcv1.VpcV1, instanceID, networkID, fipID string) (*core.DetailedResponse, error) {
	options := &vpcv1.DeleteNetworkInterfaceFloatingIpBindingOptions{}
	options.SetID(fipID)
	options.SetInstanceID(instanceID)
	options.SetNetworkInterfaceID(networkID)
	returnValue, returnValueErr := vpcService.DeleteNetworkInterfaceFloatingIpBinding(options)
	return returnValue, returnValueErr
}

// CreateNetworkInterfaceFloatingIpBinding PUT
// /instances/{instance_id}/network_interfaces/{network_interface_id}/floating_ips/{id}
// Associate a floating IP with a network interface
func CreateNetworkInterfaceFloatingIpBinding(vpcService *vpcv1.VpcV1, instanceID, networkID, fipID string) (*vpcv1.FloatingIP, *core.DetailedResponse, error) {
	options := &vpcv1.CreateNetworkInterfaceFloatingIpBindingOptions{}
	options.SetID(fipID)
	options.SetInstanceID(instanceID)
	options.SetNetworkInterfaceID(networkID)
	result, returnValue, returnValueErr := vpcService.CreateNetworkInterfaceFloatingIpBinding(options)
	return result, returnValue, returnValueErr
}

// ListVolumeAttachments GET
// /instances/{instance_id}/volume_attachments
// List all volumes attached to an instance
func ListVolumeAttachments(vpcService *vpcv1.VpcV1, id string) (*vpcv1.VolumeAttachmentCollection, *core.DetailedResponse, error) {
	options := &vpcv1.ListVolumeAttachmentsOptions{}
	options.SetInstanceID(id)
	result, returnValue, returnValueErr := vpcService.ListVolumeAttachments(options)
	return result, returnValue, returnValueErr
}

// CreateVolumeAttachment POST
// /instances/{instance_id}/volume_attachments
// Create a volume attachment, connecting a volume to an instance
func CreateVolumeAttachment(vpcService *vpcv1.VpcV1, instanceID, volumeID, name string) (*vpcv1.VolumeAttachment, *core.DetailedResponse, error) {
	options := &vpcv1.CreateVolumeAttachmentOptions{}
	options.SetInstanceID(instanceID)
	options.SetVolume(&vpcv1.VolumeIdentity{
		ID: core.StringPtr(volumeID),
	})
	options.SetName(name)
	options.SetDeleteVolumeOnInstanceDelete(false)
	result, returnValue, returnValueErr := vpcService.CreateVolumeAttachment(options)
	return result, returnValue, returnValueErr
}

// DeleteVolumeAttachment DELETE
// /instances/{instance_id}/volume_attachments/{id}
// Delete a volume attachment, detaching a volume from an instance
func DeleteVolumeAttachment(vpcService *vpcv1.VpcV1, instanceID, volumeID string) (*core.DetailedResponse, error) {
	options := &vpcv1.DeleteVolumeAttachmentOptions{}
	options.SetID(volumeID)
	options.SetInstanceID(instanceID)
	returnValue, returnValueErr := vpcService.DeleteVolumeAttachment(options)
	return returnValue, returnValueErr
}

// GetVolumeAttachment GET
// /instances/{instance_id}/volume_attachments/{id}
// Retrieve specified volume attachment
func GetVolumeAttachment(vpcService *vpcv1.VpcV1, instanceID, volumeID string) (*vpcv1.VolumeAttachment, *core.DetailedResponse, error) {
	options := &vpcv1.GetVolumeAttachmentOptions{}
	options.SetInstanceID(instanceID)
	options.SetID(volumeID)
	result, returnValue, returnValueErr := vpcService.GetVolumeAttachment(options)
	return result, returnValue, returnValueErr
}

// UpdateVolumeAttachment PATCH
// /instances/{instance_id}/volume_attachments/{id}
// Update a volume attachment
func UpdateVolumeAttachment(vpcService *vpcv1.VpcV1, instanceID, volumeID, name string) (*vpcv1.VolumeAttachment, *core.DetailedResponse, error) {
	options := &vpcv1.UpdateVolumeAttachmentOptions{}
	options.SetInstanceID(instanceID)
	options.SetID(volumeID)
	options.SetName(name)
	result, returnValue, returnValueErr := vpcService.UpdateVolumeAttachment(options)
	return result, returnValue, returnValueErr
}

/**
 * Public Gateway
 *
 */

// ListPublicGateways GET
// /public_gateways
// List all public gateways
func ListPublicGateways(vpcService *vpcv1.VpcV1) (*vpcv1.PublicGatewayCollection, *core.DetailedResponse, error) {
	options := &vpcv1.ListPublicGatewaysOptions{}
	result, returnValue, returnValueErr := vpcService.ListPublicGateways(options)
	return result, returnValue, returnValueErr
}

// CreatePublicGateway POST
// /public_gateways
// Create a public gateway
func CreatePublicGateway(vpcService *vpcv1.VpcV1, name, vpcID, zoneName string) (*vpcv1.PublicGateway, *core.DetailedResponse, error) {
	options := &vpcv1.CreatePublicGatewayOptions{}
	options.SetVpc(&vpcv1.VPCIdentity{
		ID: core.StringPtr(vpcID),
	})
	options.SetZone(&vpcv1.ZoneIdentity{
		Name: core.StringPtr(zoneName),
	})
	result, returnValue, returnValueErr := vpcService.CreatePublicGateway(options)
	return result, returnValue, returnValueErr
}

// DeletePublicGateway DELETE
// /public_gateways/{id}
// Delete specified public gateway
func DeletePublicGateway(vpcService *vpcv1.VpcV1, id string) (*core.DetailedResponse, error) {
	options := &vpcv1.DeletePublicGatewayOptions{}
	options.SetID(id)
	returnValue, returnValueErr := vpcService.DeletePublicGateway(options)
	return returnValue, returnValueErr
}

// GetPublicGateway GET
// /public_gateways/{id}
// Retrieve specified public gateway
func GetPublicGateway(vpcService *vpcv1.VpcV1, id string) (*vpcv1.PublicGateway, *core.DetailedResponse, error) {
	options := &vpcv1.GetPublicGatewayOptions{}
	options.SetID(id)
	result, returnValue, returnValueErr := vpcService.GetPublicGateway(options)
	return result, returnValue, returnValueErr
}

// UpdatePublicGateway PATCH
// /public_gateways/{id}
// Update a public gateway's name
func UpdatePublicGateway(vpcService *vpcv1.VpcV1, id, name string) (*vpcv1.PublicGateway, *core.DetailedResponse, error) {
	options := &vpcv1.UpdatePublicGatewayOptions{}
	options.SetID(id)
	options.SetName(name)
	result, returnValue, returnValueErr := vpcService.UpdatePublicGateway(options)
	return result, returnValue, returnValueErr
}

/**
 * Network ACLs not available in gen2 currently
 *
 */

// ListNetworkAcls - GET
// /network_acls
// List all network ACLs
func ListNetworkAcls(vpcService *vpcv1.VpcV1) (*vpcv1.NetworkACLCollection, *core.DetailedResponse, error) {
	options := &vpcv1.ListNetworkAclsOptions{}
	result, returnValue, returnValueErr := vpcService.ListNetworkAcls(options)
	return result, returnValue, returnValueErr
}

// CreateNetworkAcl - POST
// /network_acls
// Create a network ACL
func CreateNetworkAcl(vpcService *vpcv1.VpcV1, name, copyableAclID, vpcID string) (*vpcv1.NetworkACL, *core.DetailedResponse, error) {
	options := &vpcv1.CreateNetworkAclOptions{}
	options.SetNetworkACLPrototype(&vpcv1.NetworkACLPrototype{
		Name: core.StringPtr(name),
		SourceNetworkAcl: &vpcv1.NetworkACLIdentity{
			ID: core.StringPtr(copyableAclID),
		},
		Vpc: &vpcv1.VPCIdentity{
			ID: core.StringPtr(vpcID),
		},
	})
	result, returnValue, returnValueErr := vpcService.CreateNetworkAcl(options)
	return result, returnValue, returnValueErr
}

// DeleteNetworkAcl - DELETE
// /network_acls/{id}
// Delete specified network ACL
func DeleteNetworkAcl(vpcService *vpcv1.VpcV1, ID string) (*core.DetailedResponse, error) {
	options := &vpcv1.DeleteNetworkAclOptions{}
	options.SetID(ID)
	returnValue, returnValueErr := vpcService.DeleteNetworkAcl(options)
	return returnValue, returnValueErr
}

// GetNetworkAcl - GET
// /network_acls/{id}
// Retrieve specified network ACL
func GetNetworkAcl(vpcService *vpcv1.VpcV1, ID string) (*vpcv1.NetworkACL, *core.DetailedResponse, error) {
	options := &vpcv1.GetNetworkAclOptions{}
	options.SetID(ID)
	result, returnValue, returnValueErr := vpcService.GetNetworkAcl(options)
	return result, returnValue, returnValueErr
}

// UpdateNetworkAcl PATCH
// /network_acls/{id}
// Update a network ACL
func UpdateNetworkAcl(vpcService *vpcv1.VpcV1, id, name string) (*vpcv1.NetworkACL, *core.DetailedResponse, error) {
	options := &vpcv1.UpdateNetworkAclOptions{}
	options.SetID(id)
	options.SetName(name)
	result, returnValue, returnValueErr := vpcService.UpdateNetworkAcl(options)
	return result, returnValue, returnValueErr
}

// ListNetworkAclRules - GET
// /network_acls/{network_acl_id}/rules
// List all rules for a network ACL
func ListNetworkAclRules(vpcService *vpcv1.VpcV1, aclID string) (*vpcv1.NetworkACLRuleCollection, *core.DetailedResponse, error) {
	options := &vpcv1.ListNetworkAclRulesOptions{}
	options.SetNetworkAclID(aclID)
	result, returnValue, returnValueErr := vpcService.ListNetworkAclRules(options)
	return result, returnValue, returnValueErr
}

// CreateNetworkAclRule - POST
// /network_acls/{network_acl_id}/rules
// Create a rule
func CreateNetworkAclRule(vpcService *vpcv1.VpcV1, name, aclID string) (vpcv1.NetworkACLRuleIntf, *core.DetailedResponse, error) {
	options := &vpcv1.CreateNetworkAclRuleOptions{}
	options.SetNetworkAclID(aclID)
	options.SetNetworkACLRulePrototype(&vpcv1.NetworkACLRulePrototype{
		Action:      core.StringPtr("allow"),
		Direction:   core.StringPtr("inbound"),
		Destination: core.StringPtr("0.0.0.0/0"),
		Source:      core.StringPtr("0.0.0.0/0"),
		Protocol:    core.StringPtr("all"),
		Name:        core.StringPtr(name),
	})
	result, returnValue, returnValueErr := vpcService.CreateNetworkAclRule(options)
	return result, returnValue, returnValueErr
}

// DeleteNetworkAclRule DELETE
// /network_acls/{network_acl_id}/rules/{id}
// Delete specified rule
func DeleteNetworkAclRule(vpcService *vpcv1.VpcV1, aclID, ruleID string) (*core.DetailedResponse, error) {
	options := &vpcv1.DeleteNetworkAclRuleOptions{}
	options.SetID(ruleID)
	options.SetNetworkAclID(aclID)
	returnValue, returnValueErr := vpcService.DeleteNetworkAclRule(options)
	return returnValue, returnValueErr
}

// GetNetworkAclRule GET
// /network_acls/{network_acl_id}/rules/{id}
// Retrieve specified rule
func GetNetworkAclRule(vpcService *vpcv1.VpcV1, aclID, ruleID string) (vpcv1.NetworkACLRuleIntf, *core.DetailedResponse, error) {
	options := &vpcv1.GetNetworkAclRuleOptions{}
	options.SetID(ruleID)
	options.SetNetworkAclID(aclID)
	result, returnValue, returnValueErr := vpcService.GetNetworkAclRule(options)
	return result, returnValue, returnValueErr
}

// UpdateNetworkAclRule PATCH
// /network_acls/{network_acl_id}/rules/{id}
// Update a rule
func UpdateNetworkAclRule(vpcService *vpcv1.VpcV1, aclID, ruleID, name string) (vpcv1.NetworkACLRuleIntf, *core.DetailedResponse, error) {
	options := &vpcv1.UpdateNetworkAclRuleOptions{}
	options.SetID(ruleID)
	options.SetNetworkAclID(aclID)
	options.SetName(name)
	result, returnValue, returnValueErr := vpcService.UpdateNetworkAclRule(options)
	return result, returnValue, returnValueErr
}

/**
 * Security Groups
 *
 */

// ListSecurityGroups GET
// /security_groups
// List all security groups
func ListSecurityGroups(vpcService *vpcv1.VpcV1) (*vpcv1.SecurityGroupCollection, *core.DetailedResponse, error) {
	options := &vpcv1.ListSecurityGroupsOptions{}
	result, returnValue, returnValueErr := vpcService.ListSecurityGroups(options)
	return result, returnValue, returnValueErr
}

// CreateSecurityGroup POST
// /security_groups
// Create a security group
func CreateSecurityGroup(vpcService *vpcv1.VpcV1, name, vpcID string) (*vpcv1.SecurityGroup, *core.DetailedResponse, error) {
	options := &vpcv1.CreateSecurityGroupOptions{}

	options.SetVpc(&vpcv1.VPCIdentity{
		ID: core.StringPtr(vpcID),
	})
	options.SetName(name)
	result, returnValue, returnValueErr := vpcService.CreateSecurityGroup(options)
	return result, returnValue, returnValueErr
}

// DeleteSecurityGroup DELETE
// /security_groups/{id}
// Delete a security group
func DeleteSecurityGroup(vpcService *vpcv1.VpcV1, id string) (*core.DetailedResponse, error) {
	options := &vpcv1.DeleteSecurityGroupOptions{}

	options.SetID(id)
	returnValue, returnValueErr := vpcService.DeleteSecurityGroup(options)
	return returnValue, returnValueErr
}

// GetSecurityGroup GET
// /security_groups/{id}
// Retrieve a security group
func GetSecurityGroup(vpcService *vpcv1.VpcV1, id string) (*vpcv1.SecurityGroup, *core.DetailedResponse, error) {
	options := &vpcv1.GetSecurityGroupOptions{}
	options.SetID(id)
	result, returnValue, returnValueErr := vpcService.GetSecurityGroup(options)
	return result, returnValue, returnValueErr
}

// UpdateSecurityGroup PATCH
// /security_groups/{id}
// Update a security group
func UpdateSecurityGroup(vpcService *vpcv1.VpcV1, id, name string) (*vpcv1.SecurityGroup, *core.DetailedResponse, error) {
	options := &vpcv1.UpdateSecurityGroupOptions{}
	options.SetID(id)
	options.SetName(name)
	result, returnValue, returnValueErr := vpcService.UpdateSecurityGroup(options)
	return result, returnValue, returnValueErr
}

// ListSecurityGroupNetworkInterfaces GET
// /security_groups/{security_group_id}/network_interfaces
// List a security group's network interfaces
// ListSecurityGroupNetworkInterfaces
func ListSecurityGroupNetworkInterfaces(vpcService *vpcv1.VpcV1, sgID string) (*vpcv1.NetworkInterfaceCollection, *core.DetailedResponse, error) {
	options := &vpcv1.ListSecurityGroupNetworkInterfacesOptions{}
	options.SetSecurityGroupID(sgID)
	result, returnValue, returnValueErr := vpcService.ListSecurityGroupNetworkInterfaces(options)
	return result, returnValue, returnValueErr
}

// DeleteSecurityGroupNetworkInterfaceBinding DELETE
// /security_groups/{security_group_id}/network_interfaces/{id}
// Remove a network interface from a security group.
func DeleteSecurityGroupNetworkInterfaceBinding(vpcService *vpcv1.VpcV1, id, vnicID string) (*core.DetailedResponse, error) {
	options := &vpcv1.DeleteSecurityGroupNetworkInterfaceBindingOptions{}
	options.SetSecurityGroupID(id)
	options.SetID(vnicID)
	returnValue, returnValueErr := vpcService.DeleteSecurityGroupNetworkInterfaceBinding(options)
	return returnValue, returnValueErr
}

// GetSecurityGroupNetworkInterface GET
// /security_groups/{security_group_id}/network_interfaces/{id}
// Retrieve a network interface in a security group
func GetSecurityGroupNetworkInterface(vpcService *vpcv1.VpcV1, id, vnicID string) (*vpcv1.NetworkInterface, *core.DetailedResponse, error) {
	options := &vpcv1.GetSecurityGroupNetworkInterfaceOptions{}
	options.SetSecurityGroupID(id)
	options.SetID(vnicID)
	result, returnValue, returnValueErr := vpcService.GetSecurityGroupNetworkInterface(options)
	return result, returnValue, returnValueErr
}

// CreateSecurityGroupNetworkInterfaceBinding PUT
// /security_groups/{security_group_id}/network_interfaces/{id}
// Add a network interface to a security group
func CreateSecurityGroupNetworkInterfaceBinding(vpcService *vpcv1.VpcV1, id, vnicID string) (*vpcv1.NetworkInterface, *core.DetailedResponse, error) {
	options := &vpcv1.CreateSecurityGroupNetworkInterfaceBindingOptions{}
	options.SetSecurityGroupID(id)
	options.SetID(vnicID)
	result, returnValue, returnValueErr := vpcService.CreateSecurityGroupNetworkInterfaceBinding(options)
	return result, returnValue, returnValueErr
}

// ListSecurityGroupRules GET
// /security_groups/{security_group_id}/rules
// List all the rules of a security group
func ListSecurityGroupRules(vpcService *vpcv1.VpcV1, id string) (*vpcv1.SecurityGroupRuleCollection, *core.DetailedResponse, error) {
	options := &vpcv1.ListSecurityGroupRulesOptions{}
	options.SetSecurityGroupID(id)
	result, returnValue, returnValueErr := vpcService.ListSecurityGroupRules(options)
	return result, returnValue, returnValueErr
}

// CreateSecurityGroupRule POST
// /security_groups/{security_group_id}/rules
// Create a security group rule
func CreateSecurityGroupRule(vpcService *vpcv1.VpcV1, sgID string) (vpcv1.SecurityGroupRuleIntf, *core.DetailedResponse, error) {
	options := &vpcv1.CreateSecurityGroupRuleOptions{}
	options.SetSecurityGroupID(sgID)
	options.SetSecurityGroupRulePrototype(&vpcv1.SecurityGroupRulePrototype{
		Direction: core.StringPtr("inbound"),
		Protocol:  core.StringPtr("all"),
		IpVersion: core.StringPtr("ipv4"),
	})
	result, returnValue, returnValueErr := vpcService.CreateSecurityGroupRule(options)
	return result, returnValue, returnValueErr
}

// DeleteSecurityGroupRule DELETE
// /security_groups/{security_group_id}/rules/{id}
// Delete a security group rule
func DeleteSecurityGroupRule(vpcService *vpcv1.VpcV1, sgID, sgRuleID string) (*core.DetailedResponse, error) {
	options := &vpcv1.DeleteSecurityGroupRuleOptions{}
	options.SetSecurityGroupID(sgID)
	options.SetID(sgRuleID)
	returnValue, returnValueErr := vpcService.DeleteSecurityGroupRule(options)
	return returnValue, returnValueErr
}

// GetSecurityGroupRule GET
// /security_groups/{security_group_id}/rules/{id}
// Retrieve a security group rule
func GetSecurityGroupRule(vpcService *vpcv1.VpcV1, sgID, sgRuleID string) (vpcv1.SecurityGroupRuleIntf, *core.DetailedResponse, error) {
	options := &vpcv1.GetSecurityGroupRuleOptions{}
	options.SetSecurityGroupID(sgID)
	options.SetID(sgRuleID)
	result, returnValue, returnValueErr := vpcService.GetSecurityGroupRule(options)
	return result, returnValue, returnValueErr
}

// UpdateSecurityGroupRule PATCH
// /security_groups/{security_group_id}/rules/{id}
// Update a security group rule
func UpdateSecurityGroupRule(vpcService *vpcv1.VpcV1, sgID, sgRuleID string) (vpcv1.SecurityGroupRuleIntf, *core.DetailedResponse, error) {
	options := &vpcv1.UpdateSecurityGroupRuleOptions{}
	options.SetSecurityGroupID(sgID)
	options.SetID(sgRuleID)
	options.SetRemote(&vpcv1.SecurityGroupRulePatchRemote{
		Address: core.StringPtr("1.1.1.11"),
	})
	result, returnValue, returnValueErr := vpcService.UpdateSecurityGroupRule(options)
	return result, returnValue, returnValueErr
}

/**
 * Load Balancers
 *
 */

// ListLoadBalancers GET
// /load_balancers
// List all load balancers
func ListLoadBalancers(vpcService *vpcv1.VpcV1) (*vpcv1.LoadBalancerCollection, *core.DetailedResponse, error) {
	options := &vpcv1.ListLoadBalancersOptions{}
	result, returnValue, returnValueErr := vpcService.ListLoadBalancers(options)
	return result, returnValue, returnValueErr
}

// CreateLoadBalancer POST
// /load_balancers
// Create and provision a load balancer
func CreateLoadBalancer(vpcService *vpcv1.VpcV1, name, subnetID string) (*vpcv1.LoadBalancer, *core.DetailedResponse, error) {
	options := &vpcv1.CreateLoadBalancerOptions{}
	options.SetIsPublic(true)
	options.SetName(name)
	var subnetArray = []vpcv1.SubnetIdentityIntf{
		&vpcv1.SubnetIdentity{
			ID: core.StringPtr(subnetID),
		},
	}
	options.SetSubnets(subnetArray)
	result, returnValue, returnValueErr := vpcService.CreateLoadBalancer(options)
	return result, returnValue, returnValueErr
}

// DeleteLoadBalancer DELETE
// /load_balancers/{id}
// Delete a load balancer
func DeleteLoadBalancer(vpcService *vpcv1.VpcV1, id string) (*core.DetailedResponse, error) {
	deleteVpcOptions := &vpcv1.DeleteLoadBalancerOptions{}
	deleteVpcOptions.SetID(id)
	returnValue, returnValueErr := vpcService.DeleteLoadBalancer(deleteVpcOptions)
	return returnValue, returnValueErr
}

// GetLoadBalancer GET
// /load_balancers/{id}
// Retrieve a load balancer
func GetLoadBalancer(vpcService *vpcv1.VpcV1, id string) (*vpcv1.LoadBalancer, *core.DetailedResponse, error) {
	options := &vpcv1.GetLoadBalancerOptions{}
	options.SetID(id)
	result, returnValue, returnValueErr := vpcService.GetLoadBalancer(options)
	return result, returnValue, returnValueErr
}

// UpdateLoadBalancer PATCH
// /load_balancers/{id}
// Update a load balancer
func UpdateLoadBalancer(vpcService *vpcv1.VpcV1, id, name string) (*vpcv1.LoadBalancer, *core.DetailedResponse, error) {
	options := &vpcv1.UpdateLoadBalancerOptions{
		Name: core.StringPtr(name),
	}
	options.SetID(id)
	result, returnValue, returnValueErr := vpcService.UpdateLoadBalancer(options)
	return result, returnValue, returnValueErr
}

// GetLoadBalancerStatistics GET
// /load_balancers/{id}/statistics
// List statistics of a load balancer
func GetLoadBalancerStatistics(vpcService *vpcv1.VpcV1, id string) (*vpcv1.LoadBalancerStatistics, *core.DetailedResponse, error) {
	options := &vpcv1.GetLoadBalancerStatisticsOptions{}
	options.SetID(id)
	result, returnValue, returnValueErr := vpcService.GetLoadBalancerStatistics(options)
	return result, returnValue, returnValueErr
}

// ListLoadBalancerListeners GET
// /load_balancers/{load_balancer_id}/listeners
// List all listeners of the load balancer
func ListLoadBalancerListeners(vpcService *vpcv1.VpcV1, id string) (*vpcv1.LoadBalancerListenerCollection, *core.DetailedResponse, error) {
	options := &vpcv1.ListLoadBalancerListenersOptions{}
	options.SetLoadBalancerID(id)
	result, returnValue, returnValueErr := vpcService.ListLoadBalancerListeners(options)
	return result, returnValue, returnValueErr
}

// CreateLoadBalancerListener POST
// /load_balancers/{load_balancer_id}/listeners
// Create a listener
func CreateLoadBalancerListener(vpcService *vpcv1.VpcV1, lbID string) (*vpcv1.LoadBalancerListener, *core.DetailedResponse, error) {
	options := &vpcv1.CreateLoadBalancerListenerOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetPort(rand.Int63n(100))
	options.SetProtocol("http")
	result, returnValue, returnValueErr := vpcService.CreateLoadBalancerListener(options)
	return result, returnValue, returnValueErr
}

// DeleteLoadBalancerListener DELETE
// /load_balancers/{load_balancer_id}/listeners/{id}
// Delete a listener
func DeleteLoadBalancerListener(vpcService *vpcv1.VpcV1, lbID, listenerID string) (*core.DetailedResponse, error) {
	options := &vpcv1.DeleteLoadBalancerListenerOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetID(listenerID)
	returnValue, returnValueErr := vpcService.DeleteLoadBalancerListener(options)
	return returnValue, returnValueErr
}

// GetLoadBalancerListener GET
// /load_balancers/{load_balancer_id}/listeners/{id}
// Retrieve a listener
func GetLoadBalancerListener(vpcService *vpcv1.VpcV1, lbID, listenerID string) (*vpcv1.LoadBalancerListener, *core.DetailedResponse, error) {
	options := &vpcv1.GetLoadBalancerListenerOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetID(listenerID)
	result, returnValue, returnValueErr := vpcService.GetLoadBalancerListener(options)
	return result, returnValue, returnValueErr
}

// UpdateLoadBalancerListener PATCH
// /load_balancers/{load_balancer_id}/listeners/{id}
// Update a listener
func UpdateLoadBalancerListener(vpcService *vpcv1.VpcV1, lbID, listenerID string) (*vpcv1.LoadBalancerListener, *core.DetailedResponse, error) {
	options := &vpcv1.UpdateLoadBalancerListenerOptions{}

	options.SetLoadBalancerID(lbID)
	options.SetID(listenerID)
	options.SetProtocol("tcp")
	result, returnValue, returnValueErr := vpcService.UpdateLoadBalancerListener(options)
	return result, returnValue, returnValueErr
}

// ListLoadBalancerListenerPolicies GET
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies
// List all policies of the load balancer listener
func ListLoadBalancerListenerPolicies(vpcService *vpcv1.VpcV1, lbID, listenerID string) (*vpcv1.LoadBalancerListenerPolicyCollection, *core.DetailedResponse, error) {
	options := &vpcv1.ListLoadBalancerListenerPoliciesOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	result, returnValue, returnValueErr := vpcService.ListLoadBalancerListenerPolicies(options)
	return result, returnValue, returnValueErr
}

// CreateLoadBalancerListenerPolicy POST
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies
func CreateLoadBalancerListenerPolicy(vpcService *vpcv1.VpcV1, lbID, listenerID string) (*vpcv1.LoadBalancerListenerPolicy, *core.DetailedResponse, error) {
	options := &vpcv1.CreateLoadBalancerListenerPolicyOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	options.SetPriority(2)
	options.SetAction("reject")
	result, returnValue, returnValueErr := vpcService.CreateLoadBalancerListenerPolicy(options)
	return result, returnValue, returnValueErr
}

// DeleteLoadBalancerListenerPolicy DELETE
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies/{id}
// Delete a policy of the load balancer listener
func DeleteLoadBalancerListenerPolicy(vpcService *vpcv1.VpcV1, lbID, listenerID, policyID string) (*core.DetailedResponse, error) {
	options := &vpcv1.DeleteLoadBalancerListenerPolicyOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	options.SetID(policyID)
	returnValue, returnValueErr := vpcService.DeleteLoadBalancerListenerPolicy(options)
	return returnValue, returnValueErr
}

// GetLoadBalancerListenerPolicy GET
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies/{id}
// Retrieve a policy of the load balancer listener
func GetLoadBalancerListenerPolicy(vpcService *vpcv1.VpcV1, lbID, listenerID, policyID string) (*vpcv1.LoadBalancerListenerPolicy, *core.DetailedResponse, error) {
	options := &vpcv1.GetLoadBalancerListenerPolicyOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	options.SetID(policyID)
	result, returnValue, returnValueErr := vpcService.GetLoadBalancerListenerPolicy(options)
	return result, returnValue, returnValueErr
}

// UpdateLoadBalancerListenerPolicy PATCH
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies/{id}
// Update a policy of the load balancer listener
func UpdateLoadBalancerListenerPolicy(vpcService *vpcv1.VpcV1, lbID, listenerID, policyID, targetPoolID string) (*vpcv1.LoadBalancerListenerPolicy, *core.DetailedResponse, error) {
	options := &vpcv1.UpdateLoadBalancerListenerPolicyOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	options.SetID(policyID)

	options.SetPriority(4)
	options.SetName("some-name")
	target := &vpcv1.LoadBalancerListenerPolicyPatchTarget{
		ID: core.StringPtr(targetPoolID),
	}
	options.SetTarget(target)
	result, returnValue, returnValueErr := vpcService.UpdateLoadBalancerListenerPolicy(options)
	return result, returnValue, returnValueErr
}

// ListLoadBalancerListenerPolicyRules GET
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies/{policy_id}/rules
// List all rules of the load balancer listener policy
func ListLoadBalancerListenerPolicyRules(vpcService *vpcv1.VpcV1, lbID, listenerID, policyID string) (*vpcv1.LoadBalancerListenerPolicyRuleCollection, *core.DetailedResponse, error) {
	options := &vpcv1.ListLoadBalancerListenerPolicyRulesOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	options.SetPolicyID(policyID)
	result, returnValue, returnValueErr := vpcService.ListLoadBalancerListenerPolicyRules(options)
	return result, returnValue, returnValueErr
}

// CreateLoadBalancerListenerPolicyRule POST
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies/{policy_id}/rules
// Create a rule for the load balancer listener policy
func CreateLoadBalancerListenerPolicyRule(vpcService *vpcv1.VpcV1, lbID, listenerID, policyID string) (*vpcv1.LoadBalancerListenerPolicyRule, *core.DetailedResponse, error) {
	options := &vpcv1.CreateLoadBalancerListenerPolicyRuleOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	options.SetPolicyID(policyID)
	options.SetCondition("contains")
	options.SetType("hostname")
	options.SetValue("one")
	result, returnValue, returnValueErr := vpcService.CreateLoadBalancerListenerPolicyRule(options)
	return result, returnValue, returnValueErr
}

// DeleteLoadBalancerListenerPolicyRule DELETE
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies/{policy_id}/rules/{id}
// Delete a rule from the load balancer listener policy
func DeleteLoadBalancerListenerPolicyRule(vpcService *vpcv1.VpcV1, lbID, listenerID, policyID, ruleID string) (*core.DetailedResponse, error) {
	options := &vpcv1.DeleteLoadBalancerListenerPolicyRuleOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	options.SetPolicyID(policyID)
	options.SetID(ruleID)
	returnValue, returnValueErr := vpcService.DeleteLoadBalancerListenerPolicyRule(options)
	return returnValue, returnValueErr
}

// GetLoadBalancerListenerPolicyRule GET
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies/{policy_id}/rules/{id}
// Retrieve a rule of the load balancer listener policy
func GetLoadBalancerListenerPolicyRule(vpcService *vpcv1.VpcV1, lbID, listenerID, policyID, ruleID string) (*vpcv1.LoadBalancerListenerPolicyRule, *core.DetailedResponse, error) {
	options := &vpcv1.GetLoadBalancerListenerPolicyRuleOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	options.SetPolicyID(policyID)
	options.SetID(ruleID)
	result, returnValue, returnValueErr := vpcService.GetLoadBalancerListenerPolicyRule(options)
	return result, returnValue, returnValueErr
}

// UpdateLoadBalancerListenerPolicyRule PATCH
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies/{policy_id}/rules/{id}
// Update a rule of the load balancer listener policy
func UpdateLoadBalancerListenerPolicyRule(vpcService *vpcv1.VpcV1, lbID, listenerID, policyID, ruleID string) (*vpcv1.LoadBalancerListenerPolicyRule, *core.DetailedResponse, error) {
	options := &vpcv1.UpdateLoadBalancerListenerPolicyRuleOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	options.SetPolicyID(policyID)
	options.SetID(ruleID)
	options.SetCondition("equals")
	options.SetType("header")
	options.SetValue("1")
	options.SetField("some-name")
	result, returnValue, returnValueErr := vpcService.UpdateLoadBalancerListenerPolicyRule(options)
	return result, returnValue, returnValueErr
}

// ListLoadBalancerPools GET
// /load_balancers/{load_balancer_id}/pools
// List all pools of the load balancer
func ListLoadBalancerPools(vpcService *vpcv1.VpcV1, id string) (*vpcv1.LoadBalancerPoolCollection, *core.DetailedResponse, error) {
	options := &vpcv1.ListLoadBalancerPoolsOptions{}
	options.SetLoadBalancerID(id)
	result, returnValue, returnValueErr := vpcService.ListLoadBalancerPools(options)
	return result, returnValue, returnValueErr
}

// CreateLoadBalancerPool POST
// /load_balancers/{load_balancer_id}/pools
// Create a load balancer pool
func CreateLoadBalancerPool(vpcService *vpcv1.VpcV1, lbID, name string) (*vpcv1.LoadBalancerPool, *core.DetailedResponse, error) {
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
	result, returnValue, returnValueErr := vpcService.CreateLoadBalancerPool(options)
	return result, returnValue, returnValueErr
}

// DeleteLoadBalancerPool DELETE
// /load_balancers/{load_balancer_id}/pools/{id}
// Delete a pool
func DeleteLoadBalancerPool(vpcService *vpcv1.VpcV1, lbID, poolID string) (*core.DetailedResponse, error) {
	options := &vpcv1.DeleteLoadBalancerPoolOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetID(poolID)
	returnValue, returnValueErr := vpcService.DeleteLoadBalancerPool(options)
	return returnValue, returnValueErr
}

// GetLoadBalancerPool GET
// /load_balancers/{load_balancer_id}/pools/{id}
// Retrieve a load balancer pool
func GetLoadBalancerPool(vpcService *vpcv1.VpcV1, lbID, poolID string) (*vpcv1.LoadBalancerPool, *core.DetailedResponse, error) {
	options := &vpcv1.GetLoadBalancerPoolOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetID(poolID)
	result, returnValue, returnValueErr := vpcService.GetLoadBalancerPool(options)
	return result, returnValue, returnValueErr
}

// UpdateLoadBalancerPool PATCH
// /load_balancers/{load_balancer_id}/pools/{id}
// Update a load balancer pool
func UpdateLoadBalancerPool(vpcService *vpcv1.VpcV1, lbID, poolID string) (*vpcv1.LoadBalancerPool, *core.DetailedResponse, error) {
	options := &vpcv1.UpdateLoadBalancerPoolOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetID(poolID)
	options.SetProtocol("tcp")
	result, returnValue, returnValueErr := vpcService.UpdateLoadBalancerPool(options)
	return result, returnValue, returnValueErr
}

// ListLoadBalancerPoolMembers GET
// /load_balancers/{load_balancer_id}/pools/{pool_id}/members
// List all members of the load balancer pool
func ListLoadBalancerPoolMembers(vpcService *vpcv1.VpcV1, lbID, poolID string) (*vpcv1.LoadBalancerPoolMemberCollection, *core.DetailedResponse, error) {
	options := &vpcv1.ListLoadBalancerPoolMembersOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetPoolID(poolID)
	result, returnValue, returnValueErr := vpcService.ListLoadBalancerPoolMembers(options)
	return result, returnValue, returnValueErr
}

// CreateLoadBalancerPoolMember POST
// /load_balancers/{load_balancer_id}/pools/{pool_id}/members
// Create a member in the load balancer pool
func CreateLoadBalancerPoolMember(vpcService *vpcv1.VpcV1, lbID, poolID string) (*vpcv1.LoadBalancerPoolMember, *core.DetailedResponse, error) {
	options := &vpcv1.CreateLoadBalancerPoolMemberOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetPoolID(poolID)
	options.SetPort(1234)
	options.SetTarget(&vpcv1.LoadBalancerPoolMemberTargetPrototype{
		Address: core.StringPtr("12.12.0.0"),
	})
	result, returnValue, returnValueErr := vpcService.CreateLoadBalancerPoolMember(options)
	return result, returnValue, returnValueErr
}

// UpdateLoadBalancerPoolMembers PUT
// /load_balancers/{load_balancer_id}/pools/{pool_id}/members
// Update members of the load balancer pool
func UpdateLoadBalancerPoolMembers(vpcService *vpcv1.VpcV1, lbID, poolID string) (*vpcv1.LoadBalancerPoolMemberCollection, *core.DetailedResponse, error) {
	options := &vpcv1.UpdateLoadBalancerPoolMembersOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetPoolID(poolID)
	options.SetMembers([]vpcv1.LoadBalancerPoolMemberPrototype{
		vpcv1.LoadBalancerPoolMemberPrototype{
			Port: core.Int64Ptr(2345),
			Target: &vpcv1.LoadBalancerPoolMemberTargetPrototype{
				Address: core.StringPtr("13.13.0.0"),
			},
		},
	})
	result, returnValue, returnValueErr := vpcService.UpdateLoadBalancerPoolMembers(options)
	return result, returnValue, returnValueErr
}

// DeleteLoadBalancerPoolMember DELETE
// /load_balancers/{load_balancer_id}/pools/{pool_id}/members/{id}
// Delete a member from the load balancer pool
func DeleteLoadBalancerPoolMember(vpcService *vpcv1.VpcV1, lbID, poolID, memberID string) (*core.DetailedResponse, error) {
	options := &vpcv1.DeleteLoadBalancerPoolMemberOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetPoolID(poolID)
	options.SetID(memberID)
	returnValue, returnValueErr := vpcService.DeleteLoadBalancerPoolMember(options)
	return returnValue, returnValueErr
}

// GetLoadBalancerPoolMember GET
// /load_balancers/{load_balancer_id}/pools/{pool_id}/members/{id}
// Retrieve a member in the load balancer pool
func GetLoadBalancerPoolMember(vpcService *vpcv1.VpcV1, lbID, poolID, memberID string) (*vpcv1.LoadBalancerPoolMember, *core.DetailedResponse, error) {
	options := &vpcv1.GetLoadBalancerPoolMemberOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetPoolID(poolID)
	options.SetID(memberID)
	result, returnValue, returnValueErr := vpcService.GetLoadBalancerPoolMember(options)
	return result, returnValue, returnValueErr
}

// UpdateLoadBalancerPoolMember PATCH
// /load_balancers/{load_balancer_id}/pools/{pool_id}/members/{id}
func UpdateLoadBalancerPoolMember(vpcService *vpcv1.VpcV1, lbID, poolID, memberID string) (*vpcv1.LoadBalancerPoolMember, *core.DetailedResponse, error) {
	options := &vpcv1.UpdateLoadBalancerPoolMemberOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetPoolID(poolID)
	options.SetID(memberID)
	options.SetPort(3434)
	result, returnValue, returnValueErr := vpcService.UpdateLoadBalancerPoolMember(options)
	return result, returnValue, returnValueErr
}

/**
 * VPN
 *
 */

// ListIkePolicies GET
// /ike_policies
// List all IKE policies
func ListIkePolicies(vpcService *vpcv1.VpcV1) (*vpcv1.IKEPolicyCollection, *core.DetailedResponse, error) {
	options := vpcService.NewListIkePoliciesOptions()
	result, returnValue, returnValueErr := vpcService.ListIkePolicies(options)
	return result, returnValue, returnValueErr
}

// CreateIkePolicy POST
// /ike_policies
// Create an IKE policy
func CreateIkePolicy(vpcService *vpcv1.VpcV1, name string) (*vpcv1.IKEPolicy, *core.DetailedResponse, error) {
	options := &vpcv1.CreateIkePolicyOptions{}
	options.SetName(name)
	options.SetAuthenticationAlgorithm("md5")
	options.SetDhGroup(2)
	options.SetEncryptionAlgorithm("aes128")
	options.SetIkeVersion(1)
	result, returnValue, returnValueErr := vpcService.CreateIkePolicy(options)
	return result, returnValue, returnValueErr
}

// DeleteIkePolicy DELETE
// /ike_policies/{id}
// Delete an IKE policy
func DeleteIkePolicy(vpcService *vpcv1.VpcV1, id string) (*core.DetailedResponse, error) {
	options := vpcService.NewDeleteIkePolicyOptions(id)
	returnValue, returnValueErr := vpcService.DeleteIkePolicy(options)
	return returnValue, returnValueErr
}

// GetIkePolicy GET
// /ike_policies/{id}
// Retrieve the specified IKE policy
func GetIkePolicy(vpcService *vpcv1.VpcV1, id string) (*vpcv1.IKEPolicy, *core.DetailedResponse, error) {
	options := vpcService.NewGetIkePolicyOptions(id)
	result, returnValue, returnValueErr := vpcService.GetIkePolicy(options)
	return result, returnValue, returnValueErr
}

// UpdateIkePolicy PATCH
// /ike_policies/{id}
// Update an IKE policy
func UpdateIkePolicy(vpcService *vpcv1.VpcV1, id string) (*vpcv1.IKEPolicy, *core.DetailedResponse, error) {
	options := &vpcv1.UpdateIkePolicyOptions{
		ID:      core.StringPtr(id),
		DhGroup: core.Int64Ptr(5),
		Name:    core.StringPtr("go-ike-policy-2"),
	}
	result, returnValue, returnValueErr := vpcService.UpdateIkePolicy(options)
	return result, returnValue, returnValueErr
}

// ListVpnGatewayIkePolicyConnections GET
// /ike_policies/{id}/connections
// Lists all the connections that use the specified policy
func ListVpnGatewayIkePolicyConnections(vpcService *vpcv1.VpcV1, id string) (*vpcv1.VPNGatewayConnectionCollection, *core.DetailedResponse, error) {
	options := &vpcv1.ListVpnGatewayIkePolicyConnectionsOptions{
		ID: core.StringPtr(id),
	}
	result, returnValue, returnValueErr := vpcService.ListVpnGatewayIkePolicyConnections(options)
	return result, returnValue, returnValueErr
}

// ListIpsecPolicies GET
// /ipsec_policies
// List all IPsec policies
func ListIpsecPolicies(vpcService *vpcv1.VpcV1) (*vpcv1.IPsecPolicyCollection, *core.DetailedResponse, error) {
	options := &vpcv1.ListIpsecPoliciesOptions{}

	result, returnValue, returnValueErr := vpcService.ListIpsecPolicies(options)
	return result, returnValue, returnValueErr
}

// CreateIpsecPolicy POST
// /ipsec_policies
// Create an IPsec policy
func CreateIpsecPolicy(vpcService *vpcv1.VpcV1, name string) (*vpcv1.IPsecPolicy, *core.DetailedResponse, error) {
	options := &vpcv1.CreateIpsecPolicyOptions{}
	options.SetName(name)
	options.SetAuthenticationAlgorithm("md5")
	options.SetEncryptionAlgorithm("aes128")
	options.SetPfs("disabled")

	result, returnValue, returnValueErr := vpcService.CreateIpsecPolicy(options)
	return result, returnValue, returnValueErr
}

// DeleteIpsecPolicy DELETE
// /ipsec_policies/{id}
// Delete an IPsec policy
func DeleteIpsecPolicy(vpcService *vpcv1.VpcV1, id string) (*core.DetailedResponse, error) {
	options := vpcService.NewDeleteIpsecPolicyOptions(id)
	returnValue, returnValueErr := vpcService.DeleteIpsecPolicy(options)
	return returnValue, returnValueErr
}

// GetIpsecPolicy GET
// /ipsec_policies/{id}
// Retrieve the specified IPsec policy
func GetIpsecPolicy(vpcService *vpcv1.VpcV1, id string) (*vpcv1.IPsecPolicy, *core.DetailedResponse, error) {
	options := vpcService.NewGetIpsecPolicyOptions(id)
	result, returnValue, returnValueErr := vpcService.GetIpsecPolicy(options)
	return result, returnValue, returnValueErr
}

// UpdateIpsecPolicy PATCH
// /ipsec_policies/{id}
// Update an IPsec policy
func UpdateIpsecPolicy(vpcService *vpcv1.VpcV1, id string) (*vpcv1.IPsecPolicy, *core.DetailedResponse, error) {
	options := &vpcv1.UpdateIpsecPolicyOptions{
		ID: core.StringPtr(id),
	}
	options.SetEncryptionAlgorithm("3des")
	result, returnValue, returnValueErr := vpcService.UpdateIpsecPolicy(options)
	return result, returnValue, returnValueErr
}

// ListVpnGatewayIpsecPolicyConnections GET
// /ipsec_policies/{id}/connections
// Lists all the connections that use the specified policy
func ListVpnGatewayIpsecPolicyConnections(vpcService *vpcv1.VpcV1, id string) (*vpcv1.VPNGatewayConnectionCollection, *core.DetailedResponse, error) {
	options := &vpcv1.ListVpnGatewayIpsecPolicyConnectionsOptions{
		ID: core.StringPtr(id),
	}
	result, returnValue, returnValueErr := vpcService.ListVpnGatewayIpsecPolicyConnections(options)
	return result, returnValue, returnValueErr
}

// ListVpnGateways GET
// /vpn_gateways
// List all VPN gateways
func ListVpnGateways(vpcService *vpcv1.VpcV1) (*vpcv1.VPNGatewayCollection, *core.DetailedResponse, error) {
	options := &vpcv1.ListVpnGatewaysOptions{}
	result, returnValue, returnValueErr := vpcService.ListVpnGateways(options)
	return result, returnValue, returnValueErr
}

// CreateVpnGateway POST
// /vpn_gateways
// Create a VPN gateway
func CreateVpnGateway(vpcService *vpcv1.VpcV1, subnetID, name string) (*vpcv1.VPNGateway, *core.DetailedResponse, error) {
	options := &vpcv1.CreateVpnGatewayOptions{}
	options.SetName(name)
	options.SetSubnet(&vpcv1.SubnetIdentity{
		ID: core.StringPtr(subnetID),
	})
	result, returnValue, returnValueErr := vpcService.CreateVpnGateway(options)
	return result, returnValue, returnValueErr
}

// DeleteVpnGateway DELETE
// /vpn_gateways/{id}
// Delete a VPN gateway
func DeleteVpnGateway(vpcService *vpcv1.VpcV1, id string) (*core.DetailedResponse, error) {
	options := vpcService.NewDeleteVpnGatewayOptions(id)
	returnValue, returnValueErr := vpcService.DeleteVpnGateway(options)
	return returnValue, returnValueErr
}

// GetVpnGateway GET
// /vpn_gateways/{id}
// Retrieve the specified VPN gateway
func GetVpnGateway(vpcService *vpcv1.VpcV1, id string) (*vpcv1.VPNGateway, *core.DetailedResponse, error) {
	options := vpcService.NewGetVpnGatewayOptions(id)
	result, returnValue, returnValueErr := vpcService.GetVpnGateway(options)
	return result, returnValue, returnValueErr
}

// UpdateVpnGateway PATCH
// /vpn_gateways/{id}
// Update a VPN gateway
func UpdateVpnGateway(vpcService *vpcv1.VpcV1, id, name string) (*vpcv1.VPNGateway, *core.DetailedResponse, error) {
	options := &vpcv1.UpdateVpnGatewayOptions{
		ID:   core.StringPtr(id),
		Name: core.StringPtr(name),
	}
	result, returnValue, returnValueErr := vpcService.UpdateVpnGateway(options)
	return result, returnValue, returnValueErr
}

// ListVpnGatewayConnections GET
// /vpn_gateways/{vpn_gateway_id}/connections
// List all the connections of a VPN gateway
func ListVpnGatewayConnections(vpcService *vpcv1.VpcV1, gatewayID string) (*vpcv1.VPNGatewayConnectionCollection, *core.DetailedResponse, error) {
	options := &vpcv1.ListVpnGatewayConnectionsOptions{}
	options.SetVpnGatewayID(gatewayID)
	result, returnValue, returnValueErr := vpcService.ListVpnGatewayConnections(options)
	return result, returnValue, returnValueErr
}

// CreateVpnGatewayConnection POST
// /vpn_gateways/{vpn_gateway_id}/connections
// Create a VPN connection
func CreateVpnGatewayConnection(vpcService *vpcv1.VpcV1, gatewayID, name string) (*vpcv1.VPNGatewayConnection, *core.DetailedResponse, error) {
	options := &vpcv1.CreateVpnGatewayConnectionOptions{}
	options.SetName(name)
	options.SetVpnGatewayID(gatewayID)
	options.SetPeerAddress("192.168.0.1")
	options.SetPsk("pre-shared-key")
	local := []string{"192.132.0.0/28"}
	options.SetLocalCidrs(local)
	peer := []string{"197.155.0.0/28"}
	options.SetPeerCidrs(peer)
	result, returnValue, returnValueErr := vpcService.CreateVpnGatewayConnection(options)
	return result, returnValue, returnValueErr
}

// DeleteVpnGatewayConnection DELETE
// /vpn_gateways/{vpn_gateway_id}/connections/{id}
// Delete a VPN connection
func DeleteVpnGatewayConnection(vpcService *vpcv1.VpcV1, gatewayID, connID string) (*core.DetailedResponse, error) {
	options := &vpcv1.DeleteVpnGatewayConnectionOptions{}
	options.SetVpnGatewayID(gatewayID)
	options.SetID(connID)
	returnValue, returnValueErr := vpcService.DeleteVpnGatewayConnection(options)
	return returnValue, returnValueErr
}

// GetVpnGatewayConnection GET
// /vpn_gateways/{vpn_gateway_id}/connections/{id}
// Retrieve the specified VPN connection
func GetVpnGatewayConnection(vpcService *vpcv1.VpcV1, gatewayID, connID string) (*vpcv1.VPNGatewayConnection, *core.DetailedResponse, error) {
	options := &vpcv1.GetVpnGatewayConnectionOptions{}
	options.SetVpnGatewayID(gatewayID)
	options.SetID(connID)
	result, returnValue, returnValueErr := vpcService.GetVpnGatewayConnection(options)
	return result, returnValue, returnValueErr
}

// UpdateVpnGatewayConnection PATCH
// /vpn_gateways/{vpn_gateway_id}/connections/{id}
// Update a VPN connection
func UpdateVpnGatewayConnection(vpcService *vpcv1.VpcV1, gatewayID, connID, name string) (*vpcv1.VPNGatewayConnection, *core.DetailedResponse, error) {
	options := &vpcv1.UpdateVpnGatewayConnectionOptions{
		ID:           core.StringPtr(connID),
		VpnGatewayID: core.StringPtr(gatewayID),
		Name:         core.StringPtr(name),
	}
	result, returnValue, returnValueErr := vpcService.UpdateVpnGatewayConnection(options)
	return result, returnValue, returnValueErr
}

// ListVpnGatewayConnectionLocalCidrs GET
// /vpn_gateways/{vpn_gateway_id}/connections/{id}/local_cidrs
// List all local CIDRs for a resource
func ListVpnGatewayConnectionLocalCidrs(vpcService *vpcv1.VpcV1, gatewayID, connID string) (*vpcv1.VPNGatewayConnectionLocalCIDRs, *core.DetailedResponse, error) {
	options := &vpcv1.ListVpnGatewayConnectionLocalCidrsOptions{}
	options.SetVpnGatewayID(gatewayID)
	options.SetID(connID)
	result, returnValue, returnValueErr := vpcService.ListVpnGatewayConnectionLocalCidrs(options)
	return result, returnValue, returnValueErr
}

// DeleteVpnGatewayConnectionLocalCidr DELETE
// /vpn_gateways/{vpn_gateway_id}/connections/{id}/local_cidrs/{prefix_address}/{prefix_length}
// Remove a CIDR from a resource
func DeleteVpnGatewayConnectionLocalCidr(vpcService *vpcv1.VpcV1, gatewayID, connID, prefixAdd, prefixLen string) (*core.DetailedResponse, error) {
	options := &vpcv1.DeleteVpnGatewayConnectionLocalCidrOptions{}
	options.SetVpnGatewayID(gatewayID)
	options.SetID(connID)
	options.SetPrefixAddress(prefixAdd)
	options.SetPrefixLength(prefixLen)
	returnValue, returnValueErr := vpcService.DeleteVpnGatewayConnectionLocalCidr(options)
	return returnValue, returnValueErr
}

// GetVpnGatewayConnectionLocalCidr GET
// /vpn_gateways/{vpn_gateway_id}/connections/{id}/local_cidrs/{prefix_address}/{prefix_length}
// Check if a specific CIDR exists on a specific resource
func GetVpnGatewayConnectionLocalCidr(vpcService *vpcv1.VpcV1, gatewayID, connID, prefixAdd, prefixLen string) (*core.DetailedResponse, error) {
	options := &vpcv1.GetVpnGatewayConnectionLocalCidrOptions{}
	options.SetVpnGatewayID(gatewayID)
	options.SetID(connID)
	options.SetPrefixAddress(prefixAdd)
	options.SetPrefixLength(prefixLen)
	returnValue, returnValueErr := vpcService.GetVpnGatewayConnectionLocalCidr(options)
	return returnValue, returnValueErr
}

// SetVpnGatewayConnectionLocalCidr - PUT
// /vpn_gateways/{vpn_gateway_id}/connections/{id}/local_cidrs/{prefix_address}/{prefix_length}
// Set a CIDR on a resource
func SetVpnGatewayConnectionLocalCidr(vpcService *vpcv1.VpcV1, gatewayID, connID, prefixAdd, prefixLen string) (*core.DetailedResponse, error) {
	options := &vpcv1.SetVpnGatewayConnectionLocalCidrOptions{}
	options.SetVpnGatewayID(gatewayID)
	options.SetID(connID)
	options.SetPrefixAddress(prefixAdd)
	options.SetPrefixLength(prefixLen)
	returnValue, returnValueErr := vpcService.SetVpnGatewayConnectionLocalCidr(options)
	return returnValue, returnValueErr
}

// ListVpnGatewayConnectionPeerCidrs GET
// /vpn_gateways/{vpn_gateway_id}/connections/{id}/peer_cidrs
// List all peer CIDRs for a resource
func ListVpnGatewayConnectionPeerCidrs(vpcService *vpcv1.VpcV1, gatewayID, connID string) (*vpcv1.VPNGatewayConnectionPeerCIDRs, *core.DetailedResponse, error) {
	options := &vpcv1.ListVpnGatewayConnectionPeerCidrsOptions{}
	options.SetVpnGatewayID(gatewayID)
	options.SetID(connID)
	result, returnValue, returnValueErr := vpcService.ListVpnGatewayConnectionPeerCidrs(options)
	return result, returnValue, returnValueErr
}

// DeleteVpnGatewayConnectionPeerCidr DELETE
// /vpn_gateways/{vpn_gateway_id}/connections/{id}/peer_cidrs/{prefix_address}/{prefix_length}
// Remove a CIDR from a resource
func DeleteVpnGatewayConnectionPeerCidr(vpcService *vpcv1.VpcV1, gatewayID, connID, prefixAdd, prefixLen string) (*core.DetailedResponse, error) {
	options := &vpcv1.DeleteVpnGatewayConnectionPeerCidrOptions{}
	options.SetVpnGatewayID(gatewayID)
	options.SetID(connID)
	options.SetPrefixAddress(prefixAdd)
	options.SetPrefixLength(prefixLen)
	returnValue, returnValueErr := vpcService.DeleteVpnGatewayConnectionPeerCidr(options)
	return returnValue, returnValueErr
}

// GetVpnGatewayConnectionPeerCidr GET
// /vpn_gateways/{vpn_gateway_id}/connections/{id}/peer_cidrs/{prefix_address}/{prefix_length}
// Check if a specific CIDR exists on a specific resource
func GetVpnGatewayConnectionPeerCidr(vpcService *vpcv1.VpcV1, gatewayID, connID, prefixAdd, prefixLen string) (*core.DetailedResponse, error) {
	options := &vpcv1.GetVpnGatewayConnectionPeerCidrOptions{}
	options.SetVpnGatewayID(gatewayID)
	options.SetID(connID)
	options.SetPrefixAddress(prefixAdd)
	options.SetPrefixLength(prefixLen)
	returnValue, returnValueErr := vpcService.GetVpnGatewayConnectionPeerCidr(options)
	return returnValue, returnValueErr
}

// SetVpnGatewayConnectionPeerCidr - PUT
// /vpn_gateways/{vpn_gateway_id}/connections/{id}/peer_cidrs/{prefix_address}/{prefix_length}
// Set a CIDR on a resource
func SetVpnGatewayConnectionPeerCidr(vpcService *vpcv1.VpcV1, gatewayID, connID, prefixAdd, prefixLen string) (*core.DetailedResponse, error) {
	options := &vpcv1.SetVpnGatewayConnectionPeerCidrOptions{}
	options.SetVpnGatewayID(gatewayID)
	options.SetID(connID)
	options.SetPrefixAddress(prefixAdd)
	options.SetPrefixLength(prefixLen)
	returnValue, returnValueErr := vpcService.SetVpnGatewayConnectionPeerCidr(options)
	return returnValue, returnValueErr
}
