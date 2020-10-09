package ibm

import (
	"os"
	"sync"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/mutexkv"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// This is a global MutexKV for use within this plugin.
var ibmMutexKV = mutexkv.NewMutexKV()

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
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
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"IC_GENERATION", "IBMCLOUD_GENERATION"}, 2),
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
		},

		DataSourcesMap: map[string]*schema.Resource{
			"ibm_api_gateway":                        dataSourceIBMApiGateway(),
			"ibm_account":                            dataSourceIBMAccount(),
			"ibm_app":                                dataSourceIBMApp(),
			"ibm_app_domain_private":                 dataSourceIBMAppDomainPrivate(),
			"ibm_app_domain_shared":                  dataSourceIBMAppDomainShared(),
			"ibm_app_route":                          dataSourceIBMAppRoute(),
			"ibm_function_action":                    dataSourceIBMFunctionAction(),
			"ibm_function_package":                   dataSourceIBMFunctionPackage(),
			"ibm_function_rule":                      dataSourceIBMFunctionRule(),
			"ibm_function_trigger":                   dataSourceIBMFunctionTrigger(),
			"ibm_function_namespace":                 dataSourceIBMFunctionNamespace(),
			"ibm_certificate_manager_certificates":   dataIBMCertificateManagerCertificates(),
			"ibm_certificate_manager_certificate":    dataIBMCertificateManagerCertificate(),
			"ibm_cis":                                dataSourceIBMCISInstance(),
			"ibm_cis_dns_records":                    dataSourceIBMCISDNSRecords(),
			"ibm_cis_healthchecks":                   dataSourceIBMCISHealthChecks(),
			"ibm_cis_domain":                         dataSourceIBMCISDomain(),
			"ibm_cis_firewall":                       dataIBMCISFirewallRecord(),
			"ibm_cis_rate_limit":                     dataSourceIBMCISRateLimit(),
			"ibm_cis_ip_addresses":                   dataSourceIBMCISIP(),
			"ibm_cis_edge_functions_actions":         dataSourceIBMCISEdgeFunctionsActions(),
			"ibm_cis_edge_functions_triggers":        dataSourceIBMCISEdgeFunctionsTriggers(),
			"ibm_database":                           dataSourceIBMDatabaseInstance(),
			"ibm_compute_bare_metal":                 dataSourceIBMComputeBareMetal(),
			"ibm_compute_image_template":             dataSourceIBMComputeImageTemplate(),
			"ibm_compute_placement_group":            dataSourceIBMComputePlacementGroup(),
			"ibm_compute_ssh_key":                    dataSourceIBMComputeSSHKey(),
			"ibm_compute_vm_instance":                dataSourceIBMComputeVmInstance(),
			"ibm_container_addons":                   datasourceIBMContainerAddOns(),
			"ibm_container_alb":                      dataSourceIBMContainerALB(),
			"ibm_container_alb_cert":                 dataSourceIBMContainerALBCert(),
			"ibm_container_bind_service":             dataSourceIBMContainerBindService(),
			"ibm_container_cluster":                  dataSourceIBMContainerCluster(),
			"ibm_container_cluster_config":           dataSourceIBMContainerClusterConfig(),
			"ibm_container_cluster_versions":         dataSourceIBMContainerClusterVersions(),
			"ibm_container_cluster_worker":           dataSourceIBMContainerClusterWorker(),
			"ibm_container_vpc_cluster_alb":          dataSourceIBMContainerVPCClusterALB(),
			"ibm_container_vpc_alb":                  dataSourceIBMContainerVPCClusterALB(),
			"ibm_container_vpc_cluster":              dataSourceIBMContainerVPCCluster(),
			"ibm_container_vpc_cluster_worker":       dataSourceIBMContainerVPCClusterWorker(),
			"ibm_container_vpc_cluster_worker_pool":  dataSourceIBMContainerVpcClusterWorkerPool(),
			"ibm_container_vpc_worker_pool":          dataSourceIBMContainerVpcClusterWorkerPool(),
			"ibm_container_worker_pool":              dataSourceIBMContainerWorkerPool(),
			"ibm_cos_bucket":                         dataSourceIBMCosBucket(),
			"ibm_dns_domain_registration":            dataSourceIBMDNSDomainRegistration(),
			"ibm_dns_domain":                         dataSourceIBMDNSDomain(),
			"ibm_dns_secondary":                      dataSourceIBMDNSSecondary(),
			"ibm_event_streams_topic":                dataSourceIBMEventStreamsTopic(),
			"ibm_iam_access_group":                   dataSourceIBMIAMAccessGroup(),
			"ibm_iam_auth_token":                     dataSourceIBMIAMAuthToken(),
			"ibm_iam_role_actions":                   datasourceIBMIAMRoleAction(),
			"ibm_iam_users":                          dataSourceIBMIAMUsers(),
			"ibm_iam_roles":                          datasourceIBMIAMRole(),
			"ibm_iam_user_policy":                    dataSourceIBMIAMUserPolicy(),
			"ibm_iam_user_profile":                   dataSourceIBMIAMUserProfile(),
			"ibm_iam_service_id":                     dataSourceIBMIAMServiceID(),
			"ibm_iam_service_policy":                 dataSourceIBMIAMServicePolicy(),
			"ibm_is_floating_ip":                     dataSourceIBMISFloatingIP(),
			"ibm_is_flow_logs":                       dataSourceIBMISFlowLogs(),
			"ibm_is_image":                           dataSourceIBMISImage(),
			"ibm_is_images":                          dataSourceIBMISImages(),
			"ibm_is_instance_group":                  dataSourceIBMISInstanceGroup(),
			"ibm_is_instance_group_manager":          dataSourceIBMISInstanceGroupManager(),
			"ibm_is_instance_group_managers":         dataSourceIBMISInstanceGroupManagers(),
			"ibm_is_instance_group_manager_policies": dataSourceIBMISInstanceGroupManagerPolicies(),
			"ibm_is_instance_group_manager_policy":   dataSourceIBMISInstanceGroupManagerPolicy(),
			"ibm_is_instance_templates":              dataSourceIBMISInstanceTemplates(),
			"ibm_is_instance_profile":                dataSourceIBMISInstanceProfile(),
			"ibm_is_instance_profiles":               dataSourceIBMISInstanceProfiles(),
			"ibm_is_instance":                        dataSourceIBMISInstance(),
			"ibm_is_instances":                       dataSourceIBMISInstances(),
			"ibm_is_lb":                              dataSourceIBMISLB(),
			"ibm_is_lb_profiles":                     dataSourceIBMISLbProfiles(),
			"ibm_is_lbs":                             dataSourceIBMISLBS(),
			"ibm_is_public_gateway":                  dataSourceIBMISPublicGateway(),
			"ibm_is_region":                          dataSourceIBMISRegion(),
			"ibm_is_ssh_key":                         dataSourceIBMISSSHKey(),
			"ibm_is_subnet":                          dataSourceIBMISSubnet(),
			"ibm_is_subnets":                         dataSourceIBMISSubnets(),
			"ibm_is_security_group":                  dataSourceIBMISSecurityGroup(),
			"ibm_is_volume":                          dataSourceIBMISVolume(),
			"ibm_is_vpc":                             dataSourceIBMISVPC(),
			"ibm_is_zone":                            dataSourceIBMISZone(),
			"ibm_is_zones":                           dataSourceIBMISZones(),
			"ibm_lbaas":                              dataSourceIBMLbaas(),
			"ibm_network_vlan":                       dataSourceIBMNetworkVlan(),
			"ibm_org":                                dataSourceIBMOrg(),
			"ibm_org_quota":                          dataSourceIBMOrgQuota(),
			"ibm_kp_key":                             dataSourceIBMkey(),
			"ibm_kms_keys":                           dataSourceIBMKMSkeys(),
			"ibm_kms_key":                            dataSourceIBMKMSkey(),
			"ibm_resource_quota":                     dataSourceIBMResourceQuota(),
			"ibm_resource_group":                     dataSourceIBMResourceGroup(),
			"ibm_resource_instance":                  dataSourceIBMResourceInstance(),
			"ibm_resource_key":                       dataSourceIBMResourceKey(),
			"ibm_security_group":                     dataSourceIBMSecurityGroup(),
			"ibm_service_instance":                   dataSourceIBMServiceInstance(),
			"ibm_service_key":                        dataSourceIBMServiceKey(),
			"ibm_service_plan":                       dataSourceIBMServicePlan(),
			"ibm_space":                              dataSourceIBMSpace(),
			"ibm_schematics_workspace":               dataSourceSchematicsWorkspace(),
			"ibm_schematics_output":                  dataSourceSchematicsOut(),
			"ibm_schematics_state":                   dataSourceSchematicsState(),
			// Added for Power Resources

			"ibm_pi_key":                dataSourceIBMPIKey(),
			"ibm_pi_image":              dataSourceIBMPIImage(),
			"ibm_pi_instance":           dataSourceIBMPIInstance(),
			"ibm_pi_tenant":             dataSourceIBMPITenant(),
			"ibm_pi_network":            dataSourceIBMPINetwork(),
			"ibm_pi_volume":             dataSourceIBMPIVolume(),
			"ibm_pi_instance_volumes":   dataSourceIBMPIVolumes(),
			"ibm_pi_public_network":     dataSourceIBMPIPublicNetwork(),
			"ibm_pi_images":             dataSourceIBMPIImages(),
			"ibm_pi_instance_ip":        dataSourceIBMPIInstanceIP(),
			"ibm_pi_instance_snapshots": dataSourceIBMPISnapshots(),
			"ibm_pi_pvm_snapshots":      dataSourceIBMPISnapshot(),
			"ibm_pi_network_port":       dataSourceIBMPINetworkPort(),
			"ibm_pi_cloud_instance":     dataSourceIBMPICloudInstance(),

			// Added for private dns zones

			"ibm_dns_zones":              dataSourceIBMPrivateDNSZones(),
			"ibm_dns_permitted_networks": dataSourceIBMPrivateDNSPermittedNetworks(),
			"ibm_dns_resource_records":   dataSourceIBMPrivateDNSResourceRecords(),

			// Added for Direct Link

			"ibm_dl_gateways":        dataSourceIBMDLGateways(),
			"ibm_dl_offering_speeds": dataSourceIBMDLOfferingSpeeds(),
			"ibm_dl_port":            dataSourceIBMDirectLinkPort(),
			"ibm_dl_ports":           dataSourceIBMDirectLinkPorts(),
			"ibm_dl_gateway":         dataSourceIBMDLGateway(),
			"ibm_dl_locations":       dataSourceIBMDLLocations(),
			"ibm_dl_routers":         dataSourceIBMDLRouters(),

			//Added for Transit Gateway
			"ibm_tg_gateway":   dataSourceIBMTransitGateway(),
			"ibm_tg_gateways":  dataSourceIBMTransitGateways(),
			"ibm_tg_locations": dataSourceIBMTransitGatewaysLocations(),
			"ibm_tg_location":  dataSourceIBMTransitGatewaysLocation(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"ibm_api_gateway_endpoint":                           resourceIBMApiGatewayEndPoint(),
			"ibm_api_gateway_endpoint_subscription":              resourceIBMApiGatewayEndpointSubscription(),
			"ibm_app":                                            resourceIBMApp(),
			"ibm_app_domain_private":                             resourceIBMAppDomainPrivate(),
			"ibm_app_domain_shared":                              resourceIBMAppDomainShared(),
			"ibm_app_route":                                      resourceIBMAppRoute(),
			"ibm_function_action":                                resourceIBMFunctionAction(),
			"ibm_function_package":                               resourceIBMFunctionPackage(),
			"ibm_function_rule":                                  resourceIBMFunctionRule(),
			"ibm_function_trigger":                               resourceIBMFunctionTrigger(),
			"ibm_function_namespace":                             resourceIBMFunctionNamespace(),
			"ibm_cis":                                            resourceIBMCISInstance(),
			"ibm_database":                                       resourceIBMDatabaseInstance(),
			"ibm_certificate_manager_import":                     resourceIBMCertificateManagerImport(),
			"ibm_certificate_manager_order":                      resourceIBMCertificateManagerOrder(),
			"ibm_cis_domain":                                     resourceIBMCISDomain(),
			"ibm_cis_domain_settings":                            resourceIBMCISSettings(),
			"ibm_cis_firewall":                                   resourceIBMCISFirewallRecord(),
			"ibm_cis_healthcheck":                                resourceIBMCISHealthCheck(),
			"ibm_cis_origin_pool":                                resourceIBMCISPool(),
			"ibm_cis_global_load_balancer":                       resourceIBMCISGlb(),
			"ibm_cis_dns_record":                                 resourceIBMCISDnsRecord(),
			"ibm_cis_rate_limit":                                 resourceIBMCISRateLimit(),
			"ibm_cis_edge_functions_action":                      resourceIBMCISEdgeFunctionsAction(),
			"ibm_cis_edge_functions_trigger":                     resourceIBMCISEdgeFunctionsTrigger(),
			"ibm_cis_tls_settings":                               resourceIBMCISTLSSettings(),
			"ibm_compute_autoscale_group":                        resourceIBMComputeAutoScaleGroup(),
			"ibm_compute_autoscale_policy":                       resourceIBMComputeAutoScalePolicy(),
			"ibm_compute_bare_metal":                             resourceIBMComputeBareMetal(),
			"ibm_compute_dedicated_host":                         resourceIBMComputeDedicatedHost(),
			"ibm_compute_monitor":                                resourceIBMComputeMonitor(),
			"ibm_compute_placement_group":                        resourceIBMComputePlacementGroup(),
			"ibm_compute_provisioning_hook":                      resourceIBMComputeProvisioningHook(),
			"ibm_compute_ssh_key":                                resourceIBMComputeSSHKey(),
			"ibm_compute_ssl_certificate":                        resourceIBMComputeSSLCertificate(),
			"ibm_compute_user":                                   resourceIBMComputeUser(),
			"ibm_compute_vm_instance":                            resourceIBMComputeVmInstance(),
			"ibm_container_addons":                               resourceIBMContainerAddOns(),
			"ibm_container_alb":                                  resourceIBMContainerALB(),
			"ibm_container_vpc_alb":                              resourceIBMContainerVpcALB(),
			"ibm_container_vpc_worker_pool":                      resourceIBMContainerVpcWorkerPool(),
			"ibm_container_vpc_cluster":                          resourceIBMContainerVpcCluster(),
			"ibm_container_alb_cert":                             resourceIBMContainerALBCert(),
			"ibm_container_cluster":                              resourceIBMContainerCluster(),
			"ibm_container_cluster_feature":                      resourceIBMContainerClusterFeature(),
			"ibm_container_bind_service":                         resourceIBMContainerBindService(),
			"ibm_container_worker_pool":                          resourceIBMContainerWorkerPool(),
			"ibm_container_worker_pool_zone_attachment":          resourceIBMContainerWorkerPoolZoneAttachment(),
			"ibm_cos_bucket":                                     resourceIBMCOS(),
			"ibm_dns_domain":                                     resourceIBMDNSDomain(),
			"ibm_dns_domain_registration_nameservers":            resourceIBMDNSDomainRegistrationNameservers(),
			"ibm_dns_secondary":                                  resourceIBMDNSSecondary(),
			"ibm_dns_record":                                     resourceIBMDNSRecord(),
			"ibm_event_streams_topic":                            resourceIBMEventStreamsTopic(),
			"ibm_firewall":                                       resourceIBMFirewall(),
			"ibm_firewall_policy":                                resourceIBMFirewallPolicy(),
			"ibm_iam_access_group":                               resourceIBMIAMAccessGroup(),
			"ibm_iam_custom_role":                                resourceIBMIAMCustomRole(),
			"ibm_iam_access_group_dynamic_rule":                  resourceIBMIAMDynamicRule(),
			"ibm_iam_access_group_members":                       resourceIBMIAMAccessGroupMembers(),
			"ibm_iam_access_group_policy":                        resourceIBMIAMAccessGroupPolicy(),
			"ibm_iam_authorization_policy":                       resourceIBMIAMAuthorizationPolicy(),
			"ibm_iam_authorization_policy_detach":                resourceIBMIAMAuthorizationPolicyDetach(),
			"ibm_iam_user_policy":                                resourceIBMIAMUserPolicy(),
			"ibm_iam_user_settings":                              resourceIBMUserSettings(),
			"ibm_iam_service_id":                                 resourceIBMIAMServiceID(),
			"ibm_iam_service_policy":                             resourceIBMIAMServicePolicy(),
			"ibm_iam_user_invite":                                resourceIBMUserInvite(),
			"ibm_ipsec_vpn":                                      resourceIBMIPSecVPN(),
			"ibm_is_floating_ip":                                 resourceIBMISFloatingIP(),
			"ibm_is_flow_log":                                    resourceIBMISFlowLog(),
			"ibm_is_instance":                                    resourceIBMISInstance(),
			"ibm_is_instance_group":                              resourceIBMISInstanceGroup(),
			"ibm_is_instance_group_manager":                      resourceIBMISInstanceGroupManager(),
			"ibm_is_instance_group_manager_policy":               resourceIBMISInstanceGroupManagerPolicy(),
			"ibm_is_instance_template":                           resourceIBMISInstanceTemplate(),
			"ibm_is_ike_policy":                                  resourceIBMISIKEPolicy(),
			"ibm_is_ipsec_policy":                                resourceIBMISIPSecPolicy(),
			"ibm_is_lb":                                          resourceIBMISLB(),
			"ibm_is_lb_listener":                                 resourceIBMISLBListener(),
			"ibm_is_lb_listener_policy":                          resourceIBMISLBListenerPolicy(),
			"ibm_is_lb_listener_policy_rule":                     resourceIBMISLBListenerPolicyRule(),
			"ibm_is_lb_pool":                                     resourceIBMISLBPool(),
			"ibm_is_lb_pool_member":                              resourceIBMISLBPoolMember(),
			"ibm_is_network_acl":                                 resourceIBMISNetworkACL(),
			"ibm_is_public_gateway":                              resourceIBMISPublicGateway(),
			"ibm_is_security_group":                              resourceIBMISSecurityGroup(),
			"ibm_is_security_group_rule":                         resourceIBMISSecurityGroupRule(),
			"ibm_is_security_group_network_interface_attachment": resourceIBMISSecurityGroupNetworkInterfaceAttachment(),
			"ibm_is_subnet":                                      resourceIBMISSubnet(),
			"ibm_is_subnet_network_acl_attachment":               resourceIBMISSubnetNetworkACLAttachment(),
			"ibm_is_ssh_key":                                     resourceIBMISSSHKey(),
			"ibm_is_volume":                                      resourceIBMISVolume(),
			"ibm_is_vpn_gateway":                                 resourceIBMISVPNGateway(),
			"ibm_is_vpn_gateway_connection":                      resourceIBMISVPNGatewayConnection(),
			"ibm_is_vpc":                                         resourceIBMISVPC(),
			"ibm_is_vpc_address_prefix":                          resourceIBMISVpcAddressPrefix(),
			"ibm_is_vpc_route":                                   resourceIBMISVpcRoute(),
			"ibm_is_image":                                       resourceIBMISImage(),
			"ibm_lb":                                             resourceIBMLb(),
			"ibm_lbaas":                                          resourceIBMLbaas(),
			"ibm_lbaas_health_monitor":                           resourceIBMLbaasHealthMonitor(),
			"ibm_lbaas_server_instance_attachment":               resourceIBMLbaasServerInstanceAttachment(),
			"ibm_lb_service":                                     resourceIBMLbService(),
			"ibm_lb_service_group":                               resourceIBMLbServiceGroup(),
			"ibm_lb_vpx":                                         resourceIBMLbVpx(),
			"ibm_lb_vpx_ha":                                      resourceIBMLbVpxHa(),
			"ibm_lb_vpx_service":                                 resourceIBMLbVpxService(),
			"ibm_lb_vpx_vip":                                     resourceIBMLbVpxVip(),
			"ibm_multi_vlan_firewall":                            resourceIBMMultiVlanFirewall(),
			"ibm_network_gateway":                                resourceIBMNetworkGateway(),
			"ibm_network_gateway_vlan_association":               resourceIBMNetworkGatewayVlanAttachment(),
			"ibm_network_interface_sg_attachment":                resourceIBMNetworkInterfaceSGAttachment(),
			"ibm_network_public_ip":                              resourceIBMNetworkPublicIp(),
			"ibm_network_vlan":                                   resourceIBMNetworkVlan(),
			"ibm_network_vlan_spanning":                          resourceIBMNetworkVlanSpan(),
			"ibm_object_storage_account":                         resourceIBMObjectStorageAccount(),
			"ibm_org":                                            resourceIBMOrg(),
			"ibm_kms_key":                                        resourceIBMKmskey(),
			"ibm_kp_key":                                         resourceIBMkey(),
			"ibm_resource_group":                                 resourceIBMResourceGroup(),
			"ibm_resource_instance":                              resourceIBMResourceInstance(),
			"ibm_resource_key":                                   resourceIBMResourceKey(),
			"ibm_security_group":                                 resourceIBMSecurityGroup(),
			"ibm_security_group_rule":                            resourceIBMSecurityGroupRule(),
			"ibm_service_instance":                               resourceIBMServiceInstance(),
			"ibm_service_key":                                    resourceIBMServiceKey(),
			"ibm_space":                                          resourceIBMSpace(),
			"ibm_storage_evault":                                 resourceIBMStorageEvault(),
			"ibm_storage_block":                                  resourceIBMStorageBlock(),
			"ibm_storage_file":                                   resourceIBMStorageFile(),
			"ibm_subnet":                                         resourceIBMSubnet(),
			"ibm_dns_reverse_record":                             resourceIBMDNSReverseRecord(),
			"ibm_ssl_certificate":                                resourceIBMSSLCertificate(),
			"ibm_cdn":                                            resourceIBMCDN(),
			"ibm_hardware_firewall_shared":                       resourceIBMFirewallShared(),

			//Added for Power Colo

			"ibm_pi_key":                 resourceIBMPIKey(),
			"ibm_pi_volume":              resourceIBMPIVolume(),
			"ibm_pi_network":             resourceIBMPINetwork(),
			"ibm_pi_instance":            resourceIBMPIInstance(),
			"ibm_pi_operations":          resourceIBMPIIOperations(),
			"ibm_pi_volume_attach":       resourceIBMPIVolumeAttach(),
			"ibm_pi_capture":             resourceIBMPICapture(),
			"ibm_pi_image":               resourceIBMPIImage(),
			"ibm_pi_network_port":        resourceIBMPINetworkPort(),
			"ibm_pi_snapshot":            resourceIBMPISnapshot(),
			"ibm_pi_network_port_attach": resourceIBMPINetworkPortAttach(),

			//Private DNS related resources
			"ibm_dns_zone":              resourceIBMPrivateDNSZone(),
			"ibm_dns_permitted_network": resourceIBMPrivateDNSPermittedNetwork(),
			"ibm_dns_resource_record":   resourceIBMPrivateDNSResourceRecord(),

			//Direct Link related resources
			"ibm_dl_gateway":            resourceIBMDLGateway(),
			"ibm_dl_virtual_connection": resourceIBMDLGatewayVC(),
			//Added for Transit Gateway
			"ibm_tg_gateway":    resourceIBMTransitGateway(),
			"ibm_tg_connection": resourceIBMTransitGatewayConnection(),
		},

		ConfigureFunc: providerConfigure,
	}
}

var globalValidatorDict ValidatorDict
var initOnce sync.Once

// Validator return validator
func Validator() ValidatorDict {
	initOnce.Do(func() {
		globalValidatorDict = ValidatorDict{
			ResourceValidatorDictionary: map[string]*ResourceValidator{
				"ibm_is_vpc":                           resourceIBMISVPCValidator(),
				"ibm_is_ike_policy":                    resourceIBMISIKEValidator(),
				"ibm_is_network_acl":                   resourceIBMISNetworkACLValidator(),
				"ibm_iam_custom_role":                  resourceIBMIAMCustomRoleValidator(),
				"ibm_cis_healthcheck":                  resourceIBMCISHealthCheckValidator(),
				"ibm_cis_rate_limit":                   resourceIBMCISRateLimitValidator(),
				"ibm_cis_domain_settings":              resourceIBMCISDomainSettingValidator(),
				"ibm_cis_tls_settings":                 resourceIBMCISTLSSettingsValidator(),
				"ibm_tg_gateway":                       resourceIBMTGValidator(),
				"ibm_tg_connection":                    resourceIBMTransitGatewayConnectionValidator(),
				"ibm_dl_virtual_connection":            resourceIBMdlGatewayVCValidator(),
				"ibm_dl_gateway":                       resourceIBMDLGatewayValidator(),
				"ibm_function_package":                 resourceIBMFuncPackageValidator(),
				"ibm_function_action":                  resourceIBMFuncActionValidator(),
				"ibm_function_rule":                    resourceIBMFuncRuleValidator(),
				"ibm_function_trigger":                 resourceIBMFuncTriggerValidator(),
				"ibm_function_namespace":               resourceIBMFuncNamespaceValidator(),
				"ibm_is_flow_log":                      resourceIBMISFlowLogValidator(),
				"ibm_is_instance_group":                resourceIBMISInstanceGroupValidator(),
				"ibm_is_instance_group_manager":        resourceIBMISInstanceGroupManagerValidator(),
				"ibm_is_instance_group_manager_policy": resourceIBMISInstanceGroupManagerPolicyValidator(),
				"ibm_is_lb":                            resourceIBMISLBValidator(),
			},
			DataSourceValidatorDictionary: map[string]*ResourceValidator{
				"ibm_is_subnet":          dataSourceIBMISSubnetValidator(),
				"ibm_dl_offering_speeds": datasourceIBMDLOfferingSpeedsValidator(),
				"ibm_dl_routers":         datasourceIBMDLRoutersValidator(),
				"ibm_is_vpc":             dataSourceIBMISVpcValidator(),
				"ibm_is_volume":          dataSourceIBMISVolumeValidator(),
			},
		}
	})
	return globalValidatorDict
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	var bluemixAPIKey string
	var bluemixTimeout int
	var iamToken, iamRefreshToken string
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

	resourceGrp := d.Get("resource_group").(string)
	region := d.Get("region").(string)
	zone := d.Get("zone").(string)
	retryCount := d.Get("max_retries").(int)
	wskNameSpace := d.Get("function_namespace").(string)
	riaasEndPoint := d.Get("riaas_endpoint").(string)
	generation := d.Get("generation").(int)

	wskEnvVal, err := schema.EnvDefaultFunc("FUNCTION_NAMESPACE", "")()
	if err != nil {
		return nil, err
	}
	//Set environment variable to be used in DiffSupressFunction
	if wskEnvVal.(string) == "" {
		os.Setenv("FUNCTION_NAMESPACE", wskNameSpace)
	}

	config := Config{
		BluemixAPIKey:        bluemixAPIKey,
		Region:               region,
		ResourceGroup:        resourceGrp,
		BluemixTimeout:       time.Duration(bluemixTimeout) * time.Second,
		SoftLayerTimeout:     time.Duration(softlayerTimeout) * time.Second,
		SoftLayerUserName:    softlayerUsername,
		SoftLayerAPIKey:      softlayerAPIKey,
		RetryCount:           retryCount,
		SoftLayerEndpointURL: softlayerEndpointUrl,
		RetryDelay:           RetryAPIDelay,
		FunctionNameSpace:    wskNameSpace,
		RiaasEndPoint:        riaasEndPoint,
		Generation:           generation,
		IAMToken:             iamToken,
		IAMRefreshToken:      iamRefreshToken,
		Zone:                 zone,
		//PowerServiceInstance: powerServiceInstance,
	}

	return config.ClientSession()
}
