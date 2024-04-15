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
data "ibm_logs_e2m" "logs_e2m" {
	logs_e2m_id = "logs_e2m_id"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `logs_e2m_id` - (Required, Forces new resource, String) id of e2m to be deleted.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the logs_e2m.
* `create_time` - (String) E2M create time.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.

* `description` - (String) E2m description.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.

* `is_internal` - (Boolean) a flag that represents if the e2m is for internal usage.

* `logs_query` - (List) logs query.
Nested schema for **logs_query**:
	* `alias` - (String) alias.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
	* `applicationname_filters` - (List) application name filters.
	  * Constraints: The list items must match regular expression `/^.*$/`. The maximum length is `4096` items. The minimum length is `1` item.
	* `lucene` - (String) lucene query.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
	* `severity_filters` - (List) severity type filters.
	  * Constraints: Allowable list items are: `unspecified`, `debug`, `verbose`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `1` item.
	* `subsystemname_filters` - (List) subsystem names filters.
	  * Constraints: The list items must match regular expression `/^.*$/`. The maximum length is `4096` items. The minimum length is `1` item.

* `metric_fields` - (List) E2M metric fields.
  * Constraints: The maximum length is `10` items. The minimum length is `1` item.
Nested schema for **metric_fields**:
	* `aggregations` - (List) represents Aggregation type list.
	  * Constraints: The maximum length is `4096` items. The minimum length is `1` item.
	Nested schema for **aggregations**:
		* `agg_type` - (String) aggregation type.
		  * Constraints: The default value is `unspecified`. Allowable values are: `unspecified`, `min`, `max`, `count`, `avg`, `sum`, `histogram`, `samples`.
		* `enabled` - (Boolean) is enabled.
		* `histogram` - (List) e2m aggregate histogram type metadata.
		Nested schema for **histogram**:
			* `buckets` - (List) buckets that describe the e2m.
			  * Constraints: The maximum length is `4096` items. The minimum length is `1` item.
		* `samples` - (List) e2m sample type metadata.
		Nested schema for **samples**:
			* `sample_type` - (String) sample type min/max.
			  * Constraints: The default value is `unspecified`. Allowable values are: `unspecified`, `min`, `max`.
		* `target_metric_name` - (String) target metric field.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
	* `source_field` - (String) source field.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
	* `target_base_metric_name` - (String) target metric field.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\w/-]+$/`.

* `metric_labels` - (List) E2M metric labels.
  * Constraints: The maximum length is `4096` items. The minimum length is `1` item.
Nested schema for **metric_labels**:
	* `source_field` - (String) metric label source field.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
	* `target_label` - (String) metric label target label.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\w/-]+$/`.

* `name` - (String) E2M name.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.

* `permutations` - (List) represents E2M permutations limit.
Nested schema for **permutations**:
	* `has_exceeded_limit` - (Boolean) flag to indicate if limit was exceeded.
	* `limit` - (Integer) e2m permutation limit.

* `spans_query` - (List) spans query.
Nested schema for **spans_query**:
	* `action_filters` - (List) action filters.
	  * Constraints: The list items must match regular expression `/^.*$/`. The maximum length is `4096` items. The minimum length is `1` item.
	* `applicationname_filters` - (List) application name filters.
	  * Constraints: The list items must match regular expression `/^.*$/`. The maximum length is `4096` items. The minimum length is `1` item.
	* `lucene` - (String) lucene query.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
	* `service_filters` - (List) service filters.
	  * Constraints: The list items must match regular expression `/^.*$/`. The maximum length is `4096` items. The minimum length is `1` item.
	* `subsystemname_filters` - (List) subsystem name filters.
	  * Constraints: The list items must match regular expression `/^.*$/`. The maximum length is `4096` items. The minimum length is `1` item.

* `type` - (String) e2m type.
  * Constraints: The default value is `unspecified`. Allowable values are: `unspecified`, `logs2metrics`, `spans2metrics`.

* `update_time` - (String) E2M update time.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.

