package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMComputeBareMetalDataSource_basic(t *testing.T) {
	configName := "data.ibm_compute_bare_metal.tf-bm-ds-acc-test"
	hostname := acctest.RandString(16)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMComputeBareMetalDataSourceConfigBasic(hostname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						configName, "hostname", hostname),
					resource.TestCheckResourceAttr(
						configName, "domain", "terraformuat.ibm.com"),
					resource.TestCheckResourceAttr(
						configName, "os_reference_code", "UBUNTU_16_64"),
					resource.TestCheckResourceAttr(
						configName, "datacenter", "dal01"),
					resource.TestCheckResourceAttr(
						configName, "network_speed", "100"),
					resource.TestCheckResourceAttr(
						configName, "hourly_billing", "true"),
					resource.TestCheckResourceAttr(
						configName, "private_network_only", "false"),
					resource.TestCheckResourceAttr(
						configName, "ipv6_enabled", "true"),
					resource.TestCheckResourceAttr(
						configName, "secondary_ip_count", "4"),
					resource.TestCheckResourceAttrSet(
						configName, "secondary_ip_addresses.0"),
					resource.TestCheckResourceAttrSet(
						configName, "secondary_ip_addresses.1"),
					resource.TestCheckResourceAttrSet(
						configName, "secondary_ip_addresses.2"),
					resource.TestCheckResourceAttrSet(
						configName, "secondary_ip_addresses.3"),
					resource.TestCheckResourceAttr(
						configName, "user_metadata", "{\"value\":\"newvalue\"}"),
					resource.TestCheckResourceAttr(
						configName, "notes", "baremetal notes"),
					CheckStringSet(
						configName,
						"tags", []string{"collectd"},
					),
				),
			},
		},
	})
}

func testAccCheckIBMComputeBareMetalDataSourceConfigBasic(hostname string) string {
	return fmt.Sprintf(`
		resource "ibm_compute_bare_metal" "terraform-acceptance-test-1" {
			hostname               = "%s"
			domain                 = "terraformuat.ibm.com"
			os_reference_code      = "UBUNTU_16_64"
			datacenter             = "dal01"
			ipv6_enabled           = true
			ipv6_static_enabled    = true
			secondary_ip_count     = 4
			network_speed          = 100
			hourly_billing         = true
			private_network_only   = false
			user_metadata          = "{\"value\":\"newvalue\"}"
			fixed_config_preset    = "S1270_32GB_1X1TBSATA_NORAID"
			tags                   = ["collectd"]
			notes                  = "baremetal notes"
			}
			data "ibm_compute_bare_metal" "tf-bm-ds-acc-test" {
				hostname = "${ibm_compute_bare_metal.terraform-acceptance-test-1.hostname}"
				domain = "${ibm_compute_bare_metal.terraform-acceptance-test-1.domain}"
			}`, hostname)

}
