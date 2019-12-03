package models

// User ...
type User struct {
	Email       string `json:"email"`
	AccountRole string `json:"account_role"`
}

// UserInfo contains user info
type UserInfo struct {
	ID             string `json:"id"`
	IamID          string `json:"iam_id"`
	Realm          string `json:"realm"`
	UserID         string `json:"user_id"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	State          string `json:"state"`
	Email          string `json:"email"`
	Phonenumber    string `json:"phonenumber"`
	Altphonenumber string `json:"altphonenumber"`
	Photo          string `json:"photo"`
	AccountID      string `json:"account_id"`
}

// UserInvite ...
type UserInvite struct {
	Users             []User   `json:"users"`
	IAMPolicy         []string `json:"iam_policy,omitempty"`
	AccessGroup       []string `json:"access_groups,omitempty"`
	OrganizationRoles []string `json:"organization_roles,omitempty"`
}

// UsersList to get list of users
type UsersList struct {
	TotalUsers int        `json:"total_results"`
	Limit      int        `json:"limit"`
	FistURL    string     `json:"fist_url"`
	Resources  []UserInfo `json:"resources"`
}
