---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_app_configuration_template'
description: |-
  Manages Event Notification App Configuration Templates.
---

# ibm_en_app_configurations_template

Create, update, or delete App Configuration Template by using IBM Cloud™ Event Notifications.

## Example usage

```terraform
resource "ibm_en_app_configuration_template" "ac_template" {
  instance_guid         = ibm_resource_instance.en_terraform_test_resource.guid
  name                  = "App Configuration Template"
  type                  = "app_configuration.notification"
  description           = "AppConfiguration Template for Notifications"
      params {
        body="eyJlbmFibGVkIjogZmFsc2V9Cg=="
    }
}
```         

## Argument reference

Review the argument reference that you can specify for your resource.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `name` - (Required, String) The Message Template.

- `description` - (Optional, String) The Template description.

- `type` - (Required, String) app_configuration.notification

- `params` - (Required, List) Payload describing a template configuration

  Nested scheme for **params**:

  - `body` - (Required, String) The Body for App Configuration Template in base64 encoded format.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the `ac_template`.
- `template_id` - (String) The unique identifier of the created Template.
- `subscription_count` - (Integer) Number of subscriptions.
  - Constraints: The minimum value is `0`.
- `subscription_names` - (List) List of subscriptions.
- `updated_at` - (String) Last updated time.

## Import

You can import the `ibm_en_app_configuration_template` resource by using `id`.

The `id` property can be formed from `instance_guid`, and `template_id` in the following format:

```
<instance_guid>/<template_id>
```

- `instance_guid`: A string. Unique identifier for IBM Cloud Event Notifications instance.

- `template_id`: A string. Unique identifier for Template.

**Example**

```
$ terraform import ibm_en_app_configuration_template.ac_template <instance_guid>/<template_id>
```
