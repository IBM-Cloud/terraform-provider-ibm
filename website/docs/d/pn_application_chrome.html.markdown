---

subcategory: "Push Notifications"
layout: "ibm"
page_title: "IBM : pn_application_chrome"
description: |-
  Get push notification configuration of platform chrome web
---

# ibm_pn_application_chrome

Provides a read-only data source for platform chrome web. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "pn_application_chrome" "pn_application_chrome" {
	guid = "guid"
}
```

## Argument Reference

The following arguments are supported:

- `guid` - (Required, string) Unique guid of the application using the push service.

## Attribute Reference

The following attributes are exported:

- `id` - The unique identifier of the applications chrome.
- `server_key` - A server key that gives the push service an authorized access to Google services that is used for Chrome Web Push.

- `web_site_url` - The URL of the WebSite / WebApp that should be permitted to subscribe to WebPush.
