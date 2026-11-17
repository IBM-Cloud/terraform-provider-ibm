// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.117.3-3e91ab11-20261028-182643
 */

package secretsmanagerinstancemanagement_test

import (
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/secretsmanagerinstancemanagement"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmSmHsmPathsDataSourceBasic(t *testing.T) {
	instanceID := hsmInstanceID(t, "SM_HSM_PPS_GATEWAY_CRN")
	name := hsmRandName()
	config := hsmDestinationConfig(instanceID) + hsmPathBlock(name, "") + `
		data "ibm_sm_hsm_paths" "test" {
			instance_id = ibm_sm_hsm_path.test.instance_id
			destination_id = ibm_sm_hsm_path.test.destination_id
		}
	`
	testAccHsmDataSource(t, config, "data.ibm_sm_hsm_paths.test",
		[]string{"id", "instance_id", "destination_id", "paths.#"},
		map[string]string{"paths.0.name": name, "paths.0.crn": acc.SecretsManagerHsmPpsGatewayCrn},
	)
}

func TestAccIbmSmHsmPathsDataSourceAllArgs(t *testing.T) {
	instanceID := hsmInstanceID(t, "SM_HSM_PPS_GATEWAY_CRN")
	name := hsmRandName()
	description := "tf_description_" + name
	config := hsmDestinationConfig(instanceID) + hsmPathBlock(name, description) + `
		data "ibm_sm_hsm_paths" "test" {
			instance_id = ibm_sm_hsm_path.test.instance_id
			destination_id = ibm_sm_hsm_path.test.destination_id
		}
	`
	testAccHsmDataSource(t, config, "data.ibm_sm_hsm_paths.test",
		[]string{"id", "instance_id", "destination_id", "paths.#", "paths.0.id", "paths.0.href", "paths.0.provisioning_state", "paths.0.health_state", "paths.0.created_at", "paths.0.created_by"},
		map[string]string{"paths.0.name": name, "paths.0.description": description, "paths.0.crn": acc.SecretsManagerHsmPpsGatewayCrn},
	)
}

func TestDataSourceIbmSmHsmPathsPathToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "fa3d1234-0000-0000-0000-000000000001"
		model["href"] = "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/paths/fa3d1234-0000-0000-0000-000000000001"
		model["name"] = "hsm-node-1"
		model["description"] = "Primary HSM node for prod"
		model["crn"] = "crn:v1:bluemix:public:is:eu-fr2:a/customer123::private-path-service-gateway:hsm1"
		model["provisioning_state"] = "completed"
		model["health_state"] = "healthy"
		model["health_as_of"] = "2026-07-09T15:30:00.000Z"
		model["created_at"] = "2026-07-09T15:00:00.000Z"
		model["updated_at"] = "2026-07-09T15:30:00.000Z"
		model["created_by"] = "iam-ServiceId-abc123"
		model["updated_by"] = "iam-ServiceId-abc123"

		assert.Equal(t, result, model)
	}

	model := new(secretsmanagerinstancemanagementv2.Path)
	model.ID = core.StringPtr("fa3d1234-0000-0000-0000-000000000001")
	model.Href = core.StringPtr("https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/paths/fa3d1234-0000-0000-0000-000000000001")
	model.Name = core.StringPtr("hsm-node-1")
	model.Description = core.StringPtr("Primary HSM node for prod")
	model.Crn = core.StringPtr("crn:v1:bluemix:public:is:eu-fr2:a/customer123::private-path-service-gateway:hsm1")
	model.ProvisioningState = core.StringPtr("completed")
	model.HealthState = core.StringPtr("healthy")
	model.HealthAsOf = CreateMockDateTime("2026-07-09T15:30:00.000Z")
	model.CreatedAt = CreateMockDateTime("2026-07-09T15:00:00.000Z")
	model.UpdatedAt = CreateMockDateTime("2026-07-09T15:30:00.000Z")
	model.CreatedBy = core.StringPtr("iam-ServiceId-abc123")
	model.UpdatedBy = core.StringPtr("iam-ServiceId-abc123")

	result, err := secretsmanagerinstancemanagement.DataSourceIbmSmHsmPathsPathToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
