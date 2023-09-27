// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSccReportControlsDataSourceBasic(t *testing.T) {
	instanceID, ok := os.LookupEnv("IBMCLOUD_SCC_INSTANCE_ID")
	if !ok {
		t.Logf("Missing the env var IBMCLOUD_SCC_INSTANCE_ID.")
	}

	reportID, ok := os.LookupEnv("IBMCLOUD_SCC_REPORT_ID")
	if !ok {
		t.Logf("Missing the env var IBMCLOUD_SCC_REPORT_ID.")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckSccInstanceID(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccReportControlsDataSourceConfigBasic(instanceID, reportID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_report_controls.scc_report_controls_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_report_controls.scc_report_controls_instance", "report_id"),
				),
			},
		},
	})
}

func testAccCheckIbmSccReportControlsDataSourceConfigBasic(instanceID, reportID string) string {
	return fmt.Sprintf(`
		data "ibm_scc_report_controls" "scc_report_controls_instance" {
			instance_id = "%s"
			report_id = "%s"
		}
	`, instanceID, reportID)
}
