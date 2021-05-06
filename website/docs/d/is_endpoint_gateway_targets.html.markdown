---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Endpoint Gateway Targets"
description: |-
  Fetches endpoint gateway targets information.
---

# ibm\is_endpoint_gateway_targets

Import the details of endpoint gateway targets on IBM Cloud as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl

    data "ibm_is_endpoint_gateway_targets" "endpointGatewayTargets" {
    }

```

## Attribute Reference

The following attributes are exported:
* `crn` - The crn for this specific object.	
* `name` - Display name in the requested language.
* `parent` - The parent for this specific object. 
* `endpoint_type` - Data endpoint type of this offering.
* `full_qualified_domain_names` - Fully qualified domain names.
* `resource_type` - Resource type of this offering. 
* `service_location` - Service location of this offering.
