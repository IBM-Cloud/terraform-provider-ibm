// Copyright IBM Corp. 2017, 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMCisBotManagement_Basic(t *testing.T) {
	name := "ibm_cis_bot_management." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisBotManagementBasic1("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "fight_mode", "false"),
					resource.TestCheckResourceAttr(name, "session_score", "false"),
					resource.TestCheckResourceAttr(name, "enable_js", "false"),
					resource.TestCheckResourceAttr(name, "auth_id_logging", "false"),
					resource.TestCheckResourceAttr(name, "use_latest_model", "false"),
				),
			},
		},
	})
}

func testAccCheckCisBotManagementBasic1(id string, CisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_bot_management" "%[1]s" {
		cis_id                    = data.ibm_cis.cis.id
		domain_id                 = data.ibm_cis_domain.cis_domain.domain
		fight_mode				= false
		session_score			= false
		enable_js				= false
		auth_id_logging			= false
		use_latest_model 		= false
	  }
`, id)
}
