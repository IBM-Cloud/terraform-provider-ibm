package appid_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccAppIDIDPGoogleDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setupIBMAppIDGoogleIDPDataSourceConfig(acc.AppIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_appid_idp_google.gg", "tenant_id", acc.AppIDTenantID),
					resource.TestCheckResourceAttr("data.ibm_appid_idp_google.gg", "config.0.application_id", "test_id"),
					resource.TestCheckResourceAttr("data.ibm_appid_idp_google.gg", "config.0.application_secret", "test_secret"),
					resource.TestCheckResourceAttrSet("data.ibm_appid_idp_google.gg", "redirect_url"),
				),
			},
		},
	})
}

func setupIBMAppIDGoogleIDPDataSourceConfig(tenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_idp_google" "gg" {
			tenant_id = "%s"
			is_active = true
			
			config {
				application_id 		= "test_id"
				application_secret 	= "test_secret"
			}
		}
		data "ibm_appid_idp_google" "gg" {
			tenant_id = ibm_appid_idp_google.gg.tenant_id
			depends_on = [
				ibm_appid_idp_google.gg
			]
		}
	`, tenantID)
}
