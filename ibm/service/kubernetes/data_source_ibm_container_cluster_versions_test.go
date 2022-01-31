// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package kubernetes_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMContainerClusterVersionsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMContainerClusterVersionsDataSource(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_container_cluster_versions.versions", "valid_kube_versions.0"),
					resource.TestCheckResourceAttrSet("data.ibm_container_cluster_versions.versions", "valid_openshift_versions.0"),
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
`, acc.CsRegion)
}
