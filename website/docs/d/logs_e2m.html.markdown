---
layout: "ibm"
page_title: "IBM : ibm_logs_e2m"
description: |-
  Get information about logs_e2m
subcategory: "Cloud Logs"
---


# ibm_logs_e2m

Provides a read-only data source to retrieve information about a logs_e2m. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_logs_e2m" "logs_e2m_instance" {
  instance_id = ibm_logs_e2m.logs_e2m_instance.instance_id
  region      = ibm_logs_e2m.logs_e2m_instance.region
  logs_e2m_id = ibm_logs_e2m.logs_e2m_instance.e2m_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `instance_id` - (Required, String)  Cloud Logs Instance GUID.
* `region` - (Optional, String) Cloud Logs Instance Region.
* `logs_e2m_id` - (Required, Forces new resource, String) ID of e2m to be deleted.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the logs_e2m.
* `create_time` - (String) E2M create time.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.

* `description` - (String) Description of the E2M.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_\\-\\s]+$/`.

* `is_internal` - (Boolean) A flag that represents if the e2m is for internal usage.

* `logs_query` - (List) E2M logs query.
Nested schema for **logs_query**:
	* `alias` - (String) Alias.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
	* `applicationname_filters` - (List) Application name filters.
	  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.
	* `lucene` - (String) Lucene query.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
	* `severity_filters` - (List) Severity type filters.
	  * Constraints: Allowable list items are: `unspecified`, `debug`, `verbose`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.
	* `subsystemname_filters` - (List) Subsystem names filters.
	  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `0` items.

* `metric_fields` - (List) E2M metric fields.
  * Constraints: The maximum length is `10` items. The minimum length is `0` items.
Nested schema for **metric_fields**:
	* `aggregations` - (List) Represents Aggregation type list.
	  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
	Nested schema for **aggregations**:
		* `agg_type` - (String) Aggregation type.
		  * Constraints: Allowable values are: `unspecified`, `min`, `max`, `count`, `avg`, `sum`, `histogram`, `samples`.
		* `enabled` - (Boolean) Is enabled.
		* `histogram` - (List) E2M aggregate histogram type metadata.
		Nested schema for **histogram**:
			* `buckets` - (List) Buckets of the E2M.
			  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
		* `samples` - (List) E2M sample type metadata.
		Nested schema for **samples**:
			* `sample_type` - (String) Sample type min/max.
			  * Constraints: Allowable values are: `unspecified`, `min`, `max`.
		* `target_metric_name` - (String) Target metric field alias name.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
	* `source_field` - (String) Source field.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
	* `target_base_metric_name` - (String) Target metric field alias name.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\w\/-]+$/`.

* `metric_labels` - (List) E2M metric labels.
  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
Nested schema for **metric_labels**:
	* `source_field` - (String) Metric label source field.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
	* `target_label` - (String) Metric label target alias name.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\w\/-]+$/`.

* `name` - (String) Name of the E2M.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.

* `permutations` - (List) Represents the limit of the permutations and if the limit was exceeded.
Nested schema for **permutations**:
	* `has_exceeded_limit` - (Boolean) Flag to indicate if limit was exceeded.
	* `limit` - (Integer) E2M permutation limit.

* `type` - (String) E2M type.
  * Constraints: Allowable values are: `unspecified`, `logs2metrics`.

* `update_time` - (String) E2M update time.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.

