---
layout: "ibm"
page_title: "IBM : ibm_toolchain_tool_jenkins"
description: |-
  Manages toolchain_tool_jenkins.
subcategory: "Toolchain"
---

# ibm_toolchain_tool_jenkins

Provides a resource for toolchain_tool_jenkins. This allows toolchain_tool_jenkins to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_toolchain_tool_jenkins" "toolchain_tool_jenkins" {
  parameters {
		name = "name"
		dashboard_url = "dashboard_url"
		webhook_url = "webhook_url"
		api_user_name = "api_user_name"
		api_token = "api_token"
  }
  toolchain_id = "toolchain_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `name` - (Optional, String) Name of tool integration.
* `parameters` - (Optional, List) Tool integration parameters.
Nested scheme for **parameters**:
	* `api_token` - (Optional, String) Type the API token to use for Jenkins REST API calls so that DevOps Insights can collect data from Jenkins. You can find the API token on the configuration page of your Jenkins instance.
	* `api_user_name` - (Optional, String) Type the user name to use with the Jenkins server's API token, which is required so that DevOps Insights can collect data from Jenkins. You can find your API user name on the configuration page of your Jenkins instance.
	* `dashboard_url` - (Required, String) Type the URL of the Jenkins server that you want to open when you click the Jenkins card in your toolchain.
	* `name` - (Required, String) Type a name for this tool integration, for example: my-jenkins. This name displays on your toolchain.
	* `webhook_url` - (Optional, String) Use this webhook in your Jenkins jobs to send notifications to other tools in your toolchain. For details, see the Configuring Jenkins instructions.
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain to bind integration to.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the toolchain_tool_jenkins.
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

You can import the `ibm_toolchain_tool_jenkins` resource by using `id`.
The `id` property can be formed from `toolchain_id`, and `integration_id` in the following format:

```
<toolchain_id>/<integration_id>
```
* `toolchain_id`: A string. ID of the toolchain to bind integration to.
* `integration_id`: A string. ID of the tool integration to be deleted.

# Syntax
```
$ terraform import ibm_toolchain_tool_jenkins.toolchain_tool_jenkins <toolchain_id>/<integration_id>
```
