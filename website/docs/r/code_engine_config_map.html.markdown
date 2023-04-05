---
layout: "ibm"
page_title: "IBM : ibm_code_engine_config_map"
description: |-
  Manages code_engine_config_map.
subcategory: "Code Engine"
---

# ibm_code_engine_config_map

Provides a resource for code_engine_config_map. This allows code_engine_config_map to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_code_engine_config_map" "code_engine_config_map_instance" {
  project_id = ibm_code_engine_project.code_engine_project_instance.project_id
  name       = "my-config-map"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `data` - (Optional, Map) The key-value pair for the config map. Values must be specified in `KEY=VALUE` format. Each `KEY` field must consist of alphanumeric characters, `-`, `_` or `.` and must not be exceed a max length of 253 characters. Each `VALUE` field can consists of any character and must not be exceed a max length of 1048576 characters.
* `name` - (Required, String) The name of the config map. Use a name that is unique within the project.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([\\-a-z0-9]*[a-z0-9])?)*$/`.
* `project_id` - (Required, Forces new resource, String) The ID of the project.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the code_engine_config_map.
* `config_map_id` - (String) The identifier of the resource.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.
* `created_at` - (String) The timestamp when the resource was created.
* `entity_tag` - (String) The version of the config map instance, which is used to achieve optimistic locking.
* `href` - (String) When you provision a new config map,  a URL is created identifying the location of the instance.
  * Constraints: The maximum length is `2048` characters. The minimum length is `0` characters. The value must match regular expression `/(([^:\/?#]+):)?(\/\/([^\/?#]*))?([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `resource_type` - (String) The type of the config map.
  * Constraints: Allowable values are: `config_map_v2`.
* `etag` - ETag identifier for code_engine_config_map.

## Import

You can import the `ibm_code_engine_config_map` resource by using `name`.
The `name` property can be formed from `project_id`, and `name` in the following format:

```
<project_id>/<name>
```
* `project_id`: A string in the format `15314cc3-85b4-4338-903f-c28cdee6d005`. The ID of the project.
* `name`: A string in the format `my-config-map`. The name of your configmap.

# Syntax
```
$ terraform import ibm_code_engine_config_map.code_engine_config_map <project_id>/<name>
```
