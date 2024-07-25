---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_origin_certificates"
description: |-
  Get information on IBM Cloud Internet Services Certificates.
---

# ibm_cis_origin_certificates

Retrieve the information of an existing IBM Cloud Internet Services certificates resource. For more information about CIS origin certificates, refer to [managing origin certificates](https://cloud.ibm.com/docs/cis?topic=cis-cis-origin-certificates).

## Example usage

```terraform

data ibm_cis_origin_certificates "test" {
  cis_id    = ibm_cis.instance.id
  domain_id = ibm_cis_domain.example.id
  certificate_id = "25392180178235735583993116186144990011711092749"
}
```

## Argument reference

Review the argument references that you can specify for your data source.

- `cis_id` - (Required, String) The ID of the CIS instance.
- `domain_id` - (Required, String) The ID of the domain.
- `certificate_id` - (Optional, String) The ID of the certificate. If the ID is not provided, you will get the list of certificates. If the ID is provided, then you will get the information of that certificate.

## Attribute reference

In addition to the argument reference list, you can access the following attribute references after your data source is created.

- `origin_certificate_list` - (String) The collection of the certificates.

  origin_certificate_list
  - `certificate_id` - (String) The certificate ID.
  - `certificate` (String) Certificate associated with the origin certificate.
  - `expires_on` - (String) Expiration date of the certificate.
  - `hostnames` - (List[String]) The hosts associated with the certificate.
  - `request_type` - (String) The type of the certificate.
