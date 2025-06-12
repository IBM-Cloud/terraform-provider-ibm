---
layout: "ibm"
page_title: "IBM : ibm_logs_data_access_rules"
description: |-
  Get information about logs_data_access_rules
subcategory: "Cloud Logs"
---

# ibm_logs_data_access_rules

Provides a read-only data source to retrieve information about logs_data_access_rules. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_logs_data_access_rules" "logs_data_access_rules_instance" {
  instance_id               = "9d392fb2-b01b-40d5-9aec-fe21d02ab6ed"
  region                    = "eu-de"
  logs_data_access_rules_id = [ibm_logs_data_access_rule.logs_data_access_rule_instance.access_rule_id]
}

```

## Argument Reference

You can specify the following arguments for this data source.

* `instance_id` - (Required, String)  Cloud Logs Instance GUID.
* `region` - (Optional, String) Cloud Logs Instance Region.
* `logs_data_access_rules_id` - (Optional, List) Array of data access rule IDs.
  * Constraints: The list items must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`. The maximum length is `4096` items. The minimum length is `0` items.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the logs_data_access_rules.
* `data_access_rules` - (List) Data Access Rule details.
  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
Nested schema for **data_access_rules**:
	* `default_expression` - (String) Default expression to use when no filter matches the query.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_\\.,\\-"{}()\\[\\]=!:#\/$|'<> ]+$/`.
	* `description` - (String) Optional Data Access Rule Description.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_\\-\\s]+$/`.
	* `display_name` - (String) Data Access Rule Display Name.
	  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `^[\\p{L}\\p{N}\\p{P}\\p{Z}\\p{S}\\p{M}]+$`.
	* `filters` - (List) List of filters that the Data Access Rule is composed of.
	  * Constraints: The maximum length is `4096` items. The minimum length is `0` items.
	Nested schema for **filters**:
		* `entity_type` - (String) Filter's Entity Type.
		  * Constraints: Allowable values are: `unspecified`, `logs`.
		* `expression` - (String) Filter's Expression.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9_\\.,\\-"{}()\\[\\]=!:#\/$|'<> ]+$/`.
	* `id` - (String) Data Access Rule ID.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.

