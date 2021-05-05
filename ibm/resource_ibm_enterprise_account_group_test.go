// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/platform-services-go-sdk/enterprisemanagementv1"
)

/* To run this test case ensure the IC_API_KEY belongs to an enterprise" */
func TestAccIbmEnterpriseAccountGroupBasic(t *testing.T) {
	var conf enterprisemanagementv1.AccountGroup

	name := fmt.Sprintf("tf-gen-name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf-gen-updated-name_%d", acctest.RandIntRange(10, 100))
	//primaryContactIamIDUpdate := fmt.Sprintf("primary_contact_iam_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckEnterprise(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmEnterpriseAccountGroupConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmEnterpriseAccountGroupExists("ibm_enterprise_account_group.enterprise_account_group", conf),
					resource.TestCheckResourceAttrSet("ibm_enterprise_account_group.enterprise_account_group", "parent"),
					resource.TestCheckResourceAttr("ibm_enterprise_account_group.enterprise_account_group", "name", name),
					resource.TestCheckResourceAttrSet("ibm_enterprise_account_group.enterprise_account_group", "primary_contact_iam_id"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmEnterpriseAccountGroupConfigBasic(nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_enterprise_account_group.enterprise_account_group", "parent"),
					resource.TestCheckResourceAttr("ibm_enterprise_account_group.enterprise_account_group", "name", nameUpdate),
					resource.TestCheckResourceAttrSet("ibm_enterprise_account_group.enterprise_account_group", "primary_contact_iam_id"),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_enterprise_account_group.enterprise_account_group",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmEnterpriseAccountGroupConfigBasic(name string) string {
	return fmt.Sprintf(`
		data "ibm_enterprises" "enterprises_instance" {
		}
		resource "ibm_enterprise_account_group" "enterprise_account_group" {
			parent = data.ibm_enterprises.enterprises_instance.enterprises[0].crn
			name = "%s"
			primary_contact_iam_id = data.ibm_enterprises.enterprises_instance.enterprises[0].primary_contact_iam_id
		}
	`, name)
}

func testAccCheckIbmEnterpriseAccountGroupExists(n string, obj enterprisemanagementv1.AccountGroup) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		enterpriseManagementClient, err := testAccProvider.Meta().(ClientSession).EnterpriseManagementV1()
		if err != nil {
			return err
		}

		getAccountGroupOptions := &enterprisemanagementv1.GetAccountGroupOptions{}

		getAccountGroupOptions.SetAccountGroupID(rs.Primary.ID)

		accountGroup, _, err := enterpriseManagementClient.GetAccountGroup(getAccountGroupOptions)
		if err != nil {
			return err
		}

		obj = *accountGroup
		return nil
	}
}
