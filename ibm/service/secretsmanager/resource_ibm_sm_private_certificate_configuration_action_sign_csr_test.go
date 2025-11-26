// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/secrets-manager-go-sdk/v2/secretsmanagerv2"
)

func TestAccIbmSmPrivateCertificateConfigurationActionSignCsrBasic(t *testing.T) {
	var conf secretsmanagerv2.PrivateCertificateConfigurationIntermediateCA

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmPrivateCertificateConfigurationActionSignCsrConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmPrivateCertificateConfigurationActionSignCsrExists("ibm_sm_private_certificate_configuration_intermediate_ca.sm_private_certificate_configuration_intermediate_ca", conf),
				),
			},
		},
	})
}

func testAccCheckIbmSmPrivateCertificateConfigurationActionSignCsrConfigBasic() string {
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
			signing_method = "external"
			name = "intermediate-ca-terraform-external-intermediate-test"
		}
		
		resource "ibm_sm_private_certificate_configuration_action_sign_csr" "sm_private_certificate_configuration_action_sign_csr" {
  			instance_id   = "%s"
			region        = "%s"
			ttl = "2h"
			name = ibm_sm_private_certificate_configuration_root_ca.ibm_sm_private_certificate_configuration_root_ca_instance.name
			csr = ibm_sm_private_certificate_configuration_intermediate_ca.sm_private_certificate_configuration_intermediate_ca.data[0].csr
			max_path_length = 8
		}
		
		resource "ibm_sm_private_certificate_configuration_action_set_signed" "ibm_sm_private_certificate_configuration_action_set_signed_instance" {
			instance_id   = "%s"
			region        = "%s"
			name = ibm_sm_private_certificate_configuration_intermediate_ca.sm_private_certificate_configuration_intermediate_ca.name
			certificate = ibm_sm_private_certificate_configuration_action_sign_csr.sm_private_certificate_configuration_action_sign_csr.data[0].certificate
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}

func testAccCheckIbmSmPrivateCertificateConfigurationActionSignCsrExists(n string, obj secretsmanagerv2.PrivateCertificateConfigurationIntermediateCA) resource.TestCheckFunc {

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

		if *privateCertificateConfigurationIntermediateCA.Status != "certificate_template_required" {
			return fmt.Errorf(`attribute 'status' expected: "certificate_template_required", got: %s`, *privateCertificateConfigurationIntermediateCA.Status)
		}
		return nil
	}
}
