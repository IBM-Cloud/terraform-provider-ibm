package ibm

import (
	"fmt"
	appid "github.com/IBM/appid-management-go-sdk/appidmanagementv4"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"strings"
	"testing"
)

func TestAccIBMAppIDApplication_basic(t *testing.T) {
	appName := fmt.Sprintf("tf_testacc_app_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMAppIDApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMApplicationConfig(appIDTenantID, appName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_appid_application.test_app", "tenant_id", appIDTenantID),
					resource.TestCheckResourceAttr("ibm_appid_application.test_app", "name", appName),
					resource.TestCheckResourceAttr("ibm_appid_application.test_app", "type", "regularwebapp"),
					resource.TestCheckResourceAttrSet("ibm_appid_application.test_app", "client_id"),
				),
			},
		},
	})
}

func testAccCheckIBMApplicationConfig(tenantID string, name string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_application" "test_app" {
			tenant_id = "%s"
			name = "%s"
		}	
	`, tenantID, name)
}

func testAccCheckIBMAppIDApplicationDestroy(s *terraform.State) error {
	appIDClient, err := testAccProvider.Meta().(ClientSession).AppIDAPI()

	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_appid_application" {
			continue
		}

		id := rs.Primary.ID
		idParts := strings.Split(id, "/")

		tenantID := idParts[0]
		clientID := idParts[1]

		_, _, err := appIDClient.GetApplication(&appid.GetApplicationOptions{
			TenantID: &tenantID,
			ClientID: &clientID,
		})

		if err == nil {
			return fmt.Errorf("error checking if AppID application (%s) has been destroyed", rs.Primary.ID)
		}
	}

	return nil
}
