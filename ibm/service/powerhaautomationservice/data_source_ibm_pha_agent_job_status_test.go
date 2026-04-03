// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.113.1-d76630af-20260320-135953
*/

package powerhaautomationservice_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMPhaAgentJobStatusDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMPhaAgentJobStatusDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pha_agent_job_status.pha_agent_job_status_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_agent_job_status.pha_agent_job_status_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_pha_agent_job_status.pha_agent_job_status_instance", "job_id"),
				),
			},
		},
	})
}

func testAccCheckIBMPhaAgentJobStatusDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_pha_agent_job_status" "pha_agent_job_status_instance" {
			instance_id = "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
			job_id = "4235r23r5vdfdf-2323"
			Accept-Language = "en-US"
			If-None-Match = "abcdef"
		}
	`)
}

