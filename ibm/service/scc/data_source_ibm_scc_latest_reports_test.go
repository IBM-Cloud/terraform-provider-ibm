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

func TestAccIbmSccLatestReportsDataSourceBasic(t *testing.T) {
	instanceID, ok := os.LookupEnv("IBMCLOUD_SCC_INSTANCE_ID")
	if !ok {
		t.Logf("Missing the env var IBMCLOUD_SCC_INSTANCE_ID.")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckSccInstanceID(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccLatestReportsDataSourceConfigBasic(instanceID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_latest_reports.scc_latest_reports_instance", "id"),
				),
			},
		},
	})
}

func testAccCheckIbmSccLatestReportsDataSourceConfigBasic(instanceID string) string {
	return fmt.Sprintf(`
		data "ibm_scc_latest_reports" "scc_latest_reports_instance" {
			instance_id = "%s"
			sort = "profile_name"
		}
	`, instanceID)
}
