// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMISInstanceGroupManagerPolicy_basic(t *testing.T) {
	randInt := acctest.RandIntRange(400, 500)
	instanceGroupName := fmt.Sprintf("testinstancegroup%d", randInt)
	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC4zkmPqZ826/DpkkEIvA8VxUvJtSlP9cmAuHeofZiKczbvWbTeHBkBs2K4LKht/T53xKH8YTttmVX1AZqiHOzhi70jA7PvopbtfkcdTVxcJEJXJ6IhlTXGVcor/oreDTCn4o5KD3Y/TSAmIHi5s9+xZGfgPRijkBLCS98n0nNFqVQ2Uam8PrDkzFQox/2XsFCbrMFtjxCMo/c6DG/6Z3w/5mWi9Z4hH0kQqACaBJR6mYgM07LSmpyMu4qsrEjwQ9tKhz3EM0SOB9ueT+SFwvIeoq49j+6kYFAeZxMSUjfJ/jmrsAZS/cXsBYAekwroYr/SH0w+Mj96EnUX6IDW9YT6DqrVH91xbAaXqggwR7K5kM+WaDqxthcWYZseIsS7HNzsJKeyqEHwQy4pWAr5SHbREm+1YZ4fCGTpozNz8OKY+vizWxvbv4HJPZJtvV4X+7+rV+kkkUMh2eycWkqSjViGng0oT6wG5+FHnrRp2t4kMx+sL+/6vs2aSLEvDjTkltc= root@ffd8363b1226
	`)
	vpcName := fmt.Sprintf("testvpc%d", randInt)
	subnetName := fmt.Sprintf("testsubnet%d", randInt)
	templateName := fmt.Sprintf("testtemplate%d", randInt)
	sshKeyName := fmt.Sprintf("testsshkey%d", randInt)
	instanceGroupManager := fmt.Sprintf("testinstancegroupmanager%d", randInt)
	instanceGroupManagerPolicy := fmt.Sprintf("testinstancegroupmanagerpolicy%d", randInt)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceGroupManagerPolicyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceGroupManagerPolicyConfig(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName, instanceGroupManager, instanceGroupManagerPolicy),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"ibm_is_instance_group_manager_policy.cpuPolicy", "name", instanceGroupManagerPolicy),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_group_manager_policy.cpuPolicy", "metric_type", "cpu"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_group_manager_policy.cpuPolicy", "metric_value", "70"),
					resource.TestCheckResourceAttr(
						"ibm_is_instance_group_manager_policy.cpuPolicy", "policy_type", "target"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceGroupManagerPolicyDestroy(s *terraform.State) error {
	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_instance_group_manager_policy" {
			continue
		}
		instanceGroupID := rs.Primary.Attributes["instance_group"]
		instanceGroupManagerID := rs.Primary.Attributes["instance_group_manager"]

		getInstanceGroupManagerPolicyOptions := vpcv1.GetInstanceGroupManagerPolicyOptions{
			ID:                     &rs.Primary.ID,
			InstanceGroupID:        &instanceGroupID,
			InstanceGroupManagerID: &instanceGroupManagerID,
		}
		_, _, err := sess.GetInstanceGroupManagerPolicy(&getInstanceGroupManagerPolicyOptions)

		if err == nil {
			return fmt.Errorf("instance group manager policy still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckIBMISInstanceGroupManagerPolicyConfig(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName, instanceGroupManager, instanceGroupManagerPolicy string) string {
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
	
	resource "ibm_is_instance_group_manager_policy" "cpuPolicy" {
		instance_group = ibm_is_instance_group.instance_group.id
		instance_group_manager =  ibm_is_instance_group_manager.instance_group_manager.manager_id
		metric_type = "cpu"
		metric_value = 70
		policy_type = "target"
		name = "%s"
	}
	
	`, vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName, instanceGroupManager, instanceGroupManagerPolicy)

}
