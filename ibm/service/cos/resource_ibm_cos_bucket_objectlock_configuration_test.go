package cos_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCosBucket_Objectlock_Bucket_Enabled(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"
	objectLockEnabled := true
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Objectlock_Bucket_Enabled(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, objectLockEnabled),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "cross_region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "object_versioning.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "object_lock", "true"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Objectlock_Bucket_Enabled_Smart_tier_bucket(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform-smart-tier-ol%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us"
	bucketClass := "smart"
	bucketRegionType := "cross_region_location"
	objectLockEnabled := true
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Objectlock_Bucket_Enabled(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, objectLockEnabled),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "cross_region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "object_versioning.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "object_lock", "true"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Objectlock_Configuration_Without_Rule(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"
	objectLockEnabled := true
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Objectlock_Without_Rule(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, objectLockEnabled),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "cross_region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "object_versioning.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "object_lock", "true"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Objectlock_Configuration_Valid_Mode_and_Days(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"
	objectLockEnabled := true
	mode := "COMPLIANCE"
	days := int64(4)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Objectlock_Valid_Mode_and_Days(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, objectLockEnabled, mode, days),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "cross_region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "object_versioning.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "object_lock", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_object_lock_configuration.objectlock", "object_lock_configuration.#", "1"),
				),
			},
		},
	})
}
func TestAccIBMCosBucket_Objectlock_Configuration_Valid_Mode_and_Years(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"
	objectLockEnabled := true
	mode := "COMPLIANCE"
	years := int64(1)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Objectlock_Valid_Mode_and_Years(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, objectLockEnabled, mode, years),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "cross_region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "object_versioning.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "object_lock", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_object_lock_configuration.objectlock", "object_lock_configuration.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Objectlock_Configuration_ExistingBucket(t *testing.T) {
	bucketRegion := "us"
	bucketCRN := acc.BucketCRN
	objectLockEnabled := true
	mode := "COMPLIANCE"
	days := int64(4)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Objectlock_Existing_bucket(bucketCRN, bucketRegion, objectLockEnabled, mode, days),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cos_bucket_object_lock_configuration.objectlock", "object_lock_configuration.0.object_lock_enabled", "Enabled"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_object_lock_configuration.objectlock", "object_lock_configuration.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Objectlock_Configuration_Updating_Objectlockrule_Years(t *testing.T) {
	bucketRegion := "us"
	bucketCRN := acc.BucketCRN
	mode := "COMPLIANCE"
	years := int64(4)
	updated_years := int64(6)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Objectlock_Existing_bucket_Years(bucketCRN, bucketRegion, mode, years),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cos_bucket_object_lock_configuration.objectlock", "object_lock_configuration.0.object_lock_enabled", "Enabled"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_object_lock_configuration.objectlock", "object_lock_configuration.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_object_lock_configuration.objectlock", "object_lock_configuration.0.object_lock_rule.0.default_retention.0.years", "4"),
				),
			},
			{
				Config: testAccCheckIBMCosBucket_Objectlock_Existing_bucket_Years(bucketCRN, bucketRegion, mode, updated_years),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cos_bucket_object_lock_configuration.objectlock", "object_lock_configuration.0.object_lock_enabled", "Enabled"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_object_lock_configuration.objectlock", "object_lock_configuration.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_object_lock_configuration.objectlock", "object_lock_configuration.0.object_lock_rule.0.default_retention.0.years", "6"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Objectlock_Configuration_Updating_Objectlockrule_Days(t *testing.T) {
	bucketRegion := "us"
	bucketCRN := acc.BucketCRN
	mode := "COMPLIANCE"
	days := int64(3)
	updated_days := int64(9)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Objectlock_Existing_bucket_Days(bucketCRN, bucketRegion, mode, days),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cos_bucket_object_lock_configuration.objectlock", "object_lock_configuration.0.object_lock_enabled", "Enabled"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_object_lock_configuration.objectlock", "object_lock_configuration.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_object_lock_configuration.objectlock", "object_lock_configuration.0.object_lock_rule.0.default_retention.0.days", "6"),
				),
			},
			{
				Config: testAccCheckIBMCosBucket_Objectlock_Existing_bucket_Days(bucketCRN, bucketRegion, mode, updated_days),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cos_bucket_object_lock_configuration.objectlock", "object_lock_configuration.0.object_lock_enabled", "Enabled"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_object_lock_configuration.objectlock", "object_lock_configuration.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_object_lock_configuration.objectlock", "object_lock_configuration.0.object_lock_rule.0.default_retention.0.days", "9"),
				),
			},
		},
	})
}
func TestAccIBMCosBucket_Objectlock_Configuration_Empty(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"
	objectLockEnabled := true
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCosBucket_Objectlock_Empty(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, objectLockEnabled),
				ExpectError: regexp.MustCompile("Error: Missing required argument"),
			},
		},
	})
}
func TestAccIBMCosBucket_Objectlock_Configuration_Without_Mode(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"
	objectLockEnabled := true
	years := int64(1)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCosBucket_Objectlock_Without_Mode(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, objectLockEnabled, years),
				ExpectError: regexp.MustCompile("Error: Missing required argument"),
			},
		},
	})
}

func TestAccIBMCosBucket_Objectlock_Configuration_With_Mode_Only(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"
	objectLockEnabled := true
	mode := "COMPLIANCE"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCosBucket_Objectlock_With_Mode_Only(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, objectLockEnabled, mode),
				ExpectError: regexp.MustCompile("MalformedXML: The XML you provided was not well-formed or did not validate against our published schema."),
			},
		},
	})
}

func TestAccIBMCosBucket_Objectlock_Configuration_With_Versioning_Not_Enabled(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"
	objectLockEnabled := true
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCosBucket_Objectlock_Versioning_not_enabled(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, objectLockEnabled),
				ExpectError: regexp.MustCompile("Error: Missing required argument"),
			},
		},
	})
}

func TestAccIBMCosBucket_Objectlock_Configuration_Invalid_Mode(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"
	objectLockEnabled := true
	mode := "INVALID"
	years := int64(1)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCosBucket_Objectlock_Invalid_Mode(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, objectLockEnabled, mode, years),
				ExpectError: regexp.MustCompile("MalformedXML: The XML you provided was not well-formed or did not validate against our published schema."),
			},
		},
	})
}

func TestAccIBMCosBucket_Objectlock_Configuration_Invalid_Days(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"
	objectLockEnabled := true
	mode := "COMPLIANCE"
	days := int64(-1)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCosBucket_Objectlock_Invalid_Days(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, objectLockEnabled, mode, days),
				ExpectError: regexp.MustCompile("MalformedXML: The XML you provided was not well-formed or did not validate against our published schema."),
			},
		},
	})
}

func testAccCheckIBMCosBucket_Objectlock_Bucket_Enabled(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, objectLockEnabled bool) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		name = "Default"
	}

	resource "ibm_resource_instance" "instance" {
		name              = "%s"
		service           = "cloud-object-storage"
		plan              = "standard"
		location          = "global"
		resource_group_id = data.ibm_resource_group.cos_group.id
	}
	resource "ibm_cos_bucket" "bucket" {
		bucket_name           = "%s"
		resource_instance_id  = ibm_resource_instance.instance.id
	    cross_region_location = "%s"
		storage_class         = "%s"
		object_versioning {
			enable  = true
		}
		object_lock = "%t"
	}
	`, cosServiceName, bucketName, region, storageClass, objectLockEnabled)
}

func testAccCheckIBMCosBucket_Objectlock_Versioning_not_enabled(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, objectLockEnabled bool) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		name = "Default"
	}

	resource "ibm_resource_instance" "instance" {
		name              = "%s"
		service           = "cloud-object-storage"
		plan              = "standard"
		location          = "global"
		resource_group_id = data.ibm_resource_group.cos_group.id
	}
	resource "ibm_cos_bucket" "bucket" {
		bucket_name           = "%s"
		resource_instance_id  = ibm_resource_instance.instance.id
	    cross_region_location = "%s"
		storage_class         = "%s"
		object_lock = "%t"
	}
	`, cosServiceName, bucketName, region, storageClass, objectLockEnabled)
}
func testAccCheckIBMCosBucket_Objectlock_Without_Rule(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, objectLockEnabled bool) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		name = "Default"
	}

	resource "ibm_resource_instance" "instance" {
		name              = "%s"
		service           = "cloud-object-storage"
		plan              = "standard"
		location          = "global"
		resource_group_id = data.ibm_resource_group.cos_group.id
	}
	resource "ibm_cos_bucket" "bucket" {
		bucket_name           = "%s"
		resource_instance_id  = ibm_resource_instance.instance.id
	    cross_region_location = "%s"
		storage_class         = "%s"
		object_versioning {
			enable  = true
		}
		object_lock = "%t"
	}
	resource ibm_cos_bucket_object_lock_configuration "objectlock" {
		bucket_crn      = ibm_cos_bucket.bucket.crn
		bucket_location = ibm_cos_bucket.bucket.cross_region_location
		object_lock_configuration{
		   object_lock_enabled = "Enabled"
		 }	   
	   }
	`, cosServiceName, bucketName, region, storageClass, objectLockEnabled)
}

func testAccCheckIBMCosBucket_Objectlock_Valid_Mode_and_Days(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, objectLockEnabled bool, mode string, days int64) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		name = "Default"
	}

	resource "ibm_resource_instance" "instance" {
		name              = "%s"
		service           = "cloud-object-storage"
		plan              = "standard"
		location          = "global"
		resource_group_id = data.ibm_resource_group.cos_group.id
	}
	resource "ibm_cos_bucket" "bucket" {
		bucket_name           = "%s"
		resource_instance_id  = ibm_resource_instance.instance.id
	    cross_region_location = "%s"
		storage_class         = "%s"
		object_versioning {
			enable  = true
		}
		object_lock = "%t"
	}
	resource ibm_cos_bucket_object_lock_configuration "objectlock" {
		bucket_crn      = ibm_cos_bucket.bucket.crn
		bucket_location = ibm_cos_bucket.bucket.cross_region_location
		object_lock_configuration{
			object_lock_enabled = "Enabled"
			object_lock_rule{
			  default_retention{
				mode = "%s"
				days = "%d"
				}
			}
		  }   
	   }
	`, cosServiceName, bucketName, region, storageClass, objectLockEnabled, mode, days)
}

func testAccCheckIBMCosBucket_Objectlock_Valid_Mode_and_Years(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, objectLockEnabled bool, mode string, years int64) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		name = "Default"
	}

	resource "ibm_resource_instance" "instance" {
		name              = "%s"
		service           = "cloud-object-storage"
		plan              = "standard"
		location          = "global"
		resource_group_id = data.ibm_resource_group.cos_group.id
	}
	resource "ibm_cos_bucket" "bucket" {
		bucket_name           = "%s"
		resource_instance_id  = ibm_resource_instance.instance.id
	    cross_region_location = "%s"
		storage_class         = "%s"
		object_versioning {
			enable  = true
		}
		object_lock = "%t"
	}
	resource ibm_cos_bucket_object_lock_configuration "objectlock" {
		bucket_crn      = ibm_cos_bucket.bucket.crn
		bucket_location = ibm_cos_bucket.bucket.cross_region_location
		object_lock_configuration{
			object_lock_enabled = "Enabled"
			object_lock_rule{
			  default_retention{
				mode = "%s"
				years = "%d"
				}
			}
		  }   
	   }
	`, cosServiceName, bucketName, region, storageClass, objectLockEnabled, mode, years)
}

func testAccCheckIBMCosBucket_Objectlock_Empty(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, objectLockEnabled bool) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		name = "Default"
	}

	resource "ibm_resource_instance" "instance" {
		name              = "%s"
		service           = "cloud-object-storage"
		plan              = "standard"
		location          = "global"
		resource_group_id = data.ibm_resource_group.cos_group.id
	}
	resource "ibm_cos_bucket" "bucket" {
		bucket_name           = "%s"
		resource_instance_id  = ibm_resource_instance.instance.id
	    cross_region_location = "%s"
		storage_class         = "%s"
		object_versioning {
			enable  = true
		}
		object_lock = "%t"
	}
	resource ibm_cos_bucket_object_lock_configuration "objectlock" {
		bucket_crn      = ibm_cos_bucket.bucket.crn
		bucket_location = ibm_cos_bucket.bucket.cross_region_location
		object_lock_configuration{
		  }   
	}
	`, cosServiceName, bucketName, region, storageClass, objectLockEnabled)
}

func testAccCheckIBMCosBucket_Objectlock_Without_Mode(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, objectLockEnabled bool, years int64) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		name = "Default"
	}

	resource "ibm_resource_instance" "instance" {
		name              = "%s"
		service           = "cloud-object-storage"
		plan              = "standard"
		location          = "global"
		resource_group_id = data.ibm_resource_group.cos_group.id
	}
	resource "ibm_cos_bucket" "bucket" {
		bucket_name           = "%s"
		resource_instance_id  = ibm_resource_instance.instance.id
	    cross_region_location = "%s"
		storage_class         = "%s"
		object_versioning {
			enable  = true
		}
		object_lock = "%t"
	}
	resource ibm_cos_bucket_object_lock_configuration "objectlock" {
		bucket_crn      = ibm_cos_bucket.bucket.crn
		bucket_location = ibm_cos_bucket.bucket.cross_region_location
		object_lock_configuration{
			object_lock_enabled = "Enabled"
			object_lock_rule{
			  default_retention{
				years = "%d"
				}
			}
		  }   
	}
	`, cosServiceName, bucketName, region, storageClass, objectLockEnabled, years)
}

func testAccCheckIBMCosBucket_Objectlock_With_Mode_Only(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, objectLockEnabled bool, mode string) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		name = "Default"
	}

	resource "ibm_resource_instance" "instance" {
		name              = "%s"
		service           = "cloud-object-storage"
		plan              = "standard"
		location          = "global"
		resource_group_id = data.ibm_resource_group.cos_group.id
	}
	resource "ibm_cos_bucket" "bucket" {
		bucket_name           = "%s"
		resource_instance_id  = ibm_resource_instance.instance.id
	    cross_region_location = "%s"
		storage_class         = "%s"
		object_versioning {
			enable  = true
		}
		object_lock = "%t"
	}
	resource ibm_cos_bucket_object_lock_configuration "objectlock" {
		bucket_crn      = ibm_cos_bucket.bucket.crn
		bucket_location = ibm_cos_bucket.bucket.cross_region_location
		object_lock_configuration{
			object_lock_enabled = "Enabled"
			object_lock_rule{
			  default_retention{
				mode = "%s"
				}
			}
		  }   
	}
	`, cosServiceName, bucketName, region, storageClass, objectLockEnabled, mode)
}

func testAccCheckIBMCosBucket_Objectlock_Invalid_Mode(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, objectLockEnabled bool, mode string, years int64) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		name = "Default"
	}

	resource "ibm_resource_instance" "instance" {
		name              = "%s"
		service           = "cloud-object-storage"
		plan              = "standard"
		location          = "global"
		resource_group_id = data.ibm_resource_group.cos_group.id
	}
	resource "ibm_cos_bucket" "bucket" {
		bucket_name           = "%s"
		resource_instance_id  = ibm_resource_instance.instance.id
	    cross_region_location = "%s"
		storage_class         = "%s"
		object_versioning {
			enable  = true
		}
		object_lock = "%t"
	}
	resource ibm_cos_bucket_object_lock_configuration "objectlock" {
		bucket_crn      = ibm_cos_bucket.bucket.crn
		bucket_location = ibm_cos_bucket.bucket.cross_region_location
		object_lock_configuration{
			object_lock_enabled = "Enabled"
			object_lock_rule{
			  default_retention{
				mode = "%s"
				years = "%d"
				}
			}
		  }   
	   }
	`, cosServiceName, bucketName, region, storageClass, objectLockEnabled, mode, years)
}

func testAccCheckIBMCosBucket_Objectlock_Invalid_Days(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, objectLockEnabled bool, mode string, days int64) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		name = "Default"
	}

	resource "ibm_resource_instance" "instance" {
		name              = "%s"
		service           = "cloud-object-storage"
		plan              = "standard"
		location          = "global"
		resource_group_id = data.ibm_resource_group.cos_group.id
	}
	resource "ibm_cos_bucket" "bucket" {
		bucket_name           = "%s"
		resource_instance_id  = ibm_resource_instance.instance.id
	    cross_region_location = "%s"
		storage_class         = "%s"
		object_versioning {
			enable  = true
		}
		object_lock = "%t"
	}
	resource ibm_cos_bucket_object_lock_configuration "objectlock" {
		bucket_crn      = ibm_cos_bucket.bucket.crn
		bucket_location = ibm_cos_bucket.bucket.cross_region_location
		object_lock_configuration{
			object_lock_enabled = "Enabled"
			object_lock_rule{
			  default_retention{
				mode = "%s"
				days = "%d"
				}
			}
		  }   
	   }
	`, cosServiceName, bucketName, region, storageClass, objectLockEnabled, mode, days)
}

func testAccCheckIBMCosBucket_Objectlock_Existing_bucket(bucketCrn string, region string, objectLockEnabled bool, mode string, days int64) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		name = "Default"
	}

	resource ibm_cos_bucket_object_lock_configuration "objectlock" {
		bucket_crn      = "%s"
		bucket_location = "%s"
		object_lock_configuration{
			object_lock_enabled = "%t"
			object_lock_rule{
			  default_retention{
				mode = "%s"
				days = "%d"
				}
			}
		  }   
	   }
	`, bucketCrn, region, objectLockEnabled, mode, days)
}

func testAccCheckIBMCosBucket_Objectlock_Existing_bucket_Years(bucketCrn string, region string, mode string, years int64) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		name = "Default"
	}

	resource ibm_cos_bucket_object_lock_configuration "objectlock" {
		bucket_crn      = "%s"
		bucket_location = "%s"
		object_lock_configuration{
			object_lock_enabled = "Enabled"
			object_lock_rule{
			  default_retention{
				mode = "%s"
				years = "%d"
				}
			}
		  }   
	   }
	`, bucketCrn, region, mode, years)
}

func testAccCheckIBMCosBucket_Objectlock_Existing_bucket_Days(bucketCrn string, region string, mode string, days int64) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		name = "Default"
	}

	resource ibm_cos_bucket_object_lock_configuration "objectlock" {
		bucket_crn      = "%s"
		bucket_location = "%s"
		object_lock_configuration{
			object_lock_enabled = "Enabled"
			object_lock_rule{
			  default_retention{
				mode = "%s"
				days = "%d"
				}
			}
		  }   
	   }
	`, bucketCrn, region, mode, days)
}
