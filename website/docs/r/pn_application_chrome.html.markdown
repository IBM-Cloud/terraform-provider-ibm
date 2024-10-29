---

subcategory: "Push Notifications"
layout: "ibm"
page_title: "IBM : pn_application_chrome"
description: |-
  Create, update, and delete application settings for platform chrome web.
---

# ibm_pn_application_chrome
Configure push notifications resource for Chrome web platform. For more information, about push notifications for Chrome, see [for Chrome applications](https://cloud.ibm.com/docs/mobilepush?topic=mobilepush-push_step_2#push_step_2_chrome-apps).

## Example usage

```terraform
resource "ibm_pn_application_chrome" "application_chrome" {
  guid = "guid"
  server_key = "server_key"
  web_site_url = "web_site_url"
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `guid`-  (String)  Required - The unique GUID of the push notifications instance.
- `server_key`-  (String)  Required -  Server key that provides push notification service to authorize the access to Google services that is used for Chrome web push.
- `web_site_url`-  (String)  Required - The URL of the website or web application that should be permitted to subscribe to the web push.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id`-  (String) The unique identifier of the resource application for chrome.
