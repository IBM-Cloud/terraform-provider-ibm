---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_mtlss"
description: |-
  Get information on an IBM Cloud Internet Services mTLS.
---

# ibm_cis_mtlss

Retrieve information about an IBM Cloud Internet Services mTLS data sources. For more information, see [IBM Cloud Internet Services](https://cloud.ibm.com/docs/cis?topic=cis-about-ibm-cloud-internet-services-cis).

## Example usage

```terraform
data "ibm_cis_mtlss" "tests" {
  cis_id    = ibm_cis.instance.id
  domain_id = ibm_cis_domain.example.id
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `cis_id` - (Required, String) The ID of the CIS service instance.
- `cis_domain` - (Required, String) The Domain of the CIS service instance.


## Attributes reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `cis_id` - (String) The ID of the CIS service instance.
- `cis_domain` - (String) The Domain of the CIS service instance.
- `mtls_certificates` - (List)
   - `id` - (String) The Certificate ID.
   - `name` - (String) The Certificate Name.
   - `fingerprint` - (String) The Certificate Fingerprint.
   - `associated_hostnames` - (String) The Certificate Associated Hostnames.
   - `created_at` - (String) The Certificate Created At.
   - `updated_at` - (String) The Certificate Updated At.
   - `expires_on` - (String) The Certificate Expires On.

