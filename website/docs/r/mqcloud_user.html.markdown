---
layout: "ibm"
page_title: "IBM : ibm_mqcloud_user"
description: |-
  Manages mqcloud_user.
subcategory: "MQ on Cloud"
---

# ibm_mqcloud_user

Create, update, and delete mqcloud_users with this resource.

> **Note:** The MQ on Cloud Terraform provider access is restricted to users of the reserved deployment plan.

## Example Usage

```hcl
resource "ibm_mqcloud_user" "mqcloud_user_instance" {
  email = "testuser@ibm.com"
  name = "testuser"
  service_instance_guid = var.service_instance_guid
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `email` - (Required, Forces new resource, String) The email of the user.
  * Constraints: The maximum length is `253` characters. The minimum length is `5` characters.
* `name` - (Required, Forces new resource, String) The shortname of the user that will be used as the IBM MQ administrator in interactions with a queue manager for this service instance.
  * Constraints: The maximum length is `12` characters. The minimum length is `1` character. The value must match regular expression `/^[a-z][-a-z0-9]*$/`.
* `service_instance_guid` - (Required, Forces new resource, String) The GUID that uniquely identifies the MQ on Cloud service instance.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the mqcloud_user.
* `href` - (String) The URL for the user details.
* `user_id` - (String) The ID of the user which was allocated on creation, and can be used for delete calls.


## Import

You can import the `ibm_mqcloud_user` resource by using `id`.
The `id` property can be formed from `service_instance_guid`, and `user_id` in the following format:

<pre>
&lt;service_instance_guid&gt;/&lt;user_id&gt;
</pre>
* `service_instance_guid`: A string in the format `a2b4d4bc-dadb-4637-bcec-9b7d1e723af8`. The GUID that uniquely identifies the MQ on Cloud service instance.
* `user_id`: A string. The ID of the user which was allocated on creation, and can be used for delete calls.

# Syntax
<pre>
$ terraform import ibm_mqcloud_user.mqcloud_user &lt;service_instance_guid&gt;/&lt;user_id&gt;
</pre>
