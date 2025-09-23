// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamidentity_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
)

func TestAccIbmIamApiKeyBasic(t *testing.T) {
	var conf iamidentityv1.APIKey
	name := fmt.Sprintf("name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmIamApiKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIamApiKeyConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIamApiKeyExists("ibm_iam_api_key.iam_api_key", conf),
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key", "name", name),
				),
			},
			{
				Config: testAccCheckIbmIamApiKeyConfigBasic(nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key", "name", nameUpdate),
				),
			},
		},
	})
}

func TestAccIbmIamApiKeyAllArgs(t *testing.T) {
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
		CheckDestroy: testAccCheckIbmIamApiKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmIamApiKeyConfig(name, description, storeValue),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmIamApiKeyExists("ibm_iam_api_key.iam_api_key", conf),
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key", "name", name),
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key", "description", description),
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key", "store_value", storeValue),
				),
			},
			{
				Config: testAccCheckIbmIamApiKeyConfig(nameUpdate, descriptionUpdate, storeValueUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_iam_api_key.iam_api_key", "store_value", storeValueUpdate),
				),
			},
			{
				ResourceName:      "ibm_iam_api_key.iam_api_key",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmIamApiKeyConfigBasic(name string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_api_key" "iam_api_key" {
			name = "%s"
		}
	`, name)
}

func testAccCheckIbmIamApiKeyConfig(name string, description string, storeValue string) string {
	return fmt.Sprintf(`

		resource "ibm_iam_api_key" "iam_api_key" {
			name = "%s"
			description = "%s"
			store_value = %s
		}
	`, name, description, storeValue)
}

func testAccCheckIbmIamApiKeyExists(n string, obj iamidentityv1.APIKey) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		iamIdentityClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMIdentityV1API()
		if err != nil {
			return err
		}

		getApiKeyOptions := &iamidentityv1.GetAPIKeyOptions{}

		getApiKeyOptions.SetID(rs.Primary.ID)

		apiKey, _, err := iamIdentityClient.GetAPIKey(getApiKeyOptions)
		if err != nil {
			return err
		}

		obj = *apiKey
		return nil
	}
}

func testAccCheckIbmIamApiKeyDestroy(s *terraform.State) error {
	iamIdentityClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMIdentityV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_iam_api_key" {
			continue
		}

		getApiKeyOptions := &iamidentityv1.GetAPIKeyOptions{}

		getApiKeyOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := iamIdentityClient.GetAPIKey(getApiKeyOptions)

		if err == nil {
			return fmt.Errorf("iam_api_key still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error checking for iam_api_key (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
