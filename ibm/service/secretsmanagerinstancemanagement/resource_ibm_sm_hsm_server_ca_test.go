// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanagerinstancemanagement_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
)

func TestAccIbmSmHsmServerCa(t *testing.T) {
	testAccHsmResource(t, hsmResourceSpec{
		resourceType: "ibm_sm_hsm_server_ca",
		extraEnv:     []string{"SM_HSM_PPS_GATEWAY_CRN", "SM_HSM_SERVER_CA_PEM_PATH"},
		config: func(instanceID, name, description string) string {
			return hsmDestinationConfig(instanceID) + fmt.Sprintf(`
		resource "ibm_sm_hsm_path" "parent" {
			instance_id = ibm_sm_hsm_destination.parent.instance_id
			destination_id = ibm_sm_hsm_destination.parent.destination_id
			name = "tf_parent_path"
			crn = "%s"
		}

		resource "ibm_sm_hsm_server_ca" "test" {
			instance_id = ibm_sm_hsm_path.parent.instance_id
			destination_id = ibm_sm_hsm_path.parent.destination_id
			path_id = ibm_sm_hsm_path.parent.path_id
			name = "%s"
			%s
			certificate = file("%s")
		}
	`, acc.SecretsManagerHsmPpsGatewayCrn, name, hsmDescriptionLine(description), acc.SecretsManagerHsmServerCaPemPath)
		},
		extraChecks: []resource.TestCheckFunc{
			resource.TestCheckResourceAttrSet("ibm_sm_hsm_server_ca.test", "destination_id"),
			resource.TestCheckResourceAttrSet("ibm_sm_hsm_server_ca.test", "path_id"),
			resource.TestCheckResourceAttrSet("ibm_sm_hsm_server_ca.test", "server_ca_id"),
			resource.TestCheckResourceAttrSet("ibm_sm_hsm_server_ca.test", "expires_at"),
		},
		get: func(client *secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2, parts []string) (*core.DetailedResponse, error) {
			options := &secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationPathServerCaOptions{}
			options.SetInstanceID(parts[0])
			options.SetDestinationID(parts[1])
			options.SetPathID(parts[2])
			options.SetServerCaID(parts[3])
			_, response, err := client.GetInstanceHsmDestinationPathServerCa(options)
			return response, err
		},
	}, true)
}
