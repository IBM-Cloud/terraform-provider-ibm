package ibm

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccAppIDMFA_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setupIBMMFAConfig(appIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_appid_mfa.mf", "tenant_id", appIDTenantID),
					resource.TestCheckResourceAttr("ibm_appid_mfa.mf", "is_active", "true"),
				),
			},
		},
	})
}

func setupIBMMFAConfig(tenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_mfa" "mf" {
			tenant_id = "%s"
			is_active = true
		}
	`, tenantID)
}
