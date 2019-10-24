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
					resource.TestCheckResourceAttrSet("data.ibm_container_cluster_versions.versions", "valid_openshift_versions.0"),
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
data "ibm_container_cluster_versions" "versions" {
    region = "%s"
}
`, csRegion)
}

const testAccCheckIBMContainerClusterVersionsDataSourceWithoutOptionalFields = `
data "ibm_container_cluster_versions" "versions" {
}
`
