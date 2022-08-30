---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_origin_auths"
description: |-
  Get information on an IBM Cloud Internet Services Authentication Origin APIs.
---

# ibm_cis_origin_auths

Retrieve information about an IBM Cloud Internet Services authentication Origin data sources for both Zone level and per hostname. For more information, see [IBM Cloud Internet Services](https://cloud.ibm.com/docs/cis?topic=cis-about-ibm-cloud-internet-services-cis).

## Example usage

```terraform
data "ibm_cis_origin_auths" "tests" {
	cis_id    = ibm_cis.instance.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id

}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `cis_id` - (Required, String) The ID of the CIS service instance.
- `domain_id` - (Required, String) The Domain ID of the CIS service instance.
- `request_type` - (Optional, String) The type of request made. Can be `zone_level` or `per_hostname`. Default value :`zone_level`.
- `hostname` - (Optional, String) The Hostname of the CIS service instance. Only needed when `request_type = per_hostname`.


## Attributes reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `id` - (String) The ID of the CIS service instance.
- `domain_id` - (String) The Domain ID of the CIS service instance.
- `request_type` - (String) The type of request made. Can be `zone_level` or `per_hostname`. Default value :`zone_level`.
- `hostname` - (String) The Hostname of the CIS service instance. Only needed when `request_type = per_hostname`.
- `origin_pull_settings_enabled` - (Bool) True if the Authentication Origin Settings are enabled.
- `origin_pull_certs_list` - (List)
   - `cert_id` - (String) The Auth Origin Certificate ID.
   - `certificate` - (String) The Auth Origin Certificate Detail.
   - `cert_issuer` - (String)  The Auth Origin Certificate Issuer.
   - `cert_signature` - (Boolean)  The Auth Origin Certificate Signature.
   - `cert_status` - (String)  The Auth Origin Certificate Status
   - `cert_expires_on` - (String)  The Auth Origin Certificate Expires On.
   - `cert_uploaded_on` - (String)  The Auth Origin Certificate Uploaded On.
   - `cert_serial_number` - (String)  The Auth Origin Certificate Serial Number.

