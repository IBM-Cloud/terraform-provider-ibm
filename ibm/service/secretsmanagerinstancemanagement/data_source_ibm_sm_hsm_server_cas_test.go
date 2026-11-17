// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.117.3-3e91ab11-20261028-182643
 */

package secretsmanagerinstancemanagement_test

import (
	"testing"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/secretsmanagerinstancemanagement"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmSmHsmServerCasDataSourceBasic(t *testing.T) {
	instanceID := hsmInstanceID(t, "SM_HSM_PPS_GATEWAY_CRN", "SM_HSM_SERVER_CA_PEM_PATH")
	name := hsmRandName()
	config := hsmDestinationConfig(instanceID) + hsmServerCaBlock(name, "") + `
		data "ibm_sm_hsm_server_cas" "test" {
			instance_id = ibm_sm_hsm_server_ca.test.instance_id
			destination_id = ibm_sm_hsm_server_ca.test.destination_id
			path_id = ibm_sm_hsm_server_ca.test.path_id
		}
	`
	testAccHsmDataSource(t, config, "data.ibm_sm_hsm_server_cas.test",
		[]string{"id", "instance_id", "destination_id", "path_id", "server_cas.#"},
		map[string]string{"server_cas.0.name": name},
	)
}

func TestAccIbmSmHsmServerCasDataSourceAllArgs(t *testing.T) {
	instanceID := hsmInstanceID(t, "SM_HSM_PPS_GATEWAY_CRN", "SM_HSM_SERVER_CA_PEM_PATH")
	name := hsmRandName()
	description := "tf_description_" + name
	config := hsmDestinationConfig(instanceID) + hsmServerCaBlock(name, description) + `
		data "ibm_sm_hsm_server_cas" "test" {
			instance_id = ibm_sm_hsm_server_ca.test.instance_id
			destination_id = ibm_sm_hsm_server_ca.test.destination_id
			path_id = ibm_sm_hsm_server_ca.test.path_id
		}
	`
	testAccHsmDataSource(t, config, "data.ibm_sm_hsm_server_cas.test",
		[]string{"id", "instance_id", "destination_id", "path_id", "server_cas.#", "server_cas.0.expires_at", "server_cas.0.status", "server_cas.0.health_state", "server_cas.0.created_at", "server_cas.0.created_by"},
		map[string]string{"server_cas.0.name": name, "server_cas.0.description": description},
	)
}

func TestDataSourceIbmSmHsmServerCasServerCaToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f"
		model["name"] = "HSM1_Server_CA"
		model["description"] = "Primary CA for HSM node 1"
		model["certificate"] = "-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWg...\n-----END CERTIFICATE-----"
		model["expires_at"] = "2028-07-09T15:10:00.000Z"
		model["status"] = "active"
		model["health_state"] = "valid"
		model["health_as_of"] = "2026-07-09T15:30:00.000Z"
		model["created_at"] = "2026-07-09T15:10:00.000Z"
		model["updated_at"] = "2026-07-09T15:20:00.000Z"
		model["created_by"] = "iam-ServiceId-abc123"
		model["updated_by"] = "iam-ServiceId-abc123"

		assert.Equal(t, result, model)
	}

	model := new(secretsmanagerinstancemanagementv2.ServerCa)
	model.ID = core.StringPtr("5d1f3c2e-8a4b-4c6d-9e7f-0a1b2c3d4e5f")
	model.Name = core.StringPtr("HSM1_Server_CA")
	model.Description = core.StringPtr("Primary CA for HSM node 1")
	model.Certificate = core.StringPtr("-----BEGIN CERTIFICATE-----\nMIIDXTCCAkWg...\n-----END CERTIFICATE-----")
	model.ExpiresAt = CreateMockDateTime("2028-07-09T15:10:00.000Z")
	model.Status = core.StringPtr("active")
	model.HealthState = core.StringPtr("valid")
	model.HealthAsOf = CreateMockDateTime("2026-07-09T15:30:00.000Z")
	model.CreatedAt = CreateMockDateTime("2026-07-09T15:10:00.000Z")
	model.UpdatedAt = CreateMockDateTime("2026-07-09T15:20:00.000Z")
	model.CreatedBy = core.StringPtr("iam-ServiceId-abc123")
	model.UpdatedBy = core.StringPtr("iam-ServiceId-abc123")

	result, err := secretsmanagerinstancemanagement.DataSourceIbmSmHsmServerCasServerCaToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
