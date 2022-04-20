---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : floating_ips"
description: |-
  Fetches floating IPs information.
---

# ibm_is_floating_ips

Retrieve an information of VPC floating IPs on IBM Cloud. For more information, about floating IP, see [about floating IP](https://cloud.ibm.com/docs/vpc?topic=vpc-creating-a-vpc-using-the-rest-apis#create-floating-ip-api-tutorial).

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
data "ibm_is_floating_ips" "example" {
  name = "example-floating-ips"
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `name` - (Optional, String) The unique user-defined name for this floating IP.

## Attribute reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the Floating IPs Collection.
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
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and provides some supplementary information.
		    
			Nested scheme for **deleted**:
  			- `more_info` - (String) Link to documentation about deleted resources.
    	- `href` - (String) The URL for this network interface.
		- `id` - (String) The unique identifier for this network interface.
		- `name` - (String) The user-defined name for this network interface.
		- `primary_ip` - (List) The reserved ip reference.
		
			Nested scheme for **primary_ip**:
			- `address` - (String) The IP address. If the address has not yet been selected, the value will be 0.0.0.0. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
			- `href` - (String) The URL for this reserved IP
			- `name` - (String) The user-defined or system-provided name for this reserved IP
			- `reserved_ip` - (String) The unique identifier for this reserved IP
			- `resource_type`- (String) The resource type.		
		- `primary_ipv4_address` - (String) The primary IPv4 address. If the address has not yet been selected, the value will be `0.0.0.0`. **Same as primary_ip.0.address**
		- `resource_type` - (String) The resource type.
	- `zone` - (List) The zone this floating IP resides in.
	    
		Nested scheme for **zone**:
		- `href` - (String) The URL for this zone.
		- `name` - (String) The globally unique name for this zone.
