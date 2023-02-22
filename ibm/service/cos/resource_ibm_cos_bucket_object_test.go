// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cos_test

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"testing"

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

func TestAccIBMCOSBucketObject_VersioningEnabled(t *testing.T) {
	name := fmt.Sprintf("tf-testacc-cos-%d", acctest.RandIntRange(10, 100))
	key := "plaintext.txt"
	instanceCRN := acc.CosCRN
	objectBody1 := "Acceptance Testing"
	objectBody2 := "Acceptance Testing object 2"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCOS(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMCOSBucketBucketObject_Versioning_Enabled(name, key, instanceCRN, objectBody1, objectBody2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cos_bucket.testacc", "object_versioning.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.testacc", "object_versioning.0.enable", "true"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc_object", "id"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc_object", "content"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc_object", "content_length"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc_object", "content_type"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_object.testacc_object", "content", objectBody1),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc_object2", "id"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc_object2", "content"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc_object2", "content_length"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc_object2", "content_type"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_object.testacc_object2", "content", objectBody2),
				),
			},
		},
	})
}

func TestAccIBMCOSBucketObject_Versioning_Enabled_Sequential_Upload_on_Existing_Bucket(t *testing.T) {
	key := "plaintext.txt"
	bucketCRN := acc.BucketCRN
	objectBody1 := "Acceptance Testing"
	objectBody2 := "Acceptance Testing object 2"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCOS(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMCOSBucketBucketObjectUpload(bucketCRN, key, objectBody1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc_object", "id"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc_object", "content"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc_object", "content_length"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc_object", "content_type"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_object.testacc_object", "content", objectBody1),
				),
			},
			{
				Config: testAccIBMCOSBucketBucketObjectUpload(bucketCRN, key, objectBody2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc_object", "id"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc_object", "content"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc_object", "content_length"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc_object", "content_type"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_object.testacc_object", "content", objectBody2),
				),
			},
		},
	})
}

func TestAccIBMCOSBucketObject_Uploading_Multile_Objects_on_Existing_Bucket_without_Versioning(t *testing.T) {
	key := "plaintext.txt"
	bucketCRN := acc.BucketCRN
	objectBody1 := "Acceptance Testing"
	objectBody2 := "Acceptance Testing object 2"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCOS(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMCOSBucketBucketObjectUpload(bucketCRN, key, objectBody1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc_object", "id"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc_object", "content"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc_object", "content_length"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc_object", "content_type"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_object.testacc_object", "content", objectBody1),
				),
			},
			{
				Config: testAccIBMCOSBucketBucketObjectUpload(bucketCRN, key, objectBody2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc_object", "id"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc_object", "content"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc_object", "content_length"),
					resource.TestCheckResourceAttrSet("ibm_cos_bucket_object.testacc_object", "content_type"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_object.testacc_object", "content", objectBody2),
				),
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

func testAccIBMCOSBucketBucketObject_Versioning_Enabled(name string, key string, instanceCRN string, objectBody1 string, objectBody2 string) string {
	return fmt.Sprintf(`
		resource "ibm_cos_bucket" "testacc" {
			bucket_name          = "%[1]s"
			resource_instance_id = "%[3]s"
			region_location      = "us-south"
			storage_class        = "standard"
			object_versioning {
				enable  = true
			  }
		}
		resource "ibm_cos_bucket_object" "testacc_object" {
			bucket_crn	    = ibm_cos_bucket.testacc.crn
			bucket_location = ibm_cos_bucket.testacc.region_location
			key 			= "%[2]s"
			content			= "%[4]s"
		}
		resource "ibm_cos_bucket_object" "testacc_object2" {
			bucket_crn	    = ibm_cos_bucket.testacc.crn
			bucket_location = ibm_cos_bucket.testacc.region_location
			key 					  = "%[2]s"
			content			    = "%[5]s"
		}`, name, key, instanceCRN, objectBody1, objectBody2)
}

func testAccIBMCOSBucketBucketObjectUpload(bucketCrn string, key string, objectBody1 string) string {
	return fmt.Sprintf(`
		resource "ibm_cos_bucket_object" "testacc_object" {
			bucket_crn	    = "%[1]s"
			bucket_location = "us-south"
			key 			= "%[2]s"
			content			= "%[3]s"
		}`, bucketCrn, key, objectBody1)
}
