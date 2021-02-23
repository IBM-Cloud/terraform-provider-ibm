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

func TestAccIBMSpaceDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSpaceDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_space.testacc_ds_space", "org", cfOrganization),
					resource.TestCheckResourceAttr("data.ibm_space.testacc_ds_space", "space", cfSpace),
				),
			},
		},
	})
}

func testAccCheckIBMSpaceDataSourceConfig() string {
	return fmt.Sprintf(`
data "ibm_space" "testacc_ds_space" {
    org = "%s"
	space = "%s"
}`, cfOrganization, cfSpace)

}
