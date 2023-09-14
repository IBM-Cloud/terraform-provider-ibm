---
layout: "ibm"
page_title: "IBM : ibm_sm_en_registration"
description: |-
  Get information about NotificationsRegistration
subcategory: "Secrets Manager"
---

# ibm_sm_en_registration

Provides a read-only data source for event notification registrartion. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_sm_en_registration" "en_registration" {
  instance_id   = ibm_resource_instance.sm_instance.guid
  region        = "us-south"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `instance_id` - (Required, Forces new resource, String) The GUID of the Secrets Manager instance.
* `region` - (Optional, Forces new resource, String) The region of the Secrets Manager instance. If not provided defaults to the region defined in the IBM provider configuration.
* `endpoint_type` - (Optional, String) - The endpoint type. If not provided the endpoint type is determined by the `visibility` argument provided in the provider configuration.
  * Constraints: Allowable values are: `private`, `public`.
  
## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the data source.
* `event_notifications_instance_crn` - (String) A CRN that uniquely identifies an IBM Cloud resource.
  * Constraints: The maximum length is `512` characters. The minimum length is `9` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}$/`.

