// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/platform-services-go-sdk/atrackerv1"
)

func TestAccIBMAtrackerTargetBasic(t *testing.T) {
	var conf atrackerv1.Target

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMAtrackerTargetDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMAtrackerTargetConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAtrackerTargetExists("ibm_atracker_target.atracker_target", conf),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_atracker_target.atracker_target",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMAtrackerTargetConfigBasic() string {
	return fmt.Sprintf(`

		resource "ibm_atracker_target" "atracker_target" {
		}
	`)
}

func testAccCheckIBMAtrackerTargetExists(n string, obj atrackerv1.Target) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		atrackerClient, err := testAccProvider.Meta().(ClientSession).AtrackerV1()
		if err != nil {
			return err
		}

		getTargetOptions := &atrackerv1.GetTargetOptions{}

		parts, err := sepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getTargetOptions.SetID(parts[0])
		getTargetOptions.SetID(parts[1])

		target, _, err := atrackerClient.GetTarget(getTargetOptions)
		if err != nil {
			return err
		}

		obj = *target
		return nil
	}
}

func testAccCheckIBMAtrackerTargetDestroy(s *terraform.State) error {
	atrackerClient, err := testAccProvider.Meta().(ClientSession).AtrackerV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_atracker_target" {
			continue
		}

		getTargetOptions := &atrackerv1.GetTargetOptions{}

		parts, err := sepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getTargetOptions.SetID(parts[0])
		getTargetOptions.SetID(parts[1])

		// Try to find the key
		_, response, err := atrackerClient.GetTarget(getTargetOptions)

		if err == nil {
			return fmt.Errorf("Activity Tracking Target still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for Activity Tracking Target (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
