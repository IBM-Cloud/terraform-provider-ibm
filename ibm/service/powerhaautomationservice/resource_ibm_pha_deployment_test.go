// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package powerhaautomationservice_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/powerhaautomationservice"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/DRAutomation/dra-go-sdk/powerhaautomationservicev1"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMPhaDeploymentBasic(t *testing.T) {
	var conf powerhaautomationservicev1.PhaDeploymentResponse
	phaInstanceID := fmt.Sprintf("tf_pha_instance_id_%d", acctest.RandIntRange(10, 100))
	primaryWorkspace := fmt.Sprintf("tf_primary_workspace_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPhaDeploymentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPhaDeploymentConfigBasic(phaInstanceID, primaryWorkspace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPhaDeploymentExists("ibm_pha_deployment.pha_deployment_instance", conf),
					resource.TestCheckResourceAttr("ibm_pha_deployment.pha_deployment_instance", "pha_instance_id", phaInstanceID),
					resource.TestCheckResourceAttr("ibm_pha_deployment.pha_deployment_instance", "primary_workspace", primaryWorkspace),
				),
			},
		},
	})
}

func TestAccIBMPhaDeploymentAllArgs(t *testing.T) {
	var conf powerhaautomationservicev1.PhaDeploymentResponse
	phaInstanceID := fmt.Sprintf("tf_pha_instance_id_%d", acctest.RandIntRange(10, 100))
	acceptLanguage := fmt.Sprintf("tf_accept_language_%d", acctest.RandIntRange(10, 100))
	ifNoneMatch := fmt.Sprintf("tf_if_none_match_%d", acctest.RandIntRange(10, 100))
	primaryWorkspace := fmt.Sprintf("tf_primary_workspace_%d", acctest.RandIntRange(10, 100))
	secondaryWorkspace := fmt.Sprintf("tf_secondary_workspace_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPhaDeploymentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPhaDeploymentConfig(phaInstanceID, acceptLanguage, ifNoneMatch, primaryWorkspace, secondaryWorkspace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPhaDeploymentExists("ibm_pha_deployment.pha_deployment_instance", conf),
					resource.TestCheckResourceAttr("ibm_pha_deployment.pha_deployment_instance", "pha_instance_id", phaInstanceID),
					resource.TestCheckResourceAttr("ibm_pha_deployment.pha_deployment_instance", "accept_language", acceptLanguage),
					resource.TestCheckResourceAttr("ibm_pha_deployment.pha_deployment_instance", "if_none_match", ifNoneMatch),
					resource.TestCheckResourceAttr("ibm_pha_deployment.pha_deployment_instance", "primary_workspace", primaryWorkspace),
					resource.TestCheckResourceAttr("ibm_pha_deployment.pha_deployment_instance", "secondary_workspace", secondaryWorkspace),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_pha_deployment.pha_deployment_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMPhaDeploymentConfigBasic(phaInstanceID string, primaryWorkspace string) string {
	return fmt.Sprintf(`
		resource "ibm_pha_deployment" "pha_deployment_instance" {
			pha_instance_id = "%s"
			primary_workspace = "%s"
		}
	`, phaInstanceID, primaryWorkspace)
}

func testAccCheckIBMPhaDeploymentConfig(phaInstanceID string, acceptLanguage string, ifNoneMatch string, primaryWorkspace string, secondaryWorkspace string) string {
	return fmt.Sprintf(`

		resource "ibm_pha_deployment" "pha_deployment_instance" {
			pha_instance_id = "%s"
			accept_language = "%s"
			if_none_match = "%s"
			primary_workspace = "%s"
			secondary_workspace = "%s"
		}
	`, phaInstanceID, acceptLanguage, ifNoneMatch, primaryWorkspace, secondaryWorkspace)
}

func testAccCheckIBMPhaDeploymentExists(n string, obj powerhaautomationservicev1.PhaDeploymentResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		powerhaAutomationServiceClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PowerhaAutomationServiceV1()
		if err != nil {
			return err
		}

		getPhaDeploymentOptions := &powerhaautomationservicev1.GetPhaDeploymentOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getPhaDeploymentOptions.SetPhaInstanceID(parts[0])
		getPhaDeploymentOptions.SetPhaInstanceID(parts[1])

		phaDeploymentResponse, _, err := powerhaAutomationServiceClient.GetPhaDeployment(getPhaDeploymentOptions)
		if err != nil {
			return err
		}

		obj = *phaDeploymentResponse
		return nil
	}
}

func testAccCheckIBMPhaDeploymentDestroy(s *terraform.State) error {
	powerhaAutomationServiceClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PowerhaAutomationServiceV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pha_deployment" {
			continue
		}

		getPhaDeploymentOptions := &powerhaautomationservicev1.GetPhaDeploymentOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getPhaDeploymentOptions.SetPhaInstanceID(parts[0])
		getPhaDeploymentOptions.SetPhaInstanceID(parts[1])

		// Try to find the key
		_, response, err := powerhaAutomationServiceClient.GetPhaDeployment(getPhaDeploymentOptions)

		if err == nil {
			return fmt.Errorf("pha_deployment still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for pha_deployment (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIBMPhaDeploymentClusterNodeInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["vm_name"] = "samplevm"
		model["vm_id"] = "123e4567-e89b-12d3-a456-426614174000"
		model["ip_address"] = "10.0.1.25"
		model["cores"] = float64(1.5)
		model["vm_status"] = "ACTIVE"
		model["memory"] = int(32768)
		model["region"] = "us-south"
		model["workspace_id"] = "9f3c2b1a-4d56-789e-a123-bcdef4567890"
		model["agent_status"] = "ACTIVE"
		model["pha_level"] = "7.2.1"

		assert.Equal(t, result, model)
	}

	model := new(powerhaautomationservicev1.ClusterNodeInfo)
	model.VMName = core.StringPtr("samplevm")
	model.VMID = core.StringPtr("123e4567-e89b-12d3-a456-426614174000")
	model.IPAddress = core.StringPtr("10.0.1.25")
	model.Cores = core.Float32Ptr(float32(1.5))
	model.VMStatus = core.StringPtr("ACTIVE")
	model.Memory = core.Int64Ptr(int64(32768))
	model.Region = core.StringPtr("us-south")
	model.WorkspaceID = core.StringPtr("9f3c2b1a-4d56-789e-a123-bcdef4567890")
	model.AgentStatus = core.StringPtr("ACTIVE")
	model.PhaLevel = core.StringPtr("7.2.1")

	result, err := powerhaautomationservice.ResourceIBMPhaDeploymentClusterNodeInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
