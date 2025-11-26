package scc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSccProviderTypesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckScc(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccProviderTypesDataSourceConfigBasic(acc.SccInstanceID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_types.scc_provider_types_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_types.scc_provider_types_instance", "provider_types.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_types.scc_provider_types_instance", "provider_types.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_types.scc_provider_types_instance", "provider_types.0.id"),
				),
			},
		},
	})
}

func testAccCheckIbmSccProviderTypesDataSourceConfigBasic(instanceID string) string {
	return fmt.Sprintf(`
		data "ibm_scc_provider_types" "scc_provider_types_instance" {
			instance_id = "%s"
		}
	`, instanceID)
}
