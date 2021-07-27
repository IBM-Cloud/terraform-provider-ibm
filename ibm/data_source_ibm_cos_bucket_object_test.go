// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCOSBucketObjectDataSource_basic(t *testing.T) {
	name := "tf-testacc-cos"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCOS(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMCOSBucketObjectDataSourceConfig_basic(name, cosCRN),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cos_bucket_object.testacc", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_cos_bucket_object.testacc", "body"),
					resource.TestCheckResourceAttrSet("data.ibm_cos_bucket_object.testacc", "content_length"),
					resource.TestCheckResourceAttrSet("data.ibm_cos_bucket_object.testacc", "content_type"),
					resource.TestCheckResourceAttrSet("data.ibm_cos_bucket_object.testacc", "etag"),
					resource.TestCheckResourceAttrSet("data.ibm_cos_bucket_object.testacc", "last_modified"),
				),
			},
		},
	})
}

func testAccIBMCOSBucketObjectDataSourceConfig_basic(name string, crn string) string {
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
			content			    = "Acceptance testing"
		}
		data "ibm_cos_bucket_object" "testacc" {
			bucket_crn	    = ibm_cos_bucket.testacc.crn
			bucket_location = ibm_cos_bucket.testacc.region_location
			key             = ibm_cos_bucket_object.testacc.key
		}`, name, crn)
}
