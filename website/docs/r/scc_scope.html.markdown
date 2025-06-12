---
layout: "ibm"
page_title: "IBM : ibm_scc_scope"
description: |-
  Manages scc_scope.
subcategory: "Security and Compliance Center"
---

# ibm_scc_scope

Create, update, and delete scc_scopes with this resource.

## Example Usage

To create a scope targeting an account
```hcl
resource "ibm_scc_scope" "scc_account_scope" {
  description = "This scope allows a profile attachment to target an IBM account"
  environment = "ibm-cloud"
  instance_id = "b36c26e9-477a-43a1-9c50-19aff8e5d760"
  name        = "Sample account Scope"
  properties  = {
    scope_id = "8e042beeccee40748674442960b9eb34"
    scope_type = "account"
  }
}
```

To create a scope targeting an enterprise
```hcl
resource "ibm_scc_scope" "scc_enterprise_scope" {
  description = "This scope allows a profile attachment to target an IBM enterprise"
  environment = "ibm-cloud"
  instance_id = "b36c26e9-477a-43a1-9c50-19aff8e5d760"
  name        = "Sample enterprise Scope"
  properties  = {
    scope_id   = "6a204bd89f3c8348afd5c77c717a097a"
    scope_type = "enterprise"
  }
}
```

To create a scope targeting an account with an exclusion of a resource group
```hcl
resource "ibm_scc_scope" "scc_account_scope" {
  description = "This scope allows a profile attachment to target an IBM account"
  environment = "ibm-cloud"
  instance_id = "b36c26e9-477a-43a1-9c50-19aff8e5d760"
  name        = "Sample account Scope"
  properties  = {
    scope_id = "8e042beeccee40748674442960b9eb34"
    scope_type = "account"
  }
  exclusions {
		scope_id   = "ff6ce35b305abe1f768e3317628c0ba3"
		scope_type = "account.resource_group"
	}
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `description` - (Optional, String) The scope description.
  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-zA-Z0-9_,'\\s\\-\\.]*$/`.
* `environment` - (Required, Force New, String) The scope environment. This value details what cloud provider the scope targets.
  * Constraints: The maximum length is `128` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-zA-Z0-9_,'\\s\\-\\.]*$/`.
  * Acceptable values are:
    - `ibm-cloud`
* `instance_id` - (Required, Forces new resource, String) The ID of the Security and Compliance Center instance.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9A-Fa-f]{8}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{12}$/`.
* `name` - (Required, String) The scope name.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_,'\\s\\-\\.]*$/`.
* `properties` - (Required, Forces new resource, Map) The properties of the scope to target.

    Keys accepted in **properties**:
      * `scope_type` - (Required, String) The type of target the scope will cover
        * Constraints: Acceptable values are:
          * `account` - scope will target an IBM account
          * `account.resource_group` - scope will target a resource_group of the account which owns the Security and Compliance Center instance specified in `instance_id`
          * `enterprise.account_group` - targets an enterprise's account group
          * `enterprise` - targets an IBM enterprise
      * `scope_id` - (Required, String) The ID of the target defined in `scope_type`.
* `exclusions` - (Optional, List, Forces new resource) A list of scopes/targets to exclude from a scope.
  
  Nested schema for **exclusions**:
    * `scope_type` - (Required, String) The type of target to exclude from the scope
      * Constraints: Acceptable values are `account`, `account.resource_group`, or `enterprise.account_group`.
    * `scope_id` - (Required, String) The ID of the target defined in `scope_type`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the scc_scope.
* `account_id` - (String) The ID of the account associated with the scope.
  * Constraints: The maximum length is `32` characters. The minimum length is `32` characters. The value must match regular expression `/^[a-zA-Z0-9_\\-.]*$/`.
* `attachment_count` - (Float) The number of attachments tied to the scope.
* `created_by` - (String) The identifier of the account or service ID who created the scope.
  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.:,_\\s]*$/`.
* `created_on` - (String) The date when the scope was created.
* `scope_id` - (String) The ID of the scope.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
* `updated_by` - (String) The ID of the user or service ID who updated the scope.
  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.:,_\\s]*$/`.
* `updated_on` - (String) The date when the scope was updated.


## Import

You can import the `ibm_scc_scope` resource by using `id`.
The `id` property can be formed from `instance_id`, and `scope_id` in the following format:

```
<instance_id>/<scope_id>
```
* `instance_id`: A string in the format `acd7032c-15a3-484f-bf5b-67d41534d940`. The ID of the Security and Compliance Center instance.
* `scope_id`: A string. The ID of the scope being targeted.

# Syntax
```
$ terraform import ibm_scc_scope.scc_scope <instance_id>/<scope_id>
```
