package ibm

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/softlayer/softlayer-go/services"
	"strconv"
	"testing"
)

//Testcase for SWIFT object storage account
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

//Testcase for S3 object storage account
func TestAccIBMObjectStorageS3Account_Basic(t *testing.T) {
	var accountName string

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMObjectStorageAccountDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMObjectStorageS3AccountConfig_basic,
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
	sess := testAccProvider.Meta().(ClientSession).SoftLayerSession()
	storageService := services.GetNetworkStorageService(sess)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_object_storage_account" {
			continue
		}

		storageID, _ := strconv.Atoi(rs.Primary.ID)

		// Try to find the key
		_, err := storageService.Id(storageID).GetObject()

		if err == nil {
			return fmt.Errorf("Object Storage Account %d still exists", storageID)
		}
	}
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

var testAccCheckIBMObjectStorageS3AccountConfig_basic = `
resource "ibm_object_storage_account" "testacc_foobar" {
	accountType = "S3"
}`

var testAccCheckIBMObjectStorageAccountWithTag = `
resource "ibm_object_storage_account" "testacc_foobar" {
	tags = ["one", "two"]
}`

var testAccCheckIBMObjectStorageAccountWithUpdatedTag = `
resource "ibm_object_storage_account" "testacc_foobar" {
	tags = ["one", "two", "three"]
}`
