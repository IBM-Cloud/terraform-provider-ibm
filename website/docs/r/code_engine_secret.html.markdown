---
layout: "ibm"
page_title: "IBM : ibm_code_engine_secret"
description: |-
  Manages code_engine_secret.
subcategory: "Code Engine"
---

# ibm_code_engine_secret

Create, update, and delete code_engine_secrets with this resource.

## Example Usage

```hcl
resource "ibm_code_engine_secret" "code_engine_secret_instance" {
  project_id = "15314cc3-85b4-4338-903f-c28cdee6d005"
  name = "my-secret"
  format = "generic"

  data = {
		key1 = "value1"
		key2 = "value2"
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `data` - (Optional, Map) Data container that allows to specify config parameters and their values as a key-value map. Each key field must consist of alphanumeric characters, `-`, `_` or `.` and must not exceed a max length of 253 characters. Each value field can consists of any character and must not exceed a max length of 1048576 characters.
* `format` - (Required, Forces new resource, String) Specify the format of the secret.
  * Constraints: Allowable values are: `generic`, `ssh_auth`, `basic_auth`, `tls`, `service_access`, `registry`, `service_operator`, `other`. The value must match regular expression `/^(generic|ssh_auth|basic_auth|tls|service_access|registry|service_operator|other)$/`.
* `name` - (Required, Forces new resource, String) The name of the secret.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([\\-a-z0-9]*[a-z0-9])?)*$/`.
* `project_id` - (Required, Forces new resource, String) The ID of the project.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.
* `service_access` - (Optional, Forces new resource, List) Properties for Service Access Secrets.
Nested schema for **service_access**:
	* `resource_key` - (Required, List) The service credential associated with the secret.
	Nested schema for **resource_key**:
		* `id` - (Optional, String) ID of the service credential associated with the secret.
		  * Constraints: The maximum length is `36` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-z0-9][\\-a-z0-9]*[a-z0-9]$/`.
		* `name` - (Computed, String) Name of the service credential associated with the secret.
	* `role` - (Optional, List) A reference to the Role and Role CRN for service binding.
	Nested schema for **role**:
		* `crn` - (Optional, String) CRN of the IAM Role for this service access secret.
		  * Constraints: The maximum length is `253` characters. The minimum length is `0` characters. The value must match regular expression `/^[A-Z][a-zA-Z() ]*[a-z)]$|^crn\\:v1\\:[a-zA-Z0-9]*\\:(public|dedicated|local)\\:[\\-a-z0-9]*\\:([a-z][\\-a-z0-9_]*[a-z0-9])?\\:((a|o|s)\/[\\-a-z0-9]+)?\\:[\\-a-z0-9\/]*\\:[\\-a-zA-Z0-9]*(\\:[\\-a-zA-Z0-9\/.]*)?$/`.
		* `name` - (Computed, String) Role of the service credential.
		  * Constraints: The default value is `Writer`.
	* `service_instance` - (Required, List) The IBM Cloud service instance associated with the secret.
	Nested schema for **service_instance**:
		* `id` - (Optional, String) ID of the IBM Cloud service instance associated with the secret.
		  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-z0-9][\\-a-z0-9]*[a-z0-9]$/`.
		* `type` - (Computed, String) Type of IBM Cloud service associated with the secret.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
	* `serviceid` - (Optional, List) A reference to a Service ID.
	Nested schema for **serviceid**:
		* `crn` - (Computed, String) CRN value of a Service ID.
		  * Constraints: The maximum length is `253` characters. The minimum length is `0` characters. The value must match regular expression `/^crn\\:v1\\:[a-zA-Z0-9]*\\:(public|dedicated|local)\\:[\\-a-z0-9]*\\:([a-z][\\-a-z0-9_]*[a-z0-9])?\\:((a|o|s)\/[\\-a-z0-9]+)?\\:[\\-a-z0-9\/]*\\:[\\-a-zA-Z0-9]*(\\:[\\-a-zA-Z0-9\/.]*)?$/`.
		* `id` - (Optional, String) The ID of the Service ID.
		  * Constraints: The maximum length is `46` characters. The minimum length is `46` characters. The value must match regular expression `/^ServiceId-[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.
* `service_operator` - (Optional, List) Properties for the IBM Cloud Operator Secret.
Nested schema for **service_operator**:
	* `apikey_id` - (Computed, String) The ID of the apikey associated with the operator secret.
	  * Constraints: The maximum length is `43` characters. The minimum length is `43` characters. The value must match regular expression `/^ApiKey-[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.
	* `resource_group_ids` - (Required, List) The list of resource groups (by ID) that the operator secret can bind services in.
	  * Constraints: The list items must match regular expression `/^[a-z0-9]*$/`. The maximum length is `100` items. The minimum length is `0` items.
	* `serviceid` - (Required, List) A reference to a Service ID.
	Nested schema for **serviceid**:
		* `crn` - (Computed, String) CRN value of a Service ID.
		  * Constraints: The maximum length is `253` characters. The minimum length is `0` characters. The value must match regular expression `/^crn\\:v1\\:[a-zA-Z0-9]*\\:(public|dedicated|local)\\:[\\-a-z0-9]*\\:([a-z][\\-a-z0-9_]*[a-z0-9])?\\:((a|o|s)\/[\\-a-z0-9]+)?\\:[\\-a-z0-9\/]*\\:[\\-a-zA-Z0-9]*(\\:[\\-a-zA-Z0-9\/.]*)?$/`.
		* `id` - (Optional, String) The ID of the Service ID.
		  * Constraints: The maximum length is `46` characters. The minimum length is `46` characters. The value must match regular expression `/^ServiceId-[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.
	* `user_managed` - (Computed, Boolean) Specifies whether the operator secret is user managed.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the code_engine_secret.
* `secret_id` - (String) The identifier of the resource.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.
* `created_at` - (String) The timestamp when the resource was created.
* `entity_tag` - (String) The version of the secret instance, which is used to achieve optimistic locking.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^[\\*\\-a-z0-9]+$/`.
* `href` - (String) When you provision a new secret,  a URL is created identifying the location of the instance.
  * Constraints: The maximum length is `2048` characters. The minimum length is `0` characters. The value must match regular expression `/(([^:\/?#]+):)?(\/\/([^\/?#]*))?([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `region` - (String) The region of the project the resource is located in. Possible values: 'au-syd', 'br-sao', 'ca-tor', 'eu-de', 'eu-gb', 'jp-osa', 'jp-tok', 'us-east', 'us-south'.
* `resource_type` - (String) The type of the secret.
* `etag` - ETag identifier for code_engine_secret.

## Import

You can import the `ibm_code_engine_secret` resource by using `name`.
The `name` property can be formed from `project_id`, and `name` in the following format:

<pre>
&lt;project_id&gt;/&lt;name&gt;
</pre>
* `project_id`: A string in the format `15314cc3-85b4-4338-903f-c28cdee6d005`. The ID of the project.
* `name`: A string in the format `my-secret`. The name of the secret.

# Syntax
<pre>
$ terraform import ibm_code_engine_secret.code_engine_secret &lt;project_id&gt;/&lt;name&gt;
</pre>
