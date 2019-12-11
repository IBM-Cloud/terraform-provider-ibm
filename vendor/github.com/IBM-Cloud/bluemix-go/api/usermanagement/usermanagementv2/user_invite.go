package usermanagementv2

import (
	"fmt"
	"net/http"

	"github.com/IBM-Cloud/bluemix-go/client"
	"github.com/IBM-Cloud/bluemix-go/rest"
)

const (
	_UsersIDPath = "/v2/accounts/%s/users/%s"
	_UsersURL    = "/v2/accounts/%s/users"
)

// Users ...
type Users interface {
	GetUsers(ibmUniqueID string) (UsersList, error)
	GetUserProfile(ibmUniqueID string, userID string) (UserInfo, error)
	InviteUsers(ibmUniqueID string, users UserInvite) (UserInvite, error)
	UpdateUserProfile(ibmUniqueID string, userID string, user UserInfo) error
	RemoveUsers(ibmUniqueID string, userID string) error
}

type inviteUsersHandler struct {
	client *client.Client
}

// NewUsers
func NewUserInviteHandler(c *client.Client) Users {
	return &inviteUsersHandler{
		client: c,
	}
}

func (r *inviteUsersHandler) GetUsers(ibmUniqueID string) (UsersList, error) {
	result := UsersList{}
	URL := fmt.Sprintf(_UsersURL, ibmUniqueID)
	resp, err := r.client.Get(URL, &result)

	if resp.StatusCode == http.StatusNotFound {
		return UsersList{}, nil
	}

	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *inviteUsersHandler) GetUserProfile(ibmUniqueID string, userID string) (UserInfo, error) {
	user := UserInfo{}
	URL := fmt.Sprintf(_UsersIDPath, ibmUniqueID, userID)
	_, err := r.client.Get(URL, &user)
	if err != nil {
		return UserInfo{}, err
	}

	return user, nil
}

func (r *inviteUsersHandler) InviteUsers(ibmUniqueID string, users UserInvite) (UserInvite, error) {
	usersInvited := UserInvite{}
	URL := fmt.Sprintf(_UsersURL, ibmUniqueID)
	_, err := r.client.Post(URL, &users, &usersInvited)
	if err != nil {
		return UserInvite{}, err
	}

	return usersInvited, nil
}

func (r *inviteUsersHandler) UpdateUserProfile(ibmUniqueID string, userID string, user UserInfo) error {
	URL := fmt.Sprintf(_UsersIDPath, ibmUniqueID, userID)
	request := rest.PutRequest(*r.client.Config.Endpoint + URL)
	request = request.Body(&user)

	_, err := r.client.SendRequest(request, nil)
	if err != nil {
		return err
	}

	return nil
}

func (r *inviteUsersHandler) RemoveUsers(ibmUniqueID string, userID string) error {
	URL := fmt.Sprintf(_UsersIDPath, ibmUniqueID, userID)
	_, err := r.client.Delete(URL)
	return err
}
