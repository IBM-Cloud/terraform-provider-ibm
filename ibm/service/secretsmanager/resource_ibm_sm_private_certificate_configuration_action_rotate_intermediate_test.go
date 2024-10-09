// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/secrets-manager-go-sdk/v2/secretsmanagerv2"
)

func TestAccIbmSmPrivateCertificateConfigurationActionRotateIntermediateBasic(t *testing.T) {
	var conf secretsmanagerv2.PrivateCertificateConfigurationIntermediateCA

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmPrivateCertificateConfigurationActionRotateIntermediateConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmPrivateCertificateConfigurationActionRotateIntermediateExists("intermediate-ca-terraform-external-intermediate-test", conf),
				),
			},
		},
	})
}

func testAccCheckIbmSmPrivateCertificateConfigurationActionRotateIntermediateConfigBasic() string {
	return fmt.Sprintf(`
		
		resource "ibm_sm_private_certificate_configuration_root_ca" "ibm_sm_private_certificate_configuration_root_ca_instance" {
			instance_id   = "%s"
			region        = "%s"
			max_ttl = "180000"
			common_name = "ibm.com"
			crl_expiry = "10000h"
			name = "root-ca-terraform-exteranl-intermediate-test"
		}
		resource "ibm_sm_private_certificate_configuration_intermediate_ca" "sm_private_certificate_configuration_intermediate_ca" {
  			instance_id   = "%s"
			region        = "%s"
			max_ttl = "180000"
			common_name = "ibm.com"
			signing_method = "internal"
			name = "intermediate-ca-terraform-external-intermediate-test"
		}
		
		resource "ibm_sm_private_certificate_configuration_action_sign_csr" "sm_private_certificate_configuration_action_sign_csr" {
  			instance_id   = "%s"
			region        = "%s"
			ttl = "2h"
			name = "root-ca-terraform-exteranl-intermediate-test"
			csr = ibm_sm_private_certificate_configuration_intermediate_ca.sm_private_certificate_configuration_intermediate_ca.data[0].csr
			max_path_length = 80
		}
		
		resource "ibm_sm_private_certificate_configuration_template" "sm_private_cert_template_basic" {
			instance_id   = "%s"
			region        = "%s"
			certificate_authority = ibm_sm_private_certificate_configuration_intermediate_ca.intermediate_ca_instance.name
			name = "template-terraform-test-basic"
		}

		resource "ibm_sm_private_certificate_configuration_action_rotate_intermediate" "sm_private_certificate_configuration_action_rotate_intermediate" {
  			instance_id   = "%s"
			region        = "%s"
			name = "intermediate-ca-terraform-external-intermediate-test"
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}

func testAccCheckIbmSmPrivateCertificateConfigurationActionRotateIntermediateExists(n string, obj secretsmanagerv2.PrivateCertificateConfigurationIntermediateCA) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		secretsManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerV2()
		if err != nil {
			return err
		}

		secretsManagerClient = getClientWithInstanceEndpointTest(secretsManagerClient)

		getConfigurationOptions := &secretsmanagerv2.GetConfigurationOptions{}

		id := strings.Split(rs.Primary.ID, "/")
		configName := id[2]
		getConfigurationOptions.SetName(configName)

		privateCertificateConfigurationIntermediateCAIntf, _, err := secretsManagerClient.GetConfiguration(getConfigurationOptions)
		if err != nil {
			return err
		}

		privateCertificateConfigurationIntermediateCA := privateCertificateConfigurationIntermediateCAIntf.(*secretsmanagerv2.PrivateCertificateConfigurationIntermediateCA)
		obj = *privateCertificateConfigurationIntermediateCA

		if *privateCertificateConfigurationIntermediateCA.Status != "configured" {
			return fmt.Errorf(`attribute 'status' expected: "configured", got: %s`, *privateCertificateConfigurationIntermediateCA.Status)
		}
		return nil
	}
}
