---
layout: "ibm"
page_title: "IBM : ibm_toolchain_tool_git"
description: |-
  Manages toolchain_tool_git.
subcategory: "IBM Toolchain API"
---

# ibm_toolchain_tool_git

Provides a resource for toolchain_tool_git. This allows toolchain_tool_git to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_toolchain_tool_git" "toolchain_tool_git" {
  git_provider = "git_provider"
  initialization = {  }
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
* `git_provider` - (Required, Forces new resource, String) 
* `initialization` - (Required, List) 
Nested scheme for **initialization**:
	* `private_repo` - (Optional, Forces new resource, Boolean) Select this check box to make this repository private.
	  * Constraints: The default value is `false`.
	* `repo_name` - (Optional, Forces new resource, String)
	* `repo_url` - (Optional, Forces new resource, String) Type the URL of the repository that you are linking to.
	* `source_repo_url` - (Optional, Forces new resource, String) Type the URL of the repository that you are forking or cloning.
	* `type` - (Optional, Forces new resource, String)
	  * Constraints: Allowable values are: `new`, `fork`, `clone`, `link`.
* `parameters` - (Optional, List) 
Nested scheme for **parameters**:
	* `enable_traceability` - (Optional, Boolean)
	* `has_issues` - (Optional, Boolean)
	* `private_repo` - (Optional, Boolean) Select this check box to make this repository private.
	  * Constraints: The default value is `false`.
	* `repo_name` - (Optional, String)
	* `repo_url` - (Optional, String) Type the URL of the repository that you are linking to.
	* `source_repo_url` - (Optional, String) Type the URL of the repository that you are forking or cloning.
	* `type` - (Optional, String)
	  * Constraints: Allowable values are: `new`, `fork`, `clone`, `link`.
* `toolchain_id` - (Required, Forces new resource, String) 

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the toolchain_tool_git.
* `dashboard_url` - (Required, String) The URL of a user-facing user interface for this instance of a service.

## Import

You can import the `ibm_toolchain_tool_git` resource by using `instance_id`. The id of the created or updated service instance.

# Syntax
```
$ terraform import ibm_toolchain_tool_git.toolchain_tool_git <instance_id>
```

# Example
```
$ terraform import ibm_toolchain_tool_git.toolchain_tool_git 4f107490-3820-400b-a008-f7f38d4163ed
```
