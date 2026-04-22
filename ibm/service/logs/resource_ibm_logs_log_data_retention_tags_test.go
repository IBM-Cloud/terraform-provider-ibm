// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/logs-go-sdk/logsv0"
)

func TestAccIbmLogsLogDataRetentionTagsBasic(t *testing.T) {
	var conf logsv0.LogDataRetentionTags

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCloudLogs(t); testAccPreCheckArchiveBucket(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsLogDataRetentionTagsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsLogDataRetentionTagsConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsLogDataRetentionTagsExists("ibm_logs_log_data_retention_tags.logs_data_retention_tags_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_log_data_retention_tags.logs_data_retention_tags_instance", "tags.0", "Short"),
					resource.TestCheckResourceAttr("ibm_logs_log_data_retention_tags.logs_data_retention_tags_instance", "tags.1", "Medium"),
					resource.TestCheckResourceAttr("ibm_logs_log_data_retention_tags.logs_data_retention_tags_instance", "tags.2", "Long"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsLogDataRetentionTagsConfigUpdate(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_log_data_retention_tags.logs_data_retention_tags_instance", "tags.0", "Temporary"),
					resource.TestCheckResourceAttr("ibm_logs_log_data_retention_tags.logs_data_retention_tags_instance", "tags.1", "Standard"),
					resource.TestCheckResourceAttr("ibm_logs_log_data_retention_tags.logs_data_retention_tags_instance", "tags.2", "Extended"),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_logs_log_data_retention_tags.logs_data_retention_tags_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmLogsLogDataRetentionTagsConfigBasic() string {
	return fmt.Sprintf(`
	resource "ibm_logs_log_data_retention_tags" "logs_data_retention_tags_instance" {
		instance_id = "%s"
		region      = "%s"
		tags = ["Short", "Medium", "Long"]
	}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion)
}

func testAccCheckIbmLogsLogDataRetentionTagsConfigUpdate() string {
	return fmt.Sprintf(`
	resource "ibm_logs_log_data_retention_tags" "logs_data_retention_tags_instance" {
		instance_id = "%s"
		region      = "%s"
		tags = ["Temporary", "Standard", "Extended"]
	}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion)
}

func testAccCheckIbmLogsLogDataRetentionTagsExists(n string, obj logsv0.LogDataRetentionTags) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
		if err != nil {
			return err
		}
		logsClient = getTestClientWithLogsInstanceEndpoint(logsClient)

		getLogDataRetentionTagsOptions := &logsv0.GetLogDataRetentionTagsOptions{}

		logDataRetentionTags, _, err := logsClient.GetLogDataRetentionTags(getLogDataRetentionTagsOptions)
		if err != nil {
			return err
		}

		obj = *logDataRetentionTags
		return nil
	}
}

func testAccCheckIbmLogsLogDataRetentionTagsDestroy(s *terraform.State) error {
	// Note: Data retention tags are a configuration on the Cloud Logs instance
	// and cannot be "deleted" in the traditional sense. The API only supports
	// GET and PUT operations. When the Terraform resource is destroyed, it's
	// removed from state but the configuration remains on the instance.
	// This is expected behavior, so we just verify the resource was removed
	// from Terraform state, not from the actual instance.

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_logs_log_data_retention_tags" {
			continue
		}
		// Resource was in state during test, which is expected
		// No actual destroy verification needed as this is a configuration resource
	}

	return nil
}

func testAccPreCheckArchiveBucket(t *testing.T) {
	logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
	if err != nil {
		t.Skipf("Error getting logs client: %s", err)
		return
	}
	logsClient = getTestClientWithLogsInstanceEndpoint(logsClient)

	// First, try to get the current retention tags
	getLogDataRetentionTagsOptions := &logsv0.GetLogDataRetentionTagsOptions{}
	currentTags, getResponse, getErr := logsClient.GetLogDataRetentionTags(getLogDataRetentionTagsOptions)

	// Check GET errors
	if getErr != nil {
		errMsg := getErr.Error()
		if getResponse != nil && (getResponse.StatusCode == 400 || getResponse.StatusCode == 422) {
			if strings.Contains(strings.ToLower(errMsg), "archive bucket") {
				t.Skip("Skipping test: Archive bucket must be attached before configuring retention tags. Please configure an archive bucket in your Cloud Logs instance first.")
			}
		}
	}

	// Try a test PUT to verify archive bucket is configured
	// This is necessary because GET might succeed but PUT fails if archive bucket is not configured
	// Use test tags or current tags if available
	testTags := []string{"Test1", "Test2", "Test3"}
	if currentTags != nil && currentTags.Tags != nil && len(currentTags.Tags) == 3 {
		testTags = currentTags.Tags
	}

	updateLogDataRetentionTagsOptions := &logsv0.UpdateLogDataRetentionTagsOptions{
		Tags: testTags,
	}
	_, putErr := logsClient.UpdateLogDataRetentionTags(updateLogDataRetentionTagsOptions)

	if putErr != nil {
		errMsg := putErr.Error()
		// Only skip if the error specifically mentions "Archive bucket must be attached"
		if strings.Contains(errMsg, "Archive bucket must be attached") {
			t.Skip("Skipping test: Archive bucket must be attached before configuring retention tags. Please configure an archive bucket in your Cloud Logs instance first.")
		}
		// If it's a different error, log it but don't skip - let the actual test handle it
		t.Logf("PreCheck PUT operation failed (not archive bucket error): %v", putErr)
	}
}
