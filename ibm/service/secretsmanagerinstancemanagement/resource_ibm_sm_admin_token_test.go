// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package secretsmanagerinstancemanagement_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSmAdminTokenBasic(t *testing.T) {
	resourceName := "ibm_sm_admin_token.sm_admin_token"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmAdminTokenConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "token"),
				),
			},
		},
	})
}

func TestAccIbmSmAdminTokenResponseWrapping(t *testing.T) {
	resourceName := "ibm_sm_admin_token.sm_admin_token"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSmAdminTokenConfigResponseWrapping(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttr(resourceName, "response_wrapping", "true"),
					resource.TestCheckResourceAttrSet(resourceName, "wrapped_token"),
				),
			},
		},
	})
}

func testAccCheckIbmSmAdminTokenConfigResponseWrapping() string {
	return fmt.Sprintf(`

		resource "ibm_sm_admin_token" "sm_admin_token" {
			instance_id       = "%s"
			response_wrapping = true
		}
	`, acc.SecretsManagerDedicatedInstanceID)
}

func testAccCheckIbmSmAdminTokenConfigBasic() string {
	return fmt.Sprintf(`

		resource "ibm_sm_admin_token" "sm_admin_token" {
			instance_id   = "%s"
		}
	`, acc.SecretsManagerDedicatedInstanceID)
}
