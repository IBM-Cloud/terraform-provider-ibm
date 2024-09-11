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

func TestAccIbmBaasSourceRegistrationDataSourceBasic(t *testing.T) {
	sourceRegistrationReponseParamsXIBMTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	sourceRegistrationReponseParamsEnvironment := "kPhysical"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasSourceRegistrationDataSourceConfigBasic(sourceRegistrationReponseParamsXIBMTenantID, sourceRegistrationReponseParamsEnvironment),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registration.baas_source_registration_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registration.baas_source_registration_instance", "source_registration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registration.baas_source_registration_instance", "x_ibm_tenant_id"),
				),
			},
		},
	})
}

func TestAccIbmBaasSourceRegistrationDataSourceAllArgs(t *testing.T) {
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
				Config: testAccCheckIbmBaasSourceRegistrationDataSourceConfig(sourceRegistrationReponseParamsXIBMTenantID, sourceRegistrationReponseParamsEnvironment, sourceRegistrationReponseParamsName, sourceRegistrationReponseParamsConnectionID, sourceRegistrationReponseParamsConnectorGroupID, sourceRegistrationReponseParamsDataSourceConnectionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registration.baas_source_registration_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registration.baas_source_registration_instance", "source_registration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registration.baas_source_registration_instance", "x_ibm_tenant_id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registration.baas_source_registration_instance", "request_initiator_type"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registration.baas_source_registration_instance", "source_id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registration.baas_source_registration_instance", "source_info.#"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registration.baas_source_registration_instance", "environment"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registration.baas_source_registration_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registration.baas_source_registration_instance", "connection_id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registration.baas_source_registration_instance", "connections.#"),
					resource.TestCheckResourceAttr("data.ibm_baas_source_registration.baas_source_registration_instance", "connections.0.connection_id", sourceRegistrationReponseParamsConnectionID),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registration.baas_source_registration_instance", "connections.0.entity_id"),
					resource.TestCheckResourceAttr("data.ibm_baas_source_registration.baas_source_registration_instance", "connections.0.connector_group_id", sourceRegistrationReponseParamsConnectorGroupID),
					resource.TestCheckResourceAttr("data.ibm_baas_source_registration.baas_source_registration_instance", "connections.0.data_source_connection_id", sourceRegistrationReponseParamsDataSourceConnectionID),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registration.baas_source_registration_instance", "connector_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registration.baas_source_registration_instance", "data_source_connection_id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registration.baas_source_registration_instance", "advanced_configs.#"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registration.baas_source_registration_instance", "advanced_configs.0.key"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registration.baas_source_registration_instance", "advanced_configs.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registration.baas_source_registration_instance", "authentication_status"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registration.baas_source_registration_instance", "registration_time_msecs"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registration.baas_source_registration_instance", "last_refreshed_time_msecs"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registration.baas_source_registration_instance", "external_metadata.#"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_source_registration.baas_source_registration_instance", "physical_params.#"),
				),
			},
		},
	})
}

func testAccCheckIbmBaasSourceRegistrationDataSourceConfigBasic(sourceRegistrationReponseParamsXIBMTenantID string, sourceRegistrationReponseParamsEnvironment string) string {
	return fmt.Sprintf(`
		resource "ibm_baas_source_registration" "baas_source_registration_instance" {
			x_ibm_tenant_id = "%s"
			environment = "%s"
		}

		data "ibm_baas_source_registration" "baas_source_registration_instance" {
			source_registration_id = 2
			x_ibm_tenant_id = ibm_baas_source_registration.baas_source_registration_instance.x_ibm_tenant_id
			request_initiator_type = "UIUser"
		}
	`, sourceRegistrationReponseParamsXIBMTenantID, sourceRegistrationReponseParamsEnvironment)
}

func testAccCheckIbmBaasSourceRegistrationDataSourceConfig(sourceRegistrationReponseParamsXIBMTenantID string, sourceRegistrationReponseParamsEnvironment string, sourceRegistrationReponseParamsName string, sourceRegistrationReponseParamsConnectionID string, sourceRegistrationReponseParamsConnectorGroupID string, sourceRegistrationReponseParamsDataSourceConnectionID string) string {
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

		data "ibm_baas_source_registration" "baas_source_registration_instance" {
			source_registration_id = 2
			x_ibm_tenant_id = ibm_baas_source_registration.baas_source_registration_instance.x_ibm_tenant_id
			request_initiator_type = "UIUser"
		}
	`, sourceRegistrationReponseParamsXIBMTenantID, sourceRegistrationReponseParamsEnvironment, sourceRegistrationReponseParamsName, sourceRegistrationReponseParamsConnectionID, sourceRegistrationReponseParamsConnectorGroupID, sourceRegistrationReponseParamsDataSourceConnectionID)
}
