package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMCisDomain_basic(t *testing.T) {
	name := "ibm_cis_domain.test"

	//Fail if cis_crn not set
	if cis_crn == "" {
		panic("IBM_CIS_CRN environment variable not set - required to test CIS")
	}
	if cis_domain == "" {
		panic("IBM_CIS_DOMAIN environment variable not set - required to test CIS")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMCisDomainConfig_basic("test", cis_crn, cis_domain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "domain", cis_domain),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
				),
			},
		},
	})
}

func testAccIBMCisDomainConfig_basic(resourceID string, cis_crn string, zoneName string) string {
	return fmt.Sprintf(`
				resource "ibm_cis_domain" "%[1]s" {
					cis_id = "%[2]s"
                    domain = "%[3]s"
				}`, resourceID, cis_crn, zoneName)
}
