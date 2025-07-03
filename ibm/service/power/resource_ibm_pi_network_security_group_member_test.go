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

func TestAccIBMPINetworkSecurityGroupMemberBasic(t *testing.T) {

	typeVar := "network-interface"
	name := fmt.Sprintf("tf-nsg-name-%d", acctest.RandIntRange(10, 100))
	network_name := fmt.Sprintf("tf-nsg-network-%d", acctest.RandIntRange(10, 100))
	network_interface_name := fmt.Sprintf("tf-nsg-network-int-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkSecurityGroupMemberConfigBasic(name, typeVar, network_name, network_interface_name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPINetworkSecurityGroupMemberExists("ibm_pi_network_security_group_member.network_security_group_member"),
					resource.TestCheckResourceAttrSet("ibm_pi_network_security_group_member.network_security_group_member", power.Arg_NetworkSecurityGroupID),
					resource.TestCheckResourceAttr("ibm_pi_network_security_group_member.network_security_group_member", "pi_type", typeVar),
					resource.TestCheckResourceAttrSet("ibm_pi_network_security_group_member.network_security_group_member", power.Attr_Name),
				),
			},
		},
	})
}

func testAccCheckIBMPINetworkSecurityGroupMemberConfigBasic(name, typeVar, network_name, network_interface_name string) string {
	return testAccCheckIBMPINetworkSecurityGroupConfigBasic(name) + fmt.Sprintf(`
		resource "ibm_pi_network" "network" {
			pi_cloud_instance_id = "%[1]s"
			pi_network_name      = "%[3]s"
			pi_network_type      = "vlan"
			pi_cidr              = "192.168.17.0/24"
		}
		resource "ibm_pi_network_interface" "network_interface" {
			pi_cloud_instance_id = "%[1]s"
			pi_name = "%[4]s"
			pi_network_id = ibm_pi_network.network.network_id
			depends_on = [ibm_pi_network.network]
		}
		resource "ibm_pi_network_security_group_member" "network_security_group_member" {
			pi_cloud_instance_id = "%[1]s"
			pi_network_security_group_id = ibm_pi_network_security_group.network_security_group.network_security_group_id
			pi_target = ibm_pi_network_interface.network_interface.network_interface_id
			pi_type = "%[2]s"
		}`, acc.Pi_cloud_instance_id, typeVar, network_name, network_interface_name)
}

func testAccCheckIBMPINetworkSecurityGroupMemberExists(n string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return flex.FmtErrorf("Not found: %s", n)
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
