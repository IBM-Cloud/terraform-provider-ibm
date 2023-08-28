package cos_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCosBucket_Website_Configuration_Bucket_Basic(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraformstaticwebhosting%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"
	indexSuffix := "index.html"
	errorKey := "error.html"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Website_Configuration_Bucket_Basic(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, indexSuffix, errorKey),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "cross_region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.0.error_document.0.key", errorKey),
					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.0.index_document.0.suffix", indexSuffix),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Website_Configuration_Bucket_Without_Public_Access(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraformstaticwebhosting%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"
	indexSuffix := "index.html"
	errorKey := "error.html"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Website_Configuration_Bucket_Without_Public_Access(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, indexSuffix, errorKey),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "cross_region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.0.error_document.0.key", errorKey),
					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.0.index_document.0.suffix", indexSuffix),
				),
			},
		},
	})
}
func TestAccIBMCosBucket_Website_Configuration_Bucket_Index_Document_Only(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraformterraformstaticwebhosting%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"
	indexSuffix := "index.html"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Website_Configuration_Bucket_Index_Document_Only(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, indexSuffix),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "cross_region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.0.index_document.0.suffix", indexSuffix),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Website_Configuration_Bucket_With_Routing_Rules(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraformstaticwebhosting%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"
	indexSuffix := "index.html"
	errorKey := "error.html"
	hostName := fmt.Sprintf(bucketName, ".s3-web.us-south.cloud-object-storage.appdomain.cloud")
	protocol := "https"
	http_redirect_code := "302"
	replace_key_with := "error404.html"
	http_error_code_returned_equals := "404"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Website_Configuration_Bucket_With_Routing_Rule(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, indexSuffix, errorKey, http_error_code_returned_equals, hostName, protocol, http_redirect_code, replace_key_with),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "cross_region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.0.routing_rule.0.condition.0.http_error_code_returned_equals", http_error_code_returned_equals),
					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.0.routing_rule.0.redirect.0.host_name", hostName),
					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.0.routing_rule.0.redirect.0.protocol", protocol),
					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.0.routing_rule.0.redirect.0.http_redirect_code", http_redirect_code),
					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.0.routing_rule.0.redirect.0.replace_key_with", replace_key_with),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Website_Configuration_Bucket_With_JSON_Routing_Rules(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraformstaticwebhosting%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"
	keyPrefixEquals := "docs/"
	indexSuffix := "index.html"
	errorKey := "error.html"
	replaceKeyPrefixWith := "documents/"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Website_Configuration_Bucket_With_JSON_Routing_Rule(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, indexSuffix, errorKey, keyPrefixEquals, replaceKeyPrefixWith),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "cross_region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.0.routing_rule.0.condition.0.key_prefix_equals", keyPrefixEquals),
					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.0.routing_rule.0.condition.0.key_prefix_equals", replaceKeyPrefixWith),
				),
			},
		},
	})
}
func TestAccIBMCosBucket_Website_Configuration_Bucket_With_Routing_Rules_Condition_Only(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraformstaticwebhosting%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"
	indexSuffix := "index.html"
	errorKey := "error.html"
	http_error_code_returned_equals := "404"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Website_Configuration_Bucket_With_Routing_Rule_Condition_Only(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, indexSuffix, errorKey, http_error_code_returned_equals),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "cross_region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.0.routing_rule.0.condition.0.http_error_code_returned_equals", http_error_code_returned_equals),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Website_Configuration_Bucket_With_Routing_Rules_Redirect_Only(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraformstaticwebhosting%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"
	indexSuffix := "index.html"
	errorKey := "error.html"
	hostName := fmt.Sprintf(bucketName, ".s3-web.us-south.cloud-object-storage.appdomain.cloud")
	protocol := "https"
	http_redirect_code := "302"
	replace_key_with := "error404.html"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Website_Configuration_Bucket_With_Routing_Rule_Redirect_Only(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, indexSuffix, errorKey, hostName, protocol, http_redirect_code, replace_key_with),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "cross_region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.0.routing_rule.0.redirect.0.host_name", hostName),
					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.0.routing_rule.0.redirect.0.protocol", protocol),
					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.0.routing_rule.0.redirect.0.http_redirect_code", http_redirect_code),
					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.0.routing_rule.0.redirect.0.replace_key_with", replace_key_with),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Website_Configuration_Bucket_With_Multiple_Routing_Rules(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraformstaticwebhosting%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"
	indexSuffix := "index.html"
	errorKey := "error.html"
	hostName := fmt.Sprintf(bucketName, ".s3-web.us-south.cloud-object-storage.appdomain.cloud")
	protocol := "https"
	http_redirect_code := "302"
	replace_key_with := "error404.html"
	http_error_code_returned_equals := "404"
	http_redirect_code2 := "303"
	replace_key_with2 := "error405.html"
	http_error_code_returned_equals2 := "405"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Website_Configuration_Bucket_With_Multiple_Routing_Rule(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, indexSuffix, errorKey, http_error_code_returned_equals, hostName, protocol, http_redirect_code, replace_key_with, http_error_code_returned_equals2, http_redirect_code2, replace_key_with2),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "cross_region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.0.routing_rule.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMCosBucket_Website_Configuration_Bucket_Upload_Object_With_Redirect(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraformstaticwebhosting%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"
	indexSuffix := "index.html"
	websiteRedirectLocation := "/redirect"
	key := "key1"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCosBucket_Website_Configuration_Bucket_Upload_Object_With_Redirect(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, indexSuffix, key, websiteRedirectLocation),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "cross_region_location", bucketRegion),
					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.#", "1"),
					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.0.index_document.0.suffix", indexSuffix),
					resource.TestCheckResourceAttr("ibm_cos_bucket_object.object", "website_redirect", websiteRedirectLocation),
				),
			},
		},
	})
}

// func TestAccIBMCosBucket_Website_Configuration_Bucket_Update_Index_and_Error(t *testing.T) {
// 	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
// 	bucketName := fmt.Sprintf("terraformstaticwebhosting%d", acctest.RandIntRange(10, 100))
// 	bucketRegion := "us"
// 	bucketClass := "standard"
// 	bucketRegionType := "cross_region_location"
// 	indexSuffix := "index.html"
// 	errorKey := "error.html"
// 	updatedIndexSuffix := "index2.html"
// 	updatederrorKey := "error2.html"
// 	resource.Test(t, resource.TestCase{
// 		PreCheck:     func() { acc.TestAccPreCheck(t) },
// 		Providers:    acc.TestAccProviders,
// 		CheckDestroy: testAccCheckIBMCosBucketDestroy,
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccCheckIBMCosBucket_Website_Configuration_Bucket_Basic(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, indexSuffix, errorKey),
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
// 					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
// 					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
// 					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "cross_region_location", bucketRegion),
// 					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.#", "1"),
// 					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.0.error_document.0.key", errorKey),
// 					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.0.index_document.0.suffix", indexSuffix),
// 				),
// 			},
// 			{
// 				Config: testAccCheckIBMCosBucket_Website_Configuration_Bucket_Basic(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, updatedIndexSuffix, updatederrorKey),
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
// 					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
// 					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
// 					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "cross_region_location", bucketRegion),
// 					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.#", "1"),
// 					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.0.error_document.0.key", updatederrorKey),
// 					resource.TestCheckResourceAttr("ibm_cos_bucket_website_configuration.website", "website_configuration.0.index_document.0.suffix", updatedIndexSuffix),
// 				),
// 			},
// 		},
// 	})
// }

func TestAccIBMCosBucket_Website_Configuration_Bucket_Empty_Config(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraformstaticwebhosting%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCosBucket_Website_Configuration_Bucket_Empty(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass),
				ExpectError: regexp.MustCompile("Error: failed to put website configuration on the COS bucket"),
			},
		},
	})
}

func TestAccIBMCosBucket_Website_Configuration_Bucket_Index_And_Redirect_Together(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraformstaticwebhosting%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"
	indexSuffix := "index.html"
	hostName := fmt.Sprintf(bucketName, ".s3-web.us-south.cloud-object-storage.appdomain.cloud")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCosBucket_Website_Configuration_Bucket_Index_And_Redirect_Together(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass, indexSuffix, hostName),
				ExpectError: regexp.MustCompile("Error: Conflicting configuration arguments"),
			},
		},
	})
}

func TestAccIBMCosBucket_Website_Configuration_Bucket_No_Config(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandIntRange(10, 100))
	bucketName := fmt.Sprintf("terraformstaticwebhosting%d", acctest.RandIntRange(10, 100))
	bucketRegion := "us"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMCosBucket_Website_Configuration_Bucket_No_Config(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass),
				ExpectError: regexp.MustCompile("Error: Insufficient website_configuration blocks"),
			},
		},
	})
}

func testAccCheckIBMCosBucket_Website_Configuration_Bucket_Basic(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, indexSuffix string, errorKey string) string {

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
	}
	data "ibm_iam_access_group" "public_access_group" { 
		access_group_name = "Public Access" 
	} 
	 
resource "ibm_iam_access_group_policy" "policy" { 
	depends_on = [ibm_cos_bucket.bucket] 
	access_group_id = data.ibm_iam_access_group.public_access_group.groups[0].id 
	roles = ["Object Reader"]  
	resources { 
	service = "cloud-object-storage" 
	resource_type = "bucket" 
	resource_instance_id = "%s" 
	resource = "%s"
	} 
} 

resource ibm_cos_bucket_website_configuration "website" {
	bucket_crn = ibm_cos_bucket.bucket.crn
	bucket_location = ibm_cos_bucket.bucket.cross_region_location
	website_configuration {
	  error_document{
	    key = "%s"
	  }
	  index_document{
	    suffix = "%s"
	  }
	}
}
	 
	`, cosServiceName, bucketName, region, storageClass, cosServiceName, bucketName, errorKey, indexSuffix)
}

func testAccCheckIBMCosBucket_Website_Configuration_Bucket_Index_Document_Only(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, indexSuffix string) string {

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
	}
	data "ibm_iam_access_group" "public_access_group" { 
		access_group_name = "Public Access" 
	} 
	 
resource "ibm_iam_access_group_policy" "policy" { 
	depends_on = [ibm_cos_bucket.bucket] 
	access_group_id = data.ibm_iam_access_group.public_access_group.groups[0].id 
	roles = ["Object Reader"]  
	resources { 
	service = "cloud-object-storage" 
	resource_type = "bucket" 
	resource_instance_id = "%s" 
	resource = "%s"
	} 
} 

resource ibm_cos_bucket_website_configuration "website" {
	bucket_crn = ibm_cos_bucket.bucket.crn
	bucket_location = ibm_cos_bucket.bucket.cross_region_location
	website_configuration {
	  index_document{
	    suffix = "%s"
	  }
	}
}
	 
	`, cosServiceName, bucketName, region, storageClass, cosServiceName, bucketName, indexSuffix)
}

func testAccCheckIBMCosBucket_Website_Configuration_Bucket_With_Routing_Rule(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, indexSuffix string, errorKey string, httpErrorCodeReturnedEquals string, hostName string, protocol string, httpsRedirectCode string, replaceKeyWith string) string {

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
	}
	data "ibm_iam_access_group" "public_access_group" { 
		access_group_name = "Public Access" 
	} 
	 
resource "ibm_iam_access_group_policy" "policy" { 
	depends_on = [ibm_cos_bucket.bucket] 
	access_group_id = data.ibm_iam_access_group.public_access_group.groups[0].id 
	roles = ["Object Reader"]  
	resources { 
	service = "cloud-object-storage" 
	resource_type = "bucket" 
	resource_instance_id = "%s" 
	resource = "%s"
	} 
} 

resource ibm_cos_bucket_website_configuration "website" {
	bucket_crn = ibm_cos_bucket.bucket.crn
	bucket_location = ibm_cos_bucket.bucket.cross_region_location
	website_configuration {
		error_document{
	      key = "%s"
	    }
	    index_document{
	      suffix = "%s"
	    }
	    routing_rule{
	      condition{
	      http_error_code_returned_equals= "%s"
	    }
	    redirect{
	      host_name= "%s"
	      http_redirect_code= "%s"
	      protocol = "%s"
	      replace_key_with = "%s"
	    	}
	    }
	}
}
	 
	`, cosServiceName, bucketName, region, storageClass, cosServiceName, bucketName, errorKey, indexSuffix, httpErrorCodeReturnedEquals, hostName, httpsRedirectCode, protocol, replaceKeyWith)
}

func testAccCheckIBMCosBucket_Website_Configuration_Bucket_With_JSON_Routing_Rule(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, indexSuffix string, errorKey string, httpErrorCodeReturnedEquals string) string {

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
	}
	data "ibm_iam_access_group" "public_access_group" { 
		access_group_name = "Public Access" 
	} 
	 
resource "ibm_iam_access_group_policy" "policy" { 
	depends_on = [ibm_cos_bucket.bucket] 
	access_group_id = data.ibm_iam_access_group.public_access_group.groups[0].id 
	roles = ["Object Reader"]  
	resources { 
	service = "cloud-object-storage" 
	resource_type = "bucket" 
	resource_instance_id = "%s" 
	resource = "%s"
	} 
} 

resource ibm_cos_bucket_website_configuration "website" {
	bucket_crn = ibm_cos_bucket.bucket.crn
	bucket_location = ibm_cos_bucket.bucket.cross_region_location
	website_configuration {
		error_document{
	      key = "%s"
	    }
	    index_document{
	      suffix = "%s"
	    }
		routing_rules = <<EOF
		[{
		    "Condition": {
		        "KeyPrefixEquals": "docs/"
		     },
		     "Redirect": {
		         "ReplaceKeyPrefixWith": "documents/"
		     }
		 }]
		 EOF
	}
}
	 
	`, cosServiceName, bucketName, region, storageClass, cosServiceName, bucketName, errorKey, indexSuffix)
}

func testAccCheckIBMCosBucket_Website_Configuration_Bucket_With_Routing_Rule_Condition_Only(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, indexSuffix string, errorKey string, httpErrorCodeReturnedEquals string) string {

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
	}
	data "ibm_iam_access_group" "public_access_group" { 
		access_group_name = "Public Access" 
	} 
	 
resource "ibm_iam_access_group_policy" "policy" { 
	depends_on = [ibm_cos_bucket.bucket] 
	access_group_id = data.ibm_iam_access_group.public_access_group.groups[0].id 
	roles = ["Object Reader"]  
	resources { 
	service = "cloud-object-storage" 
	resource_type = "bucket" 
	resource_instance_id = "%s" 
	resource = "%s"
	} 
} 

resource ibm_cos_bucket_website_configuration "website" {
	bucket_crn = ibm_cos_bucket.bucket.crn
	bucket_location = ibm_cos_bucket.bucket.cross_region_location
	website_configuration {
		error_document{
	      key = "%s"
	    }
	    index_document{
	      suffix = "%s"
	    }
	    routing_rule{
	      condition{
	      http_error_code_returned_equals= "%s"
	    }
	    }
	}
}
	 
	`, cosServiceName, bucketName, region, storageClass, cosServiceName, bucketName, errorKey, indexSuffix, httpErrorCodeReturnedEquals)
}

func testAccCheckIBMCosBucket_Website_Configuration_Bucket_With_Routing_Rule_Redirect_Only(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, indexSuffix string, errorKey string, hostName string, protocol string, httpsRedirectCode string, replaceKeyWith string) string {

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
	}
	data "ibm_iam_access_group" "public_access_group" { 
		access_group_name = "Public Access" 
	} 
	 
resource "ibm_iam_access_group_policy" "policy" { 
	depends_on = [ibm_cos_bucket.bucket] 
	access_group_id = data.ibm_iam_access_group.public_access_group.groups[0].id 
	roles = ["Object Reader"]  
	resources { 
	service = "cloud-object-storage" 
	resource_type = "bucket" 
	resource_instance_id = "%s" 
	resource = "%s"
	} 
} 

resource ibm_cos_bucket_website_configuration "website" {
	bucket_crn = ibm_cos_bucket.bucket.crn
	bucket_location = ibm_cos_bucket.bucket.cross_region_location
	website_configuration {
		error_document{
	      key = "%s"
	    }
	    index_document{
	      suffix = "%s"
	    }
	    routing_rule{
	    redirect{
	      host_name= "%s"
	      http_redirect_code= "%s"
	      protocol = "%s"
	      replace_key_with = "%s"
	    	}
	    }
	}
}
	 
	`, cosServiceName, bucketName, region, storageClass, cosServiceName, bucketName, errorKey, indexSuffix, hostName, httpsRedirectCode, protocol, replaceKeyWith)
}

func testAccCheckIBMCosBucket_Website_Configuration_Bucket_With_Multiple_Routing_Rule(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, indexSuffix string, errorKey string, httpErrorCodeReturnedEquals string, hostName string, protocol string, httpsRedirectCode string, replaceKeyWith string, httpErrorCodeReturnedEquals2 string, httpsRedirectCode2 string, replaceKeyWith2 string) string {

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
	}
	data "ibm_iam_access_group" "public_access_group" { 
		access_group_name = "Public Access" 
	} 
	 
resource "ibm_iam_access_group_policy" "policy" { 
	depends_on = [ibm_cos_bucket.bucket] 
	access_group_id = data.ibm_iam_access_group.public_access_group.groups[0].id 
	roles = ["Object Reader"]  
	resources { 
	service = "cloud-object-storage" 
	resource_type = "bucket" 
	resource_instance_id = "%s" 
	resource = "%s"
	} 
} 

resource ibm_cos_bucket_website_configuration "website" {
	bucket_crn = ibm_cos_bucket.bucket.crn
	bucket_location = ibm_cos_bucket.bucket.cross_region_location
	website_configuration {
		error_document{
	      key = "%s"
	    }
	    index_document{
	      suffix = "%s"
	    }
	    routing_rule{
	      condition{
	      http_error_code_returned_equals= "%s"
	    }
	    redirect{
	      host_name= "%s"
	      http_redirect_code= "%s"
	      protocol = "%s"
	      replace_key_with = "%s"
	    	}
	    }
		routing_rule{
			condition{
			http_error_code_returned_equals= "%s"
		  }
		  redirect{
			host_name= "%s"
			http_redirect_code= "%s"
			protocol = "%s"
			replace_key_with = "%s"
			  }
		  }
	}
}
	 
	`, cosServiceName, bucketName, region, storageClass, cosServiceName, bucketName, errorKey, indexSuffix, httpErrorCodeReturnedEquals, hostName, httpsRedirectCode, protocol, replaceKeyWith, httpErrorCodeReturnedEquals2, hostName, httpsRedirectCode2, protocol, replaceKeyWith2)
}

func testAccCheckIBMCosBucket_Website_Configuration_Bucket_Without_Public_Access(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, indexSuffix string, errorKey string) string {

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
	}
	
	resource ibm_cos_bucket_website_configuration "website" {
		bucket_crn = ibm_cos_bucket.bucket.crn
		bucket_location = ibm_cos_bucket.bucket.cross_region_location
		website_configuration {
		error_document{
			key = "%s"
		}
		index_document{
			suffix = "%s"
		}
		}
	}
	 
	`, cosServiceName, bucketName, region, storageClass, errorKey, indexSuffix)
}

func testAccCheckIBMCosBucket_Website_Configuration_Bucket_Upload_Object_With_Redirect(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, indexSuffix string, key string, websiteRedirectLocation string) string {

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
	}
	data "ibm_iam_access_group" "public_access_group" { 
		access_group_name = "Public Access" 
	} 
	 
resource "ibm_iam_access_group_policy" "policy" { 
	depends_on = [ibm_cos_bucket.bucket] 
	access_group_id = data.ibm_iam_access_group.public_access_group.groups[0].id 
	roles = ["Object Reader"]  
	resources { 
	service = "cloud-object-storage" 
	resource_type = "bucket" 
	resource_instance_id = "%s" 
	resource = "%s"
	} 
} 

resource ibm_cos_bucket_website_configuration "website" {
	bucket_crn = ibm_cos_bucket.bucket.crn
	bucket_location = ibm_cos_bucket.bucket.cross_region_location
	website_configuration {
	  index_document{
	    suffix = "%s"
	  }
	}
}

resource "ibm_cos_bucket_object" "object" {
	bucket_crn	    = ibm_cos_bucket.bucket.crn
	bucket_location = ibm_cos_bucket.bucket.cross_region_location
	key 					  = "%s"
	content			    = "Acceptance testing"
	website_redirect = "%s"
}
	 
	`, cosServiceName, bucketName, region, storageClass, cosServiceName, bucketName, indexSuffix, key, websiteRedirectLocation)
}

func testAccCheckIBMCosBucket_Website_Configuration_Bucket_Empty(cosServiceName string, bucketName string, regiontype string, region string, storageClass string) string {

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
	}
	data "ibm_iam_access_group" "public_access_group" { 
		access_group_name = "Public Access" 
	} 
	 
resource "ibm_iam_access_group_policy" "policy" { 
	depends_on = [ibm_cos_bucket.bucket] 
	access_group_id = data.ibm_iam_access_group.public_access_group.groups[0].id 
	roles = ["Object Reader"]  
	resources { 
	service = "cloud-object-storage" 
	resource_type = "bucket" 
	resource_instance_id = "%s" 
	resource = "%s"
	} 
} 

resource ibm_cos_bucket_website_configuration "website" {
	bucket_crn = ibm_cos_bucket.bucket.crn
	bucket_location = ibm_cos_bucket.bucket.cross_region_location
	website_configuration {
	}
}
	 
	`, cosServiceName, bucketName, region, storageClass, cosServiceName, bucketName)
}

func testAccCheckIBMCosBucket_Website_Configuration_Bucket_Index_And_Redirect_Together(cosServiceName string, bucketName string, regiontype string, region string, storageClass string, indexSuffix string, hostName string) string {

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
	}
	data "ibm_iam_access_group" "public_access_group" { 
		access_group_name = "Public Access" 
	} 
	 
resource "ibm_iam_access_group_policy" "policy" { 
	depends_on = [ibm_cos_bucket.bucket] 
	access_group_id = data.ibm_iam_access_group.public_access_group.groups[0].id 
	roles = ["Object Reader"]  
	resources { 
	service = "cloud-object-storage" 
	resource_type = "bucket" 
	resource_instance_id = "%s" 
	resource = "%s"
	} 
} 

resource ibm_cos_bucket_website_configuration "website" {
	bucket_crn = ibm_cos_bucket.bucket.crn
	bucket_location = ibm_cos_bucket.bucket.cross_region_location
	website_configuration {
	  index_document{
	    suffix = "%s"
	  }
	  redirect_all_requests_to{
		host_name = "%s"
	}
	}
}
	 
	`, cosServiceName, bucketName, region, storageClass, cosServiceName, bucketName, indexSuffix, hostName)
}

func testAccCheckIBMCosBucket_Website_Configuration_Bucket_No_Config(cosServiceName string, bucketName string, regiontype string, region string, storageClass string) string {

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
	}
	data "ibm_iam_access_group" "public_access_group" { 
		access_group_name = "Public Access" 
	} 
	 
resource "ibm_iam_access_group_policy" "policy" { 
	depends_on = [ibm_cos_bucket.bucket] 
	access_group_id = data.ibm_iam_access_group.public_access_group.groups[0].id 
	roles = ["Object Reader"]  
	resources { 
	service = "cloud-object-storage" 
	resource_type = "bucket" 
	resource_instance_id = "%s" 
	resource = "%s"
	} 
} 

resource ibm_cos_bucket_website_configuration "website" {
	bucket_crn = ibm_cos_bucket.bucket.crn
	bucket_location = ibm_cos_bucket.bucket.cross_region_location
}
	 
	`, cosServiceName, bucketName, region, storageClass, cosServiceName, bucketName)
}
