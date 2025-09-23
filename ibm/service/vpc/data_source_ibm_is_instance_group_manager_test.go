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

func TestAccIBMISInstanceGroupManager_dataBasic(t *testing.T) {
	randInt := acctest.RandIntRange(1000, 1100)
	instanceGroupName := fmt.Sprintf("testinstancegroup%d", randInt)
	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCsylPlEMEUrnLu9XmWKdFlaFIkMr9ujSVlWTPmrD0cmdZ/TH+zFhufTvig2SnqzCtaaBz6RpJmK2lm/5HROn+bW+UKsqmr7TeEjkaStpR+34xm1eIsCRbOjDECLBD+8fHK/3ZKNvjhlz2JfTkF8U5JN1o8cvUmdBkT2Rai/uGxR2bR6oEIvLZw8CTZXvhimFJa3rWOj39arrSPhHMC9wAohO5igJRxSpvYUPrlJdVshmjxkoqYFaiyp/37DmQU16jxWfQ57ziSd1psZ+aWXlot0xz9gl8bRSMXoZxMylU9t7y05sw+KrrzoRfPvm7z9anhTnsni0yC0W/lReG5xGgkcJHg7X8nei4SHDlWCXodA5PzMUT6AEMKHbAM3SNO3pZ1sPFbwTuO1iOyXUemLwgg0ECv4Z2loaSxeH/ryu7yLw3R54azhh2eqawbyqEHfiqF5zAmLN2kGIVr7HuT15RlXBgfSRIFCruqEdqAWW32Mp8eDs0O8ZzfayBjmSW/wSE= root@ffd8363b1226
	`)
	vpcName := fmt.Sprintf("testvpc%d", randInt)
	subnetName := fmt.Sprintf("testsubnet%d", randInt)
	templateName := fmt.Sprintf("testtemplate%d", randInt)
	sshKeyName := fmt.Sprintf("testsshkey%d", randInt)
	instanceGroupManager := fmt.Sprintf("testinstancegroupmanager%d", randInt)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
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

func TestAccIBMISInstanceGroupManager_dataBasic_scheduled(t *testing.T) {
	randInt := acctest.RandIntRange(1000, 1100)
	instanceGroupName := fmt.Sprintf("testinstancegroup%d", randInt)
	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCsylPlEMEUrnLu9XmWKdFlaFIkMr9ujSVlWTPmrD0cmdZ/TH+zFhufTvig2SnqzCtaaBz6RpJmK2lm/5HROn+bW+UKsqmr7TeEjkaStpR+34xm1eIsCRbOjDECLBD+8fHK/3ZKNvjhlz2JfTkF8U5JN1o8cvUmdBkT2Rai/uGxR2bR6oEIvLZw8CTZXvhimFJa3rWOj39arrSPhHMC9wAohO5igJRxSpvYUPrlJdVshmjxkoqYFaiyp/37DmQU16jxWfQ57ziSd1psZ+aWXlot0xz9gl8bRSMXoZxMylU9t7y05sw+KrrzoRfPvm7z9anhTnsni0yC0W/lReG5xGgkcJHg7X8nei4SHDlWCXodA5PzMUT6AEMKHbAM3SNO3pZ1sPFbwTuO1iOyXUemLwgg0ECv4Z2loaSxeH/ryu7yLw3R54azhh2eqawbyqEHfiqF5zAmLN2kGIVr7HuT15RlXBgfSRIFCruqEdqAWW32Mp8eDs0O8ZzfayBjmSW/wSE= root@ffd8363b1226
	`)
	vpcName := fmt.Sprintf("testvpc%d", randInt)
	subnetName := fmt.Sprintf("testsubnet%d", randInt)
	templateName := fmt.Sprintf("testtemplate%d", randInt)
	sshKeyName := fmt.Sprintf("testsshkey%d", randInt)
	instanceGroupManager := fmt.Sprintf("testinstancegroupmanager%d", randInt)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISInstanceGroupManagerDConfigScheduled(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName, instanceGroupManager),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_group_manager.instance_group_manager", "name"),
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

func testAccCheckIBMISInstanceGroupManagerDConfigScheduled(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName, instanceGroupManager string) string {
	return testAccCheckIBMISInstanceGroupManagerConfigScheduled(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName, instanceGroupManager) + fmt.Sprintf(`

	data "ibm_is_instance_group_manager" "instance_group_manager" {
		instance_group = ibm_is_instance_group_manager.instance_group_manager.instance_group
		name = "%s"
	}

	`, instanceGroupManager)

}
