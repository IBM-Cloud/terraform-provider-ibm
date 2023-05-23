---
layout: "ibm"
page_title: "IBM : ibm_code_engine_project"
description: |-
  Get information about code_engine_project
subcategory: "Code Engine"
---

# ibm_code_engine_project

Provides a read-only data source for code_engine_project. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_code_engine_project" "code_engine_project" {
	project_id = "15314cc3-85b4-4338-903f-c28cdee6d005"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `project_id` - (Required, Forces new resource, String) The ID of the project.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the code_engine_project.
* `account_id` - (String) An alphanumeric value identifying the account ID.

* `created_at` - (String) The timestamp when the project was created.

* `crn` - (String) The CRN of the project.

* `href` - (String) When you provision a new resource, a URL is created identifying the location of the instance.
  * Constraints: The maximum length is `2048` characters. The minimum length is `0` characters. The value must match regular expression `/(([^:\/?#]+):)?(\/\/([^\/?#]*))?([^?#]*)(\\?([^#]*))?(#(.*))?$/`.

* `name` - (String) The name of the project.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^([^\\x00-\\x7F]|[a-zA-Z0-9\\-._: ])+$/`.

* `region` - (String) The region for your project deployment.
  * Constraints: Possible values are: `au-syd`, `br-sao`, `ca-tor`, `eu-de`, `eu-gb`, `jp-osa`, `jp-tok`, `us-east`, `us-south`.

* `resource_group_id` - (String) The ID of the resource group.
  * Constraints: The maximum length is `32` characters. The minimum length is `32` characters. The value must match regular expression `/^[a-z0-9]*$/`.

* `resource_type` - (String) The type of the project.
  * Constraints: Allowable values are: `project_v2`.

* `status` - (String) The current state of the project. For example, if the project is created and ready to get used, it will return `active`. After deleting a project it will remain in `status` `soft_deleted` for a seven day period, during which it will still be retrievable.
  * Constraints: Possible values are: `active`, `inactive`, `pending_removal`, `hard_deleting`, `hard_deletion_failed`, `hard_deleted`, `deleting`, `deletion_failed`, `soft_deleted`, `preparing`, `creating`, `creation_failed`.

