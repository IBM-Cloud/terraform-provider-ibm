// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.113.1-d76630af-20260320-135953
*/

package powerhaautomationservice_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/powerhaautomationservice"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.com/IBM/dra-go-sdk/powerhaautomationservicev1"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMPhaClusterNodesDataSourceBasic(t *testing.T) {
	clusterNodeResponsePhaInstanceID := fmt.Sprintf("tf_instance_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPhaClusterNodesDataSourceConfigBasic(clusterNodeResponsePhaInstanceID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "primary_node_details.#"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "secondary_node_details.#"),
				),
			},
		},
	})
}

func TestAccIBMPhaClusterNodesDataSourceAllArgs(t *testing.T) {
	clusterNodeResponsePhaInstanceID := "748943bf-7965-492a-9ce0-88309cf18451"
	clusterNodeResponseAcceptLanguage := "en"
	clusterNodeResponseIfNoneMatch := ""

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPhaClusterNodesDataSourceConfig(clusterNodeResponsePhaInstanceID, clusterNodeResponseAcceptLanguage, clusterNodeResponseIfNoneMatch),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "if_none_match"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "primary_node_details.#"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "primary_node_details.0.agent_status"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "primary_node_details.0.cores"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "primary_node_details.0.memory"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "primary_node_details.0.pha_level"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "primary_node_details.0.region"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "primary_node_details.0.vm_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "primary_node_details.0.vm_name"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "primary_node_details.0.vm_status"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "primary_node_details.0.workspace_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "secondary_node_details.#"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "secondary_node_details.0.agent_status"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "secondary_node_details.0.cores"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "secondary_node_details.0.memory"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "secondary_node_details.0.pha_level"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "secondary_node_details.0.region"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "secondary_node_details.0.vm_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "secondary_node_details.0.vm_name"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "secondary_node_details.0.vm_status"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "secondary_node_details.0.workspace_id"),
				),
			},
		},
	})
}

func testAccCheckIBMPhaClusterNodesDataSourceConfigBasic(clusterNodeResponsePhaInstanceID string) string {
	return fmt.Sprintf(`
		resource "ibm_pha_cluster_nodes" "pha_cluster_nodes_instance" {
			instance_id = "%s"
		}

		data "ibm_pha_cluster_nodes" "pha_cluster_nodes_instance" {
			instance_id = "748943bf-7965-492a-9ce0-88309cf18451"
		}
	`, clusterNodeResponsePhaInstanceID)
}

func testAccCheckIBMPhaClusterNodesDataSourceConfig(clusterNodeResponsePhaInstanceID string, clusterNodeResponseAcceptLanguage string, clusterNodeResponseIfNoneMatch string) string {
	return fmt.Sprintf(`
		resource "ibm_pha_cluster_nodes" "pha_cluster_nodes_instance" {
			instance_id = "%s"
			accept_language = "%s"
			if_none_match = "%s"
		}

		data "ibm_pha_cluster_nodes" "pha_cluster_nodes_instance" {
			instance_id = "748943bf-7965-492a-9ce0-88309cf18451"
		}
	`, clusterNodeResponsePhaInstanceID, clusterNodeResponseAcceptLanguage, clusterNodeResponseIfNoneMatch)
}

func TestDataSourceIBMPhaClusterNodesNodeDetailToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["agent_status"] = "running"
		model["cores"] = float64(8.0)
		model["ip_addresses"] = []string{"10.0.0.21", "10.0.0.22"}
		model["memory"] = float64(64.0)
		model["pha_level"] = "7.2.1"
		model["region"] = "us-south"
		model["vm_id"] = "vm-9b7c2d11"
		model["vm_name"] = "pha-node-primary-1"
		model["vm_status"] = "ACTIVE"
		model["workspace_id"] = "workspace-primary-001"

		assert.Equal(t, result, model)
	}

	model := new(powerhaautomationservicev1.NodeDetail)
	model.AgentStatus = core.StringPtr("running")
	model.Cores = core.Float32Ptr(float32(8.0))
	model.IPAddresses = []string{"10.0.0.21", "10.0.0.22"}
	model.Memory = core.Float32Ptr(float32(64.0))
	model.PhaLevel = core.StringPtr("7.2.1")
	model.Region = core.StringPtr("us-south")
	model.VMID = core.StringPtr("vm-9b7c2d11")
	model.VMName = core.StringPtr("pha-node-primary-1")
	model.VMStatus = core.StringPtr("ACTIVE")
	model.WorkspaceID = core.StringPtr("workspace-primary-001")

	result, err := powerhaautomationservice.DataSourceIBMPhaClusterNodesNodeDetailToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
