//go:build integration
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
	"flag"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/IBM/vpc-go-sdk/vpcv1"
)

var (
	Attached                  = "attached"
	bootVolAttachmentID       *string
	configLoaded              bool = false
	counter                        = Counter{0}
	createdACLID              *string
	createdACLRuleID          *string
	createdDhgID              *string
	createdDhID               *string
	createdFipID              *string
	createdImageID            *string
	createdInstanceID         *string
	createdPGWID              *string
	createdSecondVnicID       *string
	createdSgID               *string
	createdSgRuleID           *string
	createdSgVnicID           *string
	createdSSHKey             *string
	createdSubnetID           *string
	createdVnicID             *string
	createdFlowLogID          *string
	createdVolAttachmentID    *string
	createdVolumeID           *string
	createdVpcAddressPrefixID *string
	createdSubnetReservedIP   *string
	createdVpcID              *string
	createdVPCRouteID         *string
	createdEgwID              *string
	createdRtID               *string
	createdRt2ID              *string
	createdRouteID            *string
	defaultACLID              *string
	defaultImageID            *string
	defaultInstanceID         *string
	defaultInstanceProfile    *string
	defaultLBID               *string
	defaultLBListenerID       *string
	defaultLBListenerPolicyID *string
	defaultLBPoolID           *string
	defaultLBPoolMemberID     *string
	defaultLBRule             *string
	defaultOSName             *string
	defaultRegionName         *string
	defaultResourceGroupID    *string
	defaultSubnetID           *string
	defaultVnicID             *string
	defaultVolumeProfile      *string
	defaultVpcID              *string
	defaultZoneName           *string
	createdTemplateID         *string
	createdInstanceGroupID    *string
	createdIgPolicyID         *string
	createdIgManagerID        *string
	memberID                  *string
	lbProfile                 *string
	detailed                  = flag.Bool("detailed", false, "boolean")
	Running                   = "running"
	skipForMockTesting        = flag.Bool("skipForMockTesting", false, "boolean")
	Stopped                   = "stopped"
	testCount                 = flag.Bool("testCount", false, "boolean")
	timestamp                 = strconv.FormatInt(tunix, 10)
	tunix                     = time.Now().Unix()
)

const (
	externalConfigFile = "../vpc.env"
	skipMessage        = "External configuration could not be loaded, skipping..."
)

func shouldSkipTest(t *testing.T) {
	if !configLoaded {
		t.Skip(skipMessage)
	}
}

func createVpcService(t *testing.T) *vpcv1.VpcV1 {

	t.Run("Load Config", func(t *testing.T) {
		if _, err := os.Stat(externalConfigFile); err == nil {
			if err = os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile); err == nil {
				configLoaded = true
			}
		}
		shouldSkipTest(t)
	})
	var service = InstantiateVPCService()
	if service == nil {
		fmt.Println("Error creating VPC service.")
		t.Error("Error creating vpc service with error message:")
		return nil
	}
	t.Log("Success: VPC service creation complete.")
	return service
}

func TestVPCResources(t *testing.T) {
	vpcService := createVpcService(t)
	shouldSkipTest(t)

	t.Run("Geography", func(t *testing.T) {

		t.Run("All regions", func(t *testing.T) {
			res, _, err := ListRegions(vpcService)
			ValidateListResponse(t, res, err, GET, detailed, increment)
			defaultRegionName = res.Regions[0].Name
		})

		t.Run("Get region", func(t *testing.T) {
			res, _, err := GetRegion(vpcService, *defaultRegionName)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Zones within Region", func(t *testing.T) {
			if *skipForMockTesting {
				zone := "us-east" + "-1"
				defaultZoneName = &zone
				t.Skip("skipping test in travis.")
			}
			t.Run("Zones within Region", func(t *testing.T) {
				res, _, err := ListZones(vpcService, *defaultRegionName)
				ValidateListResponse(t, res, err, GET, detailed, increment)
				defaultZoneName = res.Zones[0].Name
			})

			t.Run("Get Zone", func(t *testing.T) {
				res, _, err := GetZone(vpcService, *defaultRegionName, *defaultZoneName)
				ValidateResponse(t, res, err, GET, detailed, increment)
			})
		})
	})

	t.Run("Create", func(t *testing.T) {

		t.Run("Initial Setup", func(t *testing.T) {

			// getting default resource group assuming there is atleast one VPC in the account.
			vpcs, _, err := ListVpcs(vpcService)
			if err != nil && vpcs == nil {
				fmt.Println("Error: ", err)
				t.Error("Error fetching for Resource Group with error message:", err)
				return
			}
			defaultResourceGroupID = vpcs.Vpcs[0].ResourceGroup.ID

			t.Run("List Instance Profiles", func(t *testing.T) {
				res, _, err := ListInstanceProfiles(vpcService)
				ValidateListResponse(t, res, err, GET, detailed, increment)
				defaultInstanceProfile = res.Profiles[0].Name
			})

			t.Run("List Volume Profiles", func(t *testing.T) {
				res, _, err := ListVolumeProfiles(vpcService)
				ValidateListResponse(t, res, err, GET, detailed, increment)
				defaultVolumeProfile = res.Profiles[0].Name
			})

			t.Run("Get Volume Profile", func(t *testing.T) {
				res, _, err := GetVolumeProfile(vpcService, *defaultVolumeProfile)
				ValidateResponse(t, res, err, GET, detailed, increment)
			})

			t.Run("Get Instance Profile", func(t *testing.T) {
				res, _, err := GetInstanceProfile(vpcService, *defaultInstanceProfile)
				ValidateResponse(t, res, err, GET, detailed, increment)
			})

			t.Run("List Images", func(t *testing.T) {
				res, _, err := ListImages(vpcService, "public")
				ValidateListResponse(t, res, err, GET, detailed, increment)
				defaultImageID = res.Images[0].ID
			})
			t.Run("List Operating Systems", func(t *testing.T) {
				res, _, err := ListOperatingSystems(vpcService)
				ValidateListResponse(t, res, err, GET, detailed, increment)
				defaultOSName = res.OperatingSystems[0].Name
			})

			t.Run("Get Operating System", func(t *testing.T) {
				res, _, err := GetOperatingSystem(vpcService, *defaultOSName)
				ValidateResponse(t, res, err, GET, detailed, increment)
			})

		})

		t.Run("Create VPC", func(t *testing.T) {
			name := getName("vpc")
			res, _, err := CreateVPC(vpcService, name, *defaultResourceGroupID)
			ValidateResponse(t, res, err, POST, detailed, increment)
			createdVpcID = res.ID
		})

		t.Run("Create Subnet", func(t *testing.T) {
			name := getName("subnet")
			res, _, err := CreateSubnet(vpcService, *createdVpcID, name, *defaultZoneName, *skipForMockTesting)
			ValidateResponse(t, res, err, POST, detailed, increment)
			createdSubnetID = res.ID
		})

		t.Run("Create SSH key", func(t *testing.T) {
			name := getName("key")
			res, _, err := CreateSSHKey(vpcService, name, "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCcPJwUpNQr0MplO6UM5mfV4vlvY0RpD6gcXqodzZIjsoG31+hQxoJVU9yQcSjahktHFs7Fk2Mo79jUT3wVC8Pg6A3//IDFkLjVrg/mQVpIf6+GxIYEtVg6Tk4pP3YNoksrugGlpJ4LCR3HMe3fBQTQqTzObbb0cSF6xhW5UBq8vhqIkhYKd3KLGJnnrwsIGcwb5BRk68ZFYhreAomvx4jWjaBFlH98HhE4wUEVvJLRy/qR/0w3XVjTSgOlhXywaAOEkmwye7kgSglegCpHWwYNly+NxLONjqbX9rHbFHUVRShnFKh2+M6XKE3HowT/3Y1lDd2PiVQpJY0oQmebiRxB astha.jain@ibm.com")
			ValidateResponse(t, res, err, POST, detailed, increment)
			createdSSHKey = res.ID
		})

		t.Run("Create Floating IP", func(t *testing.T) {
			name := getName("fip")
			res, _, err := CreateFloatingIP(vpcService, *defaultZoneName, name)
			ValidateResponse(t, res, err, POST, detailed, increment)
			createdFipID = res.ID
		})

		t.Run("Create Volume", func(t *testing.T) {
			name := getName("vol")
			res, _, err := CreateVolume(vpcService, name, *defaultVolumeProfile, *defaultZoneName, 10)
			ValidateResponse(t, res, err, POST, detailed, increment)
			createdVolumeID = res.ID
		})

		t.Run("Create Image", func(t *testing.T) {
			t.Skip("Skip Create Image")
			name := getName("img")
			res, _, err := CreateImage(vpcService, name)
			ValidateResponse(t, res, err, POST, detailed, increment)
			createdImageID = res.ID
		})

		t.Run("Create Instance", func(t *testing.T) {
			var profile *string
			mockProfile := "bc1-8x32"
			gtProfile := "bx2-4x16"
			if !*skipForMockTesting {
				profile = &gtProfile
			} else {
				profile = &mockProfile
			}
			name := getName("vsi")
			statusChanged := PollSubnet(vpcService, *createdSubnetID, "available", 4)
			if statusChanged {
				res, _, err := CreateInstance(vpcService, name, *profile, *defaultImageID, *defaultZoneName, *createdSubnetID, *createdSSHKey, *createdVpcID)
				ValidateResponse(t, res, err, POST, detailed, increment)
				createdInstanceID = res.ID
				createdVnicID = res.PrimaryNetworkInterface.ID
			}
		})

		t.Run("Create Instance template", func(t *testing.T) {
			var profile *string
			mockProfile := "bc1-8x32"
			gtProfile := "bx2-4x16"
			if !*skipForMockTesting {
				profile = &gtProfile
			} else {
				profile = &mockProfile
			}
			name := getName("template")
			res, _, err := CreateInstanceTemplate(vpcService, name, *defaultImageID, *profile, *defaultZoneName, *createdSubnetID, *createdVpcID)
			template := res.(*vpcv1.InstanceTemplate)
			ValidateResponse(t, template, err, POST, detailed, increment)
			createdTemplateID = template.ID
		})

		t.Run("Create Instance group", func(t *testing.T) {
			res, _, err := CreateInstanceGroup(vpcService, *createdTemplateID, getName("group"), *createdSubnetID, 1)
			ValidateResponse(t, res, err, POST, detailed, increment)
			createdInstanceGroupID = res.ID
		})

		t.Run("Create Instance group manager", func(t *testing.T) {
			res, _, err := CreateInstanceGroupManager(vpcService, *createdInstanceGroupID, getName("manager"))
			manager := res.(*vpcv1.InstanceGroupManager)
			ValidateResponse(t, manager, err, POST, detailed, increment)
			createdIgManagerID = manager.ID
		})

		t.Run("Create Instance group manager policy", func(t *testing.T) {
			res, _, err := CreateInstanceGroupManagerPolicy(vpcService, *createdInstanceGroupID, *createdIgManagerID, getName("manager"))
			managerPolicy := res.(*vpcv1.InstanceGroupManagerPolicy)
			ValidateResponse(t, managerPolicy, err, POST, detailed, increment)
			createdIgPolicyID = managerPolicy.ID
		})

	})

	t.Run("VPC Resources", func(t *testing.T) {

		t.Run("List VPCs", func(t *testing.T) {
			res, _, err := ListVpcs(vpcService)
			ValidateListResponse(t, res, err, GET, detailed, increment)
			defaultVpcID = res.Vpcs[0].ID
		})

		t.Run("List Subnets", func(t *testing.T) {
			res, _, err := ListSubnets(vpcService)
			ValidateListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("List Instances", func(t *testing.T) {
			res, _, err := ListInstances(vpcService)
			ValidateListResponse(t, res, err, GET, detailed, increment)
			defaultInstanceID = res.Instances[0].ID
			defaultVnicID = res.Instances[0].PrimaryNetworkInterface.ID
			defaultVpcID = res.Instances[0].VPC.ID
			defaultSubnetID = res.Instances[0].PrimaryNetworkInterface.Subnet.ID
		})

		t.Run("List SSH Keys", func(t *testing.T) {
			res, _, err := ListKeys(vpcService)
			ValidateListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("List Floating IPs", func(t *testing.T) {
			res, _, err := GetFloatingIPsList(vpcService)
			ValidateListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("List Volumes", func(t *testing.T) {
			res, _, err := ListVolumes(vpcService)
			ValidateListResponse(t, res, err, GET, detailed, increment)
		})
	})

	t.Run("Get a VPC Resource", func(t *testing.T) {

		t.Run("Get VPC", func(t *testing.T) {
			res, _, err := GetVPC(vpcService, *createdVpcID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get VPC Default Security Group", func(t *testing.T) {
			res, _, err := GetVPCDefaultSecurityGroup(vpcService, *createdVpcID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get VPC Default Network ACL", func(t *testing.T) {
			res, _, err := GetVPCDefaultACL(vpcService, *createdVpcID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get SSH Key", func(t *testing.T) {
			res, _, err := GetSSHKey(vpcService, *createdSSHKey)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Subnet", func(t *testing.T) {
			res, _, err := GetSubnet(vpcService, *createdSubnetID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Instance", func(t *testing.T) {
			statusChanged := PollInstance(vpcService, *createdInstanceID, Running, 7)
			if statusChanged {
				res, _, err := GetInstance(vpcService, *createdInstanceID)
				ValidateResponse(t, res, err, GET, detailed, increment)
			}
		})

		t.Run("Get Floating IP", func(t *testing.T) {
			res, _, err := GetFloatingIP(vpcService, *createdFipID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Volume", func(t *testing.T) {
			res, _, err := GetVolume(vpcService, *createdVolumeID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Image", func(t *testing.T) {
			res, _, err := GetImage(vpcService, *defaultImageID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

	})

	t.Run("Instances Network Attachments", func(t *testing.T) {
		t.Run("Get Initialization", func(t *testing.T) {
			res, _, err := GetInstanceInitialization(vpcService, *createdInstanceID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Network Interfaces", func(t *testing.T) {
			res, _, err := ListNetworkInterfaces(vpcService, *createdInstanceID)
			ValidateResponse(t, res, err, GET, detailed, increment)
			createdVnicID = res.NetworkInterfaces[0].ID
		})

		t.Run("Create Network Interfaces", func(t *testing.T) {
			res, _, err := CreateNetworkInterface(vpcService, *createdInstanceID, *createdSubnetID)
			ValidateResponse(t, res, err, GET, detailed, increment)
			createdSecondVnicID = res.ID
		})

		t.Run("Attach FIP to Vnic", func(t *testing.T) {
			res, _, err := CreateNetworkInterfaceFloatingIpBinding(vpcService, *createdInstanceID, *createdVnicID, *createdFipID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Network Interface", func(t *testing.T) {
			res, _, err := GetNetworkInterface(vpcService, *createdInstanceID, *createdVnicID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Update Network Interface", func(t *testing.T) {
			res, _, err := UpdateNetworkInterface(vpcService, *createdInstanceID, *createdVnicID, "vnic1")
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Vnic FLoating IPs", func(t *testing.T) {
			res, _, err := ListNetworkInterfaceFloatingIps(vpcService, *createdInstanceID, *createdVnicID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Vnic FLoating IP", func(t *testing.T) {
			res, _, err := GetNetworkInterfaceFloatingIp(vpcService, *createdInstanceID, *createdVnicID, *createdFipID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Delete Vnic FLoating IP", func(t *testing.T) {
			res, err := DeleteNetworkInterfaceFloatingIpBinding(vpcService, *createdInstanceID, *createdVnicID, *createdFipID)
			ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

		t.Run("Delete Network Interfaces", func(t *testing.T) {
			res, err := DeleteNetworkInterface(vpcService, *createdInstanceID, *createdSecondVnicID)
			ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

	})

	t.Run("Instances Volume Attachments", func(t *testing.T) {

		t.Run("Create Volume attachment", func(t *testing.T) {
			name := getName("vol-att")
			res, _, err := CreateVolumeAttachment(vpcService, *createdInstanceID, *createdVolumeID, name)
			ValidateResponse(t, res, err, POST, detailed, increment)
			createdVolAttachmentID = res.ID
		})

		t.Run("Get Volume attachments", func(t *testing.T) {
			res, _, err := ListVolumeAttachments(vpcService, *createdInstanceID)
			ValidateResponse(t, res, err, GET, detailed, increment)
			bootVolAttachmentID = res.VolumeAttachments[0].ID
		})

		t.Run("Get Volume attachment", func(t *testing.T) {
			res, _, err := GetVolumeAttachment(vpcService, *createdInstanceID, *createdVolAttachmentID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Update Volume attachments", func(t *testing.T) {
			name := getName("boot-att")
			res, _, err := UpdateVolumeAttachment(vpcService, *createdInstanceID, *bootVolAttachmentID, name)
			ValidateResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Delete Volume attachments", func(t *testing.T) {
			statusChanged := PollVolAttachment(vpcService, *createdInstanceID, *createdVolAttachmentID, Attached, 4)
			if statusChanged {
				res, err := DeleteVolumeAttachment(vpcService, *createdInstanceID, *createdVolAttachmentID)
				ValidateResponse(t, res, err, DELETE, detailed, increment)
			}
		})
	})

	t.Run("Subnet Bindings", func(t *testing.T) {

		t.Run("Set Subnet NetworkAcl Binding", func(t *testing.T) {
			acls, _, _ := ListNetworkAcls(vpcService)
			res, _, err := SetSubnetNetworkAclBinding(vpcService, *createdSubnetID, *acls.NetworkAcls[0].ID)
			ValidateResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Get Subnet NetworkAcl", func(t *testing.T) {
			res, _, err := GetSubnetNetworkAcl(vpcService, *createdSubnetID)
			ValidateResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Set Subnet Public Gateway Binding", func(t *testing.T) {
			name := getName("vol-att")
			pgw, _, _ := CreatePublicGateway(vpcService, name, *defaultVpcID, *defaultZoneName)
			res, _, err := CreateSubnetPublicGatewayBinding(vpcService, *createdSubnetID, *pgw.ID)
			ValidateResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Get Subnet Public Gateway", func(t *testing.T) {
			res, _, err := GetSubnetPublicGateway(vpcService, *createdSubnetID)
			ValidateResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Delete Subnet Public Gateway Binding", func(t *testing.T) {
			res, err := DeleteSubnetPublicGatewayBinding(vpcService, *createdSubnetID)
			ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

		t.Run("Create Subnet ReservedIps", func(t *testing.T) {
			name := getName("reservedIP")
			res, _, err := CreateSubnetReservedIP(vpcService, *createdSubnetID, name)
			createdSubnetReservedIP = res.ID
			ValidateResponse(t, res, err, POST, detailed, increment)
		})

		t.Run("List Subnet ReservedIps", func(t *testing.T) {
			res, _, err := ListSubnetReservedIps(vpcService, *createdSubnetID)
			ValidateListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Update Subnet ReservedIp", func(t *testing.T) {
			name := getName("reservedIP-2")
			res, _, err := UpdateSubnetReservedIP(vpcService, *createdSubnetID, *createdSubnetReservedIP, name)
			ValidateResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Get Subnet ReservedIp", func(t *testing.T) {
			res, _, err := GetSubnetReservedIP(vpcService, *createdSubnetID, *createdSubnetReservedIP)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Delete Subnet ReservedIp", func(t *testing.T) {
			res, err := DeleteSubnetReservedIP(vpcService, *createdSubnetID, *createdSubnetReservedIP)
			ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

	})

	t.Run("Update VPC Resources", func(t *testing.T) {

		t.Run("Update Floating IP", func(t *testing.T) {
			name := getName("fip-2")
			res, _, err := UpdateFloatingIP(vpcService, *createdFipID, name)
			ValidateResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Update Image", func(t *testing.T) {
			if !*skipForMockTesting {
				t.Skip("skip for stage testing")
			}
			name := getName("image-2")
			res, _, err := UpdateImage(vpcService, *defaultImageID, name)
			ValidateResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Update SSH key", func(t *testing.T) {
			name := getName("key-2")
			res, _, err := UpdateSSHKey(vpcService, *createdSSHKey, name)
			ValidateResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Update Instance", func(t *testing.T) {
			name := getName("vsi-2")
			statusChanged := PollInstance(vpcService, *createdInstanceID, Running, 4)
			if statusChanged {
				res, _, err := UpdateInstance(vpcService, *createdInstanceID, name)
				ValidateResponse(t, res, err, PATCH, detailed, increment)
			}
		})

		t.Run("Update Subnet", func(t *testing.T) {
			name := getName("subnet-2")
			res, _, err := UpdateSubnet(vpcService, *createdSubnetID, name)
			ValidateResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Update VPC", func(t *testing.T) {
			name := getName("vpc-2")
			res, _, err := UpdateVPC(vpcService, *createdVpcID, name)
			ValidateResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Update Volume", func(t *testing.T) {
			name := getName("vol-2")
			res, _, err := UpdateVolume(vpcService, *createdVolumeID, name)
			ValidateResponse(t, res, err, PATCH, detailed, increment)
		})
	})
	printTestSummary()
}

func TestVPCRoutes(t *testing.T) {
	vpcService := createVpcService(t)
	shouldSkipTest(t)

	if *createdVpcID == "" {
		res, _, err := ListInstances(vpcService)
		if err != nil {
			fmt.Println("Error: ", err)
			t.Error(err)
		}
		createdVpcID = res.Instances[0].VPC.ID
		defaultZoneName = res.Instances[0].Zone.Name
	}
	t.Run("VPC Routes", func(t *testing.T) {
		t.Run("Create VPC Route", func(t *testing.T) {
			name := getName("route-2")
			res, _, err := CreateVpcRoute(vpcService, *createdVpcID, *defaultZoneName, "5.5.0.0/16", "3.3.3.3.3", name)
			createdVPCRouteID = res.ID
			ValidateResponse(t, res, err, POST, detailed, increment)
		})

		t.Run("List VPC Routes", func(t *testing.T) {
			res, _, err := ListVpcRoutes(vpcService, *createdVpcID)
			ValidateListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Update VPC Route", func(t *testing.T) {
			name := getName("route-2")
			res, _, err := UpdateVpcRoute(vpcService, *createdVpcID, *createdVPCRouteID, name)
			ValidateResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Get VPC Route", func(t *testing.T) {
			res, _, err := GetVpcRoute(vpcService, *createdVpcID, *createdVPCRouteID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Delete VPC Route", func(t *testing.T) {
			res, err := DeleteVpcRoute(vpcService, *createdVpcID, *createdVPCRouteID)
			ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})
	})
}
func TestVPCAddressPrefix(t *testing.T) {
	vpcService := createVpcService(t)
	shouldSkipTest(t)

	if *createdVpcID == "" {
		res, _, err := ListInstances(vpcService)
		if err != nil {
			fmt.Println("Error: ", err)
			t.Error(err)
		}
		createdVpcID = res.Instances[0].VPC.ID
		defaultZoneName = res.Instances[0].Zone.Name
	}
	fmt.Println("herh", createdVpcID)
	t.Run("VPC address prefix", func(t *testing.T) {
		t.Run("List VPC Address Prefixes", func(t *testing.T) {
			res, _, err := ListVpcAddressPrefixes(vpcService, *createdVpcID)
			ValidateListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Create VPC Address Prefix", func(t *testing.T) {
			name := getName("addrprefix")
			res, _, err := CreateVpcAddressPrefix(vpcService, *createdVpcID, *defaultZoneName, "10.10.0.0/18", name)
			ValidateResponse(t, res, err, POST, detailed, increment)
			createdVpcAddressPrefixID = res.ID
		})

		t.Run("Get VPC Address Prefix", func(t *testing.T) {
			res, _, err := GetVpcAddressPrefix(vpcService, *createdVpcID, *createdVpcAddressPrefixID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})
		t.Run("Update VPC Address Prefixes", func(t *testing.T) {
			name := getName("addrprefix-2")
			res, _, err := UpdateVpcAddressPrefix(vpcService, *createdVpcID, *createdVpcAddressPrefixID, name)
			ValidateResponse(t, res, err, PATCH, detailed, increment)
		})
		t.Run("Delete VPC Address Prefixes", func(t *testing.T) {
			res, err := DeleteVpcAddressPrefix(vpcService, *createdVpcID, *createdVpcAddressPrefixID)
			ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})
	})
}
func TestVPCAccessControlLists(t *testing.T) {
	vpcService := createVpcService(t)
	shouldSkipTest(t)

	t.Run("ACL Resources", func(t *testing.T) {

		t.Run("List  ACLs", func(t *testing.T) {
			res, _, err := ListNetworkAcls(vpcService)
			ValidateListResponse(t, res, err, GET, detailed, increment)
			nacl := res.NetworkAcls[0]
			defaultACLID = nacl.ID
			defaultVpcID = nacl.VPC.ID
		})

		t.Run("Create ACL", func(t *testing.T) {
			name := getName("acl")
			res, _, err := CreateNetworkAcl(vpcService, name, *defaultACLID, *defaultVpcID)
			ValidateResponse(t, res, err, POST, detailed, increment)
			createdACLID = res.ID
		})

		t.Run("Create ACL Rule", func(t *testing.T) {
			name := getName("acl-rule")
			res, _, err := CreateNetworkAclRule(vpcService, name, *createdACLID)
			rule := res.(*vpcv1.NetworkACLRuleNetworkACLRuleProtocolAll)
			ValidateResponse(t, rule, err, POST, detailed, increment)
			createdACLRuleID = rule.ID
		})

		t.Run("List ACL Rules", func(t *testing.T) {
			res, _, err := ListNetworkAclRules(vpcService, *createdACLID)
			ValidateListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get ACL", func(t *testing.T) {
			res, _, err := GetNetworkAcl(vpcService, *createdACLID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get ACL Rules", func(t *testing.T) {
			res, _, err := GetNetworkAclRule(vpcService, *createdACLID, *createdACLRuleID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Update ACL", func(t *testing.T) {
			name := getName("acl-2")
			res, _, err := UpdateNetworkAcl(vpcService, *createdACLID, name)
			ValidateResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Update ACL Rule", func(t *testing.T) {
			name := getName("acl-rule-2")
			res, _, err := UpdateNetworkAclRule(vpcService, *createdACLID, *createdACLRuleID, name)
			ValidateResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Delete ACL Rule", func(t *testing.T) {
			res, err := DeleteNetworkAclRule(vpcService, *createdACLID, *createdACLRuleID)
			ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

		t.Run("Delete ACL", func(t *testing.T) {
			res, err := DeleteNetworkAcl(vpcService, *createdACLID)
			ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})
	})
	printTestSummary()
}
func TestVPCSecurityGroups(t *testing.T) {
	vpcService := createVpcService(t)
	shouldSkipTest(t)

	res, _, err := ListInstances(vpcService)
	if err != nil {
		fmt.Println("Error: ", err)
		t.Error(err)
	}
	var defaultVpcID = res.Instances[0].VPC.ID
	var defaultVnicID = res.Instances[0].PrimaryNetworkInterface.ID
	t.Run("SG Resources", func(t *testing.T) {

		var sgID *string
		t.Run("List Security Groups", func(t *testing.T) {
			res, _, err := ListSecurityGroups(vpcService)
			ValidateListResponse(t, res, err, GET, detailed, increment)
			sgID = res.SecurityGroups[0].ID
		})

		t.Run("List Security Group Network Interfaces", func(t *testing.T) {
			res, _, err := ListSecurityGroupNetworkInterfaces(vpcService, *sgID)
			ValidateListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("List Security Group Rules", func(t *testing.T) {
			res, _, err := ListSecurityGroupRules(vpcService, *sgID)
			ValidateListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Create Security Group", func(t *testing.T) {
			name := getName("sg")
			res, _, err := CreateSecurityGroup(vpcService, name, *defaultVpcID)
			ValidateResponse(t, res, err, POST, detailed, increment)
			createdSgID = res.ID
		})

		t.Run("Create Security Group Network Interface", func(t *testing.T) {
			res, _, err := CreateSecurityGroupNetworkInterfaceBinding(vpcService, *createdSgID, *defaultVnicID)
			ValidateResponse(t, res, err, POST, detailed, increment)
			createdSgVnicID = res.ID
		})

		t.Run("Create Security Group Rule", func(t *testing.T) {
			res, _, err := CreateSecurityGroupRule(vpcService, *createdSgID)
			sgRule := res.(*vpcv1.SecurityGroupRuleSecurityGroupRuleProtocolAll)
			ValidateResponse(t, sgRule, err, POST, detailed, increment)
			createdSgRuleID = sgRule.ID
		})

		t.Run("Get Security Group", func(t *testing.T) {
			res, _, err := GetSecurityGroup(vpcService, *sgID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Security Group Network Interface", func(t *testing.T) {
			res, _, err := GetSecurityGroupNetworkInterface(vpcService, *createdSgID, *defaultVnicID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Security Group Rules", func(t *testing.T) {
			res, _, err := GetSecurityGroupRule(vpcService, *createdSgID, *createdSgRuleID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Update Security Groups", func(t *testing.T) {
			name := getName("sg-2")
			res, _, err := UpdateSecurityGroup(vpcService, *createdSgID, name)
			ValidateResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Update Security Group Rule", func(t *testing.T) {
			res, _, err := UpdateSecurityGroupRule(vpcService, *createdSgID, *createdSgRuleID)
			ValidateResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Delete Security Group Network Interface", func(t *testing.T) {
			res, err := DeleteSecurityGroupNetworkInterfaceBinding(vpcService, *createdSgID, *defaultVnicID)
			ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

		t.Run("Delete Security Group Rule", func(t *testing.T) {
			res, err := DeleteSecurityGroupRule(vpcService, *createdSgID, *createdSgRuleID)
			ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

		t.Run("Delete Security Group", func(t *testing.T) {
			res, err := DeleteSecurityGroup(vpcService, *createdSgID)
			ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})
	})
	printTestSummary()
}
func TestVPCPublicGateways(t *testing.T) {
	vpcService := createVpcService(t)
	shouldSkipTest(t)

	res, _, err := ListInstances(vpcService)
	if err != nil {
		fmt.Println("Error: ", err)
		t.Error(err)
	}
	var defaultVpcID = res.Instances[0].VPC.ID
	var defaultZoneName = res.Instances[0].Zone.Name
	t.Run("PGW Resources", func(t *testing.T) {

		t.Run("Create Public Gateway", func(t *testing.T) {
			name := getName("pgw")
			res, _, err := CreatePublicGateway(vpcService, name, *defaultVpcID, *defaultZoneName)
			ValidateResponse(t, res, err, POST, detailed, increment)
			createdPGWID = res.ID
		})

		t.Run("List  Public Gateways", func(t *testing.T) {
			res, _, err := ListPublicGateways(vpcService)
			ValidateListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Public Gateway", func(t *testing.T) {
			res, _, err := GetPublicGateway(vpcService, *createdPGWID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Update Public Gateway", func(t *testing.T) {
			name := getName("pgw-2")
			res, _, err := UpdatePublicGateway(vpcService, *createdPGWID, name)
			ValidateResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Delete Public Gateway", func(t *testing.T) {
			res, err := DeletePublicGateway(vpcService, *createdPGWID)
			ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

	})
	printTestSummary()
}
func TestVPCAutoscale(t *testing.T) {
	vpcService := createVpcService(t)
	shouldSkipTest(t)

	t.Run("List Instance Templates", func(t *testing.T) {
		res, _, err := ListInstanceTemplates(vpcService)
		ValidateListResponse(t, res, err, GET, detailed, increment)
	})

	t.Run("Get Instance Template", func(t *testing.T) {
		res, _, err := GetInstanceTemplate(vpcService, *createdTemplateID)
		ValidateResponse(t, res, err, GET, detailed, increment)
	})

	t.Run("Update Instance Template", func(t *testing.T) {
		res, _, err := UpdateInstanceTemplate(vpcService, *createdTemplateID, getName("template-2"))
		ValidateResponse(t, res, err, GET, detailed, increment)
	})

	t.Run("List Instance Groups", func(t *testing.T) {
		res, _, err := ListInstanceGroups(vpcService)
		ValidateListResponse(t, res, err, GET, detailed, increment)
	})

	t.Run("Get Instance Groups", func(t *testing.T) {
		res, _, err := GetInstanceGroup(vpcService, *createdInstanceGroupID)
		ValidateResponse(t, res, err, GET, detailed, increment)
	})

	t.Run("Update Instance Groups", func(t *testing.T) {
		res, _, err := UpdateInstanceGroup(vpcService, *createdInstanceGroupID, getName("ig-2"))
		ValidateResponse(t, res, err, GET, detailed, increment)
	})

	t.Run("Delete IG Load balancer", func(t *testing.T) {
		res, err := DeleteInstanceGroupLoadBalancer(vpcService, *createdInstanceGroupID)
		ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
	})

	t.Run("List Managers", func(t *testing.T) {
		res, _, err := ListInstanceGroupManagers(vpcService, *createdInstanceGroupID)
		ValidateListResponse(t, res, err, GET, detailed, increment)
	})

	t.Run("Get Instance Groups Manager", func(t *testing.T) {
		res, _, err := GetInstanceGroupManager(vpcService, *createdInstanceGroupID, *createdIgManagerID, getName("igm"))
		ValidateResponse(t, res, err, GET, detailed, increment)
	})

	t.Run("Update Instance Groups Manager", func(t *testing.T) {
		res, _, err := UpdateInstanceGroupManager(vpcService, *createdInstanceGroupID, *createdIgManagerID, getName("igm-2"))
		ValidateResponse(t, res, err, GET, detailed, increment)
	})

	t.Run("List Manager Policies", func(t *testing.T) {
		res, _, err := ListInstanceGroupManagerPolicies(vpcService, *createdInstanceGroupID, *createdIgManagerID)
		ValidateListResponse(t, res, err, GET, detailed, increment)
	})

	t.Run("Get Instance Groups Manager Policy", func(t *testing.T) {
		res, _, err := GetInstanceGroupManagerPolicy(vpcService, *createdInstanceGroupID, *createdIgManagerID, *createdIgPolicyID)
		ValidateResponse(t, res, err, GET, detailed, increment)
	})

	t.Run("Update Instance Groups Policy", func(t *testing.T) {
		res, _, err := UpdateInstanceGroupManagerPolicy(vpcService, *createdInstanceGroupID, *createdIgManagerID, *createdIgPolicyID, getName("igmp-2"))
		ValidateResponse(t, res, err, GET, detailed, increment)
	})
	t.Run("Get Instance Groups Memberships", func(t *testing.T) {
		res, _, err := ListInstanceGroupMemberships(vpcService, *createdInstanceGroupID)
		ValidateResponse(t, res, err, GET, detailed, increment)
		memberID = res.Memberships[0].ID
	})

	t.Run("Get Instance Groups Membership", func(t *testing.T) {
		res, _, err := GetInstanceGroupMembership(vpcService, *createdInstanceGroupID, *memberID)
		ValidateResponse(t, res, err, GET, detailed, increment)
	})

	t.Run("Update Instance Groups Membership", func(t *testing.T) {
		res, _, err := UpdateInstanceGroupMembership(vpcService, *createdInstanceGroupID, *memberID, getName("member"))
		ValidateResponse(t, res, err, GET, detailed, increment)
	})

	t.Run("Delete Instance Groups Membership", func(t *testing.T) {
		res, err := DeleteInstanceGroupMembership(vpcService, *createdInstanceGroupID, *memberID)
		ValidateResponse(t, res, err, GET, detailed, increment)
	})

	t.Run("Get Instance Groups Memberships", func(t *testing.T) {
		res, err := DeleteInstanceGroupMemberships(vpcService, *createdInstanceGroupID)
		ValidateResponse(t, res, err, GET, detailed, increment)
	})

}
func TestVPCLoadBalancers(t *testing.T) {
	vpcService := createVpcService(t)
	shouldSkipTest(t)

	t.Run("LB Resources", func(t *testing.T) {
		var subnetID *string
		res, _, err := ListInstances(vpcService)
		if len(res.Instances) == 0 && err == nil {
			t.Error("Error retrieving subnet ID")
			return
		}
		subnetID = res.Instances[0].PrimaryNetworkInterface.Subnet.ID
		t.Run("List Load Balancers Profiles", func(t *testing.T) {
			res, _, err := ListLoadBalancerProfiles(vpcService)
			ValidateListResponse(t, res, err, GET, detailed, increment)
			lbProfile = res.Profiles[0].Name
		})

		t.Run("Get Load Balancer Profile", func(t *testing.T) {
			res, _, err := GetLoadBalancerProfile(vpcService, *lbProfile)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("List Load Balancers", func(t *testing.T) {
			res, _, err := ListLoadBalancers(vpcService)
			ValidateListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Create Load Balancer", func(t *testing.T) {
			name := "gosdk-lb-" + strconv.FormatInt(tunix, 10)
			res, _, err := CreateLoadBalancer(vpcService, name, *subnetID)
			ValidateResponse(t, res, err, POST, detailed, increment)
			defaultLBID = res.ID
		})

		t.Run("List Load Balancer Listeners", func(t *testing.T) {
			statusChanged := PollLB(vpcService, *defaultLBID, "active", 8)
			if statusChanged {
				res, _, err := ListLoadBalancerListeners(vpcService, *defaultLBID)
				ValidateListResponse(t, res, err, GET, detailed, increment)
			}
		})
		t.Run("Create Load Balancer Listener", func(t *testing.T) {
			statusChanged := PollLB(vpcService, *defaultLBID, "active", 8)
			if statusChanged {
				res, _, err := CreateLoadBalancerListener(vpcService, *defaultLBID)
				ValidateResponse(t, res, err, POST, detailed, increment)
				defaultLBListenerID = res.ID
			}
		})

		t.Run("Get Load Balancer", func(t *testing.T) {
			statusChanged := PollLB(vpcService, *defaultLBID, "active", 8)
			if statusChanged {
				res, _, err := GetLoadBalancer(vpcService, *defaultLBID)
				ValidateResponse(t, res, err, GET, detailed, increment)
			}
		})

		t.Run("Create Load Balancer Listener Policy", func(t *testing.T) {
			statusChanged := PollLB(vpcService, *defaultLBID, "active", 5)
			if statusChanged {
				res, _, err := CreateLoadBalancerListenerPolicy(vpcService, *defaultLBID, *defaultLBListenerID)
				ValidateResponse(t, res, err, POST, detailed, increment)
				defaultLBListenerPolicyID = res.ID
			}
		})

		t.Run("Create Load Balancer Listener Policy Rule", func(t *testing.T) {
			statusChanged := PollLB(vpcService, *defaultLBID, "active", 5)
			if statusChanged {
				res, _, err := CreateLoadBalancerListenerPolicyRule(vpcService, *defaultLBID, *defaultLBListenerID, *defaultLBListenerPolicyID)
				ValidateResponse(t, res, err, POST, detailed, increment)
				defaultLBRule = res.ID
			}
		})
		var poolID *string
		t.Run("Create Load Balancer Pool", func(t *testing.T) {
			statusChanged := PollLB(vpcService, *defaultLBID, "active", 8)
			if statusChanged {
				name := "gsdk-lbpool-" + timestamp
				res, _, err := CreateLoadBalancerPool(vpcService, *defaultLBID, name)
				ValidateResponse(t, res, err, POST, detailed, increment)
				defaultLBPoolID = res.ID
			}
			statusChanged = PollLB(vpcService, *defaultLBID, "active", 8)
			if statusChanged {
				name := "go-lb-pool-2-" + timestamp
				res, _, err := CreateLoadBalancerPool(vpcService, *defaultLBID, name)
				ValidateResponse(t, res, err, POST, detailed, increment)
				poolID = res.ID
			}
		})

		t.Run("List Load Balancer Listeners Policies", func(t *testing.T) {
			res, _, err := ListLoadBalancerListenerPolicies(vpcService, *defaultLBID, *defaultLBListenerID)
			ValidateListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("List Load Balancer Listeners Policy Rules", func(t *testing.T) {
			res, _, err := ListLoadBalancerListenerPolicyRules(vpcService, *defaultLBID, *defaultLBListenerID, *defaultLBListenerPolicyID)
			ValidateListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Load Balancer Statistics", func(t *testing.T) {
			res, _, err := GetLoadBalancerStatistics(vpcService, *defaultLBID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Load Balancer Listener", func(t *testing.T) {
			res, _, err := GetLoadBalancerListener(vpcService, *defaultLBID, *defaultLBListenerID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Load Balancer Listener Policy", func(t *testing.T) {
			res, _, err := GetLoadBalancerListenerPolicy(vpcService, *defaultLBID, *defaultLBListenerID, *defaultLBListenerPolicyID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Load Balancer Listener Policy Rule", func(t *testing.T) {
			res, _, err := GetLoadBalancerListenerPolicyRule(vpcService, *defaultLBID, *defaultLBListenerID, *defaultLBListenerPolicyID, *defaultLBRule)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Update Load Balancer Listener Policy Rule", func(t *testing.T) {
			res, _, err := UpdateLoadBalancerListenerPolicyRule(vpcService, *defaultLBID, *defaultLBListenerID, *defaultLBListenerPolicyID, *defaultLBRule)
			ValidateResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Update Load Balancer Listener Policy", func(t *testing.T) {
			res, _, err := UpdateLoadBalancerListenerPolicy(vpcService, *defaultLBID, *defaultLBListenerID, *defaultLBListenerPolicyID, *poolID)
			ValidateResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Update Load Balancer Listener", func(t *testing.T) {
			res, _, err := UpdateLoadBalancerListener(vpcService, *defaultLBID, *defaultLBListenerID)
			ValidateResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Update Load Balancer", func(t *testing.T) {
			name := "gsdk-lb-2-" + timestamp
			res, _, err := UpdateLoadBalancer(vpcService, *defaultLBID, name)
			ValidateResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Delete Load Balancer listener Policy Rule", func(t *testing.T) {
			statusChanged := PollLB(vpcService, *defaultLBID, "active", 5)
			if statusChanged {
				res, err := DeleteLoadBalancerListenerPolicyRule(vpcService, *defaultLBID, *defaultLBListenerID, *defaultLBListenerPolicyID, *defaultLBRule)
				ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
			}
		})

		t.Run("Delete Load Balancer listener Policy", func(t *testing.T) {
			statusChanged := PollLB(vpcService, *defaultLBID, "active", 5)
			if statusChanged {
				res, err := DeleteLoadBalancerListenerPolicy(vpcService, *defaultLBID, *defaultLBListenerID, *defaultLBListenerPolicyID)
				ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
			}
		})

		t.Run("Delete Load Balancer listener", func(t *testing.T) {
			statusChanged := PollLB(vpcService, *defaultLBID, "active", 5)
			if statusChanged {
				res, err := DeleteLoadBalancerListener(vpcService, *defaultLBID, *defaultLBListenerID)
				ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
			}
		})

		t.Run("Create Load Balancer Pool Member", func(t *testing.T) {
			statusChanged := PollLB(vpcService, *defaultLBID, "active", 5)
			if statusChanged {
				res, _, err := CreateLoadBalancerPoolMember(vpcService, *defaultLBID, *defaultLBPoolID)
				ValidateListResponse(t, res, err, POST, detailed, increment)
				defaultLBPoolMemberID = res.ID
			}
		})

		t.Run("List Load Balancer Pools", func(t *testing.T) {
			res, _, err := ListLoadBalancerPools(vpcService, *defaultLBID)
			ValidateListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("List Load Balancer Pool Members", func(t *testing.T) {
			res, _, err := ListLoadBalancerPoolMembers(vpcService, *defaultLBID, *defaultLBPoolID)
			ValidateListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Load Balancer Pool", func(t *testing.T) {
			res, _, err := GetLoadBalancerPool(vpcService, *defaultLBID, *defaultLBPoolID)
			ValidateListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Load Balancer Pool Member", func(t *testing.T) {
			res, _, err := GetLoadBalancerPoolMember(vpcService, *defaultLBID, *defaultLBPoolID, *defaultLBPoolMemberID)
			ValidateListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Update Load Balancer Pool Member", func(t *testing.T) {
			res, _, err := UpdateLoadBalancerPoolMember(vpcService, *defaultLBID, *defaultLBPoolID, *defaultLBPoolMemberID)
			ValidateResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Update Load Balancer Pool", func(t *testing.T) {
			res, _, err := UpdateLoadBalancerPool(vpcService, *defaultLBID, *defaultLBPoolID)
			ValidateResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Delete Load Balancer Pool Member", func(t *testing.T) {
			statusChanged := PollLB(vpcService, *defaultLBID, "active", 5)
			if statusChanged {
				res, err := DeleteLoadBalancerPoolMember(vpcService, *defaultLBID, *defaultLBPoolID, *defaultLBPoolMemberID)
				ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
			}
		})

		var newPoolMemberID *string
		t.Run("Update Load Balancer Add Pool Member", func(t *testing.T) {
			statusChanged := PollLB(vpcService, *defaultLBID, "active", 5)
			if statusChanged {
				res, _, err := UpdateLoadBalancerPoolMembers(vpcService, *defaultLBID, *defaultLBPoolID)
				ValidateResponse(t, res, err, PATCH, detailed, increment)
				newPoolMemberID = res.Members[0].ID
			}
		})

		t.Run("Delete Load Balancer Pool Member Added ", func(t *testing.T) {
			statusChanged := PollLB(vpcService, *defaultLBID, "active", 5)
			if statusChanged {
				res, err := DeleteLoadBalancerPoolMember(vpcService, *defaultLBID, *defaultLBPoolID, *newPoolMemberID)
				ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
			}
		})

		t.Run("Delete Load Balancer Pool", func(t *testing.T) {
			statusChanged := PollLB(vpcService, *defaultLBID, "active", 5)
			if statusChanged {
				res, err := DeleteLoadBalancerPool(vpcService, *defaultLBID, *defaultLBPoolID)
				ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
			}
		})
		t.Run("Delete Load Balancer", func(t *testing.T) {
			statusChanged := PollLB(vpcService, *defaultLBID, "active", 5)
			if statusChanged {
				res, err := DeleteLoadBalancer(vpcService, *defaultLBID)
				ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
			}
		})
	})
	printTestSummary()
}

func TestVPCVPN(t *testing.T) {
	vpcService := createVpcService(t)
	shouldSkipTest(t)

	res, _, err := ListSubnets(vpcService)
	if err != nil {
		fmt.Println("Error: ", err)
		t.Error(err)
	}
	var defaultSubnetID = res.Subnets[0].ID
	var createdIkePolicyID *string
	var createdIpsecPolicyID *string
	var createdVpnGatewayID *string
	var createdVpnGatewayConnID *string
	t.Run("VPC Resources", func(t *testing.T) {

		t.Run("Create Ike Policy", func(t *testing.T) {
			name := "go-ike-1-" + strconv.FormatInt(tunix, 10)
			res, _, err := CreateIkePolicy(vpcService, name)
			ValidateResponse(t, res, err, POST, detailed, increment)
			createdIkePolicyID = res.ID
		})

		t.Run("Create Ipsec Policy", func(t *testing.T) {
			name := "go-ipsec-1-" + strconv.FormatInt(tunix, 10)
			res, _, err := CreateIpsecPolicy(vpcService, name)
			ValidateResponse(t, res, err, POST, detailed, increment)
			createdIpsecPolicyID = res.ID
		})

		t.Run("Create VPN Gateway", func(t *testing.T) {
			name := "go-vpngateway-1-" + strconv.FormatInt(tunix, 10)
			res, _, err := CreateVPNGateway(vpcService, *defaultSubnetID, name)
			vpn := res.(*vpcv1.VPNGateway)
			ValidateResponse(t, vpn, err, POST, detailed, increment)
			createdVpnGatewayID = vpn.ID
		})

		t.Run("Create Vpn Gateway Connections", func(t *testing.T) {
			name := "go-vpngateway-conn-1-" + strconv.FormatInt(tunix, 10)
			statusChanged := PollVPNGateway(vpcService, *createdVpnGatewayID, "available", 10)
			if statusChanged {
				res, _, err := CreateVPNGatewayConnection(vpcService, *createdVpnGatewayID, name)
				vpnGatewayConnection := res.(*vpcv1.VPNGatewayConnection)
				ValidateResponse(t, vpnGatewayConnection, err, POST, detailed, increment)
				createdVpnGatewayConnID = vpnGatewayConnection.ID
			}
		})

		t.Run("List Ike Policies", func(t *testing.T) {
			res, _, err := ListIkePolicies(vpcService)
			ValidateListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("List Ipsec Policies", func(t *testing.T) {
			res, _, err := ListIpsecPolicies(vpcService)
			ValidateListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("List Vpn Gateway", func(t *testing.T) {
			res, _, err := ListVPNGateways(vpcService)
			ValidateListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("List Vpn Gateway Connections", func(t *testing.T) {
			res, _, err := ListVPNGatewayConnections(vpcService, *createdVpnGatewayID)
			ValidateListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Ike Policies", func(t *testing.T) {
			res, _, err := GetIkePolicy(vpcService, *createdIkePolicyID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Ipsec Policies", func(t *testing.T) {
			res, _, err := GetIpsecPolicy(vpcService, *createdIpsecPolicyID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Vpn Gateway", func(t *testing.T) {
			res, _, err := GetVPNGateway(vpcService, *createdVpnGatewayID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get VpnGateway Connection", func(t *testing.T) {
			res, _, err := GetVPNGatewayConnection(vpcService, *createdVpnGatewayID, *createdVpnGatewayConnID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("List VpnGateway Ipsec Policy Connections", func(t *testing.T) {
			res, _, err := ListIpsecPolicyConnections(vpcService, *createdIpsecPolicyID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("List VpnGateway Ike Policy Connections", func(t *testing.T) {
			res, _, err := ListVPNGatewayIkePolicyConnections(vpcService, *createdIkePolicyID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("List VpnGateway Connection Local CIDRs", func(t *testing.T) {
			res, _, err := ListVPNGatewayConnectionLocalCIDRs(vpcService, *createdVpnGatewayID, *createdVpnGatewayConnID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("List VpnGateway Connection Peer CIDRs", func(t *testing.T) {
			res, _, err := ListVPNGatewayConnectionPeerCIDRs(vpcService, *createdVpnGatewayID, *createdVpnGatewayConnID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Vpn Gateway Connection LocalCidr", func(t *testing.T) {
			res, err := CheckVPNGatewayConnectionLocalCidr(vpcService, *createdVpnGatewayID, *createdVpnGatewayConnID, "192.132.0.0", "28")
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get VpnGateway Connection PeerCidr", func(t *testing.T) {
			res, err := CheckVPNGatewayConnectionPeerCidr(vpcService, *createdVpnGatewayID, *createdVpnGatewayConnID, "197.155.0.0", "28")
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("SetVpnGatewayConnectionLocalCidr", func(t *testing.T) {
			res, err := SetVPNGatewayConnectionLocalCidr(vpcService, *createdVpnGatewayID, *createdVpnGatewayConnID, "192.134.0.0", "28")
			ValidateResponse(t, res, err, GET, detailed, increment)
		})
		t.Run("GetVpnGatewayConnectionLocalCidr", func(t *testing.T) {
			res, err := CheckVPNGatewayConnectionLocalCidr(vpcService, *createdVpnGatewayID, *createdVpnGatewayConnID, "192.134.0.0", "28")
			ValidateResponse(t, res, err, GET, detailed, increment)
		})
		t.Run("DeleteVpnGatewayConnectionLocalCidr", func(t *testing.T) {
			res, err := DeleteVPNGatewayConnectionLocalCidr(vpcService, *createdVpnGatewayID, *createdVpnGatewayConnID, "192.134.0.0", "28")
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("SetVpnGatewayConnectionPeerCidr", func(t *testing.T) {
			res, err := SetVPNGatewayConnectionPeerCidr(vpcService, *createdVpnGatewayID, *createdVpnGatewayConnID, "192.157.0.0", "28")
			ValidateResponse(t, res, err, GET, detailed, increment)
		})
		t.Run("GetVpnGatewayConnectionPeerCidr", func(t *testing.T) {
			res, err := CheckVPNGatewayConnectionPeerCidr(vpcService, *createdVpnGatewayID, *createdVpnGatewayConnID, "192.157.0.0", "28")
			ValidateResponse(t, res, err, GET, detailed, increment)
		})
		t.Run("DeleteVpnGatewayConnectionPeerCidr", func(t *testing.T) {
			res, err := DeleteVPNGatewayConnectionPeerCidr(vpcService, *createdVpnGatewayID, *createdVpnGatewayConnID, "192.157.0.0", "28")
			ValidateResponse(t, res, err, GET, detailed, increment)
		})
		t.Run("Update Ike Policies", func(t *testing.T) {
			res, _, err := UpdateIkePolicy(vpcService, *createdIkePolicyID)
			ValidateResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Update Ipsec Policies", func(t *testing.T) {
			res, _, err := UpdateIpsecPolicy(vpcService, *createdIpsecPolicyID)
			ValidateResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Update VPN gateway", func(t *testing.T) {
			res, _, err := UpdateVPNGateway(vpcService, *createdVpnGatewayID, "go-vpngateway-2")
			ValidateResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Update VpnGateway Connection", func(t *testing.T) {
			res, _, err := UpdateVPNGatewayConnection(vpcService, *createdVpnGatewayID, *createdVpnGatewayConnID, "go-vpngateway-connection-2")
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Delete VpnGateway Connection", func(t *testing.T) {
			res, err := DeleteVPNGatewayConnection(vpcService, *createdVpnGatewayID, *createdVpnGatewayConnID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Delete Ike Policies", func(t *testing.T) {
			res, err := DeleteIkePolicy(vpcService, *createdIkePolicyID)
			ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

		t.Run("Delete Ipsec Policies", func(t *testing.T) {
			res, err := DeleteIpsecPolicy(vpcService, *createdIpsecPolicyID)
			ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

		t.Run("Delete VPN gateway", func(t *testing.T) {
			res, err := DeleteVPNGateway(vpcService, *createdVpnGatewayID)
			ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})
	})
	printTestSummary()
}

func TestVPCFlowLogs(t *testing.T) {
	vpcService := createVpcService(t)
	shouldSkipTest(t)

	if *defaultVpcID == "" {
		res, _, err := ListInstances(vpcService)
		if err != nil {
			fmt.Println("Error: ", err)
			t.Error(err)
		}
		defaultVpcID = res.Instances[0].VPC.ID
	}
	t.Run("Flow Logs", func(t *testing.T) {

		t.Run("Create Flow Log", func(t *testing.T) {
			name := "gsdk-fl-" + timestamp
			res, _, err := CreateFlowLogCollector(vpcService, name, "bucket-name", *defaultVpcID)
			ValidateResponse(t, res, err, POST, detailed, increment)
			createdFlowLogID = res.ID
		})

		t.Run("List Flow Logs", func(t *testing.T) {
			res, _, err := ListFlowLogCollectors(vpcService)
			ValidateListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Flow Log", func(t *testing.T) {
			res, _, err := GetFlowLogCollector(vpcService, *createdFlowLogID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Update Flow Log", func(t *testing.T) {
			name := "gsdk-fl-2-" + timestamp
			res, _, err := UpdateFlowLogCollector(vpcService, *createdFlowLogID, name)
			ValidateResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Delete Flow Log", func(t *testing.T) {
			res, err := DeleteFlowLogCollector(vpcService, *createdFlowLogID)
			ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})
	})
	printTestSummary()
}

func TestVPCEndpointGateways(t *testing.T) {
	vpcService := createVpcService(t)
	if *defaultVpcID == "" {
		res, _, err := ListInstances(vpcService)
		if err != nil {
			fmt.Println("Error: ", err)
			t.Error(err)
		}
		defaultVpcID = res.Instances[0].VPC.ID
	}
	t.Run("Endpoint Gateways", func(t *testing.T) {

		t.Run("Create Endpoint Gateway", func(t *testing.T) {
			res, _, err := CreateEndpointGateway(vpcService, *createdVpcID)
			ValidateResponse(t, res, err, POST, detailed, increment)
			createdEgwID = res.ID
		})

		t.Run("List Endpoint Gateways", func(t *testing.T) {
			res, _, err := ListEndpointGateways(vpcService)
			ValidateListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Endpoint Gateway", func(t *testing.T) {
			res, _, err := GetEndpointGateway(vpcService, *createdEgwID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Update Endpoint Gateway", func(t *testing.T) {
			name := "gsdk-egw-" + timestamp
			res, _, err := UpdateEndpointGateway(vpcService, *createdEgwID, name)
			ValidateResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("List Endpoint Gateway IPs", func(t *testing.T) {
			res, _, err := ListEndpointGatewayIps(vpcService, *createdEgwID)
			ValidateListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Create Subnet ReservedIps", func(t *testing.T) {
			name := getName("reservedIP")
			res, _, err := CreateSubnetReservedIP(vpcService, *createdSubnetID, name)
			createdSubnetReservedIP = res.ID
			ValidateResponse(t, res, err, POST, detailed, increment)
		})

		t.Run("Put Endpoint Gateway IP", func(t *testing.T) {
			res, _, err := AddEndpointGatewayIP(vpcService, *createdEgwID, *createdSubnetReservedIP)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Endpoint Gateway IP", func(t *testing.T) {
			res, _, err := GetEndpointGatewayIP(vpcService, *createdEgwID, *createdSubnetReservedIP)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Remove Endpoint Gateway IP", func(t *testing.T) {
			res, err := RemoveEndpointGatewayIP(vpcService, *createdEgwID, *createdSubnetReservedIP)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Delete Endpoint Gateway", func(t *testing.T) {
			res, err := DeleteEndpointGateway(vpcService, *createdEgwID)
			ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})
	})
	printTestSummary()
}
func TestVPCRoutingTables(t *testing.T) {
	vpcService := createVpcService(t)
	if *defaultVpcID == "" {
		res, _, err := ListInstances(vpcService)
		if err != nil {
			fmt.Println("Error: ", err)
			t.Error(err)
		}
		defaultVpcID = res.Instances[0].VPC.ID
		defaultSubnetID = res.Instances[0].PrimaryNetworkInterface.Subnet.ID
	} else {
		res, _, err := ListSubnets(vpcService)
		if err != nil {
			fmt.Println("Error: ", err)
			t.Error(err)
		}
		defaultVpcID = res.Subnets[0].VPC.ID
		defaultSubnetID = res.Subnets[0].ID
	}
	t.Run("Routing Tables", func(t *testing.T) {
		t.Run("Get Subnet Routing Table", func(t *testing.T) {
			res, _, err := GetSubnetRoutingTable(vpcService, *defaultSubnetID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Routing Table", func(t *testing.T) {
			res, _, err := GetVPCDefaultRoutingTable(vpcService, *defaultVpcID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Create Routing Table", func(t *testing.T) {
			name := "gsdk-rt-" + timestamp
			res, _, err := CreateVPCRoutingTable(vpcService, *defaultVpcID, name, *defaultZoneName)
			ValidateResponse(t, res, err, POST, detailed, increment)
			createdRtID = res.ID
		})

		t.Run("Create Routing Table 2", func(t *testing.T) {
			name := "gsdk-rt2-" + timestamp
			res, _, err := CreateVPCRoutingTable(vpcService, *defaultVpcID, name, *defaultZoneName)
			ValidateResponse(t, res, err, POST, detailed, increment)
			createdRt2ID = res.ID
		})

		t.Run("Replace Subnet Routing Table", func(t *testing.T) {
			res, _, err := ReplaceSubnetRoutingTable(vpcService, *defaultSubnetID, *createdRtID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})
		t.Run("List Routing Tables", func(t *testing.T) {
			res, _, err := ListVPCRoutingTables(vpcService, *defaultVpcID)
			ValidateListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Routing Table", func(t *testing.T) {
			res, _, err := GetVPCRoutingTable(vpcService, *defaultVpcID, *createdRtID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Update Routing Table", func(t *testing.T) {
			name := "gsdk-rt2-" + timestamp
			res, _, err := UpdateVPCRoutingTable(vpcService, *defaultVpcID, *createdRtID, name)
			ValidateResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("List Routing Table Routes", func(t *testing.T) {
			res, _, err := ListVPCRoutingTableRoutes(vpcService, *defaultVpcID, *createdRtID)
			ValidateListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Create Routing Table Route", func(t *testing.T) {
			res, _, err := CreateVPCRoutingTableRoute(vpcService, *defaultVpcID, *createdRtID, *defaultZoneName)
			ValidateResponse(t, res, err, GET, detailed, increment)
			createdRouteID = res.ID
		})

		t.Run("Get Routing Table Route", func(t *testing.T) {
			res, _, err := GetVPCRoutingTableRoute(vpcService, *defaultVpcID, *createdRtID, *createdRouteID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Update Routing Table Route", func(t *testing.T) {
			name := "gsdk-route-" + timestamp
			res, _, err := UpdateVPCRoutingTableRoute(vpcService, *defaultVpcID, *createdRtID, *createdRouteID, name)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Remove Routing Table Route", func(t *testing.T) {
			res, err := DeleteVPCRoutingTableRoute(vpcService, *defaultVpcID, *createdRtID, *createdRouteID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Delete Routing Table", func(t *testing.T) {
			res, err := DeleteVPCRoutingTable(vpcService, *defaultVpcID, *createdRt2ID)
			ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

	})
	printTestSummary()
}

func TestVPCDedicatedHosts(t *testing.T) {
	vpcService := createVpcService(t)
	t.Run("Dedicated Hosts", func(t *testing.T) {
		t.Run("List DH Groups", func(t *testing.T) {
			res, _, err := ListDedicatedHostGroups(vpcService)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Create DH Group", func(t *testing.T) {
			name := "gsdk-dhg-" + timestamp
			res, _, err := CreateDedicatedHostGroup(vpcService, name, *defaultZoneName)
			ValidateResponse(t, res, err, POST, detailed, increment)
			createdDhgID = res.ID
		})

		t.Run("Get DH Group", func(t *testing.T) {
			res, _, err := GetDedicatedHostGroup(vpcService, createdDhgID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Update DH Group", func(t *testing.T) {
			name := "gsdk-dhg2-" + timestamp
			res, _, err := UpdateDedicatedHostGroup(vpcService, createdDhgID, &name)
			ValidateResponse(t, res, err, PATCH, detailed, increment)
		})

		var dhProfileName *string
		t.Run("List DH Profiles", func(t *testing.T) {
			res, _, err := ListDedicatedHostProfiles(vpcService)
			ValidateResponse(t, res, err, GET, detailed, increment)
			dhProfileName = res.Profiles[0].Name
		})

		t.Run("Get DH Profile", func(t *testing.T) {
			res, _, err := GetDedicatedHostProfile(vpcService, dhProfileName)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("List DH", func(t *testing.T) {
			res, _, err := ListDedicatedHosts(vpcService)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Create DH", func(t *testing.T) {
			name := "gsdk-dh-" + timestamp
			res, _, err := CreateDedicatedHost(vpcService, &name, dhProfileName, createdDhgID)
			ValidateResponse(t, res, err, POST, detailed, increment)
			createdDhID = res.ID
		})

		t.Run("Get DH", func(t *testing.T) {
			res, _, err := GetDedicatedHost(vpcService, *createdDhID)
			ValidateResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Update DH", func(t *testing.T) {
			name := "gsdk-dh2-" + timestamp
			res, _, err := UpdateDedicatedHost(vpcService, &name, createdDhID)
			ValidateResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Delete DH Group", func(t *testing.T) {
			res, err := DeleteDedicatedHostGroup(vpcService, createdDhgID)
			ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

		t.Run("Delete DH", func(t *testing.T) {
			res, err := DeleteDedicatedHost(vpcService, createdDhID)
			ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

	})
}
func TestVPCTeardown(t *testing.T) {
	vpcService := createVpcService(t)
	shouldSkipTest(t)

	t.Run("Delete Resources", func(t *testing.T) {
		t.Run("Delete Instance Group Manager Policy", func(t *testing.T) {
			res, err := DeleteInstanceGroupManagerPolicy(vpcService, *createdInstanceGroupID, *createdIgManagerID, *createdIgPolicyID)
			ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

		t.Run("Delete Instance Group Manager", func(t *testing.T) {
			res, err := DeleteInstanceGroupManager(vpcService, *createdInstanceGroupID, *createdIgManagerID)
			ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

		t.Run("Delete Instance Group", func(t *testing.T) {
			res, err := DeleteInstanceGroup(vpcService, *createdInstanceGroupID)
			ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

		t.Run("Stop Instance ", func(t *testing.T) {
			statusChanged := PollInstance(vpcService, *createdInstanceID, Running, 4)
			fmt.Println("Stopping Instance")
			if statusChanged {
				res, _, err := CreateInstanceAction(vpcService, *createdInstanceID, "stop")
				ValidateResponse(t, res, err, DELETE, detailed, increment)
			}
		})

		t.Run("Delete Volume", func(t *testing.T) {
			res, err := DeleteVolume(vpcService, *createdVolumeID)
			ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

		t.Run("Delete Image", func(t *testing.T) {
			t.Skip("Skip Delete Image")
			res, err := DeleteImage(vpcService, *createdImageID)
			ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

		t.Run("Delete Instance Template", func(t *testing.T) {
			res, err := DeleteInstanceTemplate(vpcService, *createdTemplateID)
			ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

		t.Run("Delete Instance", func(t *testing.T) {
			statusChanged := PollInstance(vpcService, *createdInstanceID, Stopped, 4)
			fmt.Println("Deleting Instance")
			if statusChanged {
				res, err := DeleteInstance(vpcService, *createdInstanceID)
				ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
			}
		})

		t.Run("Delete Floating IP", func(t *testing.T) {
			res, err := ReleaseFloatingIP(vpcService, *createdFipID)
			ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

		t.Run("Delete Subnet", func(t *testing.T) {
			res, err := DeleteSubnet(vpcService, *createdSubnetID)
			ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

		t.Run("Delete VPC", func(t *testing.T) {
			res, err := DeleteVPC(vpcService, *createdVpcID)
			ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

		t.Run("Delete SSH Key", func(t *testing.T) {
			res, err := DeleteSSHKey(vpcService, *createdSSHKey)
			ValidateDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})
	})
	printTestSummary()
}

// ValidateListResponse - validate response for test list APIs
// x interface{} - response from response
// err error - err from response
// operation string - HTTP operation - GET
// detailed *bool - bool to view the detailed response from API
func ValidateListResponse(t *testing.T, x interface{}, err error, operation string, detailed *bool, increment func()) {
	if err != nil && x == nil {
		fmt.Println("Error: ", err)
		t.Errorf("Error: %s %s", operation, reflect.TypeOf(x).String())
		t.Error(err)
		return
	}
	if err != nil && x != nil {
		t.Error(err)
		return
	}
	t.Log("Success: Recieved ", operation, reflect.TypeOf(x).String())
	if *detailed {
		Print(x)
	}
	increment()
}

// ValidateResponse - validate response for test get and update
// x interface{} - response from response
// err error - err from response
// operation string - HTTP operation - GET/POST/PATCH/PUT
// detailed *bool - bool to view the detailed response from API
// resourceID string - resource ID
func ValidateResponse(t *testing.T, x interface{}, err error, operation string, detailed *bool, increment func()) {
	if err != nil {
		fmt.Println("Error: ", err)
		t.Error("Error: ", operation, reflect.TypeOf(x).String())
		t.Error(err)
		return
	}
	if err != nil && x != nil {
		t.Error(err)
		return
	}
	t.Log("Success: Recieved ", operation, reflect.TypeOf(x).String())
	if *detailed {
		Print(x)
	}
	increment()
}

// ValidateDeleteResponse - validate response  for test delete
// x interface{} - response from response
// err error - err from response
// operation string - HTTP operation - DELETE
// detailed *bool - bool to view the detailed response from API
// resourceID string - resource ID
// statusCode int - status code from response
func ValidateDeleteResponse(t *testing.T, x interface{}, err error, operation string, statusCode int, detailed *bool, increment func()) {
	if err != nil && x == nil {
		fmt.Println("Error: ", err)
		t.Errorf("Error: %s %s", operation, reflect.TypeOf(x).String())
		t.Error(err)
		return
	}
	if err != nil && x != nil {
		fmt.Println("Error: ", err)
		t.Error(err)
		return
	}
	t.Log("Success: Recieved ", operation, reflect.TypeOf(x).String())
	t.Log("Status Code:", statusCode)
	if *detailed {
		Print(x)
	}
	increment()
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

func increment() {
	if *testCount {
		counter.increment()
	}
}

func printTestSummary() {
	fmt.Printf("Total test run: %d\n", counter.currentValue())
}

func getName(rtype string) string {
	return "gsdk-" + rtype + "-" + timestamp
}
