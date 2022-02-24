// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package acctest

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var AppIDTenantID string
var AppIDTestUserEmail string
var CfOrganization string
var CfSpace string
var CisDomainStatic string
var CisDomainTest string
var CisInstance string
var CisResourceGroup string
var CloudShellAccountID string
var CosCRN string
var Ibmid1 string
var Ibmid2 string
var IAMUser string
var Datacenter string
var MachineType string
var trustedMachineType string
var PublicVlanID string
var PrivateVlanID string
var PrivateSubnetID string
var PublicSubnetID string
var SubnetID string
var LbaasDatacenter string
var LbaasSubnetId string
var LbListerenerCertificateInstance string
var IpsecDatacenter string
var Customersubnetid string
var Customerpeerip string
var DedicatedHostName string
var DedicatedHostID string
var KubeVersion string
var KubeUpdateVersion string
var Zone string
var ZonePrivateVlan string
var ZonePublicVlan string
var ZoneUpdatePrivateVlan string
var ZoneUpdatePublicVlan string
var CsRegion string
var ExtendedHardwareTesting bool
var err error
var placementGroupName string
var CertCRN string
var UpdatedCertCRN string
var RegionName string
var ISZoneName string
var ISCIDR string
var ISAddressPrefixCIDR string
var InstanceName string
var InstanceProfileName string
var InstanceProfileNameUpdate string
var IsBareMetalServerProfileName string
var IsBareMetalServerImage string
var DedicatedHostProfileName string
var DedicatedHostGroupID string
var InstanceDiskProfileName string
var DedicatedHostGroupFamily string
var DedicatedHostGroupClass string
var VolumeProfileName string
var ISRouteDestination string
var ISRouteNextHop string
var WorkspaceID string
var TemplateID string
var ActionID string
var JobID string
var RepoURL string
var RepoBranch string
var imageName string
var functionNamespace string
var HpcsInstanceID string
var SecretsManagerInstanceID string
var SecretsManagerSecretType string
var SecretsManagerSecretID string
var HpcsAdmin1 string
var HpcsToken1 string
var HpcsAdmin2 string
var HpcsToken2 string
var RealmName string
var IksSa string
var IksClusterVpcID string
var IksClusterSubnetID string
var IksClusterResourceGroupID string

// For Power Colo

var Pi_image string
var Pi_image_bucket_name string
var Pi_image_bucket_file_name string
var Pi_image_bucket_access_key string
var Pi_image_bucket_secret_key string
var Pi_image_bucket_region string
var Pi_key_name string
var Pi_volume_name string
var Pi_network_name string
var Pi_cloud_instance_id string
var Pi_instance_name string
var Pi_dhcp_id string
var PiCloudConnectionName string
var PiSAPProfileID string
var Pi_placement_group_name string

var Pi_capture_storage_image_path string
var Pi_capture_cloud_storage_access_key string
var Pi_capture_cloud_storage_secret_key string

// For Image

var IsImageName string
var IsImage string
var IsImageEncryptedDataKey string
var IsImageEncryptionKey string
var IsWinImage string
var Image_cos_url string
var Image_cos_url_encrypted string
var Image_operating_system string

// Transit Gateway cross account
var Tg_cross_network_account_id string
var Tg_cross_network_id string

//Enterprise Management
var Account_to_be_imported string

//Security and Compliance Center, SI
var Scc_si_account string

//Security and Compliance Center, Posture Management
var Scc_posture_scope_id string
var Scc_posture_scan_id string
var Scc_posture_profile_id string
var Scc_posture_group_profile_id string
var Scc_posture_correlation_id string
var Scc_posture_report_setting_id string
var Scc_posture_profile_id_scansummary string
var Scc_posture_scan_id_scansummary string
var Scc_posture_credential_id_scope string
var Scc_posture_credential_id_scope_update string
var Scc_posture_collector_id_scope []string
var Scc_posture_collector_id_scope_update []string

//ROKS Cluster
var ClusterName string

func init() {
	testlogger := os.Getenv("TF_LOG")
	if testlogger != "" {
		os.Setenv("IBMCLOUD_BLUEMIX_GO_TRACE", "true")
	}

	AppIDTenantID = os.Getenv("IBM_APPID_TENANT_ID")
	if AppIDTenantID == "" {
		fmt.Println("[WARN] Set the environment variable IBM_APPID_TENANT_ID for testing AppID resources, AppID tests will fail if this is not set")
	}

	AppIDTestUserEmail = os.Getenv("IBM_APPID_TEST_USER_EMAIL")
	if AppIDTestUserEmail == "" {
		fmt.Println("[WARN] Set the environment variable IBM_APPID_TEST_USER_EMAIL for testing AppID user resources, the tests will fail if this is not set")
	}

	CfOrganization = os.Getenv("IBM_ORG")
	if CfOrganization == "" {
		fmt.Println("[WARN] Set the environment variable IBM_ORG for testing ibm_org  resource Some tests for that resource will fail if this is not set correctly")
	}
	CfSpace = os.Getenv("IBM_SPACE")
	if CfSpace == "" {
		fmt.Println("[WARN] Set the environment variable IBM_SPACE for testing ibm_space  resource Some tests for that resource will fail if this is not set correctly")
	}
	Ibmid1 = os.Getenv("IBM_ID1")
	if Ibmid1 == "" {
		fmt.Println("[WARN] Set the environment variable IBM_ID1 for testing ibm_space resource Some tests for that resource will fail if this is not set correctly")
	}

	Ibmid2 = os.Getenv("IBM_ID2")
	if Ibmid2 == "" {
		fmt.Println("[WARN] Set the environment variable IBM_ID2 for testing ibm_space resource Some tests for that resource will fail if this is not set correctly")
	}

	IAMUser = os.Getenv("IBM_IAMUSER")
	if IAMUser == "" {
		fmt.Println("[WARN] Set the environment variable IBM_IAMUSER for testing ibm_iam_user_policy resource Some tests for that resource will fail if this is not set correctly")
	}

	Datacenter = os.Getenv("IBM_DATACENTER")
	if Datacenter == "" {
		Datacenter = "par01"
		fmt.Println("[WARN] Set the environment variable IBM_DATACENTER for testing ibm_container_cluster resource else it is set to default value 'par01'")
	}
	MachineType = os.Getenv("IBM_MACHINE_TYPE")
	if MachineType == "" {
		MachineType = "b3c.4x16"
		fmt.Println("[WARN] Set the environment variable IBM_MACHINE_TYPE for testing ibm_container_cluster resource else it is set to default value 'b3c.4x16'")
	}

	CertCRN = os.Getenv("IBM_CERT_CRN")
	if CertCRN == "" {
		CertCRN = "crn:v1:bluemix:public:cloudcerts:us-south:a/52b2e14f385aca5da781baa1b9c28e53:6efac0c2-b955-49ca-939d-d7bc0cb8132f:certificate:e786b0ea2af8b5435603803ec2ff8118"
		fmt.Println("[WARN] Set the environment variable IBM_CERT_CRN for testing ibm_container_alb_cert resource else it is set to default value")
	}

	UpdatedCertCRN = os.Getenv("IBM_UPDATE_CERT_CRN")
	if UpdatedCertCRN == "" {
		UpdatedCertCRN = "crn:v1:bluemix:public:cloudcerts:eu-de:a/e9021a4d06e9b108b4a221a3cec47e3d:77e527aa-65b2-4cb3-969b-7e8714174346:certificate:1bf3d0c2b7764402dde25744218e6cba"
		fmt.Println("[WARN] Set the environment variable IBM_UPDATE_CERT_CRN for testing ibm_container_alb_cert resource else it is set to default value")
	}

	CsRegion = os.Getenv("IBM_CONTAINER_REGION")
	if CsRegion == "" {
		CsRegion = "eu-de"
		fmt.Println("[WARN] Set the environment variable IBM_CONTAINER_REGION for testing ibm_container resources else it is set to default value 'eu-de'")
	}

	CisInstance = os.Getenv("IBM_CIS_INSTANCE")
	if CisInstance == "" {
		CisInstance = ""
		fmt.Println("[WARN] Set the environment variable IBM_CIS_INSTANCE with a VALID CIS Instance NAME for testing ibm_cis resources on staging/test")
	}
	CisDomainStatic = os.Getenv("IBM_CIS_DOMAIN_STATIC")
	if CisDomainStatic == "" {
		CisDomainStatic = ""
		fmt.Println("[WARN] Set the environment variable IBM_CIS_DOMAIN_STATIC with the Domain name registered with the CIS instance on test/staging. Domain must be predefined in CIS to avoid CIS billing costs due to domain delete/create")
	}

	CisDomainTest = os.Getenv("IBM_CIS_DOMAIN_TEST")
	if CisDomainTest == "" {
		CisDomainTest = ""
		fmt.Println("[WARN] Set the environment variable IBM_CIS_DOMAIN_TEST with a VALID Domain name for testing the one time create and delete of a domain in CIS. Note each create/delete will trigger a monthly billing instance. Only to be run in staging/test")
	}

	CisResourceGroup = os.Getenv("IBM_CIS_RESOURCE_GROUP")
	if CisResourceGroup == "" {
		CisResourceGroup = ""
		fmt.Println("[WARN] Set the environment variable IBM_CIS_RESOURCE_GROUP with the resource group for the CIS Instance ")
	}

	CosCRN = os.Getenv("IBM_COS_CRN")
	if CosCRN == "" {
		CosCRN = ""
		fmt.Println("[WARN] Set the environment variable IBM_COS_CRN with a VALID COS instance CRN for testing ibm_cos_* resources")
	}

	trustedMachineType = os.Getenv("IBM_TRUSTED_MACHINE_TYPE")
	if trustedMachineType == "" {
		trustedMachineType = "mb1c.16x64"
		fmt.Println("[WARN] Set the environment variable IBM_TRUSTED_MACHINE_TYPE for testing ibm_container_cluster resource else it is set to default value 'mb1c.16x64'")
	}

	ExtendedHardwareTesting, err = strconv.ParseBool(os.Getenv("IBM_BM_EXTENDED_HW_TESTING"))
	if err != nil {
		ExtendedHardwareTesting = false
		fmt.Println("[WARN] Set the environment variable IBM_BM_EXTENDED_HW_TESTING to true/false for testing ibm_compute_bare_metal resource else it is set to default value 'false'")
	}

	PublicVlanID = os.Getenv("IBM_PUBLIC_VLAN_ID")
	if PublicVlanID == "" {
		PublicVlanID = "2393319"
		fmt.Println("[WARN] Set the environment variable IBM_PUBLIC_VLAN_ID for testing ibm_container_cluster resource else it is set to default value '2393319'")
	}

	PrivateVlanID = os.Getenv("IBM_PRIVATE_VLAN_ID")
	if PrivateVlanID == "" {
		PrivateVlanID = "2393321"
		fmt.Println("[WARN] Set the environment variable IBM_PRIVATE_VLAN_ID for testing ibm_container_cluster resource else it is set to default value '2393321'")
	}

	KubeVersion = os.Getenv("IBM_KUBE_VERSION")
	if KubeVersion == "" {
		KubeVersion = "1.18"
		fmt.Println("[WARN] Set the environment variable IBM_KUBE_VERSION for testing ibm_container_cluster resource else it is set to default value '1.18.14'")
	}

	KubeUpdateVersion = os.Getenv("IBM_KUBE_UPDATE_VERSION")
	if KubeUpdateVersion == "" {
		KubeUpdateVersion = "1.19"
		fmt.Println("[WARN] Set the environment variable IBM_KUBE_UPDATE_VERSION for testing ibm_container_cluster resource else it is set to default value '1.19.6'")
	}

	PrivateSubnetID = os.Getenv("IBM_PRIVATE_SUBNET_ID")
	if PrivateSubnetID == "" {
		PrivateSubnetID = "1636107"
		fmt.Println("[WARN] Set the environment variable IBM_PRIVATE_SUBNET_ID for testing ibm_container_cluster resource else it is set to default value '1636107'")
	}

	PublicSubnetID = os.Getenv("IBM_PUBLIC_SUBNET_ID")
	if PublicSubnetID == "" {
		PublicSubnetID = "1165645"
		fmt.Println("[WARN] Set the environment variable IBM_PUBLIC_SUBNET_ID for testing ibm_container_cluster resource else it is set to default value '1165645'")
	}

	SubnetID = os.Getenv("IBM_SUBNET_ID")
	if SubnetID == "" {
		SubnetID = "1165645"
		fmt.Println("[WARN] Set the environment variable IBM_SUBNET_ID for testing ibm_container_cluster resource else it is set to default value '1165645'")
	}

	IpsecDatacenter = os.Getenv("IBM_IPSEC_DATACENTER")
	if IpsecDatacenter == "" {
		IpsecDatacenter = "tok02"
		fmt.Println("[INFO] Set the environment variable IBM_IPSEC_DATACENTER for testing ibm_ipsec_vpn resource else it is set to default value 'tok02'")
	}

	Customersubnetid = os.Getenv("IBM_IPSEC_CUSTOMER_SUBNET_ID")
	if Customersubnetid == "" {
		Customersubnetid = "123456"
		fmt.Println("[INFO] Set the environment variable IBM_IPSEC_CUSTOMER_SUBNET_ID for testing ibm_ipsec_vpn resource else it is set to default value '123456'")
	}

	Customerpeerip = os.Getenv("IBM_IPSEC_CUSTOMER_PEER_IP")
	if Customerpeerip == "" {
		Customerpeerip = "192.168.0.1"
		fmt.Println("[INFO] Set the environment variable IBM_IPSEC_CUSTOMER_PEER_IP for testing ibm_ipsec_vpn resource else it is set to default value '192.168.0.1'")
	}

	LbaasDatacenter = os.Getenv("IBM_LBAAS_DATACENTER")
	if LbaasDatacenter == "" {
		LbaasDatacenter = "dal13"
		fmt.Println("[WARN] Set the environment variable IBM_LBAAS_DATACENTER for testing ibm_lbaas resource else it is set to default value 'dal13'")
	}

	LbaasSubnetId = os.Getenv("IBM_LBAAS_SUBNETID")
	if LbaasSubnetId == "" {
		LbaasSubnetId = "2144241"
		fmt.Println("[WARN] Set the environment variable IBM_LBAAS_SUBNETID for testing ibm_lbaas resource else it is set to default value '2144241'")
	}
	LbListerenerCertificateInstance = os.Getenv("IBM_LB_LISTENER_CERTIFICATE_INSTANCE")
	if LbListerenerCertificateInstance == "" {
		LbListerenerCertificateInstance = "crn:v1:staging:public:cloudcerts:us-south:a/2d1bace7b46e4815a81e52c6ffeba5cf:af925157-b125-4db2-b642-adacb8b9c7f5:certificate:c81627a1bf6f766379cc4b98fd2a44ed"
		fmt.Println("[WARN] Set the environment variable IBM_LB_LISTENER_CERTIFICATE_INSTANCE for testing ibm_is_lb_listener resource for https redirect else it is set to default value 'crn:v1:staging:public:cloudcerts:us-south:a/2d1bace7b46e4815a81e52c6ffeba5cf:af925157-b125-4db2-b642-adacb8b9c7f5:certificate:c81627a1bf6f766379cc4b98fd2a44ed'")
	}

	DedicatedHostName = os.Getenv("IBM_DEDICATED_HOSTNAME")
	if DedicatedHostName == "" {
		DedicatedHostName = "terraform-dedicatedhost"
		fmt.Println("[WARN] Set the environment variable IBM_DEDICATED_HOSTNAME for testing ibm_compute_vm_instance resource else it is set to default value 'terraform-dedicatedhost'")
	}

	DedicatedHostID = os.Getenv("IBM_DEDICATED_HOST_ID")
	if DedicatedHostID == "" {
		DedicatedHostID = "30301"
		fmt.Println("[WARN] Set the environment variable IBM_DEDICATED_HOST_ID for testing ibm_compute_vm_instance resource else it is set to default value '30301'")
	}

	Zone = os.Getenv("IBM_WORKER_POOL_ZONE")
	if Zone == "" {
		Zone = "ams03"
		fmt.Println("[WARN] Set the environment variable IBM_WORKER_POOL_ZONE for testing ibm_container_worker_pool_zone_attachment resource else it is set to default value 'ams03'")
	}

	ZonePrivateVlan = os.Getenv("IBM_WORKER_POOL_ZONE_PRIVATE_VLAN")
	if ZonePrivateVlan == "" {
		ZonePrivateVlan = "2538975"
		fmt.Println("[WARN] Set the environment variable IBM_WORKER_POOL_ZONE_PRIVATE_VLAN for testing ibm_container_worker_pool_zone_attachment resource else it is set to default value '2538975'")
	}

	ZonePublicVlan = os.Getenv("IBM_WORKER_POOL_ZONE_PUBLIC_VLAN")
	if ZonePublicVlan == "" {
		ZonePublicVlan = "2538967"
		fmt.Println("[WARN] Set the environment variable IBM_WORKER_POOL_ZONE_PUBLIC_VLAN for testing ibm_container_worker_pool_zone_attachment resource else it is set to default value '2538967'")
	}

	ZoneUpdatePrivateVlan = os.Getenv("IBM_WORKER_POOL_ZONE_UPDATE_PRIVATE_VLAN")
	if ZoneUpdatePrivateVlan == "" {
		ZoneUpdatePrivateVlan = "2388377"
		fmt.Println("[WARN] Set the environment variable IBM_WORKER_POOL_ZONE_UPDATE_PRIVATE_VLAN for testing ibm_container_worker_pool_zone_attachment resource else it is set to default value '2388377'")
	}

	ZoneUpdatePublicVlan = os.Getenv("IBM_WORKER_POOL_ZONE_UPDATE_PUBLIC_VLAN")
	if ZoneUpdatePublicVlan == "" {
		ZoneUpdatePublicVlan = "2388375"
		fmt.Println("[WARN] Set the environment variable IBM_WORKER_POOL_ZONE_UPDATE_PUBLIC_VLAN for testing ibm_container_worker_pool_zone_attachment resource else it is set to default value '2388375'")
	}

	placementGroupName = os.Getenv("IBM_PLACEMENT_GROUP_NAME")
	if placementGroupName == "" {
		placementGroupName = "terraform_group"
		fmt.Println("[WARN] Set the environment variable IBM_PLACEMENT_GROUP_NAME for testing ibm_compute_vm_instance resource else it is set to default value 'terraform-group'")
	}

	RegionName = os.Getenv("SL_REGION")
	if RegionName == "" {
		RegionName = "us-south"
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

	IsImage = os.Getenv("IS_IMAGE")
	if IsImage == "" {
		//IsImage = "fc538f61-7dd6-4408-978c-c6b85b69fe76" // for classic infrastructure
		IsImage = "r006-13938c0a-89e4-4370-b59b-55cd1402562d" // for next gen infrastructure
		fmt.Println("[INFO] Set the environment variable IS_IMAGE for testing ibm_is_instance, ibm_is_floating_ip else it is set to default value 'r006-ed3f775f-ad7e-4e37-ae62-7199b4988b00'")
	}

	IsWinImage = os.Getenv("IS_WIN_IMAGE")
	if IsWinImage == "" {
		//IsWinImage = "a7a0626c-f97e-4180-afbe-0331ec62f32a" // classic windows machine: ibm-windows-server-2012-full-standard-amd64-1
		IsWinImage = "r006-5f9568ae-792e-47e1-a710-5538b2bdfca7" // next gen windows machine: ibm-windows-server-2012-full-standard-amd64-3
		fmt.Println("[INFO] Set the environment variable IS_WIN_IMAGE for testing ibm_is_instance data source else it is set to default value 'r006-5f9568ae-792e-47e1-a710-5538b2bdfca7'")
	}

	InstanceName = os.Getenv("IS_INSTANCE_NAME")
	if InstanceName == "" {
		InstanceName = "placement-check-ins" // for next gen infrastructure
		fmt.Println("[INFO] Set the environment variable IS_INSTANCE_NAME for testing ibm_is_instance resource else it is set to default value 'instance-01'")
	}

	InstanceProfileName = os.Getenv("SL_INSTANCE_PROFILE")
	if InstanceProfileName == "" {
		//InstanceProfileName = "bc1-2x8" // for classic infrastructure
		InstanceProfileName = "cx2-2x4" // for next gen infrastructure
		fmt.Println("[INFO] Set the environment variable SL_INSTANCE_PROFILE for testing ibm_is_instance resource else it is set to default value 'cx2-2x4'")
	}

	InstanceProfileNameUpdate = os.Getenv("SL_INSTANCE_PROFILE_UPDATE")
	if InstanceProfileNameUpdate == "" {
		InstanceProfileNameUpdate = "cx2-4x8"
		fmt.Println("[INFO] Set the environment variable SL_INSTANCE_PROFILE_UPDATE for testing ibm_is_instance resource else it is set to default value 'cx2-4x8'")
	}

	IsBareMetalServerProfileName = os.Getenv("IS_BARE_METAL_SERVER_PROFILE")
	if IsBareMetalServerProfileName == "" {
		IsBareMetalServerProfileName = "bx2-metal-192x768" // for next gen infrastructure
		fmt.Println("[INFO] Set the environment variable IS_BARE_METAL_SERVER_PROFILE for testing ibm_is_bare_metal_server resource else it is set to default value 'bx2-metal-192x768'")
	}

	IsBareMetalServerImage = os.Getenv("IS_BARE_METAL_SERVER_IMAGE")
	if IsBareMetalServerImage == "" {
		IsBareMetalServerImage = "r006-2d1f36b0-df65-4570-82eb-df7ae5f778b1" // for next gen infrastructure
		fmt.Println("[INFO] Set the environment variable IsBareMetalServerImage for testing ibm_is_bare_metal_server resource else it is set to default value 'r006-2d1f36b0-df65-4570-82eb-df7ae5f778b1'")
	}

	DedicatedHostName = os.Getenv("IS_DEDICATED_HOST_NAME")
	if DedicatedHostName == "" {
		DedicatedHostName = "tf-dhost-01" // for next gen infrastructure
		fmt.Println("[INFO] Set the environment variable IS_DEDICATED_HOST_NAME for testing ibm_is_instance resource else it is set to default value 'tf-dhost-01'")
	}

	DedicatedHostGroupID = os.Getenv("IS_DEDICATED_HOST_GROUP_ID")
	if DedicatedHostGroupID == "" {
		DedicatedHostGroupID = "0717-9104e7b5-77ad-44ad-9eaa-091e6b6efce1" // for next gen infrastructure
		fmt.Println("[INFO] Set the environment variable IS_DEDICATED_HOST_GROUP_ID for testing ibm_is_instance resource else it is set to default value '0717-9104e7b5-77ad-44ad-9eaa-091e6b6efce1'")
	}

	DedicatedHostProfileName = os.Getenv("IS_DEDICATED_HOST_PROFILE")
	if DedicatedHostProfileName == "" {
		DedicatedHostProfileName = "bx2d-host-152x608" // for next gen infrastructure
		fmt.Println("[INFO] Set the environment variable IS_DEDICATED_HOST_PROFILE for testing ibm_is_instance resource else it is set to default value 'bx2d-host-152x608'")
	}

	DedicatedHostGroupClass = os.Getenv("IS_DEDICATED_HOST_GROUP_CLASS")
	if DedicatedHostGroupClass == "" {
		DedicatedHostGroupClass = "bx2d" // for next gen infrastructure
		fmt.Println("[INFO] Set the environment variable IS_DEDICATED_HOST_GROUP_CLASS for testing ibm_is_instance resource else it is set to default value 'bx2d'")
	}

	DedicatedHostGroupFamily = os.Getenv("IS_DEDICATED_HOST_GROUP_FAMILY")
	if DedicatedHostGroupFamily == "" {
		DedicatedHostGroupFamily = "balanced" // for next gen infrastructure
		fmt.Println("[INFO] Set the environment variable IS_DEDICATED_HOST_GROUP_FAMILY for testing ibm_is_instance resource else it is set to default value 'balanced'")
	}

	InstanceDiskProfileName = os.Getenv("IS_INSTANCE_DISK_PROFILE")
	if InstanceDiskProfileName == "" {
		//InstanceProfileName = "bc1-2x8" // for classic infrastructure
		InstanceDiskProfileName = "bx2d-16x64" // for next gen infrastructure
		fmt.Println("[INFO] Set the environment variable SL_INSTANCE_PROFILE for testing ibm_is_instance resource else it is set to default value 'bx2d-16x64'")
	}

	VolumeProfileName = os.Getenv("IS_VOLUME_PROFILE")
	if VolumeProfileName == "" {
		VolumeProfileName = "general-purpose"
		fmt.Println("[INFO] Set the environment variable IS_VOLUME_PROFILE for testing ibm_is_volume_profile else it is set to default value 'general-purpose'")
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
	Pi_image = os.Getenv("PI_IMAGE")
	if Pi_image == "" {
		Pi_image = "c93dc4c6-e85a-4da2-9ea6-f24576256122"
		fmt.Println("[INFO] Set the environment variable PI_IMAGE for testing ibm_pi_image resource else it is set to default value '7200-03-03'")
	}
	Pi_image_bucket_name = os.Getenv("PI_IMAGE_BUCKET_NAME")
	if Pi_image_bucket_name == "" {
		Pi_image_bucket_name = "images-public-bucket"
		fmt.Println("[INFO] Set the environment variable PI_IMAGE_BUCKET_NAME for testing ibm_pi_image resource else it is set to default value 'images-public-bucket'")
	}
	Pi_image_bucket_file_name = os.Getenv("PI_IMAGE_BUCKET_FILE_NAME")
	if Pi_image_bucket_file_name == "" {
		Pi_image_bucket_file_name = "rhel.ova.gz"
		fmt.Println("[INFO] Set the environment variable PI_IMAGE_BUCKET_FILE_NAME for testing ibm_pi_image resource else it is set to default value 'rhel.ova.gz'")
	}
	Pi_image_bucket_access_key = os.Getenv("PI_IMAGE_BUCKET_ACCESS_KEY")
	if Pi_image_bucket_access_key == "" {
		Pi_image_bucket_access_key = "images-bucket-access-key"
		fmt.Println("[INFO] Set the environment variable PI_IMAGE_BUCKET_ACCESS_KEY for testing ibm_pi_image_export resource else it is set to default value 'images-bucket-access-key'")
	}

	Pi_image_bucket_secret_key = os.Getenv("PI_IMAGE_BUCKET_SECRET_KEY")
	if Pi_image_bucket_secret_key == "" {
		Pi_image_bucket_secret_key = "images-bucket-secret-key"
		fmt.Println("[INFO] Set the environment variable PI_IMAGE_BUCKET_SECRET_KEY for testing ibm_pi_image_export resource else it is set to default value 'PI_IMAGE_BUCKET_SECRET_KEY'")
	}

	Pi_image_bucket_region = os.Getenv("PI_IMAGE_BUCKET_REGION")
	if Pi_image_bucket_region == "" {
		Pi_image_bucket_region = "us-east"
		fmt.Println("[INFO] Set the environment variable PI_IMAGE_BUCKET_REGION for testing ibm_pi_image_export resource else it is set to default value 'us-east'")
	}

	Pi_key_name = os.Getenv("PI_KEY_NAME")
	if Pi_key_name == "" {
		Pi_key_name = "terraform-test-power"
		fmt.Println("[INFO] Set the environment variable PI_KEY_NAME for testing ibm_pi_key_name resource else it is set to default value 'terraform-test-power'")
	}

	Pi_network_name = os.Getenv("PI_NETWORK_NAME")
	if Pi_network_name == "" {
		Pi_network_name = "terraform-test-power"
		fmt.Println("[INFO] Set the environment variable PI_NETWORK_NAME for testing ibm_pi_network_name resource else it is set to default value 'terraform-test-power'")
	}

	Pi_volume_name = os.Getenv("PI_VOLUME_NAME")
	if Pi_volume_name == "" {
		Pi_volume_name = "terraform-test-power"
		fmt.Println("[INFO] Set the environment variable PI_VOLUME_NAME for testing ibm_pi_network_name resource else it is set to default value 'terraform-test-power'")
	}

	Pi_cloud_instance_id = os.Getenv("PI_CLOUDINSTANCE_ID")
	if Pi_cloud_instance_id == "" {
		Pi_cloud_instance_id = "fd3454a3-14d8-4eb0-b075-acf3da5cd324"
		fmt.Println("[INFO] Set the environment variable PI_CLOUDINSTANCE_ID for testing ibm_pi_image resource else it is set to default value 'd16705bd-7f1a-48c9-9e0e-1c17b71e7331'")
	}

	Pi_instance_name = os.Getenv("PI_PVM_INSTANCE_NAME")
	if Pi_instance_name == "" {
		Pi_instance_name = "terraform-test-power"
		fmt.Println("[INFO] Set the environment variable PI_PVM_INSTANCE_ID for testing Pi_instance_name resource else it is set to default value 'terraform-test-power'")
	}

	Pi_dhcp_id = os.Getenv("PI_DHCP_ID")
	if Pi_dhcp_id == "" {
		Pi_dhcp_id = "terraform-test-power"
		fmt.Println("[INFO] Set the environment variable PI_DHCP_ID for testing ibm_pi_dhcp resource else it is set to default value 'terraform-test-power'")
	}

	PiCloudConnectionName = os.Getenv("PI_CLOUD_CONNECTION_NAME")
	if PiCloudConnectionName == "" {
		PiCloudConnectionName = "terraform-test-power"
		fmt.Println("[INFO] Set the environment variable PI_CLOUD_CONNECTION_NAME for testing ibm_pi_cloud_connection resource else it is set to default value 'terraform-test-power'")
	}

	PiSAPProfileID = os.Getenv("PI_SAP_PROFILE_ID")
	if PiSAPProfileID == "" {
		PiSAPProfileID = "terraform-test-power"
		fmt.Println("[INFO] Set the environment variable PI_SAP_PROFILE_ID for testing ibm_pi_sap_profile resource else it is set to default value 'terraform-test-power'")
	}

	Pi_placement_group_name = os.Getenv("PI_PLACEMENT_GROUP_NAME")
	if Pi_placement_group_name == "" {
		Pi_placement_group_name = "tf-pi-placement-group"
		fmt.Println("[WARN] Set the environment variable PI_PLACEMENT_GROUP_NAME for testing ibm_pi_placement_group resource else it is set to default value 'tf-pi-placement-group'")
	}
	// Added for resource capture instance testing
	Pi_capture_storage_image_path = os.Getenv("PI_CAPTURE_STORAGE_IMAGE_PATH")
	if Pi_capture_storage_image_path == "" {
		Pi_capture_storage_image_path = "bucket-test"
		fmt.Println("[INFO] Set the environment variable PI_CAPTURE_STORAGE_IMAGE_PATH for testing Pi_capture_storage_image_path resource else it is set to default value 'terraform-test-power'")
	}

	Pi_capture_cloud_storage_access_key = os.Getenv("PI_CAPTURE_CLOUD_STORAGE_ACCESS_KEY")
	if Pi_capture_cloud_storage_access_key == "" {
		Pi_capture_cloud_storage_access_key = "terraform-test-power"
		fmt.Println("[INFO] Set the environment variable PI_CAPTURE_CLOUD_STORAGE_ACCESS_KEY for testing Pi_capture_cloud_storage_access_key resource else it is set to default value 'terraform-test-power'")
	}

	Pi_capture_cloud_storage_secret_key = os.Getenv("PI_CAPTURE_CLOUD_STORAGE_SECRET_KEY")
	if Pi_capture_cloud_storage_secret_key == "" {
		Pi_capture_cloud_storage_secret_key = "terraform-test-power"
		fmt.Println("[INFO] Set the environment variable PI_CAPTURE_CLOUD_STORAGE_SECRET_KEY for testing Pi_capture_cloud_storage_secret_key resource else it is set to default value 'terraform-test-power'")
	}

	WorkspaceID = os.Getenv("SCHEMATICS_WORKSPACE_ID")
	if WorkspaceID == "" {
		WorkspaceID = "us-south.workspace.tf-acc-test-schematics-state-test.392cd99f"
		fmt.Println("[INFO] Set the environment variable SCHEMATICS_WORKSPACE_ID for testing schematics resources else it is set to default value")
	}
	TemplateID = os.Getenv("SCHEMATICS_TEMPLATE_ID")
	if TemplateID == "" {
		TemplateID = "c8d52331-056f-40"
		fmt.Println("[INFO] Set the environment variable SCHEMATICS_TEMPLATE_ID for testing schematics resources else it is set to default value")
	}
	ActionID = os.Getenv("SCHEMATICS_ACTION_ID")
	if ActionID == "" {
		ActionID = "us-east.ACTION.action_pm.a4ffeec3"
		fmt.Println("[INFO] Set the environment variable SCHEMATICS_ACTION_ID for testing schematics resources else it is set to default value")
	}
	JobID = os.Getenv("SCHEMATICS_JOB_ID")
	if JobID == "" {
		JobID = "us-east.ACTION.action_pm.a4ffeec3"
		fmt.Println("[INFO] Set the environment variable SCHEMATICS_JOB_ID for testing schematics resources else it is set to default value")
	}
	RepoURL = os.Getenv("SCHEMATICS_REPO_URL")
	if RepoURL == "" {
		fmt.Println("[INFO] Set the environment variable SCHEMATICS_REPO_URL for testing schematics resources else tests will fail if this is not set correctly")
	}
	RepoBranch = os.Getenv("SCHEMATICS_REPO_BRANCH")
	if RepoBranch == "" {
		fmt.Println("[INFO] Set the environment variable SCHEMATICS_REPO_BRANCH for testing schematics resources else tests will fail if this is not set correctly")
	}
	// Added for resource image testing
	Image_cos_url = os.Getenv("IMAGE_COS_URL")
	if Image_cos_url == "" {
		Image_cos_url = "cos://us-south/cosbucket-vpc-image-gen2/rhel-guest-image-7.0-20140930.0.x86_64.qcow2"
		fmt.Println("[WARN] Set the environment variable IMAGE_COS_URL with a VALID COS Image SQL URL for testing ibm_is_image resources on staging/test")
	}

	// Added for resource image testing
	Image_cos_url_encrypted = os.Getenv("IMAGE_COS_URL_ENCRYPTED")
	if Image_cos_url_encrypted == "" {
		Image_cos_url_encrypted = "cos://us-south/cosbucket-vpc-image-gen2/rhel-guest-image-7.0-encrypted.qcow2"
		fmt.Println("[WARN] Set the environment variable IMAGE_COS_URL_ENCRYPTED with a VALID COS Image SQL URL for testing ibm_is_image resources on staging/test")
	}
	Image_operating_system = os.Getenv("IMAGE_OPERATING_SYSTEM")
	if Image_operating_system == "" {
		Image_operating_system = "red-7-amd64"
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

	HpcsInstanceID = os.Getenv("HPCS_INSTANCE_ID")
	if HpcsInstanceID == "" {
		HpcsInstanceID = "5af62d5d-5d90-4b84-bbcd-90d2123ae6c8"
		fmt.Println("[INFO] Set the environment variable HPCS_INSTANCE_ID for testing data_source_ibm_kms_key_test else it is set to default value")
	}

	SecretsManagerInstanceID = os.Getenv("SECRETS_MANAGER_INSTANCE_ID")
	if SecretsManagerInstanceID == "" {
		// SecretsManagerInstanceID = "5af62d5d-5d90-4b84-bbcd-90d2123ae6c8"
		fmt.Println("[INFO] Set the environment variable SECRETS_MANAGER_INSTANCE_ID for testing data_source_ibm_secrets_manager_secrets_test else tests will fail if this is not set correctly")
	}

	SecretsManagerSecretType = os.Getenv("SECRETS_MANAGER_SECRET_TYPE")
	if SecretsManagerSecretType == "" {
		SecretsManagerSecretType = "username_password"
		fmt.Println("[INFO] Set the environment variable SECRETS_MANAGER_SECRET_TYPE for testing data_source_ibm_secrets_manager_secrets_test, else it is set to default value. For data_source_ibm_secrets_manager_secret_test, tests will fail if this is not set correctly")
	}

	SecretsManagerSecretID = os.Getenv("SECRETS_MANAGER_SECRET_ID")
	if SecretsManagerSecretID == "" {
		// SecretsManagerSecretID = "644f4a69-0d17-198f-3b58-23f2746c706d"
		fmt.Println("[WARN] Set the environment variable SECRETS_MANAGER_SECRET_ID for testing data_source_ibm_secrets_manager_secret_test else tests will fail if this is not set correctly")
	}

	Tg_cross_network_account_id = os.Getenv("IBM_TG_CROSS_ACCOUNT_ID")
	if Tg_cross_network_account_id == "" {
		fmt.Println("[INFO] Set the environment variable IBM_TG_CROSS_ACCOUNT_ID for testing ibm_tg_connection resource else  tests will fail if this is not set correctly")
	}
	Tg_cross_network_id = os.Getenv("IBM_TG_CROSS_NETWORK_ID")
	if Tg_cross_network_id == "" {
		fmt.Println("[INFO] Set the environment variable IBM_TG_CROSS_NETWORK_ID for testing ibm_tg_connection resource else  tests will fail if this is not set correctly")
	}
	Account_to_be_imported = os.Getenv("ACCOUNT_TO_BE_IMPORTED")
	if Account_to_be_imported == "" {
		fmt.Println("[INFO] Set the environment variable ACCOUNT_TO_BE_IMPORTED for testing import enterprise account resource else  tests will fail if this is not set correctly")
	}
	HpcsAdmin1 = os.Getenv("IBM_HPCS_ADMIN1")
	if HpcsAdmin1 == "" {
		fmt.Println("[WARN] Set the environment variable IBM_HPCS_ADMIN1 with a VALID HPCS Admin Key1 Path")
	}
	HpcsToken1 = os.Getenv("IBM_HPCS_TOKEN1")
	if HpcsToken1 == "" {
		fmt.Println("[WARN] Set the environment variable IBM_HPCS_TOKEN1 with a VALID token for HPCS Admin Key1")
	}
	HpcsAdmin2 = os.Getenv("IBM_HPCS_ADMIN2")
	if HpcsAdmin2 == "" {
		fmt.Println("[WARN] Set the environment variable IBM_HPCS_ADMIN2 with a VALID HPCS Admin Key2 Path")
	}
	RealmName = os.Getenv("IBM_IAM_REALM_NAME")
	if RealmName == "" {
		fmt.Println("[WARN] Set the environment variable IBM_IAM_REALM_NAME with a VALID realm name for iam trusted profile claim rule")
	}

	IksSa = os.Getenv("IBM_IAM_IKS_SA")
	if IksSa == "" {
		fmt.Println("[WARN] Set the environment variable IBM_IAM_IKS_SA with a VALID realm name for iam trusted profile link")
	}

	HpcsToken2 = os.Getenv("IBM_HPCS_TOKEN2")
	if HpcsToken2 == "" {
		fmt.Println("[WARN] Set the environment variable IBM_HPCS_TOKEN2 with a VALID token for HPCS Admin Key2")
	}
	Scc_si_account = os.Getenv("SCC_SI_ACCOUNT")
	if Scc_si_account == "" {
		fmt.Println("[INFO] Set the environment variable SCC_SI_ACCOUNT for testing SCC SI resources resource else  tests will fail if this is not set correctly")
	}

	Scc_posture_scope_id = os.Getenv("SCC_POSTURE_SCOPE_ID")
	if Scc_posture_scope_id == "" {
		fmt.Println("[INFO] Set the environment variable SCC_POSTURE_SCOPE_ID for testing SCC Posture resources or datasource resource else  tests will fail if this is not set correctly")
	}

	Scc_posture_scan_id = os.Getenv("SCC_POSTURE_SCAN_ID")
	if Scc_posture_scan_id == "" {
		fmt.Println("[INFO] Set the environment variable SCC_POSTURE_SCAN_ID for testing SCC Posture resource or datasource else  tests will fail if this is not set correctly")
	}

	Scc_posture_profile_id = os.Getenv("SCC_POSTURE_PROFILE_ID")
	if Scc_posture_profile_id == "" {
		fmt.Println("[INFO] Set the environment variable SCC_POSTURE_PROFILE_ID for testing SCC Posture resource or datasource else  tests will fail if this is not set correctly")
	}
	Scc_posture_group_profile_id = os.Getenv("SCC_POSTURE_GROUP_PROFILE_ID")
	if Scc_posture_group_profile_id == "" {
		fmt.Println("[INFO] Set the environment variable SCC_POSTURE_GROUP_PROFILE_ID for testing SCC Posture resource or datasource else  tests will fail if this is not set correctly")
	}

	Scc_posture_correlation_id = os.Getenv("SCC_POSTURE_CORRELATION_ID")
	if Scc_posture_correlation_id == "" {
		fmt.Println("[INFO] Set the environment variable SCC_POSTURE_CORRELATION_ID for testing SCC Posture resource or datasource else  tests will fail if this is not set correctly")
	}

	Scc_posture_report_setting_id = os.Getenv("SCC_POSTURE_REPORT_SETTING_ID")
	if Scc_posture_report_setting_id == "" {
		fmt.Println("[INFO] Set the environment variable SCC_POSTURE_REPORT_SETTING_ID for testing SCC Posture resource or datasource else  tests will fail if this is not set correctly")
	}

	Scc_posture_profile_id_scansummary = os.Getenv("SCC_POSTURE_PROFILE_ID_SCANSUMMARY")
	if Scc_posture_profile_id_scansummary == "" {
		fmt.Println("[INFO] Set the environment variable SCC_POSTURE_PROFILE_ID_SCANSUMMARY for testing SCC Posture resource or datasource else  tests will fail if this is not set correctly")
	}

	Scc_posture_scan_id_scansummary = os.Getenv("SCC_POSTURE_SCAN_ID_SCANSUMMARY")
	if Scc_posture_scan_id_scansummary == "" {
		fmt.Println("[INFO] Set the environment variable SCC_POSTURE_SCAN_ID_SCANSUMMARY for testing SCC Posture resource or datasource else  tests will fail if this is not set correctly")
	}

	Scc_posture_credential_id_scope = os.Getenv("SCC_POSTURE_CREDENTIAL_ID_SCOPE")
	if Scc_posture_credential_id_scope == "" {
		fmt.Println("[INFO] Set the environment variable SCC_POSTURE_CREDENTIAL_ID_SCOPE for testing SCC Posture resource or datasource else  tests will fail if this is not set correctly")
	}

	Scc_posture_credential_id_scope_update = os.Getenv("SCC_POSTURE_CREDENTIAL_ID_SCOPE_UPDATE")
	if Scc_posture_credential_id_scope_update == "" {
		fmt.Println("[INFO] Set the environment variable SCC_POSTURE_CREDENTIAL_ID_SCOPE_UPDATE for testing SCC Posture resource or datasource else  tests will fail if this is not set correctly")
	}

	Scc_posture_collector_id_scope = []string{os.Getenv("SCC_POSTURE_COLLECTOR_ID_SCOPE")}
	if os.Getenv("SCC_POSTURE_COLLECTOR_ID_SCOPE") == "" {
		fmt.Println("[INFO] Set the environment variable SCC_POSTURE_COLLECTOR_ID_SCOPE for testing SCC Posture resource or datasource else  tests will fail if this is not set correctly")
	}

	Scc_posture_collector_id_scope_update = []string{os.Getenv("SCC_POSTURE_COLLECTOR_ID_SCOPE_UPDATE")}
	if os.Getenv("SCC_POSTURE_COLLECTOR_ID_SCOPE_UPDATE") == "" {
		fmt.Println("[INFO] Set the environment variable SCC_POSTURE_COLLECTOR_ID_SCOPE_UPDATE for testing SCC Posture resource or datasource else  tests will fail if this is not set correctly")
	}

	CloudShellAccountID = os.Getenv("IBM_CLOUD_SHELL_ACCOUNT_ID")
	if CloudShellAccountID == "" {
		fmt.Println("[INFO] Set the environment variable IBM_CLOUD_SHELL_ACCOUNT_ID for ibm-cloud-shell resource or datasource else tests will fail if this is not set correctly")
	}

	IksClusterVpcID = os.Getenv("IBM_CLUSTER_VPC_ID")
	if IksClusterVpcID == "" {
		fmt.Println("[WARN] Set the environment variable IBM_CLUSTER_VPC_ID for testing ibm_container_vpc_alb_create resources, ibm_container_vpc_alb_create tests will fail if this is not set")
	}

	IksClusterSubnetID = os.Getenv("IBM_CLUSTER_VPC_SUBNET_ID")
	if IksClusterSubnetID == "" {
		fmt.Println("[WARN] Set the environment variable IBM_CLUSTER_VPC_SUBNET_ID for testing ibm_container_vpc_alb_create resources, ibm_container_vpc_alb_creates tests will fail if this is not set")
	}

	IksClusterResourceGroupID = os.Getenv("IBM_CLUSTER_VPC_RESOURCE_GROUP_ID")
	if IksClusterResourceGroupID == "" {
		fmt.Println("[WARN] Set the environment variable IBM_CLUSTER_VPC_RESOURCE_GROUP_ID for testing ibm_container_vpc_alb_create resources, ibm_container_vpc_alb_creates tests will fail if this is not set")
	}

	ClusterName = os.Getenv("IBM_CONTAINER_CLUSTER_NAME")
	if ClusterName == "" {
		fmt.Println("[INFO] Set the environment variable IBM_CONTAINER_CLUSTER_NAME for ibm_container_nlb_dns resource or datasource else tests will fail if this is not set correctly")
	}
}

var TestAccProviders map[string]*schema.Provider
var TestAccProvider *schema.Provider

func init() {
	TestAccProvider = provider.Provider()
	TestAccProviders = map[string]*schema.Provider{
		"ibm": TestAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := provider.Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = provider.Provider()
}

func TestAccPreCheck(t *testing.T) {
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

func TestAccPreCheckEnterprise(t *testing.T) {
	if v := os.Getenv("IC_API_KEY"); v == "" {
		t.Fatal("IC_API_KEY must be set for acceptance tests")
	}

}

func TestAccPreCheckEnterpriseAccountImport(t *testing.T) {
	if v := os.Getenv("IC_API_KEY"); v == "" {
		t.Fatal("IC_API_KEY must be set for acceptance tests")
	}
	if Account_to_be_imported == "" {
		t.Fatal("ACCOUNT_TO_BE_IMPORTED must be set for acceptance tests")
	}

}
func TestAccPreCheckCis(t *testing.T) {
	TestAccPreCheck(t)
	if CisInstance == "" {
		t.Fatal("IBM_CIS_INSTANCE must be set for acceptance tests")
	}
	if CisResourceGroup == "" {
		t.Fatal("IBM_CIS_RESOURCE_GROUP must be set for acceptance tests")
	}
	if CisDomainStatic == "" {
		t.Fatal("IBM_CIS_DOMAIN_STATIC must be set for acceptance tests")
	}
	if CisDomainTest == "" {
		t.Fatal("IBM_CIS_DOMAIN_TEST must be set for acceptance tests")
	}
}

func TestAccPreCheckCloudShell(t *testing.T) {
	TestAccPreCheck(t)
	if CloudShellAccountID == "" {
		t.Fatal("IBM_CLOUD_SHELL_ACCOUNT_ID must be set for acceptance tests")
	}
}

func TestAccPreCheckHPCS(t *testing.T) {
	TestAccPreCheck(t)
	if HpcsAdmin1 == "" {
		t.Fatal("IBM_HPCS_ADMIN1 must be set for acceptance tests")
	}
	if HpcsToken1 == "" {
		t.Fatal("IBM_HPCS_TOKEN1 must be set for acceptance tests")
	}
	if HpcsAdmin2 == "" {
		t.Fatal("IBM_HPCS_ADMIN2 must be set for acceptance tests")
	}
	if HpcsToken2 == "" {
		t.Fatal("IBM_HPCS_TOKEN2 must be set for acceptance tests")
	}
}
func TestAccPreCheckIAMTrustedProfile(t *testing.T) {
	TestAccPreCheck(t)
	if RealmName == "" {
		t.Fatal("IBM_IAM_REALM_NAME must be set for acceptance tests")
	}
	if IksSa == "" {
		t.Fatal("IBM_IAM_IKS_SA must be set for acceptance tests")
	}
}

func TestAccPreCheckCOS(t *testing.T) {
	TestAccPreCheck(t)
	if CosCRN == "" {
		t.Fatal("IBM_COS_CRN must be set for acceptance tests")
	}
}

func TestAccPreCheckImage(t *testing.T) {
	TestAccPreCheck(t)
	if Image_cos_url == "" {
		t.Fatal("IMAGE_COS_URL must be set for acceptance tests")
	}
	if Image_operating_system == "" {
		t.Fatal("IMAGE_OPERATING_SYSTEM must be set for acceptance tests")
	}
}
func TestAccPreCheckEncryptedImage(t *testing.T) {
	TestAccPreCheck(t)
	if Image_cos_url_encrypted == "" {
		t.Fatal("IMAGE_COS_URL_ENCRYPTED must be set for acceptance tests")
	}
	if Image_operating_system == "" {
		t.Fatal("IMAGE_OPERATING_SYSTEM must be set for acceptance tests")
	}
	if IsImageEncryptedDataKey == "" {
		t.Fatal("IS_IMAGE_ENCRYPTED_DATA_KEY must be set for acceptance tests")
	}
	if IsImageEncryptionKey == "" {
		t.Fatal("IS_IMAGE_ENCRYPTION_KEY must be set for acceptance tests")
	}
}
