---
layout: "ibm"
page_title: "IBM : ibm_logs_policy"
description: |-
  Manages logs_policy.
subcategory: "Cloud Logs"
---

~> **Beta:** This resource is in Beta, and is subject to change.

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
  log_rules {
    severities = ["info"]
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `instance_id` - (Required, Forces new resource, String)  Cloud Logs Instance GUID.
* `region` - (Optional, Forces new resource, String) Cloud Logs Instance Region.
* `endpoint_type` - (Optional, String) Cloud Logs Instance Endpoint type. Allowed values `public` and `private`.
* `application_rule` - (Optional, List) rule for matching with application.
Nested schema for **application_rule**:
	* `name` - (Required, String) value of the rule.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
	* `rule_type_id` - (Required, String) identifier of the rule.
	  * Constraints: The default value is `unspecified`. Allowable values are: `unspecified`, `is`, `is_not`, `start_with`, `includes`.
* `archive_retention` - (Optional, List) archive retention definition.
Nested schema for **archive_retention**:
	* `id` - (Required, String) references archive retention definition.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
* `description` - (Optional, String) description of policy.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
* `log_rules` - (Optional, List) log rules.
Nested schema for **log_rules**:
	* `severities` - (Optional, List) source severities to match with.
	  * Constraints: Allowable list items are: `unspecified`, `debug`, `verbose`, `info`, `warning`, `error`, `critical`. The maximum length is `4096` items. The minimum length is `1` item.
* `name` - (Required, String) name of policy.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
* `priority` - (Required, String) the data pipeline sources that match the policy rules will go through.
  * Constraints: The default value is `type_unspecified`. Allowable values are: `type_unspecified`, `type_block`, `type_low`, `type_medium`, `type_high`.
* `subsystem_rule` - (Optional, List) rule for matching with application.
Nested schema for **subsystem_rule**:
	* `name` - (Required, String) value of the rule.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
	* `rule_type_id` - (Required, String) identifier of the rule.
	  * Constraints: The default value is `unspecified`. Allowable values are: `unspecified`, `is`, `is_not`, `start_with`, `includes`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the logs_policy resource.
* `policy_id` - The unique identifier of the logs_policy.
* `company_id` - (Integer) company id.
* `created_at` - (String) created at timestamp.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
* `deleted` - (Boolean) soft deletion flag.
* `enabled` - (Boolean) enabled flag.
* `order` - (Integer) order of policy in relation to other policies.
* `updated_at` - (String) updated at timestamp.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.


## Import

You can import the `ibm_logs_policy` resource by using `id`. `id` combination of `region`, `instance_id` and `policy_id`.

# Syntax
<pre>
$ terraform import ibm_logs_policy.logs_policy <region>/<instance_id>/<policy_id>;
</pre>
