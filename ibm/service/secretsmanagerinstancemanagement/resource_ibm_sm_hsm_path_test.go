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

func TestAccIbmSmHsmPath(t *testing.T) {
	testAccHsmResource(t, hsmResourceSpec{
		resourceType: "ibm_sm_hsm_path",
		extraEnv:     []string{"SM_HSM_PPS_GATEWAY_CRN"},
		config: func(instanceID, name, description string) string {
			return hsmDestinationConfig(instanceID) + fmt.Sprintf(`
		resource "ibm_sm_hsm_path" "test" {
			instance_id = ibm_sm_hsm_destination.parent.instance_id
			destination_id = ibm_sm_hsm_destination.parent.destination_id
			name = "%s"
			%s
			crn = "%s"
		}
	`, name, hsmDescriptionLine(description), acc.SecretsManagerHsmPpsGatewayCrn)
		},
		extraChecks: []resource.TestCheckFunc{
			resource.TestCheckResourceAttrSet("ibm_sm_hsm_path.test", "destination_id"),
			resource.TestCheckResourceAttrSet("ibm_sm_hsm_path.test", "path_id"),
			resource.TestCheckResourceAttrSet("ibm_sm_hsm_path.test", "crn"),
		},
		get: func(client *secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2, parts []string) (*core.DetailedResponse, error) {
			options := &secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationPathOptions{}
			options.SetInstanceID(parts[0])
			options.SetDestinationID(parts[1])
			options.SetPathID(parts[2])
			_, response, err := client.GetInstanceHsmDestinationPath(options)
			return response, err
		},
	}, true)
}
