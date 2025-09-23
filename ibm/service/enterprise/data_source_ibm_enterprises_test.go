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

func TestAccIbmEnterprisesDataSourceBasic(t *testing.T) {
	//enterpriseSourceAccountID := fmt.Sprintf("source_account_id_%d", acctest.RandIntRange(10, 100))
	enterpriseName := fmt.Sprintf("enterprise_name_%d", acctest.RandIntRange(10, 100))
	//enterprisePrimaryContactIamID := fmt.Sprintf("primary_contact_iam_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmEnterprisesDataSourceConfigBasic(enterpriseName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_enterprises.enterprises", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_enterprises.enterprises", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_enterprises.enterprises", "enterprises.#"),
					resource.TestCheckResourceAttr("data.ibm_enterprises.enterprises", "enterprises.0.name", enterpriseName),
				),
			},
		},
	})
}

func TestAccIbmEnterprisesDataSourceAllArgs(t *testing.T) {
	//enterpriseSourceAccountID := fmt.Sprintf("source_account_id_%d", acctest.RandIntRange(10, 100))
	enterpriseName := fmt.Sprintf("enterprise_name_%d", acctest.RandIntRange(10, 100))
	//enterprisePrimaryContactIamID := fmt.Sprintf("primary_contact_iam_id_%d", acctest.RandIntRange(10, 100))
	enterpriseDomain := fmt.Sprintf("enterprise_domain_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmEnterprisesDataSourceConfig(enterpriseName, enterpriseDomain),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_enterprises.enterprises", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_enterprises.enterprises", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_enterprises.enterprises", "enterprises.#"),
					resource.TestCheckResourceAttrSet("data.ibm_enterprises.enterprises", "enterprises.0.url"),
					resource.TestCheckResourceAttrSet("data.ibm_enterprises.enterprises", "enterprises.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_enterprises.enterprises", "enterprises.0.enterprise_account_id"),
					resource.TestCheckResourceAttrSet("data.ibm_enterprises.enterprises", "enterprises.0.crn"),
					resource.TestCheckResourceAttr("data.ibm_enterprises.enterprises", "enterprises.0.name", enterpriseName),
					resource.TestCheckResourceAttr("data.ibm_enterprises.enterprises", "enterprises.0.domain", enterpriseDomain),
					resource.TestCheckResourceAttrSet("data.ibm_enterprises.enterprises", "enterprises.0.state"),
					resource.TestCheckResourceAttrSet("data.ibm_enterprises.enterprises", "enterprises.0.primary_contact_iam_id"),
					resource.TestCheckResourceAttrSet("data.ibm_enterprises.enterprises", "enterprises.0.primary_contact_email"),
					resource.TestCheckResourceAttrSet("data.ibm_enterprises.enterprises", "enterprises.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_enterprises.enterprises", "enterprises.0.created_by"),
				),
			},
		},
	})
}

func testAccCheckIbmEnterprisesDataSourceConfigBasic(enterpriseName string) string {
	return fmt.Sprintf(`
		data "ibm_iam_users" "current_account_users"{
		}
		resource "ibm_enterprise" "enterprise" {
			source_account_id = data.ibm_iam_users.current_account_users.users[0].account_id
			name = "%s"
			primary_contact_iam_id = data.ibm_iam_users.current_account_users.users[0].iam_id
		}

		data "ibm_enterprises" "enterprises" {
			depends_on  = [ibm_enterprise.enterprise]
			name = ibm_enterprise.enterprise.name
		}
	`, enterpriseName)
}

func testAccCheckIbmEnterprisesDataSourceConfig(enterpriseName string, enterpriseDomain string) string {
	return fmt.Sprintf(`
		data "ibm_iam_users" "current_account_users"{
		}
		resource "ibm_enterprise" "enterprise" {
			source_account_id = data.ibm_iam_users.current_account_users.users[0].account_id
			name = "%s"
			primary_contact_iam_id = data.ibm_iam_users.current_account_users.users[0].iam_id
			domain = "%s"
		}
		data "ibm_enterprises" "enterprises" {
			name = ibm_enterprise.enterprise.name
		}
	`, enterpriseName, enterpriseDomain)
}
