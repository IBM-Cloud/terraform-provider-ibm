// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanager_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSmServiceCredentialsSecretDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmServiceCredentialsSecretDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_service_credentials_secret.sm_service_credentials_secret", "secret_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_service_credentials_secret.sm_service_credentials_secret", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_service_credentials_secret.sm_service_credentials_secret", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_service_credentials_secret.sm_service_credentials_secret", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_service_credentials_secret.sm_service_credentials_secret", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_service_credentials_secret.sm_service_credentials_secret", "secret_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_service_credentials_secret.sm_service_credentials_secret", "secret_type"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_service_credentials_secret.sm_service_credentials_secret", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_service_credentials_secret.sm_service_credentials_secret", "retrieved_at"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_service_credentials_secret.sm_service_credentials_secret", "versions_total"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_service_credentials_secret.sm_service_credentials_secret", "rotation.#"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_service_credentials_secret.sm_service_credentials_secret", "ttl"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_service_credentials_secret.sm_service_credentials_secret", "source_service.#"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_service_credentials_secret.sm_service_credentials_secret_by_name", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_service_credentials_secret.sm_service_credentials_secret_by_name", "secret_group_name"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_service_credentials_secret.sm_service_credentials_secret_by_name", "secret_id"),
				),
			},
		},
	})
}

func testAccCheckIbmSmServiceCredentialsSecretDataSourceConfigBasic() string {
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

		data "ibm_sm_service_credentials_secret" "sm_service_credentials_secret" {
            instance_id = "%s"
			region = "%s"
			secret_id = ibm_sm_service_credentials_secret.sm_service_credentials_secret_instance.secret_id
		}

		data "ibm_sm_service_credentials_secret" "sm_service_credentials_secret_by_name" {
			instance_id   = "%s"
			region = "%s"
			name = ibm_sm_service_credentials_secret.sm_service_credentials_secret_instance.name
			secret_group_name = "default"
		}
	`, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, serviceCredentialsTtl, acc.SecretsManagerENInstanceCrn, serviceCredentialsRoleCrn, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion)
}
