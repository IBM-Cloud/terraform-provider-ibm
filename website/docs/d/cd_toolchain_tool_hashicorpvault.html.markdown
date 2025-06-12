---
layout: "ibm"
page_title: "IBM : ibm_cd_toolchain_tool_hashicorpvault"
description: |-
  Get information about cd_toolchain_tool_hashicorpvault
subcategory: "Continuous Delivery"
---

# ibm_cd_toolchain_tool_hashicorpvault

Provides a read-only data source to retrieve information about a cd_toolchain_tool_hashicorpvault. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

See the [tool integration](https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-hashicorpvault) page for more information.

## Example Usage

```hcl
data "ibm_cd_toolchain_tool_hashicorpvault" "cd_toolchain_tool_hashicorpvault" {
	tool_id = "9603dcd4-3c86-44f8-8d0a-9427369878cf"
	toolchain_id = data.ibm_cd_toolchain.cd_toolchain.id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `tool_id` - (Required, Forces new resource, String) ID of the tool bound to the toolchain.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the cd_toolchain_tool_hashicorpvault.
* `crn` - (String) Tool CRN.

* `href` - (String) URI representing the tool.

* `name` - (String) Name of the tool.
  * Constraints: The maximum length is `128` characters. The minimum length is `0` characters. The value must match regular expression `/^([^\\x00-\\x7F]|[a-zA-Z0-9-._ ])+$/`.

* `parameters` - (List) Unique key-value pairs representing parameters to be used to create the tool. A list of parameters for each tool integration can be found in the <a href="https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-integrations">Configuring tool integrations page</a>.
Nested schema for **parameters**:
	* `authentication_method` - (String) The authentication method for your HashiCorp Vault instance.
	  * Constraints: Allowable values are: `token`, `approle`, `userpass`, `github`.
	* `dashboard_url` - (String) The URL of the HashiCorp Vault server dashboard for this integration. In the graphical UI, this is the dashboard that the browser will navigate to when you click the HashiCorp Vault integration tile.
	* `default_secret` - (String) A default secret name that will be selected or used if no list of secret names are returned from your HashiCorp Vault instance.
	* `name` - (String) The name used to identify this tool integration. Secret references include this name to identify the secrets store where the secrets reside. All secrets store tools integrated into a toolchain should have a unique name to allow secret resolution to function properly.
	* `password` - (String) The authentication password for your HashiCorp Vault instance when using the 'userpass' authentication method. This parameter is ignored for other authentication methods. You can use a toolchain secret reference for this parameter. For more information, see [Protecting your sensitive data in Continuous Delivery](https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-cd_data_security#cd_secure_credentials).
	* `path` - (String) The mount path where your secrets are stored in your HashiCorp Vault instance.
	* `role_id` - (String) The authentication role ID for your HashiCorp Vault instance when using the 'approle' authentication method. This parameter is ignored for other authentication methods. Note, 'role_id' should be treated as a secret and should not be shared in plaintext. You can use a toolchain secret reference for this parameter. For more information, see [Protecting your sensitive data in Continuous Delivery](https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-cd_data_security#cd_secure_credentials).
	* `secret_filter` - (String) A regular expression to filter the list of secret names returned from your HashiCorp Vault instance.
	* `secret_id` - (String) The authentication secret ID for your HashiCorp Vault instance when using the 'approle' authentication method. This parameter is ignored for other authentication methods. You can use a toolchain secret reference for this parameter. For more information, see [Protecting your sensitive data in Continuous Delivery](https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-cd_data_security#cd_secure_credentials).
	* `server_url` - (String) The server URL for your HashiCorp Vault instance.
	* `token` - (String) The authentication token for your HashiCorp Vault instance when using the 'github' and 'token' authentication methods. This parameter is ignored for other authentication methods. You can use a toolchain secret reference for this parameter. For more information, see [Protecting your sensitive data in Continuous Delivery](https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-cd_data_security#cd_secure_credentials).
	* `username` - (String) The authentication username for your HashiCorp Vault instance when using the 'userpass' authentication method. This parameter is ignored for other authentication methods.

* `referent` - (List) Information on URIs to access this resource through the UI or API.
Nested schema for **referent**:
	* `api_href` - (String) URI representing this resource through an API.
	* `ui_href` - (String) URI representing this resource through the UI.

* `resource_group_id` - (String) Resource group where the tool is located.

* `state` - (String) Current configuration state of the tool.
  * Constraints: Allowable values are: `configured`, `configuring`, `misconfigured`, `unconfigured`.

* `toolchain_crn` - (String) CRN of toolchain which the tool is bound to.


* `updated_at` - (String) Latest tool update timestamp.

