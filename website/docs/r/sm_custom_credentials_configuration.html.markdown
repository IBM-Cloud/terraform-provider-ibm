---
layout: "ibm"
page_title: "IBM : ibm_sm_custom_credentials_configuration"
description: |-
  Manages IAMCredentialsConfiguration.
subcategory: "Secrets Manager"
---

# ibm_sm_custom_credentials_configuration

A resource for a custom credentials secret configuration. This allows custom credentials secret configuration to be created, updated and deleted. For more information, see the [docs](https://cloud.ibm.com/docs/secrets-manager?topic=secrets-manager-custom-credentials#custom-credentials-config).

## Example Usage

```hcl
resource "ibm_sm_custom_credentials_configuration" "sm_custom_credentials_configuration_instance" {
	instance_id = ibm_resource_instance.sm_instance.guid
	region = "us-south"
	name = "example-custom-credentials-config"
	api_key_ref = ibm_sm_iam_credentials_secret.my_secret_for_custom_credentials.secret_id
	code_engine {
	    project_id = ibm_code_engine_project.my_code_engine_project.project_id
	    job_name = "my_code_engine_job"
	    region = "us-south"
	}
	task_timeout = "10m"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `instance_id` - (Required, Forces new resource, String) The GUID of the Secrets Manager instance.
* `region` - (Optional, Forces new resource, String) The region of the Secrets Manager instance. If not provided defaults to the region defined in the IBM provider configuration.
* `endpoint_type` - (Optional, String) - The endpoint type. If not provided the endpoint type is determined by the `visibility` argument provided in the provider configuration.
    * Constraints: Allowable values are: `private`, `public`.
* `name` - (Required, String) A human-readable unique name to assign to your custom credentials configuration.
* `api_key_ref` - (Optional, Forces new resource, String) The IAM credentials secret ID that is used for setting up a custom credentials secret configuration.
* `code_engine` - (Required, List) The parameters required to configure Code Engine.
  Nested scheme for **code_engine**:
  * `project_id` - (Required, Forces new resource, String) The Project ID of your Code Engine project used by this custom credentials configuration.
  * `job_name` - (Required, Forces new resource, String) The Code Engine job name used by this custom credentials configuration.
  * `region` - (Required, Forces new resource, String) The region of the Code Engine project.
* `task_timeout` - (Required, String) Specifies the maximum allowed time for a Code Engine task to be completed. Consists of a number followed by a time unit, for example "3m". Supported time units are `s` (seconds), `m` (minutes) and 'h' (hours).

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of this resource.
* `code_engine_key_ref` - (String) The IAM API key used by the credentials system to access this Secrets Manager instance..
* `created_at` - (String) The date when the configuration was created. The date format follows `RFC 3339`.
* `created_by` - (String) The unique identifier that is associated with the entity that created the configuration.
  * Constraints: The maximum length is `128` characters. The minimum length is `4` characters.
* `updated_at` - (String) The date when the configuration was modified. The date format follows `RFC 3339`.
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

## Provider Configuration

The IBM Cloud provider offers a flexible means of providing credentials for authentication. The following methods are supported, in this order, and explained below:

- Static credentials
- Environment variables

To find which credentials are required for this resource, see the service table [here](https://cloud.ibm.com/docs/ibm-cloud-provider-for-terraform?topic=ibm-cloud-provider-for-terraform-provider-reference#required-parameters).

### Static credentials

You can provide your static credentials by adding the `ibmcloud_api_key`, `iaas_classic_username`, and `iaas_classic_api_key` arguments in the IBM Cloud provider block.

Usage:
```
provider "ibm" {
    ibmcloud_api_key = ""
    iaas_classic_username = ""
    iaas_classic_api_key = ""
}
```

### Environment variables

You can provide your credentials by exporting the `IC_API_KEY`, `IAAS_CLASSIC_USERNAME`, and `IAAS_CLASSIC_API_KEY` environment variables, representing your IBM Cloud platform API key, IBM Cloud Classic Infrastructure (SoftLayer) user name, and IBM Cloud infrastructure API key, respectively.

```
provider "ibm" {}
```

Usage:
```
export IC_API_KEY="ibmcloud_api_key"
export IAAS_CLASSIC_USERNAME="iaas_classic_username"
export IAAS_CLASSIC_API_KEY="iaas_classic_api_key"
terraform plan
```

Note:

1. Create or find your `ibmcloud_api_key` and `iaas_classic_api_key` [here](https://cloud.ibm.com/iam/apikeys).
  - Select `My IBM Cloud API Keys` option from view dropdown for `ibmcloud_api_key`
  - Select `Classic Infrastructure API Keys` option from view dropdown for `iaas_classic_api_key`
2. For iaas_classic_username
  - Go to [Users](https://cloud.ibm.com/iam/users)
  - Click on user.
  - Find user name in the `VPN password` section under `User Details` tab

For more informaton, see [here](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs#authentication).

## Import

You can import the `ibm_sm_custom_credentials_configuration` resource by using `region`, `instance_id`, and `name`.
For more information, see [the documentation](https://cloud.ibm.com/docs/secrets-manager)

# Syntax
```bash
$ terraform import ibm_sm_custom_credentials_configuration.sm_custom_credentials_configuration <region>/<instance_id>/<name>
```

# Example
```bash
$ terraform import ibm_sm_custom_credentials_configuration.my_config us-east/6ebc4224-e983-496a-8a54-f40a0bfa9175/example-custom-credentials-config
```
