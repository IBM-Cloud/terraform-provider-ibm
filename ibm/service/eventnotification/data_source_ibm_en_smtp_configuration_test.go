// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventnotification_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMEnSMTPConfigurationDataSourceBasic(t *testing.T) {
	smtpConfigurationInstanceID := fmt.Sprintf("tf_instance_id_%d", acctest.RandIntRange(10, 100))
	smtpConfigurationName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	smtpConfigurationDomain := fmt.Sprintf("tf_domain_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMEnSMTPConfigurationDataSourceConfigBasic(smtpConfigurationInstanceID, smtpConfigurationName, smtpConfigurationDomain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_configuration.en_smtp_configuration_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_configuration.en_smtp_configuration_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_configuration.en_smtp_configuration_instance", "en_smtp_configuration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_configuration.en_smtp_configuration_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_configuration.en_smtp_configuration_instance", "domain"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_configuration.en_smtp_configuration_instance", "config.#"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_configuration.en_smtp_configuration_instance", "updated_at"),
				),
			},
		},
	})
}

func TestAccIBMEnSMTPConfigurationDataSourceAllArgs(t *testing.T) {
	smtpConfigurationInstanceID := fmt.Sprintf("tf_instance_id_%d", acctest.RandIntRange(10, 100))
	smtpConfigurationName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	smtpConfigurationDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	smtpConfigurationDomain := fmt.Sprintf("tf_domain_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMEnSMTPConfigurationDataSourceConfig(smtpConfigurationInstanceID, smtpConfigurationName, smtpConfigurationDescription, smtpConfigurationDomain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_configuration.en_smtp_configuration_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_configuration.en_smtp_configuration_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_configuration.en_smtp_configuration_instance", "en_smtp_configuration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_configuration.en_smtp_configuration_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_configuration.en_smtp_configuration_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_configuration.en_smtp_configuration_instance", "domain"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_configuration.en_smtp_configuration_instance", "config.#"),
					resource.TestCheckResourceAttrSet("data.ibm_en_smtp_configuration.en_smtp_configuration_instance", "updated_at"),
				),
			},
		},
	})
}

func testAccCheckIBMEnSMTPConfigurationDataSourceConfigBasic(smtpConfigurationInstanceID string, smtpConfigurationName string, smtpConfigurationDomain string) string {
	return fmt.Sprintf(`
		resource "ibm_en_smtp_configuration" "en_smtp_configuration_instance" {
			instance_id = "%s"
			name = "%s"
			domain = "%s"
		}

		data "ibm_en_smtp_configuration" "en_smtp_configuration_instance" {
			instance_id = ibm_en_smtp_configuration.en_smtp_configuration_instance.instance_id
			en_smtp_configuration_id = ibm_en_smtp_configuration.en_smtp_configuration_instance.en_smtp_configuration_id
		}
	`, smtpConfigurationInstanceID, smtpConfigurationName, smtpConfigurationDomain)
}

func testAccCheckIBMEnSMTPConfigurationDataSourceConfig(smtpConfigurationInstanceID string, smtpConfigurationName string, smtpConfigurationDescription string, smtpConfigurationDomain string) string {
	return fmt.Sprintf(`
		resource "ibm_en_smtp_configuration" "en_smtp_configuration_instance" {
			instance_id = "%s"
			name = "%s"
			description = "%s"
			domain = "%s"
		}

		data "ibm_en_smtp_configuration" "en_smtp_configuration_instance" {
			instance_id = ibm_en_smtp_configuration.en_smtp_configuration_instance.instance_id
			en_smtp_configuration_id = ibm_en_smtp_configuration.en_smtp_configuration_instance.en_smtp_configuration_id
		}
	`, smtpConfigurationInstanceID, smtpConfigurationName, smtpConfigurationDescription, smtpConfigurationDomain)
}
