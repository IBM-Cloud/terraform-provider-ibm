// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package enterprise_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

/* To run this test case ensure the IC_API_KEY belongs to an enterprise" */
func TestAccIbmAccountsDataSourceBasic(t *testing.T) {
	//accountParent := fmt.Sprintf("parent_%d", acctest.RandIntRange(10, 100))
	accountName := fmt.Sprintf("name_%d", acctest.RandIntRange(10, 100))
	//accountOwnerIamID := fmt.Sprintf("owner_iam_id_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmAccountsDataSourceConfigBasic(accountName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_enterprise_accounts.accounts", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_enterprise_accounts.accounts", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_enterprise_accounts.accounts", "accounts.#"),
					resource.TestCheckResourceAttr("data.ibm_enterprise_accounts.accounts", "accounts.0.name", accountName),
				),
			},
		},
	})
}

func testAccCheckIbmAccountsDataSourceConfigBasic(accountName string) string {

	return fmt.Sprintf(`
		data "ibm_enterprises" "enterprises_instance" {
		}
		resource "ibm_enterprise_account" "enterprise_account" {
			parent = data.ibm_enterprises.enterprises_instance.enterprises[0].crn
			name = "%s"
			owner_iam_id = data.ibm_enterprises.enterprises_instance.enterprises[0].primary_contact_iam_id
		}

		data "ibm_enterprise_accounts" "accounts" {
			depends_on  = [ibm_enterprise_account.enterprise_account]
			name = ibm_enterprise_account.enterprise_account.name
		}
	`, accountName)
}
