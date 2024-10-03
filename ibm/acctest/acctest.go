// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package acctest

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	terraformsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const (
	ProviderName          = "ibm"
	ProviderNameAlternate = "ibmalternate"
)

var (
	AppIDTenantID                   string
	AppIDTestUserEmail              string
	BackupPolicyJobID               string
	BackupPolicyID                  string
	CfOrganization                  string
	CfSpace                         string
	CisDomainStatic                 string
	CisDomainTest                   string
	CisInstance                     string
	CisResourceGroup                string
	CloudShellAccountID             string
	CosCRN                          string
	BucketCRN                       string
	ActivityTrackerInstanceCRN      string
	MetricsMonitoringCRN            string
	BucketName                      string
	CosName                         string
	Ibmid1                          string
	Ibmid2                          string
	IAMUser                         string
	IAMAccountId                    string
	IAMServiceId                    string
	IAMTrustedProfileID             string
	Datacenter                      string
	MachineType                     string
	trustedMachineType              string
	PublicVlanID                    string
	PrivateVlanID                   string
	PrivateSubnetID                 string
	PublicSubnetID                  string
	SubnetID                        string
	LbaasDatacenter                 string
	LbaasSubnetId                   string
	LbListerenerCertificateInstance string
	IpsecDatacenter                 string
	Customersubnetid                string
	Customerpeerip                  string
	DedicatedHostName               string
	DedicatedHostID                 string
	KubeVersion                     string
	KubeUpdateVersion               string
	Zone                            string
	ZonePrivateVlan                 string
	ZonePublicVlan                  string
	ZoneUpdatePrivateVlan           string
	ZoneUpdatePublicVlan            string
	WorkerPoolSecondaryStorage      string
	CsRegion                        string
	ExtendedHardwareTesting         bool
	err                             error
	placementGroupName              string
	CertCRN                         string
	UpdatedCertCRN                  string
	SecretCRN                       string
	SecretCRN2                      string
	EnterpriseCRN                   string
	InstanceCRN                     string
	SecretGroupID                   string
	RegionName                      string
	ISZoneName                      string
	ISZoneName2                     string
	ISZoneName3                     string
	IsResourceGroupID               string
	ISResourceCrn                   string
	ISCIDR                          string
	ISCIDR2                         string
	ISPublicSSHKeyFilePath          string
	ISPrivateSSHKeyFilePath         string
	ISAddressPrefixCIDR             string
	InstanceName                    string
	InstanceProfileName             string
	InstanceProfileNameUpdate       string
	IsBareMetalServerProfileName    string
	IsBareMetalServerImage          string
	IsBareMetalServerImage2         string
	DNSInstanceCRN                  string
	DNSZoneID                       string
	DNSInstanceCRN1                 string
	DNSZoneID1                      string
	DedicatedHostProfileName        string
	DedicatedHostGroupID            string
	InstanceDiskProfileName         string
	DedicatedHostGroupFamily        string
	DedicatedHostGroupClass         string
	ShareProfileName                string
	SourceShareCRN                  string
	ShareEncryptionKey              string
	VNIId                           string
	VolumeProfileName               string
	VSIUnattachedBootVolumeID       string
	VSIDataVolumeID                 string
	ISRouteDestination              string
	ISRouteNextHop                  string
	ISSnapshotCRN                   string
	WorkspaceID                     string
	TemplateID                      string
	ActionID                        string
	JobID                           string
	RepoURL                         string
	RepoBranch                      string
	imageName                       string
	functionNamespace               string
	HpcsInstanceID                  string
)

// MQ on Cloud
var (
	MqcloudConfigEndpoint            string
	MqcloudInstanceID                string
	MqcloudQueueManagerID            string
	MqcloudKSCertFilePath            string
	MqcloudTSCertFilePath            string
	MqCloudQueueManagerLocation      string
	MqCloudQueueManagerVersion       string
	MqCloudQueueManagerVersionUpdate string
)

// Logs
var (
	LogsInstanceId                      string
	LogsInstanceRegion                  string
	LogsEventNotificationInstanceId     string
	LogsEventNotificationInstanceRegion string
)

// Secrets Manager
var (
	SecretsManagerInstanceID                                                        string
	SecretsManagerInstanceRegion                                                    string
	SecretsManagerENInstanceCrn                                                     string
	SecretsManagerIamCredentialsConfigurationApiKey                                 string
	SecretsManagerIamCredentialsSecretServiceId                                     string
	SecretsManagerIamCredentialsSecretServiceAccessGroup                            string
	SecretsManagerPublicCertificateLetsEncryptEnvironment                           string
	SecretsManagerPublicCertificateLetsEncryptPrivateKey                            string
	SecretsManagerPublicCertificateCisCrn                                           string
	SecretsManagerPublicCertificateClassicInfrastructureUsername                    string
	SecretsManagerPublicCertificateClassicInfrastructurePassword                    string
	SecretsManagerPublicCertificateCommonName                                       string
	SecretsManagerValidateManualDnsCisZoneId                                        string
	SecretsManagerImportedCertificatePathToCertificate                              string
	SecretsManagerServiceCredentialsCosCrn                                          string
	SecretsManagerPrivateCertificateConfigurationCryptoKeyIAMSecretServiceId        string
	SecretsManagerPrivateCertificateConfigurationCryptoKeyProviderType              string
	SecretsManagerPrivateCertificateConfigurationCryptoKeyProviderInstanceCrn       string
	SecretsManagerPrivateCertificateConfigurationCryptoKeyProviderPrivateKeystoreId string
	SecretsManagerSecretType                                                        string
	SecretsManagerSecretID                                                          string
)

var (
	HpcsAdmin1                string
	HpcsToken1                string
	HpcsAdmin2                string
	HpcsToken2                string
	HpcsRootKeyCrn            string
	RealmName                 string
	IksSa                     string
	IksClusterID              string
	IksClusterVpcID           string
	IksClusterSubnetID        string
	IksClusterResourceGroupID string
	IcdDbDeploymentId         string
	IcdDbBackupId             string
	IcdDbTaskId               string
	KmsInstanceID             string
	CrkID                     string
	KmsAccountID              string
	BaasEncryptionkeyCRN      string
)

// for snapshot encryption
var (
	IsKMSInstanceId string
	IsKMSKeyName    string
)

// For Power Colo

var (
	Pi_auxiliary_volume_name        string
	Pi_cloud_instance_id            string
	Pi_dhcp_id                      string
	Pi_host_group_id                string
	Pi_host_id                      string
	Pi_image                        string
	Pi_image_bucket_access_key      string
	Pi_image_bucket_file_name       string
	Pi_image_bucket_name            string
	Pi_image_bucket_region          string
	Pi_image_bucket_secret_key      string
	Pi_instance_name                string
	Pi_key_name                     string
	Pi_network_name                 string
	Pi_placement_group_name         string
	Pi_replication_volume_name      string
	Pi_resource_group_id            string
	Pi_sap_image                    string
	Pi_shared_processor_pool_id     string
	Pi_snapshot_id                  string
	Pi_spp_placement_group_id       string
	Pi_target_storage_tier          string
	Pi_volume_clone_task_id         string
	Pi_volume_group_id              string
	Pi_volume_group_name            string
	Pi_volume_id                    string
	Pi_volume_name                  string
	Pi_volume_onboarding_id         string
	Pi_volume_onboarding_source_crn string
	PiCloudConnectionName           string
	PiSAPProfileID                  string
	PiStoragePool                   string
	PiStorageType                   string
)

var (
	Pi_capture_storage_image_path       string
	Pi_capture_cloud_storage_access_key string
	Pi_capture_cloud_storage_secret_key string
)

var ISDelegegatedVPC string

// For Image

var (
	IsImageName             string
	IsImageName2            string
	IsImage                 string
	IsImage2                string
	IsImageEncryptedDataKey string
	IsImageEncryptionKey    string
	IsWinImage              string
	IsCosBucketName         string
	IsCosBucketCRN          string
	Image_cos_url           string
	Image_cos_url_encrypted string
	Image_operating_system  string
)

// Transit Gateway Power Virtual Server
var Tg_power_vs_network_id string

// Transit Gateway cross account
var (
	Tg_cross_network_account_id      string
	Tg_cross_network_account_api_key string
	Tg_cross_network_id              string
)

// Enterprise Management
var Account_to_be_imported string

// Billing Snapshot Configuration
var Cos_bucket string
var Cos_location string
var Cos_bucket_update string
var Cos_location_update string
var Cos_reports_folder string
var Snapshot_date_from string
var Snapshot_date_to string
var Snapshot_month string

// Secuity and Complinace Center
var (
	SccApiEndpoint            string
	SccEventNotificationsCRN  string
	SccInstanceID             string
	SccObjectStorageCRN       string
	SccObjectStorageBucket    string
	SccProviderTypeAttributes string
	SccProviderTypeID         string
	SccReportID               string
)

// ROKS Cluster
var ClusterName string

// Satellite instance
var (
	Satellite_location_id          string
	Satellite_Resource_instance_id string
)

// Dedicated host
var HostPoolID string

// Continuous Delivery
var (
	CdResourceGroupName              string
	CdAppConfigInstanceName          string
	CdKeyProtectInstanceName         string
	CdSecretsManagerInstanceName     string
	CdSlackChannelName               string
	CdSlackTeamName                  string
	CdSlackWebhook                   string
	CdJiraProjectKey                 string
	CdJiraApiUrl                     string
	CdJiraUsername                   string
	CdJiraApiToken                   string
	CdSaucelabsAccessKey             string
	CdSaucelabsUsername              string
	CdBitbucketRepoUrl               string
	CdGithubConsolidatedRepoUrl      string
	CdGitlabRepoUrl                  string
	CdHostedGitRepoUrl               string
	CdEventNotificationsInstanceName string
)

// VPN Server
var (
	ISCertificateCrn string
	ISClientCaCrn    string
)

// COS Replication Bucket
var IBM_AccountID_REPL string

// Atracker
var (
	IesApiKey    string
	IngestionKey string
	COSApiKey    string
)

// For Code Engine
var (
	CeResourceGroupID   string
	CeProjectId         string
	CeServiceInstanceID string
	CeResourceKeyID     string
	CeDomainMappingName string
	CeTLSCert           string
	CeTLSKey            string
	CeTLSKeyFilePath    string
	CeTLSCertFilePath   string
)

// Satellite tests
var (
	SatelliteSSHPubKey string
)

// for IAM Identity
var IamIdentityAssignmentTargetAccountId string

// Projects
var ProjectsConfigApiKey string

// For PAG
var (
	PagCosInstanceName         string
	PagCosBucketName           string
	PagCosBucketRegion         string
	PagVpcName                 string
	PagServicePlan             string
	PagVpcSubnetNameInstance_1 string
	PagVpcSubnetNameInstance_2 string
	PagVpcSgInstance_1         string
	PagVpcSgInstance_2         string
)

// For VMware as a Service
var (
	Vmaas_Directorsite_id      string
	Vmaas_Directorsite_pvdc_id string
)

// For IAM Access Management
var (
	TargetAccountId    string
	TargetEnterpriseId string
)

// For Partner Center Sell
var (
	PcsRegistrationAccountId                         string
	PcsOnboardingProductWithApprovedProgrammaticName string
	// one Onboarding product can only have one catalog product ever
	PcsOnboardingProductWithApprovedProgrammaticName2 string
	PcsOnboardingProductWithCatalogProduct            string
	PcsOnboardingCatalogProductId                     string
	PcsOnboardingCatalogPlanId                        string
	PcsIamServiceRegistrationId                       string
)

func init() {
	testlogger := os.Getenv("TF_LOG")
	if testlogger != "" {
		os.Setenv("IBMCLOUD_BLUEMIX_GO_TRACE", "true")
	}

	IamIdentityAssignmentTargetAccountId = os.Getenv("IAM_IDENTITY_ASSIGNMENT_TARGET_ACCOUNT")

	ProjectsConfigApiKey = os.Getenv("IBM_PROJECTS_CONFIG_APIKEY")
	if ProjectsConfigApiKey == "" {
		fmt.Println("[WARN] Set the environment variable IBM_PROJECTS_CONFIG_APIKEY for testing IBM Projects Config resources, the tests will fail if this is not set")
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

	IAMAccountId = os.Getenv("IBM_IAMACCOUNTID")
	if IAMAccountId == "" {
		fmt.Println("[WARN] Set the environment variable IBM_IAMACCOUNTID for testing ibm_iam_trusted_profile resource Some tests for that resource will fail if this is not set correctly")
	}

	IAMServiceId = os.Getenv("IBM_IAM_SERVICE_ID")
	if IAMAccountId == "" {
		fmt.Println("[WARN] Set the environment variable IBM_IAM_SERVICE_ID for testing ibm_iam_trusted_profile_identity resource Some tests for that resource will fail if this is not set correctly")
	}

	IAMTrustedProfileID = os.Getenv("IBM_IAM_TRUSTED_PROFILE_ID")
	if IAMTrustedProfileID == "" {
		fmt.Println("[WARN] Set the environment variable IBM_IAM_TRUSTED_PROFILE_ID for testing ibm_iam_trusted_profile_identity resource Some tests for that resource will fail if this is not set correctly")
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
		fmt.Println("[WARN] Set the environment variable IBM_CERT_CRN for testing ibm_container_alb_cert or ibm_container_ingress_secret_tls resource else it is set to default value")
	}

	UpdatedCertCRN = os.Getenv("IBM_UPDATE_CERT_CRN")
	if UpdatedCertCRN == "" {
		UpdatedCertCRN = "crn:v1:bluemix:public:cloudcerts:eu-de:a/e9021a4d06e9b108b4a221a3cec47e3d:77e527aa-65b2-4cb3-969b-7e8714174346:certificate:1bf3d0c2b7764402dde25744218e6cba"
		fmt.Println("[WARN] Set the environment variable IBM_UPDATE_CERT_CRN for testing ibm_container_alb_cert or ibm_container_ingress_secret_tls resource else it is set to default value")
	}

	SecretCRN = os.Getenv("IBM_SECRET_CRN")
	if SecretCRN == "" {
		SecretCRN = "crn:v1:bluemix:public:secrets-manager:us-south:a/52b2e14f385aca5da781baa1b9c28e53:6efac0c2-b955-49ca-939d-d7bc0cb8132f:secret:e786b0ea2af8b5435603803ec2ff8118"
		fmt.Println("[WARN] Set the environment variable IBM_SECRET_CRN for testing ibm_container_ingress_secret_opaque resource else it is set to default value")
	}

	SecretCRN2 = os.Getenv("IBM_SECRET_CRN_2")
	if SecretCRN2 == "" {
		SecretCRN2 = "crn:v1:bluemix:public:secrets-manager:eu-de:a/e9021a4d06e9b108b4a221a3cec47e3d:77e527aa-65b2-4cb3-969b-7e8714174346:secret:1bf3d0c2b7764402dde25744218e6cba"
		fmt.Println("[WARN] Set the environment variable IBM_SECRET_CRN_2 for testing ibm_container_ingress_secret_opaque resource else it is set to default value")
	}

	InstanceCRN = os.Getenv("IBM_INGRESS_INSTANCE_CRN")
	if InstanceCRN == "" {
		fmt.Println("[WARN] Set the environment variable IBM_INGRESS_INSTANCE_CRN for testing ibm_container_ingress_instance resource. Some tests for that resource will fail if this is not set correctly")
	}

	SecretGroupID = os.Getenv("IBM_INGRESS_INSTANCE_SECRET_GROUP_ID")
	if SecretGroupID == "" {
		fmt.Println("[WARN] Set the environment variable IBM_INGRESS_INSTANCE_SECRET_GROUP_ID for testing ibm_container_ingress_instance resource. Some tests for that resource will fail if this is not set correctly")
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
	BucketCRN = os.Getenv("IBM_COS_Bucket_CRN")
	if BucketCRN == "" {
		BucketCRN = ""
		fmt.Println("[WARN] Set the environment variable IBM_COS_Bucket_CRN with a VALID BUCKET CRN for testing ibm_cos_bucket* resources")
	}
	ActivityTrackerInstanceCRN = os.Getenv("IBM_COS_ACTIVITY_TRACKER_CRN")
	if ActivityTrackerInstanceCRN == "" {
		ActivityTrackerInstanceCRN = ""
		fmt.Println("[WARN] Set the environment variable IBM_COS_ACTIVITY_TRACKER_CRN with a VALID ACTIVITY TRACKER INSTANCE CRN in valid region for testing ibm_cos_bucket* resources")
	}
	MetricsMonitoringCRN = os.Getenv("IBM_COS_METRICS_MONITORING_CRN")
	if MetricsMonitoringCRN == "" {
		MetricsMonitoringCRN = ""
		fmt.Println("[WARN] Set the environment variable IBM_COS_METRICS_MONITORING_CRN with a VALID METRICS MONITORING CRN for testing ibm_cos_bucket* resources")
	}
	BucketName = os.Getenv("IBM_COS_BUCKET_NAME")
	if BucketName == "" {
		BucketName = ""
		fmt.Println("[WARN] Set the environment variable IBM_COS_BUCKET_NAME with a VALID BUCKET Name for testing ibm_cos_bucket* resources")
	}

	CosName = os.Getenv("IBM_COS_NAME")
	if CosName == "" {
		CosName = ""
		fmt.Println("[WARN] Set the environment variable IBM_COS_NAME with a VALID COS instance name for testing resources with cos deps")
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
		KubeVersion = "1.25.9"
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

	WorkerPoolSecondaryStorage = os.Getenv("IBM_WORKER_POOL_SECONDARY_STORAGE")
	if WorkerPoolSecondaryStorage == "" {
		fmt.Println("[WARN] Set the environment variable IBM_WORKER_POOL_SECONDARY_STORAGE for testing secondary_storage attachment to IKS workerpools")
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

	ISZoneName2 = os.Getenv("SL_ZONE_2")
	if ISZoneName2 == "" {
		ISZoneName2 = "us-south-2"
		fmt.Println("[INFO] Set the environment variable SL_ZONE_2 for testing ibm_is_zone datasource else it is set to default value 'us-south-2'")
	}

	ISZoneName3 = os.Getenv("SL_ZONE_3")
	if ISZoneName3 == "" {
		ISZoneName3 = "us-south-3"
		fmt.Println("[INFO] Set the environment variable SL_ZONE_3 for testing ibm_is_zone datasource else it is set to default value 'us-south-3'")
	}

	ISCIDR = os.Getenv("SL_CIDR")
	if ISCIDR == "" {
		ISCIDR = "10.240.0.0/24"
		fmt.Println("[INFO] Set the environment variable SL_CIDR for testing ibm_is_subnet else it is set to default value '10.240.0.0/24'")
	}

	ISCIDR2 = os.Getenv("SL_CIDR_2")
	if ISCIDR2 == "" {
		ISCIDR2 = "10.240.64.0/24"
		fmt.Println("[INFO] Set the environment variable SL_CIDR_2 for testing ibm_is_subnet else it is set to default value '10.240.64.0/24'")
	}

	ISCIDR2 = os.Getenv("SL_CIDR_2")
	if ISCIDR2 == "" {
		ISCIDR2 = "10.240.64.0/24"
		fmt.Println("[INFO] Set the environment variable SL_CIDR_2 for testing ibm_is_subnet else it is set to default value '10.240.64.0/24'")
	}

	ISPublicSSHKeyFilePath = os.Getenv("IS_PUBLIC_SSH_KEY_PATH")
	if ISPublicSSHKeyFilePath == "" {
		ISPublicSSHKeyFilePath = "./test-fixtures/.ssh/pkcs8_rsa.pub"
		fmt.Println("[INFO] Set the environment variable SL_CIDR_2 for testing ibm_is_instance datasource else it is set to default value './test-fixtures/.ssh/pkcs8_rsa.pub'")
	}

	ISPrivateSSHKeyFilePath = os.Getenv("IS_PRIVATE_SSH_KEY_PATH")
	if ISPrivateSSHKeyFilePath == "" {
		ISPrivateSSHKeyFilePath = "./test-fixtures/.ssh/pkcs8_rsa"
		fmt.Println("[INFO] Set the environment variable IS_PRIVATE_SSH_KEY_PATH for testing ibm_is_instance datasource else it is set to default value './test-fixtures/.ssh/pkcs8_rsa'")
	}

	IsResourceGroupID = os.Getenv("SL_RESOURCE_GROUP_ID")
	if IsResourceGroupID == "" {
		IsResourceGroupID = "c01d34dff4364763476834c990398zz8"
		fmt.Println("[INFO] Set the environment variable SL_RESOURCE_GROUP_ID for testing with different resource group id else it is set to default value 'c01d34dff4364763476834c990398zz8'")
	}
	ISResourceCrn = os.Getenv("IS_RESOURCE_INSTANCE_CRN")
	if ISResourceCrn == "" {
		ISResourceCrn = "crn:v1:bluemix:public:cloud-object-storage:global:a/fugeggfcgjebvrburvgurgvugfr:236764224-f48fu4-f4h84-9db3-4f94fh::"
		fmt.Println("[INFO] Set the environment variable IS_RESOURCE_CRN for testing with created resource instance")
	}

	IsImage = os.Getenv("IS_IMAGE")
	if IsImage == "" {
		// IsImage = "fc538f61-7dd6-4408-978c-c6b85b69fe76" // for classic infrastructure
		IsImage = "r006-907911a7-0ffe-467e-8821-3cc9a0d82a39" // for next gen infrastructure ibm-centos-7-9-minimal-amd64-10 image
		fmt.Println("[INFO] Set the environment variable IS_IMAGE for testing ibm_is_instance, ibm_is_floating_ip else it is set to default value 'r006-907911a7-0ffe-467e-8821-3cc9a0d82a39'")
	}

	IsImage2 = os.Getenv("IS_IMAGE2")
	if IsImage2 == "" {
		IsImage2 = "r134-f47cc24c-e020-4db5-ad96-1e5be8b5853b" // for next gen infrastructure ibm-centos-7-9-minimal-amd64-10 image
		fmt.Println("[INFO] Set the environment variable IS_IMAGE2 for testing ibm_is_instance, ibm_is_floating_ip else it is set to default value 'r134-f47cc24c-e020-4db5-ad96-1e5be8b5853b'")
	}

	IsWinImage = os.Getenv("IS_WIN_IMAGE")
	if IsWinImage == "" {
		// IsWinImage = "a7a0626c-f97e-4180-afbe-0331ec62f32a" // classic windows machine: ibm-windows-server-2012-full-standard-amd64-1
		IsWinImage = "r006-d2e0d0e9-0a4f-4c45-afd7-cab787030776" // next gen windows machine: ibm-windows-server-2022-full-standard-amd64-8
		fmt.Println("[INFO] Set the environment variable IS_WIN_IMAGE for testing ibm_is_instance data source else it is set to default value 'r006-d2e0d0e9-0a4f-4c45-afd7-cab787030776'")
	}

	IsCosBucketName = os.Getenv("IS_COS_BUCKET_NAME")
	if IsCosBucketName == "" {
		IsCosBucketName = "test-bucket"
		fmt.Println("[INFO] Set the environment variable IS_COS_BUCKET_NAME for testing ibm_is_image_export_job else it is set to default value 'bucket-27200-lwx4cfvcue'")
	}

	IsCosBucketCRN = os.Getenv("IS_COS_BUCKET_CRN")
	if IsCosBucketCRN == "" {
		IsCosBucketCRN = "crn:v1:bluemix:public:cloud-object-storage:global:a/XXXXXXXX:XXXXX-XXXX-XXXX-XXXX-XXXX:bucket:test-bucket"
		fmt.Println("[INFO] Set the environment variable IS_COS_BUCKET_CRN for testing ibm_is_image_export_job else it is set to default value 'bucket-27200-lwx4cfvcue'")
	}

	InstanceName = os.Getenv("IS_INSTANCE_NAME")
	if InstanceName == "" {
		InstanceName = "placement-check-ins" // for next gen infrastructure
		fmt.Println("[INFO] Set the environment variable IS_INSTANCE_NAME for testing ibm_is_instance resource else it is set to default value 'instance-01'")
	}

	BackupPolicyJobID = os.Getenv("IS_BACKUP_POLICY_JOB_ID")
	if BackupPolicyJobID == "" {
		fmt.Println("[INFO] Set the environment variable IS_BACKUP_POLICY_JOB_ID for testing ibm_is_backup_policy_job datasource")
	}

	BackupPolicyID = os.Getenv("IS_BACKUP_POLICY_ID")
	if BackupPolicyID == "" {
		fmt.Println("[INFO] Set the environment variable IS_BACKUP_POLICY_ID for testing ibm_is_backup_policy_jobs datasource")
	}

	BaasEncryptionkeyCRN = os.Getenv("IS_REMOTE_CP_BAAS_ENCRYPTION_KEY_CRN")
	if BaasEncryptionkeyCRN == "" {
		BaasEncryptionkeyCRN = "crn:v1:bluemix:public:kms:us-south:a/dffc98a0f1f0f95f6613b3b752286b87:e4a29d1a-2ef0-42a6-8fd2-350deb1c647e:key:5437653b-c4b1-447f-9646-b2a2a4cd6179"
		fmt.Println("[INFO] Set the environment variable IS_REMOTE_CP_BAAS_ENCRYPTION_KEY_CRN for testing remote_copies_policy with Baas plans, else it is set to default value, 'crn:v1:bluemix:public:kms:us-south:a/dffc98a0f1f0f95f6613b3b752286b87:e4a29d1a-2ef0-42a6-8fd2-350deb1c647e:key:5437653b-c4b1-447f-9646-b2a2a4cd6179'")
	}

	InstanceProfileName = os.Getenv("SL_INSTANCE_PROFILE")
	if InstanceProfileName == "" {
		// InstanceProfileName = "bc1-2x8" // for classic infrastructure
		InstanceProfileName = "cx2-2x4" // for next gen infrastructure
		fmt.Println("[INFO] Set the environment variable SL_INSTANCE_PROFILE for testing ibm_is_instance resource else it is set to default value 'cx2-2x4'")
	}

	IsKMSInstanceId = os.Getenv("SL_KMS_INSTANCE_ID")
	if IsKMSInstanceId == "" {
		IsKMSInstanceId = "30222bb5-1c6d-3834-8d78-ae6348cf8z61" // kms instance id
		fmt.Println("[INFO] Set the environment variable SL_KMS_INSTANCE_ID for testing ibm_kms_key resource else it is set to default value '30222bb5-1c6d-3834-8d78-ae6348cf8z61'")
	}

	IsKMSKeyName = os.Getenv("SL_KMS_KEY_NAME")
	if IsKMSKeyName == "" {
		IsKMSKeyName = "tfp-test-key" // kms instance key name
		fmt.Println("[INFO] Set the environment variable SL_KMS_KEY_NAME for testing ibm_kms_key resource else it is set to default value 'tfp-test-key'")
	}

	InstanceProfileNameUpdate = os.Getenv("SL_INSTANCE_PROFILE_UPDATE")
	if InstanceProfileNameUpdate == "" {
		InstanceProfileNameUpdate = "cx2-4x8"
		fmt.Println("[INFO] Set the environment variable SL_INSTANCE_PROFILE_UPDATE for testing ibm_is_instance resource else it is set to default value 'cx2-4x8'")
	}

	IsBareMetalServerProfileName = os.Getenv("IS_BARE_METAL_SERVER_PROFILE")
	if IsBareMetalServerProfileName == "" {
		IsBareMetalServerProfileName = "bx2-metal-96x384" // for next gen infrastructure
		fmt.Println("[INFO] Set the environment variable IS_BARE_METAL_SERVER_PROFILE for testing ibm_is_bare_metal_server resource else it is set to default value 'bx2-metal-96x384'")
	}

	IsBareMetalServerImage = os.Getenv("IS_BARE_METAL_SERVER_IMAGE")
	if IsBareMetalServerImage == "" {
		IsBareMetalServerImage = "r006-2d1f36b0-df65-4570-82eb-df7ae5f778b1" // for next gen infrastructure
		fmt.Println("[INFO] Set the environment variable IsBareMetalServerImage for testing ibm_is_bare_metal_server resource else it is set to default value 'r006-2d1f36b0-df65-4570-82eb-df7ae5f778b1'")
	}

	IsBareMetalServerImage2 = os.Getenv("IS_BARE_METAL_SERVER_IMAGE2")
	if IsBareMetalServerImage2 == "" {
		IsBareMetalServerImage2 = "r006-2d1f36b0-df65-4570-82eb-df7ae5f778b1" // for next gen infrastructure
		fmt.Println("[INFO] Set the environment variable IsBareMetalServerImage2 for testing ibm_is_bare_metal_server resource else it is set to default value 'r006-2d1f36b0-df65-4570-82eb-df7ae5f778b1'")
	}

	DNSInstanceCRN = os.Getenv("IS_DNS_INSTANCE_CRN")
	if DNSInstanceCRN == "" {
		DNSInstanceCRN = "crn:v1:bluemix:public:dns-svcs:global:a/7f75c7b025e54bc5635f754b2f888665:fa78ce08-a161-4703-98e5-35ed2bfe0e7c::" // for next gen infrastructure
		fmt.Println("[INFO] Set the environment variable IS_DNS_INSTANCE_CRN for testing ibm_is_lb resource else it is set to default value 'crn:v1:staging:public:dns-svcs:global:a/efe5afc483594adaa8325e2b4d1290df:82df2e3c-53a5-43c6-89ce-dcf78be18668::'")
	}

	DNSInstanceCRN1 = os.Getenv("IS_DNS_INSTANCE_CRN1")
	if DNSInstanceCRN1 == "" {
		DNSInstanceCRN1 = "crn:v1:bluemix:public:dns-svcs:global:a/7f75c7b025e54bc5635f754b2f888665:fa78ce08-a161-4703-98e5-35ed2bfe0e7c::" // for next gen infrastructure
		fmt.Println("[INFO] Set the environment variable IS_DNS_INSTANCE_CRN1 for testing ibm_is_lb resource else it is set to default value 'crn:v1:staging:public:dns-svcs:global:a/efe5afc483594adaa8325e2b4d1290df:599ae4aa-c554-4a88-8bb2-b199b9a3c046::'")
	}

	DNSZoneID = os.Getenv("IS_DNS_ZONE_ID")
	if DNSZoneID == "" {
		DNSZoneID = "9519a5f8-8827-426b-8623-22226affcb7e" // for next gen infrastructure
		fmt.Println("[INFO] Set the environment variable IS_DNS_ZONE_ID for testing ibm_is_lb resource else it is set to default value 'dd501d1d-490b-4bb4-a05d-a31954a1c59e'")
	}

	DNSZoneID1 = os.Getenv("IS_DNS_ZONE_ID_1")
	if DNSZoneID1 == "" {
		DNSZoneID1 = "c4cdfb45-c21e-4ae3-88da-c9f64ad91d22" // for next gen infrastructure
		fmt.Println("[INFO] Set the environment variable IS_DNS_ZONE_ID_1 for testing ibm_is_lb resource else it is set to default value 'b1def78d-51b3-4ea5-a746-1b64c992fcab'")
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
		// InstanceProfileName = "bc1-2x8" // for classic infrastructure
		InstanceDiskProfileName = "bx2d-16x64" // for next gen infrastructure
		fmt.Println("[INFO] Set the environment variable SL_INSTANCE_PROFILE for testing ibm_is_instance resource else it is set to default value 'bx2d-16x64'")
	}

	ShareProfileName = os.Getenv("IS_SHARE_PROFILE")
	if ShareProfileName == "" {
		ShareProfileName = "dp2" // for next gen infrastructure
		fmt.Println("[INFO] Set the environment variable IS_SHARE_PROFILE for testing ibm_is_instance resource else it is set to default value 'tier-3iops'")
	}

	SourceShareCRN = os.Getenv("IS_SOURCE_SHARE_CRN")
	if SourceShareCRN == "" {
		SourceShareCRN = "crn:v1:staging:public:is:us-east-1:a/efe5afc483594adaa8325e2b4d1290df::share:r142-a106f162-86e4-4d7f-be75-193cc55a93e9" // for next gen infrastructure
		fmt.Println("[INFO] Set the environment variable IS_SHARE_PROFILE for testing ibm_is_instance resource else it is set to default value")
	}

	ShareEncryptionKey = os.Getenv("IS_SHARE_ENCRYPTION_KEY")
	if ShareEncryptionKey == "" {
		ShareEncryptionKey = "crn:v1:staging:public:kms:us-south:a/efe5afc483594adaa8325e2b4d1290df:1be45161-6dae-44ca-b248-837f98004057:key:3dd21cc5-cc20-4f7c-bc62-8ec9a8a3d1bd" // for next gen infrastructure
		fmt.Println("[INFO] Set the environment variable IS_SHARE_PROFILE for testing ibm_is_instance resource else it is set to default value")
	}

	VolumeProfileName = os.Getenv("IS_VOLUME_PROFILE")
	if VolumeProfileName == "" {
		VolumeProfileName = "general-purpose"
		fmt.Println("[INFO] Set the environment variable IS_VOLUME_PROFILE for testing ibm_is_volume_profile else it is set to default value 'general-purpose'")
	}

	VNIId = os.Getenv("IS_VIRTUAL_NETWORK_INTERFACE")
	if VNIId == "" {
		VNIId = "c93dc4c6-e85a-4da2-9ea6-f24576256122"
		fmt.Println("[INFO] Set the environment variable IS_VIRTUAL_NETWORK_INTERFACE for testing ibm_is_virtual_network_interface else it is set to default value 'c93dc4c6-e85a-4da2-9ea6-f24576256122'")
	}

	VSIUnattachedBootVolumeID = os.Getenv("IS_VSI_UNATTACHED_BOOT_VOLUME_ID")
	if VSIUnattachedBootVolumeID == "" {
		VSIUnattachedBootVolumeID = "r006-1cbe9f0a-7101-4d25-ae72-2a2d725e530e"
		fmt.Println("[INFO] Set the environment variable IS_UNATTACHED_BOOT_VOLUME_NAME for testing ibm_is_image else it is set to default value 'r006-1cbe9f0a-7101-4d25-ae72-2a2d725e530e'")
	}
	VSIDataVolumeID = os.Getenv("IS_VSI_DATA_VOLUME_ID")
	if VSIDataVolumeID == "" {
		VSIDataVolumeID = "r006-1cbe9f0a-7101-4d25-ae72-2a2d725e530e"
		fmt.Println("[INFO] Set the environment variable IS_VSI_DATA_VOLUME_ID for testing ibm_is_image else it is set to default value 'r006-1cbe9f0a-7101-4d25-ae72-2a2d725e530e'")
	}

	ISRouteNextHop = os.Getenv("SL_ROUTE_NEXTHOP")
	if ISRouteNextHop == "" {
		ISRouteNextHop = "10.240.0.0"
		fmt.Println("[INFO] Set the environment variable SL_ROUTE_NEXTHOP else it is set to default value '10.0.0.4'")
	}

	ISSnapshotCRN = os.Getenv("IS_SNAPSHOT_CRN")
	if ISSnapshotCRN == "" {
		ISSnapshotCRN = "crn:v1:bluemix:public:is:ca-tor:a/xxxxxxxx::snapshot:xxxx-xxxxc-xxx-xxxx-xxxx-xxxxxxxxxx"
		fmt.Println("[INFO] Set the environment variable ISSnapshotCRN for ibm_is_snapshot resource else it is set to default value 'crn:v1:bluemix:public:is:ca-tor:a/xxxxxxxx::snapshot:xxxx-xxxxc-xxx-xxxx-xxxx-xxxxxxxxxx'")
	}

	IcdDbDeploymentId = os.Getenv("ICD_DB_DEPLOYMENT_ID")
	if IcdDbDeploymentId == "" {
		IcdDbDeploymentId = "crn:v1:bluemix:public:databases-for-redis:au-syd:a/40ddc34a953a8c02f10987b59085b60e:5042afe1-72c2-4231-89cc-c949e5d56251::"
		fmt.Println("[INFO] Set the environment variable ICD_DB_DEPLOYMENT_ID for testing ibm_cloud_databases else it is set to default value 'crn:v1:bluemix:public:databases-for-redis:au-syd:a/40ddc34a953a8c02f10987b59085b60e:5042afe1-72c2-4231-89cc-c949e5d56251::'")
	}

	IcdDbBackupId = os.Getenv("ICD_DB_BACKUP_ID")
	if IcdDbBackupId == "" {
		IcdDbBackupId = "crn:v1:bluemix:public:databases-for-redis:au-syd:a/40ddc34a953a8c02f10987b59085b60e:5042afe1-72c2-4231-89cc-c949e5d56251:backup:0d862fdb-4faa-42e5-aecb-5057f4d399c3"
		fmt.Println("[INFO] Set the environment variable ICD_DB_BACKUP_ID for testing ibm_cloud_databases else it is set to default value 'crn:v1:bluemix:public:databases-for-redis:au-syd:a/40ddc34a953a8c02f10987b59085b60e:5042afe1-72c2-4231-89cc-c949e5d56251:backup:0d862fdb-4faa-42e5-aecb-5057f4d399c3'")
	}

	IcdDbTaskId = os.Getenv("ICD_DB_TASK_ID")
	if IcdDbTaskId == "" {
		IcdDbTaskId = "crn:v1:bluemix:public:databases-for-redis:au-syd:a/40ddc34a953a8c02f10987b59085b60e:367b0a22-05bb-41e3-a1ed-ded1ff0889e5:task:882013a6-2751-4df7-a77a-98d258638704"
		fmt.Println("[INFO] Set the environment variable ICD_DB_TASK_ID for testing ibm_cloud_databases else it is set to default value 'crn:v1:bluemix:public:databases-for-redis:au-syd:a/40ddc34a953a8c02f10987b59085b60e:367b0a22-05bb-41e3-a1ed-ded1ff0889e5:task:882013a6-2751-4df7-a77a-98d258638704'")
	}
	// Added for Power Colo Testing
	Pi_image = os.Getenv("PI_IMAGE")
	if Pi_image == "" {
		Pi_image = "c93dc4c6-e85a-4da2-9ea6-f24576256122"
		fmt.Println("[INFO] Set the environment variable PI_IMAGE for testing ibm_pi_image resource else it is set to default value '7200-03-03'")
	}
	Pi_sap_image = os.Getenv("PI_SAP_IMAGE")
	if Pi_sap_image == "" {
		Pi_sap_image = "2e29d6d2-e5ed-4ff8-8fad-64e4be90e023"
		fmt.Println("[INFO] Set the environment variable PI_SAP_IMAGE for testing ibm_pi_image resource else it is set to default value 'Linux-RHEL-SAP-8-2'")
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

	Pi_volume_id = os.Getenv("PI_VOLUME_ID")
	if Pi_volume_id == "" {
		Pi_volume_id = "terraform-test-power"
		fmt.Println("[INFO] Set the environment variable PI_VOLUME_ID for testing ibm_pi_volume_flash_copy_mappings resource else it is set to default value 'terraform-test-power'")
	}

	Pi_replication_volume_name = os.Getenv("PI_REPLICATION_VOLUME_NAME")
	if Pi_replication_volume_name == "" {
		Pi_replication_volume_name = "terraform-test-power"
		fmt.Println("[INFO] Set the environment variable PI_REPLICATION_VOLUME_NAME for testing ibm_pi_volume resource else it is set to default value 'terraform-test-power'")
	}

	Pi_volume_onboarding_source_crn = os.Getenv("PI_VOLUME_ONBARDING_SOURCE_CRN")
	if Pi_volume_onboarding_source_crn == "" {
		Pi_volume_onboarding_source_crn = "terraform-test-power"
		fmt.Println("[INFO] Set the environment variable PI_VOLUME_ONBARDING_SOURCE_CRN for testing ibm_pi_volume_onboarding resource else it is set to default value 'terraform-test-power'")
	}

	Pi_auxiliary_volume_name = os.Getenv("PI_AUXILIARY_VOLUME_NAME")
	if Pi_auxiliary_volume_name == "" {
		Pi_auxiliary_volume_name = "terraform-test-power"
		fmt.Println("[INFO] Set the environment variable PI_AUXILIARY_VOLUME_NAME for testing ibm_pi_volume_onboarding resource else it is set to default value 'terraform-test-power'")
	}

	Pi_volume_group_name = os.Getenv("PI_VOLUME_GROUP_NAME")
	if Pi_volume_group_name == "" {
		Pi_volume_group_name = "terraform-test-power"
		fmt.Println("[INFO] Set the environment variable PI_VOLUME_GROUP_NAME for testing ibm_pi_volume_group resource else it is set to default value 'terraform-test-power'")
	}

	Pi_volume_group_id = os.Getenv("PI_VOLUME_GROUP_ID")
	if Pi_volume_group_id == "" {
		Pi_volume_group_id = "terraform-test-power"
		fmt.Println("[INFO] Set the environment variable PI_VOLUME_GROUP_ID for testing ibm_pi_volume_group_storage_details data source else it is set to default value 'terraform-test-power'")
	}

	Pi_volume_onboarding_id = os.Getenv("PI_VOLUME_ONBOARDING_ID")
	if Pi_volume_onboarding_id == "" {
		Pi_volume_onboarding_id = "terraform-test-power"
		fmt.Println("[INFO] Set the environment variable PI_VOLUME_ONBOARDING_ID for testing ibm_pi_volume_onboarding resource else it is set to default value 'terraform-test-power'")
	}

	Pi_cloud_instance_id = os.Getenv("PI_CLOUDINSTANCE_ID")
	if Pi_cloud_instance_id == "" {
		Pi_cloud_instance_id = "fd3454a3-14d8-4eb0-b075-acf3da5cd324"
		fmt.Println("[INFO] Set the environment variable PI_CLOUDINSTANCE_ID for testing ibm_pi_image resource else it is set to default value 'd16705bd-7f1a-48c9-9e0e-1c17b71e7331'")
	}

	Pi_snapshot_id = os.Getenv("PI_SNAPSHOT_ID")
	if Pi_snapshot_id == "" {
		Pi_snapshot_id = "1ea33118-4c43-4356-bfce-904d0658de82"
		fmt.Println("[INFO] Set the environment variable PI_SNAPSHOT_ID for testing ibm_pi_instance_snapshot data source else it is set to default value '1ea33118-4c43-4356-bfce-904d0658de82'")
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
	Pi_spp_placement_group_id = os.Getenv("PI_SPP_PLACEMENT_GROUP_ID")
	if Pi_spp_placement_group_id == "" {
		Pi_spp_placement_group_id = "tf-pi-spp-placement-group"
		fmt.Println("[WARN] Set the environment variable PI_SPP_PLACEMENT_GROUP_ID for testing ibm_pi_spp_placement_group resource else it is set to default value 'tf-pi-spp-placement-group'")
	}
	PiStoragePool = os.Getenv("PI_STORAGE_POOL")
	if PiStoragePool == "" {
		PiStoragePool = "terraform-test-power"
		fmt.Println("[INFO] Set the environment variable PI_STORAGE_POOL for testing ibm_pi_storage_pool_capacity else it is set to default value 'terraform-test-power'")
	}
	PiStorageType = os.Getenv("PI_STORAGE_TYPE")
	if PiStorageType == "" {
		PiStorageType = "terraform-test-power"
		fmt.Println("[INFO] Set the environment variable PI_STORAGE_TYPE for testing ibm_pi_storage_type_capacity else it is set to default value 'terraform-test-power'")
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

	Pi_shared_processor_pool_id = os.Getenv("PI_SHARED_PROCESSOR_POOL_ID")
	if Pi_shared_processor_pool_id == "" {
		Pi_shared_processor_pool_id = "tf-pi-shared-processor-pool"
		fmt.Println("[WARN] Set the environment variable PI_SHARED_PROCESSOR_POOL_ID for testing ibm_pi_shared_processor_pool resource else it is set to default value 'tf-pi-shared-processor-pool'")
	}

	Pi_target_storage_tier = os.Getenv("PI_TARGET_STORAGE_TIER")
	if Pi_target_storage_tier == "" {
		Pi_target_storage_tier = "terraform-test-tier"
		fmt.Println("[INFO] Set the environment variable PI_TARGET_STORAGE_TIER for testing Pi_target_storage_tier resource else it is set to default value 'terraform-test-tier'")
	}

	Pi_volume_clone_task_id = os.Getenv("PI_VOLUME_CLONE_TASK_ID")
	if Pi_volume_clone_task_id == "" {
		Pi_volume_clone_task_id = "terraform-test-volume-clone-task-id"
		fmt.Println("[INFO] Set the environment variable PI_VOLUME_CLONE_TASK_ID for testing Pi_volume_clone_task_id resource else it is set to default value 'terraform-test-volume-clone-task-id'")
	}

	Pi_resource_group_id = os.Getenv("PI_RESOURCE_GROUP_ID")
	if Pi_resource_group_id == "" {
		Pi_resource_group_id = ""
		fmt.Println("[WARN] Set the environment variable PI_RESOURCE_GROUP_ID for testing ibm_pi_workspace resource else it is set to default value ''")
	}
	Pi_host_group_id = os.Getenv("PI_HOST_GROUP_ID")
	if Pi_host_group_id == "" {
		Pi_host_group_id = ""
		fmt.Println("[WARN] Set the environment variable PI_HOST_GROUP_ID for testing ibm_pi_host resource else it is set to default value ''")
	}

	Pi_host_id = os.Getenv("PI_HOST_ID")
	if Pi_host_id == "" {
		Pi_host_id = ""
		fmt.Println("[WARN] Set the environment variable PI_HOST_ID for testing ibm_pi_host resource else it is set to default value ''")
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

	ISDelegegatedVPC = os.Getenv("IS_DELEGATED_VPC")
	if ISDelegegatedVPC == "" {
		ISDelegegatedVPC = "tfp-test-vpc-hub-false-del"
		fmt.Println("[WARN] Set the environment variable IS_DELEGATED_VPC with a VALID created vpc name for testing ibm_is_vpc data source on staging/test")
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
		// IsImageName = "ibm-ubuntu-18-04-2-minimal-amd64-1" // for classic infrastructure
		IsImageName = "ibm-ubuntu-22-04-1-minimal-amd64-4" // for next gen infrastructure
		fmt.Println("[INFO] Set the environment variable IS_IMAGE_NAME for testing data source ibm_is_image else it is set to default value `ibm-ubuntu-18-04-1-minimal-amd64-2`")
	}

	IsImageName2 = os.Getenv("IS_IMAGE_NAME2")
	if IsImageName2 == "" {
		IsImageName2 = "ibm-ubuntu-20-04-6-minimal-amd64-5" // for next gen infrastructure
		fmt.Println("[INFO] Set the environment variable IS_IMAGE_NAME2 for testing data source ibm_is_image else it is set to default value `ibm-ubuntu-20-04-6-minimal-amd64-5`")
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
		fmt.Println("[INFO] Set the environment variable SECRETS_MANAGER_INSTANCE_ID for testing Secrets Manager's tests else tests will fail if this is not set correctly")
	}

	SecretsManagerInstanceRegion = os.Getenv("SECRETS_MANAGER_INSTANCE_REGION")
	if SecretsManagerInstanceRegion == "" {
		fmt.Println("[INFO] Set the environment variable SECRETS_MANAGER_INSTANCE_REGION for testing Secrets Manager's tests else tests will fail if this is not set correctly")
	}

	SecretsManagerENInstanceCrn = os.Getenv("SECRETS_MANAGER_EN_INSTANCE_CRN")
	if SecretsManagerENInstanceCrn == "" {
		fmt.Println("[INFO] Set the environment variable SECRETS_MANAGER_EN_INSTANCE_CRN for testing Event Notifications for Secrets Manager tests else tests will fail if this is not set correctly")
	}

	SecretsManagerIamCredentialsConfigurationApiKey = os.Getenv("SECRETS_MANAGER_IAM_CREDENTIALS_CONFIGURATION_API_KEY")
	if SecretsManagerENInstanceCrn == "" {
		fmt.Println("[INFO] Set the environment variable SECRETS_MANAGER_EN_INSTANCE_CRN for testing IAM Credentials secret's tests else tests will assume that IAM Credentials engine is already configured and fail if not set correctly")
	}

	SecretsManagerIamCredentialsSecretServiceId = os.Getenv("SECRETS_MANAGER_IAM_CREDENTIALS_SECRET_SERVICE_ID")
	SecretsManagerIamCredentialsSecretServiceAccessGroup = os.Getenv("SECRETS_MANAGER_IAM_CREDENTIALS_SECRET_ACCESS_GROUP")
	if SecretsManagerIamCredentialsSecretServiceId == "" && SecretsManagerIamCredentialsSecretServiceAccessGroup == "" {
		fmt.Println("[INFO] Set the environment variable SECRETS_MANAGER_IAM_CREDENTIALS_SECRET_SERVICE_ID or SECRETS_MANAGER_IAM_CREDENTIALS_SECRET_ACCESS_GROUP for testing IAM Credentials secret's tests, else tests fail if not set correctly")
	}

	SecretsManagerPublicCertificateLetsEncryptEnvironment = os.Getenv("SECRETS_MANAGER_PUBLIC_CERTIFICATE_LETS_ENCRYPT_ENVIRONMENT")
	if SecretsManagerPublicCertificateLetsEncryptEnvironment == "" {
		SecretsManagerPublicCertificateLetsEncryptEnvironment = "production"
		fmt.Println("[INFO] Set the environment variable SECRETS_MANAGER_PUBLIC_CERTIFICATE_LETS_ENCRYPT_ENVIRONMENT for testing public certificate's tests, else it is set to default value ('production'). For public certificate's tests, tests will fail if this is not set correctly")
	}

	SecretsManagerPublicCertificateLetsEncryptPrivateKey = os.Getenv("SECRETS_MANAGER_PUBLIC_CERTIFICATE_LETS_ENCRYPT_PRIVATE_KEY")
	if SecretsManagerPublicCertificateLetsEncryptPrivateKey == "" {
		fmt.Println("[INFO] Set the environment variable SECRETS_MANAGER_PUBLIC_CERTIFICATE_LETS_ENCRYPT_PRIVATE_KEY for testing public certificate's tests, else tests fail if not set correctly")
	}

	SecretsManagerPublicCertificateCommonName = os.Getenv("SECRETS_MANAGER_PUBLIC_CERTIFICATE_COMMON_NAME")
	if SecretsManagerPublicCertificateCommonName == "" {
		fmt.Println("[INFO] Set the environment variable SECRETS_MANAGER_PUBLIC_CERTIFICATE_COMMON_NAME for testing public certificate's tests, else tests fail if not set correctly")
	}

	SecretsManagerValidateManualDnsCisZoneId = os.Getenv("SECRETS_MANAGER_VALIDATE_MANUAL_DNS_CIS_ZONE_ID")
	if SecretsManagerValidateManualDnsCisZoneId == "" {
		fmt.Println("[INFO] Set the environment variable SECRETS_MANAGER_VALIDATE_MANUAL_DNS_CIS_ZONE_ID for testing validate manual dns' test, else tests fail if not set correctly")
	}

	SecretsManagerPublicCertificateCisCrn = os.Getenv("SECRETS_MANAGER_PUBLIC_CERTIFICATE_CIS_CRN")
	if SecretsManagerPublicCertificateCisCrn == "" {
		fmt.Println("[INFO] Set the environment variable SECRETS_MANAGER_PUBLIC_CERTIFICATE_CIS_CRN for testing public certificate's tests, else tests fail if not set correctly")
	}

	SecretsManagerPublicCertificateClassicInfrastructureUsername = os.Getenv("SECRETS_MANAGER_PUBLIC_CLASSIC_INFRASTRUCTURE_USERNAME")
	if SecretsManagerPublicCertificateClassicInfrastructureUsername == "" {
		fmt.Println("[INFO] Set the environment variable SECRETS_MANAGER_PUBLIC_CLASSIC_INFRASTRUCTURE_USERNAME for testing public certificate's tests, else tests fail if not set correctly")
	}

	SecretsManagerPublicCertificateClassicInfrastructurePassword = os.Getenv("SECRETS_MANAGER_PUBLIC_CLASSIC_INFRASTRUCTURE_PASSWORD")
	if SecretsManagerPublicCertificateClassicInfrastructurePassword == "" {
		fmt.Println("[INFO] Set the environment variable SECRETS_MANAGER_PUBLIC_CLASSIC_INFRASTRUCTURE_PASSWORD for testing public certificate's tests, else tests fail if not set correctly")
	}

	SecretsManagerImportedCertificatePathToCertificate = os.Getenv("SECRETS_MANAGER_IMPORTED_CERTIFICATE_PATH_TO_CERTIFICATE")
	if SecretsManagerImportedCertificatePathToCertificate == "" {
		fmt.Println("[INFO] Set the environment variable SECRETS_MANAGER_IMPORTED_CERTIFICATE_PATH_TO_CERTIFICATE for testing imported certificate's tests, else tests fail if not set correctly")
	}

	SecretsManagerServiceCredentialsCosCrn = os.Getenv("SECRETS_MANAGER_SERVICE_CREDENTIALS_COS_CRN")
	if SecretsManagerServiceCredentialsCosCrn == "" {
		fmt.Println("[INFO] Set the environment variable SECRETS_MANAGER_SERVICE_CREDENTIALS_COS_CRN for testing service credentials' tests, else tests fail if not set correctly")
	}

	SecretsManagerPrivateCertificateConfigurationCryptoKeyIAMSecretServiceId = os.Getenv("SECRETS_MANAGER_PRIVATE_CERTIFICATE_CONFIGURATION_CRYPTO_KEY_IAM_SECRET_SERVICE_ID")
	if SecretsManagerPrivateCertificateConfigurationCryptoKeyIAMSecretServiceId == "" {
		fmt.Println("[INFO] Set the environment variable SECRETS_MANAGER_PRIVATE_CERTIFICATE_CONFIGURATION_CRYPTO_KEY_IAM_SECRET_SERVICE_ID for testing private certificate's configuration with crypto key tests, else tests fail if not set correctly")
	}

	SecretsManagerPrivateCertificateConfigurationCryptoKeyProviderType = os.Getenv("SECRETS_MANAGER_PRIVATE_CERTIFICATE_CONFIGURATION_CRYPTO_KEY_PROVIDER_TYPE")
	if SecretsManagerPrivateCertificateConfigurationCryptoKeyProviderType == "" {
		fmt.Println("[INFO] Set the environment variable SECRETS_MANAGER_PRIVATE_CERTIFICATE_CONFIGURATION_CRYPTO_KEY_PROVIDER_TYPE for testing private certificate's configuration with crypto key tests, else tests fail if not set correctly")
	}

	SecretsManagerPrivateCertificateConfigurationCryptoKeyProviderInstanceCrn = os.Getenv("SECRETS_MANAGER_PRIVATE_CERTIFICATE_CONFIGURATION_CRYPTO_KEY_PROVIDER_INSTANCE_CRN")
	if SecretsManagerPrivateCertificateConfigurationCryptoKeyProviderInstanceCrn == "" {
		fmt.Println("[INFO] Set the environment variable SECRETS_MANAGER_PRIVATE_CERTIFICATE_CONFIGURATION_CRYPTO_KEY_PROVIDER_INSTANCE_CRN for testing private certificate's configuration with crypto key tests, else tests fail if not set correctly")
	}

	SecretsManagerPrivateCertificateConfigurationCryptoKeyProviderPrivateKeystoreId = os.Getenv("SECRETS_MANAGER_PRIVATE_CERTIFICATE_CONFIGURATION_CRYPTO_KEY_PROVIDER_PRIVATE_KEYSTORE_ID")
	if SecretsManagerPrivateCertificateConfigurationCryptoKeyProviderPrivateKeystoreId == "" {
		fmt.Println("[INFO] Set the environment variable SECRETS_MANAGER_PRIVATE_CERTIFICATE_CONFIGURATION_CRYPTO_KEY_PROVIDER_PRIVATE_KEYSTORE_ID for testing private certificate's configuration with crypto key tests, else tests fail if not set correctly")
	}

	Tg_cross_network_account_api_key = os.Getenv("IBM_TG_CROSS_ACCOUNT_API_KEY")
	if Tg_cross_network_account_api_key == "" {
		fmt.Println("[INFO] Set the environment variable IBM_TG_CROSS_ACCOUNT_API_KEY for testing ibm_tg_connection resource else  tests will fail if this is not set correctly")
	}
	Tg_cross_network_account_id = os.Getenv("IBM_TG_CROSS_ACCOUNT_ID")
	if Tg_cross_network_account_id == "" {
		fmt.Println("[INFO] Set the environment variable IBM_TG_CROSS_ACCOUNT_ID for testing ibm_tg_connection resource else  tests will fail if this is not set correctly")
	}
	Tg_cross_network_id = os.Getenv("IBM_TG_CROSS_NETWORK_ID")
	if Tg_cross_network_id == "" {
		fmt.Println("[INFO] Set the environment variable IBM_TG_CROSS_NETWORK_ID for testing ibm_tg_connection resource else  tests will fail if this is not set correctly")
	}
	Tg_power_vs_network_id = os.Getenv("IBM_TG_POWER_VS_NETWORK_ID")
	if Tg_power_vs_network_id == "" {
		fmt.Println("[INFO] Set the environment variable IBM_TG_POWER_VS_NETWORK_ID for testing ibm_tg_connection resource else tests will fail if this is not set correctly")
	}
	Account_to_be_imported = os.Getenv("ACCOUNT_TO_BE_IMPORTED")
	if Account_to_be_imported == "" {
		fmt.Println("[INFO] Set the environment variable ACCOUNT_TO_BE_IMPORTED for testing import enterprise account resource else  tests will fail if this is not set correctly")
	}
	Cos_bucket = os.Getenv("COS_BUCKET")
	if Cos_bucket == "" {
		fmt.Println("[INFO] Set the environment variable COS_BUCKET for testing CRUD operations on billing snapshot configuration APIs")
	}
	Cos_location = os.Getenv("COS_LOCATION")
	if Cos_location == "" {
		fmt.Println("[INFO] Set the environment variable COS_LOCATION for testing CRUD operations on billing snapshot configuration APIs")
	}
	Cos_bucket_update = os.Getenv("COS_BUCKET_UPDATE")
	if Cos_bucket_update == "" {
		fmt.Println("[INFO] Set the environment variable COS_BUCKET_UPDATE for testing update operation on billing snapshot configuration API")
	}
	Cos_location_update = os.Getenv("COS_LOCATION_UPDATE")
	if Cos_location_update == "" {
		fmt.Println("[INFO] Set the environment variable COS_LOCATION_UPDATE for testing update operation on billing snapshot configuration API")
	}
	Cos_reports_folder = os.Getenv("COS_REPORTS_FOLDER")
	if Cos_reports_folder == "" {
		fmt.Println("[INFO] Set the environment variable COS_REPORTS_FOLDER for testing CRUD operations on billing snapshot configuration APIs")
	}
	Snapshot_date_from = os.Getenv("SNAPSHOT_DATE_FROM")
	if Snapshot_date_from == "" {
		fmt.Println("[INFO] Set the environment variable SNAPSHOT_DATE_FROM for testing CRUD operations on billing snapshot configuration APIs")
	}
	Snapshot_date_to = os.Getenv("SNAPSHOT_DATE_TO")
	if Snapshot_date_to == "" {
		fmt.Println("[INFO] Set the environment variable SNAPSHOT_DATE_TO for testing CRUD operations on billing snapshot configuration APIs")
	}
	Snapshot_month = os.Getenv("SNAPSHOT_MONTH")
	if Snapshot_month == "" {
		fmt.Println("[INFO] Set the environment variable SNAPSHOT_MONTH for testing CRUD operations on billing snapshot configuration APIs")
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
	HpcsRootKeyCrn = os.Getenv("IBM_HPCS_ROOTKEY_CRN")
	if HpcsRootKeyCrn == "" {
		fmt.Println("[WARN] Set the environment variable IBM_HPCS_ROOTKEY_CRN with a VALID CRN for a root key created in the HPCS instance")
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

	Satellite_location_id = os.Getenv("SATELLITE_LOCATION_ID")
	if Satellite_location_id == "" {
		fmt.Println("[INFO] Set the environment variable SATELLITE_LOCATION_ID for ibm_cos_bucket satellite location resource or datasource else tests will fail if this is not set correctly")
	}

	Satellite_Resource_instance_id = os.Getenv("SATELLITE_RESOURCE_INSTANCE_ID")
	if Satellite_Resource_instance_id == "" {
		fmt.Println("[INFO] Set the environment variable SATELLITE_RESOURCE_INSTANCE_ID for ibm_cos_bucket satellite location resource or datasource else tests will fail if this is not set correctly")
	}

	SccInstanceID = os.Getenv("IBMCLOUD_SCC_INSTANCE_ID")
	if SccInstanceID == "" {
		fmt.Println("[WARN] Set the environment variable IBMCLOUD_SCC_INSTANCE_ID with a VALID SCC INSTANCE ID")
	}

	SccApiEndpoint = os.Getenv("IBMCLOUD_SCC_API_ENDPOINT")
	if SccApiEndpoint == "" {
		fmt.Println("[WARN] Set the environment variable IBMCLOUD_SCC_API_ENDPOINT with a VALID SCC API ENDPOINT")
	}

	SccReportID = os.Getenv("IBMCLOUD_SCC_REPORT_ID")
	if SccReportID == "" {
		fmt.Println("[WARN] Set the environment variable IBMCLOUD_SCC_REPORT_ID with a VALID SCC REPORT ID")
	}

	SccProviderTypeAttributes = os.Getenv("IBMCLOUD_SCC_PROVIDER_TYPE_ATTRIBUTES")
	if SccProviderTypeAttributes == "" {
		fmt.Println("[WARN] Set the environment variable IBMCLOUD_SCC_PROVIDER_TYPE_ATTRIBUTES with a VALID SCC PROVIDER TYPE ATTRIBUTE")
	}

	SccProviderTypeID = os.Getenv("IBMCLOUD_SCC_PROVIDER_TYPE_ID")
	if SccProviderTypeID == "" {
		fmt.Println("[WARN] Set the environment variable IBMCLOUD_SCC_PROVIDER_TYPE_ID with a VALID SCC PROVIDER TYPE ID")
	}

	SccEventNotificationsCRN = os.Getenv("IBMCLOUD_SCC_EVENT_NOTIFICATION_CRN")
	if SccEventNotificationsCRN == "" {
		fmt.Println("[WARN] Set the environment variable IBMCLOUD_SCC_EVENT_NOTIFICATION_CRN")
	}

	SccObjectStorageCRN = os.Getenv("IBMCLOUD_SCC_OBJECT_STORAGE_CRN")
	if SccObjectStorageCRN == "" {
		fmt.Println("[WARN] Set the environment variable IBMCLOUD_SCC_OBJECT_STORAGE_CRN with a valid cloud object storage crn")
	}

	SccObjectStorageBucket = os.Getenv("IBMCLOUD_SCC_OBJECT_STORAGE_BUCKET")
	if SccObjectStorageBucket == "" {
		fmt.Println("[WARN] Set the environment variable IBMCLOUD_SCC_OBJECT_STORAGE_BUCKET with a valid cloud object storage bucket")
	}

	HostPoolID = os.Getenv("IBM_CONTAINER_DEDICATEDHOST_POOL_ID")
	if HostPoolID == "" {
		fmt.Println("[INFO] Set the environment variable IBM_CONTAINER_DEDICATEDHOST_POOL_ID for ibm_container_vpc_cluster resource to test dedicated host functionality")
	}

	KmsInstanceID = os.Getenv("IBM_KMS_INSTANCE_ID")
	if KmsInstanceID == "" {
		fmt.Println("[INFO] Set the environment variable IBM_KMS_INSTANCE_ID for ibm_container_vpc_cluster resource or datasource else tests will fail if this is not set correctly")
	}

	CrkID = os.Getenv("IBM_CRK_ID")
	if CrkID == "" {
		fmt.Println("[INFO] Set the environment variable IBM_CRK_ID for ibm_container_vpc_cluster resource or datasource else tests will fail if this is not set correctly")
	}

	KmsAccountID = os.Getenv("IBM_KMS_ACCOUNT_ID")
	if CrkID == "" {
		fmt.Println("[INFO] Set the environment variable IBM_KMS_ACCOUNT_ID for ibm_container_vpc_cluster resource or datasource else tests will fail if this is not set correctly")
	}

	IksClusterID = os.Getenv("IBM_CLUSTER_ID")
	if IksClusterID == "" {
		fmt.Println("[INFO] Set the environment variable IBM_CLUSTER_ID for ibm_container_vpc_worker_pool resource or datasource else tests will fail if this is not set correctly")
	}

	CdResourceGroupName = os.Getenv("IBM_CD_RESOURCE_GROUP_NAME")
	if CdResourceGroupName == "" {
		fmt.Println("[WARN] Set the environment variable IBM_CD_RESOURCE_GROUP_NAME for testing CD resources, CD tests will fail if this is not set")
	}

	CdAppConfigInstanceName = os.Getenv("IBM_CD_APPCONFIG_INSTANCE_NAME")
	if CdAppConfigInstanceName == "" {
		fmt.Println("[WARN] Set the environment variable IBM_CD_APPCONFIG_INSTANCE_NAME for testing CD resources, CD tests will fail if this is not set")
	}

	CdKeyProtectInstanceName = os.Getenv("IBM_CD_KEYPROTECT_INSTANCE_NAME")
	if CdKeyProtectInstanceName == "" {
		fmt.Println("[WARN] Set the environment variable IBM_CD_KEYPROTECT_INSTANCE_NAME for testing CD resources, CD tests will fail if this is not set")
	}

	CdSecretsManagerInstanceName = os.Getenv("IBM_CD_SECRETSMANAGER_INSTANCE_NAME")
	if CdSecretsManagerInstanceName == "" {
		fmt.Println("[WARN] Set the environment variable IBM_CD_SECRETSMANAGER_INSTANCE_NAME for testing CD resources, CD tests will fail if this is not set")
	}

	CdSlackChannelName = os.Getenv("IBM_CD_SLACK_CHANNEL_NAME")
	if CdSecretsManagerInstanceName == "" {
		fmt.Println("[WARN] Set the environment variable IBM_CD_SLACK_CHANNEL_NAME for testing CD resources, CD tests will fail if this is not set")
	}
	CdSlackTeamName = os.Getenv("IBM_CD_SLACK_TEAM_NAME")
	if CdSecretsManagerInstanceName == "" {
		fmt.Println("[WARN] Set the environment variable IBM_CD_SLACK_TEAM_NAME for testing CD resources, CD tests will fail if this is not set")
	}
	CdSlackWebhook = os.Getenv("IBM_CD_SLACK_WEBHOOK")
	if CdSecretsManagerInstanceName == "" {
		fmt.Println("[WARN] Set the environment variable IBM_CD_SLACK_WEBHOOK for testing CD resources, CD tests will fail if this is not set")
	}

	CdJiraProjectKey = os.Getenv("IBM_CD_JIRA_PROJECT_KEY")
	if CdSecretsManagerInstanceName == "" {
		fmt.Println("[WARN] Set the environment variable IBM_CD_JIRA_PROJECT_KEY for testing CD resources, CD tests will fail if this is not set")
	}
	CdJiraApiUrl = os.Getenv("IBM_CD_JIRA_API_URL")
	if CdSecretsManagerInstanceName == "" {
		fmt.Println("[WARN] Set the environment variable IBM_CD_JIRA_API_URL for testing CD resources, CD tests will fail if this is not set")
	}
	CdJiraUsername = os.Getenv("IBM_CD_JIRA_USERNAME")
	if CdSecretsManagerInstanceName == "" {
		fmt.Println("[WARN] Set the environment variable IBM_CD_JIRA_USERNAME for testing CD resources, CD tests will fail if this is not set")
	}
	CdJiraApiToken = os.Getenv("IBM_CD_JIRA_API_TOKEN")
	if CdSecretsManagerInstanceName == "" {
		fmt.Println("[WARN] Set the environment variable IBM_CD_JIRA_API_TOKEN for testing CD resources, CD tests will fail if this is not set")
	}

	CdSaucelabsAccessKey = os.Getenv("IBM_CD_SAUCELABS_ACCESS_KEY")
	if CdSecretsManagerInstanceName == "" {
		fmt.Println("[WARN] Set the environment variable IBM_CD_SAUCELABS_ACCESS_KEY for testing CD resources, CD tests will fail if this is not set")
	}
	CdSaucelabsUsername = os.Getenv("IBM_CD_SAUCELABS_USERNAME")
	if CdSecretsManagerInstanceName == "" {
		fmt.Println("[WARN] Set the environment variable IBM_CD_SAUCELABS_USERNAME for testing CD resources, CD tests will fail if this is not set")
	}

	CdBitbucketRepoUrl = os.Getenv("IBM_CD_BITBUCKET_REPO_URL")
	if CdBitbucketRepoUrl == "" {
		fmt.Println("[WARN] Set the environment variable IBM_CD_BITBUCKET_REPO_URL for testing CD resources, CD tests will fail if this is not set")
	}

	CdGithubConsolidatedRepoUrl = os.Getenv("IBM_CD_GITHUB_CONSOLIDATED_REPO_URL")
	if CdGithubConsolidatedRepoUrl == "" {
		fmt.Println("[WARN] Set the environment variable IBM_CD_GITHUB_CONSOLIDATED_REPO_URL for testing CD resources, CD tests will fail if this is not set")
	}

	CdGitlabRepoUrl = os.Getenv("IBM_CD_GITLAB_REPO_URL")
	if CdGitlabRepoUrl == "" {
		fmt.Println("[WARN] Set the environment variable IBM_CD_GITLAB_REPO_URL for testing CD resources, CD tests will fail if this is not set")
	}

	CdHostedGitRepoUrl = os.Getenv("IBM_CD_HOSTED_GIT_REPO_URL")
	if CdHostedGitRepoUrl == "" {
		fmt.Println("[WARN] Set the environment variable IBM_CD_HOSTED_GIT_REPO_URL for testing CD resources, CD tests will fail if this is not set")
	}

	CdEventNotificationsInstanceName = os.Getenv("IBM_CD_EVENTNOTIFICATIONS_INSTANCE_NAME")
	if CdEventNotificationsInstanceName == "" {
		fmt.Println("[WARN] Set the environment variable IBM_CD_EVENTNOTIFICATIONS_INSTANCE_NAME for testing CD resources, CD tests will fail if this is not set")
	}

	ISCertificateCrn = os.Getenv("IS_CERTIFICATE_CRN")
	if ISCertificateCrn == "" {
		fmt.Println("[INFO] Set the environment variable IS_CERTIFICATE_CRN for testing ibm_is_vpn_server resource")
	}

	ISClientCaCrn = os.Getenv("IS_CLIENT_CA_CRN")
	if ISClientCaCrn == "" {
		fmt.Println("[INFO] Set the environment variable IS_CLIENT_CA_CRN for testing ibm_is_vpn_server resource")
	}

	IBM_AccountID_REPL = os.Getenv("IBM_AccountID_REPL")
	if IBM_AccountID_REPL == "" {
		fmt.Println("[INFO] Set the environment variable IBM_AccountID_REPL for setting up authorization policy to enable replication feature resource or datasource else tests will fail if this is not set correctly")
	}

	COSApiKey = os.Getenv("COS_API_KEY")
	if COSApiKey == "" {
		COSApiKey = "xxxxxxxxxxxx" // pragma: allowlist secret
		fmt.Println("[WARN] Set the environment variable COS_API_KEY for testing COS targets, the tests will fail if this is not set")
	}

	IngestionKey = os.Getenv("INGESTION_KEY")
	if IngestionKey == "" {
		IngestionKey = "xxxxxxxxxxxx"
		fmt.Println("[WARN] Set the environment variable INGESTION_KEY for testing Logdna targets, the tests will fail if this is not set")
	}

	IesApiKey = os.Getenv("IES_API_KEY")
	if IesApiKey == "" {
		IesApiKey = "xxxxxxxxxxxx" // pragma: allowlist secret
		fmt.Println("[WARN] Set the environment variable IES_API_KEY for testing Event streams targets, the tests will fail if this is not set")
	}

	EnterpriseCRN = os.Getenv("ENTERPRISE_CRN")
	if EnterpriseCRN == "" {
		fmt.Println("[WARN] Set the environment variable ENTERPRISE_CRN for testing enterprise backup policy, the tests will fail if this is not set")
	}

	CeResourceGroupID = os.Getenv("IBM_CODE_ENGINE_RESOURCE_GROUP_ID")
	if CeResourceGroupID == "" {
		CeResourceGroupID = ""
		fmt.Println("[WARN] Set the environment variable IBM_CODE_ENGINE_RESOURCE_GROUP_ID with the resource group for Code Engine")
	}

	CeProjectId = os.Getenv("IBM_CODE_ENGINE_PROJECT_INSTANCE_ID")
	if CeProjectId == "" {
		CeProjectId = ""
		fmt.Println("[WARN] Set the environment variable IBM_CODE_ENGINE_PROJECT_INSTANCE_ID with the ID of a Code Engine project instance")
	}

	CeServiceInstanceID = os.Getenv("IBM_CODE_ENGINE_SERVICE_INSTANCE_ID")
	if CeServiceInstanceID == "" {
		CeServiceInstanceID = ""
		fmt.Println("[WARN] Set the environment variable IBM_CODE_ENGINE_SERVICE_INSTANCE_ID with the ID of a IBM Cloud service instance, e.g. for COS")
	}

	CeResourceKeyID = os.Getenv("IBM_CODE_ENGINE_RESOURCE_KEY_ID")
	if CeResourceKeyID == "" {
		CeResourceKeyID = ""
		fmt.Println("[WARN] Set the environment variable IBM_CODE_ENGINE_RESOURCE_KEY_ID with the ID of a resource key to access a service instance")
	}

	CeDomainMappingName = os.Getenv("IBM_CODE_ENGINE_DOMAIN_MAPPING_NAME")
	if CeDomainMappingName == "" {
		CeDomainMappingName = ""
		fmt.Println("[WARN] Set the environment variable IBM_CODE_ENGINE_DOMAIN_MAPPING_NAME with the name of a domain mapping")
	}

	CeTLSCert = os.Getenv("IBM_CODE_ENGINE_TLS_CERT")
	if CeTLSCert == "" {
		CeTLSCert = ""
		fmt.Println("[WARN] Set the environment variable IBM_CODE_ENGINE_TLS_CERT with the TLS certificate in base64 format")
	}

	CeTLSKey = os.Getenv("IBM_CODE_ENGINE_TLS_KEY")
	if CeTLSKey == "" {
		CeTLSKey = ""
		fmt.Println("[WARN] Set the environment variable IBM_CODE_ENGINE_TLS_KEY with a TLS key in base64 format")
	}

	CeTLSKeyFilePath = os.Getenv("IBM_CODE_ENGINE_TLS_CERT_KEY_PATH")
	if CeTLSKeyFilePath == "" {
		CeTLSKeyFilePath = ""
		fmt.Println("[WARN] Set the environment variable IBM_CODE_ENGINE_TLS_CERT_KEY_PATH to point to CERT KEY file path")
	}

	CeTLSCertFilePath = os.Getenv("IBM_CODE_ENGINE_TLS_CERT_PATH")
	if CeTLSCertFilePath == "" {
		CeTLSCertFilePath = ""
		fmt.Println("[WARN] Set the environment variable IBM_CODE_ENGINE_TLS_CERT_PATH to point to CERT file path")
	}

	SatelliteSSHPubKey = os.Getenv("IBM_SATELLITE_SSH_PUB_KEY")
	if SatelliteSSHPubKey == "" {
		fmt.Println("[WARN] Set the environment variable IBM_SATELLITE_SSH_PUB_KEY with a ssh public key or ibm_satellite_* tests may fail")
	}

	MqcloudConfigEndpoint = os.Getenv("IBMCLOUD_MQCLOUD_CONFIG_ENDPOINT")
	if MqcloudConfigEndpoint == "" {
		fmt.Println("[INFO] Set the environment variable IBMCLOUD_MQCLOUD_CONFIG_ENDPOINT for ibm_mqcloud service else tests will fail if this is not set correctly")
	}

	MqcloudInstanceID = os.Getenv("IBM_MQCLOUD_INSTANCE_ID")
	if MqcloudInstanceID == "" {
		fmt.Println("[INFO] Set the environment variable IBM_MQCLOUD_INSTANCE_ID for ibm_mqcloud_queue_manager resource or datasource else tests will fail if this is not set correctly")
	}
	MqcloudQueueManagerID = os.Getenv("IBM_MQCLOUD_QUEUEMANAGER_ID")
	if MqcloudQueueManagerID == "" {
		fmt.Println("[INFO] Set the environment variable IBM_MQCLOUD_QUEUEMANAGER_ID for ibm_mqcloud_queue_manager resource or datasource else tests will fail if this is not set correctly")
	}
	MqcloudKSCertFilePath = os.Getenv("IBM_MQCLOUD_KS_CERT_PATH")
	if MqcloudKSCertFilePath == "" {
		fmt.Println("[INFO] Set the environment variable IBM_MQCLOUD_KS_CERT_PATH for ibm_mqcloud_keystore_certificate resource or datasource else tests will fail if this is not set correctly")
	}
	MqcloudTSCertFilePath = os.Getenv("IBM_MQCLOUD_TS_CERT_PATH")
	if MqcloudTSCertFilePath == "" {
		fmt.Println("[INFO] Set the environment variable IBM_MQCLOUD_TS_CERT_PATH for ibm_mqcloud_truststore_certificate resource or datasource else tests will fail if this is not set correctly")
	}
	MqCloudQueueManagerLocation = os.Getenv(("IBM_MQCLOUD_QUEUEMANAGER_LOCATION"))
	if MqCloudQueueManagerLocation == "" {
		fmt.Println("[INFO] Set the environment variable IBM_MQCLOUD_QUEUEMANAGER_LOCATION for ibm_mqcloud_queue_manager resource or datasource else tests will fail if this is not set correctly")
	}
	MqCloudQueueManagerVersion = os.Getenv(("IBM_MQCLOUD_QUEUEMANAGER_VERSION"))
	if MqCloudQueueManagerVersion == "" {
		fmt.Println("[INFO] Set the environment variable IBM_MQCLOUD_QUEUEMANAGER_VERSION for ibm_mqcloud_queue_manager resource or datasource else tests will fail if this is not set correctly")
	}
	MqCloudQueueManagerVersionUpdate = os.Getenv(("IBM_MQCLOUD_QUEUEMANAGER_VERSIONUPDATE"))
	if MqCloudQueueManagerVersionUpdate == "" {
		fmt.Println("[INFO] Set the environment variable IBM_MQCLOUD_QUEUEMANAGER_VERSIONUPDATE for ibm_mqcloud_queue_manager resource or datasource else tests will fail if this is not set correctly")
	}
	LogsInstanceId = os.Getenv("IBMCLOUD_LOGS_SERVICE_INSTANCE_ID")
	if LogsInstanceId == "" {
		fmt.Println("[INFO] Set the environment variable IBMCLOUD_LOGS_SERVICE_INSTANCE_ID for testing cloud logs related operations")
	}
	LogsInstanceRegion = os.Getenv("IBMCLOUD_LOGS_SERVICE_INSTANCE_REGION")
	if LogsInstanceRegion == "" {
		fmt.Println("[INFO] Set the environment variable IBMCLOUD_LOGS_SERVICE_INSTANCE_REGION for testing cloud logs related operations")
	}
	LogsEventNotificationInstanceId = os.Getenv("IBMCLOUD_LOGS_SERVICE_EVENT_NOTIFICATIONS_INSTANCE_ID")
	if LogsEventNotificationInstanceId == "" {
		fmt.Println("[INFO] Set the environment variable IBMCLOUD_LOGS_SERVICE_EVENT_NOTIFICATIONS_INSTANCE_ID for testing cloud logs related operations")
	}
	LogsEventNotificationInstanceRegion = os.Getenv("IBMCLOUD_LOGS_SERVICE_EVENT_NOTIFICATIONS_INSTANCE_REGION")
	if LogsEventNotificationInstanceRegion == "" {
		fmt.Println("[INFO] Set the environment variable IBMCLOUD_LOGS_SERVICE_EVENT_NOTIFICATIONS_INSTANCE_REGION for testing cloud logs related operations")
	}

	PagCosInstanceName = os.Getenv("IBM_PAG_COS_INSTANCE_NAME")
	if PagCosInstanceName == "" {
		fmt.Println("[WARN] Set the environment variable IBM_PAG_COS_INSTANCE_NAME for testing IBM PAG resource, the tests will fail if this is not set")
	}

	PagCosBucketName = os.Getenv("IBM_PAG_COS_BUCKET_NAME")
	if PagCosBucketName == "" {
		fmt.Println("[WARN] Set the environment variable IBM_PAG_COS_BUCKET_NAME for testing IBM PAG resource, the tests will fail if this is not set")
	}

	PagCosBucketRegion = os.Getenv("IBM_PAG_COS_BUCKET_REGION")
	if PagCosBucketRegion == "" {
		fmt.Println("[WARN] Set the environment variable IBM_PAG_COS_BUCKET_REGION for testing IBM PAG resource, the tests will fail if this is not set")
	}

	PagVpcName = os.Getenv("IBM_PAG_VPC_NAME")
	if PagVpcName == "" {
		fmt.Println("[WARN] Set the environment variable IBM_PAG_VPC_NAME for testing IBM PAG resource, the tests will fail if this is not set")
	}

	PagServicePlan = os.Getenv("IBM_PAG_SERVICE_PLAN")
	if PagServicePlan == "" {
		fmt.Println("[WARN] Set the environment variable IBM_PAG_SERVICE_PLAN for testing IBM PAG resource, the tests will fail if this is not set")
	}

	PagVpcSubnetNameInstance_1 = os.Getenv("IBM_PAG_VPC_SUBNET_INS_1")
	if PagVpcSubnetNameInstance_1 == "" {
		fmt.Println("[WARN] Set the environment variable IBM_PAG_VPC_SUBNET_INS_1 for testing IBM PAG resource, the tests will fail if this is not set")
	}

	PagVpcSubnetNameInstance_2 = os.Getenv("IBM_PAG_VPC_SUBNET_INS_2")
	if PagVpcSubnetNameInstance_2 == "" {
		fmt.Println("[WARN] Set the environment variable IBM_PAG_VPC_SUBNET_INS_2 for testing IBM PAG resource, the tests will fail if this is not set")
	}
	PagVpcSgInstance_1 = os.Getenv("IBM_PAG_VPC_SG_SUBNET_INS_1")
	if PagVpcSubnetNameInstance_2 == "" {
		fmt.Println("[WARN] Set the environment variable IBM_PAG_VPC_SUBNET_INS_2 for testing IBM PAG resource, the tests will fail if this is not set")
	}
	PagVpcSgInstance_2 = os.Getenv("IBM_PAG_VPC_SG_SUBNET_INS_2")
	if PagVpcSubnetNameInstance_2 == "" {
		fmt.Println("[WARN] Set the environment variable IBM_PAG_VPC_SUBNET_INS_2 for testing IBM PAG resource, the tests will fail if this is not set")
	}

	// For vmware as a service
	Vmaas_Directorsite_id = os.Getenv("IBM_VMAAS_DS_ID")
	if Vmaas_Directorsite_id == "" {
		fmt.Println("[WARN] Set the environment variable IBM_VMAAS_DS_ID for testing ibm_vmaas_vdc resource else tests will fail if this is not set correctly")
	}

	Vmaas_Directorsite_pvdc_id = os.Getenv("IBM_VMAAS_DS_PVDC_ID")
	if Vmaas_Directorsite_pvdc_id == "" {
		fmt.Println("[WARN] Set the environment variable IBM_VMAAS_DS_PVDC_ID for testing ibm_vmaas_vdc resource else tests will fail if this is not set correctly")
	}

	TargetAccountId = os.Getenv("IBM_POLICY_ASSIGNMENT_TARGET_ACCOUNT_ID")
	if TargetAccountId == "" {
		fmt.Println("[INFO] Set the environment variable IBM_POLICY_ASSIGNMENT_TARGET_ACCOUNT_ID for testing ibm_iam_policy_assignment resource else tests will fail if this is not set correctly")
	}

	TargetEnterpriseId = os.Getenv("IBM_POLICY_ASSIGNMENT_TARGET_ENTERPRISE_ID")
	if TargetEnterpriseId == "" {
		fmt.Println("[INFO] Set the environment variable IBM_POLICY_ASSIGNMENT_TARGET_ENTERPRISE_ID for testing ibm_iam_policy_assignment resource else tests will fail if this is not set correctly")
	}

	PcsRegistrationAccountId = os.Getenv("PCS_REGISTRATION_ACCOUNT_ID")
	if PcsRegistrationAccountId == "" {
		fmt.Println("[WARN] Set the environment variable PCS_REGISTRATION_ACCOUNT_ID for testing iam_onboarding resource else tests will fail if this is not set correctly")
	}

	PcsOnboardingProductWithApprovedProgrammaticName = os.Getenv("PCS_PRODUCT_WITH_APPROVED_PROGRAMMATIC_NAME")
	if PcsOnboardingProductWithApprovedProgrammaticName == "" {
		fmt.Println("[WARN] Set the environment variable PCS_PRODUCT_WITH_APPROVED_PROGRAMMATIC_NAME for testing iam_onboarding resource else tests will fail if this is not set correctly")
	}

	PcsOnboardingProductWithApprovedProgrammaticName2 = os.Getenv("PCS_PRODUCT_WITH_APPROVED_PROGRAMMATIC_NAME_2")
	if PcsOnboardingProductWithApprovedProgrammaticName2 == "" {
		fmt.Println("[WARN] Set the environment variable PCS_PRODUCT_WITH_APPROVED_PROGRAMMATIC_NAME_2 for testing iam_onboarding resource else tests will fail if this is not set correctly")
	}

	PcsOnboardingProductWithCatalogProduct = os.Getenv("PCS_PRODUCT_WITH_CATALOG_PRODUCT")
	if PcsOnboardingProductWithCatalogProduct == "" {
		fmt.Println("[WARN] Set the environment variable PCS_PRODUCT_WITH_CATALOG_PRODUCT for testing iam_onboarding resource else tests will fail if this is not set correctly")
	}

	PcsOnboardingCatalogProductId = os.Getenv("PCS_CATALOG_PRODUCT")
	if PcsOnboardingCatalogProductId == "" {
		fmt.Println("[WARN] Set the environment variable PCS_CATALOG_PRODUCT for testing iam_onboarding resource else tests will fail if this is not set correctly")
	}

	PcsOnboardingCatalogPlanId = os.Getenv("PCS_CATALOG_PLAN")
	if PcsIamServiceRegistrationId == "" {
		fmt.Println("[WARN] Set the environment variable PCS_CATALOG_PLAN for testing iam_onboarding resource else tests will fail if this is not set correctly")
	}

	PcsIamServiceRegistrationId = os.Getenv("PCS_IAM_REGISTRATION_ID")
	if PcsIamServiceRegistrationId == "" {
		fmt.Println("[WARN] Set the environment variable PCS_IAM_TEGISTRATION_ID for testing iam_onboarding resource else tests will fail if this is not set correctly")
	}
}

var (
	TestAccProviders map[string]*schema.Provider
	TestAccProvider  *schema.Provider
)

// testAccProviderConfigure ensures Provider is only configured once
//
// The PreCheck(t) function is invoked for every test and this prevents
// extraneous reconfiguration to the same values each time. However, this does
// not prevent reconfiguration that may happen should the address of
// Provider be errantly reused in ProviderFactories.
var testAccProviderConfigure sync.Once

func init() {
	TestAccProvider = provider.Provider()
	TestAccProviders = map[string]*schema.Provider{
		ProviderName: TestAccProvider,
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

	testAccProviderConfigure.Do(func() {
		diags := TestAccProvider.Configure(context.Background(), terraformsdk.NewResourceConfigRaw(nil))
		if diags.HasError() {
			t.Fatalf("configuring provider: %s", diags[0].Summary)
		}
	})
}

func TestAccPreCheckEnterprise(t *testing.T) {
	if v := os.Getenv("IC_API_KEY"); v == "" {
		t.Fatal("IC_API_KEY must be set for acceptance tests")
	}
}

func TestAccPreCheckAssignmentTargetAccount(t *testing.T) {
	if v := os.Getenv("IAM_IDENTITY_ASSIGNMENT_TARGET_ACCOUNT"); v == "" {
		t.Fatal("IAM_IDENTITY_ASSIGNMENT_TARGET_ACCOUNT must be set for IAM identity assignment tests")
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
func TestAccPreCheckCloudLogs(t *testing.T) {
	if v := os.Getenv("IC_API_KEY"); v == "" {
		t.Fatal("IC_API_KEY must be set for acceptance tests")
	}
	if LogsInstanceId == "" {
		t.Fatal("IBMCLOUD_LOGS_SERVICE_INSTANCE_ID must be set for acceptance tests")
	}
	if LogsInstanceRegion == "" {
		t.Fatal("IBMCLOUD_LOGS_SERVICE_INSTANCE_REGION must be set for acceptance tests")
	}
	if LogsEventNotificationInstanceId == "" {
		t.Fatal("IBMCLOUD_LOGS_SERVICE_EVENT_NOTIFICATIONS_INSTANCE_ID must be set for acceptance tests")
	}
	if LogsEventNotificationInstanceRegion == "" {
		t.Fatal("IBMCLOUD_LOGS_SERVICE_EVENT_NOTIFICATIONS_INSTANCE_REGION must be set for acceptance tests")
	}

	testAccProviderConfigure.Do(func() {
		diags := TestAccProvider.Configure(context.Background(), terraformsdk.NewResourceConfigRaw(nil))
		if diags.HasError() {
			t.Fatalf("configuring provider: %s", diags[0].Summary)
		}
	})
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

func TestAccPreCheckCodeEngine(t *testing.T) {
	TestAccPreCheck(t)
	if CeResourceGroupID == "" {
		t.Fatal("IBM_CODE_ENGINE_RESOURCE_GROUP_ID must be set for acceptance tests")
	}
	if CeProjectId == "" {
		t.Fatal("IBM_CODE_ENGINE_PROJECT_INSTANCE_ID must be set for acceptance tests")
	}
	if CeServiceInstanceID == "" {
		t.Fatal("IBM_CODE_ENGINE_SERVICE_INSTANCE_ID must be set for acceptance tests")
	}
	if CeResourceKeyID == "" {
		t.Fatal("IBM_CODE_ENGINE_RESOURCE_KEY_ID must be set for acceptance tests")
	}
	if CeDomainMappingName == "" {
		t.Fatal("IBM_CODE_ENGINE_DOMAIN_MAPPING_NAME must be set for acceptance tests")
	}
	if CeTLSKeyFilePath == "" {
		t.Fatal("IBM_CODE_ENGINE_TLS_CERT_KEY_PATH must be set for acceptance tests")
	}
	if CeTLSCertFilePath == "" {
		t.Fatal("IBM_CODE_ENGINE_TLS_CERT_PATH must be set for acceptance tests")
	}
}

func TestAccPreCheckUsage(t *testing.T) {
	if v := os.Getenv("IC_API_KEY"); v == "" {
		t.Fatal("IC_API_KEY must be set for acceptance tests")
	}
}

func TestAccPreCheckScc(t *testing.T) {
	TestAccPreCheck(t)
	if SccApiEndpoint == "" {
		t.Fatal("IBMCLOUD_SCC_API_ENDPOINT missing. Set the environment variable IBMCLOUD_SCC_API_ENDPOINT with a VALID endpoint")
	}

	if SccProviderTypeAttributes == "" {
		t.Fatal("IBMCLOUD_SCC_PROVIDER_TYPE_ATTRIBUTES missing. Set the environment variable IBMCLOUD_SCC_PROVIDER_TYPE_ATTRIBUTES with a VALID SCC provider_type JSON object")
	}

	if SccProviderTypeID == "" {
		t.Fatal("IBMCLOUD_SCC_PROVIDER_TYPE_ID missing. Set the environment variable IBMCLOUD_SCC_PROVIDER_TYPE_ID with a VALID SCC provider_type ID")
	}

	if SccInstanceID == "" {
		t.Fatal("IBMCLOUD_SCC_INSTANCE_ID missing. Set the environment variable IBMCLOUD_SCC_INSTANCE_ID with a VALID SCC INSTANCE ID")
	}

	if SccReportID == "" {
		t.Fatal("IBMCLOUD_SCC_REPORT_ID missing. Set the environment variable IBMCLOUD_SCC_REPORT_ID with a VALID SCC REPORT_ID")
	}

	if SccEventNotificationsCRN == "" {
		t.Fatal("IBMCLOUD_SCC_EVENT_NOTIFICATION_CRN missing. Set the environment variable IBMCLOUD_SCC_EVENT_NOTIFICATION_CRN with a valid EN CRN")
	}

	if SccObjectStorageCRN == "" {
		t.Fatal("IBMCLOUD_SCC_OBJECT_STORAGE_CRN missing. Set the environment variable IBMCLOUD_SCC_OBJECT_STORAGE_CRN with a valid COS CRN")
	}

	if SccObjectStorageBucket == "" {
		t.Fatal("IBMCLOUD_SCC_OBJECT_STORAGE_CRN missing. Set the environment variable IBMCLOUD_SCC_OBJECT_STORAGE_BUCKET with a valid COS bucket")
	}
}

func TestAccPreCheckSatelliteSSH(t *testing.T) {
	TestAccPreCheck(t)
	if SatelliteSSHPubKey == "" {
		t.Fatal("IBM_SATELLITE_SSH_PUB_KEY missing. Set the environment variable IBM_SATELLITE_SSH_PUB_KEY with a VALID ssh public key")
	}
}

func TestAccPreCheckMqcloud(t *testing.T) {
	TestAccPreCheck(t)
	if MqcloudConfigEndpoint == "" {
		t.Fatal("IBMCLOUD_MQCLOUD_CONFIG_ENDPOINT must be set for acceptance tests")
	}
	if MqcloudInstanceID == "" {
		t.Fatal("IBM_MQCLOUD_INSTANCE_ID must be set for acceptance tests")
	}
	if MqcloudQueueManagerID == "" {
		t.Fatal("IBM_MQCLOUD_QUEUEMANAGER_ID must be set for acceptance tests")
	}
	if MqcloudTSCertFilePath == "" {
		t.Fatal("IBM_MQCLOUD_TS_CERT_PATH must be set for acceptance tests")
	}
	if MqcloudKSCertFilePath == "" {
		t.Fatal("IBM_MQCLOUD_KS_CERT_PATH must be set for acceptance tests")
	}
	if MqCloudQueueManagerLocation == "" {
		t.Fatal("IBM_MQCLOUD_QUEUEMANAGER_LOCATION must be set for acceptance tests")
	}
	if MqCloudQueueManagerVersion == "" {
		t.Fatal("IBM_MQCLOUD_QUEUEMANAGER_VERSION must be set for acceptance tests")
	}
	if MqCloudQueueManagerVersionUpdate == "" {
		t.Fatal("IBM_MQCLOUD_QUEUEMANAGER_VERSIONUPDATE must be set for acceptance tests")
	}
}

func TestAccPreCheckCbr(t *testing.T) {
	TestAccPreCheck(t)
	IAMAccountId = os.Getenv("IBM_IAMACCOUNTID")
	if IAMAccountId == "" {
		fmt.Println("[WARN] Set the environment variable IBM_IAMACCOUNTID for testing cbr related resources. Some tests for that resource will fail if this is not set correctly")
	}
	cbrEndpoint := os.Getenv("IBMCLOUD_CONTEXT_BASED_RESTRICTIONS_ENDPOINT")
	if cbrEndpoint == "" {
		fmt.Println("[WARN] Set the environment variable IBMCLOUD_CONTEXT_BASED_RESTRICTIONS_ENDPOINT for testing cbr related resources. Some tests for that resource will fail if this is not set correctly")
	}
}

func TestAccPreCheckVMwareService(t *testing.T) {
	if v := os.Getenv("IC_API_KEY"); v == "" {
		t.Fatal("IC_API_KEY must be set for acceptance tests")
	}
	if Vmaas_Directorsite_id == "" {
		t.Fatal("IBM_VMAAS_DS_ID must be set for acceptance tests")
	}
	if Vmaas_Directorsite_pvdc_id == "" {
		t.Fatal("IBM_VMAAS_DS_PVDC_ID must be set for acceptance tests")
	}
}

func TestAccPreCheckPartnerCenterSell(t *testing.T) {
	TestAccPreCheck(t)
	if PcsRegistrationAccountId == "" {
		t.Fatal("PCS_REGISTRATION_ACCOUNT_ID must be set for acceptance tests")
	}
	if PcsOnboardingProductWithApprovedProgrammaticName == "" {
		t.Fatal("PCS_PRODUCT_WITH_APPROVED_PROGRAMMATIC_NAME must be set for acceptance tests")
	}
	if PcsOnboardingProductWithApprovedProgrammaticName2 == "" {
		t.Fatal("PCS_PRODUCT_WITH_APPROVED_PROGRAMMATIC_NAME_2 must be set for acceptance tests")
	}
	if PcsOnboardingProductWithCatalogProduct == "" {
		t.Fatal("PCS_PRODUCT_WITH_CATALOG_PRODUCT must be set for acceptance tests")
	}
	if PcsOnboardingCatalogProductId == "" {
		t.Fatal("PCS_CATALOG_PRODUCT must be set for acceptance tests")
	}
	if PcsOnboardingCatalogPlanId == "" {
		t.Fatal("PCS_CATALOG_PLAN must be set for acceptance tests")
	}
	if PcsIamServiceRegistrationId == "" {
		t.Fatal("PCS_IAM_REGISTRATION_ID must be set for acceptance tests")
	}
}

func TestAccProviderFactories() map[string]func() (*schema.Provider, error) {
	return map[string]func() (*schema.Provider, error){
		ProviderName:          func() (*schema.Provider, error) { return provider.Provider(), nil },
		ProviderNameAlternate: func() (*schema.Provider, error) { return provider.Provider(), nil },
	}
}

func Region() string {
	region, _ := schema.MultiEnvDefaultFunc([]string{"IC_REGION", "IBMCLOUD_REGION", "BM_REGION", "BLUEMIX_REGION"}, "us-south")()

	return region.(string)
}

func RegionAlternate() string {
	region, _ := schema.MultiEnvDefaultFunc([]string{"IC_REGION_ALTERNATE", "IBMCLOUD_REGION_ALTERNATE"}, "eu-gb")()

	return region.(string)
}

func ConfigAlternateRegionProvider() string {
	return configNamedRegionalProvider(ProviderNameAlternate, RegionAlternate())
}

// ConfigCompose can be called to concatenate multiple strings to build test configurations
func ConfigCompose(config ...string) string {
	var str strings.Builder

	for _, conf := range config {
		str.WriteString(conf)
	}

	return str.String()
}

func configNamedRegionalProvider(providerName string, region string) string {
	return fmt.Sprintf(`
provider %[1]q {
  region = %[2]q
}
`, providerName, region)
}
