package appid_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	appid "github.com/IBM/appid-management-go-sdk/appidmanagementv4"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMAppIDApplication_basic(t *testing.T) {
	appName := fmt.Sprintf("tf_testacc_app_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAppIDApplicationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAppIDApplicationConfig(acc.AppIDTenantID, appName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_appid_application.test_app", "tenant_id", acc.AppIDTenantID),
					resource.TestCheckResourceAttr("ibm_appid_application.test_app", "name", appName),
					resource.TestCheckResourceAttr("ibm_appid_application.test_app", "type", "regularwebapp"),
					resource.TestCheckResourceAttrSet("ibm_appid_application.test_app", "client_id"),
				),
			},
		},
	})
}

func testAccCheckIBMAppIDApplicationConfig(tenantID string, name string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_application" "test_app" {
			tenant_id = "%s"
			name = "%s"
		}	
	`, tenantID, name)
}

func testAccCheckIBMAppIDApplicationDestroy(s *terraform.State) error {
	appIDClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).AppIDAPI()

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
			return fmt.Errorf("[ERROR] Error checking if AppID application (%s) has been destroyed", rs.Primary.ID)
		}
	}

	return nil
}
