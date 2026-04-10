---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_subscription_custom_email'
description: |-
  Manages Event Notifications Custom Email subscription.
---

# ibm_en_subscription_custom_email

Create, update, or delete a Custom Email subscription by using IBM Cloud™ Event Notifications.

## Example usage

### Subscription to Production Custom Email Destination

```terraform
resource "ibm_en_subscription_custom_email" "production_email_subscription" {
  instance_guid    = ibm_resource_instance.en_terraform_test_resource.guid
  name             = "Production Email Subscription"
  description      = "Subscription for production email notifications"
  destination_id   = ibm_en_destination_custom_email.production_destination.destination_id
  topic_id         = ibm_en_topic.topic1.topic_id
  attributes {
    add_notification_payload = true
    reply_to_mail            = "support@example.com"
    reply_to_name            = "Support Team"
    from_name                = "Production Alerts"
    from_email               = "alerts@example.com"
    invited                  = ["user1@example.com", "user2@example.com"]
  }
}
```

### Subscription to Sandbox Custom Email Destination

```terraform
resource "ibm_en_subscription_custom_email" "sandbox_email_subscription" {
  instance_guid    = ibm_resource_instance.en_terraform_test_resource.guid
  name             = "Sandbox Email Subscription"
  description      = "Subscription for testing email notifications"
  destination_id   = ibm_en_destination_custom_email.sandbox_destination.destination_id
  topic_id         = ibm_en_topic.topic1.topic_id
  attributes {
    add_notification_payload = true
    reply_to_mail            = "test@example.com"
    reply_to_name            = "Test Team"
    invited                  = ["tester1@example.com", "tester2@example.com"]
    # Note: from_name and from_email are NOT required for sandbox destinations
  }
}
```

### Updating Email Subscription (Production)

```terraform
resource "ibm_en_subscription_custom_email" "production_email_subscription" {
  instance_guid    = ibm_resource_instance.en_terraform_test_resource.guid
  name             = "Production Email Subscription Updated"
  description      = "Updated subscription for production email notifications"
  destination_id   = ibm_en_destination_custom_email.production_destination.destination_id
  topic_id         = ibm_en_topic.topic1.topic_id
  attributes {
    add_notification_payload = true
    reply_to_mail            = "support@example.com"
    reply_to_name            = "Support Team"
    from_name                = "Production Alerts"
    from_email               = "alerts@example.com"
    add                      = ["newuser@example.com"]
    remove                   = ["olduser@example.com"]
  }
}
```

## Subscription Attributes Based on Destination Type

The subscription attributes vary depending on whether the destination is a **sandbox** or **production** custom email destination:

### Production Destination Subscriptions
When subscribing to a production custom email destination (`is_sandbox = false`):
- **Required attributes**: `from_name` and `from_email`
- These attributes identify the sender of the email
- `from_email` must belong to the verified custom domain

### Sandbox Destination Subscriptions
When subscribing to a sandbox custom email destination (`is_sandbox = true`):
- **Not required**: `from_name` and `from_email`
- These attributes should be omitted for sandbox subscriptions
- Sandbox subscriptions are for testing and don't require sender verification

### Common Attributes (Both Types)
- `reply_to_mail`: Email address for replies
- `reply_to_name`: Name for reply-to field
- `invited`: List of recipient email addresses
- `add_notification_payload`: Include notification payload in email
- `template_id_notification`: Template for notifications (optional)
- `template_id_invitation`: Template for invitations (optional)

## Argument reference

Review the argument reference that you can specify for your resource.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `name` - (Requires, String) Subscription name.

- `description` - (Optional, String) Subscription description.

- `destination_id` - (Requires, String) Destination ID.

- `topic_id` - (Required, String) Topic ID.

- `attributes` - (Optional, List) Subscription attributes. The required attributes depend on the destination type (sandbox vs production).

  Nested scheme for **attributes**:

  - `reply_to_name` - (Optional, String) The email user name to reply to.

  - `reply_to_mail` - (Optional, String) The email address to reply to.

  - `from_name` - (Conditional, String) The user name from which email is addressed.
    - **Required** for production destinations (`is_sandbox = false`)
    - **Not used** for sandbox destinations (`is_sandbox = true`)

  - `from_email` - (Conditional, String) The email address from which email is addressed. Must belong to the verified custom domain.
    - **Required** for production destinations (`is_sandbox = false`)
    - **Not used** for sandbox destinations (`is_sandbox = true`)

  - `invited` - (Optional, List) The email addresses to send the email to.

  - `add` - (Optional, List) The email addresses to add when updating the list of email addresses.

  - `remove` - (Optional, List) The email addresses to remove from the subscription.

  - `add_notification_payload` - (Optional, Boolean) Whether to include the notification payload in the email. Default is `false`.

  - `template_id_notification` - (Optional, String) The template ID for notification emails.

  - `template_id_invitation` - (Optional, String) The template ID for invitation emails.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the `custom_domain_email_subscription`.

- `subscription_id` - (String) The unique identifier of the created subscription.

- `updated_at` - (String) Last updated time.

## Import

You can import the `ibm_en_subscription_custom_email` resource by using `id`.
The `id` property can be formed from `instance_guid`, and `subscription_id` in the following format:

```
<instance_guid>/<subscription_id>
```

- `instance_guid`: A string. Unique identifier for IBM Cloud Event Notifications instance.
- `subscription_id`: A string. Unique identifier for Subscription.

**Example**

```
$ terraform import ibm_en_subscription_custom_email.custom_domain_email_en_subscription <instance_guid>/<subscription_id>
```
