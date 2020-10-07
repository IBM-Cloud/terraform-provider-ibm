package ibm

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/IBM/ibm-cos-sdk-go-config/resourceconfigurationv1"
	"github.com/IBM/ibm-cos-sdk-go/aws"
	"github.com/IBM/ibm-cos-sdk-go/aws/credentials/ibmiam"

	token "github.com/IBM/ibm-cos-sdk-go/aws/credentials/ibmiam/token"
	"github.com/IBM/ibm-cos-sdk-go/aws/session"
	"github.com/IBM/ibm-cos-sdk-go/service/s3"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var singleSiteLocation = []string{
	"ams03", "che01", "hkg02", "mel01", "mex01",
	"mil01", "mon01", "osl01", "par01", "sjc04", "sao01",
	"seo01", "sng01", "tor01",
}

var regionLocation = []string{
	"au-syd", "eu-de", "eu-fr2", "eu-gb", "jp-tok", "us-east", "us-south",
}

var crossRegionLocation = []string{
	"us", "eu", "ap",
}

var storageClass = []string{
	"standard", "vault", "cold", "flex", "smart",
}

const (
	keyAlgorithm = "AES256"
)

func resourceIBMCOS() *schema.Resource {
	return &schema.Resource{
		Read:     resourceIBMCOSRead,
		Create:   resourceIBMCOSCreate,
		Update:   resourceIBMCOSUpdate,
		Delete:   resourceIBMCOSDelete,
		Exists:   resourceIBMCOSExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(60 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"bucket_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "COS Bucket name",
			},
			"resource_instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "resource instance ID",
			},
			"crn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "CRN of resource instance",
			},
			"key_protect": {
				Type:             schema.TypeString,
				DiffSuppressFunc: applyOnce,
				Optional:         true,
				Description:      "CRN of the key you want to use data at rest encryption",
			},
			"single_site_location": {
				Type:          schema.TypeString,
				Optional:      true,
				ValidateFunc:  validateAllowedStringValue(singleSiteLocation),
				ForceNew:      true,
				ConflictsWith: []string{"region_location", "cross_region_location"},
				Description:   "single site location info",
			},
			"region_location": {
				Type:          schema.TypeString,
				Optional:      true,
				ValidateFunc:  validateAllowedStringValue(regionLocation),
				ForceNew:      true,
				ConflictsWith: []string{"cross_region_location", "single_site_location"},
				Description:   "Region Location info.",
			},
			"cross_region_location": {
				Type:          schema.TypeString,
				Optional:      true,
				ValidateFunc:  validateAllowedStringValue(crossRegionLocation),
				ForceNew:      true,
				ConflictsWith: []string{"region_location", "single_site_location"},
				Description:   "Cros region location info",
			},
			"storage_class": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAllowedStringValue(storageClass),
				ForceNew:     true,
				Description:  "Storage class info",
			},
			"endpoint_type": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateFunc:     validateAllowedStringValue([]string{"public", "private"}),
				Description:      "public or private",
				DiffSuppressFunc: applyOnce,
				Default:          "public",
			},
			"s3_endpoint_public": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Public endpoint for the COS bucket",
			},
			"s3_endpoint_private": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Private endpoint for the COS bucket",
			},
			"allowed_ip": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    false,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "List of IPv4 or IPv6 addresses ",
			},
			"activity_tracking": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Enables sending log data to Activity Tracker and LogDNA to provide visibility into object read and write events",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"read_data_events": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "If set to true, all object read events will be sent to Activity Tracker.",
						},
						"write_data_events": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "If set to true, all object write events will be sent to Activity Tracker.",
						},
						"activity_tracker_crn": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The instance of Activity Tracker that will receive object event data",
						},
					},
				},
			},
			"metrics_monitoring": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Enables sending metrics to IBM Cloud Monitoring.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"usage_metrics_enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Usage metrics will be sent to the monitoring service.",
						},
						"metrics_monitoring_crn": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Instance of IBM Cloud Monitoring that will receive the bucket metrics.",
						},
					},
				},
			},
		},
	}
}

func resourceIBMCOSUpdate(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).CosConfigV1API()
	if err != nil {
		return err
	}
	endpointType := parseBucketId(d.Id(), "endpointType")
	if endpointType == "private" {
		sess.SetServiceURL("https://config.private.cloud-object-storage.cloud.ibm.com/v1")
	}

	hasChanged := false
	updateBucketConfigOptions := &resourceconfigurationv1.UpdateBucketConfigOptions{}

	//BucketName
	bucketName := d.Get("bucket_name").(string)
	updateBucketConfigOptions.Bucket = &bucketName

	if d.HasChange("allowed_ip") {
		firewall := &resourceconfigurationv1.Firewall{}
		var ips = make([]string, 0)
		if ip, ok := d.GetOk("allowed_ip"); ok && ip != nil {
			for _, i := range ip.([]interface{}) {
				ips = append(ips, i.(string))
			}
			firewall.AllowedIp = ips
		}
		hasChanged = true
		updateBucketConfigOptions.Firewall = firewall
	}

	if d.HasChange("activity_tracking") {
		activityTracker := &resourceconfigurationv1.ActivityTracking{}
		if activity, ok := d.GetOk("activity_tracking"); ok {
			activitylist := activity.([]interface{})
			for _, l := range activitylist {
				activityMap, _ := l.(map[string]interface{})

				//Read event - as its optional check for existence
				if readEvent := activityMap["read_data_events"]; readEvent != nil {
					readSet := readEvent.(bool)
					activityTracker.ReadDataEvents = &readSet
				}

				//Write Event - as its optional check for existence
				if writeEvent := activityMap["write_data_events"]; writeEvent != nil {
					writeSet := writeEvent.(bool)
					activityTracker.WriteDataEvents = &writeSet
				}

				//crn - Required field
				crn := activityMap["activity_tracker_crn"].(string)
				activityTracker.ActivityTrackerCrn = &crn
			}
		}
		hasChanged = true
		updateBucketConfigOptions.ActivityTracking = activityTracker
	}

	if d.HasChange("metrics_monitoring") {
		metricsMonitor := &resourceconfigurationv1.MetricsMonitoring{}
		if metrics, ok := d.GetOk("metrics_monitoring"); ok {
			metricslist := metrics.([]interface{})
			for _, l := range metricslist {
				metricsMap, _ := l.(map[string]interface{})

				//metrics enabled - as its optional check for existence
				if metricsSet := metricsMap["usage_metrics_enabled"]; metricsSet != nil {
					metrics := metricsSet.(bool)
					metricsMonitor.UsageMetricsEnabled = &metrics
				}

				//crn - Required field
				crn := metricsMap["metrics_monitoring_crn"].(string)
				metricsMonitor.MetricsMonitoringCrn = &crn
			}
		}
		hasChanged = true
		updateBucketConfigOptions.MetricsMonitoring = metricsMonitor
	}

	if hasChanged {
		response, err := sess.UpdateBucketConfig(updateBucketConfigOptions)
		if err != nil {
			return fmt.Errorf("Error Update COS Bucket: %s\n%s", err, response)
		}
	}
	return resourceIBMCOSRead(d, meta)
}

func resourceIBMCOSRead(d *schema.ResourceData, meta interface{}) error {
	var s3Conf *aws.Config
	rsConClient, err := meta.(ClientSession).BluemixSession()
	if err != nil {
		return err
	}
	bucketName := parseBucketId(d.Id(), "bucketName")
	serviceID := parseBucketId(d.Id(), "serviceID")
	endpointType := parseBucketId(d.Id(), "endpointType")
	apiEndpoint, apiEndpointPrivate := selectCosApi(parseBucketId(d.Id(), "apiType"), parseBucketId(d.Id(), "bLocation"))
	if endpointType == "private" {
		apiEndpoint = apiEndpointPrivate
	}
	authEndpoint, err := rsConClient.Config.EndpointLocator.IAMEndpoint()
	if err != nil {
		return err
	}
	authEndpointPath := fmt.Sprintf("%s%s", authEndpoint, "/identity/token")
	apiKey := rsConClient.Config.BluemixAPIKey
	if apiKey != "" {
		s3Conf = aws.NewConfig().WithEndpoint(envFallBack([]string{"IBMCLOUD_COS_ENDPOINT"}, apiEndpoint)).WithCredentials(ibmiam.NewStaticCredentials(aws.NewConfig(), authEndpointPath, apiKey, serviceID)).WithS3ForcePathStyle(true)
	}
	iamAccessToken := rsConClient.Config.IAMAccessToken
	if iamAccessToken != "" {
		initFunc := func() (*token.Token, error) {
			return &token.Token{
				AccessToken:  rsConClient.Config.IAMAccessToken,
				RefreshToken: rsConClient.Config.IAMRefreshToken,
				TokenType:    "Bearer",
				ExpiresIn:    int64((time.Hour * 248).Seconds()) * -1,
				Expiration:   time.Now().Add(-1 * time.Hour).Unix(),
			}, nil
		}
		s3Conf = aws.NewConfig().WithEndpoint(envFallBack([]string{"IBMCLOUD_COS_ENDPOINT"}, apiEndpoint)).WithCredentials(ibmiam.NewCustomInitFuncCredentials(aws.NewConfig(), initFunc, authEndpointPath, serviceID)).WithS3ForcePathStyle(true)
	}
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

	bucketOutput, err := s3Client.ListBucketsExtended(&s3.ListBucketsExtendedInput{})

	if err != nil {
		return err
	}
	var bLocationConstraint string
	for _, b := range bucketOutput.Buckets {
		if *b.Name == bucketName {
			bLocationConstraint = *b.LocationConstraint
		}
	}

	singleSiteLocationRegex, err := regexp.Compile("^[a-z]{3}[0-9][0-9]-[a-z]{4,8}$")
	if err != nil {
		return err
	}
	regionLocationRegex, err := regexp.Compile("^[a-z]{2}-[a-z]{2,5}-[a-z]{4,8}$")
	if err != nil {
		return err
	}
	crossRegionLocationRegex, err := regexp.Compile("^[a-z]{2}-[a-z]{4,8}$")
	if err != nil {
		return err
	}

	if singleSiteLocationRegex.MatchString(bLocationConstraint) {
		d.Set("single_site_location", strings.Split(bLocationConstraint, "-")[0])
		d.Set("storage_class", strings.Split(bLocationConstraint, "-")[1])
	}
	if regionLocationRegex.MatchString(bLocationConstraint) {
		d.Set("region_location", fmt.Sprintf("%s-%s", strings.Split(bLocationConstraint, "-")[0], strings.Split(bLocationConstraint, "-")[1]))
		d.Set("storage_class", strings.Split(bLocationConstraint, "-")[2])
	}
	if crossRegionLocationRegex.MatchString(bLocationConstraint) {
		d.Set("cross_region_location", strings.Split(bLocationConstraint, "-")[0])
		d.Set("storage_class", strings.Split(bLocationConstraint, "-")[1])
	}

	bucketCRN := fmt.Sprintf("%s:%s:%s", strings.Replace(serviceID, "::", "", -1), "bucket", bucketName)
	d.Set("crn", bucketCRN)
	d.Set("resource_instance_id", serviceID)
	d.Set("bucket_name", bucketName)
	d.Set("s3_endpoint_public", apiEndpoint)
	d.Set("s3_endpoint_private", apiEndpointPrivate)
	if endpointType != "" {
		d.Set("endpoint_type", endpointType)
	}

	getBucketConfigOptions := &resourceconfigurationv1.GetBucketConfigOptions{
		Bucket: &bucketName,
	}

	sess, err := meta.(ClientSession).CosConfigV1API()
	if err != nil {
		return err
	}
	if endpointType == "private" {
		sess.SetServiceURL("https://config.private.cloud-object-storage.cloud.ibm.com/v1")
	}

	bucketPtr, response, err := sess.GetBucketConfig(getBucketConfigOptions)
	if err != nil {
		return fmt.Errorf("Error in getting bucket info rule: %s\n%s", err, response)
	}

	if bucketPtr != nil {

		if bucketPtr.Firewall != nil {
			d.Set("allowed_ip", flattenStringList(bucketPtr.Firewall.AllowedIp))
		}
		if bucketPtr.ActivityTracking != nil {
			d.Set("activity_tracking", flattenActivityTrack(bucketPtr.ActivityTracking))
		}
		if bucketPtr.MetricsMonitoring != nil {
			d.Set("metrics_monitoring", flattenMetricsMonitor(bucketPtr.MetricsMonitoring))
		}
	}
	return nil
}

func resourceIBMCOSCreate(d *schema.ResourceData, meta interface{}) error {
	var s3Conf *aws.Config
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
	if bLocation == "" {
		return fmt.Errorf("Provide either `cross_region_location` or `region_location` or `single_site_location`")
	}
	lConstraint := fmt.Sprintf("%s-%s", bLocation, storageClass)
	var endpointType = d.Get("endpoint_type").(string)
	apiEndpoint, privateApiEndpoint := selectCosApi(apiType, bLocation)
	if endpointType == "private" {
		apiEndpoint = privateApiEndpoint
	}
	if apiEndpoint == "" {
		return fmt.Errorf("The endpoint doesn't exists for given location %s and endpoint type %s", bLocation, endpointType)
	}
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

	authEndpoint, err := rsConClient.Config.EndpointLocator.IAMEndpoint()
	if err != nil {
		return err
	}
	authEndpointPath := fmt.Sprintf("%s%s", authEndpoint, "/identity/token")
	apiKey := rsConClient.Config.BluemixAPIKey
	if apiKey != "" {
		s3Conf = aws.NewConfig().WithEndpoint(envFallBack([]string{"IBMCLOUD_COS_ENDPOINT"}, apiEndpoint)).WithCredentials(ibmiam.NewStaticCredentials(aws.NewConfig(), authEndpointPath, apiKey, serviceID)).WithS3ForcePathStyle(true)
	}
	iamAccessToken := rsConClient.Config.IAMAccessToken
	if iamAccessToken != "" {
		initFunc := func() (*token.Token, error) {
			return &token.Token{
				AccessToken:  rsConClient.Config.IAMAccessToken,
				RefreshToken: rsConClient.Config.IAMRefreshToken,
				TokenType:    "Bearer",
				ExpiresIn:    int64((time.Hour * 248).Seconds()) * -1,
				Expiration:   time.Now().Add(-1 * time.Hour).Unix(),
			}, nil
		}
		s3Conf = aws.NewConfig().WithEndpoint(envFallBack([]string{"IBMCLOUD_COS_ENDPOINT"}, apiEndpoint)).WithCredentials(ibmiam.NewCustomInitFuncCredentials(aws.NewConfig(), initFunc, authEndpointPath, serviceID)).WithS3ForcePathStyle(true)
	}

	s3Sess := session.Must(session.NewSession())
	s3Client := s3.New(s3Sess, s3Conf)

	_, err = s3Client.CreateBucket(create)
	if err != nil {
		return err
	}
	// Generating a fake id which contains every information about to get the bucket via s3 api
	bucketID := fmt.Sprintf("%s:%s:%s:meta:%s:%s:%s", strings.Replace(serviceID, "::", "", -1), "bucket", bucketName, apiType, bLocation, endpointType)
	d.SetId(bucketID)

	return resourceIBMCOSUpdate(d, meta)

}

func resourceIBMCOSDelete(d *schema.ResourceData, meta interface{}) error {
	var s3Conf *aws.Config
	rsConClient, _ := meta.(ClientSession).BluemixSession()
	bucketName := parseBucketId(d.Id(), "bucketName")
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
	endpointType := parseBucketId(d.Id(), "endpointType")
	apiEndpoint, apiEndpointPrivate := selectCosApi(apiType, bLocation)
	if endpointType == "private" {
		apiEndpoint = apiEndpointPrivate
	}
	authEndpoint, err := rsConClient.Config.EndpointLocator.IAMEndpoint()
	if err != nil {
		return err
	}
	authEndpointPath := fmt.Sprintf("%s%s", authEndpoint, "/identity/token")

	apiKey := rsConClient.Config.BluemixAPIKey
	if apiKey != "" {
		s3Conf = aws.NewConfig().WithEndpoint(envFallBack([]string{"IBMCLOUD_COS_ENDPOINT"}, apiEndpoint)).WithCredentials(ibmiam.NewStaticCredentials(aws.NewConfig(), authEndpointPath, apiKey, serviceID)).WithS3ForcePathStyle(true)
	}
	iamAccessToken := rsConClient.Config.IAMAccessToken
	if iamAccessToken != "" {
		initFunc := func() (*token.Token, error) {
			return &token.Token{
				AccessToken:  rsConClient.Config.IAMAccessToken,
				RefreshToken: rsConClient.Config.IAMRefreshToken,
				TokenType:    "Bearer",
				ExpiresIn:    int64((time.Hour * 248).Seconds()) * -1,
				Expiration:   time.Now().Add(-1 * time.Hour).Unix(),
			}, nil
		}
		s3Conf = aws.NewConfig().WithEndpoint(envFallBack([]string{"IBMCLOUD_COS_ENDPOINT"}, apiEndpoint)).WithCredentials(ibmiam.NewCustomInitFuncCredentials(aws.NewConfig(), initFunc, authEndpointPath, serviceID)).WithS3ForcePathStyle(true)
	}

	s3Sess := session.Must(session.NewSession())
	s3Client := s3.New(s3Sess, s3Conf)

	delete := &s3.DeleteBucketInput{
		Bucket: aws.String(bucketName),
	}
	_, err = s3Client.DeleteBucket(delete)
	if err != nil {
		return err
	}
	d.SetId("")
	return nil
}

func resourceIBMCOSExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	var s3Conf *aws.Config
	rsConClient, err := meta.(ClientSession).BluemixSession()
	if err != nil {
		return false, err
	}

	bucketName := parseBucketId(d.Id(), "bucketName")
	serviceID := parseBucketId(d.Id(), "serviceID")
	endpointType := parseBucketId(d.Id(), "endpointType")
	apiEndpoint, apiEndpointPrivate := selectCosApi(parseBucketId(d.Id(), "apiType"), parseBucketId(d.Id(), "bLocation"))
	if endpointType == "private" {
		apiEndpoint = apiEndpointPrivate
	}

	authEndpoint, err := rsConClient.Config.EndpointLocator.IAMEndpoint()
	if err != nil {
		return false, err
	}
	authEndpointPath := fmt.Sprintf("%s%s", authEndpoint, "/identity/token")

	apiKey := rsConClient.Config.BluemixAPIKey
	if apiKey != "" {
		s3Conf = aws.NewConfig().WithEndpoint(envFallBack([]string{"IBMCLOUD_COS_ENDPOINT"}, apiEndpoint)).WithCredentials(ibmiam.NewStaticCredentials(aws.NewConfig(), authEndpointPath, apiKey, serviceID)).WithS3ForcePathStyle(true)
	}
	iamAccessToken := rsConClient.Config.IAMAccessToken
	if iamAccessToken != "" {
		initFunc := func() (*token.Token, error) {
			return &token.Token{
				AccessToken:  rsConClient.Config.IAMAccessToken,
				RefreshToken: rsConClient.Config.IAMRefreshToken,
				TokenType:    "Bearer",
				ExpiresIn:    int64((time.Hour * 248).Seconds()) * -1,
				Expiration:   time.Now().Add(-1 * time.Hour).Unix(),
			}, nil
		}
		s3Conf = aws.NewConfig().WithEndpoint(envFallBack([]string{"IBMCLOUD_COS_ENDPOINT"}, apiEndpoint)).WithCredentials(ibmiam.NewCustomInitFuncCredentials(aws.NewConfig(), initFunc, authEndpointPath, serviceID)).WithS3ForcePathStyle(true)
	}

	s3Sess := session.Must(session.NewSession())
	s3Client := s3.New(s3Sess, s3Conf)

	bucketList, err := s3Client.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		return false, err
	}
	for _, bucket := range bucketList.Buckets {
		if *bucket.Name == bucketName {
			return true, nil
		}
	}
	return false, nil
}

func selectCosApi(apiType string, bLocation string) (string, string) {
	if apiType == "crl" {
		switch bLocation {
		case "eu":
			return "s3.eu.cloud-object-storage.appdomain.cloud", "s3.private.eu.cloud-object-storage.appdomain.cloud"
		case "ap":
			return "s3.ap.cloud-object-storage.appdomain.cloud", "s3.private.ap.cloud-object-storage.appdomain.cloud"
		case "us":
			return "s3.us.cloud-object-storage.appdomain.cloud", "s3.private.us.cloud-object-storage.appdomain.cloud"
		}
	}
	if apiType == "rl" {
		switch bLocation {
		case "au-syd":
			return "s3.au-syd.cloud-object-storage.appdomain.cloud", "s3.private.au-syd.cloud-object-storage.appdomain.cloud"
		case "eu-de":
			return "s3.eu-de.cloud-object-storage.appdomain.cloud", "s3.private.eu-de.cloud-object-storage.appdomain.cloud"
		case "eu-fr2":
			return "", "s3.private.eu-fr2.cloud-object-storage.appdomain.cloud"
		case "eu-gb":
			return "s3.eu-gb.cloud-object-storage.appdomain.cloud", "s3.private.eu-gb.cloud-object-storage.appdomain.cloud"
		case "jp-tok":
			return "s3.jp-tok.cloud-object-storage.appdomain.cloud", "s3.private.jp-tok.cloud-object-storage.appdomain.cloud"
		case "us-east":
			return "s3.us-east.cloud-object-storage.appdomain.cloud", "s3.private.us-east.cloud-object-storage.appdomain.cloud"
		case "us-south":
			return "s3.us-south.cloud-object-storage.appdomain.cloud", "s3.private.us-south.cloud-object-storage.appdomain.cloud"
		}
	}
	if apiType == "ssl" {
		switch bLocation {
		case "ams03":
			return "s3.ams03.cloud-object-storage.appdomain.cloud", "s3.private.ams03.cloud-object-storage.appdomain.cloud"
		case "che01":
			return "s3.che01.cloud-object-storage.appdomain.cloud", "s3.private.che01.cloud-object-storage.appdomain.cloud"
		case "hkg02":
			return "s3.hkg02.cloud-object-storage.appdomain.cloud", "s3.private.hkg02.cloud-object-storage.appdomain.cloud"
		case "mel01":
			return "s3.mel01.cloud-object-storage.appdomain.cloud", "s3.private.mel01.cloud-object-storage.appdomain.cloud"
		case "mex01":
			return "s3.mex01.cloud-object-storage.appdomain.cloud", "s3.private.mex01.cloud-object-storage.appdomain.cloud"
		case "mil01":
			return "s3.mil01.cloud-object-storage.appdomain.cloud", "s3.private.mil01.cloud-object-storage.appdomain.cloud"
		case "mon01":
			return "s3.mon01.cloud-object-storage.appdomain.cloud", "s3.private.mon01.cloud-object-storage.appdomain.cloud"
		case "osl01":
			return "s3.osl01.cloud-object-storage.appdomain.cloud", "s3.private.osl01.cloud-object-storage.appdomain.cloud"
		case "par01":
			return "s3.par01.cloud-object-storage.appdomain.cloud", "s3.private.par01.cloud-object-storage.appdomain.cloud"
		case "sjc04":
			return "s3.sjc04.cloud-object-storage.appdomain.cloud", "s3.private.sjc04.cloud-object-storage.appdomain.cloud"
		case "sao01":
			return "s3.sao01.cloud-object-storage.appdomain.cloud", "s3.private.sao01.cloud-object-storage.appdomain.cloud"
		case "seo01":
			return "s3.seo01.cloud-object-storage.appdomain.cloud", "s3.private.seo01.cloud-object-storage.appdomain.cloud"
		case "sng01":
			return "s3.sng01.cloud-object-storage.appdomain.cloud", "s3.private.sng01.cloud-object-storage.appdomain.cloud"
		case "tor01":
			return "s3.tor01.cloud-object-storage.appdomain.cloud", "s3.private.tor01.cloud-object-storage.appdomain.cloud"
		}
	}
	return "", ""
}

func parseBucketId(id string, info string) string {
	crn := strings.Split(id, ":meta:")[0]
	meta := strings.Split(id, ":meta:")[1]

	if info == "bucketName" {
		return strings.Split(crn, ":bucket:")[1]
	}
	if info == "serviceID" {
		return fmt.Sprintf("%s::", strings.Split(crn, ":bucket:")[0])
	}
	if info == "apiType" {
		return strings.Split(meta, ":")[0]
	}
	if info == "bLocation" {
		return strings.Split(meta, ":")[1]
	}
	if info == "endpointType" {
		s := strings.Split(meta, ":")
		if len(s) > 2 {
			return strings.Split(meta, ":")[2]
		}
		return ""

	}
	return ""
}
