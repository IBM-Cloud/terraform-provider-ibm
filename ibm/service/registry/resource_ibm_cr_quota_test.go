// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package registry_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMCrQuotaBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestAccProtoV6ProviderFactories(),
		CheckDestroy:             testAccCheckCrQuotaReset,
		Steps: []resource.TestStep{
			{
				Config: testAccCrQuotaConfig("storage_megabytes = 1024"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_cr_quota.test", "id"),
					resource.TestCheckResourceAttr("ibm_cr_quota.test", "storage_megabytes", "1024"),
					resource.TestCheckNoResourceAttr("ibm_cr_quota.test", "traffic_megabytes"),
					resource.TestCheckResourceAttrSet("ibm_cr_quota.test", "storage_usage_bytes"),
				),
			},
			{
				// Update storage and start managing traffic.
				Config: testAccCrQuotaConfig("storage_megabytes = 2048\n traffic_megabytes = 10240"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cr_quota.test", "storage_megabytes", "2048"),
					resource.TestCheckResourceAttr("ibm_cr_quota.test", "traffic_megabytes", "10240"),
				),
			},
		},
	})
}

func TestAccIBMCrQuotaRequiresOneQuota(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestAccProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config:      testAccCrQuotaConfig(""),
				ExpectError: regexp.MustCompile("Set at least one of storage_megabytes or traffic_megabytes"),
			},
		},
	})
}

func testAccCrQuotaConfig(body string) string {
	return fmt.Sprintf(`
		resource "ibm_cr_quota" "test" {
			%s
		}
	`, body)
}

// testAccCheckCrQuotaReset checks that destroy put the managed quotas back to
// the service defaults (reset_on_destroy defaults to true).
func testAccCheckCrQuotaReset(s *terraform.State) error {
	client, err := acc.TestAccProvider.Meta().(conns.ClientSession).ContainerRegistryV1()
	if err != nil {
		return err
	}
	quota, _, err := client.GetQuota(client.NewGetQuotaOptions())
	if err != nil {
		return err
	}
	const mb = 1024 * 1024
	if quota.Limit == nil || quota.Limit.StorageBytes == nil || *quota.Limit.StorageBytes != 500*mb {
		return fmt.Errorf("storage quota was not reset to 500 MB: %+v", quota.Limit)
	}
	if quota.Limit.TrafficBytes == nil || *quota.Limit.TrafficBytes != 5120*mb {
		return fmt.Errorf("traffic quota was not reset to 5120 MB: %+v", quota.Limit)
	}
	return nil
}
