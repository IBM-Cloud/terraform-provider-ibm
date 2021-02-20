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

func TestAccIBMAccountDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMAccountDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_org.testacc_ds_org", "org", cfOrganization),
					resource.TestCheckResourceAttrSet(
						"data.ibm_account.testacc_acc", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMAccountDataSourceConfig() string {
	return fmt.Sprintf(`
	
data "ibm_org" "testacc_ds_org" {
    org = "%s"
}

data "ibm_account" "testacc_acc" {
    org_guid = data.ibm_org.testacc_ds_org.id
}`, cfOrganization)

}
