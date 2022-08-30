// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSccPostureCredentialDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccPostureCredentialDataSourceConfigBasic(acc.Scc_posture_credential_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_credential.credential", "credential_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_credential.credential", "enabled"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_credential.credential", "credential_type"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_credential.credential", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_credential.credential", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_credential.credential", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_credential.credential", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_credential.credential", "updated_by"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_credential.credential", "purpose"),
				),
			},
		},
	})
}

func testAccCheckIBMSccPostureCredentialDataSourceConfigBasic(credentialId string) string {
	return fmt.Sprintf(`
		data "ibm_scc_posture_credential" "credential" {
			credential_id = "%s"
		}
	`, credentialId)
}
