# IBM Cloud VPC Flow Logs Example

This example demonstrates how to create Flow Logs for IBM Cloud VPC resources. Flow logs capture network traffic information for analysis, security auditing, and troubleshooting.

## Supported Resources

* [Flow Logs](https://cloud.ibm.com/docs/vpc?topic=vpc-flow-logs)

## Terraform Compatibility

* Terraform 0.12 or later (for current branch - `master`)
* For Terraform 0.11 compatibility, use branch `terraform_v0.11.x`

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

1. A Cloud Object Storage instance to store the flow logs
2. A bucket within that instance
3. A Flow Log collector targeting a VPC instance
4. Required permissions between the collector and the bucket

### How Flow Logs Work

VPC Flow Logs capture network traffic information (IP addresses, protocols, ports) going to and from network interfaces within your VPC. The data is sent to a specified Cloud Object Storage bucket for later retrieval and analysis.

## Example Configuration

```hcl
# Get resource group for Cloud Object Storage
data "ibm_resource_group" "cos_group" {
  name = var.resource_group
}

# Get information about an existing instance
data "ibm_is_instance" "ds_instance" {
  name = "vpc1-instance"
}

# Create a Cloud Object Storage service instance
resource "ibm_resource_instance" "instance1" {
  name              = "cos-instance"
  resource_group_id = data.ibm_resource_group.cos_group.id
  service           = "cloud-object-storage"
  plan              = "standard"
  location          = "global"
}

# Create a bucket for storing flow logs
resource "ibm_cos_bucket" "bucket1" {
   bucket_name          = "us-south-bucket-vpc1"
   resource_instance_id = ibm_resource_instance.instance1.id
   region_location = var.region
   storage_class = "standard"
}

# Create a flow log collector for an instance
resource "ibm_is_flow_log" "test_flowlog" {
  depends_on = [ibm_cos_bucket.bucket1]
  name = "test-instance-flow-log"
  target = data.ibm_is_instance.ds_instance.id
  active = true
  storage_bucket = ibm_cos_bucket.bucket1.bucket_name
}
```

## Additional Resources

* [IBM Cloud Flow Logs Documentation](https://cloud.ibm.com/docs/vpc?topic=vpc-flow-logs)
* [IBM Terraform Provider Examples](https://github.com/IBM-Cloud/terraform-provider-ibm/tree/master/examples/ibm-is-flow-log)

## Input Parameters

| Name | Description | Type | Required |
|------|-------------|------|---------|
| name | The unique user-defined name for the flow log collector | `string` | yes |
| target | The ID of the target to collect flow logs for | `string` | yes |
| storage\_bucket | The name of the Cloud Object Storage bucket for log storage | `string` | yes |
| active | Whether this collector is active (default: true) | `boolean` | no |
| resource\_group | The resource group ID for the flow log | `string` | no |

## Notes

* An IAM service authorization must be in place to grant VPC Flow Logs services write access to the bucket
* Flow logs can be created for instances, subnets, or entire VPCs
* More specific flow log collectors take precedence over less specific ones