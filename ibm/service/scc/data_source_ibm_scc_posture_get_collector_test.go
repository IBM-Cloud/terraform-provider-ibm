// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSccPostureGetCollectorDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccPostureGetCollectorDataSourceConfigBasic(acc.Scc_posture_collector_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_get_collector.get_collector", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_get_collector.get_collector", "display_name"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_get_collector.get_collector", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_get_collector.get_collector", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_get_collector.get_collector", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_get_collector.get_collector", "enabled"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_get_collector.get_collector", "registration_code"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_get_collector.get_collector", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_get_collector.get_collector", "failure_count"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_get_collector.get_collector", "use_private_endpoint"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_get_collector.get_collector", "managed_by"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_get_collector.get_collector", "status_description"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_get_collector.get_collector", "is_public"),
				),
			},
		},
	})
}

func testAccCheckIBMSccPostureGetCollectorDataSourceConfigBasic(collectorId string) string {
	return fmt.Sprintf(`
		data "ibm_scc_posture_get_collector" "get_collector" {
			id = "%s"
		}
	`, collectorId)
}
