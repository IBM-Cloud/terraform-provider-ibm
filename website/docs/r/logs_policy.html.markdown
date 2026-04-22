---
layout: "ibm"
page_title: "IBM : ibm_logs_policy"
description: |-
  Manages logs_policy.
subcategory: "Cloud Logs"
---


# ibm_logs_policy

Create, update, and delete logs_policys with this resource.

## Example Usage

```hcl
resource "ibm_logs_policy" "logs_policy_instance" {
  instance_id = ibm_resource_instance.logs_instance.guid
  region      = ibm_resource_instance.logs_instance.location
  name        = "example-policy"
  description = "example policy decription"
  priority    = "type_medium"
  application_rule {
    name         = "otel-links-test"
    rule_type_id = "start_with"
  }
  archive_retention_tag = "Default"
  log_rules {
    severities = ["info"]
  }
  before {
    id = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `instance_id` - (Required, Forces new resource, String)  Cloud Logs Instance GUID.
* `region` - (Optional, Forces new resource, String) Cloud Logs Instance Region.
* `endpoint_type` - (Optional, String) Cloud Logs Instance Endpoint type. Allowed values `public` and `private`.
* `application_rule` - (Optional, List) Rule for matching with application.
Nested schema for **application_rule**:
	* `name` - (Required, String) Name of the rule. Multiple values can be provided as comma separated string of values.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `rule_type_id` - (Required, String) Identifier of the rule.
	  * Constraints: Allowable values are: `unspecified`, `is`, `is_not`, `start_with`, `includes`.
* `archive_retention_tag` - (Optional, String) Archive retention tag. Required when retention tags are active. Cannot be set when retention tags are not active.
  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_-]+$/`.
* `before` - (Optional, List) The policy will be inserted immediately before the existing policy with this ID. If unspecified, the policy will be inserted after all existing policies.
Nested schema for **before**:
	* `id` - (Required, String) The policy will be inserted immediately before the existing policy with this ID. If unspecified, the policy will be inserted after all existing policies.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
	* `name` - (Computed, String) Policy name.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
* `description` - (Optional, String) Description of policy.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}0-9_\\-\\s]+$/`.
* `enabled` - (Optional, Boolean) Flag to enable or disable a policy. This flag is supported only while updating a policy, since the policies are always enabled during creation.
* `log_rules` - (Optional, List) Log rules.
Nested schema for **log_rules**:
	* `severities` - (Optional, List) The source severities to be used when matching.
	  * Constraints: Allowable list items are: `debug`, `verbose`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `0` items.
* `name` - (Required, String) Name of policy.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
* `priority` - (Required, String) The data pipeline sources that match the policy rules will continue to be processed by Cloud Logs.
  * Constraints: Allowable values are: `type_unspecified`, `type_block`, `type_low`, `type_medium`, `type_high`.
* `subsystem_rule` - (Optional, List) Rule for matching the application name.
Nested schema for **subsystem_rule**:
	* `name` - (Required, String) Name of the rule. Multiple values can be provided as comma separated string of values.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$/`.
	* `rule_type_id` - (Required, String) Identifier of the rule.
	  * Constraints: Allowable values are: `unspecified`, `is`, `is_not`, `start_with`, `includes`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the logs_policy.
* `policy_id` - The unique identifier of the logs_policy.
* `company_id` - (Integer) Company ID.
  * Constraints: The maximum value is `2147483647`. The minimum value is `0`.
* `created_at` - (String) Created at date at utc+0.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^"\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2}"$/`.
* `deleted` - (Boolean) Soft deletion flag.
* `order` - (Integer) Order of policy in relation to other policies.
  * Constraints: The maximum value is `2147483647`. The minimum value is `0`.
* `updated_at` - (String) Updated at date at utc+0.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^"\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2}"$/`.


## Import

You can import the `ibm_logs_policy` resource by using `id`. `id` combination of `region`, `instance_id` and `policy_id`.

# Syntax
<pre>
$ terraform import ibm_logs_policy.logs_policy < region >/< instance_id >/< policy_id >;
</pre>

# Example
```
$ terraform import ibm_logs_policy.logs_policy eu-gb/3dc02998-0b50-4ea8-b68a-4779d716fa1f/3dc02998-0b50-4ea8-b68a-4779d716fa1f
```
