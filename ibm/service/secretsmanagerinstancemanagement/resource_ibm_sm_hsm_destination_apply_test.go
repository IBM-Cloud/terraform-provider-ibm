// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanagerinstancemanagement_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
)

func TestAccIbmSmHsmDestinationApplyBasic(t *testing.T) {
	var conf secretsmanagerinstancemanagementv2.DestinationHealthSnapshot
	instanceID := hsmInstanceID(t, "SM_HSM_PPS_GATEWAY_CRN", "SM_HSM_SERVER_CA_PEM_PATH")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmSmHsmDestinationApplyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: hsmDestinationConfig(instanceID) + fmt.Sprintf(`
					resource "ibm_sm_hsm_destination_apply" "test" {
						instance_id = ibm_sm_hsm_destination.test.instance_id
						destination_id = ibm_sm_hsm_destination.test.destination_id
					}
				`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmSmHsmDestinationApplyExists("ibm_sm_hsm_destination_apply.test", conf),
					resource.TestCheckResourceAttrPair("ibm_sm_hsm_destination_apply.test", "instance_id", "ibm_sm_hsm_destination.test", "instance_id"),
					resource.TestCheckResourceAttrPair("ibm_sm_hsm_destination_apply.test", "destination_id", "ibm_sm_hsm_destination.test", "destination_id"),
					resource.TestCheckResourceAttrSet("ibm_sm_hsm_destination_apply.test", "health_state"),
				),
			},
		},
	})
}

func testAccCheckIbmSmHsmDestinationApplyExists(n string, obj secretsmanagerinstancemanagementv2.DestinationHealthSnapshot) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		secretsManagerInstanceManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SecretsManagerInstanceManagementV2()
		if err != nil {
			return err
		}

		getHsmDestinationHealthSnapshotOptions := &secretsmanagerinstancemanagementv2.GetHsmDestinationHealthSnapshotOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getHsmDestinationHealthSnapshotOptions.SetInstanceID(parts[0])
		getHsmDestinationHealthSnapshotOptions.SetDestinationID(parts[1])

		healthSnapshot, _, err := secretsManagerInstanceManagementClient.GetHsmDestinationHealthSnapshot(getHsmDestinationHealthSnapshotOptions)
		if err != nil {
			return err
		}

		obj = *healthSnapshot
		return nil
	}
}

func testAccCheckIbmSmHsmDestinationApplyDestroy(s *terraform.State) error {
	// ibm_sm_hsm_destination_apply does not create a dedicated cloud object to delete;
	// it executes apply against the destination.
	return nil
}
