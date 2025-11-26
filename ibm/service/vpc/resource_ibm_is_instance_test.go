// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMISInstance_basic(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	userData1 := "a"
	userData2 := "b"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceConfig(vpcname, subnetname, sshname, publicKey, name, userData1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData1),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.0.manufacturer"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceConfig(vpcname, subnetname, sshname, publicKey, name, userData2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData2),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "primary_network_interface.0.port_speed"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.0.manufacturer"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "numa_count"),
				),
			},
		},
	})
}
func TestAccIBMISInstance_sdpbasic(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	userData1 := "a"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceSdpConfig(vpcname, subnetname, sshname, publicKey, name, userData1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData1),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "boot_volume.0.size", "250"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "boot_volume.0.iops", "10000"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "boot_volume.0.profile", "sdp"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.0.manufacturer"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceSdpCapacityUpdateConfig(vpcname, subnetname, sshname, publicKey, name, userData1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData1),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "boot_volume.0.size", "25000"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "boot_volume.0.iops", "10000"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "boot_volume.0.profile", "sdp"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "primary_network_interface.0.port_speed"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.0.manufacturer"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "numa_count"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceSdpIopsUpdateConfig(vpcname, subnetname, sshname, publicKey, name, userData1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData1),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "boot_volume.0.size", "25000"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "boot_volume.0.iops", "28000"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "boot_volume.0.profile", "sdp"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "primary_network_interface.0.port_speed"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.0.manufacturer"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "numa_count"),
				),
			},
		},
	})
}

func TestAccIBMISInstanceWithoutKeys_basic(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	userData1 := "a"
	userData2 := "b"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceWithoutKeysConfig(vpcname, subnetname, name, userData1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData1),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.0.manufacturer"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceWithoutKeysConfig(vpcname, subnetname, name, userData2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData2),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "primary_network_interface.0.port_speed"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.0.manufacturer"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "numa_count"),
				),
			},
		},
	})
}

func TestAccIBMISInstance_concom(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	userData1 := "a"
	userData2 := "b"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceConComConfig(vpcname, subnetname, sshname, publicKey, name, userData1, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData1),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.0.manufacturer"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "confidential_compute_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "enable_secure_boot"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "enable_secure_boot", "true"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceConComConfig(vpcname, subnetname, sshname, publicKey, name, userData2, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData2),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "primary_network_interface.0.port_speed"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.0.manufacturer"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "numa_count"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "confidential_compute_mode"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "enable_secure_boot"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "enable_secure_boot", "true"),
				),
			},
		},
	})
}
func TestAccIBMISInstance_vni(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	vniname := fmt.Sprintf("tf-vni-%d", acctest.RandIntRange(10, 100))
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	userData1 := "a"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceVniConfig(vpcname, subnetname, sshname, publicKey, name, vniname, userData1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData1),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.0.manufacturer"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "primary_network_attachment.#"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "primary_network_attachment.#", "1"),
				),
			},
		},
	})
}
func TestAccIBMISInstance_vni_update(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	vniname := fmt.Sprintf("tf-vni-%d", acctest.RandIntRange(10, 100))
	inlinevniname := fmt.Sprintf("tf-inlinevni-%d", acctest.RandIntRange(10, 100))
	inlinevniupdatedname := fmt.Sprintf("tf-inlinevniupd-%d", acctest.RandIntRange(10, 100))
	pnaName := fmt.Sprintf("tf-pna-%d", acctest.RandIntRange(10, 100))
	snaName := fmt.Sprintf("tf-sna-%d", acctest.RandIntRange(10, 100))
	snaNameUpdated := fmt.Sprintf("tf-sna-upd-%d", acctest.RandIntRange(10, 100))
	pnaNameUpdated := fmt.Sprintf("tf-pna-upd-%d", acctest.RandIntRange(10, 100))
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	userData1 := "a"
	protocolStateFilteringMode := "auto"
	protocolStateFilteringUpdated := "enabled"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceVniUpdateConfig(vpcname, subnetname, sshname, publicKey, name, vniname, inlinevniname, pnaName, snaName, userData1, protocolStateFilteringMode),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData1),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.0.manufacturer"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "primary_network_attachment.#"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "primary_network_attachment.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "primary_network_attachment.0.name", pnaName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "network_attachments.0.name", snaName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "network_attachments.0.virtual_network_interface.0.name", inlinevniname),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "network_attachments.0.virtual_network_interface.0.protocol_state_filtering_mode", protocolStateFilteringMode),
				),
			},
			{
				Config: testAccCheckIBMISInstanceVniUpdateConfig(vpcname, subnetname, sshname, publicKey, name, vniname, inlinevniupdatedname, pnaNameUpdated, snaNameUpdated, userData1, protocolStateFilteringUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData1),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.0.manufacturer"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "primary_network_attachment.#"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "primary_network_attachment.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "primary_network_attachment.0.name", pnaNameUpdated),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "network_attachments.0.name", snaNameUpdated),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "network_attachments.0.virtual_network_interface.0.name", inlinevniupdatedname),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "network_attachments.0.virtual_network_interface.0.protocol_state_filtering_mode", protocolStateFilteringUpdated),
				),
			},
		},
	})
}

func TestAccIBMISInstance_enc_catalog(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	resourceName := fmt.Sprintf("tf-cosresource-%d", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("tf-key-%d", acctest.RandIntRange(10, 100))
	userData1 := "a"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceCatEncryptionConfig(vpcname, subnetname, sshname, publicKey, name, userData1, resourceName, keyName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData1),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.0.manufacturer"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "boot_volume.0.encryption"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "catalog_offering.0.version_crn"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "catalog_offering.0.plan_crn"),
				),
			},
		},
	})
}

func TestAccIBMISInstance_lifecycle(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	userData1 := "a"
	userData2 := "b"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceConfig(vpcname, subnetname, sshname, publicKey, name, userData1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData1),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "lifecycle_state", "stable"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "lifecycle_reasons.#", "0"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceConfig(vpcname, subnetname, sshname, publicKey, name, userData2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData2),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "primary_network_interface.0.port_speed"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "lifecycle_state", "stable"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "lifecycle_reasons.#", "0"),
				),
			},
		},
	})
}
func TestAccIBMISInstance_rip(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	subnetripname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	userData1 := "a"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceRipConfig(vpcname, subnetname, subnetripname, sshname, publicKey, name, userData1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData1),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
				),
			},
		},
	})
}

func TestAccIBMISInstance_ResizeBoot(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	userData1 := "a"
	resize1 := int64(220)
	resize2 := int64(250)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceResizeConfig(vpcname, subnetname, sshname, publicKey, name, userData1, resize1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData1),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "boot_volume.0.size", fmt.Sprintf("%d", resize1)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
				),
			},
			{
				Config: testAccCheckIBMISInstanceResizeConfig(vpcname, subnetname, sshname, publicKey, name, userData1, resize2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData1),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "boot_volume.0.size", fmt.Sprintf("%d", resize2)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
				),
			},
		},
	})
}
func TestAccIBMISInstance_RenameBoot(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	userData1 := "a"
	rename1 := fmt.Sprintf("tf-bootvol-%d", acctest.RandIntRange(10, 100))
	rename2 := fmt.Sprintf("tf-bootvol-update-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceRenameConfig(vpcname, subnetname, sshname, publicKey, name, userData1, rename1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData1),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "boot_volume.0.name", rename1),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
				),
			},
			{
				Config: testAccCheckIBMISInstanceRenameConfig(vpcname, subnetname, sshname, publicKey, name, userData1, rename2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData1),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "boot_volume.0.name", rename2),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
				),
			},
		},
	})
}

func TestAccIBMISInstance_bootVolumeUserTags(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	userData1 := "a"
	userTags1 := "tags-0"
	userTags2 := "tags-1"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceUserTagConfig(vpcname, subnetname, sshname, publicKey, name, userData1, userTags1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "boot_volume.0.tags.0", userTags1),
				),
			},
			{
				Config: testAccCheckIBMISInstanceUserTagConfig(vpcname, subnetname, sshname, publicKey, name, userData1, userTags2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "boot_volume.0.tags.0", userTags2),
				),
			},
		},
	})
}

func TestAccIBMISInstance_withAvailablePolicy(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	templateName := fmt.Sprintf("tf-instnace-template-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceConfigWithAvailablePolicyHostFailure_Default(vpcname, subnetname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr("ibm_is_instance.testacc_instance", "availability_policy_host_failure", "restart"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceConfigWithAvailablePolicyHostFailure_Updated(vpcname, subnetname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr("ibm_is_instance.testacc_instance", "availability_policy_host_failure", "stop"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceConfigWithAvailablePolicyHostFailure_WithTemplate(vpcname, subnetname, sshname, publicKey, templateName, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr("ibm_is_instance.testacc_instance", "availability_policy_host_failure", "stop"),
				),
			},
		},
	})
}
func TestAccIBMISInstanceBandwidth_basic(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	totalVolumeBandwidth := 1000
	totalVolumeBandwidthUpdated := 1000
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceBandwidthConfig(vpcname, subnetname, sshname, publicKey, name, totalVolumeBandwidth),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "total_volume_bandwidth", strconv.Itoa(totalVolumeBandwidth)),
				),
			},
			{
				Config: testAccCheckIBMISInstanceBandwidthUpdateConfig(vpcname, subnetname, sshname, publicKey, name, totalVolumeBandwidthUpdated),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "total_volume_bandwidth", strconv.Itoa(totalVolumeBandwidthUpdated)),
				),
			},
		},
	})
}

func TestAccIBMISInstance_QoS(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	totalVolumeBandwidth := 1000
	qosMode := "weighted"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceQOSConfig(vpcname, subnetname, sshname, publicKey, name, qosMode, totalVolumeBandwidth),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "total_volume_bandwidth", strconv.Itoa(totalVolumeBandwidth)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "volume_bandwidth_qos_mode", qosMode),
				),
			},
		},
	})
}

func TestAccIBMISInstanceWithSecurityGroup_basic(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	secGrpName := fmt.Sprintf("tf-secgrp-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceWithSecurityGroupConfig(vpcname, subnetname, sshname, publicKey, secGrpName, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "primary_network_interface.0.security_groups.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "network_interfaces.0.security_groups.#"),
				),
			},
		},
	})
}

func TestAccIBMISInstance_action(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	userData := "a"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceConfig(vpcname, subnetname, sshname, publicKey, name, userData),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
				),
			},
			{
				Config: testAccCheckIBMISInstanceConfigActionStop(vpcname, subnetname, sshname, publicKey, name, userData),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "status", "stopped"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceConfigActionStart(vpcname, subnetname, sshname, publicKey, name, userData),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "status", "running"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceConfigActionReboot(vpcname, subnetname, sshname, publicKey, name, userData),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "status", "running"),
				),
			},
		},
	})
}

func TestAccIBMISInstance_metadata_service(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDVtuCfWKVGKaRmaRG6JQZY8YdxnDgGzVOK93IrV9R5Hl0JP1oiLLWlZQS2reAKb8lBqyDVEREpaoRUDjqDqXG8J/kR42FKN51su914pjSBc86wJ02VtT1Wm1zRbSg67kT+g8/T1jCgB5XBODqbcICHVP8Z1lXkgbiHLwlUrbz6OZkGJHo/M/kD1Eme8lctceIYNz/Ilm7ewMXZA4fsidpto9AjyarrJLufrOBl4MRVcZTDSJ7rLP982aHpu9pi5eJAjOZc7Og7n4ns3NFppiCwgVMCVUQbN5GBlWhZ1OsT84ZiTf+Zy8ew+Yg5T7Il8HuC7loWnz+esQPf0s3xhC/kTsGgZreIDoh/rxJfD67wKXetNSh5RH/n5BqjaOuXPFeNXmMhKlhj9nJ8scayx/wsvOGuocEIkbyJSLj3sLUU403OafgatEdnJOwbqg6rUNNF5RIjpJpL7eEWlKIi1j9LyhmPJ+fEO7TmOES82VpCMHpLbe4gf/MhhJ/Xy8DKh9s= root@ffd8363b1226
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceWithMetaConfigDefault(vpcname, subnetname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "metadata_service.0.enabled", "false"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "metadata_service.0.protocol", "http"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "metadata_service.0.response_hop_limit", "1"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceWithMetaConfig(vpcname, subnetname, sshname, publicKey, name, true, "https", 5),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "metadata_service_enabled", "true"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "metadata_service.0.protocol", "https"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "metadata_service.0.response_hop_limit", "5"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceWithMetaConfig(vpcname, subnetname, sshname, publicKey, name, true, "http", 10),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "metadata_service_enabled", "true"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "metadata_service.0.protocol", "http"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "metadata_service.0.response_hop_limit", "10"),
				),
			},
		},
	})
}
func TestAccIBMISInstance_profile(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceConfigWithProfile(vpcname, subnetname, sshname, publicKey, name, acc.InstanceProfileName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "profile", acc.InstanceProfileName),
				),
			},

			{
				Config: testAccCheckIBMISInstanceConfigWithProfile(vpcname, subnetname, sshname, publicKey, name, acc.InstanceProfileNameUpdate),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "profile", acc.InstanceProfileNameUpdate),
				),
			},
		},
	})
}

func TestAccIBMISInstance_basicwithipv4(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	ipv4address := acc.ISIPV4Address

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceConfigwithipv4(vpcname, subnetname, sshname, publicKey, name, ipv4address),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "primary_network_interface.0.primary_ipv4_address", ipv4address),
				),
			},
		},
	})
}

func TestAccIBMISInstance_VolumeAutoDelete(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	volname := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceVolumeAutoDelete(vpcname, subnetname, sshname, publicKey, volname, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "volumes.#", "1"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "auto_delete_volume", "true"),
				),
			},
		},
	})
}

func TestAccIBMISInstance_Volume(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	volname := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceVolume(vpcname, subnetname, sshname, publicKey, volname, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "volumes.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceVolumeUpdate(vpcname, subnetname, sshname, publicKey, volname, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "volumes.#", "0"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceVolume(vpcname, subnetname, sshname, publicKey, volname, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "volumes.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMISInstance_Disk(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	volname := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceDisk(vpcname, subnetname, sshname, publicKey, volname, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "disks.#", "1"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "disks.0.name"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "disks.0.size"),
				),
			},
		},
	})
}

func TestAccIBMISInstance_Placement(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	dhostname := fmt.Sprintf("tf-dhost-%d", acctest.RandIntRange(10, 100))
	dhostgrpname := fmt.Sprintf("tf-dhostgrp-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	volname := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstancePlacement(vpcname, subnetname, sshname, publicKey, volname, name, dhostname, dhostgrpname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_dedicated_host.dhost", "instances.#"),
					resource.TestCheckResourceAttr(
						"data.ibm_is_dedicated_host.dhost", "instances.0.name", name),
				),
			},
		},
	})
}

func TestAccIBMISInstance_ByVolume(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tf-instnace1-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	volname := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	sname := fmt.Sprintf("tfsnapshot-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceByVolume(vpcname, subnetname, sshname, publicKey, volname, name, name1, sname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.instancebyvol", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.instancebyvol", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.instancebyvol", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.instancebyvol", "boot_volume.0.volume_id"),
				),
			},
		},
	})
}

func TestAccIBMISInstanceSnapshotRestore_basic(t *testing.T) {
	var instance, instanceRestore string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	snapshot := fmt.Sprintf("tf-snapshot-%d", acctest.RandIntRange(10, 100))
	vsiRestore := fmt.Sprintf("tf-instancerestore-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceSnapshotRestoreConfig(vpcname, subnetname, sshname, publicKey, name, snapshot, vsiRestore),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance_restore", instanceRestore),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_snapshot.testacc_snapshot", "name", snapshot),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance_restore", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance_restore", "name", vsiRestore),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance_restore", "boot_volume.0.name", "boot-restore"),
				),
			},
		},
	})
}
func TestAccIBMISInstanceSnapshotRestore_crn(t *testing.T) {
	var instance, instanceRestore string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	snapshot := fmt.Sprintf("tf-snapshot-%d", acctest.RandIntRange(10, 100))
	vsiRestore := fmt.Sprintf("tf-instancerestore-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceSnapshotRestoreCrnConfig(vpcname, subnetname, sshname, publicKey, name, snapshot, vsiRestore),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance_restore", instanceRestore),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_snapshot.testacc_snapshot", "name", snapshot),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance_restore", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance_restore", "name", vsiRestore),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance_restore", "boot_volume.0.name", "boot-restore"),
				),
			},
		},
	})
}

func TestAccIBMISInstanceSnapshotRestore_forcenew(t *testing.T) {
	var instance, instanceRestore string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	name2 := fmt.Sprintf("tf-instnace2-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	vsiRestore := fmt.Sprintf("tf-instancerestore-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceSnapshotRestoreForceNewConfig(vpcname, subnetname, sshname, publicKey, name, name2, name, vsiRestore),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance_restore", instanceRestore),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_snapshot.testacc_snapshot", "name", fmt.Sprintf("%s%s", name, "-snapshot")),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance_restore", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance_restore", "name", vsiRestore),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance_restore", "boot_volume.0.name", "boot-restore"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceSnapshotRestoreForceNewConfig(vpcname, subnetname, sshname, publicKey, name, name2, name2, vsiRestore),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance_restore", instanceRestore),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_snapshot.testacc_snapshot", "name", fmt.Sprintf("%s%s", name2, "-snapshot")),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance_restore", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance_restore", "name", vsiRestore),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance_restore", "boot_volume.0.name", "boot-restore"),
				),
			},
		},
	})
}

func TestAccIBMISInstance_Reservation(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instance-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceReservation(vpcname, subnetname, name, publicKey, sshname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName3),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "reservation_affinity.0.policy", "manual"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "reservation_affinity.0.pool"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceDestroy(s *terraform.State) error {

	instanceC, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_instance" {
			continue
		}
		getinsOptions := &vpcv1.GetInstanceOptions{
			ID: &rs.Primary.ID,
		}
		_, _, err := instanceC.GetInstance(getinsOptions)

		if err == nil {
			return fmt.Errorf("instance still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISInstanceWithSecurityGroupConfig(vpcname, subnetname, sshname, publicKey, secgrpname, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  resource "ibm_is_security_group" "testacc_security_group" {
		name = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id
	  }
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		  security_groups = [ibm_is_security_group.testacc_security_group.id]
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		  security_groups = [ibm_is_security_group.testacc_security_group.id]
		}
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, secgrpname, name, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName)
}

func testAccCheckIBMISInstanceExists(n string, instance string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		instanceC, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getinsOptions := &vpcv1.GetInstanceOptions{
			ID: &rs.Primary.ID,
		}
		foundins, _, err := instanceC.GetInstance(getinsOptions)
		if err != nil {
			return err
		}
		instance = *foundins.ID
		return nil
	}
}

func testAccCheckIBMISInstanceWithMetaConfigDefault(vpcname, subnetname, sshname, publicKey, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		}
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName)
}

func testAccCheckIBMISInstanceWithMetaConfig(vpcname, subnetname, sshname, publicKey, name string, metadata_service_enabled bool, protocol string, hop_limit int) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		}
		metadata_service {
			enabled = %t
			protocol = "%s"
			response_hop_limit = %d
		  }
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName, metadata_service_enabled, protocol, hop_limit)
}

func testAccCheckIBMISInstanceConfig(vpcname, subnetname, sshname, publicKey, name, userData string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
  primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
  }
		user_data = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
  network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
    name   = "eth1"
  }
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, userData, acc.ISZoneName)
}
func testAccCheckIBMISInstanceSdpConfig(vpcname, subnetname, sshname, publicKey, name, userData string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		boot_volume {
			size 	= 250 
			profile = "sdp"
			iops	= 10000
		}
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		user_data = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, userData, acc.ISZoneName)
}
func testAccCheckIBMISInstanceSdpIopsUpdateConfig(vpcname, subnetname, sshname, publicKey, name, userData string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		boot_volume {
			size 	= 25000 
			profile = "sdp"
			iops	= 28000
		}
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		user_data = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, userData, acc.ISZoneName)
}
func testAccCheckIBMISInstanceSdpCapacityUpdateConfig(vpcname, subnetname, sshname, publicKey, name, userData string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		boot_volume {
			size 	= 25000 
			profile = "sdp"
			iops	= 10000
		}
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		user_data = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, userData, acc.ISZoneName)
}

func testAccCheckIBMISInstanceWithoutKeysConfig(vpcname, subnetname, name, userData string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		user_data = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		}
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name, acc.IsImage, acc.InstanceProfileName, userData, acc.ISZoneName)
}

func testAccCheckIBMISInstanceConComConfig(vpcname, subnetname, sshname, publicKey, name, userData string, esb bool) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		enable_secure_boot = %t
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		user_data = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		}
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, esb, userData, acc.ISZoneName)
}
func testAccCheckIBMISInstanceVniConfig(vpcname, subnetname, sshname, publicKey, name, vniname, userData string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}

resource "ibm_is_subnet" "testacc_subnet" {
	name            = "%s"
	vpc             = ibm_is_vpc.testacc_vpc.id
	zone            = "%s"
	ipv4_cidr_block = "%s"
}

resource "ibm_is_ssh_key" "testacc_sshkey" {
	name       = "%s"
	public_key = "%s"
}

resource "ibm_is_virtual_network_interface" "testacc_vni"{
	name = "%s"
	allow_ip_spoofing = true
	subnet = ibm_is_subnet.testacc_subnet.id
} 

resource "ibm_is_instance" "testacc_instance" {
	name    = "%s"
	image   = "%s"
	profile = "%s"
    primary_network_attachment {
        name = "test-vni"
        virtual_network_interface { 
            id = ibm_is_virtual_network_interface.testacc_vni.id
        }
    }
	user_data = "%s"
	vpc  = ibm_is_vpc.testacc_vpc.id
	zone = "%s"
	keys = [ibm_is_ssh_key.testacc_sshkey.id]
}`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, vniname, name, acc.IsImage, acc.InstanceProfileName, userData, acc.ISZoneName)
}
func testAccCheckIBMISInstanceVniUpdateConfig(vpcname, subnetname, sshname, publicKey, name, vniname, inlinevniname, pnaName, snaName, userData, protocolStateFilteringMode string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}

resource "ibm_is_subnet" "testacc_subnet" {
	name            = "%s"
	vpc             = ibm_is_vpc.testacc_vpc.id
	zone            = "%s"
	ipv4_cidr_block = "%s"
}

resource "ibm_is_ssh_key" "testacc_sshkey" {
	name       = "%s"
	public_key = "%s"
}

resource "ibm_is_virtual_network_interface" "testacc_vni"{
	name = "%s"
	allow_ip_spoofing = true
	subnet = ibm_is_subnet.testacc_subnet.id
} 

resource "ibm_is_instance" "testacc_instance" {
	name    = "%s"
	image   = "%s"
	profile = "%s"
    primary_network_attachment {
        name = "%s"
        virtual_network_interface { 
            id = ibm_is_virtual_network_interface.testacc_vni.id
			protocol_state_filtering_mode = "%s"
        }
    }
	network_attachments {
		name = "%s"
		virtual_network_interface { 
            name = "%s"
			primary_ip {
				auto_delete 	= true
				address 		= cidrhost(ibm_is_subnet.testacc_subnet.ipv4_cidr_block, 23)
			}
			subnet = ibm_is_subnet.testacc_subnet.id
			protocol_state_filtering_mode = "%s"
        }
	}
	user_data = "%s"
	vpc  = ibm_is_vpc.testacc_vpc.id
	zone = "%s"
	keys = [ibm_is_ssh_key.testacc_sshkey.id]
}`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, vniname, name, acc.IsImage, acc.InstanceProfileName, pnaName, protocolStateFilteringMode, snaName, inlinevniname, protocolStateFilteringMode, userData, acc.ISZoneName)
}

func testAccCheckIBMISInstanceCatEncryptionConfig(vpcname, subnetname, sshname, publicKey, name, userData, resourceName, keyName string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  data ibm_is_images testacc_images {
		  catalog_managed = true
		}
	  resource "ibm_resource_instance" "testacc_resource" {
	  	name              = "%s"
	  	service           = "kms"
	  	plan              = "tiered-pricing"
	  	location          = "%s"
	  	}
	  resource "ibm_kms_key" "testacc_key" {
	  	instance_id = "${ibm_resource_instance.testacc_resource.guid}"
	  	key_name =  "%s"
	  	standard_key =  false
	  	force_delete = true
	  }
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		catalog_offering {
			version_crn = "crn:v1:staging:public:globalcatalog-collection:global::1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc:version:4f8466eb-2218-42e3-a755-bf352b559c69-global/6a73aa69-5dd9-4243-a908-3b62f467cbf8-global"
			plan_crn = "crn:v1:staging:public:globalcatalog-collection:global::1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc:plan:sw.1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.279a3cee-ba7d-42d5-ae88-6a0ebc56fa4a-global"
		}
		boot_volume {
			encryption = ibm_kms_key.testacc_key.crn
		}
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		user_data = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		}
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, resourceName, acc.RegionName, keyName, name, acc.InstanceProfileName, userData, acc.ISZoneName)
}

func testAccCheckIBMISInstanceRipConfig(vpcname, subnetname, subnetripname, sshname, publicKey, name, userData string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }

	  resource "ibm_is_subnet_reserved_ip" "testacc_rip" {
		subnet = ibm_is_subnet.testacc_subnet.id
		name = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		  primary_ip {
			reserved_ip = ibm_is_subnet_reserved_ip.testacc_rip.reserved_ip
		  }
		}
		user_data = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, subnetripname, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, userData, acc.ISZoneName)
}

func testAccCheckIBMISInstanceResizeConfig(vpcname, subnetname, sshname, publicKey, name, userData string, resize int64) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		boot_volume {
			size = %d
		}
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		user_data = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, resize, userData, acc.ISZoneName)
}
func testAccCheckIBMISInstanceRenameConfig(vpcname, subnetname, sshname, publicKey, name, userData, rename string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		boot_volume {
			name = "%s"
		}
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		user_data = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
	  }
	  `, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, rename, userData, acc.ISZoneName)
}

func testAccCheckIBMISInstanceQOSConfig(vpcname, subnetname, sshname, publicKey, name, qosMode string, bandwidth int) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
  primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
  }
		total_volume_bandwidth = %d
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
  network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
    name   = "eth1"
  }
		volume_bandwidth_qos_mode = "%s"  
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, bandwidth, acc.ISZoneName, qosMode)
}

func testAccCheckIBMISInstanceBandwidthConfig(vpcname, subnetname, sshname, publicKey, name string, bandwidth int) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		total_volume_bandwidth = %d
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		}
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, bandwidth, acc.ISZoneName)
}

func testAccCheckIBMISInstanceBandwidthUpdateConfig(vpcname, subnetname, sshname, publicKey, name string, bandwidth int) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		total_volume_bandwidth = %d
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		}
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, bandwidth, acc.ISZoneName)
}

func testAccCheckIBMISInstanceConfigActionStop(vpcname, subnetname, sshname, publicKey, name, userData string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		user_data = "%s"
		action = "stop"
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		}
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, userData, acc.ISZoneName)
}

func testAccCheckIBMISInstanceConfigActionStart(vpcname, subnetname, sshname, publicKey, name, userData string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		user_data = "%s"
		action = "start"
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		}
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, userData, acc.ISZoneName)
}

func testAccCheckIBMISInstanceConfigActionReboot(vpcname, subnetname, sshname, publicKey, name, userData string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		user_data = "%s"
		action = "reboot"
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		}
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, userData, acc.ISZoneName)
}

func testAccCheckIBMISInstanceSnapshotRestoreConfig(vpcname, subnetname, sshname, publicKey, name, snapshot, insRestore string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		total_ipv4_address_count 	= 16
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		}
	  }
	  resource "ibm_is_snapshot" "testacc_snapshot" {
		name 			= "%s"
		source_volume 	= ibm_is_instance.testacc_instance.volume_attachments[0].volume_id
	  }

	  resource "ibm_is_instance" "testacc_instance_restore" {
		name    = "%s"
		profile = "%s"
		boot_volume {
			name     = "boot-restore"
			snapshot = ibm_is_snapshot.testacc_snapshot.id
		}
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		}
	  }
	  `, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName, snapshot, insRestore, acc.InstanceProfileName, acc.ISZoneName)
}
func testAccCheckIBMISInstanceSnapshotRestoreCrnConfig(vpcname, subnetname, sshname, publicKey, name, snapshot, insRestore string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		total_ipv4_address_count 	= 16
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		}
	  }
	  resource "ibm_is_snapshot" "testacc_snapshot" {
		name 			= "%s"
		source_volume 	= ibm_is_instance.testacc_instance.volume_attachments[0].volume_id
	  }

	  resource "ibm_is_instance" "testacc_instance_restore" {
		name    = "%s"
		profile = "%s"
		boot_volume {
			name     		= "boot-restore"
			snapshot_crn 	= ibm_is_snapshot.testacc_snapshot.crn
		}
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		}
	  }
	  `, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName, snapshot, insRestore, acc.InstanceProfileName, acc.ISZoneName)
}

func testAccCheckIBMISInstanceSnapshotRestoreForceNewConfig(vpcname, subnetname, sshname, publicKey, name, name2, name3, insRestore string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		total_ipv4_address_count 	= 16
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		}
	  }
	  resource "ibm_is_instance" "testacc_instance1" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		}
	  }
	  data "ibm_is_instance" "testacc_ins_data" {
		  name = "%s"
		  depends_on = [
			ibm_is_instance.testacc_instance,
			ibm_is_instance.testacc_instance1,
		  ]
	  }
	  resource "ibm_is_snapshot" "testacc_snapshot" {
		name 			= "${data.ibm_is_instance.testacc_ins_data.name}-snapshot"
		source_volume 	= data.ibm_is_instance.testacc_ins_data.boot_volume[0].volume_id
	  }

	  resource "ibm_is_instance" "testacc_instance_restore" {
		name    = "%s"
		profile = "%s"
		boot_volume {
			name     = "boot-restore"
			snapshot = ibm_is_snapshot.testacc_snapshot.id
		}
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		}
	  }
	  `, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName, name2, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName, name3, insRestore, acc.InstanceProfileName, acc.ISZoneName)
}

func testAccCheckIBMISInstanceConfigWithProfile(vpcname, subnetname, sshname, publicKey, name, isInstanceProfileName string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		}
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, isInstanceProfileName, acc.ISZoneName)
}

func testAccCheckIBMISInstanceConfigwithipv4(vpcname, subnetname, sshname, publicKey, name, ipv4address string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		  primary_ipv4_address = "%s"
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		}
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, ipv4address, acc.ISZoneName)
}

func testAccCheckIBMISInstanceVolume(vpcname, subnetname, sshname, publicKey, volName, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_volume" "storage" {
		name    = "%s"
		profile = "10iops-tier"
		zone    = "%s"
		# capacity= 200
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet = ibm_is_subnet.testacc_subnet.id
		}
		vpc     = ibm_is_vpc.testacc_vpc.id
		zone    = "%s"
		keys    = [ibm_is_ssh_key.testacc_sshkey.id]
		volumes = [ibm_is_volume.storage.id]
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, volName, acc.ISZoneName, name, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName)
}

func testAccCheckIBMISInstanceDisk(vpcname, subnetname, sshname, publicKey, volName, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_volume" "storage" {
		name    = "%s"
		profile = "10iops-tier"
		zone    = "%s"
		# capacity= 200
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet = ibm_is_subnet.testacc_subnet.id
		}
		vpc     = ibm_is_vpc.testacc_vpc.id
		zone    = "%s"
		keys    = [ibm_is_ssh_key.testacc_sshkey.id]
		volumes = [ibm_is_volume.storage.id]
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, volName, acc.ISZoneName, name, acc.IsImage, acc.InstanceDiskProfileName, acc.ISZoneName)
}

func testAccCheckIBMISInstancePlacement(vpcname, subnetname, sshname, publicKey, volName, name, dhostname, dhostgrpname string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_volume" "storage" {
		name    = "%s"
		profile = "10iops-tier"
		zone    = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet = ibm_is_subnet.testacc_subnet.id
		}
		dedicated_host_group = "%s"
		vpc     = ibm_is_vpc.testacc_vpc.id
		zone    = "%s"
		keys    = [ibm_is_ssh_key.testacc_sshkey.id]
		
	  }
	  data "ibm_is_dedicated_host" "dhost"{
		  host_group = ibm_is_instance.testacc_instance.dedicated_host_group
		  name 		 = "%s"
	  } 
	 
	  `, vpcname, subnetname, acc.ISZoneName3, acc.ISCIDR, sshname, publicKey, volName, acc.ISZoneName3, name, acc.IsImage, acc.InstanceProfileName, acc.DedicatedHostGroupID, acc.ISZoneName, acc.DedicatedHostName)
}

func testAccCheckIBMISInstanceReservation(vpcname, subnetname, name, publickey, sshname string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	resource "ibm_is_ssh_key" "testacc_keyres" {
		name       = "%s"
		public_key = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		total_ipv4_address_count = 16
	}
	  
	resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet = ibm_is_subnet.testacc_subnet.id
		}
		vpc     = ibm_is_vpc.testacc_vpc.id
		zone    = "%s"
		reservation_affinity {
			policy = "manual"
			pool {
				id = "0735-b4a78f50-33bd-44f9-a3ff-4c33f444459d"
			}
		}
		keys = [ibm_is_ssh_key.testacc_keyres.id]
	}
	 
	  `, vpcname, sshname, publickey, subnetname, acc.ISZoneName3, name, acc.IsImage2, acc.InstanceProfileName, acc.ISZoneName3)
}

func testAccCheckIBMISInstanceByVolume(vpcname, subnetname, sshname, publicKey, volName, name, name1, sname string) string {
	return testAccCheckIBMISVolumeConfigSnapshot(vpcname, subnetname, sshname, publicKey, volName, name, sname) + fmt.Sprintf(`
	  
	  resource "ibm_is_instance" "instancebyvol" {
		name    = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet = ibm_is_subnet.testacc_subnet.id
		}
		vpc     = ibm_is_vpc.testacc_vpc.id
		zone    = "%s"
		keys    = [ibm_is_ssh_key.testacc_sshkey.id]
		
		boot_volume {
			volume_id = ibm_is_volume.storage.id
		}
	  }
	 
	  `, name1, acc.InstanceProfileName, acc.ISZoneName)
}

func testAccCheckIBMISInstanceVolumeAutoDelete(vpcname, subnetname, sshname, publicKey, volName, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_volume" "storage" {
		name    = "%s"
		profile = "10iops-tier"
		zone    = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet = ibm_is_subnet.testacc_subnet.id
		}
		vpc     			= ibm_is_vpc.testacc_vpc.id
		zone    			= "%s"
		keys    			= [ibm_is_ssh_key.testacc_sshkey.id]
		volumes 			= [ibm_is_volume.storage.id]
		auto_delete_volume 	= true
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, volName, acc.ISZoneName, name, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName)
}

func testAccCheckIBMISInstanceVolumeUpdate(vpcname, subnetname, sshname, publicKey, volName, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_volume" "storage" {
		name    = "%s"
		profile = "10iops-tier"
		zone    = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
	  }	  
	
`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, volName, acc.ISZoneName, name, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName)
}

func testAccCheckIBMISInstanceConfigWithAvailablePolicyHostFailure_Default(vpcname, subnetname, sshname, publicKey, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		}
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName)
}

func testAccCheckIBMISInstanceConfigWithAvailablePolicyHostFailure_Updated(vpcname, subnetname, sshname, publicKey, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		availability_policy_host_failure = "stop"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		}
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName)
}

func testAccCheckIBMISInstanceConfigWithAvailablePolicyHostFailure_WithTemplate(vpcname, subnetname, sshname, publicKey, templateName, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  data "ibm_is_images" "is_images" {
	  }
	  resource "ibm_is_instance_template" "instancetemplate1" {
		name    = "%s"
		image   = data.ibm_is_images.is_images.images.0.id
		profile = "bx2-8x32"
		primary_network_interface {
		  subnet = ibm_is_subnet.testacc_subnet.id
		}
		availability_policy_host_failure = "stop"
		vpc       = ibm_is_vpc.testacc_vpc.id
		zone      = "%s"
		keys      = [ibm_is_ssh_key.testacc_sshkey.id]
	  }

	  resource "ibm_is_instance" "testacc_instance" {
		name              = "%s"
  		instance_template   = ibm_is_instance_template.instancetemplate1.id
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, templateName, acc.ISZoneName, name)
}

func testAccCheckIBMISInstanceUserTagConfig(vpcname, subnetname, sshname, publicKey, name, userData, userTags1 string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }

	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }

	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }

	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		boot_volume {
			tags = ["%s"]
		}
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		user_data = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, userTags1, userData, acc.ISZoneName)
}

func TestAccIBMISInstance_catalog(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	userData1 := "a"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceCatalogImageConfig(vpcname, subnetname, sshname, publicKey, name, userData1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData1),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "catalog_offering.0.version_crn"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "catalog_offering.0.plan_crn"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_images.testacc_images", "images.0.name"),
				),
			},
		},
	})
}
func TestAccIBMISInstance_catalog_pna(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	userData1 := "a"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceCatalogImagePNAConfig(vpcname, subnetname, sshname, publicKey, name, userData1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData1),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "catalog_offering.0.version_crn"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_images.testacc_images", "images.0.name"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "primary_network_attachment.#"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceCatalogImageConfig(vpcname, subnetname, sshname, publicKey, name, userData string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }

	  data "ibm_is_images" "testacc_images" {
		catalog_managed = true
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		user_data = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_interfaces {
		  subnet = ibm_is_subnet.testacc_subnet.id
		  name   = "eth1"
		}
		catalog_offering {
			version_crn = data.ibm_is_images.testacc_images.images.0.catalog_offering.0.version.0.crn
			plan_crn = "crn:v1:bluemix:public:globalcatalog-collection:global:a/123456:51c9e0db-2911-45a6-adb0-ac5332d27cf2:plan:sw.51c9e0db-2911-45a6-adb0-ac5332d27cf2.772c0dbe-aa62-482e-adbe-a3fc20101e0e"
		}
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.InstanceProfileName, userData, acc.ISZoneName)
}
func testAccCheckIBMISInstanceCatalogImagePNAConfig(vpcname, subnetname, sshname, publicKey, name, userData string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }

	  data "ibm_is_images" "testacc_images" {
		catalog_managed = true
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		profile = "%s"
		primary_network_attachment {
			name = "testacc-instance-pna"
			virtual_network_interface {
				name = "testacc-instance-pna-vni"
				primary_ip {
					auto_delete 	= true
					address 		= cidrhost(ibm_is_subnet.testacc_subnet.ipv4_cidr_block, 23)
				}
				subnet = ibm_is_subnet.testacc_subnet.id
			}
		}
		user_data = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_attachments {
			name = "testacc-instance-sna"
			virtual_network_interface {
				name = "testacc-instance-sna-vni"
				primary_ip {
					auto_delete 	= true
					address 		= cidrhost(ibm_is_subnet.testacc_subnet.ipv4_cidr_block, 22)
				}
				subnet = ibm_is_subnet.testacc_subnet.id
			}
		}
		catalog_offering {
			version_crn = data.ibm_is_images.testacc_images.images.0.catalog_offering.0.version.0.crn
		}
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.InstanceProfileName, userData, acc.ISZoneName)
}

func TestAccIBMISInstance_volprototypes(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	userData1 := "a"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceVolumePrototypesConfig(vpcname, subnetname, sshname, publicKey, name, userData1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData1),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.0.manufacturer"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "volume_prototypes.#"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "volume_prototypes.#", "5"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceVolumePrototypesUpdate1Config(vpcname, subnetname, sshname, publicKey, name, userData1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData1),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "primary_network_interface.0.port_speed"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.0.manufacturer"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "numa_count"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "volume_prototypes.#"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "volume_prototypes.#", "6"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceVolumePrototypesUpdate2Config(vpcname, subnetname, sshname, publicKey, name, userData1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData1),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "primary_network_interface.0.port_speed"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.0.manufacturer"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "numa_count"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "volume_prototypes.#"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "volume_prototypes.#", "4"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceVolumePrototypesConfig(vpcname, subnetname, sshname, publicKey, name, userData string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		user_data = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		boot_volume {
			size = 250
			profile = "sdp"
      		iops = 10000
		}
		volume_prototypes{
		   name = "proto1"
		   delete_volume_on_instance_delete = true
		   volume_name = "proto1"
		   volume_capacity = 141
		   volume_profile = "sdp"
		}
		volume_prototypes{
		   name = "proto2"
		   delete_volume_on_instance_delete = true
		   volume_name = "proto2"
		   volume_capacity = 142
		   volume_profile = "sdp"
		}
		volume_prototypes{
		   name = "proto3"
		   delete_volume_on_instance_delete = true
		   volume_name = "proto3"
		   volume_profile = "general-purpose"
		   volume_capacity = 143
		}
		volume_prototypes{
		   name = "proto4"
		   delete_volume_on_instance_delete = true
		   volume_name = "proto4"
		   volume_capacity = 144
		   volume_iops = 10000
		   volume_profile = "sdp"
		}
	
		volume_prototypes{
		   name = "proto5"
		   delete_volume_on_instance_delete = true
		   volume_name = "proto55"
		   volume_capacity = 1455
		   volume_profile = "sdp"
		}
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, userData, acc.ISZoneName)
}
func testAccCheckIBMISInstanceVolumePrototypesUpdate1Config(vpcname, subnetname, sshname, publicKey, name, userData string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		user_data = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		boot_volume {
			size = 250
			profile = "sdp"
      		iops = 10000
		}
		volume_prototypes{
		   name = "proto1"
		   delete_volume_on_instance_delete = true
		   volume_name = "proto1"
		   volume_capacity = 141
		   # volume_iops = 1000
		   # volume_profile = "general-purpose"
		   volume_profile = "sdp"
		}
		volume_prototypes{
		   name = "proto2"
		   delete_volume_on_instance_delete = true
		   volume_name = "proto2"
		   volume_capacity = 142
		   # volume_iops = 1000
		   # volume_profile = "general-purpose"
		   volume_profile = "sdp"
		}
		volume_prototypes{
		   name = "proto3"
		   delete_volume_on_instance_delete = true
		   volume_name = "proto3"
		   volume_capacity = 143
		   # volume_iops = 1000
		   volume_profile = "general-purpose"
		   # volume_profile = "5iops-tier"
		   # volume_profile = "sdp"
		}
		volume_prototypes{
		   name = "proto4"
		   delete_volume_on_instance_delete = true
		   volume_name = "proto4"
		   volume_capacity = 144
		   volume_iops = 10000
		   # volume_profile = "general-purpose"
		   volume_profile = "sdp"
		}
	
	
		volume_prototypes{
		   name = "proto5"
		   delete_volume_on_instance_delete = true
		   volume_name = "proto55"
		   volume_capacity = 1455
		   # volume_iops = 1000
		   # volume_profile = "general-purpose"
		   volume_profile = "sdp"
		}
		volume_prototypes{
		  name = "proto6"
		  delete_volume_on_instance_delete = true
		  volume_name = "proto6"
		  volume_capacity = 146
		  volume_iops = 1000
		  # volume_profile = "general-purpose"
		  volume_profile = "sdp"
		}
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, userData, acc.ISZoneName)
}
func testAccCheckIBMISInstanceVolumePrototypesUpdate2Config(vpcname, subnetname, sshname, publicKey, name, userData string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		user_data = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		boot_volume {
			size = 250
			profile = "sdp"
      		iops = 10000
		}
		volume_prototypes{
		   name = "proto1"
		   delete_volume_on_instance_delete = true
		   volume_name = "proto1"
		   volume_capacity = 141
		   # volume_iops = 1000
		   # volume_profile = "general-purpose"
		   volume_profile = "sdp"
		}
		volume_prototypes{
		   name = "proto2"
		   delete_volume_on_instance_delete = true
		   volume_name = "proto2"
		   volume_capacity = 142
		   # volume_iops = 1000
		   # volume_profile = "general-purpose"
		   volume_profile = "sdp"
		}
		volume_prototypes{
		   name = "proto3"
		   delete_volume_on_instance_delete = true
		   volume_name = "proto3"
		   volume_capacity = 143
		   # volume_iops = 1000
		   volume_profile = "general-purpose"
		   # volume_profile = "5iops-tier"
		   # volume_profile = "sdp"
		}
		volume_prototypes{
		   name = "proto4"
		   delete_volume_on_instance_delete = true
		   volume_name = "proto4"
		   volume_capacity = 144
		   volume_iops = 10000
		   # volume_profile = "general-purpose"
		   volume_profile = "sdp"
		}
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, userData, acc.ISZoneName)
}

// cluster changes

func TestAccIBMISInstanceclusternetworkattachment_basic(t *testing.T) {
	var instance string
	randInt := acctest.RandIntRange(10, 100)
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	clustersubnetname := fmt.Sprintf("tf-clustersubnet-%d", acctest.RandIntRange(10, 100))
	clustersubnetreservedipname := fmt.Sprintf("tf-clustersubnet-reservedip-%d", acctest.RandIntRange(10, 100))
	clusterinterfacename := fmt.Sprintf("tf-clusterinterface-%d", acctest.RandIntRange(10, 100))

	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDVtuCfWKVGKaRmaRG6JQZY8YdxnDgGzVOK93IrV9R5Hl0JP1oiLLWlZQS2reAKb8lBqyDVEREpaoRUDjqDqXG8J/kR42FKN51su914pjSBc86wJ02VtT1Wm1zRbSg67kT+g8/T1jCgB5XBODqbcICHVP8Z1lXkgbiHLwlUrbz6OZkGJHo/M/kD1Eme8lctceIYNz/Ilm7ewMXZA4fsidpto9AjyarrJLufrOBl4MRVcZTDSJ7rLP982aHpu9pi5eJAjOZc7Og7n4ns3NFppiCwgVMCVUQbN5GBlWhZ1OsT84ZiTf+Zy8ew+Yg5T7Il8HuC7loWnz+esQPf0s3xhC/kTsGgZreIDoh/rxJfD67wKXetNSh5RH/n5BqjaOuXPFeNXmMhKlhj9nJ8scayx/wsvOGuocEIkbyJSLj3sLUU403OafgatEdnJOwbqg6rUNNF5RIjpJpL7eEWlKIi1j9LyhmPJ+fEO7TmOES82VpCMHpLbe4gf/MhhJ/Xy8DKh9s= root@ffd8363b1226
	`)
	subnetName := fmt.Sprintf("tf-testsubnet-%d", randInt)
	name := fmt.Sprintf("tf-testinstance-%d", randInt)
	updatedname := fmt.Sprintf("tf-testinstance-%d", randInt)
	sshKeyName := fmt.Sprintf("tf-testsshkey-%d", randInt)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceClusterNetworkAttachmentConfig(vpcname, clustersubnetname, clustersubnetreservedipname, clusterinterfacename, subnetName, sshKeyName, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.is_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.is_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.is_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.is_instance", "vcpu.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.is_instance", "vcpu.0.manufacturer"),
					resource.TestCheckResourceAttr("ibm_is_vpc.is_vpc", "name", vpcname),
					resource.TestCheckResourceAttrSet("ibm_is_vpc.is_vpc", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network.is_cluster_network_instance", "id"),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "id"),
					resource.TestCheckResourceAttr("ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance", "name", clustersubnetname),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_subnet_reserved_ip.is_cluster_network_subnet_reserved_ip_instance", "id"),
					resource.TestCheckResourceAttr("ibm_is_cluster_network_subnet_reserved_ip.is_cluster_network_subnet_reserved_ip_instance", "name", clustersubnetreservedipname),
					resource.TestCheckResourceAttrSet("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "id"),
					resource.TestCheckResourceAttr("ibm_is_cluster_network_interface.is_cluster_network_interface_instance", "name", clusterinterfacename),
					resource.TestCheckResourceAttrSet("ibm_is_subnet.is_subnet", "id"),
					resource.TestCheckResourceAttr("ibm_is_subnet.is_subnet", "name", subnetName),
					resource.TestCheckResourceAttrSet("ibm_is_ssh_key.is_sshkey", "id"),
					resource.TestCheckResourceAttr("ibm_is_ssh_key.is_sshkey", "name", sshKeyName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.is_instance", "name", name),
					resource.TestCheckResourceAttrSet("ibm_is_instance.is_instance", "profile"),
					resource.TestCheckResourceAttrSet("ibm_is_instance.is_instance", "crn"),
					resource.TestCheckResourceAttrSet("ibm_is_instance.is_instance", "image"),
					resource.TestCheckResourceAttrSet("ibm_is_instance.is_instance", "keys.#"),
					resource.TestCheckResourceAttrSet("ibm_is_instance.is_instance", "name"),
					resource.TestCheckResourceAttrSet("ibm_is_instance.is_instance", "resource_group"),
					resource.TestCheckResourceAttrSet("ibm_is_instance.is_instance", "vpc"),
					resource.TestCheckResourceAttrSet("ibm_is_instance.is_instance", "zone"),
					resource.TestCheckResourceAttrSet("ibm_is_instance.is_instance", "boot_volume.#"),
					resource.TestCheckResourceAttrSet("ibm_is_instance.is_instance", "cluster_network_attachments.#"),
					resource.TestCheckResourceAttr("ibm_is_instance.is_instance", "cluster_network_attachments.#", "8"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceClusterNetworkAttachmentConfig(vpcname, clustersubnetname, clustersubnetreservedipname, clusterinterfacename, subnetName, sshKeyName, publicKey, updatedname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "primary_network_interface.0.port_speed"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.0.manufacturer"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "numa_count"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceClusterNetworkAttachmentConfig(vpcname, clustersubnetname, clustersubnetreservedipname, clusternetworkinterfacename, subnetName, sshKeyName, publicKey, instanceName string) string {
	return fmt.Sprintf(`
		resource "ibm_is_vpc" "is_vpc" {
  			name = "%s"
		}
		resource "ibm_is_cluster_network" "is_cluster_network_instance" {
			profile = "%s"
			vpc {
				id = ibm_is_vpc.is_vpc.id
			}
			zone  = "%s"
		}
		resource "ibm_is_cluster_network_subnet" "is_cluster_network_subnet_instance" {
			cluster_network_id = ibm_is_cluster_network.is_cluster_network_instance.id
			name = "%s"
			total_ipv4_address_count = 64
		}
		resource "ibm_is_cluster_network_subnet_reserved_ip" "is_cluster_network_subnet_reserved_ip_instance" {
			cluster_network_id 			= ibm_is_cluster_network.is_cluster_network_instance.id
			cluster_network_subnet_id 	= ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
			address 					= "${replace(ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.ipv4_cidr_block, "0/26", "11")}"
  			name						= "%s"
		}
		resource "ibm_is_cluster_network_interface" "is_cluster_network_interface_instance" {
			cluster_network_id = ibm_is_cluster_network.is_cluster_network_instance.id
			name = "%s"
			primary_ip {
				id = ibm_is_cluster_network_subnet_reserved_ip.is_cluster_network_subnet_reserved_ip_instance.cluster_network_subnet_reserved_ip_id
			}
			subnet {
				id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
			}
		}
	
		resource "ibm_is_subnet" "is_subnet" {
			name            			= "%s"
			vpc             			= ibm_is_vpc.is_vpc.id
			zone            			= "%s"
			total_ipv4_address_count 	= 64
		}
		
		resource "ibm_is_ssh_key" "is_sshkey" {
			name       = "%s"
			public_key = "%s"
		}
		resource "ibm_is_instance" "is_instance" {
			name    = "%s"
			image   = "%s"
			profile = "%s"
			timeouts {
				create = "60m"
  			}
			primary_network_attachment {
				name 		= "my-pna"
				virtual_network_interface {
					auto_delete = true
					subnet      = ibm_is_subnet.is_subnet.id
				}
			}
			cluster_network_attachments {
				name = "cna-1"
				cluster_network_interface{
					auto_delete = true
					name = "cni-1"
					subnet {
						id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
					}
				}
			}
			cluster_network_attachments {
				name = "cna-2"
				cluster_network_interface{
					auto_delete = true
					name = "cni-2"
					subnet {
						id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
					}
				}
			}
			cluster_network_attachments {
				name = "cna-3"
				cluster_network_interface{
					auto_delete = true
					name = "cni-3"
					subnet {
						id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
					}
				}
			}
			cluster_network_attachments {
				name = "cna-4"
				cluster_network_interface{
					auto_delete = true
					name = "cni-4"
					subnet {
						id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
					}
				}
			}
			cluster_network_attachments {
				name = "cna-5"
				cluster_network_interface{
					auto_delete = true
					name = "cni-5"
					subnet {
						id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
					}
				}
			}
			cluster_network_attachments {
				name = "cna-6"
				cluster_network_interface{
					auto_delete = true
					name = "cni-6"
					subnet {
						id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
					}
				}
			}
			cluster_network_attachments {
				name = "cna-7"
				cluster_network_interface{
					auto_delete = true
					name = "cni-7"
					subnet {
						id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
					}
				}
			}
			cluster_network_attachments {
				name = "cna-8"
				cluster_network_interface{
					auto_delete = true
					name = "cni-8"
					subnet {
						id = ibm_is_cluster_network_subnet.is_cluster_network_subnet_instance.cluster_network_subnet_id
					}
				}
			}
			vpc       = ibm_is_vpc.is_vpc.id
			zone      = ibm_is_subnet.is_subnet.zone
			keys      = [ibm_is_ssh_key.is_sshkey.id]
		}
	`, vpcname, acc.ISClusterNetworkProfileName, acc.ISZoneName, clustersubnetname, clustersubnetreservedipname, clusternetworkinterfacename, subnetName, acc.ISZoneName, sshKeyName, publicKey, instanceName, acc.IsImage, acc.ISInstanceGPUProfileName)
}

func TestAccIBMISInstance_primary_ip_consistency(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	userData1 := "a"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstancePrimaryIpConsistencyConfig(vpcname, subnetname, sshname, publicKey, name, userData1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData1),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "vcpu.0.manufacturer"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "primary_network_attachment.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "primary_network_attachment.0.primary_ip.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "primary_network_attachment.0.primary_ip.0.address"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "network_attachments.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "network_attachments.0.primary_ip.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "network_attachments.0.primary_ip.0.address"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstancePrimaryIpConsistencyConfig(vpcname, subnetname, sshname, publicKey, name, userData string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		primary_network_attachment {
			name = "example-primarynetwork-att"
			virtual_network_interface {
				auto_delete = true
				subnet = ibm_is_subnet.testacc_subnet.id
			}
		}
		user_data = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		network_attachments {
			name = "example-network-att"
			virtual_network_interface {
				auto_delete = true
				subnet      = ibm_is_subnet.testacc_subnet.id
			}
		}
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, userData, acc.ISZoneName)
}

// TestAccIBMISInstanceResourceGroupChangeVNI tests that changing the resource group of VNI
// forces new creation of both VNI and instance
func TestAccIBMISInstanceResourceGroupChangeVNI(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	vniname := fmt.Sprintf("tf-vni-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	rg1 := acc.IsResourceGroupID
	rg2 := acc.IsResourceGroupIDUpdate

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			// Initial setup with resource group 1
			{
				Config: testAccCheckIBMISInstanceWithVNIResourceGroup(vpcname, subnetname, sshname, publicKey, name, vniname, rg1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_virtual_network_interface.testacc_vni", "resource_group", rg1),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "primary_network_attachment.#"),
				),
			},
			// Change resource group - should force new creation
			{
				Config: testAccCheckIBMISInstanceWithVNIResourceGroup(vpcname, subnetname, sshname, publicKey, name, vniname, rg2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_virtual_network_interface.testacc_vni", "resource_group", rg2),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance", "primary_network_attachment.#"),
				),
			},
		},
	})
}

// TestAccIBMISInstanceInlineVNIResourceGroupChange tests instance with inline VNI when resource group changes
func TestAccIBMISInstanceInlineVNIResourceGroupChange(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	rg1 := acc.IsResourceGroupID
	rg2 := acc.IsResourceGroupIDUpdate

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			// Initial setup with resource group 1
			{
				Config: testAccCheckIBMISInstanceWithInlineVNIResourceGroup(vpcname, subnetname, sshname, publicKey, name, rg1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "primary_network_attachment.0.virtual_network_interface.0.resource_group", rg1),
				),
			},
			// Change resource group - should force new creation
			{
				Config: testAccCheckIBMISInstanceWithInlineVNIResourceGroup(vpcname, subnetname, sshname, publicKey, name, rg2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "primary_network_attachment.0.virtual_network_interface.0.resource_group", rg2),
				),
			},
		},
	})
}

// Config generators

func testAccCheckIBMISInstanceWithVNIResourceGroup(vpcname, subnetname, sshname, publicKey, name, vniname, resourceGroup string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}

resource "ibm_is_subnet" "testacc_subnet" {
	name            = "%s"
	vpc             = ibm_is_vpc.testacc_vpc.id
	zone            = "%s"
	ipv4_cidr_block = "%s"
}

resource "ibm_is_ssh_key" "testacc_sshkey" {
	name       = "%s"
	public_key = "%s"
}

resource "ibm_is_virtual_network_interface" "testacc_vni" {
	name = "%s"
	subnet = ibm_is_subnet.testacc_subnet.id
	resource_group = "%s"
	auto_delete = false
} 

resource "ibm_is_instance" "testacc_instance" {
	name    = "%s"
	image   = "%s"
	profile = "%s"
    primary_network_attachment {
        name = "test-vni"
        virtual_network_interface { 
            id = ibm_is_virtual_network_interface.testacc_vni.id
        }
    }
	vpc  = ibm_is_vpc.testacc_vpc.id
	zone = "%s"
	keys = [ibm_is_ssh_key.testacc_sshkey.id]
}`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, vniname, resourceGroup, name, acc.IsImage, acc.InstanceProfileName, acc.ISZoneName)
}

func testAccCheckIBMISInstanceWithInlineVNIResourceGroup(vpcname, subnetname, sshname, publicKey, name, resourceGroup string) string {
	return fmt.Sprintf(`
resource "ibm_is_vpc" "testacc_vpc" {
	name = "%s"
}

resource "ibm_is_subnet" "testacc_subnet" {
	name            = "%s"
	vpc             = ibm_is_vpc.testacc_vpc.id
	zone            = "%s"
	ipv4_cidr_block = "%s"
}

resource "ibm_is_ssh_key" "testacc_sshkey" {
	name       = "%s"
	public_key = "%s"
}

resource "ibm_is_instance" "testacc_instance" {
	name    = "%s"
	image   = "%s"
	profile = "%s"
    primary_network_attachment {
        name = "test-vni-inline"
        virtual_network_interface {
			subnet = ibm_is_subnet.testacc_subnet.id
			resource_group = "%s"
        }
    }
	vpc  = ibm_is_vpc.testacc_vpc.id
	zone = "%s"
	keys = [ibm_is_ssh_key.testacc_sshkey.id]
}`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, resourceGroup, acc.ISZoneName)
}

// bandwidth changes

func TestAccIBMISInstance_BootBandwidth(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	userData1 := "a"
	bandwidth1 := int64(2200)
	bandwidth2 := int64(2500)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceBootBandwidthConfig(vpcname, subnetname, sshname, publicKey, name, userData1, bandwidth1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData1),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "boot_volume.0.bandwidth", fmt.Sprintf("%d", bandwidth1)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
				),
			},
			{
				Config: testAccCheckIBMISInstanceBootBandwidthConfig(vpcname, subnetname, sshname, publicKey, name, userData1, bandwidth2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData1),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "boot_volume.0.bandwidth", fmt.Sprintf("%d", bandwidth2)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
				),
			},
		},
	})
}
func TestAccIBMISInstance_VolumeBandwidth(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	userData1 := "a"
	bandwidth1 := int64(2200)
	bandwidth2 := int64(2500)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceVolumeBandwidthConfig(vpcname, subnetname, sshname, publicKey, name, userData1, bandwidth1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData1),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "volume_prototypes.0.volume_bandwidth", fmt.Sprintf("%d", bandwidth1)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
				),
			},
			{
				Config: testAccCheckIBMISInstanceVolumeBandwidthConfig(vpcname, subnetname, sshname, publicKey, name, userData1, bandwidth2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "user_data", userData1),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "volume_prototypes.0.volume_bandwidth", fmt.Sprintf("%d", bandwidth2)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceBootBandwidthConfig(vpcname, subnetname, sshname, publicKey, name, userData string, bandwidtth int64) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		boot_volume {
			profile 	= "sdp"
			bandwidth 	= %d
		}
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		user_data = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, bandwidtth, userData, acc.ISZoneName)
}
func testAccCheckIBMISInstanceVolumeBandwidthConfig(vpcname, subnetname, sshname, publicKey, name, userData string, bandwidtth int64) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	  }
	  
	  resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  }
	  
	  resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	  }
	  
	  resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		volume_prototypes {
			volume_profile 						= "sdp"
			delete_volume_on_instance_delete	= true
			name 								= "test-data-vol-bandwidth-att"
			volume_name 						= "test-data-vol-bandwidth"
			volume_bandwidth 					= %d
			volume_capacity 					= 150
		}
		primary_network_interface {
		  subnet     = ibm_is_subnet.testacc_subnet.id
		}
		user_data = "%s"
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, acc.InstanceProfileName, bandwidtth, userData, acc.ISZoneName)
}

// tdx testing

func TestAccIBMISInstanceTDX_basic(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instance-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceConfigTDX(vpcname, subnetname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "confidential_compute_mode", "tdx"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "profile", "bx3dc-2x10"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
				),
			},
		},
	})
}

func TestAccIBMISInstanceSGXtoTDX_basic(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instance-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceConfigSGX(vpcname, subnetname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "confidential_compute_mode", "sgx"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "profile", "bx3dc-2x10"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceActionStopSGX(vpcname, subnetname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance_action.testacc_instanceaction", "action", "stop"),
				),
			},
			{
				Config: testAccCheckIBMISInstanceConfigTDX(vpcname, subnetname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "confidential_compute_mode", "tdx"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "profile", "bx3dc-2x10"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceConfigSGX(vpcname, subnetname, sshname, publicKey, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	  
	resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	}
	  
	resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	}
	  
	resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "bx3dc-2x10"
		confidential_compute_mode = "sgx"
		primary_network_interface {
			subnet = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
	}`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, acc.ISZoneName)
}

func testAccCheckIBMISInstanceActionStopSGX(vpcname, subnetname, sshname, publicKey, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	  
	resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	}
	  
	resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	}
	  
	resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "bx3dc-2x10"
		confidential_compute_mode = "sgx"
		primary_network_interface {
			subnet = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
	}

	resource "ibm_is_instance_action" "testacc_instanceaction" {
		depends_on = [ibm_is_instance.testacc_instance]
		action = "stop"
		instance = ibm_is_instance.testacc_instance.id
	}`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, acc.ISZoneName)
}

func testAccCheckIBMISInstanceConfigTDX(vpcname, subnetname, sshname, publicKey, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	  
	resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	}
	  
	resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	}
	  
	resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "bx3dc-2x10"
		confidential_compute_mode = "tdx"
		primary_network_interface {
			subnet = ibm_is_subnet.testacc_subnet.id
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
	}`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, acc.ISZoneName)
}

// sgx
// Create a basic instance with SGX confidential compute mode
func TestAccIBMISInstanceSGX_basic(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instance-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceConfigSGX(vpcname, subnetname, sshname, publicKey, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "confidential_compute_mode", "sgx"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "profile", "bx3dc-2x10"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
				),
			},
		},
	})
}

func TestAccIBMISInstance_ProfileAndBandwidthUpdate(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instance-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	initialProfile := "cx2-4x8"   // Initial profile
	updatedProfile := "cx2-48x96" // Updated profile
	initialBandwidth := 2000      // Initial bandwidth
	updatedBandwidth := 20000     // Updated bandwidth
	prefix := fmt.Sprintf("tf-prefix-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create instance with initial profile and bandwidth
			{
				Config: testAccCheckIBMISInstanceConfigWithProfileAndBandwidth(
					vpcname, subnetname, sshname, publicKey, name, prefix, initialProfile, initialBandwidth),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "profile", initialProfile),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "total_volume_bandwidth", fmt.Sprintf("%d", initialBandwidth)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
				),
			},
			// Step 2: Update both profile and bandwidth in a single operation
			{
				Config: testAccCheckIBMISInstanceConfigWithProfileAndBandwidth(
					vpcname, subnetname, sshname, publicKey, name, prefix, updatedProfile, updatedBandwidth),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "profile", updatedProfile),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "total_volume_bandwidth", fmt.Sprintf("%d", updatedBandwidth)),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
				),
			},
		},
	})
}

// Configuration function that allows specifying both profile and bandwidth with primary network attachment
func testAccCheckIBMISInstanceConfigWithProfileAndBandwidth(vpcname, subnetname, sshname, publicKey, name, prefix, profile string, bandwidth int) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	  
	resource "ibm_is_subnet" "testacc_subnet" {
		name            = "%s"
		vpc             = ibm_is_vpc.testacc_vpc.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	}
	  
	resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	}
	  
	resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = "%s"
		profile = "%s"
		total_volume_bandwidth = %d
		primary_network_attachment {
			name = "%s-pna"
			virtual_network_interface {
				subnet = ibm_is_subnet.testacc_subnet.id
			}
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
		wait_before_delete = false
	}`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, sshname, publicKey, name, acc.IsImage, profile, bandwidth, prefix, acc.ISZoneName)
}

// volume tags

func TestAccIBMISInstance_volumeTags(t *testing.T) {
	var instance string
	prefix := fmt.Sprintf("tf-inst-%d", acctest.RandIntRange(10, 100))
	vpcname := fmt.Sprintf("%s-vpc", prefix)
	name := fmt.Sprintf("%s-instance", prefix)
	subnetname := fmt.Sprintf("%s-subnet", prefix)
	sshname := fmt.Sprintf("%s-ssh", prefix)
	dataVolumeName := fmt.Sprintf("%s-data-volume", prefix)

	publicKey := strings.TrimSpace(`
ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIBEGGaXOYllPYQE+Qj8MiRo7DOJK9j7K8OQE9VWL5VjZ terraform-test-key
`)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceVolumeTagsConfig(prefix, vpcname, subnetname, sshname, publicKey, name, dataVolumeName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					// Instance checks
					resource.TestCheckResourceAttr("ibm_is_instance.testacc_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_is_instance.testacc_instance", "profile", "bx2-2x8"),
					resource.TestCheckResourceAttr("ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
					resource.TestCheckResourceAttr("ibm_is_instance.testacc_instance", "tags.#", "1"),
					resource.TestCheckResourceAttr("ibm_is_instance.testacc_instance", "tags.0", "tagged:byuser"),

					// Volume prototype checks
					resource.TestCheckResourceAttr("ibm_is_instance.testacc_instance", "volume_prototypes.#", "1"),
					resource.TestCheckResourceAttr("ibm_is_instance.testacc_instance", "volume_prototypes.0.volume_name", dataVolumeName),
					resource.TestCheckResourceAttr("ibm_is_instance.testacc_instance", "volume_prototypes.0.volume_capacity", "141"),
					resource.TestCheckResourceAttr("ibm_is_instance.testacc_instance", "volume_prototypes.0.volume_profile", "general-purpose"),
					resource.TestCheckResourceAttr("ibm_is_instance.testacc_instance", "volume_prototypes.0.delete_volume_on_instance_delete", "true"),
					resource.TestCheckResourceAttr("ibm_is_instance.testacc_instance", "volume_prototypes.0.volume_tags.#", "1"),
					resource.TestCheckResourceAttr("ibm_is_instance.testacc_instance", "volume_prototypes.0.volume_tags.0", "tagged:byuser"),

					// Boot volume checks
					resource.TestCheckResourceAttr("ibm_is_instance.testacc_instance", "boot_volume.#", "1"),
					resource.TestCheckResourceAttr("ibm_is_instance.testacc_instance", "boot_volume.0.tags.#", "1"),
					resource.TestCheckResourceAttr("ibm_is_instance.testacc_instance", "boot_volume.0.tags.0", "tagged:byuser"),
					resource.TestCheckResourceAttrSet("ibm_is_instance.testacc_instance", "boot_volume.0.volume_id"),

					// Primary network attachment checks
					resource.TestCheckResourceAttr("ibm_is_instance.testacc_instance", "primary_network_attachment.#", "1"),
					resource.TestCheckResourceAttr("ibm_is_instance.testacc_instance", "primary_network_attachment.0.name", fmt.Sprintf("%s-pna", prefix)),
					resource.TestCheckResourceAttrSet("ibm_is_instance.testacc_instance", "primary_network_attachment.0.virtual_network_interface.0.subnet"),

					// Data source checks for boot volume tags
					resource.TestCheckResourceAttr("data.ibm_is_volume.boot", "tags.0", "tagged:byuser"),
					// Data source checks for data volume tags
					resource.TestCheckResourceAttr("data.ibm_is_volume.data", "tags.0", "tagged:byuser"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceVolumeTagsConfig(prefix, vpcname, subnetname, sshname, publicKey, name, dataVolumeName string) string {
	return fmt.Sprintf(`
	# VPC Infrastructure
	resource "ibm_is_vpc" "is_vpc" {
		name = "%s"
	}

	resource "ibm_is_subnet" "is_subnet" {
		name                     = "%s"
		vpc                      = ibm_is_vpc.is_vpc.id
		total_ipv4_address_count = 64
		zone                     = "%s"
	}

	# Data sources
	data "ibm_is_image" "is_image" {
		name = "ibm-ubuntu-20-04-6-minimal-amd64-6"
	}

	# SSH Key
	resource "ibm_is_ssh_key" "is_key" {
		name       = "%s"
		public_key = "%s"
		type       = "ed25519"
	}

	# Instance with tagged volumes
	resource "ibm_is_instance" "testacc_instance" {
		name    = "%s"
		image   = data.ibm_is_image.is_image.id
		profile = "bx2-2x8"
		
		primary_network_attachment {
			name = "%s-pna"
			virtual_network_interface {
				subnet = ibm_is_subnet.is_subnet.id
			}
		}
		
		# Boot volume with tags
		boot_volume {
			tags = ["tagged:byuser"]
		}

		# Data volume with tags
		volume_prototypes {
			name                             = "%s"
			delete_volume_on_instance_delete = true
			volume_name                      = "%s"
			volume_capacity                  = 141
			volume_profile                   = "general-purpose"
			volume_tags                      = ["tagged:byuser"]
		}
		
		vpc                = ibm_is_vpc.is_vpc.id
		zone               = ibm_is_subnet.is_subnet.zone
		keys               = [ibm_is_ssh_key.is_key.id]
		wait_before_delete = false
		
		tags = ["tagged:byuser"]
	}

	# Data sources to verify volume tags
	data "ibm_is_volume" "boot" {
		identifier = ibm_is_instance.testacc_instance.boot_volume.0.volume_id
	}

	data "ibm_is_volume" "data" {
		name = ibm_is_instance.testacc_instance.volume_prototypes.0.volume_name
	}
	`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, name, prefix, dataVolumeName, dataVolumeName)
}

func TestAccIBMISInstance_AllowedUse(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf-instance-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	volname := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfsnapshotuat-%d", acctest.RandIntRange(10, 100))
	apiVersion := "2025-07-02"
	bareMetalServer := "true"
	instanceval := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceConfig_AllowedUse(vpcname, subnetname, sshname, publicKey, volname, name, name1, apiVersion, bareMetalServer, instanceval, instanceName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance_allowed_use", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance_allowed_use", "name", instanceName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance_allowed_use", "zone", acc.ISZoneName),

					// Volume prototype checks
					resource.TestCheckResourceAttr("ibm_is_instance.testacc_instance_allowed_use", "volume_prototypes.#", "1"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance_allowed_use", "volume_prototypes.0.allowed_use.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance_allowed_use", "volume_prototypes.0.allowed_use.0.bare_metal_server"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance_allowed_use", "volume_prototypes.0.allowed_use.0.instance"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance_allowed_use", "volume_prototypes.0.allowed_use.0.api_version"),
					resource.TestCheckResourceAttr("ibm_is_instance.testacc_instance_allowed_use", "volume_prototypes.0.allowed_use.0.bare_metal_server", bareMetalServer),
					resource.TestCheckResourceAttr("ibm_is_instance.testacc_instance_allowed_use", "volume_prototypes.0.allowed_use.0.instance", instanceval),
					resource.TestCheckResourceAttr("ibm_is_instance.testacc_instance_allowed_use", "volume_prototypes.0.allowed_use.0.api_version", apiVersion),

					// Boot volume checks
					resource.TestCheckResourceAttr("ibm_is_instance.testacc_instance_allowed_use", "boot_volume.#", "1"),
					resource.TestCheckResourceAttrSet("ibm_is_instance.testacc_instance_allowed_use", "boot_volume.0.volume_id"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance_allowed_use", "boot_volume.0.allowed_use.#"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance_allowed_use", "boot_volume.0.allowed_use.0.bare_metal_server"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance_allowed_use", "boot_volume.0.allowed_use.0.instance"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance_allowed_use", "boot_volume.0.allowed_use.0.api_version"),
					resource.TestCheckResourceAttr("ibm_is_instance.testacc_instance_allowed_use", "boot_volume.0.allowed_use.0.bare_metal_server", bareMetalServer),
					resource.TestCheckResourceAttr("ibm_is_instance.testacc_instance_allowed_use", "boot_volume.0.allowed_use.0.instance", instanceval),
					resource.TestCheckResourceAttr("ibm_is_instance.testacc_instance_allowed_use", "boot_volume.0.allowed_use.0.api_version", apiVersion),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceConfig_AllowedUse(vpcname, subnetname, sshname, publicKey, volname, name, name1, apiVersion, bareMetalServer, instanceval, insName string) string {

	return testAccCheckIBMISSnapshotConfig(vpcname, subnetname, sshname, publicKey, volname, name, name1) + fmt.Sprintf(`
	resource "ibm_is_instance" "testacc_instance_allowed_use" {
	name    = "%s"
	profile = "%s"
	primary_network_interface {
		subnet = ibm_is_subnet.testacc_subnet.id
	}
	vpc  = ibm_is_vpc.testacc_vpc.id
	zone = "%s"
	keys = [ibm_is_ssh_key.testacc_sshkey.id]
	boot_volume {
		name     = "example-boot-volume"
		snapshot = ibm_is_snapshot.testacc_snapshot.id
		size = 100
		allowed_use {
			api_version       = "%s"
			instance          = "%s"
			bare_metal_server = "%s"
		}
	}
	volume_prototypes {
		name                             = "example-prototype"
		delete_volume_on_instance_delete = true
		volume_name                      = "example-volume"
		volume_capacity                  = 100
		volume_profile                   = "custom"
		volume_source_snapshot           = ibm_is_snapshot.testacc_snapshot.id
		allowed_use {
			api_version       = "%s"
			bare_metal_server = "%s"
			instance          = "%s"
		}
	}
	}
	`, insName, acc.InstanceProfileName, acc.ISZoneName, apiVersion, bareMetalServer, instanceval, apiVersion, bareMetalServer, instanceval)
}

// boot volume profile test

func TestAccIBMISInstance_BootVolumeVariations(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	sshname := fmt.Sprintf("tf-ssh-%d", acctest.RandIntRange(10, 100))
	templatename := fmt.Sprintf("tf-template-%d", acctest.RandIntRange(10, 100))
	sourceInstanceName := fmt.Sprintf("tf-instance-source-%d", acctest.RandIntRange(10, 100))
	instanceFromTemplateName := fmt.Sprintf("tf-instance-template-%d", acctest.RandIntRange(10, 100))
	instanceFromCatalogName := fmt.Sprintf("tf-instance-catalog-%d", acctest.RandIntRange(10, 100))
	snapshotname := fmt.Sprintf("tf-snapshot-%d", acctest.RandIntRange(10, 100))
	instanceFromSnapshotName := fmt.Sprintf("tf-instance-snapshot-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceBootVolumeConfig(vpcname, subnetname, sshname, publicKey, templatename, sourceInstanceName, snapshotname, instanceFromTemplateName, instanceFromCatalogName, instanceFromSnapshotName),
				Check: resource.ComposeTestCheckFunc(
					// Verify instance from template with boot volume
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance_template", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance_template", "name", instanceFromTemplateName),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance_template", "instance_template"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance_template", "boot_volume.0.profile", "sdp"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance_template", "boot_volume.0.name"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance_template", "primary_network_attachment.0.id"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance_template", "zone", acc.ISZoneName),

					// Verify instance from catalog with boot volume
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance_catalog", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance_catalog", "name", instanceFromCatalogName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance_catalog", "boot_volume.0.profile", "sdp"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance_catalog", "boot_volume.0.name"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance_catalog", "catalog_offering.0.version_crn"),
					resource.TestCheckResourceAttrSet(
						"ibm_is_instance.testacc_instance_catalog", "primary_network_attachment.0.id"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance_catalog", "zone", acc.ISZoneName),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceBootVolumeConfig(vpcname, subnetname, sshname, publicKey, templatename, sourceInstanceName, snapshotname, instanceFromTemplateName, instanceFromCatalogName, instanceFromSnapshotName string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
	}
	
	resource "ibm_is_subnet" "testacc_subnet" {
		name                     = "%s"
		vpc                      = ibm_is_vpc.testacc_vpc.id
		zone                     = "%s"
		total_ipv4_address_count = 64
	}
	
	resource "ibm_is_ssh_key" "testacc_sshkey" {
		name       = "%s"
		public_key = "%s"
	}
	
	resource "ibm_is_instance_template" "instancetemplate1" {
		name    = "%s"
		image   = "%s"
		profile = "bxf-2x8"
		primary_network_attachment {
			name = "pna-template"
			virtual_network_interface {
				subnet = ibm_is_subnet.testacc_subnet.id
			}
		}
		vpc       = ibm_is_vpc.testacc_vpc.id
		zone      = "%s"
		keys      = [ibm_is_ssh_key.testacc_sshkey.id]
	}

	resource "ibm_is_instance" "testacc_instance_source" {
		name              = "%s"
		primary_network_attachment {
			name = "pna-from-template"
			virtual_network_interface {
				subnet = ibm_is_subnet.testacc_subnet.id
			}
		}
		instance_template = ibm_is_instance_template.instancetemplate1.id
	}

	resource "ibm_is_instance" "testacc_instance_template" {
		name              = "%s"
		boot_volume {
			profile = "sdp"
		}
		primary_network_attachment {
			name = "pna-ins-from-template"
			virtual_network_interface {
				subnet = ibm_is_subnet.testacc_subnet.id
			}
		}
		instance_template = ibm_is_instance_template.instancetemplate1.id
	}

	data "ibm_is_image" "catalog_image" {
		name = "%s"
	}

	resource "ibm_is_instance" "testacc_instance_catalog" {
		name    = "%s"
		profile = "bxf-2x8"
		primary_network_attachment {
			name = "pna-catalog"
			virtual_network_interface {
				subnet = ibm_is_subnet.testacc_subnet.id
			}
		}
		boot_volume {
			profile = "sdp"
		}
		catalog_offering {
			version_crn = data.ibm_is_image.catalog_image.catalog_offering.0.version.0.crn
		}
		vpc  = ibm_is_vpc.testacc_vpc.id
		zone = ibm_is_subnet.testacc_subnet.zone
		keys = [ibm_is_ssh_key.testacc_sshkey.id]
	}

	`, vpcname, subnetname, acc.ISZoneName, sshname, publicKey, templatename, acc.IsImage, acc.ISZoneName, sourceInstanceName, instanceFromTemplateName, acc.ISCatalogImageName, instanceFromCatalogName)
}
