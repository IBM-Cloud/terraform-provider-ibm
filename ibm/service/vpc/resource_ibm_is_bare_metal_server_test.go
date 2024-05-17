// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMISBareMetalServer_basic(t *testing.T) {
	var server string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-server-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-sshname-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISBareMetalServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISBareMetalServerConfig(vpcname, subnetname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISBareMetalServerExists("ibm_is_bare_metal_server.testacc_bms", server),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "zone", acc.ISZoneName),
				),
			},
		},
	})
}
func TestAccIBMISBareMetalServerVNI_basic(t *testing.T) {
	var server string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-server-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	vniname := fmt.Sprintf("tfip-vni-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-sshname-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISBareMetalServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISBareMetalServerVNIConfig(vpcname, subnetname, sshname, publicKey, vniname, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISBareMetalServerExists("ibm_is_bare_metal_server.testacc_bms", server),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "zone", acc.ISZoneName),
				),
			},
		},
	})
}
func TestAccIBMISBareMetalServer_sg_update(t *testing.T) {
	var server string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-server-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-sshname-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISBareMetalServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISBareMetalServerConfig(vpcname, subnetname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISBareMetalServerExists("ibm_is_bare_metal_server.testacc_bms", server),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "primary_network_interface.0.security_groups.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMISBareMetalServerSgUpdateConfig(vpcname, subnetname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISBareMetalServerExists("ibm_is_bare_metal_server.testacc_bms", server),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group.testacc_sg1", "name", "test-security-group1"),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group.testacc_sg2", "name", "test-security-group2"),
					resource.TestCheckResourceAttr(
						"ibm_is_security_group.testacc_sg3", "name", "test-security-group3"),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "primary_network_interface.0.security_groups.#", "3"),
				),
			},
		},
	})
}
func TestAccIBMISBareMetalServer_SecureBoot_tpm(t *testing.T) {
	var server string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-server-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	tpm1 := "disabled"
	tpm2 := "tpm_2"
	secureBootTrue := true
	secureBootFalse := false

	sshname := fmt.Sprintf("tf-sshname-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISBareMetalServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISBareMetalServerSecureBootTpmConfig(vpcname, subnetname, sshname, publicKey, tpm2, name, secureBootFalse),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISBareMetalServerExists("ibm_is_bare_metal_server.testacc_bms", server),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "enable_secure_boot", fmt.Sprintf("%t", secureBootFalse)),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "trusted_platform_module.0.mode", tpm2),
				),
			},
			{
				Config: testAccCheckIBMISBareMetalServerSecureBootTpmConfig(vpcname, subnetname, sshname, publicKey, tpm1, name, secureBootTrue),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISBareMetalServerExists("ibm_is_bare_metal_server.testacc_bms", server),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "enable_secure_boot", fmt.Sprintf("%t", secureBootTrue)),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "trusted_platform_module.0.mode", tpm1),
				),
			},
		},
	})
}
func TestAccIBMISBareMetalServer_testZ(t *testing.T) {
	var server string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-server-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-sshname-%d", acctest.RandIntRange(10, 100))
	profileName := "mz2d-metal-2x64"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISBareMetalServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISBareMetalServerZSingleNicConfig(vpcname, subnetname, sshname, publicKey, name, profileName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISBareMetalServerExists("ibm_is_bare_metal_server.testacc_bms", server),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "memory", "64"),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "profile", profileName),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "status", "running"),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "primary_network_interface.0.enable_infrastructure_nat", "true"),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "primary_network_interface.0.interface_type", "hipersocket"),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "primary_network_interface.0.name", "test-bm-nic-1"),
				),
			},
		},
	})
}
func TestAccIBMISBareMetalServer_testMultinicZ(t *testing.T) {
	var server string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-server-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-sshname-%d", acctest.RandIntRange(10, 100))
	profileName := "mz2d-metal-16x512"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISBareMetalServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISBareMetalServerZConfig(vpcname, subnetname, sshname, publicKey, name, profileName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISBareMetalServerExists("ibm_is_bare_metal_server.testacc_bms", server),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "memory", "512"),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "profile", profileName),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "status", "running"),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "network_interfaces.0.enable_infrastructure_nat", "true"),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "network_interfaces.0.interface_type", "hipersocket"),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "network_interfaces.0.name", "test-bm-nic-2"),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "primary_network_interface.0.enable_infrastructure_nat", "true"),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "primary_network_interface.0.interface_type", "hipersocket"),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "primary_network_interface.0.name", "test-bm-nic-1"),
				),
			},
		},
	})
}
func TestAccIBMISBareMetalServer_multi_nic(t *testing.T) {
	var server string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-server-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-sshname-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISBareMetalServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISBareMetalServerMultiNicConfig(vpcname, subnetname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISBareMetalServerExists("ibm_is_bare_metal_server.testacc_bms", server),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "network_interfaces.0.vlan", "3"),
				),
			},
		},
	})
}
func TestAccIBMISBareMetalServer_mix_multi_nic(t *testing.T) {
	var server string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-server-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-sshname-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISBareMetalServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISBareMetalServerMixMultiNicConfig(vpcname, subnetname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISBareMetalServerExists("ibm_is_bare_metal_server.testacc_bms", server),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "network_interfaces.#", "3"),
				),
			},
		},
	})
}
func TestAccIBMISBareMetalServer_multi_nic_with_allow_float(t *testing.T) {
	var server string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-server-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-sshname-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISBareMetalServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISBareMetalServerMultiNicWithAllowFloatConfig(vpcname, subnetname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISBareMetalServerExists("ibm_is_bare_metal_server.testacc_bms", server),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "network_interfaces.0.vlan", "102"),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", "name", "eth21"),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", "allow_ip_spoofing", "false"),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", "allow_interface_to_float", "true"),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server_network_interface_allow_float.bms_nic", "enable_infrastructure_nat", "true"),
				),
			},
		},
	})
}
func TestAccIBMISBareMetalServer_basic_reserved_ip(t *testing.T) {
	var server string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-server-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-sshname-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISBareMetalServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISBareMetalServerReservedIpConfig(vpcname, subnetname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISBareMetalServerExists("ibm_is_bare_metal_server.testacc_bms", server),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "zone", acc.ISZoneName),
				),
			},
		},
	})
}

func testAccCheckIBMISBareMetalServerDestroy(s *terraform.State) error {

	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_bare_metal_server" {
			continue
		}

		getbmsoptions := &vpcv1.GetBareMetalServerOptions{
			ID: &rs.Primary.ID,
		}
		_, _, err := sess.GetBareMetalServer(getbmsoptions)
		if err == nil {
			return fmt.Errorf("Bare metal server still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISBareMetalServerExists(n, ip string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}
		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getbmsoptions := &vpcv1.GetBareMetalServerOptions{
			ID: &rs.Primary.ID,
		}
		bms, _, err := sess.GetBareMetalServer(getbmsoptions)
		if err != nil {
			return err
		}
		ip = *bms.ID
		return nil
	}
}

func testAccCheckIBMISBareMetalServerConfig(vpcname, subnetname, sshname, publicKey, name string) string {
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
			keys 				= [ibm_is_ssh_key.testacc_sshkey.id]
			primary_network_interface {
				subnet     		= ibm_is_subnet.testacc_subnet.id
			}
			vpc 				= ibm_is_vpc.testacc_vpc.id
		}
`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, acc.IsBareMetalServerProfileName, name, acc.IsBareMetalServerImage, acc.ISZoneName)
}
func testAccCheckIBMISBareMetalServerVNIConfig(vpcname, subnetname, sshname, publicKey, vniname, name string) string {
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
		resource "ibm_is_virtual_network_interface" "testacc_vni"{
			name = "%s"
			subnet = ibm_is_subnet.testacc_subnet.id
			enable_infrastructure_nat = true
			allow_ip_spoofing = true
		}
		resource "ibm_is_bare_metal_server" "testacc_bms" {
			profile 			= "%s"
			name 				= "%s"
			image 				= "%s"
			zone 				= "%s"
			keys 				= [ibm_is_ssh_key.testacc_sshkey.id]
			primary_network_attachment {
				name = "test-vni-100-102"
				virtual_network_interface { 
					id = ibm_is_virtual_network_interface.testacc_vni.id
				}
				allowed_vlans = [100, 102]
			}
			vpc 				= ibm_is_vpc.testacc_vpc.id
		}
`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, vniname, acc.IsBareMetalServerProfileName, name, acc.IsBareMetalServerImage, acc.ISZoneName)
}
func testAccCheckIBMISBareMetalServerSgUpdateConfig(vpcname, subnetname, sshname, publicKey, name string) string {
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


		resource "ibm_is_security_group" "testacc_sg1" {
			name = "test-security-group1"
			vpc  = ibm_is_vpc.testacc_vpc.id
		}
		resource "ibm_is_security_group" "testacc_sg2" {
			name = "test-security-group2"
			vpc  = ibm_is_vpc.testacc_vpc.id
		}
		resource "ibm_is_security_group" "testacc_sg3" {
			name = "test-security-group3"
			vpc  = ibm_is_vpc.testacc_vpc.id
		}
	  
		resource "ibm_is_bare_metal_server" "testacc_bms" {
			profile 			= "%s"
			name 				= "%s"
			image 				= "%s"
			zone 				= "%s"
			keys 				= [ibm_is_ssh_key.testacc_sshkey.id]
			primary_network_interface {
				subnet     		= ibm_is_subnet.testacc_subnet.id
				security_groups = [ibm_is_security_group.testacc_sg1.id, ibm_is_security_group.testacc_sg2.id, ibm_is_security_group.testacc_sg3.id]
			}
			vpc 				= ibm_is_vpc.testacc_vpc.id
		}
`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, acc.IsBareMetalServerProfileName, name, acc.IsBareMetalServerImage, acc.ISZoneName)
}
func testAccCheckIBMISBareMetalServerSecureBootTpmConfig(vpcname, subnetname, sshname, publicKey, tpm, name string, secureBoot bool) string {
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
			enable_secure_boot  = %t
			trusted_platform_module {
				mode = "%s"
			}
			keys 				= [ibm_is_ssh_key.testacc_sshkey.id]
			primary_network_interface {
				subnet     		= ibm_is_subnet.testacc_subnet.id
				allowed_vlans   = [101,102,103]
			}
			vpc 				= ibm_is_vpc.testacc_vpc.id
		}
`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, acc.IsBareMetalServerProfileName, name, acc.IsBareMetalServerImage, acc.ISZoneName, secureBoot, tpm)
}
func testAccCheckIBMISBareMetalServerZConfig(vpcname, subnetname, sshname, publicKey, name, profileName string) string {
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
			keys 				= [ibm_is_ssh_key.testacc_sshkey.id]
			primary_network_interface {
				subnet     					= ibm_is_subnet.testacc_subnet.id
				interface_type          	= "hipersocket"
				name                    	= "test-bm-nic-1"
				enable_infrastructure_nat 	= true
			}
			network_interfaces {
				subnet     		        	= ibm_is_subnet.testacc_subnet.id
				interface_type          	= "hipersocket"
				name                    	= "test-bm-nic-2"
				enable_infrastructure_nat 	= true
			}
			vpc 				= ibm_is_vpc.testacc_vpc.id
		}
`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, profileName, name, acc.IsBareMetalServerImage, acc.ISZoneName)
}
func testAccCheckIBMISBareMetalServerZSingleNicConfig(vpcname, subnetname, sshname, publicKey, name, profileName string) string {
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
			keys 				= [ibm_is_ssh_key.testacc_sshkey.id]
			primary_network_interface {
				subnet     					= ibm_is_subnet.testacc_subnet.id
				interface_type          	= "hipersocket"
				name                    	= "test-bm-nic-1"
				enable_infrastructure_nat 	= true
			}
			vpc 				= ibm_is_vpc.testacc_vpc.id
		}
`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, profileName, name, acc.IsBareMetalServerImage, acc.ISZoneName)
}

func testAccCheckIBMISBareMetalServerMultiNicConfig(vpcname, subnetname, sshname, publicKey, name string) string {
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
			keys 				= [ibm_is_ssh_key.testacc_sshkey.id]
			primary_network_interface {
				subnet     		= ibm_is_subnet.testacc_subnet.id
				allowed_vlans                     = [102,103]
			}
			network_interfaces {
				subnet     		= ibm_is_subnet.testacc_subnet.id
				vlan            = 102
			}
			vpc 				= ibm_is_vpc.testacc_vpc.id
		}
`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, acc.IsBareMetalServerProfileName, name, acc.IsBareMetalServerImage, acc.ISZoneName)
}
func testAccCheckIBMISBareMetalServerMixMultiNicConfig(vpcname, subnetname, sshname, publicKey, name string) string {
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
			keys 				= [ibm_is_ssh_key.testacc_sshkey.id]
			primary_network_interface {
				name			= "primary-nic"
				subnet     		= ibm_is_subnet.testacc_subnet.id
				allowed_vlans   = [102,103]
			}
			network_interfaces {
				name			= "sec-nic-vlan112"
				subnet     		= ibm_is_subnet.testacc_subnet.id
				vlan            = 102
			}
			network_interfaces {
				name			= "sec-nic-pci"
				subnet     		= ibm_is_subnet.testacc_subnet.id
			}
			network_interfaces {
				name			= "sec-nic-pci108-09"
				subnet     		= ibm_is_subnet.testacc_subnet.id
				allowed_vlans   = [108,109]

			}
			vpc 				= ibm_is_vpc.testacc_vpc.id
		}
`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, acc.IsBareMetalServerProfileName, name, acc.IsBareMetalServerImage, acc.ISZoneName)
}
func testAccCheckIBMISBareMetalServerMultiNicWithAllowFloatConfig(vpcname, subnetname, sshname, publicKey, name string) string {
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
			keys 				= [ibm_is_ssh_key.testacc_sshkey.id]
			primary_network_interface {
				subnet     		= ibm_is_subnet.testacc_subnet.id
				allowed_vlans                     = [101,102,103]
			}
			network_interfaces {
				name			= "eth20"
				subnet     		= ibm_is_subnet.testacc_subnet.id
				vlan            = 102
			}
			vpc 				= ibm_is_vpc.testacc_vpc.id
		}

		resource ibm_is_bare_metal_server_network_interface_allow_float bms_nic {
			bare_metal_server 	= ibm_is_bare_metal_server.testacc_bms.id
			
			subnet 				= ibm_is_subnet.testacc_subnet.id
			name   				= "eth21"
			vlan 				= 101
		}

`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, acc.IsBareMetalServerProfileName, name, acc.IsBareMetalServerImage, acc.ISZoneName)
}
func testAccCheckIBMISBareMetalServerReservedIpConfig(vpcname, subnetname, sshname, publicKey, name string) string {
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
			keys 				= [ibm_is_ssh_key.testacc_sshkey.id]
			primary_network_interface {
				subnet     		= ibm_is_subnet.testacc_subnet.id
				primary_ip {
					auto_delete = true
					address		= "${replace(ibm_is_subnet.testacc_subnet.ipv4_cidr_block, "0/28", "14")}"
				}
			}
			vpc 				= ibm_is_vpc.testacc_vpc.id
		}
`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, acc.IsBareMetalServerProfileName, name, acc.IsBareMetalServerImage, acc.ISZoneName)
}
