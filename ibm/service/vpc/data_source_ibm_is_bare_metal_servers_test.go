// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISBMSsDataSource_basic(t *testing.T) {
	resName := "data.ibm_is_bare_metal_servers.test1"
	var server string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-server-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-sshname-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISBMSsDataSourceConfig(vpcname, subnetname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISBareMetalServerExists("ibm_is_bare_metal_server.testacc_bms", server),
					resource.TestCheckResourceAttrSet(resName, "servers.0.name"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.id"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.memory"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.primary_network_interface.0.primary_ip.0.address"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.primary_network_interface.0.primary_ip.0.name"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.primary_network_interface.0.primary_ip.0.href"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.primary_network_interface.0.primary_ip.0.reserved_ip"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.primary_network_interface.0.primary_ip.0.resource_type"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.primary_network_interface.0.port_speed"),
				),
			},
		},
	})
}

func TestAccIBMISBMSsDataSource_firmwareUpdate(t *testing.T) {
	resName := "data.ibm_is_bare_metal_servers.test1"
	var server string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-server-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-sshname-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISBMSsDataSourceConfig(vpcname, subnetname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISBareMetalServerExists("ibm_is_bare_metal_server.testacc_bms", server),
					resource.TestCheckResourceAttrSet(resName, "servers.0.name"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.id"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.firmware_update_type_available"),
				),
			},
		},
	})
}

func TestAccIBMISBMSsDataSource_MetadataService(t *testing.T) {
	resName := "data.ibm_is_bare_metal_servers.test1"
	var server string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-server-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-sshname-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISBMSsDataSourceMetadataServiceConfig(vpcname, subnetname, sshname, publicKey, name, true, "https"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISBareMetalServerExists("ibm_is_bare_metal_server.testacc_bms", server),
					resource.TestCheckResourceAttrSet(resName, "servers.0.metadata_service.0.enabled"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.metadata_service.0.protocol"),
				),
			},
		},
	})
}

func testAccCheckIBMISBMSsDataSourceMetadataServiceConfig(vpcname, subnetname, sshname, publicKey, name string, enabled bool, protocol string) string {
	// status filter defaults to empty
	return testAccCheckIBMISBareMetalServerMetadataServiceConfig(vpcname, subnetname, sshname, publicKey, name, enabled, protocol) + fmt.Sprintf(`
      data "ibm_is_bare_metal_servers" "test1" {
	  		depends_on = [ ibm_is_bare_metal_server.testacc_bms ]
      }`)
}

func testAccCheckIBMISBMSsDataSourceConfig(vpcname, subnetname, sshname, publicKey, name string) string {
	// status filter defaults to empty
	return testAccCheckIBMISBareMetalServerConfig(vpcname, subnetname, sshname, publicKey, name) + fmt.Sprintf(`
      data "ibm_is_bare_metal_servers" "test1" {
      }`)
}
func TestAccIBMISBMSsDataSourceVNI_basic(t *testing.T) {
	resName := "data.ibm_is_bare_metal_servers.test1"
	var server string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-server-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-sshname-%d", acctest.RandIntRange(10, 100))
	vniname := fmt.Sprintf("tf-vni-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISBMSsDataSourceVNIConfig(vpcname, subnetname, sshname, publicKey, vniname, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISBareMetalServerExists("ibm_is_bare_metal_server.testacc_bms", server),
					resource.TestCheckResourceAttrSet(resName, "servers.0.name"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.id"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.memory"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.primary_network_interface.0.primary_ip.0.address"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.primary_network_interface.0.primary_ip.0.name"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.primary_network_interface.0.primary_ip.0.href"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.primary_network_interface.0.primary_ip.0.reserved_ip"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.primary_network_interface.0.primary_ip.0.resource_type"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.primary_network_interface.0.port_speed"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.primary_network_attachment.#"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.primary_network_attachment.0.href"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.primary_network_attachment.0.id"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.primary_network_attachment.0.name"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.primary_network_attachment.0.primary_ip.#"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.primary_network_attachment.0.resource_type"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.primary_network_attachment.0.subnet.#"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.primary_network_attachment.0.virtual_network_interface.#"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.network_attachments.#"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.network_attachments.0.href"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.network_attachments.0.id"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.network_attachments.0.name"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.network_attachments.0.primary_ip.#"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.network_attachments.0.resource_type"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.network_attachments.0.subnet.#"),
					resource.TestCheckResourceAttrSet(resName, "servers.0.network_attachments.0.virtual_network_interface.#"),
				),
			},
		},
	})
}

func testAccCheckIBMISBMSsDataSourceVNIConfig(vpcname, subnetname, sshname, publicKey, vniname, name string) string {
	// status filter defaults to empty
	return testAccCheckIBMISBareMetalServerVNIConfig(vpcname, subnetname, sshname, publicKey, vniname, name) + fmt.Sprintf(`
      data "ibm_is_bare_metal_servers" "test1" {
		depends_on = [ ibm_is_bare_metal_server.testacc_bms ]
      }`)
}
