// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMObjectStorageAccount_Basic(t *testing.T) {
	var accountName string

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMObjectStorageAccountDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMObjectStorageAccountConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMObjectStorageAccountExists("ibm_object_storage_account.testacc_foobar", &accountName),
					testAccCheckIBMObjectStorageAccountAttributes(&accountName),
				),
			},
		},
	})
}

func TestAccIBMObjectStorageAccountWithTag(t *testing.T) {
	var accountName string

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMObjectStorageAccountDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMObjectStorageAccountWithTag,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMObjectStorageAccountExists("ibm_object_storage_account.testacc_foobar", &accountName),
					testAccCheckIBMObjectStorageAccountAttributes(&accountName),
					resource.TestCheckResourceAttr(
						"ibm_object_storage_account.testacc_foobar", "tags.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMObjectStorageAccountWithUpdatedTag,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMObjectStorageAccountExists("ibm_object_storage_account.testacc_foobar", &accountName),
					testAccCheckIBMObjectStorageAccountAttributes(&accountName),
					resource.TestCheckResourceAttr(
						"ibm_object_storage_account.testacc_foobar", "tags.#", "3"),
				),
			},
		},
	})
}

func testAccCheckIBMObjectStorageAccountDestroy(s *terraform.State) error {
	return nil
}

func testAccCheckIBMObjectStorageAccountExists(n string, accountName *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		*accountName = rs.Primary.ID

		return nil
	}
}

func testAccCheckIBMObjectStorageAccountAttributes(accountName *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if *accountName == "" {
			return fmt.Errorf("No object storage account name")
		}

		return nil
	}
}

var testAccCheckIBMObjectStorageAccountConfig_basic = `
resource "ibm_object_storage_account" "testacc_foobar" {
}`

var testAccCheckIBMObjectStorageAccountWithTag = `
resource "ibm_object_storage_account" "testacc_foobar" {
	tags = ["one", "two"]
}`

var testAccCheckIBMObjectStorageAccountWithUpdatedTag = `
resource "ibm_object_storage_account" "testacc_foobar" {
	tags = ["one", "two", "three"]
}`
