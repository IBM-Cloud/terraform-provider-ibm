# terraform_cluster_job
This example shows how to run a shell script in IBM Openshift cluster job using terraform. This terraform template will create a new config map containing the wrapper.sh file's content and mount this config map to a volume. Then it will create  a new job/pod and execute the config map(shell script) which has been mounted, additionally it will create a secret (optional for custom images)
To run, configure your IBM Cloud provider and input your cluster id when asked.

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