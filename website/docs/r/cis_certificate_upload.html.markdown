---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_certificate_upload"
description: |-
  Provides a IBM CIS certificate upload resource.
---

# ibm_cis_certificate_upload

Provides an IBM Cloud Internet Services certificate upload resource. This resource is associated with an IBM Cloud Internet Services instance and a CIS domain resource. It allows to upload, update, and delete certificates of a domain of a CIS instance. For more information about CIS certificate upload, see [Installing an origin certificate on your server](https://cloud.ibm.com/docs/cis?topic=cis-cis-origin-certificates#cis-origin-certificates-installing).

## Example usage

```terraform
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

## Argument reference
Review the argument references that you can specify for your resource. 

- `bundle_method` - (Optional, String) The certificate bundle method. The valid values are `ubiquitous`, `optimal`, `force`.
- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `certificate` - (Required, String) The intermediate(s) certificate key.
- `domain_id` - (Required, String) The ID of the domain to add the rules certificate upload.
- `private_key` - (Required, String) The certificate private key.
- `priority` - (Optional, Integer) The order or priority in which the certificate is used in a request.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `certificate_id`-String -The certificate ID.
- `status`-String -The certificate status.
- `custom_cert_id` - (String) The certificate upload rule ID.
- `expires_on` - (Timestamp) The expiry date and time of the certificate.
- `id` - (String) The record ID. It is a combination of `<custom_cert_id>:<domain_id>:<cis_id>` attributes concatenated with `:`.
- `issuer` - (String) The certificate issuer.
- `modified_on` - (Timestamp) The modified date and time of the certificate.
- `signature` - (String) The certificate signature.
- `status` - (String) The certificate status.
- `uploaded_on` - (Timestamp) The uploaded date and time of the certificate.

## Import
The `ibm_cis_certificate_upload` resource can be imported by using the ID. The ID is formed from the certificate upload ID, the domain ID of the domain and the CRN  Concatenated  by using a `:` character.

The domain ID and CRN is located on the **Overview** page of the IBM Cloud Internet Services instance of the console domain heading, or by using the `ibmcloud cis` command line commands.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

- **Certificate upload ID** is a 32 digit character string of the form: `489d96f0da6ed76251b475971b097205c`.


**Syntax**

```
$ terraform import ibm_cis_certificate_upload.ratelimit <custm_cert_id>:<domain-id>:<crn>
```

**Example**

```
$ terraform import ibm_cis_certificate_upload.certificate 48996f0da6ed76251b475971b097205c:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
