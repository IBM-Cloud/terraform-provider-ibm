// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/IBM-Cloud/power-go-client/clients/instance"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/service/power"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMPINetworkAddressGroupMemberBasic(t *testing.T) {
	name := fmt.Sprintf("tf-nag-name-%d", acctest.RandIntRange(10, 100))
	cidr := "192.168.1.5/32"
	nagMemberRes := "ibm_pi_network_address_group_member.network_address_group_member"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPINetworkAddressGroupMemberConfigBasic(name, cidr),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMPINetworkAddressGroupMemberExists(nagMemberRes),
					resource.TestCheckResourceAttrSet(nagMemberRes, power.Arg_NetworkAddressGroupID),
					resource.TestCheckResourceAttr(nagMemberRes, power.Arg_Cidr, cidr),
					resource.TestCheckResourceAttrSet(nagMemberRes, power.Attr_Name),
				),
			},
		},
	})
}

func testAccCheckIBMPINetworkAddressGroupMemberConfigBasic(name, cidr string) string {
	return testAccCheckIBMPINetworkAddressGroupConfigBasic(name) + fmt.Sprintf(`
		resource "ibm_pi_network_address_group_member" "network_address_group_member" {
			pi_cloud_instance_id = "%[1]s"
			pi_cidr = "%[2]s"
			pi_network_address_group_id = ibm_pi_network_address_group.network_address_group.network_address_group_id
		}`, acc.Pi_cloud_instance_id, cidr)
}

func testAccCheckIBMPINetworkAddressGroupMemberExists(n string) resource.TestCheckFunc {

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
		nsgClient := instance.NewIBMPINetworkAddressGroupClient(context.Background(), sess, cloudInstanceID)
		_, err = nsgClient.Get(nsgID)
		if err != nil {
			return err
		}
		return nil
	}
}
