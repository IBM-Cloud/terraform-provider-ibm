// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package conns

import (
	"fmt"
	"log"
	gohttp "net/http"
	"os"
	"strings"
	"time"

	// Added code for the Power Colo Offering

	"github.com/IBM-Cloud/container-services-go-sdk/kubernetesserviceapiv1"
	"github.com/IBM-Cloud/container-services-go-sdk/satellitelinkv1"
	"github.com/IBM-Cloud/terraform-provider-ibm/version"
	apigateway "github.com/IBM/apigateway-go-sdk/apigatewaycontrollerapiv1"
	"github.com/IBM/appconfiguration-go-admin-sdk/appconfigurationv1"
	appid "github.com/IBM/appid-management-go-sdk/appidmanagementv4"
	"github.com/IBM/cloud-databases-go-sdk/clouddatabasesv5"
	"github.com/IBM/container-registry-go-sdk/containerregistryv1"
	"github.com/IBM/go-sdk-core/v5/core"
	cosconfig "github.com/IBM/ibm-cos-sdk-go-config/resourceconfigurationv1"
	kp "github.com/IBM/keyprotect-go-client"
	cisalertsv1 "github.com/IBM/networking-go-sdk/alertsv1"
	cisoriginpull "github.com/IBM/networking-go-sdk/authenticatedoriginpullapiv1"
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
	"github.com/IBM/platform-services-go-sdk/atrackerv1"
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
	resourcecontroller "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	resourcemanager "github.com/IBM/platform-services-go-sdk/resourcemanagerv2"
	"github.com/IBM/push-notifications-go-sdk/pushservicev1"
	"github.com/IBM/scc-go-sdk/v3/adminserviceapiv1"
	"github.com/IBM/scc-go-sdk/v3/configurationgovernancev1"
	"github.com/IBM/scc-go-sdk/v4/posturemanagementv2"
	schematicsv1 "github.com/IBM/schematics-go-sdk/schematicsv1"
	vpcbeta "github.com/IBM/vpc-beta-go-sdk/vpcbetav1"
	"github.com/IBM/vpc-go-sdk/common"
	vpc "github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/apache/openwhisk-client-go/whisk"
	jwt "github.com/golang-jwt/jwt"
	slsession "github.com/softlayer/softlayer-go/session"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/api/globaltagging/globaltaggingv3"
	resourcecontroller "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev2/managementv2"
	"github.com/IBM-Cloud/bluemix-go/authentication"
	"github.com/IBM-Cloud/bluemix-go/http"
	"github.com/IBM-Cloud/bluemix-go/rest"
	bxsession "github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/terraform-provider-ibm/version"
	"github.com/IBM/go-sdk-core/v5/core"
	jwt "github.com/golang-jwt/jwt"
	"github.com/damianovesperini/platform-services-go-sdk/projectv1"
)

//Config stores user provider input
type Config struct {
	//BluemixAPIKey is the Bluemix api key
	BluemixAPIKey string
	//Bluemix region
	Region string
	//Resource group id
	ResourceGroup string
	//Bluemix API timeout
	BluemixTimeout time.Duration

	//Retry Count for API calls
	//Unexposed in the schema at this point as they are used only during session creation for a few calls
	//When sdk implements it we an expose them for expected behaviour
	//https://github.com/softlayer/softlayer-go/issues/41
	RetryCount int
	//Constant Retry Delay for API calls
	RetryDelay time.Duration

	// Zone
	Zone       string
	Visibility string
}

//UserConfig ...
type UserConfig struct {
	userID      string
	userEmail   string
	UserAccount string
	cloudName   string `default:"bluemix"`
	cloudType   string `default:"public"`
}

//Session stores the information required for communication with the SoftLayer and Bluemix API
type Session struct {
	// BluemixSession is the the Bluemix session used to connect to the Bluemix API
	BluemixSession *bxsession.Session
}

// ClientSession ...
type ClientSession interface {
	BluemixUserDetails() (*UserConfig, error)
	GlobalTaggingAPI() (globaltaggingv3.GlobalTaggingServiceAPI, error)

	GlobalTaggingAPIv1() (globaltaggingv1.GlobalTaggingV1, error)
	GlobalSearchAPIV2() (searchv2.GlobalSearchV2, error)
	ICDAPI() (icdv4.ICDServiceAPI, error)
	CloudDatabasesV5() (*clouddatabasesv5.CloudDatabasesV5, error)
	IAMPolicyManagementV1API() (*iampolicymanagement.IamPolicyManagementV1, error)
	IAMAccessGroupsV2() (*iamaccessgroups.IamAccessGroupsV2, error)
	MccpAPI() (mccpv2.MccpServiceAPI, error)
	ResourceCatalogAPI() (catalog.ResourceCatalogAPI, error)
	ResourceManagementAPIv2() (managementv2.ResourceManagementAPIv2, error)
	ResourceControllerAPI() (controller.ResourceControllerAPI, error)
	ResourceControllerAPIV2() (controllerv2.ResourceControllerAPIV2, error)
	SoftLayerSession() *slsession.Session
	IBMPISession() (*ibmpisession.IBMPISession, error)
	UserManagementAPI() (usermanagementv2.UserManagementAPI, error)
	PushServiceV1() (*pushservicev1.PushServiceV1, error)
	EventNotificationsApiV1() (*eventnotificationsv1.EventNotificationsV1, error)
	AppConfigurationV1() (*appconfigurationv1.AppConfigurationV1, error)
	CertificateManagerAPI() (certificatemanager.CertificateManagerServiceAPI, error)
	KeyProtectAPI() (*kp.Client, error)
	KeyManagementAPI() (*kp.Client, error)
	VpcV1API() (*vpc.VpcV1, error)
	VpcV1BetaAPI() (*vpcbeta.VpcbetaV1, error)
	APIGateway() (*apigateway.ApiGatewayControllerApiV1, error)
	PrivateDNSClientSession() (*dns.DnsSvcsV1, error)
	CosConfigV1API() (*cosconfig.ResourceConfigurationV1, error)
	DirectlinkV1API() (*dl.DirectLinkV1, error)
	DirectlinkProviderV2API() (*dlProviderV2.DirectLinkProviderV2, error)
	TransitGatewayV1API() (*tg.TransitGatewayApisV1, error)
	HpcsEndpointAPI() (hpcs.HPCSV2, error)
	UkoV4() (*ukov4.UkoV4, error)
	FunctionIAMNamespaceAPI() (functions.FunctionServiceAPI, error)
	CisZonesV1ClientSession() (*ciszonesv1.ZonesV1, error)
	CisAlertsSession() (*cisalertsv1.AlertsV1, error)
	CisOrigAuthSession() (*cisoriginpull.AuthenticatedOriginPullApiV1, error)
	CisDNSRecordClientSession() (*cisdnsrecordsv1.DnsRecordsV1, error)
	CisDNSRecordBulkClientSession() (*cisdnsbulkv1.DnsRecordBulkV1, error)
	CisGLBClientSession() (*cisglbv1.GlobalLoadBalancerV1, error)
	CisGLBPoolClientSession() (*cisglbpoolv0.GlobalLoadBalancerPoolsV0, error)
	CisGLBHealthCheckClientSession() (*cisglbhealthcheckv1.GlobalLoadBalancerMonitorV1, error)
	CisIPClientSession() (*cisipv1.CisIpApiV1, error)
	CisPageRuleClientSession() (*cispagerulev1.PageRuleApiV1, error)
	CisLogpushJobsSession() (*cislogpushjobsapiv1.LogpushJobsApiV1, error)
	CisRLClientSession() (*cisratelimitv1.ZoneRateLimitsV1, error)
	CisEdgeFunctionClientSession() (*cisedgefunctionv1.EdgeFunctionsApiV1, error)
	CisSSLClientSession() (*cissslv1.SslCertificateApiV1, error)
	CisWAFPackageClientSession() (*ciswafpackagev1.WafRulePackagesApiV1, error)
	CisDomainSettingsClientSession() (*cisdomainsettingsv1.ZonesSettingsV1, error)
	CisRoutingClientSession() (*cisroutingv1.RoutingV1, error)
	CisWAFGroupClientSession() (*ciswafgroupv1.WafRuleGroupsApiV1, error)
	CisCacheClientSession() (*ciscachev1.CachingApiV1, error)
	CisMtlsSession() (*cismtlsv1.MtlsV1, error)
	CisWebhookSession() (*ciswebhooksv1.WebhooksV1, error)
	CisCustomPageClientSession() (*ciscustompagev1.CustomPagesV1, error)
	CisAccessRuleClientSession() (*cisaccessrulev1.ZoneFirewallAccessRulesV1, error)
	CisUARuleClientSession() (*cisuarulev1.UserAgentBlockingRulesV1, error)
	CisLockdownClientSession() (*cislockdownv1.ZoneLockdownV1, error)
	CisRangeAppClientSession() (*cisrangeappv1.RangeApplicationsV1, error)
	CisWAFRuleClientSession() (*ciswafrulev1.WafRulesApiV1, error)
	IAMIdentityV1API() (*iamidentity.IamIdentityV1, error)
	IBMCloudShellV1() (*ibmcloudshellv1.IBMCloudShellV1, error)
	ResourceManagerV2API() (*resourcemanager.ResourceManagerV2, error)
	CatalogManagementV1() (*catalogmanagementv1.CatalogManagementV1, error)
	EnterpriseManagementV1() (*enterprisemanagementv1.EnterpriseManagementV1, error)
	ResourceControllerV2API() (*resourcecontroller.ResourceControllerV2, error)
	ResourceManagementAPIv2() (managementv2.ResourceManagementAPIv2, error)
	ProjectV1() (*projectv1.ProjectV1, error)
}

type clientSession struct {
	session *Session

	bmxUserDetails  *UserConfig
	bmxUserFetchErr error

	globalTaggingConfigErr  error
	globalTaggingServiceAPI globaltaggingv3.GlobalTaggingServiceAPI

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
	vpcAPI     *vpc.VpcV1
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
	//IAM Identity Option
	iamIdentityErr error
	iamIdentityAPI *iamidentity.IamIdentityV1

	//Resource Manager Option
	resourceManagerErr error
	resourceManagerAPI *resourcemanager.ResourceManagerV2

	//Catalog Management Option
	catalogManagementClient    *catalogmanagementv1.CatalogManagementV1
	catalogManagementClientErr error

	enterpriseManagementClient    *enterprisemanagementv1.EnterpriseManagementV1
	enterpriseManagementClientErr error

	//Resource Controller Option
	resourceControllerErr   error
	resourceControllerAPI   *resourcecontroller.ResourceControllerV2

	projectClient     *projectv1.ProjectV1
	projectClientErr  error
}

func (session clientSession) BluemixUserDetails() (*UserConfig, error) {
	return session.bmxUserDetails, session.bmxUserFetchErr
}

// GlobalTaggingAPI provides Global Search  APIs ...
func (sess clientSession) GlobalTaggingAPI() (globaltaggingv3.GlobalTaggingServiceAPI, error) {
	return sess.globalTaggingServiceAPI, sess.globalTaggingConfigErr
}

// ResourceController Session
func (sess clientSession) ResourceControllerV2API() (*resourcecontroller.ResourceControllerV2, error) {
	return sess.resourceControllerAPI, sess.resourceControllerErr
}

// ResourceManagementAPIv2 ...
func (sess clientSession) ResourceManagementAPIv2() (managementv2.ResourceManagementAPIv2, error) {
	return sess.resourceManagementServiceAPIv2, sess.resourceManagementConfigErrv2
}


// Projects API Specification
func (session clientSession) ProjectV1() (*projectv1.ProjectV1, error) {
	return session.projectClient, session.projectClientErr
// ResourceControllerAPI ...
func (sess clientSession) ResourceControllerAPI() (controller.ResourceControllerAPI, error) {
	return sess.resourceControllerServiceAPI, sess.resourceControllerConfigErr
}

// ResourceControllerAPIv2 ...
func (sess clientSession) ResourceControllerAPIV2() (controllerv2.ResourceControllerAPIV2, error) {
	return sess.resourceControllerServiceAPIv2, sess.resourceControllerConfigErrv2
}

// SoftLayerSession providers SoftLayer Session
func (sess clientSession) SoftLayerSession() *slsession.Session {
	return sess.session.SoftLayerSession
}

// CertManagementAPI provides Certificate  management APIs ...
func (sess clientSession) CertificateManagerAPI() (certificatemanager.CertificateManagerServiceAPI, error) {
	return sess.certManagementAPI, sess.certManagementErr
}

// apigatewayAPI provides API Gateway APIs
func (sess clientSession) APIGateway() (*apigateway.ApiGatewayControllerApiV1, error) {
	return sess.apigatewayAPI, sess.apigatewayErr
}

func (session clientSession) PushServiceV1() (*pushservicev1.PushServiceV1, error) {
	return session.pushServiceClient, session.pushServiceClientErr
}

func (session clientSession) EventNotificationsApiV1() (*eventnotificationsv1.EventNotificationsV1, error) {
	return session.eventNotificationsApiClient, session.eventNotificationsApiClientErr
}

func (session clientSession) AppConfigurationV1() (*appconfigurationv1.AppConfigurationV1, error) {
	return session.appConfigurationClient, session.appConfigurationClientErr
}

func (sess clientSession) KeyProtectAPI() (*kp.Client, error) {
	return sess.kpAPI, sess.kpErr
}

func (sess clientSession) KeyManagementAPI() (*kp.Client, error) {
	if sess.kmsErr == nil {
		var clientConfig *kp.ClientConfig
		if sess.kmsAPI.Config.APIKey != "" {
			clientConfig = &kp.ClientConfig{
				BaseURL:  EnvFallBack([]string{"IBMCLOUD_KP_API_ENDPOINT"}, sess.kmsAPI.Config.BaseURL),
				APIKey:   sess.kmsAPI.Config.APIKey, //pragma: allowlist secret
				Verbose:  kp.VerboseFailOnly,
				TokenURL: sess.kmsAPI.Config.TokenURL,
			}
		} else {
			clientConfig = &kp.ClientConfig{
				BaseURL:       EnvFallBack([]string{"IBMCLOUD_KP_API_ENDPOINT"}, sess.kmsAPI.Config.BaseURL),
				Authorization: sess.session.BluemixSession.Config.IAMAccessToken, //pragma: allowlist secret
				Verbose:       kp.VerboseFailOnly,
				TokenURL:      sess.kmsAPI.Config.TokenURL,
			}
		}

		kpClient, err := kp.New(*clientConfig, DefaultTransport())
		if err != nil {
			sess.kpErr = fmt.Errorf("[ERROR] Error occured while configuring Key Protect Service: %q", err)
		}
		return kpClient, nil
	}
	return sess.kmsAPI, sess.kmsErr
}

func (sess clientSession) VpcV1API() (*vpc.VpcV1, error) {
	return sess.vpcAPI, sess.vpcErr
}

func (sess clientSession) VpcV1BetaAPI() (*vpcbeta.VpcbetaV1, error) {
	return sess.vpcBetaAPI, sess.vpcbetaErr
}

func (sess clientSession) DirectlinkV1API() (*dl.DirectLinkV1, error) {
	return sess.directlinkAPI, sess.directlinkErr
}
func (sess clientSession) DirectlinkProviderV2API() (*dlProviderV2.DirectLinkProviderV2, error) {
	return sess.dlProviderAPI, sess.dlProviderErr
}
func (sess clientSession) CosConfigV1API() (*cosconfig.ResourceConfigurationV1, error) {
	return sess.cosConfigAPI, sess.cosConfigErr
}

func (sess clientSession) TransitGatewayV1API() (*tg.TransitGatewayApisV1, error) {
	return sess.transitgatewayAPI, sess.transitgatewayErr
}

// Session to the Power Colo Service

func (sess clientSession) IBMPISession() (*ibmpisession.IBMPISession, error) {
	return sess.ibmpiSession, sess.ibmpiConfigErr
}

// Private DNS Service

func (sess clientSession) PrivateDNSClientSession() (*dns.DnsSvcsV1, error) {
	return sess.pDNSClient, sess.pDNSErr
}

// Session to the Namespace cloud function

func (sess clientSession) FunctionIAMNamespaceAPI() (functions.FunctionServiceAPI, error) {
	return sess.functionIAMNamespaceAPI, sess.functionIAMNamespaceErr
}

// CIS Zones Service
func (sess clientSession) CisZonesV1ClientSession() (*ciszonesv1.ZonesV1, error) {
	if sess.cisZonesErr != nil {
		return sess.cisZonesV1Client, sess.cisZonesErr
	}
	return sess.cisZonesV1Client.Clone(), nil
}

// CIS DNS Service
func (sess clientSession) CisDNSRecordClientSession() (*cisdnsrecordsv1.DnsRecordsV1, error) {
	if sess.cisDNSErr != nil {
		return sess.cisDNSRecordsClient, sess.cisDNSErr
	}
	return sess.cisDNSRecordsClient.Clone(), nil
}

// CIS DNS Bulk Service
func (sess clientSession) CisDNSRecordBulkClientSession() (*cisdnsbulkv1.DnsRecordBulkV1, error) {
	if sess.cisDNSBulkErr != nil {
		return sess.cisDNSRecordBulkClient, sess.cisDNSBulkErr
	}
	return sess.cisDNSRecordBulkClient.Clone(), nil
}

// CIS GLB Pool
func (sess clientSession) CisGLBPoolClientSession() (*cisglbpoolv0.GlobalLoadBalancerPoolsV0, error) {
	if sess.cisGLBPoolErr != nil {
		return sess.cisGLBPoolClient, sess.cisGLBPoolErr
	}
	return sess.cisGLBPoolClient.Clone(), nil
}

// CIS GLB
func (sess clientSession) CisGLBClientSession() (*cisglbv1.GlobalLoadBalancerV1, error) {
	if sess.cisGLBErr != nil {
		return sess.cisGLBClient, sess.cisGLBErr
	}
	return sess.cisGLBClient.Clone(), nil
}

// CIS GLB Health Check/Monitor
func (sess clientSession) CisGLBHealthCheckClientSession() (*cisglbhealthcheckv1.GlobalLoadBalancerMonitorV1, error) {
	if sess.cisGLBHealthCheckErr != nil {
		return sess.cisGLBHealthCheckClient, sess.cisGLBHealthCheckErr
	}
	return sess.cisGLBHealthCheckClient.Clone(), nil
}

// CIS Zone Rate Limits
func (sess clientSession) CisRLClientSession() (*cisratelimitv1.ZoneRateLimitsV1, error) {
	if sess.cisRLErr != nil {
		return sess.cisRLClient, sess.cisRLErr
	}
	return sess.cisRLClient.Clone(), nil
}

// CIS IP
func (sess clientSession) CisIPClientSession() (*cisipv1.CisIpApiV1, error) {
	if sess.cisIPErr != nil {
		return sess.cisIPClient, sess.cisIPErr
	}
	return sess.cisIPClient.Clone(), nil
}

// CIS Page Rules
func (sess clientSession) CisPageRuleClientSession() (*cispagerulev1.PageRuleApiV1, error) {
	if sess.cisPageRuleErr != nil {
		return sess.cisPageRuleClient, sess.cisPageRuleErr
	}
	return sess.cisPageRuleClient.Clone(), nil
}

// CIS Edge Function
func (sess clientSession) CisEdgeFunctionClientSession() (*cisedgefunctionv1.EdgeFunctionsApiV1, error) {
	if sess.cisEdgeFunctionErr != nil {
		return sess.cisEdgeFunctionClient, sess.cisEdgeFunctionErr
	}
	return sess.cisEdgeFunctionClient.Clone(), nil
}

// CIS SSL certificate
func (sess clientSession) CisSSLClientSession() (*cissslv1.SslCertificateApiV1, error) {
	if sess.cisSSLErr != nil {
		return sess.cisSSLClient, sess.cisSSLErr
	}
	return sess.cisSSLClient.Clone(), nil
}

// ClientSession configures and returns a fully initialized ClientSession
func (c *Config) ClientSession() (interface{}, error) {
	sess, err := newSession(c)
	if err != nil {
		return nil, err
	}

	session := clientSession{
		session: sess,

	return sess.cisWAFPackageClient.Clone(), nil
}

// CIS Zone Settings
func (sess clientSession) CisDomainSettingsClientSession() (*cisdomainsettingsv1.ZonesSettingsV1, error) {
	if sess.cisDomainSettingsErr != nil {
		return sess.cisDomainSettingsClient, sess.cisDomainSettingsErr
	}
	return sess.cisDomainSettingsClient.Clone(), nil
}

// CIS Alerts
func (sess clientSession) CisAlertsSession() (*cisalertsv1.AlertsV1, error) {
	if sess.cisAlertsErr != nil {
		return sess.cisAlertsClient, sess.cisAlertsErr
	}
	return sess.cisAlertsClient.Clone(), nil
}

// CIS Routing
func (sess clientSession) CisRoutingClientSession() (*cisroutingv1.RoutingV1, error) {
	if sess.cisRoutingErr != nil {
		return sess.cisRoutingClient, sess.cisRoutingErr
	}
	return sess.cisRoutingClient.Clone(), nil
}

// CIS WAF Group
func (sess clientSession) CisWAFGroupClientSession() (*ciswafgroupv1.WafRuleGroupsApiV1, error) {
	if sess.cisWAFGroupErr != nil {
		return sess.cisWAFGroupClient, sess.cisWAFGroupErr
	}
	return sess.cisWAFGroupClient.Clone(), nil
}

// CIS Cache service
func (sess clientSession) CisCacheClientSession() (*ciscachev1.CachingApiV1, error) {
	if sess.cisCacheErr != nil {
		return sess.cisCacheClient, sess.cisCacheErr
	}
	return sess.cisCacheClient.Clone(), nil
}

// CIS Zone Settings
func (sess clientSession) CisCustomPageClientSession() (*ciscustompagev1.CustomPagesV1, error) {
	if sess.cisCustomPageErr != nil {
		return sess.cisCustomPageClient, sess.cisCustomPageErr
	}
	return sess.cisCustomPageClient.Clone(), nil
}

// CIS Firewall access rule
func (sess clientSession) CisAccessRuleClientSession() (*cisaccessrulev1.ZoneFirewallAccessRulesV1, error) {
	if sess.cisAccessRuleErr != nil {
		return sess.cisAccessRuleClient, sess.cisAccessRuleErr
	}
	return sess.cisAccessRuleClient.Clone(), nil
}

// CIS User Agent Blocking rule
func (sess clientSession) CisUARuleClientSession() (*cisuarulev1.UserAgentBlockingRulesV1, error) {
	if sess.cisUARuleErr != nil {
		return sess.cisUARuleClient, sess.cisUARuleErr
	}
	return sess.cisUARuleClient.Clone(), nil
}

// CIS Firewall Lockdown rule
func (sess clientSession) CisLockdownClientSession() (*cislockdownv1.ZoneLockdownV1, error) {
	if sess.cisLockdownErr != nil {
		return sess.cisLockdownClient, sess.cisLockdownErr
	}
	return sess.cisLockdownClient.Clone(), nil
}

// CIS Range app rule
func (sess clientSession) CisRangeAppClientSession() (*cisrangeappv1.RangeApplicationsV1, error) {
	if sess.cisRangeAppErr != nil {
		return sess.cisRangeAppClient, sess.cisRangeAppErr
	}
	return sess.cisRangeAppClient.Clone(), nil
}

// CIS WAF Rule
func (sess clientSession) CisWAFRuleClientSession() (*ciswafrulev1.WafRulesApiV1, error) {
	if sess.cisWAFRuleErr != nil {
		return sess.cisWAFRuleClient, sess.cisWAFRuleErr
	}
	return sess.cisWAFRuleClient.Clone(), nil
}

// CIS Authenticated Origin Pull
func (sess clientSession) CisOrigAuthSession() (*cisoriginpull.AuthenticatedOriginPullApiV1, error) {
	if sess.cisOriginAuthPullErr != nil {
		return sess.cisOriginAuthClient, sess.cisOriginAuthPullErr
	}
	return sess.cisOriginAuthClient.Clone(), nil
}

// IAM Identity Session
func (sess clientSession) IAMIdentityV1API() (*iamidentity.IamIdentityV1, error) {
	return sess.iamIdentityAPI, sess.iamIdentityErr
}

// ResourceMAanger Session
func (sess clientSession) ResourceManagerV2API() (*resourcemanager.ResourceManagerV2, error) {
	return sess.resourceManagerAPI, sess.resourceManagerErr
}

func (session clientSession) EnterpriseManagementV1() (*enterprisemanagementv1.EnterpriseManagementV1, error) {
	return session.enterpriseManagementClient, session.enterpriseManagementClientErr
}

// ResourceController Session
func (sess clientSession) ResourceControllerV2API() (*resourcecontroller.ResourceControllerV2, error) {
	return sess.resourceControllerAPI, sess.resourceControllerErr
}

// IBM Cloud Secrets Manager V1 Basic API
func (session clientSession) SecretsManagerV1() (*secretsmanagerv1.SecretsManagerV1, error) {
	return session.secretsManagerClientV1, session.secretsManagerClientErr
}

// IBM Cloud Secrets Manager V2 Basic API
func (session clientSession) SecretsManagerV2() (*secretsmanagerv2.SecretsManagerV2, error) {
	return session.secretsManagerClient, session.secretsManagerClientErr
}

// Satellite Link
func (session clientSession) SatellitLinkClientSession() (*satellitelinkv1.SatelliteLinkV1, error) {
	return session.satelliteLinkClient, session.satelliteLinkClientErr
}

var cloudEndpoint = "cloud.ibm.com"

// Session to the Satellite client
func (sess clientSession) SatelliteClientSession() (*kubernetesserviceapiv1.KubernetesServiceApiV1, error) {
	return sess.satelliteClient, sess.satelliteClientErr
}

// CIS LogPushJob
func (sess clientSession) CisLogpushJobsSession() (*cislogpushjobsapiv1.LogpushJobsApiV1, error) {
	if sess.cisLogpushJobsErr != nil {
		return sess.cisLogpushJobsClient, sess.cisLogpushJobsErr
	}
	return sess.cisLogpushJobsClient.Clone(), nil
}

// CIS MTLS session
func (sess clientSession) CisMtlsSession() (*cismtlsv1.MtlsV1, error) {
	if sess.cisMtlsErr != nil {
		return sess.cisMtlsClient, sess.cisMtlsErr
	}
	return sess.cisMtlsClient.Clone(), nil
}

// CIS Webhooks
func (sess clientSession) CisWebhookSession() (*ciswebhooksv1.WebhooksV1, error) {
	if sess.cisWebhooksErr != nil {
		return sess.cisWebhooksClient, sess.cisWebhooksErr
	}
	return sess.cisWebhooksClient.Clone(), nil
}

// CIS Filters
func (sess clientSession) CisFiltersSession() (*cisfiltersv1.FiltersV1, error) {
	if sess.cisFiltersErr != nil {
		return sess.cisFiltersClient, sess.cisFiltersErr
	}
	return sess.cisFiltersClient.Clone(), nil
}

// CIS FirewallRules
func (sess clientSession) CisFirewallRulesSession() (*cisfirewallrulesv1.FirewallRulesV1, error) {
	if sess.cisFirewallRulesErr != nil {
		return sess.cisFirewallRulesClient, sess.cisFirewallRulesErr
	}
	return sess.cisFirewallRulesClient.Clone(), nil
}

// Activity Tracker API
func (session clientSession) AtrackerV1() (*atrackerv1.AtrackerV1, error) {
	return session.atrackerClient, session.atrackerClientErr
}

func (session clientSession) AtrackerV2() (*atrackerv2.AtrackerV2, error) {
	return session.atrackerClientV2, session.atrackerClientV2Err
}

func (session clientSession) ESschemaRegistrySession() (*schemaregistryv1.SchemaregistryV1, error) {
	return session.esSchemaRegistryClient, session.esSchemaRegistryErr
}

// Security and Compliance center Admin API
func (session clientSession) AdminServiceApiV1() (*adminserviceapiv1.AdminServiceApiV1, error) {
	return session.adminServiceApiClient, session.adminServiceApiClientErr
}

func (session clientSession) ConfigurationGovernanceV1() (*configurationgovernancev1.ConfigurationGovernanceV1, error) {
	return session.configServiceApiClient, session.configServiceApiClientErr
}

// Security and Compliance center Posture Management
func (session clientSession) PostureManagementV1() (*posturemanagementv1.PostureManagementV1, error) {
	if session.postureManagementClientErr != nil {
		return session.postureManagementClient, session.postureManagementClientErr
	}
	return session.postureManagementClient.Clone(), nil
}

// Security and Compliance center Posture Management v2
func (session clientSession) PostureManagementV2() (*posturemanagementv2.PostureManagementV2, error) {
	if session.postureManagementClientErrv2 != nil {
		return session.postureManagementClientv2, session.postureManagementClientErrv2
	}
	return session.postureManagementClientv2.Clone(), nil
}

// Context Based Restrictions
func (session clientSession) ContextBasedRestrictionsV1() (*contextbasedrestrictionsv1.ContextBasedRestrictionsV1, error) {
	return session.contextBasedRestrictionsClient, session.contextBasedRestrictionsClientErr
}

// CD Toolchain
func (session clientSession) CdToolchainV2() (*cdtoolchainv2.CdToolchainV2, error) {
	return session.cdToolchainClient, session.cdToolchainClientErr
}

// CD Tekton Pipeline
func (session clientSession) CdTektonPipelineV2() (*cdtektonpipelinev2.CdTektonPipelineV2, error) {
	return session.cdTektonPipelineClient, session.cdTektonPipelineClientErr
}

// Code Engine
func (session clientSession) CodeEngineV2() (*codeengine.CodeEngineV2, error) {
	return session.codeEngineClient, session.codeEngineClientErr
}

// ClientSession configures and returns a fully initialized ClientSession
func (c *Config) ClientSession() (interface{}, error) {
	sess, err := newSession(c)
	if err != nil {
		return nil, err
	}
	log.Printf("[INFO] Configured Region: %s\n", c.Region)
	session := clientSession{
		session: sess,
	}

	if sess.BluemixSession == nil {
		//Can be nil only  if bluemix_api_key is not provided
		log.Println("Skipping Bluemix Clients configuration")
		session.bluemixSessionErr = errEmptyBluemixCredentials
		session.accountConfigErr = errEmptyBluemixCredentials
		session.accountV1ConfigErr = errEmptyBluemixCredentials
		session.csConfigErr = errEmptyBluemixCredentials
		session.csv2ConfigErr = errEmptyBluemixCredentials
		session.containerRegistryClientErr = errEmptyBluemixCredentials
		session.kpErr = errEmptyBluemixCredentials
		session.pushServiceClientErr = errEmptyBluemixCredentials
		session.appConfigurationClientErr = errEmptyBluemixCredentials
		session.kmsErr = errEmptyBluemixCredentials
		session.cfConfigErr = errEmptyBluemixCredentials
		session.cisConfigErr = errEmptyBluemixCredentials
		session.functionConfigErr = errEmptyBluemixCredentials
		session.globalSearchConfigErr = errEmptyBluemixCredentials
		session.globalTaggingConfigErr = errEmptyBluemixCredentials
		session.globalTaggingConfigErrV1 = errEmptyBluemixCredentials
		session.hpcsEndpointErr = errEmptyBluemixCredentials
		session.iamAccessGroupsErr = errEmptyBluemixCredentials
		session.icdConfigErr = errEmptyBluemixCredentials
		session.resourceCatalogConfigErr = errEmptyBluemixCredentials
		session.resourceManagerErr = errEmptyBluemixCredentials
		session.resourceManagementConfigErrv2 = errEmptyBluemixCredentials
		session.resourceControllerConfigErr = errEmptyBluemixCredentials
		session.resourceControllerConfigErrv2 = errEmptyBluemixCredentials
		session.enterpriseManagementClientErr = errEmptyBluemixCredentials
		session.resourceControllerErr = errEmptyBluemixCredentials
		session.catalogManagementClientErr = errEmptyBluemixCredentials
		session.ibmpiConfigErr = errEmptyBluemixCredentials
		session.userManagementErr = errEmptyBluemixCredentials
		session.certManagementErr = errEmptyBluemixCredentials
		session.vpcErr = errEmptyBluemixCredentials
		session.vpcbetaErr = errEmptyBluemixCredentials
		session.apigatewayErr = errEmptyBluemixCredentials
		session.pDNSErr = errEmptyBluemixCredentials
		session.bmxUserFetchErr = errEmptyBluemixCredentials
		session.directlinkErr = errEmptyBluemixCredentials
		session.dlProviderErr = errEmptyBluemixCredentials
		session.cosConfigErr = errEmptyBluemixCredentials
		session.transitgatewayErr = errEmptyBluemixCredentials
		session.functionIAMNamespaceErr = errEmptyBluemixCredentials
		session.cisDNSErr = errEmptyBluemixCredentials
		session.cisAlertsErr = errEmptyBluemixCredentials
		session.cisDNSBulkErr = errEmptyBluemixCredentials
		session.cisGLBPoolErr = errEmptyBluemixCredentials
		session.cisGLBErr = errEmptyBluemixCredentials
		session.cisGLBHealthCheckErr = errEmptyBluemixCredentials
		session.cisIPErr = errEmptyBluemixCredentials
		session.cisZonesErr = errEmptyBluemixCredentials
		session.cisRLErr = errEmptyBluemixCredentials
		session.cisPageRuleErr = errEmptyBluemixCredentials
		session.cisEdgeFunctionErr = errEmptyBluemixCredentials
		session.cisSSLErr = errEmptyBluemixCredentials
		session.cisWAFPackageErr = errEmptyBluemixCredentials
		session.cisDomainSettingsErr = errEmptyBluemixCredentials
		session.cisRoutingErr = errEmptyBluemixCredentials
		session.cisWAFGroupErr = errEmptyBluemixCredentials
		session.cisCacheErr = errEmptyBluemixCredentials
		session.cisCustomPageErr = errEmptyBluemixCredentials
		session.cisMtlsErr = errEmptyBluemixCredentials
		session.cisAccessRuleErr = errEmptyBluemixCredentials
		session.cisUARuleErr = errEmptyBluemixCredentials
		session.cisLockdownErr = errEmptyBluemixCredentials
		session.cisRangeAppErr = errEmptyBluemixCredentials
		session.cisWAFRuleErr = errEmptyBluemixCredentials
		session.iamIdentityErr = errEmptyBluemixCredentials
		session.secretsManagerClientErr = errEmptyBluemixCredentials
		session.cisFiltersErr = errEmptyBluemixCredentials
		session.cisWebhooksErr = errEmptyBluemixCredentials
		session.cisLogpushJobsErr = errEmptyBluemixCredentials
		session.schematicsClientErr = errEmptyBluemixCredentials
		session.satelliteClientErr = errEmptyBluemixCredentials
		session.iamPolicyManagementErr = errEmptyBluemixCredentials
		session.satelliteLinkClientErr = errEmptyBluemixCredentials
		session.esSchemaRegistryErr = errEmptyBluemixCredentials
		session.contextBasedRestrictionsClientErr = errEmptyBluemixCredentials
		session.postureManagementClientErr = errEmptyBluemixCredentials
		session.postureManagementClientErrv2 = errEmptyBluemixCredentials
		session.configServiceApiClientErr = errEmptyBluemixCredentials
		session.cdTektonPipelineClientErr = errEmptyBluemixCredentials
		session.cdToolchainClientErr = errEmptyBluemixCredentials
		session.codeEngineClientErr = errEmptyBluemixCredentials

		return session, nil
	}

	if sess.BluemixSession.Config.BluemixAPIKey != "" {
		err = authenticateAPIKey(sess.BluemixSession)
		if err != nil {
			session.bmxUserFetchErr = fmt.Errorf("Error occured while fetching account user details: %q", err)
			for count := c.RetryCount; count >= 0; count-- {
				if err == nil || !isRetryable(err) {
					break
				}
				time.Sleep(c.RetryDelay)
				log.Printf("Retrying IAM Authentication %d", count)
				err = authenticateAPIKey(sess.BluemixSession)
			}
			if err != nil {
				session.bmxUserFetchErr = fmt.Errorf("[ERROR] Error occured while fetching auth key for account user details: %q", err)
				session.functionConfigErr = fmt.Errorf("[ERROR] Error occured while fetching auth key for function: %q", err)
			}
		}
		err = authenticateCF(sess.BluemixSession)
		if err != nil {
			for count := c.RetryCount; count >= 0; count-- {
				if err == nil || !isRetryable(err) {
					break
				}
				time.Sleep(c.RetryDelay)
				log.Printf("Retrying CF Authentication %d", count)
				err = authenticateCF(sess.BluemixSession)
			}
			if err != nil {
				session.functionConfigErr = fmt.Errorf("[ERROR] Error occured while fetching auth key for function: %q", err)
			}
		}
	}

	if c.IAMTrustedProfileID == "" && sess.BluemixSession.Config.IAMAccessToken != "" && sess.BluemixSession.Config.BluemixAPIKey == "" {
		err := RefreshToken(sess.BluemixSession)
		if err != nil {
			for count := c.RetryCount; count >= 0; count-- {
				if err == nil || !isRetryable(err) {
					break
				}
				time.Sleep(c.RetryDelay)
				log.Printf("Retrying refresh token %d", count)
				err = RefreshToken(sess.BluemixSession)
			}
			if err != nil {
				return nil, fmt.Errorf("[ERROR] Error occured while refreshing the token: %q", err)
			}
		}

	}
	userConfig, err := fetchUserDetails(sess.BluemixSession, c.RetryCount, c.RetryDelay)
	if err != nil {
		session.bmxUserFetchErr = fmt.Errorf("[ERROR] Error occured while fetching account user details: %q", err)
	}
	session.bmxUserDetails = userConfig

	if sess.SoftLayerSession != nil && sess.SoftLayerSession.APIKey == "" {
		log.Println("Configuring SoftLayer Session with token from IBM Cloud Session")
		sess.SoftLayerSession.IAMToken = sess.BluemixSession.Config.IAMAccessToken
		sess.SoftLayerSession.IAMRefreshToken = sess.BluemixSession.Config.IAMRefreshToken
	}

	session.functionClient, session.functionConfigErr = FunctionClient(sess.BluemixSession.Config)

	BluemixRegion = sess.BluemixSession.Config.Region
	var fileMap map[string]interface{}
	if f := EnvFallBack([]string{"IBMCLOUD_ENDPOINTS_FILE_PATH", "IC_ENDPOINTS_FILE_PATH"}, c.EndpointsFile); f != "" {
		jsonFile, err := os.Open(f)
		if err != nil {
			log.Fatalf("Unable to open Endpoints File %s", err)
		}
		defer jsonFile.Close()
		bytes, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			log.Fatalf("Unable to read Endpoints File %s", err)
		}
		err = json.Unmarshal([]byte(bytes), &fileMap)
		if err != nil {
			log.Fatalf("Unable to unmarshal Endpoints File %s", err)
		}
	}
	accv1API, err := accountv1.New(sess.BluemixSession)
	if err != nil {
		session.accountV1ConfigErr = fmt.Errorf("[ERROR] Error occured while configuring Bluemix Accountv1 Service: %q", err)
	}
	session.bmxAccountv1ServiceAPI = accv1API

	accAPI, err := accountv2.New(sess.BluemixSession)
	if err != nil {
		session.accountConfigErr = fmt.Errorf("[ERROR] Error occured while configuring  Account Service: %q", err)
	}
	session.bmxAccountServiceAPI = accAPI

	cfAPI, err := mccpv2.New(sess.BluemixSession)
	if err != nil {
		session.cfConfigErr = fmt.Errorf("[ERROR] Error occured while configuring MCCP service: %q", err)
	}
	session.cfServiceAPI = cfAPI

	clusterAPI, err := containerv1.New(sess.BluemixSession)
	if err != nil {
		session.csConfigErr = fmt.Errorf("[ERROR] Error occured while configuring Container Service for K8s cluster: %q", err)
	}
	session.csServiceAPI = clusterAPI

	v2clusterAPI, err := containerv2.New(sess.BluemixSession)
	if err != nil {
		session.csv2ConfigErr = fmt.Errorf("[ERROR] Error occured while configuring vpc Container Service for K8s cluster: %q", err)
	}
	session.csv2ServiceAPI = v2clusterAPI

	hpcsAPI, err := hpcs.New(sess.BluemixSession)
	if err != nil {
		session.hpcsEndpointErr = fmt.Errorf("[ERROR] Error occured while configuring hpcs Endpoint: %q", err)
	}
	session.hpcsEndpointAPI = hpcsAPI

	kpurl := ContructEndpoint(fmt.Sprintf("%s.kms", c.Region), cloudEndpoint)
	if c.Visibility == "private" || c.Visibility == "public-and-private" {
		kpurl = ContructEndpoint(fmt.Sprintf("private.%s.kms", c.Region), cloudEndpoint)
	}
	if fileMap != nil && c.Visibility != "public-and-private" {
		kpurl = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_KP_API_ENDPOINT", c.Region, kpurl)
	}
	var options kp.ClientConfig
	if c.BluemixAPIKey != "" {
		options = kp.ClientConfig{
			BaseURL: EnvFallBack([]string{"IBMCLOUD_KP_API_ENDPOINT"}, kpurl),
			APIKey:  sess.BluemixSession.Config.BluemixAPIKey, //pragma: allowlist secret
			// InstanceID:    "42fET57nnadurKXzXAedFLOhGqETfIGYxOmQXkFgkJV9",
			Verbose: kp.VerboseFailOnly,
		}

	} else {
		options = kp.ClientConfig{
			BaseURL:       EnvFallBack([]string{"IBMCLOUD_KP_API_ENDPOINT"}, kpurl),
			Authorization: sess.BluemixSession.Config.IAMAccessToken,
			// InstanceID:    "42fET57nnadurKXzXAedFLOhGqETfIGYxOmQXkFgkJV9",
			Verbose: kp.VerboseFailOnly,
		}
	}
	kpAPIclient, err := kp.New(options, DefaultTransport())
	if err != nil {
		session.kpErr = fmt.Errorf("[ERROR] Error occured while configuring Key Protect Service: %q", err)
	}
	session.kpAPI = kpAPIclient

	iamURL := iamidentity.DefaultServiceURL
	if c.Visibility == "private" || c.Visibility == "public-and-private" {
		if c.Region == "us-south" || c.Region == "us-east" {
			iamURL = ContructEndpoint(fmt.Sprintf("private.%s.iam", c.Region), cloudEndpoint)
		} else {
			iamURL = ContructEndpoint("private.iam", cloudEndpoint)
		}
	}
	if fileMap != nil && c.Visibility != "public-and-private" {
		iamURL = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_IAM_API_ENDPOINT", c.Region, iamURL)
	}

	// KEY MANAGEMENT Service
	kmsurl := ContructEndpoint(fmt.Sprintf("%s.kms", c.Region), cloudEndpoint)
	if c.Visibility == "private" || c.Visibility == "public-and-private" {
		kmsurl = ContructEndpoint(fmt.Sprintf("private.%s.kms", c.Region), cloudEndpoint)
	}
	if fileMap != nil && c.Visibility != "public-and-private" {
		kmsurl = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_KP_API_ENDPOINT", c.Region, kmsurl)
	}
	var kmsOptions kp.ClientConfig
	if c.BluemixAPIKey != "" {
		kmsOptions = kp.ClientConfig{
			BaseURL: EnvFallBack([]string{"IBMCLOUD_KP_API_ENDPOINT"}, kmsurl),
			APIKey:  sess.BluemixSession.Config.BluemixAPIKey, //pragma: allowlist secret
			// InstanceID:    "5af62d5d-5d90-4b84-bbcd-90d2123ae6c8",
			Verbose:  kp.VerboseFailOnly,
			TokenURL: EnvFallBack([]string{"IBMCLOUD_IAM_API_ENDPOINT"}, iamURL) + "/identity/token",
		}

	} else {
		kmsOptions = kp.ClientConfig{
			BaseURL:       EnvFallBack([]string{"IBMCLOUD_KP_API_ENDPOINT"}, kmsurl),
			Authorization: sess.BluemixSession.Config.IAMAccessToken,
			// InstanceID:    "5af62d5d-5d90-4b84-bbcd-90d2123ae6c8",
			Verbose:  kp.VerboseFailOnly,
			TokenURL: EnvFallBack([]string{"IBMCLOUD_IAM_API_ENDPOINT"}, iamURL) + "/identity/token",
		}
	}
	kmsAPIclient, err := kp.New(kmsOptions, DefaultTransport())
	if err != nil {
		session.kmsErr = fmt.Errorf("[ERROR] Error occured while configuring key Service: %q", err)
	}
	session.kmsAPI = kmsAPIclient

	var authenticator core.Authenticator

	if c.BluemixAPIKey != "" || sess.BluemixSession.Config.IAMRefreshToken != "" {
		if c.BluemixAPIKey != "" {
			authenticator = &core.IamAuthenticator{
				ApiKey: c.BluemixAPIKey,
				URL:    EnvFallBack([]string{"IBMCLOUD_IAM_API_ENDPOINT"}, iamURL),
			}
		} else {
			// Construct the IamAuthenticator with the IAM refresh token.
			authenticator = &core.IamAuthenticator{
				RefreshToken: sess.BluemixSession.Config.IAMRefreshToken,
				ClientId:     "bx",
				ClientSecret: "bx",
				URL:          EnvFallBack([]string{"IBMCLOUD_IAM_API_ENDPOINT"}, iamURL),
			}
		}
	} else if strings.HasPrefix(sess.BluemixSession.Config.IAMAccessToken, "Bearer") {
		authenticator = &core.BearerTokenAuthenticator{
			BearerToken: sess.BluemixSession.Config.IAMAccessToken[7:],
		}
	} else {
		authenticator = &core.BearerTokenAuthenticator{
			BearerToken: sess.BluemixSession.Config.IAMAccessToken,
		}
	}

	// Construct an "options" struct for creating the service client.
	ukoClientOptions := &ukov4.UkoV4Options{
		Authenticator: authenticator,
	}

	// Construct the service client.
	session.ukoClient, err = ukov4.NewUkoV4(ukoClientOptions)
	if err == nil {
		// Enable retries for API calls
		session.ukoClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		// Add custom header for analytics
		session.ukoClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	} else {
		session.ukoClientErr = fmt.Errorf("Error occurred while configuring HPCS UKO service: %q", err)
	}

	// APPID Service
	appIDEndpoint := fmt.Sprintf("https://%s.appid.cloud.ibm.com", c.Region)
	if c.Visibility == "private" {
		session.appidErr = fmt.Errorf("App Id resources doesnot support private endpoints")
	}
	if fileMap != nil && c.Visibility != "public-and-private" {
		appIDEndpoint = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_APPID_MANAGEMENT_API_ENDPOINT", c.Region, appIDEndpoint)
	}
	appIDClientOptions := &appid.AppIDManagementV4Options{
		Authenticator: authenticator,
		URL:           EnvFallBack([]string{"IBMCLOUD_APPID_MANAGEMENT_API_ENDPOINT"}, appIDEndpoint),
	}
	appIDClient, err := appid.NewAppIDManagementV4(appIDClientOptions)
	if err != nil {
		session.appidErr = fmt.Errorf("error occured while configuring AppID service: #{err}")
	}
	if appIDClient != nil && appIDClient.Service != nil {
		appIDClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		appIDClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}
	session.appidAPI = appIDClient

	// Construct an "options" struct for creating Context Based Restrictions service client.
	cbrURL := contextbasedrestrictionsv1.DefaultServiceURL
	if c.Visibility == "private" || c.Visibility == "public-and-private" {
		if c.Region == "us-south" || c.Region == "us-east" || c.Region == "eu-de" {
			cbrURL = ContructEndpoint(fmt.Sprintf("private.%s.cbr", c.Region), cloudEndpoint)
		} else {
			cbrURL = ContructEndpoint("private.cbr", cloudEndpoint)
		}
	}
	if fileMap != nil && c.Visibility != "public-and-private" {
		cbrURL = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_CONTEXT_BASED_RESTRICTIONS_ENDPOINT", c.Region, cbrURL)
	}
	contextBasedRestrictionsClientOptions := &contextbasedrestrictionsv1.ContextBasedRestrictionsV1Options{
		Authenticator: authenticator,
		URL:           EnvFallBack([]string{"IBMCLOUD_CONTEXT_BASED_RESTRICTIONS_ENDPOINT"}, cbrURL),
	}

	// Construct the service client.
	session.contextBasedRestrictionsClient, err = contextbasedrestrictionsv1.NewContextBasedRestrictionsV1(contextBasedRestrictionsClientOptions)
	if err == nil && session.contextBasedRestrictionsClient != nil {
		// Enable retries for API calls
		session.contextBasedRestrictionsClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		// Add custom header for analytics
		session.contextBasedRestrictionsClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	} else {
		session.contextBasedRestrictionsClientErr = fmt.Errorf("[ERROR] Error occurred while configuring Context Based Restrictions service: %q", err)
	}

	// CATALOG MANAGEMENT Service
	catalogManagementURL := "https://cm.globalcatalog.cloud.ibm.com/api/v1-beta"
	if c.Visibility == "private" {
		session.catalogManagementClientErr = fmt.Errorf("Catalog Management resource doesnot support private endpoints")
	}
	if fileMap != nil && c.Visibility != "public-and-private" {
		catalogManagementURL = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_CATALOG_MANAGEMENT_API_ENDPOINT", c.Region, catalogManagementURL)
	}
	catalogManagementClientOptions := &catalogmanagementv1.CatalogManagementV1Options{
		URL:           EnvFallBack([]string{"IBMCLOUD_CATALOG_MANAGEMENT_API_ENDPOINT"}, catalogManagementURL),
		Authenticator: authenticator,
	}
	// Construct the service client.
	session.catalogManagementClient, err = catalogmanagementv1.NewCatalogManagementV1(catalogManagementClientOptions)
	if err != nil {
		session.catalogManagementClientErr = fmt.Errorf("[ERROR] Error occurred while configuring Catalog Management API service: %q", err)
	}
	if session.catalogManagementClient != nil && session.catalogManagementClient.Service != nil {
		// Enable retries for API calls
		session.catalogManagementClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		// Add custom header for analytics
		session.catalogManagementClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// ATRACKER Service
	var atrackerClientURL string
	var atrackerURLErr error

	atrackerClientURL, atrackerURLErr = atrackerv1.GetServiceURLForRegion(c.Region)
	if c.Visibility == "private" || c.Visibility == "public-and-private" {
		atrackerClientURL, atrackerURLErr = atrackerv1.GetServiceURLForRegion("private." + c.Region)
		if err != nil && c.Visibility == "public-and-private" {
			atrackerClientURL, atrackerURLErr = atrackerv1.GetServiceURLForRegion(c.Region)
		}
	}

	if fileMap != nil && c.Visibility != "public-and-private" {
		atrackerClientURL = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_ATRACKER_API_ENDPOINT", c.Region, atrackerClientURL)
	}
	atrackerClientOptions := &atrackerv1.AtrackerV1Options{
		Authenticator: authenticator,
		URL:           EnvFallBack([]string{"IBMCLOUD_ATRACKER_API_ENDPOINT"}, atrackerClientURL),
	}
	// If we provide IBMCLOUD_ATRACKER_API_ENDPOINT, then ignore any missing region url
	if atrackerURLErr != nil && len(atrackerClientOptions.URL) == 0 {
		session.atrackerClientErr = atrackerURLErr
	}
	// Construct the service client.
	session.atrackerClient, err = atrackerv1.NewAtrackerV1(atrackerClientOptions)
	if err != nil {
		session.atrackerClientErr = fmt.Errorf("[ERROR] Error occurred while configuring Activity Tracker API service: %q", err)
	}
	if session.atrackerClient != nil && session.atrackerClient.Service != nil {
		// Enable retries for API calls
		session.atrackerClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		// Add custom header for analytics
		session.atrackerClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}
	// Version 2 Atracker
	var atrackerClientV2URL string
	var atrackerURLV2Err error

	if c.Visibility == "private" || c.Visibility == "public-and-private" {
		atrackerClientV2URL, atrackerURLV2Err = atrackerv2.GetServiceURLForRegion("private." + c.Region)
		if err != nil && c.Visibility == "public-and-private" {
			atrackerClientV2URL, atrackerURLV2Err = atrackerv2.GetServiceURLForRegion(c.Region)
		}
	} else {
		atrackerClientV2URL, atrackerURLV2Err = atrackerv2.GetServiceURLForRegion(c.Region)
	}
	if atrackerURLV2Err != nil {
		atrackerClientV2URL = atrackerv2.DefaultServiceURL
	}
	if fileMap != nil && c.Visibility != "public-and-private" {
		atrackerClientV2URL = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_ATRACKER_API_ENDPOINT", c.Region, atrackerClientV2URL)
	}
	atrackerClientV2Options := &atrackerv2.AtrackerV2Options{
		Authenticator: authenticator,
		URL:           EnvFallBack([]string{"IBMCLOUD_ATRACKER_API_ENDPOINT"}, atrackerClientV2URL),
	}
	// If we provide IBMCLOUD_ATRACKER_API_ENDPOINT, then ignore any missing region url, or should use the default.
	// This should technically never happen as we default this for v2
	if atrackerURLV2Err != nil && len(atrackerClientOptions.URL) == 0 {
		session.atrackerClientErr = atrackerURLErr
	}
	session.atrackerClientV2, err = atrackerv2.NewAtrackerV2(atrackerClientV2Options)
	if err == nil {
		// Enable retries for API calls
		session.atrackerClientV2.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		// Add custom header for analytics
		session.atrackerClientV2.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	} else {
		session.atrackerClientV2Err = fmt.Errorf("Error occurred while configuring Activity Tracker API Version 2 service: %q", err)
	}

	// SCC ADMIN Service
	var adminServiceApiClientURL string
	if c.Visibility == "private" || c.Visibility == "public-and-private" {
		adminServiceApiClientURL, err = adminserviceapiv1.GetServiceURLForRegion("private." + c.Region)
		if err != nil && c.Visibility == "public-and-private" {
			adminServiceApiClientURL, err = adminserviceapiv1.GetServiceURLForRegion(c.Region)
		}
	} else {
		adminServiceApiClientURL, err = adminserviceapiv1.GetServiceURLForRegion(c.Region)
	}
	if err != nil {
		adminServiceApiClientURL = adminserviceapiv1.DefaultServiceURL
	}
	adminServiceApiClientOptions := &adminserviceapiv1.AdminServiceApiV1Options{
		Authenticator: authenticator,
		URL:           EnvFallBack([]string{"IBMCLOUD_SCC_ADMIN_API_ENDPOINT"}, adminServiceApiClientURL),
	}

	// Construct the service client.
	session.adminServiceApiClient, err = adminserviceapiv1.NewAdminServiceApiV1(adminServiceApiClientOptions)
	if err == nil {
		// Enable retries for API calls
		session.adminServiceApiClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		// Add custom header for analytics
		session.adminServiceApiClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	} else {
		session.adminServiceApiClientErr = fmt.Errorf("[ERROR] Error occurred while configuring Admin Service API service: %q", err)
	}

	// SCHEMATICS Service
	// schematicsEndpoint := "https://schematics.cloud.ibm.com"
	schematicsEndpoint := ContructEndpoint(fmt.Sprintf("%s.schematics", c.Region), cloudEndpoint)
	if c.Visibility == "private" || c.Visibility == "public-and-private" {
		schematicsEndpoint = ContructEndpoint(fmt.Sprintf("private-%s.schematics", c.Region), cloudEndpoint)
	}
	if fileMap != nil && c.Visibility != "public-and-private" {
		schematicsEndpoint = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_SCHEMATICS_API_ENDPOINT", c.Region, schematicsEndpoint)
	}
	schematicsClientOptions := &schematicsv1.SchematicsV1Options{
		Authenticator: authenticator,
		URL:           EnvFallBack([]string{"IBMCLOUD_SCHEMATICS_API_ENDPOINT"}, schematicsEndpoint),
	}
	// Construct the service client.
	schematicsClient, err := schematicsv1.NewSchematicsV1(schematicsClientOptions)
	if err != nil {
		session.schematicsClientErr = fmt.Errorf("[ERROR] Error occurred while configuring Schematics Service API service: %q", err)
	}
	// Enable retries for API calls
	if schematicsClient != nil && schematicsClient.Service != nil {
		schematicsClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		schematicsClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}
	session.schematicsClient = schematicsClient

	// VPC Service
	vpcurl := ContructEndpoint(fmt.Sprintf("%s.iaas", c.Region), fmt.Sprintf("%s/v1", cloudEndpoint))
	if c.Visibility == "private" || c.Visibility == "public-and-private" {
		vpcurl = ContructEndpoint(fmt.Sprintf("%s.private.iaas", c.Region), fmt.Sprintf("%s/v1", cloudEndpoint))
	}
	if fileMap != nil && c.Visibility != "public-and-private" {
		vpcurl = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_IS_NG_API_ENDPOINT", c.Region, vpcurl)
	}
	vpcoptions := &vpc.VpcV1Options{
		URL:           EnvFallBack([]string{"IBMCLOUD_IS_NG_API_ENDPOINT"}, vpcurl),
		Authenticator: authenticator,
	}
	vpcclient, err := vpc.NewVpcV1(vpcoptions)
	if err != nil {
		session.vpcErr = fmt.Errorf("[ERROR] Error occured while configuring vpc service: %q", err)
	}
	if vpcclient != nil && vpcclient.Service != nil {
		vpcclient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		vpcclient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}
	session.vpcAPI = vpcclient

	vpcbetaoptions := &vpcbeta.VpcbetaV1Options{
		URL:           EnvFallBack([]string{"IBMCLOUD_IS_NG_API_ENDPOINT"}, vpcurl),
		Authenticator: authenticator,
	}
	vpcbetaclient, err := vpcbeta.NewVpcbetaV1(vpcbetaoptions)
	if err != nil {
		session.vpcbetaErr = fmt.Errorf("[ERROR] Error occured while configuring vpc beta service: %q", err)
	}
	if vpcbetaclient != nil && vpcbetaclient.Service != nil {
		vpcbetaclient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		vpcbetaclient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}
	session.vpcBetaAPI = vpcbetaclient

	// PUSH NOTIFICATIONS Service
	pnurl := fmt.Sprintf("https://%s.imfpush.cloud.ibm.com/imfpush/v1", c.Region)
	if c.Visibility == "private" {
		session.pushServiceClientErr = fmt.Errorf("Push Notifications Service API doesnot support private endpoints")
	}
	if fileMap != nil && c.Visibility != "public-and-private" {
		pnurl = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_PUSH_API_ENDPOINT", c.Region, pnurl)
	}
	pushNotificationOptions := &pushservicev1.PushServiceV1Options{
		URL:           EnvFallBack([]string{"IBMCLOUD_PUSH_API_ENDPOINT"}, pnurl),
		Authenticator: authenticator,
	}
	pnclient, err := pushservicev1.NewPushServiceV1(pushNotificationOptions)
	if err != nil {
		session.pushServiceClientErr = fmt.Errorf("[ERROR] Error occured while configuring Push Notifications service: %q", err)
	}
	if pnclient != nil && pnclient.Service != nil {
		// Enable retries for API calls
		pnclient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		pnclient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}
	session.pushServiceClient = pnclient
	// event notifications
	enurl := fmt.Sprintf("https://%s.event-notifications.cloud.ibm.com/event-notifications", c.Region)
	if c.Visibility == "private" {
		session.eventNotificationsApiClientErr = fmt.Errorf("Event Notifications Service does not support private endpoints")
	}
	if fileMap != nil && c.Visibility != "public-and-private" {
		enurl = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_EVENT_NOTIFICATIONS_API_ENDPOINT", c.Region, enurl)
	}
	enClientOptions := &eventnotificationsv1.EventNotificationsV1Options{
		Authenticator: authenticator,
		URL:           EnvFallBack([]string{"IBMCLOUD_EVENT_NOTIFICATIONS_API_ENDPOINT"}, enurl),
	}
	// Construct the service client.
	session.eventNotificationsApiClient, err = eventnotificationsv1.NewEventNotificationsV1(enClientOptions)
	if err != nil {
		// Enable {
		session.eventNotificationsApiClientErr = fmt.Errorf("[ERROR] Error occurred while configuring Event Notifications service: %q", err)
	}
	if session.eventNotificationsApiClient != nil && session.eventNotificationsApiClient.Service != nil {
		// Enable retries for API calls
		session.eventNotificationsApiClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.eventNotificationsApiClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// APP CONFIGURATION Service
	appconfigurl := ContructEndpoint(fmt.Sprintf("%s", c.Region), fmt.Sprintf("%s.apprapp.", cloudEndpoint))
	if c.Visibility == "private" || c.Visibility == "public-and-private" {
		appconfigurl = ContructEndpoint(fmt.Sprintf("%s.private", c.Region), fmt.Sprintf("%s.apprapp", cloudEndpoint))
	}
	if fileMap != nil && c.Visibility != "public-and-private" {
		appconfigurl = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_APP_CONFIG_ENDPOINT", c.Region, appconfigurl)
	}
	appConfigurationClientOptions := &appconfigurationv1.AppConfigurationV1Options{
		URL:           EnvFallBack([]string{"IBMCLOUD_APP_CONFIG_ENDPOINT"}, appconfigurl),
		Authenticator: authenticator,
	}

	appConfigClient, err := appconfigurationv1.NewAppConfigurationV1(appConfigurationClientOptions)
	if appConfigClient != nil {
		// Enable retries for API calls
		appConfigClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.appConfigurationClient = appConfigClient
	} else {
		session.appConfigurationClientErr = fmt.Errorf("[ERROR] Error occurred while configuring App Configuration service: %q", err)
	}

	// CONTAINER REGISTRY Service
	// Construct an "options" struct for creating the service client.
	containerRegistryClientURL, err := containerregistryv1.GetServiceURLForRegion(c.Region)
	if err != nil {
		containerRegistryClientURL = containerregistryv1.DefaultServiceURL
	}
	if c.Visibility == "private" || c.Visibility == "public-and-private" {
		containerRegistryClientURL, err = GetPrivateServiceURLForRegion(c.Region)
		if err != nil {
			containerRegistryClientURL, _ = GetPrivateServiceURLForRegion("global")
		}
	}
	if fileMap != nil && c.Visibility != "public-and-private" {
		containerRegistryClientURL = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_CR_API_ENDPOINT", c.Region, containerRegistryClientURL)
	}
	containerRegistryClientOptions := &containerregistryv1.ContainerRegistryV1Options{
		Authenticator: authenticator,
		URL:           EnvFallBack([]string{"IBMCLOUD_CR_API_ENDPOINT"}, containerRegistryClientURL),
		Account:       core.StringPtr(userConfig.UserAccount),
	}
	// Construct the service client.
	session.containerRegistryClient, err = containerregistryv1.NewContainerRegistryV1(containerRegistryClientOptions)
	if err != nil {
		session.containerRegistryClientErr = fmt.Errorf("[ERROR] Error occurred while configuring IBM Cloud Container Registry API service: %q", err)
	}
	if session.containerRegistryClient != nil && session.containerRegistryClient.Service != nil {
		// Enable retries for API calls
		session.containerRegistryClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		// Add custom header for analytics
		session.containerRegistryClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// OBJECT STORAGE Service
	cosconfigurl := "https://config.cloud-object-storage.cloud.ibm.com/v1"
	if fileMap != nil && c.Visibility != "public-and-private" {
		cosconfigurl = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_COS_CONFIG_ENDPOINT", c.Region, cosconfigurl)
	}
	cosconfigoptions := &cosconfig.ResourceConfigurationV1Options{
		Authenticator: authenticator,
		URL:           EnvFallBack([]string{"IBMCLOUD_COS_CONFIG_ENDPOINT"}, cosconfigurl),
	}
	cosconfigclient, err := cosconfig.NewResourceConfigurationV1(cosconfigoptions)
	if err != nil {
		session.cosConfigErr = fmt.Errorf("[ERROR] Error occured while configuring COS config service: %q", err)
	}
	session.cosConfigAPI = cosconfigclient

	globalSearchAPI, err := globalsearchv2.New(sess.BluemixSession)
	if err != nil {
		session.globalSearchConfigErr = fmt.Errorf("[ERROR] Error occured while configuring Global Search: %q", err)
	}
	session.globalSearchServiceAPI = globalSearchAPI
	// Global Tagging Bluemix-go
	globalTaggingAPI, err := globaltaggingv3.New(sess.BluemixSession)
	if err != nil {
		session.globalTaggingConfigErr = fmt.Errorf("[ERROR] Error occured while configuring Global Tagging: %q", err)
	}
	session.globalTaggingServiceAPI = globalTaggingAPI

	// GLOBAL TAGGING Service
	globalTaggingEndpoint := "https://tags.global-search-tagging.cloud.ibm.com"
	if c.Visibility == "private" || c.Visibility == "public-and-private" {
		var globalTaggingRegion string
		if c.Region != "us-south" && c.Region != "us-east" {
			globalTaggingRegion = "us-south"
		} else {
			globalTaggingRegion = c.Region
		}
		globalTaggingEndpoint = ContructEndpoint(fmt.Sprintf("tags.private.%s", globalTaggingRegion), fmt.Sprintf("global-search-tagging.%s", cloudEndpoint))
	}
	if fileMap != nil && c.Visibility != "public-and-private" {
		globalTaggingEndpoint = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_GT_API_ENDPOINT", c.Region, globalTaggingEndpoint)
	}
	globalTaggingV1Options := &globaltaggingv1.GlobalTaggingV1Options{
		URL:           EnvFallBack([]string{"IBMCLOUD_GT_API_ENDPOINT"}, globalTaggingEndpoint),
		Authenticator: authenticator,
	}
	globalTaggingAPIV1, err := globaltaggingv1.NewGlobalTaggingV1(globalTaggingV1Options)
	if err != nil {
		session.globalTaggingConfigErrV1 = fmt.Errorf("[ERROR] Error occured while configuring Global Tagging: %q", err)
	}
	if globalTaggingAPIV1 != nil && globalTaggingAPIV1.Service != nil {
		session.globalTaggingServiceAPIV1 = *globalTaggingAPIV1
		session.globalTaggingServiceAPIV1.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.globalTaggingServiceAPIV1.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}
	// GLOBAL TAGGING Service
	globalSearchEndpoint := "https://api.global-search-tagging.cloud.ibm.com"
	if c.Visibility == "private" || c.Visibility == "public-and-private" {
		globalSearchEndpoint = ContructEndpoint("api.private", fmt.Sprintf("global-search-tagging.%s", cloudEndpoint))
	}
	if fileMap != nil && c.Visibility != "public-and-private" {
		globalSearchEndpoint = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_GS_API_ENDPOINT", c.Region, searchv2.DefaultServiceURL)
	}
	globalSearchV2Options := &searchv2.GlobalSearchV2Options{
		URL:           EnvFallBack([]string{"IBMCLOUD_GS_API_ENDPOINT"}, globalSearchEndpoint),
		Authenticator: authenticator,
	}
	globalSearchAPIV2, err := searchv2.NewGlobalSearchV2(globalSearchV2Options)
	if err != nil {
		session.globalTaggingConfigErrV1 = fmt.Errorf("[ERROR] Error occured while configuring Global Search: %q", err)
	}
	if globalSearchAPIV2 != nil && globalSearchAPIV2.Service != nil {
		session.globalSearchServiceAPIV2 = *globalSearchAPIV2
		session.globalSearchServiceAPIV2.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.globalSearchServiceAPIV2.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	icdAPI, err := icdv4.New(sess.BluemixSession)
	if err != nil {
		session.icdConfigErr = fmt.Errorf("[ERROR] Error occured while configuring IBM Cloud Database Services: %q", err)
	}
	session.icdServiceAPI = icdAPI

	var cloudDatabasesEndpoint string

	if c.Visibility == "private" || c.Visibility == "public-and-private" {
		cloudDatabasesEndpoint = fmt.Sprintf("https://api.%s.private.databases.cloud.ibm.com/v5/ibm", c.Region)
	} else {
		cloudDatabasesEndpoint = fmt.Sprintf("https://api.%s.databases.cloud.ibm.com/v5/ibm", c.Region)
	}

	// Construct an "options" struct for creating the service client.
	cloudDatabasesClientOptions := &clouddatabasesv5.CloudDatabasesV5Options{
		URL:           EnvFallBack([]string{"IBMCLOUD_DATABASES_API_ENDPOINT"}, cloudDatabasesEndpoint),
		Authenticator: authenticator,
	}

	// Construct the service client.
	session.cloudDatabasesClient, err = clouddatabasesv5.NewCloudDatabasesV5(cloudDatabasesClientOptions)
	if err == nil {
		// Enable retries for API calls
		session.cloudDatabasesClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		// Add custom header for analytics
		session.cloudDatabasesClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	} else {
		session.cloudDatabasesClientErr = fmt.Errorf("Error occurred while configuring The IBM Cloud Databases API service: %q", err)
	}

	resourceCatalogAPI, err := catalog.New(sess.BluemixSession)
	if err != nil {
		session.resourceCatalogConfigErr = fmt.Errorf("[ERROR] Error occured while configuring Resource Catalog service: %q", err)
	}
	session.resourceCatalogServiceAPI = resourceCatalogAPI

	resourceManagementAPIv2, err := managementv2.New(sess.BluemixSession)
	if err != nil {
		session.resourceManagementConfigErrv2 = fmt.Errorf("[ERROR] Error occured while configuring Resource Management service: %q", err)
	}
	session.resourceManagementServiceAPIv2 = resourceManagementAPIv2

	resourceControllerAPI, err := controller.New(sess.BluemixSession)
	if err != nil {
		session.resourceControllerConfigErr = fmt.Errorf("[ERROR] Error occured while configuring Resource Controller service: %q", err)
	}
	session.resourceControllerServiceAPI = resourceControllerAPI

	ResourceControllerAPIv2, err := controllerv2.New(sess.BluemixSession)
	if err != nil {
		session.resourceControllerConfigErrv2 = fmt.Errorf("[ERROR] Error occured while configuring Resource Controller v2 service: %q", err)
	}
	session.resourceControllerServiceAPIv2 = ResourceControllerAPIv2

	userManagementAPI, err := usermanagementv2.New(sess.BluemixSession)
	if err != nil {
		session.userManagementErr = fmt.Errorf("[ERROR] Error occured while configuring user management service: %q", err)
	}
	session.userManagementAPI = userManagementAPI

	certManagementAPI, err := certificatemanager.New(sess.BluemixSession)
	if err != nil {
		session.certManagementErr = fmt.Errorf("[ERROR] Error occured while configuring Certificate manager service: %q", err)
	}
	session.certManagementAPI = certManagementAPI

	namespaceFunction, err := functions.New(sess.BluemixSession)
	if err != nil {
		session.functionIAMNamespaceErr = fmt.Errorf("[ERROR] Error occured while configuring Cloud Funciton Service : %q", err)
	}
	session.functionIAMNamespaceAPI = namespaceFunction

	//  API GATEWAY service
	apicurl := ContructEndpoint(fmt.Sprintf("api.%s.apigw", c.Region), fmt.Sprintf("%s/controller", cloudEndpoint))
	if c.Visibility == "private" || c.Visibility == "public-and-private" {
		apicurl = ContructEndpoint(fmt.Sprintf("api.private.%s.apigw", c.Region), fmt.Sprintf("%s/controller", cloudEndpoint))
	}
	if fileMap != nil && c.Visibility != "public-and-private" {
		apicurl = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_API_GATEWAY_ENDPOINT", c.Region, apicurl)
	}
	APIGatewayControllerAPIV1Options := &apigateway.ApiGatewayControllerApiV1Options{
		URL:           EnvFallBack([]string{"IBMCLOUD_API_GATEWAY_ENDPOINT"}, apicurl),
		Authenticator: &core.NoAuthAuthenticator{},
	}
	apigatewayAPI, err := apigateway.NewApiGatewayControllerApiV1(APIGatewayControllerAPIV1Options)
	if err != nil {
		session.apigatewayErr = fmt.Errorf("[ERROR] Error occured while configuring  APIGateway service: %q", err)
	}
	session.apigatewayAPI = apigatewayAPI

	// POWER SYSTEMS Service
	piURL := ContructEndpoint(c.Region, "power-iaas.cloud.ibm.com")
	ibmPIOptions := &ibmpisession.IBMPIOptions{
		Authenticator: authenticator,
		Debug:         os.Getenv("TF_LOG") != "",
		Region:        c.Region,
		URL:           EnvFallBack([]string{"IBMCLOUD_PI_API_ENDPOINT"}, piURL),
		UserAccount:   userConfig.UserAccount,
		Zone:          c.Zone,
	}
	ibmpisession, err := ibmpisession.NewIBMPISession(ibmPIOptions)
	if err != nil {
		session.ibmpiConfigErr = fmt.Errorf("Error occured while configuring ibmpisession: %q", err)
	}
	session.ibmpiSession = ibmpisession

	// PRIVATE DNS Service
	pdnsURL := dns.DefaultServiceURL
	if c.Visibility == "private" || c.Visibility == "public-and-private" {
		pdnsURL = ContructEndpoint("api.private.dns-svcs", fmt.Sprintf("%s/v1", cloudEndpoint))
	}
	if fileMap != nil && c.Visibility != "public-and-private" {
		pdnsURL = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_PRIVATE_DNS_API_ENDPOINT", c.Region, pdnsURL)
	}
	dnsOptions := &dns.DnsSvcsV1Options{
		URL:           EnvFallBack([]string{"IBMCLOUD_PRIVATE_DNS_API_ENDPOINT"}, pdnsURL),
		Authenticator: authenticator,
	}
	session.pDNSClient, session.pDNSErr = dns.NewDnsSvcsV1(dnsOptions)
	if session.pDNSErr != nil {
		session.pDNSErr = fmt.Errorf("[ERROR] Error occured while configuring PrivateDNS Service: %s", session.pDNSErr)
	}
	if session.pDNSClient != nil && session.pDNSClient.Service != nil {
		session.pDNSClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.pDNSClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// DIRECT LINK Service
	ver := time.Now().Format("2006-01-02")
	dlURL := dl.DefaultServiceURL
	if c.Visibility == "private" || c.Visibility == "public-and-private" {
		dlURL = ContructEndpoint("private.directlink", fmt.Sprintf("%s/v1", cloudEndpoint))
	}
	if fileMap != nil && c.Visibility != "public-and-private" {
		dlURL = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_DL_API_ENDPOINT", c.Region, dlURL)
	}
	directlinkOptions := &dl.DirectLinkV1Options{
		URL:           EnvFallBack([]string{"IBMCLOUD_DL_API_ENDPOINT"}, dlURL),
		Authenticator: authenticator,
		Version:       &ver,
	}
	session.directlinkAPI, session.directlinkErr = dl.NewDirectLinkV1(directlinkOptions)
	if session.directlinkErr != nil {
		session.directlinkErr = fmt.Errorf("[ERROR] Error occured while configuring Direct Link Service: %s", session.directlinkErr)
	}
	if session.directlinkAPI != nil && session.directlinkAPI.Service != nil {
		session.directlinkAPI.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.directlinkAPI.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// DIRECT LINK PROVIDER Service
	dlproviderURL := dlProviderV2.DefaultServiceURL
	if c.Visibility == "private" || c.Visibility == "public-and-private" {
		dlproviderURL = ContructEndpoint("private.directlink", fmt.Sprintf("%s/provider/v2", cloudEndpoint))
	}
	if fileMap != nil && c.Visibility != "public-and-private" {
		dlproviderURL = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_DL_PROVIDER_API_ENDPOINT", c.Region, dlproviderURL)
	}
	directLinkProviderV2Options := &dlProviderV2.DirectLinkProviderV2Options{
		URL:           EnvFallBack([]string{"IBMCLOUD_DL_PROVIDER_API_ENDPOINT"}, dlproviderURL),
		Authenticator: authenticator,
		Version:       &ver,
	}
	session.dlProviderAPI, session.dlProviderErr = dlProviderV2.NewDirectLinkProviderV2(directLinkProviderV2Options)
	if session.dlProviderErr != nil {
		session.dlProviderErr = fmt.Errorf("[ERROR] Error occured while configuring Direct Link Provider Service: %s", session.dlProviderErr)
	}
	if session.dlProviderAPI != nil && session.dlProviderAPI.Service != nil {
		session.dlProviderAPI.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.dlProviderAPI.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// TRANSIT GATEWAY Service
	tgURL := tg.DefaultServiceURL
	if c.Visibility == "private" || c.Visibility == "public-and-private" {
		tgURL = ContructEndpoint("private.transit", fmt.Sprintf("%s/v1", cloudEndpoint))
	}
	if fileMap != nil && c.Visibility != "public-and-private" {
		tgURL = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_TG_API_ENDPOINT", c.Region, tgURL)
	}
	transitgatewayOptions := &tg.TransitGatewayApisV1Options{
		URL:           EnvFallBack([]string{"IBMCLOUD_TG_API_ENDPOINT"}, tgURL),
		Authenticator: authenticator,
		Version:       CreateVersionDate(),
	}
	session.transitgatewayAPI, session.transitgatewayErr = tg.NewTransitGatewayApisV1(transitgatewayOptions)
	if session.transitgatewayErr != nil {
		session.transitgatewayErr = fmt.Errorf("[ERROR] Error occured while configuring Transit Gateway Service: %s", session.transitgatewayErr)
	}
	if session.transitgatewayAPI != nil && session.transitgatewayAPI.Service != nil {
		session.transitgatewayAPI.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		// session.transitgatewayAPI.SetDefaultHeaders(gohttp.Header{
		// 	"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		// })
	}

	// CIS Service instances starts here.
	cisURL := ContructEndpoint("api.cis", cloudEndpoint)
	if c.Visibility == "private" {
		// cisURL = ContructEndpoint("api.private.cis", cloudEndpoint)
		session.cisZonesErr = fmt.Errorf("CIS Service doesnt support private endpoints.")
		session.cisDNSBulkErr = fmt.Errorf("CIS Service doesnt support private endpoints.")
		session.cisGLBPoolErr = fmt.Errorf("CIS Service doesnt support private endpoints.")
		session.cisGLBErr = fmt.Errorf("CIS Service doesnt support private endpoints.")
		session.cisGLBHealthCheckErr = fmt.Errorf("CIS Service doesnt support private endpoints.")
		session.cisIPErr = fmt.Errorf("CIS Service doesnt support private endpoints.")
		session.cisRLErr = fmt.Errorf("CIS Service doesnt support private endpoints.")
		session.cisPageRuleErr = fmt.Errorf("CIS Service doesnt support private endpoints.")
		session.cisEdgeFunctionErr = fmt.Errorf("CIS Service doesnt support private endpoints.")
		session.cisSSLErr = fmt.Errorf("CIS Service doesnt support private endpoints.")
		session.cisWAFPackageErr = fmt.Errorf("CIS Service doesnt support private endpoints.")
		session.cisDomainSettingsErr = fmt.Errorf("CIS Service doesnt support private endpoints.")
		session.cisRoutingErr = fmt.Errorf("CIS Service doesnt support private endpoints.")
		session.cisWAFGroupErr = fmt.Errorf("CIS Service doesnt support private endpoints.")
		session.cisCacheErr = fmt.Errorf("CIS Service doesnt support private endpoints.")
		session.cisCustomPageErr = fmt.Errorf("CIS Service doesnt support private endpoints.")
		session.cisAccessRuleErr = fmt.Errorf("CIS Service doesnt support private endpoints.")
		session.cisUARuleErr = fmt.Errorf("CIS Service doesnt support private endpoints.")
		session.cisLockdownErr = fmt.Errorf("CIS Service doesnt support private endpoints.")
		session.cisRangeAppErr = fmt.Errorf("CIS Service doesnt support private endpoints.")
		session.cisWAFRuleErr = fmt.Errorf("CIS Service doesnt support private endpoints.")
		session.cisFiltersErr = fmt.Errorf("CIS Service doesnt support private endpoints.")
		session.cisWebhooksErr = fmt.Errorf("CIS Service doesnt support private endpoints.")
		session.cisMtlsErr = fmt.Errorf("CIS Service doesnt support private endpoints.")

	}
	if fileMap != nil && c.Visibility != "public-and-private" {
		cisURL = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_CIS_API_ENDPOINT", c.Region, cisURL)
	}
	cisEndPoint := EnvFallBack([]string{"IBMCLOUD_CIS_API_ENDPOINT"}, cisURL)

	// IBM Network CIS Zones service
	cisZonesV1Opt := &ciszonesv1.ZonesV1Options{
		URL:           cisEndPoint,
		Crn:           core.StringPtr(""),
		Authenticator: authenticator,
	}
	session.cisZonesV1Client, session.cisZonesErr = ciszonesv1.NewZonesV1(cisZonesV1Opt)
	if session.cisZonesErr != nil {
		session.cisZonesErr = fmt.Errorf(
			"Error occured while configuring CIS Zones service: %s",
			session.cisZonesErr)
	}
	if session.cisZonesV1Client != nil && session.cisZonesV1Client.Service != nil {
		session.cisZonesV1Client.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.cisZonesV1Client.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// IBM Network CIS DNS Record service
	cisDNSRecordsOpt := &cisdnsrecordsv1.DnsRecordsV1Options{
		URL:            cisEndPoint,
		Crn:            core.StringPtr(""),
		ZoneIdentifier: core.StringPtr(""),
		Authenticator:  authenticator,
	}
	session.cisDNSRecordsClient, session.cisDNSErr = cisdnsrecordsv1.NewDnsRecordsV1(cisDNSRecordsOpt)
	if session.cisDNSErr != nil {
		session.cisDNSErr = fmt.Errorf("[ERROR] Error occured while configuring CIS DNS Service: %s", session.cisDNSErr)
	}
	if session.cisDNSRecordsClient != nil && session.cisDNSRecordsClient.Service != nil {
		session.cisDNSRecordsClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.cisDNSRecordsClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// IBM Network CIS DNS Record bulk service
	cisDNSRecordBulkOpt := &cisdnsbulkv1.DnsRecordBulkV1Options{
		URL:            cisEndPoint,
		Crn:            core.StringPtr(""),
		ZoneIdentifier: core.StringPtr(""),
		Authenticator:  authenticator,
	}
	session.cisDNSRecordBulkClient, session.cisDNSBulkErr = cisdnsbulkv1.NewDnsRecordBulkV1(cisDNSRecordBulkOpt)
	if session.cisDNSBulkErr != nil {
		session.cisDNSBulkErr = fmt.Errorf(
			"Error occured while configuration CIS DNS bulk service : %s",
			session.cisDNSBulkErr)
	}
	if session.cisDNSRecordBulkClient != nil && session.cisDNSRecordBulkClient.Service != nil {
		session.cisDNSRecordBulkClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.cisDNSRecordBulkClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// IBM Network CIS Global load balancer pool
	cisGLBPoolOpt := &cisglbpoolv0.GlobalLoadBalancerPoolsV0Options{
		URL:           cisEndPoint,
		Crn:           core.StringPtr(""),
		Authenticator: authenticator,
	}
	session.cisGLBPoolClient, session.cisGLBPoolErr =
		cisglbpoolv0.NewGlobalLoadBalancerPoolsV0(cisGLBPoolOpt)
	if session.cisGLBPoolErr != nil {
		session.cisGLBPoolErr =
			fmt.Errorf("[ERROR] Error occured while configuring CIS GLB Pool service: %s",
				session.cisGLBPoolErr)
	}
	if session.cisGLBPoolClient != nil && session.cisGLBPoolClient.Service != nil {
		session.cisGLBPoolClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.cisGLBPoolClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// IBM Network CIS Global load balancer
	cisGLBOpt := &cisglbv1.GlobalLoadBalancerV1Options{
		URL:            cisEndPoint,
		Authenticator:  authenticator,
		Crn:            core.StringPtr(""),
		ZoneIdentifier: core.StringPtr(""),
	}
	session.cisGLBClient, session.cisGLBErr = cisglbv1.NewGlobalLoadBalancerV1(cisGLBOpt)
	if session.cisGLBErr != nil {
		session.cisGLBErr =
			fmt.Errorf("[ERROR] Error occured while configuring CIS GLB service: %s",
				session.cisGLBErr)
	}
	if session.cisGLBClient != nil && session.cisGLBClient.Service != nil {
		session.cisGLBClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.cisGLBClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// IBM Network CIS Global load balancer health check/monitor
	cisGLBHealthCheckOpt := &cisglbhealthcheckv1.GlobalLoadBalancerMonitorV1Options{
		URL:           cisEndPoint,
		Crn:           core.StringPtr(""),
		Authenticator: authenticator,
	}
	session.cisGLBHealthCheckClient, session.cisGLBHealthCheckErr =
		cisglbhealthcheckv1.NewGlobalLoadBalancerMonitorV1(cisGLBHealthCheckOpt)
	if session.cisGLBHealthCheckErr != nil {
		session.cisGLBHealthCheckErr =
			fmt.Errorf("[ERROR] Error occured while configuring CIS GLB Health Check service: %s",
				session.cisGLBHealthCheckErr)
	}
	if session.cisGLBHealthCheckClient != nil && session.cisGLBHealthCheckClient.Service != nil {
		session.cisGLBHealthCheckClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.cisGLBHealthCheckClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// IBM Network CIS IP
	cisIPOpt := &cisipv1.CisIpApiV1Options{
		URL:           cisEndPoint,
		Authenticator: authenticator,
	}
	session.cisIPClient, session.cisIPErr = cisipv1.NewCisIpApiV1(cisIPOpt)
	if session.cisIPErr != nil {
		session.cisIPErr = fmt.Errorf("[ERROR] Error occured while configuring CIS IP service: %s",
			session.cisIPErr)
	}
	if session.cisIPClient != nil && session.cisIPClient.Service != nil {
		session.cisIPClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.cisIPClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// IBM Network CIS Zone Rate Limit
	cisRLOpt := &cisratelimitv1.ZoneRateLimitsV1Options{
		URL:            cisEndPoint,
		Crn:            core.StringPtr(""),
		ZoneIdentifier: core.StringPtr(""),
		Authenticator:  authenticator,
	}
	session.cisRLClient, session.cisRLErr = cisratelimitv1.NewZoneRateLimitsV1(cisRLOpt)
	if session.cisRLErr != nil {
		session.cisRLErr = fmt.Errorf(
			"Error occured while cofiguring CIS Zone Rate Limit service: %s",
			session.cisRLErr)
	}
	if session.cisRLClient != nil && session.cisRLClient.Service != nil {
		session.cisRLClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.cisRLClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}
	// IBM Network CIS Alerts
	cisAlertsOpt := &cisalertsv1.AlertsV1Options{
		URL:           cisEndPoint,
		Crn:           core.StringPtr(""),
		Authenticator: authenticator,
	}
	session.cisAlertsClient, session.cisAlertsErr = cisalertsv1.NewAlertsV1(cisAlertsOpt)
	if session.cisAlertsErr != nil {
		session.cisAlertsErr =
			fmt.Errorf("[ERROR] Error occured while configuring CIS Alerts : %s",
				session.cisAlertsErr)
	}
	if session.cisAlertsClient != nil && session.cisAlertsClient.Service != nil {
		session.cisAlertsClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.cisAlertsClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// IBM Network CIS Page Rules
	cisPageRuleOpt := &cispagerulev1.PageRuleApiV1Options{
		URL:           cisEndPoint,
		Crn:           core.StringPtr(""),
		ZoneID:        core.StringPtr(""),
		Authenticator: authenticator,
	}
	session.cisPageRuleClient, session.cisPageRuleErr = cispagerulev1.NewPageRuleApiV1(cisPageRuleOpt)
	if session.cisPageRuleErr != nil {
		session.cisPageRuleErr = fmt.Errorf(
			"Error occured while cofiguring CIS Page Rule service: %s",
			session.cisPageRuleErr)
	}
	if session.cisPageRuleClient != nil && session.cisPageRuleClient.Service != nil {
		session.cisPageRuleClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.cisPageRuleClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// IBM Network CIS Edge Function
	cisEdgeFunctionOpt := &cisedgefunctionv1.EdgeFunctionsApiV1Options{
		URL:            cisEndPoint,
		Crn:            core.StringPtr(""),
		ZoneIdentifier: core.StringPtr(""),
		Authenticator:  authenticator,
	}
	session.cisEdgeFunctionClient, session.cisEdgeFunctionErr =
		cisedgefunctionv1.NewEdgeFunctionsApiV1(cisEdgeFunctionOpt)
	if session.cisEdgeFunctionErr != nil {
		session.cisEdgeFunctionErr =
			fmt.Errorf("[ERROR] Error occured while configuring CIS Edge Function service: %s",
				session.cisEdgeFunctionErr)
	}
	if session.cisEdgeFunctionClient != nil && session.cisEdgeFunctionClient.Service != nil {
		session.cisEdgeFunctionClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.cisEdgeFunctionClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// IBM Network CIS SSL certificate
	cisSSLOpt := &cissslv1.SslCertificateApiV1Options{
		URL:            cisEndPoint,
		Crn:            core.StringPtr(""),
		ZoneIdentifier: core.StringPtr(""),
		Authenticator:  authenticator,
	}

	session.cisSSLClient, session.cisSSLErr = cissslv1.NewSslCertificateApiV1(cisSSLOpt)
	if session.cisSSLErr != nil {
		session.cisSSLErr =
			fmt.Errorf("[ERROR] Error occured while configuring CIS SSL certificate service: %s",
				session.cisSSLErr)
	}
	if session.cisSSLClient != nil && session.cisSSLClient.Service != nil {
		session.cisSSLClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.cisSSLClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// IBM Network CIS WAF Package
	cisWAFPackageOpt := &ciswafpackagev1.WafRulePackagesApiV1Options{
		URL:           cisEndPoint,
		Crn:           core.StringPtr(""),
		ZoneID:        core.StringPtr(""),
		Authenticator: authenticator,
	}
	session.cisWAFPackageClient, session.cisWAFPackageErr =
		ciswafpackagev1.NewWafRulePackagesApiV1(cisWAFPackageOpt)
	if session.cisWAFPackageErr != nil {
		session.cisWAFPackageErr =
			fmt.Errorf("[ERROR] Error occured while configuration CIS WAF Package service: %s",
				session.cisWAFPackageErr)
	}
	if session.cisWAFPackageClient != nil && session.cisWAFPackageClient.Service != nil {
		session.cisWAFPackageClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.cisWAFPackageClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// IBM Network CIS Domain settings
	cisDomainSettingsOpt := &cisdomainsettingsv1.ZonesSettingsV1Options{
		URL:            cisEndPoint,
		Crn:            core.StringPtr(""),
		ZoneIdentifier: core.StringPtr(""),
		Authenticator:  authenticator,
	}
	session.cisDomainSettingsClient, session.cisDomainSettingsErr =
		cisdomainsettingsv1.NewZonesSettingsV1(cisDomainSettingsOpt)
	if session.cisDomainSettingsErr != nil {
		session.cisDomainSettingsErr =
			fmt.Errorf("[ERROR] Error occured while configuring CIS Domain Settings service: %s",
				session.cisDomainSettingsErr)
	}
	if session.cisDomainSettingsClient != nil && session.cisDomainSettingsClient.Service != nil {
		session.cisDomainSettingsClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.cisDomainSettingsClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// IBM Network CIS Routing
	cisRoutingOpt := &cisroutingv1.RoutingV1Options{
		URL:            cisEndPoint,
		Crn:            core.StringPtr(""),
		ZoneIdentifier: core.StringPtr(""),
		Authenticator:  authenticator,
	}
	session.cisRoutingClient, session.cisRoutingErr =
		cisroutingv1.NewRoutingV1(cisRoutingOpt)
	if session.cisRoutingErr != nil {
		session.cisRoutingErr =
			fmt.Errorf("[ERROR] Error occured while configuring CIS Routing service: %s",
				session.cisRoutingErr)
	}
	if session.cisRoutingClient != nil && session.cisRoutingClient.Service != nil {
		session.cisRoutingClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.cisRoutingClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// IBM Network CIS WAF Group
	cisWAFGroupOpt := &ciswafgroupv1.WafRuleGroupsApiV1Options{
		URL:           cisEndPoint,
		Crn:           core.StringPtr(""),
		ZoneID:        core.StringPtr(""),
		Authenticator: authenticator,
	}
	session.cisWAFGroupClient, session.cisWAFGroupErr =
		ciswafgroupv1.NewWafRuleGroupsApiV1(cisWAFGroupOpt)
	if session.cisWAFGroupErr != nil {
		session.cisWAFGroupErr =
			fmt.Errorf("[ERROR] Error occured while configuring CIS WAF Group service: %s",
				session.cisWAFGroupErr)
	}
	if session.cisWAFGroupClient != nil && session.cisWAFGroupClient.Service != nil {
		session.cisWAFGroupClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.cisWAFGroupClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// IBM Network CIS Cache service
	cisCacheOpt := &ciscachev1.CachingApiV1Options{
		URL:           cisEndPoint,
		Crn:           core.StringPtr(""),
		ZoneID:        core.StringPtr(""),
		Authenticator: authenticator,
	}
	session.cisCacheClient, session.cisCacheErr =
		ciscachev1.NewCachingApiV1(cisCacheOpt)
	if session.cisCacheErr != nil {
		session.cisCacheErr =
			fmt.Errorf("[ERROR] Error occured while configuring CIS Caching service: %s",
				session.cisCacheErr)
	}
	if session.cisCacheClient != nil && session.cisCacheClient.Service != nil {
		session.cisCacheClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.cisCacheClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// IBM Network CIS Custom pages service
	cisCustomPageOpt := &ciscustompagev1.CustomPagesV1Options{
		URL:            cisEndPoint,
		Crn:            core.StringPtr(""),
		ZoneIdentifier: core.StringPtr(""),
		Authenticator:  authenticator,
	}

	session.cisCustomPageClient, session.cisCustomPageErr =
		ciscustompagev1.NewCustomPagesV1(cisCustomPageOpt)
	if session.cisCustomPageErr != nil {
		session.cisCustomPageErr =
			fmt.Errorf("[ERROR] Error occured while configuring CIS Custom Pages service: %s",
				session.cisCustomPageErr)
	}
	if session.cisCustomPageClient != nil && session.cisCustomPageClient.Service != nil {
		session.cisCustomPageClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.cisCustomPageClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// IBM Network CIS Firewall Access rule
	cisAccessRuleOpt := &cisaccessrulev1.ZoneFirewallAccessRulesV1Options{
		URL:            cisEndPoint,
		Crn:            core.StringPtr(""),
		ZoneIdentifier: core.StringPtr(""),
		Authenticator:  authenticator,
	}
	session.cisAccessRuleClient, session.cisAccessRuleErr =
		cisaccessrulev1.NewZoneFirewallAccessRulesV1(cisAccessRuleOpt)
	if session.cisAccessRuleErr != nil {
		session.cisAccessRuleErr =
			fmt.Errorf("[ERROR] Error occured while configuring CIS Firewall Access Rule service: %s",
				session.cisAccessRuleErr)
	}
	if session.cisAccessRuleClient != nil && session.cisAccessRuleClient.Service != nil {
		session.cisAccessRuleClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.cisAccessRuleClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// IBM Network CIS Firewall User Agent Blocking rule
	cisUARuleOpt := &cisuarulev1.UserAgentBlockingRulesV1Options{
		URL:            cisEndPoint,
		Crn:            core.StringPtr(""),
		ZoneIdentifier: core.StringPtr(""),
		Authenticator:  authenticator,
	}
	session.cisUARuleClient, session.cisUARuleErr =
		cisuarulev1.NewUserAgentBlockingRulesV1(cisUARuleOpt)
	if session.cisUARuleErr != nil {
		session.cisUARuleErr =
			fmt.Errorf("[ERROR] Error occured while configuring CIS Firewall User Agent Blocking Rule service: %s",
				session.cisUARuleErr)
	}
	if session.cisUARuleClient != nil && session.cisUARuleClient.Service != nil {
		session.cisUARuleClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.cisUARuleClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// IBM Network CIS Firewall Lockdown rule
	cisLockdownOpt := &cislockdownv1.ZoneLockdownV1Options{
		URL:            cisEndPoint,
		Crn:            core.StringPtr(""),
		ZoneIdentifier: core.StringPtr(""),
		Authenticator:  authenticator,
	}
	session.cisLockdownClient, session.cisLockdownErr =
		cislockdownv1.NewZoneLockdownV1(cisLockdownOpt)
	if session.cisLockdownErr != nil {
		session.cisLockdownErr =
			fmt.Errorf("[ERROR] Error occured while configuring CIS Firewall Lockdown Rule service: %s",
				session.cisLockdownErr)
	}
	if session.cisLockdownClient != nil && session.cisLockdownClient.Service != nil {
		session.cisLockdownClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.cisLockdownClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// IBM Network CIS Range Application rule
	cisRangeAppOpt := &cisrangeappv1.RangeApplicationsV1Options{
		URL:            cisEndPoint,
		Crn:            core.StringPtr(""),
		ZoneIdentifier: core.StringPtr(""),
		Authenticator:  authenticator,
	}
	session.cisRangeAppClient, session.cisRangeAppErr =
		cisrangeappv1.NewRangeApplicationsV1(cisRangeAppOpt)
	if session.cisRangeAppErr != nil {
		session.cisRangeAppErr =
			fmt.Errorf("[ERROR] Error occured while configuring CIS Range Application rule service: %s",
				session.cisRangeAppErr)
	}
	if session.cisRangeAppClient != nil && session.cisRangeAppClient.Service != nil {
		session.cisRangeAppClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.cisRangeAppClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// IBM Network CIS WAF Rule Service
	cisWAFRuleOpt := &ciswafrulev1.WafRulesApiV1Options{
		URL:           cisEndPoint,
		Crn:           core.StringPtr(""),
		ZoneID:        core.StringPtr(""),
		Authenticator: authenticator,
	}
	session.cisWAFRuleClient, session.cisWAFRuleErr =
		ciswafrulev1.NewWafRulesApiV1(cisWAFRuleOpt)
	if session.cisWAFRuleErr != nil {
		session.cisWAFRuleErr = fmt.Errorf(
			"Error occured while configuring CIS WAF Rules service: %s",
			session.cisWAFRuleErr)
	}
	if session.cisWAFRuleClient != nil && session.cisWAFRuleClient.Service != nil {
		session.cisWAFRuleClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.cisWAFRuleClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// IBM Network CIS LogpushJobs
	cisLogpushJobOpt := &cislogpushjobsapiv1.LogpushJobsApiV1Options{
		URL:           cisEndPoint,
		Crn:           core.StringPtr(""),
		ZoneID:        core.StringPtr(""),
		Dataset:       core.StringPtr(""),
		Authenticator: authenticator,
	}
	session.cisLogpushJobsClient, session.cisLogpushJobsErr = cislogpushjobsapiv1.NewLogpushJobsApiV1(cisLogpushJobOpt)
	if session.cisLogpushJobsErr != nil {
		session.cisLogpushJobsErr =
			fmt.Errorf("[ERROR] Error occured while configuring CIS LogpushJobs : %s",
				session.cisLogpushJobsErr)
	}
	if session.cisLogpushJobsClient != nil && session.cisLogpushJobsClient.Service != nil {
		session.cisLogpushJobsClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.cisLogpushJobsClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// IBM MTLS Session
	cisMtlsOpt := &cismtlsv1.MtlsV1Options{
		URL:           cisEndPoint,
		Crn:           core.StringPtr(""),
		Authenticator: authenticator,
	}
	session.cisMtlsClient, session.cisMtlsErr = cismtlsv1.NewMtlsV1(cisMtlsOpt)
	if session.cisMtlsErr != nil {
		session.cisMtlsErr =
			fmt.Errorf("[ERROR] Error occured while configuring CIS MTLS : %s",
				session.cisMtlsErr)
	}
	if session.cisMtlsClient != nil && session.cisMtlsClient.Service != nil {
		session.cisMtlsClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.cisMtlsClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// IBM Network CIS Webhooks
	cisWebhooksOpt := &ciswebhooksv1.WebhooksV1Options{
		URL:           cisEndPoint,
		Crn:           core.StringPtr(""),
		Authenticator: authenticator,
	}
	session.cisWebhooksClient, session.cisWebhooksErr = ciswebhooksv1.NewWebhooksV1(cisWebhooksOpt)
	if session.cisWebhooksErr != nil {
		session.cisWebhooksErr =
			fmt.Errorf("[ERROR] Error occured while configuring CIS Webhooks : %s",
				session.cisWebhooksErr)
	}
	if session.cisWebhooksClient != nil && session.cisWebhooksClient.Service != nil {
		session.cisWebhooksClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.cisWebhooksClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}
	// IBM Network CIS Filters
	cisFiltersOpt := &cisfiltersv1.FiltersV1Options{
		URL:           cisEndPoint,
		Authenticator: authenticator,
	}
	session.cisFiltersClient, session.cisFiltersErr = cisfiltersv1.NewFiltersV1(cisFiltersOpt)
	if session.cisFiltersErr != nil {
		session.cisFiltersErr =
			fmt.Errorf("[ERROR] Error occured while configuring CIS Filters : %s",
				session.cisFiltersErr)
	}
	if session.cisFiltersClient != nil && session.cisFiltersClient.Service != nil {
		session.cisFiltersClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.cisFiltersClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// IBM Network CIS Firewall rules
	cisFirewallrulesOpt := &cisfirewallrulesv1.FirewallRulesV1Options{
		URL:           cisEndPoint,
		Authenticator: authenticator,
	}
	session.cisFirewallRulesClient, session.cisFirewallRulesErr = cisfirewallrulesv1.NewFirewallRulesV1(cisFirewallrulesOpt)
	if session.cisFirewallRulesErr != nil {
		session.cisFirewallRulesErr =
			fmt.Errorf("[ERROR] Error occured while configuring CIS Firewall rules : %s",
				session.cisFirewallRulesErr)
	}
	if session.cisFirewallRulesClient != nil && session.cisFirewallRulesClient.Service != nil {
		session.cisFirewallRulesClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.cisFirewallRulesClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// IBM Network CIS Authenticated Origin Pull
	cisOriginAuthOptions := &cisoriginpull.AuthenticatedOriginPullApiV1Options{
		URL:            cisEndPoint,
		Authenticator:  authenticator,
		Crn:            core.StringPtr(""),
		ZoneIdentifier: core.StringPtr(""),
	}

	session.cisOriginAuthClient, session.cisOriginAuthPullErr =
		cisoriginpull.NewAuthenticatedOriginPullApiV1(cisOriginAuthOptions)
	if session.cisOriginAuthPullErr != nil {
		session.cisOriginAuthPullErr = fmt.Errorf(
			"Error occured while configuring CIS Authenticated Origin Pullservice: %s",
			session.cisOriginAuthPullErr)
	}
	if session.cisOriginAuthClient != nil && session.cisOriginAuthClient.Service != nil {
		session.cisOriginAuthClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.cisOriginAuthClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// IAM IDENTITY Service
	// iamIdenityURL := fmt.Sprintf("https://%s.iam.cloud.ibm.com/v1", c.Region)
	iamIdenityURL := iamidentity.DefaultServiceURL
	if c.Visibility == "private" || c.Visibility == "public-and-private" {
		if c.Region == "us-south" || c.Region == "us-east" {
			iamIdenityURL = ContructEndpoint(fmt.Sprintf("private.%s.iam", c.Region), cloudEndpoint)
		} else {
			iamIdenityURL = ContructEndpoint("private.iam", cloudEndpoint)
		}
	}
	if fileMap != nil && c.Visibility != "public-and-private" {
		iamIdenityURL = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_IAM_API_ENDPOINT", c.Region, iamIdenityURL)
	}
	iamIdentityOptions := &iamidentity.IamIdentityV1Options{
		Authenticator: authenticator,
		URL:           EnvFallBack([]string{"IBMCLOUD_IAM_API_ENDPOINT"}, iamIdenityURL),
	}
	iamIdentityClient, err := iamidentity.NewIamIdentityV1(iamIdentityOptions)
	if err != nil {
		session.iamIdentityErr = fmt.Errorf("[ERROR] Error occured while configuring IAM Identity service: %q", err)
	}
	if iamIdentityClient != nil && iamIdentityClient.Service != nil {
		iamIdentityClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		iamIdentityClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}
	session.iamIdentityAPI = iamIdentityClient

	// IAM POLICY MANAGEMENT Service
	iamPolicyManagementURL := iampolicymanagement.DefaultServiceURL
	if c.Visibility == "private" || c.Visibility == "public-and-private" {
		if c.Region == "us-south" || c.Region == "us-east" {
			iamPolicyManagementURL = ContructEndpoint(fmt.Sprintf("private.%s.iam", c.Region), cloudEndpoint)
		} else {
			iamPolicyManagementURL = ContructEndpoint("private.iam", cloudEndpoint)
		}
	}
	if fileMap != nil && c.Visibility != "public-and-private" {
		iamPolicyManagementURL = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_IAM_API_ENDPOINT", c.Region, iamPolicyManagementURL)
	}
	iamPolicyManagementOptions := &iampolicymanagement.IamPolicyManagementV1Options{
		Authenticator: authenticator,
		URL:           EnvFallBack([]string{"IBMCLOUD_IAM_API_ENDPOINT"}, iamPolicyManagementURL),
	}
	iamPolicyManagementClient, err := iampolicymanagement.NewIamPolicyManagementV1(iamPolicyManagementOptions)
	if err != nil {
		session.iamPolicyManagementErr = fmt.Errorf("[ERROR] Error occured while configuring IAM Policy Management service: %q", err)
	}
	if iamPolicyManagementClient != nil && iamPolicyManagementClient.Service != nil {
		iamPolicyManagementClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		iamPolicyManagementClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}
	session.iamPolicyManagementAPI = iamPolicyManagementClient

	// IAM ACCESS GROUP
	iamAccessGroupsURL := iamaccessgroups.DefaultServiceURL
	if c.Visibility == "private" || c.Visibility == "public-and-private" {
		if c.Region == "us-south" || c.Region == "us-east" {
			iamAccessGroupsURL = ContructEndpoint(fmt.Sprintf("private.%s.iam", c.Region), cloudEndpoint)
		} else {
			iamAccessGroupsURL = ContructEndpoint("private.iam", cloudEndpoint)
		}
	}
	if fileMap != nil && c.Visibility != "public-and-private" {
		iamAccessGroupsURL = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_IAM_API_ENDPOINT", c.Region, iamAccessGroupsURL)
	}
	iamAccessGroupsOptions := &iamaccessgroups.IamAccessGroupsV2Options{
		Authenticator: authenticator,
		URL:           EnvFallBack([]string{"IBMCLOUD_IAM_API_ENDPOINT"}, iamAccessGroupsURL),
	}
	iamAccessGroupsClient, err := iamaccessgroups.NewIamAccessGroupsV2(iamAccessGroupsOptions)
	if err != nil {
		session.iamAccessGroupsErr = fmt.Errorf("[ERROR] Error occured while configuring IAM Access Group service: %q", err)
	}
	if iamAccessGroupsClient != nil && iamAccessGroupsClient.Service != nil {
		iamAccessGroupsClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		iamAccessGroupsClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}
	session.iamAccessGroupsAPI = iamAccessGroupsClient

	// RESOURCE MANAGEMENT Service
	rmURL := resourcemanager.DefaultServiceURL
	if c.Visibility == "private" {
		if c.Region == "us-south" || c.Region == "us-east" {
			rmURL = ContructEndpoint(fmt.Sprintf("private.%s.resource-controller", c.Region), fmt.Sprintf("%s", cloudEndpoint))
		} else {
			fmt.Println("Private Endpint supports only us-south and us-east region specific endpoint")
			rmURL = ContructEndpoint("private.us-south.resource-controller", fmt.Sprintf("%s", cloudEndpoint))
		}
	}
	if c.Visibility == "public-and-private" {
		if c.Region == "us-south" || c.Region == "us-east" {
			rmURL = ContructEndpoint(fmt.Sprintf("private.%s.resource-controller", c.Region), fmt.Sprintf("%s", cloudEndpoint))
		} else {
			rmURL = resourcemanager.DefaultServiceURL
		}
	}
	if fileMap != nil && c.Visibility != "public-and-private" {
		rmURL = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_RESOURCE_MANAGEMENT_API_ENDPOINT", c.Region, rmURL)
	}
	resourceManagerOptions := &resourcemanager.ResourceManagerV2Options{
		Authenticator: authenticator,
		URL:           EnvFallBack([]string{"IBMCLOUD_RESOURCE_MANAGEMENT_API_ENDPOINT"}, rmURL),
	}
	resourceManagerClient, err := resourcemanager.NewResourceManagerV2(resourceManagerOptions)
	if err != nil {
		session.resourceManagerErr = fmt.Errorf("[ERROR] Error occured while configuring Resource Manager service: %q", err)
	}
	if resourceManagerClient != nil && resourceManagerClient.Service != nil {
		resourceManagerClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		resourceManagerClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}
	session.resourceManagerAPI = resourceManagerClient

	//CLOUD SHELL Service
	cloudShellUrl := ibmcloudshellv1.DefaultServiceURL
	if fileMap != nil && c.Visibility != "public-and-private" {
		cloudShellUrl = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_CLOUD_SHELL_API_ENDPOINT", c.Region, cloudShellUrl)
	}
	ibmCloudShellClientOptions := &ibmcloudshellv1.IBMCloudShellV1Options{
		Authenticator: authenticator,
		URL:           EnvFallBack([]string{"IBMCLOUD_CLOUD_SHELL_API_ENDPOINT"}, cloudShellUrl),
	}
	session.ibmCloudShellClient, err = ibmcloudshellv1.NewIBMCloudShellV1(ibmCloudShellClientOptions)
	if err != nil {
		session.ibmCloudShellClientErr = fmt.Errorf("[ERROR] Error occurred while configuring IBM Cloud Shell service: %q", err)
	}
	if session.ibmCloudShellClient != nil && session.ibmCloudShellClient.Service != nil {
		session.ibmCloudShellClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		session.ibmCloudShellClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}

	// ENTERPRISE Service
	enterpriseURL := enterprisemanagementv1.DefaultServiceURL
	if c.Visibility == "private" {
		if c.Region == "us-south" || c.Region == "us-east" || c.Region == "eu-fr" {
			enterpriseURL = ContructEndpoint(fmt.Sprintf("private.%s.enterprise", c.Region), fmt.Sprintf("%s/v1", cloudEndpoint))
		} else {
			fmt.Println("Private Endpint supports only us-south and us-east region specific endpoint")
			enterpriseURL = ContructEndpoint("private.us-south.enterprise", fmt.Sprintf("%s/v1", cloudEndpoint))
		}
	}
	if c.Visibility == "public-and-private" {
		if c.Region == "us-south" || c.Region == "us-east" || c.Region == "eu-fr" {
			enterpriseURL = ContructEndpoint(fmt.Sprintf("private.%s.enterprise", c.Region),
				fmt.Sprintf("%s/v1", cloudEndpoint))
		} else {
			enterpriseURL = enterprisemanagementv1.DefaultServiceURL

		}
	}

	resourceManagementAPIv2, err := managementv2.New(sess.BluemixSession)
	if err != nil {
		session.resourceManagementConfigErrv2 = fmt.Errorf("Error occured while configuring Resource Management service: %q", err)
	}
	session.resourceManagementServiceAPIv2 = resourceManagementAPIv2

	globalTaggingAPI, err := globaltaggingv3.New(sess.BluemixSession)
	if err != nil {
		session.globalTaggingConfigErr = fmt.Errorf("Error occured while configuring Global Tagging: %q", err)
	}
	session.globalTaggingServiceAPI = globalTaggingAPI

	var authenticator core.Authenticator
	if c.BluemixAPIKey != "" {
		authenticator = &core.IamAuthenticator{
			ApiKey: c.BluemixAPIKey,
			URL:    EnvFallBack([]string{"IBMCLOUD_IAM_API_ENDPOINT"}, "https://iam.cloud.ibm.com") + "/identity/token",
		}
	} else if strings.HasPrefix(sess.BluemixSession.Config.IAMAccessToken, "Bearer") {
		authenticator = &core.BearerTokenAuthenticator{
			BearerToken: sess.BluemixSession.Config.IAMAccessToken[7:],
		}
	} else {
		authenticator = &core.BearerTokenAuthenticator{
			BearerToken: sess.BluemixSession.Config.IAMAccessToken,
		}
	}

	// Construct an "options" struct for creating the service client.
	
	// RESOURCE CONTROLLER Service
	var cloudEndpoint = "cloud.ibm.com"
	rcURL := resourcecontroller.DefaultServiceURL
	if c.Visibility == "private" {
		if c.Region == "us-south" || c.Region == "us-east" {
			rcURL = ContructEndpoint(fmt.Sprintf("private.%s.resource-controller", c.Region), cloudEndpoint)
		} else {
			fmt.Println("Private Endpint supports only us-south and us-east region specific endpoint")
			rcURL = ContructEndpoint("private.us-south.resource-controller", cloudEndpoint)
		}
	}
	if c.Visibility == "public-and-private" {
		if c.Region == "us-south" || c.Region == "us-east" {
			rcURL = ContructEndpoint(fmt.Sprintf("private.%s.resource-controller", c.Region), cloudEndpoint)
		} else {
			rcURL = resourcecontroller.DefaultServiceURL
		}
	}
	var fileMap map[string]interface{}
	if fileMap != nil && c.Visibility != "public-and-private" {
		rcURL = fileFallBack(fileMap, c.Visibility, "IBMCLOUD_RESOURCE_CONTROLLER_API_ENDPOINT", c.Region, rcURL)
	}
	resourceControllerOptions := &resourcecontroller.ResourceControllerV2Options{
		Authenticator: authenticator,
		URL:           EnvFallBack([]string{"IBMCLOUD_RESOURCE_CONTROLLER_API_ENDPOINT"}, rcURL),
	}
	resourceControllerClient, err := resourcecontroller.NewResourceControllerV2(resourceControllerOptions)
	if err != nil {
		session.resourceControllerErr = fmt.Errorf("[ERROR] Error occured while configuring Resource Controller service: %q", err)
	}
	if resourceControllerClient != nil && resourceControllerClient.Service != nil {
		resourceControllerClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		resourceControllerClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
		})
	}
	session.resourceControllerAPI = resourceControllerClient

	// Construct the service options.
	projectClientOptions := &projectv1.ProjectV1Options{
		Authenticator: authenticator,
	}

	// Construct the service client.
	session.projectClient, err = projectv1.NewProjectV1(projectClientOptions)
	if err == nil {
		// Enable retries for API calls
		session.projectClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
		// Add custom header for analytics
		session.projectClient.SetDefaultHeaders(gohttp.Header{
			"X-Original-User-Agent": { fmt.Sprintf("terraform-provider-ibm/%s", version.Version) },
		})
	} else {
		session.projectClientErr = fmt.Errorf("Error occurred while configuring Projects API Specification service: %q", err)
	}

	userConfig, err := fetchUserDetails(authenticator)
	if err != nil {
		session.bmxUserFetchErr = fmt.Errorf("Error occured while fetching account user details: %q", err)
	}
	session.bmxUserDetails = userConfig

	if os.Getenv("TF_LOG") != "" {
		logDestination := log.Writer()
		goLogger := log.New(logDestination, "", log.LstdFlags)
		core.SetLogger(core.NewLogger(core.LevelDebug, goLogger, goLogger))
	}

	// Dummy stmt to workaround "imported and not used" build error
	_ = version.Version

	return session, nil
}

func newSession(c *Config) (*Session, error) {
	ibmSession := &Session{}

	if c.BluemixAPIKey != "" {
		log.Println("Configuring IBM Cloud Session with API key")
		var sess *bxsession.Session
		bmxConfig := &bluemix.Config{
			BluemixAPIKey: c.BluemixAPIKey,
			HTTPTimeout:   c.BluemixTimeout,
			Region:        c.Region,
			ResourceGroup: c.ResourceGroup,
			RetryDelay:    &c.RetryDelay,
			MaxRetries:    &c.RetryCount,
			Visibility:    c.Visibility,
		}
		sess, err := bxsession.New(bmxConfig)
		if err != nil {
			return nil, err
		}
		ibmSession.BluemixSession = sess
	}

	return ibmSession, nil
}

func authenticateAPIKey(sess *bxsession.Session) error {
	config := sess.Config
	tokenRefresher, err := authentication.NewIAMAuthRepository(config, &rest.Client{
		DefaultHeader: gohttp.Header{
			"User-Agent": []string{http.UserAgent()},
		},
	})
	if err != nil {
		return err
	}
	return tokenRefresher.AuthenticateAPIKey(config.BluemixAPIKey)
}

func fetchUserDetails(authenticator core.Authenticator) (user *UserConfig, err error) {

	request := &gohttp.Request{
		Header: make(gohttp.Header),
	}
	err = authenticator.Authenticate(request)
	if err != nil {
		return nil, err
	}
	auth := request.Header["Authorization"][0]
	iamToken := auth[7:]

	token, err := jwt.Parse(iamToken, func(token *jwt.Token) (interface{}, error) {
		return "", nil
	})
	//TODO validate with key
	if err != nil && !strings.Contains(err.Error(), "key is of invalid type") {
		return nil, err
	}
	err = nil

	user = &UserConfig{}

	claims := token.Claims.(jwt.MapClaims)
	if email, ok := claims["email"]; ok {
		user.userEmail = email.(string)
	}

	user.userID = claims["id"].(string)
	user.UserAccount = claims["account"].(map[string]interface{})["bss"].(string)
	iss := claims["iss"].(string)
	if strings.Contains(iss, "https://iam.cloud.ibm.com") {
		user.cloudName = "bluemix"
	} else {
		user.cloudName = "staging"
	}
	user.cloudType = "public"

	return
}

//nolint deadcode Might be used by generated code.
func EnvFallBack(envs []string, defaultValue string) string {
	for _, k := range envs {
		if v := os.Getenv(k); v != "" {
			if strings.Contains(v, "https://") {
				return v
			} else {
				return fmt.Sprintf("https://%s/v1", v)
			}
		}
	}
	return defaultValue
}

func fileFallBack(fileMap map[string]interface{}, visibility, key, region, defaultValue string) string {
	if val, ok := fileMap[key]; ok {
		if v, ok := val.(map[string]interface{})[visibility]; ok {
			if r, ok := v.(map[string]interface{})[region]; ok && r.(string) != "" {
				return r.(string)
			}
		}
	}
	return defaultValue
}

func ContructEndpoint(subdomain, domain string) string {
	endpoint := fmt.Sprintf("https://%s.%s", subdomain, domain)
	return endpoint
}
