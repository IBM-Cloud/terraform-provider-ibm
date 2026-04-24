// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIamServiceidGroupBasic(t *testing.T) {
	var conf iamidentityv1.ServiceIDGroup
	accountID := fmt.Sprintf("tf_account_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	accountIDUpdate := fmt.Sprintf("tf_account_id_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIamServiceidGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamServiceidGroupConfigBasic(accountID, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIamServiceidGroupExists("ibm_iam_serviceid_group.iam_serviceid_group_instance", conf),
					resource.TestCheckResourceAttr("ibm_iam_serviceid_group.iam_serviceid_group_instance", "account_id", accountID),
					resource.TestCheckResourceAttr("ibm_iam_serviceid_group.iam_serviceid_group_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIamServiceidGroupConfigBasic(accountIDUpdate, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_serviceid_group.iam_serviceid_group_instance", "account_id", accountIDUpdate),
					resource.TestCheckResourceAttr("ibm_iam_serviceid_group.iam_serviceid_group_instance", "name", nameUpdate),
				),
			},
		},
	})
}

func TestAccIBMIamServiceidGroupAllArgs(t *testing.T) {
	var conf iamidentityv1.ServiceIDGroup
	accountID := fmt.Sprintf("tf_account_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	accountIDUpdate := fmt.Sprintf("tf_account_id_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIamServiceidGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamServiceidGroupConfig(accountID, name, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIamServiceidGroupExists("ibm_iam_serviceid_group.iam_serviceid_group_instance", conf),
					resource.TestCheckResourceAttr("ibm_iam_serviceid_group.iam_serviceid_group_instance", "account_id", accountID),
					resource.TestCheckResourceAttr("ibm_iam_serviceid_group.iam_serviceid_group_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_serviceid_group.iam_serviceid_group_instance", "description", description),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIamServiceidGroupConfig(accountIDUpdate, nameUpdate, descriptionUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_serviceid_group.iam_serviceid_group_instance", "account_id", accountIDUpdate),
					resource.TestCheckResourceAttr("ibm_iam_serviceid_group.iam_serviceid_group_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_iam_serviceid_group.iam_serviceid_group_instance", "description", descriptionUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_iam_serviceid_group.iam_serviceid_group_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMIamServiceidGroupConfigBasic(accountID string, name string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_serviceid_group" "iam_serviceid_group_instance" {
			account_id = "%s"
			name = "%s"
		}
	`, accountID, name)
}

func testAccCheckIBMIamServiceidGroupConfig(accountID string, name string, description string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_serviceid_group" "iam_serviceid_group_instance" {
			account_id = "%s"
			name = "%s"
			description = "%s"
		}
	`, accountID, name, description)
}

func testAccCheckIBMIamServiceidGroupExists(n string, obj iamidentityv1.ServiceIDGroup) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		iamIdentityClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IamIdentityV1()
		if err != nil {
			return err
		}

		getServiceIDGroupOptions := &iamidentityv1.GetServiceIDGroupOptions{}

		getServiceIDGroupOptions.SetID(rs.Primary.ID)

		serviceIDGroup, _, err := iamIdentityClient.GetServiceIDGroup(getServiceIDGroupOptions)
		if err != nil {
			return err
		}

		obj = *serviceIDGroup
		return nil
	}
}

func testAccCheckIBMIamServiceidGroupDestroy(s *terraform.State) error {
	iamIdentityClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IamIdentityV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_serviceid_group" {
			continue
		}

		getServiceIDGroupOptions := &iamidentityv1.GetServiceIDGroupOptions{}

		getServiceIDGroupOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := iamIdentityClient.GetServiceIDGroup(getServiceIDGroupOptions)

		if err == nil {
			return fmt.Errorf("iam_serviceid_group still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for iam_serviceid_group (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
