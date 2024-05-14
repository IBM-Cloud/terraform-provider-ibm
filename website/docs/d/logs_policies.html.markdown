---
layout: "ibm"
page_title: "IBM : ibm_logs_policies"
description: |-
  Get information about logs_policies
subcategory: "Cloud Logs"
---

~> **Beta:** This resource is in Beta, and is subject to change.

# ibm_logs_policies

Provides a read-only data source to retrieve information about logs_policies. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_logs_policies" "logs_policies_instance" {
	instance_id  = ibm_logs_policy.logs_policy_instance.instance_id
	region       = ibm_logs_policy.logs_policy_instance.region
	enabled_only = true
	source_type  = "logs"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `instance_id` - (Required, String)  Cloud Logs Instance GUID.
* `region` - (Optional, String) Cloud Logs Instance Region.
* `enabled_only` - (Optional, Boolean) optionally filter only enabled policies.
* `source_type` - (Optional, String) Source type to filter policies by.
  * Constraints: The default value is `logs`. Allowable values are: `unspecified`, `logs`, `spans`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the logs_policies.
* `policies` - (List) company policies.
  * Constraints: The maximum length is `4096` items. The minimum length is `1` item.
Nested schema for **policies**:
	* `application_rule` - (List) rule for matching with application.
	Nested schema for **application_rule**:
		* `name` - (String) value of the rule.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
		* `rule_type_id` - (String) identifier of the rule.
		  * Constraints: The default value is `unspecified`. Allowable values are: `unspecified`, `is`, `is_not`, `start_with`, `includes`.
	* `archive_retention` - (List) archive retention definition.
	Nested schema for **archive_retention**:
		* `id` - (String) references archive retention definition.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
	* `company_id` - (Integer) company id.
	* `created_at` - (String) created at timestamp.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
	* `deleted` - (Boolean) soft deletion flag.
	* `description` - (String) description of policy.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
	* `enabled` - (Boolean) enabled flag.
	* `id` - (String) policy id.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
	* `log_rules` - (List) log rules.
	Nested schema for **log_rules**:
		* `severities` - (List) source severities to match with.
		  * Constraints: Allowable list items are: `unspecified`, `debug`, `verbose`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `1` item.
	* `name` - (String) name of policy.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
	* `order` - (Integer) order of policy in relation to other policies.
	* `priority` - (String) the data pipeline sources that match the policy rules will go through.
	  * Constraints: The default value is `type_unspecified`. Allowable values are: `type_unspecified`, `type_block`, `type_low`, `type_medium`, `type_high`.
	* `subsystem_rule` - (List) rule for matching with application.
	Nested schema for **subsystem_rule**:
		* `name` - (String) value of the rule.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
		* `rule_type_id` - (String) identifier of the rule.
		  * Constraints: The default value is `unspecified`. Allowable values are: `unspecified`, `is`, `is_not`, `start_with`, `includes`.
	* `updated_at` - (String) updated at timestamp.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.

