package ibm

import (
	//"errors"
	"fmt"
	"testing"

	//"regexp"

	//v1 "github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	//"github.com/hashicorp/terraform/terraform"
)

func TestAccCisSettings_Basic(t *testing.T) {
	// multiple instances of this config would conflict but we only use it once
	t.Parallel()

	cisId := cis_crn

	rnd := acctest.RandString(10)
	name := "ibm_cis_domain_settings." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		//CheckDestroy: testAccCheckCisSettingsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisSettingsConfigBasic(cis_domain, rnd, cisId),
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

func testAccCheckCisSettingsConfigBasic(cis_domain string, id string, cisId string) string {
	return testAccIBMCisDomainConfig_basic("test", cisId, cis_domain) + fmt.Sprintf(`
resource "ibm_cis_domain_settings" "%[2]s" {
  cis_id = "%[3]s"	
  domain_id = "${ibm_cis_domain.test.id}"
  "waf" = "on"
  "ssl" = "full"	
  "min_tls_version" = "1.2"
}`, cis_domain, id, cisId)
}
