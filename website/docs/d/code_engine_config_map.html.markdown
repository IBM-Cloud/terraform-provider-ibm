---
layout: "ibm"
page_title: "IBM : ibm_code_engine_config_map"
description: |-
  Get information about code_engine_config_map
subcategory: "Code Engine"
---

# ibm_code_engine_config_map

Provides a read-only data source for code_engine_config_map. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_code_engine_config_map" "code_engine_config_map" {
	project_id = data.ibm_code_engine_project.code_engine_project.project_id
	name       = "my-config-map"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `name` - (Required, Forces new resource, String) The name of your configmap.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([\\-a-z0-9]*[a-z0-9])?)*$/`.
* `project_id` - (Required, Forces new resource, String) The ID of the project.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the code_engine_config_map.
* `config_map_id` - (String) The identifier of the resource.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.

* `created_at` - (String) The timestamp when the resource was created.

* `data` - (Map) The key-value pair for the config map. Values must be specified in `KEY=VALUE` format.

* `entity_tag` - (String) The version of the config map instance, which is used to achieve optimistic locking.

* `href` - (String) When you provision a new config map,  a URL is created identifying the location of the instance.
  * Constraints: The maximum length is `2048` characters. The minimum length is `0` characters. The value must match regular expression `/(([^:\/?#]+):)?(\/\/([^\/?#]*))?([^?#]*)(\\?([^#]*))?(#(.*))?$/`.

* `resource_type` - (String) The type of the config map.
  * Constraints: Allowable values are: `config_map_v2`.

