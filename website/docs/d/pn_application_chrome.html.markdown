---

subcategory: "Push Notifications"
layout: "ibm"
page_title: "IBM : pn_application_chrome"
description: |-
  Get push notification configuration of platform chrome web.
---

# ibm_pn_application_chrome
Configure push notifications resource for Chrome web platform. For more information, about push notifications for Chrome, see [Chrome applications](https://cloud.ibm.com/docs/mobilepush?topic=mobilepush-push_step_2#push_step_2_chrome-apps).

## Example usage

```terraform
data "pn_application_chrome" "pn_application_chrome" {
	guid = "guid"
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `guid`-  (String)  Required - The unique GUID of the application by using the push service.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id`-  (String) The unique identifier of the applications chrome.
- `server_key`-  (String) Server key that provides push notification service to authorize the access to Google services that is used for Chrome web push.

- `web_site_url` - The URL of the WebSite / WebApp that should be permitted to subscribe to WebPush.
