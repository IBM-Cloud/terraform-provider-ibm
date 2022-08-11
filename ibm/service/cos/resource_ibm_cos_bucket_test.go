// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cos_test

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"testing"

	"time"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/cos"

	"github.com/IBM/ibm-cos-sdk-go/aws"
	"github.com/IBM/ibm-cos-sdk-go/aws/credentials/ibmiam"
	token "github.com/IBM/ibm-cos-sdk-go/aws/credentials/ibmiam/token"
	"github.com/IBM/ibm-cos-sdk-go/aws/session"
	"github.com/IBM/ibm-cos-sdk-go/service/s3"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"gotest.tools/assert"
)

var singleSiteLocation = []string{
	"ams03", "che01", "hkg02", "mel01", "mex01",
	"mil01", "mon01", "osl01", "par01", "sjc04", "sao01",
	"seo01", "sng01", "tor01",
}

var regionLocation = []string{
	"au-syd", "ca-tor", "eu-de", "eu-gb", "jp-tok", "jp-osa", "us-east", "us-south", "br-sao",
}

var crossRegionLocation = []string{
	"us", "eu", "ap",
}

var storageClass = []string{
	"standard", "vault", "cold", "smart",
}

var singleSiteLocationRegex = regexp.MustCompile("^[a-z]{3}[0-9][0-9]-[a-z]{4,8}$")
var regionLocationRegex = regexp.MustCompile("^[a-z]{2}-[a-z]{2,5}[0-9]?-[a-z]{4,8}$")
var crossRegionLocationRegex = regexp.MustCompile("^[a-z]{2}-[a-z]{4,8}$")

func TestAccIBMCosBucket_Basic(t *testing.T) {

	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "eu"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_basic(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "cross_region_location", bucketRegion),
				),
			},
			{
				Config: testAccCheckIBMCosBucket_updateWithSameName(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "cross_region_location", bucketRegion),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Direct(t *testing.T) {

	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "eu"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_direct(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "cross_region_location", bucketRegion),
				),
			},
		},
	})
}
func TestAccIBMCosBucket_ActivityTracker_Monitor(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	activityServiceName := fmt.Sprintf("activity_tracker_%d", acctest.RandIntRange(10, 100))
	monitorServiceName := fmt.Sprintf("metrics_monitor_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "ams03"
	bucketClass := "standard"
	bucketRegionType := "single_site_location"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_activityTracker_monitor(cosServiceName, activityServiceName, monitorServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance2", "ibm_cos_bucket.bucket2", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "single_site_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMCosBucket_update_activityTracker_monitor(cosServiceName, activityServiceName, monitorServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance2", "ibm_cos_bucket.bucket2", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "single_site_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.#", "0"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.#", "0"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Archive_Expiration(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	arch_ruleId := "my-rule-id-bucket-arch"
	arch_enable := true
	archiveDays := 1
	ruleType := "GLACIER"
	exp_ruleId := "my-rule-id-bucket-expire"
	exp_enable := true
	expireDays := 1
	prefix := "prefix/"
	archDaysUpdate := 3
	expDaysUpdate := 2

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_archive_expire(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, arch_ruleId, arch_enable, archiveDays, ruleType, exp_ruleId, exp_enable, expireDays, prefix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "archive_rule.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "expire_rule.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMCosBucket_archive_expire_updateDays(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, arch_ruleId, arch_enable, archDaysUpdate, ruleType, exp_ruleId, exp_enable, expDaysUpdate, prefix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "archive_rule.0.days", fmt.Sprintf("%d", archDaysUpdate)),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "expire_rule.0.days", fmt.Sprintf("%d", expDaysUpdate)),
				),
			},
			{
				Config: testAccCheckIBMCosBucket_update_archive_expire(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, arch_ruleId, arch_enable, archiveDays, ruleType, exp_ruleId, exp_enable, expireDays, prefix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "archive_rule.#", "0"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "expire_rule.#", "0"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Archive(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	ruleId := "my-rule-id-bucket-arch"
	enable := true
	archiveDays := 1
	ruleType := "GLACIER"
	archiveDaysUpdate := 3

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_archive(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, ruleId, enable, archiveDays, ruleType),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "archive_rule.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMCosBucket_archive_updateDays(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, ruleId, enable, archiveDaysUpdate, ruleType),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "archive_rule.0.days", fmt.Sprintf("%d", archiveDaysUpdate)),
				),
			},
			{
				Config: testAccCheckIBMCosBucket_update_archive(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, ruleId, enable, archiveDays, ruleType),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "archive_rule.#", "0"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Expiredays(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	ruleId := "my-rule-id-bucket-expiredays"
	enable := true
	expireDays := 2
	prefix := "prefix/"
	expireDaysUpdate := 3

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_expiredays(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, ruleId, enable, expireDays, prefix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "expire_rule.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMCosBucket_expire_updateDays(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, ruleId, enable, expireDaysUpdate, prefix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "expire_rule.0.days", fmt.Sprintf("%d", expireDaysUpdate)),
				),
			},
			{
				Config: testAccCheckIBMCosBucket_update_expiredays(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, ruleId, enable, expireDays, prefix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "expire_rule.#", "0"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Expiredate(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	ruleId := "my-rule-id-bucket-expiredate"
	enable := true
	expireDate := "2021-11-28"
	prefix := ""
	expireDateUpdate := "2021-11-30"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_expiredate(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, ruleId, enable, expireDate, prefix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "expire_rule.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMCosBucket_expire_updateDate(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, ruleId, enable, expireDateUpdate, prefix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "expire_rule.0.date", expireDateUpdate),
				),
			},
			{
				Config: testAccCheckIBMCosBucket_update_expiredate(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, ruleId, enable, expireDate, prefix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "expire_rule.#", "0"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Expireddeletemarker(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	ruleId := "my-rule-id-bucket-expireddeletemarker"
	enable := true
	prefix := ""
	expiredObjectDeleteMarker := false

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_expiredeletemarker(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, ruleId, enable, expiredObjectDeleteMarker, prefix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "expire_rule.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_AbortIncompeleteMPU(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	ruleId := "my-rule-id-bucket-abortmpu"
	enable := true
	prefix := ""
	daysAfterInitiation := 1

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_abortincompletempu(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, ruleId, enable, daysAfterInitiation, prefix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "abort_incomplete_multipart_upload_days.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMCosBucket_update_abortincompletempu(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, ruleId, enable, daysAfterInitiation, prefix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "abort_incomplete_multipart_upload_days.#", "0"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_noncurrentversion(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	ruleId := "my-rule-id-bucket-ncversion"
	enable := true
	prefix := ""
	noncurrentDays := 1

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_noncurrentversion(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, ruleId, enable, noncurrentDays, prefix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "noncurrent_version_expiration.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMCosBucket_update_noncurrentversion(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, ruleId, enable, noncurrentDays, prefix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "noncurrent_version_expiration.#", "0"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Retention(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "jp-tok"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	default_retention := 0
	maximum_retention := 1
	minimum_retention := 0

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_retention(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, default_retention, maximum_retention, minimum_retention),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "retention_rule.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Object_Versioning(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-east"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	enable := true

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_object_versioning(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, enable),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "object_versioning.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Hard_Quota(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	hardQuota := 1024

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_hard_quota(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, hardQuota),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "hard_quota", fmt.Sprintf("%d", hardQuota)),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Smart_Type(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "eu"
	bucketClass := "smart"
	bucketRegionType := "cross_region_location"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_basic(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "cross_region_location", bucketRegion),
				),
			},
			{
				Config: testAccCheckIBMCosBucket_updateWithSameName(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "cross_region_location", bucketRegion),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_import(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "eu"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_basic(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "cross_region_location", bucketRegion),
				),
			},
			{
				ResourceName:      "ibm_cos_bucket.bucket",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"wait_time_minutes", "parameters", "force_delete"},
			},
		},
	})
}

//
// Satellite location
func TestAccIBMCosBucket_Satellite(t *testing.T) {

	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := acc.Satellite_location_id
	ResourceInstanceId := acc.Satellite_Resource_instance_id

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_basic_sat(serviceName, bucketName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucket_Satellite_Exists(ResourceInstanceId, "ibm_cos_bucket.bucket", bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
				),
			},
			{
				Config: testAccCheckIBMCosBucket_sat_updateWithSameName(serviceName, bucketName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucket_Satellite_Exists(ResourceInstanceId, "ibm_cos_bucket.bucket", bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Satellite_Expiredays(t *testing.T) {

	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := acc.Satellite_location_id
	ResourceInstanceId := acc.Satellite_Resource_instance_id
	ruleId := "my-rule-id-bucket-expiredays"
	enable := true
	expireDays := 2
	prefix := "prefix/"
	expireDaysUpdate := 3

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_expiredays_satellite(bucketName, bucketRegion, ruleId, enable, expireDays, prefix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucket_Satellite_Exists(ResourceInstanceId, "ibm_cos_bucket.bucket", bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "expire_rule.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMCosBucket_expire_updateDays_satellite(bucketName, bucketRegion, ruleId, enable, expireDaysUpdate, prefix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucket_Satellite_Exists(ResourceInstanceId, "ibm_cos_bucket.bucket", bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "expire_rule.0.days", fmt.Sprintf("%d", expireDaysUpdate)),
				),
			},
			{
				Config: testAccCheckIBMCosBucket_update_expiredays_satellite(bucketName, bucketRegion, ruleId, enable, expireDays, prefix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucket_Satellite_Exists(ResourceInstanceId, "ibm_cos_bucket.bucket", bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "expire_rule.#", "0"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Satellite_Expiredate(t *testing.T) {

	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := acc.Satellite_location_id
	ResourceInstanceId := acc.Satellite_Resource_instance_id
	ruleId := "my-rule-id-bucket-expiredate"
	enable := true
	expireDate := "2021-11-28"
	prefix := ""
	expireDateUpdate := "2021-11-30"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_sat_expiredate(bucketName, bucketRegion, ruleId, enable, expireDate, prefix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucket_Satellite_Exists(ResourceInstanceId, "ibm_cos_bucket.bucket", bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "expire_rule.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMCosBucket_sat_expire_updateDate(bucketName, bucketRegion, ruleId, enable, expireDateUpdate, prefix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucket_Satellite_Exists(ResourceInstanceId, "ibm_cos_bucket.bucket", bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "expire_rule.0.date", expireDateUpdate),
				),
			},
			{
				Config: testAccCheckIBMCosBucket_sat_update_expiredate(bucketName, bucketRegion, ruleId, enable, expireDate, prefix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucket_Satellite_Exists(ResourceInstanceId, "ibm_cos_bucket.bucket", bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "expire_rule.#", "0"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Satellite_Expireddeletemarker(t *testing.T) {

	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := acc.Satellite_location_id
	ResourceInstanceId := acc.Satellite_Resource_instance_id
	ruleId := "my-rule-id-bucket-expireddeletemarker"
	enable := true
	prefix := ""
	expiredObjectDeleteMarker := false

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_sat_expiredeletemarker(bucketName, bucketRegion, ruleId, enable, expiredObjectDeleteMarker, prefix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucket_Satellite_Exists(ResourceInstanceId, "ibm_cos_bucket.bucket", bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "expire_rule.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Satellite_AbortIncompeleteMPU(t *testing.T) {

	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := acc.Satellite_location_id
	ResourceInstanceId := acc.Satellite_Resource_instance_id
	ruleId := "my-rule-id-bucket-abortmpu"
	enable := true
	prefix := ""
	daysAfterInitiation := 1

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_sat_abortincompletempu(bucketName, bucketRegion, ruleId, enable, daysAfterInitiation, prefix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucket_Satellite_Exists(ResourceInstanceId, "ibm_cos_bucket.bucket", bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "abort_incomplete_multipart_upload_days.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMCosBucket_sat_update_abortincompletempu(bucketName, bucketRegion, ruleId, enable, daysAfterInitiation, prefix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucket_Satellite_Exists(ResourceInstanceId, "ibm_cos_bucket.bucket", bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "abort_incomplete_multipart_upload_days.#", "0"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Satellite_noncurrentversion(t *testing.T) {

	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := acc.Satellite_location_id
	ResourceInstanceId := acc.Satellite_Resource_instance_id
	ruleId := "my-rule-id-bucket-ncversion"
	enable := true
	prefix := ""
	noncurrentDays := 1

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_sat_noncurrentversion(bucketName, bucketRegion, ruleId, enable, noncurrentDays, prefix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucket_Satellite_Exists(ResourceInstanceId, "ibm_cos_bucket.bucket", bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "noncurrent_version_expiration.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMCosBucket_sat_update_noncurrentversion(bucketName, bucketRegion, ruleId, enable, noncurrentDays, prefix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucket_Satellite_Exists(ResourceInstanceId, "ibm_cos_bucket.bucket", bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "noncurrent_version_expiration.#", "0"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Satellite_Object_Versioning(t *testing.T) {

	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := acc.Satellite_location_id
	ResourceInstanceId := acc.Satellite_Resource_instance_id
	enable := true

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_sat_object_versioning(bucketName, bucketRegion, enable),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucket_Satellite_Exists(ResourceInstanceId, "ibm_cos_bucket.bucket", bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),

					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "object_versioning.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMCosBucketDestroy(s *terraform.State) error {

	var s3Conf *aws.Config
	var apiEndpoint string
	var resourceInstance string
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "ibm_cos_bucket" {
			apiEndpoint = rs.Primary.Attributes["s3_endpoint_public"]
		}
		if rs.Type == "ibm_resource_instance" && rs.Primary.Attributes["service"] == "cloud-object-storage" {
			resourceInstance = rs.Primary.Attributes["crn"]

		}
	}

	rsContClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BluemixSession()
	if err != nil {
		return err
	}

	authEndpoint, err := rsContClient.Config.EndpointLocator.IAMEndpoint()
	if err != nil {
		return err
	}
	authEndpointPath := fmt.Sprintf("%s%s", authEndpoint, "/identity/token")
	apiKey := rsContClient.Config.BluemixAPIKey
	if apiKey != "" {
		s3Conf = aws.NewConfig().WithEndpoint(apiEndpoint).WithCredentials(ibmiam.NewStaticCredentials(aws.NewConfig(), authEndpointPath, apiKey, resourceInstance)).WithS3ForcePathStyle(true)
	}
	iamAccessToken := rsContClient.Config.IAMAccessToken
	if iamAccessToken != "" {
		initFunc := func() (*token.Token, error) {
			return &token.Token{
				AccessToken:  rsContClient.Config.IAMAccessToken,
				RefreshToken: rsContClient.Config.IAMRefreshToken,
				TokenType:    "Bearer",
				ExpiresIn:    int64((time.Hour * 248).Seconds()) * -1,
				Expiration:   time.Now().Add(-1 * time.Hour).Unix(),
			}, nil
		}
		s3Conf = aws.NewConfig().WithEndpoint(apiEndpoint).WithCredentials(ibmiam.NewCustomInitFuncCredentials(aws.NewConfig(), initFunc, authEndpointPath, resourceInstance)).WithS3ForcePathStyle(true)
	}
	s3Sess := session.Must(session.NewSession())
	s3Client := s3.New(s3Sess, s3Conf)

	bucketList, _ := s3Client.ListBuckets(&s3.ListBucketsInput{})
	if len(bucketList.Buckets) > 0 {
		return errors.New("Bucket still exists")

	}
	return nil
}

// COS Satellite
func testAccCheckIBMCosBucket_Satellite_Exists(resource string, bucket string, region string, bucketname string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		var s3Conf *aws.Config

		bucket, ok := s.RootModule().Resources[bucket]
		if !ok {
			return fmt.Errorf("Bucket Not found: %s", bucket)
		}

		satloc_guid := strings.Split(resource, ":")
		bucketsatcrn := satloc_guid[7]
		resource = bucketsatcrn

		var rt string

		rt = "sl"

		apiEndpoint := cos.SelectSatlocCosApi(rt, resource, region)

		rsContClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BluemixSession()
		if err != nil {
			return err
		}

		authEndpoint, err := rsContClient.Config.EndpointLocator.IAMEndpoint()
		if err != nil {
			return err
		}
		authEndpointPath := fmt.Sprintf("%s%s", authEndpoint, "/identity/token")
		apiKey := rsContClient.Config.BluemixAPIKey
		if apiKey != "" {
			s3Conf = aws.NewConfig().WithEndpoint(apiEndpoint).WithCredentials(ibmiam.NewStaticCredentials(aws.NewConfig(), authEndpointPath, apiKey, resource)).WithS3ForcePathStyle(true)
		}
		iamAccessToken := rsContClient.Config.IAMAccessToken
		if iamAccessToken != "" {
			initFunc := func() (*token.Token, error) {
				return &token.Token{
					AccessToken:  rsContClient.Config.IAMAccessToken,
					RefreshToken: rsContClient.Config.IAMRefreshToken,
					TokenType:    "Bearer",
					ExpiresIn:    int64((time.Hour * 248).Seconds()) * -1,
					Expiration:   time.Now().Add(-1 * time.Hour).Unix(),
				}, nil
			}
			s3Conf = aws.NewConfig().WithEndpoint(apiEndpoint).WithCredentials(ibmiam.NewCustomInitFuncCredentials(aws.NewConfig(), initFunc, authEndpointPath, resource)).WithS3ForcePathStyle(true)
		}
		s3Sess := session.Must(session.NewSession())
		s3Client := s3.New(s3Sess, s3Conf)

		bucketList, _ := s3Client.ListBuckets(&s3.ListBucketsInput{})
		for _, bucket := range bucketList.Buckets {
			bn := *bucket.Name
			if bn == bucketname {
				return nil
			}
		}
		return errors.New("bucket does not exist")
	}
}

/// IBMCLOUD
func testAccCheckIBMCosBucketExists(resource string, bucket string, regiontype string, region string, bucketname string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		var s3Conf *aws.Config
		resourceInstance, ok := s.RootModule().Resources[resource]

		if !ok {
			return fmt.Errorf("Not found: %s", resource)
		}

		bucket, ok := s.RootModule().Resources[bucket]

		if !ok {
			return fmt.Errorf("Not found: %s", bucket)
		}

		var rt string
		if regiontype == "single_site_location" {
			rt = "ssl"
		}
		if regiontype == "region_location" {
			rt = "rl"
		}
		if regiontype == "cross_region_location" {
			rt = "crl"
		}

		apiEndpoint, _, _ := cos.SelectCosApi(rt, region)

		rsContClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BluemixSession()
		if err != nil {
			return err
		}

		authEndpoint, err := rsContClient.Config.EndpointLocator.IAMEndpoint()
		if err != nil {
			return err
		}
		authEndpointPath := fmt.Sprintf("%s%s", authEndpoint, "/identity/token")
		apiKey := rsContClient.Config.BluemixAPIKey
		if apiKey != "" {
			s3Conf = aws.NewConfig().WithEndpoint(apiEndpoint).WithCredentials(ibmiam.NewStaticCredentials(aws.NewConfig(), authEndpointPath, apiKey, resourceInstance.Primary.ID)).WithS3ForcePathStyle(true)
		}
		iamAccessToken := rsContClient.Config.IAMAccessToken
		if iamAccessToken != "" {
			initFunc := func() (*token.Token, error) {
				return &token.Token{
					AccessToken:  rsContClient.Config.IAMAccessToken,
					RefreshToken: rsContClient.Config.IAMRefreshToken,
					TokenType:    "Bearer",
					ExpiresIn:    int64((time.Hour * 248).Seconds()) * -1,
					Expiration:   time.Now().Add(-1 * time.Hour).Unix(),
				}, nil
			}
			s3Conf = aws.NewConfig().WithEndpoint(apiEndpoint).WithCredentials(ibmiam.NewCustomInitFuncCredentials(aws.NewConfig(), initFunc, authEndpointPath, resourceInstance.Primary.ID)).WithS3ForcePathStyle(true)
		}
		s3Sess := session.Must(session.NewSession())
		s3Client := s3.New(s3Sess, s3Conf)

		bucketList, _ := s3Client.ListBuckets(&s3.ListBucketsInput{})
		for _, bucket := range bucketList.Buckets {
			bn := *bucket.Name
			if bn == bucketname {
				return nil
			}
		}
		return errors.New("bucket does not exist")
	}
}

func testAccCheckIBMCosBucket_basic(serviceName string, bucketName string, regiontype string, region string, storageClass string) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "group" {
		is_default=true
	}
	  
	resource "ibm_resource_instance" "instance" {
		name              = "%s"
		service           = "cloud-object-storage"
		plan              = "standard"
		location          = "global"
		resource_group_id = data.ibm_resource_group.group.id
	}
	  
	resource "ibm_cos_bucket" "bucket" {
		bucket_name          = "%s"
		resource_instance_id = ibm_resource_instance.instance.id
		storage_class        = "%s"
		cross_region_location = "%s"
	}
	  
		  
	`, serviceName, bucketName, storageClass, region)
}

func testAccCheckIBMCosBucket_direct(serviceName string, bucketName string, regiontype string, region string, storageClass string) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "group" {
		is_default=true
	}
	  
	resource "ibm_resource_instance" "instance" {
		name              = "%s"
		service           = "cloud-object-storage"
		plan              = "standard"
		location          = "global"
		resource_group_id = data.ibm_resource_group.group.id
	}
	  
	resource "ibm_cos_bucket" "bucket" {
		bucket_name          = "%s"
		resource_instance_id = ibm_resource_instance.instance.id
		storage_class        = "%s"
		cross_region_location = "%s"
		endpoint_type= "direct"
	}
	  
		  
	`, serviceName, bucketName, storageClass, region)
}
func testAccCheckIBMCosBucket_updateWithSameName(serviceName string, bucketName string, regiontype string, region, storageClass string) string {

	return fmt.Sprintf(`	
	data "ibm_resource_group" "group" {
		is_default=true
	}
	  
	resource "ibm_resource_instance" "instance" {
		name              = "%s"
		service           = "cloud-object-storage"
		plan              = "standard"
		location          = "global"
		resource_group_id = data.ibm_resource_group.group.id
	}
	  
	resource "ibm_cos_bucket" "bucket" {
		bucket_name          = "%s"
		resource_instance_id = ibm_resource_instance.instance.id
		storage_class        = "%s"
		cross_region_location = "%s"
	}
	`, serviceName, bucketName, storageClass, region)
}

func testAccCheckIBMCosBucket_activityTracker_monitor(cosServiceName, activityServiceName, monitorServiceName, bucketName, regiontype, region, storageClass string) string {

	return fmt.Sprintf(`

	data "ibm_resource_group" "cos_group" {
		is_default=true
	  }
	  resource "ibm_resource_instance" "instance2" {
		name              = "%s"
		resource_group_id = data.ibm_resource_group.cos_group.id
		service           = "cloud-object-storage"
		plan              = "standard"
		location          = "global"
	  }
	  resource "ibm_resource_instance" "activity_tracker2" {
		name              = "%s"
		resource_group_id = data.ibm_resource_group.cos_group.id
		service           = "logdnaat"
		plan              = "7-day"
		location          = "us-south"
	  }
	  resource "ibm_resource_instance" "metrics_monitor2" {
		name              = "%s"
		resource_group_id = data.ibm_resource_group.cos_group.id
		service           = "sysdig-monitor"
		plan              = "graduated-tier"
		location          = "us-south"
		parameters        = {
			default_receiver = true
		}
	  }
	  resource "ibm_cos_bucket" "bucket2" {
		bucket_name          = "%s"
		resource_instance_id = ibm_resource_instance.instance2.id
		single_site_location = "%s"
		storage_class        = "%s"
		activity_tracking {
		  read_data_events     = true
		  write_data_events    = true
		  activity_tracker_crn = ibm_resource_instance.activity_tracker2.id
		}
		metrics_monitoring {
		  usage_metrics_enabled  = true
		  request_metrics_enabled = true
		  metrics_monitoring_crn = ibm_resource_instance.metrics_monitor2.id
		}
	  }  
	`, cosServiceName, activityServiceName, monitorServiceName, bucketName, region, storageClass)
}

func testAccCheckIBMCosBucket_update_activityTracker_monitor(cosServiceName, activityServiceName, monitorServiceName, bucketName, regiontype, region, storageClass string) string {

	return fmt.Sprintf(`	
	data "ibm_resource_group" "cos_group" {
		is_default=true
	}
	  
	resource "ibm_resource_instance" "instance2" {
		name              = "%s"
		resource_group_id = data.ibm_resource_group.cos_group.id
		service           = "cloud-object-storage"
		plan              = "standard"
		location          = "global"
	  }
	  
	resource "ibm_resource_instance" "activity_tracker2" {
		name              = "%s"
		resource_group_id = data.ibm_resource_group.cos_group.id
		service           = "logdnaat"
		plan              = "7-day"
		location          = "us-south"
	}
	  
	resource "ibm_resource_instance" "metrics_monitor2" {
		name              = "%s"
		resource_group_id = data.ibm_resource_group.cos_group.id
		service           = "sysdig-monitor"
		plan              = "graduated-tier"
		location          = "us-south"
		parameters        = {
			default_receiver = true
		}
	}
	resource "ibm_cos_bucket" "bucket2" {
		bucket_name          = "%s"
		resource_instance_id = ibm_resource_instance.instance2.id
		single_site_location = "%s"
		storage_class        = "%s"
	}	  
	`, cosServiceName, activityServiceName, monitorServiceName, bucketName, region, storageClass)
}

func testAccCheckIBMCosBucket_archive(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, ruleId string, enable bool, archiveDays int, ruleType string) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		is_default=true
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
	    region_location       = "%s"
		storage_class         = "%s"
		archive_rule {
		  rule_id             = "%s"
		  enable              = true
		  days                = %d
		  type                = "%s"
		}
	}
	`, cosServiceName, bucketName, region, storageClass, ruleId, archiveDays, ruleType)
}

func testAccCheckIBMCosBucket_archive_updateDays(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, ruleId string, enable bool, archiveDaysUpdate int, ruleType string) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		is_default=true
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
	    region_location       = "%s"
		storage_class         = "%s"
		archive_rule {
			rule_id             = "%s"
			enable              = true
			days                = %d
			type                = "%s"
		}
	}
	`, cosServiceName, bucketName, region, storageClass, ruleId, archiveDaysUpdate, ruleType)
}

func testAccCheckIBMCosBucket_update_archive(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, ruleId string, enable bool, archiveDays int, ruleType string) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		is_default=true
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
	    region_location       = "%s"
		storage_class         = "%s"
	}
	`, cosServiceName, bucketName, region, storageClass)
}

func testAccCheckIBMCosBucket_expiredays(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, ruleId string, enable bool, expireDays int, prefix string) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		is_default=true
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
	    region_location       = "%s"
		storage_class         = "%s"
		expire_rule {
			rule_id             = "%s"
			enable              = true
			days                = %d
			prefix              = "%s"
		}
	}
	`, cosServiceName, bucketName, region, storageClass, ruleId, expireDays, prefix)
}

func testAccCheckIBMCosBucket_expire_updateDays(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, ruleId string, enable bool, expireDaysUpdate int, prefix string) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		is_default=true
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
	    region_location       = "%s"
		storage_class         = "%s"
		expire_rule {
			rule_id             = "%s"
			enable              = true
			days                = %d
			prefix              = "%s"
		  }
	}
	`, cosServiceName, bucketName, region, storageClass, ruleId, expireDaysUpdate, prefix)

}

func testAccCheckIBMCosBucket_update_expiredays(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, ruleId string, enable bool, expireDays int, prefix string) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		is_default=true
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
	    region_location       = "%s"
		storage_class         = "%s"
	}
	`, cosServiceName, bucketName, region, storageClass)
}

func testAccCheckIBMCosBucket_expiredate(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, ruleId string, enable bool, expireDate string, prefix string) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		is_default=true
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
	    region_location       = "%s"
		storage_class         = "%s"
		expire_rule {
			rule_id             = "%s"
			enable              = true
			date                = "%s"
			prefix              = "%s"
		}
	}
	`, cosServiceName, bucketName, region, storageClass, ruleId, expireDate, prefix)
}

func testAccCheckIBMCosBucket_expire_updateDate(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, ruleId string, enable bool, expireDateUpdate string, prefix string) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		is_default=true
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
	    region_location       = "%s"
		storage_class         = "%s"
		expire_rule {
			rule_id             = "%s"
			enable              = true
			date                = "%s"
			prefix              = "%s"
		}
	}
	`, cosServiceName, bucketName, region, storageClass, ruleId, expireDateUpdate, prefix)

}

func testAccCheckIBMCosBucket_update_expiredate(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, ruleId string, enable bool, expireDate string, prefix string) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		is_default=true
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
	    region_location       = "%s"
		storage_class         = "%s"
	}
	`, cosServiceName, bucketName, region, storageClass)
}

func testAccCheckIBMCosBucket_expiredeletemarker(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, ruleId string, enable bool, expiredObjectDeleteMarker bool, prefix string) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		is_default=true
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
	    region_location       = "%s"
		storage_class         = "%s"
		expire_rule {
			rule_id                     = "%s"
			enable                      = true
			expired_object_delete_marker = true
			prefix                       = "%s"
		}
	}
	`, cosServiceName, bucketName, region, storageClass, ruleId, prefix)
}

func testAccCheckIBMCosBucket_archive_expire(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, arch_ruleId string, arch_enable bool, archiveDays int, ruleType string, exp_ruleId string, exp_enable bool, expireDays int, prefix string) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		is_default=true
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
	    region_location       = "%s"
		storage_class         = "%s"
		archive_rule {
			rule_id             = "%s"
			enable              = true
			days                = %d
			type                = "%s"
		}
		expire_rule {
			rule_id             = "%s"
			enable              = true
			days                = %d
			prefix              = "%s"
		  }
	}
	`, cosServiceName, bucketName, region, storageClass, arch_ruleId, archiveDays, ruleType, exp_ruleId, expireDays, prefix)

}
func testAccCheckIBMCosBucket_archive_expire_updateDays(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, arch_ruleId string, arch_enable bool, archDaysUpdate int, ruleType string, exp_ruleId string, exp_enable bool, expDaysUpdate int, prefix string) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		is_default=true
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
	    region_location       = "%s"
		storage_class         = "%s"
		archive_rule {
			rule_id             = "%s"
			enable              = true
			days                = %d
			type                = "%s"
		}
		expire_rule {
			rule_id             = "%s"
			enable              = true
			days                = %d
			prefix              = "%s"
		  }
	}
	`, cosServiceName, bucketName, region, storageClass, arch_ruleId, archDaysUpdate, ruleType, exp_ruleId, expDaysUpdate, prefix)

}

func testAccCheckIBMCosBucket_update_archive_expire(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, arch_ruleId string, arch_enable bool, archiveDays int, ruleType string, exp_ruleId string, exp_enable bool, expireDays int, prefix string) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		is_default=true
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
	    region_location       = "%s"
		storage_class         = "%s"
	}
	`, cosServiceName, bucketName, region, storageClass)
}

func testAccCheckIBMCosBucket_retention(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, default_retention int, maximum_retention int, minimum_retention int) string {

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
	    region_location       = "%s"
		storage_class         = "%s"
		retention_rule {
			default = %d
			maximum = %d
			minimum = %d
			permanent = false
		}
	}
	`, cosServiceName, bucketName, region, storageClass, default_retention, maximum_retention, minimum_retention)
}

func testAccCheckIBMCosBucket_object_versioning(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, enable bool) string {

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
	    region_location       = "%s"
		storage_class         = "%s"
		object_versioning {
			enable  = true
		}
	}
	`, cosServiceName, bucketName, region, storageClass)
}

func testAccCheckIBMCosBucket_hard_quota(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, hardQuota int) string {

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
	    region_location       = "%s"
		storage_class         = "%s"
		hard_quota			  = %d
	}
	`, cosServiceName, bucketName, region, storageClass, hardQuota)
}

func testAccCheckIBMCosBucket_abortincompletempu(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, ruleId string, enable bool, daysAfterInitiation int, prefix string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		is_default=true
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
	    region_location       = "%s"
		storage_class         = "%s"
		abort_incomplete_multipart_upload_days {
			rule_id             	= "%s"
			enable              	= true
			days_after_initiation   = %d
			prefix                  = "%s"
		}
	}
	`, cosServiceName, bucketName, region, storageClass, ruleId, daysAfterInitiation, prefix)
}

func testAccCheckIBMCosBucket_abortincompletempu_updateDays(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, ruleId string, enable bool, daysAfterInitiationUpdate int, prefix string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		is_default=true
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
	    region_location       = "%s"
		storage_class         = "%s"
		abort_incomplete_multipart_upload_days {
			rule_id            	    = "%s"
			enable              	= true
			days_after_initiation   = %d
			prefix                  = "%s"
		  }
	}
	`, cosServiceName, bucketName, region, storageClass, ruleId, daysAfterInitiationUpdate, prefix)

}

func testAccCheckIBMCosBucket_update_abortincompletempu(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, ruleId string, enable bool, daysAfterInitiation int, prefix string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		is_default=true
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
	    region_location       = "%s"
		storage_class         = "%s"
	}
	`, cosServiceName, bucketName, region, storageClass)
}

func testAccCheckIBMCosBucket_noncurrentversion(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, ruleId string, enable bool, noncurrentDays int, prefix string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		is_default=true
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
	    region_location       = "%s"
		storage_class         = "%s"
		noncurrent_version_expiration {
			rule_id             	= "%s"
			enable              	= true
			noncurrent_days  		= %d
			prefix                  = "%s"
		}
	}
	`, cosServiceName, bucketName, region, storageClass, ruleId, noncurrentDays, prefix)
}

func testAccCheckIBMCosBucket_update_noncurrentversion(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, ruleId string, enable bool, noncurrentDays int, prefix string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		is_default=true
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
	    region_location       = "%s"
		storage_class         = "%s"
	}
	`, cosServiceName, bucketName, region, storageClass)
}

//Satellite

func testAccCheckIBMCosBucket_basic_sat(serviceName string, bucketName string) string {
	return fmt.Sprintf(`
	
	resource "ibm_cos_bucket" "bucket" {
		bucket_name          = "%s"
		resource_instance_id = "%s"
		satellite_location_id = "%s"
	}
	  
		  
	`, bucketName, acc.Satellite_Resource_instance_id, acc.Satellite_location_id)
}
func testAccCheckIBMCosBucket_sat_updateWithSameName(serviceName string, bucketName string) string {

	return fmt.Sprintf(`	

	resource "ibm_cos_bucket" "bucket" {
		bucket_name          = "%s"
		resource_instance_id = "%s"
		satellite_location_id = "%s"
	}
	`, bucketName, acc.Satellite_Resource_instance_id, acc.Satellite_location_id)
}

func testAccCheckIBMCosBucket_sat_expiredeletemarker(bucketName string, region string, ruleId string, enable bool, expiredObjectDeleteMarker bool, prefix string) string {

	return fmt.Sprintf(`
	resource "ibm_cos_bucket" "bucket" {
		bucket_name          = "%s"
		resource_instance_id = "%s"
		satellite_location_id = "%s"
		expire_rule {
			rule_id                     = "%s"
			enable                      = true
			expired_object_delete_marker = true
			prefix                       = "%s"
		}
	}
	`, bucketName, acc.Satellite_Resource_instance_id, acc.Satellite_location_id, ruleId, prefix)
}

func testAccCheckIBMCosBucket_sat_abortincompletempu(bucketName string, region string, ruleId string, enable bool, daysAfterInitiation int, prefix string) string {
	return fmt.Sprintf(`
	resource "ibm_cos_bucket" "bucket" {
		bucket_name          = "%s"
		resource_instance_id = "%s"
		satellite_location_id = "%s"
		abort_incomplete_multipart_upload_days {
			rule_id             	= "%s"
			enable              	= true
			days_after_initiation   = %d
			prefix                  = "%s"
		}
	}
	`, bucketName, acc.Satellite_Resource_instance_id, acc.Satellite_location_id, ruleId, daysAfterInitiation, prefix)
}

func testAccCheckIBMCosBucket_sat_abortincompletempu_updateDays(bucketName string, region string, ruleId string, enable bool, daysAfterInitiationUpdate int, prefix string) string {
	return fmt.Sprintf(`
	resource "ibm_cos_bucket" "bucket" {
		bucket_name          = "%s"
		resource_instance_id = "%s"
		satellite_location_id = "%s"
		abort_incomplete_multipart_upload_days {
			rule_id            	    = "%s"
			enable              	= true
			days_after_initiation   = %d
			prefix                  = "%s"
		  }
	}
	`, bucketName, acc.Satellite_Resource_instance_id, acc.Satellite_location_id, ruleId, daysAfterInitiationUpdate, prefix)

}

func testAccCheckIBMCosBucket_sat_update_abortincompletempu(bucketName string, region string, ruleId string, enable bool, daysAfterInitiation int, prefix string) string {
	return fmt.Sprintf(`
	resource "ibm_cos_bucket" "bucket" {
		bucket_name          = "%s"
		resource_instance_id = "%s"
		satellite_location_id = "%s"
	}
	`, bucketName, acc.Satellite_Resource_instance_id, acc.Satellite_location_id)
}

func testAccCheckIBMCosBucket_sat_noncurrentversion(bucketName string, region string, ruleId string, enable bool, noncurrentDays int, prefix string) string {
	return fmt.Sprintf(`
	resource "ibm_cos_bucket" "bucket" {
		bucket_name          = "%s"
		resource_instance_id = "%s"
		satellite_location_id = "%s"
		noncurrent_version_expiration {
			rule_id             	= "%s"
			enable              	= true
			noncurrent_days  		= %d
			prefix                  = "%s"
		}
	}
	`, bucketName, acc.Satellite_Resource_instance_id, acc.Satellite_location_id, ruleId, noncurrentDays, prefix)
}

func testAccCheckIBMCosBucket_sat_update_noncurrentversion(bucketName string, region string, ruleId string, enable bool, noncurrentDays int, prefix string) string {
	return fmt.Sprintf(`
	resource "ibm_cos_bucket" "bucket" {
		bucket_name          = "%s"
		resource_instance_id = "%s"
		satellite_location_id = "%s"
	}
	`, bucketName, acc.Satellite_Resource_instance_id, acc.Satellite_location_id)
}

func testAccCheckIBMCosBucket_sat_expiredate(bucketName string, region string, ruleId string, enable bool, expireDate string, prefix string) string {

	return fmt.Sprintf(`
	resource "ibm_cos_bucket" "bucket" {
		bucket_name          = "%s"
		resource_instance_id = "%s"
		satellite_location_id = "%s"
		expire_rule {
			rule_id             = "%s"
			enable              = true
			date                = "%s"
			prefix              = "%s"
		}
	}
	`, bucketName, acc.Satellite_Resource_instance_id, acc.Satellite_location_id, ruleId, expireDate, prefix)
}

func testAccCheckIBMCosBucket_sat_expire_updateDate(bucketName string, region string, ruleId string, enable bool, expireDateUpdate string, prefix string) string {

	return fmt.Sprintf(`
	resource "ibm_cos_bucket" "bucket" {
		bucket_name          = "%s"
		resource_instance_id = "%s"
		satellite_location_id = "%s"
		expire_rule {
			rule_id             = "%s"
			enable              = true
			date                = "%s"
			prefix              = "%s"
		}
	}
	`, bucketName, acc.Satellite_Resource_instance_id, acc.Satellite_location_id, ruleId, expireDateUpdate, prefix)

}

func testAccCheckIBMCosBucket_sat_update_expiredate(bucketName string, region string, ruleId string, enable bool, expireDate string, prefix string) string {

	return fmt.Sprintf(`
	resource "ibm_cos_bucket" "bucket" {
		bucket_name          = "%s"
		resource_instance_id = "%s"
		satellite_location_id = "%s"
	}
	`, bucketName, acc.Satellite_Resource_instance_id, acc.Satellite_location_id)
}

func testAccCheckIBMCosBucket_sat_object_versioning(bucketName string, region string, enable bool) string {

	return fmt.Sprintf(`
	
	resource "ibm_cos_bucket" "bucket" {
		bucket_name          = "%s"
		resource_instance_id = "%s"
		satellite_location_id = "%s"
		object_versioning {
			enable  = true
		}
	}
	`, bucketName, acc.Satellite_Resource_instance_id, acc.Satellite_location_id)
}

func testAccCheckIBMCosBucket_expiredays_satellite(bucketName string, region string, ruleId string, enable bool, expireDays int, prefix string) string {

	return fmt.Sprintf(`
	
	resource "ibm_cos_bucket" "bucket" {
		bucket_name          = "%s"
		resource_instance_id = "%s"
		satellite_location_id = "%s"
		expire_rule {
			rule_id             = "%s"
			enable              = true
			days                = %d
			prefix              = "%s"
		}
	}

	`, bucketName, acc.Satellite_Resource_instance_id, acc.Satellite_location_id, ruleId, expireDays, prefix)
}

func testAccCheckIBMCosBucket_expire_updateDays_satellite(bucketName string, region string, ruleId string, enable bool, expireDaysUpdate int, prefix string) string {

	return fmt.Sprintf(`
	resource "ibm_cos_bucket" "bucket" {
		bucket_name          = "%s"
		resource_instance_id = "%s"
		satellite_location_id = "%s"
		expire_rule {
			rule_id             = "%s"
			enable              = true
			days                = %d
			prefix              = "%s"
		}
	}

	`, bucketName, acc.Satellite_Resource_instance_id, acc.Satellite_location_id, ruleId, expireDaysUpdate, prefix)

}

func testAccCheckIBMCosBucket_update_expiredays_satellite(bucketName string, region string, ruleId string, enable bool, expireDays int, prefix string) string {

	return fmt.Sprintf(`
	
	resource "ibm_cos_bucket" "bucket" {
		bucket_name          = "%s"
		resource_instance_id = "%s"
		satellite_location_id = "%s"
	}
	`, bucketName, acc.Satellite_Resource_instance_id, acc.Satellite_location_id)
}

func TestSingleSiteLocationRegex(t *testing.T) {
	var re = singleSiteLocationRegex
	for _, singleSite := range singleSiteLocation {
		for _, sc := range storageClass {
			assert.Equal(t, re.MatchString(singleSite+"-"+sc), true)
		}
	}

	for _, region := range regionLocation {
		assert.Equal(t, re.MatchString(region+"-standard"), false)
	}

	for _, crossRegion := range crossRegionLocation {
		assert.Equal(t, re.MatchString(crossRegion+"-standard"), false)
	}
}

func TestRegionLocationRegex(t *testing.T) {
	var re = regionLocationRegex
	for _, singleSite := range singleSiteLocation {
		assert.Equal(t, re.MatchString(singleSite+"-standard"), false)
	}

	for _, region := range regionLocation {
		for _, sc := range storageClass {
			assert.Equal(t, re.MatchString(region+"-"+sc), true)
			// test numeric suffix
			assert.Equal(t, re.MatchString(region+"2-"+sc), true)
		}
	}

	for _, crossRegion := range crossRegionLocation {
		assert.Equal(t, re.MatchString(crossRegion+"-standard"), false)
	}
}

func TestCrossRegionLocationRegex(t *testing.T) {
	var re = crossRegionLocationRegex
	for _, singleSite := range singleSiteLocation {
		assert.Equal(t, re.MatchString(singleSite+"-standard"), false)
	}

	for _, region := range regionLocation {
		assert.Equal(t, re.MatchString(region+"-standard"), false)
	}

	for _, crossRegion := range crossRegionLocation {
		for _, sc := range storageClass {
			assert.Equal(t, re.MatchString(crossRegion+"-"+sc), true)
		}
	}
}
