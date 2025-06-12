// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccReinitIBMISBareMetalServerreinitbasic(t *testing.T) {
	var server string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-server-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-sshname-%d", acctest.RandIntRange(10, 100))
	userdata := "a"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISBareMetalServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISBareMetalServerInitializationReplaceConfig(vpcname, subnetname, sshname, publicKey, name, userdata),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISBareMetalServerExists("ibm_is_bare_metal_server.testacc_bms", server),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "user_data", userdata),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "image", acc.IsBareMetalServerImage),
				),
			},
			{
				Config: testAccCheckIBMISBareMetalServerInitializationReplaceConfigUpdate(vpcname, subnetname, sshname, publicKey, name, userdata, userdata),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISBareMetalServerExists("ibm_is_bare_metal_server.testacc_bms", server),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "user_data", userdata),
					resource.TestCheckResourceAttrSet(
						"ibm_is_bare_metal_server_initialization.testacc_bms_initialization", "id"),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "image", acc.IsBareMetalServerImage),
				),
			},
		},
	})
}

func testAccCheckIBMISBareMetalServerInitializationReplaceConfig(vpcname, subnetname, sshname, publicKey, name, userData1 string) string {
	return fmt.Sprintf(`
		resource "ibm_is_vpc" "testacc_vpc" {
			name = "%s"
		}
	  
		resource "ibm_is_subnet" "testacc_subnet" {
			name            			= "%s"
			vpc             			= ibm_is_vpc.testacc_vpc.id
			zone            			= "%s"
			total_ipv4_address_count 	= 16
		}
	  
		resource "ibm_is_ssh_key" "testacc_sshkey" {
			name       			= "%s"
			public_key 			= "%s"
		}
	  
		resource "ibm_is_bare_metal_server" "testacc_bms" {
			profile 			= "%s"
			name 				= "%s"
			image 				= "%s"
			zone 				= "%s"
			user_data         	= "%s"
			keys 				= [ibm_is_ssh_key.testacc_sshkey.id]
			primary_network_interface {
				subnet     		= ibm_is_subnet.testacc_subnet.id
			}
			vpc 				= ibm_is_vpc.testacc_vpc.id
		}
		
`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, acc.IsBareMetalServerProfileName, name, acc.IsBareMetalServerImage, acc.ISZoneName, userData1)
}
func testAccCheckIBMISBareMetalServerInitializationReplaceConfigUpdate(vpcname, subnetname, sshname, publicKey, name, userData1, userData2 string) string {
	return fmt.Sprintf(`
		resource "ibm_is_vpc" "testacc_vpc" {
			name = "%s"
		}
	  
		resource "ibm_is_subnet" "testacc_subnet" {
			name            			= "%s"
			vpc             			= ibm_is_vpc.testacc_vpc.id
			zone            			= "%s"
			total_ipv4_address_count 	= 16
		}
	  
		resource "ibm_is_ssh_key" "testacc_sshkey" {
			name       			= "%s"
			public_key 			= "%s"
		}
	  
		resource "ibm_is_bare_metal_server" "testacc_bms" {
			profile 			= "%s"
			name 				= "%s"
			image 				= "%s"
			zone 				= "%s"
			user_data         	= "%s"
			keys 				= [ibm_is_ssh_key.testacc_sshkey.id]
			primary_network_interface {
				subnet     		= ibm_is_subnet.testacc_subnet.id
			}
			vpc 				= ibm_is_vpc.testacc_vpc.id
			lifecycle {
				ignore_changes = [ image, keys, user_data ]
			}
		}
		resource "ibm_is_bare_metal_server_initialization" "testacc_bms_initialization" {
			bare_metal_server 	= ibm_is_bare_metal_server.testacc_bms.id
			image				= "%s"
			keys				= [ibm_is_ssh_key.testacc_sshkey.id]
			user_data         	= "%s"
		}
		
`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, acc.IsBareMetalServerProfileName, name, acc.IsBareMetalServerImage, acc.ISZoneName, userData1, acc.IsBareMetalServerImage2, userData2)
}
