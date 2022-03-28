// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/scc-go-sdk/v3/configurationgovernancev1"
)

func TestAccIBMSccTemplateAttachmentBasic(t *testing.T) {
	var conf configurationgovernancev1.TemplateAttachment

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSccTemplateAttachmentDestroy,

		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccTemplateAttachmentConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSccTemplateAttachmentExists("ibm_scc_template_attachment.scc_template_attachment", conf),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_scc_template_attachment.scc_template_attachment",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMSccTemplateAttachmentConfigBasic() string {
	account_id := os.Getenv("SCC_GOVERNANCE_ACCOUNT_ID")
	resource_group_id := os.Getenv("IBM_SCC_RESOURCE_GROUP")
	return fmt.Sprintf(`
	resource "ibm_scc_template" "scc_template" {
		account_id = "%s"
		name = "Terraform template"
		description = "description"
		target {
		  service_name = "cloud-object-storage"
		  resource_kind = "bucket"
		  additional_target_attributes {
			name = "location"
			value = "us-south"
		  }
		}
		customized_defaults {
		  property = "activity_tracking.write_data_events"
		  value = "true"
		}
	  }

	  resource "ibm_scc_template_attachment" "scc_template_attachment" {
		template_id = ibm_scc_template.scc_template.id
		account_id = "%s"
		included_scope {
			note = "note"
			scope_id = "%s"
			scope_type = "account"
		}
		excluded_scopes {
			note = "note"
			scope_id = "%s"
			scope_type = "account.resource_group"
		}
		depends_on = [
			ibm_scc_template.scc_template
		]
	}
	`, account_id, account_id, account_id, resource_group_id)
}

func testAccCheckIBMSccTemplateAttachmentExists(n string, obj configurationgovernancev1.TemplateAttachment) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		configurationGovernanceClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ConfigurationGovernanceV1()
		if err != nil {
			return err
		}

		getTemplateAttachmentOptions := &configurationgovernancev1.GetTemplateAttachmentOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		templateID := parts[0]
		getTemplateAttachmentOptions.SetTemplateID(templateID)
		getTemplateAttachmentOptions.SetAttachmentID(parts[1])

		templateAttachment, _, err := configurationGovernanceClient.GetTemplateAttachment(getTemplateAttachmentOptions)
		if err != nil {
			return err
		}

		if *templateAttachment.TemplateID != templateID {
			return fmt.Errorf(
				"ibm_scc_template_attachment.scc_template_attachment: Attribute 'template_id' expected %#v, got %#v",
				templateID,
				templateAttachment.TemplateID,
			)
		}
		obj = *templateAttachment
		return nil
	}
}

func testAccCheckIBMSccTemplateAttachmentDestroy(s *terraform.State) error {
	configurationGovernanceClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ConfigurationGovernanceV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_scc_template_attachment" {
			continue
		}

		getTemplateAttachmentOptions := &configurationgovernancev1.GetTemplateAttachmentOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getTemplateAttachmentOptions.SetTemplateID(parts[0])
		getTemplateAttachmentOptions.SetAttachmentID(parts[1])

		// Try to find the key
		_, response, err := configurationGovernanceClient.GetTemplateAttachment(getTemplateAttachmentOptions)

		if err == nil {
			return fmt.Errorf("scc_template_attachment still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for scc_template_attachment (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
