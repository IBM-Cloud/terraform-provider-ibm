// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCOSBucketObject_basic(t *testing.T) {
	name := fmt.Sprintf("tf-testacc-cos-%d", acctest.RandIntRange(10, 100))
	instanceCRN := cosCRN
	objectBody := "Acceptance Testing"
	objectFile := "test-fixtures/cosObject.json"
	objectFileBody, _ := ioutil.ReadFile(objectFile)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCOS(t) },
		Providers: testAccProviders,
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
					resource.TestCheckResourceAttr("ibm_cos_bucket_object.testacc", "body", string(objectFileBody)),
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
