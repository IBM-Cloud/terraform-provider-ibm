// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIsReservationsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsReservationsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_reservations.example", "reservations.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reservations.example", "reservations.0.affinity_policy"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reservations.example", "reservations.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reservations.example", "reservations.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reservations.example", "reservations.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reservations.example", "reservations.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reservations.example", "reservations.0.lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reservations.example", "reservations.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reservations.example", "reservations.0.status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reservations.example", "reservations.0.zone"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reservations.example", "reservations.0.resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reservations.example", "reservations.0.profile.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reservations.example", "reservations.0.resource_group.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsReservationsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_is_reservations" "example" {
		}
	`)
}
