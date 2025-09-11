// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSmPrivateCertificateDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmPrivateCertificateDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate.sm_private_certificate", "secret_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate.sm_private_certificate", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate.sm_private_certificate", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate.sm_private_certificate", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate.sm_private_certificate", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate.sm_private_certificate", "secret_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate.sm_private_certificate", "secret_type"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate.sm_private_certificate", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate.sm_private_certificate", "retrieved_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate.sm_private_certificate", "versions_total"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate.sm_private_certificate", "signing_algorithm"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate.sm_private_certificate", "certificate_template"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate.sm_private_certificate", "common_name"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate.sm_private_certificate", "expiration_date"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate.sm_private_certificate", "issuer"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate.sm_private_certificate", "serial_number"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate.sm_private_certificate", "validity.#"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate.sm_private_certificate", "certificate"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate.sm_private_certificate", "private_key"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate.sm_private_certificate_by_name", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate.sm_private_certificate_by_name", "secret_group_name"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_private_certificate.sm_private_certificate_by_name", "secret_id"),
				),
			},
		},
	})
}

func testAccCheckIbmSmPrivateCertificateDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_sm_private_certificate_configuration_root_ca" "ibm_sm_private_certificate_configuration_root_ca_instance" {
			instance_id   = "%s"
			region        = "%s"
			max_ttl = "180000"
			common_name = "ibm.com"
			crl_expiry = "10000h"
			name = "root-ca-terraform-private-cert-datasource-test"
		}
		resource "ibm_sm_private_certificate_configuration_intermediate_ca" "ibm_sm_private_certificate_configuration_intermediate_ca_instance" {
  			instance_id   = "%s"
			region        = "%s"
			max_ttl = "180000"
			common_name = "ibm.com"
			issuer = ibm_sm_private_certificate_configuration_root_ca.ibm_sm_private_certificate_configuration_root_ca_instance.name
			signing_method = "internal"
			name = "intermediate-ca-terraform-private-cert-datasource-test"
		}
		resource "ibm_sm_private_certificate_configuration_template" "sm_private_certificate_configuration_template_instance" {
			instance_id   = "%s"
			region        = "%s"
			certificate_authority = ibm_sm_private_certificate_configuration_intermediate_ca.ibm_sm_private_certificate_configuration_intermediate_ca_instance.name
			allow_any_name = true
			name = "template-terraform-private-cert-test"
		}

		resource "ibm_sm_private_certificate" "sm_private_certificate_instance" {
			instance_id = "%s"
			region = "%s"
		  	name = "private_cert_terraform-test"
  			description = "Extended description for this secret."
  			labels = [ "my-label", "another"]
  			custom_metadata = {"key":"value"}
  			certificate_template = ibm_sm_private_certificate_configuration_template.sm_private_certificate_configuration_template_instance.name
  			common_name = "ibm.com"
  			ttl = "1800"
  			secret_group_id = "default"
		}

		data "ibm_sm_private_certificate" "sm_private_certificate" {
			instance_id   = "%s"
			region        = "%s"
			secret_id = ibm_sm_private_certificate.sm_private_certificate_instance.secret_id
		}

		data "ibm_sm_private_certificate" "sm_private_certificate_by_name" {
			instance_id   = "%s"
			region = "%s"
			name = ibm_sm_private_certificate.sm_private_certificate_instance.name
			secret_group_name = "default"
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerInstanceID,
		acc.SecretsManagerInstanceRegion, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerInstanceID,
		acc.SecretsManagerInstanceRegion, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}
