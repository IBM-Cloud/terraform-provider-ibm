---
layout: "ibm"
page_title: "IBM : ibm_cd_toolchain_tool_jenkins"
description: |-
  Manages cd_toolchain_tool_jenkins.
subcategory: "Continuous Delivery"
---

# ibm_cd_toolchain_tool_jenkins

Create, update, and delete cd_toolchain_tool_jenkinss with this resource.

See the [tool integration](https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-jenkins) page for more information.

## Example Usage

```hcl
resource "ibm_cd_toolchain_tool_jenkins" "cd_toolchain_tool_jenkins_instance" {
  parameters {
		name = "jenkins-tool-01"
		dashboard_url = "https://jenkins.mycompany.example.com"
		api_user_name = "<api_user_name>"
		api_token = "<api_token>"
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
	* `api_token` - (Optional, String) The API token to use for Jenkins REST API calls so that DevOps Insights can collect data from Jenkins. You can find the API token on the configuration page of your Jenkins instance. You can use a toolchain secret reference for this parameter. For more information, see [Protecting your sensitive data in Continuous Delivery](https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-cd_data_security#cd_secure_credentials).
	* `api_user_name` - (Optional, String) The user name to use with the Jenkins server's API token, which is required so that DevOps Insights can collect data from Jenkins. You can find your API user name on the configuration page of your Jenkins instance.
	* `dashboard_url` - (Required, String) The URL of the Jenkins server dashboard for this integration. In the graphical UI, this is the dashboard that the browser will navigate to when you click the Jenkins integration tile.
	* `name` - (Required, String) The name for this tool integration.
	* `webhook_url` - (Computed, String) The webhook to use in your Jenkins jobs to send notifications to other tools in your toolchain.
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain to bind the tool to.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the cd_toolchain_tool_jenkins.
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

You can import the `ibm_cd_toolchain_tool_jenkins` resource by using `id`.
The `id` property can be formed from `toolchain_id`, and `tool_id` in the following format:

<pre>
&lt;toolchain_id&gt;/&lt;tool_id&gt;
</pre>
* `toolchain_id`: A string. ID of the toolchain to bind the tool to.
* `tool_id`: A string. Tool ID.

# Syntax
<pre>
$ terraform import ibm_cd_toolchain_tool_jenkins.cd_toolchain_tool_jenkins &lt;toolchain_id&gt;/&lt;tool_id&gt;
</pre>
