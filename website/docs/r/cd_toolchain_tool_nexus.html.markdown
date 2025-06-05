---
layout: "ibm"
page_title: "IBM : ibm_cd_toolchain_tool_nexus"
description: |-
  Manages cd_toolchain_tool_nexus.
subcategory: "Continuous Delivery"
---

# ibm_cd_toolchain_tool_nexus

Create, update, and delete cd_toolchain_tool_nexuss with this resource.

See the [tool integration](https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-nexus) page for more information.

## Example Usage

```hcl
resource "ibm_cd_toolchain_tool_nexus" "cd_toolchain_tool_nexus_instance" {
  parameters {
		name = "my-nexus"
		type = "npm"
		user_id = "<user_id>"
		token = "<token>"
		server_url = "https://my.nexus.server.com/"
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
	* `mirror_url` - (Optional, String) The URL of the Nexus virtual repository, which is a repository that can see your private repositories and is a cache of the public repositories.
	* `name` - (Required, String) The name for this tool integration.
	* `release_url` - (Optional, String) The URL of the Nexus release repository.
	* `server_url` - (Optional, String) The URL of the Nexus server.
	* `snapshot_url` - (Optional, String) The URL of the Nexus snapshot repository.
	* `token` - (Optional, String) The password or token for authenticating to the Nexus repository. You can use a toolchain secret reference for this parameter. For more information, see [Protecting your sensitive data in Continuous Delivery](https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-cd_data_security#cd_secure_credentials).
	* `type` - (Required, String) The type of repository for the Nexus integration.
	  * Constraints: Allowable values are: `npm`, `maven`.
	* `user_id` - (Optional, String) The user id or email for authenticating to the Nexus repository.
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain to bind the tool to.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the cd_toolchain_tool_nexus.
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

You can import the `ibm_cd_toolchain_tool_nexus` resource by using `id`.
The `id` property can be formed from `toolchain_id`, and `tool_id` in the following format:

<pre>
&lt;toolchain_id&gt;/&lt;tool_id&gt;
</pre>
* `toolchain_id`: A string. ID of the toolchain to bind the tool to.
* `tool_id`: A string. Tool ID.

# Syntax
<pre>
$ terraform import ibm_cd_toolchain_tool_nexus.cd_toolchain_tool_nexus &lt;toolchain_id&gt;/&lt;tool_id&gt;
</pre>
