# IBM VPC Next Generation (ng) examples

This example demonstrates how to create VPC Next Generation resources on IBM Cloud using Terraform. The code provides a comprehensive collection of IBM Cloud VPC resource configurations including VPCs, subnets, instances, load balancers, security groups, volumes, and more.

## Prerequisites

- An IBM Cloud account
- Terraform installed (version 0.13 or later)
- IBM Cloud API key

## Configuration

Configure your IBM Cloud provider in `provider.tf` with your API key:

```hcl
variable "ibmcloud_api_key" {
  description = "Enter your IBM Cloud API Key, you can get your IBM Cloud API key using: https://cloud.ibm.com/iam#/apikeys"
}

provider "ibm" {
  ibmcloud_api_key = var.ibmcloud_api_key
  region           = "us-south"
}
```

## Running the example

### Initialize Terraform
```shell
terraform init
```

### Preview changes
```shell
terraform plan
```

### Apply changes
```shell
terraform apply
```

### Destroy resources
```shell
terraform destroy
```

## Resource Types

This example includes configurations for:

### Compute Resources
* Reservations
* Virtual server instances
* Bare metal servers
* Dedicated hosts
* Placement groups
* SSH keys
* Images
* Instance groups
* Instance templates

### Network Resources
* VPCs
* Cluster networks
* Subnets
* Virtual network interfaces
* Access control lists
* Security groups
* Routing tables
* Public gateways
* Floating IPs
* Public address ranges
* Load balancers
* Virtual private endpoint gateways
* Private Path services
* VPNs
* Flow logs

### Storage Resources
* Block storage volumes
* Block storage snapshots
* Backup policies
* File storage shares

### Geography Resources
* Regions
* Zones

## Note

These examples create actual resources in IBM Cloud that may incur costs. Be sure to destroy resources when they're no longer needed.