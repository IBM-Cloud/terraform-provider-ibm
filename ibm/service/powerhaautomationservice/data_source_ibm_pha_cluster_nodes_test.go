// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.116.0-df613dbc-20260803-154903
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
	"github.ibm.com/DRAutomation/dra-go-sdk/powerhaautomationservicev1"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMPhaClusterNodesDataSourceBasic(t *testing.T) {
	clusterNodeResponsePhaInstanceID := fmt.Sprintf("tf_pha_instance_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPhaClusterNodesDataSourceConfigBasic(clusterNodeResponsePhaInstanceID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "pha_instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "primary_node_details.#"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "secondary_node_details.#"),
				),
			},
		},
	})
}

func TestAccIBMPhaClusterNodesDataSourceAllArgs(t *testing.T) {
	clusterNodeResponsePhaInstanceID := fmt.Sprintf("tf_pha_instance_id_%d", acctest.RandIntRange(10, 100))
	clusterNodeResponseAcceptLanguage := fmt.Sprintf("tf_accept_language_%d", acctest.RandIntRange(10, 100))
	clusterNodeResponseIfNoneMatch := fmt.Sprintf("tf_if_none_match_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPhaClusterNodesDataSourceConfig(clusterNodeResponsePhaInstanceID, clusterNodeResponseAcceptLanguage, clusterNodeResponseIfNoneMatch),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "pha_instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "if_none_match"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "primary_node_details.#"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "primary_node_details.0.vm_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "primary_node_details.0.vm_name"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "primary_node_details.0.region"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "primary_node_details.0.workspace_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "primary_node_details.0.cores"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "primary_node_details.0.memory"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "primary_node_details.0.vm_status"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "primary_node_details.0.agent_status"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "primary_node_details.0.pha_level"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "primary_node_details.0.powerha_version_supported"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "secondary_node_details.#"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "secondary_node_details.0.vm_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "secondary_node_details.0.vm_name"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "secondary_node_details.0.region"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "secondary_node_details.0.workspace_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "secondary_node_details.0.cores"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "secondary_node_details.0.memory"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "secondary_node_details.0.vm_status"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "secondary_node_details.0.agent_status"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "secondary_node_details.0.pha_level"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "secondary_node_details.0.powerha_version_supported"),
				),
			},
		},
	})
}

func testAccCheckIBMPhaClusterNodesDataSourceConfigBasic(clusterNodeResponsePhaInstanceID string) string {
	return fmt.Sprintf(`
		resource "ibm_pha_cluster_nodes" "pha_cluster_nodes_instance" {
			pha_instance_id = "%s"
		}

		data "ibm_pha_cluster_nodes" "pha_cluster_nodes_instance" {
			pha_instance_id = ibm_pha_cluster_nodes.pha_cluster_nodes_instance.pha_instance_id
			if_none_match = ibm_pha_cluster_nodes.pha_cluster_nodes_instance.if_none_match
		}
	`, clusterNodeResponsePhaInstanceID)
}

func testAccCheckIBMPhaClusterNodesDataSourceConfig(clusterNodeResponsePhaInstanceID string, clusterNodeResponseAcceptLanguage string, clusterNodeResponseIfNoneMatch string) string {
	return fmt.Sprintf(`
		resource "ibm_pha_cluster_nodes" "pha_cluster_nodes_instance" {
			pha_instance_id = "%s"
			accept_language = "%s"
			if_none_match = "%s"
		}

		data "ibm_pha_cluster_nodes" "pha_cluster_nodes_instance" {
			pha_instance_id = ibm_pha_cluster_nodes.pha_cluster_nodes_instance.pha_instance_id
			if_none_match = ibm_pha_cluster_nodes.pha_cluster_nodes_instance.if_none_match
		}
	`, clusterNodeResponsePhaInstanceID, clusterNodeResponseAcceptLanguage, clusterNodeResponseIfNoneMatch)
}

func TestDataSourceIBMPhaClusterNodesNodeDetailToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["vm_id"] = "vm-1234567890"
		model["vm_name"] = "desire-01"
		model["region"] = "us-south"
		model["workspace_id"] = "2vfdv804c-79a333-5ee-94e7-993ffegve81"
		model["cores"] = float64(5)
		model["memory"] = float64(32)
		model["ip_addresses"] = []string{"10.0.0.25"}
		model["vm_status"] = "RUNNING"
		model["agent_status"] = "ACTIVE"
		model["pha_level"] = "7.2.1"
		model["powerha_version_supported"] = true

		assert.Equal(t, result, model)
	}

	model := new(powerhaautomationservicev1.NodeDetail)
	model.VMID = core.StringPtr("vm-1234567890")
	model.VMName = core.StringPtr("desire-01")
	model.Region = core.StringPtr("us-south")
	model.WorkspaceID = core.StringPtr("2vfdv804c-79a333-5ee-94e7-993ffegve81")
	model.Cores = core.Float32Ptr(float32(5))
	model.Memory = core.Float32Ptr(float32(32))
	model.IPAddresses = []string{"10.0.0.25"}
	model.VMStatus = core.StringPtr("RUNNING")
	model.AgentStatus = core.StringPtr("ACTIVE")
	model.PhaLevel = core.StringPtr("7.2.1")
	model.PowerhaVersionSupported = core.BoolPtr(true)

	result, err := powerhaautomationservice.DataSourceIBMPhaClusterNodesNodeDetailToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
