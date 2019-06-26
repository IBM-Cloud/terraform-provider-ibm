package csev2

import (
	gohttp "net/http"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/authentication"
	"github.com/IBM-Cloud/bluemix-go/client"
	"github.com/IBM-Cloud/bluemix-go/http"
	"github.com/IBM-Cloud/bluemix-go/rest"
	"github.com/IBM-Cloud/bluemix-go/session"
)

//const ErrCodeAPICreation = "APICreationError"

type CseServiceAPI interface {
	ServiceEndpoints() ServiceEndpoints
}

type cseService struct {
	*client.Client
}

func New(sess *session.Session) (CseServiceAPI, error) {
	config := sess.Config.Copy()
	err := config.ValidateConfigForService(bluemix.CseService)
	if err != nil {
		return nil, err
	}
	if config.HTTPClient == nil {
		config.HTTPClient = http.NewHTTPClient(config)
	}
	tokenRefreher, err := authentication.NewIAMAuthRepository(config, &rest.Client{
		DefaultHeader: gohttp.Header{
			"User-Agent": []string{http.UserAgent()},
		},
		HTTPClient: config.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	if config.IAMAccessToken == "" {
		err := authentication.PopulateTokens(tokenRefreher, config)
		if err != nil {
			return nil, err
		}
	}
	if config.Endpoint == nil {
		ep, err := config.EndpointLocator.CseEndpoint()
		if err != nil {
			return nil, err
		}
		config.Endpoint = &ep
	}

	return &cseService{
		Client: client.New(config, bluemix.CseService, tokenRefreher),
	}, nil
}

func (c *cseService) ServiceEndpoints() ServiceEndpoints {
	return newServiceEndpointsAPI(c.Client)
}
