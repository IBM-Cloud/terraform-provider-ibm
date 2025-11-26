---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM: instance_group_manager_policy"
description: |-
  Manages IBM VPC instance group manager policy.
---

# ibm_is_instance_group_manager_policy

Create, update or delete a policy of an instance group manager. For more information, about instance group manager policy, see [creating an instance group for auto scaling](https://cloud.ibm.com/docs/vpc?topic=vpc-creating-auto-scale-instance-group).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage
In the following example, you can create a policy for instance group manager.

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
  public_key = "SSH_KEY"
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
  name              = "example-instance-group"
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

resource "ibm_is_instance_group_manager" "example" {
  name                 = "example-instance-group-manager"
  aggregation_window   = 120
  instance_group       = ibm_is_instance_group.example.id
  cooldown             = 300
  manager_type         = "autoscale"
  enable_manager       = true
  max_membership_count = 2
  min_membership_count = 1
}

resource "ibm_is_instance_group_manager_policy" "example" {
  instance_group         = ibm_is_instance_group.example.id
  instance_group_manager = ibm_is_instance_group_manager.example.manager_id
  metric_type            = "cpu"
  metric_value           = 70
  policy_type            = "target"
  name                   = "example-ig-manager-policy"
}

```

## Argument reference
Review the argument references that you can specify for your resource. 

- `instance_group` - (Required, String) The instance group ID.
- `instance_group_manager` - (Required, String) The instance group manager ID for policy creation.
- `policy_type` - (Required, String) The type of metric to evaluate.
- `metric_type` - (Required, String) The type of metric to evaluate. The possible values for metric types are `cpu`, `memory`, `network_in`, and `network_out`.
- `metric_value`- (Required, Integer) The metric value to evaluate.
- `name` - (Optional, String) The name of the policy.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The ID in the combination of instance group ID, instance group manager ID, and instance group manager policy ID.
- `policy_id` - (String) The policy ID.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the `ibm_is_instance_group_manager_policy` resource by using `id`.
The `id` property can be formed from `instance group ID`, `instance group manager ID`, and `instance group manager policy ID`. For example:

```terraform
import {
  to = ibm_is_instance_group_manager_policy.policy
  id = "<instance_group_id>/<instance_group_manager_id>/<instance_group_manager_policy_id>"
}
```

Using `terraform import`. For example:

```console
% terraform import ibm_is_instance_group_manager_policy.policy <instance_group_id>/<instance_group_manager_id>/<instance_group_manager_policy_id>
```