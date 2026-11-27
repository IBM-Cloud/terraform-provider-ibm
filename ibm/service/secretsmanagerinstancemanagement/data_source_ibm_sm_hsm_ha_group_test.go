// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.117.3-3e91ab11-20261028-182643
 */

package secretsmanagerinstancemanagement_test

import (
	"testing"
)

func TestAccIbmSmHsmHaGroupDataSourceBasic(t *testing.T) {
	instanceID := hsmInstanceID(t, "SM_HSM_MEMBER_SERIALS")
	name := hsmRandName()
	config := hsmDestinationConfig(instanceID) + hsmHaGroupBlock(name, "") + `
		data "ibm_sm_hsm_ha_group" "test" {
			instance_id = ibm_sm_hsm_ha_group.test.instance_id
			destination_id = ibm_sm_hsm_ha_group.test.destination_id
			ha_group_id = ibm_sm_hsm_ha_group.test.ha_group_id
		}
	`
	testAccHsmDataSource(t, config, "data.ibm_sm_hsm_ha_group.test",
		[]string{"id", "instance_id", "destination_id", "ha_group_id", "name", "provisioning_state", "health_state", "health_as_of", "group_label", "group_serial", "member_serials.#", "created_at", "created_by"},
		map[string]string{"name": name, "group_label": hsmGroupLabel(), "group_serial": hsmGroupSerial()},
	)
}

func TestAccIbmSmHsmHaGroupDataSourceAllArgs(t *testing.T) {
	instanceID := hsmInstanceID(t, "SM_HSM_MEMBER_SERIALS")
	name := hsmRandName()
	description := "tf_description_" + name
	config := hsmDestinationConfig(instanceID) + hsmHaGroupBlock(name, description) + `
		data "ibm_sm_hsm_ha_group" "test" {
			instance_id = ibm_sm_hsm_ha_group.test.instance_id
			destination_id = ibm_sm_hsm_ha_group.test.destination_id
			ha_group_id = ibm_sm_hsm_ha_group.test.ha_group_id
		}
	`
	testAccHsmDataSource(t, config, "data.ibm_sm_hsm_ha_group.test",
		[]string{"id", "instance_id", "destination_id", "ha_group_id", "name", "description", "provisioning_state", "health_state", "health_as_of", "group_label", "group_serial", "member_serials.#", "created_at", "created_by"},
		map[string]string{"name": name, "description": description, "group_label": hsmGroupLabel(), "group_serial": hsmGroupSerial()},
	)
}
