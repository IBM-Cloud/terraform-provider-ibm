---
layout: "ibm"
page_title: "IBM : sa_notification_channel"
sidebar_current: "docs-ibm-resource-sa-notification-channel"
description: |-
  Manages IBM Cloud Security Advisor Notification Channel.
---

# ibm\_sa\_notification\_channel

This resource is used to order a IBM Cloud Security Advisor Notification Channel.

## Example Usage

```hcl
resource "ibm_sa_notification_channel" "channel" {
  name              = "hello"
  type              = "Webhook"
  endpoint          = "http://cloud.ibm.com"
  enabled           = "true"
  description       = "channel"
  severity          = ["low", "medium"]
  alert_source {
    provider_name = "ALL"
    finding_types = ["ALL"]
  }
}
```

## Argument Reference

The following arguments are supported:

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
