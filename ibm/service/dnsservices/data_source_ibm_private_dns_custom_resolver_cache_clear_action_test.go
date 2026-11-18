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

func TestAccIBMPrivateDNSCustomResolverCacheClearActionDataSource_basic(t *testing.T) {
	instanceID := os.Getenv("DNS_SVCS_INSTANCE_ID")
	resolverID := os.Getenv("DNS_SVCS_CUSTOM_RESOLVER_ID")
	actionID := os.Getenv("DNS_SVCS_CACHE_CLEAR_ACTION_ID")

	if instanceID == "" || resolverID == "" || actionID == "" {
		t.Skip("Skipping TestAccIBMPrivateDNSCustomResolverCacheClearActionDataSource_basic: DNS_SVCS_INSTANCE_ID, DNS_SVCS_CUSTOM_RESOLVER_ID, and DNS_SVCS_CACHE_CLEAR_ACTION_ID must be set")
	}

	node := "data.ibm_dns_custom_resolver_cache_clear_action.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPrivateDNSCustomResolverCacheClearActionDataSourceConfig(instanceID, resolverID, actionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(node, "instance_id", instanceID),
					resource.TestCheckResourceAttr(node, "resolver_id", resolverID),
					resource.TestCheckResourceAttr(node, "action_id", actionID),
					resource.TestCheckResourceAttrSet(node, "status"),
					resource.TestCheckResourceAttrSet(node, "created_on"),
					resource.TestCheckResourceAttrSet(node, "expires_on"),
				),
			},
		},
	})
}

func testAccCheckIBMPrivateDNSCustomResolverCacheClearActionDataSourceConfig(instanceID, resolverID, actionID string) string {
	return fmt.Sprintf(`
data "ibm_dns_custom_resolver_cache_clear_action" "test" {
	instance_id = "%s"
	resolver_id = "%s"
	action_id   = "%s"
}
`, instanceID, resolverID, actionID)
}
