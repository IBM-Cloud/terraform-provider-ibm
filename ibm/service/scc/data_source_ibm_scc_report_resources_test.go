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
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccReportResourcesDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_report_resources.scc_report_resources_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_report_resources.scc_report_resources_instance", "report_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_report_resources.scc_report_resources_instance", "resources.#"),
				),
			},
		},
	})
}

func testAccCheckIbmSccReportResourcesDataSourceConfigBasic() string {
	report_id := os.Getenv("IBMCLOUD_SCC_REPORT_ID")
	return fmt.Sprintf(`
		data "ibm_scc_report_resources" "scc_report_resources_instance" {
			report_id = "%s"
		}
	`, report_id)
}
