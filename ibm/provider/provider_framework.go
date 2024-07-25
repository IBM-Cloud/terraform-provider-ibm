package provider

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	bxsession "github.com/IBM-Cloud/bluemix-go/session"
	slsession "github.com/softlayer/softlayer-go/session"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/vpc"

	"github.com/IBM-Cloud/container-services-go-sdk/kubernetesserviceapiv1"
	"github.com/IBM-Cloud/container-services-go-sdk/satellitelinkv1"
	apigateway "github.com/IBM/apigateway-go-sdk/apigatewaycontrollerapiv1"
	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	appid "github.com/IBM/appid-management-go-sdk/appidmanagementv4"
	"github.com/IBM/cloud-databases-go-sdk/clouddatabasesv5"
	"github.com/IBM/container-registry-go-sdk/containerregistryv1"
	cosconfig "github.com/IBM/ibm-cos-sdk-go-config/v2/resourceconfigurationv1"
	kp "github.com/IBM/keyprotect-go-client"
	"github.com/IBM/mqcloud-go-sdk/mqcloudv1"
	cisalertsv1 "github.com/IBM/networking-go-sdk/alertsv1"
	cisoriginpull "github.com/IBM/networking-go-sdk/authenticatedoriginpullapiv1"
	cisbotanalyticsv1 "github.com/IBM/networking-go-sdk/botanalyticsv1"
	cisbotmanagementv1 "github.com/IBM/networking-go-sdk/botmanagementv1"
	ciscachev1 "github.com/IBM/networking-go-sdk/cachingapiv1"
	cisipv1 "github.com/IBM/networking-go-sdk/cisipapiv1"
	ciscustompagev1 "github.com/IBM/networking-go-sdk/custompagesv1"
	dlProviderV2 "github.com/IBM/networking-go-sdk/directlinkproviderv2"
	dl "github.com/IBM/networking-go-sdk/directlinkv1"
	cisdnsbulkv1 "github.com/IBM/networking-go-sdk/dnsrecordbulkv1"
	cisdnsrecordsv1 "github.com/IBM/networking-go-sdk/dnsrecordsv1"
	dns "github.com/IBM/networking-go-sdk/dnssvcsv1"
	cisedgefunctionv1 "github.com/IBM/networking-go-sdk/edgefunctionsapiv1"
	cisfiltersv1 "github.com/IBM/networking-go-sdk/filtersv1"
	cisfirewallrulesv1 "github.com/IBM/networking-go-sdk/firewallrulesv1"
	cisglbhealthcheckv1 "github.com/IBM/networking-go-sdk/globalloadbalancermonitorv1"
	cisglbpoolv0 "github.com/IBM/networking-go-sdk/globalloadbalancerpoolsv0"
	cisglbv1 "github.com/IBM/networking-go-sdk/globalloadbalancerv1"
	cislogpushjobsapiv1 "github.com/IBM/networking-go-sdk/logpushjobsapiv1"
	cismtlsv1 "github.com/IBM/networking-go-sdk/mtlsv1"
	cispagerulev1 "github.com/IBM/networking-go-sdk/pageruleapiv1"
	cisrangeappv1 "github.com/IBM/networking-go-sdk/rangeapplicationsv1"
	cisroutingv1 "github.com/IBM/networking-go-sdk/routingv1"
	cisrulesetsv1 "github.com/IBM/networking-go-sdk/rulesetsv1"
	cissslv1 "github.com/IBM/networking-go-sdk/sslcertificateapiv1"
	tg "github.com/IBM/networking-go-sdk/transitgatewayapisv1"
	cisuarulev1 "github.com/IBM/networking-go-sdk/useragentblockingrulesv1"
	ciswafgroupv1 "github.com/IBM/networking-go-sdk/wafrulegroupsapiv1"
	ciswafpackagev1 "github.com/IBM/networking-go-sdk/wafrulepackagesapiv1"
	ciswafrulev1 "github.com/IBM/networking-go-sdk/wafrulesapiv1"
	ciswebhooksv1 "github.com/IBM/networking-go-sdk/webhooksv1"
	cisaccessrulev1 "github.com/IBM/networking-go-sdk/zonefirewallaccessrulesv1"
	cislockdownv1 "github.com/IBM/networking-go-sdk/zonelockdownv1"
	cisratelimitv1 "github.com/IBM/networking-go-sdk/zoneratelimitsv1"
	cisdomainsettingsv1 "github.com/IBM/networking-go-sdk/zonessettingsv1"
	ciszonesv1 "github.com/IBM/networking-go-sdk/zonesv1"
	"github.com/IBM/platform-services-go-sdk/atrackerv2"
	"github.com/IBM/platform-services-go-sdk/catalogmanagementv1"
	"github.com/IBM/platform-services-go-sdk/contextbasedrestrictionsv1"
	"github.com/IBM/platform-services-go-sdk/enterprisemanagementv1"
	searchv2 "github.com/IBM/platform-services-go-sdk/globalsearchv2"
	"github.com/IBM/platform-services-go-sdk/globaltaggingv1"
	iamaccessgroups "github.com/IBM/platform-services-go-sdk/iamaccessgroupsv2"
	iamidentity "github.com/IBM/platform-services-go-sdk/iamidentityv1"
	iampolicymanagement "github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
	ibmcloudshellv1 "github.com/IBM/platform-services-go-sdk/ibmcloudshellv1"
	"github.com/IBM/platform-services-go-sdk/metricsrouterv3"
	resourcecontroller "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	resourcemanager "github.com/IBM/platform-services-go-sdk/resourcemanagerv2"
	"github.com/IBM/platform-services-go-sdk/usagereportsv4"
	project "github.com/IBM/project-go-sdk/projectv1"
	"github.com/IBM/push-notifications-go-sdk/pushservicev1"
	schematicsv1 "github.com/IBM/schematics-go-sdk/schematicsv1"
	"github.com/IBM/vmware-go-sdk/vmwarev1"
	vpcbeta "github.com/IBM/vpc-beta-go-sdk/vpcbetav1"
	vpcv1 "github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/apache/openwhisk-client-go/whisk"

	"github.com/IBM-Cloud/bluemix-go/api/account/accountv1"
	"github.com/IBM-Cloud/bluemix-go/api/account/accountv2"
	"github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
	"github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
	"github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
	"github.com/IBM-Cloud/bluemix-go/api/functions"
	"github.com/IBM-Cloud/bluemix-go/api/globalsearch/globalsearchv2"
	"github.com/IBM-Cloud/bluemix-go/api/globaltagging/globaltaggingv3"
	"github.com/IBM-Cloud/bluemix-go/api/hpcs"
	"github.com/IBM-Cloud/bluemix-go/api/icd/icdv4"
	"github.com/IBM-Cloud/bluemix-go/api/mccp/mccpv2"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/catalog"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/controller"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev2/controllerv2"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev2/managementv2"
	"github.com/IBM-Cloud/bluemix-go/api/usermanagement/usermanagementv2"
	ibmpisession "github.com/IBM-Cloud/power-go-client/ibmpisession"
	codeengine "github.com/IBM/code-engine-go-sdk/codeenginev2"
	"github.com/IBM/continuous-delivery-go-sdk/cdtektonpipelinev2"
	"github.com/IBM/continuous-delivery-go-sdk/cdtoolchainv2"
	"github.com/IBM/event-notifications-go-admin-sdk/eventnotificationsv1"
	"github.com/IBM/eventstreams-go-sdk/pkg/schemaregistryv1"
	"github.com/IBM/ibm-hpcs-uko-sdk/ukov4"
	"github.com/IBM/logs-go-sdk/logsv0"
	scc "github.com/IBM/scc-go-sdk/v5/securityandcompliancecenterapiv3"
	"github.com/IBM/secrets-manager-go-sdk/v2/secretsmanagerv2"
)

// Ensure the implementation satisfies the provider.Provider interface.
var _ provider.Provider = &ibmCloudProvider{}

type ibmCloudProvider struct {
	version       string
	clientSession interface{}
}

func NewFrameworkProvider(version string) func() provider.Provider {
	log.Printf("[INFO] UJJK New framework provider called")
	return func() provider.Provider {
		return &ibmCloudProvider{
			version: version,
		}
	}
}

func (p *ibmCloudProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "ibm"
	resp.Version = p.version
	log.Printf("[INFO] UJJK New framework provider metadata called")

}

func (p *ibmCloudProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	log.Printf("[INFO] UJJK New framework provider schema called")
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"bluemix_api_key": schema.StringAttribute{
				Optional:    true,
				Description: "The Bluemix API Key",
			},
			"ibmcloud_api_key": schema.StringAttribute{
				Optional:    true,
				Description: "The IBM Cloud API Key",
			},
			"iam_token": schema.StringAttribute{
				Optional:    true,
				Description: "IAM Authentication token",
			},
			"iam_refresh_token": schema.StringAttribute{
				Optional:    true,
				Description: "IAM Authentication refresh token",
			},
			"iam_profile_id": schema.StringAttribute{
				Optional:    true,
				Description: "IAM Trusted Profile Authentication token",
			},
			"softlayer_username": schema.StringAttribute{
				Optional:    true,
				Description: "The SoftLayer user name",
			},
			"iaas_classic_username": schema.StringAttribute{
				Optional:    true,
				Description: "The Classic Infrastructure API user name",
			},
			"softlayer_api_key": schema.StringAttribute{
				Optional:    true,
				Description: "The SoftLayer API Key",
			},
			"iaas_classic_api_key": schema.StringAttribute{
				Optional:    true,
				Description: "The Classic Infrastructure API Key",
			},
			"softlayer_endpoint_url": schema.StringAttribute{
				Optional:    true,
				Description: "The Softlayer Endpoint",
			},
			"iaas_classic_endpoint_url": schema.StringAttribute{
				Optional:    true,
				Description: "The Classic Infrastructure Endpoint",
			},
			"softlayer_timeout": schema.Int64Attribute{
				Optional:    true,
				Description: "The timeout (in seconds) to set for any SoftLayer API calls made.",
			},
			"iaas_classic_timeout": schema.Int64Attribute{
				Optional:    true,
				Description: "The timeout (in seconds) to set for any Classic Infrastructure API calls made.",
			},
			"bluemix_timeout": schema.Int64Attribute{
				Optional:    true,
				Description: "The timeout (in seconds) to set for any Bluemix API calls made.",
			},
			"ibmcloud_timeout": schema.Int64Attribute{
				Optional:    true,
				Description: "The timeout (in seconds) to set for any IBM Cloud API calls made.",
			},
			"visibility": schema.StringAttribute{
				Optional:    true,
				Description: "Visibility of the provider if it is private or public.",
				Validators: []validator.String{
					stringvalidator.OneOf("public", "private", "public-and-private"),
				},
			},
			"endpoints_file_path": schema.StringAttribute{
				Optional:    true,
				Description: "Path of the file that contains private and public regional endpoints mapping",
			},
			"resource_group": schema.StringAttribute{
				Optional:    true,
				Description: "The Resource group id.",
			},
			"region": schema.StringAttribute{
				Optional:    true,
				Description: "The IBM cloud Region (for example 'us-south').",
			},
			"zone": schema.StringAttribute{
				Optional:    true,
				Description: "The IBM cloud Region zone (for example 'us-south-1') for power resources.",
			},
			"max_retries": schema.Int64Attribute{
				Optional:    true,
				Description: "The retry count to set for API calls.",
			},
			"function_namespace": schema.StringAttribute{
				Optional:    true,
				Description: "The IBM Cloud Function namespace",
			},
			"riaas_endpoint": schema.StringAttribute{
				Optional:    true,
				Description: "The next generation infrastructure service endpoint url.",
			},
		},
	}
}

func (p *ibmCloudProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	log.Printf("[INFO] UJJK New framework provider configure called")
	log.Printf("[INFO] UJJK New framework provider configure called, configure ctx is %v, request is %v, response is %v", vpc.PrettifyPrint(ctx), vpc.PrettifyPrint(req), vpc.PrettifyPrint(resp))
	var config ProviderFrameworkModel
	if diags := req.Config.Get(ctx, &config); diags.HasError() {
		resp.Diagnostics = diags
		return
	}
	log.Printf("[INFO] UJJK New framework provider configure achieved %v", vpc.PrettifyPrint(config))
	bluemixAPIKey := config.BluemixAPIKey.ValueString()
	if bluemixAPIKey == "" {
		bluemixAPIKey = config.IBMCloudAPIKey.ValueString()
	}

	iamToken := config.IAMToken.ValueString()
	iamRefreshToken := config.IAMRefreshToken.ValueString()
	iamTrustedProfileId := config.IAMProfileID.ValueString()

	softlayerUsername := config.SoftLayerUsername.ValueString()
	if softlayerUsername == "" {
		softlayerUsername = config.IAASClassicUsername.ValueString()
	}

	softlayerAPIKey := config.SoftLayerAPIKey.ValueString()
	if softlayerAPIKey == "" {
		softlayerAPIKey = config.IAASClassicAPIKey.ValueString()
	}

	softlayerEndpointUrl := config.SoftLayerEndpointURL.ValueString()
	if softlayerEndpointUrl == "" {
		softlayerEndpointUrl = config.IAASClassicEndpointURL.ValueString()
	}

	softlayerTimeout := int(config.SoftLayerTimeout.ValueInt64())
	if softlayerTimeout == 0 {
		softlayerTimeout = int(config.IAASClassicTimeout.ValueInt64())
	}

	bluemixTimeout := int(config.BluemixTimeout.ValueInt64())
	if bluemixTimeout == 0 {
		bluemixTimeout = int(config.IBMCloudTimeout.ValueInt64())
	}

	visibility := config.Visibility.ValueString()
	endpointsFilePath := config.EndpointsFilePath.ValueString()
	resourceGroup := config.ResourceGroup.ValueString()
	region := config.Region.ValueString()
	zone := config.Zone.ValueString()
	retryCount := int(config.MaxRetries.ValueInt64())
	functionNamespace := config.FunctionNamespace.ValueString()
	riaasEndpoint := config.RIAASEndpoint.ValueString()

	if functionNamespace == "" {
		functionNamespace = os.Getenv("FUNCTION_NAMESPACE")
	}
	if functionNamespace != "" {
		os.Setenv("FUNCTION_NAMESPACE", functionNamespace)
	}

	providerConfig := conns.Config{
		BluemixAPIKey:        bluemixAPIKey,
		Region:               region,
		ResourceGroup:        resourceGroup,
		BluemixTimeout:       time.Duration(bluemixTimeout) * time.Second,
		SoftLayerTimeout:     time.Duration(softlayerTimeout) * time.Second,
		SoftLayerUserName:    softlayerUsername,
		SoftLayerAPIKey:      softlayerAPIKey,
		RetryCount:           retryCount,
		SoftLayerEndpointURL: softlayerEndpointUrl,
		RetryDelay:           conns.RetryAPIDelay,
		FunctionNameSpace:    functionNamespace,
		RiaasEndPoint:        riaasEndpoint,
		IAMToken:             iamToken,
		IAMRefreshToken:      iamRefreshToken,
		Zone:                 zone,
		Visibility:           visibility,
		EndpointsFile:        endpointsFilePath,
		IAMTrustedProfileID:  iamTrustedProfileId,
	}

	clientSession, err := providerConfig.ClientSession()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create client session",
			fmt.Sprintf("Unable to create client session: %v", err),
		)
		return
	}
	p.clientSession = clientSession
}

func (p *ibmCloudProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		// Add your data sources here
	}
}

func (p *ibmCloudProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		vpc.NewSSHKeyResource,
		vpc.NewMigrationNewResource,
	}
}

type ProviderFrameworkModel struct {
	BluemixAPIKey          types.String `tfsdk:"bluemix_api_key"`
	IBMCloudAPIKey         types.String `tfsdk:"ibmcloud_api_key"`
	IAMToken               types.String `tfsdk:"iam_token"`
	IAMRefreshToken        types.String `tfsdk:"iam_refresh_token"`
	IAMProfileID           types.String `tfsdk:"iam_profile_id"`
	SoftLayerUsername      types.String `tfsdk:"softlayer_username"`
	IAASClassicUsername    types.String `tfsdk:"iaas_classic_username"`
	SoftLayerAPIKey        types.String `tfsdk:"softlayer_api_key"`
	IAASClassicAPIKey      types.String `tfsdk:"iaas_classic_api_key"`
	SoftLayerEndpointURL   types.String `tfsdk:"softlayer_endpoint_url"`
	IAASClassicEndpointURL types.String `tfsdk:"iaas_classic_endpoint_url"`
	SoftLayerTimeout       types.Int64  `tfsdk:"softlayer_timeout"`
	IAASClassicTimeout     types.Int64  `tfsdk:"iaas_classic_timeout"`
	BluemixTimeout         types.Int64  `tfsdk:"bluemix_timeout"`
	IBMCloudTimeout        types.Int64  `tfsdk:"ibmcloud_timeout"`
	Visibility             types.String `tfsdk:"visibility"`
	EndpointsFilePath      types.String `tfsdk:"endpoints_file_path"`
	ResourceGroup          types.String `tfsdk:"resource_group"`
	Region                 types.String `tfsdk:"region"`
	Zone                   types.String `tfsdk:"zone"`
	MaxRetries             types.Int64  `tfsdk:"max_retries"`
	FunctionNamespace      types.String `tfsdk:"function_namespace"`
	RIAASEndpoint          types.String `tfsdk:"riaas_endpoint"`
}

// ClientSession struct with VPC API client and error
type ClientSession struct {
	session *Session

	appidErr error
	appidAPI *appid.AppIDManagementV4

	apigatewayErr error
	apigatewayAPI *apigateway.ApiGatewayControllerApiV1

	accountConfigErr     error
	bmxAccountServiceAPI accountv2.AccountServiceAPI

	accountV1ConfigErr     error
	bmxAccountv1ServiceAPI accountv1.AccountServiceAPI

	bmxUserDetails  *UserConfig
	bmxUserFetchErr error

	csConfigErr  error
	csServiceAPI containerv1.ContainerServiceAPI

	csv2ConfigErr  error
	csv2ServiceAPI containerv2.ContainerServiceAPI

	containerRegistryClientErr error
	containerRegistryClient    *containerregistryv1.ContainerRegistryV1

	cfConfigErr  error
	cfServiceAPI mccpv2.MccpServiceAPI

	cisConfigErr  error
	cisServiceAPI cisv1.CisServiceAPI

	functionConfigErr error
	functionClient    *whisk.Client

	globalSearchConfigErr  error
	globalSearchServiceAPI globalsearchv2.GlobalSearchServiceAPI

	globalTaggingConfigErr  error
	globalTaggingServiceAPI globaltaggingv3.GlobalTaggingServiceAPI

	globalTaggingConfigErrV1  error
	globalTaggingServiceAPIV1 globaltaggingv1.GlobalTaggingV1

	globalSearchConfigErrV2  error
	globalSearchServiceAPIV2 searchv2.GlobalSearchV2

	ibmCloudShellClient    *ibmcloudshellv1.IBMCloudShellV1
	ibmCloudShellClientErr error

	userManagementErr error
	userManagementAPI usermanagementv2.UserManagementAPI

	icdConfigErr  error
	icdServiceAPI icdv4.ICDServiceAPI

	cloudDatabasesClientErr error
	cloudDatabasesClient    *clouddatabasesv5.CloudDatabasesV5

	resourceControllerConfigErr  error
	resourceControllerServiceAPI controller.ResourceControllerAPI

	resourceControllerConfigErrv2  error
	resourceControllerServiceAPIv2 controllerv2.ResourceControllerAPIV2

	resourceManagementConfigErrv2  error
	resourceManagementServiceAPIv2 managementv2.ResourceManagementAPIv2

	resourceCatalogConfigErr  error
	resourceCatalogServiceAPI catalog.ResourceCatalogAPI

	ibmpiConfigErr error
	ibmpiSession   *ibmpisession.IBMPISession

	kpErr error
	kpAPI *kp.API

	kmsErr error
	kmsAPI *kp.API

	hpcsEndpointErr error
	hpcsEndpointAPI hpcs.HPCSV2

	ukoClient    *ukov4.UkoV4
	ukoClientErr error

	pDNSClient *dns.DnsSvcsV1
	pDNSErr    error

	bluemixSessionErr error

	pushServiceClient    *pushservicev1.PushServiceV1
	pushServiceClientErr error

	eventNotificationsApiClient    *eventnotificationsv1.EventNotificationsV1
	eventNotificationsApiClientErr error

	appConfigurationClient    *appconfigurationv1.AppConfigurationV1
	appConfigurationClientErr error

	vpcErr     error
	vpcAPI     *vpcv1.VpcV1
	vpcbetaErr error
	vpcBetaAPI *vpcbeta.VpcbetaV1

	directlinkAPI *dl.DirectLinkV1
	directlinkErr error
	dlProviderAPI *dlProviderV2.DirectLinkProviderV2
	dlProviderErr error

	cosConfigErr error
	cosConfigAPI *cosconfig.ResourceConfigurationV1

	transitgatewayAPI *tg.TransitGatewayApisV1
	transitgatewayErr error

	functionIAMNamespaceAPI functions.FunctionServiceAPI
	functionIAMNamespaceErr error

	// CIS Zones
	cisZonesErr      error
	cisZonesV1Client *ciszonesv1.ZonesV1

	// CIS Alerts
	cisAlertsClient *cisalertsv1.AlertsV1
	cisAlertsErr    error

	// CIS Rulesets
	cisRulesetsClient *cisrulesetsv1.RulesetsV1
	cisRulesetsErr    error

	// CIS Authenticated Origin Pull
	cisOriginAuthClient  *cisoriginpull.AuthenticatedOriginPullApiV1
	cisOriginAuthPullErr error

	// CIS dns service options
	cisDNSErr           error
	cisDNSRecordsClient *cisdnsrecordsv1.DnsRecordsV1

	// CIS dns bulk service options
	cisDNSBulkErr          error
	cisDNSRecordBulkClient *cisdnsbulkv1.DnsRecordBulkV1

	// CIS Global Load Balancer Pool service options
	cisGLBPoolErr    error
	cisGLBPoolClient *cisglbpoolv0.GlobalLoadBalancerPoolsV0

	// CIS GLB service options
	cisGLBErr    error
	cisGLBClient *cisglbv1.GlobalLoadBalancerV1

	// CIS GLB health check service options
	cisGLBHealthCheckErr    error
	cisGLBHealthCheckClient *cisglbhealthcheckv1.GlobalLoadBalancerMonitorV1

	// CIS IP service options
	cisIPErr    error
	cisIPClient *cisipv1.CisIpApiV1

	// CIS Zone Rate Limits service options
	cisRLErr    error
	cisRLClient *cisratelimitv1.ZoneRateLimitsV1

	// CIS Page Rules service options
	cisPageRuleErr    error
	cisPageRuleClient *cispagerulev1.PageRuleApiV1

	// CIS Edge Functions service options
	cisEdgeFunctionErr    error
	cisEdgeFunctionClient *cisedgefunctionv1.EdgeFunctionsApiV1

	// CIS SSL certificate service options
	cisSSLErr    error
	cisSSLClient *cissslv1.SslCertificateApiV1

	// CIS WAF Package service options
	cisWAFPackageErr    error
	cisWAFPackageClient *ciswafpackagev1.WafRulePackagesApiV1

	// CIS Zone Setting service options
	cisDomainSettingsErr    error
	cisDomainSettingsClient *cisdomainsettingsv1.ZonesSettingsV1

	// CIS Routing service options
	cisRoutingErr    error
	cisRoutingClient *cisroutingv1.RoutingV1

	// CIS WAF Group service options
	cisWAFGroupErr    error
	cisWAFGroupClient *ciswafgroupv1.WafRuleGroupsApiV1

	// CIS Caching service options
	cisCacheErr    error
	cisCacheClient *ciscachev1.CachingApiV1

	// CIS Custom Pages service options
	cisCustomPageErr    error
	cisCustomPageClient *ciscustompagev1.CustomPagesV1

	// CIS Firewall Access rule service option
	cisAccessRuleErr    error
	cisAccessRuleClient *cisaccessrulev1.ZoneFirewallAccessRulesV1

	// CIS User Agent Blocking Rule service option
	cisUARuleErr    error
	cisUARuleClient *cisuarulev1.UserAgentBlockingRulesV1

	// CIS Firewall Lockdwon Rule service option
	cisLockdownErr    error
	cisLockdownClient *cislockdownv1.ZoneLockdownV1

	// CIS LogpushJobs service option
	cisLogpushJobsClient *cislogpushjobsapiv1.LogpushJobsApiV1
	cisLogpushJobsErr    error

	// CIS Range app service option
	cisRangeAppErr    error
	cisRangeAppClient *cisrangeappv1.RangeApplicationsV1

	// CIS WAF rule service options
	cisWAFRuleErr    error
	cisWAFRuleClient *ciswafrulev1.WafRulesApiV1
	// IAM Identity Option
	iamIdentityErr error
	iamIdentityAPI *iamidentity.IamIdentityV1

	// Resource Manager Option
	resourceManagerErr error
	resourceManagerAPI *resourcemanager.ResourceManagerV2

	// Catalog Management Option
	catalogManagementClient    *catalogmanagementv1.CatalogManagementV1
	catalogManagementClientErr error

	enterpriseManagementClient    *enterprisemanagementv1.EnterpriseManagementV1
	enterpriseManagementClientErr error

	// Resource Controller Option
	resourceControllerErr   error
	resourceControllerAPI   *resourcecontroller.ResourceControllerV2
	secretsManagerClient    *secretsmanagerv2.SecretsManagerV2
	secretsManagerClientErr error

	// Schematics service options
	schematicsClient    *schematicsv1.SchematicsV1
	schematicsClientErr error

	// Satellite service
	satelliteClient    *kubernetesserviceapiv1.KubernetesServiceApiV1
	satelliteClientErr error

	// IAM Policy Management
	iamPolicyManagementErr error
	iamPolicyManagementAPI *iampolicymanagement.IamPolicyManagementV1

	// IAM Access Groups
	iamAccessGroupsErr error
	iamAccessGroupsAPI *iamaccessgroups.IamAccessGroupsV2

	// MTLS Session options
	cisMtlsClient *cismtlsv1.MtlsV1
	cisMtlsErr    error

	// Bot Management options
	cisBotManagementClient *cisbotmanagementv1.BotManagementV1
	cisBotManagementErr    error

	// Bot Analytics options
	cisBotAnalyticsClient *cisbotanalyticsv1.BotAnalyticsV1
	cisBotAnalyticsErr    error

	// CIS Webhooks options
	cisWebhooksClient *ciswebhooksv1.WebhooksV1
	cisWebhooksErr    error

	// CIS Filters options
	cisFiltersClient *cisfiltersv1.FiltersV1
	cisFiltersErr    error

	// CIS FirewallRules options
	cisFirewallRulesClient *cisfirewallrulesv1.FirewallRulesV1
	cisFirewallRulesErr    error

	// Atracker
	atrackerClientV2    *atrackerv2.AtrackerV2
	atrackerClientV2Err error

	// Metrics Router
	metricsRouterClient    *metricsrouterv3.MetricsRouterV3
	metricsRouterClientErr error

	// Satellite link service
	satelliteLinkClient    *satellitelinkv1.SatelliteLinkV1
	satelliteLinkClientErr error

	esSchemaRegistryClient *schemaregistryv1.SchemaregistryV1
	esSchemaRegistryErr    error

	// Security and Compliance Center (SCC)
	securityAndComplianceCenterClient    *scc.SecurityAndComplianceCenterApiV3
	securityAndComplianceCenterClientErr error

	// context Based Restrictions (CBR)
	contextBasedRestrictionsClient    *contextbasedrestrictionsv1.ContextBasedRestrictionsV1
	contextBasedRestrictionsClientErr error

	// CD Toolchain
	cdToolchainClient    *cdtoolchainv2.CdToolchainV2
	cdToolchainClientErr error

	// CD Tekton Pipeline
	cdTektonPipelineClient    *cdtektonpipelinev2.CdTektonPipelineV2
	cdTektonPipelineClientErr error

	// Code Engine options
	codeEngineClient    *codeengine.CodeEngineV2
	codeEngineClientErr error

	// Project options
	projectClient    *project.ProjectV1
	projectClientErr error

	// Usage Reports options
	usageReportsClient    *usagereportsv4.UsageReportsV4
	usageReportsClientErr error

	mqcloudClient    *mqcloudv1.MqcloudV1
	mqcloudClientErr error

	// VMware as a Service
	vmwareClient    *vmwarev1.VmwareV1
	vmwareClientErr error

	// Cloud Logs
	logsClient    *logsv0.LogsV0
	logsClientErr error
}

type Session struct {
	// SoftLayerSesssion is the the SoftLayer session used to connect to the SoftLayer API
	SoftLayerSession *slsession.Session

	// BluemixSession is the the Bluemix session used to connect to the Bluemix API
	BluemixSession *bxsession.Session
}

type UserConfig struct {
	UserID      string
	UserEmail   string
	UserAccount string
	CloudName   string `default:"bluemix"`
	cloudType   string `default:"public"`
	generation  int    `default:"2"`
}
