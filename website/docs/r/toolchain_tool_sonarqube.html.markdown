---
layout: "ibm"
page_title: "IBM : ibm_toolchain_tool_sonarqube"
description: |-
  Manages toolchain_tool_sonarqube.
subcategory: "Toolchain"
---

# ibm_toolchain_tool_sonarqube

Provides a resource for toolchain_tool_sonarqube. This allows toolchain_tool_sonarqube to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_toolchain_tool_sonarqube" "toolchain_tool_sonarqube" {
  toolchain_id = "toolchain_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `name` - (Optional, String) Name of tool integration.
* `parameters` - (Optional, List) 
Nested scheme for **parameters**:
	* `blind_connection` - (Optional, Boolean) Select this checkbox only if the server is not addressable on the public internet. IBM Cloud will not be able to validate the connection details you provide.
	  * Constraints: The default value is `false`.
	* `dashboard_url` - (Required, String) Type the URL of the SonarQube instance that you want to open when you click the SonarQube card in your toolchain.
	* `name` - (Required, String) Type a name for this tool integration, for example: my-sonarqube. This name displays on your toolchain.
	* `user_login` - (Optional, String) If you are using an authentication token, leave this field empty.
	* `user_password` - (Optional, String)
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain to bind integration to.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the toolchain_tool_sonarqube.
* `crn` - (Required, String) 
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

You can import the `ibm_toolchain_tool_sonarqube` resource by using `id`.
The `id` property can be formed from `toolchain_id`, and `integration_id` in the following format:

```
<toolchain_id>/<integration_id>
```
* `toolchain_id`: A string. ID of the toolchain to bind integration to.
* `integration_id`: A string. ID of the tool integration to be deleted.

# Syntax
```
$ terraform import ibm_toolchain_tool_sonarqube.toolchain_tool_sonarqube <toolchain_id>/<integration_id>
```
