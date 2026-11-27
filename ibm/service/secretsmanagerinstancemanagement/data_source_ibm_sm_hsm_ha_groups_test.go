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

func TestAccIbmSmHsmHaGroupsDataSourceBasic(t *testing.T) {
	instanceID := hsmInstanceID(t, "SM_HSM_MEMBER_SERIALS")
	name := hsmRandName()
	config := hsmDestinationConfig(instanceID) + hsmHaGroupBlock(name, "") + `
		data "ibm_sm_hsm_ha_groups" "test" {
			instance_id = ibm_sm_hsm_ha_group.test.instance_id
			destination_id = ibm_sm_hsm_ha_group.test.destination_id
		}
	`
	testAccHsmDataSource(t, config, "data.ibm_sm_hsm_ha_groups.test",
		[]string{"id", "instance_id", "destination_id", "ha_groups.#"},
		map[string]string{"ha_groups.0.name": name, "ha_groups.0.group_label": hsmGroupLabel(), "ha_groups.0.group_serial": hsmGroupSerial()},
	)
}

func TestAccIbmSmHsmHaGroupsDataSourceAllArgs(t *testing.T) {
	instanceID := hsmInstanceID(t, "SM_HSM_MEMBER_SERIALS")
	name := hsmRandName()
	description := "tf_description_" + name
	config := hsmDestinationConfig(instanceID) + hsmHaGroupBlock(name, description) + `
		data "ibm_sm_hsm_ha_groups" "test" {
			instance_id = ibm_sm_hsm_ha_group.test.instance_id
			destination_id = ibm_sm_hsm_ha_group.test.destination_id
		}
	`
	testAccHsmDataSource(t, config, "data.ibm_sm_hsm_ha_groups.test",
		[]string{"id", "instance_id", "destination_id", "ha_groups.#", "ha_groups.0.id", "ha_groups.0.href", "ha_groups.0.provisioning_state", "ha_groups.0.health_state", "ha_groups.0.health_as_of", "ha_groups.0.created_at", "ha_groups.0.created_by"},
		map[string]string{"ha_groups.0.name": name, "ha_groups.0.description": description, "ha_groups.0.group_label": hsmGroupLabel(), "ha_groups.0.group_serial": hsmGroupSerial()},
	)
}

func TestDataSourceIbmSmHsmHaGroupsHSMHAGroupToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "ee52ebb6-1728-4580-8ede-13f6504e3ae0"
		model["href"] = "https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/ha_groups/ee52ebb6-1728-4580-8ede-13f6504e3ae0"
		model["name"] = "production-ha-group"
		model["description"] = "Primary HA Group for prod"
		model["provisioning_state"] = "completed"
		model["health_state"] = "healthy"
		model["health_as_of"] = "2026-07-09T16:00:00.000Z"
		model["group_label"] = "prod_ha_group"
		model["group_serial"] = "11323465009866"
		model["member_serials"] = []string{"123456", "123457"}
		model["created_at"] = "2026-07-09T15:00:00.000Z"
		model["updated_at"] = "2026-07-09T16:00:00.000Z"
		model["created_by"] = "iam-ServiceId-abc123"
		model["updated_by"] = "iam-ServiceId-abc123"

		assert.Equal(t, result, model)
	}

	model := new(secretsmanagerinstancemanagementv2.HSMHAGroup)
	model.ID = core.StringPtr("ee52ebb6-1728-4580-8ede-13f6504e3ae0")
	model.Href = core.StringPtr("https://us-south.secrets-manager.cloud.ibm.com/v2/instances/bfc50c2e-d66d-4f37-9ccf-9713f8325b39/hsm_destinations/b49ad385-0cb7-3694-a71c-2579d6e94b02/ha_groups/ee52ebb6-1728-4580-8ede-13f6504e3ae0")
	model.Name = core.StringPtr("production-ha-group")
	model.Description = core.StringPtr("Primary HA Group for prod")
	model.ProvisioningState = core.StringPtr("completed")
	model.HealthState = core.StringPtr("healthy")
	model.HealthAsOf = CreateMockDateTime("2026-07-09T16:00:00.000Z")
	model.GroupLabel = core.StringPtr("prod_ha_group")
	model.GroupSerial = core.StringPtr("11323465009866")
	model.MemberSerials = []string{"123456", "123457"}
	model.CreatedAt = CreateMockDateTime("2026-07-09T15:00:00.000Z")
	model.UpdatedAt = CreateMockDateTime("2026-07-09T16:00:00.000Z")
	model.CreatedBy = core.StringPtr("iam-ServiceId-abc123")
	model.UpdatedBy = core.StringPtr("iam-ServiceId-abc123")

	result, err := secretsmanagerinstancemanagement.DataSourceIbmSmHsmHaGroupsHSMHAGroupToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
