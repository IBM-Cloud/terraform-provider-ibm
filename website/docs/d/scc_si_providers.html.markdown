---
layout: "ibm"
subcategory: "Security and Compliance Center"
page_title: "IBM : ibm_scc_si_providers"
description: |-
  Get information about scc_si_providers
---

# DEPRECATED
Security and Compliance Center - Security Insights has now deprecated, backend services are no longer available. The docs will be removed in next release.

# ibm_scc_si_providers

Provides a read-only data source for scc_si_providers. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example usage

```terraform
data "ibm_scc_si_providers" "providers" {
  limit = 4
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

* `limit` - (Optional, String) Limit the number of the returned documents to the specified number.
* `skip` - (Optional, String) The offset is the index of the item from which you want to start returning data from. Default is 0.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the scc_si_providers.
* `limit` - (Optional, Integer) The number of elements returned in the current instance. The default is 200.

* `providers` - (Optional, List) The providers requested.
Nested scheme for **providers**:
	* `name` - (Required, String) The name of the provider in the form '{account_id}/providers/{provider_id}'.
	* `id` - (Required, String) The ID of the provider.

* `skip` - (Optional, Integer) The offset is the index of the item from which you want to start returning data from. The default is 0.

* `total_count` - (Optional, Integer) The total number of providers available.

