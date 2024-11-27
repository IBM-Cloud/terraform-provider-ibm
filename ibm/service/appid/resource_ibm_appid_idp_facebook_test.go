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

func TestAccAppIDIDPFacebook_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAppIDIDPFacebookDestroy,
		Steps: []resource.TestStep{
			{
				Config: setupIBMAppIDFacebookIDPConfig(acc.AppIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_appid_idp_facebook.fb", "tenant_id", acc.AppIDTenantID),
					resource.TestCheckResourceAttr("ibm_appid_idp_facebook.fb", "config.0.application_id", "test_id"),
					resource.TestCheckResourceAttr("ibm_appid_idp_facebook.fb", "config.0.application_secret", "test_secret"),
					resource.TestCheckResourceAttrSet("ibm_appid_idp_facebook.fb", "redirect_url"),
				),
			},
		},
	})
}

func setupIBMAppIDFacebookIDPConfig(tenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_idp_facebook" "fb" {
			tenant_id = "%s"
			is_active = true
			
			config {
				application_id 		= "test_id"
				application_secret 	= "test_secret"
			}
		}
	`, tenantID)
}

func testAccCheckIBMAppIDIDPFacebookDestroy(s *terraform.State) error {
	appIDClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).AppIDAPI()

	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_appid_idp_facebook" {
			continue
		}

		tenantID := rs.Primary.ID

		config, _, err := appIDClient.GetFacebookIDP(&appid.GetFacebookIDPOptions{
			TenantID: &tenantID,
		})

		if err != nil {
			return fmt.Errorf("[ERROR] Error checking if AppID Facebook IDP configuration was reset: %s", err)
		}

		if config == nil || (config.IsActive != nil && *config.IsActive != false) {
			return fmt.Errorf("[ERROR] Error checking if AppID Facebook IDP configuration was reset")
		}
	}

	return nil
}
