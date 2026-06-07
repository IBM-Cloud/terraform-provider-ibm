package kms_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMKMSCryptoUnits_basic(t *testing.T) {
	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	sigKeyFilepath := "./test_signature_key"
	sigKeyOwner := "admin"
	sigKeyPassphrase := "test-passphrase"
	masterKeyName := "TESTKEY"
	keyShareFilepath1 := "./test_keyshare1.key"
	keyShareFilepath2 := "./test_keyshare2.key"
	keyShareToken1 := "token1"
	keyShareToken2 := "token2"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsCryptoUnitsBasicConfig(
					instanceName,
					sigKeyFilepath,
					sigKeyOwner,
					sigKeyPassphrase,
					masterKeyName,
					keyShareFilepath1,
					keyShareToken1,
					keyShareFilepath2,
					keyShareToken2,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_kms_cryptounits.test", "instance_id"),
					resource.TestCheckResourceAttr("ibm_kms_cryptounits.test", "signature_key.#", "1"),
					resource.TestCheckResourceAttr("ibm_kms_cryptounits.test", "master_key.#", "1"),
					resource.TestCheckResourceAttrSet("ibm_kms_cryptounits.test", "cryptounits.%"),
				),
			},
		},
	})
}

func TestAccIBMKMSCryptoUnits_withRegion(t *testing.T) {
	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	region := "us-south"
	sigKeyFilepath := "./test_signature_key.p8"
	sigKeyOwner := "admin"
	sigKeyPassphrase := "test-passphrase"
	masterKeyName := "TESTKEY"
	keyShareFilepath1 := "./test_keyshare1.key"
	keyShareFilepath2 := "./test_keyshare2.key"
	keyShareToken1 := "token1"
	keyShareToken2 := "token2"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsCryptoUnitsWithRegionConfig(
					instanceName,
					region,
					sigKeyFilepath,
					sigKeyOwner,
					sigKeyPassphrase,
					masterKeyName,
					keyShareFilepath1,
					keyShareToken1,
					keyShareFilepath2,
					keyShareToken2,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_kms_cryptounits.test", "instance_id"),
					resource.TestCheckResourceAttr("ibm_kms_cryptounits.test", "region", region),
					resource.TestCheckResourceAttr("ibm_kms_cryptounits.test", "signature_key.#", "1"),
					resource.TestCheckResourceAttr("ibm_kms_cryptounits.test", "master_key.#", "1"),
					resource.TestCheckResourceAttrSet("ibm_kms_cryptounits.test", "cryptounits.%"),
				),
			},
		},
	})
}

func TestAccIBMKMSCryptoUnits_withPrivateEndpoint(t *testing.T) {
	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	sigKeyFilepath := "./test_signature_key.p8"
	sigKeyOwner := "admin"
	sigKeyPassphrase := "test-passphrase"
	masterKeyName := "TESTKEY"
	keyShareFilepath1 := "./test_keyshare1.key"
	keyShareFilepath2 := "./test_keyshare2.key"
	keyShareToken1 := "token1"
	keyShareToken2 := "token2"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsCryptoUnitsWithPrivateEndpointConfig(
					instanceName,
					sigKeyFilepath,
					sigKeyOwner,
					sigKeyPassphrase,
					masterKeyName,
					keyShareFilepath1,
					keyShareToken1,
					keyShareFilepath2,
					keyShareToken2,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_kms_cryptounits.test", "instance_id"),
					resource.TestCheckResourceAttr("ibm_kms_cryptounits.test", "use_private_endpoint", "true"),
					resource.TestCheckResourceAttr("ibm_kms_cryptounits.test", "signature_key.#", "1"),
					resource.TestCheckResourceAttr("ibm_kms_cryptounits.test", "master_key.#", "1"),
					resource.TestCheckResourceAttrSet("ibm_kms_cryptounits.test", "cryptounits.%"),
				),
			},
		},
	})
}

func TestAccIBMKMSCryptoUnits_multipleKeyShares(t *testing.T) {
	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	sigKeyFilepath := "./test_signature_key.p8"
	sigKeyOwner := "admin"
	sigKeyPassphrase := "test-passphrase"
	masterKeyName := "TESTKEY"
	keyShareFilepath1 := "./test_keyshare1.key"
	keyShareFilepath2 := "./test_keyshare2.key"
	keyShareFilepath3 := "./test_keyshare3.key"
	keyShareToken1 := "token1"
	keyShareToken2 := "token2"
	keyShareToken3 := "token3"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsCryptoUnitsMultipleKeySharesConfig(
					instanceName,
					sigKeyFilepath,
					sigKeyOwner,
					sigKeyPassphrase,
					masterKeyName,
					keyShareFilepath1,
					keyShareToken1,
					keyShareFilepath2,
					keyShareToken2,
					keyShareFilepath3,
					keyShareToken3,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_kms_cryptounits.test", "instance_id"),
					resource.TestCheckResourceAttr("ibm_kms_cryptounits.test", "signature_key.#", "1"),
					resource.TestCheckResourceAttr("ibm_kms_cryptounits.test", "master_key.#", "1"),
					resource.TestCheckResourceAttrSet("ibm_kms_cryptounits.test", "cryptounits.%"),
				),
			},
		},
	})
}

func testAccCheckIBMKmsCryptoUnitsBasicConfig(instanceName, sigKeyFilepath, sigKeyOwner, sigKeyPassphrase, masterKeyName, keyShareFilepath1, keyShareToken1, keyShareFilepath2, keyShareToken2 string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kms_instance" {
		name     = "%s"
		service  = "kms"
		plan     = "tiered-pricing"
		location = "us-south"
	}

	resource "ibm_kms_cryptounits" "test" {
		instance_id = ibm_resource_instance.kms_instance.guid

		signature_key {
			filepath   = "%s"
			owner      = "%s"
			passphrase = "%s"
			exists     = false
		}

		master_key {
			keyname = "%s"
			exists  = false

			keysharefile {
				filepath = "%s"
				token    = "%s"
			}

			keysharefile {
				filepath = "%s"
				token    = "%s"
			}
		}
	}
`, addPrefixToResourceName(instanceName), sigKeyFilepath, sigKeyOwner, sigKeyPassphrase, masterKeyName, keyShareFilepath1, keyShareToken1, keyShareFilepath2, keyShareToken2)
}

func testAccCheckIBMKmsCryptoUnitsWithRegionConfig(instanceName, region, sigKeyFilepath, sigKeyOwner, sigKeyPassphrase, masterKeyName, keyShareFilepath1, keyShareToken1, keyShareFilepath2, keyShareToken2 string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kms_instance" {
		name     = "%s"
		service  = "kms"
		plan     = "tiered-pricing"
		location = "%s"
	}

	resource "ibm_kms_cryptounits" "test" {
		instance_id = ibm_resource_instance.kms_instance.guid
		region      = "%s"

		signature_key {
			filepath   = "%s"
			owner      = "%s"
			passphrase = "%s"
			exists     = false
		}

		master_key {
			keyname = "%s"
			exists  = false

			keysharefile {
				filepath = "%s"
				token    = "%s"
			}

			keysharefile {
				filepath = "%s"
				token    = "%s"
			}
		}
	}
`, addPrefixToResourceName(instanceName), region, region, sigKeyFilepath, sigKeyOwner, sigKeyPassphrase, masterKeyName, keyShareFilepath1, keyShareToken1, keyShareFilepath2, keyShareToken2)
}

func testAccCheckIBMKmsCryptoUnitsWithPrivateEndpointConfig(instanceName, sigKeyFilepath, sigKeyOwner, sigKeyPassphrase, masterKeyName, keyShareFilepath1, keyShareToken1, keyShareFilepath2, keyShareToken2 string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kms_instance" {
		name     = "%s"
		service  = "kms"
		plan     = "tiered-pricing"
		location = "us-south"
	}

	resource "ibm_kms_cryptounits" "test" {
		instance_id          = ibm_resource_instance.kms_instance.guid
		use_private_endpoint = true

		signature_key {
			filepath   = "%s"
			owner      = "%s"
			passphrase = "%s"
			exists     = false
		}

		master_key {
			keyname = "%s"
			exists  = false

			keysharefile {
				filepath = "%s"
				token    = "%s"
			}

			keysharefile {
				filepath = "%s"
				token    = "%s"
			}
		}
	}
`, addPrefixToResourceName(instanceName), sigKeyFilepath, sigKeyOwner, sigKeyPassphrase, masterKeyName, keyShareFilepath1, keyShareToken1, keyShareFilepath2, keyShareToken2)
}

func testAccCheckIBMKmsCryptoUnitsMultipleKeySharesConfig(instanceName, sigKeyFilepath, sigKeyOwner, sigKeyPassphrase, masterKeyName, keyShareFilepath1, keyShareToken1, keyShareFilepath2, keyShareToken2, keyShareFilepath3, keyShareToken3 string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kms_instance" {
		name     = "%s"
		service  = "kms"
		plan     = "tiered-pricing"
		location = "us-south"
	}

	resource "ibm_kms_cryptounits" "test" {
		instance_id = ibm_resource_instance.kms_instance.guid

		signature_key {
			filepath   = "%s"
			owner      = "%s"
			passphrase = "%s"
			exists     = false
		}

		master_key {
			keyname = "%s"
			exists  = false

			keysharefile {
				filepath = "%s"
				token    = "%s"
			}

			keysharefile {
				filepath = "%s"
				token    = "%s"
			}

			keysharefile {
				filepath = "%s"
				token    = "%s"
			}
		}
	}
`, addPrefixToResourceName(instanceName), sigKeyFilepath, sigKeyOwner, sigKeyPassphrase, masterKeyName, keyShareFilepath1, keyShareToken1, keyShareFilepath2, keyShareToken2, keyShareFilepath3, keyShareToken3)
}

// Made with Bob
