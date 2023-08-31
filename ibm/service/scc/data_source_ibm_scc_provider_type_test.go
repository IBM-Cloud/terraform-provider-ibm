// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSccProviderTypeDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccProviderTypeDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type.scc_provider_type", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type.scc_provider_type", "provider_type_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type.scc_provider_type", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type.scc_provider_type", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type.scc_provider_type", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type.scc_provider_type", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type.scc_provider_type", "s2s_enabled"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type.scc_provider_type", "instance_limit"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type.scc_provider_type", "mode"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type.scc_provider_type", "data_type"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type.scc_provider_type", "icon"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_provider_type.scc_provider_type", "attributes.%"),
				),
			},
		},
	})
}

func testAccCheckIbmSccProviderTypeDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_scc_provider_type" "scc_provider_type_instance" {
			provider_type_id = "provider_type_id"
			X-Correlation-ID = "X-Correlation-ID"
			X-Request-ID = "X-Request-ID"
		}
	`)
}
