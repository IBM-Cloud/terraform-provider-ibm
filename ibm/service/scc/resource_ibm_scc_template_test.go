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
	"github.com/IBM/scc-go-sdk/v3/configurationgovernancev1"
)

func TestAccIBMSccTemplateBasic(t *testing.T) {
	var conf configurationgovernancev1.TemplateResponse

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSccTemplateDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccTemplateConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSccTemplateExists("ibm_scc_template.scc_template", conf),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_scc_template.scc_template",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMSccTemplateConfigBasic() string {
	account_id := os.Getenv("SCC_GOVERNANCE_ACCOUNT_ID")
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
	`, account_id)
}

func testAccCheckIBMSccTemplateExists(n string, obj configurationgovernancev1.TemplateResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		configurationGovernanceClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ConfigurationGovernanceV1()
		if err != nil {
			return err
		}

		getTemplateOptions := &configurationgovernancev1.GetTemplateOptions{}

		getTemplateOptions.SetTemplateID(rs.Primary.ID)

		template, _, err := configurationGovernanceClient.GetTemplate(getTemplateOptions)
		if err != nil {
			return err
		}

		obj = *template
		return nil
	}
}

func testAccCheckIBMSccTemplateDestroy(s *terraform.State) error {
	configurationGovernanceClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ConfigurationGovernanceV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_scc_template" {
			continue
		}

		getTemplateOptions := &configurationgovernancev1.GetTemplateOptions{}

		getTemplateOptions.SetTemplateID(rs.Primary.ID)

		// Try to find the key
		_, response, err := configurationGovernanceClient.GetTemplate(getTemplateOptions)

		if err == nil {
			return fmt.Errorf("scc_template still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for scc_template (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
