// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccSatelliteClusterGroupDataSourceBasic(t *testing.T) {
	clusterGroupName := "tf-satellite-clustergroup-1"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSatelliteClusterGroupDataSourceConfig(clusterGroupName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_satellite_config_clustergroup.read_clustergroup", "uuid"),
					resource.TestCheckResourceAttrSet("data.ibm_satellite_config_clustergroup.read_clustergroup", "created"),
				),
			},
		},
	})
}

func testAccCheckSatelliteClusterGroupDataSourceConfig(clusterGroupName string) string {
	return fmt.Sprintf(`

	provider "ibm" {
		region = "us-east"
	}

	resource "ibm_satellite_config_clustergroup" "group1" {
		name = "%s"
	}

	data "ibm_satellite_config_clustergroup" "read_clustergroup" {
		name = ibm_satellite_config_clustergroup.group1.name
	}

`, clusterGroupName)
}
