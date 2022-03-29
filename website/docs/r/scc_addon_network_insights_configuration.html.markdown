---
layout: "ibm"
page_title: "IBM : ibm_scc_addon_network_insights_configuration"
description: |-
  Manages scc_addon_network_insights_configuration.
subcategory: "Security and Compliance Center"
---

# ibm_scc_addon_network_insights_configuration

Provides a resource for scc_addon_network_insights_configuration. This allows scc_addon_network_insights_configuration to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_scc_addon_network_insights_configuration" "scc_addon_network_insights_configuration" {
  region_id = "us"
  status = "enable"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `account_id` - (Optional, Forces new resource, String) Account ID is optional, if not provided value will be inferred from the token retrieved from the IBM Cloud API key.
* `region_id` - (Required, String) Region id for example - us.
* `status` - (Required, String) Enable or Disable.
  * Constraints: Allowable values are: `enable`, `disable`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the scc_addon_network_insights_configuration.
* `addon` - Type of addon.

## Import

You can import the `ibm_scc_addon_network_insights_configuration` resource by using `network_insights_config_id`.
The `network_insights_config_id` property can be formed from 

```
<account_id>/network-insights
```

* `account_id` - A string. AccountID from the resource has to be imported.

# Syntax
```
$ terraform import ibm_scc_addon_network_insights_configuration.scc_addon_network_insights_configuration 
```

# THIS RESOURCE DOES NOT SUPPORT DELETION
