/*
* IBM Confidential
* Object Code Only Source Materials
* 5747-SM3
* (c) Copyright IBM Corp. 2017,2021
*
* The source code for this program is not published or otherwise divested
* of its trade secrets, irrespective of what has been deposited with the
* U.S. Copyright Office.
 */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMOrgQuotaDataSource_basic(t *testing.T) {

	name := "qIBM"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMOrgQuotaDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_org_quota.testacc_ds_org_quota", "name", name),
				),
			},
		},
	})
}

func testAccCheckIBMOrgQuotaDataSourceConfig(name string) string {
	return fmt.Sprintf(`
	
data "ibm_org_quota" "testacc_ds_org_quota" {
    name = "%s"
}`, name)

}
