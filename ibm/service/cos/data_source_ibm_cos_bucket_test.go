// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cos_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCOSBucketDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCOS(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMCOSBucketDataSourceConfig_basic_read(acc.BucketName, acc.CosCRN, acc.RegionName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cos_bucket.testacc", "bucket_name"),
					resource.TestCheckResourceAttrSet("data.ibm_cos_bucket.testacc", "storage_class"),
					resource.TestCheckResourceAttrSet("data.ibm_cos_bucket.testacc", "region_location"),
				),
			},
		},
	})
}

func testAccIBMCOSBucketDataSourceConfig_basic_read(name string, crn string, region string) string {
	return fmt.Sprintf(`

		data "ibm_cos_bucket" "testacc" {
			bucket_name          = "%[1]s"
			resource_instance_id = "%[2]s"
			bucket_type          = "region_location"
			bucket_region        = "%[3]s"
		}`, name, crn, region)
}
