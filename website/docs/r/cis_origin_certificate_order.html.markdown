---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_origin_certificate_order"
description: |-
  Provides an IBM CIS origin certificate order resource.
---

# ibm_cis_origin_certificate_order

Provides an IBM Cloud Internet Services origin certificate order resource. This resource is associated with an IBM Cloud Internet Services instance and a CIS domain resource. It allows you to order and delete dedicated advanced certificates of a domain of a CIS instance. For more information about CIS certificate orderering, see [managing origin certificates](https://cloud.ibm.com/docs/cis?topic=cis-cis-origin-certificates).

## Example usage

```terraform

resource "ibm_cis_origin_certificate_order" "test" {
    cis_id    = data.ibm_cis.cis.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
    hostnames     = ["example.com"]
    request_type = "origin-rsa"
    requested_validity = 5475
    csr = "-----BEGIN CERTIFICATE REQUEST-----\nMIICxzCC***TA67sdbcQ==\n-----END CERTIFICATE REQUEST-----"
}

```

## Argument reference

Review the argument references that you can specify for your resource.

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `domain_id` - (Required, String) The ID of the domain.
- `hosts` - (Required, String) The hosts for the certificates to be ordered.
- `request_type` - (Required, String) The type of the certificate. Allowed values are `origin-rsa`, `origin-ecc` and `keyless-certificate`.
- `requested_validity`- (Required, Int) Validty days for the order. Allowed values are `7`, `30`, `90`, `365`, `730`, `1095`, `5475`.
- `csr` - (Required, String) The Certificate Signing Request.

## Attribute reference

In addition to the argument reference list, you can access the following attribute reference after your resource is created.

- `certificate_id`- (String) The certificate ID.
- `id` - (String) The record ID, which is a combination of `<certificate_id>,<domain_id>,<cis_id>` attributes concatenated with `:`.

## Import

The `ibm_cis_origin_certificate_order` resource can be imported using the ID. The ID is formed from the certificate ID, the domain ID of the domain and the CRN  concatenated  by using a `:` character.

The domain ID and CRN is located on the **Overview** page of the IBM Cloud Internet Services instance of the console domain heading, or by using the `ibmcloud cis` command line commands.

- **Domain ID** is a 32-digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120-digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

- **Certificate ID** is a 48-digit character string of the form: `484582976896327736468082847548136290560450732393`.

### Syntax

```terraform
terraform import ibm_cis_origin_certificate_order.test <certificate_id>:<domain-id>:<crn>
```

### Example

```terraform
terraform import ibm_cis_origin_certificate_order.test certificate_order 484582976896327736468082847548136290560450732393:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
