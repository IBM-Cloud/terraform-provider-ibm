// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package usagereports_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/platform-services-go-sdk/usagereportsv4"
)

func TestAccIBMBillingReportSnapshotBasic(t *testing.T) {
	var conf usagereportsv4.SnapshotConfig
	interval := "daily"
	cosBucket := acc.Cos_bucket
	cosLocation := acc.Cos_location
	intervalUpdate := "daily"
	cosBucketUpdate := acc.Cos_bucket_update
	cosLocationUpdate := acc.Cos_location_update

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckUsage(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMBillingReportSnapshotDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMBillingReportSnapshotConfigBasic(interval, cosBucket, cosLocation),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMBillingReportSnapshotExists("ibm_billing_report_snapshot.billing_report_snapshot_instance", conf),
					resource.TestCheckResourceAttr("ibm_billing_report_snapshot.billing_report_snapshot_instance", "interval", interval),
					resource.TestCheckResourceAttr("ibm_billing_report_snapshot.billing_report_snapshot_instance", "cos_bucket", cosBucket),
					resource.TestCheckResourceAttr("ibm_billing_report_snapshot.billing_report_snapshot_instance", "cos_location", cosLocation),
				),
			},
			{
				Config: testAccCheckIBMBillingReportSnapshotConfigBasic(intervalUpdate, cosBucketUpdate, cosLocationUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_billing_report_snapshot.billing_report_snapshot_instance", "interval", intervalUpdate),
					resource.TestCheckResourceAttr("ibm_billing_report_snapshot.billing_report_snapshot_instance", "cos_bucket", cosBucketUpdate),
					resource.TestCheckResourceAttr("ibm_billing_report_snapshot.billing_report_snapshot_instance", "cos_location", cosLocationUpdate),
				),
			},
		},
	})
}

func TestAccIBMBillingReportSnapshotAllArgs(t *testing.T) {
	var conf usagereportsv4.SnapshotConfig
	interval := "daily"
	versioning := "new"
	cosReportsFolder := fmt.Sprintf("tf_cos_reports_folder_%d", acctest.RandIntRange(10, 100))
	cosBucket := acc.Cos_bucket
	cosLocation := acc.Cos_location

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckUsage(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMBillingReportSnapshotDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMBillingReportSnapshotConfig(interval, versioning, cosReportsFolder, cosBucket, cosLocation),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMBillingReportSnapshotExists("ibm_billing_report_snapshot.billing_report_snapshot_instance", conf),
					resource.TestCheckResourceAttr("ibm_billing_report_snapshot.billing_report_snapshot_instance", "interval", interval),
					resource.TestCheckResourceAttr("ibm_billing_report_snapshot.billing_report_snapshot_instance", "versioning", versioning),
					resource.TestCheckResourceAttr("ibm_billing_report_snapshot.billing_report_snapshot_instance", "cos_reports_folder", cosReportsFolder),
					resource.TestCheckResourceAttr("ibm_billing_report_snapshot.billing_report_snapshot_instance", "cos_bucket", cosBucket),
					resource.TestCheckResourceAttr("ibm_billing_report_snapshot.billing_report_snapshot_instance", "cos_location", cosLocation),
				),
			},
		},
	})
}

func testAccCheckIBMBillingReportSnapshotConfigBasic(interval string, cosBucket string, cosLocation string) string {
	return fmt.Sprintf(`
		resource "ibm_billing_report_snapshot" "billing_report_snapshot_instance" {
			interval = "%s"
			cos_bucket = "%s"
			cos_location = "%s"
			report_types = ["account_summary", "account_resource_instance_usage"]
		}
	`, interval, cosBucket, cosLocation)
}

func testAccCheckIBMBillingReportSnapshotConfig(interval string, versioning string, cosReportsFolder string, cosBucket string, cosLocation string) string {
	return fmt.Sprintf(`

		resource "ibm_billing_report_snapshot" "billing_report_snapshot_instance" {
			interval = "%s"
			versioning = "%s"
			report_types = ["account_summary", "account_resource_instance_usage"]
			cos_reports_folder = "%s"
			cos_bucket = "%s"
			cos_location = "%s"
		}
	`, interval, versioning, cosReportsFolder, cosBucket, cosLocation)
}

func testAccCheckIBMBillingReportSnapshotExists(n string, obj usagereportsv4.SnapshotConfig) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		usageReportsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).UsageReportsV4()
		if err != nil {
			return err
		}

		getReportsSnapshotConfigOptions := &usagereportsv4.GetReportsSnapshotConfigOptions{}

		getReportsSnapshotConfigOptions.SetAccountID(rs.Primary.ID)

		snapshotConfig, _, err := usageReportsClient.GetReportsSnapshotConfig(getReportsSnapshotConfigOptions)
		if err != nil {
			return err
		}

		obj = *snapshotConfig
		return nil
	}
}

func testAccCheckIBMBillingReportSnapshotDestroy(s *terraform.State) error {
	usageReportsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).UsageReportsV4()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_billing_report_snapshot" {
			continue
		}

		getReportsSnapshotConfigOptions := &usagereportsv4.GetReportsSnapshotConfigOptions{}

		getReportsSnapshotConfigOptions.SetAccountID(rs.Primary.ID)

		// Try to find the key
		res, response, err := usageReportsClient.GetReportsSnapshotConfig(getReportsSnapshotConfigOptions)

		if !(response.StatusCode == 200 && *res.State == "disabled") {
			return fmt.Errorf("billing_report_snapshot still exists: %s", rs.Primary.ID)
		} else if err != nil {
			return fmt.Errorf("Error checking for billing_report_snapshot (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
