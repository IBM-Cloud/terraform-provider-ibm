// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMISInstancesDataSource_basic(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tfins-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfins-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tfins-ssh-%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tfins-name-%d", acctest.RandIntRange(10, 100))
	resName := "data.ibm_is_instances.ds_instances"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceConfig(vpcname, subnetname, sshname, publicKey, instanceName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", instanceName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", ISZoneName),
				),
			},
			{
				Config: testAccCheckIBMISInstancesDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "instances.0.name"),
					resource.TestCheckResourceAttrSet(resName, "instances.0.memory"),
					resource.TestCheckResourceAttrSet(resName, "instances.0.status"),
					resource.TestCheckResourceAttrSet(resName, "instances.0.resource_group"),
					resource.TestCheckResourceAttrSet(resName, "instances.0.vpc"),
					resource.TestCheckResourceAttrSet(resName, "instances.0.boot_volume.#"),
					resource.TestCheckResourceAttrSet(resName, "instances.0.volume_attachments.#"),
					resource.TestCheckResourceAttrSet(resName, "instances.0.primary_network_interface.#"),
					resource.TestCheckResourceAttrSet(resName, "instances.0.network_interfaces.#"),
					resource.TestCheckResourceAttrSet(resName, "instances.0.profile"),
					resource.TestCheckResourceAttrSet(resName, "instances.0.vcpu.#"),
					resource.TestCheckResourceAttrSet(resName, "instances.0.zone"),
				),
			},
		},
	})
}

func TestAccIBMISInstancesDataSource_vpcfilter(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tfins-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfins-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tfins-ssh-%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tfins-name-%d", acctest.RandIntRange(10, 100))
	resName := "data.ibm_is_instances.ds_instances1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceConfig(vpcname, subnetname, sshname, publicKey, instanceName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", instanceName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", ISZoneName),
				),
			},
			{
				Config: testAccCheckIBMISInstancesDataSourceConfig1(vpcname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "instances.0.name"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstancesDataSourceConfig() string {
	return fmt.Sprintf(`
	data "ibm_is_instances" "ds_instances" {
	}`)
}

func testAccCheckIBMISInstancesDataSourceConfig1(vpcname string) string {
	return fmt.Sprintf(`
	data "ibm_is_instances" "ds_instances1" {
		vpc_name = "%s"
	}`, vpcname)
}
