---
layout: "ibm"
page_title: "IBM: ibm_cis_tls_settings"
sidebar_current: "docs-ibm-resource-cis-tls-settings"
description: |-
  Provides a IBM CIS TLS Settings resource.
---

# ibm_cis_tls_settings

Provides a IBM CIS TLS Settings resource. This resource is associated with an IBM Cloud Internet Services instance and a CIS Domain resource. It allows to change TLS settings of a domain of a CIS instance

## Example Usage

```hcl
# Change TLS Settings of the domain

resource "ibm_cis_tls_settings" "%[1]s" {
	cis_id          = data.ibm_cis.cis.id
	domain_id       = data.ibm_cis_domain.cis_domain.domain_id
	tls_1_3         = "off"
	tls_1_2_only    = "on"
	min_tls_version = "1.2"
	universal_ssl   = true
}
```

## Argument Reference

The following arguments are supported:

- `cis_id` - (Required,string) The ID of the CIS service instance.
- `domain_id` - (Required,string) The ID of the domain to change TLS settings.
- `tls_1_3` - (Optional, string) The TLS 1.3 version setting. Valid values are `on`, `off`, `zrt`. `zrt` will enable `tls 1.3` and the `Zero RTT` feature.  If `on` is set, then `zrt` will be enabled by default.
- `min_tls_version` - (Optional, string) The Minimum TLS version setting. Valid values are `1.1`,`1.2`,`1.3`,`1.4`
- `universal_ssl` - (Optional, boolean) The Universal SSL enable/disable setting.
- `ssl_mode` - (Optional, string) The SSL mode settings.  `This is not supported yet`.

## Attributes Reference

The following attributes are exported:

- `id` - The record ID. It is a combination of <`domain_id`>,<`cis_id`> attributes concatenated with ":".

## Import

The `ibm_cis_tls_settings` resource can be imported using the `id`. The ID is formed from the `Domain ID` of the domain and the `CRN` (Cloud Resource Name) concatentated using a `:` character.

The Domain ID and CRN will be located on the **Overview** page of the Internet Services instance under the **Domain** heading of the UI, or via using the `ibmcloud cis` CLI commands.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

```
$ terraform import ibm_cis_tls_settings.tls_settings <domain-id>:<crn>

$ terraform import ibm_cis_tls_settings.tls_settings 9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```
