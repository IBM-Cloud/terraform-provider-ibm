---
subcategory: "Identity & Access Management (IAM)"
layout: "ibm"
page_title: "IBM : iam_users"
description: |-
  Fetches all IAM users profile information.
---

# ibm_iam_users

Retrieve information about an IAM user profile on IBM Cloud as a read-only data source. For more information, about IAM users profile information, see [assigning access to account management services](https://cloud.ibm.com/docs/account?topic=account-account-services).


## Example usage

```terraform

	data "ibm_iam_users" "users_profiles"{
  
	} 

```

## Attribute reference

You can access the following attribute references after your data source is created.

- `id` - (String) The unique identifier user.
- `users`-  (String) List of all IAM users. Each user profile has following list of arguments.

  Nested scheme for `users`:
	- `altphonenumber`-  (String) The alternative phone number of the user.
	- `account_id`-  (String) The account ID of the user.
	- `email`-  (String) The email of the user.
	- `iam_id`-  (String) The ID of the IAM user.
	- `firstname`-  (String) The first name of the user.
	- `lastname`-  (String) The last name of the user.
	- `phonenumber`-  (String) The phone for the user.
	- `realm`-  (String) The realm of the user.
	- `state`-  (String) The state of the user.
	- `user_id`-  (String) The user ID used for log in.
  
