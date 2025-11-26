// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package codeengine_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/code-engine-go-sdk/codeenginev2"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIbmCodeEngineAllowedOutboundDestinationBasic(t *testing.T) {
	var conf codeenginev2.AllowedOutboundDestination

	projectID := acc.CeProjectId
	typeVar := "cidr_block"
	cidrBlock := "192.68.3.0/24"
	name := "test-cidr"
	cidrBlockUpdate := "192.68.2.0/24"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCodeEngine(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmCodeEngineAllowedOutboundDestinationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmCodeEngineAllowedOutboundDestinationConfigBasic(projectID, typeVar, cidrBlock, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmCodeEngineAllowedOutboundDestinationExists("ibm_code_engine_allowed_outbound_destination.code_engine_allowed_outbound_destination_instance", conf),
					resource.TestCheckResourceAttr("ibm_code_engine_allowed_outbound_destination.code_engine_allowed_outbound_destination_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_allowed_outbound_destination.code_engine_allowed_outbound_destination_instance", "type", typeVar),
					resource.TestCheckResourceAttr("ibm_code_engine_allowed_outbound_destination.code_engine_allowed_outbound_destination_instance", "cidr_block", cidrBlock),
					resource.TestCheckResourceAttr("ibm_code_engine_allowed_outbound_destination.code_engine_allowed_outbound_destination_instance", "name", name),
				),
			},
			{
				Config: testAccCheckIbmCodeEngineAllowedOutboundDestinationConfigBasic(projectID, typeVar, cidrBlockUpdate, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_code_engine_allowed_outbound_destination.code_engine_allowed_outbound_destination_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_allowed_outbound_destination.code_engine_allowed_outbound_destination_instance", "type", typeVar),
					resource.TestCheckResourceAttr("ibm_code_engine_allowed_outbound_destination.code_engine_allowed_outbound_destination_instance", "cidr_block", cidrBlockUpdate),
					resource.TestCheckResourceAttr("ibm_code_engine_allowed_outbound_destination.code_engine_allowed_outbound_destination_instance", "name", name),
				),
			},
		},
	})
}

func testAccCheckIbmCodeEngineAllowedOutboundDestinationConfigBasic(projectID string, typeVar string, cidrBlock string, name string) string {
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
	`, projectID, typeVar, cidrBlock, name)
}

func testAccCheckIbmCodeEngineAllowedOutboundDestinationExists(n string, obj codeenginev2.AllowedOutboundDestination) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		codeEngineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CodeEngineV2()
		if err != nil {
			return err
		}

		getAllowedOutboundDestinationOptions := &codeenginev2.GetAllowedOutboundDestinationOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getAllowedOutboundDestinationOptions.SetProjectID(parts[0])
		getAllowedOutboundDestinationOptions.SetName(parts[1])

		allowedOutboundDestinationIntf, _, err := codeEngineClient.GetAllowedOutboundDestination(getAllowedOutboundDestinationOptions)
		if err != nil {
			return err
		}

		allowedOutboundDestination := allowedOutboundDestinationIntf.(*codeenginev2.AllowedOutboundDestination)
		obj = *allowedOutboundDestination
		return nil
	}
}

func testAccCheckIbmCodeEngineAllowedOutboundDestinationDestroy(s *terraform.State) error {
	codeEngineClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).CodeEngineV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_code_engine_allowed_outbound_destination" {
			continue
		}

		getAllowedOutboundDestinationOptions := &codeenginev2.GetAllowedOutboundDestinationOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getAllowedOutboundDestinationOptions.SetProjectID(parts[0])
		getAllowedOutboundDestinationOptions.SetName(parts[1])

		// Try to find the key
		_, response, err := codeEngineClient.GetAllowedOutboundDestination(getAllowedOutboundDestinationOptions)

		if err == nil {
			return fmt.Errorf("code_engine_allowed_outbound_destination still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for code_engine_allowed_outbound_destination (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
