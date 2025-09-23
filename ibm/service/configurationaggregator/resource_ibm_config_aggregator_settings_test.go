// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package configurationaggregator_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmConfigAggregatorSettingsBasic(t *testing.T) {

	instanceID := "instance_id"
	resourceCollectionEnabled := false
	trustedProfileID := "Profile-2546925a-7b46-40dd-81ff-48015a49ff43"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmConfigAggregatorSettingsConfigBasic(instanceID, resourceCollectionEnabled, trustedProfileID),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_config_aggregator_settings.config_aggregator_settings_instance", "resource_collection_enabled", fmt.Sprintf("%t", resourceCollectionEnabled)),
					resource.TestCheckResourceAttr("ibm_config_aggregator_settings.config_aggregator_settings_instance", "trusted_profile_id", trustedProfileID),
					resource.TestCheckResourceAttr("ibm_config_aggregator_settings.config_aggregator_settings_instance", "instance_id", instanceID),
				),
			},
		},
	})
}

func testAccCheckIbmConfigAggregatorSettingsConfigBasic(instanceID string, resourceCollectionEnabled bool, trustedProfileID string) string {
	return fmt.Sprintf(`
        resource "ibm_config_aggregator_settings" "config_aggregator_settings_instance" {
            instance_id = "%s"
            resource_collection_enabled = %t
            trusted_profile_id = "%s"
            resource_collection_regions = ["all"]
        }
    `, instanceID, resourceCollectionEnabled, trustedProfileID)
}
