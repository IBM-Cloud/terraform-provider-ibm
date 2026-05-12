---
layout: "ibm"
page_title: "IBM : ibm_notification_distribution_list_destination"
description: |-
  Manages notification_distribution_list_destination.
subcategory: "Platform Notifications"
---

# ibm_notification_distribution_list_destination

**Note - This resource is currently in beta and subject to change without notice.**

Create, update, and delete notification distribution list destinations with this resource.

## Example Usage

```hcl
resource "ibm_notification_distribution_list_destination" "notification_distribution_list_destination_instance" {
  account_id = "your_account_id"
  destination_id = "12345678-1234-1234-1234-123456789012"
  destination_type = "event_notifications"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `account_id` - (Required, Forces new resource, String) The IBM Cloud account ID.
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-zA-Z]{1,100}$/`.
* `destination_id` - (Optional, Forces new resource, String) The GUID of the Event Notifications instance.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/`.
* `destination_type` - (Required, Forces new resource, String) The type of the destination.
  * Constraints: Allowable values are: `event_notifications`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the notification_distribution_list_destination in the format `account_id/destination_id`.

## Import

You can import the `ibm_notification_distribution_list_destination` resource by using `id`. The ID is a combination of the account ID and destination ID.

# Syntax
<pre>
$ terraform import ibm_notification_distribution_list_destination.notification_distribution_list_destination_instance <account_id>/<destination_id>
</pre>

# Example
<pre>
$ terraform import ibm_notification_distribution_list_destination.notification_distribution_list_destination_instance your_account_id/12345678-1234-1234-1234-123456789012
</pre>