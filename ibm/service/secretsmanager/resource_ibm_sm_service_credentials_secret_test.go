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

var serviceCredentialsSecretName = "terraform-test-sc-secret"
var modifiedServiceCredentialsSecretName = "modified-terraform-test-sc-secret"
var serviceCredentialsParametersWithServiceId = `{"serviceid_crn": ibm_iam_service_id.ibm_iam_service_id_instance.crn}`
var serviceCredentialsTtl = "172800"
var modifiedServiceCredentialsTtl = "6048000"
var serviceCredentialsRoleCrn = "crn:v1:bluemix:public:iam::::serviceRole:Writer"

func TestAccIbmSmServiceCredentialsSecretBasic(t *testing.T) {
	resourceName := "ibm_sm_service_credentials_secret.sm_service_credentials_secret_basic"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmServiceCredentialsSecretDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: serviceCredentialsSecretConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "secret_id"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					resource.TestCheckResourceAttrSet(resourceName, "crn"),
					resource.TestCheckResourceAttrSet(resourceName, "downloaded"),
					resource.TestCheckResourceAttr(resourceName, "state", "1"),
					resource.TestCheckResourceAttr(resourceName, "versions_total", "1"),
				),
			},
			resource.TestStep{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ttl"},
			},
		},
	})
}

func TestAccIbmSmServiceCredentialsSecretAllArgs(t *testing.T) {
	resourceName := "ibm_sm_service_credentials_secret.sm_service_credentials_secret"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmServiceCredentialsSecretDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: serviceCredentialsSecretConfigAllArgs(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmServiceCredentialsSecretCreated(resourceName),
					resource.TestCheckResourceAttrSet(resourceName, "secret_id"),
					resource.TestCheckResourceAttrSet(resourceName, "created_by"),
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					resource.TestCheckResourceAttrSet(resourceName, "crn"),
					resource.TestCheckResourceAttrSet(resourceName, "downloaded"),
					resource.TestCheckResourceAttrSet(resourceName, "next_rotation_date"),
					resource.TestCheckResourceAttr(resourceName, "state", "1"),
					resource.TestCheckResourceAttr(resourceName, "versions_total", "1"),
					resource.TestCheckResourceAttr(resourceName, "ttl", serviceCredentialsTtl),
				),
			},
			resource.TestStep{
				Config: serviceCredentialsSecretConfigUpdated(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmServiceCredentialsSecretUpdated(resourceName),
				),
			},
			resource.TestStep{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"ttl"},
			},
		},
	})
}

var serviceCredentialsSecretBasicConfigFormat = `
		resource "ibm_sm_service_credentials_secret" "sm_service_credentials_secret_basic" {
			instance_id   = "%s"
  			region        = "%s"
			name = "%s"
			source_service {
				instance {
					crn = "%s"
				}
				role {
					crn = "%s"
				}
			}
			ttl = "%s"
		}`

var serviceCredentialsSecretFullConfigFormat = `
		resource "ibm_sm_service_credentials_secret" "sm_service_credentials_secret" {
			instance_id   = "%s"
  			region        = "%s"
			name = "%s"
  			description = "%s"
  			labels = ["%s"]
			source_service {
				instance {
					crn = "%s"
				}
				parameters = %s
				role {
					crn = "%s"
				}
			}
			ttl = "%s"
  			custom_metadata = %s
			secret_group_id = "default"
			rotation %s
		}`

func iamServiceIdConfig() string {
	return fmt.Sprintf(`
		resource "ibm_iam_service_id" "ibm_iam_service_id_instance" {
			name = "service-id-terraform-tests-sc"
		}`)
}

func serviceCredentialsSecretConfigBasic() string {
	return fmt.Sprintf(serviceCredentialsSecretBasicConfigFormat, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		serviceCredentialsSecretName, acc.SecretsManagerENInstanceCrn, serviceCredentialsRoleCrn, serviceCredentialsTtl)
}

func serviceCredentialsSecretConfigAllArgs() string {
	return iamServiceIdConfig() + fmt.Sprintf(serviceCredentialsSecretFullConfigFormat, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		serviceCredentialsSecretName, description, label, acc.SecretsManagerENInstanceCrn, serviceCredentialsParametersWithServiceId, serviceCredentialsRoleCrn, serviceCredentialsTtl, customMetadata, rotationPolicy)
}

func serviceCredentialsSecretConfigUpdated() string {
	return iamServiceIdConfig() + fmt.Sprintf(serviceCredentialsSecretFullConfigFormat, acc.SecretsManagerInstanceID, acc.SecretsManagerInstanceRegion,
		modifiedServiceCredentialsSecretName, modifiedDescription, modifiedLabel, acc.SecretsManagerENInstanceCrn, serviceCredentialsParametersWithServiceId, serviceCredentialsRoleCrn,
		modifiedServiceCredentialsTtl, modifiedCustomMetadata, modifiedRotationPolicy)
}

func testAccCheckIbmSmServiceCredentialsSecretCreated(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		serviceCredentialsSecretIntf, err := getSecret(s, n)
		if err != nil {
			return err
		}
		secret := serviceCredentialsSecretIntf.(*secretsmanagerv2.ServiceCredentialsSecret)

		if err := verifyAttr(*secret.Name, serviceCredentialsSecretName, "secret name"); err != nil {
			return err
		}
		if err := verifyAttr(*secret.Description, description, "secret description"); err != nil {
			return err
		}
		if len(secret.Labels) != 1 {
			return fmt.Errorf("Wrong number of labels: %d", len(secret.Labels))
		}
		if err := verifyAttr(secret.Labels[0], label, "label"); err != nil {
			return err
		}
		if err := verifyJsonAttr(secret.CustomMetadata, customMetadata, "custom metadata"); err != nil {
			return err
		}
		if err := verifyAttr(getAutoRotate(secret.Rotation), "true", "auto_rotate"); err != nil {
			return err
		}
		if err := verifyAttr(getRotationUnit(secret.Rotation), "day", "rotation unit"); err != nil {
			return err
		}
		if err := verifyAttr(getRotationInterval(secret.Rotation), "1", "rotation interval"); err != nil {
			return err
		}
		if err := verifyAttr(*secret.TTL, serviceCredentialsTtl, "ttl"); err != nil {
			return err
		}
		if err := verifyAttr(*secret.SourceService.Instance.Crn, acc.SecretsManagerENInstanceCrn, "source_service.Instance.Crn"); err != nil {
			return err
		}
		if err := verifyAttr(*secret.SourceService.Role.Crn, serviceCredentialsRoleCrn, "source_service.Role.Crn"); err != nil {
			return err
		}
		if err := verifyAttr(*secret.Credentials.IamRoleCrn, serviceCredentialsRoleCrn, "credentials.IamRoleCrn"); err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckIbmSmServiceCredentialsSecretUpdated(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		serviceCredentialsSecretIntf, err := getSecret(s, n)
		if err != nil {
			return err
		}
		secret := serviceCredentialsSecretIntf.(*secretsmanagerv2.ServiceCredentialsSecret)

		if err := verifyAttr(*secret.Name, modifiedServiceCredentialsSecretName, "secret name"); err != nil {
			return err
		}
		if err := verifyAttr(*secret.Description, modifiedDescription, "secret description after update"); err != nil {
			return err
		}
		if len(secret.Labels) != 1 {
			return fmt.Errorf("Wrong number of labels after update: %d", len(secret.Labels))
		}
		if err := verifyAttr(secret.Labels[0], modifiedLabel, "label after update"); err != nil {
			return err
		}
		if err := verifyJsonAttr(secret.CustomMetadata, modifiedCustomMetadata, "custom metadata after update"); err != nil {
			return err
		}
		if err := verifyAttr(*secret.TTL, modifiedServiceCredentialsTtl, "ttl after update"); err != nil {
			return err
		}
		if err := verifyAttr(getAutoRotate(secret.Rotation), "true", "auto_rotate after update"); err != nil {
			return err
		}
		if err := verifyAttr(getRotationUnit(secret.Rotation), "month", "rotation unit after update"); err != nil {
			return err
		}
		if err := verifyAttr(getRotationInterval(secret.Rotation), "2", "rotation interval after update"); err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckIbmSmServiceCredentialsSecretDestroy(s *terraform.State) error {
	secretsManagerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerV2()
	if err != nil {
		return err
	}

	secretsManagerClient = getClientWithInstanceEndpointTest(secretsManagerClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_sm_service_credentials_secret" {
			continue
		}

		getSecretOptions := &secretsmanagerv2.GetSecretOptions{}

		id := strings.Split(rs.Primary.ID, "/")
		secretId := id[2]
		getSecretOptions.SetID(secretId)

		// Try to find the key
		_, response, err := secretsManagerClient.GetSecret(getSecretOptions)

		if err == nil {
			return fmt.Errorf("ServiceCredentialsSecret still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for ServiceCredentialsSecret (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
