// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.94.0-fa797aec-20240814-142622
 */

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBaasSourceRegistrationsDataSourceBasic(t *testing.T) {
	sourceRegistrationReponseParamsXIBMTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	sourceRegistrationReponseParamsEnvironment := "kPhysical"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasSourceRegistrationsDataSourceConfigBasic(sourceRegistrationReponseParamsXIBMTenantID, sourceRegistrationReponseParamsEnvironment),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registrations.baas_source_registrations_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registrations.baas_source_registrations_instance", "x_ibm_tenant_id"),
				),
			},
		},
	})
}

func TestAccIbmBaasSourceRegistrationsDataSourceAllArgs(t *testing.T) {
	sourceRegistrationReponseParamsXIBMTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	sourceRegistrationReponseParamsEnvironment := "kPhysical"
	sourceRegistrationReponseParamsName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	sourceRegistrationReponseParamsConnectionID := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	sourceRegistrationReponseParamsConnectorGroupID := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	sourceRegistrationReponseParamsDataSourceConnectionID := fmt.Sprintf("tf_data_source_connection_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasSourceRegistrationsDataSourceConfig(sourceRegistrationReponseParamsXIBMTenantID, sourceRegistrationReponseParamsEnvironment, sourceRegistrationReponseParamsName, sourceRegistrationReponseParamsConnectionID, sourceRegistrationReponseParamsConnectorGroupID, sourceRegistrationReponseParamsDataSourceConnectionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registrations.baas_source_registrations_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registrations.baas_source_registrations_instance", "x_ibm_tenant_id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registrations.baas_source_registrations_instance", "ids"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registrations.baas_source_registrations_instance", "include_source_credentials"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registrations.baas_source_registrations_instance", "encryption_key"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registrations.baas_source_registrations_instance", "use_cached_data"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registrations.baas_source_registrations_instance", "include_external_metadata"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registrations.baas_source_registrations_instance", "ignore_tenant_migration_in_progress_check"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registrations.baas_source_registrations_instance", "registrations.#"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registrations.baas_source_registrations_instance", "registrations.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registrations.baas_source_registrations_instance", "registrations.0.source_id"),
					resource.TestCheckResourceAttr("data.ibm_baas_source_registrations.baas_source_registrations_instance", "registrations.0.environment", sourceRegistrationReponseParamsEnvironment),
					resource.TestCheckResourceAttr("data.ibm_baas_source_registrations.baas_source_registrations_instance", "registrations.0.name", sourceRegistrationReponseParamsName),
					resource.TestCheckResourceAttr("data.ibm_baas_source_registrations.baas_source_registrations_instance", "registrations.0.connection_id", sourceRegistrationReponseParamsConnectionID),
					resource.TestCheckResourceAttr("data.ibm_baas_source_registrations.baas_source_registrations_instance", "registrations.0.connector_group_id", sourceRegistrationReponseParamsConnectorGroupID),
					resource.TestCheckResourceAttr("data.ibm_baas_source_registrations.baas_source_registrations_instance", "registrations.0.data_source_connection_id", sourceRegistrationReponseParamsDataSourceConnectionID),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registrations.baas_source_registrations_instance", "registrations.0.authentication_status"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registrations.baas_source_registrations_instance", "registrations.0.registration_time_msecs"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registrations.baas_source_registrations_instance", "registrations.0.last_refreshed_time_msecs"),
				),
			},
		},
	})
}

func testAccCheckIbmBaasSourceRegistrationsDataSourceConfigBasic(sourceRegistrationReponseParamsXIBMTenantID string, sourceRegistrationReponseParamsEnvironment string) string {
	return fmt.Sprintf(`
		resource "ibm_baas_source_registration" "baas_source_registration_instance" {
			x_ibm_tenant_id = "%s"
			environment = "%s"
		}

		data "ibm_baas_source_registrations" "baas_source_registrations_instance" {
			x_ibm_tenant_id = ibm_baas_source_registration.baas_source_registration_instance.x_ibm_tenant_id
			ids = [ 1 ]
			include_source_credentials = true
			encryption_key = "encryption_key"
			use_cached_data = true
			include_external_metadata = true
			ignore_tenant_migration_in_progress_check = true
		}
	`, sourceRegistrationReponseParamsXIBMTenantID, sourceRegistrationReponseParamsEnvironment)
}

func testAccCheckIbmBaasSourceRegistrationsDataSourceConfig(sourceRegistrationReponseParamsXIBMTenantID string, sourceRegistrationReponseParamsEnvironment string, sourceRegistrationReponseParamsName string, sourceRegistrationReponseParamsConnectionID string, sourceRegistrationReponseParamsConnectorGroupID string, sourceRegistrationReponseParamsDataSourceConnectionID string) string {
	return fmt.Sprintf(`
		resource "ibm_baas_source_registration" "baas_source_registration_instance" {
			x_ibm_tenant_id = "%s"
			environment = "%s"
			name = "%s"
			connection_id = %s
			connections {
				connection_id = 1
				entity_id = 1
				connector_group_id = 1
				data_source_connection_id = "data_source_connection_id"
			}
			connector_group_id = %s
			data_source_connection_id = "%s"
			advanced_configs {
				key = "key"
				value = "value"
			}
			physical_params {
				endpoint = "endpoint"
				force_register = true
				host_type = "kLinux"
				physical_type = "kGroup"
				applications = [ "kSQL" ]
			}
		}

		data "ibm_baas_source_registrations" "baas_source_registrations_instance" {
			x_ibm_tenant_id = ibm_baas_source_registration.baas_source_registration_instance.x_ibm_tenant_id
			ids = [ 1 ]
			include_source_credentials = true
			encryption_key = "encryption_key"
			use_cached_data = true
			include_external_metadata = true
			ignore_tenant_migration_in_progress_check = true
		}
	`, sourceRegistrationReponseParamsXIBMTenantID, sourceRegistrationReponseParamsEnvironment, sourceRegistrationReponseParamsName, sourceRegistrationReponseParamsConnectionID, sourceRegistrationReponseParamsConnectorGroupID, sourceRegistrationReponseParamsDataSourceConnectionID)
}
