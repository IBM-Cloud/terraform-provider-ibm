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

func TestAccIBMSccAddonNetworkInsightsCosDetailsBasic(t *testing.T) {
	var conf addonmanagerv1.NiCosDetailsV2Output
	regionID := "us"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSccAddonNetworkInsightsCosDetailsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccAddonNetworkInsightsCosDetailsConfigBasic(regionID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSccAddonNetworkInsightsCosDetailsExists("ibm_scc_addon_network_insights_cos_details.scc_addon_network_insights_cos_details", conf),
					resource.TestCheckResourceAttrSet("ibm_scc_addon_network_insights_cos_details.scc_addon_network_insights_cos_details", "cos_details.#"),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_scc_addon_network_insights_cos_details.scc_addon_network_insights_cos_details",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMSccAddonNetworkInsightsCosDetailsConfigBasic(regionID string) string {
	return fmt.Sprintf(`

		resource "ibm_scc_addon_network_insights_cos_details" "scc_addon_network_insights_cos_details" {
			region_id = "%s"
			cos_details {
				cos_instance = "cos_instance"
				bucket_name = "bucket_name"
				description = "description"
				type = "network-insights"
				cos_bucket_url = "cos_bucket_url"
			}
		}
	`, regionID)
}

func testAccCheckIBMSccAddonNetworkInsightsCosDetailsExists(n string, obj addonmanagerv1.NiCosDetailsV2Output) resource.TestCheckFunc {

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

		getNetworkInsightsCosDetailsV2Options := &addonmanagerv1.GetNetworkInsightsCosDetailsV2Options{}

		niCosDetailsV2Output, _, err := addonManagerClient.GetNetworkInsightsCosDetailsV2(getNetworkInsightsCosDetailsV2Options)
		if err != nil {
			return err
		}

		for _, cosDetail := range niCosDetailsV2Output.CosDetails {
			if cosDetail.ID == &parts[1] {
				cosDetails := addonmanagerv1.NiCosDetailsV2Output{
					RegionID: niCosDetailsV2Output.RegionID,
					CosDetails: []addonmanagerv1.NiCosDetailsV2OutputCosDetailsItem{
						cosDetail,
					},
				}
				obj = cosDetails
			}
		}

		return nil
	}
}

func testAccCheckIBMSccAddonNetworkInsightsCosDetailsDestroy(s *terraform.State) error {
	addonManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).AddonManagerV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_scc_addon_network_insights_cos_details" {
			continue
		}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		addonManagerClient.AccountID = &parts[0]

		getNetworkInsightsCosDetailsV2Options := &addonmanagerv1.GetNetworkInsightsCosDetailsV2Options{}

		niCosDetailsV2Output, _, err := addonManagerClient.GetNetworkInsightsCosDetailsV2(getNetworkInsightsCosDetailsV2Options)

		if err != nil {
			return fmt.Errorf("Error checking for scc_addon_activity_insights_cos_details (%s) has been destroyed: %s", rs.Primary.ID, err)
		}

		for _, cosDetail := range niCosDetailsV2Output.CosDetails {
			if cosDetail.ID == &parts[1] {
				return fmt.Errorf("scc_addon_network_insights_cos_details still exists: %s", rs.Primary.ID)
			}
		}

	}

	return nil
}
