---
layout: "ibm"
page_title: "IBM : ibm_logs_e2ms"
description: |-
  Get information about logs_e2ms
subcategory: "Cloud Logs"
---


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
  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
Nested schema for **events2metrics**:
	* `create_time` - (String) E2M create time.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_\\.,\\-"{}()\\[\\]=!:#\/$|' ]+$/`.
	* `description` - (String) Description of the E2M.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_\\-\\s]+$/`.
	* `id` - (String) E2M unique ID, required on update requests.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
	* `is_internal` - (Boolean) A flag that represents if the e2m is for internal usage.
	* `logs_query` - (List) E2M logs query.
	Nested schema for **logs_query**:
		* `alias` - (String) Alias.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_\\.,\\-"{}()\\[\\]=!:#\/$|' ]+$/`.
		* `applicationname_filters` - (List) Application name filters.
		  * Constraints: The list items must match regular expression `/^[A-Za-z0-9_\\.,\\-"{}()\\[\\]=!:#\/$|' ]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
		* `lucene` - (String) Lucene query.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_\\.,\\-"{}()\\[\\]=!:#\/$|' ]+$/`.
		* `severity_filters` - (List) Severity type filters.
		  * Constraints: Allowable list items are: `unspecified`, `debug`, `verbose`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.
		* `subsystemname_filters` - (List) Subsystem names filters.
		  * Constraints: The list items must match regular expression `/^[A-Za-z0-9_\\.,\\-"{}()\\[\\]=!:#\/$|' ]+$/`. The maximum length is `4096` items. The minimum length is `0` items.
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
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_\\.,\\-"{}()\\[\\]=!:#\/$|' ]+$/`.
		* `source_field` - (String) Source field.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_\\.,\\-"{}()\\[\\]=!:#\/$|' ]+$/`.
		* `target_base_metric_name` - (String) Target metric field alias name.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\w\/-]+$/`.
	* `metric_labels` - (List) E2M metric labels.
	  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
	Nested schema for **metric_labels**:
		* `source_field` - (String) Metric label source field.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_\\.,\\-"{}()\\[\\]=!:#\/$|' ]+$/`.
		* `target_label` - (String) Metric label target alias name.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\w\/-]+$/`.
	* `name` - (String) Name of the E2M.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_\\.,\\-"{}()\\[\\]=!:#\/$|' ]+$/`.
	* `permutations` - (List) Represents the limit of the permutations and if the limit was exceeded.
	Nested schema for **permutations**:
		* `has_exceeded_limit` - (Boolean) Flag to indicate if limit was exceeded.
		* `limit` - (Integer) E2M permutation limit.
	* `type` - (String) E2M type.
	  * Constraints: Allowable values are: `unspecified`, `logs2metrics`.
	* `update_time` - (String) E2M update time.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_\\.,\\-"{}()\\[\\]=!:#\/$|' ]+$/`.

