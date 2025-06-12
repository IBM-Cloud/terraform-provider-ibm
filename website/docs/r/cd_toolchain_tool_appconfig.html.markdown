---
layout: "ibm"
page_title: "IBM : ibm_cd_toolchain_tool_appconfig"
description: |-
  Manages cd_toolchain_tool_appconfig.
subcategory: "Continuous Delivery"
---

# ibm_cd_toolchain_tool_appconfig

Create, update, and delete cd_toolchain_tool_appconfigs with this resource.

See the [tool integration](https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-app-configuration) page for more information.

## Example Usage

```hcl
resource "ibm_cd_toolchain_tool_appconfig" "cd_toolchain_tool_appconfig_instance" {
  parameters {
		name = "appconfig_tool_01"
		location = "us-south"
		resource_group_name = "Default"
		instance_id = "2a9e3c79-3595-45df-824d-9250aeb598c8"
		environment_id = "environment_01"
		collection_id = "collection_01"
  }
  toolchain_id = ibm_cd_toolchain.cd_toolchain.id
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `name` - (Optional, String) Name of the tool.
  * Constraints: The maximum length is `128` characters. The minimum length is `0` characters. The value must match regular expression `/^([^\\x00-\\x7F]|[a-zA-Z0-9-._ ])+$/`.
* `parameters` - (Required, List) Unique key-value pairs representing parameters to be used to create the tool. A list of parameters for each tool integration can be found in the <a href="https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-integrations">Configuring tool integrations page</a>.
Nested schema for **parameters**:
	* `collection_id` - (Required, String) The ID of the App Configuration collection.
	  * Constraints: The value must match regular expression `/\\S/`.
	* `environment_id` - (Required, String) The ID of the App Configuration environment.
	  * Constraints: The value must match regular expression `/\\S/`.
	* `instance_id` - (Required, String) The guid of the App Configuration service instance.
	  * Constraints: The value must match regular expression `/\\S/`.
	* `location` - (Required, String) The IBM Cloud location where the App Configuration service instance is located.
	* `name` - (Required, String) The name used to identify this tool integration. App Configuration references include this name to identify the App Configuration instance where the configuration values reside. All App Configuration tools integrated into a toolchain should have a unique name to allow resolution to function properly.
	* `resource_group_name` - (Required, String) The name of the resource group where the App Configuration service instance is located.
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain to bind the tool to.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the cd_toolchain_tool_appconfig.
* `crn` - (String) Tool CRN.
* `href` - (String) URI representing the tool.
* `referent` - (List) Information on URIs to access this resource through the UI or API.
Nested schema for **referent**:
	* `api_href` - (String) URI representing this resource through an API.
	* `ui_href` - (String) URI representing this resource through the UI.
* `resource_group_id` - (String) Resource group where the tool is located.
* `state` - (String) Current configuration state of the tool.
  * Constraints: Allowable values are: `configured`, `configuring`, `misconfigured`, `unconfigured`.
* `tool_id` - (String) Tool ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.
* `toolchain_crn` - (String) CRN of toolchain which the tool is bound to.
* `updated_at` - (String) Latest tool update timestamp.


## Import

You can import the `ibm_cd_toolchain_tool_appconfig` resource by using `id`.
The `id` property can be formed from `toolchain_id`, and `tool_id` in the following format:

<pre>
&lt;toolchain_id&gt;/&lt;tool_id&gt;
</pre>
* `toolchain_id`: A string. ID of the toolchain to bind the tool to.
* `tool_id`: A string. Tool ID.

# Syntax
<pre>
$ terraform import ibm_cd_toolchain_tool_appconfig.cd_toolchain_tool_appconfig &lt;toolchain_id&gt;/&lt;tool_id&gt;
</pre>
