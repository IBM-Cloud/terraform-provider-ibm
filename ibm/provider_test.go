package ibm

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var cfOrganization string
var cfSpace string
var cisDomainStatic string
var cisDomainTest string
var cisInstance string
var cisResourceGroup string
var ibmid1 string
var ibmid2 string
var IAMUser string
var datacenter string
var machineType string
var trustedMachineType string
var publicVlanID string
var privateVlanID string
var privateSubnetID string
var publicSubnetID string
var subnetID string
var lbaasDatacenter string
var lbaasSubnetId string
var ipsecDatacenter string
var customersubnetid string
var customerpeerip string
var dedicatedHostName string
var dedicatedHostID string
var kubeVersion string
var kubeUpdateVersion string
var zone string
var zonePrivateVlan string
var zonePublicVlan string
var zoneUpdatePrivateVlan string
var zoneUpdatePublicVlan string
var csRegion string
var extendedHardwareTesting bool
var err error
var placementGroupName string
var certCRN string
var updatedCertCRN string
var regionName string
var ISZoneName string
var ISCIDR string
var ISAddressPrefixCIDR string
var instanceProfileName string
var ISRouteDestination string
var ISRouteNextHop string
var workspaceID string
var templateID string
var imageName string
var functionNamespace string
var hpcsInstanceID string

// For Power Colo

var pi_image string
var pi_key_name string
var pi_volume_name string
var pi_network_name string
var pi_cloud_instance_id string
var pi_instance_name string

// For Image

var IsImageName string
var isImage string
var IsImageEncryptedDataKey string
var IsImageEncryptionKey string
var isWinImage string
var image_cos_url string
var image_cos_url_encrypted string
var image_operating_system string

// Transit Gateway cross account
var tg_cross_network_account_id string
var tg_cross_network_id string

//

func init() {
	cfOrganization = os.Getenv("IBM_ORG")
	if cfOrganization == "" {
		fmt.Println("[WARN] Set the environment variable IBM_ORG for testing ibm_org  resource Some tests for that resource will fail if this is not set correctly")
	}
	cfSpace = os.Getenv("IBM_SPACE")
	if cfSpace == "" {
		fmt.Println("[WARN] Set the environment variable IBM_SPACE for testing ibm_space  resource Some tests for that resource will fail if this is not set correctly")
	}
	ibmid1 = os.Getenv("IBM_ID1")
	if ibmid1 == "" {
		fmt.Println("[WARN] Set the environment variable IBM_ID1 for testing ibm_space resource Some tests for that resource will fail if this is not set correctly")
	}

	ibmid2 = os.Getenv("IBM_ID2")
	if ibmid2 == "" {
		fmt.Println("[WARN] Set the environment variable IBM_ID2 for testing ibm_space resource Some tests for that resource will fail if this is not set correctly")
	}

	IAMUser = os.Getenv("IBM_IAMUSER")
	if IAMUser == "" {
		fmt.Println("[WARN] Set the environment variable IBM_IAMUSER for testing ibm_iam_user_policy resource Some tests for that resource will fail if this is not set correctly")
	}

	datacenter = os.Getenv("IBM_DATACENTER")
	if datacenter == "" {
		datacenter = "par01"
		fmt.Println("[WARN] Set the environment variable IBM_DATACENTER for testing ibm_container_cluster resource else it is set to default value 'par01'")
	}

	machineType = os.Getenv("IBM_MACHINE_TYPE")
	if machineType == "" {
		machineType = "b3c.4x16"
		fmt.Println("[WARN] Set the environment variable IBM_MACHINE_TYPE for testing ibm_container_cluster resource else it is set to default value 'u2c.2x4'")
	}

	certCRN = os.Getenv("IBM_CERT_CRN")
	if certCRN == "" {
		certCRN = "crn:v1:bluemix:public:cloudcerts:us-south:a/52b2e14f385aca5da781baa1b9c28e53:6efac0c2-b955-49ca-939d-d7bc0cb8132f:certificate:e786b0ea2af8b5435603803ec2ff8118"
		fmt.Println("[WARN] Set the environment variable IBM_CERT_CRN for testing ibm_container_alb_cert resource else it is set to default value")
	}

	updatedCertCRN = os.Getenv("IBM_UPDATE_CERT_CRN")
	if updatedCertCRN == "" {
		updatedCertCRN = "crn:v1:bluemix:public:cloudcerts:eu-de:a/e9021a4d06e9b108b4a221a3cec47e3d:77e527aa-65b2-4cb3-969b-7e8714174346:certificate:1bf3d0c2b7764402dde25744218e6cba"
		fmt.Println("[WARN] Set the environment variable IBM_UPDATE_CERT_CRN for testing ibm_container_alb_cert resource else it is set to default value")
	}

	csRegion = os.Getenv("IBM_CONTAINER_REGION")
	if csRegion == "" {
		csRegion = "eu-de"
		fmt.Println("[WARN] Set the environment variable IBM_CONTAINER_REGION for testing ibm_container resources else it is set to default value 'eu-de'")
	}

	cisInstance = os.Getenv("IBM_CIS_INSTANCE")
	if cisInstance == "" {
		cisInstance = ""
		fmt.Println("[WARN] Set the environment variable IBM_CIS_INSTANCE with a VALID CIS Instance NAME for testing ibm_cis resources on staging/test")
	}
	cisDomainStatic = os.Getenv("IBM_CIS_DOMAIN_STATIC")
	if cisDomainStatic == "" {
		cisDomainStatic = ""
		fmt.Println("[WARN] Set the environment variable IBM_CIS_DOMAIN_STATIC with the Domain name registered with the CIS instance on test/staging. Domain must be predefined in CIS to avoid CIS billing costs due to domain delete/create")
	}

	cisDomainTest = os.Getenv("IBM_CIS_DOMAIN_TEST")
	if cisDomainTest == "" {
		cisDomainTest = ""
		fmt.Println("[WARN] Set the environment variable IBM_CIS_DOMAIN_TEST with a VALID Domain name for testing the one time create and delete of a domain in CIS. Note each create/delete will trigger a monthly billing instance. Only to be run in staging/test")
	}

	cisResourceGroup = os.Getenv("IBM_CIS_RESOURCE_GROUP")
	if cisResourceGroup == "" {
		cisResourceGroup = ""
		fmt.Println("[WARN] Set the environment variable IBM_CIS_RESOURCE_GROUP with the resource group for the CIS Instance ")
	}

	trustedMachineType = os.Getenv("IBM_TRUSTED_MACHINE_TYPE")
	if trustedMachineType == "" {
		trustedMachineType = "mb1c.16x64"
		fmt.Println("[WARN] Set the environment variable IBM_TRUSTED_MACHINE_TYPE for testing ibm_container_cluster resource else it is set to default value 'mb1c.16x64'")
	}

	extendedHardwareTesting, err = strconv.ParseBool(os.Getenv("IBM_BM_EXTENDED_HW_TESTING"))
	if err != nil {
		extendedHardwareTesting = false
		fmt.Println("[WARN] Set the environment variable IBM_BM_EXTENDED_HW_TESTING to true/false for testing ibm_compute_bare_metal resource else it is set to default value 'false'")
	}

	publicVlanID = os.Getenv("IBM_PUBLIC_VLAN_ID")
	if publicVlanID == "" {
		publicVlanID = "2393319"
		fmt.Println("[WARN] Set the environment variable IBM_PUBLIC_VLAN_ID for testing ibm_container_cluster resource else it is set to default value '2393319'")
	}

	privateVlanID = os.Getenv("IBM_PRIVATE_VLAN_ID")
	if privateVlanID == "" {
		privateVlanID = "2393321"
		fmt.Println("[WARN] Set the environment variable IBM_PRIVATE_VLAN_ID for testing ibm_container_cluster resource else it is set to default value '2393321'")
	}

	kubeVersion = os.Getenv("IBM_KUBE_VERSION")
	if kubeVersion == "" {
		kubeVersion = "1.17.9"
		fmt.Println("[WARN] Set the environment variable IBM_KUBE_VERSION for testing ibm_container_cluster resource else it is set to default value '1.13.10'")
	}

	kubeUpdateVersion = os.Getenv("IBM_KUBE_UPDATE_VERSION")
	if kubeUpdateVersion == "" {
		kubeUpdateVersion = "1.14.6"
		fmt.Println("[WARN] Set the environment variable IBM_KUBE_UPDATE_VERSION for testing ibm_container_cluster resource else it is set to default value '1.14.6'")
	}

	privateSubnetID = os.Getenv("IBM_PRIVATE_SUBNET_ID")
	if privateSubnetID == "" {
		privateSubnetID = "1636107"
		fmt.Println("[WARN] Set the environment variable IBM_PRIVATE_SUBNET_ID for testing ibm_container_cluster resource else it is set to default value '1636107'")
	}

	publicSubnetID = os.Getenv("IBM_PUBLIC_SUBNET_ID")
	if publicSubnetID == "" {
		publicSubnetID = "1165645"
		fmt.Println("[WARN] Set the environment variable IBM_PUBLIC_SUBNET_ID for testing ibm_container_cluster resource else it is set to default value '1165645'")
	}

	subnetID = os.Getenv("IBM_SUBNET_ID")
	if subnetID == "" {
		subnetID = "1165645"
		fmt.Println("[WARN] Set the environment variable IBM_SUBNET_ID for testing ibm_container_cluster resource else it is set to default value '1165645'")
	}

	ipsecDatacenter = os.Getenv("IBM_IPSEC_DATACENTER")
	if ipsecDatacenter == "" {
		ipsecDatacenter = "tok02"
		fmt.Println("[INFO] Set the environment variable IBM_IPSEC_DATACENTER for testing ibm_ipsec_vpn resource else it is set to default value 'tok02'")
	}

	customersubnetid = os.Getenv("IBM_IPSEC_CUSTOMER_SUBNET_ID")
	if customersubnetid == "" {
		customersubnetid = "123456"
		fmt.Println("[INFO] Set the environment variable IBM_IPSEC_CUSTOMER_SUBNET_ID for testing ibm_ipsec_vpn resource else it is set to default value '123456'")
	}

	customerpeerip = os.Getenv("IBM_IPSEC_CUSTOMER_PEER_IP")
	if customerpeerip == "" {
		customerpeerip = "192.168.0.1"
		fmt.Println("[INFO] Set the environment variable IBM_IPSEC_CUSTOMER_PEER_IP for testing ibm_ipsec_vpn resource else it is set to default value '192.168.0.1'")
	}

	lbaasDatacenter = os.Getenv("IBM_LBAAS_DATACENTER")
	if lbaasDatacenter == "" {
		lbaasDatacenter = "wdc04"
		fmt.Println("[WARN] Set the environment variable IBM_LBAAS_DATACENTER for testing ibm_lbaas resource else it is set to default value 'wdc04'")
	}

	lbaasSubnetId = os.Getenv("IBM_LBAAS_SUBNETID")
	if lbaasSubnetId == "" {
		lbaasSubnetId = "1511875"
		fmt.Println("[WARN] Set the environment variable IBM_LBAAS_SUBNETID for testing ibm_lbaas resource else it is set to default value '1511875'")
	}

	dedicatedHostName = os.Getenv("IBM_DEDICATED_HOSTNAME")
	if dedicatedHostName == "" {
		dedicatedHostName = "terraform-dedicatedhost"
		fmt.Println("[WARN] Set the environment variable IBM_DEDICATED_HOSTNAME for testing ibm_compute_vm_instance resource else it is set to default value 'terraform-dedicatedhost'")
	}

	dedicatedHostID = os.Getenv("IBM_DEDICATED_HOST_ID")
	if dedicatedHostID == "" {
		dedicatedHostID = "30301"
		fmt.Println("[WARN] Set the environment variable IBM_DEDICATED_HOST_ID for testing ibm_compute_vm_instance resource else it is set to default value '30301'")
	}

	zone = os.Getenv("IBM_WORKER_POOL_ZONE")
	if zone == "" {
		zone = "ams03"
		fmt.Println("[WARN] Set the environment variable IBM_WORKER_POOL_ZONE for testing ibm_container_worker_pool_zone_attachment resource else it is set to default value 'ams03'")
	}

	zonePrivateVlan = os.Getenv("IBM_WORKER_POOL_ZONE_PRIVATE_VLAN")
	if zonePrivateVlan == "" {
		zonePrivateVlan = "2538975"
		fmt.Println("[WARN] Set the environment variable IBM_WORKER_POOL_ZONE_PRIVATE_VLAN for testing ibm_container_worker_pool_zone_attachment resource else it is set to default value '2538975'")
	}

	zonePublicVlan = os.Getenv("IBM_WORKER_POOL_ZONE_PUBLIC_VLAN")
	if zonePublicVlan == "" {
		zonePublicVlan = "2538967"
		fmt.Println("[WARN] Set the environment variable IBM_WORKER_POOL_ZONE_PUBLIC_VLAN for testing ibm_container_worker_pool_zone_attachment resource else it is set to default value '2538967'")
	}

	zoneUpdatePrivateVlan = os.Getenv("IBM_WORKER_POOL_ZONE_UPDATE_PRIVATE_VLAN")
	if zoneUpdatePrivateVlan == "" {
		zoneUpdatePrivateVlan = "2388377"
		fmt.Println("[WARN] Set the environment variable IBM_WORKER_POOL_ZONE_UPDATE_PRIVATE_VLAN for testing ibm_container_worker_pool_zone_attachment resource else it is set to default value '2388377'")
	}

	zoneUpdatePublicVlan = os.Getenv("IBM_WORKER_POOL_ZONE_UPDATE_PUBLIC_VLAN")
	if zoneUpdatePublicVlan == "" {
		zoneUpdatePublicVlan = "2388375"
		fmt.Println("[WARN] Set the environment variable IBM_WORKER_POOL_ZONE_UPDATE_PUBLIC_VLAN for testing ibm_container_worker_pool_zone_attachment resource else it is set to default value '2388375'")
	}

	placementGroupName = os.Getenv("IBM_PLACEMENT_GROUP_NAME")
	if placementGroupName == "" {
		placementGroupName = "terraform_group"
		fmt.Println("[WARN] Set the environment variable IBM_PLACEMENT_GROUP_NAME for testing ibm_compute_vm_instance resource else it is set to default value 'terraform-group'")
	}

	regionName = os.Getenv("SL_REGION")
	if regionName == "" {
		regionName = "us-south"
		fmt.Println("[INFO] Set the environment variable SL_REGION for testing ibm_is_region datasource else it is set to default value 'us-south'")
	}

	ISZoneName = os.Getenv("SL_ZONE")
	if ISZoneName == "" {
		ISZoneName = "us-south-1"
		fmt.Println("[INFO] Set the environment variable SL_ZONE for testing ibm_is_zone datasource else it is set to default value 'us-south-1'")
	}

	ISCIDR = os.Getenv("SL_CIDR")
	if ISCIDR == "" {
		ISCIDR = "10.240.0.0/24"
		fmt.Println("[INFO] Set the environment variable SL_CIDR for testing ibm_is_subnet else it is set to default value '10.240.0.0/24'")
	}

	ISAddressPrefixCIDR = os.Getenv("SL_ADDRESS_PREFIX_CIDR")
	if ISAddressPrefixCIDR == "" {
		ISAddressPrefixCIDR = "10.120.0.0/24"
		fmt.Println("[INFO] Set the environment variable SL_ADDRESS_PREFIX_CIDR for testing ibm_is_vpc_address_prefix else it is set to default value '10.120.0.0/24'")
	}

	isImage = os.Getenv("IS_IMAGE")
	if isImage == "" {
		//isImage = "fc538f61-7dd6-4408-978c-c6b85b69fe76" // for classic infrastructure
		isImage = "r006-ed3f775f-ad7e-4e37-ae62-7199b4988b00" // for next gen infrastructure
		fmt.Println("[INFO] Set the environment variable IS_IMAGE for testing ibm_is_instance, ibm_is_floating_ip else it is set to default value 'r006-ed3f775f-ad7e-4e37-ae62-7199b4988b00'")
	}

	isWinImage = os.Getenv("IS_WIN_IMAGE")
	if isWinImage == "" {
		//isWinImage = "a7a0626c-f97e-4180-afbe-0331ec62f32a" // classic windows machine: ibm-windows-server-2012-full-standard-amd64-1
		isWinImage = "r006-5f9568ae-792e-47e1-a710-5538b2bdfca7" // next gen windows machine: ibm-windows-server-2012-full-standard-amd64-3
		fmt.Println("[INFO] Set the environment variable IS_WIN_IMAGE for testing ibm_is_instance data source else it is set to default value 'r006-5f9568ae-792e-47e1-a710-5538b2bdfca7'")
	}

	instanceProfileName = os.Getenv("SL_INSTANCE_PROFILE")
	if instanceProfileName == "" {
		//instanceProfileName = "bc1-2x8" // for classic infrastructure
		instanceProfileName = "cx2-2x4" // for next gen infrastructure
		fmt.Println("[INFO] Set the environment variable SL_INSTANCE_PROFILE for testing ibm_is_instance resource else it is set to default value 'b-2x8'")
	}

	ISRouteDestination = os.Getenv("SL_ROUTE_DESTINATION")
	if ISRouteDestination == "" {
		ISRouteDestination = "192.168.4.0/24"
		fmt.Println("[INFO] Set the environment variable SL_ROUTE_DESTINATION for testing ibm_is_vpc_route else it is set to default value '192.168.4.0/24'")
	}

	ISRouteNextHop = os.Getenv("SL_ROUTE_NEXTHOP")
	if ISRouteNextHop == "" {
		ISRouteNextHop = "10.240.0.0"
		fmt.Println("[INFO] Set the environment variable SL_ROUTE_NEXTHOP for testing ibm_is_vpc_route else it is set to default value '10.0.0.4'")
	}

	// Added for Power Colo Testing
	pi_image = os.Getenv("PI_IMAGE")
	if pi_image == "" {
		pi_image = "7200-03-03"
		fmt.Println("[INFO] Set the environment variable PI_IMAGE for testing ibm_pi_image resource else it is set to default value '7200-03-03'")
	}

	pi_key_name = os.Getenv("PI_KEY_NAME")
	if pi_key_name == "" {
		pi_key_name = "brampoc"
		fmt.Println("[INFO] Set the environment variable PI_KEY_NAME for testing ibm_pi_key_name resource else it is set to default value 'brampoc'")
	}

	pi_network_name = os.Getenv("PI_NETWORK_NAME")
	if pi_network_name == "" {
		pi_network_name = "APP"
		fmt.Println("[INFO] Set the environment variable PI_NETWORK_NAME for testing ibm_pi_network_name resource else it is set to default value 'APP'")
	}

	pi_volume_name = os.Getenv("PI_VOLUME_NAME")
	if pi_volume_name == "" {
		pi_volume_name = "vg9"
		fmt.Println("[INFO] Set the environment variable PI_VOLUME_NAME for testing ibm_pi_network_name resource else it is set to default value 'vg9'")
	}

	pi_cloud_instance_id = os.Getenv("PI_CLOUDINSTANCE_ID")
	if pi_cloud_instance_id == "" {
		//pi_cloud_instance_id = "d16705bd-7f1a-48c9-9e0e-1c17b71e7331"
		pi_cloud_instance_id = "db0e5c7d-a708-454b-8a52-f0f2f2771075"
		fmt.Println("[INFO] Set the environment variable PI_CLOUDINSTANCE_ID for testing ibm_pi_image resource else it is set to default value 'd16705bd-7f1a-48c9-9e0e-1c17b71e7331'")
	}
	workspaceID = os.Getenv("WORKSPACE_ID")
	if workspaceID == "" {
		workspaceID = "outwork-2737f163-b966-44"
		fmt.Println("[INFO] Set the environment variable WORKSPACE_ID for testing data_source_ibm_schematics_state_test else it is set to default value")
	}
	templateID = os.Getenv("TEMPLATE_ID")
	if templateID == "" {
		templateID = "653f60a4-f64f-41"
		fmt.Println("[INFO] Set the environment variable TEMPLATE_ID for testing data_source_ibm_schematics_state_test else it is set to default value")
	}

	// Added for resource image testing
	image_cos_url = os.Getenv("IMAGE_COS_URL")
	if image_cos_url == "" {
		image_cos_url = "cos://us-south/cosbucket-vpc-image-gen2/rhel-guest-image-7.0-20140930.0.x86_64.qcow2"
		fmt.Println("[WARN] Set the environment variable IMAGE_COS_URL with a VALID COS Image SQL URL for testing ibm_is_image resources on staging/test")
	}

	// Added for resource image testing
	image_cos_url_encrypted = os.Getenv("IMAGE_COS_URL_ENCRYPTED")
	if image_cos_url_encrypted == "" {
		image_cos_url_encrypted = "cos://us-south/cosbucket-vpc-image-gen2/rhel-guest-image-7.0-encrypted.qcow2"
		fmt.Println("[WARN] Set the environment variable IMAGE_COS_URL_ENCRYPTED with a VALID COS Image SQL URL for testing ibm_is_image resources on staging/test")
	}
	image_operating_system = os.Getenv("IMAGE_OPERATING_SYSTEM")
	if image_operating_system == "" {
		image_operating_system = "red-7-amd64"
		fmt.Println("[WARN] Set the environment variable IMAGE_OPERATING_SYSTEM with a VALID Operating system for testing ibm_is_image resources on staging/test")
	}

	IsImageName = os.Getenv("IS_IMAGE_NAME")
	if IsImageName == "" {
		//IsImageName = "ibm-ubuntu-18-04-2-minimal-amd64-1" // for classic infrastructure
		IsImageName = "ibm-ubuntu-18-04-1-minimal-amd64-2" // for next gen infrastructure
		fmt.Println("[INFO] Set the environment variable IS_IMAGE_NAME for testing data source ibm_is_image else it is set to default value `ibm-ubuntu-18-04-1-minimal-amd64-2`")
	}
	IsImageEncryptedDataKey = os.Getenv("IS_IMAGE_ENCRYPTED_DATA_KEY")
	if IsImageEncryptedDataKey == "" {
		IsImageEncryptedDataKey = "eyJjaXBoZXJ0ZXh0IjoidElsZnRjUXB5L0krSGJsMlVIK2ZxZ1FGK1diR3loV1dPRFk9IiwiaXYiOiJ3SlhSVklsSHUzMzFqUEY0IiwidmVyc2lvbiI6IjQuMC4wIiwiaGFuZGxlIjoiZjM2YTA2NGUtY2E2My00NmU0LThlNjAtYmJiMzEyNTY5YzM1In0="
		fmt.Println("[INFO] Set the environment variable IS_IMAGE_ENCRYPTED_DATA_KEY for testing resource ibm_is_image else it is set to default value")
	}
	IsImageEncryptionKey = os.Getenv("IS_IMAGE_ENCRYPTION_KEY")
	if IsImageEncryptionKey == "" {
		IsImageEncryptionKey = "crn:v1:bluemix:public:kms:us-south:a/52b2e14f385aca5da781baa1b9c28e53:21d9f13d-5895-49a1-9e80-b4aff69dfc1f:key:f36a064e-ca63-46e4-8e60-bbb312569c35"
		fmt.Println("[INFO] Set the environment variable IS_IMAGE_ENCRYPTION_KEY for testing resource ibm_is_image else it is set to default value")
	}

	functionNamespace = os.Getenv("IBM_FUNCTION_NAMESPACE")
	if functionNamespace == "" {
		fmt.Println("[INFO] Set the environment variable IBM_FUNCTION_NAMESPACE for testing ibm_function_package, ibm_function_action, ibm_function_rule, ibm_function_trigger resource else  tests will fail if this is not set correctly")
	}

	hpcsInstanceID = os.Getenv("HPCS_INSTANCE_ID")
	if hpcsInstanceID == "" {
		hpcsInstanceID = "5af62d5d-5d90-4b84-bbcd-90d2123ae6c8"
		fmt.Println("[INFO] Set the environment variable HPCS_INSTANCE_ID for testing data_source_ibm_kms_key_test else it is set to default value")
	}
	tg_cross_network_account_id = os.Getenv("IBM_TG_CROSS_ACCOUNT_ID")
	if tg_cross_network_account_id == "" {
		fmt.Println("[INFO] Set the environment variable IBM_TG_CROSS_ACCOUNT_ID for testing ibm_tg_connection resource else  tests will fail if this is not set correctly")
	}
	tg_cross_network_id = os.Getenv("IBM_TG_CROSS_NETWORK_ID")
	if tg_cross_network_id == "" {
		fmt.Println("[INFO] Set the environment variable IBM_TG_CROSS_NETWORK_ID for testing ibm_tg_connection resource else  tests will fail if this is not set correctly")
	}

}

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"ibm": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("IC_API_KEY"); v == "" {
		t.Fatal("IC_API_KEY must be set for acceptance tests")
	}
	if v := os.Getenv("IAAS_CLASSIC_API_KEY"); v == "" {
		t.Fatal("IAAS_CLASSIC_API_KEY must be set for acceptance tests")
	}
	if v := os.Getenv("IAAS_CLASSIC_USERNAME"); v == "" {
		t.Fatal("IAAS_CLASSIC_USERNAME must be set for acceptance tests")
	}
}

func testAccPreCheckCis(t *testing.T) {
	testAccPreCheck(t)
	if cisInstance == "" {
		t.Fatal("IBM_CIS_INSTANCE must be set for acceptance tests")
	}
	if cisResourceGroup == "" {
		t.Fatal("IBM_CIS_RESOURCE_GROUP must be set for acceptance tests")
	}
	if cisDomainStatic == "" {
		t.Fatal("IBM_CIS_DOMAIN_STATIC must be set for acceptance tests")
	}
	if cisDomainTest == "" {
		t.Fatal("IBM_CIS_DOMAIN_TEST must be set for acceptance tests")
	}
}

func testAccPreCheckImage(t *testing.T) {
	testAccPreCheck(t)
	if image_cos_url == "" {
		t.Fatal("IMAGE_COS_URL must be set for acceptance tests")
	}
	if image_operating_system == "" {
		t.Fatal("IMAGE_OPERATING_SYSTEM must be set for acceptance tests")
	}
}
func testAccPreCheckEncryptedImage(t *testing.T) {
	testAccPreCheck(t)
	if image_cos_url_encrypted == "" {
		t.Fatal("IMAGE_COS_URL_ENCRYPTED must be set for acceptance tests")
	}
	if image_operating_system == "" {
		t.Fatal("IMAGE_OPERATING_SYSTEM must be set for acceptance tests")
	}
	if IsImageEncryptedDataKey == "" {
		t.Fatal("IS_IMAGE_ENCRYPTED_DATA_KEY must be set for acceptance tests")
	}
	if IsImageEncryptionKey == "" {
		t.Fatal("IS_IMAGE_ENCRYPTION_KEY must be set for acceptance tests")
	}
}
