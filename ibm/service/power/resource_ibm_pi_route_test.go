// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMPIRouteBasic(t *testing.T) {
	name := fmt.Sprintf("tf-pi-route-%d", acctest.RandIntRange(10, 100))
	routeRes := "ibm_pi_route.route"
	nextHop := "192.112.111.1"
	destination := "192.116.111.1"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIRouteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIRouteBasicConfig(name, nextHop, destination),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIRouteExists(routeRes),
					resource.TestCheckResourceAttr(routeRes, "pi_name", name),
					resource.TestCheckResourceAttrSet(routeRes, "id"),
					resource.TestCheckResourceAttrSet(routeRes, "route_id"),
					resource.TestCheckResourceAttr(routeRes, "pi_next_hop", nextHop),
					resource.TestCheckResourceAttr(routeRes, "pi_destination", destination),
					resource.TestCheckResourceAttr(routeRes, "pi_destination_type", "ipv4-address"),
					resource.TestCheckResourceAttr(routeRes, "pi_next_hop_type", "ipv4-address"),
					resource.TestCheckResourceAttr(routeRes, "pi_enabled", "false"),
					resource.TestCheckResourceAttr(routeRes, "pi_advertise", "enable"),
					resource.TestCheckResourceAttrSet(routeRes, "state"),
				),
			},
		},
	})
}

func testAccCheckIBMPIRouteBasicConfig(name string, nextHop string, destination string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_route" "route" {
			pi_cloud_instance_id = "%[1]s"
			pi_name              = "%[2]s"
			pi_next_hop          = "%[3]s"
			pi_destination       = "%[4]s"
		}
	`, acc.Pi_cloud_instance_id, name, nextHop, destination)
}

func TestAccIBMPIRouteUpdate(t *testing.T) {
	name := fmt.Sprintf("tf-pi-route-%d", acctest.RandIntRange(10, 100))
	routeRes := "ibm_pi_route.route"
	nextHop := "192.112.111.1"
	initialDestination := "192.116.111.1"
	updatedDestination := "192.115.111.1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIRouteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIRouteUpdateConfig(name, nextHop, initialDestination),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIRouteExists(routeRes),
					resource.TestCheckResourceAttr(routeRes, "pi_name", name),
					resource.TestCheckResourceAttr(routeRes, "pi_destination", initialDestination),
				),
			},
			{
				Config: testAccCheckIBMPIRouteUpdateConfig(name, nextHop, updatedDestination),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIRouteExists(routeRes),
					resource.TestCheckResourceAttr(routeRes, "pi_name", name),
					resource.TestCheckResourceAttr(routeRes, "pi_destination", updatedDestination),
				),
			},
		},
	})
}

func testAccCheckIBMPIRouteUpdateConfig(name string, nextHop string, destination string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_route" "route" {
			pi_cloud_instance_id = "%[1]s"
			pi_name              = "%[2]s"
			pi_next_hop          = "%[3]s"
			pi_destination       = "%[4]s"
			pi_destination_type  = "ipv4-address"
			pi_next_hop_type     = "ipv4-address"
		}
	`, acc.Pi_cloud_instance_id, name, nextHop, destination)
}

func TestAccIBMPIRouteAllArgs(t *testing.T) {
	name := fmt.Sprintf("tf-pi-route-%d", acctest.RandIntRange(10, 100))
	routeRes := "ibm_pi_route.route"
	nextHop := "192.112.111.1"
	destination := "192.116.111.1"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIRouteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIRouteAllArgsConfig(name, nextHop, destination),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIRouteExists(routeRes),
					resource.TestCheckResourceAttr(routeRes, "pi_name", name),
					resource.TestCheckResourceAttrSet(routeRes, "id"),
					resource.TestCheckResourceAttrSet(routeRes, "route_id"),
					resource.TestCheckResourceAttr(routeRes, "pi_next_hop", nextHop),
					resource.TestCheckResourceAttr(routeRes, "pi_destination", destination),
					resource.TestCheckResourceAttr(routeRes, "pi_advertise", "disable"),
					resource.TestCheckResourceAttr(routeRes, "pi_action", "deliver"),
					resource.TestCheckResourceAttr(routeRes, "pi_destination_type", "ipv4-address"),
					resource.TestCheckResourceAttr(routeRes, "pi_enabled", "true"),
					resource.TestCheckResourceAttr(routeRes, "pi_next_hop_type", "ipv4-address"),
					resource.TestCheckResourceAttrSet(routeRes, "state"),
				),
			},
		},
	})
}

func testAccCheckIBMPIRouteAllArgsConfig(name string, nextHop string, destination string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_route" "route" {
			pi_cloud_instance_id    = "%[1]s"
			pi_name                 = "%[2]s"
			pi_next_hop             = "%[3]s"
			pi_destination          = "%[4]s"
			pi_advertise            = "disable"
			pi_action               = "deliver"
			pi_destination_type     = "ipv4-address"
			pi_enabled              = true
			pi_next_hop_type        = "ipv4-address"
		}
	`, acc.Pi_cloud_instance_id, name, nextHop, destination)
}

func TestAccIBMPIRouteUserTags(t *testing.T) {
	name := fmt.Sprintf("tf-pi-route-%d", acctest.RandIntRange(10, 100))
	routeRes := "ibm_pi_route.route"
	nextHop := "192.112.111.1"
	destination := "192.116.111.1"
	userTagsString := `["env:dev", "test_tag"]`
	userTagsUpdatedString := `["env:dev", "test_tag", "ibm"]`
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPIRouteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIRouteUserTagsConfig(name, nextHop, destination, userTagsString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIRouteExists(routeRes),
					resource.TestCheckResourceAttr(routeRes, "pi_name", name),
					resource.TestCheckResourceAttrSet(routeRes, "crn"),
					resource.TestCheckResourceAttr(routeRes, "pi_user_tags.#", "2"),
					resource.TestCheckTypeSetElemAttr(routeRes, "pi_user_tags.*", "env:dev"),
					resource.TestCheckTypeSetElemAttr(routeRes, "pi_user_tags.*", "test_tag"),
				),
			},
			{
				Config: testAccCheckIBMPIRouteUserTagsConfig(name, nextHop, destination, userTagsUpdatedString),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPIRouteExists(routeRes),
					resource.TestCheckResourceAttr(routeRes, "pi_name", name),
					resource.TestCheckResourceAttrSet(routeRes, "crn"),
					resource.TestCheckResourceAttr(routeRes, "pi_user_tags.#", "3"),
					resource.TestCheckTypeSetElemAttr(routeRes, "pi_user_tags.*", "env:dev"),
					resource.TestCheckTypeSetElemAttr(routeRes, "pi_user_tags.*", "test_tag"),
					resource.TestCheckTypeSetElemAttr(routeRes, "pi_user_tags.*", "ibm"),
				),
			},
		},
	})
}

func testAccCheckIBMPIRouteUserTagsConfig(name string, nextHop string, destination string, userTagsString string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_route" "route" {
			pi_cloud_instance_id = "%[1]s"
			pi_name              = "%[2]s"
			pi_next_hop          = "%[3]s"
			pi_destination       = "%[4]s"
			pi_user_tags         =  %[5]s  
		}
	`, acc.Pi_cloud_instance_id, name, nextHop, destination, userTagsString)
}

func testAccCheckIBMPIRouteDestroy(s *terraform.State) error {
	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_route" {
			continue
		}
		parts, _ := flex.IdParts(rs.Primary.ID)
		cloudinstanceid := parts[0]
		routeC := instance.NewIBMPIRouteClient(context.Background(), sess, cloudinstanceid)
		_, err = routeC.Get(parts[1])
		if err == nil {
			return fmt.Errorf("PI route still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckIBMPIRouteExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
		if err != nil {
			return err
		}
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		cloudinstanceid := parts[0]
		client := instance.NewIBMPIRouteClient(context.Background(), sess, cloudinstanceid)

		_, err = client.Get(parts[1])
		if err != nil {
			return err
		}

		return nil
	}
}
