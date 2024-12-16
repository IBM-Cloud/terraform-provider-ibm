package appid_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAppIDMFADataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setupIBMMFADataSourceConfig(acc.AppIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_appid_mfa.mf", "tenant_id", acc.AppIDTenantID),
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
