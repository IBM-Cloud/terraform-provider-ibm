package ibm

import (
	"os"
	"time"

	"github.com/hashicorp/terraform/helper/mutexkv"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
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
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"BM_API_KEY", "BLUEMIX_API_KEY"}, ""),
			},
			"bluemix_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The timeout (in seconds) to set for any Bluemix API calls made.",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"BM_TIMEOUT", "BLUEMIX_TIMEOUT"}, 60),
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Bluemix Region (for example 'us-south').",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"BM_REGION", "BLUEMIX_REGION"}, "us-south"),
			},
			"softlayer_api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The SoftLayer API Key",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"SL_API_KEY", "SOFTLAYER_API_KEY"}, ""),
			},
			"softlayer_username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The SoftLayer user name",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"SL_USERNAME", "SOFTLAYER_USERNAME"}, ""),
			},
			"softlayer_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The timeout (in seconds) to set for any SoftLayer API calls made.",
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{"SL_TIMEOUT", "SOFTLAYER_TIMEOUT"}, 60),
			},
			"max_retries": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The retry count to set for any SoftLayer API calls.",
				DefaultFunc: schema.EnvDefaultFunc("MAX_RETRIES", 5),
			},
			"function_namespace": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The IBM Cloud Function namespace",
				DefaultFunc: schema.EnvDefaultFunc("FUNCTION_NAMESPACE", ""),
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"ibm_account":                    dataSourceIBMAccount(),
			"ibm_app":                        dataSourceIBMApp(),
			"ibm_app_domain_private":         dataSourceIBMAppDomainPrivate(),
			"ibm_app_domain_shared":          dataSourceIBMAppDomainShared(),
			"ibm_app_route":                  dataSourceIBMAppRoute(),
			"ibm_function_action":            dataSourceIBMFunctionAction(),
			"ibm_function_package":           dataSourceIBMFunctionPackage(),
			"ibm_function_rule":              dataSourceIBMFunctionRule(),
			"ibm_function_trigger":           dataSourceIBMFunctionTrigger(),
			"ibm_compute_bare_metal":         dataSourceIBMComputeBareMetal(),
			"ibm_compute_image_template":     dataSourceIBMComputeImageTemplate(),
			"ibm_compute_ssh_key":            dataSourceIBMComputeSSHKey(),
			"ibm_compute_vm_instance":        dataSourceIBMComputeVmInstance(),
			"ibm_container_cluster":          dataSourceIBMContainerCluster(),
			"ibm_container_cluster_config":   dataSourceIBMContainerClusterConfig(),
			"ibm_container_cluster_versions": dataSourceIBMContainerClusterVersions(),
			"ibm_container_cluster_worker":   dataSourceIBMContainerClusterWorker(),
			"ibm_dns_domain":                 dataSourceIBMDNSDomain(),
			"ibm_dns_secondary":              dataSourceIBMDNSSecondary(),
			"ibm_iam_user_policy":            dataSourceIBMIAMUserPolicy(),
			"ibm_lbaas":                      dataSourceIBMLbaas(),
			"ibm_network_vlan":               dataSourceIBMNetworkVlan(),
			"ibm_org":                        dataSourceIBMOrg(),
			"ibm_org_quota":                  dataSourceIBMOrgQuota(),
			"ibm_resource_group":             dataSourceIBMResourceGroup(),
			"ibm_resource_instance":          dataSourceIBMResourceInstance(),
			"ibm_resource_key":               dataSourceIBMResourceKey(),
			"ibm_security_group":             dataSourceIBMSecurityGroup(),
			"ibm_service_instance":           dataSourceIBMServiceInstance(),
			"ibm_service_key":                dataSourceIBMServiceKey(),
			"ibm_service_plan":               dataSourceIBMServicePlan(),
			"ibm_space":                      dataSourceIBMSpace(),
		},

		ResourcesMap: map[string]*schema.Resource{

			"ibm_app":                              resourceIBMApp(),
			"ibm_app_domain_private":               resourceIBMAppDomainPrivate(),
			"ibm_app_domain_shared":                resourceIBMAppDomainShared(),
			"ibm_app_route":                        resourceIBMAppRoute(),
			"ibm_function_action":                  resourceIBMFunctionAction(),
			"ibm_function_package":                 resourceIBMFunctionPackage(),
			"ibm_function_rule":                    resourceIBMFunctionRule(),
			"ibm_function_trigger":                 resourceIBMFunctionTrigger(),
			"ibm_compute_autoscale_group":          resourceIBMComputeAutoScaleGroup(),
			"ibm_compute_autoscale_policy":         resourceIBMComputeAutoScalePolicy(),
			"ibm_compute_bare_metal":               resourceIBMComputeBareMetal(),
			"ibm_compute_dedicated_host":           resourceIBMComputeDedicatedHost(),
			"ibm_compute_monitor":                  resourceIBMComputeMonitor(),
			"ibm_compute_provisioning_hook":        resourceIBMComputeProvisioningHook(),
			"ibm_compute_ssh_key":                  resourceIBMComputeSSHKey(),
			"ibm_compute_ssl_certificate":          resourceIBMComputeSSLCertificate(),
			"ibm_compute_user":                     resourceIBMComputeUser(),
			"ibm_compute_vm_instance":              resourceIBMComputeVmInstance(),
			"ibm_container_cluster":                resourceIBMContainerCluster(),
			"ibm_container_bind_service":           resourceIBMContainerBindService(),
			"ibm_dns_domain":                       resourceIBMDNSDomain(),
			"ibm_dns_secondary":                    resourceIBMDNSSecondary(),
			"ibm_dns_record":                       resourceIBMDNSRecord(),
			"ibm_firewall":                         resourceIBMFirewall(),
			"ibm_firewall_policy":                  resourceIBMFirewallPolicy(),
			"ibm_iam_user_policy":                  resourceIBMIAMUserPolicy(),
			"ibm_iam_service_id":                   resourceIBMIAMServiceID(),
			"ibm_lb":                               resourceIBMLb(),
			"ibm_lbaas":                            resourceIBMLbaas(),
			"ibm_lbaas_server_instance_attachment": resourceIBMLbaasServerInstanceAttachment(),
			"ibm_lb_service":                       resourceIBMLbService(),
			"ibm_lb_service_group":                 resourceIBMLbServiceGroup(),
			"ibm_lb_vpx":                           resourceIBMLbVpx(),
			"ibm_lb_vpx_ha":                        resourceIBMLbVpxHa(),
			"ibm_lb_vpx_service":                   resourceIBMLbVpxService(),
			"ibm_lb_vpx_vip":                       resourceIBMLbVpxVip(),
			"ibm_network_gateway":                  resourceIBMNetworkGateway(),
			"ibm_network_gateway_vlan_association": resourceIBMNetworkGatewayVlanAttachment(),
			"ibm_network_interface_sg_attachment":  resourceIBMNetworkInterfaceSGAttachment(),
			"ibm_network_public_ip":                resourceIBMNetworkPublicIp(),
			"ibm_network_vlan":                     resourceIBMNetworkVlan(),
			"ibm_object_storage_account":           resourceIBMObjectStorageAccount(),
			"ibm_org":                              resourceIBMOrg(),
			"ibm_resource_instance":                resourceIBMResourceInstance(),
			"ibm_resource_key":                     resourceIBMResourceKey(),
			"ibm_security_group":                   resourceIBMSecurityGroup(),
			"ibm_security_group_rule":              resourceIBMSecurityGroupRule(),
			"ibm_service_instance":                 resourceIBMServiceInstance(),
			"ibm_service_key":                      resourceIBMServiceKey(),
			"ibm_space":                            resourceIBMSpace(),
			"ibm_storage_block":                    resourceIBMStorageBlock(),
			"ibm_storage_file":                     resourceIBMStorageFile(),
			"ibm_subnet":                           resourceIBMSubnet(),
			"ibm_multi_vlan_firewall":              resourceIBMMultiVlanFirewall(),
			"ibm_cdn":                              resourceIBMCDN(),
			"ibm_ipsec_vpn":                        resourceIBMIPSecVPN(),
			"ibm_dns_reverse_record":               resourceIBMDNSREVERSERecord(),
			"ibm_ssl_certificate":                  resourceIBMSSLCertificate(),
			"ibm_firewall_shared":                  resourceIBMFirewallShared(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	bluemixAPIKey := d.Get("bluemix_api_key").(string)
	softlayerUsername := d.Get("softlayer_username").(string)
	softlayerAPIKey := d.Get("softlayer_api_key").(string)
	softlayerTimeout := d.Get("softlayer_timeout").(int)
	bluemixTimeout := d.Get("bluemix_timeout").(int)
	region := d.Get("region").(string)
	retryCount := d.Get("max_retries").(int)
	wskNameSpace := d.Get("function_namespace").(string)

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
		BluemixTimeout:       time.Duration(bluemixTimeout) * time.Second,
		SoftLayerTimeout:     time.Duration(softlayerTimeout) * time.Second,
		SoftLayerUserName:    softlayerUsername,
		SoftLayerAPIKey:      softlayerAPIKey,
		RetryCount:           retryCount,
		SoftLayerEndpointURL: SoftlayerRestEndpoint,
		RetryDelay:           RetryAPIDelay,
		FunctionNameSpace:    wskNameSpace,
	}

	return config.ClientSession()
}
