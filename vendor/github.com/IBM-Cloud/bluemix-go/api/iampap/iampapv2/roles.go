package iampapv2

import (
	"fmt"

	"github.com/IBM-Cloud/bluemix-go/client"
)

type CreateRoleRequest struct {
	Name        string   `json:"name"`
	ServiceName string   `json:"service_name"`
	AccountID   string   `json:"account_id"`
	DisplayName string   `json:"display_name"`
	Description string   `json:"description"`
	Actions     []string `json:"actions,omitempty"`
}
type UpdateRoleRequest struct {
	DisplayName string   `json:"display_name"`
	Description string   `json:"description"`
	Actions     []string `json:"actions,omitempty"`
}

type CreateRoleResponse struct {
	CreateRoleRequest
	ID               string `json:"id"`
	Crn              string `json:"crn"`
	CreatedAt        string `json:"created_at"`
	CreatedByID      string `json:"created_by_id"`
	LastModifiedAt   string `json:"last_modified_at"`
	LastModifiedByID string `json:"last_modified_by_id"`
}

type ListResponse struct {
	CustomRoles  []CreateRoleResponse `json:"custom_roles"`
	ServiceRoles []CreateRoleResponse `json:"service_roles"`
	SystemRoles  []CreateRoleResponse `json:"system_roles"`
}

type RoleRepository interface {
	Get(roleID string) (CreateRoleResponse, string, error)
	Create(request CreateRoleRequest) (CreateRoleResponse, error)
	Update(request UpdateRoleRequest, roleID, etag string) (CreateRoleResponse, error)
	Delete(roleID string) error
	List(accountID, serviceName string) (ListResponse, error)
}

type roleRepository struct {
	client *client.Client
}

func NewRoleRepository(c *client.Client) RoleRepository {
	return &roleRepository{
		client: c,
	}
}

func (r *roleRepository) Create(request CreateRoleRequest) (CreateRoleResponse, error) {
	res := CreateRoleResponse{}
	_, err := r.client.Post(fmt.Sprintf("/v2/roles"), &request, &res)
	if err != nil {
		return CreateRoleResponse{}, err
	}
	return res, nil
}

func (r *roleRepository) Get(roleID string) (CreateRoleResponse, string, error) {
	res := CreateRoleResponse{}
	response, err := r.client.Get(fmt.Sprintf("/v2/roles/%s", roleID), &res)
	if err != nil {
		return CreateRoleResponse{}, "", err
	}
	return res, response.Header.Get("Etag"), nil
}

func (r *roleRepository) Update(request UpdateRoleRequest, roleID, etag string) (CreateRoleResponse, error) {
	res := CreateRoleResponse{}
	header := make(map[string]string)

	header["IF-Match"] = etag
	_, err := r.client.Put(fmt.Sprintf("/v2/roles/%s", roleID), &request, &res, header)
	if err != nil {
		return CreateRoleResponse{}, err
	}
	return res, nil
}

//Delete Function
func (r *roleRepository) Delete(roleID string) error {
	_, err := r.client.Delete(fmt.Sprintf("/v2/roles/%s", roleID))
	return err
}

func (r *roleRepository) List(accountID, serviceName string) (ListResponse, error) {
	res := ListResponse{}
	var requestpath string
	if accountID == "" {
		requestpath = fmt.Sprintf("/v2/roles?service_name=%s", serviceName)
	} else {
		if serviceName == "" {
			requestpath = fmt.Sprintf("/v2/roles?account_id=%s", accountID)

		} else {
			requestpath = fmt.Sprintf("/v2/roles?service_name=%s?account_id=%s", serviceName, accountID)
		}
	}

	_, err := r.client.Get(requestpath, &res)
	if err != nil {
		return ListResponse{}, err
	}
	return res, nil
}
