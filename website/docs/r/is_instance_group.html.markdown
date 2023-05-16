---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM: instance_group"
description: |-
  Manages IBM VPC instance group.
---

# ibm_is_instance_group

Create, update or delete a instance group on VPC. For more information, about instance group, see [managing an instance group](https://cloud.ibm.com/docs/vpc?topic=vpc-managing-instance-group).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage
In the following example, you can create a instance group on VPC Generation-2 infrastructure.

```terraform
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}

resource "ibm_is_subnet" "example" {
  name            = "example-subnet"
  vpc             = ibm_is_vpc.example.id
  zone            = "us-south-2"
  ipv4_cidr_block = "10.240.64.0/28"
}

resource "ibm_is_ssh_key" "example" {
  name       = "example-ssh"
  public_key = "SSH KEY"
}

resource "ibm_is_instance_template" "example" {
  name    = "example-template"
  image   = ibm_is_image.example.id
  profile = "bx2-8x32"

  primary_network_interface {
    subnet = ibm_is_subnet.example.id
  }

  vpc  = ibm_is_vpc.example.id
  zone = "us-south-2"
  keys = [ibm_is_ssh_key.example.id]
}

resource "ibm_is_instance_group" "example" {
  name              = "example-group"
  instance_template = ibm_is_instance_template.example.id
  instance_count    = 2
  subnets           = [ibm_is_subnet.example.id]

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

- `access_tags`  - (Optional, List of Strings) A list of access management tags to attach to the instance group.

  ~> **Note:** 
  **&#x2022;** You can attach only those access tags that already exists.</br>
  **&#x2022;** For more information, about creating access tags, see [working with tags](https://cloud.ibm.com/docs/account?topic=account-tag&interface=ui#create-access-console).</br>
  **&#x2022;** You must have the access listed in the [Granting users access to tag resources](https://cloud.ibm.com/docs/account?topic=account-access) for `access_tags`</br>
  **&#x2022;** `access_tags` must be in the format `key:value`.
- `application_port` - (Optional, Integer) The instance group uses when scaling up instances to supply the port for the Load Balancer pool member. The `load_balancer` and `load_balancer_pool` arguments must be specified when configured.
- `load_balancer` - (Optional, String) The load Balancer ID, the `application_port` and `load_balancer_pool` arguments must be specified when configured.
- `load_balancer_pool` - (Optional, String) The load Balancer pool ID, the `application_port` and `load_balancer` arguments must be specified when configured.
- `instance_template` - (Required, Forces new resource, String) The ID of the instance template to create the instance group.
- `instance_count` - (Optional, Integer) The number of instances to create in the instance group. 
  
  ~>**Note:** instance group manager must be in diables state to update the `instance_count`.
- `name` - (Required, String) The instance  group name.
- `resource_group` - (Optional, String) The resource group ID.
- `subnets` - (Required, List) The list of subnet IDs used by the instances.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `crn` - (String) The CRN for this instance group.
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

