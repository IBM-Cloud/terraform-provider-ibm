package ibm

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM/go-sdk-core/core"
	"github.com/apache/openwhisk-client-go/whisk"
	namespaceapi "github.ibm.com/ibmcloud/namespace-go-sdk/ibmcloudfunctionsnamespaceapiv1"
)

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://us-south.functions.cloud.ibm.com"

//FunctionClient ...
func FunctionClient(c *bluemix.Config) (*whisk.Client, error) {
	baseEndpoint := getBaseURL(c.Region)
	u, _ := url.Parse(fmt.Sprintf("%s/api", baseEndpoint))

	functionsClient, err := whisk.NewClient(http.DefaultClient, &whisk.Config{
		Host:    u.Host,
		Version: "v1",
	})

	return functionsClient, err
}

//getBaseURL ..
func getBaseURL(region string) string {
	baseEndpoint := fmt.Sprintf(DefaultServiceURL)
	if region != "us-south" {
		baseEndpoint = fmt.Sprintf("https://%s.functions.cloud.ibm.com", region)
	}

	return baseEndpoint
}

/*
 *
 * Configure a HTTP client using the OpenWhisk properties (i.e. host, auth, iamtoken)
 * Only cf-based namespaces needs auth key value.
 * iam-based namespace don't have an auth key and needs only iam token for authorization.
 *
 */
func setupOpenWhiskClientConfig(namespace string, c *bluemix.Config, wskClient *whisk.Client) (*whisk.Client, error) {
	if c.UAAAccessToken == "" && c.UAARefreshToken == "" {
		return nil, fmt.Errorf("Couldn't retrieve auth key for IBM Cloud Function")
	}

	baseEndpoint := getBaseURL(c.Region)
	apiOptions := &namespaceapi.IbmCloudFunctionsNamespaceOptions{
		URL:           fmt.Sprintf("%s", baseEndpoint),
		Authenticator: &core.NoAuthAuthenticator{},
	}

	payload := make(map[string]string)
	payload["Authorization"] = c.IAMAccessToken
	namespaceOptions := &namespaceapi.GetNamespacesOptions{
		Headers: payload,
	}

	nsService, _ := namespaceapi.NewIbmCloudFunctionsNamespaceAPIV1(apiOptions)
	nsList, _, err := nsService.GetNamespaces(namespaceOptions)
	if err != nil {
		return nil, err
	}

	var validNamespace bool
	var isCFNamespace bool
	allNamespaces := []string{}
	for _, n := range nsList.Namespaces {
		allNamespaces = append(allNamespaces, n.GetName())
		if n.GetName() == namespace || n.GetID() == namespace {
			if os.Getenv("TF_LOG") != "" {
				whisk.SetDebug(true)
			}
			if n.IsCf() {
				isCFNamespace = true
				break
			}
			validNamespace = true
			// Configure whisk properties to handle iam-based/iam-migrated  namespaces.
			if n.IsIamEnabled() {
				additionalHeaders := make(http.Header)
				additionalHeaders.Add("Authorization", c.IAMAccessToken)
				additionalHeaders.Add("X-Namespace-Id", n.GetID())

				wskClient.Config.Namespace = n.GetID()
				wskClient.Config.AdditionalHeaders = additionalHeaders
				return wskClient, nil
			}
		}
	}

	// Configure whisk properties to handle cf-based namespaces.
	if isCFNamespace {
		err := validateNamespace(namespace)
		if err != nil {
			return nil, err
		}
		payload = make(map[string]string)
		payload["accessToken"] = c.UAAAccessToken[7:len(c.UAAAccessToken)]
		payload["refreshToken"] = c.UAARefreshToken[7:len(c.UAARefreshToken)]
		namespaceOptions = &namespaceapi.GetNamespacesOptions{
			Headers: payload,
		}

		nsService, err := namespaceapi.NewIbmCloudFunctionsNamespaceAPIV1(apiOptions)
		nsList, _, err := nsService.GetCloudFoundaryNamespaces(namespaceOptions)
		if err != nil {
			return nil, err
		}

		for _, n := range nsList.Namespaces {
			if n.GetName() == namespace {
				wskClient.Config.Namespace = n.GetName()
				wskClient.Config.AuthToken = fmt.Sprintf("%s:%s", n.GetUUID(), n.GetKey())
				return wskClient, nil
			}
		}
	}

	if !validNamespace {
		return nil, fmt.Errorf("Namespace '%s' is not in the list of entitled namespaces. Available namespaces are %s", namespace, allNamespaces)
	}

	return nil, fmt.Errorf("Failed to create whisk config object for namespace '%s'", namespace)
}
