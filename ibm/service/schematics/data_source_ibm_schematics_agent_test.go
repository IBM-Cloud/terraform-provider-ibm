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

func TestAccIbmSchematicsAgentDataSourceBasic(t *testing.T) {
	agentDataName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	agentDataVersion := "1.0.0"
	agentDataSchematicsLocation := "us-south"
	agentDataAgentLocation := "eu-de"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSchematicsAgentDataSourceConfigBasic(agentDataName, agentDataVersion, agentDataSchematicsLocation, agentDataAgentLocation),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent.schematics_agent_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent.schematics_agent_instance", "agent_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent.schematics_agent_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent.schematics_agent_instance", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent.schematics_agent_instance", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent.schematics_agent_instance", "schematics_location"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent.schematics_agent_instance", "agent_location"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent.schematics_agent_instance", "agent_infrastructure.#"),
				),
			},
		},
	})
}

func TestAccIbmSchematicsAgentDataSourceAllArgs(t *testing.T) {
	agentDataName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	agentDataVersion := "1.0.0"
	agentDataSchematicsLocation := "us-south"
	agentDataAgentLocation := "eu-de"
	agentDataDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSchematicsAgentDataSourceConfig(agentDataName, agentDataVersion, agentDataSchematicsLocation, agentDataAgentLocation, agentDataDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent.schematics_agent_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent.schematics_agent_instance", "agent_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent.schematics_agent_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent.schematics_agent_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent.schematics_agent_instance", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent.schematics_agent_instance", "tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent.schematics_agent_instance", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent.schematics_agent_instance", "schematics_location"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent.schematics_agent_instance", "agent_location"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent.schematics_agent_instance", "agent_infrastructure.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent.schematics_agent_instance", "agent_metadata.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent.schematics_agent_instance", "user_state.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent.schematics_agent_instance", "agent_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent.schematics_agent_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent.schematics_agent_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent.schematics_agent_instance", "creation_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent.schematics_agent_instance", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent.schematics_agent_instance", "system_state.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent.schematics_agent_instance", "agent_kpi.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent.schematics_agent_instance", "recent_prs_job.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent.schematics_agent_instance", "recent_deploy_job.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_agent.schematics_agent_instance", "recent_health_job.#"),
				),
			},
		},
	})
}

func testAccCheckIbmSchematicsAgentDataSourceConfigBasic(agentDataName string, agentDataVersion string, agentDataSchematicsLocation string, agentDataAgentLocation string) string {
	return fmt.Sprintf(`
		resource "ibm_schematics_agent" "schematics_agent_instance" {
			name = "%s"
			resource_group = "Default"
			version = "%s"
			schematics_location = "%s"
			agent_location = "%s"
			agent_infrastructure {
				infra_type = "ibm_kubernetes"
				cluster_id = "cluster_id"
				cluster_resource_group = "cluster_resource_group"
				cos_instance_name = "cos_instance_name"
				cos_bucket_name = "cos_bucket_name"
				cos_bucket_region = "cos_bucket_region"
			}
		}

		data "ibm_schematics_agent" "schematics_agent_instance" {
			agent_id = ibm_schematics_agent.schematics_agent_instance.id
		}
	`, agentDataName, agentDataVersion, agentDataSchematicsLocation, agentDataAgentLocation)
}

func testAccCheckIbmSchematicsAgentDataSourceConfig(agentDataName string, agentDataVersion string, agentDataSchematicsLocation string, agentDataAgentLocation string, agentDataDescription string) string {
	return fmt.Sprintf(`
		resource "ibm_schematics_agent" "schematics_agent_instance" {
			name = "%s"
			resource_group = "Default"
			version = "%s"
			schematics_location = "%s"
			agent_location = "%s"
			agent_infrastructure {
				infra_type = "ibm_kubernetes"
				cluster_id = "cluster_id"
				cluster_resource_group = "cluster_resource_group"
				cos_instance_name = "cos_instance_name"
				cos_bucket_name = "cos_bucket_name"
				cos_bucket_region = "cos_bucket_region"
			}
			description = "%s"
			tags = ["agent-tag"]
			agent_metadata {
				name = "purpose"
				value = ["git", "terraform", "ansible"]
			}
			user_state {
				state = "enable"
			}
		}

		data "ibm_schematics_agent" "schematics_agent_instance" {
			agent_id = ibm_schematics_agent.schematics_agent_instance.id
		}
	`, agentDataName, agentDataVersion, agentDataSchematicsLocation, agentDataAgentLocation, agentDataDescription)
}
