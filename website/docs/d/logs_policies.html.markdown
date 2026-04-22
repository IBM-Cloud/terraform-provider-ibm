---
layout: "ibm"
page_title: "IBM : ibm_logs_policies"
description: |-
  Get information about logs_policies
subcategory: "Cloud Logs"
---

# ibm_logs_policies

Provides a read-only data source to retrieve information about logs_policies. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_logs_policies" "logs_policies_instance" {
	instance_id  = ibm_logs_policy.logs_policy_instance.instance_id
	region       = ibm_logs_policy.logs_policy_instance.region
	enabled_only = true
	source_type = "logs"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `instance_id` - (Required, String)  Cloud Logs Instance GUID.
* `region` - (Optional, String) Cloud Logs Instance Region.
* `enabled_only` - (Optional, Boolean) Optionally filter only enabled policies.
* `source_type` - (Optional, String) Source type to filter policies by.
  * Constraints: Allowable values are: `unspecified`, `logs`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the logs_policies.
* `policies` - (List) Company policies.
  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
Nested schema for **policies**:
	* `application_rule` - (List) Rule for matching with application.
	Nested schema for **application_rule**:
		* `name` - (String) Value of the rule. Multiple values can be provided as comma separated string of values.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
		* `rule_type_id` - (String) Identifier of the rule.
		  * Constraints: Allowable values are: `unspecified`, `is`, `is_not`, `start_with`, `includes`.
	* `archive_retention_tag` - (String) Archive retention tag. Required when retention tags are active. Cannot be set when retention tags are not active.
	  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_-]+$/`.
	* `before` - (List) The policy will be inserted immediately before the existing policy with this ID. If unspecified, the policy will be inserted after all existing policies.
	Nested schema for **before**:
		* `id` - (String) The policy will be inserted immediately before the existing policy with this ID. If unspecified, the policy will be inserted after all existing policies.
		  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
		* `name` - (String) Policy name.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `company_id` - (Integer) Company ID.
	  * Constraints: The maximum value is `2147483647`. The minimum value is `0`.
	* `created_at` - (String) Created at date at utc+0.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^"\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2}"$/`.
	* `deleted` - (Boolean) Soft deletion flag.
	* `description` - (String) Description of policy.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}0-9_\\-\\s]+$/`.
	* `enabled` - (Boolean) Flag to enable or disable a policy. This flag is supported only while updating a policy, since the policies are always enabled during creation.
	* `id` - (String) Policy ID.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
	* `log_rules` - (List) Log rules.
	Nested schema for **log_rules**:
		* `severities` - (List) The source severities to be used when matching.
		  * Constraints: Allowable list items are: `debug`, `verbose`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.
	* `name` - (String) Name of policy.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `order` - (Integer) Order of policy in relation to other policies.
	  * Constraints: The maximum value is `2147483647`. The minimum value is `0`.
	* `priority` - (String) The data pipeline sources that match the policy rules will continue to be processed by Cloud Logs.
	  * Constraints: Allowable values are: `type_unspecified`, `type_block`, `type_low`, `type_medium`, `type_high`.
	* `subsystem_rule` - (List) Rule for matching the application name.
	Nested schema for **subsystem_rule**:
		* `name` - (String) Name of the rule. Multiple values can be provided as comma separated string of values.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
		* `rule_type_id` - (String) Identifier of the rule.
		  * Constraints: Allowable values are: `unspecified`, `is`, `is_not`, `start_with`, `includes`.
	* `updated_at` - (String) Updated at date at utc+0.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^"\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2}"$/`.

