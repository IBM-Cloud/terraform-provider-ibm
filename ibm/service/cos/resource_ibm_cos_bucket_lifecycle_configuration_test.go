package cos_test

import (
	"fmt"
	"regexp"
	"testing"
	// "time"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCosBucket_Lifecycle_Configuration_Expiration_With_Days(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform-lifecycle-configuration%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	expirationDays := 1

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Lifecycle_Configuration_Expiration_With_Days(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, expirationDays),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.0.expiration.0.days", "1"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Lifecycle_Configuration_Expiration_With_Date(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform-lifecycle-configuration%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	expirationDate := "2024-09-05T00:00:00Z"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Lifecycle_Configuration_Expiration_With_Date(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, expirationDate),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.0.expiration.0.date", expirationDate),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Lifecycle_Configuration_Transition_With_Days_Glacier(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform-lifecycle-configuration%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	transitionDays := 1
	tStorageclass := "GLACIER"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Lifecycle_Configuration_Transition_With_Days(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, transitionDays, tStorageclass),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.0.transition.0.days", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.0.transition.0.storage_class", "GLACIER"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Lifecycle_Configuration_Transition_With_Date_Glacier(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform-lifecycle-configuration%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	transitionDate := "2024-09-05T00:00:00Z"
	tStorageclass := "GLACIER"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Lifecycle_Configuration_Transition_With_Date(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, transitionDate, tStorageclass),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.0.transition.0.date", transitionDate),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.0.transition.0.storage_class", "GLACIER"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Lifecycle_Configuration_Transition_With_Days_Accelerated(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform-lifecycle-configuration%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	transitionDays := 1
	tStorageclass := "ACCELERATED"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Lifecycle_Configuration_Transition_With_Days(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, transitionDays, tStorageclass),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.0.transition.0.days", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.0.transition.0.storage_class", "ACCELERATED"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Lifecycle_Configuration_Expiration_With_Abort_Incomplete_Multipart_Upload(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform-lifecycle-configuration%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	expirationDays := 2
	daysAfterInitiation := 1

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Lifecycle_Configuration_Expiration_With_Abort_Incomplete_Multipart_Upload(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, expirationDays, daysAfterInitiation),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.0.expiration.0.days", "2"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.0.abort_incomplete_multipart_upload.0.days_after_initiation", "1"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Lifecycle_Configuration_Abort_Incomplete_Multipart_Upload(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform-lifecycle-configuration%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	daysAfterInitiation := 1

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Lifecycle_Configuration_Abort_Incomplete_Multipart_Upload(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, daysAfterInitiation),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.0.abort_incomplete_multipart_upload.0.days_after_initiation", "1"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Lifecycle_Configuration_Transition_With_Date_Accelerated(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform-lifecycle-configuration%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	transitionDate := "2024-09-05T00:00:00Z"
	tStorageclass := "ACCELERATED"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Lifecycle_Configuration_Transition_With_Date(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, transitionDate, tStorageclass),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.0.transition.0.date", transitionDate),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.0.transition.0.storage_class", "ACCELERATED"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Lifecycle_Configuration_Expiration_With_Transition(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform-lifecycle-configuration%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	expirationDays := 2
	transitionDays := 1
	tStorageClass := "GLACIER"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Lifecycle_Configuration_Expiration_With_Transition(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, expirationDays, transitionDays, tStorageClass),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.0.expiration.0.days", "2"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.0.transition.0.days", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.0.transition.0.storage_class", "GLACIER"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Lifecycle_Configuration_Transition_With_Abort_Incomplete_Multipart_Upload(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform-lifecycle-configuration%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	transitionDays := 2
	tStorageClass := "GLACIER"
	initiationDays := 1

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Lifecycle_Configuration_Transition_With_Abort_Incomplete_Multipart_Upload(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, transitionDays, tStorageClass, initiationDays),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.0.transition.0.days", "2"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.0.transition.0.storage_class", "GLACIER"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.0.abort_incomplete_multipart_upload.0.days_after_initiation", "1"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Lifecycle_Configuration_Expiration_With_Noncurrent_Version_Expiration(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform-lifecycle-configuration%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	expirationDays := 2
	nonCurrentDays := 1

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Lifecycle_Configuration_Expiration_With_Noncurrent_Version_Expiration(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, expirationDays, nonCurrentDays),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.0.expiration.0.days", "2"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.0.noncurrent_version_expiration.0.noncurrent_days", "1"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Lifecycle_Configuration_Expiration_With_Expired_Object_Delete_Marker(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform-lifecycle-configuration%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Lifecycle_Configuration_Expiration_With_Expired_Object_Delete_Marker(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.0.expiration.0.expired_object_delete_marker", "true"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Lifecycle_Configuration_Expiration_With_Expired_Object_Delete_Marker_On_Versioning_enable_Bucket(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform-lifecycle-configuration%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Lifecycle_Configuration_Expiration_With_Expired_Object_Delete_Marker_On_Versioning_Bucket(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.0.expiration.0.expired_object_delete_marker", "true"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Lifecycle_Configuration_Expiration_Transition_Abort_Incomplete_Multipart_Upload_All(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform-lifecycle-configuration%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	expirationDays := 3
	transitionDays := 2
	tStorageclass := "GLACIER"
	daysAfterInitiation := 1

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Lifecycle_Configuration_Expiration_Transition_Abort_Incomplete_Multipart_Upload(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, expirationDays, transitionDays, tStorageclass, daysAfterInitiation),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.0.expiration.0.days", "3"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.0.transition.0.days", "2"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.0.transition.0.storage_class", "GLACIER"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.0.abort_incomplete_multipart_upload.0.days_after_initiation", "1"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Lifecycle_Configuration_With_No_Filter(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform-lifecycle-configuration%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	expirationDays := 1

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCosBucket_Lifecycle_Configuration_With_No_Filter(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, expirationDays),
				ExpectError: regexp.MustCompile("Error: Insufficient filter blocks"),
			},
		},
	})
}

func TestAccIBMCosBucket_Lifecycle_Configuration_Status_Disabled(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform-lifecycle-configuration%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	expirationDays := 1

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Lifecycle_Configuration_Status_disable(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, expirationDays),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.0.expiration.0.days", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.0.status", "disable"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Lifecycle_Configuration_Multiple_Rules(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform-lifecycle-configuration%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	expirationDays1 := 2
	daysAfterInitiation1 := 1
	expirationDays2 := 4
	daysAfterInitiation2 := 3

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Lifecycle_Configuration_Multiple_Rules(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, expirationDays1, daysAfterInitiation1, expirationDays2, daysAfterInitiation2),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.#", "2"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.0.expiration.0.days", "2"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.0.abort_incomplete_multipart_upload.0.days_after_initiation", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.1.expiration.0.days", "4"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_lifecycle_configuration.lifecycle", "lifecycle_rule.1.abort_incomplete_multipart_upload.0.days_after_initiation", "3"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Lifecycle_Configuration_Expiration_With_Invalid_Days(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform-lifecycle-configuration%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	expirationDays := -1

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCosBucket_Lifecycle_Configuration_Expiration_With_Days(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, expirationDays),
				ExpectError: regexp.MustCompile(`"lifecycle_rule.0.expiration.0.days" must contain a valid int value should be in range\(1, 3650\), got -1`),
			},
		},
	})
}

func TestAccIBMCosBucket_Lifecycle_Configuration_Transition_With_Invalid_Days(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform-lifecycle-configuration%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	transitionDays := -1
	tStorageclass := "GLACIER"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCosBucket_Lifecycle_Configuration_Transition_With_Days(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, transitionDays, tStorageclass),
				ExpectError: regexp.MustCompile(`"lifecycle_rule.0.transition.0.days" must contain a valid int value should be in range\(0, 3650\), got -1`),
			},
		},
	})
}

func TestAccIBMCosBucket_Lifecycle_Configuration_Transition_With_Invalid_Storage_Class(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform-lifecycle-configuration%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	transitionDays := 1
	tStorageclass := "Invalid"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCosBucket_Lifecycle_Configuration_Transition_With_Days(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, transitionDays, tStorageclass),
				ExpectError: regexp.MustCompile(`"lifecycle_rule.0.transition.0.storage_class" must contain a value from \[\]string\{"GLACIER", "ACCELERATED"\}, got "Invalid`),
			},
		},
	})
}
func TestAccIBMCosBucket_Lifecycle_Configuration_Transition_Multiple_Rules(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform-lifecycle-configuration%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCosBucket_Lifecycle_Configuration_Transition_With_Multiple_Rules(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass),
				ExpectError: regexp.MustCompile("MalformedXML: The XML you provided was not well-formed or did not validate against our published schema."),
			},
		},
	})
}

func testAccCheckIBMCosBucket_Lifecycle_Configuration_Expiration_With_Days(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, days int) string {

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
	    region_location = "%s"
		storage_class         = "%s"
	}
	resource "ibm_cos_bucket_lifecycle_configuration"  "lifecycle" {
		bucket_crn = ibm_cos_bucket.bucket.crn
		bucket_location = ibm_cos_bucket.bucket.region_location
		lifecycle_rule {
		   expiration{
		     days = "%d"
		   }
		   filter {
		     prefix = "foo"
		   }  
		   rule_id = "id"
		   status = "enable"
	
		 }
	  }
	`, cosServiceName, bucketName, region, storageClass, days)
}

func testAccCheckIBMCosBucket_Lifecycle_Configuration_Expiration_With_Date(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, date string) string {

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
	    region_location = "%s"
		storage_class         = "%s"
	}
	resource "ibm_cos_bucket_lifecycle_configuration"  "lifecycle" {
		bucket_crn = ibm_cos_bucket.bucket.crn
		bucket_location = ibm_cos_bucket.bucket.region_location
		lifecycle_rule {
		   expiration{
		     date = "%s"
		   }
		   filter {
		     prefix = "foo"
		   }  
		   rule_id = "id"
		   status = "enable"
	
		 }
	  }
	`, cosServiceName, bucketName, region, storageClass, date)
}

func testAccCheckIBMCosBucket_Lifecycle_Configuration_Transition_With_Days(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, days int, tStorageclass string) string {

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
	    region_location = "%s"
		storage_class         = "%s"
	}
	resource "ibm_cos_bucket_lifecycle_configuration"  "lifecycle" {
		bucket_crn = ibm_cos_bucket.bucket.crn
		bucket_location = ibm_cos_bucket.bucket.region_location
		lifecycle_rule {
		   transition{
				days = "%d"
				storage_class = "%s"
			}
		   filter {
		     prefix = ""
		   }  
		   rule_id = "id"
		   status = "enable"
	
		 }
	  }	 
	`, cosServiceName, bucketName, region, storageClass, days, tStorageclass)
}

func testAccCheckIBMCosBucket_Lifecycle_Configuration_Transition_With_Date(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, date string, tStorageclass string) string {

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
	    region_location = "%s"
		storage_class         = "%s"
	}
	resource "ibm_cos_bucket_lifecycle_configuration"  "lifecycle" {
		bucket_crn = ibm_cos_bucket.bucket.crn
		bucket_location = ibm_cos_bucket.bucket.region_location
		lifecycle_rule {
		   transition{
				date = "%s"
				storage_class = "%s"
			}
		   filter {
		     prefix = ""
		   }  
		   rule_id = "id"
		   status = "enable"
	
		 }
	  }
	`, cosServiceName, bucketName, region, storageClass, date, tStorageclass)
}

func testAccCheckIBMCosBucket_Lifecycle_Configuration_Expiration_With_Abort_Incomplete_Multipart_Upload(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, days int, initiationDays int) string {

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
	    region_location = "%s"
		storage_class         = "%s"
	}
	resource "ibm_cos_bucket_lifecycle_configuration"  "lifecycle" {
		bucket_crn = ibm_cos_bucket.bucket.crn
		bucket_location = ibm_cos_bucket.bucket.region_location
		lifecycle_rule {
		   expiration{
		     days = "%d"
		   }
		   abort_incomplete_multipart_upload{
			   days_after_initiation = "%d"
		   }
		   filter {
		     prefix = "foo"
		   }  
		   rule_id = "id"
		   status = "enable"
	
		 }
	  }
	`, cosServiceName, bucketName, region, storageClass, days, initiationDays)
}

func testAccCheckIBMCosBucket_Lifecycle_Configuration_Abort_Incomplete_Multipart_Upload(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, initiationDays int) string {

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
	    region_location = "%s"
		storage_class         = "%s"
	}
	resource "ibm_cos_bucket_lifecycle_configuration"  "lifecycle" {
		bucket_crn = ibm_cos_bucket.bucket.crn
		bucket_location = ibm_cos_bucket.bucket.region_location
		lifecycle_rule {
		   abort_incomplete_multipart_upload{
			   days_after_initiation = "%d"
		   }
		   filter {
		     prefix = "foo"
		   }  
		   rule_id = "id"
		   status = "enable"
	
		 }
	  } 
	`, cosServiceName, bucketName, region, storageClass, initiationDays)
}

func testAccCheckIBMCosBucket_Lifecycle_Configuration_Expiration_With_Transition(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, days int, tDays int, tStorageClass string) string {

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
	    region_location = "%s"
		storage_class         = "%s"
	}
	resource "ibm_cos_bucket_lifecycle_configuration"  "lifecycle" {
		bucket_crn = ibm_cos_bucket.bucket.crn
		bucket_location = ibm_cos_bucket.bucket.region_location
		lifecycle_rule {
		   expiration{
		     days = "%d"
		   }
		   transition{
			days = "%d"
			storage_class = "%s"
		}
		   filter {
		     prefix = ""
		   }  
		   rule_id = "id"
		   status = "enable"
	
		 }
	  } 
	`, cosServiceName, bucketName, region, storageClass, days, tDays, tStorageClass)
}

func testAccCheckIBMCosBucket_Lifecycle_Configuration_Transition_With_Abort_Incomplete_Multipart_Upload(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, tDays int, tStorageClass string, initiationDays int) string {

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
	    region_location = "%s"
		storage_class         = "%s"
	}
	resource "ibm_cos_bucket_lifecycle_configuration"  "lifecycle" {
		bucket_crn = ibm_cos_bucket.bucket.crn
		bucket_location = ibm_cos_bucket.bucket.region_location
		lifecycle_rule {
		   transition{
			days = "%d"
			storage_class = "%s"
		}
		abort_incomplete_multipart_upload{
			days_after_initiation = "%d"
		}
		   filter {
		     prefix = ""
		   }  
		   rule_id = "id"
		   status = "enable"
	
		 }
	  }
	`, cosServiceName, bucketName, region, storageClass, tDays, tStorageClass, initiationDays)
}

func testAccCheckIBMCosBucket_Lifecycle_Configuration_Expiration_With_Noncurrent_Version_Expiration(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, days int, ncDays int) string {

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
	    region_location = "%s"
		storage_class         = "%s"
	}
	resource "ibm_cos_bucket_lifecycle_configuration"  "lifecycle" {
		bucket_crn = ibm_cos_bucket.bucket.crn
		bucket_location = ibm_cos_bucket.bucket.region_location
		lifecycle_rule {
		   expiration{
		     days = "%d"
		   }
		   noncurrent_version_expiration{
			   noncurrent_days = "%d"
		   }
		   filter {
		     prefix = "foo"
		   }  
		   rule_id = "id"
		   status = "enable"
	
		 }
	  }
	`, cosServiceName, bucketName, region, storageClass, days, ncDays)
}
func testAccCheckIBMCosBucket_Lifecycle_Configuration_Expiration_With_Expired_Object_Delete_Marker(cosServiceName string, bucketName string, regiontype string, region string, storageClass string) string {

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
	    region_location = "%s"
		storage_class         = "%s"
	}
	resource "ibm_cos_bucket_lifecycle_configuration"  "lifecycle" {
		bucket_crn = ibm_cos_bucket.bucket.crn
		bucket_location = ibm_cos_bucket.bucket.region_location
		lifecycle_rule {
		   expiration{
			expired_object_delete_marker = true
		   }
		   filter {
		     prefix = "foo"
		   }  
		   rule_id = "id"
		   status = "enable"
	
		 }
	  }
	`, cosServiceName, bucketName, region, storageClass)
}

func testAccCheckIBMCosBucket_Lifecycle_Configuration_Expiration_With_Expired_Object_Delete_Marker_On_Versioning_Bucket(cosServiceName string, bucketName string, regiontype string, region string, storageClass string) string {

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
	    region_location = "%s"
		storage_class         = "%s"
		object_versioning {
			     enable  = true
		   }
	}
	resource "ibm_cos_bucket_lifecycle_configuration"  "lifecycle" {
		bucket_crn = ibm_cos_bucket.bucket.crn
		bucket_location = ibm_cos_bucket.bucket.region_location
		lifecycle_rule {
		   expiration{
			expired_object_delete_marker = true
		   }
		   filter {
		     prefix = "foo"
		   }  
		   rule_id = "id"
		   status = "enable"
	
		 }
	  }
	`, cosServiceName, bucketName, region, storageClass)
}

func testAccCheckIBMCosBucket_Lifecycle_Configuration_Expiration_Transition_Abort_Incomplete_Multipart_Upload(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, days int, tDays int, tStorageClass string, initiationDays int) string {

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
	    region_location = "%s"
		storage_class         = "%s"
	}
	resource "ibm_cos_bucket_lifecycle_configuration"  "lifecycle" {
		bucket_crn = ibm_cos_bucket.bucket.crn
		bucket_location = ibm_cos_bucket.bucket.region_location
		lifecycle_rule {
		   expiration{
		     days = "%d"
		   }
		   transition{
			days = "%d"
			storage_class = "%s"
		}
		abort_incomplete_multipart_upload{
			days_after_initiation = "%d"
		}
		   filter {
		     prefix = ""
		   }  
		   rule_id = "id"
		   status = "enable"
	
		 }
	  }
	`, cosServiceName, bucketName, region, storageClass, days, tDays, tStorageClass, initiationDays)
}

func testAccCheckIBMCosBucket_Lifecycle_Configuration_With_No_Filter(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, days int) string {

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
	    region_location = "%s"
		storage_class         = "%s"
	}
	resource "ibm_cos_bucket_lifecycle_configuration"  "lifecycle" {
		bucket_crn = ibm_cos_bucket.bucket.crn
		bucket_location = ibm_cos_bucket.bucket.region_location
		lifecycle_rule {
		   expiration{
		     days = "%d"
		   }
		   rule_id = "id"
		   status = "enable"
	
		 }
	  }
	`, cosServiceName, bucketName, region, storageClass, days)
}

func testAccCheckIBMCosBucket_Lifecycle_Configuration_Status_disable(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, days int) string {

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
	    region_location = "%s"
		storage_class         = "%s"
	}
	resource "ibm_cos_bucket_lifecycle_configuration"  "lifecycle" {
		bucket_crn = ibm_cos_bucket.bucket.crn
		bucket_location = ibm_cos_bucket.bucket.region_location
		lifecycle_rule {
		   expiration{
		     days = "%d"
		   }
		   filter {
		     prefix = "foo"
		   }  
		   rule_id = "id"
		   status = "disable"
	
		 }
	  }
	`, cosServiceName, bucketName, region, storageClass, days)
}

func testAccCheckIBMCosBucket_Lifecycle_Configuration_Multiple_Rules(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, days1 int, initiationDays1 int, days2 int, initiationDays2 int) string {

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
	    region_location = "%s"
		storage_class         = "%s"
	}
	resource "ibm_cos_bucket_lifecycle_configuration"  "lifecycle" {
		bucket_crn = ibm_cos_bucket.bucket.crn
		bucket_location = ibm_cos_bucket.bucket.region_location
		lifecycle_rule {
		   expiration{
		     days = "%d"
		   }
		   abort_incomplete_multipart_upload{
			   days_after_initiation = "%d"
		   }
		   filter {
		     prefix = "foo"
		   }  
		   rule_id = "id1"
		   status = "enable"
	
		 }
		 lifecycle_rule {
			expiration{
			  days = "%d"
			}
			abort_incomplete_multipart_upload{
				days_after_initiation = "%d"
			}
			filter {
			  prefix = "bar"
			}  
			rule_id = "id2"
			status = "enable"
	 
		  }
	  }
	`, cosServiceName, bucketName, region, storageClass, days1, initiationDays1, days2, initiationDays2)
}

func testAccCheckIBMCosBucket_Lifecycle_Configuration_Transition_With_Multiple_Rules(cosServiceName string, bucketName string, regiontype string, region string, storageClass string) string {

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
		region_location = "%s"
		storage_class         = "%s"
	}
	resource "ibm_cos_bucket_lifecycle_configuration"  "lifecycle" {
		bucket_crn = ibm_cos_bucket.bucket.crn
		bucket_location = ibm_cos_bucket.bucket.region_location
		lifecycle_rule {
		   transition{
			 days = 1
			 storage_class = "GLACIER"
			}
		   filter {
		     prefix = ""
		   }  
		   rule_id = "id"
		   status = "enable"
		 }
		 lifecycle_rule {
			transition{
			 days = 3
			 storage_class = "ACCELERATED"
			 }
			filter {
			  prefix = ""
			}  
			rule_id = "id2"
			status = "enable"
	 
		  }
	  }
	`, cosServiceName, bucketName, region, storageClass)
}
