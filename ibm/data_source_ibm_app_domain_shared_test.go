package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMAppDomainSharedDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMAppDomainSharedDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(
						"data.ibm_app_domain_shared.testacc_domain", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMAppDomainSharedDataSourceConfig() string {
	return fmt.Sprintf(`
	
		data "ibm_app_domain_shared" "testacc_domain" {
			name = "mybluemix.net"
		}`)

}
