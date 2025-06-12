---
layout: "ibm"
page_title: "IBM : ibm_scc_profile"
description: |-
  Manages scc_profile.
subcategory: "Security and Compliance Center"
---

# ibm_scc_profile

Create, update, and delete profiles with this resource.

~> NOTE: Security Compliance Center is a regional service. Please specify the IBM Cloud Provider attribute `region` to target another region. Else, exporting the environmental variable IBMCLOUD_SCC_API_ENDPOINT will also override which region is being targeted for all ibm providers(ex. `export IBMCLOUD_SCC_API_ENDPOINT=https://eu-es.compliance.cloud.ibm.com`).

## Example Usage

```hcl
resource "ibm_scc_profile" "scc_profile_instance" {
  instance_id = "00000000-1111-2222-3333-444444444444"
  controls {
		control_library_id = "e98a56ff-dc24-41d4-9875-1e188e2da6cd"
		control_id = "5C453578-E9A1-421E-AD0F-C6AFCDD67CCF"
		control_library_version = "control_library_version"
		control_name = "control_name"
		control_description = "control_description"
		control_category = "control_category"
		control_parent = "control_parent"
		control_requirement = true
		control_docs {
			control_docs_id = "control_docs_id"
			control_docs_type = "control_docs_type"
		}
		control_specifications_count = 1
		control_specifications {
			control_specification_id = "f3517159-889e-4781-819a-89d89b747c85"
			responsibility = "user"
			component_id = "f3517159-889e-4781-819a-89d89b747c85"
			componenet_name = "componenet_name"
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
  }
  default_parameters {
		assessment_type = "assessment_type"
		assessment_id = "assessment_id"
		parameter_name = "parameter_name"
		parameter_default_value = "parameter_default_value"
		parameter_display_name = "parameter_display_name"
		parameter_type = "string"
  }
  profile_description = "profile_description"
  profile_name = "profile_name"
  profile_type = "predefined"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `instance_id` - (Required, Forces new resource, String) The ID of the SCC instance in a particular region.
* `controls` - (Required, List) The array of controls that are used to create the profile.
  * Constraints: The maximum length is `600` items. The minimum length is `0` items.
Nested schema for **controls**:
	* `control_category` - (Optional, String) The control category.
	  * Constraints: The maximum length is `512` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
	* `control_description` - (Optional, String) The control description.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `2` characters. The value must match regular expression `[A-Za-z0-9]+//`.
	* `control_docs` - (Optional, List) The control documentation.
	Nested schema for **control_docs**:
		* `control_docs_id` - (Optional, String) The ID of the control documentation.
		  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
		* `control_docs_type` - (Optional, String) The type of control documentation.
		  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
	* `control_id` - (Optional, String) The unique ID of the control library that contains the profile.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/[A-Z0-9]+/`.
	* `control_library_id` - (Optional, String) The ID of the control library that contains the profile.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
	* `control_library_version` - (Optional, String) The most recent version of the control library.
	  * Constraints: The maximum length is `36` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
	* `control_name` - (Optional, String) The control name.
	  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
	* `control_parent` - (Optional, String) The parent control.
	  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]*/`.
	* `control_requirement` - (Optional, Boolean) Is this a control that can be automated or manually evaluated.
	* `control_specifications` - (Optional, List) The control specifications.
	  * Constraints: The maximum length is `400` items. The minimum length is `0` items.
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
		* `componenet_name` - (Optional, String) The component name.
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
	* `control_specifications_count` - (Optional, Integer) The number of control specifications.
* `default_parameters` - (Required, List) The default parameters of the profile.
  * Constraints: The maximum length is `512` items. The minimum length is `0` items.
Nested schema for **default_parameters**:
	* `assessment_id` - (Optional, String) The implementation ID of the parameter.
	  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
	* `assessment_type` - (Optional, String) The type of the implementation.
	  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
	* `parameter_default_value` - (Optional, String) The default value of the parameter.
	  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9_,'"\\s\\-\\[\\]]+$/`.
	* `parameter_display_name` - (Optional, String) The parameter display name.
	  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9_,'\\s\\-]*$/`.
	* `parameter_name` - (Optional, String) The parameter name.
	  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9_]*$/`.
	* `parameter_type` - (Optional, String) The parameter type.
	  * Constraints: Allowable values are: `string`, `numeric`, `general`, `boolean`, `string_list`, `ip_list`, `timestamp`.
* `profile_description` - (Required, String) The profile description.
  * Constraints: The maximum length is `256` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
* `profile_name` - (Required, String) The profile name.
  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
* `profile_version` - (Optional, String) The version of the profile to set. The value must match regular expression `/\d+\.\d+\.\d+/`.
* `profile_type` - (Required, String) The profile type, such as custom or predefined.
  * Constraints: Allowable values are: `predefined`, `custom`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the scc_profile.
* `profile_id` - (String) The ID that is associated with the created `profile`
* `attachments_count` - (Integer) The number of attachments related to this profile.
* `control_parents_count` - (Integer) The number of parent controls for the profile.
* `controls_count` - (Integer) The number of controls for the profile.
* `created_by` - (String) The user who created the profile.
  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.:,_\\s]*$/`.
* `created_on` - (String) The date when the profile was created.
* `hierarchy_enabled` - (Boolean) The indication of whether hierarchy is enabled for the profile.
* `instance_id` - (String) The instance ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
* `latest` - (Boolean) The latest version of the profile.
* `profile_version` - (String) The version status of the profile.
  * Constraints: The maximum length is `64` characters. The minimum length is `5` characters. The value must match regular expression `/^[a-zA-Z0-9_\\-.]*$/`.
* `updated_by` - (String) The user who updated the profile.
  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.:,_\\s]*$/`.
* `updated_on` - (String) The date when the profile was updated.
* `version_group_label` - (String) The version group label of the profile.
  * Constraints: The maximum length is `36` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.


## Import

You can import the `ibm_scc_profile` resource by using `id`.
The `id` property can be formed from `instance_id` and `profiles_id` in the following format:

```bash
<instance_id>/<profile_id>
```

* `instance_id`: A string. The instance ID.
* `profile_id`: A string. The profile ID.

# Syntax

```bash
$ terraform import ibm_scc_profile.scc_profile <instance_id>/<profile_id>
```

# Example
```bash
$ terraform import ibm_scc_profile.scc_profile 00000000-1111-2222-3333-444444444444/00000000-1111-2222-3333-444444444444
```