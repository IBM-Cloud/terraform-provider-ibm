---
layout: "ibm"
page_title: "IBM : ibm_cd_toolchain_tool_secretsmanager"
description: |-
  Manages cd_toolchain_tool_secretsmanager.
subcategory: "Continuous Delivery"
---

# ibm_cd_toolchain_tool_secretsmanager

Create, update, and delete cd_toolchain_tool_secretsmanagers with this resource.

See the [tool integration](https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-secretsmanager) page for more information.

## Example Usage

```hcl
resource "ibm_cd_toolchain_tool_secretsmanager" "cd_toolchain_tool_secretsmanager_instance" {
  parameters {
		name = "sm_tool_01"
		instance_id_type = "instance-crn"
		instance_crn = "crn:v1:bluemix:public:secrets-manager:us-south:a/00000000000000000000000000000000:00000000-0000-0000-0000-000000000000::"
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
	* `instance_crn` - (Optional, String) The Secrets Manager service instance CRN (Cloud Resource Name), only relevant when using `instance-crn` as the `instance_id_type`.
	  * Constraints: The value must match regular expression `/^crn:v1:(?:bluemix|staging):public:secrets-manager:[a-zA-Z0-9-]*\\b:a\/[0-9a-fA-F]*\\b:[0-9a-fA-F]{8}\\b-[0-9a-fA-F]{4}\\b-[0-9a-fA-F]{4}\\b-[0-9a-fA-F]{4}\\b-[0-9a-fA-F]{12}\\b::$/`.
	* `instance_id_type` - (Optional, String) The type of service instance identifier. When absent defaults to `instance-name`.
	  * Constraints: Allowable values are: `instance-name`, `instance-crn`.
	* `instance_name` - (Optional, String) The name of the Secrets Manager service instance, only relevant when using `instance-name` as the `instance_id_type`.
	  * Constraints: The value must match regular expression `/\\S/`.
	* `location` - (Optional, String) The IBM Cloud location of the Secrets Manager service instance, only relevant when using `instance-name` as the `instance_id_type`.
	* `name` - (Required, String) The name used to identify this tool integration. Secret references include this name to identify the secrets store where the secrets reside. All secrets store tools integrated into a toolchain should have a unique name to allow secret resolution to function properly.
	* `resource_group_name` - (Optional, String) The name of the resource group where the Secrets Manager service instance is located, only relevant when using `instance-name` as the `instance_id_type`.
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain to bind the tool to.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the cd_toolchain_tool_secretsmanager.
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

You can import the `ibm_cd_toolchain_tool_secretsmanager` resource by using `id`.
The `id` property can be formed from `toolchain_id`, and `tool_id` in the following format:

<pre>
&lt;toolchain_id&gt;/&lt;tool_id&gt;
</pre>
* `toolchain_id`: A string. ID of the toolchain to bind the tool to.
* `tool_id`: A string. Tool ID.

# Syntax
<pre>
$ terraform import ibm_cd_toolchain_tool_secretsmanager.cd_toolchain_tool_secretsmanager &lt;toolchain_id&gt;/&lt;tool_id&gt;
</pre>
