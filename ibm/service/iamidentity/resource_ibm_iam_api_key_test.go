// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMIamAPIKeyBasic(t *testing.T) {
	var conf iamidentityv1.APIKey
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIamAPIKeyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamAPIKeyConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIamAPIKeyExists("ibm_iam_api_key.iam_api_key_instance", conf),
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIamAPIKeyConfigBasic(nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key_instance", "name", nameUpdate),
				),
			},
		},
	})
}

func TestAccIBMIamAPIKeyAllArgs(t *testing.T) {
	var conf iamidentityv1.APIKey
	name := fmt.Sprintf("name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("description_%d", acctest.RandIntRange(10, 100))
	storeValue := "false"
	nameUpdate := fmt.Sprintf("name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("description_%d", acctest.RandIntRange(10, 100))
	storeValueUpdate := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMIamAPIKeyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIamAPIKeyConfig(name, description, storeValue),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMIamAPIKeyExists("ibm_iam_api_key.iam_api_key_instance", conf),
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key_instance", "description", description),
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key_instance", "store_value", storeValue),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMIamAPIKeyConfig(nameUpdate, descriptionUpdate, storeValueUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key_instance", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key_instance", "store_value", storeValueUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_iam_api_key.iam_api_key_instance",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"apikey",      // is only present on the initial create response
					"entity_lock", // not part of read response
					"store_value", // not part of read response
				},
			},
		},
	})
}

func testAccCheckIBMIamAPIKeyConfigBasic(name string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_api_key" "iam_api_key_instance" {
			name = "%s"
		}
	`, name)
}

func testAccCheckIBMIamAPIKeyConfig(name string, description string, storeValue string) string {
	return fmt.Sprintf(`
		resource "ibm_iam_api_key" "iam_api_key_instance" {
			name = "%s"
			description = "%s"
			store_value = %s
		}
	`, name, description, storeValue)
}

func testAccCheckIBMIamAPIKeyExists(n string, obj iamidentityv1.APIKey) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		iamIdentityClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMIdentityV1API()
		if err != nil {
			return err
		}

		getAPIKeyOptions := &iamidentityv1.GetAPIKeyOptions{}

		getAPIKeyOptions.SetID(rs.Primary.ID)

		apiKey, _, err := iamIdentityClient.GetAPIKey(getAPIKeyOptions)
		if err != nil {
			return err
		}

		obj = *apiKey
		return nil
	}
}

func testAccCheckIBMIamAPIKeyDestroy(s *terraform.State) error {
	iamIdentityClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_api_key" {
			continue
		}

		getAPIKeyOptions := &iamidentityv1.GetAPIKeyOptions{}

		getAPIKeyOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := iamIdentityClient.GetAPIKey(getAPIKeyOptions)

		if err == nil {
			return fmt.Errorf("iam_api_key still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for iam_api_key (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
