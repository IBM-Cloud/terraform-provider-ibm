// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSourceRegistrationDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSourceRegistrationDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_source_registration.source_registration_instance", "id"),
				),
			},
		},
	})
}

func testAccCheckIbmSourceRegistrationDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_source_registration" "source_registration_instance" {
			ids = [ 1 ]
			tenantIds = [ "tenantIds" ]
			includeTenants = true
			includeSourceCredentials = true
			encryptionKey = "encryptionKey"
			useCachedData = true
		}
	`)
}
