---
layout: "ibm"
page_title: "IBM : ibm_scc_control_libraries"
description: |-
  Get information about scc_control_libraries
subcategory: "Security and Compliance Center"
---

# ibm_scc_control_libraries

Retrieve information about a list of scc_control_libraries from a read-only data source. Then, you can reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

~> NOTE: Security Compliance Center is a regional service. Please specify the IBM Cloud Provider attribute `region` to target another region. Else, exporting the environmental variable IBMCLOUD_SCC_API_ENDPOINT will also override which region is being targeted for all ibm providers(ex. `export IBMCLOUD_SCC_API_ENDPOINT=https://eu-es.compliance.cloud.ibm.com`).

## Example Usage

```hcl
data "ibm_scc_control_libraries" "scc_control_libraries" {
    instance_id = "00000000-1111-2222-3333-444444444444"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `control_library_type` - (Optional, Forces new resource, String) The type of control library to query.
  * Constraints: Allowable values are: `predefined`, `custom`.
* `instance_id` - (Required, Forces new resource, String) The ID of the SCC instance in a particular region.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `control_libraries` - (List) The list of control libraries.

    Nested schema for **control_libraries**:
    * `id` - (String) The unique identifier of the scc_control_library.

    * `account_id` - (String) The account ID.

    * `control_library_description` - (String) The control library description.

    * `control_library_name` - (String) The control library name.

    * `control_library_type` - (String) The control library type.

    * `control_library_version` - (String) The control library version.

    * `control_count` - (Integer) The number of controls in the control library.

    * `created_by` - (String) The user who created the control library.

    * `created_on` - (String) The date when the control library was created.

    * `updated_by` - (String) The user who updated the control library.

    * `updated_on` - (String) The date when the control library was updated.
    
    * `version_group_label` - (String) The version group label.
