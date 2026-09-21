// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.116.0-df613dbc-20260803-154903
*/

package powerhaautomationservice_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/powerhaautomationservice"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/DRAutomation/dra-go-sdk/powerhaautomationservicev1"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMPhaDeploymentDataSourceBasic(t *testing.T) {
	phaDeploymentResponsePhaInstanceID := fmt.Sprintf("tf_pha_instance_id_%d", acctest.RandIntRange(10, 100))
	phaDeploymentResponsePrimaryWorkspace := fmt.Sprintf("tf_primary_workspace_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPhaDeploymentDataSourceConfigBasic(phaDeploymentResponsePhaInstanceID, phaDeploymentResponsePrimaryWorkspace),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "pha_instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "primary_cluster_nodes_details.#"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "secondary_cluster_nodes_details.#"),
				),
			},
		},
	})
}

func TestAccIBMPhaDeploymentDataSourceAllArgs(t *testing.T) {
	phaDeploymentResponsePhaInstanceID := fmt.Sprintf("tf_pha_instance_id_%d", acctest.RandIntRange(10, 100))
	phaDeploymentResponseAcceptLanguage := fmt.Sprintf("tf_accept_language_%d", acctest.RandIntRange(10, 100))
	phaDeploymentResponseIfNoneMatch := fmt.Sprintf("tf_if_none_match_%d", acctest.RandIntRange(10, 100))
	phaDeploymentResponsePrimaryWorkspace := fmt.Sprintf("tf_primary_workspace_%d", acctest.RandIntRange(10, 100))
	phaDeploymentResponseSecondaryWorkspace := fmt.Sprintf("tf_secondary_workspace_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPhaDeploymentDataSourceConfig(phaDeploymentResponsePhaInstanceID, phaDeploymentResponseAcceptLanguage, phaDeploymentResponseIfNoneMatch, phaDeploymentResponsePrimaryWorkspace, phaDeploymentResponseSecondaryWorkspace),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "pha_instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "if_none_match"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "primary_workspace_name"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "standby_workspace_name"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "primary_region_name"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "standby_region_name"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "powerha_cluster_type"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "powerha_cluster_name"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "resource_group_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "primary_workspace"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "secondary_workspace"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "service_description"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "resource_instance"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "region_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "guid"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "connectivity_type"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "is_duplicate"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "powerha_level"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "provision_status"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "plan_name"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "service_name"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "primary_location"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "secondary_location"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "cloud_account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "provision_start_time"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "provision_end_time"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "creation_time"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "deprovision_time"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "service_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "plan_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "user_tags"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "primary_cluster_nodes_details.#"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "primary_cluster_nodes_details.0.vm_name"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "primary_cluster_nodes_details.0.vm_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "primary_cluster_nodes_details.0.ip_address"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "primary_cluster_nodes_details.0.cores"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "primary_cluster_nodes_details.0.vm_status"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "primary_cluster_nodes_details.0.memory"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "primary_cluster_nodes_details.0.region"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "primary_cluster_nodes_details.0.workspace_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "primary_cluster_nodes_details.0.agent_status"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "primary_cluster_nodes_details.0.pha_level"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "secondary_cluster_nodes_details.#"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "secondary_cluster_nodes_details.0.vm_name"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "secondary_cluster_nodes_details.0.vm_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "secondary_cluster_nodes_details.0.ip_address"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "secondary_cluster_nodes_details.0.cores"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "secondary_cluster_nodes_details.0.vm_status"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "secondary_cluster_nodes_details.0.memory"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "secondary_cluster_nodes_details.0.region"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "secondary_cluster_nodes_details.0.workspace_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "secondary_cluster_nodes_details.0.agent_status"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "secondary_cluster_nodes_details.0.pha_level"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "primary_workspace_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_deployment.pha_deployment_instance", "secondary_workspace_crn"),
				),
			},
		},
	})
}

func testAccCheckIBMPhaDeploymentDataSourceConfigBasic(phaDeploymentResponsePhaInstanceID string, phaDeploymentResponsePrimaryWorkspace string) string {
	return fmt.Sprintf(`
		resource "ibm_pha_deployment" "pha_deployment_instance" {
			pha_instance_id = "%s"
			primary_workspace = "%s"
		}

		data "ibm_pha_deployment" "pha_deployment_instance" {
			pha_instance_id = ibm_pha_deployment.pha_deployment_instance.pha_instance_id
			if_none_match = ibm_pha_deployment.pha_deployment_instance.if_none_match
		}
	`, phaDeploymentResponsePhaInstanceID, phaDeploymentResponsePrimaryWorkspace)
}

func testAccCheckIBMPhaDeploymentDataSourceConfig(phaDeploymentResponsePhaInstanceID string, phaDeploymentResponseAcceptLanguage string, phaDeploymentResponseIfNoneMatch string, phaDeploymentResponsePrimaryWorkspace string, phaDeploymentResponseSecondaryWorkspace string) string {
	return fmt.Sprintf(`
		resource "ibm_pha_deployment" "pha_deployment_instance" {
			pha_instance_id = "%s"
			accept_language = "%s"
			if_none_match = "%s"
			primary_workspace = "%s"
			secondary_workspace = "%s"
		}

		data "ibm_pha_deployment" "pha_deployment_instance" {
			pha_instance_id = ibm_pha_deployment.pha_deployment_instance.pha_instance_id
			if_none_match = ibm_pha_deployment.pha_deployment_instance.if_none_match
		}
	`, phaDeploymentResponsePhaInstanceID, phaDeploymentResponseAcceptLanguage, phaDeploymentResponseIfNoneMatch, phaDeploymentResponsePrimaryWorkspace, phaDeploymentResponseSecondaryWorkspace)
}

func TestDataSourceIBMPhaDeploymentClusterNodeInfoToMap(t *testing.T) {
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

	result, err := powerhaautomationservice.DataSourceIBMPhaDeploymentClusterNodeInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
