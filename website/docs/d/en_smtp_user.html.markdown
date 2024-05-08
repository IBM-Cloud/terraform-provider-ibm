---
layout: "ibm"
page_title: "IBM : ibm_en_smtp_user"
description: |-
  Get information about en_smtp_user
subcategory: "Event Notifications"
---

# ibm_en_smtp_user

Provides a read-only data source to retrieve information about an en_smtp_user. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_en_smtp_user" "en_smtp_user" {
	en_smtp_config_id = "en_smtp_user_id"
	instance_id = ibm_en_smtp_user.en_smtp_user_instance.instance_id
	user_id = ibm_en_smtp_user.en_smtp_user_instance.user_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `en_smtp_config_id` - (Required, Forces new resource, String) Unique identifier for SMTP.
  * Constraints: The maximum length is `32` characters. The minimum length is `32` characters. The value must match regular expression `/[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}/`.
* `instance_id` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.
  * Constraints: The maximum length is `256` characters. The minimum length is `10` characters. The value must match regular expression `/[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]/`.
* `user_id` - (Required, Forces new resource, String) UserID.
  * Constraints: The maximum length is `256` characters. The minimum length is `5` characters. The value must match regular expression `/.*/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the en_smtp_user.
* `created_at` - (String) Updated time.

* `description` - (String) SMTP User description.
  * Constraints: The maximum length is `250` characters. The minimum length is `1` character. The value must match regular expression `/[a-zA-Z 0-9-_\/.?:'";,+=!#@$%^&*() ]*/`.

* `domain` - (String) Domain Name.
  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/.*/`.

* `smtp_config_id` - (String) SMTP confg Id.
  * Constraints: The maximum length is `100` characters. The minimum length is `32` characters. The value must match regular expression `/[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}/`.

* `updated_at` - (String) Updated time.

* `username` - (String) SMTP user name.
  * Constraints: The maximum length is `250` characters. The minimum length is `3` characters. The value must match regular expression `/.*/`.

