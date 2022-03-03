---
layout: "ibm"
page_title: "IBM : ibm_toolchain_tool_artifactory"
description: |-
  Manages toolchain_tool_artifactory.
subcategory: "IBM Toolchain API"
---

# ibm_toolchain_tool_artifactory

Provides a resource for toolchain_tool_artifactory. This allows toolchain_tool_artifactory to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_toolchain_tool_artifactory" "toolchain_tool_artifactory" {
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
	* `dashboard_url` - (Optional, String) Type the URL that you want to navigate to when you click the Artifactory integration tile.
	* `docker_config_json` - (Optional, String)
	* `mirror_url` - (Optional, String) Type the URL for your Artifactory virtual repository, which is a repository that can see your private repositories and a cache of the public repositories.
	* `name` - (Required, String) Type a name for this tool integration, for example: my-artifactory. This name displays on your toolchain.
	* `release_url` - (Optional, String) Type the URL for your Artifactory release repository.
	* `repository_name` - (Optional, String) Type the name of your artifactory repository where your docker images are located.
	* `repository_url` - (Optional, String) Type the URL of your artifactory repository where your docker images are located.
	* `snapshot_url` - (Optional, String) Type the URL for your Artifactory snapshot repository.
	* `token` - (Optional, String) Type the API key for your Artifactory repository.
	* `type` - (Required, String) Choose the type of repository for your Artifactory integration.
	  * Constraints: Allowable values are: `npm`, `maven`, `docker`.
	* `user_id` - (Optional, String) Type the User ID or email for your Artifactory repository.
* `parameters_references` - (Optional, Map) Decoded values used on provision in the broker that reference fields in the parameters.
* `toolchain_id` - (Required, Forces new resource, String) 

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the toolchain_tool_artifactory.
* `dashboard_url` - (Required, String) The URL of a user-facing user interface for this instance of a service.

## Import

You can import the `ibm_toolchain_tool_artifactory` resource by using `instance_id`. The id of the created or updated service instance.

# Syntax
```
$ terraform import ibm_toolchain_tool_artifactory.toolchain_tool_artifactory <instance_id>
```

# Example
```
$ terraform import ibm_toolchain_tool_artifactory.toolchain_tool_artifactory 4f107490-3820-400b-a008-f7f38d4163ed
```
