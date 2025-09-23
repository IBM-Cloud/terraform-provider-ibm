// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/service/power"
)

func TestAccIBMPINetworkAddressGroupBasic(t *testing.T) {
	name := fmt.Sprintf("tf-nag-name-%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf-nag-name-update-%d", acctest.RandIntRange(10, 100))
	nagRes := "ibm_pi_network_address_group.network_address_group"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPINetworkAddressGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkAddressGroupConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPINetworkAddressGroupExists(nagRes),
					resource.TestCheckResourceAttrSet(nagRes, power.Attr_ID),
					resource.TestCheckResourceAttrSet(nagRes, power.Attr_NetworkAddressGroupID),
					resource.TestCheckResourceAttr(nagRes, power.Arg_Name, name),
				),
			},
			{
				Config: testAccCheckIBMPINetworkAddressGroupConfigBasic(nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(nagRes, power.Arg_Name, nameUpdate),
					resource.TestCheckResourceAttrSet(nagRes, power.Attr_ID),
					resource.TestCheckResourceAttrSet(nagRes, power.Attr_NetworkAddressGroupID),
				),
			},
		},
	})
}

func testAccCheckIBMPINetworkAddressGroupConfigBasic(name string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_network_address_group" "network_address_group" {
			pi_cloud_instance_id = "%[1]s"	
			pi_name = "%[2]s"
		}
	`, acc.Pi_cloud_instance_id, name)
}

func TestAccIBMPINetworkAddressGroupUserTags(t *testing.T) {
	name := fmt.Sprintf("tf-nag-name-%d", acctest.RandIntRange(10, 100))
	nagRes := "ibm_pi_network_address_group.network_address_group"
	userTagsString := `["env:dev", "test_tag"]`
	userTagsStringUpdated := `["env:dev", "test_tag", "ibm"]`
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPINetworkAddressGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkAddressGroupConfigUserTags(name, userTagsString),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPINetworkAddressGroupExists(nagRes),
					resource.TestCheckResourceAttrSet(nagRes, power.Attr_ID),
					resource.TestCheckResourceAttrSet(nagRes, power.Attr_CRN),
					resource.TestCheckResourceAttr(nagRes, power.Arg_UserTags+".#", "2"),
					resource.TestCheckTypeSetElemAttr(nagRes, power.Arg_UserTags+".*", "env:dev"),
					resource.TestCheckTypeSetElemAttr(nagRes, power.Arg_UserTags+".*", "test_tag"),
				),
			},
			{
				Config: testAccCheckIBMPINetworkAddressGroupConfigUserTags(name, userTagsStringUpdated),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(nagRes, power.Attr_ID),
					resource.TestCheckResourceAttrSet(nagRes, power.Attr_CRN),
					resource.TestCheckResourceAttr(nagRes, power.Arg_UserTags+".#", "3"),
					resource.TestCheckTypeSetElemAttr(nagRes, power.Arg_UserTags+".*", "env:dev"),
					resource.TestCheckTypeSetElemAttr(nagRes, power.Arg_UserTags+".*", "test_tag"),
					resource.TestCheckTypeSetElemAttr(nagRes, power.Arg_UserTags+".*", "ibm"),
				),
			},
		},
	})
}

func testAccCheckIBMPINetworkAddressGroupConfigUserTags(name string, userTagsString string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_network_address_group" "network_address_group" {
			pi_cloud_instance_id = "%[1]s"	
			pi_name = "%[2]s"
			pi_user_tags = %[3]s
		}
	`, acc.Pi_cloud_instance_id, name, userTagsString)
}

func testAccCheckIBMPINetworkAddressGroupExists(n string) resource.TestCheckFunc {

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
		cloudInstanceID, nagID, err := splitID(rs.Primary.ID)
		if err != nil {
			return err
		}
		nagC := instance.NewIBMPINetworkAddressGroupClient(context.Background(), sess, cloudInstanceID)
		_, err = nagC.Get(nagID)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckIBMPINetworkAddressGroupDestroy(s *terraform.State) error {
	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_network_address_group" {
			continue
		}

		cloudInstanceID, nagID, err := splitID(rs.Primary.ID)
		if err != nil {
			return err
		}
		nagC := instance.NewIBMPINetworkAddressGroupClient(context.Background(), sess, cloudInstanceID)
		_, err = nagC.Get(nagID)
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), power.NotFound) {
				return nil
			}
		}
		return fmt.Errorf("network addess group still exists: %s", rs.Primary.ID)
	}
	return nil
}
