// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vmware_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vmware-go-sdk/vmwarev1"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/service/vmware"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmVmaasTransitGatewayConnectionBasic(t *testing.T) {
	var conf vmwarev1.TransitGateway
	vdcID := acc.Vmaas_vdc_id
	edgeID := acc.Vmaas_edge_id
	id := acc.Vmaas_transit_gateway_id

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckVMwareTGWService(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmVmaasTransitGatewayConnectionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmVmaasTransitGatewayConnectionConfigBasic(vdcID, edgeID, id),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmVmaasTransitGatewayConnectionExists("ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", conf),
					resource.TestCheckResourceAttr("ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "vdc_id", vdcID),
					resource.TestCheckResourceAttr("ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "edge_id", edgeID),
					resource.TestCheckResourceAttr("ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "vmaas_transit_gateway_connection_id", id),
				),
			},
		},
	})
}

func TestAccIbmVmaasTransitGatewayConnectionAllArgs(t *testing.T) {
	var conf vmwarev1.TransitGateway
	vdcID := acc.Vmaas_vdc_id
	edgeID := acc.Vmaas_edge_id
	id := acc.Vmaas_transit_gateway_id
	region := "jp-tok"
	regionUpdate := region

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckVMwareTGWService(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmVmaasTransitGatewayConnectionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmVmaasTransitGatewayConnectionConfig(vdcID, edgeID, region, id),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmVmaasTransitGatewayConnectionExists("ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", conf),
					resource.TestCheckResourceAttr("ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "vdc_id", vdcID),
					resource.TestCheckResourceAttr("ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "edge_id", edgeID),
					resource.TestCheckResourceAttr("ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "region", region),
					resource.TestCheckResourceAttr("ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "vmaas_transit_gateway_connection_id", id),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmVmaasTransitGatewayConnectionConfig(vdcID, edgeID, regionUpdate, id),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "vdc_id", vdcID),
					resource.TestCheckResourceAttr("ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "edge_id", edgeID),
					resource.TestCheckResourceAttr("ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "region", regionUpdate),
					resource.TestCheckResourceAttr("ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance", "vmaas_transit_gateway_connection_id", id),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_vmaas_transit_gateway_connection.vmaas_transit_gateway_connection_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmVmaasTransitGatewayConnectionConfigBasic(vdcID string, edgeID string, id string) string {
	return fmt.Sprintf(`
		resource "ibm_vmaas_transit_gateway_connection" "vmaas_transit_gateway_connection_instance" {
			vdc_id = "%s"
			edge_id = "%s"
			vmaas_transit_gateway_connection_id ="%s"
		}
	`, vdcID, edgeID, id)
}

func testAccCheckIbmVmaasTransitGatewayConnectionConfig(vdcID string, edgeID string, region string, id string) string {
	return fmt.Sprintf(`
		resource "ibm_vmaas_transit_gateway_connection" "vmaas_transit_gateway_connection_instance" {
			vdc_id = "%s"
			edge_id = "%s"
			region = "%s"
			vmaas_transit_gateway_connection_id ="%s"
		}
	`, vdcID, edgeID, region, id)
}

func testAccCheckIbmVmaasTransitGatewayConnectionExists(n string, obj vmwarev1.TransitGateway) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vmwareClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VmwareV1()
		if err != nil {
			return err
		}

		getVdcOptions := &vmwarev1.GetVdcOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		vdcID := parts[0]
		edgeID := parts[1]
		tgwID := parts[2]

		getVdcOptions.SetID(vdcID)

		vdc, _, err := vmwareClient.GetVdc(getVdcOptions)
		if err != nil {
			return fmt.Errorf("error fetching VDC: %s", err)
		}

		var foundEdge *vmwarev1.Edge
		for _, edge := range vdc.Edges {
			if edge.ID != nil && *edge.ID == edgeID {
				foundEdge = &edge
				break
			}
		}
		if foundEdge == nil {
			return fmt.Errorf("edge %q not found in VDC %q", edgeID, vdcID)
		}

		for _, tg := range foundEdge.TransitGateways {
			if tg.ID != nil && *tg.ID == tgwID {
				obj = tg
				return nil
			}
		}
		return fmt.Errorf("transit gateway %q not found in edge %q", tgwID, edgeID)

	}
}

func testAccCheckIbmVmaasTransitGatewayConnectionDestroy(s *terraform.State) error {
	vmwareClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VmwareV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_vmaas_transit_gateway_connection" {
			continue
		}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil || len(parts) < 3 {
			return fmt.Errorf("unexpected ID format (%q): %v", rs.Primary.ID, err)
		}
		vdcID := parts[0]
		edgeID := parts[1]
		tgwID := parts[2]

		getVdcOptions := &vmwarev1.GetVdcOptions{}
		getVdcOptions.SetID(vdcID)

		time.Sleep(120 * time.Second)

		vdc, response, err := vmwareClient.GetVdc(getVdcOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				continue
			}
			return fmt.Errorf("error retrieving VDC: %v", err)
		}

		var foundEdge *vmwarev1.Edge
		for _, edge := range vdc.Edges {
			if edge.ID != nil && *edge.ID == edgeID {
				foundEdge = &edge
				break
			}
		}
		if foundEdge == nil {
			continue
		}

		for _, tg := range foundEdge.TransitGateways {
			if tg.ID != nil && *tg.ID == tgwID {
				return fmt.Errorf("transit gateway %q still exists in edge %q", tgwID, edgeID)
			}
		}
	}
	return nil
}

func TestResourceIbmVmaasTransitGatewayConnectionTransitGatewayConnectionToMap(t *testing.T) {
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

	result, err := vmware.ResourceIbmVmaasTransitGatewayConnectionTransitGatewayConnectionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
