package ibm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMISInstanceGroupManager_dataBasic(t *testing.T) {
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
				Config: testAccCheckIBMISInstanceGroupManagerDConfig(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName, instanceGroupManager),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_group_manager.instance_group_manager", "name"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_group_manager.instance_group_manager", "max_membership_count"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_group_manager.instance_group_manager", "min_membership_count"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_group_manager.instance_group_manager", "aggregation_window"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceGroupManagerDConfig(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName, instanceGroupManager string) string {
	return testAccCheckIBMISInstanceGroupManagerConfig(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName, instanceGroupManager) + fmt.Sprintf(`

	data "ibm_is_instance_group_manager" "instance_group_manager" {
		instance_group = ibm_is_instance_group_manager.instance_group_manager.instance_group
		name = "%s"
	}

	`, instanceGroupManager)

}
