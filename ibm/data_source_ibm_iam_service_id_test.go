// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIAMServiceIDDataSource_basic(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
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
			{
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
