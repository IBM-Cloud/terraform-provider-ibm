---
layout: "ibm"
page_title: "IBM : ibm_project_environment"
description: |-
  Get information about project_environment
subcategory: "Projects"
---

# ibm_project_environment

Provides a read-only data source to retrieve information about a project_environment. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_project_environment" "project_environment" {
	project_environment_id = ibm_project_environment.project_environment_instance.project_environment_id
	project_id = ibm_project_environment.project_environment_instance.project_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `project_environment_id` - (Required, Forces new resource, String) The environment ID.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
* `project_id` - (Required, Forces new resource, String) The unique project ID.
  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the project_environment.
* `created_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ to match the date and time format as specified by RFC 3339.
* `definition` - (List) The environment definition.
Nested schema for **definition**:
	* `authorizations` - (List) The authorization details. It can authorize by using a trusted profile or an API key in Secrets Manager.
	Nested schema for **authorizations**:
		* `api_key` - (String) The IBM Cloud API Key. It can be either raw or pulled from the catalog via a `CRN` or `JSON` blob.
		  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^<>\\x00-\\x1F]*$/`.
		* `method` - (String) The authorization method. It can authorize by using a trusted profile or an API key in Secrets Manager.
		  * Constraints: The maximum length is `256` characters. The minimum length is `7` characters. The value must match regular expression `/^(ref:)[a-zA-Z0-9\\$\\-_\\.+%!\\*'\\(\\),=&?\/ ]+(authorizations\/method)$|^(api_key)$|^(trusted_profile)$/`.
		* `trusted_profile_id` - (String) The trusted profile ID.
		  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^<>\\x00-\\x1F]*$/`.
	* `compliance_profile` - (List) The profile that is required for compliance.
	Nested schema for **compliance_profile**:
		* `attachment_id` - (String) A unique ID for the attachment to a compliance profile.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\/:a-zA-Z0-9\\.\\-]+$/`.
		* `id` - (String) The unique ID for the compliance profile.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\/:a-zA-Z0-9\\.\\-]+$/`.
		* `instance_id` - (String) A unique ID for the instance of a compliance profile.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\/:a-zA-Z0-9\\.\\-]+$/`.
		* `instance_location` - (String) The location of the compliance instance.
		  * Constraints: Allowable values are: `us-south`, `us-east`, `eu-gb`, `eu-de`, `ca-tor`.
		* `profile_name` - (String) The name of the compliance profile.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^<>\\x00-\\x1F]*$/`.
		* `wp_instance_id` - (String) A unique ID for the instance of a Workload Protection.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\/:a-zA-Z0-9\\.\\-]+$/`.
		* `wp_instance_location` - (String) The location of the compliance instance.
		  * Constraints: Allowable values are: `us-south`, `us-east`, `eu-gb`, `eu-de`, `ca-tor`.
		* `wp_instance_name` - (String) The name of the Workload Protection instance.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^<>\\x00-\\x1F]*$/`.
		* `wp_policy_id` - (String) The unique ID for the Workload Protection policy.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\/:a-zA-Z0-9\\.\\-]+$/`.
		* `wp_policy_name` - (String) The name of the Workload Protection policy.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^<>\\x00-\\x1F]*$/`.
		* `wp_zone_id` - (String) A unique ID for the zone to a Workload Protection policy.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\/:a-zA-Z0-9\\.\\-]+$/`.
		* `wp_zone_name` - (String) A unique ID for the zone to a Workload Protection policy.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^<>\\x00-\\x1F]*$/`.
	* `description` - (String) The description of the environment.
	  * Constraints: The default value is `''`. The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
	* `inputs` - (Map) The input variables that are used for configuration definition and environment.
	* `name` - (String) The name of the environment. It's unique within the account across projects and regions.
	  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"<>{}\\x00-\\x1F]+$/`.
* `href` - (String) A Url.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^((http(s)?:\/\/)|\/)[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/:]+$/`.
* `modified_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ to match the date and time format as specified by RFC 3339.
* `project` - (List) The project that is referenced by this resource.
Nested schema for **project**:
	* `crn` - (String) An IBM Cloud resource name that uniquely identifies a resource.
	  * Constraints: The maximum length is `512` characters. The minimum length is `4` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)(crn)[^'"<>{}\\s\\x00-\\x1F]*$/`.
	* `definition` - (List) The definition of the project reference.
	Nested schema for **definition**:
		* `name` - (String) The name of the project.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"<>{}\\x00-\\x1F]+$/`.
	* `href` - (String) A Url.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^((http(s)?:\/\/)|\/)[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/:]+$/`.
	* `id` - (String) The unique ID.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
* `target_account` - (String) The target account ID derived from the authentication block values. The target account exists only if the environment currently has an authorization block.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9.-]+$/`.

