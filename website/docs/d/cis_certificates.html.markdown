---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_certificates"
description: |-
  Get information on IBM Cloud Internet Services Certificates.
---

# ibm_cis_certificates

Imports a read only copy of an existing Internet Services Certificates resource.

## Example Usage

```hcl
data "ibm_cis_certificates" "test" {
  cis_id    = ibm_cis.instance.id
  domain_id = ibm_cis_domain.example.id
}
```

## Argument Reference

The following arguments are supported:

- `cis_id` - The ID of the CIS service instance.
- `domain_id` - The ID of the domain.

## Attributes Reference

The following attributes are exported:

- `certificates` - The Collection of certificates
  - `id` - It is a combination of <`certificate_id`>,<`domain_id`>,<`cis_id`> attributes concatenated with ":".
  - `certificate_id` - The certificate id.
  - `type` - The certificate type.
  - `hosts` - The hosts for which certificates ordered.
  - `status` - The certificate status.
  - `primary_certificate` - The primary certificate id.
  - `certificates` - The list of certificates associated with the ordered certificate.
    - `id` - The certificate id.
    - `hosts` - The hosts which the certificates associated with.
    - `status` - The certificate status.
