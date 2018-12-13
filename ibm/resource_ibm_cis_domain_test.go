package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMCisDomain_basic(t *testing.T) {
	name := "ibm_cis_domain.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		// No requirement for CheckDestory of this resource as by reaching this test it must have already been deleted
		// correctly during the resource destroy phase of test. The destroy of resource_ibm_cis used in testAccCheckCisPoolConfigBasic
		// will fail if this resource is not correctly deleted.
		Steps: []resource.TestStep{
			{
				Config: testAccIBMCisDomainConfig_basic("test", cis_domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "domain", cis_domain),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
				),
			},
		},
	})
}

func testAccIBMCisDomainConfig_basic(resourceID string, cis_domain string) string {
	return testAccCheckIBMCISInstance_basic("test") + fmt.Sprintf(`
				resource "ibm_cis_domain" "%[1]s" {
					cis_id = "${ibm_cis.instance.id}"
                    domain = "%[2]s"
				}`, resourceID, cis_domain)
}
