---
layout: "ibm"
page_title: "IBM : ibm_code_engine_project"
description: |-
  Manages code_engine_project.
subcategory: "Code Engine"
---

# ibm_code_engine_project

Create and delete code_engine_projects with this resource.

## Example Usage

```hcl
data "ibm_resource_group" "group" {
  name = "<your_group>"
}

resource "ibm_code_engine_project" "code_engine_project_instance" {
  name              = "my-project"
  resource_group_id = data.ibm_resource_group.group.id
}
```

## Timeouts

code_engine_project provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 10 minutes) Used for creating a code_engine_project.
* `delete` - (Default Projects(/projects/{id}) minutes) Used for deleting a code_engine_project.

## Argument Reference

You can specify the following arguments for this resource.

* `name` - (Required, Forces new resource, String) The name of the project.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^([^\\x00-\\x7F]|[a-zA-Z0-9\\-._: ])+$/`.
* `resource_group_id` - (Optional, Forces new resource, String) The ID of the resource group. If this field is not defined, the default resource group of the account will be used.
  * Constraints: The maximum length is `32` characters. The minimum length is `32` characters. The value must match regular expression `/^[a-z0-9]*$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the code_engine_project.
* `project_id` - (String) The ID of the project.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.
* `account_id` - (String) An alphanumeric value identifying the account ID.
* `created_at` - (String) The timestamp when the project was created.
* `crn` - (String) The CRN of the project.
* `href` - (String) When you provision a new resource, a URL is created identifying the location of the instance.
  * Constraints: The maximum length is `2048` characters. The minimum length is `0` characters. The value must match regular expression `/(([^:\/?#]+):)?(\/\/([^\/?#]*))?([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `region` - (String) The region for your project deployment. Possible values: `au-syd`, `br-sao`, `ca-tor`, `eu-de`, `eu-es`, `eu-gb`, `jp-osa`, `jp-tok`, `us-east`, `us-south`.
* `resource_type` - (String) The type of the project.
  * Constraints: Allowable values are: `project_v2`.
* `status` - (String) The current state of the project. For example, when the project is created and is ready for use, the status of the project is active.
  * Constraints: Allowable values are: `active`, `inactive`, `pending_removal`, `hard_deleting`, `hard_deletion_failed`, `hard_deleted`, `deleting`, `deletion_failed`, `soft_deleted`, `preparing`, `creating`, `creation_failed`.

## Import

You can import the `ibm_code_engine_project` resource by using `id`. The ID of the project.

# Syntax
<pre>
$ terraform import ibm_code_engine_project.code_engine_project &lt;id&gt;
</pre>

# Example
```
$ terraform import ibm_code_engine_project.code_engine_project 4e49b3e0-27a8-48d2-a784-c7ee48bb863b
```
