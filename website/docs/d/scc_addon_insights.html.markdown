---
layout: "ibm"
page_title: "IBM : ibm_scc_addon_insights"
description: |-
  Get information about scc_addon_insights
subcategory: "Security and Compliance Center"
---

# ibm_scc_addon_insights

Provides a read-only data source for scc_addon_insights. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_scc_addon_insights" "scc_addon_insights" {
}
```


## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the scc_addon_insights.
* `type` - (Optional, List) 
  * Constraints: Allowable list items are: `network-insights`, `activity-insights`.

