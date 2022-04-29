---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : reserved_ip"
description: |-
  Shows the information for a reserved IP and subnet.
---

# ibm_is_subnet_reserved_ip
Retrieve information of an existing reserved IP in a subnet as a read only data source. For more information, about associated reserved IP subnet, see [binding and unbinding a reserved IP address](https://cloud.ibm.com/docs/vpc?topic=vpc-bind-unbind-reserved-ip).

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
data "ibm_is_subnet_reserved_ip" "example" {
  subnet      = ibm_is_subnet.example.id
  reserved_ip = ibm_is_subnet_reserved_ip.example.reserved_ip
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `subnet` - (Required, String)The ID for the subnet.
- `reserved_ip` - (Required, String)The ID for the reserved IP.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `auto_delete` -  (String) The auto_delete boolean for reserved IP.
- `created_at` -  (String) The creation timestamp for the reserved IP.
- `href` -  (String) The unique reference for the reserved IP.
- `id` -  (String) The ID for the reserved IP.
- `lifecycle_state` - (String) TThe lifecycle state of the reserved IP. [ deleting, failed, pending, stable, suspended, updating, waiting ]
- `name` -  (String) The name for the reserved IP.
- `owner` -  (String) The owner of the reserved IP.
- `reserved_ip` -  (String) The ID for the reserved IP.
- `resource_type` -  (String) The resource type.
- `subnet` -  (String) The ID of the subnet for the reserved IP.
- `target` - The ID of the target for the reserved IP.
