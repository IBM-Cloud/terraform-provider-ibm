# Cluster with citrix load balancer example

This example shows how to launch a cluster using virtual servers and citrix load balancer

This sample configuration will stand up a cluster of nodes, running on virtual guests, behind a global load balancer. The URL of the cluster is provided as an output.

Please Note: If Netscaler VPX 10.5 is used, terraform uses Netscaler's REST API([NITRO API](https://docs.citrix.com/en-us/netscaler/11/nitro-api.html)) for ibmcloud_infra_lb_vpx_vip resource management. The NITRO API is only accessable in SoftLayer private network, so it is necessary to execute terraform in SoftLayer's private network when you deploy Netscaler VPX 10.5 devices. SoftLayer [SSL VPN](http://www.softlayer.com/VPN-Access) also can be used for private network connection.

To run, configure your IBM Cloud provider

Running the example

* Pass the public key while running terraform.

For planning phase

```shell
terraform plan -var 'ssh_public_key=<public_key_value>'
```

For apply phase

```shell
terraform apply -var 'ssh_public_key=<public_key_value'
```

To remove the stack wait for few minutes and test the stack by launching a browser with cluster URL.

```shell
terraform destroy -var 'ssh_public_key=<public_key_value>'
```
