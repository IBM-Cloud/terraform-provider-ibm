---
layout: "ibm"
page_title: "IBM : ibm_en_smtp_configuration"
description: |-
  Manages en_smtp_configuration.
subcategory: "Event Notifications"
---

# ibm_en_smtp_configuration

Create, update, and delete en_smtp_configurations with this resource.

## Example Usage

```hcl
resource "ibm_en_smtp_configuration" "en_smtp_configuration_instance" {
  domain = "domain"
  instance_id = "instance_id"
  name = "name"
  description = "SMTP Configuration"
  verification_type = ""
}
```

**NOTE:**
- To perform the verification for spf, dkim and en_authorization please follow the instructions here: https://cloud.ibm.com/docs/event-notifications?topic=event-notifications-en-smtp-configurations#en-smtp-configurations-verify
- `verification_type` is SMTP Configuration update parameter which can be used to verify the status of verfication depending on the type of verification.

## Argument Reference

You can specify the following arguments for this resource.

* `description` - (Optional, String) SMTP description.
  * Constraints: The maximum length is `250` characters. The minimum length is `1` character. The value must match regular expression `/[a-zA-Z 0-9-_\/.?:'";,+=!#@$%^&*() ]*/`.
* `domain` - (Required, String) Domain Name.
  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/.*/`.
* `instance_id` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.
  * Constraints: The maximum length is `256` characters. The minimum length is `10` characters. The value must match regular expression `/[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]/`.
* `name` - (Required, String) SMTP name.
  * Constraints: The maximum length is `250` characters. The minimum length is `1` character. The value must match regular expression `/[a-zA-Z 0-9-_\/.?:'";,+=!#@$%^&*() ]*/`.
* `verification_type` - (Optional, String) The verification_type qualified values are spf/dkim/en_authorization. This is only the update parameter.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

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
* `en_smtp_configuration_id` - (String) SMTP ID.
  * Constraints: The maximum length is `100` characters. The minimum length is `32` characters. The value must match regular expression `/[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}/`.
* `updated_at` - (String) Created time.


## Import

You can import the `ibm_en_smtp_configuration` resource by using `id`.
The `id` property can be formed from `instance_id`, and `en_smtp_configuration_id` in the following format:

<pre>
&lt;instance_id&gt;/&lt;en_smtp_configuration_id&gt;
</pre>
* `instance_id`: A string. Unique identifier for IBM Cloud Event Notifications instance.
* `en_smtp_configuration_id`: A string. SMTP ID.

# Syntax
<pre>
$ terraform import ibm_en_smtp_configuration.en_smtp_configuration <instance_id>/<en_smtp_configuration_id>
</pre>
