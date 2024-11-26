---
layout: "ibm"
page_title: "IBM : ibm_scc_instance_settings"
description: |-
  Manages scc_instance_settings.
subcategory: "Security and Compliance Center"
---

# ibm_scc_instance_settings

Create, update, and delete scc_instance_settingss with this resource.

~> NOTE: Security Compliance Center is a regional service. Please specify the IBM Cloud Provider attribute `region` to target another region. Else, exporting the environmental variable IBMCLOUD_SCC_API_ENDPOINT will also override which region is being targeted for all ibm providers(ex. `export IBMCLOUD_SCC_API_ENDPOINT=https://eu-es.compliance.cloud.ibm.com`).

## Example Usage

```hcl
resource "ibm_scc_instance_settings" "scc_instance_settings_instance" {
  instance_id = "00000000-1111-2222-3333-444444444444"
  event_notifications {
		instance_crn = "<event_notifications_crn>"
  }
  object_storage {
		instance_crn = "<cloud_object_storage_crn>"
		bucket = "<cloud_object_storage_bucket>"
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `instance_id` - (Required, Forces new resource, String) The ID of the SCC instance in a particular region.
* `event_notifications` - (Optional, List) The Event Notifications settings.
Nested schema for **event_notifications**:
	* `instance_crn` - (Optional, String) The Event Notifications instance CRN.
	  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}|$/`.
	* `source_id` - (Computed, String) The connected Security and Compliance Center instance CRN.
	  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/([A-Za-z0-9]+(:[A-Za-z0-9]+)+)/`.
	* `updated_on` - (Optional, String) The date when the Event Notifications connection was updated.
	* `source_description` - (Optional,Computed, String) The description of the Event Notifications connection source.
	* `source_name` - (Optional,Computed, String) The name of the Event Notifications connection source.
* `object_storage` - (Optional, List) The Cloud Object Storage settings.
Nested schema for **object_storage**:
	* `bucket` - (Optional, String) The connected Cloud Object Storage bucket name.
	  * Constraints: The maximum length is `64` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z]+|/`.
	* `bucket_endpoint` - (Computed, String) The connected Cloud Object Storage bucket endpoint.
	  * Constraints: The maximum length is `512` characters. The minimum length is `1` character. The value must match regular expression `/([A-Za-z0-9-]+)/`.
	* `bucket_location` - (Computed, String) The connected Cloud Object Storage bucket location.
	  * Constraints: The maximum length is `32` characters. The minimum length is `0` characters. The value must match regular expression `/[A-Za-z]+/`.
	* `instance_crn` - (Optional, String) The connected Cloud Object Storage instance CRN.
	  * Constraints: The maximum length is `512` characters. The minimum length is `0` characters. The value must match regular expression `/^crn:v[0-9](:([A-Za-z0-9-._~!$&'()*+,;=@\/]|%[0-9A-Z]{2})*){8}|$/`.
	* `updated_on` - (Computed, String) The date when the bucket connection was updated.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the scc_instance_settings.

## Import

You can import the `ibm_scc_instance_settings` resource by using `instance_id`. The unique identifier of the scc_instance_settings.

# Syntax
```bash
$ terraform import ibm_scc_instance_settings.scc_instance_settings <instance_id>
```

# Example
```bash
$ terraform import ibm_scc_instance_settings.scc_instance_settings 00000000-1111-2222-3333-444444444444
```
