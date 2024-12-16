package appid_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	appid "github.com/IBM/appid-management-go-sdk/appidmanagementv4"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMAppIDAuditStatus_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAppIDAuditStatusDestroy,
		Steps: []resource.TestStep{
			{
				Config: setupAppIDAuditStatusConfig(acc.AppIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_appid_audit_status.status", "tenant_id", acc.AppIDTenantID),
					resource.TestCheckResourceAttr("ibm_appid_audit_status.status", "is_active", "true"),
				),
			},
		},
	})
}

func testAccCheckIBMAppIDAuditStatusDestroy(s *terraform.State) error {
	appIDClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).AppIDAPI()

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
			return fmt.Errorf("[ERROR] Error checking if AppID audit status (%s) has been destroyed", rs.Primary.ID)
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
