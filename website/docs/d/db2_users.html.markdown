---
layout: "ibm"
subcategory: "Db2 SaaS"
page_title: "IBM : ibm_db2_users"
description: |-
  Get information about db2_users
---

# ibm_db2_users

Provides a read-only data source to retrieve information about db2_users. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_db2_users" "db2_users" {
	x_deployment_id = "crn:v1:staging:public:dashdb-for-transactions:us-south:a/e7e3e87b512f474381c0684a5ecbba03:69db420f-33d5-4953-8bd8-1950abd356f6::"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `x_deployment_id` - (Required, String) CRN deployment id.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^crn(:[A-Za-z0-9\\-\\.]*){9}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the db2_users.
* `count` - (Integer) The total number of resources.
* `resources` - (List) A list of user resource.
Nested schema for **resources**:
	* `all_clean` - (Boolean) Indicates if the user account has no issues.
	* `authentication` - (List) Authentication details for the user.
	Nested schema for **authentication**:
		* `method` - (String) Authentication method.
		* `policy_id` - (String) Policy ID of authentication.
	* `dv_role` - (String) User's DV role.
	* `email` - (String) Email address of the user.
	* `formated_ibmid` - (String) Formatted IBM ID.
	* `iam` - (Boolean) Indicates if IAM is enabled or not.
	* `iamid` - (String) IAM ID for the user.
	* `ibmid` - (String) IBM ID of the user.
	* `id` - (String) Unique identifier for the user.
	* `init_error_msg` - (String) Initial error message.
	* `locked` - (String) Account lock status for the user.
	  * Constraints: Allowable values are: `yes`, `no`.
	* `metadata` - (Map) Metadata associated with the user.
	* `name` - (String) The display name of the user.
	* `password` - (String) User's password.
	* `permitted_actions` - (List) List of allowed actions of the user.
	* `role` - (String) Role assigned to the user.
	  * Constraints: Allowable values are: `bluadmin`, `bluuser`.

