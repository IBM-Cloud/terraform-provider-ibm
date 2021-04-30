---
subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_users"
description: |-
  Fechtes all IAM users profile information.
---

# ibm\_iam_users

Import the details of all IAM (Identity and Access Management) users profile on IBM Cloud as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl

	data "ibm_iam_users" "users_profiles"{
  
	} 

```

## Attribute Reference

In addition to all arguments above, the following attributes are exported:
* `id` - An alphanumeric value identifying the user.
* `users` - List of all IAM users. Each user profile has following list of arguments	
  * `iam_id` - An alphanumeric value identifying the user's IAM ID.
  * `realm` - The realm of the user. 
  * `user_id` - The user ID used for login.
  * `firstname` - The first name of the user.
  * `lastname` -  The last name of the user.
  * `state` - The state of the user. 
  * `email` - The email of the user.
  * `phonenumber` - The phone for the user.
  * `altphonenumber` - The alternative phone number of the user.
  * `account_id` - An alphanumeric value identifying the account ID.


  