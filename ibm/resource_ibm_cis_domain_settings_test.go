package ibm

import (
	//"errors"
	"fmt"
	"testing"

	//"regexp"

	//v1 "github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
	//"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	//"github.com/hashicorp/terraform/terraform"
)

func TestAccCisSettings_Basic(t *testing.T) {
	// multiple instances of this config would conflict but we only use it once
	t.Parallel()

	//rnd := acctest.RandString(10)
	name := "ibm_cis_domain_settings." + "test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		// Remove check destroy as this occurs after the CIS instance is deleted and fails with an auth error
		//CheckDestroy: testAccCheckCisSettingsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisSettingsConfigBasic("test", cis_domain),
				Check: resource.ComposeTestCheckFunc(
					// dont check that specified values are set, this will be evident by lack of plan diff
					// some values will get empty values
					resource.TestCheckResourceAttr(name, "waf", "on"),
					resource.TestCheckResourceAttr(name, "ssl", "full"),
					resource.TestCheckResourceAttr(name, "min_tls_version", "1.2"),
					//resource.TestCheckResourceAttr(name, "tls_1_3_setting", "off"),
				),
			},
		},
	})
}

func testAccCheckCisSettingsConfigBasic(id string, cis_domain string) string {
	return testAccIBMCisDomainConfig_basic("test", cis_domain) + fmt.Sprintf(`
resource "ibm_cis_domain_settings" "%[1]s" {
  cis_id = "${ibm_cis.instance.id}"
  domain_id = "${ibm_cis_domain.%[1]s.id}"
  "waf" = "on"
  "ssl" = "full"	
  "min_tls_version" = "1.2"
}`, id, cis_domain)
}
