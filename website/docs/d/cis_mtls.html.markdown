---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_mtlss"
description: |-
  Get information on an IBM Cloud Internet Services mTLS.
---

# ibm_cis_alerts

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
   - `cert_id` - (String) The Certificate ID.
   - `cert_name` - (String) The Certificate Name.
   - `cert_fingerprint` - (String) The Certificate Fingerprint.
   - `cert_associated_hostnames` - (String) The Certificate Associated Hostnames.
   - `cert_created_at` - (String) The Certificate Created At.
   - `cert_updated_at` - (String) The Certificate Updated At.
   - `cert_expires_on` - (String) The Certificate Expires On.

