---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : load balancer"
description: |-
  Manages IBM load balancer.
---

# ibm_is_lb
Create, update, or delete a VPC Load Balancer. For more information, about VPC load balancer, see [load balancers for VPC overview](https://cloud.ibm.com/docs/vpc?topic=vpc-nlb-vs-elb).


## Example usage
An example to create an application load balancer.

```terraform
resource "ibm_is_lb" "lb" {
  name    = "loadbalancer1"
  subnets = ["04813493-15d6-4150-9948-6cc646cb67f2"]
}

```

An example to create a network load balancer.

```terraform
resource "ibm_is_lb" "lb" {
  name    = "loadbalancer1"
  subnets = ["04813493-15d6-4150-9948-6cc646cb67f2"]
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
- `profile` - (Required, Forces new resource, String) The profile to use for this load balancer. Supported value is `network-fixed`.
- `resource_group` - (Optional, Forces new resource, String) The resource group where the load balancer to be created.
- `security_groups`  (Optional, List) A list of security groups to use for this load balancer. This option is supported only for application load balancers.
- `subnets` - (Required, List) List of the subnets IDs to connect to the load balancer.
- `tags` (Optional, Array of Strings) A list of tags that you want to add to your load balancer. Tags can help you find the load balancer more easily later.
- `type` - (Optional, Forces new resource, String) The type of the load balancer. Default value is `public`. Supported values are `public` and `private`.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `hostname` - (String) The fully qualified domain name assigned to this load balancer.
- `id` - (String) The unique identifier of the load balancer.
- `operating_status` - (String) The operating status of this load balancer.
- `public_ips` - (String) The public IP addresses assigned to this load balancer.
- `private_ips` - (String) The private IP addresses assigned to this load balancer.
- `status` - (String) The status of the load balancer.
- `security_groups_supported`- (Bool) Indicates if this load balancer supports security groups.


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
