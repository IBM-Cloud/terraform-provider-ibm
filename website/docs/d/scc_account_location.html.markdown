---
layout: "ibm"
subcategory: "Security and Compliance Center"
page_title: "IBM : ibm_scc_account_location"
description: |-
  Get information about scc_account_location
---

# ibm_scc_account_location

Provides a read-only data source for scc_account_location. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example usage

```terraform
data "ibm_scc_account_location" "scc_account_location" {
	location_id = "us"
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

* `location_id` - (Required, Forces new resource, String) The programatic ID of the location that you want to work in.
  * Constraints: Allowable values are: `us`, `eu`, `uk`.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the scc_account_location.

* `analytics_endpoint_url` - (String) The endpoint that is used to generate analytics for the Posture Management component.

* `compliance_endpoint_url` - (String) The endpoint that is used to call the Posture Management APIs.

* `governance_endpoint_url` - (String) The endpoint that is used to call the Configuration Governance APIs.

* `id` - (String) The programatic ID of the location that you want to work in.
  * Constraints: Allowable values are: `us`, `eu`, `uk`.

* `main_endpoint_url` - (String) The base URL for the service.

* `regions` - (List)
Nested scheme for **regions**:
	* `id` - (Required, String) The programatic ID of the available regions.
	  * Constraints: Allowable values are: `us`, `eu`, `uk`.

* `results_endpoint_url` - (String) The endpoint that is used to get the results for the Configuration Governance component.

* `si_endpoint_url` - (String) The endpoint that is used to call the Security Insights APIs.
