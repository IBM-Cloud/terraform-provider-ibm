---
layout: "ibm"
page_title: "IBM : ibm_pha_last_operation"
description: |-
  Get information about pha_last_operation
subcategory: "HA and DR Automation for IBM® Power® Virtual Server API reference"
---

# ibm_pha_last_operation

Retrieves the status of the last operation performed on the specified service instance, such as provisioning, updating, or deprovisioning.

## Example Usage

```hcl
data "ibm_pha_last_operation" "pha_last_operation" {
	accept_language = "en-US"
	if_none_match = "abcdef"
	pha_instance_id = "8eefautr-4c02-0009-0086-8bd4d8cf61b6"
}
```
### Path Parameters

* `pha_instance_id` - (Required, Forces new resource, String) The unique identifier of the powerha service instance.
  * Constraints: The maximum length is `1048` characters. The minimum length is `36` characters. The value must match regular expression `/^.*$/`.

### Query Parameters

* `accept_language` - (Optional, Forces new resource, String) The language in which the response should be returned.
* `accepts_incomplete` - (Optional, Forces new resource, Boolean) Indicates whether the request can be accepted before the operation is complete.
* `if_none_match` - (Optional, String) ETag for conditional requests (optional).
  * Constraints: The maximum length is `50` characters. The minimum length is `6` characters. The value must match regular expression `/^[a-zA-Z0-9\\-_,;=.*]+$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the pha_last_operation.
* `deployment_name` - (String) Name of the deployment associated with the service instance.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `href` - (String) Resource reference.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^.*$/`.
* `provision_id` - (String) Unique identifier for the provisioning operation.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `resource_group` - (String) Resource Group.
  * Constraints: The maximum length is `2048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.
* `status` - (String) Current operational status of the service instance.
  * Constraints: The maximum length is `1048` characters. The minimum length is `1` character. The value must match regular expression `/^[A-Za-z0-9._:-]+$/`.

