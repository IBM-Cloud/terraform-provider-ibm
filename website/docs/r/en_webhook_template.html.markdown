---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_webhook_template'
description: |-
  Manages Event Notification Webhook Templates.
---

# ibm_en_webhook_template

Create, update, or delete Webhook Template by using IBM Cloud™ Event Notifications.

## Example usage

```terraform
resource "ibm_en_webhook_template" "webhook_template" {
  instance_guid         = ibm_resource_instance.en_terraform_test_resource.guid
  name                  = "Notification Template"
  type                  = "webhook.notification"
  description           = "Webhook Template for Notifications"
      params {
        body="ewoJImJsb2NrcyI6IFsKCQl7CgkJCSJ0eXBlIjogInNlY3Rpb24iLAoJCQkidGV4dCI6IHsKCQkJCSJ0eXBlIjogIm1ya2R3biIsCgkJCQkidGV4dCI6ICJOZXcgUGFpZCBUaW1lIE9mZiByZXF1ZXN0IGZyb20gPGV4YW1wbGUuY29tfEZyZWQgRW5yaXF1ZXo+XG5cbjxodHRwczovL2V4YW1wbGUuY29tfFZpZXcgcmVxdWVzdD4iCgkJCX0KCQl9CgldCn0="
    }
}
```         

## Argument reference

Review the argument reference that you can specify for your resource.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `name` - (Required, String) The Message Template.

- `description` - (Optional, String) The Template description.

- `type` - (Required, String) webhook.notification

- `params` - (Required, List) Payload describing a template configuration

  Nested scheme for **params**:

  - `body` - (Required, String) The Body for Webhook Template in base64 encoded format.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the `webhook_template`.
- `template_id` - (String) The unique identifier of the created Template.
- `subscription_count` - (Integer) Number of subscriptions.
  - Constraints: The minimum value is `0`.
- `subscription_names` - (List) List of subscriptions.
- `updated_at` - (String) Last updated time.

## Import

You can import the `ibm_en_webhook_template` resource by using `id`.

The `id` property can be formed from `instance_guid`, and `template_id` in the following format:

```
<instance_guid>/<template_id>
```

- `instance_guid`: A string. Unique identifier for IBM Cloud Event Notifications instance.

- `template_id`: A string. Unique identifier for Template.

**Example**

```
$ terraform import ibm_en_webhook_template.webhook_template <instance_guid>/<template_id>
```
