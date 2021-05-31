---
subcategory: "Internet services"
layout: "ibm"
certificate_title: "IBM: ibm_cis_custom_certificates"
description: |-
  Get information on an IBM Cloud Internet Services custom certificates resource.
---

# ibm_cis_custom_certificates
Retrieve information of an existing IBM Cloud Internet Services custom certificates resource. For more information about CIS certificate order, refer to [upload custom certificates](https://cloud.ibm.com/docs/cis?topic=cis-manage-your-ibm-cis-for-optimal-security#upload-custom-certs).

## Example usage

```terraform
# Get custom certificates of the domain

data "ibm_cis_custom_certificates" "custom_certificates" {
    cis_id    = data.ibm_cis.cis.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `cis_id` - (String) The ID of the IBM Cloud Internet Services instance.
- `custom_certificates` - (String) The collection of the custom certificates.

  Nested scheme for `custom_certificates`:
	- `id` - (String) It is a combination of `<custom_cert_id>:<domain_id>:<cis_id>`.
	- `custom_cert_id` - (String) The custom certificate ID.
	- `bundle_method` - (String) The custom certificate bundle method.
	- `type` - (String) The certificate type.
	- `hosts` - (String) The list of hosts that are uploaded in a certificate.
	- `priority` - (String) The custom certificate priority.
	- `status` - (String) The custom certificate status.
	- `issuer` - (String) The custom certificate issuer.
	- `signature` - (String) The custom certificate signature.
	- `expires_on` - (String) The expiry date and time of the certificate.
	- `uploaded_on` - (String) The uploaded date and time of the certificate.
	- `modified_on` - (String) The modified date and time of the certificate.
- `domain_id` - (String) The ID of the domain to change custom certificate.
