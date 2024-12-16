// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMISInstanceGroupManagerPolicies_dataBasic(t *testing.T) {
	randInt := acctest.RandIntRange(1400, 1500)
	instanceGroupName := fmt.Sprintf("testinstancegroup%d", randInt)
	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQC7xU4ahAKfw3NakcA44MAwQ57Ztoz2+Y1gbLGoqdQVWR4F+CWnm8pG5SPz1fc23LxWH/mFL7JJDZ7aJ0IXk0SkP7hBo3aduUvfNvR2v9og8gGh8iKsRaxiORwSHhfcr0k4GaGBUNr1gXpzJnEGdKkqOm7SNnzb9kYHpN3y2DRJscv4GMj4fV4qFD9TKNd1N65fWhVRPwyMV3uXzbnDjAobbgglXB/o96Xi4WoRAHTBHiZy3SOCUHw7vEOzSTLWlB2dnwn7FE+zAvvedsX1hm0U1E5tIUP+1O2kYFeAaHdI8ZYabdYp+3fZXJdsOxfePZKRrvNsQjZA606kngjKzlhftxOUdxD2CLk1OlS9FyFrMJL9eCRzYKfBSjAv8xWibzYB8H5LtqUnCCW+wVa8dq4YJFgNg1h2GGK+K375+xioGrfvtrOAa528V/WbGztmve7eRmFxca5oBu2EHAe2GsKemGXzHu/RmlAoP49ebv+i+c5ljflPvaWMtlw7RaGI5Ik= root@ffd8363b1226
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
				Config: testAccCheckIBMISInstanceGroupManagerPoliciesConfigd(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName, instanceGroupManager, instanceGroupManagerPolicy),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_group_manager_policies.instance_group_manager_policy", "instance_group"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_group_manager_policies.instance_group_manager_policy", "instance_group_manager"),
					resource.TestCheckResourceAttr(
						"data.ibm_is_instance_group_manager_policies.instance_group_manager_policy", "instance_group_manager_policies.#", "1"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_group_manager_policies.instance_group_manager_policy", "instance_group_manager_policies.0.name"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_group_manager_policies.instance_group_manager_policy", "instance_group_manager_policies.0.metric_type"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_group_manager_policies.instance_group_manager_policy", "instance_group_manager_policies.0.metric_value"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_group_manager_policies.instance_group_manager_policy", "instance_group_manager_policies.0.policy_type"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceGroupManagerPoliciesConfigd(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName, instanceGroupManager, instanceGroupManagerPolicy string) string {
	return testAccCheckIBMISInstanceGroupManagerPolicyConfig(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName, instanceGroupManager, instanceGroupManagerPolicy) + fmt.Sprintf(`

	data "ibm_is_instance_group_manager_policies" "instance_group_manager_policy" {
		instance_group = ibm_is_instance_group_manager_policy.cpuPolicy.instance_group
		instance_group_manager = ibm_is_instance_group_manager_policy.cpuPolicy.instance_group_manager
	}
	`)

}
