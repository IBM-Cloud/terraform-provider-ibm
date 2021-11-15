---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : floating_ips"
description: |-
  Fetches floating IPs information.
---

# ibm_is_floating_ips

Retrieve an information of VPC floating IPs on IBM Cloud as a read-only data source. For more information, about floating IP, see [about floating IP](https://cloud.ibm.com/docs/vpc?topic=vpc-creating-a-vpc-using-the-rest-apis#create-floating-ip-api-tutorial).

## Example Usage

```hcl
data "ibm_is_floating_ips" "is_floating_ips" {
	name = "my-floating-ip"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

- `name` - (Optional, String) The unique user-defined name for this floating IP.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the FloatingIPCollection.
- `floating_ips` - (List) Collection of floating IPs.
  
    Nested scheme for **floating_ips**:
    - `address` - (String) The globally unique IP address.
    - `created_at` - (String) The date and time that the floating IP was created.
    - `crn` - (String) The CRN for this floating IP.
    - `href` - (String) The URL for this floating IP.
    - `id` - (String) The unique identifier for this floating IP.
    - `name` - (String) The unique user-defined name for this floating IP.
    - `resource_group` - (List) The resource group for this floating IP.
	    Nested scheme for **resource_group**:
      	- `href` - (String) The URL for this resource group.
		- `id` - (String) The unique identifier for this resource group.
		- `name` - (String) The user-defined name for this resource group.
	- `status` - (String) The status of the floating IP.
	- `target` - (List) The target of this floating IP.
	    Nested scheme for **target**:
		- `crn` - (String) The CRN for this public gateway.
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
		    Nested scheme for **deleted**:
  			- `more_info` - (String) Link to documentation about deleted resources.
    	- `href` - (String) The URL for this network interface.
		- `id` - (String) The unique identifier for this network interface.
		- `name` - (String) The user-defined name for this network interface.
		- `primary_ipv4_address` - (String) The primary IPv4 address.If the address has not yet been selected, the value will be `0.0.0.0`.
		- `resource_type` - (String) The resource type.
	- `zone` - (List) The zone this floating IP resides in.
	    Nested scheme for **zone**:
  		- `href` - (String) The URL for this zone.
		- `name` - (String) The globally unique name for this zone.

