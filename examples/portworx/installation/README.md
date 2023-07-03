# IBM Schematics example

This example shows how to deploy portworx in a Kubernetes cluster for SDS worker nodes . 
You need to have a Baremetal with sds cluster ready in iks as it comes up with one or more raw unformatted storage.
If you want to use portworx on non sds cluster you need to manually add raw, unformatted, and unmounted block storage and attach it to a volume. 
Docs- https://cloud.ibm.com/docs/containers?topic=containers-portworx#portworx_database
You also need a database for etcd service for portworx metadata. We will use kubernetes secrets for storing the secrets.
For creating a cluster go to https://cloud.ibm.com/kubernetes
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
