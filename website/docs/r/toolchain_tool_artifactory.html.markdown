---
layout: "ibm"
page_title: "IBM : ibm_toolchain_tool_artifactory"
description: |-
  Manages toolchain_tool_artifactory.
subcategory: "Toolchain"
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

* `name` - (Optional, String) Name of tool integration.
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
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain to bind integration to.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the toolchain_tool_artifactory.
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

You can import the `ibm_toolchain_tool_artifactory` resource by using `id`.
The `id` property can be formed from `toolchain_id`, and `integration_id` in the following format:

```
<toolchain_id>/<integration_id>
```
* `toolchain_id`: A string. ID of the toolchain to bind integration to.
* `integration_id`: A string. ID of the tool integration to be deleted.

# Syntax
```
$ terraform import ibm_toolchain_tool_artifactory.toolchain_tool_artifactory <toolchain_id>/<integration_id>
```
