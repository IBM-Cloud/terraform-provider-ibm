// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsVirtualNetworkInterfaceBasic(t *testing.T) {
	var conf vpcv1.VirtualNetworkInterface
	vpcname := fmt.Sprintf("tfvpngw-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfvpngw-subnet-%d", acctest.RandIntRange(10, 100))
	vniname := fmt.Sprintf("tfvpngw-createname-%d", acctest.RandIntRange(10, 100))
	tag1 := "env:test"
	tag2 := "env:dev"
	tag3 := "env:prod"
	enable_infrastructure_nat := true
	allow_ip_spoofing := true
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsVirtualNetworkInterfaceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualNetworkInterfaceConfigBasic(vpcname, subnetname, vniname, tag1, tag2, tag3, enable_infrastructure_nat, allow_ip_spoofing, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsVirtualNetworkInterfaceExists("ibm_is_virtual_network_interface.testacc_vni", conf),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "subnet"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "name"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "crn"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "vpc.#"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "zone"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "primary_ip.#"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "ips.#"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "id"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "enable_infrastructure_nat", "true"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "lifecycle_state", "stable"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "resource_type", "virtual_network_interface"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "resource_group"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "tags.#", "3"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "allow_ip_spoofing", "true"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualNetworkInterfaceConfigBasic(vpcname, subnetname, vniname, tag1, tag2, tag3, enable_infrastructure_nat, allow_ip_spoofing, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsVirtualNetworkInterfaceExists("ibm_is_virtual_network_interface.testacc_vni", conf),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "subnet"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "name"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "crn"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "vpc.#"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "zone"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "primary_ip.#"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "ips.#"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "tags.#", "2"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "id"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "enable_infrastructure_nat", "true"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "lifecycle_state", "stable"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "resource_type", "virtual_network_interface"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "resource_group"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "allow_ip_spoofing", "true"),
				),
			},
		},
	})
}
func TestAccIBMIsVirtualNetworkInterfaceResourceGroup(t *testing.T) {
	var conf vpcv1.VirtualNetworkInterface
	vpcname := fmt.Sprintf("tfvpngw-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfvpngw-subnet-%d", acctest.RandIntRange(10, 100))
	vniname := fmt.Sprintf("tfvpngw-createname-%d", acctest.RandIntRange(10, 100))
	enable_infrastructure_nat := true
	allow_ip_spoofing := true
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsVirtualNetworkInterfaceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualNetworkInterfaceConfigResourceGroup(vpcname, subnetname, vniname, enable_infrastructure_nat, allow_ip_spoofing),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsVirtualNetworkInterfaceExists("ibm_is_virtual_network_interface.testacc_vni", conf),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "subnet"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "name"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "crn"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "vpc.#"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "zone"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "primary_ip.#"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "ips.#"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "id"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "enable_infrastructure_nat", "true"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "lifecycle_state", "stable"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "resource_type", "virtual_network_interface"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "resource_group", acc.IsResourceGroupID),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "resource_group"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "allow_ip_spoofing", "true"),
				),
			},
		},
	})
}
func TestAccIBMIsVirtualNetworkInterfaceResourceGroupUpdate(t *testing.T) {
	var conf vpcv1.VirtualNetworkInterface
	vpcname := fmt.Sprintf("tfvpngw-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfvpngw-subnet-%d", acctest.RandIntRange(10, 100))
	vniname := fmt.Sprintf("tfvpngw-createname-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsVirtualNetworkInterfaceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualNetworkInterfaceConfigResourceGroupUpdated(vpcname, subnetname, vniname, acc.IsResourceGroupID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsVirtualNetworkInterfaceExists("ibm_is_virtual_network_interface.testacc_vni", conf),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "subnet"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "name"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "crn"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "vpc.#"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "zone"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "primary_ip.#"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "ips.#"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "id"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "resource_type", "virtual_network_interface"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "resource_group", acc.IsResourceGroupID),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "resource_group"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualNetworkInterfaceConfigResourceGroupUpdated(vpcname, subnetname, vniname, acc.IsResourceGroupIDUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsVirtualNetworkInterfaceExists("ibm_is_virtual_network_interface.testacc_vni", conf),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "subnet"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "name"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "crn"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "vpc.#"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "zone"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "primary_ip.#"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "ips.#"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "id"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "resource_type", "virtual_network_interface"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "resource_group", acc.IsResourceGroupIDUpdate),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "resource_group"),
				),
			},
		},
	})
}

func TestAccIBMIsVirtualNetworkInterfaceAllArgs(t *testing.T) {
	var conf vpcv1.VirtualNetworkInterface
	vpcname := fmt.Sprintf("tfvpngw-vpc-%d", acctest.RandIntRange(10, 100))
	sgname1 := fmt.Sprintf("tfvpngw-sg-%d", acctest.RandIntRange(10, 100))
	sgname2 := fmt.Sprintf("tfvpngw-sg-%d", acctest.RandIntRange(10, 100))
	reservedipname := fmt.Sprintf("tfvpngw-ip-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfvpngw-subnet-%d", acctest.RandIntRange(10, 100))
	vniname := fmt.Sprintf("tfvpngw-createname-%d", acctest.RandIntRange(10, 100))
	enable_infrastructure_nat := true
	allow_ip_spoofing := true

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsVirtualNetworkInterfaceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualNetworkInterfaceConfig(vpcname, subnetname, vniname, enable_infrastructure_nat, allow_ip_spoofing),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsVirtualNetworkInterfaceExists("ibm_is_virtual_network_interface.testacc_vni", conf),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "subnet"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "name"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "crn"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "vpc.#"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "zone"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "primary_ip.#"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "ips.#"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "ips.#", "1"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "security_groups.#", "1"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "id"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "enable_infrastructure_nat", "true"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "lifecycle_state", "stable"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "resource_type", "virtual_network_interface"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "resource_group"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "allow_ip_spoofing", "true"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualNetworkInterfaceConfigUpdate(vpcname, subnetname, vniname, sgname1, sgname2, reservedipname, enable_infrastructure_nat, allow_ip_spoofing),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "ips.#", "2"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "security_groups.#", "2"),
				),
			},
			// resource.TestStep{
			// 	ResourceName:      "ibm_is_virtual_network_interface.testacc_vni",
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// },
		},
	})
}

func TestAccIBMIsVirtualNetworkInterfacePStateFiltering(t *testing.T) {
	var conf vpcv1.VirtualNetworkInterface
	vpcname := fmt.Sprintf("tfvpngw-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfvpngw-subnet-%d", acctest.RandIntRange(10, 100))
	vniname := fmt.Sprintf("tfvpngw-createname-%d", acctest.RandIntRange(10, 100))
	tag1 := "env:test"
	tag2 := "env:dev"
	tag3 := "env:prod"
	enable_infrastructure_nat := true
	allow_ip_spoofing := true
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIsVirtualNetworkInterfaceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualNetworkInterfaceConfigPStateFiltering(vpcname, subnetname, vniname, tag1, tag2, tag3, enable_infrastructure_nat, allow_ip_spoofing, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsVirtualNetworkInterfaceExists("ibm_is_virtual_network_interface.testacc_vni", conf),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "subnet"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "name"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "crn"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "vpc.#"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "zone"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "primary_ip.#"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "ips.#"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "id"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "enable_infrastructure_nat", "true"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "lifecycle_state", "stable"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "resource_type", "virtual_network_interface"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "resource_group"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "tags.#", "3"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "allow_ip_spoofing", "true"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "protocol_state_filtering_mode", "auto"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsVirtualNetworkInterfaceConfigPStateFilteringUpdate(vpcname, subnetname, vniname, tag1, tag2, tag3, enable_infrastructure_nat, allow_ip_spoofing, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIsVirtualNetworkInterfaceExists("ibm_is_virtual_network_interface.testacc_vni", conf),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "subnet"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "name"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "created_at"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "crn"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "href"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "vpc.#"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "zone"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "primary_ip.#"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "ips.#"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "tags.#", "2"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "id"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "enable_infrastructure_nat", "true"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "lifecycle_state", "stable"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "resource_type", "virtual_network_interface"),
					resource.TestCheckResourceAttrSet("ibm_is_virtual_network_interface.testacc_vni", "resource_group"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "allow_ip_spoofing", "true"),
					resource.TestCheckResourceAttr("ibm_is_virtual_network_interface.testacc_vni", "protocol_state_filtering_mode", "disabled"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVirtualNetworkInterfaceConfigBasic(vpcname, subnetname, vniname, tag1, tag2, tag3 string, enablenat, allowipspoofing, isUpdate bool) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	
	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		total_ipv4_address_count = 16
	
	}
	
	resource "ibm_is_virtual_network_interface" "testacc_vni"{
		name = "%s"
		subnet = ibm_is_subnet.testacc_subnet.id
		enable_infrastructure_nat = %t
		allow_ip_spoofing = %t
		tags = %t ? ["%s", "%s"] : ["%s", "%s", "%s"]
	}
	`, vpcname, subnetname, acc.ISZoneName, vniname, enablenat, allowipspoofing, isUpdate, tag1, tag2, tag1, tag2, tag3)
}

func testAccCheckIBMIsVirtualNetworkInterfaceConfigPStateFiltering(vpcname, subnetname, vniname, tag1, tag2, tag3 string, enablenat, allowipspoofing, isUpdate bool) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	
	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		total_ipv4_address_count = 16
	
	}
	
	resource "ibm_is_virtual_network_interface" "testacc_vni"{
		name = "%s"
		subnet = ibm_is_subnet.testacc_subnet.id
		protocol_state_filtering_mode = "auto"
		enable_infrastructure_nat = %t
		allow_ip_spoofing = %t
		tags = %t ? ["%s", "%s"] : ["%s", "%s", "%s"]
	}
	`, vpcname, subnetname, acc.ISZoneName, vniname, enablenat, allowipspoofing, isUpdate, tag1, tag2, tag1, tag2, tag3)
}

func testAccCheckIBMIsVirtualNetworkInterfaceConfigPStateFilteringUpdate(vpcname, subnetname, vniname, tag1, tag2, tag3 string, enablenat, allowipspoofing, isUpdate bool) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	
	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		total_ipv4_address_count = 16
	
	}
	
	resource "ibm_is_virtual_network_interface" "testacc_vni"{
		name = "%s"
		subnet = ibm_is_subnet.testacc_subnet.id
		protocol_state_filtering_mode = "disabled"
		enable_infrastructure_nat = %t
		allow_ip_spoofing = %t
		tags = %t ? ["%s", "%s"] : ["%s", "%s", "%s"]
	}
	`, vpcname, subnetname, acc.ISZoneName, vniname, enablenat, allowipspoofing, isUpdate, tag1, tag2, tag1, tag2, tag3)
}

func testAccCheckIBMIsVirtualNetworkInterfaceConfigResourceGroup(vpcname, subnetname, vniname string, enablenat, allowipspoofing bool) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	
	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		total_ipv4_address_count = 16
	
	}
	
	resource "ibm_is_virtual_network_interface" "testacc_vni"{
		name = "%s"
		subnet = ibm_is_subnet.testacc_subnet.id
		enable_infrastructure_nat = %t
		allow_ip_spoofing = %t
		resource_group = "%s"
	}
	`, vpcname, subnetname, acc.ISZoneName, vniname, enablenat, allowipspoofing, acc.IsResourceGroupID)
}

func testAccCheckIBMIsVirtualNetworkInterfaceConfigResourceGroupUpdated(vpcname, subnetname, vniname, resourcegroup string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	
	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		total_ipv4_address_count = 16
	
	}
	
	resource "ibm_is_virtual_network_interface" "testacc_vni"{
		name = "%s"
		subnet = ibm_is_subnet.testacc_subnet.id
		resource_group = "%s"
	}
	`, vpcname, subnetname, acc.ISZoneName, vniname, resourcegroup)
}

func testAccCheckIBMIsVirtualNetworkInterfaceConfig(vpcname, subnetname, vniname string, enablenat, allowipspoofing bool) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	
	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		total_ipv4_address_count = 16
	
	}
	resource "ibm_is_virtual_network_interface" "testacc_vni" {
		name = "%s"
		subnet = ibm_is_subnet.testacc_subnet.id
		enable_infrastructure_nat = %t
		allow_ip_spoofing = %t
		ips {
		    address = "10.240.0.7"
		}
		primary_ip {
		    address = "10.240.0.9"
		}
	}
	`, vpcname, subnetname, acc.ISZoneName, vniname, enablenat, allowipspoofing)
}
func testAccCheckIBMIsVirtualNetworkInterfaceConfigUpdate(vpcname, subnetname, vniname, sgname1, sgname2, reservedipname string, enablenat, allowipspoofing bool) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	
	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		total_ipv4_address_count = 16
	
	}
	resource "ibm_is_security_group" "testacc_security_group1" {
		name = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id
	}
	resource "ibm_is_security_group" "testacc_security_group2" {
		name = "%s"
		vpc  =ibm_is_vpc.testacc_vpc.id
	}
	resource "ibm_is_subnet_reserved_ip" "testacc_reservedips" {
		subnet = ibm_is_subnet.testacc_subnet.id
		name = "%s"
	}
	resource "ibm_is_virtual_network_interface" "testacc_vni" {
		name = "%s"
		subnet = ibm_is_subnet.testacc_subnet.id
		enable_infrastructure_nat = %t
		allow_ip_spoofing = %t
		ips {
		    address = "10.240.0.7"
		}
		ips {
		    reserved_ip = ibm_is_subnet_reserved_ip.testacc_reservedips.reserved_ip
		}
		primary_ip {
		    address = "10.240.0.9"
		}
		security_groups = [ibm_is_security_group.testacc_security_group2.id, ibm_is_security_group.testacc_security_group1.id]
	}
	`, vpcname, subnetname, acc.ISZoneName, sgname1, sgname2, reservedipname, vniname, enablenat, allowipspoofing)
}

func testAccCheckIBMIsVirtualNetworkInterfaceExists(n string, obj vpcv1.VirtualNetworkInterface) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		sess, err := vpcClient(acc.TestAccProvider.Meta())
		if err != nil {
			return err
		}

		getVirtualNetworkInterfaceOptions := &vpcv1.GetVirtualNetworkInterfaceOptions{}

		getVirtualNetworkInterfaceOptions.SetID(rs.Primary.ID)

		virtualNetworkInterface, _, err := sess.GetVirtualNetworkInterface(getVirtualNetworkInterfaceOptions)
		if err != nil {
			return err
		}

		obj = *virtualNetworkInterface
		return nil
	}
}

func testAccCheckIBMIsVirtualNetworkInterfaceDestroy(s *terraform.State) error {
	sess, err := vpcClient(acc.TestAccProvider.Meta())
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_virtual_network_interface" {
			continue
		}

		getVirtualNetworkInterfaceOptions := &vpcv1.GetVirtualNetworkInterfaceOptions{}

		getVirtualNetworkInterfaceOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := sess.GetVirtualNetworkInterface(getVirtualNetworkInterfaceOptions)

		if err == nil {
			return fmt.Errorf("VirtualNetworkInterface still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for VirtualNetworkInterface (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
