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

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testZoneConfig("test", cis_crn, "example.org"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "domain", "example.org"),
					resource.TestCheckResourceAttr(name, "name_servers.#", "2"),
				),
			},
		},
	})
}

func testZoneConfig(resourceID string, cis_crn string, zoneName string) string {
	return fmt.Sprintf(`
				resource "ibm_cis_domain" "%s" {
					cis_id = "%s"
                    domain = "%s"
				}`, resourceID, cis_crn, zoneName)
}
