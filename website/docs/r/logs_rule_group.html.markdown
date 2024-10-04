---
layout: "ibm"
page_title: "IBM : ibm_logs_rule_group"
description: |-
  Manages logs_rule_group.
subcategory: "Cloud Logs"
---


# ibm_logs_rule_group

Create, update, and delete logs_rule_groups with this resource.

## Example Usage

```hcl
resource "ibm_logs_rule_group" "logs_rule_group_instance" {
  instance_id = ibm_resource_instance.logs_instance.guid
  region      = ibm_resource_instance.logs_instance.location
  name        = "example-rule-group"
  description = "example rule group decription"
  enabled     = false
  rule_matchers {
    subsystem_name {
      value = "mysql"
    }
  }
  rule_subgroups {
    rules {
      name         = "mysql-parse"
      source_field = "text"
      parameters {
        parse_parameters {
          destination_field = "text"
          rule              = "(?P<timestamp>[^,]+),(?P<hostname>[^,]+),(?P<username>[^,]+),(?P<ip>[^,]+),(?P<connectionId>[0-9]+),(?P<queryId>[0-9]+),(?P<operation>[^,]+),(?P<database>[^,]+),'?(?P<object>.*)'?,(?P<returnCode>[0-9]+)"
        }
      }
      enabled = true
      order   = 1
    }

    enabled = true
    order   = 1
  }
  order = 4294967
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `instance_id` - (Required, Forces new resource, String)  Cloud Logs Instance GUID.
* `region` - (Optional, Forces new resource, String) Cloud Logs Instance Region.
* `endpoint_type` - (Optional, String) Cloud Logs Instance Endpoint type. Allowed values `public` and `private`.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
* `description` - (Optional, String) A description for the rule group, should express what is the rule group purpose.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
* `enabled` - (Optional, Boolean) Whether or not the rule is enabled.
* `name` - (Required, String) The name of the rule group.
  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
* `order` - (Optional, Integer) // The order in which the rule group will be evaluated. The lower the order, the more priority the group will have. Not providing the order will by default create a group with the last order.
  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
* `rule_matchers` - (Optional, List) // Optional rule matchers which if matched will make the rule go through the rule group.
  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
Nested schema for **rule_matchers**:
	* `application_name` - (Optional, List) ApplicationName constraint.
	Nested schema for **application_name**:
		* `value` - (Required, String) Only logs with this ApplicationName value will match.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
	* `severity` - (Optional, List) Severity constraint.
	Nested schema for **severity**:
		* `value` - (Required, String) Only logs with this severity value will match.
		  * Constraints: Allowable values are: `debug_or_unspecified`, `verbose`, `info`, `warning`, `error`, `critical`.
	* `subsystem_name` - (Optional, List) SubsystemName constraint.
	Nested schema for **subsystem_name**:
		* `value` - (Required, String) Only logs with this SubsystemName value will match.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
* `rule_subgroups` - (Required, List) Rule subgroups. Will try to execute the first rule subgroup, and if not matched will try to match the next one in order.
  * Constraints: The maximum length is `4096` items. The minimum length is `1` item.
Nested schema for **rule_subgroups**:
	* `enabled` - (Optional, Boolean) Whether or not the rule subgroup is enabled.
	* `id` - (Required, String) The ID of the rule subgroup.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
	* `order` - (Required, Integer) The ordering of the rule subgroup. Lower order will run first. 0 is considered as no value.
	  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
	* `rules` - (Required, List) Rules to run on the log.
	  * Constraints: The maximum length is `4096` items. The minimum length is `1` item.
	Nested schema for **rules**:
		* `description` - (Optional, String) Description of the rule.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_\\-\\s]+$/`.
		* `enabled` - (Required, Boolean) Whether or not to execute the rule.
		* `id` - (Required, String) Unique identifier of the rule.
		  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
		* `name` - (Required, String) Name of the rule.
		  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
		* `order` - (Required, Integer) The ordering of the rule subgroup. Lower order will run first. 0 is considered as no value.
		  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
		* `parameters` - (Required, List) Parameters for a rule which specifies how it should run.
		Nested schema for **parameters**:
			* `allow_parameters` - (Optional, List) Parameters for allow rule.
			Nested schema for **allow_parameters**:
				* `keep_blocked_logs` - (Required, Boolean) If true matched logs will be blocked, otherwise matched logs will be kept.
				* `rule` - (Required, String) Regex which will match the source field and decide if the rule will apply.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
			* `block_parameters` - (Optional, List) Parameters for block rule.
			Nested schema for **block_parameters**:
				* `keep_blocked_logs` - (Required, Boolean) If true matched logs will be kept, otherwise matched logs will be blocked.
				* `rule` - (Required, String) Regex which will match the source field and decide if the rule will apply.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
			* `extract_parameters` - (Optional, List) Parameters for text extraction rule.
			Nested schema for **extract_parameters**:
				* `rule` - (Required, String) Regex which will parse the source field and extract the json keys from it while retaining the original log.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
			* `extract_timestamp_parameters` - (Optional, List) Parameters for extract timestamp rule.
			Nested schema for **extract_timestamp_parameters**:
				* `format` - (Required, String) What time format the the source field to extract from has.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
				* `standard` - (Required, String) What time format to use on the extracted time.
				  * Constraints: Allowable values are: `strftime_or_unspecified`, `javasdf`, `golang`, `secondsts`, `millits`, `microts`, `nanots`.
			* `json_extract_parameters` - (Optional, List) Parameters for json extract rule.
			Nested schema for **json_extract_parameters**:
				* `destination_field` - (Optional, String) In which metadata field to store the extracted value.
				  * Constraints: Allowable values are: `category_or_unspecified`, `classname`, `methodname`, `threadid`, `severity`.
			* `json_parse_parameters` - (Optional, List) Parameters for json parse rule.
			Nested schema for **json_parse_parameters**:
				* `delete_source` - (Optional, Boolean) Whether or not to delete the source field after running this rule.
				* `destination_field` - (Required, String) Destination field under which to put the json object.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
				* `override_dest` - (Required, Boolean) Destination field in which to put the json stringified content.
			* `json_stringify_parameters` - (Optional, List) Parameters for json stringify rule.
			Nested schema for **json_stringify_parameters**:
				* `delete_source` - (Optional, Boolean) Whether or not to delete the source field after running this rule.
				* `destination_field` - (Required, String) Destination field in which to put the json stringified content.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
			* `parse_parameters` - (Optional, List) Parameters for parse rule.
			Nested schema for **parse_parameters**:
				* `destination_field` - (Required, String) In which field to put the parsed text.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
				* `rule` - (Required, String) Regex which will parse the source field and extract the json keys from it while removing the source field.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
			* `remove_fields_parameters` - (Optional, List) Parameters for remove fields rule.
			Nested schema for **remove_fields_parameters**:
				* `fields` - (Required, List) Json field paths to drop from the log.
				  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
			* `replace_parameters` - (Optional, List) Parameters for replace rule.
			Nested schema for **replace_parameters**:
				* `destination_field` - (Required, String) In which field to put the modified text.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
				* `replace_new_val` - (Required, String) The value to replace the matched text with.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
				* `rule` - (Required, String) Regex which will match parts in the text to replace.
				  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
		* `source_field` - (Required, String) A field on which value to execute the rule.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the logs_rule_group resource.
* `rule_group_id` - The unique identifier of the logs_rule_group.


## Import

You can import the `ibm_logs_rule_group` resource by using `id`. `id` combination of `region`, `instance_id` and `rule_group_id`.

# Syntax
<pre>
$ terraform import ibm_logs_rule_group.logs_rule_group < region >/< instance_id >/< rule_group_id >;
</pre>

# Example
```
$ terraform import ibm_logs_rule_group.logs_rule_group eu-gb/3dc02998-0b50-4ea8-b68a-4779d716fa1f/3dc02998-0b50-4ea8-b68a-4779d716fa1f
```
