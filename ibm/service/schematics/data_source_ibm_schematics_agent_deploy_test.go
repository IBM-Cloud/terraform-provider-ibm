// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package schematics_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSchematicsAgentDeployDataSourceBasic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSchematicsAgentDeployDataSourceConfigBasic(acc.AgentID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent_deploy.schematics_agent_deploy_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent_deploy.schematics_agent_deploy_instance", "agent_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent_deploy.schematics_agent_deploy_instance", "job_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent_deploy.schematics_agent_deploy_instance", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent_deploy.schematics_agent_deploy_instance", "updated_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent_deploy.schematics_agent_deploy_instance", "is_redeployed"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent_deploy.schematics_agent_deploy_instance", "agent_version"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent_deploy.schematics_agent_deploy_instance", "status_code"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent_deploy.schematics_agent_deploy_instance", "status_message"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent_deploy.schematics_agent_deploy_instance", "log_url"),
				),
			},
		},
	})
}

func testAccCheckIbmSchematicsAgentDeployDataSourceConfigBasic(agentDeployJobAgentID string) string {
	return fmt.Sprintf(`
		data "ibm_schematics_agent_deploy" "schematics_agent_deploy_instance" {
			agent_id = "%s"
		}
	`, agentDeployJobAgentID)
}
