# IBM Cluster example

This example shows how to create a Kubernetes VPC Cluster under a specified resource group id, with default worker node with given zone and subnets. 
To have a multizone cluster, update the zones with new zone-name and subnet-id.
It also creates a additional worker pool on different zone. After successfull creation of cluster it binds the cloud-object-storage service.
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
