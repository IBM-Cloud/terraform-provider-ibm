// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.94.0-fa797aec-20240814-142622
 */

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBaasSourceRegistrationDataSourceBasic(t *testing.T) {
	environment := "kPhysical"
	connectionId := "4980716806983529472"
	endpoint := "172.26.1.24"
	hostType := "kLinux"
	physicalType := "kHost"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasSourceRegistrationDataSourceConfigBasic(environment, endpoint, hostType, physicalType, connectionId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_baas_source_registration.baas_source_registration_instance", "x_ibm_tenant_id", tenantId),
					resource.TestCheckResourceAttr("ibm_baas_source_registration.baas_source_registration_instance", "environment", environment),
					resource.TestCheckResourceAttrSet("ibm_baas_source_registration.baas_source_registration_instance", "id"),
					resource.TestCheckResourceAttrSet("ibm_baas_source_registration.baas_source_registration_instance", "last_refreshed_time_msecs"),
					resource.TestCheckResourceAttrSet("ibm_baas_source_registration.baas_source_registration_instance", "source_registration_id"),
					resource.TestCheckResourceAttrSet("ibm_baas_source_registration.baas_source_registration_instance", "connections.#"),
					resource.TestCheckResourceAttr("ibm_baas_source_registration.baas_source_registration_instance", "source_info.#", "1"),
					resource.TestCheckResourceAttr("ibm_baas_source_registration.baas_source_registration_instance", "source_info.0.name", endpoint),
					resource.TestCheckResourceAttrSet("ibm_baas_source_registration.baas_source_registration_instance", "source_info.0.id"),
					resource.TestCheckResourceAttr("ibm_baas_source_registration.baas_source_registration_instance", "source_info.0.os_type", hostType),
					resource.TestCheckResourceAttr("ibm_baas_source_registration.baas_source_registration_instance", "source_info.0.object_type", physicalType),
					resource.TestCheckResourceAttr("ibm_baas_source_registration.baas_source_registration_instance", "source_info.0.protection_type", "kVolume"),
					resource.TestCheckResourceAttr("ibm_baas_source_registration.baas_source_registration_instance", "physical_params.#", "1"),
					resource.TestCheckResourceAttr("ibm_baas_source_registration.baas_source_registration_instance", "physical_params.0.endpoint", endpoint),
					resource.TestCheckResourceAttr("ibm_baas_source_registration.baas_source_registration_instance", "physical_params.0.host_type", hostType),
					resource.TestCheckResourceAttr("ibm_baas_source_registration.baas_source_registration_instance", "physical_params.0.physical_type", physicalType),
				),
			},
		},
	})
}

func testAccCheckIbmBaasSourceRegistrationDataSourceConfigBasic(environment, endpoint, hostType, physicalType string, connectionId string) string {
	return fmt.Sprintf(`
			resource "ibm_baas_source_registration" "baas_source_registration_instance" {
				x_ibm_tenant_id = "%s"
				environment = "%s"
				connection_id = "%s"
				physical_params {
				endpoint = "%s"
				host_type = "%s"
				physical_type = "%s"
				}
			}

			data "ibm_baas_source_registration" "baas_source_registration_instance" {
				source_registration_id = ibm_baas_source_registration.baas_source_registration_instance.source_id
				x_ibm_tenant_id = ibm_baas_source_registration.baas_source_registration_instance.x_ibm_tenant_id
			}
	`, tenantId, environment, connectionId, endpoint, hostType, physicalType)
}
