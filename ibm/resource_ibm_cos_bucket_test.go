package ibm

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/IBM/ibm-cos-sdk-go/aws"
	"github.com/IBM/ibm-cos-sdk-go/aws/credentials/ibmiam"
	token "github.com/IBM/ibm-cos-sdk-go/aws/credentials/ibmiam/token"
	"github.com/IBM/ibm-cos-sdk-go/aws/session"
	"github.com/IBM/ibm-cos-sdk-go/service/s3"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMCosBucket_Basic(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	bucketName := fmt.Sprintf("terraform%d", acctest.RandInt())
	bucketRegion := "eu"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCosBucket_basic(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "cross_region_location", bucketRegion),
				),
			},
			resource.TestStep{
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

func TestAccIBMCosBucket_Smart_Type(t *testing.T) {
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	bucketName := fmt.Sprintf("terraform%d", acctest.RandInt())
	bucketRegion := "eu"
	bucketClass := "smart"
	bucketRegionType := "cross_region_location"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCosBucket_basic(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "cross_region_location", bucketRegion),
				),
			},
			resource.TestStep{
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
	serviceName := fmt.Sprintf("terraform_%d", acctest.RandInt())
	bucketName := fmt.Sprintf("terraform%d", acctest.RandInt())
	bucketRegion := "eu"
	bucketClass := "standard"
	bucketRegionType := "cross_region_location"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMCosBucketDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCosBucket_basic(serviceName, bucketName, bucketRegionType, bucketRegion, bucketClass),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMCosBucketExists("ibm_resource_instance.instance", "ibm_cos_bucket.bucket", bucketRegionType, bucketRegion, bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "bucket_name", bucketName),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "storage_class", bucketClass),
					resource.TestCheckResourceAttr("ibm_cos_bucket.bucket", "cross_region_location", bucketRegion),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_cos_bucket.bucket",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"wait_time_minutes", "parameters"},
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
		if rs.Type == "ibm_resource_instance" {
			resourceInstance = rs.Primary.Attributes["crn"]

		}
	}

	rsContClient, err := testAccProvider.Meta().(ClientSession).BluemixSession()
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

		apiEndpoint, _ := selectCosApi(rt, region)

		rsContClient, err := testAccProvider.Meta().(ClientSession).BluemixSession()
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
		name = "default"
	}
	  
	resource "ibm_resource_instance" "instance" {
		name              = "%s"
		service           = "cloud-object-storage"
		plan              = "lite"
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

func testAccCheckIBMCosBucket_updateWithSameName(serviceName string, bucketName string, regiontype string, region, storageClass string) string {
	return fmt.Sprintf(`	
	data "ibm_resource_group" "group" {
		name = "default"
	}
	  
	resource "ibm_resource_instance" "instance" {
		name              = "%s"
		service           = "cloud-object-storage"
		plan              = "lite"
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
