---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_mtls_app"
description: |-
  Provides an IBM Mutual TLS app-policy resource.
---

# ibm_cis_mtls_app
 Provides mutual TLS(mTLS) app-policy settings resource. The resource allows to create, update, or delete mTLS app settings of a domain of an IBM Cloud Internet Services CIS instance. For more information about mtls, see [CIS MTLS](https://cloud.ibm.com/docs/cis?topic=cis-mtls-features).

## Example usage

```terraform
# Change mTLS app-policy setting of CIS instance

resource "ibm_cis_mtls_app" "mtls_app_settings" {
  cis_id             = data.ibm_cis.cis.id
  domain_id          = data.ibm_cis_domain.cis_domain.domain_id
  name               = "MY_APP"
  domain             = "abc.abc.com"
  session_duration   = "24h"
  policy_name        = "MTLS_Policy"
}
```

## Argument reference

Review the argument references that you can specify for your resource. 

- `cis_id`                         - (Required, String) The ID of the IBM Cloud Internet Services instance.
- `domain_id`                      - (Required, String) The ID of the domain to change cache settings.
- `name`                           - (Required, String) Name for the app which you want to create.
- `domain`                         - (Required, String) Host domain for which we want to create app. 
- `policy_name`                    - (Optional, String) Valid name for a policy, default name is 'mtls-policy'.
- `session_duration `              - (Optional, String) Duration string, default is '24h'.
- `cert_rule_val`                  - (Optional, String) Valid value for certificate rule option, default is mTLS certificate name.
- `common_rule_val`                - (Optional, String) Valid value for common rule option.
- `policy_decision`                - (Optional, String) Valid policy action value e.g. 'non_identity'(default), 'allow', 'deny', 'bypass'. 


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id`                             - (String) The record ID. It is a combination of `<app_id>,<policy_id>,<domain_id>,<cis_id>` attributes concatenated with `:`.
- `app_created_at`                 - (Computed, String) Time stamp string when App is created.
- `app_updated_at`                 - (Computed, String) Time stamp string when App is modififed.
- `pol_created_at`                 - (Computed, String) Time stamp string when Policy is created.
- `pol_updated_at`                 - (Computed, String) Time stamp string when Policy is modified.
- `app_id`                         - (Computed, String) ID of created App.
- `policy_id`                      - (Computed, String) ID of created policy.

## Import
The `ibm_cis_mtls_app` resource can be imported using the ID. The ID is formed from the app_id, policy_id, domain ID of the domain and the CRN concatenated  using a `:` character.

The domain ID and CRN will be located on the overview page of the IBM Cloud Internet Services instance of the console domain heading, or by using the `ibmcloud cis` command line commands.

- **APP ID**    is a string of the form: `ac633cc7-2afc-4875-9914-c521153fee15`

- **Policy ID** is a string of the form: `fa633cc7-4afc-4875-8814-b321153fee13`

- **Domain ID** is a 32 digit character string of the form: `9caf68812ae9b3f0377fdf986751a78f`

- **CRN** is a 120 digit character string of the form: `crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::`

**Syntax**

```
$ terraform import ibm_cis_mtls_app.mtls_app_settings <app_id>:<poicy_id><domain-id>:<crn>
```

**Example**

```
$ terraform import ibm_cis_mtls_app.mtls_app_settings ac633cc7-2afc-4875-9914-c521153fee15:fa633cc7-4afc-4875-8814-b321153fee13:9caf68812ae9b3f0377fdf986751a78f:crn:v1:bluemix:public:internet-svcs:global:a/4ea1882a2d3401ed1e459979941966ea:31fa970d-51d0-4b05-893e-251cba75a7b3::
```

