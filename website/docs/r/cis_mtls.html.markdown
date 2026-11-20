---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_mtls"
description: |-
  Provides an IBM Mutual TLS(mTLS) resource.
---

# ibm_cis_mtls
 Provides mutual TLS(mTLS) certificate settings resource. The resource allows to create(upload), update, or delete mTLS certificate settings of a domain of an IBM Cloud Internet Services CIS instance. For more information about mtls, see [CIS MTLS](https://cloud.ibm.com/docs/cis?topic=cis-mtls-features).

## Example usage
```terraform
# Change mTLS certificate setting of CIS instance

resource "ibm_cis_mtls" "mtls_settings" {
  cis_id                          = data.ibm_cis.cis.id
  domain_id                       = data.ibm_cis_domain.cis_domain.domain_id
  certificate                     = EOT<<
                                  "-----BEGIN CERTIFICATE----- 
                                   ------END CERTIFICATE------"
                                  EOT
  name                            = "MTLS_Cert"
  associated_hostnames            = ["abc.abc.abc.com"]
}

# Bind one client CA to hostnames in several zones of the same CIS instance

resource "ibm_cis_mtls" "client_ca" {
  cis_id      = data.ibm_cis.cis.id
  name        = "client-ca"
  certificate = file("ca.pem")

  zone {
    domain_id = ibm_cis_domain.zone_a.domain_id
    hostnames = ["api.zone-a.example.com"]
  }
  zone {
    domain_id = ibm_cis_domain.zone_b.domain_id
    hostnames = ["api.zone-b.example.com"]
  }
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `cis_id`                  - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `domain_id`               - (Optional, String) The ID of the domain to associate the certificate with. Exactly one of `domain_id` or `zone` must be set.
- `certificate`             - (Required, Forces new resource, String) Content of valid MTLS certificate.
- `name`                    - (Required, String) Valid name for certificate. 
- `associated_hostnames`    - (Optional, []String) Valid host names for which we want to add the certificate. Use with `domain_id`; conflicts with `zone`.
- `zone`                    - (Optional, List) The zones of the CIS instance and the host names to associate in each, for a certificate used across several zones. Conflicts with `associated_hostnames`. A host name can only be added through its own zone, so the certificate is created through the first `zone` and each further zone's host names are added through that zone. Host names that belong to none of the listed zones are left unchanged.

  Nested scheme for `zone`:
  - `domain_id`             - (Required, String) The ID of the domain.
  - `hostnames`             - (Required, Set of String) Host names in this domain to associate with the certificate.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id`                      - (String) The record ID. It is a combination of `<mtls_id>,<domain_id>,<cis_id>` attributes concatenated with `:`.
- `mtls_id`                 - (Computed, String) mTLS ID.
- `created_at`              - (Computed, String) Time stamp string when Certificate is created'.
- `updated_at`              - (Computed, String) Time stamp string when Certificate is modified'.
- `expires_on`              - (Computed, String) Time stamp string when Cerftificate expires on'.
- `cert_id`                 - (Computed, String) Created certificate ID.
- `fingerprint`             - (Computed, String) Fingerprint of the certificate.



## Import
The `ibm_cis_mtls` resource can be imported using the ID. The ID is formed from the mTLS ID, domain ID of the domain and the CRN concatenated  using a `:` character.

The domain ID and CRN will be located on the overview page of the IBM Cloud Internet Services instance of the console domain heading, or by using the `ibmcloud cis` command line commands.

- **MTLS ID** is a string of the form: `fa633cc7-4afc-4875-8814-b321153fee13`

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

The API does not return the certificate contents, so after import the configured `certificate` is compared with the certificate `fingerprint`: the same certificate is adopted without replacing the resource, a different one replaces it. To manage several zones after import, add the `zone` blocks to the configuration; the next apply moves the host names into them.

**Syntax**

```
$ terraform import ibm_cis_mtls.mtls_settings <mtlsid>:<domain-id>:<crn>
```

**Example**

```
$ terraform import ibm_cis_mtls.mtls_settings  fa633cc7-4afc-4875-8814-b321153fee13:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```

