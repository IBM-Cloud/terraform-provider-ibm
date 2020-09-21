package ibm

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	gohttp "net/http"
	"os"
	"strings"
	"time"

	// Added code for the Power Colo Offering

	apigateway "github.com/IBM/apigateway-go-sdk"
	dns "github.com/IBM/dns-svcs-go-sdk/dnssvcsv1"
	"github.com/IBM/go-sdk-core/v3/core"
	cosconfig "github.com/IBM/ibm-cos-sdk-go-config/resourceconfigurationv1"
	kp "github.com/IBM/keyprotect-go-client"
	dl "github.com/IBM/networking-go-sdk/directlinkv1"
	cisdnsrecordsv1 "github.com/IBM/networking-go-sdk/dnsrecordsv1"
	cisedgefunctionv1 "github.com/IBM/networking-go-sdk/edgefunctionsapiv1"
	cisglbhealthcheckv1 "github.com/IBM/networking-go-sdk/globalloadbalancermonitorv1"
	tg "github.com/IBM/networking-go-sdk/transitgatewayapisv1"
	cisratelimitv1 "github.com/IBM/networking-go-sdk/zoneratelimitsv1"
	ciszonesv1 "github.com/IBM/networking-go-sdk/zonesv1"
	vpcclassic "github.com/IBM/vpc-go-sdk/vpcclassicv1"
	vpc "github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/apache/openwhisk-client-go/whisk"
	jwt "github.com/dgrijalva/jwt-go"
	slsession "github.com/softlayer/softlayer-go/session"
	ns "github.ibm.com/ibmcloud/namespace-go-sdk/ibmcloudfunctionsnamespaceapiv1"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/api/account/accountv1"
	"github.com/IBM-Cloud/bluemix-go/api/account/accountv2"
	"github.com/IBM-Cloud/bluemix-go/api/certificatemanager"
	"github.com/IBM-Cloud/bluemix-go/api/cis/cisv1"
	"github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
	"github.com/IBM-Cloud/bluemix-go/api/container/containerv2"
	"github.com/IBM-Cloud/bluemix-go/api/globalsearch/globalsearchv2"
	"github.com/IBM-Cloud/bluemix-go/api/globaltagging/globaltaggingv3"
	"github.com/IBM-Cloud/bluemix-go/api/hpcs"
	"github.com/IBM-Cloud/bluemix-go/api/iam/iamv1"
	"github.com/IBM-Cloud/bluemix-go/api/iampap/iampapv1"
	"github.com/IBM-Cloud/bluemix-go/api/iampap/iampapv2"
	"github.com/IBM-Cloud/bluemix-go/api/iamuum/iamuumv1"
	"github.com/IBM-Cloud/bluemix-go/api/iamuum/iamuumv2"
	"github.com/IBM-Cloud/bluemix-go/api/icd/icdv4"
	"github.com/IBM-Cloud/bluemix-go/api/mccp/mccpv2"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/catalog"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/controller"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/management"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev2/controllerv2"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev2/managementv2"
	"github.com/IBM-Cloud/bluemix-go/api/schematics"
	"github.com/IBM-Cloud/bluemix-go/api/usermanagement/usermanagementv2"
	"github.com/IBM-Cloud/bluemix-go/authentication"
	"github.com/IBM-Cloud/bluemix-go/http"
	"github.com/IBM-Cloud/bluemix-go/rest"
	bxsession "github.com/IBM-Cloud/bluemix-go/session"
	ibmpisession "github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/terraform-provider-ibm/version"
)

//RetryDelay
const RetryAPIDelay = 5 * time.Second

//BluemixRegion ...
var BluemixRegion string

var (
	errEmptySoftLayerCredentials = errors.New("iaas_classic_username and iaas_classic_api_key must be provided. Please see the documentation on how to configure them")
	errEmptyBluemixCredentials   = errors.New("ibmcloud_api_key or bluemix_api_key or iam_token and iam_refresh_token must be provided. Please see the documentation on how to configure it")
)

//UserConfig ...
type UserConfig struct {
	userID      string
	userEmail   string
	userAccount string
	cloudName   string `default:"bluemix"`
	cloudType   string `default:"public"`
	generation  int    `default:"2"`
}

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

	//Softlayer end point url
	SoftLayerEndpointURL string

	//Softlayer API timeout
	SoftLayerTimeout time.Duration

	// Softlayer User Name
	SoftLayerUserName string

	// Softlayer API Key
	SoftLayerAPIKey string

	//Retry Count for API calls
	//Unexposed in the schema at this point as they are used only during session creation for a few calls
	//When sdk implements it we an expose them for expected behaviour
	//https://github.com/softlayer/softlayer-go/issues/41
	RetryCount int
	//Constant Retry Delay for API calls
	RetryDelay time.Duration

	// FunctionNameSpace ...
	FunctionNameSpace string

	//Riaas End point
	RiaasEndPoint string

	//Generation
	Generation int

	//IAM Token
	IAMToken string

	//IAM Refresh Token
	IAMRefreshToken string

	// PowerService Instance
	PowerServiceInstance string

	// Zone
	Zone string
}

//Session stores the information required for communication with the SoftLayer and Bluemix API
type Session struct {
	// SoftLayerSesssion is the the SoftLayer session used to connect to the SoftLayer API
	SoftLayerSession *slsession.Session

	// BluemixSession is the the Bluemix session used to connect to the Bluemix API
	BluemixSession *bxsession.Session
}

// ClientSession ...
type ClientSession interface {
	BluemixSession() (*bxsession.Session, error)
	BluemixAcccountAPI() (accountv2.AccountServiceAPI, error)
	BluemixAcccountv1API() (accountv1.AccountServiceAPI, error)
	BluemixUserDetails() (*UserConfig, error)
	ContainerAPI() (containerv1.ContainerServiceAPI, error)
	VpcContainerAPI() (containerv2.ContainerServiceAPI, error)
	CisAPI() (cisv1.CisServiceAPI, error)
	FunctionClient() (*whisk.Client, error)
	GlobalSearchAPI() (globalsearchv2.GlobalSearchServiceAPI, error)
	GlobalTaggingAPI() (globaltaggingv3.GlobalTaggingServiceAPI, error)
	ICDAPI() (icdv4.ICDServiceAPI, error)
	IAMAPI() (iamv1.IAMServiceAPI, error)
	IAMPAPAPI() (iampapv1.IAMPAPAPI, error)
	IAMPAPAPIV2() (iampapv2.IAMPAPAPIV2, error)
	IAMUUMAPI() (iamuumv1.IAMUUMServiceAPI, error)
	IAMUUMAPIV2() (iamuumv2.IAMUUMServiceAPIv2, error)
	MccpAPI() (mccpv2.MccpServiceAPI, error)
	ResourceCatalogAPI() (catalog.ResourceCatalogAPI, error)
	ResourceManagementAPI() (management.ResourceManagementAPI, error)
	ResourceManagementAPIv2() (managementv2.ResourceManagementAPIv2, error)
	ResourceControllerAPI() (controller.ResourceControllerAPI, error)
	ResourceControllerAPIV2() (controllerv2.ResourceControllerAPIV2, error)
	SoftLayerSession() *slsession.Session
	IBMPISession() (*ibmpisession.IBMPISession, error)
	SchematicsAPI() (schematics.SchematicsServiceAPI, error)
	UserManagementAPI() (usermanagementv2.UserManagementAPI, error)
	CertificateManagerAPI() (certificatemanager.CertificateManagerServiceAPI, error)
	keyProtectAPI() (*kp.Client, error)
	keyManagementAPI() (*kp.Client, error)
	VpcClassicV1API() (*vpcclassic.VpcClassicV1, error)
	VpcV1API() (*vpc.VpcV1, error)
	APIGateway() (*apigateway.ApiGatewayControllerApiV1, error)
	PrivateDnsClientSession() (*dns.DnsSvcsV1, error)
	CosConfigV1API() (*cosconfig.ResourceConfigurationV1, error)
	DirectlinkV1API() (*dl.DirectLinkV1, error)
	TransitGatewayV1API() (*tg.TransitGatewayApisV1, error)
	HpcsEndpointAPI() (hpcs.HPCSV2, error)
	IAMNamespaceAPI() (*ns.IbmCloudFunctionsNamespaceAPIV1, error)
	CisZonesV1ClientSession() (*ciszonesv1.ZonesV1, error)
	CisDNSRecordClientSession() (*cisdnsrecordsv1.DnsRecordsV1, error)
	CisGLBHealthCheckClientSession() (*cisglbhealthcheckv1.GlobalLoadBalancerMonitorV1, error)
	CisRLClientSession() (*cisratelimitv1.ZoneRateLimitsV1, error)
	CisEdgeFunctionClientSession() (*cisedgefunctionv1.EdgeFunctionsApiV1, error)
}

type clientSession struct {
	session *Session

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

	stxConfigErr  error
	stxServiceAPI schematics.SchematicsServiceAPI

	certManagementErr error
	certManagementAPI certificatemanager.CertificateManagerServiceAPI

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

	iamPAPConfigErr  error
	iamPAPServiceAPI iampapv1.IAMPAPAPI

	iamPAPConfigErrv2  error
	iamPAPServiceAPIv2 iampapv2.IAMPAPAPIV2

	iamUUMConfigErr  error
	iamUUMServiceAPI iamuumv1.IAMUUMServiceAPI

	iamUUMConfigErrV2  error
	iamUUMServiceAPIV2 iamuumv2.IAMUUMServiceAPIv2

	iamConfigErr  error
	iamServiceAPI iamv1.IAMServiceAPI

	userManagementErr error
	userManagementAPI usermanagementv2.UserManagementAPI

	icdConfigErr  error
	icdServiceAPI icdv4.ICDServiceAPI

	resourceControllerConfigErr  error
	resourceControllerServiceAPI controller.ResourceControllerAPI

	resourceControllerConfigErrv2  error
	resourceControllerServiceAPIv2 controllerv2.ResourceControllerAPIV2

	resourceManagementConfigErr  error
	resourceManagementServiceAPI management.ResourceManagementAPI

	resourceManagementConfigErrv2  error
	resourceManagementServiceAPIv2 managementv2.ResourceManagementAPIv2

	resourceCatalogConfigErr  error
	resourceCatalogServiceAPI catalog.ResourceCatalogAPI

	powerConfigErr error
	ibmpiConfigErr error
	ibmpiSession   *ibmpisession.IBMPISession

	kpErr error
	kpAPI *kp.API

	kmsErr error
	kmsAPI *kp.API

	hpcsEndpointErr error
	hpcsEndpointAPI hpcs.HPCSV2

	pDnsClient *dns.DnsSvcsV1
	pDnsErr    error

	bluemixSessionErr error

	vpcClassicErr error
	vpcClassicAPI *vpcclassic.VpcClassicV1

	vpcErr error
	vpcAPI *vpc.VpcV1

	directlinkAPI *dl.DirectLinkV1
	directlinkErr error

	cosConfigErr error
	cosConfigAPI *cosconfig.ResourceConfigurationV1

	transitgatewayAPI *tg.TransitGatewayApisV1
	transitgatewayErr error

	iamNamespaceAPI *ns.IbmCloudFunctionsNamespaceAPIV1
	iamNamespaceErr error

	// CIS Zones
	cisZonesErr      error
	cisZonesV1Client *ciszonesv1.ZonesV1

	// CIS dns service options
	cisDNSErr           error
	cisDNSRecordsClient *cisdnsrecordsv1.DnsRecordsV1

	// CIS GLB health check service options
	cisGLBHealthCheckErr    error
	cisGLBHealthCheckClient *cisglbhealthcheckv1.GlobalLoadBalancerMonitorV1

	// CIS Zone Rate Limits service options
	cisRLErr    error
	cisRLClient *cisratelimitv1.ZoneRateLimitsV1

	// CIS Edge Functions service options
	cisEdgeFunctionErr    error
	cisEdgeFunctionClient *cisedgefunctionv1.EdgeFunctionsApiV1
}

// BluemixAcccountAPI ...
func (sess clientSession) BluemixAcccountAPI() (accountv2.AccountServiceAPI, error) {
	return sess.bmxAccountServiceAPI, sess.accountConfigErr
}

// BluemixAcccountAPI ...
func (sess clientSession) BluemixAcccountv1API() (accountv1.AccountServiceAPI, error) {
	return sess.bmxAccountv1ServiceAPI, sess.accountV1ConfigErr
}

// BluemixSession to provide the Bluemix Session
func (sess clientSession) BluemixSession() (*bxsession.Session, error) {
	return sess.session.BluemixSession, sess.bluemixSessionErr
}

// BluemixUserDetails ...
func (sess clientSession) BluemixUserDetails() (*UserConfig, error) {
	return sess.bmxUserDetails, sess.bmxUserFetchErr
}

// ContainerAPI provides Container Service APIs ...
func (sess clientSession) ContainerAPI() (containerv1.ContainerServiceAPI, error) {
	return sess.csServiceAPI, sess.csConfigErr
}

// VpcContainerAPI provides v2Container Service APIs ...
func (sess clientSession) VpcContainerAPI() (containerv2.ContainerServiceAPI, error) {
	return sess.csv2ServiceAPI, sess.csv2ConfigErr
}

// SchematicsAPI provides schematics Service APIs ...
func (sess clientSession) SchematicsAPI() (schematics.SchematicsServiceAPI, error) {
	return sess.stxServiceAPI, sess.stxConfigErr
}

// CisAPI provides Cloud Internet Services APIs ...
func (sess clientSession) CisAPI() (cisv1.CisServiceAPI, error) {
	return sess.cisServiceAPI, sess.cisConfigErr
}

// FunctionClient ...
func (sess clientSession) FunctionClient() (*whisk.Client, error) {
	return sess.functionClient, sess.functionConfigErr
}

// GlobalSearchAPI provides Global Search  APIs ...
func (sess clientSession) GlobalSearchAPI() (globalsearchv2.GlobalSearchServiceAPI, error) {
	return sess.globalSearchServiceAPI, sess.globalSearchConfigErr
}

// GlobalTaggingAPI provides Global Search  APIs ...
func (sess clientSession) GlobalTaggingAPI() (globaltaggingv3.GlobalTaggingServiceAPI, error) {
	return sess.globalTaggingServiceAPI, sess.globalTaggingConfigErr
}

// HpcsEndpointAPI provides Hpcs Endpoint generator APIs ...
func (sess clientSession) HpcsEndpointAPI() (hpcs.HPCSV2, error) {
	return sess.hpcsEndpointAPI, sess.hpcsEndpointErr
}

// IAMAPI provides IAM PAP APIs ...
func (sess clientSession) IAMAPI() (iamv1.IAMServiceAPI, error) {
	return sess.iamServiceAPI, sess.iamConfigErr
}

// UserManagementAPI provides User management APIs ...
func (sess clientSession) UserManagementAPI() (usermanagementv2.UserManagementAPI, error) {
	return sess.userManagementAPI, sess.userManagementErr
}

// IAMPAPAPI provides IAM PAP APIs ...
func (sess clientSession) IAMPAPAPI() (iampapv1.IAMPAPAPI, error) {
	return sess.iamPAPServiceAPI, sess.iamPAPConfigErr
}

// IAMPAPAPIV2 provides IAM PAP APIs ...
func (sess clientSession) IAMPAPAPIV2() (iampapv2.IAMPAPAPIV2, error) {
	return sess.iamPAPServiceAPIv2, sess.iamPAPConfigErrv2
}

// IAMUUMAPI provides IAM UUM APIs ...
func (sess clientSession) IAMUUMAPI() (iamuumv1.IAMUUMServiceAPI, error) {
	return sess.iamUUMServiceAPI, sess.iamUUMConfigErr
}

// IAMUUMAPIV2 provides IAM UUM APIs ...
func (sess clientSession) IAMUUMAPIV2() (iamuumv2.IAMUUMServiceAPIv2, error) {
	return sess.iamUUMServiceAPIV2, sess.iamUUMConfigErrV2
}

// IcdAPI provides IBM Cloud Databases APIs ...
func (sess clientSession) ICDAPI() (icdv4.ICDServiceAPI, error) {
	return sess.icdServiceAPI, sess.icdConfigErr
}

// MccpAPI provides Multi Cloud Controller Proxy APIs ...
func (sess clientSession) MccpAPI() (mccpv2.MccpServiceAPI, error) {
	return sess.cfServiceAPI, sess.cfConfigErr
}

// ResourceCatalogAPI ...
func (sess clientSession) ResourceCatalogAPI() (catalog.ResourceCatalogAPI, error) {
	return sess.resourceCatalogServiceAPI, sess.resourceCatalogConfigErr
}

// ResourceManagementAPI ...
func (sess clientSession) ResourceManagementAPI() (management.ResourceManagementAPI, error) {
	return sess.resourceManagementServiceAPI, sess.resourceManagementConfigErr
}

// ResourceManagementAPIv2 ...
func (sess clientSession) ResourceManagementAPIv2() (managementv2.ResourceManagementAPIv2, error) {
	return sess.resourceManagementServiceAPIv2, sess.resourceManagementConfigErrv2
}

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

//apigatewayAPI provides API Gateway APIs
func (sess clientSession) APIGateway() (*apigateway.ApiGatewayControllerApiV1, error) {
	return sess.apigatewayAPI, sess.apigatewayErr
}

func (sess clientSession) keyProtectAPI() (*kp.Client, error) {
	return sess.kpAPI, sess.kpErr
}

func (sess clientSession) keyManagementAPI() (*kp.Client, error) {
	return sess.kmsAPI, sess.kmsErr
}

func (sess clientSession) VpcClassicV1API() (*vpcclassic.VpcClassicV1, error) {
	return sess.vpcClassicAPI, sess.vpcClassicErr
}

func (sess clientSession) VpcV1API() (*vpc.VpcV1, error) {
	return sess.vpcAPI, sess.vpcErr
}

func (sess clientSession) DirectlinkV1API() (*dl.DirectLinkV1, error) {
	return sess.directlinkAPI, sess.directlinkErr
}

func (sess clientSession) CosConfigV1API() (*cosconfig.ResourceConfigurationV1, error) {
	return sess.cosConfigAPI, sess.cosConfigErr
}

func (sess clientSession) TransitGatewayV1API() (*tg.TransitGatewayApisV1, error) {
	return sess.transitgatewayAPI, sess.transitgatewayErr
}

// Session to the Power Colo Service

func (sess clientSession) IBMPISession() (*ibmpisession.IBMPISession, error) {
	return sess.ibmpiSession, sess.powerConfigErr
}

// Private DNS Service

func (sess clientSession) PrivateDnsClientSession() (*dns.DnsSvcsV1, error) {
	return sess.pDnsClient, sess.pDnsErr
}

// Session to the Namespace cloud function

func (sess clientSession) IAMNamespaceAPI() (*ns.IbmCloudFunctionsNamespaceAPIV1, error) {
	return sess.iamNamespaceAPI, sess.iamNamespaceErr
}

// CIS Zones Service
func (sess clientSession) CisZonesV1ClientSession() (*ciszonesv1.ZonesV1, error) {
	return sess.cisZonesV1Client, sess.cisZonesErr
}

// CIS DNS Service
func (sess clientSession) CisDNSRecordClientSession() (*cisdnsrecordsv1.DnsRecordsV1, error) {
	return sess.cisDNSRecordsClient, sess.cisDNSErr
}

// CIS GLB Health Check/Monitor
func (sess clientSession) CisGLBHealthCheckClientSession() (*cisglbhealthcheckv1.GlobalLoadBalancerMonitorV1, error) {
	return sess.cisGLBHealthCheckClient, sess.cisGLBHealthCheckErr
}

// CIS Zone Rate Limits
func (sess clientSession) CisRLClientSession() (*cisratelimitv1.ZoneRateLimitsV1, error) {
	return sess.cisRLClient, sess.cisRLErr
}

// cCIS Edge Function
func (sess clientSession) CisEdgeFunctionClientSession() (*cisedgefunctionv1.EdgeFunctionsApiV1, error) {
	return sess.cisEdgeFunctionClient, sess.cisEdgeFunctionErr
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
		session.kpErr = errEmptyBluemixCredentials
		session.kmsErr = errEmptyBluemixCredentials
		session.stxConfigErr = errEmptyBluemixCredentials
		session.cfConfigErr = errEmptyBluemixCredentials
		session.cisConfigErr = errEmptyBluemixCredentials
		session.functionConfigErr = errEmptyBluemixCredentials
		session.globalSearchConfigErr = errEmptyBluemixCredentials
		session.globalTaggingConfigErr = errEmptyBluemixCredentials
		session.hpcsEndpointErr = errEmptyBluemixCredentials
		session.iamConfigErr = errEmptyBluemixCredentials
		session.iamPAPConfigErr = errEmptyBluemixCredentials
		session.iamPAPConfigErrv2 = errEmptyBluemixCredentials
		session.iamUUMConfigErr = errEmptyBluemixCredentials
		session.iamUUMConfigErrV2 = errEmptyBluemixCredentials
		session.icdConfigErr = errEmptyBluemixCredentials
		session.resourceCatalogConfigErr = errEmptyBluemixCredentials
		session.resourceManagementConfigErr = errEmptyBluemixCredentials
		session.resourceManagementConfigErrv2 = errEmptyBluemixCredentials
		session.resourceControllerConfigErr = errEmptyBluemixCredentials
		session.resourceControllerConfigErrv2 = errEmptyBluemixCredentials
		session.powerConfigErr = errEmptyBluemixCredentials
		session.ibmpiConfigErr = errEmptyBluemixCredentials
		session.userManagementErr = errEmptyBluemixCredentials
		session.certManagementErr = errEmptyBluemixCredentials
		session.vpcClassicErr = errEmptyBluemixCredentials
		session.vpcErr = errEmptyBluemixCredentials
		session.apigatewayErr = errEmptyBluemixCredentials
		session.pDnsErr = errEmptyBluemixCredentials
		session.bmxUserFetchErr = errEmptyBluemixCredentials
		session.directlinkErr = errEmptyBluemixCredentials
		session.cosConfigErr = errEmptyBluemixCredentials
		session.transitgatewayErr = errEmptyBluemixCredentials
		session.iamNamespaceErr = errEmptyBluemixCredentials
		session.cisDNSErr = errEmptyBluemixCredentials
		session.cisGLBHealthCheckErr = errEmptyBluemixCredentials
		session.cisZonesErr = errEmptyBluemixCredentials
		session.cisRLErr = errEmptyBluemixCredentials
		session.cisEdgeFunctionErr = errEmptyBluemixCredentials

		return session, nil
	}

	if sess.BluemixSession.Config.BluemixAPIKey != "" {
		err = authenticateAPIKey(sess.BluemixSession)
		if err != nil {
			session.bmxUserFetchErr = fmt.Errorf("Error occured while fetching account user details: %q", err)
			session.functionConfigErr = fmt.Errorf("Error occured while fetching auth key for function: %q", err)
			session.powerConfigErr = fmt.Errorf("Error occured while fetching the auth key for power iaas: %q", err)
			session.ibmpiConfigErr = fmt.Errorf("Error occured while fetching the auth key for power iaas: %q", err)
		}
		err = authenticateCF(sess.BluemixSession)
		if err != nil {
			session.functionConfigErr = fmt.Errorf("Error occured while fetching auth key for function: %q", err)
		}
	}

	if sess.BluemixSession.Config.IAMAccessToken != "" && sess.BluemixSession.Config.BluemixAPIKey == "" {
		err := refreshToken(sess.BluemixSession)
		if err != nil {
			return nil, err
		}

	}
	userConfig, err := fetchUserDetails(sess.BluemixSession, c.Generation)
	if err != nil {
		session.bmxUserFetchErr = fmt.Errorf("Error occured while fetching account user details: %q", err)
	}
	session.bmxUserDetails = userConfig

	if sess.SoftLayerSession != nil && sess.SoftLayerSession.IAMToken != "" {
		sess.SoftLayerSession.IAMToken = sess.BluemixSession.Config.IAMAccessToken
		sess.SoftLayerSession.IAMRefreshToken = sess.BluemixSession.Config.IAMRefreshToken
	}

	session.functionClient, session.functionConfigErr = FunctionClient(sess.BluemixSession.Config)

	BluemixRegion = sess.BluemixSession.Config.Region

	accv1API, err := accountv1.New(sess.BluemixSession)
	if err != nil {
		session.accountV1ConfigErr = fmt.Errorf("Error occured while configuring Bluemix Accountv1 Service: %q", err)
	}
	session.bmxAccountv1ServiceAPI = accv1API

	accAPI, err := accountv2.New(sess.BluemixSession)
	if err != nil {
		session.accountConfigErr = fmt.Errorf("Error occured while configuring  Account Service: %q", err)
	}
	session.bmxAccountServiceAPI = accAPI

	cfAPI, err := mccpv2.New(sess.BluemixSession)
	if err != nil {
		session.cfConfigErr = fmt.Errorf("Error occured while configuring MCCP service: %q", err)
	}
	session.cfServiceAPI = cfAPI

	clusterAPI, err := containerv1.New(sess.BluemixSession)
	if err != nil {
		session.csConfigErr = fmt.Errorf("Error occured while configuring Container Service for K8s cluster: %q", err)
	}
	session.csServiceAPI = clusterAPI

	v2clusterAPI, err := containerv2.New(sess.BluemixSession)
	if err != nil {
		session.csv2ConfigErr = fmt.Errorf("Error occured while configuring vpc Container Service for K8s cluster: %q", err)
	}
	session.csv2ServiceAPI = v2clusterAPI

	hpcsAPI, err := hpcs.New(sess.BluemixSession)
	if err != nil {
		session.hpcsEndpointErr = fmt.Errorf("Error occured while configuring hpcs Endpoint: %q", err)
	}
	session.hpcsEndpointAPI = hpcsAPI

	kpurl := fmt.Sprintf("https://%s.kms.cloud.ibm.com", c.Region)
	options := kp.ClientConfig{
		BaseURL:       envFallBack([]string{"IBMCLOUD_KP_API_ENDPOINT"}, kpurl),
		Authorization: sess.BluemixSession.Config.IAMAccessToken,
		// InstanceID:    "42fET57nnadurKXzXAedFLOhGqETfIGYxOmQXkFgkJV9",
		Verbose: kp.VerboseFailOnly,
	}
	kpAPIclient, err := kp.New(options, kp.DefaultTransport())
	if err != nil {
		session.kpErr = fmt.Errorf("Error occured while configuring Key Protect Service: %q", err)
	}
	session.kpAPI = kpAPIclient

	kmsurl := fmt.Sprintf("https://%s.kms.cloud.ibm.com", c.Region)
	kmsOptions := kp.ClientConfig{
		BaseURL:       envFallBack([]string{"IBMCLOUD_KP_API_ENDPOINT"}, kmsurl),
		Authorization: sess.BluemixSession.Config.IAMAccessToken,
		// InstanceID:    "5af62d5d-5d90-4b84-bbcd-90d2123ae6c8",
		Verbose: kp.VerboseFailOnly,
	}
	kmsAPIclient, err := kp.New(kmsOptions, DefaultTransport())
	if err != nil {
		session.kmsErr = fmt.Errorf("Error occured while configuring key Service: %q", err)
	}
	session.kmsAPI = kmsAPIclient

	var authenticator *core.BearerTokenAuthenticator
	if strings.HasPrefix(sess.BluemixSession.Config.IAMAccessToken, "Bearer") {
		authenticator = &core.BearerTokenAuthenticator{
			BearerToken: sess.BluemixSession.Config.IAMAccessToken[7:],
		}
	} else {
		authenticator = &core.BearerTokenAuthenticator{
			BearerToken: sess.BluemixSession.Config.IAMAccessToken,
		}
	}

	vpcclassicurl := fmt.Sprintf("https://%s.iaas.cloud.ibm.com/v1", c.Region)
	vpcclassicoptions := &vpcclassic.VpcClassicV1Options{
		URL:           envFallBack([]string{"IBMCLOUD_IS_API_ENDPOINT"}, vpcclassicurl),
		Authenticator: authenticator,
	}
	vpcclassicclient, err := vpcclassic.NewVpcClassicV1(vpcclassicoptions)
	if err != nil {
		session.vpcErr = fmt.Errorf("Error occured while configuring vpc classic service: %q", err)
	}
	session.vpcClassicAPI = vpcclassicclient

	vpcurl := fmt.Sprintf("https://%s.iaas.cloud.ibm.com/v1", c.Region)
	vpcoptions := &vpc.VpcV1Options{
		URL:           envFallBack([]string{"IBMCLOUD_IS_NG_API_ENDPOINT"}, vpcurl),
		Authenticator: authenticator,
	}
	vpcclient, err := vpc.NewVpcV1(vpcoptions)
	if err != nil {
		session.vpcErr = fmt.Errorf("Error occured while configuring vpc service: %q", err)
	}
	session.vpcAPI = vpcclient

	//cosconfigurl := fmt.Sprintf("https://%s.iaas.cloud.ibm.com/v1", c.Region)
	cosconfigoptions := &cosconfig.ResourceConfigurationV1Options{
		Authenticator: authenticator,
	}
	cosconfigclient, err := cosconfig.NewResourceConfigurationV1(cosconfigoptions)
	if err != nil {
		session.cosConfigErr = fmt.Errorf("Error occured while configuring COS config service: %q", err)
	}
	session.cosConfigAPI = cosconfigclient

	schematicService, err := schematics.New(sess.BluemixSession)
	if err != nil {
		session.stxConfigErr = fmt.Errorf("Error occured while fetching schematics Configuration: %q", err)
	}
	session.stxServiceAPI = schematicService

	cisAPI, err := cisv1.New(sess.BluemixSession)
	if err != nil {
		session.cisConfigErr = fmt.Errorf("Error occured while configuring Cloud Internet Services: %q", err)
	}
	session.cisServiceAPI = cisAPI

	globalSearchAPI, err := globalsearchv2.New(sess.BluemixSession)
	if err != nil {
		session.globalSearchConfigErr = fmt.Errorf("Error occured while configuring Global Search: %q", err)
	}
	session.globalSearchServiceAPI = globalSearchAPI

	globalTaggingAPI, err := globaltaggingv3.New(sess.BluemixSession)
	if err != nil {
		session.globalTaggingConfigErr = fmt.Errorf("Error occured while configuring Global Tagging: %q", err)
	}
	session.globalTaggingServiceAPI = globalTaggingAPI

	iampap, err := iampapv1.New(sess.BluemixSession)
	if err != nil {
		session.iamPAPConfigErr = fmt.Errorf("Error occured while configuring Bluemix IAMPAP Service: %q", err)
	}
	session.iamPAPServiceAPI = iampap

	iampapv2, err := iampapv2.New(sess.BluemixSession)
	if err != nil {
		session.iamPAPConfigErrv2 = fmt.Errorf("Error occured while configuring Bluemix IAMPAP Service: %q", err)
	}
	session.iamPAPServiceAPIv2 = iampapv2

	iam, err := iamv1.New(sess.BluemixSession)
	if err != nil {
		session.iamConfigErr = fmt.Errorf("Error occured while configuring Bluemix IAM Service: %q", err)
	}
	session.iamServiceAPI = iam

	iamuum, err := iamuumv1.New(sess.BluemixSession)
	if err != nil {
		session.iamUUMConfigErr = fmt.Errorf("Error occured while configuring Bluemix IAMUUM Service: %q", err)
	}
	session.iamUUMServiceAPI = iamuum

	iamuumv2, err := iamuumv2.New(sess.BluemixSession)
	if err != nil {
		session.iamUUMConfigErrV2 = fmt.Errorf("Error occured while configuring Bluemix IAMUUM Service: %q", err)
	}
	session.iamUUMServiceAPIV2 = iamuumv2

	icdAPI, err := icdv4.New(sess.BluemixSession)
	if err != nil {
		session.icdConfigErr = fmt.Errorf("Error occured while configuring IBM Cloud Database Services: %q", err)
	}
	session.icdServiceAPI = icdAPI

	resourceCatalogAPI, err := catalog.New(sess.BluemixSession)
	if err != nil {
		session.resourceCatalogConfigErr = fmt.Errorf("Error occured while configuring Resource Catalog service: %q", err)
	}
	session.resourceCatalogServiceAPI = resourceCatalogAPI

	resourceManagementAPI, err := management.New(sess.BluemixSession)
	if err != nil {
		session.resourceManagementConfigErr = fmt.Errorf("Error occured while configuring Resource Management service: %q", err)
	}
	session.resourceManagementServiceAPI = resourceManagementAPI

	resourceManagementAPIv2, err := managementv2.New(sess.BluemixSession)
	if err != nil {
		session.resourceManagementConfigErrv2 = fmt.Errorf("Error occured while configuring Resource Management service: %q", err)
	}
	session.resourceManagementServiceAPIv2 = resourceManagementAPIv2

	resourceControllerAPI, err := controller.New(sess.BluemixSession)
	if err != nil {
		session.resourceControllerConfigErr = fmt.Errorf("Error occured while configuring Resource Controller service: %q", err)
	}
	session.resourceControllerServiceAPI = resourceControllerAPI

	ResourceControllerAPIv2, err := controllerv2.New(sess.BluemixSession)
	if err != nil {
		session.resourceControllerConfigErrv2 = fmt.Errorf("Error occured while configuring Resource Controller v2 service: %q", err)
	}
	session.resourceControllerServiceAPIv2 = ResourceControllerAPIv2

	userManagementAPI, err := usermanagementv2.New(sess.BluemixSession)
	if err != nil {
		session.userManagementErr = fmt.Errorf("Error occured while configuring user management service: %q", err)
	}
	session.userManagementAPI = userManagementAPI
	certManagementAPI, err := certificatemanager.New(sess.BluemixSession)
	if err != nil {
		session.certManagementErr = fmt.Errorf("Error occured while configuring Certificate manager service: %q", err)
	}
	session.certManagementAPI = certManagementAPI

	apicurl := fmt.Sprintf("https://api.%s.apigw.cloud.ibm.com/controller", c.Region)
	APIGatewayControllerAPIV1Options := &apigateway.ApiGatewayControllerApiV1Options{
		URL:           envFallBack([]string{"IBMCLOUD_API_GATEWAY_ENDPOINT"}, apicurl),
		Authenticator: &core.NoAuthAuthenticator{},
	}
	apigatewayAPI, err := apigateway.NewApiGatewayControllerApiV1(APIGatewayControllerAPIV1Options)
	if err != nil {
		session.apigatewayErr = fmt.Errorf("Error occured while configuring  APIGateway service: %q", err)
	}
	session.apigatewayAPI = apigatewayAPI

	ibmpisession, err := ibmpisession.New(sess.BluemixSession.Config.IAMAccessToken, c.Region, false, c.BluemixTimeout, session.bmxUserDetails.userAccount, c.Zone)
	if err != nil {
		session.ibmpiConfigErr = err
		return nil, err
	}

	session.ibmpiSession = ibmpisession

	bluemixToken := ""
	if strings.HasPrefix(sess.BluemixSession.Config.IAMAccessToken, "Bearer") {
		bluemixToken = sess.BluemixSession.Config.IAMAccessToken[7:len(sess.BluemixSession.Config.IAMAccessToken)]
	} else {
		bluemixToken = sess.BluemixSession.Config.IAMAccessToken
	}

	dnsOptions := &dns.DnsSvcsV1Options{
		URL: envFallBack([]string{"IBMCLOUD_PRIVATE_DNS_API_ENDPOINT"}, "https://api.dns-svcs.cloud.ibm.com/v1"),
		Authenticator: &core.BearerTokenAuthenticator{
			BearerToken: bluemixToken,
		},
	}

	session.pDnsClient, session.pDnsErr = dns.NewDnsSvcsV1(dnsOptions)
	if session.pDnsErr != nil {
		session.pDnsErr = fmt.Errorf("Error occured while configuring PrivateDNS Service: %s", session.pDnsErr)
	}
	version := time.Now().Format("2006-01-02")

	directlinkOptions := &dl.DirectLinkV1Options{
		URL: envFallBack([]string{"IBMCLOUD_DL_API_ENDPOINT"}, "https://directlink.cloud.ibm.com/v1"),
		Authenticator: &core.BearerTokenAuthenticator{
			BearerToken: bluemixToken,
		},
		Version: &version,
	}

	session.directlinkAPI, session.directlinkErr = dl.NewDirectLinkV1(directlinkOptions)
	if session.directlinkErr != nil {
		session.directlinkErr = fmt.Errorf("Error occured while configuring Direct Link Service: %s", session.directlinkErr)
	}

	transitgatewayOptions := &tg.TransitGatewayApisV1Options{
		URL: envFallBack([]string{"IBMCLOUD_TG_API_ENDPOINT"}, "https://transit.cloud.ibm.com/v1"),
		Authenticator: &core.BearerTokenAuthenticator{
			BearerToken: bluemixToken,
		},
		Version: CreateVersionDate(),
	}

	session.transitgatewayAPI, session.transitgatewayErr = tg.NewTransitGatewayApisV1(transitgatewayOptions)
	if session.transitgatewayErr != nil {
		session.transitgatewayErr = fmt.Errorf("Error occured while configuring Transit Gateway Service: %s", session.transitgatewayErr)
	}

	cfcurl := fmt.Sprintf("https://%s.functions.cloud.ibm.com/api/v1", c.Region)
	ibmCloudFunctionsNamespaceOptions := &ns.IbmCloudFunctionsNamespaceOptions{
		URL: envFallBack([]string{"IBMCLOUD_NAMESPACE_API_ENDPOINT"}, cfcurl),
		Authenticator: &core.BearerTokenAuthenticator{
			BearerToken: bluemixToken,
		},
	}

	session.iamNamespaceAPI, err = ns.NewIbmCloudFunctionsNamespaceAPIV1(ibmCloudFunctionsNamespaceOptions)
	if err != nil {
		session.iamNamespaceErr = fmt.Errorf("Error occured while configuring IAM namespace service: %q", err)
	}

	// CIS Service instances starts here.
	cisEndPoint := envFallBack([]string{"IBMCLOUD_CIS_API_ENDPOINT"}, "https://api.cis.cloud.ibm.com")

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

	// IBM Network CIS DNS Record service
	cisDNSRecordsOpt := &cisdnsrecordsv1.DnsRecordsV1Options{
		URL:            cisEndPoint,
		Crn:            core.StringPtr(""),
		ZoneIdentifier: core.StringPtr(""),
		Authenticator:  authenticator,
	}
	session.cisDNSRecordsClient, session.cisDNSErr = cisdnsrecordsv1.NewDnsRecordsV1(cisDNSRecordsOpt)
	if session.cisDNSErr != nil {
		session.cisDNSErr = fmt.Errorf("Error occured while configuring CIS DNS Service: %s", session.cisDNSErr)
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
			fmt.Errorf("Error occured while configuring CIS GLB Health Check service: %s",
				session.cisGLBHealthCheckErr)
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
			fmt.Errorf("Error occured while configuring CIS Edge Function service: %s",
				session.cisEdgeFunctionErr)
	}
	return session, nil
}

// CreateVersionDate requires mandatory version attribute. Any date from 2019-12-13 up to the currentdate may be provided. Specify the current date to request the latest version.
func CreateVersionDate() *string {
	version := time.Now().Format("2006-01-02")
	return &version
}

func newSession(c *Config) (*Session, error) {
	ibmSession := &Session{}

	softlayerSession := &slsession.Session{
		Endpoint:  c.SoftLayerEndpointURL,
		Timeout:   c.SoftLayerTimeout,
		UserName:  c.SoftLayerUserName,
		APIKey:    c.SoftLayerAPIKey,
		Debug:     os.Getenv("TF_LOG") != "",
		Retries:   c.RetryCount,
		RetryWait: c.RetryDelay,
	}

	if c.IAMToken != "" {
		log.Println("Configuring SoftLayer Session with token")
		softlayerSession.IAMToken = c.IAMToken
		softlayerSession.IAMRefreshToken = c.IAMRefreshToken
	}
	if c.SoftLayerAPIKey != "" && c.SoftLayerUserName != "" {
		log.Println("Configuring SoftLayer Session with API key")
		softlayerSession.APIKey = c.SoftLayerAPIKey
		softlayerSession.UserName = c.SoftLayerUserName
	}
	softlayerSession.AppendUserAgent(fmt.Sprintf("terraform-provider-ibm/%s", version.Version))
	ibmSession.SoftLayerSession = softlayerSession

	if (c.IAMToken != "" && c.IAMRefreshToken == "") || (c.IAMToken == "" && c.IAMRefreshToken != "") {
		return nil, fmt.Errorf("iam_token and iam_refresh_token must be provided")
	}

	if c.IAMToken != "" && c.IAMRefreshToken != "" {
		log.Println("Configuring IBM Cloud Session with token")
		var sess *bxsession.Session
		bmxConfig := &bluemix.Config{
			IAMAccessToken:  c.IAMToken,
			IAMRefreshToken: c.IAMRefreshToken,
			//Comment out debug mode for v0.12
			//Debug:           os.Getenv("TF_LOG") != "",
			HTTPTimeout:   c.BluemixTimeout,
			Region:        c.Region,
			ResourceGroup: c.ResourceGroup,
			RetryDelay:    &c.RetryDelay,
			MaxRetries:    &c.RetryCount,
		}
		sess, err := bxsession.New(bmxConfig)
		if err != nil {
			return nil, err
		}
		ibmSession.BluemixSession = sess
	}

	if c.BluemixAPIKey != "" {
		log.Println("Configuring IBM Cloud Session with API key")
		var sess *bxsession.Session
		bmxConfig := &bluemix.Config{
			BluemixAPIKey: c.BluemixAPIKey,
			//Comment out debug mode for v0.12
			//Debug:         os.Getenv("TF_LOG") != "",
			HTTPTimeout:   c.BluemixTimeout,
			Region:        c.Region,
			ResourceGroup: c.ResourceGroup,
			RetryDelay:    &c.RetryDelay,
			MaxRetries:    &c.RetryCount,
			//PowerServiceInstance: c.PowerServiceInstance,
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

func authenticateCF(sess *bxsession.Session) error {
	config := sess.Config
	tokenRefresher, err := authentication.NewUAARepository(config, &rest.Client{
		DefaultHeader: gohttp.Header{
			"User-Agent": []string{http.UserAgent()},
		},
	})
	if err != nil {
		return err
	}
	return tokenRefresher.AuthenticateAPIKey(config.BluemixAPIKey)
}

func fetchUserDetails(sess *bxsession.Session, generation int) (*UserConfig, error) {
	config := sess.Config
	user := UserConfig{}
	var bluemixToken string

	if strings.HasPrefix(config.IAMAccessToken, "Bearer") {
		bluemixToken = config.IAMAccessToken[7:len(config.IAMAccessToken)]
	} else {
		bluemixToken = config.IAMAccessToken
	}

	token, err := jwt.Parse(bluemixToken, func(token *jwt.Token) (interface{}, error) {
		return "", nil
	})
	//TODO validate with key
	if err != nil && !strings.Contains(err.Error(), "key is of invalid type") {
		return &user, err
	}
	claims := token.Claims.(jwt.MapClaims)
	if email, ok := claims["email"]; ok {
		user.userEmail = email.(string)
	}
	user.userID = claims["id"].(string)
	user.userAccount = claims["account"].(map[string]interface{})["bss"].(string)
	iss := claims["iss"].(string)
	if strings.Contains(iss, "https://iam.cloud.ibm.com") {
		user.cloudName = "bluemix"
	} else {
		user.cloudName = "staging"
	}
	user.cloudType = "public"

	user.generation = generation
	return &user, nil
}

func refreshToken(sess *bxsession.Session) error {
	config := sess.Config
	tokenRefresher, err := authentication.NewIAMAuthRepository(config, &rest.Client{
		DefaultHeader: gohttp.Header{
			"User-Agent": []string{http.UserAgent()},
		},
	})
	if err != nil {
		return err
	}
	_, err = tokenRefresher.RefreshToken()
	return err
}

func envFallBack(envs []string, defaultValue string) string {
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

// DefaultTransport ...
func DefaultTransport() gohttp.RoundTripper {
	transport := &gohttp.Transport{
		Proxy:               gohttp.ProxyFromEnvironment,
		DisableKeepAlives:   true,
		MaxIdleConnsPerHost: -1,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false,
		},
	}
	return transport
}
