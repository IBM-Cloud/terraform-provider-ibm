// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIsLbPoolMembersDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tflbpm-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tflbpmc-name-%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tfcreate%d", acctest.RandIntRange(10, 100))
	poolName := fmt.Sprintf("tflbpoolc%d", acctest.RandIntRange(10, 100))
	port := "8080"
	address := "127.0.0.1"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsLbPoolMembersDataSourceConfigBasic(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, name, poolName, port, address),
				Check: resource.ComposeTestCheckFunc(
					// resource.TestCheckResourceAttrSet("data.ibm_is_lb_pool_members.is_lb_pool_members", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pool_members.is_lb_pool_members", "lb"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pool_members.is_lb_pool_members", "pool"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_pool_members.is_lb_pool_members", "members.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsLbPoolMembersDataSourceConfigBasic(vpcname, subnetname, zone, cidr, name, poolName, port, address string) string {
	return testAccCheckIBMISLBPoolMemberConfig(vpcname, subnetname, zone, cidr, name, poolName, port, address) + fmt.Sprintf(`
        data "ibm_is_lb_pool_members" "is_lb_pool_members" {
            lb = "${ibm_is_lb.testacc_LB.id}"
            pool = "${element(split("/",ibm_is_lb_pool.testacc_lb_pool.id),1)}"
        }
    `)
}
