---
layout: "ibm"
page_title: "IBM : pn_application_chrome"
sidebar_current: "docs-ibm-resource-pn-application-chrome"
description: |-
  Manages IBM Push Notificaitions Chrome application settings.
---

# ibm_pn_application_chrome

Provides an application chrome resource. This allows to configure chrome web platform.

## Example Usage

```hcl
resource "ibm_pn_application_chrome" "application_chrome" {
  application_id = "application_id"
  "api_key" = "api_key"
  "web_site_url" = "web_site_url"
}
```

## Argument Reference

The following arguments are supported:

- `application_id` - (Required, string) The guid of the push notifications instance.
- `api_key` - (Required, string) Server key that provides Push Notification service authorized access to Google services that is used for Chrome Web Push.
- `web_site_url` - (Required, string) The URL of the website/web application that should be permitted to subscribe to Web Push.

## Attribute Reference

The following attributes are exported:

- `id` - The unique identifier of the resource application chrome.
- `api_key` - Server key that provides Push Notification service authorized access to Google services that is used for Chrome Web Push.
- `web_site_url` - The URL of the website/web application that should be permitted to subscribe to Web Push.
