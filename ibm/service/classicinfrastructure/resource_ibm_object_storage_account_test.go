// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package classicinfrastructure_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMObjectStorageAccount_Basic(t *testing.T) {
	var accountName string

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMObjectStorageAccountDestroy,
		Steps: []resource.TestStep{
			{
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
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMObjectStorageAccountDestroy,
		Steps: []resource.TestStep{
			{
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
			return fmt.Errorf("[ERROR] Not  found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("[ERROR] No Record ID is set")
		}

		*accountName = rs.Primary.ID

		return nil
	}
}

func testAccCheckIBMObjectStorageAccountAttributes(accountName *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if *accountName == "" {
			return fmt.Errorf("[ERROR] No object storage account name")
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
