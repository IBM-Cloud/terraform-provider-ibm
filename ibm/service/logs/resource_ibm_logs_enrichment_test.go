// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/logs-go-sdk/logsv0"
)

func TestAccIbmLogsEnrichmentBasic(t *testing.T) {
	var conf logsv0.Enrichment
	fieldName := "ip"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsEnrichmentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsEnrichmentConfigBasic(fieldName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsEnrichmentExists("ibm_logs_enrichment.logs_enrichment_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_enrichment.logs_enrichment_instance", "field_name", fieldName),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_logs_enrichment.logs_enrichment_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmLogsEnrichmentConfigBasic(fieldName string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_enrichment" "logs_enrichment_instance" {
			instance_id = "%s"
			region      = "%s"
			field_name  = "%s"
			enrichment_type {
				geo_ip  {  }
			}
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, fieldName)
}

func testAccCheckIbmLogsEnrichmentExists(n string, obj logsv0.Enrichment) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
		if err != nil {
			return err
		}
		logsClient = getTestClientWithLogsInstanceEndpoint(logsClient)

		resourceID, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		getEnrichmentsOptions := &logsv0.GetEnrichmentsOptions{}

		enrichmentIDInt, _ := strconv.ParseInt(resourceID[2], 10, 64)

		enrichments, _, err := logsClient.GetEnrichments(getEnrichmentsOptions)
		if err != nil {
			return err
		}
		for _, enrichment := range enrichments.Enrichments {
			if *enrichment.ID == enrichmentIDInt {
				obj = enrichment
				return nil
			}
		}

		return nil
	}
}

func testAccCheckIbmLogsEnrichmentDestroy(s *terraform.State) error {
	logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
	if err != nil {
		return err
	}
	logsClient = getTestClientWithLogsInstanceEndpoint(logsClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_logs_enrichment" {
			continue
		}
		resourceID, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		enrichmentIDInt, _ := strconv.ParseInt(resourceID[2], 10, 64)
		getEnrichmentsOptions := &logsv0.GetEnrichmentsOptions{}

		// Try to find the key
		enrichments, _, _ := logsClient.GetEnrichments(getEnrichmentsOptions)

		for _, enrichment := range enrichments.Enrichments {
			if *enrichment.ID == enrichmentIDInt {
				return fmt.Errorf("logs_enrichment still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}
