---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_cache_settings"
description: |-
  Provides an IBM Mutual TLS app-policy resource.
---

# ibm_cis_mtls
 Provides mutaul TLS(MTLS) app-policy settings resource. The resource allows to create, update, or delete cache settings of a domain of an IBM Cloud Internet Services CIS instance. For more information about mtls, see [CIS MTLS](https://cloud.ibm.com/docs/cis?topic=cis-mtls-features).

## Example usage

```terraform
# Change MTLS setting of the domain

resource "ibm_cis_mtls_app" "mtls_app_settings" {
  cis_id             = data.ibm_cis.cis.id
  domain_id          = data.ibm_cis_domain.cis_domain.domain_id
  app_name        = "MY_APP"
  url             = "abc.abc.com"
  duration        = "24h"
  policy_name     = "MTLS_Policy"
}
```

## Argument reference

Review the argument references that you can specify for your resource. 

- `cis_id`          - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `domain_id`       - (Required, String) The ID of the domain to change cache settings.
- `app_name`        - (Required, String) Name for the app which you want to create.
- `url`             - (Required, String) Host name for which we want to create app. 
- `policy_name`     - (Option, String) Valid name for a policy.
- `duration`        - (Option, String) Duraing in string, default is '24h'.
- `policy_action`   - (Option, String) Valid policuy action, default is 'non_identity'.

**Note**

Among all the purge actions `purge_all`, `purge_by-urls`, `purge_by_hosts`, and `purge_by_tags`, only one is allowed to give inside a resource.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The record ID. It is a combination of `<domain_id>,<cis_id>` attributes concatenated with `:`.

## Import
The `ibm_cis_mtls` resource can be imported using the ID. The ID is formed from the domain ID of the domain and the CRN concatenated  using a `:` character.

The domain ID and CRN will be located on the overview page of the IBM Cloud Internet Services instance of the console domain heading, or by using the `ibmcloud cis` command line commands.

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

**Syntax**

```
$ terraform import ibm_cis_mtls.mtls_settings <domain-id>:<crn>
```

**Example**

```
$ terraform import ibm_cis_mtls.mtls_settings 9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```

