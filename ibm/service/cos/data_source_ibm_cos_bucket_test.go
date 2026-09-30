// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cos_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
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

func TestAccIBMCOSBucketDataSource_KeyProtect(t *testing.T) {
	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us"
	bucketClass := "standard"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMCOSBucketDataSourceConfig_keyProtectRead(instanceName, keyName, serviceName, bucketName, bucketRegion, bucketClass),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cos_bucket.testacc", "key_protect"),
					resource.TestCheckResourceAttrSet("data.ibm_cos_bucket.testacc", "kms_key_crn"),
					resource.TestCheckResourceAttrPair("data.ibm_cos_bucket.testacc", "key_protect", "data.ibm_cos_bucket.testacc", "kms_key_crn"),
					resource.TestCheckResourceAttrPair("data.ibm_cos_bucket.testacc", "key_protect", "ibm_cos_bucket.bucket", "key_protect"),
					resource.TestCheckResourceAttrPair("data.ibm_cos_bucket.testacc", "kms_key_crn", "ibm_cos_bucket.bucket", "kms_key_crn"),
				),
			},
		},
	})
}

func testAccIBMCOSBucketDataSourceConfig_keyProtectRead(instanceName string, keyName string, serviceName string, bucketName string, bucketRegion string, bucketClass string) string {
	return fmt.Sprintf(`%s

	data "ibm_cos_bucket" "testacc" {
		bucket_name          = ibm_cos_bucket.bucket.bucket_name
		resource_instance_id = ibm_cos_bucket.bucket.resource_instance_id
		bucket_type          = "cross_region_location"
		bucket_region        = "%s"
	}`, testAccCheckIBMKeyProtectRootkeyWithCOSBucket(instanceName, keyName, serviceName, bucketName, bucketRegion, bucketClass), bucketRegion)
}
