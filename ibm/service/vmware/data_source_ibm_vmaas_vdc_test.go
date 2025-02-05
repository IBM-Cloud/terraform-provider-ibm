// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.98.0-8be2046a-20241205-162752
 */

package vmware_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vmware"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vmware-go-sdk/vmwarev1"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmVmaasVdcDataSourceBasic(t *testing.T) {
	vDCName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	ds_id := acc.Vmaas_Directorsite_id
	ds_pvdc_id := acc.Vmaas_Directorsite_pvdc_id

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckVMwareService(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmVmaasVdcDataSourceConfigBasic(ds_id, ds_pvdc_id, vDCName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "vmaas_vdc_id"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "director_site.#"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "edges.#"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "status_reasons.#"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "ordered_at"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "org_href"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "org_name"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "fast_provisioning_enabled"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "rhel_byol"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "windows_byol"),
				),
			},
		},
	})
}

func TestAccIbmVmaasVdcDataSourceAllArgs(t *testing.T) {
	ds_id := acc.Vmaas_Directorsite_id
	ds_pvdc_id := acc.Vmaas_Directorsite_pvdc_id

	vDCCpu := fmt.Sprintf("%d", acctest.RandIntRange(0, 2000))
	vDCName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	vDCRam := fmt.Sprintf("%d", acctest.RandIntRange(0, 40960))
	vDCFastProvisioningEnabled := "false"
	vDCRhelByol := "false"
	vDCWindowsByol := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckVMwareService(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmVmaasVdcDataSourceConfig(vDCCpu, ds_id, ds_pvdc_id, vDCName, vDCRam, vDCFastProvisioningEnabled, vDCRhelByol, vDCWindowsByol),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "vmaas_vdc_id"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "href"),
					// resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "provisioned_at"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "cpu"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "crn"),
					// resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "deleted_at"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "director_site.#"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "edges.#"),
					/*
						resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "edges.0.id"),
						resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "edges.0.private_only"),
						resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "edges.0.size"),
						resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "edges.0.status"),
						resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "edges.0.type"),
						resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "edges.0.version"),
					*/
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "status_reasons.#"),
					/*
						resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "status_reasons.0.code"),
						resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "status_reasons.0.message"),
						resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "status_reasons.0.more_info"),
					*/
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "ordered_at"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "org_href"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "org_name"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "ram"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "fast_provisioning_enabled"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "rhel_byol"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "windows_byol"),
				),
			},
		},
	})
}

func testAccCheckIbmVmaasVdcDataSourceConfigBasic(ds_id string, ds_pvdc_id string, vDCName string) string {
	return fmt.Sprintf(`
		resource "ibm_vmaas_vdc" "vmaas_vdc_instance" {
			director_site {
				id = "%s"
				pvdc {
					id = "%s"
					provider_type {
						name = "on_demand"
					}
				}
			}
			name = "%s"
		}

		data "ibm_vmaas_vdc" "vmaas_vdc_instance" {
			vmaas_vdc_id = ibm_vmaas_vdc.vmaas_vdc_instance.id
		}
	`, ds_id, ds_pvdc_id, vDCName)
}

func testAccCheckIbmVmaasVdcDataSourceConfig(vDCCpu string, ds_id string, ds_pvdc_id string, vDCName string, vDCRam string, vDCFastProvisioningEnabled string, vDCRhelByol string, vDCWindowsByol string) string {
	return fmt.Sprintf(`
		resource "ibm_vmaas_vdc" "vmaas_vdc_instance" {
			cpu = %s
			director_site {
				id = "%s"
				pvdc {
					id = "%s"
					provider_type {
						name = "on_demand"
					}
				}
			}
			name = "%s"
			ram = %s
			fast_provisioning_enabled = %s
			rhel_byol = %s
			windows_byol = %s
		}

		data "ibm_vmaas_vdc" "vmaas_vdc_instance" {
			vmaas_vdc_id = ibm_vmaas_vdc.vmaas_vdc_instance.id
		}
	`, vDCCpu, ds_id, ds_pvdc_id, vDCName, vDCRam, vDCFastProvisioningEnabled, vDCRhelByol, vDCWindowsByol)
}

func TestDataSourceIbmVmaasVdcVDCDirectorSiteToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		vdcProviderTypeModel := make(map[string]interface{})
		vdcProviderTypeModel["name"] = "paygo"

		directorSitePvdcModel := make(map[string]interface{})
		directorSitePvdcModel["id"] = "pvdc_id"
		directorSitePvdcModel["provider_type"] = []map[string]interface{}{vdcProviderTypeModel}

		model := make(map[string]interface{})
		model["id"] = "testString"
		model["pvdc"] = []map[string]interface{}{directorSitePvdcModel}
		model["url"] = "testString"

		assert.Equal(t, result, model)
	}

	vdcProviderTypeModel := new(vmwarev1.VDCProviderType)
	vdcProviderTypeModel.Name = core.StringPtr("paygo")

	directorSitePvdcModel := new(vmwarev1.DirectorSitePVDC)
	directorSitePvdcModel.ID = core.StringPtr("pvdc_id")
	directorSitePvdcModel.ProviderType = vdcProviderTypeModel

	model := new(vmwarev1.VDCDirectorSite)
	model.ID = core.StringPtr("testString")
	model.Pvdc = directorSitePvdcModel
	model.URL = core.StringPtr("testString")

	result, err := vmware.DataSourceIbmVmaasVdcVDCDirectorSiteToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmVmaasVdcDirectorSitePVDCToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		vdcProviderTypeModel := make(map[string]interface{})
		vdcProviderTypeModel["name"] = "paygo"

		model := make(map[string]interface{})
		model["id"] = "pvdc_id"
		model["provider_type"] = []map[string]interface{}{vdcProviderTypeModel}

		assert.Equal(t, result, model)
	}

	vdcProviderTypeModel := new(vmwarev1.VDCProviderType)
	vdcProviderTypeModel.Name = core.StringPtr("paygo")

	model := new(vmwarev1.DirectorSitePVDC)
	model.ID = core.StringPtr("pvdc_id")
	model.ProviderType = vdcProviderTypeModel

	result, err := vmware.DataSourceIbmVmaasVdcDirectorSitePVDCToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmVmaasVdcVDCProviderTypeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "paygo"

		assert.Equal(t, result, model)
	}

	model := new(vmwarev1.VDCProviderType)
	model.Name = core.StringPtr("paygo")

	result, err := vmware.DataSourceIbmVmaasVdcVDCProviderTypeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmVmaasVdcEdgeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		transitGatewayConnectionModel := make(map[string]interface{})
		transitGatewayConnectionModel["name"] = "testString"
		transitGatewayConnectionModel["transit_gateway_connection_name"] = "testString"
		transitGatewayConnectionModel["status"] = "pending"
		transitGatewayConnectionModel["local_gateway_ip"] = "testString"
		transitGatewayConnectionModel["remote_gateway_ip"] = "testString"
		transitGatewayConnectionModel["local_tunnel_ip"] = "testString"
		transitGatewayConnectionModel["remote_tunnel_ip"] = "testString"
		transitGatewayConnectionModel["local_bgp_asn"] = int(1)
		transitGatewayConnectionModel["remote_bgp_asn"] = int(1)
		transitGatewayConnectionModel["network_account_id"] = "testString"
		transitGatewayConnectionModel["network_type"] = "testString"
		transitGatewayConnectionModel["base_network_type"] = "testString"
		transitGatewayConnectionModel["zone"] = "testString"

		transitGatewayModel := make(map[string]interface{})
		transitGatewayModel["id"] = "testString"
		transitGatewayModel["connections"] = []map[string]interface{}{transitGatewayConnectionModel}
		transitGatewayModel["status"] = "pending"
		transitGatewayModel["region"] = "testString"

		model := make(map[string]interface{})
		model["id"] = "testString"
		model["public_ips"] = []string{"testString"}
		model["private_ips"] = []string{"testString"}
		model["private_only"] = true
		model["size"] = "medium"
		model["status"] = "creating"
		model["transit_gateways"] = []map[string]interface{}{transitGatewayModel}
		model["type"] = "performance"
		model["version"] = "testString"

		assert.Equal(t, result, model)
	}

	transitGatewayConnectionModel := new(vmwarev1.TransitGatewayConnection)
	transitGatewayConnectionModel.Name = core.StringPtr("testString")
	transitGatewayConnectionModel.TransitGatewayConnectionName = core.StringPtr("testString")
	transitGatewayConnectionModel.Status = core.StringPtr("pending")
	transitGatewayConnectionModel.LocalGatewayIp = core.StringPtr("testString")
	transitGatewayConnectionModel.RemoteGatewayIp = core.StringPtr("testString")
	transitGatewayConnectionModel.LocalTunnelIp = core.StringPtr("testString")
	transitGatewayConnectionModel.RemoteTunnelIp = core.StringPtr("testString")
	transitGatewayConnectionModel.LocalBgpAsn = core.Int64Ptr(int64(1))
	transitGatewayConnectionModel.RemoteBgpAsn = core.Int64Ptr(int64(1))
	transitGatewayConnectionModel.NetworkAccountID = core.StringPtr("testString")
	transitGatewayConnectionModel.NetworkType = core.StringPtr("testString")
	transitGatewayConnectionModel.BaseNetworkType = core.StringPtr("testString")
	transitGatewayConnectionModel.Zone = core.StringPtr("testString")

	transitGatewayModel := new(vmwarev1.TransitGateway)
	transitGatewayModel.ID = core.StringPtr("testString")
	transitGatewayModel.Connections = []vmwarev1.TransitGatewayConnection{*transitGatewayConnectionModel}
	transitGatewayModel.Status = core.StringPtr("pending")
	transitGatewayModel.Region = core.StringPtr("testString")

	model := new(vmwarev1.Edge)
	model.ID = core.StringPtr("testString")
	model.PublicIps = []string{"testString"}
	model.PrivateIps = []string{"testString"}
	model.PrivateOnly = core.BoolPtr(true)
	model.Size = core.StringPtr("medium")
	model.Status = core.StringPtr("creating")
	model.TransitGateways = []vmwarev1.TransitGateway{*transitGatewayModel}
	model.Type = core.StringPtr("performance")
	model.Version = core.StringPtr("testString")

	result, err := vmware.DataSourceIbmVmaasVdcEdgeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmVmaasVdcTransitGatewayToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		transitGatewayConnectionModel := make(map[string]interface{})
		transitGatewayConnectionModel["name"] = "testString"
		transitGatewayConnectionModel["transit_gateway_connection_name"] = "testString"
		transitGatewayConnectionModel["status"] = "pending"
		transitGatewayConnectionModel["local_gateway_ip"] = "testString"
		transitGatewayConnectionModel["remote_gateway_ip"] = "testString"
		transitGatewayConnectionModel["local_tunnel_ip"] = "testString"
		transitGatewayConnectionModel["remote_tunnel_ip"] = "testString"
		transitGatewayConnectionModel["local_bgp_asn"] = int(1)
		transitGatewayConnectionModel["remote_bgp_asn"] = int(1)
		transitGatewayConnectionModel["network_account_id"] = "testString"
		transitGatewayConnectionModel["network_type"] = "testString"
		transitGatewayConnectionModel["base_network_type"] = "testString"
		transitGatewayConnectionModel["zone"] = "testString"

		model := make(map[string]interface{})
		model["id"] = "testString"
		model["connections"] = []map[string]interface{}{transitGatewayConnectionModel}
		model["status"] = "pending"
		model["region"] = "testString"

		assert.Equal(t, result, model)
	}

	transitGatewayConnectionModel := new(vmwarev1.TransitGatewayConnection)
	transitGatewayConnectionModel.Name = core.StringPtr("testString")
	transitGatewayConnectionModel.TransitGatewayConnectionName = core.StringPtr("testString")
	transitGatewayConnectionModel.Status = core.StringPtr("pending")
	transitGatewayConnectionModel.LocalGatewayIp = core.StringPtr("testString")
	transitGatewayConnectionModel.RemoteGatewayIp = core.StringPtr("testString")
	transitGatewayConnectionModel.LocalTunnelIp = core.StringPtr("testString")
	transitGatewayConnectionModel.RemoteTunnelIp = core.StringPtr("testString")
	transitGatewayConnectionModel.LocalBgpAsn = core.Int64Ptr(int64(1))
	transitGatewayConnectionModel.RemoteBgpAsn = core.Int64Ptr(int64(1))
	transitGatewayConnectionModel.NetworkAccountID = core.StringPtr("testString")
	transitGatewayConnectionModel.NetworkType = core.StringPtr("testString")
	transitGatewayConnectionModel.BaseNetworkType = core.StringPtr("testString")
	transitGatewayConnectionModel.Zone = core.StringPtr("testString")

	model := new(vmwarev1.TransitGateway)
	model.ID = core.StringPtr("testString")
	model.Connections = []vmwarev1.TransitGatewayConnection{*transitGatewayConnectionModel}
	model.Status = core.StringPtr("pending")
	model.Region = core.StringPtr("testString")

	result, err := vmware.DataSourceIbmVmaasVdcTransitGatewayToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmVmaasVdcTransitGatewayConnectionToMap(t *testing.T) {
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

	result, err := vmware.DataSourceIbmVmaasVdcTransitGatewayConnectionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmVmaasVdcStatusReasonToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["code"] = "insufficent_cpu"
		model["message"] = "testString"
		model["more_info"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(vmwarev1.StatusReason)
	model.Code = core.StringPtr("insufficent_cpu")
	model.Message = core.StringPtr("testString")
	model.MoreInfo = core.StringPtr("testString")

	result, err := vmware.DataSourceIbmVmaasVdcStatusReasonToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
