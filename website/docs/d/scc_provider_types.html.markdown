---
layout: "ibm"
page_title: "IBM : ibm_scc_provider_types"
description: |-
  Get information about various scc_provider_types
subcategory: "Security and Compliance Center"
---

# ibm_scc_provider_types

Retrieve information about a provider type from a read-only data source. Then, you can reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

~> NOTE: Security Compliance Center is a regional service. Please specify the IBM Cloud Provider attribute `region` to target another region. Else, exporting the environmental variable IBMCLOUD_SCC_API_ENDPOINT will also override which region is being targeted for all ibm providers(ex. `export IBMCLOUD_SCC_API_ENDPOINT=https://eu-es.compliance.cloud.ibm.com`).

## Example Usage

```hcl
data "ibm_scc_provider_types" "scc_provider_types_instance" {
    instance_id = "00000000-1111-2222-3333-444444444444"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `instance_id` - (Required, Forces new resource, String) The ID of the SCC instance in a particular region.


## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `provider_types` - (List) The list of provider_types.

* `id` - The unique identifier of the scc_provider_type.

* `type` - (String) The type of the provider type.

* `name` - (String) The name of the provider type.

* `description` - (String) The provider type description.

* `s2s_enabled` - (Boolean) A boolean that indicates whether the provider type is s2s-enabled.
  
  **NOTE;** If the provider type is s2s-enabled, which means that if the `s2s_enabled` field is set to `true`, then a CRN field of type text is required in the attributes value object when creating a `ibm_scc_provider_type_instance`

* `attributes` - (Map) The attributes that are required when you're creating an instance of a provider type. The attributes field can have multiple  keys in its value. Each of those keys has a value  object that includes the type, and display name as keys. For example, `{type:"", display_name:""}`. 

* `created_at` - (String) The time when the resource was created.

* `data_type` - (String) The format of the results that a provider supports.

* `icon` - (String) The icon of a provider in .svg format that is encoded as a base64 string.

* `instance_limit` - (Integer) The maximum number of instances that can be created for the provider type.

* `label` - (List) The label that is associated with the provider type.
Nested schema for **label**:
	* `text` - (String) The text of the label.
	* `tip` - (String) The text to be shown when user hover overs the label.

* `mode` - (String) The mode that is used to get results from provider (`PUSH` or `PULL`).

* `updated_at` - (String) The time when the resource was updated.

