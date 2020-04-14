package vpcgen2integration

import (
	"encoding/json"
	"flag"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/IBM/go-sdk-core/v3/core"
	"github.ibm.com/ibmcloud/vpc-go-sdk/vpcv1"
)

var detailed = flag.Bool("detailed", false, "boolean")
var skipForMockTesting = flag.Bool("skipForMockTesting", false, "boolean")
var testCount = flag.Bool("testCount", false, "boolean")

var defaultFipID *string
var defaultImageID *string
var defaultOSName *string
var defaultInstanceID *string
var defaultInstanceProfile *string
var defaultRegionName *string
var defaultZoneName *string
var defaultResourceGroupID *string
var defaultVolumeProfile *string
var bootVolID *string
var bootVolAttachmentID *string
var defaultVpcID *string
var createdVpcID *string
var createdSubnetID *string
var defaultVnicID *string
var defaultSubnetID *string
var defaultACLID *string
var defaultLBID *string
var defaultLBListenerPolicyID *string
var defaultLBRule *string
var defaultLBPoolID *string
var defaultLBPoolMemberID *string
var defaultLBListenerID *string
var createdSSHKey *string
var createdVolumeID *string
var createdVPCRouteID *string
var createdVpcAddressPrefixID *string
var createdFipID *string
var createdInstanceID *string
var createdVnicID *string
var createdVolAttachmentID *string
var createdACLID *string
var createdACLRuleID *string
var createdPGWID *string
var createdSgID *string
var createdSgVnicID *string
var createdSgRuleID *string
var createdSecondVnicID *string

var Running = "running"
var Stopped = "stopped"
var Attached = "attached"
var tunix = time.Now().Unix()
var timestamp = strconv.FormatInt(tunix, 10)

var counter = Counter{0}

func increment() {
	if *testCount {
		counter.increment()
	}
}

func printTestSummary() {
	fmt.Printf("Total test run: %d\n", counter.currentValue())
}
func TestConnectVPC(t *testing.T) {
	if !*skipForMockTesting {
		var gen2 = InstantiateVPCGen2Service()
		if gen2 == nil {
			fmt.Println("Error creating VPC Gen2 service.")
			t.Error("Error creating vpc gen 2 service with error message:")
			return
		}
		t.Log("Success: VPC Gen2 service creation complete.")
	}
}

func createVpcGen2Service(t *testing.T) *vpcv1.VpcV1 {
	if *skipForMockTesting {
		testService, _ := vpcv1.NewVpcV1(&vpcv1.VpcV1Options{
			URL:           URL,
			Authenticator: &core.NoAuthAuthenticator{},
		})
		return testService
	}
	var gen2 = InstantiateVPCGen2Service()
	if gen2 == nil {
		fmt.Println("Error creating VPC Gen2 service.")
		t.Error("Error creating vpc gen 2 service with error message:")
		return nil
	}
	t.Log("Success: VPC Gen2 service creation complete.")
	return gen2
}

func TestVPCResources(t *testing.T) {
	vpcService := createVpcGen2Service(t)

	t.Run("Geography", func(t *testing.T) {

		t.Run("All regions", func(t *testing.T) {
			res, _, err := ListRegions(vpcService)
			TestListResponse(t, res, err, GET, detailed, increment)
			defaultRegionName = res.Regions[0].Name
		})

		t.Run("Get region", func(t *testing.T) {
			res, _, err := GetRegion(vpcService, *defaultRegionName)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Zones within Region", func(t *testing.T) {
			if *skipForMockTesting {
				zone := "us-east" + "-1"
				defaultZoneName = &zone
				t.Skip("skipping test in travis.")
			}
			t.Run("Zones within Region", func(t *testing.T) {
				res, _, err := ListZones(vpcService, *defaultRegionName)
				TestListResponse(t, res, err, GET, detailed, increment)
				defaultZoneName = res.Zones[0].Name
			})

			t.Run("Get Zone", func(t *testing.T) {
				res, _, err := GetZone(vpcService, *defaultRegionName, *defaultZoneName)
				TestResponse(t, res, err, GET, detailed, increment)
			})
		})
	})

	t.Run("Create", func(t *testing.T) {

		t.Run("Initial Setup", func(t *testing.T) {

			// getting default resource group assuming there is atleast one VPC in the account.
			vpcs, _, err := GetVPCsList(vpcService)
			if err != nil && vpcs == nil {
				fmt.Println("Error: ", err)
				t.Error("Error fetching for Resource Group with error message:", err)
				return
			}
			defaultResourceGroupID = vpcs.Vpcs[0].ResourceGroup.ID

			t.Run("List Instance Profiles", func(t *testing.T) {
				res, _, err := ListInstanceProfiles(vpcService)
				TestListResponse(t, res, err, GET, detailed, increment)
				defaultInstanceProfile = res.Profiles[0].Name
			})

			t.Run("List Volume Profiles", func(t *testing.T) {
				res, _, err := ListVolumeProfiles(vpcService)
				TestListResponse(t, res, err, GET, detailed, increment)
				defaultVolumeProfile = res.Profiles[0].Name
			})

			t.Run("Get Volume Profile", func(t *testing.T) {
				res, _, err := GetVolumeProfile(vpcService, *defaultVolumeProfile)
				TestResponse(t, res, err, GET, detailed, increment)
			})

			t.Run("Get Instance Profile", func(t *testing.T) {
				res, _, err := GetInstanceProfile(vpcService, *defaultInstanceProfile)
				TestResponse(t, res, err, GET, detailed, increment)
			})

			t.Run("List Images", func(t *testing.T) {
				res, _, err := ListImages(vpcService, "public")
				TestListResponse(t, res, err, GET, detailed, increment)
				defaultImageID = res.Images[0].ID
			})
			t.Run("List Operating Systems", func(t *testing.T) {
				res, _, err := ListOperatingSystems(vpcService)
				TestListResponse(t, res, err, GET, detailed, increment)
				defaultOSName = res.OperatingSystems[0].Name
			})

			t.Run("Get Operating System", func(t *testing.T) {
				res, _, err := GetOperatingSystem(vpcService, *defaultOSName)
				TestResponse(t, res, err, GET, detailed, increment)
			})

		})

		t.Run("Create VPC", func(t *testing.T) {
			name := "gosdk-vpc-" + timestamp
			res, _, err := CreateVPC(vpcService, name, *defaultResourceGroupID)
			TestResponse(t, res, err, POST, detailed, increment)
			createdVpcID = res.ID
		})

		t.Run("Create Subnet", func(t *testing.T) {
			name := "gosdk-subnet-" + timestamp
			res, _, err := CreateSubnet(vpcService, *createdVpcID, name, *defaultZoneName)
			TestResponse(t, res, err, POST, detailed, increment)
			createdSubnetID = res.ID
		})

		t.Run("Create SSH key", func(t *testing.T) {
			name := "gosdk-key-" + timestamp
			res, _, err := CreateSSHKey(vpcService, name, "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCcPJwUpNQr0MplO6UM5mfV4vlvY0RpD6gcXqodzZIjsoG31+hQxoJVU9yQcSjahktHFs7Fk2Mo79jUT3wVC8Pg6A3//IDFkLjVrg/mQVpIf6+GxIYEtVg6Tk4pP3YNoksrugGlpJ4LCR3HMe3fBQTQqTzObbb0cSF6xhW5UBq8vhqIkhYKd3KLGJnnrwsIGcwb5BRk68ZFYhreAomvx4jWjaBFlH98HhE4wUEVvJLRy/qR/0w3XVjTSgOlhXywaAOEkmwye7kgSglegCpHWwYNly+NxLONjqbX9rHbFHUVRShnFKh2+M6XKE3HowT/3Y1lDd2PiVQpJY0oQmebiRxB astha.jain@ibm.com")
			TestResponse(t, res, err, POST, detailed, increment)
			createdSSHKey = res.ID
		})

		t.Run("Create VPC Address Prefix", func(t *testing.T) {
			name := "gosdk-addprefix-" + timestamp
			res, _, err := CreateVpcAddressPrefix(vpcService, *createdVpcID, *defaultZoneName, "211.211.201.0/24", name)
			TestResponse(t, res, err, POST, detailed, increment)
			createdVpcAddressPrefixID = res.ID
		})

		t.Run("Create Floating IP", func(t *testing.T) {
			name := "gosdk-fip-" + timestamp
			res, _, err := CreateFloatingIP(vpcService, *defaultZoneName, name)
			TestResponse(t, res, err, POST, detailed, increment)
			createdFipID = res.ID
		})

		t.Run("Create Volume", func(t *testing.T) {
			name := "gosdk-vol-" + timestamp
			res, _, err := CreateVolume(vpcService, name, *defaultVolumeProfile, *defaultZoneName, 10)
			TestResponse(t, res, err, POST, detailed, increment)
			createdVolumeID = res.ID
		})

		t.Run("Create Instance", func(t *testing.T) {
			name := "gosdk-vsi-" + timestamp
			statusChanged := PollSubnet(vpcService, *createdSubnetID, "available", 4)
			if statusChanged {
				res, _, err := CreateInstance(vpcService, name, "bc1-4x16", *defaultImageID, *defaultZoneName, *createdSubnetID, *createdSSHKey, *createdVpcID)
				TestResponse(t, res, err, POST, detailed, increment)
				createdInstanceID = res.ID
				createdVnicID = res.PrimaryNetworkInterface.ID
			}
		})
	})

	t.Run("VPC Resources", func(t *testing.T) {

		t.Run("List VPCs", func(t *testing.T) {
			res, _, err := GetVPCsList(vpcService)
			TestListResponse(t, res, err, GET, detailed, increment)
			defaultVpcID = res.Vpcs[0].ID
		})

		t.Run("List Subnets", func(t *testing.T) {
			res, _, err := ListSubnets(vpcService)
			TestListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("List Instances", func(t *testing.T) {
			res, _, err := ListInstances(vpcService)
			TestListResponse(t, res, err, GET, detailed, increment)
			defaultInstanceID = res.Instances[0].ID
			defaultVnicID = res.Instances[0].PrimaryNetworkInterface.ID
			defaultVpcID = res.Instances[0].Vpc.ID
			defaultSubnetID = res.Instances[0].PrimaryNetworkInterface.Subnet.ID
		})

		t.Run("List VPC Address Prefixes", func(t *testing.T) {
			res, _, err := ListVpcAddressPrefixes(vpcService, *createdVpcID)
			TestListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("List SSH Keys", func(t *testing.T) {
			res, _, err := ListKeys(vpcService)
			TestListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("List Floating IPs", func(t *testing.T) {
			res, _, err := GetFloatingIPsList(vpcService)
			TestListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("List Volumes", func(t *testing.T) {
			res, _, err := ListVolumes(vpcService)
			TestListResponse(t, res, err, GET, detailed, increment)
		})
	})

	t.Run("Get a VPC Resource", func(t *testing.T) {

		t.Run("Get VPC", func(t *testing.T) {
			res, _, err := GetVPC(vpcService, *createdVpcID)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get VPC Default Security Group", func(t *testing.T) {
			res, _, err := GetVPCDefaultSecurityGroup(vpcService, *createdVpcID)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get VPC Default Network ACL", func(t *testing.T) {
			res, _, err := GetVPCDefaultACL(vpcService, *createdVpcID)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get VPC Address Prefix", func(t *testing.T) {
			res, _, err := GetVpcAddressPrefix(vpcService, *createdVpcID, *createdVpcAddressPrefixID)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get SSH Key", func(t *testing.T) {
			res, _, err := GetSSHKey(vpcService, *createdSSHKey)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Subnet", func(t *testing.T) {
			res, _, err := GetSubnet(vpcService, *createdSubnetID)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Instance", func(t *testing.T) {
			statusChanged := PollInstance(vpcService, *createdInstanceID, Running, 7)
			if statusChanged {
				res, _, err := GetInstance(vpcService, *createdInstanceID)
				TestResponse(t, res, err, GET, detailed, increment)
			}
		})

		t.Run("Get Floating IP", func(t *testing.T) {
			res, _, err := GetFloatingIP(vpcService, *createdFipID)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Volume", func(t *testing.T) {
			res, _, err := GetVolume(vpcService, *createdVolumeID)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Image", func(t *testing.T) {
			res, _, err := GetImage(vpcService, *defaultImageID)
			TestResponse(t, res, err, GET, detailed, increment)
		})

	})

	t.Run("VPC Routes", func(t *testing.T) {
		t.Run("Create VPC Route", func(t *testing.T) {
			name := "gosdk-route-" + timestamp
			res, _, err := CreateVpcRoute(vpcService, *createdVpcID, *defaultZoneName, "5.0.5.0/28", "100.0.1.1", name)
			createdVPCRouteID = res.ID
			TestResponse(t, res, err, POST, detailed, increment)
		})

		t.Run("List VPC Routes", func(t *testing.T) {
			res, _, err := ListVpcRoutes(vpcService, *createdVpcID)
			TestListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get VPC Routes", func(t *testing.T) {
			res, _, err := ListVpcRoutes(vpcService, *createdVpcID)
			TestListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Update VPC Route", func(t *testing.T) {
			name := "gosdk-route-2-" + timestamp
			res, _, err := UpdateVpcRoute(vpcService, *createdVpcID, *createdVPCRouteID, name)
			TestResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Delete VPC Route", func(t *testing.T) {
			res, err := DeleteVpcRoute(vpcService, *createdVpcID, *createdVPCRouteID)
			TestDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})
	})

	t.Run("Instances Network Attachments", func(t *testing.T) {
		t.Run("Get Initialization", func(t *testing.T) {
			res, _, err := GetInstanceInitialization(vpcService, *createdInstanceID)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Network Interfaces", func(t *testing.T) {
			res, _, err := ListNetworkInterfaces(vpcService, *createdInstanceID)
			TestResponse(t, res, err, GET, detailed, increment)
			createdVnicID = res.NetworkInterfaces[0].ID
		})

		t.Run("Create Network Interfaces", func(t *testing.T) {
			res, _, err := CreateNetworkInterface(vpcService, *createdInstanceID, *createdSubnetID)
			TestResponse(t, res, err, GET, detailed, increment)
			createdSecondVnicID = res.ID
		})

		t.Run("Attach FIP to Vnic", func(t *testing.T) {
			res, _, err := CreateNetworkInterfaceFloatingIpBinding(vpcService, *createdInstanceID, *createdVnicID, *createdFipID)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Network Interface", func(t *testing.T) {
			res, _, err := GetNetworkInterface(vpcService, *createdInstanceID, *createdVnicID)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Update Network Interface", func(t *testing.T) {
			res, _, err := UpdateNetworkInterface(vpcService, *createdInstanceID, *createdVnicID, "vnic1")
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Vnic FLoating IPs", func(t *testing.T) {
			res, _, err := ListNetworkInterfaceFloatingIps(vpcService, *createdInstanceID, *createdVnicID)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Vnic FLoating IP", func(t *testing.T) {
			res, _, err := GetNetworkInterfaceFloatingIp(vpcService, *createdInstanceID, *createdVnicID, *createdFipID)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Delete Network Interfaces", func(t *testing.T) {
			res, err := DeleteNetworkInterface(vpcService, *createdInstanceID, *createdSecondVnicID)
			TestDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

	})

	t.Run("Instances Volume Attachments", func(t *testing.T) {

		t.Run("Create Volume attachment", func(t *testing.T) {
			name := "gosdk-attachment-" + timestamp
			res, _, err := CreateVolumeAttachment(vpcService, *createdInstanceID, *createdVolumeID, name)
			TestResponse(t, res, err, POST, detailed, increment)
			createdVolAttachmentID = res.ID
		})

		t.Run("Get Volume attachments", func(t *testing.T) {
			res, _, err := ListVolumeAttachments(vpcService, *createdInstanceID)
			TestResponse(t, res, err, GET, detailed, increment)
			bootVolAttachmentID = res.VolumeAttachments[0].ID
		})

		t.Run("Get Volume attachment", func(t *testing.T) {
			res, _, err := GetVolumeAttachment(vpcService, *createdInstanceID, *createdVolAttachmentID)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Update Volume attachments", func(t *testing.T) {
			name := "gosdk-boot-2-" + timestamp
			res, _, err := UpdateVolumeAttachment(vpcService, *createdInstanceID, *bootVolAttachmentID, name)
			TestResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Delete Volume attachments", func(t *testing.T) {
			statusChanged := PollVolAttachment(vpcService, *createdInstanceID, *createdVolAttachmentID, Attached, 4)
			if statusChanged {
				res, err := DeleteVolumeAttachment(vpcService, *createdInstanceID, *createdVolAttachmentID)
				TestResponse(t, res, err, DELETE, detailed, increment)
			}
		})
	})

	t.Run("Subnet Bindings", func(t *testing.T) {

		t.Run("Set Subnet NetworkAcl Binding", func(t *testing.T) {
			acls, _, err := ListNetworkAcls(vpcService)

			res, _, err := SetSubnetNetworkAclBinding(vpcService, *createdSubnetID, *acls.NetworkAcls[0].ID)
			TestResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Get Subnet NetworkAcl", func(t *testing.T) {
			res, _, err := GetSubnetNetworkAcl(vpcService, *createdSubnetID)
			TestResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Set Subnet Public Gateway Binding", func(t *testing.T) {
			name := "public-gateway"
			pgw, _, err := CreatePublicGateway(vpcService, name, *defaultVpcID, *defaultZoneName)
			res, _, err := SetSubnetPublicGatewayBinding(vpcService, *createdSubnetID, *pgw.ID)
			TestResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Get Subnet Public Gateway", func(t *testing.T) {
			res, _, err := GetSubnetPublicGateway(vpcService, *createdSubnetID)
			TestResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Delete Subnet Public Gateway Binding", func(t *testing.T) {
			res, err := DeleteSubnetPublicGatewayBinding(vpcService, *createdSubnetID)
			TestDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

	})

	t.Run("Update VPC Resources", func(t *testing.T) {

		t.Run("Update Floating IP", func(t *testing.T) {
			name := "gosdk-fip-2-" + timestamp
			res, _, err := UpdateFloatingIP(vpcService, *createdFipID, name)
			TestResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Update Image", func(t *testing.T) {
			name := "gosdk-image-2-" + timestamp
			res, _, err := UpdateImage(vpcService, *defaultImageID, name)
			TestResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Update SSH key", func(t *testing.T) {
			name := "gosdk-key-2-" + timestamp
			res, _, err := UpdateSSHKey(vpcService, *createdSSHKey, name)
			TestResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Update VPC Address Prefixes", func(t *testing.T) {
			name := "gosdk-prefix-2-" + timestamp
			res, _, err := UpdateVpcAddressPrefix(vpcService, *createdVpcID, *createdVpcAddressPrefixID, name)
			TestResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Update Instance", func(t *testing.T) {
			name := "gosdk-vsi-2-" + timestamp
			statusChanged := PollInstance(vpcService, *createdInstanceID, Running, 4)
			if statusChanged {
				res, _, err := UpdateInstance(vpcService, *createdInstanceID, name)
				TestResponse(t, res, err, PATCH, detailed, increment)
			}
		})

		t.Run("Update Subnet", func(t *testing.T) {
			name := "gosdk-subnet-2-" + timestamp
			res, _, err := UpdateSubnet(vpcService, *createdSubnetID, name)
			TestResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Update VPC", func(t *testing.T) {
			name := "gosdk-vpc-2-" + timestamp
			res, _, err := UpdateVPC(vpcService, *createdVpcID, name)
			TestResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Update Volume", func(t *testing.T) {
			name := "gosdk-vol-2-" + timestamp
			res, _, err := UpdateVolume(vpcService, *createdVolumeID, name)
			TestResponse(t, res, err, PATCH, detailed, increment)
		})
	})

	t.Run("Delete Resources", func(t *testing.T) {

		t.Run("Stop Instance ", func(t *testing.T) {
			statusChanged := PollInstance(vpcService, *createdInstanceID, Running, 4)
			fmt.Println("Stopping Instance")
			if statusChanged {
				res, _, err := CreateInstanceAction(vpcService, *createdInstanceID, "stop")
				TestResponse(t, res, err, DELETE, detailed, increment)
			}
		})

		t.Run("Delete Volume", func(t *testing.T) {
			res, err := DeleteVolume(vpcService, *createdVolumeID)
			TestDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

		t.Run("Delete Instance", func(t *testing.T) {
			statusChanged := PollInstance(vpcService, *createdInstanceID, Stopped, 4)
			fmt.Println("Deleting Instance")
			if statusChanged {
				res, err := DeleteInstance(vpcService, *createdInstanceID)
				TestDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
			}
		})

		t.Run("Delete Floating IP", func(t *testing.T) {
			res, err := ReleaseFloatingIP(vpcService, *createdFipID)
			TestDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

		t.Run("Delete VPC Address Prefixes", func(t *testing.T) {
			res, err := DeleteVpcAddressPrefix(vpcService, *createdVpcID, *createdVpcAddressPrefixID)
			TestDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

		t.Run("Delete Subnet", func(t *testing.T) {
			res, err := DeleteSubnet(vpcService, *createdSubnetID)
			TestDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

		t.Run("Delete VPC", func(t *testing.T) {
			res, err := DeleteVPC(vpcService, *createdVpcID)
			TestDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

		t.Run("Delete Image", func(t *testing.T) {
			res, err := DeleteImage(vpcService, *defaultImageID)
			TestDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

		t.Run("Delete SSH Key", func(t *testing.T) {
			res, err := DeleteSSHKey(vpcService, *createdSSHKey)
			TestDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})
	})
	printTestSummary()
}

func TestVPCAccessControlLists(t *testing.T) {
	vpcService := createVpcGen2Service(t)

	t.Run("ACL Resources", func(t *testing.T) {

		t.Run("List  ACLs", func(t *testing.T) {
			res, _, err := ListNetworkAcls(vpcService)
			TestListResponse(t, res, err, GET, detailed, increment)
			nacl := res.NetworkAcls[0]
			defaultACLID = nacl.ID
			defaultVpcID = nacl.Vpc.ID
		})

		t.Run("Create ACL", func(t *testing.T) {
			name := "gosdk-acl-" + timestamp
			res, _, err := CreateNetworkAcl(vpcService, name, *defaultACLID, *defaultVpcID)
			TestResponse(t, res, err, POST, detailed, increment)
			createdACLID = res.ID
		})

		t.Run("Create ACL Rule", func(t *testing.T) {
			name := "gosdk-aclrule-" + timestamp
			res, _, err := CreateNetworkAclRule(vpcService, name, *createdACLID)
			res2B, _ := json.Marshal(res)
			rule := &vpcv1.NetworkACLRule{}
			_ = json.Unmarshal([]byte(string(res2B)), &rule)
			TestResponse(t, rule, err, POST, detailed, increment)
			createdACLRuleID = rule.ID
		})

		t.Run("List ACL Rules", func(t *testing.T) {
			res, _, err := ListNetworkAclRules(vpcService, *createdACLID)
			TestListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get ACL", func(t *testing.T) {
			res, _, err := GetNetworkAcl(vpcService, *createdACLID)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get ACL Rules", func(t *testing.T) {
			res, _, err := GetNetworkAclRule(vpcService, *createdACLID, *createdACLRuleID)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Update ACL", func(t *testing.T) {
			name := "gosdk-acl-2-" + timestamp
			res, _, err := UpdateNetworkAcl(vpcService, *createdACLID, name)
			TestResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Update ACL Rule", func(t *testing.T) {
			name := "gosdk-acl-2-" + timestamp
			res, _, err := UpdateNetworkAclRule(vpcService, *createdACLID, *createdACLRuleID, name)
			TestResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Delete ACL Rule", func(t *testing.T) {
			res, err := DeleteNetworkAclRule(vpcService, *createdACLID, *createdACLRuleID)
			TestDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

		t.Run("Delete ACL", func(t *testing.T) {
			res, err := DeleteNetworkAcl(vpcService, *createdACLID)
			TestDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})
	})
	printTestSummary()
}
func TestVPCSecurityGroups(t *testing.T) {
	vpcService := createVpcGen2Service(t)
	res, _, err := ListInstances(vpcService)
	if err != nil {
		fmt.Println("Error: ", err)
		t.Error(err)
	}
	var defaultVpcID = res.Instances[0].Vpc.ID
	var defaultVnicID = res.Instances[0].PrimaryNetworkInterface.ID
	t.Run("SG Resources", func(t *testing.T) {

		var sgID *string
		t.Run("List Security Groups", func(t *testing.T) {
			res, _, err := ListSecurityGroups(vpcService)
			TestListResponse(t, res, err, GET, detailed, increment)
			sgID = res.SecurityGroups[0].ID
		})

		t.Run("List Security Group Network Interfaces", func(t *testing.T) {
			res, _, err := ListSecurityGroupNetworkInterfaces(vpcService, *sgID)
			TestListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("List Security Group Rules", func(t *testing.T) {
			res, _, err := ListSecurityGroupRules(vpcService, *sgID)
			TestListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Create Security Group", func(t *testing.T) {
			name := "gosdk-sg-" + timestamp
			res, _, err := CreateSecurityGroup(vpcService, name, *defaultVpcID)
			TestResponse(t, res, err, POST, detailed, increment)
			createdSgID = res.ID
		})

		t.Run("Create Security Group Network Interface", func(t *testing.T) {
			res, _, err := CreateSecurityGroupNetworkInterfaceBinding(vpcService, *createdSgID, *defaultVnicID)
			TestResponse(t, res, err, POST, detailed, increment)
			createdSgVnicID = res.ID
		})

		t.Run("Create Security Group Rule", func(t *testing.T) {
			res, _, err := CreateSecurityGroupRule(vpcService, *createdSgID)
			res2B, _ := json.Marshal(res)
			rule := &vpcv1.SecurityGroupRule{}
			_ = json.Unmarshal([]byte(string(res2B)), &rule)
			TestResponse(t, rule, err, POST, detailed, increment)
			createdSgRuleID = rule.ID
		})

		t.Run("Get Security Group", func(t *testing.T) {
			res, _, err := GetSecurityGroup(vpcService, *sgID)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Security Group Network Interface", func(t *testing.T) {
			res, _, err := GetSecurityGroupNetworkInterface(vpcService, *createdSgID, *defaultVnicID)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Security Group Rules", func(t *testing.T) {
			res, _, err := GetSecurityGroupRule(vpcService, *createdSgID, *createdSgRuleID)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Update Security Groups", func(t *testing.T) {
			name := "gosdk-sg-2-" + timestamp
			res, _, err := UpdateSecurityGroup(vpcService, *createdSgID, name)
			TestResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Update Security Group Rule", func(t *testing.T) {
			res, _, err := UpdateSecurityGroupRule(vpcService, *createdSgID, *createdSgRuleID)
			TestResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Delete Security Group Network Interface", func(t *testing.T) {
			res, err := DeleteSecurityGroupNetworkInterfaceBinding(vpcService, *createdSgID, *defaultVnicID)
			TestDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

		t.Run("Delete Security Group Rule", func(t *testing.T) {
			res, err := DeleteSecurityGroupRule(vpcService, *createdSgID, *createdSgRuleID)
			TestDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

		t.Run("Delete Security Group", func(t *testing.T) {
			res, err := DeleteSecurityGroup(vpcService, *createdSgID)
			TestDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})
	})
	printTestSummary()
}
func TestVPCPublicGateways(t *testing.T) {
	vpcService := createVpcGen2Service(t)
	res, _, err := ListInstances(vpcService)
	if err != nil {
		fmt.Println("Error: ", err)
		t.Error(err)
	}
	var defaultVpcID = res.Instances[0].Vpc.ID
	var defaultZoneName = res.Instances[0].Zone.Name
	t.Run("PGW Resources", func(t *testing.T) {

		t.Run("Create Public Gateway", func(t *testing.T) {
			name := "gosdk-pgw-" + timestamp
			res, _, err := CreatePublicGateway(vpcService, name, *defaultVpcID, *defaultZoneName)
			TestResponse(t, res, err, POST, detailed, increment)
			createdPGWID = res.ID
		})

		t.Run("List  Public Gateways", func(t *testing.T) {
			res, _, err := ListPublicGateways(vpcService)
			TestListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Public Gateway", func(t *testing.T) {
			res, _, err := GetPublicGateway(vpcService, *createdPGWID)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Update Public Gateway", func(t *testing.T) {
			name := "gosdk-pgw-2-" + timestamp
			res, _, err := UpdatePublicGateway(vpcService, *createdPGWID, name)
			TestResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Delete Public Gateway", func(t *testing.T) {
			res, err := DeletePublicGateway(vpcService, *createdPGWID)
			TestDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})
	})
	printTestSummary()
}

func TestVPCLoadBalancers(t *testing.T) {
	vpcService := createVpcGen2Service(t)

	t.Run("LB Resources", func(t *testing.T) {
		var lbID *string
		var subnetID *string
		res, _, err := ListInstances(vpcService)
		if err != nil && res == nil {
			fmt.Println("Error retrieving subnet ID")
		}
		subnetID = res.Instances[0].PrimaryNetworkInterface.Subnet.ID

		t.Run("List Load Balancers", func(t *testing.T) {
			res, _, err := ListLoadBalancers(vpcService)
			TestListResponse(t, res, err, GET, detailed, increment)
			// find a lb with Provisioning status as active
			for _, i := range res.LoadBalancers {
				if *i.ProvisioningStatus == "active" {
					lbID = i.ID
					fmt.Println("Found an LB with active status")
					break
				}
			}
		})

		t.Run("Create Load Balancer", func(t *testing.T) {
			name := "gen1-gosdk-lb-" + strconv.FormatInt(tunix, 10)
			res, _, err := CreateLoadBalancer(vpcService, name, *subnetID)
			TestResponse(t, res, err, POST, detailed, increment)
			defaultLBID = res.ID
		})

		t.Run("List Load Balancer Listeners", func(t *testing.T) {
			res, _, err := ListLoadBalancerListeners(vpcService, *lbID)
			TestListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Create Load Balancer Listener", func(t *testing.T) {
			statusChanged := PollLB(vpcService, *defaultLBID, "active", 8)
			if statusChanged {
				res, _, err := CreateLoadBalancerListener(vpcService, *defaultLBID)
				TestResponse(t, res, err, POST, detailed, increment)
				defaultLBListenerID = res.ID
			}
		})

		t.Run("Get Load Balancer", func(t *testing.T) {
			res, _, err := GetLoadBalancer(vpcService, *defaultLBID)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Create Load Balancer Listener Policy", func(t *testing.T) {
			statusChanged := PollLB(vpcService, *defaultLBID, "active", 5)
			if statusChanged {
				res, _, err := CreateLoadBalancerListenerPolicy(vpcService, *defaultLBID, *defaultLBListenerID)
				TestResponse(t, res, err, POST, detailed, increment)
				defaultLBListenerPolicyID = res.ID
			}
		})

		t.Run("Create Load Balancer Listener Policy Rule", func(t *testing.T) {
			statusChanged := PollLB(vpcService, *defaultLBID, "active", 5)
			if statusChanged {
				res, _, err := CreateLoadBalancerListenerPolicyRule(vpcService, *defaultLBID, *defaultLBListenerID, *defaultLBListenerPolicyID)
				TestResponse(t, res, err, POST, detailed, increment)
				defaultLBRule = res.ID
			}
		})
		var poolID *string
		t.Run("Create Load Balancer Pool", func(t *testing.T) {
			statusChanged := PollLB(vpcService, *defaultLBID, "active", 8)
			if statusChanged {
				name := "gosdk-lbpool-" + timestamp
				res, _, err := CreateLoadBalancerPool(vpcService, *defaultLBID, name)
				TestResponse(t, res, err, POST, detailed, increment)
				defaultLBPoolID = res.ID
				name = "go-lb-pool-2-" + timestamp
				res, _, err = CreateLoadBalancerPool(vpcService, *defaultLBID, name)
				poolID = res.ID
			}
		})

		t.Run("List Load Balancer Listeners Policies", func(t *testing.T) {
			res, _, err := ListLoadBalancerListenerPolicies(vpcService, *defaultLBID, *defaultLBListenerID)
			TestListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("List Load Balancer Listeners Policy Rules", func(t *testing.T) {
			res, _, err := ListLoadBalancerListenerPolicyRules(vpcService, *defaultLBID, *defaultLBListenerID, *defaultLBListenerPolicyID)
			TestListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Load Balancer Statistics", func(t *testing.T) {
			res, _, err := GetLoadBalancerStatistics(vpcService, *defaultLBID)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Load Balancer Listener", func(t *testing.T) {
			res, _, err := GetLoadBalancerListener(vpcService, *defaultLBID, *defaultLBListenerID)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Load Balancer Listener Policy", func(t *testing.T) {
			res, _, err := GetLoadBalancerListenerPolicy(vpcService, *defaultLBID, *defaultLBListenerID, *defaultLBListenerPolicyID)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Load Balancer Listener Policy Rule", func(t *testing.T) {
			res, _, err := GetLoadBalancerListenerPolicyRule(vpcService, *defaultLBID, *defaultLBListenerID, *defaultLBListenerPolicyID, *defaultLBRule)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Update Load Balancer Listener Policy Rule", func(t *testing.T) {
			res, _, err := UpdateLoadBalancerListenerPolicyRule(vpcService, *defaultLBID, *defaultLBListenerID, *defaultLBListenerPolicyID, *defaultLBRule)
			TestResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Update Load Balancer Listener Policy", func(t *testing.T) {
			res, _, err := UpdateLoadBalancerListenerPolicy(vpcService, *defaultLBID, *defaultLBListenerID, *defaultLBListenerPolicyID, *poolID)
			TestResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Update Load Balancer Listener", func(t *testing.T) {
			res, _, err := UpdateLoadBalancerListener(vpcService, *defaultLBID, *defaultLBListenerID)
			TestResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Update Load Balancer", func(t *testing.T) {
			name := "gosdk-lb-2-" + timestamp
			res, _, err := UpdateLoadBalancer(vpcService, *defaultLBID, name)
			TestResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Delete Load Balancer listener Policy Rule", func(t *testing.T) {
			statusChanged := PollLB(vpcService, *defaultLBID, "active", 5)
			if statusChanged {
				res, err := DeleteLoadBalancerListenerPolicyRule(vpcService, *defaultLBID, *defaultLBListenerID, *defaultLBListenerPolicyID, *defaultLBRule)
				TestDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
			}
		})

		t.Run("Delete Load Balancer listener Policy", func(t *testing.T) {
			statusChanged := PollLB(vpcService, *defaultLBID, "active", 5)
			if statusChanged {
				res, err := DeleteLoadBalancerListenerPolicy(vpcService, *defaultLBID, *defaultLBListenerID, *defaultLBListenerPolicyID)
				TestDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
			}
		})

		t.Run("Delete Load Balancer listener", func(t *testing.T) {
			statusChanged := PollLB(vpcService, *defaultLBID, "active", 5)
			if statusChanged {
				res, err := DeleteLoadBalancerListener(vpcService, *defaultLBID, *defaultLBListenerID)
				TestDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
			}
		})

		t.Run("Create Load Balancer Pool Member", func(t *testing.T) {
			statusChanged := PollLB(vpcService, *defaultLBID, "active", 5)
			if statusChanged {
				res, _, err := CreateLoadBalancerPoolMember(vpcService, *defaultLBID, *defaultLBPoolID)
				TestListResponse(t, res, err, POST, detailed, increment)
				defaultLBPoolMemberID = res.ID
			}
		})

		t.Run("List Load Balancer Pools", func(t *testing.T) {
			res, _, err := ListLoadBalancerPools(vpcService, *defaultLBID)
			TestListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("List Load Balancer Pool Members", func(t *testing.T) {
			res, _, err := ListLoadBalancerPoolMembers(vpcService, *defaultLBID, *defaultLBPoolID)
			TestListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Load Balancer Pool", func(t *testing.T) {
			res, _, err := GetLoadBalancerPool(vpcService, *defaultLBID, *defaultLBPoolID)
			TestListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Load Balancer Pool Member", func(t *testing.T) {
			res, _, err := GetLoadBalancerPoolMember(vpcService, *defaultLBID, *defaultLBPoolID, *defaultLBPoolMemberID)
			TestListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Update Load Balancer Pool Member", func(t *testing.T) {
			res, _, err := UpdateLoadBalancerPoolMember(vpcService, *defaultLBID, *defaultLBPoolID, *defaultLBPoolMemberID)
			TestResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Update Load Balancer Pool", func(t *testing.T) {
			res, _, err := UpdateLoadBalancerPool(vpcService, *defaultLBID, *defaultLBPoolID)
			TestResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Delete Load Balancer Pool Member", func(t *testing.T) {
			statusChanged := PollLB(vpcService, *defaultLBID, "active", 5)
			if statusChanged {
				res, err := DeleteLoadBalancerPoolMember(vpcService, *defaultLBID, *defaultLBPoolID, *defaultLBPoolMemberID)
				TestDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
			}
		})

		var newPoolMemberID *string
		t.Run("Update Load Balancer Add Pool Member", func(t *testing.T) {
			statusChanged := PollLB(vpcService, *defaultLBID, "active", 5)
			if statusChanged {
				res, _, err := UpdateLoadBalancerPoolMembers(vpcService, *defaultLBID, *defaultLBPoolID)
				TestResponse(t, res, err, PATCH, detailed, increment)
				newPoolMemberID = res.Members[0].ID
			}
		})

		t.Run("Delete Load Balancer Pool Member Added ", func(t *testing.T) {
			statusChanged := PollLB(vpcService, *defaultLBID, "active", 5)
			if statusChanged {
				res, err := DeleteLoadBalancerPoolMember(vpcService, *defaultLBID, *defaultLBPoolID, *newPoolMemberID)
				TestDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
			}
		})

		t.Run("Delete Load Balancer Pool", func(t *testing.T) {
			statusChanged := PollLB(vpcService, *defaultLBID, "active", 5)
			if statusChanged {
				res, err := DeleteLoadBalancerPool(vpcService, *defaultLBID, *defaultLBPoolID)
				TestDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
			}
		})
		t.Run("Delete Load Balancer", func(t *testing.T) {
			statusChanged := PollLB(vpcService, *defaultLBID, "active", 5)
			if statusChanged {
				res, err := DeleteLoadBalancer(vpcService, *defaultLBID)
				TestDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
			}
		})
	})
	printTestSummary()
}

func TestVPCVpn(t *testing.T) {
	vpcService := createVpcGen2Service(t)
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
			TestResponse(t, res, err, POST, detailed, increment)
			createdIkePolicyID = res.ID
		})

		t.Run("Create Ipsec Policy", func(t *testing.T) {
			name := "go-ipsec-1-" + strconv.FormatInt(tunix, 10)
			res, _, err := CreateIpsecPolicy(vpcService, name)
			TestResponse(t, res, err, POST, detailed, increment)
			createdIpsecPolicyID = res.ID
		})

		t.Run("Create Vpn Gateway", func(t *testing.T) {
			name := "go-vpngateway-1-" + strconv.FormatInt(tunix, 10)
			res, _, err := CreateVpnGateway(vpcService, *defaultSubnetID, name)
			TestResponse(t, res, err, POST, detailed, increment)
			createdVpnGatewayID = res.ID
		})

		t.Run("Create Vpn Gateway Connections", func(t *testing.T) {
			name := "go-vpngateway-conn-1-" + strconv.FormatInt(tunix, 10)
			statusChanged := PollVpnGateway(vpcService, *createdVpnGatewayID, "available", 10)
			if statusChanged {
				res, _, err := CreateVpnGatewayConnection(vpcService, *createdVpnGatewayID, name)
				TestResponse(t, res, err, POST, detailed, increment)
				createdVpnGatewayConnID = res.ID
			}
		})

		t.Run("List Ike Policies", func(t *testing.T) {
			res, _, err := ListIkePolicies(vpcService)
			TestListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("List Ipsec Policies", func(t *testing.T) {
			res, _, err := ListIpsecPolicies(vpcService)
			TestListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("List Vpn Gateway", func(t *testing.T) {
			res, _, err := ListVpnGateways(vpcService)
			TestListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("List Vpn Gateway Connections", func(t *testing.T) {
			res, _, err := ListVpnGatewayConnections(vpcService, *createdVpnGatewayID)
			TestListResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Ike Policies", func(t *testing.T) {
			res, _, err := GetIkePolicy(vpcService, *createdIkePolicyID)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Ipsec Policies", func(t *testing.T) {
			res, _, err := GetIpsecPolicy(vpcService, *createdIpsecPolicyID)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Vpn Gateway", func(t *testing.T) {
			res, _, err := GetVpnGateway(vpcService, *createdVpnGatewayID)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get VpnGateway Connection", func(t *testing.T) {
			res, _, err := GetVpnGatewayConnection(vpcService, *createdVpnGatewayID, *createdVpnGatewayConnID)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("List VpnGateway Ipsec Policy Connections", func(t *testing.T) {
			res, _, err := ListVpnGatewayIpsecPolicyConnections(vpcService, *createdIpsecPolicyID)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("List VpnGateway Ike Policy Connections", func(t *testing.T) {
			res, _, err := ListVpnGatewayIkePolicyConnections(vpcService, *createdIkePolicyID)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("List VpnGateway Connection Local Cidrs", func(t *testing.T) {
			res, _, err := ListVpnGatewayConnectionLocalCidrs(vpcService, *createdVpnGatewayID, *createdVpnGatewayConnID)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("List VpnGateway Connection Peer Cidrs", func(t *testing.T) {
			res, _, err := ListVpnGatewayConnectionPeerCidrs(vpcService, *createdVpnGatewayID, *createdVpnGatewayConnID)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get Vpn Gateway Connection LocalCidr", func(t *testing.T) {
			res, err := GetVpnGatewayConnectionLocalCidr(vpcService, *createdVpnGatewayID, *createdVpnGatewayConnID, "192.132.0.0", "28")
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Get VpnGateway Connection PeerCidr", func(t *testing.T) {
			res, err := GetVpnGatewayConnectionPeerCidr(vpcService, *createdVpnGatewayID, *createdVpnGatewayConnID, "197.155.0.0", "28")
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("SetVpnGatewayConnectionLocalCidr", func(t *testing.T) {
			res, err := SetVpnGatewayConnectionLocalCidr(vpcService, *createdVpnGatewayID, *createdVpnGatewayConnID, "192.134.0.0", "28")
			TestResponse(t, res, err, GET, detailed, increment)
		})
		t.Run("GetVpnGatewayConnectionLocalCidr", func(t *testing.T) {
			res, err := GetVpnGatewayConnectionLocalCidr(vpcService, *createdVpnGatewayID, *createdVpnGatewayConnID, "192.134.0.0", "28")
			TestResponse(t, res, err, GET, detailed, increment)
		})
		t.Run("DeleteVpnGatewayConnectionLocalCidr", func(t *testing.T) {
			res, err := DeleteVpnGatewayConnectionLocalCidr(vpcService, *createdVpnGatewayID, *createdVpnGatewayConnID, "192.134.0.0", "28")
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("SetVpnGatewayConnectionPeerCidr", func(t *testing.T) {
			res, err := SetVpnGatewayConnectionPeerCidr(vpcService, *createdVpnGatewayID, *createdVpnGatewayConnID, "192.157.0.0", "28")
			TestResponse(t, res, err, GET, detailed, increment)
		})
		t.Run("GetVpnGatewayConnectionPeerCidr", func(t *testing.T) {
			res, err := GetVpnGatewayConnectionPeerCidr(vpcService, *createdVpnGatewayID, *createdVpnGatewayConnID, "192.157.0.0", "28")
			TestResponse(t, res, err, GET, detailed, increment)
		})
		t.Run("DeleteVpnGatewayConnectionPeerCidr", func(t *testing.T) {
			res, err := DeleteVpnGatewayConnectionPeerCidr(vpcService, *createdVpnGatewayID, *createdVpnGatewayConnID, "192.157.0.0", "28")
			TestResponse(t, res, err, GET, detailed, increment)
		})
		t.Run("Update Ike Policies", func(t *testing.T) {
			res, _, err := UpdateIkePolicy(vpcService, *createdIkePolicyID)
			TestResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Update Ipsec Policies", func(t *testing.T) {
			res, _, err := UpdateIpsecPolicy(vpcService, *createdIpsecPolicyID)
			TestResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Update VPN gateway", func(t *testing.T) {
			res, _, err := UpdateVpnGateway(vpcService, *createdVpnGatewayID, "go-vpngateway-2")
			TestResponse(t, res, err, PATCH, detailed, increment)
		})

		t.Run("Update VpnGateway Connection", func(t *testing.T) {
			res, _, err := UpdateVpnGatewayConnection(vpcService, *createdVpnGatewayID, *createdVpnGatewayConnID, "go-vpngateway-connection-2")
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Delete VpnGateway Connection", func(t *testing.T) {
			res, err := DeleteVpnGatewayConnection(vpcService, *createdVpnGatewayID, *createdVpnGatewayConnID)
			TestResponse(t, res, err, GET, detailed, increment)
		})

		t.Run("Delete Ike Policies", func(t *testing.T) {
			res, err := DeleteIkePolicy(vpcService, *createdIkePolicyID)
			TestDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

		t.Run("Delete Ipsec Policies", func(t *testing.T) {
			res, err := DeleteIpsecPolicy(vpcService, *createdIpsecPolicyID)
			TestDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

		t.Run("Delete VPN gateway", func(t *testing.T) {
			res, err := DeleteVpnGateway(vpcService, *createdVpnGatewayID)
			TestDeleteResponse(t, res, err, DELETE, res.StatusCode, detailed, increment)
		})

	})
	printTestSummary()
}
