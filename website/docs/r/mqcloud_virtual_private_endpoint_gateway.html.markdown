---
layout: "ibm"
page_title: "IBM : ibm_mqcloud_virtual_private_endpoint_gateway"
description: |-
  Manages mqcloud_virtual_private_endpoint_gateway.
subcategory: "MQ SaaS"
---

# ibm_mqcloud_virtual_private_endpoint_gateway

Create, update, and delete mqcloud_virtual_private_endpoint_gateways with this resource.

> **Note:** The MQaaS Terraform provider access is restricted to users of the reserved deployment, reserved capacity, and reserved capacity subscription plans.

## Example Usage

```hcl
resource "ibm_mqcloud_virtual_private_endpoint_gateway" "mqcloud_virtual_private_endpoint_gateway_instance" {
  name = "vpe_gateway1-to-vpe_gateway2"
  service_instance_guid = "a2b4d4bc-dadb-4637-bcec-9b7d1e723af8"
  target_crn = "crn:v1:bluemix:public:mqcloud:eu-de:::endpoint:qm1.private.eu-de.mq2.test.appdomain.cloud"
  trusted_profile = "crn:v1:bluemix:public:iam-identity::a/5d5ff2a9001c4055ab1408e9bf185f48::profile:Profile-1c0a8609-ca25-4ad2-a09b-aea472d34afc"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `name` - (Required, Forces new resource, String) The name of the virtual private endpoint gateway, created by the user.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z]|[a-z][-a-z0-9]*[a-z0-9]$/`.
* `service_instance_guid` - (Required, Forces new resource, String) The GUID that uniquely identifies the MQ SaaS service instance.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/`.
* `target_crn` - (Required, Forces new resource, String) The CRN of the reserved capacity service instance the user is trying to connect to.
  * Constraints: The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\/[a-z0-9-]+)?:[a-z0-9-]*:[a-z0-9-]*:[a-zA-Z0-9-_\\.\/]*$|^crn:\\[\\.\\.\\.\\]$/`.
* `trusted_profile` - (Optional, Forces new resource, String) The CRN of the trusted profile to assume for this request. This can only be retrieved using the CLI using `ibmcloud iam tp <profile_id> -o json`.
  * Constraints: The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\/[a-z0-9-]+)?:[a-z0-9-]*:[a-z0-9-]*:[a-zA-Z0-9-_\\.\/]*$|^crn:\\[\\.\\.\\.\\]$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the mqcloud_virtual_private_endpoint_gateway.
* `href` - (String) URL for the details of the virtual private endpoint gateway.
  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
* `status` - (String) The lifecycle state of this virtual privage endpoint.
  * Constraints: The maximum length is `12` characters. The minimum length is `2` characters. The value must match regular expression `/^deleting$|failed$|pending$|stable$|suspended$|updating$|waiting$|approved$/`.
* `virtual_private_endpoint_gateway_guid` - (String) The ID of the virtual private endpoint gateway which was allocated on creation.
  * Constraints: The maximum length is `41` characters. The minimum length is `41` characters. The value must match regular expression `/^[0-9a-zA-Z]{4}-[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/`.


## Import

You can import the `ibm_mqcloud_virtual_private_endpoint_gateway` resource by using `id`.
The `id` property can be formed from `service_instance_guid`, and `virtual_private_endpoint_gateway_guid` in the following format:

<pre>
&lt;service_instance_guid&gt;/&lt;virtual_private_endpoint_gateway_guid&gt;
</pre>
* `service_instance_guid`: A string in the format `a2b4d4bc-dadb-4637-bcec-9b7d1e723af8`. The GUID that uniquely identifies the MQ SaaS service instance.
* `virtual_private_endpoint_gateway_guid`: A string in the format `r010-ebab3c08-c9a8-40c4-8869-61c09ddf7b44`. The ID of the virtual private endpoint gateway which was allocated on creation.

# Syntax
<pre>
$ terraform import ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway &lt;service_instance_guid&gt;/&lt;virtual_private_endpoint_gateway_guid&gt;
</pre>
