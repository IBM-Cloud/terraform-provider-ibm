package ibm

import (
	"fmt"
	appid "github.com/IBM/appid-management-go-sdk/appidmanagementv4"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

func TestAccIBMAppIDAuditStatus_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMAppIDAuditStatusDestroy,
		Steps: []resource.TestStep{
			{
				Config: setupAppIDAuditStatusConfig(appIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_appid_audit_status.status", "tenant_id", appIDTenantID),
					resource.TestCheckResourceAttr("ibm_appid_audit_status.status", "is_active", "true"),
				),
			},
		},
	})
}

func testAccCheckIBMAppIDAuditStatusDestroy(s *terraform.State) error {
	appIDClient, err := testAccProvider.Meta().(ClientSession).AppIDAPI()

	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_appid_audit_status" {
			continue
		}

		tenantID := rs.Primary.ID

		cfg, _, err := appIDClient.GetAuditStatus(&appid.GetAuditStatusOptions{
			TenantID: &tenantID,
		})

		// default for audit status is `false`
		if err != nil || (cfg.IsActive != nil && *cfg.IsActive != false) {
			return fmt.Errorf("error checking if AppID audit status (%s) has been destroyed", rs.Primary.ID)
		}
	}

	return nil
}

func setupAppIDAuditStatusConfig(tenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_audit_status" "status" {
			tenant_id = "%s"
			is_active = true
		}
	`, tenantID)
}
