---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_certificates"
description: |-
  Get information on IBM Cloud Internet Services Certificates.
---

# ibm_cis_certificates
Retrieve an information of an existing IBM Cloud Internet Services certificates resource. For more information about CIS certificate order, refer to [managing origin certificates](https://cloud.ibm.com/docs/cis?topic=cis-cis-origin-certificates).

## Example usage

```terraform
data "ibm_cis_certificates" "test" {
  cis_id    = ibm_cis.instance.id
  domain_id = ibm_cis_domain.example.id
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `cis_id` - (Required, String) The ID of the CIS instance.
- `domain_id` - (Required, String) The ID of the domain.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `certificates` - (String) The collection of the certificates.

   Nested scheme for `certificates`:

   - `certificate_id` - (String) The certificate ID.
   - `cprimary_certificate` - (String) The primary certificate ID.
   - `certificates` (List) The list of certificates associated with the ordered certificate.

       Nested scheme for `certificates`:
	   - `id` - (String) The certificate ID.
	   - `hosts` - (String) The hosts of the associated with the certificates.
	   - `status` - (String) The certificate status.
   - `hosts` - (String) The hosts of the ordered certificates.
   - `id` - (String) It is a combination of `<certificate_id>:<domain_id>:<cis_id>`.
   - `status` - (String) The certificate status.
   - `type` - (String) The certificate type.

