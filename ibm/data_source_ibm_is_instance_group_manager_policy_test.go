package ibm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMISInstanceGroupManagerPolicy_dataBasic(t *testing.T) {
	instanceGroupName := fmt.Sprintf("testinstancegroup%d", acctest.RandIntRange(10, 100))
	instanceGroupManagerName := fmt.Sprintf("igmanager%d", acctest.RandIntRange(200, 300))
	instanceGroupManagerPolicy := fmt.Sprintf("igmanager%d", acctest.RandIntRange(400, 500))
	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
	`)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceGroupManagerPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceGroupManagerPolicyConfigd(instanceGroupManagerPolicy, instanceGroupManagerName, instanceGroupName, publicKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_group_manager_policy.instance_group_manager_policy", "instance_group"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_group_manager_policy.instance_group_manager_policy", "instance_group_manager"),
					resource.TestCheckResourceAttr(
						"data.ibm_is_instance_group_manager_policy.instance_group_manager_policy", "instance_group_manager_policies.#", "1"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_group_manager_policy.instance_group_manager_policy", "instance_group_manager_policies.0.name"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_group_manager_policy.instance_group_manager_policy", "instance_group_manager_policies.0.metric_type"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_group_manager_policy.instance_group_manager_policy", "instance_group_manager_policies.0.metric_value"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_group_manager_policy.instance_group_manager_policy", "instance_group_manager_policies.0.policy_type"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceGroupManagerPolicyConfigd(name, instancegroupManager, instancegroup, publicKey string) string {
	return fmt.Sprintf(`
	provider "ibm" {
		generation = 2
	}
	
	resource "ibm_is_vpc" "vpc2" {
	  name = "vpc2test"
	}
	
	resource "ibm_is_subnet" "subnet2" {
	  name            = "subnet2"
	  vpc             = ibm_is_vpc.vpc2.id
	  zone            = "us-south-2"
	  ipv4_cidr_block = "10.240.64.0/28"
	}
	
	resource "ibm_is_ssh_key" "sshkey" {
	  name       = "ssh1"
	  public_key = "%s"
	}
	
	resource "ibm_is_instance_template" "instancetemplate1" {
	   name    = "testtemplate"
	   image   = "r006-14140f94-fcc4-11e9-96e7-a72723715315"
	   profile = "bx2-8x32"
	
	   primary_network_interface {
		 subnet = ibm_is_subnet.subnet2.id
	   }
	
	   vpc       = ibm_is_vpc.vpc2.id
	   zone      = "us-south-2"
	   keys      = [ibm_is_ssh_key.sshkey.id]
	 }
	
	resource "ibm_is_instance_group" "instance_group" {
		name =  "%s"
		instance_template = ibm_is_instance_template.instancetemplate1.id
		instance_count =  2
		subnets = [ibm_is_subnet.subnet2.id]
	}

	resource "ibm_is_instance_group_manager" "instance_group_manager" {
		name = "%s"
		aggregation_window = 120
		instance_group = ibm_is_instance_group.instance_group.id
		cooldown = 300
		manager_type = "autoscale"
		enable_manager = true
		max_membership_count = 2
		min_membership_count = 1
	}
	
	resource "ibm_is_instance_group_manager_policy" "cpuPolicy" {
		instance_group = ibm_is_instance_group.instance_group.id
		instance_group_manager =  ibm_is_instance_group_manager.instance_group_manager.id
		metric_type = "cpu"
		metric_value = 70
		policy_type = "target"
		name = "%s"
	}

	data "ibm_is_instance_group_manager_policy" "instance_group_manager_policy" {
		instance_group = ibm_is_instance_group_manager_policy.cpuPolicy.instance_group
		instance_group_manager = ibm_is_instance_group_manager_policy.cpuPolicy.instance_group_manager
	}

	`, publicKey, instancegroup, instancegroupManager, name)

}
