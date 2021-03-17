---

subcategory: "Push Notifications"
layout: "ibm"
page_title: "IBM : pn_application_chrome"
description: |-
  Create, Update, and Delete application settings for platform chrome web.
---

# ibm_pn_application_chrome

Provides an application chrome web resource. This allows to configure chrome web platform for push notification.

## Example Usage

```hcl
resource "ibm_pn_application_chrome" "application_chrome" {
  guid = "guid"
  server_key = "server_key"
  web_site_url = "web_site_url"
}
```

## Argument Reference

The following arguments are supported:

- `guid` - (Required, string) The unique guid of the push notifications instance.
- `server_key` - (Required, string) Server key that provides Push Notification service authorized access to Google services that is used for Chrome Web Push.
- `web_site_url` - (Required, string) The URL of the website/web application that should be permitted to subscribe to Web Push.

## Attribute Reference

The following attributes are exported:

- `id` - The unique identifier of the resource application chrome.
