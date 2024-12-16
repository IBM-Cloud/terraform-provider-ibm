// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package satellite_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/IBM-Cloud/container-services-go-sdk/satellitelinkv1"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIbmSatelliteLinkBasic(t *testing.T) {
	var conf satellitelinkv1.Location
	locationID := fmt.Sprintf("tf-location-%d", acctest.RandIntRange(10, 100))
	locationIDUpdate := fmt.Sprintf("tf-location-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSatelliteLinkDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSatelliteLinkConfig(locationID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSatelliteLinkExists("ibm_satellite_link.satellite_link", conf),
					resource.TestCheckResourceAttr("ibm_satellite_link.satellite_link", "location", locationID),
				),
			},
			{
				Config: testAccCheckIbmSatelliteLinkConfig(locationIDUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_satellite_link.satellite_link", "location", locationIDUpdate),
				),
			},
			{
				ResourceName:      "ibm_satellite_link.satellite_link",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmSatelliteLinkConfig(locationID string) string {
	return fmt.Sprintf(`

		data "ibm_satellite_location" "location" {
			location = "%s"
	  	}

		resource "ibm_satellite_link" "satellite_link" {
			crn = data.ibm_satellite_location.location.crn
			location = "%s"
		}
	`, locationID, locationID)
}

func testAccCheckIbmSatelliteLinkExists(n string, obj satellitelinkv1.Location) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		satelliteLinkClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SatellitLinkClientSession()
		if err != nil {
			return err
		}

		getLinkOptions := &satellitelinkv1.GetLinkOptions{}

		getLinkOptions.SetLocationID(rs.Primary.ID)

		location, _, err := satelliteLinkClient.GetLink(getLinkOptions)
		if err != nil {
			return err
		}

		obj = *location
		return nil
	}
}

func testAccCheckIbmSatelliteLinkDestroy(s *terraform.State) error {
	satelliteLinkClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SatellitLinkClientSession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_satellite_link" {
			continue
		}

		getLinkOptions := &satellitelinkv1.GetLinkOptions{}

		getLinkOptions.SetLocationID(rs.Primary.ID)

		// Try to find the key
		_, response, err := satelliteLinkClient.GetLink(getLinkOptions)

		if err == nil {
			return fmt.Errorf("satellite_link still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error checking for satellite_link (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
