// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventnotification_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMEnMetricsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMEnMetricsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_en_metrics.en_metrics_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_metrics.en_metrics_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_metrics.en_metrics_instance", "destination_type"),
					resource.TestCheckResourceAttrSet("data.ibm_en_metrics.en_metrics_instance", "gte"),
					resource.TestCheckResourceAttrSet("data.ibm_en_metrics.en_metrics_instance", "lte"),
					resource.TestCheckResourceAttrSet("data.ibm_en_metrics.en_metrics_instance", "metrics.#"),
				),
			},
		},
	})
}

func testAccCheckIBMEnMetricsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_en_metrics" "en_metrics_instance" {
			instance_id = "instance_id"
			destination_type = "smtp_custom"
			gte = "gte"
			lte = "lte"
			destination_id = "destination_id"
			source_id = "source_id"
			email_to = "email_to"
			notification_id = "notification_id"
			subject = "subject"
		}
	`)
}
