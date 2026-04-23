// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.113.1-d76630af-20260320-135953
 */

package powerhaautomationservice_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/powerhaautomationservice"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/DRAutomation/dra-go-sdk/powerhaautomationservicev1"
)

func TestAccIBMPhaDeploymentDataSourceBasic(t *testing.T) {
	phaDeploymentResponsePhaInstanceID := "748943bf-7965-492a-9ce0-88309cf18451"
	phaDeploymentResponsePrimaryWorkspace := "xxxx3e8a-xxxx-xxxx-xxxx-xxxxxf3010"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPhaDeploymentDataSourceConfigBasic(phaDeploymentResponsePhaInstanceID, phaDeploymentResponsePrimaryWorkspace),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "custom_network.#"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "secondary_cluster_nodes.#"),
				),
			},
		},
	})
}

func TestAccIBMPhaDeploymentDataSourceAllArgs(t *testing.T) {
	phaDeploymentResponsePhaInstanceID := "748943bf-7965-492a-9ce0-88309cf18451"
	phaDeploymentResponseAcceptLanguage := "en"
	phaDeploymentResponseIfNoneMatch := ""
	phaDeploymentResponsePrimaryLocation := "us-south"
	phaDeploymentResponsePrimaryWorkspace := "xxxx3e8a-xxxx-xxxx-xxxx-xxxxxf3010"
	phaDeploymentResponseSecondaryLocation := ""
	phaDeploymentResponseSecondaryWorkspace := ""

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPhaDeploymentDataSourceConfig(phaDeploymentResponsePhaInstanceID, phaDeploymentResponseAcceptLanguage, phaDeploymentResponseIfNoneMatch, phaDeploymentResponsePrimaryLocation, phaDeploymentResponsePrimaryWorkspace, phaDeploymentResponseSecondaryLocation, phaDeploymentResponseSecondaryWorkspace),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "if_none_match"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "cloud_account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "connectivity_type"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "creation_time"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "custom_network.#"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "deprovision_time"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "guid"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "is_duplicate"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "plan_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "plan_name"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "powerha_cluster_name"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "powerha_cluster_type"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "powerha_level"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "primary_cluster_nodes_details.#"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "primary_cluster_nodes_details.0.agent_status"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "primary_cluster_nodes_details.0.cores"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "primary_cluster_nodes_details.0.ip_address"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "primary_cluster_nodes_details.0.memory"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "primary_cluster_nodes_details.0.pha_level"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "primary_cluster_nodes_details.0.region"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "primary_cluster_nodes_details.0.vm_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "primary_cluster_nodes_details.0.vm_name"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "primary_cluster_nodes_details.0.vm_status"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "primary_cluster_nodes_details.0.workspace_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "primary_location"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "primary_region_name"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "primary_workspace"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "primary_workspace_name"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "provision_end_time"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "provision_start_time"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "provision_status"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "region_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "resource_group_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "resource_instance"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "secondary_cluster_nodes.#"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "secondary_cluster_nodes.0.agent_status"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "secondary_cluster_nodes.0.cores"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "secondary_cluster_nodes.0.ip_address"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "secondary_cluster_nodes.0.memory"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "secondary_cluster_nodes.0.pha_level"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "secondary_cluster_nodes.0.region"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "secondary_cluster_nodes.0.vm_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "secondary_cluster_nodes.0.vm_name"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "secondary_cluster_nodes.0.vm_status"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "secondary_cluster_nodes.0.workspace_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "secondary_location"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "secondary_workspace"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "service_description"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "service_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "service_name"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "standby_region_name"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "standby_workspace_name"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "user_tags"),
				),
			},
		},
	})
}

func testAccCheckIBMPhaDeploymentDataSourceConfigBasic(phaDeploymentResponsePhaInstanceID string, phaDeploymentResponsePrimaryWorkspace string) string {
	return fmt.Sprintf(`
		resource "ibm_pha_deployment" "pha_deployment_instance" {
			instance_id = "%s"
			primary_workspace = "%s"
		}

		data "ibm_pha_deployment" "pha_deployment_instance" {
			instance_id = "748943bf-7965-492a-9ce0-88309cf18451"
		}
	`, phaDeploymentResponsePhaInstanceID, phaDeploymentResponsePrimaryWorkspace)
}

func testAccCheckIBMPhaDeploymentDataSourceConfig(phaDeploymentResponsePhaInstanceID string, phaDeploymentResponseAcceptLanguage string, phaDeploymentResponseIfNoneMatch string, phaDeploymentResponsePrimaryLocation string, phaDeploymentResponsePrimaryWorkspace string, phaDeploymentResponseSecondaryLocation string, phaDeploymentResponseSecondaryWorkspace string) string {
	return fmt.Sprintf(`
		resource "ibm_pha_deployment" "pha_deployment_instance" {
			instance_id = "%s"
			accept_language = "%s"
			if_none_match = "%s"
			primary_location = "%s"
			primary_workspace = "%s"
			secondary_location = "%s"
			secondary_workspace = "%s"
		}

		data "ibm_pha_deployment" "pha_deployment_instance" {
			instance_id = "748943bf-7965-492a-9ce0-88309cf18451"
		}
	`, phaDeploymentResponsePhaInstanceID, phaDeploymentResponseAcceptLanguage, phaDeploymentResponseIfNoneMatch, phaDeploymentResponsePrimaryLocation, phaDeploymentResponsePrimaryWorkspace, phaDeploymentResponseSecondaryLocation, phaDeploymentResponseSecondaryWorkspace)
}

func TestDataSourceIBMPhaDeploymentClusterNodeInfoToMap(t *testing.T) {
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

	result, err := powerhaautomationservice.DataSourceIBMPhaDeploymentClusterNodeInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
