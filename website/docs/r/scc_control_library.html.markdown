---
layout: "ibm"
page_title: "IBM : ibm_scc_control_library"
description: |-
  Manages scc_control_library.
subcategory: "Security and Compliance Center"
---

# ibm_scc_control_library

Create, update, and delete control libraries by using this resource.

~> NOTE: Security Compliance Center is a regional service. Please specify the IBM Cloud Provider attribute `region` to target another region. Else, exporting the environmental variable IBMCLOUD_SCC_API_ENDPOINT will also override which region is being targeted for all ibm providers(ex. `export IBMCLOUD_SCC_API_ENDPOINT=https://eu-es.compliance.cloud.ibm.com`).

## Example Usage

```hcl
resource "ibm_scc_control_library" "scc_control_library_instance" {
  instance_id = "00000000-1111-2222-3333-444444444444"
  control_library_description = "control_library_description"
  control_library_name = "control_library_name"
  control_library_type = "predefined"
  controls {
		control_name = "control_name"
		control_id = "1fa45e17-9322-4e6c-bbd6-1c51db08e790"
		control_description = "control_description"
		control_category = "control_category"
		control_parent = "control_parent"
		control_tags = [ "control_tags" ]
		control_specifications {
			control_specification_id = "f3517159-889e-4781-819a-89d89b747c85"
			responsibility = "user"
			component_id = "f3517159-889e-4781-819a-89d89b747c85"
			component_name = "componenet_name"
			environment = "environment"
			control_specification_description = "control_specification_description"
			assessments_count = 1
			assessments {
				assessment_id = "assessment_id"
				assessment_method = "assessment_method"
				assessment_type = "assessment_type"
				assessment_description = "assessment_description"
				parameter_count = 1
				parameters {
					parameter_name = "parameter_name"
					parameter_display_name = "parameter_display_name"
					parameter_type = "string"
				}
			}
		}
		control_docs {
			control_docs_id = "control_docs_id"
			control_docs_type = "control_docs_type"
		}
		control_requirement = true
		status = "enabled"
  }
  version_group_label = "e0923045-f00d-44de-b49b-6f1f0e8033cc"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `instance_id` - (Required, Forces new resource, String) The ID of the SCC instance in a particular region.
* `control_library_description` - (Required, String) The control library description.
  * Constraints: The maximum length is `256` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
* `control_library_name` - (Required, String) The control library name.
  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9_\\s\\-]*$/`.
* `control_library_type` - (Required, String) The control library type. Use `custom` in most cases.
  * Constraints: Allowable values are: `predefined`, `custom`.
* `control_library_version` - (Optional, String) The control library version.
  * Constraints: The maximum length is `64` characters. The minimum length is `5` characters. The value must match regular expression `/^[a-zA-Z0-9_\\-.]*$/`.
* `controls` - (Required, List) The list of controls in a control library.
  * Constraints: The maximum length is `1200` items. The minimum length is `0` items.
Nested schema for **controls**:
	* `control_category` - (Optional, String) The control category.
	  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9_,\\-\\s]*$/`.
	* `control_description` - (Optional, String) The control description.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9_,'"\\s\\-\\[\\]]+$/`.
	* `control_docs` - (Optional, List) The control documentation.
	Nested schema for **control_docs**:
		* `control_docs_id` - (Optional, String) The ID of the control documentation.
		  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
		* `control_docs_type` - (Optional, String) The type of control documentation.
		  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
	* `control_id` - (Optional, String) The control name.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
	* `control_name` - (Optional, String) The ID of the control library that contains the profile.
	  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
	* `control_parent` - (Optional, String) The parent control.
	  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]*/`.
	* `control_requirement` - (Optional, Boolean) Is this a control that can be automated or manually evaluated.
	* `control_specifications` - (Optional, List) The control specifications.
	  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
	Nested schema for **control_specifications**:
		* `assessments` - (Optional, List) The assessments.
		  * Constraints: The maximum length is `10` items. The minimum length is `0` items.
		Nested schema for **assessments**:
			* `assessment_description` - (Optional, String) The assessment description.
			  * Constraints: The maximum length is `256` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9_,'\\s\\-]*$/`.
			* `assessment_id` - (Optional, String) The assessment ID.
			  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
			* `assessment_method` - (Optional, String) The assessment method.
			  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
			* `assessment_type` - (Optional, String) The assessment type.
			  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
			* `parameter_count` - (Optional, Integer) The parameter count.
			* `parameters` - (Optional, List) The parameters.
			  * Constraints: The maximum length is `512` items. The minimum length is `0` items.
			Nested schema for **parameters**:
				* `parameter_display_name` - (Optional, String) The parameter display name.
				  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9_,'\\s\\-]*$/`.
				* `parameter_name` - (Optional, String) The parameter name.
				  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9_\\s\\-]*$/`.
				* `parameter_type` - (Optional, String) The parameter type.
				  * Constraints: Allowable values are: `string`, `numeric`, `general`, `boolean`, `string_list`, `ip_list`, `timestamp`.
		* `assessments_count` - (Optional, Integer) The number of assessments.
		* `component_name` - (Optional, String) The component name.
		  * Constraints: The maximum length is `512` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
		* `component_id` - (Optional, String) The component ID.
		  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9-]*$/`.
		* `control_specification_description` - (Optional, String) The control specifications description.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9_,'\\s\\-]*$/`.
		* `control_specification_id` - (Optional, String) The control specification ID.
		  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
		* `environment` - (Optional, String) The control specifications environment.
		  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9-]*$/`.
		* `responsibility` - (Optional, String) The responsibility for managing the control.
		  * Constraints: Allowable values are: `user`.
	* `control_tags` - (Optional, List) The control tags.
	  * Constraints: The list items must match regular expression `/^[a-zA-Z0-9_,'"\\s\\-\\[\\]]+$/`. The maximum length is `512` items. The minimum length is `0` items.
	* `status` - (Optional, String) The control status. Set to `enabled` to other resources to use this control library, `disabled` otherwise.
	  * Constraints: Allowable values are: `enabled`, `disabled`.

* `latest` - (Optional, Boolean) The latest version of the control library.
* `version_group_label` - (Computed, String) The version group label. This is string is the unique identifier for the current version of the Control Library
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `controls_count` - (Optional, Integer) The number of controls.
* `id` - The unique identifier of the scc_control_library.
* `control_library_id` - (String) The ID that is associated with the created `control_library`
* `account_id` - (String) The account ID.
  * Constraints: The maximum length is `32` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-zA-Z0-9-]*$/`.
* `control_parents_count` - (Integer) The number of parent controls in the control library.
* `created_by` - (String) The user who created the control library.
  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.:,_\\s]*$/`.
* `created_on` - (String) The date when the control library was created.
* `hierarchy_enabled` - (Boolean) The indication of whether hierarchy is enabled for the control library.
* `updated_by` - (String) The user who updated the control library.
  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.:,_\\s]*$/`.
* `updated_on` - (String) The date when the control library was updated.


## Import

You can import the `ibm_scc_control_library` resource by using `id`.
The `id` property can be formed from `instance_id` and `control_library_id` in the following format:
```bash
<instance_id>/<control_library_id>
```
* `instance_id`: A string. The instance ID.
* `control_library_id`: A string. The control library ID.

# Syntax
```bash
$ terraform import ibm_scc_control_library.scc_control_library <instance_id>/<control_library_id>
```

# Example
```bash
$ terraform import ibm_scc_control_library.scc_control_library 00000000-1111-2222-3333-444444444444/f3517159-889e-4781-819a-89d89b747c85
```
