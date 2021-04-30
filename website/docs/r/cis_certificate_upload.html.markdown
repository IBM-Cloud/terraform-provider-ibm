---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_certificate_upload"
description: |-
  Provides a IBM CIS Certificate Upload resource.
---

# ibm_cis_certificate_upload

Provides a IBM CIS Certificate upload resource. This resource is associated with an IBM Cloud Internet Services instance and a CIS Domain resource. It allows to upload, update, delete certificates of a domain of a CIS instance

## Example Usage

```hcl
# Upload a certificate for a domain

resource "ibm_cis_certificate_upload" "cert" {
    cis_id        = data.ibm_cis.cis.id
    domain_id     = data.ibm_cis_domain.cis_domain.domain_id
    certificate   = "xxxxx"
    private_key   = "xxxxx
    bundle_method = "ubiquitous"
    priority      = 20
}
```

## Argument Reference

The following arguments are supported:

- `cis_id` - (Required,string) The ID of the CIS service instance
- `domain_id` - (Required,string) The ID of the domain to add the Certificate Upload rule.
- `certificate` - (Required,string) The intermediate(s) certificate key.
- `private_key` - (Required,string) The certificate private key.
- `bundle_method` - (Optional,string) The certificate bundle method. The valid values are: `ubiquitous`, `optimal`, `force`.
- `priority` - (Optional,integer) The order/priority in which the certificate is used in a request.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- `id` - The record ID. It is a combination of <`custom_cert_id`>,<`domain_id`>,<`cis_id`> attributes concatenated with ":".
- `custom_cert_id` - The Certificate Upload Rule ID.
- `status` - The ceritificate status.
- `issuer` - The certificate issuer.
- `signature` - The certificate signature.
- `expires_on` - The certificate exipres date and time.
- `uploaded_on` - The certificate uploaded date and time.
- `modified_on` - The certificate modified date and time.

## Import

The `ibm_cis_certificate_upload` resource can be imported using the `id`. The ID is formed from the `Certificate Upload ID`, the `Domain ID` of the domain and the `CRN` (Cloud Resource Name) concatentated using a `:` character.

The Domain ID and CRN will be located on the **Overview** page of the Internet Services instance under the **Domain** heading of the UI, or via using the `ibmcloud cis` CLI commands.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

- **Certificate Upload ID** is a 32 digit character string of the form: `489d96f0da6ed76251b475971b097205c`.

```
$ terraform import ibm_cis_certificate_upload.ratelimit <custm_cert_id>:<domain-id>:<crn>

$ terraform import ibm_cis_certificate_upload.certificate 48996f0da6ed76251b475971b097205c:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
