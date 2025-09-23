// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package enterprise_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

/* To run this test case ensure the IC_API_KEY belongs to an enterprise" */
func TestAccIbmAccountGroupsDataSourceBasic(t *testing.T) {
	//accountGroupParent := fmt.Sprintf("parent_%d", acctest.RandIntRange(10, 100))
	accountGroupName := fmt.Sprintf("tf_gen_name_%d", acctest.RandIntRange(10, 100))
	//accountGroupPrimaryContactIamID := fmt.Sprintf("primary_contact_iam_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmAccountGroupsDataSourceConfigBasic(accountGroupName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_enterprise_account_groups.account_groups", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_enterprise_account_groups.account_groups", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_enterprise_account_groups.account_groups", "account_groups.#"),
					resource.TestCheckResourceAttr("data.ibm_enterprise_account_groups.account_groups", "account_groups.0.name", accountGroupName),
				),
			},
		},
	})
}

func testAccCheckIbmAccountGroupsDataSourceConfigBasic(accountGroupName string) string {
	return fmt.Sprintf(`
		data "ibm_enterprises" "enterprises_instance" {
		}
		resource "ibm_enterprise_account_group" "enterprise_account_group" {
			parent = data.ibm_enterprises.enterprises_instance.enterprises[0].crn
			name = "%s"
			primary_contact_iam_id = data.ibm_enterprises.enterprises_instance.enterprises[0].primary_contact_iam_id
		}
		data "ibm_enterprise_account_groups" "account_groups" {
			depends_on =[ibm_enterprise_account_group.enterprise_account_group]
			name = ibm_enterprise_account_group.enterprise_account_group.name
		}
	`, accountGroupName)
}
