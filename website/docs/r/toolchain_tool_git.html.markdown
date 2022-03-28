---
layout: "ibm"
page_title: "IBM : ibm_toolchain_tool_git"
description: |-
  Manages toolchain_tool_git.
subcategory: "Toolchain"
---

# ibm_toolchain_tool_git

Provides a resource for toolchain_tool_git. This allows toolchain_tool_git to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_toolchain_tool_git" "toolchain_tool_git" {
  git_provider = "githubintegrated"
  toolchain_id = "toolchain_id"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `git_provider` - (Required, Forces new resource, String) The Git provider.
  * Constraints: The maximum length is `100` characters.
* `initialization` - (Optional, List) 
Nested scheme for **initialization**:
	* `git_id` - (Optional, Forces new resource, String)
	* `owner_id` - (Optional, Forces new resource, String)
	* `private_repo` - (Optional, Forces new resource, Boolean) Select this check box to make this repository private.
	  * Constraints: The default value is `false`.
	* `repo_name` - (Optional, Forces new resource, String)
	* `repo_url` - (Optional, Forces new resource, String) Type the URL of the repository that you are linking to.
	* `source_repo_url` - (Optional, Forces new resource, String) Type the URL of the repository that you are forking or cloning.
	* `type` - (Optional, Forces new resource, String)
	  * Constraints: Allowable values are: `new`, `fork`, `clone`, `link`.
* `name` - (Optional, String) Name of tool integration.
* `parameters` - (Optional, List) 
Nested scheme for **parameters**:
	* `enable_traceability` - (Optional, Boolean)
	* `git_id` - (Optional, String)
	* `has_issues` - (Optional, Boolean)
	* `owner_id` - (Optional, String)
	* `private_repo` - (Optional, Boolean) Select this check box to make this repository private.
	  * Constraints: The default value is `false`.
	* `repo_name` - (Optional, String)
	* `repo_url` - (Optional, String) Type the URL of the repository that you are linking to.
	* `source_repo_url` - (Optional, String) Type the URL of the repository that you are forking or cloning.
	* `type` - (Optional, String)
	  * Constraints: Allowable values are: `new`, `fork`, `clone`, `link`.
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain to bind integration to.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the toolchain_tool_git.
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

You can import the `ibm_toolchain_tool_git` resource by using `id`.
The `id` property can be formed from `toolchain_id`, and `integration_id` in the following format:

```
<toolchain_id>/<integration_id>
```
* `toolchain_id`: A string. ID of the toolchain to bind integration to.
* `integration_id`: A string. ID of the tool integration to be deleted.

# Syntax
```
$ terraform import ibm_toolchain_tool_git.toolchain_tool_git <toolchain_id>/<integration_id>
```
