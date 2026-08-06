# IBM Cloud VPC Bare Metal Worker Reload Example

This example demonstrates how to reload a bare metal worker node in an IBM Cloud VPC Kubernetes cluster using the `ibm_container_vpc_bare_metal_worker_reload` action.

## Usage

To run this example, execute the following commands:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

### Example Usage

Create a bare metal cluster:

```hcl

resource "random_id" "name" {
  byte_length = 2
}

locals {
  ZONE = "${var.region}-1"
}

resource "ibm_is_vpc" "vpc" {
  name = "vpc-${random_id.name.hex}"
}

resource "ibm_is_subnet" "subnet" {
  name                     = "subnet-${random_id.name.hex}"
  vpc                      = ibm_is_vpc.vpc.id
  zone                     = local.ZONE
  total_ipv4_address_count = 256
}

data "ibm_resource_group" "resource_group" {
  name = var.resource_group
}

resource "ibm_container_vpc_cluster" "cluster" {
  name              = "${var.cluster_name}-${random_id.name.hex}"
  vpc_id            = ibm_is_vpc.vpc.id
  flavor            = var.flavor
  worker_count      = 1
  resource_group_id = data.ibm_resource_group.resource_group.id
  wait_till         = "OneWorkerNodeReady"

  zones {
    subnet_id = ibm_is_subnet.subnet.id
    name      = local.ZONE
  }
}

data "ibm_container_vpc_cluster" "cluster_data" {
  name              = ibm_container_vpc_cluster.cluster.name
  resource_group_id = data.ibm_resource_group.resource_group.id
}

data "ibm_container_cluster_config" "cluster_config" {
  cluster_name_id = ibm_container_vpc_cluster.cluster.id
}
```

```hcl
# Example 1: Reload bare metal worker and wait for completion
action "ibm_container_vpc_bare_metal_worker_reload" "reload" {
  config {
    cluster_name_id      = ibm_container_vpc_cluster.cluster.id
    bare_metal_server_id = data.ibm_container_vpc_cluster.cluster_data.workers[0]
    timeout              = "1h"
  }
}

# Example 2: Reload bare metal worker without waiting (fire-and-forget)
action "ibm_container_vpc_bare_metal_worker_reload" "reload_no_wait" {
  config {
    cluster_name_id      = ibm_container_vpc_cluster.cluster.id
    bare_metal_server_id = data.ibm_container_vpc_cluster.cluster_data.workers[0]
    no_wait              = true
  }
}
```

Actions are invoked explicitly using the `-invoke` flag:

```bash
# Invoke the reload action and wait for completion
terraform apply -invoke action.ibm_container_vpc_bare_metal_worker_reload.reload

# Invoke the no-wait action (fire-and-forget)
terraform apply -invoke action.ibm_container_vpc_bare_metal_worker_reload.reload_no_wait
```

Run `terraform destroy` when you don't need these resources.

## Behavior

When invoked, this action:
1. Triggers a reload operation on the specified bare metal worker
2. If `no_wait` is false (default):
   - Waits until worker reaches one of the target states ("deployed", "deploy_failed", or "reloading_failed") or timeout
   - Reports success or failure via Terraform diagnostics
3. If `no_wait` is true:
   - Returns immediately after triggering the reload
   - Worker reload continues in the background

## Examples

* [VPC bare metal worker reload](https://github.com/IBM-Cloud/terraform-provider-ibm/tree/master/examples/ibm-cluster/vpc-baremetal-worker-reload)

## Requirements

| Name | Version |
|------|---------|
| terraform | >=1.14.0, <2.0 |

## Providers

| Name | Version |
|------|---------|
| ibm  | latest |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| name | Name of the cluster. | `string` | yes |
| flavor | The flavor of the VPC worker node that you want to use. | `string` | yes |
| worker_count | The number of worker nodes per zone in the default worker pool. Default value `1`.| `integer` | no |
| zone | Name of the zone.| `string` | yes |
| resource_group | Name of the resource group.| `string` | yes |
{: caption="inputs"}

### Action Inputs

| Name | Description | Type | Required | Default |
|------|-------------|------|---------|---------|
| cluster_name_id | The name or ID of the VPC cluster containing the worker. | `string` | yes | - |
| bare_metal_server_id | The ID of the bare metal server to reload (same as worker ID for bare metal workers). | `string` | yes | - |
| timeout | Maximum time to wait for the reload operation (e.g., "1h", "30m"). | `string` | no | "45m" |
| no_wait | If true, returns immediately after triggering the reload without waiting for completion. | `bool` | no | false |

## Outputs

| Name | Description |
|------|-------------|
| cluster_config_file_path | Path to the cluster config file for kubectl access |