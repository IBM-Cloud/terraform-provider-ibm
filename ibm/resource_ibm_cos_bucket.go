package ibm

import (
	"fmt"
	"strings"
	"time"

	"github.com/IBM/ibm-cos-sdk-go/aws"
	"github.com/IBM/ibm-cos-sdk-go/aws/credentials/ibmiam"
	"github.com/IBM/ibm-cos-sdk-go/aws/session"
	"github.com/IBM/ibm-cos-sdk-go/service/s3"
	"github.com/hashicorp/terraform/helper/schema"
)

var singleSiteLocation = []string{
	"ams03", "che01", "hkg02", "mel01", "mex01",
	"mil01", "mon01", "osl01", "sjc04", "sao01",
	"seo01", "tor01",
}

var regionLocation = []string{
	"au-syd", "eu-de", "eu-gb", "jp-tok", "us-east", "us-south",
}

var crossRegionLocation = []string{
	"us", "eu", "ap",
}

var storageClass = []string{
	"standard", "vault", "cold", "flex",
}

const (
	authEndpoint = "https://iam.cloud.ibm.com/identity/token"
	keyAlgorithm = "AES256"
)

func resourceIBMCOS() *schema.Resource {
	return &schema.Resource{
		Read:     resourceIBMCOSRead,
		Create:   resourceIBMCOSCreate,
		Delete:   resourceIBMCOSDelete,
		Exists:   resourceIBMCOSExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"bucket_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"resource_instance_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "CRN of resource instance",
			},
			"key_protect": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Description: "CRN of the key you want to use data at rest encryption",
			},
			"single_site_location": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ValidateFunc:  validateAllowedStringValue(singleSiteLocation),
				ForceNew:      true,
				ConflictsWith: []string{"region_location", "cross_region_location"},
			},
			"region_location": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ValidateFunc:  validateAllowedStringValue(regionLocation),
				ForceNew:      true,
				ConflictsWith: []string{"cross_region_location", "single_site_location"},
			},
			"cross_region_location": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ValidateFunc:  validateAllowedStringValue(crossRegionLocation),
				ForceNew:      true,
				ConflictsWith: []string{"region_location", "single_site_location"},
			},
			"storage_class": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue(storageClass),
				ForceNew:     true,
			},
		},
	}
}

func resourceIBMCOSRead(d *schema.ResourceData, meta interface{}) error {
	rsConClient, err := meta.(ClientSession).BluemixSession()
	if err != nil {
		return err
	}
	bucketName := d.Get("bucket_name").(string)
	serviceID := d.Get("resource_instance_id").(string)
	var bLocation string
	var apiType string
	if bucketLocation, ok := d.GetOk("cross_region_location"); ok {
		bLocation = bucketLocation.(string)
		apiType = "crl"
	}
	if bucketLocation, ok := d.GetOk("region_location"); ok {
		bLocation = bucketLocation.(string)
		apiType = "rl"
	}
	if bucketLocation, ok := d.GetOk("single_site_location"); ok {
		bLocation = bucketLocation.(string)
		apiType = "ssl"
	}
	apiEndpoint := selectCosApi(apiType, bLocation)

	apiKey := rsConClient.Config.BluemixAPIKey
	s3Conf := aws.NewConfig().WithEndpoint(apiEndpoint).WithCredentials(ibmiam.NewStaticCredentials(aws.NewConfig(), authEndpoint, apiKey, serviceID)).WithS3ForcePathStyle(true)

	s3Sess := session.Must(session.NewSession())
	s3Client := s3.New(s3Sess, s3Conf)

	headInput := &s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	}
	err = s3Client.WaitUntilBucketExists(headInput)
	if err != nil {
		return fmt.Errorf("failed waiting for bucket %s to be created, %v",
			bucketName, err)
	}

	d.Set("bucket_name", d.Get("bucket_name").(string))

	return nil
}

func resourceIBMCOSCreate(d *schema.ResourceData, meta interface{}) error {
	rsConClient, err := meta.(ClientSession).BluemixSession()
	if err != nil {
		return err
	}
	bucketName := d.Get("bucket_name").(string)
	storageClass := d.Get("storage_class").(string)
	var bLocation string
	var apiType string
	serviceID := d.Get("resource_instance_id").(string)

	if bucketLocation, ok := d.GetOk("cross_region_location"); ok {
		bLocation = bucketLocation.(string)
		apiType = "crl"
	}
	if bucketLocation, ok := d.GetOk("region_location"); ok {
		bLocation = bucketLocation.(string)
		apiType = "rl"
	}
	if bucketLocation, ok := d.GetOk("single_site_location"); ok {
		bLocation = bucketLocation.(string)
		apiType = "ssl"
	}
	lConstraint := fmt.Sprintf("%s-%s", bLocation, storageClass)
	apiEndpoint := selectCosApi(apiType, bLocation)
	create := &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
		CreateBucketConfiguration: &s3.CreateBucketConfiguration{
			LocationConstraint: aws.String(lConstraint),
		},
	}

	if keyprotect, ok := d.GetOk("key_protect"); ok {
		create.IBMSSEKPCustomerRootKeyCrn = aws.String(keyprotect.(string))
		create.IBMSSEKPEncryptionAlgorithm = aws.String(keyAlgorithm)
	}

	apiKey := rsConClient.Config.BluemixAPIKey
	s3Conf := aws.NewConfig().WithEndpoint(apiEndpoint).WithCredentials(ibmiam.NewStaticCredentials(aws.NewConfig(), authEndpoint, apiKey, serviceID)).WithS3ForcePathStyle(true)

	s3Sess := session.Must(session.NewSession())
	s3Client := s3.New(s3Sess, s3Conf)

	_, err = s3Client.CreateBucket(create)
	if err != nil {
		return err
	}
	// Generating CRN to create a "fake" id for TF
	bucketCRN := fmt.Sprintf("%s:%s:%s", strings.Replace(serviceID, "::", "", -1), "bucket", bucketName)
	d.SetId(bucketCRN)
	return resourceIBMCOSRead(d, meta)
}

func resourceIBMCOSDelete(d *schema.ResourceData, meta interface{}) error {
	rsConClient, _ := meta.(ClientSession).BluemixSession()
	bucketName := d.Get("bucket_name").(string)
	serviceID := d.Get("resource_instance_id").(string)
	var bLocation string
	var apiType string
	if bucketLocation, ok := d.GetOk("cross_region_location"); ok {
		bLocation = bucketLocation.(string)
		apiType = "crl"
	}
	if bucketLocation, ok := d.GetOk("region_location"); ok {
		bLocation = bucketLocation.(string)
		apiType = "rl"
	}
	if bucketLocation, ok := d.GetOk("single_site_location"); ok {
		bLocation = bucketLocation.(string)
		apiType = "ssl"
	}
	apiEndpoint := selectCosApi(apiType, bLocation)

	apiKey := rsConClient.Config.BluemixAPIKey
	s3Conf := aws.NewConfig().WithEndpoint(apiEndpoint).WithCredentials(ibmiam.NewStaticCredentials(aws.NewConfig(), authEndpoint, apiKey, serviceID)).WithS3ForcePathStyle(true)

	s3Sess := session.Must(session.NewSession())
	s3Client := s3.New(s3Sess, s3Conf)

	delete := &s3.DeleteBucketInput{
		Bucket: aws.String(bucketName),
	}
	s3Client.DeleteBucket(delete)
	d.SetId("")
	return nil
}

func resourceIBMCOSExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	rsConClient, err := meta.(ClientSession).BluemixSession()
	if err != nil {
		return false, err
	}
	bucketName := d.Get("bucket_name").(string)
	serviceID := d.Get("resource_instance_id").(string)
	var bLocation string
	var apiType string
	if bucketLocation, ok := d.GetOk("cross_region_location"); ok {
		bLocation = bucketLocation.(string)
		apiType = "crl"
	}
	if bucketLocation, ok := d.GetOk("region_location"); ok {
		bLocation = bucketLocation.(string)
		apiType = "rl"
	}
	if bucketLocation, ok := d.GetOk("single_site_location"); ok {
		bLocation = bucketLocation.(string)
		apiType = "ssl"
	}
	apiEndpoint := selectCosApi(apiType, bLocation)

	apiKey := rsConClient.Config.BluemixAPIKey
	s3Conf := aws.NewConfig().WithEndpoint(apiEndpoint).WithCredentials(ibmiam.NewStaticCredentials(aws.NewConfig(), authEndpoint, apiKey, serviceID)).WithS3ForcePathStyle(true)

	s3Sess := session.Must(session.NewSession())
	s3Client := s3.New(s3Sess, s3Conf)

	bucketList, _ := s3Client.ListBuckets(&s3.ListBucketsInput{})
	for _, bucket := range bucketList.Buckets {
		if bucket.Name == aws.String(bucketName) {
			return true, nil
		}
	}
	return false, nil
}

func selectCosApi(apiType string, bLocation string) string {
	if apiType == "crl" {
		switch bLocation {
		case "eu":
			return "s3.eu.cloud-object-storage.appdomain.cloud"
		case "ap":
			return "s3.ap.cloud-object-storage.appdomain.cloud"
		case "us":
			return "s3.us.cloud-object-storage.appdomain.cloud"
		}
	}
	if apiType == "rl" {
		switch bLocation {
		case "au-syd":
			return "s3.au-syd.cloud-object-storage.appdomain.cloud"
		case "eu-de":
			return "s3.eu-de.cloud-object-storage.appdomain.cloud"
		case "eu-gb":
			return "s3.eu-gb.cloud-object-storage.appdomain.cloud"
		case "jp-tok":
			return "s3.jp-tok.cloud-object-storage.appdomain.cloud"
		case "us-east":
			return "s3.us-east.cloud-object-storage.appdomain.cloud"
		case "us-south":
			return "s3.us-south.cloud-object-storage.appdomain.cloud"
		}
	}
	if apiType == "ssl" {
		switch bLocation {
		case "ams03":
			return "s3.ams03.cloud-object-storage.appdomain.cloud"
		case "che01":
			return "s3.che01.cloud-object-storage.appdomain.cloud"
		case "hkg02":
			return "s3.hkg02.cloud-object-storage.appdomain.cloud"
		case "mel01":
			return "s3.mel01.cloud-object-storage.appdomain.cloud"
		case "mex01":
			return "s3.mex01.cloud-object-storage.appdomain.cloud"
		case "mil01":
			return "s3.mil01.cloud-object-storage.appdomain.cloud"
		case "mon01":
			return "s3.mon01.cloud-object-storage.appdomain.cloud"
		case "osl01":
			return "s3.osl01.cloud-object-storage.appdomain.cloud"
		case "sjc04":
			return "s3.sjc04.cloud-object-storage.appdomain.cloud"
		case "sao01":
			return "s3.sao01.cloud-object-storage.appdomain.cloud"
		case "seo01":
			return "s3.seo01.cloud-object-storage.appdomain.cloud"
		case "tor01":
			return "s3.tor01.cloud-object-storage.appdomain.cloud"
		}
	}
	return ""
}
