// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"errors"
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMISReservationActivate_basic(t *testing.T) {
	var reservation string
	name := fmt.Sprintf("tfresa-name-%d", acctest.RandIntRange(10, 100))
	zone := "us-east-3"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckImage(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: checkReservationActivateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISReservationActivateConfig(name, zone),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISReservationActivateExists("ibm_is_reservation.isExampleReservation", reservation),
					resource.TestCheckResourceAttr(
						"ibm_is_reservation.isExampleReservation", "name", name),
					resource.TestCheckResourceAttr(
						"ibm_is_reservation.isExampleReservation", "status", "active"),
				),
			},
		},
	})
}

func checkReservationActivateDestroy(s *terraform.State) error {

	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_reservation" {
			continue
		}

		getresoptions := &vpcv1.GetReservationOptions{
			ID: &rs.Primary.ID,
		}
		_, _, err := sess.GetReservation(getresoptions)
		if err == nil {
			fmt.Printf("Reservation %s still exists: %s", rs.Primary.ID, err.Error())
		}
	}

	return nil
}

func testAccCheckIBMISReservationActivateExists(n, image string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getimgoptions := &vpcv1.GetReservationOptions{
			ID: &rs.Primary.ID,
		}
		foundReservation, _, err := sess.GetReservation(getimgoptions)
		if err != nil {
			return err
		}
		image = *foundReservation.ID

		return nil
	}
}

func testAccCheckIBMISReservationActivateConfig(name, zone string) string {
	return fmt.Sprintf(`
	resource "ibm_is_reservation" "isExampleReservation" {
		capacity {
			total = 10
		  }
		  committed_use {
			term = "one_year"
		  }
		profile {
			name = "ba2-2x8"
			resource_type = "instance_profile"
		  }
		name = "%s"
		zone = "%s"
	}
	resource ibm_is_reservation_activate activate {
		reservation = ibm_is_reservation.isExampleReservation.id
	}
	`, name, zone)
}
