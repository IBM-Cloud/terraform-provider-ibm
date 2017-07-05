package iampapv1

import (
	gohttp "net/http"

	bluemix "github.com/IBM-Bluemix/bluemix-go"
	"github.com/IBM-Bluemix/bluemix-go/authentication"
	"github.com/IBM-Bluemix/bluemix-go/client"
	"github.com/IBM-Bluemix/bluemix-go/http"
	"github.com/IBM-Bluemix/bluemix-go/rest"
	"github.com/IBM-Bluemix/bluemix-go/session"
)

//IAMPAPAPI is the IAMpapv2 client ...
type IAMPAPAPI interface {
	IAMPolicy() IAMPolicy
	IAMService() IAMService
}

//ErrCodeAPICreation ...
const ErrCodeAPICreation = "APICreationError"

//IamPapService holds the client
type iampapService struct {
	*client.Client
}

//New ...
func New(sess *session.Session) (IAMPAPAPI, error) {
	config := sess.Config.Copy()
	err := config.ValidateConfigForService(bluemix.IAMPAPService)
	if err != nil {
		return nil, err
	}
	tokenRefreher, err := authentication.NewIAMAuthRepository(config, &rest.Client{
		DefaultHeader: gohttp.Header{
			"User-Agent": []string{http.UserAgent()},
		},
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

	if config.HTTPClient == nil {
		config.HTTPClient = http.NewHTTPClient(config)
	}

	if config.Endpoint == nil {
		ep, err := config.EndpointLocator.IAMPAPEndpoint()
		if err != nil {
			return nil, err
		}
		config.Endpoint = &ep
	}
	return &iampapService{
		Client: client.New(config, bluemix.IAMPAPService, tokenRefreher, nil),
	}, nil
}

//IAMPolicy API
func (a *iampapService) IAMPolicy() IAMPolicy {
	return newIAMPolicyAPI(a.Client)
}

//IAMService API
func (a *iampapService) IAMService() IAMService {
	return newIAMServiceAPI(a.Client)
}
