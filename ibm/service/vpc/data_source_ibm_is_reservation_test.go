// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISReservationDatasource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testDSCheckIBMISReservationConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_reservation.ds_res", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reservation.ds_res", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reservation.ds_res", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reservation.ds_res", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reservation.ds_res", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reservation.ds_res", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reservation.ds_res", "profile.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reservation.ds_res", "resource_group.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reservation.ds_res", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_reservation.ds_res", "zone"),
				),
			},
		},
	})
}

func testDSCheckIBMISReservationConfig() string {
	return fmt.Sprintf(`
		resource "ibm_is_reservation" "res" {
			capacity {
				total = 5
			  }
			  committed_use {
				term = "one_year"
			  }
			  profile {
				name = "ba2-2x8"
				resource_type = "instance_profile"
			  }
			  zone = "us-east-3"
			  name = "reservation-name"
		}
		data "ibm_is_reservation" "ds_res" {
		    name = ibm_is_reservation.res.id
		}`)
}
