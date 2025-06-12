---
layout: "ibm"
page_title: "IBM : ibm_logs_rule_groups"
description: |-
  Get information about logs_rule_groups
subcategory: "Cloud Logs"
---


# ibm_logs_rule_groups

Provides a read-only data source to retrieve information about logs_rule_groups. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_logs_rule_groups" "logs_rule_groups_instance" {
  instance_id = ibm_logs_rule_group.logs_rule_group_instance.instance_id
  region      = ibm_logs_rule_group.logs_rule_group_instance.region
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `instance_id` - (Required, String)  Cloud Logs Instance GUID.
* `region` - (Optional, String) Cloud Logs Instance Region.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the logs_rule_groups.
* `rulegroups` - (List) The rule groups.
  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
Nested schema for **rulegroups**:
	* `description` - (String) A description for the rule group, should express what is the rule group purpose.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
	* `enabled` - (Boolean) Whether or not the rule is enabled.
	* `id` - (String) The ID of the rule group.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
	* `name` - (String) The name of the rule group.
	  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
	* `order` - (Integer) // The order in which the rule group will be evaluated. The lower the order, the more priority the group will have. Not providing the order will by default create a group with the last order.
	  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
	* `rule_matchers` - (List) // Optional rule matchers which if matched will make the rule go through the rule group.
	  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
	Nested schema for **rule_matchers**:
		* `application_name` - (List) ApplicationName constraint.
		Nested schema for **application_name**:
			* `value` - (String) Only logs with this ApplicationName value will match.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
		* `severity` - (List) Severity constraint.
		Nested schema for **severity**:
			* `value` - (String) Only logs with this severity value will match.
			  * Constraints: Allowable values are: `debug_or_unspecified`, `verbose`, `info`, `warning`, `error`, `critical`.
		* `subsystem_name` - (List) SubsystemName constraint.
		Nested schema for **subsystem_name**:
			* `value` - (String) Only logs with this SubsystemName value will match.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
	* `rule_subgroups` - (List) Rule subgroups. Will try to execute the first rule subgroup, and if not matched will try to match the next one in order.
	  * Constraints: The maximum length is `4096` items. The minimum length is `1` item.
	Nested schema for **rule_subgroups**:
		* `enabled` - (Boolean) Whether or not the rule subgroup is enabled.
		* `id` - (String) The ID of the rule subgroup.
		  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
		* `order` - (Integer) The ordering of the rule subgroup. Lower order will run first. 0 is considered as no value.
		  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
		* `rules` - (List) Rules to run on the log.
		  * Constraints: The maximum length is `4096` items. The minimum length is `1` item.
		Nested schema for **rules**:
			* `description` - (String) Description of the rule.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_\\-\\s]+$/`.
			* `enabled` - (Boolean) Whether or not to execute the rule.
			* `id` - (String) Unique identifier of the rule.
			  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
			* `name` - (String) Name of the rule.
			  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
			* `order` - (Integer) The ordering of the rule subgroup. Lower order will run first. 0 is considered as no value.
			  * Constraints: The maximum value is `4294967295`. The minimum value is `0`.
			* `parameters` - (List) Parameters for a rule which specifies how it should run.
			Nested schema for **parameters**:
				* `allow_parameters` - (List) Parameters for allow rule.
				Nested schema for **allow_parameters**:
					* `keep_blocked_logs` - (Boolean) If true matched logs will be blocked, otherwise matched logs will be kept.
					* `rule` - (String) Regex which will match the source field and decide if the rule will apply.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
				* `block_parameters` - (List) Parameters for block rule.
				Nested schema for **block_parameters**:
					* `keep_blocked_logs` - (Boolean) If true matched logs will be kept, otherwise matched logs will be blocked.
					* `rule` - (String) Regex which will match the source field and decide if the rule will apply.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
				* `extract_parameters` - (List) Parameters for text extraction rule.
				Nested schema for **extract_parameters**:
					* `rule` - (String) Regex which will parse the source field and extract the json keys from it while retaining the original log.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
				* `extract_timestamp_parameters` - (List) Parameters for extract timestamp rule.
				Nested schema for **extract_timestamp_parameters**:
					* `format` - (String) What time format the the source field to extract from has.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
					* `standard` - (String) What time format to use on the extracted time.
					  * Constraints: Allowable values are: `strftime_or_unspecified`, `javasdf`, `golang`, `secondsts`, `millits`, `microts`, `nanots`.
				* `json_extract_parameters` - (List) Parameters for json extract rule.
				Nested schema for **json_extract_parameters**:
					* `destination_field` - (String) In which metadata field to store the extracted value.
					  * Constraints: Allowable values are: `category_or_unspecified`, `classname`, `methodname`, `threadid`, `severity`.
				* `json_parse_parameters` - (List) Parameters for json parse rule.
				Nested schema for **json_parse_parameters**:
					* `delete_source` - (Boolean) Whether or not to delete the source field after running this rule.
					* `destination_field` - (String) Destination field under which to put the json object.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
					* `override_dest` - (Boolean) Destination field in which to put the json stringified content.
				* `json_stringify_parameters` - (List) Parameters for json stringify rule.
				Nested schema for **json_stringify_parameters**:
					* `delete_source` - (Boolean) Whether or not to delete the source field after running this rule.
					* `destination_field` - (String) Destination field in which to put the json stringified content.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
				* `parse_parameters` - (List) Parameters for parse rule.
				Nested schema for **parse_parameters**:
					* `destination_field` - (String) In which field to put the parsed text.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
					* `rule` - (String) Regex which will parse the source field and extract the json keys from it while removing the source field.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
				* `remove_fields_parameters` - (List) Parameters for remove fields rule.
				Nested schema for **remove_fields_parameters**:
					* `fields` - (List) Json field paths to drop from the log.
					  * Constraints: The list items must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`. The maximum length is `4096` items. The minimum length is `1` item.
				* `replace_parameters` - (List) Parameters for replace rule.
				Nested schema for **replace_parameters**:
					* `destination_field` - (String) In which field to put the modified text.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
					* `replace_new_val` - (String) The value to replace the matched text with.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
					* `rule` - (String) Regex which will match parts in the text to replace.
					  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
			* `source_field` - (String) A field on which value to execute the rule.
			  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.

