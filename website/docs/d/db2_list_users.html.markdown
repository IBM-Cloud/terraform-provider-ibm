---
subcategory: "Db2 SaaS"
layout: "ibm"
page_title: "IBM : ibm_db2_users"
description: |-
  Get information about IBM Db2 users
---

# ibm_db2_users

Retrieve information about users of an existing [IBM Db2 Instance](https://cloud.ibm.com/docs/Db2onCloud).

## Example Usage

```hcl
data "ibm_db2_users" "db2_users" {
	x_deployment_id = "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `x_deployment_id` - (Required, String) CRN of the instance this list of users relates to.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `count` - (Integer) The total number of resources.
* `resources` - (List) A list of user resource.
Nested schema for **resources**:
	* `all_clean` - (Boolean) Indicates if the user account has no issues.
	* `authentication` - (List) Authentication details for the user.
	Nested schema for **authentication**:
		* `method` - (String) The Authentication method used.
		* `policy_id` - (String) The Policy ID of the authentication.
	* `dv_role` - (String) User's DV role.
	* `email` - (String) Email address of the IBM Db2 user.
	* `formated_ibmid` - (String) Formatted IBM ID.
	* `iam` - (Boolean) Indicates if IAM is enabled or not.
	* `iamid` - (String) IAM ID for the IBM Db2 user.
	* `ibmid` - (String) IBM ID of the IBM Db2 user.
	* `id` - (String) Unique identifier for the IBM Db2 user.
	* `init_error_msg` - (String) Initial error message.
	* `locked` - (String) Account lock status for the IBM Db2 user.
	  * Constraints: Allowable values are: `yes`, `no`.
	* `metadata` - (Map) Metadata associated with the IBM Db2 user.
	* `name` - (String) The display name of the user.
	* `password` - (String) User's password.
	* `permitted_actions` - (List) List of allowed actions of the IBM Db2  user.
	* `role` - (String) Role assigned to the user.
	  * Constraints: Allowable values are: `bluadmin`, `bluuser`.