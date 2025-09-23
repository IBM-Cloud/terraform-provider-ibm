// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func TestAccIbmIsPlacementGroupBasic(t *testing.T) {
	var conf vpcv1.PlacementGroup
	strategy := "host_spread"
	strategyUpdate := "power_spread"
	name := fmt.Sprintf("tf-pg-name%d", acctest.RandIntRange(10, 100))
	nameupdate := fmt.Sprintf("tf-pg-name%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmIsPlacementGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIsPlacementGroupConfigBasic(strategy, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIsPlacementGroupExists("ibm_is_placement_group.is_placement_group", conf),
					resource.TestCheckResourceAttr("ibm_is_placement_group.is_placement_group", "strategy", strategy),
				),
			},
			{
				Config: testAccCheckIbmIsPlacementGroupConfigBasic(strategyUpdate, nameupdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_placement_group.is_placement_group", "strategy", strategyUpdate),
				),
			},
		},
	})
}

func TestAccIbmIsPlacementGroupAllArgs(t *testing.T) {
	var conf vpcv1.PlacementGroup
	strategy := "host_spread"
	name := fmt.Sprintf("tf-pg-name%d", acctest.RandIntRange(10, 100))
	strategyUpdate := "power_spread"
	nameUpdate := fmt.Sprintf("tf-pg-name-%d", acctest.RandIntRange(10, 100))
	tag1 := "stageplgrp"
	tag2 := "intplgrp"

	tagupdate1 := "prodplgrp"
	tagupdate2 := "devplgrp"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmIsPlacementGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIsPlacementGroupConfig(strategy, name, tag1, tag2),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIsPlacementGroupExists("ibm_is_placement_group.is_placement_group", conf),
					resource.TestCheckResourceAttr("ibm_is_placement_group.is_placement_group", "strategy", strategy),
					resource.TestCheckResourceAttr("ibm_is_placement_group.is_placement_group", "name", name),
					resource.TestCheckResourceAttr("ibm_is_placement_group.is_placement_group", "tags.0", tag2),
					resource.TestCheckResourceAttr("ibm_is_placement_group.is_placement_group", "tags.1", tag1),
				),
			},
			{
				Config: testAccCheckIbmIsPlacementGroupConfig(strategyUpdate, nameUpdate, tagupdate1, tagupdate2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_is_placement_group.is_placement_group", "strategy", strategyUpdate),
					resource.TestCheckResourceAttr("ibm_is_placement_group.is_placement_group", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_is_placement_group.is_placement_group", "tags.0", tagupdate2),
					resource.TestCheckResourceAttr("ibm_is_placement_group.is_placement_group", "tags.1", tagupdate1),
				),
			},
			{
				ResourceName:      "ibm_is_placement_group.is_placement_group",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmIsPlacementGroupConfigBasic(strategy, name string) string {
	return fmt.Sprintf(`

		resource "ibm_is_placement_group" "is_placement_group" {
			strategy = "%s"
			name = "%s"
			
		}
	`, strategy, name)
}

func testAccCheckIbmIsPlacementGroupConfig(strategy, name, tag1, tag2 string) string {
	return fmt.Sprintf(`
		data "ibm_resource_group" "default" {
			is_default=true
		}
		resource "ibm_is_placement_group" "is_placement_group" {
			strategy = "%s"
			name = "%s"
			resource_group = data.ibm_resource_group.default.id
			tags = ["%s", "%s"]
		}
	`, strategy, name, tag1, tag2)
}

func testAccCheckIbmIsPlacementGroupExists(n string, obj vpcv1.PlacementGroup) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		getPlacementGroupOptions := &vpcv1.GetPlacementGroupOptions{}

		getPlacementGroupOptions.SetID(rs.Primary.ID)

		placementGroup, _, err := vpcClient.GetPlacementGroup(getPlacementGroupOptions)
		if err != nil {
			return err
		}

		obj = *placementGroup
		return nil
	}
}

func testAccCheckIbmIsPlacementGroupDestroy(s *terraform.State) error {
	vpcClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_placement_group" {
			continue
		}

		getPlacementGroupOptions := &vpcv1.GetPlacementGroupOptions{}

		getPlacementGroupOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := vpcClient.GetPlacementGroup(getPlacementGroupOptions)

		if err == nil {
			return fmt.Errorf("PlacementGroup still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error checking for PlacementGroup (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
