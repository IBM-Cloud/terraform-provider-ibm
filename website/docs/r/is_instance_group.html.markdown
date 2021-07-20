---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM: instance_group"
description: |-
  Manages IBM VPC instance group.
---

# ibm_is_instance_group

Create, update or delete a instance group on VPC. For more information, about instance group, see [managing an instance group](https://cloud.ibm.com/docs/vpc?topic=vpc-managing-instance-group).

## Example usage
In the following example, you can create a instance group on VPC Generation-2 infrastructure.

```terraform
resource "ibm_is_vpc" "vpc2" {
  name = "vpc2test"
}

resource "ibm_is_subnet" "subnet2" {
  name            = "subnet2"
  vpc             = ibm_is_vpc.vpc2.id
  zone            = "us-south-2"
  ipv4_cidr_block = "10.240.64.0/28"
}

resource "ibm_is_ssh_key" "sshkey" {
  name       = "ssh1"
  public_key = "SSH KEY"
}

resource "ibm_is_instance_template" "instancetemplate1" {
  name    = "testtemplate"
  image   = "r006-14140f94-fcc4-11e9-96e7-a72723715315"
  profile = "bx2-8x32"

  primary_network_interface {
    subnet = ibm_is_subnet.subnet2.id
  }

  vpc  = ibm_is_vpc.vpc2.id
  zone = "us-south-2"
  keys = [ibm_is_ssh_key.sshkey.id]
}

resource "ibm_is_instance_group" "instance_group" {
  name              = "testgroup"
  instance_template = ibm_is_instance_template.instancetemplate1.id
  instance_count    = 2
  subnets           = [ibm_is_subnet.subnet2.id]

  //User can configure timeouts
  timeouts {
    create = "15m"
    delete = "15m"
    update = "10m"
  }
}
```

## Timeouts

The `ibm_is_instance_group` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create**: The creation of the instance group is considered `failed` if no response is received for 15 minutes.
- **delete**: The deletion of the instance group is considered `failed` if no response is received for 15 minutes.
- **update**: The creation of the instance group is considered `failed` if no response is received for 10 minutes. 

## Argument reference
Review the argument references that you can specify for your resource. 

- `application_port` - (Optional, Integer) The instance group uses when scaling up instances to supply the port for the Load Balancer pool member. The `load_balancer` and `load_balancer_pool` arguments must be specified when configured.
- `load_balancer` - (Optional, String) The load Balancer ID, the `application_port` and `load_balancer_pool` arguments must be specified when configured.
- `load_balancer_pool` - (Optional, String) The load Balancer pool ID, the `application_port` and `load_balancer` arguments must be specified when configured.
- `instance_template` - (Required, Forces new resource, String) The ID of the instance template to create the instance group.
- `instance_count` - (Optional, Integer) The number of instances to create in the instance group. **Note** instance group manager must be in diables state to update the `instance_count`.
- `name` - (Required, String) The instance  group name.
- `resource_group` - (Optional, String) The resource group ID.
- `subnets` - (Required, List) The list of subnet IDs used by the instances.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The ID of an instance group.
- `instances` - (String) The number of instances in the instances group.
- `managers` - (String) List of managers associated with the instance group.
- `status` - (String) Status of an instance group.
- `vpc` - (String) The VPC ID.

## Import
The `ibm_is_instance_group` resource can be imported by using the instance group ID.

```
$ terraform import ibm_is_instance_group.instance_group r006-14140f94-fcc4-11e9-96e7-a7272asd122112315
```

