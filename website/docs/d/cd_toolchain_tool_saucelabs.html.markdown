---
layout: "ibm"
page_title: "IBM : ibm_cd_toolchain_tool_saucelabs"
description: |-
  Get information about cd_toolchain_tool_saucelabs
subcategory: "Continuous Delivery"
---

# ibm_cd_toolchain_tool_saucelabs

Provides a read-only data source to retrieve information about cd_toolchain_tool_saucelabs. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

See the [tool integration](https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-saucelabs) page for more information.

## Example Usage

```hcl
data "ibm_cd_toolchain_tool_saucelabs" "cd_toolchain_tool_saucelabs" {
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

* `id` - The unique identifier of the cd_toolchain_tool_saucelabs.
* `crn` - (String) Tool CRN.

* `href` - (String) URI representing the tool.

* `name` - (String) Name of the tool.
  * Constraints: The maximum length is `128` characters. The minimum length is `0` characters. The value must match regular expression `/^([^\\x00-\\x7F]|[a-zA-Z0-9-._ ])+$/`.

* `parameters` - (List) Unique key-value pairs representing parameters to be used to create the tool. A list of parameters for each tool integration can be found in the <a href="https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-integrations">Configuring tool integrations page</a>.
Nested schema for **parameters**:
	* `access_key` - (String) The access key for the Sauce Labs account. You can use a toolchain secret reference for this parameter. For more information, see [Protecting your sensitive data in Continuous Delivery](https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-cd_data_security#cd_secure_credentials).
	* `username` - (String) The user name for the Sauce Labs account.

* `referent` - (List) Information on URIs to access this resource through the UI or API.
Nested schema for **referent**:
	* `api_href` - (String) URI representing this resource through an API.
	* `ui_href` - (String) URI representing this resource through the UI.

* `resource_group_id` - (String) Resource group where the tool is located.

* `state` - (String) Current configuration state of the tool.
  * Constraints: Allowable values are: `configured`, `configuring`, `misconfigured`, `unconfigured`.

* `toolchain_crn` - (String) CRN of toolchain which the tool is bound to.


* `updated_at` - (String) Latest tool update timestamp.

