---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : reserved_ips"
description: |-
  Lists all the information in reserved IP for subnet.
---

# ibm_is_subnet_reserved_ips
Retrieve information about a reserved IP in a subnet as a read only data source. For more information, about associated reserved IP subnet, see [binding and unbinding a reserved IP address](https://cloud.ibm.com/docs/vpc?topic=vpc-bind-unbind-reserved-ip).

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
data "ibm_is_subnet_reserved_ips" "example" {
  subnet = ibm_is_subnet.example.id
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `subnet` - (Required, String) The ID for the subnet.

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `reserved_ips` -  (List) The collection of all the reserved IPs in the subnet.

  Nested scheme for `reserved_ips`:
  - `address` -  (String) The IP bound for the reserved IP.
  - `auto_delete` -  (String) If reserved IP shall be deleted automatically.
  - `created_at` -  (String) The date and time that the reserved IP was created.
  - `href` -  (String) The URL for this reserved IP.
  - `lifecycle_state` - (String) The lifecycle state of the reserved IP. [ deleting, failed, pending, stable, suspended, updating, waiting ]
  - `reserved_ip` -  (String) The unique identifier for this reserved IP.
  - `name` -  (String) The user defined or system provided name for this reserved IP.
  - `owner` -  (String) The owner of a reserved IP, defining whether it is managed by the user or the provider.
  - `resource_type` -  (String) The resource type.
  - `target` - (String) The ID of the target for the reserved IP.
  - `target_crn` - (String) The crn of the target for the reserved IP.
  - `target_reference` - (List) The target this reserved IP is bound to. If absent, this reserved IP is provider-owned or unbound.
	  
      Nested schema for **target_reference**:
      - `crn` - (String) The CRN for this endpoint gateway.
      - `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.

          Nested schema for **deleted**:
          - `more_info` - (String) Link to documentation about deleted resources.
      - `href` - (String) The URL for this endpoint gateway.
      - `id` - (String) The unique identifier for this endpoint gateway.
      - `name` - (String) The name for this endpoint gateway. The name is unique across all endpoint gateways in the VPC.
      - `resource_type` - (String) The resource type.
- `total_count` -  (String) The total number of resources across all pages.