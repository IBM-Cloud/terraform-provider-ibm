---
layout: "ibm"
page_title: "IBM : ibm_pdr_validate_clustertype"
description: |-
  Get information about pdr_validate_clustertype
subcategory: "DrAutomation Service"
---

# ibm_pdr_validate_clustertype

Provides a read-only data source to retrieve information about a pdr_validate_clustertype. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_pdr_validate_clustertype" "pdr_validate_clustertype" {
	instance_id = "crn:v1:staging:public:power-dr-automation:global:a/a123456fb04ceebfb4a9fd38c22334455:123456d3-1122-3344-b67d-4389b44b7bf9::"
	orchestrator_cluster_type = "on-premises"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `accept_language` - (Optional, String) The language requested for the return document.
* `if_none_match` - (Optional, String) ETag for conditional requests (optional).
* `instance_id` - (Required, Forces new resource, String) instance id of instance to provision.
* `orchestrator_cluster_type` - (Required, String) orchestrator cluster type value.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the pdr_validate_clustertype.
* `description` - (String) 
* `status` - (String) 

