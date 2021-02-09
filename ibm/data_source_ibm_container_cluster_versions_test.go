/**
 * © Copyright IBM Corporation 2020. All Rights Reserved.
 *
 * Licensed under the Mozilla Public License, version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at https://mozilla.org/MPL/2.0/
 */

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
