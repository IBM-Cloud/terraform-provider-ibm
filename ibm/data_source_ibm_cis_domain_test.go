package ibm

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAccIBMCisDomainDataSource_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCisDomainDataSourceConfig_basic1("test_acc", cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_cis_domain.test_acc", "status", "pending"),
					resource.TestCheckResourceAttr("data.ibm_cis_domain.test_acc", "original_name_servers.#", "2"),
					resource.TestCheckResourceAttr("data.ibm_cis_domain.test_acc", "name_servers.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIBMCisDomainDataSourceConfig_basic1(resourceName string, domain string) string {
	return fmt.Sprintf(`
				data "ibm_cis_domain" "%[1]s" {
					cis_id = "${data.ibm_cis.%[4]s.id}"
                    domain = "%[2]s"
				}
				data "ibm_resource_group" "test_acc" {
				  name = "%[3]s"
				}

				data "ibm_cis" "%[4]s" {
				  resource_group_id = "${data.ibm_resource_group.test_acc.id}"	
				  name = "%[4]s"
				}`, resourceName, domain, cisResourceGroup, cisInstance)
}
