package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMOrgQuotaDataSource_basic(t *testing.T) {

	name := "qIBM"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMOrgQuotaDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_org_quota.testacc_ds_org_quota", "name", name),
				),
			},
		},
	})
}

func testAccCheckIBMOrgQuotaDataSourceConfig(name string) string {
	return fmt.Sprintf(`
	
data "ibm_org_quota" "testacc_ds_org_quota" {
    name = "%s"
}`, name)

}
