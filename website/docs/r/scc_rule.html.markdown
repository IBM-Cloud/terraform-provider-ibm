---
layout: "ibm"
page_title: "IBM : ibm_scc_rule"
description: |-
  Manages scc_rule.
subcategory: "Security and Compliance Center"
---

# ibm_scc_rule

Create, update, and delete rules with this resource.

~> NOTE: Security Compliance Center is a regional service. Please specify the IBM Cloud Provider attribute `region` to target another region. Else, exporting the environmental variable IBMCLOUD_SCC_API_ENDPOINT will also override which region is being targeted for all ibm providers(ex. `export IBMCLOUD_SCC_API_ENDPOINT=https://eu-es.compliance.cloud.ibm.com`).

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
```hcl
resource "ibm_scc_rule" "scc_rule_instance" {
			instance_id = "00000000-1111-2222-3333-444444444444"
			description = "This rule will determine if a cloud object storage bucket is configured my way"
			version = "1.0.0"
			import {
				parameters {
					name = "name"
					display_name = "display_name"
					description = "description"
					type = "string"
				}
			}
			target {
				service_name = "cloud-object-storage"
				resource_kind = "bucket"
				additional_target_attributes {
					name = "location"
					operator = "string_equals"
					value = "$${name}"
				}
			}
			required_config {
				and {
					or {
						description = "description 1"
						property = "storage_class"
						operator = "string_equals"
						value = "smart"
					}
					or {
						description = "description 2"
						property = "storage_class"
						operator = "string_equals"
						value = "cold"
					}
					or {
						description = "description 3"
						property = "firewall.allowed_ip"
						operator = "ips_equals"
						value = jsonencode(["0.0.0.0/0"])
					}
				}
			}
			labels = ["FIXME"]
		}
```
```hcl
resource "ibm_scc_rule" "scc_rule_instance" {
			instance_id = "00000000-1111-2222-3333-444444444444"
			description = "This rule will determine if Activity Tracker is correctly configured"
			version = "0.1.0"
			target {
				service_name = "atracker"
				resource_kind = "target"
				reference_name = "this-target"
				additional_target_attributes {
					name = "type"
					operator = "string_equals"
					value = "cloud_object_storage"
				}
			}
			required_config {
				or {
					property = "route_attached"
					operator = "is_false"
				}
				or {
					any_if {
						target {
							service_name = "cloud-object-storage"
							resource_kind = "bucket"
							additional_target_attributes {
								name = "location"
								operator = "strings_in_list"
								value = "$${this-target}.bucket_name"
							}
						}
						required_config {
							property = "location"
							operator = "strings_in_list"
							value = jsonencode(["us-south","us-east"])
						}
					}
				}
			}
			labels = ["FIXME"]
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
* `required_config` - (Required, List) The specified settings by which your target service will be evaluated against. See the [required_config](#required_config) section for more details.
* `target` - (Required, List) The service or resource used in the assessment. See [rule_target](#rule_target) for more details.
* `version` - (Optional, String) The version number of a rule.
  * Constraints: The maximum length is `10` characters. The minimum length is `5` characters. The value must match regular expression `/^[0-9][0-9.]*$/`.

### required_config
The `required_config` is specified setting by which the target will be evaluated against. 

The `required_config` block supports any of the following schemas:
 - [base_condition](#base_condtion)
 - [list_condition](#list_condition)
 - [sub_rule](#sub_rule)

### base_condition
A base_condition is the basic object in `required_config` block. It details the expected specifications of a service/resource. 

```hcl
required_config {
	property = "location"
	operator = "strings_in_list"
	description = "Must be in the region us-south or us-east"
	value = jsonencode(["us-south","us-east"])
}
```

* `description` - (Optional, String) The details of the expected setting.
* `property` - (Required, String) The property.
	* Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
* `operator` - (Required, String) The operator.
	* Constraints: Allowable values are: 
		* `string_equals`
		* `string_not_equals` 
		* `string_match` 
		* `string_not_match`
		* `string_contains`
		* `string_not_contains`
		* `num_equals`
		* `num_not_equals`
		* `num_less_than`
		* `num_less_than_equals`
		* `num_greater_than`
		* `num_greater_than_equals`
		* `is_empty`
		* `is_not_empty`
		* `is_true`
		* `is_false`
		* `strings_in_list`
		* `strings_allowed`
		* `strings_required`
		* `ips_in_range`
		* `ips_equals`
		* `ips_not_equals`
		* `days_less_than`
		
		The maximum length is `23` characters. The minimum length is `7` characters.
* `value` - (Optional, String) Value of the condition to satisfy. Target/Imports can be used here.

	~> NOTE: If the value requires a list/array, please use the terraform function `jsonencode`
	```hcl
	jsonencode(["us-south","us-east"])
	```

### list_condition
A list_condition is a collection of `and/or` conditons. One or all condtions must be satisfied for the rule to be compliant.
```hcl
required_config {
	...
	and {
		property = "app_lb_pools_with_multiple_members_count"
		operator = "num_not_equals"
		value = "0"
	}
	and {
		property = "app_lb_pools_without_multiple_members_count"
		operator = "num_not_equals"
		value = "0"
	}
}
```
One of the following attributes can be used to denote a list_condition
* `and` - (Optional, List) A list of conditions where all conditions listed must be satisfied. 
* `or` - (Optional, List) A list of conditions where one condition listed must be satisfied.

`and`, `or` supports a combination of the following schemas in the list:
* [base_condition](#base_condition)
* [list_condition](#list_condition)
* [sub_rule_condition](#sub_rule)

### sub_rule
A sub_rule is a rule condition within a rule used to evaluate a target.
```hcl
required_config {
	...
	any_if {
		target {
			service_name = "cloud-object-storage"
			resource_kind = "bucket"
		}
		required_config {
			property = "location"
			operator = "strings_in_list"
			value = jsonencode(["us-south","us-east"])
		}
	}
}
```
One of the following attributes can be used to denote a subrule: 
* `all` - (Optional) A subrule where all targets must satisfy the settings specified in the `required_config` argument. If no targets are found during evaluation, the subrule condition will fail.
* `all_if` - (Optional) A subrule where all targets must satisfy the setting specified in the `required_config` argument. If no targets are found, the subrule condition will pass.
* `any` - (Optional) A subrule where one of the targets must satisfy the setting specified in the `required_config` argument. If no targets are found, the subrule condition will fail.
* `any_if` - (Optional) A subrule where one of the targets must satisfy the setting specified in the `required_config` argument. If no targets are found, the subrule condition will fail. 

The arguments for `all`, `all_if`, `any`, `any_if`: 
* `target` - (Required) see the section [rule_target](#rule_target) for more details
* `required_config` - (Required) see the section [required_config](#required_config) for more details

### rule_target
The `rule_target` is the target the rule is evaluating. This target can be a service or a resource.

Nested schema for **rule_target**:
* `additional_target_attributes` - (Optional, List) The list of targets supported properties.
	* Constraints: The maximum length is `99999` items. The minimum length is `0` items.
	Nested schema for **additional_target_attributes**:
	* `name` - (Optional, String) The additional target attribute name.
		* Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
	* `operator` - (Optional, String) The operator.
		* Constraints: Allowable values are:
			- `string_equals`
			- `string_not_equals`
			- `string_match`
			- `string_not_match`
			- `string_contains`
			- `string_not_contains`
			- `num_equals`
			- `num_not_equals`
			- `num_less_than`
			- `num_less_than_equals`
			- `num_greater_than`
			- `num_greater_than_equals`
			- `is_empty`
			- `is_not_empty`
			- `is_true`
			- `is_false`
			- `strings_in_list`
			- `strings_allowed`
			- `strings_required`
			- `ips_in_range`
			- `ips_equals`
			- `ips_not_equals`
			- `days_less_than`
	* `value` - (Optional, String) The value.
		* Constraints: The maximum length is `256` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
* `resource_kind` - (Required, String) The target resource kind.
	* Constraints: The maximum length is `99999` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
* `reference_name` - (Optional, String) The variable that can be used in the `required_config`. 
* `service_display_name` - (Optional, String) The display name of the target service.
	* Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
* `service_name` - (Required, String) The target service name.
	* Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]+/`.

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
