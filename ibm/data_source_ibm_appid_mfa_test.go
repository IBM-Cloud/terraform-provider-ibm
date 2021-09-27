package ibm

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccAppIDMFADataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setupIBMMFADataSourceConfig(appIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_appid_mfa.mf", "tenant_id", appIDTenantID),
					resource.TestCheckResourceAttr("data.ibm_appid_mfa.mf", "is_active", "true"),
				),
			},
		},
	})
}

func setupIBMMFADataSourceConfig(tenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_mfa" "mf" {
			tenant_id = "%s"
			is_active = true
		}
		data "ibm_appid_mfa" "mf" {
			tenant_id = ibm_appid_mfa.mf.tenant_id
			depends_on = [
				ibm_appid_mfa.mf
			]
		}
	`, tenantID)
}
