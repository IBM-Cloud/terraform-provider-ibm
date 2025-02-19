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

```hcl
resource "ibm_scc_scope" "scc_account_scope" {
  description = "This scope allows a profile attachment to target an IBM account"
  environment = "ibm-cloud"
  instance_id = "b36c26e9-477a-43a1-9c50-19aff8e5d760"
  name        = "Sample account Scope"
  properties {
    account_id = "8e042beeccee40748674442960b9eb34"
  }
}

resource "ibm_scc_scope" "scc_enterprise_scope" {
  description = "This scope allows a profile attachment to target an IBM enterprise"
  environment = "ibm-cloud"
  instance_id = "b36c26e9-477a-43a1-9c50-19aff8e5d760"
  name        = "Sample account Scope"
  properties {
    enterprise_id = "6a204bd89f3c8348afd5c77c717a097a"
  }
}

resource "ibm_scc_scope" "scc_ibm_facts_scope" {
  description = "This scope allows a profile attachment to target a facts provider instance"
  environment = "ibm-cloud"
  instance_id = "b36c26e9-477a-43a1-9c50-19aff8e5d760"
  name        = "Sample facts Scope"
  properties {
    ibm_facts_api_instance_id = ibm_scc_provider_type_instance.facts_provider_type_instance.provider_type_instance_id
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `description` - (Optional, String) The scope description.
  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-zA-Z0-9_,'\\s\\-\\.]*$/`.
* `environment` - (Optional, String) The scope environment. This value details what cloud provider the scope targets.
  * Constraints: The maximum length is `128` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-zA-Z0-9_,'\\s\\-\\.]*$/`.
* `instance_id` - (Required, Forces new resource, String) The ID of the Security and Compliance Center instance.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9A-Fa-f]{8}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{12}$/`.
* `name` - (Optional, String) The scope name.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_,'\\s\\-\\.]*$/`.
* `properties` - (Optional, List) The span for the scope to target.
Nested schema for **properties**:
    * `account_id` - (Optional, ForceNew, String) The ID of the IBM account ID.
    * `enterprise_id` - (Optional, ForceNew, String) The ID of the IBM enterprise ID.
    * `resource_group_id` - (Optional, ForceNew, String) The ID of the IBM resource group tied to an account
    * `account_group_id` - (Optional, ForceNew, String) The ID of an account group tied to an enterprise
    * `ibm_facts_api_instance_id` - (Optional, ForceNew, String) The ID of ibm_cloud_facts_api provider type instance
      * Constraints: Using `ibm_facts_api_instance_id` is restricted to certain users. 

* `exclusions` - (Optional, List) A list of scopes/targets to exclude from a scope.
Nested schema for **exclusions**:
    * `account_id` - (Optional, String) The account ID to exclude.
    * `resource_group_id` - (Optional, String) The ID of the IBM resource group in an account to exclude
    * `account_group_id` - (Optional, String) The ID of an account group in an enterprise.

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
