package appid_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMAppIDApplicationsDataSource_basic(t *testing.T) {
	appName1 := fmt.Sprintf("tf_testacc_app_1_%d", acctest.RandIntRange(10, 100))
	appName2 := fmt.Sprintf("tf_testacc_app_2_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAppIDApplicationsDataSourceConfig(acc.AppIDTenantID, appName1, appName2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_appid_applications.apps", "applications.#", "2"),
					resource.TestCheckResourceAttr("data.ibm_appid_applications.apps", "applications.0.name", appName1),
					resource.TestCheckResourceAttr("data.ibm_appid_applications.apps", "applications.1.name", appName2),
				),
			},
		},
	})
}

func testAccCheckIBMAppIDApplicationsDataSourceConfig(tenantID, appName1, appName2 string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_application" "app1" {
			tenant_id = "%s"
			name = "%s"  
			type = "singlepageapp"
		}

		resource "ibm_appid_application" "app2" {
			tenant_id = ibm_appid_application.app1.tenant_id
			name = "%s"
		}

		data "ibm_appid_applications" "apps" {
			tenant_id = ibm_appid_application.app1.tenant_id	

			depends_on = [
				ibm_appid_application.app1,
				ibm_appid_application.app2,
			]
		}
	`, tenantID, appName1, appName2)
}
