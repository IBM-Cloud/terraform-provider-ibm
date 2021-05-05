// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0
package ibm

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"

	"github.com/IBM/platform-services-go-sdk/enterprisemanagementv1"
)

/* To run this test case ensure the IC_API_KEY belongs to an enterprise" */
func TestAccIbmEnterpriseAccountBasic(t *testing.T) {
	var conf enterprisemanagementv1.Account
	//parent := fmt.Sprintf("parent_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf-gen-account-name_%d", acctest.RandIntRange(10, 100))
	//ownerIamID := fmt.Sprintf("owner_iam_id_%d", acctest.RandIntRange(10, 100))
	//parentUpdate := fmt.Sprintf("parent_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckEnterprise(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmEnterpriseAccountConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmEnterpriseAccountExists("ibm_enterprise_account.enterprise_account", conf),
					resource.TestCheckResourceAttrSet("ibm_enterprise_account.enterprise_account", "parent"),
					resource.TestCheckResourceAttr("ibm_enterprise_account.enterprise_account", "name", name),
					resource.TestCheckResourceAttrSet("ibm_enterprise_account.enterprise_account", "owner_iam_id"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmEnterpriseAccountConfigUpdateBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_enterprise_account.enterprise_account", "parent"),
					resource.TestCheckResourceAttrSet("ibm_enterprise_account.enterprise_account", "name"),
					resource.TestCheckResourceAttrSet("ibm_enterprise_account.enterprise_account", "owner_iam_id"),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_enterprise_account.enterprise_account",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

/* To run this test case ensure the IC_API_KEY belongs to an enterprise.
ACCOUNT_TO_BE_IMPORTED should invite enterprise and grant relevant iam policies before running this test case" */
func TestAccIbmEnterpriseImportAccountBasic(t *testing.T) {
	var conf enterprisemanagementv1.Account
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckEnterpriseAccountImport(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmAccountsDataSourceConfigImportBasic(account_to_be_imported),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmEnterpriseAccountExists("ibm_enterprise_account.enterprise_account_import", conf),
					resource.TestCheckResourceAttrSet("ibm_enterprise_account.enterprise_account_import", "parent"),
					resource.TestCheckResourceAttr("ibm_enterprise_account.enterprise_account_import", "account_id", account_to_be_imported),
					resource.TestCheckResourceAttrSet("ibm_enterprise_account.enterprise_account_import", "owner_iam_id"),
				),
			},
		},
	})
}

func testAccCheckIbmEnterpriseAccountConfigBasic(name string) string {
	return fmt.Sprintf(`
		data "ibm_enterprises" "enterprises_instance" {
		}
		resource "ibm_enterprise_account" "enterprise_account" {
			parent = data.ibm_enterprises.enterprises_instance.enterprises[0].crn
			name = "%s"
			owner_iam_id = data.ibm_enterprises.enterprises_instance.enterprises[0].primary_contact_iam_id
		}
	`, name)
}

func testAccCheckIbmEnterpriseAccountConfigUpdateBasic(name string) string {
	return fmt.Sprintf(`
		data "ibm_enterprise_account_groups" "account_groups_instance" {
		}
		resource "ibm_enterprise_account" "enterprise_account" {
			parent = data.ibm_enterprise_account_groups.account_groups_instance.account_groups[0].crn
			name = "%s"
			owner_iam_id = data.ibm_enterprise_account_groups.account_groups_instance.account_groups[0].primary_contact_iam_id
		}
	`, name)
}

func testAccCheckIbmAccountsDataSourceConfigImportBasic(accountToBeImported string) string {

	return fmt.Sprintf(`
		data "ibm_enterprises" "enterprises_instance" {
		}
		resource "ibm_enterprise_account" "enterprise_account_import" {
			enterprise_id = data.ibm_enterprises.enterprises_instance.enterprises[0].id
			account_id = "%s"
			parent = data.ibm_enterprises.enterprises_instance.enterprises[0].crn
		}
	`, accountToBeImported)
}

func testAccCheckIbmEnterpriseAccountExists(n string, obj enterprisemanagementv1.Account) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		enterpriseManagementClient, err := testAccProvider.Meta().(ClientSession).EnterpriseManagementV1()
		if err != nil {
			return err
		}

		getAccountOptions := &enterprisemanagementv1.GetAccountOptions{}

		getAccountOptions.SetAccountID(rs.Primary.ID)

		account, _, err := enterpriseManagementClient.GetAccount(getAccountOptions)
		if err != nil {
			return err
		}

		obj = *account
		return nil
	}
}
