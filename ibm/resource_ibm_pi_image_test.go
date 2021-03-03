// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	st "github.com/IBM-Cloud/power-go-client/clients/instance"
)

func TestAccIBMPIImagebasic(t *testing.T) {

	name := fmt.Sprintf("tf-pi-image-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMPIImageDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIImageConfig(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIImageExists("ibm_pi_image.power_image"),
					resource.TestCheckResourceAttr(
						"ibm_pi_image.power_image", "pi_image_name", name),
				),
			},
		},
	})
}
func testAccCheckIBMPIImageDestroy(s *terraform.State) error {

	sess, err := testAccProvider.Meta().(ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_image" {
			continue
		}
		parts, err := idParts(rs.Primary.ID)
		powerinstanceid := parts[0]
		networkC := st.NewIBMPIImageClient(sess, powerinstanceid)
		_, err = networkC.Get(parts[1], powerinstanceid)
		if err == nil {
			return fmt.Errorf("PI Image still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}
func testAccCheckIBMPIImageExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		fmt.Println("siv ", s.RootModule().Resources)
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, err := testAccProvider.Meta().(ClientSession).IBMPISession()
		if err != nil {
			return err
		}
		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		powerinstanceid := parts[0]
		client := st.NewIBMPIImageClient(sess, powerinstanceid)

		image, err := client.Get(parts[1], powerinstanceid)
		if err != nil {
			return err
		}
		parts[1] = *image.ImageID
		return nil

	}
}

func testAccCheckIBMPIImageConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_pi_image" "power_image" {
		pi_image_name       = "%s"
		pi_image_id         = "cfc02954-8f6f-4e6b-96ae-40b24c90bd54"
		pi_cloud_instance_id = "%s"
	  }
	`, name, pi_cloud_instance_id)
}
