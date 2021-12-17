---

subcategory: "Power Systems"
layout: "ibm"
page_title: "IBM: pi_sap_profiles"
description: |-
  Manages SAP profiles in the Power Virtual Server cloud.
---

# ibm_pi_sap_profiles
Retrieve information about all SAP profiles. For more information, see [getting started with IBM Power Systems Virtual Servers](https://cloud.ibm.com/docs/power-iaas?topic=power-iaas-getting-started).

## Example usage

```terraform
data "ibm_pi_sap_profiles" "example" {
  pi_cloud_instance_id = "<value of the cloud_instance_id>"
}
```

**Notes**

* Please find [supported Regions](https://cloud.ibm.com/apidocs/power-cloud#endpoint) for endpoints.
* If a Power cloud instance is provisioned at `lon04`, The provider level attributes should be as follows:
  * `region` - `lon`
  * `zone` - `lon04`

Example usage:

  ```terraform
    provider "ibm" {
      region    =   "lon"
      zone      =   "lon04"
    }
  ```
  
## Argument reference
Review the argument references that you can specify for your data source.

- `pi_cloud_instance_id` - (Required, String) The GUID of the service instance associated with an account.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `profiles` - (List) List of all the SAP Profiles.

  Nested scheme for `profiles`:
  - `certified` - (Boolean) Has certification been performed on profile.
  - `cores` - (Integer) Amount of cores.
  - `memory` - (Integer) Amount of memory (in GB).
  - `profile_id` - (String) SAP Profile ID.
  - `type` - (String) Type of profile.
