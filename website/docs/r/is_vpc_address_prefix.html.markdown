---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : vpc-address-prefix"
description: |-
  Manages IBM IS VPC address prefix.
---

# ibm_is_vpc_address_prefix
Create, update, or delete an IP address prefix. For more information, about IS VPC address prefix, see [address prefixes](https://cloud.ibm.com/docs/vpc?topic=vpc-vpc-behind-the-curtain#address-prefixes).

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
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_vpc_address_prefix" "example" {
  name = "example-address-prefix"
  zone = "us-south-1"
  vpc  = ibm_is_vpc.example.id
  cidr = "10.240.0.0/24"
}

```


## Argument reference
Review the argument references that you can specify for your resource. 

- `cidr` - (Required, Forces new resource, String) The CIDR block for the address prefix.
- `is_default` - (Optional, Boolean) Makes the prefix as default prefix for this zone in this VPC. Default is `false`
- `name` - (Required, String) The address prefix name.No.
- `vpc` - (Required, Forces new resource, String) The VPC ID.
- `zone` - (Required, Forces new resource, String) The name of the zone.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The ID of the address prefix. The ID is composed of `<vpc_id>/<address_prefix_id>`.
- `has_subnets`- (Bool) Indicates whether subnets exist with addresses from this prefix.
- `address_prefix` - (String) the unique identifier of the address prefix.
- `related_crn` - (String) CRN of the VPC this address prefix belongs to.

## Import
The `ibm_is_vpc_address_prefix` resource can be imported by using the VPC ID and VPC address prefix ID.

**Syntax**

```
$ terraform import ibm_is_vpc_address_prefix.example <vpc_ID>/<address_prefix_ID>
```
