package ibm

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/IBM-Cloud/bluemix-go/api/iam/iamv1"

	gohttp "net/http"

	"github.com/apache/incubator-openwhisk-client-go/whisk"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/IBM-Cloud/terraform-provider-ibm/version"
	slsession "github.com/softlayer/softlayer-go/session"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/api/account/accountv1"
	"github.com/IBM-Cloud/bluemix-go/api/account/accountv2"
	"github.com/IBM-Cloud/bluemix-go/api/container/containerv1"
	"github.com/IBM-Cloud/bluemix-go/api/iampap/iampapv1"
	"github.com/IBM-Cloud/bluemix-go/api/iamuum/iamuumv1"
	"github.com/IBM-Cloud/bluemix-go/api/mccp/mccpv2"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/catalog"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/controller"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev1/management"
	"github.com/IBM-Cloud/bluemix-go/authentication"
	"github.com/IBM-Cloud/bluemix-go/http"
	"github.com/IBM-Cloud/bluemix-go/rest"

	bxsession "github.com/IBM-Cloud/bluemix-go/session"
)

//RetryDelay
const RetryAPIDelay = 5 * time.Second

//BluemixRegion ...
var BluemixRegion string

var (
	errEmptySoftLayerCredentials = errors.New("softlayer_username and softlayer_api_key must be provided. Please see the documentation on how to configure them")
	errEmptyBluemixCredentials   = errors.New("bluemix_api_key must be provided. Please see the documentation on how to configure it")
)

//UserConfig ...
type UserConfig struct {
	userID      string
	userEmail   string
	userAccount string
}

//Config stores user provider input
type Config struct {
	//BluemixAPIKey is the Bluemix api key
	BluemixAPIKey string
	//Bluemix region
	Region string
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
	SoftLayerSession() *slsession.Session
	BluemixSession() (*bxsession.Session, error)
	ContainerAPI() (containerv1.ContainerServiceAPI, error)
	IAMAPI() (iamv1.IAMServiceAPI, error)
	IAMPAPAPI() (iampapv1.IAMPAPAPI, error)
	IAMUUMAPI() (iamuumv1.IAMUUMServiceAPI, error)
	MccpAPI() (mccpv2.MccpServiceAPI, error)
	BluemixAcccountAPI() (accountv2.AccountServiceAPI, error)
	BluemixAcccountv1API() (accountv1.AccountServiceAPI, error)
	BluemixUserDetails() (*UserConfig, error)
	FunctionClient() (*whisk.Client, error)
	ResourceCatalogAPI() (catalog.ResourceCatalogAPI, error)
	ResourceManagementAPI() (management.ResourceManagementAPI, error)
	ResourceControllerAPI() (controller.ResourceControllerAPI, error)
}

type clientSession struct {
	session *Session

	csConfigErr  error
	csServiceAPI containerv1.ContainerServiceAPI

	cfConfigErr  error
	cfServiceAPI mccpv2.MccpServiceAPI

	iamPAPConfigErr  error
	iamPAPServiceAPI iampapv1.IAMPAPAPI

	iamUUMConfigErr  error
	iamUUMServiceAPI iamuumv1.IAMUUMServiceAPI

	iamConfigErr  error
	iamServiceAPI iamv1.IAMServiceAPI

	accountConfigErr     error
	bmxAccountServiceAPI accountv2.AccountServiceAPI

	accountV1ConfigErr     error
	bmxAccountv1ServiceAPI accountv1.AccountServiceAPI

	bmxUserDetails  *UserConfig
	bmxUserFetchErr error

	functionConfigErr error
	functionClient    *whisk.Client

	resourceControllerConfigErr  error
	resourceControllerServiceAPI controller.ResourceControllerAPI

	resourceManagementConfigErr  error
	resourceManagementServiceAPI management.ResourceManagementAPI

	resourceCatalogConfigErr  error
	resourceCatalogServiceAPI catalog.ResourceCatalogAPI
}

// SoftLayerSession providers SoftLayer Session
func (sess clientSession) SoftLayerSession() *slsession.Session {
	return sess.session.SoftLayerSession
}

// MccpAPI provides Multi Cloud Controller Proxy APIs ...
func (sess clientSession) MccpAPI() (mccpv2.MccpServiceAPI, error) {
	return sess.cfServiceAPI, sess.cfConfigErr
}

// BluemixAcccountAPI ...
func (sess clientSession) BluemixAcccountAPI() (accountv2.AccountServiceAPI, error) {
	return sess.bmxAccountServiceAPI, sess.accountConfigErr
}

// BluemixAcccountAPI ...
func (sess clientSession) BluemixAcccountv1API() (accountv1.AccountServiceAPI, error) {
	return sess.bmxAccountv1ServiceAPI, sess.accountV1ConfigErr
}

// IAMAPI provides IAM PAP APIs ...
func (sess clientSession) IAMAPI() (iamv1.IAMServiceAPI, error) {
	return sess.iamServiceAPI, sess.iamConfigErr
}

// IAMPAPAPI provides IAM PAP APIs ...
func (sess clientSession) IAMPAPAPI() (iampapv1.IAMPAPAPI, error) {
	return sess.iamPAPServiceAPI, sess.iamPAPConfigErr
}

// IAMUUMAPI provides IAM UUM APIs ...
func (sess clientSession) IAMUUMAPI() (iamuumv1.IAMUUMServiceAPI, error) {
	return sess.iamUUMServiceAPI, sess.iamUUMConfigErr
}

// ContainerAPI provides Container Service APIs ...
func (sess clientSession) ContainerAPI() (containerv1.ContainerServiceAPI, error) {
	return sess.csServiceAPI, sess.csConfigErr
}

// BluemixSession to provide the Bluemix Session
func (sess clientSession) BluemixSession() (*bxsession.Session, error) {
	return sess.session.BluemixSession, sess.cfConfigErr
}

// BluemixUserDetails ...
func (sess clientSession) BluemixUserDetails() (*UserConfig, error) {
	return sess.bmxUserDetails, sess.bmxUserFetchErr
}

// FunctionClient ...
func (sess clientSession) FunctionClient() (*whisk.Client, error) {
	return sess.functionClient, sess.functionConfigErr
}

// ResourceCatalogAPI ...
func (sess clientSession) ResourceCatalogAPI() (catalog.ResourceCatalogAPI, error) {
	return sess.resourceCatalogServiceAPI, sess.resourceCatalogConfigErr
}

// ResourceManagementAPI ...
func (sess clientSession) ResourceManagementAPI() (management.ResourceManagementAPI, error) {
	return sess.resourceManagementServiceAPI, sess.resourceManagementConfigErr
}

// ResourceControllerAPI ...
func (sess clientSession) ResourceControllerAPI() (controller.ResourceControllerAPI, error) {
	return sess.resourceControllerServiceAPI, sess.resourceControllerConfigErr
}

// ClientSession configures and returns a fully initialized ClientSession
func (c *Config) ClientSession() (interface{}, error) {
	sess, err := newSession(c)
	if err != nil {
		return nil, err
	}
	session := clientSession{
		session: sess,
	}

	if sess.BluemixSession == nil {
		//Can be nil only  if bluemix_api_key is not provided
		log.Println("Skipping Bluemix Clients configuration")
		session.csConfigErr = errEmptyBluemixCredentials
		session.cfConfigErr = errEmptyBluemixCredentials
		session.accountConfigErr = errEmptyBluemixCredentials
		session.accountV1ConfigErr = errEmptyBluemixCredentials
		session.iamConfigErr = errEmptyBluemixCredentials
		session.functionConfigErr = errEmptyBluemixCredentials
		session.iamPAPConfigErr = errEmptyBluemixCredentials
		session.iamUUMConfigErr = errEmptyBluemixCredentials
		session.resourceCatalogConfigErr = errEmptyBluemixCredentials
		session.resourceManagementConfigErr = errEmptyBluemixCredentials
		session.resourceCatalogConfigErr = errEmptyBluemixCredentials

		return session, nil
	}
	err = authenticateAPIKey(sess.BluemixSession)
	if err != nil {
		session.bmxUserFetchErr = fmt.Errorf("Error occured while fetching account user details: %q", err)
		session.functionConfigErr = fmt.Errorf("Error occured while fetching auth key for function: %q", err)
	} else {
		userConfig, err := fetchUserDetails(sess.BluemixSession)
		if err != nil {
			session.bmxUserFetchErr = fmt.Errorf("Error occured while fetching account user details: %q", err)
		}
		session.bmxUserDetails = userConfig

		session.functionClient, session.functionConfigErr = FunctionClient(sess.BluemixSession.Config, c.FunctionNameSpace)
	}
	sess.BluemixSession.Config.UAAAccessToken = ""
	sess.BluemixSession.Config.UAARefreshToken = ""

	BluemixRegion = sess.BluemixSession.Config.Region
	cfAPI, err := mccpv2.New(sess.BluemixSession)
	if err != nil {
		session.cfConfigErr = fmt.Errorf("Error occured while configuring MCCP service: %q", err)
	}
	session.cfServiceAPI = cfAPI

	accAPI, err := accountv2.New(sess.BluemixSession)
	if err != nil {
		session.accountConfigErr = fmt.Errorf("Error occured while configuring  Account Service: %q", err)
	}
	session.bmxAccountServiceAPI = accAPI

	clusterAPI, err := containerv1.New(sess.BluemixSession)
	if err != nil {
		session.csConfigErr = fmt.Errorf("Error occured while configuring Container Service for K8s cluster: %q", err)
	}
	session.csServiceAPI = clusterAPI

	accv1API, err := accountv1.New(sess.BluemixSession)
	if err != nil {
		session.accountV1ConfigErr = fmt.Errorf("Error occured while configuring Bluemix Accountv1 Service: %q", err)
	}
	session.bmxAccountv1ServiceAPI = accv1API

	iampap, err := iampapv1.New(sess.BluemixSession)
	if err != nil {
		session.iamPAPConfigErr = fmt.Errorf("Error occured while configuring Bluemix IAMPAP Service: %q", err)
	}
	session.iamPAPServiceAPI = iampap

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

	resourceControllerAPI, err := controller.New(sess.BluemixSession)
	if err != nil {
		session.resourceControllerConfigErr = fmt.Errorf("Error occured while configuring Resource Controller service: %q", err)
	}
	session.resourceControllerServiceAPI = resourceControllerAPI

	return session, nil
}

func newSession(c *Config) (*Session, error) {
	ibmSession := &Session{}

	log.Println("Configuring SoftLayer Session")
	softlayerSession := &slsession.Session{
		Endpoint:  c.SoftLayerEndpointURL,
		Timeout:   c.SoftLayerTimeout,
		UserName:  c.SoftLayerUserName,
		APIKey:    c.SoftLayerAPIKey,
		Debug:     os.Getenv("TF_LOG") != "",
		Retries:   c.RetryCount,
		RetryWait: c.RetryDelay,
	}
	softlayerSession.AppendUserAgent(fmt.Sprintf("terraform-provider-ibm/%s", version.Version))
	ibmSession.SoftLayerSession = softlayerSession

	if c.BluemixAPIKey != "" {
		log.Println("Configuring Bluemix Session")
		var sess *bxsession.Session
		bmxConfig := &bluemix.Config{
			BluemixAPIKey: c.BluemixAPIKey,
			Debug:         os.Getenv("TF_LOG") != "",
			HTTPTimeout:   c.BluemixTimeout,
			Region:        c.Region,
			RetryDelay:    &c.RetryDelay,
			MaxRetries:    &c.RetryCount,
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

func fetchUserDetails(sess *bxsession.Session) (*UserConfig, error) {
	config := sess.Config
	user := UserConfig{}

	bluemixToken := config.IAMAccessToken[7:len(config.IAMAccessToken)]
	token, err := jwt.Parse(bluemixToken, func(token *jwt.Token) (interface{}, error) {
		return "", nil
	})
	//TODO validate with key
	if err != nil && !strings.Contains(err.Error(), "key is of invalid type") {
		return &user, err
	}
	claims := token.Claims.(jwt.MapClaims)
	user.userEmail = claims["email"].(string)
	user.userID = claims["id"].(string)
	user.userAccount = claims["account"].(map[string]interface{})["bss"].(string)
	return &user, nil
}
