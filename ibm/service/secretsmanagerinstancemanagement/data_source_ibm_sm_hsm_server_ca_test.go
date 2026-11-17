// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.117.3-3e91ab11-20261028-182643
 */

package secretsmanagerinstancemanagement_test

import (
	"testing"
)

func TestAccIbmSmHsmServerCaDataSourceBasic(t *testing.T) {
	instanceID := hsmInstanceID(t, "SM_HSM_PPS_GATEWAY_CRN", "SM_HSM_SERVER_CA_PEM_PATH")
	name := hsmRandName()
	config := hsmDestinationConfig(instanceID) + hsmServerCaBlock(name, "") + `
		data "ibm_sm_hsm_server_ca" "test" {
			instance_id = ibm_sm_hsm_server_ca.test.instance_id
			destination_id = ibm_sm_hsm_server_ca.test.destination_id
			path_id = ibm_sm_hsm_server_ca.test.path_id
			server_ca_id = ibm_sm_hsm_server_ca.test.server_ca_id
		}
	`
	testAccHsmDataSource(t, config, "data.ibm_sm_hsm_server_ca.test",
		[]string{"id", "instance_id", "destination_id", "path_id", "server_ca_id", "name", "expires_at", "status", "health_state", "created_at", "created_by"},
		map[string]string{"name": name},
	)
}

func TestAccIbmSmHsmServerCaDataSourceAllArgs(t *testing.T) {
	instanceID := hsmInstanceID(t, "SM_HSM_PPS_GATEWAY_CRN", "SM_HSM_SERVER_CA_PEM_PATH")
	name := hsmRandName()
	description := "tf_description_" + name
	config := hsmDestinationConfig(instanceID) + hsmServerCaBlock(name, description) + `
		data "ibm_sm_hsm_server_ca" "test" {
			instance_id = ibm_sm_hsm_server_ca.test.instance_id
			destination_id = ibm_sm_hsm_server_ca.test.destination_id
			path_id = ibm_sm_hsm_server_ca.test.path_id
			server_ca_id = ibm_sm_hsm_server_ca.test.server_ca_id
		}
	`
	testAccHsmDataSource(t, config, "data.ibm_sm_hsm_server_ca.test",
		[]string{"id", "instance_id", "destination_id", "path_id", "server_ca_id", "name", "description", "expires_at", "status", "health_state", "created_at", "created_by"},
		map[string]string{"name": name, "description": description},
	)
}
