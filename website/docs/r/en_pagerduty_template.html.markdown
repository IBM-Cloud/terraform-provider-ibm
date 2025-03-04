---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_pagerduty_template'
description: |-
  Manages Event Notification PagerDuty Templates.
---

# ibm_en_pagerduty_template

Create, update, or delete PagerDuty Template by using IBM Cloud™ Event Notifications.

## Example usage

```terraform
resource "ibm_en_pagerduty_template" "pagerduty_template" {
  instance_guid         = ibm_resource_instance.en_terraform_test_resource.guid
  name                  = "Notification Template"
  type                  = "pagerduty.notification"
  description           = "PagerDuty Template for Notifications"
      params {
        body="ewogICJwYXlsb2FkIjogewogICAgInN1bW1hcnkiOiAie3sgZGF0YS5hbGVydF9kZWZpbml0aW9uLm5hbWV9fSIsCiAgICAidGltZXN0YW1wIjogInt7dGltZX19IiwKICAgICJzZXZlcml0eSI6ICJpbmZvIiwKICAgICJzb3VyY2UiOiAie3sgc291cmNlIH19IgogIH0sCiAgImRlZHVwX2tleSI6ICJ7eyBpZCB9fSIsCiAge3sjZXF1YWwgZGF0YS5zdGF0dXMgInRyaWdnZXJlZCJ9fQogICJldmVudF9hY3Rpb24iOiAidHJpZ2dlciIKICAge3svZXF1YWx9fQoKICB7eyNlcXVhbCBkYXRhLnN0YXR1cyAicmVzb2x2ZWQifX0KICAiZXZlbnRfYWN0aW9uIjogInJlc29sdmUiCiAge3svZXF1YWx9fQoKICAge3sjZXF1YWwgZGF0YS5zdGF0dXMgImFja25vd2xlZGdlZCJ9fQogICAiZXZlbnRfYWN0aW9uIjogImFja25vd2xlZGdlIgogICB7ey9lcXVhbH19Cn0="
    }
}
```         

## Argument reference

Review the argument reference that you can specify for your resource.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `name` - (Required, String) The Message Template.

- `description` - (Optional, String) The Template description.

- `type` - (Required, String) pagerduty.notification

- `params` - (Required, List) Payload describing a template configuration

  Nested scheme for **params**:

  - `body` - (Required, String) The Body for PagerDuty Template in base64 encoded format.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the `pagerduty_template`.
- `template_id` - (String) The unique identifier of the created Template.
- `subscription_count` - (Integer) Number of subscriptions.
  - Constraints: The minimum value is `0`.
- `subscription_names` - (List) List of subscriptions.
- `updated_at` - (String) Last updated time.

## Import

You can import the `ibm_en_pagerduty_template` resource by using `id`.

The `id` property can be formed from `instance_guid`, and `template_id` in the following format:

```
<instance_guid>/<template_id>
```

- `instance_guid`: A string. Unique identifier for IBM Cloud Event Notifications instance.

- `template_id`: A string. Unique identifier for Template.

**Example**

```
$ terraform import ibm_en_pagerduty_template.pagerduty_template <instance_guid>/<template_id>
```
