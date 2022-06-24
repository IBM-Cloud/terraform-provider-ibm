// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSccPostureGetCredentialDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccPostureGetCredentialDataSourceConfigBasic(scc_posture_credential_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_get_credential.get_credential", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_get_credential.get_credential", "enabled"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_get_credential.get_credential", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_get_credential.get_credential", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_get_credential.get_credential", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_get_credential.get_credential", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_get_credential.get_credential", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_get_credential.get_credential", "updated_by"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_get_credential.get_credential", "purpose"),
				),
			},
		},
	})
}

func testAccCheckIBMSccPostureGetCredentialDataSourceConfigBasic(credentialId string) string {
	return fmt.Sprintf(`
		data "ibm_scc_posture_get_credential" "get_credential" {
			id = "%s"
		}
	`, credentialId)
}
