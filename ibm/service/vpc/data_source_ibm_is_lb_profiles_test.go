// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISLBProfilesDatasource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{

				Config: testDSCheckIBMISLBProfilesConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profiles.test_profiles", "lb_profiles.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profiles.test_profiles", "lb_profiles.0.access_modes.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profiles.test_profiles", "lb_profiles.0.access_modes.0.values.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profiles.test_profiles", "lb_profiles.0.targetable_load_balancer_profiles.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profiles.test_profiles", "lb_profiles.0.targetable_resource_types.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profiles.test_profiles", "lb_profiles.0.route_mode_supported"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profiles.test_profiles", "lb_profiles.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profiles.test_profiles", "lb_profiles.0.udp_supported"),
				),
			},
		},
	})
}
func TestAccIBMISLBProfilesDatasource_failsafepolicyactions(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{

				Config: testDSCheckIBMISLBProfilesFailsafepolicyactionsConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profiles.test_profiles", "lb_profiles.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profiles.test_profiles", "lb_profiles.0.failsafe_policy_actions.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profiles.test_profiles", "lb_profiles.0.failsafe_policy_actions.0.default"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profiles.test_profiles", "lb_profiles.0.failsafe_policy_actions.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profiles.test_profiles", "lb_profiles.0.failsafe_policy_actions.0.values.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profiles.test_profiles", "lb_profiles.0.family"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profiles.test_profiles", "lb_profiles.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profiles.test_profiles", "lb_profiles.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profiles.test_profiles", "lb_profiles.0.route_mode_supported"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profiles.test_profiles", "lb_profiles.0.route_mode_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profiles.test_profiles", "lb_profiles.0.udp_supported"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profiles.test_profiles", "lb_profiles.0.udp_supported_type"),
				),
			},
		},
	})
}
func TestAccIBMISLBProfilesDatasource_filter(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{

				Config: testDSCheckIBMISLBProfilesFilterConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_is_lb_profiles.test_profiles", "lb_profiles.0.name", "network-fixed"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_profiles.test_profiles", "lb_profiles.0.family", "network"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_profiles.test_profiles", "lb_profiles.0.route_mode_supported", "true"),
					resource.TestCheckResourceAttr("data.ibm_is_lb_profiles.test_profiles", "lb_profiles.0.targetable_resource_types.0.values.0", "instance"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profiles.test_profiles", "lb_profiles.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_lb_profiles.test_profiles", "lb_profiles.0.udp_supported"),
				),
			},
		},
	})
}
func testDSCheckIBMISLBProfilesConfig() string {
	return fmt.Sprintf(`
	data "ibm_is_lb_profiles" "test_profiles" {
	} `)
}
func testDSCheckIBMISLBProfilesFilterConfig() string {
	return fmt.Sprintf(`
	data "ibm_is_lb_profiles" "test_profiles" {
		name = "network-fixed"
	} `)
}
func testDSCheckIBMISLBProfilesFailsafepolicyactionsConfig() string {
	return fmt.Sprintf(`
	data "ibm_is_lb_profiles" "test_profiles" {
	} `)
}
