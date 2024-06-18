---
layout: "ibm"
page_title: "IBM : ibm_code_engine_config_map"
description: |-
  Get information about code_engine_config_map
subcategory: "Code Engine"
---

# ibm_code_engine_config_map

Provides a read-only data source to retrieve information about a code_engine_config_map. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_code_engine_config_map" "code_engine_config_map" {
	project_id = data.ibm_code_engine_project.code_engine_project.project_id
	name       = "my-config-map"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `name` - (Required, Forces new resource, String) The name of your configmap.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([\\-a-z0-9]*[a-z0-9])?)*$/`.
* `project_id` - (Required, Forces new resource, String) The ID of the project.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the code_engine_config_map.

* `config_map_id` - (String) The identifier of the resource.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.

* `created_at` - (String) The timestamp when the resource was created.

* `data` - (Map) The key-value pair for the config map. Values must be specified in `KEY=VALUE` format.

* `entity_tag` - (String) The version of the config map instance, which is used to achieve optimistic locking.

* `href` - (String) When you provision a new config map,  a URL is created identifying the location of the instance.
  * Constraints: The maximum length is `2048` characters. The minimum length is `0` characters. The value must match regular expression `/(([^:\/?#]+):)?(\/\/([^\/?#]*))?([^?#]*)(\\?([^#]*))?(#(.*))?$/`.

* `region` - (String) The region of the project the resource is located in. Possible values: 'au-syd', 'br-sao', 'ca-tor', 'eu-de', 'eu-gb', 'jp-osa', 'jp-tok', 'us-east', 'us-south'.

* `resource_type` - (String) The type of the config map.
  * Constraints: Allowable values are: `config_map_v2`.

