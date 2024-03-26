---
layout: "ibm"
page_title: "IBM : ibm_scc_instance_settings"
description: |-
  Manages scc_instance_settings.
subcategory: "Security and Compliance Center"
---

# ibm_scc_instance_settings

Provides a read-only data source to retrieve information about scc_instance_settings. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

~> NOTE: if you specify the `region` in the provider, that region will become the default URL. Else, exporting the environmental variable IBMCLOUD_SCC_API_ENDPOINT will override any URL(ex. `export IBMCLOUD_SCC_API_ENDPOINT=https://us-south.compliance.cloud.ibm.com`).

## Example Usage

```hcl
data "ibm_scc_instance_settings" "scc_instance_settings_instance" {
  instance_id = "00000000-1111-2222-3333-444444444444"
}
```
## Argument Reference


## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `event_notifications` - (List) The Event Notifications settings.
Nested schema for **event_notifications**:
	* `instance_crn` - (String) The Event Notifications instance CRN.
	* `source_id` - (String) The connected Security and Compliance Center instance CRN.
	* `updated_on` - (String) The date when the Event Notifications connection was updated.
* `object_storage` - (List) The Cloud Object Storage settings.
Nested schema for **object_storage**:
	* `bucket` - (String) The connected Cloud Object Storage bucket name.
	* `bucket_endpoint` - (String) The connected Cloud Object Storage bucket endpoint.
	* `bucket_location` - (String) The connected Cloud Object Storage bucket location.
	* `instance_crn` - (String) The connected Cloud Object Storage instance CRN.
	* `updated_on` - (String) The date when the bucket connection was updated.
