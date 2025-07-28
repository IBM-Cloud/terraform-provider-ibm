// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.97.2-fc613b62-20241203-155509
 */

package vmware_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vmware"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vmware-go-sdk/vmwarev1"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmVmaasTransitGatewayConnectionDataSourceBasic(t *testing.T) {
	transitGatewayVdcID := acc.Vmaas_vdc_id
	transitGatewayEdgeID := acc.Vmaas_edge_id
	transitGatewayID := acc.Vmaas_transit_gateway_id

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckVMwareService(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmVmaasTransitGatewayConnectionDataSourceConfigBasic(transitGatewayVdcID, transitGatewayEdgeID, transitGatewayID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "vmaas_transit_gateway_connection_id"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "connections.#"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "region"),
				),
			},
		},
	})
}

func TestAccIbmVmaasTransitGatewayConnectionDataSourceAllArgs(t *testing.T) {
	transitGatewayVdcID := acc.Vmaas_vdc_id
	transitGatewayEdgeID := acc.Vmaas_edge_id
	transitGatewayAcceptLanguage := "en-us"
	transitGatewayRegion := "jp-tok"
	transitGatewayID := acc.Vmaas_transit_gateway_id

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckVMwareService(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmVmaasTransitGatewayConnectionDataSourceConfig(transitGatewayVdcID, transitGatewayEdgeID, transitGatewayAcceptLanguage, transitGatewayRegion, transitGatewayID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "vmaas_transit_gateway_connection_id"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "connections.#"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "connections.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "connections.0.transit_gateway_connection_name"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "connections.0.status"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "connections.0.local_gateway_ip"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "connections.0.remote_gateway_ip"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "connections.0.local_tunnel_ip"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "connections.0.remote_tunnel_ip"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "connections.0.local_bgp_asn"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "connections.0.remote_bgp_asn"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "connections.0.network_account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "connections.0.network_type"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "connections.0.base_network_type"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "connections.0.zone"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "region"),
				),
			},
		},
	})
}

func testAccCheckIbmVmaasTransitGatewayConnectionDataSourceConfigBasic(transitGatewayVdcID string, transitGatewayEdgeID string, transitGatewayID string) string {
	return fmt.Sprintf(`
		resource "ibm_vmaas_transit_gateway_connection" "vmaas_transit_gateway_connection_instance" {
			vdc_id = "%s"
			edge_id = "%s"
			vmaas_transit_gateway_connection_id ="%s"
		}

		data "ibm_vmaas_transit_gateway_connection" "vmaas_transit_gateway_connection_instance" {
			vmaas_transit_gateway_connection_id = ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance.vmaas_transit_gateway_connection_id
		}
	`, transitGatewayVdcID, transitGatewayEdgeID, transitGatewayID)
}

func testAccCheckIbmVmaasTransitGatewayConnectionDataSourceConfig(transitGatewayVdcID string, transitGatewayEdgeID string, transitGatewayAcceptLanguage string, transitGatewayRegion string, transitGatewayID string) string {
	return fmt.Sprintf(`
		resource "ibm_vmaas_transit_gateway_connection" "vmaas_transit_gateway_connection_instance" {
			vdc_id = "%s"
			edge_id = "%s"
			accept_language = "%s"
			region = "%s"
			vmaas_transit_gateway_connection_id ="%s"
		}

		data "ibm_vmaas_transit_gateway_connection" "vmaas_transit_gateway_connection_instance" {
			vmaas_transit_gateway_connection_id = ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance.vmaas_transit_gateway_connection_id
		}
	`, transitGatewayVdcID, transitGatewayEdgeID, transitGatewayAcceptLanguage, transitGatewayRegion, transitGatewayID)
}

func TestDataSourceIbmVmaasTransitGatewayConnectionTransitGatewayConnectionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["transit_gateway_connection_name"] = "testString"
		model["status"] = "pending"
		model["local_gateway_ip"] = "testString"
		model["remote_gateway_ip"] = "testString"
		model["local_tunnel_ip"] = "testString"
		model["remote_tunnel_ip"] = "testString"
		model["local_bgp_asn"] = int(1)
		model["remote_bgp_asn"] = int(1)
		model["network_account_id"] = "testString"
		model["network_type"] = "testString"
		model["base_network_type"] = "testString"
		model["zone"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(vmwarev1.TransitGatewayConnection)
	model.Name = core.StringPtr("testString")
	model.TransitGatewayConnectionName = core.StringPtr("testString")
	model.Status = core.StringPtr("pending")
	model.LocalGatewayIp = core.StringPtr("testString")
	model.RemoteGatewayIp = core.StringPtr("testString")
	model.LocalTunnelIp = core.StringPtr("testString")
	model.RemoteTunnelIp = core.StringPtr("testString")
	model.LocalBgpAsn = core.Int64Ptr(int64(1))
	model.RemoteBgpAsn = core.Int64Ptr(int64(1))
	model.NetworkAccountID = core.StringPtr("testString")
	model.NetworkType = core.StringPtr("testString")
	model.BaseNetworkType = core.StringPtr("testString")
	model.Zone = core.StringPtr("testString")

	result, err := vmware.DataSourceIbmVmaasTransitGatewayConnectionTransitGatewayConnectionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
