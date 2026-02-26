---
layout: "ibm"
page_title: "IBM : ibm_pha_get_last_operation"
description: |-
  Get information about pha_get_last_operation
subcategory: "PowerhaAutomation Service"
---

# ibm_pha_get_last_operation

Provides a read-only data source to retrieve information about a pha_get_last_operation. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_pha_get_last_operation" "pha_get_last_operation" {
	accept_language = "en-US"
	if_none_match = "abcdef"
	pha_instance_id = "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `accept_language` - (Optional, String) The language requested for the return document.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_,;=.*]+$/`.
* `if_none_match` - (Optional, String) ETag for conditional requests (optional).
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9\\-_,;=.*]+$/`.
* `pha_instance_id` - (Required, Forces new resource, String) instance id of instance to provision.
  * Constraints: The maximum length is `50` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9-]+$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the pha_get_last_operation.
* `deployment_name` - (String) Name of the deployment associated with the service instance.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `provision_id` - (String) Unique identifier for the provisioning operation.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `resource_group` - (String) Resource Group.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `status` - (String) Current operational status of the service instance.
  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.

