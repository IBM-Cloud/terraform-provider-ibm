---
layout: "ibm"
page_title: "IBM : ibm_mqcloud_keystore_certificate"
description: |-
  Manages mqcloud_keystore_certificate.
subcategory: "MQaaS"
---

# ibm_mqcloud_keystore_certificate

Create, update, and delete mqcloud_keystore_certificates with this resource.

> **Note:** The MQaaS Terraform provider access is restricted to users of the reserved deployment, reserved capacity, and reserved capacity subscription plans.

## Example Usage

```hcl
resource "ibm_mqcloud_keystore_certificate" "mqcloud_keystore_certificate_instance" {
  certificate_file = filebase64("certificate_file.data")
  label = "certlabel"
  queue_manager_id = var.queue_manager_id
  service_instance_guid = var.service_instance_guid
  config {
    ams {
      channels {
        name = "CLOUD.APP.SVRCONN"
      }
    }
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `certificate_file` - (Required, Forces new resource, String) The filename and path of the certificate to be uploaded.
  * Constraints: The maximum length is `65537` characters. The minimum length is `1500` characters.
* `label` - (Required, Forces new resource, String) The label to use for the certificate to be uploaded.
  * Constraints: The maximum length is `64` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_.]*$/`.
* `queue_manager_id` - (Required, Forces new resource, String) The id of the queue manager to retrieve its full details.
  * Constraints: The maximum length is `32` characters. The minimum length is `32` characters. The value must match regular expression `/^[0-9a-fA-F]{32}$/`.
* `service_instance_guid` - (Required, Forces new resource, String) The GUID that uniquely identifies the MQaaS service instance.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the mqcloud_keystore_certificate.
* `certificate_id` - (String) ID of the certificate.
  * Constraints: The maximum length is `16` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-fA-F]*$/`.
* `certificate_type` - (String) The type of certificate.
  * Constraints: Allowable values are: `key_store`.
* `config` - (List) The configuration details for this certificate.
Nested schema for **config**:
	* `ams` - (List) A list of channels that are configured with this certificate.
	Nested schema for **ams**:
		* `channels` - (List) A list of channels that are configured with this certificate.
		  * Constraints: The maximum length is `999999999` items. The minimum length is `0` items.
		Nested schema for **channels**:
			* `name` - (String) The name of the channel.
			  * Constraints: The maximum length is `20` characters. The minimum length is `1` character. The value must match regular expression `/^[a-zA-Z0-9_.\/%]*$/`.
* `dns_names` - (List) The list of DNS names.
  * Constraints: The maximum length is `200` items. The minimum length is `0` items.
* `dns_names_total_count` - (Integer) The total count of dns names.
* `expiry` - (String) Expiry date for the certificate.
* `fingerprint_sha256` - (String) Fingerprint SHA256.
  * Constraints: The value must match regular expression `/^[A-F0-9]{2}(:[A-F0-9]{2}){31}$/`.
* `href` - (String) The URL for this key store certificate.
* `is_default` - (Boolean) Indicates whether it is the queue manager's default certificate.
* `issued` - (String) Date certificate was issued.
* `issuer_cn` - (String) Issuer's Common Name.
* `issuer_dn` - (String) Issuer's Distinguished Name.
* `subject_cn` - (String) Subject's Common Name.
* `subject_dn` - (String) Subject's Distinguished Name.


## Import

You can import the `ibm_mqcloud_keystore_certificate` resource by using `id`.
The `id` property can be formed from `service_instance_guid`, `queue_manager_id`, and `certificate_id` in the following format:

<pre>
&lt;service_instance_guid&gt;/&lt;queue_manager_id&gt;/&lt;certificate_id&gt;
</pre>
* `service_instance_guid`: A string in the format `a2b4d4bc-dadb-4637-bcec-9b7d1e723af8`. The GUID that uniquely identifies the MQaaS service instance.
* `queue_manager_id`: A string in the format `b8e1aeda078009cf3db74e90d5d42328`. The id of the queue manager to retrieve its full details.
* `certificate_id`: A string. ID of the certificate.

> ### Important Note
> When configuring the `ibm_mqcloud_keystore_certificate` resource in the root module:
> Ensure to set the `certificate_file` value to an empty string (`certificate_file=""`). This step is crucial as we are not downloading the certificate to the local system.

# Syntax
<pre>
$ terraform import ibm_mqcloud_keystore_certificate.mqcloud_keystore_certificate &lt;service_instance_guid&gt;/&lt;queue_manager_id&gt;/&lt;certificate_id&gt;
</pre>
