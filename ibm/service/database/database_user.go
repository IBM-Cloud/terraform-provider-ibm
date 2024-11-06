// Copyright IBM Corp. 2017, 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/cloud-databases-go-sdk/clouddatabasesv5"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	passwordSpecialChars      = "~!@#$%^&*()=+[]{}|;:,.<>/?_-"
	redisRBACRoleRegexPattern = `[+-]@(?P<category>[a-z]+)`
)

type DatabaseUser struct {
	Username string
	Password string
	Role     *string
	Type     string
}

type databaseUserValidationError struct {
	user *DatabaseUser
	errs []error
}

func (e *databaseUserValidationError) Error() string {
	if len(e.errs) == 0 {
		return ""
	}

	var b []byte
	for i, err := range e.errs {
		if i > 0 {
			b = append(b, '\n')
		}
		b = append(b, err.Error()...)
	}

	return fmt.Sprintf("database user (%s) validation error:\n%s", e.user.Username, string(b))
}

func (e *databaseUserValidationError) Unwrap() error {
	if e == nil || len(e.errs) == 0 {
		return nil
	}

	// only return the first
	return e.errs[0]
}

type userChange struct {
	Old, New *DatabaseUser
}

func redisRBACAllowedRoles() []string {
	return []string{"all", "admin", "read", "write"}
}

func opsManagerRoles() []string {
	return []string{"group_read_only", "group_data_access_admin"}
}

func validateUsersDiff(_ context.Context, diff *schema.ResourceDiff, meta interface{}) (err error) {
	service := diff.Get("service").(string)

	var versionStr string
	var version int

	if _version, ok := diff.GetOk("version"); ok {
		versionStr = _version.(string)
	}

	if versionStr == "" {
		// Latest Version
		version = 0
	} else {
		_v, err := strconv.ParseFloat(versionStr, 64)

		if err != nil {
			return fmt.Errorf("invalid version: %s", versionStr)
		}

		version = int(_v)
	}

	oldUsers, newUsers := diff.GetChange("users")
	userChanges := expandUserChanges(oldUsers.(*schema.Set).List(), newUsers.(*schema.Set).List())

	for _, change := range userChanges {
		if change.isDelete() {
			continue
		}

		if change.isCreate() || change.isUpdate() {
			err = change.New.ValidatePassword()

			if err != nil {
				return err
			}

			// TODO: Use Capability API
			// RBAC roles supported for Redis 6.0 and above
			if (service == "databases-for-redis") && !(version > 0 && version < 6) {
				err = change.New.ValidateRBACRole()
			} else if service == "databases-for-mongodb" && change.New.Type == "ops_manager" {
				err = change.New.ValidateOpsManagerRole()
			} else {
				if change.New.Role != nil {
					if *change.New.Role != "" {
						err = errors.New("role is not supported for this deployment or user type")
						err = &databaseUserValidationError{user: change.New, errs: []error{err}}
					}
				}
			}

			if err != nil {
				return err
			}
		}
	}

	return
}

func expandUsers(_users []interface{}) []*DatabaseUser {
	if len(_users) == 0 {
		return nil
	}

	users := make([]*DatabaseUser, 0, len(_users))

	for _, userRaw := range _users {
		if tfUser, ok := userRaw.(map[string]interface{}); ok {

			user := DatabaseUser{
				Username: tfUser["name"].(string),
				Password: tfUser["password"].(string),
				Type:     tfUser["type"].(string),
			}

			// NOTE: cannot differentiate nil vs empty string
			// https://github.com/hashicorp/terraform-plugin-sdk/issues/741
			if role, ok := tfUser["role"].(string); ok {
				if tfUser["role"] != "" {
					user.Role = &role
				}
			}

			users = append(users, &user)
		}
	}

	return users
}

func expandUserChanges(_oldUsers []interface{}, _newUsers []interface{}) (userChanges []*userChange) {
	oldUsers := expandUsers(_oldUsers)
	newUsers := expandUsers(_newUsers)

	userChangeMap := make(map[string]*userChange)

	for _, user := range oldUsers {
		userChangeMap[user.ID()] = &userChange{Old: user}
	}

	for _, user := range newUsers {
		if _, ok := userChangeMap[user.ID()]; !ok {
			userChangeMap[user.ID()] = &userChange{}
		}
		userChangeMap[user.ID()].New = user
	}

	userChanges = make([]*userChange, 0, len(userChangeMap))

	for _, change := range userChangeMap {
		userChanges = append(userChanges, change)
	}

	return userChanges
}

func (c *userChange) isDelete() bool {
	return c.Old != nil && c.New == nil
}

func (c *userChange) isCreate() bool {
	return c.Old == nil && c.New != nil
}

func (c *userChange) isUpdate() bool {
	return c.New != nil &&
		c.Old != nil &&
		((c.Old.Password != c.New.Password) ||
			(c.Old.Role != c.New.Role))
}

func (u *DatabaseUser) ID() (id string) {
	return fmt.Sprintf("%s-%s", u.Type, u.Username)
}

func (u *DatabaseUser) Create(instanceID string, d *schema.ResourceData, meta interface{}) (err error) {
	cloudDatabasesClient, err := meta.(conns.ClientSession).CloudDatabasesV5()
	if err != nil {
		return fmt.Errorf("[ERROR] Error getting database client settings: %w", err)
	}

	//Attempt to create user
	userEntry := &clouddatabasesv5.User{
		Username: core.StringPtr(u.Username),
		Password: core.StringPtr(u.Password),
	}

	// User Role only for ops_manager user type and Redis 6.0 and above
	if u.Role != nil {
		userEntry.Role = u.Role
	}

	createDatabaseUserOptions := &clouddatabasesv5.CreateDatabaseUserOptions{
		ID:       &instanceID,
		UserType: core.StringPtr(u.Type),
		User:     userEntry,
	}

	createDatabaseUserResponse, response, err := cloudDatabasesClient.CreateDatabaseUser(createDatabaseUserOptions)
	if err != nil {
		return fmt.Errorf("[ERROR] CreateDatabaseUser (%s) failed %w\n%s", *userEntry.Username, err, response)
	}

	taskID := *createDatabaseUserResponse.Task.ID
	_, err = waitForDatabaseTaskComplete(taskID, d, meta, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf(
			"[ERROR] Error waiting for database (%s) user (%s) create task to complete: %w", instanceID, *userEntry.Username, err)
	}

	return nil
}

func (u *DatabaseUser) Update(instanceID string, d *schema.ResourceData, meta interface{}) (err error) {
	cloudDatabasesClient, err := meta.(conns.ClientSession).CloudDatabasesV5()
	if err != nil {
		return fmt.Errorf("[ERROR] Error getting database client settings: %s", err)
	}

	// Attempt to update user password
	user := &clouddatabasesv5.UserUpdate{
		Password: core.StringPtr(u.Password),
	}

	if u.Role != nil {
		user.Role = u.Role
	}

	updateUserOptions := &clouddatabasesv5.UpdateUserOptions{
		ID:       &instanceID,
		UserType: core.StringPtr(u.Type),
		Username: core.StringPtr(u.Username),
		User:     user,
	}

	updateUserResponse, response, err := cloudDatabasesClient.UpdateUser(updateUserOptions)

	// user was found but an error occurs while triggering task
	if err != nil || (response.StatusCode < 200 || response.StatusCode >= 300) {
		return fmt.Errorf("[ERROR] UpdateUser (%s) failed %w\n%s", *updateUserOptions.Username, err, response)
	}

	taskID := *updateUserResponse.Task.ID
	_, err = waitForDatabaseTaskComplete(taskID, d, meta, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf(
			"[ERROR] Error waiting for database (%s) user (%s) create task to complete: %w", instanceID, *updateUserOptions.Username, err)
	}

	return nil
}

func (u *DatabaseUser) Delete(instanceID string, d *schema.ResourceData, meta interface{}) (err error) {
	cloudDatabasesClient, err := meta.(conns.ClientSession).CloudDatabasesV5()
	if err != nil {
		return fmt.Errorf("[ERROR] Error getting database client settings: %s", err)
	}

	deleteDatabaseUserOptions := &clouddatabasesv5.DeleteDatabaseUserOptions{
		ID:       &instanceID,
		UserType: core.StringPtr(u.Type),
		Username: core.StringPtr(u.Username),
	}

	deleteDatabaseUserResponse, response, err := cloudDatabasesClient.DeleteDatabaseUser(deleteDatabaseUserOptions)

	if err != nil {
		return fmt.Errorf(
			"[ERROR] DeleteDatabaseUser (%s) failed %s\n%s", *deleteDatabaseUserOptions.Username, err, response)

	}

	taskID := *deleteDatabaseUserResponse.Task.ID
	_, err = waitForDatabaseTaskComplete(taskID, d, meta, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf(
			"[ERROR] Error waiting for database (%s) user (%s) delete task to complete: %s", instanceID, *deleteDatabaseUserOptions.Username, err)
	}

	return nil
}

func (u *DatabaseUser) isUpdatable() bool {
	return u.Type != "ops_manager"
}

func (u *DatabaseUser) ValidatePassword() (err error) {
	var errs []error

	// Format for regexp
	var specialCharPattern string
	var bs strings.Builder
	for i, c := range strings.Split(passwordSpecialChars, "") {
		if i > 0 {
			bs.WriteByte('|')
		}
		bs.WriteString(regexp.QuoteMeta(c))
	}

	specialCharPattern = bs.String()

	var allowedCharacters = regexp.MustCompile(fmt.Sprintf("^(?:[a-zA-Z0-9]|%s)+$", specialCharPattern))
	var beginWithSpecialChar = regexp.MustCompile(fmt.Sprintf("^(?:%s)", specialCharPattern))
	var containsLower = regexp.MustCompile("[a-z]")
	var containsUpper = regexp.MustCompile("[A-Z]")
	var containsNumber = regexp.MustCompile("[0-9]")
	var containsSpecialChar = regexp.MustCompile(fmt.Sprintf("(?:%s)", specialCharPattern))

	if u.Type == "ops_manager" && !containsSpecialChar.MatchString(u.Password) {
		errs = append(errs, fmt.Errorf(
			"password must contain at least one special character (%s)", passwordSpecialChars))
	}

	if u.Type == "database" && beginWithSpecialChar.MatchString(u.Password) {
		errs = append(errs, fmt.Errorf(
			"password must not begin with a special character (%s)", passwordSpecialChars))
	}

	if !containsLower.MatchString(u.Password) {
		errs = append(errs, errors.New("password must contain at least one lower case letter"))
	}

	if !containsUpper.MatchString(u.Password) {
		errs = append(errs, errors.New("password must contain at least one upper case letter"))
	}

	if !containsNumber.MatchString(u.Password) {
		errs = append(errs, errors.New("password must contain at least one number"))
	}

	if !allowedCharacters.MatchString(u.Password) {
		errs = append(errs, errors.New("password must not contain invalid characters"))
	}

	if len(errs) == 0 {
		return
	}

	return &databaseUserValidationError{user: u, errs: errs}
}

func (u *DatabaseUser) ValidateRBACRole() (err error) {
	var errs []error

	if u.Role == nil || *u.Role == "" {
		return
	}

	if u.Type != "database" {
		errs = append(errs, errors.New("role is only allowed for the database user"))
		return &databaseUserValidationError{user: u, errs: errs}
	}

	redisRBACCategoryRegex := regexp.MustCompile(redisRBACRoleRegexPattern)
	redisRBACRoleRegex := regexp.MustCompile(fmt.Sprintf(`^(%s\s?)+$`, redisRBACRoleRegexPattern))

	if !redisRBACRoleRegex.MatchString(*u.Role) {
		errs = append(errs, errors.New("role must be in the format +@category or -@category"))
	}

	matches := redisRBACCategoryRegex.FindAllStringSubmatch(*u.Role, -1)

	for _, match := range matches {
		valid := false
		role := match[1]
		for _, allowed := range redisRBACAllowedRoles() {
			if role == allowed {
				valid = true
				break
			}
		}

		if !valid {
			errs = append(errs, fmt.Errorf("role must contain only allowed categories: %s", strings.Join(redisRBACAllowedRoles()[:], ",")))
			break
		}
	}

	if len(errs) == 0 {
		return
	}

	return &databaseUserValidationError{user: u, errs: errs}
}

func (u *DatabaseUser) ValidateOpsManagerRole() (err error) {
	if u.Role == nil {
		return
	}

	if u.Type != "ops_manager" {
		return
	}

	if *u.Role == "" {
		return
	}

	for _, str := range opsManagerRoles() {
		if *u.Role == str {
			return
		}
	}

	err = fmt.Errorf("role must be a valid ops_manager role: %s", strings.Join(opsManagerRoles()[:], ","))

	return &databaseUserValidationError{user: u, errs: []error{err}}
}

func DatabaseUserPasswordValidator(userType string) schema.SchemaValidateFunc {
	return func(i interface{}, k string) (warnings []string, errors []error) {
		user := &DatabaseUser{Username: "admin", Type: userType, Password: i.(string)}
		err := user.ValidatePassword()
		if err != nil {
			errors = append(errors, err)
		}
		return
	}
}
