// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISLBProfileDatasource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{

				Config: testDSCheckIBMISLBProfileBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_is_lb_profile.test_profile", "name", "network-fixed"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_profile.test_profile", "family", "Network"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_profile.test_profile", "route_mode_supported", "true"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profile.test_profile", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profile.test_profile", "udp_supported"),
				),
			},
		},
	})
}
func testDSCheckIBMISLBProfileBasicConfig() string {
	return fmt.Sprintf(`
	data "ibm_is_lb_profile" "test_profile" {
		name = "network-fixed"
	} `)
}
