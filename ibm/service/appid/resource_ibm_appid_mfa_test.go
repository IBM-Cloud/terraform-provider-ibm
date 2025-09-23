package appid_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAppIDMFA_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setupIBMMFAConfig(acc.AppIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_appid_mfa.mf", "tenant_id", acc.AppIDTenantID),
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
