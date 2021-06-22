---

subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_tls_settings"
description: |-
  Provides a IBM CIS TLS settings resource.
---

# ibm_cis_tls_settings
Create, update, or delete an IBM Cloud Internet Services TLS settings resources. This resource is associated with an IBM Cloud Internet Services instance and an IBM Cloud Internet Services Domain resource. For more information, about CIS TLS settings, see [setting TLS options](https://cloud.ibm.com/docs/cis?topic=cis-cis-tls-options).

## Example usage

```terraform
# Change TLS Settings of the domain

resource "ibm_cis_tls_settings" "tls_settings" {
	cis_id          = data.ibm_cis.cis.id
	domain_id       = data.ibm_cis_domain.cis_domain.domain_id
	tls_1_3         = "off"
	tls_1_2_only    = "on"
	min_tls_version = "1.2"
	universal_ssl   = true
}
```

## Argument reference
Review the argument references that you can specify for your resource.

- `cis_id` - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `domain_id` - (Required, String) The ID of the domain to change TLS settings.
- `min_tls_version` - (Optional, String) The Minimum TLS version setting. Valid values are `1.1`, `1.2`, `1.3`, or `1.4`.
- `ssl_mode` - (Optional, String) The SSL mode settings. This is yet to support.
- `tls_1_3` - (Optional, String) The TLS 1.3 version setting. Valid values are `on`, `off`, `zrt`. `zrt` will enable TLS 1.3 and the Zero RTT feature. If `on` is set, then `zrt` is enabled by default.
- `universal_ssl` - (Optional, Bool) The Universal SSL `enable` or `disable` setting.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The record ID. It is a combination of <domain_id>,<cis_id> attributes concatenated with `:`.

## Import

The `ibm_cis_tls_settings` resource can be imported using the `id`. The ID is formed from the `Domain ID` of the domain and the `CRN` (Cloud Resource Name) concatentated using a `:` character.

The Domain ID and CRN will be located on the **Overview** page of the Internet Services instance under the **Domain** heading of the UI, or via using the `ibmcloud cis` CLI commands.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

**Syntax**

```
$ terraform import ibm_cis_tls_settings.tls_settings <domain-id>:<crn>
```

**Example**

```
$ terraform import ibm_cis_tls_settings.tls_settings 9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
