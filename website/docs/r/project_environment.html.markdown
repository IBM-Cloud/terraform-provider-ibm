---
layout: "ibm"
page_title: "IBM : ibm_project_environment"
description: |-
  Manages project_environment.
subcategory: "Projects"
---

# ibm_project_environment

Create, update, and delete project_environments with this resource.

## Example Usage

```hcl
resource "ibm_project_environment" "project_environment" {
  definition {
    name = "environment-stage"
    description = "environment for stage project"
    authorizations {
      method = "api_key"
      api_key = "<your_apikey_here>"
    }
  }
  project_id = ibm_project.project_instance.id
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `definition` - (Required, List) The environment definition.
Nested schema for **definition**:
	* `authorizations` - (Optional, List) The authorization details. You can authorize by using a trusted profile or an API key in Secrets Manager.
	Nested schema for **authorizations**:
		* `api_key` - (Optional, String) The IBM Cloud API Key. It can be either raw or pulled from the catalog via a `CRN` or `JSON` blob.
		  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^<>\\x00-\\x1F]*$/`.
		* `method` - (Optional, String) The authorization method. You can authorize by using a trusted profile or an API key in Secrets Manager.
		  * Constraints: Allowable values are: `api_key`, `trusted_profile`.
		* `trusted_profile_id` - (Optional, String) The trusted profile ID.
		  * Constraints: The maximum length is `512` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
	* `compliance_profile` - (Optional, List) The profile that is required for compliance.
	Nested schema for **compliance_profile**:
		* `attachment_id` - (Optional, String) A unique ID for the attachment to a compliance profile.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `id` - (Optional, String) The unique ID for the compliance profile.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `instance_id` - (Optional, String) A unique ID for the instance of a compliance profile.
		  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
		* `instance_location` - (Optional, String) The location of the compliance instance.
		  * Constraints: Allowable values are: `us-south`, `us-east`, `eu-gb`, `eu-de`, `ca-tor`.
		* `profile_name` - (Optional, String) The name of the compliance profile.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^<>\\x00-\\x1F]*$/`.
	* `description` - (Required, String) The description of the environment.
	  * Constraints: The default value is `''`. The maximum length is `1024` characters. The minimum length is `0` characters. The value must match regular expression `/^$|^(?!\\s)(?!.*\\s$)[^\\x00-\\x1F]*$/`.
	* `inputs` - (Optional, Map) The input variables that are used for configuration definition and environment.
	* `name` - (Required, String) The name of the environment. It's unique within the account across projects and regions.
	  * Constraints: The maximum length is `128` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"<>{}\\x00-\\x1F]+$/`.
* `project_id` - (Required, Forces new resource, String) The unique project ID.
  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the project_environment.
* `created_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ to match the date and time format as specified by RFC 3339.
* `href` - (String) A URL.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(http(s)?:\/\/)[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/]+$/`.
* `modified_at` - (String) A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ to match the date and time format as specified by RFC 3339.
* `project` - (List) The project that is referenced by this resource.
Nested schema for **project**:
	* `crn` - (String) An IBM Cloud resource name that uniquely identifies a resource.
	  * Constraints: The maximum length is `512` characters. The minimum length is `4` characters. The value must match regular expression `/(?!\\s)(?!.*\\s$)^(crn)[^'"<>{}\\s\\x00-\\x1F]*/`.
	* `definition` - (List) The definition of the project reference.
	Nested schema for **definition**:
		* `name` - (String) The name of the project.
		  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$)[^'"<>{}\\x00-\\x1F]+$/`.
	* `href` - (String) A URL.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(http(s)?:\/\/)[a-zA-Z0-9\\$\\-_\\.+!\\*'\\(\\),=&?\/]+$/`.
	* `id` - (String) The unique ID.
	  * Constraints: The maximum length is `128` characters. The value must match regular expression `/^[\\.\\-0-9a-zA-Z]+$/`.
* `project_environment_id` - (String) The environment ID as a friendly name.
  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^(?!\\s)(?!.*\\s$).+$/`.
* `target_account` - (String) The target account ID derived from the authentication block values. The target account exists only if the environment currently has an authorization block.
  * Constraints: The maximum length is `64` characters. The value must match regular expression `/^[a-zA-Z0-9.-]+$/`.


## Import

You can import the `ibm_project_environment` resource by using `id`.
The `id` property can be formed from `project_id`, and `project_environment_id` in the following format:

<pre>
&lt;project_id&gt;/&lt;project_environment_id&gt;
</pre>
* `project_id`: A string. The unique project ID.
* `project_environment_id`: A string. The environment ID as a friendly name.

# Syntax
<pre>
$ terraform import ibm_project_environment.project_environment &lt;project_id&gt;/&lt;project_environment_id&gt;
</pre>
