---
subcategory: "Internet services"
layout: "ibm"
certificate_title: "IBM: ibm_cis_custom_certificates"
description: |-
  Get information on an IBM Cloud Internet Services Custom Certificates resource.
---

# ibm_cis_custom_certificates

Imports a read only copy of an existing Internet Services custom certificates resource.

## Example Usage

```hcl
# Get custom certificates of the domain

data "ibm_cis_custom_certificates" "custom_certificates" {
    cis_id    = data.ibm_cis.cis.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
}
```

## Argument Reference

The following arguments are supported:

- `cis_id` - The ID of the CIS service instance.
- `domain_id` - The ID of the domain to change Custom Certificate.
- `custom_certificates` - The collection of custom certificates.
  - `id` - The custom certificate ID. It is a combination of <`custom_cert_id`>,<`domain_id`>,<`cis_id`> attributes concatenated with ":".
  - `custom_cert_id` - The custom certificate id.
  - `bundle_method` - The custom certificate bundle method.
  - `hosts` - The list of hosts which certificated uploaded.
  - `priority` - The custom certificate priority.
  - `status` - The custom certificate status.
  - `issuer` - The custom certificate issuer.
  - `signature` - The custom certificate signature.
  - `expires_on` - The custom certificate exipres date and time.
  - `uploaded_on` - The custom certificate uploaded date and time.
  - `modified_on` - The custom certificate modified date and time.
