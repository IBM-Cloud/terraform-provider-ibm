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
	"github.com/IBM/schematics-go-sdk/schematicsv1"
)

func TestAccIbmSchematicsAgentBasic(t *testing.T) {
	var conf schematicsv1.AgentData
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	version := "1.0.0-beta2"
	schematicsLocation := "us-south"
	agentLocation := "eu-de"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	versionUpdate := "1.0.0-beta2"
	schematicsLocationUpdate := "us-east"
	agentLocationUpdate := "eu-gb"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSchematicsAgentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSchematicsAgentConfigBasic(name, version, schematicsLocation, agentLocation),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSchematicsAgentExists("ibm_schematics_agent.schematics_agent_instance", conf),
					resource.TestCheckResourceAttr("ibm_schematics_agent.schematics_agent_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_schematics_agent.schematics_agent_instance", "version", version),
					resource.TestCheckResourceAttr("ibm_schematics_agent.schematics_agent_instance", "schematics_location", schematicsLocation),
					resource.TestCheckResourceAttr("ibm_schematics_agent.schematics_agent_instance", "agent_location", agentLocation),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmSchematicsAgentConfigBasic(nameUpdate, versionUpdate, schematicsLocationUpdate, agentLocationUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_schematics_agent.schematics_agent_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_agent.schematics_agent_instance", "version", versionUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_agent.schematics_agent_instance", "schematics_location", schematicsLocationUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_agent.schematics_agent_instance", "agent_location", agentLocationUpdate),
				),
			},
		},
	})
}

func TestAccIbmSchematicsAgentAllArgs(t *testing.T) {
	var conf schematicsv1.AgentData
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	version := "1.0.0-beta2"
	schematicsLocation := "us-south"
	agentLocation := "eu-de"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	versionUpdate := "1.0.0-beta2"
	schematicsLocationUpdate := "us-east"
	agentLocationUpdate := "eu-gb"
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSchematicsAgentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSchematicsAgentConfig(name, version, schematicsLocation, agentLocation, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSchematicsAgentExists("ibm_schematics_agent.schematics_agent_instance", conf),
					resource.TestCheckResourceAttr("ibm_schematics_agent.schematics_agent_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_schematics_agent.schematics_agent_instance", "version", version),
					resource.TestCheckResourceAttr("ibm_schematics_agent.schematics_agent_instance", "schematics_location", schematicsLocation),
					resource.TestCheckResourceAttr("ibm_schematics_agent.schematics_agent_instance", "agent_location", agentLocation),
					resource.TestCheckResourceAttr("ibm_schematics_agent.schematics_agent_instance", "description", description),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmSchematicsAgentConfig(nameUpdate, versionUpdate, schematicsLocationUpdate, agentLocationUpdate, descriptionUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_schematics_agent.schematics_agent_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_agent.schematics_agent_instance", "version", versionUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_agent.schematics_agent_instance", "schematics_location", schematicsLocationUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_agent.schematics_agent_instance", "agent_location", agentLocationUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_agent.schematics_agent_instance", "description", descriptionUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_schematics_agent.schematics_agent_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmSchematicsAgentConfigBasic(name string, version string, schematicsLocation string, agentLocation string) string {
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
	`, name, version, schematicsLocation, agentLocation)
}

func testAccCheckIbmSchematicsAgentConfig(name string, version string, schematicsLocation string, agentLocation string, description string) string {
	return fmt.Sprintf(`

		resource "ibm_schematics_agent" "schematics_agent_instance" {
			name = "%s"
			resource_group = "default"
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
			tags = "FIXME"
			agent_metadata {
				name = "purpose"
				value = ["git", "terraform", "ansible"]
			}
			agent_inputs {
				name = "name"
				value = "value"
				use_default = true
				metadata {
					type = "boolean"
					aliases = [ "aliases" ]
					description = "description"
					cloud_data_type = "cloud_data_type"
					default_value = "default_value"
					link_status = "normal"
					secure = true
					immutable = true
					hidden = true
					required = true
					options = [ "options" ]
					min_value = 1
					max_value = 1
					min_length = 1
					max_length = 1
					matches = "matches"
					position = 1
					group_by = "group_by"
					source = "source"
				}
				link = "link"
			}
			user_state {
				state = "enable"
				set_by = "set_by"
				set_at = "2021-01-31T09:44:12Z"
			}
			agent_kpi {
				availability_indicator = "available"
				lifecycle_indicator = "consistent"
				percent_usage_indicator = "percent_usage_indicator"
				application_indicators = [ null ]
				infra_indicators = [ null ]
			}
		}
	`, name, version, schematicsLocation, agentLocation, description)
}

func testAccCheckIbmSchematicsAgentExists(n string, obj schematicsv1.AgentData) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		schematicsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SchematicsV1()
		if err != nil {
			return err
		}

		getAgentDataOptions := &schematicsv1.GetAgentDataOptions{}

		getAgentDataOptions.SetAgentID(rs.Primary.ID)

		agentData, _, err := schematicsClient.GetAgentData(getAgentDataOptions)
		if err != nil {
			return err
		}

		obj = *agentData
		return nil
	}
}

func testAccCheckIbmSchematicsAgentDestroy(s *terraform.State) error {
	schematicsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SchematicsV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_schematics_agent" {
			continue
		}

		getAgentDataOptions := &schematicsv1.GetAgentDataOptions{}

		getAgentDataOptions.SetAgentID(rs.Primary.ID)

		// Try to find the key
		_, response, err := schematicsClient.GetAgentData(getAgentDataOptions)
		if err == nil {
			return nil
			// TODO: uncomment the following lines of code once the agent delete actually deletes agentdoc
			// return fmt.Errorf("schematics_agent still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for schematics_agent (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
