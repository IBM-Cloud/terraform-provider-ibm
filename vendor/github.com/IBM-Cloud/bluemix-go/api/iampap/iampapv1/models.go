package iampapv1

import "github.com/IBM-Cloud/bluemix-go/models"

// Policy is the model of IAM PAP policy
type Policy struct {
	ID               string     `json:"id,omitempty"`
	Type             string     `json:"type"`
	Subjects         []Subject  `json:"subjects"`
	Roles            []Role     `json:"roles"`
	Resources        []Resource `json:"resources"`
	Href             string     `json:"href,omitempty"`
	CreatedAt        string     `json:"created_at,omitempty"`
	CreatedByID      string     `json:"created_by_id,omitempty"`
	LastModifiedAt   string     `json:"last_modified_at,omitempty"`
	LastModifiedByID string     `json:"last_modified_by_id,omitempty"`
	Version          string     `json:"-"`
}

// Role is the role model used by policy
type Role struct {
	RoleID      string `json:"role_id"`
	Name        string `json:"display_name,omitempty"`
	Description string `json:"description,omitempty"`
}

func fromModel(role models.PolicyRole) Role {
	return Role{
		RoleID: role.ID.String(),
		// When create/update, "name" and "description" are not allowed
		// Name:        role.Name,
		// Description: role.Description,
	}
}

// ConvertRoleModels will transform role models returned from "/v1/roles" to the model used by policy
func ConvertRoleModels(roles []models.PolicyRole) []Role {
	results := make([]Role, len(roles))
	for i, r := range roles {
		results[i] = fromModel(r)
	}
	return results
}

// Subject is the target to which is assigned policy
type Subject struct {
	Attributes []Attribute `json:"attributes"`
}

// GetAttribute returns an attribute of policy subject
func (s *Subject) GetAttribute(name string) string {
	for _, a := range s.Attributes {
		if a.Name == name {
			return a.Value
		}
	}
	return ""
}

// SetAttribute sets value of an attribute of policy subject
func (s *Subject) SetAttribute(name string, value string) {
	for _, a := range s.Attributes {
		if a.Name == name {
			a.Value = value
			return
		}
	}
	s.Attributes = append(s.Attributes, Attribute{
		Name:  name,
		Value: value,
	})
}

// AccessGroupID returns access group ID attribute of policy subject if exists
func (s *Subject) AccessGroupID() string {
	return s.GetAttribute("access_group_id")
}

// AccountID returns account ID attribute of policy subject if exists
func (s *Subject) AccountID() string {
	return s.GetAttribute("accountId")
}

// IAMID returns IAM ID attribute of policy subject if exists
func (s *Subject) IAMID() string {
	return s.GetAttribute("iam_id")
}

// ServiceName returns service name attribute of policy subject if exists
func (s *Subject) ServiceName() string {
	return s.GetAttribute("serviceName")
}

// ServiceInstance returns service instance attribute of policy subject if exists
func (s *Subject) ServiceInstance() string {
	return s.GetAttribute("serviceInstance")
}

// SetAccessGroupID sets value of access group ID attribute of policy subject
func (s *Subject) SetAccessGroupID(value string) {
	s.SetAttribute("access_group_id", value)
}

// SetAccountID sets value of account ID attribute of policy subject
func (s *Subject) SetAccountID(value string) {
	s.SetAttribute("accountId", value)
}

// SetIAMID sets value of IAM ID attribute of policy subject
func (s *Subject) SetIAMID(value string) {
	s.SetAttribute("iam_id", value)
}

// SetServiceName sets value of service name attribute of policy subject
func (s *Subject) SetServiceName(value string) {
	s.SetAttribute("serviceName", value)
}

// SetServiceInstance sets value of service instance attribute of policy subject
func (s *Subject) SetServiceInstance(value string) {
	s.SetAttribute("serviceInstance", value)
}

// Resource is the object controlled by the policy
type Resource struct {
	Attributes []Attribute `json:"attributes"`
}

// GetAttribute returns an attribute of policy resource
func (r *Resource) GetAttribute(name string) string {
	for _, a := range r.Attributes {
		if a.Name == name {
			return a.Value
		}
	}
	return ""
}

// SetAttribute sets value of an attribute of policy resource
func (r *Resource) SetAttribute(name string, value string) {
	for _, a := range r.Attributes {
		if a.Name == name {
			a.Value = value
			return
		}
	}
	r.Attributes = append(r.Attributes, Attribute{
		Name:  name,
		Value: value,
	})
}

// AccessGroupID returns access group ID attribute of policy resource if exists
func (r *Resource) AccessGroupID() string {
	return r.GetAttribute("accessGroupId")
}

// AccountID returns account ID attribute of policy resource if exists
func (r *Resource) AccountID() string {
	return r.GetAttribute("accountId")
}

// OrganizationID returns organization ID attribute of policy resource if exists
func (r *Resource) OrganizationID() string {
	return r.GetAttribute("organizationId")
}

// Region returns region attribute of policy resource if exists
func (r *Resource) Region() string {
	return r.GetAttribute("region")
}

// Resource returns resource attribute of policy resource if exists
func (r *Resource) Resource() string {
	return r.GetAttribute("resource")
}

// ResourceType returns resource type attribute of policy resource if exists
func (r *Resource) ResourceType() string {
	return r.GetAttribute("resourceType")
}

// ResourceGroupID returns resource group ID attribute of policy resource if exists
func (r *Resource) ResourceGroupID() string {
	return r.GetAttribute("resourceGroupId")
}

// ServiceName returns service name attribute of policy resource if exists
func (r *Resource) ServiceName() string {
	return r.GetAttribute("serviceName")
}

// ServiceInstance returns service instance attribute of policy resource if exists
func (r *Resource) ServiceInstance() string {
	return r.GetAttribute("serviceInstance")
}

// SpaceID returns space ID attribute of policy resource if exists
func (r *Resource) SpaceID() string {
	return r.GetAttribute("spaceId")
}

// ServiceType returns service type attribute of policy resource if exists
func (r *Resource) ServiceType() string {
	return r.GetAttribute("serviceType")
}

// SetAccessGroupID sets value of access group ID attribute of policy resource
func (r *Resource) SetAccessGroupID(value string) {
	r.SetAttribute("accessGroupId", value)
}

// SetAccountID sets value of account ID attribute of policy resource
func (r *Resource) SetAccountID(value string) {
	r.SetAttribute("accountId", value)
}

// SetOrganizationID sets value of organization ID attribute of policy resource
func (r *Resource) SetOrganizationID(value string) {
	r.SetAttribute("organizationId", value)
}

// SetRegion sets value of region attribute of policy resource
func (r *Resource) SetRegion(value string) {
	r.SetAttribute("region", value)
}

// SetResource sets value of resource attribute of policy resource
func (r *Resource) SetResource(value string) {
	r.SetAttribute("resource", value)
}

// SetResourceType sets value of resource type attribute of policy resource
func (r *Resource) SetResourceType(value string) {
	r.SetAttribute("resourceType", value)
}

// SetResourceGroupID sets value of resource group ID attribute of policy resource
func (r *Resource) SetResourceGroupID(value string) {
	r.SetAttribute("resourceGroupId", value)
}

// SetServiceName sets value of service name attribute of policy resource
func (r *Resource) SetServiceName(value string) {
	r.SetAttribute("serviceName", value)
}

// SetServiceInstance sets value of service instance attribute of policy resource
func (r *Resource) SetServiceInstance(value string) {
	r.SetAttribute("serviceInstance", value)
}

// SetSpaceID sets value of space ID attribute of policy resource
func (r *Resource) SetSpaceID(value string) {
	r.SetAttribute("spaceID", value)
}

// SetServiceType sets value of service type attribute of policy resource
func (r *Resource) SetServiceType(value string) {
	r.SetAttribute("serviceType", value)
}

// Attribute is part of policy subject and resource
type Attribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
