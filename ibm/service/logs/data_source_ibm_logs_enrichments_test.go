// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmLogsEnrichmentsDataSourceBasic(t *testing.T) {
	enrichmentFieldName := "ip"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsEnrichmentsDataSourceConfigBasic(enrichmentFieldName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_enrichments.logs_enrichments_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_enrichments.logs_enrichments_instance", "enrichments.#"),
					resource.TestCheckResourceAttr("data.ibm_logs_enrichments.logs_enrichments_instance", "enrichments.0.field_name", enrichmentFieldName),
				),
			},
		},
	})
}

func testAccCheckIbmLogsEnrichmentsDataSourceConfigBasic(enrichmentFieldName string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_enrichment" "logs_enrichment_instance" {
			instance_id = "%s"
			region      = "%s"
			field_name  = "%s"
			enrichment_type {
				geo_ip  {  }
			}
		}

		data "ibm_logs_enrichments" "logs_enrichments_instance" {
			instance_id = ibm_logs_enrichment.logs_enrichment_instance.instance_id
			region      = ibm_logs_enrichment.logs_enrichment_instance.region
			depends_on  = [
				ibm_logs_enrichment.logs_enrichment_instance
			]
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, enrichmentFieldName)
}
