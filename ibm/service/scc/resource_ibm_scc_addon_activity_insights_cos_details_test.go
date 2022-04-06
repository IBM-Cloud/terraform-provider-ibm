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

func TestAccIBMSccAddonActivityInsightsCosDetailsBasic(t *testing.T) {
	var conf addonmanagerv1.ActivityInsightsCosDetailsOutput
	regionID := "us"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSccAddonActivityInsightsCosDetailsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccAddonActivityInsightsCosDetailsConfigBasic(regionID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSccAddonActivityInsightsCosDetailsExists("ibm_scc_addon_activity_insights_cos_details.scc_addon_activity_insights_cos_details", conf),
					resource.TestCheckResourceAttrSet("ibm_scc_addon_activity_insights_cos_details.scc_addon_activity_insights_cos_details", "cos_details.#"),
				),
			},
			resource.TestStep{
				ResourceName:            "ibm_scc_addon_activity_insights_cos_details.scc_addon_activity_insights_cos_details",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"region_id"},
			},
		},
	})
}

func testAccCheckIBMSccAddonActivityInsightsCosDetailsConfigBasic(regionID string) string {
	return fmt.Sprintf(`

		resource "ibm_scc_addon_activity_insights_cos_details" "scc_addon_activity_insights_cos_details" {
			region_id = "%s"
			cos_details {
				cos_instance = "cos_instance"
				bucket_name = "bucket_name"
				description = "description"
				type = "activity_insights"
				cos_bucket_url = "cos_bucket_url"
			}
		}
	`, regionID)
}

func testAccCheckIBMSccAddonActivityInsightsCosDetailsExists(n string, obj addonmanagerv1.ActivityInsightsCosDetailsOutput) resource.TestCheckFunc {

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

		getActivityInsightsCosDetailsV2Options := &addonmanagerv1.GetActivityInsightsCosDetailsV2Options{}

		aiCosDetailsV2Output, _, err := addonManagerClient.GetActivityInsightsCosDetailsV2(getActivityInsightsCosDetailsV2Options)
		if err != nil {
			return err
		}

		for _, cosDetail := range aiCosDetailsV2Output.CosDetails {
			if cosDetail.ID == &parts[1] {
				cosDetails := addonmanagerv1.ActivityInsightsCosDetailsOutput{
					CosDetails: []addonmanagerv1.CosDetailsWithID{
						cosDetail,
					},
				}
				obj = cosDetails
			}
		}

		return nil
	}
}

func testAccCheckIBMSccAddonActivityInsightsCosDetailsDestroy(s *terraform.State) error {
	addonManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).AddonManagerV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_scc_addon_activity_insights_cos_details" {
			continue
		}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		addonManagerClient.AccountID = &parts[0]

		getActivityInsightsCosDetailsV2Options := &addonmanagerv1.GetActivityInsightsCosDetailsV2Options{}

		aiCosDetailsV2Output, _, err := addonManagerClient.GetActivityInsightsCosDetailsV2(getActivityInsightsCosDetailsV2Options)

		if err != nil {
			return fmt.Errorf("Error checking for scc_addon_activity_insights_cos_details (%s) has been destroyed: %s", rs.Primary.ID, err)
		}

		for _, cosDetail := range aiCosDetailsV2Output.CosDetails {
			if cosDetail.ID == &parts[1] {
				return fmt.Errorf("scc_addon_activity_insights_cos_details still exists: %s", rs.Primary.ID)
			}
		}

	}

	return nil
}
