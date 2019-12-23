# Auto Scale Group example

This example shows how to launch a cluster using Auto Scaling Groups and Local load balancer.

This creates a cluster with the ability to automate the manual scaling process associated with adding or removing virtual servers to support your business applications.
It distribute processing and communications evenly across multiple servers within a data center so that a single device does not carry the entire load.
The cluster installs nginx and it listens on port 80.

To run, configure your IBM Cloud provider

Running the example

* Pass the public key while running terraform.

For planning phase

```shell
terraform plan -var 'ssh_public_key=<public_key_value>'
```

For apply phase

```shell
terraform apply -var 'ssh_public_key=<public_key_value>'
```

To remove the stack wait for few minutes and test the stack by launching a browser with cluster url.

```shell
terraform destroy -var 'ssh_public_key=<public_key_value>'
```
