---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_destination_ios'
description: |-
  Manages Event Notifications IOS destination.
---

# ibm_en_destination_ios

Create, update, or delete IOS destination by using IBM Cloud™ Event Notifications.

## Example usage for P8

```terraform
resource "ibm_en_destination_ios" "ios_en_destination" {
  instance_guid = ibm_resource_instance.en_terraform_test_resource.guid
  name        = "IOS Destination Auth"
  type        = "push_ios"
  certificate_content_type = "p8"
  certificate = "${path.module}/Certificates/Auth.p8"
  description = "IOS destination with P8"
  config {
    params {
      cert_type = "p8"
      is_sandbox = true
      key_id = production
      team_id = "2347"
      bundle_id = "testp8"
    }
  }
}
```
## Example usage for P12

```terraform
resource "ibm_en_destination_ios" "ios_en_destination" {
  instance_guid = "ibm_resource_instance.en_terraform_test_resource.guid"
  name        = "IOS Destination "
  type        = "push_ios"
  certificate_content_type = "p12"
  certificate = "${path.module}/Certificates/prod.p12"
  description = "IOS destination with P12"
  config {
    params {
      cert_type = "p12"
      is_sandbox = true
      password = "apnscertpassword"
    }
  }
}
```

## Argument reference

Review the argument reference that you can specify for your resource.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `name` - (Required, String) The Destintion name.

- `description` - (Optional, String) The Destination description.

- `type` - (Required, String) push_ios.

- `certificate_content_type` - (Required, String) The type of certificate, Values are p8/p12.

- `certificate` - (Required, binary) Certificate file. The file type allowed is .p8 and .p12

- `config` - (Required, List) Payload describing a destination configuration.

  Nested scheme for **config**:

  - `params` - (Required, List)

  Nested scheme for **params**:

  - `cert_type` - (Required, String) The Certificate type. Values are p8/p12.

  - `is_sandbox` - (Required, boolean) The flag for sandbox/production environment.

  - `password` - (String) The password string for p12 certificate. Required in case 0f p12.

  - `team_id` - (String) The team_id value in case P8 certificate. Required in case of p8.

  - `key_id` - (String) The team_id value in case P8 certificate. Required in case of p8.

  - `bundle_id` - (String) The team_id value in case P8 certificate. Required in case of p8.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the `ios_en_destination`.
- `destination_id` - (String) The unique identifier of the created destination.
- `subscription_count` - (Integer) Number of subscriptions.
  - Constraints: The minimum value is `0`.
- `subscription_names` - (List) List of subscriptions.
- `updated_at` - (String) Last updated time.

## Import

You can import the `ibm_en_destination_ios` resource by using `id`.

The `id` property can be formed from `instance_guid`, and `destination_id` in the following format:

```
<instance_guid>/<destination_id>
```

- `instance_guid`: A string. Unique identifier for IBM Cloud Event Notifications instance.

- `destination_id`: A string. Unique identifier for Destination.

**Example**

```
$ terraform import ibm_en_destination_ios.ios_en_destination <instance_guid>/<destination_id>
```
