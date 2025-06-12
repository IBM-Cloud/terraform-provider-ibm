---
layout: "ibm"
page_title: "IBM : ibm_en_smtp_user"
description: |-
  Manages en_smtp_user.
subcategory: "Event Notifications"
---

# ibm_en_smtp_user

Create, update, and delete en_smtp_users with this resource.

## Example Usage

```hcl
resource "ibm_en_smtp_user" "en_smtp_user_instance" {
  instance_id = "instance_id"
  description   = "test-user"
  smtp_config_id = ibm_en_smtp_configuration.tf_smtp_config.en_smtp_configuration_id
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `description` - (Optional, String) SMTP User description.
  * Constraints: The maximum length is `250` characters. The minimum length is `1` character. The value must match regular expression `/[a-zA-Z 0-9-_\/.?:'";,+=!#@$%^&*() ]*/`.
* `instance_id` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.
  * Constraints: The maximum length is `256` characters. The minimum length is `10` characters. The value must match regular expression `/[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the en_smtp_user.
* `created_at` - (String) Updated time.
* `domain` - (String) Domain Name.
  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/.*/`.
* `smtp_config_id` - (String) SMTP confg Id.
  * Constraints: The maximum length is `100` characters. The minimum length is `32` characters. The value must match regular expression `/[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}/`.
* `updated_at` - (String) Updated time.
* `user_id` - (String) Id.
  * Constraints: The maximum length is `100` characters. The minimum length is `32` characters. The value must match regular expression `/[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}/`.
* `username` - (String) SMTP user name.
  * Constraints: The maximum length is `250` characters. The minimum length is `3` characters. The value must match regular expression `/.*/`.
* `password` - (String) SMTP user password.
  * Constraints: The maximum length is `250` characters. The minimum length is `3` characters. The value must match regular expression `/.*/`.  


## Import

You can import the `ibm_en_smtp_user` resource by using `id`.
The `id` property can be formed from `instance_id`, and `user_id` in the following format:

<pre>
&lt;instance_id&gt;/&lt;user_id&gt;
</pre>
* `instance_id`: A string. Unique identifier for IBM Cloud Event Notifications instance.
* `user_id`: A string. Id.

# Syntax
<pre>
$ terraform import ibm_en_smtp_user.en_smtp_user <instance_id>/<user_id>
</pre>
