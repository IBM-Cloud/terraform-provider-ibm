---
layout: "ibm"
page_title: "IBM : ibm_scc_profiles"
description: |-
  Get information about scc_profiles
subcategory: "Security and Compliance Center"
---

# ibm_scc_profiles

Retrieve information about a list of profiles from a read-only data source. Then, you can reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

~> NOTE: if you specify the `region` in the provider, that region will become the default URL. Else, exporting the environmental variable IBMCLOUD_SCC_API_ENDPOINT will override any URL(ex. `export IBMCLOUD_SCC_API_ENDPOINT=https://us-south.compliance.cloud.ibm.com`).

## Example Usage

```hcl
data "ibm_scc_profiles" "scc_profiles" {
    instance_id = "00000000-1111-2222-3333-444444444444"
    profile_id = ibm_scc_profile.scc_profile_instance.profile_id
}
```