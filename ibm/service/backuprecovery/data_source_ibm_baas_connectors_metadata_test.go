// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.94.0-fa797aec-20240814-142622
 */

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBaasConnectorsMetadataDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasConnectorsMetadataDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_baas_connectors_metadata.baas_connectors_metadata_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_connectors_metadata.baas_connectors_metadata_instance", "tenant_id"),
				),
			},
		},
	})
}

func testAccCheckIbmBaasConnectorsMetadataDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_baas_connectors_metadata" "baas_connectors_metadata_instance" {
		}

		data "ibm_baas_connectors_metadata" "baas_connectors_metadata_instance" {
			tenant_id = 8
		}
	`)
}
