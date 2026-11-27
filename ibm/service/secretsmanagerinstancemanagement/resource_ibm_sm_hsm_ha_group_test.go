// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanagerinstancemanagement_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
)

func TestAccIbmSmHsmHaGroup(t *testing.T) {
	testAccHsmResource(t, hsmResourceSpec{
		resourceType: "ibm_sm_hsm_ha_group",
		extraEnv:     []string{"SM_HSM_MEMBER_SERIALS"},
		config: func(instanceID, name, description string) string {
			return hsmDestinationConfig(instanceID) + fmt.Sprintf(`
		resource "ibm_sm_hsm_ha_group" "test" {
			instance_id = ibm_sm_hsm_destination.parent.instance_id
			destination_id = ibm_sm_hsm_destination.parent.destination_id
			name = "%s"
			%s
			group_label = "%s"
			group_serial = "%s"
			member_serials = %s
		}
	`, name, hsmDescriptionLine(description), hsmGroupLabel(), hsmGroupSerial(), hsmMemberSerialsHCL())
		},
		extraChecks: []resource.TestCheckFunc{
			resource.TestCheckResourceAttr("ibm_sm_hsm_ha_group.test", "group_label", hsmGroupLabel()),
			resource.TestCheckResourceAttr("ibm_sm_hsm_ha_group.test", "group_serial", hsmGroupSerial()),
			resource.TestCheckResourceAttr("ibm_sm_hsm_ha_group.test", "member_serials.#", fmt.Sprintf("%d", len(hsmMemberSerials()))),
			resource.TestCheckResourceAttrSet("ibm_sm_hsm_ha_group.test", "destination_id"),
			resource.TestCheckResourceAttrSet("ibm_sm_hsm_ha_group.test", "ha_group_id"),
		},
		get: func(client *secretsmanagerinstancemanagementv2.SecretsManagerInstanceManagementV2, parts []string) (*core.DetailedResponse, error) {
			options := &secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationHaGroupOptions{}
			options.SetInstanceID(parts[0])
			options.SetDestinationID(parts[1])
			options.SetHaGroupID(parts[2])
			_, response, err := client.GetInstanceHsmDestinationHaGroup(options)
			return response, err
		},
	}, true)
}
