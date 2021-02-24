---
layout: "ibm"
page_title: "IBM : sa_notification_channel"
sidebar_current: "docs-ibm-datasources-sa-notification-channel"
description: |-
  Manages IBM Cloud Security Advisor Notification Channels.
---

# ibm_sa_notification_channel

Import the details of an existing IBM Cloud Security Advisor Notification channel as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_sa_notification_channel" "channel" {
  channel_id = var.channel_id
}
```

## Argument Reference

The following arguments are supported:

- `channel_id` - (Required, string) The ID of the channel.

## Attribute Reference

The following attributes are exported:

- `channel` - Object of Security Advisor Notification Channel.
  - `name` - The name of the notification channel in the form "v1/{account_id}/notifications/channelName"
  - `description` - A one sentence description of this Channel.
  - `type` - Type of callback URL. Possible values: [ Webhook ]
  - `endpoint` - The callback URL which receives the notification
  - `enabled` - Channel is enabled or not. Default is disabled
  - `frequency` - Default is 0
  - `severity` - Severity of the notification.
    - `critical` - Critical Severity
    - `high` - High Severity
    - `medium` - Medium Severity
    - `low` - Low Severity
  - `alert_source` - List of the alert sources. They identify the providers and their finding types which makes the findings available to Security Advisor
    - `provider_name` - Name of the provide. Possible values: [ VA, NA, ATA, CERT, ALL ]
    - `finding_types` - An array of the finding types of the provider_name or “ALL” to specify all finding types under that provider. Possible values: [ VA, NA, ATA, CERT, config-advisor, ALL ]
