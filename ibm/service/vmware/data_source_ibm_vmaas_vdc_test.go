// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vmware_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
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
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "provisioned_at"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "cpu"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "crn"),
					// resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "deleted_at"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "director_site.#"),
					resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "edges.#"),
					/*
						resource.TestCheckResourceAttrSet("data.ibm_vmaas_vdc.vmaas_vdc_instance", "edges.0.id"),
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
