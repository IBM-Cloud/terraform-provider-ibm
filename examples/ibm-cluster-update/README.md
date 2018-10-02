# IBM Cluster example

This example shows how to update the existing cluster and worker to the available latest version. In this example we are using the kubectl drain and uncordon to
prepare the workers for maintenance.
The given node will be marked unschedulable to prevent new pods from arriving. Then drain deletes all pods except mirror pods (which cannot be deleted through the API server). If there are DaemonSet-managed pods, drain will not proceed without –ignore-daemonsets, and regardless it will not delete any DaemonSet-managed pods, because those pods would be immediately replaced by the DaemonSet controller, which ignores unschedulable markings. If there are any pods that are neither mirror pods nor managed–by ReplicationController, DaemonSet or Job–, then drain will not delete any pods unless you use –force.
Once the clusters and workers are ready put the workers back into service, useing kubectl uncordon, which will make the workers schedulable again.

Please refer to https://console.bluemix.net/docs/containers/cs_cluster_update.html#master for overall process on updating cluster and workers version.

To run, configure your IBM Cloud provider

Running the example

For planning phase

```shell
terraform plan
```

For apply phase

```shell
terraform apply
```

For destroy

```shell
terraform destroy
```
