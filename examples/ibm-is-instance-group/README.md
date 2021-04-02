# Example for VPC instance group resources

This example shows how to create VPC instance group, instance group manager and isnatcne group manager policy resources to  achieve autoscale feature.

Following types of resources are supported:

* [Instance Group](https://cloud.ibm.com/docs/terraform)


## Terraform versions

Terraform 0.12. Pin module version to `~> v1.12.0`. Branch - `master`.

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## Example Usage

Create a VPC Instance group with autoscale feature:

```hcl
resource "ibm_is_vpc" "vpc2" {
  name = var.vpc_name
}

resource "ibm_is_subnet" "subnet2" {
  name            = var.subnet_name
  vpc             = ibm_is_vpc.vpc2.id
  zone            = var.zone
  ipv4_cidr_block = "10.240.64.0/28"
}

resource "ibm_is_ssh_key" "sshkey" {
  name       = var.ssh_key_name
  public_key = var.ssh_key
}

resource "ibm_is_instance_template" "instancetemplate1" {
  name    = var.template_name
  image   = var.image_id
  profile = var.profile

  primary_network_interface {
    subnet = ibm_is_subnet.subnet2.id
  }

  vpc  = ibm_is_vpc.vpc2.id
  zone = var.zone
  keys = [ibm_is_ssh_key.sshkey.id]
}

resource "ibm_is_instance_group" "instance_group" {
  name              = var.instance_group_name
  instance_template = ibm_is_instance_template.instancetemplate1.id
  instance_count    = var.instance_count
  subnets           = [ibm_is_subnet.subnet2.id]
}

resource "ibm_is_instance_group_manager" "instance_group_manager" {
  name                 = var.instance_group_manager_name
  aggregation_window   = var.aggregation_window
  instance_group       = ibm_is_instance_group.instance_group.id
  cooldown             = var.cooldown
  manager_type         = var.manager_type
  enable_manager       = var.enable_manager
  max_membership_count = var.max_membership_count
  min_membership_count = var.min_membership_count
}

resource "ibm_is_instance_group_manager_policy" "cpuPolicy" {
  instance_group         = ibm_is_instance_group.instance_group.id
  instance_group_manager = ibm_is_instance_group_manager.instance_group_manager.manager_id
  metric_type            = "cpu"
  metric_value           = var.metric_value
  policy_type            = "target"
  name                   = var.policy_name
}

data "ibm_is_instance_group" "instance_group_data" {
  name = ibm_is_instance_group.instance_group.name
}

data "ibm_is_instance_group_manager" "instance_group_manager" {
  instance_group = ibm_is_instance_group_manager.instance_group_manager.instance_group
  name           = ibm_is_instance_group_manager.instance_group_manager.name
}

data "ibm_is_instance_group_manager_policy" "instance_group_manager_policy" {
  instance_group         = ibm_is_instance_group_manager_policy.cpuPolicy.instance_group
  instance_group_manager = ibm_is_instance_group_manager_policy.cpuPolicy.instance_group_manager
  name                   = ibm_is_instance_group_manager_policy.cpuPolicy.name
}
```

## Examples

* [ Instance Group ](https://github.com/IBM-Cloud/terraform-provider-ibm/tree/master/examples/ibm-is-instance-group)

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | n/a |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| vpc_name | The unique user-defined name for the vpc. | `string` | no |
| subnet\_name | The unique user-defined name for the vpc subnet. | `string` | no |
| ssh\_key | The SSH RSA Public key to access the instances. | `string` | yes |
| ssh\_key\_name | The name of the SSH key. | `string` | no |
| template\_name | The instance template name to create instance template. | `string` | no |
| image\_id | Image identifier to create the instance template. | `string` | yes |
| profile | Instance profile type. | `string` | no |
| zone | VPC Zone name where instance template is created. | `string` | no |
| instance\_group\_name | Name of the vpc instance group. | `string` | no |
| instance\_count | The number of instances managed in the instance group. | `integer` | no |
| instance\_group\_manager\_name | The manager name under instance group. | `string` | no |
| aggregation\_window | The time window in seconds to aggregate metrics prior to evaluation | `integer` | no |
| cooldown | The duration of time in seconds to pause further scale actions after scaling has taken place | `integer` | no |
| manager\_type | The type of instance group manager | `string` | no |
| enable\_manager | enable or disable the autoscale behavior of instance group. | `bool` | no |
| max\_membership\_count | The upper threshold value for autoscaling. Based on the metrics collected, the instances scaled to this maximum number. | `integer` | yes |
| min\_membership\_count | The lower threshold value set to instance group manager to scale the nubner of instances to least value. | `integer` | yes |
| policy\_name | The instance group manager's policy name. | `string` | no |
| metric\_value | Metric value to be set to evaluated by instance group manager. | `integer` | no |

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
