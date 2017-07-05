package accountv1

import (
	gohttp "net/http"

	bluemix "github.com/IBM-Bluemix/bluemix-go"
	"github.com/IBM-Bluemix/bluemix-go/api/account/accountv2"
	"github.com/IBM-Bluemix/bluemix-go/authentication"
	"github.com/IBM-Bluemix/bluemix-go/client"
	"github.com/IBM-Bluemix/bluemix-go/http"
	"github.com/IBM-Bluemix/bluemix-go/rest"
	"github.com/IBM-Bluemix/bluemix-go/session"
)

//AccountServiceAPI is the accountv2 client ...
type AccountServiceAPI interface {
	Accounts() Accounts
}

//ErrCodeNoAccountExists ...
const ErrCodeNoAccountExists = "NoAccountExists"

//CfService holds the client
type accountService struct {
	*client.Client
}

//New ...
func New(sess *session.Session) (AccountServiceAPI, error) {
	config := sess.Config.Copy()
	err := config.ValidateConfigForService(bluemix.AccountServicev1)
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
		ep, err := config.EndpointLocator.AccountManagementEndpoint()
		if err != nil {
			return nil, err
		}
		config.Endpoint = &ep
	}

	return &accountService{
		Client: client.New(config, bluemix.AccountServicev1, tokenRefreher, accountv2.Paginate),
	}, nil
}

//Accounts API
func (a *accountService) Accounts() Accounts {
	return newAccountAPI(a.Client)
}
