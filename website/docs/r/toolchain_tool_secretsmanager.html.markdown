---
layout: "ibm"
page_title: "IBM : ibm_toolchain_tool_secretsmanager"
description: |-
  Manages toolchain_tool_secretsmanager.
subcategory: "Toolchain"
---

# ibm_toolchain_tool_secretsmanager

Provides a resource for toolchain_tool_secretsmanager. This allows toolchain_tool_secretsmanager to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_toolchain_tool_secretsmanager" "toolchain_tool_secretsmanager" {
  parameters {
		name = "name"
		region = "region"
		resource-group = "resource-group"
		instance-name = "instance-name"
		integration-status = "integration-status"
  }
  toolchain_id = "toolchain_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `name` - (Optional, String) Name of tool integration.
* `parameters` - (Optional, List) Tool integration parameters.
Nested scheme for **parameters**:
	* `instance_name` - (Required, String) The name of your Secrets Manager instance. You should choose an entry from the list provided based on the selected region and resource group. e.g: Secrets Manager-01.
	  * Constraints: The value must match regular expression `/\\S/`.
	* `integration_status` - (Optional, String)
	* `name` - (Required, String) Enter a name for this tool integration. This name is displayed on your toolchain.
	* `region` - (Required, String) Region.
	* `resource_group` - (Required, String) Resource group.
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain to bind integration to.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the toolchain_tool_secretsmanager.
* `crn` - (Required, String) 
* `get_integration_by_id_response_id` - (Required, String) 
* `href` - (Required, String) 
* `referent` - (Required, List) 
Nested scheme for **referent**:
	* `api_href` - (Optional, String)
	* `ui_href` - (Optional, String)
* `resource_group_id` - (Required, String) 
* `state` - (Required, String) 
  * Constraints: Allowable values are: `configured`, `configuring`, `misconfigured`, `unconfigured`.
* `toolchain_crn` - (Required, String) 
* `updated_at` - (Required, String) 

## Import

You can import the `ibm_toolchain_tool_secretsmanager` resource by using `id`.
The `id` property can be formed from `toolchain_id`, and `integration_id` in the following format:

```
<toolchain_id>/<integration_id>
```
* `toolchain_id`: A string. ID of the toolchain to bind integration to.
* `integration_id`: A string. ID of the tool integration to be deleted.

# Syntax
```
$ terraform import ibm_toolchain_tool_secretsmanager.toolchain_tool_secretsmanager <toolchain_id>/<integration_id>
```
