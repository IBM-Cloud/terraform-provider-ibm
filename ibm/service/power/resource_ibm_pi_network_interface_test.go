// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/power"
)

func TestAccIBMPINetworkInterfaceBasic(t *testing.T) {
	name := fmt.Sprintf("tf-pi-name-%d", acctest.RandIntRange(10, 100))
	netInterRes := "ibm_pi_network_interface.network_interface"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPINetworkInterfaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkInterfaceConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPINetworkInterfaceExists(netInterRes),
					resource.TestCheckResourceAttr(netInterRes, power.Arg_NetworkID, acc.Pi_network_id),
					resource.TestCheckResourceAttrSet(netInterRes, power.Arg_NetworkID),
					resource.TestCheckResourceAttr(netInterRes, power.Attr_Name, name),
				),
			},
		},
	})
}

func TestAccIBMPINetworkInterfaceAllArgs(t *testing.T) {
	name := fmt.Sprintf("tf-pi-name-%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf-pi-name-update-%d", acctest.RandIntRange(10, 100))
	userTags := `["tf-ni-tag-1", "tf-ni-tag-2"]`
	userTagsUpdated := `["tf-ni-tag-1","tf-ni-tag-2", "tf-ni-tag-3"]`
	netInterRes := "ibm_pi_network_interface.network_interface"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPINetworkInterfaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkInterfaceConfig(name, userTags),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPINetworkInterfaceExists(netInterRes),
					resource.TestCheckResourceAttr(netInterRes, power.Attr_Name, name),
					resource.TestCheckResourceAttrSet(netInterRes, power.Arg_NetworkID),
					resource.TestCheckResourceAttrSet(netInterRes, power.Attr_IPAddress),
				),
			},
			{
				Config: testAccCheckIBMPINetworkInterfaceConfig(nameUpdate, userTagsUpdated),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(netInterRes, power.Attr_Name, nameUpdate),
					resource.TestCheckResourceAttrSet(netInterRes, power.Arg_NetworkID),
					resource.TestCheckResourceAttrSet(netInterRes, power.Attr_IPAddress),
				),
			},
		},
	})
}

func testAccCheckIBMPINetworkInterfaceConfigBasic(name string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_network_interface" "network_interface" {
			pi_cloud_instance_id = "%[1]s"
			pi_name = "%[2]s"
			pi_network_id = "%[3]s"
		}`, acc.Pi_cloud_instance_id, name, acc.Pi_network_id)
}

func testAccCheckIBMPINetworkInterfaceConfig(name, userTags string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_network_interface" "network_interface" {
			pi_cloud_instance_id = "%[1]s"
			pi_network_id = "%[2]s"
			pi_name = "%[3]s"
			pi_user_tags = %[4]s
		}`, acc.Pi_cloud_instance_id, acc.Pi_network_id, name, userTags)
}

func testAccCheckIBMPINetworkInterfaceExists(n string) resource.TestCheckFunc {
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
		networkClient := instance.NewIBMPINetworkClient(context.Background(), sess, parts[0])

		_, err = networkClient.GetNetworkInterface(parts[1], parts[2])
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckIBMPINetworkInterfaceDestroy(s *terraform.State) error {
	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_network_interface" {
			continue
		}
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		networkClient := instance.NewIBMPINetworkClient(context.Background(), sess, parts[0])
		_, err = networkClient.GetNetworkInterface(parts[1], parts[2])
		if err == nil {
			return fmt.Errorf("pi_network_interface still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}
