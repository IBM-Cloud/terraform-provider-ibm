---
layout: "ibm"
page_title: "IBM : ibm_sm_instance"
description: |-
  Get information about sm_instance
subcategory: "IBM Cloud Secrets Manager Instance Management API"
---

# ibm_sm_instance

Provides a read-only data source to retrieve information about a sm_instance. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_sm_instance" "sm_instance" {
	instance_id = "60b40daa-1fd3-4f35-a994-2409cc0f270c"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `instance_id` - (Required, Forces new resource, String) The service instance ID.
  * Constraints: Length must be `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the sm_instance.
* `encryption` - (List) Vault encryption configuration for Vault Dedicated instances.
Nested schema for **encryption**:
	* `key_crn` - (String) Vault encryption key CRN (only present for customer_managed mode).
	  * Constraints: The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.
	* `mode` - (String) Vault encryption mode.
	  * Constraints: Allowable values are: `customer_managed`, `service_managed`.
	* `provider` - (String) Vault encryption provider (only present for customer_managed mode). Valid value - 'key_protect'.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z_]+$/`.
* `endpoints` - (List) Instance endpoints for Vault Dedicated instances.
Nested schema for **endpoints**:
	* `private` - (List) Endpoint URLs for accessing the Vault Dedicated instance.
	Nested schema for **private**:
		* `vault_api` - (String) Vault API endpoint URL.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*?$/`.
		* `vault_ui` - (String) Vault UI endpoint URL.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*?$/`.
	* `public` - (List) Endpoint URLs for accessing the Vault Dedicated instance.
	Nested schema for **public**:
		* `vault_api` - (String) Vault API endpoint URL.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*?$/`.
		* `vault_ui` - (String) Vault UI endpoint URL.
		  * Constraints: The maximum length is `4096` characters. The minimum length is `1` character. The value must match regular expression `/^.*?$/`.
* `instance_crn` - (String) The instance CRN identifier.
  * Constraints: The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.
* `plan` - (String) Instance plan name.
  * Constraints: Allowable values are: `dedicated`.
* `vault_cluster` - (List) Vault cluster information for Vault Dedicated instances.
Nested schema for **vault_cluster**:
	* `status` - (String) Vault cluster status. Possible values:- sealed: The Vault cluster is sealed and requires unsealing to access secrets- not_initialized: The Vault cluster has not been initialized yet- healthy: The Vault cluster is operational and ready to serve requests.
	  * Constraints: Allowable values are: `sealed`, `not_initialized`, `healthy`.
	* `version` - (String) Vault cluster version.
	  * Constraints: The maximum length is `256` characters. The minimum length is `1` character. The value must match regular expression `/^[\\w\\.\\+\\-]+$/`.

