// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package usagereports_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMBillingSnapshotListDataSourceBasic(t *testing.T) {
	month := acc.Snapshot_month
	date_from := acc.Snapshot_date_from
	date_to := acc.Snapshot_date_to
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckUsage(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMBillingSnapshotListDataSourceConfigBasic(month, date_from, date_to),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_billing_snapshot_list.billing_snapshot_list_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_billing_snapshot_list.billing_snapshot_list_instance", "month"),
				),
			},
		},
	})
}

func testAccCheckIBMBillingSnapshotListDataSourceConfigBasic(month string, date_from string, date_to string) string {
	from, _ := strconv.ParseInt(date_from, 10, 64)
	to, _ := strconv.ParseInt(date_to, 10, 64)
	return fmt.Sprintf(`
		data "ibm_billing_snapshot_list" "billing_snapshot_list_instance" {
			month = "%s"
			date_from = "%d"
			date_to = "%d"
		}
	`, month, from, to)
}
