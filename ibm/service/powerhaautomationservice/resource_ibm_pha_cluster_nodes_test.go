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

	// "github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/powerhaautomationservice"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/DRAutomation/dra-go-sdk/powerhaautomationservicev1"
)

func TestAccIBMPhaClusterNodesBasic(t *testing.T) {
	var conf powerhaautomationservicev1.ClusterNodeResponse
	phaInstanceID := "2cfb7a06-623b-4eb9-a9ac-daa03dc0b5a6"
	// primary_cluster_nodes :="d6feda7b-f679-4869-9e9e-133d9467ba5c"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPhaClusterNodesConfigBasic(phaInstanceID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPhaClusterNodesExists("ibm_pha_cluster_nodes.pha_cluster_nodes_instance", conf),
					resource.TestCheckResourceAttr("ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "pha_instance_id", phaInstanceID),
					// resource.TestCheckResourceAttr("ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "primary_cluster_nodes.0", primary_cluster_nodes),
					resource.TestCheckTypeSetElemAttr(
						"ibm_pha_cluster_nodes.pha_cluster_nodes_instance",
						"primary_cluster_nodes.*",
						"049a8b09-a1ff-4434-acda-92946f3f4ab5",
					),
				),
			},
		},
	})
}

func TestAccIBMPhaClusterNodesAllArgs(t *testing.T) {
	var conf powerhaautomationservicev1.ClusterNodeResponse
	phaInstanceID := "2cfb7a06-623b-4eb9-a9ac-daa03dc0b5a6"
	acceptLanguage := fmt.Sprintf("tf_accept_language_%d", acctest.RandIntRange(10, 100))
	ifNoneMatch := fmt.Sprintf("tf_if_none_match_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPhaClusterNodesConfig(phaInstanceID, acceptLanguage, ifNoneMatch),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPhaClusterNodesExists("ibm_pha_cluster_nodes.pha_cluster_nodes_instance", conf),
					resource.TestCheckResourceAttr("ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "pha_instance_id", phaInstanceID),
					resource.TestCheckResourceAttr("ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "accept_language", acceptLanguage),
					resource.TestCheckResourceAttr("ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "if_none_match", ifNoneMatch),
				),
			},
			// resource.TestStep{
			// 	ResourceName:      "ibm_pha_cluster_nodes.pha_cluster_nodes_instance",
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// },
		},
	})
}

func testAccCheckIBMPhaClusterNodesConfigBasic(phaInstanceID string) string {
	return fmt.Sprintf(`
		resource "ibm_pha_cluster_nodes" "pha_cluster_nodes_instance" {
			pha_instance_id = "%s"
			primary_cluster_nodes = ["049a8b09-a1ff-4434-acda-92946f3f4ab5"]
		}
	`, phaInstanceID)
}

func testAccCheckIBMPhaClusterNodesConfig(phaInstanceID string, acceptLanguage string, ifNoneMatch string) string {
	return fmt.Sprintf(`

		resource "ibm_pha_cluster_nodes" "pha_cluster_nodes_instance" {
			pha_instance_id = "%s"
			primary_cluster_nodes = ["9cdd4756-da0f-449e-a1a6-bb24291bea16"]
			accept_language = "%s"
			if_none_match = "%s"
		}
	`, phaInstanceID, acceptLanguage, ifNoneMatch)
}

func testAccCheckIBMPhaClusterNodesExists(n string, obj powerhaautomationservicev1.ClusterNodeResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		powerhaAutomationServiceClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PowerhaAutomationServiceV1()
		if err != nil {
			return err
		}

		getClusterNodeOptions := &powerhaautomationservicev1.GetClusterNodeOptions{}

		getClusterNodeOptions.SetPhaInstanceID(rs.Primary.ID)

		clusterNodeResponse, _, err := powerhaAutomationServiceClient.GetClusterNode(getClusterNodeOptions)
		if err != nil {
			return err
		}

		obj = *clusterNodeResponse
		return nil
	}
}

func testAccCheckIBMPhaClusterNodesDestroy(s *terraform.State) error {
	powerhaAutomationServiceClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PowerhaAutomationServiceV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pha_cluster_nodes" {
			continue
		}

		getClusterNodeOptions := &powerhaautomationservicev1.GetClusterNodeOptions{}

		getClusterNodeOptions.SetPhaInstanceID(rs.Primary.ID)

		// Try to find the key
		_, response, err := powerhaAutomationServiceClient.GetClusterNode(getClusterNodeOptions)

		if err == nil {
			return fmt.Errorf("pha_cluster_nodes still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for pha_cluster_nodes (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIBMPhaClusterNodesNodeDetailToMap(t *testing.T) {
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

	result, err := powerhaautomationservice.ResourceIBMPhaClusterNodesNodeDetailToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
