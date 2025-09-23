// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISInstanceGroupManagerAction_dataBasic(t *testing.T) {
	randInt := acctest.RandIntRange(1200, 1300)
	instanceGroupName := fmt.Sprintf("testinstancegroup%d", randInt)
	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDmdgitxJde3s6PFDmWbyoF8S5YsC8+l3Qo3vRzBzXf05n5b3JL0t1sswZ7XNNyO2y8jTX7sCpGMzv4Q3WksCzkU12OPbr89Zmf+mC+11o3Lp/NpejiNtYf8hVtHWUAUrKLNywjFnm28pn64pf9KFgdkkp9quBZQgis8osfeygknYaSBBzkZKaZPszGuixTqaRAaomfwDP7QJJvS3Bo8bAe2kK+4EsW2DfP7h1G6BhHoxjoinVshbfE1nsJ2zlQigidjyjFL5YbCUYygjz5kq2khoxWmaNNKPVxAZ8fqvIHNi8F8sLCKW6VTruxPQIlW2A/D1YIJ4ME/Y6Goje9l40dA1W/mnygD0mZVYiLtYtlUM6ylKoQKNGeV+ugA554UK0lA++FVg5xOm8SNSvWWf6hyN/mK6atbpSBzRLoUQc95XsG1u7eQtz/zA1+pKaVsASpCMMbZFcTaeOiPLIcIVlzYDBcKap0MglFlsJsoKSRJ4uxIGm+CHoCWdc7VgaVvz8= root@ffd8363b1226
	`)
	vpcName := fmt.Sprintf("testvpc%d", randInt)
	subnetName := fmt.Sprintf("testsubnet%d", randInt)
	templateName := fmt.Sprintf("testtemplate%d", randInt)
	sshKeyName := fmt.Sprintf("testsshkey%d", randInt)
	instanceGroupManager := fmt.Sprintf("testinstancegroupmanager%d", randInt)
	instanceGroupManageraction := fmt.Sprintf("testinstancegroupmanageraction%d", randInt)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISInstanceGroupManagerActionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceGroupManagerActionConfigd(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName, instanceGroupManager, instanceGroupManageraction),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_group_manager_action.instance_group_manager_action", "instance_group"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_group_manager_action.instance_group_manager_action", "instance_group_manager"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_group_manager_action.instance_group_manager_action", "name"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_group_manager_action.instance_group_manager_action", "action_id"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceGroupManagerActionConfigd(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName, instanceGroupManager, instanceGroupManageraction string) string {
	return testAccCheckIBMISInstanceGroupManagerActionConfig(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName, instanceGroupManager, instanceGroupManageraction) + fmt.Sprintf(`

	data "ibm_is_instance_group_manager_action" "instance_group_manager_action" {
		instance_group = ibm_is_instance_group_manager_action.instance_group_manager_action.instance_group
		instance_group_manager = ibm_is_instance_group_manager_action.instance_group_manager_action.instance_group_manager
		name = "%s"
	}
	`, instanceGroupManageraction)
}
