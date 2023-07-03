// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSccPostureListCredentialsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccPostureListCredentialsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_credentials.list_credentials", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_credentials.list_credentials", "first.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_credentials.list_credentials", "last.#"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_credentials.list_credentials", "credentials.#"),
				),
			},
		},
	})
}

func testAccCheckIBMSccPostureListCredentialsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_scc_posture_credentials" "list_credentials" {
		}
	`)
}
