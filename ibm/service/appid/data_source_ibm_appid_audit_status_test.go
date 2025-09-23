package appid_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMAppIDAuditStatusDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setupAppIDAuditStatusDataSourceConfig(acc.AppIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_appid_audit_status.status", "tenant_id", acc.AppIDTenantID),
					resource.TestCheckResourceAttr("data.ibm_appid_audit_status.status", "is_active", "true"),
				),
			},
		},
	})
}

func setupAppIDAuditStatusDataSourceConfig(tenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_audit_status" "status" {
			tenant_id = "%s"
			is_active = true
		}

		data "ibm_appid_audit_status" "status" {
			tenant_id = ibm_appid_audit_status.status.tenant_id

			depends_on = [
				ibm_appid_audit_status.status
			]
		}
	`, tenantID)
}
