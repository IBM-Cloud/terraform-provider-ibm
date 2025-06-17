// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package codeengine_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmCodeEngineAllowedOutboundDestinationDataSourceBasic(t *testing.T) {
	projectID := acc.CeProjectId
	typeVar := "cidr_block"
	cidrBlock := "192.68.3.0/24"
	name := "test-cidr"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCodeEngine(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmCodeEngineAllowedOutboundDestinationDataSourceConfigBasic(projectID, typeVar, cidrBlock, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_code_engine_allowed_outbound_destination.code_engine_allowed_outbound_destination_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("data.ibm_code_engine_allowed_outbound_destination.code_engine_allowed_outbound_destination_instance", "name", name),
					resource.TestCheckResourceAttr("data.ibm_code_engine_allowed_outbound_destination.code_engine_allowed_outbound_destination_instance", "type", typeVar),
					resource.TestCheckResourceAttr("data.ibm_code_engine_allowed_outbound_destination.code_engine_allowed_outbound_destination_instance", "cidr_block", cidrBlock),
				),
			},
		},
	})
}

func testAccCheckIbmCodeEngineAllowedOutboundDestinationDataSourceConfigBasic(projectID string, typeVar string, cidrBlock string, name string) string {
	return fmt.Sprintf(`
		data "ibm_code_engine_project" "code_engine_project_instance" {
			project_id = "%s"
		}

		resource "ibm_code_engine_allowed_outbound_destination" "code_engine_allowed_outbound_destination_instance" {
			project_id = data.ibm_code_engine_project.code_engine_project_instance.project_id
			type = "%s"
			cidr_block = "%s"
			name = "%s"
		}

		data "ibm_code_engine_allowed_outbound_destination" "code_engine_allowed_outbound_destination_instance" {
			project_id = ibm_code_engine_allowed_outbound_destination.code_engine_allowed_outbound_destination_instance.project_id
			name = ibm_code_engine_allowed_outbound_destination.code_engine_allowed_outbound_destination_instance.name
		}
	`, projectID, typeVar, cidrBlock, name)
}
