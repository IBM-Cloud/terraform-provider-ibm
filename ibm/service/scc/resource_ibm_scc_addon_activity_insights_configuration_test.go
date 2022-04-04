// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/scc-go-sdk/v4/addonmanagerv1"
)

func TestAccIBMSccAddonActivityInsightsConfigurationBasic(t *testing.T) {
	var conf addonmanagerv1.ActivityInsightsStatusConfigOutput
	regionID := "us"
	status := "enable"
	regionIDUpdate := "us"
	statusUpdate := "disable"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccAddonActivityInsightsConfigurationConfigBasic(regionID, status),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSccAddonActivityInsightsConfigurationExists("ibm_scc_addon_activity_insights_configuration.scc_addon_activity_insights_configuration", conf),
					resource.TestCheckResourceAttr("ibm_scc_addon_activity_insights_configuration.scc_addon_activity_insights_configuration", "region_id", regionID),
					resource.TestCheckResourceAttr("ibm_scc_addon_activity_insights_configuration.scc_addon_activity_insights_configuration", "status", status),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMSccAddonActivityInsightsConfigurationConfigBasic(regionIDUpdate, statusUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_scc_addon_activity_insights_configuration.scc_addon_activity_insights_configuration", "region_id", regionIDUpdate),
					resource.TestCheckResourceAttr("ibm_scc_addon_activity_insights_configuration.scc_addon_activity_insights_configuration", "status", statusUpdate),
				),
			},
			resource.TestStep{
				ResourceName:            "ibm_scc_addon_activity_insights_configuration.scc_addon_activity_insights_configuration",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region_id"},
			},
		},
	})
}

func testAccCheckIBMSccAddonActivityInsightsConfigurationConfigBasic(regionID string, status string) string {
	return fmt.Sprintf(`

		resource "ibm_scc_addon_activity_insights_configuration" "scc_addon_activity_insights_configuration" {
			region_id = "%s"
			status = "%s"
		}
	`, regionID, status)
}

func testAccCheckIBMSccAddonActivityInsightsConfigurationExists(n string, obj addonmanagerv1.ActivityInsightsStatusConfigOutput) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		addonManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).AddonManagerV1()
		if err != nil {
			return err
		}

		addonManagerClient.AccountID = &parts[0]

		getActivityInsightStatusV2Options := &addonmanagerv1.GetActivityInsightStatusV2Options{}

		aiEnableAddOn, _, err := addonManagerClient.GetActivityInsightStatusV2(getActivityInsightStatusV2Options)
		if err != nil {
			return err
		}

		obj = *aiEnableAddOn
		return nil
	}
}
