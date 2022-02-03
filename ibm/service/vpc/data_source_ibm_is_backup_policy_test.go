// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIsBackupPolicyDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIsBackupPolicyDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "match_resource_types.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "match_user_tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "plans.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "resource_group.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "resource_type"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIsBackupPolicyDataSourceConfigBasicName(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "match_resource_types.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "match_user_tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "plans.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "resource_group.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_backup_policy.is_backup_policy", "resource_type"),
				),
			},
		},
	})
}

func testAccCheckIBMIsBackupPolicyDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_is_backup_policy" "is_backup_policy" {
			identifier = "r134-59921a4a-bd6e-4b0c-88e4-2dff44c9142f"
		}
	`)
}

func testAccCheckIBMIsBackupPolicyDataSourceConfigBasicName() string {
	return fmt.Sprintf(`
		data "ibm_is_backup_policy" "is_backup_policy" {
			name = "copier-darkness-submerge-parmesan"
		}
	`)
}
