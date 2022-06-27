---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_mtls_apps"
description: |-
  Get information on an IBM Cloud Internet Services mTLS Applications and Policies.
---

# ibm_cis_mtlss

Retrieve information about an IBM Cloud Internet Services mTLS Applications data sources and fetch Policies data source, with respect to Application ID. For more information, see [IBM Cloud Internet Services](https://cloud.ibm.com/docs/cis?topic=cis-about-ibm-cloud-internet-services-cis).

## Example usage

```terraform
data "ibm_cis_mtls_apps" "tests" {
   cis_id    = ibm_cis.instance.id
   domain_id = ibm_cis_domain.example.id
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `cis_id` - (Required, String) The ID of the CIS service instance.
- `cis_domain` - (Required, String) The Domain of the CIS service instance.


## Attributes reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `cis_id` - (String) The ID of the CIS service instance.
- `cis_domain` - (String) The Domain of the CIS service instance.
- `mtls_access_apps` - (List)
   - `app_id` - (String) The Application ID.
   - `app_name` - (String) The Application Name.
   - `app_domain` - (String) The Application Domain.
   - `app_aud` - (String) The Application Aud.
   - `allowed_idps` - (List) The List of allowed idps.
   - `auto_redirect_to_identity` - (Bool) Auto Redirect to Identity.
   - `session_duration` - (String) The Session Duration.
   - `app_type` - (String) The Session Type.
   - `app_uid` - (String) The Application Uid.
   - `app_created_at` - (String) The Application Created At.
   - `app_updated_at` - (String) The Application Updated At.
- `mtls_access_app_policies` - (List)
   - `policy_id` - (String) The Policy ID.
   - `policy_name` - (String) The Policy Name.
   - `policy_decision` - (String) The Policy Decision.
   - `policy_precedence` - (Int) The Policy Precedence.
   - `policy_uid` - (String) The Policy Uid.
   - `policy_created_at` - (String) The Policy Created At.
   - `policy_updated_at` - (String) The Policy Updated At.

