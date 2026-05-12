// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package codeengine_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/code-engine-go-sdk/codeenginev2"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIbmCodeEngineAllowedOutboundDestinationCIDRBlock(t *testing.T) {
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
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineAllowedOutboundDestinationCIDRBlock(projectID, typeVar, cidrBlock, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmCodeEngineAllowedOutboundDestinationExists("ibm_code_engine_allowed_outbound_destination.code_engine_allowed_outbound_destination_instance", conf),
					resource.TestCheckResourceAttr("ibm_code_engine_allowed_outbound_destination.code_engine_allowed_outbound_destination_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_allowed_outbound_destination.code_engine_allowed_outbound_destination_instance", "type", typeVar),
					resource.TestCheckResourceAttr("ibm_code_engine_allowed_outbound_destination.code_engine_allowed_outbound_destination_instance", "cidr_block", cidrBlock),
					resource.TestCheckResourceAttr("ibm_code_engine_allowed_outbound_destination.code_engine_allowed_outbound_destination_instance", "name", name),
				),
			},
			{
				Config: testAccCheckIbmCodeEngineAllowedOutboundDestinationCIDRBlock(projectID, typeVar, cidrBlockUpdate, name),
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

func TestAccIbmCodeEngineAllowedOutboundDestinationPrivatePathServiceGateway(t *testing.T) {
	projectID := acc.CeProjectId
	name := "aod-name"
	aodType := "private_path_service_gateway"
	privatePathServiceGatewayCrn := acctest.CePrivatePathServiceGatewayCrn

	ppsgCreate := codeenginev2.AllowedOutboundDestinationPrototype{
		Name:                         &name,
		Type:                         &aodType,
		PrivatePathServiceGatewayCrn: &privatePathServiceGatewayCrn,
		IsolationPolicy:              core.StringPtr("dedicated"),
	}

	ppsgUpdate := codeenginev2.AllowedOutboundDestinationPrototype{
		Name:                         &name,
		Type:                         &aodType,
		PrivatePathServiceGatewayCrn: &privatePathServiceGatewayCrn,
		IsolationPolicy:              core.StringPtr("shared"),
	}

	var conf codeenginev2.AllowedOutboundDestination
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCodeEngine(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmCodeEngineAllowedOutboundDestinationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmCodeEngineAllowedOutboundDestinationPrivatePathServiceGateway(projectID, ppsgCreate),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmCodeEngineAllowedOutboundDestinationExists("ibm_code_engine_allowed_outbound_destination.code_engine_allowed_outbound_destination_instance", conf),
					resource.TestCheckResourceAttr("ibm_code_engine_allowed_outbound_destination.code_engine_allowed_outbound_destination_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_allowed_outbound_destination.code_engine_allowed_outbound_destination_instance", "private_path_service_gateway_crn", privatePathServiceGatewayCrn),
					resource.TestCheckResourceAttr("ibm_code_engine_allowed_outbound_destination.code_engine_allowed_outbound_destination_instance", "isolation_policy", "dedicated"),
				),
			},
			{
				Config: testAccCheckIbmCodeEngineAllowedOutboundDestinationPrivatePathServiceGateway(projectID, ppsgUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmCodeEngineAllowedOutboundDestinationExists("ibm_code_engine_allowed_outbound_destination.code_engine_allowed_outbound_destination_instance", conf),
					resource.TestCheckResourceAttr("ibm_code_engine_allowed_outbound_destination.code_engine_allowed_outbound_destination_instance", "project_id", projectID),
					resource.TestCheckResourceAttr("ibm_code_engine_allowed_outbound_destination.code_engine_allowed_outbound_destination_instance", "private_path_service_gateway_crn", privatePathServiceGatewayCrn),
					resource.TestCheckResourceAttr("ibm_code_engine_allowed_outbound_destination.code_engine_allowed_outbound_destination_instance", "isolation_policy", "shared"),
				),
			},
		},
	})
}

func testAccCheckIbmCodeEngineAllowedOutboundDestinationCIDRBlock(projectID string, typeVar string, cidrBlock string, name string) string {
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

func testAccCheckIbmCodeEngineAllowedOutboundDestinationPrivatePathServiceGateway(projectID string, allowedOutboundDestination codeenginev2.AllowedOutboundDestinationPrototype) string {
	return fmt.Sprintf(`
		data "ibm_code_engine_project" "code_engine_project_instance" {
			project_id = "%s"
		}

		resource "ibm_code_engine_allowed_outbound_destination" "code_engine_allowed_outbound_destination_instance" {
			name = "%s"
			project_id = data.ibm_code_engine_project.code_engine_project_instance.project_id
			type = "%s"
			private_path_service_gateway_crn = "%s"
			isolation_policy = "%s"
		}
	`, projectID, *allowedOutboundDestination.Name, *allowedOutboundDestination.Type, *allowedOutboundDestination.PrivatePathServiceGatewayCrn, *allowedOutboundDestination.IsolationPolicy)
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

		lookUp := func() *retry.RetryError {
			_, response, err := codeEngineClient.GetAllowedOutboundDestination(getAllowedOutboundDestinationOptions)

			if err != nil && response.StatusCode == 404 { // Resource successfully deleted
				return nil
			}

			if err != nil && response.StatusCode != 404 { // Unexpected error
				return retry.NonRetryableError(fmt.Errorf("Error checking for code_engine_allowed_outbound_destination (%s) has been destroyed: %s", rs.Primary.ID, err))
			}

			return retry.RetryableError(fmt.Errorf("code_engine_allowed_outbound_destination still exists: %s", rs.Primary.ID))
		}

		// Poll for up to 30 seconds to verify the resource is deleted (asynchronous deletion)
		if err := retry.RetryContext(context.Background(), 30*time.Second, lookUp); err != nil {
			return err
		}
	}

	return nil
}
