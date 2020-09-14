package ibm

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMISInstanceGroup_dataBasic(t *testing.T) {
	randInt := acctest.RandIntRange(10, 100)
	instanceGroupName := fmt.Sprintf("testinstancegroup%d", randInt)
	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCKVmnMOlHKcZK8tpt3MP1lqOLAcqcJzhsvJcjscgVERRN7/9484SOBJ3HSKxxNG5JN8owAjy5f9yYwcUg+JaUVuytn5Pv3aeYROHGGg+5G346xaq3DAwX6Y5ykr2fvjObgncQBnuU5KHWCECO/4h8uWuwh/kfniXPVjFToc+gnkqA+3RKpAecZhFXwfalQ9mMuYGFxn+fwn8cYEApsJbsEmb0iJwPiZ5hjFC8wREuiTlhPHDgkBLOiycd20op2nXzDbHfCHInquEe/gYxEitALONxm0swBOwJZwlTDOB7C6y2dzlrtxr1L59m7pCkWI4EtTRLvleehBoj3u7jB4usR
	`)
	vpcName := fmt.Sprintf("testvpc%d", randInt)
	subnetName := fmt.Sprintf("testsubnet%d", randInt)
	templateName := fmt.Sprintf("testtemplate%d", randInt)
	sshKeyName := fmt.Sprintf("testsshkey%d", randInt)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceGroupDConfig(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ibm_is_instance_group.instance_group_data", "name", instanceGroupName),
					resource.TestCheckResourceAttr(
						"data.ibm_is_instance_group.instance_group_data", "status", "healthy"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceGroupDConfig(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName string) string {
	return testAccCheckIBMISInstanceGroupConfig(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName) + fmt.Sprintf(`

	data "ibm_is_instance_group" "instance_group_data" {
		name = ibm_is_instance_group.instance_group.name
	}
	`)
}
