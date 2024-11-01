---
layout: "ibm"
page_title: "IBM : ibm_cbr_rule"
description: |-
  Get information about cbr_rule
subcategory: "Context Based Restrictions"
---

# ibm_cbr_rule

Provides a read-only data source to retrieve information about a cbr_rule. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_cbr_rule" "cbr_rule" {
	rule_id = "rule_id"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `rule_id` - (Required, Forces new resource, String) The ID of a rule.
  * Constraints: The maximum length is `32` characters. The minimum length is `32` characters. The value must match regular expression `/^[a-fA-F0-9]{32}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the cbr_rule.
* `contexts` - (List) The contexts this rule applies to.
  * Constraints: The maximum length is `1000` items. The minimum length is `0` items.
Nested schema for **contexts**:
	* `attributes` - (List) The attributes.
	  * Constraints: The minimum length is `1` item.
	Nested schema for **attributes**:
		* `name` - (String) The attribute name.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9]+$/`.
		* `value` - (String) The attribute value.
		  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[\S\s]+$/`.

* `created_at` - (String) The time the resource was created.
* `created_by_id` - (String) IAM ID of the user or service which created the resource.
* `crn` - (String) The rule CRN.
* `description` - (String) The description of the rule.
  * Constraints: The maximum length is `300` characters. The minimum length is `0` characters. The value must match regular expression `/^[\x20-\xFE]*$/`.
* `enforcement_mode` - (String) The rule enforcement mode: * `enabled` - The restrictions are enforced and reported. This is the default. * `disabled` - The restrictions are disabled. Nothing is enforced or reported. * `report` - The restrictions are evaluated and reported, but not enforced.
  * Constraints: The default value is `enabled`. Allowable values are: `enabled`, `disabled`, `report`.
* `href` - (String) The href link to the resource.
* `id` - (String) The globally unique ID of the rule.
* `last_modified_at` - (String) The last time the resource was modified.
* `last_modified_by_id` - (String) IAM ID of the user or service which modified the resource.
* `operations` - (List) The operations this rule applies to.
Nested schema for **operations**:
	* `api_types` - (List) The API types this rule applies to.
	  * Constraints: The maximum length is `100` items. The minimum length is `1` item.
	Nested schema for **api_types**:
		* `api_type_id` - (String)
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_.\-:]+$/`.
* `resources` - (List) The resources this rule apply to.
  * Constraints: The maximum length is `1` item. The minimum length is `1` item.
Nested schema for **resources**:
	* `attributes` - (List) The resource attributes.
	  * Constraints: The minimum length is `1` item.
	Nested schema for **attributes**:
		* `name` - (String) The attribute name.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_]+$/`.
		* `operator` - (String) The attribute operator.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9]+$/`.
		* `value` - (String) The attribute value.
		  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[\S\s]+$/`.
	* `tags` - (List) The optional resource tags.
	  * Constraints: The maximum length is `10` items. The minimum length is `1` item.
	Nested schema for **tags**:
		* `name` - (String) The tag attribute name.
		  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9 _.-]+$/`.
		* `operator` - (String) The attribute operator.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9]+$/`.
		* `value` - (String) The tag attribute value.
		  * Constraints: The maximum length is `1000` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9 _*?.-]+$/`.

