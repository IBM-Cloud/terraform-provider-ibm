---
layout: "ibm"
page_title: "IBM : ibm_code_engine_secret"
description: |-
  Get information about code_engine_secret
subcategory: "Code Engine"
---

# ibm_code_engine_secret

Provides a read-only data source to retrieve information about a code_engine_secret. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_code_engine_secret" "code_engine_secret" {
	project_id = data.ibm_code_engine_project.code_engine_project.project_id
	name       = "my-secret"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `name` - (Required, Forces new resource, String) The name of your secret.
  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z0-9]([\\-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([\\-a-z0-9]*[a-z0-9])?)*$/`.
* `project_id` - (Required, Forces new resource, String) The ID of the project.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the code_engine_secret.

* `secret_id` - (String) The identifier of the resource.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.

* `created_at` - (String) The timestamp when the resource was created.

* `data` - (Map) Data container that allows to specify config parameters and their values as a key-value map. Each key field must consist of alphanumeric characters, `-`, `_` or `.` and must not exceed a max length of 253 characters. Each value field can consists of any character and must not exceed a max length of 1048576 characters.

* `entity_tag` - (String) The version of the secret instance, which is used to achieve optimistic locking.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^[\\*\\-a-z0-9]+$/`.

* `format` - (Forces new resource, String) Specify the format of the secret.
  * Constraints: Allowable values are: `generic`, `ssh_auth`, `basic_auth`, `tls`, `service_access`, `registry`, `service_operator`, `other`. The value must match regular expression `/^(generic|ssh_auth|basic_auth|tls|service_access|registry|service_operator|other)$/`.

* `href` - (String) When you provision a new secret,  a URL is created identifying the location of the instance.
  * Constraints: The maximum length is `2048` characters. The minimum length is `0` characters. The value must match regular expression `/(([^:\/?#]+):)?(\/\/([^\/?#]*))?([^?#]*)(\\?([^#]*))?(#(.*))?$/`.

* `region` - (String) The region of the project the resource is located in. Possible values: 'au-syd', 'br-sao', 'ca-tor', 'eu-de', 'eu-gb', 'jp-osa', 'jp-tok', 'us-east', 'us-south'.

* `resource_type` - (String) The type of the secret.

* `service_access` - (Forces new resource, List) Properties for Service Access Secrets.
Nested schema for **service_access**:
	* `resource_key` - (List) The service credential associated with the secret.
	Nested schema for **resource_key**:
		* `id` - (String) ID of the service credential associated with the secret.
		  * Constraints: The maximum length is `36` characters. The minimum length is `0` characters. The value must match regular expression `/^[a-z0-9][\\-a-z0-9]*[a-z0-9]$/`.
		* `name` - (String) Name of the service credential associated with the secret.
	* `role` - (List) A reference to the Role and Role CRN for service binding.
	Nested schema for **role**:
		* `crn` - (String) CRN of the IAM Role for this service access secret.
		  * Constraints: The maximum length is `253` characters. The minimum length is `0` characters. The value must match regular expression `/^[A-Z][a-zA-Z() ]*[a-z)]$|^crn\\:v1\\:[a-zA-Z0-9]*\\:(public|dedicated|local)\\:[\\-a-z0-9]*\\:([a-z][\\-a-z0-9_]*[a-z0-9])?\\:((a|o|s)\/[\\-a-z0-9]+)?\\:[\\-a-z0-9\/]*\\:[\\-a-zA-Z0-9]*(\\:[\\-a-zA-Z0-9\/.]*)?$/`.
		* `name` - (String) Role of the service credential.
		  * Constraints: The default value is `Writer`.
	* `service_instance` - (List) The IBM Cloud service instance associated with the secret.
	Nested schema for **service_instance**:
		* `id` - (String) ID of the IBM Cloud service instance associated with the secret.
		  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-z0-9][\\-a-z0-9]*[a-z0-9]$/`.
		* `type` - (String) Type of IBM Cloud service associated with the secret.
		  * Constraints: The maximum length is `253` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
	* `serviceid` - (List) A reference to a Service ID.
	Nested schema for **serviceid**:
		* `crn` - (String) CRN value of a Service ID.
		  * Constraints: The maximum length is `253` characters. The minimum length is `0` characters. The value must match regular expression `/^crn\\:v1\\:[a-zA-Z0-9]*\\:(public|dedicated|local)\\:[\\-a-z0-9]*\\:([a-z][\\-a-z0-9_]*[a-z0-9])?\\:((a|o|s)\/[\\-a-z0-9]+)?\\:[\\-a-z0-9\/]*\\:[\\-a-zA-Z0-9]*(\\:[\\-a-zA-Z0-9\/.]*)?$/`.
		* `id` - (String) The ID of the Service ID.
		  * Constraints: The maximum length is `46` characters. The minimum length is `46` characters. The value must match regular expression `/^ServiceId-[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.

* `service_operator` - (List) Properties for the IBM Cloud Operator Secret.
Nested schema for **service_operator**:
	* `apikey_id` - (String) The ID of the apikey associated with the operator secret.
	  * Constraints: The maximum length is `43` characters. The minimum length is `43` characters. The value must match regular expression `/^ApiKey-[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.
	* `resource_group_ids` - (List) The list of resource groups (by ID) that the operator secret can bind services in.
	  * Constraints: The list items must match regular expression `/^[a-z0-9]*$/`. The maximum length is `100` items. The minimum length is `0` items.
	* `serviceid` - (List) A reference to a Service ID.
	Nested schema for **serviceid**:
		* `crn` - (String) CRN value of a Service ID.
		  * Constraints: The maximum length is `253` characters. The minimum length is `0` characters. The value must match regular expression `/^crn\\:v1\\:[a-zA-Z0-9]*\\:(public|dedicated|local)\\:[\\-a-z0-9]*\\:([a-z][\\-a-z0-9_]*[a-z0-9])?\\:((a|o|s)\/[\\-a-z0-9]+)?\\:[\\-a-z0-9\/]*\\:[\\-a-zA-Z0-9]*(\\:[\\-a-zA-Z0-9\/.]*)?$/`.
		* `id` - (String) The ID of the Service ID.
		  * Constraints: The maximum length is `46` characters. The minimum length is `46` characters. The value must match regular expression `/^ServiceId-[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$/`.
	* `user_managed` - (Boolean) Specifies whether the operator secret is user managed.

