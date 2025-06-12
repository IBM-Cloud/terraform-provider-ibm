package scc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSccControlLibrariesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckScc(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccControlLibrariesDataSourceConfigBasic(acc.SccInstanceID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_libraries.scc_control_libraries_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_libraries.scc_control_libraries_instance", "control_libraries.#"),
				),
			},
		},
	})
}

func TestAccIbmSccControlLibrariesDataSourceAllArgs(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckScc(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccControlLibrariesDataSourceConfigAllArgs(acc.SccInstanceID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_libraries.scc_control_libraries_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_control_libraries.scc_control_libraries_instance", "control_libraries.#"),
				),
			},
		},
	})
}

func testAccCheckIbmSccControlLibrariesDataSourceConfigBasic(instanceID string) string {
	return fmt.Sprintf(`
		data "ibm_scc_control_libraries" "scc_control_libraries_instance" {
			instance_id = "%s"
		}
	`, instanceID)
}

func testAccCheckIbmSccControlLibrariesDataSourceConfigAllArgs(instanceID string) string {
	return fmt.Sprintf(`
		data "ibm_scc_control_libraries" "scc_control_libraries_instance" {
      control_library_type = "predefined"
			instance_id = "%s"
		}
	`, instanceID)
}
