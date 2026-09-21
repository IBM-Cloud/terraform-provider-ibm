// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package drautomationservice_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.ibm.com/DRAutomation/dra-go-sdk/drautomationservicev1"
)

func TestAccIBMPdrIBMMaintainedDeploymentBasic(t *testing.T) {
	var conf drautomationservicev1.IBMOrchestratorGetDeploymentResponse
	instanceID := fmt.Sprintf("tf_instance_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPdrIBMMaintainedDeploymentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPdrIBMMaintainedDeploymentConfigBasic(instanceID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPdrIBMMaintainedDeploymentExists("ibm_pdr_ibm_maintained_deployment.pdr_ibm_maintained_deployment_instance", conf),
					resource.TestCheckResourceAttr("ibm_pdr_ibm_maintained_deployment.pdr_ibm_maintained_deployment_instance", "instance_id", instanceID),
				),
			},
		},
	})
}

func TestAccIBMPdrIBMMaintainedDeploymentAllArgs(t *testing.T) {
	var conf drautomationservicev1.IBMOrchestratorGetDeploymentResponse
	instanceID := fmt.Sprintf("tf_instance_id_%d", acctest.RandIntRange(10, 100))
	standByRedeploy := fmt.Sprintf("tf_stand_by_redeploy_%d", acctest.RandIntRange(10, 100))
	acceptLanguage := fmt.Sprintf("tf_accept_language_%d", acctest.RandIntRange(10, 100))
	acceptsIncomplete := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPdrIBMMaintainedDeploymentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPdrIBMMaintainedDeploymentConfig(instanceID, standByRedeploy, acceptLanguage, acceptsIncomplete),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPdrIBMMaintainedDeploymentExists("ibm_pdr_ibm_maintained_deployment.pdr_ibm_maintained_deployment_instance", conf),
					resource.TestCheckResourceAttr("ibm_pdr_ibm_maintained_deployment.pdr_ibm_maintained_deployment_instance", "instance_id", instanceID),
					resource.TestCheckResourceAttr("ibm_pdr_ibm_maintained_deployment.pdr_ibm_maintained_deployment_instance", "stand_by_redeploy", standByRedeploy),
					resource.TestCheckResourceAttr("ibm_pdr_ibm_maintained_deployment.pdr_ibm_maintained_deployment_instance", "accept_language", acceptLanguage),
					resource.TestCheckResourceAttr("ibm_pdr_ibm_maintained_deployment.pdr_ibm_maintained_deployment_instance", "accepts_incomplete", acceptsIncomplete),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_pdr_ibm_maintained_deployment.pdr_ibm_maintained_deployment_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMPdrIBMMaintainedDeploymentConfigBasic(instanceID string) string {
	return fmt.Sprintf(`
		resource "ibm_pdr_ibm_maintained_deployment" "pdr_ibm_maintained_deployment_instance" {
			instance_id = "%s"
		}
	`, instanceID)
}

func testAccCheckIBMPdrIBMMaintainedDeploymentConfig(instanceID string, standByRedeploy string, acceptLanguage string, acceptsIncomplete string) string {
	return fmt.Sprintf(`

		resource "ibm_pdr_ibm_maintained_deployment" "pdr_ibm_maintained_deployment_instance" {
			instance_id = "%s"
			stand_by_redeploy = "%s"
			accept_language = "%s"
			accepts_incomplete = %s
		}
	`, instanceID, standByRedeploy, acceptLanguage, acceptsIncomplete)
}

func testAccCheckIBMPdrIBMMaintainedDeploymentExists(n string, obj drautomationservicev1.IBMOrchestratorGetDeploymentResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		drAutomationServiceClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).DrAutomationServiceV1()
		if err != nil {
			return err
		}

		getIBMMaintainedDeploymentOptions := &drautomationservicev1.GetIBMMaintainedDeploymentOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getIBMMaintainedDeploymentOptions.SetInstanceID(parts[0])
		getIBMMaintainedDeploymentOptions.SetInstanceID(parts[1])

		ibmOrchestratorResponse, _, err := drAutomationServiceClient.GetIBMMaintainedDeployment(getIBMMaintainedDeploymentOptions)
		if err != nil {
			return err
		}

		obj = *ibmOrchestratorResponse
		return nil
	}
}

func testAccCheckIBMPdrIBMMaintainedDeploymentDestroy(s *terraform.State) error {
	drAutomationServiceClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pdr_ibm_maintained_deployment" {
			continue
		}

		getIBMMaintainedDeploymentOptions := &drautomationservicev1.GetIBMMaintainedDeploymentOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getIBMMaintainedDeploymentOptions.SetInstanceID(parts[0])
		getIBMMaintainedDeploymentOptions.SetInstanceID(parts[1])

		// Try to find the key
		_, response, err := drAutomationServiceClient.GetIBMMaintainedDeployment(getIBMMaintainedDeploymentOptions)

		if err == nil {
			return fmt.Errorf("pdr_ibm_maintained_deployment still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for pdr_ibm_maintained_deployment (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
