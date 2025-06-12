---
layout: "ibm"
page_title: "IBM : ibm_cd_toolchain_tool_slack"
description: |-
  Manages cd_toolchain_tool_slack.
subcategory: "Continuous Delivery"
---

# ibm_cd_toolchain_tool_slack

Create, update, and delete cd_toolchain_tool_slacks with this resource.

See the [tool integration](https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-slack) page for more information.

## Example Usage

```hcl
resource "ibm_cd_toolchain_tool_slack" "cd_toolchain_tool_slack_instance" {
  parameters {
		channel_name = "#my_channel"
		pipeline_start = true
		pipeline_success = true
		pipeline_fail = true
		toolchain_bind = true
		toolchain_unbind = true
		webhook = "https://hooks.slack.com/services/A5EWRN5WK/A726ZQWT68G/TsdTjp6q4i6wFQTICTasjkE8"
		team_name = "my_team"
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
	* `channel_name` - (Required, String) The Slack channel that notifications will be posted to.
	* `pipeline_fail` - (Optional, Boolean) Generate `pipeline failed` notifications.
	  * Constraints: The default value is `true`.
	* `pipeline_start` - (Optional, Boolean) Generate `pipeline start` notifications.
	  * Constraints: The default value is `true`.
	* `pipeline_success` - (Optional, Boolean) Generate `pipeline succeeded` notifications.
	  * Constraints: The default value is `true`.
	* `team_name` - (Optional, String) The Slack team name, which is the word or phrase before _.slack.com_ in the team URL.
	* `toolchain_bind` - (Optional, Boolean) Generate `tool added to toolchain` notifications.
	  * Constraints: The default value is `true`.
	* `toolchain_unbind` - (Optional, Boolean) Generate `tool removed from toolchain` notifications.
	  * Constraints: The default value is `true`.
	* `webhook` - (Required, String) The incoming webhook used by Slack to receive events. You can use a toolchain secret reference for this parameter. For more information, see [Protecting your sensitive data in Continuous Delivery](https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-cd_data_security#cd_secure_credentials).
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain to bind the tool to.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the cd_toolchain_tool_slack.
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

You can import the `ibm_cd_toolchain_tool_slack` resource by using `id`.
The `id` property can be formed from `toolchain_id`, and `tool_id` in the following format:

<pre>
&lt;toolchain_id&gt;/&lt;tool_id&gt;
</pre>
* `toolchain_id`: A string. ID of the toolchain to bind the tool to.
* `tool_id`: A string. Tool ID.

# Syntax
<pre>
$ terraform import ibm_cd_toolchain_tool_slack.cd_toolchain_tool_slack &lt;toolchain_id&gt;/&lt;tool_id&gt;
</pre>
