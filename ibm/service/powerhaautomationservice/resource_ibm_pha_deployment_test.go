// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package powerhaautomationservice_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/powerhaautomationservice"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/DRAutomation/dra-go-sdk/powerhaautomationservicev1"
)

func TestAccIBMPhaDeploymentBasic(t *testing.T) {
	var conf powerhaautomationservicev1.PhaDeploymentResponse
	phaInstanceID := "2cfb7a06-623b-4eb9-a9ac-daa03dc0b5a6"
	primaryWorkspace := "9aa63e8a-6cd8-4998-95c0-d2bf121f3010"
	location_id := "us-south"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPhaDeploymentConfigBasic(phaInstanceID, primaryWorkspace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPhaDeploymentExists("ibm_pha_deployment.pha_deployment_instance", conf),
					resource.TestCheckResourceAttr("ibm_pha_deployment.pha_deployment_instance", "pha_instance_id", phaInstanceID),
					resource.TestCheckResourceAttr("ibm_pha_deployment.pha_deployment_instance", "primary_workspace", primaryWorkspace),
					resource.TestCheckResourceAttr("ibm_pha_deployment.pha_deployment_instance", "location_id", location_id),
				),
			},
		},
	})
}

func TestAccIBMPhaDeploymentAllArgs(t *testing.T) {
	var conf powerhaautomationservicev1.PhaDeploymentResponse
	phaInstanceID := "2cfb7a06-623b-4eb9-a9ac-daa03dc0b5a6"
	acceptLanguage := "en"
	ifNoneMatch := fmt.Sprintf("tf_if_none_match_%d", acctest.RandIntRange(10, 100))
	primaryLocation := "us-south"
	primaryWorkspace := "9aa63e8a-6cd8-4998-95c0-d2bf121f3010"
	location_id := "us-south"
	secondaryLocation := "eu-de-1"
	secondaryWorkspace := fmt.Sprintf("tf_secondary_workspace_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPhaDeploymentConfig(phaInstanceID, acceptLanguage, ifNoneMatch, primaryLocation, primaryWorkspace, secondaryLocation, secondaryWorkspace),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPhaDeploymentExists("ibm_pha_deployment.pha_deployment_instance", conf),
					resource.TestCheckResourceAttr("ibm_pha_deployment.pha_deployment_instance", "pha_instance_id", phaInstanceID),
					resource.TestCheckResourceAttr("ibm_pha_deployment.pha_deployment_instance", "accept_language", acceptLanguage),
					resource.TestCheckResourceAttr("ibm_pha_deployment.pha_deployment_instance", "if_none_match", ifNoneMatch),
					resource.TestCheckResourceAttr("ibm_pha_deployment.pha_deployment_instance", "primary_location", primaryLocation),
					resource.TestCheckResourceAttr("ibm_pha_deployment.pha_deployment_instance", "primary_workspace", primaryWorkspace),
					resource.TestCheckResourceAttr("ibm_pha_deployment.pha_deployment_instance", "secondary_location", secondaryLocation),
					resource.TestCheckResourceAttr("ibm_pha_deployment.pha_deployment_instance", "secondary_workspace", secondaryWorkspace),
					resource.TestCheckResourceAttr("ibm_pha_deployment.pha_deployment_instance", "secondary_workspace", secondaryWorkspace),
					resource.TestCheckResourceAttr("ibm_pha_deployment.pha_deployment_instance", "location_id", location_id),
				),
			},
			// resource.TestStep{
			// 	ResourceName:      "ibm_pha_deployment.pha_deployment_instance",
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// },
		},
	})
}

func testAccCheckIBMPhaDeploymentConfigBasic(phaInstanceID string, primaryWorkspace string) string {
	return fmt.Sprintf(`
		resource "ibm_pha_deployment" "pha_deployment_instance" {
			pha_instance_id = "%s"
			location_id = "us-south"
			primary_workspace = "%s"
		}
	`, phaInstanceID, primaryWorkspace)
}

func testAccCheckIBMPhaDeploymentConfig(phaInstanceID string, acceptLanguage string, ifNoneMatch string, primaryLocation string, primaryWorkspace string, secondaryLocation string, secondaryWorkspace string) string {
	return fmt.Sprintf(`

		resource "ibm_pha_deployment" "pha_deployment_instance" {
			pha_instance_id = "%s"
			accept_language = "%s"
			location_id= "us-south"
			if_none_match = "%s"
			primary_location = "%s"
			primary_workspace = "%s"
			secondary_location = "%s"
			secondary_workspace = "%s"
		}
	`, phaInstanceID, acceptLanguage, ifNoneMatch, primaryLocation, primaryWorkspace, secondaryLocation, secondaryWorkspace)
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

		// parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		// if err != nil {
		// 	return err
		// }

		// getPhaDeploymentOptions.SetPhaInstanceID(parts[0])
		getPhaDeploymentOptions.SetPhaInstanceID(rs.Primary.ID)

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

		// parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		// if err != nil {
		// 	return err
		// }

		// getPhaDeploymentOptions.SetPhaInstanceID(parts[0])
		getPhaDeploymentOptions.SetPhaInstanceID(rs.Primary.ID)

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
		model["agent_status"] = "RUNNING"
		model["cores"] = float64(8.0)
		model["ip_address"] = "10.0.2.45"
		model["memory"] = int(32)
		model["pha_level"] = "7.2.1"
		model["region"] = "us-south"
		model["vm_id"] = "vm-3c91af27"
		model["vm_name"] = "pha-node-01"
		model["vm_status"] = "ACTIVE"
		model["workspace_id"] = "workspace-pha-prod"

		assert.Equal(t, result, model)
	}

	model := new(powerhaautomationservicev1.ClusterNodeInfo)
	model.AgentStatus = core.StringPtr("RUNNING")
	model.Cores = core.Float32Ptr(float32(8.0))
	model.IPAddress = core.StringPtr("10.0.2.45")
	model.Memory = core.Int64Ptr(int64(32))
	model.PhaLevel = core.StringPtr("7.2.1")
	model.Region = core.StringPtr("us-south")
	model.VMID = core.StringPtr("vm-3c91af27")
	model.VMName = core.StringPtr("pha-node-01")
	model.VMStatus = core.StringPtr("ACTIVE")
	model.WorkspaceID = core.StringPtr("workspace-pha-prod")

	result, err := powerhaautomationservice.ResourceIBMPhaDeploymentClusterNodeInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
