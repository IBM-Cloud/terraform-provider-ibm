---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM : ibm_container_vpc_bare_metal_worker_reload"
description: |-
  Executes a bare metal worker reload action for an IBM Cloud Kubernetes Service VPC cluster.
---

# ibm_container_vpc_bare_metal_worker_reload

Use the `ibm_container_vpc_bare_metal_worker_reload` action to reload a bare metal worker node in an IBM Cloud Kubernetes Service VPC cluster.

## Example usage

### Invoke an action from the CLI

The following example reloads a bare metal worker node and waits for the operation to complete.

```terraform
action "ibm_container_vpc_bare_metal_worker_reload" "reload" {
  config {
    cluster_name_id      = ibm_container_vpc_cluster.cluster.id
    bare_metal_server_id = data.ibm_container_vpc_cluster.cluster_data.workers[0]
  }
}
```

The following example submits the reload request and returns immediately without waiting for completion.

```terraform
action "ibm_container_vpc_bare_metal_worker_reload" "reload_no_wait" {
  config {
    cluster_name_id      = ibm_container_vpc_cluster.cluster.id
    bare_metal_server_id = data.ibm_container_vpc_cluster.cluster_data.workers[0]
    no_wait              = true
  }
}
```

Invoke the action explicitly by using the `-invoke` flag.

```bash
terraform apply -invoke action.ibm_container_vpc_bare_metal_worker_reload.reload
terraform apply -invoke action.ibm_container_vpc_bare_metal_worker_reload.reload_no_wait
```

## Argument reference

Review the argument references that you can specify for the action configuration.

- `cluster_name_id` - (Required, String) The ID or name of the VPC cluster that contains the bare metal worker node.
- `bare_metal_server_id` - (Required, String) The ID of the bare metal server to reload.
- `timeout` - (Optional, String) The maximum time to wait for the reload operation to complete, such as `30m` or `1h`. If not specified, the default value is `45m`. This argument is ignored when `no_wait` is `true`.
- `no_wait` - (Optional, Boolean) If set to `true`, the action returns immediately after the reload request is submitted without waiting for completion. The default value is `false`.

## Behavior

When invoked, this action performs the following steps:

1. Sends a reload request for the specified bare metal worker node.
2. If `no_wait` is `false`, waits until the worker reaches a terminal state or the timeout is reached.
3. If `no_wait` is `true`, returns immediately after the reload request is accepted.

This action does not return output values.

## Related information

For more information about using Terraform actions, see the [HashiCorp documentation](https://developer.hashicorp.com/terraform/language/invoke-actions).
