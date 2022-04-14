---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : load balancer"
description: |-
  Manages IBM load balancer.
---

# ibm_is_lb
Create, update, or delete a VPC Load Balancer. For more information, about VPC load balancer, see [load balancers for VPC overview](https://cloud.ibm.com/docs/vpc?topic=vpc-nlb-vs-elb).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```


## Example usage
An example to create an application load balancer.

```terraform
resource "ibm_is_lb" "example" {
  name    = "example-load-balancer"
  subnets = [ibm_is_subnet.example.id]
}

```

An example to create a network load balancer.

```terraform
resource "ibm_is_lb" "example" {
  name    = "example-load-balancer"
  subnets = [ibm_is_subnet.example.id]
  profile = "network-fixed"
}

```

## Timeouts
The `ibm_is_lb` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 30 minutes) Used for creating Instance.
- **delete** - (Default 30 minutes) Used for deleting Instance.


## Argument reference
Review the argument references that you can specify for your resource. 

- `logging`- (Optional, Bool) Enable or disable datapath logging for the load balancer. This is applicable only for application load balancer. Supported values are **true** or **false**. Default value is **false**.
- `name` - (Required, String) The name of the VPC load balancer.
- `profile` - (Optional, Forces new resource, String) For a Network Load Balancer, this attribute is required and should be set to `network-fixed`. For Application Load Balancer, profile is not a required attribute.
- `resource_group` - (Optional, Forces new resource, String) The resource group where the load balancer to be created.
- `route_mode` - (Optional, Forces new resource, Bool) Indicates whether route mode is enabled for this load balancer.

  ~> **NOTE:** Currently, `route_mode` enabled is supported only by private network load balancers.
- `security_groups`  (Optional, List) A list of security groups to use for this load balancer. This option is supported only for application load balancers.
- `subnets` - (Required, Forces new resource, List) List of the subnets IDs to connect to the load balancer.
- `tags` (Optional, Array of Strings) A list of tags that you want to add to your load balancer. Tags can help you find the load balancer more easily later.
- `type` - (Optional, Forces new resource, String) The type of the load balancer. Default value is `public`. Supported values are `public` and `private`.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `crn` - (String) The CRN for this load balancer.
- `hostname` - (String) The fully qualified domain name assigned to this load balancer.
- `id` - (String) The unique identifier of the load balancer.
- `operating_status` - (String) The operating status of this load balancer.
- `public_ips` - (String) The public IP addresses assigned to this load balancer.
- `private_ip` - (List) The Reserved IP address reference assigned to this load balancer.

  Nested scheme for `private_ip`:
    - `address` - (String) IPv4 The IP address. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
    - `href` - (String) The URL for this reserved ip
    - `reserved_ip`- (String) The unique identifier for this reserved IP.
    - `name`- (String) The user-defined or system-provided name for this reserved IP

- `private_ips` - (String) The private IP addresses (Reserved IP address reference) assigned to this load balancer.
- `status` - (String) The status of the load balancer.
- `security_groups_supported`- (Bool) Indicates if this load balancer supports security groups.
- `udp_supported`- (Bool) Indicates whether this load balancer supports UDP.


## Import
The `ibm_is_lb` resource can be imported by using the load balancer ID. 

**Syntax**

```
$ terraform import ibm_is_lb.example <lb_ID>
```

**Example**

```
$ terraform import ibm_is_lb.example d7bec597-4726-451f-8a63-e62e6f133332c
``` 
