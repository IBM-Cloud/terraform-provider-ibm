---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : public_gateways"
description: |-
  Manages IBM public gateways.
---

# ibm_is_public_gateways
Retrieve information of an existing public gateways as a read only data source. For more information, about an VPC public gateway, see [about networking](https://cloud.ibm.com/docs/vpc?topic=vpc-about-networking-for-vpc).

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
data "ibm_is_public_gateways" "example"{
}

```

## Argument reference

Review the argument references that you can specify for your data source. 

- `resource_group` - (String) The ID of the Resource group this public gateway belongs to.

## Attribute reference
Review the attribute references that you can access after you retrieve your data source.

- `public_gateways` - (List) List of all Public Gateways in the IBM Cloud infrastructure region.

  Nested scheme for `public_gateways`:
  - `access_tags`  - (List) Access management tags associated for the public gateway.
  - `crn` - (String) The CRN for this public gateway.
  - `id` - (String) The ID of the public gateway.
  - `status` - (String) The status of the public gateway.
  - `vpc` - (String) The VPC ID of the public gateway.
  - `zone` - (String) The public gateway zone name.
  - `tags` - (String) Tags associated with the public gateway.
  - `name` - (String) The name of the public gateway.
  - `floating_ip` - (List) A nested block describing the floating IP of the public gateway.
  
    Nested scheme for `floating_ip`:
    - `id` - (String) ID of the floating ip bound to the public gateway.
    - `address` - (String) IP address of the floating ip bound to the public gateway.
