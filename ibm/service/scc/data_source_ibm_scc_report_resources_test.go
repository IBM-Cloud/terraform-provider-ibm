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
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccReportResourcesDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_report_resources.scc_report_resources", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_report_resources.scc_report_resources", "report_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_report_resources.scc_report_resources", "first.#"),
				),
			},
		},
	})
}

func testAccCheckIbmSccReportResourcesDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_scc_report_resources" "scc_report_resources_instance" {
			report_id = "report_id"
			id = "id"
			resource_name = "resource_name"
			account_id = "account_id"
			component_id = "component_id"
			status = "compliant"
			sort = "account_id"
		}
	`)
}
