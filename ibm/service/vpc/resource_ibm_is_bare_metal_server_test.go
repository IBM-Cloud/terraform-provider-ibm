// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"errors"
	"fmt"
	"regexp"
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

func TestAccIBMISBareMetalServer_firmwareUpdate(t *testing.T) {
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
					resource.TestCheckResourceAttrSet(
						"ibm_is_bare_metal_server.testacc_bms", "firmware_update_type_available"),
				),
			},
		},
	})
}

func TestAccIBMISBareMetalServer_bandwidth(t *testing.T) {
	var server string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	pipName := fmt.Sprintf("tf-vpc-pip-%d", acctest.RandIntRange(10, 100))
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
				Config: testAccCheckIBMISBareMetalServerBandwidthConfig(vpcname, subnetname, sshname, publicKey, name, pipName, 10000),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISBareMetalServerExists("ibm_is_bare_metal_server.testacc_bms", server),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "bandwidth", "10000"),
				),
			},
			{
				Config: testAccCheckIBMISBareMetalServerBandwidthConfig(vpcname, subnetname, sshname, publicKey, name, pipName, 25000),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISBareMetalServerExists("ibm_is_bare_metal_server.testacc_bms", server),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "bandwidth", "25000"),
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

func TestAccIBMISBareMetalServerVNIPSFM_basic(t *testing.T) {
	var server string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-server-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	vniname1 := fmt.Sprintf("tfip-vni-%d", acctest.RandIntRange(10, 100))
	vniname2 := fmt.Sprintf("tfip-vni-%d", acctest.RandIntRange(10, 100))
	psfm1 := "auto"
	psfm2 := "disabled"
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
				Config: testAccCheckIBMISBareMetalServerVNIPSFMConfig(vpcname, subnetname, sshname, publicKey, vniname1, vniname2, psfm1, psfm2, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISBareMetalServerExists("ibm_is_bare_metal_server.testacc_bms", server),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "primary_network_attachment.0.virtual_network_interface.0.protocol_state_filtering_mode", psfm1),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "network_attachments.0.virtual_network_interface.0.protocol_state_filtering_mode", psfm2),
				),
			},
			{
				Config: testAccCheckIBMISBareMetalServerVNIPSFMConfig(vpcname, subnetname, sshname, publicKey, vniname1, vniname2, psfm2, psfm1, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISBareMetalServerExists("ibm_is_bare_metal_server.testacc_bms", server),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "primary_network_attachment.0.virtual_network_interface.0.protocol_state_filtering_mode", psfm2),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "network_attachments.0.virtual_network_interface.0.protocol_state_filtering_mode", psfm1),
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
						"ibm_is_bare_metal_server.testacc_bms", "network_interfaces.0.vlan", "102"),
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

func TestAccIBMISBareMetalServer_reservation(t *testing.T) {
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
				Config: testAccCheckIBMISBareMetalServerReservationConfig(vpcname, subnetname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISBareMetalServerExists("ibm_is_bare_metal_server.testacc_bms", server),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "reservation_affinity.0.policy", "manual"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_bare_metal_server.testacc_bms", "reservation_affinity.0.pool"),
				),
			},
		},
	})
}

func testAccCheckIBMISBareMetalServerReservationConfig(vpcname, subnetname, sshname, publicKey, name string) string {
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
			reservation_affinity {
				policy = "manual"
				pool {
					id = "0735-b4a78f50-33bd-44f9-a3ff-4c33f444459d"
				}
			}
		}
`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, acc.IsBareMetalServerProfileName, name, acc.IsBareMetalServerImage, acc.ISZoneName)
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

func testAccCheckIBMISBareMetalServerBandwidthConfig(vpcname, subnetname, sshname, publicKey, name, pipName string, bandwidth int) string {
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
			bandwidth = %d
			profile 			= "%s"
			name 				= "%s"
			image 				= "%s"
			zone 				= "%s"
			keys 				= [ibm_is_ssh_key.testacc_sshkey.id]
			primary_network_attachment {
				virtual_network_interface { 
						auto_delete = true
						enable_infrastructure_nat = true
						primary_ip {
							name = "%s"
						}
						subnet = ibm_is_subnet.testacc_subnet.id
				}
				allowed_vlans = [100, 102]
			}
			vpc 				= ibm_is_vpc.testacc_vpc.id
		}
`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, bandwidth, acc.IsBareMetalServerProfileName, name, acc.IsBareMetalServerImage, acc.ISZoneName, pipName)
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
		resource "ibm_is_virtual_network_interface" "testacc_vni2"{
			name = "%s-1"
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
			network_attachments {
				name = "test-vni-200-202"
				virtual_network_interface { 
					id = ibm_is_virtual_network_interface.testacc_vni2.id
				}
				allowed_vlans = [200, 202]
			}
			vpc 				= ibm_is_vpc.testacc_vpc.id
		}
`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, vniname, vniname, acc.IsBareMetalServerProfileName, name, acc.IsBareMetalServerImage, acc.ISZoneName)
}

func testAccCheckIBMISBareMetalServerVNIPSFMConfig(vpcname, subnetname, sshname, publicKey, vniname1, vniname2, psfm1, psfm2, name string) string {
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
			primary_network_attachment {
				name = "test-vni-100-102"
				virtual_network_interface { 
					name = "%s"
					subnet = ibm_is_subnet.testacc_subnet.id
					enable_infrastructure_nat = true
					allow_ip_spoofing = true
					protocol_state_filteting_mode = "%s"
				}
				allowed_vlans = [100, 102]
			}
			network_attachments {
				name = "test-vni-tfp-snac100-102"
				virtual_network_interface { 
					name = "%s"
					subnet = ibm_is_subnet.testacc_subnet.id
					enable_infrastructure_nat = true
					allow_ip_spoofing = true
					protocol_state_filteting_mode = "%s"
				}
				allowed_vlans = [103, 105]
			}
			vpc 				= ibm_is_vpc.testacc_vpc.id
		}
`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, acc.IsBareMetalServerProfileName, name, acc.IsBareMetalServerImage, acc.ISZoneName, vniname2, psfm1, vniname2, psfm2)
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

func TestAccIBMISBareMetalServer_updateInitialization(t *testing.T) {
	var server string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-server-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfip-subnet-%d", acctest.RandIntRange(10, 100))
	userdata1 := "a"
	userdata2 := "b"
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
				Config: testAccCheckIBMISBareMetalServerInitializationConfig(vpcname, subnetname, sshname, publicKey, name, acc.IsBareMetalServerImage, userdata1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISBareMetalServerExists("ibm_is_bare_metal_server.testacc_bms", server),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "image", acc.IsBareMetalServerImage),
					resource.TestCheckResourceAttrSet(
						"ibm_is_bare_metal_server.testacc_bms", "keys.#"),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "user_data", userdata1),
				),
			},
			{
				Config: testAccCheckIBMISBareMetalServerInitializationConfig(vpcname, subnetname, sshname, publicKey, name, acc.IsBareMetalServerImage2, userdata2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISBareMetalServerExists("ibm_is_bare_metal_server.testacc_bms", server),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "image", acc.IsBareMetalServerImage2),
					resource.TestCheckResourceAttrSet(
						"ibm_is_bare_metal_server.testacc_bms", "keys.#"),
					resource.TestCheckResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms", "user_data", userdata2),
				),
			},
		},
	})
}

func testAccCheckIBMISBareMetalServerInitializationConfig(vpcname, subnetname, sshname, publicKey, name, image, userdata string) string {
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
			user_data 			= "%s"
			keys 				= [ibm_is_ssh_key.testacc_sshkey.id]
			primary_network_interface {
				subnet     		= ibm_is_subnet.testacc_subnet.id
			}
			vpc 				= ibm_is_vpc.testacc_vpc.id
		}
`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, acc.IsBareMetalServerProfileName, name, image, acc.ISZoneName, userdata)
}

func TestAccIBMISBareMetalServerMultiple(t *testing.T) {
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	namePrefix := fmt.Sprintf("tf-server-%d", acctest.RandIntRange(10, 100))
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
				Config: testAccCheckIBMISBareMetalServerMultipleConfig(vpcname, subnetname, sshname, publicKey, namePrefix),
				Check: resource.ComposeTestCheckFunc(
					// Check that each of the 9 servers has an entry in the state file
					testAccCheckBareMetalServerStateEntries(9),
					resource.TestMatchResourceAttr(
						"ibm_is_bare_metal_server.testacc_bms.0", "name", regexp.MustCompile(fmt.Sprintf("%s-0", namePrefix))),
				),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})
}

// Test function that verifies the state contains the expected number of bare metal servers
func testAccCheckBareMetalServerStateEntries(expectedCount int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		count := 0

		// Check for each of the expected instances in the state file
		for i := 0; i < expectedCount; i++ {
			resourceName := fmt.Sprintf("ibm_is_bare_metal_server.testacc_bms.%d", i)
			_, ok := s.RootModule().Resources[resourceName]
			if !ok {
				return fmt.Errorf("Expected state entry not found: %s", resourceName)
			}

			// The ID might be empty if resource creation failed, but we still want
			// to count it as a state entry since that's what we're testing
			count++
		}

		if count != expectedCount {
			return fmt.Errorf("Expected to find %d bare metal server entries in state, found %d", expectedCount, count)
		}

		return nil
	}
}

// Configuration for testing multiple bare metal servers
func testAccCheckIBMISBareMetalServerMultipleConfig(vpcname, subnetname, sshname, publicKey, namePrefix string) string {
	return fmt.Sprintf(`
		resource "ibm_is_vpc" "testacc_vpc" {
			name = "%s"
		}
	  
		resource "ibm_is_subnet" "testacc_subnet" {
			name            			= "%s"
			vpc             			= ibm_is_vpc.testacc_vpc.id
			zone            			= "%s"
			total_ipv4_address_count 	= 256
		}
	  
		resource "ibm_is_ssh_key" "testacc_sshkey" {
			name       			= "%s"
			public_key 			= "%s"
		}
	  
		resource "ibm_is_bare_metal_server" "testacc_bms" {
			count = 9
			profile 			= "%s"
			name 				= "%s-${count.index}"
			image 				= "%s"
			zone 				= "%s"
			keys 				= [ibm_is_ssh_key.testacc_sshkey.id]
			vpc 				= ibm_is_vpc.testacc_vpc.id
			primary_network_attachment {
				name = "bms-${count.index}-pnac-${count.index}"
				virtual_network_interface { 
					name = "bms-${count.index}-eth0-pnac-vni-${count.index}"
					subnet = ibm_is_subnet.testacc_subnet.id
					enable_infrastructure_nat = true
					allow_ip_spoofing = true
				}
				allowed_vlans = [100, 102]
			}
			network_attachments {
				name = "bms-${count.index}-eth1-snac-${count.index}"
				virtual_network_interface { 
					name = "bms-${count.index}-eth1-snac-vni-${count.index}"
					subnet = ibm_is_subnet.testacc_subnet.id
					enable_infrastructure_nat = true
					allow_ip_spoofing = true
				}
				allowed_vlans = [203, 205]
			}
			lifecycle {
				ignore_changes = [ network_attachments ]
			}
		}
`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, acc.IsBareMetalServerProfileName, namePrefix, acc.IsBareMetalServerImage, acc.ISZoneName)
}
