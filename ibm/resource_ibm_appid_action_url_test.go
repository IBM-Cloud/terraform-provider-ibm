package ibm

import (
	"fmt"
	appid "github.com/IBM/appid-management-go-sdk/appidmanagementv4"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"strings"
	"testing"
)

func TestAccIBMAppIDActionURL_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMAppIDActionURLDestroy,
		Steps: []resource.TestStep{
			{
				Config: setupAppIDActionURLConfig(appIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_appid_action_url.url", "tenant_id", appIDTenantID),
					resource.TestCheckResourceAttr("ibm_appid_action_url.url", "action", "on_reset_password"),
					resource.TestCheckResourceAttr("ibm_appid_action_url.url", "url", "https://www.example.com/psw-reset"),
				),
			},
		},
	})
}

func setupAppIDActionURLConfig(tenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_action_url" "url" {
			tenant_id = "%s"
			action = "on_reset_password"
			url = "https://www.example.com/psw-reset"
		}	
	`, tenantID)
}

func testAccCheckIBMAppIDActionURLDestroy(s *terraform.State) error {
	appIDClient, err := testAccProvider.Meta().(ClientSession).AppIDAPI()

	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_appid_action_url" {
			continue
		}

		id := rs.Primary.ID
		idParts := strings.Split(id, "/")

		tenantID := idParts[0]
		action := idParts[1]

		cfg, _, err := appIDClient.GetCloudDirectoryActionURL(&appid.GetCloudDirectoryActionURLOptions{
			TenantID: &tenantID,
			Action:   &action,
		})

		// when action URL is deleted it is returned as an empty string e.g. `{ actionUrl: "" }`
		if err != nil || (cfg.ActionURL != nil && *cfg.ActionURL != "") {
			return fmt.Errorf("error checking if AppID action URL (%s) has been destroyed", rs.Primary.ID)
		}
	}

	return nil
}
