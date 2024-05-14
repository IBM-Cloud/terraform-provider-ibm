---
layout: "ibm"
page_title: "IBM : ibm_logs_e2m"
description: |-
  Manages logs_e2m.
subcategory: "Cloud Logs"
---

~> **Beta:** This resource is in Beta, and is subject to change.

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
* `description` - (Optional, String) E2m description.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
* `logs_query` - (Optional, List) logs query.
Nested schema for **logs_query**:
	* `alias` - (Optional, String) alias.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
	* `applicationname_filters` - (Optional, List) application name filters.
	  * Constraints: The list items must match regular expression `/^.*$/`. The maximum length is `4096` items. The minimum length is `1` item.
	* `lucene` - (Optional, String) lucene query.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
	* `severity_filters` - (Optional, List) severity type filters.
	  * Constraints: Allowable list items are: `unspecified`, `debug`, `verbose`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `1` item.
	* `subsystemname_filters` - (Optional, List) subsystem names filters.
	  * Constraints: The list items must match regular expression `/^.*$/`. The maximum length is `4096` items. The minimum length is `1` item.
* `metric_fields` - (Optional, List) E2M metric fields.
  * Constraints: The maximum length is `10` items. The minimum length is `1` item.
Nested schema for **metric_fields**:
	* `aggregations` - (Optional, List) represents Aggregation type list.
	  * Constraints: The maximum length is `4096` items. The minimum length is `1` item.
	Nested schema for **aggregations**:
		* `agg_type` - (Optional, String) aggregation type.
		  * Constraints: The default value is `unspecified`. Allowable values are: `unspecified`, `min`, `max`, `count`, `avg`, `sum`, `histogram`, `samples`.
		* `enabled` - (Optional, Boolean) is enabled.
		* `histogram` - (Optional, List) e2m aggregate histogram type metadata.
		Nested schema for **histogram**:
			* `buckets` - (Optional, List) buckets that describe the e2m.
			  * Constraints: The maximum length is `4096` items. The minimum length is `1` item.
		* `samples` - (Optional, List) e2m sample type metadata.
		Nested schema for **samples**:
			* `sample_type` - (Optional, String) sample type min/max.
			  * Constraints: The default value is `unspecified`. Allowable values are: `unspecified`, `min`, `max`.
		* `target_metric_name` - (Optional, String) target metric field.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
	* `source_field` - (Optional, String) source field.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
	* `target_base_metric_name` - (Optional, String) target metric field.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\w/-]+$/`.
* `metric_labels` - (Optional, List) E2M metric labels.
  * Constraints: The maximum length is `4096` items. The minimum length is `1` item.
Nested schema for **metric_labels**:
	* `source_field` - (Optional, String) metric label source field.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
	* `target_label` - (Optional, String) metric label target label.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\w/-]+$/`.
* `name` - (Required, String) E2M name.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
* `spans_query` - (Optional, List) spans query.
Nested schema for **spans_query**:
	* `action_filters` - (Optional, List) action filters.
	  * Constraints: The list items must match regular expression `/^.*$/`. The maximum length is `4096` items. The minimum length is `1` item.
	* `applicationname_filters` - (Optional, List) application name filters.
	  * Constraints: The list items must match regular expression `/^.*$/`. The maximum length is `4096` items. The minimum length is `1` item.
	* `lucene` - (Optional, String) lucene query.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
	* `service_filters` - (Optional, List) service filters.
	  * Constraints: The list items must match regular expression `/^.*$/`. The maximum length is `4096` items. The minimum length is `1` item.
	* `subsystemname_filters` - (Optional, List) subsystem name filters.
	  * Constraints: The list items must match regular expression `/^.*$/`. The maximum length is `4096` items. The minimum length is `1` item.
* `type` - (Optional, String) e2m type.
  * Constraints: The default value is `unspecified`. Allowable values are: `unspecified`, `logs2metrics`, `spans2metrics`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the logs_e2m resource.
* `e2m_id` - The unique identifier of the logs_e2m.
* `create_time` - (String) E2M create time.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
* `is_internal` - (Boolean) a flag that represents if the e2m is for internal usage.
* `permutations` - (List) represents E2M permutations limit.
Nested schema for **permutations**:
	* `has_exceeded_limit` - (Boolean) flag to indicate if limit was exceeded.
	* `limit` - (Integer) e2m permutation limit.
* `update_time` - (String) E2M update time.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.


## Import

You can import the `ibm_logs_e2m` resource by using `id`. `id` combination of `region`, `instance_id` and `e2m_id`.

# Syntax
<pre>
$ terraform import ibm_logs_e2m.logs_e2m <region>/<instance_id>/<e2m_id>;
</pre>

# Example
```
$ terraform import ibm_logs_e2m.logs_e2m eu-gb/3dc02998-0b50-4ea8-b68a-4779d716fa1f/d6a3658e-78d2-47d0-9b81-b2c551f01b09
```
