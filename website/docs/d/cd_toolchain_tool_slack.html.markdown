---
layout: "ibm"
page_title: "IBM : ibm_cd_toolchain_tool_slack"
description: |-
  Get information about cd_toolchain_tool_slack
subcategory: "Continuous Delivery"
---

# ibm_cd_toolchain_tool_slack

Provides a read-only data source for cd_toolchain_tool_slack. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

See the [tool integration](https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-slack) page for more information.

## Example Usage

```hcl
data "ibm_cd_toolchain_tool_slack" "cd_toolchain_tool_slack" {
	tool_id = "9603dcd4-3c86-44f8-8d0a-9427369878cf"
	toolchain_id = data.ibm_cd_toolchain.cd_toolchain.id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `tool_id` - (Required, Forces new resource, String) ID of the tool bound to the toolchain.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the cd_toolchain_tool_slack.
* `crn` - (String) Tool CRN.

* `href` - (String) URI representing the tool.

* `name` - (String) Tool name.

* `parameters` - (List) Unique key-value pairs representing parameters to be used to create the tool. A list of parameters for each tool integration can be found in the <a href="https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-integrations">Configuring tool integrations page</a>.
Nested scheme for **parameters**:
	* `channel_name` - (String) The Slack channel that notifications will be posted to.
	* `pipeline_fail` - (Boolean) Generate `pipeline failed` notifications.
	  * Constraints: The default value is `true`.
	* `pipeline_start` - (Boolean) Generate `pipeline start` notifications.
	  * Constraints: The default value is `true`.
	* `pipeline_success` - (Boolean) Generate `pipeline succeeded` notifications.
	  * Constraints: The default value is `true`.
	* `team_name` - (String) The Slack team name, which is the word or phrase before _.slack.com_ in the team URL.
	* `toolchain_bind` - (Boolean) Generate `tool added to toolchain` notifications.
	  * Constraints: The default value is `true`.
	* `toolchain_unbind` - (Boolean) Generate `tool removed from toolchain` notifications.
	  * Constraints: The default value is `true`.
	* `webhook` - (String) The incoming webhook used by Slack to receive events. You can use a toolchain secret reference for this parameter. For more information, see [Protecting your sensitive data in Continuous Delivery](https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-cd_data_security#cd_secure_credentials).

* `referent` - (List) Information on URIs to access this resource through the UI or API.
Nested scheme for **referent**:
	* `api_href` - (String) URI representing this resource through an API.
	* `ui_href` - (String) URI representing this resource through the UI.

* `resource_group_id` - (String) Resource group where the tool is located.

* `state` - (String) Current configuration state of the tool.
  * Constraints: Allowable values are: `configured`, `configuring`, `misconfigured`, `unconfigured`.

* `toolchain_crn` - (String) CRN of toolchain which the tool is bound to.


* `updated_at` - (String) Latest tool update timestamp.

