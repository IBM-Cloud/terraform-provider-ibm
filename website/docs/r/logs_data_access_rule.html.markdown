---
layout: "ibm"
page_title: "IBM : ibm_logs_data_access_rule"
description: |-
  Manages logs_data_access_rule.
subcategory: "Cloud Logs"
---

# ibm_logs_data_access_rule

Create, update, and delete logs_data_access_rules with this resource.

## Example Usage

```hcl
resource "ibm_logs_data_access_rule" "logs_data_access_rule_instance" {
  instance_id  = "9d392fb2-b01b-40d5-9aec-fe21d02ab6ed"
  region       = "eu-de"
  display_name = "Test Data Access Rule"
  description  = "Data Access Rule intended for testing"
  filters {
    entity_type = "logs"
    expression  = "<v1> foo == 'bar'"
  }
  default_expression = "<v1>true"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `instance_id` - (Required, Forces new resource, String) Cloud Logs Instance GUID.
* `region` - (Optional, Forces new resource, String) Cloud Logs Instance Region.
* `endpoint_type` - (Optional, String) Cloud Logs Instance Endpoint type. Allowed values `public` and `private`.
* `default_expression` - (Required, String) Default expression to use when no filter matches the query.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_\\.,\\-"{}()\\[\\]=!:#\/$|'<> ]+$/`.
* `description` - (Optional, String) Optional Data Access Rule Description.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_\\-\\s]+$/`.
* `display_name` - (Required, String) Data Access Rule Display Name.
  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
* `filters` - (Required, List) List of filters that the Data Access Rule is composed of.
  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
Nested schema for **filters**:
	* `entity_type` - (Required, String) Filter's Entity Type.
	  * Constraints: Allowable values are: `unspecified`, `logs`.
	* `expression` - (Required, String) Filter's Expression.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_\\.,\\-"{}()\\[\\]=!:#\/$|'<> ]+$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the logs_data_access_rule.
* `access_rule_id` - The unique identifier of the logs data access rule.

## Import

You can import the `ibm_logs_data_access_rule` resource by using `id`. `id` combination of `region`, `instance_id` and `access_rule_id`.

# Syntax
<pre>
$ terraform import ibm_logs_dashboard_folder.logs_dashboard_folder < region >/< instance_id >/< access_rule_id >;
</pre>

# Example
```
$ terraform import ibm_logs_data_access_rule.logs_data_access_rule eu-gb/3dc02998-0b50-4ea8-b68a-4779d716fa1f/d6a3658e-78d2-47d0-9b81-b2c551f01b09
```
