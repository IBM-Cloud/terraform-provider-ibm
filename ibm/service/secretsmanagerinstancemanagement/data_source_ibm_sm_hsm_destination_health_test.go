// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.117.3-3e91ab11-20261028-182643
 */

package secretsmanagerinstancemanagement_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/secretsmanagerinstancemanagement"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmSmHsmDestinationHealthDataSourceBasic(t *testing.T) {
	instanceID := hsmInstanceID(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmHsmDestinationHealthDataSourceConfigBasic(instanceID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destination_health.sm_hsm_destination_health_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destination_health.sm_hsm_destination_health_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destination_health.sm_hsm_destination_health_instance", "destination_id"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destination_health.sm_hsm_destination_health_instance", "health_state"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destination_health.sm_hsm_destination_health_instance", "health_as_of"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destination_health.sm_hsm_destination_health_instance", "client_certificates.#"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destination_health.sm_hsm_destination_health_instance", "ha_groups.#"),
					resource.TestCheckResourceAttrSet("data.ibm_sm_hsm_destination_health.sm_hsm_destination_health_instance", "paths.#"),
				),
			},
		},
	})
}

func testAccCheckIbmSmHsmDestinationHealthDataSourceConfigBasic(instanceID string) string {
	return hsmDestinationConfig(instanceID) + `
		data "ibm_sm_hsm_destination_health" "sm_hsm_destination_health_instance" {
			instance_id = ibm_sm_hsm_destination.parent.instance_id
			destination_id = ibm_sm_hsm_destination.parent.destination_id
		}
	`
}

func TestDataSourceIbmSmHsmDestinationHealthHealthSnapshotClientCertToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "7c9e6679-7425-40de-944b-e07fc1f90ae7"
		model["name"] = "hsm-client-cert-prod"
		model["health_state"] = "valid"

		assert.Equal(t, result, model)
	}

	model := new(secretsmanagerinstancemanagementv2.HealthSnapshotClientCert)
	model.ID = core.StringPtr("7c9e6679-7425-40de-944b-e07fc1f90ae7")
	model.Name = core.StringPtr("hsm-client-cert-prod")
	model.HealthState = core.StringPtr("valid")

	result, err := secretsmanagerinstancemanagement.DataSourceIbmSmHsmDestinationHealthHealthSnapshotClientCertToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmSmHsmDestinationHealthHealthSnapshotHSMHAGroupToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "ee52ebb6-1728-4580-8ede-13f6504e3ae0"
		model["name"] = "production-ha-group"
		model["health_state"] = "healthy"

		assert.Equal(t, result, model)
	}

	model := new(secretsmanagerinstancemanagementv2.HealthSnapshotHSMHAGroup)
	model.ID = core.StringPtr("ee52ebb6-1728-4580-8ede-13f6504e3ae0")
	model.Name = core.StringPtr("production-ha-group")
	model.HealthState = core.StringPtr("healthy")

	result, err := secretsmanagerinstancemanagement.DataSourceIbmSmHsmDestinationHealthHealthSnapshotHSMHAGroupToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmSmHsmDestinationHealthHealthSnapshotConnectionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		healthSnapshotServerCaModel := make(map[string]interface{})
		healthSnapshotServerCaModel["id"] = "5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f"
		healthSnapshotServerCaModel["name"] = "HSM1_Server_CA"
		healthSnapshotServerCaModel["health_state"] = "valid"

		model := make(map[string]interface{})
		model["id"] = "fa3d1234-0000-0000-0000-000000000001"
		model["name"] = "hsm-node-1"
		model["health_state"] = "healthy"
		model["server_cas"] = []map[string]interface{}{healthSnapshotServerCaModel}

		assert.Equal(t, result, model)
	}

	healthSnapshotServerCaModel := new(secretsmanagerinstancemanagementv2.HealthSnapshotServerCa)
	healthSnapshotServerCaModel.ID = core.StringPtr("5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f")
	healthSnapshotServerCaModel.Name = core.StringPtr("HSM1_Server_CA")
	healthSnapshotServerCaModel.HealthState = core.StringPtr("valid")

	model := new(secretsmanagerinstancemanagementv2.HealthSnapshotConnection)
	model.ID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
	model.Name = core.StringPtr("hsm-node-1")
	model.HealthState = core.StringPtr("healthy")
	model.ServerCas = []secretsmanagerinstancemanagementv2.HealthSnapshotServerCa{*healthSnapshotServerCaModel}

	result, err := secretsmanagerinstancemanagement.DataSourceIbmSmHsmDestinationHealthHealthSnapshotConnectionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmSmHsmDestinationHealthHealthSnapshotServerCaToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f"
		model["name"] = "HSM1_Server_CA"
		model["health_state"] = "valid"

		assert.Equal(t, result, model)
	}

	model := new(secretsmanagerinstancemanagementv2.HealthSnapshotServerCa)
	model.ID = core.StringPtr("5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f")
	model.Name = core.StringPtr("HSM1_Server_CA")
	model.HealthState = core.StringPtr("valid")

	result, err := secretsmanagerinstancemanagement.DataSourceIbmSmHsmDestinationHealthHealthSnapshotServerCaToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
