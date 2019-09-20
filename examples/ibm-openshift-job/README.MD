# terraform_cluster_job
This example shows how to run a shell script in IBM Openshift cluster job using terraform. 
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

It will mount the shell script(wrapper.sh) as a config map into your cluster and execute it in the job/pod