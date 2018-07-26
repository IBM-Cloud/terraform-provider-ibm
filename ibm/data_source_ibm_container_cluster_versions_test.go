package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMContainerClusterVersionsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMContainerClusterVersionsDataSource(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_container_cluster_versions.versions", "valid_kube_versions.0"),
				),
			},
		},
	})
}

func TestAccIBMContainerClusterVersionsDataSource_WithoutOptionalFields(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMContainerClusterVersionsDataSourceWithoutOptionalFields,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_container_cluster_versions.versions", "valid_kube_versions.0"),
				),
			},
		},
	})
}

func testAccCheckIBMContainerClusterVersionsDataSource() string {
	return fmt.Sprintf(`
data "ibm_org" "testacc_ds_org" {
    org = "%s"
}
data "ibm_space" "testacc_ds_space" {
    org = "%s"
    space = "%s"
}
data "ibm_account" "testacc_acc" {
    org_guid = "${data.ibm_org.testacc_ds_org.id}"
}
data "ibm_container_cluster_versions" "versions" {
	org_guid = "${data.ibm_org.testacc_ds_org.id}"
    space_guid = "${data.ibm_space.testacc_ds_space.id}"
    account_guid = "${data.ibm_account.testacc_acc.id}"
    region = "%s"
}
`, cfOrganization, cfOrganization, cfSpace, csRegion)
}

const testAccCheckIBMContainerClusterVersionsDataSourceWithoutOptionalFields = `
data "ibm_container_cluster_versions" "versions" {
}
`
