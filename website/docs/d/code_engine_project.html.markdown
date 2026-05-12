---
layout: "ibm"
page_title: "IBM : ibm_code_engine_project"
description: |-
  Get information about code_engine_project
subcategory: "Code Engine"
---

# ibm_code_engine_project

Provides a read-only data source to retrieve information about a code_engine_project. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_code_engine_project" "code_engine_project" {
	project_id = "15314cc3-85b4-4338-903f-c28cdee6d005"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `project_id` - (Required, Forces new resource, String) The ID of the project.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the code_engine_project.

* `account_id` - (String) An alphanumeric value identifying the account ID.

* `created_at` - (String) The timestamp when the project was created.

* `crn` - (String) The CRN of the project.

* `href` - (String) When you provision a new resource, a URL is created identifying the location of the instance.
  * Constraints: The maximum length is `2048` characters. The minimum length is `0` characters. The value must match regular expression `/(([^:\/?#]+):)?(\/\/([^\/?#]*))?([^?#]*)(\\?([^#]*))?(#(.*))?$/`.

* `name` - (Forces new resource, String) The name of the project.
  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^([^\\x00-\\x7F]|[a-zA-Z0-9\\-._: ])+$/`.

* `region` - (String) The region for your project deployment. Possible values: `au-syd`, `br-sao`, `ca-tor`, `eu-de`, `eu-es`, `eu-gb`, `jp-osa`, `jp-tok`, `us-east`, `us-south`.

* `resource_group_id` - (String) The ID of the resource group.
  * Constraints: The maximum length is `32` characters. The minimum length is `32` characters. The value must match regular expression `/^[a-z0-9]*$/`.

* `resource_type` - (String) The type of the project.
  * Constraints: Allowable values are: `project_v2`.

* `status` - (String) The current state of the project. For example, when the project is created and is ready for use, the status of the project is active. After deleting a project it will remain in `status` `soft_deleted` for a seven day period, during which it will still be retrievable.
  * Constraints: Allowable values are: `active`, `inactive`, `pending_removal`, `hard_deleting`, `hard_deletion_failed`, `hard_deleted`, `deleting`, `deletion_failed`, `soft_deleted`, `preparing`, `creating`, `creation_failed`.

