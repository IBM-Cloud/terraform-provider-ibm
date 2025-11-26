// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventnotification_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/event-notifications-go-admin-sdk/eventnotificationsv1"
)

func TestAccIBMEnSMTPConfigurationBasic(t *testing.T) {
	var conf eventnotificationsv1.SMTPConfiguration
	instanceID := fmt.Sprintf("tf_instance_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	domain := fmt.Sprintf("tf_domain_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	domainUpdate := fmt.Sprintf("tf_domain_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMEnSMTPConfigurationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMEnSMTPConfigurationConfigBasic(instanceID, name, domain),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEnSMTPConfigurationExists("ibm_en_smtp_configuration.en_smtp_configuration_instance", conf),
					resource.TestCheckResourceAttr("ibm_en_smtp_configuration.en_smtp_configuration_instance", "instance_id", instanceID),
					resource.TestCheckResourceAttr("ibm_en_smtp_configuration.en_smtp_configuration_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_en_smtp_configuration.en_smtp_configuration_instance", "domain", domain),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMEnSMTPConfigurationConfigBasic(instanceID, nameUpdate, domainUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_en_smtp_configuration.en_smtp_configuration_instance", "instance_id", instanceID),
					resource.TestCheckResourceAttr("ibm_en_smtp_configuration.en_smtp_configuration_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_en_smtp_configuration.en_smtp_configuration_instance", "domain", domainUpdate),
				),
			},
		},
	})
}

func TestAccIBMEnSMTPConfigurationAllArgs(t *testing.T) {
	var conf eventnotificationsv1.SMTPConfiguration
	instanceID := fmt.Sprintf("tf_instance_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	domain := fmt.Sprintf("tf_domain_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	domainUpdate := fmt.Sprintf("tf_domain_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMEnSMTPConfigurationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMEnSMTPConfigurationConfig(instanceID, name, description, domain),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEnSMTPConfigurationExists("ibm_en_smtp_configuration.en_smtp_configuration_instance", conf),
					resource.TestCheckResourceAttr("ibm_en_smtp_configuration.en_smtp_configuration_instance", "instance_id", instanceID),
					resource.TestCheckResourceAttr("ibm_en_smtp_configuration.en_smtp_configuration_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_en_smtp_configuration.en_smtp_configuration_instance", "description", description),
					resource.TestCheckResourceAttr("ibm_en_smtp_configuration.en_smtp_configuration_instance", "domain", domain),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMEnSMTPConfigurationConfig(instanceID, nameUpdate, descriptionUpdate, domainUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_en_smtp_configuration.en_smtp_configuration_instance", "instance_id", instanceID),
					resource.TestCheckResourceAttr("ibm_en_smtp_configuration.en_smtp_configuration_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_en_smtp_configuration.en_smtp_configuration_instance", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_en_smtp_configuration.en_smtp_configuration_instance", "domain", domainUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_en_smtp_configuration.en_smtp_configuration",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMEnSMTPConfigurationConfigBasic(instanceID string, name string, domain string) string {
	return fmt.Sprintf(`
		resource "ibm_en_smtp_configuration" "en_smtp_configuration_instance" {
			instance_id = "%s"
			name = "%s"
			domain = "%s"
		}
	`, instanceID, name, domain)
}

func testAccCheckIBMEnSMTPConfigurationConfig(instanceID string, name string, description string, domain string) string {
	return fmt.Sprintf(`

		resource "ibm_en_smtp_configuration" "en_smtp_configuration_instance" {
			instance_id = "%s"
			name = "%s"
			description = "%s"
			domain = "%s"
		}
	`, instanceID, name, description, domain)
}

func testAccCheckIBMEnSMTPConfigurationExists(n string, obj eventnotificationsv1.SMTPConfiguration) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		eventNotificationsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).EventNotificationsApiV1()
		if err != nil {
			return err
		}

		getSMTPConfigurationOptions := &eventnotificationsv1.GetSMTPConfigurationOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getSMTPConfigurationOptions.SetInstanceID(parts[0])
		getSMTPConfigurationOptions.SetID(parts[1])

		smtpConfiguration, _, err := eventNotificationsClient.GetSMTPConfiguration(getSMTPConfigurationOptions)
		if err != nil {
			return err
		}

		obj = *smtpConfiguration
		return nil
	}
}

func testAccCheckIBMEnSMTPConfigurationDestroy(s *terraform.State) error {
	eventNotificationsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).EventNotificationsApiV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_en_smtp_configuration" {
			continue
		}

		getSMTPConfigurationOptions := &eventnotificationsv1.GetSMTPConfigurationOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getSMTPConfigurationOptions.SetInstanceID(parts[0])
		getSMTPConfigurationOptions.SetID(parts[1])

		// Try to find the key
		_, response, err := eventNotificationsClient.GetSMTPConfiguration(getSMTPConfigurationOptions)

		if err == nil {
			return fmt.Errorf("en_smtp_configuration still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for en_smtp_configuration (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
