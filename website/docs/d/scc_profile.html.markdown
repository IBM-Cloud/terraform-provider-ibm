---
layout: "ibm"
page_title: "IBM : ibm_scc_profile"
description: |-
  Get information about scc_profile
subcategory: "Security and Compliance Center"
---

# ibm_scc_profile

Retrieve information about a profile from a read-only data source. Then, you can reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

~> NOTE: Security Compliance Center is a regional service. Please specify the IBM Cloud Provider attribute `region` to target another region. Else, exporting the environmental variable IBMCLOUD_SCC_API_ENDPOINT will also override which region is being targeted for all ibm providers(ex. `export IBMCLOUD_SCC_API_ENDPOINT=https://eu-es.compliance.cloud.ibm.com`).

## Example Usage

```hcl
data "ibm_scc_profile" "scc_profile" {
    instance_id = "00000000-1111-2222-3333-444444444444"
    profile_id = ibm_scc_profile.scc_profile_instance.profile_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `instance_id` - (Required, Forces new resource, String) The ID of the SCC instance in a particular region.
* `profile_id` - (Required, Forces new resource, String) The profile ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `attachments_count` - (Integer) The number of attachments related to this profile.

* `control_parents_count` - (Integer) The number of parent controls for the profile.

* `controls` - (List) The array of controls that are used to create the profile.
  * Constraints: The maximum length is `600` items. The minimum length is `0` items.
	
	Nested schema for **controls**:
	* `control_category` - (String) The control category.
	  * Constraints: The maximum length is `512` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
	* `control_description` - (String) The control description.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
	* `control_docs` - (List) The control documentation.
	Nested schema for **control_docs**:
		* `control_docs_id` - (String) The ID of the control documentation.
		  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
		* `control_docs_type` - (String) The type of control documentation.
		  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
	* `control_id` - (String) The unique ID of the control library that contains the profile.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/[A-Z0-9]+/`.
	* `control_library_id` - (String) The ID of the control library that contains the profile.
	  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
	* `control_library_version` - (String) The most recent version of the control library.
	  * Constraints: The maximum length is `36` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
	* `control_name` - (String) The control name.
	  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
	* `control_parent` - (String) The parent control.
	  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z0-9]*/`.
	* `control_requirement` - (Boolean) Is this a control that can be automated or manually evaluated.
	* `control_specifications` - (List) The control specifications.
	  * Constraints: The maximum length is `400` items. The minimum length is `0` items.
		
		Nested schema for **control_specifications**:
		* `assessments` - (List) The assessments.
		  * Constraints: The maximum length is `10` items. The minimum length is `0` items.
		Nested schema for **assessments**:
			* `assessment_description` - (String) The assessment description.
			  * Constraints: The maximum length is `256` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9_,'\\s\\-]*$/`.
			* `assessment_id` - (String) The assessment ID.
			  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
			* `assessment_method` - (String) The assessment method.
			  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
			* `assessment_type` - (String) The assessment type.
			  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
			* `parameter_count` - (Integer) The parameter count.
			* `parameters` - (List) The parameters.
			  * Constraints: The maximum length is `512` items. The minimum length is `0` items.
			
				Nested schema for **parameters**:
				* `parameter_display_name` - (String) The parameter display name.
				  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9_,'\\s\\-]*$/`.
				* `parameter_name` - (String) The parameter name.
				  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9_\\s\\-]*$/`.
				* `parameter_type` - (String) The parameter type.
				  * Constraints: Allowable values are: `string`, `numeric`, `general`, `boolean`, `string_list`, `ip_list`, `timestamp`.
		* `assessments_count` - (Integer) The number of assessments.
		* `componenet_name` - (String) The component name.
		  * Constraints: The maximum length is `512` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
		* `component_id` - (String) The component ID.
		  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9-]*$/`.
		* `control_specification_description` - (String) The control specifications description.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9_,'\\s\\-]*$/`.
		* `control_specification_id` - (String) The control specification ID.
		  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
		* `environment` - (String) The control specifications environment.
		  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9-]*$/`.
		* `responsibility` - (String) The responsibility for managing the control.
		  * Constraints: Allowable values are: `user`.
	* `control_specifications_count` - (Integer) The number of control specifications.

* `controls_count` - (Integer) The number of controls for the profile.

* `created_by` - (String) The user who created the profile.
  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.:,_\\s]*$/`.

* `created_on` - (String) The date when the profile was created.

* `default_parameters` - (List) The default parameters of the profile.
  * Constraints: The maximum length is `512` items. The minimum length is `0` items.

	Nested schema for **default_parameters**:
	* `assessment_id` - (String) The implementation ID of the parameter.
	  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
	* `assessment_type` - (String) The type of the implementation.
	  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.
	* `parameter_default_value` - (String) The default value of the parameter.
	  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9_,'"\\s\\-\\[\\]]+$/`.
	* `parameter_display_name` - (String) The parameter display name.
	  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9_,'\\s\\-]*$/`.
	* `parameter_name` - (String) The parameter name.
	  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/^[a-zA-Z0-9_]*$/`.
	* `parameter_type` - (String) The parameter type.
	  * Constraints: Allowable values are: `string`, `numeric`, `general`, `boolean`, `string_list`, `ip_list`, `timestamp`.

* `hierarchy_enabled` - (Boolean) The indication of whether hierarchy is enabled for the profile.

* `id` - (String) The unique ID of the profile.
  * Constraints: The maximum length is `36` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.

* `instance_id` - (String) The instance ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.

* `latest` - (Boolean) The latest version of the profile.

* `profile_description` - (String) The profile description.
  * Constraints: The maximum length is `256` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.

* `profile_name` - (String) The profile name.
  * Constraints: The maximum length is `64` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.

* `profile_type` - (String) The profile type, such as custom or predefined.
  * Constraints: Allowable values are: `predefined`, `custom`.

* `profile_version` - (String) The version status of the profile.
  * Constraints: The maximum length is `64` characters. The minimum length is `5` characters. The value must match regular expression `/^[a-zA-Z0-9_\\-.]*$/`.

* `updated_by` - (String) The user who updated the profile.
  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-\\.:,_\\s]*$/`.

* `updated_on` - (String) The date when the profile was updated.

* `version_group_label` - (String) The version group label of the profile.
  * Constraints: The maximum length is `36` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.

