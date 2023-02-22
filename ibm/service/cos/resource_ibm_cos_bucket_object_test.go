// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cos_test

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"regexp"
	"testing"
	"time"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCOSBucketObject_basic(t *testing.T) {
	name := fmt.Sprintf("tf-testacc-cos-%d", acctest.RandIntRange(10, 100))
	instanceCRN := acc.CosCRN
	objectBody := "Acceptance Testing"
	objectFile := "../../test-fixtures/cosObject.json"
	objectFileBody, _ := ioutil.ReadFile(objectFile)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCOS(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMCOSBucketObjectConfig_plaintext(name, instanceCRN, objectBody),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc", "id"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc", "content"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc", "content_length"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc", "content_type"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc", "etag"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc", "last_modified"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc", "object_sql_url"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_object.testacc", "body", objectBody),
				),
			},
			{
				Config: testAccIBMCOSBucketObjectConfig_base64(name, instanceCRN, objectBody),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc", "id"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc", "content_base64"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc", "content_length"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc", "content_type"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc", "etag"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc", "last_modified"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc", "object_sql_url"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_object.testacc", "body", objectBody),
				),
			},
			{
				Config: testAccIBMCOSBucketObjectConfig_file(name, instanceCRN, objectFile),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc", "id"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc", "content_file"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc", "content_length"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc", "content_type"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc", "etag"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc", "last_modified"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc", "object_sql_url"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_object.testacc", "body", string(objectFileBody)),
				),
			},
		},
	})
}

func TestAccIBMCOSBucketObjectlock_Retention_Without_Mode(t *testing.T) {
	name := fmt.Sprintf("tf-testacc-cos-%d", acctest.RandIntRange(10, 100))
	instanceCRN := acc.CosCRN
	objectBody := "Acceptance Testing"
	retainUntilDate := time.Now().Local().Add(time.Second * 5)
	retainUntilDateString := retainUntilDate.Format(time.RFC3339)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCOS(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccIBMCOSBucketObjectlock_retention_without_mode(name, instanceCRN, objectBody, retainUntilDateString),
				ExpectError: regexp.MustCompile("MalformedXML: The XML you provided was not well-formed or did not validate against our published schema."),
			},
		},
	})
}

func TestAccIBMCOSBucketObjectlock_retention_without_objectlock_enabled(t *testing.T) {
	name := fmt.Sprintf("tf-testacc-cos-%d", acctest.RandIntRange(10, 100))
	instanceCRN := acc.CosCRN
	retainUntilDate := time.Now().Local().Add(time.Second * 22)
	retainUntilDateString := retainUntilDate.Format(time.RFC3339)
	objectBody := "Acceptance Testing"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCOS(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccIBMCOSBucketObjectlock_retention_without_objectlock_enabled(name, instanceCRN, objectBody, retainUntilDateString),
				ExpectError: regexp.MustCompile("InvalidRequest: Bucket is missing Object Lock Configuration"),
			},
		},
	})
}

func TestAccIBMCOSBucketObjectlock_legalhold_without_objectlock_enabled(t *testing.T) {
	name := fmt.Sprintf("tf-testacc-cos-%d", acctest.RandIntRange(10, 100))
	instanceCRN := acc.CosCRN
	retainUntilDate := time.Now().Local().Add(time.Second * 22)
	retainUntilDateString := retainUntilDate.Format(time.RFC3339)
	objectBody := "Acceptance Testing"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCOS(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccIBMCOSBucketObjectlock_legalhold_without_objectlock_enabled(name, instanceCRN, objectBody, retainUntilDateString),
				ExpectError: regexp.MustCompile("InvalidRequest: Bucket is missing Object Lock Configuration"),
			},
		},
	})
}
func TestAccIBMCOSBucketObjectlock_Retention_Invalid_Mode(t *testing.T) {
	name := fmt.Sprintf("tf-testacc-cos-%d", acctest.RandIntRange(10, 100))
	instanceCRN := acc.CosCRN
	objectBody := "Acceptance Testing"
	retainUntilDate := time.Now().Local().Add(time.Second * 5)
	retainUntilDateString := retainUntilDate.Format(time.RFC3339)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCOS(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccIBMCOSBucketObjectlock_retention_invalid_mode(name, instanceCRN, objectBody, retainUntilDateString),
				ExpectError: regexp.MustCompile("MalformedXML: The XML you provided was not well-formed or did not validate against our published schema."),
			},
		},
	})
}

func TestAccIBMCOSBucketObjectlock_Retention_Retainuntildate_Past(t *testing.T) {
	name := fmt.Sprintf("tf-testacc-cos-%d", acctest.RandIntRange(10, 100))
	instanceCRN := acc.CosCRN
	objectBody := "Acceptance Testing"
	retainUntilDate := time.Now().Local().Add(time.Second * 5)
	retainUntilDateString := retainUntilDate.Format(time.RFC3339)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCOS(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccIBMCOSBucketObjectlock_Retention_Retainuntildate_Past(name, instanceCRN, objectBody, retainUntilDateString),
				ExpectError: regexp.MustCompile("InvalidArgument: The retain until date must be in the future!"),
			},
		},
	})
}

func TestAccIBMCOSBucketObjectlock_Retention_Without_Retainuntildate(t *testing.T) {
	name := fmt.Sprintf("tf-testacc-cos-%d", acctest.RandIntRange(10, 100))
	instanceCRN := acc.CosCRN
	objectBody := "Acceptance Testing"
	retainUntilDate := time.Now().Local().Add(time.Second * 5)
	retainUntilDateString := retainUntilDate.Format(time.RFC3339)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCOS(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccIBMCOSBucketObjectlock_Retention_Without_Retainuntildate(name, instanceCRN, objectBody, retainUntilDateString),
				ExpectError: regexp.MustCompile("MalformedXML: The XML you provided was not well-formed or did not validate against our published schema."),
			},
		},
	})
}

func testAccIBMCOSBucketObjectConfig_plaintext(name string, instanceCRN string, objectBody string) string {
	return fmt.Sprintf(`
		resource "ibm_cos_bucket" "testacc" {
			bucket_name          = "%[1]s"
			resource_instance_id = "%[2]s"
			region_location      = "us-east"
			storage_class        = "standard"
		}
		resource "ibm_cos_bucket_object" "testacc" {
			bucket_crn	    = ibm_cos_bucket.testacc.crn
			bucket_location = ibm_cos_bucket.testacc.region_location
			key 					  = "%[1]s.txt"
			content			    = "%[3]s"
		}`, name, instanceCRN, objectBody)
}

func testAccIBMCOSBucketObjectConfig_base64(name string, instanceCRN string, objectBody string) string {
	return fmt.Sprintf(`
		resource "ibm_cos_bucket" "testacc" {
			bucket_name          = "%[1]s"
			resource_instance_id = "%[2]s"
			region_location      = "us-east"
			storage_class        = "standard"
		}
		resource "ibm_cos_bucket_object" "testacc" {
			bucket_crn	    = ibm_cos_bucket.testacc.crn
			bucket_location = ibm_cos_bucket.testacc.region_location
			key 					  = "%[1]s.txt"
			content_base64  = "%[3]s"
		}`, name, instanceCRN, base64.StdEncoding.EncodeToString([]byte(objectBody)))
}

func testAccIBMCOSBucketObjectConfig_file(name string, instanceCRN string, objectFile string) string {
	return fmt.Sprintf(`
		resource "ibm_cos_bucket" "testacc" {
			bucket_name          = "%[1]s"
			resource_instance_id = "%[2]s"
			region_location      = "us-east"
			storage_class        = "standard"
		}
		resource "ibm_cos_bucket_object" "testacc" {
			bucket_crn	    = ibm_cos_bucket.testacc.crn
			bucket_location = ibm_cos_bucket.testacc.region_location
			key 					  = "%[1]s.txt"
			content_file	  = "%[3]s"
		}`, name, instanceCRN, objectFile)
}

func testAccIBMCOSBucketObjectlock_retention_without_objectlock_enabled(name string, instanceCRN string, objectBody string, retainUntilDate string) string {
	return fmt.Sprintf(`
		resource "ibm_cos_bucket" "testacc" {
			bucket_name          = "%[1]s"
			resource_instance_id = "%[2]s"
			cross_region_location      = "us"
			storage_class        = "standard"
			object_versioning {
				enable  = true
			  }
		}
		resource "ibm_cos_bucket_object" "testacc" {
			bucket_crn	    = ibm_cos_bucket.testacc.crn
			bucket_location = ibm_cos_bucket.testacc.cross_region_location
			key 					  = "%[1]s.txt"
			content			    = "%[3]s"
			object_lock_mode    = "COMPLIANCE"
			object_lock_retain_until_date = "%[4]s"
   			force_delete = true
		}`, name, instanceCRN, objectBody, retainUntilDate)
}

func testAccIBMCOSBucketObjectlock_legalhold_without_objectlock_enabled(name string, instanceCRN string, objectBody string, retainUntilDate string) string {
	return fmt.Sprintf(`
		resource "ibm_cos_bucket" "testacc" {
			bucket_name          = "%[1]s"
			resource_instance_id = "%[2]s"
			cross_region_location      = "us"
			storage_class        = "standard"
			object_versioning {
				enable  = true
			  }
		}
		resource "ibm_cos_bucket_object" "testacc" {
			bucket_crn	    = ibm_cos_bucket.testacc.crn
			bucket_location = ibm_cos_bucket.testacc.cross_region_location
			key 					  = "%[1]s.txt"
			content			    = "%[3]s"
			object_lock_legal_hold_status = "ON"
   			force_delete = true
		}`, name, instanceCRN, objectBody, retainUntilDate)
}
func testAccIBMCOSBucketObjectlock_retention_without_mode(name string, instanceCRN string, objectBody string, retainUntilDate string) string {
	return fmt.Sprintf(`
		resource "ibm_cos_bucket" "testacc" {
			bucket_name          = "%[1]s"
			resource_instance_id = "%[2]s"
			cross_region_location      = "us"
			storage_class        = "standard"
			object_versioning {
				enable  = true
			  }
			  object_lock = true
		}
		resource "ibm_cos_bucket_object" "testacc" {
			bucket_crn	    = ibm_cos_bucket.testacc.crn
			bucket_location = ibm_cos_bucket.testacc.cross_region_location
			key 					  = "%[1]s.txt"
			content			    = "%[3]s"
			object_lock_retain_until_date = "%[4]s"
   			force_delete = true
		}`, name, instanceCRN, objectBody, retainUntilDate)
}

func testAccIBMCOSBucketObjectlock_retention_invalid_mode(name string, instanceCRN string, objectBody string, retainUntilDate string) string {
	return fmt.Sprintf(`
		resource "ibm_cos_bucket" "testacc" {
			bucket_name          = "%[1]s"
			resource_instance_id = "%[2]s"
			cross_region_location      = "us"
			storage_class        = "standard"
			object_versioning {
				enable  = true
			  }
			  object_lock = true
		}
		resource "ibm_cos_bucket_object" "testacc" {
			bucket_crn	    = ibm_cos_bucket.testacc.crn
			bucket_location = ibm_cos_bucket.testacc.cross_region_location
			key 					  = "%[1]s.txt"
			content			    = "%[3]s"
			object_lock_mode              = "Invalid"
			object_lock_retain_until_date = "%[4]s"
   			force_delete = true
		}`, name, instanceCRN, objectBody, retainUntilDate)
}

func testAccIBMCOSBucketObjectlock_Retention_Retainuntildate_Past(name string, instanceCRN string, objectBody string, retainUntilDate string) string {
	return fmt.Sprintf(`
		resource "ibm_cos_bucket" "testacc" {
			bucket_name          = "%[1]s"
			resource_instance_id = "%[2]s"
			cross_region_location      = "us"
			storage_class        = "standard"
			object_versioning {
				enable  = true
			  }
			  object_lock = true
		}
		resource "ibm_cos_bucket_object" "testacc" {
			bucket_crn	    = ibm_cos_bucket.testacc.crn
			bucket_location = ibm_cos_bucket.testacc.cross_region_location
			key 					  = "%[1]s.txt"
			content			    = "%[3]s"
			object_lock_mode              = "COMPLIANCE"
			object_lock_retain_until_date = "%[4]s"
   			force_delete = true
		}`, name, instanceCRN, objectBody, retainUntilDate)
}

func testAccIBMCOSBucketObjectlock_Retention_Without_Retainuntildate(name string, instanceCRN string, objectBody string, retainUntilDate string) string {
	return fmt.Sprintf(`
		resource "ibm_cos_bucket" "testacc" {
			bucket_name          = "%[1]s"
			resource_instance_id = "%[2]s"
			cross_region_location      = "us"
			storage_class        = "standard"
			object_versioning {
				enable  = true
			  }
			  object_lock = true
		}
		resource "ibm_cos_bucket_object" "testacc" {
			bucket_crn	    = ibm_cos_bucket.testacc.crn
			bucket_location = ibm_cos_bucket.testacc.cross_region_location
			key 					  = "%[1]s.txt"
			content			    = "%[3]s"
			object_lock_mode              = "COMPLIANCE"
   			force_delete = true
		}`, name, instanceCRN, objectBody, retainUntilDate)
}
func testAccIBMCOSBucketObjectlock_legalhold_off(name string, instanceCRN string, objectBody string) string {
	return fmt.Sprintf(`
		resource "ibm_cos_bucket" "testacc" {
			bucket_name          = "%[1]s"
			resource_instance_id = "%[2]s"
			cross_region_location      = "us"
			storage_class        = "standard"
			object_versioning {
				enable  = true
			  }
			  object_lock = true
		}
		resource "ibm_cos_bucket_object" "testacc" {
			bucket_crn	    = ibm_cos_bucket.testacc.crn
			bucket_location = ibm_cos_bucket.testacc.cross_region_location
			key 					  = "%[1]s.txt"
			content			    = "%[3]s"
			object_lock_legal_hold_status = "OFF"
   			force_delete = true
		}`, name, instanceCRN, objectBody)
}
