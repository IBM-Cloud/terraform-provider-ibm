package ibm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMISInstanceGroupManager_basic(t *testing.T) {
	randInt := acctest.RandIntRange(10, 100)
	instanceGroupName := fmt.Sprintf("testinstancegroup%d", randInt)
	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
	`)
	vpcName := fmt.Sprintf("testvpc%d", randInt)
	subnetName := fmt.Sprintf("testsubnet%d", randInt)
	templateName := fmt.Sprintf("testtemplate%d", randInt)
	sshKeyName := fmt.Sprintf("testsshkey%d", randInt)
	instanceGroupManager := fmt.Sprintf("testinstancegroupmanager%d", randInt)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceGroupManagerConfig(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName, instanceGroupManager),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance_group_manager.instance_group_manager", "name", instanceGroupManager),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_group_manager.instance_group_manager", "max_membership_count", "2"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_group_manager.instance_group_manager", "min_membership_count", "1"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_group_manager.instance_group_manager", "aggregation_window", "120"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceGroupManagerDestroy(s *terraform.State) error {
	sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_instance_group_manager" {
			continue
		}
		instanceGroup := rs.Primary.Attributes["instance_group"]
		getInstanceGroupManagerOptions := vpcv1.GetInstanceGroupManagerOptions{
			ID:              &rs.Primary.ID,
			InstanceGroupID: &instanceGroup,
		}
		_, _, err := sess.GetInstanceGroupManager(&getInstanceGroupManagerOptions)
		if err == nil {
			return fmt.Errorf("instance group manager still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckIBMISInstanceGroupManagerConfig(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName, instanceGroupManager string) string {
	return fmt.Sprintf(`
	provider "ibm" {
		generation = 2
	}
	
	resource "ibm_is_vpc" "vpc2" {
	  name = "%s"
	}
	
	resource "ibm_is_subnet" "subnet2" {
	  name            = "%s"
	  vpc             = ibm_is_vpc.vpc2.id
	  zone            = "us-south-2"
	  ipv4_cidr_block = "10.240.64.0/28"
	}
	
	resource "ibm_is_ssh_key" "sshkey" {
	  name       = "%s"
	  public_key = "%s"
	}
	
	resource "ibm_is_instance_template" "instancetemplate1" {
	   name    = "%s"
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

	`, vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName, instanceGroupManager)

}
