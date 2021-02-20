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

func TestAccIBMIAMServiceIDDataSource_basic(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMServiceIDDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_iam_service_id.testacc_ds_service_id", "name", name),
					resource.TestCheckResourceAttr(
						"data.ibm_iam_service_id.testacc_ds_service_id", "service_ids.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMIAMServiceIDDataSource_same_name(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMServiceIDDataSourceSameName(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_iam_service_id.testacc_ds_service_id", "name", name),
					resource.TestCheckResourceAttr(
						"data.ibm_iam_service_id.testacc_ds_service_id", "service_ids.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIBMIAMServiceIDDataSourceConfig(name string) string {
	return fmt.Sprintf(`

resource "ibm_iam_service_id" "serviceID" {
  name        = "%s"
  description = "ServiceID for test"
}

data "ibm_iam_service_id" "testacc_ds_service_id" {
  name = ibm_iam_service_id.serviceID.name
}
`, name)

}

func testAccCheckIBMIAMServiceIDDataSourceSameName(name string) string {
	return fmt.Sprintf(`

resource "ibm_iam_service_id" "serviceID" {
  name        = "%s"
  description = "ServiceID for test"
}

resource "ibm_iam_service_id" "serviceID2" {
  name = "%s"
}

data "ibm_iam_service_id" "testacc_ds_service_id" {
  name = ibm_iam_service_id.serviceID.name
}`, name, name)

}
