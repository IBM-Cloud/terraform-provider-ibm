// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package dnsservices_test

import (
	"fmt"
	"os"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPrivateDNSCustomResolverCacheClearActionsDataSource_basic(t *testing.T) {
	instanceID := os.Getenv("DNS_SVCS_INSTANCE_ID")
	resolverID := os.Getenv("DNS_SVCS_CUSTOM_RESOLVER_ID")

	if instanceID == "" || resolverID == "" {
		t.Skip("Skipping TestAccIBMPrivateDNSCustomResolverCacheClearActionsDataSource_basic: DNS_SVCS_INSTANCE_ID and DNS_SVCS_CUSTOM_RESOLVER_ID must be set")
	}

	node := "data.ibm_dns_custom_resolver_cache_clear_actions.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSCustomResolverCacheClearActionsDataSourceConfig(instanceID, resolverID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(node, "instance_id", instanceID),
					resource.TestCheckResourceAttr(node, "resolver_id", resolverID),
					resource.TestCheckResourceAttrSet(node, "total_count"),
				),
			},
		},
	})
}

func testAccCheckIBMPrivateDNSCustomResolverCacheClearActionsDataSourceConfig(instanceID, resolverID string) string {
	return fmt.Sprintf(`
data "ibm_dns_custom_resolver_cache_clear_actions" "test" {
	instance_id = "%s"
	resolver_id = "%s"
}
`, instanceID, resolverID)
}
