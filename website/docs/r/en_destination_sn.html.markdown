---
subcategory: 'Event Notifications'
layout: 'ibm'
page_title: 'IBM : ibm_en_destination_sn'
description: |-
  Manages Event Notification Service Now destinations.
---

# ibm_en_destination_sn

Create, update, or delete a servicenow destination by using IBM Cloud™ Event Notifications.

## Example usage

```terraform
resource "ibm_en_destination_sn" "servicenow_en_destination" {
  instance_guid         = ibm_resource_instance.en_terraform_test_resource.guid
  name                  = "Service Now Destination"
  type                  = "servicenow"
  collect_failed_events = false
  description           = "Destination servicenow for event notification"
  config {
    params {
      client_id  = "321efc4f9c974b03d0fd959a81d8d8e8fyeyee"
      client_secret = "636gdkgvfwepefy9we[]"
      username = "testservicenowcredsuser"
      password = "testservicenowpassword"
      instance_name = "dhivhiifvgfewgewf"
      
    }
  }
}
```

## Argument reference

Review the argument reference that you can specify for your resource.

- `instance_guid` - (Required, Forces new resource, String) Unique identifier for IBM Cloud Event Notifications instance.

- `name` - (Required, String) The Destintion name.

- `description` - (Optional, String) The Destination description.

- `type` - (Required, String) pagerduty.

- `collect_failed_events` - (boolean) Toggle switch to enable collect failed event in Cloud Object Storage bucket.

- `config` - (Optional, List) Payload describing a destination configuration.

  Nested scheme for **config**:

  - `params` - (Required, List)

  Nested scheme for **params**:

  - `client_id` - (Required, string) ClientID for the ServiceNow account oauth.
  - `client_secret` - (Required, string) ClientSecret for the ServiceNow account oauth..
  - `username` - (Required, string) Username for ServiceNow account REST API.
  - `password` - (Required, string) Password for ServiceNow account REST API.
  - `instance_name` - (Required, string) Instance name for ServiceNow account.
## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - (String) The unique identifier of the `servicenow_en_destination`.
- `destination_id` - (String) The unique identifier of the created destination.
- `subscription_count` - (Integer) Number of subscriptions.
  - Constraints: The minimum value is `0`.
- `subscription_names` - (List) List of subscriptions.
- `updated_at` - (String) Last updated time.

## Import

You can import the `ibm_en_destination_sn` resource by using `id`.

The `id` property can be formed from `instance_guid`, and `destination_id` in the following format:

```
<instance_guid>/<destination_id>
```

- `instance_guid`: A string. Unique identifier for IBM Cloud Event Notifications instance.

- `destination_id`: A string. Unique identifier for Destination.

**Example**

```
$ terraform import ibm_en_destination_sn.servicenow_en_destination <instance_guid>/<destination_id>
```
