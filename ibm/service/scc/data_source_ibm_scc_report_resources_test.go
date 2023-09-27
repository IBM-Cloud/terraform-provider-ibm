// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSccReportResourcesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckScc(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccReportResourcesDataSourceConfigBasic(acc.SccInstanceID, acc.SccReportID),
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
