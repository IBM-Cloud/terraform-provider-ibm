// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package schematics_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSchematicsAgentHealthDataSourceBasic(t *testing.T) {
	agentHealthJobAgentID := fmt.Sprintf("tf_agent_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSchematicsAgentHealthDataSourceConfigBasic(agentHealthJobAgentID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent_health.schematics_agent_health_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent_health.schematics_agent_health_instance", "agent_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent_health.schematics_agent_health_instance", "job_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent_health.schematics_agent_health_instance", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent_health.schematics_agent_health_instance", "updated_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent_health.schematics_agent_health_instance", "agent_version"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent_health.schematics_agent_health_instance", "status_code"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent_health.schematics_agent_health_instance", "status_message"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent_health.schematics_agent_health_instance", "log_url"),
				),
			},
		},
	})
}

func testAccCheckIbmSchematicsAgentHealthDataSourceConfigBasic(agentHealthJobAgentID string) string {
	return fmt.Sprintf(`
		data "ibm_schematics_agent_health" "schematics_agent_health_instance" {
			agent_id = "%s"
		}
	`, agentHealthJobAgentID)
}
