---
layout: "ibm"
page_title: "IBM : ibm_sm_en_registration"
description: |-
  Get information about NotificationsRegistration
subcategory: "Secrets Manager"
---

# ibm_sm_en_registration

Provides a read-only data source for NotificationsRegistration. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_sm_en_registration" {
  instance_id   = "6ebc4224-e983-496a-8a54-f40a0bfa9175"
  region        = "us-south"
}
```


## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the NotificationsRegistration.
* `event_notifications_instance_crn` - (String) A CRN that uniquely identifies an IBM Cloud resource.
  * Constraints: The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.

