/* IBM Confidential
*  Object Code Only Source Materials
*  5747-SM3
*  (c) Copyright IBM Corp. 2017,2021
*
*  The source code for this program is not published or otherwise divested
*  of its trade secrets, irrespective of what has been deposited with the
*  U.S. Copyright Office. */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
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
