package ibm

import (
	"fmt"
	appid "github.com/IBM/appid-management-go-sdk/appidmanagementv4"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestAccIBMAppIDCloudDirectoryUser_basic(t *testing.T) {
	userName := fmt.Sprintf("tf_testacc_user_%d", acctest.RandIntRange(10, 100))
	email := fmt.Sprintf("%s@mail.com", userName)
	lockedUntil := time.Now().Add(time.Hour*2).UnixNano() / int64(time.Millisecond)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMAppIDCloudDirectoryUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAppIDCloudDirectoryUserConfig(appIDTenantID, userName, email, lockedUntil),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_appid_cloud_directory_user.user", "active", "false"),
					resource.TestCheckResourceAttr("ibm_appid_cloud_directory_user.user", "locked_until", strconv.Itoa(int(lockedUntil))),
					resource.TestCheckResourceAttrSet("ibm_appid_cloud_directory_user.user", "user_id"),
					resource.TestCheckResourceAttr("ibm_appid_cloud_directory_user.user", "display_name", "Test TF User"),
					resource.TestCheckResourceAttr("ibm_appid_cloud_directory_user.user", "status", "PENDING"),
					resource.TestCheckResourceAttr("ibm_appid_cloud_directory_user.user", "email.#", "1"),
					resource.TestCheckResourceAttr("ibm_appid_cloud_directory_user.user", "email.0.value", email),
					resource.TestCheckResourceAttr("ibm_appid_cloud_directory_user.user", "email.0.primary", "true"),
				),
			},
		},
	})
}

func testAccCheckIBMAppIDCloudDirectoryUserConfig(tenantID, userName, email string, lockedUntil int64) string {
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
	`, tenantID, userName, email, lockedUntil)
}

func testAccCheckIBMAppIDCloudDirectoryUserDestroy(s *terraform.State) error {
	appIDClient, err := testAccProvider.Meta().(ClientSession).AppIDAPI()

	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_appid_cloud_directory_user" {
			continue
		}

		id := rs.Primary.ID
		idParts := strings.Split(id, "/")

		tenantID := idParts[0]
		userID := idParts[1]

		_, _, err := appIDClient.GetCloudDirectoryUser(&appid.GetCloudDirectoryUserOptions{
			TenantID: &tenantID,
			UserID:   &userID,
		})

		if err == nil {
			return fmt.Errorf("error checking if AppID Cloud Directory user (%s) has been destroyed", rs.Primary.ID)
		}
	}

	return nil
}
