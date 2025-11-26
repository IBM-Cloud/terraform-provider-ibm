// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
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
					resource.TestCheckResourceAttr("data.ibm_is_lb_profile.test_profile", "family", "network"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_profile.test_profile", "route_mode_supported", "true"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profile.test_profile", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profile.test_profile", "targetable_resource_types.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profile.test_profile", "udp_supported"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profile.test_profile", "access_modes.0.values.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profile.test_profile", "targetable_load_balancer_profiles.#"),
				),
			},
		},
	})
}
func TestAccIBMISLBProfileDatasource_failsafepolicyactions(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{

				Config: testDSCheckIBMISLBProfileBasicConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_is_lb_profile.test_profile", "name", "network-fixed"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_profile.test_profile", "family", "network"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_profile.test_profile", "route_mode_supported", "true"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profile.test_profile", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profile.test_profile", "udp_supported"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profile.test_profile", "failsafe_policy_actions.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profile.test_profile", "failsafe_policy_actions.0.default"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profile.test_profile", "failsafe_policy_actions.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profile.test_profile", "failsafe_policy_actions.0.values.#"),
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
