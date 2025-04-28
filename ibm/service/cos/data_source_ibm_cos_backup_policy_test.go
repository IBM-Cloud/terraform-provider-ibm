// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cos_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCOSBackupPolicyDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCOS(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMCOSBackupPolicyDataSourceConfig_basic_read(acc.BucketName, acc.CosBackupPolicyID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cos_backup_policy.policy", "bucket_name"),
					resource.TestCheckResourceAttrSet("data.ibm_cos_backup_policy.policy", "policy_id"),
					resource.TestCheckResourceAttr("data.ibm_cos_backup_policy.policy", "backup_type", "continuous"),
				),
			},
		},
	})
}

func testAccIBMCOSBackupPolicyDataSourceConfig_basic_read(name string, policy_id string) string {
	return fmt.Sprintf(`

		data "ibm_cos_backup_policy" "policy" {
			bucket_name          = "%s"
			policy_id = "%s"
		}`, name, policy_id)
}
