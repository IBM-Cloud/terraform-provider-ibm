---
layout: "ibm"
page_title: "IBM : ibm_scc_rule"
description: |-
  Manages scc_rule.
subcategory: "Security and Compliance Center"
---

# ibm_scc_rule

Create, update, and delete rules with this resource.

~> NOTE: if you specify the `region` in the provider, that region will become the default URL. Else, exporting the environmental variable IBMCLOUD_SCC_API_ENDPOINT will override any URL(ex. `export IBMCLOUD_SCC_API_ENDPOINT=https://us-south.compliance.cloud.ibm.com`).

## Example Usage

```hcl
resource "ibm_scc_rule" "scc_rule_instance" {
  instance_id = "00000000-1111-2222-3333-444444444444"
  description = "Example rule"
  import {
		parameters {
			name = "name"
			display_name = "display_name"
			description = "description"
			type = "string"
		}
  }
  required_config {
		description = "description"
		and {
			or {
				description = "description"
				property = "property"
				operator = "string_equals"
				value = "anything as a string"
			}
		}
  }
  target {
		service_name = "service_name"
		service_display_name = "service_display_name"
		resource_kind = "resource_kind"
		additional_target_attributes {
			name = "name"
			operator = "string_equals"
			value = "value"
		}
  }
  version = "1.0.0"
}
```

## Timeouts

scc_rule provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 60 minutes) Used for creating a scc_rule.
* `update` - (Default 60 minutes) Used for updating a scc_rule.
* `delete` - (Default 20 minutes) Used for deleting a scc_rule.

## Argument Reference

You can specify the following arguments for this resource.

* `instance_id` - (Required, Forces new resource, String) The ID of the SCC instance in a particular region.
* `description` - (Required, String) The details of a rule's response.
  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
* `import` - (Optional, List) The collection of import parameters.
Nested schema for **import**:
	* `parameters` - (Optional, List) The list of import parameters.
	  * Constraints: The maximum length is `8` items. The minimum length is `0` items.
	Nested schema for **parameters**:
		* `description` - (Optional, String) The propery description.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
		* `display_name` - (Optional, String) The display name of the property.
		  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
		* `name` - (Optional, String) The import parameter name.
		  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
		* `type` - (Optional, String) The property type.
		  * Constraints: Allowable values are: `string`, `numeric`, `general`, `boolean`, `string_list`, `ip_list`, `timestamp`. The maximum length is `11` characters. The minimum length is `6` characters. The value must match regular expression `/[A-Za-z]+/`.
* `labels` - (Optional, List) The list of labels.
  * Constraints: The list items must match regular expression `/[A-Za-z0-9]+/`. The maximum length is `32` items. The minimum length is `0` items.
* `required_config` - (Required, List) The required configurations.
Nested schema for **required_config**:
	* `and` - (Optional, List) The `AND` required configurations.
	  * Constraints: The maximum length is `64` items. The minimum length is `1` item.
	Nested schema for **and**:
		* `and` - (Optional, List) The `AND` required configurations.
		  * Constraints: The maximum length is `64` items. The minimum length is `1` item.
		Nested schema for **and**:
			* `description` - (Optional, String) The required config description.
			  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
			* `operator` - (Required, String) The operator.
			  * Constraints: Allowable values are: `string_equals`, `string_not_equals`, `string_match`, `string_not_match`, `string_contains`, `string_not_contains`, `num_equals`, `num_not_equals`, `num_less_than`, `num_less_than_equals`, `num_greater_than`, `num_greater_than_equals`, `is_empty`, `is_not_empty`, `is_true`, `is_false`, `strings_in_list`, `strings_allowed`, `strings_required`, `ips_in_range`, `ips_equals`, `ips_not_equals`, `days_less_than`. The maximum length is `23` characters. The minimum length is `7` characters.
			* `property` - (Required, String) The property.
			  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
			* `value` - (Optional, String) Schema for any JSON type.
		* `description` - (Optional, String) The required config description.
		  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
		* `operator` - (Optional, String) The operator.
		  * Constraints: Allowable values are: `string_equals`, `string_not_equals`, `string_match`, `string_not_match`, `string_contains`, `string_not_contains`, `num_equals`, `num_not_equals`, `num_less_than`, `num_less_than_equals`, `num_greater_than`, `num_greater_than_equals`, `is_empty`, `is_not_empty`, `is_true`, `is_false`, `strings_in_list`, `strings_allowed`, `strings_required`, `ips_in_range`, `ips_equals`, `ips_not_equals`, `days_less_than`. The maximum length is `23` characters. The minimum length is `7` characters.
		* `or` - (Optional, List) The `OR` required configurations.
		  * Constraints: The maximum length is `64` items. The minimum length is `1` item.
		Nested schema for **or**:
			* `description` - (Optional, String) The required config description.
			  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
			* `operator` - (Required, String) The operator.
			  * Constraints: Allowable values are: `string_equals`, `string_not_equals`, `string_match`, `string_not_match`, `string_contains`, `string_not_contains`, `num_equals`, `num_not_equals`, `num_less_than`, `num_less_than_equals`, `num_greater_than`, `num_greater_than_equals`, `is_empty`, `is_not_empty`, `is_true`, `is_false`, `strings_in_list`, `strings_allowed`, `strings_required`, `ips_in_range`, `ips_equals`, `ips_not_equals`, `days_less_than`. The maximum length is `23` characters. The minimum length is `7` characters.
			* `property` - (Required, String) The property.
			  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
			* `value` - (Optional, String) Schema for any JSON type.
		* `property` - (Optional, String) The property.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
		* `value` - (Optional, String) Schema for any JSON type.
	* `description` - (Optional, String) The required config description.
	  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
	* `operator` - (Optional, String) The operator.
	  * Constraints: Allowable values are: `string_equals`, `string_not_equals`, `string_match`, `string_not_match`, `string_contains`, `string_not_contains`, `num_equals`, `num_not_equals`, `num_less_than`, `num_less_than_equals`, `num_greater_than`, `num_greater_than_equals`, `is_empty`, `is_not_empty`, `is_true`, `is_false`, `strings_in_list`, `strings_allowed`, `strings_required`, `ips_in_range`, `ips_equals`, `ips_not_equals`, `days_less_than`. The maximum length is `23` characters. The minimum length is `7` characters.
	* `or` - (Optional, List) The `OR` required configurations.
	  * Constraints: The maximum length is `64` items. The minimum length is `1` item.
	Nested schema for **or**:
		* `and` - (Optional, List) The `AND` required configurations.
		  * Constraints: The maximum length is `64` items. The minimum length is `1` item.
		Nested schema for **and**:
			* `description` - (Optional, String) The required config description.
			  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
			* `operator` - (Required, String) The operator.
			  * Constraints: Allowable values are: `string_equals`, `string_not_equals`, `string_match`, `string_not_match`, `string_contains`, `string_not_contains`, `num_equals`, `num_not_equals`, `num_less_than`, `num_less_than_equals`, `num_greater_than`, `num_greater_than_equals`, `is_empty`, `is_not_empty`, `is_true`, `is_false`, `strings_in_list`, `strings_allowed`, `strings_required`, `ips_in_range`, `ips_equals`, `ips_not_equals`, `days_less_than`. The maximum length is `23` characters. The minimum length is `7` characters.
			* `property` - (Required, String) The property.
			  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
			* `value` - (Optional, String) Schema for any JSON type.
		* `description` - (Optional, String) The required config description.
		  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
		* `operator` - (Optional, String) The operator.
		  * Constraints: Allowable values are: `string_equals`, `string_not_equals`, `string_match`, `string_not_match`, `string_contains`, `string_not_contains`, `num_equals`, `num_not_equals`, `num_less_than`, `num_less_than_equals`, `num_greater_than`, `num_greater_than_equals`, `is_empty`, `is_not_empty`, `is_true`, `is_false`, `strings_in_list`, `strings_allowed`, `strings_required`, `ips_in_range`, `ips_equals`, `ips_not_equals`, `days_less_than`. The maximum length is `23` characters. The minimum length is `7` characters.
		* `or` - (Optional, List) The `OR` required configurations.
		  * Constraints: The maximum length is `64` items. The minimum length is `1` item.
		Nested schema for **or**:
			* `description` - (Optional, String) The required config description.
			  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
			* `operator` - (Required, String) The operator.
			  * Constraints: Allowable values are: `string_equals`, `string_not_equals`, `string_match`, `string_not_match`, `string_contains`, `string_not_contains`, `num_equals`, `num_not_equals`, `num_less_than`, `num_less_than_equals`, `num_greater_than`, `num_greater_than_equals`, `is_empty`, `is_not_empty`, `is_true`, `is_false`, `strings_in_list`, `strings_allowed`, `strings_required`, `ips_in_range`, `ips_equals`, `ips_not_equals`, `days_less_than`. The maximum length is `23` characters. The minimum length is `7` characters.
			* `property` - (Required, String) The property.
			  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
			* `value` - (Optional, String) Schema for any JSON type.
		* `property` - (Optional, String) The property.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
		* `value` - (Optional, String) Schema for any JSON type.
	* `property` - (Optional, String) The property.
	  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
	* `value` - (Optional, String) Schema for any JSON type.
* `target` - (Required, List) The rule target.
Nested schema for **target**:
	* `additional_target_attributes` - (Optional, List) The list of targets supported properties.
	  * Constraints: The maximum length is `99999` items. The minimum length is `0` items.
	Nested schema for **additional_target_attributes**:
		* `name` - (Optional, String) The additional target attribute name.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
		* `operator` - (Optional, String) The operator.
		  * Constraints: Allowable values are: `string_equals`, `string_not_equals`, `string_match`, `string_not_match`, `string_contains`, `string_not_contains`, `num_equals`, `num_not_equals`, `num_less_than`, `num_less_than_equals`, `num_greater_than`, `num_greater_than_equals`, `is_empty`, `is_not_empty`, `is_true`, `is_false`, `strings_in_list`, `strings_allowed`, `strings_required`, `ips_in_range`, `ips_equals`, `ips_not_equals`, `days_less_than`.
		* `value` - (Optional, String) The value.
		  * Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
	* `resource_kind` - (Required, String) The target resource kind.
	  * Constraints: The maximum length is `99999` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
	* `service_display_name` - (Optional, String) The display name of the target service.
	  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
	* `service_name` - (Required, String) The target service name.
	  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
* `version` - (Optional, String) The version number of a rule.
  * Constraints: The maximum length is `10` characters. The minimum length is `5` characters. The value must match regular expression `/^[0-9][0-9.]*$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the scc_rule.
* `rule_id` - (String) The ID that is associated with the created `rule`
* `account_id` - (String) The account ID.
  * Constraints: The maximum length is `32` characters. The minimum length is `3` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
* `created_by` - (String) The user who created the rule.
  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
* `created_on` - (String) The date when the rule was created.
* `type` - (String) The rule type (allowable values are `user_defined` or `system_defined`).
  * Constraints: Allowable values are: `user_defined`, `system_defined`. The maximum length is `14` characters. The minimum length is `12` characters. The value must match regular expression `/[A-Za-z]+_[A-Za-z]+/`.
* `updated_by` - (String) The user who modified the rule.
  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
* `updated_on` - (String) The date when the rule was modified.


## Import

You can import the `ibm_scc_rule` resource by using `id`. The rule ID.
The `id` property can be formed from `instance_id` and `rule_id` in the following format:

```bash
<instance_id>/<rule_id>
```
* `instance_id`: A string. The instance ID.
* `rule_id`: A string. The rule ID.

# Syntax

```bash
$ terraform import ibm_scc_rule.scc_rule <instance_id>/<rule_id>
```

# Example
```bash
$ terraform import ibm_scc_rule.scc_rule 00000000-1111-2222-3333-444444444444/00000000-1111-2222-3333-444444444444
```