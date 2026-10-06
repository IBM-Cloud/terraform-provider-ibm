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

func TestAccIBMPhaClusterNodesBasic(t *testing.T) {
	var conf powerhaautomationservicev1.ClusterNodeResponse
	phaInstanceID := fmt.Sprintf("tf_pha_instance_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPhaClusterNodesDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPhaClusterNodesConfigBasic(phaInstanceID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPhaClusterNodesExists("ibm_pha_cluster_nodes.pha_cluster_nodes_instance", conf),
					resource.TestCheckResourceAttr("ibm_pha_cluster_nodes.pha_cluster_nodes_instance", "pha_instance_id", phaInstanceID),
				),
			},
		},
	})
}

func TestAccIBMPhaClusterNodesAllArgs(t *testing.T) {
	var conf powerhaautomationservicev1.ClusterNodeResponse
	phaInstanceID := fmt.Sprintf("tf_pha_instance_id_%d", acctest.RandIntRange(10, 100))
	acceptLanguage := fmt.Sprintf("tf_accept_language_%d", acctest.RandIntRange(10, 100))
	ifNoneMatch := fmt.Sprintf("tf_if_none_match_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPhaClusterNodesDestroy,
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
			resource.TestStep{
				ResourceName:      "ibm_pha_cluster_nodes.pha_cluster_nodes_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMPhaClusterNodesConfigBasic(phaInstanceID string) string {
	return fmt.Sprintf(`
		resource "ibm_pha_cluster_nodes" "pha_cluster_nodes_instance" {
			pha_instance_id = "%s"
		}
	`, phaInstanceID)
}

func testAccCheckIBMPhaClusterNodesConfig(phaInstanceID string, acceptLanguage string, ifNoneMatch string) string {
	return fmt.Sprintf(`

		resource "ibm_pha_cluster_nodes" "pha_cluster_nodes_instance" {
			pha_instance_id = "%s"
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

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getClusterNodeOptions.SetPhaInstanceID(parts[0])
		getClusterNodeOptions.SetPhaInstanceID(parts[1])

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

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getClusterNodeOptions.SetPhaInstanceID(parts[0])
		getClusterNodeOptions.SetPhaInstanceID(parts[1])

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

	result, err := powerhaautomationservice.ResourceIBMPhaClusterNodesNodeDetailToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
