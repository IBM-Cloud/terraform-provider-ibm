---
layout: "ibm"
page_title: "IBM : ibm_sm_custom_credentials_configuration"
description: |-
Get information about CustomCredentialsConfiguration
subcategory: "Secrets Manager"
---

# ibm_sm_custom_credentials_configuration

Provides a read-only data source for a custom credentials secret configuration. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_sm_custom_credentials_configuration" "sm_custom_credentials_configuration_instance" {
	instance_id = ibm_resource_instance.sm_instance.guid
	region = "us-south"
	name = "example-custom-credentials-config"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `instance_id` - (Required, Forces new resource, String) The GUID of the Secrets Manager instance.
* `region` - (Optional, Forces new resource, String) The region of the Secrets Manager instance. If not provided defaults to the region defined in the IBM provider configuration.
* `endpoint_type` - (Optional, String) - The endpoint type. If not provided the endpoint type is determined by the `visibility` argument provided in the provider configuration.
    * Constraints: Allowable values are: `private`, `public`.
* `name` - (Required, String) The name of the custom credentials configuration.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of this data source.
* `api_key_ref` - (String) The IAM credentials secret ID that is used for setting up a custom credentials secret configuration.
* `code_engine` - (List) The parameters required to configure Code Engine.
  Nested scheme for **code_engine**:
  * `project_id` - (String) The Project ID of your Code Engine project used by this custom credentials configuration.
  * `job_name` - (String) The Code Engine job name used by this custom credentials configuration.
  * `region` - (String) The region of the Code Engine project.
* `code_engine_key_ref` - (String) The IAM API key used by the credentials system to access this Secrets Manager instance..
* `created_at` - (String) The date when the configuration was created. The date format follows `RFC 3339`.
* `created_by` - (String) The unique identifier that is associated with the entity that created the configuration.
  * Constraints: The maximum length is `128` characters. The minimum length is `4` characters.
* `schema` - (List) The schema that defines the format of the input and output parameters  (the credentials) of the Code Engine job.
  Nested scheme for **schema**:
  * `parameters` - (List) The schema of the input parameters.
    Nested scheme for **parameters**:
    * `name` - (String) The name of the parameter.
    * `format` - (String) The format of the parameter, for example 'required:true, type:string'.
    * `env_variable_name` - (String) The name of the environment variable associated with the configuration schema parameter.
  * `credentials` - (List) The schema of the credentials.
    Nested scheme for **credentials**:
    * `name` - (String) The name of the credential.
    * `format` - (String) The format of the credential, for example 'required:true, type:string'.
* `task_timeout` - (String) Specifies the maximum allowed time for a Code Engine task to be completed. Consists of a number followed by a time unit, for example "3m". Supported time units are `s` (seconds), `m` (minutes) and 'h' (hours).
* `updated_at` - (String) The date when the configuration was modified. The date format follows `RFC 3339`.
