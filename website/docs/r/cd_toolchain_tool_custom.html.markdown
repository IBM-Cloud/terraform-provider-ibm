---
layout: "ibm"
page_title: "IBM : ibm_cd_toolchain_tool_custom"
description: |-
  Manages cd_toolchain_tool_custom.
subcategory: "Continuous Delivery"
---

# ibm_cd_toolchain_tool_custom

Create, update, and delete cd_toolchain_tool_customs with this resource.

See the [tool integration](https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-othertool) page for more information.

## Example Usage

```hcl
resource "ibm_cd_toolchain_tool_custom" "cd_toolchain_tool_custom_instance" {
  parameters {
		type = "Delivery Pipeline"
		lifecycle_phase = "DELIVER"
		name = "My Build and Deploy Pipeline"
		dashboard_url = "https://cloud.ibm.com/devops/pipelines/tekton/ae47390c-9495-4b0b-a489-78464685acdd"
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
	* `additional_properties` - (Optional, String) Any information that is needed to integrate with other tools in the toolchain.
	* `dashboard_url` - (Required, String) The URL of the dashboard for this integration. In the graphical UI, this is the dashboard that the browser will navigate to when you click the integration tile.
	* `description` - (Optional, String) A description outlining the function of this tool.
	* `documentation_url` - (Optional, String) The URL for this tool's documentation.
	* `image_url` - (Optional, String) The URL of the icon shown on the tool integration card in the graphical UI.
	* `lifecycle_phase` - (Required, String) The lifecycle phase of the IBM Cloud Garage Method that is the most closely associated with this tool.
	  * Constraints: Allowable values are: `THINK`, `CODE`, `DELIVER`, `RUN`, `MANAGE`, `LEARN`, `CULTURE`.
	* `name` - (Required, String) The name for this tool integration.
	* `type` - (Required, String) The type of tool that this custom tool is integrating with.
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain to bind the tool to.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the cd_toolchain_tool_custom.
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

You can import the `ibm_cd_toolchain_tool_custom` resource by using `id`.
The `id` property can be formed from `toolchain_id`, and `tool_id` in the following format:

<pre>
&lt;toolchain_id&gt;/&lt;tool_id&gt;
</pre>
* `toolchain_id`: A string. ID of the toolchain to bind the tool to.
* `tool_id`: A string. Tool ID.

# Syntax
<pre>
$ terraform import ibm_cd_toolchain_tool_custom.cd_toolchain_tool_custom &lt;toolchain_id&gt;/&lt;tool_id&gt;
</pre>
