---
layout: "ibm"
page_title: "IBM : ibm_logs_e2m"
description: |-
  Manages logs_e2m.
subcategory: "Cloud Logs"
---


# ibm_logs_e2m

Create, update, and delete logs_e2ms with this resource.

## Example Usage

```hcl
resource "ibm_logs_e2m" "logs_e2m_instance" {
  instance_id = ibm_resource_instance.logs_instance.guid
  region      = ibm_resource_instance.logs_instance.location
  name        = "example-E2M"
  description = "example E2M decription"
  logs_query {
    applicationname_filters = []
    severity_filters = [
      "debug", "error"
    ]
    subsystemname_filters = []
  }
  type = "logs2metrics"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `instance_id` - (Required, Forces new resource, String)  Cloud Logs Instance GUID.
* `region` - (Optional, Forces new resource, String) Cloud Logs Instance Region.
* `endpoint_type` - (Optional, String) Cloud Logs Instance Endpoint type. Allowed values `public` and `private`.
* `description` - (Optional, String) Description of the E2M.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_\\-\\s]+$/`.
* `logs_query` - (Optional, List) E2M logs query.
Nested schema for **logs_query**:
	* `alias` - (Optional, String) Alias.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
	* `applicationname_filters` - (Optional, List) Application name filters.
	  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
	* `lucene` - (Optional, String) Lucene query.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
	* `severity_filters` - (Optional, List) Severity type filters.
	  * Constraints: Allowable list items are: `unspecified`, `debug`, `verbose`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.
	* `subsystemname_filters` - (Optional, List) Subsystem names filters.
	  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
* `metric_fields` - (Optional, List) E2M metric fields.
  * Constraints: The maximum length is `10` items. The minimum length is `0` items.
Nested schema for **metric_fields**:
	* `aggregations` - (Optional, List) Represents Aggregation type list.
	  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
	Nested schema for **aggregations**:
		* `agg_type` - (Optional, String) Aggregation type.
		  * Constraints: Allowable values are: `unspecified`, `min`, `max`, `count`, `avg`, `sum`, `histogram`, `samples`.
		* `enabled` - (Optional, Boolean) Is enabled.
		* `histogram` - (Optional, List) E2M aggregate histogram type metadata.
		Nested schema for **histogram**:
			* `buckets` - (Optional, List) Buckets of the E2M.
			  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
		* `samples` - (Optional, List) E2M sample type metadata.
		Nested schema for **samples**:
			* `sample_type` - (Optional, String) Sample type min/max.
			  * Constraints: Allowable values are: `unspecified`, `min`, `max`.
		* `target_metric_name` - (Optional, String) Target metric field alias name.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
	* `source_field` - (Optional, String) Source field.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
	* `target_base_metric_name` - (Optional, String) Target metric field alias name.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\w\/-]+$/`.
* `metric_labels` - (Optional, List) E2M metric labels.
  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
Nested schema for **metric_labels**:
	* `source_field` - (Optional, String) Metric label source field.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
	* `target_label` - (Optional, String) Metric label target alias name.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\w\/-]+$/`.
* `name` - (Required, String) Name of the E2M.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
* `type` - (Optional, String) E2M type.
  * Constraints: Allowable values are: `unspecified`, `logs2metrics`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the logs_e2m resource.
* `e2m_id` - The unique identifier of the logs e2m.
* `create_time` - (String) E2M create time.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
* `is_internal` - (Boolean) A flag that represents if the e2m is for internal usage.
* `permutations` - (List) Represents the limit of the permutations and if the limit was exceeded.
Nested schema for **permutations**:
	* `has_exceeded_limit` - (Boolean) Flag to indicate if limit was exceeded.
	* `limit` - (Integer) E2M permutation limit.
* `update_time` - (String) E2M update time.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.


## Import

You can import the `ibm_logs_e2m` resource by using `id`. `id` combination of `region`, `instance_id` and `e2m_id`.

# Syntax
<pre>
$ terraform import ibm_logs_e2m.logs_e2m < region >/< instance_id >/< e2m_id >;
</pre>

# Example
```
$ terraform import ibm_logs_e2m.logs_e2m eu-gb/3dc02998-0b50-4ea8-b68a-4779d716fa1f/d6a3658e-78d2-47d0-9b81-b2c551f01b09
```
