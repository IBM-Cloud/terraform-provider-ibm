// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package conns

import (
	"fmt"
	"log"
	gohttp "net/http"
	"os"
	"strings"
	"time"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/api/globaltagging/globaltaggingv3"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev2/managementv2"
	"github.com/IBM-Cloud/bluemix-go/authentication"
	"github.com/IBM-Cloud/bluemix-go/http"
	"github.com/IBM-Cloud/bluemix-go/rest"
	bxsession "github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/terraform-provider-ibm/version"
	"github.com/IBM/go-sdk-core/v5/core"
	resourcecontroller "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.ibm.com/DRAutomation/dra-go-sdk/drautomationservicev1"
)

// Config stores user provider input.
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

// UserConfig ...
type UserConfig struct {
	userID      string
	userEmail   string
	UserAccount string
	cloudName   string `default:"bluemix"`
	cloudType   string `default:"public"`
}

// Session stores the information required for communication with the SoftLayer and Bluemix API
type Session struct {
	// BluemixSession is the the Bluemix session used to connect to the Bluemix API
	BluemixSession *bxsession.Session
}

// ClientSession ...
type ClientSession interface {
	BluemixUserDetails() (*UserConfig, error)
	GlobalTaggingAPI() (globaltaggingv3.GlobalTaggingServiceAPI, error)
	ResourceControllerV2API() (*resourcecontroller.ResourceControllerV2, error)
	ResourceManagementAPIv2() (managementv2.ResourceManagementAPIv2, error)
	DrAutomationServiceV1() (*drautomationservicev1.DrAutomationServiceV1, error)
}

type clientSession struct {
	session *Session

	bmxUserDetails  *UserConfig
	bmxUserFetchErr error

	// Global Tagging API from Bluemix Go SDK.
	globalTaggingConfigErr  error
	globalTaggingServiceAPI globaltaggingv3.GlobalTaggingServiceAPI

	// Resource Controller API from Platform-Services Go SDK.
	resourceControllerErr error
	resourceControllerAPI *resourcecontroller.ResourceControllerV2

	// Resource Management API from Bluemix Go SDK.
	resourceManagementConfigErrv2  error
	resourceManagementServiceAPIv2 managementv2.ResourceManagementAPIv2

	drAutomationServiceClient    *drautomationservicev1.DrAutomationServiceV1
	drAutomationServiceClientErr error
}

func (session clientSession) BluemixUserDetails() (*UserConfig, error) {
	return session.bmxUserDetails, session.bmxUserFetchErr
}

// GlobalTaggingAPI provides Global Tagging API from the Bluemix Go SDK.
func (sess clientSession) GlobalTaggingAPI() (globaltaggingv3.GlobalTaggingServiceAPI, error) {
	return sess.globalTaggingServiceAPI, sess.globalTaggingConfigErr
}

// ResourceControllerV2API provides the Resource Controller service from the Platform-Services Go SDK.
func (sess clientSession) ResourceControllerV2API() (*resourcecontroller.ResourceControllerV2, error) {
	return sess.resourceControllerAPI, sess.resourceControllerErr
}

// ResourceManagementAPIv2 provides the Resource Management service from the Bluemix Go SDK.
func (sess clientSession) ResourceManagementAPIv2() (managementv2.ResourceManagementAPIv2, error) {
	return sess.resourceManagementServiceAPIv2, sess.resourceManagementConfigErrv2
}

// DrAutomation Service
func (session clientSession) DrAutomationServiceV1() (*drautomationservicev1.DrAutomationServiceV1, error) {
	return session.drAutomationServiceClient, session.drAutomationServiceClientErr
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

	if sess.BluemixSession != nil && sess.BluemixSession.Config.BluemixAPIKey != "" {
		err = authenticateAPIKey(sess.BluemixSession)
		if err != nil {
			session.bmxUserFetchErr = fmt.Errorf("Error occured while fetching account user details: %q", err)
		}
	}

	resourceManagementAPIv2, err := managementv2.New(sess.BluemixSession)
	if err != nil {
		session.resourceManagementConfigErrv2 = fmt.Errorf("Error occured while configuring Resource Management service: %q", err)
	}
	session.resourceManagementServiceAPIv2 = resourceManagementAPIv2

	// Global Tagging API from Bluemix Go SDK.
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

	var fileMap map[string]interface{}

	// RESOURCE CONTROLLER API from Platform-Services Go SDK.
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

	// Construct an instance of the 'DrAutomation Service' service.
	if session.drAutomationServiceClientErr == nil {
		// Construct the service options.
		drAutomationServiceClientOptions := &drautomationservicev1.DrAutomationServiceV1Options{
			Authenticator: authenticator,
		}

		// Construct the service client.
		session.drAutomationServiceClient, err = drautomationservicev1.NewDrAutomationServiceV1(drAutomationServiceClientOptions)
		if err == nil {
			// Enable retries for API calls
			session.drAutomationServiceClient.Service.EnableRetries(c.RetryCount, c.RetryDelay)
			// Add custom header for analytics
			session.drAutomationServiceClient.SetDefaultHeaders(gohttp.Header{
				"X-Original-User-Agent": {fmt.Sprintf("terraform-provider-ibm/%s", version.Version)},
			})
		} else {
			session.drAutomationServiceClientErr = fmt.Errorf("Error occurred while constructing 'DrAutomation Service' service client: %q", err)
		}
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

// nolint deadcode Might be used by generated code.
func EnvFallBack(envs []string, defaultValue string) string {
	for _, k := range envs {
		if v := os.Getenv(k); v != "" {
			return v
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
