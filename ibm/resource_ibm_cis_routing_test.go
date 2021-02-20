/* IBM Confidential
*  Object Code Only Source Materials
*  5747-SM3
*  (c) Copyright IBM Corp. 2017,2021
*
*  The source code for this program is not published or otherwise divested
*  of its trade secrets, irrespective of what has been deposited with the
*  U.S. Copyright Office. */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMCisRouting_Basic(t *testing.T) {
	name := "ibm_cis_routing." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCis(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisRoutingConfigBasic1("test", cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "smart_routing", "on"),
				),
			},
			{
				Config: testAccCheckCisRoutingConfigBasic2("test", cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "smart_routing", "off"),
				),
			},
		},
	})
}

func TestAccIBMCisRouting_Import(t *testing.T) {
	name := "ibm_cis_routing." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisRoutingConfigBasic2("test", cisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "smart_routing", "off"),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckCisRoutingConfigBasic1(id string, cisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_routing" "%[1]s" {
		cis_id          = data.ibm_cis.cis.id
		domain_id       = data.ibm_cis_domain.cis_domain.domain_id
		smart_routing   = "on"
	  }
`, id)
}
func testAccCheckCisRoutingConfigBasic2(id string, cisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_routing" "%[1]s" {
		cis_id          = data.ibm_cis.cis.id
		domain_id       = data.ibm_cis_domain.cis_domain.domain_id
		smart_routing   = "off"
	  }
`, id)
}
