package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMAccountDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMAccountDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_org.testacc_ds_org", "org", cfOrganization),
					resource.TestCheckResourceAttrSet(
						"data.ibm_account.testacc_acc", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMAccountDataSourceConfig() string {
	return fmt.Sprintf(`
	
data "ibm_org" "testacc_ds_org" {
    org = "%s"
}

data "ibm_account" "testacc_acc" {
    org_guid = data.ibm_org.testacc_ds_org.id
}`, cfOrganization)

}
