---
layout: "ibm"
page_title: "IBM : ibm_scc_profiles"
description: |-
  Get information about scc_profiles
subcategory: "Security and Compliance Center"
---

# ibm_scc_profiles

Retrieve information about a list of profiles from a read-only data source. Then, you can reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

~> NOTE: Security Compliance Center is a regional service. Please specify the IBM Cloud Provider attribute `region` to target another region. Else, exporting the environmental variable IBMCLOUD_SCC_API_ENDPOINT will also override which region is being targeted for all ibm providers(ex. `export IBMCLOUD_SCC_API_ENDPOINT=https://eu-es.compliance.cloud.ibm.com`).

## Example Usage

```hcl
data "ibm_scc_profiles" "scc_profiles_instace" {
    instance_id = "00000000-1111-2222-3333-444444444444"
    profile_type = ibm_scc_profile.scc_profile_instance.profile_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `profile_type` - (Optional, Forces new resource, String) The type of profiles to query.
  * Constraints: Allowable values are: `predefined`, `custom`.
* `instance_id` - (Required, Forces new resource, String) The ID of the SCC instance in a particular region.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `profiles` - (List) The list of profiles.

    Nested schema for **profiles**:
    * `id` - The unique identifier of the scc_profile.

    * `attachments_count` - (Integer) The number of attachments related to this profile.

    * `control_parents_count` - (Integer) The number of parent controls for the profile.    

    * `instance_id` - (String) The instance ID.

    * `latest` - (Boolean) The latest version of the profile.

    * `profile_description` - (String) The profile description.
      * Constraints: The maximum length is `256` characters. The minimum length is `2` characters. The value must match regular expression `/[A-Za-z0-9]+/`.

    * `profile_name` - (String) The profile name.


    * `profile_type` - (String) The profile type, such as custom or predefined.

    * `profile_version` - (String) The version status of the profile.

    * `version_group_label` - (String) The version group label of the profile.

    * `latest` - (Boolean) The latest version of the profile.

    * `hierarchy_enabled` - (Boolean) The indication of whether hierarchy is enabled for the profile.

    * `created_by` - (String) The user who created the profile.

    * `created_on` - (String) The date when the profile was created.

    * `controls_count` - (Integer) The number of controls for the profile.

    * `control_parents_count` - (Integer) The number of parent controls for the profile.

    * `attachments_count` - (Integer) The number of attachments related to this profile.