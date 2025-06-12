---
layout: "ibm"
page_title: "IBM : ibm_scc_rule"
description: |-
  Get information about scc_rule
subcategory: "Security and Compliance Center"
---

# ibm_scc_rule

Retrieve information about a rule from a read-only data source. Then, you can reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

~> NOTE: Security Compliance Center is a regional service. Please specify the IBM Cloud Provider attribute `region` to target another region. Else, exporting the environmental variable IBMCLOUD_SCC_API_ENDPOINT will also override which region is being targeted for all ibm providers(ex. `export IBMCLOUD_SCC_API_ENDPOINT=https://eu-es.compliance.cloud.ibm.com`).

## Example Usage

```hcl
data "ibm_scc_rule" "scc_rule" {
    instance_id = "00000000-1111-2222-3333-444444444444"
    rule_id = ibm_scc_rule.scc_rule_instance.rule_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `instance_id` - (Required, Forces new resource, String) The ID of the SCC instance in a particular region.
* `rule_id` - (Required, Forces new resource, String) The ID of the corresponding rule.
  * Constraints: The maximum length is `41` characters. The minimum length is `41` characters. The value must match regular expression `/rule-[0-9A-Fa-f]{8}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{12}/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the scc_rule.
* `account_id` - (String) The account ID.
  * Constraints: The maximum length is `32` characters. The minimum length is `3` characters. The value must match regular expression `/[A-Za-z0-9]+/`.

* `created_by` - (String) The user who created the rule.
  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.

* `created_on` - (String) The date when the rule was created.

* `description` - (String) The details of a rule's response.
  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.

* `id` - (String) The rule ID.
  * Constraints: The maximum length is `41` characters. The minimum length is `41` characters. The value must match regular expression `/rule-[0-9A-Fa-f]{8}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{4}-[0-9A-Fa-f]{12}/`.

* `import` - (List) The collection of import parameters.
Nested schema for **import**:
	* `parameters` - (List) The list of import parameters.
	  * Constraints: The maximum length is `8` items. The minimum length is `0` items.
	Nested schema for **parameters**:
		* `description` - (String) The propery description.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
		* `display_name` - (String) The display name of the property.
		  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
		* `name` - (String) The import parameter name.
		  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
		* `type` - (String) The property type.
		  * Constraints: Allowable values are: `string`, `numeric`, `general`, `boolean`, `string_list`, `ip_list`, `timestamp`. The maximum length is `11` characters. The minimum length is `6` characters. The value must match regular expression `/[A-Za-z]+/`.

* `labels` - (List) The list of labels.
  * Constraints: The list items must match regular expression `/[A-Za-z0-9]+/`. The maximum length is `32` items. The minimum length is `0` items.

* `required_config` - (List) The required configurations.
Nested schema for **required_config**:
	* `and` - (List) The `AND` required configurations.
	  * Constraints: The maximum length is `64` items. The minimum length is `1` item.
	Nested schema for **and**:
		* `and` - (List) The `AND` required configurations.
		  * Constraints: The maximum length is `64` items. The minimum length is `1` item.
		Nested schema for **and**:
			* `description` - (String) The required config description.
			  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
			* `operator` - (String) The operator.
			  * Constraints: Allowable values are: `string_equals`, `string_not_equals`, `string_match`, `string_not_match`, `string_contains`, `string_not_contains`, `num_equals`, `num_not_equals`, `num_less_than`, `num_less_than_equals`, `num_greater_than`, `num_greater_than_equals`, `is_empty`, `is_not_empty`, `is_true`, `is_false`, `strings_in_list`, `strings_allowed`, `strings_required`, `ips_in_range`, `ips_equals`, `ips_not_equals`, `days_less_than`. The maximum length is `23` characters. The minimum length is `7` characters.
			* `property` - (String) The property.
			  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
			* `value` - (String) Schema for any JSON type.
		* `description` - (String) The required config description.
		  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
		* `operator` - (String) The operator.
		  * Constraints: Allowable values are: `string_equals`, `string_not_equals`, `string_match`, `string_not_match`, `string_contains`, `string_not_contains`, `num_equals`, `num_not_equals`, `num_less_than`, `num_less_than_equals`, `num_greater_than`, `num_greater_than_equals`, `is_empty`, `is_not_empty`, `is_true`, `is_false`, `strings_in_list`, `strings_allowed`, `strings_required`, `ips_in_range`, `ips_equals`, `ips_not_equals`, `days_less_than`. The maximum length is `23` characters. The minimum length is `7` characters.
		* `or` - (List) The `OR` required configurations.
		  * Constraints: The maximum length is `64` items. The minimum length is `1` item.
		Nested schema for **or**:
			* `description` - (String) The required config description.
			  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
			* `operator` - (String) The operator.
			  * Constraints: Allowable values are: `string_equals`, `string_not_equals`, `string_match`, `string_not_match`, `string_contains`, `string_not_contains`, `num_equals`, `num_not_equals`, `num_less_than`, `num_less_than_equals`, `num_greater_than`, `num_greater_than_equals`, `is_empty`, `is_not_empty`, `is_true`, `is_false`, `strings_in_list`, `strings_allowed`, `strings_required`, `ips_in_range`, `ips_equals`, `ips_not_equals`, `days_less_than`. The maximum length is `23` characters. The minimum length is `7` characters.
			* `property` - (String) The property.
			  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
			* `value` - (String) Schema for any JSON type.
		* `property` - (String) The property.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
		* `value` - (String) Schema for any JSON type.
	* `description` - (String) The required config description.
	  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
	* `operator` - (String) The operator.
	  * Constraints: Allowable values are: `string_equals`, `string_not_equals`, `string_match`, `string_not_match`, `string_contains`, `string_not_contains`, `num_equals`, `num_not_equals`, `num_less_than`, `num_less_than_equals`, `num_greater_than`, `num_greater_than_equals`, `is_empty`, `is_not_empty`, `is_true`, `is_false`, `strings_in_list`, `strings_allowed`, `strings_required`, `ips_in_range`, `ips_equals`, `ips_not_equals`, `days_less_than`. The maximum length is `23` characters. The minimum length is `7` characters.
	* `or` - (List) The `OR` required configurations.
	  * Constraints: The maximum length is `64` items. The minimum length is `1` item.
	Nested schema for **or**:
		* `and` - (List) The `AND` required configurations.
		  * Constraints: The maximum length is `64` items. The minimum length is `1` item.
		Nested schema for **and**:
			* `description` - (String) The required config description.
			  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
			* `operator` - (String) The operator.
			  * Constraints: Allowable values are: `string_equals`, `string_not_equals`, `string_match`, `string_not_match`, `string_contains`, `string_not_contains`, `num_equals`, `num_not_equals`, `num_less_than`, `num_less_than_equals`, `num_greater_than`, `num_greater_than_equals`, `is_empty`, `is_not_empty`, `is_true`, `is_false`, `strings_in_list`, `strings_allowed`, `strings_required`, `ips_in_range`, `ips_equals`, `ips_not_equals`, `days_less_than`. The maximum length is `23` characters. The minimum length is `7` characters.
			* `property` - (String) The property.
			  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
			* `value` - (String) Schema for any JSON type.
		* `description` - (String) The required config description.
		  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
		* `operator` - (String) The operator.
		  * Constraints: Allowable values are: `string_equals`, `string_not_equals`, `string_match`, `string_not_match`, `string_contains`, `string_not_contains`, `num_equals`, `num_not_equals`, `num_less_than`, `num_less_than_equals`, `num_greater_than`, `num_greater_than_equals`, `is_empty`, `is_not_empty`, `is_true`, `is_false`, `strings_in_list`, `strings_allowed`, `strings_required`, `ips_in_range`, `ips_equals`, `ips_not_equals`, `days_less_than`. The maximum length is `23` characters. The minimum length is `7` characters.
		* `or` - (List) The `OR` required configurations.
		  * Constraints: The maximum length is `64` items. The minimum length is `1` item.
		Nested schema for **or**:
			* `description` - (String) The required config description.
			  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
			* `operator` - (String) The operator.
			  * Constraints: Allowable values are: `string_equals`, `string_not_equals`, `string_match`, `string_not_match`, `string_contains`, `string_not_contains`, `num_equals`, `num_not_equals`, `num_less_than`, `num_less_than_equals`, `num_greater_than`, `num_greater_than_equals`, `is_empty`, `is_not_empty`, `is_true`, `is_false`, `strings_in_list`, `strings_allowed`, `strings_required`, `ips_in_range`, `ips_equals`, `ips_not_equals`, `days_less_than`. The maximum length is `23` characters. The minimum length is `7` characters.
			* `property` - (String) The property.
			  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
			* `value` - (String) Schema for any JSON type.
		* `property` - (String) The property.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
		* `value` - (String) Schema for any JSON type.
	* `property` - (String) The property.
	  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
	* `value` - (String) Schema for any JSON type.

* `target` - (List) The rule target.
Nested schema for **target**:
	* `additional_target_attributes` - (List) The list of targets supported properties.
	  * Constraints: The maximum length is `99999` items. The minimum length is `0` items.
	Nested schema for **additional_target_attributes**:
		* `name` - (String) The additional target attribute name.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
		* `operator` - (String) The operator.
		  * Constraints: Allowable values are: `string_equals`, `string_not_equals`, `string_match`, `string_not_match`, `string_contains`, `string_not_contains`, `num_equals`, `num_not_equals`, `num_less_than`, `num_less_than_equals`, `num_greater_than`, `num_greater_than_equals`, `is_empty`, `is_not_empty`, `is_true`, `is_false`, `strings_in_list`, `strings_allowed`, `strings_required`, `ips_in_range`, `ips_equals`, `ips_not_equals`, `days_less_than`.
		* `value` - (String) The value.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
	* `resource_kind` - (String) The target resource kind.
	  * Constraints: The maximum length is `99999` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
	* `service_display_name` - (String) The display name of the target service.
	  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
	* `service_name` - (String) The target service name.
	  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.

* `type` - (String) The rule type (allowable values are `user_defined` or `system_defined`).
  * Constraints: Allowable values are: `user_defined`, `system_defined`. The maximum length is `14` characters. The minimum length is `12` characters. The value must match regular expression `/[A-Za-z]+_[A-Za-z]+/`.

* `updated_by` - (String) The user who modified the rule.
  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.

* `updated_on` - (String) The date when the rule was modified.

* `version` - (String) The version number of a rule.
  * Constraints: The maximum length is `10` characters. The minimum length is `5` characters. The value must match regular expression `/^[0-9][0-9.]*$/`.

