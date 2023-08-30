---

subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_vpc_worker"
description: |-
  Manages IBM container VPC worker.
---

# ibm_container_vpc_worker

Replace a worker. The worker will be replaced & updated to the latest patch in the specified cluster. For more information, about VPC worker updates, see [Updating VPC worker nodes](https://cloud.ibm.com/docs/containers?topic=containers-update&interface=ui#vpc_worker_node)


## Example usage
In the following example, you can replace worker in a vpc cluster:

```terraform
resource "ibm_container_vpc_worker" "test_worker" {
    cluster_name        = "my_vpc_cluster"
    replace_worker      = "kube-clusterid-mycluster-default-00001"
    resource_group_id   = "6015365a-9d93-4bb4-8248-79ae0db2dc21"
    kube_config_path    = "my_vpc_cluster.yaml"
    check_ptx_status    = "true"
    ptx_timeout         = "5m"
}
```

## Timeouts

The `ibm_container_vpc_worker` provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **Create** The creation of the worker is considered failed when no response is received for 30 minutes. 
- **Delete** The deletion of the worker is considered failed when no response is received for 30 minutes. 

## Argument reference
Review the argument references that you can specify for your resource. 

- `cluster_name` - (Required, Forces new resource, String) The name or ID of the cluster.
- `replace_worker` - (Required, Forces new resource, String) The ID of the worker that needs to be replaced.
- `resource_group_id` - (Optional, Forces new resource, String) The ID of the resource group. To retrieve the ID, run `ibmcloud resource groups` or use the `ibm_resource_group` data source. If no value is provided, the `default` resource group is used.
- `check_ptx_status` - (Optional, String) Boolean value to check the status of Portworx on the replaced worker instance. By default, this variable is set as `false`.
- `kube_config_path` - (Optional, String) The Cluster config with absolute path. If `check_ptx_status` is true, this variable should hold a valid value. To retrieve the cluster config, run `ibmcloud cluster config -c <Cluster_ID>` or use the `ibm_container_cluster_config` data source.
- `ptx_timeout` - (Optional, String) The Status of Portworx on the replaced worker is considered failed when no response is received for 15 minutes.
- `sds` - (Optional, String) Software Defined Storage (SDS) parameter performs worker replace based on the installed SDS solution in the cluster. Supported value `ODF`
- `sds_timeout` - (Optional, String) The Status of the Software Defined Storage on the replaced worker is considered failed when no response is received for 30 minutes.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the worker.
- `ip` - (String) The IP of the worker.

## Note
- This resource is different from all other resource of IBM Cloud. Worker replace has 2 operations, i.e. Delete old worker & Create a new worker. On `terraform apply`, Replace operation is being handled where both the deletion & creation happens whereas on the `terraform destroy`, only the state is cleared but not the actual resource.
- When the worker list is being provided as inputs, the list must be user generated and should not be passed from the `ibm_container_cluster` data source.
- If `terraform apply` fails during worker replace or while checking the portworx status, perform any one of the following actions before retrying.
  - Resolve the issue manually and perform `terraform untaint` to proceed with the subsequent workers in the list.
  - If worker replace is still needed, update the input list by replacing the existing worker id with the new worker id.
- The `sds` option is currently in development. To perform Worker Replace for `ODF`, you can test and utilise it. Please ignore the parameter otherwise.