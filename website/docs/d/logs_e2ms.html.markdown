---
layout: "ibm"
page_title: "IBM : ibm_logs_e2ms"
description: |-
  Get information about logs_e2ms
subcategory: "Cloud Logs"
---

~> **Beta:** This resource is in Beta, and is subject to change.

# ibm_logs_e2ms

Provides a read-only data source to retrieve information about logs_e2ms. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
	data "ibm_logs_e2ms" "logs_e2ms_instance" {
		instance_id = ibm_logs_e2m.logs_e2m_instance.instance_id
		region      = ibm_logs_e2m.logs_e2m_instance.region
	}
```

## Argument Reference

You can specify the following arguments for this data source.

* `instance_id` - (Required, String)  Cloud Logs Instance GUID.
* `region` - (Optional, String) Cloud Logs Instance Region.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the logs_e2ms.
* `events2metrics` - (List) List of event to metrics definitions.
  * Constraints: The maximum length is `4096` items. The minimum length is `1` item.
Nested schema for **events2metrics**:
	* `create_time` - (String) E2M create time.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
	* `description` - (String) E2m description.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
	* `id` - (String) E2M id, required on update requests.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
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

