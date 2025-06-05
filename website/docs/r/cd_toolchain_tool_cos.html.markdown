---
layout: "ibm"
page_title: "IBM : ibm_cd_toolchain_tool_cos"
description: |-
  Manages cd_toolchain_tool_cos.
subcategory: "Continuous Delivery"
---

# ibm_cd_toolchain_tool_cos

Create, update, and delete cd_toolchain_tool_coss with this resource.

See the [tool integration](https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-cos_integration) page for more information.

## Example Usage

```hcl
resource "ibm_cd_toolchain_tool_cos" "cd_toolchain_tool_cos_instance" {
  parameters {
		name = "cos_tool_01"
		auth_type = "apikey"
		endpoint = "s3.direct.us-south.cloud-object-storage.appdomain.cloud"
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
	* `auth_type` - (Optional, String) The authentication type. Options are `apikey` IBM Cloud API Key or `hmac` HMAC (Hash Message Authentication Code). The default is `apikey`.
	  * Constraints: Allowable values are: `apikey`, `hmac`.
	* `bucket_name` - (Optional, String) The name of the Cloud Object Storage service bucket.
	  * Constraints: The value must match regular expression `/\\S/`.
	* `cos_api_key` - (Optional, String) The IBM Cloud API key used to access the Cloud Object Storage service. Only relevant when using `apikey` as the `auth_type`.
	* `endpoint` - (Optional, String) The [Cloud Object Storage endpoint](https://cloud.ibm.com/docs/cloud-object-storage?topic=cloud-object-storage-endpoints) in IBM Cloud or other endpoint. For example for IBM Cloud Object Storage: `s3.direct.us-south.cloud-object-storage.appdomain.cloud`.
	  * Constraints: The value must match regular expression `/\\S/`.
	* `hmac_access_key_id` - (Optional, String) The HMAC Access Key ID which is part of an HMAC (Hash Message Authentication Code) credential set. HMAC is identified by a combination of an Access Key ID and a Secret Access Key. Only relevant when `auth_type` is set to `hmac`.
	* `hmac_secret_access_key` - (Optional, String) The HMAC Secret Access Key which is part of an HMAC (Hash Message Authentication Code) credential set. HMAC is identified by a combination of an Access Key ID and a Secret Access Key. Only relevant when `auth_type` is set to `hmac`.
	* `instance_crn` - (Optional, String) The CRN (Cloud Resource Name) of the IBM Cloud Object Storage service instance, only relevant when using `apikey` as the `auth_type`.
	  * Constraints: The value must match regular expression `/^crn:v1:(?:bluemix|staging):public:cloud-object-storage:[a-zA-Z0-9-]*\\b:a\/[0-9a-fA-F]*\\b:[0-9a-fA-F]{8}\\b-[0-9a-fA-F]{4}\\b-[0-9a-fA-F]{4}\\b-[0-9a-fA-F]{4}\\b-[0-9a-fA-F]{12}\\b::$/`.
	* `name` - (Required, String) The name used to identify this tool integration.
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain to bind the tool to.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the cd_toolchain_tool_cos.
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

You can import the `ibm_cd_toolchain_tool_cos` resource by using `id`.
The `id` property can be formed from `toolchain_id`, and `tool_id` in the following format:

<pre>
&lt;toolchain_id&gt;/&lt;tool_id&gt;
</pre>
* `toolchain_id`: A string. ID of the toolchain to bind the tool to.
* `tool_id`: A string. Tool ID.

# Syntax
<pre>
$ terraform import ibm_cd_toolchain_tool_cos.cd_toolchain_tool_cos &lt;toolchain_id&gt;/&lt;tool_id&gt;
</pre>
