// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMISInstanceGroupManagerAction_basic(t *testing.T) {
	randInt := acctest.RandIntRange(200, 300)
	instanceGroupName := fmt.Sprintf("testinstancegroup%d", randInt)
	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDVtuCfWKVGKaRmaRG6JQZY8YdxnDgGzVOK93IrV9R5Hl0JP1oiLLWlZQS2reAKb8lBqyDVEREpaoRUDjqDqXG8J/kR42FKN51su914pjSBc86wJ02VtT1Wm1zRbSg67kT+g8/T1jCgB5XBODqbcICHVP8Z1lXkgbiHLwlUrbz6OZkGJHo/M/kD1Eme8lctceIYNz/Ilm7ewMXZA4fsidpto9AjyarrJLufrOBl4MRVcZTDSJ7rLP982aHpu9pi5eJAjOZc7Og7n4ns3NFppiCwgVMCVUQbN5GBlWhZ1OsT84ZiTf+Zy8ew+Yg5T7Il8HuC7loWnz+esQPf0s3xhC/kTsGgZreIDoh/rxJfD67wKXetNSh5RH/n5BqjaOuXPFeNXmMhKlhj9nJ8scayx/wsvOGuocEIkbyJSLj3sLUU403OafgatEdnJOwbqg6rUNNF5RIjpJpL7eEWlKIi1j9LyhmPJ+fEO7TmOES82VpCMHpLbe4gf/MhhJ/Xy8DKh9s= root@ffd8363b1226
	`)
	vpcName := fmt.Sprintf("testvpc%d", randInt)
	subnetName := fmt.Sprintf("testsubnet%d", randInt)
	templateName := fmt.Sprintf("testtemplate%d", randInt)
	sshKeyName := fmt.Sprintf("testsshkey%d", randInt)
	instanceGroupManager := fmt.Sprintf("testinstancegroupmanager%d", randInt)
	instanceGroupManagerAction := fmt.Sprintf("testinstancegroupmanageraction%d", randInt)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceGroupManagerActionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceGroupManagerActionConfig(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName, instanceGroupManager, instanceGroupManagerAction),
				Check: resource.ComposeTestCheckFunc(
					// resource.TestCheckResourceAttr(
					// 	"ibm_is_instance_group_manager.instance_group_manager", "name", instanceGroupManager),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_group_manager_action.instance_group_manager_action", "name", instanceGroupManagerAction),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceGroupManagerActionDestroy(s *terraform.State) error {
	sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {

		if rs.Type != "ibm_ibm_is_instance_group_manager_action" {
			continue
		}

		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		instanceGroupID := parts[0]
		instancegroupmanagerscheduledID := parts[1]
		instanceGroupManagerActionID := parts[2]

		getInstanceGroupManagerActionOptions := &vpcv1.GetInstanceGroupManagerActionOptions{
			InstanceGroupID:        &instanceGroupID,
			InstanceGroupManagerID: &instancegroupmanagerscheduledID,
			ID:                     &instanceGroupManagerActionID,
		}

		_, _, err = sess.GetInstanceGroupManagerAction(getInstanceGroupManagerActionOptions)

		if err == nil {
			return fmt.Errorf("ibm_is_instance_group_manager_action still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckIBMISInstanceGroupManagerActionConfig(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName, instanceGroupManager, instanceGroupManagerAction string) string {
	// return fmt.Sprintf(`
	// provider "ibm" {
	// 	generation = 2
	// }

	// resource "ibm_is_vpc" "vpc2" {
	//   name = "%s"
	// }

	// resource "ibm_is_subnet" "subnet2" {
	//   name            = "%s"
	//   vpc             = ibm_is_vpc.vpc2.id
	//   zone            = "us-south-2"
	//   ipv4_cidr_block = "10.240.64.0/28"
	// }

	// resource "ibm_is_ssh_key" "sshkey" {
	//   name       = "%s"
	//   public_key = "%s"
	// }

	// resource "ibm_is_instance_template" "instancetemplate1" {
	//    name    = "%s"
	//    image   = "r006-14140f94-fcc4-11e9-96e7-a72723715315"
	//    profile = "bx2-8x32"

	//    primary_network_interface {
	// 	 subnet = ibm_is_subnet.subnet2.id
	//    }

	//    vpc       = ibm_is_vpc.vpc2.id
	//    zone      = "us-south-2"
	//    keys      = [ibm_is_ssh_key.sshkey.id]
	//  }

	// resource "ibm_is_instance_group" "instance_group" {
	// 	name =  "%s"
	// 	instance_template = ibm_is_instance_template.instancetemplate1.id
	// 	instance_count =  2
	// 	subnets = [ibm_is_subnet.subnet2.id]
	// }

	// resource "ibm_is_instance_group_manager" "instance_group_manager" {
	// 	name = "%s"
	// 	instance_group = ibm_is_instance_group.instance_group.id
	// 	manager_type = "scheduled"
	// 	enable_manager = true
	// }

	// resource "ibm_is_instance_group_manager_action" "instance_group_manager_action" {
	// 	name = "%s"
	// 	instance_group = ibm_is_instance_group.instance_group.id
	// 	instance_group_manager_scheduled = ibm_is_instance_group_manager.instance_group_manager.manager_id
	// 	cron_spec = "*/5 1,2,3 * * *"
	// 	membership_count = 1
	// }

	// `, vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName, instanceGroupManager, instanceGroupManagerAction)

	return fmt.Sprintf(`
	resource "ibm_is_instance_group" "instance_group" {
		name =  "sunithainstancegroup"
		instance_template = "0726-b749453f-751c-4154-b3f7-03a361e2b072"
		instance_count =  2
		subnets = ["0726-bcc3cf54-c88b-4c60-b9c8-24d64d8460a2"]
	}

	resource "ibm_is_instance_group_manager" "instance_group_manager" {
		name =  "sunithainstancegroupmanagerscheduled"
		instance_group = ibm_is_instance_group.instance_group.id
		manager_type = "scheduled"
		enable_manager = true
	}

	resource "ibm_is_instance_group_manager_action" "instance_group_manager_action" {
		name = "%s"
		instance_group = ibm_is_instance_group.instance_group.id
		instance_group_manager_scheduled = ibm_is_instance_group_manager.instance_group_manager.manager_id
		run_at= "2022-01-01T12:00:00.000Z"
		membership_count = 1
	}

	`, instanceGroupManagerAction)

}
