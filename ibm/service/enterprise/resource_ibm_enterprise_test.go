// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package enterprise_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/IBM/platform-services-go-sdk/enterprisemanagementv1"
)

func TestAccIbmEnterpriseBasic(t *testing.T) {
	var conf enterprisemanagementv1.Enterprise
	//sourceAccountID := fmt.Sprintf("source_account_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("name_%d", acctest.RandIntRange(10, 100))
	//primaryContactIamID := fmt.Sprintf("primary_contact_iam_id_%d", acctest.RandIntRange(10, 100))
	//sourceAccountIDUpdate := fmt.Sprintf("source_account_id_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("name_%d", acctest.RandIntRange(10, 100))
	//primaryContactIamIDUpdate := fmt.Sprintf("primary_contact_iam_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmEnterpriseConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmEnterpriseExists("ibm_enterprise.enterprise", conf),
					resource.TestCheckResourceAttrSet("ibm_enterprise.enterprise", "source_account_id"),
					resource.TestCheckResourceAttr("ibm_enterprise.enterprise", "name", name),
					resource.TestCheckResourceAttrSet("ibm_enterprise.enterprise", "primary_contact_iam_id"),
				),
			},
			{
				Config: testAccCheckIbmEnterpriseConfigBasic(nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_enterprise.enterprise", "source_account_id"),
					resource.TestCheckResourceAttr("ibm_enterprise.enterprise", "name", nameUpdate),
					resource.TestCheckResourceAttrSet("ibm_enterprise.enterprise", "primary_contact_iam_id"),
				),
			},
		},
	})
}

func TestAccIbmEnterpriseAllArgs(t *testing.T) {
	var conf enterprisemanagementv1.Enterprise
	//sourceAccountID := fmt.Sprintf("source_account_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_test_generated_name_%d", acctest.RandIntRange(10, 100))
	//primaryContactIamID := fmt.Sprintf("primary_contact_iam_id_%d", acctest.RandIntRange(10, 100))
	domain := fmt.Sprintf("tf_test_generated_domain_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_updated_test_generated_name_%d", acctest.RandIntRange(10, 100))
	//primaryContactIamIDUpdate := fmt.Sprintf("primary_contact_iam_id_%d", acctest.RandIntRange(10, 100))
	domainUpdate := fmt.Sprintf("tf_updated_test_domain_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmEnterpriseConfig(name, domain),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmEnterpriseExists("ibm_enterprise.enterprise", conf),
					resource.TestCheckResourceAttrSet("ibm_enterprise.enterprise", "source_account_id"),
					resource.TestCheckResourceAttr("ibm_enterprise.enterprise", "name", name),
					resource.TestCheckResourceAttrSet("ibm_enterprise.enterprise", "primary_contact_iam_id"),
					resource.TestCheckResourceAttr("ibm_enterprise.enterprise", "domain", domain),
				),
			},
			{
				Config: testAccCheckIbmEnterpriseConfig(nameUpdate, domainUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_enterprise.enterprise", "source_account_id"),
					resource.TestCheckResourceAttr("ibm_enterprise.enterprise", "name", nameUpdate),
					resource.TestCheckResourceAttrSet("ibm_enterprise.enterprise", "primary_contact_iam_id"),
					resource.TestCheckResourceAttr("ibm_enterprise.enterprise", "domain", domainUpdate),
				),
			},
			{
				ResourceName:      "ibm_enterprise.enterprise",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmEnterpriseConfigBasic(name string) string {
	return fmt.Sprintf(`
		data "ibm_iam_users" "current_account_users"{
		}
		resource "ibm_enterprise" "enterprise" {
			source_account_id = data.ibm_iam_users.current_account_users.users[0].account_id
			name = "%s"
			primary_contact_iam_id = data.ibm_iam_users.current_account_users.users[0].iam_id
			
		}
	`, name)
}

func testAccCheckIbmEnterpriseConfig(name string, domain string) string {
	return fmt.Sprintf(`
		data "ibm_iam_users" "current_account_users"{
		}
		resource "ibm_enterprise" "enterprise" {
			source_account_id = data.ibm_iam_users.current_account_users.users[0].account_id
			name = "%s"
			primary_contact_iam_id = data.ibm_iam_users.current_account_users.users[0].iam_id
			domain = "%s"
		}
	`, name, domain)
}

func testAccCheckIbmEnterpriseExists(n string, obj enterprisemanagementv1.Enterprise) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		enterpriseManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).EnterpriseManagementV1()
		if err != nil {
			return err
		}

		getEnterpriseOptions := &enterprisemanagementv1.GetEnterpriseOptions{}

		getEnterpriseOptions.SetEnterpriseID(rs.Primary.ID)

		enterprise, _, err := enterpriseManagementClient.GetEnterprise(getEnterpriseOptions)
		if err != nil {
			return err
		}

		obj = *enterprise
		return nil
	}
}
