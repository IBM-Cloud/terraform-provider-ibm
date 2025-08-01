---
layout: "ibm"
page_title: "IBM : ibm_mqcloud_truststore_certificate"
description: |-
  Get information about mqcloud_truststore_certificate
subcategory: "MQ SaaS"
---

# ibm_mqcloud_truststore_certificate

Provides a read-only data source to retrieve information about a mqcloud_truststore_certificate. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

> **Note:** The MQaaS Terraform provider access is restricted to users of the reserved deployment, reserved capacity, and reserved capacity subscription plans.

## Example Usage

```hcl
data "ibm_mqcloud_truststore_certificate" "mqcloud_truststore_certificate" {
	label = ibm_mqcloud_truststore_certificate.mqcloud_truststore_certificate_instance.label
	queue_manager_id = ibm_mqcloud_truststore_certificate.mqcloud_truststore_certificate_instance.queue_manager_id
	service_instance_guid = ibm_mqcloud_truststore_certificate.mqcloud_truststore_certificate_instance.service_instance_guid
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `label` - (Optional, String) Certificate label in queue manager store.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_.]*$/`.
* `queue_manager_id` - (Required, Forces new resource, String) The id of the queue manager to retrieve its full details.
  * Constraints: The maximum length is `32` characters. The minimum length is `32` characters. The value must match regular expression `/^[0-9a-fA-F]{32}$/`.
* `service_instance_guid` - (Required, Forces new resource, String) The GUID that uniquely identifies the MQ SaaS service instance.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the mqcloud_truststore_certificate.
* `total_count` - (Integer) The total count of trust store certificates.
* `trust_store` - (List) The list of trust store certificates.
  * Constraints: The maximum length is `50` items. The minimum length is `0` items.
Nested schema for **trust_store**:
	* `certificate_type` - (String) The type of certificate.
	  * Constraints: Allowable values are: `trust_store`.
	* `expiry` - (String) Expiry date for the certificate.
	* `fingerprint_sha256` - (String) Fingerprint SHA256.
	  * Constraints: The value must match regular expression `/^[A-F0-9]{2}(:[A-F0-9]{2}){31}$/`.
	* `href` - (String) The URL for this trust store certificate.
	* `id` - (String) Id of the certificate.
	  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-fA-F]*$/`.
	* `issued` - (String) The Date the certificate was issued.
	* `issuer_cn` - (String) Issuer's Common Name.
	* `issuer_dn` - (String) Issuer's Distinguished Name.
	* `label` - (String) Certificate label in queue manager store.
	  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_.]*$/`.
	* `subject_cn` - (String) Subject's Common Name.
	* `subject_dn` - (String) Subject's Distinguished Name.
	* `trusted` - (Boolean) Indicates whether a certificate is trusted.

