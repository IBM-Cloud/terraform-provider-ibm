# IBM Cloud VPC Instance Group Example

This example demonstrates how to create a VPC instance group with autoscaling capabilities on IBM Cloud. Instance groups allow you to manage multiple virtual server instances with uniform configurations and provide automatic scaling based on metrics or schedules.

## Supported Resources

* [Instance Group](https://cloud.ibm.com/docs/vpc?topic=vpc-creating-auto-scale-instance-group)
* [Instance Group Manager](https://cloud.ibm.com/docs/vpc?topic=vpc-creating-auto-scale-instance-group#creating-instance-group-manager)
* [Instance Group Manager Policy](https://cloud.ibm.com/docs/vpc?topic=vpc-creating-auto-scale-instance-group#creating-instance-group-manager-policies)
* [Instance Group Manager Action](https://cloud.ibm.com/docs/vpc?topic=vpc-creating-auto-scale-instance-group#creating-instance-group-scheduled-actions)

## Terraform Compatibility

* Terraform 0.12 or later (for current branch - `master`)

## Usage

To run this example, execute:

```bash
terraform init
terraform plan
terraform apply
```

To remove the created resources:

```bash
terraform destroy
```

## Implementation Details

This example creates:

1. Base infrastructure:
   - A VPC
   - A subnet within the VPC
   - An SSH key for instance access

2. Instance scaling infrastructure:
   - An instance template defining the configuration for all instances
   - An instance group based on the template
   - Two instance group managers:
     - An autoscale manager for adjusting instance count based on metrics
     - A scheduled manager for time-based scaling
   - A CPU-based scaling policy
   - A scheduled action for specific time-based scaling

### Autoscaling Explained

The autoscaling system works through these components:
- The **Instance Group** maintains a collection of identical instances
- The **Instance Group Manager** controls the scaling behavior
- The **Manager Policy** defines when to scale based on metrics (e.g., CPU usage)
- The **Manager Action** defines scheduled scaling operations

When CPU usage exceeds the defined threshold, the autoscale manager will automatically add instances up to the defined maximum. When usage decreases, it will reduce instances to the defined minimum after the cooldown period.

## Example Configuration

```hcl
# Create a VPC for the instance group
resource "ibm_is_vpc" "vpc2" {
  name = var.vpc_name
}

# Create a subnet within the VPC
resource "ibm_is_subnet" "subnet2" {
  name            = var.subnet_name
  vpc             = ibm_is_vpc.vpc2.id
  zone            = var.zone
  ipv4_cidr_block = "10.240.64.0/28"
}

# Create an instance template for the instance group
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

# Create an instance group using the template
resource "ibm_is_instance_group" "instance_group" {
  name              = var.instance_group_name
  instance_template = ibm_is_instance_template.instancetemplate1.id
  instance_count    = var.instance_count
  subnets           = [ibm_is_subnet.subnet2.id]
}

# Create an autoscale manager for the instance group
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

# Create a CPU-based scaling policy
resource "ibm_is_instance_group_manager_policy" "cpuPolicy" {
  instance_group         = ibm_is_instance_group.instance_group.id
  instance_group_manager = ibm_is_instance_group_manager.instance_group_manager.manager_id
  metric_type            = "cpu"
  metric_value           = var.metric_value
  policy_type            = "target"
  name                   = var.policy_name
}
```

## Additional Resources

* [IBM Cloud VPC Instance Groups Documentation](https://cloud.ibm.com/docs/vpc?topic=vpc-creating-auto-scale-instance-group)
* [IBM Terraform Provider Examples](https://github.com/Mavrickk3/terraform-provider-ibm/tree/master/examples/ibm-is-instance-group)

## Input Parameters

| Name | Description | Type | Required |
|------|-------------|------|---------|
| vpc_name | Name for the VPC | `string` | no |
| subnet_name | Name for the subnet | `string` | no |
| ssh_key | SSH RSA Public key for instance access | `string` | yes |
| ssh_key_name | Name for the SSH key | `string` | no |
| template_name | Name for the instance template | `string` | no |
| image_id | Image ID for instances | `string` | yes |
| profile | Instance profile type | `string` | no |
| zone | VPC Zone for resources | `string` | no |
| instance_group_name | Name for the instance group | `string` | no |
| instance_count | Initial number of instances | `integer` | no |
| instance_group_manager_name | Name for the autoscale manager | `string` | no |
| aggregation_window | Time window in seconds to aggregate metrics | `integer` | no |
| cooldown | Duration in seconds to pause scaling after an action | `integer` | no |
| manager_type | Type of instance group manager | `string` | no |
| enable_manager | Enable/disable autoscaling behavior | `bool` | no |
| max_membership_count | Maximum number of instances | `integer` | yes |
| min_membership_count | Minimum number of instances | `integer` | yes |
| policy_name | Name for the scaling policy | `string` | no |
| metric_value | Target metric value (e.g., CPU percentage) | `integer` | no |
| cron_spec | Cron specification for scheduled actions | `string` | no |

## Notes

* Instance groups require a VPC, subnet, and SSH key
* The instance template defines the configuration for all instances in the group
* Two types of managers are supported: autoscale and scheduled
* Policies can monitor CPU, memory, network_in, or network_out metrics
* Scheduled actions use cron syntax and override the autoscale manager temporarily