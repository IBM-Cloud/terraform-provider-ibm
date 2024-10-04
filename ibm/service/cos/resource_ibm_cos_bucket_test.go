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
	"standard", "vault", "cold", "smart", "onerate_active",
}

var singleSiteLocationRegex = regexp.MustCompile("^[a-z]{3}[0-9][0-9]-[a-z_a-z]{4,14}$")
var regionLocationRegex = regexp.MustCompile("^[a-z]{2}-[a-z]{2,5}[0-9]?-[a-z_a-z]{4,14}$")
var crossRegionLocationRegex = regexp.MustCompile("^[a-z]{2}-[a-z_a-z]{4,14}$")

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

func TestAccIBMCosBucket_Basic_Single_Site_Location(t *testing.T) {

	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "ams03"
	bucketClass := "standard"
	bucketRegionType := "single_site_location"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_basic_ssl(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "single_site_location", bucketRegion),
				),
			},
		},
	})
}
func TestAccIBMCosBucket_AllowedIP(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"
	allowedIp1 := "103.208.71.79"
	allowedIp2 := "172.30.8.121"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_allowedip(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, allowedIp1, allowedIp2),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "cross_region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "allowed_ip.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMCosBucket_allowedipremoved(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "cross_region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "allowed_ip.#", "0"),
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

// *** F1881 Activity tracker test cases  ***
func TestAccIBMCosBucket_ActivityTracker_Read_True_Write_False_ActivityTrackerCrn_NotSet_ManagementEvents_False(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	readDataEvents := true
	writeDataEvents := false
	managementEvents := false

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_activityTracker_Without_Crn(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, readDataEvents, writeDataEvents, managementEvents),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance2", "ibm_cos_bucket.bucket2", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.read_data_events", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.write_data_events", "false"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.management_events", "false"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_ActivityTracker_Read_False_Write_True_ActivityTrackerCrn_NotSet_ManagementEvents_False(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	readDataEvents := false
	writeDataEvents := true
	managementEvents := false

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_activityTracker_Without_Crn(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, readDataEvents, writeDataEvents, managementEvents),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance2", "ibm_cos_bucket.bucket2", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.read_data_events", "false"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.write_data_events", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.management_events", "false"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_ActivityTracker_Read_False_Write_False_ActivityTrackerCrn_NotSet_ManagementEvents_True(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	readDataEvents := false
	writeDataEvents := false
	managementEvents := true

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_activityTracker_Without_Crn(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, readDataEvents, writeDataEvents, managementEvents),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance2", "ibm_cos_bucket.bucket2", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.read_data_events", "false"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.write_data_events", "false"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.management_events", "true"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_ActivityTracker_Read_True_Write_True_ActivityTrackerCrn_NotSet_ManagementEvents_False(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	readDataEvents := true
	writeDataEvents := true
	managementEvents := false

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_activityTracker_Without_Crn(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, readDataEvents, writeDataEvents, managementEvents),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance2", "ibm_cos_bucket.bucket2", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.read_data_events", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.write_data_events", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.management_events", "false"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_ActivityTracker_Read_False_Write_False_ActivityTrackerCrn_NotSet_ManagementEvents_False(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	readDataEvents := false
	writeDataEvents := false
	managementEvents := false

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_activityTracker_Without_Crn(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, readDataEvents, writeDataEvents, managementEvents),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance2", "ibm_cos_bucket.bucket2", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.read_data_events", "false"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.write_data_events", "false"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.management_events", "false"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_ActivityTracker_Read_True_Write_False_ActivityTrackerCrn_NotSet_ManagementEvents_True(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	readDataEvents := true
	writeDataEvents := false
	managementEvents := true

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_activityTracker_Without_Crn(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, readDataEvents, writeDataEvents, managementEvents),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance2", "ibm_cos_bucket.bucket2", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.read_data_events", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.write_data_events", "false"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.management_events", "true"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_ActivityTracker_Read_False_Write_True_ActivityTrackerCrn_NotSet_ManagementEvents_True(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	readDataEvents := false
	writeDataEvents := true
	managementEvents := true

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_activityTracker_Without_Crn(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, readDataEvents, writeDataEvents, managementEvents),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance2", "ibm_cos_bucket.bucket2", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.read_data_events", "false"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.write_data_events", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.management_events", "true"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_ActivityTracker_Read_True_Write_True_ActivityTrackerCrn_NotSet_ManagementEvents_True(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	readDataEvents := true
	writeDataEvents := true
	managementEvents := true

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_activityTracker_Without_Crn(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, readDataEvents, writeDataEvents, managementEvents),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance2", "ibm_cos_bucket.bucket2", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.read_data_events", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.write_data_events", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.management_events", "true"),
				),
			},
		},
	})
}

// *** with crn ***

func TestAccIBMCosBucket_ActivityTracker_Read_True_Write_False_ActivityTrackerCrn_Set_ManagementEvents_Not_Set(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	readDataEvents := true
	writeDataEvents := false
	activityTrackerInstanceCRN := acc.ActivityTrackerInstanceCRN
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_activityTracker_With_Crn_ManagementEvents_NotSet(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, activityTrackerInstanceCRN, readDataEvents, writeDataEvents),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance2", "ibm_cos_bucket.bucket2", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.read_data_events", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.write_data_events", "false"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.activity_tracker_crn", activityTrackerInstanceCRN),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_ActivityTracker_Read_False_Write_True_ActivityTrackerCrn_Set_ManagementEvents_Not_Set(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	readDataEvents := false
	writeDataEvents := true
	activityTrackerInstanceCRN := acc.ActivityTrackerInstanceCRN
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_activityTracker_With_Crn_ManagementEvents_NotSet(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, activityTrackerInstanceCRN, readDataEvents, writeDataEvents),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance2", "ibm_cos_bucket.bucket2", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.read_data_events", "false"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.write_data_events", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.management_events", "false"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.activity_tracker_crn", activityTrackerInstanceCRN),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_ActivityTracker_Read_True_Write_True_ActivityTrackerCrn_Set_ManagementEvents_Not_Set(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	readDataEvents := true
	writeDataEvents := true
	activityTrackerInstanceCRN := acc.ActivityTrackerInstanceCRN

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_activityTracker_With_Crn_ManagementEvents_NotSet(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, activityTrackerInstanceCRN, readDataEvents, writeDataEvents),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance2", "ibm_cos_bucket.bucket2", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.read_data_events", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.write_data_events", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.management_events", "false"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.activity_tracker_crn", activityTrackerInstanceCRN),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_ActivityTracker_Read_False_Write_False_ActivityTrackerCrn_Set_ManagementEvents_Not_Set(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	readDataEvents := false
	writeDataEvents := false
	activityTrackerInstanceCRN := acc.ActivityTrackerInstanceCRN
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_activityTracker_With_Crn_ManagementEvents_NotSet(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, activityTrackerInstanceCRN, readDataEvents, writeDataEvents),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance2", "ibm_cos_bucket.bucket2", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.read_data_events", "false"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.write_data_events", "false"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.management_events", "false"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.activity_tracker_crn", activityTrackerInstanceCRN),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_ActivityTracker_Read_True_Write_False_ActivityTrackerCrn_Set_ManagementEvents_True(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	readDataEvents := true
	writeDataEvents := false
	managementEvents := true
	activityTrackerInstanceCRN := acc.ActivityTrackerInstanceCRN
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_activityTracker_With_Crn(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, activityTrackerInstanceCRN, readDataEvents, writeDataEvents, managementEvents),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance2", "ibm_cos_bucket.bucket2", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.read_data_events", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.write_data_events", "false"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.management_events", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.activity_tracker_crn", activityTrackerInstanceCRN),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_ActivityTracker_Read_False_Write_True_ActivityTrackerCrn_Set_ManagementEvents_True(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	readDataEvents := false
	writeDataEvents := true
	managementEvents := true
	activityTrackerInstanceCRN := acc.ActivityTrackerInstanceCRN
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_activityTracker_With_Crn(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, activityTrackerInstanceCRN, readDataEvents, writeDataEvents, managementEvents),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance2", "ibm_cos_bucket.bucket2", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.read_data_events", "false"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.write_data_events", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.management_events", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.activity_tracker_crn", activityTrackerInstanceCRN),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_ActivityTracker_Read_True_Write_True_ActivityTrackerCrn_Set_ManagementEvents_True(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	readDataEvents := true
	writeDataEvents := true
	managementEvents := true
	activityTrackerInstanceCRN := acc.ActivityTrackerInstanceCRN
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_activityTracker_With_Crn(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, activityTrackerInstanceCRN, readDataEvents, writeDataEvents, managementEvents),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance2", "ibm_cos_bucket.bucket2", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.read_data_events", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.write_data_events", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.management_events", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.activity_tracker_crn", activityTrackerInstanceCRN),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_ActivityTracker_Read_False_Write_False_ActivityTrackerCrn_Set_ManagementEvents_True(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	readDataEvents := false
	writeDataEvents := false
	managementEvents := true
	activityTrackerInstanceCRN := acc.ActivityTrackerInstanceCRN
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_activityTracker_With_Crn(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, activityTrackerInstanceCRN, readDataEvents, writeDataEvents, managementEvents),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance2", "ibm_cos_bucket.bucket2", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.read_data_events", "false"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.write_data_events", "false"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.management_events", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.activity_tracker_crn", activityTrackerInstanceCRN),
				),
			},
		},
	})
}
func TestAccIBMCosBucket_ActivityTracker_Read_False_Write_False_ManagementEvents_False_ActivityTrackerCrn_Set(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	readDataEvents := false
	writeDataEvents := false
	managementEvents := false
	activityTrackerInstanceCRN := acc.ActivityTrackerInstanceCRN
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCosBucket_activityTracker_With_Crn(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, activityTrackerInstanceCRN, readDataEvents, writeDataEvents, managementEvents),
				ExpectError: regexp.MustCompile("Error Update COS Bucket: Cannot have an Activity Tracking CRN without opting for management events"),
			},
		},
	})
}
func TestAccIBMCosBucket_Upload_Object_Activity_Tracker_Enabled_With_CRN(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	key := fmt.Sprintf("tf-testacc-cos-%d", acctest.RandIntRange(10, 100))
	bucketRegionType := "region_location"
	objectBody := "Acceptance Testing"
	activityTrackerInstanceCRN := acc.ActivityTrackerInstanceCRN
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Upload_Object_Activity_Tracker_Enabled_With_CRN(cosServiceName, activityTrackerInstanceCRN, bucketName, bucketRegionType, bucketRegion, bucketClass, key, objectBody),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance2", "ibm_cos_bucket.bucket2", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.read_data_events", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.write_data_events", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.management_events", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.activity_tracker_crn", activityTrackerInstanceCRN),
					resource.TestCheckResourceAttr("ibm_cos_bucket_object.testacc", "body", objectBody),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Upload_Object_Activity_Tracker_Enabled_Without_CRN(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	key := fmt.Sprintf("tf-testacc-cos-%d", acctest.RandIntRange(10, 100))
	bucketRegionType := "region_location"
	objectBody := "Acceptance Testing"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_upload_Object_Activity_Tracker_Enabled_Without_CRN(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, key, objectBody),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance2", "ibm_cos_bucket.bucket2", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.read_data_events", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.write_data_events", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "activity_tracking.0.management_events", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_object.testacc", "body", objectBody),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_ActivityTracker_Read_Invalid_Write_Invalid_ManagementEvents_Invalid_With_Crn(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	activityTrackerInstanceCRN := acc.ActivityTrackerInstanceCRN
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCosBucket_activityTracker_Read_Invalid_Write_Invalid_ManagementEvents_Invalid_With_Crn(cosServiceName, activityTrackerInstanceCRN, bucketName, bucketRegionType, bucketRegion, bucketClass),
				ExpectError: regexp.MustCompile("Error: Incorrect attribute value type"),
			},
		},
	})
}

func TestAccIBMCosBucket_ActivityTracker_Read_True_Write_True_ManagementEvents_True_With_Crn_Invalid(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	readDataEvents := true
	writeDataEvents := true
	managementEvents := true
	crnValue := "Invalid"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCosBucket_activityTracker_With_Crn(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, crnValue, readDataEvents, writeDataEvents, managementEvents),
				ExpectError: regexp.MustCompile("Error Update COS Bucket: Malformed activity tracker CRN"),
			},
		},
	})
}

//Metrics monitoring test cases:

func TestAccIBMCosBucket_MetricsMonitoring_RequestEnabled_True_UsageEnabled_False_MonitoringCrn_NotSet(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	requestEnabled := true
	usageEnabled := false

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_metricsMonitoring_Without_Crn(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, requestEnabled, usageEnabled),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance2", "ibm_cos_bucket.bucket2", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.0.request_metrics_enabled", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.0.usage_metrics_enabled", "false"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_MetricsMonitoring_RequestEnabled_False_UsageEnabled_True_MonitoringCrn_NotSet(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	requestEnabled := false
	usageEnabled := true

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_metricsMonitoring_Without_Crn(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, requestEnabled, usageEnabled),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance2", "ibm_cos_bucket.bucket2", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.0.request_metrics_enabled", "false"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.0.usage_metrics_enabled", "true"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_MetricsMonitoring_RequestEnabled_True_UsageEnabled_True_MonitoringCrn_NotSet(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	requestEnabled := true
	usageEnabled := true

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_metricsMonitoring_Without_Crn(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, requestEnabled, usageEnabled),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance2", "ibm_cos_bucket.bucket2", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.0.request_metrics_enabled", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.0.usage_metrics_enabled", "true"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_MetricsMonitoring_RequestEnabled_False_UsageEnabled_False_MonitoringCrn_NotSet(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	requestEnabled := false
	usageEnabled := false

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_metricsMonitoring_Without_Crn(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, requestEnabled, usageEnabled),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance2", "ibm_cos_bucket.bucket2", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.0.request_metrics_enabled", "false"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.0.usage_metrics_enabled", "false"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_MetricsMonitoring_RequestEnabled_True_UsageEnabled_False_MonitoringCrn_Set(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	requestEnabled := true
	usageEnabled := false
	metricsMonitoringCrn := acc.MetricsMonitoringCRN

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_metricsMonitoring_With_Crn(cosServiceName, metricsMonitoringCrn, bucketName, bucketRegionType, bucketRegion, bucketClass, requestEnabled, usageEnabled),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance2", "ibm_cos_bucket.bucket2", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.0.request_metrics_enabled", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.0.usage_metrics_enabled", "false"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.0.metrics_monitoring_crn", metricsMonitoringCrn),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_MetricsMonitoring_RequestEnabled_False_UsageEnabled_True_MonitoringCrn_Set(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	requestEnabled := false
	usageEnabled := true
	metricsMonitoringCrn := acc.MetricsMonitoringCRN
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_metricsMonitoring_With_Crn(cosServiceName, metricsMonitoringCrn, bucketName, bucketRegionType, bucketRegion, bucketClass, requestEnabled, usageEnabled),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance2", "ibm_cos_bucket.bucket2", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.0.request_metrics_enabled", "false"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.0.usage_metrics_enabled", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.0.metrics_monitoring_crn", metricsMonitoringCrn),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_MetricsMonitoring_RequestEnabled_True_UsageEnabled_True_MonitoringCrn_Set(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	requestEnabled := true
	usageEnabled := true
	metricsMonitoringCrn := acc.MetricsMonitoringCRN
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_metricsMonitoring_With_Crn(cosServiceName, metricsMonitoringCrn, bucketName, bucketRegionType, bucketRegion, bucketClass, requestEnabled, usageEnabled),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance2", "ibm_cos_bucket.bucket2", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.0.request_metrics_enabled", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.0.usage_metrics_enabled", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.0.metrics_monitoring_crn", metricsMonitoringCrn),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_MetricsMonitoring_RequestEnabled_False_UsageEnabled_False_MonitoringCrn_Set(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	requestEnabled := false
	usageEnabled := false
	metricsMonitoringCrn := acc.MetricsMonitoringCRN
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_metricsMonitoring_With_Crn(cosServiceName, metricsMonitoringCrn, bucketName, bucketRegionType, bucketRegion, bucketClass, requestEnabled, usageEnabled),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance2", "ibm_cos_bucket.bucket2", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.0.request_metrics_enabled", "false"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.0.usage_metrics_enabled", "false"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.0.metrics_monitoring_crn", metricsMonitoringCrn),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_MetricsMonitoring_Upload_Object_RequestMetrics_True_UsageMetrics_True_MonitoringCrn_Set(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	key := fmt.Sprintf("tf-testacc-cos-%d", acctest.RandIntRange(10, 100))
	objectBody := "Acceptance Testing"
	metricsMonitoringCrn := acc.MetricsMonitoringCRN
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_metricsMonitoring_Upload_Object_RequestMetrics_True_UsageMetrics_True_With_Crn(cosServiceName, metricsMonitoringCrn, bucketName, bucketRegionType, bucketRegion, bucketClass, key, objectBody),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance2", "ibm_cos_bucket.bucket2", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.0.request_metrics_enabled", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.0.usage_metrics_enabled", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.0.metrics_monitoring_crn", metricsMonitoringCrn),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_MetricsMonitoring_Upload_Object_RequestMetrics_True_UsageMetrics_True_Without_Crn(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	key := fmt.Sprintf("tf-testacc-cos-%d", acctest.RandIntRange(10, 100))
	objectBody := "Acceptance Testing"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_metricsMonitoring_Upload_Object_RequestMetrics_True_UsageMetrics_True_Without_Crn(cosServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass, key, objectBody),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance2", "ibm_cos_bucket.bucket2", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.0.request_metrics_enabled", "true"),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.0.usage_metrics_enabled", "true"),
				),
			},
		},
	})
}
func TestAccIBMCosBucket_MetricsMonitoring_RequestMetrics_Invalid_UsageMetrics_Invalid_MonitoringCrn_Set(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	metricsMonitoringCrn := acc.MetricsMonitoringCRN
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCosBucket_metricsMonitoring_RequestMetrics_Invalid_UsageMetrics_Invalid_With_Crn(cosServiceName, metricsMonitoringCrn, bucketName, bucketRegionType, bucketRegion, bucketClass),
				ExpectError: regexp.MustCompile("Error: Incorrect attribute value type"),
			},
		},
	})
}

func TestAccIBMCosBucket_MetricsMonitoring_Crn_Invalid(t *testing.T) {

	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"
	metricsMonitoringCrn := acc.MetricsMonitoringCRN
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCosBucket_metricsMonitoring_Crn_Invalid(cosServiceName, metricsMonitoringCrn, bucketName, bucketRegionType, bucketRegion, bucketClass),
				ExpectError: regexp.MustCompile("Error Update COS Bucket: Malformed Monitoring CRN."),
			},
		},
	})
}

// func TestAccIBMCosBucket_MetricsMonitoring_New_Instance(t *testing.T) {

// 	cosServiceName := fmt.Sprintf("cos_instance_%d", acctest.RandIntRange(10, 100))
// 	bucketName := fmt.Sprintf("tf-bucket%d", acctest.RandIntRange(10, 100))
// 	bucketRegion := "us-south"
// 	bucketClass := "standard"
// 	bucketRegionType := "region_location"
// 	monitorServiceName := fmt.Sprintf("metrics_monitor_%d", acctest.RandIntRange(10, 100))
// 	resource.Test(t, resource.TestCase{
// 		PreCheck:     func() { acc.TestAccPreCheck(t) },
// 		Providers:    acc.TestAccProviders,
// 		CheckDestroy: testAccCheckIBMCosBucketDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccCheckIBMCosBucket_metricsMonitoring_RequestMetrics_New_Instance(cosServiceName, monitorServiceName, bucketName, bucketRegionType, bucketRegion, bucketClass),
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance2", "ibm_cos_bucket.bucket2", bucketRegionType, bucketRegion, bucketName),
// 					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "bucket_name", bucketName),
// 					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "storage_class", bucketClass),
// 					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "region_location", bucketRegion),
// 					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.#", "1"),
// 					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.0.request_metrics_enabled", "true"),
// 					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket2", "metrics_monitoring.0.usage_metrics_enabled", "true"),
// 				),
// 			},
// 		},
// 	})
// }

// *** f1881 tests cases end ***

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

func TestAccIBMCosBucket_Retention_Rule_Existing_bucket(t *testing.T) {
	cosCrn := acc.CosCRN
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
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
				Config: testAccCheckIBMCosBucket_retention_basic_bucket(bucketName, cosCrn, bucketRegionType, bucketRegion, bucketClass),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
				),
			},
			{
				Config: testAccCheckIBMCosBucket_retention_existing_bucket(bucketName, cosCrn, bucketRegionType, bucketRegion, bucketClass, default_retention, maximum_retention, minimum_retention),
				Check: resource.ComposeAggregateTestCheckFunc(
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

func TestAccIBMCosBucket_OneRate_With_Storageclass(t *testing.T) {

	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "onerate_active"
	bucketRegionType := "region_location"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Onerate_With_Storageclass(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_OneRate_Without_Storage_class(t *testing.T) {

	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketRegionType := "region_location"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Onerate_Without_Storage_class(serviceName, bucketName, bucketRegionType, bucketRegion),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "region_location", bucketRegion),
				),
			},
		},
	})
}
func TestAccIBMCosBucket_OneRate_With_Invalid_Storageclass(t *testing.T) {

	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "invalidstorageclass"
	bucketRegionType := "region_location"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCosBucket_Onerate_With_Invalid_Storageclass(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass),
				ExpectError: regexp.MustCompile("\"storage_class\" must contain a value from \\[\\]string{\"standard\", \"vault\", \"cold\", \"smart\", \"flex\", \"onerate_active\"}, got \"invalidstorageclass\""),
			},
		},
	})
}

func TestAccIBMCosBucket_COS_Plan_Storageclass_Mismatch_Type1(t *testing.T) {

	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "region_location"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCosBucket_COS_Plan_Storageclass_Mismatch_Type1(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass),
				ExpectError: regexp.MustCompile("InvalidLocationConstraint: Storage class not allowed for one rate user"),
			},
		},
	})
}
func TestAccIBMCosBucket_COS_Plan_Storageclass_Mismatch_Type2(t *testing.T) {

	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "onerate_active"
	bucketRegionType := "region_location"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCosBucket_COS_Plan_Storageclass_Mismatch_Type2(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass),
				ExpectError: regexp.MustCompile("InvalidLocationConstraint: Storage class not allowed for standard or cloud lite user"),
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

// / IBMCLOUD
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

		apiEndpoint, _, _ := cos.SelectCosApi(rt, region, false)

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
func TestAccIBMCOSKP(t *testing.T) {
	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKeyProtectRootkeyWithCOSBucket(instanceName, keyName, serviceName, bucketName, bucketRegion, bucketClass),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_kms_key.test", "key_name", keyName),
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
				),
			},
		},
	})
}

func TestAccIBMCOSHPCS(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMHPCSRootkeyWithCOSBucket(keyName, serviceName, bucketName, bucketRegion, bucketClass),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "key_protect", acc.HpcsRootKeyCrn),
				),
			},
		},
	})
}

// new hpcs
func TestAccIBMCOSKPKmsParamValid(t *testing.T) {

	instanceName := fmt.Sprintf("kms_%d", acctest.RandIntRange(10, 100))
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMKeyProtectRootkeyWithCOSBucketKmsParam(instanceName, keyName, serviceName, bucketName, bucketRegion, bucketClass),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
				),
			},
		},
	})
}
func TestAccIBMCOSKPKmsParamWithInvalidCRN(t *testing.T) {

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
				Config:      testAccCheckIBMKeyProtectRootkeyWithCOSBucketKmsParamWithInvalidCRN(instanceName, keyName, serviceName, bucketName, bucketRegion, bucketClass),
				ExpectError: regexp.MustCompile("InvalidArgument: Invalid ibm-sse-kp-customer-root-key-crn: received only 7 of required 10 segments"),
			},
		},
	})
}

func TestAccIBMCOSHPCSKmsParam(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMHPCSRootkeyWithCOSBucketKmsParam(keyName, serviceName, bucketName, bucketRegion, bucketClass),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "kms_key_crn", acc.HpcsRootKeyCrn),
				),
			},
		},
	})
}

func TestAccIBMCOSHPCSKmsParamWithInvalidCRN(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMHPCSRootkeyWithCOSBucketKmsParamWithInvalidCRN(keyName, serviceName, bucketName, bucketRegion, bucketClass),
				ExpectError: regexp.MustCompile("InvalidArgument: Invalid ibm-sse-kp-customer-root-key-crn: received only 7 of required 10 segments"),
			},
		},
	})
}

func TestAccIBMCOSKMSBothParamProvided(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraform%d", acctest.RandIntRange(10, 100))
	keyName := fmt.Sprintf("key_%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us-south"
	bucketClass := "standard"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMHPCSRootkeyWithCOSBucketKMSBothParamProvided(keyName, serviceName, bucketName, bucketRegion, bucketClass),
				ExpectError: regexp.MustCompile("Error: Conflicting configuration arguments"),
			},
		},
	})
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

func testAccCheckIBMCosBucket_basic_ssl(serviceName string, bucketName string, regiontype string, region string, storageClass string) string {

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
		single_site_location = "%s"
	}
	`, serviceName, bucketName, storageClass, region)
}

func testAccCheckIBMCosBucket_Onerate_With_Storageclass(serviceName string, bucketName string, regiontype string, region string, storageClass string) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "group" {
		is_default=true
	}
	  
	resource "ibm_resource_instance" "instance" {
		name              = "%s"
		service           = "cloud-object-storage"
		plan              = "cos-one-rate-plan"
		location          = "global"
		resource_group_id = data.ibm_resource_group.group.id
	}
	  
	resource "ibm_cos_bucket" "bucket" {
		bucket_name          = "%s"
		resource_instance_id = ibm_resource_instance.instance.id
		storage_class        = "%s"
		region_location = "%s"
	}
	  
		  
	`, serviceName, bucketName, storageClass, region)
}

func testAccCheckIBMCosBucket_Onerate_Without_Storage_class(serviceName string, bucketName string, regiontype string, region string) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "group" {
		is_default=true
	}

	resource "ibm_resource_instance" "instance" {
		name              = "%s"
		service           = "cloud-object-storage"
		plan              = "cos-one-rate-plan"
		location          = "global"
		resource_group_id = data.ibm_resource_group.group.id
	}

	resource "ibm_cos_bucket" "bucket" {
		bucket_name          = "%s"
		resource_instance_id = ibm_resource_instance.instance.id
		region_location = "%s"
	}

	`, serviceName, bucketName, region)
}

func testAccCheckIBMCosBucket_Onerate_With_Invalid_Storageclass(serviceName string, bucketName string, regiontype string, region string, storageClass string) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "group" {
		is_default=true
	}

	resource "ibm_resource_instance" "instance" {
		name              = "%s"
		service           = "cloud-object-storage"
		plan              = "cos-one-rate-plan"
		location          = "global"
		resource_group_id = data.ibm_resource_group.group.id
	}

	resource "ibm_cos_bucket" "bucket" {
		bucket_name          = "%s"
		resource_instance_id = ibm_resource_instance.instance.id
		storage_class        = "%s"
		region_location = "%s"
	}

	`, serviceName, bucketName, storageClass, region)
}

func testAccCheckIBMCosBucket_COS_Plan_Storageclass_Mismatch_Type1(serviceName string, bucketName string, regiontype string, region string, storageClass string) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "group" {
		is_default=true
	}

	resource "ibm_resource_instance" "instance" {
		name              = "%s"
		service           = "cloud-object-storage"
		plan              = "cos-one-rate-plan"
		location          = "global"
		resource_group_id = data.ibm_resource_group.group.id
	}

	resource "ibm_cos_bucket" "bucket" {
		bucket_name          = "%s"
		resource_instance_id = ibm_resource_instance.instance.id
		storage_class        = "%s"
		region_location = "%s"
	}

	`, serviceName, bucketName, storageClass, region)
}
func testAccCheckIBMCosBucket_COS_Plan_Storageclass_Mismatch_Type2(serviceName string, bucketName string, regiontype string, region string, storageClass string) string {

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
		region_location = "%s"
	}

	`, serviceName, bucketName, storageClass, region)
}

func testAccCheckIBMCosBucket_allowedip(serviceName string, bucketName string, regiontype string, region string, storageClass string, allowedIp1 string, allowedIp2 string) string {

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
		allowed_ip = ["%s","%s"]
	}
	  
	`, serviceName, bucketName, storageClass, region, allowedIp1, allowedIp2)
}

func testAccCheckIBMCosBucket_allowedipremoved(serviceName string, bucketName string, regiontype string, region string, storageClass string) string {

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

// f1881

func testAccCheckIBMCosBucket_activityTracker_Without_Crn(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, readDataEvents bool, writeDataEvents bool, managementEvents bool) string {

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
	  resource "ibm_cos_bucket" "bucket2" {
		bucket_name          = "%s"
		resource_instance_id = ibm_resource_instance.instance2.id
		region_location = "%s"
		storage_class        = "%s"
		activity_tracking {
		  read_data_events      = %v
		  write_data_events     = %v
		  management_events     = %v
		}
	  }  
	`, cosServiceName, bucketName, region, storageClass, readDataEvents, writeDataEvents, managementEvents)
}

func testAccCheckIBMCosBucket_activityTracker_With_Crn(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, activityTrackerCRN string, readDataEvents bool, writeDataEvents bool, managementEvents bool) string {

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
	  resource "ibm_cos_bucket" "bucket2" {
		bucket_name          = "%s"
		resource_instance_id = ibm_resource_instance.instance2.id
		region_location = "%s"
		storage_class        = "%s"
		activity_tracking {
		  read_data_events      = %v
		  write_data_events     = %v
		  management_events     = %v
		  activity_tracker_crn = "%s"
		}
	  }  
	`, cosServiceName, bucketName, region, storageClass, readDataEvents, writeDataEvents, managementEvents, activityTrackerCRN)
}

func testAccCheckIBMCosBucket_activityTracker_With_Crn_ManagementEvents_NotSet(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, activityTrackerCRN string, readDataEvents bool, writeDataEvents bool) string {

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
	  resource "ibm_cos_bucket" "bucket2" {
		bucket_name          = "%s"
		resource_instance_id = ibm_resource_instance.instance2.id
		region_location = "%s"
		storage_class        = "%s"
		activity_tracking {
		  read_data_events      = %v
		  write_data_events     = %v
		  activity_tracker_crn = "%s"
		}
	  }  
	`, cosServiceName, bucketName, region, storageClass, readDataEvents, writeDataEvents, activityTrackerCRN)
}

func testAccCheckIBMCosBucket_activityTracker_Read_Invalid_Write_Invalid_ManagementEvents_Invalid_With_Crn(cosServiceName, activityTrackerCRN, bucketName, regiontype, region, storageClass string) string {

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
	  resource "ibm_cos_bucket" "bucket2" {
		bucket_name          = "%s"
		resource_instance_id = ibm_resource_instance.instance2.id
		region_location = "%s"
		storage_class        = "%s"
		activity_tracking {
		  read_data_events     = "invalid"
		  write_data_events     = "invalid"
		  management_events     = "invalid"
		  activity_tracker_crn = "%s"
		}
	  }  
	`, cosServiceName, bucketName, region, storageClass, activityTrackerCRN)
}

func testAccCheckIBMCosBucket_Upload_Object_Activity_Tracker_Enabled_With_CRN(cosServiceName, activityTrackerCRN, bucketName, regiontype, region, storageClass, key, object_body string) string {

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
	  resource "ibm_cos_bucket" "bucket2" {
		bucket_name          = "%s"
		resource_instance_id = ibm_resource_instance.instance2.id
		region_location = "%s"
		storage_class        = "%s"
		activity_tracking {
		  read_data_events     = true
		  write_data_events     = true
		  management_events     = true
		  activity_tracker_crn = "%s"
		}
	  } 
	  resource "ibm_cos_bucket_object" "testacc" {
		bucket_crn	    = ibm_cos_bucket.bucket2.crn
		bucket_location = ibm_cos_bucket.bucket2.region_location
		key 					  = "%s.txt"
		content			    = "%s"
	} 
	`, cosServiceName, bucketName, region, storageClass, activityTrackerCRN, key, object_body)
}

func testAccCheckIBMCosBucket_upload_Object_Activity_Tracker_Enabled_Without_CRN(cosServiceName, bucketName, regiontype, region, storageClass, key, object_body string) string {

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
	  resource "ibm_cos_bucket" "bucket2" {
		bucket_name          = "%s"
		resource_instance_id = ibm_resource_instance.instance2.id
		region_location = "%s"
		storage_class        = "%s"
		activity_tracking {
		  read_data_events     = true
		  write_data_events     = true
		  management_events     = true
		}
	  }
	  resource "ibm_cos_bucket_object" "testacc" {
		bucket_crn	    = ibm_cos_bucket.bucket2.crn
		bucket_location = ibm_cos_bucket.bucket2.region_location
		key 					  = "%s.txt"
		content			    = "%s"
	}
	`, cosServiceName, bucketName, region, storageClass, key, object_body)
}

func testAccCheckIBMCosBucket_metricsMonitoring_Without_Crn(cosServiceName, bucketName, regiontype, region, storageClass string, requestMetricsEnabled, usageMetricsEnabled bool) string {

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
	  resource "ibm_cos_bucket" "bucket2" {
		bucket_name          = "%s"
		resource_instance_id = ibm_resource_instance.instance2.id
		region_location = "%s"
		storage_class        = "%s"
		metrics_monitoring {
		  request_metrics_enabled = %v
		  usage_metrics_enabled = %v
		}
	  }  
	`, cosServiceName, bucketName, region, storageClass, requestMetricsEnabled, usageMetricsEnabled)
}

func testAccCheckIBMCosBucket_metricsMonitoring_With_Crn(cosServiceName, metricsMonitoringCrn, bucketName, regiontype, region, storageClass string, requestMetricsEnabled, usageMetricsEnabled bool) string {

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
	  resource "ibm_cos_bucket" "bucket2" {
		bucket_name          = "%s"
		resource_instance_id = ibm_resource_instance.instance2.id
		region_location = "%s"
		storage_class        = "%s"
		metrics_monitoring {
		  request_metrics_enabled = %v
		  usage_metrics_enabled = %v
		  metrics_monitoring_crn = "%s"
		}
	  }  
	`, cosServiceName, bucketName, region, storageClass, requestMetricsEnabled, usageMetricsEnabled, metricsMonitoringCrn)
}

func testAccCheckIBMCosBucket_metricsMonitoring_RequestMetrics_Invalid_UsageMetrics_Invalid_With_Crn(cosServiceName, metricsMonitoringCrn, bucketName, regiontype, region, storageClass string) string {

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
	  resource "ibm_cos_bucket" "bucket2" {
		bucket_name          = "%s"
		resource_instance_id = ibm_resource_instance.instance2.id
		region_location = "%s"
		storage_class        = "%s"
		metrics_monitoring {
		  usage_metrics_enabled = "invalid"
		  request_metrics_enabled = "invalid"
		  metrics_monitoring_crn = "%s"
		}
	  }  
	`, cosServiceName, bucketName, region, storageClass, metricsMonitoringCrn)
}

func testAccCheckIBMCosBucket_metricsMonitoring_Crn_Invalid(cosServiceName, metricsMonitoringCrn, bucketName, regiontype, region, storageClass string) string {

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
	  resource "ibm_cos_bucket" "bucket2" {
		bucket_name          = "%s"
		resource_instance_id = ibm_resource_instance.instance2.id
		region_location = "%s"
		storage_class        = "%s"
		metrics_monitoring {
		  usage_metrics_enabled = true
		  request_metrics_enabled = true
		  metrics_monitoring_crn = "invalid"
		}
	  }  
	`, cosServiceName, bucketName, region, storageClass)
}

func testAccCheckIBMCosBucket_metricsMonitoring_Upload_Object_RequestMetrics_True_UsageMetrics_True_With_Crn(cosServiceName, metricsMonitoringCrn, bucketName, regiontype, region, storageClass, key, objectBody string) string {

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
	  resource "ibm_cos_bucket" "bucket2" {
		bucket_name          = "%s"
		resource_instance_id = ibm_resource_instance.instance2.id
		region_location = "%s"
		storage_class        = "%s"
		metrics_monitoring {
		  usage_metrics_enabled = true
		  request_metrics_enabled = true
		  metrics_monitoring_crn = "%s"
		}
	  }  
	  resource "ibm_cos_bucket_object" "testacc" {
		bucket_crn	    = ibm_cos_bucket.bucket2.crn
		bucket_location = ibm_cos_bucket.bucket2.region_location
		key 			    = "%s.txt"
		content			    = "%s"
	} 
	`, cosServiceName, bucketName, region, storageClass, metricsMonitoringCrn, key, objectBody)
}

func testAccCheckIBMCosBucket_metricsMonitoring_Upload_Object_RequestMetrics_True_UsageMetrics_True_Without_Crn(cosServiceName, bucketName, regiontype, region, storageClass, key, objectBody string) string {

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
	  resource "ibm_cos_bucket" "bucket2" {
		bucket_name          = "%s"
		resource_instance_id = ibm_resource_instance.instance2.id
		region_location = "%s"
		storage_class        = "%s"
		metrics_monitoring {
		  usage_metrics_enabled = true
		  request_metrics_enabled = true
		}
	  }  
	  resource "ibm_cos_bucket_object" "testacc" {
		bucket_crn	    = ibm_cos_bucket.bucket2.crn
		bucket_location = ibm_cos_bucket.bucket2.region_location
		key 			    = "%s.txt"
		content			    = "%s"
	} 
	`, cosServiceName, bucketName, region, storageClass, key, objectBody)
}

// func testAccCheckIBMCosBucket_metricsMonitoring_RequestMetrics_New_Instance(cosServiceName, metricsMonitoringName, bucketName, regiontype, region, storageClass string) string {

// 	return fmt.Sprintf(`

//		data "ibm_resource_group" "cos_group" {
//			is_default=true
//		  }
//		  resource "ibm_resource_instance" "instance2" {
//			name              = "%s"
//			resource_group_id = data.ibm_resource_group.cos_group.id
//			service           = "cloud-object-storage"
//			plan              = "standard"
//			location          = "global"
//		  }
//		  resource "ibm_resource_instance" "metrics_monitor2" {
//			name              = "%s"
//			resource_group_id = data.ibm_resource_group.cos_group.id
//			service           = "sysdig-monitor"
//			plan              = "graduated-tier"
//			location          = "us-south"
//			parameters        = {
//				default_receiver = true
//			}
//		}
//		  resource "ibm_cos_bucket" "bucket2" {
//			bucket_name          = "%s"
//			resource_instance_id = ibm_resource_instance.instance2.id
//			region_location = "%s"
//			storage_class        = "%s"
//			metrics_monitoring {
//			  usage_metrics_enabled = true
//			  request_metrics_enabled = true
//			  metrics_monitoring_crn = ibm_resource_instance.metrics_monitor2.id
//			}
//		  }
//		`, cosServiceName, metricsMonitoringName, bucketName, region, storageClass)
//	}
//
// f1881 end
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

func testAccCheckIBMCosBucket_retention_basic_bucket(bucketName string, cosCrn string, regiontype string, region string, storageClass string) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		name = "Default"
	}

	resource "ibm_cos_bucket" "bucket" {
		bucket_name           = "%s"
		resource_instance_id  = "%s"
	    region_location       = "%s"
		storage_class         = "%s"
	}
	`, bucketName, cosCrn, region, storageClass)
}
func testAccCheckIBMCosBucket_retention_existing_bucket(bucketName string, cosCrn string, regiontype string, region string, storageClass string, default_retention int, maximum_retention int, minimum_retention int) string {

	return fmt.Sprintf(`
	data "ibm_resource_group" "cos_group" {
		name = "Default"
	}

	resource "ibm_cos_bucket" "bucket" {
		bucket_name           = "%s"
		resource_instance_id  = "%s"
	    region_location       = "%s"
		storage_class         = "%s"
		retention_rule {
			default = %d
			maximum = %d
			minimum = %d
			permanent = false
		}
	}
	`, bucketName, cosCrn, region, storageClass, default_retention, maximum_retention, minimum_retention)
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

func testAccCheckIBMKeyProtectRootkeyWithCOSBucket(instanceName, KeyName, serviceName, bucketName, bucketRegion, bucketClass string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "group" {
		is_default=true
	}
	resource "ibm_resource_instance" "kms_instance1" {
		name              = "%s"
		service           = "kms"
		plan              = "tiered-pricing"
		location          = "us-south"
	  }
	  resource "ibm_iam_authorization_policy" "policy1" {
		source_service_name = "cloud-object-storage"
		target_service_name = "kms"
		roles               = ["Reader"]
	  }
	  resource "ibm_kms_key" "test" {
		instance_id = "${ibm_resource_instance.kms_instance1.guid}"
		key_name = "%s"
		standard_key =  false
		force_delete = true
	}

	resource "ibm_resource_instance" "instance" {
		name     = "%s"
		service  = "cloud-object-storage"
		plan     = "standard"
		location = "global"
	}

	resource "ibm_cos_bucket" "bucket" {
		depends_on           = [ibm_iam_authorization_policy.policy1]
		bucket_name          = "%s"
		resource_instance_id = ibm_resource_instance.instance.id
		cross_region_location = "%s"
		storage_class        = "%s"
		key_protect          = ibm_kms_key.test.id
	}
`, instanceName, KeyName, serviceName, bucketName, bucketRegion, bucketClass)
}

func testAccCheckIBMHPCSRootkeyWithCOSBucket(KeyName, serviceName, bucketName, bucketRegion, bucketClass string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "group" {
		is_default=true
	}
	resource "ibm_iam_authorization_policy" "policy1" {
		source_service_name = "cloud-object-storage"
		target_service_name = "hs-crypto"
		roles               = ["Reader"]
	  }

	resource "ibm_resource_instance" "instance" {
		name     = "%s"
		service  = "cloud-object-storage"
		plan     = "standard"
		location = "global"
	}

	resource "ibm_cos_bucket" "bucket" {
		depends_on           = [ibm_iam_authorization_policy.policy1]
		bucket_name          = "%s"
		resource_instance_id = ibm_resource_instance.instance.id
		region_location 	= "%s"
		storage_class       = "%s"
		key_protect			= "%s"
	}
`, serviceName, bucketName, bucketRegion, bucketClass, acc.HpcsRootKeyCrn)
}
func testAccCheckIBMKeyProtectRootkeyWithCOSBucketKmsParam(instanceName, KeyName, serviceName, bucketName, bucketRegion, bucketClass string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "group" {
		is_default=true
	}
	resource "ibm_resource_instance" "kms_instance1" {
		name              = "%s"
		service           = "kms"
		plan              = "tiered-pricing"
		location          = "us-south"
	  }

	  resource "ibm_kms_key" "test" {
		instance_id = "${ibm_resource_instance.kms_instance1.guid}"
		key_name = "%s"
		standard_key =  false
		force_delete = true
	}

	resource "ibm_resource_instance" "instance" {
		name     = "%s"
		service  = "cloud-object-storage"
		plan     = "standard"
		location = "global"
	}

	resource "ibm_cos_bucket" "bucket" {
		bucket_name          = "%s"
		resource_instance_id = ibm_resource_instance.instance.id
		cross_region_location = "%s"
		storage_class        = "%s"
		kms_key_crn          = ibm_kms_key.test.id
	}
`, instanceName, KeyName, serviceName, bucketName, bucketRegion, bucketClass)
}
func testAccCheckIBMKeyProtectRootkeyWithCOSBucketKmsParamWithInvalidCRN(instanceName, KeyName, serviceName, bucketName, bucketRegion, bucketClass string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "group" {
		is_default=true
	}
	
	resource "ibm_resource_instance" "instance" {
		name     = "%s"
		service  = "cloud-object-storage"
		plan     = "standard"
		location = "global"
	}

	resource "ibm_cos_bucket" "bucket" {
		bucket_name          = "%s"
		resource_instance_id = ibm_resource_instance.instance.id
		cross_region_location = "%s"
		storage_class        = "%s"
		kms_key_crn          = "crn:v1:staging:public:kms:us-south:invalid"
	}
`, instanceName, bucketName, bucketRegion, bucketClass)
}

func testAccCheckIBMHPCSRootkeyWithCOSBucketKmsParam(KeyName, serviceName, bucketName, bucketRegion, bucketClass string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "group" {
		is_default=true
	}
	resource "ibm_resource_instance" "instance" {
		name     = "%s"
		service  = "cloud-object-storage"
		plan     = "standard"
		location = "global"
	}

	resource "ibm_cos_bucket" "bucket" {
		bucket_name          = "%s"
		resource_instance_id = ibm_resource_instance.instance.id
		region_location 	= "%s"
		storage_class       = "%s"
		kms_key_crn			= "%s"
	}
`, serviceName, bucketName, bucketRegion, bucketClass, acc.HpcsRootKeyCrn)
}

func testAccCheckIBMHPCSRootkeyWithCOSBucketKMSBothParamProvided(KeyName, serviceName, bucketName, bucketRegion, bucketClass string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "group" {
		is_default=true
	}
	resource "ibm_resource_instance" "instance" {
		name     = "%s"
		service  = "cloud-object-storage"
		plan     = "standard"
		location = "global"
	}

	resource "ibm_cos_bucket" "bucket" {
		bucket_name          = "%s"
		resource_instance_id = ibm_resource_instance.instance.id
		region_location 	= "%s"
		storage_class       = "%s"
		kms_key_crn			= "%s"
		key_protect 		= "%s"
	}
`, serviceName, bucketName, bucketRegion, bucketClass, acc.HpcsRootKeyCrn, acc.HpcsRootKeyCrn)
}

func testAccCheckIBMHPCSRootkeyWithCOSBucketKmsParamWithInvalidCRN(KeyName, serviceName, bucketName, bucketRegion, bucketClass string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "group" {
		is_default=true
	}
	resource "ibm_resource_instance" "instance" {
		name     = "%s"
		service  = "cloud-object-storage"
		plan     = "standard"
		location = "global"
	}

	resource "ibm_cos_bucket" "bucket" {
		bucket_name          = "%s"
		resource_instance_id = ibm_resource_instance.instance.id
		region_location 	= "%s"
		storage_class       = "%s"
		kms_key_crn			= "crn:v1:staging:public:hs-crypto:us-south:invalid"
	}
`, serviceName, bucketName, bucketRegion, bucketClass)
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
