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
	userData := "a"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceConfig(vpcname, subnetname, sshname, publicKey, instanceName, userData),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", instanceName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
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
					resource.TestCheckResourceAttrSet(resName, "instances.0.availability_policy.#"),
				),
			},
		},
	})
}
func TestAccIBMISInstancesDataSource_ReservedIp(t *testing.T) {
	var instance string
	vpcname := fmt.Sprintf("tfins-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfins-subnet-%d", acctest.RandIntRange(10, 100))
	publicKey := strings.TrimSpace(`
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
`)
	sshname := fmt.Sprintf("tfins-ssh-%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tfins-name-%d", acctest.RandIntRange(10, 100))
	resName := "data.ibm_is_instances.ds_instances"
	userData := "a"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceConfig(vpcname, subnetname, sshname, publicKey, instanceName, userData),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", instanceName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
				),
			},
			{
				Config: testAccCheckIBMISInstancesDataSourceReservedIpConfig(vpcname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "instances.0.name"),
					resource.TestCheckResourceAttrSet(resName, "instances.0.memory"),
					resource.TestCheckResourceAttrSet(resName, "instances.0.status"),
					resource.TestCheckResourceAttrSet(resName, "instances.0.resource_group"),
					resource.TestCheckResourceAttrSet(resName, "instances.0.vpc"),
					resource.TestCheckResourceAttrSet(resName, "instances.0.boot_volume.#"),
					resource.TestCheckResourceAttrSet(resName, "instances.0.volume_attachments.#"),
					resource.TestCheckResourceAttrSet(resName, "instances.0.primary_network_interface.#"),
					resource.TestCheckResourceAttrSet(resName, "instances.0.primary_network_interface.0.primary_ip.0.reserved_ip"),
					resource.TestCheckResourceAttrSet(resName, "instances.0.primary_network_interface.0.primary_ip.0.href"),
					resource.TestCheckResourceAttrSet(resName, "instances.0.primary_network_interface.0.primary_ip.0.address"),
					resource.TestCheckResourceAttrSet(resName, "instances.0.primary_network_interface.0.primary_ip.0.resource_type"),
					resource.TestCheckResourceAttrSet(resName, "instances.0.primary_network_interface.0.primary_ip.0.name"),
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
	userData := "a"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceConfig(vpcname, subnetname, sshname, publicKey, instanceName, userData),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISInstanceExists("ibm_is_instance.testacc_instance", instance),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "name", instanceName),
					resource.TestCheckResourceAttr(
						"ibm_is_instance.testacc_instance", "zone", acc.ISZoneName),
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

func TestAccIBMISInstancesDataSource_InsGroupfilter(t *testing.T) {

	randInt := acctest.RandIntRange(10, 100)
	instanceGroupName := fmt.Sprintf("testinstancegroup%d", randInt)
	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDVtuCfWKVGKaRmaRG6JQZY8YdxnDgGzVOK93IrV9R5Hl0JP1oiLLWlZQS2reAKb8lBqyDVEREpaoRUDjqDqXG8J/kR42FKN51su914pjSBc86wJ02VtT1Wm1zRbSg67kT+g8/T1jCgB5XBODqbcICHVP8Z1lXkgbiHLwlUrbz6OZkGJHo/M/kD1Eme8lctceIYNz/Ilm7ewMXZA4fsidpto9AjyarrJLufrOBl4MRVcZTDSJ7rLP982aHpu9pi5eJAjOZc7Og7n4ns3NFppiCwgVMCVUQbN5GBlWhZ1OsT84ZiTf+Zy8ew+Yg5T7Il8HuC7loWnz+esQPf0s3xhC/kTsGgZreIDoh/rxJfD67wKXetNSh5RH/n5BqjaOuXPFeNXmMhKlhj9nJ8scayx/wsvOGuocEIkbyJSLj3sLUU403OafgatEdnJOwbqg6rUNNF5RIjpJpL7eEWlKIi1j9LyhmPJ+fEO7TmOES82VpCMHpLbe4gf/MhhJ/Xy8DKh9s= root@ffd8363b1226
	`)
	vpcName := fmt.Sprintf("testvpc%d", randInt)
	subnetName := fmt.Sprintf("testsubnet%d", randInt)
	templateName := fmt.Sprintf("testtemplate%d", randInt)
	sshKeyName := fmt.Sprintf("testsshkey%d", randInt)
	resName := "data.ibm_is_instances.ds_instances1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceGroupConfig(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance_group.instance_group", "name", instanceGroupName),
				),
			},
			{
				Config: testAccCheckIBMISInstancesDataSourceConfigInstanceGroup(instanceGroupName),
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

func testAccCheckIBMISInstancesDataSourceReservedIpConfig(vpcname string) string {
	return fmt.Sprintf(`
	data "ibm_is_instances" "ds_instances" {
		vpc_name = "%s"
	}`, vpcname)
}

func testAccCheckIBMISInstancesDataSourceConfig1(vpcname string) string {
	return fmt.Sprintf(`
	data "ibm_is_instances" "ds_instances1" {
		vpc_name = "%s"
	}`, vpcname)
}
func testAccCheckIBMISInstancesDataSourceConfigInstanceGroup(insGrpName string) string {
	return fmt.Sprintf(`
	data "ibm_is_instances" "ds_instances1" {
		instance_group_name = "%s"
	}`, insGrpName)
}
