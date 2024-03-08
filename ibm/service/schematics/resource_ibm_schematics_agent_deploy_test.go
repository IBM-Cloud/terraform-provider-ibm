// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package schematics_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/schematics-go-sdk/schematicsv1"
)

func TestAccIbmSchematicsAgentDeployBasic(t *testing.T) {
	var conf *schematicsv1.AgentDataRecentDeployJob
	agentID := fmt.Sprintf("tf_agent_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSchematicsAgentDeployDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSchematicsAgentDeployConfigBasic(agentID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSchematicsAgentDeployExists("ibm_schematics_agent_deploy.schematics_agent_deploy_instance", conf),
					resource.TestCheckResourceAttr("ibm_schematics_agent_deploy.schematics_agent_deploy_instance", "agent_id", agentID),
				),
			},
		},
	})
}

func testAccCheckIbmSchematicsAgentDeployConfigBasic(agentID string) string {
	return fmt.Sprintf(`

		resource "ibm_schematics_agent_deploy" "schematics_agent_deploy_instance" {
			agent_id = "%s"
		}
	`, agentID)
}
func testAccCheckIbmSchematicsAgentDeployExists(n string, obj *schematicsv1.AgentDataRecentDeployJob) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		schematicsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SchematicsV1()
		if err != nil {
			return err
		}

		getAgentDataOptions := &schematicsv1.GetAgentDataOptions{
			Profile: core.StringPtr("detailed"),
		}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getAgentDataOptions.SetAgentID(parts[0])

		agentData, _, err := schematicsClient.GetAgentData(getAgentDataOptions)
		if err != nil {
			return err
		}

		obj = agentData.RecentDeployJob
		return nil
	}
}

func testAccCheckIbmSchematicsAgentDeployDestroy(s *terraform.State) error {
	schematicsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SchematicsV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_schematics_agent_deploy" {
			continue
		}

		getAgentDataOptions := &schematicsv1.GetAgentDataOptions{
			Profile: core.StringPtr("detailed"),
		}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getAgentDataOptions.SetAgentID(parts[0])

		agent, response, err := schematicsClient.GetAgentData(getAgentDataOptions)

		if err == nil && agent.RecentDeployJob != nil {
			// Agent Deploy Job can never really truely be deleted
			return nil
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for schematics_agent (%s) has been destroyed: %s", parts[0], err)
		}
	}

	return nil
}
