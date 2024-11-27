// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cdtoolchain_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMCdToolchainsDataSourceBasic(t *testing.T) {
	tcName := fmt.Sprintf("tf_tc_ds_name_%d", acctest.RandIntRange(10, 100))
	rgName := acc.CdResourceGroupName
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCdToolchainsDataSourceConfigBasic(tcName, rgName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchains.cd_toolchains", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchains.cd_toolchains", "resource_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cd_toolchains.cd_toolchains", "toolchains.0.name"),
					resource.TestCheckResourceAttr("data.ibm_cd_toolchains.cd_toolchains", "toolchains.0.name", tcName),
				),
			},
		},
	})
}

func testAccCheckIBMCdToolchainsDataSourceConfigBasic(tcName string, rgName string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "resource_group" {
			name = "%s"
		}
	
		resource "ibm_cd_toolchain" "cd_toolchain" {
			name = "%s"
			resource_group_id = data.ibm_resource_group.resource_group.id
		}

		data "ibm_cd_toolchains" "cd_toolchains" {
			resource_group_id = data.ibm_resource_group.resource_group.id
			name = "%s"
			depends_on = [
				ibm_cd_toolchain.cd_toolchain
			]
		}
	`, rgName, tcName, tcName)
}
