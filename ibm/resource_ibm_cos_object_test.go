// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCOSObject_basic(t *testing.T) {
	name := "tf-testacc-cos"
	instanceCRN := cosCRN
	objectBody := "Acceptance Testing"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCOS(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMCOSObjectConfig_basic(name, instanceCRN, objectBody),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_cos_object.testacc", "id"),
					resource.TestCheckResourceAttrSet("ibm_cos_object.testacc", "content_length"),
					resource.TestCheckResourceAttrSet("ibm_cos_object.testacc", "content_type"),
					resource.TestCheckResourceAttrSet("ibm_cos_object.testacc", "etag"),
					resource.TestCheckResourceAttrSet("ibm_cos_object.testacc", "last_modified"),
					resource.TestCheckResourceAttr("ibm_cos_object.testacc", "body", objectBody),
				),
			},
		},
	})
}

func testAccIBMCOSObjectConfig_basic(name string, instanceCRN string, objectBody string) string {
	return fmt.Sprintf(`
		resource "ibm_cos_bucket" "testacc" {
			bucket_name          = "%[1]s"
			resource_instance_id = "%[2]s"
			region_location      = "us-east"
			storage_class        = "standard"
		}
		resource "ibm_cos_object" "testacc" {
			bucket_crn	    = ibm_cos_bucket.testacc.crn
			bucket_location = ibm_cos_bucket.testacc.region_location
			key 					  = "%[1]s.txt"
			content			    = "%[3]s"
		}`, name, instanceCRN, objectBody)
}
