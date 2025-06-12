---
layout: "ibm"
page_title: "IBM : ibm_mqcloud_user"
description: |-
  Get information about mqcloud_user
subcategory: "MQaaS"
---

# ibm_mqcloud_user

Provides a read-only data source to retrieve information about a mqcloud_user. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

> **Note:** The MQaaS Terraform provider access is restricted to users of the reserved deployment, reserved capacity, and reserved capacity subscription plans.

## Example Usage

```hcl
data "ibm_mqcloud_user" "mqcloud_user" {
	name = ibm_mqcloud_user.mqcloud_user_instance.name
	service_instance_guid = ibm_mqcloud_user.mqcloud_user_instance.service_instance_guid
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `name` - (Optional, String) The shortname of the user that will be used as the IBM MQ administrator in interactions with a queue manager for this service instance.
  * Constraints: The maximum length is `12` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][-a-z0-9]*$/`.
* `service_instance_guid` - (Required, Forces new resource, String) The GUID that uniquely identifies the MQaaS service instance.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the mqcloud_user.
* `users` - (List) List of users.
  * Constraints: The maximum length is `50` items. The minimum length is `0` items.
Nested schema for **users**:
	* `email` - (String) The email of the user.
	  * Constraints: The maximum length is `253` characters. The minimum length is `5` characters.
	* `href` - (String) The URL for the user details.
	* `id` - (String) The ID of the user which was allocated on creation, and can be used for delete calls.
	* `name` - (String) The shortname of the user that will be used as the IBM MQ administrator in interactions with a queue manager for this service instance.
	  * Constraints: The maximum length is `12` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][-a-z0-9]*$/`.

