---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : floating_ip"
description: |-
  Fetches floating IP information.
---

# ibm_floating_ip
Retrieve an information of VPC floating IP on IBM Cloud as a read-only data source. For more information, about floating IP, see [about floating IP](https://cloud.ibm.com/docs/vpc?topic=vpc-creating-a-vpc-using-the-rest-apis#create-floating-ip-api-tutorial).

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
data "ibm_is_floating_ip" "example" {
  name = "example-floating-ip"
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `name` - (Required, String) The name of the floating IP.

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `address` - (String) The floating IP address that is created.
- `crn` - (String) The CRN for this floating IP.
- `id` - (String) The unique identifier of the floating IP.
- `status` - (String) Provisioning status of the floating IP address.
- `tags` - (String) The tags associated with the floating IP.
- `target` - (String) The unique identifier for the target to allocate the floating IP address.
- `target_list` - (List) The target of this floating IP.
    Nested scheme for **target_list**:
    - `crn` - (String) The CRN if target is a public gateway.
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and provides some supplementary information.
		    
			Nested scheme for **deleted**:
  			- `more_info` - (String) Link to documentation about deleted resources.
    - `href` - (String) The URL for this target.
    - `id` - (String) The unique identifier for this target.
    - `name` - (String) The user-defined name for this target.
    - `primary_ip` - (List) The reserved ip reference.
    
      Nested scheme for **primary_ip**:
        - `address` - (String) The IP address. If the address has not yet been selected, the value will be 0.0.0.0. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
        - `href`- (String) The URL for this reserved IP
        - `name`- (String) The user-defined or system-provided name for this reserved IP
        - `reserved_ip`- (String) The unique identifier for this reserved IP
        - `resource_type`- (String) The resource type.
- `resource_type` - (String) The resource type.
- `zone` - (String) The zone name where to create the floating IP address.
