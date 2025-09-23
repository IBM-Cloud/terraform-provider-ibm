package appid_test

import (
	"fmt"
	"strconv"
	"testing"

	"time"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMAppIDCloudDirectoryUserDataSource_basic(t *testing.T) {
	userName := fmt.Sprintf("tf_testacc_user_%d", acctest.RandIntRange(10, 100))
	email := fmt.Sprintf("%s@mail.com", userName)
	lockedUntil := time.Now().Add(time.Hour*2).UnixNano() / int64(time.Millisecond)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAppIDCloudDirectoryUserDataSourceConfig(acc.AppIDTenantID, userName, email, lockedUntil),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_appid_cloud_directory_user.user", "active", "false"),
					resource.TestCheckResourceAttr("data.ibm_appid_cloud_directory_user.user", "locked_until", strconv.Itoa(int(lockedUntil))),
					resource.TestCheckResourceAttrSet("data.ibm_appid_cloud_directory_user.user", "user_id"),
					resource.TestCheckResourceAttrSet("data.ibm_appid_cloud_directory_user.user", "subject"),
					resource.TestCheckResourceAttr("data.ibm_appid_cloud_directory_user.user", "display_name", "Test TF User"),
					resource.TestCheckResourceAttr("data.ibm_appid_cloud_directory_user.user", "status", "PENDING"),
					resource.TestCheckResourceAttr("data.ibm_appid_cloud_directory_user.user", "email.#", "1"),
					resource.TestCheckResourceAttr("data.ibm_appid_cloud_directory_user.user", "email.0.value", email),
					resource.TestCheckResourceAttr("data.ibm_appid_cloud_directory_user.user", "email.0.primary", "true"),
				),
			},
		},
	})
}

func testAccCheckIBMAppIDCloudDirectoryUserDataSourceConfig(tenantID, userName, email string, lockedUntil int64) string {
	return fmt.Sprintf(`
		resource "ibm_appid_cloud_directory_user" "user" {
			tenant_id = "%s"
			user_name = "%s"
			
			email {
				value = "%s"
				primary = true
			}

			active = false
			locked_until = %d

			password = "P@ssw0rd"
			
			display_name = "Test TF User"
		}

		data "ibm_appid_cloud_directory_user" "user" {
			tenant_id = ibm_appid_cloud_directory_user.user.tenant_id
			user_id = ibm_appid_cloud_directory_user.user.user_id

			depends_on = [
				ibm_appid_cloud_directory_user.user	
			]
		}
	`, tenantID, userName, email, lockedUntil)
}
