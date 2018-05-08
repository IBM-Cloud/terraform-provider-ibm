/**
 * Copyright 2016 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/**
 * AUTOMATICALLY GENERATED CODE - DO NOT MODIFY
 */

package datatypes

// no documentation yet
type BluePages_Container_EmployeeProfile struct {
	Entity

	// Employee address
	Address1 *string `json:"address1,omitempty" xmlrpc:"address1,omitempty"`

	// Employee address
	Address2 *string `json:"address2,omitempty" xmlrpc:"address2,omitempty"`

	// Country of employee's address
	AddressCountry *string `json:"addressCountry,omitempty" xmlrpc:"addressCountry,omitempty"`

	// Employee city
	City *string `json:"city,omitempty" xmlrpc:"city,omitempty"`

	// Employee department code
	DepartmentCode *string `json:"departmentCode,omitempty" xmlrpc:"departmentCode,omitempty"`

	// Employee department country code
	DepartmentCountry *string `json:"departmentCountry,omitempty" xmlrpc:"departmentCountry,omitempty"`

	// Employee division code
	DivisionCode *string `json:"divisionCode,omitempty" xmlrpc:"divisionCode,omitempty"`

	// Employee email address
	EmailAddress *string `json:"emailAddress,omitempty" xmlrpc:"emailAddress,omitempty"`

	// Employee first name
	FirstName *string `json:"firstName,omitempty" xmlrpc:"firstName,omitempty"`

	// Employee last name
	LastName *string `json:"lastName,omitempty" xmlrpc:"lastName,omitempty"`

	// Email of employee's manager
	ManagerEmailAddress *string `json:"managerEmailAddress,omitempty" xmlrpc:"managerEmailAddress,omitempty"`

	// Employee's manager's first name
	ManagerFirstName *string `json:"managerFirstName,omitempty" xmlrpc:"managerFirstName,omitempty"`

	// Employee's manager's last name
	ManagerLastName *string `json:"managerLastName,omitempty" xmlrpc:"managerLastName,omitempty"`

	// Employee' manager's identifier
	ManagerUid *string `json:"managerUid,omitempty" xmlrpc:"managerUid,omitempty"`

	// Employee phone number
	Phone *string `json:"phone,omitempty" xmlrpc:"phone,omitempty"`

	// Employee postal code
	PostalCode *string `json:"postalCode,omitempty" xmlrpc:"postalCode,omitempty"`

	// Employee state
	State *string `json:"state,omitempty" xmlrpc:"state,omitempty"`

	// Employee identifier
	Uid *string `json:"uid,omitempty" xmlrpc:"uid,omitempty"`
}

// no documentation yet
type BluePages_Search struct {
	Entity
}
