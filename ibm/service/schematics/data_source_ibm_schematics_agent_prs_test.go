// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package schematics_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSchematicsAgentPrsDataSourceBasic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSchematicsAgentPrsDataSourceConfigBasic(acc.AgentID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent_prs.schematics_agent_prs_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent_prs.schematics_agent_prs_instance", "agent_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent_prs.schematics_agent_prs_instance", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent_prs.schematics_agent_prs_instance", "updated_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent_prs.schematics_agent_prs_instance", "agent_version"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent_prs.schematics_agent_prs_instance", "status_code"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent_prs.schematics_agent_prs_instance", "status_message"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent_prs.schematics_agent_prs_instance", "log_url"),
				),
			},
		},
	})
}

func testAccCheckIbmSchematicsAgentPrsDataSourceConfigBasic(agentPRSJobAgentID string) string {
	return fmt.Sprintf(`
		data "ibm_schematics_agent_prs" "schematics_agent_prs_instance" {
			agent_id = "%s"
		}
	`, agentPRSJobAgentID)
}
