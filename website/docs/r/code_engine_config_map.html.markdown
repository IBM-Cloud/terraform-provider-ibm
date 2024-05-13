---
layout: "ibm"
page_title: "IBM : ibm_code_engine_config_map"
description: |-
  Manages code_engine_config_map.
subcategory: "Code Engine"
---

# ibm_code_engine_config_map

Create, update, and delete code_engine_config_maps with this resource.

## Example Usage

```hcl
resource "ibm_code_engine_config_map" "code_engine_config_map_instance" {
  project_id = ibm_code_engine_project.code_engine_project_instance.project_id
  name       = "my-config-map"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `data` - (Optional, Map) The key-value pair for the config map. Values must be specified in `KEY=VALUE` format. Each `KEY` field must consist of alphanumeric characters, `-`, `_` or `.` and must not be exceed a max length of 253 characters. Each `VALUE` field can consists of any character and must not be exceed a max length of 1048576 characters.
* `name` - (Required, Forces new resource, String) The name of the config map.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([\\-a-z0-9]*[a-z0-9])?)*$/`.
* `project_id` - (Required, Forces new resource, String) The ID of the project.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the code_engine_config_map.
* `config_map_id` - (String) The identifier of the resource.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.
* `created_at` - (String) The timestamp when the resource was created.
* `entity_tag` - (String) The version of the config map instance, which is used to achieve optimistic locking.
* `href` - (String) When you provision a new config map,  a URL is created identifying the location of the instance.
  * Constraints: The maximum length is `2048` characters. The minimum length is `0` characters. The value must match regular expression `/(([^:\/?#]+):)?(\/\/([^\/?#]*))?([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `region` - (String) The region of the project the resource is located in. Possible values: 'au-syd', 'br-sao', 'ca-tor', 'eu-de', 'eu-gb', 'jp-osa', 'jp-tok', 'us-east', 'us-south'.
* `resource_type` - (String) The type of the config map.
  * Constraints: Allowable values are: `config_map_v2`.
* `etag` - ETag identifier for code_engine_config_map.

## Import

You can import the `ibm_code_engine_config_map` resource by using `name`.
The `name` property can be formed from `project_id`, and `name` in the following format:

<pre>
&lt;project_id&gt;/&lt;name&gt;
</pre>
* `project_id`: A string in the format `15314cc3-85b4-4338-903f-c28cdee6d005`. The ID of the project.
* `name`: A string in the format `my-config-map`. The name of the config map.

# Syntax
<pre>
$ terraform import ibm_code_engine_config_map.code_engine_config_map &lt;project_id&gt;/&lt;name&gt;
</pre>
