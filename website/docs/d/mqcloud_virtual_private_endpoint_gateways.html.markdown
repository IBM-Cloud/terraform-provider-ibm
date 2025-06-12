---
layout: "ibm"
page_title: "IBM : ibm_mqcloud_virtual_private_endpoint_gateways"
description: |-
  Get information about mqcloud_virtual_private_endpoint_gateways
subcategory: "MQaaS"
---

# ibm_mqcloud_virtual_private_endpoint_gateways

Provides a read-only data source to retrieve information about mqcloud_virtual_private_endpoint_gateways. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

> **Note:** The MQaaS Terraform provider access is restricted to users of the reserved deployment, reserved capacity, and reserved capacity subscription plans.

## Example Usage

```hcl
data "ibm_mqcloud_virtual_private_endpoint_gateways" "mqcloud_virtual_private_endpoint_gateways" {
	name = ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance.name
	service_instance_guid = ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance.service_instance_guid
	trusted_profile = ibm_mqcloud_virtual_private_endpoint_gateway.mqcloud_virtual_private_endpoint_gateway_instance.trusted_profile
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `name` - (Optional, Forces new resource, String) The name of the virtual private endpoint gateway, created by the user.
  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z]|[a-z][-a-z0-9]*[a-z0-9]$/`.
* `service_instance_guid` - (Required, Forces new resource, String) The GUID that uniquely identifies the MQaaS service instance.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/`.
* `trusted_profile` - (Optional, String) The CRN of the trusted profile to assume for this request.
  * Constraints: The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\/[a-z0-9-]+)?:[a-z0-9-]*:[a-z0-9-]*:[a-zA-Z0-9-_\\.\/]*$|^crn:\\[\\.\\.\\.\\]$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the mqcloud_virtual_private_endpoint_gateways.
* `virtual_private_endpoint_gateways` - (List) List of virtual private endpoint gateways.
  * Constraints: The maximum length is `50` items. The minimum length is `0` items.
Nested schema for **virtual_private_endpoint_gateways**:
	* `href` - (String) URL for the details of the virtual private endpoint gateway.
	  * Constraints: The maximum length is `8000` characters. The minimum length is `10` characters. The value must match regular expression `/^http(s)?:\/\/([^\/?#]*)([^?#]*)(\\?([^#]*))?(#(.*))?$/`.
	* `id` - (String) The ID of the virtual private endpoint gateway which was allocated on creation.
	  * Constraints: The maximum length is `41` characters. The minimum length is `41` characters. The value must match regular expression `/^[0-9a-zA-Z]{4}-[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/`.
	* `name` - (Forces new resource, String) The name of the virtual private endpoint gateway, created by the user.
	  * Constraints: The maximum length is `63` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z]|[a-z][-a-z0-9]*[a-z0-9]$/`.
	* `status` - (String) The lifecycle state of this virtual privage endpoint.
	  * Constraints: The maximum length is `12` characters. The minimum length is `2` characters. The value must match regular expression `/^deleting$|failed$|pending$|stable$|suspended$|updating$|waiting$/`.
	* `target_crn` - (String) The CRN of the reserved capacity service instance the user is trying to connect to.
	  * Constraints: The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]+:[a-z0-9-]*:([a-z]\/[a-z0-9-]+)?:[a-z0-9-]*:[a-z0-9-]*:[a-zA-Z0-9-_\\.\/]*$|^crn:\\[\\.\\.\\.\\]$/`.

