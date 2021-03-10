---
layout: "ibm"
page_title: "IBM : pn_application_chrome"
sidebar_current: "docs-ibm-datasource-pn-application-chrome"
description: |-
  Manages IBM Push Notificaitions Chrome application settings.
---

# ibm_pn_application_chrome

Import the details of an existing IBM push notifications chrome application settings from IBM Cloud as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_pn_application_chrome" "application_chrome" {
  service_instance_guid = "<value of the push notifications instance guid>"
}
```

## Argument Reference

The following arguments are supported:

- `service_instance_guid` - (Required, string) The guid of the push notifications instance.

## Attribute Reference

The following attributes are exported:

- `id` - The unique identifier of the push notifications instance.
- `server_key` - Server key that provides Push Notification service authorized access to Google services that is used for Chrome Web Push.
- `website_url` - The URL of the website/web application that should be permitted to subscribe to Web Push.
