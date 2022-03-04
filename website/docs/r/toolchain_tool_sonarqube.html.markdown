---
layout: "ibm"
page_title: "IBM : ibm_toolchain_tool_sonarqube"
description: |-
  Manages toolchain_tool_sonarqube.
subcategory: "IBM Toolchain API"
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

* `container` - (Optional, Forces new resource, List) 
Nested scheme for **container**:
	* `guid` - (Required, String)
	* `type` - (Required, String)
	  * Constraints: Allowable values are: `organization_guid`, `resource_group_id`.
* `parameters` - (Optional, List) 
Nested scheme for **parameters**:
	* `blind_connection` - (Optional, Boolean) Select this checkbox only if the server is not addressable on the public internet. IBM Cloud will not be able to validate the connection details you provide.
	  * Constraints: The default value is `false`.
	* `dashboard_url` - (Required, String) Type the URL of the SonarQube instance that you want to open when you click the SonarQube card in your toolchain.
	* `name` - (Required, String) Type a name for this tool integration, for example: my-sonarqube. This name displays on your toolchain.
	* `user_login` - (Optional, String) If you are using an authentication token, leave this field empty.
	* `user_password` - (Optional, String)
* `parameters_references` - (Optional, Map) Decoded values used on provision in the broker that reference fields in the parameters.
* `toolchain_id` - (Required, Forces new resource, String) 

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the toolchain_tool_sonarqube.
* `dashboard_url` - (Required, String) The URL of a user-facing user interface for this instance of a service.

## Import

You can import the `ibm_toolchain_tool_sonarqube` resource by using `instance_id`. The id of the created or updated service instance.

# Syntax
```
$ terraform import ibm_toolchain_tool_sonarqube.toolchain_tool_sonarqube <instance_id>
```

# Example
```
$ terraform import ibm_toolchain_tool_sonarqube.toolchain_tool_sonarqube 4f107490-3820-400b-a008-f7f38d4163ed
```
