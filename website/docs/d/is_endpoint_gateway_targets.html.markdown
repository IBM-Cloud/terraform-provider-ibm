---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Endpoint Gateway Targets"
description: |-
  Fetches endpoint gateway targets information.
---

# ibm_is_endpoint_gateway_targets
Retrieve an information of an endpoint gateway targets on IBM Cloud as a read-only data source. For more information, about VPC endpoint gateway target, see [creating an endpoint gateway](https://cloud.ibm.com/docs/vpc?topic=vpc-ordering-endpoint-gateway).

## Example usage

```terraform

    data "ibm_is_endpoint_gateway_targets" "endpointGatewayTargets" {
    }

```

## Attribute reference
You can access the following attribute references after your data source is created. 

- `crn` - (String) The CRN for this specific object.	
- `endpoint_type` - (String) The data endpoint type of this offering.
- `full_qualified_domain_names` - (String) The fully qualified domain names.
- `name` - (String) The display name in the requested language.
- `parent` - (String) The parent for this specific object. 
- `resource_type` - (String) The resource type of this offering. 
- `service_location` - (String) The service location of this offering.
