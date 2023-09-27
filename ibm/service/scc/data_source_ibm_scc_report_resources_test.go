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

func TestAccIbmSccReportResourcesDataSourceBasic(t *testing.T) {
	instanceID, ok := os.LookupEnv("IBMCLOUD_SCC_INSTANCE_ID")
	if !ok {
		t.Logf("Missing the env var IBMCLOUD_SCC_INSTANCE_ID.")
	}

	reportID, ok := os.LookupEnv("IBMCLOUD_SCC_REPORT_ID")
	if !ok {
		t.Logf("Missing the env var IBMCLOUD_SCC_REPORT_ID.")
	}
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccReportResourcesDataSourceConfigBasic(instanceID, reportID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_report_resources.scc_report_resources_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_report_resources.scc_report_resources_instance", "report_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_report_resources.scc_report_resources_instance", "resources.#"),
				),
			},
		},
	})
}

func testAccCheckIbmSccReportResourcesDataSourceConfigBasic(instanceID, reportID string) string {
	return fmt.Sprintf(`
		data "ibm_scc_report_resources" "scc_report_resources_instance" {
			instance_id = "%s"
			report_id = "%s"
		}
	`, instanceID, reportID)
}
