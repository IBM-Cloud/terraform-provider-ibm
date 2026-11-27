// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.117.3-3e91ab11-20261028-182643
 */

package secretsmanagerinstancemanagement_test

import (
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSmHsmPathDataSourceBasic(t *testing.T) {
	instanceID := hsmInstanceID(t, "SM_HSM_PPS_GATEWAY_CRN")
	name := hsmRandName()
	config := hsmDestinationConfig(instanceID) + hsmPathBlock(name, "") + `
		data "ibm_sm_hsm_path" "test" {
			instance_id = ibm_sm_hsm_path.test.instance_id
			destination_id = ibm_sm_hsm_path.test.destination_id
			path_id = ibm_sm_hsm_path.test.path_id
		}
	`
	testAccHsmDataSource(t, config, "data.ibm_sm_hsm_path.test",
		[]string{"id", "instance_id", "destination_id", "path_id", "name", "crn", "provisioning_state", "health_state", "created_at", "created_by"},
		map[string]string{"name": name, "crn": acc.SecretsManagerHsmPpsGatewayCrn},
	)
}

func TestAccIbmSmHsmPathDataSourceAllArgs(t *testing.T) {
	instanceID := hsmInstanceID(t, "SM_HSM_PPS_GATEWAY_CRN")
	name := hsmRandName()
	description := "tf_description_" + name
	config := hsmDestinationConfig(instanceID) + hsmPathBlock(name, description) + `
		data "ibm_sm_hsm_path" "test" {
			instance_id = ibm_sm_hsm_path.test.instance_id
			destination_id = ibm_sm_hsm_path.test.destination_id
			path_id = ibm_sm_hsm_path.test.path_id
		}
	`
	testAccHsmDataSource(t, config, "data.ibm_sm_hsm_path.test",
		[]string{"id", "instance_id", "destination_id", "path_id", "name", "description", "crn", "provisioning_state", "health_state", "created_at", "created_by"},
		map[string]string{"name": name, "description": description, "crn": acc.SecretsManagerHsmPpsGatewayCrn},
	)
}
