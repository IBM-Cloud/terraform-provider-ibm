package accountv1

import (
	"fmt"

	"github.com/IBM-Bluemix/bluemix-go/bmxerror"
	"github.com/IBM-Bluemix/bluemix-go/client"
)

type AccountUser struct {
	UserId      string `json:"userId"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	State       string `json:"state"`
	IbmUniqueId string `json:"ibmUniqueId"`
	Email       string `json:"email"`
	Phonenumber string `json:"phonenumber"`
	CreatedOn   string `json:"createdOn"`
	VerifiedOn  string `json:"verifiedOn"`
	Id          string `json:"id"`
	UaaGuid     string `json:"uaaGuid"`
	AccountId   string `json:"accountId"`
	Role        string `json:"role"`
	InvitedOn   string `json:"invitedOn"`
	Photo       string `json:"photo"`
}

//Accounts ...
type Accounts interface {
	GetAccountUsers(accountGuid string) ([]AccountUser, error)
	InviteAccountUser(accountGuid string, userEmail string) (AccountInviteResponse, error)
	DeleteAccountUser(accountGuid string, userGuid string) error
}

type account struct {
	client *client.Client
}

type AccountUserResource struct {
	Metadata AccountUserMetadata
	Entity   AccountUserEntity
}

type Metadata struct {
	Guid       string   `json:"guid"`
	Url        string   `json:"url"`
	CreatedAt  string   `json:"created_at"`
	UpdatedAt  string   `json:"updated_at"`
	VerifiedAt string   `json:"verified_at"`
	Identity   Identity `json:"identity"`
}

type AccountUserEntity struct {
	AccountId   string `json:"account_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	State       string `json:"state"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phonenumber"`
	Role        string `json:"role"`
	Photo       string `json:"photo"`
}

type AccountUserMetadata struct {
	Metadata
	VerifiedAt string `json:"verified_at"`
}

type Identity struct {
	Id         string `json:"id"`
	UserName   string `json:"username"`
	Realmid    string `json:"realmid"`
	Identifier string `json:"identifier"`
}

// Account Invites ...
type AccountInviteResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	State string `json:"state"`
}

func (resource AccountUserResource) ToModel() AccountUser {
	m := resource.Metadata
	e := resource.Entity

	return AccountUser{
		UserId:      m.Identity.UserName,
		CreatedOn:   m.CreatedAt,
		VerifiedOn:  m.VerifiedAt,
		FirstName:   e.FirstName,
		LastName:    e.LastName,
		IbmUniqueId: m.Identity.Id,
		State:       e.State,
		Email:       e.Email,
		Phonenumber: e.PhoneNumber,
		Id:          m.Guid,
		AccountId:   e.AccountId,
		Role:        e.Role,
		Photo:       e.Photo,
	}
}

func newAccountAPI(c *client.Client) Accounts {
	return &account{
		client: c,
	}
}

//GetAccountUser ...
func (a *account) GetAccountUsers(accountGuid string) ([]AccountUser, error) {
	var users []AccountUser

	resp, err := a.client.GetPaginated(fmt.Sprintf("/v1/accounts/%s/users", accountGuid),
		AccountUserResource{},
		func(resource interface{}) bool {
			if accountUser, ok := resource.(AccountUserResource); ok {
				users = append(users, accountUser.ToModel())
				return true
			}
			return false
		})

	if resp.StatusCode == 404 {
		return []AccountUser{}, bmxerror.New(ErrCodeNoAccountExists,
			fmt.Sprintf("No Account exists with account id:%q", accountGuid))
	}

	return users, err
}

func (a *account) InviteAccountUser(accountGuid string, userEmail string) (AccountInviteResponse, error) {
	type userEntity struct {
		Email       string `json:"email"`
		AccountRole string `json:"account_role"`
	}

	payload := struct {
		Users []userEntity `json:"users"`
	}{
		Users: []userEntity{
			{
				Email:       userEmail,
				AccountRole: "MEMBER",
			},
		},
	}

	resp := AccountInviteResponse{}

	_, err := a.client.Post(fmt.Sprintf("/v1/accounts/%s/users", accountGuid), payload, &resp)
	return resp, err
}

func (a *account) DeleteAccountUser(accountGuid string, userGuid string) error {
	_, err := a.client.Delete(fmt.Sprintf("/v1/accounts/%s/users/%s", accountGuid, userGuid))

	return err
}
