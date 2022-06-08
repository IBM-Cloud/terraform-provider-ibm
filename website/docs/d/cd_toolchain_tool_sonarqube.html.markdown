---
layout: "ibm"
page_title: "IBM : ibm_cd_toolchain_tool_sonarqube"
description: |-
  Get information about cd_toolchain_tool_sonarqube
subcategory: "CD Toolchain"
---

# ibm_cd_toolchain_tool_sonarqube

Provides a read-only data source for cd_toolchain_tool_sonarqube. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_cd_toolchain_tool_sonarqube" "cd_toolchain_tool_sonarqube" {
	tool_id = "tool_id"
	toolchain_id = ibm_cd_toolchain_tool_sonarqube.cd_toolchain_tool_sonarqube.toolchain_id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `tool_id` - (Required, Forces new resource, String) ID of the tool bound to the toolchain.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the cd_toolchain_tool_sonarqube.
* `crn` - (Required, String) Tool CRN.

* `get_tool_by_id_response_id` - (Required, String) Tool ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

* `href` - (Required, String) URI representing the tool.

* `name` - (Optional, String) Tool name.

* `parameters` - (Required, List) Parameters to be used to create the tool.
Nested scheme for **parameters**:
	* `blind_connection` - (Optional, Boolean) Select this checkbox only if the server is not addressable on the public internet. IBM Cloud will not be able to validate the connection details you provide.
	  * Constraints: The default value is `false`.
	* `dashboard_url` - (Required, String) Type the URL of the SonarQube instance that you want to open when you click the SonarQube card in your toolchain.
	* `name` - (Required, String) Type a name for this tool integration, for example: my-sonarqube. This name displays on your toolchain.
	* `user_login` - (Optional, String) If you are using an authentication token, leave this field empty.
	* `user_password` - (Optional, String)

* `referent` - (Required, List) Information on URIs to access this resource through the UI or API.
Nested scheme for **referent**:
	* `api_href` - (Optional, String) URI representing the this resource through an API.
	* `ui_href` - (Optional, String) URI representing the this resource through the UI.

* `resource_group_id` - (Required, String) Resource group where tool can be found.

* `state` - (Required, String) Current configuration state of the tool.
  * Constraints: Allowable values are: `configured`, `configuring`, `misconfigured`, `unconfigured`.

* `toolchain_crn` - (Required, String) CRN of toolchain which the tool is bound to.

* `updated_at` - (Required, String) Latest tool update timestamp.

