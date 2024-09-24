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
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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
	ipv4address := "10.240.0.6"

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
