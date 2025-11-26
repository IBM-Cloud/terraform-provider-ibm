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

func TestAccIBMISInstanceGroupManagers_dataBasic(t *testing.T) {
	randInt := acctest.RandIntRange(800, 900)
	instanceGroupName := fmt.Sprintf("testinstancegroup%d", randInt)
	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDEXuhV4aJJzeFYL7vfVWnEvXgpXBJs0RD1ITxQvNGXbu6gCvd08WWjN9lzCpSGqZvGy3ZJ1tgQekBLPSPOlpSkwskt//pvSncLkMiBXPq+cTHrI2QL1b0mScxvGBRllgzs1sBKN6EFWiSdVOGmo0z1oYh9GKAxly5+7yE7s9NCzTJ2JYB7wMfdti3FhFK6plqRnSxPQ/phjoPvvcfXCwNRe7CA+nLR3cyBXoFBHtP9SsfwCH+dNUbPy3q/TvOcWJoLgAd+Jt8NnuS4DItgUeu1pFWO/Jcw1j+vHY8PN3yOLi7MSH2AYFOkWqodI5s9d41sBHZQkVrADy0JGWXLaTWjYXSmF4vjPMYTRVSQZFojpQ2iQbzw2D9ITQEs1U+Zcbdx04PPjXMoNtsF5V3bzjAqRepHKHv1ld/ReXcbl9v71Bz29ppFLI5U6dMKl7YOoBBF2U5qGT2ASMRILKjosCjHaD0qp09qRJTyq+78+7bQUergG4PoAHB/B9iNboZUckU= root@ffd8363b1226
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
				Config: testAccCheckIBMISInstanceGroupManagersDConfig(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName, instanceGroupManager),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_group_managers.instance_group_manager", "instance_group_managers.0.name"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_group_managers.instance_group_manager", "instance_group_managers.0.max_membership_count"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_group_managers.instance_group_manager", "instance_group_managers.0.min_membership_count"),
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_group_managers.instance_group_manager", "instance_group_managers.0.aggregation_window"),
				),
			},
		},
	})
}

func TestAccIBMISInstanceGroupManagers_dataBasic_scheduled(t *testing.T) {
	randInt := acctest.RandIntRange(800, 900)
	instanceGroupName := fmt.Sprintf("testinstancegroup%d", randInt)
	publicKey := strings.TrimSpace(`
	ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDEXuhV4aJJzeFYL7vfVWnEvXgpXBJs0RD1ITxQvNGXbu6gCvd08WWjN9lzCpSGqZvGy3ZJ1tgQekBLPSPOlpSkwskt//pvSncLkMiBXPq+cTHrI2QL1b0mScxvGBRllgzs1sBKN6EFWiSdVOGmo0z1oYh9GKAxly5+7yE7s9NCzTJ2JYB7wMfdti3FhFK6plqRnSxPQ/phjoPvvcfXCwNRe7CA+nLR3cyBXoFBHtP9SsfwCH+dNUbPy3q/TvOcWJoLgAd+Jt8NnuS4DItgUeu1pFWO/Jcw1j+vHY8PN3yOLi7MSH2AYFOkWqodI5s9d41sBHZQkVrADy0JGWXLaTWjYXSmF4vjPMYTRVSQZFojpQ2iQbzw2D9ITQEs1U+Zcbdx04PPjXMoNtsF5V3bzjAqRepHKHv1ld/ReXcbl9v71Bz29ppFLI5U6dMKl7YOoBBF2U5qGT2ASMRILKjosCjHaD0qp09qRJTyq+78+7bQUergG4PoAHB/B9iNboZUckU= root@ffd8363b1226
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
				Config: testAccCheckIBMISInstanceGroupManagersDConfigScheduled(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName, instanceGroupManager),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.ibm_is_instance_group_managers.instance_group_manager", "instance_group_managers.0.name"),
				),
			},
		},
	})
}

func testAccCheckIBMISInstanceGroupManagersDConfig(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName, instanceGroupManager string) string {
	return testAccCheckIBMISInstanceGroupManagerConfig(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName, instanceGroupManager) + fmt.Sprintf(`

	data "ibm_is_instance_group_managers" "instance_group_manager" {
		instance_group = ibm_is_instance_group_manager.instance_group_manager.instance_group
	}

	`)

}

func testAccCheckIBMISInstanceGroupManagersDConfigScheduled(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName, instanceGroupManager string) string {
	return testAccCheckIBMISInstanceGroupManagerConfigScheduled(vpcName, subnetName, sshKeyName, publicKey, templateName, instanceGroupName, instanceGroupManager) + fmt.Sprintf(`

	data "ibm_is_instance_group_managers" "instance_group_manager" {
		instance_group = ibm_is_instance_group_manager.instance_group_manager.instance_group
	}

	`)

}
