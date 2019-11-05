# IBM Cluster example

This example shows how to create a Kubernetes Cluster under a specified resource group id, with a default worker pool with 2 workers, edit the default worker pool to add a new zone to it, add a worker pool with different zone with 2 workers and binds a service instance to a cluster.

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
