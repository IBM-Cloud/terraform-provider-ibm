---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_origin_auth"
description: |-
  Provides an IBM CIS origin auth resource.
---

# ibm_cis_origin_auth
 Provides origin auth settings resource. The resource allows to create, update, or delete cache settings of a domain of an IBM Cloud Internet Services CIS instance. For more information about CIS origin auth, see [CIS ORIGIN AUTH](https://cloud.ibm.com/docs/cis?topic=cis-cli-plugin-cis-cli#authenticated-origin-pull).

## Example usage

```terraform
# Change origin auth setting of CIS instance

resource "ibm_cis_origin_auth" "orig_auth_settings" {
  cis_id                          = data.ibm_cis.cis.id
  domain_id                       = data.ibm_cis_domain.cis_domain.domain_id
  certificate                     = EOT<<
                                  "-----BEGIN CERTIFICATE------ 
                                   ------END CERTIFICATE-------"
                                  EOT
  private_key                     = <<EOT # pragma: whitelist secret
                                  "-----BEGIN------ 
                                  -------END--------"
                                  EOT
  hostname                        = "abc.abc.abc.com"
  level                           = "zone"
}
```

## Argument reference

Review the argument references that you can specify for your resource. 

- `cis_id`                  - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `domain_id`               - (Required, String) The ID of the domain to change cache settings.
- `certificate`             - (Required, String) Content of certificate.
- `private_key`             - (Required, String) Content of private key. # pragma: whitelist secret.
- `level  `                 - (Required, String) Origin Auth setting level  zone or hostname.
- `hostname`                - (optional, String) Valid host names for host level origin auth processing.
- `enabled`                 - (optional, Bool)   Default is true, it enables/disables the host and zone level origin auth setting.



## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id`           - (Computed, String) The record ID. It is a combination of `<auth_id>,<level>,<domain_id>,<cis_id>` attributes concatenated with `:`.
- `updated_at`   - (Computed, String) Time stamp string when Certificate is modified'.
- `expires_on`   - (Computed, String) Time stamp string when Cerftificate expires on'.
- `cert_id`      - (Computed, String) Uploaded certificate ID.
- `status`       - (Computed, String) Origin auth status enbled or not.


## Import
The `ibm_cis_origin_auth` resource can be imported using the ID. The ID is formed from auth ID, level, the domain ID of the domain and the CRN concatenated  using a `:` character.

The domain ID and CRN will be located on the overview page of the IBM Cloud Internet Services instance of the console domain heading, or by using the `ibmcloud cis` command line commands.

- **Auth ID** is a string of the form: `fa633cc7-4afc-4875-8814-b321153fee13`

- **Level** is a string of the form: `zone`

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`


**Syntax**

```
$ terraform import ibm_cis_origin_auth.origin_auth_settings <auth_id>:<level>:<domain-id>:<crn>
```

**Example**

```
$ terraform import ibm_cis_origin_auth.origin_auth_settings fa633cc7-4afc-4875-8814-b321153fee13:zone:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```

