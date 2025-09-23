package appid_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"

	appid "github.com/IBM/appid-management-go-sdk/appidmanagementv4"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMAppIDIDPCustom_basic(t *testing.T) {
	publicKey := `-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAzb19EC2vJfTLaJTs3/4F\ndmoHnpYHJo4Q5SJYJK2YfclwRJc49zs1juoNGvXsUOsEi58PHarot3aAUpzBk8g9\n1RdDoovQDKBhMbT7BXP291qp5WQsvrv5W6xPoTbNONYPmAWTN75e3AvvvQElgv9N\n4BBkXZ962bf/OM1Ccm786laop9fC03D7vmUUypISPMZ61O6aA3dRI2JSvHh+VL4s\nEtXkZvLR7DvvWl4sl4oA5EvpYqw5/qbXTp4bnllfiQuCuwgYz/MH1mQA4qGWEVTN\nE4z3b0jsHNHVAzsPfB3Bnok/Zvgtxc3cjVlm3el+bie9O3vW1jFQf1JCke/qusj7\neQIDAQAB\n-----END PUBLIC KEY-----\n`

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAppIDIDPCustomDestroy,
		Steps: []resource.TestStep{
			{
				Config: setupAppIDCustomIDPConfig(acc.AppIDTenantID, publicKey),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_appid_idp_custom.idp", "tenant_id", acc.AppIDTenantID),
					resource.TestCheckResourceAttr("ibm_appid_idp_custom.idp", "is_active", "true"),
					resource.TestCheckResourceAttr("ibm_appid_idp_custom.idp", "public_key", strings.Replace(publicKey, "\\n", "\n", -1)),
				),
			},
		},
	})
}

func setupAppIDCustomIDPConfig(tenantID string, publicKey string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_idp_custom" "idp" {
			tenant_id = "%s"
			is_active = true
			public_key = "%s"
		}	
	`, tenantID, publicKey)
}

func testAccCheckIBMAppIDIDPCustomDestroy(s *terraform.State) error {
	appIDClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).AppIDAPI()

	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_appid_idp_custom" {
			continue
		}

		tenantID := rs.Primary.ID

		cfg, _, err := appIDClient.GetCustomIDP(&appid.GetCustomIDPOptions{
			TenantID: &tenantID,
		})

		if err != nil || (cfg.IsActive != nil && *cfg.IsActive != false) {
			return fmt.Errorf("[ERROR] Error checking if AppID custom IDP configuration (%s) has been destroyed", rs.Primary.ID)
		}
	}

	return nil
}
