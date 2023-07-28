// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package provider

import (
	"os"
	"sync"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/apigateway"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appconfiguration"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/appid"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/atracker"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/catalogmanagement"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/cdtektonpipeline"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/cdtoolchain"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/cis"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/classicinfrastructure"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/cloudant"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/cloudfoundry"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/cloudshell"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/codeengine"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/contextbasedrestrictions"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/cos"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/database"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/directlink"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/dnsservices"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/enterprise"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/eventnotification"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/eventstreams"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/functions"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/globaltagging"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/hpcs"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/iamaccessgroup"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/iamidentity"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/iampolicy"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/kms"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/kubernetes"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/metricsrouter"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/power"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/project"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/pushnotification"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/registry"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/resourcecontroller"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/resourcemanager"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/satellite"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/scc"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/schematics"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/secretsmanager"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/transitgateway"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider returns a *schema.Provider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"bluemix_api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Bluemix API Key",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"BM_API_KEY", "BLUEMIX_API_KEY"}, nil),
				Deprecated:  "This field is deprecated please use ibmcloud_api_key",
			},
			"bluemix_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The timeout (in seconds) to set for any Bluemix API calls made.",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"BM_TIMEOUT", "BLUEMIX_TIMEOUT"}, nil),
				Deprecated:  "This field is deprecated please use ibmcloud_timeout",
			},
			"ibmcloud_api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The IBM Cloud API Key",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"IC_API_KEY", "IBMCLOUD_API_KEY"}, nil),
			},
			"ibmcloud_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The timeout (in seconds) to set for any IBM Cloud API calls made.",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"IC_TIMEOUT", "IBMCLOUD_TIMEOUT"}, 60),
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The IBM cloud Region (for example 'us-south').",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"IC_REGION", "IBMCLOUD_REGION", "BM_REGION", "BLUEMIX_REGION"}, "us-south"),
			},
			"zone": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The IBM cloud Region zone (for example 'us-south-1') for power resources.",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"IC_ZONE", "IBMCLOUD_ZONE"}, ""),
			},
			"resource_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Resource group id.",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"IC_RESOURCE_GROUP", "IBMCLOUD_RESOURCE_GROUP", "BM_RESOURCE_GROUP", "BLUEMIX_RESOURCE_GROUP"}, ""),
			},
			"softlayer_api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The SoftLayer API Key",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"SL_API_KEY", "SOFTLAYER_API_KEY"}, nil),
				Deprecated:  "This field is deprecated please use iaas_classic_api_key",
			},
			"softlayer_username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The SoftLayer user name",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"SL_USERNAME", "SOFTLAYER_USERNAME"}, nil),
				Deprecated:  "This field is deprecated please use iaas_classic_username",
			},
			"softlayer_endpoint_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Softlayer Endpoint",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"SL_ENDPOINT_URL", "SOFTLAYER_ENDPOINT_URL"}, nil),
				Deprecated:  "This field is deprecated please use iaas_classic_endpoint_url",
			},
			"softlayer_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The timeout (in seconds) to set for any SoftLayer API calls made.",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"SL_TIMEOUT", "SOFTLAYER_TIMEOUT"}, nil),
				Deprecated:  "This field is deprecated please use iaas_classic_timeout",
			},
			"iaas_classic_api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Classic Infrastructure API Key",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"IAAS_CLASSIC_API_KEY"}, nil),
			},
			"iaas_classic_username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Classic Infrastructure API user name",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"IAAS_CLASSIC_USERNAME"}, nil),
			},
			"iaas_classic_endpoint_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Classic Infrastructure Endpoint",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"IAAS_CLASSIC_ENDPOINT_URL"}, "https://api.softlayer.com/rest/v3"),
			},
			"iaas_classic_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The timeout (in seconds) to set for any Classic Infrastructure API calls made.",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"IAAS_CLASSIC_TIMEOUT"}, 60),
			},
			"max_retries": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The retry count to set for API calls.",
				DefaultFunc: schema.EnvDefaultFunc("MAX_RETRIES", 10),
			},
			"function_namespace": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The IBM Cloud Function namespace",
				DefaultFunc: schema.EnvDefaultFunc("FUNCTION_NAMESPACE", nil),
				Deprecated:  "This field will be deprecated soon",
			},
			"riaas_endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The next generation infrastructure service endpoint url.",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"RIAAS_ENDPOINT"}, nil),
				Deprecated:  "This field is deprecated use generation",
			},
			"generation": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Generation of Virtual Private Cloud. Default is 2",
				//DefaultFunc: schema.MultiEnvDefaultFunc([]string{"IC_GENERATION", "IBMCLOUD_GENERATION"}, nil),
				Deprecated: "The generation field is deprecated and will be removed after couple of releases",
			},
			"iam_profile_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IAM Trusted Profile Authentication token",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"IC_IAM_PROFILE_ID", "IBMCLOUD_IAM_PROFILE_ID"}, nil),
			},
			"iam_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IAM Authentication token",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"IC_IAM_TOKEN", "IBMCLOUD_IAM_TOKEN"}, nil),
			},
			"iam_refresh_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IAM Authentication refresh token",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"IC_IAM_REFRESH_TOKEN", "IBMCLOUD_IAM_REFRESH_TOKEN"}, nil),
			},
			"visibility": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.ValidateAllowedStringValues([]string{"public", "private", "public-and-private"}),
				Description:  "Visibility of the provider if it is private or public.",
				DefaultFunc:  schema.MultiEnvDefaultFunc([]string{"IC_VISIBILITY", "IBMCLOUD_VISIBILITY"}, "public"),
			},
			"endpoints_file_path": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Path of the file that contains private and public regional endpoints mapping",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"IC_ENDPOINTS_FILE_PATH", "IBMCLOUD_ENDPOINTS_FILE_PATH"}, nil),
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"ibm_api_gateway":        apigateway.DataSourceIBMApiGateway(),
			"ibm_account":            cloudfoundry.DataSourceIBMAccount(),
			"ibm_app":                cloudfoundry.DataSourceIBMApp(),
			"ibm_app_domain_private": cloudfoundry.DataSourceIBMAppDomainPrivate(),
			"ibm_app_domain_shared":  cloudfoundry.DataSourceIBMAppDomainShared(),
			"ibm_app_route":          cloudfoundry.DataSourceIBMAppRoute(),

			// // AppID
			"ibm_appid_action_url":               appid.DataSourceIBMAppIDActionURL(),
			"ibm_appid_apm":                      appid.DataSourceIBMAppIDAPM(),
			"ibm_appid_application":              appid.DataSourceIBMAppIDApplication(),
			"ibm_appid_application_scopes":       appid.DataSourceIBMAppIDApplicationScopes(),
			"ibm_appid_application_roles":        appid.DataSourceIBMAppIDApplicationRoles(),
			"ibm_appid_applications":             appid.DataSourceIBMAppIDApplications(),
			"ibm_appid_audit_status":             appid.DataSourceIBMAppIDAuditStatus(),
			"ibm_appid_cloud_directory_template": appid.DataSourceIBMAppIDCloudDirectoryTemplate(),
			"ibm_appid_cloud_directory_user":     appid.DataSourceIBMAppIDCloudDirectoryUser(),
			"ibm_appid_idp_cloud_directory":      appid.DataSourceIBMAppIDIDPCloudDirectory(),
			"ibm_appid_idp_custom":               appid.DataSourceIBMAppIDIDPCustom(),
			"ibm_appid_idp_facebook":             appid.DataSourceIBMAppIDIDPFacebook(),
			"ibm_appid_idp_google":               appid.DataSourceIBMAppIDIDPGoogle(),
			"ibm_appid_idp_saml":                 appid.DataSourceIBMAppIDIDPSAML(),
			"ibm_appid_idp_saml_metadata":        appid.DataSourceIBMAppIDIDPSAMLMetadata(),
			"ibm_appid_languages":                appid.DataSourceIBMAppIDLanguages(),
			"ibm_appid_mfa":                      appid.DataSourceIBMAppIDMFA(),
			"ibm_appid_mfa_channel":              appid.DataSourceIBMAppIDMFAChannel(),
			"ibm_appid_password_regex":           appid.DataSourceIBMAppIDPasswordRegex(),
			"ibm_appid_token_config":             appid.DataSourceIBMAppIDTokenConfig(),
			"ibm_appid_redirect_urls":            appid.DataSourceIBMAppIDRedirectURLs(),
			"ibm_appid_role":                     appid.DataSourceIBMAppIDRole(),
			"ibm_appid_roles":                    appid.DataSourceIBMAppIDRoles(),
			"ibm_appid_theme_color":              appid.DataSourceIBMAppIDThemeColor(),
			"ibm_appid_theme_text":               appid.DataSourceIBMAppIDThemeText(),
			"ibm_appid_user_roles":               appid.DataSourceIBMAppIDUserRoles(),

			"ibm_function_action":                   functions.DataSourceIBMFunctionAction(),
			"ibm_function_package":                  functions.DataSourceIBMFunctionPackage(),
			"ibm_function_rule":                     functions.DataSourceIBMFunctionRule(),
			"ibm_function_trigger":                  functions.DataSourceIBMFunctionTrigger(),
			"ibm_function_namespace":                functions.DataSourceIBMFunctionNamespace(),
			"ibm_cis":                               cis.DataSourceIBMCISInstance(),
			"ibm_cis_dns_records":                   cis.DataSourceIBMCISDNSRecords(),
			"ibm_cis_certificates":                  cis.DataSourceIBMCISCertificates(),
			"ibm_cis_global_load_balancers":         cis.DataSourceIBMCISGlbs(),
			"ibm_cis_origin_pools":                  cis.DataSourceIBMCISOriginPools(),
			"ibm_cis_healthchecks":                  cis.DataSourceIBMCISHealthChecks(),
			"ibm_cis_domain":                        cis.DataSourceIBMCISDomain(),
			"ibm_cis_firewall":                      cis.DataSourceIBMCISFirewallsRecord(),
			"ibm_cis_cache_settings":                cis.DataSourceIBMCISCacheSetting(),
			"ibm_cis_waf_packages":                  cis.DataSourceIBMCISWAFPackages(),
			"ibm_cis_range_apps":                    cis.DataSourceIBMCISRangeApps(),
			"ibm_cis_custom_certificates":           cis.DataSourceIBMCISCustomCertificates(),
			"ibm_cis_rate_limit":                    cis.DataSourceIBMCISRateLimit(),
			"ibm_cis_ip_addresses":                  cis.DataSourceIBMCISIP(),
			"ibm_cis_waf_groups":                    cis.DataSourceIBMCISWAFGroups(),
			"ibm_cis_alerts":                        cis.DataSourceIBMCISAlert(),
			"ibm_cis_origin_auths":                  cis.DataSourceIBMCISOriginAuthPull(),
			"ibm_cis_mtlss":                         cis.DataSourceIBMCISMtls(),
			"ibm_cis_mtls_apps":                     cis.DataSourceIBMCISMtlsApp(),
			"ibm_cis_bot_managements":               cis.DataSourceIBMCISBotManagement(),
			"ibm_cis_bot_analytics":                 cis.DataSourceIBMCISBotAnalytics(),
			"ibm_cis_webhooks":                      cis.DataSourceIBMCISWebhooks(),
			"ibm_cis_logpush_jobs":                  cis.DataSourceIBMCISLogPushJobs(),
			"ibm_cis_edge_functions_actions":        cis.DataSourceIBMCISEdgeFunctionsActions(),
			"ibm_cis_edge_functions_triggers":       cis.DataSourceIBMCISEdgeFunctionsTriggers(),
			"ibm_cis_custom_pages":                  cis.DataSourceIBMCISCustomPages(),
			"ibm_cis_page_rules":                    cis.DataSourceIBMCISPageRules(),
			"ibm_cis_waf_rules":                     cis.DataSourceIBMCISWAFRules(),
			"ibm_cis_filters":                       cis.DataSourceIBMCISFilters(),
			"ibm_cis_firewall_rules":                cis.DataSourceIBMCISFirewallRules(),
			"ibm_cloudant":                          cloudant.DataSourceIBMCloudant(),
			"ibm_cloudant_database":                 cloudant.DataSourceIBMCloudantDatabase(),
			"ibm_database":                          database.DataSourceIBMDatabaseInstance(),
			"ibm_database_connection":               database.DataSourceIBMDatabaseConnection(),
			"ibm_database_point_in_time_recovery":   database.DataSourceIBMDatabasePointInTimeRecovery(),
			"ibm_database_remotes":                  database.DataSourceIBMDatabaseRemotes(),
			"ibm_database_task":                     database.DataSourceIBMDatabaseTask(),
			"ibm_database_tasks":                    database.DataSourceIBMDatabaseTasks(),
			"ibm_database_backup":                   database.DataSourceIBMDatabaseBackup(),
			"ibm_database_backups":                  database.DataSourceIBMDatabaseBackups(),
			"ibm_compute_bare_metal":                classicinfrastructure.DataSourceIBMComputeBareMetal(),
			"ibm_compute_image_template":            classicinfrastructure.DataSourceIBMComputeImageTemplate(),
			"ibm_compute_placement_group":           classicinfrastructure.DataSourceIBMComputePlacementGroup(),
			"ibm_compute_reserved_capacity":         classicinfrastructure.DataSourceIBMComputeReservedCapacity(),
			"ibm_compute_ssh_key":                   classicinfrastructure.DataSourceIBMComputeSSHKey(),
			"ibm_compute_vm_instance":               classicinfrastructure.DataSourceIBMComputeVmInstance(),
			"ibm_container_addons":                  kubernetes.DataSourceIBMContainerAddOns(),
			"ibm_container_alb":                     kubernetes.DataSourceIBMContainerALB(),
			"ibm_container_alb_cert":                kubernetes.DataSourceIBMContainerALBCert(),
			"ibm_container_ingress_instance":        kubernetes.DataSourceIBMContainerIngressInstance(),
			"ibm_container_ingress_secret_tls":      kubernetes.DataSourceIBMContainerIngressSecretTLS(),
			"ibm_container_ingress_secret_opaque":   kubernetes.DataSourceIBMContainerIngressSecretOpaque(),
			"ibm_container_bind_service":            kubernetes.DataSourceIBMContainerBindService(),
			"ibm_container_cluster":                 kubernetes.DataSourceIBMContainerCluster(),
			"ibm_container_cluster_config":          kubernetes.DataSourceIBMContainerClusterConfig(),
			"ibm_container_cluster_versions":        kubernetes.DataSourceIBMContainerClusterVersions(),
			"ibm_container_cluster_worker":          kubernetes.DataSourceIBMContainerClusterWorker(),
			"ibm_container_nlb_dns":                 kubernetes.DataSourceIBMContainerNLBDNS(),
			"ibm_container_vpc_cluster_alb":         kubernetes.DataSourceIBMContainerVPCClusterALB(),
			"ibm_container_vpc_alb":                 kubernetes.DataSourceIBMContainerVPCClusterALB(),
			"ibm_container_vpc_cluster":             kubernetes.DataSourceIBMContainerVPCCluster(),
			"ibm_container_vpc_cluster_worker":      kubernetes.DataSourceIBMContainerVPCClusterWorker(),
			"ibm_container_vpc_cluster_worker_pool": kubernetes.DataSourceIBMContainerVpcClusterWorkerPool(),
			"ibm_container_vpc_worker_pool":         kubernetes.DataSourceIBMContainerVpcClusterWorkerPool(),
			"ibm_container_worker_pool":             kubernetes.DataSourceIBMContainerWorkerPool(),
			"ibm_container_storage_attachment":      kubernetes.DataSourceIBMContainerVpcWorkerVolumeAttachment(),
			"ibm_container_dedicated_host_pool":     kubernetes.DataSourceIBMContainerDedicatedHostPool(),
			"ibm_container_dedicated_host_flavor":   kubernetes.DataSourceIBMContainerDedicatedHostFlavor(),
			"ibm_container_dedicated_host_flavors":  kubernetes.DataSourceIBMContainerDedicatedHostFlavors(),
			"ibm_container_dedicated_host":          kubernetes.DataSourceIBMContainerDedicatedHost(),
			"ibm_cr_namespaces":                     registry.DataIBMContainerRegistryNamespaces(),
			"ibm_cloud_shell_account_settings":      cloudshell.DataSourceIBMCloudShellAccountSettings(),
			"ibm_cos_bucket":                        cos.DataSourceIBMCosBucket(),
			"ibm_cos_bucket_object":                 cos.DataSourceIBMCosBucketObject(),
			"ibm_dns_domain_registration":           classicinfrastructure.DataSourceIBMDNSDomainRegistration(),
			"ibm_dns_domain":                        classicinfrastructure.DataSourceIBMDNSDomain(),
			"ibm_dns_secondary":                     classicinfrastructure.DataSourceIBMDNSSecondary(),
			"ibm_event_streams_topic":               eventstreams.DataSourceIBMEventStreamsTopic(),
			"ibm_event_streams_schema":              eventstreams.DataSourceIBMEventStreamsSchema(),
			"ibm_hpcs":                              hpcs.DataSourceIBMHPCS(),
			"ibm_hpcs_managed_key":                  hpcs.DataSourceIbmManagedKey(),
			"ibm_hpcs_key_template":                 hpcs.DataSourceIbmKeyTemplate(),
			"ibm_hpcs_keystore":                     hpcs.DataSourceIbmKeystore(),
			"ibm_hpcs_vault":                        hpcs.DataSourceIbmVault(),
			"ibm_iam_access_group":                  iamaccessgroup.DataSourceIBMIAMAccessGroup(),
			"ibm_iam_access_group_policy":           iampolicy.DataSourceIBMIAMAccessGroupPolicy(),
			"ibm_iam_account_settings":              iamidentity.DataSourceIBMIAMAccountSettings(),
			"ibm_iam_auth_token":                    iamidentity.DataSourceIBMIAMAuthToken(),
			"ibm_iam_role_actions":                  iampolicy.DataSourceIBMIAMRoleAction(),
			"ibm_iam_users":                         iamidentity.DataSourceIBMIAMUsers(),
			"ibm_iam_roles":                         iampolicy.DataSourceIBMIAMRole(),
			"ibm_iam_user_policy":                   iampolicy.DataSourceIBMIAMUserPolicy(),
			"ibm_iam_authorization_policies":        iampolicy.DataSourceIBMIAMAuthorizationPolicies(),
			"ibm_iam_user_profile":                  iamidentity.DataSourceIBMIAMUserProfile(),
			"ibm_iam_service_id":                    iamidentity.DataSourceIBMIAMServiceID(),
			"ibm_iam_service_policy":                iampolicy.DataSourceIBMIAMServicePolicy(),
			"ibm_iam_api_key":                       iamidentity.DataSourceIBMIamApiKey(),
			"ibm_iam_trusted_profile":               iamidentity.DataSourceIBMIamTrustedProfile(),
			"ibm_iam_trusted_profile_claim_rule":    iamidentity.DataSourceIBMIamTrustedProfileClaimRule(),
			"ibm_iam_trusted_profile_link":          iamidentity.DataSourceIBMIamTrustedProfileLink(),
			"ibm_iam_trusted_profile_claim_rules":   iamidentity.DataSourceIBMIamTrustedProfileClaimRules(),
			"ibm_iam_trusted_profile_links":         iamidentity.DataSourceIBMIamTrustedProfileLinks(),
			"ibm_iam_trusted_profiles":              iamidentity.DataSourceIBMIamTrustedProfiles(),
			"ibm_iam_trusted_profile_policy":        iampolicy.DataSourceIBMIAMTrustedProfilePolicy(),
			"ibm_iam_user_mfa_enrollments":          iamidentity.DataSourceIBMIamUserMfaEnrollments(),

			"ibm_lbaas":                   classicinfrastructure.DataSourceIBMLbaas(),
			"ibm_network_vlan":            classicinfrastructure.DataSourceIBMNetworkVlan(),
			"ibm_org":                     cloudfoundry.DataSourceIBMOrg(),
			"ibm_org_quota":               cloudfoundry.DataSourceIBMOrgQuota(),
			"ibm_kms_instance_policies":   kms.DataSourceIBMKmsInstancePolicies(),
			"ibm_kp_key":                  kms.DataSourceIBMkey(),
			"ibm_kms_key_rings":           kms.DataSourceIBMKMSkeyRings(),
			"ibm_kms_key_policies":        kms.DataSourceIBMKMSkeyPolicies(),
			"ibm_kms_keys":                kms.DataSourceIBMKMSkeys(),
			"ibm_kms_key":                 kms.DataSourceIBMKMSkey(),
			"ibm_pn_application_chrome":   pushnotification.DataSourceIBMPNApplicationChrome(),
			"ibm_app_config_environment":  appconfiguration.DataSourceIBMAppConfigEnvironment(),
			"ibm_app_config_environments": appconfiguration.DataSourceIBMAppConfigEnvironments(),
			"ibm_app_config_collection":   appconfiguration.DataSourceIBMAppConfigCollection(),
			"ibm_app_config_collections":  appconfiguration.DataSourceIBMAppConfigCollections(),
			"ibm_app_config_feature":      appconfiguration.DataSourceIBMAppConfigFeature(),
			"ibm_app_config_features":     appconfiguration.DataSourceIBMAppConfigFeatures(),
			"ibm_app_config_property":     appconfiguration.DataSourceIBMAppConfigProperty(),
			"ibm_app_config_properties":   appconfiguration.DataSourceIBMAppConfigProperties(),
			"ibm_app_config_segment":      appconfiguration.DataSourceIBMAppConfigSegment(),
			"ibm_app_config_segments":     appconfiguration.DataSourceIBMAppConfigSegments(),
			"ibm_app_config_snapshot":     appconfiguration.DataSourceIBMAppConfigSnapshot(),
			"ibm_app_config_snapshots":    appconfiguration.DataSourceIBMAppConfigSnapshots(),

			"ibm_resource_quota":    resourcecontroller.DataSourceIBMResourceQuota(),
			"ibm_resource_group":    resourcemanager.DataSourceIBMResourceGroup(),
			"ibm_resource_instance": resourcecontroller.DataSourceIBMResourceInstance(),
			"ibm_resource_key":      resourcecontroller.DataSourceIBMResourceKey(),
			"ibm_security_group":    classicinfrastructure.DataSourceIBMSecurityGroup(),
			"ibm_service_instance":  cloudfoundry.DataSourceIBMServiceInstance(),
			"ibm_service_key":       cloudfoundry.DataSourceIBMServiceKey(),
			"ibm_service_plan":      cloudfoundry.DataSourceIBMServicePlan(),
			"ibm_space":             cloudfoundry.DataSourceIBMSpace(),

			// Added for Schematics
			"ibm_schematics_workspace":      schematics.DataSourceIBMSchematicsWorkspace(),
			"ibm_schematics_output":         schematics.DataSourceIBMSchematicsOutput(),
			"ibm_schematics_state":          schematics.DataSourceIBMSchematicsState(),
			"ibm_schematics_action":         schematics.DataSourceIBMSchematicsAction(),
			"ibm_schematics_job":            schematics.DataSourceIBMSchematicsJob(),
			"ibm_schematics_inventory":      schematics.DataSourceIBMSchematicsInventory(),
			"ibm_schematics_resource_query": schematics.DataSourceIBMSchematicsResourceQuery(),

			// // Added for Power Resources

			"ibm_pi_catalog_images":                         power.DataSourceIBMPICatalogImages(),
			"ibm_pi_cloud_connection":                       power.DataSourceIBMPICloudConnection(),
			"ibm_pi_cloud_connections":                      power.DataSourceIBMPICloudConnections(),
			"ibm_pi_cloud_instance":                         power.DataSourceIBMPICloudInstance(),
			"ibm_pi_console_languages":                      power.DataSourceIBMPIInstanceConsoleLanguages(),
			"ibm_pi_dhcp":                                   power.DataSourceIBMPIDhcp(),
			"ibm_pi_dhcps":                                  power.DataSourceIBMPIDhcps(),
			"ibm_pi_disaster_recovery_location":             power.DataSourceIBMPIDisasterRecoveryLocation(),
			"ibm_pi_disaster_recovery_locations":            power.DataSourceIBMPIDisasterRecoveryLocations(),
			"ibm_pi_image":                                  power.DataSourceIBMPIImage(),
			"ibm_pi_images":                                 power.DataSourceIBMPIImages(),
			"ibm_pi_instance":                               power.DataSourceIBMPIInstance(),
			"ibm_pi_instances":                              power.DataSourceIBMPIInstances(),
			"ibm_pi_instance_ip":                            power.DataSourceIBMPIInstanceIP(),
			"ibm_pi_instance_snapshots":                     power.DataSourceIBMPISnapshots(),
			"ibm_pi_instance_volumes":                       power.DataSourceIBMPIInstanceVolumes(),
			"ibm_pi_key":                                    power.DataSourceIBMPIKey(),
			"ibm_pi_keys":                                   power.DataSourceIBMPIKeys(),
			"ibm_pi_network":                                power.DataSourceIBMPINetwork(),
			"ibm_pi_network_port":                           power.DataSourceIBMPINetworkPort(),
			"ibm_pi_placement_group":                        power.DataSourceIBMPIPlacementGroup(),
			"ibm_pi_placement_groups":                       power.DataSourceIBMPIPlacementGroups(),
			"ibm_pi_public_network":                         power.DataSourceIBMPIPublicNetwork(),
			"ibm_pi_pvm_snapshots":                          power.DataSourceIBMPISnapshot(),
			"ibm_pi_sap_profile":                            power.DataSourceIBMPISAPProfile(),
			"ibm_pi_sap_profiles":                           power.DataSourceIBMPISAPProfiles(),
			"ibm_pi_shared_processor_pool":                  power.DataSourceIBMPISharedProcessorPool(),
			"ibm_pi_shared_processor_pools":                 power.DataSourceIBMPISharedProcessorPools(),
			"ibm_pi_spp_placement_group":                    power.DataSourceIBMPISPPPlacementGroup(),
			"ibm_pi_spp_placement_groups":                   power.DataSourceIBMPISPPPlacementGroups(),
			"ibm_pi_storage_pool_capacity":                  power.DataSourceIBMPIStoragePoolCapacity(),
			"ibm_pi_storage_pools_capacity":                 power.DataSourceIBMPIStoragePoolsCapacity(),
			"ibm_pi_storage_type_capacity":                  power.DataSourceIBMPIStorageTypeCapacity(),
			"ibm_pi_storage_types_capacity":                 power.DataSourceIBMPIStorageTypesCapacity(),
			"ibm_pi_system_pools":                           power.DataSourceIBMPISystemPools(),
			"ibm_pi_tenant":                                 power.DataSourceIBMPITenant(),
			"ibm_pi_volume":                                 power.DataSourceIBMPIVolume(),
			"ibm_pi_volume_group":                           power.DataSourceIBMPIVolumeGroup(),
			"ibm_pi_volume_groups":                          power.DataSourceIBMPIVolumeGroups(),
			"ibm_pi_volume_group_details":                   power.DataSourceIBMPIVolumeGroupDetails(),
			"ibm_pi_volume_groups_details":                  power.DataSourceIBMPIVolumeGroupsDetails(),
			"ibm_pi_volume_group_storage_details":           power.DataSourceIBMPIVolumeGroupStorageDetails(),
			"ibm_pi_volume_group_remote_copy_relationships": power.DataSourceIBMPIVolumeGroupRemoteCopyRelationships(),
			"ibm_pi_volume_flash_copy_mappings":             power.DataSourceIBMPIVolumeFlashCopyMappings(),
			"ibm_pi_volume_remote_copy_relationship":        power.DataSourceIBMPIVolumeRemoteCopyRelationship(),
			"ibm_pi_volume_onboardings":                     power.DataSourceIBMPIVolumeOnboardings(),
			"ibm_pi_volume_onboarding":                      power.DataSourceIBMPIVolumeOnboarding(),

			// // Added for private dns zones

			"ibm_dns_zones":                            dnsservices.DataSourceIBMPrivateDNSZones(),
			"ibm_dns_permitted_networks":               dnsservices.DataSourceIBMPrivateDNSPermittedNetworks(),
			"ibm_dns_resource_records":                 dnsservices.DataSourceIBMPrivateDNSResourceRecords(),
			"ibm_dns_glb_monitors":                     dnsservices.DataSourceIBMPrivateDNSGLBMonitors(),
			"ibm_dns_glb_pools":                        dnsservices.DataSourceIBMPrivateDNSGLBPools(),
			"ibm_dns_glbs":                             dnsservices.DataSourceIBMPrivateDNSGLBs(),
			"ibm_dns_custom_resolvers":                 dnsservices.DataSourceIBMPrivateDNSCustomResolver(),
			"ibm_dns_custom_resolver_forwarding_rules": dnsservices.DataSourceIBMPrivateDNSForwardingRules(),
			"ibm_dns_custom_resolver_secondary_zones":  dnsservices.DataSourceIBMPrivateDNSSecondaryZones(),

			// // Added for Direct Link

			"ibm_dl_gateways":             directlink.DataSourceIBMDLGateways(),
			"ibm_dl_offering_speeds":      directlink.DataSourceIBMDLOfferingSpeeds(),
			"ibm_dl_port":                 directlink.DataSourceIBMDirectLinkPort(),
			"ibm_dl_ports":                directlink.DataSourceIBMDirectLinkPorts(),
			"ibm_dl_gateway":              directlink.DataSourceIBMDLGateway(),
			"ibm_dl_locations":            directlink.DataSourceIBMDLLocations(),
			"ibm_dl_routers":              directlink.DataSourceIBMDLRouters(),
			"ibm_dl_provider_ports":       directlink.DataSourceIBMDirectLinkProviderPorts(),
			"ibm_dl_provider_gateways":    directlink.DataSourceIBMDirectLinkProviderGateways(),
			"ibm_dl_route_reports":        directlink.DataSourceIBMDLRouteReports(),
			"ibm_dl_route_report":         directlink.DataSourceIBMDLRouteReport(),
			"ibm_dl_export_route_filters": directlink.DataSourceIBMDLExportRouteFilters(),
			"ibm_dl_export_route_filter":  directlink.DataSourceIBMDLExportRouteFilter(),
			"ibm_dl_import_route_filters": directlink.DataSourceIBMDLImportRouteFilters(),
			"ibm_dl_import_route_filter":  directlink.DataSourceIBMDLImportRouteFilter(),

			// //Added for Transit Gateway
			"ibm_tg_gateway":                   transitgateway.DataSourceIBMTransitGateway(),
			"ibm_tg_gateways":                  transitgateway.DataSourceIBMTransitGateways(),
			"ibm_tg_connection_prefix_filter":  transitgateway.DataSourceIBMTransitGatewayConnectionPrefixFilter(),
			"ibm_tg_connection_prefix_filters": transitgateway.DataSourceIBMTransitGatewayConnectionPrefixFilters(),
			"ibm_tg_locations":                 transitgateway.DataSourceIBMTransitGatewaysLocations(),
			"ibm_tg_location":                  transitgateway.DataSourceIBMTransitGatewaysLocation(),
			"ibm_tg_route_report":              transitgateway.DataSourceIBMTransitGatewayRouteReport(),
			"ibm_tg_route_reports":             transitgateway.DataSourceIBMTransitGatewayRouteReports(),

			// //Added for BSS Enterprise
			"ibm_enterprises":               enterprise.DataSourceIBMEnterprises(),
			"ibm_enterprise_account_groups": enterprise.DataSourceIBMEnterpriseAccountGroups(),
			"ibm_enterprise_accounts":       enterprise.DataSourceIBMEnterpriseAccounts(),

			// //Added for Secrets Manager
			// V1 data sources:
			"ibm_secrets_manager_secrets": secretsmanager.DataSourceIBMSecretsManagerSecrets(),
			"ibm_secrets_manager_secret":  secretsmanager.DataSourceIBMSecretsManagerSecret(),

			// V2 data sources
			"ibm_sm_secret_group":  secretsmanager.AddInstanceFields(secretsmanager.DataSourceIbmSmSecretGroup()),
			"ibm_sm_secret_groups": secretsmanager.AddInstanceFields(secretsmanager.DataSourceIbmSmSecretGroups()),
			"ibm_sm_private_certificate_configuration_intermediate_ca":           secretsmanager.AddInstanceFields(secretsmanager.DataSourceIbmSmPrivateCertificateConfigurationIntermediateCA()),
			"ibm_sm_private_certificate_configuration_root_ca":                   secretsmanager.AddInstanceFields(secretsmanager.DataSourceIbmSmPrivateCertificateConfigurationRootCA()),
			"ibm_sm_private_certificate_configuration_template":                  secretsmanager.AddInstanceFields(secretsmanager.DataSourceIbmSmPrivateCertificateConfigurationTemplate()),
			"ibm_sm_public_certificate_configuration_ca_lets_encrypt":            secretsmanager.AddInstanceFields(secretsmanager.DataSourceIbmSmPublicCertificateConfigurationCALetsEncrypt()),
			"ibm_sm_public_certificate_configuration_dns_cis":                    secretsmanager.AddInstanceFields(secretsmanager.DataSourceIbmSmConfigurationPublicCertificateDNSCis()),
			"ibm_sm_public_certificate_configuration_dns_classic_infrastructure": secretsmanager.AddInstanceFields(secretsmanager.DataSourceIbmSmPublicCertificateConfigurationDNSClassicInfrastructure()),
			"ibm_sm_iam_credentials_configuration":                               secretsmanager.AddInstanceFields(secretsmanager.DataSourceIbmSmIamCredentialsConfiguration()),
			"ibm_sm_configurations":                                              secretsmanager.AddInstanceFields(secretsmanager.DataSourceIbmSmConfigurations()),
			"ibm_sm_secrets":                                                     secretsmanager.AddInstanceFields(secretsmanager.DataSourceIbmSmSecrets()),
			"ibm_sm_arbitrary_secret_metadata":                                   secretsmanager.AddInstanceFields(secretsmanager.DataSourceIbmSmArbitrarySecretMetadata()),
			"ibm_sm_imported_certificate_metadata":                               secretsmanager.AddInstanceFields(secretsmanager.DataSourceIbmSmImportedCertificateMetadata()),
			"ibm_sm_public_certificate_metadata":                                 secretsmanager.AddInstanceFields(secretsmanager.DataSourceIbmSmPublicCertificateMetadata()),
			"ibm_sm_private_certificate_metadata":                                secretsmanager.AddInstanceFields(secretsmanager.DataSourceIbmSmPrivateCertificateMetadata()),
			"ibm_sm_iam_credentials_secret_metadata":                             secretsmanager.AddInstanceFields(secretsmanager.DataSourceIbmSmIamCredentialsSecretMetadata()),
			"ibm_sm_kv_secret_metadata":                                          secretsmanager.AddInstanceFields(secretsmanager.DataSourceIbmSmKvSecretMetadata()),
			"ibm_sm_username_password_secret_metadata":                           secretsmanager.AddInstanceFields(secretsmanager.DataSourceIbmSmUsernamePasswordSecretMetadata()),
			"ibm_sm_arbitrary_secret":                                            secretsmanager.AddInstanceFields(secretsmanager.DataSourceIbmSmArbitrarySecret()),
			"ibm_sm_imported_certificate":                                        secretsmanager.AddInstanceFields(secretsmanager.DataSourceIbmSmImportedCertificate()),
			"ibm_sm_public_certificate":                                          secretsmanager.AddInstanceFields(secretsmanager.DataSourceIbmSmPublicCertificate()),
			"ibm_sm_private_certificate":                                         secretsmanager.AddInstanceFields(secretsmanager.DataSourceIbmSmPrivateCertificate()),
			"ibm_sm_iam_credentials_secret":                                      secretsmanager.AddInstanceFields(secretsmanager.DataSourceIbmSmIamCredentialsSecret()),
			"ibm_sm_kv_secret":                                                   secretsmanager.AddInstanceFields(secretsmanager.DataSourceIbmSmKvSecret()),
			"ibm_sm_username_password_secret":                                    secretsmanager.AddInstanceFields(secretsmanager.DataSourceIbmSmUsernamePasswordSecret()),
			"ibm_sm_en_registration":                                             secretsmanager.AddInstanceFields(secretsmanager.DataSourceIbmSmEnRegistration()),

			// //Added for Satellite
			"ibm_satellite_location":                            satellite.DataSourceIBMSatelliteLocation(),
			"ibm_satellite_location_nlb_dns":                    satellite.DataSourceIBMSatelliteLocationNLBDNS(),
			"ibm_satellite_attach_host_script":                  satellite.DataSourceIBMSatelliteAttachHostScript(),
			"ibm_satellite_cluster":                             satellite.DataSourceIBMSatelliteCluster(),
			"ibm_satellite_cluster_worker_pool":                 satellite.DataSourceIBMSatelliteClusterWorkerPool(),
			"ibm_satellite_link":                                satellite.DataSourceIBMSatelliteLink(),
			"ibm_satellite_endpoint":                            satellite.DataSourceIBMSatelliteEndpoint(),
			"ibm_satellite_cluster_worker_pool_zone_attachment": satellite.DataSourceIBMSatelliteClusterWorkerPoolAttachment(),

			// // Catalog related resources
			"ibm_cm_catalog":           catalogmanagement.DataSourceIBMCmCatalog(),
			"ibm_cm_offering":          catalogmanagement.DataSourceIBMCmOffering(),
			"ibm_cm_version":           catalogmanagement.DataSourceIBMCmVersion(),
			"ibm_cm_offering_instance": catalogmanagement.DataSourceIBMCmOfferingInstance(),
			"ibm_cm_preset":            catalogmanagement.DataSourceIBMCmPreset(),
			"ibm_cm_object":            catalogmanagement.DataSourceIBMCmObject(),

			// //Added for Resource Tag
			"ibm_resource_tag": globaltagging.DataSourceIBMResourceTag(),

			// // Atracker
			"ibm_atracker_targets":   atracker.DataSourceIBMAtrackerTargets(),
			"ibm_atracker_routes":    atracker.DataSourceIBMAtrackerRoutes(),
			"ibm_atracker_endpoints": atracker.DataSourceIBMAtrackerEndpoints(),

			//  Metrics Router
			"ibm_metrics_router_targets": metricsrouter.DataSourceIBMMetricsRouterTargets(),
			"ibm_metrics_router_routes":  metricsrouter.DataSourceIBMMetricsRouterRoutes(),

			//Security and Compliance Center
			"ibm_scc_account_location":              scc.DataSourceIBMSccAccountLocation(),
			"ibm_scc_account_locations":             scc.DataSourceIBMSccAccountLocations(),
			"ibm_scc_account_location_settings":     scc.DataSourceIBMSccAccountLocationSettings(),
			"ibm_scc_account_notification_settings": scc.DataSourceIBMSccNotificationSettings(),

			// // Added for Context Based Restrictions
			"ibm_cbr_zone": contextbasedrestrictions.DataSourceIBMCbrZone(),
			"ibm_cbr_rule": contextbasedrestrictions.DataSourceIBMCbrRule(),

			// // Added for Event Notifications
			"ibm_en_source":                 eventnotification.DataSourceIBMEnSource(),
			"ibm_en_destinations":           eventnotification.DataSourceIBMEnDestinations(),
			"ibm_en_topic":                  eventnotification.DataSourceIBMEnTopic(),
			"ibm_en_topics":                 eventnotification.DataSourceIBMEnTopics(),
			"ibm_en_subscriptions":          eventnotification.DataSourceIBMEnSubscriptions(),
			"ibm_en_destination_webhook":    eventnotification.DataSourceIBMEnWebhookDestination(),
			"ibm_en_destination_android":    eventnotification.DataSourceIBMEnFCMDestination(),
			"ibm_en_destination_ios":        eventnotification.DataSourceIBMEnAPNSDestination(),
			"ibm_en_destination_chrome":     eventnotification.DataSourceIBMEnChromeDestination(),
			"ibm_en_destination_firefox":    eventnotification.DataSourceIBMEnFirefoxDestination(),
			"ibm_en_destination_slack":      eventnotification.DataSourceIBMEnSlackDestination(),
			"ibm_en_subscription_sms":       eventnotification.DataSourceIBMEnSMSSubscription(),
			"ibm_en_subscription_email":     eventnotification.DataSourceIBMEnEmailSubscription(),
			"ibm_en_subscription_webhook":   eventnotification.DataSourceIBMEnWebhookSubscription(),
			"ibm_en_subscription_android":   eventnotification.DataSourceIBMEnFCMSubscription(),
			"ibm_en_subscription_ios":       eventnotification.DataSourceIBMEnFCMSubscription(),
			"ibm_en_subscription_chrome":    eventnotification.DataSourceIBMEnFCMSubscription(),
			"ibm_en_subscription_firefox":   eventnotification.DataSourceIBMEnFCMSubscription(),
			"ibm_en_subscription_slack":     eventnotification.DataSourceIBMEnSlackSubscription(),
			"ibm_en_subscription_safari":    eventnotification.DataSourceIBMEnFCMSubscription(),
			"ibm_en_destination_safari":     eventnotification.DataSourceIBMEnSafariDestination(),
			"ibm_en_destination_msteams":    eventnotification.DataSourceIBMEnMSTeamsDestination(),
			"ibm_en_subscription_msteams":   eventnotification.DataSourceIBMEnFCMSubscription(),
			"ibm_en_destination_cf":         eventnotification.DataSourceIBMEnCFDestination(),
			"ibm_en_subscription_cf":        eventnotification.DataSourceIBMEnFCMSubscription(),
			"ibm_en_destination_pagerduty":  eventnotification.DataSourceIBMEnPagerDutyDestination(),
			"ibm_en_subscription_pagerduty": eventnotification.DataSourceIBMEnFCMSubscription(),
			"ibm_en_integration":            eventnotification.DataSourceIBMEnIntegration(),
			"ibm_en_integrations":           eventnotification.DataSourceIBMEnIntegrations(),
			"ibm_en_destination_sn":         eventnotification.DataSourceIBMEnServiceNowDestination(),
			"ibm_en_subscription_sn":        eventnotification.DataSourceIBMEnFCMSubscription(),
			"ibm_en_destination_ce":         eventnotification.DataSourceIBMEnCodeEngineDestination(),
			"ibm_en_subscription_ce":        eventnotification.DataSourceIBMEnFCMSubscription(),
			"ibm_en_destination_cos":        eventnotification.DataSourceIBMEnCOSDestination(),
			"ibm_en_subscription_cos":       eventnotification.DataSourceIBMEnFCMSubscription(),

			// // Added for Toolchain
			"ibm_cd_toolchain":                         cdtoolchain.DataSourceIBMCdToolchain(),
			"ibm_cd_toolchain_tool_keyprotect":         cdtoolchain.DataSourceIBMCdToolchainToolKeyprotect(),
			"ibm_cd_toolchain_tool_secretsmanager":     cdtoolchain.DataSourceIBMCdToolchainToolSecretsmanager(),
			"ibm_cd_toolchain_tool_bitbucketgit":       cdtoolchain.DataSourceIBMCdToolchainToolBitbucketgit(),
			"ibm_cd_toolchain_tool_githubconsolidated": cdtoolchain.DataSourceIBMCdToolchainToolGithubconsolidated(),
			"ibm_cd_toolchain_tool_gitlab":             cdtoolchain.DataSourceIBMCdToolchainToolGitlab(),
			"ibm_cd_toolchain_tool_hostedgit":          cdtoolchain.DataSourceIBMCdToolchainToolHostedgit(),
			"ibm_cd_toolchain_tool_artifactory":        cdtoolchain.DataSourceIBMCdToolchainToolArtifactory(),
			"ibm_cd_toolchain_tool_custom":             cdtoolchain.DataSourceIBMCdToolchainToolCustom(),
			"ibm_cd_toolchain_tool_pipeline":           cdtoolchain.DataSourceIBMCdToolchainToolPipeline(),
			"ibm_cd_toolchain_tool_devopsinsights":     cdtoolchain.DataSourceIBMCdToolchainToolDevopsinsights(),
			"ibm_cd_toolchain_tool_slack":              cdtoolchain.DataSourceIBMCdToolchainToolSlack(),
			"ibm_cd_toolchain_tool_sonarqube":          cdtoolchain.DataSourceIBMCdToolchainToolSonarqube(),
			"ibm_cd_toolchain_tool_hashicorpvault":     cdtoolchain.DataSourceIBMCdToolchainToolHashicorpvault(),
			"ibm_cd_toolchain_tool_securitycompliance": cdtoolchain.DataSourceIBMCdToolchainToolSecuritycompliance(),
			"ibm_cd_toolchain_tool_privateworker":      cdtoolchain.DataSourceIBMCdToolchainToolPrivateworker(),
			"ibm_cd_toolchain_tool_appconfig":          cdtoolchain.DataSourceIBMCdToolchainToolAppconfig(),
			"ibm_cd_toolchain_tool_jenkins":            cdtoolchain.DataSourceIBMCdToolchainToolJenkins(),
			"ibm_cd_toolchain_tool_nexus":              cdtoolchain.DataSourceIBMCdToolchainToolNexus(),
			"ibm_cd_toolchain_tool_pagerduty":          cdtoolchain.DataSourceIBMCdToolchainToolPagerduty(),
			"ibm_cd_toolchain_tool_saucelabs":          cdtoolchain.DataSourceIBMCdToolchainToolSaucelabs(),
			"ibm_cd_toolchain_tool_jira":               cdtoolchain.DataSourceIBMCdToolchainToolJira(),
			"ibm_cd_toolchain_tool_eventnotifications": cdtoolchain.DataSourceIBMCdToolchainToolEventnotifications(),

			// Added for Tekton Pipeline
			"ibm_cd_tekton_pipeline_definition":       cdtektonpipeline.DataSourceIBMCdTektonPipelineDefinition(),
			"ibm_cd_tekton_pipeline_trigger_property": cdtektonpipeline.DataSourceIBMCdTektonPipelineTriggerProperty(),
			"ibm_cd_tekton_pipeline_property":         cdtektonpipeline.DataSourceIBMCdTektonPipelineProperty(),
			"ibm_cd_tekton_pipeline_trigger":          cdtektonpipeline.DataSourceIBMCdTektonPipelineTrigger(),
			"ibm_cd_tekton_pipeline":                  cdtektonpipeline.DataSourceIBMCdTektonPipeline(),

			// Added for Code Engine
			"ibm_code_engine_app":        codeengine.DataSourceIbmCodeEngineApp(),
			"ibm_code_engine_binding":    codeengine.DataSourceIbmCodeEngineBinding(),
			"ibm_code_engine_build":      codeengine.DataSourceIbmCodeEngineBuild(),
			"ibm_code_engine_config_map": codeengine.DataSourceIbmCodeEngineConfigMap(),
			"ibm_code_engine_job":        codeengine.DataSourceIbmCodeEngineJob(),
			"ibm_code_engine_project":    codeengine.DataSourceIbmCodeEngineProject(),
			"ibm_code_engine_secret":     codeengine.DataSourceIbmCodeEngineSecret(),

			// Added for Project
			"ibm_project":        project.DataSourceIbmProject(),
			"ibm_project_config": project.DataSourceIbmProjectConfig(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"ibm_api_gateway_endpoint":              apigateway.ResourceIBMApiGatewayEndPoint(),
			"ibm_api_gateway_endpoint_subscription": apigateway.ResourceIBMApiGatewayEndpointSubscription(),
			"ibm_app":                               cloudfoundry.ResourceIBMApp(),
			"ibm_app_domain_private":                cloudfoundry.ResourceIBMAppDomainPrivate(),
			"ibm_app_domain_shared":                 cloudfoundry.ResourceIBMAppDomainShared(),
			"ibm_app_route":                         cloudfoundry.ResourceIBMAppRoute(),

			// // AppID
			"ibm_appid_action_url":               appid.ResourceIBMAppIDActionURL(),
			"ibm_appid_apm":                      appid.ResourceIBMAppIDAPM(),
			"ibm_appid_application":              appid.ResourceIBMAppIDApplication(),
			"ibm_appid_application_scopes":       appid.ResourceIBMAppIDApplicationScopes(),
			"ibm_appid_application_roles":        appid.ResourceIBMAppIDApplicationRoles(),
			"ibm_appid_audit_status":             appid.ResourceIBMAppIDAuditStatus(),
			"ibm_appid_cloud_directory_template": appid.ResourceIBMAppIDCloudDirectoryTemplate(),
			"ibm_appid_cloud_directory_user":     appid.ResourceIBMAppIDCloudDirectoryUser(),
			"ibm_appid_idp_cloud_directory":      appid.ResourceIBMAppIDIDPCloudDirectory(),
			"ibm_appid_idp_custom":               appid.ResourceIBMAppIDIDPCustom(),
			"ibm_appid_idp_facebook":             appid.ResourceIBMAppIDIDPFacebook(),
			"ibm_appid_idp_google":               appid.ResourceIBMAppIDIDPGoogle(),
			"ibm_appid_idp_saml":                 appid.ResourceIBMAppIDIDPSAML(),
			"ibm_appid_languages":                appid.ResourceIBMAppIDLanguages(),
			"ibm_appid_mfa":                      appid.ResourceIBMAppIDMFA(),
			"ibm_appid_mfa_channel":              appid.ResourceIBMAppIDMFAChannel(),
			"ibm_appid_password_regex":           appid.ResourceIBMAppIDPasswordRegex(),
			"ibm_appid_token_config":             appid.ResourceIBMAppIDTokenConfig(),
			"ibm_appid_redirect_urls":            appid.ResourceIBMAppIDRedirectURLs(),
			"ibm_appid_role":                     appid.ResourceIBMAppIDRole(),
			"ibm_appid_theme_color":              appid.ResourceIBMAppIDThemeColor(),
			"ibm_appid_theme_text":               appid.ResourceIBMAppIDThemeText(),
			"ibm_appid_user_roles":               appid.ResourceIBMAppIDUserRoles(),

			"ibm_function_action":                       functions.ResourceIBMFunctionAction(),
			"ibm_function_package":                      functions.ResourceIBMFunctionPackage(),
			"ibm_function_rule":                         functions.ResourceIBMFunctionRule(),
			"ibm_function_trigger":                      functions.ResourceIBMFunctionTrigger(),
			"ibm_function_namespace":                    functions.ResourceIBMFunctionNamespace(),
			"ibm_cis":                                   cis.ResourceIBMCISInstance(),
			"ibm_database":                              database.ResourceIBMDatabaseInstance(),
			"ibm_cis_domain":                            cis.ResourceIBMCISDomain(),
			"ibm_cis_domain_settings":                   cis.ResourceIBMCISSettings(),
			"ibm_cis_firewall":                          cis.ResourceIBMCISFirewallRecord(),
			"ibm_cis_range_app":                         cis.ResourceIBMCISRangeApp(),
			"ibm_cis_healthcheck":                       cis.ResourceIBMCISHealthCheck(),
			"ibm_cis_origin_pool":                       cis.ResourceIBMCISPool(),
			"ibm_cis_global_load_balancer":              cis.ResourceIBMCISGlb(),
			"ibm_cis_certificate_upload":                cis.ResourceIBMCISCertificateUpload(),
			"ibm_cis_dns_record":                        cis.ResourceIBMCISDnsRecord(),
			"ibm_cis_dns_records_import":                cis.ResourceIBMCISDNSRecordsImport(),
			"ibm_cis_rate_limit":                        cis.ResourceIBMCISRateLimit(),
			"ibm_cis_page_rule":                         cis.ResourceIBMCISPageRule(),
			"ibm_cis_edge_functions_action":             cis.ResourceIBMCISEdgeFunctionsAction(),
			"ibm_cis_edge_functions_trigger":            cis.ResourceIBMCISEdgeFunctionsTrigger(),
			"ibm_cis_tls_settings":                      cis.ResourceIBMCISTLSSettings(),
			"ibm_cis_waf_package":                       cis.ResourceIBMCISWAFPackage(),
			"ibm_cis_webhook":                           cis.ResourceIBMCISWebhooks(),
			"ibm_cis_origin_auth":                       cis.ResourceIBMCISOriginAuthPull(),
			"ibm_cis_mtls":                              cis.ResourceIBMCISMtls(),
			"ibm_cis_mtls_app":                          cis.ResourceIBMCISMtlsApp(),
			"ibm_cis_bot_management":                    cis.ResourceIBMCISBotManagement(),
			"ibm_cis_logpush_job":                       cis.ResourceIBMCISLogPushJob(),
			"ibm_cis_alert":                             cis.ResourceIBMCISAlert(),
			"ibm_cis_routing":                           cis.ResourceIBMCISRouting(),
			"ibm_cis_waf_group":                         cis.ResourceIBMCISWAFGroup(),
			"ibm_cis_cache_settings":                    cis.ResourceIBMCISCacheSettings(),
			"ibm_cis_custom_page":                       cis.ResourceIBMCISCustomPage(),
			"ibm_cis_waf_rule":                          cis.ResourceIBMCISWAFRule(),
			"ibm_cis_certificate_order":                 cis.ResourceIBMCISCertificateOrder(),
			"ibm_cis_filter":                            cis.ResourceIBMCISFilter(),
			"ibm_cis_firewall_rule":                     cis.ResourceIBMCISFirewallrules(),
			"ibm_cloudant":                              cloudant.ResourceIBMCloudant(),
			"ibm_cloudant_database":                     cloudant.ResourceIBMCloudantDatabase(),
			"ibm_cloud_shell_account_settings":          cloudshell.ResourceIBMCloudShellAccountSettings(),
			"ibm_compute_autoscale_group":               classicinfrastructure.ResourceIBMComputeAutoScaleGroup(),
			"ibm_compute_autoscale_policy":              classicinfrastructure.ResourceIBMComputeAutoScalePolicy(),
			"ibm_compute_bare_metal":                    classicinfrastructure.ResourceIBMComputeBareMetal(),
			"ibm_compute_dedicated_host":                classicinfrastructure.ResourceIBMComputeDedicatedHost(),
			"ibm_compute_monitor":                       classicinfrastructure.ResourceIBMComputeMonitor(),
			"ibm_compute_placement_group":               classicinfrastructure.ResourceIBMComputePlacementGroup(),
			"ibm_compute_reserved_capacity":             classicinfrastructure.ResourceIBMComputeReservedCapacity(),
			"ibm_compute_provisioning_hook":             classicinfrastructure.ResourceIBMComputeProvisioningHook(),
			"ibm_compute_ssh_key":                       classicinfrastructure.ResourceIBMComputeSSHKey(),
			"ibm_compute_ssl_certificate":               classicinfrastructure.ResourceIBMComputeSSLCertificate(),
			"ibm_compute_user":                          classicinfrastructure.ResourceIBMComputeUser(),
			"ibm_compute_vm_instance":                   classicinfrastructure.ResourceIBMComputeVmInstance(),
			"ibm_container_addons":                      kubernetes.ResourceIBMContainerAddOns(),
			"ibm_container_alb":                         kubernetes.ResourceIBMContainerALB(),
			"ibm_container_alb_create":                  kubernetes.ResourceIBMContainerAlbCreate(),
			"ibm_container_api_key_reset":               kubernetes.ResourceIBMContainerAPIKeyReset(),
			"ibm_container_vpc_alb":                     kubernetes.ResourceIBMContainerVpcALB(),
			"ibm_container_vpc_alb_create":              kubernetes.ResourceIBMContainerVpcAlbCreateNew(),
			"ibm_container_vpc_worker_pool":             kubernetes.ResourceIBMContainerVpcWorkerPool(),
			"ibm_container_vpc_worker":                  kubernetes.ResourceIBMContainerVpcWorker(),
			"ibm_container_vpc_cluster":                 kubernetes.ResourceIBMContainerVpcCluster(),
			"ibm_container_alb_cert":                    kubernetes.ResourceIBMContainerALBCert(),
			"ibm_container_ingress_instance":            kubernetes.ResourceIBMContainerIngressInstance(),
			"ibm_container_ingress_secret_tls":          kubernetes.ResourceIBMContainerIngressSecretTLS(),
			"ibm_container_ingress_secret_opaque":       kubernetes.ResourceIBMContainerIngressSecretOpaque(),
			"ibm_container_cluster":                     kubernetes.ResourceIBMContainerCluster(),
			"ibm_container_cluster_feature":             kubernetes.ResourceIBMContainerClusterFeature(),
			"ibm_container_bind_service":                kubernetes.ResourceIBMContainerBindService(),
			"ibm_container_worker_pool":                 kubernetes.ResourceIBMContainerWorkerPool(),
			"ibm_container_worker_pool_zone_attachment": kubernetes.ResourceIBMContainerWorkerPoolZoneAttachment(),
			"ibm_container_storage_attachment":          kubernetes.ResourceIBMContainerVpcWorkerVolumeAttachment(),
			"ibm_container_nlb_dns":                     kubernetes.ResourceIBMContainerNlbDns(),
			"ibm_container_dedicated_host_pool":         kubernetes.ResourceIBMContainerDedicatedHostPool(),
			"ibm_container_dedicated_host":              kubernetes.ResourceIBMContainerDedicatedHost(),
			"ibm_cr_namespace":                          registry.ResourceIBMCrNamespace(),
			"ibm_cr_retention_policy":                   registry.ResourceIBMCrRetentionPolicy(),
			"ibm_ob_logging":                            kubernetes.ResourceIBMObLogging(),
			"ibm_ob_monitoring":                         kubernetes.ResourceIBMObMonitoring(),
			"ibm_cos_bucket":                            cos.ResourceIBMCOSBucket(),
			"ibm_cos_bucket_replication_rule":           cos.ResourceIBMCOSBucketReplicationConfiguration(),
			"ibm_cos_bucket_object":                     cos.ResourceIBMCOSBucketObject(),
			"ibm_cos_bucket_object_lock_configuration":  cos.ResourceIBMCOSBucketObjectlock(),
			"ibm_dns_domain":                            classicinfrastructure.ResourceIBMDNSDomain(),
			"ibm_dns_domain_registration_nameservers":   classicinfrastructure.ResourceIBMDNSDomainRegistrationNameservers(),
			"ibm_dns_secondary":                         classicinfrastructure.ResourceIBMDNSSecondary(),
			"ibm_dns_record":                            classicinfrastructure.ResourceIBMDNSRecord(),
			"ibm_event_streams_topic":                   eventstreams.ResourceIBMEventStreamsTopic(),
			"ibm_event_streams_schema":                  eventstreams.ResourceIBMEventStreamsSchema(),
			"ibm_firewall":                              classicinfrastructure.ResourceIBMFirewall(),
			"ibm_firewall_policy":                       classicinfrastructure.ResourceIBMFirewallPolicy(),
			"ibm_hpcs":                                  hpcs.ResourceIBMHPCS(),
			"ibm_hpcs_managed_key":                      hpcs.ResourceIbmManagedKey(),
			"ibm_hpcs_key_template":                     hpcs.ResourceIbmKeyTemplate(),
			"ibm_hpcs_keystore":                         hpcs.ResourceIbmKeystore(),
			"ibm_hpcs_vault":                            hpcs.ResourceIbmVault(),
			"ibm_iam_access_group":                      iamaccessgroup.ResourceIBMIAMAccessGroup(),
			"ibm_iam_access_group_account_settings":     iamaccessgroup.ResourceIBMIAMAccessGroupAccountSettings(),
			"ibm_iam_account_settings":                  iamidentity.ResourceIBMIAMAccountSettings(),
			"ibm_iam_custom_role":                       iampolicy.ResourceIBMIAMCustomRole(),
			"ibm_iam_access_group_dynamic_rule":         iamaccessgroup.ResourceIBMIAMDynamicRule(),
			"ibm_iam_access_group_members":              iamaccessgroup.ResourceIBMIAMAccessGroupMembers(),
			"ibm_iam_access_group_policy":               iampolicy.ResourceIBMIAMAccessGroupPolicy(),
			"ibm_iam_authorization_policy":              iampolicy.ResourceIBMIAMAuthorizationPolicy(),
			"ibm_iam_authorization_policy_detach":       iampolicy.ResourceIBMIAMAuthorizationPolicyDetach(),
			"ibm_iam_user_policy":                       iampolicy.ResourceIBMIAMUserPolicy(),
			"ibm_iam_user_settings":                     iamidentity.ResourceIBMIAMUserSettings(),
			"ibm_iam_service_id":                        iamidentity.ResourceIBMIAMServiceID(),
			"ibm_iam_service_api_key":                   iamidentity.ResourceIBMIAMServiceAPIKey(),
			"ibm_iam_service_policy":                    iampolicy.ResourceIBMIAMServicePolicy(),
			"ibm_iam_user_invite":                       iampolicy.ResourceIBMIAMUserInvite(),
			"ibm_iam_api_key":                           iamidentity.ResourceIBMIAMApiKey(),
			"ibm_iam_trusted_profile":                   iamidentity.ResourceIBMIAMTrustedProfile(),
			"ibm_iam_trusted_profile_claim_rule":        iamidentity.ResourceIBMIAMTrustedProfileClaimRule(),
			"ibm_iam_trusted_profile_link":              iamidentity.ResourceIBMIAMTrustedProfileLink(),
			"ibm_iam_trusted_profile_policy":            iampolicy.ResourceIBMIAMTrustedProfilePolicy(),
			"ibm_ipsec_vpn":                             classicinfrastructure.ResourceIBMIPSecVPN(),

			"ibm_lb":                               classicinfrastructure.ResourceIBMLb(),
			"ibm_lbaas":                            classicinfrastructure.ResourceIBMLbaas(),
			"ibm_lbaas_health_monitor":             classicinfrastructure.ResourceIBMLbaasHealthMonitor(),
			"ibm_lbaas_server_instance_attachment": classicinfrastructure.ResourceIBMLbaasServerInstanceAttachment(),
			"ibm_lb_service":                       classicinfrastructure.ResourceIBMLbService(),
			"ibm_lb_service_group":                 classicinfrastructure.ResourceIBMLbServiceGroup(),
			"ibm_lb_vpx":                           classicinfrastructure.ResourceIBMLbVpx(),
			"ibm_lb_vpx_ha":                        classicinfrastructure.ResourceIBMLbVpxHa(),
			"ibm_lb_vpx_service":                   classicinfrastructure.ResourceIBMLbVpxService(),
			"ibm_lb_vpx_vip":                       classicinfrastructure.ResourceIBMLbVpxVip(),
			"ibm_multi_vlan_firewall":              classicinfrastructure.ResourceIBMMultiVlanFirewall(),
			"ibm_network_gateway":                  classicinfrastructure.ResourceIBMNetworkGateway(),
			"ibm_network_gateway_vlan_association": classicinfrastructure.ResourceIBMNetworkGatewayVlanAttachment(),
			"ibm_network_interface_sg_attachment":  classicinfrastructure.ResourceIBMNetworkInterfaceSGAttachment(),
			"ibm_network_public_ip":                classicinfrastructure.ResourceIBMNetworkPublicIp(),
			"ibm_network_vlan":                     classicinfrastructure.ResourceIBMNetworkVlan(),
			"ibm_network_vlan_spanning":            classicinfrastructure.ResourceIBMNetworkVlanSpan(),
			"ibm_object_storage_account":           classicinfrastructure.ResourceIBMObjectStorageAccount(),
			"ibm_org":                              cloudfoundry.ResourceIBMOrg(),
			"ibm_pn_application_chrome":            pushnotification.ResourceIBMPNApplicationChrome(),
			"ibm_app_config_environment":           appconfiguration.ResourceIBMAppConfigEnvironment(),
			"ibm_app_config_collection":            appconfiguration.ResourceIBMAppConfigCollection(),
			"ibm_app_config_feature":               appconfiguration.ResourceIBMIbmAppConfigFeature(),
			"ibm_app_config_property":              appconfiguration.ResourceIBMIbmAppConfigProperty(),
			"ibm_app_config_segment":               appconfiguration.ResourceIBMIbmAppConfigSegment(),
			"ibm_app_config_snapshot":              appconfiguration.ResourceIBMIbmAppConfigSnapshot(),
			"ibm_kms_key":                          kms.ResourceIBMKmskey(),
			"ibm_kms_key_with_policy_overrides":    kms.ResourceIBMKmsKeyWithPolicyOverrides(),
			"ibm_kms_key_alias":                    kms.ResourceIBMKmskeyAlias(),
			"ibm_kms_key_rings":                    kms.ResourceIBMKmskeyRings(),
			"ibm_kms_key_policies":                 kms.ResourceIBMKmskeyPolicies(),
			"ibm_kp_key":                           kms.ResourceIBMkey(),
			"ibm_kms_instance_policies":            kms.ResourceIBMKmsInstancePolicy(),
			"ibm_resource_group":                   resourcemanager.ResourceIBMResourceGroup(),
			"ibm_resource_instance":                resourcecontroller.ResourceIBMResourceInstance(),
			"ibm_resource_key":                     resourcecontroller.ResourceIBMResourceKey(),
			"ibm_security_group":                   classicinfrastructure.ResourceIBMSecurityGroup(),
			"ibm_security_group_rule":              classicinfrastructure.ResourceIBMSecurityGroupRule(),
			"ibm_service_instance":                 cloudfoundry.ResourceIBMServiceInstance(),
			"ibm_service_key":                      cloudfoundry.ResourceIBMServiceKey(),
			"ibm_space":                            cloudfoundry.ResourceIBMSpace(),
			"ibm_storage_evault":                   classicinfrastructure.ResourceIBMStorageEvault(),
			"ibm_storage_block":                    classicinfrastructure.ResourceIBMStorageBlock(),
			"ibm_storage_file":                     classicinfrastructure.ResourceIBMStorageFile(),
			"ibm_subnet":                           classicinfrastructure.ResourceIBMSubnet(),
			"ibm_dns_reverse_record":               classicinfrastructure.ResourceIBMDNSReverseRecord(),
			"ibm_ssl_certificate":                  classicinfrastructure.ResourceIBMSSLCertificate(),
			"ibm_cdn":                              classicinfrastructure.ResourceIBMCDN(),
			"ibm_hardware_firewall_shared":         classicinfrastructure.ResourceIBMFirewallShared(),

			// //Added for Power Colo

			"ibm_pi_key":                             power.ResourceIBMPIKey(),
			"ibm_pi_volume":                          power.ResourceIBMPIVolume(),
			"ibm_pi_volume_onboarding":               power.ResourceIBMPIVolumeOnboarding(),
			"ibm_pi_volume_group":                    power.ResourceIBMPIVolumeGroup(),
			"ibm_pi_volume_group_action":             power.ResourceIBMPIVolumeGroupAction(),
			"ibm_pi_network":                         power.ResourceIBMPINetwork(),
			"ibm_pi_instance":                        power.ResourceIBMPIInstance(),
			"ibm_pi_instance_action":                 power.ResourceIBMPIInstanceAction(),
			"ibm_pi_volume_attach":                   power.ResourceIBMPIVolumeAttach(),
			"ibm_pi_capture":                         power.ResourceIBMPICapture(),
			"ibm_pi_image":                           power.ResourceIBMPIImage(),
			"ibm_pi_image_export":                    power.ResourceIBMPIImageExport(),
			"ibm_pi_network_port":                    power.ResourceIBMPINetworkPort(),
			"ibm_pi_snapshot":                        power.ResourceIBMPISnapshot(),
			"ibm_pi_network_port_attach":             power.ResourceIBMPINetworkPortAttach(),
			"ibm_pi_dhcp":                            power.ResourceIBMPIDhcp(),
			"ibm_pi_cloud_connection":                power.ResourceIBMPICloudConnection(),
			"ibm_pi_cloud_connection_network_attach": power.ResourceIBMPICloudConnectionNetworkAttach(),
			"ibm_pi_ike_policy":                      power.ResourceIBMPIIKEPolicy(),
			"ibm_pi_ipsec_policy":                    power.ResourceIBMPIIPSecPolicy(),
			"ibm_pi_vpn_connection":                  power.ResourceIBMPIVPNConnection(),
			"ibm_pi_console_language":                power.ResourceIBMPIInstanceConsoleLanguage(),
			"ibm_pi_placement_group":                 power.ResourceIBMPIPlacementGroup(),
			"ibm_pi_spp_placement_group":             power.ResourceIBMPISPPPlacementGroup(),
			"ibm_pi_shared_processor_pool":           power.ResourceIBMPISharedProcessorPool(),

			// //Private DNS related resources
			"ibm_dns_zone":              dnsservices.ResourceIBMPrivateDNSZone(),
			"ibm_dns_permitted_network": dnsservices.ResourceIBMPrivateDNSPermittedNetwork(),
			"ibm_dns_resource_record":   dnsservices.ResourceIBMPrivateDNSResourceRecord(),
			"ibm_dns_glb_monitor":       dnsservices.ResourceIBMPrivateDNSGLBMonitor(),
			"ibm_dns_glb_pool":          dnsservices.ResourceIBMPrivateDNSGLBPool(),
			"ibm_dns_glb":               dnsservices.ResourceIBMPrivateDNSGLB(),

			// //Added for Custom Resolver
			"ibm_dns_custom_resolver":                 dnsservices.ResourceIBMPrivateDNSCustomResolver(),
			"ibm_dns_custom_resolver_location":        dnsservices.ResourceIBMPrivateDNSCRLocation(),
			"ibm_dns_custom_resolver_forwarding_rule": dnsservices.ResourceIBMPrivateDNSForwardingRule(),
			"ibm_dns_custom_resolver_secondary_zone":  dnsservices.ResourceIBMPrivateDNSSecondaryZone(),
			"ibm_dns_linked_zone":                     dnsservices.ResourceIBMDNSLinkedZone(),

			// //Direct Link related resources
			"ibm_dl_gateway":            directlink.ResourceIBMDLGateway(),
			"ibm_dl_virtual_connection": directlink.ResourceIBMDLGatewayVC(),
			"ibm_dl_provider_gateway":   directlink.ResourceIBMDLProviderGateway(),
			"ibm_dl_route_report":       directlink.ResourceIBMDLGatewayRouteReport(),
			"ibm_dl_gateway_action":     directlink.ResourceIBMDLGatewayAction(),

			// //Added for Transit Gateway
			"ibm_tg_gateway":                  transitgateway.ResourceIBMTransitGateway(),
			"ibm_tg_connection":               transitgateway.ResourceIBMTransitGatewayConnection(),
			"ibm_tg_connection_action":        transitgateway.ResourceIBMTransitGatewayConnectionAction(),
			"ibm_tg_connection_prefix_filter": transitgateway.ResourceIBMTransitGatewayConnectionPrefixFilter(),
			"ibm_tg_route_report":             transitgateway.ResourceIBMTransitGatewayRouteReport(),

			// //Catalog related resources
			"ibm_cm_offering_instance": catalogmanagement.ResourceIBMCmOfferingInstance(),
			"ibm_cm_catalog":           catalogmanagement.ResourceIBMCmCatalog(),
			"ibm_cm_offering":          catalogmanagement.ResourceIBMCmOffering(),
			"ibm_cm_version":           catalogmanagement.ResourceIBMCmVersion(),
			"ibm_cm_validation":        catalogmanagement.ResourceIBMCmValidation(),
			"ibm_cm_object":            catalogmanagement.ResourceIBMCmObject(),

			// //Added for enterprise
			"ibm_enterprise":               enterprise.ResourceIBMEnterprise(),
			"ibm_enterprise_account_group": enterprise.ResourceIBMEnterpriseAccountGroup(),
			"ibm_enterprise_account":       enterprise.ResourceIBMEnterpriseAccount(),

			//Added for Schematics
			"ibm_schematics_workspace":      schematics.ResourceIBMSchematicsWorkspace(),
			"ibm_schematics_action":         schematics.ResourceIBMSchematicsAction(),
			"ibm_schematics_job":            schematics.ResourceIBMSchematicsJob(),
			"ibm_schematics_inventory":      schematics.ResourceIBMSchematicsInventory(),
			"ibm_schematics_resource_query": schematics.ResourceIBMSchematicsResourceQuery(),

			// //Added for Secrets Manager
			"ibm_sm_secret_group":                                                secretsmanager.AddInstanceFields(secretsmanager.ResourceIbmSmSecretGroup()),
			"ibm_sm_arbitrary_secret":                                            secretsmanager.AddInstanceFields(secretsmanager.ResourceIbmSmArbitrarySecret()),
			"ibm_sm_imported_certificate":                                        secretsmanager.AddInstanceFields(secretsmanager.ResourceIbmSmImportedCertificate()),
			"ibm_sm_public_certificate":                                          secretsmanager.AddInstanceFields(secretsmanager.ResourceIbmSmPublicCertificate()),
			"ibm_sm_private_certificate":                                         secretsmanager.AddInstanceFields(secretsmanager.ResourceIbmSmPrivateCertificate()),
			"ibm_sm_iam_credentials_secret":                                      secretsmanager.AddInstanceFields(secretsmanager.ResourceIbmSmIamCredentialsSecret()),
			"ibm_sm_username_password_secret":                                    secretsmanager.AddInstanceFields(secretsmanager.ResourceIbmSmUsernamePasswordSecret()),
			"ibm_sm_kv_secret":                                                   secretsmanager.AddInstanceFields(secretsmanager.ResourceIbmSmKvSecret()),
			"ibm_sm_public_certificate_configuration_ca_lets_encrypt":            secretsmanager.AddInstanceFields(secretsmanager.ResourceIbmSmPublicCertificateConfigurationCALetsEncrypt()),
			"ibm_sm_public_certificate_configuration_dns_cis":                    secretsmanager.AddInstanceFields(secretsmanager.ResourceIbmSmConfigurationPublicCertificateDNSCis()),
			"ibm_sm_public_certificate_configuration_dns_classic_infrastructure": secretsmanager.AddInstanceFields(secretsmanager.ResourceIbmSmPublicCertificateConfigurationDNSClassicInfrastructure()),
			"ibm_sm_private_certificate_configuration_root_ca":                   secretsmanager.AddInstanceFields(secretsmanager.ResourceIbmSmPrivateCertificateConfigurationRootCA()),
			"ibm_sm_private_certificate_configuration_intermediate_ca":           secretsmanager.AddInstanceFields(secretsmanager.ResourceIbmSmPrivateCertificateConfigurationIntermediateCA()),
			"ibm_sm_private_certificate_configuration_template":                  secretsmanager.AddInstanceFields(secretsmanager.ResourceIbmSmPrivateCertificateConfigurationTemplate()),
			"ibm_sm_iam_credentials_configuration":                               secretsmanager.AddInstanceFields(secretsmanager.ResourceIbmSmIamCredentialsConfiguration()),
			"ibm_sm_public_certificate_action_validate_manual_dns":               secretsmanager.AddInstanceFields(secretsmanager.ResourceIbmSmPublicCertificateActionValidateManualDns()),
			"ibm_sm_en_registration":                                             secretsmanager.AddInstanceFields(secretsmanager.ResourceIbmSmEnRegistration()),
			"ibm_sm_private_certificate_configuration_action_sign_csr":           secretsmanager.AddInstanceFields(secretsmanager.ResourceIbmSmPrivateCertificateConfigurationActionSignCsr()),
			"ibm_sm_private_certificate_configuration_action_set_signed":         secretsmanager.AddInstanceFields(secretsmanager.ResourceIbmSmPrivateCertificateConfigurationActionSetSigned()),

			// //satellite  resources
			"ibm_satellite_location":                            satellite.ResourceIBMSatelliteLocation(),
			"ibm_satellite_host":                                satellite.ResourceIBMSatelliteHost(),
			"ibm_satellite_cluster":                             satellite.ResourceIBMSatelliteCluster(),
			"ibm_satellite_cluster_worker_pool":                 satellite.ResourceIBMSatelliteClusterWorkerPool(),
			"ibm_satellite_link":                                satellite.ResourceIBMSatelliteLink(),
			"ibm_satellite_endpoint":                            satellite.ResourceIBMSatelliteEndpoint(),
			"ibm_satellite_location_nlb_dns":                    satellite.ResourceIBMSatelliteLocationNlbDns(),
			"ibm_satellite_cluster_worker_pool_zone_attachment": satellite.ResourceIbmSatelliteClusterWorkerPoolZoneAttachment(),

			//Added for Resource Tag
			"ibm_resource_tag": globaltagging.ResourceIBMResourceTag(),

			// // Atracker
			"ibm_atracker_target":   atracker.ResourceIBMAtrackerTarget(),
			"ibm_atracker_route":    atracker.ResourceIBMAtrackerRoute(),
			"ibm_atracker_settings": atracker.ResourceIBMAtrackerSettings(),

			// Metrics Router
			"ibm_metrics_router_target":   metricsrouter.ResourceIBMMetricsRouterTarget(),
			"ibm_metrics_router_route":    metricsrouter.ResourceIBMMetricsRouterRoute(),
			"ibm_metrics_router_settings": metricsrouter.ResourceIBMMetricsRouterSettings(),

			// //Security and Compliance Center
			"ibm_scc_account_settings":    scc.ResourceIBMSccAccountSettings(),
			"ibm_scc_rule":                scc.ResourceIBMSccRule(),
			"ibm_scc_rule_attachment":     scc.ResourceIBMSccRuleAttachment(),
			"ibm_scc_template":            scc.ResourceIBMSccTemplate(),
			"ibm_scc_template_attachment": scc.ResourceIBMSccTemplateAttachment(),

			// // Added for Context Based Restrictions
			"ibm_cbr_zone": contextbasedrestrictions.ResourceIBMCbrZone(),
			"ibm_cbr_rule": contextbasedrestrictions.ResourceIBMCbrRule(),

			// // Added for Event Notifications
			"ibm_en_source":                 eventnotification.ResourceIBMEnSource(),
			"ibm_en_topic":                  eventnotification.ResourceIBMEnTopic(),
			"ibm_en_destination_webhook":    eventnotification.ResourceIBMEnWebhookDestination(),
			"ibm_en_destination_android":    eventnotification.ResourceIBMEnFCMDestination(),
			"ibm_en_destination_chrome":     eventnotification.ResourceIBMEnChromeDestination(),
			"ibm_en_destination_firefox":    eventnotification.ResourceIBMEnFirefoxDestination(),
			"ibm_en_destination_ios":        eventnotification.ResourceIBMEnAPNSDestination(),
			"ibm_en_destination_slack":      eventnotification.ResourceIBMEnSlackDestination(),
			"ibm_en_subscription_sms":       eventnotification.ResourceIBMEnSMSSubscription(),
			"ibm_en_subscription_email":     eventnotification.ResourceIBMEnEmailSubscription(),
			"ibm_en_subscription_webhook":   eventnotification.ResourceIBMEnWebhookSubscription(),
			"ibm_en_subscription_android":   eventnotification.ResourceIBMEnFCMSubscription(),
			"ibm_en_subscription_ios":       eventnotification.ResourceIBMEnFCMSubscription(),
			"ibm_en_subscription_chrome":    eventnotification.ResourceIBMEnFCMSubscription(),
			"ibm_en_subscription_firefox":   eventnotification.ResourceIBMEnFCMSubscription(),
			"ibm_en_subscription_slack":     eventnotification.ResourceIBMEnSlackSubscription(),
			"ibm_en_subscription_safari":    eventnotification.ResourceIBMEnFCMSubscription(),
			"ibm_en_destination_safari":     eventnotification.ResourceIBMEnSafariDestination(),
			"ibm_en_destination_msteams":    eventnotification.ResourceIBMEnMSTeamsDestination(),
			"ibm_en_subscription_msteams":   eventnotification.ResourceIBMEnFCMSubscription(),
			"ibm_en_destination_cf":         eventnotification.ResourceIBMEnCFDestination(),
			"ibm_en_subscription_cf":        eventnotification.ResourceIBMEnFCMSubscription(),
			"ibm_en_destination_pagerduty":  eventnotification.ResourceIBMEnPagerDutyDestination(),
			"ibm_en_subscription_pagerduty": eventnotification.ResourceIBMEnFCMSubscription(),
			"ibm_en_integration":            eventnotification.ResourceIBMEnIntegration(),
			"ibm_en_destination_sn":         eventnotification.ResourceIBMEnServiceNowDestination(),
			"ibm_en_subscription_sn":        eventnotification.ResourceIBMEnFCMSubscription(),
			"ibm_en_destination_ce":         eventnotification.ResourceIBMEnCodeEngineDestination(),
			"ibm_en_subscription_ce":        eventnotification.ResourceIBMEnFCMSubscription(),
			"ibm_en_destination_cos":        eventnotification.ResourceIBMEnCOSDestination(),
			"ibm_en_subscription_cos":       eventnotification.ResourceIBMEnFCMSubscription(),

			// // Added for Toolchain
			"ibm_cd_toolchain":                         cdtoolchain.ResourceIBMCdToolchain(),
			"ibm_cd_toolchain_tool_keyprotect":         cdtoolchain.ResourceIBMCdToolchainToolKeyprotect(),
			"ibm_cd_toolchain_tool_secretsmanager":     cdtoolchain.ResourceIBMCdToolchainToolSecretsmanager(),
			"ibm_cd_toolchain_tool_bitbucketgit":       cdtoolchain.ResourceIBMCdToolchainToolBitbucketgit(),
			"ibm_cd_toolchain_tool_githubconsolidated": cdtoolchain.ResourceIBMCdToolchainToolGithubconsolidated(),
			"ibm_cd_toolchain_tool_gitlab":             cdtoolchain.ResourceIBMCdToolchainToolGitlab(),
			"ibm_cd_toolchain_tool_hostedgit":          cdtoolchain.ResourceIBMCdToolchainToolHostedgit(),
			"ibm_cd_toolchain_tool_artifactory":        cdtoolchain.ResourceIBMCdToolchainToolArtifactory(),
			"ibm_cd_toolchain_tool_custom":             cdtoolchain.ResourceIBMCdToolchainToolCustom(),
			"ibm_cd_toolchain_tool_pipeline":           cdtoolchain.ResourceIBMCdToolchainToolPipeline(),
			"ibm_cd_toolchain_tool_devopsinsights":     cdtoolchain.ResourceIBMCdToolchainToolDevopsinsights(),
			"ibm_cd_toolchain_tool_slack":              cdtoolchain.ResourceIBMCdToolchainToolSlack(),
			"ibm_cd_toolchain_tool_sonarqube":          cdtoolchain.ResourceIBMCdToolchainToolSonarqube(),
			"ibm_cd_toolchain_tool_hashicorpvault":     cdtoolchain.ResourceIBMCdToolchainToolHashicorpvault(),
			"ibm_cd_toolchain_tool_securitycompliance": cdtoolchain.ResourceIBMCdToolchainToolSecuritycompliance(),
			"ibm_cd_toolchain_tool_privateworker":      cdtoolchain.ResourceIBMCdToolchainToolPrivateworker(),
			"ibm_cd_toolchain_tool_appconfig":          cdtoolchain.ResourceIBMCdToolchainToolAppconfig(),
			"ibm_cd_toolchain_tool_jenkins":            cdtoolchain.ResourceIBMCdToolchainToolJenkins(),
			"ibm_cd_toolchain_tool_nexus":              cdtoolchain.ResourceIBMCdToolchainToolNexus(),
			"ibm_cd_toolchain_tool_pagerduty":          cdtoolchain.ResourceIBMCdToolchainToolPagerduty(),
			"ibm_cd_toolchain_tool_saucelabs":          cdtoolchain.ResourceIBMCdToolchainToolSaucelabs(),
			"ibm_cd_toolchain_tool_jira":               cdtoolchain.ResourceIBMCdToolchainToolJira(),
			"ibm_cd_toolchain_tool_eventnotifications": cdtoolchain.ResourceIBMCdToolchainToolEventnotifications(),

			// // Added for Tekton Pipeline
			"ibm_cd_tekton_pipeline_definition":       cdtektonpipeline.ResourceIBMCdTektonPipelineDefinition(),
			"ibm_cd_tekton_pipeline_trigger_property": cdtektonpipeline.ResourceIBMCdTektonPipelineTriggerProperty(),
			"ibm_cd_tekton_pipeline_property":         cdtektonpipeline.ResourceIBMCdTektonPipelineProperty(),
			"ibm_cd_tekton_pipeline_trigger":          cdtektonpipeline.ResourceIBMCdTektonPipelineTrigger(),
			"ibm_cd_tekton_pipeline":                  cdtektonpipeline.ResourceIBMCdTektonPipeline(),

			// // Added for Code Engine
			"ibm_code_engine_app":        codeengine.ResourceIbmCodeEngineApp(),
			"ibm_code_engine_binding":    codeengine.ResourceIbmCodeEngineBinding(),
			"ibm_code_engine_build":      codeengine.ResourceIbmCodeEngineBuild(),
			"ibm_code_engine_config_map": codeengine.ResourceIbmCodeEngineConfigMap(),
			"ibm_code_engine_job":        codeengine.ResourceIbmCodeEngineJob(),
			"ibm_code_engine_project":    codeengine.ResourceIbmCodeEngineProject(),
			"ibm_code_engine_secret":     codeengine.ResourceIbmCodeEngineSecret(),

			// Added for Project
			"ibm_project":        project.ResourceIbmProject(),
			"ibm_project_config": project.ResourceIbmProjectConfig(),
		},

		ConfigureFunc: providerConfigure,
	}
}

var globalValidatorDict validate.ValidatorDict
var initOnce sync.Once

func init() {
	validate.SetValidatorDict(Validator())
}

// Validator return validator
func Validator() validate.ValidatorDict {
	initOnce.Do(func() {
		globalValidatorDict = validate.ValidatorDict{
			ResourceValidatorDictionary: map[string]*validate.ResourceValidator{
				"ibm_iam_account_settings":        iamidentity.ResourceIBMIAMAccountSettingsValidator(),
				"ibm_iam_custom_role":             iampolicy.ResourceIBMIAMCustomRoleValidator(),
				"ibm_cis_healthcheck":             cis.ResourceIBMCISHealthCheckValidator(),
				"ibm_cis_rate_limit":              cis.ResourceIBMCISRateLimitValidator(),
				"ibm_cis":                         cis.ResourceIBMCISValidator(),
				"ibm_cis_domain_settings":         cis.ResourceIBMCISDomainSettingValidator(),
				"ibm_cis_domain":                  cis.ResourceIBMCISDomainValidator(),
				"ibm_cis_tls_settings":            cis.ResourceIBMCISTLSSettingsValidator(),
				"ibm_cis_routing":                 cis.ResourceIBMCISRoutingValidator(),
				"ibm_cis_page_rule":               cis.ResourceIBMCISPageRuleValidator(),
				"ibm_cis_waf_package":             cis.ResourceIBMCISWAFPackageValidator(),
				"ibm_cis_waf_group":               cis.ResourceIBMCISWAFGroupValidator(),
				"ibm_cis_certificate_upload":      cis.ResourceIBMCISCertificateUploadValidator(),
				"ibm_cis_cache_settings":          cis.ResourceIBMCISCacheSettingsValidator(),
				"ibm_cis_custom_page":             cis.ResourceIBMCISCustomPageValidator(),
				"ibm_cis_firewall":                cis.ResourceIBMCISFirewallValidator(),
				"ibm_cis_range_app":               cis.ResourceIBMCISRangeAppValidator(),
				"ibm_cis_waf_rule":                cis.ResourceIBMCISWAFRuleValidator(),
				"ibm_cis_certificate_order":       cis.ResourceIBMCISCertificateOrderValidator(),
				"ibm_cis_filter":                  cis.ResourceIBMCISFilterValidator(),
				"ibm_cis_firewall_rules":          cis.ResourceIBMCISFirewallrulesValidator(),
				"ibm_cis_webhook":                 cis.ResourceIBMCISWebhooksValidator(),
				"ibm_cis_alert":                   cis.ResourceIBMCISAlertValidator(),
				"ibm_cis_dns_record":              cis.ResourceIBMCISDnsRecordValidator(),
				"ibm_cis_dns_records_import":      cis.ResourceIBMCISDnsRecordsImportValidator(),
				"ibm_cis_edge_functions_action":   cis.ResourceIBMCISEdgeFunctionsActionValidator(),
				"ibm_cis_edge_functions_trigger":  cis.ResourceIBMCISEdgeFunctionsTriggerValidator(),
				"ibm_cis_global_load_balancer":    cis.ResourceIBMCISGlbValidator(),
				"ibm_cis_logpush_job":             cis.ResourceIBMCISLogPushJobValidator(),
				"ibm_cis_mtls_app":                cis.ResourceIBMCISMtlsAppValidator(),
				"ibm_cis_mtls":                    cis.ResourceIBMCISMtlsValidator(),
				"ibm_cis_bot_management":          cis.ResourceIBMCISBotManagementValidator(),
				"ibm_cis_origin_auth":             cis.ResourceIBMCISOriginAuthPullValidator(),
				"ibm_cis_origin_pool":             cis.ResourceIBMCISPoolValidator(),
				"ibm_container_cluster":           kubernetes.ResourceIBMContainerClusterValidator(),
				"ibm_container_worker_pool":       kubernetes.ResourceIBMContainerWorkerPoolValidator(),
				"ibm_container_vpc_worker_pool":   kubernetes.ResourceIBMContainerVPCWorkerPoolValidator(),
				"ibm_container_vpc_worker":        kubernetes.ResourceIBMContainerVPCWorkerValidator(),
				"ibm_container_vpc_cluster":       kubernetes.ResourceIBMContainerVpcClusterValidator(),
				"ibm_cos_bucket":                  cos.ResourceIBMCOSBucketValidator(),
				"ibm_cr_namespace":                registry.ResourceIBMCrNamespaceValidator(),
				"ibm_tg_gateway":                  transitgateway.ResourceIBMTGValidator(),
				"ibm_app_config_feature":          appconfiguration.ResourceIBMAppConfigFeatureValidator(),
				"ibm_tg_connection":               transitgateway.ResourceIBMTransitGatewayConnectionValidator(),
				"ibm_tg_connection_action":        transitgateway.ResourceIBMTransitGatewayConnectionActionValidator(),
				"ibm_tg_connection_prefix_filter": transitgateway.ResourceIBMTransitGatewayConnectionPrefixFilterValidator(),
				"ibm_dl_virtual_connection":       directlink.ResourceIBMDLGatewayVCValidator(),
				"ibm_dl_gateway":                  directlink.ResourceIBMDLGatewayValidator(),
				"ibm_dl_provider_gateway":         directlink.ResourceIBMDLProviderGatewayValidator(),
				"ibm_dl_gateway_action":           directlink.ResourceIBMDLGatewayActionValidator(),
				"ibm_database":                    database.ResourceIBMICDValidator(),
				"ibm_function_package":            functions.ResourceIBMFuncPackageValidator(),
				"ibm_function_action":             functions.ResourceIBMFuncActionValidator(),
				"ibm_function_rule":               functions.ResourceIBMFuncRuleValidator(),
				"ibm_function_trigger":            functions.ResourceIBMFuncTriggerValidator(),
				"ibm_function_namespace":          functions.ResourceIBMFuncNamespaceValidator(),
				"ibm_hpcs":                        hpcs.ResourceIBMHPCSValidator(),
				"ibm_hpcs_managed_key":            hpcs.ResourceIbmManagedKeyValidator(),
				"ibm_hpcs_keystore":               hpcs.ResourceIbmKeystoreValidator(),
				"ibm_hpcs_key_template":           hpcs.ResourceIbmKeyTemplateValidator(),
				"ibm_hpcs_vault":                  hpcs.ResourceIbmVaultValidator(),

				"ibm_kms_key_rings":                       kms.ResourceIBMKeyRingValidator(),
				"ibm_dns_glb_monitor":                     dnsservices.ResourceIBMPrivateDNSGLBMonitorValidator(),
				"ibm_dns_custom_resolver_forwarding_rule": dnsservices.ResourceIBMPrivateDNSForwardingRuleValidator(),
				"ibm_schematics_action":                   schematics.ResourceIBMSchematicsActionValidator(),
				"ibm_schematics_job":                      schematics.ResourceIBMSchematicsJobValidator(),
				"ibm_schematics_workspace":                schematics.ResourceIBMSchematicsWorkspaceValidator(),
				"ibm_schematics_inventory":                schematics.ResourceIBMSchematicsInventoryValidator(),
				"ibm_schematics_resource_query":           schematics.ResourceIBMSchematicsResourceQueryValidator(),
				"ibm_resource_instance":                   resourcecontroller.ResourceIBMResourceInstanceValidator(),
				"ibm_resource_key":                        resourcecontroller.ResourceIBMResourceKeyValidator(),
				"ibm_resource_tag":                        globaltagging.ResourceIBMResourceTagValidator(),
				"ibm_satellite_location":                  satellite.ResourceIBMSatelliteLocationValidator(),
				"ibm_satellite_cluster":                   satellite.ResourceIBMSatelliteClusterValidator(),
				"ibm_pi_volume":                           power.ResourceIBMPIVolumeValidator(),
				"ibm_atracker_target":                     atracker.ResourceIBMAtrackerTargetValidator(),
				"ibm_atracker_route":                      atracker.ResourceIBMAtrackerRouteValidator(),
				"ibm_atracker_settings":                   atracker.ResourceIBMAtrackerSettingsValidator(),
				"ibm_metrics_router_target":               metricsrouter.ResourceIBMMetricsRouterTargetValidator(),
				"ibm_metrics_router_route":                metricsrouter.ResourceIBMMetricsRouterRouteValidator(),
				"ibm_metrics_router_settings":             metricsrouter.ResourceIBMMetricsRouterSettingsValidator(),
				"ibm_satellite_endpoint":                  satellite.ResourceIBMSatelliteEndpointValidator(),
				"ibm_scc_account_settings":                scc.ResourceIBMSccAccountSettingsValidator(),
				"ibm_scc_rule":                            scc.ResourceIBMSccRuleValidator(),
				"ibm_scc_rule_attachment":                 scc.ResourceIBMSccRuleAttachmentValidator(),
				"ibm_scc_template":                        scc.ResourceIBMSccTemplateValidator(),
				"ibm_scc_template_attachment":             scc.ResourceIBMSccTemplateAttachmentValidator(),
				"ibm_cbr_zone":                            contextbasedrestrictions.ResourceIBMCbrZoneValidator(),
				"ibm_cbr_rule":                            contextbasedrestrictions.ResourceIBMCbrRuleValidator(),
				"ibm_satellite_host":                      satellite.ResourceIBMSatelliteHostValidator(),

				// // Added for Toolchains
				"ibm_cd_toolchain":                         cdtoolchain.ResourceIBMCdToolchainValidator(),
				"ibm_cd_toolchain_tool_keyprotect":         cdtoolchain.ResourceIBMCdToolchainToolKeyprotectValidator(),
				"ibm_cd_toolchain_tool_secretsmanager":     cdtoolchain.ResourceIBMCdToolchainToolSecretsmanagerValidator(),
				"ibm_cd_toolchain_tool_bitbucketgit":       cdtoolchain.ResourceIBMCdToolchainToolBitbucketgitValidator(),
				"ibm_cd_toolchain_tool_githubconsolidated": cdtoolchain.ResourceIBMCdToolchainToolGithubconsolidatedValidator(),
				"ibm_cd_toolchain_tool_gitlab":             cdtoolchain.ResourceIBMCdToolchainToolGitlabValidator(),
				"ibm_cd_toolchain_tool_hostedgit":          cdtoolchain.ResourceIBMCdToolchainToolHostedgitValidator(),
				"ibm_cd_toolchain_tool_artifactory":        cdtoolchain.ResourceIBMCdToolchainToolArtifactoryValidator(),
				"ibm_cd_toolchain_tool_custom":             cdtoolchain.ResourceIBMCdToolchainToolCustomValidator(),
				"ibm_cd_toolchain_tool_pipeline":           cdtoolchain.ResourceIBMCdToolchainToolPipelineValidator(),
				"ibm_cd_toolchain_tool_slack":              cdtoolchain.ResourceIBMCdToolchainToolSlackValidator(),
				"ibm_cd_toolchain_tool_devopsinsights":     cdtoolchain.ResourceIBMCdToolchainToolDevopsinsightsValidator(),
				"ibm_cd_toolchain_tool_sonarqube":          cdtoolchain.ResourceIBMCdToolchainToolSonarqubeValidator(),
				"ibm_cd_toolchain_tool_hashicorpvault":     cdtoolchain.ResourceIBMCdToolchainToolHashicorpvaultValidator(),
				"ibm_cd_toolchain_tool_securitycompliance": cdtoolchain.ResourceIBMCdToolchainToolSecuritycomplianceValidator(),
				"ibm_cd_toolchain_tool_privateworker":      cdtoolchain.ResourceIBMCdToolchainToolPrivateworkerValidator(),
				"ibm_cd_toolchain_tool_appconfig":          cdtoolchain.ResourceIBMCdToolchainToolAppconfigValidator(),
				"ibm_cd_toolchain_tool_jenkins":            cdtoolchain.ResourceIBMCdToolchainToolJenkinsValidator(),
				"ibm_cd_toolchain_tool_nexus":              cdtoolchain.ResourceIBMCdToolchainToolNexusValidator(),
				"ibm_cd_toolchain_tool_pagerduty":          cdtoolchain.ResourceIBMCdToolchainToolPagerdutyValidator(),
				"ibm_cd_toolchain_tool_saucelabs":          cdtoolchain.ResourceIBMCdToolchainToolSaucelabsValidator(),
				"ibm_cd_toolchain_tool_jira":               cdtoolchain.ResourceIBMCdToolchainToolJiraValidator(),
				"ibm_cd_toolchain_tool_eventnotifications": cdtoolchain.ResourceIBMCdToolchainToolEventnotificationsValidator(),

				// // Added for Tekton Pipeline
				"ibm_cd_tekton_pipeline_definition":       cdtektonpipeline.ResourceIBMCdTektonPipelineDefinitionValidator(),
				"ibm_cd_tekton_pipeline_trigger_property": cdtektonpipeline.ResourceIBMCdTektonPipelineTriggerPropertyValidator(),
				"ibm_cd_tekton_pipeline_property":         cdtektonpipeline.ResourceIBMCdTektonPipelinePropertyValidator(),
				"ibm_cd_tekton_pipeline_trigger":          cdtektonpipeline.ResourceIBMCdTektonPipelineTriggerValidator(),

				"ibm_container_addons":                      kubernetes.ResourceIBMContainerAddOnsValidator(),
				"ibm_container_alb_create":                  kubernetes.ResourceIBMContainerAlbCreateValidator(),
				"ibm_container_nlb_dns":                     kubernetes.ResourceIBMContainerNlbDnsValidator(),
				"ibm_container_vpc_alb_create":              kubernetes.ResourceIBMContainerVpcAlbCreateNewValidator(),
				"ibm_container_storage_attachment":          kubernetes.ResourceIBMContainerVpcWorkerVolumeAttachmentValidator(),
				"ibm_container_worker_pool_zone_attachment": kubernetes.ResourceIBMContainerWorkerPoolZoneAttachmentValidator(),
				"ibm_container_bind_service":                kubernetes.ResourceIBMContainerBindServiceValidator(),
				"ibm_container_alb_cert":                    kubernetes.ResourceIBMContainerALBCertValidator(),
				"ibm_container_ingress_instance":            kubernetes.ResourceIBMContainerIngressInstanceValidator(),
				"ibm_container_ingress_secret_tls":          kubernetes.ResourceIBMContainerIngressSecretTLSValidator(),
				"ibm_container_ingress_secret_opaque":       kubernetes.ResourceIBMContainerIngressSecretOpaqueValidator(),
				"ibm_container_cluster_feature":             kubernetes.ResourceIBMContainerClusterFeatureValidator(),

				"ibm_iam_access_group_dynamic_rule": iamaccessgroup.ResourceIBMIAMDynamicRuleValidator(),
				"ibm_iam_access_group_members":      iamaccessgroup.ResourceIBMIAMAccessGroupMembersValidator(),

				"ibm_iam_trusted_profile_claim_rule": iamidentity.ResourceIBMIAMTrustedProfileClaimRuleValidator(),
				"ibm_iam_trusted_profile_link":       iamidentity.ResourceIBMIAMTrustedProfileLinkValidator(),
				"ibm_iam_service_api_key":            iamidentity.ResourceIBMIAMServiceAPIKeyValidator(),

				"ibm_iam_trusted_profile_policy": iampolicy.ResourceIBMIAMTrustedProfilePolicyValidator(),
				"ibm_iam_access_group_policy":    iampolicy.ResourceIBMIAMAccessGroupPolicyValidator(),
				"ibm_iam_service_policy":         iampolicy.ResourceIBMIAMServicePolicyValidator(),
				"ibm_iam_authorization_policy":   iampolicy.ResourceIBMIAMAuthorizationPolicyValidator(),

				// // Added for Secrets Manager
				"ibm_sm_secret_group":                                                secretsmanager.ResourceIbmSmSecretGroupValidator(),
				"ibm_sm_en_registration":                                             secretsmanager.ResourceIbmSmEnRegistrationValidator(),
				"ibm_sm_public_certificate_configuration_dns_cis":                    secretsmanager.ResourceIbmSmConfigurationPublicCertificateDNSCisValidator(),
				"ibm_sm_public_certificate_configuration_dns_classic_infrastructure": secretsmanager.ResourceIbmSmPublicCertificateConfigurationDNSClassicInfrastructureValidator(),

				// // Added for Code Engine
				"ibm_code_engine_app":        codeengine.ResourceIbmCodeEngineAppValidator(),
				"ibm_code_engine_binding":    codeengine.ResourceIbmCodeEngineBindingValidator(),
				"ibm_code_engine_build":      codeengine.ResourceIbmCodeEngineBuildValidator(),
				"ibm_code_engine_config_map": codeengine.ResourceIbmCodeEngineConfigMapValidator(),
				"ibm_code_engine_job":        codeengine.ResourceIbmCodeEngineJobValidator(),
				"ibm_code_engine_project":    codeengine.ResourceIbmCodeEngineProjectValidator(),
				"ibm_code_engine_secret":     codeengine.ResourceIbmCodeEngineSecretValidator(),

				// Added for Project
				"ibm_project":        project.ResourceIbmProjectValidator(),
				"ibm_project_config": project.ResourceIbmProjectConfigValidator(),
			},
			DataSourceValidatorDictionary: map[string]*validate.ResourceValidator{
				"ibm_dl_offering_speeds": directlink.DataSourceIBMDLOfferingSpeedsValidator(),
				"ibm_dl_routers":         directlink.DataSourceIBMDLRoutersValidator(),
				"ibm_resource_instance":  resourcecontroller.DataSourceIBMResourceInstanceValidator(),
				"ibm_resource_key":       resourcecontroller.DataSourceIBMResourceKeyValidator(),
				"ibm_resource_group":     resourcemanager.DataSourceIBMResourceGroupValidator(),

				"ibm_secrets_manager_secret":      secretsmanager.DataSourceIBMSecretsManagerSecretValidator(),
				"ibm_secrets_manager_secrets":     secretsmanager.DataSourceIBMSecretsManagerSecretsValidator(),
				"ibm_cis_webhooks":                cis.DataSourceIBMCISAlertWebhooksValidator(),
				"ibm_cis_alerts":                  cis.DataSourceIBMCISAlertsValidator(),
				"ibm_cis_bot_managements":         cis.DataSourceIBMCISBotManagementValidator(),
				"ibm_cis_bot_analytics":           cis.DataSourceIBMCISBotAnalyticsValidator(),
				"ibm_cis_cache_settings":          cis.DataSourceIBMCISCacheSettingsValidator(),
				"ibm_cis_custom_certificates":     cis.DataSourceIBMCISCustomCertificatesValidator(),
				"ibm_cis_custom_pages":            cis.DataSourceIBMCISCustomPagesValidator(),
				"ibm_cis_dns_records":             cis.DataSourceIBMCISDNSRecordsValidator(),
				"ibm_cis_domain":                  cis.DataSourceIBMCISDomainValidator(),
				"ibm_cis_certificates":            cis.DataSourceIBMCISCertificatesValidator(),
				"ibm_cis_edge_functions_actions":  cis.DataSourceIBMCISEdgeFunctionsActionsValidator(),
				"ibm_cis_edge_functions_triggers": cis.DataSourceIBMCISEdgeFunctionsTriggersValidator(),
				"ibm_cis_filters":                 cis.DataSourceIBMCISFiltersValidator(),
				"ibm_cis_firewall_rules":          cis.DataSourceIBMCISFirewallRulesValidator(),
				"ibm_cis_firewall":                cis.DataSourceIBMCISFirewallsRecordValidator(),
				"ibm_cis_global_load_balancers":   cis.DataSourceIBMCISGlbsValidator(),
				"ibm_cis_healthchecks":            cis.DataSourceIBMCISHealthChecksValidator(),
				"ibm_cis_mtls_apps":               cis.DataSourceIBMCISMtlsAppValidator(),
				"ibm_cis_mtlss":                   cis.DataSourceIBMCISMtlsValidator(),
				"ibm_cis_origin_auths":            cis.DataSourceIBMCISOriginAuthPullValidator(),
				"ibm_cis_origin_pools":            cis.DataSourceIBMCISOriginPoolsValidator(),
				"ibm_cis_page_rules":              cis.DataSourceIBMCISPageRulesValidator(),
				"ibm_cis_range_apps":              cis.DataSourceIBMCISRangeAppsValidator(),
				"ibm_cis_rate_limit":              cis.DataSourceIBMCISRateLimitValidator(),
				"ibm_cis_waf_groups":              cis.DataSourceIBMCISWAFGroupsValidator(),
				"ibm_cis_waf_packages":            cis.DataSourceIBMCISWAFPackagesValidator(),
				"ibm_cis_waf_rules":               cis.DataSourceIBMCISWAFRulesValidator(),
				"ibm_cis_logpush_jobs":            cis.DataSourceIBMCISLogPushJobsValidator(),

				"ibm_cos_bucket": cos.DataSourceIBMCosBucketValidator(),

				"ibm_database_backups":                database.DataSourceIBMDatabaseBackupsValidator(),
				"ibm_database_connection":             database.DataSourceIBMDatabaseConnectionValidator(),
				"ibm_database_point_in_time_recovery": database.DataSourceIBMDatabasePointInTimeRecoveryValidator(),
				"ibm_database_remotes":                database.DataSourceIBMDatabaseRemotesValidator(),
				"ibm_database_tasks":                  database.DataSourceIBMDatabaseTasksValidator(),
				"ibm_database":                        database.DataSourceIBMDatabaseInstanceValidator(),

				"ibm_container_addons":                  kubernetes.DataSourceIBMContainerAddOnsValidator(),
				"ibm_container_nlb_dns":                 kubernetes.DataSourceIBMContainerNLBDNSValidator(),
				"ibm_container_storage_attachment":      kubernetes.DataSourceIBMContainerVpcWorkerVolumeAttachmentValidator(),
				"ibm_container_vpc_cluster_worker_pool": kubernetes.DataSourceIBMContainerVpcClusterWorkerPoolValidator(),
				"ibm_container_worker_pool":             kubernetes.DataSourceIBMContainerWorkerPoolValidator(),
				"ibm_container_bind_service":            kubernetes.DataSourceIBMContainerBindServiceValidator(),
				"ibm_container_cluster_config":          kubernetes.DataSourceIBMContainerClusterConfigValidator(),
				"ibm_container_cluster":                 kubernetes.DataSourceIBMContainerClusterValidator(),
				"ibm_container_vpc_cluster_worker":      kubernetes.DataSourceIBMContainerVPCClusterWorkerValidator(),
				"ibm_container_vpc_cluster":             kubernetes.DataSourceIBMContainerVPCClusterValidator(),
				"ibm_container_alb_cert":                kubernetes.DataSourceIBMContainerALBCertValidator(),
				"ibm_container_ingress_instance":        kubernetes.DataSourceIBMContainerIngressInstanceValidator(),
				"ibm_container_ingress_secret_tls":      kubernetes.DataSourceIBMContainerIngressSecretTLSValidator(),
				"ibm_container_ingress_secret_opaque":   kubernetes.DataSourceIBMContainerIngressSecretOpaqueValidator(),

				"ibm_iam_access_group": iamaccessgroup.DataSourceIBMIAMAccessGroupValidator(),

				"ibm_iam_service_id":                  iamidentity.DataSourceIBMIAMServiceIDValidator(),
				"ibm_iam_trusted_profile_claim_rule":  iamidentity.DataSourceIBMIamTrustedProfileClaimRuleValidator(),
				"ibm_iam_trusted_profile_link":        iamidentity.DataSourceIBMIamTrustedProfileLinkValidator(),
				"ibm_iam_trusted_profile_links":       iamidentity.DataSourceIBMIamTrustedProfileLinksValidator(),
				"ibm_iam_trusted_profile":             iamidentity.DataSourceIBMIamTrustedProfileValidator(),
				"ibm_iam_trusted_profile_claim_rules": iamidentity.DataSourceIBMIamTrustedProfileClaimRulesValidator(),
				"ibm_iam_trusted_profiles":            iamidentity.DataSourceIBMIamTrustedProfilesValidator(),

				"ibm_iam_access_group_policy":    iampolicy.DataSourceIBMIAMAccessGroupPolicyValidator(),
				"ibm_iam_service_policy":         iampolicy.DataSourceIBMIAMServicePolicyValidator(),
				"ibm_iam_trusted_profile_policy": iampolicy.DataSourceIBMIAMTrustedProfilePolicyValidator(),
			},
		}
	})
	return globalValidatorDict
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	var bluemixAPIKey string
	var bluemixTimeout int
	var iamToken, iamRefreshToken, iamTrustedProfileId string
	if key, ok := d.GetOk("bluemix_api_key"); ok {
		bluemixAPIKey = key.(string)
	}
	if key, ok := d.GetOk("ibmcloud_api_key"); ok {
		bluemixAPIKey = key.(string)
	}
	if itoken, ok := d.GetOk("iam_token"); ok {
		iamToken = itoken.(string)
	}
	if rtoken, ok := d.GetOk("iam_refresh_token"); ok {
		iamRefreshToken = rtoken.(string)
	}
	if ttoken, ok := d.GetOk("iam_profile_id"); ok {
		iamTrustedProfileId = ttoken.(string)
	}
	var softlayerUsername, softlayerAPIKey, softlayerEndpointUrl string
	var softlayerTimeout int
	if username, ok := d.GetOk("softlayer_username"); ok {
		softlayerUsername = username.(string)
	}
	if username, ok := d.GetOk("iaas_classic_username"); ok {
		softlayerUsername = username.(string)
	}
	if apikey, ok := d.GetOk("softlayer_api_key"); ok {
		softlayerAPIKey = apikey.(string)
	}
	if apikey, ok := d.GetOk("iaas_classic_api_key"); ok {
		softlayerAPIKey = apikey.(string)
	}
	if endpoint, ok := d.GetOk("softlayer_endpoint_url"); ok {
		softlayerEndpointUrl = endpoint.(string)
	}
	if endpoint, ok := d.GetOk("iaas_classic_endpoint_url"); ok {
		softlayerEndpointUrl = endpoint.(string)
	}
	if tm, ok := d.GetOk("softlayer_timeout"); ok {
		softlayerTimeout = tm.(int)
	}
	if tm, ok := d.GetOk("iaas_classic_timeout"); ok {
		softlayerTimeout = tm.(int)
	}

	if tm, ok := d.GetOk("bluemix_timeout"); ok {
		bluemixTimeout = tm.(int)
	}
	if tm, ok := d.GetOk("ibmcloud_timeout"); ok {
		bluemixTimeout = tm.(int)
	}
	var visibility string
	if v, ok := d.GetOk("visibility"); ok {
		visibility = v.(string)
	}
	var file string
	if f, ok := d.GetOk("endpoints_file_path"); ok {
		file = f.(string)
	}

	resourceGrp := d.Get("resource_group").(string)
	region := d.Get("region").(string)
	zone := d.Get("zone").(string)
	retryCount := d.Get("max_retries").(int)
	wskNameSpace := d.Get("function_namespace").(string)
	riaasEndPoint := d.Get("riaas_endpoint").(string)

	wskEnvVal, err := schema.EnvDefaultFunc("FUNCTION_NAMESPACE", "")()
	if err != nil {
		return nil, err
	}
	//Set environment variable to be used in DiffSupressFunction
	if wskEnvVal.(string) == "" {
		os.Setenv("FUNCTION_NAMESPACE", wskNameSpace)
	}

	config := conns.Config{
		BluemixAPIKey:        bluemixAPIKey,
		Region:               region,
		ResourceGroup:        resourceGrp,
		BluemixTimeout:       time.Duration(bluemixTimeout) * time.Second,
		SoftLayerTimeout:     time.Duration(softlayerTimeout) * time.Second,
		SoftLayerUserName:    softlayerUsername,
		SoftLayerAPIKey:      softlayerAPIKey,
		RetryCount:           retryCount,
		SoftLayerEndpointURL: softlayerEndpointUrl,
		RetryDelay:           conns.RetryAPIDelay,
		FunctionNameSpace:    wskNameSpace,
		RiaasEndPoint:        riaasEndPoint,
		IAMToken:             iamToken,
		IAMRefreshToken:      iamRefreshToken,
		Zone:                 zone,
		Visibility:           visibility,
		EndpointsFile:        file,
		IAMTrustedProfileID:  iamTrustedProfileId,
	}

	return config.ClientSession()
}
