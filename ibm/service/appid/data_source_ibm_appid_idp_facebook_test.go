package appid_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAppIDIDPFacebookDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setupIBMAppIDFacebookIDPDataSourceConfig(acc.AppIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_appid_idp_facebook.fb", "tenant_id", acc.AppIDTenantID),
					resource.TestCheckResourceAttr("data.ibm_appid_idp_facebook.fb", "config.0.application_id", "test_id"),
					resource.TestCheckResourceAttr("data.ibm_appid_idp_facebook.fb", "config.0.application_secret", "test_secret"),
					resource.TestCheckResourceAttrSet("data.ibm_appid_idp_facebook.fb", "redirect_url"),
				),
			},
		},
	})
}

func setupIBMAppIDFacebookIDPDataSourceConfig(tenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_idp_facebook" "fb" {
			tenant_id = "%s"
			is_active = true
			
			config {
				application_id 		= "test_id"
				application_secret 	= "test_secret"
			}
		}

		data "ibm_appid_idp_facebook" "fb" {
			tenant_id = ibm_appid_idp_facebook.fb.tenant_id

			depends_on = [
				ibm_appid_idp_facebook.fb
			]
		}
	`, tenantID)
}
