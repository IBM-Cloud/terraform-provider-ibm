// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccSatelliteClusterGroup_Basic(t *testing.T) {
	clusterGroupName := "tf-satellite-clustergroup-1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSatelliteClusterGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckSatelliteClusterGroupCreate(clusterGroupName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSatelliteClusterGroupExists("ibm_satellite_config_clustergroup.group1"),
					resource.TestCheckResourceAttr("ibm_satellite_config_clustergroup.group1", "name", clusterGroupName),
				),
			},
		},
	})
}

func testAccCheckSatelliteClusterGroupExists(n string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		satconClient, err := testAccProvider.Meta().(ClientSession).SatellitConfigClientSession()
		if err != nil {
			return err
		}

		satconGroupAPI := satconClient.Groups

		userDetails, err := testAccProvider.Meta().(ClientSession).BluemixUserDetails()
		if err != nil {
			return err
		}

		groupName := rs.Primary.ID
		_, err = satconGroupAPI.GroupByName(userDetails.userAccount, groupName)
		if err != nil {
			return fmt.Errorf("Satellite Cluster Group doesn't exist: %s", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckSatelliteClusterGroupDestroy(s *terraform.State) error {
	satconClient, err := testAccProvider.Meta().(ClientSession).SatellitConfigClientSession()
	if err != nil {
		return err
	}

	satconGroupAPI := satconClient.Groups

	userDetails, err := testAccProvider.Meta().(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_satellite_config_clustergroup" {
			continue
		}

		groupName := rs.Primary.ID
		_, err := satconGroupAPI.GroupByName(userDetails.userAccount, groupName)
		if err == nil {
			return fmt.Errorf("Satellite Cluster Group still exists: %s", rs.Primary.ID)
		}

	}
	return nil
}

func testAccCheckSatelliteClusterGroupCreate(clusterGroupName string) string {
	return fmt.Sprintf(`

	provider "ibm" {
		region = "us-east"
	}

	resource "ibm_satellite_config_clustergroup" "group1" {
		name = "%s"
	}

`, clusterGroupName)
}
