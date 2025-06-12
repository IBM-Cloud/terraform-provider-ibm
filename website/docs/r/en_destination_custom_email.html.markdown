---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_destination_custom_email'
description: |-
  Manages Event Notification Custom Email destinations.
---

# ibm_en_destination_custom_email

Create, update, or delete a Custom Email destination by using IBM Cloud™ Event Notifications.

## Example usage

```terraform
resource "ibm_en_destination_custom_email" "custom_domain_en_destination" {
  instance_guid         = ibm_resource_instance.en_terraform_test_resource.guid
  name                  = "Custom Email EN Destination"
  type                  = "smtp_custom"
  collect_failed_events = true
  description           = "Destination Custom Email for event notification"
    config {
      params {
        domain  = "mailx.com"
    }
  }
}

**NOTE:**
- To perform the verification for spf and dkim please follow the instructions here: https://cloud.ibm.com/docs/event-notifications?topic=event-notifications-en-destinations-custom-email
- `verification_type` is Custom Email Destination update parameter which can be used to verify the status of verfication depending on the type of verification.

Process To do the Custom Domain Configuration and Verification.

- Select the configure overflow menu for the destination you want to verify.

- Create Sender Policy Framework (SPF), which is used to authenticate the sender of an email. SPF specifies the mail servers that are allowed to send email for your domain.
    - Open your DNS hosting provider for the domain name configured
    - Create a new TXT record with your domain name registerer with the name and value provided in the configure screen for SPF

- Create DomainKeys Identified Mail (DKIM), which allows an organization to take responsibility for transmitting a message by signing it. DKIM allows the receiver to check the email that claimed to have come from a specific domain, is authorized by the owner of that domain.
    - Open your DNS hosting provider for the domain name configured
    - Create a new TXT record with your domain name registerer with the name and value provided in the configure screen for DKIM

- Save the TXT records

- In the destination verify screen, click on Verify buttons for both SPF and DKIM.         

## Argument reference

Review the argument reference that you can specify for your resource.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `name` - (Required, String) The Destintion name.

- `description` - (Optional, String) The Destination description.

- `type` - (Required, String) smtp_custom.

- `collect_failed_events` - (boolean) Toggle switch to enable collect failed event in Cloud Object Storage bucket.

- `config` - (Optional, List) Payload describing a destination configuration.

  Nested scheme for **config**:

  - `params` - (Required, List)

  Nested scheme for **params**:

  - `domain` - (Required, String) The Custom Domain.
  - `spf` - (Optional, List) The SPF attributes.
		Nested schema for **spf**:
			* `txt_name` - (Optional, String) spf text name.
			  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/.*/`.
			* `txt_value` - (Optional, String) spf text value.
			  * Constraints: The maximum length is `500` characters. The minimum length is `1` character. The value must match regular expression `/.*/`.
			* `verification` - (Optional, String) spf verification.
			  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/.*/`.
  - `dkim` - (Optional, List) The DKIM attributes.
		Nested schema for **dkim**:
			* `public_key` - (Optional, String) dkim public key.
			  * Constraints: The maximum length is `500` characters. The minimum length is `1` character. The value must match regular expression `/.*/`.
			* `selector` - (Optional, String) dkim selector.
			  * Constraints: The maximum length is `500` characters. The minimum length is `1` character. The value must match regular expression `/.*/`.
			* `verification` - (Optional, String) dkim verification.
			  * Constraints: The maximum length is `255` characters. The minimum length is `1` character. The value must match regular expression `/.*/`.      

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the `custom_domain_en_destination`.
- `destination_id` - (String) The unique identifier of the created destination.
- `subscription_count` - (Integer) Number of subscriptions.
  - Constraints: The minimum value is `0`.
- `subscription_names` - (List) List of subscriptions.
- `updated_at` - (String) Last updated time.

## Import

You can import the `ibm_en_destination_custom_email` resource by using `id`.

The `id` property can be formed from `instance_guid`, and `destination_id` in the following format:

```
<instance_guid>/<destination_id>
```

- `instance_guid`: A string. Unique identifier for IBM Cloud Event Notifications instance.

- `destination_id`: A string. Unique identifier for Destination.

**Example**

```
$ terraform import ibm_en_destination_custom_email.custom_domain_email_en_destination <instance_guid>/<destination_id>
```
