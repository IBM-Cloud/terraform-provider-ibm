---
layout: "ibm"
page_title: "IBM : ibm_logs_policy"
description: |-
  Get information about logs_policy
subcategory: "Cloud Logs"
---


# ibm_logs_policy

Provides a read-only data source to retrieve information about a logs_policy. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_logs_policy" "logs_policy_instance" {
	instance_id    = ibm_logs_policy.logs_policy_instance.instance_id
	region         = ibm_logs_policy.logs_policy_instance.region
	logs_policy_id = ibm_logs_policy.logs_policy_instance.policy_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `instance_id` - (Required, String)  Cloud Logs Instance GUID.
* `region` - (Optional, String) Cloud Logs Instance Region.
* `logs_policy_id` - (Required, Forces new resource, String) ID of policy.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the logs_policy.
* `application_rule` - (List) Rule for matching with application.
Nested schema for **application_rule**:
	* `name` - (String) Value of the rule.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_\\.,\\-"{}()\\[\\]=!:#\/$|' ]+$/`.
	* `rule_type_id` - (String) Identifier of the rule.
	  * Constraints: Allowable values are: `unspecified`, `is`, `is_not`, `start_with`, `includes`.

* `archive_retention` - (List) Archive retention definition.
Nested schema for **archive_retention**:
	* `id` - (String) References archive retention definition.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.

* `company_id` - (Integer) Company ID.

* `created_at` - (String) Created at date at utc+0.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^"\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2}"$/`.

* `deleted` - (Boolean) Soft deletion flag.

* `description` - (String) Description of policy.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_\\-\\s]+$/`.

* `enabled` - (Boolean) Enabled flag.

* `log_rules` - (List) Log rules.
Nested schema for **log_rules**:
	* `severities` - (List) Source severities to match with.
	  * Constraints: Allowable list items are: `unspecified`, `debug`, `verbose`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.

* `name` - (String) Name of policy.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_\\.,\\-"{}()\\[\\]=!:#\/$|' ]+$/`.

* `order` - (Integer) Order of policy in relation to other policies.

* `priority` - (String) The data pipeline sources that match the policy rules will go through.
  * Constraints: Allowable values are: `type_unspecified`, `type_block`, `type_low`, `type_medium`, `type_high`.

* `subsystem_rule` - (List) Rule for matching with application.
Nested schema for **subsystem_rule**:
	* `name` - (String) Value of the rule.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_\\.,\\-"{}()\\[\\]=!:#\/$|' ]+$/`.
	* `rule_type_id` - (String) Identifier of the rule.
	  * Constraints: Allowable values are: `unspecified`, `is`, `is_not`, `start_with`, `includes`.

* `updated_at` - (String) Updated at date at utc+0.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^"\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2}"$/`.

