// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vmware_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/vmware-go-sdk/vmwarev1"
)

func TestAccIbmVmaasVdcBasic(t *testing.T) {
	var conf vmwarev1.VDC

	ds_id := acc.Vmaas_Directorsite_id
	ds_pvdc_id := acc.Vmaas_Directorsite_pvdc_id
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckVMwareService(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmVmaasVdcDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmVmaasVdcConfigBasic(ds_id, ds_pvdc_id, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmVmaasVdcExists("ibm_vmaas_vdc.vmaas_vdc_instance", conf),
					resource.TestCheckResourceAttr("ibm_vmaas_vdc.vmaas_vdc_instance", "name", name),
				),
			},
		},
	})
}

func TestAccIbmVmaasVdcAllArgs(t *testing.T) {
	var conf vmwarev1.VDC

	ds_id := acc.Vmaas_Directorsite_id
	ds_pvdc_id := acc.Vmaas_Directorsite_pvdc_id

	cpu := fmt.Sprintf("%d", acctest.RandIntRange(0, 2000))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	ram := fmt.Sprintf("%d", acctest.RandIntRange(0, 40960))
	fastProvisioningEnabled := "false"
	rhelByol := "false"
	windowsByol := "false"

	cpuUpdate := fmt.Sprintf("%d", acctest.RandIntRange(0, 2000))
	nameUpdate := name
	ramUpdate := fmt.Sprintf("%d", acctest.RandIntRange(0, 40960))
	fastProvisioningEnabledUpdate := fastProvisioningEnabled
	rhelByolUpdate := rhelByol
	windowsByolUpdate := windowsByol

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckVMwareService(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmVmaasVdcDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmVmaasVdcConfig(cpu, ds_id, ds_pvdc_id, name, ram, fastProvisioningEnabled, rhelByol, windowsByol),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmVmaasVdcExists("ibm_vmaas_vdc.vmaas_vdc_instance", conf),
					resource.TestCheckResourceAttr("ibm_vmaas_vdc.vmaas_vdc_instance", "cpu", cpu),
					resource.TestCheckResourceAttr("ibm_vmaas_vdc.vmaas_vdc_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_vmaas_vdc.vmaas_vdc_instance", "ram", ram),
					resource.TestCheckResourceAttr("ibm_vmaas_vdc.vmaas_vdc_instance", "fast_provisioning_enabled", fastProvisioningEnabled),
					resource.TestCheckResourceAttr("ibm_vmaas_vdc.vmaas_vdc_instance", "rhel_byol", rhelByol),
					resource.TestCheckResourceAttr("ibm_vmaas_vdc.vmaas_vdc_instance", "windows_byol", windowsByol),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmVmaasVdcConfig(cpuUpdate, ds_id, ds_pvdc_id, nameUpdate, ramUpdate, fastProvisioningEnabledUpdate, rhelByolUpdate, windowsByolUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_vmaas_vdc.vmaas_vdc_instance", "cpu", cpuUpdate),
					resource.TestCheckResourceAttr("ibm_vmaas_vdc.vmaas_vdc_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_vmaas_vdc.vmaas_vdc_instance", "ram", ramUpdate),
					resource.TestCheckResourceAttr("ibm_vmaas_vdc.vmaas_vdc_instance", "fast_provisioning_enabled", fastProvisioningEnabledUpdate),
					resource.TestCheckResourceAttr("ibm_vmaas_vdc.vmaas_vdc_instance", "rhel_byol", rhelByolUpdate),
					resource.TestCheckResourceAttr("ibm_vmaas_vdc.vmaas_vdc_instance", "windows_byol", windowsByolUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_vmaas_vdc.vmaas_vdc_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmVmaasVdcConfigBasic(ds_id string, ds_pvdc_id string, name string) string {
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
	`, ds_id, ds_pvdc_id, name)
}

func testAccCheckIbmVmaasVdcConfig(cpu string, ds_id string, ds_pvdc_id string, name string, ram string, fastProvisioningEnabled string, rhelByol string, windowsByol string) string {
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
	`, cpu, ds_id, ds_pvdc_id, name, ram, fastProvisioningEnabled, rhelByol, windowsByol)
}

func testAccCheckIbmVmaasVdcExists(n string, obj vmwarev1.VDC) resource.TestCheckFunc {

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

		getVdcOptions.SetID(rs.Primary.ID)

		vDC, _, err := vmwareClient.GetVdc(getVdcOptions)
		if err != nil {
			return err
		}

		obj = *vDC
		return nil
	}
}

func testAccCheckIbmVmaasVdcDestroy(s *terraform.State) error {
	vmwareClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VmwareV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_vmaas_vdc" {
			continue
		}

		getVdcOptions := &vmwarev1.GetVdcOptions{}

		getVdcOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := vmwareClient.GetVdc(getVdcOptions)

		if err == nil {
			return fmt.Errorf("vmaas_vdc still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for vmaas_vdc (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
