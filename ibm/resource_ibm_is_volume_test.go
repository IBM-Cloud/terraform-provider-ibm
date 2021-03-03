// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/IBM/vpc-go-sdk/vpcclassicv1"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMISVolume_basic(t *testing.T) {
	var vol string
	name := fmt.Sprintf("tf-vol-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tf-vol-upd-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISVolumeDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISVolumeConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVolumeExists("ibm_is_volume.storage", vol),
					resource.TestCheckResourceAttr(
						"ibm_is_volume.storage", "name", name),
				),
			},

			resource.TestStep{
				Config: testAccCheckIBMISVolumeConfig(name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVolumeExists("ibm_is_volume.storage", vol),
					resource.TestCheckResourceAttr(
						"ibm_is_volume.storage", "name", name1),
				),
			},
		},
	})
}

func testAccCheckIBMISVolumeDestroy(s *terraform.State) error {
	userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()

	if userDetails.generation == 1 {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcClassicV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_vol" {
				continue
			}

			getvolumeoptions := &vpcclassicv1.GetVolumeOptions{
				ID: &rs.Primary.ID,
			}
			_, _, err := sess.GetVolume(getvolumeoptions)

			if err == nil {
				return fmt.Errorf("Volume still exists: %s", rs.Primary.ID)
			}
		}
	} else {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_vol" {
				continue
			}

			getvolumeoptions := &vpcv1.GetVolumeOptions{
				ID: &rs.Primary.ID,
			}
			_, _, err := sess.GetVolume(getvolumeoptions)

			if err == nil {
				return fmt.Errorf("Volume still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}

func testAccCheckIBMISVolumeExists(n, volID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()

		if userDetails.generation == 1 {
			sess, _ := testAccProvider.Meta().(ClientSession).VpcClassicV1API()
			getvolumeoptions := &vpcclassicv1.GetVolumeOptions{
				ID: &rs.Primary.ID,
			}
			foundvol, _, err := sess.GetVolume(getvolumeoptions)
			if err != nil {
				return err
			}
			volID = *foundvol.ID
		} else {
			sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
			getvolumeoptions := &vpcv1.GetVolumeOptions{
				ID: &rs.Primary.ID,
			}
			foundvol, _, err := sess.GetVolume(getvolumeoptions)
			if err != nil {
				return err
			}
			volID = *foundvol.ID
		}
		return nil
	}
}

func testAccCheckIBMISVolumeConfig(name string) string {
	return fmt.Sprintf(
		`resource "ibm_is_volume" "storage"{
    name = "%s"
    profile = "10iops-tier"
    zone = "us-south-3"
    # capacity= 200
}`, name)

}
