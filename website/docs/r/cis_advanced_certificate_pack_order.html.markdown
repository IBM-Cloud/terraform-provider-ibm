---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_advanced_certificate_pack_order"
description: |-
  Provides an IBM CIS certificate order resource.
---

# ibm_cis_advanced_certificate_pack_order

 Provides an IBM Cloud Internet Services advanced certificate order resource. This resource is associated with an IBM Cloud Internet Services instance and a CIS domain resource. It allows you to order and delete dedicated advanced certificates of a domain of a CIS instance. For more information about CIS certificate ordering, see [managing edge certificates](https://cloud.ibm.com/docs/cis?topic=cis-managing-edge-certs).

## Example usage

```terraform
resource "ibm_cis_advanced_certificate_pack_order" "test" {
    cis_id    = data.ibm_cis.cis.id
    domain_id = data.ibm_cis_domain.cis_domain.domain_id
    hosts     = ["example.com"]
    certificate_authority = "lets_encrypt"
    cloudflare_branding = false
    validation_method = "txt"
    validity = 90
}
```

## Argument reference

Review the argument references that you can specify for your resource.

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `domain_id` - (Required, String) The ID of the domain.
- `hosts` - (Required, String) The hosts for the certificates to be ordered.
- `certificate_authority` - (Required, String) The certificate authority selected for the order. Allowed values are `google` and `lets_encrypt`
- `cloudflare_branding` - (Optional, Boolean) Whether to add Cloudflare branding for the order.
- `validation_method` - (Required, String) Validation methond selected for the order. Allowed values are `txt`, `http`, and `email`.
- `validity`- (Required, Int) Validty days for the order. Allowed values are `14`, `30`, `90`, `365`.

## Attribute reference

In addition to the argument reference list, you can access the following attribute reference after your resource is created.

- `certificate_id`- (String) The certificate ID.
- `id` - (String) The record ID, which is a combination of `<certificate_id>,<domain_id>,<cis_id>` attributes concatenated with `:`.
- `status`- (String) The certificate status.
