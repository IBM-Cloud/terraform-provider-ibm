package kms_test

// TF_LOG=1 TF_ACC=1 go test -v -run "TestAccIBMKMSCryptoUnits" ./ibm/service/kms/

import (
	"context"
	"fmt"
	"os"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	kpCryptoUnit "github.com/IBM/keyprotect-go-client/dedicated"
)

func TestAccIBMKMSCryptoUnits_basic(t *testing.T) {
	instanceName := os.Getenv("KP_INSTANCE_ID")
	url := os.Getenv("KP_URL")
	sigKeyFilepath := "./test_signature_key"
	sigKeyOwner := "admin"
	sigKeyPassphrase := ""
	masterKeyName := "TESTKEY"
	keyShareFilepath1 := "./test_keyshare1.key"
	keyShareFilepath2 := "./test_keyshare2.key"
	keyShareToken1 := "token1"
	keyShareToken2 := "token2"

	t.Cleanup(func() {
		os.Remove(sigKeyFilepath)
		os.Remove(keyShareFilepath1)
		os.Remove(keyShareFilepath2)
	})

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckKmsCrypto(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMKmsCryptoUnitsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsCryptoUnitsBasicConfig(
					instanceName,
					url,
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
					resource.TestCheckResourceAttrSet("ibm_kms_cryptounits.test", "cryptounits.#"),
				),
			},
		},
	})
}

func TestAccIBMKMSCryptoUnits_withRegion(t *testing.T) {
	instanceName := os.Getenv("KP_INSTANCE_ID")
	url := os.Getenv("KP_URL")
	region := "eu-gb"
	sigKeyFilepath := "./test_signature_key_region.p8"
	sigKeyOwner := "admin"
	sigKeyPassphrase := ""
	masterKeyName := "TESTKEY"
	keyShareFilepath1 := "./test_keyshare_region1.key"
	keyShareFilepath2 := "./test_keyshare_region2.key"
	keyShareToken1 := "token1"
	keyShareToken2 := "token2"

	t.Cleanup(func() {
		os.Remove(sigKeyFilepath)
		os.Remove(keyShareFilepath1)
		os.Remove(keyShareFilepath2)
	})

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckKmsCrypto(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMKmsCryptoUnitsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsCryptoUnitsWithRegionConfig(
					instanceName,
					url,
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
					resource.TestCheckResourceAttrSet("ibm_kms_cryptounits.test", "cryptounits.#"),
				),
			},
		},
	})
}

func TestAccIBMKMSCryptoUnits_multipleKeyShares(t *testing.T) {
	instanceName := os.Getenv("KP_INSTANCE_ID")
	url := os.Getenv("KP_URL")
	sigKeyFilepath := "./test_signature_key_multi.p8"
	sigKeyOwner := "admin"
	sigKeyPassphrase := ""
	masterKeyName := "TESTKEY"
	keyShareFilepath1 := "./test_keyshare_multi1.key"
	keyShareFilepath2 := "./test_keyshare_multi2.key"
	keyShareFilepath3 := "./test_keyshare_multi3.key"
	keyShareToken1 := "token1"
	keyShareToken2 := "token2"
	keyShareToken3 := "token3"

	t.Cleanup(func() {
		os.Remove(sigKeyFilepath)
		os.Remove(keyShareFilepath1)
		os.Remove(keyShareFilepath2)
		os.Remove(keyShareFilepath3)
	})

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckKmsCrypto(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMKmsCryptoUnitsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKmsCryptoUnitsMultipleKeySharesConfig(
					instanceName,
					url,
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
					resource.TestCheckResourceAttrSet("ibm_kms_cryptounits.test", "cryptounits.#"),
				),
			},
		},
	})
}

func testAccCheckIBMKmsCryptoUnitsBasicConfig(instanceName, url, sigKeyFilepath, sigKeyOwner, sigKeyPassphrase, masterKeyName, keyShareFilepath1, keyShareToken1, keyShareFilepath2, keyShareToken2 string) string {
	return fmt.Sprintf(`
	resource "ibm_kms_cryptounits" "test" {
		instance_id = "%s"
		url         = "%s"
		should_zeroize = true

		signature_key {
			filepath   = "%s"
			owner      = "%s"
			passphrase = "%s"
		}

		master_key {
			keyname = "%s"

			keysharefile {
				filepath = "%s"
				passphrase    = "%s"
			}

			keysharefile {
				filepath = "%s"
				passphrase    = "%s"
			}
		}
	}
`, instanceName, url, sigKeyFilepath, sigKeyOwner, sigKeyPassphrase, masterKeyName, keyShareFilepath1, keyShareToken1, keyShareFilepath2, keyShareToken2)
}

func testAccCheckIBMKmsCryptoUnitsWithRegionConfig(instanceName, url, sigKeyFilepath, sigKeyOwner, sigKeyPassphrase, masterKeyName, keyShareFilepath1, keyShareToken1, keyShareFilepath2, keyShareToken2 string) string {
	return fmt.Sprintf(`
	resource "ibm_kms_cryptounits" "test" {
		instance_id = "%s"
		url = "%s"
		should_zeroize = true

		signature_key {
			filepath   = "%s"
			owner      = "%s"
			passphrase = "%s"
		}

		master_key {
			keyname = "%s"

			keysharefile {
				filepath = "%s"
				passphrase    = "%s"
			}

			keysharefile {
				filepath = "%s"
				passphrase    = "%s"
			}
		}
	}
`, instanceName, url, sigKeyFilepath, sigKeyOwner, sigKeyPassphrase, masterKeyName, keyShareFilepath1, keyShareToken1, keyShareFilepath2, keyShareToken2)
}

func testAccCheckIBMKmsCryptoUnitsMultipleKeySharesConfig(instanceName, url, sigKeyFilepath, sigKeyOwner, sigKeyPassphrase, masterKeyName, keyShareFilepath1, keyShareToken1, keyShareFilepath2, keyShareToken2, keyShareFilepath3, keyShareToken3 string) string {
	return fmt.Sprintf(`
	resource "ibm_kms_cryptounits" "test" {
		instance_id = "%s"
		url         = "%s"
		should_zeroize = true

		signature_key {
			filepath   = "%s"
			owner      = "%s"
			passphrase = "%s"
		}

		master_key {
			keyname = "%s"

			keysharefile {
				filepath = "%s"
				passphrase    = "%s"
			}

			keysharefile {
				filepath = "%s"
				passphrase    = "%s"
			}

			keysharefile {
				filepath = "%s"
				passphrase    = "%s"
			}
		}
	}
`, instanceName, url, sigKeyFilepath, sigKeyOwner, sigKeyPassphrase, masterKeyName, keyShareFilepath1, keyShareToken1, keyShareFilepath2, keyShareToken2, keyShareFilepath3, keyShareToken3)
}

// testAccCheckIBMKmsCryptoUnitsDestroy verifies that all crypto units associated
// with the ibm_kms_cryptounits resource have been zeroized after destroy.
func testAccCheckIBMKmsCryptoUnitsDestroy(s *terraform.State) error {
	ctx := context.Background()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_kms_cryptounits" {
			continue
		}

		instanceID := rs.Primary.Attributes["instance_id"]
		region := rs.Primary.Attributes["region"]
		url := rs.Primary.Attributes["url"]

		kpOpts, err := kpCryptoUnit.NewKeyProtectCryptoUnitAPIOptions(url)
		if err != nil || kpOpts.InstanceID == "" {
			// Fall back to building URL from instance_id + region
			if instanceID == "" || region == "" {
				return fmt.Errorf("ibm_kms_cryptounits destroy check: cannot determine endpoint (instance_id=%s, region=%s)", instanceID, region)
			}
			serviceURL, urlErr := kpCryptoUnit.GetServiceURLForRegion(instanceID, region, false)
			if urlErr != nil {
				return fmt.Errorf("ibm_kms_cryptounits destroy check: failed to resolve URL: %v", urlErr)
			}
			kpOpts, err = kpCryptoUnit.NewKeyProtectCryptoUnitAPIOptions(serviceURL)
			if err != nil {
				return fmt.Errorf("ibm_kms_cryptounits destroy check: failed to create options: %v", err)
			}
		}
		if kpOpts.InstanceID == "" {
			kpOpts.InstanceID = instanceID
		}
		if kpOpts.Region == "" {
			kpOpts.Region = region
		}

		client, err := acc.TestAccProvider.Meta().(conns.ClientSession).KeyProtectCryptoUnitAPI(ctx, kpOpts)
		if err != nil {
			return fmt.Errorf("ibm_kms_cryptounits destroy check: failed to create client: %v", err)
		}

		cryptoUnitsResponse, _, err := client.ListCryptoUnitsWithContext(ctx)
		if err != nil {
			return fmt.Errorf("ibm_kms_cryptounits destroy check: failed to list crypto units: %v", err)
		}

		for _, cu := range cryptoUnitsResponse.CryptoUnits {
			if cu.State != kpCryptoUnit.CryptoUnitStateReserved {
				return fmt.Errorf("ibm_kms_cryptounits destroy check: crypto unit %s is in state %q, expected %q",
					cu.ID, cu.State, kpCryptoUnit.CryptoUnitStateReserved)
			}
		}
	}

	return nil
}

// Made with Bob
