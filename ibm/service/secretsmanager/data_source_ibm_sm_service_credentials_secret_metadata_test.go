// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSmServiceCredentialsSecretMetadataDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmServiceCredentialsSecretMetadataDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_service_credentials_secret_metadata.sm_service_credentials_secret_metadata", "secret_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_service_credentials_secret_metadata.sm_service_credentials_secret_metadata", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_service_credentials_secret_metadata.sm_service_credentials_secret_metadata", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_service_credentials_secret_metadata.sm_service_credentials_secret_metadata", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_service_credentials_secret_metadata.sm_service_credentials_secret_metadata", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_service_credentials_secret_metadata.sm_service_credentials_secret_metadata", "secret_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_service_credentials_secret_metadata.sm_service_credentials_secret_metadata", "secret_type"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_service_credentials_secret_metadata.sm_service_credentials_secret_metadata", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_service_credentials_secret_metadata.sm_service_credentials_secret_metadata", "retrieved_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_service_credentials_secret_metadata.sm_service_credentials_secret_metadata", "versions_total"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_service_credentials_secret_metadata.sm_service_credentials_secret_metadata", "rotation.#"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_service_credentials_secret_metadata.sm_service_credentials_secret_metadata", "ttl"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_service_credentials_secret_metadata.sm_service_credentials_secret_metadata", "source_service.#"),
				),
			},
		},
	})
}

func testAccCheckIbmSmServiceCredentialsSecretMetadataDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_sm_service_credentials_secret" "sm_service_credentials_secret_instance" {
			  instance_id   = "%s"
              region        = "%s"
              custom_metadata = {"key":"value"}
              description = "Extended description for this secret."
              labels = ["my-label"]
              rotation {
                auto_rotate = true
                interval = 1
                unit = "day"
              }
              secret_group_id = "default"
			  name = "service_credentials-datasource-terraform-test"
			  ttl = "%s"
   			  source_service {
			  	instance {
					crn = "%s"
				}
				role {
					crn = "%s"
				}
			}
		}

		data "ibm_sm_service_credentials_secret_metadata" "sm_service_credentials_secret_metadata" {
            instance_id = "%s"
			region = "%s"
			secret_id = ibm_sm_service_credentials_secret.sm_service_credentials_secret_instance.secret_id
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, serviceCredentialsTtl, acc.SecretsManagerENInstanceCrn, serviceCredentialsRoleCrn, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}
