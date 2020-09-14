package ibm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMISInstanceGroupManagerPolicies_dataBasic(t *testing.T) {
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
	instanceGroupManagerPolicy := fmt.Sprintf("testinstancegroupmanagerpolicy%d", randInt)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
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
