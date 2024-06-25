---
layout: "ibm"
page_title: "IBM : ibm_en_smtp_configuration"
description: |-
  Get information about en_smtp_configuration
subcategory: "Event Notifications"
---

# ibm_en_smtp_configuration

Provides a read-only data source to retrieve information about an en_smtp_configuration. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_en_smtp_configuration" "en_smtp_configuration" {
	en_smtp_configuration_id = ibm_en_smtp_configuration.en_smtp_configuration_instance.en_smtp_configuration_id
	instance_id = ibm_en_smtp_configuration.en_smtp_configuration_instance.instance_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `en_smtp_configuration_id` - (Required, Forces new resource, String) Unique identifier for SMTP.
  * Constraints: The maximum length is `32` characters. The minimum length is `32` characters. The value must match regular expression `/[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}/`.
* `instance_id` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.
  * Constraints: The maximum length is `256` characters. The minimum length is `10` characters. The value must match regular expression `/[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the en_smtp_configuration.
* `config` - (List) Payload describing a SMTP configuration.
Nested schema for **config**:
	* `dkim` - (List) The DKIM attributes.
	Nested schema for **dkim**:
		* `txt_name` - (String) DMIM text name.
		  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/.*/`.
		* `txt_value` - (String) DMIM text value.
		  * Constraints: The maximum length is `500` characters. The minimum length is `1` character. The value must match regular expression `/.*/`.
		* `verification` - (String) dkim verification.
		  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/.*/`.
	* `en_authorization` - (List) The en_authorization attributes.
	Nested schema for **en_authorization**:
		* `verification` - (String) en_authorization verification.
		  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/.*/`.
	* `spf` - (List) The SPF attributes.
	Nested schema for **spf**:
		* `txt_name` - (String) spf text name.
		  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/.*/`.
		* `txt_value` - (String) spf text value.
		  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/.*/`.
		* `verification` - (String) spf verification.
		  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/.*/`.

* `description` - (String) SMTP description.
  * Constraints: The maximum length is `250` characters. The minimum length is `1` character. The value must match regular expression `/[a-zA-Z 0-9-_\/.?:'";,+=!#@$%^&*() ]*/`.

* `domain` - (String) Domain Name.
  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/.*/`.

* `name` - (String) SMTP name.
  * Constraints: The maximum length is `250` characters. The minimum length is `1` character. The value must match regular expression `/[a-zA-Z 0-9-_\/.?:'";,+=!#@$%^&*() ]*/`.

* `updated_at` - (String) Created time.

