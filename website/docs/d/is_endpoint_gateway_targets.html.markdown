---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : Endpoint Gateway Targets"
description: |-
  Fetches endpoint gateway targets information.
---

# ibm_is_endpoint_gateway_targets
Retrieve an information of an endpoint gateway targets on IBM Cloud as a read-only data source. For more information, about VPC endpoint gateway target, see [creating an endpoint gateway](https://cloud.ibm.com/docs/vpc?topic=vpc-ordering-endpoint-gateway).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage

```terraform
data "ibm_is_endpoint_gateway_targets" "example" {
}
```

## Attribute reference
You can access the following attribute references after your data source is created. 
- `resources` -  (List) Collection of resources to be set as endpoint gateway target. Nested `resources` blocks have the following structure.

  Nested scheme for `resources`:
  - `crn` - (String) The CRN for the specific object.	
  - `endpoint_type` - (String) The data endpoint type of the offering.
  - `full_qualified_domain_names` - (String) The fully qualified domain names.
  - `name` - (String) The display name in the requested language.
  - `parent` - (String) The parent for the specific object. 
  - `resource_type` - (String) The resource type of the offering. 
  - `service_location` - (String) The service location of the offering.
