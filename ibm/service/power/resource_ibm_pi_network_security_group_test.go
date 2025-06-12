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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/power"
)

func TestAccIBMPINetworkSecurityGroupBasic(t *testing.T) {
	name := fmt.Sprintf("tf-nsg-name-%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf-nsg-name-update-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPINetworkSecurityGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkSecurityGroupConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPINetworkSecurityGroupExists("ibm_pi_network_security_group.network_security_group"),
					resource.TestCheckResourceAttr("ibm_pi_network_security_group.network_security_group", power.Arg_Name, name),
				),
			},
			{
				Config: testAccCheckIBMPINetworkSecurityGroupConfigBasic(nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_pi_network_security_group.network_security_group", power.Arg_Name, nameUpdate),
				),
			},
		},
	})
}

func testAccCheckIBMPINetworkSecurityGroupConfigBasic(name string) string {
	return fmt.Sprintf(`
		resource "ibm_pi_network_security_group" "network_security_group" {
			pi_cloud_instance_id = "%[1]s"
			pi_name = "%[2]s"
			pi_user_tags = ["tag:test"]
		}`, acc.Pi_cloud_instance_id, name)
}

func testAccCheckIBMPINetworkSecurityGroupExists(n string) resource.TestCheckFunc {

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
		cloudInstanceID, nsgID, err := splitID(rs.Primary.ID)
		if err != nil {
			return err
		}
		nsgClient := instance.NewIBMIPINetworkSecurityGroupClient(context.Background(), sess, cloudInstanceID)
		_, err = nsgClient.Get(nsgID)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckIBMPINetworkSecurityGroupDestroy(s *terraform.State) error {
	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).IBMPISession()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pi_network_security_group" {
			continue
		}

		cloudInstanceID, nsgID, err := splitID(rs.Primary.ID)
		if err != nil {
			return err
		}
		nsgClient := instance.NewIBMIPINetworkSecurityGroupClient(context.Background(), sess, cloudInstanceID)
		_, err = nsgClient.Get(nsgID)
		if err == nil {
			return fmt.Errorf("network_security_group still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}
