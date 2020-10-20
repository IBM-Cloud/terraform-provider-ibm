package ibm

import (
	"fmt"
	"math/rand"
	"regexp"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMKMSResource_basic(t *testing.T) {
	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	cosInstanceName := fmt.Sprintf("cos_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("bucket-test77")
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))
	payload := "LqMWNtSi3Snr4gFNO0PsFFLFRNs57mSXCQE7O2oE+g0="
	hpcskeyName := fmt.Sprintf("hpcs_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMKmsResourceStandardConfig(instanceName, keyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMKmsResourceImportStandardConfig(instanceName, keyName, payload),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMKmsResourceRootkeyWithCOSConfig(instanceName, keyName, cosInstanceName, bucketName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMKmsResourceHpcsConfig(hpcsInstanceID, hpcskeyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key.hpcstest", "key_name", hpcskeyName),
				),
			},
		},
	})
}

// Test for valid expiration date for create key operation
func TestAccIBMKMSResource_ValidExpDate(t *testing.T) {

	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))

	hours := time.Duration(rand.Intn(24) + 1)
	mins := time.Duration(rand.Intn(60) + 1)
	sec := time.Duration(rand.Intn(60) + 1)
	expirationDateValid := ((time.Now().Add(time.Hour*hours + time.Minute*mins + time.Second*sec)).Format(time.RFC3339))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMKmsCreateStandardKeyConfig(instanceName, keyName, expirationDateValid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName),
					resource.TestCheckResourceAttr("ibm_kms_key.test", "expiration_date", expirationDateValid),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMKmsCreateRootKeyConfig(instanceName, keyName, expirationDateValid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName),
					resource.TestCheckResourceAttr("ibm_kms_key.test", "expiration_date", expirationDateValid),
				),
			},
		},
	})
}

// Test for invalid expiration date for create key operation
func TestAccIBMKMSResource_InvalidExpDate(t *testing.T) {
	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))

	hours := time.Duration(rand.Intn(24) + 1)
	mins := time.Duration(rand.Intn(60) + 1)
	sec := time.Duration(rand.Intn(60) + 1)
	expirationDateInvalid := (time.Now().Add(time.Hour*hours + time.Minute*mins + time.Second*sec)).String()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:      testAccCheckIBMKmsCreateStandardKeyConfig(instanceName, keyName, expirationDateInvalid),
				ExpectError: regexp.MustCompile("Invalid time format"),
			},
			resource.TestStep{
				Config:      testAccCheckIBMKmsCreateRootKeyConfig(instanceName, keyName, expirationDateInvalid),
				ExpectError: regexp.MustCompile("Invalid time format"),
			},
		},
	})
}

func testAccCheckIBMKmsResourceStandardConfig(instanceName, KeyName string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kms_instance" {
		name              = "%s"
		service           = "kms"
		plan              = "tiered-pricing"
		location          = "us-south"
	  }
	  resource "ibm_kms_key" "test" {
		instance_id = "${ibm_resource_instance.kms_instance.guid}"
		key_name = "%s"
		standard_key =  true
		force_delete = true
	}
	
`, instanceName, KeyName)
}

func testAccCheckIBMKmsResourceImportStandardConfig(instanceName, KeyName, payload string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kms_instance" {
		name              = "%s"
		service           = "kms"
		plan              = "tiered-pricing"
		location          = "us-south"
	  }
	  resource "ibm_kms_key" "test" {
		instance_id = "${ibm_resource_instance.kms_instance.guid}"
		key_name = "%s"
		standard_key =  true
		payload = "%s"
		force_delete = true
	}

`, instanceName, KeyName, payload)
}

func testAccCheckIBMKmsResourceRootkeyWithCOSConfig(instanceName, KeyName, cosInstanceName, bucketName string) string {
	return fmt.Sprintf(`
	provider "ibm" {
		region = "us-south"
	}	
	resource "ibm_resource_instance" "kms_instance1" {
		name              = "%s"
		service           = "kms"
		plan              = "tiered-pricing"
		location          = "us-south"
	  }
	  resource "ibm_kms_key" "test" {
		instance_id = "${ibm_resource_instance.kms_instance1.guid}"
		key_name = "%s"
		standard_key =  false
		force_delete = true
	}

	resource "ibm_resource_instance" "cos_instance" {
		name     = "%s"
		service  = "cloud-object-storage"
		plan     = "standard"
		location = "global"
	}
	
	resource "ibm_iam_authorization_policy" "policy" {
		source_service_name = "cloud-object-storage"
		target_service_name = "kms"
		roles               = ["Reader"]
	}
	
	resource "ibm_cos_bucket" "flex-us-south" {
		depends_on           = [ibm_iam_authorization_policy.policy]
		bucket_name          = "%s"
		resource_instance_id = ibm_resource_instance.cos_instance.id
		region_location      = "us-south"
		storage_class        = "flex"
		key_protect          = ibm_kms_key.test.id
	}
	
`, instanceName, KeyName, cosInstanceName, bucketName)
}

func testAccCheckIBMKmsResourceHpcsConfig(hpcsInstanceID, KeyName string) string {
	return fmt.Sprintf(`
	  resource "ibm_kms_key" "hpcstest" {
		instance_id = "%s"
		key_name = "%s"
		standard_key =  true
		force_delete = true
	}

`, hpcsInstanceID, KeyName)
}

func testAccCheckIBMKmsCreateStandardKeyConfig(instanceName, KeyName, expirationDate string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kms_instance" {
		name              = "%s"
		service           = "kms"
		plan              = "tiered-pricing"
		location          = "us-south"
	  }
	  resource "ibm_kms_key" "test" {
		instance_id = "${ibm_resource_instance.kms_instance.guid}"
		key_name = "%s"
		standard_key =  true
		force_delete = true
		expiration_date = "%s"
	}
	
`, instanceName, KeyName, expirationDate)
}

func testAccCheckIBMKmsCreateRootKeyConfig(instanceName, KeyName, expirationDate string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "kms_instance" {
		name              = "%s"
		service           = "kms"
		plan              = "tiered-pricing"
		location          = "us-south"
	  }
	  resource "ibm_kms_key" "test" {
		instance_id = "${ibm_resource_instance.kms_instance.guid}"
		key_name = "%s"
		standard_key =  false
		force_delete = true
		expiration_date = "%s"
	}
	
`, instanceName, KeyName, expirationDate)
}
