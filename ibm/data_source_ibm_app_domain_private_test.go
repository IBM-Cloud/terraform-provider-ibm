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

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMAppDomainPrivateDataSource_basic(t *testing.T) {
	name := fmt.Sprintf("terraform%d.com", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMAppDomainPrivateDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.ibm_app_domain_private.testacc_domain", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMAppDomainPrivateDataSourceConfig(name string) string {
	return fmt.Sprintf(`
	
		data "ibm_org" "orgdata" {
			org    = "%s"
		}

		resource "ibm_app_domain_private" "domain" {
			name = "%s"
			org_guid = data.ibm_org.orgdata.id
		}
	
		data "ibm_app_domain_private" "testacc_domain" {
			name = ibm_app_domain_private.domain.name
		}`, cfOrganization, name)

}
